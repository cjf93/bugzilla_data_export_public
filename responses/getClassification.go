package bugzillaDataImporterRespones

//GetClassification is the bugzilla response from api call /rest/classification/<name>
type GetClassification struct {
	Classifications []struct {
		Description string `json:"description"`
		ID          int    `json:"id"`
		Name        string `json:"name"`
		Products    []struct {
			Description string `json:"description"`
			ID          int    `json:"id"`
			Name        string `json:"name"`
		} `json:"products"`
		SortKey int `json:"sort_key"`
	} `json:"classifications"`
}
