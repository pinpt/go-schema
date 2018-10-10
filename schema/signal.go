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

var _ Model = (*Signal)(nil)
var _ CSVWriter = (*Signal)(nil)
var _ JSONWriter = (*Signal)(nil)

// SignalTableName is the name of the table in SQL
const SignalTableName = "signal"

var SignalColumns = []string{
	"id",
	"name",
	"value",
	"timeunit",
	"date",
	"metadata",
	"customer_id",
	"ref_type",
	"ref_id",
}

// Signal table
type Signal struct {
	CustomerID string  `json:"customer_id"`
	Date       string  `json:"date"`
	ID         string  `json:"id"`
	Metadata   *string `json:"metadata,omitempty"`
	Name       string  `json:"name"`
	RefID      *string `json:"ref_id,omitempty"`
	RefType    *string `json:"ref_type,omitempty"`
	Timeunit   int32   `json:"timeunit"`
	Value      float64 `json:"value"`
}

// TableName returns the SQL table name for Signal and satifies the Model interface
func (t *Signal) TableName() string {
	return SignalTableName
}

// ToCSV will serialize the Signal instance to a CSV compatible array of strings
func (t *Signal) ToCSV() []string {
	return []string{
		t.ID,
		t.Name,
		toCSVString(t.Value),
		toCSVString(t.Timeunit),
		t.Date,
		toCSVString(t.Metadata),
		t.CustomerID,
		toCSVString(t.RefType),
		toCSVString(t.RefID),
	}
}

// WriteCSV will serialize the Signal instance to the writer as CSV and satisfies the CSVWriter interface
func (t *Signal) WriteCSV(w *csv.Writer) error {
	return w.Write(t.ToCSV())
}

