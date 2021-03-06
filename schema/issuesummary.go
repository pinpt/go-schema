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

var _ Model = (*IssueSummary)(nil)
var _ CSVWriter = (*IssueSummary)(nil)
var _ JSONWriter = (*IssueSummary)(nil)
var _ Checksum = (*IssueSummary)(nil)

// IssueSummaryTableName is the name of the table in SQL
const IssueSummaryTableName = "issue_summary"

var IssueSummaryColumns = []string{
	"id",
	"checksum",
	"issue_id",
	"total_issues",
	"new30_days",
	"total_closed",
	"closed30_days",
	"estimated_work_months",
	"estimated_work_months30_days",
	"title",
	"url",
	"priority",
	"priority_id",
	"status",
	"status_id",
	"issue_type",
	"issue_type_id",
	"resolution",
	"resolution_id",
	"state",
	"custom_field_ids",
	"teams",
	"parent_issue_id",
	"parents_issue_ids",
	"metadata",
	"project_id",
	"sprints",
	"labels",
	"top_level",
	"is_leaf",
	"path",
	"in_progress_duration",
	"verification_duration",
	"in_progress_count",
	"reopen_count",
	"mapped_type",
	"strategic_parent_id",
	"sprint_id",
	"issue_project_name",
	"users",
	"initial_start_date",
	"total_duration",
	"created_at",
	"closed_at",
	"planned_end_date",
	"customer_id",
	"ref_type",
	"ref_id",
	"custom_field_ids_virtual",
	"release_duration",
	"completed",
	"completed_date",
}

// IssueSummary table
type IssueSummary struct {
	Checksum                  *string `json:"checksum,omitempty"`
	Closed30Days              int32   `json:"closed30_days"`
	ClosedAt                  *int64  `json:"closed_at,omitempty"`
	Completed                 bool    `json:"completed"`
	CompletedDate             *int64  `json:"completed_date,omitempty"`
	CreatedAt                 *int64  `json:"created_at,omitempty"`
	CustomFieldIds            *string `json:"custom_field_ids,omitempty"`
	CustomFieldIdsVirtual     *string `json:"custom_field_ids_virtual,omitempty"`
	CustomerID                string  `json:"customer_id"`
	EstimatedWorkMonths       float64 `json:"estimated_work_months"`
	EstimatedWorkMonths30Days float64 `json:"estimated_work_months30_days"`
	ID                        string  `json:"id"`
	InProgressCount           int32   `json:"in_progress_count"`
	InProgressDuration        int64   `json:"in_progress_duration"`
	InitialStartDate          int64   `json:"initial_start_date"`
	IsLeaf                    bool    `json:"is_leaf"`
	IssueID                   string  `json:"issue_id"`
	IssueProjectName          string  `json:"issue_project_name"`
	IssueType                 string  `json:"issue_type"`
	IssueTypeID               *string `json:"issue_type_id,omitempty"`
	Labels                    *string `json:"labels,omitempty"`
	MappedType                string  `json:"mapped_type"`
	Metadata                  *string `json:"metadata,omitempty"`
	New30Days                 int32   `json:"new30_days"`
	ParentIssueID             *string `json:"parent_issue_id,omitempty"`
	ParentsIssueIds           *string `json:"parents_issue_ids,omitempty"`
	Path                      string  `json:"path"`
	PlannedEndDate            *int64  `json:"planned_end_date,omitempty"`
	Priority                  *string `json:"priority,omitempty"`
	PriorityID                *string `json:"priority_id,omitempty"`
	ProjectID                 string  `json:"project_id"`
	RefID                     string  `json:"ref_id"`
	RefType                   string  `json:"ref_type"`
	ReleaseDuration           int64   `json:"release_duration"`
	ReopenCount               int32   `json:"reopen_count"`
	Resolution                *string `json:"resolution,omitempty"`
	ResolutionID              *string `json:"resolution_id,omitempty"`
	SprintID                  *string `json:"sprint_id,omitempty"`
	Sprints                   *string `json:"sprints,omitempty"`
	State                     string  `json:"state"`
	Status                    *string `json:"status,omitempty"`
	StatusID                  *string `json:"status_id,omitempty"`
	StrategicParentID         *string `json:"strategic_parent_id,omitempty"`
	Teams                     *string `json:"teams,omitempty"`
	Title                     string  `json:"title"`
	TopLevel                  bool    `json:"top_level"`
	TotalClosed               int32   `json:"total_closed"`
	TotalDuration             int64   `json:"total_duration"`
	TotalIssues               int32   `json:"total_issues"`
	URL                       *string `json:"url,omitempty"`
	Users                     *string `json:"users,omitempty"`
	VerificationDuration      int64   `json:"verification_duration"`
}

// TableName returns the SQL table name for IssueSummary and satifies the Model interface
func (t *IssueSummary) TableName() string {
	return IssueSummaryTableName
}

// ToCSV will serialize the IssueSummary instance to a CSV compatible array of strings
func (t *IssueSummary) ToCSV() []string {
	return []string{
		t.ID,
		t.CalculateChecksum(),
		t.IssueID,
		toCSVString(t.TotalIssues),
		toCSVString(t.New30Days),
		toCSVString(t.TotalClosed),
		toCSVString(t.Closed30Days),
		toCSVString(t.EstimatedWorkMonths),
		toCSVString(t.EstimatedWorkMonths30Days),
		t.Title,
		toCSVString(t.URL),
		toCSVString(t.Priority),
		toCSVString(t.PriorityID),
		toCSVString(t.Status),
		toCSVString(t.StatusID),
		t.IssueType,
		toCSVString(t.IssueTypeID),
		toCSVString(t.Resolution),
		toCSVString(t.ResolutionID),
		t.State,
		toCSVString(t.CustomFieldIds),
		toCSVString(t.Teams),
		toCSVString(t.ParentIssueID),
		toCSVString(t.ParentsIssueIds),
		toCSVString(t.Metadata),
		t.ProjectID,
		toCSVString(t.Sprints),
		toCSVString(t.Labels),
		toCSVBool(t.TopLevel),
		toCSVBool(t.IsLeaf),
		t.Path,
		toCSVString(t.InProgressDuration),
		toCSVString(t.VerificationDuration),
		toCSVString(t.InProgressCount),
		toCSVString(t.ReopenCount),
		t.MappedType,
		toCSVString(t.StrategicParentID),
		toCSVString(t.SprintID),
		t.IssueProjectName,
		toCSVString(t.Users),
		toCSVString(t.InitialStartDate),
		toCSVString(t.TotalDuration),
		toCSVString(t.CreatedAt),
		toCSVString(t.ClosedAt),
		toCSVString(t.PlannedEndDate),
		t.CustomerID,
		t.RefType,
		t.RefID,
		toCSVString(t.CustomFieldIdsVirtual),
		toCSVString(t.ReleaseDuration),
		toCSVBool(t.Completed),
		toCSVString(t.CompletedDate),
	}
}

// WriteCSV will serialize the IssueSummary instance to the writer as CSV and satisfies the CSVWriter interface
func (t *IssueSummary) WriteCSV(w *csv.Writer) error {
	return w.Write(t.ToCSV())
}

// WriteJSON will serialize the IssueSummary instance to the writer as JSON and satisfies the JSONWriter interface
func (t *IssueSummary) WriteJSON(w io.Writer, indent ...bool) error {
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

// NewIssueSummaryReader creates a JSON reader which can read in IssueSummary objects serialized as JSON either as an array, single object or json new lines
// and writes each IssueSummary to the channel provided
func NewIssueSummaryReader(r io.Reader, ch chan<- IssueSummary) error {
	return orm.Deserialize(r, func(buf json.RawMessage) error {
		dec := json.NewDecoder(bytes.NewBuffer(buf))
		e := IssueSummary{}
		if err := dec.Decode(&e); err != nil {
			return err
		}
		ch <- e
		return nil
	})
}

// NewCSVIssueSummaryReaderDir will read the reader as CSV and emit each record to the channel provided
func NewCSVIssueSummaryReader(r io.Reader, ch chan<- IssueSummary) error {
	cr := csv.NewReader(r)
	for {
		record, err := cr.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		ch <- IssueSummary{
			ID:                        record[0],
			Checksum:                  fromStringPointer(record[1]),
			IssueID:                   record[2],
			TotalIssues:               fromCSVInt32(record[3]),
			New30Days:                 fromCSVInt32(record[4]),
			TotalClosed:               fromCSVInt32(record[5]),
			Closed30Days:              fromCSVInt32(record[6]),
			EstimatedWorkMonths:       fromCSVFloat64(record[7]),
			EstimatedWorkMonths30Days: fromCSVFloat64(record[8]),
			Title:                     record[9],
			URL:                       fromStringPointer(record[10]),
			Priority:                  fromStringPointer(record[11]),
			PriorityID:                fromStringPointer(record[12]),
			Status:                    fromStringPointer(record[13]),
			StatusID:                  fromStringPointer(record[14]),
			IssueType:                 record[15],
			IssueTypeID:               fromStringPointer(record[16]),
			Resolution:                fromStringPointer(record[17]),
			ResolutionID:              fromStringPointer(record[18]),
			State:                     record[19],
			CustomFieldIds:            fromStringPointer(record[20]),
			Teams:                     fromStringPointer(record[21]),
			ParentIssueID:             fromStringPointer(record[22]),
			ParentsIssueIds:           fromStringPointer(record[23]),
			Metadata:                  fromStringPointer(record[24]),
			ProjectID:                 record[25],
			Sprints:                   fromStringPointer(record[26]),
			Labels:                    fromStringPointer(record[27]),
			TopLevel:                  fromCSVBool(record[28]),
			IsLeaf:                    fromCSVBool(record[29]),
			Path:                      record[30],
			InProgressDuration:        fromCSVInt64(record[31]),
			VerificationDuration:      fromCSVInt64(record[32]),
			InProgressCount:           fromCSVInt32(record[33]),
			ReopenCount:               fromCSVInt32(record[34]),
			MappedType:                record[35],
			StrategicParentID:         fromStringPointer(record[36]),
			SprintID:                  fromStringPointer(record[37]),
			IssueProjectName:          record[38],
			Users:                     fromStringPointer(record[39]),
			InitialStartDate:          fromCSVInt64(record[40]),
			TotalDuration:             fromCSVInt64(record[41]),
			CreatedAt:                 fromCSVInt64Pointer(record[42]),
			ClosedAt:                  fromCSVInt64Pointer(record[43]),
			PlannedEndDate:            fromCSVInt64Pointer(record[44]),
			CustomerID:                record[45],
			RefType:                   record[46],
			RefID:                     record[47],
			CustomFieldIdsVirtual:     fromStringPointer(record[48]),
			ReleaseDuration:           fromCSVInt64(record[49]),
			Completed:                 fromCSVBool(record[50]),
			CompletedDate:             fromCSVInt64Pointer(record[51]),
		}
	}
	return nil
}

// NewCSVIssueSummaryReaderFile will read the file as a CSV and emit each record to the channel provided
func NewCSVIssueSummaryReaderFile(fp string, ch chan<- IssueSummary) error {
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
	return NewCSVIssueSummaryReader(fc, ch)
}

// NewCSVIssueSummaryReaderDir will read the issue_summary.csv.gz file as a CSV and emit each record to the channel provided
func NewCSVIssueSummaryReaderDir(dir string, ch chan<- IssueSummary) error {
	return NewCSVIssueSummaryReaderFile(filepath.Join(dir, "issue_summary.csv.gz"), ch)
}

// IssueSummaryCSVDeduper is a function callback which takes the existing value (a) and the new value (b)
// and the return value should be the one to use (or a new one, if applicable). return nil
// to skip processing of this record
type IssueSummaryCSVDeduper func(a IssueSummary, b IssueSummary) *IssueSummary

// IssueSummaryCSVDedupeDisabled is set on whether the CSV writer should de-dupe values by key
var IssueSummaryCSVDedupeDisabled bool

// NewIssueSummaryCSVWriterSize creates a batch writer that will write each IssueSummary into a CSV file
// this method will automatically de-duplicate entries using the primary key. if the checksum
// for a newer item with the same primary key doesn't match a previously sent item, the newer
// one will be used
func NewIssueSummaryCSVWriterSize(w io.Writer, size int, dedupers ...IssueSummaryCSVDeduper) (chan IssueSummary, chan bool, error) {
	cw := csv.NewWriter(w)
	ch := make(chan IssueSummary, size)
	done := make(chan bool)
	go func() {
		defer func() { done <- true }()
		dodedupe := !IssueSummaryCSVDedupeDisabled
		var kv map[string]*IssueSummary
		var deduper IssueSummaryCSVDeduper
		if dedupers != nil && len(dedupers) > 0 {
			deduper = dedupers[0]
			dodedupe = true
		}
		if dodedupe {
			kv = make(map[string]*IssueSummary)
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

// IssueSummaryCSVDefaultSize is the default channel buffer size if not provided
var IssueSummaryCSVDefaultSize = 100

// NewIssueSummaryCSVWriter creates a batch writer that will write each IssueSummary into a CSV file
func NewIssueSummaryCSVWriter(w io.Writer, dedupers ...IssueSummaryCSVDeduper) (chan IssueSummary, chan bool, error) {
	return NewIssueSummaryCSVWriterSize(w, IssueSummaryCSVDefaultSize, dedupers...)
}

// NewIssueSummaryCSVWriterDir creates a batch writer that will write each IssueSummary into a CSV file named issue_summary.csv.gz in dir
func NewIssueSummaryCSVWriterDir(dir string, dedupers ...IssueSummaryCSVDeduper) (chan IssueSummary, chan bool, error) {
	return NewIssueSummaryCSVWriterFile(filepath.Join(dir, "issue_summary.csv.gz"), dedupers...)
}

// NewIssueSummaryCSVWriterFile creates a batch writer that will write each IssueSummary into a CSV file
func NewIssueSummaryCSVWriterFile(fn string, dedupers ...IssueSummaryCSVDeduper) (chan IssueSummary, chan bool, error) {
	f, err := os.Create(fn)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening CSV file %s. %v", fn, err)
	}
	var fc io.WriteCloser = f
	if filepath.Ext(fn) == ".gz" {
		w, _ := gzip.NewWriterLevel(f, gzip.BestCompression)
		fc = w
	}
	ch, done, err := NewIssueSummaryCSVWriter(fc, dedupers...)
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

type IssueSummaryDBAction func(ctx context.Context, db DB, record IssueSummary) error

// NewIssueSummaryDBWriterSize creates a DB writer that will write each issue into the DB
func NewIssueSummaryDBWriterSize(ctx context.Context, db DB, errors chan<- error, size int, actions ...IssueSummaryDBAction) (chan IssueSummary, chan bool, error) {
	ch := make(chan IssueSummary, size)
	done := make(chan bool)
	var action IssueSummaryDBAction
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

// NewIssueSummaryDBWriter creates a DB writer that will write each issue into the DB
func NewIssueSummaryDBWriter(ctx context.Context, db DB, errors chan<- error, actions ...IssueSummaryDBAction) (chan IssueSummary, chan bool, error) {
	return NewIssueSummaryDBWriterSize(ctx, db, errors, 100, actions...)
}

// IssueSummaryColumnID is the ID SQL column name for the IssueSummary table
const IssueSummaryColumnID = "id"

// IssueSummaryEscapedColumnID is the escaped ID SQL column name for the IssueSummary table
const IssueSummaryEscapedColumnID = "`id`"

// IssueSummaryColumnChecksum is the Checksum SQL column name for the IssueSummary table
const IssueSummaryColumnChecksum = "checksum"

// IssueSummaryEscapedColumnChecksum is the escaped Checksum SQL column name for the IssueSummary table
const IssueSummaryEscapedColumnChecksum = "`checksum`"

// IssueSummaryColumnIssueID is the IssueID SQL column name for the IssueSummary table
const IssueSummaryColumnIssueID = "issue_id"

// IssueSummaryEscapedColumnIssueID is the escaped IssueID SQL column name for the IssueSummary table
const IssueSummaryEscapedColumnIssueID = "`issue_id`"

// IssueSummaryColumnTotalIssues is the TotalIssues SQL column name for the IssueSummary table
const IssueSummaryColumnTotalIssues = "total_issues"

// IssueSummaryEscapedColumnTotalIssues is the escaped TotalIssues SQL column name for the IssueSummary table
const IssueSummaryEscapedColumnTotalIssues = "`total_issues`"

// IssueSummaryColumnNew30Days is the New30Days SQL column name for the IssueSummary table
const IssueSummaryColumnNew30Days = "new30_days"

// IssueSummaryEscapedColumnNew30Days is the escaped New30Days SQL column name for the IssueSummary table
const IssueSummaryEscapedColumnNew30Days = "`new30_days`"

// IssueSummaryColumnTotalClosed is the TotalClosed SQL column name for the IssueSummary table
const IssueSummaryColumnTotalClosed = "total_closed"

// IssueSummaryEscapedColumnTotalClosed is the escaped TotalClosed SQL column name for the IssueSummary table
const IssueSummaryEscapedColumnTotalClosed = "`total_closed`"

// IssueSummaryColumnClosed30Days is the Closed30Days SQL column name for the IssueSummary table
const IssueSummaryColumnClosed30Days = "closed30_days"

// IssueSummaryEscapedColumnClosed30Days is the escaped Closed30Days SQL column name for the IssueSummary table
const IssueSummaryEscapedColumnClosed30Days = "`closed30_days`"

// IssueSummaryColumnEstimatedWorkMonths is the EstimatedWorkMonths SQL column name for the IssueSummary table
const IssueSummaryColumnEstimatedWorkMonths = "estimated_work_months"

// IssueSummaryEscapedColumnEstimatedWorkMonths is the escaped EstimatedWorkMonths SQL column name for the IssueSummary table
const IssueSummaryEscapedColumnEstimatedWorkMonths = "`estimated_work_months`"

// IssueSummaryColumnEstimatedWorkMonths30Days is the EstimatedWorkMonths30Days SQL column name for the IssueSummary table
const IssueSummaryColumnEstimatedWorkMonths30Days = "estimated_work_months30_days"

// IssueSummaryEscapedColumnEstimatedWorkMonths30Days is the escaped EstimatedWorkMonths30Days SQL column name for the IssueSummary table
const IssueSummaryEscapedColumnEstimatedWorkMonths30Days = "`estimated_work_months30_days`"

// IssueSummaryColumnTitle is the Title SQL column name for the IssueSummary table
const IssueSummaryColumnTitle = "title"

// IssueSummaryEscapedColumnTitle is the escaped Title SQL column name for the IssueSummary table
const IssueSummaryEscapedColumnTitle = "`title`"

// IssueSummaryColumnURL is the URL SQL column name for the IssueSummary table
const IssueSummaryColumnURL = "url"

// IssueSummaryEscapedColumnURL is the escaped URL SQL column name for the IssueSummary table
const IssueSummaryEscapedColumnURL = "`url`"

// IssueSummaryColumnPriority is the Priority SQL column name for the IssueSummary table
const IssueSummaryColumnPriority = "priority"

// IssueSummaryEscapedColumnPriority is the escaped Priority SQL column name for the IssueSummary table
const IssueSummaryEscapedColumnPriority = "`priority`"

// IssueSummaryColumnPriorityID is the PriorityID SQL column name for the IssueSummary table
const IssueSummaryColumnPriorityID = "priority_id"

// IssueSummaryEscapedColumnPriorityID is the escaped PriorityID SQL column name for the IssueSummary table
const IssueSummaryEscapedColumnPriorityID = "`priority_id`"

// IssueSummaryColumnStatus is the Status SQL column name for the IssueSummary table
const IssueSummaryColumnStatus = "status"

// IssueSummaryEscapedColumnStatus is the escaped Status SQL column name for the IssueSummary table
const IssueSummaryEscapedColumnStatus = "`status`"

// IssueSummaryColumnStatusID is the StatusID SQL column name for the IssueSummary table
const IssueSummaryColumnStatusID = "status_id"

// IssueSummaryEscapedColumnStatusID is the escaped StatusID SQL column name for the IssueSummary table
const IssueSummaryEscapedColumnStatusID = "`status_id`"

// IssueSummaryColumnIssueType is the IssueType SQL column name for the IssueSummary table
const IssueSummaryColumnIssueType = "issue_type"

// IssueSummaryEscapedColumnIssueType is the escaped IssueType SQL column name for the IssueSummary table
const IssueSummaryEscapedColumnIssueType = "`issue_type`"

// IssueSummaryColumnIssueTypeID is the IssueTypeID SQL column name for the IssueSummary table
const IssueSummaryColumnIssueTypeID = "issue_type_id"

// IssueSummaryEscapedColumnIssueTypeID is the escaped IssueTypeID SQL column name for the IssueSummary table
const IssueSummaryEscapedColumnIssueTypeID = "`issue_type_id`"

// IssueSummaryColumnResolution is the Resolution SQL column name for the IssueSummary table
const IssueSummaryColumnResolution = "resolution"

// IssueSummaryEscapedColumnResolution is the escaped Resolution SQL column name for the IssueSummary table
const IssueSummaryEscapedColumnResolution = "`resolution`"

// IssueSummaryColumnResolutionID is the ResolutionID SQL column name for the IssueSummary table
const IssueSummaryColumnResolutionID = "resolution_id"

// IssueSummaryEscapedColumnResolutionID is the escaped ResolutionID SQL column name for the IssueSummary table
const IssueSummaryEscapedColumnResolutionID = "`resolution_id`"

// IssueSummaryColumnState is the State SQL column name for the IssueSummary table
const IssueSummaryColumnState = "state"

// IssueSummaryEscapedColumnState is the escaped State SQL column name for the IssueSummary table
const IssueSummaryEscapedColumnState = "`state`"

// IssueSummaryColumnCustomFieldIds is the CustomFieldIds SQL column name for the IssueSummary table
const IssueSummaryColumnCustomFieldIds = "custom_field_ids"

// IssueSummaryEscapedColumnCustomFieldIds is the escaped CustomFieldIds SQL column name for the IssueSummary table
const IssueSummaryEscapedColumnCustomFieldIds = "`custom_field_ids`"

// IssueSummaryColumnTeams is the Teams SQL column name for the IssueSummary table
const IssueSummaryColumnTeams = "teams"

// IssueSummaryEscapedColumnTeams is the escaped Teams SQL column name for the IssueSummary table
const IssueSummaryEscapedColumnTeams = "`teams`"

// IssueSummaryColumnParentIssueID is the ParentIssueID SQL column name for the IssueSummary table
const IssueSummaryColumnParentIssueID = "parent_issue_id"

// IssueSummaryEscapedColumnParentIssueID is the escaped ParentIssueID SQL column name for the IssueSummary table
const IssueSummaryEscapedColumnParentIssueID = "`parent_issue_id`"

// IssueSummaryColumnParentsIssueIds is the ParentsIssueIds SQL column name for the IssueSummary table
const IssueSummaryColumnParentsIssueIds = "parents_issue_ids"

// IssueSummaryEscapedColumnParentsIssueIds is the escaped ParentsIssueIds SQL column name for the IssueSummary table
const IssueSummaryEscapedColumnParentsIssueIds = "`parents_issue_ids`"

// IssueSummaryColumnMetadata is the Metadata SQL column name for the IssueSummary table
const IssueSummaryColumnMetadata = "metadata"

// IssueSummaryEscapedColumnMetadata is the escaped Metadata SQL column name for the IssueSummary table
const IssueSummaryEscapedColumnMetadata = "`metadata`"

// IssueSummaryColumnProjectID is the ProjectID SQL column name for the IssueSummary table
const IssueSummaryColumnProjectID = "project_id"

// IssueSummaryEscapedColumnProjectID is the escaped ProjectID SQL column name for the IssueSummary table
const IssueSummaryEscapedColumnProjectID = "`project_id`"

// IssueSummaryColumnSprints is the Sprints SQL column name for the IssueSummary table
const IssueSummaryColumnSprints = "sprints"

// IssueSummaryEscapedColumnSprints is the escaped Sprints SQL column name for the IssueSummary table
const IssueSummaryEscapedColumnSprints = "`sprints`"

// IssueSummaryColumnLabels is the Labels SQL column name for the IssueSummary table
const IssueSummaryColumnLabels = "labels"

// IssueSummaryEscapedColumnLabels is the escaped Labels SQL column name for the IssueSummary table
const IssueSummaryEscapedColumnLabels = "`labels`"

// IssueSummaryColumnTopLevel is the TopLevel SQL column name for the IssueSummary table
const IssueSummaryColumnTopLevel = "top_level"

// IssueSummaryEscapedColumnTopLevel is the escaped TopLevel SQL column name for the IssueSummary table
const IssueSummaryEscapedColumnTopLevel = "`top_level`"

// IssueSummaryColumnIsLeaf is the IsLeaf SQL column name for the IssueSummary table
const IssueSummaryColumnIsLeaf = "is_leaf"

// IssueSummaryEscapedColumnIsLeaf is the escaped IsLeaf SQL column name for the IssueSummary table
const IssueSummaryEscapedColumnIsLeaf = "`is_leaf`"

// IssueSummaryColumnPath is the Path SQL column name for the IssueSummary table
const IssueSummaryColumnPath = "path"

// IssueSummaryEscapedColumnPath is the escaped Path SQL column name for the IssueSummary table
const IssueSummaryEscapedColumnPath = "`path`"

// IssueSummaryColumnInProgressDuration is the InProgressDuration SQL column name for the IssueSummary table
const IssueSummaryColumnInProgressDuration = "in_progress_duration"

// IssueSummaryEscapedColumnInProgressDuration is the escaped InProgressDuration SQL column name for the IssueSummary table
const IssueSummaryEscapedColumnInProgressDuration = "`in_progress_duration`"

// IssueSummaryColumnVerificationDuration is the VerificationDuration SQL column name for the IssueSummary table
const IssueSummaryColumnVerificationDuration = "verification_duration"

// IssueSummaryEscapedColumnVerificationDuration is the escaped VerificationDuration SQL column name for the IssueSummary table
const IssueSummaryEscapedColumnVerificationDuration = "`verification_duration`"

// IssueSummaryColumnInProgressCount is the InProgressCount SQL column name for the IssueSummary table
const IssueSummaryColumnInProgressCount = "in_progress_count"

// IssueSummaryEscapedColumnInProgressCount is the escaped InProgressCount SQL column name for the IssueSummary table
const IssueSummaryEscapedColumnInProgressCount = "`in_progress_count`"

// IssueSummaryColumnReopenCount is the ReopenCount SQL column name for the IssueSummary table
const IssueSummaryColumnReopenCount = "reopen_count"

// IssueSummaryEscapedColumnReopenCount is the escaped ReopenCount SQL column name for the IssueSummary table
const IssueSummaryEscapedColumnReopenCount = "`reopen_count`"

// IssueSummaryColumnMappedType is the MappedType SQL column name for the IssueSummary table
const IssueSummaryColumnMappedType = "mapped_type"

// IssueSummaryEscapedColumnMappedType is the escaped MappedType SQL column name for the IssueSummary table
const IssueSummaryEscapedColumnMappedType = "`mapped_type`"

// IssueSummaryColumnStrategicParentID is the StrategicParentID SQL column name for the IssueSummary table
const IssueSummaryColumnStrategicParentID = "strategic_parent_id"

// IssueSummaryEscapedColumnStrategicParentID is the escaped StrategicParentID SQL column name for the IssueSummary table
const IssueSummaryEscapedColumnStrategicParentID = "`strategic_parent_id`"

// IssueSummaryColumnSprintID is the SprintID SQL column name for the IssueSummary table
const IssueSummaryColumnSprintID = "sprint_id"

// IssueSummaryEscapedColumnSprintID is the escaped SprintID SQL column name for the IssueSummary table
const IssueSummaryEscapedColumnSprintID = "`sprint_id`"

// IssueSummaryColumnIssueProjectName is the IssueProjectName SQL column name for the IssueSummary table
const IssueSummaryColumnIssueProjectName = "issue_project_name"

// IssueSummaryEscapedColumnIssueProjectName is the escaped IssueProjectName SQL column name for the IssueSummary table
const IssueSummaryEscapedColumnIssueProjectName = "`issue_project_name`"

// IssueSummaryColumnUsers is the Users SQL column name for the IssueSummary table
const IssueSummaryColumnUsers = "users"

// IssueSummaryEscapedColumnUsers is the escaped Users SQL column name for the IssueSummary table
const IssueSummaryEscapedColumnUsers = "`users`"

// IssueSummaryColumnInitialStartDate is the InitialStartDate SQL column name for the IssueSummary table
const IssueSummaryColumnInitialStartDate = "initial_start_date"

// IssueSummaryEscapedColumnInitialStartDate is the escaped InitialStartDate SQL column name for the IssueSummary table
const IssueSummaryEscapedColumnInitialStartDate = "`initial_start_date`"

// IssueSummaryColumnTotalDuration is the TotalDuration SQL column name for the IssueSummary table
const IssueSummaryColumnTotalDuration = "total_duration"

// IssueSummaryEscapedColumnTotalDuration is the escaped TotalDuration SQL column name for the IssueSummary table
const IssueSummaryEscapedColumnTotalDuration = "`total_duration`"

// IssueSummaryColumnCreatedAt is the CreatedAt SQL column name for the IssueSummary table
const IssueSummaryColumnCreatedAt = "created_at"

// IssueSummaryEscapedColumnCreatedAt is the escaped CreatedAt SQL column name for the IssueSummary table
const IssueSummaryEscapedColumnCreatedAt = "`created_at`"

// IssueSummaryColumnClosedAt is the ClosedAt SQL column name for the IssueSummary table
const IssueSummaryColumnClosedAt = "closed_at"

// IssueSummaryEscapedColumnClosedAt is the escaped ClosedAt SQL column name for the IssueSummary table
const IssueSummaryEscapedColumnClosedAt = "`closed_at`"

// IssueSummaryColumnPlannedEndDate is the PlannedEndDate SQL column name for the IssueSummary table
const IssueSummaryColumnPlannedEndDate = "planned_end_date"

// IssueSummaryEscapedColumnPlannedEndDate is the escaped PlannedEndDate SQL column name for the IssueSummary table
const IssueSummaryEscapedColumnPlannedEndDate = "`planned_end_date`"

// IssueSummaryColumnCustomerID is the CustomerID SQL column name for the IssueSummary table
const IssueSummaryColumnCustomerID = "customer_id"

// IssueSummaryEscapedColumnCustomerID is the escaped CustomerID SQL column name for the IssueSummary table
const IssueSummaryEscapedColumnCustomerID = "`customer_id`"

// IssueSummaryColumnRefType is the RefType SQL column name for the IssueSummary table
const IssueSummaryColumnRefType = "ref_type"

// IssueSummaryEscapedColumnRefType is the escaped RefType SQL column name for the IssueSummary table
const IssueSummaryEscapedColumnRefType = "`ref_type`"

// IssueSummaryColumnRefID is the RefID SQL column name for the IssueSummary table
const IssueSummaryColumnRefID = "ref_id"

// IssueSummaryEscapedColumnRefID is the escaped RefID SQL column name for the IssueSummary table
const IssueSummaryEscapedColumnRefID = "`ref_id`"

// IssueSummaryColumnCustomFieldIdsVirtual is the CustomFieldIdsVirtual SQL column name for the IssueSummary table
const IssueSummaryColumnCustomFieldIdsVirtual = "custom_field_ids_virtual"

// IssueSummaryEscapedColumnCustomFieldIdsVirtual is the escaped CustomFieldIdsVirtual SQL column name for the IssueSummary table
const IssueSummaryEscapedColumnCustomFieldIdsVirtual = "`custom_field_ids_virtual`"

// IssueSummaryColumnReleaseDuration is the ReleaseDuration SQL column name for the IssueSummary table
const IssueSummaryColumnReleaseDuration = "release_duration"

// IssueSummaryEscapedColumnReleaseDuration is the escaped ReleaseDuration SQL column name for the IssueSummary table
const IssueSummaryEscapedColumnReleaseDuration = "`release_duration`"

// IssueSummaryColumnCompleted is the Completed SQL column name for the IssueSummary table
const IssueSummaryColumnCompleted = "completed"

// IssueSummaryEscapedColumnCompleted is the escaped Completed SQL column name for the IssueSummary table
const IssueSummaryEscapedColumnCompleted = "`completed`"

// IssueSummaryColumnCompletedDate is the CompletedDate SQL column name for the IssueSummary table
const IssueSummaryColumnCompletedDate = "completed_date"

// IssueSummaryEscapedColumnCompletedDate is the escaped CompletedDate SQL column name for the IssueSummary table
const IssueSummaryEscapedColumnCompletedDate = "`completed_date`"

// GetID will return the IssueSummary ID value
func (t *IssueSummary) GetID() string {
	return t.ID
}

// SetID will set the IssueSummary ID value
func (t *IssueSummary) SetID(v string) {
	t.ID = v
}

// FindIssueSummaryByID will find a IssueSummary by ID
func FindIssueSummaryByID(ctx context.Context, db DB, value string) (*IssueSummary, error) {
	q := "SELECT `issue_summary`.`id`,`issue_summary`.`checksum`,`issue_summary`.`issue_id`,`issue_summary`.`total_issues`,`issue_summary`.`new30_days`,`issue_summary`.`total_closed`,`issue_summary`.`closed30_days`,`issue_summary`.`estimated_work_months`,`issue_summary`.`estimated_work_months30_days`,`issue_summary`.`title`,`issue_summary`.`url`,`issue_summary`.`priority`,`issue_summary`.`priority_id`,`issue_summary`.`status`,`issue_summary`.`status_id`,`issue_summary`.`issue_type`,`issue_summary`.`issue_type_id`,`issue_summary`.`resolution`,`issue_summary`.`resolution_id`,`issue_summary`.`state`,`issue_summary`.`custom_field_ids`,`issue_summary`.`teams`,`issue_summary`.`parent_issue_id`,`issue_summary`.`parents_issue_ids`,`issue_summary`.`metadata`,`issue_summary`.`project_id`,`issue_summary`.`sprints`,`issue_summary`.`labels`,`issue_summary`.`top_level`,`issue_summary`.`is_leaf`,`issue_summary`.`path`,`issue_summary`.`in_progress_duration`,`issue_summary`.`verification_duration`,`issue_summary`.`in_progress_count`,`issue_summary`.`reopen_count`,`issue_summary`.`mapped_type`,`issue_summary`.`strategic_parent_id`,`issue_summary`.`sprint_id`,`issue_summary`.`issue_project_name`,`issue_summary`.`users`,`issue_summary`.`initial_start_date`,`issue_summary`.`total_duration`,`issue_summary`.`created_at`,`issue_summary`.`closed_at`,`issue_summary`.`planned_end_date`,`issue_summary`.`customer_id`,`issue_summary`.`ref_type`,`issue_summary`.`ref_id`,`issue_summary`.`custom_field_ids_virtual`,`issue_summary`.`release_duration`,`issue_summary`.`completed`,`issue_summary`.`completed_date` FROM `issue_summary` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _IssueID sql.NullString
	var _TotalIssues sql.NullInt64
	var _New30Days sql.NullInt64
	var _TotalClosed sql.NullInt64
	var _Closed30Days sql.NullInt64
	var _EstimatedWorkMonths sql.NullFloat64
	var _EstimatedWorkMonths30Days sql.NullFloat64
	var _Title sql.NullString
	var _URL sql.NullString
	var _Priority sql.NullString
	var _PriorityID sql.NullString
	var _Status sql.NullString
	var _StatusID sql.NullString
	var _IssueType sql.NullString
	var _IssueTypeID sql.NullString
	var _Resolution sql.NullString
	var _ResolutionID sql.NullString
	var _State sql.NullString
	var _CustomFieldIds sql.NullString
	var _Teams sql.NullString
	var _ParentIssueID sql.NullString
	var _ParentsIssueIds sql.NullString
	var _Metadata sql.NullString
	var _ProjectID sql.NullString
	var _Sprints sql.NullString
	var _Labels sql.NullString
	var _TopLevel sql.NullBool
	var _IsLeaf sql.NullBool
	var _Path sql.NullString
	var _InProgressDuration sql.NullInt64
	var _VerificationDuration sql.NullInt64
	var _InProgressCount sql.NullInt64
	var _ReopenCount sql.NullInt64
	var _MappedType sql.NullString
	var _StrategicParentID sql.NullString
	var _SprintID sql.NullString
	var _IssueProjectName sql.NullString
	var _Users sql.NullString
	var _InitialStartDate sql.NullInt64
	var _TotalDuration sql.NullInt64
	var _CreatedAt sql.NullInt64
	var _ClosedAt sql.NullInt64
	var _PlannedEndDate sql.NullInt64
	var _CustomerID sql.NullString
	var _RefType sql.NullString
	var _RefID sql.NullString
	var _CustomFieldIdsVirtual sql.NullString
	var _ReleaseDuration sql.NullInt64
	var _Completed sql.NullBool
	var _CompletedDate sql.NullInt64
	err := db.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_IssueID,
		&_TotalIssues,
		&_New30Days,
		&_TotalClosed,
		&_Closed30Days,
		&_EstimatedWorkMonths,
		&_EstimatedWorkMonths30Days,
		&_Title,
		&_URL,
		&_Priority,
		&_PriorityID,
		&_Status,
		&_StatusID,
		&_IssueType,
		&_IssueTypeID,
		&_Resolution,
		&_ResolutionID,
		&_State,
		&_CustomFieldIds,
		&_Teams,
		&_ParentIssueID,
		&_ParentsIssueIds,
		&_Metadata,
		&_ProjectID,
		&_Sprints,
		&_Labels,
		&_TopLevel,
		&_IsLeaf,
		&_Path,
		&_InProgressDuration,
		&_VerificationDuration,
		&_InProgressCount,
		&_ReopenCount,
		&_MappedType,
		&_StrategicParentID,
		&_SprintID,
		&_IssueProjectName,
		&_Users,
		&_InitialStartDate,
		&_TotalDuration,
		&_CreatedAt,
		&_ClosedAt,
		&_PlannedEndDate,
		&_CustomerID,
		&_RefType,
		&_RefID,
		&_CustomFieldIdsVirtual,
		&_ReleaseDuration,
		&_Completed,
		&_CompletedDate,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &IssueSummary{}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _Checksum.Valid {
		t.SetChecksum(_Checksum.String)
	}
	if _IssueID.Valid {
		t.SetIssueID(_IssueID.String)
	}
	if _TotalIssues.Valid {
		t.SetTotalIssues(int32(_TotalIssues.Int64))
	}
	if _New30Days.Valid {
		t.SetNew30Days(int32(_New30Days.Int64))
	}
	if _TotalClosed.Valid {
		t.SetTotalClosed(int32(_TotalClosed.Int64))
	}
	if _Closed30Days.Valid {
		t.SetClosed30Days(int32(_Closed30Days.Int64))
	}
	if _EstimatedWorkMonths.Valid {
		t.SetEstimatedWorkMonths(_EstimatedWorkMonths.Float64)
	}
	if _EstimatedWorkMonths30Days.Valid {
		t.SetEstimatedWorkMonths30Days(_EstimatedWorkMonths30Days.Float64)
	}
	if _Title.Valid {
		t.SetTitle(_Title.String)
	}
	if _URL.Valid {
		t.SetURL(_URL.String)
	}
	if _Priority.Valid {
		t.SetPriority(_Priority.String)
	}
	if _PriorityID.Valid {
		t.SetPriorityID(_PriorityID.String)
	}
	if _Status.Valid {
		t.SetStatus(_Status.String)
	}
	if _StatusID.Valid {
		t.SetStatusID(_StatusID.String)
	}
	if _IssueType.Valid {
		t.SetIssueType(_IssueType.String)
	}
	if _IssueTypeID.Valid {
		t.SetIssueTypeID(_IssueTypeID.String)
	}
	if _Resolution.Valid {
		t.SetResolution(_Resolution.String)
	}
	if _ResolutionID.Valid {
		t.SetResolutionID(_ResolutionID.String)
	}
	if _State.Valid {
		t.SetState(_State.String)
	}
	if _CustomFieldIds.Valid {
		t.SetCustomFieldIds(_CustomFieldIds.String)
	}
	if _Teams.Valid {
		t.SetTeams(_Teams.String)
	}
	if _ParentIssueID.Valid {
		t.SetParentIssueID(_ParentIssueID.String)
	}
	if _ParentsIssueIds.Valid {
		t.SetParentsIssueIds(_ParentsIssueIds.String)
	}
	if _Metadata.Valid {
		t.SetMetadata(_Metadata.String)
	}
	if _ProjectID.Valid {
		t.SetProjectID(_ProjectID.String)
	}
	if _Sprints.Valid {
		t.SetSprints(_Sprints.String)
	}
	if _Labels.Valid {
		t.SetLabels(_Labels.String)
	}
	if _TopLevel.Valid {
		t.SetTopLevel(_TopLevel.Bool)
	}
	if _IsLeaf.Valid {
		t.SetIsLeaf(_IsLeaf.Bool)
	}
	if _Path.Valid {
		t.SetPath(_Path.String)
	}
	if _InProgressDuration.Valid {
		t.SetInProgressDuration(_InProgressDuration.Int64)
	}
	if _VerificationDuration.Valid {
		t.SetVerificationDuration(_VerificationDuration.Int64)
	}
	if _InProgressCount.Valid {
		t.SetInProgressCount(int32(_InProgressCount.Int64))
	}
	if _ReopenCount.Valid {
		t.SetReopenCount(int32(_ReopenCount.Int64))
	}
	if _MappedType.Valid {
		t.SetMappedType(_MappedType.String)
	}
	if _StrategicParentID.Valid {
		t.SetStrategicParentID(_StrategicParentID.String)
	}
	if _SprintID.Valid {
		t.SetSprintID(_SprintID.String)
	}
	if _IssueProjectName.Valid {
		t.SetIssueProjectName(_IssueProjectName.String)
	}
	if _Users.Valid {
		t.SetUsers(_Users.String)
	}
	if _InitialStartDate.Valid {
		t.SetInitialStartDate(_InitialStartDate.Int64)
	}
	if _TotalDuration.Valid {
		t.SetTotalDuration(_TotalDuration.Int64)
	}
	if _CreatedAt.Valid {
		t.SetCreatedAt(_CreatedAt.Int64)
	}
	if _ClosedAt.Valid {
		t.SetClosedAt(_ClosedAt.Int64)
	}
	if _PlannedEndDate.Valid {
		t.SetPlannedEndDate(_PlannedEndDate.Int64)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _RefType.Valid {
		t.SetRefType(_RefType.String)
	}
	if _RefID.Valid {
		t.SetRefID(_RefID.String)
	}
	if _CustomFieldIdsVirtual.Valid {
		t.SetCustomFieldIdsVirtual(_CustomFieldIdsVirtual.String)
	}
	if _ReleaseDuration.Valid {
		t.SetReleaseDuration(_ReleaseDuration.Int64)
	}
	if _Completed.Valid {
		t.SetCompleted(_Completed.Bool)
	}
	if _CompletedDate.Valid {
		t.SetCompletedDate(_CompletedDate.Int64)
	}
	return t, nil
}

// FindIssueSummaryByIDTx will find a IssueSummary by ID using the provided transaction
func FindIssueSummaryByIDTx(ctx context.Context, tx Tx, value string) (*IssueSummary, error) {
	q := "SELECT `issue_summary`.`id`,`issue_summary`.`checksum`,`issue_summary`.`issue_id`,`issue_summary`.`total_issues`,`issue_summary`.`new30_days`,`issue_summary`.`total_closed`,`issue_summary`.`closed30_days`,`issue_summary`.`estimated_work_months`,`issue_summary`.`estimated_work_months30_days`,`issue_summary`.`title`,`issue_summary`.`url`,`issue_summary`.`priority`,`issue_summary`.`priority_id`,`issue_summary`.`status`,`issue_summary`.`status_id`,`issue_summary`.`issue_type`,`issue_summary`.`issue_type_id`,`issue_summary`.`resolution`,`issue_summary`.`resolution_id`,`issue_summary`.`state`,`issue_summary`.`custom_field_ids`,`issue_summary`.`teams`,`issue_summary`.`parent_issue_id`,`issue_summary`.`parents_issue_ids`,`issue_summary`.`metadata`,`issue_summary`.`project_id`,`issue_summary`.`sprints`,`issue_summary`.`labels`,`issue_summary`.`top_level`,`issue_summary`.`is_leaf`,`issue_summary`.`path`,`issue_summary`.`in_progress_duration`,`issue_summary`.`verification_duration`,`issue_summary`.`in_progress_count`,`issue_summary`.`reopen_count`,`issue_summary`.`mapped_type`,`issue_summary`.`strategic_parent_id`,`issue_summary`.`sprint_id`,`issue_summary`.`issue_project_name`,`issue_summary`.`users`,`issue_summary`.`initial_start_date`,`issue_summary`.`total_duration`,`issue_summary`.`created_at`,`issue_summary`.`closed_at`,`issue_summary`.`planned_end_date`,`issue_summary`.`customer_id`,`issue_summary`.`ref_type`,`issue_summary`.`ref_id`,`issue_summary`.`custom_field_ids_virtual`,`issue_summary`.`release_duration`,`issue_summary`.`completed`,`issue_summary`.`completed_date` FROM `issue_summary` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _IssueID sql.NullString
	var _TotalIssues sql.NullInt64
	var _New30Days sql.NullInt64
	var _TotalClosed sql.NullInt64
	var _Closed30Days sql.NullInt64
	var _EstimatedWorkMonths sql.NullFloat64
	var _EstimatedWorkMonths30Days sql.NullFloat64
	var _Title sql.NullString
	var _URL sql.NullString
	var _Priority sql.NullString
	var _PriorityID sql.NullString
	var _Status sql.NullString
	var _StatusID sql.NullString
	var _IssueType sql.NullString
	var _IssueTypeID sql.NullString
	var _Resolution sql.NullString
	var _ResolutionID sql.NullString
	var _State sql.NullString
	var _CustomFieldIds sql.NullString
	var _Teams sql.NullString
	var _ParentIssueID sql.NullString
	var _ParentsIssueIds sql.NullString
	var _Metadata sql.NullString
	var _ProjectID sql.NullString
	var _Sprints sql.NullString
	var _Labels sql.NullString
	var _TopLevel sql.NullBool
	var _IsLeaf sql.NullBool
	var _Path sql.NullString
	var _InProgressDuration sql.NullInt64
	var _VerificationDuration sql.NullInt64
	var _InProgressCount sql.NullInt64
	var _ReopenCount sql.NullInt64
	var _MappedType sql.NullString
	var _StrategicParentID sql.NullString
	var _SprintID sql.NullString
	var _IssueProjectName sql.NullString
	var _Users sql.NullString
	var _InitialStartDate sql.NullInt64
	var _TotalDuration sql.NullInt64
	var _CreatedAt sql.NullInt64
	var _ClosedAt sql.NullInt64
	var _PlannedEndDate sql.NullInt64
	var _CustomerID sql.NullString
	var _RefType sql.NullString
	var _RefID sql.NullString
	var _CustomFieldIdsVirtual sql.NullString
	var _ReleaseDuration sql.NullInt64
	var _Completed sql.NullBool
	var _CompletedDate sql.NullInt64
	err := tx.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_IssueID,
		&_TotalIssues,
		&_New30Days,
		&_TotalClosed,
		&_Closed30Days,
		&_EstimatedWorkMonths,
		&_EstimatedWorkMonths30Days,
		&_Title,
		&_URL,
		&_Priority,
		&_PriorityID,
		&_Status,
		&_StatusID,
		&_IssueType,
		&_IssueTypeID,
		&_Resolution,
		&_ResolutionID,
		&_State,
		&_CustomFieldIds,
		&_Teams,
		&_ParentIssueID,
		&_ParentsIssueIds,
		&_Metadata,
		&_ProjectID,
		&_Sprints,
		&_Labels,
		&_TopLevel,
		&_IsLeaf,
		&_Path,
		&_InProgressDuration,
		&_VerificationDuration,
		&_InProgressCount,
		&_ReopenCount,
		&_MappedType,
		&_StrategicParentID,
		&_SprintID,
		&_IssueProjectName,
		&_Users,
		&_InitialStartDate,
		&_TotalDuration,
		&_CreatedAt,
		&_ClosedAt,
		&_PlannedEndDate,
		&_CustomerID,
		&_RefType,
		&_RefID,
		&_CustomFieldIdsVirtual,
		&_ReleaseDuration,
		&_Completed,
		&_CompletedDate,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &IssueSummary{}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _Checksum.Valid {
		t.SetChecksum(_Checksum.String)
	}
	if _IssueID.Valid {
		t.SetIssueID(_IssueID.String)
	}
	if _TotalIssues.Valid {
		t.SetTotalIssues(int32(_TotalIssues.Int64))
	}
	if _New30Days.Valid {
		t.SetNew30Days(int32(_New30Days.Int64))
	}
	if _TotalClosed.Valid {
		t.SetTotalClosed(int32(_TotalClosed.Int64))
	}
	if _Closed30Days.Valid {
		t.SetClosed30Days(int32(_Closed30Days.Int64))
	}
	if _EstimatedWorkMonths.Valid {
		t.SetEstimatedWorkMonths(_EstimatedWorkMonths.Float64)
	}
	if _EstimatedWorkMonths30Days.Valid {
		t.SetEstimatedWorkMonths30Days(_EstimatedWorkMonths30Days.Float64)
	}
	if _Title.Valid {
		t.SetTitle(_Title.String)
	}
	if _URL.Valid {
		t.SetURL(_URL.String)
	}
	if _Priority.Valid {
		t.SetPriority(_Priority.String)
	}
	if _PriorityID.Valid {
		t.SetPriorityID(_PriorityID.String)
	}
	if _Status.Valid {
		t.SetStatus(_Status.String)
	}
	if _StatusID.Valid {
		t.SetStatusID(_StatusID.String)
	}
	if _IssueType.Valid {
		t.SetIssueType(_IssueType.String)
	}
	if _IssueTypeID.Valid {
		t.SetIssueTypeID(_IssueTypeID.String)
	}
	if _Resolution.Valid {
		t.SetResolution(_Resolution.String)
	}
	if _ResolutionID.Valid {
		t.SetResolutionID(_ResolutionID.String)
	}
	if _State.Valid {
		t.SetState(_State.String)
	}
	if _CustomFieldIds.Valid {
		t.SetCustomFieldIds(_CustomFieldIds.String)
	}
	if _Teams.Valid {
		t.SetTeams(_Teams.String)
	}
	if _ParentIssueID.Valid {
		t.SetParentIssueID(_ParentIssueID.String)
	}
	if _ParentsIssueIds.Valid {
		t.SetParentsIssueIds(_ParentsIssueIds.String)
	}
	if _Metadata.Valid {
		t.SetMetadata(_Metadata.String)
	}
	if _ProjectID.Valid {
		t.SetProjectID(_ProjectID.String)
	}
	if _Sprints.Valid {
		t.SetSprints(_Sprints.String)
	}
	if _Labels.Valid {
		t.SetLabels(_Labels.String)
	}
	if _TopLevel.Valid {
		t.SetTopLevel(_TopLevel.Bool)
	}
	if _IsLeaf.Valid {
		t.SetIsLeaf(_IsLeaf.Bool)
	}
	if _Path.Valid {
		t.SetPath(_Path.String)
	}
	if _InProgressDuration.Valid {
		t.SetInProgressDuration(_InProgressDuration.Int64)
	}
	if _VerificationDuration.Valid {
		t.SetVerificationDuration(_VerificationDuration.Int64)
	}
	if _InProgressCount.Valid {
		t.SetInProgressCount(int32(_InProgressCount.Int64))
	}
	if _ReopenCount.Valid {
		t.SetReopenCount(int32(_ReopenCount.Int64))
	}
	if _MappedType.Valid {
		t.SetMappedType(_MappedType.String)
	}
	if _StrategicParentID.Valid {
		t.SetStrategicParentID(_StrategicParentID.String)
	}
	if _SprintID.Valid {
		t.SetSprintID(_SprintID.String)
	}
	if _IssueProjectName.Valid {
		t.SetIssueProjectName(_IssueProjectName.String)
	}
	if _Users.Valid {
		t.SetUsers(_Users.String)
	}
	if _InitialStartDate.Valid {
		t.SetInitialStartDate(_InitialStartDate.Int64)
	}
	if _TotalDuration.Valid {
		t.SetTotalDuration(_TotalDuration.Int64)
	}
	if _CreatedAt.Valid {
		t.SetCreatedAt(_CreatedAt.Int64)
	}
	if _ClosedAt.Valid {
		t.SetClosedAt(_ClosedAt.Int64)
	}
	if _PlannedEndDate.Valid {
		t.SetPlannedEndDate(_PlannedEndDate.Int64)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _RefType.Valid {
		t.SetRefType(_RefType.String)
	}
	if _RefID.Valid {
		t.SetRefID(_RefID.String)
	}
	if _CustomFieldIdsVirtual.Valid {
		t.SetCustomFieldIdsVirtual(_CustomFieldIdsVirtual.String)
	}
	if _ReleaseDuration.Valid {
		t.SetReleaseDuration(_ReleaseDuration.Int64)
	}
	if _Completed.Valid {
		t.SetCompleted(_Completed.Bool)
	}
	if _CompletedDate.Valid {
		t.SetCompletedDate(_CompletedDate.Int64)
	}
	return t, nil
}

// GetChecksum will return the IssueSummary Checksum value
func (t *IssueSummary) GetChecksum() string {
	if t.Checksum == nil {
		return ""
	}
	return *t.Checksum
}

// SetChecksum will set the IssueSummary Checksum value
func (t *IssueSummary) SetChecksum(v string) {
	t.Checksum = &v
}

// GetIssueID will return the IssueSummary IssueID value
func (t *IssueSummary) GetIssueID() string {
	return t.IssueID
}

// SetIssueID will set the IssueSummary IssueID value
func (t *IssueSummary) SetIssueID(v string) {
	t.IssueID = v
}

// FindIssueSummariesByIssueID will find all IssueSummarys by the IssueID value
func FindIssueSummariesByIssueID(ctx context.Context, db DB, value string) ([]*IssueSummary, error) {
	q := "SELECT `issue_summary`.`id`,`issue_summary`.`checksum`,`issue_summary`.`issue_id`,`issue_summary`.`total_issues`,`issue_summary`.`new30_days`,`issue_summary`.`total_closed`,`issue_summary`.`closed30_days`,`issue_summary`.`estimated_work_months`,`issue_summary`.`estimated_work_months30_days`,`issue_summary`.`title`,`issue_summary`.`url`,`issue_summary`.`priority`,`issue_summary`.`priority_id`,`issue_summary`.`status`,`issue_summary`.`status_id`,`issue_summary`.`issue_type`,`issue_summary`.`issue_type_id`,`issue_summary`.`resolution`,`issue_summary`.`resolution_id`,`issue_summary`.`state`,`issue_summary`.`custom_field_ids`,`issue_summary`.`teams`,`issue_summary`.`parent_issue_id`,`issue_summary`.`parents_issue_ids`,`issue_summary`.`metadata`,`issue_summary`.`project_id`,`issue_summary`.`sprints`,`issue_summary`.`labels`,`issue_summary`.`top_level`,`issue_summary`.`is_leaf`,`issue_summary`.`path`,`issue_summary`.`in_progress_duration`,`issue_summary`.`verification_duration`,`issue_summary`.`in_progress_count`,`issue_summary`.`reopen_count`,`issue_summary`.`mapped_type`,`issue_summary`.`strategic_parent_id`,`issue_summary`.`sprint_id`,`issue_summary`.`issue_project_name`,`issue_summary`.`users`,`issue_summary`.`initial_start_date`,`issue_summary`.`total_duration`,`issue_summary`.`created_at`,`issue_summary`.`closed_at`,`issue_summary`.`planned_end_date`,`issue_summary`.`customer_id`,`issue_summary`.`ref_type`,`issue_summary`.`ref_id`,`issue_summary`.`custom_field_ids_virtual`,`issue_summary`.`release_duration`,`issue_summary`.`completed`,`issue_summary`.`completed_date` FROM `issue_summary` WHERE `issue_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*IssueSummary, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _IssueID sql.NullString
		var _TotalIssues sql.NullInt64
		var _New30Days sql.NullInt64
		var _TotalClosed sql.NullInt64
		var _Closed30Days sql.NullInt64
		var _EstimatedWorkMonths sql.NullFloat64
		var _EstimatedWorkMonths30Days sql.NullFloat64
		var _Title sql.NullString
		var _URL sql.NullString
		var _Priority sql.NullString
		var _PriorityID sql.NullString
		var _Status sql.NullString
		var _StatusID sql.NullString
		var _IssueType sql.NullString
		var _IssueTypeID sql.NullString
		var _Resolution sql.NullString
		var _ResolutionID sql.NullString
		var _State sql.NullString
		var _CustomFieldIds sql.NullString
		var _Teams sql.NullString
		var _ParentIssueID sql.NullString
		var _ParentsIssueIds sql.NullString
		var _Metadata sql.NullString
		var _ProjectID sql.NullString
		var _Sprints sql.NullString
		var _Labels sql.NullString
		var _TopLevel sql.NullBool
		var _IsLeaf sql.NullBool
		var _Path sql.NullString
		var _InProgressDuration sql.NullInt64
		var _VerificationDuration sql.NullInt64
		var _InProgressCount sql.NullInt64
		var _ReopenCount sql.NullInt64
		var _MappedType sql.NullString
		var _StrategicParentID sql.NullString
		var _SprintID sql.NullString
		var _IssueProjectName sql.NullString
		var _Users sql.NullString
		var _InitialStartDate sql.NullInt64
		var _TotalDuration sql.NullInt64
		var _CreatedAt sql.NullInt64
		var _ClosedAt sql.NullInt64
		var _PlannedEndDate sql.NullInt64
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefID sql.NullString
		var _CustomFieldIdsVirtual sql.NullString
		var _ReleaseDuration sql.NullInt64
		var _Completed sql.NullBool
		var _CompletedDate sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_IssueID,
			&_TotalIssues,
			&_New30Days,
			&_TotalClosed,
			&_Closed30Days,
			&_EstimatedWorkMonths,
			&_EstimatedWorkMonths30Days,
			&_Title,
			&_URL,
			&_Priority,
			&_PriorityID,
			&_Status,
			&_StatusID,
			&_IssueType,
			&_IssueTypeID,
			&_Resolution,
			&_ResolutionID,
			&_State,
			&_CustomFieldIds,
			&_Teams,
			&_ParentIssueID,
			&_ParentsIssueIds,
			&_Metadata,
			&_ProjectID,
			&_Sprints,
			&_Labels,
			&_TopLevel,
			&_IsLeaf,
			&_Path,
			&_InProgressDuration,
			&_VerificationDuration,
			&_InProgressCount,
			&_ReopenCount,
			&_MappedType,
			&_StrategicParentID,
			&_SprintID,
			&_IssueProjectName,
			&_Users,
			&_InitialStartDate,
			&_TotalDuration,
			&_CreatedAt,
			&_ClosedAt,
			&_PlannedEndDate,
			&_CustomerID,
			&_RefType,
			&_RefID,
			&_CustomFieldIdsVirtual,
			&_ReleaseDuration,
			&_Completed,
			&_CompletedDate,
		)
		if err != nil {
			return nil, err
		}
		t := &IssueSummary{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _TotalIssues.Valid {
			t.SetTotalIssues(int32(_TotalIssues.Int64))
		}
		if _New30Days.Valid {
			t.SetNew30Days(int32(_New30Days.Int64))
		}
		if _TotalClosed.Valid {
			t.SetTotalClosed(int32(_TotalClosed.Int64))
		}
		if _Closed30Days.Valid {
			t.SetClosed30Days(int32(_Closed30Days.Int64))
		}
		if _EstimatedWorkMonths.Valid {
			t.SetEstimatedWorkMonths(_EstimatedWorkMonths.Float64)
		}
		if _EstimatedWorkMonths30Days.Valid {
			t.SetEstimatedWorkMonths30Days(_EstimatedWorkMonths30Days.Float64)
		}
		if _Title.Valid {
			t.SetTitle(_Title.String)
		}
		if _URL.Valid {
			t.SetURL(_URL.String)
		}
		if _Priority.Valid {
			t.SetPriority(_Priority.String)
		}
		if _PriorityID.Valid {
			t.SetPriorityID(_PriorityID.String)
		}
		if _Status.Valid {
			t.SetStatus(_Status.String)
		}
		if _StatusID.Valid {
			t.SetStatusID(_StatusID.String)
		}
		if _IssueType.Valid {
			t.SetIssueType(_IssueType.String)
		}
		if _IssueTypeID.Valid {
			t.SetIssueTypeID(_IssueTypeID.String)
		}
		if _Resolution.Valid {
			t.SetResolution(_Resolution.String)
		}
		if _ResolutionID.Valid {
			t.SetResolutionID(_ResolutionID.String)
		}
		if _State.Valid {
			t.SetState(_State.String)
		}
		if _CustomFieldIds.Valid {
			t.SetCustomFieldIds(_CustomFieldIds.String)
		}
		if _Teams.Valid {
			t.SetTeams(_Teams.String)
		}
		if _ParentIssueID.Valid {
			t.SetParentIssueID(_ParentIssueID.String)
		}
		if _ParentsIssueIds.Valid {
			t.SetParentsIssueIds(_ParentsIssueIds.String)
		}
		if _Metadata.Valid {
			t.SetMetadata(_Metadata.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _Sprints.Valid {
			t.SetSprints(_Sprints.String)
		}
		if _Labels.Valid {
			t.SetLabels(_Labels.String)
		}
		if _TopLevel.Valid {
			t.SetTopLevel(_TopLevel.Bool)
		}
		if _IsLeaf.Valid {
			t.SetIsLeaf(_IsLeaf.Bool)
		}
		if _Path.Valid {
			t.SetPath(_Path.String)
		}
		if _InProgressDuration.Valid {
			t.SetInProgressDuration(_InProgressDuration.Int64)
		}
		if _VerificationDuration.Valid {
			t.SetVerificationDuration(_VerificationDuration.Int64)
		}
		if _InProgressCount.Valid {
			t.SetInProgressCount(int32(_InProgressCount.Int64))
		}
		if _ReopenCount.Valid {
			t.SetReopenCount(int32(_ReopenCount.Int64))
		}
		if _MappedType.Valid {
			t.SetMappedType(_MappedType.String)
		}
		if _StrategicParentID.Valid {
			t.SetStrategicParentID(_StrategicParentID.String)
		}
		if _SprintID.Valid {
			t.SetSprintID(_SprintID.String)
		}
		if _IssueProjectName.Valid {
			t.SetIssueProjectName(_IssueProjectName.String)
		}
		if _Users.Valid {
			t.SetUsers(_Users.String)
		}
		if _InitialStartDate.Valid {
			t.SetInitialStartDate(_InitialStartDate.Int64)
		}
		if _TotalDuration.Valid {
			t.SetTotalDuration(_TotalDuration.Int64)
		}
		if _CreatedAt.Valid {
			t.SetCreatedAt(_CreatedAt.Int64)
		}
		if _ClosedAt.Valid {
			t.SetClosedAt(_ClosedAt.Int64)
		}
		if _PlannedEndDate.Valid {
			t.SetPlannedEndDate(_PlannedEndDate.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefType.Valid {
			t.SetRefType(_RefType.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		if _CustomFieldIdsVirtual.Valid {
			t.SetCustomFieldIdsVirtual(_CustomFieldIdsVirtual.String)
		}
		if _ReleaseDuration.Valid {
			t.SetReleaseDuration(_ReleaseDuration.Int64)
		}
		if _Completed.Valid {
			t.SetCompleted(_Completed.Bool)
		}
		if _CompletedDate.Valid {
			t.SetCompletedDate(_CompletedDate.Int64)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindIssueSummariesByIssueIDTx will find all IssueSummarys by the IssueID value using the provided transaction
func FindIssueSummariesByIssueIDTx(ctx context.Context, tx Tx, value string) ([]*IssueSummary, error) {
	q := "SELECT `issue_summary`.`id`,`issue_summary`.`checksum`,`issue_summary`.`issue_id`,`issue_summary`.`total_issues`,`issue_summary`.`new30_days`,`issue_summary`.`total_closed`,`issue_summary`.`closed30_days`,`issue_summary`.`estimated_work_months`,`issue_summary`.`estimated_work_months30_days`,`issue_summary`.`title`,`issue_summary`.`url`,`issue_summary`.`priority`,`issue_summary`.`priority_id`,`issue_summary`.`status`,`issue_summary`.`status_id`,`issue_summary`.`issue_type`,`issue_summary`.`issue_type_id`,`issue_summary`.`resolution`,`issue_summary`.`resolution_id`,`issue_summary`.`state`,`issue_summary`.`custom_field_ids`,`issue_summary`.`teams`,`issue_summary`.`parent_issue_id`,`issue_summary`.`parents_issue_ids`,`issue_summary`.`metadata`,`issue_summary`.`project_id`,`issue_summary`.`sprints`,`issue_summary`.`labels`,`issue_summary`.`top_level`,`issue_summary`.`is_leaf`,`issue_summary`.`path`,`issue_summary`.`in_progress_duration`,`issue_summary`.`verification_duration`,`issue_summary`.`in_progress_count`,`issue_summary`.`reopen_count`,`issue_summary`.`mapped_type`,`issue_summary`.`strategic_parent_id`,`issue_summary`.`sprint_id`,`issue_summary`.`issue_project_name`,`issue_summary`.`users`,`issue_summary`.`initial_start_date`,`issue_summary`.`total_duration`,`issue_summary`.`created_at`,`issue_summary`.`closed_at`,`issue_summary`.`planned_end_date`,`issue_summary`.`customer_id`,`issue_summary`.`ref_type`,`issue_summary`.`ref_id`,`issue_summary`.`custom_field_ids_virtual`,`issue_summary`.`release_duration`,`issue_summary`.`completed`,`issue_summary`.`completed_date` FROM `issue_summary` WHERE `issue_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*IssueSummary, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _IssueID sql.NullString
		var _TotalIssues sql.NullInt64
		var _New30Days sql.NullInt64
		var _TotalClosed sql.NullInt64
		var _Closed30Days sql.NullInt64
		var _EstimatedWorkMonths sql.NullFloat64
		var _EstimatedWorkMonths30Days sql.NullFloat64
		var _Title sql.NullString
		var _URL sql.NullString
		var _Priority sql.NullString
		var _PriorityID sql.NullString
		var _Status sql.NullString
		var _StatusID sql.NullString
		var _IssueType sql.NullString
		var _IssueTypeID sql.NullString
		var _Resolution sql.NullString
		var _ResolutionID sql.NullString
		var _State sql.NullString
		var _CustomFieldIds sql.NullString
		var _Teams sql.NullString
		var _ParentIssueID sql.NullString
		var _ParentsIssueIds sql.NullString
		var _Metadata sql.NullString
		var _ProjectID sql.NullString
		var _Sprints sql.NullString
		var _Labels sql.NullString
		var _TopLevel sql.NullBool
		var _IsLeaf sql.NullBool
		var _Path sql.NullString
		var _InProgressDuration sql.NullInt64
		var _VerificationDuration sql.NullInt64
		var _InProgressCount sql.NullInt64
		var _ReopenCount sql.NullInt64
		var _MappedType sql.NullString
		var _StrategicParentID sql.NullString
		var _SprintID sql.NullString
		var _IssueProjectName sql.NullString
		var _Users sql.NullString
		var _InitialStartDate sql.NullInt64
		var _TotalDuration sql.NullInt64
		var _CreatedAt sql.NullInt64
		var _ClosedAt sql.NullInt64
		var _PlannedEndDate sql.NullInt64
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefID sql.NullString
		var _CustomFieldIdsVirtual sql.NullString
		var _ReleaseDuration sql.NullInt64
		var _Completed sql.NullBool
		var _CompletedDate sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_IssueID,
			&_TotalIssues,
			&_New30Days,
			&_TotalClosed,
			&_Closed30Days,
			&_EstimatedWorkMonths,
			&_EstimatedWorkMonths30Days,
			&_Title,
			&_URL,
			&_Priority,
			&_PriorityID,
			&_Status,
			&_StatusID,
			&_IssueType,
			&_IssueTypeID,
			&_Resolution,
			&_ResolutionID,
			&_State,
			&_CustomFieldIds,
			&_Teams,
			&_ParentIssueID,
			&_ParentsIssueIds,
			&_Metadata,
			&_ProjectID,
			&_Sprints,
			&_Labels,
			&_TopLevel,
			&_IsLeaf,
			&_Path,
			&_InProgressDuration,
			&_VerificationDuration,
			&_InProgressCount,
			&_ReopenCount,
			&_MappedType,
			&_StrategicParentID,
			&_SprintID,
			&_IssueProjectName,
			&_Users,
			&_InitialStartDate,
			&_TotalDuration,
			&_CreatedAt,
			&_ClosedAt,
			&_PlannedEndDate,
			&_CustomerID,
			&_RefType,
			&_RefID,
			&_CustomFieldIdsVirtual,
			&_ReleaseDuration,
			&_Completed,
			&_CompletedDate,
		)
		if err != nil {
			return nil, err
		}
		t := &IssueSummary{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _TotalIssues.Valid {
			t.SetTotalIssues(int32(_TotalIssues.Int64))
		}
		if _New30Days.Valid {
			t.SetNew30Days(int32(_New30Days.Int64))
		}
		if _TotalClosed.Valid {
			t.SetTotalClosed(int32(_TotalClosed.Int64))
		}
		if _Closed30Days.Valid {
			t.SetClosed30Days(int32(_Closed30Days.Int64))
		}
		if _EstimatedWorkMonths.Valid {
			t.SetEstimatedWorkMonths(_EstimatedWorkMonths.Float64)
		}
		if _EstimatedWorkMonths30Days.Valid {
			t.SetEstimatedWorkMonths30Days(_EstimatedWorkMonths30Days.Float64)
		}
		if _Title.Valid {
			t.SetTitle(_Title.String)
		}
		if _URL.Valid {
			t.SetURL(_URL.String)
		}
		if _Priority.Valid {
			t.SetPriority(_Priority.String)
		}
		if _PriorityID.Valid {
			t.SetPriorityID(_PriorityID.String)
		}
		if _Status.Valid {
			t.SetStatus(_Status.String)
		}
		if _StatusID.Valid {
			t.SetStatusID(_StatusID.String)
		}
		if _IssueType.Valid {
			t.SetIssueType(_IssueType.String)
		}
		if _IssueTypeID.Valid {
			t.SetIssueTypeID(_IssueTypeID.String)
		}
		if _Resolution.Valid {
			t.SetResolution(_Resolution.String)
		}
		if _ResolutionID.Valid {
			t.SetResolutionID(_ResolutionID.String)
		}
		if _State.Valid {
			t.SetState(_State.String)
		}
		if _CustomFieldIds.Valid {
			t.SetCustomFieldIds(_CustomFieldIds.String)
		}
		if _Teams.Valid {
			t.SetTeams(_Teams.String)
		}
		if _ParentIssueID.Valid {
			t.SetParentIssueID(_ParentIssueID.String)
		}
		if _ParentsIssueIds.Valid {
			t.SetParentsIssueIds(_ParentsIssueIds.String)
		}
		if _Metadata.Valid {
			t.SetMetadata(_Metadata.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _Sprints.Valid {
			t.SetSprints(_Sprints.String)
		}
		if _Labels.Valid {
			t.SetLabels(_Labels.String)
		}
		if _TopLevel.Valid {
			t.SetTopLevel(_TopLevel.Bool)
		}
		if _IsLeaf.Valid {
			t.SetIsLeaf(_IsLeaf.Bool)
		}
		if _Path.Valid {
			t.SetPath(_Path.String)
		}
		if _InProgressDuration.Valid {
			t.SetInProgressDuration(_InProgressDuration.Int64)
		}
		if _VerificationDuration.Valid {
			t.SetVerificationDuration(_VerificationDuration.Int64)
		}
		if _InProgressCount.Valid {
			t.SetInProgressCount(int32(_InProgressCount.Int64))
		}
		if _ReopenCount.Valid {
			t.SetReopenCount(int32(_ReopenCount.Int64))
		}
		if _MappedType.Valid {
			t.SetMappedType(_MappedType.String)
		}
		if _StrategicParentID.Valid {
			t.SetStrategicParentID(_StrategicParentID.String)
		}
		if _SprintID.Valid {
			t.SetSprintID(_SprintID.String)
		}
		if _IssueProjectName.Valid {
			t.SetIssueProjectName(_IssueProjectName.String)
		}
		if _Users.Valid {
			t.SetUsers(_Users.String)
		}
		if _InitialStartDate.Valid {
			t.SetInitialStartDate(_InitialStartDate.Int64)
		}
		if _TotalDuration.Valid {
			t.SetTotalDuration(_TotalDuration.Int64)
		}
		if _CreatedAt.Valid {
			t.SetCreatedAt(_CreatedAt.Int64)
		}
		if _ClosedAt.Valid {
			t.SetClosedAt(_ClosedAt.Int64)
		}
		if _PlannedEndDate.Valid {
			t.SetPlannedEndDate(_PlannedEndDate.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefType.Valid {
			t.SetRefType(_RefType.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		if _CustomFieldIdsVirtual.Valid {
			t.SetCustomFieldIdsVirtual(_CustomFieldIdsVirtual.String)
		}
		if _ReleaseDuration.Valid {
			t.SetReleaseDuration(_ReleaseDuration.Int64)
		}
		if _Completed.Valid {
			t.SetCompleted(_Completed.Bool)
		}
		if _CompletedDate.Valid {
			t.SetCompletedDate(_CompletedDate.Int64)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetTotalIssues will return the IssueSummary TotalIssues value
func (t *IssueSummary) GetTotalIssues() int32 {
	return t.TotalIssues
}

// SetTotalIssues will set the IssueSummary TotalIssues value
func (t *IssueSummary) SetTotalIssues(v int32) {
	t.TotalIssues = v
}

// GetNew30Days will return the IssueSummary New30Days value
func (t *IssueSummary) GetNew30Days() int32 {
	return t.New30Days
}

// SetNew30Days will set the IssueSummary New30Days value
func (t *IssueSummary) SetNew30Days(v int32) {
	t.New30Days = v
}

// GetTotalClosed will return the IssueSummary TotalClosed value
func (t *IssueSummary) GetTotalClosed() int32 {
	return t.TotalClosed
}

// SetTotalClosed will set the IssueSummary TotalClosed value
func (t *IssueSummary) SetTotalClosed(v int32) {
	t.TotalClosed = v
}

// GetClosed30Days will return the IssueSummary Closed30Days value
func (t *IssueSummary) GetClosed30Days() int32 {
	return t.Closed30Days
}

// SetClosed30Days will set the IssueSummary Closed30Days value
func (t *IssueSummary) SetClosed30Days(v int32) {
	t.Closed30Days = v
}

// GetEstimatedWorkMonths will return the IssueSummary EstimatedWorkMonths value
func (t *IssueSummary) GetEstimatedWorkMonths() float64 {
	return t.EstimatedWorkMonths
}

// SetEstimatedWorkMonths will set the IssueSummary EstimatedWorkMonths value
func (t *IssueSummary) SetEstimatedWorkMonths(v float64) {
	t.EstimatedWorkMonths = v
}

// GetEstimatedWorkMonths30Days will return the IssueSummary EstimatedWorkMonths30Days value
func (t *IssueSummary) GetEstimatedWorkMonths30Days() float64 {
	return t.EstimatedWorkMonths30Days
}

// SetEstimatedWorkMonths30Days will set the IssueSummary EstimatedWorkMonths30Days value
func (t *IssueSummary) SetEstimatedWorkMonths30Days(v float64) {
	t.EstimatedWorkMonths30Days = v
}

// GetTitle will return the IssueSummary Title value
func (t *IssueSummary) GetTitle() string {
	return t.Title
}

// SetTitle will set the IssueSummary Title value
func (t *IssueSummary) SetTitle(v string) {
	t.Title = v
}

// GetURL will return the IssueSummary URL value
func (t *IssueSummary) GetURL() string {
	if t.URL == nil {
		return ""
	}
	return *t.URL
}

// SetURL will set the IssueSummary URL value
func (t *IssueSummary) SetURL(v string) {
	t.URL = &v
}

// GetPriority will return the IssueSummary Priority value
func (t *IssueSummary) GetPriority() string {
	if t.Priority == nil {
		return ""
	}
	return *t.Priority
}

// SetPriority will set the IssueSummary Priority value
func (t *IssueSummary) SetPriority(v string) {
	t.Priority = &v
}

// GetPriorityID will return the IssueSummary PriorityID value
func (t *IssueSummary) GetPriorityID() string {
	if t.PriorityID == nil {
		return ""
	}
	return *t.PriorityID
}

// SetPriorityID will set the IssueSummary PriorityID value
func (t *IssueSummary) SetPriorityID(v string) {
	t.PriorityID = &v
}

// FindIssueSummariesByPriorityID will find all IssueSummarys by the PriorityID value
func FindIssueSummariesByPriorityID(ctx context.Context, db DB, value string) ([]*IssueSummary, error) {
	q := "SELECT `issue_summary`.`id`,`issue_summary`.`checksum`,`issue_summary`.`issue_id`,`issue_summary`.`total_issues`,`issue_summary`.`new30_days`,`issue_summary`.`total_closed`,`issue_summary`.`closed30_days`,`issue_summary`.`estimated_work_months`,`issue_summary`.`estimated_work_months30_days`,`issue_summary`.`title`,`issue_summary`.`url`,`issue_summary`.`priority`,`issue_summary`.`priority_id`,`issue_summary`.`status`,`issue_summary`.`status_id`,`issue_summary`.`issue_type`,`issue_summary`.`issue_type_id`,`issue_summary`.`resolution`,`issue_summary`.`resolution_id`,`issue_summary`.`state`,`issue_summary`.`custom_field_ids`,`issue_summary`.`teams`,`issue_summary`.`parent_issue_id`,`issue_summary`.`parents_issue_ids`,`issue_summary`.`metadata`,`issue_summary`.`project_id`,`issue_summary`.`sprints`,`issue_summary`.`labels`,`issue_summary`.`top_level`,`issue_summary`.`is_leaf`,`issue_summary`.`path`,`issue_summary`.`in_progress_duration`,`issue_summary`.`verification_duration`,`issue_summary`.`in_progress_count`,`issue_summary`.`reopen_count`,`issue_summary`.`mapped_type`,`issue_summary`.`strategic_parent_id`,`issue_summary`.`sprint_id`,`issue_summary`.`issue_project_name`,`issue_summary`.`users`,`issue_summary`.`initial_start_date`,`issue_summary`.`total_duration`,`issue_summary`.`created_at`,`issue_summary`.`closed_at`,`issue_summary`.`planned_end_date`,`issue_summary`.`customer_id`,`issue_summary`.`ref_type`,`issue_summary`.`ref_id`,`issue_summary`.`custom_field_ids_virtual`,`issue_summary`.`release_duration`,`issue_summary`.`completed`,`issue_summary`.`completed_date` FROM `issue_summary` WHERE `priority_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*IssueSummary, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _IssueID sql.NullString
		var _TotalIssues sql.NullInt64
		var _New30Days sql.NullInt64
		var _TotalClosed sql.NullInt64
		var _Closed30Days sql.NullInt64
		var _EstimatedWorkMonths sql.NullFloat64
		var _EstimatedWorkMonths30Days sql.NullFloat64
		var _Title sql.NullString
		var _URL sql.NullString
		var _Priority sql.NullString
		var _PriorityID sql.NullString
		var _Status sql.NullString
		var _StatusID sql.NullString
		var _IssueType sql.NullString
		var _IssueTypeID sql.NullString
		var _Resolution sql.NullString
		var _ResolutionID sql.NullString
		var _State sql.NullString
		var _CustomFieldIds sql.NullString
		var _Teams sql.NullString
		var _ParentIssueID sql.NullString
		var _ParentsIssueIds sql.NullString
		var _Metadata sql.NullString
		var _ProjectID sql.NullString
		var _Sprints sql.NullString
		var _Labels sql.NullString
		var _TopLevel sql.NullBool
		var _IsLeaf sql.NullBool
		var _Path sql.NullString
		var _InProgressDuration sql.NullInt64
		var _VerificationDuration sql.NullInt64
		var _InProgressCount sql.NullInt64
		var _ReopenCount sql.NullInt64
		var _MappedType sql.NullString
		var _StrategicParentID sql.NullString
		var _SprintID sql.NullString
		var _IssueProjectName sql.NullString
		var _Users sql.NullString
		var _InitialStartDate sql.NullInt64
		var _TotalDuration sql.NullInt64
		var _CreatedAt sql.NullInt64
		var _ClosedAt sql.NullInt64
		var _PlannedEndDate sql.NullInt64
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefID sql.NullString
		var _CustomFieldIdsVirtual sql.NullString
		var _ReleaseDuration sql.NullInt64
		var _Completed sql.NullBool
		var _CompletedDate sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_IssueID,
			&_TotalIssues,
			&_New30Days,
			&_TotalClosed,
			&_Closed30Days,
			&_EstimatedWorkMonths,
			&_EstimatedWorkMonths30Days,
			&_Title,
			&_URL,
			&_Priority,
			&_PriorityID,
			&_Status,
			&_StatusID,
			&_IssueType,
			&_IssueTypeID,
			&_Resolution,
			&_ResolutionID,
			&_State,
			&_CustomFieldIds,
			&_Teams,
			&_ParentIssueID,
			&_ParentsIssueIds,
			&_Metadata,
			&_ProjectID,
			&_Sprints,
			&_Labels,
			&_TopLevel,
			&_IsLeaf,
			&_Path,
			&_InProgressDuration,
			&_VerificationDuration,
			&_InProgressCount,
			&_ReopenCount,
			&_MappedType,
			&_StrategicParentID,
			&_SprintID,
			&_IssueProjectName,
			&_Users,
			&_InitialStartDate,
			&_TotalDuration,
			&_CreatedAt,
			&_ClosedAt,
			&_PlannedEndDate,
			&_CustomerID,
			&_RefType,
			&_RefID,
			&_CustomFieldIdsVirtual,
			&_ReleaseDuration,
			&_Completed,
			&_CompletedDate,
		)
		if err != nil {
			return nil, err
		}
		t := &IssueSummary{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _TotalIssues.Valid {
			t.SetTotalIssues(int32(_TotalIssues.Int64))
		}
		if _New30Days.Valid {
			t.SetNew30Days(int32(_New30Days.Int64))
		}
		if _TotalClosed.Valid {
			t.SetTotalClosed(int32(_TotalClosed.Int64))
		}
		if _Closed30Days.Valid {
			t.SetClosed30Days(int32(_Closed30Days.Int64))
		}
		if _EstimatedWorkMonths.Valid {
			t.SetEstimatedWorkMonths(_EstimatedWorkMonths.Float64)
		}
		if _EstimatedWorkMonths30Days.Valid {
			t.SetEstimatedWorkMonths30Days(_EstimatedWorkMonths30Days.Float64)
		}
		if _Title.Valid {
			t.SetTitle(_Title.String)
		}
		if _URL.Valid {
			t.SetURL(_URL.String)
		}
		if _Priority.Valid {
			t.SetPriority(_Priority.String)
		}
		if _PriorityID.Valid {
			t.SetPriorityID(_PriorityID.String)
		}
		if _Status.Valid {
			t.SetStatus(_Status.String)
		}
		if _StatusID.Valid {
			t.SetStatusID(_StatusID.String)
		}
		if _IssueType.Valid {
			t.SetIssueType(_IssueType.String)
		}
		if _IssueTypeID.Valid {
			t.SetIssueTypeID(_IssueTypeID.String)
		}
		if _Resolution.Valid {
			t.SetResolution(_Resolution.String)
		}
		if _ResolutionID.Valid {
			t.SetResolutionID(_ResolutionID.String)
		}
		if _State.Valid {
			t.SetState(_State.String)
		}
		if _CustomFieldIds.Valid {
			t.SetCustomFieldIds(_CustomFieldIds.String)
		}
		if _Teams.Valid {
			t.SetTeams(_Teams.String)
		}
		if _ParentIssueID.Valid {
			t.SetParentIssueID(_ParentIssueID.String)
		}
		if _ParentsIssueIds.Valid {
			t.SetParentsIssueIds(_ParentsIssueIds.String)
		}
		if _Metadata.Valid {
			t.SetMetadata(_Metadata.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _Sprints.Valid {
			t.SetSprints(_Sprints.String)
		}
		if _Labels.Valid {
			t.SetLabels(_Labels.String)
		}
		if _TopLevel.Valid {
			t.SetTopLevel(_TopLevel.Bool)
		}
		if _IsLeaf.Valid {
			t.SetIsLeaf(_IsLeaf.Bool)
		}
		if _Path.Valid {
			t.SetPath(_Path.String)
		}
		if _InProgressDuration.Valid {
			t.SetInProgressDuration(_InProgressDuration.Int64)
		}
		if _VerificationDuration.Valid {
			t.SetVerificationDuration(_VerificationDuration.Int64)
		}
		if _InProgressCount.Valid {
			t.SetInProgressCount(int32(_InProgressCount.Int64))
		}
		if _ReopenCount.Valid {
			t.SetReopenCount(int32(_ReopenCount.Int64))
		}
		if _MappedType.Valid {
			t.SetMappedType(_MappedType.String)
		}
		if _StrategicParentID.Valid {
			t.SetStrategicParentID(_StrategicParentID.String)
		}
		if _SprintID.Valid {
			t.SetSprintID(_SprintID.String)
		}
		if _IssueProjectName.Valid {
			t.SetIssueProjectName(_IssueProjectName.String)
		}
		if _Users.Valid {
			t.SetUsers(_Users.String)
		}
		if _InitialStartDate.Valid {
			t.SetInitialStartDate(_InitialStartDate.Int64)
		}
		if _TotalDuration.Valid {
			t.SetTotalDuration(_TotalDuration.Int64)
		}
		if _CreatedAt.Valid {
			t.SetCreatedAt(_CreatedAt.Int64)
		}
		if _ClosedAt.Valid {
			t.SetClosedAt(_ClosedAt.Int64)
		}
		if _PlannedEndDate.Valid {
			t.SetPlannedEndDate(_PlannedEndDate.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefType.Valid {
			t.SetRefType(_RefType.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		if _CustomFieldIdsVirtual.Valid {
			t.SetCustomFieldIdsVirtual(_CustomFieldIdsVirtual.String)
		}
		if _ReleaseDuration.Valid {
			t.SetReleaseDuration(_ReleaseDuration.Int64)
		}
		if _Completed.Valid {
			t.SetCompleted(_Completed.Bool)
		}
		if _CompletedDate.Valid {
			t.SetCompletedDate(_CompletedDate.Int64)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindIssueSummariesByPriorityIDTx will find all IssueSummarys by the PriorityID value using the provided transaction
func FindIssueSummariesByPriorityIDTx(ctx context.Context, tx Tx, value string) ([]*IssueSummary, error) {
	q := "SELECT `issue_summary`.`id`,`issue_summary`.`checksum`,`issue_summary`.`issue_id`,`issue_summary`.`total_issues`,`issue_summary`.`new30_days`,`issue_summary`.`total_closed`,`issue_summary`.`closed30_days`,`issue_summary`.`estimated_work_months`,`issue_summary`.`estimated_work_months30_days`,`issue_summary`.`title`,`issue_summary`.`url`,`issue_summary`.`priority`,`issue_summary`.`priority_id`,`issue_summary`.`status`,`issue_summary`.`status_id`,`issue_summary`.`issue_type`,`issue_summary`.`issue_type_id`,`issue_summary`.`resolution`,`issue_summary`.`resolution_id`,`issue_summary`.`state`,`issue_summary`.`custom_field_ids`,`issue_summary`.`teams`,`issue_summary`.`parent_issue_id`,`issue_summary`.`parents_issue_ids`,`issue_summary`.`metadata`,`issue_summary`.`project_id`,`issue_summary`.`sprints`,`issue_summary`.`labels`,`issue_summary`.`top_level`,`issue_summary`.`is_leaf`,`issue_summary`.`path`,`issue_summary`.`in_progress_duration`,`issue_summary`.`verification_duration`,`issue_summary`.`in_progress_count`,`issue_summary`.`reopen_count`,`issue_summary`.`mapped_type`,`issue_summary`.`strategic_parent_id`,`issue_summary`.`sprint_id`,`issue_summary`.`issue_project_name`,`issue_summary`.`users`,`issue_summary`.`initial_start_date`,`issue_summary`.`total_duration`,`issue_summary`.`created_at`,`issue_summary`.`closed_at`,`issue_summary`.`planned_end_date`,`issue_summary`.`customer_id`,`issue_summary`.`ref_type`,`issue_summary`.`ref_id`,`issue_summary`.`custom_field_ids_virtual`,`issue_summary`.`release_duration`,`issue_summary`.`completed`,`issue_summary`.`completed_date` FROM `issue_summary` WHERE `priority_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*IssueSummary, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _IssueID sql.NullString
		var _TotalIssues sql.NullInt64
		var _New30Days sql.NullInt64
		var _TotalClosed sql.NullInt64
		var _Closed30Days sql.NullInt64
		var _EstimatedWorkMonths sql.NullFloat64
		var _EstimatedWorkMonths30Days sql.NullFloat64
		var _Title sql.NullString
		var _URL sql.NullString
		var _Priority sql.NullString
		var _PriorityID sql.NullString
		var _Status sql.NullString
		var _StatusID sql.NullString
		var _IssueType sql.NullString
		var _IssueTypeID sql.NullString
		var _Resolution sql.NullString
		var _ResolutionID sql.NullString
		var _State sql.NullString
		var _CustomFieldIds sql.NullString
		var _Teams sql.NullString
		var _ParentIssueID sql.NullString
		var _ParentsIssueIds sql.NullString
		var _Metadata sql.NullString
		var _ProjectID sql.NullString
		var _Sprints sql.NullString
		var _Labels sql.NullString
		var _TopLevel sql.NullBool
		var _IsLeaf sql.NullBool
		var _Path sql.NullString
		var _InProgressDuration sql.NullInt64
		var _VerificationDuration sql.NullInt64
		var _InProgressCount sql.NullInt64
		var _ReopenCount sql.NullInt64
		var _MappedType sql.NullString
		var _StrategicParentID sql.NullString
		var _SprintID sql.NullString
		var _IssueProjectName sql.NullString
		var _Users sql.NullString
		var _InitialStartDate sql.NullInt64
		var _TotalDuration sql.NullInt64
		var _CreatedAt sql.NullInt64
		var _ClosedAt sql.NullInt64
		var _PlannedEndDate sql.NullInt64
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefID sql.NullString
		var _CustomFieldIdsVirtual sql.NullString
		var _ReleaseDuration sql.NullInt64
		var _Completed sql.NullBool
		var _CompletedDate sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_IssueID,
			&_TotalIssues,
			&_New30Days,
			&_TotalClosed,
			&_Closed30Days,
			&_EstimatedWorkMonths,
			&_EstimatedWorkMonths30Days,
			&_Title,
			&_URL,
			&_Priority,
			&_PriorityID,
			&_Status,
			&_StatusID,
			&_IssueType,
			&_IssueTypeID,
			&_Resolution,
			&_ResolutionID,
			&_State,
			&_CustomFieldIds,
			&_Teams,
			&_ParentIssueID,
			&_ParentsIssueIds,
			&_Metadata,
			&_ProjectID,
			&_Sprints,
			&_Labels,
			&_TopLevel,
			&_IsLeaf,
			&_Path,
			&_InProgressDuration,
			&_VerificationDuration,
			&_InProgressCount,
			&_ReopenCount,
			&_MappedType,
			&_StrategicParentID,
			&_SprintID,
			&_IssueProjectName,
			&_Users,
			&_InitialStartDate,
			&_TotalDuration,
			&_CreatedAt,
			&_ClosedAt,
			&_PlannedEndDate,
			&_CustomerID,
			&_RefType,
			&_RefID,
			&_CustomFieldIdsVirtual,
			&_ReleaseDuration,
			&_Completed,
			&_CompletedDate,
		)
		if err != nil {
			return nil, err
		}
		t := &IssueSummary{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _TotalIssues.Valid {
			t.SetTotalIssues(int32(_TotalIssues.Int64))
		}
		if _New30Days.Valid {
			t.SetNew30Days(int32(_New30Days.Int64))
		}
		if _TotalClosed.Valid {
			t.SetTotalClosed(int32(_TotalClosed.Int64))
		}
		if _Closed30Days.Valid {
			t.SetClosed30Days(int32(_Closed30Days.Int64))
		}
		if _EstimatedWorkMonths.Valid {
			t.SetEstimatedWorkMonths(_EstimatedWorkMonths.Float64)
		}
		if _EstimatedWorkMonths30Days.Valid {
			t.SetEstimatedWorkMonths30Days(_EstimatedWorkMonths30Days.Float64)
		}
		if _Title.Valid {
			t.SetTitle(_Title.String)
		}
		if _URL.Valid {
			t.SetURL(_URL.String)
		}
		if _Priority.Valid {
			t.SetPriority(_Priority.String)
		}
		if _PriorityID.Valid {
			t.SetPriorityID(_PriorityID.String)
		}
		if _Status.Valid {
			t.SetStatus(_Status.String)
		}
		if _StatusID.Valid {
			t.SetStatusID(_StatusID.String)
		}
		if _IssueType.Valid {
			t.SetIssueType(_IssueType.String)
		}
		if _IssueTypeID.Valid {
			t.SetIssueTypeID(_IssueTypeID.String)
		}
		if _Resolution.Valid {
			t.SetResolution(_Resolution.String)
		}
		if _ResolutionID.Valid {
			t.SetResolutionID(_ResolutionID.String)
		}
		if _State.Valid {
			t.SetState(_State.String)
		}
		if _CustomFieldIds.Valid {
			t.SetCustomFieldIds(_CustomFieldIds.String)
		}
		if _Teams.Valid {
			t.SetTeams(_Teams.String)
		}
		if _ParentIssueID.Valid {
			t.SetParentIssueID(_ParentIssueID.String)
		}
		if _ParentsIssueIds.Valid {
			t.SetParentsIssueIds(_ParentsIssueIds.String)
		}
		if _Metadata.Valid {
			t.SetMetadata(_Metadata.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _Sprints.Valid {
			t.SetSprints(_Sprints.String)
		}
		if _Labels.Valid {
			t.SetLabels(_Labels.String)
		}
		if _TopLevel.Valid {
			t.SetTopLevel(_TopLevel.Bool)
		}
		if _IsLeaf.Valid {
			t.SetIsLeaf(_IsLeaf.Bool)
		}
		if _Path.Valid {
			t.SetPath(_Path.String)
		}
		if _InProgressDuration.Valid {
			t.SetInProgressDuration(_InProgressDuration.Int64)
		}
		if _VerificationDuration.Valid {
			t.SetVerificationDuration(_VerificationDuration.Int64)
		}
		if _InProgressCount.Valid {
			t.SetInProgressCount(int32(_InProgressCount.Int64))
		}
		if _ReopenCount.Valid {
			t.SetReopenCount(int32(_ReopenCount.Int64))
		}
		if _MappedType.Valid {
			t.SetMappedType(_MappedType.String)
		}
		if _StrategicParentID.Valid {
			t.SetStrategicParentID(_StrategicParentID.String)
		}
		if _SprintID.Valid {
			t.SetSprintID(_SprintID.String)
		}
		if _IssueProjectName.Valid {
			t.SetIssueProjectName(_IssueProjectName.String)
		}
		if _Users.Valid {
			t.SetUsers(_Users.String)
		}
		if _InitialStartDate.Valid {
			t.SetInitialStartDate(_InitialStartDate.Int64)
		}
		if _TotalDuration.Valid {
			t.SetTotalDuration(_TotalDuration.Int64)
		}
		if _CreatedAt.Valid {
			t.SetCreatedAt(_CreatedAt.Int64)
		}
		if _ClosedAt.Valid {
			t.SetClosedAt(_ClosedAt.Int64)
		}
		if _PlannedEndDate.Valid {
			t.SetPlannedEndDate(_PlannedEndDate.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefType.Valid {
			t.SetRefType(_RefType.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		if _CustomFieldIdsVirtual.Valid {
			t.SetCustomFieldIdsVirtual(_CustomFieldIdsVirtual.String)
		}
		if _ReleaseDuration.Valid {
			t.SetReleaseDuration(_ReleaseDuration.Int64)
		}
		if _Completed.Valid {
			t.SetCompleted(_Completed.Bool)
		}
		if _CompletedDate.Valid {
			t.SetCompletedDate(_CompletedDate.Int64)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetStatus will return the IssueSummary Status value
func (t *IssueSummary) GetStatus() string {
	if t.Status == nil {
		return ""
	}
	return *t.Status
}

// SetStatus will set the IssueSummary Status value
func (t *IssueSummary) SetStatus(v string) {
	t.Status = &v
}

// GetStatusID will return the IssueSummary StatusID value
func (t *IssueSummary) GetStatusID() string {
	if t.StatusID == nil {
		return ""
	}
	return *t.StatusID
}

// SetStatusID will set the IssueSummary StatusID value
func (t *IssueSummary) SetStatusID(v string) {
	t.StatusID = &v
}

// GetIssueType will return the IssueSummary IssueType value
func (t *IssueSummary) GetIssueType() string {
	return t.IssueType
}

// SetIssueType will set the IssueSummary IssueType value
func (t *IssueSummary) SetIssueType(v string) {
	t.IssueType = v
}

// GetIssueTypeID will return the IssueSummary IssueTypeID value
func (t *IssueSummary) GetIssueTypeID() string {
	if t.IssueTypeID == nil {
		return ""
	}
	return *t.IssueTypeID
}

// SetIssueTypeID will set the IssueSummary IssueTypeID value
func (t *IssueSummary) SetIssueTypeID(v string) {
	t.IssueTypeID = &v
}

// FindIssueSummariesByIssueTypeID will find all IssueSummarys by the IssueTypeID value
func FindIssueSummariesByIssueTypeID(ctx context.Context, db DB, value string) ([]*IssueSummary, error) {
	q := "SELECT `issue_summary`.`id`,`issue_summary`.`checksum`,`issue_summary`.`issue_id`,`issue_summary`.`total_issues`,`issue_summary`.`new30_days`,`issue_summary`.`total_closed`,`issue_summary`.`closed30_days`,`issue_summary`.`estimated_work_months`,`issue_summary`.`estimated_work_months30_days`,`issue_summary`.`title`,`issue_summary`.`url`,`issue_summary`.`priority`,`issue_summary`.`priority_id`,`issue_summary`.`status`,`issue_summary`.`status_id`,`issue_summary`.`issue_type`,`issue_summary`.`issue_type_id`,`issue_summary`.`resolution`,`issue_summary`.`resolution_id`,`issue_summary`.`state`,`issue_summary`.`custom_field_ids`,`issue_summary`.`teams`,`issue_summary`.`parent_issue_id`,`issue_summary`.`parents_issue_ids`,`issue_summary`.`metadata`,`issue_summary`.`project_id`,`issue_summary`.`sprints`,`issue_summary`.`labels`,`issue_summary`.`top_level`,`issue_summary`.`is_leaf`,`issue_summary`.`path`,`issue_summary`.`in_progress_duration`,`issue_summary`.`verification_duration`,`issue_summary`.`in_progress_count`,`issue_summary`.`reopen_count`,`issue_summary`.`mapped_type`,`issue_summary`.`strategic_parent_id`,`issue_summary`.`sprint_id`,`issue_summary`.`issue_project_name`,`issue_summary`.`users`,`issue_summary`.`initial_start_date`,`issue_summary`.`total_duration`,`issue_summary`.`created_at`,`issue_summary`.`closed_at`,`issue_summary`.`planned_end_date`,`issue_summary`.`customer_id`,`issue_summary`.`ref_type`,`issue_summary`.`ref_id`,`issue_summary`.`custom_field_ids_virtual`,`issue_summary`.`release_duration`,`issue_summary`.`completed`,`issue_summary`.`completed_date` FROM `issue_summary` WHERE `issue_type_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*IssueSummary, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _IssueID sql.NullString
		var _TotalIssues sql.NullInt64
		var _New30Days sql.NullInt64
		var _TotalClosed sql.NullInt64
		var _Closed30Days sql.NullInt64
		var _EstimatedWorkMonths sql.NullFloat64
		var _EstimatedWorkMonths30Days sql.NullFloat64
		var _Title sql.NullString
		var _URL sql.NullString
		var _Priority sql.NullString
		var _PriorityID sql.NullString
		var _Status sql.NullString
		var _StatusID sql.NullString
		var _IssueType sql.NullString
		var _IssueTypeID sql.NullString
		var _Resolution sql.NullString
		var _ResolutionID sql.NullString
		var _State sql.NullString
		var _CustomFieldIds sql.NullString
		var _Teams sql.NullString
		var _ParentIssueID sql.NullString
		var _ParentsIssueIds sql.NullString
		var _Metadata sql.NullString
		var _ProjectID sql.NullString
		var _Sprints sql.NullString
		var _Labels sql.NullString
		var _TopLevel sql.NullBool
		var _IsLeaf sql.NullBool
		var _Path sql.NullString
		var _InProgressDuration sql.NullInt64
		var _VerificationDuration sql.NullInt64
		var _InProgressCount sql.NullInt64
		var _ReopenCount sql.NullInt64
		var _MappedType sql.NullString
		var _StrategicParentID sql.NullString
		var _SprintID sql.NullString
		var _IssueProjectName sql.NullString
		var _Users sql.NullString
		var _InitialStartDate sql.NullInt64
		var _TotalDuration sql.NullInt64
		var _CreatedAt sql.NullInt64
		var _ClosedAt sql.NullInt64
		var _PlannedEndDate sql.NullInt64
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefID sql.NullString
		var _CustomFieldIdsVirtual sql.NullString
		var _ReleaseDuration sql.NullInt64
		var _Completed sql.NullBool
		var _CompletedDate sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_IssueID,
			&_TotalIssues,
			&_New30Days,
			&_TotalClosed,
			&_Closed30Days,
			&_EstimatedWorkMonths,
			&_EstimatedWorkMonths30Days,
			&_Title,
			&_URL,
			&_Priority,
			&_PriorityID,
			&_Status,
			&_StatusID,
			&_IssueType,
			&_IssueTypeID,
			&_Resolution,
			&_ResolutionID,
			&_State,
			&_CustomFieldIds,
			&_Teams,
			&_ParentIssueID,
			&_ParentsIssueIds,
			&_Metadata,
			&_ProjectID,
			&_Sprints,
			&_Labels,
			&_TopLevel,
			&_IsLeaf,
			&_Path,
			&_InProgressDuration,
			&_VerificationDuration,
			&_InProgressCount,
			&_ReopenCount,
			&_MappedType,
			&_StrategicParentID,
			&_SprintID,
			&_IssueProjectName,
			&_Users,
			&_InitialStartDate,
			&_TotalDuration,
			&_CreatedAt,
			&_ClosedAt,
			&_PlannedEndDate,
			&_CustomerID,
			&_RefType,
			&_RefID,
			&_CustomFieldIdsVirtual,
			&_ReleaseDuration,
			&_Completed,
			&_CompletedDate,
		)
		if err != nil {
			return nil, err
		}
		t := &IssueSummary{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _TotalIssues.Valid {
			t.SetTotalIssues(int32(_TotalIssues.Int64))
		}
		if _New30Days.Valid {
			t.SetNew30Days(int32(_New30Days.Int64))
		}
		if _TotalClosed.Valid {
			t.SetTotalClosed(int32(_TotalClosed.Int64))
		}
		if _Closed30Days.Valid {
			t.SetClosed30Days(int32(_Closed30Days.Int64))
		}
		if _EstimatedWorkMonths.Valid {
			t.SetEstimatedWorkMonths(_EstimatedWorkMonths.Float64)
		}
		if _EstimatedWorkMonths30Days.Valid {
			t.SetEstimatedWorkMonths30Days(_EstimatedWorkMonths30Days.Float64)
		}
		if _Title.Valid {
			t.SetTitle(_Title.String)
		}
		if _URL.Valid {
			t.SetURL(_URL.String)
		}
		if _Priority.Valid {
			t.SetPriority(_Priority.String)
		}
		if _PriorityID.Valid {
			t.SetPriorityID(_PriorityID.String)
		}
		if _Status.Valid {
			t.SetStatus(_Status.String)
		}
		if _StatusID.Valid {
			t.SetStatusID(_StatusID.String)
		}
		if _IssueType.Valid {
			t.SetIssueType(_IssueType.String)
		}
		if _IssueTypeID.Valid {
			t.SetIssueTypeID(_IssueTypeID.String)
		}
		if _Resolution.Valid {
			t.SetResolution(_Resolution.String)
		}
		if _ResolutionID.Valid {
			t.SetResolutionID(_ResolutionID.String)
		}
		if _State.Valid {
			t.SetState(_State.String)
		}
		if _CustomFieldIds.Valid {
			t.SetCustomFieldIds(_CustomFieldIds.String)
		}
		if _Teams.Valid {
			t.SetTeams(_Teams.String)
		}
		if _ParentIssueID.Valid {
			t.SetParentIssueID(_ParentIssueID.String)
		}
		if _ParentsIssueIds.Valid {
			t.SetParentsIssueIds(_ParentsIssueIds.String)
		}
		if _Metadata.Valid {
			t.SetMetadata(_Metadata.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _Sprints.Valid {
			t.SetSprints(_Sprints.String)
		}
		if _Labels.Valid {
			t.SetLabels(_Labels.String)
		}
		if _TopLevel.Valid {
			t.SetTopLevel(_TopLevel.Bool)
		}
		if _IsLeaf.Valid {
			t.SetIsLeaf(_IsLeaf.Bool)
		}
		if _Path.Valid {
			t.SetPath(_Path.String)
		}
		if _InProgressDuration.Valid {
			t.SetInProgressDuration(_InProgressDuration.Int64)
		}
		if _VerificationDuration.Valid {
			t.SetVerificationDuration(_VerificationDuration.Int64)
		}
		if _InProgressCount.Valid {
			t.SetInProgressCount(int32(_InProgressCount.Int64))
		}
		if _ReopenCount.Valid {
			t.SetReopenCount(int32(_ReopenCount.Int64))
		}
		if _MappedType.Valid {
			t.SetMappedType(_MappedType.String)
		}
		if _StrategicParentID.Valid {
			t.SetStrategicParentID(_StrategicParentID.String)
		}
		if _SprintID.Valid {
			t.SetSprintID(_SprintID.String)
		}
		if _IssueProjectName.Valid {
			t.SetIssueProjectName(_IssueProjectName.String)
		}
		if _Users.Valid {
			t.SetUsers(_Users.String)
		}
		if _InitialStartDate.Valid {
			t.SetInitialStartDate(_InitialStartDate.Int64)
		}
		if _TotalDuration.Valid {
			t.SetTotalDuration(_TotalDuration.Int64)
		}
		if _CreatedAt.Valid {
			t.SetCreatedAt(_CreatedAt.Int64)
		}
		if _ClosedAt.Valid {
			t.SetClosedAt(_ClosedAt.Int64)
		}
		if _PlannedEndDate.Valid {
			t.SetPlannedEndDate(_PlannedEndDate.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefType.Valid {
			t.SetRefType(_RefType.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		if _CustomFieldIdsVirtual.Valid {
			t.SetCustomFieldIdsVirtual(_CustomFieldIdsVirtual.String)
		}
		if _ReleaseDuration.Valid {
			t.SetReleaseDuration(_ReleaseDuration.Int64)
		}
		if _Completed.Valid {
			t.SetCompleted(_Completed.Bool)
		}
		if _CompletedDate.Valid {
			t.SetCompletedDate(_CompletedDate.Int64)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindIssueSummariesByIssueTypeIDTx will find all IssueSummarys by the IssueTypeID value using the provided transaction
func FindIssueSummariesByIssueTypeIDTx(ctx context.Context, tx Tx, value string) ([]*IssueSummary, error) {
	q := "SELECT `issue_summary`.`id`,`issue_summary`.`checksum`,`issue_summary`.`issue_id`,`issue_summary`.`total_issues`,`issue_summary`.`new30_days`,`issue_summary`.`total_closed`,`issue_summary`.`closed30_days`,`issue_summary`.`estimated_work_months`,`issue_summary`.`estimated_work_months30_days`,`issue_summary`.`title`,`issue_summary`.`url`,`issue_summary`.`priority`,`issue_summary`.`priority_id`,`issue_summary`.`status`,`issue_summary`.`status_id`,`issue_summary`.`issue_type`,`issue_summary`.`issue_type_id`,`issue_summary`.`resolution`,`issue_summary`.`resolution_id`,`issue_summary`.`state`,`issue_summary`.`custom_field_ids`,`issue_summary`.`teams`,`issue_summary`.`parent_issue_id`,`issue_summary`.`parents_issue_ids`,`issue_summary`.`metadata`,`issue_summary`.`project_id`,`issue_summary`.`sprints`,`issue_summary`.`labels`,`issue_summary`.`top_level`,`issue_summary`.`is_leaf`,`issue_summary`.`path`,`issue_summary`.`in_progress_duration`,`issue_summary`.`verification_duration`,`issue_summary`.`in_progress_count`,`issue_summary`.`reopen_count`,`issue_summary`.`mapped_type`,`issue_summary`.`strategic_parent_id`,`issue_summary`.`sprint_id`,`issue_summary`.`issue_project_name`,`issue_summary`.`users`,`issue_summary`.`initial_start_date`,`issue_summary`.`total_duration`,`issue_summary`.`created_at`,`issue_summary`.`closed_at`,`issue_summary`.`planned_end_date`,`issue_summary`.`customer_id`,`issue_summary`.`ref_type`,`issue_summary`.`ref_id`,`issue_summary`.`custom_field_ids_virtual`,`issue_summary`.`release_duration`,`issue_summary`.`completed`,`issue_summary`.`completed_date` FROM `issue_summary` WHERE `issue_type_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*IssueSummary, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _IssueID sql.NullString
		var _TotalIssues sql.NullInt64
		var _New30Days sql.NullInt64
		var _TotalClosed sql.NullInt64
		var _Closed30Days sql.NullInt64
		var _EstimatedWorkMonths sql.NullFloat64
		var _EstimatedWorkMonths30Days sql.NullFloat64
		var _Title sql.NullString
		var _URL sql.NullString
		var _Priority sql.NullString
		var _PriorityID sql.NullString
		var _Status sql.NullString
		var _StatusID sql.NullString
		var _IssueType sql.NullString
		var _IssueTypeID sql.NullString
		var _Resolution sql.NullString
		var _ResolutionID sql.NullString
		var _State sql.NullString
		var _CustomFieldIds sql.NullString
		var _Teams sql.NullString
		var _ParentIssueID sql.NullString
		var _ParentsIssueIds sql.NullString
		var _Metadata sql.NullString
		var _ProjectID sql.NullString
		var _Sprints sql.NullString
		var _Labels sql.NullString
		var _TopLevel sql.NullBool
		var _IsLeaf sql.NullBool
		var _Path sql.NullString
		var _InProgressDuration sql.NullInt64
		var _VerificationDuration sql.NullInt64
		var _InProgressCount sql.NullInt64
		var _ReopenCount sql.NullInt64
		var _MappedType sql.NullString
		var _StrategicParentID sql.NullString
		var _SprintID sql.NullString
		var _IssueProjectName sql.NullString
		var _Users sql.NullString
		var _InitialStartDate sql.NullInt64
		var _TotalDuration sql.NullInt64
		var _CreatedAt sql.NullInt64
		var _ClosedAt sql.NullInt64
		var _PlannedEndDate sql.NullInt64
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefID sql.NullString
		var _CustomFieldIdsVirtual sql.NullString
		var _ReleaseDuration sql.NullInt64
		var _Completed sql.NullBool
		var _CompletedDate sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_IssueID,
			&_TotalIssues,
			&_New30Days,
			&_TotalClosed,
			&_Closed30Days,
			&_EstimatedWorkMonths,
			&_EstimatedWorkMonths30Days,
			&_Title,
			&_URL,
			&_Priority,
			&_PriorityID,
			&_Status,
			&_StatusID,
			&_IssueType,
			&_IssueTypeID,
			&_Resolution,
			&_ResolutionID,
			&_State,
			&_CustomFieldIds,
			&_Teams,
			&_ParentIssueID,
			&_ParentsIssueIds,
			&_Metadata,
			&_ProjectID,
			&_Sprints,
			&_Labels,
			&_TopLevel,
			&_IsLeaf,
			&_Path,
			&_InProgressDuration,
			&_VerificationDuration,
			&_InProgressCount,
			&_ReopenCount,
			&_MappedType,
			&_StrategicParentID,
			&_SprintID,
			&_IssueProjectName,
			&_Users,
			&_InitialStartDate,
			&_TotalDuration,
			&_CreatedAt,
			&_ClosedAt,
			&_PlannedEndDate,
			&_CustomerID,
			&_RefType,
			&_RefID,
			&_CustomFieldIdsVirtual,
			&_ReleaseDuration,
			&_Completed,
			&_CompletedDate,
		)
		if err != nil {
			return nil, err
		}
		t := &IssueSummary{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _TotalIssues.Valid {
			t.SetTotalIssues(int32(_TotalIssues.Int64))
		}
		if _New30Days.Valid {
			t.SetNew30Days(int32(_New30Days.Int64))
		}
		if _TotalClosed.Valid {
			t.SetTotalClosed(int32(_TotalClosed.Int64))
		}
		if _Closed30Days.Valid {
			t.SetClosed30Days(int32(_Closed30Days.Int64))
		}
		if _EstimatedWorkMonths.Valid {
			t.SetEstimatedWorkMonths(_EstimatedWorkMonths.Float64)
		}
		if _EstimatedWorkMonths30Days.Valid {
			t.SetEstimatedWorkMonths30Days(_EstimatedWorkMonths30Days.Float64)
		}
		if _Title.Valid {
			t.SetTitle(_Title.String)
		}
		if _URL.Valid {
			t.SetURL(_URL.String)
		}
		if _Priority.Valid {
			t.SetPriority(_Priority.String)
		}
		if _PriorityID.Valid {
			t.SetPriorityID(_PriorityID.String)
		}
		if _Status.Valid {
			t.SetStatus(_Status.String)
		}
		if _StatusID.Valid {
			t.SetStatusID(_StatusID.String)
		}
		if _IssueType.Valid {
			t.SetIssueType(_IssueType.String)
		}
		if _IssueTypeID.Valid {
			t.SetIssueTypeID(_IssueTypeID.String)
		}
		if _Resolution.Valid {
			t.SetResolution(_Resolution.String)
		}
		if _ResolutionID.Valid {
			t.SetResolutionID(_ResolutionID.String)
		}
		if _State.Valid {
			t.SetState(_State.String)
		}
		if _CustomFieldIds.Valid {
			t.SetCustomFieldIds(_CustomFieldIds.String)
		}
		if _Teams.Valid {
			t.SetTeams(_Teams.String)
		}
		if _ParentIssueID.Valid {
			t.SetParentIssueID(_ParentIssueID.String)
		}
		if _ParentsIssueIds.Valid {
			t.SetParentsIssueIds(_ParentsIssueIds.String)
		}
		if _Metadata.Valid {
			t.SetMetadata(_Metadata.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _Sprints.Valid {
			t.SetSprints(_Sprints.String)
		}
		if _Labels.Valid {
			t.SetLabels(_Labels.String)
		}
		if _TopLevel.Valid {
			t.SetTopLevel(_TopLevel.Bool)
		}
		if _IsLeaf.Valid {
			t.SetIsLeaf(_IsLeaf.Bool)
		}
		if _Path.Valid {
			t.SetPath(_Path.String)
		}
		if _InProgressDuration.Valid {
			t.SetInProgressDuration(_InProgressDuration.Int64)
		}
		if _VerificationDuration.Valid {
			t.SetVerificationDuration(_VerificationDuration.Int64)
		}
		if _InProgressCount.Valid {
			t.SetInProgressCount(int32(_InProgressCount.Int64))
		}
		if _ReopenCount.Valid {
			t.SetReopenCount(int32(_ReopenCount.Int64))
		}
		if _MappedType.Valid {
			t.SetMappedType(_MappedType.String)
		}
		if _StrategicParentID.Valid {
			t.SetStrategicParentID(_StrategicParentID.String)
		}
		if _SprintID.Valid {
			t.SetSprintID(_SprintID.String)
		}
		if _IssueProjectName.Valid {
			t.SetIssueProjectName(_IssueProjectName.String)
		}
		if _Users.Valid {
			t.SetUsers(_Users.String)
		}
		if _InitialStartDate.Valid {
			t.SetInitialStartDate(_InitialStartDate.Int64)
		}
		if _TotalDuration.Valid {
			t.SetTotalDuration(_TotalDuration.Int64)
		}
		if _CreatedAt.Valid {
			t.SetCreatedAt(_CreatedAt.Int64)
		}
		if _ClosedAt.Valid {
			t.SetClosedAt(_ClosedAt.Int64)
		}
		if _PlannedEndDate.Valid {
			t.SetPlannedEndDate(_PlannedEndDate.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefType.Valid {
			t.SetRefType(_RefType.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		if _CustomFieldIdsVirtual.Valid {
			t.SetCustomFieldIdsVirtual(_CustomFieldIdsVirtual.String)
		}
		if _ReleaseDuration.Valid {
			t.SetReleaseDuration(_ReleaseDuration.Int64)
		}
		if _Completed.Valid {
			t.SetCompleted(_Completed.Bool)
		}
		if _CompletedDate.Valid {
			t.SetCompletedDate(_CompletedDate.Int64)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetResolution will return the IssueSummary Resolution value
func (t *IssueSummary) GetResolution() string {
	if t.Resolution == nil {
		return ""
	}
	return *t.Resolution
}

// SetResolution will set the IssueSummary Resolution value
func (t *IssueSummary) SetResolution(v string) {
	t.Resolution = &v
}

// GetResolutionID will return the IssueSummary ResolutionID value
func (t *IssueSummary) GetResolutionID() string {
	if t.ResolutionID == nil {
		return ""
	}
	return *t.ResolutionID
}

// SetResolutionID will set the IssueSummary ResolutionID value
func (t *IssueSummary) SetResolutionID(v string) {
	t.ResolutionID = &v
}

// FindIssueSummariesByResolutionID will find all IssueSummarys by the ResolutionID value
func FindIssueSummariesByResolutionID(ctx context.Context, db DB, value string) ([]*IssueSummary, error) {
	q := "SELECT `issue_summary`.`id`,`issue_summary`.`checksum`,`issue_summary`.`issue_id`,`issue_summary`.`total_issues`,`issue_summary`.`new30_days`,`issue_summary`.`total_closed`,`issue_summary`.`closed30_days`,`issue_summary`.`estimated_work_months`,`issue_summary`.`estimated_work_months30_days`,`issue_summary`.`title`,`issue_summary`.`url`,`issue_summary`.`priority`,`issue_summary`.`priority_id`,`issue_summary`.`status`,`issue_summary`.`status_id`,`issue_summary`.`issue_type`,`issue_summary`.`issue_type_id`,`issue_summary`.`resolution`,`issue_summary`.`resolution_id`,`issue_summary`.`state`,`issue_summary`.`custom_field_ids`,`issue_summary`.`teams`,`issue_summary`.`parent_issue_id`,`issue_summary`.`parents_issue_ids`,`issue_summary`.`metadata`,`issue_summary`.`project_id`,`issue_summary`.`sprints`,`issue_summary`.`labels`,`issue_summary`.`top_level`,`issue_summary`.`is_leaf`,`issue_summary`.`path`,`issue_summary`.`in_progress_duration`,`issue_summary`.`verification_duration`,`issue_summary`.`in_progress_count`,`issue_summary`.`reopen_count`,`issue_summary`.`mapped_type`,`issue_summary`.`strategic_parent_id`,`issue_summary`.`sprint_id`,`issue_summary`.`issue_project_name`,`issue_summary`.`users`,`issue_summary`.`initial_start_date`,`issue_summary`.`total_duration`,`issue_summary`.`created_at`,`issue_summary`.`closed_at`,`issue_summary`.`planned_end_date`,`issue_summary`.`customer_id`,`issue_summary`.`ref_type`,`issue_summary`.`ref_id`,`issue_summary`.`custom_field_ids_virtual`,`issue_summary`.`release_duration`,`issue_summary`.`completed`,`issue_summary`.`completed_date` FROM `issue_summary` WHERE `resolution_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*IssueSummary, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _IssueID sql.NullString
		var _TotalIssues sql.NullInt64
		var _New30Days sql.NullInt64
		var _TotalClosed sql.NullInt64
		var _Closed30Days sql.NullInt64
		var _EstimatedWorkMonths sql.NullFloat64
		var _EstimatedWorkMonths30Days sql.NullFloat64
		var _Title sql.NullString
		var _URL sql.NullString
		var _Priority sql.NullString
		var _PriorityID sql.NullString
		var _Status sql.NullString
		var _StatusID sql.NullString
		var _IssueType sql.NullString
		var _IssueTypeID sql.NullString
		var _Resolution sql.NullString
		var _ResolutionID sql.NullString
		var _State sql.NullString
		var _CustomFieldIds sql.NullString
		var _Teams sql.NullString
		var _ParentIssueID sql.NullString
		var _ParentsIssueIds sql.NullString
		var _Metadata sql.NullString
		var _ProjectID sql.NullString
		var _Sprints sql.NullString
		var _Labels sql.NullString
		var _TopLevel sql.NullBool
		var _IsLeaf sql.NullBool
		var _Path sql.NullString
		var _InProgressDuration sql.NullInt64
		var _VerificationDuration sql.NullInt64
		var _InProgressCount sql.NullInt64
		var _ReopenCount sql.NullInt64
		var _MappedType sql.NullString
		var _StrategicParentID sql.NullString
		var _SprintID sql.NullString
		var _IssueProjectName sql.NullString
		var _Users sql.NullString
		var _InitialStartDate sql.NullInt64
		var _TotalDuration sql.NullInt64
		var _CreatedAt sql.NullInt64
		var _ClosedAt sql.NullInt64
		var _PlannedEndDate sql.NullInt64
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefID sql.NullString
		var _CustomFieldIdsVirtual sql.NullString
		var _ReleaseDuration sql.NullInt64
		var _Completed sql.NullBool
		var _CompletedDate sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_IssueID,
			&_TotalIssues,
			&_New30Days,
			&_TotalClosed,
			&_Closed30Days,
			&_EstimatedWorkMonths,
			&_EstimatedWorkMonths30Days,
			&_Title,
			&_URL,
			&_Priority,
			&_PriorityID,
			&_Status,
			&_StatusID,
			&_IssueType,
			&_IssueTypeID,
			&_Resolution,
			&_ResolutionID,
			&_State,
			&_CustomFieldIds,
			&_Teams,
			&_ParentIssueID,
			&_ParentsIssueIds,
			&_Metadata,
			&_ProjectID,
			&_Sprints,
			&_Labels,
			&_TopLevel,
			&_IsLeaf,
			&_Path,
			&_InProgressDuration,
			&_VerificationDuration,
			&_InProgressCount,
			&_ReopenCount,
			&_MappedType,
			&_StrategicParentID,
			&_SprintID,
			&_IssueProjectName,
			&_Users,
			&_InitialStartDate,
			&_TotalDuration,
			&_CreatedAt,
			&_ClosedAt,
			&_PlannedEndDate,
			&_CustomerID,
			&_RefType,
			&_RefID,
			&_CustomFieldIdsVirtual,
			&_ReleaseDuration,
			&_Completed,
			&_CompletedDate,
		)
		if err != nil {
			return nil, err
		}
		t := &IssueSummary{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _TotalIssues.Valid {
			t.SetTotalIssues(int32(_TotalIssues.Int64))
		}
		if _New30Days.Valid {
			t.SetNew30Days(int32(_New30Days.Int64))
		}
		if _TotalClosed.Valid {
			t.SetTotalClosed(int32(_TotalClosed.Int64))
		}
		if _Closed30Days.Valid {
			t.SetClosed30Days(int32(_Closed30Days.Int64))
		}
		if _EstimatedWorkMonths.Valid {
			t.SetEstimatedWorkMonths(_EstimatedWorkMonths.Float64)
		}
		if _EstimatedWorkMonths30Days.Valid {
			t.SetEstimatedWorkMonths30Days(_EstimatedWorkMonths30Days.Float64)
		}
		if _Title.Valid {
			t.SetTitle(_Title.String)
		}
		if _URL.Valid {
			t.SetURL(_URL.String)
		}
		if _Priority.Valid {
			t.SetPriority(_Priority.String)
		}
		if _PriorityID.Valid {
			t.SetPriorityID(_PriorityID.String)
		}
		if _Status.Valid {
			t.SetStatus(_Status.String)
		}
		if _StatusID.Valid {
			t.SetStatusID(_StatusID.String)
		}
		if _IssueType.Valid {
			t.SetIssueType(_IssueType.String)
		}
		if _IssueTypeID.Valid {
			t.SetIssueTypeID(_IssueTypeID.String)
		}
		if _Resolution.Valid {
			t.SetResolution(_Resolution.String)
		}
		if _ResolutionID.Valid {
			t.SetResolutionID(_ResolutionID.String)
		}
		if _State.Valid {
			t.SetState(_State.String)
		}
		if _CustomFieldIds.Valid {
			t.SetCustomFieldIds(_CustomFieldIds.String)
		}
		if _Teams.Valid {
			t.SetTeams(_Teams.String)
		}
		if _ParentIssueID.Valid {
			t.SetParentIssueID(_ParentIssueID.String)
		}
		if _ParentsIssueIds.Valid {
			t.SetParentsIssueIds(_ParentsIssueIds.String)
		}
		if _Metadata.Valid {
			t.SetMetadata(_Metadata.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _Sprints.Valid {
			t.SetSprints(_Sprints.String)
		}
		if _Labels.Valid {
			t.SetLabels(_Labels.String)
		}
		if _TopLevel.Valid {
			t.SetTopLevel(_TopLevel.Bool)
		}
		if _IsLeaf.Valid {
			t.SetIsLeaf(_IsLeaf.Bool)
		}
		if _Path.Valid {
			t.SetPath(_Path.String)
		}
		if _InProgressDuration.Valid {
			t.SetInProgressDuration(_InProgressDuration.Int64)
		}
		if _VerificationDuration.Valid {
			t.SetVerificationDuration(_VerificationDuration.Int64)
		}
		if _InProgressCount.Valid {
			t.SetInProgressCount(int32(_InProgressCount.Int64))
		}
		if _ReopenCount.Valid {
			t.SetReopenCount(int32(_ReopenCount.Int64))
		}
		if _MappedType.Valid {
			t.SetMappedType(_MappedType.String)
		}
		if _StrategicParentID.Valid {
			t.SetStrategicParentID(_StrategicParentID.String)
		}
		if _SprintID.Valid {
			t.SetSprintID(_SprintID.String)
		}
		if _IssueProjectName.Valid {
			t.SetIssueProjectName(_IssueProjectName.String)
		}
		if _Users.Valid {
			t.SetUsers(_Users.String)
		}
		if _InitialStartDate.Valid {
			t.SetInitialStartDate(_InitialStartDate.Int64)
		}
		if _TotalDuration.Valid {
			t.SetTotalDuration(_TotalDuration.Int64)
		}
		if _CreatedAt.Valid {
			t.SetCreatedAt(_CreatedAt.Int64)
		}
		if _ClosedAt.Valid {
			t.SetClosedAt(_ClosedAt.Int64)
		}
		if _PlannedEndDate.Valid {
			t.SetPlannedEndDate(_PlannedEndDate.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefType.Valid {
			t.SetRefType(_RefType.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		if _CustomFieldIdsVirtual.Valid {
			t.SetCustomFieldIdsVirtual(_CustomFieldIdsVirtual.String)
		}
		if _ReleaseDuration.Valid {
			t.SetReleaseDuration(_ReleaseDuration.Int64)
		}
		if _Completed.Valid {
			t.SetCompleted(_Completed.Bool)
		}
		if _CompletedDate.Valid {
			t.SetCompletedDate(_CompletedDate.Int64)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindIssueSummariesByResolutionIDTx will find all IssueSummarys by the ResolutionID value using the provided transaction
func FindIssueSummariesByResolutionIDTx(ctx context.Context, tx Tx, value string) ([]*IssueSummary, error) {
	q := "SELECT `issue_summary`.`id`,`issue_summary`.`checksum`,`issue_summary`.`issue_id`,`issue_summary`.`total_issues`,`issue_summary`.`new30_days`,`issue_summary`.`total_closed`,`issue_summary`.`closed30_days`,`issue_summary`.`estimated_work_months`,`issue_summary`.`estimated_work_months30_days`,`issue_summary`.`title`,`issue_summary`.`url`,`issue_summary`.`priority`,`issue_summary`.`priority_id`,`issue_summary`.`status`,`issue_summary`.`status_id`,`issue_summary`.`issue_type`,`issue_summary`.`issue_type_id`,`issue_summary`.`resolution`,`issue_summary`.`resolution_id`,`issue_summary`.`state`,`issue_summary`.`custom_field_ids`,`issue_summary`.`teams`,`issue_summary`.`parent_issue_id`,`issue_summary`.`parents_issue_ids`,`issue_summary`.`metadata`,`issue_summary`.`project_id`,`issue_summary`.`sprints`,`issue_summary`.`labels`,`issue_summary`.`top_level`,`issue_summary`.`is_leaf`,`issue_summary`.`path`,`issue_summary`.`in_progress_duration`,`issue_summary`.`verification_duration`,`issue_summary`.`in_progress_count`,`issue_summary`.`reopen_count`,`issue_summary`.`mapped_type`,`issue_summary`.`strategic_parent_id`,`issue_summary`.`sprint_id`,`issue_summary`.`issue_project_name`,`issue_summary`.`users`,`issue_summary`.`initial_start_date`,`issue_summary`.`total_duration`,`issue_summary`.`created_at`,`issue_summary`.`closed_at`,`issue_summary`.`planned_end_date`,`issue_summary`.`customer_id`,`issue_summary`.`ref_type`,`issue_summary`.`ref_id`,`issue_summary`.`custom_field_ids_virtual`,`issue_summary`.`release_duration`,`issue_summary`.`completed`,`issue_summary`.`completed_date` FROM `issue_summary` WHERE `resolution_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*IssueSummary, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _IssueID sql.NullString
		var _TotalIssues sql.NullInt64
		var _New30Days sql.NullInt64
		var _TotalClosed sql.NullInt64
		var _Closed30Days sql.NullInt64
		var _EstimatedWorkMonths sql.NullFloat64
		var _EstimatedWorkMonths30Days sql.NullFloat64
		var _Title sql.NullString
		var _URL sql.NullString
		var _Priority sql.NullString
		var _PriorityID sql.NullString
		var _Status sql.NullString
		var _StatusID sql.NullString
		var _IssueType sql.NullString
		var _IssueTypeID sql.NullString
		var _Resolution sql.NullString
		var _ResolutionID sql.NullString
		var _State sql.NullString
		var _CustomFieldIds sql.NullString
		var _Teams sql.NullString
		var _ParentIssueID sql.NullString
		var _ParentsIssueIds sql.NullString
		var _Metadata sql.NullString
		var _ProjectID sql.NullString
		var _Sprints sql.NullString
		var _Labels sql.NullString
		var _TopLevel sql.NullBool
		var _IsLeaf sql.NullBool
		var _Path sql.NullString
		var _InProgressDuration sql.NullInt64
		var _VerificationDuration sql.NullInt64
		var _InProgressCount sql.NullInt64
		var _ReopenCount sql.NullInt64
		var _MappedType sql.NullString
		var _StrategicParentID sql.NullString
		var _SprintID sql.NullString
		var _IssueProjectName sql.NullString
		var _Users sql.NullString
		var _InitialStartDate sql.NullInt64
		var _TotalDuration sql.NullInt64
		var _CreatedAt sql.NullInt64
		var _ClosedAt sql.NullInt64
		var _PlannedEndDate sql.NullInt64
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefID sql.NullString
		var _CustomFieldIdsVirtual sql.NullString
		var _ReleaseDuration sql.NullInt64
		var _Completed sql.NullBool
		var _CompletedDate sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_IssueID,
			&_TotalIssues,
			&_New30Days,
			&_TotalClosed,
			&_Closed30Days,
			&_EstimatedWorkMonths,
			&_EstimatedWorkMonths30Days,
			&_Title,
			&_URL,
			&_Priority,
			&_PriorityID,
			&_Status,
			&_StatusID,
			&_IssueType,
			&_IssueTypeID,
			&_Resolution,
			&_ResolutionID,
			&_State,
			&_CustomFieldIds,
			&_Teams,
			&_ParentIssueID,
			&_ParentsIssueIds,
			&_Metadata,
			&_ProjectID,
			&_Sprints,
			&_Labels,
			&_TopLevel,
			&_IsLeaf,
			&_Path,
			&_InProgressDuration,
			&_VerificationDuration,
			&_InProgressCount,
			&_ReopenCount,
			&_MappedType,
			&_StrategicParentID,
			&_SprintID,
			&_IssueProjectName,
			&_Users,
			&_InitialStartDate,
			&_TotalDuration,
			&_CreatedAt,
			&_ClosedAt,
			&_PlannedEndDate,
			&_CustomerID,
			&_RefType,
			&_RefID,
			&_CustomFieldIdsVirtual,
			&_ReleaseDuration,
			&_Completed,
			&_CompletedDate,
		)
		if err != nil {
			return nil, err
		}
		t := &IssueSummary{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _TotalIssues.Valid {
			t.SetTotalIssues(int32(_TotalIssues.Int64))
		}
		if _New30Days.Valid {
			t.SetNew30Days(int32(_New30Days.Int64))
		}
		if _TotalClosed.Valid {
			t.SetTotalClosed(int32(_TotalClosed.Int64))
		}
		if _Closed30Days.Valid {
			t.SetClosed30Days(int32(_Closed30Days.Int64))
		}
		if _EstimatedWorkMonths.Valid {
			t.SetEstimatedWorkMonths(_EstimatedWorkMonths.Float64)
		}
		if _EstimatedWorkMonths30Days.Valid {
			t.SetEstimatedWorkMonths30Days(_EstimatedWorkMonths30Days.Float64)
		}
		if _Title.Valid {
			t.SetTitle(_Title.String)
		}
		if _URL.Valid {
			t.SetURL(_URL.String)
		}
		if _Priority.Valid {
			t.SetPriority(_Priority.String)
		}
		if _PriorityID.Valid {
			t.SetPriorityID(_PriorityID.String)
		}
		if _Status.Valid {
			t.SetStatus(_Status.String)
		}
		if _StatusID.Valid {
			t.SetStatusID(_StatusID.String)
		}
		if _IssueType.Valid {
			t.SetIssueType(_IssueType.String)
		}
		if _IssueTypeID.Valid {
			t.SetIssueTypeID(_IssueTypeID.String)
		}
		if _Resolution.Valid {
			t.SetResolution(_Resolution.String)
		}
		if _ResolutionID.Valid {
			t.SetResolutionID(_ResolutionID.String)
		}
		if _State.Valid {
			t.SetState(_State.String)
		}
		if _CustomFieldIds.Valid {
			t.SetCustomFieldIds(_CustomFieldIds.String)
		}
		if _Teams.Valid {
			t.SetTeams(_Teams.String)
		}
		if _ParentIssueID.Valid {
			t.SetParentIssueID(_ParentIssueID.String)
		}
		if _ParentsIssueIds.Valid {
			t.SetParentsIssueIds(_ParentsIssueIds.String)
		}
		if _Metadata.Valid {
			t.SetMetadata(_Metadata.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _Sprints.Valid {
			t.SetSprints(_Sprints.String)
		}
		if _Labels.Valid {
			t.SetLabels(_Labels.String)
		}
		if _TopLevel.Valid {
			t.SetTopLevel(_TopLevel.Bool)
		}
		if _IsLeaf.Valid {
			t.SetIsLeaf(_IsLeaf.Bool)
		}
		if _Path.Valid {
			t.SetPath(_Path.String)
		}
		if _InProgressDuration.Valid {
			t.SetInProgressDuration(_InProgressDuration.Int64)
		}
		if _VerificationDuration.Valid {
			t.SetVerificationDuration(_VerificationDuration.Int64)
		}
		if _InProgressCount.Valid {
			t.SetInProgressCount(int32(_InProgressCount.Int64))
		}
		if _ReopenCount.Valid {
			t.SetReopenCount(int32(_ReopenCount.Int64))
		}
		if _MappedType.Valid {
			t.SetMappedType(_MappedType.String)
		}
		if _StrategicParentID.Valid {
			t.SetStrategicParentID(_StrategicParentID.String)
		}
		if _SprintID.Valid {
			t.SetSprintID(_SprintID.String)
		}
		if _IssueProjectName.Valid {
			t.SetIssueProjectName(_IssueProjectName.String)
		}
		if _Users.Valid {
			t.SetUsers(_Users.String)
		}
		if _InitialStartDate.Valid {
			t.SetInitialStartDate(_InitialStartDate.Int64)
		}
		if _TotalDuration.Valid {
			t.SetTotalDuration(_TotalDuration.Int64)
		}
		if _CreatedAt.Valid {
			t.SetCreatedAt(_CreatedAt.Int64)
		}
		if _ClosedAt.Valid {
			t.SetClosedAt(_ClosedAt.Int64)
		}
		if _PlannedEndDate.Valid {
			t.SetPlannedEndDate(_PlannedEndDate.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefType.Valid {
			t.SetRefType(_RefType.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		if _CustomFieldIdsVirtual.Valid {
			t.SetCustomFieldIdsVirtual(_CustomFieldIdsVirtual.String)
		}
		if _ReleaseDuration.Valid {
			t.SetReleaseDuration(_ReleaseDuration.Int64)
		}
		if _Completed.Valid {
			t.SetCompleted(_Completed.Bool)
		}
		if _CompletedDate.Valid {
			t.SetCompletedDate(_CompletedDate.Int64)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetState will return the IssueSummary State value
func (t *IssueSummary) GetState() string {
	return t.State
}

// SetState will set the IssueSummary State value
func (t *IssueSummary) SetState(v string) {
	t.State = v
}

// GetCustomFieldIds will return the IssueSummary CustomFieldIds value
func (t *IssueSummary) GetCustomFieldIds() string {
	if t.CustomFieldIds == nil {
		return ""
	}
	return *t.CustomFieldIds
}

// SetCustomFieldIds will set the IssueSummary CustomFieldIds value
func (t *IssueSummary) SetCustomFieldIds(v string) {
	t.CustomFieldIds = &v
}

// GetTeams will return the IssueSummary Teams value
func (t *IssueSummary) GetTeams() string {
	if t.Teams == nil {
		return ""
	}
	return *t.Teams
}

// SetTeams will set the IssueSummary Teams value
func (t *IssueSummary) SetTeams(v string) {
	t.Teams = &v
}

// GetParentIssueID will return the IssueSummary ParentIssueID value
func (t *IssueSummary) GetParentIssueID() string {
	if t.ParentIssueID == nil {
		return ""
	}
	return *t.ParentIssueID
}

// SetParentIssueID will set the IssueSummary ParentIssueID value
func (t *IssueSummary) SetParentIssueID(v string) {
	t.ParentIssueID = &v
}

// FindIssueSummariesByParentIssueID will find all IssueSummarys by the ParentIssueID value
func FindIssueSummariesByParentIssueID(ctx context.Context, db DB, value string) ([]*IssueSummary, error) {
	q := "SELECT `issue_summary`.`id`,`issue_summary`.`checksum`,`issue_summary`.`issue_id`,`issue_summary`.`total_issues`,`issue_summary`.`new30_days`,`issue_summary`.`total_closed`,`issue_summary`.`closed30_days`,`issue_summary`.`estimated_work_months`,`issue_summary`.`estimated_work_months30_days`,`issue_summary`.`title`,`issue_summary`.`url`,`issue_summary`.`priority`,`issue_summary`.`priority_id`,`issue_summary`.`status`,`issue_summary`.`status_id`,`issue_summary`.`issue_type`,`issue_summary`.`issue_type_id`,`issue_summary`.`resolution`,`issue_summary`.`resolution_id`,`issue_summary`.`state`,`issue_summary`.`custom_field_ids`,`issue_summary`.`teams`,`issue_summary`.`parent_issue_id`,`issue_summary`.`parents_issue_ids`,`issue_summary`.`metadata`,`issue_summary`.`project_id`,`issue_summary`.`sprints`,`issue_summary`.`labels`,`issue_summary`.`top_level`,`issue_summary`.`is_leaf`,`issue_summary`.`path`,`issue_summary`.`in_progress_duration`,`issue_summary`.`verification_duration`,`issue_summary`.`in_progress_count`,`issue_summary`.`reopen_count`,`issue_summary`.`mapped_type`,`issue_summary`.`strategic_parent_id`,`issue_summary`.`sprint_id`,`issue_summary`.`issue_project_name`,`issue_summary`.`users`,`issue_summary`.`initial_start_date`,`issue_summary`.`total_duration`,`issue_summary`.`created_at`,`issue_summary`.`closed_at`,`issue_summary`.`planned_end_date`,`issue_summary`.`customer_id`,`issue_summary`.`ref_type`,`issue_summary`.`ref_id`,`issue_summary`.`custom_field_ids_virtual`,`issue_summary`.`release_duration`,`issue_summary`.`completed`,`issue_summary`.`completed_date` FROM `issue_summary` WHERE `parent_issue_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*IssueSummary, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _IssueID sql.NullString
		var _TotalIssues sql.NullInt64
		var _New30Days sql.NullInt64
		var _TotalClosed sql.NullInt64
		var _Closed30Days sql.NullInt64
		var _EstimatedWorkMonths sql.NullFloat64
		var _EstimatedWorkMonths30Days sql.NullFloat64
		var _Title sql.NullString
		var _URL sql.NullString
		var _Priority sql.NullString
		var _PriorityID sql.NullString
		var _Status sql.NullString
		var _StatusID sql.NullString
		var _IssueType sql.NullString
		var _IssueTypeID sql.NullString
		var _Resolution sql.NullString
		var _ResolutionID sql.NullString
		var _State sql.NullString
		var _CustomFieldIds sql.NullString
		var _Teams sql.NullString
		var _ParentIssueID sql.NullString
		var _ParentsIssueIds sql.NullString
		var _Metadata sql.NullString
		var _ProjectID sql.NullString
		var _Sprints sql.NullString
		var _Labels sql.NullString
		var _TopLevel sql.NullBool
		var _IsLeaf sql.NullBool
		var _Path sql.NullString
		var _InProgressDuration sql.NullInt64
		var _VerificationDuration sql.NullInt64
		var _InProgressCount sql.NullInt64
		var _ReopenCount sql.NullInt64
		var _MappedType sql.NullString
		var _StrategicParentID sql.NullString
		var _SprintID sql.NullString
		var _IssueProjectName sql.NullString
		var _Users sql.NullString
		var _InitialStartDate sql.NullInt64
		var _TotalDuration sql.NullInt64
		var _CreatedAt sql.NullInt64
		var _ClosedAt sql.NullInt64
		var _PlannedEndDate sql.NullInt64
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefID sql.NullString
		var _CustomFieldIdsVirtual sql.NullString
		var _ReleaseDuration sql.NullInt64
		var _Completed sql.NullBool
		var _CompletedDate sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_IssueID,
			&_TotalIssues,
			&_New30Days,
			&_TotalClosed,
			&_Closed30Days,
			&_EstimatedWorkMonths,
			&_EstimatedWorkMonths30Days,
			&_Title,
			&_URL,
			&_Priority,
			&_PriorityID,
			&_Status,
			&_StatusID,
			&_IssueType,
			&_IssueTypeID,
			&_Resolution,
			&_ResolutionID,
			&_State,
			&_CustomFieldIds,
			&_Teams,
			&_ParentIssueID,
			&_ParentsIssueIds,
			&_Metadata,
			&_ProjectID,
			&_Sprints,
			&_Labels,
			&_TopLevel,
			&_IsLeaf,
			&_Path,
			&_InProgressDuration,
			&_VerificationDuration,
			&_InProgressCount,
			&_ReopenCount,
			&_MappedType,
			&_StrategicParentID,
			&_SprintID,
			&_IssueProjectName,
			&_Users,
			&_InitialStartDate,
			&_TotalDuration,
			&_CreatedAt,
			&_ClosedAt,
			&_PlannedEndDate,
			&_CustomerID,
			&_RefType,
			&_RefID,
			&_CustomFieldIdsVirtual,
			&_ReleaseDuration,
			&_Completed,
			&_CompletedDate,
		)
		if err != nil {
			return nil, err
		}
		t := &IssueSummary{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _TotalIssues.Valid {
			t.SetTotalIssues(int32(_TotalIssues.Int64))
		}
		if _New30Days.Valid {
			t.SetNew30Days(int32(_New30Days.Int64))
		}
		if _TotalClosed.Valid {
			t.SetTotalClosed(int32(_TotalClosed.Int64))
		}
		if _Closed30Days.Valid {
			t.SetClosed30Days(int32(_Closed30Days.Int64))
		}
		if _EstimatedWorkMonths.Valid {
			t.SetEstimatedWorkMonths(_EstimatedWorkMonths.Float64)
		}
		if _EstimatedWorkMonths30Days.Valid {
			t.SetEstimatedWorkMonths30Days(_EstimatedWorkMonths30Days.Float64)
		}
		if _Title.Valid {
			t.SetTitle(_Title.String)
		}
		if _URL.Valid {
			t.SetURL(_URL.String)
		}
		if _Priority.Valid {
			t.SetPriority(_Priority.String)
		}
		if _PriorityID.Valid {
			t.SetPriorityID(_PriorityID.String)
		}
		if _Status.Valid {
			t.SetStatus(_Status.String)
		}
		if _StatusID.Valid {
			t.SetStatusID(_StatusID.String)
		}
		if _IssueType.Valid {
			t.SetIssueType(_IssueType.String)
		}
		if _IssueTypeID.Valid {
			t.SetIssueTypeID(_IssueTypeID.String)
		}
		if _Resolution.Valid {
			t.SetResolution(_Resolution.String)
		}
		if _ResolutionID.Valid {
			t.SetResolutionID(_ResolutionID.String)
		}
		if _State.Valid {
			t.SetState(_State.String)
		}
		if _CustomFieldIds.Valid {
			t.SetCustomFieldIds(_CustomFieldIds.String)
		}
		if _Teams.Valid {
			t.SetTeams(_Teams.String)
		}
		if _ParentIssueID.Valid {
			t.SetParentIssueID(_ParentIssueID.String)
		}
		if _ParentsIssueIds.Valid {
			t.SetParentsIssueIds(_ParentsIssueIds.String)
		}
		if _Metadata.Valid {
			t.SetMetadata(_Metadata.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _Sprints.Valid {
			t.SetSprints(_Sprints.String)
		}
		if _Labels.Valid {
			t.SetLabels(_Labels.String)
		}
		if _TopLevel.Valid {
			t.SetTopLevel(_TopLevel.Bool)
		}
		if _IsLeaf.Valid {
			t.SetIsLeaf(_IsLeaf.Bool)
		}
		if _Path.Valid {
			t.SetPath(_Path.String)
		}
		if _InProgressDuration.Valid {
			t.SetInProgressDuration(_InProgressDuration.Int64)
		}
		if _VerificationDuration.Valid {
			t.SetVerificationDuration(_VerificationDuration.Int64)
		}
		if _InProgressCount.Valid {
			t.SetInProgressCount(int32(_InProgressCount.Int64))
		}
		if _ReopenCount.Valid {
			t.SetReopenCount(int32(_ReopenCount.Int64))
		}
		if _MappedType.Valid {
			t.SetMappedType(_MappedType.String)
		}
		if _StrategicParentID.Valid {
			t.SetStrategicParentID(_StrategicParentID.String)
		}
		if _SprintID.Valid {
			t.SetSprintID(_SprintID.String)
		}
		if _IssueProjectName.Valid {
			t.SetIssueProjectName(_IssueProjectName.String)
		}
		if _Users.Valid {
			t.SetUsers(_Users.String)
		}
		if _InitialStartDate.Valid {
			t.SetInitialStartDate(_InitialStartDate.Int64)
		}
		if _TotalDuration.Valid {
			t.SetTotalDuration(_TotalDuration.Int64)
		}
		if _CreatedAt.Valid {
			t.SetCreatedAt(_CreatedAt.Int64)
		}
		if _ClosedAt.Valid {
			t.SetClosedAt(_ClosedAt.Int64)
		}
		if _PlannedEndDate.Valid {
			t.SetPlannedEndDate(_PlannedEndDate.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefType.Valid {
			t.SetRefType(_RefType.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		if _CustomFieldIdsVirtual.Valid {
			t.SetCustomFieldIdsVirtual(_CustomFieldIdsVirtual.String)
		}
		if _ReleaseDuration.Valid {
			t.SetReleaseDuration(_ReleaseDuration.Int64)
		}
		if _Completed.Valid {
			t.SetCompleted(_Completed.Bool)
		}
		if _CompletedDate.Valid {
			t.SetCompletedDate(_CompletedDate.Int64)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindIssueSummariesByParentIssueIDTx will find all IssueSummarys by the ParentIssueID value using the provided transaction
func FindIssueSummariesByParentIssueIDTx(ctx context.Context, tx Tx, value string) ([]*IssueSummary, error) {
	q := "SELECT `issue_summary`.`id`,`issue_summary`.`checksum`,`issue_summary`.`issue_id`,`issue_summary`.`total_issues`,`issue_summary`.`new30_days`,`issue_summary`.`total_closed`,`issue_summary`.`closed30_days`,`issue_summary`.`estimated_work_months`,`issue_summary`.`estimated_work_months30_days`,`issue_summary`.`title`,`issue_summary`.`url`,`issue_summary`.`priority`,`issue_summary`.`priority_id`,`issue_summary`.`status`,`issue_summary`.`status_id`,`issue_summary`.`issue_type`,`issue_summary`.`issue_type_id`,`issue_summary`.`resolution`,`issue_summary`.`resolution_id`,`issue_summary`.`state`,`issue_summary`.`custom_field_ids`,`issue_summary`.`teams`,`issue_summary`.`parent_issue_id`,`issue_summary`.`parents_issue_ids`,`issue_summary`.`metadata`,`issue_summary`.`project_id`,`issue_summary`.`sprints`,`issue_summary`.`labels`,`issue_summary`.`top_level`,`issue_summary`.`is_leaf`,`issue_summary`.`path`,`issue_summary`.`in_progress_duration`,`issue_summary`.`verification_duration`,`issue_summary`.`in_progress_count`,`issue_summary`.`reopen_count`,`issue_summary`.`mapped_type`,`issue_summary`.`strategic_parent_id`,`issue_summary`.`sprint_id`,`issue_summary`.`issue_project_name`,`issue_summary`.`users`,`issue_summary`.`initial_start_date`,`issue_summary`.`total_duration`,`issue_summary`.`created_at`,`issue_summary`.`closed_at`,`issue_summary`.`planned_end_date`,`issue_summary`.`customer_id`,`issue_summary`.`ref_type`,`issue_summary`.`ref_id`,`issue_summary`.`custom_field_ids_virtual`,`issue_summary`.`release_duration`,`issue_summary`.`completed`,`issue_summary`.`completed_date` FROM `issue_summary` WHERE `parent_issue_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*IssueSummary, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _IssueID sql.NullString
		var _TotalIssues sql.NullInt64
		var _New30Days sql.NullInt64
		var _TotalClosed sql.NullInt64
		var _Closed30Days sql.NullInt64
		var _EstimatedWorkMonths sql.NullFloat64
		var _EstimatedWorkMonths30Days sql.NullFloat64
		var _Title sql.NullString
		var _URL sql.NullString
		var _Priority sql.NullString
		var _PriorityID sql.NullString
		var _Status sql.NullString
		var _StatusID sql.NullString
		var _IssueType sql.NullString
		var _IssueTypeID sql.NullString
		var _Resolution sql.NullString
		var _ResolutionID sql.NullString
		var _State sql.NullString
		var _CustomFieldIds sql.NullString
		var _Teams sql.NullString
		var _ParentIssueID sql.NullString
		var _ParentsIssueIds sql.NullString
		var _Metadata sql.NullString
		var _ProjectID sql.NullString
		var _Sprints sql.NullString
		var _Labels sql.NullString
		var _TopLevel sql.NullBool
		var _IsLeaf sql.NullBool
		var _Path sql.NullString
		var _InProgressDuration sql.NullInt64
		var _VerificationDuration sql.NullInt64
		var _InProgressCount sql.NullInt64
		var _ReopenCount sql.NullInt64
		var _MappedType sql.NullString
		var _StrategicParentID sql.NullString
		var _SprintID sql.NullString
		var _IssueProjectName sql.NullString
		var _Users sql.NullString
		var _InitialStartDate sql.NullInt64
		var _TotalDuration sql.NullInt64
		var _CreatedAt sql.NullInt64
		var _ClosedAt sql.NullInt64
		var _PlannedEndDate sql.NullInt64
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefID sql.NullString
		var _CustomFieldIdsVirtual sql.NullString
		var _ReleaseDuration sql.NullInt64
		var _Completed sql.NullBool
		var _CompletedDate sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_IssueID,
			&_TotalIssues,
			&_New30Days,
			&_TotalClosed,
			&_Closed30Days,
			&_EstimatedWorkMonths,
			&_EstimatedWorkMonths30Days,
			&_Title,
			&_URL,
			&_Priority,
			&_PriorityID,
			&_Status,
			&_StatusID,
			&_IssueType,
			&_IssueTypeID,
			&_Resolution,
			&_ResolutionID,
			&_State,
			&_CustomFieldIds,
			&_Teams,
			&_ParentIssueID,
			&_ParentsIssueIds,
			&_Metadata,
			&_ProjectID,
			&_Sprints,
			&_Labels,
			&_TopLevel,
			&_IsLeaf,
			&_Path,
			&_InProgressDuration,
			&_VerificationDuration,
			&_InProgressCount,
			&_ReopenCount,
			&_MappedType,
			&_StrategicParentID,
			&_SprintID,
			&_IssueProjectName,
			&_Users,
			&_InitialStartDate,
			&_TotalDuration,
			&_CreatedAt,
			&_ClosedAt,
			&_PlannedEndDate,
			&_CustomerID,
			&_RefType,
			&_RefID,
			&_CustomFieldIdsVirtual,
			&_ReleaseDuration,
			&_Completed,
			&_CompletedDate,
		)
		if err != nil {
			return nil, err
		}
		t := &IssueSummary{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _TotalIssues.Valid {
			t.SetTotalIssues(int32(_TotalIssues.Int64))
		}
		if _New30Days.Valid {
			t.SetNew30Days(int32(_New30Days.Int64))
		}
		if _TotalClosed.Valid {
			t.SetTotalClosed(int32(_TotalClosed.Int64))
		}
		if _Closed30Days.Valid {
			t.SetClosed30Days(int32(_Closed30Days.Int64))
		}
		if _EstimatedWorkMonths.Valid {
			t.SetEstimatedWorkMonths(_EstimatedWorkMonths.Float64)
		}
		if _EstimatedWorkMonths30Days.Valid {
			t.SetEstimatedWorkMonths30Days(_EstimatedWorkMonths30Days.Float64)
		}
		if _Title.Valid {
			t.SetTitle(_Title.String)
		}
		if _URL.Valid {
			t.SetURL(_URL.String)
		}
		if _Priority.Valid {
			t.SetPriority(_Priority.String)
		}
		if _PriorityID.Valid {
			t.SetPriorityID(_PriorityID.String)
		}
		if _Status.Valid {
			t.SetStatus(_Status.String)
		}
		if _StatusID.Valid {
			t.SetStatusID(_StatusID.String)
		}
		if _IssueType.Valid {
			t.SetIssueType(_IssueType.String)
		}
		if _IssueTypeID.Valid {
			t.SetIssueTypeID(_IssueTypeID.String)
		}
		if _Resolution.Valid {
			t.SetResolution(_Resolution.String)
		}
		if _ResolutionID.Valid {
			t.SetResolutionID(_ResolutionID.String)
		}
		if _State.Valid {
			t.SetState(_State.String)
		}
		if _CustomFieldIds.Valid {
			t.SetCustomFieldIds(_CustomFieldIds.String)
		}
		if _Teams.Valid {
			t.SetTeams(_Teams.String)
		}
		if _ParentIssueID.Valid {
			t.SetParentIssueID(_ParentIssueID.String)
		}
		if _ParentsIssueIds.Valid {
			t.SetParentsIssueIds(_ParentsIssueIds.String)
		}
		if _Metadata.Valid {
			t.SetMetadata(_Metadata.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _Sprints.Valid {
			t.SetSprints(_Sprints.String)
		}
		if _Labels.Valid {
			t.SetLabels(_Labels.String)
		}
		if _TopLevel.Valid {
			t.SetTopLevel(_TopLevel.Bool)
		}
		if _IsLeaf.Valid {
			t.SetIsLeaf(_IsLeaf.Bool)
		}
		if _Path.Valid {
			t.SetPath(_Path.String)
		}
		if _InProgressDuration.Valid {
			t.SetInProgressDuration(_InProgressDuration.Int64)
		}
		if _VerificationDuration.Valid {
			t.SetVerificationDuration(_VerificationDuration.Int64)
		}
		if _InProgressCount.Valid {
			t.SetInProgressCount(int32(_InProgressCount.Int64))
		}
		if _ReopenCount.Valid {
			t.SetReopenCount(int32(_ReopenCount.Int64))
		}
		if _MappedType.Valid {
			t.SetMappedType(_MappedType.String)
		}
		if _StrategicParentID.Valid {
			t.SetStrategicParentID(_StrategicParentID.String)
		}
		if _SprintID.Valid {
			t.SetSprintID(_SprintID.String)
		}
		if _IssueProjectName.Valid {
			t.SetIssueProjectName(_IssueProjectName.String)
		}
		if _Users.Valid {
			t.SetUsers(_Users.String)
		}
		if _InitialStartDate.Valid {
			t.SetInitialStartDate(_InitialStartDate.Int64)
		}
		if _TotalDuration.Valid {
			t.SetTotalDuration(_TotalDuration.Int64)
		}
		if _CreatedAt.Valid {
			t.SetCreatedAt(_CreatedAt.Int64)
		}
		if _ClosedAt.Valid {
			t.SetClosedAt(_ClosedAt.Int64)
		}
		if _PlannedEndDate.Valid {
			t.SetPlannedEndDate(_PlannedEndDate.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefType.Valid {
			t.SetRefType(_RefType.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		if _CustomFieldIdsVirtual.Valid {
			t.SetCustomFieldIdsVirtual(_CustomFieldIdsVirtual.String)
		}
		if _ReleaseDuration.Valid {
			t.SetReleaseDuration(_ReleaseDuration.Int64)
		}
		if _Completed.Valid {
			t.SetCompleted(_Completed.Bool)
		}
		if _CompletedDate.Valid {
			t.SetCompletedDate(_CompletedDate.Int64)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetParentsIssueIds will return the IssueSummary ParentsIssueIds value
func (t *IssueSummary) GetParentsIssueIds() string {
	if t.ParentsIssueIds == nil {
		return ""
	}
	return *t.ParentsIssueIds
}

// SetParentsIssueIds will set the IssueSummary ParentsIssueIds value
func (t *IssueSummary) SetParentsIssueIds(v string) {
	t.ParentsIssueIds = &v
}

// GetMetadata will return the IssueSummary Metadata value
func (t *IssueSummary) GetMetadata() string {
	if t.Metadata == nil {
		return ""
	}
	return *t.Metadata
}

// SetMetadata will set the IssueSummary Metadata value
func (t *IssueSummary) SetMetadata(v string) {
	t.Metadata = &v
}

// GetProjectID will return the IssueSummary ProjectID value
func (t *IssueSummary) GetProjectID() string {
	return t.ProjectID
}

// SetProjectID will set the IssueSummary ProjectID value
func (t *IssueSummary) SetProjectID(v string) {
	t.ProjectID = v
}

// FindIssueSummariesByProjectID will find all IssueSummarys by the ProjectID value
func FindIssueSummariesByProjectID(ctx context.Context, db DB, value string) ([]*IssueSummary, error) {
	q := "SELECT `issue_summary`.`id`,`issue_summary`.`checksum`,`issue_summary`.`issue_id`,`issue_summary`.`total_issues`,`issue_summary`.`new30_days`,`issue_summary`.`total_closed`,`issue_summary`.`closed30_days`,`issue_summary`.`estimated_work_months`,`issue_summary`.`estimated_work_months30_days`,`issue_summary`.`title`,`issue_summary`.`url`,`issue_summary`.`priority`,`issue_summary`.`priority_id`,`issue_summary`.`status`,`issue_summary`.`status_id`,`issue_summary`.`issue_type`,`issue_summary`.`issue_type_id`,`issue_summary`.`resolution`,`issue_summary`.`resolution_id`,`issue_summary`.`state`,`issue_summary`.`custom_field_ids`,`issue_summary`.`teams`,`issue_summary`.`parent_issue_id`,`issue_summary`.`parents_issue_ids`,`issue_summary`.`metadata`,`issue_summary`.`project_id`,`issue_summary`.`sprints`,`issue_summary`.`labels`,`issue_summary`.`top_level`,`issue_summary`.`is_leaf`,`issue_summary`.`path`,`issue_summary`.`in_progress_duration`,`issue_summary`.`verification_duration`,`issue_summary`.`in_progress_count`,`issue_summary`.`reopen_count`,`issue_summary`.`mapped_type`,`issue_summary`.`strategic_parent_id`,`issue_summary`.`sprint_id`,`issue_summary`.`issue_project_name`,`issue_summary`.`users`,`issue_summary`.`initial_start_date`,`issue_summary`.`total_duration`,`issue_summary`.`created_at`,`issue_summary`.`closed_at`,`issue_summary`.`planned_end_date`,`issue_summary`.`customer_id`,`issue_summary`.`ref_type`,`issue_summary`.`ref_id`,`issue_summary`.`custom_field_ids_virtual`,`issue_summary`.`release_duration`,`issue_summary`.`completed`,`issue_summary`.`completed_date` FROM `issue_summary` WHERE `project_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*IssueSummary, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _IssueID sql.NullString
		var _TotalIssues sql.NullInt64
		var _New30Days sql.NullInt64
		var _TotalClosed sql.NullInt64
		var _Closed30Days sql.NullInt64
		var _EstimatedWorkMonths sql.NullFloat64
		var _EstimatedWorkMonths30Days sql.NullFloat64
		var _Title sql.NullString
		var _URL sql.NullString
		var _Priority sql.NullString
		var _PriorityID sql.NullString
		var _Status sql.NullString
		var _StatusID sql.NullString
		var _IssueType sql.NullString
		var _IssueTypeID sql.NullString
		var _Resolution sql.NullString
		var _ResolutionID sql.NullString
		var _State sql.NullString
		var _CustomFieldIds sql.NullString
		var _Teams sql.NullString
		var _ParentIssueID sql.NullString
		var _ParentsIssueIds sql.NullString
		var _Metadata sql.NullString
		var _ProjectID sql.NullString
		var _Sprints sql.NullString
		var _Labels sql.NullString
		var _TopLevel sql.NullBool
		var _IsLeaf sql.NullBool
		var _Path sql.NullString
		var _InProgressDuration sql.NullInt64
		var _VerificationDuration sql.NullInt64
		var _InProgressCount sql.NullInt64
		var _ReopenCount sql.NullInt64
		var _MappedType sql.NullString
		var _StrategicParentID sql.NullString
		var _SprintID sql.NullString
		var _IssueProjectName sql.NullString
		var _Users sql.NullString
		var _InitialStartDate sql.NullInt64
		var _TotalDuration sql.NullInt64
		var _CreatedAt sql.NullInt64
		var _ClosedAt sql.NullInt64
		var _PlannedEndDate sql.NullInt64
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefID sql.NullString
		var _CustomFieldIdsVirtual sql.NullString
		var _ReleaseDuration sql.NullInt64
		var _Completed sql.NullBool
		var _CompletedDate sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_IssueID,
			&_TotalIssues,
			&_New30Days,
			&_TotalClosed,
			&_Closed30Days,
			&_EstimatedWorkMonths,
			&_EstimatedWorkMonths30Days,
			&_Title,
			&_URL,
			&_Priority,
			&_PriorityID,
			&_Status,
			&_StatusID,
			&_IssueType,
			&_IssueTypeID,
			&_Resolution,
			&_ResolutionID,
			&_State,
			&_CustomFieldIds,
			&_Teams,
			&_ParentIssueID,
			&_ParentsIssueIds,
			&_Metadata,
			&_ProjectID,
			&_Sprints,
			&_Labels,
			&_TopLevel,
			&_IsLeaf,
			&_Path,
			&_InProgressDuration,
			&_VerificationDuration,
			&_InProgressCount,
			&_ReopenCount,
			&_MappedType,
			&_StrategicParentID,
			&_SprintID,
			&_IssueProjectName,
			&_Users,
			&_InitialStartDate,
			&_TotalDuration,
			&_CreatedAt,
			&_ClosedAt,
			&_PlannedEndDate,
			&_CustomerID,
			&_RefType,
			&_RefID,
			&_CustomFieldIdsVirtual,
			&_ReleaseDuration,
			&_Completed,
			&_CompletedDate,
		)
		if err != nil {
			return nil, err
		}
		t := &IssueSummary{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _TotalIssues.Valid {
			t.SetTotalIssues(int32(_TotalIssues.Int64))
		}
		if _New30Days.Valid {
			t.SetNew30Days(int32(_New30Days.Int64))
		}
		if _TotalClosed.Valid {
			t.SetTotalClosed(int32(_TotalClosed.Int64))
		}
		if _Closed30Days.Valid {
			t.SetClosed30Days(int32(_Closed30Days.Int64))
		}
		if _EstimatedWorkMonths.Valid {
			t.SetEstimatedWorkMonths(_EstimatedWorkMonths.Float64)
		}
		if _EstimatedWorkMonths30Days.Valid {
			t.SetEstimatedWorkMonths30Days(_EstimatedWorkMonths30Days.Float64)
		}
		if _Title.Valid {
			t.SetTitle(_Title.String)
		}
		if _URL.Valid {
			t.SetURL(_URL.String)
		}
		if _Priority.Valid {
			t.SetPriority(_Priority.String)
		}
		if _PriorityID.Valid {
			t.SetPriorityID(_PriorityID.String)
		}
		if _Status.Valid {
			t.SetStatus(_Status.String)
		}
		if _StatusID.Valid {
			t.SetStatusID(_StatusID.String)
		}
		if _IssueType.Valid {
			t.SetIssueType(_IssueType.String)
		}
		if _IssueTypeID.Valid {
			t.SetIssueTypeID(_IssueTypeID.String)
		}
		if _Resolution.Valid {
			t.SetResolution(_Resolution.String)
		}
		if _ResolutionID.Valid {
			t.SetResolutionID(_ResolutionID.String)
		}
		if _State.Valid {
			t.SetState(_State.String)
		}
		if _CustomFieldIds.Valid {
			t.SetCustomFieldIds(_CustomFieldIds.String)
		}
		if _Teams.Valid {
			t.SetTeams(_Teams.String)
		}
		if _ParentIssueID.Valid {
			t.SetParentIssueID(_ParentIssueID.String)
		}
		if _ParentsIssueIds.Valid {
			t.SetParentsIssueIds(_ParentsIssueIds.String)
		}
		if _Metadata.Valid {
			t.SetMetadata(_Metadata.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _Sprints.Valid {
			t.SetSprints(_Sprints.String)
		}
		if _Labels.Valid {
			t.SetLabels(_Labels.String)
		}
		if _TopLevel.Valid {
			t.SetTopLevel(_TopLevel.Bool)
		}
		if _IsLeaf.Valid {
			t.SetIsLeaf(_IsLeaf.Bool)
		}
		if _Path.Valid {
			t.SetPath(_Path.String)
		}
		if _InProgressDuration.Valid {
			t.SetInProgressDuration(_InProgressDuration.Int64)
		}
		if _VerificationDuration.Valid {
			t.SetVerificationDuration(_VerificationDuration.Int64)
		}
		if _InProgressCount.Valid {
			t.SetInProgressCount(int32(_InProgressCount.Int64))
		}
		if _ReopenCount.Valid {
			t.SetReopenCount(int32(_ReopenCount.Int64))
		}
		if _MappedType.Valid {
			t.SetMappedType(_MappedType.String)
		}
		if _StrategicParentID.Valid {
			t.SetStrategicParentID(_StrategicParentID.String)
		}
		if _SprintID.Valid {
			t.SetSprintID(_SprintID.String)
		}
		if _IssueProjectName.Valid {
			t.SetIssueProjectName(_IssueProjectName.String)
		}
		if _Users.Valid {
			t.SetUsers(_Users.String)
		}
		if _InitialStartDate.Valid {
			t.SetInitialStartDate(_InitialStartDate.Int64)
		}
		if _TotalDuration.Valid {
			t.SetTotalDuration(_TotalDuration.Int64)
		}
		if _CreatedAt.Valid {
			t.SetCreatedAt(_CreatedAt.Int64)
		}
		if _ClosedAt.Valid {
			t.SetClosedAt(_ClosedAt.Int64)
		}
		if _PlannedEndDate.Valid {
			t.SetPlannedEndDate(_PlannedEndDate.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefType.Valid {
			t.SetRefType(_RefType.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		if _CustomFieldIdsVirtual.Valid {
			t.SetCustomFieldIdsVirtual(_CustomFieldIdsVirtual.String)
		}
		if _ReleaseDuration.Valid {
			t.SetReleaseDuration(_ReleaseDuration.Int64)
		}
		if _Completed.Valid {
			t.SetCompleted(_Completed.Bool)
		}
		if _CompletedDate.Valid {
			t.SetCompletedDate(_CompletedDate.Int64)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindIssueSummariesByProjectIDTx will find all IssueSummarys by the ProjectID value using the provided transaction
func FindIssueSummariesByProjectIDTx(ctx context.Context, tx Tx, value string) ([]*IssueSummary, error) {
	q := "SELECT `issue_summary`.`id`,`issue_summary`.`checksum`,`issue_summary`.`issue_id`,`issue_summary`.`total_issues`,`issue_summary`.`new30_days`,`issue_summary`.`total_closed`,`issue_summary`.`closed30_days`,`issue_summary`.`estimated_work_months`,`issue_summary`.`estimated_work_months30_days`,`issue_summary`.`title`,`issue_summary`.`url`,`issue_summary`.`priority`,`issue_summary`.`priority_id`,`issue_summary`.`status`,`issue_summary`.`status_id`,`issue_summary`.`issue_type`,`issue_summary`.`issue_type_id`,`issue_summary`.`resolution`,`issue_summary`.`resolution_id`,`issue_summary`.`state`,`issue_summary`.`custom_field_ids`,`issue_summary`.`teams`,`issue_summary`.`parent_issue_id`,`issue_summary`.`parents_issue_ids`,`issue_summary`.`metadata`,`issue_summary`.`project_id`,`issue_summary`.`sprints`,`issue_summary`.`labels`,`issue_summary`.`top_level`,`issue_summary`.`is_leaf`,`issue_summary`.`path`,`issue_summary`.`in_progress_duration`,`issue_summary`.`verification_duration`,`issue_summary`.`in_progress_count`,`issue_summary`.`reopen_count`,`issue_summary`.`mapped_type`,`issue_summary`.`strategic_parent_id`,`issue_summary`.`sprint_id`,`issue_summary`.`issue_project_name`,`issue_summary`.`users`,`issue_summary`.`initial_start_date`,`issue_summary`.`total_duration`,`issue_summary`.`created_at`,`issue_summary`.`closed_at`,`issue_summary`.`planned_end_date`,`issue_summary`.`customer_id`,`issue_summary`.`ref_type`,`issue_summary`.`ref_id`,`issue_summary`.`custom_field_ids_virtual`,`issue_summary`.`release_duration`,`issue_summary`.`completed`,`issue_summary`.`completed_date` FROM `issue_summary` WHERE `project_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*IssueSummary, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _IssueID sql.NullString
		var _TotalIssues sql.NullInt64
		var _New30Days sql.NullInt64
		var _TotalClosed sql.NullInt64
		var _Closed30Days sql.NullInt64
		var _EstimatedWorkMonths sql.NullFloat64
		var _EstimatedWorkMonths30Days sql.NullFloat64
		var _Title sql.NullString
		var _URL sql.NullString
		var _Priority sql.NullString
		var _PriorityID sql.NullString
		var _Status sql.NullString
		var _StatusID sql.NullString
		var _IssueType sql.NullString
		var _IssueTypeID sql.NullString
		var _Resolution sql.NullString
		var _ResolutionID sql.NullString
		var _State sql.NullString
		var _CustomFieldIds sql.NullString
		var _Teams sql.NullString
		var _ParentIssueID sql.NullString
		var _ParentsIssueIds sql.NullString
		var _Metadata sql.NullString
		var _ProjectID sql.NullString
		var _Sprints sql.NullString
		var _Labels sql.NullString
		var _TopLevel sql.NullBool
		var _IsLeaf sql.NullBool
		var _Path sql.NullString
		var _InProgressDuration sql.NullInt64
		var _VerificationDuration sql.NullInt64
		var _InProgressCount sql.NullInt64
		var _ReopenCount sql.NullInt64
		var _MappedType sql.NullString
		var _StrategicParentID sql.NullString
		var _SprintID sql.NullString
		var _IssueProjectName sql.NullString
		var _Users sql.NullString
		var _InitialStartDate sql.NullInt64
		var _TotalDuration sql.NullInt64
		var _CreatedAt sql.NullInt64
		var _ClosedAt sql.NullInt64
		var _PlannedEndDate sql.NullInt64
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefID sql.NullString
		var _CustomFieldIdsVirtual sql.NullString
		var _ReleaseDuration sql.NullInt64
		var _Completed sql.NullBool
		var _CompletedDate sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_IssueID,
			&_TotalIssues,
			&_New30Days,
			&_TotalClosed,
			&_Closed30Days,
			&_EstimatedWorkMonths,
			&_EstimatedWorkMonths30Days,
			&_Title,
			&_URL,
			&_Priority,
			&_PriorityID,
			&_Status,
			&_StatusID,
			&_IssueType,
			&_IssueTypeID,
			&_Resolution,
			&_ResolutionID,
			&_State,
			&_CustomFieldIds,
			&_Teams,
			&_ParentIssueID,
			&_ParentsIssueIds,
			&_Metadata,
			&_ProjectID,
			&_Sprints,
			&_Labels,
			&_TopLevel,
			&_IsLeaf,
			&_Path,
			&_InProgressDuration,
			&_VerificationDuration,
			&_InProgressCount,
			&_ReopenCount,
			&_MappedType,
			&_StrategicParentID,
			&_SprintID,
			&_IssueProjectName,
			&_Users,
			&_InitialStartDate,
			&_TotalDuration,
			&_CreatedAt,
			&_ClosedAt,
			&_PlannedEndDate,
			&_CustomerID,
			&_RefType,
			&_RefID,
			&_CustomFieldIdsVirtual,
			&_ReleaseDuration,
			&_Completed,
			&_CompletedDate,
		)
		if err != nil {
			return nil, err
		}
		t := &IssueSummary{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _TotalIssues.Valid {
			t.SetTotalIssues(int32(_TotalIssues.Int64))
		}
		if _New30Days.Valid {
			t.SetNew30Days(int32(_New30Days.Int64))
		}
		if _TotalClosed.Valid {
			t.SetTotalClosed(int32(_TotalClosed.Int64))
		}
		if _Closed30Days.Valid {
			t.SetClosed30Days(int32(_Closed30Days.Int64))
		}
		if _EstimatedWorkMonths.Valid {
			t.SetEstimatedWorkMonths(_EstimatedWorkMonths.Float64)
		}
		if _EstimatedWorkMonths30Days.Valid {
			t.SetEstimatedWorkMonths30Days(_EstimatedWorkMonths30Days.Float64)
		}
		if _Title.Valid {
			t.SetTitle(_Title.String)
		}
		if _URL.Valid {
			t.SetURL(_URL.String)
		}
		if _Priority.Valid {
			t.SetPriority(_Priority.String)
		}
		if _PriorityID.Valid {
			t.SetPriorityID(_PriorityID.String)
		}
		if _Status.Valid {
			t.SetStatus(_Status.String)
		}
		if _StatusID.Valid {
			t.SetStatusID(_StatusID.String)
		}
		if _IssueType.Valid {
			t.SetIssueType(_IssueType.String)
		}
		if _IssueTypeID.Valid {
			t.SetIssueTypeID(_IssueTypeID.String)
		}
		if _Resolution.Valid {
			t.SetResolution(_Resolution.String)
		}
		if _ResolutionID.Valid {
			t.SetResolutionID(_ResolutionID.String)
		}
		if _State.Valid {
			t.SetState(_State.String)
		}
		if _CustomFieldIds.Valid {
			t.SetCustomFieldIds(_CustomFieldIds.String)
		}
		if _Teams.Valid {
			t.SetTeams(_Teams.String)
		}
		if _ParentIssueID.Valid {
			t.SetParentIssueID(_ParentIssueID.String)
		}
		if _ParentsIssueIds.Valid {
			t.SetParentsIssueIds(_ParentsIssueIds.String)
		}
		if _Metadata.Valid {
			t.SetMetadata(_Metadata.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _Sprints.Valid {
			t.SetSprints(_Sprints.String)
		}
		if _Labels.Valid {
			t.SetLabels(_Labels.String)
		}
		if _TopLevel.Valid {
			t.SetTopLevel(_TopLevel.Bool)
		}
		if _IsLeaf.Valid {
			t.SetIsLeaf(_IsLeaf.Bool)
		}
		if _Path.Valid {
			t.SetPath(_Path.String)
		}
		if _InProgressDuration.Valid {
			t.SetInProgressDuration(_InProgressDuration.Int64)
		}
		if _VerificationDuration.Valid {
			t.SetVerificationDuration(_VerificationDuration.Int64)
		}
		if _InProgressCount.Valid {
			t.SetInProgressCount(int32(_InProgressCount.Int64))
		}
		if _ReopenCount.Valid {
			t.SetReopenCount(int32(_ReopenCount.Int64))
		}
		if _MappedType.Valid {
			t.SetMappedType(_MappedType.String)
		}
		if _StrategicParentID.Valid {
			t.SetStrategicParentID(_StrategicParentID.String)
		}
		if _SprintID.Valid {
			t.SetSprintID(_SprintID.String)
		}
		if _IssueProjectName.Valid {
			t.SetIssueProjectName(_IssueProjectName.String)
		}
		if _Users.Valid {
			t.SetUsers(_Users.String)
		}
		if _InitialStartDate.Valid {
			t.SetInitialStartDate(_InitialStartDate.Int64)
		}
		if _TotalDuration.Valid {
			t.SetTotalDuration(_TotalDuration.Int64)
		}
		if _CreatedAt.Valid {
			t.SetCreatedAt(_CreatedAt.Int64)
		}
		if _ClosedAt.Valid {
			t.SetClosedAt(_ClosedAt.Int64)
		}
		if _PlannedEndDate.Valid {
			t.SetPlannedEndDate(_PlannedEndDate.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefType.Valid {
			t.SetRefType(_RefType.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		if _CustomFieldIdsVirtual.Valid {
			t.SetCustomFieldIdsVirtual(_CustomFieldIdsVirtual.String)
		}
		if _ReleaseDuration.Valid {
			t.SetReleaseDuration(_ReleaseDuration.Int64)
		}
		if _Completed.Valid {
			t.SetCompleted(_Completed.Bool)
		}
		if _CompletedDate.Valid {
			t.SetCompletedDate(_CompletedDate.Int64)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetSprints will return the IssueSummary Sprints value
func (t *IssueSummary) GetSprints() string {
	if t.Sprints == nil {
		return ""
	}
	return *t.Sprints
}

// SetSprints will set the IssueSummary Sprints value
func (t *IssueSummary) SetSprints(v string) {
	t.Sprints = &v
}

// GetLabels will return the IssueSummary Labels value
func (t *IssueSummary) GetLabels() string {
	if t.Labels == nil {
		return ""
	}
	return *t.Labels
}

// SetLabels will set the IssueSummary Labels value
func (t *IssueSummary) SetLabels(v string) {
	t.Labels = &v
}

// GetTopLevel will return the IssueSummary TopLevel value
func (t *IssueSummary) GetTopLevel() bool {
	return t.TopLevel
}

// SetTopLevel will set the IssueSummary TopLevel value
func (t *IssueSummary) SetTopLevel(v bool) {
	t.TopLevel = v
}

// FindIssueSummariesByTopLevel will find all IssueSummarys by the TopLevel value
func FindIssueSummariesByTopLevel(ctx context.Context, db DB, value bool) ([]*IssueSummary, error) {
	q := "SELECT `issue_summary`.`id`,`issue_summary`.`checksum`,`issue_summary`.`issue_id`,`issue_summary`.`total_issues`,`issue_summary`.`new30_days`,`issue_summary`.`total_closed`,`issue_summary`.`closed30_days`,`issue_summary`.`estimated_work_months`,`issue_summary`.`estimated_work_months30_days`,`issue_summary`.`title`,`issue_summary`.`url`,`issue_summary`.`priority`,`issue_summary`.`priority_id`,`issue_summary`.`status`,`issue_summary`.`status_id`,`issue_summary`.`issue_type`,`issue_summary`.`issue_type_id`,`issue_summary`.`resolution`,`issue_summary`.`resolution_id`,`issue_summary`.`state`,`issue_summary`.`custom_field_ids`,`issue_summary`.`teams`,`issue_summary`.`parent_issue_id`,`issue_summary`.`parents_issue_ids`,`issue_summary`.`metadata`,`issue_summary`.`project_id`,`issue_summary`.`sprints`,`issue_summary`.`labels`,`issue_summary`.`top_level`,`issue_summary`.`is_leaf`,`issue_summary`.`path`,`issue_summary`.`in_progress_duration`,`issue_summary`.`verification_duration`,`issue_summary`.`in_progress_count`,`issue_summary`.`reopen_count`,`issue_summary`.`mapped_type`,`issue_summary`.`strategic_parent_id`,`issue_summary`.`sprint_id`,`issue_summary`.`issue_project_name`,`issue_summary`.`users`,`issue_summary`.`initial_start_date`,`issue_summary`.`total_duration`,`issue_summary`.`created_at`,`issue_summary`.`closed_at`,`issue_summary`.`planned_end_date`,`issue_summary`.`customer_id`,`issue_summary`.`ref_type`,`issue_summary`.`ref_id`,`issue_summary`.`custom_field_ids_virtual`,`issue_summary`.`release_duration`,`issue_summary`.`completed`,`issue_summary`.`completed_date` FROM `issue_summary` WHERE `top_level` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*IssueSummary, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _IssueID sql.NullString
		var _TotalIssues sql.NullInt64
		var _New30Days sql.NullInt64
		var _TotalClosed sql.NullInt64
		var _Closed30Days sql.NullInt64
		var _EstimatedWorkMonths sql.NullFloat64
		var _EstimatedWorkMonths30Days sql.NullFloat64
		var _Title sql.NullString
		var _URL sql.NullString
		var _Priority sql.NullString
		var _PriorityID sql.NullString
		var _Status sql.NullString
		var _StatusID sql.NullString
		var _IssueType sql.NullString
		var _IssueTypeID sql.NullString
		var _Resolution sql.NullString
		var _ResolutionID sql.NullString
		var _State sql.NullString
		var _CustomFieldIds sql.NullString
		var _Teams sql.NullString
		var _ParentIssueID sql.NullString
		var _ParentsIssueIds sql.NullString
		var _Metadata sql.NullString
		var _ProjectID sql.NullString
		var _Sprints sql.NullString
		var _Labels sql.NullString
		var _TopLevel sql.NullBool
		var _IsLeaf sql.NullBool
		var _Path sql.NullString
		var _InProgressDuration sql.NullInt64
		var _VerificationDuration sql.NullInt64
		var _InProgressCount sql.NullInt64
		var _ReopenCount sql.NullInt64
		var _MappedType sql.NullString
		var _StrategicParentID sql.NullString
		var _SprintID sql.NullString
		var _IssueProjectName sql.NullString
		var _Users sql.NullString
		var _InitialStartDate sql.NullInt64
		var _TotalDuration sql.NullInt64
		var _CreatedAt sql.NullInt64
		var _ClosedAt sql.NullInt64
		var _PlannedEndDate sql.NullInt64
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefID sql.NullString
		var _CustomFieldIdsVirtual sql.NullString
		var _ReleaseDuration sql.NullInt64
		var _Completed sql.NullBool
		var _CompletedDate sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_IssueID,
			&_TotalIssues,
			&_New30Days,
			&_TotalClosed,
			&_Closed30Days,
			&_EstimatedWorkMonths,
			&_EstimatedWorkMonths30Days,
			&_Title,
			&_URL,
			&_Priority,
			&_PriorityID,
			&_Status,
			&_StatusID,
			&_IssueType,
			&_IssueTypeID,
			&_Resolution,
			&_ResolutionID,
			&_State,
			&_CustomFieldIds,
			&_Teams,
			&_ParentIssueID,
			&_ParentsIssueIds,
			&_Metadata,
			&_ProjectID,
			&_Sprints,
			&_Labels,
			&_TopLevel,
			&_IsLeaf,
			&_Path,
			&_InProgressDuration,
			&_VerificationDuration,
			&_InProgressCount,
			&_ReopenCount,
			&_MappedType,
			&_StrategicParentID,
			&_SprintID,
			&_IssueProjectName,
			&_Users,
			&_InitialStartDate,
			&_TotalDuration,
			&_CreatedAt,
			&_ClosedAt,
			&_PlannedEndDate,
			&_CustomerID,
			&_RefType,
			&_RefID,
			&_CustomFieldIdsVirtual,
			&_ReleaseDuration,
			&_Completed,
			&_CompletedDate,
		)
		if err != nil {
			return nil, err
		}
		t := &IssueSummary{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _TotalIssues.Valid {
			t.SetTotalIssues(int32(_TotalIssues.Int64))
		}
		if _New30Days.Valid {
			t.SetNew30Days(int32(_New30Days.Int64))
		}
		if _TotalClosed.Valid {
			t.SetTotalClosed(int32(_TotalClosed.Int64))
		}
		if _Closed30Days.Valid {
			t.SetClosed30Days(int32(_Closed30Days.Int64))
		}
		if _EstimatedWorkMonths.Valid {
			t.SetEstimatedWorkMonths(_EstimatedWorkMonths.Float64)
		}
		if _EstimatedWorkMonths30Days.Valid {
			t.SetEstimatedWorkMonths30Days(_EstimatedWorkMonths30Days.Float64)
		}
		if _Title.Valid {
			t.SetTitle(_Title.String)
		}
		if _URL.Valid {
			t.SetURL(_URL.String)
		}
		if _Priority.Valid {
			t.SetPriority(_Priority.String)
		}
		if _PriorityID.Valid {
			t.SetPriorityID(_PriorityID.String)
		}
		if _Status.Valid {
			t.SetStatus(_Status.String)
		}
		if _StatusID.Valid {
			t.SetStatusID(_StatusID.String)
		}
		if _IssueType.Valid {
			t.SetIssueType(_IssueType.String)
		}
		if _IssueTypeID.Valid {
			t.SetIssueTypeID(_IssueTypeID.String)
		}
		if _Resolution.Valid {
			t.SetResolution(_Resolution.String)
		}
		if _ResolutionID.Valid {
			t.SetResolutionID(_ResolutionID.String)
		}
		if _State.Valid {
			t.SetState(_State.String)
		}
		if _CustomFieldIds.Valid {
			t.SetCustomFieldIds(_CustomFieldIds.String)
		}
		if _Teams.Valid {
			t.SetTeams(_Teams.String)
		}
		if _ParentIssueID.Valid {
			t.SetParentIssueID(_ParentIssueID.String)
		}
		if _ParentsIssueIds.Valid {
			t.SetParentsIssueIds(_ParentsIssueIds.String)
		}
		if _Metadata.Valid {
			t.SetMetadata(_Metadata.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _Sprints.Valid {
			t.SetSprints(_Sprints.String)
		}
		if _Labels.Valid {
			t.SetLabels(_Labels.String)
		}
		if _TopLevel.Valid {
			t.SetTopLevel(_TopLevel.Bool)
		}
		if _IsLeaf.Valid {
			t.SetIsLeaf(_IsLeaf.Bool)
		}
		if _Path.Valid {
			t.SetPath(_Path.String)
		}
		if _InProgressDuration.Valid {
			t.SetInProgressDuration(_InProgressDuration.Int64)
		}
		if _VerificationDuration.Valid {
			t.SetVerificationDuration(_VerificationDuration.Int64)
		}
		if _InProgressCount.Valid {
			t.SetInProgressCount(int32(_InProgressCount.Int64))
		}
		if _ReopenCount.Valid {
			t.SetReopenCount(int32(_ReopenCount.Int64))
		}
		if _MappedType.Valid {
			t.SetMappedType(_MappedType.String)
		}
		if _StrategicParentID.Valid {
			t.SetStrategicParentID(_StrategicParentID.String)
		}
		if _SprintID.Valid {
			t.SetSprintID(_SprintID.String)
		}
		if _IssueProjectName.Valid {
			t.SetIssueProjectName(_IssueProjectName.String)
		}
		if _Users.Valid {
			t.SetUsers(_Users.String)
		}
		if _InitialStartDate.Valid {
			t.SetInitialStartDate(_InitialStartDate.Int64)
		}
		if _TotalDuration.Valid {
			t.SetTotalDuration(_TotalDuration.Int64)
		}
		if _CreatedAt.Valid {
			t.SetCreatedAt(_CreatedAt.Int64)
		}
		if _ClosedAt.Valid {
			t.SetClosedAt(_ClosedAt.Int64)
		}
		if _PlannedEndDate.Valid {
			t.SetPlannedEndDate(_PlannedEndDate.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefType.Valid {
			t.SetRefType(_RefType.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		if _CustomFieldIdsVirtual.Valid {
			t.SetCustomFieldIdsVirtual(_CustomFieldIdsVirtual.String)
		}
		if _ReleaseDuration.Valid {
			t.SetReleaseDuration(_ReleaseDuration.Int64)
		}
		if _Completed.Valid {
			t.SetCompleted(_Completed.Bool)
		}
		if _CompletedDate.Valid {
			t.SetCompletedDate(_CompletedDate.Int64)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindIssueSummariesByTopLevelTx will find all IssueSummarys by the TopLevel value using the provided transaction
func FindIssueSummariesByTopLevelTx(ctx context.Context, tx Tx, value bool) ([]*IssueSummary, error) {
	q := "SELECT `issue_summary`.`id`,`issue_summary`.`checksum`,`issue_summary`.`issue_id`,`issue_summary`.`total_issues`,`issue_summary`.`new30_days`,`issue_summary`.`total_closed`,`issue_summary`.`closed30_days`,`issue_summary`.`estimated_work_months`,`issue_summary`.`estimated_work_months30_days`,`issue_summary`.`title`,`issue_summary`.`url`,`issue_summary`.`priority`,`issue_summary`.`priority_id`,`issue_summary`.`status`,`issue_summary`.`status_id`,`issue_summary`.`issue_type`,`issue_summary`.`issue_type_id`,`issue_summary`.`resolution`,`issue_summary`.`resolution_id`,`issue_summary`.`state`,`issue_summary`.`custom_field_ids`,`issue_summary`.`teams`,`issue_summary`.`parent_issue_id`,`issue_summary`.`parents_issue_ids`,`issue_summary`.`metadata`,`issue_summary`.`project_id`,`issue_summary`.`sprints`,`issue_summary`.`labels`,`issue_summary`.`top_level`,`issue_summary`.`is_leaf`,`issue_summary`.`path`,`issue_summary`.`in_progress_duration`,`issue_summary`.`verification_duration`,`issue_summary`.`in_progress_count`,`issue_summary`.`reopen_count`,`issue_summary`.`mapped_type`,`issue_summary`.`strategic_parent_id`,`issue_summary`.`sprint_id`,`issue_summary`.`issue_project_name`,`issue_summary`.`users`,`issue_summary`.`initial_start_date`,`issue_summary`.`total_duration`,`issue_summary`.`created_at`,`issue_summary`.`closed_at`,`issue_summary`.`planned_end_date`,`issue_summary`.`customer_id`,`issue_summary`.`ref_type`,`issue_summary`.`ref_id`,`issue_summary`.`custom_field_ids_virtual`,`issue_summary`.`release_duration`,`issue_summary`.`completed`,`issue_summary`.`completed_date` FROM `issue_summary` WHERE `top_level` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*IssueSummary, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _IssueID sql.NullString
		var _TotalIssues sql.NullInt64
		var _New30Days sql.NullInt64
		var _TotalClosed sql.NullInt64
		var _Closed30Days sql.NullInt64
		var _EstimatedWorkMonths sql.NullFloat64
		var _EstimatedWorkMonths30Days sql.NullFloat64
		var _Title sql.NullString
		var _URL sql.NullString
		var _Priority sql.NullString
		var _PriorityID sql.NullString
		var _Status sql.NullString
		var _StatusID sql.NullString
		var _IssueType sql.NullString
		var _IssueTypeID sql.NullString
		var _Resolution sql.NullString
		var _ResolutionID sql.NullString
		var _State sql.NullString
		var _CustomFieldIds sql.NullString
		var _Teams sql.NullString
		var _ParentIssueID sql.NullString
		var _ParentsIssueIds sql.NullString
		var _Metadata sql.NullString
		var _ProjectID sql.NullString
		var _Sprints sql.NullString
		var _Labels sql.NullString
		var _TopLevel sql.NullBool
		var _IsLeaf sql.NullBool
		var _Path sql.NullString
		var _InProgressDuration sql.NullInt64
		var _VerificationDuration sql.NullInt64
		var _InProgressCount sql.NullInt64
		var _ReopenCount sql.NullInt64
		var _MappedType sql.NullString
		var _StrategicParentID sql.NullString
		var _SprintID sql.NullString
		var _IssueProjectName sql.NullString
		var _Users sql.NullString
		var _InitialStartDate sql.NullInt64
		var _TotalDuration sql.NullInt64
		var _CreatedAt sql.NullInt64
		var _ClosedAt sql.NullInt64
		var _PlannedEndDate sql.NullInt64
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefID sql.NullString
		var _CustomFieldIdsVirtual sql.NullString
		var _ReleaseDuration sql.NullInt64
		var _Completed sql.NullBool
		var _CompletedDate sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_IssueID,
			&_TotalIssues,
			&_New30Days,
			&_TotalClosed,
			&_Closed30Days,
			&_EstimatedWorkMonths,
			&_EstimatedWorkMonths30Days,
			&_Title,
			&_URL,
			&_Priority,
			&_PriorityID,
			&_Status,
			&_StatusID,
			&_IssueType,
			&_IssueTypeID,
			&_Resolution,
			&_ResolutionID,
			&_State,
			&_CustomFieldIds,
			&_Teams,
			&_ParentIssueID,
			&_ParentsIssueIds,
			&_Metadata,
			&_ProjectID,
			&_Sprints,
			&_Labels,
			&_TopLevel,
			&_IsLeaf,
			&_Path,
			&_InProgressDuration,
			&_VerificationDuration,
			&_InProgressCount,
			&_ReopenCount,
			&_MappedType,
			&_StrategicParentID,
			&_SprintID,
			&_IssueProjectName,
			&_Users,
			&_InitialStartDate,
			&_TotalDuration,
			&_CreatedAt,
			&_ClosedAt,
			&_PlannedEndDate,
			&_CustomerID,
			&_RefType,
			&_RefID,
			&_CustomFieldIdsVirtual,
			&_ReleaseDuration,
			&_Completed,
			&_CompletedDate,
		)
		if err != nil {
			return nil, err
		}
		t := &IssueSummary{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _TotalIssues.Valid {
			t.SetTotalIssues(int32(_TotalIssues.Int64))
		}
		if _New30Days.Valid {
			t.SetNew30Days(int32(_New30Days.Int64))
		}
		if _TotalClosed.Valid {
			t.SetTotalClosed(int32(_TotalClosed.Int64))
		}
		if _Closed30Days.Valid {
			t.SetClosed30Days(int32(_Closed30Days.Int64))
		}
		if _EstimatedWorkMonths.Valid {
			t.SetEstimatedWorkMonths(_EstimatedWorkMonths.Float64)
		}
		if _EstimatedWorkMonths30Days.Valid {
			t.SetEstimatedWorkMonths30Days(_EstimatedWorkMonths30Days.Float64)
		}
		if _Title.Valid {
			t.SetTitle(_Title.String)
		}
		if _URL.Valid {
			t.SetURL(_URL.String)
		}
		if _Priority.Valid {
			t.SetPriority(_Priority.String)
		}
		if _PriorityID.Valid {
			t.SetPriorityID(_PriorityID.String)
		}
		if _Status.Valid {
			t.SetStatus(_Status.String)
		}
		if _StatusID.Valid {
			t.SetStatusID(_StatusID.String)
		}
		if _IssueType.Valid {
			t.SetIssueType(_IssueType.String)
		}
		if _IssueTypeID.Valid {
			t.SetIssueTypeID(_IssueTypeID.String)
		}
		if _Resolution.Valid {
			t.SetResolution(_Resolution.String)
		}
		if _ResolutionID.Valid {
			t.SetResolutionID(_ResolutionID.String)
		}
		if _State.Valid {
			t.SetState(_State.String)
		}
		if _CustomFieldIds.Valid {
			t.SetCustomFieldIds(_CustomFieldIds.String)
		}
		if _Teams.Valid {
			t.SetTeams(_Teams.String)
		}
		if _ParentIssueID.Valid {
			t.SetParentIssueID(_ParentIssueID.String)
		}
		if _ParentsIssueIds.Valid {
			t.SetParentsIssueIds(_ParentsIssueIds.String)
		}
		if _Metadata.Valid {
			t.SetMetadata(_Metadata.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _Sprints.Valid {
			t.SetSprints(_Sprints.String)
		}
		if _Labels.Valid {
			t.SetLabels(_Labels.String)
		}
		if _TopLevel.Valid {
			t.SetTopLevel(_TopLevel.Bool)
		}
		if _IsLeaf.Valid {
			t.SetIsLeaf(_IsLeaf.Bool)
		}
		if _Path.Valid {
			t.SetPath(_Path.String)
		}
		if _InProgressDuration.Valid {
			t.SetInProgressDuration(_InProgressDuration.Int64)
		}
		if _VerificationDuration.Valid {
			t.SetVerificationDuration(_VerificationDuration.Int64)
		}
		if _InProgressCount.Valid {
			t.SetInProgressCount(int32(_InProgressCount.Int64))
		}
		if _ReopenCount.Valid {
			t.SetReopenCount(int32(_ReopenCount.Int64))
		}
		if _MappedType.Valid {
			t.SetMappedType(_MappedType.String)
		}
		if _StrategicParentID.Valid {
			t.SetStrategicParentID(_StrategicParentID.String)
		}
		if _SprintID.Valid {
			t.SetSprintID(_SprintID.String)
		}
		if _IssueProjectName.Valid {
			t.SetIssueProjectName(_IssueProjectName.String)
		}
		if _Users.Valid {
			t.SetUsers(_Users.String)
		}
		if _InitialStartDate.Valid {
			t.SetInitialStartDate(_InitialStartDate.Int64)
		}
		if _TotalDuration.Valid {
			t.SetTotalDuration(_TotalDuration.Int64)
		}
		if _CreatedAt.Valid {
			t.SetCreatedAt(_CreatedAt.Int64)
		}
		if _ClosedAt.Valid {
			t.SetClosedAt(_ClosedAt.Int64)
		}
		if _PlannedEndDate.Valid {
			t.SetPlannedEndDate(_PlannedEndDate.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefType.Valid {
			t.SetRefType(_RefType.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		if _CustomFieldIdsVirtual.Valid {
			t.SetCustomFieldIdsVirtual(_CustomFieldIdsVirtual.String)
		}
		if _ReleaseDuration.Valid {
			t.SetReleaseDuration(_ReleaseDuration.Int64)
		}
		if _Completed.Valid {
			t.SetCompleted(_Completed.Bool)
		}
		if _CompletedDate.Valid {
			t.SetCompletedDate(_CompletedDate.Int64)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetIsLeaf will return the IssueSummary IsLeaf value
func (t *IssueSummary) GetIsLeaf() bool {
	return t.IsLeaf
}

// SetIsLeaf will set the IssueSummary IsLeaf value
func (t *IssueSummary) SetIsLeaf(v bool) {
	t.IsLeaf = v
}

// GetPath will return the IssueSummary Path value
func (t *IssueSummary) GetPath() string {
	return t.Path
}

// SetPath will set the IssueSummary Path value
func (t *IssueSummary) SetPath(v string) {
	t.Path = v
}

// GetInProgressDuration will return the IssueSummary InProgressDuration value
func (t *IssueSummary) GetInProgressDuration() int64 {
	return t.InProgressDuration
}

// SetInProgressDuration will set the IssueSummary InProgressDuration value
func (t *IssueSummary) SetInProgressDuration(v int64) {
	t.InProgressDuration = v
}

// GetVerificationDuration will return the IssueSummary VerificationDuration value
func (t *IssueSummary) GetVerificationDuration() int64 {
	return t.VerificationDuration
}

// SetVerificationDuration will set the IssueSummary VerificationDuration value
func (t *IssueSummary) SetVerificationDuration(v int64) {
	t.VerificationDuration = v
}

// GetInProgressCount will return the IssueSummary InProgressCount value
func (t *IssueSummary) GetInProgressCount() int32 {
	return t.InProgressCount
}

// SetInProgressCount will set the IssueSummary InProgressCount value
func (t *IssueSummary) SetInProgressCount(v int32) {
	t.InProgressCount = v
}

// GetReopenCount will return the IssueSummary ReopenCount value
func (t *IssueSummary) GetReopenCount() int32 {
	return t.ReopenCount
}

// SetReopenCount will set the IssueSummary ReopenCount value
func (t *IssueSummary) SetReopenCount(v int32) {
	t.ReopenCount = v
}

// GetMappedType will return the IssueSummary MappedType value
func (t *IssueSummary) GetMappedType() string {
	return t.MappedType
}

// SetMappedType will set the IssueSummary MappedType value
func (t *IssueSummary) SetMappedType(v string) {
	t.MappedType = v
}

// GetStrategicParentID will return the IssueSummary StrategicParentID value
func (t *IssueSummary) GetStrategicParentID() string {
	if t.StrategicParentID == nil {
		return ""
	}
	return *t.StrategicParentID
}

// SetStrategicParentID will set the IssueSummary StrategicParentID value
func (t *IssueSummary) SetStrategicParentID(v string) {
	t.StrategicParentID = &v
}

// FindIssueSummariesByStrategicParentID will find all IssueSummarys by the StrategicParentID value
func FindIssueSummariesByStrategicParentID(ctx context.Context, db DB, value string) ([]*IssueSummary, error) {
	q := "SELECT `issue_summary`.`id`,`issue_summary`.`checksum`,`issue_summary`.`issue_id`,`issue_summary`.`total_issues`,`issue_summary`.`new30_days`,`issue_summary`.`total_closed`,`issue_summary`.`closed30_days`,`issue_summary`.`estimated_work_months`,`issue_summary`.`estimated_work_months30_days`,`issue_summary`.`title`,`issue_summary`.`url`,`issue_summary`.`priority`,`issue_summary`.`priority_id`,`issue_summary`.`status`,`issue_summary`.`status_id`,`issue_summary`.`issue_type`,`issue_summary`.`issue_type_id`,`issue_summary`.`resolution`,`issue_summary`.`resolution_id`,`issue_summary`.`state`,`issue_summary`.`custom_field_ids`,`issue_summary`.`teams`,`issue_summary`.`parent_issue_id`,`issue_summary`.`parents_issue_ids`,`issue_summary`.`metadata`,`issue_summary`.`project_id`,`issue_summary`.`sprints`,`issue_summary`.`labels`,`issue_summary`.`top_level`,`issue_summary`.`is_leaf`,`issue_summary`.`path`,`issue_summary`.`in_progress_duration`,`issue_summary`.`verification_duration`,`issue_summary`.`in_progress_count`,`issue_summary`.`reopen_count`,`issue_summary`.`mapped_type`,`issue_summary`.`strategic_parent_id`,`issue_summary`.`sprint_id`,`issue_summary`.`issue_project_name`,`issue_summary`.`users`,`issue_summary`.`initial_start_date`,`issue_summary`.`total_duration`,`issue_summary`.`created_at`,`issue_summary`.`closed_at`,`issue_summary`.`planned_end_date`,`issue_summary`.`customer_id`,`issue_summary`.`ref_type`,`issue_summary`.`ref_id`,`issue_summary`.`custom_field_ids_virtual`,`issue_summary`.`release_duration`,`issue_summary`.`completed`,`issue_summary`.`completed_date` FROM `issue_summary` WHERE `strategic_parent_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*IssueSummary, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _IssueID sql.NullString
		var _TotalIssues sql.NullInt64
		var _New30Days sql.NullInt64
		var _TotalClosed sql.NullInt64
		var _Closed30Days sql.NullInt64
		var _EstimatedWorkMonths sql.NullFloat64
		var _EstimatedWorkMonths30Days sql.NullFloat64
		var _Title sql.NullString
		var _URL sql.NullString
		var _Priority sql.NullString
		var _PriorityID sql.NullString
		var _Status sql.NullString
		var _StatusID sql.NullString
		var _IssueType sql.NullString
		var _IssueTypeID sql.NullString
		var _Resolution sql.NullString
		var _ResolutionID sql.NullString
		var _State sql.NullString
		var _CustomFieldIds sql.NullString
		var _Teams sql.NullString
		var _ParentIssueID sql.NullString
		var _ParentsIssueIds sql.NullString
		var _Metadata sql.NullString
		var _ProjectID sql.NullString
		var _Sprints sql.NullString
		var _Labels sql.NullString
		var _TopLevel sql.NullBool
		var _IsLeaf sql.NullBool
		var _Path sql.NullString
		var _InProgressDuration sql.NullInt64
		var _VerificationDuration sql.NullInt64
		var _InProgressCount sql.NullInt64
		var _ReopenCount sql.NullInt64
		var _MappedType sql.NullString
		var _StrategicParentID sql.NullString
		var _SprintID sql.NullString
		var _IssueProjectName sql.NullString
		var _Users sql.NullString
		var _InitialStartDate sql.NullInt64
		var _TotalDuration sql.NullInt64
		var _CreatedAt sql.NullInt64
		var _ClosedAt sql.NullInt64
		var _PlannedEndDate sql.NullInt64
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefID sql.NullString
		var _CustomFieldIdsVirtual sql.NullString
		var _ReleaseDuration sql.NullInt64
		var _Completed sql.NullBool
		var _CompletedDate sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_IssueID,
			&_TotalIssues,
			&_New30Days,
			&_TotalClosed,
			&_Closed30Days,
			&_EstimatedWorkMonths,
			&_EstimatedWorkMonths30Days,
			&_Title,
			&_URL,
			&_Priority,
			&_PriorityID,
			&_Status,
			&_StatusID,
			&_IssueType,
			&_IssueTypeID,
			&_Resolution,
			&_ResolutionID,
			&_State,
			&_CustomFieldIds,
			&_Teams,
			&_ParentIssueID,
			&_ParentsIssueIds,
			&_Metadata,
			&_ProjectID,
			&_Sprints,
			&_Labels,
			&_TopLevel,
			&_IsLeaf,
			&_Path,
			&_InProgressDuration,
			&_VerificationDuration,
			&_InProgressCount,
			&_ReopenCount,
			&_MappedType,
			&_StrategicParentID,
			&_SprintID,
			&_IssueProjectName,
			&_Users,
			&_InitialStartDate,
			&_TotalDuration,
			&_CreatedAt,
			&_ClosedAt,
			&_PlannedEndDate,
			&_CustomerID,
			&_RefType,
			&_RefID,
			&_CustomFieldIdsVirtual,
			&_ReleaseDuration,
			&_Completed,
			&_CompletedDate,
		)
		if err != nil {
			return nil, err
		}
		t := &IssueSummary{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _TotalIssues.Valid {
			t.SetTotalIssues(int32(_TotalIssues.Int64))
		}
		if _New30Days.Valid {
			t.SetNew30Days(int32(_New30Days.Int64))
		}
		if _TotalClosed.Valid {
			t.SetTotalClosed(int32(_TotalClosed.Int64))
		}
		if _Closed30Days.Valid {
			t.SetClosed30Days(int32(_Closed30Days.Int64))
		}
		if _EstimatedWorkMonths.Valid {
			t.SetEstimatedWorkMonths(_EstimatedWorkMonths.Float64)
		}
		if _EstimatedWorkMonths30Days.Valid {
			t.SetEstimatedWorkMonths30Days(_EstimatedWorkMonths30Days.Float64)
		}
		if _Title.Valid {
			t.SetTitle(_Title.String)
		}
		if _URL.Valid {
			t.SetURL(_URL.String)
		}
		if _Priority.Valid {
			t.SetPriority(_Priority.String)
		}
		if _PriorityID.Valid {
			t.SetPriorityID(_PriorityID.String)
		}
		if _Status.Valid {
			t.SetStatus(_Status.String)
		}
		if _StatusID.Valid {
			t.SetStatusID(_StatusID.String)
		}
		if _IssueType.Valid {
			t.SetIssueType(_IssueType.String)
		}
		if _IssueTypeID.Valid {
			t.SetIssueTypeID(_IssueTypeID.String)
		}
		if _Resolution.Valid {
			t.SetResolution(_Resolution.String)
		}
		if _ResolutionID.Valid {
			t.SetResolutionID(_ResolutionID.String)
		}
		if _State.Valid {
			t.SetState(_State.String)
		}
		if _CustomFieldIds.Valid {
			t.SetCustomFieldIds(_CustomFieldIds.String)
		}
		if _Teams.Valid {
			t.SetTeams(_Teams.String)
		}
		if _ParentIssueID.Valid {
			t.SetParentIssueID(_ParentIssueID.String)
		}
		if _ParentsIssueIds.Valid {
			t.SetParentsIssueIds(_ParentsIssueIds.String)
		}
		if _Metadata.Valid {
			t.SetMetadata(_Metadata.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _Sprints.Valid {
			t.SetSprints(_Sprints.String)
		}
		if _Labels.Valid {
			t.SetLabels(_Labels.String)
		}
		if _TopLevel.Valid {
			t.SetTopLevel(_TopLevel.Bool)
		}
		if _IsLeaf.Valid {
			t.SetIsLeaf(_IsLeaf.Bool)
		}
		if _Path.Valid {
			t.SetPath(_Path.String)
		}
		if _InProgressDuration.Valid {
			t.SetInProgressDuration(_InProgressDuration.Int64)
		}
		if _VerificationDuration.Valid {
			t.SetVerificationDuration(_VerificationDuration.Int64)
		}
		if _InProgressCount.Valid {
			t.SetInProgressCount(int32(_InProgressCount.Int64))
		}
		if _ReopenCount.Valid {
			t.SetReopenCount(int32(_ReopenCount.Int64))
		}
		if _MappedType.Valid {
			t.SetMappedType(_MappedType.String)
		}
		if _StrategicParentID.Valid {
			t.SetStrategicParentID(_StrategicParentID.String)
		}
		if _SprintID.Valid {
			t.SetSprintID(_SprintID.String)
		}
		if _IssueProjectName.Valid {
			t.SetIssueProjectName(_IssueProjectName.String)
		}
		if _Users.Valid {
			t.SetUsers(_Users.String)
		}
		if _InitialStartDate.Valid {
			t.SetInitialStartDate(_InitialStartDate.Int64)
		}
		if _TotalDuration.Valid {
			t.SetTotalDuration(_TotalDuration.Int64)
		}
		if _CreatedAt.Valid {
			t.SetCreatedAt(_CreatedAt.Int64)
		}
		if _ClosedAt.Valid {
			t.SetClosedAt(_ClosedAt.Int64)
		}
		if _PlannedEndDate.Valid {
			t.SetPlannedEndDate(_PlannedEndDate.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefType.Valid {
			t.SetRefType(_RefType.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		if _CustomFieldIdsVirtual.Valid {
			t.SetCustomFieldIdsVirtual(_CustomFieldIdsVirtual.String)
		}
		if _ReleaseDuration.Valid {
			t.SetReleaseDuration(_ReleaseDuration.Int64)
		}
		if _Completed.Valid {
			t.SetCompleted(_Completed.Bool)
		}
		if _CompletedDate.Valid {
			t.SetCompletedDate(_CompletedDate.Int64)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindIssueSummariesByStrategicParentIDTx will find all IssueSummarys by the StrategicParentID value using the provided transaction
func FindIssueSummariesByStrategicParentIDTx(ctx context.Context, tx Tx, value string) ([]*IssueSummary, error) {
	q := "SELECT `issue_summary`.`id`,`issue_summary`.`checksum`,`issue_summary`.`issue_id`,`issue_summary`.`total_issues`,`issue_summary`.`new30_days`,`issue_summary`.`total_closed`,`issue_summary`.`closed30_days`,`issue_summary`.`estimated_work_months`,`issue_summary`.`estimated_work_months30_days`,`issue_summary`.`title`,`issue_summary`.`url`,`issue_summary`.`priority`,`issue_summary`.`priority_id`,`issue_summary`.`status`,`issue_summary`.`status_id`,`issue_summary`.`issue_type`,`issue_summary`.`issue_type_id`,`issue_summary`.`resolution`,`issue_summary`.`resolution_id`,`issue_summary`.`state`,`issue_summary`.`custom_field_ids`,`issue_summary`.`teams`,`issue_summary`.`parent_issue_id`,`issue_summary`.`parents_issue_ids`,`issue_summary`.`metadata`,`issue_summary`.`project_id`,`issue_summary`.`sprints`,`issue_summary`.`labels`,`issue_summary`.`top_level`,`issue_summary`.`is_leaf`,`issue_summary`.`path`,`issue_summary`.`in_progress_duration`,`issue_summary`.`verification_duration`,`issue_summary`.`in_progress_count`,`issue_summary`.`reopen_count`,`issue_summary`.`mapped_type`,`issue_summary`.`strategic_parent_id`,`issue_summary`.`sprint_id`,`issue_summary`.`issue_project_name`,`issue_summary`.`users`,`issue_summary`.`initial_start_date`,`issue_summary`.`total_duration`,`issue_summary`.`created_at`,`issue_summary`.`closed_at`,`issue_summary`.`planned_end_date`,`issue_summary`.`customer_id`,`issue_summary`.`ref_type`,`issue_summary`.`ref_id`,`issue_summary`.`custom_field_ids_virtual`,`issue_summary`.`release_duration`,`issue_summary`.`completed`,`issue_summary`.`completed_date` FROM `issue_summary` WHERE `strategic_parent_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*IssueSummary, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _IssueID sql.NullString
		var _TotalIssues sql.NullInt64
		var _New30Days sql.NullInt64
		var _TotalClosed sql.NullInt64
		var _Closed30Days sql.NullInt64
		var _EstimatedWorkMonths sql.NullFloat64
		var _EstimatedWorkMonths30Days sql.NullFloat64
		var _Title sql.NullString
		var _URL sql.NullString
		var _Priority sql.NullString
		var _PriorityID sql.NullString
		var _Status sql.NullString
		var _StatusID sql.NullString
		var _IssueType sql.NullString
		var _IssueTypeID sql.NullString
		var _Resolution sql.NullString
		var _ResolutionID sql.NullString
		var _State sql.NullString
		var _CustomFieldIds sql.NullString
		var _Teams sql.NullString
		var _ParentIssueID sql.NullString
		var _ParentsIssueIds sql.NullString
		var _Metadata sql.NullString
		var _ProjectID sql.NullString
		var _Sprints sql.NullString
		var _Labels sql.NullString
		var _TopLevel sql.NullBool
		var _IsLeaf sql.NullBool
		var _Path sql.NullString
		var _InProgressDuration sql.NullInt64
		var _VerificationDuration sql.NullInt64
		var _InProgressCount sql.NullInt64
		var _ReopenCount sql.NullInt64
		var _MappedType sql.NullString
		var _StrategicParentID sql.NullString
		var _SprintID sql.NullString
		var _IssueProjectName sql.NullString
		var _Users sql.NullString
		var _InitialStartDate sql.NullInt64
		var _TotalDuration sql.NullInt64
		var _CreatedAt sql.NullInt64
		var _ClosedAt sql.NullInt64
		var _PlannedEndDate sql.NullInt64
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefID sql.NullString
		var _CustomFieldIdsVirtual sql.NullString
		var _ReleaseDuration sql.NullInt64
		var _Completed sql.NullBool
		var _CompletedDate sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_IssueID,
			&_TotalIssues,
			&_New30Days,
			&_TotalClosed,
			&_Closed30Days,
			&_EstimatedWorkMonths,
			&_EstimatedWorkMonths30Days,
			&_Title,
			&_URL,
			&_Priority,
			&_PriorityID,
			&_Status,
			&_StatusID,
			&_IssueType,
			&_IssueTypeID,
			&_Resolution,
			&_ResolutionID,
			&_State,
			&_CustomFieldIds,
			&_Teams,
			&_ParentIssueID,
			&_ParentsIssueIds,
			&_Metadata,
			&_ProjectID,
			&_Sprints,
			&_Labels,
			&_TopLevel,
			&_IsLeaf,
			&_Path,
			&_InProgressDuration,
			&_VerificationDuration,
			&_InProgressCount,
			&_ReopenCount,
			&_MappedType,
			&_StrategicParentID,
			&_SprintID,
			&_IssueProjectName,
			&_Users,
			&_InitialStartDate,
			&_TotalDuration,
			&_CreatedAt,
			&_ClosedAt,
			&_PlannedEndDate,
			&_CustomerID,
			&_RefType,
			&_RefID,
			&_CustomFieldIdsVirtual,
			&_ReleaseDuration,
			&_Completed,
			&_CompletedDate,
		)
		if err != nil {
			return nil, err
		}
		t := &IssueSummary{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _TotalIssues.Valid {
			t.SetTotalIssues(int32(_TotalIssues.Int64))
		}
		if _New30Days.Valid {
			t.SetNew30Days(int32(_New30Days.Int64))
		}
		if _TotalClosed.Valid {
			t.SetTotalClosed(int32(_TotalClosed.Int64))
		}
		if _Closed30Days.Valid {
			t.SetClosed30Days(int32(_Closed30Days.Int64))
		}
		if _EstimatedWorkMonths.Valid {
			t.SetEstimatedWorkMonths(_EstimatedWorkMonths.Float64)
		}
		if _EstimatedWorkMonths30Days.Valid {
			t.SetEstimatedWorkMonths30Days(_EstimatedWorkMonths30Days.Float64)
		}
		if _Title.Valid {
			t.SetTitle(_Title.String)
		}
		if _URL.Valid {
			t.SetURL(_URL.String)
		}
		if _Priority.Valid {
			t.SetPriority(_Priority.String)
		}
		if _PriorityID.Valid {
			t.SetPriorityID(_PriorityID.String)
		}
		if _Status.Valid {
			t.SetStatus(_Status.String)
		}
		if _StatusID.Valid {
			t.SetStatusID(_StatusID.String)
		}
		if _IssueType.Valid {
			t.SetIssueType(_IssueType.String)
		}
		if _IssueTypeID.Valid {
			t.SetIssueTypeID(_IssueTypeID.String)
		}
		if _Resolution.Valid {
			t.SetResolution(_Resolution.String)
		}
		if _ResolutionID.Valid {
			t.SetResolutionID(_ResolutionID.String)
		}
		if _State.Valid {
			t.SetState(_State.String)
		}
		if _CustomFieldIds.Valid {
			t.SetCustomFieldIds(_CustomFieldIds.String)
		}
		if _Teams.Valid {
			t.SetTeams(_Teams.String)
		}
		if _ParentIssueID.Valid {
			t.SetParentIssueID(_ParentIssueID.String)
		}
		if _ParentsIssueIds.Valid {
			t.SetParentsIssueIds(_ParentsIssueIds.String)
		}
		if _Metadata.Valid {
			t.SetMetadata(_Metadata.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _Sprints.Valid {
			t.SetSprints(_Sprints.String)
		}
		if _Labels.Valid {
			t.SetLabels(_Labels.String)
		}
		if _TopLevel.Valid {
			t.SetTopLevel(_TopLevel.Bool)
		}
		if _IsLeaf.Valid {
			t.SetIsLeaf(_IsLeaf.Bool)
		}
		if _Path.Valid {
			t.SetPath(_Path.String)
		}
		if _InProgressDuration.Valid {
			t.SetInProgressDuration(_InProgressDuration.Int64)
		}
		if _VerificationDuration.Valid {
			t.SetVerificationDuration(_VerificationDuration.Int64)
		}
		if _InProgressCount.Valid {
			t.SetInProgressCount(int32(_InProgressCount.Int64))
		}
		if _ReopenCount.Valid {
			t.SetReopenCount(int32(_ReopenCount.Int64))
		}
		if _MappedType.Valid {
			t.SetMappedType(_MappedType.String)
		}
		if _StrategicParentID.Valid {
			t.SetStrategicParentID(_StrategicParentID.String)
		}
		if _SprintID.Valid {
			t.SetSprintID(_SprintID.String)
		}
		if _IssueProjectName.Valid {
			t.SetIssueProjectName(_IssueProjectName.String)
		}
		if _Users.Valid {
			t.SetUsers(_Users.String)
		}
		if _InitialStartDate.Valid {
			t.SetInitialStartDate(_InitialStartDate.Int64)
		}
		if _TotalDuration.Valid {
			t.SetTotalDuration(_TotalDuration.Int64)
		}
		if _CreatedAt.Valid {
			t.SetCreatedAt(_CreatedAt.Int64)
		}
		if _ClosedAt.Valid {
			t.SetClosedAt(_ClosedAt.Int64)
		}
		if _PlannedEndDate.Valid {
			t.SetPlannedEndDate(_PlannedEndDate.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefType.Valid {
			t.SetRefType(_RefType.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		if _CustomFieldIdsVirtual.Valid {
			t.SetCustomFieldIdsVirtual(_CustomFieldIdsVirtual.String)
		}
		if _ReleaseDuration.Valid {
			t.SetReleaseDuration(_ReleaseDuration.Int64)
		}
		if _Completed.Valid {
			t.SetCompleted(_Completed.Bool)
		}
		if _CompletedDate.Valid {
			t.SetCompletedDate(_CompletedDate.Int64)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetSprintID will return the IssueSummary SprintID value
func (t *IssueSummary) GetSprintID() string {
	if t.SprintID == nil {
		return ""
	}
	return *t.SprintID
}

// SetSprintID will set the IssueSummary SprintID value
func (t *IssueSummary) SetSprintID(v string) {
	t.SprintID = &v
}

// GetIssueProjectName will return the IssueSummary IssueProjectName value
func (t *IssueSummary) GetIssueProjectName() string {
	return t.IssueProjectName
}

// SetIssueProjectName will set the IssueSummary IssueProjectName value
func (t *IssueSummary) SetIssueProjectName(v string) {
	t.IssueProjectName = v
}

// GetUsers will return the IssueSummary Users value
func (t *IssueSummary) GetUsers() string {
	if t.Users == nil {
		return ""
	}
	return *t.Users
}

// SetUsers will set the IssueSummary Users value
func (t *IssueSummary) SetUsers(v string) {
	t.Users = &v
}

// GetInitialStartDate will return the IssueSummary InitialStartDate value
func (t *IssueSummary) GetInitialStartDate() int64 {
	return t.InitialStartDate
}

// SetInitialStartDate will set the IssueSummary InitialStartDate value
func (t *IssueSummary) SetInitialStartDate(v int64) {
	t.InitialStartDate = v
}

// GetTotalDuration will return the IssueSummary TotalDuration value
func (t *IssueSummary) GetTotalDuration() int64 {
	return t.TotalDuration
}

// SetTotalDuration will set the IssueSummary TotalDuration value
func (t *IssueSummary) SetTotalDuration(v int64) {
	t.TotalDuration = v
}

// GetCreatedAt will return the IssueSummary CreatedAt value
func (t *IssueSummary) GetCreatedAt() int64 {
	if t.CreatedAt == nil {
		return int64(0)
	}
	return *t.CreatedAt
}

// SetCreatedAt will set the IssueSummary CreatedAt value
func (t *IssueSummary) SetCreatedAt(v int64) {
	t.CreatedAt = &v
}

// GetClosedAt will return the IssueSummary ClosedAt value
func (t *IssueSummary) GetClosedAt() int64 {
	if t.ClosedAt == nil {
		return int64(0)
	}
	return *t.ClosedAt
}

// SetClosedAt will set the IssueSummary ClosedAt value
func (t *IssueSummary) SetClosedAt(v int64) {
	t.ClosedAt = &v
}

// GetPlannedEndDate will return the IssueSummary PlannedEndDate value
func (t *IssueSummary) GetPlannedEndDate() int64 {
	if t.PlannedEndDate == nil {
		return int64(0)
	}
	return *t.PlannedEndDate
}

// SetPlannedEndDate will set the IssueSummary PlannedEndDate value
func (t *IssueSummary) SetPlannedEndDate(v int64) {
	t.PlannedEndDate = &v
}

// GetCustomerID will return the IssueSummary CustomerID value
func (t *IssueSummary) GetCustomerID() string {
	return t.CustomerID
}

// SetCustomerID will set the IssueSummary CustomerID value
func (t *IssueSummary) SetCustomerID(v string) {
	t.CustomerID = v
}

// FindIssueSummariesByCustomerID will find all IssueSummarys by the CustomerID value
func FindIssueSummariesByCustomerID(ctx context.Context, db DB, value string) ([]*IssueSummary, error) {
	q := "SELECT `issue_summary`.`id`,`issue_summary`.`checksum`,`issue_summary`.`issue_id`,`issue_summary`.`total_issues`,`issue_summary`.`new30_days`,`issue_summary`.`total_closed`,`issue_summary`.`closed30_days`,`issue_summary`.`estimated_work_months`,`issue_summary`.`estimated_work_months30_days`,`issue_summary`.`title`,`issue_summary`.`url`,`issue_summary`.`priority`,`issue_summary`.`priority_id`,`issue_summary`.`status`,`issue_summary`.`status_id`,`issue_summary`.`issue_type`,`issue_summary`.`issue_type_id`,`issue_summary`.`resolution`,`issue_summary`.`resolution_id`,`issue_summary`.`state`,`issue_summary`.`custom_field_ids`,`issue_summary`.`teams`,`issue_summary`.`parent_issue_id`,`issue_summary`.`parents_issue_ids`,`issue_summary`.`metadata`,`issue_summary`.`project_id`,`issue_summary`.`sprints`,`issue_summary`.`labels`,`issue_summary`.`top_level`,`issue_summary`.`is_leaf`,`issue_summary`.`path`,`issue_summary`.`in_progress_duration`,`issue_summary`.`verification_duration`,`issue_summary`.`in_progress_count`,`issue_summary`.`reopen_count`,`issue_summary`.`mapped_type`,`issue_summary`.`strategic_parent_id`,`issue_summary`.`sprint_id`,`issue_summary`.`issue_project_name`,`issue_summary`.`users`,`issue_summary`.`initial_start_date`,`issue_summary`.`total_duration`,`issue_summary`.`created_at`,`issue_summary`.`closed_at`,`issue_summary`.`planned_end_date`,`issue_summary`.`customer_id`,`issue_summary`.`ref_type`,`issue_summary`.`ref_id`,`issue_summary`.`custom_field_ids_virtual`,`issue_summary`.`release_duration`,`issue_summary`.`completed`,`issue_summary`.`completed_date` FROM `issue_summary` WHERE `customer_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*IssueSummary, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _IssueID sql.NullString
		var _TotalIssues sql.NullInt64
		var _New30Days sql.NullInt64
		var _TotalClosed sql.NullInt64
		var _Closed30Days sql.NullInt64
		var _EstimatedWorkMonths sql.NullFloat64
		var _EstimatedWorkMonths30Days sql.NullFloat64
		var _Title sql.NullString
		var _URL sql.NullString
		var _Priority sql.NullString
		var _PriorityID sql.NullString
		var _Status sql.NullString
		var _StatusID sql.NullString
		var _IssueType sql.NullString
		var _IssueTypeID sql.NullString
		var _Resolution sql.NullString
		var _ResolutionID sql.NullString
		var _State sql.NullString
		var _CustomFieldIds sql.NullString
		var _Teams sql.NullString
		var _ParentIssueID sql.NullString
		var _ParentsIssueIds sql.NullString
		var _Metadata sql.NullString
		var _ProjectID sql.NullString
		var _Sprints sql.NullString
		var _Labels sql.NullString
		var _TopLevel sql.NullBool
		var _IsLeaf sql.NullBool
		var _Path sql.NullString
		var _InProgressDuration sql.NullInt64
		var _VerificationDuration sql.NullInt64
		var _InProgressCount sql.NullInt64
		var _ReopenCount sql.NullInt64
		var _MappedType sql.NullString
		var _StrategicParentID sql.NullString
		var _SprintID sql.NullString
		var _IssueProjectName sql.NullString
		var _Users sql.NullString
		var _InitialStartDate sql.NullInt64
		var _TotalDuration sql.NullInt64
		var _CreatedAt sql.NullInt64
		var _ClosedAt sql.NullInt64
		var _PlannedEndDate sql.NullInt64
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefID sql.NullString
		var _CustomFieldIdsVirtual sql.NullString
		var _ReleaseDuration sql.NullInt64
		var _Completed sql.NullBool
		var _CompletedDate sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_IssueID,
			&_TotalIssues,
			&_New30Days,
			&_TotalClosed,
			&_Closed30Days,
			&_EstimatedWorkMonths,
			&_EstimatedWorkMonths30Days,
			&_Title,
			&_URL,
			&_Priority,
			&_PriorityID,
			&_Status,
			&_StatusID,
			&_IssueType,
			&_IssueTypeID,
			&_Resolution,
			&_ResolutionID,
			&_State,
			&_CustomFieldIds,
			&_Teams,
			&_ParentIssueID,
			&_ParentsIssueIds,
			&_Metadata,
			&_ProjectID,
			&_Sprints,
			&_Labels,
			&_TopLevel,
			&_IsLeaf,
			&_Path,
			&_InProgressDuration,
			&_VerificationDuration,
			&_InProgressCount,
			&_ReopenCount,
			&_MappedType,
			&_StrategicParentID,
			&_SprintID,
			&_IssueProjectName,
			&_Users,
			&_InitialStartDate,
			&_TotalDuration,
			&_CreatedAt,
			&_ClosedAt,
			&_PlannedEndDate,
			&_CustomerID,
			&_RefType,
			&_RefID,
			&_CustomFieldIdsVirtual,
			&_ReleaseDuration,
			&_Completed,
			&_CompletedDate,
		)
		if err != nil {
			return nil, err
		}
		t := &IssueSummary{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _TotalIssues.Valid {
			t.SetTotalIssues(int32(_TotalIssues.Int64))
		}
		if _New30Days.Valid {
			t.SetNew30Days(int32(_New30Days.Int64))
		}
		if _TotalClosed.Valid {
			t.SetTotalClosed(int32(_TotalClosed.Int64))
		}
		if _Closed30Days.Valid {
			t.SetClosed30Days(int32(_Closed30Days.Int64))
		}
		if _EstimatedWorkMonths.Valid {
			t.SetEstimatedWorkMonths(_EstimatedWorkMonths.Float64)
		}
		if _EstimatedWorkMonths30Days.Valid {
			t.SetEstimatedWorkMonths30Days(_EstimatedWorkMonths30Days.Float64)
		}
		if _Title.Valid {
			t.SetTitle(_Title.String)
		}
		if _URL.Valid {
			t.SetURL(_URL.String)
		}
		if _Priority.Valid {
			t.SetPriority(_Priority.String)
		}
		if _PriorityID.Valid {
			t.SetPriorityID(_PriorityID.String)
		}
		if _Status.Valid {
			t.SetStatus(_Status.String)
		}
		if _StatusID.Valid {
			t.SetStatusID(_StatusID.String)
		}
		if _IssueType.Valid {
			t.SetIssueType(_IssueType.String)
		}
		if _IssueTypeID.Valid {
			t.SetIssueTypeID(_IssueTypeID.String)
		}
		if _Resolution.Valid {
			t.SetResolution(_Resolution.String)
		}
		if _ResolutionID.Valid {
			t.SetResolutionID(_ResolutionID.String)
		}
		if _State.Valid {
			t.SetState(_State.String)
		}
		if _CustomFieldIds.Valid {
			t.SetCustomFieldIds(_CustomFieldIds.String)
		}
		if _Teams.Valid {
			t.SetTeams(_Teams.String)
		}
		if _ParentIssueID.Valid {
			t.SetParentIssueID(_ParentIssueID.String)
		}
		if _ParentsIssueIds.Valid {
			t.SetParentsIssueIds(_ParentsIssueIds.String)
		}
		if _Metadata.Valid {
			t.SetMetadata(_Metadata.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _Sprints.Valid {
			t.SetSprints(_Sprints.String)
		}
		if _Labels.Valid {
			t.SetLabels(_Labels.String)
		}
		if _TopLevel.Valid {
			t.SetTopLevel(_TopLevel.Bool)
		}
		if _IsLeaf.Valid {
			t.SetIsLeaf(_IsLeaf.Bool)
		}
		if _Path.Valid {
			t.SetPath(_Path.String)
		}
		if _InProgressDuration.Valid {
			t.SetInProgressDuration(_InProgressDuration.Int64)
		}
		if _VerificationDuration.Valid {
			t.SetVerificationDuration(_VerificationDuration.Int64)
		}
		if _InProgressCount.Valid {
			t.SetInProgressCount(int32(_InProgressCount.Int64))
		}
		if _ReopenCount.Valid {
			t.SetReopenCount(int32(_ReopenCount.Int64))
		}
		if _MappedType.Valid {
			t.SetMappedType(_MappedType.String)
		}
		if _StrategicParentID.Valid {
			t.SetStrategicParentID(_StrategicParentID.String)
		}
		if _SprintID.Valid {
			t.SetSprintID(_SprintID.String)
		}
		if _IssueProjectName.Valid {
			t.SetIssueProjectName(_IssueProjectName.String)
		}
		if _Users.Valid {
			t.SetUsers(_Users.String)
		}
		if _InitialStartDate.Valid {
			t.SetInitialStartDate(_InitialStartDate.Int64)
		}
		if _TotalDuration.Valid {
			t.SetTotalDuration(_TotalDuration.Int64)
		}
		if _CreatedAt.Valid {
			t.SetCreatedAt(_CreatedAt.Int64)
		}
		if _ClosedAt.Valid {
			t.SetClosedAt(_ClosedAt.Int64)
		}
		if _PlannedEndDate.Valid {
			t.SetPlannedEndDate(_PlannedEndDate.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefType.Valid {
			t.SetRefType(_RefType.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		if _CustomFieldIdsVirtual.Valid {
			t.SetCustomFieldIdsVirtual(_CustomFieldIdsVirtual.String)
		}
		if _ReleaseDuration.Valid {
			t.SetReleaseDuration(_ReleaseDuration.Int64)
		}
		if _Completed.Valid {
			t.SetCompleted(_Completed.Bool)
		}
		if _CompletedDate.Valid {
			t.SetCompletedDate(_CompletedDate.Int64)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindIssueSummariesByCustomerIDTx will find all IssueSummarys by the CustomerID value using the provided transaction
func FindIssueSummariesByCustomerIDTx(ctx context.Context, tx Tx, value string) ([]*IssueSummary, error) {
	q := "SELECT `issue_summary`.`id`,`issue_summary`.`checksum`,`issue_summary`.`issue_id`,`issue_summary`.`total_issues`,`issue_summary`.`new30_days`,`issue_summary`.`total_closed`,`issue_summary`.`closed30_days`,`issue_summary`.`estimated_work_months`,`issue_summary`.`estimated_work_months30_days`,`issue_summary`.`title`,`issue_summary`.`url`,`issue_summary`.`priority`,`issue_summary`.`priority_id`,`issue_summary`.`status`,`issue_summary`.`status_id`,`issue_summary`.`issue_type`,`issue_summary`.`issue_type_id`,`issue_summary`.`resolution`,`issue_summary`.`resolution_id`,`issue_summary`.`state`,`issue_summary`.`custom_field_ids`,`issue_summary`.`teams`,`issue_summary`.`parent_issue_id`,`issue_summary`.`parents_issue_ids`,`issue_summary`.`metadata`,`issue_summary`.`project_id`,`issue_summary`.`sprints`,`issue_summary`.`labels`,`issue_summary`.`top_level`,`issue_summary`.`is_leaf`,`issue_summary`.`path`,`issue_summary`.`in_progress_duration`,`issue_summary`.`verification_duration`,`issue_summary`.`in_progress_count`,`issue_summary`.`reopen_count`,`issue_summary`.`mapped_type`,`issue_summary`.`strategic_parent_id`,`issue_summary`.`sprint_id`,`issue_summary`.`issue_project_name`,`issue_summary`.`users`,`issue_summary`.`initial_start_date`,`issue_summary`.`total_duration`,`issue_summary`.`created_at`,`issue_summary`.`closed_at`,`issue_summary`.`planned_end_date`,`issue_summary`.`customer_id`,`issue_summary`.`ref_type`,`issue_summary`.`ref_id`,`issue_summary`.`custom_field_ids_virtual`,`issue_summary`.`release_duration`,`issue_summary`.`completed`,`issue_summary`.`completed_date` FROM `issue_summary` WHERE `customer_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*IssueSummary, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _IssueID sql.NullString
		var _TotalIssues sql.NullInt64
		var _New30Days sql.NullInt64
		var _TotalClosed sql.NullInt64
		var _Closed30Days sql.NullInt64
		var _EstimatedWorkMonths sql.NullFloat64
		var _EstimatedWorkMonths30Days sql.NullFloat64
		var _Title sql.NullString
		var _URL sql.NullString
		var _Priority sql.NullString
		var _PriorityID sql.NullString
		var _Status sql.NullString
		var _StatusID sql.NullString
		var _IssueType sql.NullString
		var _IssueTypeID sql.NullString
		var _Resolution sql.NullString
		var _ResolutionID sql.NullString
		var _State sql.NullString
		var _CustomFieldIds sql.NullString
		var _Teams sql.NullString
		var _ParentIssueID sql.NullString
		var _ParentsIssueIds sql.NullString
		var _Metadata sql.NullString
		var _ProjectID sql.NullString
		var _Sprints sql.NullString
		var _Labels sql.NullString
		var _TopLevel sql.NullBool
		var _IsLeaf sql.NullBool
		var _Path sql.NullString
		var _InProgressDuration sql.NullInt64
		var _VerificationDuration sql.NullInt64
		var _InProgressCount sql.NullInt64
		var _ReopenCount sql.NullInt64
		var _MappedType sql.NullString
		var _StrategicParentID sql.NullString
		var _SprintID sql.NullString
		var _IssueProjectName sql.NullString
		var _Users sql.NullString
		var _InitialStartDate sql.NullInt64
		var _TotalDuration sql.NullInt64
		var _CreatedAt sql.NullInt64
		var _ClosedAt sql.NullInt64
		var _PlannedEndDate sql.NullInt64
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefID sql.NullString
		var _CustomFieldIdsVirtual sql.NullString
		var _ReleaseDuration sql.NullInt64
		var _Completed sql.NullBool
		var _CompletedDate sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_IssueID,
			&_TotalIssues,
			&_New30Days,
			&_TotalClosed,
			&_Closed30Days,
			&_EstimatedWorkMonths,
			&_EstimatedWorkMonths30Days,
			&_Title,
			&_URL,
			&_Priority,
			&_PriorityID,
			&_Status,
			&_StatusID,
			&_IssueType,
			&_IssueTypeID,
			&_Resolution,
			&_ResolutionID,
			&_State,
			&_CustomFieldIds,
			&_Teams,
			&_ParentIssueID,
			&_ParentsIssueIds,
			&_Metadata,
			&_ProjectID,
			&_Sprints,
			&_Labels,
			&_TopLevel,
			&_IsLeaf,
			&_Path,
			&_InProgressDuration,
			&_VerificationDuration,
			&_InProgressCount,
			&_ReopenCount,
			&_MappedType,
			&_StrategicParentID,
			&_SprintID,
			&_IssueProjectName,
			&_Users,
			&_InitialStartDate,
			&_TotalDuration,
			&_CreatedAt,
			&_ClosedAt,
			&_PlannedEndDate,
			&_CustomerID,
			&_RefType,
			&_RefID,
			&_CustomFieldIdsVirtual,
			&_ReleaseDuration,
			&_Completed,
			&_CompletedDate,
		)
		if err != nil {
			return nil, err
		}
		t := &IssueSummary{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _TotalIssues.Valid {
			t.SetTotalIssues(int32(_TotalIssues.Int64))
		}
		if _New30Days.Valid {
			t.SetNew30Days(int32(_New30Days.Int64))
		}
		if _TotalClosed.Valid {
			t.SetTotalClosed(int32(_TotalClosed.Int64))
		}
		if _Closed30Days.Valid {
			t.SetClosed30Days(int32(_Closed30Days.Int64))
		}
		if _EstimatedWorkMonths.Valid {
			t.SetEstimatedWorkMonths(_EstimatedWorkMonths.Float64)
		}
		if _EstimatedWorkMonths30Days.Valid {
			t.SetEstimatedWorkMonths30Days(_EstimatedWorkMonths30Days.Float64)
		}
		if _Title.Valid {
			t.SetTitle(_Title.String)
		}
		if _URL.Valid {
			t.SetURL(_URL.String)
		}
		if _Priority.Valid {
			t.SetPriority(_Priority.String)
		}
		if _PriorityID.Valid {
			t.SetPriorityID(_PriorityID.String)
		}
		if _Status.Valid {
			t.SetStatus(_Status.String)
		}
		if _StatusID.Valid {
			t.SetStatusID(_StatusID.String)
		}
		if _IssueType.Valid {
			t.SetIssueType(_IssueType.String)
		}
		if _IssueTypeID.Valid {
			t.SetIssueTypeID(_IssueTypeID.String)
		}
		if _Resolution.Valid {
			t.SetResolution(_Resolution.String)
		}
		if _ResolutionID.Valid {
			t.SetResolutionID(_ResolutionID.String)
		}
		if _State.Valid {
			t.SetState(_State.String)
		}
		if _CustomFieldIds.Valid {
			t.SetCustomFieldIds(_CustomFieldIds.String)
		}
		if _Teams.Valid {
			t.SetTeams(_Teams.String)
		}
		if _ParentIssueID.Valid {
			t.SetParentIssueID(_ParentIssueID.String)
		}
		if _ParentsIssueIds.Valid {
			t.SetParentsIssueIds(_ParentsIssueIds.String)
		}
		if _Metadata.Valid {
			t.SetMetadata(_Metadata.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _Sprints.Valid {
			t.SetSprints(_Sprints.String)
		}
		if _Labels.Valid {
			t.SetLabels(_Labels.String)
		}
		if _TopLevel.Valid {
			t.SetTopLevel(_TopLevel.Bool)
		}
		if _IsLeaf.Valid {
			t.SetIsLeaf(_IsLeaf.Bool)
		}
		if _Path.Valid {
			t.SetPath(_Path.String)
		}
		if _InProgressDuration.Valid {
			t.SetInProgressDuration(_InProgressDuration.Int64)
		}
		if _VerificationDuration.Valid {
			t.SetVerificationDuration(_VerificationDuration.Int64)
		}
		if _InProgressCount.Valid {
			t.SetInProgressCount(int32(_InProgressCount.Int64))
		}
		if _ReopenCount.Valid {
			t.SetReopenCount(int32(_ReopenCount.Int64))
		}
		if _MappedType.Valid {
			t.SetMappedType(_MappedType.String)
		}
		if _StrategicParentID.Valid {
			t.SetStrategicParentID(_StrategicParentID.String)
		}
		if _SprintID.Valid {
			t.SetSprintID(_SprintID.String)
		}
		if _IssueProjectName.Valid {
			t.SetIssueProjectName(_IssueProjectName.String)
		}
		if _Users.Valid {
			t.SetUsers(_Users.String)
		}
		if _InitialStartDate.Valid {
			t.SetInitialStartDate(_InitialStartDate.Int64)
		}
		if _TotalDuration.Valid {
			t.SetTotalDuration(_TotalDuration.Int64)
		}
		if _CreatedAt.Valid {
			t.SetCreatedAt(_CreatedAt.Int64)
		}
		if _ClosedAt.Valid {
			t.SetClosedAt(_ClosedAt.Int64)
		}
		if _PlannedEndDate.Valid {
			t.SetPlannedEndDate(_PlannedEndDate.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefType.Valid {
			t.SetRefType(_RefType.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		if _CustomFieldIdsVirtual.Valid {
			t.SetCustomFieldIdsVirtual(_CustomFieldIdsVirtual.String)
		}
		if _ReleaseDuration.Valid {
			t.SetReleaseDuration(_ReleaseDuration.Int64)
		}
		if _Completed.Valid {
			t.SetCompleted(_Completed.Bool)
		}
		if _CompletedDate.Valid {
			t.SetCompletedDate(_CompletedDate.Int64)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetRefType will return the IssueSummary RefType value
func (t *IssueSummary) GetRefType() string {
	return t.RefType
}

// SetRefType will set the IssueSummary RefType value
func (t *IssueSummary) SetRefType(v string) {
	t.RefType = v
}

// GetRefID will return the IssueSummary RefID value
func (t *IssueSummary) GetRefID() string {
	return t.RefID
}

// SetRefID will set the IssueSummary RefID value
func (t *IssueSummary) SetRefID(v string) {
	t.RefID = v
}

// FindIssueSummariesByRefID will find all IssueSummarys by the RefID value
func FindIssueSummariesByRefID(ctx context.Context, db DB, value string) ([]*IssueSummary, error) {
	q := "SELECT `issue_summary`.`id`,`issue_summary`.`checksum`,`issue_summary`.`issue_id`,`issue_summary`.`total_issues`,`issue_summary`.`new30_days`,`issue_summary`.`total_closed`,`issue_summary`.`closed30_days`,`issue_summary`.`estimated_work_months`,`issue_summary`.`estimated_work_months30_days`,`issue_summary`.`title`,`issue_summary`.`url`,`issue_summary`.`priority`,`issue_summary`.`priority_id`,`issue_summary`.`status`,`issue_summary`.`status_id`,`issue_summary`.`issue_type`,`issue_summary`.`issue_type_id`,`issue_summary`.`resolution`,`issue_summary`.`resolution_id`,`issue_summary`.`state`,`issue_summary`.`custom_field_ids`,`issue_summary`.`teams`,`issue_summary`.`parent_issue_id`,`issue_summary`.`parents_issue_ids`,`issue_summary`.`metadata`,`issue_summary`.`project_id`,`issue_summary`.`sprints`,`issue_summary`.`labels`,`issue_summary`.`top_level`,`issue_summary`.`is_leaf`,`issue_summary`.`path`,`issue_summary`.`in_progress_duration`,`issue_summary`.`verification_duration`,`issue_summary`.`in_progress_count`,`issue_summary`.`reopen_count`,`issue_summary`.`mapped_type`,`issue_summary`.`strategic_parent_id`,`issue_summary`.`sprint_id`,`issue_summary`.`issue_project_name`,`issue_summary`.`users`,`issue_summary`.`initial_start_date`,`issue_summary`.`total_duration`,`issue_summary`.`created_at`,`issue_summary`.`closed_at`,`issue_summary`.`planned_end_date`,`issue_summary`.`customer_id`,`issue_summary`.`ref_type`,`issue_summary`.`ref_id`,`issue_summary`.`custom_field_ids_virtual`,`issue_summary`.`release_duration`,`issue_summary`.`completed`,`issue_summary`.`completed_date` FROM `issue_summary` WHERE `ref_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*IssueSummary, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _IssueID sql.NullString
		var _TotalIssues sql.NullInt64
		var _New30Days sql.NullInt64
		var _TotalClosed sql.NullInt64
		var _Closed30Days sql.NullInt64
		var _EstimatedWorkMonths sql.NullFloat64
		var _EstimatedWorkMonths30Days sql.NullFloat64
		var _Title sql.NullString
		var _URL sql.NullString
		var _Priority sql.NullString
		var _PriorityID sql.NullString
		var _Status sql.NullString
		var _StatusID sql.NullString
		var _IssueType sql.NullString
		var _IssueTypeID sql.NullString
		var _Resolution sql.NullString
		var _ResolutionID sql.NullString
		var _State sql.NullString
		var _CustomFieldIds sql.NullString
		var _Teams sql.NullString
		var _ParentIssueID sql.NullString
		var _ParentsIssueIds sql.NullString
		var _Metadata sql.NullString
		var _ProjectID sql.NullString
		var _Sprints sql.NullString
		var _Labels sql.NullString
		var _TopLevel sql.NullBool
		var _IsLeaf sql.NullBool
		var _Path sql.NullString
		var _InProgressDuration sql.NullInt64
		var _VerificationDuration sql.NullInt64
		var _InProgressCount sql.NullInt64
		var _ReopenCount sql.NullInt64
		var _MappedType sql.NullString
		var _StrategicParentID sql.NullString
		var _SprintID sql.NullString
		var _IssueProjectName sql.NullString
		var _Users sql.NullString
		var _InitialStartDate sql.NullInt64
		var _TotalDuration sql.NullInt64
		var _CreatedAt sql.NullInt64
		var _ClosedAt sql.NullInt64
		var _PlannedEndDate sql.NullInt64
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefID sql.NullString
		var _CustomFieldIdsVirtual sql.NullString
		var _ReleaseDuration sql.NullInt64
		var _Completed sql.NullBool
		var _CompletedDate sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_IssueID,
			&_TotalIssues,
			&_New30Days,
			&_TotalClosed,
			&_Closed30Days,
			&_EstimatedWorkMonths,
			&_EstimatedWorkMonths30Days,
			&_Title,
			&_URL,
			&_Priority,
			&_PriorityID,
			&_Status,
			&_StatusID,
			&_IssueType,
			&_IssueTypeID,
			&_Resolution,
			&_ResolutionID,
			&_State,
			&_CustomFieldIds,
			&_Teams,
			&_ParentIssueID,
			&_ParentsIssueIds,
			&_Metadata,
			&_ProjectID,
			&_Sprints,
			&_Labels,
			&_TopLevel,
			&_IsLeaf,
			&_Path,
			&_InProgressDuration,
			&_VerificationDuration,
			&_InProgressCount,
			&_ReopenCount,
			&_MappedType,
			&_StrategicParentID,
			&_SprintID,
			&_IssueProjectName,
			&_Users,
			&_InitialStartDate,
			&_TotalDuration,
			&_CreatedAt,
			&_ClosedAt,
			&_PlannedEndDate,
			&_CustomerID,
			&_RefType,
			&_RefID,
			&_CustomFieldIdsVirtual,
			&_ReleaseDuration,
			&_Completed,
			&_CompletedDate,
		)
		if err != nil {
			return nil, err
		}
		t := &IssueSummary{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _TotalIssues.Valid {
			t.SetTotalIssues(int32(_TotalIssues.Int64))
		}
		if _New30Days.Valid {
			t.SetNew30Days(int32(_New30Days.Int64))
		}
		if _TotalClosed.Valid {
			t.SetTotalClosed(int32(_TotalClosed.Int64))
		}
		if _Closed30Days.Valid {
			t.SetClosed30Days(int32(_Closed30Days.Int64))
		}
		if _EstimatedWorkMonths.Valid {
			t.SetEstimatedWorkMonths(_EstimatedWorkMonths.Float64)
		}
		if _EstimatedWorkMonths30Days.Valid {
			t.SetEstimatedWorkMonths30Days(_EstimatedWorkMonths30Days.Float64)
		}
		if _Title.Valid {
			t.SetTitle(_Title.String)
		}
		if _URL.Valid {
			t.SetURL(_URL.String)
		}
		if _Priority.Valid {
			t.SetPriority(_Priority.String)
		}
		if _PriorityID.Valid {
			t.SetPriorityID(_PriorityID.String)
		}
		if _Status.Valid {
			t.SetStatus(_Status.String)
		}
		if _StatusID.Valid {
			t.SetStatusID(_StatusID.String)
		}
		if _IssueType.Valid {
			t.SetIssueType(_IssueType.String)
		}
		if _IssueTypeID.Valid {
			t.SetIssueTypeID(_IssueTypeID.String)
		}
		if _Resolution.Valid {
			t.SetResolution(_Resolution.String)
		}
		if _ResolutionID.Valid {
			t.SetResolutionID(_ResolutionID.String)
		}
		if _State.Valid {
			t.SetState(_State.String)
		}
		if _CustomFieldIds.Valid {
			t.SetCustomFieldIds(_CustomFieldIds.String)
		}
		if _Teams.Valid {
			t.SetTeams(_Teams.String)
		}
		if _ParentIssueID.Valid {
			t.SetParentIssueID(_ParentIssueID.String)
		}
		if _ParentsIssueIds.Valid {
			t.SetParentsIssueIds(_ParentsIssueIds.String)
		}
		if _Metadata.Valid {
			t.SetMetadata(_Metadata.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _Sprints.Valid {
			t.SetSprints(_Sprints.String)
		}
		if _Labels.Valid {
			t.SetLabels(_Labels.String)
		}
		if _TopLevel.Valid {
			t.SetTopLevel(_TopLevel.Bool)
		}
		if _IsLeaf.Valid {
			t.SetIsLeaf(_IsLeaf.Bool)
		}
		if _Path.Valid {
			t.SetPath(_Path.String)
		}
		if _InProgressDuration.Valid {
			t.SetInProgressDuration(_InProgressDuration.Int64)
		}
		if _VerificationDuration.Valid {
			t.SetVerificationDuration(_VerificationDuration.Int64)
		}
		if _InProgressCount.Valid {
			t.SetInProgressCount(int32(_InProgressCount.Int64))
		}
		if _ReopenCount.Valid {
			t.SetReopenCount(int32(_ReopenCount.Int64))
		}
		if _MappedType.Valid {
			t.SetMappedType(_MappedType.String)
		}
		if _StrategicParentID.Valid {
			t.SetStrategicParentID(_StrategicParentID.String)
		}
		if _SprintID.Valid {
			t.SetSprintID(_SprintID.String)
		}
		if _IssueProjectName.Valid {
			t.SetIssueProjectName(_IssueProjectName.String)
		}
		if _Users.Valid {
			t.SetUsers(_Users.String)
		}
		if _InitialStartDate.Valid {
			t.SetInitialStartDate(_InitialStartDate.Int64)
		}
		if _TotalDuration.Valid {
			t.SetTotalDuration(_TotalDuration.Int64)
		}
		if _CreatedAt.Valid {
			t.SetCreatedAt(_CreatedAt.Int64)
		}
		if _ClosedAt.Valid {
			t.SetClosedAt(_ClosedAt.Int64)
		}
		if _PlannedEndDate.Valid {
			t.SetPlannedEndDate(_PlannedEndDate.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefType.Valid {
			t.SetRefType(_RefType.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		if _CustomFieldIdsVirtual.Valid {
			t.SetCustomFieldIdsVirtual(_CustomFieldIdsVirtual.String)
		}
		if _ReleaseDuration.Valid {
			t.SetReleaseDuration(_ReleaseDuration.Int64)
		}
		if _Completed.Valid {
			t.SetCompleted(_Completed.Bool)
		}
		if _CompletedDate.Valid {
			t.SetCompletedDate(_CompletedDate.Int64)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindIssueSummariesByRefIDTx will find all IssueSummarys by the RefID value using the provided transaction
func FindIssueSummariesByRefIDTx(ctx context.Context, tx Tx, value string) ([]*IssueSummary, error) {
	q := "SELECT `issue_summary`.`id`,`issue_summary`.`checksum`,`issue_summary`.`issue_id`,`issue_summary`.`total_issues`,`issue_summary`.`new30_days`,`issue_summary`.`total_closed`,`issue_summary`.`closed30_days`,`issue_summary`.`estimated_work_months`,`issue_summary`.`estimated_work_months30_days`,`issue_summary`.`title`,`issue_summary`.`url`,`issue_summary`.`priority`,`issue_summary`.`priority_id`,`issue_summary`.`status`,`issue_summary`.`status_id`,`issue_summary`.`issue_type`,`issue_summary`.`issue_type_id`,`issue_summary`.`resolution`,`issue_summary`.`resolution_id`,`issue_summary`.`state`,`issue_summary`.`custom_field_ids`,`issue_summary`.`teams`,`issue_summary`.`parent_issue_id`,`issue_summary`.`parents_issue_ids`,`issue_summary`.`metadata`,`issue_summary`.`project_id`,`issue_summary`.`sprints`,`issue_summary`.`labels`,`issue_summary`.`top_level`,`issue_summary`.`is_leaf`,`issue_summary`.`path`,`issue_summary`.`in_progress_duration`,`issue_summary`.`verification_duration`,`issue_summary`.`in_progress_count`,`issue_summary`.`reopen_count`,`issue_summary`.`mapped_type`,`issue_summary`.`strategic_parent_id`,`issue_summary`.`sprint_id`,`issue_summary`.`issue_project_name`,`issue_summary`.`users`,`issue_summary`.`initial_start_date`,`issue_summary`.`total_duration`,`issue_summary`.`created_at`,`issue_summary`.`closed_at`,`issue_summary`.`planned_end_date`,`issue_summary`.`customer_id`,`issue_summary`.`ref_type`,`issue_summary`.`ref_id`,`issue_summary`.`custom_field_ids_virtual`,`issue_summary`.`release_duration`,`issue_summary`.`completed`,`issue_summary`.`completed_date` FROM `issue_summary` WHERE `ref_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*IssueSummary, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _IssueID sql.NullString
		var _TotalIssues sql.NullInt64
		var _New30Days sql.NullInt64
		var _TotalClosed sql.NullInt64
		var _Closed30Days sql.NullInt64
		var _EstimatedWorkMonths sql.NullFloat64
		var _EstimatedWorkMonths30Days sql.NullFloat64
		var _Title sql.NullString
		var _URL sql.NullString
		var _Priority sql.NullString
		var _PriorityID sql.NullString
		var _Status sql.NullString
		var _StatusID sql.NullString
		var _IssueType sql.NullString
		var _IssueTypeID sql.NullString
		var _Resolution sql.NullString
		var _ResolutionID sql.NullString
		var _State sql.NullString
		var _CustomFieldIds sql.NullString
		var _Teams sql.NullString
		var _ParentIssueID sql.NullString
		var _ParentsIssueIds sql.NullString
		var _Metadata sql.NullString
		var _ProjectID sql.NullString
		var _Sprints sql.NullString
		var _Labels sql.NullString
		var _TopLevel sql.NullBool
		var _IsLeaf sql.NullBool
		var _Path sql.NullString
		var _InProgressDuration sql.NullInt64
		var _VerificationDuration sql.NullInt64
		var _InProgressCount sql.NullInt64
		var _ReopenCount sql.NullInt64
		var _MappedType sql.NullString
		var _StrategicParentID sql.NullString
		var _SprintID sql.NullString
		var _IssueProjectName sql.NullString
		var _Users sql.NullString
		var _InitialStartDate sql.NullInt64
		var _TotalDuration sql.NullInt64
		var _CreatedAt sql.NullInt64
		var _ClosedAt sql.NullInt64
		var _PlannedEndDate sql.NullInt64
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefID sql.NullString
		var _CustomFieldIdsVirtual sql.NullString
		var _ReleaseDuration sql.NullInt64
		var _Completed sql.NullBool
		var _CompletedDate sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_IssueID,
			&_TotalIssues,
			&_New30Days,
			&_TotalClosed,
			&_Closed30Days,
			&_EstimatedWorkMonths,
			&_EstimatedWorkMonths30Days,
			&_Title,
			&_URL,
			&_Priority,
			&_PriorityID,
			&_Status,
			&_StatusID,
			&_IssueType,
			&_IssueTypeID,
			&_Resolution,
			&_ResolutionID,
			&_State,
			&_CustomFieldIds,
			&_Teams,
			&_ParentIssueID,
			&_ParentsIssueIds,
			&_Metadata,
			&_ProjectID,
			&_Sprints,
			&_Labels,
			&_TopLevel,
			&_IsLeaf,
			&_Path,
			&_InProgressDuration,
			&_VerificationDuration,
			&_InProgressCount,
			&_ReopenCount,
			&_MappedType,
			&_StrategicParentID,
			&_SprintID,
			&_IssueProjectName,
			&_Users,
			&_InitialStartDate,
			&_TotalDuration,
			&_CreatedAt,
			&_ClosedAt,
			&_PlannedEndDate,
			&_CustomerID,
			&_RefType,
			&_RefID,
			&_CustomFieldIdsVirtual,
			&_ReleaseDuration,
			&_Completed,
			&_CompletedDate,
		)
		if err != nil {
			return nil, err
		}
		t := &IssueSummary{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _TotalIssues.Valid {
			t.SetTotalIssues(int32(_TotalIssues.Int64))
		}
		if _New30Days.Valid {
			t.SetNew30Days(int32(_New30Days.Int64))
		}
		if _TotalClosed.Valid {
			t.SetTotalClosed(int32(_TotalClosed.Int64))
		}
		if _Closed30Days.Valid {
			t.SetClosed30Days(int32(_Closed30Days.Int64))
		}
		if _EstimatedWorkMonths.Valid {
			t.SetEstimatedWorkMonths(_EstimatedWorkMonths.Float64)
		}
		if _EstimatedWorkMonths30Days.Valid {
			t.SetEstimatedWorkMonths30Days(_EstimatedWorkMonths30Days.Float64)
		}
		if _Title.Valid {
			t.SetTitle(_Title.String)
		}
		if _URL.Valid {
			t.SetURL(_URL.String)
		}
		if _Priority.Valid {
			t.SetPriority(_Priority.String)
		}
		if _PriorityID.Valid {
			t.SetPriorityID(_PriorityID.String)
		}
		if _Status.Valid {
			t.SetStatus(_Status.String)
		}
		if _StatusID.Valid {
			t.SetStatusID(_StatusID.String)
		}
		if _IssueType.Valid {
			t.SetIssueType(_IssueType.String)
		}
		if _IssueTypeID.Valid {
			t.SetIssueTypeID(_IssueTypeID.String)
		}
		if _Resolution.Valid {
			t.SetResolution(_Resolution.String)
		}
		if _ResolutionID.Valid {
			t.SetResolutionID(_ResolutionID.String)
		}
		if _State.Valid {
			t.SetState(_State.String)
		}
		if _CustomFieldIds.Valid {
			t.SetCustomFieldIds(_CustomFieldIds.String)
		}
		if _Teams.Valid {
			t.SetTeams(_Teams.String)
		}
		if _ParentIssueID.Valid {
			t.SetParentIssueID(_ParentIssueID.String)
		}
		if _ParentsIssueIds.Valid {
			t.SetParentsIssueIds(_ParentsIssueIds.String)
		}
		if _Metadata.Valid {
			t.SetMetadata(_Metadata.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _Sprints.Valid {
			t.SetSprints(_Sprints.String)
		}
		if _Labels.Valid {
			t.SetLabels(_Labels.String)
		}
		if _TopLevel.Valid {
			t.SetTopLevel(_TopLevel.Bool)
		}
		if _IsLeaf.Valid {
			t.SetIsLeaf(_IsLeaf.Bool)
		}
		if _Path.Valid {
			t.SetPath(_Path.String)
		}
		if _InProgressDuration.Valid {
			t.SetInProgressDuration(_InProgressDuration.Int64)
		}
		if _VerificationDuration.Valid {
			t.SetVerificationDuration(_VerificationDuration.Int64)
		}
		if _InProgressCount.Valid {
			t.SetInProgressCount(int32(_InProgressCount.Int64))
		}
		if _ReopenCount.Valid {
			t.SetReopenCount(int32(_ReopenCount.Int64))
		}
		if _MappedType.Valid {
			t.SetMappedType(_MappedType.String)
		}
		if _StrategicParentID.Valid {
			t.SetStrategicParentID(_StrategicParentID.String)
		}
		if _SprintID.Valid {
			t.SetSprintID(_SprintID.String)
		}
		if _IssueProjectName.Valid {
			t.SetIssueProjectName(_IssueProjectName.String)
		}
		if _Users.Valid {
			t.SetUsers(_Users.String)
		}
		if _InitialStartDate.Valid {
			t.SetInitialStartDate(_InitialStartDate.Int64)
		}
		if _TotalDuration.Valid {
			t.SetTotalDuration(_TotalDuration.Int64)
		}
		if _CreatedAt.Valid {
			t.SetCreatedAt(_CreatedAt.Int64)
		}
		if _ClosedAt.Valid {
			t.SetClosedAt(_ClosedAt.Int64)
		}
		if _PlannedEndDate.Valid {
			t.SetPlannedEndDate(_PlannedEndDate.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefType.Valid {
			t.SetRefType(_RefType.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		if _CustomFieldIdsVirtual.Valid {
			t.SetCustomFieldIdsVirtual(_CustomFieldIdsVirtual.String)
		}
		if _ReleaseDuration.Valid {
			t.SetReleaseDuration(_ReleaseDuration.Int64)
		}
		if _Completed.Valid {
			t.SetCompleted(_Completed.Bool)
		}
		if _CompletedDate.Valid {
			t.SetCompletedDate(_CompletedDate.Int64)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetCustomFieldIdsVirtual will return the IssueSummary CustomFieldIdsVirtual value
func (t *IssueSummary) GetCustomFieldIdsVirtual() string {
	if t.CustomFieldIdsVirtual == nil {
		return ""
	}
	return *t.CustomFieldIdsVirtual
}

// SetCustomFieldIdsVirtual will set the IssueSummary CustomFieldIdsVirtual value
func (t *IssueSummary) SetCustomFieldIdsVirtual(v string) {
	t.CustomFieldIdsVirtual = &v
}

// GetReleaseDuration will return the IssueSummary ReleaseDuration value
func (t *IssueSummary) GetReleaseDuration() int64 {
	return t.ReleaseDuration
}

// SetReleaseDuration will set the IssueSummary ReleaseDuration value
func (t *IssueSummary) SetReleaseDuration(v int64) {
	t.ReleaseDuration = v
}

// GetCompleted will return the IssueSummary Completed value
func (t *IssueSummary) GetCompleted() bool {
	return t.Completed
}

// SetCompleted will set the IssueSummary Completed value
func (t *IssueSummary) SetCompleted(v bool) {
	t.Completed = v
}

// GetCompletedDate will return the IssueSummary CompletedDate value
func (t *IssueSummary) GetCompletedDate() int64 {
	if t.CompletedDate == nil {
		return int64(0)
	}
	return *t.CompletedDate
}

// SetCompletedDate will set the IssueSummary CompletedDate value
func (t *IssueSummary) SetCompletedDate(v int64) {
	t.CompletedDate = &v
}

func (t *IssueSummary) toTimestamp(value time.Time) *timestamp.Timestamp {
	ts, _ := ptypes.TimestampProto(value)
	return ts
}

// DBCreateIssueSummaryTable will create the IssueSummary table
func DBCreateIssueSummaryTable(ctx context.Context, db DB) error {
	q := "CREATE TABLE `issue_summary` (`id` VARCHAR(64) NOT NULL PRIMARY KEY,`checksum` CHAR(64),`issue_id` VARCHAR(64) NOT NULL,`total_issues` INT UNSIGNED NOT NULL,`new30_days`INT UNSIGNED NOT NULL,`total_closed` INT UNSIGNED NOT NULL,`closed30_days`INT UNSIGNED NOT NULL,`estimated_work_months` FLOAT NOT NULL,`estimated_work_months30_days` FLOAT NOT NULL,`title` TEXT NOT NULL,`url` TEXT,`priority` VARCHAR(100),`priority_id` VARCHAR(64),`status` VARCHAR(100),`status_id` VARCHAR(64),`issue_type`VARCHAR(100) NOT NULL,`issue_type_id`VARCHAR(64),`resolution`VARCHAR(100),`resolution_id`VARCHAR(64),`state` VARCHAR(10) NOT NULL,`custom_field_ids`JSON,`teams` JSON,`parent_issue_id` VARCHAR(64),`parents_issue_ids` JSON,`metadata` JSON,`project_id`VARCHAR(64) NOT NULL,`sprints`JSON,`labels` JSON,`top_level` TINYINT UNSIGNED NOT NULL,`is_leaf`TINYINT UNSIGNED NOT NULL,`path`VARCHAR(1024) NOT NULL,`in_progress_duration` BIGINT NOT NULL,`verification_duration` BIGINT NOT NULL,`in_progress_count` INT NOT NULL,`reopen_count` INT NOT NULL,`mapped_type` VARCHAR(75) NOT NULL,`strategic_parent_id`VARCHAR(64),`sprint_id` JSON,`issue_project_name` VARCHAR(255) NOT NULL,`users` JSON,`initial_start_date` BIGINT UNSIGNED NOT NULL,`total_duration` BIGINT UNSIGNED NOT NULL,`created_at`BIGINT UNSIGNED,`closed_at` BIGINT UNSIGNED,`planned_end_date`BIGINT UNSIGNED,`customer_id` VARCHAR(64) NOT NULL,`ref_type` VARCHAR(20) NOT NULL,`ref_id` VARCHAR(64) NOT NULL,`custom_field_ids_virtual` TEXT,`release_duration`BIGINT NOT NULL,`completed` BOOL NOT NULL,`completed_date` BIGINT UNSIGNED,INDEX issue_summary_issue_id_index (`issue_id`),INDEX issue_summary_priority_id_index (`priority_id`),INDEX issue_summary_issue_type_id_index (`issue_type_id`),INDEX issue_summary_resolution_id_index (`resolution_id`),INDEX issue_summary_parent_issue_id_index (`parent_issue_id`),INDEX issue_summary_project_id_index (`project_id`),INDEX issue_summary_top_level_index (`top_level`),INDEX issue_summary_strategic_parent_id_index (`strategic_parent_id`),INDEX issue_summary_customer_id_index (`customer_id`),INDEX issue_summary_ref_id_index (`ref_id`),INDEX issue_summary_customer_id_parent_issue_id_index (`customer_id`,`parent_issue_id`),INDEX issue_summary_customer_id_top_level_index (`customer_id`,`top_level`),INDEX issue_summary_customer_id_top_level_issue_type_id_index (`customer_id`,`top_level`,`issue_type_id`),INDEX issue_summary_customer_id_top_level_issue_type_id_priority_id_in (`customer_id`,`top_level`,`issue_type_id`,`priority_id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBCreateIssueSummaryTableTx will create the IssueSummary table using the provided transction
func DBCreateIssueSummaryTableTx(ctx context.Context, tx Tx) error {
	q := "CREATE TABLE `issue_summary` (`id` VARCHAR(64) NOT NULL PRIMARY KEY,`checksum` CHAR(64),`issue_id` VARCHAR(64) NOT NULL,`total_issues` INT UNSIGNED NOT NULL,`new30_days`INT UNSIGNED NOT NULL,`total_closed` INT UNSIGNED NOT NULL,`closed30_days`INT UNSIGNED NOT NULL,`estimated_work_months` FLOAT NOT NULL,`estimated_work_months30_days` FLOAT NOT NULL,`title` TEXT NOT NULL,`url` TEXT,`priority` VARCHAR(100),`priority_id` VARCHAR(64),`status` VARCHAR(100),`status_id` VARCHAR(64),`issue_type`VARCHAR(100) NOT NULL,`issue_type_id`VARCHAR(64),`resolution`VARCHAR(100),`resolution_id`VARCHAR(64),`state` VARCHAR(10) NOT NULL,`custom_field_ids`JSON,`teams` JSON,`parent_issue_id` VARCHAR(64),`parents_issue_ids` JSON,`metadata` JSON,`project_id`VARCHAR(64) NOT NULL,`sprints`JSON,`labels` JSON,`top_level` TINYINT UNSIGNED NOT NULL,`is_leaf`TINYINT UNSIGNED NOT NULL,`path`VARCHAR(1024) NOT NULL,`in_progress_duration` BIGINT NOT NULL,`verification_duration` BIGINT NOT NULL,`in_progress_count` INT NOT NULL,`reopen_count` INT NOT NULL,`mapped_type` VARCHAR(75) NOT NULL,`strategic_parent_id`VARCHAR(64),`sprint_id` JSON,`issue_project_name` VARCHAR(255) NOT NULL,`users` JSON,`initial_start_date` BIGINT UNSIGNED NOT NULL,`total_duration` BIGINT UNSIGNED NOT NULL,`created_at`BIGINT UNSIGNED,`closed_at` BIGINT UNSIGNED,`planned_end_date`BIGINT UNSIGNED,`customer_id` VARCHAR(64) NOT NULL,`ref_type` VARCHAR(20) NOT NULL,`ref_id` VARCHAR(64) NOT NULL,`custom_field_ids_virtual` TEXT,`release_duration`BIGINT NOT NULL,`completed` BOOL NOT NULL,`completed_date` BIGINT UNSIGNED,INDEX issue_summary_issue_id_index (`issue_id`),INDEX issue_summary_priority_id_index (`priority_id`),INDEX issue_summary_issue_type_id_index (`issue_type_id`),INDEX issue_summary_resolution_id_index (`resolution_id`),INDEX issue_summary_parent_issue_id_index (`parent_issue_id`),INDEX issue_summary_project_id_index (`project_id`),INDEX issue_summary_top_level_index (`top_level`),INDEX issue_summary_strategic_parent_id_index (`strategic_parent_id`),INDEX issue_summary_customer_id_index (`customer_id`),INDEX issue_summary_ref_id_index (`ref_id`),INDEX issue_summary_customer_id_parent_issue_id_index (`customer_id`,`parent_issue_id`),INDEX issue_summary_customer_id_top_level_index (`customer_id`,`top_level`),INDEX issue_summary_customer_id_top_level_issue_type_id_index (`customer_id`,`top_level`,`issue_type_id`),INDEX issue_summary_customer_id_top_level_issue_type_id_priority_id_in (`customer_id`,`top_level`,`issue_type_id`,`priority_id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// DBDropIssueSummaryTable will drop the IssueSummary table
func DBDropIssueSummaryTable(ctx context.Context, db DB) error {
	q := "DROP TABLE IF EXISTS `issue_summary`"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBDropIssueSummaryTableTx will drop the IssueSummary table using the provided transaction
func DBDropIssueSummaryTableTx(ctx context.Context, tx Tx) error {
	q := "DROP TABLE IF EXISTS `issue_summary`"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// CalculateChecksum will calculate a checksum of the SHA1 of all field values
func (t *IssueSummary) CalculateChecksum() string {
	return orm.HashStrings(
		orm.ToString(t.ID),
		orm.ToString(t.IssueID),
		orm.ToString(t.TotalIssues),
		orm.ToString(t.New30Days),
		orm.ToString(t.TotalClosed),
		orm.ToString(t.Closed30Days),
		orm.ToString(t.EstimatedWorkMonths),
		orm.ToString(t.EstimatedWorkMonths30Days),
		orm.ToString(t.Title),
		orm.ToString(t.URL),
		orm.ToString(t.Priority),
		orm.ToString(t.PriorityID),
		orm.ToString(t.Status),
		orm.ToString(t.StatusID),
		orm.ToString(t.IssueType),
		orm.ToString(t.IssueTypeID),
		orm.ToString(t.Resolution),
		orm.ToString(t.ResolutionID),
		orm.ToString(t.State),
		orm.ToString(t.CustomFieldIds),
		orm.ToString(t.Teams),
		orm.ToString(t.ParentIssueID),
		orm.ToString(t.ParentsIssueIds),
		orm.ToString(t.Metadata),
		orm.ToString(t.ProjectID),
		orm.ToString(t.Sprints),
		orm.ToString(t.Labels),
		orm.ToString(t.TopLevel),
		orm.ToString(t.IsLeaf),
		orm.ToString(t.Path),
		orm.ToString(t.InProgressDuration),
		orm.ToString(t.VerificationDuration),
		orm.ToString(t.InProgressCount),
		orm.ToString(t.ReopenCount),
		orm.ToString(t.MappedType),
		orm.ToString(t.StrategicParentID),
		orm.ToString(t.SprintID),
		orm.ToString(t.IssueProjectName),
		orm.ToString(t.Users),
		orm.ToString(t.InitialStartDate),
		orm.ToString(t.TotalDuration),
		orm.ToString(t.CreatedAt),
		orm.ToString(t.ClosedAt),
		orm.ToString(t.PlannedEndDate),
		orm.ToString(t.CustomerID),
		orm.ToString(t.RefType),
		orm.ToString(t.RefID),
		orm.ToString(t.CustomFieldIdsVirtual),
		orm.ToString(t.ReleaseDuration),
		orm.ToString(t.Completed),
		orm.ToString(t.CompletedDate),
	)
}

// DBCreate will create a new IssueSummary record in the database
func (t *IssueSummary) DBCreate(ctx context.Context, db DB) (sql.Result, error) {
	q := "INSERT INTO `issue_summary` (`issue_summary`.`id`,`issue_summary`.`checksum`,`issue_summary`.`issue_id`,`issue_summary`.`total_issues`,`issue_summary`.`new30_days`,`issue_summary`.`total_closed`,`issue_summary`.`closed30_days`,`issue_summary`.`estimated_work_months`,`issue_summary`.`estimated_work_months30_days`,`issue_summary`.`title`,`issue_summary`.`url`,`issue_summary`.`priority`,`issue_summary`.`priority_id`,`issue_summary`.`status`,`issue_summary`.`status_id`,`issue_summary`.`issue_type`,`issue_summary`.`issue_type_id`,`issue_summary`.`resolution`,`issue_summary`.`resolution_id`,`issue_summary`.`state`,`issue_summary`.`custom_field_ids`,`issue_summary`.`teams`,`issue_summary`.`parent_issue_id`,`issue_summary`.`parents_issue_ids`,`issue_summary`.`metadata`,`issue_summary`.`project_id`,`issue_summary`.`sprints`,`issue_summary`.`labels`,`issue_summary`.`top_level`,`issue_summary`.`is_leaf`,`issue_summary`.`path`,`issue_summary`.`in_progress_duration`,`issue_summary`.`verification_duration`,`issue_summary`.`in_progress_count`,`issue_summary`.`reopen_count`,`issue_summary`.`mapped_type`,`issue_summary`.`strategic_parent_id`,`issue_summary`.`sprint_id`,`issue_summary`.`issue_project_name`,`issue_summary`.`users`,`issue_summary`.`initial_start_date`,`issue_summary`.`total_duration`,`issue_summary`.`created_at`,`issue_summary`.`closed_at`,`issue_summary`.`planned_end_date`,`issue_summary`.`customer_id`,`issue_summary`.`ref_type`,`issue_summary`.`ref_id`,`issue_summary`.`custom_field_ids_virtual`,`issue_summary`.`release_duration`,`issue_summary`.`completed`,`issue_summary`.`completed_date`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.IssueID),
		orm.ToSQLInt64(t.TotalIssues),
		orm.ToSQLInt64(t.New30Days),
		orm.ToSQLInt64(t.TotalClosed),
		orm.ToSQLInt64(t.Closed30Days),
		orm.ToSQLFloat64(t.EstimatedWorkMonths),
		orm.ToSQLFloat64(t.EstimatedWorkMonths30Days),
		orm.ToSQLString(t.Title),
		orm.ToSQLString(t.URL),
		orm.ToSQLString(t.Priority),
		orm.ToSQLString(t.PriorityID),
		orm.ToSQLString(t.Status),
		orm.ToSQLString(t.StatusID),
		orm.ToSQLString(t.IssueType),
		orm.ToSQLString(t.IssueTypeID),
		orm.ToSQLString(t.Resolution),
		orm.ToSQLString(t.ResolutionID),
		orm.ToSQLString(t.State),
		orm.ToSQLString(t.CustomFieldIds),
		orm.ToSQLString(t.Teams),
		orm.ToSQLString(t.ParentIssueID),
		orm.ToSQLString(t.ParentsIssueIds),
		orm.ToSQLString(t.Metadata),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLString(t.Sprints),
		orm.ToSQLString(t.Labels),
		orm.ToSQLBool(t.TopLevel),
		orm.ToSQLBool(t.IsLeaf),
		orm.ToSQLString(t.Path),
		orm.ToSQLInt64(t.InProgressDuration),
		orm.ToSQLInt64(t.VerificationDuration),
		orm.ToSQLInt64(t.InProgressCount),
		orm.ToSQLInt64(t.ReopenCount),
		orm.ToSQLString(t.MappedType),
		orm.ToSQLString(t.StrategicParentID),
		orm.ToSQLString(t.SprintID),
		orm.ToSQLString(t.IssueProjectName),
		orm.ToSQLString(t.Users),
		orm.ToSQLInt64(t.InitialStartDate),
		orm.ToSQLInt64(t.TotalDuration),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.ClosedAt),
		orm.ToSQLInt64(t.PlannedEndDate),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefType),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.CustomFieldIdsVirtual),
		orm.ToSQLInt64(t.ReleaseDuration),
		orm.ToSQLBool(t.Completed),
		orm.ToSQLInt64(t.CompletedDate),
	)
}

// DBCreateTx will create a new IssueSummary record in the database using the provided transaction
func (t *IssueSummary) DBCreateTx(ctx context.Context, tx Tx) (sql.Result, error) {
	q := "INSERT INTO `issue_summary` (`issue_summary`.`id`,`issue_summary`.`checksum`,`issue_summary`.`issue_id`,`issue_summary`.`total_issues`,`issue_summary`.`new30_days`,`issue_summary`.`total_closed`,`issue_summary`.`closed30_days`,`issue_summary`.`estimated_work_months`,`issue_summary`.`estimated_work_months30_days`,`issue_summary`.`title`,`issue_summary`.`url`,`issue_summary`.`priority`,`issue_summary`.`priority_id`,`issue_summary`.`status`,`issue_summary`.`status_id`,`issue_summary`.`issue_type`,`issue_summary`.`issue_type_id`,`issue_summary`.`resolution`,`issue_summary`.`resolution_id`,`issue_summary`.`state`,`issue_summary`.`custom_field_ids`,`issue_summary`.`teams`,`issue_summary`.`parent_issue_id`,`issue_summary`.`parents_issue_ids`,`issue_summary`.`metadata`,`issue_summary`.`project_id`,`issue_summary`.`sprints`,`issue_summary`.`labels`,`issue_summary`.`top_level`,`issue_summary`.`is_leaf`,`issue_summary`.`path`,`issue_summary`.`in_progress_duration`,`issue_summary`.`verification_duration`,`issue_summary`.`in_progress_count`,`issue_summary`.`reopen_count`,`issue_summary`.`mapped_type`,`issue_summary`.`strategic_parent_id`,`issue_summary`.`sprint_id`,`issue_summary`.`issue_project_name`,`issue_summary`.`users`,`issue_summary`.`initial_start_date`,`issue_summary`.`total_duration`,`issue_summary`.`created_at`,`issue_summary`.`closed_at`,`issue_summary`.`planned_end_date`,`issue_summary`.`customer_id`,`issue_summary`.`ref_type`,`issue_summary`.`ref_id`,`issue_summary`.`custom_field_ids_virtual`,`issue_summary`.`release_duration`,`issue_summary`.`completed`,`issue_summary`.`completed_date`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.IssueID),
		orm.ToSQLInt64(t.TotalIssues),
		orm.ToSQLInt64(t.New30Days),
		orm.ToSQLInt64(t.TotalClosed),
		orm.ToSQLInt64(t.Closed30Days),
		orm.ToSQLFloat64(t.EstimatedWorkMonths),
		orm.ToSQLFloat64(t.EstimatedWorkMonths30Days),
		orm.ToSQLString(t.Title),
		orm.ToSQLString(t.URL),
		orm.ToSQLString(t.Priority),
		orm.ToSQLString(t.PriorityID),
		orm.ToSQLString(t.Status),
		orm.ToSQLString(t.StatusID),
		orm.ToSQLString(t.IssueType),
		orm.ToSQLString(t.IssueTypeID),
		orm.ToSQLString(t.Resolution),
		orm.ToSQLString(t.ResolutionID),
		orm.ToSQLString(t.State),
		orm.ToSQLString(t.CustomFieldIds),
		orm.ToSQLString(t.Teams),
		orm.ToSQLString(t.ParentIssueID),
		orm.ToSQLString(t.ParentsIssueIds),
		orm.ToSQLString(t.Metadata),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLString(t.Sprints),
		orm.ToSQLString(t.Labels),
		orm.ToSQLBool(t.TopLevel),
		orm.ToSQLBool(t.IsLeaf),
		orm.ToSQLString(t.Path),
		orm.ToSQLInt64(t.InProgressDuration),
		orm.ToSQLInt64(t.VerificationDuration),
		orm.ToSQLInt64(t.InProgressCount),
		orm.ToSQLInt64(t.ReopenCount),
		orm.ToSQLString(t.MappedType),
		orm.ToSQLString(t.StrategicParentID),
		orm.ToSQLString(t.SprintID),
		orm.ToSQLString(t.IssueProjectName),
		orm.ToSQLString(t.Users),
		orm.ToSQLInt64(t.InitialStartDate),
		orm.ToSQLInt64(t.TotalDuration),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.ClosedAt),
		orm.ToSQLInt64(t.PlannedEndDate),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefType),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.CustomFieldIdsVirtual),
		orm.ToSQLInt64(t.ReleaseDuration),
		orm.ToSQLBool(t.Completed),
		orm.ToSQLInt64(t.CompletedDate),
	)
}

// DBCreateIgnoreDuplicate will upsert the IssueSummary record in the database
func (t *IssueSummary) DBCreateIgnoreDuplicate(ctx context.Context, db DB) (sql.Result, error) {
	q := "INSERT INTO `issue_summary` (`issue_summary`.`id`,`issue_summary`.`checksum`,`issue_summary`.`issue_id`,`issue_summary`.`total_issues`,`issue_summary`.`new30_days`,`issue_summary`.`total_closed`,`issue_summary`.`closed30_days`,`issue_summary`.`estimated_work_months`,`issue_summary`.`estimated_work_months30_days`,`issue_summary`.`title`,`issue_summary`.`url`,`issue_summary`.`priority`,`issue_summary`.`priority_id`,`issue_summary`.`status`,`issue_summary`.`status_id`,`issue_summary`.`issue_type`,`issue_summary`.`issue_type_id`,`issue_summary`.`resolution`,`issue_summary`.`resolution_id`,`issue_summary`.`state`,`issue_summary`.`custom_field_ids`,`issue_summary`.`teams`,`issue_summary`.`parent_issue_id`,`issue_summary`.`parents_issue_ids`,`issue_summary`.`metadata`,`issue_summary`.`project_id`,`issue_summary`.`sprints`,`issue_summary`.`labels`,`issue_summary`.`top_level`,`issue_summary`.`is_leaf`,`issue_summary`.`path`,`issue_summary`.`in_progress_duration`,`issue_summary`.`verification_duration`,`issue_summary`.`in_progress_count`,`issue_summary`.`reopen_count`,`issue_summary`.`mapped_type`,`issue_summary`.`strategic_parent_id`,`issue_summary`.`sprint_id`,`issue_summary`.`issue_project_name`,`issue_summary`.`users`,`issue_summary`.`initial_start_date`,`issue_summary`.`total_duration`,`issue_summary`.`created_at`,`issue_summary`.`closed_at`,`issue_summary`.`planned_end_date`,`issue_summary`.`customer_id`,`issue_summary`.`ref_type`,`issue_summary`.`ref_id`,`issue_summary`.`custom_field_ids_virtual`,`issue_summary`.`release_duration`,`issue_summary`.`completed`,`issue_summary`.`completed_date`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.IssueID),
		orm.ToSQLInt64(t.TotalIssues),
		orm.ToSQLInt64(t.New30Days),
		orm.ToSQLInt64(t.TotalClosed),
		orm.ToSQLInt64(t.Closed30Days),
		orm.ToSQLFloat64(t.EstimatedWorkMonths),
		orm.ToSQLFloat64(t.EstimatedWorkMonths30Days),
		orm.ToSQLString(t.Title),
		orm.ToSQLString(t.URL),
		orm.ToSQLString(t.Priority),
		orm.ToSQLString(t.PriorityID),
		orm.ToSQLString(t.Status),
		orm.ToSQLString(t.StatusID),
		orm.ToSQLString(t.IssueType),
		orm.ToSQLString(t.IssueTypeID),
		orm.ToSQLString(t.Resolution),
		orm.ToSQLString(t.ResolutionID),
		orm.ToSQLString(t.State),
		orm.ToSQLString(t.CustomFieldIds),
		orm.ToSQLString(t.Teams),
		orm.ToSQLString(t.ParentIssueID),
		orm.ToSQLString(t.ParentsIssueIds),
		orm.ToSQLString(t.Metadata),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLString(t.Sprints),
		orm.ToSQLString(t.Labels),
		orm.ToSQLBool(t.TopLevel),
		orm.ToSQLBool(t.IsLeaf),
		orm.ToSQLString(t.Path),
		orm.ToSQLInt64(t.InProgressDuration),
		orm.ToSQLInt64(t.VerificationDuration),
		orm.ToSQLInt64(t.InProgressCount),
		orm.ToSQLInt64(t.ReopenCount),
		orm.ToSQLString(t.MappedType),
		orm.ToSQLString(t.StrategicParentID),
		orm.ToSQLString(t.SprintID),
		orm.ToSQLString(t.IssueProjectName),
		orm.ToSQLString(t.Users),
		orm.ToSQLInt64(t.InitialStartDate),
		orm.ToSQLInt64(t.TotalDuration),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.ClosedAt),
		orm.ToSQLInt64(t.PlannedEndDate),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefType),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.CustomFieldIdsVirtual),
		orm.ToSQLInt64(t.ReleaseDuration),
		orm.ToSQLBool(t.Completed),
		orm.ToSQLInt64(t.CompletedDate),
	)
}

// DBCreateIgnoreDuplicateTx will upsert the IssueSummary record in the database using the provided transaction
func (t *IssueSummary) DBCreateIgnoreDuplicateTx(ctx context.Context, tx Tx) (sql.Result, error) {
	q := "INSERT INTO `issue_summary` (`issue_summary`.`id`,`issue_summary`.`checksum`,`issue_summary`.`issue_id`,`issue_summary`.`total_issues`,`issue_summary`.`new30_days`,`issue_summary`.`total_closed`,`issue_summary`.`closed30_days`,`issue_summary`.`estimated_work_months`,`issue_summary`.`estimated_work_months30_days`,`issue_summary`.`title`,`issue_summary`.`url`,`issue_summary`.`priority`,`issue_summary`.`priority_id`,`issue_summary`.`status`,`issue_summary`.`status_id`,`issue_summary`.`issue_type`,`issue_summary`.`issue_type_id`,`issue_summary`.`resolution`,`issue_summary`.`resolution_id`,`issue_summary`.`state`,`issue_summary`.`custom_field_ids`,`issue_summary`.`teams`,`issue_summary`.`parent_issue_id`,`issue_summary`.`parents_issue_ids`,`issue_summary`.`metadata`,`issue_summary`.`project_id`,`issue_summary`.`sprints`,`issue_summary`.`labels`,`issue_summary`.`top_level`,`issue_summary`.`is_leaf`,`issue_summary`.`path`,`issue_summary`.`in_progress_duration`,`issue_summary`.`verification_duration`,`issue_summary`.`in_progress_count`,`issue_summary`.`reopen_count`,`issue_summary`.`mapped_type`,`issue_summary`.`strategic_parent_id`,`issue_summary`.`sprint_id`,`issue_summary`.`issue_project_name`,`issue_summary`.`users`,`issue_summary`.`initial_start_date`,`issue_summary`.`total_duration`,`issue_summary`.`created_at`,`issue_summary`.`closed_at`,`issue_summary`.`planned_end_date`,`issue_summary`.`customer_id`,`issue_summary`.`ref_type`,`issue_summary`.`ref_id`,`issue_summary`.`custom_field_ids_virtual`,`issue_summary`.`release_duration`,`issue_summary`.`completed`,`issue_summary`.`completed_date`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.IssueID),
		orm.ToSQLInt64(t.TotalIssues),
		orm.ToSQLInt64(t.New30Days),
		orm.ToSQLInt64(t.TotalClosed),
		orm.ToSQLInt64(t.Closed30Days),
		orm.ToSQLFloat64(t.EstimatedWorkMonths),
		orm.ToSQLFloat64(t.EstimatedWorkMonths30Days),
		orm.ToSQLString(t.Title),
		orm.ToSQLString(t.URL),
		orm.ToSQLString(t.Priority),
		orm.ToSQLString(t.PriorityID),
		orm.ToSQLString(t.Status),
		orm.ToSQLString(t.StatusID),
		orm.ToSQLString(t.IssueType),
		orm.ToSQLString(t.IssueTypeID),
		orm.ToSQLString(t.Resolution),
		orm.ToSQLString(t.ResolutionID),
		orm.ToSQLString(t.State),
		orm.ToSQLString(t.CustomFieldIds),
		orm.ToSQLString(t.Teams),
		orm.ToSQLString(t.ParentIssueID),
		orm.ToSQLString(t.ParentsIssueIds),
		orm.ToSQLString(t.Metadata),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLString(t.Sprints),
		orm.ToSQLString(t.Labels),
		orm.ToSQLBool(t.TopLevel),
		orm.ToSQLBool(t.IsLeaf),
		orm.ToSQLString(t.Path),
		orm.ToSQLInt64(t.InProgressDuration),
		orm.ToSQLInt64(t.VerificationDuration),
		orm.ToSQLInt64(t.InProgressCount),
		orm.ToSQLInt64(t.ReopenCount),
		orm.ToSQLString(t.MappedType),
		orm.ToSQLString(t.StrategicParentID),
		orm.ToSQLString(t.SprintID),
		orm.ToSQLString(t.IssueProjectName),
		orm.ToSQLString(t.Users),
		orm.ToSQLInt64(t.InitialStartDate),
		orm.ToSQLInt64(t.TotalDuration),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.ClosedAt),
		orm.ToSQLInt64(t.PlannedEndDate),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefType),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.CustomFieldIdsVirtual),
		orm.ToSQLInt64(t.ReleaseDuration),
		orm.ToSQLBool(t.Completed),
		orm.ToSQLInt64(t.CompletedDate),
	)
}

// DeleteAllIssueSummaries deletes all IssueSummary records in the database with optional filters
func DeleteAllIssueSummaries(ctx context.Context, db DB, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(IssueSummaryTableName),
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

// DeleteAllIssueSummariesTx deletes all IssueSummary records in the database with optional filters using the provided transaction
func DeleteAllIssueSummariesTx(ctx context.Context, tx Tx, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(IssueSummaryTableName),
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

// DBDelete will delete this IssueSummary record in the database
func (t *IssueSummary) DBDelete(ctx context.Context, db DB) (bool, error) {
	q := "DELETE FROM `issue_summary` WHERE `id` = ?"
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

// DBDeleteTx will delete this IssueSummary record in the database using the provided transaction
func (t *IssueSummary) DBDeleteTx(ctx context.Context, tx Tx) (bool, error) {
	q := "DELETE FROM `issue_summary` WHERE `id` = ?"
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

// DBUpdate will update the IssueSummary record in the database
func (t *IssueSummary) DBUpdate(ctx context.Context, db DB) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `issue_summary` SET `checksum`=?,`issue_id`=?,`total_issues`=?,`new30_days`=?,`total_closed`=?,`closed30_days`=?,`estimated_work_months`=?,`estimated_work_months30_days`=?,`title`=?,`url`=?,`priority`=?,`priority_id`=?,`status`=?,`status_id`=?,`issue_type`=?,`issue_type_id`=?,`resolution`=?,`resolution_id`=?,`state`=?,`custom_field_ids`=?,`teams`=?,`parent_issue_id`=?,`parents_issue_ids`=?,`metadata`=?,`project_id`=?,`sprints`=?,`labels`=?,`top_level`=?,`is_leaf`=?,`path`=?,`in_progress_duration`=?,`verification_duration`=?,`in_progress_count`=?,`reopen_count`=?,`mapped_type`=?,`strategic_parent_id`=?,`sprint_id`=?,`issue_project_name`=?,`users`=?,`initial_start_date`=?,`total_duration`=?,`created_at`=?,`closed_at`=?,`planned_end_date`=?,`customer_id`=?,`ref_type`=?,`ref_id`=?,`custom_field_ids_virtual`=?,`release_duration`=?,`completed`=?,`completed_date`=? WHERE `id`=?"
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.IssueID),
		orm.ToSQLInt64(t.TotalIssues),
		orm.ToSQLInt64(t.New30Days),
		orm.ToSQLInt64(t.TotalClosed),
		orm.ToSQLInt64(t.Closed30Days),
		orm.ToSQLFloat64(t.EstimatedWorkMonths),
		orm.ToSQLFloat64(t.EstimatedWorkMonths30Days),
		orm.ToSQLString(t.Title),
		orm.ToSQLString(t.URL),
		orm.ToSQLString(t.Priority),
		orm.ToSQLString(t.PriorityID),
		orm.ToSQLString(t.Status),
		orm.ToSQLString(t.StatusID),
		orm.ToSQLString(t.IssueType),
		orm.ToSQLString(t.IssueTypeID),
		orm.ToSQLString(t.Resolution),
		orm.ToSQLString(t.ResolutionID),
		orm.ToSQLString(t.State),
		orm.ToSQLString(t.CustomFieldIds),
		orm.ToSQLString(t.Teams),
		orm.ToSQLString(t.ParentIssueID),
		orm.ToSQLString(t.ParentsIssueIds),
		orm.ToSQLString(t.Metadata),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLString(t.Sprints),
		orm.ToSQLString(t.Labels),
		orm.ToSQLBool(t.TopLevel),
		orm.ToSQLBool(t.IsLeaf),
		orm.ToSQLString(t.Path),
		orm.ToSQLInt64(t.InProgressDuration),
		orm.ToSQLInt64(t.VerificationDuration),
		orm.ToSQLInt64(t.InProgressCount),
		orm.ToSQLInt64(t.ReopenCount),
		orm.ToSQLString(t.MappedType),
		orm.ToSQLString(t.StrategicParentID),
		orm.ToSQLString(t.SprintID),
		orm.ToSQLString(t.IssueProjectName),
		orm.ToSQLString(t.Users),
		orm.ToSQLInt64(t.InitialStartDate),
		orm.ToSQLInt64(t.TotalDuration),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.ClosedAt),
		orm.ToSQLInt64(t.PlannedEndDate),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefType),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.CustomFieldIdsVirtual),
		orm.ToSQLInt64(t.ReleaseDuration),
		orm.ToSQLBool(t.Completed),
		orm.ToSQLInt64(t.CompletedDate),
		orm.ToSQLString(t.ID),
	)
}

// DBUpdateTx will update the IssueSummary record in the database using the provided transaction
func (t *IssueSummary) DBUpdateTx(ctx context.Context, tx Tx) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `issue_summary` SET `checksum`=?,`issue_id`=?,`total_issues`=?,`new30_days`=?,`total_closed`=?,`closed30_days`=?,`estimated_work_months`=?,`estimated_work_months30_days`=?,`title`=?,`url`=?,`priority`=?,`priority_id`=?,`status`=?,`status_id`=?,`issue_type`=?,`issue_type_id`=?,`resolution`=?,`resolution_id`=?,`state`=?,`custom_field_ids`=?,`teams`=?,`parent_issue_id`=?,`parents_issue_ids`=?,`metadata`=?,`project_id`=?,`sprints`=?,`labels`=?,`top_level`=?,`is_leaf`=?,`path`=?,`in_progress_duration`=?,`verification_duration`=?,`in_progress_count`=?,`reopen_count`=?,`mapped_type`=?,`strategic_parent_id`=?,`sprint_id`=?,`issue_project_name`=?,`users`=?,`initial_start_date`=?,`total_duration`=?,`created_at`=?,`closed_at`=?,`planned_end_date`=?,`customer_id`=?,`ref_type`=?,`ref_id`=?,`custom_field_ids_virtual`=?,`release_duration`=?,`completed`=?,`completed_date`=? WHERE `id`=?"
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.IssueID),
		orm.ToSQLInt64(t.TotalIssues),
		orm.ToSQLInt64(t.New30Days),
		orm.ToSQLInt64(t.TotalClosed),
		orm.ToSQLInt64(t.Closed30Days),
		orm.ToSQLFloat64(t.EstimatedWorkMonths),
		orm.ToSQLFloat64(t.EstimatedWorkMonths30Days),
		orm.ToSQLString(t.Title),
		orm.ToSQLString(t.URL),
		orm.ToSQLString(t.Priority),
		orm.ToSQLString(t.PriorityID),
		orm.ToSQLString(t.Status),
		orm.ToSQLString(t.StatusID),
		orm.ToSQLString(t.IssueType),
		orm.ToSQLString(t.IssueTypeID),
		orm.ToSQLString(t.Resolution),
		orm.ToSQLString(t.ResolutionID),
		orm.ToSQLString(t.State),
		orm.ToSQLString(t.CustomFieldIds),
		orm.ToSQLString(t.Teams),
		orm.ToSQLString(t.ParentIssueID),
		orm.ToSQLString(t.ParentsIssueIds),
		orm.ToSQLString(t.Metadata),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLString(t.Sprints),
		orm.ToSQLString(t.Labels),
		orm.ToSQLBool(t.TopLevel),
		orm.ToSQLBool(t.IsLeaf),
		orm.ToSQLString(t.Path),
		orm.ToSQLInt64(t.InProgressDuration),
		orm.ToSQLInt64(t.VerificationDuration),
		orm.ToSQLInt64(t.InProgressCount),
		orm.ToSQLInt64(t.ReopenCount),
		orm.ToSQLString(t.MappedType),
		orm.ToSQLString(t.StrategicParentID),
		orm.ToSQLString(t.SprintID),
		orm.ToSQLString(t.IssueProjectName),
		orm.ToSQLString(t.Users),
		orm.ToSQLInt64(t.InitialStartDate),
		orm.ToSQLInt64(t.TotalDuration),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.ClosedAt),
		orm.ToSQLInt64(t.PlannedEndDate),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefType),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.CustomFieldIdsVirtual),
		orm.ToSQLInt64(t.ReleaseDuration),
		orm.ToSQLBool(t.Completed),
		orm.ToSQLInt64(t.CompletedDate),
		orm.ToSQLString(t.ID),
	)
}

// DBUpsert will upsert the IssueSummary record in the database
func (t *IssueSummary) DBUpsert(ctx context.Context, db DB, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `issue_summary` (`issue_summary`.`id`,`issue_summary`.`checksum`,`issue_summary`.`issue_id`,`issue_summary`.`total_issues`,`issue_summary`.`new30_days`,`issue_summary`.`total_closed`,`issue_summary`.`closed30_days`,`issue_summary`.`estimated_work_months`,`issue_summary`.`estimated_work_months30_days`,`issue_summary`.`title`,`issue_summary`.`url`,`issue_summary`.`priority`,`issue_summary`.`priority_id`,`issue_summary`.`status`,`issue_summary`.`status_id`,`issue_summary`.`issue_type`,`issue_summary`.`issue_type_id`,`issue_summary`.`resolution`,`issue_summary`.`resolution_id`,`issue_summary`.`state`,`issue_summary`.`custom_field_ids`,`issue_summary`.`teams`,`issue_summary`.`parent_issue_id`,`issue_summary`.`parents_issue_ids`,`issue_summary`.`metadata`,`issue_summary`.`project_id`,`issue_summary`.`sprints`,`issue_summary`.`labels`,`issue_summary`.`top_level`,`issue_summary`.`is_leaf`,`issue_summary`.`path`,`issue_summary`.`in_progress_duration`,`issue_summary`.`verification_duration`,`issue_summary`.`in_progress_count`,`issue_summary`.`reopen_count`,`issue_summary`.`mapped_type`,`issue_summary`.`strategic_parent_id`,`issue_summary`.`sprint_id`,`issue_summary`.`issue_project_name`,`issue_summary`.`users`,`issue_summary`.`initial_start_date`,`issue_summary`.`total_duration`,`issue_summary`.`created_at`,`issue_summary`.`closed_at`,`issue_summary`.`planned_end_date`,`issue_summary`.`customer_id`,`issue_summary`.`ref_type`,`issue_summary`.`ref_id`,`issue_summary`.`custom_field_ids_virtual`,`issue_summary`.`release_duration`,`issue_summary`.`completed`,`issue_summary`.`completed_date`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `issue_summary` (`issue_summary`.`id`,`issue_summary`.`checksum`,`issue_summary`.`issue_id`,`issue_summary`.`total_issues`,`issue_summary`.`new30_days`,`issue_summary`.`total_closed`,`issue_summary`.`closed30_days`,`issue_summary`.`estimated_work_months`,`issue_summary`.`estimated_work_months30_days`,`issue_summary`.`title`,`issue_summary`.`url`,`issue_summary`.`priority`,`issue_summary`.`priority_id`,`issue_summary`.`status`,`issue_summary`.`status_id`,`issue_summary`.`issue_type`,`issue_summary`.`issue_type_id`,`issue_summary`.`resolution`,`issue_summary`.`resolution_id`,`issue_summary`.`state`,`issue_summary`.`custom_field_ids`,`issue_summary`.`teams`,`issue_summary`.`parent_issue_id`,`issue_summary`.`parents_issue_ids`,`issue_summary`.`metadata`,`issue_summary`.`project_id`,`issue_summary`.`sprints`,`issue_summary`.`labels`,`issue_summary`.`top_level`,`issue_summary`.`is_leaf`,`issue_summary`.`path`,`issue_summary`.`in_progress_duration`,`issue_summary`.`verification_duration`,`issue_summary`.`in_progress_count`,`issue_summary`.`reopen_count`,`issue_summary`.`mapped_type`,`issue_summary`.`strategic_parent_id`,`issue_summary`.`sprint_id`,`issue_summary`.`issue_project_name`,`issue_summary`.`users`,`issue_summary`.`initial_start_date`,`issue_summary`.`total_duration`,`issue_summary`.`created_at`,`issue_summary`.`closed_at`,`issue_summary`.`planned_end_date`,`issue_summary`.`customer_id`,`issue_summary`.`ref_type`,`issue_summary`.`ref_id`,`issue_summary`.`custom_field_ids_virtual`,`issue_summary`.`release_duration`,`issue_summary`.`completed`,`issue_summary`.`completed_date`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`issue_id`=VALUES(`issue_id`),`total_issues`=VALUES(`total_issues`),`new30_days`=VALUES(`new30_days`),`total_closed`=VALUES(`total_closed`),`closed30_days`=VALUES(`closed30_days`),`estimated_work_months`=VALUES(`estimated_work_months`),`estimated_work_months30_days`=VALUES(`estimated_work_months30_days`),`title`=VALUES(`title`),`url`=VALUES(`url`),`priority`=VALUES(`priority`),`priority_id`=VALUES(`priority_id`),`status`=VALUES(`status`),`status_id`=VALUES(`status_id`),`issue_type`=VALUES(`issue_type`),`issue_type_id`=VALUES(`issue_type_id`),`resolution`=VALUES(`resolution`),`resolution_id`=VALUES(`resolution_id`),`state`=VALUES(`state`),`custom_field_ids`=VALUES(`custom_field_ids`),`teams`=VALUES(`teams`),`parent_issue_id`=VALUES(`parent_issue_id`),`parents_issue_ids`=VALUES(`parents_issue_ids`),`metadata`=VALUES(`metadata`),`project_id`=VALUES(`project_id`),`sprints`=VALUES(`sprints`),`labels`=VALUES(`labels`),`top_level`=VALUES(`top_level`),`is_leaf`=VALUES(`is_leaf`),`path`=VALUES(`path`),`in_progress_duration`=VALUES(`in_progress_duration`),`verification_duration`=VALUES(`verification_duration`),`in_progress_count`=VALUES(`in_progress_count`),`reopen_count`=VALUES(`reopen_count`),`mapped_type`=VALUES(`mapped_type`),`strategic_parent_id`=VALUES(`strategic_parent_id`),`sprint_id`=VALUES(`sprint_id`),`issue_project_name`=VALUES(`issue_project_name`),`users`=VALUES(`users`),`initial_start_date`=VALUES(`initial_start_date`),`total_duration`=VALUES(`total_duration`),`created_at`=VALUES(`created_at`),`closed_at`=VALUES(`closed_at`),`planned_end_date`=VALUES(`planned_end_date`),`customer_id`=VALUES(`customer_id`),`ref_type`=VALUES(`ref_type`),`ref_id`=VALUES(`ref_id`),`custom_field_ids_virtual`=VALUES(`custom_field_ids_virtual`),`release_duration`=VALUES(`release_duration`),`completed`=VALUES(`completed`),`completed_date`=VALUES(`completed_date`)"
	}
	r, err := db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.IssueID),
		orm.ToSQLInt64(t.TotalIssues),
		orm.ToSQLInt64(t.New30Days),
		orm.ToSQLInt64(t.TotalClosed),
		orm.ToSQLInt64(t.Closed30Days),
		orm.ToSQLFloat64(t.EstimatedWorkMonths),
		orm.ToSQLFloat64(t.EstimatedWorkMonths30Days),
		orm.ToSQLString(t.Title),
		orm.ToSQLString(t.URL),
		orm.ToSQLString(t.Priority),
		orm.ToSQLString(t.PriorityID),
		orm.ToSQLString(t.Status),
		orm.ToSQLString(t.StatusID),
		orm.ToSQLString(t.IssueType),
		orm.ToSQLString(t.IssueTypeID),
		orm.ToSQLString(t.Resolution),
		orm.ToSQLString(t.ResolutionID),
		orm.ToSQLString(t.State),
		orm.ToSQLString(t.CustomFieldIds),
		orm.ToSQLString(t.Teams),
		orm.ToSQLString(t.ParentIssueID),
		orm.ToSQLString(t.ParentsIssueIds),
		orm.ToSQLString(t.Metadata),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLString(t.Sprints),
		orm.ToSQLString(t.Labels),
		orm.ToSQLBool(t.TopLevel),
		orm.ToSQLBool(t.IsLeaf),
		orm.ToSQLString(t.Path),
		orm.ToSQLInt64(t.InProgressDuration),
		orm.ToSQLInt64(t.VerificationDuration),
		orm.ToSQLInt64(t.InProgressCount),
		orm.ToSQLInt64(t.ReopenCount),
		orm.ToSQLString(t.MappedType),
		orm.ToSQLString(t.StrategicParentID),
		orm.ToSQLString(t.SprintID),
		orm.ToSQLString(t.IssueProjectName),
		orm.ToSQLString(t.Users),
		orm.ToSQLInt64(t.InitialStartDate),
		orm.ToSQLInt64(t.TotalDuration),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.ClosedAt),
		orm.ToSQLInt64(t.PlannedEndDate),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefType),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.CustomFieldIdsVirtual),
		orm.ToSQLInt64(t.ReleaseDuration),
		orm.ToSQLBool(t.Completed),
		orm.ToSQLInt64(t.CompletedDate),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBUpsertTx will upsert the IssueSummary record in the database using the provided transaction
func (t *IssueSummary) DBUpsertTx(ctx context.Context, tx Tx, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `issue_summary` (`issue_summary`.`id`,`issue_summary`.`checksum`,`issue_summary`.`issue_id`,`issue_summary`.`total_issues`,`issue_summary`.`new30_days`,`issue_summary`.`total_closed`,`issue_summary`.`closed30_days`,`issue_summary`.`estimated_work_months`,`issue_summary`.`estimated_work_months30_days`,`issue_summary`.`title`,`issue_summary`.`url`,`issue_summary`.`priority`,`issue_summary`.`priority_id`,`issue_summary`.`status`,`issue_summary`.`status_id`,`issue_summary`.`issue_type`,`issue_summary`.`issue_type_id`,`issue_summary`.`resolution`,`issue_summary`.`resolution_id`,`issue_summary`.`state`,`issue_summary`.`custom_field_ids`,`issue_summary`.`teams`,`issue_summary`.`parent_issue_id`,`issue_summary`.`parents_issue_ids`,`issue_summary`.`metadata`,`issue_summary`.`project_id`,`issue_summary`.`sprints`,`issue_summary`.`labels`,`issue_summary`.`top_level`,`issue_summary`.`is_leaf`,`issue_summary`.`path`,`issue_summary`.`in_progress_duration`,`issue_summary`.`verification_duration`,`issue_summary`.`in_progress_count`,`issue_summary`.`reopen_count`,`issue_summary`.`mapped_type`,`issue_summary`.`strategic_parent_id`,`issue_summary`.`sprint_id`,`issue_summary`.`issue_project_name`,`issue_summary`.`users`,`issue_summary`.`initial_start_date`,`issue_summary`.`total_duration`,`issue_summary`.`created_at`,`issue_summary`.`closed_at`,`issue_summary`.`planned_end_date`,`issue_summary`.`customer_id`,`issue_summary`.`ref_type`,`issue_summary`.`ref_id`,`issue_summary`.`custom_field_ids_virtual`,`issue_summary`.`release_duration`,`issue_summary`.`completed`,`issue_summary`.`completed_date`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `issue_summary` (`issue_summary`.`id`,`issue_summary`.`checksum`,`issue_summary`.`issue_id`,`issue_summary`.`total_issues`,`issue_summary`.`new30_days`,`issue_summary`.`total_closed`,`issue_summary`.`closed30_days`,`issue_summary`.`estimated_work_months`,`issue_summary`.`estimated_work_months30_days`,`issue_summary`.`title`,`issue_summary`.`url`,`issue_summary`.`priority`,`issue_summary`.`priority_id`,`issue_summary`.`status`,`issue_summary`.`status_id`,`issue_summary`.`issue_type`,`issue_summary`.`issue_type_id`,`issue_summary`.`resolution`,`issue_summary`.`resolution_id`,`issue_summary`.`state`,`issue_summary`.`custom_field_ids`,`issue_summary`.`teams`,`issue_summary`.`parent_issue_id`,`issue_summary`.`parents_issue_ids`,`issue_summary`.`metadata`,`issue_summary`.`project_id`,`issue_summary`.`sprints`,`issue_summary`.`labels`,`issue_summary`.`top_level`,`issue_summary`.`is_leaf`,`issue_summary`.`path`,`issue_summary`.`in_progress_duration`,`issue_summary`.`verification_duration`,`issue_summary`.`in_progress_count`,`issue_summary`.`reopen_count`,`issue_summary`.`mapped_type`,`issue_summary`.`strategic_parent_id`,`issue_summary`.`sprint_id`,`issue_summary`.`issue_project_name`,`issue_summary`.`users`,`issue_summary`.`initial_start_date`,`issue_summary`.`total_duration`,`issue_summary`.`created_at`,`issue_summary`.`closed_at`,`issue_summary`.`planned_end_date`,`issue_summary`.`customer_id`,`issue_summary`.`ref_type`,`issue_summary`.`ref_id`,`issue_summary`.`custom_field_ids_virtual`,`issue_summary`.`release_duration`,`issue_summary`.`completed`,`issue_summary`.`completed_date`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`issue_id`=VALUES(`issue_id`),`total_issues`=VALUES(`total_issues`),`new30_days`=VALUES(`new30_days`),`total_closed`=VALUES(`total_closed`),`closed30_days`=VALUES(`closed30_days`),`estimated_work_months`=VALUES(`estimated_work_months`),`estimated_work_months30_days`=VALUES(`estimated_work_months30_days`),`title`=VALUES(`title`),`url`=VALUES(`url`),`priority`=VALUES(`priority`),`priority_id`=VALUES(`priority_id`),`status`=VALUES(`status`),`status_id`=VALUES(`status_id`),`issue_type`=VALUES(`issue_type`),`issue_type_id`=VALUES(`issue_type_id`),`resolution`=VALUES(`resolution`),`resolution_id`=VALUES(`resolution_id`),`state`=VALUES(`state`),`custom_field_ids`=VALUES(`custom_field_ids`),`teams`=VALUES(`teams`),`parent_issue_id`=VALUES(`parent_issue_id`),`parents_issue_ids`=VALUES(`parents_issue_ids`),`metadata`=VALUES(`metadata`),`project_id`=VALUES(`project_id`),`sprints`=VALUES(`sprints`),`labels`=VALUES(`labels`),`top_level`=VALUES(`top_level`),`is_leaf`=VALUES(`is_leaf`),`path`=VALUES(`path`),`in_progress_duration`=VALUES(`in_progress_duration`),`verification_duration`=VALUES(`verification_duration`),`in_progress_count`=VALUES(`in_progress_count`),`reopen_count`=VALUES(`reopen_count`),`mapped_type`=VALUES(`mapped_type`),`strategic_parent_id`=VALUES(`strategic_parent_id`),`sprint_id`=VALUES(`sprint_id`),`issue_project_name`=VALUES(`issue_project_name`),`users`=VALUES(`users`),`initial_start_date`=VALUES(`initial_start_date`),`total_duration`=VALUES(`total_duration`),`created_at`=VALUES(`created_at`),`closed_at`=VALUES(`closed_at`),`planned_end_date`=VALUES(`planned_end_date`),`customer_id`=VALUES(`customer_id`),`ref_type`=VALUES(`ref_type`),`ref_id`=VALUES(`ref_id`),`custom_field_ids_virtual`=VALUES(`custom_field_ids_virtual`),`release_duration`=VALUES(`release_duration`),`completed`=VALUES(`completed`),`completed_date`=VALUES(`completed_date`)"
	}
	r, err := tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.IssueID),
		orm.ToSQLInt64(t.TotalIssues),
		orm.ToSQLInt64(t.New30Days),
		orm.ToSQLInt64(t.TotalClosed),
		orm.ToSQLInt64(t.Closed30Days),
		orm.ToSQLFloat64(t.EstimatedWorkMonths),
		orm.ToSQLFloat64(t.EstimatedWorkMonths30Days),
		orm.ToSQLString(t.Title),
		orm.ToSQLString(t.URL),
		orm.ToSQLString(t.Priority),
		orm.ToSQLString(t.PriorityID),
		orm.ToSQLString(t.Status),
		orm.ToSQLString(t.StatusID),
		orm.ToSQLString(t.IssueType),
		orm.ToSQLString(t.IssueTypeID),
		orm.ToSQLString(t.Resolution),
		orm.ToSQLString(t.ResolutionID),
		orm.ToSQLString(t.State),
		orm.ToSQLString(t.CustomFieldIds),
		orm.ToSQLString(t.Teams),
		orm.ToSQLString(t.ParentIssueID),
		orm.ToSQLString(t.ParentsIssueIds),
		orm.ToSQLString(t.Metadata),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLString(t.Sprints),
		orm.ToSQLString(t.Labels),
		orm.ToSQLBool(t.TopLevel),
		orm.ToSQLBool(t.IsLeaf),
		orm.ToSQLString(t.Path),
		orm.ToSQLInt64(t.InProgressDuration),
		orm.ToSQLInt64(t.VerificationDuration),
		orm.ToSQLInt64(t.InProgressCount),
		orm.ToSQLInt64(t.ReopenCount),
		orm.ToSQLString(t.MappedType),
		orm.ToSQLString(t.StrategicParentID),
		orm.ToSQLString(t.SprintID),
		orm.ToSQLString(t.IssueProjectName),
		orm.ToSQLString(t.Users),
		orm.ToSQLInt64(t.InitialStartDate),
		orm.ToSQLInt64(t.TotalDuration),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.ClosedAt),
		orm.ToSQLInt64(t.PlannedEndDate),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefType),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.CustomFieldIdsVirtual),
		orm.ToSQLInt64(t.ReleaseDuration),
		orm.ToSQLBool(t.Completed),
		orm.ToSQLInt64(t.CompletedDate),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBFindOne will find a IssueSummary record in the database with the primary key
func (t *IssueSummary) DBFindOne(ctx context.Context, db DB, value string) (bool, error) {
	q := "SELECT `issue_summary`.`id`,`issue_summary`.`checksum`,`issue_summary`.`issue_id`,`issue_summary`.`total_issues`,`issue_summary`.`new30_days`,`issue_summary`.`total_closed`,`issue_summary`.`closed30_days`,`issue_summary`.`estimated_work_months`,`issue_summary`.`estimated_work_months30_days`,`issue_summary`.`title`,`issue_summary`.`url`,`issue_summary`.`priority`,`issue_summary`.`priority_id`,`issue_summary`.`status`,`issue_summary`.`status_id`,`issue_summary`.`issue_type`,`issue_summary`.`issue_type_id`,`issue_summary`.`resolution`,`issue_summary`.`resolution_id`,`issue_summary`.`state`,`issue_summary`.`custom_field_ids`,`issue_summary`.`teams`,`issue_summary`.`parent_issue_id`,`issue_summary`.`parents_issue_ids`,`issue_summary`.`metadata`,`issue_summary`.`project_id`,`issue_summary`.`sprints`,`issue_summary`.`labels`,`issue_summary`.`top_level`,`issue_summary`.`is_leaf`,`issue_summary`.`path`,`issue_summary`.`in_progress_duration`,`issue_summary`.`verification_duration`,`issue_summary`.`in_progress_count`,`issue_summary`.`reopen_count`,`issue_summary`.`mapped_type`,`issue_summary`.`strategic_parent_id`,`issue_summary`.`sprint_id`,`issue_summary`.`issue_project_name`,`issue_summary`.`users`,`issue_summary`.`initial_start_date`,`issue_summary`.`total_duration`,`issue_summary`.`created_at`,`issue_summary`.`closed_at`,`issue_summary`.`planned_end_date`,`issue_summary`.`customer_id`,`issue_summary`.`ref_type`,`issue_summary`.`ref_id`,`issue_summary`.`custom_field_ids_virtual`,`issue_summary`.`release_duration`,`issue_summary`.`completed`,`issue_summary`.`completed_date` FROM `issue_summary` WHERE `id` = ? LIMIT 1"
	row := db.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _IssueID sql.NullString
	var _TotalIssues sql.NullInt64
	var _New30Days sql.NullInt64
	var _TotalClosed sql.NullInt64
	var _Closed30Days sql.NullInt64
	var _EstimatedWorkMonths sql.NullFloat64
	var _EstimatedWorkMonths30Days sql.NullFloat64
	var _Title sql.NullString
	var _URL sql.NullString
	var _Priority sql.NullString
	var _PriorityID sql.NullString
	var _Status sql.NullString
	var _StatusID sql.NullString
	var _IssueType sql.NullString
	var _IssueTypeID sql.NullString
	var _Resolution sql.NullString
	var _ResolutionID sql.NullString
	var _State sql.NullString
	var _CustomFieldIds sql.NullString
	var _Teams sql.NullString
	var _ParentIssueID sql.NullString
	var _ParentsIssueIds sql.NullString
	var _Metadata sql.NullString
	var _ProjectID sql.NullString
	var _Sprints sql.NullString
	var _Labels sql.NullString
	var _TopLevel sql.NullBool
	var _IsLeaf sql.NullBool
	var _Path sql.NullString
	var _InProgressDuration sql.NullInt64
	var _VerificationDuration sql.NullInt64
	var _InProgressCount sql.NullInt64
	var _ReopenCount sql.NullInt64
	var _MappedType sql.NullString
	var _StrategicParentID sql.NullString
	var _SprintID sql.NullString
	var _IssueProjectName sql.NullString
	var _Users sql.NullString
	var _InitialStartDate sql.NullInt64
	var _TotalDuration sql.NullInt64
	var _CreatedAt sql.NullInt64
	var _ClosedAt sql.NullInt64
	var _PlannedEndDate sql.NullInt64
	var _CustomerID sql.NullString
	var _RefType sql.NullString
	var _RefID sql.NullString
	var _CustomFieldIdsVirtual sql.NullString
	var _ReleaseDuration sql.NullInt64
	var _Completed sql.NullBool
	var _CompletedDate sql.NullInt64
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_IssueID,
		&_TotalIssues,
		&_New30Days,
		&_TotalClosed,
		&_Closed30Days,
		&_EstimatedWorkMonths,
		&_EstimatedWorkMonths30Days,
		&_Title,
		&_URL,
		&_Priority,
		&_PriorityID,
		&_Status,
		&_StatusID,
		&_IssueType,
		&_IssueTypeID,
		&_Resolution,
		&_ResolutionID,
		&_State,
		&_CustomFieldIds,
		&_Teams,
		&_ParentIssueID,
		&_ParentsIssueIds,
		&_Metadata,
		&_ProjectID,
		&_Sprints,
		&_Labels,
		&_TopLevel,
		&_IsLeaf,
		&_Path,
		&_InProgressDuration,
		&_VerificationDuration,
		&_InProgressCount,
		&_ReopenCount,
		&_MappedType,
		&_StrategicParentID,
		&_SprintID,
		&_IssueProjectName,
		&_Users,
		&_InitialStartDate,
		&_TotalDuration,
		&_CreatedAt,
		&_ClosedAt,
		&_PlannedEndDate,
		&_CustomerID,
		&_RefType,
		&_RefID,
		&_CustomFieldIdsVirtual,
		&_ReleaseDuration,
		&_Completed,
		&_CompletedDate,
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
	if _TotalIssues.Valid {
		t.SetTotalIssues(int32(_TotalIssues.Int64))
	}
	if _New30Days.Valid {
		t.SetNew30Days(int32(_New30Days.Int64))
	}
	if _TotalClosed.Valid {
		t.SetTotalClosed(int32(_TotalClosed.Int64))
	}
	if _Closed30Days.Valid {
		t.SetClosed30Days(int32(_Closed30Days.Int64))
	}
	if _EstimatedWorkMonths.Valid {
		t.SetEstimatedWorkMonths(_EstimatedWorkMonths.Float64)
	}
	if _EstimatedWorkMonths30Days.Valid {
		t.SetEstimatedWorkMonths30Days(_EstimatedWorkMonths30Days.Float64)
	}
	if _Title.Valid {
		t.SetTitle(_Title.String)
	}
	if _URL.Valid {
		t.SetURL(_URL.String)
	}
	if _Priority.Valid {
		t.SetPriority(_Priority.String)
	}
	if _PriorityID.Valid {
		t.SetPriorityID(_PriorityID.String)
	}
	if _Status.Valid {
		t.SetStatus(_Status.String)
	}
	if _StatusID.Valid {
		t.SetStatusID(_StatusID.String)
	}
	if _IssueType.Valid {
		t.SetIssueType(_IssueType.String)
	}
	if _IssueTypeID.Valid {
		t.SetIssueTypeID(_IssueTypeID.String)
	}
	if _Resolution.Valid {
		t.SetResolution(_Resolution.String)
	}
	if _ResolutionID.Valid {
		t.SetResolutionID(_ResolutionID.String)
	}
	if _State.Valid {
		t.SetState(_State.String)
	}
	if _CustomFieldIds.Valid {
		t.SetCustomFieldIds(_CustomFieldIds.String)
	}
	if _Teams.Valid {
		t.SetTeams(_Teams.String)
	}
	if _ParentIssueID.Valid {
		t.SetParentIssueID(_ParentIssueID.String)
	}
	if _ParentsIssueIds.Valid {
		t.SetParentsIssueIds(_ParentsIssueIds.String)
	}
	if _Metadata.Valid {
		t.SetMetadata(_Metadata.String)
	}
	if _ProjectID.Valid {
		t.SetProjectID(_ProjectID.String)
	}
	if _Sprints.Valid {
		t.SetSprints(_Sprints.String)
	}
	if _Labels.Valid {
		t.SetLabels(_Labels.String)
	}
	if _TopLevel.Valid {
		t.SetTopLevel(_TopLevel.Bool)
	}
	if _IsLeaf.Valid {
		t.SetIsLeaf(_IsLeaf.Bool)
	}
	if _Path.Valid {
		t.SetPath(_Path.String)
	}
	if _InProgressDuration.Valid {
		t.SetInProgressDuration(_InProgressDuration.Int64)
	}
	if _VerificationDuration.Valid {
		t.SetVerificationDuration(_VerificationDuration.Int64)
	}
	if _InProgressCount.Valid {
		t.SetInProgressCount(int32(_InProgressCount.Int64))
	}
	if _ReopenCount.Valid {
		t.SetReopenCount(int32(_ReopenCount.Int64))
	}
	if _MappedType.Valid {
		t.SetMappedType(_MappedType.String)
	}
	if _StrategicParentID.Valid {
		t.SetStrategicParentID(_StrategicParentID.String)
	}
	if _SprintID.Valid {
		t.SetSprintID(_SprintID.String)
	}
	if _IssueProjectName.Valid {
		t.SetIssueProjectName(_IssueProjectName.String)
	}
	if _Users.Valid {
		t.SetUsers(_Users.String)
	}
	if _InitialStartDate.Valid {
		t.SetInitialStartDate(_InitialStartDate.Int64)
	}
	if _TotalDuration.Valid {
		t.SetTotalDuration(_TotalDuration.Int64)
	}
	if _CreatedAt.Valid {
		t.SetCreatedAt(_CreatedAt.Int64)
	}
	if _ClosedAt.Valid {
		t.SetClosedAt(_ClosedAt.Int64)
	}
	if _PlannedEndDate.Valid {
		t.SetPlannedEndDate(_PlannedEndDate.Int64)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _RefType.Valid {
		t.SetRefType(_RefType.String)
	}
	if _RefID.Valid {
		t.SetRefID(_RefID.String)
	}
	if _CustomFieldIdsVirtual.Valid {
		t.SetCustomFieldIdsVirtual(_CustomFieldIdsVirtual.String)
	}
	if _ReleaseDuration.Valid {
		t.SetReleaseDuration(_ReleaseDuration.Int64)
	}
	if _Completed.Valid {
		t.SetCompleted(_Completed.Bool)
	}
	if _CompletedDate.Valid {
		t.SetCompletedDate(_CompletedDate.Int64)
	}
	return true, nil
}

// DBFindOneTx will find a IssueSummary record in the database with the primary key using the provided transaction
func (t *IssueSummary) DBFindOneTx(ctx context.Context, tx Tx, value string) (bool, error) {
	q := "SELECT `issue_summary`.`id`,`issue_summary`.`checksum`,`issue_summary`.`issue_id`,`issue_summary`.`total_issues`,`issue_summary`.`new30_days`,`issue_summary`.`total_closed`,`issue_summary`.`closed30_days`,`issue_summary`.`estimated_work_months`,`issue_summary`.`estimated_work_months30_days`,`issue_summary`.`title`,`issue_summary`.`url`,`issue_summary`.`priority`,`issue_summary`.`priority_id`,`issue_summary`.`status`,`issue_summary`.`status_id`,`issue_summary`.`issue_type`,`issue_summary`.`issue_type_id`,`issue_summary`.`resolution`,`issue_summary`.`resolution_id`,`issue_summary`.`state`,`issue_summary`.`custom_field_ids`,`issue_summary`.`teams`,`issue_summary`.`parent_issue_id`,`issue_summary`.`parents_issue_ids`,`issue_summary`.`metadata`,`issue_summary`.`project_id`,`issue_summary`.`sprints`,`issue_summary`.`labels`,`issue_summary`.`top_level`,`issue_summary`.`is_leaf`,`issue_summary`.`path`,`issue_summary`.`in_progress_duration`,`issue_summary`.`verification_duration`,`issue_summary`.`in_progress_count`,`issue_summary`.`reopen_count`,`issue_summary`.`mapped_type`,`issue_summary`.`strategic_parent_id`,`issue_summary`.`sprint_id`,`issue_summary`.`issue_project_name`,`issue_summary`.`users`,`issue_summary`.`initial_start_date`,`issue_summary`.`total_duration`,`issue_summary`.`created_at`,`issue_summary`.`closed_at`,`issue_summary`.`planned_end_date`,`issue_summary`.`customer_id`,`issue_summary`.`ref_type`,`issue_summary`.`ref_id`,`issue_summary`.`custom_field_ids_virtual`,`issue_summary`.`release_duration`,`issue_summary`.`completed`,`issue_summary`.`completed_date` FROM `issue_summary` WHERE `id` = ? LIMIT 1"
	row := tx.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _IssueID sql.NullString
	var _TotalIssues sql.NullInt64
	var _New30Days sql.NullInt64
	var _TotalClosed sql.NullInt64
	var _Closed30Days sql.NullInt64
	var _EstimatedWorkMonths sql.NullFloat64
	var _EstimatedWorkMonths30Days sql.NullFloat64
	var _Title sql.NullString
	var _URL sql.NullString
	var _Priority sql.NullString
	var _PriorityID sql.NullString
	var _Status sql.NullString
	var _StatusID sql.NullString
	var _IssueType sql.NullString
	var _IssueTypeID sql.NullString
	var _Resolution sql.NullString
	var _ResolutionID sql.NullString
	var _State sql.NullString
	var _CustomFieldIds sql.NullString
	var _Teams sql.NullString
	var _ParentIssueID sql.NullString
	var _ParentsIssueIds sql.NullString
	var _Metadata sql.NullString
	var _ProjectID sql.NullString
	var _Sprints sql.NullString
	var _Labels sql.NullString
	var _TopLevel sql.NullBool
	var _IsLeaf sql.NullBool
	var _Path sql.NullString
	var _InProgressDuration sql.NullInt64
	var _VerificationDuration sql.NullInt64
	var _InProgressCount sql.NullInt64
	var _ReopenCount sql.NullInt64
	var _MappedType sql.NullString
	var _StrategicParentID sql.NullString
	var _SprintID sql.NullString
	var _IssueProjectName sql.NullString
	var _Users sql.NullString
	var _InitialStartDate sql.NullInt64
	var _TotalDuration sql.NullInt64
	var _CreatedAt sql.NullInt64
	var _ClosedAt sql.NullInt64
	var _PlannedEndDate sql.NullInt64
	var _CustomerID sql.NullString
	var _RefType sql.NullString
	var _RefID sql.NullString
	var _CustomFieldIdsVirtual sql.NullString
	var _ReleaseDuration sql.NullInt64
	var _Completed sql.NullBool
	var _CompletedDate sql.NullInt64
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_IssueID,
		&_TotalIssues,
		&_New30Days,
		&_TotalClosed,
		&_Closed30Days,
		&_EstimatedWorkMonths,
		&_EstimatedWorkMonths30Days,
		&_Title,
		&_URL,
		&_Priority,
		&_PriorityID,
		&_Status,
		&_StatusID,
		&_IssueType,
		&_IssueTypeID,
		&_Resolution,
		&_ResolutionID,
		&_State,
		&_CustomFieldIds,
		&_Teams,
		&_ParentIssueID,
		&_ParentsIssueIds,
		&_Metadata,
		&_ProjectID,
		&_Sprints,
		&_Labels,
		&_TopLevel,
		&_IsLeaf,
		&_Path,
		&_InProgressDuration,
		&_VerificationDuration,
		&_InProgressCount,
		&_ReopenCount,
		&_MappedType,
		&_StrategicParentID,
		&_SprintID,
		&_IssueProjectName,
		&_Users,
		&_InitialStartDate,
		&_TotalDuration,
		&_CreatedAt,
		&_ClosedAt,
		&_PlannedEndDate,
		&_CustomerID,
		&_RefType,
		&_RefID,
		&_CustomFieldIdsVirtual,
		&_ReleaseDuration,
		&_Completed,
		&_CompletedDate,
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
	if _TotalIssues.Valid {
		t.SetTotalIssues(int32(_TotalIssues.Int64))
	}
	if _New30Days.Valid {
		t.SetNew30Days(int32(_New30Days.Int64))
	}
	if _TotalClosed.Valid {
		t.SetTotalClosed(int32(_TotalClosed.Int64))
	}
	if _Closed30Days.Valid {
		t.SetClosed30Days(int32(_Closed30Days.Int64))
	}
	if _EstimatedWorkMonths.Valid {
		t.SetEstimatedWorkMonths(_EstimatedWorkMonths.Float64)
	}
	if _EstimatedWorkMonths30Days.Valid {
		t.SetEstimatedWorkMonths30Days(_EstimatedWorkMonths30Days.Float64)
	}
	if _Title.Valid {
		t.SetTitle(_Title.String)
	}
	if _URL.Valid {
		t.SetURL(_URL.String)
	}
	if _Priority.Valid {
		t.SetPriority(_Priority.String)
	}
	if _PriorityID.Valid {
		t.SetPriorityID(_PriorityID.String)
	}
	if _Status.Valid {
		t.SetStatus(_Status.String)
	}
	if _StatusID.Valid {
		t.SetStatusID(_StatusID.String)
	}
	if _IssueType.Valid {
		t.SetIssueType(_IssueType.String)
	}
	if _IssueTypeID.Valid {
		t.SetIssueTypeID(_IssueTypeID.String)
	}
	if _Resolution.Valid {
		t.SetResolution(_Resolution.String)
	}
	if _ResolutionID.Valid {
		t.SetResolutionID(_ResolutionID.String)
	}
	if _State.Valid {
		t.SetState(_State.String)
	}
	if _CustomFieldIds.Valid {
		t.SetCustomFieldIds(_CustomFieldIds.String)
	}
	if _Teams.Valid {
		t.SetTeams(_Teams.String)
	}
	if _ParentIssueID.Valid {
		t.SetParentIssueID(_ParentIssueID.String)
	}
	if _ParentsIssueIds.Valid {
		t.SetParentsIssueIds(_ParentsIssueIds.String)
	}
	if _Metadata.Valid {
		t.SetMetadata(_Metadata.String)
	}
	if _ProjectID.Valid {
		t.SetProjectID(_ProjectID.String)
	}
	if _Sprints.Valid {
		t.SetSprints(_Sprints.String)
	}
	if _Labels.Valid {
		t.SetLabels(_Labels.String)
	}
	if _TopLevel.Valid {
		t.SetTopLevel(_TopLevel.Bool)
	}
	if _IsLeaf.Valid {
		t.SetIsLeaf(_IsLeaf.Bool)
	}
	if _Path.Valid {
		t.SetPath(_Path.String)
	}
	if _InProgressDuration.Valid {
		t.SetInProgressDuration(_InProgressDuration.Int64)
	}
	if _VerificationDuration.Valid {
		t.SetVerificationDuration(_VerificationDuration.Int64)
	}
	if _InProgressCount.Valid {
		t.SetInProgressCount(int32(_InProgressCount.Int64))
	}
	if _ReopenCount.Valid {
		t.SetReopenCount(int32(_ReopenCount.Int64))
	}
	if _MappedType.Valid {
		t.SetMappedType(_MappedType.String)
	}
	if _StrategicParentID.Valid {
		t.SetStrategicParentID(_StrategicParentID.String)
	}
	if _SprintID.Valid {
		t.SetSprintID(_SprintID.String)
	}
	if _IssueProjectName.Valid {
		t.SetIssueProjectName(_IssueProjectName.String)
	}
	if _Users.Valid {
		t.SetUsers(_Users.String)
	}
	if _InitialStartDate.Valid {
		t.SetInitialStartDate(_InitialStartDate.Int64)
	}
	if _TotalDuration.Valid {
		t.SetTotalDuration(_TotalDuration.Int64)
	}
	if _CreatedAt.Valid {
		t.SetCreatedAt(_CreatedAt.Int64)
	}
	if _ClosedAt.Valid {
		t.SetClosedAt(_ClosedAt.Int64)
	}
	if _PlannedEndDate.Valid {
		t.SetPlannedEndDate(_PlannedEndDate.Int64)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _RefType.Valid {
		t.SetRefType(_RefType.String)
	}
	if _RefID.Valid {
		t.SetRefID(_RefID.String)
	}
	if _CustomFieldIdsVirtual.Valid {
		t.SetCustomFieldIdsVirtual(_CustomFieldIdsVirtual.String)
	}
	if _ReleaseDuration.Valid {
		t.SetReleaseDuration(_ReleaseDuration.Int64)
	}
	if _Completed.Valid {
		t.SetCompleted(_Completed.Bool)
	}
	if _CompletedDate.Valid {
		t.SetCompletedDate(_CompletedDate.Int64)
	}
	return true, nil
}

// FindIssueSummaries will find a IssueSummary record in the database with the provided parameters
func FindIssueSummaries(ctx context.Context, db DB, _params ...interface{}) ([]*IssueSummary, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("issue_id"),
		orm.Column("total_issues"),
		orm.Column("new30_days"),
		orm.Column("total_closed"),
		orm.Column("closed30_days"),
		orm.Column("estimated_work_months"),
		orm.Column("estimated_work_months30_days"),
		orm.Column("title"),
		orm.Column("url"),
		orm.Column("priority"),
		orm.Column("priority_id"),
		orm.Column("status"),
		orm.Column("status_id"),
		orm.Column("issue_type"),
		orm.Column("issue_type_id"),
		orm.Column("resolution"),
		orm.Column("resolution_id"),
		orm.Column("state"),
		orm.Column("custom_field_ids"),
		orm.Column("teams"),
		orm.Column("parent_issue_id"),
		orm.Column("parents_issue_ids"),
		orm.Column("metadata"),
		orm.Column("project_id"),
		orm.Column("sprints"),
		orm.Column("labels"),
		orm.Column("top_level"),
		orm.Column("is_leaf"),
		orm.Column("path"),
		orm.Column("in_progress_duration"),
		orm.Column("verification_duration"),
		orm.Column("in_progress_count"),
		orm.Column("reopen_count"),
		orm.Column("mapped_type"),
		orm.Column("strategic_parent_id"),
		orm.Column("sprint_id"),
		orm.Column("issue_project_name"),
		orm.Column("users"),
		orm.Column("initial_start_date"),
		orm.Column("total_duration"),
		orm.Column("created_at"),
		orm.Column("closed_at"),
		orm.Column("planned_end_date"),
		orm.Column("customer_id"),
		orm.Column("ref_type"),
		orm.Column("ref_id"),
		orm.Column("custom_field_ids_virtual"),
		orm.Column("release_duration"),
		orm.Column("completed"),
		orm.Column("completed_date"),
		orm.Table(IssueSummaryTableName),
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
	results := make([]*IssueSummary, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _IssueID sql.NullString
		var _TotalIssues sql.NullInt64
		var _New30Days sql.NullInt64
		var _TotalClosed sql.NullInt64
		var _Closed30Days sql.NullInt64
		var _EstimatedWorkMonths sql.NullFloat64
		var _EstimatedWorkMonths30Days sql.NullFloat64
		var _Title sql.NullString
		var _URL sql.NullString
		var _Priority sql.NullString
		var _PriorityID sql.NullString
		var _Status sql.NullString
		var _StatusID sql.NullString
		var _IssueType sql.NullString
		var _IssueTypeID sql.NullString
		var _Resolution sql.NullString
		var _ResolutionID sql.NullString
		var _State sql.NullString
		var _CustomFieldIds sql.NullString
		var _Teams sql.NullString
		var _ParentIssueID sql.NullString
		var _ParentsIssueIds sql.NullString
		var _Metadata sql.NullString
		var _ProjectID sql.NullString
		var _Sprints sql.NullString
		var _Labels sql.NullString
		var _TopLevel sql.NullBool
		var _IsLeaf sql.NullBool
		var _Path sql.NullString
		var _InProgressDuration sql.NullInt64
		var _VerificationDuration sql.NullInt64
		var _InProgressCount sql.NullInt64
		var _ReopenCount sql.NullInt64
		var _MappedType sql.NullString
		var _StrategicParentID sql.NullString
		var _SprintID sql.NullString
		var _IssueProjectName sql.NullString
		var _Users sql.NullString
		var _InitialStartDate sql.NullInt64
		var _TotalDuration sql.NullInt64
		var _CreatedAt sql.NullInt64
		var _ClosedAt sql.NullInt64
		var _PlannedEndDate sql.NullInt64
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefID sql.NullString
		var _CustomFieldIdsVirtual sql.NullString
		var _ReleaseDuration sql.NullInt64
		var _Completed sql.NullBool
		var _CompletedDate sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_IssueID,
			&_TotalIssues,
			&_New30Days,
			&_TotalClosed,
			&_Closed30Days,
			&_EstimatedWorkMonths,
			&_EstimatedWorkMonths30Days,
			&_Title,
			&_URL,
			&_Priority,
			&_PriorityID,
			&_Status,
			&_StatusID,
			&_IssueType,
			&_IssueTypeID,
			&_Resolution,
			&_ResolutionID,
			&_State,
			&_CustomFieldIds,
			&_Teams,
			&_ParentIssueID,
			&_ParentsIssueIds,
			&_Metadata,
			&_ProjectID,
			&_Sprints,
			&_Labels,
			&_TopLevel,
			&_IsLeaf,
			&_Path,
			&_InProgressDuration,
			&_VerificationDuration,
			&_InProgressCount,
			&_ReopenCount,
			&_MappedType,
			&_StrategicParentID,
			&_SprintID,
			&_IssueProjectName,
			&_Users,
			&_InitialStartDate,
			&_TotalDuration,
			&_CreatedAt,
			&_ClosedAt,
			&_PlannedEndDate,
			&_CustomerID,
			&_RefType,
			&_RefID,
			&_CustomFieldIdsVirtual,
			&_ReleaseDuration,
			&_Completed,
			&_CompletedDate,
		)
		if err != nil {
			return nil, err
		}
		t := &IssueSummary{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _TotalIssues.Valid {
			t.SetTotalIssues(int32(_TotalIssues.Int64))
		}
		if _New30Days.Valid {
			t.SetNew30Days(int32(_New30Days.Int64))
		}
		if _TotalClosed.Valid {
			t.SetTotalClosed(int32(_TotalClosed.Int64))
		}
		if _Closed30Days.Valid {
			t.SetClosed30Days(int32(_Closed30Days.Int64))
		}
		if _EstimatedWorkMonths.Valid {
			t.SetEstimatedWorkMonths(_EstimatedWorkMonths.Float64)
		}
		if _EstimatedWorkMonths30Days.Valid {
			t.SetEstimatedWorkMonths30Days(_EstimatedWorkMonths30Days.Float64)
		}
		if _Title.Valid {
			t.SetTitle(_Title.String)
		}
		if _URL.Valid {
			t.SetURL(_URL.String)
		}
		if _Priority.Valid {
			t.SetPriority(_Priority.String)
		}
		if _PriorityID.Valid {
			t.SetPriorityID(_PriorityID.String)
		}
		if _Status.Valid {
			t.SetStatus(_Status.String)
		}
		if _StatusID.Valid {
			t.SetStatusID(_StatusID.String)
		}
		if _IssueType.Valid {
			t.SetIssueType(_IssueType.String)
		}
		if _IssueTypeID.Valid {
			t.SetIssueTypeID(_IssueTypeID.String)
		}
		if _Resolution.Valid {
			t.SetResolution(_Resolution.String)
		}
		if _ResolutionID.Valid {
			t.SetResolutionID(_ResolutionID.String)
		}
		if _State.Valid {
			t.SetState(_State.String)
		}
		if _CustomFieldIds.Valid {
			t.SetCustomFieldIds(_CustomFieldIds.String)
		}
		if _Teams.Valid {
			t.SetTeams(_Teams.String)
		}
		if _ParentIssueID.Valid {
			t.SetParentIssueID(_ParentIssueID.String)
		}
		if _ParentsIssueIds.Valid {
			t.SetParentsIssueIds(_ParentsIssueIds.String)
		}
		if _Metadata.Valid {
			t.SetMetadata(_Metadata.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _Sprints.Valid {
			t.SetSprints(_Sprints.String)
		}
		if _Labels.Valid {
			t.SetLabels(_Labels.String)
		}
		if _TopLevel.Valid {
			t.SetTopLevel(_TopLevel.Bool)
		}
		if _IsLeaf.Valid {
			t.SetIsLeaf(_IsLeaf.Bool)
		}
		if _Path.Valid {
			t.SetPath(_Path.String)
		}
		if _InProgressDuration.Valid {
			t.SetInProgressDuration(_InProgressDuration.Int64)
		}
		if _VerificationDuration.Valid {
			t.SetVerificationDuration(_VerificationDuration.Int64)
		}
		if _InProgressCount.Valid {
			t.SetInProgressCount(int32(_InProgressCount.Int64))
		}
		if _ReopenCount.Valid {
			t.SetReopenCount(int32(_ReopenCount.Int64))
		}
		if _MappedType.Valid {
			t.SetMappedType(_MappedType.String)
		}
		if _StrategicParentID.Valid {
			t.SetStrategicParentID(_StrategicParentID.String)
		}
		if _SprintID.Valid {
			t.SetSprintID(_SprintID.String)
		}
		if _IssueProjectName.Valid {
			t.SetIssueProjectName(_IssueProjectName.String)
		}
		if _Users.Valid {
			t.SetUsers(_Users.String)
		}
		if _InitialStartDate.Valid {
			t.SetInitialStartDate(_InitialStartDate.Int64)
		}
		if _TotalDuration.Valid {
			t.SetTotalDuration(_TotalDuration.Int64)
		}
		if _CreatedAt.Valid {
			t.SetCreatedAt(_CreatedAt.Int64)
		}
		if _ClosedAt.Valid {
			t.SetClosedAt(_ClosedAt.Int64)
		}
		if _PlannedEndDate.Valid {
			t.SetPlannedEndDate(_PlannedEndDate.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefType.Valid {
			t.SetRefType(_RefType.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		if _CustomFieldIdsVirtual.Valid {
			t.SetCustomFieldIdsVirtual(_CustomFieldIdsVirtual.String)
		}
		if _ReleaseDuration.Valid {
			t.SetReleaseDuration(_ReleaseDuration.Int64)
		}
		if _Completed.Valid {
			t.SetCompleted(_Completed.Bool)
		}
		if _CompletedDate.Valid {
			t.SetCompletedDate(_CompletedDate.Int64)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindIssueSummariesTx will find a IssueSummary record in the database with the provided parameters using the provided transaction
func FindIssueSummariesTx(ctx context.Context, tx Tx, _params ...interface{}) ([]*IssueSummary, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("issue_id"),
		orm.Column("total_issues"),
		orm.Column("new30_days"),
		orm.Column("total_closed"),
		orm.Column("closed30_days"),
		orm.Column("estimated_work_months"),
		orm.Column("estimated_work_months30_days"),
		orm.Column("title"),
		orm.Column("url"),
		orm.Column("priority"),
		orm.Column("priority_id"),
		orm.Column("status"),
		orm.Column("status_id"),
		orm.Column("issue_type"),
		orm.Column("issue_type_id"),
		orm.Column("resolution"),
		orm.Column("resolution_id"),
		orm.Column("state"),
		orm.Column("custom_field_ids"),
		orm.Column("teams"),
		orm.Column("parent_issue_id"),
		orm.Column("parents_issue_ids"),
		orm.Column("metadata"),
		orm.Column("project_id"),
		orm.Column("sprints"),
		orm.Column("labels"),
		orm.Column("top_level"),
		orm.Column("is_leaf"),
		orm.Column("path"),
		orm.Column("in_progress_duration"),
		orm.Column("verification_duration"),
		orm.Column("in_progress_count"),
		orm.Column("reopen_count"),
		orm.Column("mapped_type"),
		orm.Column("strategic_parent_id"),
		orm.Column("sprint_id"),
		orm.Column("issue_project_name"),
		orm.Column("users"),
		orm.Column("initial_start_date"),
		orm.Column("total_duration"),
		orm.Column("created_at"),
		orm.Column("closed_at"),
		orm.Column("planned_end_date"),
		orm.Column("customer_id"),
		orm.Column("ref_type"),
		orm.Column("ref_id"),
		orm.Column("custom_field_ids_virtual"),
		orm.Column("release_duration"),
		orm.Column("completed"),
		orm.Column("completed_date"),
		orm.Table(IssueSummaryTableName),
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
	results := make([]*IssueSummary, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _IssueID sql.NullString
		var _TotalIssues sql.NullInt64
		var _New30Days sql.NullInt64
		var _TotalClosed sql.NullInt64
		var _Closed30Days sql.NullInt64
		var _EstimatedWorkMonths sql.NullFloat64
		var _EstimatedWorkMonths30Days sql.NullFloat64
		var _Title sql.NullString
		var _URL sql.NullString
		var _Priority sql.NullString
		var _PriorityID sql.NullString
		var _Status sql.NullString
		var _StatusID sql.NullString
		var _IssueType sql.NullString
		var _IssueTypeID sql.NullString
		var _Resolution sql.NullString
		var _ResolutionID sql.NullString
		var _State sql.NullString
		var _CustomFieldIds sql.NullString
		var _Teams sql.NullString
		var _ParentIssueID sql.NullString
		var _ParentsIssueIds sql.NullString
		var _Metadata sql.NullString
		var _ProjectID sql.NullString
		var _Sprints sql.NullString
		var _Labels sql.NullString
		var _TopLevel sql.NullBool
		var _IsLeaf sql.NullBool
		var _Path sql.NullString
		var _InProgressDuration sql.NullInt64
		var _VerificationDuration sql.NullInt64
		var _InProgressCount sql.NullInt64
		var _ReopenCount sql.NullInt64
		var _MappedType sql.NullString
		var _StrategicParentID sql.NullString
		var _SprintID sql.NullString
		var _IssueProjectName sql.NullString
		var _Users sql.NullString
		var _InitialStartDate sql.NullInt64
		var _TotalDuration sql.NullInt64
		var _CreatedAt sql.NullInt64
		var _ClosedAt sql.NullInt64
		var _PlannedEndDate sql.NullInt64
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefID sql.NullString
		var _CustomFieldIdsVirtual sql.NullString
		var _ReleaseDuration sql.NullInt64
		var _Completed sql.NullBool
		var _CompletedDate sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_IssueID,
			&_TotalIssues,
			&_New30Days,
			&_TotalClosed,
			&_Closed30Days,
			&_EstimatedWorkMonths,
			&_EstimatedWorkMonths30Days,
			&_Title,
			&_URL,
			&_Priority,
			&_PriorityID,
			&_Status,
			&_StatusID,
			&_IssueType,
			&_IssueTypeID,
			&_Resolution,
			&_ResolutionID,
			&_State,
			&_CustomFieldIds,
			&_Teams,
			&_ParentIssueID,
			&_ParentsIssueIds,
			&_Metadata,
			&_ProjectID,
			&_Sprints,
			&_Labels,
			&_TopLevel,
			&_IsLeaf,
			&_Path,
			&_InProgressDuration,
			&_VerificationDuration,
			&_InProgressCount,
			&_ReopenCount,
			&_MappedType,
			&_StrategicParentID,
			&_SprintID,
			&_IssueProjectName,
			&_Users,
			&_InitialStartDate,
			&_TotalDuration,
			&_CreatedAt,
			&_ClosedAt,
			&_PlannedEndDate,
			&_CustomerID,
			&_RefType,
			&_RefID,
			&_CustomFieldIdsVirtual,
			&_ReleaseDuration,
			&_Completed,
			&_CompletedDate,
		)
		if err != nil {
			return nil, err
		}
		t := &IssueSummary{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _TotalIssues.Valid {
			t.SetTotalIssues(int32(_TotalIssues.Int64))
		}
		if _New30Days.Valid {
			t.SetNew30Days(int32(_New30Days.Int64))
		}
		if _TotalClosed.Valid {
			t.SetTotalClosed(int32(_TotalClosed.Int64))
		}
		if _Closed30Days.Valid {
			t.SetClosed30Days(int32(_Closed30Days.Int64))
		}
		if _EstimatedWorkMonths.Valid {
			t.SetEstimatedWorkMonths(_EstimatedWorkMonths.Float64)
		}
		if _EstimatedWorkMonths30Days.Valid {
			t.SetEstimatedWorkMonths30Days(_EstimatedWorkMonths30Days.Float64)
		}
		if _Title.Valid {
			t.SetTitle(_Title.String)
		}
		if _URL.Valid {
			t.SetURL(_URL.String)
		}
		if _Priority.Valid {
			t.SetPriority(_Priority.String)
		}
		if _PriorityID.Valid {
			t.SetPriorityID(_PriorityID.String)
		}
		if _Status.Valid {
			t.SetStatus(_Status.String)
		}
		if _StatusID.Valid {
			t.SetStatusID(_StatusID.String)
		}
		if _IssueType.Valid {
			t.SetIssueType(_IssueType.String)
		}
		if _IssueTypeID.Valid {
			t.SetIssueTypeID(_IssueTypeID.String)
		}
		if _Resolution.Valid {
			t.SetResolution(_Resolution.String)
		}
		if _ResolutionID.Valid {
			t.SetResolutionID(_ResolutionID.String)
		}
		if _State.Valid {
			t.SetState(_State.String)
		}
		if _CustomFieldIds.Valid {
			t.SetCustomFieldIds(_CustomFieldIds.String)
		}
		if _Teams.Valid {
			t.SetTeams(_Teams.String)
		}
		if _ParentIssueID.Valid {
			t.SetParentIssueID(_ParentIssueID.String)
		}
		if _ParentsIssueIds.Valid {
			t.SetParentsIssueIds(_ParentsIssueIds.String)
		}
		if _Metadata.Valid {
			t.SetMetadata(_Metadata.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _Sprints.Valid {
			t.SetSprints(_Sprints.String)
		}
		if _Labels.Valid {
			t.SetLabels(_Labels.String)
		}
		if _TopLevel.Valid {
			t.SetTopLevel(_TopLevel.Bool)
		}
		if _IsLeaf.Valid {
			t.SetIsLeaf(_IsLeaf.Bool)
		}
		if _Path.Valid {
			t.SetPath(_Path.String)
		}
		if _InProgressDuration.Valid {
			t.SetInProgressDuration(_InProgressDuration.Int64)
		}
		if _VerificationDuration.Valid {
			t.SetVerificationDuration(_VerificationDuration.Int64)
		}
		if _InProgressCount.Valid {
			t.SetInProgressCount(int32(_InProgressCount.Int64))
		}
		if _ReopenCount.Valid {
			t.SetReopenCount(int32(_ReopenCount.Int64))
		}
		if _MappedType.Valid {
			t.SetMappedType(_MappedType.String)
		}
		if _StrategicParentID.Valid {
			t.SetStrategicParentID(_StrategicParentID.String)
		}
		if _SprintID.Valid {
			t.SetSprintID(_SprintID.String)
		}
		if _IssueProjectName.Valid {
			t.SetIssueProjectName(_IssueProjectName.String)
		}
		if _Users.Valid {
			t.SetUsers(_Users.String)
		}
		if _InitialStartDate.Valid {
			t.SetInitialStartDate(_InitialStartDate.Int64)
		}
		if _TotalDuration.Valid {
			t.SetTotalDuration(_TotalDuration.Int64)
		}
		if _CreatedAt.Valid {
			t.SetCreatedAt(_CreatedAt.Int64)
		}
		if _ClosedAt.Valid {
			t.SetClosedAt(_ClosedAt.Int64)
		}
		if _PlannedEndDate.Valid {
			t.SetPlannedEndDate(_PlannedEndDate.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefType.Valid {
			t.SetRefType(_RefType.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		if _CustomFieldIdsVirtual.Valid {
			t.SetCustomFieldIdsVirtual(_CustomFieldIdsVirtual.String)
		}
		if _ReleaseDuration.Valid {
			t.SetReleaseDuration(_ReleaseDuration.Int64)
		}
		if _Completed.Valid {
			t.SetCompleted(_Completed.Bool)
		}
		if _CompletedDate.Valid {
			t.SetCompletedDate(_CompletedDate.Int64)
		}
		results = append(results, t)
	}
	return results, nil
}

// DBFind will find a IssueSummary record in the database with the provided parameters
func (t *IssueSummary) DBFind(ctx context.Context, db DB, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("issue_id"),
		orm.Column("total_issues"),
		orm.Column("new30_days"),
		orm.Column("total_closed"),
		orm.Column("closed30_days"),
		orm.Column("estimated_work_months"),
		orm.Column("estimated_work_months30_days"),
		orm.Column("title"),
		orm.Column("url"),
		orm.Column("priority"),
		orm.Column("priority_id"),
		orm.Column("status"),
		orm.Column("status_id"),
		orm.Column("issue_type"),
		orm.Column("issue_type_id"),
		orm.Column("resolution"),
		orm.Column("resolution_id"),
		orm.Column("state"),
		orm.Column("custom_field_ids"),
		orm.Column("teams"),
		orm.Column("parent_issue_id"),
		orm.Column("parents_issue_ids"),
		orm.Column("metadata"),
		orm.Column("project_id"),
		orm.Column("sprints"),
		orm.Column("labels"),
		orm.Column("top_level"),
		orm.Column("is_leaf"),
		orm.Column("path"),
		orm.Column("in_progress_duration"),
		orm.Column("verification_duration"),
		orm.Column("in_progress_count"),
		orm.Column("reopen_count"),
		orm.Column("mapped_type"),
		orm.Column("strategic_parent_id"),
		orm.Column("sprint_id"),
		orm.Column("issue_project_name"),
		orm.Column("users"),
		orm.Column("initial_start_date"),
		orm.Column("total_duration"),
		orm.Column("created_at"),
		orm.Column("closed_at"),
		orm.Column("planned_end_date"),
		orm.Column("customer_id"),
		orm.Column("ref_type"),
		orm.Column("ref_id"),
		orm.Column("custom_field_ids_virtual"),
		orm.Column("release_duration"),
		orm.Column("completed"),
		orm.Column("completed_date"),
		orm.Table(IssueSummaryTableName),
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
	var _TotalIssues sql.NullInt64
	var _New30Days sql.NullInt64
	var _TotalClosed sql.NullInt64
	var _Closed30Days sql.NullInt64
	var _EstimatedWorkMonths sql.NullFloat64
	var _EstimatedWorkMonths30Days sql.NullFloat64
	var _Title sql.NullString
	var _URL sql.NullString
	var _Priority sql.NullString
	var _PriorityID sql.NullString
	var _Status sql.NullString
	var _StatusID sql.NullString
	var _IssueType sql.NullString
	var _IssueTypeID sql.NullString
	var _Resolution sql.NullString
	var _ResolutionID sql.NullString
	var _State sql.NullString
	var _CustomFieldIds sql.NullString
	var _Teams sql.NullString
	var _ParentIssueID sql.NullString
	var _ParentsIssueIds sql.NullString
	var _Metadata sql.NullString
	var _ProjectID sql.NullString
	var _Sprints sql.NullString
	var _Labels sql.NullString
	var _TopLevel sql.NullBool
	var _IsLeaf sql.NullBool
	var _Path sql.NullString
	var _InProgressDuration sql.NullInt64
	var _VerificationDuration sql.NullInt64
	var _InProgressCount sql.NullInt64
	var _ReopenCount sql.NullInt64
	var _MappedType sql.NullString
	var _StrategicParentID sql.NullString
	var _SprintID sql.NullString
	var _IssueProjectName sql.NullString
	var _Users sql.NullString
	var _InitialStartDate sql.NullInt64
	var _TotalDuration sql.NullInt64
	var _CreatedAt sql.NullInt64
	var _ClosedAt sql.NullInt64
	var _PlannedEndDate sql.NullInt64
	var _CustomerID sql.NullString
	var _RefType sql.NullString
	var _RefID sql.NullString
	var _CustomFieldIdsVirtual sql.NullString
	var _ReleaseDuration sql.NullInt64
	var _Completed sql.NullBool
	var _CompletedDate sql.NullInt64
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_IssueID,
		&_TotalIssues,
		&_New30Days,
		&_TotalClosed,
		&_Closed30Days,
		&_EstimatedWorkMonths,
		&_EstimatedWorkMonths30Days,
		&_Title,
		&_URL,
		&_Priority,
		&_PriorityID,
		&_Status,
		&_StatusID,
		&_IssueType,
		&_IssueTypeID,
		&_Resolution,
		&_ResolutionID,
		&_State,
		&_CustomFieldIds,
		&_Teams,
		&_ParentIssueID,
		&_ParentsIssueIds,
		&_Metadata,
		&_ProjectID,
		&_Sprints,
		&_Labels,
		&_TopLevel,
		&_IsLeaf,
		&_Path,
		&_InProgressDuration,
		&_VerificationDuration,
		&_InProgressCount,
		&_ReopenCount,
		&_MappedType,
		&_StrategicParentID,
		&_SprintID,
		&_IssueProjectName,
		&_Users,
		&_InitialStartDate,
		&_TotalDuration,
		&_CreatedAt,
		&_ClosedAt,
		&_PlannedEndDate,
		&_CustomerID,
		&_RefType,
		&_RefID,
		&_CustomFieldIdsVirtual,
		&_ReleaseDuration,
		&_Completed,
		&_CompletedDate,
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
	if _TotalIssues.Valid {
		t.SetTotalIssues(int32(_TotalIssues.Int64))
	}
	if _New30Days.Valid {
		t.SetNew30Days(int32(_New30Days.Int64))
	}
	if _TotalClosed.Valid {
		t.SetTotalClosed(int32(_TotalClosed.Int64))
	}
	if _Closed30Days.Valid {
		t.SetClosed30Days(int32(_Closed30Days.Int64))
	}
	if _EstimatedWorkMonths.Valid {
		t.SetEstimatedWorkMonths(_EstimatedWorkMonths.Float64)
	}
	if _EstimatedWorkMonths30Days.Valid {
		t.SetEstimatedWorkMonths30Days(_EstimatedWorkMonths30Days.Float64)
	}
	if _Title.Valid {
		t.SetTitle(_Title.String)
	}
	if _URL.Valid {
		t.SetURL(_URL.String)
	}
	if _Priority.Valid {
		t.SetPriority(_Priority.String)
	}
	if _PriorityID.Valid {
		t.SetPriorityID(_PriorityID.String)
	}
	if _Status.Valid {
		t.SetStatus(_Status.String)
	}
	if _StatusID.Valid {
		t.SetStatusID(_StatusID.String)
	}
	if _IssueType.Valid {
		t.SetIssueType(_IssueType.String)
	}
	if _IssueTypeID.Valid {
		t.SetIssueTypeID(_IssueTypeID.String)
	}
	if _Resolution.Valid {
		t.SetResolution(_Resolution.String)
	}
	if _ResolutionID.Valid {
		t.SetResolutionID(_ResolutionID.String)
	}
	if _State.Valid {
		t.SetState(_State.String)
	}
	if _CustomFieldIds.Valid {
		t.SetCustomFieldIds(_CustomFieldIds.String)
	}
	if _Teams.Valid {
		t.SetTeams(_Teams.String)
	}
	if _ParentIssueID.Valid {
		t.SetParentIssueID(_ParentIssueID.String)
	}
	if _ParentsIssueIds.Valid {
		t.SetParentsIssueIds(_ParentsIssueIds.String)
	}
	if _Metadata.Valid {
		t.SetMetadata(_Metadata.String)
	}
	if _ProjectID.Valid {
		t.SetProjectID(_ProjectID.String)
	}
	if _Sprints.Valid {
		t.SetSprints(_Sprints.String)
	}
	if _Labels.Valid {
		t.SetLabels(_Labels.String)
	}
	if _TopLevel.Valid {
		t.SetTopLevel(_TopLevel.Bool)
	}
	if _IsLeaf.Valid {
		t.SetIsLeaf(_IsLeaf.Bool)
	}
	if _Path.Valid {
		t.SetPath(_Path.String)
	}
	if _InProgressDuration.Valid {
		t.SetInProgressDuration(_InProgressDuration.Int64)
	}
	if _VerificationDuration.Valid {
		t.SetVerificationDuration(_VerificationDuration.Int64)
	}
	if _InProgressCount.Valid {
		t.SetInProgressCount(int32(_InProgressCount.Int64))
	}
	if _ReopenCount.Valid {
		t.SetReopenCount(int32(_ReopenCount.Int64))
	}
	if _MappedType.Valid {
		t.SetMappedType(_MappedType.String)
	}
	if _StrategicParentID.Valid {
		t.SetStrategicParentID(_StrategicParentID.String)
	}
	if _SprintID.Valid {
		t.SetSprintID(_SprintID.String)
	}
	if _IssueProjectName.Valid {
		t.SetIssueProjectName(_IssueProjectName.String)
	}
	if _Users.Valid {
		t.SetUsers(_Users.String)
	}
	if _InitialStartDate.Valid {
		t.SetInitialStartDate(_InitialStartDate.Int64)
	}
	if _TotalDuration.Valid {
		t.SetTotalDuration(_TotalDuration.Int64)
	}
	if _CreatedAt.Valid {
		t.SetCreatedAt(_CreatedAt.Int64)
	}
	if _ClosedAt.Valid {
		t.SetClosedAt(_ClosedAt.Int64)
	}
	if _PlannedEndDate.Valid {
		t.SetPlannedEndDate(_PlannedEndDate.Int64)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _RefType.Valid {
		t.SetRefType(_RefType.String)
	}
	if _RefID.Valid {
		t.SetRefID(_RefID.String)
	}
	if _CustomFieldIdsVirtual.Valid {
		t.SetCustomFieldIdsVirtual(_CustomFieldIdsVirtual.String)
	}
	if _ReleaseDuration.Valid {
		t.SetReleaseDuration(_ReleaseDuration.Int64)
	}
	if _Completed.Valid {
		t.SetCompleted(_Completed.Bool)
	}
	if _CompletedDate.Valid {
		t.SetCompletedDate(_CompletedDate.Int64)
	}
	return true, nil
}

// DBFindTx will find a IssueSummary record in the database with the provided parameters using the provided transaction
func (t *IssueSummary) DBFindTx(ctx context.Context, tx Tx, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("issue_id"),
		orm.Column("total_issues"),
		orm.Column("new30_days"),
		orm.Column("total_closed"),
		orm.Column("closed30_days"),
		orm.Column("estimated_work_months"),
		orm.Column("estimated_work_months30_days"),
		orm.Column("title"),
		orm.Column("url"),
		orm.Column("priority"),
		orm.Column("priority_id"),
		orm.Column("status"),
		orm.Column("status_id"),
		orm.Column("issue_type"),
		orm.Column("issue_type_id"),
		orm.Column("resolution"),
		orm.Column("resolution_id"),
		orm.Column("state"),
		orm.Column("custom_field_ids"),
		orm.Column("teams"),
		orm.Column("parent_issue_id"),
		orm.Column("parents_issue_ids"),
		orm.Column("metadata"),
		orm.Column("project_id"),
		orm.Column("sprints"),
		orm.Column("labels"),
		orm.Column("top_level"),
		orm.Column("is_leaf"),
		orm.Column("path"),
		orm.Column("in_progress_duration"),
		orm.Column("verification_duration"),
		orm.Column("in_progress_count"),
		orm.Column("reopen_count"),
		orm.Column("mapped_type"),
		orm.Column("strategic_parent_id"),
		orm.Column("sprint_id"),
		orm.Column("issue_project_name"),
		orm.Column("users"),
		orm.Column("initial_start_date"),
		orm.Column("total_duration"),
		orm.Column("created_at"),
		orm.Column("closed_at"),
		orm.Column("planned_end_date"),
		orm.Column("customer_id"),
		orm.Column("ref_type"),
		orm.Column("ref_id"),
		orm.Column("custom_field_ids_virtual"),
		orm.Column("release_duration"),
		orm.Column("completed"),
		orm.Column("completed_date"),
		orm.Table(IssueSummaryTableName),
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
	var _TotalIssues sql.NullInt64
	var _New30Days sql.NullInt64
	var _TotalClosed sql.NullInt64
	var _Closed30Days sql.NullInt64
	var _EstimatedWorkMonths sql.NullFloat64
	var _EstimatedWorkMonths30Days sql.NullFloat64
	var _Title sql.NullString
	var _URL sql.NullString
	var _Priority sql.NullString
	var _PriorityID sql.NullString
	var _Status sql.NullString
	var _StatusID sql.NullString
	var _IssueType sql.NullString
	var _IssueTypeID sql.NullString
	var _Resolution sql.NullString
	var _ResolutionID sql.NullString
	var _State sql.NullString
	var _CustomFieldIds sql.NullString
	var _Teams sql.NullString
	var _ParentIssueID sql.NullString
	var _ParentsIssueIds sql.NullString
	var _Metadata sql.NullString
	var _ProjectID sql.NullString
	var _Sprints sql.NullString
	var _Labels sql.NullString
	var _TopLevel sql.NullBool
	var _IsLeaf sql.NullBool
	var _Path sql.NullString
	var _InProgressDuration sql.NullInt64
	var _VerificationDuration sql.NullInt64
	var _InProgressCount sql.NullInt64
	var _ReopenCount sql.NullInt64
	var _MappedType sql.NullString
	var _StrategicParentID sql.NullString
	var _SprintID sql.NullString
	var _IssueProjectName sql.NullString
	var _Users sql.NullString
	var _InitialStartDate sql.NullInt64
	var _TotalDuration sql.NullInt64
	var _CreatedAt sql.NullInt64
	var _ClosedAt sql.NullInt64
	var _PlannedEndDate sql.NullInt64
	var _CustomerID sql.NullString
	var _RefType sql.NullString
	var _RefID sql.NullString
	var _CustomFieldIdsVirtual sql.NullString
	var _ReleaseDuration sql.NullInt64
	var _Completed sql.NullBool
	var _CompletedDate sql.NullInt64
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_IssueID,
		&_TotalIssues,
		&_New30Days,
		&_TotalClosed,
		&_Closed30Days,
		&_EstimatedWorkMonths,
		&_EstimatedWorkMonths30Days,
		&_Title,
		&_URL,
		&_Priority,
		&_PriorityID,
		&_Status,
		&_StatusID,
		&_IssueType,
		&_IssueTypeID,
		&_Resolution,
		&_ResolutionID,
		&_State,
		&_CustomFieldIds,
		&_Teams,
		&_ParentIssueID,
		&_ParentsIssueIds,
		&_Metadata,
		&_ProjectID,
		&_Sprints,
		&_Labels,
		&_TopLevel,
		&_IsLeaf,
		&_Path,
		&_InProgressDuration,
		&_VerificationDuration,
		&_InProgressCount,
		&_ReopenCount,
		&_MappedType,
		&_StrategicParentID,
		&_SprintID,
		&_IssueProjectName,
		&_Users,
		&_InitialStartDate,
		&_TotalDuration,
		&_CreatedAt,
		&_ClosedAt,
		&_PlannedEndDate,
		&_CustomerID,
		&_RefType,
		&_RefID,
		&_CustomFieldIdsVirtual,
		&_ReleaseDuration,
		&_Completed,
		&_CompletedDate,
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
	if _TotalIssues.Valid {
		t.SetTotalIssues(int32(_TotalIssues.Int64))
	}
	if _New30Days.Valid {
		t.SetNew30Days(int32(_New30Days.Int64))
	}
	if _TotalClosed.Valid {
		t.SetTotalClosed(int32(_TotalClosed.Int64))
	}
	if _Closed30Days.Valid {
		t.SetClosed30Days(int32(_Closed30Days.Int64))
	}
	if _EstimatedWorkMonths.Valid {
		t.SetEstimatedWorkMonths(_EstimatedWorkMonths.Float64)
	}
	if _EstimatedWorkMonths30Days.Valid {
		t.SetEstimatedWorkMonths30Days(_EstimatedWorkMonths30Days.Float64)
	}
	if _Title.Valid {
		t.SetTitle(_Title.String)
	}
	if _URL.Valid {
		t.SetURL(_URL.String)
	}
	if _Priority.Valid {
		t.SetPriority(_Priority.String)
	}
	if _PriorityID.Valid {
		t.SetPriorityID(_PriorityID.String)
	}
	if _Status.Valid {
		t.SetStatus(_Status.String)
	}
	if _StatusID.Valid {
		t.SetStatusID(_StatusID.String)
	}
	if _IssueType.Valid {
		t.SetIssueType(_IssueType.String)
	}
	if _IssueTypeID.Valid {
		t.SetIssueTypeID(_IssueTypeID.String)
	}
	if _Resolution.Valid {
		t.SetResolution(_Resolution.String)
	}
	if _ResolutionID.Valid {
		t.SetResolutionID(_ResolutionID.String)
	}
	if _State.Valid {
		t.SetState(_State.String)
	}
	if _CustomFieldIds.Valid {
		t.SetCustomFieldIds(_CustomFieldIds.String)
	}
	if _Teams.Valid {
		t.SetTeams(_Teams.String)
	}
	if _ParentIssueID.Valid {
		t.SetParentIssueID(_ParentIssueID.String)
	}
	if _ParentsIssueIds.Valid {
		t.SetParentsIssueIds(_ParentsIssueIds.String)
	}
	if _Metadata.Valid {
		t.SetMetadata(_Metadata.String)
	}
	if _ProjectID.Valid {
		t.SetProjectID(_ProjectID.String)
	}
	if _Sprints.Valid {
		t.SetSprints(_Sprints.String)
	}
	if _Labels.Valid {
		t.SetLabels(_Labels.String)
	}
	if _TopLevel.Valid {
		t.SetTopLevel(_TopLevel.Bool)
	}
	if _IsLeaf.Valid {
		t.SetIsLeaf(_IsLeaf.Bool)
	}
	if _Path.Valid {
		t.SetPath(_Path.String)
	}
	if _InProgressDuration.Valid {
		t.SetInProgressDuration(_InProgressDuration.Int64)
	}
	if _VerificationDuration.Valid {
		t.SetVerificationDuration(_VerificationDuration.Int64)
	}
	if _InProgressCount.Valid {
		t.SetInProgressCount(int32(_InProgressCount.Int64))
	}
	if _ReopenCount.Valid {
		t.SetReopenCount(int32(_ReopenCount.Int64))
	}
	if _MappedType.Valid {
		t.SetMappedType(_MappedType.String)
	}
	if _StrategicParentID.Valid {
		t.SetStrategicParentID(_StrategicParentID.String)
	}
	if _SprintID.Valid {
		t.SetSprintID(_SprintID.String)
	}
	if _IssueProjectName.Valid {
		t.SetIssueProjectName(_IssueProjectName.String)
	}
	if _Users.Valid {
		t.SetUsers(_Users.String)
	}
	if _InitialStartDate.Valid {
		t.SetInitialStartDate(_InitialStartDate.Int64)
	}
	if _TotalDuration.Valid {
		t.SetTotalDuration(_TotalDuration.Int64)
	}
	if _CreatedAt.Valid {
		t.SetCreatedAt(_CreatedAt.Int64)
	}
	if _ClosedAt.Valid {
		t.SetClosedAt(_ClosedAt.Int64)
	}
	if _PlannedEndDate.Valid {
		t.SetPlannedEndDate(_PlannedEndDate.Int64)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _RefType.Valid {
		t.SetRefType(_RefType.String)
	}
	if _RefID.Valid {
		t.SetRefID(_RefID.String)
	}
	if _CustomFieldIdsVirtual.Valid {
		t.SetCustomFieldIdsVirtual(_CustomFieldIdsVirtual.String)
	}
	if _ReleaseDuration.Valid {
		t.SetReleaseDuration(_ReleaseDuration.Int64)
	}
	if _Completed.Valid {
		t.SetCompleted(_Completed.Bool)
	}
	if _CompletedDate.Valid {
		t.SetCompletedDate(_CompletedDate.Int64)
	}
	return true, nil
}

// CountIssueSummaries will find the count of IssueSummary records in the database
func CountIssueSummaries(ctx context.Context, db DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(IssueSummaryTableName),
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

// CountIssueSummariesTx will find the count of IssueSummary records in the database using the provided transaction
func CountIssueSummariesTx(ctx context.Context, tx Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(IssueSummaryTableName),
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

// DBCount will find the count of IssueSummary records in the database
func (t *IssueSummary) DBCount(ctx context.Context, db DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(IssueSummaryTableName),
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

// DBCountTx will find the count of IssueSummary records in the database using the provided transaction
func (t *IssueSummary) DBCountTx(ctx context.Context, tx Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(IssueSummaryTableName),
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

// DBExists will return true if the IssueSummary record exists in the database
func (t *IssueSummary) DBExists(ctx context.Context, db DB) (bool, error) {
	q := "SELECT `id` FROM `issue_summary` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := db.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// DBExistsTx will return true if the IssueSummary record exists in the database using the provided transaction
func (t *IssueSummary) DBExistsTx(ctx context.Context, tx Tx) (bool, error) {
	q := "SELECT `id` FROM `issue_summary` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := tx.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// PrimaryKeyColumn returns the column name for the primary key
func (t *IssueSummary) PrimaryKeyColumn() string {
	return IssueSummaryColumnID
}

// PrimaryKeyColumnType returns the primary key column Go type as a string
func (t *IssueSummary) PrimaryKeyColumnType() string {
	return "string"
}

// PrimaryKey returns the primary key column value
func (t *IssueSummary) PrimaryKey() interface{} {
	return t.ID
}
