package onesub

import "math"

func oneSub(v int) int {
	if v == 0 {
		return 1
	}
	allOnes := v&(v+1) == 0
	if allOnes {
		return int(math.Log2(float64(v+1))) + 1
	}
	var max, running, sinceZero int

	var usedZero bool
	comparer := 1

	for comparer <= v {
		isOne := (comparer & v) != 0
		if isOne {
			running++
			if running > max {
				max = running
			}
			if usedZero {
				sinceZero++
			}
		} else {
			if !usedZero {
				usedZero = true
				running++
				if running > max {
					max = running
				}
			} else {
				running = sinceZero + 1
				sinceZero = 0
			}
		}
		comparer <<= 1
	}
	return max
}
