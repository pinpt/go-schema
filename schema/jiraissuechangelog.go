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

var _ Model = (*JiraIssueChangeLog)(nil)
var _ CSVWriter = (*JiraIssueChangeLog)(nil)
var _ JSONWriter = (*JiraIssueChangeLog)(nil)

// JiraIssueChangeLogTableName is the name of the table in SQL
const JiraIssueChangeLogTableName = "jira_issue_change_log"

var JiraIssueChangeLogColumns = []string{
	"id",
	"user_id",
	"assignee_id",
	"created_at",
	"field",
	"field_type",
	"from",
	"from_string",
	"to",
	"to_string",
	"ref_id",
	"customer_id",
	"issue_id",
	"project_id",
	"ordinal",
}

// JiraIssueChangeLog table
type JiraIssueChangeLog struct {
	AssigneeID *string `json:"assignee_id,omitempty"`
	CreatedAt  *int64  `json:"created_at,omitempty"`
	CustomerID string  `json:"customer_id"`
	Field      string  `json:"field"`
	FieldType  *string `json:"field_type,omitempty"`
	From       *string `json:"from,omitempty"`
	FromString *string `json:"from_string,omitempty"`
	ID         string  `json:"id"`
	IssueID    string  `json:"issue_id"`
	Ordinal    int32   `json:"ordinal"`
	ProjectID  string  `json:"project_id"`
	RefID      string  `json:"ref_id"`
	To         *string `json:"to,omitempty"`
	ToString   *string `json:"to_string,omitempty"`
	UserID     *string `json:"user_id,omitempty"`
}

// TableName returns the SQL table name for JiraIssueChangeLog and satifies the Model interface
func (t *JiraIssueChangeLog) TableName() string {
	return JiraIssueChangeLogTableName
}

// ToCSV will serialize the JiraIssueChangeLog instance to a CSV compatible array of strings
func (t *JiraIssueChangeLog) ToCSV() []string {
	return []string{
		t.ID,
		toCSVString(t.UserID),
		toCSVString(t.AssigneeID),
		toCSVString(t.CreatedAt),
		t.Field,
		toCSVString(t.FieldType),
		toCSVString(t.From),
		toCSVString(t.FromString),
		toCSVString(t.To),
		toCSVString(t.ToString),
		t.RefID,
		t.CustomerID,
		t.IssueID,
		t.ProjectID,
		toCSVString(t.Ordinal),
	}
}

// WriteCSV will serialize the JiraIssueChangeLog instance to the writer as CSV and satisfies the CSVWriter interface
func (t *JiraIssueChangeLog) WriteCSV(w *csv.Writer) error {
	return w.Write(t.ToCSV())
}

