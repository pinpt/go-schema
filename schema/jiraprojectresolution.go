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

var _ Model = (*JiraProjectResolution)(nil)
var _ CSVWriter = (*JiraProjectResolution)(nil)
var _ JSONWriter = (*JiraProjectResolution)(nil)
var _ Checksum = (*JiraProjectResolution)(nil)

// JiraProjectResolutionTableName is the name of the table in SQL
const JiraProjectResolutionTableName = "jira_project_resolution"

var JiraProjectResolutionColumns = []string{
	"id",
	"checksum",
	"resolution_id",
	"project_id",
	"name",
	"description",
	"customer_id",
}

// JiraProjectResolution table
type JiraProjectResolution struct {
	Checksum     *string `json:"checksum,omitempty"`
	CustomerID   string  `json:"customer_id"`
	Description  *string `json:"description,omitempty"`
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	ProjectID    string  `json:"project_id"`
	ResolutionID string  `json:"resolution_id"`
}

// TableName returns the SQL table name for JiraProjectResolution and satifies the Model interface
func (t *JiraProjectResolution) TableName() string {
	return JiraProjectResolutionTableName
}

// ToCSV will serialize the JiraProjectResolution instance to a CSV compatible array of strings
func (t *JiraProjectResolution) ToCSV() []string {
	return []string{
		t.ID,
		t.CalculateChecksum(),
		t.ResolutionID,
		t.ProjectID,
		t.Name,
		toCSVString(t.Description),
		t.CustomerID,
	}
}

// WriteCSV will serialize the JiraProjectResolution instance to the writer as CSV and satisfies the CSVWriter interface
func (t *JiraProjectResolution) WriteCSV(w *csv.Writer) error {
	return w.Write(t.ToCSV())
}

