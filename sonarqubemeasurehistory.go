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

var _ Model = (*SonarqubeMeasureHistory)(nil)
var _ CSVWriter = (*SonarqubeMeasureHistory)(nil)
var _ JSONWriter = (*SonarqubeMeasureHistory)(nil)
var _ Checksum = (*SonarqubeMeasureHistory)(nil)

// SonarqubeMeasureHistoryTableName is the name of the table in SQL
const SonarqubeMeasureHistoryTableName = "sonarqube_measure_history"

var SonarqubeMeasureHistoryColumns = []string{
	"id",
	"checksum",
	"customer_id",
	"component_id_ext_id",
	"component_id_id",
	"component_key_ext_id",
	"ext_id",
	"history",
}

// SonarqubeMeasureHistory table
type SonarqubeMeasureHistory struct {
	Checksum          *string `json:"checksum,omitempty"`
	ComponentIDExtID  string  `json:"component_id_ext_id"`
	ComponentIDID     string  `json:"component_id_id"`
	ComponentKeyExtID string  `json:"component_key_ext_id"`
	CustomerID        string  `json:"customer_id"`
	ExtID             string  `json:"ext_id"`
	History           string  `json:"history"`
	ID                string  `json:"id"`
}

// TableName returns the SQL table name for SonarqubeMeasureHistory and satifies the Model interface
func (t *SonarqubeMeasureHistory) TableName() string {
	return SonarqubeMeasureHistoryTableName
}

// ToCSV will serialize the SonarqubeMeasureHistory instance to a CSV compatible array of strings
func (t *SonarqubeMeasureHistory) ToCSV() []string {
	return []string{
		t.ID,
		t.CalculateChecksum(),
		t.CustomerID,
		t.ComponentIDExtID,
		t.ComponentIDID,
		t.ComponentKeyExtID,
		t.ExtID,
		t.History,
	}
}

// WriteCSV will serialize the SonarqubeMeasureHistory instance to the writer as CSV and satisfies the CSVWriter interface
func (t *SonarqubeMeasureHistory) WriteCSV(w *csv.Writer) error {
	return w.Write(t.ToCSV())
}