// WriteJSON will serialize the JiraIssueChangeLog instance to the writer as JSON and satisfies the JSONWriter interface
func (t *JiraIssueChangeLog) WriteJSON(w io.Writer, indent ...bool) error {
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

// NewJiraIssueChangeLogReader creates a JSON reader which can read in JiraIssueChangeLog objects serialized as JSON either as an array, single object or json new lines
// and writes each JiraIssueChangeLog to the channel provided
func NewJiraIssueChangeLogReader(r io.Reader, ch chan<- JiraIssueChangeLog) error {
	return orm.Deserialize(r, func(buf json.RawMessage) error {
		dec := json.NewDecoder(bytes.NewBuffer(buf))
		e := JiraIssueChangeLog{}
		if err := dec.Decode(&e); err != nil {
			return err
		}
		ch <- e
		return nil
	})
}

// NewCSVJiraIssueChangeLogReaderDir will read the reader as CSV and emit each record to the channel provided
func NewCSVJiraIssueChangeLogReader(r io.Reader, ch chan<- JiraIssueChangeLog) error {
	cr := csv.NewReader(r)
	for {
		record, err := cr.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		ch <- JiraIssueChangeLog{
			ID:         record[0],
			UserID:     fromStringPointer(record[1]),
			AssigneeID: fromStringPointer(record[2]),
			CreatedAt:  fromCSVInt64Pointer(record[3]),
			Field:      record[4],
			FieldType:  fromStringPointer(record[5]),
			From:       fromStringPointer(record[6]),
			FromString: fromStringPointer(record[7]),
			To:         fromStringPointer(record[8]),
			ToString:   fromStringPointer(record[9]),
			RefID:      record[10],
			CustomerID: record[11],
			IssueID:    record[12],
			ProjectID:  record[13],
			Ordinal:    fromCSVInt32(record[14]),
		}
	}
	return nil
}

// NewCSVJiraIssueChangeLogReaderFile will read the file as a CSV and emit each record to the channel provided
func NewCSVJiraIssueChangeLogReaderFile(fp string, ch chan<- JiraIssueChangeLog) error {
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
	return NewCSVJiraIssueChangeLogReader(fc, ch)
}

// NewCSVJiraIssueChangeLogReaderDir will read the jira_issue_change_log.csv.gz file as a CSV and emit each record to the channel provided
func NewCSVJiraIssueChangeLogReaderDir(dir string, ch chan<- JiraIssueChangeLog) error {
	return NewCSVJiraIssueChangeLogReaderFile(filepath.Join(dir, "jira_issue_change_log.csv.gz"), ch)
}

// JiraIssueChangeLogCSVDeduper is a function callback which takes the existing value (a) and the new value (b)
// and the return value should be the one to use (or a new one, if applicable). return nil
// to skip processing of this record
type JiraIssueChangeLogCSVDeduper func(a JiraIssueChangeLog, b JiraIssueChangeLog) *JiraIssueChangeLog

// JiraIssueChangeLogCSVDedupeDisabled is set on whether the CSV writer should de-dupe values by key
var JiraIssueChangeLogCSVDedupeDisabled bool

// NewJiraIssueChangeLogCSVWriterSize creates a batch writer that will write each JiraIssueChangeLog into a CSV file
func NewJiraIssueChangeLogCSVWriterSize(w io.Writer, size int, dedupers ...JiraIssueChangeLogCSVDeduper) (chan JiraIssueChangeLog, chan bool, error) {
	cw := csv.NewWriter(w)
	ch := make(chan JiraIssueChangeLog, size)
	done := make(chan bool)
	go func() {
		defer func() { done <- true }()
		dodedupe := !JiraIssueChangeLogCSVDedupeDisabled
		var kv map[string]*JiraIssueChangeLog
		var deduper JiraIssueChangeLogCSVDeduper
		if dedupers != nil && len(dedupers) > 0 {
			deduper = dedupers[0]
			dodedupe = true
		}
		if dodedupe {
			kv = make(map[string]*JiraIssueChangeLog)
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

// JiraIssueChangeLogCSVDefaultSize is the default channel buffer size if not provided
var JiraIssueChangeLogCSVDefaultSize = 100

// NewJiraIssueChangeLogCSVWriter creates a batch writer that will write each JiraIssueChangeLog into a CSV file
func NewJiraIssueChangeLogCSVWriter(w io.Writer, dedupers ...JiraIssueChangeLogCSVDeduper) (chan JiraIssueChangeLog, chan bool, error) {
	return NewJiraIssueChangeLogCSVWriterSize(w, JiraIssueChangeLogCSVDefaultSize, dedupers...)
}

// NewJiraIssueChangeLogCSVWriterDir creates a batch writer that will write each JiraIssueChangeLog into a CSV file named jira_issue_change_log.csv.gz in dir
func NewJiraIssueChangeLogCSVWriterDir(dir string, dedupers ...JiraIssueChangeLogCSVDeduper) (chan JiraIssueChangeLog, chan bool, error) {
	return NewJiraIssueChangeLogCSVWriterFile(filepath.Join(dir, "jira_issue_change_log.csv.gz"), dedupers...)
}

// NewJiraIssueChangeLogCSVWriterFile creates a batch writer that will write each JiraIssueChangeLog into a CSV file
func NewJiraIssueChangeLogCSVWriterFile(fn string, dedupers ...JiraIssueChangeLogCSVDeduper) (chan JiraIssueChangeLog, chan bool, error) {
	f, err := os.Create(fn)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening CSV file %s. %v", fn, err)
	}
	var fc io.WriteCloser = f
	if filepath.Ext(fn) == ".gz" {
		w, _ := gzip.NewWriterLevel(f, gzip.BestCompression)
		fc = w
	}
	ch, done, err := NewJiraIssueChangeLogCSVWriter(fc, dedupers...)
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

type JiraIssueChangeLogDBAction func(ctx context.Context, db DB, record JiraIssueChangeLog) error

// NewJiraIssueChangeLogDBWriterSize creates a DB writer that will write each issue into the DB
func NewJiraIssueChangeLogDBWriterSize(ctx context.Context, db DB, errors chan<- error, size int, actions ...JiraIssueChangeLogDBAction) (chan JiraIssueChangeLog, chan bool, error) {
	ch := make(chan JiraIssueChangeLog, size)
	done := make(chan bool)
	var action JiraIssueChangeLogDBAction
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

// NewJiraIssueChangeLogDBWriter creates a DB writer that will write each issue into the DB
func NewJiraIssueChangeLogDBWriter(ctx context.Context, db DB, errors chan<- error, actions ...JiraIssueChangeLogDBAction) (chan JiraIssueChangeLog, chan bool, error) {
	return NewJiraIssueChangeLogDBWriterSize(ctx, db, errors, 100, actions...)
}

// JiraIssueChangeLogColumnID is the ID SQL column name for the JiraIssueChangeLog table
const JiraIssueChangeLogColumnID = "id"

// JiraIssueChangeLogEscapedColumnID is the escaped ID SQL column name for the JiraIssueChangeLog table
const JiraIssueChangeLogEscapedColumnID = "`id`"

// JiraIssueChangeLogColumnUserID is the UserID SQL column name for the JiraIssueChangeLog table
const JiraIssueChangeLogColumnUserID = "user_id"

// JiraIssueChangeLogEscapedColumnUserID is the escaped UserID SQL column name for the JiraIssueChangeLog table
const JiraIssueChangeLogEscapedColumnUserID = "`user_id`"

// JiraIssueChangeLogColumnAssigneeID is the AssigneeID SQL column name for the JiraIssueChangeLog table
const JiraIssueChangeLogColumnAssigneeID = "assignee_id"

// JiraIssueChangeLogEscapedColumnAssigneeID is the escaped AssigneeID SQL column name for the JiraIssueChangeLog table
const JiraIssueChangeLogEscapedColumnAssigneeID = "`assignee_id`"

// JiraIssueChangeLogColumnCreatedAt is the CreatedAt SQL column name for the JiraIssueChangeLog table
const JiraIssueChangeLogColumnCreatedAt = "created_at"

// JiraIssueChangeLogEscapedColumnCreatedAt is the escaped CreatedAt SQL column name for the JiraIssueChangeLog table
const JiraIssueChangeLogEscapedColumnCreatedAt = "`created_at`"

// JiraIssueChangeLogColumnField is the Field SQL column name for the JiraIssueChangeLog table
const JiraIssueChangeLogColumnField = "field"

// JiraIssueChangeLogEscapedColumnField is the escaped Field SQL column name for the JiraIssueChangeLog table
const JiraIssueChangeLogEscapedColumnField = "`field`"

// JiraIssueChangeLogColumnFieldType is the FieldType SQL column name for the JiraIssueChangeLog table
const JiraIssueChangeLogColumnFieldType = "field_type"

// JiraIssueChangeLogEscapedColumnFieldType is the escaped FieldType SQL column name for the JiraIssueChangeLog table
const JiraIssueChangeLogEscapedColumnFieldType = "`field_type`"

// JiraIssueChangeLogColumnFrom is the From SQL column name for the JiraIssueChangeLog table
const JiraIssueChangeLogColumnFrom = "from"

// JiraIssueChangeLogEscapedColumnFrom is the escaped From SQL column name for the JiraIssueChangeLog table
const JiraIssueChangeLogEscapedColumnFrom = "`from`"

// JiraIssueChangeLogColumnFromString is the FromString SQL column name for the JiraIssueChangeLog table
const JiraIssueChangeLogColumnFromString = "from_string"

// JiraIssueChangeLogEscapedColumnFromString is the escaped FromString SQL column name for the JiraIssueChangeLog table
const JiraIssueChangeLogEscapedColumnFromString = "`from_string`"

// JiraIssueChangeLogColumnTo is the To SQL column name for the JiraIssueChangeLog table
const JiraIssueChangeLogColumnTo = "to"

// JiraIssueChangeLogEscapedColumnTo is the escaped To SQL column name for the JiraIssueChangeLog table
const JiraIssueChangeLogEscapedColumnTo = "`to`"

// JiraIssueChangeLogColumnToString is the ToString SQL column name for the JiraIssueChangeLog table
const JiraIssueChangeLogColumnToString = "to_string"

// JiraIssueChangeLogEscapedColumnToString is the escaped ToString SQL column name for the JiraIssueChangeLog table
const JiraIssueChangeLogEscapedColumnToString = "`to_string`"

// JiraIssueChangeLogColumnRefID is the RefID SQL column name for the JiraIssueChangeLog table
const JiraIssueChangeLogColumnRefID = "ref_id"

// JiraIssueChangeLogEscapedColumnRefID is the escaped RefID SQL column name for the JiraIssueChangeLog table
const JiraIssueChangeLogEscapedColumnRefID = "`ref_id`"

// JiraIssueChangeLogColumnCustomerID is the CustomerID SQL column name for the JiraIssueChangeLog table
const JiraIssueChangeLogColumnCustomerID = "customer_id"

// JiraIssueChangeLogEscapedColumnCustomerID is the escaped CustomerID SQL column name for the JiraIssueChangeLog table
const JiraIssueChangeLogEscapedColumnCustomerID = "`customer_id`"

// JiraIssueChangeLogColumnIssueID is the IssueID SQL column name for the JiraIssueChangeLog table
const JiraIssueChangeLogColumnIssueID = "issue_id"

// JiraIssueChangeLogEscapedColumnIssueID is the escaped IssueID SQL column name for the JiraIssueChangeLog table
const JiraIssueChangeLogEscapedColumnIssueID = "`issue_id`"

// JiraIssueChangeLogColumnProjectID is the ProjectID SQL column name for the JiraIssueChangeLog table
const JiraIssueChangeLogColumnProjectID = "project_id"

// JiraIssueChangeLogEscapedColumnProjectID is the escaped ProjectID SQL column name for the JiraIssueChangeLog table
const JiraIssueChangeLogEscapedColumnProjectID = "`project_id`"

// JiraIssueChangeLogColumnOrdinal is the Ordinal SQL column name for the JiraIssueChangeLog table
const JiraIssueChangeLogColumnOrdinal = "ordinal"

// JiraIssueChangeLogEscapedColumnOrdinal is the escaped Ordinal SQL column name for the JiraIssueChangeLog table
const JiraIssueChangeLogEscapedColumnOrdinal = "`ordinal`"

// GetID will return the JiraIssueChangeLog ID value
func (t *JiraIssueChangeLog) GetID() string {
	return t.ID
}

// SetID will set the JiraIssueChangeLog ID value
func (t *JiraIssueChangeLog) SetID(v string) {
	t.ID = v
}

// FindJiraIssueChangeLogByID will find a JiraIssueChangeLog by ID
func FindJiraIssueChangeLogByID(ctx context.Context, db DB, value string) (*JiraIssueChangeLog, error) {
	q := "SELECT * FROM `jira_issue_change_log` WHERE `id` = ?"
	var _ID sql.NullString
	var _UserID sql.NullString
	var _AssigneeID sql.NullString
	var _CreatedAt sql.NullInt64
	var _Field sql.NullString
	var _FieldType sql.NullString
	var _From sql.NullString
	var _FromString sql.NullString
	var _To sql.NullString
	var _ToString sql.NullString
	var _RefID sql.NullString
	var _CustomerID sql.NullString
	var _IssueID sql.NullString
	var _ProjectID sql.NullString
	var _Ordinal sql.NullInt64
	err := db.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_UserID,
		&_AssigneeID,
		&_CreatedAt,
		&_Field,
		&_FieldType,
		&_From,
		&_FromString,
		&_To,
		&_ToString,
		&_RefID,
		&_CustomerID,
		&_IssueID,
		&_ProjectID,
		&_Ordinal,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &JiraIssueChangeLog{}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _UserID.Valid {
		t.SetUserID(_UserID.String)
	}
	if _AssigneeID.Valid {
		t.SetAssigneeID(_AssigneeID.String)
	}
	if _CreatedAt.Valid {
		t.SetCreatedAt(_CreatedAt.Int64)
	}
	if _Field.Valid {
		t.SetField(_Field.String)
	}
	if _FieldType.Valid {
		t.SetFieldType(_FieldType.String)
	}
	if _From.Valid {
		t.SetFrom(_From.String)
	}
	if _FromString.Valid {
		t.SetFromString(_FromString.String)
	}
	if _To.Valid {
		t.SetTo(_To.String)
	}
	if _ToString.Valid {
		t.SetToString(_ToString.String)
	}
	if _RefID.Valid {
		t.SetRefID(_RefID.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _IssueID.Valid {
		t.SetIssueID(_IssueID.String)
	}
	if _ProjectID.Valid {
		t.SetProjectID(_ProjectID.String)
	}
	if _Ordinal.Valid {
		t.SetOrdinal(int32(_Ordinal.Int64))
	}
	return t, nil
}

// FindJiraIssueChangeLogByIDTx will find a JiraIssueChangeLog by ID using the provided transaction
func FindJiraIssueChangeLogByIDTx(ctx context.Context, tx Tx, value string) (*JiraIssueChangeLog, error) {
	q := "SELECT * FROM `jira_issue_change_log` WHERE `id` = ?"
	var _ID sql.NullString
	var _UserID sql.NullString
	var _AssigneeID sql.NullString
	var _CreatedAt sql.NullInt64
	var _Field sql.NullString
	var _FieldType sql.NullString
	var _From sql.NullString
	var _FromString sql.NullString
	var _To sql.NullString
	var _ToString sql.NullString
	var _RefID sql.NullString
	var _CustomerID sql.NullString
	var _IssueID sql.NullString
	var _ProjectID sql.NullString
	var _Ordinal sql.NullInt64
	err := tx.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_UserID,
		&_AssigneeID,
		&_CreatedAt,
		&_Field,
		&_FieldType,
		&_From,
		&_FromString,
		&_To,
		&_ToString,
		&_RefID,
		&_CustomerID,
		&_IssueID,
		&_ProjectID,
		&_Ordinal,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &JiraIssueChangeLog{}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _UserID.Valid {
		t.SetUserID(_UserID.String)
	}
	if _AssigneeID.Valid {
		t.SetAssigneeID(_AssigneeID.String)
	}
	if _CreatedAt.Valid {
		t.SetCreatedAt(_CreatedAt.Int64)
	}
	if _Field.Valid {
		t.SetField(_Field.String)
	}
	if _FieldType.Valid {
		t.SetFieldType(_FieldType.String)
	}
	if _From.Valid {
		t.SetFrom(_From.String)
	}
	if _FromString.Valid {
		t.SetFromString(_FromString.String)
	}
	if _To.Valid {
		t.SetTo(_To.String)
	}
	if _ToString.Valid {
		t.SetToString(_ToString.String)
	}
	if _RefID.Valid {
		t.SetRefID(_RefID.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _IssueID.Valid {
		t.SetIssueID(_IssueID.String)
	}
	if _ProjectID.Valid {
		t.SetProjectID(_ProjectID.String)
	}
	if _Ordinal.Valid {
		t.SetOrdinal(int32(_Ordinal.Int64))
	}
	return t, nil
}

// GetUserID will return the JiraIssueChangeLog UserID value
func (t *JiraIssueChangeLog) GetUserID() string {
	if t.UserID == nil {
		return ""
	}
	return *t.UserID
}

// SetUserID will set the JiraIssueChangeLog UserID value
func (t *JiraIssueChangeLog) SetUserID(v string) {
	t.UserID = &v
}

// GetAssigneeID will return the JiraIssueChangeLog AssigneeID value
func (t *JiraIssueChangeLog) GetAssigneeID() string {
	if t.AssigneeID == nil {
		return ""
	}
	return *t.AssigneeID
}

// SetAssigneeID will set the JiraIssueChangeLog AssigneeID value
func (t *JiraIssueChangeLog) SetAssigneeID(v string) {
	t.AssigneeID = &v
}

// GetCreatedAt will return the JiraIssueChangeLog CreatedAt value
func (t *JiraIssueChangeLog) GetCreatedAt() int64 {
	if t.CreatedAt == nil {
		return int64(0)
	}
	return *t.CreatedAt
}

// SetCreatedAt will set the JiraIssueChangeLog CreatedAt value
func (t *JiraIssueChangeLog) SetCreatedAt(v int64) {
	t.CreatedAt = &v
}

// GetField will return the JiraIssueChangeLog Field value
func (t *JiraIssueChangeLog) GetField() string {
	return t.Field
}

// SetField will set the JiraIssueChangeLog Field value
func (t *JiraIssueChangeLog) SetField(v string) {
	t.Field = v
}

// FindJiraIssueChangeLogsByField will find all JiraIssueChangeLogs by the Field value
func FindJiraIssueChangeLogsByField(ctx context.Context, db DB, value string) ([]*JiraIssueChangeLog, error) {
	q := "SELECT * FROM `jira_issue_change_log` WHERE `field` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraIssueChangeLog, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _UserID sql.NullString
		var _AssigneeID sql.NullString
		var _CreatedAt sql.NullInt64
		var _Field sql.NullString
		var _FieldType sql.NullString
		var _From sql.NullString
		var _FromString sql.NullString
		var _To sql.NullString
		var _ToString sql.NullString
		var _RefID sql.NullString
		var _CustomerID sql.NullString
		var _IssueID sql.NullString
		var _ProjectID sql.NullString
		var _Ordinal sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_UserID,
			&_AssigneeID,
			&_CreatedAt,
			&_Field,
			&_FieldType,
			&_From,
			&_FromString,
			&_To,
			&_ToString,
			&_RefID,
			&_CustomerID,
			&_IssueID,
			&_ProjectID,
			&_Ordinal,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraIssueChangeLog{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _AssigneeID.Valid {
			t.SetAssigneeID(_AssigneeID.String)
		}
		if _CreatedAt.Valid {
			t.SetCreatedAt(_CreatedAt.Int64)
		}
		if _Field.Valid {
			t.SetField(_Field.String)
		}
		if _FieldType.Valid {
			t.SetFieldType(_FieldType.String)
		}
		if _From.Valid {
			t.SetFrom(_From.String)
		}
		if _FromString.Valid {
			t.SetFromString(_FromString.String)
		}
		if _To.Valid {
			t.SetTo(_To.String)
		}
		if _ToString.Valid {
			t.SetToString(_ToString.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _Ordinal.Valid {
			t.SetOrdinal(int32(_Ordinal.Int64))
		}
		results = append(results, t)
	}
	return results, nil
}

// FindJiraIssueChangeLogsByFieldTx will find all JiraIssueChangeLogs by the Field value using the provided transaction
func FindJiraIssueChangeLogsByFieldTx(ctx context.Context, tx Tx, value string) ([]*JiraIssueChangeLog, error) {
	q := "SELECT * FROM `jira_issue_change_log` WHERE `field` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraIssueChangeLog, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _UserID sql.NullString
		var _AssigneeID sql.NullString
		var _CreatedAt sql.NullInt64
		var _Field sql.NullString
		var _FieldType sql.NullString
		var _From sql.NullString
		var _FromString sql.NullString
		var _To sql.NullString
		var _ToString sql.NullString
		var _RefID sql.NullString
		var _CustomerID sql.NullString
		var _IssueID sql.NullString
		var _ProjectID sql.NullString
		var _Ordinal sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_UserID,
			&_AssigneeID,
			&_CreatedAt,
			&_Field,
			&_FieldType,
			&_From,
			&_FromString,
			&_To,
			&_ToString,
			&_RefID,
			&_CustomerID,
			&_IssueID,
			&_ProjectID,
			&_Ordinal,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraIssueChangeLog{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _AssigneeID.Valid {
			t.SetAssigneeID(_AssigneeID.String)
		}
		if _CreatedAt.Valid {
			t.SetCreatedAt(_CreatedAt.Int64)
		}
		if _Field.Valid {
			t.SetField(_Field.String)
		}
		if _FieldType.Valid {
			t.SetFieldType(_FieldType.String)
		}
		if _From.Valid {
			t.SetFrom(_From.String)
		}
		if _FromString.Valid {
			t.SetFromString(_FromString.String)
		}
		if _To.Valid {
			t.SetTo(_To.String)
		}
		if _ToString.Valid {
			t.SetToString(_ToString.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _Ordinal.Valid {
			t.SetOrdinal(int32(_Ordinal.Int64))
		}
		results = append(results, t)
	}
	return results, nil
}

// GetFieldType will return the JiraIssueChangeLog FieldType value
func (t *JiraIssueChangeLog) GetFieldType() string {
	if t.FieldType == nil {
		return ""
	}
	return *t.FieldType
}

// SetFieldType will set the JiraIssueChangeLog FieldType value
func (t *JiraIssueChangeLog) SetFieldType(v string) {
	t.FieldType = &v
}

// GetFrom will return the JiraIssueChangeLog From value
func (t *JiraIssueChangeLog) GetFrom() string {
	if t.From == nil {
		return ""
	}
	return *t.From
}

// SetFrom will set the JiraIssueChangeLog From value
func (t *JiraIssueChangeLog) SetFrom(v string) {
	t.From = &v
}

// GetFromString will return the JiraIssueChangeLog FromString value
func (t *JiraIssueChangeLog) GetFromString() string {
	if t.FromString == nil {
		return ""
	}
	return *t.FromString
}

// SetFromString will set the JiraIssueChangeLog FromString value
func (t *JiraIssueChangeLog) SetFromString(v string) {
	t.FromString = &v
}

// GetTo will return the JiraIssueChangeLog To value
func (t *JiraIssueChangeLog) GetTo() string {
	if t.To == nil {
		return ""
	}
	return *t.To
}

// SetTo will set the JiraIssueChangeLog To value
func (t *JiraIssueChangeLog) SetTo(v string) {
	t.To = &v
}

// GetToString will return the JiraIssueChangeLog ToString value
func (t *JiraIssueChangeLog) GetToString() string {
	if t.ToString == nil {
		return ""
	}
	return *t.ToString
}

// SetToString will set the JiraIssueChangeLog ToString value
func (t *JiraIssueChangeLog) SetToString(v string) {
	t.ToString = &v
}

// GetRefID will return the JiraIssueChangeLog RefID value
func (t *JiraIssueChangeLog) GetRefID() string {
	return t.RefID
}

// SetRefID will set the JiraIssueChangeLog RefID value
func (t *JiraIssueChangeLog) SetRefID(v string) {
	t.RefID = v
}

// GetCustomerID will return the JiraIssueChangeLog CustomerID value
func (t *JiraIssueChangeLog) GetCustomerID() string {
	return t.CustomerID
}

// SetCustomerID will set the JiraIssueChangeLog CustomerID value
func (t *JiraIssueChangeLog) SetCustomerID(v string) {
	t.CustomerID = v
}

// FindJiraIssueChangeLogsByCustomerID will find all JiraIssueChangeLogs by the CustomerID value
func FindJiraIssueChangeLogsByCustomerID(ctx context.Context, db DB, value string) ([]*JiraIssueChangeLog, error) {
	q := "SELECT * FROM `jira_issue_change_log` WHERE `customer_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraIssueChangeLog, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _UserID sql.NullString
		var _AssigneeID sql.NullString
		var _CreatedAt sql.NullInt64
		var _Field sql.NullString
		var _FieldType sql.NullString
		var _From sql.NullString
		var _FromString sql.NullString
		var _To sql.NullString
		var _ToString sql.NullString
		var _RefID sql.NullString
		var _CustomerID sql.NullString
		var _IssueID sql.NullString
		var _ProjectID sql.NullString
		var _Ordinal sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_UserID,
			&_AssigneeID,
			&_CreatedAt,
			&_Field,
			&_FieldType,
			&_From,
			&_FromString,
			&_To,
			&_ToString,
			&_RefID,
			&_CustomerID,
			&_IssueID,
			&_ProjectID,
			&_Ordinal,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraIssueChangeLog{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _AssigneeID.Valid {
			t.SetAssigneeID(_AssigneeID.String)
		}
		if _CreatedAt.Valid {
			t.SetCreatedAt(_CreatedAt.Int64)
		}
		if _Field.Valid {
			t.SetField(_Field.String)
		}
		if _FieldType.Valid {
			t.SetFieldType(_FieldType.String)
		}
		if _From.Valid {
			t.SetFrom(_From.String)
		}
		if _FromString.Valid {
			t.SetFromString(_FromString.String)
		}
		if _To.Valid {
			t.SetTo(_To.String)
		}
		if _ToString.Valid {
			t.SetToString(_ToString.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _Ordinal.Valid {
			t.SetOrdinal(int32(_Ordinal.Int64))
		}
		results = append(results, t)
	}
	return results, nil
}

// FindJiraIssueChangeLogsByCustomerIDTx will find all JiraIssueChangeLogs by the CustomerID value using the provided transaction
func FindJiraIssueChangeLogsByCustomerIDTx(ctx context.Context, tx Tx, value string) ([]*JiraIssueChangeLog, error) {
	q := "SELECT * FROM `jira_issue_change_log` WHERE `customer_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraIssueChangeLog, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _UserID sql.NullString
		var _AssigneeID sql.NullString
		var _CreatedAt sql.NullInt64
		var _Field sql.NullString
		var _FieldType sql.NullString
		var _From sql.NullString
		var _FromString sql.NullString
		var _To sql.NullString
		var _ToString sql.NullString
		var _RefID sql.NullString
		var _CustomerID sql.NullString
		var _IssueID sql.NullString
		var _ProjectID sql.NullString
		var _Ordinal sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_UserID,
			&_AssigneeID,
			&_CreatedAt,
			&_Field,
			&_FieldType,
			&_From,
			&_FromString,
			&_To,
			&_ToString,
			&_RefID,
			&_CustomerID,
			&_IssueID,
			&_ProjectID,
			&_Ordinal,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraIssueChangeLog{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _AssigneeID.Valid {
			t.SetAssigneeID(_AssigneeID.String)
		}
		if _CreatedAt.Valid {
			t.SetCreatedAt(_CreatedAt.Int64)
		}
		if _Field.Valid {
			t.SetField(_Field.String)
		}
		if _FieldType.Valid {
			t.SetFieldType(_FieldType.String)
		}
		if _From.Valid {
			t.SetFrom(_From.String)
		}
		if _FromString.Valid {
			t.SetFromString(_FromString.String)
		}
		if _To.Valid {
			t.SetTo(_To.String)
		}
		if _ToString.Valid {
			t.SetToString(_ToString.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _Ordinal.Valid {
			t.SetOrdinal(int32(_Ordinal.Int64))
		}
		results = append(results, t)
	}
	return results, nil
}

// GetIssueID will return the JiraIssueChangeLog IssueID value
func (t *JiraIssueChangeLog) GetIssueID() string {
	return t.IssueID
}

// SetIssueID will set the JiraIssueChangeLog IssueID value
func (t *JiraIssueChangeLog) SetIssueID(v string) {
	t.IssueID = v
}

// FindJiraIssueChangeLogsByIssueID will find all JiraIssueChangeLogs by the IssueID value
func FindJiraIssueChangeLogsByIssueID(ctx context.Context, db DB, value string) ([]*JiraIssueChangeLog, error) {
	q := "SELECT * FROM `jira_issue_change_log` WHERE `issue_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraIssueChangeLog, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _UserID sql.NullString
		var _AssigneeID sql.NullString
		var _CreatedAt sql.NullInt64
		var _Field sql.NullString
		var _FieldType sql.NullString
		var _From sql.NullString
		var _FromString sql.NullString
		var _To sql.NullString
		var _ToString sql.NullString
		var _RefID sql.NullString
		var _CustomerID sql.NullString
		var _IssueID sql.NullString
		var _ProjectID sql.NullString
		var _Ordinal sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_UserID,
			&_AssigneeID,
			&_CreatedAt,
			&_Field,
			&_FieldType,
			&_From,
			&_FromString,
			&_To,
			&_ToString,
			&_RefID,
			&_CustomerID,
			&_IssueID,
			&_ProjectID,
			&_Ordinal,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraIssueChangeLog{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _AssigneeID.Valid {
			t.SetAssigneeID(_AssigneeID.String)
		}
		if _CreatedAt.Valid {
			t.SetCreatedAt(_CreatedAt.Int64)
		}
		if _Field.Valid {
			t.SetField(_Field.String)
		}
		if _FieldType.Valid {
			t.SetFieldType(_FieldType.String)
		}
		if _From.Valid {
			t.SetFrom(_From.String)
		}
		if _FromString.Valid {
			t.SetFromString(_FromString.String)
		}
		if _To.Valid {
			t.SetTo(_To.String)
		}
		if _ToString.Valid {
			t.SetToString(_ToString.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _Ordinal.Valid {
			t.SetOrdinal(int32(_Ordinal.Int64))
		}
		results = append(results, t)
	}
	return results, nil
}

// FindJiraIssueChangeLogsByIssueIDTx will find all JiraIssueChangeLogs by the IssueID value using the provided transaction
func FindJiraIssueChangeLogsByIssueIDTx(ctx context.Context, tx Tx, value string) ([]*JiraIssueChangeLog, error) {
	q := "SELECT * FROM `jira_issue_change_log` WHERE `issue_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraIssueChangeLog, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _UserID sql.NullString
		var _AssigneeID sql.NullString
		var _CreatedAt sql.NullInt64
		var _Field sql.NullString
		var _FieldType sql.NullString
		var _From sql.NullString
		var _FromString sql.NullString
		var _To sql.NullString
		var _ToString sql.NullString
		var _RefID sql.NullString
		var _CustomerID sql.NullString
		var _IssueID sql.NullString
		var _ProjectID sql.NullString
		var _Ordinal sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_UserID,
			&_AssigneeID,
			&_CreatedAt,
			&_Field,
			&_FieldType,
			&_From,
			&_FromString,
			&_To,
			&_ToString,
			&_RefID,
			&_CustomerID,
			&_IssueID,
			&_ProjectID,
			&_Ordinal,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraIssueChangeLog{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _AssigneeID.Valid {
			t.SetAssigneeID(_AssigneeID.String)
		}
		if _CreatedAt.Valid {
			t.SetCreatedAt(_CreatedAt.Int64)
		}
		if _Field.Valid {
			t.SetField(_Field.String)
		}
		if _FieldType.Valid {
			t.SetFieldType(_FieldType.String)
		}
		if _From.Valid {
			t.SetFrom(_From.String)
		}
		if _FromString.Valid {
			t.SetFromString(_FromString.String)
		}
		if _To.Valid {
			t.SetTo(_To.String)
		}
		if _ToString.Valid {
			t.SetToString(_ToString.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _Ordinal.Valid {
			t.SetOrdinal(int32(_Ordinal.Int64))
		}
		results = append(results, t)
	}
	return results, nil
}

// GetProjectID will return the JiraIssueChangeLog ProjectID value
func (t *JiraIssueChangeLog) GetProjectID() string {
	return t.ProjectID
}

// SetProjectID will set the JiraIssueChangeLog ProjectID value
func (t *JiraIssueChangeLog) SetProjectID(v string) {
	t.ProjectID = v
}

// GetOrdinal will return the JiraIssueChangeLog Ordinal value
func (t *JiraIssueChangeLog) GetOrdinal() int32 {
	return t.Ordinal
}

// SetOrdinal will set the JiraIssueChangeLog Ordinal value
func (t *JiraIssueChangeLog) SetOrdinal(v int32) {
	t.Ordinal = v
}

func (t *JiraIssueChangeLog) toTimestamp(value time.Time) *timestamp.Timestamp {
	ts, _ := ptypes.TimestampProto(value)
	return ts
}

// DBCreateJiraIssueChangeLogTable will create the JiraIssueChangeLog table
func DBCreateJiraIssueChangeLogTable(ctx context.Context, db DB) error {
	q := "CREATE TABLE `jira_issue_change_log` (`id` VARCHAR(64) NOT NULL PRIMARY KEY,`user_id`VARCHAR(64),`assignee_id` VARCHAR(64),`created_at`BIGINT UNSIGNED,`field` VARCHAR(255) NOT NULL,`field_type`TEXT,`from`LONGTEXT,`from_string` LONGTEXT,`to` LONGTEXT,`to_string` LONGTEXT,`ref_id` VARCHAR(64) NOT NULL,`customer_id` VARCHAR(64) NOT NULL,`issue_id` VARCHAR(64) NOT NULL,`project_id`VARCHAR(64) NOT NULL,`ordinal`INT NOT NULL,INDEX jira_issue_change_log_field_index (`field`),INDEX jira_issue_change_log_customer_id_index (`customer_id`),INDEX jira_issue_change_log_issue_id_index (`issue_id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBCreateJiraIssueChangeLogTableTx will create the JiraIssueChangeLog table using the provided transction
func DBCreateJiraIssueChangeLogTableTx(ctx context.Context, tx Tx) error {
	q := "CREATE TABLE `jira_issue_change_log` (`id` VARCHAR(64) NOT NULL PRIMARY KEY,`user_id`VARCHAR(64),`assignee_id` VARCHAR(64),`created_at`BIGINT UNSIGNED,`field` VARCHAR(255) NOT NULL,`field_type`TEXT,`from`LONGTEXT,`from_string` LONGTEXT,`to` LONGTEXT,`to_string` LONGTEXT,`ref_id` VARCHAR(64) NOT NULL,`customer_id` VARCHAR(64) NOT NULL,`issue_id` VARCHAR(64) NOT NULL,`project_id`VARCHAR(64) NOT NULL,`ordinal`INT NOT NULL,INDEX jira_issue_change_log_field_index (`field`),INDEX jira_issue_change_log_customer_id_index (`customer_id`),INDEX jira_issue_change_log_issue_id_index (`issue_id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// DBDropJiraIssueChangeLogTable will drop the JiraIssueChangeLog table
func DBDropJiraIssueChangeLogTable(ctx context.Context, db DB) error {
	q := "DROP TABLE IF EXISTS `jira_issue_change_log`"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBDropJiraIssueChangeLogTableTx will drop the JiraIssueChangeLog table using the provided transaction
func DBDropJiraIssueChangeLogTableTx(ctx context.Context, tx Tx) error {
	q := "DROP TABLE IF EXISTS `jira_issue_change_log`"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// DBCreate will create a new JiraIssueChangeLog record in the database
func (t *JiraIssueChangeLog) DBCreate(ctx context.Context, db DB) (sql.Result, error) {
	q := "INSERT INTO `jira_issue_change_log` (`jira_issue_change_log`.`id`,`jira_issue_change_log`.`user_id`,`jira_issue_change_log`.`assignee_id`,`jira_issue_change_log`.`created_at`,`jira_issue_change_log`.`field`,`jira_issue_change_log`.`field_type`,`jira_issue_change_log`.`from`,`jira_issue_change_log`.`from_string`,`jira_issue_change_log`.`to`,`jira_issue_change_log`.`to_string`,`jira_issue_change_log`.`ref_id`,`jira_issue_change_log`.`customer_id`,`jira_issue_change_log`.`issue_id`,`jira_issue_change_log`.`project_id`,`jira_issue_change_log`.`ordinal`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.UserID),
		orm.ToSQLString(t.AssigneeID),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLString(t.Field),
		orm.ToSQLString(t.FieldType),
		orm.ToSQLString(t.From),
		orm.ToSQLString(t.FromString),
		orm.ToSQLString(t.To),
		orm.ToSQLString(t.ToString),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.IssueID),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLInt64(t.Ordinal),
	)
}

// DBCreateTx will create a new JiraIssueChangeLog record in the database using the provided transaction
func (t *JiraIssueChangeLog) DBCreateTx(ctx context.Context, tx Tx) (sql.Result, error) {
	q := "INSERT INTO `jira_issue_change_log` (`jira_issue_change_log`.`id`,`jira_issue_change_log`.`user_id`,`jira_issue_change_log`.`assignee_id`,`jira_issue_change_log`.`created_at`,`jira_issue_change_log`.`field`,`jira_issue_change_log`.`field_type`,`jira_issue_change_log`.`from`,`jira_issue_change_log`.`from_string`,`jira_issue_change_log`.`to`,`jira_issue_change_log`.`to_string`,`jira_issue_change_log`.`ref_id`,`jira_issue_change_log`.`customer_id`,`jira_issue_change_log`.`issue_id`,`jira_issue_change_log`.`project_id`,`jira_issue_change_log`.`ordinal`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.UserID),
		orm.ToSQLString(t.AssigneeID),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLString(t.Field),
		orm.ToSQLString(t.FieldType),
		orm.ToSQLString(t.From),
		orm.ToSQLString(t.FromString),
		orm.ToSQLString(t.To),
		orm.ToSQLString(t.ToString),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.IssueID),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLInt64(t.Ordinal),
	)
}

// DBCreateIgnoreDuplicate will upsert the JiraIssueChangeLog record in the database
func (t *JiraIssueChangeLog) DBCreateIgnoreDuplicate(ctx context.Context, db DB) (sql.Result, error) {
	q := "INSERT INTO `jira_issue_change_log` (`jira_issue_change_log`.`id`,`jira_issue_change_log`.`user_id`,`jira_issue_change_log`.`assignee_id`,`jira_issue_change_log`.`created_at`,`jira_issue_change_log`.`field`,`jira_issue_change_log`.`field_type`,`jira_issue_change_log`.`from`,`jira_issue_change_log`.`from_string`,`jira_issue_change_log`.`to`,`jira_issue_change_log`.`to_string`,`jira_issue_change_log`.`ref_id`,`jira_issue_change_log`.`customer_id`,`jira_issue_change_log`.`issue_id`,`jira_issue_change_log`.`project_id`,`jira_issue_change_log`.`ordinal`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.UserID),
		orm.ToSQLString(t.AssigneeID),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLString(t.Field),
		orm.ToSQLString(t.FieldType),
		orm.ToSQLString(t.From),
		orm.ToSQLString(t.FromString),
		orm.ToSQLString(t.To),
		orm.ToSQLString(t.ToString),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.IssueID),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLInt64(t.Ordinal),
	)
}

// DBCreateIgnoreDuplicateTx will upsert the JiraIssueChangeLog record in the database using the provided transaction
func (t *JiraIssueChangeLog) DBCreateIgnoreDuplicateTx(ctx context.Context, tx Tx) (sql.Result, error) {
	q := "INSERT INTO `jira_issue_change_log` (`jira_issue_change_log`.`id`,`jira_issue_change_log`.`user_id`,`jira_issue_change_log`.`assignee_id`,`jira_issue_change_log`.`created_at`,`jira_issue_change_log`.`field`,`jira_issue_change_log`.`field_type`,`jira_issue_change_log`.`from`,`jira_issue_change_log`.`from_string`,`jira_issue_change_log`.`to`,`jira_issue_change_log`.`to_string`,`jira_issue_change_log`.`ref_id`,`jira_issue_change_log`.`customer_id`,`jira_issue_change_log`.`issue_id`,`jira_issue_change_log`.`project_id`,`jira_issue_change_log`.`ordinal`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.UserID),
		orm.ToSQLString(t.AssigneeID),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLString(t.Field),
		orm.ToSQLString(t.FieldType),
		orm.ToSQLString(t.From),
		orm.ToSQLString(t.FromString),
		orm.ToSQLString(t.To),
		orm.ToSQLString(t.ToString),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.IssueID),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLInt64(t.Ordinal),
	)
}

// DeleteAllJiraIssueChangeLogs deletes all JiraIssueChangeLog records in the database with optional filters
func DeleteAllJiraIssueChangeLogs(ctx context.Context, db DB, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(JiraIssueChangeLogTableName),
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

// DeleteAllJiraIssueChangeLogsTx deletes all JiraIssueChangeLog records in the database with optional filters using the provided transaction
func DeleteAllJiraIssueChangeLogsTx(ctx context.Context, tx Tx, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(JiraIssueChangeLogTableName),
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

// DBDelete will delete this JiraIssueChangeLog record in the database
func (t *JiraIssueChangeLog) DBDelete(ctx context.Context, db DB) (bool, error) {
	q := "DELETE FROM `jira_issue_change_log` WHERE `id` = ?"
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

// DBDeleteTx will delete this JiraIssueChangeLog record in the database using the provided transaction
func (t *JiraIssueChangeLog) DBDeleteTx(ctx context.Context, tx Tx) (bool, error) {
	q := "DELETE FROM `jira_issue_change_log` WHERE `id` = ?"
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

// DBUpdate will update the JiraIssueChangeLog record in the database
func (t *JiraIssueChangeLog) DBUpdate(ctx context.Context, db DB) (sql.Result, error) {
	q := "UPDATE `jira_issue_change_log` SET `user_id`=?,`assignee_id`=?,`created_at`=?,`field`=?,`field_type`=?,`from`=?,`from_string`=?,`to`=?,`to_string`=?,`ref_id`=?,`customer_id`=?,`issue_id`=?,`project_id`=?,`ordinal`=? WHERE `id`=?"
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.UserID),
		orm.ToSQLString(t.AssigneeID),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLString(t.Field),
		orm.ToSQLString(t.FieldType),
		orm.ToSQLString(t.From),
		orm.ToSQLString(t.FromString),
		orm.ToSQLString(t.To),
		orm.ToSQLString(t.ToString),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.IssueID),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLInt64(t.Ordinal),
		orm.ToSQLString(t.ID),
	)
}

// DBUpdateTx will update the JiraIssueChangeLog record in the database using the provided transaction
func (t *JiraIssueChangeLog) DBUpdateTx(ctx context.Context, tx Tx) (sql.Result, error) {
	q := "UPDATE `jira_issue_change_log` SET `user_id`=?,`assignee_id`=?,`created_at`=?,`field`=?,`field_type`=?,`from`=?,`from_string`=?,`to`=?,`to_string`=?,`ref_id`=?,`customer_id`=?,`issue_id`=?,`project_id`=?,`ordinal`=? WHERE `id`=?"
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.UserID),
		orm.ToSQLString(t.AssigneeID),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLString(t.Field),
		orm.ToSQLString(t.FieldType),
		orm.ToSQLString(t.From),
		orm.ToSQLString(t.FromString),
		orm.ToSQLString(t.To),
		orm.ToSQLString(t.ToString),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.IssueID),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLInt64(t.Ordinal),
		orm.ToSQLString(t.ID),
	)
}

// DBUpsert will upsert the JiraIssueChangeLog record in the database
func (t *JiraIssueChangeLog) DBUpsert(ctx context.Context, db DB, conditions ...interface{}) (bool, bool, error) {
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `jira_issue_change_log` (`jira_issue_change_log`.`id`,`jira_issue_change_log`.`user_id`,`jira_issue_change_log`.`assignee_id`,`jira_issue_change_log`.`created_at`,`jira_issue_change_log`.`field`,`jira_issue_change_log`.`field_type`,`jira_issue_change_log`.`from`,`jira_issue_change_log`.`from_string`,`jira_issue_change_log`.`to`,`jira_issue_change_log`.`to_string`,`jira_issue_change_log`.`ref_id`,`jira_issue_change_log`.`customer_id`,`jira_issue_change_log`.`issue_id`,`jira_issue_change_log`.`project_id`,`jira_issue_change_log`.`ordinal`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `jira_issue_change_log` (`jira_issue_change_log`.`id`,`jira_issue_change_log`.`user_id`,`jira_issue_change_log`.`assignee_id`,`jira_issue_change_log`.`created_at`,`jira_issue_change_log`.`field`,`jira_issue_change_log`.`field_type`,`jira_issue_change_log`.`from`,`jira_issue_change_log`.`from_string`,`jira_issue_change_log`.`to`,`jira_issue_change_log`.`to_string`,`jira_issue_change_log`.`ref_id`,`jira_issue_change_log`.`customer_id`,`jira_issue_change_log`.`issue_id`,`jira_issue_change_log`.`project_id`,`jira_issue_change_log`.`ordinal`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `user_id`=VALUES(`user_id`),`assignee_id`=VALUES(`assignee_id`),`created_at`=VALUES(`created_at`),`field`=VALUES(`field`),`field_type`=VALUES(`field_type`),`from`=VALUES(`from`),`from_string`=VALUES(`from_string`),`to`=VALUES(`to`),`to_string`=VALUES(`to_string`),`ref_id`=VALUES(`ref_id`),`customer_id`=VALUES(`customer_id`),`issue_id`=VALUES(`issue_id`),`project_id`=VALUES(`project_id`),`ordinal`=VALUES(`ordinal`)"
	}
	r, err := db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.UserID),
		orm.ToSQLString(t.AssigneeID),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLString(t.Field),
		orm.ToSQLString(t.FieldType),
		orm.ToSQLString(t.From),
		orm.ToSQLString(t.FromString),
		orm.ToSQLString(t.To),
		orm.ToSQLString(t.ToString),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.IssueID),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLInt64(t.Ordinal),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBUpsertTx will upsert the JiraIssueChangeLog record in the database using the provided transaction
func (t *JiraIssueChangeLog) DBUpsertTx(ctx context.Context, tx Tx, conditions ...interface{}) (bool, bool, error) {
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `jira_issue_change_log` (`jira_issue_change_log`.`id`,`jira_issue_change_log`.`user_id`,`jira_issue_change_log`.`assignee_id`,`jira_issue_change_log`.`created_at`,`jira_issue_change_log`.`field`,`jira_issue_change_log`.`field_type`,`jira_issue_change_log`.`from`,`jira_issue_change_log`.`from_string`,`jira_issue_change_log`.`to`,`jira_issue_change_log`.`to_string`,`jira_issue_change_log`.`ref_id`,`jira_issue_change_log`.`customer_id`,`jira_issue_change_log`.`issue_id`,`jira_issue_change_log`.`project_id`,`jira_issue_change_log`.`ordinal`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `jira_issue_change_log` (`jira_issue_change_log`.`id`,`jira_issue_change_log`.`user_id`,`jira_issue_change_log`.`assignee_id`,`jira_issue_change_log`.`created_at`,`jira_issue_change_log`.`field`,`jira_issue_change_log`.`field_type`,`jira_issue_change_log`.`from`,`jira_issue_change_log`.`from_string`,`jira_issue_change_log`.`to`,`jira_issue_change_log`.`to_string`,`jira_issue_change_log`.`ref_id`,`jira_issue_change_log`.`customer_id`,`jira_issue_change_log`.`issue_id`,`jira_issue_change_log`.`project_id`,`jira_issue_change_log`.`ordinal`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `user_id`=VALUES(`user_id`),`assignee_id`=VALUES(`assignee_id`),`created_at`=VALUES(`created_at`),`field`=VALUES(`field`),`field_type`=VALUES(`field_type`),`from`=VALUES(`from`),`from_string`=VALUES(`from_string`),`to`=VALUES(`to`),`to_string`=VALUES(`to_string`),`ref_id`=VALUES(`ref_id`),`customer_id`=VALUES(`customer_id`),`issue_id`=VALUES(`issue_id`),`project_id`=VALUES(`project_id`),`ordinal`=VALUES(`ordinal`)"
	}
	r, err := tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.UserID),
		orm.ToSQLString(t.AssigneeID),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLString(t.Field),
		orm.ToSQLString(t.FieldType),
		orm.ToSQLString(t.From),
		orm.ToSQLString(t.FromString),
		orm.ToSQLString(t.To),
		orm.ToSQLString(t.ToString),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.IssueID),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLInt64(t.Ordinal),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBFindOne will find a JiraIssueChangeLog record in the database with the primary key
func (t *JiraIssueChangeLog) DBFindOne(ctx context.Context, db DB, value string) (bool, error) {
	q := "SELECT * FROM `jira_issue_change_log` WHERE `id` = ? LIMIT 1"
	row := db.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _UserID sql.NullString
	var _AssigneeID sql.NullString
	var _CreatedAt sql.NullInt64
	var _Field sql.NullString
	var _FieldType sql.NullString
	var _From sql.NullString
	var _FromString sql.NullString
	var _To sql.NullString
	var _ToString sql.NullString
	var _RefID sql.NullString
	var _CustomerID sql.NullString
	var _IssueID sql.NullString
	var _ProjectID sql.NullString
	var _Ordinal sql.NullInt64
	err := row.Scan(
		&_ID,
		&_UserID,
		&_AssigneeID,
		&_CreatedAt,
		&_Field,
		&_FieldType,
		&_From,
		&_FromString,
		&_To,
		&_ToString,
		&_RefID,
		&_CustomerID,
		&_IssueID,
		&_ProjectID,
		&_Ordinal,
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
	if _UserID.Valid {
		t.SetUserID(_UserID.String)
	}
	if _AssigneeID.Valid {
		t.SetAssigneeID(_AssigneeID.String)
	}
	if _CreatedAt.Valid {
		t.SetCreatedAt(_CreatedAt.Int64)
	}
	if _Field.Valid {
		t.SetField(_Field.String)
	}
	if _FieldType.Valid {
		t.SetFieldType(_FieldType.String)
	}
	if _From.Valid {
		t.SetFrom(_From.String)
	}
	if _FromString.Valid {
		t.SetFromString(_FromString.String)
	}
	if _To.Valid {
		t.SetTo(_To.String)
	}
	if _ToString.Valid {
		t.SetToString(_ToString.String)
	}
	if _RefID.Valid {
		t.SetRefID(_RefID.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _IssueID.Valid {
		t.SetIssueID(_IssueID.String)
	}
	if _ProjectID.Valid {
		t.SetProjectID(_ProjectID.String)
	}
	if _Ordinal.Valid {
		t.SetOrdinal(int32(_Ordinal.Int64))
	}
	return true, nil
}

// DBFindOneTx will find a JiraIssueChangeLog record in the database with the primary key using the provided transaction
func (t *JiraIssueChangeLog) DBFindOneTx(ctx context.Context, tx Tx, value string) (bool, error) {
	q := "SELECT * FROM `jira_issue_change_log` WHERE `id` = ? LIMIT 1"
	row := tx.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _UserID sql.NullString
	var _AssigneeID sql.NullString
	var _CreatedAt sql.NullInt64
	var _Field sql.NullString
	var _FieldType sql.NullString
	var _From sql.NullString
	var _FromString sql.NullString
	var _To sql.NullString
	var _ToString sql.NullString
	var _RefID sql.NullString
	var _CustomerID sql.NullString
	var _IssueID sql.NullString
	var _ProjectID sql.NullString
	var _Ordinal sql.NullInt64
	err := row.Scan(
		&_ID,
		&_UserID,
		&_AssigneeID,
		&_CreatedAt,
		&_Field,
		&_FieldType,
		&_From,
		&_FromString,
		&_To,
		&_ToString,
		&_RefID,
		&_CustomerID,
		&_IssueID,
		&_ProjectID,
		&_Ordinal,
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
	if _UserID.Valid {
		t.SetUserID(_UserID.String)
	}
	if _AssigneeID.Valid {
		t.SetAssigneeID(_AssigneeID.String)
	}
	if _CreatedAt.Valid {
		t.SetCreatedAt(_CreatedAt.Int64)
	}
	if _Field.Valid {
		t.SetField(_Field.String)
	}
	if _FieldType.Valid {
		t.SetFieldType(_FieldType.String)
	}
	if _From.Valid {
		t.SetFrom(_From.String)
	}
	if _FromString.Valid {
		t.SetFromString(_FromString.String)
	}
	if _To.Valid {
		t.SetTo(_To.String)
	}
	if _ToString.Valid {
		t.SetToString(_ToString.String)
	}
	if _RefID.Valid {
		t.SetRefID(_RefID.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _IssueID.Valid {
		t.SetIssueID(_IssueID.String)
	}
	if _ProjectID.Valid {
		t.SetProjectID(_ProjectID.String)
	}
	if _Ordinal.Valid {
		t.SetOrdinal(int32(_Ordinal.Int64))
	}
	return true, nil
}

// FindJiraIssueChangeLogs will find a JiraIssueChangeLog record in the database with the provided parameters
func FindJiraIssueChangeLogs(ctx context.Context, db DB, _params ...interface{}) ([]*JiraIssueChangeLog, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("user_id"),
		orm.Column("assignee_id"),
		orm.Column("created_at"),
		orm.Column("field"),
		orm.Column("field_type"),
		orm.Column("from"),
		orm.Column("from_string"),
		orm.Column("to"),
		orm.Column("to_string"),
		orm.Column("ref_id"),
		orm.Column("customer_id"),
		orm.Column("issue_id"),
		orm.Column("project_id"),
		orm.Column("ordinal"),
		orm.Table(JiraIssueChangeLogTableName),
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
	results := make([]*JiraIssueChangeLog, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _UserID sql.NullString
		var _AssigneeID sql.NullString
		var _CreatedAt sql.NullInt64
		var _Field sql.NullString
		var _FieldType sql.NullString
		var _From sql.NullString
		var _FromString sql.NullString
		var _To sql.NullString
		var _ToString sql.NullString
		var _RefID sql.NullString
		var _CustomerID sql.NullString
		var _IssueID sql.NullString
		var _ProjectID sql.NullString
		var _Ordinal sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_UserID,
			&_AssigneeID,
			&_CreatedAt,
			&_Field,
			&_FieldType,
			&_From,
			&_FromString,
			&_To,
			&_ToString,
			&_RefID,
			&_CustomerID,
			&_IssueID,
			&_ProjectID,
			&_Ordinal,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraIssueChangeLog{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _AssigneeID.Valid {
			t.SetAssigneeID(_AssigneeID.String)
		}
		if _CreatedAt.Valid {
			t.SetCreatedAt(_CreatedAt.Int64)
		}
		if _Field.Valid {
			t.SetField(_Field.String)
		}
		if _FieldType.Valid {
			t.SetFieldType(_FieldType.String)
		}
		if _From.Valid {
			t.SetFrom(_From.String)
		}
		if _FromString.Valid {
			t.SetFromString(_FromString.String)
		}
		if _To.Valid {
			t.SetTo(_To.String)
		}
		if _ToString.Valid {
			t.SetToString(_ToString.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _Ordinal.Valid {
			t.SetOrdinal(int32(_Ordinal.Int64))
		}
		results = append(results, t)
	}
	return results, nil
}

// FindJiraIssueChangeLogsTx will find a JiraIssueChangeLog record in the database with the provided parameters using the provided transaction
func FindJiraIssueChangeLogsTx(ctx context.Context, tx Tx, _params ...interface{}) ([]*JiraIssueChangeLog, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("user_id"),
		orm.Column("assignee_id"),
		orm.Column("created_at"),
		orm.Column("field"),
		orm.Column("field_type"),
		orm.Column("from"),
		orm.Column("from_string"),
		orm.Column("to"),
		orm.Column("to_string"),
		orm.Column("ref_id"),
		orm.Column("customer_id"),
		orm.Column("issue_id"),
		orm.Column("project_id"),
		orm.Column("ordinal"),
		orm.Table(JiraIssueChangeLogTableName),
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
	results := make([]*JiraIssueChangeLog, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _UserID sql.NullString
		var _AssigneeID sql.NullString
		var _CreatedAt sql.NullInt64
		var _Field sql.NullString
		var _FieldType sql.NullString
		var _From sql.NullString
		var _FromString sql.NullString
		var _To sql.NullString
		var _ToString sql.NullString
		var _RefID sql.NullString
		var _CustomerID sql.NullString
		var _IssueID sql.NullString
		var _ProjectID sql.NullString
		var _Ordinal sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_UserID,
			&_AssigneeID,
			&_CreatedAt,
			&_Field,
			&_FieldType,
			&_From,
			&_FromString,
			&_To,
			&_ToString,
			&_RefID,
			&_CustomerID,
			&_IssueID,
			&_ProjectID,
			&_Ordinal,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraIssueChangeLog{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _AssigneeID.Valid {
			t.SetAssigneeID(_AssigneeID.String)
		}
		if _CreatedAt.Valid {
			t.SetCreatedAt(_CreatedAt.Int64)
		}
		if _Field.Valid {
			t.SetField(_Field.String)
		}
		if _FieldType.Valid {
			t.SetFieldType(_FieldType.String)
		}
		if _From.Valid {
			t.SetFrom(_From.String)
		}
		if _FromString.Valid {
			t.SetFromString(_FromString.String)
		}
		if _To.Valid {
			t.SetTo(_To.String)
		}
		if _ToString.Valid {
			t.SetToString(_ToString.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _Ordinal.Valid {
			t.SetOrdinal(int32(_Ordinal.Int64))
		}
		results = append(results, t)
	}
	return results, nil
}

// DBFind will find a JiraIssueChangeLog record in the database with the provided parameters
func (t *JiraIssueChangeLog) DBFind(ctx context.Context, db DB, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("user_id"),
		orm.Column("assignee_id"),
		orm.Column("created_at"),
		orm.Column("field"),
		orm.Column("field_type"),
		orm.Column("from"),
		orm.Column("from_string"),
		orm.Column("to"),
		orm.Column("to_string"),
		orm.Column("ref_id"),
		orm.Column("customer_id"),
		orm.Column("issue_id"),
		orm.Column("project_id"),
		orm.Column("ordinal"),
		orm.Table(JiraIssueChangeLogTableName),
	}
	if len(_params) > 0 {
		for _, param := range _params {
			params = append(params, param)
		}
	}
	q, p := orm.BuildQuery(params...)
	row := db.QueryRowContext(ctx, q, p...)
	var _ID sql.NullString
	var _UserID sql.NullString
	var _AssigneeID sql.NullString
	var _CreatedAt sql.NullInt64
	var _Field sql.NullString
	var _FieldType sql.NullString
	var _From sql.NullString
	var _FromString sql.NullString
	var _To sql.NullString
	var _ToString sql.NullString
	var _RefID sql.NullString
	var _CustomerID sql.NullString
	var _IssueID sql.NullString
	var _ProjectID sql.NullString
	var _Ordinal sql.NullInt64
	err := row.Scan(
		&_ID,
		&_UserID,
		&_AssigneeID,
		&_CreatedAt,
		&_Field,
		&_FieldType,
		&_From,
		&_FromString,
		&_To,
		&_ToString,
		&_RefID,
		&_CustomerID,
		&_IssueID,
		&_ProjectID,
		&_Ordinal,
	)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _UserID.Valid {
		t.SetUserID(_UserID.String)
	}
	if _AssigneeID.Valid {
		t.SetAssigneeID(_AssigneeID.String)
	}
	if _CreatedAt.Valid {
		t.SetCreatedAt(_CreatedAt.Int64)
	}
	if _Field.Valid {
		t.SetField(_Field.String)
	}
	if _FieldType.Valid {
		t.SetFieldType(_FieldType.String)
	}
	if _From.Valid {
		t.SetFrom(_From.String)
	}
	if _FromString.Valid {
		t.SetFromString(_FromString.String)
	}
	if _To.Valid {
		t.SetTo(_To.String)
	}
	if _ToString.Valid {
		t.SetToString(_ToString.String)
	}
	if _RefID.Valid {
		t.SetRefID(_RefID.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _IssueID.Valid {
		t.SetIssueID(_IssueID.String)
	}
	if _ProjectID.Valid {
		t.SetProjectID(_ProjectID.String)
	}
	if _Ordinal.Valid {
		t.SetOrdinal(int32(_Ordinal.Int64))
	}
	return true, nil
}

// DBFindTx will find a JiraIssueChangeLog record in the database with the provided parameters using the provided transaction
func (t *JiraIssueChangeLog) DBFindTx(ctx context.Context, tx Tx, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("user_id"),
		orm.Column("assignee_id"),
		orm.Column("created_at"),
		orm.Column("field"),
		orm.Column("field_type"),
		orm.Column("from"),
		orm.Column("from_string"),
		orm.Column("to"),
		orm.Column("to_string"),
		orm.Column("ref_id"),
		orm.Column("customer_id"),
		orm.Column("issue_id"),
		orm.Column("project_id"),
		orm.Column("ordinal"),
		orm.Table(JiraIssueChangeLogTableName),
	}
	if len(_params) > 0 {
		for _, param := range _params {
			params = append(params, param)
		}
	}
	q, p := orm.BuildQuery(params...)
	row := tx.QueryRowContext(ctx, q, p...)
	var _ID sql.NullString
	var _UserID sql.NullString
	var _AssigneeID sql.NullString
	var _CreatedAt sql.NullInt64
	var _Field sql.NullString
	var _FieldType sql.NullString
	var _From sql.NullString
	var _FromString sql.NullString
	var _To sql.NullString
	var _ToString sql.NullString
	var _RefID sql.NullString
	var _CustomerID sql.NullString
	var _IssueID sql.NullString
	var _ProjectID sql.NullString
	var _Ordinal sql.NullInt64
	err := row.Scan(
		&_ID,
		&_UserID,
		&_AssigneeID,
		&_CreatedAt,
		&_Field,
		&_FieldType,
		&_From,
		&_FromString,
		&_To,
		&_ToString,
		&_RefID,
		&_CustomerID,
		&_IssueID,
		&_ProjectID,
		&_Ordinal,
	)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _UserID.Valid {
		t.SetUserID(_UserID.String)
	}
	if _AssigneeID.Valid {
		t.SetAssigneeID(_AssigneeID.String)
	}
	if _CreatedAt.Valid {
		t.SetCreatedAt(_CreatedAt.Int64)
	}
	if _Field.Valid {
		t.SetField(_Field.String)
	}
	if _FieldType.Valid {
		t.SetFieldType(_FieldType.String)
	}
	if _From.Valid {
		t.SetFrom(_From.String)
	}
	if _FromString.Valid {
		t.SetFromString(_FromString.String)
	}
	if _To.Valid {
		t.SetTo(_To.String)
	}
	if _ToString.Valid {
		t.SetToString(_ToString.String)
	}
	if _RefID.Valid {
		t.SetRefID(_RefID.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _IssueID.Valid {
		t.SetIssueID(_IssueID.String)
	}
	if _ProjectID.Valid {
		t.SetProjectID(_ProjectID.String)
	}
	if _Ordinal.Valid {
		t.SetOrdinal(int32(_Ordinal.Int64))
	}
	return true, nil
}

// CountJiraIssueChangeLogs will find the count of JiraIssueChangeLog records in the database
func CountJiraIssueChangeLogs(ctx context.Context, db DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(JiraIssueChangeLogTableName),
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

// CountJiraIssueChangeLogsTx will find the count of JiraIssueChangeLog records in the database using the provided transaction
func CountJiraIssueChangeLogsTx(ctx context.Context, tx Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(JiraIssueChangeLogTableName),
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

// DBCount will find the count of JiraIssueChangeLog records in the database
func (t *JiraIssueChangeLog) DBCount(ctx context.Context, db DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(JiraIssueChangeLogTableName),
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

// DBCountTx will find the count of JiraIssueChangeLog records in the database using the provided transaction
func (t *JiraIssueChangeLog) DBCountTx(ctx context.Context, tx Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(JiraIssueChangeLogTableName),
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

// DBExists will return true if the JiraIssueChangeLog record exists in the database
func (t *JiraIssueChangeLog) DBExists(ctx context.Context, db DB) (bool, error) {
	q := "SELECT * FROM `jira_issue_change_log` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := db.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// DBExistsTx will return true if the JiraIssueChangeLog record exists in the database using the provided transaction
func (t *JiraIssueChangeLog) DBExistsTx(ctx context.Context, tx Tx) (bool, error) {
	q := "SELECT * FROM `jira_issue_change_log` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := tx.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// PrimaryKeyColumn returns the column name for the primary key
func (t *JiraIssueChangeLog) PrimaryKeyColumn() string {
	return JiraIssueChangeLogColumnID
}

// PrimaryKeyColumnType returns the primary key column Go type as a string
func (t *JiraIssueChangeLog) PrimaryKeyColumnType() string {
	return "string"
}

// PrimaryKey returns the primary key column value
func (t *JiraIssueChangeLog) PrimaryKey() interface{} {
	return t.ID
}
