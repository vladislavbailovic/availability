package data

type Scanner interface {
	Scan(dest ...any) error
}

type Scanners []Scanner

type Collector interface {
	Query(args ...any) (*Scanners, error)
}
