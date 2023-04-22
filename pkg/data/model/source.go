package model

func (x *Source) IsValid() bool {
	return x.SiteID != 0 && x.URL != ""
}
