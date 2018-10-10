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

var _ Model = (*JiraIssueType)(nil)
var _ CSVWriter = (*JiraIssueType)(nil)
var _ JSONWriter = (*JiraIssueType)(nil)
var _ Checksum = (*JiraIssueType)(nil)

// JiraIssueTypeTableName is the name of the table in SQL
const JiraIssueTypeTableName = "jira_issue_type"

var JiraIssueTypeColumns = []string{
	"id",
	"checksum",
	"issue_type_id",
	"name",
	"description",
	"icon_url",
	"url",
	"subtask",
	"customer_id",
}

// JiraIssueType table
type JiraIssueType struct {
	Checksum    *string `json:"checksum,omitempty"`
	CustomerID  string  `json:"customer_id"`
	Description *string `json:"description,omitempty"`
	IconURL     *string `json:"icon_url,omitempty"`
	ID          string  `json:"id"`
	IssueTypeID string  `json:"issue_type_id"`
	Name        string  `json:"name"`
	Subtask     bool    `json:"subtask"`
	URL         *string `json:"url,omitempty"`
}

// TableName returns the SQL table name for JiraIssueType and satifies the Model interface
func (t *JiraIssueType) TableName() string {
	return JiraIssueTypeTableName
}

// ToCSV will serialize the JiraIssueType instance to a CSV compatible array of strings
func (t *JiraIssueType) ToCSV() []string {
	return []string{
		t.ID,
		t.CalculateChecksum(),
		t.IssueTypeID,
		t.Name,
		toCSVString(t.Description),
		toCSVString(t.IconURL),
		toCSVString(t.URL),
		toCSVBool(t.Subtask),
		t.CustomerID,
	}
}

// WriteCSV will serialize the JiraIssueType instance to the writer as CSV and satisfies the CSVWriter interface
func (t *JiraIssueType) WriteCSV(w *csv.Writer) error {
	return w.Write(t.ToCSV())
}

