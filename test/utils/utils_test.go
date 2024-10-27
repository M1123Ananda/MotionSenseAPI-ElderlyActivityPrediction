package utils

import (
	"MotionSense/utils"
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestModeFiltering_Case1(t *testing.T) {
	sequence := []string{"walk", "walk", "sit", "sit", "sit", "sit", "sit", "walk", "walk", "walk", "sit", "sit", "sit", "sit", "sit", "walk", "walk", "sit", "sit", "sit", "sit", "sit", "sit", "walk", "walk", "walk", "walk", "walk", "sit", "sit"}
	windowSize := 5
	expected := []string{"sit", "sit", "sit", "sit", "sit", "sit", "sit", "walk", "walk", "walk", "sit", "sit", "sit", "sit", "sit", "sit", "sit", "sit", "sit", "sit", "sit", "sit", "sit", "walk", "walk", "walk", "walk", "walk", "walk", "walk"}

	result := utils.ModeFilter(&sequence, windowSize)

	assert.Equal(t, expected, *result)
}

func TestModeFiltering_Case2(t *testing.T) {
	sequence := []string{"walk", "sit", "walk", "sit", "walk", "sit", "walk", "sit", "walk", "sit", "walk", "sit", "walk", "sit", "walk", "sit", "walk", "sit", "walk", "sit", "walk", "sit", "walk", "sit", "walk", "sit", "walk", "sit", "walk", "sit"}
	windowSize := 5
	expected := []string{"walk", "walk", "walk", "sit", "walk", "sit", "walk", "sit", "walk", "sit", "walk", "sit", "walk", "sit", "walk", "sit", "walk", "sit", "walk", "sit", "walk", "sit", "walk", "sit", "walk", "sit", "walk", "sit", "sit", "sit"}

	result := utils.ModeFilter(&sequence, windowSize)

	assert.Equal(t, expected, *result)
}

func TestModeFiltering_Case3(t *testing.T) {
	sequence := []string{"sit", "sit", "walk", "walk", "sit", "sit", "sit", "walk", "walk", "walk", "sit", "sit", "walk", "sit", "sit", "walk", "walk", "walk", "sit", "sit", "sit", "sit", "sit", "walk", "walk", "walk", "sit", "sit", "walk", "sit"}
	windowSize := 5
	expected := []string{"sit", "sit", "sit", "sit", "sit", "sit", "sit", "walk", "walk", "walk", "walk", "sit", "sit", "sit", "walk", "walk", "walk", "walk", "sit", "sit", "sit", "sit", "sit", "walk", "walk", "walk", "walk", "sit", "sit", "sit"}

	result := utils.ModeFilter(&sequence, windowSize)

	assert.Equal(t, expected, *result)
}

func TestGroupData_Case1(t *testing.T) {
	sequence := []string{
		"walk", "walk", "walk", "sit",
		"sit", "lie", "lie", "standing",
		"standing", "standing", "lie", "standing",
		"walk", "sit", "lie", "standing"}
	rooms := []string{
		"B", "B", "B", "A",
		"A", "Toilet", "B", "Toilet",
		"Toilet", "Toilet", "B", "C",
		"B", "A", "C", "Toilet"}
	timeStamps := []string{
		"2024-09-04 03:32:35.557727+00",
		"2024-09-04 04:15:12.123456+00",
		"2024-09-04 05:45:09.987654+00",
		"2024-09-04 06:30:45.123789+00",
		"2024-09-04 07:15:27.456321+00",
		"2024-09-04 08:00:00.000001+00",
		"2024-09-04 08:45:59.654321+00",
		"2024-09-04 09:30:30.111213+00",
		"2024-09-04 10:15:44.789012+00",
		"2024-09-04 11:00:55.246813+00",
		"2024-09-04 12:45:10.159753+00",
		"2024-09-04 13:30:22.365478+00",
		"2024-09-04 14:15:35.874125+00",
		"2024-09-04 15:00:48.481516+00",
		"2024-09-04 15:45:59.134679+00",
		"2024-09-04 16:30:12.278455+00"}

	expectedSequence := []string{
		"walk", "sit", "lie", "lie",
		"standing", "lie", "standing", "walk",
		"sit", "lie", "standing",
	}
	expectedRooms := []string{
		"B", "A", "Toilet", "B",
		"Toilet", "B", "C", "B",
		"A", "C", "Toilet",
	}
	expectedTimes := []string{
		"2024-09-04 03:32:35.557727+00",
		"2024-09-04 06:30:45.123789+00",
		"2024-09-04 08:00:00.000001+00",
		"2024-09-04 08:45:59.654321+00",
		"2024-09-04 09:30:30.111213+00",
		"2024-09-04 12:45:10.159753+00",
		"2024-09-04 13:30:22.365478+00",
		"2024-09-04 14:15:35.874125+00",
		"2024-09-04 15:00:48.481516+00",
		"2024-09-04 15:45:59.134679+00",
		"2024-09-04 16:30:12.278455+00"}

	resultSequence, resultRooms, resultTimes := utils.GroupData(&sequence, &rooms, &timeStamps)
	assert.Equal(t, expectedRooms, resultRooms)
	assert.Equal(t, expectedSequence, resultSequence)
	assert.Equal(t, expectedTimes, resultTimes)
}

func TestGroupData_Case2(t *testing.T) {
	sequence := []string{
		"walk", "walk", "walk", "walk"}

	rooms := []string{
		"A", "B", "B", "B"}

	timeStamps := []string{
		"2024-09-04 03:32:35.557727+00",
		"2024-09-04 04:15:12.123456+00",
		"2024-09-04 05:45:09.987654+00",
		"2024-09-04 16:30:12.278455+00",
	}

	expectedSequence := []string{
		"walk", "walk"}
	expectedRooms := []string{
		"A", "B"}
	expectedTimes := []string{
		"2024-09-04 03:32:35.557727+00",
		"2024-09-04 04:15:12.123456+00",
	}

	resultSequence, resultRooms, resultTimes := utils.GroupData(&sequence, &rooms, &timeStamps)
	assert.Equal(t, expectedRooms, resultRooms)
	assert.Equal(t, expectedSequence, resultSequence)
	assert.Equal(t, expectedTimes, resultTimes)
}

func TestGroupData_Case3(t *testing.T) {
	sequence := []string{
		"walk"}

	rooms := []string{
		"A"}

	timeStamps := []string{
		"2024-09-04 03:32:35.557727+00",
	}

	expectedSequence := []string{
		"walk"}
	expectedRooms := []string{
		"A"}
	expectedTimes := []string{
		"2024-09-04 03:32:35.557727+00",
	}

	resultSequence, resultRooms, resultTimes := utils.GroupData(&sequence, &rooms, &timeStamps)
	assert.Equal(t, expectedRooms, resultRooms)
	assert.Equal(t, expectedSequence, resultSequence)
	assert.Equal(t, expectedTimes, resultTimes)
}

func TestGroupData_Case4(t *testing.T) {
	sequence := []string{
		"walk", "walk", "sit", "lie"}

	rooms := []string{
		"A", "B", "B", "B"}

	timeStamps := []string{
		"2024-09-04 03:32:35.557727+00",
		"2024-09-04 04:15:12.123456+00",
		"2024-09-04 05:45:09.987654+00",
		"2024-09-04 16:30:12.278455+00",
	}

	expectedSequence := []string{
		"walk", "walk", "sit", "lie"}
	expectedRooms := []string{
		"A", "B", "B", "B"}
	expectedTimes := []string{
		"2024-09-04 03:32:35.557727+00",
		"2024-09-04 04:15:12.123456+00",
		"2024-09-04 05:45:09.987654+00",
		"2024-09-04 16:30:12.278455+00",
	}

	resultSequence, resultRooms, resultTimes := utils.GroupData(&sequence, &rooms, &timeStamps)
	assert.Equal(t, expectedRooms, resultRooms)
	assert.Equal(t, expectedSequence, resultSequence)
	assert.Equal(t, expectedTimes, resultTimes)
}

func TestCombineAndTransform_Case1(t *testing.T) {
	sequence := []string{
		"WALKING", "WALKING", "SITTING", "LAYING"}

	rooms := []string{
		"A", "B", "B", "B"}

	timeStamps := []string{
		"2024-09-04 03:32:35.557727+00",
		"2024-09-04 04:15:12.123456+00",
		"2024-09-04 05:45:09.987654+00",
		"2024-09-04 16:30:12.278455+00",
	}

	expectedSequence := []string{
		"WALKING_A_B", "InActivity_B"}

	expectedTimes := []string{
		"2024-09-04 03:32:35.557727+00",
		"2024-09-04 05:45:09.987654+00",
	}

	resultSequence, resultTimes := utils.CombineAndTransform(&sequence, &rooms, &timeStamps)
	assert.Equal(t, &expectedSequence, resultSequence)
	assert.Equal(t, &expectedTimes, resultTimes)
}

func TestCombineAndTransform_Case2(t *testing.T) {
	sequence := []string{
		"WALKING_DOWNSTAIRS", "WALKING_UPSTAIRS", "SITTING", "LAYING"}

	rooms := []string{
		"A", "B", "B", "B"}

	timeStamps := []string{
		"2024-09-04 03:32:35.557727+00",
		"2024-09-04 04:15:12.123456+00",
		"2024-09-04 05:45:09.987654+00",
		"2024-09-04 16:30:12.278455+00",
	}

	expectedSequence := []string{
		"WALKING_A_B", "InActivity_B"}

	expectedTimes := []string{
		"2024-09-04 03:32:35.557727+00",
		"2024-09-04 05:45:09.987654+00",
	}

	resultSequence, resultTimes := utils.CombineAndTransform(&sequence, &rooms, &timeStamps)
	assert.Equal(t, &expectedSequence, resultSequence)
	assert.Equal(t, &expectedTimes, resultTimes)
}

func TestCombineAndTransform_Case3(t *testing.T) {
	sequence := []string{
		"WALKING", "WALKING", "SITTING", "STANDING"}

	rooms := []string{
		"A", "B", "Toilet", "Toilet"}

	timeStamps := []string{
		"2024-09-04 03:32:35.557727+00",
		"2024-09-04 04:15:12.123456+00",
		"2024-09-04 05:45:09.987654+00",
		"2024-09-04 16:30:12.278455+00",
	}

	expectedSequence := []string{
		"WALKING_A_B", "UsingToilet"}

	expectedTimes := []string{
		"2024-09-04 03:32:35.557727+00",
		"2024-09-04 05:45:09.987654+00",
	}

	resultSequence, resultTimes := utils.CombineAndTransform(&sequence, &rooms, &timeStamps)
	assert.Equal(t, &expectedSequence, resultSequence)
	assert.Equal(t, &expectedTimes, resultTimes)
}