// WriteJSON will serialize the JiraProjectResolution instance to the writer as JSON and satisfies the JSONWriter interface
func (t *JiraProjectResolution) WriteJSON(w io.Writer, indent ...bool) error {
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

// NewJiraProjectResolutionReader creates a JSON reader which can read in JiraProjectResolution objects serialized as JSON either as an array, single object or json new lines
// and writes each JiraProjectResolution to the channel provided
func NewJiraProjectResolutionReader(r io.Reader, ch chan<- JiraProjectResolution) error {
	return orm.Deserialize(r, func(buf json.RawMessage) error {
		dec := json.NewDecoder(bytes.NewBuffer(buf))
		e := JiraProjectResolution{}
		if err := dec.Decode(&e); err != nil {
			return err
		}
		ch <- e
		return nil
	})
}

// NewCSVJiraProjectResolutionReaderDir will read the reader as CSV and emit each record to the channel provided
func NewCSVJiraProjectResolutionReader(r io.Reader, ch chan<- JiraProjectResolution) error {
	cr := csv.NewReader(r)
	for {
		record, err := cr.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		ch <- JiraProjectResolution{
			ID:           record[0],
			Checksum:     fromStringPointer(record[1]),
			ResolutionID: record[2],
			ProjectID:    record[3],
			Name:         record[4],
			Description:  fromStringPointer(record[5]),
			CustomerID:   record[6],
		}
	}
	return nil
}

// NewCSVJiraProjectResolutionReaderFile will read the file as a CSV and emit each record to the channel provided
func NewCSVJiraProjectResolutionReaderFile(fp string, ch chan<- JiraProjectResolution) error {
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
	return NewCSVJiraProjectResolutionReader(fc, ch)
}

// NewCSVJiraProjectResolutionReaderDir will read the jira_project_resolution.csv.gz file as a CSV and emit each record to the channel provided
func NewCSVJiraProjectResolutionReaderDir(dir string, ch chan<- JiraProjectResolution) error {
	return NewCSVJiraProjectResolutionReaderFile(filepath.Join(dir, "jira_project_resolution.csv.gz"), ch)
}

// JiraProjectResolutionCSVDeduper is a function callback which takes the existing value (a) and the new value (b)
// and the return value should be the one to use (or a new one, if applicable). return nil
// to skip processing of this record
type JiraProjectResolutionCSVDeduper func(a JiraProjectResolution, b JiraProjectResolution) *JiraProjectResolution

// JiraProjectResolutionCSVDedupeDisabled is set on whether the CSV writer should de-dupe values by key
var JiraProjectResolutionCSVDedupeDisabled bool

// NewJiraProjectResolutionCSVWriterSize creates a batch writer that will write each JiraProjectResolution into a CSV file
// this method will automatically de-duplicate entries using the primary key. if the checksum
// for a newer item with the same primary key doesn't match a previously sent item, the newer
// one will be used
func NewJiraProjectResolutionCSVWriterSize(w io.Writer, size int, dedupers ...JiraProjectResolutionCSVDeduper) (chan JiraProjectResolution, chan bool, error) {
	cw := csv.NewWriter(w)
	ch := make(chan JiraProjectResolution, size)
	done := make(chan bool)
	go func() {
		defer func() { done <- true }()
		dodedupe := !JiraProjectResolutionCSVDedupeDisabled
		var kv map[string]*JiraProjectResolution
		var deduper JiraProjectResolutionCSVDeduper
		if dedupers != nil && len(dedupers) > 0 {
			deduper = dedupers[0]
			dodedupe = true
		}
		if dodedupe {
			kv = make(map[string]*JiraProjectResolution)
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

// JiraProjectResolutionCSVDefaultSize is the default channel buffer size if not provided
var JiraProjectResolutionCSVDefaultSize = 100

// NewJiraProjectResolutionCSVWriter creates a batch writer that will write each JiraProjectResolution into a CSV file
func NewJiraProjectResolutionCSVWriter(w io.Writer, dedupers ...JiraProjectResolutionCSVDeduper) (chan JiraProjectResolution, chan bool, error) {
	return NewJiraProjectResolutionCSVWriterSize(w, JiraProjectResolutionCSVDefaultSize, dedupers...)
}

// NewJiraProjectResolutionCSVWriterDir creates a batch writer that will write each JiraProjectResolution into a CSV file named jira_project_resolution.csv.gz in dir
func NewJiraProjectResolutionCSVWriterDir(dir string, dedupers ...JiraProjectResolutionCSVDeduper) (chan JiraProjectResolution, chan bool, error) {
	return NewJiraProjectResolutionCSVWriterFile(filepath.Join(dir, "jira_project_resolution.csv.gz"), dedupers...)
}

// NewJiraProjectResolutionCSVWriterFile creates a batch writer that will write each JiraProjectResolution into a CSV file
func NewJiraProjectResolutionCSVWriterFile(fn string, dedupers ...JiraProjectResolutionCSVDeduper) (chan JiraProjectResolution, chan bool, error) {
	f, err := os.Create(fn)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening CSV file %s. %v", fn, err)
	}
	var fc io.WriteCloser = f
	if filepath.Ext(fn) == ".gz" {
		w, _ := gzip.NewWriterLevel(f, gzip.BestCompression)
		fc = w
	}
	ch, done, err := NewJiraProjectResolutionCSVWriter(fc, dedupers...)
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

type JiraProjectResolutionDBAction func(ctx context.Context, db DB, record JiraProjectResolution) error

// NewJiraProjectResolutionDBWriterSize creates a DB writer that will write each issue into the DB
func NewJiraProjectResolutionDBWriterSize(ctx context.Context, db DB, errors chan<- error, size int, actions ...JiraProjectResolutionDBAction) (chan JiraProjectResolution, chan bool, error) {
	ch := make(chan JiraProjectResolution, size)
	done := make(chan bool)
	var action JiraProjectResolutionDBAction
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

// NewJiraProjectResolutionDBWriter creates a DB writer that will write each issue into the DB
func NewJiraProjectResolutionDBWriter(ctx context.Context, db DB, errors chan<- error, actions ...JiraProjectResolutionDBAction) (chan JiraProjectResolution, chan bool, error) {
	return NewJiraProjectResolutionDBWriterSize(ctx, db, errors, 100, actions...)
}

// JiraProjectResolutionColumnID is the ID SQL column name for the JiraProjectResolution table
const JiraProjectResolutionColumnID = "id"

// JiraProjectResolutionEscapedColumnID is the escaped ID SQL column name for the JiraProjectResolution table
const JiraProjectResolutionEscapedColumnID = "`id`"

// JiraProjectResolutionColumnChecksum is the Checksum SQL column name for the JiraProjectResolution table
const JiraProjectResolutionColumnChecksum = "checksum"

// JiraProjectResolutionEscapedColumnChecksum is the escaped Checksum SQL column name for the JiraProjectResolution table
const JiraProjectResolutionEscapedColumnChecksum = "`checksum`"

// JiraProjectResolutionColumnResolutionID is the ResolutionID SQL column name for the JiraProjectResolution table
const JiraProjectResolutionColumnResolutionID = "resolution_id"

// JiraProjectResolutionEscapedColumnResolutionID is the escaped ResolutionID SQL column name for the JiraProjectResolution table
const JiraProjectResolutionEscapedColumnResolutionID = "`resolution_id`"

// JiraProjectResolutionColumnProjectID is the ProjectID SQL column name for the JiraProjectResolution table
const JiraProjectResolutionColumnProjectID = "project_id"

// JiraProjectResolutionEscapedColumnProjectID is the escaped ProjectID SQL column name for the JiraProjectResolution table
const JiraProjectResolutionEscapedColumnProjectID = "`project_id`"

// JiraProjectResolutionColumnName is the Name SQL column name for the JiraProjectResolution table
const JiraProjectResolutionColumnName = "name"

// JiraProjectResolutionEscapedColumnName is the escaped Name SQL column name for the JiraProjectResolution table
const JiraProjectResolutionEscapedColumnName = "`name`"

// JiraProjectResolutionColumnDescription is the Description SQL column name for the JiraProjectResolution table
const JiraProjectResolutionColumnDescription = "description"

// JiraProjectResolutionEscapedColumnDescription is the escaped Description SQL column name for the JiraProjectResolution table
const JiraProjectResolutionEscapedColumnDescription = "`description`"

// JiraProjectResolutionColumnCustomerID is the CustomerID SQL column name for the JiraProjectResolution table
const JiraProjectResolutionColumnCustomerID = "customer_id"

// JiraProjectResolutionEscapedColumnCustomerID is the escaped CustomerID SQL column name for the JiraProjectResolution table
const JiraProjectResolutionEscapedColumnCustomerID = "`customer_id`"

// GetID will return the JiraProjectResolution ID value
func (t *JiraProjectResolution) GetID() string {
	return t.ID
}

// SetID will set the JiraProjectResolution ID value
func (t *JiraProjectResolution) SetID(v string) {
	t.ID = v
}

// FindJiraProjectResolutionByID will find a JiraProjectResolution by ID
func FindJiraProjectResolutionByID(ctx context.Context, db DB, value string) (*JiraProjectResolution, error) {
	q := "SELECT `jira_project_resolution`.`id`,`jira_project_resolution`.`checksum`,`jira_project_resolution`.`resolution_id`,`jira_project_resolution`.`project_id`,`jira_project_resolution`.`name`,`jira_project_resolution`.`description`,`jira_project_resolution`.`customer_id` FROM `jira_project_resolution` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _ResolutionID sql.NullString
	var _ProjectID sql.NullString
	var _Name sql.NullString
	var _Description sql.NullString
	var _CustomerID sql.NullString
	err := db.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_ResolutionID,
		&_ProjectID,
		&_Name,
		&_Description,
		&_CustomerID,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &JiraProjectResolution{}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _Checksum.Valid {
		t.SetChecksum(_Checksum.String)
	}
	if _ResolutionID.Valid {
		t.SetResolutionID(_ResolutionID.String)
	}
	if _ProjectID.Valid {
		t.SetProjectID(_ProjectID.String)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _Description.Valid {
		t.SetDescription(_Description.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	return t, nil
}

// FindJiraProjectResolutionByIDTx will find a JiraProjectResolution by ID using the provided transaction
func FindJiraProjectResolutionByIDTx(ctx context.Context, tx Tx, value string) (*JiraProjectResolution, error) {
	q := "SELECT `jira_project_resolution`.`id`,`jira_project_resolution`.`checksum`,`jira_project_resolution`.`resolution_id`,`jira_project_resolution`.`project_id`,`jira_project_resolution`.`name`,`jira_project_resolution`.`description`,`jira_project_resolution`.`customer_id` FROM `jira_project_resolution` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _ResolutionID sql.NullString
	var _ProjectID sql.NullString
	var _Name sql.NullString
	var _Description sql.NullString
	var _CustomerID sql.NullString
	err := tx.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_ResolutionID,
		&_ProjectID,
		&_Name,
		&_Description,
		&_CustomerID,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &JiraProjectResolution{}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _Checksum.Valid {
		t.SetChecksum(_Checksum.String)
	}
	if _ResolutionID.Valid {
		t.SetResolutionID(_ResolutionID.String)
	}
	if _ProjectID.Valid {
		t.SetProjectID(_ProjectID.String)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _Description.Valid {
		t.SetDescription(_Description.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	return t, nil
}

// GetChecksum will return the JiraProjectResolution Checksum value
func (t *JiraProjectResolution) GetChecksum() string {
	if t.Checksum == nil {
		return ""
	}
	return *t.Checksum
}

// SetChecksum will set the JiraProjectResolution Checksum value
func (t *JiraProjectResolution) SetChecksum(v string) {
	t.Checksum = &v
}

// GetResolutionID will return the JiraProjectResolution ResolutionID value
func (t *JiraProjectResolution) GetResolutionID() string {
	return t.ResolutionID
}

// SetResolutionID will set the JiraProjectResolution ResolutionID value
func (t *JiraProjectResolution) SetResolutionID(v string) {
	t.ResolutionID = v
}

// GetProjectID will return the JiraProjectResolution ProjectID value
func (t *JiraProjectResolution) GetProjectID() string {
	return t.ProjectID
}

// SetProjectID will set the JiraProjectResolution ProjectID value
func (t *JiraProjectResolution) SetProjectID(v string) {
	t.ProjectID = v
}

// FindJiraProjectResolutionsByProjectID will find all JiraProjectResolutions by the ProjectID value
func FindJiraProjectResolutionsByProjectID(ctx context.Context, db DB, value string) ([]*JiraProjectResolution, error) {
	q := "SELECT `jira_project_resolution`.`id`,`jira_project_resolution`.`checksum`,`jira_project_resolution`.`resolution_id`,`jira_project_resolution`.`project_id`,`jira_project_resolution`.`name`,`jira_project_resolution`.`description`,`jira_project_resolution`.`customer_id` FROM `jira_project_resolution` WHERE `project_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraProjectResolution, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _ResolutionID sql.NullString
		var _ProjectID sql.NullString
		var _Name sql.NullString
		var _Description sql.NullString
		var _CustomerID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_ResolutionID,
			&_ProjectID,
			&_Name,
			&_Description,
			&_CustomerID,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraProjectResolution{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _ResolutionID.Valid {
			t.SetResolutionID(_ResolutionID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _Description.Valid {
			t.SetDescription(_Description.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindJiraProjectResolutionsByProjectIDTx will find all JiraProjectResolutions by the ProjectID value using the provided transaction
func FindJiraProjectResolutionsByProjectIDTx(ctx context.Context, tx Tx, value string) ([]*JiraProjectResolution, error) {
	q := "SELECT `jira_project_resolution`.`id`,`jira_project_resolution`.`checksum`,`jira_project_resolution`.`resolution_id`,`jira_project_resolution`.`project_id`,`jira_project_resolution`.`name`,`jira_project_resolution`.`description`,`jira_project_resolution`.`customer_id` FROM `jira_project_resolution` WHERE `project_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraProjectResolution, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _ResolutionID sql.NullString
		var _ProjectID sql.NullString
		var _Name sql.NullString
		var _Description sql.NullString
		var _CustomerID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_ResolutionID,
			&_ProjectID,
			&_Name,
			&_Description,
			&_CustomerID,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraProjectResolution{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _ResolutionID.Valid {
			t.SetResolutionID(_ResolutionID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _Description.Valid {
			t.SetDescription(_Description.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetName will return the JiraProjectResolution Name value
func (t *JiraProjectResolution) GetName() string {
	return t.Name
}

// SetName will set the JiraProjectResolution Name value
func (t *JiraProjectResolution) SetName(v string) {
	t.Name = v
}

// FindJiraProjectResolutionsByName will find all JiraProjectResolutions by the Name value
func FindJiraProjectResolutionsByName(ctx context.Context, db DB, value string) ([]*JiraProjectResolution, error) {
	q := "SELECT `jira_project_resolution`.`id`,`jira_project_resolution`.`checksum`,`jira_project_resolution`.`resolution_id`,`jira_project_resolution`.`project_id`,`jira_project_resolution`.`name`,`jira_project_resolution`.`description`,`jira_project_resolution`.`customer_id` FROM `jira_project_resolution` WHERE `name` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraProjectResolution, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _ResolutionID sql.NullString
		var _ProjectID sql.NullString
		var _Name sql.NullString
		var _Description sql.NullString
		var _CustomerID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_ResolutionID,
			&_ProjectID,
			&_Name,
			&_Description,
			&_CustomerID,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraProjectResolution{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _ResolutionID.Valid {
			t.SetResolutionID(_ResolutionID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _Description.Valid {
			t.SetDescription(_Description.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindJiraProjectResolutionsByNameTx will find all JiraProjectResolutions by the Name value using the provided transaction
func FindJiraProjectResolutionsByNameTx(ctx context.Context, tx Tx, value string) ([]*JiraProjectResolution, error) {
	q := "SELECT `jira_project_resolution`.`id`,`jira_project_resolution`.`checksum`,`jira_project_resolution`.`resolution_id`,`jira_project_resolution`.`project_id`,`jira_project_resolution`.`name`,`jira_project_resolution`.`description`,`jira_project_resolution`.`customer_id` FROM `jira_project_resolution` WHERE `name` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraProjectResolution, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _ResolutionID sql.NullString
		var _ProjectID sql.NullString
		var _Name sql.NullString
		var _Description sql.NullString
		var _CustomerID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_ResolutionID,
			&_ProjectID,
			&_Name,
			&_Description,
			&_CustomerID,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraProjectResolution{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _ResolutionID.Valid {
			t.SetResolutionID(_ResolutionID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _Description.Valid {
			t.SetDescription(_Description.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetDescription will return the JiraProjectResolution Description value
func (t *JiraProjectResolution) GetDescription() string {
	if t.Description == nil {
		return ""
	}
	return *t.Description
}

// SetDescription will set the JiraProjectResolution Description value
func (t *JiraProjectResolution) SetDescription(v string) {
	t.Description = &v
}

// GetCustomerID will return the JiraProjectResolution CustomerID value
func (t *JiraProjectResolution) GetCustomerID() string {
	return t.CustomerID
}

// SetCustomerID will set the JiraProjectResolution CustomerID value
func (t *JiraProjectResolution) SetCustomerID(v string) {
	t.CustomerID = v
}

// FindJiraProjectResolutionsByCustomerID will find all JiraProjectResolutions by the CustomerID value
func FindJiraProjectResolutionsByCustomerID(ctx context.Context, db DB, value string) ([]*JiraProjectResolution, error) {
	q := "SELECT `jira_project_resolution`.`id`,`jira_project_resolution`.`checksum`,`jira_project_resolution`.`resolution_id`,`jira_project_resolution`.`project_id`,`jira_project_resolution`.`name`,`jira_project_resolution`.`description`,`jira_project_resolution`.`customer_id` FROM `jira_project_resolution` WHERE `customer_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraProjectResolution, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _ResolutionID sql.NullString
		var _ProjectID sql.NullString
		var _Name sql.NullString
		var _Description sql.NullString
		var _CustomerID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_ResolutionID,
			&_ProjectID,
			&_Name,
			&_Description,
			&_CustomerID,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraProjectResolution{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _ResolutionID.Valid {
			t.SetResolutionID(_ResolutionID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _Description.Valid {
			t.SetDescription(_Description.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindJiraProjectResolutionsByCustomerIDTx will find all JiraProjectResolutions by the CustomerID value using the provided transaction
func FindJiraProjectResolutionsByCustomerIDTx(ctx context.Context, tx Tx, value string) ([]*JiraProjectResolution, error) {
	q := "SELECT `jira_project_resolution`.`id`,`jira_project_resolution`.`checksum`,`jira_project_resolution`.`resolution_id`,`jira_project_resolution`.`project_id`,`jira_project_resolution`.`name`,`jira_project_resolution`.`description`,`jira_project_resolution`.`customer_id` FROM `jira_project_resolution` WHERE `customer_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraProjectResolution, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _ResolutionID sql.NullString
		var _ProjectID sql.NullString
		var _Name sql.NullString
		var _Description sql.NullString
		var _CustomerID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_ResolutionID,
			&_ProjectID,
			&_Name,
			&_Description,
			&_CustomerID,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraProjectResolution{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _ResolutionID.Valid {
			t.SetResolutionID(_ResolutionID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _Description.Valid {
			t.SetDescription(_Description.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

func (t *JiraProjectResolution) toTimestamp(value time.Time) *timestamp.Timestamp {
	ts, _ := ptypes.TimestampProto(value)
	return ts
}

// DBCreateJiraProjectResolutionTable will create the JiraProjectResolution table
func DBCreateJiraProjectResolutionTable(ctx context.Context, db DB) error {
	q := "CREATE TABLE `jira_project_resolution` (`id` VARCHAR(64) NOT NULL PRIMARY KEY,`checksum` CHAR(64),`resolution_id` VARCHAR(64) NOT NULL,`project_id` VARCHAR(64) NOT NULL,`name` VARCHAR(255) NOT NULL,`description` TEXT,`customer_id` VARCHAR(64) NOT NULL,INDEX jira_project_resolution_project_id_index (`project_id`),INDEX jira_project_resolution_name_index (`name`),INDEX jira_project_resolution_customer_id_index (`customer_id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBCreateJiraProjectResolutionTableTx will create the JiraProjectResolution table using the provided transction
func DBCreateJiraProjectResolutionTableTx(ctx context.Context, tx Tx) error {
	q := "CREATE TABLE `jira_project_resolution` (`id` VARCHAR(64) NOT NULL PRIMARY KEY,`checksum` CHAR(64),`resolution_id` VARCHAR(64) NOT NULL,`project_id` VARCHAR(64) NOT NULL,`name` VARCHAR(255) NOT NULL,`description` TEXT,`customer_id` VARCHAR(64) NOT NULL,INDEX jira_project_resolution_project_id_index (`project_id`),INDEX jira_project_resolution_name_index (`name`),INDEX jira_project_resolution_customer_id_index (`customer_id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// DBDropJiraProjectResolutionTable will drop the JiraProjectResolution table
func DBDropJiraProjectResolutionTable(ctx context.Context, db DB) error {
	q := "DROP TABLE IF EXISTS `jira_project_resolution`"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBDropJiraProjectResolutionTableTx will drop the JiraProjectResolution table using the provided transaction
func DBDropJiraProjectResolutionTableTx(ctx context.Context, tx Tx) error {
	q := "DROP TABLE IF EXISTS `jira_project_resolution`"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// CalculateChecksum will calculate a checksum of the SHA1 of all field values
func (t *JiraProjectResolution) CalculateChecksum() string {
	return orm.HashStrings(
		orm.ToString(t.ID),
		orm.ToString(t.ResolutionID),
		orm.ToString(t.ProjectID),
		orm.ToString(t.Name),
		orm.ToString(t.Description),
		orm.ToString(t.CustomerID),
	)
}

// DBCreate will create a new JiraProjectResolution record in the database
func (t *JiraProjectResolution) DBCreate(ctx context.Context, db DB) (sql.Result, error) {
	q := "INSERT INTO `jira_project_resolution` (`jira_project_resolution`.`id`,`jira_project_resolution`.`checksum`,`jira_project_resolution`.`resolution_id`,`jira_project_resolution`.`project_id`,`jira_project_resolution`.`name`,`jira_project_resolution`.`description`,`jira_project_resolution`.`customer_id`) VALUES (?,?,?,?,?,?,?)"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.ResolutionID),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.Description),
		orm.ToSQLString(t.CustomerID),
	)
}

// DBCreateTx will create a new JiraProjectResolution record in the database using the provided transaction
func (t *JiraProjectResolution) DBCreateTx(ctx context.Context, tx Tx) (sql.Result, error) {
	q := "INSERT INTO `jira_project_resolution` (`jira_project_resolution`.`id`,`jira_project_resolution`.`checksum`,`jira_project_resolution`.`resolution_id`,`jira_project_resolution`.`project_id`,`jira_project_resolution`.`name`,`jira_project_resolution`.`description`,`jira_project_resolution`.`customer_id`) VALUES (?,?,?,?,?,?,?)"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.ResolutionID),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.Description),
		orm.ToSQLString(t.CustomerID),
	)
}

// DBCreateIgnoreDuplicate will upsert the JiraProjectResolution record in the database
func (t *JiraProjectResolution) DBCreateIgnoreDuplicate(ctx context.Context, db DB) (sql.Result, error) {
	q := "INSERT INTO `jira_project_resolution` (`jira_project_resolution`.`id`,`jira_project_resolution`.`checksum`,`jira_project_resolution`.`resolution_id`,`jira_project_resolution`.`project_id`,`jira_project_resolution`.`name`,`jira_project_resolution`.`description`,`jira_project_resolution`.`customer_id`) VALUES (?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.ResolutionID),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.Description),
		orm.ToSQLString(t.CustomerID),
	)
}

// DBCreateIgnoreDuplicateTx will upsert the JiraProjectResolution record in the database using the provided transaction
func (t *JiraProjectResolution) DBCreateIgnoreDuplicateTx(ctx context.Context, tx Tx) (sql.Result, error) {
	q := "INSERT INTO `jira_project_resolution` (`jira_project_resolution`.`id`,`jira_project_resolution`.`checksum`,`jira_project_resolution`.`resolution_id`,`jira_project_resolution`.`project_id`,`jira_project_resolution`.`name`,`jira_project_resolution`.`description`,`jira_project_resolution`.`customer_id`) VALUES (?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.ResolutionID),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.Description),
		orm.ToSQLString(t.CustomerID),
	)
}

// DeleteAllJiraProjectResolutions deletes all JiraProjectResolution records in the database with optional filters
func DeleteAllJiraProjectResolutions(ctx context.Context, db DB, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(JiraProjectResolutionTableName),
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

// DeleteAllJiraProjectResolutionsTx deletes all JiraProjectResolution records in the database with optional filters using the provided transaction
func DeleteAllJiraProjectResolutionsTx(ctx context.Context, tx Tx, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(JiraProjectResolutionTableName),
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

// DBDelete will delete this JiraProjectResolution record in the database
func (t *JiraProjectResolution) DBDelete(ctx context.Context, db DB) (bool, error) {
	q := "DELETE FROM `jira_project_resolution` WHERE `id` = ?"
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

// DBDeleteTx will delete this JiraProjectResolution record in the database using the provided transaction
func (t *JiraProjectResolution) DBDeleteTx(ctx context.Context, tx Tx) (bool, error) {
	q := "DELETE FROM `jira_project_resolution` WHERE `id` = ?"
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

// DBUpdate will update the JiraProjectResolution record in the database
func (t *JiraProjectResolution) DBUpdate(ctx context.Context, db DB) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `jira_project_resolution` SET `checksum`=?,`resolution_id`=?,`project_id`=?,`name`=?,`description`=?,`customer_id`=? WHERE `id`=?"
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.ResolutionID),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.Description),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.ID),
	)
}

// DBUpdateTx will update the JiraProjectResolution record in the database using the provided transaction
func (t *JiraProjectResolution) DBUpdateTx(ctx context.Context, tx Tx) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `jira_project_resolution` SET `checksum`=?,`resolution_id`=?,`project_id`=?,`name`=?,`description`=?,`customer_id`=? WHERE `id`=?"
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.ResolutionID),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.Description),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.ID),
	)
}

// DBUpsert will upsert the JiraProjectResolution record in the database
func (t *JiraProjectResolution) DBUpsert(ctx context.Context, db DB, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `jira_project_resolution` (`jira_project_resolution`.`id`,`jira_project_resolution`.`checksum`,`jira_project_resolution`.`resolution_id`,`jira_project_resolution`.`project_id`,`jira_project_resolution`.`name`,`jira_project_resolution`.`description`,`jira_project_resolution`.`customer_id`) VALUES (?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `jira_project_resolution` (`jira_project_resolution`.`id`,`jira_project_resolution`.`checksum`,`jira_project_resolution`.`resolution_id`,`jira_project_resolution`.`project_id`,`jira_project_resolution`.`name`,`jira_project_resolution`.`description`,`jira_project_resolution`.`customer_id`) VALUES (?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`resolution_id`=VALUES(`resolution_id`),`project_id`=VALUES(`project_id`),`name`=VALUES(`name`),`description`=VALUES(`description`),`customer_id`=VALUES(`customer_id`)"
	}
	r, err := db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.ResolutionID),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.Description),
		orm.ToSQLString(t.CustomerID),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBUpsertTx will upsert the JiraProjectResolution record in the database using the provided transaction
func (t *JiraProjectResolution) DBUpsertTx(ctx context.Context, tx Tx, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `jira_project_resolution` (`jira_project_resolution`.`id`,`jira_project_resolution`.`checksum`,`jira_project_resolution`.`resolution_id`,`jira_project_resolution`.`project_id`,`jira_project_resolution`.`name`,`jira_project_resolution`.`description`,`jira_project_resolution`.`customer_id`) VALUES (?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `jira_project_resolution` (`jira_project_resolution`.`id`,`jira_project_resolution`.`checksum`,`jira_project_resolution`.`resolution_id`,`jira_project_resolution`.`project_id`,`jira_project_resolution`.`name`,`jira_project_resolution`.`description`,`jira_project_resolution`.`customer_id`) VALUES (?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`resolution_id`=VALUES(`resolution_id`),`project_id`=VALUES(`project_id`),`name`=VALUES(`name`),`description`=VALUES(`description`),`customer_id`=VALUES(`customer_id`)"
	}
	r, err := tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.ResolutionID),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.Description),
		orm.ToSQLString(t.CustomerID),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBFindOne will find a JiraProjectResolution record in the database with the primary key
func (t *JiraProjectResolution) DBFindOne(ctx context.Context, db DB, value string) (bool, error) {
	q := "SELECT `jira_project_resolution`.`id`,`jira_project_resolution`.`checksum`,`jira_project_resolution`.`resolution_id`,`jira_project_resolution`.`project_id`,`jira_project_resolution`.`name`,`jira_project_resolution`.`description`,`jira_project_resolution`.`customer_id` FROM `jira_project_resolution` WHERE `id` = ? LIMIT 1"
	row := db.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _ResolutionID sql.NullString
	var _ProjectID sql.NullString
	var _Name sql.NullString
	var _Description sql.NullString
	var _CustomerID sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_ResolutionID,
		&_ProjectID,
		&_Name,
		&_Description,
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
	if _ResolutionID.Valid {
		t.SetResolutionID(_ResolutionID.String)
	}
	if _ProjectID.Valid {
		t.SetProjectID(_ProjectID.String)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _Description.Valid {
		t.SetDescription(_Description.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	return true, nil
}

// DBFindOneTx will find a JiraProjectResolution record in the database with the primary key using the provided transaction
func (t *JiraProjectResolution) DBFindOneTx(ctx context.Context, tx Tx, value string) (bool, error) {
	q := "SELECT `jira_project_resolution`.`id`,`jira_project_resolution`.`checksum`,`jira_project_resolution`.`resolution_id`,`jira_project_resolution`.`project_id`,`jira_project_resolution`.`name`,`jira_project_resolution`.`description`,`jira_project_resolution`.`customer_id` FROM `jira_project_resolution` WHERE `id` = ? LIMIT 1"
	row := tx.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _ResolutionID sql.NullString
	var _ProjectID sql.NullString
	var _Name sql.NullString
	var _Description sql.NullString
	var _CustomerID sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_ResolutionID,
		&_ProjectID,
		&_Name,
		&_Description,
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
	if _ResolutionID.Valid {
		t.SetResolutionID(_ResolutionID.String)
	}
	if _ProjectID.Valid {
		t.SetProjectID(_ProjectID.String)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _Description.Valid {
		t.SetDescription(_Description.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	return true, nil
}

// FindJiraProjectResolutions will find a JiraProjectResolution record in the database with the provided parameters
func FindJiraProjectResolutions(ctx context.Context, db DB, _params ...interface{}) ([]*JiraProjectResolution, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("resolution_id"),
		orm.Column("project_id"),
		orm.Column("name"),
		orm.Column("description"),
		orm.Column("customer_id"),
		orm.Table(JiraProjectResolutionTableName),
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
	results := make([]*JiraProjectResolution, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _ResolutionID sql.NullString
		var _ProjectID sql.NullString
		var _Name sql.NullString
		var _Description sql.NullString
		var _CustomerID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_ResolutionID,
			&_ProjectID,
			&_Name,
			&_Description,
			&_CustomerID,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraProjectResolution{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _ResolutionID.Valid {
			t.SetResolutionID(_ResolutionID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _Description.Valid {
			t.SetDescription(_Description.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindJiraProjectResolutionsTx will find a JiraProjectResolution record in the database with the provided parameters using the provided transaction
func FindJiraProjectResolutionsTx(ctx context.Context, tx Tx, _params ...interface{}) ([]*JiraProjectResolution, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("resolution_id"),
		orm.Column("project_id"),
		orm.Column("name"),
		orm.Column("description"),
		orm.Column("customer_id"),
		orm.Table(JiraProjectResolutionTableName),
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
	results := make([]*JiraProjectResolution, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _ResolutionID sql.NullString
		var _ProjectID sql.NullString
		var _Name sql.NullString
		var _Description sql.NullString
		var _CustomerID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_ResolutionID,
			&_ProjectID,
			&_Name,
			&_Description,
			&_CustomerID,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraProjectResolution{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _ResolutionID.Valid {
			t.SetResolutionID(_ResolutionID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _Description.Valid {
			t.SetDescription(_Description.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// DBFind will find a JiraProjectResolution record in the database with the provided parameters
func (t *JiraProjectResolution) DBFind(ctx context.Context, db DB, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("resolution_id"),
		orm.Column("project_id"),
		orm.Column("name"),
		orm.Column("description"),
		orm.Column("customer_id"),
		orm.Table(JiraProjectResolutionTableName),
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
	var _ResolutionID sql.NullString
	var _ProjectID sql.NullString
	var _Name sql.NullString
	var _Description sql.NullString
	var _CustomerID sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_ResolutionID,
		&_ProjectID,
		&_Name,
		&_Description,
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
	if _ResolutionID.Valid {
		t.SetResolutionID(_ResolutionID.String)
	}
	if _ProjectID.Valid {
		t.SetProjectID(_ProjectID.String)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _Description.Valid {
		t.SetDescription(_Description.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	return true, nil
}

// DBFindTx will find a JiraProjectResolution record in the database with the provided parameters using the provided transaction
func (t *JiraProjectResolution) DBFindTx(ctx context.Context, tx Tx, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("resolution_id"),
		orm.Column("project_id"),
		orm.Column("name"),
		orm.Column("description"),
		orm.Column("customer_id"),
		orm.Table(JiraProjectResolutionTableName),
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
	var _ResolutionID sql.NullString
	var _ProjectID sql.NullString
	var _Name sql.NullString
	var _Description sql.NullString
	var _CustomerID sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_ResolutionID,
		&_ProjectID,
		&_Name,
		&_Description,
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
	if _ResolutionID.Valid {
		t.SetResolutionID(_ResolutionID.String)
	}
	if _ProjectID.Valid {
		t.SetProjectID(_ProjectID.String)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _Description.Valid {
		t.SetDescription(_Description.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	return true, nil
}

// CountJiraProjectResolutions will find the count of JiraProjectResolution records in the database
func CountJiraProjectResolutions(ctx context.Context, db DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(JiraProjectResolutionTableName),
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

// CountJiraProjectResolutionsTx will find the count of JiraProjectResolution records in the database using the provided transaction
func CountJiraProjectResolutionsTx(ctx context.Context, tx Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(JiraProjectResolutionTableName),
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

// DBCount will find the count of JiraProjectResolution records in the database
func (t *JiraProjectResolution) DBCount(ctx context.Context, db DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(JiraProjectResolutionTableName),
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

// DBCountTx will find the count of JiraProjectResolution records in the database using the provided transaction
func (t *JiraProjectResolution) DBCountTx(ctx context.Context, tx Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(JiraProjectResolutionTableName),
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

// DBExists will return true if the JiraProjectResolution record exists in the database
func (t *JiraProjectResolution) DBExists(ctx context.Context, db DB) (bool, error) {
	q := "SELECT `id` FROM `jira_project_resolution` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := db.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// DBExistsTx will return true if the JiraProjectResolution record exists in the database using the provided transaction
func (t *JiraProjectResolution) DBExistsTx(ctx context.Context, tx Tx) (bool, error) {
	q := "SELECT `id` FROM `jira_project_resolution` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := tx.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// PrimaryKeyColumn returns the column name for the primary key
func (t *JiraProjectResolution) PrimaryKeyColumn() string {
	return JiraProjectResolutionColumnID
}

// PrimaryKeyColumnType returns the primary key column Go type as a string
func (t *JiraProjectResolution) PrimaryKeyColumnType() string {
	return "string"
}

// PrimaryKey returns the primary key column value
func (t *JiraProjectResolution) PrimaryKey() interface{} {
	return t.ID
}
