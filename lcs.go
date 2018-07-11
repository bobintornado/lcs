package lcs

type ResultComparator interface {
	Evaluate([]string, []string) float64
}

type Comparator struct{}

func (r Comparator) Evaluate(originResult, newResult []string) float64 {
	if len(originResult) == 0 || len(newResult) == 0 {
		return 0.0
	}

	maxlength := LcsLength(originResult, newResult)
	return float64(maxlength) / float64(len(originResult))
}

func MaxInt(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func LcsLength(a []string, b []string) int {
	aL := len(a)
	bL := len(b)

	var shortSlice []string
	var longSlice []string

	if aL < bL {
		shortSlice = a
		longSlice = b
	} else {
		shortSlice = b
		longSlice = a
	}

	preCount := 0
	start := 0
	shortEnd := len(shortSlice) - 1
	longEnd := len(longSlice) - 1

	// trim off the matching items at the beginning
	for start <= shortEnd && start <= longEnd && shortSlice[start] == longSlice[start] {
		start = start + 1
		preCount++
	}

	// trim off the matching items at the end
	for start <= shortEnd && start <= longEnd && shortSlice[shortEnd] == longSlice[longEnd] {
		shortEnd = shortEnd - 1
		longEnd = longEnd - 1
		preCount++
	}

	// prefix and suffix match all
	if shortEnd < 0 || longEnd < 0 {
		return preCount
	}

	compareLength := shortEnd - start + 1
	var previousVal int
	currentRow := make([]int, compareLength, compareLength)

	for idxl := start; idxl <= longEnd; idxl++ {
		previousVal = 0
		for idxs := start; idxs <= shortEnd; idxs++ {
			comparePos := idxs - start
			backup := currentRow[comparePos]

			if longSlice[idxl] == shortSlice[idxs] {
				currentRow[comparePos] = 1 + previousVal
			} else {
				if comparePos != 0 {
					currentRow[comparePos] = MaxInt(currentRow[comparePos-1], currentRow[comparePos])
				}
			}

			previousVal = backup
		}
	}

	return currentRow[compareLength-1] + preCount
}
