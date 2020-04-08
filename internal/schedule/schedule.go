package schedule

import (
	"sort"
	"time"
)

const TimeFormat = "02/01/2006"

type Duty struct {
	Primary Employee
	Backup  Employee
}

func (d *Duty) Equals(other Duty) bool {
	return d.Primary.Equals(other.Primary) && d.Backup.Equals(other.Backup)
}

func (d *Duty) HasEmployee(e Employee) bool {
	return d.IsPrimary(e) || d.IsBackup(e)
}

func (d *Duty) IsPrimary(e Employee) bool {
	return d.Primary.Equals(e)
}

func (d *Duty) IsBackup(e Employee) bool {
	return d.Backup.Equals(e)
}

func (d *Duty) HasPrimary() bool {
	return d.Primary.Name != ""
}

func (d *Duty) HasBackup() bool {
	return d.Backup.Name != ""
}

type Schedule map[time.Time]Duty

func (s Schedule) Equals(other Schedule) bool {
	if len(s) != len(other) {
		return false
	}

	for t, d := range s {
		if _, has := other[t]; !has {
			return false
		}

		if !d.Equals(other[t]) {
			return false
		}
	}

	return true
}

// ToTable returns sorted shedule table representaion.
func (s Schedule) ToTable() [][]string {
	result := make([][]string, len(s))
	i := 0

	for date, duty := range s {
		result[i] = make([]string, 3)
		result[i][0] = date.Format(TimeFormat)
		result[i][1] = duty.Primary.Name
		result[i][2] = duty.Backup.Name
		i += 1
	}

	sort.SliceStable(result, func(i, j int) bool {
		a, _ := time.Parse(TimeFormat, result[i][0])
		b, _ := time.Parse(TimeFormat, result[j][0])
		return a.Before(b)
	})

	return result
}

func CreateScheduleFromTable(table [][]string) (Schedule, error) {
	sched := Schedule{}

	for _, record := range table {
		date, err := time.Parse(TimeFormat, record[0])
		if err != nil {
			return nil, err
		}

		sched[date] = Duty{
			Primary: Employee{Name: record[1]},
			Backup:  Employee{Name: record[2]},
		}
	}

	return sched, nil
}
