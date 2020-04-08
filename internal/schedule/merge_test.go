package schedule

import (
	"testing"
	"time"
)

func makeTime(t string) time.Time {
	result, _ := time.Parse(TimeFormat, t)
	return result
}

func makeSchedule(table [][]string) Schedule {
	result := make(Schedule)

	for _, row := range table {
		date := makeTime(row[0])
		result[date] = Duty{
			Primary: Employee{row[1], nil},
			Backup:  Employee{row[2], nil},
		}
	}

	return result
}

func makeEmployees(employees []simpleEmployee) []Employee {
	result := make([]Employee, len(employees))

	for i, e := range employees {
		absence := make([]time.Time, len(e.absence))

		for j, t := range e.absence {
			absence[j] = makeTime(t)
		}

		employee := Employee{
			Name:    e.name,
			Absence: absence,
		}

		result[i] = employee
	}

	return result
}

type simpleEmployee struct {
	name    string
	absence []string
}

func TestMergeSchedules(t *testing.T) {
	for _, tc := range []struct {
		from      string
		to        string
		employees []simpleEmployee
		first     [][]string
		second    [][]string
		expected  [][]string
	}{
		{
			from: "01/01/2020",
			to:   "01/02/2020",
			employees: []simpleEmployee{
				{"a", nil},
				{"b", nil},
				{"c", nil},
			},
			first: [][]string{
				{"01/01/2020", "a", "b"},
				{"02/01/2020", "b", "c"},
				{"03/01/2020", "c", "d"},
			},
			second: [][]string{
				{"01/01/2020", "1", "2"},
				{"02/01/2020", "2", "3"},
				{"03/01/2020", "3", "4"},
				{"04/01/2020", "4", "a"},
				{"05/01/2020", "a", "5"},
				{"06/01/2020", "5", "6"},
			},
			expected: [][]string{
				{"01/01/2020", "a", "b"},
				{"02/01/2020", "b", "c"},
				{"03/01/2020", "c", ""},
				{"04/01/2020", "", "a"},
				{"05/01/2020", "a", ""},
				{"06/01/2020", "", ""},
			},
		},
	} {
		spec := MergeSpec{
			DateFrom:  makeTime(tc.from),
			DateTo:    makeTime(tc.to),
			Employees: makeEmployees(tc.employees),
			First:     makeSchedule(tc.first),
			Second:    makeSchedule(tc.second),
		}

		expected := makeSchedule(tc.expected)

		actual, _ := Merge(spec)
		if !actual.Equals(expected) {
			t.Errorf("got schedule %v but expected %v", actual, expected)
		}
	}
}
