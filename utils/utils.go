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

func GroupData(cleanPredictions *[]string, cleanPositions *[]string, time *[]string) (groupedPredictions *[]string, groupedPositions *[]string, groupedTime *[]string) {
	var previousAct string = (*cleanPredictions)[0]
	var previousPos string = (*cleanPositions)[0]
	var l int = 0

	groupedPredictions = &[]string{previousAct}
	groupedPositions = &[]string{previousPos}
	groupedTime = &[]string{(*time)[0]}

	for i := 1; i < len(*cleanPredictions); i++ {

		if ((*cleanPredictions)[i] != previousAct) || ((*cleanPositions)[i] != previousPos) {
			*groupedPredictions = append(*groupedPredictions, (*cleanPredictions)[i])
			*groupedPositions = append(*groupedPositions, (*cleanPositions)[i])
			*groupedTime = append(*groupedTime, (*time)[i])
			l++
		}
		previousAct = (*groupedPredictions)[l]
		previousPos = (*groupedPositions)[l]
	}

	return groupedPredictions, groupedPositions, groupedTime
}

// func transformAcitivity(act *string, room *string) {
// 	var static = []string{"SITTING", "STANDIG", "LAYING"}

// 	if slices.Contains(static, *act) {
// 		if *room == "Toilet" {
// 			*act = "UsingToilet"
// 		} else {
// 			*act = "InActivity"
// 		}
// 	} else {
// 		*act = "WALK_" + *room
// 	}
// }

// func CombineAndTransform(groupedPredictions *[]string, groupedPositions *[]string, groupedTime *[]string) (highLevelPredictions *[]string, highLevelTime *[]string) {
// 	var static = []string{"SITTING", "STANDIG", "LAYING"}
// 	var dynamic = []string{"WALKING", "WALKING_UPSTAIRS", "WALKING_DOWNSTAIRS"}
// 	var previousAct string = (*groupedPredictions)[0]
// 	retPredictions := []string{previousAct}
// 	retTimes := []string{(*groupedTime)[0]}
// 	transformAcitivity(&previousAct)

// 	for i := 1; i < len(*groupedPredictions); i++ {
// 		if ()
// 	}

// 	return nil, nil
// }
