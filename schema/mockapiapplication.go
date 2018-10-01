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

var _ Model = (*MockapiApplication)(nil)
var _ CSVWriter = (*MockapiApplication)(nil)
var _ JSONWriter = (*MockapiApplication)(nil)
var _ Checksum = (*MockapiApplication)(nil)

// MockapiApplicationTableName is the name of the table in SQL
const MockapiApplicationTableName = "mockapi_application"

var MockapiApplicationColumns = []string{
	"id",
	"checksum",
	"customer_id",
	"ext_id",
	"name",
}

// MockapiApplication table
type MockapiApplication struct {
	Checksum   *string `json:"checksum,omitempty"`
	CustomerID string  `json:"customer_id"`
	ExtID      int64   `json:"ext_id"`
	ID         string  `json:"id"`
	Name       string  `json:"name"`
}

// TableName returns the SQL table name for MockapiApplication and satifies the Model interface
func (t *MockapiApplication) TableName() string {
	return MockapiApplicationTableName
}

// ToCSV will serialize the MockapiApplication instance to a CSV compatible array of strings
func (t *MockapiApplication) ToCSV() []string {
	return []string{
		t.ID,
		t.CalculateChecksum(),
		t.CustomerID,
		toCSVString(t.ExtID),
		t.Name,
	}
}

// WriteCSV will serialize the MockapiApplication instance to the writer as CSV and satisfies the CSVWriter interface
func (t *MockapiApplication) WriteCSV(w *csv.Writer) error {
	return w.Write(t.ToCSV())
}

