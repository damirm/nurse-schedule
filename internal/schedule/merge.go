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
	Conflicts map[time.Time][]Duty
}

func (s *MergeStat) AddConflict(t time.Time, duties []Duty) {
	if s.Conflicts == nil {
		s.Conflicts = make(map[time.Time][]Duty)
	}

	arr, has := s.Conflicts[t]
	if !has {
		arr = []Duty{}
	}

	arr = append(arr, duties...)
	s.Conflicts[t] = arr
}

type Merger struct {
	Stat MergeStat
	Spec MergeSpec

	cache employeeCache
}

func newMerger(spec MergeSpec) Merger {
	return Merger{
		Stat:  MergeStat{},
		Spec:  spec,
		cache: newEmployeeCache(spec.Employees),
	}
}

func Merge(spec MergeSpec) (Schedule, MergeStat) {
	m := newMerger(spec)
	return m.Merge(), m.Stat
}

func (m *Merger) Merge() Schedule {
	result := make(Schedule)

	for date := m.Spec.DateFrom; date.Before(m.Spec.DateTo); date = date.AddDate(0, 0, 1) {
		fd, fok := m.Spec.First[date]
		sd, sok := m.Spec.Second[date]

		if !fok && !sok {
			continue
		}

		duty := Duty{}

		if fok {
			if m.cache.hasEmployee(fd.Primary) {
				duty.Primary = fd.Primary
			}

			if m.cache.hasEmployee(fd.Backup) {
				duty.Backup = fd.Backup
			}
		}

		if sok {
			if m.cache.hasEmployee(sd.Primary) {
				if duty.HasPrimary() {
					m.Stat.AddConflict(date, []Duty{duty, sd})
				}

				duty.Primary = sd.Primary
			}

			if m.cache.hasEmployee(sd.Backup) {
				if duty.HasBackup() {
					m.Stat.AddConflict(date, []Duty{duty, sd})
				}

				duty.Backup = sd.Backup
			}
		}

		result[date] = duty
	}

	return result
}
