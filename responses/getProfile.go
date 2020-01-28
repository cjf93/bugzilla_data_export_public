package bugzillaDataImporterRespones

type GetProfile struct {
	Users []struct {
		ID       int    `json:"id"`
		Name     string `json:"name"`
		RealName string `json:"real_name"`
	} `json:"users"`
}
