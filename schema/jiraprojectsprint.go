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

var _ Model = (*JiraProjectSprint)(nil)
var _ CSVWriter = (*JiraProjectSprint)(nil)
var _ JSONWriter = (*JiraProjectSprint)(nil)
var _ Checksum = (*JiraProjectSprint)(nil)

// JiraProjectSprintTableName is the name of the table in SQL
const JiraProjectSprintTableName = "jira_project_sprint"

var JiraProjectSprintColumns = []string{
	"id",
	"checksum",
	"project_id",
	"sprint_id",
	"name",
	"customer_id",
	"complete_date",
	"end_date",
	"start_date",
	"state",
	"goal",
	"initial_issue_ids",
	"final_issue_ids",
	"final_closed_issue_ids",
	"initial_issue_count",
	"final_issue_count",
	"closed_issue_count",
	"added_issue_count",
	"removed_issue_count",
	"initial_issues_closed",
	"initial_issues_in_final_count",
}

// JiraProjectSprint table
type JiraProjectSprint struct {
	AddedIssueCount           int32   `json:"added_issue_count"`
	Checksum                  *string `json:"checksum,omitempty"`
	ClosedIssueCount          int32   `json:"closed_issue_count"`
	CompleteDate              *int64  `json:"complete_date,omitempty"`
	CustomerID                string  `json:"customer_id"`
	EndDate                   *int64  `json:"end_date,omitempty"`
	FinalClosedIssueIds       *string `json:"final_closed_issue_ids,omitempty"`
	FinalIssueCount           int32   `json:"final_issue_count"`
	FinalIssueIds             *string `json:"final_issue_ids,omitempty"`
	Goal                      *string `json:"goal,omitempty"`
	ID                        string  `json:"id"`
	InitialIssueCount         int32   `json:"initial_issue_count"`
	InitialIssueIds           *string `json:"initial_issue_ids,omitempty"`
	InitialIssuesClosed       int32   `json:"initial_issues_closed"`
	InitialIssuesInFinalCount int32   `json:"initial_issues_in_final_count"`
	Name                      string  `json:"name"`
	ProjectID                 string  `json:"project_id"`
	RemovedIssueCount         int32   `json:"removed_issue_count"`
	SprintID                  string  `json:"sprint_id"`
	StartDate                 *int64  `json:"start_date,omitempty"`
	State                     *string `json:"state,omitempty"`
}

// TableName returns the SQL table name for JiraProjectSprint and satifies the Model interface
func (t *JiraProjectSprint) TableName() string {
	return JiraProjectSprintTableName
}

// ToCSV will serialize the JiraProjectSprint instance to a CSV compatible array of strings
func (t *JiraProjectSprint) ToCSV() []string {
	return []string{
		t.ID,
		t.CalculateChecksum(),
		t.ProjectID,
		t.SprintID,
		t.Name,
		t.CustomerID,
		toCSVString(t.CompleteDate),
		toCSVString(t.EndDate),
		toCSVString(t.StartDate),
		toCSVString(t.State),
		toCSVString(t.Goal),
		toCSVString(t.InitialIssueIds),
		toCSVString(t.FinalIssueIds),
		toCSVString(t.FinalClosedIssueIds),
		toCSVString(t.InitialIssueCount),
		toCSVString(t.FinalIssueCount),
		toCSVString(t.ClosedIssueCount),
		toCSVString(t.AddedIssueCount),
		toCSVString(t.RemovedIssueCount),
		toCSVString(t.InitialIssuesClosed),
		toCSVString(t.InitialIssuesInFinalCount),
	}
}

// WriteCSV will serialize the JiraProjectSprint instance to the writer as CSV and satisfies the CSVWriter interface
func (t *JiraProjectSprint) WriteCSV(w *csv.Writer) error {
	return w.Write(t.ToCSV())
}

