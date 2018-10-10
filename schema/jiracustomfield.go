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

var _ Model = (*JiraCustomField)(nil)
var _ CSVWriter = (*JiraCustomField)(nil)
var _ JSONWriter = (*JiraCustomField)(nil)
var _ Checksum = (*JiraCustomField)(nil)

// JiraCustomFieldTableName is the name of the table in SQL
const JiraCustomFieldTableName = "jira_custom_field"

var JiraCustomFieldColumns = []string{
	"id",
	"checksum",
	"field_id",
	"name",
	"schema_type",
	"schema_custom",
	"schema_custom_id",
	"customer_id",
}

// JiraCustomField table
type JiraCustomField struct {
	Checksum       *string `json:"checksum,omitempty"`
	CustomerID     string  `json:"customer_id"`
	FieldID        string  `json:"field_id"`
	ID             string  `json:"id"`
	Name           string  `json:"name"`
	SchemaCustom   string  `json:"schema_custom"`
	SchemaCustomID *int64  `json:"schema_custom_id,omitempty"`
	SchemaType     string  `json:"schema_type"`
}

// TableName returns the SQL table name for JiraCustomField and satifies the Model interface
func (t *JiraCustomField) TableName() string {
	return JiraCustomFieldTableName
}

// ToCSV will serialize the JiraCustomField instance to a CSV compatible array of strings
func (t *JiraCustomField) ToCSV() []string {
	return []string{
		t.ID,
		t.CalculateChecksum(),
		t.FieldID,
		t.Name,
		t.SchemaType,
		t.SchemaCustom,
		toCSVString(t.SchemaCustomID),
		t.CustomerID,
	}
}

// WriteCSV will serialize the JiraCustomField instance to the writer as CSV and satisfies the CSVWriter interface
func (t *JiraCustomField) WriteCSV(w *csv.Writer) error {
	return w.Write(t.ToCSV())
}