// WriteJSON will serialize the Signal instance to the writer as JSON and satisfies the JSONWriter interface
func (t *Signal) WriteJSON(w io.Writer, indent ...bool) error {
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

// NewSignalReader creates a JSON reader which can read in Signal objects serialized as JSON either as an array, single object or json new lines
// and writes each Signal to the channel provided
func NewSignalReader(r io.Reader, ch chan<- Signal) error {
	return orm.Deserialize(r, func(buf json.RawMessage) error {
		dec := json.NewDecoder(bytes.NewBuffer(buf))
		e := Signal{}
		if err := dec.Decode(&e); err != nil {
			return err
		}
		ch <- e
		return nil
	})
}

// NewCSVSignalReaderDir will read the reader as CSV and emit each record to the channel provided
func NewCSVSignalReader(r io.Reader, ch chan<- Signal) error {
	cr := csv.NewReader(r)
	for {
		record, err := cr.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		ch <- Signal{
			ID:         record[0],
			Name:       record[1],
			Value:      fromCSVFloat64(record[2]),
			Timeunit:   fromCSVInt32(record[3]),
			Date:       record[4],
			Metadata:   fromStringPointer(record[5]),
			CustomerID: record[6],
			RefType:    fromStringPointer(record[7]),
			RefID:      fromStringPointer(record[8]),
		}
	}
	return nil
}

// NewCSVSignalReaderFile will read the file as a CSV and emit each record to the channel provided
func NewCSVSignalReaderFile(fp string, ch chan<- Signal) error {
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
	return NewCSVSignalReader(fc, ch)
}

// NewCSVSignalReaderDir will read the signal.csv.gz file as a CSV and emit each record to the channel provided
func NewCSVSignalReaderDir(dir string, ch chan<- Signal) error {
	return NewCSVSignalReaderFile(filepath.Join(dir, "signal.csv.gz"), ch)
}

// SignalCSVDeduper is a function callback which takes the existing value (a) and the new value (b)
// and the return value should be the one to use (or a new one, if applicable). return nil
// to skip processing of this record
type SignalCSVDeduper func(a Signal, b Signal) *Signal

// SignalCSVDedupeDisabled is set on whether the CSV writer should de-dupe values by key
var SignalCSVDedupeDisabled bool

// NewSignalCSVWriterSize creates a batch writer that will write each Signal into a CSV file
func NewSignalCSVWriterSize(w io.Writer, size int, dedupers ...SignalCSVDeduper) (chan Signal, chan bool, error) {
	cw := csv.NewWriter(w)
	ch := make(chan Signal, size)
	done := make(chan bool)
	go func() {
		defer func() { done <- true }()
		dodedupe := !SignalCSVDedupeDisabled
		var kv map[string]*Signal
		var deduper SignalCSVDeduper
		if dedupers != nil && len(dedupers) > 0 {
			deduper = dedupers[0]
			dodedupe = true
		}
		if dodedupe {
			kv = make(map[string]*Signal)
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

// SignalCSVDefaultSize is the default channel buffer size if not provided
var SignalCSVDefaultSize = 100

// NewSignalCSVWriter creates a batch writer that will write each Signal into a CSV file
func NewSignalCSVWriter(w io.Writer, dedupers ...SignalCSVDeduper) (chan Signal, chan bool, error) {
	return NewSignalCSVWriterSize(w, SignalCSVDefaultSize, dedupers...)
}

// NewSignalCSVWriterDir creates a batch writer that will write each Signal into a CSV file named signal.csv.gz in dir
func NewSignalCSVWriterDir(dir string, dedupers ...SignalCSVDeduper) (chan Signal, chan bool, error) {
	return NewSignalCSVWriterFile(filepath.Join(dir, "signal.csv.gz"), dedupers...)
}

// NewSignalCSVWriterFile creates a batch writer that will write each Signal into a CSV file
func NewSignalCSVWriterFile(fn string, dedupers ...SignalCSVDeduper) (chan Signal, chan bool, error) {
	f, err := os.Create(fn)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening CSV file %s. %v", fn, err)
	}
	var fc io.WriteCloser = f
	if filepath.Ext(fn) == ".gz" {
		w, _ := gzip.NewWriterLevel(f, gzip.BestCompression)
		fc = w
	}
	ch, done, err := NewSignalCSVWriter(fc, dedupers...)
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

type SignalDBAction func(ctx context.Context, db DB, record Signal) error

// NewSignalDBWriterSize creates a DB writer that will write each issue into the DB
func NewSignalDBWriterSize(ctx context.Context, db DB, errors chan<- error, size int, actions ...SignalDBAction) (chan Signal, chan bool, error) {
	ch := make(chan Signal, size)
	done := make(chan bool)
	var action SignalDBAction
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

// NewSignalDBWriter creates a DB writer that will write each issue into the DB
func NewSignalDBWriter(ctx context.Context, db DB, errors chan<- error, actions ...SignalDBAction) (chan Signal, chan bool, error) {
	return NewSignalDBWriterSize(ctx, db, errors, 100, actions...)
}

// SignalColumnID is the ID SQL column name for the Signal table
const SignalColumnID = "id"

// SignalEscapedColumnID is the escaped ID SQL column name for the Signal table
const SignalEscapedColumnID = "`id`"

// SignalColumnName is the Name SQL column name for the Signal table
const SignalColumnName = "name"

// SignalEscapedColumnName is the escaped Name SQL column name for the Signal table
const SignalEscapedColumnName = "`name`"

// SignalColumnValue is the Value SQL column name for the Signal table
const SignalColumnValue = "value"

// SignalEscapedColumnValue is the escaped Value SQL column name for the Signal table
const SignalEscapedColumnValue = "`value`"

// SignalColumnTimeunit is the Timeunit SQL column name for the Signal table
const SignalColumnTimeunit = "timeunit"

// SignalEscapedColumnTimeunit is the escaped Timeunit SQL column name for the Signal table
const SignalEscapedColumnTimeunit = "`timeunit`"

// SignalColumnDate is the Date SQL column name for the Signal table
const SignalColumnDate = "date"

// SignalEscapedColumnDate is the escaped Date SQL column name for the Signal table
const SignalEscapedColumnDate = "`date`"

// SignalColumnMetadata is the Metadata SQL column name for the Signal table
const SignalColumnMetadata = "metadata"

// SignalEscapedColumnMetadata is the escaped Metadata SQL column name for the Signal table
const SignalEscapedColumnMetadata = "`metadata`"

// SignalColumnCustomerID is the CustomerID SQL column name for the Signal table
const SignalColumnCustomerID = "customer_id"

// SignalEscapedColumnCustomerID is the escaped CustomerID SQL column name for the Signal table
const SignalEscapedColumnCustomerID = "`customer_id`"

// SignalColumnRefType is the RefType SQL column name for the Signal table
const SignalColumnRefType = "ref_type"

// SignalEscapedColumnRefType is the escaped RefType SQL column name for the Signal table
const SignalEscapedColumnRefType = "`ref_type`"

// SignalColumnRefID is the RefID SQL column name for the Signal table
const SignalColumnRefID = "ref_id"

// SignalEscapedColumnRefID is the escaped RefID SQL column name for the Signal table
const SignalEscapedColumnRefID = "`ref_id`"

// GetID will return the Signal ID value
func (t *Signal) GetID() string {
	return t.ID
}

// SetID will set the Signal ID value
func (t *Signal) SetID(v string) {
	t.ID = v
}

// FindSignalByID will find a Signal by ID
func FindSignalByID(ctx context.Context, db DB, value string) (*Signal, error) {
	q := "SELECT * FROM `signal` WHERE `id` = ?"
	var _ID sql.NullString
	var _Name sql.NullString
	var _Value sql.NullFloat64
	var _Timeunit sql.NullInt64
	var _Date sql.NullString
	var _Metadata sql.NullString
	var _CustomerID sql.NullString
	var _RefType sql.NullString
	var _RefID sql.NullString
	err := db.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Name,
		&_Value,
		&_Timeunit,
		&_Date,
		&_Metadata,
		&_CustomerID,
		&_RefType,
		&_RefID,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &Signal{}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _Value.Valid {
		t.SetValue(_Value.Float64)
	}
	if _Timeunit.Valid {
		t.SetTimeunit(int32(_Timeunit.Int64))
	}
	if _Date.Valid {
		t.SetDate(_Date.String)
	}
	if _Metadata.Valid {
		t.SetMetadata(_Metadata.String)
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
	return t, nil
}

// FindSignalByIDTx will find a Signal by ID using the provided transaction
func FindSignalByIDTx(ctx context.Context, tx Tx, value string) (*Signal, error) {
	q := "SELECT * FROM `signal` WHERE `id` = ?"
	var _ID sql.NullString
	var _Name sql.NullString
	var _Value sql.NullFloat64
	var _Timeunit sql.NullInt64
	var _Date sql.NullString
	var _Metadata sql.NullString
	var _CustomerID sql.NullString
	var _RefType sql.NullString
	var _RefID sql.NullString
	err := tx.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Name,
		&_Value,
		&_Timeunit,
		&_Date,
		&_Metadata,
		&_CustomerID,
		&_RefType,
		&_RefID,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &Signal{}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _Value.Valid {
		t.SetValue(_Value.Float64)
	}
	if _Timeunit.Valid {
		t.SetTimeunit(int32(_Timeunit.Int64))
	}
	if _Date.Valid {
		t.SetDate(_Date.String)
	}
	if _Metadata.Valid {
		t.SetMetadata(_Metadata.String)
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
	return t, nil
}

// GetName will return the Signal Name value
func (t *Signal) GetName() string {
	return t.Name
}

// SetName will set the Signal Name value
func (t *Signal) SetName(v string) {
	t.Name = v
}

// FindSignalsByName will find all Signals by the Name value
func FindSignalsByName(ctx context.Context, db DB, value string) ([]*Signal, error) {
	q := "SELECT * FROM `signal` WHERE `name` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*Signal, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Name sql.NullString
		var _Value sql.NullFloat64
		var _Timeunit sql.NullInt64
		var _Date sql.NullString
		var _Metadata sql.NullString
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Name,
			&_Value,
			&_Timeunit,
			&_Date,
			&_Metadata,
			&_CustomerID,
			&_RefType,
			&_RefID,
		)
		if err != nil {
			return nil, err
		}
		t := &Signal{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _Value.Valid {
			t.SetValue(_Value.Float64)
		}
		if _Timeunit.Valid {
			t.SetTimeunit(int32(_Timeunit.Int64))
		}
		if _Date.Valid {
			t.SetDate(_Date.String)
		}
		if _Metadata.Valid {
			t.SetMetadata(_Metadata.String)
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
		results = append(results, t)
	}
	return results, nil
}

// FindSignalsByNameTx will find all Signals by the Name value using the provided transaction
func FindSignalsByNameTx(ctx context.Context, tx Tx, value string) ([]*Signal, error) {
	q := "SELECT * FROM `signal` WHERE `name` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*Signal, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Name sql.NullString
		var _Value sql.NullFloat64
		var _Timeunit sql.NullInt64
		var _Date sql.NullString
		var _Metadata sql.NullString
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Name,
			&_Value,
			&_Timeunit,
			&_Date,
			&_Metadata,
			&_CustomerID,
			&_RefType,
			&_RefID,
		)
		if err != nil {
			return nil, err
		}
		t := &Signal{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _Value.Valid {
			t.SetValue(_Value.Float64)
		}
		if _Timeunit.Valid {
			t.SetTimeunit(int32(_Timeunit.Int64))
		}
		if _Date.Valid {
			t.SetDate(_Date.String)
		}
		if _Metadata.Valid {
			t.SetMetadata(_Metadata.String)
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
		results = append(results, t)
	}
	return results, nil
}

// GetValue will return the Signal Value value
func (t *Signal) GetValue() float64 {
	return t.Value
}

// SetValue will set the Signal Value value
func (t *Signal) SetValue(v float64) {
	t.Value = v
}

// GetTimeunit will return the Signal Timeunit value
func (t *Signal) GetTimeunit() int32 {
	return t.Timeunit
}

// SetTimeunit will set the Signal Timeunit value
func (t *Signal) SetTimeunit(v int32) {
	t.Timeunit = v
}

// GetDate will return the Signal Date value
func (t *Signal) GetDate() string {
	return t.Date
}

// SetDate will set the Signal Date value
func (t *Signal) SetDate(v string) {
	t.Date = v
}

// GetMetadata will return the Signal Metadata value
func (t *Signal) GetMetadata() string {
	if t.Metadata == nil {
		return ""
	}
	return *t.Metadata
}

// SetMetadata will set the Signal Metadata value
func (t *Signal) SetMetadata(v string) {
	t.Metadata = &v
}

// GetCustomerID will return the Signal CustomerID value
func (t *Signal) GetCustomerID() string {
	return t.CustomerID
}

// SetCustomerID will set the Signal CustomerID value
func (t *Signal) SetCustomerID(v string) {
	t.CustomerID = v
}

// FindSignalsByCustomerID will find all Signals by the CustomerID value
func FindSignalsByCustomerID(ctx context.Context, db DB, value string) ([]*Signal, error) {
	q := "SELECT * FROM `signal` WHERE `customer_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*Signal, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Name sql.NullString
		var _Value sql.NullFloat64
		var _Timeunit sql.NullInt64
		var _Date sql.NullString
		var _Metadata sql.NullString
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Name,
			&_Value,
			&_Timeunit,
			&_Date,
			&_Metadata,
			&_CustomerID,
			&_RefType,
			&_RefID,
		)
		if err != nil {
			return nil, err
		}
		t := &Signal{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _Value.Valid {
			t.SetValue(_Value.Float64)
		}
		if _Timeunit.Valid {
			t.SetTimeunit(int32(_Timeunit.Int64))
		}
		if _Date.Valid {
			t.SetDate(_Date.String)
		}
		if _Metadata.Valid {
			t.SetMetadata(_Metadata.String)
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
		results = append(results, t)
	}
	return results, nil
}

// FindSignalsByCustomerIDTx will find all Signals by the CustomerID value using the provided transaction
func FindSignalsByCustomerIDTx(ctx context.Context, tx Tx, value string) ([]*Signal, error) {
	q := "SELECT * FROM `signal` WHERE `customer_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*Signal, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Name sql.NullString
		var _Value sql.NullFloat64
		var _Timeunit sql.NullInt64
		var _Date sql.NullString
		var _Metadata sql.NullString
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Name,
			&_Value,
			&_Timeunit,
			&_Date,
			&_Metadata,
			&_CustomerID,
			&_RefType,
			&_RefID,
		)
		if err != nil {
			return nil, err
		}
		t := &Signal{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _Value.Valid {
			t.SetValue(_Value.Float64)
		}
		if _Timeunit.Valid {
			t.SetTimeunit(int32(_Timeunit.Int64))
		}
		if _Date.Valid {
			t.SetDate(_Date.String)
		}
		if _Metadata.Valid {
			t.SetMetadata(_Metadata.String)
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
		results = append(results, t)
	}
	return results, nil
}

// GetRefType will return the Signal RefType value
func (t *Signal) GetRefType() string {
	if t.RefType == nil {
		return ""
	}
	return *t.RefType
}

// SetRefType will set the Signal RefType value
func (t *Signal) SetRefType(v string) {
	t.RefType = &v
}

// FindSignalsByRefType will find all Signals by the RefType value
func FindSignalsByRefType(ctx context.Context, db DB, value string) ([]*Signal, error) {
	q := "SELECT * FROM `signal` WHERE `ref_type` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*Signal, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Name sql.NullString
		var _Value sql.NullFloat64
		var _Timeunit sql.NullInt64
		var _Date sql.NullString
		var _Metadata sql.NullString
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Name,
			&_Value,
			&_Timeunit,
			&_Date,
			&_Metadata,
			&_CustomerID,
			&_RefType,
			&_RefID,
		)
		if err != nil {
			return nil, err
		}
		t := &Signal{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _Value.Valid {
			t.SetValue(_Value.Float64)
		}
		if _Timeunit.Valid {
			t.SetTimeunit(int32(_Timeunit.Int64))
		}
		if _Date.Valid {
			t.SetDate(_Date.String)
		}
		if _Metadata.Valid {
			t.SetMetadata(_Metadata.String)
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
		results = append(results, t)
	}
	return results, nil
}

// FindSignalsByRefTypeTx will find all Signals by the RefType value using the provided transaction
func FindSignalsByRefTypeTx(ctx context.Context, tx Tx, value string) ([]*Signal, error) {
	q := "SELECT * FROM `signal` WHERE `ref_type` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*Signal, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Name sql.NullString
		var _Value sql.NullFloat64
		var _Timeunit sql.NullInt64
		var _Date sql.NullString
		var _Metadata sql.NullString
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Name,
			&_Value,
			&_Timeunit,
			&_Date,
			&_Metadata,
			&_CustomerID,
			&_RefType,
			&_RefID,
		)
		if err != nil {
			return nil, err
		}
		t := &Signal{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _Value.Valid {
			t.SetValue(_Value.Float64)
		}
		if _Timeunit.Valid {
			t.SetTimeunit(int32(_Timeunit.Int64))
		}
		if _Date.Valid {
			t.SetDate(_Date.String)
		}
		if _Metadata.Valid {
			t.SetMetadata(_Metadata.String)
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
		results = append(results, t)
	}
	return results, nil
}

// GetRefID will return the Signal RefID value
func (t *Signal) GetRefID() string {
	if t.RefID == nil {
		return ""
	}
	return *t.RefID
}

// SetRefID will set the Signal RefID value
func (t *Signal) SetRefID(v string) {
	t.RefID = &v
}

// FindSignalsByRefID will find all Signals by the RefID value
func FindSignalsByRefID(ctx context.Context, db DB, value string) ([]*Signal, error) {
	q := "SELECT * FROM `signal` WHERE `ref_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*Signal, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Name sql.NullString
		var _Value sql.NullFloat64
		var _Timeunit sql.NullInt64
		var _Date sql.NullString
		var _Metadata sql.NullString
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Name,
			&_Value,
			&_Timeunit,
			&_Date,
			&_Metadata,
			&_CustomerID,
			&_RefType,
			&_RefID,
		)
		if err != nil {
			return nil, err
		}
		t := &Signal{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _Value.Valid {
			t.SetValue(_Value.Float64)
		}
		if _Timeunit.Valid {
			t.SetTimeunit(int32(_Timeunit.Int64))
		}
		if _Date.Valid {
			t.SetDate(_Date.String)
		}
		if _Metadata.Valid {
			t.SetMetadata(_Metadata.String)
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
		results = append(results, t)
	}
	return results, nil
}

// FindSignalsByRefIDTx will find all Signals by the RefID value using the provided transaction
func FindSignalsByRefIDTx(ctx context.Context, tx Tx, value string) ([]*Signal, error) {
	q := "SELECT * FROM `signal` WHERE `ref_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*Signal, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Name sql.NullString
		var _Value sql.NullFloat64
		var _Timeunit sql.NullInt64
		var _Date sql.NullString
		var _Metadata sql.NullString
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Name,
			&_Value,
			&_Timeunit,
			&_Date,
			&_Metadata,
			&_CustomerID,
			&_RefType,
			&_RefID,
		)
		if err != nil {
			return nil, err
		}
		t := &Signal{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _Value.Valid {
			t.SetValue(_Value.Float64)
		}
		if _Timeunit.Valid {
			t.SetTimeunit(int32(_Timeunit.Int64))
		}
		if _Date.Valid {
			t.SetDate(_Date.String)
		}
		if _Metadata.Valid {
			t.SetMetadata(_Metadata.String)
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
		results = append(results, t)
	}
	return results, nil
}

func (t *Signal) toTimestamp(value time.Time) *timestamp.Timestamp {
	ts, _ := ptypes.TimestampProto(value)
	return ts
}

// DBCreateSignalTable will create the Signal table
func DBCreateSignalTable(ctx context.Context, db DB) error {
	q := "CREATE TABLE `signal` (`id` VARCHAR(64) NOT NULL PRIMARY KEY,`name`CHAR(60) NOT NULL,`value` REAL NOT NULL DEFAULT 0,`timeunit` INT NOT NULL,`date`DATE NOT NULL,`metadata` JSON,`customer_id` VARCHAR(64) NOT NULL,`ref_type` VARCHAR(20),`ref_id` VARCHAR(64),INDEX signal_name_index (`name`),INDEX signal_customer_id_index (`customer_id`),INDEX signal_ref_type_index (`ref_type`),INDEX signal_ref_id_index (`ref_id`),INDEX signal_name_timeunit_customer_id_index (`name`,`timeunit`,`customer_id`),INDEX signal_name_timeunit_customer_id_date_index (`name`,`timeunit`,`customer_id`,`date`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBCreateSignalTableTx will create the Signal table using the provided transction
func DBCreateSignalTableTx(ctx context.Context, tx Tx) error {
	q := "CREATE TABLE `signal` (`id` VARCHAR(64) NOT NULL PRIMARY KEY,`name`CHAR(60) NOT NULL,`value` REAL NOT NULL DEFAULT 0,`timeunit` INT NOT NULL,`date`DATE NOT NULL,`metadata` JSON,`customer_id` VARCHAR(64) NOT NULL,`ref_type` VARCHAR(20),`ref_id` VARCHAR(64),INDEX signal_name_index (`name`),INDEX signal_customer_id_index (`customer_id`),INDEX signal_ref_type_index (`ref_type`),INDEX signal_ref_id_index (`ref_id`),INDEX signal_name_timeunit_customer_id_index (`name`,`timeunit`,`customer_id`),INDEX signal_name_timeunit_customer_id_date_index (`name`,`timeunit`,`customer_id`,`date`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// DBDropSignalTable will drop the Signal table
func DBDropSignalTable(ctx context.Context, db DB) error {
	q := "DROP TABLE IF EXISTS `signal`"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBDropSignalTableTx will drop the Signal table using the provided transaction
func DBDropSignalTableTx(ctx context.Context, tx Tx) error {
	q := "DROP TABLE IF EXISTS `signal`"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// DBCreate will create a new Signal record in the database
func (t *Signal) DBCreate(ctx context.Context, db DB) (sql.Result, error) {
	q := "INSERT INTO `signal` (`signal`.`id`,`signal`.`name`,`signal`.`value`,`signal`.`timeunit`,`signal`.`date`,`signal`.`metadata`,`signal`.`customer_id`,`signal`.`ref_type`,`signal`.`ref_id`) VALUES (?,?,?,?,?,?,?,?,?)"
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Name),
		orm.ToSQLFloat64(t.Value),
		orm.ToSQLInt64(t.Timeunit),
		orm.ToSQLString(t.Date),
		orm.ToSQLString(t.Metadata),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefType),
		orm.ToSQLString(t.RefID),
	)
}

// DBCreateTx will create a new Signal record in the database using the provided transaction
func (t *Signal) DBCreateTx(ctx context.Context, tx Tx) (sql.Result, error) {
	q := "INSERT INTO `signal` (`signal`.`id`,`signal`.`name`,`signal`.`value`,`signal`.`timeunit`,`signal`.`date`,`signal`.`metadata`,`signal`.`customer_id`,`signal`.`ref_type`,`signal`.`ref_id`) VALUES (?,?,?,?,?,?,?,?,?)"
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Name),
		orm.ToSQLFloat64(t.Value),
		orm.ToSQLInt64(t.Timeunit),
		orm.ToSQLString(t.Date),
		orm.ToSQLString(t.Metadata),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefType),
		orm.ToSQLString(t.RefID),
	)
}

// DBCreateIgnoreDuplicate will upsert the Signal record in the database
func (t *Signal) DBCreateIgnoreDuplicate(ctx context.Context, db DB) (sql.Result, error) {
	q := "INSERT INTO `signal` (`signal`.`id`,`signal`.`name`,`signal`.`value`,`signal`.`timeunit`,`signal`.`date`,`signal`.`metadata`,`signal`.`customer_id`,`signal`.`ref_type`,`signal`.`ref_id`) VALUES (?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Name),
		orm.ToSQLFloat64(t.Value),
		orm.ToSQLInt64(t.Timeunit),
		orm.ToSQLString(t.Date),
		orm.ToSQLString(t.Metadata),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefType),
		orm.ToSQLString(t.RefID),
	)
}

// DBCreateIgnoreDuplicateTx will upsert the Signal record in the database using the provided transaction
func (t *Signal) DBCreateIgnoreDuplicateTx(ctx context.Context, tx Tx) (sql.Result, error) {
	q := "INSERT INTO `signal` (`signal`.`id`,`signal`.`name`,`signal`.`value`,`signal`.`timeunit`,`signal`.`date`,`signal`.`metadata`,`signal`.`customer_id`,`signal`.`ref_type`,`signal`.`ref_id`) VALUES (?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Name),
		orm.ToSQLFloat64(t.Value),
		orm.ToSQLInt64(t.Timeunit),
		orm.ToSQLString(t.Date),
		orm.ToSQLString(t.Metadata),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefType),
		orm.ToSQLString(t.RefID),
	)
}

// DeleteAllSignals deletes all Signal records in the database with optional filters
func DeleteAllSignals(ctx context.Context, db DB, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(SignalTableName),
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

// DeleteAllSignalsTx deletes all Signal records in the database with optional filters using the provided transaction
func DeleteAllSignalsTx(ctx context.Context, tx Tx, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(SignalTableName),
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

// DBDelete will delete this Signal record in the database
func (t *Signal) DBDelete(ctx context.Context, db DB) (bool, error) {
	q := "DELETE FROM `signal` WHERE `id` = ?"
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

// DBDeleteTx will delete this Signal record in the database using the provided transaction
func (t *Signal) DBDeleteTx(ctx context.Context, tx Tx) (bool, error) {
	q := "DELETE FROM `signal` WHERE `id` = ?"
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

// DBUpdate will update the Signal record in the database
func (t *Signal) DBUpdate(ctx context.Context, db DB) (sql.Result, error) {
	q := "UPDATE `signal` SET `name`=?,`value`=?,`timeunit`=?,`date`=?,`metadata`=?,`customer_id`=?,`ref_type`=?,`ref_id`=? WHERE `id`=?"
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.Name),
		orm.ToSQLFloat64(t.Value),
		orm.ToSQLInt64(t.Timeunit),
		orm.ToSQLString(t.Date),
		orm.ToSQLString(t.Metadata),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefType),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.ID),
	)
}

// DBUpdateTx will update the Signal record in the database using the provided transaction
func (t *Signal) DBUpdateTx(ctx context.Context, tx Tx) (sql.Result, error) {
	q := "UPDATE `signal` SET `name`=?,`value`=?,`timeunit`=?,`date`=?,`metadata`=?,`customer_id`=?,`ref_type`=?,`ref_id`=? WHERE `id`=?"
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.Name),
		orm.ToSQLFloat64(t.Value),
		orm.ToSQLInt64(t.Timeunit),
		orm.ToSQLString(t.Date),
		orm.ToSQLString(t.Metadata),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefType),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.ID),
	)
}

