package models

import (
	"strconv"
	"strings"
)

type Bug struct {
	bug_id              int
	assigned_to         int
	bug_file_loc        string
	bug_severity        string
	bug_status          string
	creation_ts         string
	delta_ts            string
	short_desc          string
	op_sys              string
	priority            string
	product_id          int
	rep_platform        string
	reporter            int
	version             string
	component_id        int
	resolution          string
	qa_contact          int
	target_milestone    string
	status_whiteboard   string
	lastdiffed          string
	everconfirmed       bool
	reporter_accessible bool
	cclist_accessible   bool
	estimated_time      float32
	remaining_time      float32
	//deadline            string
}

type BugList []Bug

type BugMap map[int]Bug

func CreateBug(id, assigned_to, product_id, reporter, component_id, qa_contact int,
	bug_severity, bug_status, creation_ts /*, delta_ts*/, short_desc, op_sys, priority, rep_platform, version, resolution, target_milestone, status_whiteboard,
	lastdiffed /*, deadline*/ string /**estimated_time, remaining_time float32,*/, everconfirmed, reporter_accessible, cclist_accessible bool) Bug {
	var newBug Bug
	newBug.bug_id = id
	newBug.assigned_to = assigned_to
	newBug.bug_file_loc = ""
	newBug.bug_severity = bug_severity
	newBug.bug_status = bug_status
	newBug.creation_ts = creation_ts
	newBug.delta_ts = creation_ts
	str := strings.Replace(short_desc, "\\", "\\\\", -1)
	str = strings.Replace(str, "'", "\\'", -1)
	str = strings.Replace(str, "\"", "\\\"", -1)
	newBug.short_desc = str
	newBug.op_sys = op_sys
	newBug.priority = priority
	newBug.product_id = product_id
	newBug.rep_platform = rep_platform
	newBug.reporter = reporter
	newBug.version = version
	newBug.component_id = component_id
	newBug.resolution = resolution
	newBug.qa_contact = qa_contact
	newBug.target_milestone = target_milestone
	newBug.status_whiteboard = status_whiteboard
	newBug.lastdiffed = lastdiffed
	newBug.everconfirmed = everconfirmed
	newBug.reporter_accessible = reporter_accessible
	newBug.cclist_accessible = cclist_accessible
	//newBug.estimated_time = estimated_time
	//newBug.remaining_time = remaining_time
	//newBug.deadline = deadline

	return newBug
}

func (b *Bug) GenerateInsert() string {
	if b.qa_contact == -1 {
		query := "INSERT INTO `bugs`.`bugs` (`bug_id`, `assigned_to`, `bug_file_loc`, `bug_severity`, `bug_status`, `creation_ts`, `delta_ts`, `short_desc`,`op_sys`, `priority`, `product_id`,`rep_platform`, `reporter`, `version`, `component_id`, `resolution`, `target_milestone`, `qa_contact`,`status_whiteboard`, `lastdiffed`, `everconfirmed`,`reporter_accessible`, `cclist_accessible`, `estimated_time`, `remaining_time`) VALUES (" +
			strconv.Itoa(b.bug_id) + "," + strconv.Itoa(b.assigned_to) + "," + "\"" + b.bug_file_loc + "\"" + "," + "\"" + b.bug_severity + "\"" + "," + "\"" + b.bug_status + "\"" +
			"," + "\"" + b.creation_ts + "\"" + "," + "\"" + b.delta_ts + "\"" + "," + "\"" + b.short_desc + "\"" + "," + "\"" + b.op_sys + "\"" + "," + "\"" + b.priority + "\"" + "," +
			strconv.Itoa(b.product_id) + "," + "\"" + b.rep_platform + "\"" + "," + strconv.Itoa(b.reporter) + "," + "\"" + b.version + "\"" + "," + strconv.Itoa(b.component_id) + "," +
			"\"" + b.resolution + "\"" + "," + "\"" + b.target_milestone + "\"" + ", NULL ," + "\"" + b.status_whiteboard + "\"" + "," +
			"\"" + b.lastdiffed + "\"" + "," + "\"" + strconv.FormatBool(b.everconfirmed) + "\"" + "," + "\"" + strconv.FormatBool(b.reporter_accessible) + "\"" + "," +
			"\"" + strconv.FormatBool(b.cclist_accessible) + "\"" + ", 0.00 , 0.00);"

		return query
	}

	query := "INSERT INTO `bugs`.`bugs` (`bug_id`, `assigned_to`, `bug_file_loc`, `bug_severity`, `bug_status`, `creation_ts`, `delta_ts`, `short_desc`,`op_sys`, `priority`, `product_id`,`rep_platform`, `reporter`, `version`, `component_id`, `resolution`, `target_milestone`, `qa_contact`,`status_whiteboard`, `lastdiffed`, `everconfirmed`,`reporter_accessible`, `cclist_accessible`, `estimated_time`, `remaining_time`) VALUES (" +
		strconv.Itoa(b.bug_id) + "," + strconv.Itoa(b.assigned_to) + "," + "\"" + b.bug_file_loc + "\"" + "," + "\"" + b.bug_severity + "\"" + "," + "\"" + b.bug_status + "\"" +
		"," + "\"" + b.creation_ts + "\"" + "," + "\"" + b.delta_ts + "\"" + "," + "\"" + b.short_desc + "\"" + "," + "\"" + b.op_sys + "\"" + "," + "\"" + b.priority + "\"" + "," +
		strconv.Itoa(b.product_id) + "," + "\"" + b.rep_platform + "\"" + "," + strconv.Itoa(b.reporter) + "," + "\"" + b.version + "\"" + "," + strconv.Itoa(b.component_id) + "," +
		"\"" + b.resolution + "\"" + "," + "\"" + b.target_milestone + "\"" + "," + strconv.Itoa(b.qa_contact) + "," + "\"" + b.status_whiteboard + "\"" + "," +
		"\"" + b.lastdiffed + "\"" + "," + "\"" + strconv.FormatBool(b.everconfirmed) + "\"" + "," + "\"" + strconv.FormatBool(b.reporter_accessible) + "\"" + "," +
		"\"" + strconv.FormatBool(b.cclist_accessible) + "\"" + ", 0.00 , 0.00);"

	return query

}
