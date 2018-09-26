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

var _ Model = (*NewrelicMetric)(nil)
var _ CSVWriter = (*NewrelicMetric)(nil)
var _ JSONWriter = (*NewrelicMetric)(nil)
var _ Checksum = (*NewrelicMetric)(nil)

// NewrelicMetricTableName is the name of the table in SQL
const NewrelicMetricTableName = "newrelic_metric"

var NewrelicMetricColumns = []string{
	"id",
	"checksum",
	"customer_id",
	"application_ext_id",
	"application_id",
	"name",
	"value_name",
	"start",
	"duration_sec",
	"value",
}

// NewrelicMetric table
type NewrelicMetric struct {
	ApplicationExtID int64   `json:"application_ext_id"`
	ApplicationID    string  `json:"application_id"`
	Checksum         *string `json:"checksum,omitempty"`
	CustomerID       string  `json:"customer_id"`
	DurationSec      int64   `json:"duration_sec"`
	ID               string  `json:"id"`
	Name             string  `json:"name"`
	Start            int64   `json:"start"`
	Value            float64 `json:"value"`
	ValueName        string  `json:"value_name"`
}

// TableName returns the SQL table name for NewrelicMetric and satifies the Model interface
func (t *NewrelicMetric) TableName() string {
	return NewrelicMetricTableName
}

// ToCSV will serialize the NewrelicMetric instance to a CSV compatible array of strings
func (t *NewrelicMetric) ToCSV() []string {
	return []string{
		t.ID,
		t.CalculateChecksum(),
		t.CustomerID,
		toCSVString(t.ApplicationExtID),
		t.ApplicationID,
		t.Name,
		t.ValueName,
		toCSVString(t.Start),
		toCSVString(t.DurationSec),
		toCSVString(t.Value),
	}
}

// WriteCSV will serialize the NewrelicMetric instance to the writer as CSV and satisfies the CSVWriter interface
func (t *NewrelicMetric) WriteCSV(w *csv.Writer) error {
	return w.Write(t.ToCSV())
}

