package schedule

import (
	"encoding/csv"
	"io"
	"os"
	"strings"
	"time"
)

const DefaultHeader = "Date,Primary,Backup"

// cvs format:
// Date,Primary,Backup
// 01/08/2018,a,b
// 02/08/2018,b,c

func isHeader(row []string) bool {
	return strings.Join(row, ",") == DefaultHeader
}

func CreateScheduleFromCSVFile(path string) (Schedule, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return CreateScheduleFromReader(f)
}

func CreateScheduleFromReader(reader io.Reader) (Schedule, error) {
	sched := Schedule{}

	r := csv.NewReader(reader)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		if isHeader(record) {
			continue
		}

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

func SaveToCSVFile(sched Schedule, path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}

	defer f.Close()

	return SaveToWriter(sched, f)
}

func SaveToWriter(sched Schedule, writer io.Writer) error {
	w := csv.NewWriter(writer)

	err := w.Write(strings.Split(DefaultHeader, ","))
	if err != nil {
		return err
	}

	for _, row := range sched.ToTable() {
		err := w.Write(row)
		if err != nil {
			return err
		}
	}

	w.Flush()

	return nil
}