// WriteJSON will serialize the SonarqubeMeasureHistory instance to the writer as JSON and satisfies the JSONWriter interface
func (t *SonarqubeMeasureHistory) WriteJSON(w io.Writer, indent ...bool) error {
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

// NewSonarqubeMeasureHistoryReader creates a JSON reader which can read in SonarqubeMeasureHistory objects serialized as JSON either as an array, single object or json new lines
// and writes each SonarqubeMeasureHistory to the channel provided
func NewSonarqubeMeasureHistoryReader(r io.Reader, ch chan<- SonarqubeMeasureHistory) error {
	return orm.Deserialize(r, func(buf json.RawMessage) error {
		dec := json.NewDecoder(bytes.NewBuffer(buf))
		e := SonarqubeMeasureHistory{}
		if err := dec.Decode(&e); err != nil {
			return err
		}
		ch <- e
		return nil
	})
}

// NewCSVSonarqubeMeasureHistoryReaderDir will read the reader as CSV and emit each record to the channel provided
func NewCSVSonarqubeMeasureHistoryReader(r io.Reader, ch chan<- SonarqubeMeasureHistory) error {
	cr := csv.NewReader(r)
	for {
		record, err := cr.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		ch <- SonarqubeMeasureHistory{
			ID:                record[0],
			Checksum:          fromStringPointer(record[1]),
			CustomerID:        record[2],
			ComponentIDExtID:  record[3],
			ComponentIDID:     record[4],
			ComponentKeyExtID: record[5],
			ExtID:             record[6],
			History:           record[7],
		}
	}
	return nil
}

// NewCSVSonarqubeMeasureHistoryReaderFile will read the file as a CSV and emit each record to the channel provided
func NewCSVSonarqubeMeasureHistoryReaderFile(fp string, ch chan<- SonarqubeMeasureHistory) error {
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
	return NewCSVSonarqubeMeasureHistoryReader(fc, ch)
}

// NewCSVSonarqubeMeasureHistoryReaderDir will read the sonarqube_measure_history.csv.gz file as a CSV and emit each record to the channel provided
func NewCSVSonarqubeMeasureHistoryReaderDir(dir string, ch chan<- SonarqubeMeasureHistory) error {
	return NewCSVSonarqubeMeasureHistoryReaderFile(filepath.Join(dir, "sonarqube_measure_history.csv.gz"), ch)
}

// SonarqubeMeasureHistoryCSVDeduper is a function callback which takes the existing value (a) and the new value (b)
// and the return value should be the one to use (or a new one, if applicable). return nil
// to skip processing of this record
type SonarqubeMeasureHistoryCSVDeduper func(a SonarqubeMeasureHistory, b SonarqubeMeasureHistory) *SonarqubeMeasureHistory

// SonarqubeMeasureHistoryCSVDedupeDisabled is set on whether the CSV writer should de-dupe values by key
var SonarqubeMeasureHistoryCSVDedupeDisabled bool

// NewSonarqubeMeasureHistoryCSVWriterSize creates a batch writer that will write each SonarqubeMeasureHistory into a CSV file
// this method will automatically de-duplicate entries using the primary key. if the checksum
// for a newer item with the same primary key doesn't match a previously sent item, the newer
// one will be used
func NewSonarqubeMeasureHistoryCSVWriterSize(w io.Writer, size int, dedupers ...SonarqubeMeasureHistoryCSVDeduper) (chan SonarqubeMeasureHistory, chan bool, error) {
	cw := csv.NewWriter(w)
	ch := make(chan SonarqubeMeasureHistory, size)
	done := make(chan bool)
	go func() {
		defer func() { done <- true }()
		dodedupe := !SonarqubeMeasureHistoryCSVDedupeDisabled
		var kv map[string]*SonarqubeMeasureHistory
		var deduper SonarqubeMeasureHistoryCSVDeduper
		if dedupers != nil && len(dedupers) > 0 {
			deduper = dedupers[0]
			dodedupe = true
		}
		if dodedupe {
			kv = make(map[string]*SonarqubeMeasureHistory)
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

// SonarqubeMeasureHistoryCSVDefaultSize is the default channel buffer size if not provided
var SonarqubeMeasureHistoryCSVDefaultSize = 100

// NewSonarqubeMeasureHistoryCSVWriter creates a batch writer that will write each SonarqubeMeasureHistory into a CSV file
func NewSonarqubeMeasureHistoryCSVWriter(w io.Writer, dedupers ...SonarqubeMeasureHistoryCSVDeduper) (chan SonarqubeMeasureHistory, chan bool, error) {
	return NewSonarqubeMeasureHistoryCSVWriterSize(w, SonarqubeMeasureHistoryCSVDefaultSize, dedupers...)
}

// NewSonarqubeMeasureHistoryCSVWriterDir creates a batch writer that will write each SonarqubeMeasureHistory into a CSV file named sonarqube_measure_history.csv.gz in dir
func NewSonarqubeMeasureHistoryCSVWriterDir(dir string, dedupers ...SonarqubeMeasureHistoryCSVDeduper) (chan SonarqubeMeasureHistory, chan bool, error) {
	return NewSonarqubeMeasureHistoryCSVWriterFile(filepath.Join(dir, "sonarqube_measure_history.csv.gz"), dedupers...)
}

// NewSonarqubeMeasureHistoryCSVWriterFile creates a batch writer that will write each SonarqubeMeasureHistory into a CSV file
func NewSonarqubeMeasureHistoryCSVWriterFile(fn string, dedupers ...SonarqubeMeasureHistoryCSVDeduper) (chan SonarqubeMeasureHistory, chan bool, error) {
	f, err := os.Create(fn)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening CSV file %s. %v", fn, err)
	}
	var fc io.WriteCloser = f
	if filepath.Ext(fn) == ".gz" {
		w, _ := gzip.NewWriterLevel(f, gzip.BestCompression)
		fc = w
	}
	ch, done, err := NewSonarqubeMeasureHistoryCSVWriter(fc, dedupers...)
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

type SonarqubeMeasureHistoryDBAction func(ctx context.Context, db *sql.DB, record SonarqubeMeasureHistory) error

// NewSonarqubeMeasureHistoryDBWriterSize creates a DB writer that will write each issue into the DB
func NewSonarqubeMeasureHistoryDBWriterSize(ctx context.Context, db *sql.DB, errors chan<- error, size int, actions ...SonarqubeMeasureHistoryDBAction) (chan SonarqubeMeasureHistory, chan bool, error) {
	ch := make(chan SonarqubeMeasureHistory, size)
	done := make(chan bool)
	var action SonarqubeMeasureHistoryDBAction
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

// NewSonarqubeMeasureHistoryDBWriter creates a DB writer that will write each issue into the DB
func NewSonarqubeMeasureHistoryDBWriter(ctx context.Context, db *sql.DB, errors chan<- error, actions ...SonarqubeMeasureHistoryDBAction) (chan SonarqubeMeasureHistory, chan bool, error) {
	return NewSonarqubeMeasureHistoryDBWriterSize(ctx, db, errors, 100, actions...)
}

// SonarqubeMeasureHistoryColumnID is the ID SQL column name for the SonarqubeMeasureHistory table
const SonarqubeMeasureHistoryColumnID = "id"

// SonarqubeMeasureHistoryEscapedColumnID is the escaped ID SQL column name for the SonarqubeMeasureHistory table
const SonarqubeMeasureHistoryEscapedColumnID = "`id`"

// SonarqubeMeasureHistoryColumnChecksum is the Checksum SQL column name for the SonarqubeMeasureHistory table
const SonarqubeMeasureHistoryColumnChecksum = "checksum"

// SonarqubeMeasureHistoryEscapedColumnChecksum is the escaped Checksum SQL column name for the SonarqubeMeasureHistory table
const SonarqubeMeasureHistoryEscapedColumnChecksum = "`checksum`"

// SonarqubeMeasureHistoryColumnCustomerID is the CustomerID SQL column name for the SonarqubeMeasureHistory table
const SonarqubeMeasureHistoryColumnCustomerID = "customer_id"

// SonarqubeMeasureHistoryEscapedColumnCustomerID is the escaped CustomerID SQL column name for the SonarqubeMeasureHistory table
const SonarqubeMeasureHistoryEscapedColumnCustomerID = "`customer_id`"

// SonarqubeMeasureHistoryColumnComponentIDExtID is the ComponentIDExtID SQL column name for the SonarqubeMeasureHistory table
const SonarqubeMeasureHistoryColumnComponentIDExtID = "component_id_ext_id"

// SonarqubeMeasureHistoryEscapedColumnComponentIDExtID is the escaped ComponentIDExtID SQL column name for the SonarqubeMeasureHistory table
const SonarqubeMeasureHistoryEscapedColumnComponentIDExtID = "`component_id_ext_id`"

// SonarqubeMeasureHistoryColumnComponentIDID is the ComponentIDID SQL column name for the SonarqubeMeasureHistory table
const SonarqubeMeasureHistoryColumnComponentIDID = "component_id_id"

// SonarqubeMeasureHistoryEscapedColumnComponentIDID is the escaped ComponentIDID SQL column name for the SonarqubeMeasureHistory table
const SonarqubeMeasureHistoryEscapedColumnComponentIDID = "`component_id_id`"

// SonarqubeMeasureHistoryColumnComponentKeyExtID is the ComponentKeyExtID SQL column name for the SonarqubeMeasureHistory table
const SonarqubeMeasureHistoryColumnComponentKeyExtID = "component_key_ext_id"

// SonarqubeMeasureHistoryEscapedColumnComponentKeyExtID is the escaped ComponentKeyExtID SQL column name for the SonarqubeMeasureHistory table
const SonarqubeMeasureHistoryEscapedColumnComponentKeyExtID = "`component_key_ext_id`"

// SonarqubeMeasureHistoryColumnExtID is the ExtID SQL column name for the SonarqubeMeasureHistory table
const SonarqubeMeasureHistoryColumnExtID = "ext_id"

// SonarqubeMeasureHistoryEscapedColumnExtID is the escaped ExtID SQL column name for the SonarqubeMeasureHistory table
const SonarqubeMeasureHistoryEscapedColumnExtID = "`ext_id`"

// SonarqubeMeasureHistoryColumnHistory is the History SQL column name for the SonarqubeMeasureHistory table
const SonarqubeMeasureHistoryColumnHistory = "history"

// SonarqubeMeasureHistoryEscapedColumnHistory is the escaped History SQL column name for the SonarqubeMeasureHistory table
const SonarqubeMeasureHistoryEscapedColumnHistory = "`history`"

// GetID will return the SonarqubeMeasureHistory ID value
func (t *SonarqubeMeasureHistory) GetID() string {
	return t.ID
}

// SetID will set the SonarqubeMeasureHistory ID value
func (t *SonarqubeMeasureHistory) SetID(v string) {
	t.ID = v
}

// FindSonarqubeMeasureHistoryByID will find a SonarqubeMeasureHistory by ID
func FindSonarqubeMeasureHistoryByID(ctx context.Context, db *sql.DB, value string) (*SonarqubeMeasureHistory, error) {
	q := "SELECT `sonarqube_measure_history`.`id`,`sonarqube_measure_history`.`checksum`,`sonarqube_measure_history`.`customer_id`,`sonarqube_measure_history`.`component_id_ext_id`,`sonarqube_measure_history`.`component_id_id`,`sonarqube_measure_history`.`component_key_ext_id`,`sonarqube_measure_history`.`ext_id`,`sonarqube_measure_history`.`history` FROM `sonarqube_measure_history` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CustomerID sql.NullString
	var _ComponentIDExtID sql.NullString
	var _ComponentIDID sql.NullString
	var _ComponentKeyExtID sql.NullString
	var _ExtID sql.NullString
	var _History sql.NullString
	err := db.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_ComponentIDExtID,
		&_ComponentIDID,
		&_ComponentKeyExtID,
		&_ExtID,
		&_History,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &SonarqubeMeasureHistory{}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _Checksum.Valid {
		t.SetChecksum(_Checksum.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _ComponentIDExtID.Valid {
		t.SetComponentIDExtID(_ComponentIDExtID.String)
	}
	if _ComponentIDID.Valid {
		t.SetComponentIDID(_ComponentIDID.String)
	}
	if _ComponentKeyExtID.Valid {
		t.SetComponentKeyExtID(_ComponentKeyExtID.String)
	}
	if _ExtID.Valid {
		t.SetExtID(_ExtID.String)
	}
	if _History.Valid {
		t.SetHistory(_History.String)
	}
	return t, nil
}

// FindSonarqubeMeasureHistoryByIDTx will find a SonarqubeMeasureHistory by ID using the provided transaction
func FindSonarqubeMeasureHistoryByIDTx(ctx context.Context, tx *sql.Tx, value string) (*SonarqubeMeasureHistory, error) {
	q := "SELECT `sonarqube_measure_history`.`id`,`sonarqube_measure_history`.`checksum`,`sonarqube_measure_history`.`customer_id`,`sonarqube_measure_history`.`component_id_ext_id`,`sonarqube_measure_history`.`component_id_id`,`sonarqube_measure_history`.`component_key_ext_id`,`sonarqube_measure_history`.`ext_id`,`sonarqube_measure_history`.`history` FROM `sonarqube_measure_history` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CustomerID sql.NullString
	var _ComponentIDExtID sql.NullString
	var _ComponentIDID sql.NullString
	var _ComponentKeyExtID sql.NullString
	var _ExtID sql.NullString
	var _History sql.NullString
	err := tx.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_ComponentIDExtID,
		&_ComponentIDID,
		&_ComponentKeyExtID,
		&_ExtID,
		&_History,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &SonarqubeMeasureHistory{}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _Checksum.Valid {
		t.SetChecksum(_Checksum.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _ComponentIDExtID.Valid {
		t.SetComponentIDExtID(_ComponentIDExtID.String)
	}
	if _ComponentIDID.Valid {
		t.SetComponentIDID(_ComponentIDID.String)
	}
	if _ComponentKeyExtID.Valid {
		t.SetComponentKeyExtID(_ComponentKeyExtID.String)
	}
	if _ExtID.Valid {
		t.SetExtID(_ExtID.String)
	}
	if _History.Valid {
		t.SetHistory(_History.String)
	}
	return t, nil
}

// GetChecksum will return the SonarqubeMeasureHistory Checksum value
func (t *SonarqubeMeasureHistory) GetChecksum() string {
	if t.Checksum == nil {
		return ""
	}
	return *t.Checksum
}

// SetChecksum will set the SonarqubeMeasureHistory Checksum value
func (t *SonarqubeMeasureHistory) SetChecksum(v string) {
	t.Checksum = &v
}

// GetCustomerID will return the SonarqubeMeasureHistory CustomerID value
func (t *SonarqubeMeasureHistory) GetCustomerID() string {
	return t.CustomerID
}

// SetCustomerID will set the SonarqubeMeasureHistory CustomerID value
func (t *SonarqubeMeasureHistory) SetCustomerID(v string) {
	t.CustomerID = v
}

// FindSonarqubeMeasureHistoriesByCustomerID will find all SonarqubeMeasureHistorys by the CustomerID value
func FindSonarqubeMeasureHistoriesByCustomerID(ctx context.Context, db *sql.DB, value string) ([]*SonarqubeMeasureHistory, error) {
	q := "SELECT `sonarqube_measure_history`.`id`,`sonarqube_measure_history`.`checksum`,`sonarqube_measure_history`.`customer_id`,`sonarqube_measure_history`.`component_id_ext_id`,`sonarqube_measure_history`.`component_id_id`,`sonarqube_measure_history`.`component_key_ext_id`,`sonarqube_measure_history`.`ext_id`,`sonarqube_measure_history`.`history` FROM `sonarqube_measure_history` WHERE `customer_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*SonarqubeMeasureHistory, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _ComponentIDExtID sql.NullString
		var _ComponentIDID sql.NullString
		var _ComponentKeyExtID sql.NullString
		var _ExtID sql.NullString
		var _History sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_ComponentIDExtID,
			&_ComponentIDID,
			&_ComponentKeyExtID,
			&_ExtID,
			&_History,
		)
		if err != nil {
			return nil, err
		}
		t := &SonarqubeMeasureHistory{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _ComponentIDExtID.Valid {
			t.SetComponentIDExtID(_ComponentIDExtID.String)
		}
		if _ComponentIDID.Valid {
			t.SetComponentIDID(_ComponentIDID.String)
		}
		if _ComponentKeyExtID.Valid {
			t.SetComponentKeyExtID(_ComponentKeyExtID.String)
		}
		if _ExtID.Valid {
			t.SetExtID(_ExtID.String)
		}
		if _History.Valid {
			t.SetHistory(_History.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindSonarqubeMeasureHistoriesByCustomerIDTx will find all SonarqubeMeasureHistorys by the CustomerID value using the provided transaction
func FindSonarqubeMeasureHistoriesByCustomerIDTx(ctx context.Context, tx *sql.Tx, value string) ([]*SonarqubeMeasureHistory, error) {
	q := "SELECT `sonarqube_measure_history`.`id`,`sonarqube_measure_history`.`checksum`,`sonarqube_measure_history`.`customer_id`,`sonarqube_measure_history`.`component_id_ext_id`,`sonarqube_measure_history`.`component_id_id`,`sonarqube_measure_history`.`component_key_ext_id`,`sonarqube_measure_history`.`ext_id`,`sonarqube_measure_history`.`history` FROM `sonarqube_measure_history` WHERE `customer_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*SonarqubeMeasureHistory, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _ComponentIDExtID sql.NullString
		var _ComponentIDID sql.NullString
		var _ComponentKeyExtID sql.NullString
		var _ExtID sql.NullString
		var _History sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_ComponentIDExtID,
			&_ComponentIDID,
			&_ComponentKeyExtID,
			&_ExtID,
			&_History,
		)
		if err != nil {
			return nil, err
		}
		t := &SonarqubeMeasureHistory{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _ComponentIDExtID.Valid {
			t.SetComponentIDExtID(_ComponentIDExtID.String)
		}
		if _ComponentIDID.Valid {
			t.SetComponentIDID(_ComponentIDID.String)
		}
		if _ComponentKeyExtID.Valid {
			t.SetComponentKeyExtID(_ComponentKeyExtID.String)
		}
		if _ExtID.Valid {
			t.SetExtID(_ExtID.String)
		}
		if _History.Valid {
			t.SetHistory(_History.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetComponentIDExtID will return the SonarqubeMeasureHistory ComponentIDExtID value
func (t *SonarqubeMeasureHistory) GetComponentIDExtID() string {
	return t.ComponentIDExtID
}

// SetComponentIDExtID will set the SonarqubeMeasureHistory ComponentIDExtID value
func (t *SonarqubeMeasureHistory) SetComponentIDExtID(v string) {
	t.ComponentIDExtID = v
}

// GetComponentIDID will return the SonarqubeMeasureHistory ComponentIDID value
func (t *SonarqubeMeasureHistory) GetComponentIDID() string {
	return t.ComponentIDID
}

// SetComponentIDID will set the SonarqubeMeasureHistory ComponentIDID value
func (t *SonarqubeMeasureHistory) SetComponentIDID(v string) {
	t.ComponentIDID = v
}

// FindSonarqubeMeasureHistoriesByComponentIDID will find all SonarqubeMeasureHistorys by the ComponentIDID value
func FindSonarqubeMeasureHistoriesByComponentIDID(ctx context.Context, db *sql.DB, value string) ([]*SonarqubeMeasureHistory, error) {
	q := "SELECT `sonarqube_measure_history`.`id`,`sonarqube_measure_history`.`checksum`,`sonarqube_measure_history`.`customer_id`,`sonarqube_measure_history`.`component_id_ext_id`,`sonarqube_measure_history`.`component_id_id`,`sonarqube_measure_history`.`component_key_ext_id`,`sonarqube_measure_history`.`ext_id`,`sonarqube_measure_history`.`history` FROM `sonarqube_measure_history` WHERE `component_id_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*SonarqubeMeasureHistory, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _ComponentIDExtID sql.NullString
		var _ComponentIDID sql.NullString
		var _ComponentKeyExtID sql.NullString
		var _ExtID sql.NullString
		var _History sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_ComponentIDExtID,
			&_ComponentIDID,
			&_ComponentKeyExtID,
			&_ExtID,
			&_History,
		)
		if err != nil {
			return nil, err
		}
		t := &SonarqubeMeasureHistory{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _ComponentIDExtID.Valid {
			t.SetComponentIDExtID(_ComponentIDExtID.String)
		}
		if _ComponentIDID.Valid {
			t.SetComponentIDID(_ComponentIDID.String)
		}
		if _ComponentKeyExtID.Valid {
			t.SetComponentKeyExtID(_ComponentKeyExtID.String)
		}
		if _ExtID.Valid {
			t.SetExtID(_ExtID.String)
		}
		if _History.Valid {
			t.SetHistory(_History.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindSonarqubeMeasureHistoriesByComponentIDIDTx will find all SonarqubeMeasureHistorys by the ComponentIDID value using the provided transaction
func FindSonarqubeMeasureHistoriesByComponentIDIDTx(ctx context.Context, tx *sql.Tx, value string) ([]*SonarqubeMeasureHistory, error) {
	q := "SELECT `sonarqube_measure_history`.`id`,`sonarqube_measure_history`.`checksum`,`sonarqube_measure_history`.`customer_id`,`sonarqube_measure_history`.`component_id_ext_id`,`sonarqube_measure_history`.`component_id_id`,`sonarqube_measure_history`.`component_key_ext_id`,`sonarqube_measure_history`.`ext_id`,`sonarqube_measure_history`.`history` FROM `sonarqube_measure_history` WHERE `component_id_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*SonarqubeMeasureHistory, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _ComponentIDExtID sql.NullString
		var _ComponentIDID sql.NullString
		var _ComponentKeyExtID sql.NullString
		var _ExtID sql.NullString
		var _History sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_ComponentIDExtID,
			&_ComponentIDID,
			&_ComponentKeyExtID,
			&_ExtID,
			&_History,
		)
		if err != nil {
			return nil, err
		}
		t := &SonarqubeMeasureHistory{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _ComponentIDExtID.Valid {
			t.SetComponentIDExtID(_ComponentIDExtID.String)
		}
		if _ComponentIDID.Valid {
			t.SetComponentIDID(_ComponentIDID.String)
		}
		if _ComponentKeyExtID.Valid {
			t.SetComponentKeyExtID(_ComponentKeyExtID.String)
		}
		if _ExtID.Valid {
			t.SetExtID(_ExtID.String)
		}
		if _History.Valid {
			t.SetHistory(_History.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetComponentKeyExtID will return the SonarqubeMeasureHistory ComponentKeyExtID value
func (t *SonarqubeMeasureHistory) GetComponentKeyExtID() string {
	return t.ComponentKeyExtID
}

// SetComponentKeyExtID will set the SonarqubeMeasureHistory ComponentKeyExtID value
func (t *SonarqubeMeasureHistory) SetComponentKeyExtID(v string) {
	t.ComponentKeyExtID = v
}

// GetExtID will return the SonarqubeMeasureHistory ExtID value
func (t *SonarqubeMeasureHistory) GetExtID() string {
	return t.ExtID
}

// SetExtID will set the SonarqubeMeasureHistory ExtID value
func (t *SonarqubeMeasureHistory) SetExtID(v string) {
	t.ExtID = v
}

// GetHistory will return the SonarqubeMeasureHistory History value
func (t *SonarqubeMeasureHistory) GetHistory() string {
	return t.History
}

// SetHistory will set the SonarqubeMeasureHistory History value
func (t *SonarqubeMeasureHistory) SetHistory(v string) {
	t.History = v
}

func (t *SonarqubeMeasureHistory) toTimestamp(value time.Time) *timestamp.Timestamp {
	ts, _ := ptypes.TimestampProto(value)
	return ts
}

// DBCreateSonarqubeMeasureHistoryTable will create the SonarqubeMeasureHistory table
func DBCreateSonarqubeMeasureHistoryTable(ctx context.Context, db *sql.DB) error {
	q := "CREATE TABLE `sonarqube_measure_history` (`id`VARCHAR(64) NOT NULL PRIMARY KEY,`checksum`CHAR(64),`customer_id`VARCHAR(64) NOT NULL,`component_id_ext_id` VARCHAR(255) NOT NULL,`component_id_id` VARCHAR(64) NOT NULL,`component_key_ext_id` VARCHAR(255) NOT NULL,`ext_id` TEXT NOT NULL,`history` JSON NOT NULL,INDEX sonarqube_measure_history_customer_id_index (`customer_id`),INDEX sonarqube_measure_history_component_id_id_index (`component_id_id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBCreateSonarqubeMeasureHistoryTableTx will create the SonarqubeMeasureHistory table using the provided transction
func DBCreateSonarqubeMeasureHistoryTableTx(ctx context.Context, tx *sql.Tx) error {
	q := "CREATE TABLE `sonarqube_measure_history` (`id`VARCHAR(64) NOT NULL PRIMARY KEY,`checksum`CHAR(64),`customer_id`VARCHAR(64) NOT NULL,`component_id_ext_id` VARCHAR(255) NOT NULL,`component_id_id` VARCHAR(64) NOT NULL,`component_key_ext_id` VARCHAR(255) NOT NULL,`ext_id` TEXT NOT NULL,`history` JSON NOT NULL,INDEX sonarqube_measure_history_customer_id_index (`customer_id`),INDEX sonarqube_measure_history_component_id_id_index (`component_id_id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// DBDropSonarqubeMeasureHistoryTable will drop the SonarqubeMeasureHistory table
func DBDropSonarqubeMeasureHistoryTable(ctx context.Context, db *sql.DB) error {
	q := "DROP TABLE IF EXISTS `sonarqube_measure_history`"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBDropSonarqubeMeasureHistoryTableTx will drop the SonarqubeMeasureHistory table using the provided transaction
func DBDropSonarqubeMeasureHistoryTableTx(ctx context.Context, tx *sql.Tx) error {
	q := "DROP TABLE IF EXISTS `sonarqube_measure_history`"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// CalculateChecksum will calculate a checksum of the SHA1 of all field values
func (t *SonarqubeMeasureHistory) CalculateChecksum() string {
	return orm.HashStrings(
		orm.ToString(t.ID),
		orm.ToString(t.CustomerID),
		orm.ToString(t.ComponentIDExtID),
		orm.ToString(t.ComponentIDID),
		orm.ToString(t.ComponentKeyExtID),
		orm.ToString(t.ExtID),
		orm.ToString(t.History),
	)
}

// DBCreate will create a new SonarqubeMeasureHistory record in the database
func (t *SonarqubeMeasureHistory) DBCreate(ctx context.Context, db *sql.DB) (sql.Result, error) {
	q := "INSERT INTO `sonarqube_measure_history` (`sonarqube_measure_history`.`id`,`sonarqube_measure_history`.`checksum`,`sonarqube_measure_history`.`customer_id`,`sonarqube_measure_history`.`component_id_ext_id`,`sonarqube_measure_history`.`component_id_id`,`sonarqube_measure_history`.`component_key_ext_id`,`sonarqube_measure_history`.`ext_id`,`sonarqube_measure_history`.`history`) VALUES (?,?,?,?,?,?,?,?)"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.ComponentIDExtID),
		orm.ToSQLString(t.ComponentIDID),
		orm.ToSQLString(t.ComponentKeyExtID),
		orm.ToSQLString(t.ExtID),
		orm.ToSQLString(t.History),
	)
}

// DBCreateTx will create a new SonarqubeMeasureHistory record in the database using the provided transaction
func (t *SonarqubeMeasureHistory) DBCreateTx(ctx context.Context, tx *sql.Tx) (sql.Result, error) {
	q := "INSERT INTO `sonarqube_measure_history` (`sonarqube_measure_history`.`id`,`sonarqube_measure_history`.`checksum`,`sonarqube_measure_history`.`customer_id`,`sonarqube_measure_history`.`component_id_ext_id`,`sonarqube_measure_history`.`component_id_id`,`sonarqube_measure_history`.`component_key_ext_id`,`sonarqube_measure_history`.`ext_id`,`sonarqube_measure_history`.`history`) VALUES (?,?,?,?,?,?,?,?)"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.ComponentIDExtID),
		orm.ToSQLString(t.ComponentIDID),
		orm.ToSQLString(t.ComponentKeyExtID),
		orm.ToSQLString(t.ExtID),
		orm.ToSQLString(t.History),
	)
}

// DBCreateIgnoreDuplicate will upsert the SonarqubeMeasureHistory record in the database
func (t *SonarqubeMeasureHistory) DBCreateIgnoreDuplicate(ctx context.Context, db *sql.DB) (sql.Result, error) {
	q := "INSERT INTO `sonarqube_measure_history` (`sonarqube_measure_history`.`id`,`sonarqube_measure_history`.`checksum`,`sonarqube_measure_history`.`customer_id`,`sonarqube_measure_history`.`component_id_ext_id`,`sonarqube_measure_history`.`component_id_id`,`sonarqube_measure_history`.`component_key_ext_id`,`sonarqube_measure_history`.`ext_id`,`sonarqube_measure_history`.`history`) VALUES (?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.ComponentIDExtID),
		orm.ToSQLString(t.ComponentIDID),
		orm.ToSQLString(t.ComponentKeyExtID),
		orm.ToSQLString(t.ExtID),
		orm.ToSQLString(t.History),
	)
}

// DBCreateIgnoreDuplicateTx will upsert the SonarqubeMeasureHistory record in the database using the provided transaction
func (t *SonarqubeMeasureHistory) DBCreateIgnoreDuplicateTx(ctx context.Context, tx *sql.Tx) (sql.Result, error) {
	q := "INSERT INTO `sonarqube_measure_history` (`sonarqube_measure_history`.`id`,`sonarqube_measure_history`.`checksum`,`sonarqube_measure_history`.`customer_id`,`sonarqube_measure_history`.`component_id_ext_id`,`sonarqube_measure_history`.`component_id_id`,`sonarqube_measure_history`.`component_key_ext_id`,`sonarqube_measure_history`.`ext_id`,`sonarqube_measure_history`.`history`) VALUES (?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.ComponentIDExtID),
		orm.ToSQLString(t.ComponentIDID),
		orm.ToSQLString(t.ComponentKeyExtID),
		orm.ToSQLString(t.ExtID),
		orm.ToSQLString(t.History),
	)
}

// DeleteAllSonarqubeMeasureHistories deletes all SonarqubeMeasureHistory records in the database with optional filters
func DeleteAllSonarqubeMeasureHistories(ctx context.Context, db *sql.DB, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(SonarqubeMeasureHistoryTableName),
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

// DeleteAllSonarqubeMeasureHistoriesTx deletes all SonarqubeMeasureHistory records in the database with optional filters using the provided transaction
func DeleteAllSonarqubeMeasureHistoriesTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(SonarqubeMeasureHistoryTableName),
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

// DBDelete will delete this SonarqubeMeasureHistory record in the database
func (t *SonarqubeMeasureHistory) DBDelete(ctx context.Context, db *sql.DB) (bool, error) {
	q := "DELETE FROM `sonarqube_measure_history` WHERE `id` = ?"
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

// DBDeleteTx will delete this SonarqubeMeasureHistory record in the database using the provided transaction
func (t *SonarqubeMeasureHistory) DBDeleteTx(ctx context.Context, tx *sql.Tx) (bool, error) {
	q := "DELETE FROM `sonarqube_measure_history` WHERE `id` = ?"
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

// DBUpdate will update the SonarqubeMeasureHistory record in the database
func (t *SonarqubeMeasureHistory) DBUpdate(ctx context.Context, db *sql.DB) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `sonarqube_measure_history` SET `checksum`=?,`customer_id`=?,`component_id_ext_id`=?,`component_id_id`=?,`component_key_ext_id`=?,`ext_id`=?,`history`=? WHERE `id`=?"
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.ComponentIDExtID),
		orm.ToSQLString(t.ComponentIDID),
		orm.ToSQLString(t.ComponentKeyExtID),
		orm.ToSQLString(t.ExtID),
		orm.ToSQLString(t.History),
		orm.ToSQLString(t.ID),
	)
}

// DBUpdateTx will update the SonarqubeMeasureHistory record in the database using the provided transaction
func (t *SonarqubeMeasureHistory) DBUpdateTx(ctx context.Context, tx *sql.Tx) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `sonarqube_measure_history` SET `checksum`=?,`customer_id`=?,`component_id_ext_id`=?,`component_id_id`=?,`component_key_ext_id`=?,`ext_id`=?,`history`=? WHERE `id`=?"
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.ComponentIDExtID),
		orm.ToSQLString(t.ComponentIDID),
		orm.ToSQLString(t.ComponentKeyExtID),
		orm.ToSQLString(t.ExtID),
		orm.ToSQLString(t.History),
		orm.ToSQLString(t.ID),
	)
}

// DBUpsert will upsert the SonarqubeMeasureHistory record in the database
func (t *SonarqubeMeasureHistory) DBUpsert(ctx context.Context, db *sql.DB, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `sonarqube_measure_history` (`sonarqube_measure_history`.`id`,`sonarqube_measure_history`.`checksum`,`sonarqube_measure_history`.`customer_id`,`sonarqube_measure_history`.`component_id_ext_id`,`sonarqube_measure_history`.`component_id_id`,`sonarqube_measure_history`.`component_key_ext_id`,`sonarqube_measure_history`.`ext_id`,`sonarqube_measure_history`.`history`) VALUES (?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `sonarqube_measure_history` (`sonarqube_measure_history`.`id`,`sonarqube_measure_history`.`checksum`,`sonarqube_measure_history`.`customer_id`,`sonarqube_measure_history`.`component_id_ext_id`,`sonarqube_measure_history`.`component_id_id`,`sonarqube_measure_history`.`component_key_ext_id`,`sonarqube_measure_history`.`ext_id`,`sonarqube_measure_history`.`history`) VALUES (?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`customer_id`=VALUES(`customer_id`),`component_id_ext_id`=VALUES(`component_id_ext_id`),`component_id_id`=VALUES(`component_id_id`),`component_key_ext_id`=VALUES(`component_key_ext_id`),`ext_id`=VALUES(`ext_id`),`history`=VALUES(`history`)"
	}
	r, err := db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.ComponentIDExtID),
		orm.ToSQLString(t.ComponentIDID),
		orm.ToSQLString(t.ComponentKeyExtID),
		orm.ToSQLString(t.ExtID),
		orm.ToSQLString(t.History),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBUpsertTx will upsert the SonarqubeMeasureHistory record in the database using the provided transaction
func (t *SonarqubeMeasureHistory) DBUpsertTx(ctx context.Context, tx *sql.Tx, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `sonarqube_measure_history` (`sonarqube_measure_history`.`id`,`sonarqube_measure_history`.`checksum`,`sonarqube_measure_history`.`customer_id`,`sonarqube_measure_history`.`component_id_ext_id`,`sonarqube_measure_history`.`component_id_id`,`sonarqube_measure_history`.`component_key_ext_id`,`sonarqube_measure_history`.`ext_id`,`sonarqube_measure_history`.`history`) VALUES (?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `sonarqube_measure_history` (`sonarqube_measure_history`.`id`,`sonarqube_measure_history`.`checksum`,`sonarqube_measure_history`.`customer_id`,`sonarqube_measure_history`.`component_id_ext_id`,`sonarqube_measure_history`.`component_id_id`,`sonarqube_measure_history`.`component_key_ext_id`,`sonarqube_measure_history`.`ext_id`,`sonarqube_measure_history`.`history`) VALUES (?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`customer_id`=VALUES(`customer_id`),`component_id_ext_id`=VALUES(`component_id_ext_id`),`component_id_id`=VALUES(`component_id_id`),`component_key_ext_id`=VALUES(`component_key_ext_id`),`ext_id`=VALUES(`ext_id`),`history`=VALUES(`history`)"
	}
	r, err := tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.ComponentIDExtID),
		orm.ToSQLString(t.ComponentIDID),
		orm.ToSQLString(t.ComponentKeyExtID),
		orm.ToSQLString(t.ExtID),
		orm.ToSQLString(t.History),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBFindOne will find a SonarqubeMeasureHistory record in the database with the primary key
func (t *SonarqubeMeasureHistory) DBFindOne(ctx context.Context, db *sql.DB, value string) (bool, error) {
	q := "SELECT `sonarqube_measure_history`.`id`,`sonarqube_measure_history`.`checksum`,`sonarqube_measure_history`.`customer_id`,`sonarqube_measure_history`.`component_id_ext_id`,`sonarqube_measure_history`.`component_id_id`,`sonarqube_measure_history`.`component_key_ext_id`,`sonarqube_measure_history`.`ext_id`,`sonarqube_measure_history`.`history` FROM `sonarqube_measure_history` WHERE `id` = ? LIMIT 1"
	row := db.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CustomerID sql.NullString
	var _ComponentIDExtID sql.NullString
	var _ComponentIDID sql.NullString
	var _ComponentKeyExtID sql.NullString
	var _ExtID sql.NullString
	var _History sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_ComponentIDExtID,
		&_ComponentIDID,
		&_ComponentKeyExtID,
		&_ExtID,
		&_History,
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
	if _ComponentIDExtID.Valid {
		t.SetComponentIDExtID(_ComponentIDExtID.String)
	}
	if _ComponentIDID.Valid {
		t.SetComponentIDID(_ComponentIDID.String)
	}
	if _ComponentKeyExtID.Valid {
		t.SetComponentKeyExtID(_ComponentKeyExtID.String)
	}
	if _ExtID.Valid {
		t.SetExtID(_ExtID.String)
	}
	if _History.Valid {
		t.SetHistory(_History.String)
	}
	return true, nil
}

// DBFindOneTx will find a SonarqubeMeasureHistory record in the database with the primary key using the provided transaction
func (t *SonarqubeMeasureHistory) DBFindOneTx(ctx context.Context, tx *sql.Tx, value string) (bool, error) {
	q := "SELECT `sonarqube_measure_history`.`id`,`sonarqube_measure_history`.`checksum`,`sonarqube_measure_history`.`customer_id`,`sonarqube_measure_history`.`component_id_ext_id`,`sonarqube_measure_history`.`component_id_id`,`sonarqube_measure_history`.`component_key_ext_id`,`sonarqube_measure_history`.`ext_id`,`sonarqube_measure_history`.`history` FROM `sonarqube_measure_history` WHERE `id` = ? LIMIT 1"
	row := tx.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CustomerID sql.NullString
	var _ComponentIDExtID sql.NullString
	var _ComponentIDID sql.NullString
	var _ComponentKeyExtID sql.NullString
	var _ExtID sql.NullString
	var _History sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_ComponentIDExtID,
		&_ComponentIDID,
		&_ComponentKeyExtID,
		&_ExtID,
		&_History,
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
	if _ComponentIDExtID.Valid {
		t.SetComponentIDExtID(_ComponentIDExtID.String)
	}
	if _ComponentIDID.Valid {
		t.SetComponentIDID(_ComponentIDID.String)
	}
	if _ComponentKeyExtID.Valid {
		t.SetComponentKeyExtID(_ComponentKeyExtID.String)
	}
	if _ExtID.Valid {
		t.SetExtID(_ExtID.String)
	}
	if _History.Valid {
		t.SetHistory(_History.String)
	}
	return true, nil
}

// FindSonarqubeMeasureHistories will find a SonarqubeMeasureHistory record in the database with the provided parameters
func FindSonarqubeMeasureHistories(ctx context.Context, db *sql.DB, _params ...interface{}) ([]*SonarqubeMeasureHistory, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("customer_id"),
		orm.Column("component_id_ext_id"),
		orm.Column("component_id_id"),
		orm.Column("component_key_ext_id"),
		orm.Column("ext_id"),
		orm.Column("history"),
		orm.Table(SonarqubeMeasureHistoryTableName),
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
	results := make([]*SonarqubeMeasureHistory, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _ComponentIDExtID sql.NullString
		var _ComponentIDID sql.NullString
		var _ComponentKeyExtID sql.NullString
		var _ExtID sql.NullString
		var _History sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_ComponentIDExtID,
			&_ComponentIDID,
			&_ComponentKeyExtID,
			&_ExtID,
			&_History,
		)
		if err != nil {
			return nil, err
		}
		t := &SonarqubeMeasureHistory{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _ComponentIDExtID.Valid {
			t.SetComponentIDExtID(_ComponentIDExtID.String)
		}
		if _ComponentIDID.Valid {
			t.SetComponentIDID(_ComponentIDID.String)
		}
		if _ComponentKeyExtID.Valid {
			t.SetComponentKeyExtID(_ComponentKeyExtID.String)
		}
		if _ExtID.Valid {
			t.SetExtID(_ExtID.String)
		}
		if _History.Valid {
			t.SetHistory(_History.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindSonarqubeMeasureHistoriesTx will find a SonarqubeMeasureHistory record in the database with the provided parameters using the provided transaction
func FindSonarqubeMeasureHistoriesTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) ([]*SonarqubeMeasureHistory, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("customer_id"),
		orm.Column("component_id_ext_id"),
		orm.Column("component_id_id"),
		orm.Column("component_key_ext_id"),
		orm.Column("ext_id"),
		orm.Column("history"),
		orm.Table(SonarqubeMeasureHistoryTableName),
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
	results := make([]*SonarqubeMeasureHistory, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _ComponentIDExtID sql.NullString
		var _ComponentIDID sql.NullString
		var _ComponentKeyExtID sql.NullString
		var _ExtID sql.NullString
		var _History sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_ComponentIDExtID,
			&_ComponentIDID,
			&_ComponentKeyExtID,
			&_ExtID,
			&_History,
		)
		if err != nil {
			return nil, err
		}
		t := &SonarqubeMeasureHistory{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _ComponentIDExtID.Valid {
			t.SetComponentIDExtID(_ComponentIDExtID.String)
		}
		if _ComponentIDID.Valid {
			t.SetComponentIDID(_ComponentIDID.String)
		}
		if _ComponentKeyExtID.Valid {
			t.SetComponentKeyExtID(_ComponentKeyExtID.String)
		}
		if _ExtID.Valid {
			t.SetExtID(_ExtID.String)
		}
		if _History.Valid {
			t.SetHistory(_History.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// DBFind will find a SonarqubeMeasureHistory record in the database with the provided parameters
func (t *SonarqubeMeasureHistory) DBFind(ctx context.Context, db *sql.DB, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("customer_id"),
		orm.Column("component_id_ext_id"),
		orm.Column("component_id_id"),
		orm.Column("component_key_ext_id"),
		orm.Column("ext_id"),
		orm.Column("history"),
		orm.Table(SonarqubeMeasureHistoryTableName),
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
	var _ComponentIDExtID sql.NullString
	var _ComponentIDID sql.NullString
	var _ComponentKeyExtID sql.NullString
	var _ExtID sql.NullString
	var _History sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_ComponentIDExtID,
		&_ComponentIDID,
		&_ComponentKeyExtID,
		&_ExtID,
		&_History,
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
	if _ComponentIDExtID.Valid {
		t.SetComponentIDExtID(_ComponentIDExtID.String)
	}
	if _ComponentIDID.Valid {
		t.SetComponentIDID(_ComponentIDID.String)
	}
	if _ComponentKeyExtID.Valid {
		t.SetComponentKeyExtID(_ComponentKeyExtID.String)
	}
	if _ExtID.Valid {
		t.SetExtID(_ExtID.String)
	}
	if _History.Valid {
		t.SetHistory(_History.String)
	}
	return true, nil
}

// DBFindTx will find a SonarqubeMeasureHistory record in the database with the provided parameters using the provided transaction
func (t *SonarqubeMeasureHistory) DBFindTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("customer_id"),
		orm.Column("component_id_ext_id"),
		orm.Column("component_id_id"),
		orm.Column("component_key_ext_id"),
		orm.Column("ext_id"),
		orm.Column("history"),
		orm.Table(SonarqubeMeasureHistoryTableName),
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
	var _ComponentIDExtID sql.NullString
	var _ComponentIDID sql.NullString
	var _ComponentKeyExtID sql.NullString
	var _ExtID sql.NullString
	var _History sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_ComponentIDExtID,
		&_ComponentIDID,
		&_ComponentKeyExtID,
		&_ExtID,
		&_History,
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
	if _ComponentIDExtID.Valid {
		t.SetComponentIDExtID(_ComponentIDExtID.String)
	}
	if _ComponentIDID.Valid {
		t.SetComponentIDID(_ComponentIDID.String)
	}
	if _ComponentKeyExtID.Valid {
		t.SetComponentKeyExtID(_ComponentKeyExtID.String)
	}
	if _ExtID.Valid {
		t.SetExtID(_ExtID.String)
	}
	if _History.Valid {
		t.SetHistory(_History.String)
	}
	return true, nil
}

// CountSonarqubeMeasureHistories will find the count of SonarqubeMeasureHistory records in the database
func CountSonarqubeMeasureHistories(ctx context.Context, db *sql.DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(SonarqubeMeasureHistoryTableName),
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

// CountSonarqubeMeasureHistoriesTx will find the count of SonarqubeMeasureHistory records in the database using the provided transaction
func CountSonarqubeMeasureHistoriesTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(SonarqubeMeasureHistoryTableName),
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

// DBCount will find the count of SonarqubeMeasureHistory records in the database
func (t *SonarqubeMeasureHistory) DBCount(ctx context.Context, db *sql.DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(SonarqubeMeasureHistoryTableName),
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

// DBCountTx will find the count of SonarqubeMeasureHistory records in the database using the provided transaction
func (t *SonarqubeMeasureHistory) DBCountTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(SonarqubeMeasureHistoryTableName),
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

// DBExists will return true if the SonarqubeMeasureHistory record exists in the database
func (t *SonarqubeMeasureHistory) DBExists(ctx context.Context, db *sql.DB) (bool, error) {
	q := "SELECT `id` FROM `sonarqube_measure_history` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := db.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// DBExistsTx will return true if the SonarqubeMeasureHistory record exists in the database using the provided transaction
func (t *SonarqubeMeasureHistory) DBExistsTx(ctx context.Context, tx *sql.Tx) (bool, error) {
	q := "SELECT `id` FROM `sonarqube_measure_history` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := tx.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// PrimaryKeyColumn returns the column name for the primary key
func (t *SonarqubeMeasureHistory) PrimaryKeyColumn() string {
	return SonarqubeMeasureHistoryColumnID
}

// PrimaryKeyColumnType returns the primary key column Go type as a string
func (t *SonarqubeMeasureHistory) PrimaryKeyColumnType() string {
	return "string"
}

// PrimaryKey returns the primary key column value
func (t *SonarqubeMeasureHistory) PrimaryKey() interface{} {
	return t.ID
}
