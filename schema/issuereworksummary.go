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

var _ Model = (*issueReworkSummary)(nil)
var _ CSVWriter = (*issueReworkSummary)(nil)
var _ JSONWriter = (*issueReworkSummary)(nil)
var _ Checksum = (*issueReworkSummary)(nil)

// issueReworkSummaryTableName is the name of the table in SQL
const issueReworkSummaryTableName = "issue_rework_summary"

var issueReworkSummaryColumns = []string{
	"id",
	"checksum",
	"customer_id",
	"project_id",
	"user_id",
	"issue_id",
	"path",
	"date",
}

// issueReworkSummary table
type issueReworkSummary struct {
	Checksum   *string `json:"checksum,omitempty"`
	CustomerID string  `json:"customer_id"`
	Date       *int64  `json:"date,omitempty"`
	ID         string  `json:"id"`
	IssueID    string  `json:"issue_id"`
	Path       string  `json:"path"`
	ProjectID  string  `json:"project_id"`
	UserID     *string `json:"user_id,omitempty"`
}

// TableName returns the SQL table name for issueReworkSummary and satifies the Model interface
func (t *issueReworkSummary) TableName() string {
	return issueReworkSummaryTableName
}

// ToCSV will serialize the issueReworkSummary instance to a CSV compatible array of strings
func (t *issueReworkSummary) ToCSV() []string {
	return []string{
		t.ID,
		t.CalculateChecksum(),
		t.CustomerID,
		t.ProjectID,
		toCSVString(t.UserID),
		t.IssueID,
		t.Path,
		toCSVString(t.Date),
	}
}

// WriteCSV will serialize the issueReworkSummary instance to the writer as CSV and satisfies the CSVWriter interface
func (t *issueReworkSummary) WriteCSV(w *csv.Writer) error {
	return w.Write(t.ToCSV())
}