// DBUpsert will upsert the Signal record in the database
func (t *Signal) DBUpsert(ctx context.Context, db DB, conditions ...interface{}) (bool, bool, error) {
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `signal` (`signal`.`id`,`signal`.`name`,`signal`.`value`,`signal`.`timeunit`,`signal`.`date`,`signal`.`metadata`,`signal`.`customer_id`,`signal`.`ref_type`,`signal`.`ref_id`) VALUES (?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `signal` (`signal`.`id`,`signal`.`name`,`signal`.`value`,`signal`.`timeunit`,`signal`.`date`,`signal`.`metadata`,`signal`.`customer_id`,`signal`.`ref_type`,`signal`.`ref_id`) VALUES (?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `name`=VALUES(`name`),`value`=VALUES(`value`),`timeunit`=VALUES(`timeunit`),`date`=VALUES(`date`),`metadata`=VALUES(`metadata`),`customer_id`=VALUES(`customer_id`),`ref_type`=VALUES(`ref_type`),`ref_id`=VALUES(`ref_id`)"
	}
	r, err := db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Name),
		orm.ToSQLFloat64(t.Value),
		orm.ToSQLInt64(t.Timeunit),
		orm.ToSQLString(t.Date),
		orm.ToSQLString(t.Metadata),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefType),
		orm.ToSQLString(t.RefID),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBUpsertTx will upsert the Signal record in the database using the provided transaction
