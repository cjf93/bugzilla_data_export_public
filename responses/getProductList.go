package bugzillaDataImporterRespones

//GetProductList is the bugzilla response from api call /rest/product_selectable
type GetProductList struct {
	Ids []string `json:"ids"`
}