// WriteJSON will serialize the issueReworkSummary instance to the writer as JSON and satisfies the JSONWriter interface
func (t *issueReworkSummary) WriteJSON(w io.Writer, indent ...bool) error {
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

// NewissueReworkSummaryReader creates a JSON reader which can read in issueReworkSummary objects serialized as JSON either as an array, single object or json new lines
// and writes each issueReworkSummary to the channel provided
func NewissueReworkSummaryReader(r io.Reader, ch chan<- issueReworkSummary) error {
	return orm.Deserialize(r, func(buf json.RawMessage) error {
		dec := json.NewDecoder(bytes.NewBuffer(buf))
		e := issueReworkSummary{}
		if err := dec.Decode(&e); err != nil {
			return err
		}
		ch <- e
		return nil
	})
}

// NewCSVissueReworkSummaryReaderDir will read the reader as CSV and emit each record to the channel provided
func NewCSVissueReworkSummaryReader(r io.Reader, ch chan<- issueReworkSummary) error {
	cr := csv.NewReader(r)
	for {
		record, err := cr.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		ch <- issueReworkSummary{
			ID:         record[0],
			Checksum:   fromStringPointer(record[1]),
			CustomerID: record[2],
			ProjectID:  record[3],
			UserID:     fromStringPointer(record[4]),
			IssueID:    record[5],
			Path:       record[6],
			Date:       fromCSVInt64Pointer(record[7]),
		}
	}
	return nil
}

// NewCSVissueReworkSummaryReaderFile will read the file as a CSV and emit each record to the channel provided
func NewCSVissueReworkSummaryReaderFile(fp string, ch chan<- issueReworkSummary) error {
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
	return NewCSVissueReworkSummaryReader(fc, ch)
}

// NewCSVissueReworkSummaryReaderDir will read the issue_rework_summary.csv.gz file as a CSV and emit each record to the channel provided
func NewCSVissueReworkSummaryReaderDir(dir string, ch chan<- issueReworkSummary) error {
	return NewCSVissueReworkSummaryReaderFile(filepath.Join(dir, "issue_rework_summary.csv.gz"), ch)
}

// issueReworkSummaryCSVDeduper is a function callback which takes the existing value (a) and the new value (b)
// and the return value should be the one to use (or a new one, if applicable). return nil
// to skip processing of this record
type issueReworkSummaryCSVDeduper func(a issueReworkSummary, b issueReworkSummary) *issueReworkSummary

// issueReworkSummaryCSVDedupeDisabled is set on whether the CSV writer should de-dupe values by key
var issueReworkSummaryCSVDedupeDisabled bool

// NewissueReworkSummaryCSVWriterSize creates a batch writer that will write each issueReworkSummary into a CSV file
// this method will automatically de-duplicate entries using the primary key. if the checksum
// for a newer item with the same primary key doesn't match a previously sent item, the newer
// one will be used
func NewissueReworkSummaryCSVWriterSize(w io.Writer, size int, dedupers ...issueReworkSummaryCSVDeduper) (chan issueReworkSummary, chan bool, error) {
	cw := csv.NewWriter(w)
	ch := make(chan issueReworkSummary, size)
	done := make(chan bool)
	go func() {
		defer func() { done <- true }()
		dodedupe := !issueReworkSummaryCSVDedupeDisabled
		var kv map[string]*issueReworkSummary
		var deduper issueReworkSummaryCSVDeduper
		if dedupers != nil && len(dedupers) > 0 {
			deduper = dedupers[0]
			dodedupe = true
		}
		if dodedupe {
			kv = make(map[string]*issueReworkSummary)
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

// issueReworkSummaryCSVDefaultSize is the default channel buffer size if not provided
var issueReworkSummaryCSVDefaultSize = 100

// NewissueReworkSummaryCSVWriter creates a batch writer that will write each issueReworkSummary into a CSV file
func NewissueReworkSummaryCSVWriter(w io.Writer, dedupers ...issueReworkSummaryCSVDeduper) (chan issueReworkSummary, chan bool, error) {
	return NewissueReworkSummaryCSVWriterSize(w, issueReworkSummaryCSVDefaultSize, dedupers...)
}

// NewissueReworkSummaryCSVWriterDir creates a batch writer that will write each issueReworkSummary into a CSV file named issue_rework_summary.csv.gz in dir
func NewissueReworkSummaryCSVWriterDir(dir string, dedupers ...issueReworkSummaryCSVDeduper) (chan issueReworkSummary, chan bool, error) {
	return NewissueReworkSummaryCSVWriterFile(filepath.Join(dir, "issue_rework_summary.csv.gz"), dedupers...)
}

// NewissueReworkSummaryCSVWriterFile creates a batch writer that will write each issueReworkSummary into a CSV file
func NewissueReworkSummaryCSVWriterFile(fn string, dedupers ...issueReworkSummaryCSVDeduper) (chan issueReworkSummary, chan bool, error) {
	f, err := os.Create(fn)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening CSV file %s. %v", fn, err)
	}
	var fc io.WriteCloser = f
	if filepath.Ext(fn) == ".gz" {
		w, _ := gzip.NewWriterLevel(f, gzip.BestCompression)
		fc = w
	}
	ch, done, err := NewissueReworkSummaryCSVWriter(fc, dedupers...)
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

type issueReworkSummaryDBAction func(ctx context.Context, db DB, record issueReworkSummary) error

// NewissueReworkSummaryDBWriterSize creates a DB writer that will write each issue into the DB
func NewissueReworkSummaryDBWriterSize(ctx context.Context, db DB, errors chan<- error, size int, actions ...issueReworkSummaryDBAction) (chan issueReworkSummary, chan bool, error) {
	ch := make(chan issueReworkSummary, size)
	done := make(chan bool)
	var action issueReworkSummaryDBAction
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

// NewissueReworkSummaryDBWriter creates a DB writer that will write each issue into the DB
func NewissueReworkSummaryDBWriter(ctx context.Context, db DB, errors chan<- error, actions ...issueReworkSummaryDBAction) (chan issueReworkSummary, chan bool, error) {
	return NewissueReworkSummaryDBWriterSize(ctx, db, errors, 100, actions...)
}

// issueReworkSummaryColumnID is the ID SQL column name for the issueReworkSummary table
const issueReworkSummaryColumnID = "id"

// issueReworkSummaryEscapedColumnID is the escaped ID SQL column name for the issueReworkSummary table
const issueReworkSummaryEscapedColumnID = "`id`"

// issueReworkSummaryColumnChecksum is the Checksum SQL column name for the issueReworkSummary table
const issueReworkSummaryColumnChecksum = "checksum"

// issueReworkSummaryEscapedColumnChecksum is the escaped Checksum SQL column name for the issueReworkSummary table
const issueReworkSummaryEscapedColumnChecksum = "`checksum`"

// issueReworkSummaryColumnCustomerID is the CustomerID SQL column name for the issueReworkSummary table
const issueReworkSummaryColumnCustomerID = "customer_id"

// issueReworkSummaryEscapedColumnCustomerID is the escaped CustomerID SQL column name for the issueReworkSummary table
const issueReworkSummaryEscapedColumnCustomerID = "`customer_id`"

// issueReworkSummaryColumnProjectID is the ProjectID SQL column name for the issueReworkSummary table
const issueReworkSummaryColumnProjectID = "project_id"

// issueReworkSummaryEscapedColumnProjectID is the escaped ProjectID SQL column name for the issueReworkSummary table
const issueReworkSummaryEscapedColumnProjectID = "`project_id`"

// issueReworkSummaryColumnUserID is the UserID SQL column name for the issueReworkSummary table
const issueReworkSummaryColumnUserID = "user_id"

// issueReworkSummaryEscapedColumnUserID is the escaped UserID SQL column name for the issueReworkSummary table
const issueReworkSummaryEscapedColumnUserID = "`user_id`"

// issueReworkSummaryColumnIssueID is the IssueID SQL column name for the issueReworkSummary table
const issueReworkSummaryColumnIssueID = "issue_id"

// issueReworkSummaryEscapedColumnIssueID is the escaped IssueID SQL column name for the issueReworkSummary table
const issueReworkSummaryEscapedColumnIssueID = "`issue_id`"

// issueReworkSummaryColumnPath is the Path SQL column name for the issueReworkSummary table
const issueReworkSummaryColumnPath = "path"

// issueReworkSummaryEscapedColumnPath is the escaped Path SQL column name for the issueReworkSummary table
const issueReworkSummaryEscapedColumnPath = "`path`"

// issueReworkSummaryColumnDate is the Date SQL column name for the issueReworkSummary table
const issueReworkSummaryColumnDate = "date"

// issueReworkSummaryEscapedColumnDate is the escaped Date SQL column name for the issueReworkSummary table
const issueReworkSummaryEscapedColumnDate = "`date`"

// GetID will return the issueReworkSummary ID value
func (t *issueReworkSummary) GetID() string {
	return t.ID
}

// SetID will set the issueReworkSummary ID value
func (t *issueReworkSummary) SetID(v string) {
	t.ID = v
}

// FindIssueReworkSummaryByID will find a issueReworkSummary by ID
func FindIssueReworkSummaryByID(ctx context.Context, db DB, value string) (*issueReworkSummary, error) {
	q := "SELECT `issue_rework_summary`.`id`,`issue_rework_summary`.`checksum`,`issue_rework_summary`.`customer_id`,`issue_rework_summary`.`project_id`,`issue_rework_summary`.`user_id`,`issue_rework_summary`.`issue_id`,`issue_rework_summary`.`path`,`issue_rework_summary`.`date` FROM `issue_rework_summary` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CustomerID sql.NullString
	var _ProjectID sql.NullString
	var _UserID sql.NullString
	var _IssueID sql.NullString
	var _Path sql.NullString
	var _Date sql.NullInt64
	err := db.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_ProjectID,
		&_UserID,
		&_IssueID,
		&_Path,
		&_Date,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &issueReworkSummary{}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _Checksum.Valid {
		t.SetChecksum(_Checksum.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _ProjectID.Valid {
		t.SetProjectID(_ProjectID.String)
	}
	if _UserID.Valid {
		t.SetUserID(_UserID.String)
	}
	if _IssueID.Valid {
		t.SetIssueID(_IssueID.String)
	}
	if _Path.Valid {
		t.SetPath(_Path.String)
	}
	if _Date.Valid {
		t.SetDate(_Date.Int64)
	}
	return t, nil
}

// FindIssueReworkSummaryByIDTx will find a issueReworkSummary by ID using the provided transaction
func FindIssueReworkSummaryByIDTx(ctx context.Context, tx Tx, value string) (*issueReworkSummary, error) {
	q := "SELECT `issue_rework_summary`.`id`,`issue_rework_summary`.`checksum`,`issue_rework_summary`.`customer_id`,`issue_rework_summary`.`project_id`,`issue_rework_summary`.`user_id`,`issue_rework_summary`.`issue_id`,`issue_rework_summary`.`path`,`issue_rework_summary`.`date` FROM `issue_rework_summary` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CustomerID sql.NullString
	var _ProjectID sql.NullString
	var _UserID sql.NullString
	var _IssueID sql.NullString
	var _Path sql.NullString
	var _Date sql.NullInt64
	err := tx.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_ProjectID,
		&_UserID,
		&_IssueID,
		&_Path,
		&_Date,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &issueReworkSummary{}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _Checksum.Valid {
		t.SetChecksum(_Checksum.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _ProjectID.Valid {
		t.SetProjectID(_ProjectID.String)
	}
	if _UserID.Valid {
		t.SetUserID(_UserID.String)
	}
	if _IssueID.Valid {
		t.SetIssueID(_IssueID.String)
	}
	if _Path.Valid {
		t.SetPath(_Path.String)
	}
	if _Date.Valid {
		t.SetDate(_Date.Int64)
	}
	return t, nil
}

// GetChecksum will return the issueReworkSummary Checksum value
func (t *issueReworkSummary) GetChecksum() string {
	if t.Checksum == nil {
		return ""
	}
	return *t.Checksum
}

// SetChecksum will set the issueReworkSummary Checksum value
func (t *issueReworkSummary) SetChecksum(v string) {
	t.Checksum = &v
}

// GetCustomerID will return the issueReworkSummary CustomerID value
func (t *issueReworkSummary) GetCustomerID() string {
	return t.CustomerID
}

// SetCustomerID will set the issueReworkSummary CustomerID value
func (t *issueReworkSummary) SetCustomerID(v string) {
	t.CustomerID = v
}

// FindIssueReworkSummariesByCustomerID will find all issueReworkSummarys by the CustomerID value
func FindIssueReworkSummariesByCustomerID(ctx context.Context, db DB, value string) ([]*issueReworkSummary, error) {
	q := "SELECT `issue_rework_summary`.`id`,`issue_rework_summary`.`checksum`,`issue_rework_summary`.`customer_id`,`issue_rework_summary`.`project_id`,`issue_rework_summary`.`user_id`,`issue_rework_summary`.`issue_id`,`issue_rework_summary`.`path`,`issue_rework_summary`.`date` FROM `issue_rework_summary` WHERE `customer_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*issueReworkSummary, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _ProjectID sql.NullString
		var _UserID sql.NullString
		var _IssueID sql.NullString
		var _Path sql.NullString
		var _Date sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_ProjectID,
			&_UserID,
			&_IssueID,
			&_Path,
			&_Date,
		)
		if err != nil {
			return nil, err
		}
		t := &issueReworkSummary{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _Path.Valid {
			t.SetPath(_Path.String)
		}
		if _Date.Valid {
			t.SetDate(_Date.Int64)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindIssueReworkSummariesByCustomerIDTx will find all issueReworkSummarys by the CustomerID value using the provided transaction
func FindIssueReworkSummariesByCustomerIDTx(ctx context.Context, tx Tx, value string) ([]*issueReworkSummary, error) {
	q := "SELECT `issue_rework_summary`.`id`,`issue_rework_summary`.`checksum`,`issue_rework_summary`.`customer_id`,`issue_rework_summary`.`project_id`,`issue_rework_summary`.`user_id`,`issue_rework_summary`.`issue_id`,`issue_rework_summary`.`path`,`issue_rework_summary`.`date` FROM `issue_rework_summary` WHERE `customer_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*issueReworkSummary, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _ProjectID sql.NullString
		var _UserID sql.NullString
		var _IssueID sql.NullString
		var _Path sql.NullString
		var _Date sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_ProjectID,
			&_UserID,
			&_IssueID,
			&_Path,
			&_Date,
		)
		if err != nil {
			return nil, err
		}
		t := &issueReworkSummary{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _Path.Valid {
			t.SetPath(_Path.String)
		}
		if _Date.Valid {
			t.SetDate(_Date.Int64)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetProjectID will return the issueReworkSummary ProjectID value
func (t *issueReworkSummary) GetProjectID() string {
	return t.ProjectID
}

// SetProjectID will set the issueReworkSummary ProjectID value
func (t *issueReworkSummary) SetProjectID(v string) {
	t.ProjectID = v
}

// FindIssueReworkSummariesByProjectID will find all issueReworkSummarys by the ProjectID value
func FindIssueReworkSummariesByProjectID(ctx context.Context, db DB, value string) ([]*issueReworkSummary, error) {
	q := "SELECT `issue_rework_summary`.`id`,`issue_rework_summary`.`checksum`,`issue_rework_summary`.`customer_id`,`issue_rework_summary`.`project_id`,`issue_rework_summary`.`user_id`,`issue_rework_summary`.`issue_id`,`issue_rework_summary`.`path`,`issue_rework_summary`.`date` FROM `issue_rework_summary` WHERE `project_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*issueReworkSummary, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _ProjectID sql.NullString
		var _UserID sql.NullString
		var _IssueID sql.NullString
		var _Path sql.NullString
		var _Date sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_ProjectID,
			&_UserID,
			&_IssueID,
			&_Path,
			&_Date,
		)
		if err != nil {
			return nil, err
		}
		t := &issueReworkSummary{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _Path.Valid {
			t.SetPath(_Path.String)
		}
		if _Date.Valid {
			t.SetDate(_Date.Int64)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindIssueReworkSummariesByProjectIDTx will find all issueReworkSummarys by the ProjectID value using the provided transaction
func FindIssueReworkSummariesByProjectIDTx(ctx context.Context, tx Tx, value string) ([]*issueReworkSummary, error) {
	q := "SELECT `issue_rework_summary`.`id`,`issue_rework_summary`.`checksum`,`issue_rework_summary`.`customer_id`,`issue_rework_summary`.`project_id`,`issue_rework_summary`.`user_id`,`issue_rework_summary`.`issue_id`,`issue_rework_summary`.`path`,`issue_rework_summary`.`date` FROM `issue_rework_summary` WHERE `project_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*issueReworkSummary, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _ProjectID sql.NullString
		var _UserID sql.NullString
		var _IssueID sql.NullString
		var _Path sql.NullString
		var _Date sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_ProjectID,
			&_UserID,
			&_IssueID,
			&_Path,
			&_Date,
		)
		if err != nil {
			return nil, err
		}
		t := &issueReworkSummary{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _Path.Valid {
			t.SetPath(_Path.String)
		}
		if _Date.Valid {
			t.SetDate(_Date.Int64)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetUserID will return the issueReworkSummary UserID value
func (t *issueReworkSummary) GetUserID() string {
	if t.UserID == nil {
		return ""
	}
	return *t.UserID
}

// SetUserID will set the issueReworkSummary UserID value
func (t *issueReworkSummary) SetUserID(v string) {
	t.UserID = &v
}

// FindIssueReworkSummariesByUserID will find all issueReworkSummarys by the UserID value
func FindIssueReworkSummariesByUserID(ctx context.Context, db DB, value string) ([]*issueReworkSummary, error) {
	q := "SELECT `issue_rework_summary`.`id`,`issue_rework_summary`.`checksum`,`issue_rework_summary`.`customer_id`,`issue_rework_summary`.`project_id`,`issue_rework_summary`.`user_id`,`issue_rework_summary`.`issue_id`,`issue_rework_summary`.`path`,`issue_rework_summary`.`date` FROM `issue_rework_summary` WHERE `user_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*issueReworkSummary, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _ProjectID sql.NullString
		var _UserID sql.NullString
		var _IssueID sql.NullString
		var _Path sql.NullString
		var _Date sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_ProjectID,
			&_UserID,
			&_IssueID,
			&_Path,
			&_Date,
		)
		if err != nil {
			return nil, err
		}
		t := &issueReworkSummary{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _Path.Valid {
			t.SetPath(_Path.String)
		}
		if _Date.Valid {
			t.SetDate(_Date.Int64)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindIssueReworkSummariesByUserIDTx will find all issueReworkSummarys by the UserID value using the provided transaction
func FindIssueReworkSummariesByUserIDTx(ctx context.Context, tx Tx, value string) ([]*issueReworkSummary, error) {
	q := "SELECT `issue_rework_summary`.`id`,`issue_rework_summary`.`checksum`,`issue_rework_summary`.`customer_id`,`issue_rework_summary`.`project_id`,`issue_rework_summary`.`user_id`,`issue_rework_summary`.`issue_id`,`issue_rework_summary`.`path`,`issue_rework_summary`.`date` FROM `issue_rework_summary` WHERE `user_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*issueReworkSummary, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _ProjectID sql.NullString
		var _UserID sql.NullString
		var _IssueID sql.NullString
		var _Path sql.NullString
		var _Date sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_ProjectID,
			&_UserID,
			&_IssueID,
			&_Path,
			&_Date,
		)
		if err != nil {
			return nil, err
		}
		t := &issueReworkSummary{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _Path.Valid {
			t.SetPath(_Path.String)
		}
		if _Date.Valid {
			t.SetDate(_Date.Int64)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetIssueID will return the issueReworkSummary IssueID value
func (t *issueReworkSummary) GetIssueID() string {
	return t.IssueID
}

// SetIssueID will set the issueReworkSummary IssueID value
func (t *issueReworkSummary) SetIssueID(v string) {
	t.IssueID = v
}

// FindIssueReworkSummariesByIssueID will find all issueReworkSummarys by the IssueID value
func FindIssueReworkSummariesByIssueID(ctx context.Context, db DB, value string) ([]*issueReworkSummary, error) {
	q := "SELECT `issue_rework_summary`.`id`,`issue_rework_summary`.`checksum`,`issue_rework_summary`.`customer_id`,`issue_rework_summary`.`project_id`,`issue_rework_summary`.`user_id`,`issue_rework_summary`.`issue_id`,`issue_rework_summary`.`path`,`issue_rework_summary`.`date` FROM `issue_rework_summary` WHERE `issue_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*issueReworkSummary, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _ProjectID sql.NullString
		var _UserID sql.NullString
		var _IssueID sql.NullString
		var _Path sql.NullString
		var _Date sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_ProjectID,
			&_UserID,
			&_IssueID,
			&_Path,
			&_Date,
		)
		if err != nil {
			return nil, err
		}
		t := &issueReworkSummary{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _Path.Valid {
			t.SetPath(_Path.String)
		}
		if _Date.Valid {
			t.SetDate(_Date.Int64)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindIssueReworkSummariesByIssueIDTx will find all issueReworkSummarys by the IssueID value using the provided transaction
func FindIssueReworkSummariesByIssueIDTx(ctx context.Context, tx Tx, value string) ([]*issueReworkSummary, error) {
	q := "SELECT `issue_rework_summary`.`id`,`issue_rework_summary`.`checksum`,`issue_rework_summary`.`customer_id`,`issue_rework_summary`.`project_id`,`issue_rework_summary`.`user_id`,`issue_rework_summary`.`issue_id`,`issue_rework_summary`.`path`,`issue_rework_summary`.`date` FROM `issue_rework_summary` WHERE `issue_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*issueReworkSummary, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _ProjectID sql.NullString
		var _UserID sql.NullString
		var _IssueID sql.NullString
		var _Path sql.NullString
		var _Date sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_ProjectID,
			&_UserID,
			&_IssueID,
			&_Path,
			&_Date,
		)
		if err != nil {
			return nil, err
		}
		t := &issueReworkSummary{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _Path.Valid {
			t.SetPath(_Path.String)
		}
		if _Date.Valid {
			t.SetDate(_Date.Int64)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetPath will return the issueReworkSummary Path value
func (t *issueReworkSummary) GetPath() string {
	return t.Path
}

// SetPath will set the issueReworkSummary Path value
func (t *issueReworkSummary) SetPath(v string) {
	t.Path = v
}

// GetDate will return the issueReworkSummary Date value
func (t *issueReworkSummary) GetDate() int64 {
	if t.Date == nil {
		return int64(0)
	}
	return *t.Date
}

// SetDate will set the issueReworkSummary Date value
func (t *issueReworkSummary) SetDate(v int64) {
	t.Date = &v
}

func (t *issueReworkSummary) toTimestamp(value time.Time) *timestamp.Timestamp {
	ts, _ := ptypes.TimestampProto(value)
	return ts
}

// DBCreateissueReworkSummaryTable will create the issueReworkSummary table
func DBCreateissueReworkSummaryTable(ctx context.Context, db DB) error {
	q := "CREATE TABLE `issue_rework_summary` (`id` VARCHAR(64) NOT NULL PRIMARY KEY,`checksum` CHAR(64),`customer_id` VARCHAR(64) NOT NULL,`project_id`VARCHAR(64) NOT NULL,`user_id`VARCHAR(64),`issue_id` VARCHAR(64) NOT NULL,`path`VARCHAR(615) NOT NULL,`date`BIGINT UNSIGNED,INDEX issue_rework_summary_customer_id_index (`customer_id`),INDEX issue_rework_summary_project_id_index (`project_id`),INDEX issue_rework_summary_user_id_index (`user_id`),INDEX issue_rework_summary_issue_id_index (`issue_id`),INDEX issue_rework_summary_customer_id_project_id_user_id_index (`customer_id`,`project_id`,`user_id`),INDEX issue_rework_summary_customer_id_user_id_path_index (`customer_id`,`user_id`,`path`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBCreateissueReworkSummaryTableTx will create the issueReworkSummary table using the provided transction
func DBCreateissueReworkSummaryTableTx(ctx context.Context, tx Tx) error {
	q := "CREATE TABLE `issue_rework_summary` (`id` VARCHAR(64) NOT NULL PRIMARY KEY,`checksum` CHAR(64),`customer_id` VARCHAR(64) NOT NULL,`project_id`VARCHAR(64) NOT NULL,`user_id`VARCHAR(64),`issue_id` VARCHAR(64) NOT NULL,`path`VARCHAR(615) NOT NULL,`date`BIGINT UNSIGNED,INDEX issue_rework_summary_customer_id_index (`customer_id`),INDEX issue_rework_summary_project_id_index (`project_id`),INDEX issue_rework_summary_user_id_index (`user_id`),INDEX issue_rework_summary_issue_id_index (`issue_id`),INDEX issue_rework_summary_customer_id_project_id_user_id_index (`customer_id`,`project_id`,`user_id`),INDEX issue_rework_summary_customer_id_user_id_path_index (`customer_id`,`user_id`,`path`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// DBDropissueReworkSummaryTable will drop the issueReworkSummary table
func DBDropissueReworkSummaryTable(ctx context.Context, db DB) error {
	q := "DROP TABLE IF EXISTS `issue_rework_summary`"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBDropissueReworkSummaryTableTx will drop the issueReworkSummary table using the provided transaction
func DBDropissueReworkSummaryTableTx(ctx context.Context, tx Tx) error {
	q := "DROP TABLE IF EXISTS `issue_rework_summary`"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// CalculateChecksum will calculate a checksum of the SHA1 of all field values
func (t *issueReworkSummary) CalculateChecksum() string {
	return orm.HashStrings(
		orm.ToString(t.ID),
		orm.ToString(t.CustomerID),
		orm.ToString(t.ProjectID),
		orm.ToString(t.UserID),
		orm.ToString(t.IssueID),
		orm.ToString(t.Path),
		orm.ToString(t.Date),
	)
}

// DBCreate will create a new issueReworkSummary record in the database
func (t *issueReworkSummary) DBCreate(ctx context.Context, db DB) (sql.Result, error) {
	q := "INSERT INTO `issue_rework_summary` (`issue_rework_summary`.`id`,`issue_rework_summary`.`checksum`,`issue_rework_summary`.`customer_id`,`issue_rework_summary`.`project_id`,`issue_rework_summary`.`user_id`,`issue_rework_summary`.`issue_id`,`issue_rework_summary`.`path`,`issue_rework_summary`.`date`) VALUES (?,?,?,?,?,?,?,?)"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLString(t.UserID),
		orm.ToSQLString(t.IssueID),
		orm.ToSQLString(t.Path),
		orm.ToSQLInt64(t.Date),
	)
}

// DBCreateTx will create a new issueReworkSummary record in the database using the provided transaction
func (t *issueReworkSummary) DBCreateTx(ctx context.Context, tx Tx) (sql.Result, error) {
	q := "INSERT INTO `issue_rework_summary` (`issue_rework_summary`.`id`,`issue_rework_summary`.`checksum`,`issue_rework_summary`.`customer_id`,`issue_rework_summary`.`project_id`,`issue_rework_summary`.`user_id`,`issue_rework_summary`.`issue_id`,`issue_rework_summary`.`path`,`issue_rework_summary`.`date`) VALUES (?,?,?,?,?,?,?,?)"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLString(t.UserID),
		orm.ToSQLString(t.IssueID),
		orm.ToSQLString(t.Path),
		orm.ToSQLInt64(t.Date),
	)
}

// DBCreateIgnoreDuplicate will upsert the issueReworkSummary record in the database
func (t *issueReworkSummary) DBCreateIgnoreDuplicate(ctx context.Context, db DB) (sql.Result, error) {
	q := "INSERT INTO `issue_rework_summary` (`issue_rework_summary`.`id`,`issue_rework_summary`.`checksum`,`issue_rework_summary`.`customer_id`,`issue_rework_summary`.`project_id`,`issue_rework_summary`.`user_id`,`issue_rework_summary`.`issue_id`,`issue_rework_summary`.`path`,`issue_rework_summary`.`date`) VALUES (?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLString(t.UserID),
		orm.ToSQLString(t.IssueID),
		orm.ToSQLString(t.Path),
		orm.ToSQLInt64(t.Date),
	)
}

// DBCreateIgnoreDuplicateTx will upsert the issueReworkSummary record in the database using the provided transaction
func (t *issueReworkSummary) DBCreateIgnoreDuplicateTx(ctx context.Context, tx Tx) (sql.Result, error) {
	q := "INSERT INTO `issue_rework_summary` (`issue_rework_summary`.`id`,`issue_rework_summary`.`checksum`,`issue_rework_summary`.`customer_id`,`issue_rework_summary`.`project_id`,`issue_rework_summary`.`user_id`,`issue_rework_summary`.`issue_id`,`issue_rework_summary`.`path`,`issue_rework_summary`.`date`) VALUES (?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLString(t.UserID),
		orm.ToSQLString(t.IssueID),
		orm.ToSQLString(t.Path),
		orm.ToSQLInt64(t.Date),
	)
}

// DeleteAllIssueReworkSummaries deletes all issueReworkSummary records in the database with optional filters
func DeleteAllIssueReworkSummaries(ctx context.Context, db DB, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(issueReworkSummaryTableName),
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

// DeleteAllIssueReworkSummariesTx deletes all issueReworkSummary records in the database with optional filters using the provided transaction
func DeleteAllIssueReworkSummariesTx(ctx context.Context, tx Tx, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(issueReworkSummaryTableName),
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

// DBDelete will delete this issueReworkSummary record in the database
func (t *issueReworkSummary) DBDelete(ctx context.Context, db DB) (bool, error) {
	q := "DELETE FROM `issue_rework_summary` WHERE `id` = ?"
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

// DBDeleteTx will delete this issueReworkSummary record in the database using the provided transaction
func (t *issueReworkSummary) DBDeleteTx(ctx context.Context, tx Tx) (bool, error) {
	q := "DELETE FROM `issue_rework_summary` WHERE `id` = ?"
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

// DBUpdate will update the issueReworkSummary record in the database
func (t *issueReworkSummary) DBUpdate(ctx context.Context, db DB) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `issue_rework_summary` SET `checksum`=?,`customer_id`=?,`project_id`=?,`user_id`=?,`issue_id`=?,`path`=?,`date`=? WHERE `id`=?"
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLString(t.UserID),
		orm.ToSQLString(t.IssueID),
		orm.ToSQLString(t.Path),
		orm.ToSQLInt64(t.Date),
		orm.ToSQLString(t.ID),
	)
}

// DBUpdateTx will update the issueReworkSummary record in the database using the provided transaction
func (t *issueReworkSummary) DBUpdateTx(ctx context.Context, tx Tx) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `issue_rework_summary` SET `checksum`=?,`customer_id`=?,`project_id`=?,`user_id`=?,`issue_id`=?,`path`=?,`date`=? WHERE `id`=?"
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLString(t.UserID),
		orm.ToSQLString(t.IssueID),
		orm.ToSQLString(t.Path),
		orm.ToSQLInt64(t.Date),
		orm.ToSQLString(t.ID),
	)
}

// DBUpsert will upsert the issueReworkSummary record in the database
func (t *issueReworkSummary) DBUpsert(ctx context.Context, db DB, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `issue_rework_summary` (`issue_rework_summary`.`id`,`issue_rework_summary`.`checksum`,`issue_rework_summary`.`customer_id`,`issue_rework_summary`.`project_id`,`issue_rework_summary`.`user_id`,`issue_rework_summary`.`issue_id`,`issue_rework_summary`.`path`,`issue_rework_summary`.`date`) VALUES (?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `issue_rework_summary` (`issue_rework_summary`.`id`,`issue_rework_summary`.`checksum`,`issue_rework_summary`.`customer_id`,`issue_rework_summary`.`project_id`,`issue_rework_summary`.`user_id`,`issue_rework_summary`.`issue_id`,`issue_rework_summary`.`path`,`issue_rework_summary`.`date`) VALUES (?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`customer_id`=VALUES(`customer_id`),`project_id`=VALUES(`project_id`),`user_id`=VALUES(`user_id`),`issue_id`=VALUES(`issue_id`),`path`=VALUES(`path`),`date`=VALUES(`date`)"
	}
	r, err := db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLString(t.UserID),
		orm.ToSQLString(t.IssueID),
		orm.ToSQLString(t.Path),
		orm.ToSQLInt64(t.Date),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBUpsertTx will upsert the issueReworkSummary record in the database using the provided transaction
func (t *issueReworkSummary) DBUpsertTx(ctx context.Context, tx Tx, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `issue_rework_summary` (`issue_rework_summary`.`id`,`issue_rework_summary`.`checksum`,`issue_rework_summary`.`customer_id`,`issue_rework_summary`.`project_id`,`issue_rework_summary`.`user_id`,`issue_rework_summary`.`issue_id`,`issue_rework_summary`.`path`,`issue_rework_summary`.`date`) VALUES (?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `issue_rework_summary` (`issue_rework_summary`.`id`,`issue_rework_summary`.`checksum`,`issue_rework_summary`.`customer_id`,`issue_rework_summary`.`project_id`,`issue_rework_summary`.`user_id`,`issue_rework_summary`.`issue_id`,`issue_rework_summary`.`path`,`issue_rework_summary`.`date`) VALUES (?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`customer_id`=VALUES(`customer_id`),`project_id`=VALUES(`project_id`),`user_id`=VALUES(`user_id`),`issue_id`=VALUES(`issue_id`),`path`=VALUES(`path`),`date`=VALUES(`date`)"
	}
	r, err := tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLString(t.UserID),
		orm.ToSQLString(t.IssueID),
		orm.ToSQLString(t.Path),
		orm.ToSQLInt64(t.Date),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBFindOne will find a issueReworkSummary record in the database with the primary key
func (t *issueReworkSummary) DBFindOne(ctx context.Context, db DB, value string) (bool, error) {
	q := "SELECT `issue_rework_summary`.`id`,`issue_rework_summary`.`checksum`,`issue_rework_summary`.`customer_id`,`issue_rework_summary`.`project_id`,`issue_rework_summary`.`user_id`,`issue_rework_summary`.`issue_id`,`issue_rework_summary`.`path`,`issue_rework_summary`.`date` FROM `issue_rework_summary` WHERE `id` = ? LIMIT 1"
	row := db.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CustomerID sql.NullString
	var _ProjectID sql.NullString
	var _UserID sql.NullString
	var _IssueID sql.NullString
	var _Path sql.NullString
	var _Date sql.NullInt64
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_ProjectID,
		&_UserID,
		&_IssueID,
		&_Path,
		&_Date,
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
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _ProjectID.Valid {
		t.SetProjectID(_ProjectID.String)
	}
	if _UserID.Valid {
		t.SetUserID(_UserID.String)
	}
	if _IssueID.Valid {
		t.SetIssueID(_IssueID.String)
	}
	if _Path.Valid {
		t.SetPath(_Path.String)
	}
	if _Date.Valid {
		t.SetDate(_Date.Int64)
	}
	return true, nil
}

// DBFindOneTx will find a issueReworkSummary record in the database with the primary key using the provided transaction
func (t *issueReworkSummary) DBFindOneTx(ctx context.Context, tx Tx, value string) (bool, error) {
	q := "SELECT `issue_rework_summary`.`id`,`issue_rework_summary`.`checksum`,`issue_rework_summary`.`customer_id`,`issue_rework_summary`.`project_id`,`issue_rework_summary`.`user_id`,`issue_rework_summary`.`issue_id`,`issue_rework_summary`.`path`,`issue_rework_summary`.`date` FROM `issue_rework_summary` WHERE `id` = ? LIMIT 1"
	row := tx.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CustomerID sql.NullString
	var _ProjectID sql.NullString
	var _UserID sql.NullString
	var _IssueID sql.NullString
	var _Path sql.NullString
	var _Date sql.NullInt64
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_ProjectID,
		&_UserID,
		&_IssueID,
		&_Path,
		&_Date,
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
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _ProjectID.Valid {
		t.SetProjectID(_ProjectID.String)
	}
	if _UserID.Valid {
		t.SetUserID(_UserID.String)
	}
	if _IssueID.Valid {
		t.SetIssueID(_IssueID.String)
	}
	if _Path.Valid {
		t.SetPath(_Path.String)
	}
	if _Date.Valid {
		t.SetDate(_Date.Int64)
	}
	return true, nil
}

// FindIssueReworkSummaries will find a issueReworkSummary record in the database with the provided parameters
func FindIssueReworkSummaries(ctx context.Context, db DB, _params ...interface{}) ([]*issueReworkSummary, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("customer_id"),
		orm.Column("project_id"),
		orm.Column("user_id"),
		orm.Column("issue_id"),
		orm.Column("path"),
		orm.Column("date"),
		orm.Table(issueReworkSummaryTableName),
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
	results := make([]*issueReworkSummary, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _ProjectID sql.NullString
		var _UserID sql.NullString
		var _IssueID sql.NullString
		var _Path sql.NullString
		var _Date sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_ProjectID,
			&_UserID,
			&_IssueID,
			&_Path,
			&_Date,
		)
		if err != nil {
			return nil, err
		}
		t := &issueReworkSummary{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _Path.Valid {
			t.SetPath(_Path.String)
		}
		if _Date.Valid {
			t.SetDate(_Date.Int64)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindIssueReworkSummariesTx will find a issueReworkSummary record in the database with the provided parameters using the provided transaction
func FindIssueReworkSummariesTx(ctx context.Context, tx Tx, _params ...interface{}) ([]*issueReworkSummary, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("customer_id"),
		orm.Column("project_id"),
		orm.Column("user_id"),
		orm.Column("issue_id"),
		orm.Column("path"),
		orm.Column("date"),
		orm.Table(issueReworkSummaryTableName),
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
	results := make([]*issueReworkSummary, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _ProjectID sql.NullString
		var _UserID sql.NullString
		var _IssueID sql.NullString
		var _Path sql.NullString
		var _Date sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_ProjectID,
			&_UserID,
			&_IssueID,
			&_Path,
			&_Date,
		)
		if err != nil {
			return nil, err
		}
		t := &issueReworkSummary{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _Path.Valid {
			t.SetPath(_Path.String)
		}
		if _Date.Valid {
			t.SetDate(_Date.Int64)
		}
		results = append(results, t)
	}
	return results, nil
}

// DBFind will find a issueReworkSummary record in the database with the provided parameters
func (t *issueReworkSummary) DBFind(ctx context.Context, db DB, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("customer_id"),
		orm.Column("project_id"),
		orm.Column("user_id"),
		orm.Column("issue_id"),
		orm.Column("path"),
		orm.Column("date"),
		orm.Table(issueReworkSummaryTableName),
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
	var _CustomerID sql.NullString
	var _ProjectID sql.NullString
	var _UserID sql.NullString
	var _IssueID sql.NullString
	var _Path sql.NullString
	var _Date sql.NullInt64
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_ProjectID,
		&_UserID,
		&_IssueID,
		&_Path,
		&_Date,
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
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _ProjectID.Valid {
		t.SetProjectID(_ProjectID.String)
	}
	if _UserID.Valid {
		t.SetUserID(_UserID.String)
	}
	if _IssueID.Valid {
		t.SetIssueID(_IssueID.String)
	}
	if _Path.Valid {
		t.SetPath(_Path.String)
	}
	if _Date.Valid {
		t.SetDate(_Date.Int64)
	}
	return true, nil
}

// DBFindTx will find a issueReworkSummary record in the database with the provided parameters using the provided transaction
func (t *issueReworkSummary) DBFindTx(ctx context.Context, tx Tx, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("customer_id"),
		orm.Column("project_id"),
		orm.Column("user_id"),
		orm.Column("issue_id"),
		orm.Column("path"),
		orm.Column("date"),
		orm.Table(issueReworkSummaryTableName),
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
	var _CustomerID sql.NullString
	var _ProjectID sql.NullString
	var _UserID sql.NullString
	var _IssueID sql.NullString
	var _Path sql.NullString
	var _Date sql.NullInt64
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_ProjectID,
		&_UserID,
		&_IssueID,
		&_Path,
		&_Date,
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
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _ProjectID.Valid {
		t.SetProjectID(_ProjectID.String)
	}
	if _UserID.Valid {
		t.SetUserID(_UserID.String)
	}
	if _IssueID.Valid {
		t.SetIssueID(_IssueID.String)
	}
	if _Path.Valid {
		t.SetPath(_Path.String)
	}
	if _Date.Valid {
		t.SetDate(_Date.Int64)
	}
	return true, nil
}

// CountIssueReworkSummaries will find the count of issueReworkSummary records in the database
func CountIssueReworkSummaries(ctx context.Context, db DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(issueReworkSummaryTableName),
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

// CountIssueReworkSummariesTx will find the count of issueReworkSummary records in the database using the provided transaction
func CountIssueReworkSummariesTx(ctx context.Context, tx Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(issueReworkSummaryTableName),
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

// DBCount will find the count of issueReworkSummary records in the database
func (t *issueReworkSummary) DBCount(ctx context.Context, db DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(issueReworkSummaryTableName),
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

// DBCountTx will find the count of issueReworkSummary records in the database using the provided transaction
func (t *issueReworkSummary) DBCountTx(ctx context.Context, tx Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(issueReworkSummaryTableName),
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

// DBExists will return true if the issueReworkSummary record exists in the database
func (t *issueReworkSummary) DBExists(ctx context.Context, db DB) (bool, error) {
	q := "SELECT `id` FROM `issue_rework_summary` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := db.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// DBExistsTx will return true if the issueReworkSummary record exists in the database using the provided transaction
func (t *issueReworkSummary) DBExistsTx(ctx context.Context, tx Tx) (bool, error) {
	q := "SELECT `id` FROM `issue_rework_summary` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := tx.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// PrimaryKeyColumn returns the column name for the primary key
func (t *issueReworkSummary) PrimaryKeyColumn() string {
	return issueReworkSummaryColumnID
}

// PrimaryKeyColumnType returns the primary key column Go type as a string
func (t *issueReworkSummary) PrimaryKeyColumnType() string {
	return "string"
}

// PrimaryKey returns the primary key column value
func (t *issueReworkSummary) PrimaryKey() interface{} {
	return t.ID
}