func (t *Signal) DBUpsertTx(ctx context.Context, tx Tx, conditions ...interface{}) (bool, bool, error) {
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `signal` (`signal`.`id`,`signal`.`name`,`signal`.`value`,`signal`.`timeunit`,`signal`.`date`,`signal`.`metadata`,`signal`.`customer_id`,`signal`.`ref_type`,`signal`.`ref_id`) VALUES (?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `signal` (`signal`.`id`,`signal`.`name`,`signal`.`value`,`signal`.`timeunit`,`signal`.`date`,`signal`.`metadata`,`signal`.`customer_id`,`signal`.`ref_type`,`signal`.`ref_id`) VALUES (?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `name`=VALUES(`name`),`value`=VALUES(`value`),`timeunit`=VALUES(`timeunit`),`date`=VALUES(`date`),`metadata`=VALUES(`metadata`),`customer_id`=VALUES(`customer_id`),`ref_type`=VALUES(`ref_type`),`ref_id`=VALUES(`ref_id`)"
	}
	r, err := tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Name),
		orm.ToSQLFloat64(t.Value),
		orm.ToSQLInt64(t.Timeunit),
		orm.ToSQLString(t.Date),
		orm.ToSQLString(t.Metadata),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefType),
		orm.ToSQLString(t.RefID),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBFindOne will find a Signal record in the database with the primary key
