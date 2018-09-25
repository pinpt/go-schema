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

var _ Model = (*NewrelicAlertsIncident)(nil)
var _ CSVWriter = (*NewrelicAlertsIncident)(nil)
var _ JSONWriter = (*NewrelicAlertsIncident)(nil)
var _ Checksum = (*NewrelicAlertsIncident)(nil)

// NewrelicAlertsIncidentTableName is the name of the table in SQL
const NewrelicAlertsIncidentTableName = "newrelic_alerts_incident"

var NewrelicAlertsIncidentColumns = []string{
	"id",
	"checksum",
	"customer_id",
	"ext_id",
	"opened_at",
	"closed_at",
	"incident_preference",
	"violation_ext_ids",
	"violation_ids",
	"policy_ext_id",
	"policy_id",
}

// NewrelicAlertsIncident table
type NewrelicAlertsIncident struct {
	Checksum           *string `json:"checksum,omitempty"`
	ClosedAt           *int64  `json:"closed_at,omitempty"`
	CustomerID         string  `json:"customer_id"`
	ExtID              int64   `json:"ext_id"`
	ID                 string  `json:"id"`
	IncidentPreference string  `json:"incident_preference"`
	OpenedAt           int64   `json:"opened_at"`
	PolicyExtID        *int64  `json:"policy_ext_id,omitempty"`
	PolicyID           *string `json:"policy_id,omitempty"`
	ViolationExtIds    string  `json:"violation_ext_ids"`
	ViolationIds       string  `json:"violation_ids"`
}

// TableName returns the SQL table name for NewrelicAlertsIncident and satifies the Model interface
func (t *NewrelicAlertsIncident) TableName() string {
	return NewrelicAlertsIncidentTableName
}

// ToCSV will serialize the NewrelicAlertsIncident instance to a CSV compatible array of strings
func (t *NewrelicAlertsIncident) ToCSV() []string {
	return []string{
		t.ID,
		t.CalculateChecksum(),
		t.CustomerID,
		toCSVString(t.ExtID),
		toCSVString(t.OpenedAt),
		toCSVString(t.ClosedAt),
		t.IncidentPreference,
		t.ViolationExtIds,
		t.ViolationIds,
		toCSVString(t.PolicyExtID),
		toCSVString(t.PolicyID),
	}
}

// WriteCSV will serialize the NewrelicAlertsIncident instance to the writer as CSV and satisfies the CSVWriter interface
func (t *NewrelicAlertsIncident) WriteCSV(w *csv.Writer) error {
	return w.Write(t.ToCSV())
}