// WriteJSON will serialize the JiraProjectSprint instance to the writer as JSON and satisfies the JSONWriter interface
func (t *JiraProjectSprint) WriteJSON(w io.Writer, indent ...bool) error {
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

// NewJiraProjectSprintReader creates a JSON reader which can read in JiraProjectSprint objects serialized as JSON either as an array, single object or json new lines
// and writes each JiraProjectSprint to the channel provided
func NewJiraProjectSprintReader(r io.Reader, ch chan<- JiraProjectSprint) error {
	return orm.Deserialize(r, func(buf json.RawMessage) error {
		dec := json.NewDecoder(bytes.NewBuffer(buf))
		e := JiraProjectSprint{}
		if err := dec.Decode(&e); err != nil {
			return err
		}
		ch <- e
		return nil
	})
}

// NewCSVJiraProjectSprintReaderDir will read the reader as CSV and emit each record to the channel provided
func NewCSVJiraProjectSprintReader(r io.Reader, ch chan<- JiraProjectSprint) error {
	cr := csv.NewReader(r)
	for {
		record, err := cr.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		ch <- JiraProjectSprint{
			ID:                        record[0],
			Checksum:                  fromStringPointer(record[1]),
			ProjectID:                 record[2],
			SprintID:                  record[3],
			Name:                      record[4],
			CustomerID:                record[5],
			CompleteDate:              fromCSVInt64Pointer(record[6]),
			EndDate:                   fromCSVInt64Pointer(record[7]),
			StartDate:                 fromCSVInt64Pointer(record[8]),
			State:                     fromStringPointer(record[9]),
			Goal:                      fromStringPointer(record[10]),
			InitialIssueIds:           fromStringPointer(record[11]),
			FinalIssueIds:             fromStringPointer(record[12]),
			FinalClosedIssueIds:       fromStringPointer(record[13]),
			InitialIssueCount:         fromCSVInt32(record[14]),
			FinalIssueCount:           fromCSVInt32(record[15]),
			ClosedIssueCount:          fromCSVInt32(record[16]),
			AddedIssueCount:           fromCSVInt32(record[17]),
			RemovedIssueCount:         fromCSVInt32(record[18]),
			InitialIssuesClosed:       fromCSVInt32(record[19]),
			InitialIssuesInFinalCount: fromCSVInt32(record[20]),
		}
	}
	return nil
}

// NewCSVJiraProjectSprintReaderFile will read the file as a CSV and emit each record to the channel provided
func NewCSVJiraProjectSprintReaderFile(fp string, ch chan<- JiraProjectSprint) error {
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
	return NewCSVJiraProjectSprintReader(fc, ch)
}

// NewCSVJiraProjectSprintReaderDir will read the jira_project_sprint.csv.gz file as a CSV and emit each record to the channel provided
func NewCSVJiraProjectSprintReaderDir(dir string, ch chan<- JiraProjectSprint) error {
	return NewCSVJiraProjectSprintReaderFile(filepath.Join(dir, "jira_project_sprint.csv.gz"), ch)
}

// JiraProjectSprintCSVDeduper is a function callback which takes the existing value (a) and the new value (b)
// and the return value should be the one to use (or a new one, if applicable). return nil
// to skip processing of this record
type JiraProjectSprintCSVDeduper func(a JiraProjectSprint, b JiraProjectSprint) *JiraProjectSprint

// JiraProjectSprintCSVDedupeDisabled is set on whether the CSV writer should de-dupe values by key
var JiraProjectSprintCSVDedupeDisabled bool

// NewJiraProjectSprintCSVWriterSize creates a batch writer that will write each JiraProjectSprint into a CSV file
// this method will automatically de-duplicate entries using the primary key. if the checksum
// for a newer item with the same primary key doesn't match a previously sent item, the newer
// one will be used
func NewJiraProjectSprintCSVWriterSize(w io.Writer, size int, dedupers ...JiraProjectSprintCSVDeduper) (chan JiraProjectSprint, chan bool, error) {
	cw := csv.NewWriter(w)
	ch := make(chan JiraProjectSprint, size)
	done := make(chan bool)
	go func() {
		defer func() { done <- true }()
		dodedupe := !JiraProjectSprintCSVDedupeDisabled
		var kv map[string]*JiraProjectSprint
		var deduper JiraProjectSprintCSVDeduper
		if dedupers != nil && len(dedupers) > 0 {
			deduper = dedupers[0]
			dodedupe = true
		}
		if dodedupe {
			kv = make(map[string]*JiraProjectSprint)
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

// JiraProjectSprintCSVDefaultSize is the default channel buffer size if not provided
var JiraProjectSprintCSVDefaultSize = 100

// NewJiraProjectSprintCSVWriter creates a batch writer that will write each JiraProjectSprint into a CSV file
func NewJiraProjectSprintCSVWriter(w io.Writer, dedupers ...JiraProjectSprintCSVDeduper) (chan JiraProjectSprint, chan bool, error) {
	return NewJiraProjectSprintCSVWriterSize(w, JiraProjectSprintCSVDefaultSize, dedupers...)
}

// NewJiraProjectSprintCSVWriterDir creates a batch writer that will write each JiraProjectSprint into a CSV file named jira_project_sprint.csv.gz in dir
func NewJiraProjectSprintCSVWriterDir(dir string, dedupers ...JiraProjectSprintCSVDeduper) (chan JiraProjectSprint, chan bool, error) {
	return NewJiraProjectSprintCSVWriterFile(filepath.Join(dir, "jira_project_sprint.csv.gz"), dedupers...)
}

// NewJiraProjectSprintCSVWriterFile creates a batch writer that will write each JiraProjectSprint into a CSV file
func NewJiraProjectSprintCSVWriterFile(fn string, dedupers ...JiraProjectSprintCSVDeduper) (chan JiraProjectSprint, chan bool, error) {
	f, err := os.Create(fn)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening CSV file %s. %v", fn, err)
	}
	var fc io.WriteCloser = f
	if filepath.Ext(fn) == ".gz" {
		w, _ := gzip.NewWriterLevel(f, gzip.BestCompression)
		fc = w
	}
	ch, done, err := NewJiraProjectSprintCSVWriter(fc, dedupers...)
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

type JiraProjectSprintDBAction func(ctx context.Context, db *sql.DB, record JiraProjectSprint) error

// NewJiraProjectSprintDBWriterSize creates a DB writer that will write each issue into the DB
func NewJiraProjectSprintDBWriterSize(ctx context.Context, db *sql.DB, errors chan<- error, size int, actions ...JiraProjectSprintDBAction) (chan JiraProjectSprint, chan bool, error) {
	ch := make(chan JiraProjectSprint, size)
	done := make(chan bool)
	var action JiraProjectSprintDBAction
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

// NewJiraProjectSprintDBWriter creates a DB writer that will write each issue into the DB
func NewJiraProjectSprintDBWriter(ctx context.Context, db *sql.DB, errors chan<- error, actions ...JiraProjectSprintDBAction) (chan JiraProjectSprint, chan bool, error) {
	return NewJiraProjectSprintDBWriterSize(ctx, db, errors, 100, actions...)
}

// JiraProjectSprintColumnID is the ID SQL column name for the JiraProjectSprint table
const JiraProjectSprintColumnID = "id"

// JiraProjectSprintEscapedColumnID is the escaped ID SQL column name for the JiraProjectSprint table
const JiraProjectSprintEscapedColumnID = "`id`"

// JiraProjectSprintColumnChecksum is the Checksum SQL column name for the JiraProjectSprint table
const JiraProjectSprintColumnChecksum = "checksum"

// JiraProjectSprintEscapedColumnChecksum is the escaped Checksum SQL column name for the JiraProjectSprint table
const JiraProjectSprintEscapedColumnChecksum = "`checksum`"

// JiraProjectSprintColumnProjectID is the ProjectID SQL column name for the JiraProjectSprint table
const JiraProjectSprintColumnProjectID = "project_id"

// JiraProjectSprintEscapedColumnProjectID is the escaped ProjectID SQL column name for the JiraProjectSprint table
const JiraProjectSprintEscapedColumnProjectID = "`project_id`"

// JiraProjectSprintColumnSprintID is the SprintID SQL column name for the JiraProjectSprint table
const JiraProjectSprintColumnSprintID = "sprint_id"

// JiraProjectSprintEscapedColumnSprintID is the escaped SprintID SQL column name for the JiraProjectSprint table
const JiraProjectSprintEscapedColumnSprintID = "`sprint_id`"

// JiraProjectSprintColumnName is the Name SQL column name for the JiraProjectSprint table
const JiraProjectSprintColumnName = "name"

// JiraProjectSprintEscapedColumnName is the escaped Name SQL column name for the JiraProjectSprint table
const JiraProjectSprintEscapedColumnName = "`name`"

// JiraProjectSprintColumnCustomerID is the CustomerID SQL column name for the JiraProjectSprint table
const JiraProjectSprintColumnCustomerID = "customer_id"

// JiraProjectSprintEscapedColumnCustomerID is the escaped CustomerID SQL column name for the JiraProjectSprint table
const JiraProjectSprintEscapedColumnCustomerID = "`customer_id`"

// JiraProjectSprintColumnCompleteDate is the CompleteDate SQL column name for the JiraProjectSprint table
const JiraProjectSprintColumnCompleteDate = "complete_date"

// JiraProjectSprintEscapedColumnCompleteDate is the escaped CompleteDate SQL column name for the JiraProjectSprint table
const JiraProjectSprintEscapedColumnCompleteDate = "`complete_date`"

// JiraProjectSprintColumnEndDate is the EndDate SQL column name for the JiraProjectSprint table
const JiraProjectSprintColumnEndDate = "end_date"

// JiraProjectSprintEscapedColumnEndDate is the escaped EndDate SQL column name for the JiraProjectSprint table
const JiraProjectSprintEscapedColumnEndDate = "`end_date`"

// JiraProjectSprintColumnStartDate is the StartDate SQL column name for the JiraProjectSprint table
const JiraProjectSprintColumnStartDate = "start_date"

// JiraProjectSprintEscapedColumnStartDate is the escaped StartDate SQL column name for the JiraProjectSprint table
const JiraProjectSprintEscapedColumnStartDate = "`start_date`"

// JiraProjectSprintColumnState is the State SQL column name for the JiraProjectSprint table
const JiraProjectSprintColumnState = "state"

// JiraProjectSprintEscapedColumnState is the escaped State SQL column name for the JiraProjectSprint table
const JiraProjectSprintEscapedColumnState = "`state`"

// JiraProjectSprintColumnGoal is the Goal SQL column name for the JiraProjectSprint table
const JiraProjectSprintColumnGoal = "goal"

// JiraProjectSprintEscapedColumnGoal is the escaped Goal SQL column name for the JiraProjectSprint table
const JiraProjectSprintEscapedColumnGoal = "`goal`"

// JiraProjectSprintColumnInitialIssueIds is the InitialIssueIds SQL column name for the JiraProjectSprint table
const JiraProjectSprintColumnInitialIssueIds = "initial_issue_ids"

// JiraProjectSprintEscapedColumnInitialIssueIds is the escaped InitialIssueIds SQL column name for the JiraProjectSprint table
const JiraProjectSprintEscapedColumnInitialIssueIds = "`initial_issue_ids`"

// JiraProjectSprintColumnFinalIssueIds is the FinalIssueIds SQL column name for the JiraProjectSprint table
const JiraProjectSprintColumnFinalIssueIds = "final_issue_ids"

// JiraProjectSprintEscapedColumnFinalIssueIds is the escaped FinalIssueIds SQL column name for the JiraProjectSprint table
const JiraProjectSprintEscapedColumnFinalIssueIds = "`final_issue_ids`"

// JiraProjectSprintColumnFinalClosedIssueIds is the FinalClosedIssueIds SQL column name for the JiraProjectSprint table
const JiraProjectSprintColumnFinalClosedIssueIds = "final_closed_issue_ids"

// JiraProjectSprintEscapedColumnFinalClosedIssueIds is the escaped FinalClosedIssueIds SQL column name for the JiraProjectSprint table
const JiraProjectSprintEscapedColumnFinalClosedIssueIds = "`final_closed_issue_ids`"

// JiraProjectSprintColumnInitialIssueCount is the InitialIssueCount SQL column name for the JiraProjectSprint table
const JiraProjectSprintColumnInitialIssueCount = "initial_issue_count"

// JiraProjectSprintEscapedColumnInitialIssueCount is the escaped InitialIssueCount SQL column name for the JiraProjectSprint table
const JiraProjectSprintEscapedColumnInitialIssueCount = "`initial_issue_count`"

// JiraProjectSprintColumnFinalIssueCount is the FinalIssueCount SQL column name for the JiraProjectSprint table
const JiraProjectSprintColumnFinalIssueCount = "final_issue_count"

// JiraProjectSprintEscapedColumnFinalIssueCount is the escaped FinalIssueCount SQL column name for the JiraProjectSprint table
const JiraProjectSprintEscapedColumnFinalIssueCount = "`final_issue_count`"

// JiraProjectSprintColumnClosedIssueCount is the ClosedIssueCount SQL column name for the JiraProjectSprint table
const JiraProjectSprintColumnClosedIssueCount = "closed_issue_count"

// JiraProjectSprintEscapedColumnClosedIssueCount is the escaped ClosedIssueCount SQL column name for the JiraProjectSprint table
const JiraProjectSprintEscapedColumnClosedIssueCount = "`closed_issue_count`"

// JiraProjectSprintColumnAddedIssueCount is the AddedIssueCount SQL column name for the JiraProjectSprint table
const JiraProjectSprintColumnAddedIssueCount = "added_issue_count"

// JiraProjectSprintEscapedColumnAddedIssueCount is the escaped AddedIssueCount SQL column name for the JiraProjectSprint table
const JiraProjectSprintEscapedColumnAddedIssueCount = "`added_issue_count`"

// JiraProjectSprintColumnRemovedIssueCount is the RemovedIssueCount SQL column name for the JiraProjectSprint table
const JiraProjectSprintColumnRemovedIssueCount = "removed_issue_count"

// JiraProjectSprintEscapedColumnRemovedIssueCount is the escaped RemovedIssueCount SQL column name for the JiraProjectSprint table
const JiraProjectSprintEscapedColumnRemovedIssueCount = "`removed_issue_count`"

// JiraProjectSprintColumnInitialIssuesClosed is the InitialIssuesClosed SQL column name for the JiraProjectSprint table
const JiraProjectSprintColumnInitialIssuesClosed = "initial_issues_closed"

// JiraProjectSprintEscapedColumnInitialIssuesClosed is the escaped InitialIssuesClosed SQL column name for the JiraProjectSprint table
const JiraProjectSprintEscapedColumnInitialIssuesClosed = "`initial_issues_closed`"

// JiraProjectSprintColumnInitialIssuesInFinalCount is the InitialIssuesInFinalCount SQL column name for the JiraProjectSprint table
const JiraProjectSprintColumnInitialIssuesInFinalCount = "initial_issues_in_final_count"

// JiraProjectSprintEscapedColumnInitialIssuesInFinalCount is the escaped InitialIssuesInFinalCount SQL column name for the JiraProjectSprint table
const JiraProjectSprintEscapedColumnInitialIssuesInFinalCount = "`initial_issues_in_final_count`"

// GetID will return the JiraProjectSprint ID value
func (t *JiraProjectSprint) GetID() string {
	return t.ID
}

// SetID will set the JiraProjectSprint ID value
func (t *JiraProjectSprint) SetID(v string) {
	t.ID = v
}

// FindJiraProjectSprintByID will find a JiraProjectSprint by ID
func FindJiraProjectSprintByID(ctx context.Context, db *sql.DB, value string) (*JiraProjectSprint, error) {
	q := "SELECT `jira_project_sprint`.`id`,`jira_project_sprint`.`checksum`,`jira_project_sprint`.`project_id`,`jira_project_sprint`.`sprint_id`,`jira_project_sprint`.`name`,`jira_project_sprint`.`customer_id`,`jira_project_sprint`.`complete_date`,`jira_project_sprint`.`end_date`,`jira_project_sprint`.`start_date`,`jira_project_sprint`.`state`,`jira_project_sprint`.`goal`,`jira_project_sprint`.`initial_issue_ids`,`jira_project_sprint`.`final_issue_ids`,`jira_project_sprint`.`final_closed_issue_ids`,`jira_project_sprint`.`initial_issue_count`,`jira_project_sprint`.`final_issue_count`,`jira_project_sprint`.`closed_issue_count`,`jira_project_sprint`.`added_issue_count`,`jira_project_sprint`.`removed_issue_count`,`jira_project_sprint`.`initial_issues_closed`,`jira_project_sprint`.`initial_issues_in_final_count` FROM `jira_project_sprint` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _ProjectID sql.NullString
	var _SprintID sql.NullString
	var _Name sql.NullString
	var _CustomerID sql.NullString
	var _CompleteDate sql.NullInt64
	var _EndDate sql.NullInt64
	var _StartDate sql.NullInt64
	var _State sql.NullString
	var _Goal sql.NullString
	var _InitialIssueIds sql.NullString
	var _FinalIssueIds sql.NullString
	var _FinalClosedIssueIds sql.NullString
	var _InitialIssueCount sql.NullInt64
	var _FinalIssueCount sql.NullInt64
	var _ClosedIssueCount sql.NullInt64
	var _AddedIssueCount sql.NullInt64
	var _RemovedIssueCount sql.NullInt64
	var _InitialIssuesClosed sql.NullInt64
	var _InitialIssuesInFinalCount sql.NullInt64
	err := db.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_ProjectID,
		&_SprintID,
		&_Name,
		&_CustomerID,
		&_CompleteDate,
		&_EndDate,
		&_StartDate,
		&_State,
		&_Goal,
		&_InitialIssueIds,
		&_FinalIssueIds,
		&_FinalClosedIssueIds,
		&_InitialIssueCount,
		&_FinalIssueCount,
		&_ClosedIssueCount,
		&_AddedIssueCount,
		&_RemovedIssueCount,
		&_InitialIssuesClosed,
		&_InitialIssuesInFinalCount,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &JiraProjectSprint{}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _Checksum.Valid {
		t.SetChecksum(_Checksum.String)
	}
	if _ProjectID.Valid {
		t.SetProjectID(_ProjectID.String)
	}
	if _SprintID.Valid {
		t.SetSprintID(_SprintID.String)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _CompleteDate.Valid {
		t.SetCompleteDate(_CompleteDate.Int64)
	}
	if _EndDate.Valid {
		t.SetEndDate(_EndDate.Int64)
	}
	if _StartDate.Valid {
		t.SetStartDate(_StartDate.Int64)
	}
	if _State.Valid {
		t.SetState(_State.String)
	}
	if _Goal.Valid {
		t.SetGoal(_Goal.String)
	}
	if _InitialIssueIds.Valid {
		t.SetInitialIssueIds(_InitialIssueIds.String)
	}
	if _FinalIssueIds.Valid {
		t.SetFinalIssueIds(_FinalIssueIds.String)
	}
	if _FinalClosedIssueIds.Valid {
		t.SetFinalClosedIssueIds(_FinalClosedIssueIds.String)
	}
	if _InitialIssueCount.Valid {
		t.SetInitialIssueCount(int32(_InitialIssueCount.Int64))
	}
	if _FinalIssueCount.Valid {
		t.SetFinalIssueCount(int32(_FinalIssueCount.Int64))
	}
	if _ClosedIssueCount.Valid {
		t.SetClosedIssueCount(int32(_ClosedIssueCount.Int64))
	}
	if _AddedIssueCount.Valid {
		t.SetAddedIssueCount(int32(_AddedIssueCount.Int64))
	}
	if _RemovedIssueCount.Valid {
		t.SetRemovedIssueCount(int32(_RemovedIssueCount.Int64))
	}
	if _InitialIssuesClosed.Valid {
		t.SetInitialIssuesClosed(int32(_InitialIssuesClosed.Int64))
	}
	if _InitialIssuesInFinalCount.Valid {
		t.SetInitialIssuesInFinalCount(int32(_InitialIssuesInFinalCount.Int64))
	}
	return t, nil
}

// FindJiraProjectSprintByIDTx will find a JiraProjectSprint by ID using the provided transaction
func FindJiraProjectSprintByIDTx(ctx context.Context, tx *sql.Tx, value string) (*JiraProjectSprint, error) {
	q := "SELECT `jira_project_sprint`.`id`,`jira_project_sprint`.`checksum`,`jira_project_sprint`.`project_id`,`jira_project_sprint`.`sprint_id`,`jira_project_sprint`.`name`,`jira_project_sprint`.`customer_id`,`jira_project_sprint`.`complete_date`,`jira_project_sprint`.`end_date`,`jira_project_sprint`.`start_date`,`jira_project_sprint`.`state`,`jira_project_sprint`.`goal`,`jira_project_sprint`.`initial_issue_ids`,`jira_project_sprint`.`final_issue_ids`,`jira_project_sprint`.`final_closed_issue_ids`,`jira_project_sprint`.`initial_issue_count`,`jira_project_sprint`.`final_issue_count`,`jira_project_sprint`.`closed_issue_count`,`jira_project_sprint`.`added_issue_count`,`jira_project_sprint`.`removed_issue_count`,`jira_project_sprint`.`initial_issues_closed`,`jira_project_sprint`.`initial_issues_in_final_count` FROM `jira_project_sprint` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _ProjectID sql.NullString
	var _SprintID sql.NullString
	var _Name sql.NullString
	var _CustomerID sql.NullString
	var _CompleteDate sql.NullInt64
	var _EndDate sql.NullInt64
	var _StartDate sql.NullInt64
	var _State sql.NullString
	var _Goal sql.NullString
	var _InitialIssueIds sql.NullString
	var _FinalIssueIds sql.NullString
	var _FinalClosedIssueIds sql.NullString
	var _InitialIssueCount sql.NullInt64
	var _FinalIssueCount sql.NullInt64
	var _ClosedIssueCount sql.NullInt64
	var _AddedIssueCount sql.NullInt64
	var _RemovedIssueCount sql.NullInt64
	var _InitialIssuesClosed sql.NullInt64
	var _InitialIssuesInFinalCount sql.NullInt64
	err := tx.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_ProjectID,
		&_SprintID,
		&_Name,
		&_CustomerID,
		&_CompleteDate,
		&_EndDate,
		&_StartDate,
		&_State,
		&_Goal,
		&_InitialIssueIds,
		&_FinalIssueIds,
		&_FinalClosedIssueIds,
		&_InitialIssueCount,
		&_FinalIssueCount,
		&_ClosedIssueCount,
		&_AddedIssueCount,
		&_RemovedIssueCount,
		&_InitialIssuesClosed,
		&_InitialIssuesInFinalCount,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &JiraProjectSprint{}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _Checksum.Valid {
		t.SetChecksum(_Checksum.String)
	}
	if _ProjectID.Valid {
		t.SetProjectID(_ProjectID.String)
	}
	if _SprintID.Valid {
		t.SetSprintID(_SprintID.String)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _CompleteDate.Valid {
		t.SetCompleteDate(_CompleteDate.Int64)
	}
	if _EndDate.Valid {
		t.SetEndDate(_EndDate.Int64)
	}
	if _StartDate.Valid {
		t.SetStartDate(_StartDate.Int64)
	}
	if _State.Valid {
		t.SetState(_State.String)
	}
	if _Goal.Valid {
		t.SetGoal(_Goal.String)
	}
	if _InitialIssueIds.Valid {
		t.SetInitialIssueIds(_InitialIssueIds.String)
	}
	if _FinalIssueIds.Valid {
		t.SetFinalIssueIds(_FinalIssueIds.String)
	}
	if _FinalClosedIssueIds.Valid {
		t.SetFinalClosedIssueIds(_FinalClosedIssueIds.String)
	}
	if _InitialIssueCount.Valid {
		t.SetInitialIssueCount(int32(_InitialIssueCount.Int64))
	}
	if _FinalIssueCount.Valid {
		t.SetFinalIssueCount(int32(_FinalIssueCount.Int64))
	}
	if _ClosedIssueCount.Valid {
		t.SetClosedIssueCount(int32(_ClosedIssueCount.Int64))
	}
	if _AddedIssueCount.Valid {
		t.SetAddedIssueCount(int32(_AddedIssueCount.Int64))
	}
	if _RemovedIssueCount.Valid {
		t.SetRemovedIssueCount(int32(_RemovedIssueCount.Int64))
	}
	if _InitialIssuesClosed.Valid {
		t.SetInitialIssuesClosed(int32(_InitialIssuesClosed.Int64))
	}
	if _InitialIssuesInFinalCount.Valid {
		t.SetInitialIssuesInFinalCount(int32(_InitialIssuesInFinalCount.Int64))
	}
	return t, nil
}

// GetChecksum will return the JiraProjectSprint Checksum value
func (t *JiraProjectSprint) GetChecksum() string {
	if t.Checksum == nil {
		return ""
	}
	return *t.Checksum
}

// SetChecksum will set the JiraProjectSprint Checksum value
func (t *JiraProjectSprint) SetChecksum(v string) {
	t.Checksum = &v
}

// GetProjectID will return the JiraProjectSprint ProjectID value
func (t *JiraProjectSprint) GetProjectID() string {
	return t.ProjectID
}

// SetProjectID will set the JiraProjectSprint ProjectID value
func (t *JiraProjectSprint) SetProjectID(v string) {
	t.ProjectID = v
}

// FindJiraProjectSprintsByProjectID will find all JiraProjectSprints by the ProjectID value
func FindJiraProjectSprintsByProjectID(ctx context.Context, db *sql.DB, value string) ([]*JiraProjectSprint, error) {
	q := "SELECT `jira_project_sprint`.`id`,`jira_project_sprint`.`checksum`,`jira_project_sprint`.`project_id`,`jira_project_sprint`.`sprint_id`,`jira_project_sprint`.`name`,`jira_project_sprint`.`customer_id`,`jira_project_sprint`.`complete_date`,`jira_project_sprint`.`end_date`,`jira_project_sprint`.`start_date`,`jira_project_sprint`.`state`,`jira_project_sprint`.`goal`,`jira_project_sprint`.`initial_issue_ids`,`jira_project_sprint`.`final_issue_ids`,`jira_project_sprint`.`final_closed_issue_ids`,`jira_project_sprint`.`initial_issue_count`,`jira_project_sprint`.`final_issue_count`,`jira_project_sprint`.`closed_issue_count`,`jira_project_sprint`.`added_issue_count`,`jira_project_sprint`.`removed_issue_count`,`jira_project_sprint`.`initial_issues_closed`,`jira_project_sprint`.`initial_issues_in_final_count` FROM `jira_project_sprint` WHERE `project_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraProjectSprint, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _ProjectID sql.NullString
		var _SprintID sql.NullString
		var _Name sql.NullString
		var _CustomerID sql.NullString
		var _CompleteDate sql.NullInt64
		var _EndDate sql.NullInt64
		var _StartDate sql.NullInt64
		var _State sql.NullString
		var _Goal sql.NullString
		var _InitialIssueIds sql.NullString
		var _FinalIssueIds sql.NullString
		var _FinalClosedIssueIds sql.NullString
		var _InitialIssueCount sql.NullInt64
		var _FinalIssueCount sql.NullInt64
		var _ClosedIssueCount sql.NullInt64
		var _AddedIssueCount sql.NullInt64
		var _RemovedIssueCount sql.NullInt64
		var _InitialIssuesClosed sql.NullInt64
		var _InitialIssuesInFinalCount sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_ProjectID,
			&_SprintID,
			&_Name,
			&_CustomerID,
			&_CompleteDate,
			&_EndDate,
			&_StartDate,
			&_State,
			&_Goal,
			&_InitialIssueIds,
			&_FinalIssueIds,
			&_FinalClosedIssueIds,
			&_InitialIssueCount,
			&_FinalIssueCount,
			&_ClosedIssueCount,
			&_AddedIssueCount,
			&_RemovedIssueCount,
			&_InitialIssuesClosed,
			&_InitialIssuesInFinalCount,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraProjectSprint{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _SprintID.Valid {
			t.SetSprintID(_SprintID.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _CompleteDate.Valid {
			t.SetCompleteDate(_CompleteDate.Int64)
		}
		if _EndDate.Valid {
			t.SetEndDate(_EndDate.Int64)
		}
		if _StartDate.Valid {
			t.SetStartDate(_StartDate.Int64)
		}
		if _State.Valid {
			t.SetState(_State.String)
		}
		if _Goal.Valid {
			t.SetGoal(_Goal.String)
		}
		if _InitialIssueIds.Valid {
			t.SetInitialIssueIds(_InitialIssueIds.String)
		}
		if _FinalIssueIds.Valid {
			t.SetFinalIssueIds(_FinalIssueIds.String)
		}
		if _FinalClosedIssueIds.Valid {
			t.SetFinalClosedIssueIds(_FinalClosedIssueIds.String)
		}
		if _InitialIssueCount.Valid {
			t.SetInitialIssueCount(int32(_InitialIssueCount.Int64))
		}
		if _FinalIssueCount.Valid {
			t.SetFinalIssueCount(int32(_FinalIssueCount.Int64))
		}
		if _ClosedIssueCount.Valid {
			t.SetClosedIssueCount(int32(_ClosedIssueCount.Int64))
		}
		if _AddedIssueCount.Valid {
			t.SetAddedIssueCount(int32(_AddedIssueCount.Int64))
		}
		if _RemovedIssueCount.Valid {
			t.SetRemovedIssueCount(int32(_RemovedIssueCount.Int64))
		}
		if _InitialIssuesClosed.Valid {
			t.SetInitialIssuesClosed(int32(_InitialIssuesClosed.Int64))
		}
		if _InitialIssuesInFinalCount.Valid {
			t.SetInitialIssuesInFinalCount(int32(_InitialIssuesInFinalCount.Int64))
		}
		results = append(results, t)
	}
	return results, nil
}

// FindJiraProjectSprintsByProjectIDTx will find all JiraProjectSprints by the ProjectID value using the provided transaction
func FindJiraProjectSprintsByProjectIDTx(ctx context.Context, tx *sql.Tx, value string) ([]*JiraProjectSprint, error) {
	q := "SELECT `jira_project_sprint`.`id`,`jira_project_sprint`.`checksum`,`jira_project_sprint`.`project_id`,`jira_project_sprint`.`sprint_id`,`jira_project_sprint`.`name`,`jira_project_sprint`.`customer_id`,`jira_project_sprint`.`complete_date`,`jira_project_sprint`.`end_date`,`jira_project_sprint`.`start_date`,`jira_project_sprint`.`state`,`jira_project_sprint`.`goal`,`jira_project_sprint`.`initial_issue_ids`,`jira_project_sprint`.`final_issue_ids`,`jira_project_sprint`.`final_closed_issue_ids`,`jira_project_sprint`.`initial_issue_count`,`jira_project_sprint`.`final_issue_count`,`jira_project_sprint`.`closed_issue_count`,`jira_project_sprint`.`added_issue_count`,`jira_project_sprint`.`removed_issue_count`,`jira_project_sprint`.`initial_issues_closed`,`jira_project_sprint`.`initial_issues_in_final_count` FROM `jira_project_sprint` WHERE `project_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraProjectSprint, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _ProjectID sql.NullString
		var _SprintID sql.NullString
		var _Name sql.NullString
		var _CustomerID sql.NullString
		var _CompleteDate sql.NullInt64
		var _EndDate sql.NullInt64
		var _StartDate sql.NullInt64
		var _State sql.NullString
		var _Goal sql.NullString
		var _InitialIssueIds sql.NullString
		var _FinalIssueIds sql.NullString
		var _FinalClosedIssueIds sql.NullString
		var _InitialIssueCount sql.NullInt64
		var _FinalIssueCount sql.NullInt64
		var _ClosedIssueCount sql.NullInt64
		var _AddedIssueCount sql.NullInt64
		var _RemovedIssueCount sql.NullInt64
		var _InitialIssuesClosed sql.NullInt64
		var _InitialIssuesInFinalCount sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_ProjectID,
			&_SprintID,
			&_Name,
			&_CustomerID,
			&_CompleteDate,
			&_EndDate,
			&_StartDate,
			&_State,
			&_Goal,
			&_InitialIssueIds,
			&_FinalIssueIds,
			&_FinalClosedIssueIds,
			&_InitialIssueCount,
			&_FinalIssueCount,
			&_ClosedIssueCount,
			&_AddedIssueCount,
			&_RemovedIssueCount,
			&_InitialIssuesClosed,
			&_InitialIssuesInFinalCount,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraProjectSprint{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _SprintID.Valid {
			t.SetSprintID(_SprintID.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _CompleteDate.Valid {
			t.SetCompleteDate(_CompleteDate.Int64)
		}
		if _EndDate.Valid {
			t.SetEndDate(_EndDate.Int64)
		}
		if _StartDate.Valid {
			t.SetStartDate(_StartDate.Int64)
		}
		if _State.Valid {
			t.SetState(_State.String)
		}
		if _Goal.Valid {
			t.SetGoal(_Goal.String)
		}
		if _InitialIssueIds.Valid {
			t.SetInitialIssueIds(_InitialIssueIds.String)
		}
		if _FinalIssueIds.Valid {
			t.SetFinalIssueIds(_FinalIssueIds.String)
		}
		if _FinalClosedIssueIds.Valid {
			t.SetFinalClosedIssueIds(_FinalClosedIssueIds.String)
		}
		if _InitialIssueCount.Valid {
			t.SetInitialIssueCount(int32(_InitialIssueCount.Int64))
		}
		if _FinalIssueCount.Valid {
			t.SetFinalIssueCount(int32(_FinalIssueCount.Int64))
		}
		if _ClosedIssueCount.Valid {
			t.SetClosedIssueCount(int32(_ClosedIssueCount.Int64))
		}
		if _AddedIssueCount.Valid {
			t.SetAddedIssueCount(int32(_AddedIssueCount.Int64))
		}
		if _RemovedIssueCount.Valid {
			t.SetRemovedIssueCount(int32(_RemovedIssueCount.Int64))
		}
		if _InitialIssuesClosed.Valid {
			t.SetInitialIssuesClosed(int32(_InitialIssuesClosed.Int64))
		}
		if _InitialIssuesInFinalCount.Valid {
			t.SetInitialIssuesInFinalCount(int32(_InitialIssuesInFinalCount.Int64))
		}
		results = append(results, t)
	}
	return results, nil
}

// GetSprintID will return the JiraProjectSprint SprintID value
func (t *JiraProjectSprint) GetSprintID() string {
	return t.SprintID
}

// SetSprintID will set the JiraProjectSprint SprintID value
func (t *JiraProjectSprint) SetSprintID(v string) {
	t.SprintID = v
}

// GetName will return the JiraProjectSprint Name value
func (t *JiraProjectSprint) GetName() string {
	return t.Name
}

// SetName will set the JiraProjectSprint Name value
func (t *JiraProjectSprint) SetName(v string) {
	t.Name = v
}

// GetCustomerID will return the JiraProjectSprint CustomerID value
func (t *JiraProjectSprint) GetCustomerID() string {
	return t.CustomerID
}

// SetCustomerID will set the JiraProjectSprint CustomerID value
func (t *JiraProjectSprint) SetCustomerID(v string) {
	t.CustomerID = v
}

// FindJiraProjectSprintsByCustomerID will find all JiraProjectSprints by the CustomerID value
func FindJiraProjectSprintsByCustomerID(ctx context.Context, db *sql.DB, value string) ([]*JiraProjectSprint, error) {
	q := "SELECT `jira_project_sprint`.`id`,`jira_project_sprint`.`checksum`,`jira_project_sprint`.`project_id`,`jira_project_sprint`.`sprint_id`,`jira_project_sprint`.`name`,`jira_project_sprint`.`customer_id`,`jira_project_sprint`.`complete_date`,`jira_project_sprint`.`end_date`,`jira_project_sprint`.`start_date`,`jira_project_sprint`.`state`,`jira_project_sprint`.`goal`,`jira_project_sprint`.`initial_issue_ids`,`jira_project_sprint`.`final_issue_ids`,`jira_project_sprint`.`final_closed_issue_ids`,`jira_project_sprint`.`initial_issue_count`,`jira_project_sprint`.`final_issue_count`,`jira_project_sprint`.`closed_issue_count`,`jira_project_sprint`.`added_issue_count`,`jira_project_sprint`.`removed_issue_count`,`jira_project_sprint`.`initial_issues_closed`,`jira_project_sprint`.`initial_issues_in_final_count` FROM `jira_project_sprint` WHERE `customer_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraProjectSprint, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _ProjectID sql.NullString
		var _SprintID sql.NullString
		var _Name sql.NullString
		var _CustomerID sql.NullString
		var _CompleteDate sql.NullInt64
		var _EndDate sql.NullInt64
		var _StartDate sql.NullInt64
		var _State sql.NullString
		var _Goal sql.NullString
		var _InitialIssueIds sql.NullString
		var _FinalIssueIds sql.NullString
		var _FinalClosedIssueIds sql.NullString
		var _InitialIssueCount sql.NullInt64
		var _FinalIssueCount sql.NullInt64
		var _ClosedIssueCount sql.NullInt64
		var _AddedIssueCount sql.NullInt64
		var _RemovedIssueCount sql.NullInt64
		var _InitialIssuesClosed sql.NullInt64
		var _InitialIssuesInFinalCount sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_ProjectID,
			&_SprintID,
			&_Name,
			&_CustomerID,
			&_CompleteDate,
			&_EndDate,
			&_StartDate,
			&_State,
			&_Goal,
			&_InitialIssueIds,
			&_FinalIssueIds,
			&_FinalClosedIssueIds,
			&_InitialIssueCount,
			&_FinalIssueCount,
			&_ClosedIssueCount,
			&_AddedIssueCount,
			&_RemovedIssueCount,
			&_InitialIssuesClosed,
			&_InitialIssuesInFinalCount,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraProjectSprint{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _SprintID.Valid {
			t.SetSprintID(_SprintID.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _CompleteDate.Valid {
			t.SetCompleteDate(_CompleteDate.Int64)
		}
		if _EndDate.Valid {
			t.SetEndDate(_EndDate.Int64)
		}
		if _StartDate.Valid {
			t.SetStartDate(_StartDate.Int64)
		}
		if _State.Valid {
			t.SetState(_State.String)
		}
		if _Goal.Valid {
			t.SetGoal(_Goal.String)
		}
		if _InitialIssueIds.Valid {
			t.SetInitialIssueIds(_InitialIssueIds.String)
		}
		if _FinalIssueIds.Valid {
			t.SetFinalIssueIds(_FinalIssueIds.String)
		}
		if _FinalClosedIssueIds.Valid {
			t.SetFinalClosedIssueIds(_FinalClosedIssueIds.String)
		}
		if _InitialIssueCount.Valid {
			t.SetInitialIssueCount(int32(_InitialIssueCount.Int64))
		}
		if _FinalIssueCount.Valid {
			t.SetFinalIssueCount(int32(_FinalIssueCount.Int64))
		}
		if _ClosedIssueCount.Valid {
			t.SetClosedIssueCount(int32(_ClosedIssueCount.Int64))
		}
		if _AddedIssueCount.Valid {
			t.SetAddedIssueCount(int32(_AddedIssueCount.Int64))
		}
		if _RemovedIssueCount.Valid {
			t.SetRemovedIssueCount(int32(_RemovedIssueCount.Int64))
		}
		if _InitialIssuesClosed.Valid {
			t.SetInitialIssuesClosed(int32(_InitialIssuesClosed.Int64))
		}
		if _InitialIssuesInFinalCount.Valid {
			t.SetInitialIssuesInFinalCount(int32(_InitialIssuesInFinalCount.Int64))
		}
		results = append(results, t)
	}
	return results, nil
}

// FindJiraProjectSprintsByCustomerIDTx will find all JiraProjectSprints by the CustomerID value using the provided transaction
func FindJiraProjectSprintsByCustomerIDTx(ctx context.Context, tx *sql.Tx, value string) ([]*JiraProjectSprint, error) {
	q := "SELECT `jira_project_sprint`.`id`,`jira_project_sprint`.`checksum`,`jira_project_sprint`.`project_id`,`jira_project_sprint`.`sprint_id`,`jira_project_sprint`.`name`,`jira_project_sprint`.`customer_id`,`jira_project_sprint`.`complete_date`,`jira_project_sprint`.`end_date`,`jira_project_sprint`.`start_date`,`jira_project_sprint`.`state`,`jira_project_sprint`.`goal`,`jira_project_sprint`.`initial_issue_ids`,`jira_project_sprint`.`final_issue_ids`,`jira_project_sprint`.`final_closed_issue_ids`,`jira_project_sprint`.`initial_issue_count`,`jira_project_sprint`.`final_issue_count`,`jira_project_sprint`.`closed_issue_count`,`jira_project_sprint`.`added_issue_count`,`jira_project_sprint`.`removed_issue_count`,`jira_project_sprint`.`initial_issues_closed`,`jira_project_sprint`.`initial_issues_in_final_count` FROM `jira_project_sprint` WHERE `customer_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraProjectSprint, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _ProjectID sql.NullString
		var _SprintID sql.NullString
		var _Name sql.NullString
		var _CustomerID sql.NullString
		var _CompleteDate sql.NullInt64
		var _EndDate sql.NullInt64
		var _StartDate sql.NullInt64
		var _State sql.NullString
		var _Goal sql.NullString
		var _InitialIssueIds sql.NullString
		var _FinalIssueIds sql.NullString
		var _FinalClosedIssueIds sql.NullString
		var _InitialIssueCount sql.NullInt64
		var _FinalIssueCount sql.NullInt64
		var _ClosedIssueCount sql.NullInt64
		var _AddedIssueCount sql.NullInt64
		var _RemovedIssueCount sql.NullInt64
		var _InitialIssuesClosed sql.NullInt64
		var _InitialIssuesInFinalCount sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_ProjectID,
			&_SprintID,
			&_Name,
			&_CustomerID,
			&_CompleteDate,
			&_EndDate,
			&_StartDate,
			&_State,
			&_Goal,
			&_InitialIssueIds,
			&_FinalIssueIds,
			&_FinalClosedIssueIds,
			&_InitialIssueCount,
			&_FinalIssueCount,
			&_ClosedIssueCount,
			&_AddedIssueCount,
			&_RemovedIssueCount,
			&_InitialIssuesClosed,
			&_InitialIssuesInFinalCount,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraProjectSprint{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _SprintID.Valid {
			t.SetSprintID(_SprintID.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _CompleteDate.Valid {
			t.SetCompleteDate(_CompleteDate.Int64)
		}
		if _EndDate.Valid {
			t.SetEndDate(_EndDate.Int64)
		}
		if _StartDate.Valid {
			t.SetStartDate(_StartDate.Int64)
		}
		if _State.Valid {
			t.SetState(_State.String)
		}
		if _Goal.Valid {
			t.SetGoal(_Goal.String)
		}
		if _InitialIssueIds.Valid {
			t.SetInitialIssueIds(_InitialIssueIds.String)
		}
		if _FinalIssueIds.Valid {
			t.SetFinalIssueIds(_FinalIssueIds.String)
		}
		if _FinalClosedIssueIds.Valid {
			t.SetFinalClosedIssueIds(_FinalClosedIssueIds.String)
		}
		if _InitialIssueCount.Valid {
			t.SetInitialIssueCount(int32(_InitialIssueCount.Int64))
		}
		if _FinalIssueCount.Valid {
			t.SetFinalIssueCount(int32(_FinalIssueCount.Int64))
		}
		if _ClosedIssueCount.Valid {
			t.SetClosedIssueCount(int32(_ClosedIssueCount.Int64))
		}
		if _AddedIssueCount.Valid {
			t.SetAddedIssueCount(int32(_AddedIssueCount.Int64))
		}
		if _RemovedIssueCount.Valid {
			t.SetRemovedIssueCount(int32(_RemovedIssueCount.Int64))
		}
		if _InitialIssuesClosed.Valid {
			t.SetInitialIssuesClosed(int32(_InitialIssuesClosed.Int64))
		}
		if _InitialIssuesInFinalCount.Valid {
			t.SetInitialIssuesInFinalCount(int32(_InitialIssuesInFinalCount.Int64))
		}
		results = append(results, t)
	}
	return results, nil
}

// GetCompleteDate will return the JiraProjectSprint CompleteDate value
func (t *JiraProjectSprint) GetCompleteDate() int64 {
	if t.CompleteDate == nil {
		return int64(0)
	}
	return *t.CompleteDate
}

// SetCompleteDate will set the JiraProjectSprint CompleteDate value
func (t *JiraProjectSprint) SetCompleteDate(v int64) {
	t.CompleteDate = &v
}

// GetEndDate will return the JiraProjectSprint EndDate value
func (t *JiraProjectSprint) GetEndDate() int64 {
	if t.EndDate == nil {
		return int64(0)
	}
	return *t.EndDate
}

// SetEndDate will set the JiraProjectSprint EndDate value
func (t *JiraProjectSprint) SetEndDate(v int64) {
	t.EndDate = &v
}

// GetStartDate will return the JiraProjectSprint StartDate value
func (t *JiraProjectSprint) GetStartDate() int64 {
	if t.StartDate == nil {
		return int64(0)
	}
	return *t.StartDate
}

// SetStartDate will set the JiraProjectSprint StartDate value
func (t *JiraProjectSprint) SetStartDate(v int64) {
	t.StartDate = &v
}

// GetState will return the JiraProjectSprint State value
func (t *JiraProjectSprint) GetState() string {
	if t.State == nil {
		return ""
	}
	return *t.State
}

// SetState will set the JiraProjectSprint State value
func (t *JiraProjectSprint) SetState(v string) {
	t.State = &v
}

// GetGoal will return the JiraProjectSprint Goal value
func (t *JiraProjectSprint) GetGoal() string {
	if t.Goal == nil {
		return ""
	}
	return *t.Goal
}

// SetGoal will set the JiraProjectSprint Goal value
func (t *JiraProjectSprint) SetGoal(v string) {
	t.Goal = &v
}

// GetInitialIssueIds will return the JiraProjectSprint InitialIssueIds value
func (t *JiraProjectSprint) GetInitialIssueIds() string {
	if t.InitialIssueIds == nil {
		return ""
	}
	return *t.InitialIssueIds
}

// SetInitialIssueIds will set the JiraProjectSprint InitialIssueIds value
func (t *JiraProjectSprint) SetInitialIssueIds(v string) {
	t.InitialIssueIds = &v
}

// GetFinalIssueIds will return the JiraProjectSprint FinalIssueIds value
func (t *JiraProjectSprint) GetFinalIssueIds() string {
	if t.FinalIssueIds == nil {
		return ""
	}
	return *t.FinalIssueIds
}

// SetFinalIssueIds will set the JiraProjectSprint FinalIssueIds value
func (t *JiraProjectSprint) SetFinalIssueIds(v string) {
	t.FinalIssueIds = &v
}

// GetFinalClosedIssueIds will return the JiraProjectSprint FinalClosedIssueIds value
func (t *JiraProjectSprint) GetFinalClosedIssueIds() string {
	if t.FinalClosedIssueIds == nil {
		return ""
	}
	return *t.FinalClosedIssueIds
}

// SetFinalClosedIssueIds will set the JiraProjectSprint FinalClosedIssueIds value
func (t *JiraProjectSprint) SetFinalClosedIssueIds(v string) {
	t.FinalClosedIssueIds = &v
}

// GetInitialIssueCount will return the JiraProjectSprint InitialIssueCount value
func (t *JiraProjectSprint) GetInitialIssueCount() int32 {
	return t.InitialIssueCount
}

// SetInitialIssueCount will set the JiraProjectSprint InitialIssueCount value
func (t *JiraProjectSprint) SetInitialIssueCount(v int32) {
	t.InitialIssueCount = v
}

// GetFinalIssueCount will return the JiraProjectSprint FinalIssueCount value
func (t *JiraProjectSprint) GetFinalIssueCount() int32 {
	return t.FinalIssueCount
}

// SetFinalIssueCount will set the JiraProjectSprint FinalIssueCount value
func (t *JiraProjectSprint) SetFinalIssueCount(v int32) {
	t.FinalIssueCount = v
}

// GetClosedIssueCount will return the JiraProjectSprint ClosedIssueCount value
func (t *JiraProjectSprint) GetClosedIssueCount() int32 {
	return t.ClosedIssueCount
}

// SetClosedIssueCount will set the JiraProjectSprint ClosedIssueCount value
func (t *JiraProjectSprint) SetClosedIssueCount(v int32) {
	t.ClosedIssueCount = v
}

// GetAddedIssueCount will return the JiraProjectSprint AddedIssueCount value
func (t *JiraProjectSprint) GetAddedIssueCount() int32 {
	return t.AddedIssueCount
}

// SetAddedIssueCount will set the JiraProjectSprint AddedIssueCount value
func (t *JiraProjectSprint) SetAddedIssueCount(v int32) {
	t.AddedIssueCount = v
}

// GetRemovedIssueCount will return the JiraProjectSprint RemovedIssueCount value
func (t *JiraProjectSprint) GetRemovedIssueCount() int32 {
	return t.RemovedIssueCount
}

// SetRemovedIssueCount will set the JiraProjectSprint RemovedIssueCount value
func (t *JiraProjectSprint) SetRemovedIssueCount(v int32) {
	t.RemovedIssueCount = v
}

// GetInitialIssuesClosed will return the JiraProjectSprint InitialIssuesClosed value
func (t *JiraProjectSprint) GetInitialIssuesClosed() int32 {
	return t.InitialIssuesClosed
}

// SetInitialIssuesClosed will set the JiraProjectSprint InitialIssuesClosed value
func (t *JiraProjectSprint) SetInitialIssuesClosed(v int32) {
	t.InitialIssuesClosed = v
}

// GetInitialIssuesInFinalCount will return the JiraProjectSprint InitialIssuesInFinalCount value
func (t *JiraProjectSprint) GetInitialIssuesInFinalCount() int32 {
	return t.InitialIssuesInFinalCount
}

// SetInitialIssuesInFinalCount will set the JiraProjectSprint InitialIssuesInFinalCount value
func (t *JiraProjectSprint) SetInitialIssuesInFinalCount(v int32) {
	t.InitialIssuesInFinalCount = v
}

func (t *JiraProjectSprint) toTimestamp(value time.Time) *timestamp.Timestamp {
	ts, _ := ptypes.TimestampProto(value)
	return ts
}

// DBCreateJiraProjectSprintTable will create the JiraProjectSprint table
func DBCreateJiraProjectSprintTable(ctx context.Context, db *sql.DB) error {
	q := "CREATE TABLE `jira_project_sprint` (`id` VARCHAR(64) NOT NULL PRIMARY KEY,`checksum` CHAR(64),`project_id`VARCHAR(64) NOT NULL,`sprint_id` TEXT NOT NULL,`name`TEXT NOT NULL,`customer_id` VARCHAR(64) NOT NULL,`complete_date`BIGINT UNSIGNED,`end_date` BIGINT UNSIGNED,`start_date`BIGINT UNSIGNED,`state` TEXT,`goal`TEXT,`initial_issue_ids` JSON,`final_issue_ids` JSON,`final_closed_issue_ids`JSON,`initial_issue_count`INT NOT NULL,`final_issue_count` INT NOT NULL,`closed_issue_count` INT NOT NULL,`added_issue_count` INT NOT NULL,`removed_issue_count`INT NOT NULL,`initial_issues_closed` INT NOT NULL,`initial_issues_in_final_count` INT NOT NULL,INDEX jira_project_sprint_project_id_index (`project_id`),INDEX jira_project_sprint_customer_id_index (`customer_id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBCreateJiraProjectSprintTableTx will create the JiraProjectSprint table using the provided transction
func DBCreateJiraProjectSprintTableTx(ctx context.Context, tx *sql.Tx) error {
	q := "CREATE TABLE `jira_project_sprint` (`id` VARCHAR(64) NOT NULL PRIMARY KEY,`checksum` CHAR(64),`project_id`VARCHAR(64) NOT NULL,`sprint_id` TEXT NOT NULL,`name`TEXT NOT NULL,`customer_id` VARCHAR(64) NOT NULL,`complete_date`BIGINT UNSIGNED,`end_date` BIGINT UNSIGNED,`start_date`BIGINT UNSIGNED,`state` TEXT,`goal`TEXT,`initial_issue_ids` JSON,`final_issue_ids` JSON,`final_closed_issue_ids`JSON,`initial_issue_count`INT NOT NULL,`final_issue_count` INT NOT NULL,`closed_issue_count` INT NOT NULL,`added_issue_count` INT NOT NULL,`removed_issue_count`INT NOT NULL,`initial_issues_closed` INT NOT NULL,`initial_issues_in_final_count` INT NOT NULL,INDEX jira_project_sprint_project_id_index (`project_id`),INDEX jira_project_sprint_customer_id_index (`customer_id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// DBDropJiraProjectSprintTable will drop the JiraProjectSprint table
func DBDropJiraProjectSprintTable(ctx context.Context, db *sql.DB) error {
	q := "DROP TABLE IF EXISTS `jira_project_sprint`"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBDropJiraProjectSprintTableTx will drop the JiraProjectSprint table using the provided transaction
func DBDropJiraProjectSprintTableTx(ctx context.Context, tx *sql.Tx) error {
	q := "DROP TABLE IF EXISTS `jira_project_sprint`"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// CalculateChecksum will calculate a checksum of the SHA1 of all field values
func (t *JiraProjectSprint) CalculateChecksum() string {
	return orm.HashStrings(
		orm.ToString(t.ID),
		orm.ToString(t.ProjectID),
		orm.ToString(t.SprintID),
		orm.ToString(t.Name),
		orm.ToString(t.CustomerID),
		orm.ToString(t.CompleteDate),
		orm.ToString(t.EndDate),
		orm.ToString(t.StartDate),
		orm.ToString(t.State),
		orm.ToString(t.Goal),
		orm.ToString(t.InitialIssueIds),
		orm.ToString(t.FinalIssueIds),
		orm.ToString(t.FinalClosedIssueIds),
		orm.ToString(t.InitialIssueCount),
		orm.ToString(t.FinalIssueCount),
		orm.ToString(t.ClosedIssueCount),
		orm.ToString(t.AddedIssueCount),
		orm.ToString(t.RemovedIssueCount),
		orm.ToString(t.InitialIssuesClosed),
		orm.ToString(t.InitialIssuesInFinalCount),
	)
}

// DBCreate will create a new JiraProjectSprint record in the database
func (t *JiraProjectSprint) DBCreate(ctx context.Context, db *sql.DB) (sql.Result, error) {
	q := "INSERT INTO `jira_project_sprint` (`jira_project_sprint`.`id`,`jira_project_sprint`.`checksum`,`jira_project_sprint`.`project_id`,`jira_project_sprint`.`sprint_id`,`jira_project_sprint`.`name`,`jira_project_sprint`.`customer_id`,`jira_project_sprint`.`complete_date`,`jira_project_sprint`.`end_date`,`jira_project_sprint`.`start_date`,`jira_project_sprint`.`state`,`jira_project_sprint`.`goal`,`jira_project_sprint`.`initial_issue_ids`,`jira_project_sprint`.`final_issue_ids`,`jira_project_sprint`.`final_closed_issue_ids`,`jira_project_sprint`.`initial_issue_count`,`jira_project_sprint`.`final_issue_count`,`jira_project_sprint`.`closed_issue_count`,`jira_project_sprint`.`added_issue_count`,`jira_project_sprint`.`removed_issue_count`,`jira_project_sprint`.`initial_issues_closed`,`jira_project_sprint`.`initial_issues_in_final_count`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLString(t.SprintID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLInt64(t.CompleteDate),
		orm.ToSQLInt64(t.EndDate),
		orm.ToSQLInt64(t.StartDate),
		orm.ToSQLString(t.State),
		orm.ToSQLString(t.Goal),
		orm.ToSQLString(t.InitialIssueIds),
		orm.ToSQLString(t.FinalIssueIds),
		orm.ToSQLString(t.FinalClosedIssueIds),
		orm.ToSQLInt64(t.InitialIssueCount),
		orm.ToSQLInt64(t.FinalIssueCount),
		orm.ToSQLInt64(t.ClosedIssueCount),
		orm.ToSQLInt64(t.AddedIssueCount),
		orm.ToSQLInt64(t.RemovedIssueCount),
		orm.ToSQLInt64(t.InitialIssuesClosed),
		orm.ToSQLInt64(t.InitialIssuesInFinalCount),
	)
}

// DBCreateTx will create a new JiraProjectSprint record in the database using the provided transaction
func (t *JiraProjectSprint) DBCreateTx(ctx context.Context, tx *sql.Tx) (sql.Result, error) {
	q := "INSERT INTO `jira_project_sprint` (`jira_project_sprint`.`id`,`jira_project_sprint`.`checksum`,`jira_project_sprint`.`project_id`,`jira_project_sprint`.`sprint_id`,`jira_project_sprint`.`name`,`jira_project_sprint`.`customer_id`,`jira_project_sprint`.`complete_date`,`jira_project_sprint`.`end_date`,`jira_project_sprint`.`start_date`,`jira_project_sprint`.`state`,`jira_project_sprint`.`goal`,`jira_project_sprint`.`initial_issue_ids`,`jira_project_sprint`.`final_issue_ids`,`jira_project_sprint`.`final_closed_issue_ids`,`jira_project_sprint`.`initial_issue_count`,`jira_project_sprint`.`final_issue_count`,`jira_project_sprint`.`closed_issue_count`,`jira_project_sprint`.`added_issue_count`,`jira_project_sprint`.`removed_issue_count`,`jira_project_sprint`.`initial_issues_closed`,`jira_project_sprint`.`initial_issues_in_final_count`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLString(t.SprintID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLInt64(t.CompleteDate),
		orm.ToSQLInt64(t.EndDate),
		orm.ToSQLInt64(t.StartDate),
		orm.ToSQLString(t.State),
		orm.ToSQLString(t.Goal),
		orm.ToSQLString(t.InitialIssueIds),
		orm.ToSQLString(t.FinalIssueIds),
		orm.ToSQLString(t.FinalClosedIssueIds),
		orm.ToSQLInt64(t.InitialIssueCount),
		orm.ToSQLInt64(t.FinalIssueCount),
		orm.ToSQLInt64(t.ClosedIssueCount),
		orm.ToSQLInt64(t.AddedIssueCount),
		orm.ToSQLInt64(t.RemovedIssueCount),
		orm.ToSQLInt64(t.InitialIssuesClosed),
		orm.ToSQLInt64(t.InitialIssuesInFinalCount),
	)
}

// DBCreateIgnoreDuplicate will upsert the JiraProjectSprint record in the database
func (t *JiraProjectSprint) DBCreateIgnoreDuplicate(ctx context.Context, db *sql.DB) (sql.Result, error) {
	q := "INSERT INTO `jira_project_sprint` (`jira_project_sprint`.`id`,`jira_project_sprint`.`checksum`,`jira_project_sprint`.`project_id`,`jira_project_sprint`.`sprint_id`,`jira_project_sprint`.`name`,`jira_project_sprint`.`customer_id`,`jira_project_sprint`.`complete_date`,`jira_project_sprint`.`end_date`,`jira_project_sprint`.`start_date`,`jira_project_sprint`.`state`,`jira_project_sprint`.`goal`,`jira_project_sprint`.`initial_issue_ids`,`jira_project_sprint`.`final_issue_ids`,`jira_project_sprint`.`final_closed_issue_ids`,`jira_project_sprint`.`initial_issue_count`,`jira_project_sprint`.`final_issue_count`,`jira_project_sprint`.`closed_issue_count`,`jira_project_sprint`.`added_issue_count`,`jira_project_sprint`.`removed_issue_count`,`jira_project_sprint`.`initial_issues_closed`,`jira_project_sprint`.`initial_issues_in_final_count`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLString(t.SprintID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLInt64(t.CompleteDate),
		orm.ToSQLInt64(t.EndDate),
		orm.ToSQLInt64(t.StartDate),
		orm.ToSQLString(t.State),
		orm.ToSQLString(t.Goal),
		orm.ToSQLString(t.InitialIssueIds),
		orm.ToSQLString(t.FinalIssueIds),
		orm.ToSQLString(t.FinalClosedIssueIds),
		orm.ToSQLInt64(t.InitialIssueCount),
		orm.ToSQLInt64(t.FinalIssueCount),
		orm.ToSQLInt64(t.ClosedIssueCount),
		orm.ToSQLInt64(t.AddedIssueCount),
		orm.ToSQLInt64(t.RemovedIssueCount),
		orm.ToSQLInt64(t.InitialIssuesClosed),
		orm.ToSQLInt64(t.InitialIssuesInFinalCount),
	)
}

// DBCreateIgnoreDuplicateTx will upsert the JiraProjectSprint record in the database using the provided transaction
func (t *JiraProjectSprint) DBCreateIgnoreDuplicateTx(ctx context.Context, tx *sql.Tx) (sql.Result, error) {
	q := "INSERT INTO `jira_project_sprint` (`jira_project_sprint`.`id`,`jira_project_sprint`.`checksum`,`jira_project_sprint`.`project_id`,`jira_project_sprint`.`sprint_id`,`jira_project_sprint`.`name`,`jira_project_sprint`.`customer_id`,`jira_project_sprint`.`complete_date`,`jira_project_sprint`.`end_date`,`jira_project_sprint`.`start_date`,`jira_project_sprint`.`state`,`jira_project_sprint`.`goal`,`jira_project_sprint`.`initial_issue_ids`,`jira_project_sprint`.`final_issue_ids`,`jira_project_sprint`.`final_closed_issue_ids`,`jira_project_sprint`.`initial_issue_count`,`jira_project_sprint`.`final_issue_count`,`jira_project_sprint`.`closed_issue_count`,`jira_project_sprint`.`added_issue_count`,`jira_project_sprint`.`removed_issue_count`,`jira_project_sprint`.`initial_issues_closed`,`jira_project_sprint`.`initial_issues_in_final_count`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLString(t.SprintID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLInt64(t.CompleteDate),
		orm.ToSQLInt64(t.EndDate),
		orm.ToSQLInt64(t.StartDate),
		orm.ToSQLString(t.State),
		orm.ToSQLString(t.Goal),
		orm.ToSQLString(t.InitialIssueIds),
		orm.ToSQLString(t.FinalIssueIds),
		orm.ToSQLString(t.FinalClosedIssueIds),
		orm.ToSQLInt64(t.InitialIssueCount),
		orm.ToSQLInt64(t.FinalIssueCount),
		orm.ToSQLInt64(t.ClosedIssueCount),
		orm.ToSQLInt64(t.AddedIssueCount),
		orm.ToSQLInt64(t.RemovedIssueCount),
		orm.ToSQLInt64(t.InitialIssuesClosed),
		orm.ToSQLInt64(t.InitialIssuesInFinalCount),
	)
}

// DeleteAllJiraProjectSprints deletes all JiraProjectSprint records in the database with optional filters
func DeleteAllJiraProjectSprints(ctx context.Context, db *sql.DB, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(JiraProjectSprintTableName),
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

// DeleteAllJiraProjectSprintsTx deletes all JiraProjectSprint records in the database with optional filters using the provided transaction
func DeleteAllJiraProjectSprintsTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(JiraProjectSprintTableName),
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

// DBDelete will delete this JiraProjectSprint record in the database
func (t *JiraProjectSprint) DBDelete(ctx context.Context, db *sql.DB) (bool, error) {
	q := "DELETE FROM `jira_project_sprint` WHERE `id` = ?"
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

// DBDeleteTx will delete this JiraProjectSprint record in the database using the provided transaction
func (t *JiraProjectSprint) DBDeleteTx(ctx context.Context, tx *sql.Tx) (bool, error) {
	q := "DELETE FROM `jira_project_sprint` WHERE `id` = ?"
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

// DBUpdate will update the JiraProjectSprint record in the database
func (t *JiraProjectSprint) DBUpdate(ctx context.Context, db *sql.DB) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `jira_project_sprint` SET `checksum`=?,`project_id`=?,`sprint_id`=?,`name`=?,`customer_id`=?,`complete_date`=?,`end_date`=?,`start_date`=?,`state`=?,`goal`=?,`initial_issue_ids`=?,`final_issue_ids`=?,`final_closed_issue_ids`=?,`initial_issue_count`=?,`final_issue_count`=?,`closed_issue_count`=?,`added_issue_count`=?,`removed_issue_count`=?,`initial_issues_closed`=?,`initial_issues_in_final_count`=? WHERE `id`=?"
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLString(t.SprintID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLInt64(t.CompleteDate),
		orm.ToSQLInt64(t.EndDate),
		orm.ToSQLInt64(t.StartDate),
		orm.ToSQLString(t.State),
		orm.ToSQLString(t.Goal),
		orm.ToSQLString(t.InitialIssueIds),
		orm.ToSQLString(t.FinalIssueIds),
		orm.ToSQLString(t.FinalClosedIssueIds),
		orm.ToSQLInt64(t.InitialIssueCount),
		orm.ToSQLInt64(t.FinalIssueCount),
		orm.ToSQLInt64(t.ClosedIssueCount),
		orm.ToSQLInt64(t.AddedIssueCount),
		orm.ToSQLInt64(t.RemovedIssueCount),
		orm.ToSQLInt64(t.InitialIssuesClosed),
		orm.ToSQLInt64(t.InitialIssuesInFinalCount),
		orm.ToSQLString(t.ID),
	)
}

// DBUpdateTx will update the JiraProjectSprint record in the database using the provided transaction
func (t *JiraProjectSprint) DBUpdateTx(ctx context.Context, tx *sql.Tx) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `jira_project_sprint` SET `checksum`=?,`project_id`=?,`sprint_id`=?,`name`=?,`customer_id`=?,`complete_date`=?,`end_date`=?,`start_date`=?,`state`=?,`goal`=?,`initial_issue_ids`=?,`final_issue_ids`=?,`final_closed_issue_ids`=?,`initial_issue_count`=?,`final_issue_count`=?,`closed_issue_count`=?,`added_issue_count`=?,`removed_issue_count`=?,`initial_issues_closed`=?,`initial_issues_in_final_count`=? WHERE `id`=?"
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLString(t.SprintID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLInt64(t.CompleteDate),
		orm.ToSQLInt64(t.EndDate),
		orm.ToSQLInt64(t.StartDate),
		orm.ToSQLString(t.State),
		orm.ToSQLString(t.Goal),
		orm.ToSQLString(t.InitialIssueIds),
		orm.ToSQLString(t.FinalIssueIds),
		orm.ToSQLString(t.FinalClosedIssueIds),
		orm.ToSQLInt64(t.InitialIssueCount),
		orm.ToSQLInt64(t.FinalIssueCount),
		orm.ToSQLInt64(t.ClosedIssueCount),
		orm.ToSQLInt64(t.AddedIssueCount),
		orm.ToSQLInt64(t.RemovedIssueCount),
		orm.ToSQLInt64(t.InitialIssuesClosed),
		orm.ToSQLInt64(t.InitialIssuesInFinalCount),
		orm.ToSQLString(t.ID),
	)
}

// DBUpsert will upsert the JiraProjectSprint record in the database
func (t *JiraProjectSprint) DBUpsert(ctx context.Context, db *sql.DB, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `jira_project_sprint` (`jira_project_sprint`.`id`,`jira_project_sprint`.`checksum`,`jira_project_sprint`.`project_id`,`jira_project_sprint`.`sprint_id`,`jira_project_sprint`.`name`,`jira_project_sprint`.`customer_id`,`jira_project_sprint`.`complete_date`,`jira_project_sprint`.`end_date`,`jira_project_sprint`.`start_date`,`jira_project_sprint`.`state`,`jira_project_sprint`.`goal`,`jira_project_sprint`.`initial_issue_ids`,`jira_project_sprint`.`final_issue_ids`,`jira_project_sprint`.`final_closed_issue_ids`,`jira_project_sprint`.`initial_issue_count`,`jira_project_sprint`.`final_issue_count`,`jira_project_sprint`.`closed_issue_count`,`jira_project_sprint`.`added_issue_count`,`jira_project_sprint`.`removed_issue_count`,`jira_project_sprint`.`initial_issues_closed`,`jira_project_sprint`.`initial_issues_in_final_count`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `jira_project_sprint` (`jira_project_sprint`.`id`,`jira_project_sprint`.`checksum`,`jira_project_sprint`.`project_id`,`jira_project_sprint`.`sprint_id`,`jira_project_sprint`.`name`,`jira_project_sprint`.`customer_id`,`jira_project_sprint`.`complete_date`,`jira_project_sprint`.`end_date`,`jira_project_sprint`.`start_date`,`jira_project_sprint`.`state`,`jira_project_sprint`.`goal`,`jira_project_sprint`.`initial_issue_ids`,`jira_project_sprint`.`final_issue_ids`,`jira_project_sprint`.`final_closed_issue_ids`,`jira_project_sprint`.`initial_issue_count`,`jira_project_sprint`.`final_issue_count`,`jira_project_sprint`.`closed_issue_count`,`jira_project_sprint`.`added_issue_count`,`jira_project_sprint`.`removed_issue_count`,`jira_project_sprint`.`initial_issues_closed`,`jira_project_sprint`.`initial_issues_in_final_count`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`project_id`=VALUES(`project_id`),`sprint_id`=VALUES(`sprint_id`),`name`=VALUES(`name`),`customer_id`=VALUES(`customer_id`),`complete_date`=VALUES(`complete_date`),`end_date`=VALUES(`end_date`),`start_date`=VALUES(`start_date`),`state`=VALUES(`state`),`goal`=VALUES(`goal`),`initial_issue_ids`=VALUES(`initial_issue_ids`),`final_issue_ids`=VALUES(`final_issue_ids`),`final_closed_issue_ids`=VALUES(`final_closed_issue_ids`),`initial_issue_count`=VALUES(`initial_issue_count`),`final_issue_count`=VALUES(`final_issue_count`),`closed_issue_count`=VALUES(`closed_issue_count`),`added_issue_count`=VALUES(`added_issue_count`),`removed_issue_count`=VALUES(`removed_issue_count`),`initial_issues_closed`=VALUES(`initial_issues_closed`),`initial_issues_in_final_count`=VALUES(`initial_issues_in_final_count`)"
	}
	r, err := db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLString(t.SprintID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLInt64(t.CompleteDate),
		orm.ToSQLInt64(t.EndDate),
		orm.ToSQLInt64(t.StartDate),
		orm.ToSQLString(t.State),
		orm.ToSQLString(t.Goal),
		orm.ToSQLString(t.InitialIssueIds),
		orm.ToSQLString(t.FinalIssueIds),
		orm.ToSQLString(t.FinalClosedIssueIds),
		orm.ToSQLInt64(t.InitialIssueCount),
		orm.ToSQLInt64(t.FinalIssueCount),
		orm.ToSQLInt64(t.ClosedIssueCount),
		orm.ToSQLInt64(t.AddedIssueCount),
		orm.ToSQLInt64(t.RemovedIssueCount),
		orm.ToSQLInt64(t.InitialIssuesClosed),
		orm.ToSQLInt64(t.InitialIssuesInFinalCount),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBUpsertTx will upsert the JiraProjectSprint record in the database using the provided transaction
func (t *JiraProjectSprint) DBUpsertTx(ctx context.Context, tx *sql.Tx, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `jira_project_sprint` (`jira_project_sprint`.`id`,`jira_project_sprint`.`checksum`,`jira_project_sprint`.`project_id`,`jira_project_sprint`.`sprint_id`,`jira_project_sprint`.`name`,`jira_project_sprint`.`customer_id`,`jira_project_sprint`.`complete_date`,`jira_project_sprint`.`end_date`,`jira_project_sprint`.`start_date`,`jira_project_sprint`.`state`,`jira_project_sprint`.`goal`,`jira_project_sprint`.`initial_issue_ids`,`jira_project_sprint`.`final_issue_ids`,`jira_project_sprint`.`final_closed_issue_ids`,`jira_project_sprint`.`initial_issue_count`,`jira_project_sprint`.`final_issue_count`,`jira_project_sprint`.`closed_issue_count`,`jira_project_sprint`.`added_issue_count`,`jira_project_sprint`.`removed_issue_count`,`jira_project_sprint`.`initial_issues_closed`,`jira_project_sprint`.`initial_issues_in_final_count`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `jira_project_sprint` (`jira_project_sprint`.`id`,`jira_project_sprint`.`checksum`,`jira_project_sprint`.`project_id`,`jira_project_sprint`.`sprint_id`,`jira_project_sprint`.`name`,`jira_project_sprint`.`customer_id`,`jira_project_sprint`.`complete_date`,`jira_project_sprint`.`end_date`,`jira_project_sprint`.`start_date`,`jira_project_sprint`.`state`,`jira_project_sprint`.`goal`,`jira_project_sprint`.`initial_issue_ids`,`jira_project_sprint`.`final_issue_ids`,`jira_project_sprint`.`final_closed_issue_ids`,`jira_project_sprint`.`initial_issue_count`,`jira_project_sprint`.`final_issue_count`,`jira_project_sprint`.`closed_issue_count`,`jira_project_sprint`.`added_issue_count`,`jira_project_sprint`.`removed_issue_count`,`jira_project_sprint`.`initial_issues_closed`,`jira_project_sprint`.`initial_issues_in_final_count`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`project_id`=VALUES(`project_id`),`sprint_id`=VALUES(`sprint_id`),`name`=VALUES(`name`),`customer_id`=VALUES(`customer_id`),`complete_date`=VALUES(`complete_date`),`end_date`=VALUES(`end_date`),`start_date`=VALUES(`start_date`),`state`=VALUES(`state`),`goal`=VALUES(`goal`),`initial_issue_ids`=VALUES(`initial_issue_ids`),`final_issue_ids`=VALUES(`final_issue_ids`),`final_closed_issue_ids`=VALUES(`final_closed_issue_ids`),`initial_issue_count`=VALUES(`initial_issue_count`),`final_issue_count`=VALUES(`final_issue_count`),`closed_issue_count`=VALUES(`closed_issue_count`),`added_issue_count`=VALUES(`added_issue_count`),`removed_issue_count`=VALUES(`removed_issue_count`),`initial_issues_closed`=VALUES(`initial_issues_closed`),`initial_issues_in_final_count`=VALUES(`initial_issues_in_final_count`)"
	}
	r, err := tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLString(t.SprintID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLInt64(t.CompleteDate),
		orm.ToSQLInt64(t.EndDate),
		orm.ToSQLInt64(t.StartDate),
		orm.ToSQLString(t.State),
		orm.ToSQLString(t.Goal),
		orm.ToSQLString(t.InitialIssueIds),
		orm.ToSQLString(t.FinalIssueIds),
		orm.ToSQLString(t.FinalClosedIssueIds),
		orm.ToSQLInt64(t.InitialIssueCount),
		orm.ToSQLInt64(t.FinalIssueCount),
		orm.ToSQLInt64(t.ClosedIssueCount),
		orm.ToSQLInt64(t.AddedIssueCount),
		orm.ToSQLInt64(t.RemovedIssueCount),
		orm.ToSQLInt64(t.InitialIssuesClosed),
		orm.ToSQLInt64(t.InitialIssuesInFinalCount),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBFindOne will find a JiraProjectSprint record in the database with the primary key
func (t *JiraProjectSprint) DBFindOne(ctx context.Context, db *sql.DB, value string) (bool, error) {
	q := "SELECT `jira_project_sprint`.`id`,`jira_project_sprint`.`checksum`,`jira_project_sprint`.`project_id`,`jira_project_sprint`.`sprint_id`,`jira_project_sprint`.`name`,`jira_project_sprint`.`customer_id`,`jira_project_sprint`.`complete_date`,`jira_project_sprint`.`end_date`,`jira_project_sprint`.`start_date`,`jira_project_sprint`.`state`,`jira_project_sprint`.`goal`,`jira_project_sprint`.`initial_issue_ids`,`jira_project_sprint`.`final_issue_ids`,`jira_project_sprint`.`final_closed_issue_ids`,`jira_project_sprint`.`initial_issue_count`,`jira_project_sprint`.`final_issue_count`,`jira_project_sprint`.`closed_issue_count`,`jira_project_sprint`.`added_issue_count`,`jira_project_sprint`.`removed_issue_count`,`jira_project_sprint`.`initial_issues_closed`,`jira_project_sprint`.`initial_issues_in_final_count` FROM `jira_project_sprint` WHERE `id` = ? LIMIT 1"
	row := db.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _ProjectID sql.NullString
	var _SprintID sql.NullString
	var _Name sql.NullString
	var _CustomerID sql.NullString
	var _CompleteDate sql.NullInt64
	var _EndDate sql.NullInt64
	var _StartDate sql.NullInt64
	var _State sql.NullString
	var _Goal sql.NullString
	var _InitialIssueIds sql.NullString
	var _FinalIssueIds sql.NullString
	var _FinalClosedIssueIds sql.NullString
	var _InitialIssueCount sql.NullInt64
	var _FinalIssueCount sql.NullInt64
	var _ClosedIssueCount sql.NullInt64
	var _AddedIssueCount sql.NullInt64
	var _RemovedIssueCount sql.NullInt64
	var _InitialIssuesClosed sql.NullInt64
	var _InitialIssuesInFinalCount sql.NullInt64
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_ProjectID,
		&_SprintID,
		&_Name,
		&_CustomerID,
		&_CompleteDate,
		&_EndDate,
		&_StartDate,
		&_State,
		&_Goal,
		&_InitialIssueIds,
		&_FinalIssueIds,
		&_FinalClosedIssueIds,
		&_InitialIssueCount,
		&_FinalIssueCount,
		&_ClosedIssueCount,
		&_AddedIssueCount,
		&_RemovedIssueCount,
		&_InitialIssuesClosed,
		&_InitialIssuesInFinalCount,
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
	if _ProjectID.Valid {
		t.SetProjectID(_ProjectID.String)
	}
	if _SprintID.Valid {
		t.SetSprintID(_SprintID.String)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _CompleteDate.Valid {
		t.SetCompleteDate(_CompleteDate.Int64)
	}
	if _EndDate.Valid {
		t.SetEndDate(_EndDate.Int64)
	}
	if _StartDate.Valid {
		t.SetStartDate(_StartDate.Int64)
	}
	if _State.Valid {
		t.SetState(_State.String)
	}
	if _Goal.Valid {
		t.SetGoal(_Goal.String)
	}
	if _InitialIssueIds.Valid {
		t.SetInitialIssueIds(_InitialIssueIds.String)
	}
	if _FinalIssueIds.Valid {
		t.SetFinalIssueIds(_FinalIssueIds.String)
	}
	if _FinalClosedIssueIds.Valid {
		t.SetFinalClosedIssueIds(_FinalClosedIssueIds.String)
	}
	if _InitialIssueCount.Valid {
		t.SetInitialIssueCount(int32(_InitialIssueCount.Int64))
	}
	if _FinalIssueCount.Valid {
		t.SetFinalIssueCount(int32(_FinalIssueCount.Int64))
	}
	if _ClosedIssueCount.Valid {
		t.SetClosedIssueCount(int32(_ClosedIssueCount.Int64))
	}
	if _AddedIssueCount.Valid {
		t.SetAddedIssueCount(int32(_AddedIssueCount.Int64))
	}
	if _RemovedIssueCount.Valid {
		t.SetRemovedIssueCount(int32(_RemovedIssueCount.Int64))
	}
	if _InitialIssuesClosed.Valid {
		t.SetInitialIssuesClosed(int32(_InitialIssuesClosed.Int64))
	}
	if _InitialIssuesInFinalCount.Valid {
		t.SetInitialIssuesInFinalCount(int32(_InitialIssuesInFinalCount.Int64))
	}
	return true, nil
}

// DBFindOneTx will find a JiraProjectSprint record in the database with the primary key using the provided transaction
func (t *JiraProjectSprint) DBFindOneTx(ctx context.Context, tx *sql.Tx, value string) (bool, error) {
	q := "SELECT `jira_project_sprint`.`id`,`jira_project_sprint`.`checksum`,`jira_project_sprint`.`project_id`,`jira_project_sprint`.`sprint_id`,`jira_project_sprint`.`name`,`jira_project_sprint`.`customer_id`,`jira_project_sprint`.`complete_date`,`jira_project_sprint`.`end_date`,`jira_project_sprint`.`start_date`,`jira_project_sprint`.`state`,`jira_project_sprint`.`goal`,`jira_project_sprint`.`initial_issue_ids`,`jira_project_sprint`.`final_issue_ids`,`jira_project_sprint`.`final_closed_issue_ids`,`jira_project_sprint`.`initial_issue_count`,`jira_project_sprint`.`final_issue_count`,`jira_project_sprint`.`closed_issue_count`,`jira_project_sprint`.`added_issue_count`,`jira_project_sprint`.`removed_issue_count`,`jira_project_sprint`.`initial_issues_closed`,`jira_project_sprint`.`initial_issues_in_final_count` FROM `jira_project_sprint` WHERE `id` = ? LIMIT 1"
	row := tx.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _ProjectID sql.NullString
	var _SprintID sql.NullString
	var _Name sql.NullString
	var _CustomerID sql.NullString
	var _CompleteDate sql.NullInt64
	var _EndDate sql.NullInt64
	var _StartDate sql.NullInt64
	var _State sql.NullString
	var _Goal sql.NullString
	var _InitialIssueIds sql.NullString
	var _FinalIssueIds sql.NullString
	var _FinalClosedIssueIds sql.NullString
	var _InitialIssueCount sql.NullInt64
	var _FinalIssueCount sql.NullInt64
	var _ClosedIssueCount sql.NullInt64
	var _AddedIssueCount sql.NullInt64
	var _RemovedIssueCount sql.NullInt64
	var _InitialIssuesClosed sql.NullInt64
	var _InitialIssuesInFinalCount sql.NullInt64
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_ProjectID,
		&_SprintID,
		&_Name,
		&_CustomerID,
		&_CompleteDate,
		&_EndDate,
		&_StartDate,
		&_State,
		&_Goal,
		&_InitialIssueIds,
		&_FinalIssueIds,
		&_FinalClosedIssueIds,
		&_InitialIssueCount,
		&_FinalIssueCount,
		&_ClosedIssueCount,
		&_AddedIssueCount,
		&_RemovedIssueCount,
		&_InitialIssuesClosed,
		&_InitialIssuesInFinalCount,
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
	if _ProjectID.Valid {
		t.SetProjectID(_ProjectID.String)
	}
	if _SprintID.Valid {
		t.SetSprintID(_SprintID.String)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _CompleteDate.Valid {
		t.SetCompleteDate(_CompleteDate.Int64)
	}
	if _EndDate.Valid {
		t.SetEndDate(_EndDate.Int64)
	}
	if _StartDate.Valid {
		t.SetStartDate(_StartDate.Int64)
	}
	if _State.Valid {
		t.SetState(_State.String)
	}
	if _Goal.Valid {
		t.SetGoal(_Goal.String)
	}
	if _InitialIssueIds.Valid {
		t.SetInitialIssueIds(_InitialIssueIds.String)
	}
	if _FinalIssueIds.Valid {
		t.SetFinalIssueIds(_FinalIssueIds.String)
	}
	if _FinalClosedIssueIds.Valid {
		t.SetFinalClosedIssueIds(_FinalClosedIssueIds.String)
	}
	if _InitialIssueCount.Valid {
		t.SetInitialIssueCount(int32(_InitialIssueCount.Int64))
	}
	if _FinalIssueCount.Valid {
		t.SetFinalIssueCount(int32(_FinalIssueCount.Int64))
	}
	if _ClosedIssueCount.Valid {
		t.SetClosedIssueCount(int32(_ClosedIssueCount.Int64))
	}
	if _AddedIssueCount.Valid {
		t.SetAddedIssueCount(int32(_AddedIssueCount.Int64))
	}
	if _RemovedIssueCount.Valid {
		t.SetRemovedIssueCount(int32(_RemovedIssueCount.Int64))
	}
	if _InitialIssuesClosed.Valid {
		t.SetInitialIssuesClosed(int32(_InitialIssuesClosed.Int64))
	}
	if _InitialIssuesInFinalCount.Valid {
		t.SetInitialIssuesInFinalCount(int32(_InitialIssuesInFinalCount.Int64))
	}
	return true, nil
}

// FindJiraProjectSprints will find a JiraProjectSprint record in the database with the provided parameters
func FindJiraProjectSprints(ctx context.Context, db *sql.DB, _params ...interface{}) ([]*JiraProjectSprint, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("project_id"),
		orm.Column("sprint_id"),
		orm.Column("name"),
		orm.Column("customer_id"),
		orm.Column("complete_date"),
		orm.Column("end_date"),
		orm.Column("start_date"),
		orm.Column("state"),
		orm.Column("goal"),
		orm.Column("initial_issue_ids"),
		orm.Column("final_issue_ids"),
		orm.Column("final_closed_issue_ids"),
		orm.Column("initial_issue_count"),
		orm.Column("final_issue_count"),
		orm.Column("closed_issue_count"),
		orm.Column("added_issue_count"),
		orm.Column("removed_issue_count"),
		orm.Column("initial_issues_closed"),
		orm.Column("initial_issues_in_final_count"),
		orm.Table(JiraProjectSprintTableName),
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
	results := make([]*JiraProjectSprint, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _ProjectID sql.NullString
		var _SprintID sql.NullString
		var _Name sql.NullString
		var _CustomerID sql.NullString
		var _CompleteDate sql.NullInt64
		var _EndDate sql.NullInt64
		var _StartDate sql.NullInt64
		var _State sql.NullString
		var _Goal sql.NullString
		var _InitialIssueIds sql.NullString
		var _FinalIssueIds sql.NullString
		var _FinalClosedIssueIds sql.NullString
		var _InitialIssueCount sql.NullInt64
		var _FinalIssueCount sql.NullInt64
		var _ClosedIssueCount sql.NullInt64
		var _AddedIssueCount sql.NullInt64
		var _RemovedIssueCount sql.NullInt64
		var _InitialIssuesClosed sql.NullInt64
		var _InitialIssuesInFinalCount sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_ProjectID,
			&_SprintID,
			&_Name,
			&_CustomerID,
			&_CompleteDate,
			&_EndDate,
			&_StartDate,
			&_State,
			&_Goal,
			&_InitialIssueIds,
			&_FinalIssueIds,
			&_FinalClosedIssueIds,
			&_InitialIssueCount,
			&_FinalIssueCount,
			&_ClosedIssueCount,
			&_AddedIssueCount,
			&_RemovedIssueCount,
			&_InitialIssuesClosed,
			&_InitialIssuesInFinalCount,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraProjectSprint{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _SprintID.Valid {
			t.SetSprintID(_SprintID.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _CompleteDate.Valid {
			t.SetCompleteDate(_CompleteDate.Int64)
		}
		if _EndDate.Valid {
			t.SetEndDate(_EndDate.Int64)
		}
		if _StartDate.Valid {
			t.SetStartDate(_StartDate.Int64)
		}
		if _State.Valid {
			t.SetState(_State.String)
		}
		if _Goal.Valid {
			t.SetGoal(_Goal.String)
		}
		if _InitialIssueIds.Valid {
			t.SetInitialIssueIds(_InitialIssueIds.String)
		}
		if _FinalIssueIds.Valid {
			t.SetFinalIssueIds(_FinalIssueIds.String)
		}
		if _FinalClosedIssueIds.Valid {
			t.SetFinalClosedIssueIds(_FinalClosedIssueIds.String)
		}
		if _InitialIssueCount.Valid {
			t.SetInitialIssueCount(int32(_InitialIssueCount.Int64))
		}
		if _FinalIssueCount.Valid {
			t.SetFinalIssueCount(int32(_FinalIssueCount.Int64))
		}
		if _ClosedIssueCount.Valid {
			t.SetClosedIssueCount(int32(_ClosedIssueCount.Int64))
		}
		if _AddedIssueCount.Valid {
			t.SetAddedIssueCount(int32(_AddedIssueCount.Int64))
		}
		if _RemovedIssueCount.Valid {
			t.SetRemovedIssueCount(int32(_RemovedIssueCount.Int64))
		}
		if _InitialIssuesClosed.Valid {
			t.SetInitialIssuesClosed(int32(_InitialIssuesClosed.Int64))
		}
		if _InitialIssuesInFinalCount.Valid {
			t.SetInitialIssuesInFinalCount(int32(_InitialIssuesInFinalCount.Int64))
		}
		results = append(results, t)
	}
	return results, nil
}

// FindJiraProjectSprintsTx will find a JiraProjectSprint record in the database with the provided parameters using the provided transaction
func FindJiraProjectSprintsTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) ([]*JiraProjectSprint, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("project_id"),
		orm.Column("sprint_id"),
		orm.Column("name"),
		orm.Column("customer_id"),
		orm.Column("complete_date"),
		orm.Column("end_date"),
		orm.Column("start_date"),
		orm.Column("state"),
		orm.Column("goal"),
		orm.Column("initial_issue_ids"),
		orm.Column("final_issue_ids"),
		orm.Column("final_closed_issue_ids"),
		orm.Column("initial_issue_count"),
		orm.Column("final_issue_count"),
		orm.Column("closed_issue_count"),
		orm.Column("added_issue_count"),
		orm.Column("removed_issue_count"),
		orm.Column("initial_issues_closed"),
		orm.Column("initial_issues_in_final_count"),
		orm.Table(JiraProjectSprintTableName),
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
	results := make([]*JiraProjectSprint, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _ProjectID sql.NullString
		var _SprintID sql.NullString
		var _Name sql.NullString
		var _CustomerID sql.NullString
		var _CompleteDate sql.NullInt64
		var _EndDate sql.NullInt64
		var _StartDate sql.NullInt64
		var _State sql.NullString
		var _Goal sql.NullString
		var _InitialIssueIds sql.NullString
		var _FinalIssueIds sql.NullString
		var _FinalClosedIssueIds sql.NullString
		var _InitialIssueCount sql.NullInt64
		var _FinalIssueCount sql.NullInt64
		var _ClosedIssueCount sql.NullInt64
		var _AddedIssueCount sql.NullInt64
		var _RemovedIssueCount sql.NullInt64
		var _InitialIssuesClosed sql.NullInt64
		var _InitialIssuesInFinalCount sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_ProjectID,
			&_SprintID,
			&_Name,
			&_CustomerID,
			&_CompleteDate,
			&_EndDate,
			&_StartDate,
			&_State,
			&_Goal,
			&_InitialIssueIds,
			&_FinalIssueIds,
			&_FinalClosedIssueIds,
			&_InitialIssueCount,
			&_FinalIssueCount,
			&_ClosedIssueCount,
			&_AddedIssueCount,
			&_RemovedIssueCount,
			&_InitialIssuesClosed,
			&_InitialIssuesInFinalCount,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraProjectSprint{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _SprintID.Valid {
			t.SetSprintID(_SprintID.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _CompleteDate.Valid {
			t.SetCompleteDate(_CompleteDate.Int64)
		}
		if _EndDate.Valid {
			t.SetEndDate(_EndDate.Int64)
		}
		if _StartDate.Valid {
			t.SetStartDate(_StartDate.Int64)
		}
		if _State.Valid {
			t.SetState(_State.String)
		}
		if _Goal.Valid {
			t.SetGoal(_Goal.String)
		}
		if _InitialIssueIds.Valid {
			t.SetInitialIssueIds(_InitialIssueIds.String)
		}
		if _FinalIssueIds.Valid {
			t.SetFinalIssueIds(_FinalIssueIds.String)
		}
		if _FinalClosedIssueIds.Valid {
			t.SetFinalClosedIssueIds(_FinalClosedIssueIds.String)
		}
		if _InitialIssueCount.Valid {
			t.SetInitialIssueCount(int32(_InitialIssueCount.Int64))
		}
		if _FinalIssueCount.Valid {
			t.SetFinalIssueCount(int32(_FinalIssueCount.Int64))
		}
		if _ClosedIssueCount.Valid {
			t.SetClosedIssueCount(int32(_ClosedIssueCount.Int64))
		}
		if _AddedIssueCount.Valid {
			t.SetAddedIssueCount(int32(_AddedIssueCount.Int64))
		}
		if _RemovedIssueCount.Valid {
			t.SetRemovedIssueCount(int32(_RemovedIssueCount.Int64))
		}
		if _InitialIssuesClosed.Valid {
			t.SetInitialIssuesClosed(int32(_InitialIssuesClosed.Int64))
		}
		if _InitialIssuesInFinalCount.Valid {
			t.SetInitialIssuesInFinalCount(int32(_InitialIssuesInFinalCount.Int64))
		}
		results = append(results, t)
	}
	return results, nil
}

// DBFind will find a JiraProjectSprint record in the database with the provided parameters
func (t *JiraProjectSprint) DBFind(ctx context.Context, db *sql.DB, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("project_id"),
		orm.Column("sprint_id"),
		orm.Column("name"),
		orm.Column("customer_id"),
		orm.Column("complete_date"),
		orm.Column("end_date"),
		orm.Column("start_date"),
		orm.Column("state"),
		orm.Column("goal"),
		orm.Column("initial_issue_ids"),
		orm.Column("final_issue_ids"),
		orm.Column("final_closed_issue_ids"),
		orm.Column("initial_issue_count"),
		orm.Column("final_issue_count"),
		orm.Column("closed_issue_count"),
		orm.Column("added_issue_count"),
		orm.Column("removed_issue_count"),
		orm.Column("initial_issues_closed"),
		orm.Column("initial_issues_in_final_count"),
		orm.Table(JiraProjectSprintTableName),
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
	var _ProjectID sql.NullString
	var _SprintID sql.NullString
	var _Name sql.NullString
	var _CustomerID sql.NullString
	var _CompleteDate sql.NullInt64
	var _EndDate sql.NullInt64
	var _StartDate sql.NullInt64
	var _State sql.NullString
	var _Goal sql.NullString
	var _InitialIssueIds sql.NullString
	var _FinalIssueIds sql.NullString
	var _FinalClosedIssueIds sql.NullString
	var _InitialIssueCount sql.NullInt64
	var _FinalIssueCount sql.NullInt64
	var _ClosedIssueCount sql.NullInt64
	var _AddedIssueCount sql.NullInt64
	var _RemovedIssueCount sql.NullInt64
	var _InitialIssuesClosed sql.NullInt64
	var _InitialIssuesInFinalCount sql.NullInt64
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_ProjectID,
		&_SprintID,
		&_Name,
		&_CustomerID,
		&_CompleteDate,
		&_EndDate,
		&_StartDate,
		&_State,
		&_Goal,
		&_InitialIssueIds,
		&_FinalIssueIds,
		&_FinalClosedIssueIds,
		&_InitialIssueCount,
		&_FinalIssueCount,
		&_ClosedIssueCount,
		&_AddedIssueCount,
		&_RemovedIssueCount,
		&_InitialIssuesClosed,
		&_InitialIssuesInFinalCount,
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
	if _ProjectID.Valid {
		t.SetProjectID(_ProjectID.String)
	}
	if _SprintID.Valid {
		t.SetSprintID(_SprintID.String)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _CompleteDate.Valid {
		t.SetCompleteDate(_CompleteDate.Int64)
	}
	if _EndDate.Valid {
		t.SetEndDate(_EndDate.Int64)
	}
	if _StartDate.Valid {
		t.SetStartDate(_StartDate.Int64)
	}
	if _State.Valid {
		t.SetState(_State.String)
	}
	if _Goal.Valid {
		t.SetGoal(_Goal.String)
	}
	if _InitialIssueIds.Valid {
		t.SetInitialIssueIds(_InitialIssueIds.String)
	}
	if _FinalIssueIds.Valid {
		t.SetFinalIssueIds(_FinalIssueIds.String)
	}
	if _FinalClosedIssueIds.Valid {
		t.SetFinalClosedIssueIds(_FinalClosedIssueIds.String)
	}
	if _InitialIssueCount.Valid {
		t.SetInitialIssueCount(int32(_InitialIssueCount.Int64))
	}
	if _FinalIssueCount.Valid {
		t.SetFinalIssueCount(int32(_FinalIssueCount.Int64))
	}
	if _ClosedIssueCount.Valid {
		t.SetClosedIssueCount(int32(_ClosedIssueCount.Int64))
	}
	if _AddedIssueCount.Valid {
		t.SetAddedIssueCount(int32(_AddedIssueCount.Int64))
	}
	if _RemovedIssueCount.Valid {
		t.SetRemovedIssueCount(int32(_RemovedIssueCount.Int64))
	}
	if _InitialIssuesClosed.Valid {
		t.SetInitialIssuesClosed(int32(_InitialIssuesClosed.Int64))
	}
	if _InitialIssuesInFinalCount.Valid {
		t.SetInitialIssuesInFinalCount(int32(_InitialIssuesInFinalCount.Int64))
	}
	return true, nil
}

// DBFindTx will find a JiraProjectSprint record in the database with the provided parameters using the provided transaction
func (t *JiraProjectSprint) DBFindTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("project_id"),
		orm.Column("sprint_id"),
		orm.Column("name"),
		orm.Column("customer_id"),
		orm.Column("complete_date"),
		orm.Column("end_date"),
		orm.Column("start_date"),
		orm.Column("state"),
		orm.Column("goal"),
		orm.Column("initial_issue_ids"),
		orm.Column("final_issue_ids"),
		orm.Column("final_closed_issue_ids"),
		orm.Column("initial_issue_count"),
		orm.Column("final_issue_count"),
		orm.Column("closed_issue_count"),
		orm.Column("added_issue_count"),
		orm.Column("removed_issue_count"),
		orm.Column("initial_issues_closed"),
		orm.Column("initial_issues_in_final_count"),
		orm.Table(JiraProjectSprintTableName),
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
	var _ProjectID sql.NullString
	var _SprintID sql.NullString
	var _Name sql.NullString
	var _CustomerID sql.NullString
	var _CompleteDate sql.NullInt64
	var _EndDate sql.NullInt64
	var _StartDate sql.NullInt64
	var _State sql.NullString
	var _Goal sql.NullString
	var _InitialIssueIds sql.NullString
	var _FinalIssueIds sql.NullString
	var _FinalClosedIssueIds sql.NullString
	var _InitialIssueCount sql.NullInt64
	var _FinalIssueCount sql.NullInt64
	var _ClosedIssueCount sql.NullInt64
	var _AddedIssueCount sql.NullInt64
	var _RemovedIssueCount sql.NullInt64
	var _InitialIssuesClosed sql.NullInt64
	var _InitialIssuesInFinalCount sql.NullInt64
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_ProjectID,
		&_SprintID,
		&_Name,
		&_CustomerID,
		&_CompleteDate,
		&_EndDate,
		&_StartDate,
		&_State,
		&_Goal,
		&_InitialIssueIds,
		&_FinalIssueIds,
		&_FinalClosedIssueIds,
		&_InitialIssueCount,
		&_FinalIssueCount,
		&_ClosedIssueCount,
		&_AddedIssueCount,
		&_RemovedIssueCount,
		&_InitialIssuesClosed,
		&_InitialIssuesInFinalCount,
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
	if _ProjectID.Valid {
		t.SetProjectID(_ProjectID.String)
	}
	if _SprintID.Valid {
		t.SetSprintID(_SprintID.String)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _CompleteDate.Valid {
		t.SetCompleteDate(_CompleteDate.Int64)
	}
	if _EndDate.Valid {
		t.SetEndDate(_EndDate.Int64)
	}
	if _StartDate.Valid {
		t.SetStartDate(_StartDate.Int64)
	}
	if _State.Valid {
		t.SetState(_State.String)
	}
	if _Goal.Valid {
		t.SetGoal(_Goal.String)
	}
	if _InitialIssueIds.Valid {
		t.SetInitialIssueIds(_InitialIssueIds.String)
	}
	if _FinalIssueIds.Valid {
		t.SetFinalIssueIds(_FinalIssueIds.String)
	}
	if _FinalClosedIssueIds.Valid {
		t.SetFinalClosedIssueIds(_FinalClosedIssueIds.String)
	}
	if _InitialIssueCount.Valid {
		t.SetInitialIssueCount(int32(_InitialIssueCount.Int64))
	}
	if _FinalIssueCount.Valid {
		t.SetFinalIssueCount(int32(_FinalIssueCount.Int64))
	}
	if _ClosedIssueCount.Valid {
		t.SetClosedIssueCount(int32(_ClosedIssueCount.Int64))
	}
	if _AddedIssueCount.Valid {
		t.SetAddedIssueCount(int32(_AddedIssueCount.Int64))
	}
	if _RemovedIssueCount.Valid {
		t.SetRemovedIssueCount(int32(_RemovedIssueCount.Int64))
	}
	if _InitialIssuesClosed.Valid {
		t.SetInitialIssuesClosed(int32(_InitialIssuesClosed.Int64))
	}
	if _InitialIssuesInFinalCount.Valid {
		t.SetInitialIssuesInFinalCount(int32(_InitialIssuesInFinalCount.Int64))
	}
	return true, nil
}

// CountJiraProjectSprints will find the count of JiraProjectSprint records in the database
func CountJiraProjectSprints(ctx context.Context, db *sql.DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(JiraProjectSprintTableName),
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

// CountJiraProjectSprintsTx will find the count of JiraProjectSprint records in the database using the provided transaction
func CountJiraProjectSprintsTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(JiraProjectSprintTableName),
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

// DBCount will find the count of JiraProjectSprint records in the database
func (t *JiraProjectSprint) DBCount(ctx context.Context, db *sql.DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(JiraProjectSprintTableName),
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

// DBCountTx will find the count of JiraProjectSprint records in the database using the provided transaction
func (t *JiraProjectSprint) DBCountTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(JiraProjectSprintTableName),
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

// DBExists will return true if the JiraProjectSprint record exists in the database
func (t *JiraProjectSprint) DBExists(ctx context.Context, db *sql.DB) (bool, error) {
	q := "SELECT `id` FROM `jira_project_sprint` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := db.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// DBExistsTx will return true if the JiraProjectSprint record exists in the database using the provided transaction
func (t *JiraProjectSprint) DBExistsTx(ctx context.Context, tx *sql.Tx) (bool, error) {
	q := "SELECT `id` FROM `jira_project_sprint` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := tx.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// PrimaryKeyColumn returns the column name for the primary key
func (t *JiraProjectSprint) PrimaryKeyColumn() string {
	return JiraProjectSprintColumnID
}

// PrimaryKeyColumnType returns the primary key column Go type as a string
func (t *JiraProjectSprint) PrimaryKeyColumnType() string {
	return "string"
}

// PrimaryKey returns the primary key column value
func (t *JiraProjectSprint) PrimaryKey() interface{} {
	return t.ID
}
