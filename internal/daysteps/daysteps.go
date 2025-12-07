package daysteps

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
	errStepsZero    = errors.New("the number of steps is less than or equal to zero")
	errDurationZero = errors.New("the duration of activity is less than or equal to zero")
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	slice := strings.Split(datastring, ",")
	if len(slice) != 2 {
		return errLenSlice
	}
	steps, err := strconv.Atoi(slice[0])
	if steps <= 0 {
		return errStepsZero
	}
	if err != nil {
		return err
	}
	ds.Steps = steps
	duration, err := time.ParseDuration(slice[1])
	if duration <= 0 {
		return errDurationZero
	}
	if err != nil {
		return err
	}
	ds.Duration = duration
	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	distance := spentenergy.Distance(ds.Steps, ds.Personal.Height)
	var (
		calories float64
		err      error
	)
	calories, err = spentenergy.WalkingSpentCalories(ds.Steps, ds.Personal.Weight, ds.Personal.Height, ds.Duration)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, distance, calories), nil
}
