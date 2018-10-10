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

var _ Model = (*JiraCustomFieldValue)(nil)
var _ CSVWriter = (*JiraCustomFieldValue)(nil)
var _ JSONWriter = (*JiraCustomFieldValue)(nil)
var _ Checksum = (*JiraCustomFieldValue)(nil)

// JiraCustomFieldValueTableName is the name of the table in SQL
const JiraCustomFieldValueTableName = "jira_custom_field_value"

var JiraCustomFieldValueColumns = []string{
	"id",
	"checksum",
	"schema_type",
	"value",
	"custom_field_id",
	"customer_id",
}

// JiraCustomFieldValue table
type JiraCustomFieldValue struct {
	Checksum      *string `json:"checksum,omitempty"`
	CustomFieldID string  `json:"custom_field_id"`
	CustomerID    string  `json:"customer_id"`
	ID            string  `json:"id"`
	SchemaType    string  `json:"schema_type"`
	Value         string  `json:"value"`
}

// TableName returns the SQL table name for JiraCustomFieldValue and satifies the Model interface
func (t *JiraCustomFieldValue) TableName() string {
	return JiraCustomFieldValueTableName
}

// ToCSV will serialize the JiraCustomFieldValue instance to a CSV compatible array of strings
func (t *JiraCustomFieldValue) ToCSV() []string {
	return []string{
		t.ID,
		t.CalculateChecksum(),
		t.SchemaType,
		t.Value,
		t.CustomFieldID,
		t.CustomerID,
	}
}

// WriteCSV will serialize the JiraCustomFieldValue instance to the writer as CSV and satisfies the CSVWriter interface
func (t *JiraCustomFieldValue) WriteCSV(w *csv.Writer) error {
	return w.Write(t.ToCSV())
}

