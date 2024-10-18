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

//func TestGroupData_Case1(t *testing.T) {
//	//Mos
//	assert.Equal(t, "1", "2")
//}

//func TestCombineAndTransform_Case1(t *testing.T) {
//	//Mos
//	assert.Equal(t, "1", "2")
//}
