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

var _ Model = (*NewrelicAlertsCondition)(nil)
var _ CSVWriter = (*NewrelicAlertsCondition)(nil)
var _ JSONWriter = (*NewrelicAlertsCondition)(nil)
var _ Checksum = (*NewrelicAlertsCondition)(nil)

// NewrelicAlertsConditionTableName is the name of the table in SQL
const NewrelicAlertsConditionTableName = "newrelic_alerts_condition"

var NewrelicAlertsConditionColumns = []string{
	"id",
	"checksum",
	"customer_id",
	"policy_ext_id",
	"policy_id",
	"ext_id",
	"type",
	"name",
	"enabled",
	"metric",
	"terms",
}

// NewrelicAlertsCondition table
type NewrelicAlertsCondition struct {
	Checksum    *string `json:"checksum,omitempty"`
	CustomerID  string  `json:"customer_id"`
	Enabled     bool    `json:"enabled"`
	ExtID       int64   `json:"ext_id"`
	ID          string  `json:"id"`
	Metric      string  `json:"metric"`
	Name        string  `json:"name"`
	PolicyExtID int64   `json:"policy_ext_id"`
	PolicyID    string  `json:"policy_id"`
	Terms       string  `json:"terms"`
	Type        string  `json:"type"`
}

// TableName returns the SQL table name for NewrelicAlertsCondition and satifies the Model interface
func (t *NewrelicAlertsCondition) TableName() string {
	return NewrelicAlertsConditionTableName
}

// ToCSV will serialize the NewrelicAlertsCondition instance to a CSV compatible array of strings
func (t *NewrelicAlertsCondition) ToCSV() []string {
	return []string{
		t.ID,
		t.CalculateChecksum(),
		t.CustomerID,
		toCSVString(t.PolicyExtID),
		t.PolicyID,
		toCSVString(t.ExtID),
		t.Type,
		t.Name,
		toCSVBool(t.Enabled),
		t.Metric,
		t.Terms,
	}
}

// WriteCSV will serialize the NewrelicAlertsCondition instance to the writer as CSV and satisfies the CSVWriter interface
func (t *NewrelicAlertsCondition) WriteCSV(w *csv.Writer) error {
	return w.Write(t.ToCSV())
}

