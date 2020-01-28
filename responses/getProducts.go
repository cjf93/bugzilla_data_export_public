package bugzillaDataImporterRespones

//GetProducts is the bugzilla response from api call /rest/product
type GetProducts struct {
	Products []struct {
		Classification string `json:"classification"`
		Components     []struct {
			DefaultAssignedTo string `json:"default_assigned_to"`
			DefaultQaContact  string `json:"default_qa_contact"`
			Description       string `json:"description"`
			FlagTypes         struct {
				Attachment []struct {
					CcList          string `json:"cc_list"`
					Description     string `json:"description"`
					GrantGroup      int    `json:"grant_group"`
					ID              int    `json:"id"`
					IsActive        bool   `json:"is_active"`
					IsMultiplicable bool   `json:"is_multiplicable"`
					IsRequestable   bool   `json:"is_requestable"`
					IsRequesteeble  bool   `json:"is_requesteeble"`
					Name            string `json:"name"`
					RequestGroup    int    `json:"request_group"`
					SortKey         int    `json:"sort_key"`
				} `json:"attachment"`
				Bug []struct {
					CcList          string      `json:"cc_list"`
					Description     string      `json:"description"`
					GrantGroup      int         `json:"grant_group"`
					ID              int         `json:"id"`
					IsActive        bool        `json:"is_active"`
					IsMultiplicable bool        `json:"is_multiplicable"`
					IsRequestable   bool        `json:"is_requestable"`
					IsRequesteeble  bool        `json:"is_requesteeble"`
					Name            string      `json:"name"`
					RequestGroup    interface{} `json:"request_group"`
					SortKey         int         `json:"sort_key"`
				} `json:"bug"`
			} `json:"flag_types"`
			ID       int    `json:"id"`
			IsActive bool   `json:"is_active"`
			Name     string `json:"name"`
			SortKey  int    `json:"sort_key"`
		} `json:"components"`
		DefaultMilestone string `json:"default_milestone"`
		Description      string `json:"description"`
		HasUnconfirmed   bool   `json:"has_unconfirmed"`
		ID               int    `json:"id"`
		IsActive         bool   `json:"is_active"`
		Milestones       []struct {
			ID       int    `json:"id"`
			IsActive bool   `json:"is_active"`
			Name     string `json:"name"`
			SortKey  int    `json:"sort_key"`
		} `json:"milestones"`
		Name     string `json:"name"`
		Versions []struct {
			ID       int    `json:"id"`
			IsActive bool   `json:"is_active"`
			Name     string `json:"name"`
			SortKey  int    `json:"sort_key"`
		} `json:"versions"`
	} `json:"products"`
}