// WriteJSON will serialize the MockapiApplication instance to the writer as JSON and satisfies the JSONWriter interface
func (t *MockapiApplication) WriteJSON(w io.Writer, indent ...bool) error {
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

// NewMockapiApplicationReader creates a JSON reader which can read in MockapiApplication objects serialized as JSON either as an array, single object or json new lines
// and writes each MockapiApplication to the channel provided
func NewMockapiApplicationReader(r io.Reader, ch chan<- MockapiApplication) error {
	return orm.Deserialize(r, func(buf json.RawMessage) error {
		dec := json.NewDecoder(bytes.NewBuffer(buf))
		e := MockapiApplication{}
		if err := dec.Decode(&e); err != nil {
			return err
		}
		ch <- e
		return nil
	})
}

// NewCSVMockapiApplicationReaderDir will read the reader as CSV and emit each record to the channel provided
func NewCSVMockapiApplicationReader(r io.Reader, ch chan<- MockapiApplication) error {
	cr := csv.NewReader(r)
	for {
		record, err := cr.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		ch <- MockapiApplication{
			ID:         record[0],
			Checksum:   fromStringPointer(record[1]),
			CustomerID: record[2],
			ExtID:      fromCSVInt64(record[3]),
			Name:       record[4],
		}
	}
	return nil
}

// NewCSVMockapiApplicationReaderFile will read the file as a CSV and emit each record to the channel provided
func NewCSVMockapiApplicationReaderFile(fp string, ch chan<- MockapiApplication) error {
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
	return NewCSVMockapiApplicationReader(fc, ch)
}

// NewCSVMockapiApplicationReaderDir will read the mockapi_application.csv.gz file as a CSV and emit each record to the channel provided
func NewCSVMockapiApplicationReaderDir(dir string, ch chan<- MockapiApplication) error {
	return NewCSVMockapiApplicationReaderFile(filepath.Join(dir, "mockapi_application.csv.gz"), ch)
}

// MockapiApplicationCSVDeduper is a function callback which takes the existing value (a) and the new value (b)
// and the return value should be the one to use (or a new one, if applicable). return nil
// to skip processing of this record
type MockapiApplicationCSVDeduper func(a MockapiApplication, b MockapiApplication) *MockapiApplication

// MockapiApplicationCSVDedupeDisabled is set on whether the CSV writer should de-dupe values by key
var MockapiApplicationCSVDedupeDisabled bool

// NewMockapiApplicationCSVWriterSize creates a batch writer that will write each MockapiApplication into a CSV file
// this method will automatically de-duplicate entries using the primary key. if the checksum
// for a newer item with the same primary key doesn't match a previously sent item, the newer
// one will be used
func NewMockapiApplicationCSVWriterSize(w io.Writer, size int, dedupers ...MockapiApplicationCSVDeduper) (chan MockapiApplication, chan bool, error) {
	cw := csv.NewWriter(w)
	ch := make(chan MockapiApplication, size)
	done := make(chan bool)
	go func() {
		defer func() { done <- true }()
		dodedupe := !MockapiApplicationCSVDedupeDisabled
		var kv map[string]*MockapiApplication
		var deduper MockapiApplicationCSVDeduper
		if dedupers != nil && len(dedupers) > 0 {
			deduper = dedupers[0]
			dodedupe = true
		}
		if dodedupe {
			kv = make(map[string]*MockapiApplication)
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

// MockapiApplicationCSVDefaultSize is the default channel buffer size if not provided
var MockapiApplicationCSVDefaultSize = 100

// NewMockapiApplicationCSVWriter creates a batch writer that will write each MockapiApplication into a CSV file
func NewMockapiApplicationCSVWriter(w io.Writer, dedupers ...MockapiApplicationCSVDeduper) (chan MockapiApplication, chan bool, error) {
	return NewMockapiApplicationCSVWriterSize(w, MockapiApplicationCSVDefaultSize, dedupers...)
}

// NewMockapiApplicationCSVWriterDir creates a batch writer that will write each MockapiApplication into a CSV file named mockapi_application.csv.gz in dir
func NewMockapiApplicationCSVWriterDir(dir string, dedupers ...MockapiApplicationCSVDeduper) (chan MockapiApplication, chan bool, error) {
	return NewMockapiApplicationCSVWriterFile(filepath.Join(dir, "mockapi_application.csv.gz"), dedupers...)
}

// NewMockapiApplicationCSVWriterFile creates a batch writer that will write each MockapiApplication into a CSV file
func NewMockapiApplicationCSVWriterFile(fn string, dedupers ...MockapiApplicationCSVDeduper) (chan MockapiApplication, chan bool, error) {
	f, err := os.Create(fn)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening CSV file %s. %v", fn, err)
	}
	var fc io.WriteCloser = f
	if filepath.Ext(fn) == ".gz" {
		w, _ := gzip.NewWriterLevel(f, gzip.BestCompression)
		fc = w
	}
	ch, done, err := NewMockapiApplicationCSVWriter(fc, dedupers...)
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

type MockapiApplicationDBAction func(ctx context.Context, db DB, record MockapiApplication) error

// NewMockapiApplicationDBWriterSize creates a DB writer that will write each issue into the DB
func NewMockapiApplicationDBWriterSize(ctx context.Context, db DB, errors chan<- error, size int, actions ...MockapiApplicationDBAction) (chan MockapiApplication, chan bool, error) {
	ch := make(chan MockapiApplication, size)
	done := make(chan bool)
	var action MockapiApplicationDBAction
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

// NewMockapiApplicationDBWriter creates a DB writer that will write each issue into the DB
func NewMockapiApplicationDBWriter(ctx context.Context, db DB, errors chan<- error, actions ...MockapiApplicationDBAction) (chan MockapiApplication, chan bool, error) {
	return NewMockapiApplicationDBWriterSize(ctx, db, errors, 100, actions...)
}

// MockapiApplicationColumnID is the ID SQL column name for the MockapiApplication table
const MockapiApplicationColumnID = "id"

// MockapiApplicationEscapedColumnID is the escaped ID SQL column name for the MockapiApplication table
const MockapiApplicationEscapedColumnID = "`id`"

// MockapiApplicationColumnChecksum is the Checksum SQL column name for the MockapiApplication table
const MockapiApplicationColumnChecksum = "checksum"

// MockapiApplicationEscapedColumnChecksum is the escaped Checksum SQL column name for the MockapiApplication table
const MockapiApplicationEscapedColumnChecksum = "`checksum`"

// MockapiApplicationColumnCustomerID is the CustomerID SQL column name for the MockapiApplication table
const MockapiApplicationColumnCustomerID = "customer_id"

// MockapiApplicationEscapedColumnCustomerID is the escaped CustomerID SQL column name for the MockapiApplication table
const MockapiApplicationEscapedColumnCustomerID = "`customer_id`"

// MockapiApplicationColumnExtID is the ExtID SQL column name for the MockapiApplication table
const MockapiApplicationColumnExtID = "ext_id"

// MockapiApplicationEscapedColumnExtID is the escaped ExtID SQL column name for the MockapiApplication table
const MockapiApplicationEscapedColumnExtID = "`ext_id`"

// MockapiApplicationColumnName is the Name SQL column name for the MockapiApplication table
const MockapiApplicationColumnName = "name"

// MockapiApplicationEscapedColumnName is the escaped Name SQL column name for the MockapiApplication table
const MockapiApplicationEscapedColumnName = "`name`"

// GetID will return the MockapiApplication ID value
func (t *MockapiApplication) GetID() string {
	return t.ID
}

// SetID will set the MockapiApplication ID value
func (t *MockapiApplication) SetID(v string) {
	t.ID = v
}

// FindMockapiApplicationByID will find a MockapiApplication by ID
func FindMockapiApplicationByID(ctx context.Context, db DB, value string) (*MockapiApplication, error) {
	q := "SELECT `mockapi_application`.`id`,`mockapi_application`.`checksum`,`mockapi_application`.`customer_id`,`mockapi_application`.`ext_id`,`mockapi_application`.`name` FROM `mockapi_application` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CustomerID sql.NullString
	var _ExtID sql.NullInt64
	var _Name sql.NullString
	err := db.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_ExtID,
		&_Name,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &MockapiApplication{}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _Checksum.Valid {
		t.SetChecksum(_Checksum.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _ExtID.Valid {
		t.SetExtID(_ExtID.Int64)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	return t, nil
}

// FindMockapiApplicationByIDTx will find a MockapiApplication by ID using the provided transaction
func FindMockapiApplicationByIDTx(ctx context.Context, tx Tx, value string) (*MockapiApplication, error) {
	q := "SELECT `mockapi_application`.`id`,`mockapi_application`.`checksum`,`mockapi_application`.`customer_id`,`mockapi_application`.`ext_id`,`mockapi_application`.`name` FROM `mockapi_application` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CustomerID sql.NullString
	var _ExtID sql.NullInt64
	var _Name sql.NullString
	err := tx.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_ExtID,
		&_Name,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &MockapiApplication{}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _Checksum.Valid {
		t.SetChecksum(_Checksum.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _ExtID.Valid {
		t.SetExtID(_ExtID.Int64)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	return t, nil
}

// GetChecksum will return the MockapiApplication Checksum value
func (t *MockapiApplication) GetChecksum() string {
	if t.Checksum == nil {
		return ""
	}
	return *t.Checksum
}

// SetChecksum will set the MockapiApplication Checksum value
func (t *MockapiApplication) SetChecksum(v string) {
	t.Checksum = &v
}

// GetCustomerID will return the MockapiApplication CustomerID value
func (t *MockapiApplication) GetCustomerID() string {
	return t.CustomerID
}

// SetCustomerID will set the MockapiApplication CustomerID value
func (t *MockapiApplication) SetCustomerID(v string) {
	t.CustomerID = v
}

// FindMockapiApplicationsByCustomerID will find all MockapiApplications by the CustomerID value
func FindMockapiApplicationsByCustomerID(ctx context.Context, db DB, value string) ([]*MockapiApplication, error) {
	q := "SELECT `mockapi_application`.`id`,`mockapi_application`.`checksum`,`mockapi_application`.`customer_id`,`mockapi_application`.`ext_id`,`mockapi_application`.`name` FROM `mockapi_application` WHERE `customer_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*MockapiApplication, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _ExtID sql.NullInt64
		var _Name sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_ExtID,
			&_Name,
		)
		if err != nil {
			return nil, err
		}
		t := &MockapiApplication{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
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

// FindMockapiApplicationsByCustomerIDTx will find all MockapiApplications by the CustomerID value using the provided transaction
func FindMockapiApplicationsByCustomerIDTx(ctx context.Context, tx Tx, value string) ([]*MockapiApplication, error) {
	q := "SELECT `mockapi_application`.`id`,`mockapi_application`.`checksum`,`mockapi_application`.`customer_id`,`mockapi_application`.`ext_id`,`mockapi_application`.`name` FROM `mockapi_application` WHERE `customer_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*MockapiApplication, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _ExtID sql.NullInt64
		var _Name sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_ExtID,
			&_Name,
		)
		if err != nil {
			return nil, err
		}
		t := &MockapiApplication{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
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

// GetExtID will return the MockapiApplication ExtID value
func (t *MockapiApplication) GetExtID() int64 {
	return t.ExtID
}

// SetExtID will set the MockapiApplication ExtID value
func (t *MockapiApplication) SetExtID(v int64) {
	t.ExtID = v
}

// GetName will return the MockapiApplication Name value
func (t *MockapiApplication) GetName() string {
	return t.Name
}

// SetName will set the MockapiApplication Name value
func (t *MockapiApplication) SetName(v string) {
	t.Name = v
}

func (t *MockapiApplication) toTimestamp(value time.Time) *timestamp.Timestamp {
	ts, _ := ptypes.TimestampProto(value)
	return ts
}

// DBCreateMockapiApplicationTable will create the MockapiApplication table
func DBCreateMockapiApplicationTable(ctx context.Context, db DB) error {
	q := "CREATE TABLE `mockapi_application` (`id` VARCHAR(64) NOT NULL PRIMARY KEY,`checksum` CHAR(64),`customer_id` VARCHAR(64) NOT NULL,`ext_id` BIGINT NOT NULL,`name`TEXT NOT NULL,INDEX mockapi_application_customer_id_index (`customer_id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBCreateMockapiApplicationTableTx will create the MockapiApplication table using the provided transction
func DBCreateMockapiApplicationTableTx(ctx context.Context, tx Tx) error {
	q := "CREATE TABLE `mockapi_application` (`id` VARCHAR(64) NOT NULL PRIMARY KEY,`checksum` CHAR(64),`customer_id` VARCHAR(64) NOT NULL,`ext_id` BIGINT NOT NULL,`name`TEXT NOT NULL,INDEX mockapi_application_customer_id_index (`customer_id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// DBDropMockapiApplicationTable will drop the MockapiApplication table
func DBDropMockapiApplicationTable(ctx context.Context, db DB) error {
	q := "DROP TABLE IF EXISTS `mockapi_application`"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBDropMockapiApplicationTableTx will drop the MockapiApplication table using the provided transaction
func DBDropMockapiApplicationTableTx(ctx context.Context, tx Tx) error {
	q := "DROP TABLE IF EXISTS `mockapi_application`"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// CalculateChecksum will calculate a checksum of the SHA1 of all field values
func (t *MockapiApplication) CalculateChecksum() string {
	return orm.HashStrings(
		orm.ToString(t.ID),
		orm.ToString(t.CustomerID),
		orm.ToString(t.ExtID),
		orm.ToString(t.Name),
	)
}

// DBCreate will create a new MockapiApplication record in the database
func (t *MockapiApplication) DBCreate(ctx context.Context, db DB) (sql.Result, error) {
	q := "INSERT INTO `mockapi_application` (`mockapi_application`.`id`,`mockapi_application`.`checksum`,`mockapi_application`.`customer_id`,`mockapi_application`.`ext_id`,`mockapi_application`.`name`) VALUES (?,?,?,?,?)"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLInt64(t.ExtID),
		orm.ToSQLString(t.Name),
	)
}

// DBCreateTx will create a new MockapiApplication record in the database using the provided transaction
func (t *MockapiApplication) DBCreateTx(ctx context.Context, tx Tx) (sql.Result, error) {
	q := "INSERT INTO `mockapi_application` (`mockapi_application`.`id`,`mockapi_application`.`checksum`,`mockapi_application`.`customer_id`,`mockapi_application`.`ext_id`,`mockapi_application`.`name`) VALUES (?,?,?,?,?)"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLInt64(t.ExtID),
		orm.ToSQLString(t.Name),
	)
}

// DBCreateIgnoreDuplicate will upsert the MockapiApplication record in the database
func (t *MockapiApplication) DBCreateIgnoreDuplicate(ctx context.Context, db DB) (sql.Result, error) {
	q := "INSERT INTO `mockapi_application` (`mockapi_application`.`id`,`mockapi_application`.`checksum`,`mockapi_application`.`customer_id`,`mockapi_application`.`ext_id`,`mockapi_application`.`name`) VALUES (?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLInt64(t.ExtID),
		orm.ToSQLString(t.Name),
	)
}

// DBCreateIgnoreDuplicateTx will upsert the MockapiApplication record in the database using the provided transaction
func (t *MockapiApplication) DBCreateIgnoreDuplicateTx(ctx context.Context, tx Tx) (sql.Result, error) {
	q := "INSERT INTO `mockapi_application` (`mockapi_application`.`id`,`mockapi_application`.`checksum`,`mockapi_application`.`customer_id`,`mockapi_application`.`ext_id`,`mockapi_application`.`name`) VALUES (?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLInt64(t.ExtID),
		orm.ToSQLString(t.Name),
	)
}

// DeleteAllMockapiApplications deletes all MockapiApplication records in the database with optional filters
func DeleteAllMockapiApplications(ctx context.Context, db DB, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(MockapiApplicationTableName),
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

// DeleteAllMockapiApplicationsTx deletes all MockapiApplication records in the database with optional filters using the provided transaction
func DeleteAllMockapiApplicationsTx(ctx context.Context, tx Tx, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(MockapiApplicationTableName),
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

// DBDelete will delete this MockapiApplication record in the database
func (t *MockapiApplication) DBDelete(ctx context.Context, db DB) (bool, error) {
	q := "DELETE FROM `mockapi_application` WHERE `id` = ?"
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

// DBDeleteTx will delete this MockapiApplication record in the database using the provided transaction
func (t *MockapiApplication) DBDeleteTx(ctx context.Context, tx Tx) (bool, error) {
	q := "DELETE FROM `mockapi_application` WHERE `id` = ?"
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

// DBUpdate will update the MockapiApplication record in the database
func (t *MockapiApplication) DBUpdate(ctx context.Context, db DB) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `mockapi_application` SET `checksum`=?,`customer_id`=?,`ext_id`=?,`name`=? WHERE `id`=?"
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLInt64(t.ExtID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.ID),
	)
}

// DBUpdateTx will update the MockapiApplication record in the database using the provided transaction
func (t *MockapiApplication) DBUpdateTx(ctx context.Context, tx Tx) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `mockapi_application` SET `checksum`=?,`customer_id`=?,`ext_id`=?,`name`=? WHERE `id`=?"
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLInt64(t.ExtID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.ID),
	)
}

// DBUpsert will upsert the MockapiApplication record in the database
func (t *MockapiApplication) DBUpsert(ctx context.Context, db DB, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `mockapi_application` (`mockapi_application`.`id`,`mockapi_application`.`checksum`,`mockapi_application`.`customer_id`,`mockapi_application`.`ext_id`,`mockapi_application`.`name`) VALUES (?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `mockapi_application` (`mockapi_application`.`id`,`mockapi_application`.`checksum`,`mockapi_application`.`customer_id`,`mockapi_application`.`ext_id`,`mockapi_application`.`name`) VALUES (?,?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`customer_id`=VALUES(`customer_id`),`ext_id`=VALUES(`ext_id`),`name`=VALUES(`name`)"
	}
	r, err := db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLInt64(t.ExtID),
		orm.ToSQLString(t.Name),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBUpsertTx will upsert the MockapiApplication record in the database using the provided transaction
func (t *MockapiApplication) DBUpsertTx(ctx context.Context, tx Tx, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `mockapi_application` (`mockapi_application`.`id`,`mockapi_application`.`checksum`,`mockapi_application`.`customer_id`,`mockapi_application`.`ext_id`,`mockapi_application`.`name`) VALUES (?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `mockapi_application` (`mockapi_application`.`id`,`mockapi_application`.`checksum`,`mockapi_application`.`customer_id`,`mockapi_application`.`ext_id`,`mockapi_application`.`name`) VALUES (?,?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`customer_id`=VALUES(`customer_id`),`ext_id`=VALUES(`ext_id`),`name`=VALUES(`name`)"
	}
	r, err := tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLInt64(t.ExtID),
		orm.ToSQLString(t.Name),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBFindOne will find a MockapiApplication record in the database with the primary key
func (t *MockapiApplication) DBFindOne(ctx context.Context, db DB, value string) (bool, error) {
	q := "SELECT `mockapi_application`.`id`,`mockapi_application`.`checksum`,`mockapi_application`.`customer_id`,`mockapi_application`.`ext_id`,`mockapi_application`.`name` FROM `mockapi_application` WHERE `id` = ? LIMIT 1"
	row := db.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CustomerID sql.NullString
	var _ExtID sql.NullInt64
	var _Name sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
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
	if _ExtID.Valid {
		t.SetExtID(_ExtID.Int64)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	return true, nil
}

// DBFindOneTx will find a MockapiApplication record in the database with the primary key using the provided transaction
func (t *MockapiApplication) DBFindOneTx(ctx context.Context, tx Tx, value string) (bool, error) {
	q := "SELECT `mockapi_application`.`id`,`mockapi_application`.`checksum`,`mockapi_application`.`customer_id`,`mockapi_application`.`ext_id`,`mockapi_application`.`name` FROM `mockapi_application` WHERE `id` = ? LIMIT 1"
	row := tx.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CustomerID sql.NullString
	var _ExtID sql.NullInt64
	var _Name sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
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
	if _ExtID.Valid {
		t.SetExtID(_ExtID.Int64)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	return true, nil
}

// FindMockapiApplications will find a MockapiApplication record in the database with the provided parameters
func FindMockapiApplications(ctx context.Context, db DB, _params ...interface{}) ([]*MockapiApplication, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("customer_id"),
		orm.Column("ext_id"),
		orm.Column("name"),
		orm.Table(MockapiApplicationTableName),
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
	results := make([]*MockapiApplication, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _ExtID sql.NullInt64
		var _Name sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_ExtID,
			&_Name,
		)
		if err != nil {
			return nil, err
		}
		t := &MockapiApplication{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
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

// FindMockapiApplicationsTx will find a MockapiApplication record in the database with the provided parameters using the provided transaction
func FindMockapiApplicationsTx(ctx context.Context, tx Tx, _params ...interface{}) ([]*MockapiApplication, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("customer_id"),
		orm.Column("ext_id"),
		orm.Column("name"),
		orm.Table(MockapiApplicationTableName),
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
	results := make([]*MockapiApplication, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _ExtID sql.NullInt64
		var _Name sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_ExtID,
			&_Name,
		)
		if err != nil {
			return nil, err
		}
		t := &MockapiApplication{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
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

// DBFind will find a MockapiApplication record in the database with the provided parameters
func (t *MockapiApplication) DBFind(ctx context.Context, db DB, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("customer_id"),
		orm.Column("ext_id"),
		orm.Column("name"),
		orm.Table(MockapiApplicationTableName),
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
	var _ExtID sql.NullInt64
	var _Name sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
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
	if _ExtID.Valid {
		t.SetExtID(_ExtID.Int64)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	return true, nil
}

// DBFindTx will find a MockapiApplication record in the database with the provided parameters using the provided transaction
func (t *MockapiApplication) DBFindTx(ctx context.Context, tx Tx, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("customer_id"),
		orm.Column("ext_id"),
		orm.Column("name"),
		orm.Table(MockapiApplicationTableName),
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
	var _ExtID sql.NullInt64
	var _Name sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
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
	if _ExtID.Valid {
		t.SetExtID(_ExtID.Int64)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	return true, nil
}

// CountMockapiApplications will find the count of MockapiApplication records in the database
func CountMockapiApplications(ctx context.Context, db DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(MockapiApplicationTableName),
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

// CountMockapiApplicationsTx will find the count of MockapiApplication records in the database using the provided transaction
func CountMockapiApplicationsTx(ctx context.Context, tx Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(MockapiApplicationTableName),
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

// DBCount will find the count of MockapiApplication records in the database
func (t *MockapiApplication) DBCount(ctx context.Context, db DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(MockapiApplicationTableName),
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

// DBCountTx will find the count of MockapiApplication records in the database using the provided transaction
func (t *MockapiApplication) DBCountTx(ctx context.Context, tx Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(MockapiApplicationTableName),
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

// DBExists will return true if the MockapiApplication record exists in the database
func (t *MockapiApplication) DBExists(ctx context.Context, db DB) (bool, error) {
	q := "SELECT `id` FROM `mockapi_application` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := db.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// DBExistsTx will return true if the MockapiApplication record exists in the database using the provided transaction
func (t *MockapiApplication) DBExistsTx(ctx context.Context, tx Tx) (bool, error) {
	q := "SELECT `id` FROM `mockapi_application` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := tx.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// PrimaryKeyColumn returns the column name for the primary key
func (t *MockapiApplication) PrimaryKeyColumn() string {
	return MockapiApplicationColumnID
}

// PrimaryKeyColumnType returns the primary key column Go type as a string
func (t *MockapiApplication) PrimaryKeyColumnType() string {
	return "string"
}

// PrimaryKey returns the primary key column value
func (t *MockapiApplication) PrimaryKey() interface{} {
	return t.ID
}
