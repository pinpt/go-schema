// GENERATED CODE. DO NOT EDIT

package schema

import (
	"bytes"
	"compress/gzip"
	"context"
	"database/sql"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/jhaynie/go-gator/orm"
)

// compiler checks for interface implementations. if the generated model
// doesn't implement these interfaces for some reason you'll get a compiler error

var _ Model = (*JiraIssue)(nil)
var _ CSVWriter = (*JiraIssue)(nil)
var _ JSONWriter = (*JiraIssue)(nil)
var _ Checksum = (*JiraIssue)(nil)

// JiraIssueTableName is the name of the table in SQL
const JiraIssueTableName = "jira_issue"

var JiraIssueColumns = []string{
	"id",
	"checksum",
	"issue_id",
	"project_id",
	"issue_type_id",
	"user_id",
	"assignee_id",
	"priority_id",
	"status_id",
	"resolution_id",
	"fix_version_ids",
	"version_ids",
	"environment",
	"component_ids",
	"label_ids",
	"duedate_at",
	"planned_start_at",
	"planned_end_at",
	"key",
	"custom_field_ids",
	"sprint_id",
	"epic_id",
	"parent_id",
	"strategic_parent_id",
	"in_progress_count",
	"reopen_count",
	"in_progress_duration",
	"verification_duration",
	"in_progress_start_at",
	"customer_id",
	"ref_id",
	"user_ref_id",
	"cost",
}

// JiraIssue table
type JiraIssue struct {
	AssigneeID           *string `json:"assignee_id,omitempty"`
	Checksum             *string `json:"checksum,omitempty"`
	ComponentIds         *string `json:"component_ids,omitempty"`
	Cost                 float64 `json:"cost"`
	CustomFieldIds       *string `json:"custom_field_ids,omitempty"`
	CustomerID           string  `json:"customer_id"`
	DuedateAt            *int64  `json:"duedate_at,omitempty"`
	Environment          *string `json:"environment,omitempty"`
	EpicID               *string `json:"epic_id,omitempty"`
	FixVersionIds        *string `json:"fix_version_ids,omitempty"`
	ID                   string  `json:"id"`
	InProgressCount      int32   `json:"in_progress_count"`
	InProgressDuration   int64   `json:"in_progress_duration"`
	InProgressStartAt    *int64  `json:"in_progress_start_at,omitempty"`
	IssueID              string  `json:"issue_id"`
	IssueTypeID          string  `json:"issue_type_id"`
	Key                  string  `json:"key"`
	LabelIds             *string `json:"label_ids,omitempty"`
	ParentID             *string `json:"parent_id,omitempty"`
	PlannedEndAt         *int64  `json:"planned_end_at,omitempty"`
	PlannedStartAt       *int64  `json:"planned_start_at,omitempty"`
	PriorityID           *string `json:"priority_id,omitempty"`
	ProjectID            string  `json:"project_id"`
	RefID                string  `json:"ref_id"`
	ReopenCount          int32   `json:"reopen_count"`
	ResolutionID         *string `json:"resolution_id,omitempty"`
	SprintID             *string `json:"sprint_id,omitempty"`
	StatusID             string  `json:"status_id"`
	StrategicParentID    *string `json:"strategic_parent_id,omitempty"`
	UserID               *string `json:"user_id,omitempty"`
	UserRefID            string  `json:"user_ref_id"`
	VerificationDuration int64   `json:"verification_duration"`
	VersionIds           *string `json:"version_ids,omitempty"`
}

// TableName returns the SQL table name for JiraIssue and satifies the Model interface
func (t *JiraIssue) TableName() string {
	return JiraIssueTableName
}

// ToCSV will serialize the JiraIssue instance to a CSV compatible array of strings
func (t *JiraIssue) ToCSV() []string {
	return []string{
		t.ID,
		t.CalculateChecksum(),
		t.IssueID,
		t.ProjectID,
		t.IssueTypeID,
		toCSVString(t.UserID),
		toCSVString(t.AssigneeID),
		toCSVString(t.PriorityID),
		t.StatusID,
		toCSVString(t.ResolutionID),
		toCSVString(t.FixVersionIds),
		toCSVString(t.VersionIds),
		toCSVString(t.Environment),
		toCSVString(t.ComponentIds),
		toCSVString(t.LabelIds),
		toCSVString(t.DuedateAt),
		toCSVString(t.PlannedStartAt),
		toCSVString(t.PlannedEndAt),
		t.Key,
		toCSVString(t.CustomFieldIds),
		toCSVString(t.SprintID),
		toCSVString(t.EpicID),
		toCSVString(t.ParentID),
		toCSVString(t.StrategicParentID),
		toCSVString(t.InProgressCount),
		toCSVString(t.ReopenCount),
		toCSVString(t.InProgressDuration),
		toCSVString(t.VerificationDuration),
		toCSVString(t.InProgressStartAt),
		t.CustomerID,
		t.RefID,
		t.UserRefID,
		toCSVString(t.Cost),
	}
}

// WriteCSV will serialize the JiraIssue instance to the writer as CSV and satisfies the CSVWriter interface
func (t *JiraIssue) WriteCSV(w *csv.Writer) error {
	return w.Write(t.ToCSV())
}

// WriteJSON will serialize the JiraIssue instance to the writer as JSON and satisfies the JSONWriter interface
func (t *JiraIssue) WriteJSON(w io.Writer, indent ...bool) error {
	if indent != nil && len(indent) > 0 {
		buf, err := json.MarshalIndent(t, "", "\t")
		if err != nil {
			return err
		}
		if _, err := w.Write(buf); err != nil {
			return err
		}
		if _, err := w.Write([]byte("\n")); err != nil {
			return err
		}
		return nil
	}
	buf, err := json.Marshal(t)
	if err != nil {
		return nil
	}
	if _, err := w.Write(buf); err != nil {
		return err
	}
	if _, err := w.Write([]byte("\n")); err != nil {
		return err
	}
	return nil
}

// NewJiraIssueReader creates a JSON reader which can read in JiraIssue objects serialized as JSON either as an array, single object or json new lines
// and writes each JiraIssue to the channel provided
func NewJiraIssueReader(r io.Reader, ch chan<- JiraIssue) error {
	return orm.Deserialize(r, func(buf json.RawMessage) error {
		dec := json.NewDecoder(bytes.NewBuffer(buf))
		e := JiraIssue{}
		if err := dec.Decode(&e); err != nil {
			return err
		}
		ch <- e
		return nil
	})
}

// NewCSVJiraIssueReaderDir will read the reader as CSV and emit each record to the channel provided
func NewCSVJiraIssueReader(r io.Reader, ch chan<- JiraIssue) error {
	cr := csv.NewReader(r)
	for {
		record, err := cr.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		ch <- JiraIssue{
			ID:                   record[0],
			Checksum:             fromStringPointer(record[1]),
			IssueID:              record[2],
			ProjectID:            record[3],
			IssueTypeID:          record[4],
			UserID:               fromStringPointer(record[5]),
			AssigneeID:           fromStringPointer(record[6]),
			PriorityID:           fromStringPointer(record[7]),
			StatusID:             record[8],
			ResolutionID:         fromStringPointer(record[9]),
			FixVersionIds:        fromStringPointer(record[10]),
			VersionIds:           fromStringPointer(record[11]),
			Environment:          fromStringPointer(record[12]),
			ComponentIds:         fromStringPointer(record[13]),
			LabelIds:             fromStringPointer(record[14]),
			DuedateAt:            fromCSVInt64Pointer(record[15]),
			PlannedStartAt:       fromCSVInt64Pointer(record[16]),
			PlannedEndAt:         fromCSVInt64Pointer(record[17]),
			Key:                  record[18],
			CustomFieldIds:       fromStringPointer(record[19]),
			SprintID:             fromStringPointer(record[20]),
			EpicID:               fromStringPointer(record[21]),
			ParentID:             fromStringPointer(record[22]),
			StrategicParentID:    fromStringPointer(record[23]),
			InProgressCount:      fromCSVInt32(record[24]),
			ReopenCount:          fromCSVInt32(record[25]),
			InProgressDuration:   fromCSVInt64(record[26]),
			VerificationDuration: fromCSVInt64(record[27]),
			InProgressStartAt:    fromCSVInt64Pointer(record[28]),
			CustomerID:           record[29],
			RefID:                record[30],
			UserRefID:            record[31],
			Cost:                 fromCSVFloat64(record[32]),
		}
	}
	return nil
}

// NewCSVJiraIssueReaderFile will read the file as a CSV and emit each record to the channel provided
func NewCSVJiraIssueReaderFile(fp string, ch chan<- JiraIssue) error {
	f, err := os.Open(fp)
	if err != nil {
		return fmt.Errorf("error opening CSV file at %s. %v", fp, err)
	}
	var fc io.ReadCloser = f
	if filepath.Ext(fp) == ".gz" {
		gr, err := gzip.NewReader(f)
		if err != nil {
			return fmt.Errorf("error opening CSV file at %s. %v", fp, err)
		}
		fc = gr
	}
	defer f.Close()
	defer fc.Close()
	return NewCSVJiraIssueReader(fc, ch)
}

// NewCSVJiraIssueReaderDir will read the jira_issue.csv.gz file as a CSV and emit each record to the channel provided
func NewCSVJiraIssueReaderDir(dir string, ch chan<- JiraIssue) error {
	return NewCSVJiraIssueReaderFile(filepath.Join(dir, "jira_issue.csv.gz"), ch)
}

// JiraIssueCSVDeduper is a function callback which takes the existing value (a) and the new value (b)
// and the return value should be the one to use (or a new one, if applicable). return nil
// to skip processing of this record
type JiraIssueCSVDeduper func(a JiraIssue, b JiraIssue) *JiraIssue

// JiraIssueCSVDedupeDisabled is set on whether the CSV writer should de-dupe values by key
var JiraIssueCSVDedupeDisabled bool

// NewJiraIssueCSVWriterSize creates a batch writer that will write each JiraIssue into a CSV file
// this method will automatically de-duplicate entries using the primary key. if the checksum
// for a newer item with the same primary key doesn't match a previously sent item, the newer
// one will be used
func NewJiraIssueCSVWriterSize(w io.Writer, size int, dedupers ...JiraIssueCSVDeduper) (chan JiraIssue, chan bool, error) {
	cw := csv.NewWriter(w)
	ch := make(chan JiraIssue, size)
	done := make(chan bool)
	go func() {
		defer func() { done <- true }()
		dodedupe := !JiraIssueCSVDedupeDisabled
		var kv map[string]*JiraIssue
		var deduper JiraIssueCSVDeduper
		if dedupers != nil && len(dedupers) > 0 {
			deduper = dedupers[0]
			dodedupe = true
		}
		if dodedupe {
			kv = make(map[string]*JiraIssue)
		}
		for c := range ch {
			if dodedupe {
				// get the address and then make a copy so that
				// we mutate on the copy and store it not the source
				cp := &c
				e := *cp
				pk := e.ID
				v := kv[pk]
				if v == nil {
					kv[pk] = &e
					continue
				}
				if deduper != nil {
					r := deduper(e, *v)
					if r != nil {
						kv[pk] = r
					}
					continue
				}
				if v.CalculateChecksum() != e.CalculateChecksum() {
					kv[pk] = &e
					continue
				}
			} else {
				// if not de-duping, just immediately write to CSV
				c.WriteCSV(cw)
			}
		}
		if dodedupe {
			for _, e := range kv {
				e.WriteCSV(cw)
			}
		}
		cw.Flush()
	}()
	return ch, done, nil
}

// JiraIssueCSVDefaultSize is the default channel buffer size if not provided
var JiraIssueCSVDefaultSize = 100

// NewJiraIssueCSVWriter creates a batch writer that will write each JiraIssue into a CSV file
func NewJiraIssueCSVWriter(w io.Writer, dedupers ...JiraIssueCSVDeduper) (chan JiraIssue, chan bool, error) {
	return NewJiraIssueCSVWriterSize(w, JiraIssueCSVDefaultSize, dedupers...)
}

// NewJiraIssueCSVWriterDir creates a batch writer that will write each JiraIssue into a CSV file named jira_issue.csv.gz in dir
func NewJiraIssueCSVWriterDir(dir string, dedupers ...JiraIssueCSVDeduper) (chan JiraIssue, chan bool, error) {
	return NewJiraIssueCSVWriterFile(filepath.Join(dir, "jira_issue.csv.gz"), dedupers...)
}

// NewJiraIssueCSVWriterFile creates a batch writer that will write each JiraIssue into a CSV file
func NewJiraIssueCSVWriterFile(fn string, dedupers ...JiraIssueCSVDeduper) (chan JiraIssue, chan bool, error) {
	f, err := os.Create(fn)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening CSV file %s. %v", fn, err)
	}
	var fc io.WriteCloser = f
	if filepath.Ext(fn) == ".gz" {
		w, _ := gzip.NewWriterLevel(f, gzip.BestCompression)
		fc = w
	}
	ch, done, err := NewJiraIssueCSVWriter(fc, dedupers...)
	if err != nil {
		fc.Close()
		f.Close()
		return nil, nil, fmt.Errorf("error creating CSV writer for %s. %v", fn, err)
	}
	sdone := make(chan bool)
	go func() {
		// wait for our writer to finish
		<-done
		// close our files
		fc.Close()
		f.Close()
		// signal our delegate channel
		sdone <- true
	}()
	return ch, sdone, nil
}

type JiraIssueDBAction func(ctx context.Context, db DB, record JiraIssue) error

// NewJiraIssueDBWriterSize creates a DB writer that will write each issue into the DB
func NewJiraIssueDBWriterSize(ctx context.Context, db DB, errors chan<- error, size int, actions ...JiraIssueDBAction) (chan JiraIssue, chan bool, error) {
	ch := make(chan JiraIssue, size)
	done := make(chan bool)
	var action JiraIssueDBAction
	if actions != nil && len(actions) > 0 {
		action = actions[0]
	}
	go func() {
		defer func() { done <- true }()
		for e := range ch {
			if action != nil {
				if err := action(ctx, db, e); err != nil {
					errors <- err
				}
			} else {
				if _, _, err := e.DBUpsert(ctx, db); err != nil {
					errors <- err
				}
			}
		}
	}()
	return ch, done, nil
}

// NewJiraIssueDBWriter creates a DB writer that will write each issue into the DB
func NewJiraIssueDBWriter(ctx context.Context, db DB, errors chan<- error, actions ...JiraIssueDBAction) (chan JiraIssue, chan bool, error) {
	return NewJiraIssueDBWriterSize(ctx, db, errors, 100, actions...)
}

// JiraIssueColumnID is the ID SQL column name for the JiraIssue table
const JiraIssueColumnID = "id"

// JiraIssueEscapedColumnID is the escaped ID SQL column name for the JiraIssue table
const JiraIssueEscapedColumnID = "`id`"

// JiraIssueColumnChecksum is the Checksum SQL column name for the JiraIssue table
const JiraIssueColumnChecksum = "checksum"

// JiraIssueEscapedColumnChecksum is the escaped Checksum SQL column name for the JiraIssue table
const JiraIssueEscapedColumnChecksum = "`checksum`"

// JiraIssueColumnIssueID is the IssueID SQL column name for the JiraIssue table
const JiraIssueColumnIssueID = "issue_id"

// JiraIssueEscapedColumnIssueID is the escaped IssueID SQL column name for the JiraIssue table
const JiraIssueEscapedColumnIssueID = "`issue_id`"

// JiraIssueColumnProjectID is the ProjectID SQL column name for the JiraIssue table
const JiraIssueColumnProjectID = "project_id"

// JiraIssueEscapedColumnProjectID is the escaped ProjectID SQL column name for the JiraIssue table
const JiraIssueEscapedColumnProjectID = "`project_id`"

// JiraIssueColumnIssueTypeID is the IssueTypeID SQL column name for the JiraIssue table
const JiraIssueColumnIssueTypeID = "issue_type_id"

// JiraIssueEscapedColumnIssueTypeID is the escaped IssueTypeID SQL column name for the JiraIssue table
const JiraIssueEscapedColumnIssueTypeID = "`issue_type_id`"

// JiraIssueColumnUserID is the UserID SQL column name for the JiraIssue table
const JiraIssueColumnUserID = "user_id"

// JiraIssueEscapedColumnUserID is the escaped UserID SQL column name for the JiraIssue table
const JiraIssueEscapedColumnUserID = "`user_id`"

// JiraIssueColumnAssigneeID is the AssigneeID SQL column name for the JiraIssue table
const JiraIssueColumnAssigneeID = "assignee_id"

// JiraIssueEscapedColumnAssigneeID is the escaped AssigneeID SQL column name for the JiraIssue table
const JiraIssueEscapedColumnAssigneeID = "`assignee_id`"

// JiraIssueColumnPriorityID is the PriorityID SQL column name for the JiraIssue table
const JiraIssueColumnPriorityID = "priority_id"

// JiraIssueEscapedColumnPriorityID is the escaped PriorityID SQL column name for the JiraIssue table
const JiraIssueEscapedColumnPriorityID = "`priority_id`"

// JiraIssueColumnStatusID is the StatusID SQL column name for the JiraIssue table
const JiraIssueColumnStatusID = "status_id"

// JiraIssueEscapedColumnStatusID is the escaped StatusID SQL column name for the JiraIssue table
const JiraIssueEscapedColumnStatusID = "`status_id`"

// JiraIssueColumnResolutionID is the ResolutionID SQL column name for the JiraIssue table
const JiraIssueColumnResolutionID = "resolution_id"

// JiraIssueEscapedColumnResolutionID is the escaped ResolutionID SQL column name for the JiraIssue table
const JiraIssueEscapedColumnResolutionID = "`resolution_id`"

// JiraIssueColumnFixVersionIds is the FixVersionIds SQL column name for the JiraIssue table
const JiraIssueColumnFixVersionIds = "fix_version_ids"

// JiraIssueEscapedColumnFixVersionIds is the escaped FixVersionIds SQL column name for the JiraIssue table
const JiraIssueEscapedColumnFixVersionIds = "`fix_version_ids`"

// JiraIssueColumnVersionIds is the VersionIds SQL column name for the JiraIssue table
const JiraIssueColumnVersionIds = "version_ids"

// JiraIssueEscapedColumnVersionIds is the escaped VersionIds SQL column name for the JiraIssue table
const JiraIssueEscapedColumnVersionIds = "`version_ids`"

// JiraIssueColumnEnvironment is the Environment SQL column name for the JiraIssue table
const JiraIssueColumnEnvironment = "environment"

// JiraIssueEscapedColumnEnvironment is the escaped Environment SQL column name for the JiraIssue table
const JiraIssueEscapedColumnEnvironment = "`environment`"

// JiraIssueColumnComponentIds is the ComponentIds SQL column name for the JiraIssue table
const JiraIssueColumnComponentIds = "component_ids"

// JiraIssueEscapedColumnComponentIds is the escaped ComponentIds SQL column name for the JiraIssue table
const JiraIssueEscapedColumnComponentIds = "`component_ids`"

// JiraIssueColumnLabelIds is the LabelIds SQL column name for the JiraIssue table
const JiraIssueColumnLabelIds = "label_ids"

// JiraIssueEscapedColumnLabelIds is the escaped LabelIds SQL column name for the JiraIssue table
const JiraIssueEscapedColumnLabelIds = "`label_ids`"

// JiraIssueColumnDuedateAt is the DuedateAt SQL column name for the JiraIssue table
const JiraIssueColumnDuedateAt = "duedate_at"

// JiraIssueEscapedColumnDuedateAt is the escaped DuedateAt SQL column name for the JiraIssue table
const JiraIssueEscapedColumnDuedateAt = "`duedate_at`"

// JiraIssueColumnPlannedStartAt is the PlannedStartAt SQL column name for the JiraIssue table
const JiraIssueColumnPlannedStartAt = "planned_start_at"

// JiraIssueEscapedColumnPlannedStartAt is the escaped PlannedStartAt SQL column name for the JiraIssue table
const JiraIssueEscapedColumnPlannedStartAt = "`planned_start_at`"

// JiraIssueColumnPlannedEndAt is the PlannedEndAt SQL column name for the JiraIssue table
const JiraIssueColumnPlannedEndAt = "planned_end_at"

// JiraIssueEscapedColumnPlannedEndAt is the escaped PlannedEndAt SQL column name for the JiraIssue table
const JiraIssueEscapedColumnPlannedEndAt = "`planned_end_at`"

// JiraIssueColumnKey is the Key SQL column name for the JiraIssue table
const JiraIssueColumnKey = "key"

// JiraIssueEscapedColumnKey is the escaped Key SQL column name for the JiraIssue table
const JiraIssueEscapedColumnKey = "`key`"

// JiraIssueColumnCustomFieldIds is the CustomFieldIds SQL column name for the JiraIssue table
const JiraIssueColumnCustomFieldIds = "custom_field_ids"

// JiraIssueEscapedColumnCustomFieldIds is the escaped CustomFieldIds SQL column name for the JiraIssue table
const JiraIssueEscapedColumnCustomFieldIds = "`custom_field_ids`"

// JiraIssueColumnSprintID is the SprintID SQL column name for the JiraIssue table
const JiraIssueColumnSprintID = "sprint_id"

// JiraIssueEscapedColumnSprintID is the escaped SprintID SQL column name for the JiraIssue table
const JiraIssueEscapedColumnSprintID = "`sprint_id`"

// JiraIssueColumnEpicID is the EpicID SQL column name for the JiraIssue table
const JiraIssueColumnEpicID = "epic_id"

// JiraIssueEscapedColumnEpicID is the escaped EpicID SQL column name for the JiraIssue table
const JiraIssueEscapedColumnEpicID = "`epic_id`"

// JiraIssueColumnParentID is the ParentID SQL column name for the JiraIssue table
const JiraIssueColumnParentID = "parent_id"

// JiraIssueEscapedColumnParentID is the escaped ParentID SQL column name for the JiraIssue table
const JiraIssueEscapedColumnParentID = "`parent_id`"

// JiraIssueColumnStrategicParentID is the StrategicParentID SQL column name for the JiraIssue table
const JiraIssueColumnStrategicParentID = "strategic_parent_id"

// JiraIssueEscapedColumnStrategicParentID is the escaped StrategicParentID SQL column name for the JiraIssue table
const JiraIssueEscapedColumnStrategicParentID = "`strategic_parent_id`"

// JiraIssueColumnInProgressCount is the InProgressCount SQL column name for the JiraIssue table
const JiraIssueColumnInProgressCount = "in_progress_count"

// JiraIssueEscapedColumnInProgressCount is the escaped InProgressCount SQL column name for the JiraIssue table
const JiraIssueEscapedColumnInProgressCount = "`in_progress_count`"

// JiraIssueColumnReopenCount is the ReopenCount SQL column name for the JiraIssue table
const JiraIssueColumnReopenCount = "reopen_count"

// JiraIssueEscapedColumnReopenCount is the escaped ReopenCount SQL column name for the JiraIssue table
const JiraIssueEscapedColumnReopenCount = "`reopen_count`"

// JiraIssueColumnInProgressDuration is the InProgressDuration SQL column name for the JiraIssue table
const JiraIssueColumnInProgressDuration = "in_progress_duration"

// JiraIssueEscapedColumnInProgressDuration is the escaped InProgressDuration SQL column name for the JiraIssue table
const JiraIssueEscapedColumnInProgressDuration = "`in_progress_duration`"

// JiraIssueColumnVerificationDuration is the VerificationDuration SQL column name for the JiraIssue table
const JiraIssueColumnVerificationDuration = "verification_duration"

// JiraIssueEscapedColumnVerificationDuration is the escaped VerificationDuration SQL column name for the JiraIssue table
const JiraIssueEscapedColumnVerificationDuration = "`verification_duration`"

// JiraIssueColumnInProgressStartAt is the InProgressStartAt SQL column name for the JiraIssue table
const JiraIssueColumnInProgressStartAt = "in_progress_start_at"

// JiraIssueEscapedColumnInProgressStartAt is the escaped InProgressStartAt SQL column name for the JiraIssue table
const JiraIssueEscapedColumnInProgressStartAt = "`in_progress_start_at`"

// JiraIssueColumnCustomerID is the CustomerID SQL column name for the JiraIssue table
const JiraIssueColumnCustomerID = "customer_id"

// JiraIssueEscapedColumnCustomerID is the escaped CustomerID SQL column name for the JiraIssue table
const JiraIssueEscapedColumnCustomerID = "`customer_id`"

// JiraIssueColumnRefID is the RefID SQL column name for the JiraIssue table
const JiraIssueColumnRefID = "ref_id"

// JiraIssueEscapedColumnRefID is the escaped RefID SQL column name for the JiraIssue table
const JiraIssueEscapedColumnRefID = "`ref_id`"

// JiraIssueColumnUserRefID is the UserRefID SQL column name for the JiraIssue table
const JiraIssueColumnUserRefID = "user_ref_id"

// JiraIssueEscapedColumnUserRefID is the escaped UserRefID SQL column name for the JiraIssue table
const JiraIssueEscapedColumnUserRefID = "`user_ref_id`"

// JiraIssueColumnCost is the Cost SQL column name for the JiraIssue table
const JiraIssueColumnCost = "cost"

// JiraIssueEscapedColumnCost is the escaped Cost SQL column name for the JiraIssue table
const JiraIssueEscapedColumnCost = "`cost`"

// GetID will return the JiraIssue ID value
func (t *JiraIssue) GetID() string {
	return t.ID
}

// SetID will set the JiraIssue ID value
func (t *JiraIssue) SetID(v string) {
	t.ID = v
}

