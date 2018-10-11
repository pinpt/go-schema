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

var _ Model = (*JiraIssueProgress)(nil)
var _ CSVWriter = (*JiraIssueProgress)(nil)
var _ JSONWriter = (*JiraIssueProgress)(nil)
var _ Checksum = (*JiraIssueProgress)(nil)

// JiraIssueProgressTableName is the name of the table in SQL
const JiraIssueProgressTableName = "jira_issue_progress"

var JiraIssueProgressColumns = []string{
	"id",
	"checksum",
	"user_id",
	"issue_id",
	"start_date",
	"end_date",
	"duration",
	"customer_id",
	"ref_id",
}

// JiraIssueProgress table
type JiraIssueProgress struct {
	Checksum   *string `json:"checksum,omitempty"`
	CustomerID string  `json:"customer_id"`
	Duration   *int64  `json:"duration,omitempty"`
	EndDate    *int64  `json:"end_date,omitempty"`
	ID         string  `json:"id"`
	IssueID    string  `json:"issue_id"`
	RefID      string  `json:"ref_id"`
	StartDate  *int64  `json:"start_date,omitempty"`
	UserID     *string `json:"user_id,omitempty"`
}

// TableName returns the SQL table name for JiraIssueProgress and satifies the Model interface
func (t *JiraIssueProgress) TableName() string {
	return JiraIssueProgressTableName
}

// ToCSV will serialize the JiraIssueProgress instance to a CSV compatible array of strings
func (t *JiraIssueProgress) ToCSV() []string {
	return []string{
		t.ID,
		t.CalculateChecksum(),
		toCSVString(t.UserID),
		t.IssueID,
		toCSVString(t.StartDate),
		toCSVString(t.EndDate),
		toCSVString(t.Duration),
		t.CustomerID,
		t.RefID,
	}
}

// WriteCSV will serialize the JiraIssueProgress instance to the writer as CSV and satisfies the CSVWriter interface
func (t *JiraIssueProgress) WriteCSV(w *csv.Writer) error {
	return w.Write(t.ToCSV())
}

