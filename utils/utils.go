package utils

func ModeFilter(sequence *[]string, windowSize int) *[]string {
	smoothed := make([]string, len(*sequence))
	halfWindow := windowSize / 2

	for i := range *sequence {
		start := i - halfWindow
		end := i + halfWindow + 1

		if start < 0 {
			start = 0
			end = windowSize
		}

		if end > len(*sequence) {
			end = len(*sequence)
			start = len(*sequence) - windowSize
		}

		window := (*sequence)[start:end]
		smoothed[i] = findMode(&window)
	}

	return &smoothed
}

func findMode(window *[]string) string {
	freqMap := make(map[string]int)

	for _, val := range *window {
		freqMap[val]++
	}

	var mode string
	maxCount := 0
	for val, count := range freqMap {
		if count > maxCount {
			mode = val
			maxCount = count
		}
	}

	return mode
}

func GroupData(cleanPredictions *[]string, cleanPositions *[]string, time *[]string) (groupedPredictions *[]string, groupedPositions *[]string, groupedTimeIntervals *[][]string) {
	//To be implemented
	return nil, nil, nil
}

func CombineAndTransform(groupedPredictions *[]string, groupedPositions *[]string, groupedTime *[]string) (highLevelPredictions *[]string, highLevelTime *[][]string) {
	//To be implemented
	return nil, nil
}