// WriteJSON will serialize the JiraCustomFieldValue instance to the writer as JSON and satisfies the JSONWriter interface
func (t *JiraCustomFieldValue) WriteJSON(w io.Writer, indent ...bool) error {
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

// NewJiraCustomFieldValueReader creates a JSON reader which can read in JiraCustomFieldValue objects serialized as JSON either as an array, single object or json new lines
// and writes each JiraCustomFieldValue to the channel provided
func NewJiraCustomFieldValueReader(r io.Reader, ch chan<- JiraCustomFieldValue) error {
	return orm.Deserialize(r, func(buf json.RawMessage) error {
		dec := json.NewDecoder(bytes.NewBuffer(buf))
		e := JiraCustomFieldValue{}
		if err := dec.Decode(&e); err != nil {
			return err
		}
		ch <- e
		return nil
	})
}

// NewCSVJiraCustomFieldValueReaderDir will read the reader as CSV and emit each record to the channel provided
func NewCSVJiraCustomFieldValueReader(r io.Reader, ch chan<- JiraCustomFieldValue) error {
	cr := csv.NewReader(r)
	for {
		record, err := cr.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		ch <- JiraCustomFieldValue{
			ID:            record[0],
			Checksum:      fromStringPointer(record[1]),
			SchemaType:    record[2],
			Value:         record[3],
			CustomFieldID: record[4],
			CustomerID:    record[5],
		}
	}
	return nil
}

// NewCSVJiraCustomFieldValueReaderFile will read the file as a CSV and emit each record to the channel provided
func NewCSVJiraCustomFieldValueReaderFile(fp string, ch chan<- JiraCustomFieldValue) error {
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
	return NewCSVJiraCustomFieldValueReader(fc, ch)
}

// NewCSVJiraCustomFieldValueReaderDir will read the jira_custom_field_value.csv.gz file as a CSV and emit each record to the channel provided
func NewCSVJiraCustomFieldValueReaderDir(dir string, ch chan<- JiraCustomFieldValue) error {
	return NewCSVJiraCustomFieldValueReaderFile(filepath.Join(dir, "jira_custom_field_value.csv.gz"), ch)
}

// JiraCustomFieldValueCSVDeduper is a function callback which takes the existing value (a) and the new value (b)
// and the return value should be the one to use (or a new one, if applicable). return nil
// to skip processing of this record
type JiraCustomFieldValueCSVDeduper func(a JiraCustomFieldValue, b JiraCustomFieldValue) *JiraCustomFieldValue

// JiraCustomFieldValueCSVDedupeDisabled is set on whether the CSV writer should de-dupe values by key
var JiraCustomFieldValueCSVDedupeDisabled bool

// NewJiraCustomFieldValueCSVWriterSize creates a batch writer that will write each JiraCustomFieldValue into a CSV file
// this method will automatically de-duplicate entries using the primary key. if the checksum
// for a newer item with the same primary key doesn't match a previously sent item, the newer
// one will be used
func NewJiraCustomFieldValueCSVWriterSize(w io.Writer, size int, dedupers ...JiraCustomFieldValueCSVDeduper) (chan JiraCustomFieldValue, chan bool, error) {
	cw := csv.NewWriter(w)
	ch := make(chan JiraCustomFieldValue, size)
	done := make(chan bool)
	go func() {
		defer func() { done <- true }()
		dodedupe := !JiraCustomFieldValueCSVDedupeDisabled
		var kv map[string]*JiraCustomFieldValue
		var deduper JiraCustomFieldValueCSVDeduper
		if dedupers != nil && len(dedupers) > 0 {
			deduper = dedupers[0]
			dodedupe = true
		}
		if dodedupe {
			kv = make(map[string]*JiraCustomFieldValue)
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

// JiraCustomFieldValueCSVDefaultSize is the default channel buffer size if not provided
var JiraCustomFieldValueCSVDefaultSize = 100

// NewJiraCustomFieldValueCSVWriter creates a batch writer that will write each JiraCustomFieldValue into a CSV file
func NewJiraCustomFieldValueCSVWriter(w io.Writer, dedupers ...JiraCustomFieldValueCSVDeduper) (chan JiraCustomFieldValue, chan bool, error) {
	return NewJiraCustomFieldValueCSVWriterSize(w, JiraCustomFieldValueCSVDefaultSize, dedupers...)
}

// NewJiraCustomFieldValueCSVWriterDir creates a batch writer that will write each JiraCustomFieldValue into a CSV file named jira_custom_field_value.csv.gz in dir
func NewJiraCustomFieldValueCSVWriterDir(dir string, dedupers ...JiraCustomFieldValueCSVDeduper) (chan JiraCustomFieldValue, chan bool, error) {
	return NewJiraCustomFieldValueCSVWriterFile(filepath.Join(dir, "jira_custom_field_value.csv.gz"), dedupers...)
}

// NewJiraCustomFieldValueCSVWriterFile creates a batch writer that will write each JiraCustomFieldValue into a CSV file
func NewJiraCustomFieldValueCSVWriterFile(fn string, dedupers ...JiraCustomFieldValueCSVDeduper) (chan JiraCustomFieldValue, chan bool, error) {
	f, err := os.Create(fn)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening CSV file %s. %v", fn, err)
	}
	var fc io.WriteCloser = f
	if filepath.Ext(fn) == ".gz" {
		w, _ := gzip.NewWriterLevel(f, gzip.BestCompression)
		fc = w
	}
	ch, done, err := NewJiraCustomFieldValueCSVWriter(fc, dedupers...)
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

type JiraCustomFieldValueDBAction func(ctx context.Context, db DB, record JiraCustomFieldValue) error

// NewJiraCustomFieldValueDBWriterSize creates a DB writer that will write each issue into the DB
func NewJiraCustomFieldValueDBWriterSize(ctx context.Context, db DB, errors chan<- error, size int, actions ...JiraCustomFieldValueDBAction) (chan JiraCustomFieldValue, chan bool, error) {
	ch := make(chan JiraCustomFieldValue, size)
	done := make(chan bool)
	var action JiraCustomFieldValueDBAction
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

// NewJiraCustomFieldValueDBWriter creates a DB writer that will write each issue into the DB
func NewJiraCustomFieldValueDBWriter(ctx context.Context, db DB, errors chan<- error, actions ...JiraCustomFieldValueDBAction) (chan JiraCustomFieldValue, chan bool, error) {
	return NewJiraCustomFieldValueDBWriterSize(ctx, db, errors, 100, actions...)
}

// JiraCustomFieldValueColumnID is the ID SQL column name for the JiraCustomFieldValue table
const JiraCustomFieldValueColumnID = "id"

// JiraCustomFieldValueEscapedColumnID is the escaped ID SQL column name for the JiraCustomFieldValue table
const JiraCustomFieldValueEscapedColumnID = "`id`"

// JiraCustomFieldValueColumnChecksum is the Checksum SQL column name for the JiraCustomFieldValue table
const JiraCustomFieldValueColumnChecksum = "checksum"

// JiraCustomFieldValueEscapedColumnChecksum is the escaped Checksum SQL column name for the JiraCustomFieldValue table
const JiraCustomFieldValueEscapedColumnChecksum = "`checksum`"

// JiraCustomFieldValueColumnSchemaType is the SchemaType SQL column name for the JiraCustomFieldValue table
const JiraCustomFieldValueColumnSchemaType = "schema_type"

// JiraCustomFieldValueEscapedColumnSchemaType is the escaped SchemaType SQL column name for the JiraCustomFieldValue table
const JiraCustomFieldValueEscapedColumnSchemaType = "`schema_type`"

// JiraCustomFieldValueColumnValue is the Value SQL column name for the JiraCustomFieldValue table
const JiraCustomFieldValueColumnValue = "value"

// JiraCustomFieldValueEscapedColumnValue is the escaped Value SQL column name for the JiraCustomFieldValue table
const JiraCustomFieldValueEscapedColumnValue = "`value`"

// JiraCustomFieldValueColumnCustomFieldID is the CustomFieldID SQL column name for the JiraCustomFieldValue table
const JiraCustomFieldValueColumnCustomFieldID = "custom_field_id"

// JiraCustomFieldValueEscapedColumnCustomFieldID is the escaped CustomFieldID SQL column name for the JiraCustomFieldValue table
const JiraCustomFieldValueEscapedColumnCustomFieldID = "`custom_field_id`"

// JiraCustomFieldValueColumnCustomerID is the CustomerID SQL column name for the JiraCustomFieldValue table
const JiraCustomFieldValueColumnCustomerID = "customer_id"

// JiraCustomFieldValueEscapedColumnCustomerID is the escaped CustomerID SQL column name for the JiraCustomFieldValue table
const JiraCustomFieldValueEscapedColumnCustomerID = "`customer_id`"

// GetID will return the JiraCustomFieldValue ID value
func (t *JiraCustomFieldValue) GetID() string {
	return t.ID
}

// SetID will set the JiraCustomFieldValue ID value
func (t *JiraCustomFieldValue) SetID(v string) {
	t.ID = v
}

// FindJiraCustomFieldValueByID will find a JiraCustomFieldValue by ID
func FindJiraCustomFieldValueByID(ctx context.Context, db DB, value string) (*JiraCustomFieldValue, error) {
	q := "SELECT * FROM `jira_custom_field_value` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _SchemaType sql.NullString
	var _Value sql.NullString
	var _CustomFieldID sql.NullString
	var _CustomerID sql.NullString
	err := db.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_SchemaType,
		&_Value,
		&_CustomFieldID,
		&_CustomerID,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &JiraCustomFieldValue{}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _Checksum.Valid {
		t.SetChecksum(_Checksum.String)
	}
	if _SchemaType.Valid {
		t.SetSchemaType(_SchemaType.String)
	}
	if _Value.Valid {
		t.SetValue(_Value.String)
	}
	if _CustomFieldID.Valid {
		t.SetCustomFieldID(_CustomFieldID.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	return t, nil
}

// FindJiraCustomFieldValueByIDTx will find a JiraCustomFieldValue by ID using the provided transaction
func FindJiraCustomFieldValueByIDTx(ctx context.Context, tx Tx, value string) (*JiraCustomFieldValue, error) {
	q := "SELECT * FROM `jira_custom_field_value` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _SchemaType sql.NullString
	var _Value sql.NullString
	var _CustomFieldID sql.NullString
	var _CustomerID sql.NullString
	err := tx.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_SchemaType,
		&_Value,
		&_CustomFieldID,
		&_CustomerID,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &JiraCustomFieldValue{}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _Checksum.Valid {
		t.SetChecksum(_Checksum.String)
	}
	if _SchemaType.Valid {
		t.SetSchemaType(_SchemaType.String)
	}
	if _Value.Valid {
		t.SetValue(_Value.String)
	}
	if _CustomFieldID.Valid {
		t.SetCustomFieldID(_CustomFieldID.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	return t, nil
}

// GetChecksum will return the JiraCustomFieldValue Checksum value
func (t *JiraCustomFieldValue) GetChecksum() string {
	if t.Checksum == nil {
		return ""
	}
	return *t.Checksum
}

// SetChecksum will set the JiraCustomFieldValue Checksum value
func (t *JiraCustomFieldValue) SetChecksum(v string) {
	t.Checksum = &v
}

// GetSchemaType will return the JiraCustomFieldValue SchemaType value
func (t *JiraCustomFieldValue) GetSchemaType() string {
	return t.SchemaType
}

// SetSchemaType will set the JiraCustomFieldValue SchemaType value
func (t *JiraCustomFieldValue) SetSchemaType(v string) {
	t.SchemaType = v
}

// GetValue will return the JiraCustomFieldValue Value value
func (t *JiraCustomFieldValue) GetValue() string {
	return t.Value
}

// SetValue will set the JiraCustomFieldValue Value value
func (t *JiraCustomFieldValue) SetValue(v string) {
	t.Value = v
}

// GetCustomFieldID will return the JiraCustomFieldValue CustomFieldID value
func (t *JiraCustomFieldValue) GetCustomFieldID() string {
	return t.CustomFieldID
}

// SetCustomFieldID will set the JiraCustomFieldValue CustomFieldID value
func (t *JiraCustomFieldValue) SetCustomFieldID(v string) {
	t.CustomFieldID = v
}

// FindJiraCustomFieldValuesByCustomFieldID will find all JiraCustomFieldValues by the CustomFieldID value
func FindJiraCustomFieldValuesByCustomFieldID(ctx context.Context, db DB, value string) ([]*JiraCustomFieldValue, error) {
	q := "SELECT * FROM `jira_custom_field_value` WHERE `custom_field_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraCustomFieldValue, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _SchemaType sql.NullString
		var _Value sql.NullString
		var _CustomFieldID sql.NullString
		var _CustomerID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_SchemaType,
			&_Value,
			&_CustomFieldID,
			&_CustomerID,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraCustomFieldValue{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _SchemaType.Valid {
			t.SetSchemaType(_SchemaType.String)
		}
		if _Value.Valid {
			t.SetValue(_Value.String)
		}
		if _CustomFieldID.Valid {
			t.SetCustomFieldID(_CustomFieldID.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindJiraCustomFieldValuesByCustomFieldIDTx will find all JiraCustomFieldValues by the CustomFieldID value using the provided transaction
func FindJiraCustomFieldValuesByCustomFieldIDTx(ctx context.Context, tx Tx, value string) ([]*JiraCustomFieldValue, error) {
	q := "SELECT * FROM `jira_custom_field_value` WHERE `custom_field_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraCustomFieldValue, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _SchemaType sql.NullString
		var _Value sql.NullString
		var _CustomFieldID sql.NullString
		var _CustomerID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_SchemaType,
			&_Value,
			&_CustomFieldID,
			&_CustomerID,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraCustomFieldValue{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _SchemaType.Valid {
			t.SetSchemaType(_SchemaType.String)
		}
		if _Value.Valid {
			t.SetValue(_Value.String)
		}
		if _CustomFieldID.Valid {
			t.SetCustomFieldID(_CustomFieldID.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetCustomerID will return the JiraCustomFieldValue CustomerID value
func (t *JiraCustomFieldValue) GetCustomerID() string {
	return t.CustomerID
}

// SetCustomerID will set the JiraCustomFieldValue CustomerID value
func (t *JiraCustomFieldValue) SetCustomerID(v string) {
	t.CustomerID = v
}

// FindJiraCustomFieldValuesByCustomerID will find all JiraCustomFieldValues by the CustomerID value
func FindJiraCustomFieldValuesByCustomerID(ctx context.Context, db DB, value string) ([]*JiraCustomFieldValue, error) {
	q := "SELECT * FROM `jira_custom_field_value` WHERE `customer_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraCustomFieldValue, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _SchemaType sql.NullString
		var _Value sql.NullString
		var _CustomFieldID sql.NullString
		var _CustomerID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_SchemaType,
			&_Value,
			&_CustomFieldID,
			&_CustomerID,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraCustomFieldValue{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _SchemaType.Valid {
			t.SetSchemaType(_SchemaType.String)
		}
		if _Value.Valid {
			t.SetValue(_Value.String)
		}
		if _CustomFieldID.Valid {
			t.SetCustomFieldID(_CustomFieldID.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindJiraCustomFieldValuesByCustomerIDTx will find all JiraCustomFieldValues by the CustomerID value using the provided transaction
func FindJiraCustomFieldValuesByCustomerIDTx(ctx context.Context, tx Tx, value string) ([]*JiraCustomFieldValue, error) {
	q := "SELECT * FROM `jira_custom_field_value` WHERE `customer_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraCustomFieldValue, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _SchemaType sql.NullString
		var _Value sql.NullString
		var _CustomFieldID sql.NullString
		var _CustomerID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_SchemaType,
			&_Value,
			&_CustomFieldID,
			&_CustomerID,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraCustomFieldValue{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _SchemaType.Valid {
			t.SetSchemaType(_SchemaType.String)
		}
		if _Value.Valid {
			t.SetValue(_Value.String)
		}
		if _CustomFieldID.Valid {
			t.SetCustomFieldID(_CustomFieldID.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

func (t *JiraCustomFieldValue) toTimestamp(value time.Time) *timestamp.Timestamp {
	ts, _ := ptypes.TimestampProto(value)
	return ts
}

// DBCreateJiraCustomFieldValueTable will create the JiraCustomFieldValue table
func DBCreateJiraCustomFieldValueTable(ctx context.Context, db DB) error {
	q := "CREATE TABLE `jira_custom_field_value` (`id` VARCHAR(64) NOT NULL PRIMARY KEY,`checksum` CHAR(64),`schema_type` VARCHAR(100) NOT NULL,`value` LONGTEXT NOT NULL,`custom_field_id` VARCHAR(64) NOT NULL,`customer_id` VARCHAR(64) NOT NULL,INDEX jira_custom_field_value_custom_field_id_index (`custom_field_id`),INDEX jira_custom_field_value_customer_id_index (`customer_id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBCreateJiraCustomFieldValueTableTx will create the JiraCustomFieldValue table using the provided transction
func DBCreateJiraCustomFieldValueTableTx(ctx context.Context, tx Tx) error {
	q := "CREATE TABLE `jira_custom_field_value` (`id` VARCHAR(64) NOT NULL PRIMARY KEY,`checksum` CHAR(64),`schema_type` VARCHAR(100) NOT NULL,`value` LONGTEXT NOT NULL,`custom_field_id` VARCHAR(64) NOT NULL,`customer_id` VARCHAR(64) NOT NULL,INDEX jira_custom_field_value_custom_field_id_index (`custom_field_id`),INDEX jira_custom_field_value_customer_id_index (`customer_id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// DBDropJiraCustomFieldValueTable will drop the JiraCustomFieldValue table
func DBDropJiraCustomFieldValueTable(ctx context.Context, db DB) error {
	q := "DROP TABLE IF EXISTS `jira_custom_field_value`"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBDropJiraCustomFieldValueTableTx will drop the JiraCustomFieldValue table using the provided transaction
func DBDropJiraCustomFieldValueTableTx(ctx context.Context, tx Tx) error {
	q := "DROP TABLE IF EXISTS `jira_custom_field_value`"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// CalculateChecksum will calculate a checksum of the SHA1 of all field values
func (t *JiraCustomFieldValue) CalculateChecksum() string {
	return orm.HashStrings(
		orm.ToString(t.ID),
		orm.ToString(t.SchemaType),
		orm.ToString(t.Value),
		orm.ToString(t.CustomFieldID),
		orm.ToString(t.CustomerID),
	)
}

// DBCreate will create a new JiraCustomFieldValue record in the database
func (t *JiraCustomFieldValue) DBCreate(ctx context.Context, db DB) (sql.Result, error) {
	q := "INSERT INTO `jira_custom_field_value` (`jira_custom_field_value`.`id`,`jira_custom_field_value`.`checksum`,`jira_custom_field_value`.`schema_type`,`jira_custom_field_value`.`value`,`jira_custom_field_value`.`custom_field_id`,`jira_custom_field_value`.`customer_id`) VALUES (?,?,?,?,?,?)"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.SchemaType),
		orm.ToSQLString(t.Value),
		orm.ToSQLString(t.CustomFieldID),
		orm.ToSQLString(t.CustomerID),
	)
}

// DBCreateTx will create a new JiraCustomFieldValue record in the database using the provided transaction
func (t *JiraCustomFieldValue) DBCreateTx(ctx context.Context, tx Tx) (sql.Result, error) {
	q := "INSERT INTO `jira_custom_field_value` (`jira_custom_field_value`.`id`,`jira_custom_field_value`.`checksum`,`jira_custom_field_value`.`schema_type`,`jira_custom_field_value`.`value`,`jira_custom_field_value`.`custom_field_id`,`jira_custom_field_value`.`customer_id`) VALUES (?,?,?,?,?,?)"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.SchemaType),
		orm.ToSQLString(t.Value),
		orm.ToSQLString(t.CustomFieldID),
		orm.ToSQLString(t.CustomerID),
	)
}

// DBCreateIgnoreDuplicate will upsert the JiraCustomFieldValue record in the database
func (t *JiraCustomFieldValue) DBCreateIgnoreDuplicate(ctx context.Context, db DB) (sql.Result, error) {
	q := "INSERT INTO `jira_custom_field_value` (`jira_custom_field_value`.`id`,`jira_custom_field_value`.`checksum`,`jira_custom_field_value`.`schema_type`,`jira_custom_field_value`.`value`,`jira_custom_field_value`.`custom_field_id`,`jira_custom_field_value`.`customer_id`) VALUES (?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.SchemaType),
		orm.ToSQLString(t.Value),
		orm.ToSQLString(t.CustomFieldID),
		orm.ToSQLString(t.CustomerID),
	)
}

// DBCreateIgnoreDuplicateTx will upsert the JiraCustomFieldValue record in the database using the provided transaction
func (t *JiraCustomFieldValue) DBCreateIgnoreDuplicateTx(ctx context.Context, tx Tx) (sql.Result, error) {
	q := "INSERT INTO `jira_custom_field_value` (`jira_custom_field_value`.`id`,`jira_custom_field_value`.`checksum`,`jira_custom_field_value`.`schema_type`,`jira_custom_field_value`.`value`,`jira_custom_field_value`.`custom_field_id`,`jira_custom_field_value`.`customer_id`) VALUES (?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.SchemaType),
		orm.ToSQLString(t.Value),
		orm.ToSQLString(t.CustomFieldID),
		orm.ToSQLString(t.CustomerID),
	)
}

// DeleteAllJiraCustomFieldValues deletes all JiraCustomFieldValue records in the database with optional filters
func DeleteAllJiraCustomFieldValues(ctx context.Context, db DB, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(JiraCustomFieldValueTableName),
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

// DeleteAllJiraCustomFieldValuesTx deletes all JiraCustomFieldValue records in the database with optional filters using the provided transaction
func DeleteAllJiraCustomFieldValuesTx(ctx context.Context, tx Tx, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(JiraCustomFieldValueTableName),
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

// DBDelete will delete this JiraCustomFieldValue record in the database
func (t *JiraCustomFieldValue) DBDelete(ctx context.Context, db DB) (bool, error) {
	q := "DELETE FROM `jira_custom_field_value` WHERE `id` = ?"
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

// DBDeleteTx will delete this JiraCustomFieldValue record in the database using the provided transaction
func (t *JiraCustomFieldValue) DBDeleteTx(ctx context.Context, tx Tx) (bool, error) {
	q := "DELETE FROM `jira_custom_field_value` WHERE `id` = ?"
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

// DBUpdate will update the JiraCustomFieldValue record in the database
func (t *JiraCustomFieldValue) DBUpdate(ctx context.Context, db DB) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `jira_custom_field_value` SET `checksum`=?,`schema_type`=?,`value`=?,`custom_field_id`=?,`customer_id`=? WHERE `id`=?"
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.SchemaType),
		orm.ToSQLString(t.Value),
		orm.ToSQLString(t.CustomFieldID),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.ID),
	)
}

// DBUpdateTx will update the JiraCustomFieldValue record in the database using the provided transaction
func (t *JiraCustomFieldValue) DBUpdateTx(ctx context.Context, tx Tx) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `jira_custom_field_value` SET `checksum`=?,`schema_type`=?,`value`=?,`custom_field_id`=?,`customer_id`=? WHERE `id`=?"
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.SchemaType),
		orm.ToSQLString(t.Value),
		orm.ToSQLString(t.CustomFieldID),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.ID),
	)
}

// DBUpsert will upsert the JiraCustomFieldValue record in the database
func (t *JiraCustomFieldValue) DBUpsert(ctx context.Context, db DB, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `jira_custom_field_value` (`jira_custom_field_value`.`id`,`jira_custom_field_value`.`checksum`,`jira_custom_field_value`.`schema_type`,`jira_custom_field_value`.`value`,`jira_custom_field_value`.`custom_field_id`,`jira_custom_field_value`.`customer_id`) VALUES (?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `jira_custom_field_value` (`jira_custom_field_value`.`id`,`jira_custom_field_value`.`checksum`,`jira_custom_field_value`.`schema_type`,`jira_custom_field_value`.`value`,`jira_custom_field_value`.`custom_field_id`,`jira_custom_field_value`.`customer_id`) VALUES (?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`schema_type`=VALUES(`schema_type`),`value`=VALUES(`value`),`custom_field_id`=VALUES(`custom_field_id`),`customer_id`=VALUES(`customer_id`)"
	}
	r, err := db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.SchemaType),
		orm.ToSQLString(t.Value),
		orm.ToSQLString(t.CustomFieldID),
		orm.ToSQLString(t.CustomerID),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBUpsertTx will upsert the JiraCustomFieldValue record in the database using the provided transaction
func (t *JiraCustomFieldValue) DBUpsertTx(ctx context.Context, tx Tx, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `jira_custom_field_value` (`jira_custom_field_value`.`id`,`jira_custom_field_value`.`checksum`,`jira_custom_field_value`.`schema_type`,`jira_custom_field_value`.`value`,`jira_custom_field_value`.`custom_field_id`,`jira_custom_field_value`.`customer_id`) VALUES (?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `jira_custom_field_value` (`jira_custom_field_value`.`id`,`jira_custom_field_value`.`checksum`,`jira_custom_field_value`.`schema_type`,`jira_custom_field_value`.`value`,`jira_custom_field_value`.`custom_field_id`,`jira_custom_field_value`.`customer_id`) VALUES (?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`schema_type`=VALUES(`schema_type`),`value`=VALUES(`value`),`custom_field_id`=VALUES(`custom_field_id`),`customer_id`=VALUES(`customer_id`)"
	}
	r, err := tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.SchemaType),
		orm.ToSQLString(t.Value),
		orm.ToSQLString(t.CustomFieldID),
		orm.ToSQLString(t.CustomerID),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBFindOne will find a JiraCustomFieldValue record in the database with the primary key
func (t *JiraCustomFieldValue) DBFindOne(ctx context.Context, db DB, value string) (bool, error) {
	q := "SELECT * FROM `jira_custom_field_value` WHERE `id` = ? LIMIT 1"
	row := db.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _SchemaType sql.NullString
	var _Value sql.NullString
	var _CustomFieldID sql.NullString
	var _CustomerID sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_SchemaType,
		&_Value,
		&_CustomFieldID,
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
	if _SchemaType.Valid {
		t.SetSchemaType(_SchemaType.String)
	}
	if _Value.Valid {
		t.SetValue(_Value.String)
	}
	if _CustomFieldID.Valid {
		t.SetCustomFieldID(_CustomFieldID.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	return true, nil
}

// DBFindOneTx will find a JiraCustomFieldValue record in the database with the primary key using the provided transaction
func (t *JiraCustomFieldValue) DBFindOneTx(ctx context.Context, tx Tx, value string) (bool, error) {
	q := "SELECT * FROM `jira_custom_field_value` WHERE `id` = ? LIMIT 1"
	row := tx.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _SchemaType sql.NullString
	var _Value sql.NullString
	var _CustomFieldID sql.NullString
	var _CustomerID sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_SchemaType,
		&_Value,
		&_CustomFieldID,
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
	if _SchemaType.Valid {
		t.SetSchemaType(_SchemaType.String)
	}
	if _Value.Valid {
		t.SetValue(_Value.String)
	}
	if _CustomFieldID.Valid {
		t.SetCustomFieldID(_CustomFieldID.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	return true, nil
}

// FindJiraCustomFieldValues will find a JiraCustomFieldValue record in the database with the provided parameters
func FindJiraCustomFieldValues(ctx context.Context, db DB, _params ...interface{}) ([]*JiraCustomFieldValue, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("schema_type"),
		orm.Column("value"),
		orm.Column("custom_field_id"),
		orm.Column("customer_id"),
		orm.Table(JiraCustomFieldValueTableName),
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
	results := make([]*JiraCustomFieldValue, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _SchemaType sql.NullString
		var _Value sql.NullString
		var _CustomFieldID sql.NullString
		var _CustomerID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_SchemaType,
			&_Value,
			&_CustomFieldID,
			&_CustomerID,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraCustomFieldValue{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _SchemaType.Valid {
			t.SetSchemaType(_SchemaType.String)
		}
		if _Value.Valid {
			t.SetValue(_Value.String)
		}
		if _CustomFieldID.Valid {
			t.SetCustomFieldID(_CustomFieldID.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindJiraCustomFieldValuesTx will find a JiraCustomFieldValue record in the database with the provided parameters using the provided transaction
func FindJiraCustomFieldValuesTx(ctx context.Context, tx Tx, _params ...interface{}) ([]*JiraCustomFieldValue, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("schema_type"),
		orm.Column("value"),
		orm.Column("custom_field_id"),
		orm.Column("customer_id"),
		orm.Table(JiraCustomFieldValueTableName),
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
	results := make([]*JiraCustomFieldValue, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _SchemaType sql.NullString
		var _Value sql.NullString
		var _CustomFieldID sql.NullString
		var _CustomerID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_SchemaType,
			&_Value,
			&_CustomFieldID,
			&_CustomerID,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraCustomFieldValue{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _SchemaType.Valid {
			t.SetSchemaType(_SchemaType.String)
		}
		if _Value.Valid {
			t.SetValue(_Value.String)
		}
		if _CustomFieldID.Valid {
			t.SetCustomFieldID(_CustomFieldID.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// DBFind will find a JiraCustomFieldValue record in the database with the provided parameters
func (t *JiraCustomFieldValue) DBFind(ctx context.Context, db DB, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("schema_type"),
		orm.Column("value"),
		orm.Column("custom_field_id"),
		orm.Column("customer_id"),
		orm.Table(JiraCustomFieldValueTableName),
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
	var _SchemaType sql.NullString
	var _Value sql.NullString
	var _CustomFieldID sql.NullString
	var _CustomerID sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_SchemaType,
		&_Value,
		&_CustomFieldID,
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
	if _SchemaType.Valid {
		t.SetSchemaType(_SchemaType.String)
	}
	if _Value.Valid {
		t.SetValue(_Value.String)
	}
	if _CustomFieldID.Valid {
		t.SetCustomFieldID(_CustomFieldID.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	return true, nil
}

// DBFindTx will find a JiraCustomFieldValue record in the database with the provided parameters using the provided transaction
func (t *JiraCustomFieldValue) DBFindTx(ctx context.Context, tx Tx, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("schema_type"),
		orm.Column("value"),
		orm.Column("custom_field_id"),
		orm.Column("customer_id"),
		orm.Table(JiraCustomFieldValueTableName),
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
	var _SchemaType sql.NullString
	var _Value sql.NullString
	var _CustomFieldID sql.NullString
	var _CustomerID sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_SchemaType,
		&_Value,
		&_CustomFieldID,
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
	if _SchemaType.Valid {
		t.SetSchemaType(_SchemaType.String)
	}
	if _Value.Valid {
		t.SetValue(_Value.String)
	}
	if _CustomFieldID.Valid {
		t.SetCustomFieldID(_CustomFieldID.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	return true, nil
}

// CountJiraCustomFieldValues will find the count of JiraCustomFieldValue records in the database
func CountJiraCustomFieldValues(ctx context.Context, db DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(JiraCustomFieldValueTableName),
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

// CountJiraCustomFieldValuesTx will find the count of JiraCustomFieldValue records in the database using the provided transaction
func CountJiraCustomFieldValuesTx(ctx context.Context, tx Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(JiraCustomFieldValueTableName),
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

// DBCount will find the count of JiraCustomFieldValue records in the database
func (t *JiraCustomFieldValue) DBCount(ctx context.Context, db DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(JiraCustomFieldValueTableName),
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

// DBCountTx will find the count of JiraCustomFieldValue records in the database using the provided transaction
func (t *JiraCustomFieldValue) DBCountTx(ctx context.Context, tx Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(JiraCustomFieldValueTableName),
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

// DBExists will return true if the JiraCustomFieldValue record exists in the database
func (t *JiraCustomFieldValue) DBExists(ctx context.Context, db DB) (bool, error) {
	q := "SELECT * FROM `jira_custom_field_value` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := db.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// DBExistsTx will return true if the JiraCustomFieldValue record exists in the database using the provided transaction
func (t *JiraCustomFieldValue) DBExistsTx(ctx context.Context, tx Tx) (bool, error) {
	q := "SELECT * FROM `jira_custom_field_value` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := tx.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// PrimaryKeyColumn returns the column name for the primary key
func (t *JiraCustomFieldValue) PrimaryKeyColumn() string {
	return JiraCustomFieldValueColumnID
}

// PrimaryKeyColumnType returns the primary key column Go type as a string
func (t *JiraCustomFieldValue) PrimaryKeyColumnType() string {
	return "string"
}

// PrimaryKey returns the primary key column value
func (t *JiraCustomFieldValue) PrimaryKey() interface{} {
	return t.ID
}
