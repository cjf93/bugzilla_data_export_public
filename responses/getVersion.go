package bugzillaDataImporterRespones

//GetVersion is the bugzilla response from api call /rest/version
type GetVersion struct {
	Version string `json:"version"`
}
