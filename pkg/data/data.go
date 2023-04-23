package data

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
