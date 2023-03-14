package parser

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/trixky/krpsim/algo/core"
)

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

func parseStock(line string) (product core.Product, err error) {
	product.Quantity = -1

	for i := 0; i < len(line); i++ {
		// Peek next token
		token, newOffset := peekToken(line, i)
		i = newOffset

		if product.Name != "" && product.Quantity >= 0 {
			return product, fmt.Errorf("too many product informations")
		}

		// Use token
		if token == ":" && product.Name == "" {
			return product, fmt.Errorf("missing product name")
		}
		if token == ":" {
			continue
		}
		if product.Name == "" {
			product.Name = token
		} else if product.Quantity < 0 {
			quantity, err := strconv.Atoi(token)
			if err != nil {
				return product, err
			}
			if quantity < 0 {
				return product, fmt.Errorf("invalid stock quantity `%d`", quantity)
			}
			product.Quantity = quantity
		} else {
			return product, fmt.Errorf("too many product informations")
		}
	}

	if product.Name == "" || product.Quantity < 0 {
		return product, fmt.Errorf("missing product name or quantity")
	}

	return product, nil
}

func parseProcess(line string) (process core.Process, err error) {
	process.Inputs = make(map[string]int)
	process.Outputs = make(map[string]int)

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
			return process, fmt.Errorf("too many process informations")
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
			return process, fmt.Errorf("expected separator `:` but got `%s`", token)
		}

		// Inputs
		if step == 2 && token == ":" {
			if !insideParenthesis {
				step = 3
				continue
			} else if currentIdentifier == "" {
				return process, fmt.Errorf("missing input product name")
			} else {
				continue // expect quantity next
			}
		}
		if step == 2 {
			if token == "(" {
				if insideParenthesis {
					return process, fmt.Errorf("unexpected `(` inside of parenthesis")
				}
				insideParenthesis = true
				continue
			}
			if token == ")" {
				if !insideParenthesis || closed || currentIdentifier != "" {
					return process, fmt.Errorf("unexpected `)` token outside of parenthesis")
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
				if err != nil {
					return process, err
				}
				if quantity < 0 {
					return process, fmt.Errorf("invalid input product quantity `%d`", quantity)
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
				return process, fmt.Errorf("missing output product name")
			} else {
				continue // expect quantity next
			}
		}
		if step == 3 {
			if token == "(" {
				if insideParenthesis {
					return process, fmt.Errorf("unexpected `(` inside of parenthesis")
				}
				insideParenthesis = true
				continue
			}
			if token == ")" {
				if !insideParenthesis || closed || currentIdentifier != "" {
					return process, fmt.Errorf("unexpected `)` token outside of parenthesis")
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
				if err != nil {
					return process, err
				}
				if quantity < 0 {
					return process, fmt.Errorf("invalid output product quantity `%d`", quantity)
				}
				process.Outputs[currentIdentifier] = quantity
				currentIdentifier = ""
				continue
			}
		}

		// Delay
		if step == 4 {
			delay, err := strconv.Atoi(token)
			if err != nil {
				return process, err
			}
			if delay < 0 {
				return process, fmt.Errorf("invalid process delay `%d`", delay)
			}
			process.Delay = delay
			step = 5
			continue
		}
	}

	if step < 5 {
		return process, fmt.Errorf("missing process informations")
	}

	return process, nil
}

func parseOptimize(line string) (optimizeFor map[string]bool, err error) {
	optimizeFor = make(map[string]bool)

	step := 0
	shouldEnd := false
	leftToken := ""
	expectTwo := -1
	rightToken := ""

	for i := 0; i < len(line); i++ {
		// Peek next token
		token, newOffset := peekToken(line, i)
		i = newOffset

		if shouldEnd {
			return optimizeFor, fmt.Errorf("too many informations in optimize line")
		}

		// First part
		if token == "optimize" {
			step = 1
			continue
		}
		if step < 1 {
			return optimizeFor, fmt.Errorf("expected `optimize` token but got `%s`", token)
		}

		// Separator
		if step == 1 && token == ":" {
			step = 2
			continue
		}
		if step < 2 {
			return optimizeFor, fmt.Errorf("expected `:` token but got `%s`", token)
		}

		// Separator
		if step == 2 && token == "(" {
			step = 3
			continue
		}
		if step < 3 {
			return optimizeFor, fmt.Errorf("expected `(` token but got `%s`", token)
		}

		// Optimize for
		if token == ")" {
			if leftToken != "" && rightToken != "" {
				if rightToken == "time" {
					return optimizeFor, fmt.Errorf("unexpected `time` optimize in right side of optimization")
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
				return optimizeFor, fmt.Errorf("missing left side of optimization")
			}
			if expectTwo > 0 && rightToken == "" {
				return optimizeFor, fmt.Errorf("missing right side of optimization")
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
				return optimizeFor, fmt.Errorf("unexpected `|` token in already separated optimization")
			}
			if leftToken == "" {
				return optimizeFor, fmt.Errorf("missing left side of optimization")
			}
			expectTwo = 1
			continue
		} else {
			if rightToken != "" {
				return optimizeFor, fmt.Errorf("too many informations in optimization")
			}
			if leftToken == "" {
				leftToken = token
			} else {
				if token == "time" {
					return optimizeFor, fmt.Errorf("unexpected `time` optimize in right side of optimization")
				}
				if leftToken != "time" {
					return optimizeFor, fmt.Errorf("unexpected `%s` optimize in left side of optimization but expected `time`", leftToken)
				}
				rightToken = token
			}
		}
	}

	if step < 3 || !shouldEnd {
		return optimizeFor, fmt.Errorf("missing informations in optimize line")
	}
	if len(optimizeFor) == 0 {
		return optimizeFor, fmt.Errorf("nothing to optimize for")
	}

	return optimizeFor, nil
}

func ParseSimulationFile(input string) (sm core.InitialContext, err error) {
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

		asStock, stockErr := parseStock(line)
		if stockErr == nil {
			sm.Stock[asStock.Name] = asStock.Quantity
			continue
		}

		asProcess, processErr := parseProcess(line)
		if processErr == nil {
			for key := range asProcess.Inputs {
				if !sm.Stock.Exists(key) {
					sm.Stock[key] = 0
				}
			}
			for key := range asProcess.Outputs {
				if !sm.Stock.Exists(key) {
					sm.Stock[key] = 0
				}
			}
			sm.Processes = append(sm.Processes, asProcess)
			continue
		}

		asOptimize, optimizeErr := parseOptimize(line)
		if optimizeErr == nil && len(asOptimize) > 0 {
			for product := range asOptimize {
				if product == "time" {
					continue
				} else if !sm.IsInOutput(product) && !sm.Stock.Exists(product) {
					return sm, fmt.Errorf("unexpected optimize for `%s`, not in any process output", product)
				}
			}
			sm.Optimize = asOptimize
			continue
		}

		errString := processErr.Error()
		if strings.HasPrefix(strings.TrimSpace(line), "optimize") {
			errString = optimizeErr.Error()
		} else if len(strings.Split(line, ":")) == 2 {
			errString = stockErr.Error()
		}
		return sm, fmt.Errorf("invalid line %d: %s: %s", line_number+1, line, errString)
	}

	if sm.Stock == nil {
		return sm, fmt.Errorf("invalid Stock")
	}
	if sm.Processes == nil {
		return sm, fmt.Errorf("invalid Processes")
	} else if len(sm.Processes) == 0 {
		return sm, fmt.Errorf("no Processes")
	}
	if sm.Optimize == nil {
		return sm, fmt.Errorf("invalid Optimize")
	} else if len(sm.Optimize) == 0 {
		return sm, fmt.Errorf("nothing to optimize for")
	}

	return sm, nil
}
