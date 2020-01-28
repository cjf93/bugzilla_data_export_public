package bugzillaDataImporterRespones

// ErrorResponse is the Bugzilla API response from every error
type ErrorResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}
