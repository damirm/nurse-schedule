package schedule

import (
	"sort"
	"time"
)

type Employee struct {
	Name string
	// Absence must be sorted.
	Absence []time.Time
}

func (e *Employee) IsAbsent(t time.Time) bool {
	return sort.Search(len(e.Absence), func(i int) bool {
		return e.Absence[i] == t
	}) != -1
}

func (e *Employee) Equals(other Employee) bool {
	return e.Name == other.Name
}

func CreateEmployeesFromNames(names []string) []Employee {
	result := make([]Employee, len(names))

	for i, name := range names {
		result[i] = Employee{Name: name}
	}

	return result
}
