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

var _ Model = (*NewrelicApplication)(nil)
var _ CSVWriter = (*NewrelicApplication)(nil)
var _ JSONWriter = (*NewrelicApplication)(nil)
var _ Checksum = (*NewrelicApplication)(nil)

// NewrelicApplicationTableName is the name of the table in SQL
const NewrelicApplicationTableName = "newrelic_application"

var NewrelicApplicationColumns = []string{
	"id",
	"checksum",
	"customer_id",
	"ext_id",
	"name",
	"health_status",
	"reporting",
}

// NewrelicApplication table
type NewrelicApplication struct {
	Checksum     *string `json:"checksum,omitempty"`
	CustomerID   string  `json:"customer_id"`
	ExtID        int64   `json:"ext_id"`
	HealthStatus *string `json:"health_status,omitempty"`
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	Reporting    bool    `json:"reporting"`
}

// TableName returns the SQL table name for NewrelicApplication and satifies the Model interface
func (t *NewrelicApplication) TableName() string {
	return NewrelicApplicationTableName
}

// ToCSV will serialize the NewrelicApplication instance to a CSV compatible array of strings
func (t *NewrelicApplication) ToCSV() []string {
	return []string{
		t.ID,
		t.CalculateChecksum(),
		t.CustomerID,
		toCSVString(t.ExtID),
		t.Name,
		toCSVString(t.HealthStatus),
		toCSVBool(t.Reporting),
	}
}

// WriteCSV will serialize the NewrelicApplication instance to the writer as CSV and satisfies the CSVWriter interface
func (t *NewrelicApplication) WriteCSV(w *csv.Writer) error {
	return w.Write(t.ToCSV())
}