// WriteJSON will serialize the JiraCustomField instance to the writer as JSON and satisfies the JSONWriter interface
func (t *JiraCustomField) WriteJSON(w io.Writer, indent ...bool) error {
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

// NewJiraCustomFieldReader creates a JSON reader which can read in JiraCustomField objects serialized as JSON either as an array, single object or json new lines
// and writes each JiraCustomField to the channel provided
func NewJiraCustomFieldReader(r io.Reader, ch chan<- JiraCustomField) error {
	return orm.Deserialize(r, func(buf json.RawMessage) error {
		dec := json.NewDecoder(bytes.NewBuffer(buf))
		e := JiraCustomField{}
		if err := dec.Decode(&e); err != nil {
			return err
		}
		ch <- e
		return nil
	})
}

// NewCSVJiraCustomFieldReaderDir will read the reader as CSV and emit each record to the channel provided
func NewCSVJiraCustomFieldReader(r io.Reader, ch chan<- JiraCustomField) error {
	cr := csv.NewReader(r)
	for {
		record, err := cr.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		ch <- JiraCustomField{
			ID:             record[0],
			Checksum:       fromStringPointer(record[1]),
			FieldID:        record[2],
			Name:           record[3],
			SchemaType:     record[4],
			SchemaCustom:   record[5],
			SchemaCustomID: fromCSVInt64Pointer(record[6]),
			CustomerID:     record[7],
		}
	}
	return nil
}

// NewCSVJiraCustomFieldReaderFile will read the file as a CSV and emit each record to the channel provided
func NewCSVJiraCustomFieldReaderFile(fp string, ch chan<- JiraCustomField) error {
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
	return NewCSVJiraCustomFieldReader(fc, ch)
}

// NewCSVJiraCustomFieldReaderDir will read the jira_custom_field.csv.gz file as a CSV and emit each record to the channel provided
func NewCSVJiraCustomFieldReaderDir(dir string, ch chan<- JiraCustomField) error {
	return NewCSVJiraCustomFieldReaderFile(filepath.Join(dir, "jira_custom_field.csv.gz"), ch)
}

// JiraCustomFieldCSVDeduper is a function callback which takes the existing value (a) and the new value (b)
// and the return value should be the one to use (or a new one, if applicable). return nil
// to skip processing of this record
type JiraCustomFieldCSVDeduper func(a JiraCustomField, b JiraCustomField) *JiraCustomField

// JiraCustomFieldCSVDedupeDisabled is set on whether the CSV writer should de-dupe values by key
var JiraCustomFieldCSVDedupeDisabled bool

// NewJiraCustomFieldCSVWriterSize creates a batch writer that will write each JiraCustomField into a CSV file
// this method will automatically de-duplicate entries using the primary key. if the checksum
// for a newer item with the same primary key doesn't match a previously sent item, the newer
// one will be used
func NewJiraCustomFieldCSVWriterSize(w io.Writer, size int, dedupers ...JiraCustomFieldCSVDeduper) (chan JiraCustomField, chan bool, error) {
	cw := csv.NewWriter(w)
	ch := make(chan JiraCustomField, size)
	done := make(chan bool)
	go func() {
		defer func() { done <- true }()
		dodedupe := !JiraCustomFieldCSVDedupeDisabled
		var kv map[string]*JiraCustomField
		var deduper JiraCustomFieldCSVDeduper
		if dedupers != nil && len(dedupers) > 0 {
			deduper = dedupers[0]
			dodedupe = true
		}
		if dodedupe {
			kv = make(map[string]*JiraCustomField)
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

// JiraCustomFieldCSVDefaultSize is the default channel buffer size if not provided
var JiraCustomFieldCSVDefaultSize = 100

// NewJiraCustomFieldCSVWriter creates a batch writer that will write each JiraCustomField into a CSV file
func NewJiraCustomFieldCSVWriter(w io.Writer, dedupers ...JiraCustomFieldCSVDeduper) (chan JiraCustomField, chan bool, error) {
	return NewJiraCustomFieldCSVWriterSize(w, JiraCustomFieldCSVDefaultSize, dedupers...)
}

// NewJiraCustomFieldCSVWriterDir creates a batch writer that will write each JiraCustomField into a CSV file named jira_custom_field.csv.gz in dir
func NewJiraCustomFieldCSVWriterDir(dir string, dedupers ...JiraCustomFieldCSVDeduper) (chan JiraCustomField, chan bool, error) {
	return NewJiraCustomFieldCSVWriterFile(filepath.Join(dir, "jira_custom_field.csv.gz"), dedupers...)
}

// NewJiraCustomFieldCSVWriterFile creates a batch writer that will write each JiraCustomField into a CSV file
func NewJiraCustomFieldCSVWriterFile(fn string, dedupers ...JiraCustomFieldCSVDeduper) (chan JiraCustomField, chan bool, error) {
	f, err := os.Create(fn)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening CSV file %s. %v", fn, err)
	}
	var fc io.WriteCloser = f
	if filepath.Ext(fn) == ".gz" {
		w, _ := gzip.NewWriterLevel(f, gzip.BestCompression)
		fc = w
	}
	ch, done, err := NewJiraCustomFieldCSVWriter(fc, dedupers...)
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

type JiraCustomFieldDBAction func(ctx context.Context, db DB, record JiraCustomField) error

// NewJiraCustomFieldDBWriterSize creates a DB writer that will write each issue into the DB
func NewJiraCustomFieldDBWriterSize(ctx context.Context, db DB, errors chan<- error, size int, actions ...JiraCustomFieldDBAction) (chan JiraCustomField, chan bool, error) {
	ch := make(chan JiraCustomField, size)
	done := make(chan bool)
	var action JiraCustomFieldDBAction
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

// NewJiraCustomFieldDBWriter creates a DB writer that will write each issue into the DB
func NewJiraCustomFieldDBWriter(ctx context.Context, db DB, errors chan<- error, actions ...JiraCustomFieldDBAction) (chan JiraCustomField, chan bool, error) {
	return NewJiraCustomFieldDBWriterSize(ctx, db, errors, 100, actions...)
}

// JiraCustomFieldColumnID is the ID SQL column name for the JiraCustomField table
const JiraCustomFieldColumnID = "id"

// JiraCustomFieldEscapedColumnID is the escaped ID SQL column name for the JiraCustomField table
const JiraCustomFieldEscapedColumnID = "`id`"

// JiraCustomFieldColumnChecksum is the Checksum SQL column name for the JiraCustomField table
const JiraCustomFieldColumnChecksum = "checksum"

// JiraCustomFieldEscapedColumnChecksum is the escaped Checksum SQL column name for the JiraCustomField table
const JiraCustomFieldEscapedColumnChecksum = "`checksum`"

// JiraCustomFieldColumnFieldID is the FieldID SQL column name for the JiraCustomField table
const JiraCustomFieldColumnFieldID = "field_id"

// JiraCustomFieldEscapedColumnFieldID is the escaped FieldID SQL column name for the JiraCustomField table
const JiraCustomFieldEscapedColumnFieldID = "`field_id`"

// JiraCustomFieldColumnName is the Name SQL column name for the JiraCustomField table
const JiraCustomFieldColumnName = "name"

// JiraCustomFieldEscapedColumnName is the escaped Name SQL column name for the JiraCustomField table
const JiraCustomFieldEscapedColumnName = "`name`"

// JiraCustomFieldColumnSchemaType is the SchemaType SQL column name for the JiraCustomField table
const JiraCustomFieldColumnSchemaType = "schema_type"

// JiraCustomFieldEscapedColumnSchemaType is the escaped SchemaType SQL column name for the JiraCustomField table
const JiraCustomFieldEscapedColumnSchemaType = "`schema_type`"

// JiraCustomFieldColumnSchemaCustom is the SchemaCustom SQL column name for the JiraCustomField table
const JiraCustomFieldColumnSchemaCustom = "schema_custom"

// JiraCustomFieldEscapedColumnSchemaCustom is the escaped SchemaCustom SQL column name for the JiraCustomField table
const JiraCustomFieldEscapedColumnSchemaCustom = "`schema_custom`"

// JiraCustomFieldColumnSchemaCustomID is the SchemaCustomID SQL column name for the JiraCustomField table
const JiraCustomFieldColumnSchemaCustomID = "schema_custom_id"

// JiraCustomFieldEscapedColumnSchemaCustomID is the escaped SchemaCustomID SQL column name for the JiraCustomField table
const JiraCustomFieldEscapedColumnSchemaCustomID = "`schema_custom_id`"

// JiraCustomFieldColumnCustomerID is the CustomerID SQL column name for the JiraCustomField table
const JiraCustomFieldColumnCustomerID = "customer_id"

// JiraCustomFieldEscapedColumnCustomerID is the escaped CustomerID SQL column name for the JiraCustomField table
const JiraCustomFieldEscapedColumnCustomerID = "`customer_id`"

// GetID will return the JiraCustomField ID value
func (t *JiraCustomField) GetID() string {
	return t.ID
}

// SetID will set the JiraCustomField ID value
func (t *JiraCustomField) SetID(v string) {
	t.ID = v
}

// FindJiraCustomFieldByID will find a JiraCustomField by ID
func FindJiraCustomFieldByID(ctx context.Context, db DB, value string) (*JiraCustomField, error) {
	q := "SELECT * FROM `jira_custom_field` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _FieldID sql.NullString
	var _Name sql.NullString
	var _SchemaType sql.NullString
	var _SchemaCustom sql.NullString
	var _SchemaCustomID sql.NullInt64
	var _CustomerID sql.NullString
	err := db.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_FieldID,
		&_Name,
		&_SchemaType,
		&_SchemaCustom,
		&_SchemaCustomID,
		&_CustomerID,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &JiraCustomField{}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _Checksum.Valid {
		t.SetChecksum(_Checksum.String)
	}
	if _FieldID.Valid {
		t.SetFieldID(_FieldID.String)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _SchemaType.Valid {
		t.SetSchemaType(_SchemaType.String)
	}
	if _SchemaCustom.Valid {
		t.SetSchemaCustom(_SchemaCustom.String)
	}
	if _SchemaCustomID.Valid {
		t.SetSchemaCustomID(_SchemaCustomID.Int64)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	return t, nil
}

// FindJiraCustomFieldByIDTx will find a JiraCustomField by ID using the provided transaction
func FindJiraCustomFieldByIDTx(ctx context.Context, tx Tx, value string) (*JiraCustomField, error) {
	q := "SELECT * FROM `jira_custom_field` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _FieldID sql.NullString
	var _Name sql.NullString
	var _SchemaType sql.NullString
	var _SchemaCustom sql.NullString
	var _SchemaCustomID sql.NullInt64
	var _CustomerID sql.NullString
	err := tx.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_FieldID,
		&_Name,
		&_SchemaType,
		&_SchemaCustom,
		&_SchemaCustomID,
		&_CustomerID,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &JiraCustomField{}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _Checksum.Valid {
		t.SetChecksum(_Checksum.String)
	}
	if _FieldID.Valid {
		t.SetFieldID(_FieldID.String)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _SchemaType.Valid {
		t.SetSchemaType(_SchemaType.String)
	}
	if _SchemaCustom.Valid {
		t.SetSchemaCustom(_SchemaCustom.String)
	}
	if _SchemaCustomID.Valid {
		t.SetSchemaCustomID(_SchemaCustomID.Int64)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	return t, nil
}

// GetChecksum will return the JiraCustomField Checksum value
func (t *JiraCustomField) GetChecksum() string {
	if t.Checksum == nil {
		return ""
	}
	return *t.Checksum
}

// SetChecksum will set the JiraCustomField Checksum value
func (t *JiraCustomField) SetChecksum(v string) {
	t.Checksum = &v
}

// GetFieldID will return the JiraCustomField FieldID value
func (t *JiraCustomField) GetFieldID() string {
	return t.FieldID
}

// SetFieldID will set the JiraCustomField FieldID value
func (t *JiraCustomField) SetFieldID(v string) {
	t.FieldID = v
}

// GetName will return the JiraCustomField Name value
func (t *JiraCustomField) GetName() string {
	return t.Name
}

// SetName will set the JiraCustomField Name value
func (t *JiraCustomField) SetName(v string) {
	t.Name = v
}

// FindJiraCustomFieldsByName will find all JiraCustomFields by the Name value
func FindJiraCustomFieldsByName(ctx context.Context, db DB, value string) ([]*JiraCustomField, error) {
	q := "SELECT * FROM `jira_custom_field` WHERE `name` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraCustomField, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _FieldID sql.NullString
		var _Name sql.NullString
		var _SchemaType sql.NullString
		var _SchemaCustom sql.NullString
		var _SchemaCustomID sql.NullInt64
		var _CustomerID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_FieldID,
			&_Name,
			&_SchemaType,
			&_SchemaCustom,
			&_SchemaCustomID,
			&_CustomerID,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraCustomField{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _FieldID.Valid {
			t.SetFieldID(_FieldID.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _SchemaType.Valid {
			t.SetSchemaType(_SchemaType.String)
		}
		if _SchemaCustom.Valid {
			t.SetSchemaCustom(_SchemaCustom.String)
		}
		if _SchemaCustomID.Valid {
			t.SetSchemaCustomID(_SchemaCustomID.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindJiraCustomFieldsByNameTx will find all JiraCustomFields by the Name value using the provided transaction
func FindJiraCustomFieldsByNameTx(ctx context.Context, tx Tx, value string) ([]*JiraCustomField, error) {
	q := "SELECT * FROM `jira_custom_field` WHERE `name` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraCustomField, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _FieldID sql.NullString
		var _Name sql.NullString
		var _SchemaType sql.NullString
		var _SchemaCustom sql.NullString
		var _SchemaCustomID sql.NullInt64
		var _CustomerID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_FieldID,
			&_Name,
			&_SchemaType,
			&_SchemaCustom,
			&_SchemaCustomID,
			&_CustomerID,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraCustomField{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _FieldID.Valid {
			t.SetFieldID(_FieldID.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _SchemaType.Valid {
			t.SetSchemaType(_SchemaType.String)
		}
		if _SchemaCustom.Valid {
			t.SetSchemaCustom(_SchemaCustom.String)
		}
		if _SchemaCustomID.Valid {
			t.SetSchemaCustomID(_SchemaCustomID.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetSchemaType will return the JiraCustomField SchemaType value
func (t *JiraCustomField) GetSchemaType() string {
	return t.SchemaType
}

// SetSchemaType will set the JiraCustomField SchemaType value
func (t *JiraCustomField) SetSchemaType(v string) {
	t.SchemaType = v
}

// GetSchemaCustom will return the JiraCustomField SchemaCustom value
func (t *JiraCustomField) GetSchemaCustom() string {
	return t.SchemaCustom
}

// SetSchemaCustom will set the JiraCustomField SchemaCustom value
func (t *JiraCustomField) SetSchemaCustom(v string) {
	t.SchemaCustom = v
}

// GetSchemaCustomID will return the JiraCustomField SchemaCustomID value
func (t *JiraCustomField) GetSchemaCustomID() int64 {
	if t.SchemaCustomID == nil {
		return int64(0)
	}
	return *t.SchemaCustomID
}

// SetSchemaCustomID will set the JiraCustomField SchemaCustomID value
func (t *JiraCustomField) SetSchemaCustomID(v int64) {
	t.SchemaCustomID = &v
}

// GetCustomerID will return the JiraCustomField CustomerID value
func (t *JiraCustomField) GetCustomerID() string {
	return t.CustomerID
}

// SetCustomerID will set the JiraCustomField CustomerID value
func (t *JiraCustomField) SetCustomerID(v string) {
	t.CustomerID = v
}

// FindJiraCustomFieldsByCustomerID will find all JiraCustomFields by the CustomerID value
func FindJiraCustomFieldsByCustomerID(ctx context.Context, db DB, value string) ([]*JiraCustomField, error) {
	q := "SELECT * FROM `jira_custom_field` WHERE `customer_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraCustomField, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _FieldID sql.NullString
		var _Name sql.NullString
		var _SchemaType sql.NullString
		var _SchemaCustom sql.NullString
		var _SchemaCustomID sql.NullInt64
		var _CustomerID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_FieldID,
			&_Name,
			&_SchemaType,
			&_SchemaCustom,
			&_SchemaCustomID,
			&_CustomerID,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraCustomField{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _FieldID.Valid {
			t.SetFieldID(_FieldID.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _SchemaType.Valid {
			t.SetSchemaType(_SchemaType.String)
		}
		if _SchemaCustom.Valid {
			t.SetSchemaCustom(_SchemaCustom.String)
		}
		if _SchemaCustomID.Valid {
			t.SetSchemaCustomID(_SchemaCustomID.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindJiraCustomFieldsByCustomerIDTx will find all JiraCustomFields by the CustomerID value using the provided transaction
func FindJiraCustomFieldsByCustomerIDTx(ctx context.Context, tx Tx, value string) ([]*JiraCustomField, error) {
	q := "SELECT * FROM `jira_custom_field` WHERE `customer_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraCustomField, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _FieldID sql.NullString
		var _Name sql.NullString
		var _SchemaType sql.NullString
		var _SchemaCustom sql.NullString
		var _SchemaCustomID sql.NullInt64
		var _CustomerID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_FieldID,
			&_Name,
			&_SchemaType,
			&_SchemaCustom,
			&_SchemaCustomID,
			&_CustomerID,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraCustomField{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _FieldID.Valid {
			t.SetFieldID(_FieldID.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _SchemaType.Valid {
			t.SetSchemaType(_SchemaType.String)
		}
		if _SchemaCustom.Valid {
			t.SetSchemaCustom(_SchemaCustom.String)
		}
		if _SchemaCustomID.Valid {
			t.SetSchemaCustomID(_SchemaCustomID.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

func (t *JiraCustomField) toTimestamp(value time.Time) *timestamp.Timestamp {
	ts, _ := ptypes.TimestampProto(value)
	return ts
}

// DBCreateJiraCustomFieldTable will create the JiraCustomField table
func DBCreateJiraCustomFieldTable(ctx context.Context, db DB) error {
	q := "CREATE TABLE `jira_custom_field` (`id`VARCHAR(64) NOT NULL PRIMARY KEY,`checksum`CHAR(64),`field_id`VARCHAR(64) NOT NULL,`name` VARCHAR(255) NOT NULL,`schema_type`VARCHAR(100) NOT NULL,`schema_custom` VARCHAR(255) NOT NULL,`schema_custom_id` BIGINT UNSIGNED,`customer_id`VARCHAR(64) NOT NULL,INDEX jira_custom_field_name_index (`name`),INDEX jira_custom_field_customer_id_index (`customer_id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBCreateJiraCustomFieldTableTx will create the JiraCustomField table using the provided transction
func DBCreateJiraCustomFieldTableTx(ctx context.Context, tx Tx) error {
	q := "CREATE TABLE `jira_custom_field` (`id`VARCHAR(64) NOT NULL PRIMARY KEY,`checksum`CHAR(64),`field_id`VARCHAR(64) NOT NULL,`name` VARCHAR(255) NOT NULL,`schema_type`VARCHAR(100) NOT NULL,`schema_custom` VARCHAR(255) NOT NULL,`schema_custom_id` BIGINT UNSIGNED,`customer_id`VARCHAR(64) NOT NULL,INDEX jira_custom_field_name_index (`name`),INDEX jira_custom_field_customer_id_index (`customer_id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// DBDropJiraCustomFieldTable will drop the JiraCustomField table
func DBDropJiraCustomFieldTable(ctx context.Context, db DB) error {
	q := "DROP TABLE IF EXISTS `jira_custom_field`"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBDropJiraCustomFieldTableTx will drop the JiraCustomField table using the provided transaction
func DBDropJiraCustomFieldTableTx(ctx context.Context, tx Tx) error {
	q := "DROP TABLE IF EXISTS `jira_custom_field`"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// CalculateChecksum will calculate a checksum of the SHA1 of all field values
func (t *JiraCustomField) CalculateChecksum() string {
	return orm.HashStrings(
		orm.ToString(t.ID),
		orm.ToString(t.FieldID),
		orm.ToString(t.Name),
		orm.ToString(t.SchemaType),
		orm.ToString(t.SchemaCustom),
		orm.ToString(t.SchemaCustomID),
		orm.ToString(t.CustomerID),
	)
}

// DBCreate will create a new JiraCustomField record in the database
func (t *JiraCustomField) DBCreate(ctx context.Context, db DB) (sql.Result, error) {
	q := "INSERT INTO `jira_custom_field` (`jira_custom_field`.`id`,`jira_custom_field`.`checksum`,`jira_custom_field`.`field_id`,`jira_custom_field`.`name`,`jira_custom_field`.`schema_type`,`jira_custom_field`.`schema_custom`,`jira_custom_field`.`schema_custom_id`,`jira_custom_field`.`customer_id`) VALUES (?,?,?,?,?,?,?,?)"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.FieldID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.SchemaType),
		orm.ToSQLString(t.SchemaCustom),
		orm.ToSQLInt64(t.SchemaCustomID),
		orm.ToSQLString(t.CustomerID),
	)
}

// DBCreateTx will create a new JiraCustomField record in the database using the provided transaction
func (t *JiraCustomField) DBCreateTx(ctx context.Context, tx Tx) (sql.Result, error) {
	q := "INSERT INTO `jira_custom_field` (`jira_custom_field`.`id`,`jira_custom_field`.`checksum`,`jira_custom_field`.`field_id`,`jira_custom_field`.`name`,`jira_custom_field`.`schema_type`,`jira_custom_field`.`schema_custom`,`jira_custom_field`.`schema_custom_id`,`jira_custom_field`.`customer_id`) VALUES (?,?,?,?,?,?,?,?)"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.FieldID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.SchemaType),
		orm.ToSQLString(t.SchemaCustom),
		orm.ToSQLInt64(t.SchemaCustomID),
		orm.ToSQLString(t.CustomerID),
	)
}

// DBCreateIgnoreDuplicate will upsert the JiraCustomField record in the database
func (t *JiraCustomField) DBCreateIgnoreDuplicate(ctx context.Context, db DB) (sql.Result, error) {
	q := "INSERT INTO `jira_custom_field` (`jira_custom_field`.`id`,`jira_custom_field`.`checksum`,`jira_custom_field`.`field_id`,`jira_custom_field`.`name`,`jira_custom_field`.`schema_type`,`jira_custom_field`.`schema_custom`,`jira_custom_field`.`schema_custom_id`,`jira_custom_field`.`customer_id`) VALUES (?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.FieldID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.SchemaType),
		orm.ToSQLString(t.SchemaCustom),
		orm.ToSQLInt64(t.SchemaCustomID),
		orm.ToSQLString(t.CustomerID),
	)
}

// DBCreateIgnoreDuplicateTx will upsert the JiraCustomField record in the database using the provided transaction
func (t *JiraCustomField) DBCreateIgnoreDuplicateTx(ctx context.Context, tx Tx) (sql.Result, error) {
	q := "INSERT INTO `jira_custom_field` (`jira_custom_field`.`id`,`jira_custom_field`.`checksum`,`jira_custom_field`.`field_id`,`jira_custom_field`.`name`,`jira_custom_field`.`schema_type`,`jira_custom_field`.`schema_custom`,`jira_custom_field`.`schema_custom_id`,`jira_custom_field`.`customer_id`) VALUES (?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.FieldID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.SchemaType),
		orm.ToSQLString(t.SchemaCustom),
		orm.ToSQLInt64(t.SchemaCustomID),
		orm.ToSQLString(t.CustomerID),
	)
}

// DeleteAllJiraCustomFields deletes all JiraCustomField records in the database with optional filters
func DeleteAllJiraCustomFields(ctx context.Context, db DB, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(JiraCustomFieldTableName),
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

// DeleteAllJiraCustomFieldsTx deletes all JiraCustomField records in the database with optional filters using the provided transaction
func DeleteAllJiraCustomFieldsTx(ctx context.Context, tx Tx, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(JiraCustomFieldTableName),
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

// DBDelete will delete this JiraCustomField record in the database
func (t *JiraCustomField) DBDelete(ctx context.Context, db DB) (bool, error) {
	q := "DELETE FROM `jira_custom_field` WHERE `id` = ?"
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

// DBDeleteTx will delete this JiraCustomField record in the database using the provided transaction
func (t *JiraCustomField) DBDeleteTx(ctx context.Context, tx Tx) (bool, error) {
	q := "DELETE FROM `jira_custom_field` WHERE `id` = ?"
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

// DBUpdate will update the JiraCustomField record in the database
func (t *JiraCustomField) DBUpdate(ctx context.Context, db DB) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `jira_custom_field` SET `checksum`=?,`field_id`=?,`name`=?,`schema_type`=?,`schema_custom`=?,`schema_custom_id`=?,`customer_id`=? WHERE `id`=?"
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.FieldID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.SchemaType),
		orm.ToSQLString(t.SchemaCustom),
		orm.ToSQLInt64(t.SchemaCustomID),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.ID),
	)
}

// DBUpdateTx will update the JiraCustomField record in the database using the provided transaction
func (t *JiraCustomField) DBUpdateTx(ctx context.Context, tx Tx) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `jira_custom_field` SET `checksum`=?,`field_id`=?,`name`=?,`schema_type`=?,`schema_custom`=?,`schema_custom_id`=?,`customer_id`=? WHERE `id`=?"
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.FieldID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.SchemaType),
		orm.ToSQLString(t.SchemaCustom),
		orm.ToSQLInt64(t.SchemaCustomID),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.ID),
	)
}

// DBUpsert will upsert the JiraCustomField record in the database
func (t *JiraCustomField) DBUpsert(ctx context.Context, db DB, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `jira_custom_field` (`jira_custom_field`.`id`,`jira_custom_field`.`checksum`,`jira_custom_field`.`field_id`,`jira_custom_field`.`name`,`jira_custom_field`.`schema_type`,`jira_custom_field`.`schema_custom`,`jira_custom_field`.`schema_custom_id`,`jira_custom_field`.`customer_id`) VALUES (?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `jira_custom_field` (`jira_custom_field`.`id`,`jira_custom_field`.`checksum`,`jira_custom_field`.`field_id`,`jira_custom_field`.`name`,`jira_custom_field`.`schema_type`,`jira_custom_field`.`schema_custom`,`jira_custom_field`.`schema_custom_id`,`jira_custom_field`.`customer_id`) VALUES (?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`field_id`=VALUES(`field_id`),`name`=VALUES(`name`),`schema_type`=VALUES(`schema_type`),`schema_custom`=VALUES(`schema_custom`),`schema_custom_id`=VALUES(`schema_custom_id`),`customer_id`=VALUES(`customer_id`)"
	}
	r, err := db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.FieldID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.SchemaType),
		orm.ToSQLString(t.SchemaCustom),
		orm.ToSQLInt64(t.SchemaCustomID),
		orm.ToSQLString(t.CustomerID),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBUpsertTx will upsert the JiraCustomField record in the database using the provided transaction
func (t *JiraCustomField) DBUpsertTx(ctx context.Context, tx Tx, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `jira_custom_field` (`jira_custom_field`.`id`,`jira_custom_field`.`checksum`,`jira_custom_field`.`field_id`,`jira_custom_field`.`name`,`jira_custom_field`.`schema_type`,`jira_custom_field`.`schema_custom`,`jira_custom_field`.`schema_custom_id`,`jira_custom_field`.`customer_id`) VALUES (?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `jira_custom_field` (`jira_custom_field`.`id`,`jira_custom_field`.`checksum`,`jira_custom_field`.`field_id`,`jira_custom_field`.`name`,`jira_custom_field`.`schema_type`,`jira_custom_field`.`schema_custom`,`jira_custom_field`.`schema_custom_id`,`jira_custom_field`.`customer_id`) VALUES (?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`field_id`=VALUES(`field_id`),`name`=VALUES(`name`),`schema_type`=VALUES(`schema_type`),`schema_custom`=VALUES(`schema_custom`),`schema_custom_id`=VALUES(`schema_custom_id`),`customer_id`=VALUES(`customer_id`)"
	}
	r, err := tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.FieldID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.SchemaType),
		orm.ToSQLString(t.SchemaCustom),
		orm.ToSQLInt64(t.SchemaCustomID),
		orm.ToSQLString(t.CustomerID),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBFindOne will find a JiraCustomField record in the database with the primary key
func (t *JiraCustomField) DBFindOne(ctx context.Context, db DB, value string) (bool, error) {
	q := "SELECT * FROM `jira_custom_field` WHERE `id` = ? LIMIT 1"
	row := db.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _FieldID sql.NullString
	var _Name sql.NullString
	var _SchemaType sql.NullString
	var _SchemaCustom sql.NullString
	var _SchemaCustomID sql.NullInt64
	var _CustomerID sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_FieldID,
		&_Name,
		&_SchemaType,
		&_SchemaCustom,
		&_SchemaCustomID,
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
	if _FieldID.Valid {
		t.SetFieldID(_FieldID.String)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _SchemaType.Valid {
		t.SetSchemaType(_SchemaType.String)
	}
	if _SchemaCustom.Valid {
		t.SetSchemaCustom(_SchemaCustom.String)
	}
	if _SchemaCustomID.Valid {
		t.SetSchemaCustomID(_SchemaCustomID.Int64)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	return true, nil
}

// DBFindOneTx will find a JiraCustomField record in the database with the primary key using the provided transaction
func (t *JiraCustomField) DBFindOneTx(ctx context.Context, tx Tx, value string) (bool, error) {
	q := "SELECT * FROM `jira_custom_field` WHERE `id` = ? LIMIT 1"
	row := tx.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _FieldID sql.NullString
	var _Name sql.NullString
	var _SchemaType sql.NullString
	var _SchemaCustom sql.NullString
	var _SchemaCustomID sql.NullInt64
	var _CustomerID sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_FieldID,
		&_Name,
		&_SchemaType,
		&_SchemaCustom,
		&_SchemaCustomID,
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
	if _FieldID.Valid {
		t.SetFieldID(_FieldID.String)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _SchemaType.Valid {
		t.SetSchemaType(_SchemaType.String)
	}
	if _SchemaCustom.Valid {
		t.SetSchemaCustom(_SchemaCustom.String)
	}
	if _SchemaCustomID.Valid {
		t.SetSchemaCustomID(_SchemaCustomID.Int64)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	return true, nil
}

// FindJiraCustomFields will find a JiraCustomField record in the database with the provided parameters
func FindJiraCustomFields(ctx context.Context, db DB, _params ...interface{}) ([]*JiraCustomField, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("field_id"),
		orm.Column("name"),
		orm.Column("schema_type"),
		orm.Column("schema_custom"),
		orm.Column("schema_custom_id"),
		orm.Column("customer_id"),
		orm.Table(JiraCustomFieldTableName),
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
	results := make([]*JiraCustomField, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _FieldID sql.NullString
		var _Name sql.NullString
		var _SchemaType sql.NullString
		var _SchemaCustom sql.NullString
		var _SchemaCustomID sql.NullInt64
		var _CustomerID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_FieldID,
			&_Name,
			&_SchemaType,
			&_SchemaCustom,
			&_SchemaCustomID,
			&_CustomerID,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraCustomField{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _FieldID.Valid {
			t.SetFieldID(_FieldID.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _SchemaType.Valid {
			t.SetSchemaType(_SchemaType.String)
		}
		if _SchemaCustom.Valid {
			t.SetSchemaCustom(_SchemaCustom.String)
		}
		if _SchemaCustomID.Valid {
			t.SetSchemaCustomID(_SchemaCustomID.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindJiraCustomFieldsTx will find a JiraCustomField record in the database with the provided parameters using the provided transaction
func FindJiraCustomFieldsTx(ctx context.Context, tx Tx, _params ...interface{}) ([]*JiraCustomField, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("field_id"),
		orm.Column("name"),
		orm.Column("schema_type"),
		orm.Column("schema_custom"),
		orm.Column("schema_custom_id"),
		orm.Column("customer_id"),
		orm.Table(JiraCustomFieldTableName),
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
	results := make([]*JiraCustomField, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _FieldID sql.NullString
		var _Name sql.NullString
		var _SchemaType sql.NullString
		var _SchemaCustom sql.NullString
		var _SchemaCustomID sql.NullInt64
		var _CustomerID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_FieldID,
			&_Name,
			&_SchemaType,
			&_SchemaCustom,
			&_SchemaCustomID,
			&_CustomerID,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraCustomField{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _FieldID.Valid {
			t.SetFieldID(_FieldID.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _SchemaType.Valid {
			t.SetSchemaType(_SchemaType.String)
		}
		if _SchemaCustom.Valid {
			t.SetSchemaCustom(_SchemaCustom.String)
		}
		if _SchemaCustomID.Valid {
			t.SetSchemaCustomID(_SchemaCustomID.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// DBFind will find a JiraCustomField record in the database with the provided parameters
func (t *JiraCustomField) DBFind(ctx context.Context, db DB, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("field_id"),
		orm.Column("name"),
		orm.Column("schema_type"),
		orm.Column("schema_custom"),
		orm.Column("schema_custom_id"),
		orm.Column("customer_id"),
		orm.Table(JiraCustomFieldTableName),
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
	var _FieldID sql.NullString
	var _Name sql.NullString
	var _SchemaType sql.NullString
	var _SchemaCustom sql.NullString
	var _SchemaCustomID sql.NullInt64
	var _CustomerID sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_FieldID,
		&_Name,
		&_SchemaType,
		&_SchemaCustom,
		&_SchemaCustomID,
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
	if _FieldID.Valid {
		t.SetFieldID(_FieldID.String)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _SchemaType.Valid {
		t.SetSchemaType(_SchemaType.String)
	}
	if _SchemaCustom.Valid {
		t.SetSchemaCustom(_SchemaCustom.String)
	}
	if _SchemaCustomID.Valid {
		t.SetSchemaCustomID(_SchemaCustomID.Int64)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	return true, nil
}

// DBFindTx will find a JiraCustomField record in the database with the provided parameters using the provided transaction
func (t *JiraCustomField) DBFindTx(ctx context.Context, tx Tx, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("field_id"),
		orm.Column("name"),
		orm.Column("schema_type"),
		orm.Column("schema_custom"),
		orm.Column("schema_custom_id"),
		orm.Column("customer_id"),
		orm.Table(JiraCustomFieldTableName),
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
	var _FieldID sql.NullString
	var _Name sql.NullString
	var _SchemaType sql.NullString
	var _SchemaCustom sql.NullString
	var _SchemaCustomID sql.NullInt64
	var _CustomerID sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_FieldID,
		&_Name,
		&_SchemaType,
		&_SchemaCustom,
		&_SchemaCustomID,
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
	if _FieldID.Valid {
		t.SetFieldID(_FieldID.String)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _SchemaType.Valid {
		t.SetSchemaType(_SchemaType.String)
	}
	if _SchemaCustom.Valid {
		t.SetSchemaCustom(_SchemaCustom.String)
	}
	if _SchemaCustomID.Valid {
		t.SetSchemaCustomID(_SchemaCustomID.Int64)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	return true, nil
}

// CountJiraCustomFields will find the count of JiraCustomField records in the database
func CountJiraCustomFields(ctx context.Context, db DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(JiraCustomFieldTableName),
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

// CountJiraCustomFieldsTx will find the count of JiraCustomField records in the database using the provided transaction
func CountJiraCustomFieldsTx(ctx context.Context, tx Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(JiraCustomFieldTableName),
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

// DBCount will find the count of JiraCustomField records in the database
func (t *JiraCustomField) DBCount(ctx context.Context, db DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(JiraCustomFieldTableName),
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

// DBCountTx will find the count of JiraCustomField records in the database using the provided transaction
func (t *JiraCustomField) DBCountTx(ctx context.Context, tx Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(JiraCustomFieldTableName),
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

// DBExists will return true if the JiraCustomField record exists in the database
func (t *JiraCustomField) DBExists(ctx context.Context, db DB) (bool, error) {
	q := "SELECT * FROM `jira_custom_field` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := db.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// DBExistsTx will return true if the JiraCustomField record exists in the database using the provided transaction
func (t *JiraCustomField) DBExistsTx(ctx context.Context, tx Tx) (bool, error) {
	q := "SELECT * FROM `jira_custom_field` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := tx.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// PrimaryKeyColumn returns the column name for the primary key
func (t *JiraCustomField) PrimaryKeyColumn() string {
	return JiraCustomFieldColumnID
}

// PrimaryKeyColumnType returns the primary key column Go type as a string
func (t *JiraCustomField) PrimaryKeyColumnType() string {
	return "string"
}

// PrimaryKey returns the primary key column value
func (t *JiraCustomField) PrimaryKey() interface{} {
	return t.ID
}
