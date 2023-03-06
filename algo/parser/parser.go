package parser

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Product struct {
	Name     string
	Quantity int
}

type Process struct {
	Name    string
	Inputs  map[string]int
	Outputs map[string]int
	Delay   int
}

type SimulationParameters struct {
	Stock     map[string]int
	Processes []Process
	Optimize  map[string]bool
}

func (sm *SimulationParameters) hasStock(product string) bool {
	_, exists := sm.Stock[product]
	return exists
}

func (sm *SimulationParameters) canExecuteProcess(process Process) bool {
	for product, quantity := range process.Inputs {
		if sm.Stock[product] < quantity {
			return false
		}
	}
	return true
}

func (sm *SimulationParameters) isInOutput(product string) bool {
	for _, process := range sm.Processes {
		for outputProduct := range process.Outputs {
			if product == outputProduct {
				return true
			}
		}
	}
	return false
}

var lineStartWithComment = regexp.MustCompile(`^\s*#`)
var lineEndWithComment = regexp.MustCompile(`(.+)\s*#`)

func peekToken(line string, offset int) (string, int) {
	token := ""
	j := offset

	if j > len(line) {
		return "", -1
	}

	for ; j < len(line); j++ {
		if line[j] == ':' || line[j] == ';' || line[j] == '(' || line[j] == ')' || line[j] == '|' {
			if len(token) == 0 {
				token = token + string(line[j])
			} else {
				j -= 1
			}
			break
		}
		token = token + string(line[j])
	}

	return token, j
}

// TODO return Product, err ?
func parseStock(line string) *Product {
	product := Product{
		Name:     "",
		Quantity: -1,
	}

	for i := 0; i < len(line); i++ {
		// Peek next token
		token, newOffset := peekToken(line, i)
		i = newOffset

		if product.Name != "" && product.Quantity >= 0 {
			return nil
		}

		// Use token
		if token == ":" && product.Name == "" {
			return nil
		}
		if token == ":" {
			continue
		}
		if product.Name == "" {
			product.Name = token
		} else if product.Quantity < 0 {
			quantity, err := strconv.Atoi(token)
			if err != nil || quantity < 0 {
				return nil
			}
			product.Quantity = quantity
		} else {
			return nil
		}
	}

	if product.Name == "" || product.Quantity < 0 {
		return nil
	}

	return &product
}

// TODO return Process, err ?
func parseProcess(line string) *Process {
	process := Process{
		Name:    "",
		Inputs:  make(map[string]int),
		Outputs: make(map[string]int),
		Delay:   0,
	}
	currentIdentifier := ""
	closed := false
	step := 0
	insideParenthesis := false

	for i := 0; i < len(line); i++ {
		// Peek next token
		token, newOffset := peekToken(line, i)
		i = newOffset

		// Too much tokens
		if step == 5 {
			return nil
		}

		// Process name
		if step == 0 {
			process.Name = token
			step = 1
			continue
		}

		// Separator
		if step == 1 && token == ":" {
			step = 2
			continue
		}
		if step < 2 {
			return nil
		}

		// Inputs
		if step == 2 && token == ":" {
			if !insideParenthesis {
				step = 3
				continue
			} else if currentIdentifier == "" {
				return nil
			} else {
				continue // expect quantity next
			}
		}
		if step == 2 {
			if token == "(" {
				if insideParenthesis {
					return nil
				}
				insideParenthesis = true
				continue
			}
			if token == ")" {
				if !insideParenthesis || closed || currentIdentifier != "" {
					return nil
				}
				insideParenthesis = false
				continue
			}
			if token == ";" {
				closed = true
				continue
			}
			if currentIdentifier == "" {
				currentIdentifier = token
				closed = false
				continue
			} else {
				quantity, err := strconv.Atoi(token)
				if err != nil || quantity < 0 {
					return nil
				}
				process.Inputs[currentIdentifier] = quantity
				currentIdentifier = ""
				continue
			}
		}

		// Outputs
		if step == 3 && token == ":" {
			if !insideParenthesis {
				step = 4
				continue
			} else if currentIdentifier == "" {
				return nil
			} else {
				continue // expect quantity next
			}
		}
		if step == 3 {
			if token == "(" {
				if insideParenthesis {
					return nil
				}
				insideParenthesis = true
				continue
			}
			if token == ")" {
				if !insideParenthesis || closed || currentIdentifier != "" {
					return nil
				}
				insideParenthesis = false
				continue
			}
			if token == ";" {
				closed = true
				continue
			}
			if currentIdentifier == "" {
				currentIdentifier = token
				closed = false
				continue
			} else {
				quantity, err := strconv.Atoi(token)
				if err != nil || quantity < 0 {
					return nil
				}
				process.Outputs[currentIdentifier] = quantity
				currentIdentifier = ""
				continue
			}
		}

		// Delay
		if step == 4 {
			delay, err := strconv.Atoi(token)
			if err != nil || delay < 0 {
				return nil
			}
			process.Delay = delay
			step = 5
			continue
		}
	}

	if step < 5 {
		return nil
	}

	return &process
}