// WriteJSON will serialize the NewrelicApplication instance to the writer as JSON and satisfies the JSONWriter interface
func (t *NewrelicApplication) WriteJSON(w io.Writer, indent ...bool) error {
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

// NewNewrelicApplicationReader creates a JSON reader which can read in NewrelicApplication objects serialized as JSON either as an array, single object or json new lines
// and writes each NewrelicApplication to the channel provided
func NewNewrelicApplicationReader(r io.Reader, ch chan<- NewrelicApplication) error {
	return orm.Deserialize(r, func(buf json.RawMessage) error {
		dec := json.NewDecoder(bytes.NewBuffer(buf))
		e := NewrelicApplication{}
		if err := dec.Decode(&e); err != nil {
			return err
		}
		ch <- e
		return nil
	})
}

// NewCSVNewrelicApplicationReaderDir will read the reader as CSV and emit each record to the channel provided
func NewCSVNewrelicApplicationReader(r io.Reader, ch chan<- NewrelicApplication) error {
	cr := csv.NewReader(r)
	for {
		record, err := cr.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		ch <- NewrelicApplication{
			ID:           record[0],
			Checksum:     fromStringPointer(record[1]),
			CustomerID:   record[2],
			ExtID:        fromCSVInt64(record[3]),
			Name:         record[4],
			HealthStatus: fromStringPointer(record[5]),
			Reporting:    fromCSVBool(record[6]),
		}
	}
	return nil
}

// NewCSVNewrelicApplicationReaderFile will read the file as a CSV and emit each record to the channel provided
func NewCSVNewrelicApplicationReaderFile(fp string, ch chan<- NewrelicApplication) error {
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
	return NewCSVNewrelicApplicationReader(fc, ch)
}

// NewCSVNewrelicApplicationReaderDir will read the newrelic_application.csv.gz file as a CSV and emit each record to the channel provided
func NewCSVNewrelicApplicationReaderDir(dir string, ch chan<- NewrelicApplication) error {
	return NewCSVNewrelicApplicationReaderFile(filepath.Join(dir, "newrelic_application.csv.gz"), ch)
}

// NewrelicApplicationCSVDeduper is a function callback which takes the existing value (a) and the new value (b)
// and the return value should be the one to use (or a new one, if applicable). return nil
// to skip processing of this record
type NewrelicApplicationCSVDeduper func(a NewrelicApplication, b NewrelicApplication) *NewrelicApplication

// NewrelicApplicationCSVDedupeDisabled is set on whether the CSV writer should de-dupe values by key
var NewrelicApplicationCSVDedupeDisabled bool

// NewNewrelicApplicationCSVWriterSize creates a batch writer that will write each NewrelicApplication into a CSV file
// this method will automatically de-duplicate entries using the primary key. if the checksum
// for a newer item with the same primary key doesn't match a previously sent item, the newer
// one will be used
func NewNewrelicApplicationCSVWriterSize(w io.Writer, size int, dedupers ...NewrelicApplicationCSVDeduper) (chan NewrelicApplication, chan bool, error) {
	cw := csv.NewWriter(w)
	ch := make(chan NewrelicApplication, size)
	done := make(chan bool)
	go func() {
		defer func() { done <- true }()
		dodedupe := !NewrelicApplicationCSVDedupeDisabled
		var kv map[string]*NewrelicApplication
		var deduper NewrelicApplicationCSVDeduper
		if dedupers != nil && len(dedupers) > 0 {
			deduper = dedupers[0]
			dodedupe = true
		}
		if dodedupe {
			kv = make(map[string]*NewrelicApplication)
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

// NewrelicApplicationCSVDefaultSize is the default channel buffer size if not provided
var NewrelicApplicationCSVDefaultSize = 100

// NewNewrelicApplicationCSVWriter creates a batch writer that will write each NewrelicApplication into a CSV file
func NewNewrelicApplicationCSVWriter(w io.Writer, dedupers ...NewrelicApplicationCSVDeduper) (chan NewrelicApplication, chan bool, error) {
	return NewNewrelicApplicationCSVWriterSize(w, NewrelicApplicationCSVDefaultSize, dedupers...)
}

// NewNewrelicApplicationCSVWriterDir creates a batch writer that will write each NewrelicApplication into a CSV file named newrelic_application.csv.gz in dir
func NewNewrelicApplicationCSVWriterDir(dir string, dedupers ...NewrelicApplicationCSVDeduper) (chan NewrelicApplication, chan bool, error) {
	return NewNewrelicApplicationCSVWriterFile(filepath.Join(dir, "newrelic_application.csv.gz"), dedupers...)
}

// NewNewrelicApplicationCSVWriterFile creates a batch writer that will write each NewrelicApplication into a CSV file
func NewNewrelicApplicationCSVWriterFile(fn string, dedupers ...NewrelicApplicationCSVDeduper) (chan NewrelicApplication, chan bool, error) {
	f, err := os.Create(fn)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening CSV file %s. %v", fn, err)
	}
	var fc io.WriteCloser = f
	if filepath.Ext(fn) == ".gz" {
		w, _ := gzip.NewWriterLevel(f, gzip.BestCompression)
		fc = w
	}
	ch, done, err := NewNewrelicApplicationCSVWriter(fc, dedupers...)
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

type NewrelicApplicationDBAction func(ctx context.Context, db *sql.DB, record NewrelicApplication) error

// NewNewrelicApplicationDBWriterSize creates a DB writer that will write each issue into the DB
func NewNewrelicApplicationDBWriterSize(ctx context.Context, db *sql.DB, errors chan<- error, size int, actions ...NewrelicApplicationDBAction) (chan NewrelicApplication, chan bool, error) {
	ch := make(chan NewrelicApplication, size)
	done := make(chan bool)
	var action NewrelicApplicationDBAction
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

// NewNewrelicApplicationDBWriter creates a DB writer that will write each issue into the DB
func NewNewrelicApplicationDBWriter(ctx context.Context, db *sql.DB, errors chan<- error, actions ...NewrelicApplicationDBAction) (chan NewrelicApplication, chan bool, error) {
	return NewNewrelicApplicationDBWriterSize(ctx, db, errors, 100, actions...)
}

// NewrelicApplicationColumnID is the ID SQL column name for the NewrelicApplication table
const NewrelicApplicationColumnID = "id"

// NewrelicApplicationEscapedColumnID is the escaped ID SQL column name for the NewrelicApplication table
const NewrelicApplicationEscapedColumnID = "`id`"

// NewrelicApplicationColumnChecksum is the Checksum SQL column name for the NewrelicApplication table
const NewrelicApplicationColumnChecksum = "checksum"

// NewrelicApplicationEscapedColumnChecksum is the escaped Checksum SQL column name for the NewrelicApplication table
const NewrelicApplicationEscapedColumnChecksum = "`checksum`"

// NewrelicApplicationColumnCustomerID is the CustomerID SQL column name for the NewrelicApplication table
const NewrelicApplicationColumnCustomerID = "customer_id"

// NewrelicApplicationEscapedColumnCustomerID is the escaped CustomerID SQL column name for the NewrelicApplication table
const NewrelicApplicationEscapedColumnCustomerID = "`customer_id`"

// NewrelicApplicationColumnExtID is the ExtID SQL column name for the NewrelicApplication table
const NewrelicApplicationColumnExtID = "ext_id"

// NewrelicApplicationEscapedColumnExtID is the escaped ExtID SQL column name for the NewrelicApplication table
const NewrelicApplicationEscapedColumnExtID = "`ext_id`"

// NewrelicApplicationColumnName is the Name SQL column name for the NewrelicApplication table
const NewrelicApplicationColumnName = "name"

// NewrelicApplicationEscapedColumnName is the escaped Name SQL column name for the NewrelicApplication table
const NewrelicApplicationEscapedColumnName = "`name`"

// NewrelicApplicationColumnHealthStatus is the HealthStatus SQL column name for the NewrelicApplication table
const NewrelicApplicationColumnHealthStatus = "health_status"

// NewrelicApplicationEscapedColumnHealthStatus is the escaped HealthStatus SQL column name for the NewrelicApplication table
const NewrelicApplicationEscapedColumnHealthStatus = "`health_status`"

// NewrelicApplicationColumnReporting is the Reporting SQL column name for the NewrelicApplication table
const NewrelicApplicationColumnReporting = "reporting"

// NewrelicApplicationEscapedColumnReporting is the escaped Reporting SQL column name for the NewrelicApplication table
const NewrelicApplicationEscapedColumnReporting = "`reporting`"

// GetID will return the NewrelicApplication ID value
func (t *NewrelicApplication) GetID() string {
	return t.ID
}

// SetID will set the NewrelicApplication ID value
func (t *NewrelicApplication) SetID(v string) {
	t.ID = v
}

// FindNewrelicApplicationByID will find a NewrelicApplication by ID
func FindNewrelicApplicationByID(ctx context.Context, db *sql.DB, value string) (*NewrelicApplication, error) {
	q := "SELECT `newrelic_application`.`id`,`newrelic_application`.`checksum`,`newrelic_application`.`customer_id`,`newrelic_application`.`ext_id`,`newrelic_application`.`name`,`newrelic_application`.`health_status`,`newrelic_application`.`reporting` FROM `newrelic_application` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CustomerID sql.NullString
	var _ExtID sql.NullInt64
	var _Name sql.NullString
	var _HealthStatus sql.NullString
	var _Reporting sql.NullBool
	err := db.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_ExtID,
		&_Name,
		&_HealthStatus,
		&_Reporting,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &NewrelicApplication{}
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
	if _HealthStatus.Valid {
		t.SetHealthStatus(_HealthStatus.String)
	}
	if _Reporting.Valid {
		t.SetReporting(_Reporting.Bool)
	}
	return t, nil
}

// FindNewrelicApplicationByIDTx will find a NewrelicApplication by ID using the provided transaction
func FindNewrelicApplicationByIDTx(ctx context.Context, tx *sql.Tx, value string) (*NewrelicApplication, error) {
	q := "SELECT `newrelic_application`.`id`,`newrelic_application`.`checksum`,`newrelic_application`.`customer_id`,`newrelic_application`.`ext_id`,`newrelic_application`.`name`,`newrelic_application`.`health_status`,`newrelic_application`.`reporting` FROM `newrelic_application` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CustomerID sql.NullString
	var _ExtID sql.NullInt64
	var _Name sql.NullString
	var _HealthStatus sql.NullString
	var _Reporting sql.NullBool
	err := tx.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_ExtID,
		&_Name,
		&_HealthStatus,
		&_Reporting,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &NewrelicApplication{}
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
	if _HealthStatus.Valid {
		t.SetHealthStatus(_HealthStatus.String)
	}
	if _Reporting.Valid {
		t.SetReporting(_Reporting.Bool)
	}
	return t, nil
}

// GetChecksum will return the NewrelicApplication Checksum value
func (t *NewrelicApplication) GetChecksum() string {
	if t.Checksum == nil {
		return ""
	}
	return *t.Checksum
}

// SetChecksum will set the NewrelicApplication Checksum value
func (t *NewrelicApplication) SetChecksum(v string) {
	t.Checksum = &v
}

// GetCustomerID will return the NewrelicApplication CustomerID value
func (t *NewrelicApplication) GetCustomerID() string {
	return t.CustomerID
}

// SetCustomerID will set the NewrelicApplication CustomerID value
func (t *NewrelicApplication) SetCustomerID(v string) {
	t.CustomerID = v
}

// FindNewrelicApplicationsByCustomerID will find all NewrelicApplications by the CustomerID value
func FindNewrelicApplicationsByCustomerID(ctx context.Context, db *sql.DB, value string) ([]*NewrelicApplication, error) {
	q := "SELECT `newrelic_application`.`id`,`newrelic_application`.`checksum`,`newrelic_application`.`customer_id`,`newrelic_application`.`ext_id`,`newrelic_application`.`name`,`newrelic_application`.`health_status`,`newrelic_application`.`reporting` FROM `newrelic_application` WHERE `customer_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*NewrelicApplication, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _ExtID sql.NullInt64
		var _Name sql.NullString
		var _HealthStatus sql.NullString
		var _Reporting sql.NullBool
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_ExtID,
			&_Name,
			&_HealthStatus,
			&_Reporting,
		)
		if err != nil {
			return nil, err
		}
		t := &NewrelicApplication{}
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
		if _HealthStatus.Valid {
			t.SetHealthStatus(_HealthStatus.String)
		}
		if _Reporting.Valid {
			t.SetReporting(_Reporting.Bool)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindNewrelicApplicationsByCustomerIDTx will find all NewrelicApplications by the CustomerID value using the provided transaction
func FindNewrelicApplicationsByCustomerIDTx(ctx context.Context, tx *sql.Tx, value string) ([]*NewrelicApplication, error) {
	q := "SELECT `newrelic_application`.`id`,`newrelic_application`.`checksum`,`newrelic_application`.`customer_id`,`newrelic_application`.`ext_id`,`newrelic_application`.`name`,`newrelic_application`.`health_status`,`newrelic_application`.`reporting` FROM `newrelic_application` WHERE `customer_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*NewrelicApplication, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _ExtID sql.NullInt64
		var _Name sql.NullString
		var _HealthStatus sql.NullString
		var _Reporting sql.NullBool
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_ExtID,
			&_Name,
			&_HealthStatus,
			&_Reporting,
		)
		if err != nil {
			return nil, err
		}
		t := &NewrelicApplication{}
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
		if _HealthStatus.Valid {
			t.SetHealthStatus(_HealthStatus.String)
		}
		if _Reporting.Valid {
			t.SetReporting(_Reporting.Bool)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetExtID will return the NewrelicApplication ExtID value
func (t *NewrelicApplication) GetExtID() int64 {
	return t.ExtID
}

// SetExtID will set the NewrelicApplication ExtID value
func (t *NewrelicApplication) SetExtID(v int64) {
	t.ExtID = v
}

// GetName will return the NewrelicApplication Name value
func (t *NewrelicApplication) GetName() string {
	return t.Name
}

// SetName will set the NewrelicApplication Name value
func (t *NewrelicApplication) SetName(v string) {
	t.Name = v
}

// GetHealthStatus will return the NewrelicApplication HealthStatus value
func (t *NewrelicApplication) GetHealthStatus() string {
	if t.HealthStatus == nil {
		return ""
	}
	return *t.HealthStatus
}

// SetHealthStatus will set the NewrelicApplication HealthStatus value
func (t *NewrelicApplication) SetHealthStatus(v string) {
	t.HealthStatus = &v
}

// GetReporting will return the NewrelicApplication Reporting value
func (t *NewrelicApplication) GetReporting() bool {
	return t.Reporting
}

// SetReporting will set the NewrelicApplication Reporting value
func (t *NewrelicApplication) SetReporting(v bool) {
	t.Reporting = v
}

func (t *NewrelicApplication) toTimestamp(value time.Time) *timestamp.Timestamp {
	ts, _ := ptypes.TimestampProto(value)
	return ts
}

// DBCreateNewrelicApplicationTable will create the NewrelicApplication table
func DBCreateNewrelicApplicationTable(ctx context.Context, db *sql.DB) error {
	q := "CREATE TABLE `newrelic_application` (`id` VARCHAR(64) NOT NULL PRIMARY KEY,`checksum` CHAR(64),`customer_id` VARCHAR(64) NOT NULL,`ext_id`BIGINT NOT NULL,`name` TEXT NOT NULL,`health_status` TEXT,`reporting`BOOL NOT NULL,INDEX newrelic_application_customer_id_index (`customer_id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBCreateNewrelicApplicationTableTx will create the NewrelicApplication table using the provided transction
func DBCreateNewrelicApplicationTableTx(ctx context.Context, tx *sql.Tx) error {
	q := "CREATE TABLE `newrelic_application` (`id` VARCHAR(64) NOT NULL PRIMARY KEY,`checksum` CHAR(64),`customer_id` VARCHAR(64) NOT NULL,`ext_id`BIGINT NOT NULL,`name` TEXT NOT NULL,`health_status` TEXT,`reporting`BOOL NOT NULL,INDEX newrelic_application_customer_id_index (`customer_id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// DBDropNewrelicApplicationTable will drop the NewrelicApplication table
func DBDropNewrelicApplicationTable(ctx context.Context, db *sql.DB) error {
	q := "DROP TABLE IF EXISTS `newrelic_application`"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBDropNewrelicApplicationTableTx will drop the NewrelicApplication table using the provided transaction
func DBDropNewrelicApplicationTableTx(ctx context.Context, tx *sql.Tx) error {
	q := "DROP TABLE IF EXISTS `newrelic_application`"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// CalculateChecksum will calculate a checksum of the SHA1 of all field values
func (t *NewrelicApplication) CalculateChecksum() string {
	return orm.HashStrings(
		orm.ToString(t.ID),
		orm.ToString(t.CustomerID),
		orm.ToString(t.ExtID),
		orm.ToString(t.Name),
		orm.ToString(t.HealthStatus),
		orm.ToString(t.Reporting),
	)
}

// DBCreate will create a new NewrelicApplication record in the database
func (t *NewrelicApplication) DBCreate(ctx context.Context, db *sql.DB) (sql.Result, error) {
	q := "INSERT INTO `newrelic_application` (`newrelic_application`.`id`,`newrelic_application`.`checksum`,`newrelic_application`.`customer_id`,`newrelic_application`.`ext_id`,`newrelic_application`.`name`,`newrelic_application`.`health_status`,`newrelic_application`.`reporting`) VALUES (?,?,?,?,?,?,?)"
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
		orm.ToSQLString(t.HealthStatus),
		orm.ToSQLBool(t.Reporting),
	)
}

// DBCreateTx will create a new NewrelicApplication record in the database using the provided transaction
func (t *NewrelicApplication) DBCreateTx(ctx context.Context, tx *sql.Tx) (sql.Result, error) {
	q := "INSERT INTO `newrelic_application` (`newrelic_application`.`id`,`newrelic_application`.`checksum`,`newrelic_application`.`customer_id`,`newrelic_application`.`ext_id`,`newrelic_application`.`name`,`newrelic_application`.`health_status`,`newrelic_application`.`reporting`) VALUES (?,?,?,?,?,?,?)"
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
		orm.ToSQLString(t.HealthStatus),
		orm.ToSQLBool(t.Reporting),
	)
}

// DBCreateIgnoreDuplicate will upsert the NewrelicApplication record in the database
func (t *NewrelicApplication) DBCreateIgnoreDuplicate(ctx context.Context, db *sql.DB) (sql.Result, error) {
	q := "INSERT INTO `newrelic_application` (`newrelic_application`.`id`,`newrelic_application`.`checksum`,`newrelic_application`.`customer_id`,`newrelic_application`.`ext_id`,`newrelic_application`.`name`,`newrelic_application`.`health_status`,`newrelic_application`.`reporting`) VALUES (?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
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
		orm.ToSQLString(t.HealthStatus),
		orm.ToSQLBool(t.Reporting),
	)
}

// DBCreateIgnoreDuplicateTx will upsert the NewrelicApplication record in the database using the provided transaction
func (t *NewrelicApplication) DBCreateIgnoreDuplicateTx(ctx context.Context, tx *sql.Tx) (sql.Result, error) {
	q := "INSERT INTO `newrelic_application` (`newrelic_application`.`id`,`newrelic_application`.`checksum`,`newrelic_application`.`customer_id`,`newrelic_application`.`ext_id`,`newrelic_application`.`name`,`newrelic_application`.`health_status`,`newrelic_application`.`reporting`) VALUES (?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
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
		orm.ToSQLString(t.HealthStatus),
		orm.ToSQLBool(t.Reporting),
	)
}

// DeleteAllNewrelicApplications deletes all NewrelicApplication records in the database with optional filters
func DeleteAllNewrelicApplications(ctx context.Context, db *sql.DB, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(NewrelicApplicationTableName),
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

// DeleteAllNewrelicApplicationsTx deletes all NewrelicApplication records in the database with optional filters using the provided transaction
func DeleteAllNewrelicApplicationsTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(NewrelicApplicationTableName),
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

// DBDelete will delete this NewrelicApplication record in the database
func (t *NewrelicApplication) DBDelete(ctx context.Context, db *sql.DB) (bool, error) {
	q := "DELETE FROM `newrelic_application` WHERE `id` = ?"
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

// DBDeleteTx will delete this NewrelicApplication record in the database using the provided transaction
func (t *NewrelicApplication) DBDeleteTx(ctx context.Context, tx *sql.Tx) (bool, error) {
	q := "DELETE FROM `newrelic_application` WHERE `id` = ?"
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

// DBUpdate will update the NewrelicApplication record in the database
func (t *NewrelicApplication) DBUpdate(ctx context.Context, db *sql.DB) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `newrelic_application` SET `checksum`=?,`customer_id`=?,`ext_id`=?,`name`=?,`health_status`=?,`reporting`=? WHERE `id`=?"
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLInt64(t.ExtID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.HealthStatus),
		orm.ToSQLBool(t.Reporting),
		orm.ToSQLString(t.ID),
	)
}

// DBUpdateTx will update the NewrelicApplication record in the database using the provided transaction
func (t *NewrelicApplication) DBUpdateTx(ctx context.Context, tx *sql.Tx) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `newrelic_application` SET `checksum`=?,`customer_id`=?,`ext_id`=?,`name`=?,`health_status`=?,`reporting`=? WHERE `id`=?"
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLInt64(t.ExtID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.HealthStatus),
		orm.ToSQLBool(t.Reporting),
		orm.ToSQLString(t.ID),
	)
}

// DBUpsert will upsert the NewrelicApplication record in the database
func (t *NewrelicApplication) DBUpsert(ctx context.Context, db *sql.DB, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `newrelic_application` (`newrelic_application`.`id`,`newrelic_application`.`checksum`,`newrelic_application`.`customer_id`,`newrelic_application`.`ext_id`,`newrelic_application`.`name`,`newrelic_application`.`health_status`,`newrelic_application`.`reporting`) VALUES (?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `newrelic_application` (`newrelic_application`.`id`,`newrelic_application`.`checksum`,`newrelic_application`.`customer_id`,`newrelic_application`.`ext_id`,`newrelic_application`.`name`,`newrelic_application`.`health_status`,`newrelic_application`.`reporting`) VALUES (?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`customer_id`=VALUES(`customer_id`),`ext_id`=VALUES(`ext_id`),`name`=VALUES(`name`),`health_status`=VALUES(`health_status`),`reporting`=VALUES(`reporting`)"
	}
	r, err := db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLInt64(t.ExtID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.HealthStatus),
		orm.ToSQLBool(t.Reporting),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBUpsertTx will upsert the NewrelicApplication record in the database using the provided transaction
func (t *NewrelicApplication) DBUpsertTx(ctx context.Context, tx *sql.Tx, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `newrelic_application` (`newrelic_application`.`id`,`newrelic_application`.`checksum`,`newrelic_application`.`customer_id`,`newrelic_application`.`ext_id`,`newrelic_application`.`name`,`newrelic_application`.`health_status`,`newrelic_application`.`reporting`) VALUES (?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `newrelic_application` (`newrelic_application`.`id`,`newrelic_application`.`checksum`,`newrelic_application`.`customer_id`,`newrelic_application`.`ext_id`,`newrelic_application`.`name`,`newrelic_application`.`health_status`,`newrelic_application`.`reporting`) VALUES (?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`customer_id`=VALUES(`customer_id`),`ext_id`=VALUES(`ext_id`),`name`=VALUES(`name`),`health_status`=VALUES(`health_status`),`reporting`=VALUES(`reporting`)"
	}
	r, err := tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLInt64(t.ExtID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.HealthStatus),
		orm.ToSQLBool(t.Reporting),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBFindOne will find a NewrelicApplication record in the database with the primary key
func (t *NewrelicApplication) DBFindOne(ctx context.Context, db *sql.DB, value string) (bool, error) {
	q := "SELECT `newrelic_application`.`id`,`newrelic_application`.`checksum`,`newrelic_application`.`customer_id`,`newrelic_application`.`ext_id`,`newrelic_application`.`name`,`newrelic_application`.`health_status`,`newrelic_application`.`reporting` FROM `newrelic_application` WHERE `id` = ? LIMIT 1"
	row := db.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CustomerID sql.NullString
	var _ExtID sql.NullInt64
	var _Name sql.NullString
	var _HealthStatus sql.NullString
	var _Reporting sql.NullBool
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_ExtID,
		&_Name,
		&_HealthStatus,
		&_Reporting,
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
	if _HealthStatus.Valid {
		t.SetHealthStatus(_HealthStatus.String)
	}
	if _Reporting.Valid {
		t.SetReporting(_Reporting.Bool)
	}
	return true, nil
}

// DBFindOneTx will find a NewrelicApplication record in the database with the primary key using the provided transaction
func (t *NewrelicApplication) DBFindOneTx(ctx context.Context, tx *sql.Tx, value string) (bool, error) {
	q := "SELECT `newrelic_application`.`id`,`newrelic_application`.`checksum`,`newrelic_application`.`customer_id`,`newrelic_application`.`ext_id`,`newrelic_application`.`name`,`newrelic_application`.`health_status`,`newrelic_application`.`reporting` FROM `newrelic_application` WHERE `id` = ? LIMIT 1"
	row := tx.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CustomerID sql.NullString
	var _ExtID sql.NullInt64
	var _Name sql.NullString
	var _HealthStatus sql.NullString
	var _Reporting sql.NullBool
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_ExtID,
		&_Name,
		&_HealthStatus,
		&_Reporting,
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
	if _HealthStatus.Valid {
		t.SetHealthStatus(_HealthStatus.String)
	}
	if _Reporting.Valid {
		t.SetReporting(_Reporting.Bool)
	}
	return true, nil
}

// FindNewrelicApplications will find a NewrelicApplication record in the database with the provided parameters
func FindNewrelicApplications(ctx context.Context, db *sql.DB, _params ...interface{}) ([]*NewrelicApplication, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("customer_id"),
		orm.Column("ext_id"),
		orm.Column("name"),
		orm.Column("health_status"),
		orm.Column("reporting"),
		orm.Table(NewrelicApplicationTableName),
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
	results := make([]*NewrelicApplication, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _ExtID sql.NullInt64
		var _Name sql.NullString
		var _HealthStatus sql.NullString
		var _Reporting sql.NullBool
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_ExtID,
			&_Name,
			&_HealthStatus,
			&_Reporting,
		)
		if err != nil {
			return nil, err
		}
		t := &NewrelicApplication{}
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
		if _HealthStatus.Valid {
			t.SetHealthStatus(_HealthStatus.String)
		}
		if _Reporting.Valid {
			t.SetReporting(_Reporting.Bool)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindNewrelicApplicationsTx will find a NewrelicApplication record in the database with the provided parameters using the provided transaction
func FindNewrelicApplicationsTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) ([]*NewrelicApplication, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("customer_id"),
		orm.Column("ext_id"),
		orm.Column("name"),
		orm.Column("health_status"),
		orm.Column("reporting"),
		orm.Table(NewrelicApplicationTableName),
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
	results := make([]*NewrelicApplication, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _ExtID sql.NullInt64
		var _Name sql.NullString
		var _HealthStatus sql.NullString
		var _Reporting sql.NullBool
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_ExtID,
			&_Name,
			&_HealthStatus,
			&_Reporting,
		)
		if err != nil {
			return nil, err
		}
		t := &NewrelicApplication{}
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
		if _HealthStatus.Valid {
			t.SetHealthStatus(_HealthStatus.String)
		}
		if _Reporting.Valid {
			t.SetReporting(_Reporting.Bool)
		}
		results = append(results, t)
	}
	return results, nil
}

// DBFind will find a NewrelicApplication record in the database with the provided parameters
func (t *NewrelicApplication) DBFind(ctx context.Context, db *sql.DB, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("customer_id"),
		orm.Column("ext_id"),
		orm.Column("name"),
		orm.Column("health_status"),
		orm.Column("reporting"),
		orm.Table(NewrelicApplicationTableName),
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
	var _HealthStatus sql.NullString
	var _Reporting sql.NullBool
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_ExtID,
		&_Name,
		&_HealthStatus,
		&_Reporting,
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
	if _HealthStatus.Valid {
		t.SetHealthStatus(_HealthStatus.String)
	}
	if _Reporting.Valid {
		t.SetReporting(_Reporting.Bool)
	}
	return true, nil
}

// DBFindTx will find a NewrelicApplication record in the database with the provided parameters using the provided transaction
func (t *NewrelicApplication) DBFindTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("customer_id"),
		orm.Column("ext_id"),
		orm.Column("name"),
		orm.Column("health_status"),
		orm.Column("reporting"),
		orm.Table(NewrelicApplicationTableName),
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
	var _HealthStatus sql.NullString
	var _Reporting sql.NullBool
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_ExtID,
		&_Name,
		&_HealthStatus,
		&_Reporting,
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
	if _HealthStatus.Valid {
		t.SetHealthStatus(_HealthStatus.String)
	}
	if _Reporting.Valid {
		t.SetReporting(_Reporting.Bool)
	}
	return true, nil
}

// CountNewrelicApplications will find the count of NewrelicApplication records in the database
func CountNewrelicApplications(ctx context.Context, db *sql.DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(NewrelicApplicationTableName),
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

// CountNewrelicApplicationsTx will find the count of NewrelicApplication records in the database using the provided transaction
func CountNewrelicApplicationsTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(NewrelicApplicationTableName),
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

// DBCount will find the count of NewrelicApplication records in the database
func (t *NewrelicApplication) DBCount(ctx context.Context, db *sql.DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(NewrelicApplicationTableName),
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

// DBCountTx will find the count of NewrelicApplication records in the database using the provided transaction
func (t *NewrelicApplication) DBCountTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(NewrelicApplicationTableName),
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

// DBExists will return true if the NewrelicApplication record exists in the database
func (t *NewrelicApplication) DBExists(ctx context.Context, db *sql.DB) (bool, error) {
	q := "SELECT `id` FROM `newrelic_application` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := db.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// DBExistsTx will return true if the NewrelicApplication record exists in the database using the provided transaction
func (t *NewrelicApplication) DBExistsTx(ctx context.Context, tx *sql.Tx) (bool, error) {
	q := "SELECT `id` FROM `newrelic_application` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := tx.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// PrimaryKeyColumn returns the column name for the primary key
func (t *NewrelicApplication) PrimaryKeyColumn() string {
	return NewrelicApplicationColumnID
}

// PrimaryKeyColumnType returns the primary key column Go type as a string
func (t *NewrelicApplication) PrimaryKeyColumnType() string {
	return "string"
}

// PrimaryKey returns the primary key column value
func (t *NewrelicApplication) PrimaryKey() interface{} {
	return t.ID
}
