package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

var (
	errLenSlice     = errors.New("the slice length is not equal to 3")
	errTraining     = errors.New("unknown type of training")
	errStepsZero    = errors.New("the number of steps is less than or equal to zero")
	errDurationZero = errors.New("the duration of activity is less than or equal to zero")
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	slice := strings.Split(datastring, ",")
	if len(slice) != 3 {
		return errLenSlice
	}
	steps, err := strconv.Atoi(slice[0])
	if steps <= 0 {
		return errStepsZero
	}
	if err != nil {
		return err
	}
	t.Steps = steps
	t.TrainingType = slice[1]
	duration, err := time.ParseDuration(slice[2])
	if duration <= 0 {
		return errDurationZero
	}
	if err != nil {
		return err
	}
	t.Duration = duration
	return nil
}

func (t Training) ActionInfo() (string, error) {
	distance := spentenergy.Distance(t.Steps, t.Personal.Height)
	meanSpeed := spentenergy.MeanSpeed(t.Steps, t.Personal.Height, t.Duration)
	var (
		calories float64
		err      error
	)
	switch t.TrainingType {
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
		if err != nil {
			return "", err
		}
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
		if err != nil {
			return "", err
		}
	default:
		return "", errTraining
	}
	return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", t.TrainingType, t.Duration.Hours(), distance, meanSpeed, calories), nil
}
