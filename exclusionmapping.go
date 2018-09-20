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

var _ Model = (*ExclusionMapping)(nil)
var _ CSVWriter = (*ExclusionMapping)(nil)
var _ JSONWriter = (*ExclusionMapping)(nil)

// ExclusionMappingTableName is the name of the table in SQL
const ExclusionMappingTableName = "exclusion_mapping"

var ExclusionMappingColumns = []string{
	"id",
	"ref_id",
	"ref_type",
	"user_id",
	"date",
	"customer_id",
}

// ExclusionMapping table
type ExclusionMapping struct {
	CustomerID string  `json:"customer_id"`
	Date       *int64  `json:"date,omitempty"`
	ID         string  `json:"id"`
	RefID      string  `json:"ref_id"`
	RefType    string  `json:"ref_type"`
	UserID     *string `json:"user_id,omitempty"`
}

// TableName returns the SQL table name for ExclusionMapping and satifies the Model interface
func (t *ExclusionMapping) TableName() string {
	return ExclusionMappingTableName
}

// ToCSV will serialize the ExclusionMapping instance to a CSV compatible array of strings
func (t *ExclusionMapping) ToCSV() []string {
	return []string{
		t.ID,
		t.RefID,
		t.RefType,
		toCSVString(t.UserID),
		toCSVString(t.Date),
		t.CustomerID,
	}
}

// WriteCSV will serialize the ExclusionMapping instance to the writer as CSV and satisfies the CSVWriter interface
func (t *ExclusionMapping) WriteCSV(w *csv.Writer) error {
	return w.Write(t.ToCSV())
}

