package bugzillaDataImporterRespones

import "time"

type GetBugsv2 struct {
	Bugs []struct {
		Flags               []interface{} `json:"flags"`
		Severity            string        `json:"severity"`
		ID                  int           `json:"id"`
		IsCreatorAccessible bool          `json:"is_creator_accessible"`
		CreatorDetail       struct {
			RealName string `json:"real_name"`
			ID       int    `json:"id"`
			Name     string `json:"name"`
			Email    string `json:"email"`
		} `json:"creator_detail"`
		LastChangeTime   time.Time   `json:"last_change_time"`
		Deadline         interface{} `json:"deadline"`
		AssignedToDetail struct {
			Name     string `json:"name"`
			Email    string `json:"email"`
			RealName string `json:"real_name"`
			ID       int    `json:"id"`
		} `json:"assigned_to_detail"`
		Component       string        `json:"component"`
		Priority        string        `json:"priority"`
		CreationTime    time.Time     `json:"creation_time"`
		Cc              []string      `json:"cc"`
		Alias           []interface{} `json:"alias"`
		Summary         string        `json:"summary"`
		QaContact       string        `json:"qa_contact"`
		Status          string        `json:"status"`
		Creator         string        `json:"creator"`
		Blocks          []interface{} `json:"blocks"`
		Groups          []interface{} `json:"groups"`
		Resolution      string        `json:"resolution"`
		OpSys           string        `json:"op_sys"`
		AssignedTo      string        `json:"assigned_to"`
		Product         string        `json:"product"`
		QaContactDetail struct {
			Email    string `json:"email"`
			Name     string `json:"name"`
			ID       int    `json:"id"`
			RealName string `json:"real_name"`
		} `json:"qa_contact_detail"`
		Whiteboard      string        `json:"whiteboard"`
		Classification  string        `json:"classification"`
		TargetMilestone string        `json:"target_milestone"`
		Platform        string        `json:"platform"`
		Version         string        `json:"version"`
		DupeOf          interface{}   `json:"dupe_of"`
		DependsOn       []interface{} `json:"depends_on"`
		Keywords        []interface{} `json:"keywords"`
		URL             string        `json:"url"`
		IsConfirmed     bool          `json:"is_confirmed"`
		IsCcAccessible  bool          `json:"is_cc_accessible"`
		SeeAlso         []interface{} `json:"see_also"`
		IsOpen          bool          `json:"is_open"`
		CcDetail        []struct {
			Email    string `json:"email"`
			Name     string `json:"name"`
			ID       int    `json:"id"`
			RealName string `json:"real_name"`
		} `json:"cc_detail"`
	} `json:"bugs"`
}