// TODO return []string, err ?
func parseOptimize(line string) *map[string]bool {
	step := 0
	shouldEnd := false
	leftToken := ""
	expectTwo := -1
	rightToken := ""
	optimizeFor := make(map[string]bool)

	for i := 0; i < len(line); i++ {
		// Peek next token
		token, newOffset := peekToken(line, i)
		i = newOffset

		if shouldEnd {
			return nil
		}

		// First part
		if token == "optimize" {
			step = 1
			continue
		}
		if step < 1 {
			return nil
		}

		// Separator
		if step == 1 && token == ":" {
			step = 2
			continue
		}
		if step < 2 {
			return nil
		}

		// Separator
		if step == 2 && token == "(" {
			step = 3
			continue
		}
		if step < 3 {
			return nil
		}

		// Optimize for
		if token == ")" {
			if leftToken != "" && rightToken != "" {
				if rightToken == "time" {
					return nil
				}
				optimizeFor[rightToken] = true
			} else {
				optimizeFor[leftToken] = false
			}
			shouldEnd = true
			continue
		}
		if token == ";" {
			if leftToken == "" {
				return nil
			}
			if expectTwo > 0 && rightToken == "" {
				return nil
			}
			if leftToken != "" && rightToken != "" {
				optimizeFor[rightToken] = true
			} else {
				optimizeFor[leftToken] = false
			}
			leftToken = ""
			rightToken = ""
			expectTwo = -1
			continue
		} else if token == "|" {
			if expectTwo > 0 {
				return nil
			}
			if leftToken == "" {
				return nil
			}
			expectTwo = 1
			continue
		} else {
			if rightToken != "" {
				return nil
			}
			if leftToken == "" {
				leftToken = token
			} else {
				if token == "time" {
					return nil
				}
				if leftToken != "time" {
					return nil
				}
				rightToken = token
			}
		}
	}

	if step < 3 || !shouldEnd {
		return nil
	}
	if len(optimizeFor) == 0 {
		return nil
	}

	return &optimizeFor
}

func ParseSimulationFile(input string) (sm SimulationParameters, err error) {
	sm.Stock = make(map[string]int)

	// Split and strip comments
	var lines = strings.Split(input, "\n")
	var filtered []string

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if len(trimmed) == 0 {
			filtered = append(filtered, "")
			continue
		}
		start_with_comment := lineStartWithComment.MatchString(trimmed)
		if !start_with_comment {
			trimmed = strings.TrimSpace(lineEndWithComment.ReplaceAllString(trimmed, "$1"))
			if len(trimmed) == 0 {
				filtered = append(filtered, "")
			} else {
				filtered = append(filtered, trimmed)
			}
		} else {
			filtered = append(filtered, "")
		}
	}

	// Parse each lines and update SimulationParameters
	for line_number, line := range filtered {
		if line == "" {
			continue
		}

		asStock := parseStock(line)
		if asStock != nil {
			sm.Stock[asStock.Name] = asStock.Quantity
			continue
		}

		asProcess := parseProcess(line)
		if asProcess != nil {
			for key := range asProcess.Inputs {
				if !sm.hasStock(key) {
					sm.Stock[key] = 0
				}
			}
			for key := range asProcess.Outputs {
				if !sm.hasStock(key) {
					sm.Stock[key] = 0
				}
			}
			sm.Processes = append(sm.Processes, *asProcess)
			continue
		}

		asOptimize := parseOptimize(line)
		if asOptimize != nil && len(*asOptimize) > 0 {
			for product := range *asOptimize {
				if product == "time" {
					continue
				} else if !sm.isInOutput(product) && !sm.hasStock(product) {
					return sm, fmt.Errorf("parser: Unexpected optimize for %s, not in any process output", product)
				}
			}
			sm.Optimize = *asOptimize
			continue
		}

		return sm, fmt.Errorf("parser: Invalid line %d: %s", line_number+1, line)
	}

	if sm.Stock == nil {
		return sm, fmt.Errorf("parser: Invalid Stock")
	}
	if sm.Processes == nil {
		return sm, fmt.Errorf("parser: Invalid Processes")
	} else if len(sm.Processes) == 0 {
		return sm, fmt.Errorf("parser: No Processes")
	}
	if sm.Optimize == nil {
		return sm, fmt.Errorf("parser: Invalid Optimize")
	} else if len(sm.Optimize) == 0 {
		return sm, fmt.Errorf("parser: Nothing to optimize for")
	}

	return sm, nil
}
