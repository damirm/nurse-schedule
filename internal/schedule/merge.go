package schedule

import (
	"time"
)

type MergeSpec struct {
	DateFrom  time.Time
	DateTo    time.Time
	Employees []Employee
	First     Schedule
	Second    Schedule
}

type employeeCache map[string]Employee

func newEmployeeCache(employees []Employee) employeeCache {
	cache := make(employeeCache, len(employees))
	for _, e := range employees {
		cache[e.Name] = e
	}
	return cache
}

func (ec employeeCache) hasEmployee(e Employee) bool {
	_, ok := ec[e.Name]
	return ok
}

type MergeStat struct {
	Overlaps map[time.Time][]Duty
}

func Merge(spec MergeSpec) (Schedule, MergeStat) {
	result := make(Schedule)
	cache := newEmployeeCache(spec.Employees)

	for date := spec.DateFrom; date.Before(spec.DateTo); date = date.AddDate(0, 0, 1) {
		fd, fok := spec.First[date]
		sd, sok := spec.Second[date]

		if !fok && !sok {
			continue
		}

		duty := Duty{}

		if fok {
			if cache.hasEmployee(fd.Primary) {
				duty.Primary = fd.Primary
			}

			if cache.hasEmployee(fd.Backup) {
				duty.Backup = fd.Backup
			}
		} else if sok {
			if cache.hasEmployee(sd.Primary) {
				duty.Primary = sd.Primary
			}

			if cache.hasEmployee(sd.Backup) {
				duty.Backup = sd.Backup
			}
		}

		result[date] = duty
	}

	return result, MergeStat{}
}
