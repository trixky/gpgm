package checker

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/trixky/krpsim/checker/core"
	"github.com/trixky/krpsim/checker/parser"
)

type LocalExpectedStock struct {
	Product          string
	Quantity         int
	AvailableAtCycle int
}

type TokenWithQuantity struct {
	Name     string
	Quantity int
}

func parseOutputLine(sm core.InitialContext, line string) (products []TokenWithQuantity, err error) {
	pairs := strings.Split(line, ";")
	for _, pair := range pairs {
		infos := strings.Split(pair, ":")
		if len(infos) != 2 {
			return products, fmt.Errorf("invalid token `%s`", pair)
		}
		quantity, err := strconv.Atoi(infos[1])
		if err != nil {
			return products, err
		}
		products = append(products, TokenWithQuantity{
			Name:     infos[0],
			Quantity: quantity,
		})
	}
	return products, nil
}

func CheckOutput(simulationFile string, outputFile string) (res bool, err error) {
	sm, err := parser.ParseSimulationFile(simulationFile)
	if err != nil {
		return false, err
	}

	// Format:
	// cycle: process:amount;[process:amount;...]
	// stock: product:amount;[product:amount;...]
	shouldEnd := false
	simulationStock := sm.Stock.DeepCopy()
	expectedStock := []LocalExpectedStock{}
	finalStock := core.Stock{}
	lastCycle := -1
	lines := strings.Split(outputFile, "\n")
	for _, line := range lines {
		// Ignore empty lines
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		// Extra content
		if shouldEnd {
			return false, fmt.Errorf("unexpected line after final stock")
		}

		// Parse final stock
		if strings.HasPrefix(line, "stock: ") {
			products, err := parseOutputLine(sm, line[7:])
			if err != nil {
				return false, err
			}
			for _, product := range products {
				if !sm.Stock.ResourceExists(product.Name) {
					return false, fmt.Errorf("invalid inexisting product `%s`", product.Name)
				}
				finalStock[product.Name] = product.Quantity
			}
			shouldEnd = true
		} else { // Parse each cycle processes
			// Split by the cycle separator
			parts := strings.Split(line, ": ")
			if len(parts) != 2 {
				return false, fmt.Errorf("invalid cycle line `%s`", line)
			}

			// Check cycle validity
			cycle, err := strconv.Atoi(parts[0])
			if err != nil {
				return false, err
			}
			if cycle < 0 {
				return false, fmt.Errorf("invalid cycle `%s`", parts[0])
			}
			if cycle < lastCycle {
				return false, fmt.Errorf("duplicate or unordered cycle `%s` (%d)", parts[0], lastCycle)
			}
			lastCycle = cycle

			// Update stock to add all expected stocks for the cycle
			remaining := []LocalExpectedStock{}
			for _, e := range expectedStock {
				if e.AvailableAtCycle <= cycle {
					simulationStock.AddResource(e.Product, e.Quantity)
				} else {
					remaining = append(remaining, e)
				}
			}
			expectedStock = remaining

			// Check if each processes are valid and can be executed
			processes, err := parseOutputLine(sm, parts[1])
			if err != nil {
				return false, err
			}
			for _, processToken := range processes {
				var process *core.Process
				for _, cProcess := range sm.Processes {
					if cProcess.Name == processToken.Name {
						process = &cProcess
						break
					}
				}
				if process == nil {
					return false, fmt.Errorf("invalid process `%s`", processToken.Name)
				}

				// Execute process and update stock
				if !process.CanBeExecutedXTimes(simulationStock, processToken.Quantity) {
					return false, fmt.Errorf("can't execute process `%s` on cycle %d", processToken.Name, cycle)
				}
				for name, quantity := range process.Inputs {
					simulationStock.RemoveResource(name, quantity*processToken.Quantity)
				}
				for name, quantity := range process.Outputs {
					expectedStock = append(expectedStock, LocalExpectedStock{
						Product:          name,
						Quantity:         quantity * processToken.Quantity,
						AvailableAtCycle: cycle + process.Delay,
					})
				}
			}
		}
	}
	if !shouldEnd {
		return false, fmt.Errorf("missing final stock in output")
	}

	// Finish the remaining processes
	for _, e := range expectedStock {
		simulationStock.AddResource(e.Product, e.Quantity)
	}

	// Compare final stock with the simulation stock
	for name, quantity := range finalStock {
		if simulationStock.GetResource(name) != quantity {
			return false, fmt.Errorf("invalid final stock for a product, expected %d for `%s` but got %d", simulationStock.GetResource(name), name, quantity)
		}
	}

	return true, nil
}
