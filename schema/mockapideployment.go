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

var _ Model = (*MockapiDeployment)(nil)
var _ CSVWriter = (*MockapiDeployment)(nil)
var _ JSONWriter = (*MockapiDeployment)(nil)
var _ Checksum = (*MockapiDeployment)(nil)

// MockapiDeploymentTableName is the name of the table in SQL
const MockapiDeploymentTableName = "mockapi_deployment"

var MockapiDeploymentColumns = []string{
	"id",
	"checksum",
	"customer_id",
	"application_ext_id_ext_id",
	"application_ext_id_id",
	"ext_id",
	"name",
}

// MockapiDeployment table
type MockapiDeployment struct {
	ApplicationExtIDExtID int64   `json:"application_ext_id_ext_id"`
	ApplicationExtIDID    string  `json:"application_ext_id_id"`
	Checksum              *string `json:"checksum,omitempty"`
	CustomerID            string  `json:"customer_id"`
	ExtID                 int64   `json:"ext_id"`
	ID                    string  `json:"id"`
	Name                  string  `json:"name"`
}

// TableName returns the SQL table name for MockapiDeployment and satifies the Model interface
func (t *MockapiDeployment) TableName() string {
	return MockapiDeploymentTableName
}

// ToCSV will serialize the MockapiDeployment instance to a CSV compatible array of strings
func (t *MockapiDeployment) ToCSV() []string {
	return []string{
		t.ID,
		t.CalculateChecksum(),
		t.CustomerID,
		toCSVString(t.ApplicationExtIDExtID),
		t.ApplicationExtIDID,
		toCSVString(t.ExtID),
		t.Name,
	}
}

// WriteCSV will serialize the MockapiDeployment instance to the writer as CSV and satisfies the CSVWriter interface
func (t *MockapiDeployment) WriteCSV(w *csv.Writer) error {
	return w.Write(t.ToCSV())
}

