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

var _ Model = (*NewrelicAlertsPolicy)(nil)
var _ CSVWriter = (*NewrelicAlertsPolicy)(nil)
var _ JSONWriter = (*NewrelicAlertsPolicy)(nil)
var _ Checksum = (*NewrelicAlertsPolicy)(nil)

// NewrelicAlertsPolicyTableName is the name of the table in SQL
const NewrelicAlertsPolicyTableName = "newrelic_alerts_policy"

var NewrelicAlertsPolicyColumns = []string{
	"id",
	"checksum",
	"customer_id",
	"ext_id",
	"incident_preference",
	"name",
	"created_at",
	"updated_at",
}

// NewrelicAlertsPolicy table
type NewrelicAlertsPolicy struct {
	Checksum           *string `json:"checksum,omitempty"`
	CreatedAt          int64   `json:"created_at"`
	CustomerID         string  `json:"customer_id"`
	ExtID              int64   `json:"ext_id"`
	ID                 string  `json:"id"`
	IncidentPreference string  `json:"incident_preference"`
	Name               string  `json:"name"`
	UpdatedAt          int64   `json:"updated_at"`
}

// TableName returns the SQL table name for NewrelicAlertsPolicy and satifies the Model interface
func (t *NewrelicAlertsPolicy) TableName() string {
	return NewrelicAlertsPolicyTableName
}

// ToCSV will serialize the NewrelicAlertsPolicy instance to a CSV compatible array of strings
func (t *NewrelicAlertsPolicy) ToCSV() []string {
	return []string{
		t.ID,
		t.CalculateChecksum(),
		t.CustomerID,
		toCSVString(t.ExtID),
		t.IncidentPreference,
		t.Name,
		toCSVString(t.CreatedAt),
		toCSVString(t.UpdatedAt),
	}
}

// WriteCSV will serialize the NewrelicAlertsPolicy instance to the writer as CSV and satisfies the CSVWriter interface
func (t *NewrelicAlertsPolicy) WriteCSV(w *csv.Writer) error {
	return w.Write(t.ToCSV())
}