// WriteJSON will serialize the ExclusionMapping instance to the writer as JSON and satisfies the JSONWriter interface
func (t *ExclusionMapping) WriteJSON(w io.Writer, indent ...bool) error {
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

// NewExclusionMappingReader creates a JSON reader which can read in ExclusionMapping objects serialized as JSON either as an array, single object or json new lines
// and writes each ExclusionMapping to the channel provided
func NewExclusionMappingReader(r io.Reader, ch chan<- ExclusionMapping) error {
	return orm.Deserialize(r, func(buf json.RawMessage) error {
		dec := json.NewDecoder(bytes.NewBuffer(buf))
		e := ExclusionMapping{}
		if err := dec.Decode(&e); err != nil {
			return err
		}
		ch <- e
		return nil
	})
}

// NewCSVExclusionMappingReaderDir will read the reader as CSV and emit each record to the channel provided
func NewCSVExclusionMappingReader(r io.Reader, ch chan<- ExclusionMapping) error {
	cr := csv.NewReader(r)
	for {
		record, err := cr.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		ch <- ExclusionMapping{
			ID:         record[0],
			RefID:      record[1],
			RefType:    record[2],
			UserID:     fromStringPointer(record[3]),
			Date:       fromCSVInt64Pointer(record[4]),
			CustomerID: record[5],
		}
	}
	return nil
}

// NewCSVExclusionMappingReaderFile will read the file as a CSV and emit each record to the channel provided
func NewCSVExclusionMappingReaderFile(fp string, ch chan<- ExclusionMapping) error {
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
	return NewCSVExclusionMappingReader(fc, ch)
}

// NewCSVExclusionMappingReaderDir will read the exclusion_mapping.csv.gz file as a CSV and emit each record to the channel provided
func NewCSVExclusionMappingReaderDir(dir string, ch chan<- ExclusionMapping) error {
	return NewCSVExclusionMappingReaderFile(filepath.Join(dir, "exclusion_mapping.csv.gz"), ch)
}

// ExclusionMappingCSVDeduper is a function callback which takes the existing value (a) and the new value (b)
// and the return value should be the one to use (or a new one, if applicable). return nil
// to skip processing of this record
type ExclusionMappingCSVDeduper func(a ExclusionMapping, b ExclusionMapping) *ExclusionMapping

// ExclusionMappingCSVDedupeDisabled is set on whether the CSV writer should de-dupe values by key
var ExclusionMappingCSVDedupeDisabled bool

// NewExclusionMappingCSVWriterSize creates a batch writer that will write each ExclusionMapping into a CSV file
func NewExclusionMappingCSVWriterSize(w io.Writer, size int, dedupers ...ExclusionMappingCSVDeduper) (chan ExclusionMapping, chan bool, error) {
	cw := csv.NewWriter(w)
	ch := make(chan ExclusionMapping, size)
	done := make(chan bool)
	go func() {
		defer func() { done <- true }()
		dodedupe := !ExclusionMappingCSVDedupeDisabled
		var kv map[string]*ExclusionMapping
		var deduper ExclusionMappingCSVDeduper
		if dedupers != nil && len(dedupers) > 0 {
			deduper = dedupers[0]
			dodedupe = true
		}
		if dodedupe {
			kv = make(map[string]*ExclusionMapping)
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

// ExclusionMappingCSVDefaultSize is the default channel buffer size if not provided
var ExclusionMappingCSVDefaultSize = 100

// NewExclusionMappingCSVWriter creates a batch writer that will write each ExclusionMapping into a CSV file
func NewExclusionMappingCSVWriter(w io.Writer, dedupers ...ExclusionMappingCSVDeduper) (chan ExclusionMapping, chan bool, error) {
	return NewExclusionMappingCSVWriterSize(w, ExclusionMappingCSVDefaultSize, dedupers...)
}

// NewExclusionMappingCSVWriterDir creates a batch writer that will write each ExclusionMapping into a CSV file named exclusion_mapping.csv.gz in dir
func NewExclusionMappingCSVWriterDir(dir string, dedupers ...ExclusionMappingCSVDeduper) (chan ExclusionMapping, chan bool, error) {
	return NewExclusionMappingCSVWriterFile(filepath.Join(dir, "exclusion_mapping.csv.gz"), dedupers...)
}

// NewExclusionMappingCSVWriterFile creates a batch writer that will write each ExclusionMapping into a CSV file
func NewExclusionMappingCSVWriterFile(fn string, dedupers ...ExclusionMappingCSVDeduper) (chan ExclusionMapping, chan bool, error) {
	f, err := os.Create(fn)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening CSV file %s. %v", fn, err)
	}
	var fc io.WriteCloser = f
	if filepath.Ext(fn) == ".gz" {
		w, _ := gzip.NewWriterLevel(f, gzip.BestCompression)
		fc = w
	}
	ch, done, err := NewExclusionMappingCSVWriter(fc, dedupers...)
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

type ExclusionMappingDBAction func(ctx context.Context, db *sql.DB, record ExclusionMapping) error

// NewExclusionMappingDBWriterSize creates a DB writer that will write each issue into the DB
func NewExclusionMappingDBWriterSize(ctx context.Context, db *sql.DB, errors chan<- error, size int, actions ...ExclusionMappingDBAction) (chan ExclusionMapping, chan bool, error) {
	ch := make(chan ExclusionMapping, size)
	done := make(chan bool)
	var action ExclusionMappingDBAction
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

// NewExclusionMappingDBWriter creates a DB writer that will write each issue into the DB
func NewExclusionMappingDBWriter(ctx context.Context, db *sql.DB, errors chan<- error, actions ...ExclusionMappingDBAction) (chan ExclusionMapping, chan bool, error) {
	return NewExclusionMappingDBWriterSize(ctx, db, errors, 100, actions...)
}

// ExclusionMappingColumnID is the ID SQL column name for the ExclusionMapping table
const ExclusionMappingColumnID = "id"

// ExclusionMappingEscapedColumnID is the escaped ID SQL column name for the ExclusionMapping table
const ExclusionMappingEscapedColumnID = "`id`"

// ExclusionMappingColumnRefID is the RefID SQL column name for the ExclusionMapping table
const ExclusionMappingColumnRefID = "ref_id"

// ExclusionMappingEscapedColumnRefID is the escaped RefID SQL column name for the ExclusionMapping table
const ExclusionMappingEscapedColumnRefID = "`ref_id`"

// ExclusionMappingColumnRefType is the RefType SQL column name for the ExclusionMapping table
const ExclusionMappingColumnRefType = "ref_type"

// ExclusionMappingEscapedColumnRefType is the escaped RefType SQL column name for the ExclusionMapping table
const ExclusionMappingEscapedColumnRefType = "`ref_type`"

// ExclusionMappingColumnUserID is the UserID SQL column name for the ExclusionMapping table
const ExclusionMappingColumnUserID = "user_id"

// ExclusionMappingEscapedColumnUserID is the escaped UserID SQL column name for the ExclusionMapping table
const ExclusionMappingEscapedColumnUserID = "`user_id`"

// ExclusionMappingColumnDate is the Date SQL column name for the ExclusionMapping table
const ExclusionMappingColumnDate = "date"

// ExclusionMappingEscapedColumnDate is the escaped Date SQL column name for the ExclusionMapping table
const ExclusionMappingEscapedColumnDate = "`date`"

// ExclusionMappingColumnCustomerID is the CustomerID SQL column name for the ExclusionMapping table
const ExclusionMappingColumnCustomerID = "customer_id"

// ExclusionMappingEscapedColumnCustomerID is the escaped CustomerID SQL column name for the ExclusionMapping table
const ExclusionMappingEscapedColumnCustomerID = "`customer_id`"

// GetID will return the ExclusionMapping ID value
func (t *ExclusionMapping) GetID() string {
	return t.ID
}

// SetID will set the ExclusionMapping ID value
func (t *ExclusionMapping) SetID(v string) {
	t.ID = v
}

// FindExclusionMappingByID will find a ExclusionMapping by ID
func FindExclusionMappingByID(ctx context.Context, db *sql.DB, value string) (*ExclusionMapping, error) {
	q := "SELECT `exclusion_mapping`.`id`,`exclusion_mapping`.`ref_id`,`exclusion_mapping`.`ref_type`,`exclusion_mapping`.`user_id`,`exclusion_mapping`.`date`,`exclusion_mapping`.`customer_id` FROM `exclusion_mapping` WHERE `id` = ?"
	var _ID sql.NullString
	var _RefID sql.NullString
	var _RefType sql.NullString
	var _UserID sql.NullString
	var _Date sql.NullInt64
	var _CustomerID sql.NullString
	err := db.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_RefID,
		&_RefType,
		&_UserID,
		&_Date,
		&_CustomerID,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &ExclusionMapping{}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _RefID.Valid {
		t.SetRefID(_RefID.String)
	}
	if _RefType.Valid {
		t.SetRefType(_RefType.String)
	}
	if _UserID.Valid {
		t.SetUserID(_UserID.String)
	}
	if _Date.Valid {
		t.SetDate(_Date.Int64)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	return t, nil
}

// FindExclusionMappingByIDTx will find a ExclusionMapping by ID using the provided transaction
func FindExclusionMappingByIDTx(ctx context.Context, tx *sql.Tx, value string) (*ExclusionMapping, error) {
	q := "SELECT `exclusion_mapping`.`id`,`exclusion_mapping`.`ref_id`,`exclusion_mapping`.`ref_type`,`exclusion_mapping`.`user_id`,`exclusion_mapping`.`date`,`exclusion_mapping`.`customer_id` FROM `exclusion_mapping` WHERE `id` = ?"
	var _ID sql.NullString
	var _RefID sql.NullString
	var _RefType sql.NullString
	var _UserID sql.NullString
	var _Date sql.NullInt64
	var _CustomerID sql.NullString
	err := tx.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_RefID,
		&_RefType,
		&_UserID,
		&_Date,
		&_CustomerID,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &ExclusionMapping{}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _RefID.Valid {
		t.SetRefID(_RefID.String)
	}
	if _RefType.Valid {
		t.SetRefType(_RefType.String)
	}
	if _UserID.Valid {
		t.SetUserID(_UserID.String)
	}
	if _Date.Valid {
		t.SetDate(_Date.Int64)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	return t, nil
}

// GetRefID will return the ExclusionMapping RefID value
func (t *ExclusionMapping) GetRefID() string {
	return t.RefID
}

// SetRefID will set the ExclusionMapping RefID value
func (t *ExclusionMapping) SetRefID(v string) {
	t.RefID = v
}

// FindExclusionMappingsByRefID will find all ExclusionMappings by the RefID value
func FindExclusionMappingsByRefID(ctx context.Context, db *sql.DB, value string) ([]*ExclusionMapping, error) {
	q := "SELECT `exclusion_mapping`.`id`,`exclusion_mapping`.`ref_id`,`exclusion_mapping`.`ref_type`,`exclusion_mapping`.`user_id`,`exclusion_mapping`.`date`,`exclusion_mapping`.`customer_id` FROM `exclusion_mapping` WHERE `ref_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*ExclusionMapping, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _RefID sql.NullString
		var _RefType sql.NullString
		var _UserID sql.NullString
		var _Date sql.NullInt64
		var _CustomerID sql.NullString
		err := rows.Scan(
			&_ID,
			&_RefID,
			&_RefType,
			&_UserID,
			&_Date,
			&_CustomerID,
		)
		if err != nil {
			return nil, err
		}
		t := &ExclusionMapping{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		if _RefType.Valid {
			t.SetRefType(_RefType.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _Date.Valid {
			t.SetDate(_Date.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindExclusionMappingsByRefIDTx will find all ExclusionMappings by the RefID value using the provided transaction
func FindExclusionMappingsByRefIDTx(ctx context.Context, tx *sql.Tx, value string) ([]*ExclusionMapping, error) {
	q := "SELECT `exclusion_mapping`.`id`,`exclusion_mapping`.`ref_id`,`exclusion_mapping`.`ref_type`,`exclusion_mapping`.`user_id`,`exclusion_mapping`.`date`,`exclusion_mapping`.`customer_id` FROM `exclusion_mapping` WHERE `ref_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*ExclusionMapping, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _RefID sql.NullString
		var _RefType sql.NullString
		var _UserID sql.NullString
		var _Date sql.NullInt64
		var _CustomerID sql.NullString
		err := rows.Scan(
			&_ID,
			&_RefID,
			&_RefType,
			&_UserID,
			&_Date,
			&_CustomerID,
		)
		if err != nil {
			return nil, err
		}
		t := &ExclusionMapping{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		if _RefType.Valid {
			t.SetRefType(_RefType.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _Date.Valid {
			t.SetDate(_Date.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetRefType will return the ExclusionMapping RefType value
func (t *ExclusionMapping) GetRefType() string {
	return t.RefType
}

// SetRefType will set the ExclusionMapping RefType value
func (t *ExclusionMapping) SetRefType(v string) {
	t.RefType = v
}

// FindExclusionMappingsByRefType will find all ExclusionMappings by the RefType value
func FindExclusionMappingsByRefType(ctx context.Context, db *sql.DB, value string) ([]*ExclusionMapping, error) {
	q := "SELECT `exclusion_mapping`.`id`,`exclusion_mapping`.`ref_id`,`exclusion_mapping`.`ref_type`,`exclusion_mapping`.`user_id`,`exclusion_mapping`.`date`,`exclusion_mapping`.`customer_id` FROM `exclusion_mapping` WHERE `ref_type` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*ExclusionMapping, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _RefID sql.NullString
		var _RefType sql.NullString
		var _UserID sql.NullString
		var _Date sql.NullInt64
		var _CustomerID sql.NullString
		err := rows.Scan(
			&_ID,
			&_RefID,
			&_RefType,
			&_UserID,
			&_Date,
			&_CustomerID,
		)
		if err != nil {
			return nil, err
		}
		t := &ExclusionMapping{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		if _RefType.Valid {
			t.SetRefType(_RefType.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _Date.Valid {
			t.SetDate(_Date.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindExclusionMappingsByRefTypeTx will find all ExclusionMappings by the RefType value using the provided transaction
func FindExclusionMappingsByRefTypeTx(ctx context.Context, tx *sql.Tx, value string) ([]*ExclusionMapping, error) {
	q := "SELECT `exclusion_mapping`.`id`,`exclusion_mapping`.`ref_id`,`exclusion_mapping`.`ref_type`,`exclusion_mapping`.`user_id`,`exclusion_mapping`.`date`,`exclusion_mapping`.`customer_id` FROM `exclusion_mapping` WHERE `ref_type` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*ExclusionMapping, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _RefID sql.NullString
		var _RefType sql.NullString
		var _UserID sql.NullString
		var _Date sql.NullInt64
		var _CustomerID sql.NullString
		err := rows.Scan(
			&_ID,
			&_RefID,
			&_RefType,
			&_UserID,
			&_Date,
			&_CustomerID,
		)
		if err != nil {
			return nil, err
		}
		t := &ExclusionMapping{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		if _RefType.Valid {
			t.SetRefType(_RefType.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _Date.Valid {
			t.SetDate(_Date.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetUserID will return the ExclusionMapping UserID value
func (t *ExclusionMapping) GetUserID() string {
	if t.UserID == nil {
		return ""
	}
	return *t.UserID
}

// SetUserID will set the ExclusionMapping UserID value
func (t *ExclusionMapping) SetUserID(v string) {
	t.UserID = &v
}

// GetDate will return the ExclusionMapping Date value
func (t *ExclusionMapping) GetDate() int64 {
	if t.Date == nil {
		return int64(0)
	}
	return *t.Date
}

// SetDate will set the ExclusionMapping Date value
func (t *ExclusionMapping) SetDate(v int64) {
	t.Date = &v
}

// GetCustomerID will return the ExclusionMapping CustomerID value
func (t *ExclusionMapping) GetCustomerID() string {
	return t.CustomerID
}

// SetCustomerID will set the ExclusionMapping CustomerID value
func (t *ExclusionMapping) SetCustomerID(v string) {
	t.CustomerID = v
}

// FindExclusionMappingsByCustomerID will find all ExclusionMappings by the CustomerID value
func FindExclusionMappingsByCustomerID(ctx context.Context, db *sql.DB, value string) ([]*ExclusionMapping, error) {
	q := "SELECT `exclusion_mapping`.`id`,`exclusion_mapping`.`ref_id`,`exclusion_mapping`.`ref_type`,`exclusion_mapping`.`user_id`,`exclusion_mapping`.`date`,`exclusion_mapping`.`customer_id` FROM `exclusion_mapping` WHERE `customer_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*ExclusionMapping, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _RefID sql.NullString
		var _RefType sql.NullString
		var _UserID sql.NullString
		var _Date sql.NullInt64
		var _CustomerID sql.NullString
		err := rows.Scan(
			&_ID,
			&_RefID,
			&_RefType,
			&_UserID,
			&_Date,
			&_CustomerID,
		)
		if err != nil {
			return nil, err
		}
		t := &ExclusionMapping{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		if _RefType.Valid {
			t.SetRefType(_RefType.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _Date.Valid {
			t.SetDate(_Date.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindExclusionMappingsByCustomerIDTx will find all ExclusionMappings by the CustomerID value using the provided transaction
func FindExclusionMappingsByCustomerIDTx(ctx context.Context, tx *sql.Tx, value string) ([]*ExclusionMapping, error) {
	q := "SELECT `exclusion_mapping`.`id`,`exclusion_mapping`.`ref_id`,`exclusion_mapping`.`ref_type`,`exclusion_mapping`.`user_id`,`exclusion_mapping`.`date`,`exclusion_mapping`.`customer_id` FROM `exclusion_mapping` WHERE `customer_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*ExclusionMapping, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _RefID sql.NullString
		var _RefType sql.NullString
		var _UserID sql.NullString
		var _Date sql.NullInt64
		var _CustomerID sql.NullString
		err := rows.Scan(
			&_ID,
			&_RefID,
			&_RefType,
			&_UserID,
			&_Date,
			&_CustomerID,
		)
		if err != nil {
			return nil, err
		}
		t := &ExclusionMapping{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		if _RefType.Valid {
			t.SetRefType(_RefType.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _Date.Valid {
			t.SetDate(_Date.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

func (t *ExclusionMapping) toTimestamp(value time.Time) *timestamp.Timestamp {
	ts, _ := ptypes.TimestampProto(value)
	return ts
}

// DBCreateExclusionMappingTable will create the ExclusionMapping table
func DBCreateExclusionMappingTable(ctx context.Context, db *sql.DB) error {
	q := "CREATE TABLE `exclusion_mapping` (`id` VARCHAR(64) NOT NULL PRIMARY KEY,`ref_id` VARCHAR(64) NOT NULL,`ref_type` VARCHAR(20) NOT NULL,`user_id`VARCHAR(64),`date`BIGINT UNSIGNED,`customer_id` VARCHAR(64) NOT NULL,INDEX exclusion_mapping_ref_id_index (`ref_id`),INDEX exclusion_mapping_ref_type_index (`ref_type`),INDEX exclusion_mapping_customer_id_index (`customer_id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBCreateExclusionMappingTableTx will create the ExclusionMapping table using the provided transction
func DBCreateExclusionMappingTableTx(ctx context.Context, tx *sql.Tx) error {
	q := "CREATE TABLE `exclusion_mapping` (`id` VARCHAR(64) NOT NULL PRIMARY KEY,`ref_id` VARCHAR(64) NOT NULL,`ref_type` VARCHAR(20) NOT NULL,`user_id`VARCHAR(64),`date`BIGINT UNSIGNED,`customer_id` VARCHAR(64) NOT NULL,INDEX exclusion_mapping_ref_id_index (`ref_id`),INDEX exclusion_mapping_ref_type_index (`ref_type`),INDEX exclusion_mapping_customer_id_index (`customer_id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// DBDropExclusionMappingTable will drop the ExclusionMapping table
func DBDropExclusionMappingTable(ctx context.Context, db *sql.DB) error {
	q := "DROP TABLE IF EXISTS `exclusion_mapping`"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBDropExclusionMappingTableTx will drop the ExclusionMapping table using the provided transaction
func DBDropExclusionMappingTableTx(ctx context.Context, tx *sql.Tx) error {
	q := "DROP TABLE IF EXISTS `exclusion_mapping`"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// DBCreate will create a new ExclusionMapping record in the database
func (t *ExclusionMapping) DBCreate(ctx context.Context, db *sql.DB) (sql.Result, error) {
	q := "INSERT INTO `exclusion_mapping` (`exclusion_mapping`.`id`,`exclusion_mapping`.`ref_id`,`exclusion_mapping`.`ref_type`,`exclusion_mapping`.`user_id`,`exclusion_mapping`.`date`,`exclusion_mapping`.`customer_id`) VALUES (?,?,?,?,?,?)"
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.RefType),
		orm.ToSQLString(t.UserID),
		orm.ToSQLInt64(t.Date),
		orm.ToSQLString(t.CustomerID),
	)
}

// DBCreateTx will create a new ExclusionMapping record in the database using the provided transaction
func (t *ExclusionMapping) DBCreateTx(ctx context.Context, tx *sql.Tx) (sql.Result, error) {
	q := "INSERT INTO `exclusion_mapping` (`exclusion_mapping`.`id`,`exclusion_mapping`.`ref_id`,`exclusion_mapping`.`ref_type`,`exclusion_mapping`.`user_id`,`exclusion_mapping`.`date`,`exclusion_mapping`.`customer_id`) VALUES (?,?,?,?,?,?)"
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.RefType),
		orm.ToSQLString(t.UserID),
		orm.ToSQLInt64(t.Date),
		orm.ToSQLString(t.CustomerID),
	)
}

// DBCreateIgnoreDuplicate will upsert the ExclusionMapping record in the database
func (t *ExclusionMapping) DBCreateIgnoreDuplicate(ctx context.Context, db *sql.DB) (sql.Result, error) {
	q := "INSERT INTO `exclusion_mapping` (`exclusion_mapping`.`id`,`exclusion_mapping`.`ref_id`,`exclusion_mapping`.`ref_type`,`exclusion_mapping`.`user_id`,`exclusion_mapping`.`date`,`exclusion_mapping`.`customer_id`) VALUES (?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.RefType),
		orm.ToSQLString(t.UserID),
		orm.ToSQLInt64(t.Date),
		orm.ToSQLString(t.CustomerID),
	)
}

// DBCreateIgnoreDuplicateTx will upsert the ExclusionMapping record in the database using the provided transaction
func (t *ExclusionMapping) DBCreateIgnoreDuplicateTx(ctx context.Context, tx *sql.Tx) (sql.Result, error) {
	q := "INSERT INTO `exclusion_mapping` (`exclusion_mapping`.`id`,`exclusion_mapping`.`ref_id`,`exclusion_mapping`.`ref_type`,`exclusion_mapping`.`user_id`,`exclusion_mapping`.`date`,`exclusion_mapping`.`customer_id`) VALUES (?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.RefType),
		orm.ToSQLString(t.UserID),
		orm.ToSQLInt64(t.Date),
		orm.ToSQLString(t.CustomerID),
	)
}

// DeleteAllExclusionMappings deletes all ExclusionMapping records in the database with optional filters
func DeleteAllExclusionMappings(ctx context.Context, db *sql.DB, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(ExclusionMappingTableName),
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

// DeleteAllExclusionMappingsTx deletes all ExclusionMapping records in the database with optional filters using the provided transaction
func DeleteAllExclusionMappingsTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(ExclusionMappingTableName),
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

// DBDelete will delete this ExclusionMapping record in the database
func (t *ExclusionMapping) DBDelete(ctx context.Context, db *sql.DB) (bool, error) {
	q := "DELETE FROM `exclusion_mapping` WHERE `id` = ?"
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

// DBDeleteTx will delete this ExclusionMapping record in the database using the provided transaction
func (t *ExclusionMapping) DBDeleteTx(ctx context.Context, tx *sql.Tx) (bool, error) {
	q := "DELETE FROM `exclusion_mapping` WHERE `id` = ?"
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

// DBUpdate will update the ExclusionMapping record in the database
func (t *ExclusionMapping) DBUpdate(ctx context.Context, db *sql.DB) (sql.Result, error) {
	q := "UPDATE `exclusion_mapping` SET `ref_id`=?,`ref_type`=?,`user_id`=?,`date`=?,`customer_id`=? WHERE `id`=?"
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.RefType),
		orm.ToSQLString(t.UserID),
		orm.ToSQLInt64(t.Date),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.ID),
	)
}

// DBUpdateTx will update the ExclusionMapping record in the database using the provided transaction
func (t *ExclusionMapping) DBUpdateTx(ctx context.Context, tx *sql.Tx) (sql.Result, error) {
	q := "UPDATE `exclusion_mapping` SET `ref_id`=?,`ref_type`=?,`user_id`=?,`date`=?,`customer_id`=? WHERE `id`=?"
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.RefType),
		orm.ToSQLString(t.UserID),
		orm.ToSQLInt64(t.Date),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.ID),
	)
}

// DBUpsert will upsert the ExclusionMapping record in the database
func (t *ExclusionMapping) DBUpsert(ctx context.Context, db *sql.DB, conditions ...interface{}) (bool, bool, error) {
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `exclusion_mapping` (`exclusion_mapping`.`id`,`exclusion_mapping`.`ref_id`,`exclusion_mapping`.`ref_type`,`exclusion_mapping`.`user_id`,`exclusion_mapping`.`date`,`exclusion_mapping`.`customer_id`) VALUES (?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `exclusion_mapping` (`exclusion_mapping`.`id`,`exclusion_mapping`.`ref_id`,`exclusion_mapping`.`ref_type`,`exclusion_mapping`.`user_id`,`exclusion_mapping`.`date`,`exclusion_mapping`.`customer_id`) VALUES (?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `ref_id`=VALUES(`ref_id`),`ref_type`=VALUES(`ref_type`),`user_id`=VALUES(`user_id`),`date`=VALUES(`date`),`customer_id`=VALUES(`customer_id`)"
	}
	r, err := db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.RefType),
		orm.ToSQLString(t.UserID),
		orm.ToSQLInt64(t.Date),
		orm.ToSQLString(t.CustomerID),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBUpsertTx will upsert the ExclusionMapping record in the database using the provided transaction
func (t *ExclusionMapping) DBUpsertTx(ctx context.Context, tx *sql.Tx, conditions ...interface{}) (bool, bool, error) {
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `exclusion_mapping` (`exclusion_mapping`.`id`,`exclusion_mapping`.`ref_id`,`exclusion_mapping`.`ref_type`,`exclusion_mapping`.`user_id`,`exclusion_mapping`.`date`,`exclusion_mapping`.`customer_id`) VALUES (?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `exclusion_mapping` (`exclusion_mapping`.`id`,`exclusion_mapping`.`ref_id`,`exclusion_mapping`.`ref_type`,`exclusion_mapping`.`user_id`,`exclusion_mapping`.`date`,`exclusion_mapping`.`customer_id`) VALUES (?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `ref_id`=VALUES(`ref_id`),`ref_type`=VALUES(`ref_type`),`user_id`=VALUES(`user_id`),`date`=VALUES(`date`),`customer_id`=VALUES(`customer_id`)"
	}
	r, err := tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.RefType),
		orm.ToSQLString(t.UserID),
		orm.ToSQLInt64(t.Date),
		orm.ToSQLString(t.CustomerID),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBFindOne will find a ExclusionMapping record in the database with the primary key
func (t *ExclusionMapping) DBFindOne(ctx context.Context, db *sql.DB, value string) (bool, error) {
	q := "SELECT `exclusion_mapping`.`id`,`exclusion_mapping`.`ref_id`,`exclusion_mapping`.`ref_type`,`exclusion_mapping`.`user_id`,`exclusion_mapping`.`date`,`exclusion_mapping`.`customer_id` FROM `exclusion_mapping` WHERE `id` = ? LIMIT 1"
	row := db.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _RefID sql.NullString
	var _RefType sql.NullString
	var _UserID sql.NullString
	var _Date sql.NullInt64
	var _CustomerID sql.NullString
	err := row.Scan(
		&_ID,
		&_RefID,
		&_RefType,
		&_UserID,
		&_Date,
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
	if _RefID.Valid {
		t.SetRefID(_RefID.String)
	}
	if _RefType.Valid {
		t.SetRefType(_RefType.String)
	}
	if _UserID.Valid {
		t.SetUserID(_UserID.String)
	}
	if _Date.Valid {
		t.SetDate(_Date.Int64)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	return true, nil
}

// DBFindOneTx will find a ExclusionMapping record in the database with the primary key using the provided transaction
func (t *ExclusionMapping) DBFindOneTx(ctx context.Context, tx *sql.Tx, value string) (bool, error) {
	q := "SELECT `exclusion_mapping`.`id`,`exclusion_mapping`.`ref_id`,`exclusion_mapping`.`ref_type`,`exclusion_mapping`.`user_id`,`exclusion_mapping`.`date`,`exclusion_mapping`.`customer_id` FROM `exclusion_mapping` WHERE `id` = ? LIMIT 1"
	row := tx.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _RefID sql.NullString
	var _RefType sql.NullString
	var _UserID sql.NullString
	var _Date sql.NullInt64
	var _CustomerID sql.NullString
	err := row.Scan(
		&_ID,
		&_RefID,
		&_RefType,
		&_UserID,
		&_Date,
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
	if _RefID.Valid {
		t.SetRefID(_RefID.String)
	}
	if _RefType.Valid {
		t.SetRefType(_RefType.String)
	}
	if _UserID.Valid {
		t.SetUserID(_UserID.String)
	}
	if _Date.Valid {
		t.SetDate(_Date.Int64)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	return true, nil
}

// FindExclusionMappings will find a ExclusionMapping record in the database with the provided parameters
func FindExclusionMappings(ctx context.Context, db *sql.DB, _params ...interface{}) ([]*ExclusionMapping, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("ref_id"),
		orm.Column("ref_type"),
		orm.Column("user_id"),
		orm.Column("date"),
		orm.Column("customer_id"),
		orm.Table(ExclusionMappingTableName),
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
	results := make([]*ExclusionMapping, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _RefID sql.NullString
		var _RefType sql.NullString
		var _UserID sql.NullString
		var _Date sql.NullInt64
		var _CustomerID sql.NullString
		err := rows.Scan(
			&_ID,
			&_RefID,
			&_RefType,
			&_UserID,
			&_Date,
			&_CustomerID,
		)
		if err != nil {
			return nil, err
		}
		t := &ExclusionMapping{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		if _RefType.Valid {
			t.SetRefType(_RefType.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _Date.Valid {
			t.SetDate(_Date.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindExclusionMappingsTx will find a ExclusionMapping record in the database with the provided parameters using the provided transaction
func FindExclusionMappingsTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) ([]*ExclusionMapping, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("ref_id"),
		orm.Column("ref_type"),
		orm.Column("user_id"),
		orm.Column("date"),
		orm.Column("customer_id"),
		orm.Table(ExclusionMappingTableName),
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
	results := make([]*ExclusionMapping, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _RefID sql.NullString
		var _RefType sql.NullString
		var _UserID sql.NullString
		var _Date sql.NullInt64
		var _CustomerID sql.NullString
		err := rows.Scan(
			&_ID,
			&_RefID,
			&_RefType,
			&_UserID,
			&_Date,
			&_CustomerID,
		)
		if err != nil {
			return nil, err
		}
		t := &ExclusionMapping{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		if _RefType.Valid {
			t.SetRefType(_RefType.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _Date.Valid {
			t.SetDate(_Date.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// DBFind will find a ExclusionMapping record in the database with the provided parameters
func (t *ExclusionMapping) DBFind(ctx context.Context, db *sql.DB, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("ref_id"),
		orm.Column("ref_type"),
		orm.Column("user_id"),
		orm.Column("date"),
		orm.Column("customer_id"),
		orm.Table(ExclusionMappingTableName),
	}
	if len(_params) > 0 {
		for _, param := range _params {
			params = append(params, param)
		}
	}
	q, p := orm.BuildQuery(params...)
	row := db.QueryRowContext(ctx, q, p...)
	var _ID sql.NullString
	var _RefID sql.NullString
	var _RefType sql.NullString
	var _UserID sql.NullString
	var _Date sql.NullInt64
	var _CustomerID sql.NullString
	err := row.Scan(
		&_ID,
		&_RefID,
		&_RefType,
		&_UserID,
		&_Date,
		&_CustomerID,
	)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _RefID.Valid {
		t.SetRefID(_RefID.String)
	}
	if _RefType.Valid {
		t.SetRefType(_RefType.String)
	}
	if _UserID.Valid {
		t.SetUserID(_UserID.String)
	}
	if _Date.Valid {
		t.SetDate(_Date.Int64)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	return true, nil
}

// DBFindTx will find a ExclusionMapping record in the database with the provided parameters using the provided transaction
func (t *ExclusionMapping) DBFindTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("ref_id"),
		orm.Column("ref_type"),
		orm.Column("user_id"),
		orm.Column("date"),
		orm.Column("customer_id"),
		orm.Table(ExclusionMappingTableName),
	}
	if len(_params) > 0 {
		for _, param := range _params {
			params = append(params, param)
		}
	}
	q, p := orm.BuildQuery(params...)
	row := tx.QueryRowContext(ctx, q, p...)
	var _ID sql.NullString
	var _RefID sql.NullString
	var _RefType sql.NullString
	var _UserID sql.NullString
	var _Date sql.NullInt64
	var _CustomerID sql.NullString
	err := row.Scan(
		&_ID,
		&_RefID,
		&_RefType,
		&_UserID,
		&_Date,
		&_CustomerID,
	)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _RefID.Valid {
		t.SetRefID(_RefID.String)
	}
	if _RefType.Valid {
		t.SetRefType(_RefType.String)
	}
	if _UserID.Valid {
		t.SetUserID(_UserID.String)
	}
	if _Date.Valid {
		t.SetDate(_Date.Int64)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	return true, nil
}

// CountExclusionMappings will find the count of ExclusionMapping records in the database
func CountExclusionMappings(ctx context.Context, db *sql.DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(ExclusionMappingTableName),
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

// CountExclusionMappingsTx will find the count of ExclusionMapping records in the database using the provided transaction
func CountExclusionMappingsTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(ExclusionMappingTableName),
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

// DBCount will find the count of ExclusionMapping records in the database
func (t *ExclusionMapping) DBCount(ctx context.Context, db *sql.DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(ExclusionMappingTableName),
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

// DBCountTx will find the count of ExclusionMapping records in the database using the provided transaction
func (t *ExclusionMapping) DBCountTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(ExclusionMappingTableName),
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

// DBExists will return true if the ExclusionMapping record exists in the database
func (t *ExclusionMapping) DBExists(ctx context.Context, db *sql.DB) (bool, error) {
	q := "SELECT `id` FROM `exclusion_mapping` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := db.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// DBExistsTx will return true if the ExclusionMapping record exists in the database using the provided transaction
func (t *ExclusionMapping) DBExistsTx(ctx context.Context, tx *sql.Tx) (bool, error) {
	q := "SELECT `id` FROM `exclusion_mapping` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := tx.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// PrimaryKeyColumn returns the column name for the primary key
func (t *ExclusionMapping) PrimaryKeyColumn() string {
	return ExclusionMappingColumnID
}

// PrimaryKeyColumnType returns the primary key column Go type as a string
func (t *ExclusionMapping) PrimaryKeyColumnType() string {
	return "string"
}

// PrimaryKey returns the primary key column value
func (t *ExclusionMapping) PrimaryKey() interface{} {
	return t.ID
}
