package day07

type Equation struct {
	Result int64
	Values []int64
}

type Handler func(int64, *int64, *int64)

func (e *Equation) Validate(operators int64, handler Handler) bool {
	maximum := intPow(operators, int64(len(e.Values)-1))

	// State is a sequence of operators ("+, +, +" = 0 or "+, +, *" = 1 or ...)
	var state int64
	for state = range maximum {
		result := e.Values[0]

		// Go through every value
		for idx := int64(1); idx < int64(len(e.Values)); idx++ {
			current := e.Values[idx]
			// Extract operator for current position from state
			op := (state / intPow(operators, idx-1)) % operators

			handler(op, &result, &current)
		}

		// If equation satisfied, return valid
		if result == e.Result {
			return true
		}
	}

	return false
}

func intPow(base, expo int64) int64 {
	if expo == 0 {
		return 1
	} else if expo == 1 {
		return base
	}

	result := base

	var iter int64
	for iter = 2; iter <= expo; iter++ {
		result *= base
	}
	return result
}
