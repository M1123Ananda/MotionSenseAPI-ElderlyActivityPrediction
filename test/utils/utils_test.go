package utils

import (
	"MotionSense/utils"
	"testing"

	"github.com/go-playground/assert/v2"
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
	time_stamps := []string{
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

	expected_sequence := []string{
		"walk", "sit", "lie", "lie",
		"standing", "lie", "standing", "walk",
		"sit", "lie", "standing",
	}
	expected_rooms := []string{
		"B", "A", "Toilet", "B",
		"Toilet", "B", "C", "B",
		"A", "C", "Toilet",
	}
	expected_times := []string{
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

	result_sequence, result_rooms, result_times := utils.GroupData(&sequence, &rooms, &time_stamps)
	assert.Equal(t, expected_rooms, result_rooms)
	assert.Equal(t, expected_sequence, result_sequence)
	assert.Equal(t, expected_times, result_times)
}

func TestGroupData_Case2(t *testing.T) {
	sequence := []string{
		"walk", "walk", "walk", "walk"}

	rooms := []string{
		"A", "B", "B", "B"}

	time_stamps := []string{
		"2024-09-04 03:32:35.557727+00",
		"2024-09-04 04:15:12.123456+00",
		"2024-09-04 05:45:09.987654+00",
		"2024-09-04 16:30:12.278455+00",
	}

	expected_sequence := []string{
		"walk", "walk"}
	expected_rooms := []string{
		"A", "B"}
	expected_times := []string{
		"2024-09-04 03:32:35.557727+00",
		"2024-09-04 04:15:12.123456+00",
	}

	result_sequence, result_rooms, result_times := utils.GroupData(&sequence, &rooms, &time_stamps)
	assert.Equal(t, expected_rooms, result_rooms)
	assert.Equal(t, expected_sequence, result_sequence)
	assert.Equal(t, expected_times, result_times)
}

func TestGroupData_Case3(t *testing.T) {
	sequence := []string{
		"walk"}

	rooms := []string{
		"A"}

	time_stamps := []string{
		"2024-09-04 03:32:35.557727+00",
	}

	expected_sequence := []string{
		"walk"}
	expected_rooms := []string{
		"A"}
	expected_times := []string{
		"2024-09-04 03:32:35.557727+00",
	}

	result_sequence, result_rooms, result_times := utils.GroupData(&sequence, &rooms, &time_stamps)
	assert.Equal(t, expected_rooms, result_rooms)
	assert.Equal(t, expected_sequence, result_sequence)
	assert.Equal(t, expected_times, result_times)
}

func TestGroupData_Case4(t *testing.T) {
	sequence := []string{
		"walk", "walk", "sit", "lie"}

	rooms := []string{
		"A", "B", "B", "B"}

	time_stamps := []string{
		"2024-09-04 03:32:35.557727+00",
		"2024-09-04 04:15:12.123456+00",
		"2024-09-04 05:45:09.987654+00",
		"2024-09-04 16:30:12.278455+00",
	}

	expected_sequence := []string{
		"walk", "walk", "sit", "lie"}
	expected_rooms := []string{
		"A", "B", "B", "B"}
	expected_times := []string{
		"2024-09-04 03:32:35.557727+00",
		"2024-09-04 04:15:12.123456+00",
		"2024-09-04 05:45:09.987654+00",
		"2024-09-04 16:30:12.278455+00",
	}

	result_sequence, result_rooms, result_times := utils.GroupData(&sequence, &rooms, &time_stamps)
	assert.Equal(t, expected_rooms, result_rooms)
	assert.Equal(t, expected_sequence, result_sequence)
	assert.Equal(t, expected_times, result_times)
}

func TestCombineAndTransform_Case1(t *testing.T) {
	//Mos
	sequence := []string{
		"WALKING", "WALKING", "SITTING", "LAYING"}

	rooms := []string{
		"A", "B", "B", "B"}

	time_stamps := []string{
		"2024-09-04 03:32:35.557727+00",
		"2024-09-04 04:15:12.123456+00",
		"2024-09-04 05:45:09.987654+00",
		"2024-09-04 16:30:12.278455+00",
	}

	expected_sequence := []string{
		"WALKING_A_B", "InActivity_B"}

	expected_times := []string{
		"2024-09-04 03:32:35.557727+00",
		"2024-09-04 05:45:09.987654+00",
	}

	result_sequence, result_times := utils.CombineAndTransform(&sequence, &rooms, &time_stamps)
	assert.Equal(t, &expected_sequence, result_sequence)
	assert.Equal(t, &expected_times, result_times)
}

func TestCombineAndTransform_Case2(t *testing.T) {
	//Mos
	sequence := []string{
		"WALKING_DOWNSTAIRS", "WALKING_UPSTAIRS", "SITTING", "LAYING"}

	rooms := []string{
		"A", "B", "B", "B"}

	time_stamps := []string{
		"2024-09-04 03:32:35.557727+00",
		"2024-09-04 04:15:12.123456+00",
		"2024-09-04 05:45:09.987654+00",
		"2024-09-04 16:30:12.278455+00",
	}

	expected_sequence := []string{
		"WALKING_A_B", "InActivity_B"}

	expected_times := []string{
		"2024-09-04 03:32:35.557727+00",
		"2024-09-04 05:45:09.987654+00",
	}

	result_sequence, result_times := utils.CombineAndTransform(&sequence, &rooms, &time_stamps)
	assert.Equal(t, &expected_sequence, result_sequence)
	assert.Equal(t, &expected_times, result_times)
}

func TestCombineAndTransform_Case3(t *testing.T) {
	//Mos
	sequence := []string{
		"WALKING", "WALKING", "SITTING", "STANDING"}

	rooms := []string{
		"A", "B", "Toilet", "Toilet"}

	time_stamps := []string{
		"2024-09-04 03:32:35.557727+00",
		"2024-09-04 04:15:12.123456+00",
		"2024-09-04 05:45:09.987654+00",
		"2024-09-04 16:30:12.278455+00",
	}

	expected_sequence := []string{
		"WALKING_A_B", "UsingToilet"}

	expected_times := []string{
		"2024-09-04 03:32:35.557727+00",
		"2024-09-04 05:45:09.987654+00",
	}

	result_sequence, result_times := utils.CombineAndTransform(&sequence, &rooms, &time_stamps)
	assert.Equal(t, &expected_sequence, result_sequence)
	assert.Equal(t, &expected_times, result_times)
}
