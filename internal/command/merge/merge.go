package merge

import (
	"fmt"
	"time"

	"github.com/damirm/nurse-schedule/internal/schedule"
)

type Config struct {
	// Paths is csv files paths.
	Paths []string

	// Employees are list of employee names.
	Employees []string

	// Out is csv file output.
	Out string

	DateFrom time.Time
	DateTo   time.Time
}

func Merge(config Config) error {
	if len(config.Paths) < 2 {
		return fmt.Errorf("invalid file paths")
	}

	scheds := make([]schedule.Schedule, len(config.Paths))

	merged, err := schedule.CreateScheduleFromCSVFile(config.Paths[0])
	if err != nil {
		return err
	}

	var stat schedule.MergeStat

	for i := 1; i < len(scheds); i += 1 {
		second, err := schedule.CreateScheduleFromCSVFile(config.Paths[i])
		if err != nil {
			return err
		}

		spec := schedule.MergeSpec{
			DateFrom:  config.DateFrom,
			DateTo:    config.DateTo,
			Employees: schedule.CreateEmployeesFromNames(config.Employees),
			First:     merged,
			Second:    second,
		}

		merged, stat = schedule.Merge(spec)

		fmt.Println(stat)
	}

	if config.Out == "" {
		return nil
	}

	return schedule.SaveToCSVFile(merged, config.Out)
}
