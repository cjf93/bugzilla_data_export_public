package bugzillaDataImporterRespones

import "time"

//GetBugs is the bugzilla response from api call /rest/bug
type GetBugs struct {
	Bugs []struct {
		Alias            []interface{} `json:"alias"`
		AssignedTo       string        `json:"assigned_to"`
		AssignedToDetail struct {
			Email    string `json:"email"`
			ID       int    `json:"id"`
			Name     string `json:"name"`
			RealName string `json:"real_name"`
		} `json:"assigned_to_detail"`
		Blocks   []interface{} `json:"blocks"`
		Cc       []string      `json:"cc"`
		CcDetail []struct {
			Email    string `json:"email"`
			ID       int    `json:"id"`
			Name     string `json:"name"`
			RealName string `json:"real_name"`
		} `json:"cc_detail"`
		Classification string    `json:"classification"`
		Component      string    `json:"component"`
		CreationTime   time.Time `json:"creation_time"`
		Creator        string    `json:"creator"`
		CreatorDetail  struct {
			Email    string `json:"email"`
			ID       int    `json:"id"`
			Name     string `json:"name"`
			RealName string `json:"real_name"`
		} `json:"creator_detail"`
		Deadline            interface{}   `json:"deadline"`
		DependsOn           []interface{} `json:"depends_on"`
		DupeOf              interface{}   `json:"dupe_of"`
		Flags               []interface{} `json:"flags"`
		Groups              []interface{} `json:"groups"`
		ID                  int           `json:"id"`
		IsCcAccessible      bool          `json:"is_cc_accessible"`
		IsConfirmed         bool          `json:"is_confirmed"`
		IsCreatorAccessible bool          `json:"is_creator_accessible"`
		IsOpen              bool          `json:"is_open"`
		Keywords            []interface{} `json:"keywords"`
		LastChangeTime      time.Time     `json:"last_change_time"`
		OpSys               string        `json:"op_sys"`
		Platform            string        `json:"platform"`
		Priority            string        `json:"priority"`
		Product             string        `json:"product"`
		QaContact           string        `json:"qa_contact"`
		Resolution          string        `json:"resolution"`
		SeeAlso             []string      `json:"see_also"`
		Severity            string        `json:"severity"`
		Status              string        `json:"status"`
		Summary             string        `json:"summary"`
		TargetMilestone     string        `json:"target_milestone"`
		URL                 string        `json:"url"`
		Version             string        `json:"version"`
		Whiteboard          string        `json:"whiteboard"`
		//ActualTime          float32       `json:"actual_time"`
		//EstimatedTime       float32       `json:"estimated_time"`
		//RemainingTime       float32       `json:"remaining_time"`  This only appear when logged in a time tracking group
	} `json:"bugs"`
}
