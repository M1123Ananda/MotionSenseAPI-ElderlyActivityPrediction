package utils

import (
	"slices"
	"strings"
)

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

func transformAcitivity(acts *[]string, rooms *[]string) {
	var static = []string{"SITTING", "STANDING", "LAYING"}
	for i := 0; i < len(*acts); i++ {
		if slices.Contains(static, (*acts)[i]) {
			if (*rooms)[i] == "Toilet" {
				if (*acts)[i] != "LAYING" {
					(*acts)[i] = "UsingToilet"
				}
			} else if (*rooms)[i] == "Bedroom" {
				if (*acts)[i] == "LAYING" {
					(*acts)[i] = "Sleep"
				}
			} else { // Static Activity = InActivity
				(*acts)[i] = "InActivity"
			}
		} else { // Dynamic = Walking
			(*acts)[i] = "WALKING"
		}
	}
}

func concatRoom(act *string, room *string) {
	words := strings.Split(*act, "_")
	if len(words) > 1 {
		if words[len(words)-1] != *room {
			*act = words[0] + "_" + words[1] + "_" + *room
		}
	} else {
		if *act != "UsingToilet" {
			*act = *act + "_" + *room
		}
	}
}

func CombineAndTransform(groupedPredictions *[]string, groupedPositions *[]string, groupedTime *[]string) (highLevelPredictions *[]string, highLevelTime *[]string) {
	acts := *groupedPredictions
	rooms := *groupedPositions
	transformAcitivity(&acts, &rooms)
	var currentAct string = acts[0]
	var currentRoom string = rooms[0]
	concatRoom(&currentAct, &currentRoom)
	retPredictions := []string{currentAct}
	retTimes := []string{(*groupedTime)[0]}
	l := 0

	for i := 1; i < len(*groupedPredictions); i++ {
		currentAct = acts[i]
		currentRoom = rooms[i]
		if acts[i-1] == currentAct && currentAct != "InActivity" {
			concatRoom(&retPredictions[l], &currentRoom)
		} else {
			concatRoom(&currentAct, &currentRoom)
			if retPredictions[l] != currentAct {
				retPredictions = append(retPredictions, currentAct)
				retTimes = append(retTimes, (*groupedTime)[i])
				l++
			}
		}
	}

	return &retPredictions, &retTimes
}