// WriteJSON will serialize the NewrelicMetric instance to the writer as JSON and satisfies the JSONWriter interface
func (t *NewrelicMetric) WriteJSON(w io.Writer, indent ...bool) error {
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

// NewNewrelicMetricReader creates a JSON reader which can read in NewrelicMetric objects serialized as JSON either as an array, single object or json new lines
// and writes each NewrelicMetric to the channel provided
func NewNewrelicMetricReader(r io.Reader, ch chan<- NewrelicMetric) error {
	return orm.Deserialize(r, func(buf json.RawMessage) error {
		dec := json.NewDecoder(bytes.NewBuffer(buf))
		e := NewrelicMetric{}
		if err := dec.Decode(&e); err != nil {
			return err
		}
		ch <- e
		return nil
	})
}

// NewCSVNewrelicMetricReaderDir will read the reader as CSV and emit each record to the channel provided
func NewCSVNewrelicMetricReader(r io.Reader, ch chan<- NewrelicMetric) error {
	cr := csv.NewReader(r)
	for {
		record, err := cr.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		ch <- NewrelicMetric{
			ID:               record[0],
			Checksum:         fromStringPointer(record[1]),
			CustomerID:       record[2],
			ApplicationExtID: fromCSVInt64(record[3]),
			ApplicationID:    record[4],
			Name:             record[5],
			ValueName:        record[6],
			Start:            fromCSVInt64(record[7]),
			DurationSec:      fromCSVInt64(record[8]),
			Value:            fromCSVFloat64(record[9]),
		}
	}
	return nil
}

// NewCSVNewrelicMetricReaderFile will read the file as a CSV and emit each record to the channel provided
func NewCSVNewrelicMetricReaderFile(fp string, ch chan<- NewrelicMetric) error {
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
	return NewCSVNewrelicMetricReader(fc, ch)
}

// NewCSVNewrelicMetricReaderDir will read the newrelic_metric.csv.gz file as a CSV and emit each record to the channel provided
func NewCSVNewrelicMetricReaderDir(dir string, ch chan<- NewrelicMetric) error {
	return NewCSVNewrelicMetricReaderFile(filepath.Join(dir, "newrelic_metric.csv.gz"), ch)
}

// NewrelicMetricCSVDeduper is a function callback which takes the existing value (a) and the new value (b)
// and the return value should be the one to use (or a new one, if applicable). return nil
// to skip processing of this record
type NewrelicMetricCSVDeduper func(a NewrelicMetric, b NewrelicMetric) *NewrelicMetric

// NewrelicMetricCSVDedupeDisabled is set on whether the CSV writer should de-dupe values by key
var NewrelicMetricCSVDedupeDisabled bool

// NewNewrelicMetricCSVWriterSize creates a batch writer that will write each NewrelicMetric into a CSV file
// this method will automatically de-duplicate entries using the primary key. if the checksum
// for a newer item with the same primary key doesn't match a previously sent item, the newer
// one will be used
func NewNewrelicMetricCSVWriterSize(w io.Writer, size int, dedupers ...NewrelicMetricCSVDeduper) (chan NewrelicMetric, chan bool, error) {
	cw := csv.NewWriter(w)
	ch := make(chan NewrelicMetric, size)
	done := make(chan bool)
	go func() {
		defer func() { done <- true }()
		dodedupe := !NewrelicMetricCSVDedupeDisabled
		var kv map[string]*NewrelicMetric
		var deduper NewrelicMetricCSVDeduper
		if dedupers != nil && len(dedupers) > 0 {
			deduper = dedupers[0]
			dodedupe = true
		}
		if dodedupe {
			kv = make(map[string]*NewrelicMetric)
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

// NewrelicMetricCSVDefaultSize is the default channel buffer size if not provided
var NewrelicMetricCSVDefaultSize = 100

// NewNewrelicMetricCSVWriter creates a batch writer that will write each NewrelicMetric into a CSV file
func NewNewrelicMetricCSVWriter(w io.Writer, dedupers ...NewrelicMetricCSVDeduper) (chan NewrelicMetric, chan bool, error) {
	return NewNewrelicMetricCSVWriterSize(w, NewrelicMetricCSVDefaultSize, dedupers...)
}

// NewNewrelicMetricCSVWriterDir creates a batch writer that will write each NewrelicMetric into a CSV file named newrelic_metric.csv.gz in dir
func NewNewrelicMetricCSVWriterDir(dir string, dedupers ...NewrelicMetricCSVDeduper) (chan NewrelicMetric, chan bool, error) {
	return NewNewrelicMetricCSVWriterFile(filepath.Join(dir, "newrelic_metric.csv.gz"), dedupers...)
}

// NewNewrelicMetricCSVWriterFile creates a batch writer that will write each NewrelicMetric into a CSV file
func NewNewrelicMetricCSVWriterFile(fn string, dedupers ...NewrelicMetricCSVDeduper) (chan NewrelicMetric, chan bool, error) {
	f, err := os.Create(fn)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening CSV file %s. %v", fn, err)
	}
	var fc io.WriteCloser = f
	if filepath.Ext(fn) == ".gz" {
		w, _ := gzip.NewWriterLevel(f, gzip.BestCompression)
		fc = w
	}
	ch, done, err := NewNewrelicMetricCSVWriter(fc, dedupers...)
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

type NewrelicMetricDBAction func(ctx context.Context, db *sql.DB, record NewrelicMetric) error

// NewNewrelicMetricDBWriterSize creates a DB writer that will write each issue into the DB
func NewNewrelicMetricDBWriterSize(ctx context.Context, db *sql.DB, errors chan<- error, size int, actions ...NewrelicMetricDBAction) (chan NewrelicMetric, chan bool, error) {
	ch := make(chan NewrelicMetric, size)
	done := make(chan bool)
	var action NewrelicMetricDBAction
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

// NewNewrelicMetricDBWriter creates a DB writer that will write each issue into the DB
func NewNewrelicMetricDBWriter(ctx context.Context, db *sql.DB, errors chan<- error, actions ...NewrelicMetricDBAction) (chan NewrelicMetric, chan bool, error) {
	return NewNewrelicMetricDBWriterSize(ctx, db, errors, 100, actions...)
}

// NewrelicMetricColumnID is the ID SQL column name for the NewrelicMetric table
const NewrelicMetricColumnID = "id"

// NewrelicMetricEscapedColumnID is the escaped ID SQL column name for the NewrelicMetric table
const NewrelicMetricEscapedColumnID = "`id`"

// NewrelicMetricColumnChecksum is the Checksum SQL column name for the NewrelicMetric table
const NewrelicMetricColumnChecksum = "checksum"

// NewrelicMetricEscapedColumnChecksum is the escaped Checksum SQL column name for the NewrelicMetric table
const NewrelicMetricEscapedColumnChecksum = "`checksum`"

// NewrelicMetricColumnCustomerID is the CustomerID SQL column name for the NewrelicMetric table
const NewrelicMetricColumnCustomerID = "customer_id"

// NewrelicMetricEscapedColumnCustomerID is the escaped CustomerID SQL column name for the NewrelicMetric table
const NewrelicMetricEscapedColumnCustomerID = "`customer_id`"

// NewrelicMetricColumnApplicationExtID is the ApplicationExtID SQL column name for the NewrelicMetric table
const NewrelicMetricColumnApplicationExtID = "application_ext_id"

// NewrelicMetricEscapedColumnApplicationExtID is the escaped ApplicationExtID SQL column name for the NewrelicMetric table
const NewrelicMetricEscapedColumnApplicationExtID = "`application_ext_id`"

// NewrelicMetricColumnApplicationID is the ApplicationID SQL column name for the NewrelicMetric table
const NewrelicMetricColumnApplicationID = "application_id"

// NewrelicMetricEscapedColumnApplicationID is the escaped ApplicationID SQL column name for the NewrelicMetric table
const NewrelicMetricEscapedColumnApplicationID = "`application_id`"

// NewrelicMetricColumnName is the Name SQL column name for the NewrelicMetric table
const NewrelicMetricColumnName = "name"

// NewrelicMetricEscapedColumnName is the escaped Name SQL column name for the NewrelicMetric table
const NewrelicMetricEscapedColumnName = "`name`"

// NewrelicMetricColumnValueName is the ValueName SQL column name for the NewrelicMetric table
const NewrelicMetricColumnValueName = "value_name"

// NewrelicMetricEscapedColumnValueName is the escaped ValueName SQL column name for the NewrelicMetric table
const NewrelicMetricEscapedColumnValueName = "`value_name`"

// NewrelicMetricColumnStart is the Start SQL column name for the NewrelicMetric table
const NewrelicMetricColumnStart = "start"

// NewrelicMetricEscapedColumnStart is the escaped Start SQL column name for the NewrelicMetric table
const NewrelicMetricEscapedColumnStart = "`start`"

// NewrelicMetricColumnDurationSec is the DurationSec SQL column name for the NewrelicMetric table
const NewrelicMetricColumnDurationSec = "duration_sec"

// NewrelicMetricEscapedColumnDurationSec is the escaped DurationSec SQL column name for the NewrelicMetric table
const NewrelicMetricEscapedColumnDurationSec = "`duration_sec`"

// NewrelicMetricColumnValue is the Value SQL column name for the NewrelicMetric table
const NewrelicMetricColumnValue = "value"

// NewrelicMetricEscapedColumnValue is the escaped Value SQL column name for the NewrelicMetric table
const NewrelicMetricEscapedColumnValue = "`value`"

// GetID will return the NewrelicMetric ID value
func (t *NewrelicMetric) GetID() string {
	return t.ID
}

// SetID will set the NewrelicMetric ID value
func (t *NewrelicMetric) SetID(v string) {
	t.ID = v
}

// FindNewrelicMetricByID will find a NewrelicMetric by ID
func FindNewrelicMetricByID(ctx context.Context, db *sql.DB, value string) (*NewrelicMetric, error) {
	q := "SELECT `newrelic_metric`.`id`,`newrelic_metric`.`checksum`,`newrelic_metric`.`customer_id`,`newrelic_metric`.`application_ext_id`,`newrelic_metric`.`application_id`,`newrelic_metric`.`name`,`newrelic_metric`.`value_name`,`newrelic_metric`.`start`,`newrelic_metric`.`duration_sec`,`newrelic_metric`.`value` FROM `newrelic_metric` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CustomerID sql.NullString
	var _ApplicationExtID sql.NullInt64
	var _ApplicationID sql.NullString
	var _Name sql.NullString
	var _ValueName sql.NullString
	var _Start sql.NullInt64
	var _DurationSec sql.NullInt64
	var _Value sql.NullFloat64
	err := db.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_ApplicationExtID,
		&_ApplicationID,
		&_Name,
		&_ValueName,
		&_Start,
		&_DurationSec,
		&_Value,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &NewrelicMetric{}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _Checksum.Valid {
		t.SetChecksum(_Checksum.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _ApplicationExtID.Valid {
		t.SetApplicationExtID(_ApplicationExtID.Int64)
	}
	if _ApplicationID.Valid {
		t.SetApplicationID(_ApplicationID.String)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _ValueName.Valid {
		t.SetValueName(_ValueName.String)
	}
	if _Start.Valid {
		t.SetStart(_Start.Int64)
	}
	if _DurationSec.Valid {
		t.SetDurationSec(_DurationSec.Int64)
	}
	if _Value.Valid {
		t.SetValue(_Value.Float64)
	}
	return t, nil
}

// FindNewrelicMetricByIDTx will find a NewrelicMetric by ID using the provided transaction
func FindNewrelicMetricByIDTx(ctx context.Context, tx *sql.Tx, value string) (*NewrelicMetric, error) {
	q := "SELECT `newrelic_metric`.`id`,`newrelic_metric`.`checksum`,`newrelic_metric`.`customer_id`,`newrelic_metric`.`application_ext_id`,`newrelic_metric`.`application_id`,`newrelic_metric`.`name`,`newrelic_metric`.`value_name`,`newrelic_metric`.`start`,`newrelic_metric`.`duration_sec`,`newrelic_metric`.`value` FROM `newrelic_metric` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CustomerID sql.NullString
	var _ApplicationExtID sql.NullInt64
	var _ApplicationID sql.NullString
	var _Name sql.NullString
	var _ValueName sql.NullString
	var _Start sql.NullInt64
	var _DurationSec sql.NullInt64
	var _Value sql.NullFloat64
	err := tx.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_ApplicationExtID,
		&_ApplicationID,
		&_Name,
		&_ValueName,
		&_Start,
		&_DurationSec,
		&_Value,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &NewrelicMetric{}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _Checksum.Valid {
		t.SetChecksum(_Checksum.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _ApplicationExtID.Valid {
		t.SetApplicationExtID(_ApplicationExtID.Int64)
	}
	if _ApplicationID.Valid {
		t.SetApplicationID(_ApplicationID.String)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _ValueName.Valid {
		t.SetValueName(_ValueName.String)
	}
	if _Start.Valid {
		t.SetStart(_Start.Int64)
	}
	if _DurationSec.Valid {
		t.SetDurationSec(_DurationSec.Int64)
	}
	if _Value.Valid {
		t.SetValue(_Value.Float64)
	}
	return t, nil
}

// GetChecksum will return the NewrelicMetric Checksum value
func (t *NewrelicMetric) GetChecksum() string {
	if t.Checksum == nil {
		return ""
	}
	return *t.Checksum
}

// SetChecksum will set the NewrelicMetric Checksum value
func (t *NewrelicMetric) SetChecksum(v string) {
	t.Checksum = &v
}

// GetCustomerID will return the NewrelicMetric CustomerID value
func (t *NewrelicMetric) GetCustomerID() string {
	return t.CustomerID
}

// SetCustomerID will set the NewrelicMetric CustomerID value
func (t *NewrelicMetric) SetCustomerID(v string) {
	t.CustomerID = v
}

// FindNewrelicMetricsByCustomerID will find all NewrelicMetrics by the CustomerID value
func FindNewrelicMetricsByCustomerID(ctx context.Context, db *sql.DB, value string) ([]*NewrelicMetric, error) {
	q := "SELECT `newrelic_metric`.`id`,`newrelic_metric`.`checksum`,`newrelic_metric`.`customer_id`,`newrelic_metric`.`application_ext_id`,`newrelic_metric`.`application_id`,`newrelic_metric`.`name`,`newrelic_metric`.`value_name`,`newrelic_metric`.`start`,`newrelic_metric`.`duration_sec`,`newrelic_metric`.`value` FROM `newrelic_metric` WHERE `customer_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*NewrelicMetric, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _ApplicationExtID sql.NullInt64
		var _ApplicationID sql.NullString
		var _Name sql.NullString
		var _ValueName sql.NullString
		var _Start sql.NullInt64
		var _DurationSec sql.NullInt64
		var _Value sql.NullFloat64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_ApplicationExtID,
			&_ApplicationID,
			&_Name,
			&_ValueName,
			&_Start,
			&_DurationSec,
			&_Value,
		)
		if err != nil {
			return nil, err
		}
		t := &NewrelicMetric{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _ApplicationExtID.Valid {
			t.SetApplicationExtID(_ApplicationExtID.Int64)
		}
		if _ApplicationID.Valid {
			t.SetApplicationID(_ApplicationID.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _ValueName.Valid {
			t.SetValueName(_ValueName.String)
		}
		if _Start.Valid {
			t.SetStart(_Start.Int64)
		}
		if _DurationSec.Valid {
			t.SetDurationSec(_DurationSec.Int64)
		}
		if _Value.Valid {
			t.SetValue(_Value.Float64)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindNewrelicMetricsByCustomerIDTx will find all NewrelicMetrics by the CustomerID value using the provided transaction
func FindNewrelicMetricsByCustomerIDTx(ctx context.Context, tx *sql.Tx, value string) ([]*NewrelicMetric, error) {
	q := "SELECT `newrelic_metric`.`id`,`newrelic_metric`.`checksum`,`newrelic_metric`.`customer_id`,`newrelic_metric`.`application_ext_id`,`newrelic_metric`.`application_id`,`newrelic_metric`.`name`,`newrelic_metric`.`value_name`,`newrelic_metric`.`start`,`newrelic_metric`.`duration_sec`,`newrelic_metric`.`value` FROM `newrelic_metric` WHERE `customer_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*NewrelicMetric, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _ApplicationExtID sql.NullInt64
		var _ApplicationID sql.NullString
		var _Name sql.NullString
		var _ValueName sql.NullString
		var _Start sql.NullInt64
		var _DurationSec sql.NullInt64
		var _Value sql.NullFloat64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_ApplicationExtID,
			&_ApplicationID,
			&_Name,
			&_ValueName,
			&_Start,
			&_DurationSec,
			&_Value,
		)
		if err != nil {
			return nil, err
		}
		t := &NewrelicMetric{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _ApplicationExtID.Valid {
			t.SetApplicationExtID(_ApplicationExtID.Int64)
		}
		if _ApplicationID.Valid {
			t.SetApplicationID(_ApplicationID.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _ValueName.Valid {
			t.SetValueName(_ValueName.String)
		}
		if _Start.Valid {
			t.SetStart(_Start.Int64)
		}
		if _DurationSec.Valid {
			t.SetDurationSec(_DurationSec.Int64)
		}
		if _Value.Valid {
			t.SetValue(_Value.Float64)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetApplicationExtID will return the NewrelicMetric ApplicationExtID value
func (t *NewrelicMetric) GetApplicationExtID() int64 {
	return t.ApplicationExtID
}

// SetApplicationExtID will set the NewrelicMetric ApplicationExtID value
func (t *NewrelicMetric) SetApplicationExtID(v int64) {
	t.ApplicationExtID = v
}

// FindNewrelicMetricsByApplicationExtID will find all NewrelicMetrics by the ApplicationExtID value
func FindNewrelicMetricsByApplicationExtID(ctx context.Context, db *sql.DB, value int64) ([]*NewrelicMetric, error) {
	q := "SELECT `newrelic_metric`.`id`,`newrelic_metric`.`checksum`,`newrelic_metric`.`customer_id`,`newrelic_metric`.`application_ext_id`,`newrelic_metric`.`application_id`,`newrelic_metric`.`name`,`newrelic_metric`.`value_name`,`newrelic_metric`.`start`,`newrelic_metric`.`duration_sec`,`newrelic_metric`.`value` FROM `newrelic_metric` WHERE `application_ext_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*NewrelicMetric, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _ApplicationExtID sql.NullInt64
		var _ApplicationID sql.NullString
		var _Name sql.NullString
		var _ValueName sql.NullString
		var _Start sql.NullInt64
		var _DurationSec sql.NullInt64
		var _Value sql.NullFloat64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_ApplicationExtID,
			&_ApplicationID,
			&_Name,
			&_ValueName,
			&_Start,
			&_DurationSec,
			&_Value,
		)
		if err != nil {
			return nil, err
		}
		t := &NewrelicMetric{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _ApplicationExtID.Valid {
			t.SetApplicationExtID(_ApplicationExtID.Int64)
		}
		if _ApplicationID.Valid {
			t.SetApplicationID(_ApplicationID.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _ValueName.Valid {
			t.SetValueName(_ValueName.String)
		}
		if _Start.Valid {
			t.SetStart(_Start.Int64)
		}
		if _DurationSec.Valid {
			t.SetDurationSec(_DurationSec.Int64)
		}
		if _Value.Valid {
			t.SetValue(_Value.Float64)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindNewrelicMetricsByApplicationExtIDTx will find all NewrelicMetrics by the ApplicationExtID value using the provided transaction
func FindNewrelicMetricsByApplicationExtIDTx(ctx context.Context, tx *sql.Tx, value int64) ([]*NewrelicMetric, error) {
	q := "SELECT `newrelic_metric`.`id`,`newrelic_metric`.`checksum`,`newrelic_metric`.`customer_id`,`newrelic_metric`.`application_ext_id`,`newrelic_metric`.`application_id`,`newrelic_metric`.`name`,`newrelic_metric`.`value_name`,`newrelic_metric`.`start`,`newrelic_metric`.`duration_sec`,`newrelic_metric`.`value` FROM `newrelic_metric` WHERE `application_ext_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*NewrelicMetric, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _ApplicationExtID sql.NullInt64
		var _ApplicationID sql.NullString
		var _Name sql.NullString
		var _ValueName sql.NullString
		var _Start sql.NullInt64
		var _DurationSec sql.NullInt64
		var _Value sql.NullFloat64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_ApplicationExtID,
			&_ApplicationID,
			&_Name,
			&_ValueName,
			&_Start,
			&_DurationSec,
			&_Value,
		)
		if err != nil {
			return nil, err
		}
		t := &NewrelicMetric{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _ApplicationExtID.Valid {
			t.SetApplicationExtID(_ApplicationExtID.Int64)
		}
		if _ApplicationID.Valid {
			t.SetApplicationID(_ApplicationID.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _ValueName.Valid {
			t.SetValueName(_ValueName.String)
		}
		if _Start.Valid {
			t.SetStart(_Start.Int64)
		}
		if _DurationSec.Valid {
			t.SetDurationSec(_DurationSec.Int64)
		}
		if _Value.Valid {
			t.SetValue(_Value.Float64)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetApplicationID will return the NewrelicMetric ApplicationID value
func (t *NewrelicMetric) GetApplicationID() string {
	return t.ApplicationID
}

// SetApplicationID will set the NewrelicMetric ApplicationID value
func (t *NewrelicMetric) SetApplicationID(v string) {
	t.ApplicationID = v
}

// FindNewrelicMetricsByApplicationID will find all NewrelicMetrics by the ApplicationID value
func FindNewrelicMetricsByApplicationID(ctx context.Context, db *sql.DB, value string) ([]*NewrelicMetric, error) {
	q := "SELECT `newrelic_metric`.`id`,`newrelic_metric`.`checksum`,`newrelic_metric`.`customer_id`,`newrelic_metric`.`application_ext_id`,`newrelic_metric`.`application_id`,`newrelic_metric`.`name`,`newrelic_metric`.`value_name`,`newrelic_metric`.`start`,`newrelic_metric`.`duration_sec`,`newrelic_metric`.`value` FROM `newrelic_metric` WHERE `application_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*NewrelicMetric, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _ApplicationExtID sql.NullInt64
		var _ApplicationID sql.NullString
		var _Name sql.NullString
		var _ValueName sql.NullString
		var _Start sql.NullInt64
		var _DurationSec sql.NullInt64
		var _Value sql.NullFloat64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_ApplicationExtID,
			&_ApplicationID,
			&_Name,
			&_ValueName,
			&_Start,
			&_DurationSec,
			&_Value,
		)
		if err != nil {
			return nil, err
		}
		t := &NewrelicMetric{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _ApplicationExtID.Valid {
			t.SetApplicationExtID(_ApplicationExtID.Int64)
		}
		if _ApplicationID.Valid {
			t.SetApplicationID(_ApplicationID.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _ValueName.Valid {
			t.SetValueName(_ValueName.String)
		}
		if _Start.Valid {
			t.SetStart(_Start.Int64)
		}
		if _DurationSec.Valid {
			t.SetDurationSec(_DurationSec.Int64)
		}
		if _Value.Valid {
			t.SetValue(_Value.Float64)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindNewrelicMetricsByApplicationIDTx will find all NewrelicMetrics by the ApplicationID value using the provided transaction
func FindNewrelicMetricsByApplicationIDTx(ctx context.Context, tx *sql.Tx, value string) ([]*NewrelicMetric, error) {
	q := "SELECT `newrelic_metric`.`id`,`newrelic_metric`.`checksum`,`newrelic_metric`.`customer_id`,`newrelic_metric`.`application_ext_id`,`newrelic_metric`.`application_id`,`newrelic_metric`.`name`,`newrelic_metric`.`value_name`,`newrelic_metric`.`start`,`newrelic_metric`.`duration_sec`,`newrelic_metric`.`value` FROM `newrelic_metric` WHERE `application_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*NewrelicMetric, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _ApplicationExtID sql.NullInt64
		var _ApplicationID sql.NullString
		var _Name sql.NullString
		var _ValueName sql.NullString
		var _Start sql.NullInt64
		var _DurationSec sql.NullInt64
		var _Value sql.NullFloat64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_ApplicationExtID,
			&_ApplicationID,
			&_Name,
			&_ValueName,
			&_Start,
			&_DurationSec,
			&_Value,
		)
		if err != nil {
			return nil, err
		}
		t := &NewrelicMetric{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _ApplicationExtID.Valid {
			t.SetApplicationExtID(_ApplicationExtID.Int64)
		}
		if _ApplicationID.Valid {
			t.SetApplicationID(_ApplicationID.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _ValueName.Valid {
			t.SetValueName(_ValueName.String)
		}
		if _Start.Valid {
			t.SetStart(_Start.Int64)
		}
		if _DurationSec.Valid {
			t.SetDurationSec(_DurationSec.Int64)
		}
		if _Value.Valid {
			t.SetValue(_Value.Float64)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetName will return the NewrelicMetric Name value
func (t *NewrelicMetric) GetName() string {
	return t.Name
}

// SetName will set the NewrelicMetric Name value
func (t *NewrelicMetric) SetName(v string) {
	t.Name = v
}

// GetValueName will return the NewrelicMetric ValueName value
func (t *NewrelicMetric) GetValueName() string {
	return t.ValueName
}

// SetValueName will set the NewrelicMetric ValueName value
func (t *NewrelicMetric) SetValueName(v string) {
	t.ValueName = v
}

// GetStart will return the NewrelicMetric Start value
func (t *NewrelicMetric) GetStart() int64 {
	return t.Start
}

// SetStart will set the NewrelicMetric Start value
func (t *NewrelicMetric) SetStart(v int64) {
	t.Start = v
}

// GetDurationSec will return the NewrelicMetric DurationSec value
func (t *NewrelicMetric) GetDurationSec() int64 {
	return t.DurationSec
}

// SetDurationSec will set the NewrelicMetric DurationSec value
func (t *NewrelicMetric) SetDurationSec(v int64) {
	t.DurationSec = v
}

// GetValue will return the NewrelicMetric Value value
func (t *NewrelicMetric) GetValue() float64 {
	return t.Value
}

// SetValue will set the NewrelicMetric Value value
func (t *NewrelicMetric) SetValue(v float64) {
	t.Value = v
}

func (t *NewrelicMetric) toTimestamp(value time.Time) *timestamp.Timestamp {
	ts, _ := ptypes.TimestampProto(value)
	return ts
}

// DBCreateNewrelicMetricTable will create the NewrelicMetric table
func DBCreateNewrelicMetricTable(ctx context.Context, db *sql.DB) error {
	q := "CREATE TABLE `newrelic_metric` (`id` VARCHAR(64) NOT NULL PRIMARY KEY,`checksum` CHAR(64),`customer_id` VARCHAR(64) NOT NULL,`application_ext_id` BIGINT(20) NOT NULL,`application_id` VARCHAR(64) NOT NULL,`name`TEXT NOT NULL,`value_name`TEXT NOT NULL,`start` BIGINT(20) NOT NULL,`duration_sec` BIGINT(20) NOT NULL,`value` DOUBLE NOT NULL,INDEX newrelic_metric_customer_id_index (`customer_id`),INDEX newrelic_metric_application_ext_id_index (`application_ext_id`),INDEX newrelic_metric_application_id_index (`application_id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBCreateNewrelicMetricTableTx will create the NewrelicMetric table using the provided transction
func DBCreateNewrelicMetricTableTx(ctx context.Context, tx *sql.Tx) error {
	q := "CREATE TABLE `newrelic_metric` (`id` VARCHAR(64) NOT NULL PRIMARY KEY,`checksum` CHAR(64),`customer_id` VARCHAR(64) NOT NULL,`application_ext_id` BIGINT(20) NOT NULL,`application_id` VARCHAR(64) NOT NULL,`name`TEXT NOT NULL,`value_name`TEXT NOT NULL,`start` BIGINT(20) NOT NULL,`duration_sec` BIGINT(20) NOT NULL,`value` DOUBLE NOT NULL,INDEX newrelic_metric_customer_id_index (`customer_id`),INDEX newrelic_metric_application_ext_id_index (`application_ext_id`),INDEX newrelic_metric_application_id_index (`application_id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// DBDropNewrelicMetricTable will drop the NewrelicMetric table
func DBDropNewrelicMetricTable(ctx context.Context, db *sql.DB) error {
	q := "DROP TABLE IF EXISTS `newrelic_metric`"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBDropNewrelicMetricTableTx will drop the NewrelicMetric table using the provided transaction
func DBDropNewrelicMetricTableTx(ctx context.Context, tx *sql.Tx) error {
	q := "DROP TABLE IF EXISTS `newrelic_metric`"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// CalculateChecksum will calculate a checksum of the SHA1 of all field values
func (t *NewrelicMetric) CalculateChecksum() string {
	return orm.HashStrings(
		orm.ToString(t.ID),
		orm.ToString(t.CustomerID),
		orm.ToString(t.ApplicationExtID),
		orm.ToString(t.ApplicationID),
		orm.ToString(t.Name),
		orm.ToString(t.ValueName),
		orm.ToString(t.Start),
		orm.ToString(t.DurationSec),
		orm.ToString(t.Value),
	)
}

// DBCreate will create a new NewrelicMetric record in the database
func (t *NewrelicMetric) DBCreate(ctx context.Context, db *sql.DB) (sql.Result, error) {
	q := "INSERT INTO `newrelic_metric` (`newrelic_metric`.`id`,`newrelic_metric`.`checksum`,`newrelic_metric`.`customer_id`,`newrelic_metric`.`application_ext_id`,`newrelic_metric`.`application_id`,`newrelic_metric`.`name`,`newrelic_metric`.`value_name`,`newrelic_metric`.`start`,`newrelic_metric`.`duration_sec`,`newrelic_metric`.`value`) VALUES (?,?,?,?,?,?,?,?,?,?)"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLInt64(t.ApplicationExtID),
		orm.ToSQLString(t.ApplicationID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.ValueName),
		orm.ToSQLInt64(t.Start),
		orm.ToSQLInt64(t.DurationSec),
		orm.ToSQLFloat64(t.Value),
	)
}

// DBCreateTx will create a new NewrelicMetric record in the database using the provided transaction
func (t *NewrelicMetric) DBCreateTx(ctx context.Context, tx *sql.Tx) (sql.Result, error) {
	q := "INSERT INTO `newrelic_metric` (`newrelic_metric`.`id`,`newrelic_metric`.`checksum`,`newrelic_metric`.`customer_id`,`newrelic_metric`.`application_ext_id`,`newrelic_metric`.`application_id`,`newrelic_metric`.`name`,`newrelic_metric`.`value_name`,`newrelic_metric`.`start`,`newrelic_metric`.`duration_sec`,`newrelic_metric`.`value`) VALUES (?,?,?,?,?,?,?,?,?,?)"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLInt64(t.ApplicationExtID),
		orm.ToSQLString(t.ApplicationID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.ValueName),
		orm.ToSQLInt64(t.Start),
		orm.ToSQLInt64(t.DurationSec),
		orm.ToSQLFloat64(t.Value),
	)
}

// DBCreateIgnoreDuplicate will upsert the NewrelicMetric record in the database
func (t *NewrelicMetric) DBCreateIgnoreDuplicate(ctx context.Context, db *sql.DB) (sql.Result, error) {
	q := "INSERT INTO `newrelic_metric` (`newrelic_metric`.`id`,`newrelic_metric`.`checksum`,`newrelic_metric`.`customer_id`,`newrelic_metric`.`application_ext_id`,`newrelic_metric`.`application_id`,`newrelic_metric`.`name`,`newrelic_metric`.`value_name`,`newrelic_metric`.`start`,`newrelic_metric`.`duration_sec`,`newrelic_metric`.`value`) VALUES (?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLInt64(t.ApplicationExtID),
		orm.ToSQLString(t.ApplicationID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.ValueName),
		orm.ToSQLInt64(t.Start),
		orm.ToSQLInt64(t.DurationSec),
		orm.ToSQLFloat64(t.Value),
	)
}

// DBCreateIgnoreDuplicateTx will upsert the NewrelicMetric record in the database using the provided transaction
func (t *NewrelicMetric) DBCreateIgnoreDuplicateTx(ctx context.Context, tx *sql.Tx) (sql.Result, error) {
	q := "INSERT INTO `newrelic_metric` (`newrelic_metric`.`id`,`newrelic_metric`.`checksum`,`newrelic_metric`.`customer_id`,`newrelic_metric`.`application_ext_id`,`newrelic_metric`.`application_id`,`newrelic_metric`.`name`,`newrelic_metric`.`value_name`,`newrelic_metric`.`start`,`newrelic_metric`.`duration_sec`,`newrelic_metric`.`value`) VALUES (?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLInt64(t.ApplicationExtID),
		orm.ToSQLString(t.ApplicationID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.ValueName),
		orm.ToSQLInt64(t.Start),
		orm.ToSQLInt64(t.DurationSec),
		orm.ToSQLFloat64(t.Value),
	)
}

// DeleteAllNewrelicMetrics deletes all NewrelicMetric records in the database with optional filters
func DeleteAllNewrelicMetrics(ctx context.Context, db *sql.DB, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(NewrelicMetricTableName),
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

// DeleteAllNewrelicMetricsTx deletes all NewrelicMetric records in the database with optional filters using the provided transaction
func DeleteAllNewrelicMetricsTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(NewrelicMetricTableName),
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

// DBDelete will delete this NewrelicMetric record in the database
func (t *NewrelicMetric) DBDelete(ctx context.Context, db *sql.DB) (bool, error) {
	q := "DELETE FROM `newrelic_metric` WHERE `id` = ?"
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

// DBDeleteTx will delete this NewrelicMetric record in the database using the provided transaction
func (t *NewrelicMetric) DBDeleteTx(ctx context.Context, tx *sql.Tx) (bool, error) {
	q := "DELETE FROM `newrelic_metric` WHERE `id` = ?"
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

// DBUpdate will update the NewrelicMetric record in the database
func (t *NewrelicMetric) DBUpdate(ctx context.Context, db *sql.DB) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `newrelic_metric` SET `checksum`=?,`customer_id`=?,`application_ext_id`=?,`application_id`=?,`name`=?,`value_name`=?,`start`=?,`duration_sec`=?,`value`=? WHERE `id`=?"
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLInt64(t.ApplicationExtID),
		orm.ToSQLString(t.ApplicationID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.ValueName),
		orm.ToSQLInt64(t.Start),
		orm.ToSQLInt64(t.DurationSec),
		orm.ToSQLFloat64(t.Value),
		orm.ToSQLString(t.ID),
	)
}

// DBUpdateTx will update the NewrelicMetric record in the database using the provided transaction
func (t *NewrelicMetric) DBUpdateTx(ctx context.Context, tx *sql.Tx) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `newrelic_metric` SET `checksum`=?,`customer_id`=?,`application_ext_id`=?,`application_id`=?,`name`=?,`value_name`=?,`start`=?,`duration_sec`=?,`value`=? WHERE `id`=?"
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLInt64(t.ApplicationExtID),
		orm.ToSQLString(t.ApplicationID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.ValueName),
		orm.ToSQLInt64(t.Start),
		orm.ToSQLInt64(t.DurationSec),
		orm.ToSQLFloat64(t.Value),
		orm.ToSQLString(t.ID),
	)
}

// DBUpsert will upsert the NewrelicMetric record in the database
func (t *NewrelicMetric) DBUpsert(ctx context.Context, db *sql.DB, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `newrelic_metric` (`newrelic_metric`.`id`,`newrelic_metric`.`checksum`,`newrelic_metric`.`customer_id`,`newrelic_metric`.`application_ext_id`,`newrelic_metric`.`application_id`,`newrelic_metric`.`name`,`newrelic_metric`.`value_name`,`newrelic_metric`.`start`,`newrelic_metric`.`duration_sec`,`newrelic_metric`.`value`) VALUES (?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `newrelic_metric` (`newrelic_metric`.`id`,`newrelic_metric`.`checksum`,`newrelic_metric`.`customer_id`,`newrelic_metric`.`application_ext_id`,`newrelic_metric`.`application_id`,`newrelic_metric`.`name`,`newrelic_metric`.`value_name`,`newrelic_metric`.`start`,`newrelic_metric`.`duration_sec`,`newrelic_metric`.`value`) VALUES (?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`customer_id`=VALUES(`customer_id`),`application_ext_id`=VALUES(`application_ext_id`),`application_id`=VALUES(`application_id`),`name`=VALUES(`name`),`value_name`=VALUES(`value_name`),`start`=VALUES(`start`),`duration_sec`=VALUES(`duration_sec`),`value`=VALUES(`value`)"
	}
	r, err := db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLInt64(t.ApplicationExtID),
		orm.ToSQLString(t.ApplicationID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.ValueName),
		orm.ToSQLInt64(t.Start),
		orm.ToSQLInt64(t.DurationSec),
		orm.ToSQLFloat64(t.Value),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBUpsertTx will upsert the NewrelicMetric record in the database using the provided transaction
func (t *NewrelicMetric) DBUpsertTx(ctx context.Context, tx *sql.Tx, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `newrelic_metric` (`newrelic_metric`.`id`,`newrelic_metric`.`checksum`,`newrelic_metric`.`customer_id`,`newrelic_metric`.`application_ext_id`,`newrelic_metric`.`application_id`,`newrelic_metric`.`name`,`newrelic_metric`.`value_name`,`newrelic_metric`.`start`,`newrelic_metric`.`duration_sec`,`newrelic_metric`.`value`) VALUES (?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `newrelic_metric` (`newrelic_metric`.`id`,`newrelic_metric`.`checksum`,`newrelic_metric`.`customer_id`,`newrelic_metric`.`application_ext_id`,`newrelic_metric`.`application_id`,`newrelic_metric`.`name`,`newrelic_metric`.`value_name`,`newrelic_metric`.`start`,`newrelic_metric`.`duration_sec`,`newrelic_metric`.`value`) VALUES (?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`customer_id`=VALUES(`customer_id`),`application_ext_id`=VALUES(`application_ext_id`),`application_id`=VALUES(`application_id`),`name`=VALUES(`name`),`value_name`=VALUES(`value_name`),`start`=VALUES(`start`),`duration_sec`=VALUES(`duration_sec`),`value`=VALUES(`value`)"
	}
	r, err := tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLInt64(t.ApplicationExtID),
		orm.ToSQLString(t.ApplicationID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.ValueName),
		orm.ToSQLInt64(t.Start),
		orm.ToSQLInt64(t.DurationSec),
		orm.ToSQLFloat64(t.Value),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBFindOne will find a NewrelicMetric record in the database with the primary key
func (t *NewrelicMetric) DBFindOne(ctx context.Context, db *sql.DB, value string) (bool, error) {
	q := "SELECT `newrelic_metric`.`id`,`newrelic_metric`.`checksum`,`newrelic_metric`.`customer_id`,`newrelic_metric`.`application_ext_id`,`newrelic_metric`.`application_id`,`newrelic_metric`.`name`,`newrelic_metric`.`value_name`,`newrelic_metric`.`start`,`newrelic_metric`.`duration_sec`,`newrelic_metric`.`value` FROM `newrelic_metric` WHERE `id` = ? LIMIT 1"
	row := db.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CustomerID sql.NullString
	var _ApplicationExtID sql.NullInt64
	var _ApplicationID sql.NullString
	var _Name sql.NullString
	var _ValueName sql.NullString
	var _Start sql.NullInt64
	var _DurationSec sql.NullInt64
	var _Value sql.NullFloat64
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_ApplicationExtID,
		&_ApplicationID,
		&_Name,
		&_ValueName,
		&_Start,
		&_DurationSec,
		&_Value,
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
	if _ApplicationExtID.Valid {
		t.SetApplicationExtID(_ApplicationExtID.Int64)
	}
	if _ApplicationID.Valid {
		t.SetApplicationID(_ApplicationID.String)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _ValueName.Valid {
		t.SetValueName(_ValueName.String)
	}
	if _Start.Valid {
		t.SetStart(_Start.Int64)
	}
	if _DurationSec.Valid {
		t.SetDurationSec(_DurationSec.Int64)
	}
	if _Value.Valid {
		t.SetValue(_Value.Float64)
	}
	return true, nil
}

// DBFindOneTx will find a NewrelicMetric record in the database with the primary key using the provided transaction
func (t *NewrelicMetric) DBFindOneTx(ctx context.Context, tx *sql.Tx, value string) (bool, error) {
	q := "SELECT `newrelic_metric`.`id`,`newrelic_metric`.`checksum`,`newrelic_metric`.`customer_id`,`newrelic_metric`.`application_ext_id`,`newrelic_metric`.`application_id`,`newrelic_metric`.`name`,`newrelic_metric`.`value_name`,`newrelic_metric`.`start`,`newrelic_metric`.`duration_sec`,`newrelic_metric`.`value` FROM `newrelic_metric` WHERE `id` = ? LIMIT 1"
	row := tx.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CustomerID sql.NullString
	var _ApplicationExtID sql.NullInt64
	var _ApplicationID sql.NullString
	var _Name sql.NullString
	var _ValueName sql.NullString
	var _Start sql.NullInt64
	var _DurationSec sql.NullInt64
	var _Value sql.NullFloat64
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_ApplicationExtID,
		&_ApplicationID,
		&_Name,
		&_ValueName,
		&_Start,
		&_DurationSec,
		&_Value,
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
	if _ApplicationExtID.Valid {
		t.SetApplicationExtID(_ApplicationExtID.Int64)
	}
	if _ApplicationID.Valid {
		t.SetApplicationID(_ApplicationID.String)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _ValueName.Valid {
		t.SetValueName(_ValueName.String)
	}
	if _Start.Valid {
		t.SetStart(_Start.Int64)
	}
	if _DurationSec.Valid {
		t.SetDurationSec(_DurationSec.Int64)
	}
	if _Value.Valid {
		t.SetValue(_Value.Float64)
	}
	return true, nil
}

// FindNewrelicMetrics will find a NewrelicMetric record in the database with the provided parameters
func FindNewrelicMetrics(ctx context.Context, db *sql.DB, _params ...interface{}) ([]*NewrelicMetric, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("customer_id"),
		orm.Column("application_ext_id"),
		orm.Column("application_id"),
		orm.Column("name"),
		orm.Column("value_name"),
		orm.Column("start"),
		orm.Column("duration_sec"),
		orm.Column("value"),
		orm.Table(NewrelicMetricTableName),
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
	results := make([]*NewrelicMetric, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _ApplicationExtID sql.NullInt64
		var _ApplicationID sql.NullString
		var _Name sql.NullString
		var _ValueName sql.NullString
		var _Start sql.NullInt64
		var _DurationSec sql.NullInt64
		var _Value sql.NullFloat64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_ApplicationExtID,
			&_ApplicationID,
			&_Name,
			&_ValueName,
			&_Start,
			&_DurationSec,
			&_Value,
		)
		if err != nil {
			return nil, err
		}
		t := &NewrelicMetric{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _ApplicationExtID.Valid {
			t.SetApplicationExtID(_ApplicationExtID.Int64)
		}
		if _ApplicationID.Valid {
			t.SetApplicationID(_ApplicationID.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _ValueName.Valid {
			t.SetValueName(_ValueName.String)
		}
		if _Start.Valid {
			t.SetStart(_Start.Int64)
		}
		if _DurationSec.Valid {
			t.SetDurationSec(_DurationSec.Int64)
		}
		if _Value.Valid {
			t.SetValue(_Value.Float64)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindNewrelicMetricsTx will find a NewrelicMetric record in the database with the provided parameters using the provided transaction
func FindNewrelicMetricsTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) ([]*NewrelicMetric, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("customer_id"),
		orm.Column("application_ext_id"),
		orm.Column("application_id"),
		orm.Column("name"),
		orm.Column("value_name"),
		orm.Column("start"),
		orm.Column("duration_sec"),
		orm.Column("value"),
		orm.Table(NewrelicMetricTableName),
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
	results := make([]*NewrelicMetric, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _ApplicationExtID sql.NullInt64
		var _ApplicationID sql.NullString
		var _Name sql.NullString
		var _ValueName sql.NullString
		var _Start sql.NullInt64
		var _DurationSec sql.NullInt64
		var _Value sql.NullFloat64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_ApplicationExtID,
			&_ApplicationID,
			&_Name,
			&_ValueName,
			&_Start,
			&_DurationSec,
			&_Value,
		)
		if err != nil {
			return nil, err
		}
		t := &NewrelicMetric{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _ApplicationExtID.Valid {
			t.SetApplicationExtID(_ApplicationExtID.Int64)
		}
		if _ApplicationID.Valid {
			t.SetApplicationID(_ApplicationID.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _ValueName.Valid {
			t.SetValueName(_ValueName.String)
		}
		if _Start.Valid {
			t.SetStart(_Start.Int64)
		}
		if _DurationSec.Valid {
			t.SetDurationSec(_DurationSec.Int64)
		}
		if _Value.Valid {
			t.SetValue(_Value.Float64)
		}
		results = append(results, t)
	}
	return results, nil
}

// DBFind will find a NewrelicMetric record in the database with the provided parameters
func (t *NewrelicMetric) DBFind(ctx context.Context, db *sql.DB, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("customer_id"),
		orm.Column("application_ext_id"),
		orm.Column("application_id"),
		orm.Column("name"),
		orm.Column("value_name"),
		orm.Column("start"),
		orm.Column("duration_sec"),
		orm.Column("value"),
		orm.Table(NewrelicMetricTableName),
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
	var _ApplicationExtID sql.NullInt64
	var _ApplicationID sql.NullString
	var _Name sql.NullString
	var _ValueName sql.NullString
	var _Start sql.NullInt64
	var _DurationSec sql.NullInt64
	var _Value sql.NullFloat64
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_ApplicationExtID,
		&_ApplicationID,
		&_Name,
		&_ValueName,
		&_Start,
		&_DurationSec,
		&_Value,
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
	if _ApplicationExtID.Valid {
		t.SetApplicationExtID(_ApplicationExtID.Int64)
	}
	if _ApplicationID.Valid {
		t.SetApplicationID(_ApplicationID.String)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _ValueName.Valid {
		t.SetValueName(_ValueName.String)
	}
	if _Start.Valid {
		t.SetStart(_Start.Int64)
	}
	if _DurationSec.Valid {
		t.SetDurationSec(_DurationSec.Int64)
	}
	if _Value.Valid {
		t.SetValue(_Value.Float64)
	}
	return true, nil
}

// DBFindTx will find a NewrelicMetric record in the database with the provided parameters using the provided transaction
func (t *NewrelicMetric) DBFindTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("customer_id"),
		orm.Column("application_ext_id"),
		orm.Column("application_id"),
		orm.Column("name"),
		orm.Column("value_name"),
		orm.Column("start"),
		orm.Column("duration_sec"),
		orm.Column("value"),
		orm.Table(NewrelicMetricTableName),
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
	var _ApplicationExtID sql.NullInt64
	var _ApplicationID sql.NullString
	var _Name sql.NullString
	var _ValueName sql.NullString
	var _Start sql.NullInt64
	var _DurationSec sql.NullInt64
	var _Value sql.NullFloat64
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_ApplicationExtID,
		&_ApplicationID,
		&_Name,
		&_ValueName,
		&_Start,
		&_DurationSec,
		&_Value,
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
	if _ApplicationExtID.Valid {
		t.SetApplicationExtID(_ApplicationExtID.Int64)
	}
	if _ApplicationID.Valid {
		t.SetApplicationID(_ApplicationID.String)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _ValueName.Valid {
		t.SetValueName(_ValueName.String)
	}
	if _Start.Valid {
		t.SetStart(_Start.Int64)
	}
	if _DurationSec.Valid {
		t.SetDurationSec(_DurationSec.Int64)
	}
	if _Value.Valid {
		t.SetValue(_Value.Float64)
	}
	return true, nil
}

// CountNewrelicMetrics will find the count of NewrelicMetric records in the database
func CountNewrelicMetrics(ctx context.Context, db *sql.DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(NewrelicMetricTableName),
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

// CountNewrelicMetricsTx will find the count of NewrelicMetric records in the database using the provided transaction
func CountNewrelicMetricsTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(NewrelicMetricTableName),
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

// DBCount will find the count of NewrelicMetric records in the database
func (t *NewrelicMetric) DBCount(ctx context.Context, db *sql.DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(NewrelicMetricTableName),
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

// DBCountTx will find the count of NewrelicMetric records in the database using the provided transaction
func (t *NewrelicMetric) DBCountTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(NewrelicMetricTableName),
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

// DBExists will return true if the NewrelicMetric record exists in the database
func (t *NewrelicMetric) DBExists(ctx context.Context, db *sql.DB) (bool, error) {
	q := "SELECT `id` FROM `newrelic_metric` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := db.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// DBExistsTx will return true if the NewrelicMetric record exists in the database using the provided transaction
func (t *NewrelicMetric) DBExistsTx(ctx context.Context, tx *sql.Tx) (bool, error) {
	q := "SELECT `id` FROM `newrelic_metric` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := tx.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// PrimaryKeyColumn returns the column name for the primary key
func (t *NewrelicMetric) PrimaryKeyColumn() string {
	return NewrelicMetricColumnID
}

// PrimaryKeyColumnType returns the primary key column Go type as a string
func (t *NewrelicMetric) PrimaryKeyColumnType() string {
	return "string"
}

// PrimaryKey returns the primary key column value
func (t *NewrelicMetric) PrimaryKey() interface{} {
	return t.ID
}