// WriteJSON will serialize the NewrelicAlertsCondition instance to the writer as JSON and satisfies the JSONWriter interface
func (t *NewrelicAlertsCondition) WriteJSON(w io.Writer, indent ...bool) error {
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

// NewNewrelicAlertsConditionReader creates a JSON reader which can read in NewrelicAlertsCondition objects serialized as JSON either as an array, single object or json new lines
// and writes each NewrelicAlertsCondition to the channel provided
func NewNewrelicAlertsConditionReader(r io.Reader, ch chan<- NewrelicAlertsCondition) error {
	return orm.Deserialize(r, func(buf json.RawMessage) error {
		dec := json.NewDecoder(bytes.NewBuffer(buf))
		e := NewrelicAlertsCondition{}
		if err := dec.Decode(&e); err != nil {
			return err
		}
		ch <- e
		return nil
	})
}

// NewCSVNewrelicAlertsConditionReaderDir will read the reader as CSV and emit each record to the channel provided
func NewCSVNewrelicAlertsConditionReader(r io.Reader, ch chan<- NewrelicAlertsCondition) error {
	cr := csv.NewReader(r)
	for {
		record, err := cr.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		ch <- NewrelicAlertsCondition{
			ID:          record[0],
			Checksum:    fromStringPointer(record[1]),
			CustomerID:  record[2],
			PolicyExtID: fromCSVInt64(record[3]),
			PolicyID:    record[4],
			ExtID:       fromCSVInt64(record[5]),
			Type:        record[6],
			Name:        record[7],
			Enabled:     fromCSVBool(record[8]),
			Metric:      record[9],
			Terms:       record[10],
		}
	}
	return nil
}

// NewCSVNewrelicAlertsConditionReaderFile will read the file as a CSV and emit each record to the channel provided
func NewCSVNewrelicAlertsConditionReaderFile(fp string, ch chan<- NewrelicAlertsCondition) error {
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
	return NewCSVNewrelicAlertsConditionReader(fc, ch)
}

// NewCSVNewrelicAlertsConditionReaderDir will read the newrelic_alerts_condition.csv.gz file as a CSV and emit each record to the channel provided
func NewCSVNewrelicAlertsConditionReaderDir(dir string, ch chan<- NewrelicAlertsCondition) error {
	return NewCSVNewrelicAlertsConditionReaderFile(filepath.Join(dir, "newrelic_alerts_condition.csv.gz"), ch)
}

// NewrelicAlertsConditionCSVDeduper is a function callback which takes the existing value (a) and the new value (b)
// and the return value should be the one to use (or a new one, if applicable). return nil
// to skip processing of this record
type NewrelicAlertsConditionCSVDeduper func(a NewrelicAlertsCondition, b NewrelicAlertsCondition) *NewrelicAlertsCondition

// NewrelicAlertsConditionCSVDedupeDisabled is set on whether the CSV writer should de-dupe values by key
var NewrelicAlertsConditionCSVDedupeDisabled bool

// NewNewrelicAlertsConditionCSVWriterSize creates a batch writer that will write each NewrelicAlertsCondition into a CSV file
// this method will automatically de-duplicate entries using the primary key. if the checksum
// for a newer item with the same primary key doesn't match a previously sent item, the newer
// one will be used
func NewNewrelicAlertsConditionCSVWriterSize(w io.Writer, size int, dedupers ...NewrelicAlertsConditionCSVDeduper) (chan NewrelicAlertsCondition, chan bool, error) {
	cw := csv.NewWriter(w)
	ch := make(chan NewrelicAlertsCondition, size)
	done := make(chan bool)
	go func() {
		defer func() { done <- true }()
		dodedupe := !NewrelicAlertsConditionCSVDedupeDisabled
		var kv map[string]*NewrelicAlertsCondition
		var deduper NewrelicAlertsConditionCSVDeduper
		if dedupers != nil && len(dedupers) > 0 {
			deduper = dedupers[0]
			dodedupe = true
		}
		if dodedupe {
			kv = make(map[string]*NewrelicAlertsCondition)
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

// NewrelicAlertsConditionCSVDefaultSize is the default channel buffer size if not provided
var NewrelicAlertsConditionCSVDefaultSize = 100

// NewNewrelicAlertsConditionCSVWriter creates a batch writer that will write each NewrelicAlertsCondition into a CSV file
func NewNewrelicAlertsConditionCSVWriter(w io.Writer, dedupers ...NewrelicAlertsConditionCSVDeduper) (chan NewrelicAlertsCondition, chan bool, error) {
	return NewNewrelicAlertsConditionCSVWriterSize(w, NewrelicAlertsConditionCSVDefaultSize, dedupers...)
}

// NewNewrelicAlertsConditionCSVWriterDir creates a batch writer that will write each NewrelicAlertsCondition into a CSV file named newrelic_alerts_condition.csv.gz in dir
func NewNewrelicAlertsConditionCSVWriterDir(dir string, dedupers ...NewrelicAlertsConditionCSVDeduper) (chan NewrelicAlertsCondition, chan bool, error) {
	return NewNewrelicAlertsConditionCSVWriterFile(filepath.Join(dir, "newrelic_alerts_condition.csv.gz"), dedupers...)
}

// NewNewrelicAlertsConditionCSVWriterFile creates a batch writer that will write each NewrelicAlertsCondition into a CSV file
func NewNewrelicAlertsConditionCSVWriterFile(fn string, dedupers ...NewrelicAlertsConditionCSVDeduper) (chan NewrelicAlertsCondition, chan bool, error) {
	f, err := os.Create(fn)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening CSV file %s. %v", fn, err)
	}
	var fc io.WriteCloser = f
	if filepath.Ext(fn) == ".gz" {
		w, _ := gzip.NewWriterLevel(f, gzip.BestCompression)
		fc = w
	}
	ch, done, err := NewNewrelicAlertsConditionCSVWriter(fc, dedupers...)
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

type NewrelicAlertsConditionDBAction func(ctx context.Context, db *sql.DB, record NewrelicAlertsCondition) error

// NewNewrelicAlertsConditionDBWriterSize creates a DB writer that will write each issue into the DB
func NewNewrelicAlertsConditionDBWriterSize(ctx context.Context, db *sql.DB, errors chan<- error, size int, actions ...NewrelicAlertsConditionDBAction) (chan NewrelicAlertsCondition, chan bool, error) {
	ch := make(chan NewrelicAlertsCondition, size)
	done := make(chan bool)
	var action NewrelicAlertsConditionDBAction
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

// NewNewrelicAlertsConditionDBWriter creates a DB writer that will write each issue into the DB
func NewNewrelicAlertsConditionDBWriter(ctx context.Context, db *sql.DB, errors chan<- error, actions ...NewrelicAlertsConditionDBAction) (chan NewrelicAlertsCondition, chan bool, error) {
	return NewNewrelicAlertsConditionDBWriterSize(ctx, db, errors, 100, actions...)
}

// NewrelicAlertsConditionColumnID is the ID SQL column name for the NewrelicAlertsCondition table
const NewrelicAlertsConditionColumnID = "id"

// NewrelicAlertsConditionEscapedColumnID is the escaped ID SQL column name for the NewrelicAlertsCondition table
const NewrelicAlertsConditionEscapedColumnID = "`id`"

// NewrelicAlertsConditionColumnChecksum is the Checksum SQL column name for the NewrelicAlertsCondition table
const NewrelicAlertsConditionColumnChecksum = "checksum"

// NewrelicAlertsConditionEscapedColumnChecksum is the escaped Checksum SQL column name for the NewrelicAlertsCondition table
const NewrelicAlertsConditionEscapedColumnChecksum = "`checksum`"

// NewrelicAlertsConditionColumnCustomerID is the CustomerID SQL column name for the NewrelicAlertsCondition table
const NewrelicAlertsConditionColumnCustomerID = "customer_id"

// NewrelicAlertsConditionEscapedColumnCustomerID is the escaped CustomerID SQL column name for the NewrelicAlertsCondition table
const NewrelicAlertsConditionEscapedColumnCustomerID = "`customer_id`"

// NewrelicAlertsConditionColumnPolicyExtID is the PolicyExtID SQL column name for the NewrelicAlertsCondition table
const NewrelicAlertsConditionColumnPolicyExtID = "policy_ext_id"

// NewrelicAlertsConditionEscapedColumnPolicyExtID is the escaped PolicyExtID SQL column name for the NewrelicAlertsCondition table
const NewrelicAlertsConditionEscapedColumnPolicyExtID = "`policy_ext_id`"

// NewrelicAlertsConditionColumnPolicyID is the PolicyID SQL column name for the NewrelicAlertsCondition table
const NewrelicAlertsConditionColumnPolicyID = "policy_id"

// NewrelicAlertsConditionEscapedColumnPolicyID is the escaped PolicyID SQL column name for the NewrelicAlertsCondition table
const NewrelicAlertsConditionEscapedColumnPolicyID = "`policy_id`"

// NewrelicAlertsConditionColumnExtID is the ExtID SQL column name for the NewrelicAlertsCondition table
const NewrelicAlertsConditionColumnExtID = "ext_id"

// NewrelicAlertsConditionEscapedColumnExtID is the escaped ExtID SQL column name for the NewrelicAlertsCondition table
const NewrelicAlertsConditionEscapedColumnExtID = "`ext_id`"

// NewrelicAlertsConditionColumnType is the Type SQL column name for the NewrelicAlertsCondition table
const NewrelicAlertsConditionColumnType = "type"

// NewrelicAlertsConditionEscapedColumnType is the escaped Type SQL column name for the NewrelicAlertsCondition table
const NewrelicAlertsConditionEscapedColumnType = "`type`"

// NewrelicAlertsConditionColumnName is the Name SQL column name for the NewrelicAlertsCondition table
const NewrelicAlertsConditionColumnName = "name"

// NewrelicAlertsConditionEscapedColumnName is the escaped Name SQL column name for the NewrelicAlertsCondition table
const NewrelicAlertsConditionEscapedColumnName = "`name`"

// NewrelicAlertsConditionColumnEnabled is the Enabled SQL column name for the NewrelicAlertsCondition table
const NewrelicAlertsConditionColumnEnabled = "enabled"

// NewrelicAlertsConditionEscapedColumnEnabled is the escaped Enabled SQL column name for the NewrelicAlertsCondition table
const NewrelicAlertsConditionEscapedColumnEnabled = "`enabled`"

// NewrelicAlertsConditionColumnMetric is the Metric SQL column name for the NewrelicAlertsCondition table
const NewrelicAlertsConditionColumnMetric = "metric"

// NewrelicAlertsConditionEscapedColumnMetric is the escaped Metric SQL column name for the NewrelicAlertsCondition table
const NewrelicAlertsConditionEscapedColumnMetric = "`metric`"

// NewrelicAlertsConditionColumnTerms is the Terms SQL column name for the NewrelicAlertsCondition table
const NewrelicAlertsConditionColumnTerms = "terms"

// NewrelicAlertsConditionEscapedColumnTerms is the escaped Terms SQL column name for the NewrelicAlertsCondition table
const NewrelicAlertsConditionEscapedColumnTerms = "`terms`"

// GetID will return the NewrelicAlertsCondition ID value
func (t *NewrelicAlertsCondition) GetID() string {
	return t.ID
}

// SetID will set the NewrelicAlertsCondition ID value
func (t *NewrelicAlertsCondition) SetID(v string) {
	t.ID = v
}

// FindNewrelicAlertsConditionByID will find a NewrelicAlertsCondition by ID
func FindNewrelicAlertsConditionByID(ctx context.Context, db *sql.DB, value string) (*NewrelicAlertsCondition, error) {
	q := "SELECT `newrelic_alerts_condition`.`id`,`newrelic_alerts_condition`.`checksum`,`newrelic_alerts_condition`.`customer_id`,`newrelic_alerts_condition`.`policy_ext_id`,`newrelic_alerts_condition`.`policy_id`,`newrelic_alerts_condition`.`ext_id`,`newrelic_alerts_condition`.`type`,`newrelic_alerts_condition`.`name`,`newrelic_alerts_condition`.`enabled`,`newrelic_alerts_condition`.`metric`,`newrelic_alerts_condition`.`terms` FROM `newrelic_alerts_condition` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CustomerID sql.NullString
	var _PolicyExtID sql.NullInt64
	var _PolicyID sql.NullString
	var _ExtID sql.NullInt64
	var _Type sql.NullString
	var _Name sql.NullString
	var _Enabled sql.NullBool
	var _Metric sql.NullString
	var _Terms sql.NullString
	err := db.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_PolicyExtID,
		&_PolicyID,
		&_ExtID,
		&_Type,
		&_Name,
		&_Enabled,
		&_Metric,
		&_Terms,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &NewrelicAlertsCondition{}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _Checksum.Valid {
		t.SetChecksum(_Checksum.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _PolicyExtID.Valid {
		t.SetPolicyExtID(_PolicyExtID.Int64)
	}
	if _PolicyID.Valid {
		t.SetPolicyID(_PolicyID.String)
	}
	if _ExtID.Valid {
		t.SetExtID(_ExtID.Int64)
	}
	if _Type.Valid {
		t.SetType(_Type.String)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _Enabled.Valid {
		t.SetEnabled(_Enabled.Bool)
	}
	if _Metric.Valid {
		t.SetMetric(_Metric.String)
	}
	if _Terms.Valid {
		t.SetTerms(_Terms.String)
	}
	return t, nil
}

// FindNewrelicAlertsConditionByIDTx will find a NewrelicAlertsCondition by ID using the provided transaction
func FindNewrelicAlertsConditionByIDTx(ctx context.Context, tx *sql.Tx, value string) (*NewrelicAlertsCondition, error) {
	q := "SELECT `newrelic_alerts_condition`.`id`,`newrelic_alerts_condition`.`checksum`,`newrelic_alerts_condition`.`customer_id`,`newrelic_alerts_condition`.`policy_ext_id`,`newrelic_alerts_condition`.`policy_id`,`newrelic_alerts_condition`.`ext_id`,`newrelic_alerts_condition`.`type`,`newrelic_alerts_condition`.`name`,`newrelic_alerts_condition`.`enabled`,`newrelic_alerts_condition`.`metric`,`newrelic_alerts_condition`.`terms` FROM `newrelic_alerts_condition` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CustomerID sql.NullString
	var _PolicyExtID sql.NullInt64
	var _PolicyID sql.NullString
	var _ExtID sql.NullInt64
	var _Type sql.NullString
	var _Name sql.NullString
	var _Enabled sql.NullBool
	var _Metric sql.NullString
	var _Terms sql.NullString
	err := tx.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_PolicyExtID,
		&_PolicyID,
		&_ExtID,
		&_Type,
		&_Name,
		&_Enabled,
		&_Metric,
		&_Terms,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &NewrelicAlertsCondition{}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _Checksum.Valid {
		t.SetChecksum(_Checksum.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _PolicyExtID.Valid {
		t.SetPolicyExtID(_PolicyExtID.Int64)
	}
	if _PolicyID.Valid {
		t.SetPolicyID(_PolicyID.String)
	}
	if _ExtID.Valid {
		t.SetExtID(_ExtID.Int64)
	}
	if _Type.Valid {
		t.SetType(_Type.String)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _Enabled.Valid {
		t.SetEnabled(_Enabled.Bool)
	}
	if _Metric.Valid {
		t.SetMetric(_Metric.String)
	}
	if _Terms.Valid {
		t.SetTerms(_Terms.String)
	}
	return t, nil
}

// GetChecksum will return the NewrelicAlertsCondition Checksum value
func (t *NewrelicAlertsCondition) GetChecksum() string {
	if t.Checksum == nil {
		return ""
	}
	return *t.Checksum
}

// SetChecksum will set the NewrelicAlertsCondition Checksum value
func (t *NewrelicAlertsCondition) SetChecksum(v string) {
	t.Checksum = &v
}

// GetCustomerID will return the NewrelicAlertsCondition CustomerID value
func (t *NewrelicAlertsCondition) GetCustomerID() string {
	return t.CustomerID
}

// SetCustomerID will set the NewrelicAlertsCondition CustomerID value
func (t *NewrelicAlertsCondition) SetCustomerID(v string) {
	t.CustomerID = v
}

// FindNewrelicAlertsConditionsByCustomerID will find all NewrelicAlertsConditions by the CustomerID value
func FindNewrelicAlertsConditionsByCustomerID(ctx context.Context, db *sql.DB, value string) ([]*NewrelicAlertsCondition, error) {
	q := "SELECT `newrelic_alerts_condition`.`id`,`newrelic_alerts_condition`.`checksum`,`newrelic_alerts_condition`.`customer_id`,`newrelic_alerts_condition`.`policy_ext_id`,`newrelic_alerts_condition`.`policy_id`,`newrelic_alerts_condition`.`ext_id`,`newrelic_alerts_condition`.`type`,`newrelic_alerts_condition`.`name`,`newrelic_alerts_condition`.`enabled`,`newrelic_alerts_condition`.`metric`,`newrelic_alerts_condition`.`terms` FROM `newrelic_alerts_condition` WHERE `customer_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*NewrelicAlertsCondition, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _PolicyExtID sql.NullInt64
		var _PolicyID sql.NullString
		var _ExtID sql.NullInt64
		var _Type sql.NullString
		var _Name sql.NullString
		var _Enabled sql.NullBool
		var _Metric sql.NullString
		var _Terms sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_PolicyExtID,
			&_PolicyID,
			&_ExtID,
			&_Type,
			&_Name,
			&_Enabled,
			&_Metric,
			&_Terms,
		)
		if err != nil {
			return nil, err
		}
		t := &NewrelicAlertsCondition{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _PolicyExtID.Valid {
			t.SetPolicyExtID(_PolicyExtID.Int64)
		}
		if _PolicyID.Valid {
			t.SetPolicyID(_PolicyID.String)
		}
		if _ExtID.Valid {
			t.SetExtID(_ExtID.Int64)
		}
		if _Type.Valid {
			t.SetType(_Type.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _Enabled.Valid {
			t.SetEnabled(_Enabled.Bool)
		}
		if _Metric.Valid {
			t.SetMetric(_Metric.String)
		}
		if _Terms.Valid {
			t.SetTerms(_Terms.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindNewrelicAlertsConditionsByCustomerIDTx will find all NewrelicAlertsConditions by the CustomerID value using the provided transaction
func FindNewrelicAlertsConditionsByCustomerIDTx(ctx context.Context, tx *sql.Tx, value string) ([]*NewrelicAlertsCondition, error) {
	q := "SELECT `newrelic_alerts_condition`.`id`,`newrelic_alerts_condition`.`checksum`,`newrelic_alerts_condition`.`customer_id`,`newrelic_alerts_condition`.`policy_ext_id`,`newrelic_alerts_condition`.`policy_id`,`newrelic_alerts_condition`.`ext_id`,`newrelic_alerts_condition`.`type`,`newrelic_alerts_condition`.`name`,`newrelic_alerts_condition`.`enabled`,`newrelic_alerts_condition`.`metric`,`newrelic_alerts_condition`.`terms` FROM `newrelic_alerts_condition` WHERE `customer_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*NewrelicAlertsCondition, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _PolicyExtID sql.NullInt64
		var _PolicyID sql.NullString
		var _ExtID sql.NullInt64
		var _Type sql.NullString
		var _Name sql.NullString
		var _Enabled sql.NullBool
		var _Metric sql.NullString
		var _Terms sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_PolicyExtID,
			&_PolicyID,
			&_ExtID,
			&_Type,
			&_Name,
			&_Enabled,
			&_Metric,
			&_Terms,
		)
		if err != nil {
			return nil, err
		}
		t := &NewrelicAlertsCondition{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _PolicyExtID.Valid {
			t.SetPolicyExtID(_PolicyExtID.Int64)
		}
		if _PolicyID.Valid {
			t.SetPolicyID(_PolicyID.String)
		}
		if _ExtID.Valid {
			t.SetExtID(_ExtID.Int64)
		}
		if _Type.Valid {
			t.SetType(_Type.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _Enabled.Valid {
			t.SetEnabled(_Enabled.Bool)
		}
		if _Metric.Valid {
			t.SetMetric(_Metric.String)
		}
		if _Terms.Valid {
			t.SetTerms(_Terms.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetPolicyExtID will return the NewrelicAlertsCondition PolicyExtID value
func (t *NewrelicAlertsCondition) GetPolicyExtID() int64 {
	return t.PolicyExtID
}

// SetPolicyExtID will set the NewrelicAlertsCondition PolicyExtID value
func (t *NewrelicAlertsCondition) SetPolicyExtID(v int64) {
	t.PolicyExtID = v
}

// FindNewrelicAlertsConditionsByPolicyExtID will find all NewrelicAlertsConditions by the PolicyExtID value
func FindNewrelicAlertsConditionsByPolicyExtID(ctx context.Context, db *sql.DB, value int64) ([]*NewrelicAlertsCondition, error) {
	q := "SELECT `newrelic_alerts_condition`.`id`,`newrelic_alerts_condition`.`checksum`,`newrelic_alerts_condition`.`customer_id`,`newrelic_alerts_condition`.`policy_ext_id`,`newrelic_alerts_condition`.`policy_id`,`newrelic_alerts_condition`.`ext_id`,`newrelic_alerts_condition`.`type`,`newrelic_alerts_condition`.`name`,`newrelic_alerts_condition`.`enabled`,`newrelic_alerts_condition`.`metric`,`newrelic_alerts_condition`.`terms` FROM `newrelic_alerts_condition` WHERE `policy_ext_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*NewrelicAlertsCondition, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _PolicyExtID sql.NullInt64
		var _PolicyID sql.NullString
		var _ExtID sql.NullInt64
		var _Type sql.NullString
		var _Name sql.NullString
		var _Enabled sql.NullBool
		var _Metric sql.NullString
		var _Terms sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_PolicyExtID,
			&_PolicyID,
			&_ExtID,
			&_Type,
			&_Name,
			&_Enabled,
			&_Metric,
			&_Terms,
		)
		if err != nil {
			return nil, err
		}
		t := &NewrelicAlertsCondition{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _PolicyExtID.Valid {
			t.SetPolicyExtID(_PolicyExtID.Int64)
		}
		if _PolicyID.Valid {
			t.SetPolicyID(_PolicyID.String)
		}
		if _ExtID.Valid {
			t.SetExtID(_ExtID.Int64)
		}
		if _Type.Valid {
			t.SetType(_Type.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _Enabled.Valid {
			t.SetEnabled(_Enabled.Bool)
		}
		if _Metric.Valid {
			t.SetMetric(_Metric.String)
		}
		if _Terms.Valid {
			t.SetTerms(_Terms.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindNewrelicAlertsConditionsByPolicyExtIDTx will find all NewrelicAlertsConditions by the PolicyExtID value using the provided transaction
func FindNewrelicAlertsConditionsByPolicyExtIDTx(ctx context.Context, tx *sql.Tx, value int64) ([]*NewrelicAlertsCondition, error) {
	q := "SELECT `newrelic_alerts_condition`.`id`,`newrelic_alerts_condition`.`checksum`,`newrelic_alerts_condition`.`customer_id`,`newrelic_alerts_condition`.`policy_ext_id`,`newrelic_alerts_condition`.`policy_id`,`newrelic_alerts_condition`.`ext_id`,`newrelic_alerts_condition`.`type`,`newrelic_alerts_condition`.`name`,`newrelic_alerts_condition`.`enabled`,`newrelic_alerts_condition`.`metric`,`newrelic_alerts_condition`.`terms` FROM `newrelic_alerts_condition` WHERE `policy_ext_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*NewrelicAlertsCondition, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _PolicyExtID sql.NullInt64
		var _PolicyID sql.NullString
		var _ExtID sql.NullInt64
		var _Type sql.NullString
		var _Name sql.NullString
		var _Enabled sql.NullBool
		var _Metric sql.NullString
		var _Terms sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_PolicyExtID,
			&_PolicyID,
			&_ExtID,
			&_Type,
			&_Name,
			&_Enabled,
			&_Metric,
			&_Terms,
		)
		if err != nil {
			return nil, err
		}
		t := &NewrelicAlertsCondition{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _PolicyExtID.Valid {
			t.SetPolicyExtID(_PolicyExtID.Int64)
		}
		if _PolicyID.Valid {
			t.SetPolicyID(_PolicyID.String)
		}
		if _ExtID.Valid {
			t.SetExtID(_ExtID.Int64)
		}
		if _Type.Valid {
			t.SetType(_Type.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _Enabled.Valid {
			t.SetEnabled(_Enabled.Bool)
		}
		if _Metric.Valid {
			t.SetMetric(_Metric.String)
		}
		if _Terms.Valid {
			t.SetTerms(_Terms.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetPolicyID will return the NewrelicAlertsCondition PolicyID value
func (t *NewrelicAlertsCondition) GetPolicyID() string {
	return t.PolicyID
}

// SetPolicyID will set the NewrelicAlertsCondition PolicyID value
func (t *NewrelicAlertsCondition) SetPolicyID(v string) {
	t.PolicyID = v
}

// FindNewrelicAlertsConditionsByPolicyID will find all NewrelicAlertsConditions by the PolicyID value
func FindNewrelicAlertsConditionsByPolicyID(ctx context.Context, db *sql.DB, value string) ([]*NewrelicAlertsCondition, error) {
	q := "SELECT `newrelic_alerts_condition`.`id`,`newrelic_alerts_condition`.`checksum`,`newrelic_alerts_condition`.`customer_id`,`newrelic_alerts_condition`.`policy_ext_id`,`newrelic_alerts_condition`.`policy_id`,`newrelic_alerts_condition`.`ext_id`,`newrelic_alerts_condition`.`type`,`newrelic_alerts_condition`.`name`,`newrelic_alerts_condition`.`enabled`,`newrelic_alerts_condition`.`metric`,`newrelic_alerts_condition`.`terms` FROM `newrelic_alerts_condition` WHERE `policy_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*NewrelicAlertsCondition, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _PolicyExtID sql.NullInt64
		var _PolicyID sql.NullString
		var _ExtID sql.NullInt64
		var _Type sql.NullString
		var _Name sql.NullString
		var _Enabled sql.NullBool
		var _Metric sql.NullString
		var _Terms sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_PolicyExtID,
			&_PolicyID,
			&_ExtID,
			&_Type,
			&_Name,
			&_Enabled,
			&_Metric,
			&_Terms,
		)
		if err != nil {
			return nil, err
		}
		t := &NewrelicAlertsCondition{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _PolicyExtID.Valid {
			t.SetPolicyExtID(_PolicyExtID.Int64)
		}
		if _PolicyID.Valid {
			t.SetPolicyID(_PolicyID.String)
		}
		if _ExtID.Valid {
			t.SetExtID(_ExtID.Int64)
		}
		if _Type.Valid {
			t.SetType(_Type.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _Enabled.Valid {
			t.SetEnabled(_Enabled.Bool)
		}
		if _Metric.Valid {
			t.SetMetric(_Metric.String)
		}
		if _Terms.Valid {
			t.SetTerms(_Terms.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindNewrelicAlertsConditionsByPolicyIDTx will find all NewrelicAlertsConditions by the PolicyID value using the provided transaction
func FindNewrelicAlertsConditionsByPolicyIDTx(ctx context.Context, tx *sql.Tx, value string) ([]*NewrelicAlertsCondition, error) {
	q := "SELECT `newrelic_alerts_condition`.`id`,`newrelic_alerts_condition`.`checksum`,`newrelic_alerts_condition`.`customer_id`,`newrelic_alerts_condition`.`policy_ext_id`,`newrelic_alerts_condition`.`policy_id`,`newrelic_alerts_condition`.`ext_id`,`newrelic_alerts_condition`.`type`,`newrelic_alerts_condition`.`name`,`newrelic_alerts_condition`.`enabled`,`newrelic_alerts_condition`.`metric`,`newrelic_alerts_condition`.`terms` FROM `newrelic_alerts_condition` WHERE `policy_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*NewrelicAlertsCondition, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _PolicyExtID sql.NullInt64
		var _PolicyID sql.NullString
		var _ExtID sql.NullInt64
		var _Type sql.NullString
		var _Name sql.NullString
		var _Enabled sql.NullBool
		var _Metric sql.NullString
		var _Terms sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_PolicyExtID,
			&_PolicyID,
			&_ExtID,
			&_Type,
			&_Name,
			&_Enabled,
			&_Metric,
			&_Terms,
		)
		if err != nil {
			return nil, err
		}
		t := &NewrelicAlertsCondition{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _PolicyExtID.Valid {
			t.SetPolicyExtID(_PolicyExtID.Int64)
		}
		if _PolicyID.Valid {
			t.SetPolicyID(_PolicyID.String)
		}
		if _ExtID.Valid {
			t.SetExtID(_ExtID.Int64)
		}
		if _Type.Valid {
			t.SetType(_Type.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _Enabled.Valid {
			t.SetEnabled(_Enabled.Bool)
		}
		if _Metric.Valid {
			t.SetMetric(_Metric.String)
		}
		if _Terms.Valid {
			t.SetTerms(_Terms.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetExtID will return the NewrelicAlertsCondition ExtID value
func (t *NewrelicAlertsCondition) GetExtID() int64 {
	return t.ExtID
}

// SetExtID will set the NewrelicAlertsCondition ExtID value
func (t *NewrelicAlertsCondition) SetExtID(v int64) {
	t.ExtID = v
}

// GetType will return the NewrelicAlertsCondition Type value
func (t *NewrelicAlertsCondition) GetType() string {
	return t.Type
}

// SetType will set the NewrelicAlertsCondition Type value
func (t *NewrelicAlertsCondition) SetType(v string) {
	t.Type = v
}

// GetName will return the NewrelicAlertsCondition Name value
func (t *NewrelicAlertsCondition) GetName() string {
	return t.Name
}

// SetName will set the NewrelicAlertsCondition Name value
func (t *NewrelicAlertsCondition) SetName(v string) {
	t.Name = v
}

// GetEnabled will return the NewrelicAlertsCondition Enabled value
func (t *NewrelicAlertsCondition) GetEnabled() bool {
	return t.Enabled
}

// SetEnabled will set the NewrelicAlertsCondition Enabled value
func (t *NewrelicAlertsCondition) SetEnabled(v bool) {
	t.Enabled = v
}

// GetMetric will return the NewrelicAlertsCondition Metric value
func (t *NewrelicAlertsCondition) GetMetric() string {
	return t.Metric
}

// SetMetric will set the NewrelicAlertsCondition Metric value
func (t *NewrelicAlertsCondition) SetMetric(v string) {
	t.Metric = v
}

// GetTerms will return the NewrelicAlertsCondition Terms value
func (t *NewrelicAlertsCondition) GetTerms() string {
	return t.Terms
}

// SetTerms will set the NewrelicAlertsCondition Terms value
func (t *NewrelicAlertsCondition) SetTerms(v string) {
	t.Terms = v
}

func (t *NewrelicAlertsCondition) toTimestamp(value time.Time) *timestamp.Timestamp {
	ts, _ := ptypes.TimestampProto(value)
	return ts
}

// DBCreateNewrelicAlertsConditionTable will create the NewrelicAlertsCondition table
func DBCreateNewrelicAlertsConditionTable(ctx context.Context, db *sql.DB) error {
	q := "CREATE TABLE `newrelic_alerts_condition` (`id`VARCHAR(64) NOT NULL PRIMARY KEY,`checksum`CHAR(64),`customer_id`VARCHAR(64) NOT NULL,`policy_ext_id` BIGINT(20) NOT NULL,`policy_id` VARCHAR(64) NOT NULL,`ext_id` BIGINT(20) NOT NULL,`type` TEXT NOT NULL,`name` TEXT NOT NULL,`enabled` BOOL NOT NULL,`metric` TEXT NOT NULL,`terms`JSON NOT NULL,INDEX newrelic_alerts_condition_customer_id_index (`customer_id`),INDEX newrelic_alerts_condition_policy_ext_id_index (`policy_ext_id`),INDEX newrelic_alerts_condition_policy_id_index (`policy_id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBCreateNewrelicAlertsConditionTableTx will create the NewrelicAlertsCondition table using the provided transction
func DBCreateNewrelicAlertsConditionTableTx(ctx context.Context, tx *sql.Tx) error {
	q := "CREATE TABLE `newrelic_alerts_condition` (`id`VARCHAR(64) NOT NULL PRIMARY KEY,`checksum`CHAR(64),`customer_id`VARCHAR(64) NOT NULL,`policy_ext_id` BIGINT(20) NOT NULL,`policy_id` VARCHAR(64) NOT NULL,`ext_id` BIGINT(20) NOT NULL,`type` TEXT NOT NULL,`name` TEXT NOT NULL,`enabled` BOOL NOT NULL,`metric` TEXT NOT NULL,`terms`JSON NOT NULL,INDEX newrelic_alerts_condition_customer_id_index (`customer_id`),INDEX newrelic_alerts_condition_policy_ext_id_index (`policy_ext_id`),INDEX newrelic_alerts_condition_policy_id_index (`policy_id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// DBDropNewrelicAlertsConditionTable will drop the NewrelicAlertsCondition table
func DBDropNewrelicAlertsConditionTable(ctx context.Context, db *sql.DB) error {
	q := "DROP TABLE IF EXISTS `newrelic_alerts_condition`"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBDropNewrelicAlertsConditionTableTx will drop the NewrelicAlertsCondition table using the provided transaction
func DBDropNewrelicAlertsConditionTableTx(ctx context.Context, tx *sql.Tx) error {
	q := "DROP TABLE IF EXISTS `newrelic_alerts_condition`"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// CalculateChecksum will calculate a checksum of the SHA1 of all field values
func (t *NewrelicAlertsCondition) CalculateChecksum() string {
	return orm.HashStrings(
		orm.ToString(t.ID),
		orm.ToString(t.CustomerID),
		orm.ToString(t.PolicyExtID),
		orm.ToString(t.PolicyID),
		orm.ToString(t.ExtID),
		orm.ToString(t.Type),
		orm.ToString(t.Name),
		orm.ToString(t.Enabled),
		orm.ToString(t.Metric),
		orm.ToString(t.Terms),
	)
}

// DBCreate will create a new NewrelicAlertsCondition record in the database
func (t *NewrelicAlertsCondition) DBCreate(ctx context.Context, db *sql.DB) (sql.Result, error) {
	q := "INSERT INTO `newrelic_alerts_condition` (`newrelic_alerts_condition`.`id`,`newrelic_alerts_condition`.`checksum`,`newrelic_alerts_condition`.`customer_id`,`newrelic_alerts_condition`.`policy_ext_id`,`newrelic_alerts_condition`.`policy_id`,`newrelic_alerts_condition`.`ext_id`,`newrelic_alerts_condition`.`type`,`newrelic_alerts_condition`.`name`,`newrelic_alerts_condition`.`enabled`,`newrelic_alerts_condition`.`metric`,`newrelic_alerts_condition`.`terms`) VALUES (?,?,?,?,?,?,?,?,?,?,?)"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLInt64(t.PolicyExtID),
		orm.ToSQLString(t.PolicyID),
		orm.ToSQLInt64(t.ExtID),
		orm.ToSQLString(t.Type),
		orm.ToSQLString(t.Name),
		orm.ToSQLBool(t.Enabled),
		orm.ToSQLString(t.Metric),
		orm.ToSQLString(t.Terms),
	)
}

// DBCreateTx will create a new NewrelicAlertsCondition record in the database using the provided transaction
func (t *NewrelicAlertsCondition) DBCreateTx(ctx context.Context, tx *sql.Tx) (sql.Result, error) {
	q := "INSERT INTO `newrelic_alerts_condition` (`newrelic_alerts_condition`.`id`,`newrelic_alerts_condition`.`checksum`,`newrelic_alerts_condition`.`customer_id`,`newrelic_alerts_condition`.`policy_ext_id`,`newrelic_alerts_condition`.`policy_id`,`newrelic_alerts_condition`.`ext_id`,`newrelic_alerts_condition`.`type`,`newrelic_alerts_condition`.`name`,`newrelic_alerts_condition`.`enabled`,`newrelic_alerts_condition`.`metric`,`newrelic_alerts_condition`.`terms`) VALUES (?,?,?,?,?,?,?,?,?,?,?)"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLInt64(t.PolicyExtID),
		orm.ToSQLString(t.PolicyID),
		orm.ToSQLInt64(t.ExtID),
		orm.ToSQLString(t.Type),
		orm.ToSQLString(t.Name),
		orm.ToSQLBool(t.Enabled),
		orm.ToSQLString(t.Metric),
		orm.ToSQLString(t.Terms),
	)
}

// DBCreateIgnoreDuplicate will upsert the NewrelicAlertsCondition record in the database
func (t *NewrelicAlertsCondition) DBCreateIgnoreDuplicate(ctx context.Context, db *sql.DB) (sql.Result, error) {
	q := "INSERT INTO `newrelic_alerts_condition` (`newrelic_alerts_condition`.`id`,`newrelic_alerts_condition`.`checksum`,`newrelic_alerts_condition`.`customer_id`,`newrelic_alerts_condition`.`policy_ext_id`,`newrelic_alerts_condition`.`policy_id`,`newrelic_alerts_condition`.`ext_id`,`newrelic_alerts_condition`.`type`,`newrelic_alerts_condition`.`name`,`newrelic_alerts_condition`.`enabled`,`newrelic_alerts_condition`.`metric`,`newrelic_alerts_condition`.`terms`) VALUES (?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLInt64(t.PolicyExtID),
		orm.ToSQLString(t.PolicyID),
		orm.ToSQLInt64(t.ExtID),
		orm.ToSQLString(t.Type),
		orm.ToSQLString(t.Name),
		orm.ToSQLBool(t.Enabled),
		orm.ToSQLString(t.Metric),
		orm.ToSQLString(t.Terms),
	)
}

// DBCreateIgnoreDuplicateTx will upsert the NewrelicAlertsCondition record in the database using the provided transaction
func (t *NewrelicAlertsCondition) DBCreateIgnoreDuplicateTx(ctx context.Context, tx *sql.Tx) (sql.Result, error) {
	q := "INSERT INTO `newrelic_alerts_condition` (`newrelic_alerts_condition`.`id`,`newrelic_alerts_condition`.`checksum`,`newrelic_alerts_condition`.`customer_id`,`newrelic_alerts_condition`.`policy_ext_id`,`newrelic_alerts_condition`.`policy_id`,`newrelic_alerts_condition`.`ext_id`,`newrelic_alerts_condition`.`type`,`newrelic_alerts_condition`.`name`,`newrelic_alerts_condition`.`enabled`,`newrelic_alerts_condition`.`metric`,`newrelic_alerts_condition`.`terms`) VALUES (?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLInt64(t.PolicyExtID),
		orm.ToSQLString(t.PolicyID),
		orm.ToSQLInt64(t.ExtID),
		orm.ToSQLString(t.Type),
		orm.ToSQLString(t.Name),
		orm.ToSQLBool(t.Enabled),
		orm.ToSQLString(t.Metric),
		orm.ToSQLString(t.Terms),
	)
}

// DeleteAllNewrelicAlertsConditions deletes all NewrelicAlertsCondition records in the database with optional filters
func DeleteAllNewrelicAlertsConditions(ctx context.Context, db *sql.DB, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(NewrelicAlertsConditionTableName),
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

// DeleteAllNewrelicAlertsConditionsTx deletes all NewrelicAlertsCondition records in the database with optional filters using the provided transaction
func DeleteAllNewrelicAlertsConditionsTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(NewrelicAlertsConditionTableName),
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

// DBDelete will delete this NewrelicAlertsCondition record in the database
func (t *NewrelicAlertsCondition) DBDelete(ctx context.Context, db *sql.DB) (bool, error) {
	q := "DELETE FROM `newrelic_alerts_condition` WHERE `id` = ?"
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

// DBDeleteTx will delete this NewrelicAlertsCondition record in the database using the provided transaction
func (t *NewrelicAlertsCondition) DBDeleteTx(ctx context.Context, tx *sql.Tx) (bool, error) {
	q := "DELETE FROM `newrelic_alerts_condition` WHERE `id` = ?"
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

// DBUpdate will update the NewrelicAlertsCondition record in the database
func (t *NewrelicAlertsCondition) DBUpdate(ctx context.Context, db *sql.DB) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `newrelic_alerts_condition` SET `checksum`=?,`customer_id`=?,`policy_ext_id`=?,`policy_id`=?,`ext_id`=?,`type`=?,`name`=?,`enabled`=?,`metric`=?,`terms`=? WHERE `id`=?"
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLInt64(t.PolicyExtID),
		orm.ToSQLString(t.PolicyID),
		orm.ToSQLInt64(t.ExtID),
		orm.ToSQLString(t.Type),
		orm.ToSQLString(t.Name),
		orm.ToSQLBool(t.Enabled),
		orm.ToSQLString(t.Metric),
		orm.ToSQLString(t.Terms),
		orm.ToSQLString(t.ID),
	)
}

// DBUpdateTx will update the NewrelicAlertsCondition record in the database using the provided transaction
func (t *NewrelicAlertsCondition) DBUpdateTx(ctx context.Context, tx *sql.Tx) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `newrelic_alerts_condition` SET `checksum`=?,`customer_id`=?,`policy_ext_id`=?,`policy_id`=?,`ext_id`=?,`type`=?,`name`=?,`enabled`=?,`metric`=?,`terms`=? WHERE `id`=?"
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLInt64(t.PolicyExtID),
		orm.ToSQLString(t.PolicyID),
		orm.ToSQLInt64(t.ExtID),
		orm.ToSQLString(t.Type),
		orm.ToSQLString(t.Name),
		orm.ToSQLBool(t.Enabled),
		orm.ToSQLString(t.Metric),
		orm.ToSQLString(t.Terms),
		orm.ToSQLString(t.ID),
	)
}

// DBUpsert will upsert the NewrelicAlertsCondition record in the database
func (t *NewrelicAlertsCondition) DBUpsert(ctx context.Context, db *sql.DB, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `newrelic_alerts_condition` (`newrelic_alerts_condition`.`id`,`newrelic_alerts_condition`.`checksum`,`newrelic_alerts_condition`.`customer_id`,`newrelic_alerts_condition`.`policy_ext_id`,`newrelic_alerts_condition`.`policy_id`,`newrelic_alerts_condition`.`ext_id`,`newrelic_alerts_condition`.`type`,`newrelic_alerts_condition`.`name`,`newrelic_alerts_condition`.`enabled`,`newrelic_alerts_condition`.`metric`,`newrelic_alerts_condition`.`terms`) VALUES (?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `newrelic_alerts_condition` (`newrelic_alerts_condition`.`id`,`newrelic_alerts_condition`.`checksum`,`newrelic_alerts_condition`.`customer_id`,`newrelic_alerts_condition`.`policy_ext_id`,`newrelic_alerts_condition`.`policy_id`,`newrelic_alerts_condition`.`ext_id`,`newrelic_alerts_condition`.`type`,`newrelic_alerts_condition`.`name`,`newrelic_alerts_condition`.`enabled`,`newrelic_alerts_condition`.`metric`,`newrelic_alerts_condition`.`terms`) VALUES (?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`customer_id`=VALUES(`customer_id`),`policy_ext_id`=VALUES(`policy_ext_id`),`policy_id`=VALUES(`policy_id`),`ext_id`=VALUES(`ext_id`),`type`=VALUES(`type`),`name`=VALUES(`name`),`enabled`=VALUES(`enabled`),`metric`=VALUES(`metric`),`terms`=VALUES(`terms`)"
	}
	r, err := db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLInt64(t.PolicyExtID),
		orm.ToSQLString(t.PolicyID),
		orm.ToSQLInt64(t.ExtID),
		orm.ToSQLString(t.Type),
		orm.ToSQLString(t.Name),
		orm.ToSQLBool(t.Enabled),
		orm.ToSQLString(t.Metric),
		orm.ToSQLString(t.Terms),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBUpsertTx will upsert the NewrelicAlertsCondition record in the database using the provided transaction
func (t *NewrelicAlertsCondition) DBUpsertTx(ctx context.Context, tx *sql.Tx, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `newrelic_alerts_condition` (`newrelic_alerts_condition`.`id`,`newrelic_alerts_condition`.`checksum`,`newrelic_alerts_condition`.`customer_id`,`newrelic_alerts_condition`.`policy_ext_id`,`newrelic_alerts_condition`.`policy_id`,`newrelic_alerts_condition`.`ext_id`,`newrelic_alerts_condition`.`type`,`newrelic_alerts_condition`.`name`,`newrelic_alerts_condition`.`enabled`,`newrelic_alerts_condition`.`metric`,`newrelic_alerts_condition`.`terms`) VALUES (?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `newrelic_alerts_condition` (`newrelic_alerts_condition`.`id`,`newrelic_alerts_condition`.`checksum`,`newrelic_alerts_condition`.`customer_id`,`newrelic_alerts_condition`.`policy_ext_id`,`newrelic_alerts_condition`.`policy_id`,`newrelic_alerts_condition`.`ext_id`,`newrelic_alerts_condition`.`type`,`newrelic_alerts_condition`.`name`,`newrelic_alerts_condition`.`enabled`,`newrelic_alerts_condition`.`metric`,`newrelic_alerts_condition`.`terms`) VALUES (?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`customer_id`=VALUES(`customer_id`),`policy_ext_id`=VALUES(`policy_ext_id`),`policy_id`=VALUES(`policy_id`),`ext_id`=VALUES(`ext_id`),`type`=VALUES(`type`),`name`=VALUES(`name`),`enabled`=VALUES(`enabled`),`metric`=VALUES(`metric`),`terms`=VALUES(`terms`)"
	}
	r, err := tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLInt64(t.PolicyExtID),
		orm.ToSQLString(t.PolicyID),
		orm.ToSQLInt64(t.ExtID),
		orm.ToSQLString(t.Type),
		orm.ToSQLString(t.Name),
		orm.ToSQLBool(t.Enabled),
		orm.ToSQLString(t.Metric),
		orm.ToSQLString(t.Terms),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBFindOne will find a NewrelicAlertsCondition record in the database with the primary key
func (t *NewrelicAlertsCondition) DBFindOne(ctx context.Context, db *sql.DB, value string) (bool, error) {
	q := "SELECT `newrelic_alerts_condition`.`id`,`newrelic_alerts_condition`.`checksum`,`newrelic_alerts_condition`.`customer_id`,`newrelic_alerts_condition`.`policy_ext_id`,`newrelic_alerts_condition`.`policy_id`,`newrelic_alerts_condition`.`ext_id`,`newrelic_alerts_condition`.`type`,`newrelic_alerts_condition`.`name`,`newrelic_alerts_condition`.`enabled`,`newrelic_alerts_condition`.`metric`,`newrelic_alerts_condition`.`terms` FROM `newrelic_alerts_condition` WHERE `id` = ? LIMIT 1"
	row := db.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CustomerID sql.NullString
	var _PolicyExtID sql.NullInt64
	var _PolicyID sql.NullString
	var _ExtID sql.NullInt64
	var _Type sql.NullString
	var _Name sql.NullString
	var _Enabled sql.NullBool
	var _Metric sql.NullString
	var _Terms sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_PolicyExtID,
		&_PolicyID,
		&_ExtID,
		&_Type,
		&_Name,
		&_Enabled,
		&_Metric,
		&_Terms,
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
	if _PolicyExtID.Valid {
		t.SetPolicyExtID(_PolicyExtID.Int64)
	}
	if _PolicyID.Valid {
		t.SetPolicyID(_PolicyID.String)
	}
	if _ExtID.Valid {
		t.SetExtID(_ExtID.Int64)
	}
	if _Type.Valid {
		t.SetType(_Type.String)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _Enabled.Valid {
		t.SetEnabled(_Enabled.Bool)
	}
	if _Metric.Valid {
		t.SetMetric(_Metric.String)
	}
	if _Terms.Valid {
		t.SetTerms(_Terms.String)
	}
	return true, nil
}

// DBFindOneTx will find a NewrelicAlertsCondition record in the database with the primary key using the provided transaction
func (t *NewrelicAlertsCondition) DBFindOneTx(ctx context.Context, tx *sql.Tx, value string) (bool, error) {
	q := "SELECT `newrelic_alerts_condition`.`id`,`newrelic_alerts_condition`.`checksum`,`newrelic_alerts_condition`.`customer_id`,`newrelic_alerts_condition`.`policy_ext_id`,`newrelic_alerts_condition`.`policy_id`,`newrelic_alerts_condition`.`ext_id`,`newrelic_alerts_condition`.`type`,`newrelic_alerts_condition`.`name`,`newrelic_alerts_condition`.`enabled`,`newrelic_alerts_condition`.`metric`,`newrelic_alerts_condition`.`terms` FROM `newrelic_alerts_condition` WHERE `id` = ? LIMIT 1"
	row := tx.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CustomerID sql.NullString
	var _PolicyExtID sql.NullInt64
	var _PolicyID sql.NullString
	var _ExtID sql.NullInt64
	var _Type sql.NullString
	var _Name sql.NullString
	var _Enabled sql.NullBool
	var _Metric sql.NullString
	var _Terms sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_PolicyExtID,
		&_PolicyID,
		&_ExtID,
		&_Type,
		&_Name,
		&_Enabled,
		&_Metric,
		&_Terms,
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
	if _PolicyExtID.Valid {
		t.SetPolicyExtID(_PolicyExtID.Int64)
	}
	if _PolicyID.Valid {
		t.SetPolicyID(_PolicyID.String)
	}
	if _ExtID.Valid {
		t.SetExtID(_ExtID.Int64)
	}
	if _Type.Valid {
		t.SetType(_Type.String)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _Enabled.Valid {
		t.SetEnabled(_Enabled.Bool)
	}
	if _Metric.Valid {
		t.SetMetric(_Metric.String)
	}
	if _Terms.Valid {
		t.SetTerms(_Terms.String)
	}
	return true, nil
}

// FindNewrelicAlertsConditions will find a NewrelicAlertsCondition record in the database with the provided parameters
func FindNewrelicAlertsConditions(ctx context.Context, db *sql.DB, _params ...interface{}) ([]*NewrelicAlertsCondition, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("customer_id"),
		orm.Column("policy_ext_id"),
		orm.Column("policy_id"),
		orm.Column("ext_id"),
		orm.Column("type"),
		orm.Column("name"),
		orm.Column("enabled"),
		orm.Column("metric"),
		orm.Column("terms"),
		orm.Table(NewrelicAlertsConditionTableName),
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
	results := make([]*NewrelicAlertsCondition, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _PolicyExtID sql.NullInt64
		var _PolicyID sql.NullString
		var _ExtID sql.NullInt64
		var _Type sql.NullString
		var _Name sql.NullString
		var _Enabled sql.NullBool
		var _Metric sql.NullString
		var _Terms sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_PolicyExtID,
			&_PolicyID,
			&_ExtID,
			&_Type,
			&_Name,
			&_Enabled,
			&_Metric,
			&_Terms,
		)
		if err != nil {
			return nil, err
		}
		t := &NewrelicAlertsCondition{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _PolicyExtID.Valid {
			t.SetPolicyExtID(_PolicyExtID.Int64)
		}
		if _PolicyID.Valid {
			t.SetPolicyID(_PolicyID.String)
		}
		if _ExtID.Valid {
			t.SetExtID(_ExtID.Int64)
		}
		if _Type.Valid {
			t.SetType(_Type.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _Enabled.Valid {
			t.SetEnabled(_Enabled.Bool)
		}
		if _Metric.Valid {
			t.SetMetric(_Metric.String)
		}
		if _Terms.Valid {
			t.SetTerms(_Terms.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindNewrelicAlertsConditionsTx will find a NewrelicAlertsCondition record in the database with the provided parameters using the provided transaction
func FindNewrelicAlertsConditionsTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) ([]*NewrelicAlertsCondition, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("customer_id"),
		orm.Column("policy_ext_id"),
		orm.Column("policy_id"),
		orm.Column("ext_id"),
		orm.Column("type"),
		orm.Column("name"),
		orm.Column("enabled"),
		orm.Column("metric"),
		orm.Column("terms"),
		orm.Table(NewrelicAlertsConditionTableName),
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
	results := make([]*NewrelicAlertsCondition, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _PolicyExtID sql.NullInt64
		var _PolicyID sql.NullString
		var _ExtID sql.NullInt64
		var _Type sql.NullString
		var _Name sql.NullString
		var _Enabled sql.NullBool
		var _Metric sql.NullString
		var _Terms sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_PolicyExtID,
			&_PolicyID,
			&_ExtID,
			&_Type,
			&_Name,
			&_Enabled,
			&_Metric,
			&_Terms,
		)
		if err != nil {
			return nil, err
		}
		t := &NewrelicAlertsCondition{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _PolicyExtID.Valid {
			t.SetPolicyExtID(_PolicyExtID.Int64)
		}
		if _PolicyID.Valid {
			t.SetPolicyID(_PolicyID.String)
		}
		if _ExtID.Valid {
			t.SetExtID(_ExtID.Int64)
		}
		if _Type.Valid {
			t.SetType(_Type.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _Enabled.Valid {
			t.SetEnabled(_Enabled.Bool)
		}
		if _Metric.Valid {
			t.SetMetric(_Metric.String)
		}
		if _Terms.Valid {
			t.SetTerms(_Terms.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// DBFind will find a NewrelicAlertsCondition record in the database with the provided parameters
func (t *NewrelicAlertsCondition) DBFind(ctx context.Context, db *sql.DB, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("customer_id"),
		orm.Column("policy_ext_id"),
		orm.Column("policy_id"),
		orm.Column("ext_id"),
		orm.Column("type"),
		orm.Column("name"),
		orm.Column("enabled"),
		orm.Column("metric"),
		orm.Column("terms"),
		orm.Table(NewrelicAlertsConditionTableName),
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
	var _PolicyExtID sql.NullInt64
	var _PolicyID sql.NullString
	var _ExtID sql.NullInt64
	var _Type sql.NullString
	var _Name sql.NullString
	var _Enabled sql.NullBool
	var _Metric sql.NullString
	var _Terms sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_PolicyExtID,
		&_PolicyID,
		&_ExtID,
		&_Type,
		&_Name,
		&_Enabled,
		&_Metric,
		&_Terms,
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
	if _PolicyExtID.Valid {
		t.SetPolicyExtID(_PolicyExtID.Int64)
	}
	if _PolicyID.Valid {
		t.SetPolicyID(_PolicyID.String)
	}
	if _ExtID.Valid {
		t.SetExtID(_ExtID.Int64)
	}
	if _Type.Valid {
		t.SetType(_Type.String)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _Enabled.Valid {
		t.SetEnabled(_Enabled.Bool)
	}
	if _Metric.Valid {
		t.SetMetric(_Metric.String)
	}
	if _Terms.Valid {
		t.SetTerms(_Terms.String)
	}
	return true, nil
}

// DBFindTx will find a NewrelicAlertsCondition record in the database with the provided parameters using the provided transaction
func (t *NewrelicAlertsCondition) DBFindTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("customer_id"),
		orm.Column("policy_ext_id"),
		orm.Column("policy_id"),
		orm.Column("ext_id"),
		orm.Column("type"),
		orm.Column("name"),
		orm.Column("enabled"),
		orm.Column("metric"),
		orm.Column("terms"),
		orm.Table(NewrelicAlertsConditionTableName),
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
	var _PolicyExtID sql.NullInt64
	var _PolicyID sql.NullString
	var _ExtID sql.NullInt64
	var _Type sql.NullString
	var _Name sql.NullString
	var _Enabled sql.NullBool
	var _Metric sql.NullString
	var _Terms sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_PolicyExtID,
		&_PolicyID,
		&_ExtID,
		&_Type,
		&_Name,
		&_Enabled,
		&_Metric,
		&_Terms,
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
	if _PolicyExtID.Valid {
		t.SetPolicyExtID(_PolicyExtID.Int64)
	}
	if _PolicyID.Valid {
		t.SetPolicyID(_PolicyID.String)
	}
	if _ExtID.Valid {
		t.SetExtID(_ExtID.Int64)
	}
	if _Type.Valid {
		t.SetType(_Type.String)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _Enabled.Valid {
		t.SetEnabled(_Enabled.Bool)
	}
	if _Metric.Valid {
		t.SetMetric(_Metric.String)
	}
	if _Terms.Valid {
		t.SetTerms(_Terms.String)
	}
	return true, nil
}

// CountNewrelicAlertsConditions will find the count of NewrelicAlertsCondition records in the database
func CountNewrelicAlertsConditions(ctx context.Context, db *sql.DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(NewrelicAlertsConditionTableName),
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

// CountNewrelicAlertsConditionsTx will find the count of NewrelicAlertsCondition records in the database using the provided transaction
func CountNewrelicAlertsConditionsTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(NewrelicAlertsConditionTableName),
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

// DBCount will find the count of NewrelicAlertsCondition records in the database
func (t *NewrelicAlertsCondition) DBCount(ctx context.Context, db *sql.DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(NewrelicAlertsConditionTableName),
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

// DBCountTx will find the count of NewrelicAlertsCondition records in the database using the provided transaction
func (t *NewrelicAlertsCondition) DBCountTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(NewrelicAlertsConditionTableName),
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

// DBExists will return true if the NewrelicAlertsCondition record exists in the database
func (t *NewrelicAlertsCondition) DBExists(ctx context.Context, db *sql.DB) (bool, error) {
	q := "SELECT `id` FROM `newrelic_alerts_condition` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := db.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// DBExistsTx will return true if the NewrelicAlertsCondition record exists in the database using the provided transaction
func (t *NewrelicAlertsCondition) DBExistsTx(ctx context.Context, tx *sql.Tx) (bool, error) {
	q := "SELECT `id` FROM `newrelic_alerts_condition` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := tx.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// PrimaryKeyColumn returns the column name for the primary key
func (t *NewrelicAlertsCondition) PrimaryKeyColumn() string {
	return NewrelicAlertsConditionColumnID
}

// PrimaryKeyColumnType returns the primary key column Go type as a string
func (t *NewrelicAlertsCondition) PrimaryKeyColumnType() string {
	return "string"
}

// PrimaryKey returns the primary key column value
func (t *NewrelicAlertsCondition) PrimaryKey() interface{} {
	return t.ID
}