// WriteJSON will serialize the NewrelicAlertsPolicy instance to the writer as JSON and satisfies the JSONWriter interface
func (t *NewrelicAlertsPolicy) WriteJSON(w io.Writer, indent ...bool) error {
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

// NewNewrelicAlertsPolicyReader creates a JSON reader which can read in NewrelicAlertsPolicy objects serialized as JSON either as an array, single object or json new lines
// and writes each NewrelicAlertsPolicy to the channel provided
func NewNewrelicAlertsPolicyReader(r io.Reader, ch chan<- NewrelicAlertsPolicy) error {
	return orm.Deserialize(r, func(buf json.RawMessage) error {
		dec := json.NewDecoder(bytes.NewBuffer(buf))
		e := NewrelicAlertsPolicy{}
		if err := dec.Decode(&e); err != nil {
			return err
		}
		ch <- e
		return nil
	})
}

// NewCSVNewrelicAlertsPolicyReaderDir will read the reader as CSV and emit each record to the channel provided
func NewCSVNewrelicAlertsPolicyReader(r io.Reader, ch chan<- NewrelicAlertsPolicy) error {
	cr := csv.NewReader(r)
	for {
		record, err := cr.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		ch <- NewrelicAlertsPolicy{
			ID:                 record[0],
			Checksum:           fromStringPointer(record[1]),
			CustomerID:         record[2],
			ExtID:              fromCSVInt64(record[3]),
			IncidentPreference: record[4],
			Name:               record[5],
			CreatedAt:          fromCSVInt64(record[6]),
			UpdatedAt:          fromCSVInt64(record[7]),
		}
	}
	return nil
}

// NewCSVNewrelicAlertsPolicyReaderFile will read the file as a CSV and emit each record to the channel provided
func NewCSVNewrelicAlertsPolicyReaderFile(fp string, ch chan<- NewrelicAlertsPolicy) error {
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
	return NewCSVNewrelicAlertsPolicyReader(fc, ch)
}

// NewCSVNewrelicAlertsPolicyReaderDir will read the newrelic_alerts_policy.csv.gz file as a CSV and emit each record to the channel provided
func NewCSVNewrelicAlertsPolicyReaderDir(dir string, ch chan<- NewrelicAlertsPolicy) error {
	return NewCSVNewrelicAlertsPolicyReaderFile(filepath.Join(dir, "newrelic_alerts_policy.csv.gz"), ch)
}

// NewrelicAlertsPolicyCSVDeduper is a function callback which takes the existing value (a) and the new value (b)
// and the return value should be the one to use (or a new one, if applicable). return nil
// to skip processing of this record
type NewrelicAlertsPolicyCSVDeduper func(a NewrelicAlertsPolicy, b NewrelicAlertsPolicy) *NewrelicAlertsPolicy

// NewrelicAlertsPolicyCSVDedupeDisabled is set on whether the CSV writer should de-dupe values by key
var NewrelicAlertsPolicyCSVDedupeDisabled bool

// NewNewrelicAlertsPolicyCSVWriterSize creates a batch writer that will write each NewrelicAlertsPolicy into a CSV file
// this method will automatically de-duplicate entries using the primary key. if the checksum
// for a newer item with the same primary key doesn't match a previously sent item, the newer
// one will be used
func NewNewrelicAlertsPolicyCSVWriterSize(w io.Writer, size int, dedupers ...NewrelicAlertsPolicyCSVDeduper) (chan NewrelicAlertsPolicy, chan bool, error) {
	cw := csv.NewWriter(w)
	ch := make(chan NewrelicAlertsPolicy, size)
	done := make(chan bool)
	go func() {
		defer func() { done <- true }()
		dodedupe := !NewrelicAlertsPolicyCSVDedupeDisabled
		var kv map[string]*NewrelicAlertsPolicy
		var deduper NewrelicAlertsPolicyCSVDeduper
		if dedupers != nil && len(dedupers) > 0 {
			deduper = dedupers[0]
			dodedupe = true
		}
		if dodedupe {
			kv = make(map[string]*NewrelicAlertsPolicy)
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

// NewrelicAlertsPolicyCSVDefaultSize is the default channel buffer size if not provided
var NewrelicAlertsPolicyCSVDefaultSize = 100

// NewNewrelicAlertsPolicyCSVWriter creates a batch writer that will write each NewrelicAlertsPolicy into a CSV file
func NewNewrelicAlertsPolicyCSVWriter(w io.Writer, dedupers ...NewrelicAlertsPolicyCSVDeduper) (chan NewrelicAlertsPolicy, chan bool, error) {
	return NewNewrelicAlertsPolicyCSVWriterSize(w, NewrelicAlertsPolicyCSVDefaultSize, dedupers...)
}

// NewNewrelicAlertsPolicyCSVWriterDir creates a batch writer that will write each NewrelicAlertsPolicy into a CSV file named newrelic_alerts_policy.csv.gz in dir
func NewNewrelicAlertsPolicyCSVWriterDir(dir string, dedupers ...NewrelicAlertsPolicyCSVDeduper) (chan NewrelicAlertsPolicy, chan bool, error) {
	return NewNewrelicAlertsPolicyCSVWriterFile(filepath.Join(dir, "newrelic_alerts_policy.csv.gz"), dedupers...)
}

// NewNewrelicAlertsPolicyCSVWriterFile creates a batch writer that will write each NewrelicAlertsPolicy into a CSV file
func NewNewrelicAlertsPolicyCSVWriterFile(fn string, dedupers ...NewrelicAlertsPolicyCSVDeduper) (chan NewrelicAlertsPolicy, chan bool, error) {
	f, err := os.Create(fn)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening CSV file %s. %v", fn, err)
	}
	var fc io.WriteCloser = f
	if filepath.Ext(fn) == ".gz" {
		w, _ := gzip.NewWriterLevel(f, gzip.BestCompression)
		fc = w
	}
	ch, done, err := NewNewrelicAlertsPolicyCSVWriter(fc, dedupers...)
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

type NewrelicAlertsPolicyDBAction func(ctx context.Context, db *sql.DB, record NewrelicAlertsPolicy) error

// NewNewrelicAlertsPolicyDBWriterSize creates a DB writer that will write each issue into the DB
func NewNewrelicAlertsPolicyDBWriterSize(ctx context.Context, db *sql.DB, errors chan<- error, size int, actions ...NewrelicAlertsPolicyDBAction) (chan NewrelicAlertsPolicy, chan bool, error) {
	ch := make(chan NewrelicAlertsPolicy, size)
	done := make(chan bool)
	var action NewrelicAlertsPolicyDBAction
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

// NewNewrelicAlertsPolicyDBWriter creates a DB writer that will write each issue into the DB
func NewNewrelicAlertsPolicyDBWriter(ctx context.Context, db *sql.DB, errors chan<- error, actions ...NewrelicAlertsPolicyDBAction) (chan NewrelicAlertsPolicy, chan bool, error) {
	return NewNewrelicAlertsPolicyDBWriterSize(ctx, db, errors, 100, actions...)
}

// NewrelicAlertsPolicyColumnID is the ID SQL column name for the NewrelicAlertsPolicy table
const NewrelicAlertsPolicyColumnID = "id"

// NewrelicAlertsPolicyEscapedColumnID is the escaped ID SQL column name for the NewrelicAlertsPolicy table
const NewrelicAlertsPolicyEscapedColumnID = "`id`"

// NewrelicAlertsPolicyColumnChecksum is the Checksum SQL column name for the NewrelicAlertsPolicy table
const NewrelicAlertsPolicyColumnChecksum = "checksum"

// NewrelicAlertsPolicyEscapedColumnChecksum is the escaped Checksum SQL column name for the NewrelicAlertsPolicy table
const NewrelicAlertsPolicyEscapedColumnChecksum = "`checksum`"

// NewrelicAlertsPolicyColumnCustomerID is the CustomerID SQL column name for the NewrelicAlertsPolicy table
const NewrelicAlertsPolicyColumnCustomerID = "customer_id"

// NewrelicAlertsPolicyEscapedColumnCustomerID is the escaped CustomerID SQL column name for the NewrelicAlertsPolicy table
const NewrelicAlertsPolicyEscapedColumnCustomerID = "`customer_id`"

// NewrelicAlertsPolicyColumnExtID is the ExtID SQL column name for the NewrelicAlertsPolicy table
const NewrelicAlertsPolicyColumnExtID = "ext_id"

// NewrelicAlertsPolicyEscapedColumnExtID is the escaped ExtID SQL column name for the NewrelicAlertsPolicy table
const NewrelicAlertsPolicyEscapedColumnExtID = "`ext_id`"

// NewrelicAlertsPolicyColumnIncidentPreference is the IncidentPreference SQL column name for the NewrelicAlertsPolicy table
const NewrelicAlertsPolicyColumnIncidentPreference = "incident_preference"

// NewrelicAlertsPolicyEscapedColumnIncidentPreference is the escaped IncidentPreference SQL column name for the NewrelicAlertsPolicy table
const NewrelicAlertsPolicyEscapedColumnIncidentPreference = "`incident_preference`"

// NewrelicAlertsPolicyColumnName is the Name SQL column name for the NewrelicAlertsPolicy table
const NewrelicAlertsPolicyColumnName = "name"

// NewrelicAlertsPolicyEscapedColumnName is the escaped Name SQL column name for the NewrelicAlertsPolicy table
const NewrelicAlertsPolicyEscapedColumnName = "`name`"

// NewrelicAlertsPolicyColumnCreatedAt is the CreatedAt SQL column name for the NewrelicAlertsPolicy table
const NewrelicAlertsPolicyColumnCreatedAt = "created_at"

// NewrelicAlertsPolicyEscapedColumnCreatedAt is the escaped CreatedAt SQL column name for the NewrelicAlertsPolicy table
const NewrelicAlertsPolicyEscapedColumnCreatedAt = "`created_at`"

// NewrelicAlertsPolicyColumnUpdatedAt is the UpdatedAt SQL column name for the NewrelicAlertsPolicy table
const NewrelicAlertsPolicyColumnUpdatedAt = "updated_at"

// NewrelicAlertsPolicyEscapedColumnUpdatedAt is the escaped UpdatedAt SQL column name for the NewrelicAlertsPolicy table
const NewrelicAlertsPolicyEscapedColumnUpdatedAt = "`updated_at`"

// GetID will return the NewrelicAlertsPolicy ID value
func (t *NewrelicAlertsPolicy) GetID() string {
	return t.ID
}

// SetID will set the NewrelicAlertsPolicy ID value
func (t *NewrelicAlertsPolicy) SetID(v string) {
	t.ID = v
}

// FindNewrelicAlertsPolicyByID will find a NewrelicAlertsPolicy by ID
func FindNewrelicAlertsPolicyByID(ctx context.Context, db *sql.DB, value string) (*NewrelicAlertsPolicy, error) {
	q := "SELECT `newrelic_alerts_policy`.`id`,`newrelic_alerts_policy`.`checksum`,`newrelic_alerts_policy`.`customer_id`,`newrelic_alerts_policy`.`ext_id`,`newrelic_alerts_policy`.`incident_preference`,`newrelic_alerts_policy`.`name`,`newrelic_alerts_policy`.`created_at`,`newrelic_alerts_policy`.`updated_at` FROM `newrelic_alerts_policy` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CustomerID sql.NullString
	var _ExtID sql.NullInt64
	var _IncidentPreference sql.NullString
	var _Name sql.NullString
	var _CreatedAt sql.NullInt64
	var _UpdatedAt sql.NullInt64
	err := db.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_ExtID,
		&_IncidentPreference,
		&_Name,
		&_CreatedAt,
		&_UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &NewrelicAlertsPolicy{}
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
	if _IncidentPreference.Valid {
		t.SetIncidentPreference(_IncidentPreference.String)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _CreatedAt.Valid {
		t.SetCreatedAt(_CreatedAt.Int64)
	}
	if _UpdatedAt.Valid {
		t.SetUpdatedAt(_UpdatedAt.Int64)
	}
	return t, nil
}

// FindNewrelicAlertsPolicyByIDTx will find a NewrelicAlertsPolicy by ID using the provided transaction
func FindNewrelicAlertsPolicyByIDTx(ctx context.Context, tx *sql.Tx, value string) (*NewrelicAlertsPolicy, error) {
	q := "SELECT `newrelic_alerts_policy`.`id`,`newrelic_alerts_policy`.`checksum`,`newrelic_alerts_policy`.`customer_id`,`newrelic_alerts_policy`.`ext_id`,`newrelic_alerts_policy`.`incident_preference`,`newrelic_alerts_policy`.`name`,`newrelic_alerts_policy`.`created_at`,`newrelic_alerts_policy`.`updated_at` FROM `newrelic_alerts_policy` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CustomerID sql.NullString
	var _ExtID sql.NullInt64
	var _IncidentPreference sql.NullString
	var _Name sql.NullString
	var _CreatedAt sql.NullInt64
	var _UpdatedAt sql.NullInt64
	err := tx.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_ExtID,
		&_IncidentPreference,
		&_Name,
		&_CreatedAt,
		&_UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &NewrelicAlertsPolicy{}
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
	if _IncidentPreference.Valid {
		t.SetIncidentPreference(_IncidentPreference.String)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _CreatedAt.Valid {
		t.SetCreatedAt(_CreatedAt.Int64)
	}
	if _UpdatedAt.Valid {
		t.SetUpdatedAt(_UpdatedAt.Int64)
	}
	return t, nil
}

// GetChecksum will return the NewrelicAlertsPolicy Checksum value
func (t *NewrelicAlertsPolicy) GetChecksum() string {
	if t.Checksum == nil {
		return ""
	}
	return *t.Checksum
}

// SetChecksum will set the NewrelicAlertsPolicy Checksum value
func (t *NewrelicAlertsPolicy) SetChecksum(v string) {
	t.Checksum = &v
}

// GetCustomerID will return the NewrelicAlertsPolicy CustomerID value
func (t *NewrelicAlertsPolicy) GetCustomerID() string {
	return t.CustomerID
}

// SetCustomerID will set the NewrelicAlertsPolicy CustomerID value
func (t *NewrelicAlertsPolicy) SetCustomerID(v string) {
	t.CustomerID = v
}

// FindNewrelicAlertsPoliciesByCustomerID will find all NewrelicAlertsPolicys by the CustomerID value
func FindNewrelicAlertsPoliciesByCustomerID(ctx context.Context, db *sql.DB, value string) ([]*NewrelicAlertsPolicy, error) {
	q := "SELECT `newrelic_alerts_policy`.`id`,`newrelic_alerts_policy`.`checksum`,`newrelic_alerts_policy`.`customer_id`,`newrelic_alerts_policy`.`ext_id`,`newrelic_alerts_policy`.`incident_preference`,`newrelic_alerts_policy`.`name`,`newrelic_alerts_policy`.`created_at`,`newrelic_alerts_policy`.`updated_at` FROM `newrelic_alerts_policy` WHERE `customer_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*NewrelicAlertsPolicy, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _ExtID sql.NullInt64
		var _IncidentPreference sql.NullString
		var _Name sql.NullString
		var _CreatedAt sql.NullInt64
		var _UpdatedAt sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_ExtID,
			&_IncidentPreference,
			&_Name,
			&_CreatedAt,
			&_UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		t := &NewrelicAlertsPolicy{}
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
		if _IncidentPreference.Valid {
			t.SetIncidentPreference(_IncidentPreference.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _CreatedAt.Valid {
			t.SetCreatedAt(_CreatedAt.Int64)
		}
		if _UpdatedAt.Valid {
			t.SetUpdatedAt(_UpdatedAt.Int64)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindNewrelicAlertsPoliciesByCustomerIDTx will find all NewrelicAlertsPolicys by the CustomerID value using the provided transaction
func FindNewrelicAlertsPoliciesByCustomerIDTx(ctx context.Context, tx *sql.Tx, value string) ([]*NewrelicAlertsPolicy, error) {
	q := "SELECT `newrelic_alerts_policy`.`id`,`newrelic_alerts_policy`.`checksum`,`newrelic_alerts_policy`.`customer_id`,`newrelic_alerts_policy`.`ext_id`,`newrelic_alerts_policy`.`incident_preference`,`newrelic_alerts_policy`.`name`,`newrelic_alerts_policy`.`created_at`,`newrelic_alerts_policy`.`updated_at` FROM `newrelic_alerts_policy` WHERE `customer_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*NewrelicAlertsPolicy, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _ExtID sql.NullInt64
		var _IncidentPreference sql.NullString
		var _Name sql.NullString
		var _CreatedAt sql.NullInt64
		var _UpdatedAt sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_ExtID,
			&_IncidentPreference,
			&_Name,
			&_CreatedAt,
			&_UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		t := &NewrelicAlertsPolicy{}
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
		if _IncidentPreference.Valid {
			t.SetIncidentPreference(_IncidentPreference.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _CreatedAt.Valid {
			t.SetCreatedAt(_CreatedAt.Int64)
		}
		if _UpdatedAt.Valid {
			t.SetUpdatedAt(_UpdatedAt.Int64)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetExtID will return the NewrelicAlertsPolicy ExtID value
func (t *NewrelicAlertsPolicy) GetExtID() int64 {
	return t.ExtID
}

// SetExtID will set the NewrelicAlertsPolicy ExtID value
func (t *NewrelicAlertsPolicy) SetExtID(v int64) {
	t.ExtID = v
}

// GetIncidentPreference will return the NewrelicAlertsPolicy IncidentPreference value
func (t *NewrelicAlertsPolicy) GetIncidentPreference() string {
	return t.IncidentPreference
}

// SetIncidentPreference will set the NewrelicAlertsPolicy IncidentPreference value
func (t *NewrelicAlertsPolicy) SetIncidentPreference(v string) {
	t.IncidentPreference = v
}

// GetName will return the NewrelicAlertsPolicy Name value
func (t *NewrelicAlertsPolicy) GetName() string {
	return t.Name
}

// SetName will set the NewrelicAlertsPolicy Name value
func (t *NewrelicAlertsPolicy) SetName(v string) {
	t.Name = v
}

// GetCreatedAt will return the NewrelicAlertsPolicy CreatedAt value
func (t *NewrelicAlertsPolicy) GetCreatedAt() int64 {
	return t.CreatedAt
}

// SetCreatedAt will set the NewrelicAlertsPolicy CreatedAt value
func (t *NewrelicAlertsPolicy) SetCreatedAt(v int64) {
	t.CreatedAt = v
}

// GetUpdatedAt will return the NewrelicAlertsPolicy UpdatedAt value
func (t *NewrelicAlertsPolicy) GetUpdatedAt() int64 {
	return t.UpdatedAt
}

// SetUpdatedAt will set the NewrelicAlertsPolicy UpdatedAt value
func (t *NewrelicAlertsPolicy) SetUpdatedAt(v int64) {
	t.UpdatedAt = v
}

func (t *NewrelicAlertsPolicy) toTimestamp(value time.Time) *timestamp.Timestamp {
	ts, _ := ptypes.TimestampProto(value)
	return ts
}

// DBCreateNewrelicAlertsPolicyTable will create the NewrelicAlertsPolicy table
func DBCreateNewrelicAlertsPolicyTable(ctx context.Context, db *sql.DB) error {
	q := "CREATE TABLE `newrelic_alerts_policy` (`id` VARCHAR(64) NOT NULL PRIMARY KEY,`checksum` CHAR(64),`customer_id` VARCHAR(64) NOT NULL,`ext_id`BIGINT(20) NOT NULL,`incident_preference` TEXT NOT NULL,`name` TEXT NOT NULL,`created_at` BIGINT(20) NOT NULL,`updated_at` BIGINT(20) NOT NULL,INDEX newrelic_alerts_policy_customer_id_index (`customer_id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBCreateNewrelicAlertsPolicyTableTx will create the NewrelicAlertsPolicy table using the provided transction
func DBCreateNewrelicAlertsPolicyTableTx(ctx context.Context, tx *sql.Tx) error {
	q := "CREATE TABLE `newrelic_alerts_policy` (`id` VARCHAR(64) NOT NULL PRIMARY KEY,`checksum` CHAR(64),`customer_id` VARCHAR(64) NOT NULL,`ext_id`BIGINT(20) NOT NULL,`incident_preference` TEXT NOT NULL,`name` TEXT NOT NULL,`created_at` BIGINT(20) NOT NULL,`updated_at` BIGINT(20) NOT NULL,INDEX newrelic_alerts_policy_customer_id_index (`customer_id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// DBDropNewrelicAlertsPolicyTable will drop the NewrelicAlertsPolicy table
func DBDropNewrelicAlertsPolicyTable(ctx context.Context, db *sql.DB) error {
	q := "DROP TABLE IF EXISTS `newrelic_alerts_policy`"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBDropNewrelicAlertsPolicyTableTx will drop the NewrelicAlertsPolicy table using the provided transaction
func DBDropNewrelicAlertsPolicyTableTx(ctx context.Context, tx *sql.Tx) error {
	q := "DROP TABLE IF EXISTS `newrelic_alerts_policy`"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// CalculateChecksum will calculate a checksum of the SHA1 of all field values
func (t *NewrelicAlertsPolicy) CalculateChecksum() string {
	return orm.HashStrings(
		orm.ToString(t.ID),
		orm.ToString(t.CustomerID),
		orm.ToString(t.ExtID),
		orm.ToString(t.IncidentPreference),
		orm.ToString(t.Name),
		orm.ToString(t.CreatedAt),
		orm.ToString(t.UpdatedAt),
	)
}

// DBCreate will create a new NewrelicAlertsPolicy record in the database
func (t *NewrelicAlertsPolicy) DBCreate(ctx context.Context, db *sql.DB) (sql.Result, error) {
	q := "INSERT INTO `newrelic_alerts_policy` (`newrelic_alerts_policy`.`id`,`newrelic_alerts_policy`.`checksum`,`newrelic_alerts_policy`.`customer_id`,`newrelic_alerts_policy`.`ext_id`,`newrelic_alerts_policy`.`incident_preference`,`newrelic_alerts_policy`.`name`,`newrelic_alerts_policy`.`created_at`,`newrelic_alerts_policy`.`updated_at`) VALUES (?,?,?,?,?,?,?,?)"
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
		orm.ToSQLString(t.IncidentPreference),
		orm.ToSQLString(t.Name),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
	)
}

// DBCreateTx will create a new NewrelicAlertsPolicy record in the database using the provided transaction
func (t *NewrelicAlertsPolicy) DBCreateTx(ctx context.Context, tx *sql.Tx) (sql.Result, error) {
	q := "INSERT INTO `newrelic_alerts_policy` (`newrelic_alerts_policy`.`id`,`newrelic_alerts_policy`.`checksum`,`newrelic_alerts_policy`.`customer_id`,`newrelic_alerts_policy`.`ext_id`,`newrelic_alerts_policy`.`incident_preference`,`newrelic_alerts_policy`.`name`,`newrelic_alerts_policy`.`created_at`,`newrelic_alerts_policy`.`updated_at`) VALUES (?,?,?,?,?,?,?,?)"
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
		orm.ToSQLString(t.IncidentPreference),
		orm.ToSQLString(t.Name),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
	)
}

// DBCreateIgnoreDuplicate will upsert the NewrelicAlertsPolicy record in the database
func (t *NewrelicAlertsPolicy) DBCreateIgnoreDuplicate(ctx context.Context, db *sql.DB) (sql.Result, error) {
	q := "INSERT INTO `newrelic_alerts_policy` (`newrelic_alerts_policy`.`id`,`newrelic_alerts_policy`.`checksum`,`newrelic_alerts_policy`.`customer_id`,`newrelic_alerts_policy`.`ext_id`,`newrelic_alerts_policy`.`incident_preference`,`newrelic_alerts_policy`.`name`,`newrelic_alerts_policy`.`created_at`,`newrelic_alerts_policy`.`updated_at`) VALUES (?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
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
		orm.ToSQLString(t.IncidentPreference),
		orm.ToSQLString(t.Name),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
	)
}

// DBCreateIgnoreDuplicateTx will upsert the NewrelicAlertsPolicy record in the database using the provided transaction
func (t *NewrelicAlertsPolicy) DBCreateIgnoreDuplicateTx(ctx context.Context, tx *sql.Tx) (sql.Result, error) {
	q := "INSERT INTO `newrelic_alerts_policy` (`newrelic_alerts_policy`.`id`,`newrelic_alerts_policy`.`checksum`,`newrelic_alerts_policy`.`customer_id`,`newrelic_alerts_policy`.`ext_id`,`newrelic_alerts_policy`.`incident_preference`,`newrelic_alerts_policy`.`name`,`newrelic_alerts_policy`.`created_at`,`newrelic_alerts_policy`.`updated_at`) VALUES (?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
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
		orm.ToSQLString(t.IncidentPreference),
		orm.ToSQLString(t.Name),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
	)
}

// DeleteAllNewrelicAlertsPolicies deletes all NewrelicAlertsPolicy records in the database with optional filters
func DeleteAllNewrelicAlertsPolicies(ctx context.Context, db *sql.DB, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(NewrelicAlertsPolicyTableName),
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

// DeleteAllNewrelicAlertsPoliciesTx deletes all NewrelicAlertsPolicy records in the database with optional filters using the provided transaction
func DeleteAllNewrelicAlertsPoliciesTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(NewrelicAlertsPolicyTableName),
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

// DBDelete will delete this NewrelicAlertsPolicy record in the database
func (t *NewrelicAlertsPolicy) DBDelete(ctx context.Context, db *sql.DB) (bool, error) {
	q := "DELETE FROM `newrelic_alerts_policy` WHERE `id` = ?"
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

// DBDeleteTx will delete this NewrelicAlertsPolicy record in the database using the provided transaction
func (t *NewrelicAlertsPolicy) DBDeleteTx(ctx context.Context, tx *sql.Tx) (bool, error) {
	q := "DELETE FROM `newrelic_alerts_policy` WHERE `id` = ?"
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

// DBUpdate will update the NewrelicAlertsPolicy record in the database
func (t *NewrelicAlertsPolicy) DBUpdate(ctx context.Context, db *sql.DB) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `newrelic_alerts_policy` SET `checksum`=?,`customer_id`=?,`ext_id`=?,`incident_preference`=?,`name`=?,`created_at`=?,`updated_at`=? WHERE `id`=?"
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLInt64(t.ExtID),
		orm.ToSQLString(t.IncidentPreference),
		orm.ToSQLString(t.Name),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
		orm.ToSQLString(t.ID),
	)
}

// DBUpdateTx will update the NewrelicAlertsPolicy record in the database using the provided transaction
func (t *NewrelicAlertsPolicy) DBUpdateTx(ctx context.Context, tx *sql.Tx) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `newrelic_alerts_policy` SET `checksum`=?,`customer_id`=?,`ext_id`=?,`incident_preference`=?,`name`=?,`created_at`=?,`updated_at`=? WHERE `id`=?"
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLInt64(t.ExtID),
		orm.ToSQLString(t.IncidentPreference),
		orm.ToSQLString(t.Name),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
		orm.ToSQLString(t.ID),
	)
}

// DBUpsert will upsert the NewrelicAlertsPolicy record in the database
func (t *NewrelicAlertsPolicy) DBUpsert(ctx context.Context, db *sql.DB, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `newrelic_alerts_policy` (`newrelic_alerts_policy`.`id`,`newrelic_alerts_policy`.`checksum`,`newrelic_alerts_policy`.`customer_id`,`newrelic_alerts_policy`.`ext_id`,`newrelic_alerts_policy`.`incident_preference`,`newrelic_alerts_policy`.`name`,`newrelic_alerts_policy`.`created_at`,`newrelic_alerts_policy`.`updated_at`) VALUES (?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `newrelic_alerts_policy` (`newrelic_alerts_policy`.`id`,`newrelic_alerts_policy`.`checksum`,`newrelic_alerts_policy`.`customer_id`,`newrelic_alerts_policy`.`ext_id`,`newrelic_alerts_policy`.`incident_preference`,`newrelic_alerts_policy`.`name`,`newrelic_alerts_policy`.`created_at`,`newrelic_alerts_policy`.`updated_at`) VALUES (?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`customer_id`=VALUES(`customer_id`),`ext_id`=VALUES(`ext_id`),`incident_preference`=VALUES(`incident_preference`),`name`=VALUES(`name`),`created_at`=VALUES(`created_at`),`updated_at`=VALUES(`updated_at`)"
	}
	r, err := db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLInt64(t.ExtID),
		orm.ToSQLString(t.IncidentPreference),
		orm.ToSQLString(t.Name),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBUpsertTx will upsert the NewrelicAlertsPolicy record in the database using the provided transaction
func (t *NewrelicAlertsPolicy) DBUpsertTx(ctx context.Context, tx *sql.Tx, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `newrelic_alerts_policy` (`newrelic_alerts_policy`.`id`,`newrelic_alerts_policy`.`checksum`,`newrelic_alerts_policy`.`customer_id`,`newrelic_alerts_policy`.`ext_id`,`newrelic_alerts_policy`.`incident_preference`,`newrelic_alerts_policy`.`name`,`newrelic_alerts_policy`.`created_at`,`newrelic_alerts_policy`.`updated_at`) VALUES (?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `newrelic_alerts_policy` (`newrelic_alerts_policy`.`id`,`newrelic_alerts_policy`.`checksum`,`newrelic_alerts_policy`.`customer_id`,`newrelic_alerts_policy`.`ext_id`,`newrelic_alerts_policy`.`incident_preference`,`newrelic_alerts_policy`.`name`,`newrelic_alerts_policy`.`created_at`,`newrelic_alerts_policy`.`updated_at`) VALUES (?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`customer_id`=VALUES(`customer_id`),`ext_id`=VALUES(`ext_id`),`incident_preference`=VALUES(`incident_preference`),`name`=VALUES(`name`),`created_at`=VALUES(`created_at`),`updated_at`=VALUES(`updated_at`)"
	}
	r, err := tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLInt64(t.ExtID),
		orm.ToSQLString(t.IncidentPreference),
		orm.ToSQLString(t.Name),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBFindOne will find a NewrelicAlertsPolicy record in the database with the primary key
func (t *NewrelicAlertsPolicy) DBFindOne(ctx context.Context, db *sql.DB, value string) (bool, error) {
	q := "SELECT `newrelic_alerts_policy`.`id`,`newrelic_alerts_policy`.`checksum`,`newrelic_alerts_policy`.`customer_id`,`newrelic_alerts_policy`.`ext_id`,`newrelic_alerts_policy`.`incident_preference`,`newrelic_alerts_policy`.`name`,`newrelic_alerts_policy`.`created_at`,`newrelic_alerts_policy`.`updated_at` FROM `newrelic_alerts_policy` WHERE `id` = ? LIMIT 1"
	row := db.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CustomerID sql.NullString
	var _ExtID sql.NullInt64
	var _IncidentPreference sql.NullString
	var _Name sql.NullString
	var _CreatedAt sql.NullInt64
	var _UpdatedAt sql.NullInt64
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_ExtID,
		&_IncidentPreference,
		&_Name,
		&_CreatedAt,
		&_UpdatedAt,
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
	if _IncidentPreference.Valid {
		t.SetIncidentPreference(_IncidentPreference.String)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _CreatedAt.Valid {
		t.SetCreatedAt(_CreatedAt.Int64)
	}
	if _UpdatedAt.Valid {
		t.SetUpdatedAt(_UpdatedAt.Int64)
	}
	return true, nil
}

// DBFindOneTx will find a NewrelicAlertsPolicy record in the database with the primary key using the provided transaction
func (t *NewrelicAlertsPolicy) DBFindOneTx(ctx context.Context, tx *sql.Tx, value string) (bool, error) {
	q := "SELECT `newrelic_alerts_policy`.`id`,`newrelic_alerts_policy`.`checksum`,`newrelic_alerts_policy`.`customer_id`,`newrelic_alerts_policy`.`ext_id`,`newrelic_alerts_policy`.`incident_preference`,`newrelic_alerts_policy`.`name`,`newrelic_alerts_policy`.`created_at`,`newrelic_alerts_policy`.`updated_at` FROM `newrelic_alerts_policy` WHERE `id` = ? LIMIT 1"
	row := tx.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CustomerID sql.NullString
	var _ExtID sql.NullInt64
	var _IncidentPreference sql.NullString
	var _Name sql.NullString
	var _CreatedAt sql.NullInt64
	var _UpdatedAt sql.NullInt64
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_ExtID,
		&_IncidentPreference,
		&_Name,
		&_CreatedAt,
		&_UpdatedAt,
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
	if _IncidentPreference.Valid {
		t.SetIncidentPreference(_IncidentPreference.String)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _CreatedAt.Valid {
		t.SetCreatedAt(_CreatedAt.Int64)
	}
	if _UpdatedAt.Valid {
		t.SetUpdatedAt(_UpdatedAt.Int64)
	}
	return true, nil
}

// FindNewrelicAlertsPolicies will find a NewrelicAlertsPolicy record in the database with the provided parameters
func FindNewrelicAlertsPolicies(ctx context.Context, db *sql.DB, _params ...interface{}) ([]*NewrelicAlertsPolicy, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("customer_id"),
		orm.Column("ext_id"),
		orm.Column("incident_preference"),
		orm.Column("name"),
		orm.Column("created_at"),
		orm.Column("updated_at"),
		orm.Table(NewrelicAlertsPolicyTableName),
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
	results := make([]*NewrelicAlertsPolicy, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _ExtID sql.NullInt64
		var _IncidentPreference sql.NullString
		var _Name sql.NullString
		var _CreatedAt sql.NullInt64
		var _UpdatedAt sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_ExtID,
			&_IncidentPreference,
			&_Name,
			&_CreatedAt,
			&_UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		t := &NewrelicAlertsPolicy{}
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
		if _IncidentPreference.Valid {
			t.SetIncidentPreference(_IncidentPreference.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _CreatedAt.Valid {
			t.SetCreatedAt(_CreatedAt.Int64)
		}
		if _UpdatedAt.Valid {
			t.SetUpdatedAt(_UpdatedAt.Int64)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindNewrelicAlertsPoliciesTx will find a NewrelicAlertsPolicy record in the database with the provided parameters using the provided transaction
func FindNewrelicAlertsPoliciesTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) ([]*NewrelicAlertsPolicy, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("customer_id"),
		orm.Column("ext_id"),
		orm.Column("incident_preference"),
		orm.Column("name"),
		orm.Column("created_at"),
		orm.Column("updated_at"),
		orm.Table(NewrelicAlertsPolicyTableName),
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
	results := make([]*NewrelicAlertsPolicy, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _ExtID sql.NullInt64
		var _IncidentPreference sql.NullString
		var _Name sql.NullString
		var _CreatedAt sql.NullInt64
		var _UpdatedAt sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_ExtID,
			&_IncidentPreference,
			&_Name,
			&_CreatedAt,
			&_UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		t := &NewrelicAlertsPolicy{}
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
		if _IncidentPreference.Valid {
			t.SetIncidentPreference(_IncidentPreference.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _CreatedAt.Valid {
			t.SetCreatedAt(_CreatedAt.Int64)
		}
		if _UpdatedAt.Valid {
			t.SetUpdatedAt(_UpdatedAt.Int64)
		}
		results = append(results, t)
	}
	return results, nil
}

// DBFind will find a NewrelicAlertsPolicy record in the database with the provided parameters
func (t *NewrelicAlertsPolicy) DBFind(ctx context.Context, db *sql.DB, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("customer_id"),
		orm.Column("ext_id"),
		orm.Column("incident_preference"),
		orm.Column("name"),
		orm.Column("created_at"),
		orm.Column("updated_at"),
		orm.Table(NewrelicAlertsPolicyTableName),
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
	var _IncidentPreference sql.NullString
	var _Name sql.NullString
	var _CreatedAt sql.NullInt64
	var _UpdatedAt sql.NullInt64
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_ExtID,
		&_IncidentPreference,
		&_Name,
		&_CreatedAt,
		&_UpdatedAt,
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
	if _IncidentPreference.Valid {
		t.SetIncidentPreference(_IncidentPreference.String)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _CreatedAt.Valid {
		t.SetCreatedAt(_CreatedAt.Int64)
	}
	if _UpdatedAt.Valid {
		t.SetUpdatedAt(_UpdatedAt.Int64)
	}
	return true, nil
}

// DBFindTx will find a NewrelicAlertsPolicy record in the database with the provided parameters using the provided transaction
func (t *NewrelicAlertsPolicy) DBFindTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("customer_id"),
		orm.Column("ext_id"),
		orm.Column("incident_preference"),
		orm.Column("name"),
		orm.Column("created_at"),
		orm.Column("updated_at"),
		orm.Table(NewrelicAlertsPolicyTableName),
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
	var _IncidentPreference sql.NullString
	var _Name sql.NullString
	var _CreatedAt sql.NullInt64
	var _UpdatedAt sql.NullInt64
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_ExtID,
		&_IncidentPreference,
		&_Name,
		&_CreatedAt,
		&_UpdatedAt,
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
	if _IncidentPreference.Valid {
		t.SetIncidentPreference(_IncidentPreference.String)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _CreatedAt.Valid {
		t.SetCreatedAt(_CreatedAt.Int64)
	}
	if _UpdatedAt.Valid {
		t.SetUpdatedAt(_UpdatedAt.Int64)
	}
	return true, nil
}

// CountNewrelicAlertsPolicies will find the count of NewrelicAlertsPolicy records in the database
func CountNewrelicAlertsPolicies(ctx context.Context, db *sql.DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(NewrelicAlertsPolicyTableName),
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

// CountNewrelicAlertsPoliciesTx will find the count of NewrelicAlertsPolicy records in the database using the provided transaction
func CountNewrelicAlertsPoliciesTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(NewrelicAlertsPolicyTableName),
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

// DBCount will find the count of NewrelicAlertsPolicy records in the database
func (t *NewrelicAlertsPolicy) DBCount(ctx context.Context, db *sql.DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(NewrelicAlertsPolicyTableName),
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

// DBCountTx will find the count of NewrelicAlertsPolicy records in the database using the provided transaction
func (t *NewrelicAlertsPolicy) DBCountTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(NewrelicAlertsPolicyTableName),
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

// DBExists will return true if the NewrelicAlertsPolicy record exists in the database
func (t *NewrelicAlertsPolicy) DBExists(ctx context.Context, db *sql.DB) (bool, error) {
	q := "SELECT `id` FROM `newrelic_alerts_policy` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := db.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// DBExistsTx will return true if the NewrelicAlertsPolicy record exists in the database using the provided transaction
func (t *NewrelicAlertsPolicy) DBExistsTx(ctx context.Context, tx *sql.Tx) (bool, error) {
	q := "SELECT `id` FROM `newrelic_alerts_policy` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := tx.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// PrimaryKeyColumn returns the column name for the primary key
func (t *NewrelicAlertsPolicy) PrimaryKeyColumn() string {
	return NewrelicAlertsPolicyColumnID
}

// PrimaryKeyColumnType returns the primary key column Go type as a string
func (t *NewrelicAlertsPolicy) PrimaryKeyColumnType() string {
	return "string"
}

// PrimaryKey returns the primary key column value
func (t *NewrelicAlertsPolicy) PrimaryKey() interface{} {
	return t.ID
}
