package statistics

import (
	"errors"
	"os"
	"strconv"
	"strings"
	"time"
)

type Statistic struct {
	fields  []Field
	state   StatisticState
	csvFile *os.File
}

type Field struct {
	name      string
	value     string
	startTime time.Time
}

type StatisticState int8

const (
	Collecting StatisticState = iota
	Ready
	Closed
)

func NewStatistic(filename string, fieldNames ...string) Statistic {
	fields := make([]Field, 0)
	for _, fieldName := range fieldNames {
		field := Field{fieldName, "", time.Now()}
		fields = append(fields, field)
	}
	file, _ := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
	fileInfo, _ := file.Stat()
	if fileInfo.Size() == 0 {
		sb := strings.Builder{}
		sb.WriteString("ID")
		for _, fieldName := range fieldNames {
			sb.WriteString("," + fieldName)
		}
		file.WriteString(sb.String())
	}
	return Statistic{
		fields:  fields,
		state:   Ready,
		csvFile: file,
	}
}

func (statistic *Statistic) BeginnColum() error {
	switch statistic.state {
	case Collecting:
		return errors.New("colum has to be ended before beginning a new one")
	case Closed:
		return errors.New("this operation is not allowed, because this statistic is closed")
	}
	statistic.state = Collecting
	now := time.Now()
	for _, field := range statistic.fields {
		field.value = ""
		field.startTime = now
	}
	return nil
}

func (statistic *Statistic) SetValue(fieldName string, value string) error {
	switch statistic.state {
	case Ready:
		return errors.New("start new colum before setting a value")
	case Closed:
		return errors.New("this operation is not allowed, because this statistic is closed")
	}

	for _, field := range statistic.fields {
		if field.name == fieldName {
			field.value = value
		}
	}
	return nil
}

func (statistic *Statistic) StartTimer(fieldName string) error {
	switch statistic.state {
	case Ready:
		return errors.New("start new colum before starting a timer")
	case Closed:
		return errors.New("this operation is not allowed, because this statistic is closed")
	}

	for _, field := range statistic.fields {
		if field.name == fieldName {
			field.startTime = time.Now()
		}
	}

	return nil
}

func (statistic *Statistic) StopTimerAndSetDuration(fieldName string) error {
	switch statistic.state {
	case Ready:
		return errors.New("start new colum before stopping a timer")
	case Closed:
		return errors.New("this operation is not allowed, because this statistic is closed")
	}

	for _, field := range statistic.fields {
		if field.name == fieldName {
			duration := time.Since(field.startTime)
			field.value = strconv.FormatInt(duration.Milliseconds(), 10)
		}
	}

	return nil
}

func (statistic *Statistic) EndColum() error {
	switch statistic.state {
	case Ready:
		return errors.New("start new colum before ending a colum")
	case Closed:
		return errors.New("this operation is not allowed, because this statistic is closed")
	}
	statistic.state = Ready

	sb := strings.Builder{}
	for i, field := range statistic.fields {
		if i == 0 {
			sb.WriteString(field.value)
		} else {
			sb.WriteString("," + field.value)
		}
	}

	_, err := statistic.csvFile.WriteString(sb.String())
	return err

}

func (statistic *Statistic) Close() error {
	if statistic.state == Closed {
		return errors.New("this statistic is already closed")
	}
	statistic.state = Closed
	return statistic.csvFile.Close()

}
