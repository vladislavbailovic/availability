package data

import (
	"log"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type DataID int64

func (x DataID) ToItemID() int32 {
	return int32(x)
}

func (x DataID) ToNumericID() int {
	return int(x)
}

func (x DataID) IsValid() bool {
	return x.ToNumericID() > 0
}

type Scanner interface {
	Scan(...any) error
}

type Scanners []Scanner

type Collector interface {
	Query(...any) (*Scanners, error)
}

type Selector interface {
	Query(...any) (Scanner, error)
}

type MultiInserter interface {
	Insert(...any) (DataID, error)
}

type Inserter interface {
	Insert(any) (DataID, error)
}

type Updater interface {
	Update(any) error
}

func IntArgAt(args []any, pos int) int {
	var x int
	if len(args) >= pos {
		if y, ok := args[pos].(int); ok {
			x = y
		}
	}
	return x
}

func DurationArgAt(args []any, pos int) time.Duration {
	var x time.Duration
	if len(args) >= pos {
		if y, ok := args[pos].(time.Duration); ok {
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
