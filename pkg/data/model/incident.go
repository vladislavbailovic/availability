package model

func NewIncident(siteID, downProbeId int) *Incident {
	incident := new(Incident)
	incident.SiteID = int32(siteID)
	incident.DownProbeID = int32(downProbeId)
	return incident
}

func (x *Incident) Close(upProbeID int) {
	x.UpProbeID = int32(upProbeID)
}

func (x *IncidentReport) IsValid() bool {
	return x.SiteID > 0
}
