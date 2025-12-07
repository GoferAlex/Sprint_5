package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

var (
	errStepsZero    = errors.New("the number of steps is less than or equal to zero")
	errDurationZero = errors.New("the duration of activity is less than or equal to zero")
	errWeightZero   = errors.New("the weight is less than or equal to zero")
	errHeightZero   = errors.New("the height is less than or equal to zero")
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, errStepsZero
	}
	if duration <= 0 {
		return 0, errDurationZero
	}
	if weight <= 0 {
		return 0, errWeightZero
	}
	if height <= 0 {
		return 0, errHeightZero
	}
	return (weight * MeanSpeed(steps, height, duration) * duration.Minutes() * walkingCaloriesCoefficient) / minInH, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, errStepsZero
	}
	if duration <= 0 {
		return 0, errDurationZero
	}
	return (weight * MeanSpeed(steps, height, duration) * duration.Minutes()) / minInH, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if steps <= 0 || duration <= 0 {
		return 0
	}
	return Distance(steps, height) / duration.Hours()
}

func Distance(steps int, height float64) float64 {
	return (height * stepLengthCoefficient * float64(steps)) / mInKm
}