// WriteJSON will serialize the JiraIssueProgress instance to the writer as JSON and satisfies the JSONWriter interface
func (t *JiraIssueProgress) WriteJSON(w io.Writer, indent ...bool) error {
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

// NewJiraIssueProgressReader creates a JSON reader which can read in JiraIssueProgress objects serialized as JSON either as an array, single object or json new lines
// and writes each JiraIssueProgress to the channel provided
func NewJiraIssueProgressReader(r io.Reader, ch chan<- JiraIssueProgress) error {
	return orm.Deserialize(r, func(buf json.RawMessage) error {
		dec := json.NewDecoder(bytes.NewBuffer(buf))
		e := JiraIssueProgress{}
		if err := dec.Decode(&e); err != nil {
			return err
		}
		ch <- e
		return nil
	})
}

// NewCSVJiraIssueProgressReaderDir will read the reader as CSV and emit each record to the channel provided
func NewCSVJiraIssueProgressReader(r io.Reader, ch chan<- JiraIssueProgress) error {
	cr := csv.NewReader(r)
	for {
		record, err := cr.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		ch <- JiraIssueProgress{
			ID:         record[0],
			Checksum:   fromStringPointer(record[1]),
			UserID:     fromStringPointer(record[2]),
			IssueID:    record[3],
			StartDate:  fromCSVInt64Pointer(record[4]),
			EndDate:    fromCSVInt64Pointer(record[5]),
			Duration:   fromCSVInt64Pointer(record[6]),
			CustomerID: record[7],
			RefID:      record[8],
		}
	}
	return nil
}

// NewCSVJiraIssueProgressReaderFile will read the file as a CSV and emit each record to the channel provided
func NewCSVJiraIssueProgressReaderFile(fp string, ch chan<- JiraIssueProgress) error {
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
	return NewCSVJiraIssueProgressReader(fc, ch)
}

// NewCSVJiraIssueProgressReaderDir will read the jira_issue_progress.csv.gz file as a CSV and emit each record to the channel provided
func NewCSVJiraIssueProgressReaderDir(dir string, ch chan<- JiraIssueProgress) error {
	return NewCSVJiraIssueProgressReaderFile(filepath.Join(dir, "jira_issue_progress.csv.gz"), ch)
}

// JiraIssueProgressCSVDeduper is a function callback which takes the existing value (a) and the new value (b)
// and the return value should be the one to use (or a new one, if applicable). return nil
// to skip processing of this record
type JiraIssueProgressCSVDeduper func(a JiraIssueProgress, b JiraIssueProgress) *JiraIssueProgress

// JiraIssueProgressCSVDedupeDisabled is set on whether the CSV writer should de-dupe values by key
var JiraIssueProgressCSVDedupeDisabled bool

// NewJiraIssueProgressCSVWriterSize creates a batch writer that will write each JiraIssueProgress into a CSV file
// this method will automatically de-duplicate entries using the primary key. if the checksum
// for a newer item with the same primary key doesn't match a previously sent item, the newer
// one will be used
func NewJiraIssueProgressCSVWriterSize(w io.Writer, size int, dedupers ...JiraIssueProgressCSVDeduper) (chan JiraIssueProgress, chan bool, error) {
	cw := csv.NewWriter(w)
	ch := make(chan JiraIssueProgress, size)
	done := make(chan bool)
	go func() {
		defer func() { done <- true }()
		dodedupe := !JiraIssueProgressCSVDedupeDisabled
		var kv map[string]*JiraIssueProgress
		var deduper JiraIssueProgressCSVDeduper
		if dedupers != nil && len(dedupers) > 0 {
			deduper = dedupers[0]
			dodedupe = true
		}
		if dodedupe {
			kv = make(map[string]*JiraIssueProgress)
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

// JiraIssueProgressCSVDefaultSize is the default channel buffer size if not provided
var JiraIssueProgressCSVDefaultSize = 100

// NewJiraIssueProgressCSVWriter creates a batch writer that will write each JiraIssueProgress into a CSV file
func NewJiraIssueProgressCSVWriter(w io.Writer, dedupers ...JiraIssueProgressCSVDeduper) (chan JiraIssueProgress, chan bool, error) {
	return NewJiraIssueProgressCSVWriterSize(w, JiraIssueProgressCSVDefaultSize, dedupers...)
}

// NewJiraIssueProgressCSVWriterDir creates a batch writer that will write each JiraIssueProgress into a CSV file named jira_issue_progress.csv.gz in dir
func NewJiraIssueProgressCSVWriterDir(dir string, dedupers ...JiraIssueProgressCSVDeduper) (chan JiraIssueProgress, chan bool, error) {
	return NewJiraIssueProgressCSVWriterFile(filepath.Join(dir, "jira_issue_progress.csv.gz"), dedupers...)
}

// NewJiraIssueProgressCSVWriterFile creates a batch writer that will write each JiraIssueProgress into a CSV file
func NewJiraIssueProgressCSVWriterFile(fn string, dedupers ...JiraIssueProgressCSVDeduper) (chan JiraIssueProgress, chan bool, error) {
	f, err := os.Create(fn)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening CSV file %s. %v", fn, err)
	}
	var fc io.WriteCloser = f
	if filepath.Ext(fn) == ".gz" {
		w, _ := gzip.NewWriterLevel(f, gzip.BestCompression)
		fc = w
	}
	ch, done, err := NewJiraIssueProgressCSVWriter(fc, dedupers...)
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

type JiraIssueProgressDBAction func(ctx context.Context, db DB, record JiraIssueProgress) error

// NewJiraIssueProgressDBWriterSize creates a DB writer that will write each issue into the DB
func NewJiraIssueProgressDBWriterSize(ctx context.Context, db DB, errors chan<- error, size int, actions ...JiraIssueProgressDBAction) (chan JiraIssueProgress, chan bool, error) {
	ch := make(chan JiraIssueProgress, size)
	done := make(chan bool)
	var action JiraIssueProgressDBAction
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

// NewJiraIssueProgressDBWriter creates a DB writer that will write each issue into the DB
func NewJiraIssueProgressDBWriter(ctx context.Context, db DB, errors chan<- error, actions ...JiraIssueProgressDBAction) (chan JiraIssueProgress, chan bool, error) {
	return NewJiraIssueProgressDBWriterSize(ctx, db, errors, 100, actions...)
}

// JiraIssueProgressColumnID is the ID SQL column name for the JiraIssueProgress table
const JiraIssueProgressColumnID = "id"

// JiraIssueProgressEscapedColumnID is the escaped ID SQL column name for the JiraIssueProgress table
const JiraIssueProgressEscapedColumnID = "`id`"

// JiraIssueProgressColumnChecksum is the Checksum SQL column name for the JiraIssueProgress table
const JiraIssueProgressColumnChecksum = "checksum"

// JiraIssueProgressEscapedColumnChecksum is the escaped Checksum SQL column name for the JiraIssueProgress table
const JiraIssueProgressEscapedColumnChecksum = "`checksum`"

// JiraIssueProgressColumnUserID is the UserID SQL column name for the JiraIssueProgress table
const JiraIssueProgressColumnUserID = "user_id"

// JiraIssueProgressEscapedColumnUserID is the escaped UserID SQL column name for the JiraIssueProgress table
const JiraIssueProgressEscapedColumnUserID = "`user_id`"

// JiraIssueProgressColumnIssueID is the IssueID SQL column name for the JiraIssueProgress table
const JiraIssueProgressColumnIssueID = "issue_id"

// JiraIssueProgressEscapedColumnIssueID is the escaped IssueID SQL column name for the JiraIssueProgress table
const JiraIssueProgressEscapedColumnIssueID = "`issue_id`"

// JiraIssueProgressColumnStartDate is the StartDate SQL column name for the JiraIssueProgress table
const JiraIssueProgressColumnStartDate = "start_date"

// JiraIssueProgressEscapedColumnStartDate is the escaped StartDate SQL column name for the JiraIssueProgress table
const JiraIssueProgressEscapedColumnStartDate = "`start_date`"

// JiraIssueProgressColumnEndDate is the EndDate SQL column name for the JiraIssueProgress table
const JiraIssueProgressColumnEndDate = "end_date"

// JiraIssueProgressEscapedColumnEndDate is the escaped EndDate SQL column name for the JiraIssueProgress table
const JiraIssueProgressEscapedColumnEndDate = "`end_date`"

// JiraIssueProgressColumnDuration is the Duration SQL column name for the JiraIssueProgress table
const JiraIssueProgressColumnDuration = "duration"

// JiraIssueProgressEscapedColumnDuration is the escaped Duration SQL column name for the JiraIssueProgress table
const JiraIssueProgressEscapedColumnDuration = "`duration`"

// JiraIssueProgressColumnCustomerID is the CustomerID SQL column name for the JiraIssueProgress table
const JiraIssueProgressColumnCustomerID = "customer_id"

// JiraIssueProgressEscapedColumnCustomerID is the escaped CustomerID SQL column name for the JiraIssueProgress table
const JiraIssueProgressEscapedColumnCustomerID = "`customer_id`"

// JiraIssueProgressColumnRefID is the RefID SQL column name for the JiraIssueProgress table
const JiraIssueProgressColumnRefID = "ref_id"

// JiraIssueProgressEscapedColumnRefID is the escaped RefID SQL column name for the JiraIssueProgress table
const JiraIssueProgressEscapedColumnRefID = "`ref_id`"

// GetID will return the JiraIssueProgress ID value
func (t *JiraIssueProgress) GetID() string {
	return t.ID
}

// SetID will set the JiraIssueProgress ID value
func (t *JiraIssueProgress) SetID(v string) {
	t.ID = v
}

// FindJiraIssueProgressByID will find a JiraIssueProgress by ID
func FindJiraIssueProgressByID(ctx context.Context, db DB, value string) (*JiraIssueProgress, error) {
	q := "SELECT `jira_issue_progress`.`id`,`jira_issue_progress`.`checksum`,`jira_issue_progress`.`user_id`,`jira_issue_progress`.`issue_id`,`jira_issue_progress`.`start_date`,`jira_issue_progress`.`end_date`,`jira_issue_progress`.`duration`,`jira_issue_progress`.`customer_id`,`jira_issue_progress`.`ref_id` FROM `jira_issue_progress` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _UserID sql.NullString
	var _IssueID sql.NullString
	var _StartDate sql.NullInt64
	var _EndDate sql.NullInt64
	var _Duration sql.NullInt64
	var _CustomerID sql.NullString
	var _RefID sql.NullString
	err := db.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_UserID,
		&_IssueID,
		&_StartDate,
		&_EndDate,
		&_Duration,
		&_CustomerID,
		&_RefID,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &JiraIssueProgress{}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _Checksum.Valid {
		t.SetChecksum(_Checksum.String)
	}
	if _UserID.Valid {
		t.SetUserID(_UserID.String)
	}
	if _IssueID.Valid {
		t.SetIssueID(_IssueID.String)
	}
	if _StartDate.Valid {
		t.SetStartDate(_StartDate.Int64)
	}
	if _EndDate.Valid {
		t.SetEndDate(_EndDate.Int64)
	}
	if _Duration.Valid {
		t.SetDuration(_Duration.Int64)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _RefID.Valid {
		t.SetRefID(_RefID.String)
	}
	return t, nil
}

// FindJiraIssueProgressByIDTx will find a JiraIssueProgress by ID using the provided transaction
func FindJiraIssueProgressByIDTx(ctx context.Context, tx Tx, value string) (*JiraIssueProgress, error) {
	q := "SELECT `jira_issue_progress`.`id`,`jira_issue_progress`.`checksum`,`jira_issue_progress`.`user_id`,`jira_issue_progress`.`issue_id`,`jira_issue_progress`.`start_date`,`jira_issue_progress`.`end_date`,`jira_issue_progress`.`duration`,`jira_issue_progress`.`customer_id`,`jira_issue_progress`.`ref_id` FROM `jira_issue_progress` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _UserID sql.NullString
	var _IssueID sql.NullString
	var _StartDate sql.NullInt64
	var _EndDate sql.NullInt64
	var _Duration sql.NullInt64
	var _CustomerID sql.NullString
	var _RefID sql.NullString
	err := tx.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_UserID,
		&_IssueID,
		&_StartDate,
		&_EndDate,
		&_Duration,
		&_CustomerID,
		&_RefID,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &JiraIssueProgress{}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _Checksum.Valid {
		t.SetChecksum(_Checksum.String)
	}
	if _UserID.Valid {
		t.SetUserID(_UserID.String)
	}
	if _IssueID.Valid {
		t.SetIssueID(_IssueID.String)
	}
	if _StartDate.Valid {
		t.SetStartDate(_StartDate.Int64)
	}
	if _EndDate.Valid {
		t.SetEndDate(_EndDate.Int64)
	}
	if _Duration.Valid {
		t.SetDuration(_Duration.Int64)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _RefID.Valid {
		t.SetRefID(_RefID.String)
	}
	return t, nil
}

// GetChecksum will return the JiraIssueProgress Checksum value
func (t *JiraIssueProgress) GetChecksum() string {
	if t.Checksum == nil {
		return ""
	}
	return *t.Checksum
}

// SetChecksum will set the JiraIssueProgress Checksum value
func (t *JiraIssueProgress) SetChecksum(v string) {
	t.Checksum = &v
}

// GetUserID will return the JiraIssueProgress UserID value
func (t *JiraIssueProgress) GetUserID() string {
	if t.UserID == nil {
		return ""
	}
	return *t.UserID
}

// SetUserID will set the JiraIssueProgress UserID value
func (t *JiraIssueProgress) SetUserID(v string) {
	t.UserID = &v
}

// FindJiraIssueProgressesByUserID will find all JiraIssueProgresss by the UserID value
func FindJiraIssueProgressesByUserID(ctx context.Context, db DB, value string) ([]*JiraIssueProgress, error) {
	q := "SELECT `jira_issue_progress`.`id`,`jira_issue_progress`.`checksum`,`jira_issue_progress`.`user_id`,`jira_issue_progress`.`issue_id`,`jira_issue_progress`.`start_date`,`jira_issue_progress`.`end_date`,`jira_issue_progress`.`duration`,`jira_issue_progress`.`customer_id`,`jira_issue_progress`.`ref_id` FROM `jira_issue_progress` WHERE `user_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraIssueProgress, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _UserID sql.NullString
		var _IssueID sql.NullString
		var _StartDate sql.NullInt64
		var _EndDate sql.NullInt64
		var _Duration sql.NullInt64
		var _CustomerID sql.NullString
		var _RefID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_UserID,
			&_IssueID,
			&_StartDate,
			&_EndDate,
			&_Duration,
			&_CustomerID,
			&_RefID,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraIssueProgress{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _StartDate.Valid {
			t.SetStartDate(_StartDate.Int64)
		}
		if _EndDate.Valid {
			t.SetEndDate(_EndDate.Int64)
		}
		if _Duration.Valid {
			t.SetDuration(_Duration.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindJiraIssueProgressesByUserIDTx will find all JiraIssueProgresss by the UserID value using the provided transaction
func FindJiraIssueProgressesByUserIDTx(ctx context.Context, tx Tx, value string) ([]*JiraIssueProgress, error) {
	q := "SELECT `jira_issue_progress`.`id`,`jira_issue_progress`.`checksum`,`jira_issue_progress`.`user_id`,`jira_issue_progress`.`issue_id`,`jira_issue_progress`.`start_date`,`jira_issue_progress`.`end_date`,`jira_issue_progress`.`duration`,`jira_issue_progress`.`customer_id`,`jira_issue_progress`.`ref_id` FROM `jira_issue_progress` WHERE `user_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraIssueProgress, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _UserID sql.NullString
		var _IssueID sql.NullString
		var _StartDate sql.NullInt64
		var _EndDate sql.NullInt64
		var _Duration sql.NullInt64
		var _CustomerID sql.NullString
		var _RefID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_UserID,
			&_IssueID,
			&_StartDate,
			&_EndDate,
			&_Duration,
			&_CustomerID,
			&_RefID,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraIssueProgress{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _StartDate.Valid {
			t.SetStartDate(_StartDate.Int64)
		}
		if _EndDate.Valid {
			t.SetEndDate(_EndDate.Int64)
		}
		if _Duration.Valid {
			t.SetDuration(_Duration.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetIssueID will return the JiraIssueProgress IssueID value
func (t *JiraIssueProgress) GetIssueID() string {
	return t.IssueID
}

// SetIssueID will set the JiraIssueProgress IssueID value
func (t *JiraIssueProgress) SetIssueID(v string) {
	t.IssueID = v
}

// FindJiraIssueProgressesByIssueID will find all JiraIssueProgresss by the IssueID value
func FindJiraIssueProgressesByIssueID(ctx context.Context, db DB, value string) ([]*JiraIssueProgress, error) {
	q := "SELECT `jira_issue_progress`.`id`,`jira_issue_progress`.`checksum`,`jira_issue_progress`.`user_id`,`jira_issue_progress`.`issue_id`,`jira_issue_progress`.`start_date`,`jira_issue_progress`.`end_date`,`jira_issue_progress`.`duration`,`jira_issue_progress`.`customer_id`,`jira_issue_progress`.`ref_id` FROM `jira_issue_progress` WHERE `issue_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraIssueProgress, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _UserID sql.NullString
		var _IssueID sql.NullString
		var _StartDate sql.NullInt64
		var _EndDate sql.NullInt64
		var _Duration sql.NullInt64
		var _CustomerID sql.NullString
		var _RefID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_UserID,
			&_IssueID,
			&_StartDate,
			&_EndDate,
			&_Duration,
			&_CustomerID,
			&_RefID,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraIssueProgress{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _StartDate.Valid {
			t.SetStartDate(_StartDate.Int64)
		}
		if _EndDate.Valid {
			t.SetEndDate(_EndDate.Int64)
		}
		if _Duration.Valid {
			t.SetDuration(_Duration.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindJiraIssueProgressesByIssueIDTx will find all JiraIssueProgresss by the IssueID value using the provided transaction
func FindJiraIssueProgressesByIssueIDTx(ctx context.Context, tx Tx, value string) ([]*JiraIssueProgress, error) {
	q := "SELECT `jira_issue_progress`.`id`,`jira_issue_progress`.`checksum`,`jira_issue_progress`.`user_id`,`jira_issue_progress`.`issue_id`,`jira_issue_progress`.`start_date`,`jira_issue_progress`.`end_date`,`jira_issue_progress`.`duration`,`jira_issue_progress`.`customer_id`,`jira_issue_progress`.`ref_id` FROM `jira_issue_progress` WHERE `issue_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraIssueProgress, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _UserID sql.NullString
		var _IssueID sql.NullString
		var _StartDate sql.NullInt64
		var _EndDate sql.NullInt64
		var _Duration sql.NullInt64
		var _CustomerID sql.NullString
		var _RefID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_UserID,
			&_IssueID,
			&_StartDate,
			&_EndDate,
			&_Duration,
			&_CustomerID,
			&_RefID,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraIssueProgress{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _StartDate.Valid {
			t.SetStartDate(_StartDate.Int64)
		}
		if _EndDate.Valid {
			t.SetEndDate(_EndDate.Int64)
		}
		if _Duration.Valid {
			t.SetDuration(_Duration.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetStartDate will return the JiraIssueProgress StartDate value
func (t *JiraIssueProgress) GetStartDate() int64 {
	if t.StartDate == nil {
		return int64(0)
	}
	return *t.StartDate
}

// SetStartDate will set the JiraIssueProgress StartDate value
func (t *JiraIssueProgress) SetStartDate(v int64) {
	t.StartDate = &v
}

// GetEndDate will return the JiraIssueProgress EndDate value
func (t *JiraIssueProgress) GetEndDate() int64 {
	if t.EndDate == nil {
		return int64(0)
	}
	return *t.EndDate
}

// SetEndDate will set the JiraIssueProgress EndDate value
func (t *JiraIssueProgress) SetEndDate(v int64) {
	t.EndDate = &v
}

// GetDuration will return the JiraIssueProgress Duration value
func (t *JiraIssueProgress) GetDuration() int64 {
	if t.Duration == nil {
		return int64(0)
	}
	return *t.Duration
}

// SetDuration will set the JiraIssueProgress Duration value
func (t *JiraIssueProgress) SetDuration(v int64) {
	t.Duration = &v
}

// GetCustomerID will return the JiraIssueProgress CustomerID value
func (t *JiraIssueProgress) GetCustomerID() string {
	return t.CustomerID
}

// SetCustomerID will set the JiraIssueProgress CustomerID value
func (t *JiraIssueProgress) SetCustomerID(v string) {
	t.CustomerID = v
}

// FindJiraIssueProgressesByCustomerID will find all JiraIssueProgresss by the CustomerID value
func FindJiraIssueProgressesByCustomerID(ctx context.Context, db DB, value string) ([]*JiraIssueProgress, error) {
	q := "SELECT `jira_issue_progress`.`id`,`jira_issue_progress`.`checksum`,`jira_issue_progress`.`user_id`,`jira_issue_progress`.`issue_id`,`jira_issue_progress`.`start_date`,`jira_issue_progress`.`end_date`,`jira_issue_progress`.`duration`,`jira_issue_progress`.`customer_id`,`jira_issue_progress`.`ref_id` FROM `jira_issue_progress` WHERE `customer_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraIssueProgress, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _UserID sql.NullString
		var _IssueID sql.NullString
		var _StartDate sql.NullInt64
		var _EndDate sql.NullInt64
		var _Duration sql.NullInt64
		var _CustomerID sql.NullString
		var _RefID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_UserID,
			&_IssueID,
			&_StartDate,
			&_EndDate,
			&_Duration,
			&_CustomerID,
			&_RefID,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraIssueProgress{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _StartDate.Valid {
			t.SetStartDate(_StartDate.Int64)
		}
		if _EndDate.Valid {
			t.SetEndDate(_EndDate.Int64)
		}
		if _Duration.Valid {
			t.SetDuration(_Duration.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindJiraIssueProgressesByCustomerIDTx will find all JiraIssueProgresss by the CustomerID value using the provided transaction
func FindJiraIssueProgressesByCustomerIDTx(ctx context.Context, tx Tx, value string) ([]*JiraIssueProgress, error) {
	q := "SELECT `jira_issue_progress`.`id`,`jira_issue_progress`.`checksum`,`jira_issue_progress`.`user_id`,`jira_issue_progress`.`issue_id`,`jira_issue_progress`.`start_date`,`jira_issue_progress`.`end_date`,`jira_issue_progress`.`duration`,`jira_issue_progress`.`customer_id`,`jira_issue_progress`.`ref_id` FROM `jira_issue_progress` WHERE `customer_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraIssueProgress, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _UserID sql.NullString
		var _IssueID sql.NullString
		var _StartDate sql.NullInt64
		var _EndDate sql.NullInt64
		var _Duration sql.NullInt64
		var _CustomerID sql.NullString
		var _RefID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_UserID,
			&_IssueID,
			&_StartDate,
			&_EndDate,
			&_Duration,
			&_CustomerID,
			&_RefID,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraIssueProgress{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _StartDate.Valid {
			t.SetStartDate(_StartDate.Int64)
		}
		if _EndDate.Valid {
			t.SetEndDate(_EndDate.Int64)
		}
		if _Duration.Valid {
			t.SetDuration(_Duration.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetRefID will return the JiraIssueProgress RefID value
func (t *JiraIssueProgress) GetRefID() string {
	return t.RefID
}

// SetRefID will set the JiraIssueProgress RefID value
func (t *JiraIssueProgress) SetRefID(v string) {
	t.RefID = v
}

// FindJiraIssueProgressesByRefID will find all JiraIssueProgresss by the RefID value
func FindJiraIssueProgressesByRefID(ctx context.Context, db DB, value string) ([]*JiraIssueProgress, error) {
	q := "SELECT `jira_issue_progress`.`id`,`jira_issue_progress`.`checksum`,`jira_issue_progress`.`user_id`,`jira_issue_progress`.`issue_id`,`jira_issue_progress`.`start_date`,`jira_issue_progress`.`end_date`,`jira_issue_progress`.`duration`,`jira_issue_progress`.`customer_id`,`jira_issue_progress`.`ref_id` FROM `jira_issue_progress` WHERE `ref_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraIssueProgress, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _UserID sql.NullString
		var _IssueID sql.NullString
		var _StartDate sql.NullInt64
		var _EndDate sql.NullInt64
		var _Duration sql.NullInt64
		var _CustomerID sql.NullString
		var _RefID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_UserID,
			&_IssueID,
			&_StartDate,
			&_EndDate,
			&_Duration,
			&_CustomerID,
			&_RefID,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraIssueProgress{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _StartDate.Valid {
			t.SetStartDate(_StartDate.Int64)
		}
		if _EndDate.Valid {
			t.SetEndDate(_EndDate.Int64)
		}
		if _Duration.Valid {
			t.SetDuration(_Duration.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindJiraIssueProgressesByRefIDTx will find all JiraIssueProgresss by the RefID value using the provided transaction
func FindJiraIssueProgressesByRefIDTx(ctx context.Context, tx Tx, value string) ([]*JiraIssueProgress, error) {
	q := "SELECT `jira_issue_progress`.`id`,`jira_issue_progress`.`checksum`,`jira_issue_progress`.`user_id`,`jira_issue_progress`.`issue_id`,`jira_issue_progress`.`start_date`,`jira_issue_progress`.`end_date`,`jira_issue_progress`.`duration`,`jira_issue_progress`.`customer_id`,`jira_issue_progress`.`ref_id` FROM `jira_issue_progress` WHERE `ref_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraIssueProgress, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _UserID sql.NullString
		var _IssueID sql.NullString
		var _StartDate sql.NullInt64
		var _EndDate sql.NullInt64
		var _Duration sql.NullInt64
		var _CustomerID sql.NullString
		var _RefID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_UserID,
			&_IssueID,
			&_StartDate,
			&_EndDate,
			&_Duration,
			&_CustomerID,
			&_RefID,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraIssueProgress{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _StartDate.Valid {
			t.SetStartDate(_StartDate.Int64)
		}
		if _EndDate.Valid {
			t.SetEndDate(_EndDate.Int64)
		}
		if _Duration.Valid {
			t.SetDuration(_Duration.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

func (t *JiraIssueProgress) toTimestamp(value time.Time) *timestamp.Timestamp {
	ts, _ := ptypes.TimestampProto(value)
	return ts
}

// DBCreateJiraIssueProgressTable will create the JiraIssueProgress table
func DBCreateJiraIssueProgressTable(ctx context.Context, db DB) error {
	q := "CREATE TABLE `jira_issue_progress` (`id` VARCHAR(64) NOT NULL PRIMARY KEY,`checksum` CHAR(64),`user_id`VARCHAR(64),`issue_id` VARCHAR(64) NOT NULL,`start_date`BIGINT UNSIGNED,`end_date` BIGINT UNSIGNED,`duration` BIGINT UNSIGNED,`customer_id` VARCHAR(64) NOT NULL,`ref_id` VARCHAR(64) NOT NULL,INDEX jira_issue_progress_user_id_index (`user_id`),INDEX jira_issue_progress_issue_id_index (`issue_id`),INDEX jira_issue_progress_customer_id_index (`customer_id`),INDEX jira_issue_progress_ref_id_index (`ref_id`),INDEX jira_issue_progress_user_id_customer_id_start_date_index (`user_id`,`customer_id`,`start_date`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBCreateJiraIssueProgressTableTx will create the JiraIssueProgress table using the provided transction
func DBCreateJiraIssueProgressTableTx(ctx context.Context, tx Tx) error {
	q := "CREATE TABLE `jira_issue_progress` (`id` VARCHAR(64) NOT NULL PRIMARY KEY,`checksum` CHAR(64),`user_id`VARCHAR(64),`issue_id` VARCHAR(64) NOT NULL,`start_date`BIGINT UNSIGNED,`end_date` BIGINT UNSIGNED,`duration` BIGINT UNSIGNED,`customer_id` VARCHAR(64) NOT NULL,`ref_id` VARCHAR(64) NOT NULL,INDEX jira_issue_progress_user_id_index (`user_id`),INDEX jira_issue_progress_issue_id_index (`issue_id`),INDEX jira_issue_progress_customer_id_index (`customer_id`),INDEX jira_issue_progress_ref_id_index (`ref_id`),INDEX jira_issue_progress_user_id_customer_id_start_date_index (`user_id`,`customer_id`,`start_date`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// DBDropJiraIssueProgressTable will drop the JiraIssueProgress table
func DBDropJiraIssueProgressTable(ctx context.Context, db DB) error {
	q := "DROP TABLE IF EXISTS `jira_issue_progress`"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBDropJiraIssueProgressTableTx will drop the JiraIssueProgress table using the provided transaction
func DBDropJiraIssueProgressTableTx(ctx context.Context, tx Tx) error {
	q := "DROP TABLE IF EXISTS `jira_issue_progress`"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// CalculateChecksum will calculate a checksum of the SHA1 of all field values
func (t *JiraIssueProgress) CalculateChecksum() string {
	return orm.HashStrings(
		orm.ToString(t.ID),
		orm.ToString(t.UserID),
		orm.ToString(t.IssueID),
		orm.ToString(t.StartDate),
		orm.ToString(t.EndDate),
		orm.ToString(t.Duration),
		orm.ToString(t.CustomerID),
		orm.ToString(t.RefID),
	)
}

// DBCreate will create a new JiraIssueProgress record in the database
func (t *JiraIssueProgress) DBCreate(ctx context.Context, db DB) (sql.Result, error) {
	q := "INSERT INTO `jira_issue_progress` (`jira_issue_progress`.`id`,`jira_issue_progress`.`checksum`,`jira_issue_progress`.`user_id`,`jira_issue_progress`.`issue_id`,`jira_issue_progress`.`start_date`,`jira_issue_progress`.`end_date`,`jira_issue_progress`.`duration`,`jira_issue_progress`.`customer_id`,`jira_issue_progress`.`ref_id`) VALUES (?,?,?,?,?,?,?,?,?)"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.UserID),
		orm.ToSQLString(t.IssueID),
		orm.ToSQLInt64(t.StartDate),
		orm.ToSQLInt64(t.EndDate),
		orm.ToSQLInt64(t.Duration),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefID),
	)
}

// DBCreateTx will create a new JiraIssueProgress record in the database using the provided transaction
func (t *JiraIssueProgress) DBCreateTx(ctx context.Context, tx Tx) (sql.Result, error) {
	q := "INSERT INTO `jira_issue_progress` (`jira_issue_progress`.`id`,`jira_issue_progress`.`checksum`,`jira_issue_progress`.`user_id`,`jira_issue_progress`.`issue_id`,`jira_issue_progress`.`start_date`,`jira_issue_progress`.`end_date`,`jira_issue_progress`.`duration`,`jira_issue_progress`.`customer_id`,`jira_issue_progress`.`ref_id`) VALUES (?,?,?,?,?,?,?,?,?)"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.UserID),
		orm.ToSQLString(t.IssueID),
		orm.ToSQLInt64(t.StartDate),
		orm.ToSQLInt64(t.EndDate),
		orm.ToSQLInt64(t.Duration),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefID),
	)
}

// DBCreateIgnoreDuplicate will upsert the JiraIssueProgress record in the database
func (t *JiraIssueProgress) DBCreateIgnoreDuplicate(ctx context.Context, db DB) (sql.Result, error) {
	q := "INSERT INTO `jira_issue_progress` (`jira_issue_progress`.`id`,`jira_issue_progress`.`checksum`,`jira_issue_progress`.`user_id`,`jira_issue_progress`.`issue_id`,`jira_issue_progress`.`start_date`,`jira_issue_progress`.`end_date`,`jira_issue_progress`.`duration`,`jira_issue_progress`.`customer_id`,`jira_issue_progress`.`ref_id`) VALUES (?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.UserID),
		orm.ToSQLString(t.IssueID),
		orm.ToSQLInt64(t.StartDate),
		orm.ToSQLInt64(t.EndDate),
		orm.ToSQLInt64(t.Duration),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefID),
	)
}

// DBCreateIgnoreDuplicateTx will upsert the JiraIssueProgress record in the database using the provided transaction
func (t *JiraIssueProgress) DBCreateIgnoreDuplicateTx(ctx context.Context, tx Tx) (sql.Result, error) {
	q := "INSERT INTO `jira_issue_progress` (`jira_issue_progress`.`id`,`jira_issue_progress`.`checksum`,`jira_issue_progress`.`user_id`,`jira_issue_progress`.`issue_id`,`jira_issue_progress`.`start_date`,`jira_issue_progress`.`end_date`,`jira_issue_progress`.`duration`,`jira_issue_progress`.`customer_id`,`jira_issue_progress`.`ref_id`) VALUES (?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.UserID),
		orm.ToSQLString(t.IssueID),
		orm.ToSQLInt64(t.StartDate),
		orm.ToSQLInt64(t.EndDate),
		orm.ToSQLInt64(t.Duration),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefID),
	)
}

// DeleteAllJiraIssueProgresses deletes all JiraIssueProgress records in the database with optional filters
func DeleteAllJiraIssueProgresses(ctx context.Context, db DB, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(JiraIssueProgressTableName),
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

// DeleteAllJiraIssueProgressesTx deletes all JiraIssueProgress records in the database with optional filters using the provided transaction
func DeleteAllJiraIssueProgressesTx(ctx context.Context, tx Tx, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(JiraIssueProgressTableName),
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

// DBDelete will delete this JiraIssueProgress record in the database
func (t *JiraIssueProgress) DBDelete(ctx context.Context, db DB) (bool, error) {
	q := "DELETE FROM `jira_issue_progress` WHERE `id` = ?"
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

// DBDeleteTx will delete this JiraIssueProgress record in the database using the provided transaction
func (t *JiraIssueProgress) DBDeleteTx(ctx context.Context, tx Tx) (bool, error) {
	q := "DELETE FROM `jira_issue_progress` WHERE `id` = ?"
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

// DBUpdate will update the JiraIssueProgress record in the database
func (t *JiraIssueProgress) DBUpdate(ctx context.Context, db DB) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `jira_issue_progress` SET `checksum`=?,`user_id`=?,`issue_id`=?,`start_date`=?,`end_date`=?,`duration`=?,`customer_id`=?,`ref_id`=? WHERE `id`=?"
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.UserID),
		orm.ToSQLString(t.IssueID),
		orm.ToSQLInt64(t.StartDate),
		orm.ToSQLInt64(t.EndDate),
		orm.ToSQLInt64(t.Duration),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.ID),
	)
}

// DBUpdateTx will update the JiraIssueProgress record in the database using the provided transaction
func (t *JiraIssueProgress) DBUpdateTx(ctx context.Context, tx Tx) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `jira_issue_progress` SET `checksum`=?,`user_id`=?,`issue_id`=?,`start_date`=?,`end_date`=?,`duration`=?,`customer_id`=?,`ref_id`=? WHERE `id`=?"
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.UserID),
		orm.ToSQLString(t.IssueID),
		orm.ToSQLInt64(t.StartDate),
		orm.ToSQLInt64(t.EndDate),
		orm.ToSQLInt64(t.Duration),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.ID),
	)
}

// DBUpsert will upsert the JiraIssueProgress record in the database
func (t *JiraIssueProgress) DBUpsert(ctx context.Context, db DB, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `jira_issue_progress` (`jira_issue_progress`.`id`,`jira_issue_progress`.`checksum`,`jira_issue_progress`.`user_id`,`jira_issue_progress`.`issue_id`,`jira_issue_progress`.`start_date`,`jira_issue_progress`.`end_date`,`jira_issue_progress`.`duration`,`jira_issue_progress`.`customer_id`,`jira_issue_progress`.`ref_id`) VALUES (?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `jira_issue_progress` (`jira_issue_progress`.`id`,`jira_issue_progress`.`checksum`,`jira_issue_progress`.`user_id`,`jira_issue_progress`.`issue_id`,`jira_issue_progress`.`start_date`,`jira_issue_progress`.`end_date`,`jira_issue_progress`.`duration`,`jira_issue_progress`.`customer_id`,`jira_issue_progress`.`ref_id`) VALUES (?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`user_id`=VALUES(`user_id`),`issue_id`=VALUES(`issue_id`),`start_date`=VALUES(`start_date`),`end_date`=VALUES(`end_date`),`duration`=VALUES(`duration`),`customer_id`=VALUES(`customer_id`),`ref_id`=VALUES(`ref_id`)"
	}
	r, err := db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.UserID),
		orm.ToSQLString(t.IssueID),
		orm.ToSQLInt64(t.StartDate),
		orm.ToSQLInt64(t.EndDate),
		orm.ToSQLInt64(t.Duration),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefID),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBUpsertTx will upsert the JiraIssueProgress record in the database using the provided transaction
func (t *JiraIssueProgress) DBUpsertTx(ctx context.Context, tx Tx, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `jira_issue_progress` (`jira_issue_progress`.`id`,`jira_issue_progress`.`checksum`,`jira_issue_progress`.`user_id`,`jira_issue_progress`.`issue_id`,`jira_issue_progress`.`start_date`,`jira_issue_progress`.`end_date`,`jira_issue_progress`.`duration`,`jira_issue_progress`.`customer_id`,`jira_issue_progress`.`ref_id`) VALUES (?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `jira_issue_progress` (`jira_issue_progress`.`id`,`jira_issue_progress`.`checksum`,`jira_issue_progress`.`user_id`,`jira_issue_progress`.`issue_id`,`jira_issue_progress`.`start_date`,`jira_issue_progress`.`end_date`,`jira_issue_progress`.`duration`,`jira_issue_progress`.`customer_id`,`jira_issue_progress`.`ref_id`) VALUES (?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`user_id`=VALUES(`user_id`),`issue_id`=VALUES(`issue_id`),`start_date`=VALUES(`start_date`),`end_date`=VALUES(`end_date`),`duration`=VALUES(`duration`),`customer_id`=VALUES(`customer_id`),`ref_id`=VALUES(`ref_id`)"
	}
	r, err := tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.UserID),
		orm.ToSQLString(t.IssueID),
		orm.ToSQLInt64(t.StartDate),
		orm.ToSQLInt64(t.EndDate),
		orm.ToSQLInt64(t.Duration),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefID),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBFindOne will find a JiraIssueProgress record in the database with the primary key
func (t *JiraIssueProgress) DBFindOne(ctx context.Context, db DB, value string) (bool, error) {
	q := "SELECT `jira_issue_progress`.`id`,`jira_issue_progress`.`checksum`,`jira_issue_progress`.`user_id`,`jira_issue_progress`.`issue_id`,`jira_issue_progress`.`start_date`,`jira_issue_progress`.`end_date`,`jira_issue_progress`.`duration`,`jira_issue_progress`.`customer_id`,`jira_issue_progress`.`ref_id` FROM `jira_issue_progress` WHERE `id` = ? LIMIT 1"
	row := db.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _UserID sql.NullString
	var _IssueID sql.NullString
	var _StartDate sql.NullInt64
	var _EndDate sql.NullInt64
	var _Duration sql.NullInt64
	var _CustomerID sql.NullString
	var _RefID sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_UserID,
		&_IssueID,
		&_StartDate,
		&_EndDate,
		&_Duration,
		&_CustomerID,
		&_RefID,
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
	if _UserID.Valid {
		t.SetUserID(_UserID.String)
	}
	if _IssueID.Valid {
		t.SetIssueID(_IssueID.String)
	}
	if _StartDate.Valid {
		t.SetStartDate(_StartDate.Int64)
	}
	if _EndDate.Valid {
		t.SetEndDate(_EndDate.Int64)
	}
	if _Duration.Valid {
		t.SetDuration(_Duration.Int64)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _RefID.Valid {
		t.SetRefID(_RefID.String)
	}
	return true, nil
}

// DBFindOneTx will find a JiraIssueProgress record in the database with the primary key using the provided transaction
func (t *JiraIssueProgress) DBFindOneTx(ctx context.Context, tx Tx, value string) (bool, error) {
	q := "SELECT `jira_issue_progress`.`id`,`jira_issue_progress`.`checksum`,`jira_issue_progress`.`user_id`,`jira_issue_progress`.`issue_id`,`jira_issue_progress`.`start_date`,`jira_issue_progress`.`end_date`,`jira_issue_progress`.`duration`,`jira_issue_progress`.`customer_id`,`jira_issue_progress`.`ref_id` FROM `jira_issue_progress` WHERE `id` = ? LIMIT 1"
	row := tx.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _UserID sql.NullString
	var _IssueID sql.NullString
	var _StartDate sql.NullInt64
	var _EndDate sql.NullInt64
	var _Duration sql.NullInt64
	var _CustomerID sql.NullString
	var _RefID sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_UserID,
		&_IssueID,
		&_StartDate,
		&_EndDate,
		&_Duration,
		&_CustomerID,
		&_RefID,
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
	if _UserID.Valid {
		t.SetUserID(_UserID.String)
	}
	if _IssueID.Valid {
		t.SetIssueID(_IssueID.String)
	}
	if _StartDate.Valid {
		t.SetStartDate(_StartDate.Int64)
	}
	if _EndDate.Valid {
		t.SetEndDate(_EndDate.Int64)
	}
	if _Duration.Valid {
		t.SetDuration(_Duration.Int64)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _RefID.Valid {
		t.SetRefID(_RefID.String)
	}
	return true, nil
}

// FindJiraIssueProgresses will find a JiraIssueProgress record in the database with the provided parameters
func FindJiraIssueProgresses(ctx context.Context, db DB, _params ...interface{}) ([]*JiraIssueProgress, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("user_id"),
		orm.Column("issue_id"),
		orm.Column("start_date"),
		orm.Column("end_date"),
		orm.Column("duration"),
		orm.Column("customer_id"),
		orm.Column("ref_id"),
		orm.Table(JiraIssueProgressTableName),
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
	results := make([]*JiraIssueProgress, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _UserID sql.NullString
		var _IssueID sql.NullString
		var _StartDate sql.NullInt64
		var _EndDate sql.NullInt64
		var _Duration sql.NullInt64
		var _CustomerID sql.NullString
		var _RefID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_UserID,
			&_IssueID,
			&_StartDate,
			&_EndDate,
			&_Duration,
			&_CustomerID,
			&_RefID,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraIssueProgress{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _StartDate.Valid {
			t.SetStartDate(_StartDate.Int64)
		}
		if _EndDate.Valid {
			t.SetEndDate(_EndDate.Int64)
		}
		if _Duration.Valid {
			t.SetDuration(_Duration.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindJiraIssueProgressesTx will find a JiraIssueProgress record in the database with the provided parameters using the provided transaction
func FindJiraIssueProgressesTx(ctx context.Context, tx Tx, _params ...interface{}) ([]*JiraIssueProgress, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("user_id"),
		orm.Column("issue_id"),
		orm.Column("start_date"),
		orm.Column("end_date"),
		orm.Column("duration"),
		orm.Column("customer_id"),
		orm.Column("ref_id"),
		orm.Table(JiraIssueProgressTableName),
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
	results := make([]*JiraIssueProgress, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _UserID sql.NullString
		var _IssueID sql.NullString
		var _StartDate sql.NullInt64
		var _EndDate sql.NullInt64
		var _Duration sql.NullInt64
		var _CustomerID sql.NullString
		var _RefID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_UserID,
			&_IssueID,
			&_StartDate,
			&_EndDate,
			&_Duration,
			&_CustomerID,
			&_RefID,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraIssueProgress{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _StartDate.Valid {
			t.SetStartDate(_StartDate.Int64)
		}
		if _EndDate.Valid {
			t.SetEndDate(_EndDate.Int64)
		}
		if _Duration.Valid {
			t.SetDuration(_Duration.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// DBFind will find a JiraIssueProgress record in the database with the provided parameters
func (t *JiraIssueProgress) DBFind(ctx context.Context, db DB, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("user_id"),
		orm.Column("issue_id"),
		orm.Column("start_date"),
		orm.Column("end_date"),
		orm.Column("duration"),
		orm.Column("customer_id"),
		orm.Column("ref_id"),
		orm.Table(JiraIssueProgressTableName),
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
	var _UserID sql.NullString
	var _IssueID sql.NullString
	var _StartDate sql.NullInt64
	var _EndDate sql.NullInt64
	var _Duration sql.NullInt64
	var _CustomerID sql.NullString
	var _RefID sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_UserID,
		&_IssueID,
		&_StartDate,
		&_EndDate,
		&_Duration,
		&_CustomerID,
		&_RefID,
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
	if _UserID.Valid {
		t.SetUserID(_UserID.String)
	}
	if _IssueID.Valid {
		t.SetIssueID(_IssueID.String)
	}
	if _StartDate.Valid {
		t.SetStartDate(_StartDate.Int64)
	}
	if _EndDate.Valid {
		t.SetEndDate(_EndDate.Int64)
	}
	if _Duration.Valid {
		t.SetDuration(_Duration.Int64)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _RefID.Valid {
		t.SetRefID(_RefID.String)
	}
	return true, nil
}

// DBFindTx will find a JiraIssueProgress record in the database with the provided parameters using the provided transaction
func (t *JiraIssueProgress) DBFindTx(ctx context.Context, tx Tx, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("user_id"),
		orm.Column("issue_id"),
		orm.Column("start_date"),
		orm.Column("end_date"),
		orm.Column("duration"),
		orm.Column("customer_id"),
		orm.Column("ref_id"),
		orm.Table(JiraIssueProgressTableName),
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
	var _UserID sql.NullString
	var _IssueID sql.NullString
	var _StartDate sql.NullInt64
	var _EndDate sql.NullInt64
	var _Duration sql.NullInt64
	var _CustomerID sql.NullString
	var _RefID sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_UserID,
		&_IssueID,
		&_StartDate,
		&_EndDate,
		&_Duration,
		&_CustomerID,
		&_RefID,
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
	if _UserID.Valid {
		t.SetUserID(_UserID.String)
	}
	if _IssueID.Valid {
		t.SetIssueID(_IssueID.String)
	}
	if _StartDate.Valid {
		t.SetStartDate(_StartDate.Int64)
	}
	if _EndDate.Valid {
		t.SetEndDate(_EndDate.Int64)
	}
	if _Duration.Valid {
		t.SetDuration(_Duration.Int64)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _RefID.Valid {
		t.SetRefID(_RefID.String)
	}
	return true, nil
}

// CountJiraIssueProgresses will find the count of JiraIssueProgress records in the database
func CountJiraIssueProgresses(ctx context.Context, db DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(JiraIssueProgressTableName),
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

// CountJiraIssueProgressesTx will find the count of JiraIssueProgress records in the database using the provided transaction
func CountJiraIssueProgressesTx(ctx context.Context, tx Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(JiraIssueProgressTableName),
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

// DBCount will find the count of JiraIssueProgress records in the database
func (t *JiraIssueProgress) DBCount(ctx context.Context, db DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(JiraIssueProgressTableName),
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

// DBCountTx will find the count of JiraIssueProgress records in the database using the provided transaction
func (t *JiraIssueProgress) DBCountTx(ctx context.Context, tx Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(JiraIssueProgressTableName),
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

// DBExists will return true if the JiraIssueProgress record exists in the database
func (t *JiraIssueProgress) DBExists(ctx context.Context, db DB) (bool, error) {
	q := "SELECT `id` FROM `jira_issue_progress` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := db.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// DBExistsTx will return true if the JiraIssueProgress record exists in the database using the provided transaction
func (t *JiraIssueProgress) DBExistsTx(ctx context.Context, tx Tx) (bool, error) {
	q := "SELECT `id` FROM `jira_issue_progress` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := tx.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// PrimaryKeyColumn returns the column name for the primary key
func (t *JiraIssueProgress) PrimaryKeyColumn() string {
	return JiraIssueProgressColumnID
}

// PrimaryKeyColumnType returns the primary key column Go type as a string
func (t *JiraIssueProgress) PrimaryKeyColumnType() string {
	return "string"
}

// PrimaryKey returns the primary key column value
func (t *JiraIssueProgress) PrimaryKey() interface{} {
	return t.ID
}