// WriteJSON will serialize the MockapiDeployment instance to the writer as JSON and satisfies the JSONWriter interface
func (t *MockapiDeployment) WriteJSON(w io.Writer, indent ...bool) error {
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

// NewMockapiDeploymentReader creates a JSON reader which can read in MockapiDeployment objects serialized as JSON either as an array, single object or json new lines
// and writes each MockapiDeployment to the channel provided
func NewMockapiDeploymentReader(r io.Reader, ch chan<- MockapiDeployment) error {
	return orm.Deserialize(r, func(buf json.RawMessage) error {
		dec := json.NewDecoder(bytes.NewBuffer(buf))
		e := MockapiDeployment{}
		if err := dec.Decode(&e); err != nil {
			return err
		}
		ch <- e
		return nil
	})
}

// NewCSVMockapiDeploymentReaderDir will read the reader as CSV and emit each record to the channel provided
func NewCSVMockapiDeploymentReader(r io.Reader, ch chan<- MockapiDeployment) error {
	cr := csv.NewReader(r)
	for {
		record, err := cr.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		ch <- MockapiDeployment{
			ID:                    record[0],
			Checksum:              fromStringPointer(record[1]),
			CustomerID:            record[2],
			ApplicationExtIDExtID: fromCSVInt64(record[3]),
			ApplicationExtIDID:    record[4],
			ExtID:                 fromCSVInt64(record[5]),
			Name:                  record[6],
		}
	}
	return nil
}

// NewCSVMockapiDeploymentReaderFile will read the file as a CSV and emit each record to the channel provided
func NewCSVMockapiDeploymentReaderFile(fp string, ch chan<- MockapiDeployment) error {
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
	return NewCSVMockapiDeploymentReader(fc, ch)
}

// NewCSVMockapiDeploymentReaderDir will read the mockapi_deployment.csv.gz file as a CSV and emit each record to the channel provided
func NewCSVMockapiDeploymentReaderDir(dir string, ch chan<- MockapiDeployment) error {
	return NewCSVMockapiDeploymentReaderFile(filepath.Join(dir, "mockapi_deployment.csv.gz"), ch)
}

// MockapiDeploymentCSVDeduper is a function callback which takes the existing value (a) and the new value (b)
// and the return value should be the one to use (or a new one, if applicable). return nil
// to skip processing of this record
type MockapiDeploymentCSVDeduper func(a MockapiDeployment, b MockapiDeployment) *MockapiDeployment

// MockapiDeploymentCSVDedupeDisabled is set on whether the CSV writer should de-dupe values by key
var MockapiDeploymentCSVDedupeDisabled bool

// NewMockapiDeploymentCSVWriterSize creates a batch writer that will write each MockapiDeployment into a CSV file
// this method will automatically de-duplicate entries using the primary key. if the checksum
// for a newer item with the same primary key doesn't match a previously sent item, the newer
// one will be used
func NewMockapiDeploymentCSVWriterSize(w io.Writer, size int, dedupers ...MockapiDeploymentCSVDeduper) (chan MockapiDeployment, chan bool, error) {
	cw := csv.NewWriter(w)
	ch := make(chan MockapiDeployment, size)
	done := make(chan bool)
	go func() {
		defer func() { done <- true }()
		dodedupe := !MockapiDeploymentCSVDedupeDisabled
		var kv map[string]*MockapiDeployment
		var deduper MockapiDeploymentCSVDeduper
		if dedupers != nil && len(dedupers) > 0 {
			deduper = dedupers[0]
			dodedupe = true
		}
		if dodedupe {
			kv = make(map[string]*MockapiDeployment)
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

// MockapiDeploymentCSVDefaultSize is the default channel buffer size if not provided
var MockapiDeploymentCSVDefaultSize = 100

// NewMockapiDeploymentCSVWriter creates a batch writer that will write each MockapiDeployment into a CSV file
func NewMockapiDeploymentCSVWriter(w io.Writer, dedupers ...MockapiDeploymentCSVDeduper) (chan MockapiDeployment, chan bool, error) {
	return NewMockapiDeploymentCSVWriterSize(w, MockapiDeploymentCSVDefaultSize, dedupers...)
}

// NewMockapiDeploymentCSVWriterDir creates a batch writer that will write each MockapiDeployment into a CSV file named mockapi_deployment.csv.gz in dir
func NewMockapiDeploymentCSVWriterDir(dir string, dedupers ...MockapiDeploymentCSVDeduper) (chan MockapiDeployment, chan bool, error) {
	return NewMockapiDeploymentCSVWriterFile(filepath.Join(dir, "mockapi_deployment.csv.gz"), dedupers...)
}

// NewMockapiDeploymentCSVWriterFile creates a batch writer that will write each MockapiDeployment into a CSV file
func NewMockapiDeploymentCSVWriterFile(fn string, dedupers ...MockapiDeploymentCSVDeduper) (chan MockapiDeployment, chan bool, error) {
	f, err := os.Create(fn)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening CSV file %s. %v", fn, err)
	}
	var fc io.WriteCloser = f
	if filepath.Ext(fn) == ".gz" {
		w, _ := gzip.NewWriterLevel(f, gzip.BestCompression)
		fc = w
	}
	ch, done, err := NewMockapiDeploymentCSVWriter(fc, dedupers...)
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

type MockapiDeploymentDBAction func(ctx context.Context, db *sql.DB, record MockapiDeployment) error

// NewMockapiDeploymentDBWriterSize creates a DB writer that will write each issue into the DB
func NewMockapiDeploymentDBWriterSize(ctx context.Context, db *sql.DB, errors chan<- error, size int, actions ...MockapiDeploymentDBAction) (chan MockapiDeployment, chan bool, error) {
	ch := make(chan MockapiDeployment, size)
	done := make(chan bool)
	var action MockapiDeploymentDBAction
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

// NewMockapiDeploymentDBWriter creates a DB writer that will write each issue into the DB
func NewMockapiDeploymentDBWriter(ctx context.Context, db *sql.DB, errors chan<- error, actions ...MockapiDeploymentDBAction) (chan MockapiDeployment, chan bool, error) {
	return NewMockapiDeploymentDBWriterSize(ctx, db, errors, 100, actions...)
}

// MockapiDeploymentColumnID is the ID SQL column name for the MockapiDeployment table
const MockapiDeploymentColumnID = "id"

// MockapiDeploymentEscapedColumnID is the escaped ID SQL column name for the MockapiDeployment table
const MockapiDeploymentEscapedColumnID = "`id`"

// MockapiDeploymentColumnChecksum is the Checksum SQL column name for the MockapiDeployment table
const MockapiDeploymentColumnChecksum = "checksum"

// MockapiDeploymentEscapedColumnChecksum is the escaped Checksum SQL column name for the MockapiDeployment table
const MockapiDeploymentEscapedColumnChecksum = "`checksum`"

// MockapiDeploymentColumnCustomerID is the CustomerID SQL column name for the MockapiDeployment table
const MockapiDeploymentColumnCustomerID = "customer_id"

// MockapiDeploymentEscapedColumnCustomerID is the escaped CustomerID SQL column name for the MockapiDeployment table
const MockapiDeploymentEscapedColumnCustomerID = "`customer_id`"

// MockapiDeploymentColumnApplicationExtIDExtID is the ApplicationExtIDExtID SQL column name for the MockapiDeployment table
const MockapiDeploymentColumnApplicationExtIDExtID = "application_ext_id_ext_id"

// MockapiDeploymentEscapedColumnApplicationExtIDExtID is the escaped ApplicationExtIDExtID SQL column name for the MockapiDeployment table
const MockapiDeploymentEscapedColumnApplicationExtIDExtID = "`application_ext_id_ext_id`"

// MockapiDeploymentColumnApplicationExtIDID is the ApplicationExtIDID SQL column name for the MockapiDeployment table
const MockapiDeploymentColumnApplicationExtIDID = "application_ext_id_id"

// MockapiDeploymentEscapedColumnApplicationExtIDID is the escaped ApplicationExtIDID SQL column name for the MockapiDeployment table
const MockapiDeploymentEscapedColumnApplicationExtIDID = "`application_ext_id_id`"

// MockapiDeploymentColumnExtID is the ExtID SQL column name for the MockapiDeployment table
const MockapiDeploymentColumnExtID = "ext_id"

// MockapiDeploymentEscapedColumnExtID is the escaped ExtID SQL column name for the MockapiDeployment table
const MockapiDeploymentEscapedColumnExtID = "`ext_id`"

// MockapiDeploymentColumnName is the Name SQL column name for the MockapiDeployment table
const MockapiDeploymentColumnName = "name"

// MockapiDeploymentEscapedColumnName is the escaped Name SQL column name for the MockapiDeployment table
const MockapiDeploymentEscapedColumnName = "`name`"

// GetID will return the MockapiDeployment ID value
func (t *MockapiDeployment) GetID() string {
	return t.ID
}

// SetID will set the MockapiDeployment ID value
func (t *MockapiDeployment) SetID(v string) {
	t.ID = v
}

// FindMockapiDeploymentByID will find a MockapiDeployment by ID
func FindMockapiDeploymentByID(ctx context.Context, db *sql.DB, value string) (*MockapiDeployment, error) {
	q := "SELECT `mockapi_deployment`.`id`,`mockapi_deployment`.`checksum`,`mockapi_deployment`.`customer_id`,`mockapi_deployment`.`application_ext_id_ext_id`,`mockapi_deployment`.`application_ext_id_id`,`mockapi_deployment`.`ext_id`,`mockapi_deployment`.`name` FROM `mockapi_deployment` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CustomerID sql.NullString
	var _ApplicationExtIDExtID sql.NullInt64
	var _ApplicationExtIDID sql.NullString
	var _ExtID sql.NullInt64
	var _Name sql.NullString
	err := db.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_ApplicationExtIDExtID,
		&_ApplicationExtIDID,
		&_ExtID,
		&_Name,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &MockapiDeployment{}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _Checksum.Valid {
		t.SetChecksum(_Checksum.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _ApplicationExtIDExtID.Valid {
		t.SetApplicationExtIDExtID(_ApplicationExtIDExtID.Int64)
	}
	if _ApplicationExtIDID.Valid {
		t.SetApplicationExtIDID(_ApplicationExtIDID.String)
	}
	if _ExtID.Valid {
		t.SetExtID(_ExtID.Int64)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	return t, nil
}

// FindMockapiDeploymentByIDTx will find a MockapiDeployment by ID using the provided transaction
func FindMockapiDeploymentByIDTx(ctx context.Context, tx *sql.Tx, value string) (*MockapiDeployment, error) {
	q := "SELECT `mockapi_deployment`.`id`,`mockapi_deployment`.`checksum`,`mockapi_deployment`.`customer_id`,`mockapi_deployment`.`application_ext_id_ext_id`,`mockapi_deployment`.`application_ext_id_id`,`mockapi_deployment`.`ext_id`,`mockapi_deployment`.`name` FROM `mockapi_deployment` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CustomerID sql.NullString
	var _ApplicationExtIDExtID sql.NullInt64
	var _ApplicationExtIDID sql.NullString
	var _ExtID sql.NullInt64
	var _Name sql.NullString
	err := tx.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_ApplicationExtIDExtID,
		&_ApplicationExtIDID,
		&_ExtID,
		&_Name,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &MockapiDeployment{}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _Checksum.Valid {
		t.SetChecksum(_Checksum.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _ApplicationExtIDExtID.Valid {
		t.SetApplicationExtIDExtID(_ApplicationExtIDExtID.Int64)
	}
	if _ApplicationExtIDID.Valid {
		t.SetApplicationExtIDID(_ApplicationExtIDID.String)
	}
	if _ExtID.Valid {
		t.SetExtID(_ExtID.Int64)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	return t, nil
}

// GetChecksum will return the MockapiDeployment Checksum value
func (t *MockapiDeployment) GetChecksum() string {
	if t.Checksum == nil {
		return ""
	}
	return *t.Checksum
}

// SetChecksum will set the MockapiDeployment Checksum value
func (t *MockapiDeployment) SetChecksum(v string) {
	t.Checksum = &v
}

// GetCustomerID will return the MockapiDeployment CustomerID value
func (t *MockapiDeployment) GetCustomerID() string {
	return t.CustomerID
}

// SetCustomerID will set the MockapiDeployment CustomerID value
func (t *MockapiDeployment) SetCustomerID(v string) {
	t.CustomerID = v
}

// FindMockapiDeploymentsByCustomerID will find all MockapiDeployments by the CustomerID value
func FindMockapiDeploymentsByCustomerID(ctx context.Context, db *sql.DB, value string) ([]*MockapiDeployment, error) {
	q := "SELECT `mockapi_deployment`.`id`,`mockapi_deployment`.`checksum`,`mockapi_deployment`.`customer_id`,`mockapi_deployment`.`application_ext_id_ext_id`,`mockapi_deployment`.`application_ext_id_id`,`mockapi_deployment`.`ext_id`,`mockapi_deployment`.`name` FROM `mockapi_deployment` WHERE `customer_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*MockapiDeployment, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _ApplicationExtIDExtID sql.NullInt64
		var _ApplicationExtIDID sql.NullString
		var _ExtID sql.NullInt64
		var _Name sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_ApplicationExtIDExtID,
			&_ApplicationExtIDID,
			&_ExtID,
			&_Name,
		)
		if err != nil {
			return nil, err
		}
		t := &MockapiDeployment{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _ApplicationExtIDExtID.Valid {
			t.SetApplicationExtIDExtID(_ApplicationExtIDExtID.Int64)
		}
		if _ApplicationExtIDID.Valid {
			t.SetApplicationExtIDID(_ApplicationExtIDID.String)
		}
		if _ExtID.Valid {
			t.SetExtID(_ExtID.Int64)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindMockapiDeploymentsByCustomerIDTx will find all MockapiDeployments by the CustomerID value using the provided transaction
func FindMockapiDeploymentsByCustomerIDTx(ctx context.Context, tx *sql.Tx, value string) ([]*MockapiDeployment, error) {
	q := "SELECT `mockapi_deployment`.`id`,`mockapi_deployment`.`checksum`,`mockapi_deployment`.`customer_id`,`mockapi_deployment`.`application_ext_id_ext_id`,`mockapi_deployment`.`application_ext_id_id`,`mockapi_deployment`.`ext_id`,`mockapi_deployment`.`name` FROM `mockapi_deployment` WHERE `customer_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*MockapiDeployment, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _ApplicationExtIDExtID sql.NullInt64
		var _ApplicationExtIDID sql.NullString
		var _ExtID sql.NullInt64
		var _Name sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_ApplicationExtIDExtID,
			&_ApplicationExtIDID,
			&_ExtID,
			&_Name,
		)
		if err != nil {
			return nil, err
		}
		t := &MockapiDeployment{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _ApplicationExtIDExtID.Valid {
			t.SetApplicationExtIDExtID(_ApplicationExtIDExtID.Int64)
		}
		if _ApplicationExtIDID.Valid {
			t.SetApplicationExtIDID(_ApplicationExtIDID.String)
		}
		if _ExtID.Valid {
			t.SetExtID(_ExtID.Int64)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetApplicationExtIDExtID will return the MockapiDeployment ApplicationExtIDExtID value
func (t *MockapiDeployment) GetApplicationExtIDExtID() int64 {
	return t.ApplicationExtIDExtID
}

// SetApplicationExtIDExtID will set the MockapiDeployment ApplicationExtIDExtID value
func (t *MockapiDeployment) SetApplicationExtIDExtID(v int64) {
	t.ApplicationExtIDExtID = v
}

// FindMockapiDeploymentsByApplicationExtIDExtID will find all MockapiDeployments by the ApplicationExtIDExtID value
func FindMockapiDeploymentsByApplicationExtIDExtID(ctx context.Context, db *sql.DB, value int64) ([]*MockapiDeployment, error) {
	q := "SELECT `mockapi_deployment`.`id`,`mockapi_deployment`.`checksum`,`mockapi_deployment`.`customer_id`,`mockapi_deployment`.`application_ext_id_ext_id`,`mockapi_deployment`.`application_ext_id_id`,`mockapi_deployment`.`ext_id`,`mockapi_deployment`.`name` FROM `mockapi_deployment` WHERE `application_ext_id_ext_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*MockapiDeployment, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _ApplicationExtIDExtID sql.NullInt64
		var _ApplicationExtIDID sql.NullString
		var _ExtID sql.NullInt64
		var _Name sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_ApplicationExtIDExtID,
			&_ApplicationExtIDID,
			&_ExtID,
			&_Name,
		)
		if err != nil {
			return nil, err
		}
		t := &MockapiDeployment{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _ApplicationExtIDExtID.Valid {
			t.SetApplicationExtIDExtID(_ApplicationExtIDExtID.Int64)
		}
		if _ApplicationExtIDID.Valid {
			t.SetApplicationExtIDID(_ApplicationExtIDID.String)
		}
		if _ExtID.Valid {
			t.SetExtID(_ExtID.Int64)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindMockapiDeploymentsByApplicationExtIDExtIDTx will find all MockapiDeployments by the ApplicationExtIDExtID value using the provided transaction
func FindMockapiDeploymentsByApplicationExtIDExtIDTx(ctx context.Context, tx *sql.Tx, value int64) ([]*MockapiDeployment, error) {
	q := "SELECT `mockapi_deployment`.`id`,`mockapi_deployment`.`checksum`,`mockapi_deployment`.`customer_id`,`mockapi_deployment`.`application_ext_id_ext_id`,`mockapi_deployment`.`application_ext_id_id`,`mockapi_deployment`.`ext_id`,`mockapi_deployment`.`name` FROM `mockapi_deployment` WHERE `application_ext_id_ext_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*MockapiDeployment, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _ApplicationExtIDExtID sql.NullInt64
		var _ApplicationExtIDID sql.NullString
		var _ExtID sql.NullInt64
		var _Name sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_ApplicationExtIDExtID,
			&_ApplicationExtIDID,
			&_ExtID,
			&_Name,
		)
		if err != nil {
			return nil, err
		}
		t := &MockapiDeployment{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _ApplicationExtIDExtID.Valid {
			t.SetApplicationExtIDExtID(_ApplicationExtIDExtID.Int64)
		}
		if _ApplicationExtIDID.Valid {
			t.SetApplicationExtIDID(_ApplicationExtIDID.String)
		}
		if _ExtID.Valid {
			t.SetExtID(_ExtID.Int64)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetApplicationExtIDID will return the MockapiDeployment ApplicationExtIDID value
func (t *MockapiDeployment) GetApplicationExtIDID() string {
	return t.ApplicationExtIDID
}

// SetApplicationExtIDID will set the MockapiDeployment ApplicationExtIDID value
func (t *MockapiDeployment) SetApplicationExtIDID(v string) {
	t.ApplicationExtIDID = v
}

// FindMockapiDeploymentsByApplicationExtIDID will find all MockapiDeployments by the ApplicationExtIDID value
func FindMockapiDeploymentsByApplicationExtIDID(ctx context.Context, db *sql.DB, value string) ([]*MockapiDeployment, error) {
	q := "SELECT `mockapi_deployment`.`id`,`mockapi_deployment`.`checksum`,`mockapi_deployment`.`customer_id`,`mockapi_deployment`.`application_ext_id_ext_id`,`mockapi_deployment`.`application_ext_id_id`,`mockapi_deployment`.`ext_id`,`mockapi_deployment`.`name` FROM `mockapi_deployment` WHERE `application_ext_id_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*MockapiDeployment, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _ApplicationExtIDExtID sql.NullInt64
		var _ApplicationExtIDID sql.NullString
		var _ExtID sql.NullInt64
		var _Name sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_ApplicationExtIDExtID,
			&_ApplicationExtIDID,
			&_ExtID,
			&_Name,
		)
		if err != nil {
			return nil, err
		}
		t := &MockapiDeployment{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _ApplicationExtIDExtID.Valid {
			t.SetApplicationExtIDExtID(_ApplicationExtIDExtID.Int64)
		}
		if _ApplicationExtIDID.Valid {
			t.SetApplicationExtIDID(_ApplicationExtIDID.String)
		}
		if _ExtID.Valid {
			t.SetExtID(_ExtID.Int64)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindMockapiDeploymentsByApplicationExtIDIDTx will find all MockapiDeployments by the ApplicationExtIDID value using the provided transaction
func FindMockapiDeploymentsByApplicationExtIDIDTx(ctx context.Context, tx *sql.Tx, value string) ([]*MockapiDeployment, error) {
	q := "SELECT `mockapi_deployment`.`id`,`mockapi_deployment`.`checksum`,`mockapi_deployment`.`customer_id`,`mockapi_deployment`.`application_ext_id_ext_id`,`mockapi_deployment`.`application_ext_id_id`,`mockapi_deployment`.`ext_id`,`mockapi_deployment`.`name` FROM `mockapi_deployment` WHERE `application_ext_id_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*MockapiDeployment, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _ApplicationExtIDExtID sql.NullInt64
		var _ApplicationExtIDID sql.NullString
		var _ExtID sql.NullInt64
		var _Name sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_ApplicationExtIDExtID,
			&_ApplicationExtIDID,
			&_ExtID,
			&_Name,
		)
		if err != nil {
			return nil, err
		}
		t := &MockapiDeployment{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _ApplicationExtIDExtID.Valid {
			t.SetApplicationExtIDExtID(_ApplicationExtIDExtID.Int64)
		}
		if _ApplicationExtIDID.Valid {
			t.SetApplicationExtIDID(_ApplicationExtIDID.String)
		}
		if _ExtID.Valid {
			t.SetExtID(_ExtID.Int64)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetExtID will return the MockapiDeployment ExtID value
func (t *MockapiDeployment) GetExtID() int64 {
	return t.ExtID
}

// SetExtID will set the MockapiDeployment ExtID value
func (t *MockapiDeployment) SetExtID(v int64) {
	t.ExtID = v
}

// GetName will return the MockapiDeployment Name value
func (t *MockapiDeployment) GetName() string {
	return t.Name
}

// SetName will set the MockapiDeployment Name value
func (t *MockapiDeployment) SetName(v string) {
	t.Name = v
}

func (t *MockapiDeployment) toTimestamp(value time.Time) *timestamp.Timestamp {
	ts, _ := ptypes.TimestampProto(value)
	return ts
}

// DBCreateMockapiDeploymentTable will create the MockapiDeployment table
func DBCreateMockapiDeploymentTable(ctx context.Context, db *sql.DB) error {
	q := "CREATE TABLE `mockapi_deployment` (`id` VARCHAR(64) NOT NULL PRIMARY KEY,`checksum` CHAR(64),`customer_id` VARCHAR(64) NOT NULL,`application_ext_id_ext_id` BIGINT NOT NULL,`application_ext_id_id`VARCHAR(64) NOT NULL,`ext_id`BIGINT NOT NULL,`name` TEXT NOT NULL,INDEX mockapi_deployment_customer_id_index (`customer_id`),INDEX mockapi_deployment_application_ext_id_ext_id_index (`application_ext_id_ext_id`),INDEX mockapi_deployment_application_ext_id_id_index (`application_ext_id_id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBCreateMockapiDeploymentTableTx will create the MockapiDeployment table using the provided transction
func DBCreateMockapiDeploymentTableTx(ctx context.Context, tx *sql.Tx) error {
	q := "CREATE TABLE `mockapi_deployment` (`id` VARCHAR(64) NOT NULL PRIMARY KEY,`checksum` CHAR(64),`customer_id` VARCHAR(64) NOT NULL,`application_ext_id_ext_id` BIGINT NOT NULL,`application_ext_id_id`VARCHAR(64) NOT NULL,`ext_id`BIGINT NOT NULL,`name` TEXT NOT NULL,INDEX mockapi_deployment_customer_id_index (`customer_id`),INDEX mockapi_deployment_application_ext_id_ext_id_index (`application_ext_id_ext_id`),INDEX mockapi_deployment_application_ext_id_id_index (`application_ext_id_id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// DBDropMockapiDeploymentTable will drop the MockapiDeployment table
func DBDropMockapiDeploymentTable(ctx context.Context, db *sql.DB) error {
	q := "DROP TABLE IF EXISTS `mockapi_deployment`"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBDropMockapiDeploymentTableTx will drop the MockapiDeployment table using the provided transaction
func DBDropMockapiDeploymentTableTx(ctx context.Context, tx *sql.Tx) error {
	q := "DROP TABLE IF EXISTS `mockapi_deployment`"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// CalculateChecksum will calculate a checksum of the SHA1 of all field values
func (t *MockapiDeployment) CalculateChecksum() string {
	return orm.HashStrings(
		orm.ToString(t.ID),
		orm.ToString(t.CustomerID),
		orm.ToString(t.ApplicationExtIDExtID),
		orm.ToString(t.ApplicationExtIDID),
		orm.ToString(t.ExtID),
		orm.ToString(t.Name),
	)
}

// DBCreate will create a new MockapiDeployment record in the database
func (t *MockapiDeployment) DBCreate(ctx context.Context, db *sql.DB) (sql.Result, error) {
	q := "INSERT INTO `mockapi_deployment` (`mockapi_deployment`.`id`,`mockapi_deployment`.`checksum`,`mockapi_deployment`.`customer_id`,`mockapi_deployment`.`application_ext_id_ext_id`,`mockapi_deployment`.`application_ext_id_id`,`mockapi_deployment`.`ext_id`,`mockapi_deployment`.`name`) VALUES (?,?,?,?,?,?,?)"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLInt64(t.ApplicationExtIDExtID),
		orm.ToSQLString(t.ApplicationExtIDID),
		orm.ToSQLInt64(t.ExtID),
		orm.ToSQLString(t.Name),
	)
}

// DBCreateTx will create a new MockapiDeployment record in the database using the provided transaction
func (t *MockapiDeployment) DBCreateTx(ctx context.Context, tx *sql.Tx) (sql.Result, error) {
	q := "INSERT INTO `mockapi_deployment` (`mockapi_deployment`.`id`,`mockapi_deployment`.`checksum`,`mockapi_deployment`.`customer_id`,`mockapi_deployment`.`application_ext_id_ext_id`,`mockapi_deployment`.`application_ext_id_id`,`mockapi_deployment`.`ext_id`,`mockapi_deployment`.`name`) VALUES (?,?,?,?,?,?,?)"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLInt64(t.ApplicationExtIDExtID),
		orm.ToSQLString(t.ApplicationExtIDID),
		orm.ToSQLInt64(t.ExtID),
		orm.ToSQLString(t.Name),
	)
}

// DBCreateIgnoreDuplicate will upsert the MockapiDeployment record in the database
func (t *MockapiDeployment) DBCreateIgnoreDuplicate(ctx context.Context, db *sql.DB) (sql.Result, error) {
	q := "INSERT INTO `mockapi_deployment` (`mockapi_deployment`.`id`,`mockapi_deployment`.`checksum`,`mockapi_deployment`.`customer_id`,`mockapi_deployment`.`application_ext_id_ext_id`,`mockapi_deployment`.`application_ext_id_id`,`mockapi_deployment`.`ext_id`,`mockapi_deployment`.`name`) VALUES (?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLInt64(t.ApplicationExtIDExtID),
		orm.ToSQLString(t.ApplicationExtIDID),
		orm.ToSQLInt64(t.ExtID),
		orm.ToSQLString(t.Name),
	)
}

// DBCreateIgnoreDuplicateTx will upsert the MockapiDeployment record in the database using the provided transaction
func (t *MockapiDeployment) DBCreateIgnoreDuplicateTx(ctx context.Context, tx *sql.Tx) (sql.Result, error) {
	q := "INSERT INTO `mockapi_deployment` (`mockapi_deployment`.`id`,`mockapi_deployment`.`checksum`,`mockapi_deployment`.`customer_id`,`mockapi_deployment`.`application_ext_id_ext_id`,`mockapi_deployment`.`application_ext_id_id`,`mockapi_deployment`.`ext_id`,`mockapi_deployment`.`name`) VALUES (?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLInt64(t.ApplicationExtIDExtID),
		orm.ToSQLString(t.ApplicationExtIDID),
		orm.ToSQLInt64(t.ExtID),
		orm.ToSQLString(t.Name),
	)
}

// DeleteAllMockapiDeployments deletes all MockapiDeployment records in the database with optional filters
func DeleteAllMockapiDeployments(ctx context.Context, db *sql.DB, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(MockapiDeploymentTableName),
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

// DeleteAllMockapiDeploymentsTx deletes all MockapiDeployment records in the database with optional filters using the provided transaction
func DeleteAllMockapiDeploymentsTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(MockapiDeploymentTableName),
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

// DBDelete will delete this MockapiDeployment record in the database
func (t *MockapiDeployment) DBDelete(ctx context.Context, db *sql.DB) (bool, error) {
	q := "DELETE FROM `mockapi_deployment` WHERE `id` = ?"
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

// DBDeleteTx will delete this MockapiDeployment record in the database using the provided transaction
func (t *MockapiDeployment) DBDeleteTx(ctx context.Context, tx *sql.Tx) (bool, error) {
	q := "DELETE FROM `mockapi_deployment` WHERE `id` = ?"
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

// DBUpdate will update the MockapiDeployment record in the database
func (t *MockapiDeployment) DBUpdate(ctx context.Context, db *sql.DB) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `mockapi_deployment` SET `checksum`=?,`customer_id`=?,`application_ext_id_ext_id`=?,`application_ext_id_id`=?,`ext_id`=?,`name`=? WHERE `id`=?"
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLInt64(t.ApplicationExtIDExtID),
		orm.ToSQLString(t.ApplicationExtIDID),
		orm.ToSQLInt64(t.ExtID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.ID),
	)
}

// DBUpdateTx will update the MockapiDeployment record in the database using the provided transaction
func (t *MockapiDeployment) DBUpdateTx(ctx context.Context, tx *sql.Tx) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `mockapi_deployment` SET `checksum`=?,`customer_id`=?,`application_ext_id_ext_id`=?,`application_ext_id_id`=?,`ext_id`=?,`name`=? WHERE `id`=?"
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLInt64(t.ApplicationExtIDExtID),
		orm.ToSQLString(t.ApplicationExtIDID),
		orm.ToSQLInt64(t.ExtID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.ID),
	)
}

// DBUpsert will upsert the MockapiDeployment record in the database
func (t *MockapiDeployment) DBUpsert(ctx context.Context, db *sql.DB, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `mockapi_deployment` (`mockapi_deployment`.`id`,`mockapi_deployment`.`checksum`,`mockapi_deployment`.`customer_id`,`mockapi_deployment`.`application_ext_id_ext_id`,`mockapi_deployment`.`application_ext_id_id`,`mockapi_deployment`.`ext_id`,`mockapi_deployment`.`name`) VALUES (?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `mockapi_deployment` (`mockapi_deployment`.`id`,`mockapi_deployment`.`checksum`,`mockapi_deployment`.`customer_id`,`mockapi_deployment`.`application_ext_id_ext_id`,`mockapi_deployment`.`application_ext_id_id`,`mockapi_deployment`.`ext_id`,`mockapi_deployment`.`name`) VALUES (?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`customer_id`=VALUES(`customer_id`),`application_ext_id_ext_id`=VALUES(`application_ext_id_ext_id`),`application_ext_id_id`=VALUES(`application_ext_id_id`),`ext_id`=VALUES(`ext_id`),`name`=VALUES(`name`)"
	}
	r, err := db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLInt64(t.ApplicationExtIDExtID),
		orm.ToSQLString(t.ApplicationExtIDID),
		orm.ToSQLInt64(t.ExtID),
		orm.ToSQLString(t.Name),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBUpsertTx will upsert the MockapiDeployment record in the database using the provided transaction
func (t *MockapiDeployment) DBUpsertTx(ctx context.Context, tx *sql.Tx, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `mockapi_deployment` (`mockapi_deployment`.`id`,`mockapi_deployment`.`checksum`,`mockapi_deployment`.`customer_id`,`mockapi_deployment`.`application_ext_id_ext_id`,`mockapi_deployment`.`application_ext_id_id`,`mockapi_deployment`.`ext_id`,`mockapi_deployment`.`name`) VALUES (?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `mockapi_deployment` (`mockapi_deployment`.`id`,`mockapi_deployment`.`checksum`,`mockapi_deployment`.`customer_id`,`mockapi_deployment`.`application_ext_id_ext_id`,`mockapi_deployment`.`application_ext_id_id`,`mockapi_deployment`.`ext_id`,`mockapi_deployment`.`name`) VALUES (?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`customer_id`=VALUES(`customer_id`),`application_ext_id_ext_id`=VALUES(`application_ext_id_ext_id`),`application_ext_id_id`=VALUES(`application_ext_id_id`),`ext_id`=VALUES(`ext_id`),`name`=VALUES(`name`)"
	}
	r, err := tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLInt64(t.ApplicationExtIDExtID),
		orm.ToSQLString(t.ApplicationExtIDID),
		orm.ToSQLInt64(t.ExtID),
		orm.ToSQLString(t.Name),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBFindOne will find a MockapiDeployment record in the database with the primary key
func (t *MockapiDeployment) DBFindOne(ctx context.Context, db *sql.DB, value string) (bool, error) {
	q := "SELECT `mockapi_deployment`.`id`,`mockapi_deployment`.`checksum`,`mockapi_deployment`.`customer_id`,`mockapi_deployment`.`application_ext_id_ext_id`,`mockapi_deployment`.`application_ext_id_id`,`mockapi_deployment`.`ext_id`,`mockapi_deployment`.`name` FROM `mockapi_deployment` WHERE `id` = ? LIMIT 1"
	row := db.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CustomerID sql.NullString
	var _ApplicationExtIDExtID sql.NullInt64
	var _ApplicationExtIDID sql.NullString
	var _ExtID sql.NullInt64
	var _Name sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_ApplicationExtIDExtID,
		&_ApplicationExtIDID,
		&_ExtID,
		&_Name,
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
	if _ApplicationExtIDExtID.Valid {
		t.SetApplicationExtIDExtID(_ApplicationExtIDExtID.Int64)
	}
	if _ApplicationExtIDID.Valid {
		t.SetApplicationExtIDID(_ApplicationExtIDID.String)
	}
	if _ExtID.Valid {
		t.SetExtID(_ExtID.Int64)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	return true, nil
}

// DBFindOneTx will find a MockapiDeployment record in the database with the primary key using the provided transaction
func (t *MockapiDeployment) DBFindOneTx(ctx context.Context, tx *sql.Tx, value string) (bool, error) {
	q := "SELECT `mockapi_deployment`.`id`,`mockapi_deployment`.`checksum`,`mockapi_deployment`.`customer_id`,`mockapi_deployment`.`application_ext_id_ext_id`,`mockapi_deployment`.`application_ext_id_id`,`mockapi_deployment`.`ext_id`,`mockapi_deployment`.`name` FROM `mockapi_deployment` WHERE `id` = ? LIMIT 1"
	row := tx.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CustomerID sql.NullString
	var _ApplicationExtIDExtID sql.NullInt64
	var _ApplicationExtIDID sql.NullString
	var _ExtID sql.NullInt64
	var _Name sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_ApplicationExtIDExtID,
		&_ApplicationExtIDID,
		&_ExtID,
		&_Name,
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
	if _ApplicationExtIDExtID.Valid {
		t.SetApplicationExtIDExtID(_ApplicationExtIDExtID.Int64)
	}
	if _ApplicationExtIDID.Valid {
		t.SetApplicationExtIDID(_ApplicationExtIDID.String)
	}
	if _ExtID.Valid {
		t.SetExtID(_ExtID.Int64)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	return true, nil
}

// FindMockapiDeployments will find a MockapiDeployment record in the database with the provided parameters
func FindMockapiDeployments(ctx context.Context, db *sql.DB, _params ...interface{}) ([]*MockapiDeployment, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("customer_id"),
		orm.Column("application_ext_id_ext_id"),
		orm.Column("application_ext_id_id"),
		orm.Column("ext_id"),
		orm.Column("name"),
		orm.Table(MockapiDeploymentTableName),
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
	results := make([]*MockapiDeployment, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _ApplicationExtIDExtID sql.NullInt64
		var _ApplicationExtIDID sql.NullString
		var _ExtID sql.NullInt64
		var _Name sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_ApplicationExtIDExtID,
			&_ApplicationExtIDID,
			&_ExtID,
			&_Name,
		)
		if err != nil {
			return nil, err
		}
		t := &MockapiDeployment{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _ApplicationExtIDExtID.Valid {
			t.SetApplicationExtIDExtID(_ApplicationExtIDExtID.Int64)
		}
		if _ApplicationExtIDID.Valid {
			t.SetApplicationExtIDID(_ApplicationExtIDID.String)
		}
		if _ExtID.Valid {
			t.SetExtID(_ExtID.Int64)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindMockapiDeploymentsTx will find a MockapiDeployment record in the database with the provided parameters using the provided transaction
func FindMockapiDeploymentsTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) ([]*MockapiDeployment, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("customer_id"),
		orm.Column("application_ext_id_ext_id"),
		orm.Column("application_ext_id_id"),
		orm.Column("ext_id"),
		orm.Column("name"),
		orm.Table(MockapiDeploymentTableName),
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
	results := make([]*MockapiDeployment, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _ApplicationExtIDExtID sql.NullInt64
		var _ApplicationExtIDID sql.NullString
		var _ExtID sql.NullInt64
		var _Name sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_ApplicationExtIDExtID,
			&_ApplicationExtIDID,
			&_ExtID,
			&_Name,
		)
		if err != nil {
			return nil, err
		}
		t := &MockapiDeployment{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _ApplicationExtIDExtID.Valid {
			t.SetApplicationExtIDExtID(_ApplicationExtIDExtID.Int64)
		}
		if _ApplicationExtIDID.Valid {
			t.SetApplicationExtIDID(_ApplicationExtIDID.String)
		}
		if _ExtID.Valid {
			t.SetExtID(_ExtID.Int64)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// DBFind will find a MockapiDeployment record in the database with the provided parameters
func (t *MockapiDeployment) DBFind(ctx context.Context, db *sql.DB, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("customer_id"),
		orm.Column("application_ext_id_ext_id"),
		orm.Column("application_ext_id_id"),
		orm.Column("ext_id"),
		orm.Column("name"),
		orm.Table(MockapiDeploymentTableName),
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
	var _ApplicationExtIDExtID sql.NullInt64
	var _ApplicationExtIDID sql.NullString
	var _ExtID sql.NullInt64
	var _Name sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_ApplicationExtIDExtID,
		&_ApplicationExtIDID,
		&_ExtID,
		&_Name,
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
	if _ApplicationExtIDExtID.Valid {
		t.SetApplicationExtIDExtID(_ApplicationExtIDExtID.Int64)
	}
	if _ApplicationExtIDID.Valid {
		t.SetApplicationExtIDID(_ApplicationExtIDID.String)
	}
	if _ExtID.Valid {
		t.SetExtID(_ExtID.Int64)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	return true, nil
}

// DBFindTx will find a MockapiDeployment record in the database with the provided parameters using the provided transaction
func (t *MockapiDeployment) DBFindTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("customer_id"),
		orm.Column("application_ext_id_ext_id"),
		orm.Column("application_ext_id_id"),
		orm.Column("ext_id"),
		orm.Column("name"),
		orm.Table(MockapiDeploymentTableName),
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
	var _ApplicationExtIDExtID sql.NullInt64
	var _ApplicationExtIDID sql.NullString
	var _ExtID sql.NullInt64
	var _Name sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_ApplicationExtIDExtID,
		&_ApplicationExtIDID,
		&_ExtID,
		&_Name,
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
	if _ApplicationExtIDExtID.Valid {
		t.SetApplicationExtIDExtID(_ApplicationExtIDExtID.Int64)
	}
	if _ApplicationExtIDID.Valid {
		t.SetApplicationExtIDID(_ApplicationExtIDID.String)
	}
	if _ExtID.Valid {
		t.SetExtID(_ExtID.Int64)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	return true, nil
}

// CountMockapiDeployments will find the count of MockapiDeployment records in the database
func CountMockapiDeployments(ctx context.Context, db *sql.DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(MockapiDeploymentTableName),
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

// CountMockapiDeploymentsTx will find the count of MockapiDeployment records in the database using the provided transaction
func CountMockapiDeploymentsTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(MockapiDeploymentTableName),
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

// DBCount will find the count of MockapiDeployment records in the database
func (t *MockapiDeployment) DBCount(ctx context.Context, db *sql.DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(MockapiDeploymentTableName),
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

// DBCountTx will find the count of MockapiDeployment records in the database using the provided transaction
func (t *MockapiDeployment) DBCountTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(MockapiDeploymentTableName),
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

// DBExists will return true if the MockapiDeployment record exists in the database
func (t *MockapiDeployment) DBExists(ctx context.Context, db *sql.DB) (bool, error) {
	q := "SELECT `id` FROM `mockapi_deployment` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := db.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// DBExistsTx will return true if the MockapiDeployment record exists in the database using the provided transaction
func (t *MockapiDeployment) DBExistsTx(ctx context.Context, tx *sql.Tx) (bool, error) {
	q := "SELECT `id` FROM `mockapi_deployment` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := tx.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// PrimaryKeyColumn returns the column name for the primary key
func (t *MockapiDeployment) PrimaryKeyColumn() string {
	return MockapiDeploymentColumnID
}

// PrimaryKeyColumnType returns the primary key column Go type as a string
func (t *MockapiDeployment) PrimaryKeyColumnType() string {
	return "string"
}

// PrimaryKey returns the primary key column value
func (t *MockapiDeployment) PrimaryKey() interface{} {
	return t.ID
}