// WriteJSON will serialize the NewrelicAlertsIncident instance to the writer as JSON and satisfies the JSONWriter interface
func (t *NewrelicAlertsIncident) WriteJSON(w io.Writer, indent ...bool) error {
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

// NewNewrelicAlertsIncidentReader creates a JSON reader which can read in NewrelicAlertsIncident objects serialized as JSON either as an array, single object or json new lines
// and writes each NewrelicAlertsIncident to the channel provided
func NewNewrelicAlertsIncidentReader(r io.Reader, ch chan<- NewrelicAlertsIncident) error {
	return orm.Deserialize(r, func(buf json.RawMessage) error {
		dec := json.NewDecoder(bytes.NewBuffer(buf))
		e := NewrelicAlertsIncident{}
		if err := dec.Decode(&e); err != nil {
			return err
		}
		ch <- e
		return nil
	})
}

// NewCSVNewrelicAlertsIncidentReaderDir will read the reader as CSV and emit each record to the channel provided
func NewCSVNewrelicAlertsIncidentReader(r io.Reader, ch chan<- NewrelicAlertsIncident) error {
	cr := csv.NewReader(r)
	for {
		record, err := cr.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		ch <- NewrelicAlertsIncident{
			ID:                 record[0],
			Checksum:           fromStringPointer(record[1]),
			CustomerID:         record[2],
			ExtID:              fromCSVInt64(record[3]),
			OpenedAt:           fromCSVInt64(record[4]),
			ClosedAt:           fromCSVInt64Pointer(record[5]),
			IncidentPreference: record[6],
			ViolationExtIds:    record[7],
			ViolationIds:       record[8],
			PolicyExtID:        fromCSVInt64Pointer(record[9]),
			PolicyID:           fromStringPointer(record[10]),
		}
	}
	return nil
}

// NewCSVNewrelicAlertsIncidentReaderFile will read the file as a CSV and emit each record to the channel provided
func NewCSVNewrelicAlertsIncidentReaderFile(fp string, ch chan<- NewrelicAlertsIncident) error {
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
	return NewCSVNewrelicAlertsIncidentReader(fc, ch)
}

// NewCSVNewrelicAlertsIncidentReaderDir will read the newrelic_alerts_incident.csv.gz file as a CSV and emit each record to the channel provided
func NewCSVNewrelicAlertsIncidentReaderDir(dir string, ch chan<- NewrelicAlertsIncident) error {
	return NewCSVNewrelicAlertsIncidentReaderFile(filepath.Join(dir, "newrelic_alerts_incident.csv.gz"), ch)
}

// NewrelicAlertsIncidentCSVDeduper is a function callback which takes the existing value (a) and the new value (b)
// and the return value should be the one to use (or a new one, if applicable). return nil
// to skip processing of this record
type NewrelicAlertsIncidentCSVDeduper func(a NewrelicAlertsIncident, b NewrelicAlertsIncident) *NewrelicAlertsIncident

// NewrelicAlertsIncidentCSVDedupeDisabled is set on whether the CSV writer should de-dupe values by key
var NewrelicAlertsIncidentCSVDedupeDisabled bool

// NewNewrelicAlertsIncidentCSVWriterSize creates a batch writer that will write each NewrelicAlertsIncident into a CSV file
// this method will automatically de-duplicate entries using the primary key. if the checksum
// for a newer item with the same primary key doesn't match a previously sent item, the newer
// one will be used
func NewNewrelicAlertsIncidentCSVWriterSize(w io.Writer, size int, dedupers ...NewrelicAlertsIncidentCSVDeduper) (chan NewrelicAlertsIncident, chan bool, error) {
	cw := csv.NewWriter(w)
	ch := make(chan NewrelicAlertsIncident, size)
	done := make(chan bool)
	go func() {
		defer func() { done <- true }()
		dodedupe := !NewrelicAlertsIncidentCSVDedupeDisabled
		var kv map[string]*NewrelicAlertsIncident
		var deduper NewrelicAlertsIncidentCSVDeduper
		if dedupers != nil && len(dedupers) > 0 {
			deduper = dedupers[0]
			dodedupe = true
		}
		if dodedupe {
			kv = make(map[string]*NewrelicAlertsIncident)
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

// NewrelicAlertsIncidentCSVDefaultSize is the default channel buffer size if not provided
var NewrelicAlertsIncidentCSVDefaultSize = 100

// NewNewrelicAlertsIncidentCSVWriter creates a batch writer that will write each NewrelicAlertsIncident into a CSV file
func NewNewrelicAlertsIncidentCSVWriter(w io.Writer, dedupers ...NewrelicAlertsIncidentCSVDeduper) (chan NewrelicAlertsIncident, chan bool, error) {
	return NewNewrelicAlertsIncidentCSVWriterSize(w, NewrelicAlertsIncidentCSVDefaultSize, dedupers...)
}

// NewNewrelicAlertsIncidentCSVWriterDir creates a batch writer that will write each NewrelicAlertsIncident into a CSV file named newrelic_alerts_incident.csv.gz in dir
func NewNewrelicAlertsIncidentCSVWriterDir(dir string, dedupers ...NewrelicAlertsIncidentCSVDeduper) (chan NewrelicAlertsIncident, chan bool, error) {
	return NewNewrelicAlertsIncidentCSVWriterFile(filepath.Join(dir, "newrelic_alerts_incident.csv.gz"), dedupers...)
}

// NewNewrelicAlertsIncidentCSVWriterFile creates a batch writer that will write each NewrelicAlertsIncident into a CSV file
func NewNewrelicAlertsIncidentCSVWriterFile(fn string, dedupers ...NewrelicAlertsIncidentCSVDeduper) (chan NewrelicAlertsIncident, chan bool, error) {
	f, err := os.Create(fn)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening CSV file %s. %v", fn, err)
	}
	var fc io.WriteCloser = f
	if filepath.Ext(fn) == ".gz" {
		w, _ := gzip.NewWriterLevel(f, gzip.BestCompression)
		fc = w
	}
	ch, done, err := NewNewrelicAlertsIncidentCSVWriter(fc, dedupers...)
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

type NewrelicAlertsIncidentDBAction func(ctx context.Context, db *sql.DB, record NewrelicAlertsIncident) error

// NewNewrelicAlertsIncidentDBWriterSize creates a DB writer that will write each issue into the DB
func NewNewrelicAlertsIncidentDBWriterSize(ctx context.Context, db *sql.DB, errors chan<- error, size int, actions ...NewrelicAlertsIncidentDBAction) (chan NewrelicAlertsIncident, chan bool, error) {
	ch := make(chan NewrelicAlertsIncident, size)
	done := make(chan bool)
	var action NewrelicAlertsIncidentDBAction
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

// NewNewrelicAlertsIncidentDBWriter creates a DB writer that will write each issue into the DB
func NewNewrelicAlertsIncidentDBWriter(ctx context.Context, db *sql.DB, errors chan<- error, actions ...NewrelicAlertsIncidentDBAction) (chan NewrelicAlertsIncident, chan bool, error) {
	return NewNewrelicAlertsIncidentDBWriterSize(ctx, db, errors, 100, actions...)
}

// NewrelicAlertsIncidentColumnID is the ID SQL column name for the NewrelicAlertsIncident table
const NewrelicAlertsIncidentColumnID = "id"

// NewrelicAlertsIncidentEscapedColumnID is the escaped ID SQL column name for the NewrelicAlertsIncident table
const NewrelicAlertsIncidentEscapedColumnID = "`id`"

// NewrelicAlertsIncidentColumnChecksum is the Checksum SQL column name for the NewrelicAlertsIncident table
const NewrelicAlertsIncidentColumnChecksum = "checksum"

// NewrelicAlertsIncidentEscapedColumnChecksum is the escaped Checksum SQL column name for the NewrelicAlertsIncident table
const NewrelicAlertsIncidentEscapedColumnChecksum = "`checksum`"

// NewrelicAlertsIncidentColumnCustomerID is the CustomerID SQL column name for the NewrelicAlertsIncident table
const NewrelicAlertsIncidentColumnCustomerID = "customer_id"

// NewrelicAlertsIncidentEscapedColumnCustomerID is the escaped CustomerID SQL column name for the NewrelicAlertsIncident table
const NewrelicAlertsIncidentEscapedColumnCustomerID = "`customer_id`"

// NewrelicAlertsIncidentColumnExtID is the ExtID SQL column name for the NewrelicAlertsIncident table
const NewrelicAlertsIncidentColumnExtID = "ext_id"

// NewrelicAlertsIncidentEscapedColumnExtID is the escaped ExtID SQL column name for the NewrelicAlertsIncident table
const NewrelicAlertsIncidentEscapedColumnExtID = "`ext_id`"

// NewrelicAlertsIncidentColumnOpenedAt is the OpenedAt SQL column name for the NewrelicAlertsIncident table
const NewrelicAlertsIncidentColumnOpenedAt = "opened_at"

// NewrelicAlertsIncidentEscapedColumnOpenedAt is the escaped OpenedAt SQL column name for the NewrelicAlertsIncident table
const NewrelicAlertsIncidentEscapedColumnOpenedAt = "`opened_at`"

// NewrelicAlertsIncidentColumnClosedAt is the ClosedAt SQL column name for the NewrelicAlertsIncident table
const NewrelicAlertsIncidentColumnClosedAt = "closed_at"

// NewrelicAlertsIncidentEscapedColumnClosedAt is the escaped ClosedAt SQL column name for the NewrelicAlertsIncident table
const NewrelicAlertsIncidentEscapedColumnClosedAt = "`closed_at`"

// NewrelicAlertsIncidentColumnIncidentPreference is the IncidentPreference SQL column name for the NewrelicAlertsIncident table
const NewrelicAlertsIncidentColumnIncidentPreference = "incident_preference"

// NewrelicAlertsIncidentEscapedColumnIncidentPreference is the escaped IncidentPreference SQL column name for the NewrelicAlertsIncident table
const NewrelicAlertsIncidentEscapedColumnIncidentPreference = "`incident_preference`"

// NewrelicAlertsIncidentColumnViolationExtIds is the ViolationExtIds SQL column name for the NewrelicAlertsIncident table
const NewrelicAlertsIncidentColumnViolationExtIds = "violation_ext_ids"

// NewrelicAlertsIncidentEscapedColumnViolationExtIds is the escaped ViolationExtIds SQL column name for the NewrelicAlertsIncident table
const NewrelicAlertsIncidentEscapedColumnViolationExtIds = "`violation_ext_ids`"

// NewrelicAlertsIncidentColumnViolationIds is the ViolationIds SQL column name for the NewrelicAlertsIncident table
const NewrelicAlertsIncidentColumnViolationIds = "violation_ids"

// NewrelicAlertsIncidentEscapedColumnViolationIds is the escaped ViolationIds SQL column name for the NewrelicAlertsIncident table
const NewrelicAlertsIncidentEscapedColumnViolationIds = "`violation_ids`"

// NewrelicAlertsIncidentColumnPolicyExtID is the PolicyExtID SQL column name for the NewrelicAlertsIncident table
const NewrelicAlertsIncidentColumnPolicyExtID = "policy_ext_id"

// NewrelicAlertsIncidentEscapedColumnPolicyExtID is the escaped PolicyExtID SQL column name for the NewrelicAlertsIncident table
const NewrelicAlertsIncidentEscapedColumnPolicyExtID = "`policy_ext_id`"

// NewrelicAlertsIncidentColumnPolicyID is the PolicyID SQL column name for the NewrelicAlertsIncident table
const NewrelicAlertsIncidentColumnPolicyID = "policy_id"

// NewrelicAlertsIncidentEscapedColumnPolicyID is the escaped PolicyID SQL column name for the NewrelicAlertsIncident table
const NewrelicAlertsIncidentEscapedColumnPolicyID = "`policy_id`"

// GetID will return the NewrelicAlertsIncident ID value
func (t *NewrelicAlertsIncident) GetID() string {
	return t.ID
}

// SetID will set the NewrelicAlertsIncident ID value
func (t *NewrelicAlertsIncident) SetID(v string) {
	t.ID = v
}

// FindNewrelicAlertsIncidentByID will find a NewrelicAlertsIncident by ID
func FindNewrelicAlertsIncidentByID(ctx context.Context, db *sql.DB, value string) (*NewrelicAlertsIncident, error) {
	q := "SELECT `newrelic_alerts_incident`.`id`,`newrelic_alerts_incident`.`checksum`,`newrelic_alerts_incident`.`customer_id`,`newrelic_alerts_incident`.`ext_id`,`newrelic_alerts_incident`.`opened_at`,`newrelic_alerts_incident`.`closed_at`,`newrelic_alerts_incident`.`incident_preference`,`newrelic_alerts_incident`.`violation_ext_ids`,`newrelic_alerts_incident`.`violation_ids`,`newrelic_alerts_incident`.`policy_ext_id`,`newrelic_alerts_incident`.`policy_id` FROM `newrelic_alerts_incident` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CustomerID sql.NullString
	var _ExtID sql.NullInt64
	var _OpenedAt sql.NullInt64
	var _ClosedAt sql.NullInt64
	var _IncidentPreference sql.NullString
	var _ViolationExtIds sql.NullString
	var _ViolationIds sql.NullString
	var _PolicyExtID sql.NullInt64
	var _PolicyID sql.NullString
	err := db.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_ExtID,
		&_OpenedAt,
		&_ClosedAt,
		&_IncidentPreference,
		&_ViolationExtIds,
		&_ViolationIds,
		&_PolicyExtID,
		&_PolicyID,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &NewrelicAlertsIncident{}
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
	if _OpenedAt.Valid {
		t.SetOpenedAt(_OpenedAt.Int64)
	}
	if _ClosedAt.Valid {
		t.SetClosedAt(_ClosedAt.Int64)
	}
	if _IncidentPreference.Valid {
		t.SetIncidentPreference(_IncidentPreference.String)
	}
	if _ViolationExtIds.Valid {
		t.SetViolationExtIds(_ViolationExtIds.String)
	}
	if _ViolationIds.Valid {
		t.SetViolationIds(_ViolationIds.String)
	}
	if _PolicyExtID.Valid {
		t.SetPolicyExtID(_PolicyExtID.Int64)
	}
	if _PolicyID.Valid {
		t.SetPolicyID(_PolicyID.String)
	}
	return t, nil
}

// FindNewrelicAlertsIncidentByIDTx will find a NewrelicAlertsIncident by ID using the provided transaction
func FindNewrelicAlertsIncidentByIDTx(ctx context.Context, tx *sql.Tx, value string) (*NewrelicAlertsIncident, error) {
	q := "SELECT `newrelic_alerts_incident`.`id`,`newrelic_alerts_incident`.`checksum`,`newrelic_alerts_incident`.`customer_id`,`newrelic_alerts_incident`.`ext_id`,`newrelic_alerts_incident`.`opened_at`,`newrelic_alerts_incident`.`closed_at`,`newrelic_alerts_incident`.`incident_preference`,`newrelic_alerts_incident`.`violation_ext_ids`,`newrelic_alerts_incident`.`violation_ids`,`newrelic_alerts_incident`.`policy_ext_id`,`newrelic_alerts_incident`.`policy_id` FROM `newrelic_alerts_incident` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CustomerID sql.NullString
	var _ExtID sql.NullInt64
	var _OpenedAt sql.NullInt64
	var _ClosedAt sql.NullInt64
	var _IncidentPreference sql.NullString
	var _ViolationExtIds sql.NullString
	var _ViolationIds sql.NullString
	var _PolicyExtID sql.NullInt64
	var _PolicyID sql.NullString
	err := tx.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_ExtID,
		&_OpenedAt,
		&_ClosedAt,
		&_IncidentPreference,
		&_ViolationExtIds,
		&_ViolationIds,
		&_PolicyExtID,
		&_PolicyID,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &NewrelicAlertsIncident{}
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
	if _OpenedAt.Valid {
		t.SetOpenedAt(_OpenedAt.Int64)
	}
	if _ClosedAt.Valid {
		t.SetClosedAt(_ClosedAt.Int64)
	}
	if _IncidentPreference.Valid {
		t.SetIncidentPreference(_IncidentPreference.String)
	}
	if _ViolationExtIds.Valid {
		t.SetViolationExtIds(_ViolationExtIds.String)
	}
	if _ViolationIds.Valid {
		t.SetViolationIds(_ViolationIds.String)
	}
	if _PolicyExtID.Valid {
		t.SetPolicyExtID(_PolicyExtID.Int64)
	}
	if _PolicyID.Valid {
		t.SetPolicyID(_PolicyID.String)
	}
	return t, nil
}

// GetChecksum will return the NewrelicAlertsIncident Checksum value
func (t *NewrelicAlertsIncident) GetChecksum() string {
	if t.Checksum == nil {
		return ""
	}
	return *t.Checksum
}

// SetChecksum will set the NewrelicAlertsIncident Checksum value
func (t *NewrelicAlertsIncident) SetChecksum(v string) {
	t.Checksum = &v
}

// GetCustomerID will return the NewrelicAlertsIncident CustomerID value
func (t *NewrelicAlertsIncident) GetCustomerID() string {
	return t.CustomerID
}

// SetCustomerID will set the NewrelicAlertsIncident CustomerID value
func (t *NewrelicAlertsIncident) SetCustomerID(v string) {
	t.CustomerID = v
}

// FindNewrelicAlertsIncidentsByCustomerID will find all NewrelicAlertsIncidents by the CustomerID value
func FindNewrelicAlertsIncidentsByCustomerID(ctx context.Context, db *sql.DB, value string) ([]*NewrelicAlertsIncident, error) {
	q := "SELECT `newrelic_alerts_incident`.`id`,`newrelic_alerts_incident`.`checksum`,`newrelic_alerts_incident`.`customer_id`,`newrelic_alerts_incident`.`ext_id`,`newrelic_alerts_incident`.`opened_at`,`newrelic_alerts_incident`.`closed_at`,`newrelic_alerts_incident`.`incident_preference`,`newrelic_alerts_incident`.`violation_ext_ids`,`newrelic_alerts_incident`.`violation_ids`,`newrelic_alerts_incident`.`policy_ext_id`,`newrelic_alerts_incident`.`policy_id` FROM `newrelic_alerts_incident` WHERE `customer_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*NewrelicAlertsIncident, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _ExtID sql.NullInt64
		var _OpenedAt sql.NullInt64
		var _ClosedAt sql.NullInt64
		var _IncidentPreference sql.NullString
		var _ViolationExtIds sql.NullString
		var _ViolationIds sql.NullString
		var _PolicyExtID sql.NullInt64
		var _PolicyID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_ExtID,
			&_OpenedAt,
			&_ClosedAt,
			&_IncidentPreference,
			&_ViolationExtIds,
			&_ViolationIds,
			&_PolicyExtID,
			&_PolicyID,
		)
		if err != nil {
			return nil, err
		}
		t := &NewrelicAlertsIncident{}
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
		if _OpenedAt.Valid {
			t.SetOpenedAt(_OpenedAt.Int64)
		}
		if _ClosedAt.Valid {
			t.SetClosedAt(_ClosedAt.Int64)
		}
		if _IncidentPreference.Valid {
			t.SetIncidentPreference(_IncidentPreference.String)
		}
		if _ViolationExtIds.Valid {
			t.SetViolationExtIds(_ViolationExtIds.String)
		}
		if _ViolationIds.Valid {
			t.SetViolationIds(_ViolationIds.String)
		}
		if _PolicyExtID.Valid {
			t.SetPolicyExtID(_PolicyExtID.Int64)
		}
		if _PolicyID.Valid {
			t.SetPolicyID(_PolicyID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindNewrelicAlertsIncidentsByCustomerIDTx will find all NewrelicAlertsIncidents by the CustomerID value using the provided transaction
func FindNewrelicAlertsIncidentsByCustomerIDTx(ctx context.Context, tx *sql.Tx, value string) ([]*NewrelicAlertsIncident, error) {
	q := "SELECT `newrelic_alerts_incident`.`id`,`newrelic_alerts_incident`.`checksum`,`newrelic_alerts_incident`.`customer_id`,`newrelic_alerts_incident`.`ext_id`,`newrelic_alerts_incident`.`opened_at`,`newrelic_alerts_incident`.`closed_at`,`newrelic_alerts_incident`.`incident_preference`,`newrelic_alerts_incident`.`violation_ext_ids`,`newrelic_alerts_incident`.`violation_ids`,`newrelic_alerts_incident`.`policy_ext_id`,`newrelic_alerts_incident`.`policy_id` FROM `newrelic_alerts_incident` WHERE `customer_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*NewrelicAlertsIncident, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _ExtID sql.NullInt64
		var _OpenedAt sql.NullInt64
		var _ClosedAt sql.NullInt64
		var _IncidentPreference sql.NullString
		var _ViolationExtIds sql.NullString
		var _ViolationIds sql.NullString
		var _PolicyExtID sql.NullInt64
		var _PolicyID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_ExtID,
			&_OpenedAt,
			&_ClosedAt,
			&_IncidentPreference,
			&_ViolationExtIds,
			&_ViolationIds,
			&_PolicyExtID,
			&_PolicyID,
		)
		if err != nil {
			return nil, err
		}
		t := &NewrelicAlertsIncident{}
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
		if _OpenedAt.Valid {
			t.SetOpenedAt(_OpenedAt.Int64)
		}
		if _ClosedAt.Valid {
			t.SetClosedAt(_ClosedAt.Int64)
		}
		if _IncidentPreference.Valid {
			t.SetIncidentPreference(_IncidentPreference.String)
		}
		if _ViolationExtIds.Valid {
			t.SetViolationExtIds(_ViolationExtIds.String)
		}
		if _ViolationIds.Valid {
			t.SetViolationIds(_ViolationIds.String)
		}
		if _PolicyExtID.Valid {
			t.SetPolicyExtID(_PolicyExtID.Int64)
		}
		if _PolicyID.Valid {
			t.SetPolicyID(_PolicyID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetExtID will return the NewrelicAlertsIncident ExtID value
func (t *NewrelicAlertsIncident) GetExtID() int64 {
	return t.ExtID
}

// SetExtID will set the NewrelicAlertsIncident ExtID value
func (t *NewrelicAlertsIncident) SetExtID(v int64) {
	t.ExtID = v
}

// GetOpenedAt will return the NewrelicAlertsIncident OpenedAt value
func (t *NewrelicAlertsIncident) GetOpenedAt() int64 {
	return t.OpenedAt
}

// SetOpenedAt will set the NewrelicAlertsIncident OpenedAt value
func (t *NewrelicAlertsIncident) SetOpenedAt(v int64) {
	t.OpenedAt = v
}

// GetClosedAt will return the NewrelicAlertsIncident ClosedAt value
func (t *NewrelicAlertsIncident) GetClosedAt() int64 {
	if t.ClosedAt == nil {
		return int64(0)
	}
	return *t.ClosedAt
}

// SetClosedAt will set the NewrelicAlertsIncident ClosedAt value
func (t *NewrelicAlertsIncident) SetClosedAt(v int64) {
	t.ClosedAt = &v
}

// GetIncidentPreference will return the NewrelicAlertsIncident IncidentPreference value
func (t *NewrelicAlertsIncident) GetIncidentPreference() string {
	return t.IncidentPreference
}

// SetIncidentPreference will set the NewrelicAlertsIncident IncidentPreference value
func (t *NewrelicAlertsIncident) SetIncidentPreference(v string) {
	t.IncidentPreference = v
}

// GetViolationExtIds will return the NewrelicAlertsIncident ViolationExtIds value
func (t *NewrelicAlertsIncident) GetViolationExtIds() string {
	return t.ViolationExtIds
}

// SetViolationExtIds will set the NewrelicAlertsIncident ViolationExtIds value
func (t *NewrelicAlertsIncident) SetViolationExtIds(v string) {
	t.ViolationExtIds = v
}

// GetViolationIds will return the NewrelicAlertsIncident ViolationIds value
func (t *NewrelicAlertsIncident) GetViolationIds() string {
	return t.ViolationIds
}

// SetViolationIds will set the NewrelicAlertsIncident ViolationIds value
func (t *NewrelicAlertsIncident) SetViolationIds(v string) {
	t.ViolationIds = v
}

// GetPolicyExtID will return the NewrelicAlertsIncident PolicyExtID value
func (t *NewrelicAlertsIncident) GetPolicyExtID() int64 {
	if t.PolicyExtID == nil {
		return int64(0)
	}
	return *t.PolicyExtID
}

// SetPolicyExtID will set the NewrelicAlertsIncident PolicyExtID value
func (t *NewrelicAlertsIncident) SetPolicyExtID(v int64) {
	t.PolicyExtID = &v
}

// FindNewrelicAlertsIncidentsByPolicyExtID will find all NewrelicAlertsIncidents by the PolicyExtID value
func FindNewrelicAlertsIncidentsByPolicyExtID(ctx context.Context, db *sql.DB, value int64) ([]*NewrelicAlertsIncident, error) {
	q := "SELECT `newrelic_alerts_incident`.`id`,`newrelic_alerts_incident`.`checksum`,`newrelic_alerts_incident`.`customer_id`,`newrelic_alerts_incident`.`ext_id`,`newrelic_alerts_incident`.`opened_at`,`newrelic_alerts_incident`.`closed_at`,`newrelic_alerts_incident`.`incident_preference`,`newrelic_alerts_incident`.`violation_ext_ids`,`newrelic_alerts_incident`.`violation_ids`,`newrelic_alerts_incident`.`policy_ext_id`,`newrelic_alerts_incident`.`policy_id` FROM `newrelic_alerts_incident` WHERE `policy_ext_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*NewrelicAlertsIncident, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _ExtID sql.NullInt64
		var _OpenedAt sql.NullInt64
		var _ClosedAt sql.NullInt64
		var _IncidentPreference sql.NullString
		var _ViolationExtIds sql.NullString
		var _ViolationIds sql.NullString
		var _PolicyExtID sql.NullInt64
		var _PolicyID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_ExtID,
			&_OpenedAt,
			&_ClosedAt,
			&_IncidentPreference,
			&_ViolationExtIds,
			&_ViolationIds,
			&_PolicyExtID,
			&_PolicyID,
		)
		if err != nil {
			return nil, err
		}
		t := &NewrelicAlertsIncident{}
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
		if _OpenedAt.Valid {
			t.SetOpenedAt(_OpenedAt.Int64)
		}
		if _ClosedAt.Valid {
			t.SetClosedAt(_ClosedAt.Int64)
		}
		if _IncidentPreference.Valid {
			t.SetIncidentPreference(_IncidentPreference.String)
		}
		if _ViolationExtIds.Valid {
			t.SetViolationExtIds(_ViolationExtIds.String)
		}
		if _ViolationIds.Valid {
			t.SetViolationIds(_ViolationIds.String)
		}
		if _PolicyExtID.Valid {
			t.SetPolicyExtID(_PolicyExtID.Int64)
		}
		if _PolicyID.Valid {
			t.SetPolicyID(_PolicyID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindNewrelicAlertsIncidentsByPolicyExtIDTx will find all NewrelicAlertsIncidents by the PolicyExtID value using the provided transaction
func FindNewrelicAlertsIncidentsByPolicyExtIDTx(ctx context.Context, tx *sql.Tx, value int64) ([]*NewrelicAlertsIncident, error) {
	q := "SELECT `newrelic_alerts_incident`.`id`,`newrelic_alerts_incident`.`checksum`,`newrelic_alerts_incident`.`customer_id`,`newrelic_alerts_incident`.`ext_id`,`newrelic_alerts_incident`.`opened_at`,`newrelic_alerts_incident`.`closed_at`,`newrelic_alerts_incident`.`incident_preference`,`newrelic_alerts_incident`.`violation_ext_ids`,`newrelic_alerts_incident`.`violation_ids`,`newrelic_alerts_incident`.`policy_ext_id`,`newrelic_alerts_incident`.`policy_id` FROM `newrelic_alerts_incident` WHERE `policy_ext_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*NewrelicAlertsIncident, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _ExtID sql.NullInt64
		var _OpenedAt sql.NullInt64
		var _ClosedAt sql.NullInt64
		var _IncidentPreference sql.NullString
		var _ViolationExtIds sql.NullString
		var _ViolationIds sql.NullString
		var _PolicyExtID sql.NullInt64
		var _PolicyID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_ExtID,
			&_OpenedAt,
			&_ClosedAt,
			&_IncidentPreference,
			&_ViolationExtIds,
			&_ViolationIds,
			&_PolicyExtID,
			&_PolicyID,
		)
		if err != nil {
			return nil, err
		}
		t := &NewrelicAlertsIncident{}
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
		if _OpenedAt.Valid {
			t.SetOpenedAt(_OpenedAt.Int64)
		}
		if _ClosedAt.Valid {
			t.SetClosedAt(_ClosedAt.Int64)
		}
		if _IncidentPreference.Valid {
			t.SetIncidentPreference(_IncidentPreference.String)
		}
		if _ViolationExtIds.Valid {
			t.SetViolationExtIds(_ViolationExtIds.String)
		}
		if _ViolationIds.Valid {
			t.SetViolationIds(_ViolationIds.String)
		}
		if _PolicyExtID.Valid {
			t.SetPolicyExtID(_PolicyExtID.Int64)
		}
		if _PolicyID.Valid {
			t.SetPolicyID(_PolicyID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetPolicyID will return the NewrelicAlertsIncident PolicyID value
func (t *NewrelicAlertsIncident) GetPolicyID() string {
	if t.PolicyID == nil {
		return ""
	}
	return *t.PolicyID
}

// SetPolicyID will set the NewrelicAlertsIncident PolicyID value
func (t *NewrelicAlertsIncident) SetPolicyID(v string) {
	t.PolicyID = &v
}

// FindNewrelicAlertsIncidentsByPolicyID will find all NewrelicAlertsIncidents by the PolicyID value
func FindNewrelicAlertsIncidentsByPolicyID(ctx context.Context, db *sql.DB, value string) ([]*NewrelicAlertsIncident, error) {
	q := "SELECT `newrelic_alerts_incident`.`id`,`newrelic_alerts_incident`.`checksum`,`newrelic_alerts_incident`.`customer_id`,`newrelic_alerts_incident`.`ext_id`,`newrelic_alerts_incident`.`opened_at`,`newrelic_alerts_incident`.`closed_at`,`newrelic_alerts_incident`.`incident_preference`,`newrelic_alerts_incident`.`violation_ext_ids`,`newrelic_alerts_incident`.`violation_ids`,`newrelic_alerts_incident`.`policy_ext_id`,`newrelic_alerts_incident`.`policy_id` FROM `newrelic_alerts_incident` WHERE `policy_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*NewrelicAlertsIncident, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _ExtID sql.NullInt64
		var _OpenedAt sql.NullInt64
		var _ClosedAt sql.NullInt64
		var _IncidentPreference sql.NullString
		var _ViolationExtIds sql.NullString
		var _ViolationIds sql.NullString
		var _PolicyExtID sql.NullInt64
		var _PolicyID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_ExtID,
			&_OpenedAt,
			&_ClosedAt,
			&_IncidentPreference,
			&_ViolationExtIds,
			&_ViolationIds,
			&_PolicyExtID,
			&_PolicyID,
		)
		if err != nil {
			return nil, err
		}
		t := &NewrelicAlertsIncident{}
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
		if _OpenedAt.Valid {
			t.SetOpenedAt(_OpenedAt.Int64)
		}
		if _ClosedAt.Valid {
			t.SetClosedAt(_ClosedAt.Int64)
		}
		if _IncidentPreference.Valid {
			t.SetIncidentPreference(_IncidentPreference.String)
		}
		if _ViolationExtIds.Valid {
			t.SetViolationExtIds(_ViolationExtIds.String)
		}
		if _ViolationIds.Valid {
			t.SetViolationIds(_ViolationIds.String)
		}
		if _PolicyExtID.Valid {
			t.SetPolicyExtID(_PolicyExtID.Int64)
		}
		if _PolicyID.Valid {
			t.SetPolicyID(_PolicyID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindNewrelicAlertsIncidentsByPolicyIDTx will find all NewrelicAlertsIncidents by the PolicyID value using the provided transaction
func FindNewrelicAlertsIncidentsByPolicyIDTx(ctx context.Context, tx *sql.Tx, value string) ([]*NewrelicAlertsIncident, error) {
	q := "SELECT `newrelic_alerts_incident`.`id`,`newrelic_alerts_incident`.`checksum`,`newrelic_alerts_incident`.`customer_id`,`newrelic_alerts_incident`.`ext_id`,`newrelic_alerts_incident`.`opened_at`,`newrelic_alerts_incident`.`closed_at`,`newrelic_alerts_incident`.`incident_preference`,`newrelic_alerts_incident`.`violation_ext_ids`,`newrelic_alerts_incident`.`violation_ids`,`newrelic_alerts_incident`.`policy_ext_id`,`newrelic_alerts_incident`.`policy_id` FROM `newrelic_alerts_incident` WHERE `policy_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*NewrelicAlertsIncident, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _ExtID sql.NullInt64
		var _OpenedAt sql.NullInt64
		var _ClosedAt sql.NullInt64
		var _IncidentPreference sql.NullString
		var _ViolationExtIds sql.NullString
		var _ViolationIds sql.NullString
		var _PolicyExtID sql.NullInt64
		var _PolicyID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_ExtID,
			&_OpenedAt,
			&_ClosedAt,
			&_IncidentPreference,
			&_ViolationExtIds,
			&_ViolationIds,
			&_PolicyExtID,
			&_PolicyID,
		)
		if err != nil {
			return nil, err
		}
		t := &NewrelicAlertsIncident{}
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
		if _OpenedAt.Valid {
			t.SetOpenedAt(_OpenedAt.Int64)
		}
		if _ClosedAt.Valid {
			t.SetClosedAt(_ClosedAt.Int64)
		}
		if _IncidentPreference.Valid {
			t.SetIncidentPreference(_IncidentPreference.String)
		}
		if _ViolationExtIds.Valid {
			t.SetViolationExtIds(_ViolationExtIds.String)
		}
		if _ViolationIds.Valid {
			t.SetViolationIds(_ViolationIds.String)
		}
		if _PolicyExtID.Valid {
			t.SetPolicyExtID(_PolicyExtID.Int64)
		}
		if _PolicyID.Valid {
			t.SetPolicyID(_PolicyID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

func (t *NewrelicAlertsIncident) toTimestamp(value time.Time) *timestamp.Timestamp {
	ts, _ := ptypes.TimestampProto(value)
	return ts
}

// DBCreateNewrelicAlertsIncidentTable will create the NewrelicAlertsIncident table
func DBCreateNewrelicAlertsIncidentTable(ctx context.Context, db *sql.DB) error {
	q := "CREATE TABLE `newrelic_alerts_incident` (`id` VARCHAR(64) NOT NULL PRIMARY KEY,`checksum` CHAR(64),`customer_id` VARCHAR(64) NOT NULL,`ext_id`BIGINT(20) NOT NULL,`opened_at`BIGINT(20) NOT NULL,`closed_at`BIGINT(20),`incident_preference` TEXT NOT NULL,`violation_ext_ids` JSON NOT NULL,`violation_ids` JSON NOT NULL,`policy_ext_id` BIGINT(20),`policy_id`VARCHAR(64),INDEX newrelic_alerts_incident_customer_id_index (`customer_id`),INDEX newrelic_alerts_incident_policy_ext_id_index (`policy_ext_id`),INDEX newrelic_alerts_incident_policy_id_index (`policy_id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBCreateNewrelicAlertsIncidentTableTx will create the NewrelicAlertsIncident table using the provided transction
func DBCreateNewrelicAlertsIncidentTableTx(ctx context.Context, tx *sql.Tx) error {
	q := "CREATE TABLE `newrelic_alerts_incident` (`id` VARCHAR(64) NOT NULL PRIMARY KEY,`checksum` CHAR(64),`customer_id` VARCHAR(64) NOT NULL,`ext_id`BIGINT(20) NOT NULL,`opened_at`BIGINT(20) NOT NULL,`closed_at`BIGINT(20),`incident_preference` TEXT NOT NULL,`violation_ext_ids` JSON NOT NULL,`violation_ids` JSON NOT NULL,`policy_ext_id` BIGINT(20),`policy_id`VARCHAR(64),INDEX newrelic_alerts_incident_customer_id_index (`customer_id`),INDEX newrelic_alerts_incident_policy_ext_id_index (`policy_ext_id`),INDEX newrelic_alerts_incident_policy_id_index (`policy_id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// DBDropNewrelicAlertsIncidentTable will drop the NewrelicAlertsIncident table
func DBDropNewrelicAlertsIncidentTable(ctx context.Context, db *sql.DB) error {
	q := "DROP TABLE IF EXISTS `newrelic_alerts_incident`"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBDropNewrelicAlertsIncidentTableTx will drop the NewrelicAlertsIncident table using the provided transaction
func DBDropNewrelicAlertsIncidentTableTx(ctx context.Context, tx *sql.Tx) error {
	q := "DROP TABLE IF EXISTS `newrelic_alerts_incident`"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// CalculateChecksum will calculate a checksum of the SHA1 of all field values
func (t *NewrelicAlertsIncident) CalculateChecksum() string {
	return orm.HashStrings(
		orm.ToString(t.ID),
		orm.ToString(t.CustomerID),
		orm.ToString(t.ExtID),
		orm.ToString(t.OpenedAt),
		orm.ToString(t.ClosedAt),
		orm.ToString(t.IncidentPreference),
		orm.ToString(t.ViolationExtIds),
		orm.ToString(t.ViolationIds),
		orm.ToString(t.PolicyExtID),
		orm.ToString(t.PolicyID),
	)
}

// DBCreate will create a new NewrelicAlertsIncident record in the database
func (t *NewrelicAlertsIncident) DBCreate(ctx context.Context, db *sql.DB) (sql.Result, error) {
	q := "INSERT INTO `newrelic_alerts_incident` (`newrelic_alerts_incident`.`id`,`newrelic_alerts_incident`.`checksum`,`newrelic_alerts_incident`.`customer_id`,`newrelic_alerts_incident`.`ext_id`,`newrelic_alerts_incident`.`opened_at`,`newrelic_alerts_incident`.`closed_at`,`newrelic_alerts_incident`.`incident_preference`,`newrelic_alerts_incident`.`violation_ext_ids`,`newrelic_alerts_incident`.`violation_ids`,`newrelic_alerts_incident`.`policy_ext_id`,`newrelic_alerts_incident`.`policy_id`) VALUES (?,?,?,?,?,?,?,?,?,?,?)"
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
		orm.ToSQLInt64(t.OpenedAt),
		orm.ToSQLInt64(t.ClosedAt),
		orm.ToSQLString(t.IncidentPreference),
		orm.ToSQLString(t.ViolationExtIds),
		orm.ToSQLString(t.ViolationIds),
		orm.ToSQLInt64(t.PolicyExtID),
		orm.ToSQLString(t.PolicyID),
	)
}

// DBCreateTx will create a new NewrelicAlertsIncident record in the database using the provided transaction
func (t *NewrelicAlertsIncident) DBCreateTx(ctx context.Context, tx *sql.Tx) (sql.Result, error) {
	q := "INSERT INTO `newrelic_alerts_incident` (`newrelic_alerts_incident`.`id`,`newrelic_alerts_incident`.`checksum`,`newrelic_alerts_incident`.`customer_id`,`newrelic_alerts_incident`.`ext_id`,`newrelic_alerts_incident`.`opened_at`,`newrelic_alerts_incident`.`closed_at`,`newrelic_alerts_incident`.`incident_preference`,`newrelic_alerts_incident`.`violation_ext_ids`,`newrelic_alerts_incident`.`violation_ids`,`newrelic_alerts_incident`.`policy_ext_id`,`newrelic_alerts_incident`.`policy_id`) VALUES (?,?,?,?,?,?,?,?,?,?,?)"
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
		orm.ToSQLInt64(t.OpenedAt),
		orm.ToSQLInt64(t.ClosedAt),
		orm.ToSQLString(t.IncidentPreference),
		orm.ToSQLString(t.ViolationExtIds),
		orm.ToSQLString(t.ViolationIds),
		orm.ToSQLInt64(t.PolicyExtID),
		orm.ToSQLString(t.PolicyID),
	)
}

// DBCreateIgnoreDuplicate will upsert the NewrelicAlertsIncident record in the database
func (t *NewrelicAlertsIncident) DBCreateIgnoreDuplicate(ctx context.Context, db *sql.DB) (sql.Result, error) {
	q := "INSERT INTO `newrelic_alerts_incident` (`newrelic_alerts_incident`.`id`,`newrelic_alerts_incident`.`checksum`,`newrelic_alerts_incident`.`customer_id`,`newrelic_alerts_incident`.`ext_id`,`newrelic_alerts_incident`.`opened_at`,`newrelic_alerts_incident`.`closed_at`,`newrelic_alerts_incident`.`incident_preference`,`newrelic_alerts_incident`.`violation_ext_ids`,`newrelic_alerts_incident`.`violation_ids`,`newrelic_alerts_incident`.`policy_ext_id`,`newrelic_alerts_incident`.`policy_id`) VALUES (?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
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
		orm.ToSQLInt64(t.OpenedAt),
		orm.ToSQLInt64(t.ClosedAt),
		orm.ToSQLString(t.IncidentPreference),
		orm.ToSQLString(t.ViolationExtIds),
		orm.ToSQLString(t.ViolationIds),
		orm.ToSQLInt64(t.PolicyExtID),
		orm.ToSQLString(t.PolicyID),
	)
}

// DBCreateIgnoreDuplicateTx will upsert the NewrelicAlertsIncident record in the database using the provided transaction
func (t *NewrelicAlertsIncident) DBCreateIgnoreDuplicateTx(ctx context.Context, tx *sql.Tx) (sql.Result, error) {
	q := "INSERT INTO `newrelic_alerts_incident` (`newrelic_alerts_incident`.`id`,`newrelic_alerts_incident`.`checksum`,`newrelic_alerts_incident`.`customer_id`,`newrelic_alerts_incident`.`ext_id`,`newrelic_alerts_incident`.`opened_at`,`newrelic_alerts_incident`.`closed_at`,`newrelic_alerts_incident`.`incident_preference`,`newrelic_alerts_incident`.`violation_ext_ids`,`newrelic_alerts_incident`.`violation_ids`,`newrelic_alerts_incident`.`policy_ext_id`,`newrelic_alerts_incident`.`policy_id`) VALUES (?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
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
		orm.ToSQLInt64(t.OpenedAt),
		orm.ToSQLInt64(t.ClosedAt),
		orm.ToSQLString(t.IncidentPreference),
		orm.ToSQLString(t.ViolationExtIds),
		orm.ToSQLString(t.ViolationIds),
		orm.ToSQLInt64(t.PolicyExtID),
		orm.ToSQLString(t.PolicyID),
	)
}

// DeleteAllNewrelicAlertsIncidents deletes all NewrelicAlertsIncident records in the database with optional filters
func DeleteAllNewrelicAlertsIncidents(ctx context.Context, db *sql.DB, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(NewrelicAlertsIncidentTableName),
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

// DeleteAllNewrelicAlertsIncidentsTx deletes all NewrelicAlertsIncident records in the database with optional filters using the provided transaction
func DeleteAllNewrelicAlertsIncidentsTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(NewrelicAlertsIncidentTableName),
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

// DBDelete will delete this NewrelicAlertsIncident record in the database
func (t *NewrelicAlertsIncident) DBDelete(ctx context.Context, db *sql.DB) (bool, error) {
	q := "DELETE FROM `newrelic_alerts_incident` WHERE `id` = ?"
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

// DBDeleteTx will delete this NewrelicAlertsIncident record in the database using the provided transaction
func (t *NewrelicAlertsIncident) DBDeleteTx(ctx context.Context, tx *sql.Tx) (bool, error) {
	q := "DELETE FROM `newrelic_alerts_incident` WHERE `id` = ?"
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

// DBUpdate will update the NewrelicAlertsIncident record in the database
func (t *NewrelicAlertsIncident) DBUpdate(ctx context.Context, db *sql.DB) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `newrelic_alerts_incident` SET `checksum`=?,`customer_id`=?,`ext_id`=?,`opened_at`=?,`closed_at`=?,`incident_preference`=?,`violation_ext_ids`=?,`violation_ids`=?,`policy_ext_id`=?,`policy_id`=? WHERE `id`=?"
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLInt64(t.ExtID),
		orm.ToSQLInt64(t.OpenedAt),
		orm.ToSQLInt64(t.ClosedAt),
		orm.ToSQLString(t.IncidentPreference),
		orm.ToSQLString(t.ViolationExtIds),
		orm.ToSQLString(t.ViolationIds),
		orm.ToSQLInt64(t.PolicyExtID),
		orm.ToSQLString(t.PolicyID),
		orm.ToSQLString(t.ID),
	)
}

// DBUpdateTx will update the NewrelicAlertsIncident record in the database using the provided transaction
func (t *NewrelicAlertsIncident) DBUpdateTx(ctx context.Context, tx *sql.Tx) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `newrelic_alerts_incident` SET `checksum`=?,`customer_id`=?,`ext_id`=?,`opened_at`=?,`closed_at`=?,`incident_preference`=?,`violation_ext_ids`=?,`violation_ids`=?,`policy_ext_id`=?,`policy_id`=? WHERE `id`=?"
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLInt64(t.ExtID),
		orm.ToSQLInt64(t.OpenedAt),
		orm.ToSQLInt64(t.ClosedAt),
		orm.ToSQLString(t.IncidentPreference),
		orm.ToSQLString(t.ViolationExtIds),
		orm.ToSQLString(t.ViolationIds),
		orm.ToSQLInt64(t.PolicyExtID),
		orm.ToSQLString(t.PolicyID),
		orm.ToSQLString(t.ID),
	)
}

// DBUpsert will upsert the NewrelicAlertsIncident record in the database
func (t *NewrelicAlertsIncident) DBUpsert(ctx context.Context, db *sql.DB, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `newrelic_alerts_incident` (`newrelic_alerts_incident`.`id`,`newrelic_alerts_incident`.`checksum`,`newrelic_alerts_incident`.`customer_id`,`newrelic_alerts_incident`.`ext_id`,`newrelic_alerts_incident`.`opened_at`,`newrelic_alerts_incident`.`closed_at`,`newrelic_alerts_incident`.`incident_preference`,`newrelic_alerts_incident`.`violation_ext_ids`,`newrelic_alerts_incident`.`violation_ids`,`newrelic_alerts_incident`.`policy_ext_id`,`newrelic_alerts_incident`.`policy_id`) VALUES (?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `newrelic_alerts_incident` (`newrelic_alerts_incident`.`id`,`newrelic_alerts_incident`.`checksum`,`newrelic_alerts_incident`.`customer_id`,`newrelic_alerts_incident`.`ext_id`,`newrelic_alerts_incident`.`opened_at`,`newrelic_alerts_incident`.`closed_at`,`newrelic_alerts_incident`.`incident_preference`,`newrelic_alerts_incident`.`violation_ext_ids`,`newrelic_alerts_incident`.`violation_ids`,`newrelic_alerts_incident`.`policy_ext_id`,`newrelic_alerts_incident`.`policy_id`) VALUES (?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`customer_id`=VALUES(`customer_id`),`ext_id`=VALUES(`ext_id`),`opened_at`=VALUES(`opened_at`),`closed_at`=VALUES(`closed_at`),`incident_preference`=VALUES(`incident_preference`),`violation_ext_ids`=VALUES(`violation_ext_ids`),`violation_ids`=VALUES(`violation_ids`),`policy_ext_id`=VALUES(`policy_ext_id`),`policy_id`=VALUES(`policy_id`)"
	}
	r, err := db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLInt64(t.ExtID),
		orm.ToSQLInt64(t.OpenedAt),
		orm.ToSQLInt64(t.ClosedAt),
		orm.ToSQLString(t.IncidentPreference),
		orm.ToSQLString(t.ViolationExtIds),
		orm.ToSQLString(t.ViolationIds),
		orm.ToSQLInt64(t.PolicyExtID),
		orm.ToSQLString(t.PolicyID),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBUpsertTx will upsert the NewrelicAlertsIncident record in the database using the provided transaction
func (t *NewrelicAlertsIncident) DBUpsertTx(ctx context.Context, tx *sql.Tx, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `newrelic_alerts_incident` (`newrelic_alerts_incident`.`id`,`newrelic_alerts_incident`.`checksum`,`newrelic_alerts_incident`.`customer_id`,`newrelic_alerts_incident`.`ext_id`,`newrelic_alerts_incident`.`opened_at`,`newrelic_alerts_incident`.`closed_at`,`newrelic_alerts_incident`.`incident_preference`,`newrelic_alerts_incident`.`violation_ext_ids`,`newrelic_alerts_incident`.`violation_ids`,`newrelic_alerts_incident`.`policy_ext_id`,`newrelic_alerts_incident`.`policy_id`) VALUES (?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `newrelic_alerts_incident` (`newrelic_alerts_incident`.`id`,`newrelic_alerts_incident`.`checksum`,`newrelic_alerts_incident`.`customer_id`,`newrelic_alerts_incident`.`ext_id`,`newrelic_alerts_incident`.`opened_at`,`newrelic_alerts_incident`.`closed_at`,`newrelic_alerts_incident`.`incident_preference`,`newrelic_alerts_incident`.`violation_ext_ids`,`newrelic_alerts_incident`.`violation_ids`,`newrelic_alerts_incident`.`policy_ext_id`,`newrelic_alerts_incident`.`policy_id`) VALUES (?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`customer_id`=VALUES(`customer_id`),`ext_id`=VALUES(`ext_id`),`opened_at`=VALUES(`opened_at`),`closed_at`=VALUES(`closed_at`),`incident_preference`=VALUES(`incident_preference`),`violation_ext_ids`=VALUES(`violation_ext_ids`),`violation_ids`=VALUES(`violation_ids`),`policy_ext_id`=VALUES(`policy_ext_id`),`policy_id`=VALUES(`policy_id`)"
	}
	r, err := tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLInt64(t.ExtID),
		orm.ToSQLInt64(t.OpenedAt),
		orm.ToSQLInt64(t.ClosedAt),
		orm.ToSQLString(t.IncidentPreference),
		orm.ToSQLString(t.ViolationExtIds),
		orm.ToSQLString(t.ViolationIds),
		orm.ToSQLInt64(t.PolicyExtID),
		orm.ToSQLString(t.PolicyID),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBFindOne will find a NewrelicAlertsIncident record in the database with the primary key
func (t *NewrelicAlertsIncident) DBFindOne(ctx context.Context, db *sql.DB, value string) (bool, error) {
	q := "SELECT `newrelic_alerts_incident`.`id`,`newrelic_alerts_incident`.`checksum`,`newrelic_alerts_incident`.`customer_id`,`newrelic_alerts_incident`.`ext_id`,`newrelic_alerts_incident`.`opened_at`,`newrelic_alerts_incident`.`closed_at`,`newrelic_alerts_incident`.`incident_preference`,`newrelic_alerts_incident`.`violation_ext_ids`,`newrelic_alerts_incident`.`violation_ids`,`newrelic_alerts_incident`.`policy_ext_id`,`newrelic_alerts_incident`.`policy_id` FROM `newrelic_alerts_incident` WHERE `id` = ? LIMIT 1"
	row := db.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CustomerID sql.NullString
	var _ExtID sql.NullInt64
	var _OpenedAt sql.NullInt64
	var _ClosedAt sql.NullInt64
	var _IncidentPreference sql.NullString
	var _ViolationExtIds sql.NullString
	var _ViolationIds sql.NullString
	var _PolicyExtID sql.NullInt64
	var _PolicyID sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_ExtID,
		&_OpenedAt,
		&_ClosedAt,
		&_IncidentPreference,
		&_ViolationExtIds,
		&_ViolationIds,
		&_PolicyExtID,
		&_PolicyID,
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
	if _OpenedAt.Valid {
		t.SetOpenedAt(_OpenedAt.Int64)
	}
	if _ClosedAt.Valid {
		t.SetClosedAt(_ClosedAt.Int64)
	}
	if _IncidentPreference.Valid {
		t.SetIncidentPreference(_IncidentPreference.String)
	}
	if _ViolationExtIds.Valid {
		t.SetViolationExtIds(_ViolationExtIds.String)
	}
	if _ViolationIds.Valid {
		t.SetViolationIds(_ViolationIds.String)
	}
	if _PolicyExtID.Valid {
		t.SetPolicyExtID(_PolicyExtID.Int64)
	}
	if _PolicyID.Valid {
		t.SetPolicyID(_PolicyID.String)
	}
	return true, nil
}

// DBFindOneTx will find a NewrelicAlertsIncident record in the database with the primary key using the provided transaction
func (t *NewrelicAlertsIncident) DBFindOneTx(ctx context.Context, tx *sql.Tx, value string) (bool, error) {
	q := "SELECT `newrelic_alerts_incident`.`id`,`newrelic_alerts_incident`.`checksum`,`newrelic_alerts_incident`.`customer_id`,`newrelic_alerts_incident`.`ext_id`,`newrelic_alerts_incident`.`opened_at`,`newrelic_alerts_incident`.`closed_at`,`newrelic_alerts_incident`.`incident_preference`,`newrelic_alerts_incident`.`violation_ext_ids`,`newrelic_alerts_incident`.`violation_ids`,`newrelic_alerts_incident`.`policy_ext_id`,`newrelic_alerts_incident`.`policy_id` FROM `newrelic_alerts_incident` WHERE `id` = ? LIMIT 1"
	row := tx.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CustomerID sql.NullString
	var _ExtID sql.NullInt64
	var _OpenedAt sql.NullInt64
	var _ClosedAt sql.NullInt64
	var _IncidentPreference sql.NullString
	var _ViolationExtIds sql.NullString
	var _ViolationIds sql.NullString
	var _PolicyExtID sql.NullInt64
	var _PolicyID sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_ExtID,
		&_OpenedAt,
		&_ClosedAt,
		&_IncidentPreference,
		&_ViolationExtIds,
		&_ViolationIds,
		&_PolicyExtID,
		&_PolicyID,
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
	if _OpenedAt.Valid {
		t.SetOpenedAt(_OpenedAt.Int64)
	}
	if _ClosedAt.Valid {
		t.SetClosedAt(_ClosedAt.Int64)
	}
	if _IncidentPreference.Valid {
		t.SetIncidentPreference(_IncidentPreference.String)
	}
	if _ViolationExtIds.Valid {
		t.SetViolationExtIds(_ViolationExtIds.String)
	}
	if _ViolationIds.Valid {
		t.SetViolationIds(_ViolationIds.String)
	}
	if _PolicyExtID.Valid {
		t.SetPolicyExtID(_PolicyExtID.Int64)
	}
	if _PolicyID.Valid {
		t.SetPolicyID(_PolicyID.String)
	}
	return true, nil
}

// FindNewrelicAlertsIncidents will find a NewrelicAlertsIncident record in the database with the provided parameters
func FindNewrelicAlertsIncidents(ctx context.Context, db *sql.DB, _params ...interface{}) ([]*NewrelicAlertsIncident, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("customer_id"),
		orm.Column("ext_id"),
		orm.Column("opened_at"),
		orm.Column("closed_at"),
		orm.Column("incident_preference"),
		orm.Column("violation_ext_ids"),
		orm.Column("violation_ids"),
		orm.Column("policy_ext_id"),
		orm.Column("policy_id"),
		orm.Table(NewrelicAlertsIncidentTableName),
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
	results := make([]*NewrelicAlertsIncident, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _ExtID sql.NullInt64
		var _OpenedAt sql.NullInt64
		var _ClosedAt sql.NullInt64
		var _IncidentPreference sql.NullString
		var _ViolationExtIds sql.NullString
		var _ViolationIds sql.NullString
		var _PolicyExtID sql.NullInt64
		var _PolicyID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_ExtID,
			&_OpenedAt,
			&_ClosedAt,
			&_IncidentPreference,
			&_ViolationExtIds,
			&_ViolationIds,
			&_PolicyExtID,
			&_PolicyID,
		)
		if err != nil {
			return nil, err
		}
		t := &NewrelicAlertsIncident{}
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
		if _OpenedAt.Valid {
			t.SetOpenedAt(_OpenedAt.Int64)
		}
		if _ClosedAt.Valid {
			t.SetClosedAt(_ClosedAt.Int64)
		}
		if _IncidentPreference.Valid {
			t.SetIncidentPreference(_IncidentPreference.String)
		}
		if _ViolationExtIds.Valid {
			t.SetViolationExtIds(_ViolationExtIds.String)
		}
		if _ViolationIds.Valid {
			t.SetViolationIds(_ViolationIds.String)
		}
		if _PolicyExtID.Valid {
			t.SetPolicyExtID(_PolicyExtID.Int64)
		}
		if _PolicyID.Valid {
			t.SetPolicyID(_PolicyID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindNewrelicAlertsIncidentsTx will find a NewrelicAlertsIncident record in the database with the provided parameters using the provided transaction
func FindNewrelicAlertsIncidentsTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) ([]*NewrelicAlertsIncident, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("customer_id"),
		orm.Column("ext_id"),
		orm.Column("opened_at"),
		orm.Column("closed_at"),
		orm.Column("incident_preference"),
		orm.Column("violation_ext_ids"),
		orm.Column("violation_ids"),
		orm.Column("policy_ext_id"),
		orm.Column("policy_id"),
		orm.Table(NewrelicAlertsIncidentTableName),
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
	results := make([]*NewrelicAlertsIncident, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _ExtID sql.NullInt64
		var _OpenedAt sql.NullInt64
		var _ClosedAt sql.NullInt64
		var _IncidentPreference sql.NullString
		var _ViolationExtIds sql.NullString
		var _ViolationIds sql.NullString
		var _PolicyExtID sql.NullInt64
		var _PolicyID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_ExtID,
			&_OpenedAt,
			&_ClosedAt,
			&_IncidentPreference,
			&_ViolationExtIds,
			&_ViolationIds,
			&_PolicyExtID,
			&_PolicyID,
		)
		if err != nil {
			return nil, err
		}
		t := &NewrelicAlertsIncident{}
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
		if _OpenedAt.Valid {
			t.SetOpenedAt(_OpenedAt.Int64)
		}
		if _ClosedAt.Valid {
			t.SetClosedAt(_ClosedAt.Int64)
		}
		if _IncidentPreference.Valid {
			t.SetIncidentPreference(_IncidentPreference.String)
		}
		if _ViolationExtIds.Valid {
			t.SetViolationExtIds(_ViolationExtIds.String)
		}
		if _ViolationIds.Valid {
			t.SetViolationIds(_ViolationIds.String)
		}
		if _PolicyExtID.Valid {
			t.SetPolicyExtID(_PolicyExtID.Int64)
		}
		if _PolicyID.Valid {
			t.SetPolicyID(_PolicyID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// DBFind will find a NewrelicAlertsIncident record in the database with the provided parameters
func (t *NewrelicAlertsIncident) DBFind(ctx context.Context, db *sql.DB, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("customer_id"),
		orm.Column("ext_id"),
		orm.Column("opened_at"),
		orm.Column("closed_at"),
		orm.Column("incident_preference"),
		orm.Column("violation_ext_ids"),
		orm.Column("violation_ids"),
		orm.Column("policy_ext_id"),
		orm.Column("policy_id"),
		orm.Table(NewrelicAlertsIncidentTableName),
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
	var _OpenedAt sql.NullInt64
	var _ClosedAt sql.NullInt64
	var _IncidentPreference sql.NullString
	var _ViolationExtIds sql.NullString
	var _ViolationIds sql.NullString
	var _PolicyExtID sql.NullInt64
	var _PolicyID sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_ExtID,
		&_OpenedAt,
		&_ClosedAt,
		&_IncidentPreference,
		&_ViolationExtIds,
		&_ViolationIds,
		&_PolicyExtID,
		&_PolicyID,
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
	if _OpenedAt.Valid {
		t.SetOpenedAt(_OpenedAt.Int64)
	}
	if _ClosedAt.Valid {
		t.SetClosedAt(_ClosedAt.Int64)
	}
	if _IncidentPreference.Valid {
		t.SetIncidentPreference(_IncidentPreference.String)
	}
	if _ViolationExtIds.Valid {
		t.SetViolationExtIds(_ViolationExtIds.String)
	}
	if _ViolationIds.Valid {
		t.SetViolationIds(_ViolationIds.String)
	}
	if _PolicyExtID.Valid {
		t.SetPolicyExtID(_PolicyExtID.Int64)
	}
	if _PolicyID.Valid {
		t.SetPolicyID(_PolicyID.String)
	}
	return true, nil
}

// DBFindTx will find a NewrelicAlertsIncident record in the database with the provided parameters using the provided transaction
func (t *NewrelicAlertsIncident) DBFindTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("customer_id"),
		orm.Column("ext_id"),
		orm.Column("opened_at"),
		orm.Column("closed_at"),
		orm.Column("incident_preference"),
		orm.Column("violation_ext_ids"),
		orm.Column("violation_ids"),
		orm.Column("policy_ext_id"),
		orm.Column("policy_id"),
		orm.Table(NewrelicAlertsIncidentTableName),
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
	var _OpenedAt sql.NullInt64
	var _ClosedAt sql.NullInt64
	var _IncidentPreference sql.NullString
	var _ViolationExtIds sql.NullString
	var _ViolationIds sql.NullString
	var _PolicyExtID sql.NullInt64
	var _PolicyID sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_ExtID,
		&_OpenedAt,
		&_ClosedAt,
		&_IncidentPreference,
		&_ViolationExtIds,
		&_ViolationIds,
		&_PolicyExtID,
		&_PolicyID,
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
	if _OpenedAt.Valid {
		t.SetOpenedAt(_OpenedAt.Int64)
	}
	if _ClosedAt.Valid {
		t.SetClosedAt(_ClosedAt.Int64)
	}
	if _IncidentPreference.Valid {
		t.SetIncidentPreference(_IncidentPreference.String)
	}
	if _ViolationExtIds.Valid {
		t.SetViolationExtIds(_ViolationExtIds.String)
	}
	if _ViolationIds.Valid {
		t.SetViolationIds(_ViolationIds.String)
	}
	if _PolicyExtID.Valid {
		t.SetPolicyExtID(_PolicyExtID.Int64)
	}
	if _PolicyID.Valid {
		t.SetPolicyID(_PolicyID.String)
	}
	return true, nil
}

// CountNewrelicAlertsIncidents will find the count of NewrelicAlertsIncident records in the database
func CountNewrelicAlertsIncidents(ctx context.Context, db *sql.DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(NewrelicAlertsIncidentTableName),
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

// CountNewrelicAlertsIncidentsTx will find the count of NewrelicAlertsIncident records in the database using the provided transaction
func CountNewrelicAlertsIncidentsTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(NewrelicAlertsIncidentTableName),
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

// DBCount will find the count of NewrelicAlertsIncident records in the database
func (t *NewrelicAlertsIncident) DBCount(ctx context.Context, db *sql.DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(NewrelicAlertsIncidentTableName),
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

// DBCountTx will find the count of NewrelicAlertsIncident records in the database using the provided transaction
func (t *NewrelicAlertsIncident) DBCountTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(NewrelicAlertsIncidentTableName),
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

// DBExists will return true if the NewrelicAlertsIncident record exists in the database
func (t *NewrelicAlertsIncident) DBExists(ctx context.Context, db *sql.DB) (bool, error) {
	q := "SELECT `id` FROM `newrelic_alerts_incident` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := db.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// DBExistsTx will return true if the NewrelicAlertsIncident record exists in the database using the provided transaction
func (t *NewrelicAlertsIncident) DBExistsTx(ctx context.Context, tx *sql.Tx) (bool, error) {
	q := "SELECT `id` FROM `newrelic_alerts_incident` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := tx.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// PrimaryKeyColumn returns the column name for the primary key
func (t *NewrelicAlertsIncident) PrimaryKeyColumn() string {
	return NewrelicAlertsIncidentColumnID
}

// PrimaryKeyColumnType returns the primary key column Go type as a string
func (t *NewrelicAlertsIncident) PrimaryKeyColumnType() string {
	return "string"
}

// PrimaryKey returns the primary key column value
func (t *NewrelicAlertsIncident) PrimaryKey() interface{} {
	return t.ID
}