// FindJiraIssueByID will find a JiraIssue by ID
func FindJiraIssueByID(ctx context.Context, db DB, value string) (*JiraIssue, error) {
	q := "SELECT * FROM `jira_issue` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _IssueID sql.NullString
	var _ProjectID sql.NullString
	var _IssueTypeID sql.NullString
	var _UserID sql.NullString
	var _AssigneeID sql.NullString
	var _PriorityID sql.NullString
	var _StatusID sql.NullString
	var _ResolutionID sql.NullString
	var _FixVersionIds sql.NullString
	var _VersionIds sql.NullString
	var _Environment sql.NullString
	var _ComponentIds sql.NullString
	var _LabelIds sql.NullString
	var _DuedateAt sql.NullInt64
	var _PlannedStartAt sql.NullInt64
	var _PlannedEndAt sql.NullInt64
	var _Key sql.NullString
	var _CustomFieldIds sql.NullString
	var _SprintID sql.NullString
	var _EpicID sql.NullString
	var _ParentID sql.NullString
	var _StrategicParentID sql.NullString
	var _InProgressCount sql.NullInt64
	var _ReopenCount sql.NullInt64
	var _InProgressDuration sql.NullInt64
	var _VerificationDuration sql.NullInt64
	var _InProgressStartAt sql.NullInt64
	var _CustomerID sql.NullString
	var _RefID sql.NullString
	var _UserRefID sql.NullString
	var _Cost sql.NullFloat64
	err := db.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_IssueID,
		&_ProjectID,
		&_IssueTypeID,
		&_UserID,
		&_AssigneeID,
		&_PriorityID,
		&_StatusID,
		&_ResolutionID,
		&_FixVersionIds,
		&_VersionIds,
		&_Environment,
		&_ComponentIds,
		&_LabelIds,
		&_DuedateAt,
		&_PlannedStartAt,
		&_PlannedEndAt,
		&_Key,
		&_CustomFieldIds,
		&_SprintID,
		&_EpicID,
		&_ParentID,
		&_StrategicParentID,
		&_InProgressCount,
		&_ReopenCount,
		&_InProgressDuration,
		&_VerificationDuration,
		&_InProgressStartAt,
		&_CustomerID,
		&_RefID,
		&_UserRefID,
		&_Cost,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &JiraIssue{}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _Checksum.Valid {
		t.SetChecksum(_Checksum.String)
	}
	if _IssueID.Valid {
		t.SetIssueID(_IssueID.String)
	}
	if _ProjectID.Valid {
		t.SetProjectID(_ProjectID.String)
	}
	if _IssueTypeID.Valid {
		t.SetIssueTypeID(_IssueTypeID.String)
	}
	if _UserID.Valid {
		t.SetUserID(_UserID.String)
	}
	if _AssigneeID.Valid {
		t.SetAssigneeID(_AssigneeID.String)
	}
	if _PriorityID.Valid {
		t.SetPriorityID(_PriorityID.String)
	}
	if _StatusID.Valid {
		t.SetStatusID(_StatusID.String)
	}
	if _ResolutionID.Valid {
		t.SetResolutionID(_ResolutionID.String)
	}
	if _FixVersionIds.Valid {
		t.SetFixVersionIds(_FixVersionIds.String)
	}
	if _VersionIds.Valid {
		t.SetVersionIds(_VersionIds.String)
	}
	if _Environment.Valid {
		t.SetEnvironment(_Environment.String)
	}
	if _ComponentIds.Valid {
		t.SetComponentIds(_ComponentIds.String)
	}
	if _LabelIds.Valid {
		t.SetLabelIds(_LabelIds.String)
	}
	if _DuedateAt.Valid {
		t.SetDuedateAt(_DuedateAt.Int64)
	}
	if _PlannedStartAt.Valid {
		t.SetPlannedStartAt(_PlannedStartAt.Int64)
	}
	if _PlannedEndAt.Valid {
		t.SetPlannedEndAt(_PlannedEndAt.Int64)
	}
	if _Key.Valid {
		t.SetKey(_Key.String)
	}
	if _CustomFieldIds.Valid {
		t.SetCustomFieldIds(_CustomFieldIds.String)
	}
	if _SprintID.Valid {
		t.SetSprintID(_SprintID.String)
	}
	if _EpicID.Valid {
		t.SetEpicID(_EpicID.String)
	}
	if _ParentID.Valid {
		t.SetParentID(_ParentID.String)
	}
	if _StrategicParentID.Valid {
		t.SetStrategicParentID(_StrategicParentID.String)
	}
	if _InProgressCount.Valid {
		t.SetInProgressCount(int32(_InProgressCount.Int64))
	}
	if _ReopenCount.Valid {
		t.SetReopenCount(int32(_ReopenCount.Int64))
	}
	if _InProgressDuration.Valid {
		t.SetInProgressDuration(_InProgressDuration.Int64)
	}
	if _VerificationDuration.Valid {
		t.SetVerificationDuration(_VerificationDuration.Int64)
	}
	if _InProgressStartAt.Valid {
		t.SetInProgressStartAt(_InProgressStartAt.Int64)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _RefID.Valid {
		t.SetRefID(_RefID.String)
	}
	if _UserRefID.Valid {
		t.SetUserRefID(_UserRefID.String)
	}
	if _Cost.Valid {
		t.SetCost(_Cost.Float64)
	}
	return t, nil
}

// FindJiraIssueByIDTx will find a JiraIssue by ID using the provided transaction
func FindJiraIssueByIDTx(ctx context.Context, tx Tx, value string) (*JiraIssue, error) {
	q := "SELECT * FROM `jira_issue` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _IssueID sql.NullString
	var _ProjectID sql.NullString
	var _IssueTypeID sql.NullString
	var _UserID sql.NullString
	var _AssigneeID sql.NullString
	var _PriorityID sql.NullString
	var _StatusID sql.NullString
	var _ResolutionID sql.NullString
	var _FixVersionIds sql.NullString
	var _VersionIds sql.NullString
	var _Environment sql.NullString
	var _ComponentIds sql.NullString
	var _LabelIds sql.NullString
	var _DuedateAt sql.NullInt64
	var _PlannedStartAt sql.NullInt64
	var _PlannedEndAt sql.NullInt64
	var _Key sql.NullString
	var _CustomFieldIds sql.NullString
	var _SprintID sql.NullString
	var _EpicID sql.NullString
	var _ParentID sql.NullString
	var _StrategicParentID sql.NullString
	var _InProgressCount sql.NullInt64
	var _ReopenCount sql.NullInt64
	var _InProgressDuration sql.NullInt64
	var _VerificationDuration sql.NullInt64
	var _InProgressStartAt sql.NullInt64
	var _CustomerID sql.NullString
	var _RefID sql.NullString
	var _UserRefID sql.NullString
	var _Cost sql.NullFloat64
	err := tx.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_IssueID,
		&_ProjectID,
		&_IssueTypeID,
		&_UserID,
		&_AssigneeID,
		&_PriorityID,
		&_StatusID,
		&_ResolutionID,
		&_FixVersionIds,
		&_VersionIds,
		&_Environment,
		&_ComponentIds,
		&_LabelIds,
		&_DuedateAt,
		&_PlannedStartAt,
		&_PlannedEndAt,
		&_Key,
		&_CustomFieldIds,
		&_SprintID,
		&_EpicID,
		&_ParentID,
		&_StrategicParentID,
		&_InProgressCount,
		&_ReopenCount,
		&_InProgressDuration,
		&_VerificationDuration,
		&_InProgressStartAt,
		&_CustomerID,
		&_RefID,
		&_UserRefID,
		&_Cost,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &JiraIssue{}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _Checksum.Valid {
		t.SetChecksum(_Checksum.String)
	}
	if _IssueID.Valid {
		t.SetIssueID(_IssueID.String)
	}
	if _ProjectID.Valid {
		t.SetProjectID(_ProjectID.String)
	}
	if _IssueTypeID.Valid {
		t.SetIssueTypeID(_IssueTypeID.String)
	}
	if _UserID.Valid {
		t.SetUserID(_UserID.String)
	}
	if _AssigneeID.Valid {
		t.SetAssigneeID(_AssigneeID.String)
	}
	if _PriorityID.Valid {
		t.SetPriorityID(_PriorityID.String)
	}
	if _StatusID.Valid {
		t.SetStatusID(_StatusID.String)
	}
	if _ResolutionID.Valid {
		t.SetResolutionID(_ResolutionID.String)
	}
	if _FixVersionIds.Valid {
		t.SetFixVersionIds(_FixVersionIds.String)
	}
	if _VersionIds.Valid {
		t.SetVersionIds(_VersionIds.String)
	}
	if _Environment.Valid {
		t.SetEnvironment(_Environment.String)
	}
	if _ComponentIds.Valid {
		t.SetComponentIds(_ComponentIds.String)
	}
	if _LabelIds.Valid {
		t.SetLabelIds(_LabelIds.String)
	}
	if _DuedateAt.Valid {
		t.SetDuedateAt(_DuedateAt.Int64)
	}
	if _PlannedStartAt.Valid {
		t.SetPlannedStartAt(_PlannedStartAt.Int64)
	}
	if _PlannedEndAt.Valid {
		t.SetPlannedEndAt(_PlannedEndAt.Int64)
	}
	if _Key.Valid {
		t.SetKey(_Key.String)
	}
	if _CustomFieldIds.Valid {
		t.SetCustomFieldIds(_CustomFieldIds.String)
	}
	if _SprintID.Valid {
		t.SetSprintID(_SprintID.String)
	}
	if _EpicID.Valid {
		t.SetEpicID(_EpicID.String)
	}
	if _ParentID.Valid {
		t.SetParentID(_ParentID.String)
	}
	if _StrategicParentID.Valid {
		t.SetStrategicParentID(_StrategicParentID.String)
	}
	if _InProgressCount.Valid {
		t.SetInProgressCount(int32(_InProgressCount.Int64))
	}
	if _ReopenCount.Valid {
		t.SetReopenCount(int32(_ReopenCount.Int64))
	}
	if _InProgressDuration.Valid {
		t.SetInProgressDuration(_InProgressDuration.Int64)
	}
	if _VerificationDuration.Valid {
		t.SetVerificationDuration(_VerificationDuration.Int64)
	}
	if _InProgressStartAt.Valid {
		t.SetInProgressStartAt(_InProgressStartAt.Int64)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _RefID.Valid {
		t.SetRefID(_RefID.String)
	}
	if _UserRefID.Valid {
		t.SetUserRefID(_UserRefID.String)
	}
	if _Cost.Valid {
		t.SetCost(_Cost.Float64)
	}
	return t, nil
}

// GetChecksum will return the JiraIssue Checksum value
func (t *JiraIssue) GetChecksum() string {
	if t.Checksum == nil {
		return ""
	}
	return *t.Checksum
}

// SetChecksum will set the JiraIssue Checksum value
func (t *JiraIssue) SetChecksum(v string) {
	t.Checksum = &v
}

// GetIssueID will return the JiraIssue IssueID value
func (t *JiraIssue) GetIssueID() string {
	return t.IssueID
}

// SetIssueID will set the JiraIssue IssueID value
func (t *JiraIssue) SetIssueID(v string) {
	t.IssueID = v
}

// FindJiraIssuesByIssueID will find all JiraIssues by the IssueID value
func FindJiraIssuesByIssueID(ctx context.Context, db DB, value string) ([]*JiraIssue, error) {
	q := "SELECT * FROM `jira_issue` WHERE `issue_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraIssue, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _IssueID sql.NullString
		var _ProjectID sql.NullString
		var _IssueTypeID sql.NullString
		var _UserID sql.NullString
		var _AssigneeID sql.NullString
		var _PriorityID sql.NullString
		var _StatusID sql.NullString
		var _ResolutionID sql.NullString
		var _FixVersionIds sql.NullString
		var _VersionIds sql.NullString
		var _Environment sql.NullString
		var _ComponentIds sql.NullString
		var _LabelIds sql.NullString
		var _DuedateAt sql.NullInt64
		var _PlannedStartAt sql.NullInt64
		var _PlannedEndAt sql.NullInt64
		var _Key sql.NullString
		var _CustomFieldIds sql.NullString
		var _SprintID sql.NullString
		var _EpicID sql.NullString
		var _ParentID sql.NullString
		var _StrategicParentID sql.NullString
		var _InProgressCount sql.NullInt64
		var _ReopenCount sql.NullInt64
		var _InProgressDuration sql.NullInt64
		var _VerificationDuration sql.NullInt64
		var _InProgressStartAt sql.NullInt64
		var _CustomerID sql.NullString
		var _RefID sql.NullString
		var _UserRefID sql.NullString
		var _Cost sql.NullFloat64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_IssueID,
			&_ProjectID,
			&_IssueTypeID,
			&_UserID,
			&_AssigneeID,
			&_PriorityID,
			&_StatusID,
			&_ResolutionID,
			&_FixVersionIds,
			&_VersionIds,
			&_Environment,
			&_ComponentIds,
			&_LabelIds,
			&_DuedateAt,
			&_PlannedStartAt,
			&_PlannedEndAt,
			&_Key,
			&_CustomFieldIds,
			&_SprintID,
			&_EpicID,
			&_ParentID,
			&_StrategicParentID,
			&_InProgressCount,
			&_ReopenCount,
			&_InProgressDuration,
			&_VerificationDuration,
			&_InProgressStartAt,
			&_CustomerID,
			&_RefID,
			&_UserRefID,
			&_Cost,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraIssue{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _IssueTypeID.Valid {
			t.SetIssueTypeID(_IssueTypeID.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _AssigneeID.Valid {
			t.SetAssigneeID(_AssigneeID.String)
		}
		if _PriorityID.Valid {
			t.SetPriorityID(_PriorityID.String)
		}
		if _StatusID.Valid {
			t.SetStatusID(_StatusID.String)
		}
		if _ResolutionID.Valid {
			t.SetResolutionID(_ResolutionID.String)
		}
		if _FixVersionIds.Valid {
			t.SetFixVersionIds(_FixVersionIds.String)
		}
		if _VersionIds.Valid {
			t.SetVersionIds(_VersionIds.String)
		}
		if _Environment.Valid {
			t.SetEnvironment(_Environment.String)
		}
		if _ComponentIds.Valid {
			t.SetComponentIds(_ComponentIds.String)
		}
		if _LabelIds.Valid {
			t.SetLabelIds(_LabelIds.String)
		}
		if _DuedateAt.Valid {
			t.SetDuedateAt(_DuedateAt.Int64)
		}
		if _PlannedStartAt.Valid {
			t.SetPlannedStartAt(_PlannedStartAt.Int64)
		}
		if _PlannedEndAt.Valid {
			t.SetPlannedEndAt(_PlannedEndAt.Int64)
		}
		if _Key.Valid {
			t.SetKey(_Key.String)
		}
		if _CustomFieldIds.Valid {
			t.SetCustomFieldIds(_CustomFieldIds.String)
		}
		if _SprintID.Valid {
			t.SetSprintID(_SprintID.String)
		}
		if _EpicID.Valid {
			t.SetEpicID(_EpicID.String)
		}
		if _ParentID.Valid {
			t.SetParentID(_ParentID.String)
		}
		if _StrategicParentID.Valid {
			t.SetStrategicParentID(_StrategicParentID.String)
		}
		if _InProgressCount.Valid {
			t.SetInProgressCount(int32(_InProgressCount.Int64))
		}
		if _ReopenCount.Valid {
			t.SetReopenCount(int32(_ReopenCount.Int64))
		}
		if _InProgressDuration.Valid {
			t.SetInProgressDuration(_InProgressDuration.Int64)
		}
		if _VerificationDuration.Valid {
			t.SetVerificationDuration(_VerificationDuration.Int64)
		}
		if _InProgressStartAt.Valid {
			t.SetInProgressStartAt(_InProgressStartAt.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		if _UserRefID.Valid {
			t.SetUserRefID(_UserRefID.String)
		}
		if _Cost.Valid {
			t.SetCost(_Cost.Float64)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindJiraIssuesByIssueIDTx will find all JiraIssues by the IssueID value using the provided transaction
func FindJiraIssuesByIssueIDTx(ctx context.Context, tx Tx, value string) ([]*JiraIssue, error) {
	q := "SELECT * FROM `jira_issue` WHERE `issue_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraIssue, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _IssueID sql.NullString
		var _ProjectID sql.NullString
		var _IssueTypeID sql.NullString
		var _UserID sql.NullString
		var _AssigneeID sql.NullString
		var _PriorityID sql.NullString
		var _StatusID sql.NullString
		var _ResolutionID sql.NullString
		var _FixVersionIds sql.NullString
		var _VersionIds sql.NullString
		var _Environment sql.NullString
		var _ComponentIds sql.NullString
		var _LabelIds sql.NullString
		var _DuedateAt sql.NullInt64
		var _PlannedStartAt sql.NullInt64
		var _PlannedEndAt sql.NullInt64
		var _Key sql.NullString
		var _CustomFieldIds sql.NullString
		var _SprintID sql.NullString
		var _EpicID sql.NullString
		var _ParentID sql.NullString
		var _StrategicParentID sql.NullString
		var _InProgressCount sql.NullInt64
		var _ReopenCount sql.NullInt64
		var _InProgressDuration sql.NullInt64
		var _VerificationDuration sql.NullInt64
		var _InProgressStartAt sql.NullInt64
		var _CustomerID sql.NullString
		var _RefID sql.NullString
		var _UserRefID sql.NullString
		var _Cost sql.NullFloat64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_IssueID,
			&_ProjectID,
			&_IssueTypeID,
			&_UserID,
			&_AssigneeID,
			&_PriorityID,
			&_StatusID,
			&_ResolutionID,
			&_FixVersionIds,
			&_VersionIds,
			&_Environment,
			&_ComponentIds,
			&_LabelIds,
			&_DuedateAt,
			&_PlannedStartAt,
			&_PlannedEndAt,
			&_Key,
			&_CustomFieldIds,
			&_SprintID,
			&_EpicID,
			&_ParentID,
			&_StrategicParentID,
			&_InProgressCount,
			&_ReopenCount,
			&_InProgressDuration,
			&_VerificationDuration,
			&_InProgressStartAt,
			&_CustomerID,
			&_RefID,
			&_UserRefID,
			&_Cost,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraIssue{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _IssueTypeID.Valid {
			t.SetIssueTypeID(_IssueTypeID.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _AssigneeID.Valid {
			t.SetAssigneeID(_AssigneeID.String)
		}
		if _PriorityID.Valid {
			t.SetPriorityID(_PriorityID.String)
		}
		if _StatusID.Valid {
			t.SetStatusID(_StatusID.String)
		}
		if _ResolutionID.Valid {
			t.SetResolutionID(_ResolutionID.String)
		}
		if _FixVersionIds.Valid {
			t.SetFixVersionIds(_FixVersionIds.String)
		}
		if _VersionIds.Valid {
			t.SetVersionIds(_VersionIds.String)
		}
		if _Environment.Valid {
			t.SetEnvironment(_Environment.String)
		}
		if _ComponentIds.Valid {
			t.SetComponentIds(_ComponentIds.String)
		}
		if _LabelIds.Valid {
			t.SetLabelIds(_LabelIds.String)
		}
		if _DuedateAt.Valid {
			t.SetDuedateAt(_DuedateAt.Int64)
		}
		if _PlannedStartAt.Valid {
			t.SetPlannedStartAt(_PlannedStartAt.Int64)
		}
		if _PlannedEndAt.Valid {
			t.SetPlannedEndAt(_PlannedEndAt.Int64)
		}
		if _Key.Valid {
			t.SetKey(_Key.String)
		}
		if _CustomFieldIds.Valid {
			t.SetCustomFieldIds(_CustomFieldIds.String)
		}
		if _SprintID.Valid {
			t.SetSprintID(_SprintID.String)
		}
		if _EpicID.Valid {
			t.SetEpicID(_EpicID.String)
		}
		if _ParentID.Valid {
			t.SetParentID(_ParentID.String)
		}
		if _StrategicParentID.Valid {
			t.SetStrategicParentID(_StrategicParentID.String)
		}
		if _InProgressCount.Valid {
			t.SetInProgressCount(int32(_InProgressCount.Int64))
		}
		if _ReopenCount.Valid {
			t.SetReopenCount(int32(_ReopenCount.Int64))
		}
		if _InProgressDuration.Valid {
			t.SetInProgressDuration(_InProgressDuration.Int64)
		}
		if _VerificationDuration.Valid {
			t.SetVerificationDuration(_VerificationDuration.Int64)
		}
		if _InProgressStartAt.Valid {
			t.SetInProgressStartAt(_InProgressStartAt.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		if _UserRefID.Valid {
			t.SetUserRefID(_UserRefID.String)
		}
		if _Cost.Valid {
			t.SetCost(_Cost.Float64)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetProjectID will return the JiraIssue ProjectID value
func (t *JiraIssue) GetProjectID() string {
	return t.ProjectID
}

// SetProjectID will set the JiraIssue ProjectID value
func (t *JiraIssue) SetProjectID(v string) {
	t.ProjectID = v
}

// FindJiraIssuesByProjectID will find all JiraIssues by the ProjectID value
func FindJiraIssuesByProjectID(ctx context.Context, db DB, value string) ([]*JiraIssue, error) {
	q := "SELECT * FROM `jira_issue` WHERE `project_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraIssue, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _IssueID sql.NullString
		var _ProjectID sql.NullString
		var _IssueTypeID sql.NullString
		var _UserID sql.NullString
		var _AssigneeID sql.NullString
		var _PriorityID sql.NullString
		var _StatusID sql.NullString
		var _ResolutionID sql.NullString
		var _FixVersionIds sql.NullString
		var _VersionIds sql.NullString
		var _Environment sql.NullString
		var _ComponentIds sql.NullString
		var _LabelIds sql.NullString
		var _DuedateAt sql.NullInt64
		var _PlannedStartAt sql.NullInt64
		var _PlannedEndAt sql.NullInt64
		var _Key sql.NullString
		var _CustomFieldIds sql.NullString
		var _SprintID sql.NullString
		var _EpicID sql.NullString
		var _ParentID sql.NullString
		var _StrategicParentID sql.NullString
		var _InProgressCount sql.NullInt64
		var _ReopenCount sql.NullInt64
		var _InProgressDuration sql.NullInt64
		var _VerificationDuration sql.NullInt64
		var _InProgressStartAt sql.NullInt64
		var _CustomerID sql.NullString
		var _RefID sql.NullString
		var _UserRefID sql.NullString
		var _Cost sql.NullFloat64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_IssueID,
			&_ProjectID,
			&_IssueTypeID,
			&_UserID,
			&_AssigneeID,
			&_PriorityID,
			&_StatusID,
			&_ResolutionID,
			&_FixVersionIds,
			&_VersionIds,
			&_Environment,
			&_ComponentIds,
			&_LabelIds,
			&_DuedateAt,
			&_PlannedStartAt,
			&_PlannedEndAt,
			&_Key,
			&_CustomFieldIds,
			&_SprintID,
			&_EpicID,
			&_ParentID,
			&_StrategicParentID,
			&_InProgressCount,
			&_ReopenCount,
			&_InProgressDuration,
			&_VerificationDuration,
			&_InProgressStartAt,
			&_CustomerID,
			&_RefID,
			&_UserRefID,
			&_Cost,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraIssue{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _IssueTypeID.Valid {
			t.SetIssueTypeID(_IssueTypeID.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _AssigneeID.Valid {
			t.SetAssigneeID(_AssigneeID.String)
		}
		if _PriorityID.Valid {
			t.SetPriorityID(_PriorityID.String)
		}
		if _StatusID.Valid {
			t.SetStatusID(_StatusID.String)
		}
		if _ResolutionID.Valid {
			t.SetResolutionID(_ResolutionID.String)
		}
		if _FixVersionIds.Valid {
			t.SetFixVersionIds(_FixVersionIds.String)
		}
		if _VersionIds.Valid {
			t.SetVersionIds(_VersionIds.String)
		}
		if _Environment.Valid {
			t.SetEnvironment(_Environment.String)
		}
		if _ComponentIds.Valid {
			t.SetComponentIds(_ComponentIds.String)
		}
		if _LabelIds.Valid {
			t.SetLabelIds(_LabelIds.String)
		}
		if _DuedateAt.Valid {
			t.SetDuedateAt(_DuedateAt.Int64)
		}
		if _PlannedStartAt.Valid {
			t.SetPlannedStartAt(_PlannedStartAt.Int64)
		}
		if _PlannedEndAt.Valid {
			t.SetPlannedEndAt(_PlannedEndAt.Int64)
		}
		if _Key.Valid {
			t.SetKey(_Key.String)
		}
		if _CustomFieldIds.Valid {
			t.SetCustomFieldIds(_CustomFieldIds.String)
		}
		if _SprintID.Valid {
			t.SetSprintID(_SprintID.String)
		}
		if _EpicID.Valid {
			t.SetEpicID(_EpicID.String)
		}
		if _ParentID.Valid {
			t.SetParentID(_ParentID.String)
		}
		if _StrategicParentID.Valid {
			t.SetStrategicParentID(_StrategicParentID.String)
		}
		if _InProgressCount.Valid {
			t.SetInProgressCount(int32(_InProgressCount.Int64))
		}
		if _ReopenCount.Valid {
			t.SetReopenCount(int32(_ReopenCount.Int64))
		}
		if _InProgressDuration.Valid {
			t.SetInProgressDuration(_InProgressDuration.Int64)
		}
		if _VerificationDuration.Valid {
			t.SetVerificationDuration(_VerificationDuration.Int64)
		}
		if _InProgressStartAt.Valid {
			t.SetInProgressStartAt(_InProgressStartAt.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		if _UserRefID.Valid {
			t.SetUserRefID(_UserRefID.String)
		}
		if _Cost.Valid {
			t.SetCost(_Cost.Float64)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindJiraIssuesByProjectIDTx will find all JiraIssues by the ProjectID value using the provided transaction
func FindJiraIssuesByProjectIDTx(ctx context.Context, tx Tx, value string) ([]*JiraIssue, error) {
	q := "SELECT * FROM `jira_issue` WHERE `project_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraIssue, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _IssueID sql.NullString
		var _ProjectID sql.NullString
		var _IssueTypeID sql.NullString
		var _UserID sql.NullString
		var _AssigneeID sql.NullString
		var _PriorityID sql.NullString
		var _StatusID sql.NullString
		var _ResolutionID sql.NullString
		var _FixVersionIds sql.NullString
		var _VersionIds sql.NullString
		var _Environment sql.NullString
		var _ComponentIds sql.NullString
		var _LabelIds sql.NullString
		var _DuedateAt sql.NullInt64
		var _PlannedStartAt sql.NullInt64
		var _PlannedEndAt sql.NullInt64
		var _Key sql.NullString
		var _CustomFieldIds sql.NullString
		var _SprintID sql.NullString
		var _EpicID sql.NullString
		var _ParentID sql.NullString
		var _StrategicParentID sql.NullString
		var _InProgressCount sql.NullInt64
		var _ReopenCount sql.NullInt64
		var _InProgressDuration sql.NullInt64
		var _VerificationDuration sql.NullInt64
		var _InProgressStartAt sql.NullInt64
		var _CustomerID sql.NullString
		var _RefID sql.NullString
		var _UserRefID sql.NullString
		var _Cost sql.NullFloat64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_IssueID,
			&_ProjectID,
			&_IssueTypeID,
			&_UserID,
			&_AssigneeID,
			&_PriorityID,
			&_StatusID,
			&_ResolutionID,
			&_FixVersionIds,
			&_VersionIds,
			&_Environment,
			&_ComponentIds,
			&_LabelIds,
			&_DuedateAt,
			&_PlannedStartAt,
			&_PlannedEndAt,
			&_Key,
			&_CustomFieldIds,
			&_SprintID,
			&_EpicID,
			&_ParentID,
			&_StrategicParentID,
			&_InProgressCount,
			&_ReopenCount,
			&_InProgressDuration,
			&_VerificationDuration,
			&_InProgressStartAt,
			&_CustomerID,
			&_RefID,
			&_UserRefID,
			&_Cost,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraIssue{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _IssueTypeID.Valid {
			t.SetIssueTypeID(_IssueTypeID.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _AssigneeID.Valid {
			t.SetAssigneeID(_AssigneeID.String)
		}
		if _PriorityID.Valid {
			t.SetPriorityID(_PriorityID.String)
		}
		if _StatusID.Valid {
			t.SetStatusID(_StatusID.String)
		}
		if _ResolutionID.Valid {
			t.SetResolutionID(_ResolutionID.String)
		}
		if _FixVersionIds.Valid {
			t.SetFixVersionIds(_FixVersionIds.String)
		}
		if _VersionIds.Valid {
			t.SetVersionIds(_VersionIds.String)
		}
		if _Environment.Valid {
			t.SetEnvironment(_Environment.String)
		}
		if _ComponentIds.Valid {
			t.SetComponentIds(_ComponentIds.String)
		}
		if _LabelIds.Valid {
			t.SetLabelIds(_LabelIds.String)
		}
		if _DuedateAt.Valid {
			t.SetDuedateAt(_DuedateAt.Int64)
		}
		if _PlannedStartAt.Valid {
			t.SetPlannedStartAt(_PlannedStartAt.Int64)
		}
		if _PlannedEndAt.Valid {
			t.SetPlannedEndAt(_PlannedEndAt.Int64)
		}
		if _Key.Valid {
			t.SetKey(_Key.String)
		}
		if _CustomFieldIds.Valid {
			t.SetCustomFieldIds(_CustomFieldIds.String)
		}
		if _SprintID.Valid {
			t.SetSprintID(_SprintID.String)
		}
		if _EpicID.Valid {
			t.SetEpicID(_EpicID.String)
		}
		if _ParentID.Valid {
			t.SetParentID(_ParentID.String)
		}
		if _StrategicParentID.Valid {
			t.SetStrategicParentID(_StrategicParentID.String)
		}
		if _InProgressCount.Valid {
			t.SetInProgressCount(int32(_InProgressCount.Int64))
		}
		if _ReopenCount.Valid {
			t.SetReopenCount(int32(_ReopenCount.Int64))
		}
		if _InProgressDuration.Valid {
			t.SetInProgressDuration(_InProgressDuration.Int64)
		}
		if _VerificationDuration.Valid {
			t.SetVerificationDuration(_VerificationDuration.Int64)
		}
		if _InProgressStartAt.Valid {
			t.SetInProgressStartAt(_InProgressStartAt.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		if _UserRefID.Valid {
			t.SetUserRefID(_UserRefID.String)
		}
		if _Cost.Valid {
			t.SetCost(_Cost.Float64)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetIssueTypeID will return the JiraIssue IssueTypeID value
func (t *JiraIssue) GetIssueTypeID() string {
	return t.IssueTypeID
}

// SetIssueTypeID will set the JiraIssue IssueTypeID value
func (t *JiraIssue) SetIssueTypeID(v string) {
	t.IssueTypeID = v
}

// FindJiraIssuesByIssueTypeID will find all JiraIssues by the IssueTypeID value
func FindJiraIssuesByIssueTypeID(ctx context.Context, db DB, value string) ([]*JiraIssue, error) {
	q := "SELECT * FROM `jira_issue` WHERE `issue_type_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraIssue, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _IssueID sql.NullString
		var _ProjectID sql.NullString
		var _IssueTypeID sql.NullString
		var _UserID sql.NullString
		var _AssigneeID sql.NullString
		var _PriorityID sql.NullString
		var _StatusID sql.NullString
		var _ResolutionID sql.NullString
		var _FixVersionIds sql.NullString
		var _VersionIds sql.NullString
		var _Environment sql.NullString
		var _ComponentIds sql.NullString
		var _LabelIds sql.NullString
		var _DuedateAt sql.NullInt64
		var _PlannedStartAt sql.NullInt64
		var _PlannedEndAt sql.NullInt64
		var _Key sql.NullString
		var _CustomFieldIds sql.NullString
		var _SprintID sql.NullString
		var _EpicID sql.NullString
		var _ParentID sql.NullString
		var _StrategicParentID sql.NullString
		var _InProgressCount sql.NullInt64
		var _ReopenCount sql.NullInt64
		var _InProgressDuration sql.NullInt64
		var _VerificationDuration sql.NullInt64
		var _InProgressStartAt sql.NullInt64
		var _CustomerID sql.NullString
		var _RefID sql.NullString
		var _UserRefID sql.NullString
		var _Cost sql.NullFloat64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_IssueID,
			&_ProjectID,
			&_IssueTypeID,
			&_UserID,
			&_AssigneeID,
			&_PriorityID,
			&_StatusID,
			&_ResolutionID,
			&_FixVersionIds,
			&_VersionIds,
			&_Environment,
			&_ComponentIds,
			&_LabelIds,
			&_DuedateAt,
			&_PlannedStartAt,
			&_PlannedEndAt,
			&_Key,
			&_CustomFieldIds,
			&_SprintID,
			&_EpicID,
			&_ParentID,
			&_StrategicParentID,
			&_InProgressCount,
			&_ReopenCount,
			&_InProgressDuration,
			&_VerificationDuration,
			&_InProgressStartAt,
			&_CustomerID,
			&_RefID,
			&_UserRefID,
			&_Cost,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraIssue{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _IssueTypeID.Valid {
			t.SetIssueTypeID(_IssueTypeID.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _AssigneeID.Valid {
			t.SetAssigneeID(_AssigneeID.String)
		}
		if _PriorityID.Valid {
			t.SetPriorityID(_PriorityID.String)
		}
		if _StatusID.Valid {
			t.SetStatusID(_StatusID.String)
		}
		if _ResolutionID.Valid {
			t.SetResolutionID(_ResolutionID.String)
		}
		if _FixVersionIds.Valid {
			t.SetFixVersionIds(_FixVersionIds.String)
		}
		if _VersionIds.Valid {
			t.SetVersionIds(_VersionIds.String)
		}
		if _Environment.Valid {
			t.SetEnvironment(_Environment.String)
		}
		if _ComponentIds.Valid {
			t.SetComponentIds(_ComponentIds.String)
		}
		if _LabelIds.Valid {
			t.SetLabelIds(_LabelIds.String)
		}
		if _DuedateAt.Valid {
			t.SetDuedateAt(_DuedateAt.Int64)
		}
		if _PlannedStartAt.Valid {
			t.SetPlannedStartAt(_PlannedStartAt.Int64)
		}
		if _PlannedEndAt.Valid {
			t.SetPlannedEndAt(_PlannedEndAt.Int64)
		}
		if _Key.Valid {
			t.SetKey(_Key.String)
		}
		if _CustomFieldIds.Valid {
			t.SetCustomFieldIds(_CustomFieldIds.String)
		}
		if _SprintID.Valid {
			t.SetSprintID(_SprintID.String)
		}
		if _EpicID.Valid {
			t.SetEpicID(_EpicID.String)
		}
		if _ParentID.Valid {
			t.SetParentID(_ParentID.String)
		}
		if _StrategicParentID.Valid {
			t.SetStrategicParentID(_StrategicParentID.String)
		}
		if _InProgressCount.Valid {
			t.SetInProgressCount(int32(_InProgressCount.Int64))
		}
		if _ReopenCount.Valid {
			t.SetReopenCount(int32(_ReopenCount.Int64))
		}
		if _InProgressDuration.Valid {
			t.SetInProgressDuration(_InProgressDuration.Int64)
		}
		if _VerificationDuration.Valid {
			t.SetVerificationDuration(_VerificationDuration.Int64)
		}
		if _InProgressStartAt.Valid {
			t.SetInProgressStartAt(_InProgressStartAt.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		if _UserRefID.Valid {
			t.SetUserRefID(_UserRefID.String)
		}
		if _Cost.Valid {
			t.SetCost(_Cost.Float64)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindJiraIssuesByIssueTypeIDTx will find all JiraIssues by the IssueTypeID value using the provided transaction
func FindJiraIssuesByIssueTypeIDTx(ctx context.Context, tx Tx, value string) ([]*JiraIssue, error) {
	q := "SELECT * FROM `jira_issue` WHERE `issue_type_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraIssue, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _IssueID sql.NullString
		var _ProjectID sql.NullString
		var _IssueTypeID sql.NullString
		var _UserID sql.NullString
		var _AssigneeID sql.NullString
		var _PriorityID sql.NullString
		var _StatusID sql.NullString
		var _ResolutionID sql.NullString
		var _FixVersionIds sql.NullString
		var _VersionIds sql.NullString
		var _Environment sql.NullString
		var _ComponentIds sql.NullString
		var _LabelIds sql.NullString
		var _DuedateAt sql.NullInt64
		var _PlannedStartAt sql.NullInt64
		var _PlannedEndAt sql.NullInt64
		var _Key sql.NullString
		var _CustomFieldIds sql.NullString
		var _SprintID sql.NullString
		var _EpicID sql.NullString
		var _ParentID sql.NullString
		var _StrategicParentID sql.NullString
		var _InProgressCount sql.NullInt64
		var _ReopenCount sql.NullInt64
		var _InProgressDuration sql.NullInt64
		var _VerificationDuration sql.NullInt64
		var _InProgressStartAt sql.NullInt64
		var _CustomerID sql.NullString
		var _RefID sql.NullString
		var _UserRefID sql.NullString
		var _Cost sql.NullFloat64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_IssueID,
			&_ProjectID,
			&_IssueTypeID,
			&_UserID,
			&_AssigneeID,
			&_PriorityID,
			&_StatusID,
			&_ResolutionID,
			&_FixVersionIds,
			&_VersionIds,
			&_Environment,
			&_ComponentIds,
			&_LabelIds,
			&_DuedateAt,
			&_PlannedStartAt,
			&_PlannedEndAt,
			&_Key,
			&_CustomFieldIds,
			&_SprintID,
			&_EpicID,
			&_ParentID,
			&_StrategicParentID,
			&_InProgressCount,
			&_ReopenCount,
			&_InProgressDuration,
			&_VerificationDuration,
			&_InProgressStartAt,
			&_CustomerID,
			&_RefID,
			&_UserRefID,
			&_Cost,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraIssue{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _IssueTypeID.Valid {
			t.SetIssueTypeID(_IssueTypeID.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _AssigneeID.Valid {
			t.SetAssigneeID(_AssigneeID.String)
		}
		if _PriorityID.Valid {
			t.SetPriorityID(_PriorityID.String)
		}
		if _StatusID.Valid {
			t.SetStatusID(_StatusID.String)
		}
		if _ResolutionID.Valid {
			t.SetResolutionID(_ResolutionID.String)
		}
		if _FixVersionIds.Valid {
			t.SetFixVersionIds(_FixVersionIds.String)
		}
		if _VersionIds.Valid {
			t.SetVersionIds(_VersionIds.String)
		}
		if _Environment.Valid {
			t.SetEnvironment(_Environment.String)
		}
		if _ComponentIds.Valid {
			t.SetComponentIds(_ComponentIds.String)
		}
		if _LabelIds.Valid {
			t.SetLabelIds(_LabelIds.String)
		}
		if _DuedateAt.Valid {
			t.SetDuedateAt(_DuedateAt.Int64)
		}
		if _PlannedStartAt.Valid {
			t.SetPlannedStartAt(_PlannedStartAt.Int64)
		}
		if _PlannedEndAt.Valid {
			t.SetPlannedEndAt(_PlannedEndAt.Int64)
		}
		if _Key.Valid {
			t.SetKey(_Key.String)
		}
		if _CustomFieldIds.Valid {
			t.SetCustomFieldIds(_CustomFieldIds.String)
		}
		if _SprintID.Valid {
			t.SetSprintID(_SprintID.String)
		}
		if _EpicID.Valid {
			t.SetEpicID(_EpicID.String)
		}
		if _ParentID.Valid {
			t.SetParentID(_ParentID.String)
		}
		if _StrategicParentID.Valid {
			t.SetStrategicParentID(_StrategicParentID.String)
		}
		if _InProgressCount.Valid {
			t.SetInProgressCount(int32(_InProgressCount.Int64))
		}
		if _ReopenCount.Valid {
			t.SetReopenCount(int32(_ReopenCount.Int64))
		}
		if _InProgressDuration.Valid {
			t.SetInProgressDuration(_InProgressDuration.Int64)
		}
		if _VerificationDuration.Valid {
			t.SetVerificationDuration(_VerificationDuration.Int64)
		}
		if _InProgressStartAt.Valid {
			t.SetInProgressStartAt(_InProgressStartAt.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		if _UserRefID.Valid {
			t.SetUserRefID(_UserRefID.String)
		}
		if _Cost.Valid {
			t.SetCost(_Cost.Float64)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetUserID will return the JiraIssue UserID value
func (t *JiraIssue) GetUserID() string {
	if t.UserID == nil {
		return ""
	}
	return *t.UserID
}

// SetUserID will set the JiraIssue UserID value
func (t *JiraIssue) SetUserID(v string) {
	t.UserID = &v
}

// GetAssigneeID will return the JiraIssue AssigneeID value
func (t *JiraIssue) GetAssigneeID() string {
	if t.AssigneeID == nil {
		return ""
	}
	return *t.AssigneeID
}

// SetAssigneeID will set the JiraIssue AssigneeID value
func (t *JiraIssue) SetAssigneeID(v string) {
	t.AssigneeID = &v
}

// GetPriorityID will return the JiraIssue PriorityID value
func (t *JiraIssue) GetPriorityID() string {
	if t.PriorityID == nil {
		return ""
	}
	return *t.PriorityID
}

// SetPriorityID will set the JiraIssue PriorityID value
func (t *JiraIssue) SetPriorityID(v string) {
	t.PriorityID = &v
}

// GetStatusID will return the JiraIssue StatusID value
func (t *JiraIssue) GetStatusID() string {
	return t.StatusID
}

// SetStatusID will set the JiraIssue StatusID value
func (t *JiraIssue) SetStatusID(v string) {
	t.StatusID = v
}

// GetResolutionID will return the JiraIssue ResolutionID value
func (t *JiraIssue) GetResolutionID() string {
	if t.ResolutionID == nil {
		return ""
	}
	return *t.ResolutionID
}

// SetResolutionID will set the JiraIssue ResolutionID value
func (t *JiraIssue) SetResolutionID(v string) {
	t.ResolutionID = &v
}

// GetFixVersionIds will return the JiraIssue FixVersionIds value
func (t *JiraIssue) GetFixVersionIds() string {
	if t.FixVersionIds == nil {
		return ""
	}
	return *t.FixVersionIds
}

// SetFixVersionIds will set the JiraIssue FixVersionIds value
func (t *JiraIssue) SetFixVersionIds(v string) {
	t.FixVersionIds = &v
}

// GetVersionIds will return the JiraIssue VersionIds value
func (t *JiraIssue) GetVersionIds() string {
	if t.VersionIds == nil {
		return ""
	}
	return *t.VersionIds
}

// SetVersionIds will set the JiraIssue VersionIds value
func (t *JiraIssue) SetVersionIds(v string) {
	t.VersionIds = &v
}

// GetEnvironment will return the JiraIssue Environment value
func (t *JiraIssue) GetEnvironment() string {
	if t.Environment == nil {
		return ""
	}
	return *t.Environment
}

// SetEnvironment will set the JiraIssue Environment value
func (t *JiraIssue) SetEnvironment(v string) {
	t.Environment = &v
}

// GetComponentIds will return the JiraIssue ComponentIds value
func (t *JiraIssue) GetComponentIds() string {
	if t.ComponentIds == nil {
		return ""
	}
	return *t.ComponentIds
}

// SetComponentIds will set the JiraIssue ComponentIds value
func (t *JiraIssue) SetComponentIds(v string) {
	t.ComponentIds = &v
}

// GetLabelIds will return the JiraIssue LabelIds value
func (t *JiraIssue) GetLabelIds() string {
	if t.LabelIds == nil {
		return ""
	}
	return *t.LabelIds
}

// SetLabelIds will set the JiraIssue LabelIds value
func (t *JiraIssue) SetLabelIds(v string) {
	t.LabelIds = &v
}

// GetDuedateAt will return the JiraIssue DuedateAt value
func (t *JiraIssue) GetDuedateAt() int64 {
	if t.DuedateAt == nil {
		return int64(0)
	}
	return *t.DuedateAt
}

// SetDuedateAt will set the JiraIssue DuedateAt value
func (t *JiraIssue) SetDuedateAt(v int64) {
	t.DuedateAt = &v
}

// GetPlannedStartAt will return the JiraIssue PlannedStartAt value
func (t *JiraIssue) GetPlannedStartAt() int64 {
	if t.PlannedStartAt == nil {
		return int64(0)
	}
	return *t.PlannedStartAt
}

// SetPlannedStartAt will set the JiraIssue PlannedStartAt value
func (t *JiraIssue) SetPlannedStartAt(v int64) {
	t.PlannedStartAt = &v
}

// GetPlannedEndAt will return the JiraIssue PlannedEndAt value
func (t *JiraIssue) GetPlannedEndAt() int64 {
	if t.PlannedEndAt == nil {
		return int64(0)
	}
	return *t.PlannedEndAt
}

// SetPlannedEndAt will set the JiraIssue PlannedEndAt value
func (t *JiraIssue) SetPlannedEndAt(v int64) {
	t.PlannedEndAt = &v
}

// GetKey will return the JiraIssue Key value
func (t *JiraIssue) GetKey() string {
	return t.Key
}

// SetKey will set the JiraIssue Key value
func (t *JiraIssue) SetKey(v string) {
	t.Key = v
}

// FindJiraIssuesByKey will find all JiraIssues by the Key value
func FindJiraIssuesByKey(ctx context.Context, db DB, value string) ([]*JiraIssue, error) {
	q := "SELECT * FROM `jira_issue` WHERE `key` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraIssue, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _IssueID sql.NullString
		var _ProjectID sql.NullString
		var _IssueTypeID sql.NullString
		var _UserID sql.NullString
		var _AssigneeID sql.NullString
		var _PriorityID sql.NullString
		var _StatusID sql.NullString
		var _ResolutionID sql.NullString
		var _FixVersionIds sql.NullString
		var _VersionIds sql.NullString
		var _Environment sql.NullString
		var _ComponentIds sql.NullString
		var _LabelIds sql.NullString
		var _DuedateAt sql.NullInt64
		var _PlannedStartAt sql.NullInt64
		var _PlannedEndAt sql.NullInt64
		var _Key sql.NullString
		var _CustomFieldIds sql.NullString
		var _SprintID sql.NullString
		var _EpicID sql.NullString
		var _ParentID sql.NullString
		var _StrategicParentID sql.NullString
		var _InProgressCount sql.NullInt64
		var _ReopenCount sql.NullInt64
		var _InProgressDuration sql.NullInt64
		var _VerificationDuration sql.NullInt64
		var _InProgressStartAt sql.NullInt64
		var _CustomerID sql.NullString
		var _RefID sql.NullString
		var _UserRefID sql.NullString
		var _Cost sql.NullFloat64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_IssueID,
			&_ProjectID,
			&_IssueTypeID,
			&_UserID,
			&_AssigneeID,
			&_PriorityID,
			&_StatusID,
			&_ResolutionID,
			&_FixVersionIds,
			&_VersionIds,
			&_Environment,
			&_ComponentIds,
			&_LabelIds,
			&_DuedateAt,
			&_PlannedStartAt,
			&_PlannedEndAt,
			&_Key,
			&_CustomFieldIds,
			&_SprintID,
			&_EpicID,
			&_ParentID,
			&_StrategicParentID,
			&_InProgressCount,
			&_ReopenCount,
			&_InProgressDuration,
			&_VerificationDuration,
			&_InProgressStartAt,
			&_CustomerID,
			&_RefID,
			&_UserRefID,
			&_Cost,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraIssue{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _IssueTypeID.Valid {
			t.SetIssueTypeID(_IssueTypeID.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _AssigneeID.Valid {
			t.SetAssigneeID(_AssigneeID.String)
		}
		if _PriorityID.Valid {
			t.SetPriorityID(_PriorityID.String)
		}
		if _StatusID.Valid {
			t.SetStatusID(_StatusID.String)
		}
		if _ResolutionID.Valid {
			t.SetResolutionID(_ResolutionID.String)
		}
		if _FixVersionIds.Valid {
			t.SetFixVersionIds(_FixVersionIds.String)
		}
		if _VersionIds.Valid {
			t.SetVersionIds(_VersionIds.String)
		}
		if _Environment.Valid {
			t.SetEnvironment(_Environment.String)
		}
		if _ComponentIds.Valid {
			t.SetComponentIds(_ComponentIds.String)
		}
		if _LabelIds.Valid {
			t.SetLabelIds(_LabelIds.String)
		}
		if _DuedateAt.Valid {
			t.SetDuedateAt(_DuedateAt.Int64)
		}
		if _PlannedStartAt.Valid {
			t.SetPlannedStartAt(_PlannedStartAt.Int64)
		}
		if _PlannedEndAt.Valid {
			t.SetPlannedEndAt(_PlannedEndAt.Int64)
		}
		if _Key.Valid {
			t.SetKey(_Key.String)
		}
		if _CustomFieldIds.Valid {
			t.SetCustomFieldIds(_CustomFieldIds.String)
		}
		if _SprintID.Valid {
			t.SetSprintID(_SprintID.String)
		}
		if _EpicID.Valid {
			t.SetEpicID(_EpicID.String)
		}
		if _ParentID.Valid {
			t.SetParentID(_ParentID.String)
		}
		if _StrategicParentID.Valid {
			t.SetStrategicParentID(_StrategicParentID.String)
		}
		if _InProgressCount.Valid {
			t.SetInProgressCount(int32(_InProgressCount.Int64))
		}
		if _ReopenCount.Valid {
			t.SetReopenCount(int32(_ReopenCount.Int64))
		}
		if _InProgressDuration.Valid {
			t.SetInProgressDuration(_InProgressDuration.Int64)
		}
		if _VerificationDuration.Valid {
			t.SetVerificationDuration(_VerificationDuration.Int64)
		}
		if _InProgressStartAt.Valid {
			t.SetInProgressStartAt(_InProgressStartAt.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		if _UserRefID.Valid {
			t.SetUserRefID(_UserRefID.String)
		}
		if _Cost.Valid {
			t.SetCost(_Cost.Float64)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindJiraIssuesByKeyTx will find all JiraIssues by the Key value using the provided transaction
func FindJiraIssuesByKeyTx(ctx context.Context, tx Tx, value string) ([]*JiraIssue, error) {
	q := "SELECT * FROM `jira_issue` WHERE `key` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraIssue, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _IssueID sql.NullString
		var _ProjectID sql.NullString
		var _IssueTypeID sql.NullString
		var _UserID sql.NullString
		var _AssigneeID sql.NullString
		var _PriorityID sql.NullString
		var _StatusID sql.NullString
		var _ResolutionID sql.NullString
		var _FixVersionIds sql.NullString
		var _VersionIds sql.NullString
		var _Environment sql.NullString
		var _ComponentIds sql.NullString
		var _LabelIds sql.NullString
		var _DuedateAt sql.NullInt64
		var _PlannedStartAt sql.NullInt64
		var _PlannedEndAt sql.NullInt64
		var _Key sql.NullString
		var _CustomFieldIds sql.NullString
		var _SprintID sql.NullString
		var _EpicID sql.NullString
		var _ParentID sql.NullString
		var _StrategicParentID sql.NullString
		var _InProgressCount sql.NullInt64
		var _ReopenCount sql.NullInt64
		var _InProgressDuration sql.NullInt64
		var _VerificationDuration sql.NullInt64
		var _InProgressStartAt sql.NullInt64
		var _CustomerID sql.NullString
		var _RefID sql.NullString
		var _UserRefID sql.NullString
		var _Cost sql.NullFloat64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_IssueID,
			&_ProjectID,
			&_IssueTypeID,
			&_UserID,
			&_AssigneeID,
			&_PriorityID,
			&_StatusID,
			&_ResolutionID,
			&_FixVersionIds,
			&_VersionIds,
			&_Environment,
			&_ComponentIds,
			&_LabelIds,
			&_DuedateAt,
			&_PlannedStartAt,
			&_PlannedEndAt,
			&_Key,
			&_CustomFieldIds,
			&_SprintID,
			&_EpicID,
			&_ParentID,
			&_StrategicParentID,
			&_InProgressCount,
			&_ReopenCount,
			&_InProgressDuration,
			&_VerificationDuration,
			&_InProgressStartAt,
			&_CustomerID,
			&_RefID,
			&_UserRefID,
			&_Cost,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraIssue{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _IssueTypeID.Valid {
			t.SetIssueTypeID(_IssueTypeID.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _AssigneeID.Valid {
			t.SetAssigneeID(_AssigneeID.String)
		}
		if _PriorityID.Valid {
			t.SetPriorityID(_PriorityID.String)
		}
		if _StatusID.Valid {
			t.SetStatusID(_StatusID.String)
		}
		if _ResolutionID.Valid {
			t.SetResolutionID(_ResolutionID.String)
		}
		if _FixVersionIds.Valid {
			t.SetFixVersionIds(_FixVersionIds.String)
		}
		if _VersionIds.Valid {
			t.SetVersionIds(_VersionIds.String)
		}
		if _Environment.Valid {
			t.SetEnvironment(_Environment.String)
		}
		if _ComponentIds.Valid {
			t.SetComponentIds(_ComponentIds.String)
		}
		if _LabelIds.Valid {
			t.SetLabelIds(_LabelIds.String)
		}
		if _DuedateAt.Valid {
			t.SetDuedateAt(_DuedateAt.Int64)
		}
		if _PlannedStartAt.Valid {
			t.SetPlannedStartAt(_PlannedStartAt.Int64)
		}
		if _PlannedEndAt.Valid {
			t.SetPlannedEndAt(_PlannedEndAt.Int64)
		}
		if _Key.Valid {
			t.SetKey(_Key.String)
		}
		if _CustomFieldIds.Valid {
			t.SetCustomFieldIds(_CustomFieldIds.String)
		}
		if _SprintID.Valid {
			t.SetSprintID(_SprintID.String)
		}
		if _EpicID.Valid {
			t.SetEpicID(_EpicID.String)
		}
		if _ParentID.Valid {
			t.SetParentID(_ParentID.String)
		}
		if _StrategicParentID.Valid {
			t.SetStrategicParentID(_StrategicParentID.String)
		}
		if _InProgressCount.Valid {
			t.SetInProgressCount(int32(_InProgressCount.Int64))
		}
		if _ReopenCount.Valid {
			t.SetReopenCount(int32(_ReopenCount.Int64))
		}
		if _InProgressDuration.Valid {
			t.SetInProgressDuration(_InProgressDuration.Int64)
		}
		if _VerificationDuration.Valid {
			t.SetVerificationDuration(_VerificationDuration.Int64)
		}
		if _InProgressStartAt.Valid {
			t.SetInProgressStartAt(_InProgressStartAt.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		if _UserRefID.Valid {
			t.SetUserRefID(_UserRefID.String)
		}
		if _Cost.Valid {
			t.SetCost(_Cost.Float64)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetCustomFieldIds will return the JiraIssue CustomFieldIds value
func (t *JiraIssue) GetCustomFieldIds() string {
	if t.CustomFieldIds == nil {
		return ""
	}
	return *t.CustomFieldIds
}

// SetCustomFieldIds will set the JiraIssue CustomFieldIds value
func (t *JiraIssue) SetCustomFieldIds(v string) {
	t.CustomFieldIds = &v
}

// GetSprintID will return the JiraIssue SprintID value
func (t *JiraIssue) GetSprintID() string {
	if t.SprintID == nil {
		return ""
	}
	return *t.SprintID
}

// SetSprintID will set the JiraIssue SprintID value
func (t *JiraIssue) SetSprintID(v string) {
	t.SprintID = &v
}

// GetEpicID will return the JiraIssue EpicID value
func (t *JiraIssue) GetEpicID() string {
	if t.EpicID == nil {
		return ""
	}
	return *t.EpicID
}

// SetEpicID will set the JiraIssue EpicID value
func (t *JiraIssue) SetEpicID(v string) {
	t.EpicID = &v
}

// FindJiraIssuesByEpicID will find all JiraIssues by the EpicID value
func FindJiraIssuesByEpicID(ctx context.Context, db DB, value string) ([]*JiraIssue, error) {
	q := "SELECT * FROM `jira_issue` WHERE `epic_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraIssue, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _IssueID sql.NullString
		var _ProjectID sql.NullString
		var _IssueTypeID sql.NullString
		var _UserID sql.NullString
		var _AssigneeID sql.NullString
		var _PriorityID sql.NullString
		var _StatusID sql.NullString
		var _ResolutionID sql.NullString
		var _FixVersionIds sql.NullString
		var _VersionIds sql.NullString
		var _Environment sql.NullString
		var _ComponentIds sql.NullString
		var _LabelIds sql.NullString
		var _DuedateAt sql.NullInt64
		var _PlannedStartAt sql.NullInt64
		var _PlannedEndAt sql.NullInt64
		var _Key sql.NullString
		var _CustomFieldIds sql.NullString
		var _SprintID sql.NullString
		var _EpicID sql.NullString
		var _ParentID sql.NullString
		var _StrategicParentID sql.NullString
		var _InProgressCount sql.NullInt64
		var _ReopenCount sql.NullInt64
		var _InProgressDuration sql.NullInt64
		var _VerificationDuration sql.NullInt64
		var _InProgressStartAt sql.NullInt64
		var _CustomerID sql.NullString
		var _RefID sql.NullString
		var _UserRefID sql.NullString
		var _Cost sql.NullFloat64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_IssueID,
			&_ProjectID,
			&_IssueTypeID,
			&_UserID,
			&_AssigneeID,
			&_PriorityID,
			&_StatusID,
			&_ResolutionID,
			&_FixVersionIds,
			&_VersionIds,
			&_Environment,
			&_ComponentIds,
			&_LabelIds,
			&_DuedateAt,
			&_PlannedStartAt,
			&_PlannedEndAt,
			&_Key,
			&_CustomFieldIds,
			&_SprintID,
			&_EpicID,
			&_ParentID,
			&_StrategicParentID,
			&_InProgressCount,
			&_ReopenCount,
			&_InProgressDuration,
			&_VerificationDuration,
			&_InProgressStartAt,
			&_CustomerID,
			&_RefID,
			&_UserRefID,
			&_Cost,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraIssue{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _IssueTypeID.Valid {
			t.SetIssueTypeID(_IssueTypeID.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _AssigneeID.Valid {
			t.SetAssigneeID(_AssigneeID.String)
		}
		if _PriorityID.Valid {
			t.SetPriorityID(_PriorityID.String)
		}
		if _StatusID.Valid {
			t.SetStatusID(_StatusID.String)
		}
		if _ResolutionID.Valid {
			t.SetResolutionID(_ResolutionID.String)
		}
		if _FixVersionIds.Valid {
			t.SetFixVersionIds(_FixVersionIds.String)
		}
		if _VersionIds.Valid {
			t.SetVersionIds(_VersionIds.String)
		}
		if _Environment.Valid {
			t.SetEnvironment(_Environment.String)
		}
		if _ComponentIds.Valid {
			t.SetComponentIds(_ComponentIds.String)
		}
		if _LabelIds.Valid {
			t.SetLabelIds(_LabelIds.String)
		}
		if _DuedateAt.Valid {
			t.SetDuedateAt(_DuedateAt.Int64)
		}
		if _PlannedStartAt.Valid {
			t.SetPlannedStartAt(_PlannedStartAt.Int64)
		}
		if _PlannedEndAt.Valid {
			t.SetPlannedEndAt(_PlannedEndAt.Int64)
		}
		if _Key.Valid {
			t.SetKey(_Key.String)
		}
		if _CustomFieldIds.Valid {
			t.SetCustomFieldIds(_CustomFieldIds.String)
		}
		if _SprintID.Valid {
			t.SetSprintID(_SprintID.String)
		}
		if _EpicID.Valid {
			t.SetEpicID(_EpicID.String)
		}
		if _ParentID.Valid {
			t.SetParentID(_ParentID.String)
		}
		if _StrategicParentID.Valid {
			t.SetStrategicParentID(_StrategicParentID.String)
		}
		if _InProgressCount.Valid {
			t.SetInProgressCount(int32(_InProgressCount.Int64))
		}
		if _ReopenCount.Valid {
			t.SetReopenCount(int32(_ReopenCount.Int64))
		}
		if _InProgressDuration.Valid {
			t.SetInProgressDuration(_InProgressDuration.Int64)
		}
		if _VerificationDuration.Valid {
			t.SetVerificationDuration(_VerificationDuration.Int64)
		}
		if _InProgressStartAt.Valid {
			t.SetInProgressStartAt(_InProgressStartAt.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		if _UserRefID.Valid {
			t.SetUserRefID(_UserRefID.String)
		}
		if _Cost.Valid {
			t.SetCost(_Cost.Float64)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindJiraIssuesByEpicIDTx will find all JiraIssues by the EpicID value using the provided transaction
func FindJiraIssuesByEpicIDTx(ctx context.Context, tx Tx, value string) ([]*JiraIssue, error) {
	q := "SELECT * FROM `jira_issue` WHERE `epic_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraIssue, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _IssueID sql.NullString
		var _ProjectID sql.NullString
		var _IssueTypeID sql.NullString
		var _UserID sql.NullString
		var _AssigneeID sql.NullString
		var _PriorityID sql.NullString
		var _StatusID sql.NullString
		var _ResolutionID sql.NullString
		var _FixVersionIds sql.NullString
		var _VersionIds sql.NullString
		var _Environment sql.NullString
		var _ComponentIds sql.NullString
		var _LabelIds sql.NullString
		var _DuedateAt sql.NullInt64
		var _PlannedStartAt sql.NullInt64
		var _PlannedEndAt sql.NullInt64
		var _Key sql.NullString
		var _CustomFieldIds sql.NullString
		var _SprintID sql.NullString
		var _EpicID sql.NullString
		var _ParentID sql.NullString
		var _StrategicParentID sql.NullString
		var _InProgressCount sql.NullInt64
		var _ReopenCount sql.NullInt64
		var _InProgressDuration sql.NullInt64
		var _VerificationDuration sql.NullInt64
		var _InProgressStartAt sql.NullInt64
		var _CustomerID sql.NullString
		var _RefID sql.NullString
		var _UserRefID sql.NullString
		var _Cost sql.NullFloat64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_IssueID,
			&_ProjectID,
			&_IssueTypeID,
			&_UserID,
			&_AssigneeID,
			&_PriorityID,
			&_StatusID,
			&_ResolutionID,
			&_FixVersionIds,
			&_VersionIds,
			&_Environment,
			&_ComponentIds,
			&_LabelIds,
			&_DuedateAt,
			&_PlannedStartAt,
			&_PlannedEndAt,
			&_Key,
			&_CustomFieldIds,
			&_SprintID,
			&_EpicID,
			&_ParentID,
			&_StrategicParentID,
			&_InProgressCount,
			&_ReopenCount,
			&_InProgressDuration,
			&_VerificationDuration,
			&_InProgressStartAt,
			&_CustomerID,
			&_RefID,
			&_UserRefID,
			&_Cost,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraIssue{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _IssueTypeID.Valid {
			t.SetIssueTypeID(_IssueTypeID.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _AssigneeID.Valid {
			t.SetAssigneeID(_AssigneeID.String)
		}
		if _PriorityID.Valid {
			t.SetPriorityID(_PriorityID.String)
		}
		if _StatusID.Valid {
			t.SetStatusID(_StatusID.String)
		}
		if _ResolutionID.Valid {
			t.SetResolutionID(_ResolutionID.String)
		}
		if _FixVersionIds.Valid {
			t.SetFixVersionIds(_FixVersionIds.String)
		}
		if _VersionIds.Valid {
			t.SetVersionIds(_VersionIds.String)
		}
		if _Environment.Valid {
			t.SetEnvironment(_Environment.String)
		}
		if _ComponentIds.Valid {
			t.SetComponentIds(_ComponentIds.String)
		}
		if _LabelIds.Valid {
			t.SetLabelIds(_LabelIds.String)
		}
		if _DuedateAt.Valid {
			t.SetDuedateAt(_DuedateAt.Int64)
		}
		if _PlannedStartAt.Valid {
			t.SetPlannedStartAt(_PlannedStartAt.Int64)
		}
		if _PlannedEndAt.Valid {
			t.SetPlannedEndAt(_PlannedEndAt.Int64)
		}
		if _Key.Valid {
			t.SetKey(_Key.String)
		}
		if _CustomFieldIds.Valid {
			t.SetCustomFieldIds(_CustomFieldIds.String)
		}
		if _SprintID.Valid {
			t.SetSprintID(_SprintID.String)
		}
		if _EpicID.Valid {
			t.SetEpicID(_EpicID.String)
		}
		if _ParentID.Valid {
			t.SetParentID(_ParentID.String)
		}
		if _StrategicParentID.Valid {
			t.SetStrategicParentID(_StrategicParentID.String)
		}
		if _InProgressCount.Valid {
			t.SetInProgressCount(int32(_InProgressCount.Int64))
		}
		if _ReopenCount.Valid {
			t.SetReopenCount(int32(_ReopenCount.Int64))
		}
		if _InProgressDuration.Valid {
			t.SetInProgressDuration(_InProgressDuration.Int64)
		}
		if _VerificationDuration.Valid {
			t.SetVerificationDuration(_VerificationDuration.Int64)
		}
		if _InProgressStartAt.Valid {
			t.SetInProgressStartAt(_InProgressStartAt.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		if _UserRefID.Valid {
			t.SetUserRefID(_UserRefID.String)
		}
		if _Cost.Valid {
			t.SetCost(_Cost.Float64)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetParentID will return the JiraIssue ParentID value
func (t *JiraIssue) GetParentID() string {
	if t.ParentID == nil {
		return ""
	}
	return *t.ParentID
}

// SetParentID will set the JiraIssue ParentID value
func (t *JiraIssue) SetParentID(v string) {
	t.ParentID = &v
}

// FindJiraIssuesByParentID will find all JiraIssues by the ParentID value
func FindJiraIssuesByParentID(ctx context.Context, db DB, value string) ([]*JiraIssue, error) {
	q := "SELECT * FROM `jira_issue` WHERE `parent_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraIssue, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _IssueID sql.NullString
		var _ProjectID sql.NullString
		var _IssueTypeID sql.NullString
		var _UserID sql.NullString
		var _AssigneeID sql.NullString
		var _PriorityID sql.NullString
		var _StatusID sql.NullString
		var _ResolutionID sql.NullString
		var _FixVersionIds sql.NullString
		var _VersionIds sql.NullString
		var _Environment sql.NullString
		var _ComponentIds sql.NullString
		var _LabelIds sql.NullString
		var _DuedateAt sql.NullInt64
		var _PlannedStartAt sql.NullInt64
		var _PlannedEndAt sql.NullInt64
		var _Key sql.NullString
		var _CustomFieldIds sql.NullString
		var _SprintID sql.NullString
		var _EpicID sql.NullString
		var _ParentID sql.NullString
		var _StrategicParentID sql.NullString
		var _InProgressCount sql.NullInt64
		var _ReopenCount sql.NullInt64
		var _InProgressDuration sql.NullInt64
		var _VerificationDuration sql.NullInt64
		var _InProgressStartAt sql.NullInt64
		var _CustomerID sql.NullString
		var _RefID sql.NullString
		var _UserRefID sql.NullString
		var _Cost sql.NullFloat64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_IssueID,
			&_ProjectID,
			&_IssueTypeID,
			&_UserID,
			&_AssigneeID,
			&_PriorityID,
			&_StatusID,
			&_ResolutionID,
			&_FixVersionIds,
			&_VersionIds,
			&_Environment,
			&_ComponentIds,
			&_LabelIds,
			&_DuedateAt,
			&_PlannedStartAt,
			&_PlannedEndAt,
			&_Key,
			&_CustomFieldIds,
			&_SprintID,
			&_EpicID,
			&_ParentID,
			&_StrategicParentID,
			&_InProgressCount,
			&_ReopenCount,
			&_InProgressDuration,
			&_VerificationDuration,
			&_InProgressStartAt,
			&_CustomerID,
			&_RefID,
			&_UserRefID,
			&_Cost,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraIssue{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _IssueTypeID.Valid {
			t.SetIssueTypeID(_IssueTypeID.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _AssigneeID.Valid {
			t.SetAssigneeID(_AssigneeID.String)
		}
		if _PriorityID.Valid {
			t.SetPriorityID(_PriorityID.String)
		}
		if _StatusID.Valid {
			t.SetStatusID(_StatusID.String)
		}
		if _ResolutionID.Valid {
			t.SetResolutionID(_ResolutionID.String)
		}
		if _FixVersionIds.Valid {
			t.SetFixVersionIds(_FixVersionIds.String)
		}
		if _VersionIds.Valid {
			t.SetVersionIds(_VersionIds.String)
		}
		if _Environment.Valid {
			t.SetEnvironment(_Environment.String)
		}
		if _ComponentIds.Valid {
			t.SetComponentIds(_ComponentIds.String)
		}
		if _LabelIds.Valid {
			t.SetLabelIds(_LabelIds.String)
		}
		if _DuedateAt.Valid {
			t.SetDuedateAt(_DuedateAt.Int64)
		}
		if _PlannedStartAt.Valid {
			t.SetPlannedStartAt(_PlannedStartAt.Int64)
		}
		if _PlannedEndAt.Valid {
			t.SetPlannedEndAt(_PlannedEndAt.Int64)
		}
		if _Key.Valid {
			t.SetKey(_Key.String)
		}
		if _CustomFieldIds.Valid {
			t.SetCustomFieldIds(_CustomFieldIds.String)
		}
		if _SprintID.Valid {
			t.SetSprintID(_SprintID.String)
		}
		if _EpicID.Valid {
			t.SetEpicID(_EpicID.String)
		}
		if _ParentID.Valid {
			t.SetParentID(_ParentID.String)
		}
		if _StrategicParentID.Valid {
			t.SetStrategicParentID(_StrategicParentID.String)
		}
		if _InProgressCount.Valid {
			t.SetInProgressCount(int32(_InProgressCount.Int64))
		}
		if _ReopenCount.Valid {
			t.SetReopenCount(int32(_ReopenCount.Int64))
		}
		if _InProgressDuration.Valid {
			t.SetInProgressDuration(_InProgressDuration.Int64)
		}
		if _VerificationDuration.Valid {
			t.SetVerificationDuration(_VerificationDuration.Int64)
		}
		if _InProgressStartAt.Valid {
			t.SetInProgressStartAt(_InProgressStartAt.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		if _UserRefID.Valid {
			t.SetUserRefID(_UserRefID.String)
		}
		if _Cost.Valid {
			t.SetCost(_Cost.Float64)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindJiraIssuesByParentIDTx will find all JiraIssues by the ParentID value using the provided transaction
func FindJiraIssuesByParentIDTx(ctx context.Context, tx Tx, value string) ([]*JiraIssue, error) {
	q := "SELECT * FROM `jira_issue` WHERE `parent_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraIssue, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _IssueID sql.NullString
		var _ProjectID sql.NullString
		var _IssueTypeID sql.NullString
		var _UserID sql.NullString
		var _AssigneeID sql.NullString
		var _PriorityID sql.NullString
		var _StatusID sql.NullString
		var _ResolutionID sql.NullString
		var _FixVersionIds sql.NullString
		var _VersionIds sql.NullString
		var _Environment sql.NullString
		var _ComponentIds sql.NullString
		var _LabelIds sql.NullString
		var _DuedateAt sql.NullInt64
		var _PlannedStartAt sql.NullInt64
		var _PlannedEndAt sql.NullInt64
		var _Key sql.NullString
		var _CustomFieldIds sql.NullString
		var _SprintID sql.NullString
		var _EpicID sql.NullString
		var _ParentID sql.NullString
		var _StrategicParentID sql.NullString
		var _InProgressCount sql.NullInt64
		var _ReopenCount sql.NullInt64
		var _InProgressDuration sql.NullInt64
		var _VerificationDuration sql.NullInt64
		var _InProgressStartAt sql.NullInt64
		var _CustomerID sql.NullString
		var _RefID sql.NullString
		var _UserRefID sql.NullString
		var _Cost sql.NullFloat64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_IssueID,
			&_ProjectID,
			&_IssueTypeID,
			&_UserID,
			&_AssigneeID,
			&_PriorityID,
			&_StatusID,
			&_ResolutionID,
			&_FixVersionIds,
			&_VersionIds,
			&_Environment,
			&_ComponentIds,
			&_LabelIds,
			&_DuedateAt,
			&_PlannedStartAt,
			&_PlannedEndAt,
			&_Key,
			&_CustomFieldIds,
			&_SprintID,
			&_EpicID,
			&_ParentID,
			&_StrategicParentID,
			&_InProgressCount,
			&_ReopenCount,
			&_InProgressDuration,
			&_VerificationDuration,
			&_InProgressStartAt,
			&_CustomerID,
			&_RefID,
			&_UserRefID,
			&_Cost,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraIssue{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _IssueTypeID.Valid {
			t.SetIssueTypeID(_IssueTypeID.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _AssigneeID.Valid {
			t.SetAssigneeID(_AssigneeID.String)
		}
		if _PriorityID.Valid {
			t.SetPriorityID(_PriorityID.String)
		}
		if _StatusID.Valid {
			t.SetStatusID(_StatusID.String)
		}
		if _ResolutionID.Valid {
			t.SetResolutionID(_ResolutionID.String)
		}
		if _FixVersionIds.Valid {
			t.SetFixVersionIds(_FixVersionIds.String)
		}
		if _VersionIds.Valid {
			t.SetVersionIds(_VersionIds.String)
		}
		if _Environment.Valid {
			t.SetEnvironment(_Environment.String)
		}
		if _ComponentIds.Valid {
			t.SetComponentIds(_ComponentIds.String)
		}
		if _LabelIds.Valid {
			t.SetLabelIds(_LabelIds.String)
		}
		if _DuedateAt.Valid {
			t.SetDuedateAt(_DuedateAt.Int64)
		}
		if _PlannedStartAt.Valid {
			t.SetPlannedStartAt(_PlannedStartAt.Int64)
		}
		if _PlannedEndAt.Valid {
			t.SetPlannedEndAt(_PlannedEndAt.Int64)
		}
		if _Key.Valid {
			t.SetKey(_Key.String)
		}
		if _CustomFieldIds.Valid {
			t.SetCustomFieldIds(_CustomFieldIds.String)
		}
		if _SprintID.Valid {
			t.SetSprintID(_SprintID.String)
		}
		if _EpicID.Valid {
			t.SetEpicID(_EpicID.String)
		}
		if _ParentID.Valid {
			t.SetParentID(_ParentID.String)
		}
		if _StrategicParentID.Valid {
			t.SetStrategicParentID(_StrategicParentID.String)
		}
		if _InProgressCount.Valid {
			t.SetInProgressCount(int32(_InProgressCount.Int64))
		}
		if _ReopenCount.Valid {
			t.SetReopenCount(int32(_ReopenCount.Int64))
		}
		if _InProgressDuration.Valid {
			t.SetInProgressDuration(_InProgressDuration.Int64)
		}
		if _VerificationDuration.Valid {
			t.SetVerificationDuration(_VerificationDuration.Int64)
		}
		if _InProgressStartAt.Valid {
			t.SetInProgressStartAt(_InProgressStartAt.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		if _UserRefID.Valid {
			t.SetUserRefID(_UserRefID.String)
		}
		if _Cost.Valid {
			t.SetCost(_Cost.Float64)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetStrategicParentID will return the JiraIssue StrategicParentID value
func (t *JiraIssue) GetStrategicParentID() string {
	if t.StrategicParentID == nil {
		return ""
	}
	return *t.StrategicParentID
}

// SetStrategicParentID will set the JiraIssue StrategicParentID value
func (t *JiraIssue) SetStrategicParentID(v string) {
	t.StrategicParentID = &v
}

// FindJiraIssuesByStrategicParentID will find all JiraIssues by the StrategicParentID value
func FindJiraIssuesByStrategicParentID(ctx context.Context, db DB, value string) ([]*JiraIssue, error) {
	q := "SELECT * FROM `jira_issue` WHERE `strategic_parent_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraIssue, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _IssueID sql.NullString
		var _ProjectID sql.NullString
		var _IssueTypeID sql.NullString
		var _UserID sql.NullString
		var _AssigneeID sql.NullString
		var _PriorityID sql.NullString
		var _StatusID sql.NullString
		var _ResolutionID sql.NullString
		var _FixVersionIds sql.NullString
		var _VersionIds sql.NullString
		var _Environment sql.NullString
		var _ComponentIds sql.NullString
		var _LabelIds sql.NullString
		var _DuedateAt sql.NullInt64
		var _PlannedStartAt sql.NullInt64
		var _PlannedEndAt sql.NullInt64
		var _Key sql.NullString
		var _CustomFieldIds sql.NullString
		var _SprintID sql.NullString
		var _EpicID sql.NullString
		var _ParentID sql.NullString
		var _StrategicParentID sql.NullString
		var _InProgressCount sql.NullInt64
		var _ReopenCount sql.NullInt64
		var _InProgressDuration sql.NullInt64
		var _VerificationDuration sql.NullInt64
		var _InProgressStartAt sql.NullInt64
		var _CustomerID sql.NullString
		var _RefID sql.NullString
		var _UserRefID sql.NullString
		var _Cost sql.NullFloat64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_IssueID,
			&_ProjectID,
			&_IssueTypeID,
			&_UserID,
			&_AssigneeID,
			&_PriorityID,
			&_StatusID,
			&_ResolutionID,
			&_FixVersionIds,
			&_VersionIds,
			&_Environment,
			&_ComponentIds,
			&_LabelIds,
			&_DuedateAt,
			&_PlannedStartAt,
			&_PlannedEndAt,
			&_Key,
			&_CustomFieldIds,
			&_SprintID,
			&_EpicID,
			&_ParentID,
			&_StrategicParentID,
			&_InProgressCount,
			&_ReopenCount,
			&_InProgressDuration,
			&_VerificationDuration,
			&_InProgressStartAt,
			&_CustomerID,
			&_RefID,
			&_UserRefID,
			&_Cost,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraIssue{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _IssueTypeID.Valid {
			t.SetIssueTypeID(_IssueTypeID.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _AssigneeID.Valid {
			t.SetAssigneeID(_AssigneeID.String)
		}
		if _PriorityID.Valid {
			t.SetPriorityID(_PriorityID.String)
		}
		if _StatusID.Valid {
			t.SetStatusID(_StatusID.String)
		}
		if _ResolutionID.Valid {
			t.SetResolutionID(_ResolutionID.String)
		}
		if _FixVersionIds.Valid {
			t.SetFixVersionIds(_FixVersionIds.String)
		}
		if _VersionIds.Valid {
			t.SetVersionIds(_VersionIds.String)
		}
		if _Environment.Valid {
			t.SetEnvironment(_Environment.String)
		}
		if _ComponentIds.Valid {
			t.SetComponentIds(_ComponentIds.String)
		}
		if _LabelIds.Valid {
			t.SetLabelIds(_LabelIds.String)
		}
		if _DuedateAt.Valid {
			t.SetDuedateAt(_DuedateAt.Int64)
		}
		if _PlannedStartAt.Valid {
			t.SetPlannedStartAt(_PlannedStartAt.Int64)
		}
		if _PlannedEndAt.Valid {
			t.SetPlannedEndAt(_PlannedEndAt.Int64)
		}
		if _Key.Valid {
			t.SetKey(_Key.String)
		}
		if _CustomFieldIds.Valid {
			t.SetCustomFieldIds(_CustomFieldIds.String)
		}
		if _SprintID.Valid {
			t.SetSprintID(_SprintID.String)
		}
		if _EpicID.Valid {
			t.SetEpicID(_EpicID.String)
		}
		if _ParentID.Valid {
			t.SetParentID(_ParentID.String)
		}
		if _StrategicParentID.Valid {
			t.SetStrategicParentID(_StrategicParentID.String)
		}
		if _InProgressCount.Valid {
			t.SetInProgressCount(int32(_InProgressCount.Int64))
		}
		if _ReopenCount.Valid {
			t.SetReopenCount(int32(_ReopenCount.Int64))
		}
		if _InProgressDuration.Valid {
			t.SetInProgressDuration(_InProgressDuration.Int64)
		}
		if _VerificationDuration.Valid {
			t.SetVerificationDuration(_VerificationDuration.Int64)
		}
		if _InProgressStartAt.Valid {
			t.SetInProgressStartAt(_InProgressStartAt.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		if _UserRefID.Valid {
			t.SetUserRefID(_UserRefID.String)
		}
		if _Cost.Valid {
			t.SetCost(_Cost.Float64)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindJiraIssuesByStrategicParentIDTx will find all JiraIssues by the StrategicParentID value using the provided transaction
func FindJiraIssuesByStrategicParentIDTx(ctx context.Context, tx Tx, value string) ([]*JiraIssue, error) {
	q := "SELECT * FROM `jira_issue` WHERE `strategic_parent_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraIssue, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _IssueID sql.NullString
		var _ProjectID sql.NullString
		var _IssueTypeID sql.NullString
		var _UserID sql.NullString
		var _AssigneeID sql.NullString
		var _PriorityID sql.NullString
		var _StatusID sql.NullString
		var _ResolutionID sql.NullString
		var _FixVersionIds sql.NullString
		var _VersionIds sql.NullString
		var _Environment sql.NullString
		var _ComponentIds sql.NullString
		var _LabelIds sql.NullString
		var _DuedateAt sql.NullInt64
		var _PlannedStartAt sql.NullInt64
		var _PlannedEndAt sql.NullInt64
		var _Key sql.NullString
		var _CustomFieldIds sql.NullString
		var _SprintID sql.NullString
		var _EpicID sql.NullString
		var _ParentID sql.NullString
		var _StrategicParentID sql.NullString
		var _InProgressCount sql.NullInt64
		var _ReopenCount sql.NullInt64
		var _InProgressDuration sql.NullInt64
		var _VerificationDuration sql.NullInt64
		var _InProgressStartAt sql.NullInt64
		var _CustomerID sql.NullString
		var _RefID sql.NullString
		var _UserRefID sql.NullString
		var _Cost sql.NullFloat64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_IssueID,
			&_ProjectID,
			&_IssueTypeID,
			&_UserID,
			&_AssigneeID,
			&_PriorityID,
			&_StatusID,
			&_ResolutionID,
			&_FixVersionIds,
			&_VersionIds,
			&_Environment,
			&_ComponentIds,
			&_LabelIds,
			&_DuedateAt,
			&_PlannedStartAt,
			&_PlannedEndAt,
			&_Key,
			&_CustomFieldIds,
			&_SprintID,
			&_EpicID,
			&_ParentID,
			&_StrategicParentID,
			&_InProgressCount,
			&_ReopenCount,
			&_InProgressDuration,
			&_VerificationDuration,
			&_InProgressStartAt,
			&_CustomerID,
			&_RefID,
			&_UserRefID,
			&_Cost,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraIssue{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _IssueTypeID.Valid {
			t.SetIssueTypeID(_IssueTypeID.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _AssigneeID.Valid {
			t.SetAssigneeID(_AssigneeID.String)
		}
		if _PriorityID.Valid {
			t.SetPriorityID(_PriorityID.String)
		}
		if _StatusID.Valid {
			t.SetStatusID(_StatusID.String)
		}
		if _ResolutionID.Valid {
			t.SetResolutionID(_ResolutionID.String)
		}
		if _FixVersionIds.Valid {
			t.SetFixVersionIds(_FixVersionIds.String)
		}
		if _VersionIds.Valid {
			t.SetVersionIds(_VersionIds.String)
		}
		if _Environment.Valid {
			t.SetEnvironment(_Environment.String)
		}
		if _ComponentIds.Valid {
			t.SetComponentIds(_ComponentIds.String)
		}
		if _LabelIds.Valid {
			t.SetLabelIds(_LabelIds.String)
		}
		if _DuedateAt.Valid {
			t.SetDuedateAt(_DuedateAt.Int64)
		}
		if _PlannedStartAt.Valid {
			t.SetPlannedStartAt(_PlannedStartAt.Int64)
		}
		if _PlannedEndAt.Valid {
			t.SetPlannedEndAt(_PlannedEndAt.Int64)
		}
		if _Key.Valid {
			t.SetKey(_Key.String)
		}
		if _CustomFieldIds.Valid {
			t.SetCustomFieldIds(_CustomFieldIds.String)
		}
		if _SprintID.Valid {
			t.SetSprintID(_SprintID.String)
		}
		if _EpicID.Valid {
			t.SetEpicID(_EpicID.String)
		}
		if _ParentID.Valid {
			t.SetParentID(_ParentID.String)
		}
		if _StrategicParentID.Valid {
			t.SetStrategicParentID(_StrategicParentID.String)
		}
		if _InProgressCount.Valid {
			t.SetInProgressCount(int32(_InProgressCount.Int64))
		}
		if _ReopenCount.Valid {
			t.SetReopenCount(int32(_ReopenCount.Int64))
		}
		if _InProgressDuration.Valid {
			t.SetInProgressDuration(_InProgressDuration.Int64)
		}
		if _VerificationDuration.Valid {
			t.SetVerificationDuration(_VerificationDuration.Int64)
		}
		if _InProgressStartAt.Valid {
			t.SetInProgressStartAt(_InProgressStartAt.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		if _UserRefID.Valid {
			t.SetUserRefID(_UserRefID.String)
		}
		if _Cost.Valid {
			t.SetCost(_Cost.Float64)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetInProgressCount will return the JiraIssue InProgressCount value
func (t *JiraIssue) GetInProgressCount() int32 {
	return t.InProgressCount
}

// SetInProgressCount will set the JiraIssue InProgressCount value
func (t *JiraIssue) SetInProgressCount(v int32) {
	t.InProgressCount = v
}

// GetReopenCount will return the JiraIssue ReopenCount value
func (t *JiraIssue) GetReopenCount() int32 {
	return t.ReopenCount
}

// SetReopenCount will set the JiraIssue ReopenCount value
func (t *JiraIssue) SetReopenCount(v int32) {
	t.ReopenCount = v
}

// GetInProgressDuration will return the JiraIssue InProgressDuration value
func (t *JiraIssue) GetInProgressDuration() int64 {
	return t.InProgressDuration
}

// SetInProgressDuration will set the JiraIssue InProgressDuration value
func (t *JiraIssue) SetInProgressDuration(v int64) {
	t.InProgressDuration = v
}

// GetVerificationDuration will return the JiraIssue VerificationDuration value
func (t *JiraIssue) GetVerificationDuration() int64 {
	return t.VerificationDuration
}

// SetVerificationDuration will set the JiraIssue VerificationDuration value
func (t *JiraIssue) SetVerificationDuration(v int64) {
	t.VerificationDuration = v
}

// GetInProgressStartAt will return the JiraIssue InProgressStartAt value
func (t *JiraIssue) GetInProgressStartAt() int64 {
	if t.InProgressStartAt == nil {
		return int64(0)
	}
	return *t.InProgressStartAt
}

// SetInProgressStartAt will set the JiraIssue InProgressStartAt value
func (t *JiraIssue) SetInProgressStartAt(v int64) {
	t.InProgressStartAt = &v
}

// GetCustomerID will return the JiraIssue CustomerID value
func (t *JiraIssue) GetCustomerID() string {
	return t.CustomerID
}

// SetCustomerID will set the JiraIssue CustomerID value
func (t *JiraIssue) SetCustomerID(v string) {
	t.CustomerID = v
}

// FindJiraIssuesByCustomerID will find all JiraIssues by the CustomerID value
func FindJiraIssuesByCustomerID(ctx context.Context, db DB, value string) ([]*JiraIssue, error) {
	q := "SELECT * FROM `jira_issue` WHERE `customer_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraIssue, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _IssueID sql.NullString
		var _ProjectID sql.NullString
		var _IssueTypeID sql.NullString
		var _UserID sql.NullString
		var _AssigneeID sql.NullString
		var _PriorityID sql.NullString
		var _StatusID sql.NullString
		var _ResolutionID sql.NullString
		var _FixVersionIds sql.NullString
		var _VersionIds sql.NullString
		var _Environment sql.NullString
		var _ComponentIds sql.NullString
		var _LabelIds sql.NullString
		var _DuedateAt sql.NullInt64
		var _PlannedStartAt sql.NullInt64
		var _PlannedEndAt sql.NullInt64
		var _Key sql.NullString
		var _CustomFieldIds sql.NullString
		var _SprintID sql.NullString
		var _EpicID sql.NullString
		var _ParentID sql.NullString
		var _StrategicParentID sql.NullString
		var _InProgressCount sql.NullInt64
		var _ReopenCount sql.NullInt64
		var _InProgressDuration sql.NullInt64
		var _VerificationDuration sql.NullInt64
		var _InProgressStartAt sql.NullInt64
		var _CustomerID sql.NullString
		var _RefID sql.NullString
		var _UserRefID sql.NullString
		var _Cost sql.NullFloat64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_IssueID,
			&_ProjectID,
			&_IssueTypeID,
			&_UserID,
			&_AssigneeID,
			&_PriorityID,
			&_StatusID,
			&_ResolutionID,
			&_FixVersionIds,
			&_VersionIds,
			&_Environment,
			&_ComponentIds,
			&_LabelIds,
			&_DuedateAt,
			&_PlannedStartAt,
			&_PlannedEndAt,
			&_Key,
			&_CustomFieldIds,
			&_SprintID,
			&_EpicID,
			&_ParentID,
			&_StrategicParentID,
			&_InProgressCount,
			&_ReopenCount,
			&_InProgressDuration,
			&_VerificationDuration,
			&_InProgressStartAt,
			&_CustomerID,
			&_RefID,
			&_UserRefID,
			&_Cost,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraIssue{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _IssueTypeID.Valid {
			t.SetIssueTypeID(_IssueTypeID.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _AssigneeID.Valid {
			t.SetAssigneeID(_AssigneeID.String)
		}
		if _PriorityID.Valid {
			t.SetPriorityID(_PriorityID.String)
		}
		if _StatusID.Valid {
			t.SetStatusID(_StatusID.String)
		}
		if _ResolutionID.Valid {
			t.SetResolutionID(_ResolutionID.String)
		}
		if _FixVersionIds.Valid {
			t.SetFixVersionIds(_FixVersionIds.String)
		}
		if _VersionIds.Valid {
			t.SetVersionIds(_VersionIds.String)
		}
		if _Environment.Valid {
			t.SetEnvironment(_Environment.String)
		}
		if _ComponentIds.Valid {
			t.SetComponentIds(_ComponentIds.String)
		}
		if _LabelIds.Valid {
			t.SetLabelIds(_LabelIds.String)
		}
		if _DuedateAt.Valid {
			t.SetDuedateAt(_DuedateAt.Int64)
		}
		if _PlannedStartAt.Valid {
			t.SetPlannedStartAt(_PlannedStartAt.Int64)
		}
		if _PlannedEndAt.Valid {
			t.SetPlannedEndAt(_PlannedEndAt.Int64)
		}
		if _Key.Valid {
			t.SetKey(_Key.String)
		}
		if _CustomFieldIds.Valid {
			t.SetCustomFieldIds(_CustomFieldIds.String)
		}
		if _SprintID.Valid {
			t.SetSprintID(_SprintID.String)
		}
		if _EpicID.Valid {
			t.SetEpicID(_EpicID.String)
		}
		if _ParentID.Valid {
			t.SetParentID(_ParentID.String)
		}
		if _StrategicParentID.Valid {
			t.SetStrategicParentID(_StrategicParentID.String)
		}
		if _InProgressCount.Valid {
			t.SetInProgressCount(int32(_InProgressCount.Int64))
		}
		if _ReopenCount.Valid {
			t.SetReopenCount(int32(_ReopenCount.Int64))
		}
		if _InProgressDuration.Valid {
			t.SetInProgressDuration(_InProgressDuration.Int64)
		}
		if _VerificationDuration.Valid {
			t.SetVerificationDuration(_VerificationDuration.Int64)
		}
		if _InProgressStartAt.Valid {
			t.SetInProgressStartAt(_InProgressStartAt.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		if _UserRefID.Valid {
			t.SetUserRefID(_UserRefID.String)
		}
		if _Cost.Valid {
			t.SetCost(_Cost.Float64)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindJiraIssuesByCustomerIDTx will find all JiraIssues by the CustomerID value using the provided transaction
func FindJiraIssuesByCustomerIDTx(ctx context.Context, tx Tx, value string) ([]*JiraIssue, error) {
	q := "SELECT * FROM `jira_issue` WHERE `customer_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraIssue, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _IssueID sql.NullString
		var _ProjectID sql.NullString
		var _IssueTypeID sql.NullString
		var _UserID sql.NullString
		var _AssigneeID sql.NullString
		var _PriorityID sql.NullString
		var _StatusID sql.NullString
		var _ResolutionID sql.NullString
		var _FixVersionIds sql.NullString
		var _VersionIds sql.NullString
		var _Environment sql.NullString
		var _ComponentIds sql.NullString
		var _LabelIds sql.NullString
		var _DuedateAt sql.NullInt64
		var _PlannedStartAt sql.NullInt64
		var _PlannedEndAt sql.NullInt64
		var _Key sql.NullString
		var _CustomFieldIds sql.NullString
		var _SprintID sql.NullString
		var _EpicID sql.NullString
		var _ParentID sql.NullString
		var _StrategicParentID sql.NullString
		var _InProgressCount sql.NullInt64
		var _ReopenCount sql.NullInt64
		var _InProgressDuration sql.NullInt64
		var _VerificationDuration sql.NullInt64
		var _InProgressStartAt sql.NullInt64
		var _CustomerID sql.NullString
		var _RefID sql.NullString
		var _UserRefID sql.NullString
		var _Cost sql.NullFloat64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_IssueID,
			&_ProjectID,
			&_IssueTypeID,
			&_UserID,
			&_AssigneeID,
			&_PriorityID,
			&_StatusID,
			&_ResolutionID,
			&_FixVersionIds,
			&_VersionIds,
			&_Environment,
			&_ComponentIds,
			&_LabelIds,
			&_DuedateAt,
			&_PlannedStartAt,
			&_PlannedEndAt,
			&_Key,
			&_CustomFieldIds,
			&_SprintID,
			&_EpicID,
			&_ParentID,
			&_StrategicParentID,
			&_InProgressCount,
			&_ReopenCount,
			&_InProgressDuration,
			&_VerificationDuration,
			&_InProgressStartAt,
			&_CustomerID,
			&_RefID,
			&_UserRefID,
			&_Cost,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraIssue{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _IssueTypeID.Valid {
			t.SetIssueTypeID(_IssueTypeID.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _AssigneeID.Valid {
			t.SetAssigneeID(_AssigneeID.String)
		}
		if _PriorityID.Valid {
			t.SetPriorityID(_PriorityID.String)
		}
		if _StatusID.Valid {
			t.SetStatusID(_StatusID.String)
		}
		if _ResolutionID.Valid {
			t.SetResolutionID(_ResolutionID.String)
		}
		if _FixVersionIds.Valid {
			t.SetFixVersionIds(_FixVersionIds.String)
		}
		if _VersionIds.Valid {
			t.SetVersionIds(_VersionIds.String)
		}
		if _Environment.Valid {
			t.SetEnvironment(_Environment.String)
		}
		if _ComponentIds.Valid {
			t.SetComponentIds(_ComponentIds.String)
		}
		if _LabelIds.Valid {
			t.SetLabelIds(_LabelIds.String)
		}
		if _DuedateAt.Valid {
			t.SetDuedateAt(_DuedateAt.Int64)
		}
		if _PlannedStartAt.Valid {
			t.SetPlannedStartAt(_PlannedStartAt.Int64)
		}
		if _PlannedEndAt.Valid {
			t.SetPlannedEndAt(_PlannedEndAt.Int64)
		}
		if _Key.Valid {
			t.SetKey(_Key.String)
		}
		if _CustomFieldIds.Valid {
			t.SetCustomFieldIds(_CustomFieldIds.String)
		}
		if _SprintID.Valid {
			t.SetSprintID(_SprintID.String)
		}
		if _EpicID.Valid {
			t.SetEpicID(_EpicID.String)
		}
		if _ParentID.Valid {
			t.SetParentID(_ParentID.String)
		}
		if _StrategicParentID.Valid {
			t.SetStrategicParentID(_StrategicParentID.String)
		}
		if _InProgressCount.Valid {
			t.SetInProgressCount(int32(_InProgressCount.Int64))
		}
		if _ReopenCount.Valid {
			t.SetReopenCount(int32(_ReopenCount.Int64))
		}
		if _InProgressDuration.Valid {
			t.SetInProgressDuration(_InProgressDuration.Int64)
		}
		if _VerificationDuration.Valid {
			t.SetVerificationDuration(_VerificationDuration.Int64)
		}
		if _InProgressStartAt.Valid {
			t.SetInProgressStartAt(_InProgressStartAt.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		if _UserRefID.Valid {
			t.SetUserRefID(_UserRefID.String)
		}
		if _Cost.Valid {
			t.SetCost(_Cost.Float64)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetRefID will return the JiraIssue RefID value
func (t *JiraIssue) GetRefID() string {
	return t.RefID
}

// SetRefID will set the JiraIssue RefID value
func (t *JiraIssue) SetRefID(v string) {
	t.RefID = v
}

// FindJiraIssuesByRefID will find all JiraIssues by the RefID value
func FindJiraIssuesByRefID(ctx context.Context, db DB, value string) ([]*JiraIssue, error) {
	q := "SELECT * FROM `jira_issue` WHERE `ref_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraIssue, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _IssueID sql.NullString
		var _ProjectID sql.NullString
		var _IssueTypeID sql.NullString
		var _UserID sql.NullString
		var _AssigneeID sql.NullString
		var _PriorityID sql.NullString
		var _StatusID sql.NullString
		var _ResolutionID sql.NullString
		var _FixVersionIds sql.NullString
		var _VersionIds sql.NullString
		var _Environment sql.NullString
		var _ComponentIds sql.NullString
		var _LabelIds sql.NullString
		var _DuedateAt sql.NullInt64
		var _PlannedStartAt sql.NullInt64
		var _PlannedEndAt sql.NullInt64
		var _Key sql.NullString
		var _CustomFieldIds sql.NullString
		var _SprintID sql.NullString
		var _EpicID sql.NullString
		var _ParentID sql.NullString
		var _StrategicParentID sql.NullString
		var _InProgressCount sql.NullInt64
		var _ReopenCount sql.NullInt64
		var _InProgressDuration sql.NullInt64
		var _VerificationDuration sql.NullInt64
		var _InProgressStartAt sql.NullInt64
		var _CustomerID sql.NullString
		var _RefID sql.NullString
		var _UserRefID sql.NullString
		var _Cost sql.NullFloat64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_IssueID,
			&_ProjectID,
			&_IssueTypeID,
			&_UserID,
			&_AssigneeID,
			&_PriorityID,
			&_StatusID,
			&_ResolutionID,
			&_FixVersionIds,
			&_VersionIds,
			&_Environment,
			&_ComponentIds,
			&_LabelIds,
			&_DuedateAt,
			&_PlannedStartAt,
			&_PlannedEndAt,
			&_Key,
			&_CustomFieldIds,
			&_SprintID,
			&_EpicID,
			&_ParentID,
			&_StrategicParentID,
			&_InProgressCount,
			&_ReopenCount,
			&_InProgressDuration,
			&_VerificationDuration,
			&_InProgressStartAt,
			&_CustomerID,
			&_RefID,
			&_UserRefID,
			&_Cost,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraIssue{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _IssueTypeID.Valid {
			t.SetIssueTypeID(_IssueTypeID.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _AssigneeID.Valid {
			t.SetAssigneeID(_AssigneeID.String)
		}
		if _PriorityID.Valid {
			t.SetPriorityID(_PriorityID.String)
		}
		if _StatusID.Valid {
			t.SetStatusID(_StatusID.String)
		}
		if _ResolutionID.Valid {
			t.SetResolutionID(_ResolutionID.String)
		}
		if _FixVersionIds.Valid {
			t.SetFixVersionIds(_FixVersionIds.String)
		}
		if _VersionIds.Valid {
			t.SetVersionIds(_VersionIds.String)
		}
		if _Environment.Valid {
			t.SetEnvironment(_Environment.String)
		}
		if _ComponentIds.Valid {
			t.SetComponentIds(_ComponentIds.String)
		}
		if _LabelIds.Valid {
			t.SetLabelIds(_LabelIds.String)
		}
		if _DuedateAt.Valid {
			t.SetDuedateAt(_DuedateAt.Int64)
		}
		if _PlannedStartAt.Valid {
			t.SetPlannedStartAt(_PlannedStartAt.Int64)
		}
		if _PlannedEndAt.Valid {
			t.SetPlannedEndAt(_PlannedEndAt.Int64)
		}
		if _Key.Valid {
			t.SetKey(_Key.String)
		}
		if _CustomFieldIds.Valid {
			t.SetCustomFieldIds(_CustomFieldIds.String)
		}
		if _SprintID.Valid {
			t.SetSprintID(_SprintID.String)
		}
		if _EpicID.Valid {
			t.SetEpicID(_EpicID.String)
		}
		if _ParentID.Valid {
			t.SetParentID(_ParentID.String)
		}
		if _StrategicParentID.Valid {
			t.SetStrategicParentID(_StrategicParentID.String)
		}
		if _InProgressCount.Valid {
			t.SetInProgressCount(int32(_InProgressCount.Int64))
		}
		if _ReopenCount.Valid {
			t.SetReopenCount(int32(_ReopenCount.Int64))
		}
		if _InProgressDuration.Valid {
			t.SetInProgressDuration(_InProgressDuration.Int64)
		}
		if _VerificationDuration.Valid {
			t.SetVerificationDuration(_VerificationDuration.Int64)
		}
		if _InProgressStartAt.Valid {
			t.SetInProgressStartAt(_InProgressStartAt.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		if _UserRefID.Valid {
			t.SetUserRefID(_UserRefID.String)
		}
		if _Cost.Valid {
			t.SetCost(_Cost.Float64)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindJiraIssuesByRefIDTx will find all JiraIssues by the RefID value using the provided transaction
func FindJiraIssuesByRefIDTx(ctx context.Context, tx Tx, value string) ([]*JiraIssue, error) {
	q := "SELECT * FROM `jira_issue` WHERE `ref_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraIssue, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _IssueID sql.NullString
		var _ProjectID sql.NullString
		var _IssueTypeID sql.NullString
		var _UserID sql.NullString
		var _AssigneeID sql.NullString
		var _PriorityID sql.NullString
		var _StatusID sql.NullString
		var _ResolutionID sql.NullString
		var _FixVersionIds sql.NullString
		var _VersionIds sql.NullString
		var _Environment sql.NullString
		var _ComponentIds sql.NullString
		var _LabelIds sql.NullString
		var _DuedateAt sql.NullInt64
		var _PlannedStartAt sql.NullInt64
		var _PlannedEndAt sql.NullInt64
		var _Key sql.NullString
		var _CustomFieldIds sql.NullString
		var _SprintID sql.NullString
		var _EpicID sql.NullString
		var _ParentID sql.NullString
		var _StrategicParentID sql.NullString
		var _InProgressCount sql.NullInt64
		var _ReopenCount sql.NullInt64
		var _InProgressDuration sql.NullInt64
		var _VerificationDuration sql.NullInt64
		var _InProgressStartAt sql.NullInt64
		var _CustomerID sql.NullString
		var _RefID sql.NullString
		var _UserRefID sql.NullString
		var _Cost sql.NullFloat64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_IssueID,
			&_ProjectID,
			&_IssueTypeID,
			&_UserID,
			&_AssigneeID,
			&_PriorityID,
			&_StatusID,
			&_ResolutionID,
			&_FixVersionIds,
			&_VersionIds,
			&_Environment,
			&_ComponentIds,
			&_LabelIds,
			&_DuedateAt,
			&_PlannedStartAt,
			&_PlannedEndAt,
			&_Key,
			&_CustomFieldIds,
			&_SprintID,
			&_EpicID,
			&_ParentID,
			&_StrategicParentID,
			&_InProgressCount,
			&_ReopenCount,
			&_InProgressDuration,
			&_VerificationDuration,
			&_InProgressStartAt,
			&_CustomerID,
			&_RefID,
			&_UserRefID,
			&_Cost,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraIssue{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _IssueTypeID.Valid {
			t.SetIssueTypeID(_IssueTypeID.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _AssigneeID.Valid {
			t.SetAssigneeID(_AssigneeID.String)
		}
		if _PriorityID.Valid {
			t.SetPriorityID(_PriorityID.String)
		}
		if _StatusID.Valid {
			t.SetStatusID(_StatusID.String)
		}
		if _ResolutionID.Valid {
			t.SetResolutionID(_ResolutionID.String)
		}
		if _FixVersionIds.Valid {
			t.SetFixVersionIds(_FixVersionIds.String)
		}
		if _VersionIds.Valid {
			t.SetVersionIds(_VersionIds.String)
		}
		if _Environment.Valid {
			t.SetEnvironment(_Environment.String)
		}
		if _ComponentIds.Valid {
			t.SetComponentIds(_ComponentIds.String)
		}
		if _LabelIds.Valid {
			t.SetLabelIds(_LabelIds.String)
		}
		if _DuedateAt.Valid {
			t.SetDuedateAt(_DuedateAt.Int64)
		}
		if _PlannedStartAt.Valid {
			t.SetPlannedStartAt(_PlannedStartAt.Int64)
		}
		if _PlannedEndAt.Valid {
			t.SetPlannedEndAt(_PlannedEndAt.Int64)
		}
		if _Key.Valid {
			t.SetKey(_Key.String)
		}
		if _CustomFieldIds.Valid {
			t.SetCustomFieldIds(_CustomFieldIds.String)
		}
		if _SprintID.Valid {
			t.SetSprintID(_SprintID.String)
		}
		if _EpicID.Valid {
			t.SetEpicID(_EpicID.String)
		}
		if _ParentID.Valid {
			t.SetParentID(_ParentID.String)
		}
		if _StrategicParentID.Valid {
			t.SetStrategicParentID(_StrategicParentID.String)
		}
		if _InProgressCount.Valid {
			t.SetInProgressCount(int32(_InProgressCount.Int64))
		}
		if _ReopenCount.Valid {
			t.SetReopenCount(int32(_ReopenCount.Int64))
		}
		if _InProgressDuration.Valid {
			t.SetInProgressDuration(_InProgressDuration.Int64)
		}
		if _VerificationDuration.Valid {
			t.SetVerificationDuration(_VerificationDuration.Int64)
		}
		if _InProgressStartAt.Valid {
			t.SetInProgressStartAt(_InProgressStartAt.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		if _UserRefID.Valid {
			t.SetUserRefID(_UserRefID.String)
		}
		if _Cost.Valid {
			t.SetCost(_Cost.Float64)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetUserRefID will return the JiraIssue UserRefID value
func (t *JiraIssue) GetUserRefID() string {
	return t.UserRefID
}

// SetUserRefID will set the JiraIssue UserRefID value
func (t *JiraIssue) SetUserRefID(v string) {
	t.UserRefID = v
}

// FindJiraIssuesByUserRefID will find all JiraIssues by the UserRefID value
func FindJiraIssuesByUserRefID(ctx context.Context, db DB, value string) ([]*JiraIssue, error) {
	q := "SELECT * FROM `jira_issue` WHERE `user_ref_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraIssue, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _IssueID sql.NullString
		var _ProjectID sql.NullString
		var _IssueTypeID sql.NullString
		var _UserID sql.NullString
		var _AssigneeID sql.NullString
		var _PriorityID sql.NullString
		var _StatusID sql.NullString
		var _ResolutionID sql.NullString
		var _FixVersionIds sql.NullString
		var _VersionIds sql.NullString
		var _Environment sql.NullString
		var _ComponentIds sql.NullString
		var _LabelIds sql.NullString
		var _DuedateAt sql.NullInt64
		var _PlannedStartAt sql.NullInt64
		var _PlannedEndAt sql.NullInt64
		var _Key sql.NullString
		var _CustomFieldIds sql.NullString
		var _SprintID sql.NullString
		var _EpicID sql.NullString
		var _ParentID sql.NullString
		var _StrategicParentID sql.NullString
		var _InProgressCount sql.NullInt64
		var _ReopenCount sql.NullInt64
		var _InProgressDuration sql.NullInt64
		var _VerificationDuration sql.NullInt64
		var _InProgressStartAt sql.NullInt64
		var _CustomerID sql.NullString
		var _RefID sql.NullString
		var _UserRefID sql.NullString
		var _Cost sql.NullFloat64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_IssueID,
			&_ProjectID,
			&_IssueTypeID,
			&_UserID,
			&_AssigneeID,
			&_PriorityID,
			&_StatusID,
			&_ResolutionID,
			&_FixVersionIds,
			&_VersionIds,
			&_Environment,
			&_ComponentIds,
			&_LabelIds,
			&_DuedateAt,
			&_PlannedStartAt,
			&_PlannedEndAt,
			&_Key,
			&_CustomFieldIds,
			&_SprintID,
			&_EpicID,
			&_ParentID,
			&_StrategicParentID,
			&_InProgressCount,
			&_ReopenCount,
			&_InProgressDuration,
			&_VerificationDuration,
			&_InProgressStartAt,
			&_CustomerID,
			&_RefID,
			&_UserRefID,
			&_Cost,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraIssue{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _IssueTypeID.Valid {
			t.SetIssueTypeID(_IssueTypeID.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _AssigneeID.Valid {
			t.SetAssigneeID(_AssigneeID.String)
		}
		if _PriorityID.Valid {
			t.SetPriorityID(_PriorityID.String)
		}
		if _StatusID.Valid {
			t.SetStatusID(_StatusID.String)
		}
		if _ResolutionID.Valid {
			t.SetResolutionID(_ResolutionID.String)
		}
		if _FixVersionIds.Valid {
			t.SetFixVersionIds(_FixVersionIds.String)
		}
		if _VersionIds.Valid {
			t.SetVersionIds(_VersionIds.String)
		}
		if _Environment.Valid {
			t.SetEnvironment(_Environment.String)
		}
		if _ComponentIds.Valid {
			t.SetComponentIds(_ComponentIds.String)
		}
		if _LabelIds.Valid {
			t.SetLabelIds(_LabelIds.String)
		}
		if _DuedateAt.Valid {
			t.SetDuedateAt(_DuedateAt.Int64)
		}
		if _PlannedStartAt.Valid {
			t.SetPlannedStartAt(_PlannedStartAt.Int64)
		}
		if _PlannedEndAt.Valid {
			t.SetPlannedEndAt(_PlannedEndAt.Int64)
		}
		if _Key.Valid {
			t.SetKey(_Key.String)
		}
		if _CustomFieldIds.Valid {
			t.SetCustomFieldIds(_CustomFieldIds.String)
		}
		if _SprintID.Valid {
			t.SetSprintID(_SprintID.String)
		}
		if _EpicID.Valid {
			t.SetEpicID(_EpicID.String)
		}
		if _ParentID.Valid {
			t.SetParentID(_ParentID.String)
		}
		if _StrategicParentID.Valid {
			t.SetStrategicParentID(_StrategicParentID.String)
		}
		if _InProgressCount.Valid {
			t.SetInProgressCount(int32(_InProgressCount.Int64))
		}
		if _ReopenCount.Valid {
			t.SetReopenCount(int32(_ReopenCount.Int64))
		}
		if _InProgressDuration.Valid {
			t.SetInProgressDuration(_InProgressDuration.Int64)
		}
		if _VerificationDuration.Valid {
			t.SetVerificationDuration(_VerificationDuration.Int64)
		}
		if _InProgressStartAt.Valid {
			t.SetInProgressStartAt(_InProgressStartAt.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		if _UserRefID.Valid {
			t.SetUserRefID(_UserRefID.String)
		}
		if _Cost.Valid {
			t.SetCost(_Cost.Float64)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindJiraIssuesByUserRefIDTx will find all JiraIssues by the UserRefID value using the provided transaction
func FindJiraIssuesByUserRefIDTx(ctx context.Context, tx Tx, value string) ([]*JiraIssue, error) {
	q := "SELECT * FROM `jira_issue` WHERE `user_ref_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraIssue, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _IssueID sql.NullString
		var _ProjectID sql.NullString
		var _IssueTypeID sql.NullString
		var _UserID sql.NullString
		var _AssigneeID sql.NullString
		var _PriorityID sql.NullString
		var _StatusID sql.NullString
		var _ResolutionID sql.NullString
		var _FixVersionIds sql.NullString
		var _VersionIds sql.NullString
		var _Environment sql.NullString
		var _ComponentIds sql.NullString
		var _LabelIds sql.NullString
		var _DuedateAt sql.NullInt64
		var _PlannedStartAt sql.NullInt64
		var _PlannedEndAt sql.NullInt64
		var _Key sql.NullString
		var _CustomFieldIds sql.NullString
		var _SprintID sql.NullString
		var _EpicID sql.NullString
		var _ParentID sql.NullString
		var _StrategicParentID sql.NullString
		var _InProgressCount sql.NullInt64
		var _ReopenCount sql.NullInt64
		var _InProgressDuration sql.NullInt64
		var _VerificationDuration sql.NullInt64
		var _InProgressStartAt sql.NullInt64
		var _CustomerID sql.NullString
		var _RefID sql.NullString
		var _UserRefID sql.NullString
		var _Cost sql.NullFloat64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_IssueID,
			&_ProjectID,
			&_IssueTypeID,
			&_UserID,
			&_AssigneeID,
			&_PriorityID,
			&_StatusID,
			&_ResolutionID,
			&_FixVersionIds,
			&_VersionIds,
			&_Environment,
			&_ComponentIds,
			&_LabelIds,
			&_DuedateAt,
			&_PlannedStartAt,
			&_PlannedEndAt,
			&_Key,
			&_CustomFieldIds,
			&_SprintID,
			&_EpicID,
			&_ParentID,
			&_StrategicParentID,
			&_InProgressCount,
			&_ReopenCount,
			&_InProgressDuration,
			&_VerificationDuration,
			&_InProgressStartAt,
			&_CustomerID,
			&_RefID,
			&_UserRefID,
			&_Cost,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraIssue{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _IssueTypeID.Valid {
			t.SetIssueTypeID(_IssueTypeID.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _AssigneeID.Valid {
			t.SetAssigneeID(_AssigneeID.String)
		}
		if _PriorityID.Valid {
			t.SetPriorityID(_PriorityID.String)
		}
		if _StatusID.Valid {
			t.SetStatusID(_StatusID.String)
		}
		if _ResolutionID.Valid {
			t.SetResolutionID(_ResolutionID.String)
		}
		if _FixVersionIds.Valid {
			t.SetFixVersionIds(_FixVersionIds.String)
		}
		if _VersionIds.Valid {
			t.SetVersionIds(_VersionIds.String)
		}
		if _Environment.Valid {
			t.SetEnvironment(_Environment.String)
		}
		if _ComponentIds.Valid {
			t.SetComponentIds(_ComponentIds.String)
		}
		if _LabelIds.Valid {
			t.SetLabelIds(_LabelIds.String)
		}
		if _DuedateAt.Valid {
			t.SetDuedateAt(_DuedateAt.Int64)
		}
		if _PlannedStartAt.Valid {
			t.SetPlannedStartAt(_PlannedStartAt.Int64)
		}
		if _PlannedEndAt.Valid {
			t.SetPlannedEndAt(_PlannedEndAt.Int64)
		}
		if _Key.Valid {
			t.SetKey(_Key.String)
		}
		if _CustomFieldIds.Valid {
			t.SetCustomFieldIds(_CustomFieldIds.String)
		}
		if _SprintID.Valid {
			t.SetSprintID(_SprintID.String)
		}
		if _EpicID.Valid {
			t.SetEpicID(_EpicID.String)
		}
		if _ParentID.Valid {
			t.SetParentID(_ParentID.String)
		}
		if _StrategicParentID.Valid {
			t.SetStrategicParentID(_StrategicParentID.String)
		}
		if _InProgressCount.Valid {
			t.SetInProgressCount(int32(_InProgressCount.Int64))
		}
		if _ReopenCount.Valid {
			t.SetReopenCount(int32(_ReopenCount.Int64))
		}
		if _InProgressDuration.Valid {
			t.SetInProgressDuration(_InProgressDuration.Int64)
		}
		if _VerificationDuration.Valid {
			t.SetVerificationDuration(_VerificationDuration.Int64)
		}
		if _InProgressStartAt.Valid {
			t.SetInProgressStartAt(_InProgressStartAt.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		if _UserRefID.Valid {
			t.SetUserRefID(_UserRefID.String)
		}
		if _Cost.Valid {
			t.SetCost(_Cost.Float64)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetCost will return the JiraIssue Cost value
func (t *JiraIssue) GetCost() float64 {
	return t.Cost
}

// SetCost will set the JiraIssue Cost value
func (t *JiraIssue) SetCost(v float64) {
	t.Cost = v
}

func (t *JiraIssue) toTimestamp(value time.Time) *timestamp.Timestamp {
	ts, _ := ptypes.TimestampProto(value)
	return ts
}

// DBCreateJiraIssueTable will create the JiraIssue table
func DBCreateJiraIssueTable(ctx context.Context, db DB) error {
	q := "CREATE TABLE `jira_issue` (`id`VARCHAR(64) NOT NULL PRIMARY KEY,`checksum`CHAR(64),`issue_id`VARCHAR(255) NOT NULL,`project_id` VARCHAR(64) NOT NULL,`issue_type_id` VARCHAR(64) NOT NULL,`user_id` VARCHAR(64),`assignee_id`VARCHAR(64),`priority_id`VARCHAR(64),`status_id` VARCHAR(64) NOT NULL,`resolution_id` VARCHAR(64),`fix_version_ids` JSON,`version_ids`JSON,`environment`TEXT,`component_ids` JSON,`label_ids` JSON,`duedate_at` BIGINT UNSIGNED,`planned_start_at` BIGINT UNSIGNED,`planned_end_at`BIGINT UNSIGNED,`key` VARCHAR(20) NOT NULL,`custom_field_ids` JSON,`sprint_id` JSON,`epic_id` VARCHAR(64),`parent_id` VARCHAR(64),`strategic_parent_id` VARCHAR(64),`in_progress_count`INT NOT NULL,`reopen_count` INT NOT NULL,`in_progress_duration`BIGINT NOT NULL,`verification_duration` BIGINT NOT NULL,`in_progress_start_at`BIGINT,`customer_id`VARCHAR(64) NOT NULL,`ref_id` VARCHAR(64) NOT NULL,`user_ref_id`VARCHAR(64) NOT NULL,`cost` REAL NOT NULL DEFAULT 0,INDEX jira_issue_issue_id_index (`issue_id`),INDEX jira_issue_project_id_index (`project_id`),INDEX jira_issue_issue_type_id_index (`issue_type_id`),INDEX jira_issue_key_index (`key`),INDEX jira_issue_epic_id_index (`epic_id`),INDEX jira_issue_parent_id_index (`parent_id`),INDEX jira_issue_strategic_parent_id_index (`strategic_parent_id`),INDEX jira_issue_customer_id_index (`customer_id`),INDEX jira_issue_ref_id_index (`ref_id`),INDEX jira_issue_user_ref_id_index (`user_ref_id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBCreateJiraIssueTableTx will create the JiraIssue table using the provided transction
func DBCreateJiraIssueTableTx(ctx context.Context, tx Tx) error {
	q := "CREATE TABLE `jira_issue` (`id`VARCHAR(64) NOT NULL PRIMARY KEY,`checksum`CHAR(64),`issue_id`VARCHAR(255) NOT NULL,`project_id` VARCHAR(64) NOT NULL,`issue_type_id` VARCHAR(64) NOT NULL,`user_id` VARCHAR(64),`assignee_id`VARCHAR(64),`priority_id`VARCHAR(64),`status_id` VARCHAR(64) NOT NULL,`resolution_id` VARCHAR(64),`fix_version_ids` JSON,`version_ids`JSON,`environment`TEXT,`component_ids` JSON,`label_ids` JSON,`duedate_at` BIGINT UNSIGNED,`planned_start_at` BIGINT UNSIGNED,`planned_end_at`BIGINT UNSIGNED,`key` VARCHAR(20) NOT NULL,`custom_field_ids` JSON,`sprint_id` JSON,`epic_id` VARCHAR(64),`parent_id` VARCHAR(64),`strategic_parent_id` VARCHAR(64),`in_progress_count`INT NOT NULL,`reopen_count` INT NOT NULL,`in_progress_duration`BIGINT NOT NULL,`verification_duration` BIGINT NOT NULL,`in_progress_start_at`BIGINT,`customer_id`VARCHAR(64) NOT NULL,`ref_id` VARCHAR(64) NOT NULL,`user_ref_id`VARCHAR(64) NOT NULL,`cost` REAL NOT NULL DEFAULT 0,INDEX jira_issue_issue_id_index (`issue_id`),INDEX jira_issue_project_id_index (`project_id`),INDEX jira_issue_issue_type_id_index (`issue_type_id`),INDEX jira_issue_key_index (`key`),INDEX jira_issue_epic_id_index (`epic_id`),INDEX jira_issue_parent_id_index (`parent_id`),INDEX jira_issue_strategic_parent_id_index (`strategic_parent_id`),INDEX jira_issue_customer_id_index (`customer_id`),INDEX jira_issue_ref_id_index (`ref_id`),INDEX jira_issue_user_ref_id_index (`user_ref_id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// DBDropJiraIssueTable will drop the JiraIssue table
func DBDropJiraIssueTable(ctx context.Context, db DB) error {
	q := "DROP TABLE IF EXISTS `jira_issue`"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBDropJiraIssueTableTx will drop the JiraIssue table using the provided transaction
func DBDropJiraIssueTableTx(ctx context.Context, tx Tx) error {
	q := "DROP TABLE IF EXISTS `jira_issue`"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// CalculateChecksum will calculate a checksum of the SHA1 of all field values
func (t *JiraIssue) CalculateChecksum() string {
	return orm.HashStrings(
		orm.ToString(t.ID),
		orm.ToString(t.IssueID),
		orm.ToString(t.ProjectID),
		orm.ToString(t.IssueTypeID),
		orm.ToString(t.UserID),
		orm.ToString(t.AssigneeID),
		orm.ToString(t.PriorityID),
		orm.ToString(t.StatusID),
		orm.ToString(t.ResolutionID),
		orm.ToString(t.FixVersionIds),
		orm.ToString(t.VersionIds),
		orm.ToString(t.Environment),
		orm.ToString(t.ComponentIds),
		orm.ToString(t.LabelIds),
		orm.ToString(t.DuedateAt),
		orm.ToString(t.PlannedStartAt),
		orm.ToString(t.PlannedEndAt),
		orm.ToString(t.Key),
		orm.ToString(t.CustomFieldIds),
		orm.ToString(t.SprintID),
		orm.ToString(t.EpicID),
		orm.ToString(t.ParentID),
		orm.ToString(t.StrategicParentID),
		orm.ToString(t.InProgressCount),
		orm.ToString(t.ReopenCount),
		orm.ToString(t.InProgressDuration),
		orm.ToString(t.VerificationDuration),
		orm.ToString(t.InProgressStartAt),
		orm.ToString(t.CustomerID),
		orm.ToString(t.RefID),
		orm.ToString(t.UserRefID),
		orm.ToString(t.Cost),
	)
}

// DBCreate will create a new JiraIssue record in the database
func (t *JiraIssue) DBCreate(ctx context.Context, db DB) (sql.Result, error) {
	q := "INSERT INTO `jira_issue` (`jira_issue`.`id`,`jira_issue`.`checksum`,`jira_issue`.`issue_id`,`jira_issue`.`project_id`,`jira_issue`.`issue_type_id`,`jira_issue`.`user_id`,`jira_issue`.`assignee_id`,`jira_issue`.`priority_id`,`jira_issue`.`status_id`,`jira_issue`.`resolution_id`,`jira_issue`.`fix_version_ids`,`jira_issue`.`version_ids`,`jira_issue`.`environment`,`jira_issue`.`component_ids`,`jira_issue`.`label_ids`,`jira_issue`.`duedate_at`,`jira_issue`.`planned_start_at`,`jira_issue`.`planned_end_at`,`jira_issue`.`key`,`jira_issue`.`custom_field_ids`,`jira_issue`.`sprint_id`,`jira_issue`.`epic_id`,`jira_issue`.`parent_id`,`jira_issue`.`strategic_parent_id`,`jira_issue`.`in_progress_count`,`jira_issue`.`reopen_count`,`jira_issue`.`in_progress_duration`,`jira_issue`.`verification_duration`,`jira_issue`.`in_progress_start_at`,`jira_issue`.`customer_id`,`jira_issue`.`ref_id`,`jira_issue`.`user_ref_id`,`jira_issue`.`cost`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.IssueID),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLString(t.IssueTypeID),
		orm.ToSQLString(t.UserID),
		orm.ToSQLString(t.AssigneeID),
		orm.ToSQLString(t.PriorityID),
		orm.ToSQLString(t.StatusID),
		orm.ToSQLString(t.ResolutionID),
		orm.ToSQLString(t.FixVersionIds),
		orm.ToSQLString(t.VersionIds),
		orm.ToSQLString(t.Environment),
		orm.ToSQLString(t.ComponentIds),
		orm.ToSQLString(t.LabelIds),
		orm.ToSQLInt64(t.DuedateAt),
		orm.ToSQLInt64(t.PlannedStartAt),
		orm.ToSQLInt64(t.PlannedEndAt),
		orm.ToSQLString(t.Key),
		orm.ToSQLString(t.CustomFieldIds),
		orm.ToSQLString(t.SprintID),
		orm.ToSQLString(t.EpicID),
		orm.ToSQLString(t.ParentID),
		orm.ToSQLString(t.StrategicParentID),
		orm.ToSQLInt64(t.InProgressCount),
		orm.ToSQLInt64(t.ReopenCount),
		orm.ToSQLInt64(t.InProgressDuration),
		orm.ToSQLInt64(t.VerificationDuration),
		orm.ToSQLInt64(t.InProgressStartAt),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.UserRefID),
		orm.ToSQLFloat64(t.Cost),
	)
}

// DBCreateTx will create a new JiraIssue record in the database using the provided transaction
func (t *JiraIssue) DBCreateTx(ctx context.Context, tx Tx) (sql.Result, error) {
	q := "INSERT INTO `jira_issue` (`jira_issue`.`id`,`jira_issue`.`checksum`,`jira_issue`.`issue_id`,`jira_issue`.`project_id`,`jira_issue`.`issue_type_id`,`jira_issue`.`user_id`,`jira_issue`.`assignee_id`,`jira_issue`.`priority_id`,`jira_issue`.`status_id`,`jira_issue`.`resolution_id`,`jira_issue`.`fix_version_ids`,`jira_issue`.`version_ids`,`jira_issue`.`environment`,`jira_issue`.`component_ids`,`jira_issue`.`label_ids`,`jira_issue`.`duedate_at`,`jira_issue`.`planned_start_at`,`jira_issue`.`planned_end_at`,`jira_issue`.`key`,`jira_issue`.`custom_field_ids`,`jira_issue`.`sprint_id`,`jira_issue`.`epic_id`,`jira_issue`.`parent_id`,`jira_issue`.`strategic_parent_id`,`jira_issue`.`in_progress_count`,`jira_issue`.`reopen_count`,`jira_issue`.`in_progress_duration`,`jira_issue`.`verification_duration`,`jira_issue`.`in_progress_start_at`,`jira_issue`.`customer_id`,`jira_issue`.`ref_id`,`jira_issue`.`user_ref_id`,`jira_issue`.`cost`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.IssueID),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLString(t.IssueTypeID),
		orm.ToSQLString(t.UserID),
		orm.ToSQLString(t.AssigneeID),
		orm.ToSQLString(t.PriorityID),
		orm.ToSQLString(t.StatusID),
		orm.ToSQLString(t.ResolutionID),
		orm.ToSQLString(t.FixVersionIds),
		orm.ToSQLString(t.VersionIds),
		orm.ToSQLString(t.Environment),
		orm.ToSQLString(t.ComponentIds),
		orm.ToSQLString(t.LabelIds),
		orm.ToSQLInt64(t.DuedateAt),
		orm.ToSQLInt64(t.PlannedStartAt),
		orm.ToSQLInt64(t.PlannedEndAt),
		orm.ToSQLString(t.Key),
		orm.ToSQLString(t.CustomFieldIds),
		orm.ToSQLString(t.SprintID),
		orm.ToSQLString(t.EpicID),
		orm.ToSQLString(t.ParentID),
		orm.ToSQLString(t.StrategicParentID),
		orm.ToSQLInt64(t.InProgressCount),
		orm.ToSQLInt64(t.ReopenCount),
		orm.ToSQLInt64(t.InProgressDuration),
		orm.ToSQLInt64(t.VerificationDuration),
		orm.ToSQLInt64(t.InProgressStartAt),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.UserRefID),
		orm.ToSQLFloat64(t.Cost),
	)
}

// DBCreateIgnoreDuplicate will upsert the JiraIssue record in the database
func (t *JiraIssue) DBCreateIgnoreDuplicate(ctx context.Context, db DB) (sql.Result, error) {
	q := "INSERT INTO `jira_issue` (`jira_issue`.`id`,`jira_issue`.`checksum`,`jira_issue`.`issue_id`,`jira_issue`.`project_id`,`jira_issue`.`issue_type_id`,`jira_issue`.`user_id`,`jira_issue`.`assignee_id`,`jira_issue`.`priority_id`,`jira_issue`.`status_id`,`jira_issue`.`resolution_id`,`jira_issue`.`fix_version_ids`,`jira_issue`.`version_ids`,`jira_issue`.`environment`,`jira_issue`.`component_ids`,`jira_issue`.`label_ids`,`jira_issue`.`duedate_at`,`jira_issue`.`planned_start_at`,`jira_issue`.`planned_end_at`,`jira_issue`.`key`,`jira_issue`.`custom_field_ids`,`jira_issue`.`sprint_id`,`jira_issue`.`epic_id`,`jira_issue`.`parent_id`,`jira_issue`.`strategic_parent_id`,`jira_issue`.`in_progress_count`,`jira_issue`.`reopen_count`,`jira_issue`.`in_progress_duration`,`jira_issue`.`verification_duration`,`jira_issue`.`in_progress_start_at`,`jira_issue`.`customer_id`,`jira_issue`.`ref_id`,`jira_issue`.`user_ref_id`,`jira_issue`.`cost`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.IssueID),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLString(t.IssueTypeID),
		orm.ToSQLString(t.UserID),
		orm.ToSQLString(t.AssigneeID),
		orm.ToSQLString(t.PriorityID),
		orm.ToSQLString(t.StatusID),
		orm.ToSQLString(t.ResolutionID),
		orm.ToSQLString(t.FixVersionIds),
		orm.ToSQLString(t.VersionIds),
		orm.ToSQLString(t.Environment),
		orm.ToSQLString(t.ComponentIds),
		orm.ToSQLString(t.LabelIds),
		orm.ToSQLInt64(t.DuedateAt),
		orm.ToSQLInt64(t.PlannedStartAt),
		orm.ToSQLInt64(t.PlannedEndAt),
		orm.ToSQLString(t.Key),
		orm.ToSQLString(t.CustomFieldIds),
		orm.ToSQLString(t.SprintID),
		orm.ToSQLString(t.EpicID),
		orm.ToSQLString(t.ParentID),
		orm.ToSQLString(t.StrategicParentID),
		orm.ToSQLInt64(t.InProgressCount),
		orm.ToSQLInt64(t.ReopenCount),
		orm.ToSQLInt64(t.InProgressDuration),
		orm.ToSQLInt64(t.VerificationDuration),
		orm.ToSQLInt64(t.InProgressStartAt),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.UserRefID),
		orm.ToSQLFloat64(t.Cost),
	)
}

// DBCreateIgnoreDuplicateTx will upsert the JiraIssue record in the database using the provided transaction
func (t *JiraIssue) DBCreateIgnoreDuplicateTx(ctx context.Context, tx Tx) (sql.Result, error) {
	q := "INSERT INTO `jira_issue` (`jira_issue`.`id`,`jira_issue`.`checksum`,`jira_issue`.`issue_id`,`jira_issue`.`project_id`,`jira_issue`.`issue_type_id`,`jira_issue`.`user_id`,`jira_issue`.`assignee_id`,`jira_issue`.`priority_id`,`jira_issue`.`status_id`,`jira_issue`.`resolution_id`,`jira_issue`.`fix_version_ids`,`jira_issue`.`version_ids`,`jira_issue`.`environment`,`jira_issue`.`component_ids`,`jira_issue`.`label_ids`,`jira_issue`.`duedate_at`,`jira_issue`.`planned_start_at`,`jira_issue`.`planned_end_at`,`jira_issue`.`key`,`jira_issue`.`custom_field_ids`,`jira_issue`.`sprint_id`,`jira_issue`.`epic_id`,`jira_issue`.`parent_id`,`jira_issue`.`strategic_parent_id`,`jira_issue`.`in_progress_count`,`jira_issue`.`reopen_count`,`jira_issue`.`in_progress_duration`,`jira_issue`.`verification_duration`,`jira_issue`.`in_progress_start_at`,`jira_issue`.`customer_id`,`jira_issue`.`ref_id`,`jira_issue`.`user_ref_id`,`jira_issue`.`cost`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.IssueID),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLString(t.IssueTypeID),
		orm.ToSQLString(t.UserID),
		orm.ToSQLString(t.AssigneeID),
		orm.ToSQLString(t.PriorityID),
		orm.ToSQLString(t.StatusID),
		orm.ToSQLString(t.ResolutionID),
		orm.ToSQLString(t.FixVersionIds),
		orm.ToSQLString(t.VersionIds),
		orm.ToSQLString(t.Environment),
		orm.ToSQLString(t.ComponentIds),
		orm.ToSQLString(t.LabelIds),
		orm.ToSQLInt64(t.DuedateAt),
		orm.ToSQLInt64(t.PlannedStartAt),
		orm.ToSQLInt64(t.PlannedEndAt),
		orm.ToSQLString(t.Key),
		orm.ToSQLString(t.CustomFieldIds),
		orm.ToSQLString(t.SprintID),
		orm.ToSQLString(t.EpicID),
		orm.ToSQLString(t.ParentID),
		orm.ToSQLString(t.StrategicParentID),
		orm.ToSQLInt64(t.InProgressCount),
		orm.ToSQLInt64(t.ReopenCount),
		orm.ToSQLInt64(t.InProgressDuration),
		orm.ToSQLInt64(t.VerificationDuration),
		orm.ToSQLInt64(t.InProgressStartAt),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.UserRefID),
		orm.ToSQLFloat64(t.Cost),
	)
}

// DeleteAllJiraIssues deletes all JiraIssue records in the database with optional filters
func DeleteAllJiraIssues(ctx context.Context, db DB, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(JiraIssueTableName),
	}
	if len(_params) > 0 {
		for _, param := range _params {
			params = append(params, param)
		}
	}
	q, p := orm.BuildQuery(params...)
	_, err := db.ExecContext(ctx, "DELETE "+q, p...)
	return err
}

// DeleteAllJiraIssuesTx deletes all JiraIssue records in the database with optional filters using the provided transaction
func DeleteAllJiraIssuesTx(ctx context.Context, tx Tx, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(JiraIssueTableName),
	}
	if len(_params) > 0 {
		for _, param := range _params {
			params = append(params, param)
		}
	}
	q, p := orm.BuildQuery(params...)
	_, err := tx.ExecContext(ctx, "DELETE "+q, p...)
	return err
}

// DBDelete will delete this JiraIssue record in the database
func (t *JiraIssue) DBDelete(ctx context.Context, db DB) (bool, error) {
	q := "DELETE FROM `jira_issue` WHERE `id` = ?"
	r, err := db.ExecContext(ctx, q, orm.ToSQLString(t.ID))
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	if err == sql.ErrNoRows {
		return false, nil
	}
	c, _ := r.RowsAffected()
	return c > 0, nil
}

// DBDeleteTx will delete this JiraIssue record in the database using the provided transaction
func (t *JiraIssue) DBDeleteTx(ctx context.Context, tx Tx) (bool, error) {
	q := "DELETE FROM `jira_issue` WHERE `id` = ?"
	r, err := tx.ExecContext(ctx, q, orm.ToSQLString(t.ID))
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	if err == sql.ErrNoRows {
		return false, nil
	}
	c, _ := r.RowsAffected()
	return c > 0, nil
}

// DBUpdate will update the JiraIssue record in the database
func (t *JiraIssue) DBUpdate(ctx context.Context, db DB) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `jira_issue` SET `checksum`=?,`issue_id`=?,`project_id`=?,`issue_type_id`=?,`user_id`=?,`assignee_id`=?,`priority_id`=?,`status_id`=?,`resolution_id`=?,`fix_version_ids`=?,`version_ids`=?,`environment`=?,`component_ids`=?,`label_ids`=?,`duedate_at`=?,`planned_start_at`=?,`planned_end_at`=?,`key`=?,`custom_field_ids`=?,`sprint_id`=?,`epic_id`=?,`parent_id`=?,`strategic_parent_id`=?,`in_progress_count`=?,`reopen_count`=?,`in_progress_duration`=?,`verification_duration`=?,`in_progress_start_at`=?,`customer_id`=?,`ref_id`=?,`user_ref_id`=?,`cost`=? WHERE `id`=?"
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.IssueID),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLString(t.IssueTypeID),
		orm.ToSQLString(t.UserID),
		orm.ToSQLString(t.AssigneeID),
		orm.ToSQLString(t.PriorityID),
		orm.ToSQLString(t.StatusID),
		orm.ToSQLString(t.ResolutionID),
		orm.ToSQLString(t.FixVersionIds),
		orm.ToSQLString(t.VersionIds),
		orm.ToSQLString(t.Environment),
		orm.ToSQLString(t.ComponentIds),
		orm.ToSQLString(t.LabelIds),
		orm.ToSQLInt64(t.DuedateAt),
		orm.ToSQLInt64(t.PlannedStartAt),
		orm.ToSQLInt64(t.PlannedEndAt),
		orm.ToSQLString(t.Key),
		orm.ToSQLString(t.CustomFieldIds),
		orm.ToSQLString(t.SprintID),
		orm.ToSQLString(t.EpicID),
		orm.ToSQLString(t.ParentID),
		orm.ToSQLString(t.StrategicParentID),
		orm.ToSQLInt64(t.InProgressCount),
		orm.ToSQLInt64(t.ReopenCount),
		orm.ToSQLInt64(t.InProgressDuration),
		orm.ToSQLInt64(t.VerificationDuration),
		orm.ToSQLInt64(t.InProgressStartAt),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.UserRefID),
		orm.ToSQLFloat64(t.Cost),
		orm.ToSQLString(t.ID),
	)
}

// DBUpdateTx will update the JiraIssue record in the database using the provided transaction
func (t *JiraIssue) DBUpdateTx(ctx context.Context, tx Tx) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `jira_issue` SET `checksum`=?,`issue_id`=?,`project_id`=?,`issue_type_id`=?,`user_id`=?,`assignee_id`=?,`priority_id`=?,`status_id`=?,`resolution_id`=?,`fix_version_ids`=?,`version_ids`=?,`environment`=?,`component_ids`=?,`label_ids`=?,`duedate_at`=?,`planned_start_at`=?,`planned_end_at`=?,`key`=?,`custom_field_ids`=?,`sprint_id`=?,`epic_id`=?,`parent_id`=?,`strategic_parent_id`=?,`in_progress_count`=?,`reopen_count`=?,`in_progress_duration`=?,`verification_duration`=?,`in_progress_start_at`=?,`customer_id`=?,`ref_id`=?,`user_ref_id`=?,`cost`=? WHERE `id`=?"
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.IssueID),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLString(t.IssueTypeID),
		orm.ToSQLString(t.UserID),
		orm.ToSQLString(t.AssigneeID),
		orm.ToSQLString(t.PriorityID),
		orm.ToSQLString(t.StatusID),
		orm.ToSQLString(t.ResolutionID),
		orm.ToSQLString(t.FixVersionIds),
		orm.ToSQLString(t.VersionIds),
		orm.ToSQLString(t.Environment),
		orm.ToSQLString(t.ComponentIds),
		orm.ToSQLString(t.LabelIds),
		orm.ToSQLInt64(t.DuedateAt),
		orm.ToSQLInt64(t.PlannedStartAt),
		orm.ToSQLInt64(t.PlannedEndAt),
		orm.ToSQLString(t.Key),
		orm.ToSQLString(t.CustomFieldIds),
		orm.ToSQLString(t.SprintID),
		orm.ToSQLString(t.EpicID),
		orm.ToSQLString(t.ParentID),
		orm.ToSQLString(t.StrategicParentID),
		orm.ToSQLInt64(t.InProgressCount),
		orm.ToSQLInt64(t.ReopenCount),
		orm.ToSQLInt64(t.InProgressDuration),
		orm.ToSQLInt64(t.VerificationDuration),
		orm.ToSQLInt64(t.InProgressStartAt),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.UserRefID),
		orm.ToSQLFloat64(t.Cost),
		orm.ToSQLString(t.ID),
	)
}

// DBUpsert will upsert the JiraIssue record in the database
func (t *JiraIssue) DBUpsert(ctx context.Context, db DB, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `jira_issue` (`jira_issue`.`id`,`jira_issue`.`checksum`,`jira_issue`.`issue_id`,`jira_issue`.`project_id`,`jira_issue`.`issue_type_id`,`jira_issue`.`user_id`,`jira_issue`.`assignee_id`,`jira_issue`.`priority_id`,`jira_issue`.`status_id`,`jira_issue`.`resolution_id`,`jira_issue`.`fix_version_ids`,`jira_issue`.`version_ids`,`jira_issue`.`environment`,`jira_issue`.`component_ids`,`jira_issue`.`label_ids`,`jira_issue`.`duedate_at`,`jira_issue`.`planned_start_at`,`jira_issue`.`planned_end_at`,`jira_issue`.`key`,`jira_issue`.`custom_field_ids`,`jira_issue`.`sprint_id`,`jira_issue`.`epic_id`,`jira_issue`.`parent_id`,`jira_issue`.`strategic_parent_id`,`jira_issue`.`in_progress_count`,`jira_issue`.`reopen_count`,`jira_issue`.`in_progress_duration`,`jira_issue`.`verification_duration`,`jira_issue`.`in_progress_start_at`,`jira_issue`.`customer_id`,`jira_issue`.`ref_id`,`jira_issue`.`user_ref_id`,`jira_issue`.`cost`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `jira_issue` (`jira_issue`.`id`,`jira_issue`.`checksum`,`jira_issue`.`issue_id`,`jira_issue`.`project_id`,`jira_issue`.`issue_type_id`,`jira_issue`.`user_id`,`jira_issue`.`assignee_id`,`jira_issue`.`priority_id`,`jira_issue`.`status_id`,`jira_issue`.`resolution_id`,`jira_issue`.`fix_version_ids`,`jira_issue`.`version_ids`,`jira_issue`.`environment`,`jira_issue`.`component_ids`,`jira_issue`.`label_ids`,`jira_issue`.`duedate_at`,`jira_issue`.`planned_start_at`,`jira_issue`.`planned_end_at`,`jira_issue`.`key`,`jira_issue`.`custom_field_ids`,`jira_issue`.`sprint_id`,`jira_issue`.`epic_id`,`jira_issue`.`parent_id`,`jira_issue`.`strategic_parent_id`,`jira_issue`.`in_progress_count`,`jira_issue`.`reopen_count`,`jira_issue`.`in_progress_duration`,`jira_issue`.`verification_duration`,`jira_issue`.`in_progress_start_at`,`jira_issue`.`customer_id`,`jira_issue`.`ref_id`,`jira_issue`.`user_ref_id`,`jira_issue`.`cost`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`issue_id`=VALUES(`issue_id`),`project_id`=VALUES(`project_id`),`issue_type_id`=VALUES(`issue_type_id`),`user_id`=VALUES(`user_id`),`assignee_id`=VALUES(`assignee_id`),`priority_id`=VALUES(`priority_id`),`status_id`=VALUES(`status_id`),`resolution_id`=VALUES(`resolution_id`),`fix_version_ids`=VALUES(`fix_version_ids`),`version_ids`=VALUES(`version_ids`),`environment`=VALUES(`environment`),`component_ids`=VALUES(`component_ids`),`label_ids`=VALUES(`label_ids`),`duedate_at`=VALUES(`duedate_at`),`planned_start_at`=VALUES(`planned_start_at`),`planned_end_at`=VALUES(`planned_end_at`),`key`=VALUES(`key`),`custom_field_ids`=VALUES(`custom_field_ids`),`sprint_id`=VALUES(`sprint_id`),`epic_id`=VALUES(`epic_id`),`parent_id`=VALUES(`parent_id`),`strategic_parent_id`=VALUES(`strategic_parent_id`),`in_progress_count`=VALUES(`in_progress_count`),`reopen_count`=VALUES(`reopen_count`),`in_progress_duration`=VALUES(`in_progress_duration`),`verification_duration`=VALUES(`verification_duration`),`in_progress_start_at`=VALUES(`in_progress_start_at`),`customer_id`=VALUES(`customer_id`),`ref_id`=VALUES(`ref_id`),`user_ref_id`=VALUES(`user_ref_id`),`cost`=VALUES(`cost`)"
	}
	r, err := db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.IssueID),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLString(t.IssueTypeID),
		orm.ToSQLString(t.UserID),
		orm.ToSQLString(t.AssigneeID),
		orm.ToSQLString(t.PriorityID),
		orm.ToSQLString(t.StatusID),
		orm.ToSQLString(t.ResolutionID),
		orm.ToSQLString(t.FixVersionIds),
		orm.ToSQLString(t.VersionIds),
		orm.ToSQLString(t.Environment),
		orm.ToSQLString(t.ComponentIds),
		orm.ToSQLString(t.LabelIds),
		orm.ToSQLInt64(t.DuedateAt),
		orm.ToSQLInt64(t.PlannedStartAt),
		orm.ToSQLInt64(t.PlannedEndAt),
		orm.ToSQLString(t.Key),
		orm.ToSQLString(t.CustomFieldIds),
		orm.ToSQLString(t.SprintID),
		orm.ToSQLString(t.EpicID),
		orm.ToSQLString(t.ParentID),
		orm.ToSQLString(t.StrategicParentID),
		orm.ToSQLInt64(t.InProgressCount),
		orm.ToSQLInt64(t.ReopenCount),
		orm.ToSQLInt64(t.InProgressDuration),
		orm.ToSQLInt64(t.VerificationDuration),
		orm.ToSQLInt64(t.InProgressStartAt),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.UserRefID),
		orm.ToSQLFloat64(t.Cost),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBUpsertTx will upsert the JiraIssue record in the database using the provided transaction
func (t *JiraIssue) DBUpsertTx(ctx context.Context, tx Tx, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `jira_issue` (`jira_issue`.`id`,`jira_issue`.`checksum`,`jira_issue`.`issue_id`,`jira_issue`.`project_id`,`jira_issue`.`issue_type_id`,`jira_issue`.`user_id`,`jira_issue`.`assignee_id`,`jira_issue`.`priority_id`,`jira_issue`.`status_id`,`jira_issue`.`resolution_id`,`jira_issue`.`fix_version_ids`,`jira_issue`.`version_ids`,`jira_issue`.`environment`,`jira_issue`.`component_ids`,`jira_issue`.`label_ids`,`jira_issue`.`duedate_at`,`jira_issue`.`planned_start_at`,`jira_issue`.`planned_end_at`,`jira_issue`.`key`,`jira_issue`.`custom_field_ids`,`jira_issue`.`sprint_id`,`jira_issue`.`epic_id`,`jira_issue`.`parent_id`,`jira_issue`.`strategic_parent_id`,`jira_issue`.`in_progress_count`,`jira_issue`.`reopen_count`,`jira_issue`.`in_progress_duration`,`jira_issue`.`verification_duration`,`jira_issue`.`in_progress_start_at`,`jira_issue`.`customer_id`,`jira_issue`.`ref_id`,`jira_issue`.`user_ref_id`,`jira_issue`.`cost`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `jira_issue` (`jira_issue`.`id`,`jira_issue`.`checksum`,`jira_issue`.`issue_id`,`jira_issue`.`project_id`,`jira_issue`.`issue_type_id`,`jira_issue`.`user_id`,`jira_issue`.`assignee_id`,`jira_issue`.`priority_id`,`jira_issue`.`status_id`,`jira_issue`.`resolution_id`,`jira_issue`.`fix_version_ids`,`jira_issue`.`version_ids`,`jira_issue`.`environment`,`jira_issue`.`component_ids`,`jira_issue`.`label_ids`,`jira_issue`.`duedate_at`,`jira_issue`.`planned_start_at`,`jira_issue`.`planned_end_at`,`jira_issue`.`key`,`jira_issue`.`custom_field_ids`,`jira_issue`.`sprint_id`,`jira_issue`.`epic_id`,`jira_issue`.`parent_id`,`jira_issue`.`strategic_parent_id`,`jira_issue`.`in_progress_count`,`jira_issue`.`reopen_count`,`jira_issue`.`in_progress_duration`,`jira_issue`.`verification_duration`,`jira_issue`.`in_progress_start_at`,`jira_issue`.`customer_id`,`jira_issue`.`ref_id`,`jira_issue`.`user_ref_id`,`jira_issue`.`cost`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`issue_id`=VALUES(`issue_id`),`project_id`=VALUES(`project_id`),`issue_type_id`=VALUES(`issue_type_id`),`user_id`=VALUES(`user_id`),`assignee_id`=VALUES(`assignee_id`),`priority_id`=VALUES(`priority_id`),`status_id`=VALUES(`status_id`),`resolution_id`=VALUES(`resolution_id`),`fix_version_ids`=VALUES(`fix_version_ids`),`version_ids`=VALUES(`version_ids`),`environment`=VALUES(`environment`),`component_ids`=VALUES(`component_ids`),`label_ids`=VALUES(`label_ids`),`duedate_at`=VALUES(`duedate_at`),`planned_start_at`=VALUES(`planned_start_at`),`planned_end_at`=VALUES(`planned_end_at`),`key`=VALUES(`key`),`custom_field_ids`=VALUES(`custom_field_ids`),`sprint_id`=VALUES(`sprint_id`),`epic_id`=VALUES(`epic_id`),`parent_id`=VALUES(`parent_id`),`strategic_parent_id`=VALUES(`strategic_parent_id`),`in_progress_count`=VALUES(`in_progress_count`),`reopen_count`=VALUES(`reopen_count`),`in_progress_duration`=VALUES(`in_progress_duration`),`verification_duration`=VALUES(`verification_duration`),`in_progress_start_at`=VALUES(`in_progress_start_at`),`customer_id`=VALUES(`customer_id`),`ref_id`=VALUES(`ref_id`),`user_ref_id`=VALUES(`user_ref_id`),`cost`=VALUES(`cost`)"
	}
	r, err := tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.IssueID),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLString(t.IssueTypeID),
		orm.ToSQLString(t.UserID),
		orm.ToSQLString(t.AssigneeID),
		orm.ToSQLString(t.PriorityID),
		orm.ToSQLString(t.StatusID),
		orm.ToSQLString(t.ResolutionID),
		orm.ToSQLString(t.FixVersionIds),
		orm.ToSQLString(t.VersionIds),
		orm.ToSQLString(t.Environment),
		orm.ToSQLString(t.ComponentIds),
		orm.ToSQLString(t.LabelIds),
		orm.ToSQLInt64(t.DuedateAt),
		orm.ToSQLInt64(t.PlannedStartAt),
		orm.ToSQLInt64(t.PlannedEndAt),
		orm.ToSQLString(t.Key),
		orm.ToSQLString(t.CustomFieldIds),
		orm.ToSQLString(t.SprintID),
		orm.ToSQLString(t.EpicID),
		orm.ToSQLString(t.ParentID),
		orm.ToSQLString(t.StrategicParentID),
		orm.ToSQLInt64(t.InProgressCount),
		orm.ToSQLInt64(t.ReopenCount),
		orm.ToSQLInt64(t.InProgressDuration),
		orm.ToSQLInt64(t.VerificationDuration),
		orm.ToSQLInt64(t.InProgressStartAt),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.UserRefID),
		orm.ToSQLFloat64(t.Cost),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBFindOne will find a JiraIssue record in the database with the primary key
func (t *JiraIssue) DBFindOne(ctx context.Context, db DB, value string) (bool, error) {
	q := "SELECT * FROM `jira_issue` WHERE `id` = ? LIMIT 1"
	row := db.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _IssueID sql.NullString
	var _ProjectID sql.NullString
	var _IssueTypeID sql.NullString
	var _UserID sql.NullString
	var _AssigneeID sql.NullString
	var _PriorityID sql.NullString
	var _StatusID sql.NullString
	var _ResolutionID sql.NullString
	var _FixVersionIds sql.NullString
	var _VersionIds sql.NullString
	var _Environment sql.NullString
	var _ComponentIds sql.NullString
	var _LabelIds sql.NullString
	var _DuedateAt sql.NullInt64
	var _PlannedStartAt sql.NullInt64
	var _PlannedEndAt sql.NullInt64
	var _Key sql.NullString
	var _CustomFieldIds sql.NullString
	var _SprintID sql.NullString
	var _EpicID sql.NullString
	var _ParentID sql.NullString
	var _StrategicParentID sql.NullString
	var _InProgressCount sql.NullInt64
	var _ReopenCount sql.NullInt64
	var _InProgressDuration sql.NullInt64
	var _VerificationDuration sql.NullInt64
	var _InProgressStartAt sql.NullInt64
	var _CustomerID sql.NullString
	var _RefID sql.NullString
	var _UserRefID sql.NullString
	var _Cost sql.NullFloat64
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_IssueID,
		&_ProjectID,
		&_IssueTypeID,
		&_UserID,
		&_AssigneeID,
		&_PriorityID,
		&_StatusID,
		&_ResolutionID,
		&_FixVersionIds,
		&_VersionIds,
		&_Environment,
		&_ComponentIds,
		&_LabelIds,
		&_DuedateAt,
		&_PlannedStartAt,
		&_PlannedEndAt,
		&_Key,
		&_CustomFieldIds,
		&_SprintID,
		&_EpicID,
		&_ParentID,
		&_StrategicParentID,
		&_InProgressCount,
		&_ReopenCount,
		&_InProgressDuration,
		&_VerificationDuration,
		&_InProgressStartAt,
		&_CustomerID,
		&_RefID,
		&_UserRefID,
		&_Cost,
	)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	if _ID.Valid == false {
		return false, nil
	}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _Checksum.Valid {
		t.SetChecksum(_Checksum.String)
	}
	if _IssueID.Valid {
		t.SetIssueID(_IssueID.String)
	}
	if _ProjectID.Valid {
		t.SetProjectID(_ProjectID.String)
	}
	if _IssueTypeID.Valid {
		t.SetIssueTypeID(_IssueTypeID.String)
	}
	if _UserID.Valid {
		t.SetUserID(_UserID.String)
	}
	if _AssigneeID.Valid {
		t.SetAssigneeID(_AssigneeID.String)
	}
	if _PriorityID.Valid {
		t.SetPriorityID(_PriorityID.String)
	}
	if _StatusID.Valid {
		t.SetStatusID(_StatusID.String)
	}
	if _ResolutionID.Valid {
		t.SetResolutionID(_ResolutionID.String)
	}
	if _FixVersionIds.Valid {
		t.SetFixVersionIds(_FixVersionIds.String)
	}
	if _VersionIds.Valid {
		t.SetVersionIds(_VersionIds.String)
	}
	if _Environment.Valid {
		t.SetEnvironment(_Environment.String)
	}
	if _ComponentIds.Valid {
		t.SetComponentIds(_ComponentIds.String)
	}
	if _LabelIds.Valid {
		t.SetLabelIds(_LabelIds.String)
	}
	if _DuedateAt.Valid {
		t.SetDuedateAt(_DuedateAt.Int64)
	}
	if _PlannedStartAt.Valid {
		t.SetPlannedStartAt(_PlannedStartAt.Int64)
	}
	if _PlannedEndAt.Valid {
		t.SetPlannedEndAt(_PlannedEndAt.Int64)
	}
	if _Key.Valid {
		t.SetKey(_Key.String)
	}
	if _CustomFieldIds.Valid {
		t.SetCustomFieldIds(_CustomFieldIds.String)
	}
	if _SprintID.Valid {
		t.SetSprintID(_SprintID.String)
	}
	if _EpicID.Valid {
		t.SetEpicID(_EpicID.String)
	}
	if _ParentID.Valid {
		t.SetParentID(_ParentID.String)
	}
	if _StrategicParentID.Valid {
		t.SetStrategicParentID(_StrategicParentID.String)
	}
	if _InProgressCount.Valid {
		t.SetInProgressCount(int32(_InProgressCount.Int64))
	}
	if _ReopenCount.Valid {
		t.SetReopenCount(int32(_ReopenCount.Int64))
	}
	if _InProgressDuration.Valid {
		t.SetInProgressDuration(_InProgressDuration.Int64)
	}
	if _VerificationDuration.Valid {
		t.SetVerificationDuration(_VerificationDuration.Int64)
	}
	if _InProgressStartAt.Valid {
		t.SetInProgressStartAt(_InProgressStartAt.Int64)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _RefID.Valid {
		t.SetRefID(_RefID.String)
	}
	if _UserRefID.Valid {
		t.SetUserRefID(_UserRefID.String)
	}
	if _Cost.Valid {
		t.SetCost(_Cost.Float64)
	}
	return true, nil
}

// DBFindOneTx will find a JiraIssue record in the database with the primary key using the provided transaction
func (t *JiraIssue) DBFindOneTx(ctx context.Context, tx Tx, value string) (bool, error) {
	q := "SELECT * FROM `jira_issue` WHERE `id` = ? LIMIT 1"
	row := tx.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _IssueID sql.NullString
	var _ProjectID sql.NullString
	var _IssueTypeID sql.NullString
	var _UserID sql.NullString
	var _AssigneeID sql.NullString
	var _PriorityID sql.NullString
	var _StatusID sql.NullString
	var _ResolutionID sql.NullString
	var _FixVersionIds sql.NullString
	var _VersionIds sql.NullString
	var _Environment sql.NullString
	var _ComponentIds sql.NullString
	var _LabelIds sql.NullString
	var _DuedateAt sql.NullInt64
	var _PlannedStartAt sql.NullInt64
	var _PlannedEndAt sql.NullInt64
	var _Key sql.NullString
	var _CustomFieldIds sql.NullString
	var _SprintID sql.NullString
	var _EpicID sql.NullString
	var _ParentID sql.NullString
	var _StrategicParentID sql.NullString
	var _InProgressCount sql.NullInt64
	var _ReopenCount sql.NullInt64
	var _InProgressDuration sql.NullInt64
	var _VerificationDuration sql.NullInt64
	var _InProgressStartAt sql.NullInt64
	var _CustomerID sql.NullString
	var _RefID sql.NullString
	var _UserRefID sql.NullString
	var _Cost sql.NullFloat64
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_IssueID,
		&_ProjectID,
		&_IssueTypeID,
		&_UserID,
		&_AssigneeID,
		&_PriorityID,
		&_StatusID,
		&_ResolutionID,
		&_FixVersionIds,
		&_VersionIds,
		&_Environment,
		&_ComponentIds,
		&_LabelIds,
		&_DuedateAt,
		&_PlannedStartAt,
		&_PlannedEndAt,
		&_Key,
		&_CustomFieldIds,
		&_SprintID,
		&_EpicID,
		&_ParentID,
		&_StrategicParentID,
		&_InProgressCount,
		&_ReopenCount,
		&_InProgressDuration,
		&_VerificationDuration,
		&_InProgressStartAt,
		&_CustomerID,
		&_RefID,
		&_UserRefID,
		&_Cost,
	)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	if _ID.Valid == false {
		return false, nil
	}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _Checksum.Valid {
		t.SetChecksum(_Checksum.String)
	}
	if _IssueID.Valid {
		t.SetIssueID(_IssueID.String)
	}
	if _ProjectID.Valid {
		t.SetProjectID(_ProjectID.String)
	}
	if _IssueTypeID.Valid {
		t.SetIssueTypeID(_IssueTypeID.String)
	}
	if _UserID.Valid {
		t.SetUserID(_UserID.String)
	}
	if _AssigneeID.Valid {
		t.SetAssigneeID(_AssigneeID.String)
	}
	if _PriorityID.Valid {
		t.SetPriorityID(_PriorityID.String)
	}
	if _StatusID.Valid {
		t.SetStatusID(_StatusID.String)
	}
	if _ResolutionID.Valid {
		t.SetResolutionID(_ResolutionID.String)
	}
	if _FixVersionIds.Valid {
		t.SetFixVersionIds(_FixVersionIds.String)
	}
	if _VersionIds.Valid {
		t.SetVersionIds(_VersionIds.String)
	}
	if _Environment.Valid {
		t.SetEnvironment(_Environment.String)
	}
	if _ComponentIds.Valid {
		t.SetComponentIds(_ComponentIds.String)
	}
	if _LabelIds.Valid {
		t.SetLabelIds(_LabelIds.String)
	}
	if _DuedateAt.Valid {
		t.SetDuedateAt(_DuedateAt.Int64)
	}
	if _PlannedStartAt.Valid {
		t.SetPlannedStartAt(_PlannedStartAt.Int64)
	}
	if _PlannedEndAt.Valid {
		t.SetPlannedEndAt(_PlannedEndAt.Int64)
	}
	if _Key.Valid {
		t.SetKey(_Key.String)
	}
	if _CustomFieldIds.Valid {
		t.SetCustomFieldIds(_CustomFieldIds.String)
	}
	if _SprintID.Valid {
		t.SetSprintID(_SprintID.String)
	}
	if _EpicID.Valid {
		t.SetEpicID(_EpicID.String)
	}
	if _ParentID.Valid {
		t.SetParentID(_ParentID.String)
	}
	if _StrategicParentID.Valid {
		t.SetStrategicParentID(_StrategicParentID.String)
	}
	if _InProgressCount.Valid {
		t.SetInProgressCount(int32(_InProgressCount.Int64))
	}
	if _ReopenCount.Valid {
		t.SetReopenCount(int32(_ReopenCount.Int64))
	}
	if _InProgressDuration.Valid {
		t.SetInProgressDuration(_InProgressDuration.Int64)
	}
	if _VerificationDuration.Valid {
		t.SetVerificationDuration(_VerificationDuration.Int64)
	}
	if _InProgressStartAt.Valid {
		t.SetInProgressStartAt(_InProgressStartAt.Int64)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _RefID.Valid {
		t.SetRefID(_RefID.String)
	}
	if _UserRefID.Valid {
		t.SetUserRefID(_UserRefID.String)
	}
	if _Cost.Valid {
		t.SetCost(_Cost.Float64)
	}
	return true, nil
}

// FindJiraIssues will find a JiraIssue record in the database with the provided parameters
func FindJiraIssues(ctx context.Context, db DB, _params ...interface{}) ([]*JiraIssue, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("issue_id"),
		orm.Column("project_id"),
		orm.Column("issue_type_id"),
		orm.Column("user_id"),
		orm.Column("assignee_id"),
		orm.Column("priority_id"),
		orm.Column("status_id"),
		orm.Column("resolution_id"),
		orm.Column("fix_version_ids"),
		orm.Column("version_ids"),
		orm.Column("environment"),
		orm.Column("component_ids"),
		orm.Column("label_ids"),
		orm.Column("duedate_at"),
		orm.Column("planned_start_at"),
		orm.Column("planned_end_at"),
		orm.Column("key"),
		orm.Column("custom_field_ids"),
		orm.Column("sprint_id"),
		orm.Column("epic_id"),
		orm.Column("parent_id"),
		orm.Column("strategic_parent_id"),
		orm.Column("in_progress_count"),
		orm.Column("reopen_count"),
		orm.Column("in_progress_duration"),
		orm.Column("verification_duration"),
		orm.Column("in_progress_start_at"),
		orm.Column("customer_id"),
		orm.Column("ref_id"),
		orm.Column("user_ref_id"),
		orm.Column("cost"),
		orm.Table(JiraIssueTableName),
	}
	if len(_params) > 0 {
		for _, param := range _params {
			params = append(params, param)
		}
	}
	q, p := orm.BuildQuery(params...)
	rows, err := db.QueryContext(ctx, q, p...)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraIssue, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _IssueID sql.NullString
		var _ProjectID sql.NullString
		var _IssueTypeID sql.NullString
		var _UserID sql.NullString
		var _AssigneeID sql.NullString
		var _PriorityID sql.NullString
		var _StatusID sql.NullString
		var _ResolutionID sql.NullString
		var _FixVersionIds sql.NullString
		var _VersionIds sql.NullString
		var _Environment sql.NullString
		var _ComponentIds sql.NullString
		var _LabelIds sql.NullString
		var _DuedateAt sql.NullInt64
		var _PlannedStartAt sql.NullInt64
		var _PlannedEndAt sql.NullInt64
		var _Key sql.NullString
		var _CustomFieldIds sql.NullString
		var _SprintID sql.NullString
		var _EpicID sql.NullString
		var _ParentID sql.NullString
		var _StrategicParentID sql.NullString
		var _InProgressCount sql.NullInt64
		var _ReopenCount sql.NullInt64
		var _InProgressDuration sql.NullInt64
		var _VerificationDuration sql.NullInt64
		var _InProgressStartAt sql.NullInt64
		var _CustomerID sql.NullString
		var _RefID sql.NullString
		var _UserRefID sql.NullString
		var _Cost sql.NullFloat64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_IssueID,
			&_ProjectID,
			&_IssueTypeID,
			&_UserID,
			&_AssigneeID,
			&_PriorityID,
			&_StatusID,
			&_ResolutionID,
			&_FixVersionIds,
			&_VersionIds,
			&_Environment,
			&_ComponentIds,
			&_LabelIds,
			&_DuedateAt,
			&_PlannedStartAt,
			&_PlannedEndAt,
			&_Key,
			&_CustomFieldIds,
			&_SprintID,
			&_EpicID,
			&_ParentID,
			&_StrategicParentID,
			&_InProgressCount,
			&_ReopenCount,
			&_InProgressDuration,
			&_VerificationDuration,
			&_InProgressStartAt,
			&_CustomerID,
			&_RefID,
			&_UserRefID,
			&_Cost,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraIssue{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _IssueTypeID.Valid {
			t.SetIssueTypeID(_IssueTypeID.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _AssigneeID.Valid {
			t.SetAssigneeID(_AssigneeID.String)
		}
		if _PriorityID.Valid {
			t.SetPriorityID(_PriorityID.String)
		}
		if _StatusID.Valid {
			t.SetStatusID(_StatusID.String)
		}
		if _ResolutionID.Valid {
			t.SetResolutionID(_ResolutionID.String)
		}
		if _FixVersionIds.Valid {
			t.SetFixVersionIds(_FixVersionIds.String)
		}
		if _VersionIds.Valid {
			t.SetVersionIds(_VersionIds.String)
		}
		if _Environment.Valid {
			t.SetEnvironment(_Environment.String)
		}
		if _ComponentIds.Valid {
			t.SetComponentIds(_ComponentIds.String)
		}
		if _LabelIds.Valid {
			t.SetLabelIds(_LabelIds.String)
		}
		if _DuedateAt.Valid {
			t.SetDuedateAt(_DuedateAt.Int64)
		}
		if _PlannedStartAt.Valid {
			t.SetPlannedStartAt(_PlannedStartAt.Int64)
		}
		if _PlannedEndAt.Valid {
			t.SetPlannedEndAt(_PlannedEndAt.Int64)
		}
		if _Key.Valid {
			t.SetKey(_Key.String)
		}
		if _CustomFieldIds.Valid {
			t.SetCustomFieldIds(_CustomFieldIds.String)
		}
		if _SprintID.Valid {
			t.SetSprintID(_SprintID.String)
		}
		if _EpicID.Valid {
			t.SetEpicID(_EpicID.String)
		}
		if _ParentID.Valid {
			t.SetParentID(_ParentID.String)
		}
		if _StrategicParentID.Valid {
			t.SetStrategicParentID(_StrategicParentID.String)
		}
		if _InProgressCount.Valid {
			t.SetInProgressCount(int32(_InProgressCount.Int64))
		}
		if _ReopenCount.Valid {
			t.SetReopenCount(int32(_ReopenCount.Int64))
		}
		if _InProgressDuration.Valid {
			t.SetInProgressDuration(_InProgressDuration.Int64)
		}
		if _VerificationDuration.Valid {
			t.SetVerificationDuration(_VerificationDuration.Int64)
		}
		if _InProgressStartAt.Valid {
			t.SetInProgressStartAt(_InProgressStartAt.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		if _UserRefID.Valid {
			t.SetUserRefID(_UserRefID.String)
		}
		if _Cost.Valid {
			t.SetCost(_Cost.Float64)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindJiraIssuesTx will find a JiraIssue record in the database with the provided parameters using the provided transaction
func FindJiraIssuesTx(ctx context.Context, tx Tx, _params ...interface{}) ([]*JiraIssue, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("issue_id"),
		orm.Column("project_id"),
		orm.Column("issue_type_id"),
		orm.Column("user_id"),
		orm.Column("assignee_id"),
		orm.Column("priority_id"),
		orm.Column("status_id"),
		orm.Column("resolution_id"),
		orm.Column("fix_version_ids"),
		orm.Column("version_ids"),
		orm.Column("environment"),
		orm.Column("component_ids"),
		orm.Column("label_ids"),
		orm.Column("duedate_at"),
		orm.Column("planned_start_at"),
		orm.Column("planned_end_at"),
		orm.Column("key"),
		orm.Column("custom_field_ids"),
		orm.Column("sprint_id"),
		orm.Column("epic_id"),
		orm.Column("parent_id"),
		orm.Column("strategic_parent_id"),
		orm.Column("in_progress_count"),
		orm.Column("reopen_count"),
		orm.Column("in_progress_duration"),
		orm.Column("verification_duration"),
		orm.Column("in_progress_start_at"),
		orm.Column("customer_id"),
		orm.Column("ref_id"),
		orm.Column("user_ref_id"),
		orm.Column("cost"),
		orm.Table(JiraIssueTableName),
	}
	if len(_params) > 0 {
		for _, param := range _params {
			params = append(params, param)
		}
	}
	q, p := orm.BuildQuery(params...)
	rows, err := tx.QueryContext(ctx, q, p...)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraIssue, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _IssueID sql.NullString
		var _ProjectID sql.NullString
		var _IssueTypeID sql.NullString
		var _UserID sql.NullString
		var _AssigneeID sql.NullString
		var _PriorityID sql.NullString
		var _StatusID sql.NullString
		var _ResolutionID sql.NullString
		var _FixVersionIds sql.NullString
		var _VersionIds sql.NullString
		var _Environment sql.NullString
		var _ComponentIds sql.NullString
		var _LabelIds sql.NullString
		var _DuedateAt sql.NullInt64
		var _PlannedStartAt sql.NullInt64
		var _PlannedEndAt sql.NullInt64
		var _Key sql.NullString
		var _CustomFieldIds sql.NullString
		var _SprintID sql.NullString
		var _EpicID sql.NullString
		var _ParentID sql.NullString
		var _StrategicParentID sql.NullString
		var _InProgressCount sql.NullInt64
		var _ReopenCount sql.NullInt64
		var _InProgressDuration sql.NullInt64
		var _VerificationDuration sql.NullInt64
		var _InProgressStartAt sql.NullInt64
		var _CustomerID sql.NullString
		var _RefID sql.NullString
		var _UserRefID sql.NullString
		var _Cost sql.NullFloat64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_IssueID,
			&_ProjectID,
			&_IssueTypeID,
			&_UserID,
			&_AssigneeID,
			&_PriorityID,
			&_StatusID,
			&_ResolutionID,
			&_FixVersionIds,
			&_VersionIds,
			&_Environment,
			&_ComponentIds,
			&_LabelIds,
			&_DuedateAt,
			&_PlannedStartAt,
			&_PlannedEndAt,
			&_Key,
			&_CustomFieldIds,
			&_SprintID,
			&_EpicID,
			&_ParentID,
			&_StrategicParentID,
			&_InProgressCount,
			&_ReopenCount,
			&_InProgressDuration,
			&_VerificationDuration,
			&_InProgressStartAt,
			&_CustomerID,
			&_RefID,
			&_UserRefID,
			&_Cost,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraIssue{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _IssueTypeID.Valid {
			t.SetIssueTypeID(_IssueTypeID.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _AssigneeID.Valid {
			t.SetAssigneeID(_AssigneeID.String)
		}
		if _PriorityID.Valid {
			t.SetPriorityID(_PriorityID.String)
		}
		if _StatusID.Valid {
			t.SetStatusID(_StatusID.String)
		}
		if _ResolutionID.Valid {
			t.SetResolutionID(_ResolutionID.String)
		}
		if _FixVersionIds.Valid {
			t.SetFixVersionIds(_FixVersionIds.String)
		}
		if _VersionIds.Valid {
			t.SetVersionIds(_VersionIds.String)
		}
		if _Environment.Valid {
			t.SetEnvironment(_Environment.String)
		}
		if _ComponentIds.Valid {
			t.SetComponentIds(_ComponentIds.String)
		}
		if _LabelIds.Valid {
			t.SetLabelIds(_LabelIds.String)
		}
		if _DuedateAt.Valid {
			t.SetDuedateAt(_DuedateAt.Int64)
		}
		if _PlannedStartAt.Valid {
			t.SetPlannedStartAt(_PlannedStartAt.Int64)
		}
		if _PlannedEndAt.Valid {
			t.SetPlannedEndAt(_PlannedEndAt.Int64)
		}
		if _Key.Valid {
			t.SetKey(_Key.String)
		}
		if _CustomFieldIds.Valid {
			t.SetCustomFieldIds(_CustomFieldIds.String)
		}
		if _SprintID.Valid {
			t.SetSprintID(_SprintID.String)
		}
		if _EpicID.Valid {
			t.SetEpicID(_EpicID.String)
		}
		if _ParentID.Valid {
			t.SetParentID(_ParentID.String)
		}
		if _StrategicParentID.Valid {
			t.SetStrategicParentID(_StrategicParentID.String)
		}
		if _InProgressCount.Valid {
			t.SetInProgressCount(int32(_InProgressCount.Int64))
		}
		if _ReopenCount.Valid {
			t.SetReopenCount(int32(_ReopenCount.Int64))
		}
		if _InProgressDuration.Valid {
			t.SetInProgressDuration(_InProgressDuration.Int64)
		}
		if _VerificationDuration.Valid {
			t.SetVerificationDuration(_VerificationDuration.Int64)
		}
		if _InProgressStartAt.Valid {
			t.SetInProgressStartAt(_InProgressStartAt.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		if _UserRefID.Valid {
			t.SetUserRefID(_UserRefID.String)
		}
		if _Cost.Valid {
			t.SetCost(_Cost.Float64)
		}
		results = append(results, t)
	}
	return results, nil
}

// DBFind will find a JiraIssue record in the database with the provided parameters
func (t *JiraIssue) DBFind(ctx context.Context, db DB, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("issue_id"),
		orm.Column("project_id"),
		orm.Column("issue_type_id"),
		orm.Column("user_id"),
		orm.Column("assignee_id"),
		orm.Column("priority_id"),
		orm.Column("status_id"),
		orm.Column("resolution_id"),
		orm.Column("fix_version_ids"),
		orm.Column("version_ids"),
		orm.Column("environment"),
		orm.Column("component_ids"),
		orm.Column("label_ids"),
		orm.Column("duedate_at"),
		orm.Column("planned_start_at"),
		orm.Column("planned_end_at"),
		orm.Column("key"),
		orm.Column("custom_field_ids"),
		orm.Column("sprint_id"),
		orm.Column("epic_id"),
		orm.Column("parent_id"),
		orm.Column("strategic_parent_id"),
		orm.Column("in_progress_count"),
		orm.Column("reopen_count"),
		orm.Column("in_progress_duration"),
		orm.Column("verification_duration"),
		orm.Column("in_progress_start_at"),
		orm.Column("customer_id"),
		orm.Column("ref_id"),
		orm.Column("user_ref_id"),
		orm.Column("cost"),
		orm.Table(JiraIssueTableName),
	}
	if len(_params) > 0 {
		for _, param := range _params {
			params = append(params, param)
		}
	}
	q, p := orm.BuildQuery(params...)
	row := db.QueryRowContext(ctx, q, p...)
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _IssueID sql.NullString
	var _ProjectID sql.NullString
	var _IssueTypeID sql.NullString
	var _UserID sql.NullString
	var _AssigneeID sql.NullString
	var _PriorityID sql.NullString
	var _StatusID sql.NullString
	var _ResolutionID sql.NullString
	var _FixVersionIds sql.NullString
	var _VersionIds sql.NullString
	var _Environment sql.NullString
	var _ComponentIds sql.NullString
	var _LabelIds sql.NullString
	var _DuedateAt sql.NullInt64
	var _PlannedStartAt sql.NullInt64
	var _PlannedEndAt sql.NullInt64
	var _Key sql.NullString
	var _CustomFieldIds sql.NullString
	var _SprintID sql.NullString
	var _EpicID sql.NullString
	var _ParentID sql.NullString
	var _StrategicParentID sql.NullString
	var _InProgressCount sql.NullInt64
	var _ReopenCount sql.NullInt64
	var _InProgressDuration sql.NullInt64
	var _VerificationDuration sql.NullInt64
	var _InProgressStartAt sql.NullInt64
	var _CustomerID sql.NullString
	var _RefID sql.NullString
	var _UserRefID sql.NullString
	var _Cost sql.NullFloat64
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_IssueID,
		&_ProjectID,
		&_IssueTypeID,
		&_UserID,
		&_AssigneeID,
		&_PriorityID,
		&_StatusID,
		&_ResolutionID,
		&_FixVersionIds,
		&_VersionIds,
		&_Environment,
		&_ComponentIds,
		&_LabelIds,
		&_DuedateAt,
		&_PlannedStartAt,
		&_PlannedEndAt,
		&_Key,
		&_CustomFieldIds,
		&_SprintID,
		&_EpicID,
		&_ParentID,
		&_StrategicParentID,
		&_InProgressCount,
		&_ReopenCount,
		&_InProgressDuration,
		&_VerificationDuration,
		&_InProgressStartAt,
		&_CustomerID,
		&_RefID,
		&_UserRefID,
		&_Cost,
	)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _Checksum.Valid {
		t.SetChecksum(_Checksum.String)
	}
	if _IssueID.Valid {
		t.SetIssueID(_IssueID.String)
	}
	if _ProjectID.Valid {
		t.SetProjectID(_ProjectID.String)
	}
	if _IssueTypeID.Valid {
		t.SetIssueTypeID(_IssueTypeID.String)
	}
	if _UserID.Valid {
		t.SetUserID(_UserID.String)
	}
	if _AssigneeID.Valid {
		t.SetAssigneeID(_AssigneeID.String)
	}
	if _PriorityID.Valid {
		t.SetPriorityID(_PriorityID.String)
	}
	if _StatusID.Valid {
		t.SetStatusID(_StatusID.String)
	}
	if _ResolutionID.Valid {
		t.SetResolutionID(_ResolutionID.String)
	}
	if _FixVersionIds.Valid {
		t.SetFixVersionIds(_FixVersionIds.String)
	}
	if _VersionIds.Valid {
		t.SetVersionIds(_VersionIds.String)
	}
	if _Environment.Valid {
		t.SetEnvironment(_Environment.String)
	}
	if _ComponentIds.Valid {
		t.SetComponentIds(_ComponentIds.String)
	}
	if _LabelIds.Valid {
		t.SetLabelIds(_LabelIds.String)
	}
	if _DuedateAt.Valid {
		t.SetDuedateAt(_DuedateAt.Int64)
	}
	if _PlannedStartAt.Valid {
		t.SetPlannedStartAt(_PlannedStartAt.Int64)
	}
	if _PlannedEndAt.Valid {
		t.SetPlannedEndAt(_PlannedEndAt.Int64)
	}
	if _Key.Valid {
		t.SetKey(_Key.String)
	}
	if _CustomFieldIds.Valid {
		t.SetCustomFieldIds(_CustomFieldIds.String)
	}
	if _SprintID.Valid {
		t.SetSprintID(_SprintID.String)
	}
	if _EpicID.Valid {
		t.SetEpicID(_EpicID.String)
	}
	if _ParentID.Valid {
		t.SetParentID(_ParentID.String)
	}
	if _StrategicParentID.Valid {
		t.SetStrategicParentID(_StrategicParentID.String)
	}
	if _InProgressCount.Valid {
		t.SetInProgressCount(int32(_InProgressCount.Int64))
	}
	if _ReopenCount.Valid {
		t.SetReopenCount(int32(_ReopenCount.Int64))
	}
	if _InProgressDuration.Valid {
		t.SetInProgressDuration(_InProgressDuration.Int64)
	}
	if _VerificationDuration.Valid {
		t.SetVerificationDuration(_VerificationDuration.Int64)
	}
	if _InProgressStartAt.Valid {
		t.SetInProgressStartAt(_InProgressStartAt.Int64)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _RefID.Valid {
		t.SetRefID(_RefID.String)
	}
	if _UserRefID.Valid {
		t.SetUserRefID(_UserRefID.String)
	}
	if _Cost.Valid {
		t.SetCost(_Cost.Float64)
	}
	return true, nil
}

// DBFindTx will find a JiraIssue record in the database with the provided parameters using the provided transaction
func (t *JiraIssue) DBFindTx(ctx context.Context, tx Tx, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("issue_id"),
		orm.Column("project_id"),
		orm.Column("issue_type_id"),
		orm.Column("user_id"),
		orm.Column("assignee_id"),
		orm.Column("priority_id"),
		orm.Column("status_id"),
		orm.Column("resolution_id"),
		orm.Column("fix_version_ids"),
		orm.Column("version_ids"),
		orm.Column("environment"),
		orm.Column("component_ids"),
		orm.Column("label_ids"),
		orm.Column("duedate_at"),
		orm.Column("planned_start_at"),
		orm.Column("planned_end_at"),
		orm.Column("key"),
		orm.Column("custom_field_ids"),
		orm.Column("sprint_id"),
		orm.Column("epic_id"),
		orm.Column("parent_id"),
		orm.Column("strategic_parent_id"),
		orm.Column("in_progress_count"),
		orm.Column("reopen_count"),
		orm.Column("in_progress_duration"),
		orm.Column("verification_duration"),
		orm.Column("in_progress_start_at"),
		orm.Column("customer_id"),
		orm.Column("ref_id"),
		orm.Column("user_ref_id"),
		orm.Column("cost"),
		orm.Table(JiraIssueTableName),
	}
	if len(_params) > 0 {
		for _, param := range _params {
			params = append(params, param)
		}
	}
	q, p := orm.BuildQuery(params...)
	row := tx.QueryRowContext(ctx, q, p...)
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _IssueID sql.NullString
	var _ProjectID sql.NullString
	var _IssueTypeID sql.NullString
	var _UserID sql.NullString
	var _AssigneeID sql.NullString
	var _PriorityID sql.NullString
	var _StatusID sql.NullString
	var _ResolutionID sql.NullString
	var _FixVersionIds sql.NullString
	var _VersionIds sql.NullString
	var _Environment sql.NullString
	var _ComponentIds sql.NullString
	var _LabelIds sql.NullString
	var _DuedateAt sql.NullInt64
	var _PlannedStartAt sql.NullInt64
	var _PlannedEndAt sql.NullInt64
	var _Key sql.NullString
	var _CustomFieldIds sql.NullString
	var _SprintID sql.NullString
	var _EpicID sql.NullString
	var _ParentID sql.NullString
	var _StrategicParentID sql.NullString
	var _InProgressCount sql.NullInt64
	var _ReopenCount sql.NullInt64
	var _InProgressDuration sql.NullInt64
	var _VerificationDuration sql.NullInt64
	var _InProgressStartAt sql.NullInt64
	var _CustomerID sql.NullString
	var _RefID sql.NullString
	var _UserRefID sql.NullString
	var _Cost sql.NullFloat64
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_IssueID,
		&_ProjectID,
		&_IssueTypeID,
		&_UserID,
		&_AssigneeID,
		&_PriorityID,
		&_StatusID,
		&_ResolutionID,
		&_FixVersionIds,
		&_VersionIds,
		&_Environment,
		&_ComponentIds,
		&_LabelIds,
		&_DuedateAt,
		&_PlannedStartAt,
		&_PlannedEndAt,
		&_Key,
		&_CustomFieldIds,
		&_SprintID,
		&_EpicID,
		&_ParentID,
		&_StrategicParentID,
		&_InProgressCount,
		&_ReopenCount,
		&_InProgressDuration,
		&_VerificationDuration,
		&_InProgressStartAt,
		&_CustomerID,
		&_RefID,
		&_UserRefID,
		&_Cost,
	)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _Checksum.Valid {
		t.SetChecksum(_Checksum.String)
	}
	if _IssueID.Valid {
		t.SetIssueID(_IssueID.String)
	}
	if _ProjectID.Valid {
		t.SetProjectID(_ProjectID.String)
	}
	if _IssueTypeID.Valid {
		t.SetIssueTypeID(_IssueTypeID.String)
	}
	if _UserID.Valid {
		t.SetUserID(_UserID.String)
	}
	if _AssigneeID.Valid {
		t.SetAssigneeID(_AssigneeID.String)
	}
	if _PriorityID.Valid {
		t.SetPriorityID(_PriorityID.String)
	}
	if _StatusID.Valid {
		t.SetStatusID(_StatusID.String)
	}
	if _ResolutionID.Valid {
		t.SetResolutionID(_ResolutionID.String)
	}
	if _FixVersionIds.Valid {
		t.SetFixVersionIds(_FixVersionIds.String)
	}
	if _VersionIds.Valid {
		t.SetVersionIds(_VersionIds.String)
	}
	if _Environment.Valid {
		t.SetEnvironment(_Environment.String)
	}
	if _ComponentIds.Valid {
		t.SetComponentIds(_ComponentIds.String)
	}
	if _LabelIds.Valid {
		t.SetLabelIds(_LabelIds.String)
	}
	if _DuedateAt.Valid {
		t.SetDuedateAt(_DuedateAt.Int64)
	}
	if _PlannedStartAt.Valid {
		t.SetPlannedStartAt(_PlannedStartAt.Int64)
	}
	if _PlannedEndAt.Valid {
		t.SetPlannedEndAt(_PlannedEndAt.Int64)
	}
	if _Key.Valid {
		t.SetKey(_Key.String)
	}
	if _CustomFieldIds.Valid {
		t.SetCustomFieldIds(_CustomFieldIds.String)
	}
	if _SprintID.Valid {
		t.SetSprintID(_SprintID.String)
	}
	if _EpicID.Valid {
		t.SetEpicID(_EpicID.String)
	}
	if _ParentID.Valid {
		t.SetParentID(_ParentID.String)
	}
	if _StrategicParentID.Valid {
		t.SetStrategicParentID(_StrategicParentID.String)
	}
	if _InProgressCount.Valid {
		t.SetInProgressCount(int32(_InProgressCount.Int64))
	}
	if _ReopenCount.Valid {
		t.SetReopenCount(int32(_ReopenCount.Int64))
	}
	if _InProgressDuration.Valid {
		t.SetInProgressDuration(_InProgressDuration.Int64)
	}
	if _VerificationDuration.Valid {
		t.SetVerificationDuration(_VerificationDuration.Int64)
	}
	if _InProgressStartAt.Valid {
		t.SetInProgressStartAt(_InProgressStartAt.Int64)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _RefID.Valid {
		t.SetRefID(_RefID.String)
	}
	if _UserRefID.Valid {
		t.SetUserRefID(_UserRefID.String)
	}
	if _Cost.Valid {
		t.SetCost(_Cost.Float64)
	}
	return true, nil
}

// CountJiraIssues will find the count of JiraIssue records in the database
func CountJiraIssues(ctx context.Context, db DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(JiraIssueTableName),
	}
	if len(_params) > 0 {
		for _, param := range _params {
			params = append(params, param)
		}
	}
	q, p := orm.BuildQuery(params...)
	var count sql.NullInt64
	err := db.QueryRowContext(ctx, q, p...).Scan(&count)
	if err != nil && err != sql.ErrNoRows {
		return 0, err
	}
	return count.Int64, nil
}

// CountJiraIssuesTx will find the count of JiraIssue records in the database using the provided transaction
func CountJiraIssuesTx(ctx context.Context, tx Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(JiraIssueTableName),
	}
	if len(_params) > 0 {
		for _, param := range _params {
			params = append(params, param)
		}
	}
	q, p := orm.BuildQuery(params...)
	var count sql.NullInt64
	err := tx.QueryRowContext(ctx, q, p...).Scan(&count)
	if err != nil && err != sql.ErrNoRows {
		return 0, err
	}
	return count.Int64, nil
}

// DBCount will find the count of JiraIssue records in the database
func (t *JiraIssue) DBCount(ctx context.Context, db DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(JiraIssueTableName),
	}
	if len(_params) > 0 {
		for _, param := range _params {
			params = append(params, param)
		}
	}
	q, p := orm.BuildQuery(params...)
	var count sql.NullInt64
	err := db.QueryRowContext(ctx, q, p...).Scan(&count)
	if err != nil && err != sql.ErrNoRows {
		return 0, err
	}
	return count.Int64, nil
}

// DBCountTx will find the count of JiraIssue records in the database using the provided transaction
func (t *JiraIssue) DBCountTx(ctx context.Context, tx Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(JiraIssueTableName),
	}
	if len(_params) > 0 {
		for _, param := range _params {
			params = append(params, param)
		}
	}
	q, p := orm.BuildQuery(params...)
	var count sql.NullInt64
	err := tx.QueryRowContext(ctx, q, p...).Scan(&count)
	if err != nil && err != sql.ErrNoRows {
		return 0, err
	}
	return count.Int64, nil
}

// DBExists will return true if the JiraIssue record exists in the database
func (t *JiraIssue) DBExists(ctx context.Context, db DB) (bool, error) {
	q := "SELECT * FROM `jira_issue` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := db.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// DBExistsTx will return true if the JiraIssue record exists in the database using the provided transaction
func (t *JiraIssue) DBExistsTx(ctx context.Context, tx Tx) (bool, error) {
	q := "SELECT * FROM `jira_issue` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := tx.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// PrimaryKeyColumn returns the column name for the primary key
func (t *JiraIssue) PrimaryKeyColumn() string {
	return JiraIssueColumnID
}

// PrimaryKeyColumnType returns the primary key column Go type as a string
func (t *JiraIssue) PrimaryKeyColumnType() string {
	return "string"
}

// PrimaryKey returns the primary key column value
func (t *JiraIssue) PrimaryKey() interface{} {
	return t.ID
}
