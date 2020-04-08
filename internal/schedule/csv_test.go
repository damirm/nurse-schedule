package schedule

import (
	"strings"
	"testing"
)

func tablesAreEqual(a, b [][]string) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i += 1 {
		for j := 0; j < 2; j += 1 {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}

	return true
}

func TestCreateScheduleFromReader(t *testing.T) {
	for _, tc := range []struct {
		in       string
		expected [][]string
		err      bool
	}{
		{
			in: `Date,Primary,Backup
01/01/2020,a,b
02/01/2020,b,c
03/01/2020,c,d`,
			expected: [][]string{
				{"01/01/2020", "a", "b"},
				{"02/01/2020", "b", "c"},
				{"03/01/2020", "c", "d"},
			},
			err: false,
		},
	} {
		sched, err := CreateScheduleFromReader(strings.NewReader(tc.in))
		if err != nil && tc.err != true {
			t.Errorf("got unexpected error %#v", err)
		}

		table := sched.ToTable()
		if !tablesAreEqual(table, tc.expected) {
			t.Errorf("got table %v but expected %v", table, tc.expected)
		}
	}
}
