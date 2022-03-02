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
	Paused
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
		sb.WriteString("\n")
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
	case Paused:
		return errors.New("this operation is not allowed, because this statistic is paused")
	}
	statistic.state = Collecting
	now := time.Now()
	for i, _ := range statistic.fields {
		statistic.fields[i].value = ""
		statistic.fields[i].startTime = now
	}
	return nil
}

func (statistic *Statistic) SetValue(fieldName string, value string) error {
	switch statistic.state {
	case Ready:
		return errors.New("start new colum before setting a value")
	case Closed:
		return errors.New("this operation is not allowed, because this statistic is closed")
	case Paused:
		return errors.New("this operation is not allowed, because this statistic is paused")
	}

	for i, field := range statistic.fields {
		if field.name == fieldName {
			statistic.fields[i].value = value
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
	case Paused:
		return errors.New("this operation is not allowed, because this statistic is paused")
	}

	for i, field := range statistic.fields {
		if field.name == fieldName {
			statistic.fields[i].startTime = time.Now()
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
	case Paused:
		return errors.New("this operation is not allowed, because this statistic is paused")
	}
	fields := statistic.fields
	for i, field := range statistic.fields {
		if field.name == fieldName {
			duration := time.Since(field.startTime)
			fields[i].value = strconv.FormatInt(duration.Milliseconds(), 10)
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
	case Paused:
		return errors.New("this operation is not allowed, because this statistic is paused")
	}
	statistic.state = Ready

	sb := strings.Builder{}
	sb.WriteString(time.Now().Format(time.RFC3339))
	for _, field := range statistic.fields {
		sb.WriteString("," + field.value)
	}
	sb.WriteString("\n")
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

func (statistic *Statistic) Pause() error {
	switch statistic.state {
	case Closed:
		return errors.New("this operation is not allowed, because this statistic is closed")
	case Paused:
		return errors.New("this operation is not allowed, because this statistic is paused")
	case Collecting:
		return errors.New("statistic cant be paused, while collecting data. end colum to pause")
	}
	statistic.state = Paused
	return nil
}