func (t *Signal) DBFindOne(ctx context.Context, db DB, value string) (bool, error) {
	q := "SELECT * FROM `signal` WHERE `id` = ? LIMIT 1"
	row := db.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Name sql.NullString
	var _Value sql.NullFloat64
	var _Timeunit sql.NullInt64
	var _Date sql.NullString
	var _Metadata sql.NullString
	var _CustomerID sql.NullString
	var _RefType sql.NullString
	var _RefID sql.NullString
	err := row.Scan(
		&_ID,
		&_Name,
		&_Value,
		&_Timeunit,
		&_Date,
		&_Metadata,
		&_CustomerID,
		&_RefType,
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
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _Value.Valid {
		t.SetValue(_Value.Float64)
	}
	if _Timeunit.Valid {
		t.SetTimeunit(int32(_Timeunit.Int64))
	}
	if _Date.Valid {
		t.SetDate(_Date.String)
	}
	if _Metadata.Valid {
		t.SetMetadata(_Metadata.String)
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
	return true, nil
}

// DBFindOneTx will find a Signal record in the database with the primary key using the provided transaction
func (t *Signal) DBFindOneTx(ctx context.Context, tx Tx, value string) (bool, error) {
	q := "SELECT * FROM `signal` WHERE `id` = ? LIMIT 1"
	row := tx.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Name sql.NullString
	var _Value sql.NullFloat64
	var _Timeunit sql.NullInt64
	var _Date sql.NullString
	var _Metadata sql.NullString
	var _CustomerID sql.NullString
	var _RefType sql.NullString
	var _RefID sql.NullString
	err := row.Scan(
		&_ID,
		&_Name,
		&_Value,
		&_Timeunit,
		&_Date,
		&_Metadata,
		&_CustomerID,
		&_RefType,
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
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _Value.Valid {
		t.SetValue(_Value.Float64)
	}
	if _Timeunit.Valid {
		t.SetTimeunit(int32(_Timeunit.Int64))
	}
	if _Date.Valid {
		t.SetDate(_Date.String)
	}
	if _Metadata.Valid {
		t.SetMetadata(_Metadata.String)
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
	return true, nil
}

// FindSignals will find a Signal record in the database with the provided parameters
func FindSignals(ctx context.Context, db DB, _params ...interface{}) ([]*Signal, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("name"),
		orm.Column("value"),
		orm.Column("timeunit"),
		orm.Column("date"),
		orm.Column("metadata"),
		orm.Column("customer_id"),
		orm.Column("ref_type"),
		orm.Column("ref_id"),
		orm.Table(SignalTableName),
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
	results := make([]*Signal, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Name sql.NullString
		var _Value sql.NullFloat64
		var _Timeunit sql.NullInt64
		var _Date sql.NullString
		var _Metadata sql.NullString
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Name,
			&_Value,
			&_Timeunit,
			&_Date,
			&_Metadata,
			&_CustomerID,
			&_RefType,
			&_RefID,
		)
		if err != nil {
			return nil, err
		}
		t := &Signal{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _Value.Valid {
			t.SetValue(_Value.Float64)
		}
		if _Timeunit.Valid {
			t.SetTimeunit(int32(_Timeunit.Int64))
		}
		if _Date.Valid {
			t.SetDate(_Date.String)
		}
		if _Metadata.Valid {
			t.SetMetadata(_Metadata.String)
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
		results = append(results, t)
	}
	return results, nil
}

// FindSignalsTx will find a Signal record in the database with the provided parameters using the provided transaction
func FindSignalsTx(ctx context.Context, tx Tx, _params ...interface{}) ([]*Signal, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("name"),
		orm.Column("value"),
		orm.Column("timeunit"),
		orm.Column("date"),
		orm.Column("metadata"),
		orm.Column("customer_id"),
		orm.Column("ref_type"),
		orm.Column("ref_id"),
		orm.Table(SignalTableName),
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
	results := make([]*Signal, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Name sql.NullString
		var _Value sql.NullFloat64
		var _Timeunit sql.NullInt64
		var _Date sql.NullString
		var _Metadata sql.NullString
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Name,
			&_Value,
			&_Timeunit,
			&_Date,
			&_Metadata,
			&_CustomerID,
			&_RefType,
			&_RefID,
		)
		if err != nil {
			return nil, err
		}
		t := &Signal{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _Value.Valid {
			t.SetValue(_Value.Float64)
		}
		if _Timeunit.Valid {
			t.SetTimeunit(int32(_Timeunit.Int64))
		}
		if _Date.Valid {
			t.SetDate(_Date.String)
		}
		if _Metadata.Valid {
			t.SetMetadata(_Metadata.String)
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
		results = append(results, t)
	}
	return results, nil
}

// DBFind will find a Signal record in the database with the provided parameters
func (t *Signal) DBFind(ctx context.Context, db DB, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("name"),
		orm.Column("value"),
		orm.Column("timeunit"),
		orm.Column("date"),
		orm.Column("metadata"),
		orm.Column("customer_id"),
		orm.Column("ref_type"),
		orm.Column("ref_id"),
		orm.Table(SignalTableName),
	}
	if len(_params) > 0 {
		for _, param := range _params {
			params = append(params, param)
		}
	}
	q, p := orm.BuildQuery(params...)
	row := db.QueryRowContext(ctx, q, p...)
	var _ID sql.NullString
	var _Name sql.NullString
	var _Value sql.NullFloat64
	var _Timeunit sql.NullInt64
	var _Date sql.NullString
	var _Metadata sql.NullString
	var _CustomerID sql.NullString
	var _RefType sql.NullString
	var _RefID sql.NullString
	err := row.Scan(
		&_ID,
		&_Name,
		&_Value,
		&_Timeunit,
		&_Date,
		&_Metadata,
		&_CustomerID,
		&_RefType,
		&_RefID,
	)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _Value.Valid {
		t.SetValue(_Value.Float64)
	}
	if _Timeunit.Valid {
		t.SetTimeunit(int32(_Timeunit.Int64))
	}
	if _Date.Valid {
		t.SetDate(_Date.String)
	}
	if _Metadata.Valid {
		t.SetMetadata(_Metadata.String)
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
	return true, nil
}

// DBFindTx will find a Signal record in the database with the provided parameters using the provided transaction
func (t *Signal) DBFindTx(ctx context.Context, tx Tx, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("name"),
		orm.Column("value"),
		orm.Column("timeunit"),
		orm.Column("date"),
		orm.Column("metadata"),
		orm.Column("customer_id"),
		orm.Column("ref_type"),
		orm.Column("ref_id"),
		orm.Table(SignalTableName),
	}
	if len(_params) > 0 {
		for _, param := range _params {
			params = append(params, param)
		}
	}
	q, p := orm.BuildQuery(params...)
	row := tx.QueryRowContext(ctx, q, p...)
	var _ID sql.NullString
	var _Name sql.NullString
	var _Value sql.NullFloat64
	var _Timeunit sql.NullInt64
	var _Date sql.NullString
	var _Metadata sql.NullString
	var _CustomerID sql.NullString
	var _RefType sql.NullString
	var _RefID sql.NullString
	err := row.Scan(
		&_ID,
		&_Name,
		&_Value,
		&_Timeunit,
		&_Date,
		&_Metadata,
		&_CustomerID,
		&_RefType,
		&_RefID,
	)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _Value.Valid {
		t.SetValue(_Value.Float64)
	}
	if _Timeunit.Valid {
		t.SetTimeunit(int32(_Timeunit.Int64))
	}
	if _Date.Valid {
		t.SetDate(_Date.String)
	}
	if _Metadata.Valid {
		t.SetMetadata(_Metadata.String)
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
	return true, nil
}

// CountSignals will find the count of Signal records in the database
func CountSignals(ctx context.Context, db DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(SignalTableName),
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

// CountSignalsTx will find the count of Signal records in the database using the provided transaction
func CountSignalsTx(ctx context.Context, tx Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(SignalTableName),
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

// DBCount will find the count of Signal records in the database
func (t *Signal) DBCount(ctx context.Context, db DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(SignalTableName),
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

// DBCountTx will find the count of Signal records in the database using the provided transaction
func (t *Signal) DBCountTx(ctx context.Context, tx Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(SignalTableName),
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

// DBExists will return true if the Signal record exists in the database
func (t *Signal) DBExists(ctx context.Context, db DB) (bool, error) {
	q := "SELECT * FROM `signal` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := db.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// DBExistsTx will return true if the Signal record exists in the database using the provided transaction
func (t *Signal) DBExistsTx(ctx context.Context, tx Tx) (bool, error) {
	q := "SELECT * FROM `signal` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := tx.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// PrimaryKeyColumn returns the column name for the primary key
func (t *Signal) PrimaryKeyColumn() string {
	return SignalColumnID
}

// PrimaryKeyColumnType returns the primary key column Go type as a string
func (t *Signal) PrimaryKeyColumnType() string {
	return "string"
}

// PrimaryKey returns the primary key column value
func (t *Signal) PrimaryKey() interface{} {
	return t.ID
}