// WriteJSON will serialize the JiraIssueType instance to the writer as JSON and satisfies the JSONWriter interface
func (t *JiraIssueType) WriteJSON(w io.Writer, indent ...bool) error {
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

// NewJiraIssueTypeReader creates a JSON reader which can read in JiraIssueType objects serialized as JSON either as an array, single object or json new lines
// and writes each JiraIssueType to the channel provided
func NewJiraIssueTypeReader(r io.Reader, ch chan<- JiraIssueType) error {
	return orm.Deserialize(r, func(buf json.RawMessage) error {
		dec := json.NewDecoder(bytes.NewBuffer(buf))
		e := JiraIssueType{}
		if err := dec.Decode(&e); err != nil {
			return err
		}
		ch <- e
		return nil
	})
}

// NewCSVJiraIssueTypeReaderDir will read the reader as CSV and emit each record to the channel provided
func NewCSVJiraIssueTypeReader(r io.Reader, ch chan<- JiraIssueType) error {
	cr := csv.NewReader(r)
	for {
		record, err := cr.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		ch <- JiraIssueType{
			ID:          record[0],
			Checksum:    fromStringPointer(record[1]),
			IssueTypeID: record[2],
			Name:        record[3],
			Description: fromStringPointer(record[4]),
			IconURL:     fromStringPointer(record[5]),
			URL:         fromStringPointer(record[6]),
			Subtask:     fromCSVBool(record[7]),
			CustomerID:  record[8],
		}
	}
	return nil
}

// NewCSVJiraIssueTypeReaderFile will read the file as a CSV and emit each record to the channel provided
func NewCSVJiraIssueTypeReaderFile(fp string, ch chan<- JiraIssueType) error {
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
	return NewCSVJiraIssueTypeReader(fc, ch)
}

// NewCSVJiraIssueTypeReaderDir will read the jira_issue_type.csv.gz file as a CSV and emit each record to the channel provided
func NewCSVJiraIssueTypeReaderDir(dir string, ch chan<- JiraIssueType) error {
	return NewCSVJiraIssueTypeReaderFile(filepath.Join(dir, "jira_issue_type.csv.gz"), ch)
}

// JiraIssueTypeCSVDeduper is a function callback which takes the existing value (a) and the new value (b)
// and the return value should be the one to use (or a new one, if applicable). return nil
// to skip processing of this record
type JiraIssueTypeCSVDeduper func(a JiraIssueType, b JiraIssueType) *JiraIssueType

// JiraIssueTypeCSVDedupeDisabled is set on whether the CSV writer should de-dupe values by key
var JiraIssueTypeCSVDedupeDisabled bool

// NewJiraIssueTypeCSVWriterSize creates a batch writer that will write each JiraIssueType into a CSV file
// this method will automatically de-duplicate entries using the primary key. if the checksum
// for a newer item with the same primary key doesn't match a previously sent item, the newer
// one will be used
func NewJiraIssueTypeCSVWriterSize(w io.Writer, size int, dedupers ...JiraIssueTypeCSVDeduper) (chan JiraIssueType, chan bool, error) {
	cw := csv.NewWriter(w)
	ch := make(chan JiraIssueType, size)
	done := make(chan bool)
	go func() {
		defer func() { done <- true }()
		dodedupe := !JiraIssueTypeCSVDedupeDisabled
		var kv map[string]*JiraIssueType
		var deduper JiraIssueTypeCSVDeduper
		if dedupers != nil && len(dedupers) > 0 {
			deduper = dedupers[0]
			dodedupe = true
		}
		if dodedupe {
			kv = make(map[string]*JiraIssueType)
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

// JiraIssueTypeCSVDefaultSize is the default channel buffer size if not provided
var JiraIssueTypeCSVDefaultSize = 100

// NewJiraIssueTypeCSVWriter creates a batch writer that will write each JiraIssueType into a CSV file
func NewJiraIssueTypeCSVWriter(w io.Writer, dedupers ...JiraIssueTypeCSVDeduper) (chan JiraIssueType, chan bool, error) {
	return NewJiraIssueTypeCSVWriterSize(w, JiraIssueTypeCSVDefaultSize, dedupers...)
}

// NewJiraIssueTypeCSVWriterDir creates a batch writer that will write each JiraIssueType into a CSV file named jira_issue_type.csv.gz in dir
func NewJiraIssueTypeCSVWriterDir(dir string, dedupers ...JiraIssueTypeCSVDeduper) (chan JiraIssueType, chan bool, error) {
	return NewJiraIssueTypeCSVWriterFile(filepath.Join(dir, "jira_issue_type.csv.gz"), dedupers...)
}

// NewJiraIssueTypeCSVWriterFile creates a batch writer that will write each JiraIssueType into a CSV file
func NewJiraIssueTypeCSVWriterFile(fn string, dedupers ...JiraIssueTypeCSVDeduper) (chan JiraIssueType, chan bool, error) {
	f, err := os.Create(fn)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening CSV file %s. %v", fn, err)
	}
	var fc io.WriteCloser = f
	if filepath.Ext(fn) == ".gz" {
		w, _ := gzip.NewWriterLevel(f, gzip.BestCompression)
		fc = w
	}
	ch, done, err := NewJiraIssueTypeCSVWriter(fc, dedupers...)
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

type JiraIssueTypeDBAction func(ctx context.Context, db DB, record JiraIssueType) error

// NewJiraIssueTypeDBWriterSize creates a DB writer that will write each issue into the DB
func NewJiraIssueTypeDBWriterSize(ctx context.Context, db DB, errors chan<- error, size int, actions ...JiraIssueTypeDBAction) (chan JiraIssueType, chan bool, error) {
	ch := make(chan JiraIssueType, size)
	done := make(chan bool)
	var action JiraIssueTypeDBAction
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

// NewJiraIssueTypeDBWriter creates a DB writer that will write each issue into the DB
func NewJiraIssueTypeDBWriter(ctx context.Context, db DB, errors chan<- error, actions ...JiraIssueTypeDBAction) (chan JiraIssueType, chan bool, error) {
	return NewJiraIssueTypeDBWriterSize(ctx, db, errors, 100, actions...)
}

// JiraIssueTypeColumnID is the ID SQL column name for the JiraIssueType table
const JiraIssueTypeColumnID = "id"

// JiraIssueTypeEscapedColumnID is the escaped ID SQL column name for the JiraIssueType table
const JiraIssueTypeEscapedColumnID = "`id`"

// JiraIssueTypeColumnChecksum is the Checksum SQL column name for the JiraIssueType table
const JiraIssueTypeColumnChecksum = "checksum"

// JiraIssueTypeEscapedColumnChecksum is the escaped Checksum SQL column name for the JiraIssueType table
const JiraIssueTypeEscapedColumnChecksum = "`checksum`"

// JiraIssueTypeColumnIssueTypeID is the IssueTypeID SQL column name for the JiraIssueType table
const JiraIssueTypeColumnIssueTypeID = "issue_type_id"

// JiraIssueTypeEscapedColumnIssueTypeID is the escaped IssueTypeID SQL column name for the JiraIssueType table
const JiraIssueTypeEscapedColumnIssueTypeID = "`issue_type_id`"

// JiraIssueTypeColumnName is the Name SQL column name for the JiraIssueType table
const JiraIssueTypeColumnName = "name"

// JiraIssueTypeEscapedColumnName is the escaped Name SQL column name for the JiraIssueType table
const JiraIssueTypeEscapedColumnName = "`name`"

// JiraIssueTypeColumnDescription is the Description SQL column name for the JiraIssueType table
const JiraIssueTypeColumnDescription = "description"

// JiraIssueTypeEscapedColumnDescription is the escaped Description SQL column name for the JiraIssueType table
const JiraIssueTypeEscapedColumnDescription = "`description`"

// JiraIssueTypeColumnIconURL is the IconURL SQL column name for the JiraIssueType table
const JiraIssueTypeColumnIconURL = "icon_url"

// JiraIssueTypeEscapedColumnIconURL is the escaped IconURL SQL column name for the JiraIssueType table
const JiraIssueTypeEscapedColumnIconURL = "`icon_url`"

// JiraIssueTypeColumnURL is the URL SQL column name for the JiraIssueType table
const JiraIssueTypeColumnURL = "url"

// JiraIssueTypeEscapedColumnURL is the escaped URL SQL column name for the JiraIssueType table
const JiraIssueTypeEscapedColumnURL = "`url`"

// JiraIssueTypeColumnSubtask is the Subtask SQL column name for the JiraIssueType table
const JiraIssueTypeColumnSubtask = "subtask"

// JiraIssueTypeEscapedColumnSubtask is the escaped Subtask SQL column name for the JiraIssueType table
const JiraIssueTypeEscapedColumnSubtask = "`subtask`"

// JiraIssueTypeColumnCustomerID is the CustomerID SQL column name for the JiraIssueType table
const JiraIssueTypeColumnCustomerID = "customer_id"

// JiraIssueTypeEscapedColumnCustomerID is the escaped CustomerID SQL column name for the JiraIssueType table
const JiraIssueTypeEscapedColumnCustomerID = "`customer_id`"

// GetID will return the JiraIssueType ID value
func (t *JiraIssueType) GetID() string {
	return t.ID
}

// SetID will set the JiraIssueType ID value
func (t *JiraIssueType) SetID(v string) {
	t.ID = v
}

// FindJiraIssueTypeByID will find a JiraIssueType by ID
func FindJiraIssueTypeByID(ctx context.Context, db DB, value string) (*JiraIssueType, error) {
	q := "SELECT * FROM `jira_issue_type` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _IssueTypeID sql.NullString
	var _Name sql.NullString
	var _Description sql.NullString
	var _IconURL sql.NullString
	var _URL sql.NullString
	var _Subtask sql.NullBool
	var _CustomerID sql.NullString
	err := db.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_IssueTypeID,
		&_Name,
		&_Description,
		&_IconURL,
		&_URL,
		&_Subtask,
		&_CustomerID,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &JiraIssueType{}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _Checksum.Valid {
		t.SetChecksum(_Checksum.String)
	}
	if _IssueTypeID.Valid {
		t.SetIssueTypeID(_IssueTypeID.String)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _Description.Valid {
		t.SetDescription(_Description.String)
	}
	if _IconURL.Valid {
		t.SetIconURL(_IconURL.String)
	}
	if _URL.Valid {
		t.SetURL(_URL.String)
	}
	if _Subtask.Valid {
		t.SetSubtask(_Subtask.Bool)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	return t, nil
}

// FindJiraIssueTypeByIDTx will find a JiraIssueType by ID using the provided transaction
func FindJiraIssueTypeByIDTx(ctx context.Context, tx Tx, value string) (*JiraIssueType, error) {
	q := "SELECT * FROM `jira_issue_type` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _IssueTypeID sql.NullString
	var _Name sql.NullString
	var _Description sql.NullString
	var _IconURL sql.NullString
	var _URL sql.NullString
	var _Subtask sql.NullBool
	var _CustomerID sql.NullString
	err := tx.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_IssueTypeID,
		&_Name,
		&_Description,
		&_IconURL,
		&_URL,
		&_Subtask,
		&_CustomerID,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &JiraIssueType{}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _Checksum.Valid {
		t.SetChecksum(_Checksum.String)
	}
	if _IssueTypeID.Valid {
		t.SetIssueTypeID(_IssueTypeID.String)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _Description.Valid {
		t.SetDescription(_Description.String)
	}
	if _IconURL.Valid {
		t.SetIconURL(_IconURL.String)
	}
	if _URL.Valid {
		t.SetURL(_URL.String)
	}
	if _Subtask.Valid {
		t.SetSubtask(_Subtask.Bool)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	return t, nil
}

// GetChecksum will return the JiraIssueType Checksum value
func (t *JiraIssueType) GetChecksum() string {
	if t.Checksum == nil {
		return ""
	}
	return *t.Checksum
}

// SetChecksum will set the JiraIssueType Checksum value
func (t *JiraIssueType) SetChecksum(v string) {
	t.Checksum = &v
}

// GetIssueTypeID will return the JiraIssueType IssueTypeID value
func (t *JiraIssueType) GetIssueTypeID() string {
	return t.IssueTypeID
}

// SetIssueTypeID will set the JiraIssueType IssueTypeID value
func (t *JiraIssueType) SetIssueTypeID(v string) {
	t.IssueTypeID = v
}

// GetName will return the JiraIssueType Name value
func (t *JiraIssueType) GetName() string {
	return t.Name
}

// SetName will set the JiraIssueType Name value
func (t *JiraIssueType) SetName(v string) {
	t.Name = v
}

// FindJiraIssueTypesByName will find all JiraIssueTypes by the Name value
func FindJiraIssueTypesByName(ctx context.Context, db DB, value string) ([]*JiraIssueType, error) {
	q := "SELECT * FROM `jira_issue_type` WHERE `name` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraIssueType, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _IssueTypeID sql.NullString
		var _Name sql.NullString
		var _Description sql.NullString
		var _IconURL sql.NullString
		var _URL sql.NullString
		var _Subtask sql.NullBool
		var _CustomerID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_IssueTypeID,
			&_Name,
			&_Description,
			&_IconURL,
			&_URL,
			&_Subtask,
			&_CustomerID,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraIssueType{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _IssueTypeID.Valid {
			t.SetIssueTypeID(_IssueTypeID.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _Description.Valid {
			t.SetDescription(_Description.String)
		}
		if _IconURL.Valid {
			t.SetIconURL(_IconURL.String)
		}
		if _URL.Valid {
			t.SetURL(_URL.String)
		}
		if _Subtask.Valid {
			t.SetSubtask(_Subtask.Bool)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindJiraIssueTypesByNameTx will find all JiraIssueTypes by the Name value using the provided transaction
func FindJiraIssueTypesByNameTx(ctx context.Context, tx Tx, value string) ([]*JiraIssueType, error) {
	q := "SELECT * FROM `jira_issue_type` WHERE `name` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraIssueType, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _IssueTypeID sql.NullString
		var _Name sql.NullString
		var _Description sql.NullString
		var _IconURL sql.NullString
		var _URL sql.NullString
		var _Subtask sql.NullBool
		var _CustomerID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_IssueTypeID,
			&_Name,
			&_Description,
			&_IconURL,
			&_URL,
			&_Subtask,
			&_CustomerID,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraIssueType{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _IssueTypeID.Valid {
			t.SetIssueTypeID(_IssueTypeID.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _Description.Valid {
			t.SetDescription(_Description.String)
		}
		if _IconURL.Valid {
			t.SetIconURL(_IconURL.String)
		}
		if _URL.Valid {
			t.SetURL(_URL.String)
		}
		if _Subtask.Valid {
			t.SetSubtask(_Subtask.Bool)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetDescription will return the JiraIssueType Description value
func (t *JiraIssueType) GetDescription() string {
	if t.Description == nil {
		return ""
	}
	return *t.Description
}

// SetDescription will set the JiraIssueType Description value
func (t *JiraIssueType) SetDescription(v string) {
	t.Description = &v
}

// GetIconURL will return the JiraIssueType IconURL value
func (t *JiraIssueType) GetIconURL() string {
	if t.IconURL == nil {
		return ""
	}
	return *t.IconURL
}

// SetIconURL will set the JiraIssueType IconURL value
func (t *JiraIssueType) SetIconURL(v string) {
	t.IconURL = &v
}

// GetURL will return the JiraIssueType URL value
func (t *JiraIssueType) GetURL() string {
	if t.URL == nil {
		return ""
	}
	return *t.URL
}

// SetURL will set the JiraIssueType URL value
func (t *JiraIssueType) SetURL(v string) {
	t.URL = &v
}

// GetSubtask will return the JiraIssueType Subtask value
func (t *JiraIssueType) GetSubtask() bool {
	return t.Subtask
}

// SetSubtask will set the JiraIssueType Subtask value
func (t *JiraIssueType) SetSubtask(v bool) {
	t.Subtask = v
}

// GetCustomerID will return the JiraIssueType CustomerID value
func (t *JiraIssueType) GetCustomerID() string {
	return t.CustomerID
}

// SetCustomerID will set the JiraIssueType CustomerID value
func (t *JiraIssueType) SetCustomerID(v string) {
	t.CustomerID = v
}

// FindJiraIssueTypesByCustomerID will find all JiraIssueTypes by the CustomerID value
func FindJiraIssueTypesByCustomerID(ctx context.Context, db DB, value string) ([]*JiraIssueType, error) {
	q := "SELECT * FROM `jira_issue_type` WHERE `customer_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraIssueType, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _IssueTypeID sql.NullString
		var _Name sql.NullString
		var _Description sql.NullString
		var _IconURL sql.NullString
		var _URL sql.NullString
		var _Subtask sql.NullBool
		var _CustomerID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_IssueTypeID,
			&_Name,
			&_Description,
			&_IconURL,
			&_URL,
			&_Subtask,
			&_CustomerID,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraIssueType{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _IssueTypeID.Valid {
			t.SetIssueTypeID(_IssueTypeID.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _Description.Valid {
			t.SetDescription(_Description.String)
		}
		if _IconURL.Valid {
			t.SetIconURL(_IconURL.String)
		}
		if _URL.Valid {
			t.SetURL(_URL.String)
		}
		if _Subtask.Valid {
			t.SetSubtask(_Subtask.Bool)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindJiraIssueTypesByCustomerIDTx will find all JiraIssueTypes by the CustomerID value using the provided transaction
func FindJiraIssueTypesByCustomerIDTx(ctx context.Context, tx Tx, value string) ([]*JiraIssueType, error) {
	q := "SELECT * FROM `jira_issue_type` WHERE `customer_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraIssueType, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _IssueTypeID sql.NullString
		var _Name sql.NullString
		var _Description sql.NullString
		var _IconURL sql.NullString
		var _URL sql.NullString
		var _Subtask sql.NullBool
		var _CustomerID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_IssueTypeID,
			&_Name,
			&_Description,
			&_IconURL,
			&_URL,
			&_Subtask,
			&_CustomerID,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraIssueType{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _IssueTypeID.Valid {
			t.SetIssueTypeID(_IssueTypeID.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _Description.Valid {
			t.SetDescription(_Description.String)
		}
		if _IconURL.Valid {
			t.SetIconURL(_IconURL.String)
		}
		if _URL.Valid {
			t.SetURL(_URL.String)
		}
		if _Subtask.Valid {
			t.SetSubtask(_Subtask.Bool)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

func (t *JiraIssueType) toTimestamp(value time.Time) *timestamp.Timestamp {
	ts, _ := ptypes.TimestampProto(value)
	return ts
}

// DBCreateJiraIssueTypeTable will create the JiraIssueType table
func DBCreateJiraIssueTypeTable(ctx context.Context, db DB) error {
	q := "CREATE TABLE `jira_issue_type` (`id`VARCHAR(64) NOT NULL PRIMARY KEY,`checksum`CHAR(64),`issue_type_id` VARCHAR(64) NOT NULL,`name` VARCHAR(255) NOT NULL,`description`TEXT,`icon_url`VARCHAR(255),`url` VARCHAR(255),`subtask` BOOL NOT NULL,`customer_id`VARCHAR(64) NOT NULL,INDEX jira_issue_type_name_index (`name`),INDEX jira_issue_type_customer_id_index (`customer_id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBCreateJiraIssueTypeTableTx will create the JiraIssueType table using the provided transction
func DBCreateJiraIssueTypeTableTx(ctx context.Context, tx Tx) error {
	q := "CREATE TABLE `jira_issue_type` (`id`VARCHAR(64) NOT NULL PRIMARY KEY,`checksum`CHAR(64),`issue_type_id` VARCHAR(64) NOT NULL,`name` VARCHAR(255) NOT NULL,`description`TEXT,`icon_url`VARCHAR(255),`url` VARCHAR(255),`subtask` BOOL NOT NULL,`customer_id`VARCHAR(64) NOT NULL,INDEX jira_issue_type_name_index (`name`),INDEX jira_issue_type_customer_id_index (`customer_id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// DBDropJiraIssueTypeTable will drop the JiraIssueType table
func DBDropJiraIssueTypeTable(ctx context.Context, db DB) error {
	q := "DROP TABLE IF EXISTS `jira_issue_type`"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBDropJiraIssueTypeTableTx will drop the JiraIssueType table using the provided transaction
func DBDropJiraIssueTypeTableTx(ctx context.Context, tx Tx) error {
	q := "DROP TABLE IF EXISTS `jira_issue_type`"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// CalculateChecksum will calculate a checksum of the SHA1 of all field values
func (t *JiraIssueType) CalculateChecksum() string {
	return orm.HashStrings(
		orm.ToString(t.ID),
		orm.ToString(t.IssueTypeID),
		orm.ToString(t.Name),
		orm.ToString(t.Description),
		orm.ToString(t.IconURL),
		orm.ToString(t.URL),
		orm.ToString(t.Subtask),
		orm.ToString(t.CustomerID),
	)
}

// DBCreate will create a new JiraIssueType record in the database
func (t *JiraIssueType) DBCreate(ctx context.Context, db DB) (sql.Result, error) {
	q := "INSERT INTO `jira_issue_type` (`jira_issue_type`.`id`,`jira_issue_type`.`checksum`,`jira_issue_type`.`issue_type_id`,`jira_issue_type`.`name`,`jira_issue_type`.`description`,`jira_issue_type`.`icon_url`,`jira_issue_type`.`url`,`jira_issue_type`.`subtask`,`jira_issue_type`.`customer_id`) VALUES (?,?,?,?,?,?,?,?,?)"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.IssueTypeID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.Description),
		orm.ToSQLString(t.IconURL),
		orm.ToSQLString(t.URL),
		orm.ToSQLBool(t.Subtask),
		orm.ToSQLString(t.CustomerID),
	)
}

// DBCreateTx will create a new JiraIssueType record in the database using the provided transaction
func (t *JiraIssueType) DBCreateTx(ctx context.Context, tx Tx) (sql.Result, error) {
	q := "INSERT INTO `jira_issue_type` (`jira_issue_type`.`id`,`jira_issue_type`.`checksum`,`jira_issue_type`.`issue_type_id`,`jira_issue_type`.`name`,`jira_issue_type`.`description`,`jira_issue_type`.`icon_url`,`jira_issue_type`.`url`,`jira_issue_type`.`subtask`,`jira_issue_type`.`customer_id`) VALUES (?,?,?,?,?,?,?,?,?)"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.IssueTypeID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.Description),
		orm.ToSQLString(t.IconURL),
		orm.ToSQLString(t.URL),
		orm.ToSQLBool(t.Subtask),
		orm.ToSQLString(t.CustomerID),
	)
}

// DBCreateIgnoreDuplicate will upsert the JiraIssueType record in the database
func (t *JiraIssueType) DBCreateIgnoreDuplicate(ctx context.Context, db DB) (sql.Result, error) {
	q := "INSERT INTO `jira_issue_type` (`jira_issue_type`.`id`,`jira_issue_type`.`checksum`,`jira_issue_type`.`issue_type_id`,`jira_issue_type`.`name`,`jira_issue_type`.`description`,`jira_issue_type`.`icon_url`,`jira_issue_type`.`url`,`jira_issue_type`.`subtask`,`jira_issue_type`.`customer_id`) VALUES (?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.IssueTypeID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.Description),
		orm.ToSQLString(t.IconURL),
		orm.ToSQLString(t.URL),
		orm.ToSQLBool(t.Subtask),
		orm.ToSQLString(t.CustomerID),
	)
}

// DBCreateIgnoreDuplicateTx will upsert the JiraIssueType record in the database using the provided transaction
func (t *JiraIssueType) DBCreateIgnoreDuplicateTx(ctx context.Context, tx Tx) (sql.Result, error) {
	q := "INSERT INTO `jira_issue_type` (`jira_issue_type`.`id`,`jira_issue_type`.`checksum`,`jira_issue_type`.`issue_type_id`,`jira_issue_type`.`name`,`jira_issue_type`.`description`,`jira_issue_type`.`icon_url`,`jira_issue_type`.`url`,`jira_issue_type`.`subtask`,`jira_issue_type`.`customer_id`) VALUES (?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.IssueTypeID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.Description),
		orm.ToSQLString(t.IconURL),
		orm.ToSQLString(t.URL),
		orm.ToSQLBool(t.Subtask),
		orm.ToSQLString(t.CustomerID),
	)
}

// DeleteAllJiraIssueTypes deletes all JiraIssueType records in the database with optional filters
func DeleteAllJiraIssueTypes(ctx context.Context, db DB, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(JiraIssueTypeTableName),
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

// DeleteAllJiraIssueTypesTx deletes all JiraIssueType records in the database with optional filters using the provided transaction
func DeleteAllJiraIssueTypesTx(ctx context.Context, tx Tx, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(JiraIssueTypeTableName),
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

// DBDelete will delete this JiraIssueType record in the database
func (t *JiraIssueType) DBDelete(ctx context.Context, db DB) (bool, error) {
	q := "DELETE FROM `jira_issue_type` WHERE `id` = ?"
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

// DBDeleteTx will delete this JiraIssueType record in the database using the provided transaction
func (t *JiraIssueType) DBDeleteTx(ctx context.Context, tx Tx) (bool, error) {
	q := "DELETE FROM `jira_issue_type` WHERE `id` = ?"
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

// DBUpdate will update the JiraIssueType record in the database
func (t *JiraIssueType) DBUpdate(ctx context.Context, db DB) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `jira_issue_type` SET `checksum`=?,`issue_type_id`=?,`name`=?,`description`=?,`icon_url`=?,`url`=?,`subtask`=?,`customer_id`=? WHERE `id`=?"
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.IssueTypeID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.Description),
		orm.ToSQLString(t.IconURL),
		orm.ToSQLString(t.URL),
		orm.ToSQLBool(t.Subtask),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.ID),
	)
}

// DBUpdateTx will update the JiraIssueType record in the database using the provided transaction
func (t *JiraIssueType) DBUpdateTx(ctx context.Context, tx Tx) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `jira_issue_type` SET `checksum`=?,`issue_type_id`=?,`name`=?,`description`=?,`icon_url`=?,`url`=?,`subtask`=?,`customer_id`=? WHERE `id`=?"
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.IssueTypeID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.Description),
		orm.ToSQLString(t.IconURL),
		orm.ToSQLString(t.URL),
		orm.ToSQLBool(t.Subtask),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.ID),
	)
}

// DBUpsert will upsert the JiraIssueType record in the database
func (t *JiraIssueType) DBUpsert(ctx context.Context, db DB, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `jira_issue_type` (`jira_issue_type`.`id`,`jira_issue_type`.`checksum`,`jira_issue_type`.`issue_type_id`,`jira_issue_type`.`name`,`jira_issue_type`.`description`,`jira_issue_type`.`icon_url`,`jira_issue_type`.`url`,`jira_issue_type`.`subtask`,`jira_issue_type`.`customer_id`) VALUES (?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `jira_issue_type` (`jira_issue_type`.`id`,`jira_issue_type`.`checksum`,`jira_issue_type`.`issue_type_id`,`jira_issue_type`.`name`,`jira_issue_type`.`description`,`jira_issue_type`.`icon_url`,`jira_issue_type`.`url`,`jira_issue_type`.`subtask`,`jira_issue_type`.`customer_id`) VALUES (?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`issue_type_id`=VALUES(`issue_type_id`),`name`=VALUES(`name`),`description`=VALUES(`description`),`icon_url`=VALUES(`icon_url`),`url`=VALUES(`url`),`subtask`=VALUES(`subtask`),`customer_id`=VALUES(`customer_id`)"
	}
	r, err := db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.IssueTypeID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.Description),
		orm.ToSQLString(t.IconURL),
		orm.ToSQLString(t.URL),
		orm.ToSQLBool(t.Subtask),
		orm.ToSQLString(t.CustomerID),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBUpsertTx will upsert the JiraIssueType record in the database using the provided transaction
func (t *JiraIssueType) DBUpsertTx(ctx context.Context, tx Tx, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `jira_issue_type` (`jira_issue_type`.`id`,`jira_issue_type`.`checksum`,`jira_issue_type`.`issue_type_id`,`jira_issue_type`.`name`,`jira_issue_type`.`description`,`jira_issue_type`.`icon_url`,`jira_issue_type`.`url`,`jira_issue_type`.`subtask`,`jira_issue_type`.`customer_id`) VALUES (?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `jira_issue_type` (`jira_issue_type`.`id`,`jira_issue_type`.`checksum`,`jira_issue_type`.`issue_type_id`,`jira_issue_type`.`name`,`jira_issue_type`.`description`,`jira_issue_type`.`icon_url`,`jira_issue_type`.`url`,`jira_issue_type`.`subtask`,`jira_issue_type`.`customer_id`) VALUES (?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`issue_type_id`=VALUES(`issue_type_id`),`name`=VALUES(`name`),`description`=VALUES(`description`),`icon_url`=VALUES(`icon_url`),`url`=VALUES(`url`),`subtask`=VALUES(`subtask`),`customer_id`=VALUES(`customer_id`)"
	}
	r, err := tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.IssueTypeID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.Description),
		orm.ToSQLString(t.IconURL),
		orm.ToSQLString(t.URL),
		orm.ToSQLBool(t.Subtask),
		orm.ToSQLString(t.CustomerID),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBFindOne will find a JiraIssueType record in the database with the primary key
func (t *JiraIssueType) DBFindOne(ctx context.Context, db DB, value string) (bool, error) {
	q := "SELECT * FROM `jira_issue_type` WHERE `id` = ? LIMIT 1"
	row := db.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _IssueTypeID sql.NullString
	var _Name sql.NullString
	var _Description sql.NullString
	var _IconURL sql.NullString
	var _URL sql.NullString
	var _Subtask sql.NullBool
	var _CustomerID sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_IssueTypeID,
		&_Name,
		&_Description,
		&_IconURL,
		&_URL,
		&_Subtask,
		&_CustomerID,
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
	if _IssueTypeID.Valid {
		t.SetIssueTypeID(_IssueTypeID.String)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _Description.Valid {
		t.SetDescription(_Description.String)
	}
	if _IconURL.Valid {
		t.SetIconURL(_IconURL.String)
	}
	if _URL.Valid {
		t.SetURL(_URL.String)
	}
	if _Subtask.Valid {
		t.SetSubtask(_Subtask.Bool)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	return true, nil
}

// DBFindOneTx will find a JiraIssueType record in the database with the primary key using the provided transaction
func (t *JiraIssueType) DBFindOneTx(ctx context.Context, tx Tx, value string) (bool, error) {
	q := "SELECT * FROM `jira_issue_type` WHERE `id` = ? LIMIT 1"
	row := tx.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _IssueTypeID sql.NullString
	var _Name sql.NullString
	var _Description sql.NullString
	var _IconURL sql.NullString
	var _URL sql.NullString
	var _Subtask sql.NullBool
	var _CustomerID sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_IssueTypeID,
		&_Name,
		&_Description,
		&_IconURL,
		&_URL,
		&_Subtask,
		&_CustomerID,
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
	if _IssueTypeID.Valid {
		t.SetIssueTypeID(_IssueTypeID.String)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _Description.Valid {
		t.SetDescription(_Description.String)
	}
	if _IconURL.Valid {
		t.SetIconURL(_IconURL.String)
	}
	if _URL.Valid {
		t.SetURL(_URL.String)
	}
	if _Subtask.Valid {
		t.SetSubtask(_Subtask.Bool)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	return true, nil
}

// FindJiraIssueTypes will find a JiraIssueType record in the database with the provided parameters
func FindJiraIssueTypes(ctx context.Context, db DB, _params ...interface{}) ([]*JiraIssueType, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("issue_type_id"),
		orm.Column("name"),
		orm.Column("description"),
		orm.Column("icon_url"),
		orm.Column("url"),
		orm.Column("subtask"),
		orm.Column("customer_id"),
		orm.Table(JiraIssueTypeTableName),
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
	results := make([]*JiraIssueType, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _IssueTypeID sql.NullString
		var _Name sql.NullString
		var _Description sql.NullString
		var _IconURL sql.NullString
		var _URL sql.NullString
		var _Subtask sql.NullBool
		var _CustomerID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_IssueTypeID,
			&_Name,
			&_Description,
			&_IconURL,
			&_URL,
			&_Subtask,
			&_CustomerID,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraIssueType{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _IssueTypeID.Valid {
			t.SetIssueTypeID(_IssueTypeID.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _Description.Valid {
			t.SetDescription(_Description.String)
		}
		if _IconURL.Valid {
			t.SetIconURL(_IconURL.String)
		}
		if _URL.Valid {
			t.SetURL(_URL.String)
		}
		if _Subtask.Valid {
			t.SetSubtask(_Subtask.Bool)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindJiraIssueTypesTx will find a JiraIssueType record in the database with the provided parameters using the provided transaction
func FindJiraIssueTypesTx(ctx context.Context, tx Tx, _params ...interface{}) ([]*JiraIssueType, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("issue_type_id"),
		orm.Column("name"),
		orm.Column("description"),
		orm.Column("icon_url"),
		orm.Column("url"),
		orm.Column("subtask"),
		orm.Column("customer_id"),
		orm.Table(JiraIssueTypeTableName),
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
	results := make([]*JiraIssueType, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _IssueTypeID sql.NullString
		var _Name sql.NullString
		var _Description sql.NullString
		var _IconURL sql.NullString
		var _URL sql.NullString
		var _Subtask sql.NullBool
		var _CustomerID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_IssueTypeID,
			&_Name,
			&_Description,
			&_IconURL,
			&_URL,
			&_Subtask,
			&_CustomerID,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraIssueType{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _IssueTypeID.Valid {
			t.SetIssueTypeID(_IssueTypeID.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _Description.Valid {
			t.SetDescription(_Description.String)
		}
		if _IconURL.Valid {
			t.SetIconURL(_IconURL.String)
		}
		if _URL.Valid {
			t.SetURL(_URL.String)
		}
		if _Subtask.Valid {
			t.SetSubtask(_Subtask.Bool)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// DBFind will find a JiraIssueType record in the database with the provided parameters
func (t *JiraIssueType) DBFind(ctx context.Context, db DB, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("issue_type_id"),
		orm.Column("name"),
		orm.Column("description"),
		orm.Column("icon_url"),
		orm.Column("url"),
		orm.Column("subtask"),
		orm.Column("customer_id"),
		orm.Table(JiraIssueTypeTableName),
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
	var _IssueTypeID sql.NullString
	var _Name sql.NullString
	var _Description sql.NullString
	var _IconURL sql.NullString
	var _URL sql.NullString
	var _Subtask sql.NullBool
	var _CustomerID sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_IssueTypeID,
		&_Name,
		&_Description,
		&_IconURL,
		&_URL,
		&_Subtask,
		&_CustomerID,
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
	if _IssueTypeID.Valid {
		t.SetIssueTypeID(_IssueTypeID.String)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _Description.Valid {
		t.SetDescription(_Description.String)
	}
	if _IconURL.Valid {
		t.SetIconURL(_IconURL.String)
	}
	if _URL.Valid {
		t.SetURL(_URL.String)
	}
	if _Subtask.Valid {
		t.SetSubtask(_Subtask.Bool)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	return true, nil
}

// DBFindTx will find a JiraIssueType record in the database with the provided parameters using the provided transaction
func (t *JiraIssueType) DBFindTx(ctx context.Context, tx Tx, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("issue_type_id"),
		orm.Column("name"),
		orm.Column("description"),
		orm.Column("icon_url"),
		orm.Column("url"),
		orm.Column("subtask"),
		orm.Column("customer_id"),
		orm.Table(JiraIssueTypeTableName),
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
	var _IssueTypeID sql.NullString
	var _Name sql.NullString
	var _Description sql.NullString
	var _IconURL sql.NullString
	var _URL sql.NullString
	var _Subtask sql.NullBool
	var _CustomerID sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_IssueTypeID,
		&_Name,
		&_Description,
		&_IconURL,
		&_URL,
		&_Subtask,
		&_CustomerID,
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
	if _IssueTypeID.Valid {
		t.SetIssueTypeID(_IssueTypeID.String)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _Description.Valid {
		t.SetDescription(_Description.String)
	}
	if _IconURL.Valid {
		t.SetIconURL(_IconURL.String)
	}
	if _URL.Valid {
		t.SetURL(_URL.String)
	}
	if _Subtask.Valid {
		t.SetSubtask(_Subtask.Bool)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	return true, nil
}

// CountJiraIssueTypes will find the count of JiraIssueType records in the database
func CountJiraIssueTypes(ctx context.Context, db DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(JiraIssueTypeTableName),
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

// CountJiraIssueTypesTx will find the count of JiraIssueType records in the database using the provided transaction
func CountJiraIssueTypesTx(ctx context.Context, tx Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(JiraIssueTypeTableName),
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

// DBCount will find the count of JiraIssueType records in the database
func (t *JiraIssueType) DBCount(ctx context.Context, db DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(JiraIssueTypeTableName),
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

// DBCountTx will find the count of JiraIssueType records in the database using the provided transaction
func (t *JiraIssueType) DBCountTx(ctx context.Context, tx Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(JiraIssueTypeTableName),
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

// DBExists will return true if the JiraIssueType record exists in the database
func (t *JiraIssueType) DBExists(ctx context.Context, db DB) (bool, error) {
	q := "SELECT * FROM `jira_issue_type` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := db.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// DBExistsTx will return true if the JiraIssueType record exists in the database using the provided transaction
func (t *JiraIssueType) DBExistsTx(ctx context.Context, tx Tx) (bool, error) {
	q := "SELECT * FROM `jira_issue_type` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := tx.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// PrimaryKeyColumn returns the column name for the primary key
func (t *JiraIssueType) PrimaryKeyColumn() string {
	return JiraIssueTypeColumnID
}

// PrimaryKeyColumnType returns the primary key column Go type as a string
func (t *JiraIssueType) PrimaryKeyColumnType() string {
	return "string"
}

// PrimaryKey returns the primary key column value
func (t *JiraIssueType) PrimaryKey() interface{} {
	return t.ID
}
