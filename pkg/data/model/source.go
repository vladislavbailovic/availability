package model

type SourceStatus byte

const (
	SourceInactive SourceStatus = iota
	SourceActive
)

func (x *Source) IsValid() bool {
	return x.SiteID != 0 && x.URL != ""
}

func (x *NewSource) IsValid() bool {
	return x.SiteID != 0 && x.URL != ""
}
