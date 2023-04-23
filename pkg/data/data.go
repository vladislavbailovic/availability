package data

import (
	"log"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type Scanner interface {
	Scan(dest ...any) error
}

type Scanners []Scanner

type Collector interface {
	Query(args ...any) (*Scanners, error)
}

type Inserter interface {
	Insert(items ...any) error
}

func IntArgAt(args []any, pos int) int {
	var x int
	if len(args) > 0 {
		if y, ok := args[0].(int); ok {
			x = y
		}
	}
	return x
}

const (
	SourceDatetimeFormat string = "2006-01-02 15:04:05"
)

func TimestampFromDatetime(raw string) *timestamppb.Timestamp {
	ts, err := time.Parse(SourceDatetimeFormat, raw)
	if err != nil {
		log.Printf("WARNING unable to parse timestamp %s because %v", raw, err)
		return nil
	}
	return timestamppb.New(ts)
}

func DatetimeToTimestamp(ts *timestamppb.Timestamp) string {
	dt := time.Now()
	if ts != nil {
		dt = ts.AsTime()
	}
	return dt.Format(SourceDatetimeFormat)
}
