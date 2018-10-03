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

var _ Model = (*SonarqubeMetric)(nil)
var _ CSVWriter = (*SonarqubeMetric)(nil)
var _ JSONWriter = (*SonarqubeMetric)(nil)
var _ Checksum = (*SonarqubeMetric)(nil)

// SonarqubeMetricTableName is the name of the table in SQL
const SonarqubeMetricTableName = "sonarqube_metric"

var SonarqubeMetricColumns = []string{
	"id",
	"checksum",
	"customer_id",
	"project_ext_id",
	"project_id",
	"project_key_ext_id",
	"ext_id",
	"date",
	"metric",
	"value",
}

// SonarqubeMetric table
type SonarqubeMetric struct {
	Checksum        *string `json:"checksum,omitempty"`
	CustomerID      string  `json:"customer_id"`
	Date            int64   `json:"date"`
	ExtID           string  `json:"ext_id"`
	ID              string  `json:"id"`
	Metric          string  `json:"metric"`
	ProjectExtID    string  `json:"project_ext_id"`
	ProjectID       string  `json:"project_id"`
	ProjectKeyExtID string  `json:"project_key_ext_id"`
	Value           string  `json:"value"`
}

// TableName returns the SQL table name for SonarqubeMetric and satifies the Model interface
func (t *SonarqubeMetric) TableName() string {
	return SonarqubeMetricTableName
}

// ToCSV will serialize the SonarqubeMetric instance to a CSV compatible array of strings
func (t *SonarqubeMetric) ToCSV() []string {
	return []string{
		t.ID,
		t.CalculateChecksum(),
		t.CustomerID,
		t.ProjectExtID,
		t.ProjectID,
		t.ProjectKeyExtID,
		t.ExtID,
		toCSVString(t.Date),
		t.Metric,
		t.Value,
	}
}

// WriteCSV will serialize the SonarqubeMetric instance to the writer as CSV and satisfies the CSVWriter interface
func (t *SonarqubeMetric) WriteCSV(w *csv.Writer) error {
	return w.Write(t.ToCSV())
}

// WriteJSON will serialize the SonarqubeMetric instance to the writer as JSON and satisfies the JSONWriter interface
func (t *SonarqubeMetric) WriteJSON(w io.Writer, indent ...bool) error {
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

// NewSonarqubeMetricReader creates a JSON reader which can read in SonarqubeMetric objects serialized as JSON either as an array, single object or json new lines
// and writes each SonarqubeMetric to the channel provided
func NewSonarqubeMetricReader(r io.Reader, ch chan<- SonarqubeMetric) error {
	return orm.Deserialize(r, func(buf json.RawMessage) error {
		dec := json.NewDecoder(bytes.NewBuffer(buf))
		e := SonarqubeMetric{}
		if err := dec.Decode(&e); err != nil {
			return err
		}
		ch <- e
		return nil
	})
}

// NewCSVSonarqubeMetricReaderDir will read the reader as CSV and emit each record to the channel provided
func NewCSVSonarqubeMetricReader(r io.Reader, ch chan<- SonarqubeMetric) error {
	cr := csv.NewReader(r)
	for {
		record, err := cr.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		ch <- SonarqubeMetric{
			ID:              record[0],
			Checksum:        fromStringPointer(record[1]),
			CustomerID:      record[2],
			ProjectExtID:    record[3],
			ProjectID:       record[4],
			ProjectKeyExtID: record[5],
			ExtID:           record[6],
			Date:            fromCSVInt64(record[7]),
			Metric:          record[8],
			Value:           record[9],
		}
	}
	return nil
}

// NewCSVSonarqubeMetricReaderFile will read the file as a CSV and emit each record to the channel provided
func NewCSVSonarqubeMetricReaderFile(fp string, ch chan<- SonarqubeMetric) error {
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
	return NewCSVSonarqubeMetricReader(fc, ch)
}

// NewCSVSonarqubeMetricReaderDir will read the sonarqube_metric.csv.gz file as a CSV and emit each record to the channel provided
func NewCSVSonarqubeMetricReaderDir(dir string, ch chan<- SonarqubeMetric) error {
	return NewCSVSonarqubeMetricReaderFile(filepath.Join(dir, "sonarqube_metric.csv.gz"), ch)
}

// SonarqubeMetricCSVDeduper is a function callback which takes the existing value (a) and the new value (b)
// and the return value should be the one to use (or a new one, if applicable). return nil
// to skip processing of this record
type SonarqubeMetricCSVDeduper func(a SonarqubeMetric, b SonarqubeMetric) *SonarqubeMetric

// SonarqubeMetricCSVDedupeDisabled is set on whether the CSV writer should de-dupe values by key
var SonarqubeMetricCSVDedupeDisabled bool

// NewSonarqubeMetricCSVWriterSize creates a batch writer that will write each SonarqubeMetric into a CSV file
// this method will automatically de-duplicate entries using the primary key. if the checksum
// for a newer item with the same primary key doesn't match a previously sent item, the newer
// one will be used
func NewSonarqubeMetricCSVWriterSize(w io.Writer, size int, dedupers ...SonarqubeMetricCSVDeduper) (chan SonarqubeMetric, chan bool, error) {
	cw := csv.NewWriter(w)
	ch := make(chan SonarqubeMetric, size)
	done := make(chan bool)
	go func() {
		defer func() { done <- true }()
		dodedupe := !SonarqubeMetricCSVDedupeDisabled
		var kv map[string]*SonarqubeMetric
		var deduper SonarqubeMetricCSVDeduper
		if dedupers != nil && len(dedupers) > 0 {
			deduper = dedupers[0]
			dodedupe = true
		}
		if dodedupe {
			kv = make(map[string]*SonarqubeMetric)
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

// SonarqubeMetricCSVDefaultSize is the default channel buffer size if not provided
var SonarqubeMetricCSVDefaultSize = 100

// NewSonarqubeMetricCSVWriter creates a batch writer that will write each SonarqubeMetric into a CSV file
func NewSonarqubeMetricCSVWriter(w io.Writer, dedupers ...SonarqubeMetricCSVDeduper) (chan SonarqubeMetric, chan bool, error) {
	return NewSonarqubeMetricCSVWriterSize(w, SonarqubeMetricCSVDefaultSize, dedupers...)
}

// NewSonarqubeMetricCSVWriterDir creates a batch writer that will write each SonarqubeMetric into a CSV file named sonarqube_metric.csv.gz in dir
func NewSonarqubeMetricCSVWriterDir(dir string, dedupers ...SonarqubeMetricCSVDeduper) (chan SonarqubeMetric, chan bool, error) {
	return NewSonarqubeMetricCSVWriterFile(filepath.Join(dir, "sonarqube_metric.csv.gz"), dedupers...)
}

// NewSonarqubeMetricCSVWriterFile creates a batch writer that will write each SonarqubeMetric into a CSV file
func NewSonarqubeMetricCSVWriterFile(fn string, dedupers ...SonarqubeMetricCSVDeduper) (chan SonarqubeMetric, chan bool, error) {
	f, err := os.Create(fn)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening CSV file %s. %v", fn, err)
	}
	var fc io.WriteCloser = f
	if filepath.Ext(fn) == ".gz" {
		w, _ := gzip.NewWriterLevel(f, gzip.BestCompression)
		fc = w
	}
	ch, done, err := NewSonarqubeMetricCSVWriter(fc, dedupers...)
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

type SonarqubeMetricDBAction func(ctx context.Context, db *sql.DB, record SonarqubeMetric) error

// NewSonarqubeMetricDBWriterSize creates a DB writer that will write each issue into the DB
func NewSonarqubeMetricDBWriterSize(ctx context.Context, db *sql.DB, errors chan<- error, size int, actions ...SonarqubeMetricDBAction) (chan SonarqubeMetric, chan bool, error) {
	ch := make(chan SonarqubeMetric, size)
	done := make(chan bool)
	var action SonarqubeMetricDBAction
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

// NewSonarqubeMetricDBWriter creates a DB writer that will write each issue into the DB
func NewSonarqubeMetricDBWriter(ctx context.Context, db *sql.DB, errors chan<- error, actions ...SonarqubeMetricDBAction) (chan SonarqubeMetric, chan bool, error) {
	return NewSonarqubeMetricDBWriterSize(ctx, db, errors, 100, actions...)
}

// SonarqubeMetricColumnID is the ID SQL column name for the SonarqubeMetric table
const SonarqubeMetricColumnID = "id"

// SonarqubeMetricEscapedColumnID is the escaped ID SQL column name for the SonarqubeMetric table
const SonarqubeMetricEscapedColumnID = "`id`"

// SonarqubeMetricColumnChecksum is the Checksum SQL column name for the SonarqubeMetric table
const SonarqubeMetricColumnChecksum = "checksum"

// SonarqubeMetricEscapedColumnChecksum is the escaped Checksum SQL column name for the SonarqubeMetric table
const SonarqubeMetricEscapedColumnChecksum = "`checksum`"

// SonarqubeMetricColumnCustomerID is the CustomerID SQL column name for the SonarqubeMetric table
const SonarqubeMetricColumnCustomerID = "customer_id"

// SonarqubeMetricEscapedColumnCustomerID is the escaped CustomerID SQL column name for the SonarqubeMetric table
const SonarqubeMetricEscapedColumnCustomerID = "`customer_id`"

// SonarqubeMetricColumnProjectExtID is the ProjectExtID SQL column name for the SonarqubeMetric table
const SonarqubeMetricColumnProjectExtID = "project_ext_id"

// SonarqubeMetricEscapedColumnProjectExtID is the escaped ProjectExtID SQL column name for the SonarqubeMetric table
const SonarqubeMetricEscapedColumnProjectExtID = "`project_ext_id`"

// SonarqubeMetricColumnProjectID is the ProjectID SQL column name for the SonarqubeMetric table
const SonarqubeMetricColumnProjectID = "project_id"

// SonarqubeMetricEscapedColumnProjectID is the escaped ProjectID SQL column name for the SonarqubeMetric table
const SonarqubeMetricEscapedColumnProjectID = "`project_id`"

// SonarqubeMetricColumnProjectKeyExtID is the ProjectKeyExtID SQL column name for the SonarqubeMetric table
const SonarqubeMetricColumnProjectKeyExtID = "project_key_ext_id"

// SonarqubeMetricEscapedColumnProjectKeyExtID is the escaped ProjectKeyExtID SQL column name for the SonarqubeMetric table
const SonarqubeMetricEscapedColumnProjectKeyExtID = "`project_key_ext_id`"

// SonarqubeMetricColumnExtID is the ExtID SQL column name for the SonarqubeMetric table
const SonarqubeMetricColumnExtID = "ext_id"

// SonarqubeMetricEscapedColumnExtID is the escaped ExtID SQL column name for the SonarqubeMetric table
const SonarqubeMetricEscapedColumnExtID = "`ext_id`"

// SonarqubeMetricColumnDate is the Date SQL column name for the SonarqubeMetric table
const SonarqubeMetricColumnDate = "date"

// SonarqubeMetricEscapedColumnDate is the escaped Date SQL column name for the SonarqubeMetric table
const SonarqubeMetricEscapedColumnDate = "`date`"

// SonarqubeMetricColumnMetric is the Metric SQL column name for the SonarqubeMetric table
const SonarqubeMetricColumnMetric = "metric"

// SonarqubeMetricEscapedColumnMetric is the escaped Metric SQL column name for the SonarqubeMetric table
const SonarqubeMetricEscapedColumnMetric = "`metric`"

// SonarqubeMetricColumnValue is the Value SQL column name for the SonarqubeMetric table
const SonarqubeMetricColumnValue = "value"

// SonarqubeMetricEscapedColumnValue is the escaped Value SQL column name for the SonarqubeMetric table
const SonarqubeMetricEscapedColumnValue = "`value`"

// GetID will return the SonarqubeMetric ID value
func (t *SonarqubeMetric) GetID() string {
	return t.ID
}

// SetID will set the SonarqubeMetric ID value
func (t *SonarqubeMetric) SetID(v string) {
	t.ID = v
}

// FindSonarqubeMetricByID will find a SonarqubeMetric by ID
func FindSonarqubeMetricByID(ctx context.Context, db *sql.DB, value string) (*SonarqubeMetric, error) {
	q := "SELECT `sonarqube_metric`.`id`,`sonarqube_metric`.`checksum`,`sonarqube_metric`.`customer_id`,`sonarqube_metric`.`project_ext_id`,`sonarqube_metric`.`project_id`,`sonarqube_metric`.`project_key_ext_id`,`sonarqube_metric`.`ext_id`,`sonarqube_metric`.`date`,`sonarqube_metric`.`metric`,`sonarqube_metric`.`value` FROM `sonarqube_metric` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CustomerID sql.NullString
	var _ProjectExtID sql.NullString
	var _ProjectID sql.NullString
	var _ProjectKeyExtID sql.NullString
	var _ExtID sql.NullString
	var _Date sql.NullInt64
	var _Metric sql.NullString
	var _Value sql.NullString
	err := db.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_ProjectExtID,
		&_ProjectID,
		&_ProjectKeyExtID,
		&_ExtID,
		&_Date,
		&_Metric,
		&_Value,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &SonarqubeMetric{}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _Checksum.Valid {
		t.SetChecksum(_Checksum.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _ProjectExtID.Valid {
		t.SetProjectExtID(_ProjectExtID.String)
	}
	if _ProjectID.Valid {
		t.SetProjectID(_ProjectID.String)
	}
	if _ProjectKeyExtID.Valid {
		t.SetProjectKeyExtID(_ProjectKeyExtID.String)
	}
	if _ExtID.Valid {
		t.SetExtID(_ExtID.String)
	}
	if _Date.Valid {
		t.SetDate(_Date.Int64)
	}
	if _Metric.Valid {
		t.SetMetric(_Metric.String)
	}
	if _Value.Valid {
		t.SetValue(_Value.String)
	}
	return t, nil
}

// FindSonarqubeMetricByIDTx will find a SonarqubeMetric by ID using the provided transaction
func FindSonarqubeMetricByIDTx(ctx context.Context, tx *sql.Tx, value string) (*SonarqubeMetric, error) {
	q := "SELECT `sonarqube_metric`.`id`,`sonarqube_metric`.`checksum`,`sonarqube_metric`.`customer_id`,`sonarqube_metric`.`project_ext_id`,`sonarqube_metric`.`project_id`,`sonarqube_metric`.`project_key_ext_id`,`sonarqube_metric`.`ext_id`,`sonarqube_metric`.`date`,`sonarqube_metric`.`metric`,`sonarqube_metric`.`value` FROM `sonarqube_metric` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CustomerID sql.NullString
	var _ProjectExtID sql.NullString
	var _ProjectID sql.NullString
	var _ProjectKeyExtID sql.NullString
	var _ExtID sql.NullString
	var _Date sql.NullInt64
	var _Metric sql.NullString
	var _Value sql.NullString
	err := tx.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_ProjectExtID,
		&_ProjectID,
		&_ProjectKeyExtID,
		&_ExtID,
		&_Date,
		&_Metric,
		&_Value,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &SonarqubeMetric{}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _Checksum.Valid {
		t.SetChecksum(_Checksum.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _ProjectExtID.Valid {
		t.SetProjectExtID(_ProjectExtID.String)
	}
	if _ProjectID.Valid {
		t.SetProjectID(_ProjectID.String)
	}
	if _ProjectKeyExtID.Valid {
		t.SetProjectKeyExtID(_ProjectKeyExtID.String)
	}
	if _ExtID.Valid {
		t.SetExtID(_ExtID.String)
	}
	if _Date.Valid {
		t.SetDate(_Date.Int64)
	}
	if _Metric.Valid {
		t.SetMetric(_Metric.String)
	}
	if _Value.Valid {
		t.SetValue(_Value.String)
	}
	return t, nil
}

// GetChecksum will return the SonarqubeMetric Checksum value
func (t *SonarqubeMetric) GetChecksum() string {
	if t.Checksum == nil {
		return ""
	}
	return *t.Checksum
}

// SetChecksum will set the SonarqubeMetric Checksum value
func (t *SonarqubeMetric) SetChecksum(v string) {
	t.Checksum = &v
}

// GetCustomerID will return the SonarqubeMetric CustomerID value
func (t *SonarqubeMetric) GetCustomerID() string {
	return t.CustomerID
}

// SetCustomerID will set the SonarqubeMetric CustomerID value
func (t *SonarqubeMetric) SetCustomerID(v string) {
	t.CustomerID = v
}

// FindSonarqubeMetricsByCustomerID will find all SonarqubeMetrics by the CustomerID value
func FindSonarqubeMetricsByCustomerID(ctx context.Context, db *sql.DB, value string) ([]*SonarqubeMetric, error) {
	q := "SELECT `sonarqube_metric`.`id`,`sonarqube_metric`.`checksum`,`sonarqube_metric`.`customer_id`,`sonarqube_metric`.`project_ext_id`,`sonarqube_metric`.`project_id`,`sonarqube_metric`.`project_key_ext_id`,`sonarqube_metric`.`ext_id`,`sonarqube_metric`.`date`,`sonarqube_metric`.`metric`,`sonarqube_metric`.`value` FROM `sonarqube_metric` WHERE `customer_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*SonarqubeMetric, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _ProjectExtID sql.NullString
		var _ProjectID sql.NullString
		var _ProjectKeyExtID sql.NullString
		var _ExtID sql.NullString
		var _Date sql.NullInt64
		var _Metric sql.NullString
		var _Value sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_ProjectExtID,
			&_ProjectID,
			&_ProjectKeyExtID,
			&_ExtID,
			&_Date,
			&_Metric,
			&_Value,
		)
		if err != nil {
			return nil, err
		}
		t := &SonarqubeMetric{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _ProjectExtID.Valid {
			t.SetProjectExtID(_ProjectExtID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _ProjectKeyExtID.Valid {
			t.SetProjectKeyExtID(_ProjectKeyExtID.String)
		}
		if _ExtID.Valid {
			t.SetExtID(_ExtID.String)
		}
		if _Date.Valid {
			t.SetDate(_Date.Int64)
		}
		if _Metric.Valid {
			t.SetMetric(_Metric.String)
		}
		if _Value.Valid {
			t.SetValue(_Value.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindSonarqubeMetricsByCustomerIDTx will find all SonarqubeMetrics by the CustomerID value using the provided transaction
func FindSonarqubeMetricsByCustomerIDTx(ctx context.Context, tx *sql.Tx, value string) ([]*SonarqubeMetric, error) {
	q := "SELECT `sonarqube_metric`.`id`,`sonarqube_metric`.`checksum`,`sonarqube_metric`.`customer_id`,`sonarqube_metric`.`project_ext_id`,`sonarqube_metric`.`project_id`,`sonarqube_metric`.`project_key_ext_id`,`sonarqube_metric`.`ext_id`,`sonarqube_metric`.`date`,`sonarqube_metric`.`metric`,`sonarqube_metric`.`value` FROM `sonarqube_metric` WHERE `customer_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*SonarqubeMetric, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _ProjectExtID sql.NullString
		var _ProjectID sql.NullString
		var _ProjectKeyExtID sql.NullString
		var _ExtID sql.NullString
		var _Date sql.NullInt64
		var _Metric sql.NullString
		var _Value sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_ProjectExtID,
			&_ProjectID,
			&_ProjectKeyExtID,
			&_ExtID,
			&_Date,
			&_Metric,
			&_Value,
		)
		if err != nil {
			return nil, err
		}
		t := &SonarqubeMetric{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _ProjectExtID.Valid {
			t.SetProjectExtID(_ProjectExtID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _ProjectKeyExtID.Valid {
			t.SetProjectKeyExtID(_ProjectKeyExtID.String)
		}
		if _ExtID.Valid {
			t.SetExtID(_ExtID.String)
		}
		if _Date.Valid {
			t.SetDate(_Date.Int64)
		}
		if _Metric.Valid {
			t.SetMetric(_Metric.String)
		}
		if _Value.Valid {
			t.SetValue(_Value.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetProjectExtID will return the SonarqubeMetric ProjectExtID value
func (t *SonarqubeMetric) GetProjectExtID() string {
	return t.ProjectExtID
}

// SetProjectExtID will set the SonarqubeMetric ProjectExtID value
func (t *SonarqubeMetric) SetProjectExtID(v string) {
	t.ProjectExtID = v
}

// GetProjectID will return the SonarqubeMetric ProjectID value
func (t *SonarqubeMetric) GetProjectID() string {
	return t.ProjectID
}

// SetProjectID will set the SonarqubeMetric ProjectID value
func (t *SonarqubeMetric) SetProjectID(v string) {
	t.ProjectID = v
}

// FindSonarqubeMetricsByProjectID will find all SonarqubeMetrics by the ProjectID value
func FindSonarqubeMetricsByProjectID(ctx context.Context, db *sql.DB, value string) ([]*SonarqubeMetric, error) {
	q := "SELECT `sonarqube_metric`.`id`,`sonarqube_metric`.`checksum`,`sonarqube_metric`.`customer_id`,`sonarqube_metric`.`project_ext_id`,`sonarqube_metric`.`project_id`,`sonarqube_metric`.`project_key_ext_id`,`sonarqube_metric`.`ext_id`,`sonarqube_metric`.`date`,`sonarqube_metric`.`metric`,`sonarqube_metric`.`value` FROM `sonarqube_metric` WHERE `project_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*SonarqubeMetric, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _ProjectExtID sql.NullString
		var _ProjectID sql.NullString
		var _ProjectKeyExtID sql.NullString
		var _ExtID sql.NullString
		var _Date sql.NullInt64
		var _Metric sql.NullString
		var _Value sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_ProjectExtID,
			&_ProjectID,
			&_ProjectKeyExtID,
			&_ExtID,
			&_Date,
			&_Metric,
			&_Value,
		)
		if err != nil {
			return nil, err
		}
		t := &SonarqubeMetric{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _ProjectExtID.Valid {
			t.SetProjectExtID(_ProjectExtID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _ProjectKeyExtID.Valid {
			t.SetProjectKeyExtID(_ProjectKeyExtID.String)
		}
		if _ExtID.Valid {
			t.SetExtID(_ExtID.String)
		}
		if _Date.Valid {
			t.SetDate(_Date.Int64)
		}
		if _Metric.Valid {
			t.SetMetric(_Metric.String)
		}
		if _Value.Valid {
			t.SetValue(_Value.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindSonarqubeMetricsByProjectIDTx will find all SonarqubeMetrics by the ProjectID value using the provided transaction
func FindSonarqubeMetricsByProjectIDTx(ctx context.Context, tx *sql.Tx, value string) ([]*SonarqubeMetric, error) {
	q := "SELECT `sonarqube_metric`.`id`,`sonarqube_metric`.`checksum`,`sonarqube_metric`.`customer_id`,`sonarqube_metric`.`project_ext_id`,`sonarqube_metric`.`project_id`,`sonarqube_metric`.`project_key_ext_id`,`sonarqube_metric`.`ext_id`,`sonarqube_metric`.`date`,`sonarqube_metric`.`metric`,`sonarqube_metric`.`value` FROM `sonarqube_metric` WHERE `project_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*SonarqubeMetric, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _ProjectExtID sql.NullString
		var _ProjectID sql.NullString
		var _ProjectKeyExtID sql.NullString
		var _ExtID sql.NullString
		var _Date sql.NullInt64
		var _Metric sql.NullString
		var _Value sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_ProjectExtID,
			&_ProjectID,
			&_ProjectKeyExtID,
			&_ExtID,
			&_Date,
			&_Metric,
			&_Value,
		)
		if err != nil {
			return nil, err
		}
		t := &SonarqubeMetric{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _ProjectExtID.Valid {
			t.SetProjectExtID(_ProjectExtID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _ProjectKeyExtID.Valid {
			t.SetProjectKeyExtID(_ProjectKeyExtID.String)
		}
		if _ExtID.Valid {
			t.SetExtID(_ExtID.String)
		}
		if _Date.Valid {
			t.SetDate(_Date.Int64)
		}
		if _Metric.Valid {
			t.SetMetric(_Metric.String)
		}
		if _Value.Valid {
			t.SetValue(_Value.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetProjectKeyExtID will return the SonarqubeMetric ProjectKeyExtID value
func (t *SonarqubeMetric) GetProjectKeyExtID() string {
	return t.ProjectKeyExtID
}

// SetProjectKeyExtID will set the SonarqubeMetric ProjectKeyExtID value
func (t *SonarqubeMetric) SetProjectKeyExtID(v string) {
	t.ProjectKeyExtID = v
}

// GetExtID will return the SonarqubeMetric ExtID value
func (t *SonarqubeMetric) GetExtID() string {
	return t.ExtID
}

// SetExtID will set the SonarqubeMetric ExtID value
func (t *SonarqubeMetric) SetExtID(v string) {
	t.ExtID = v
}

// GetDate will return the SonarqubeMetric Date value
func (t *SonarqubeMetric) GetDate() int64 {
	return t.Date
}

// SetDate will set the SonarqubeMetric Date value
func (t *SonarqubeMetric) SetDate(v int64) {
	t.Date = v
}

// GetMetric will return the SonarqubeMetric Metric value
func (t *SonarqubeMetric) GetMetric() string {
	return t.Metric
}

// SetMetric will set the SonarqubeMetric Metric value
func (t *SonarqubeMetric) SetMetric(v string) {
	t.Metric = v
}

// GetValue will return the SonarqubeMetric Value value
func (t *SonarqubeMetric) GetValue() string {
	return t.Value
}

// SetValue will set the SonarqubeMetric Value value
func (t *SonarqubeMetric) SetValue(v string) {
	t.Value = v
}

func (t *SonarqubeMetric) toTimestamp(value time.Time) *timestamp.Timestamp {
	ts, _ := ptypes.TimestampProto(value)
	return ts
}

// DBCreateSonarqubeMetricTable will create the SonarqubeMetric table
func DBCreateSonarqubeMetricTable(ctx context.Context, db *sql.DB) error {
	q := "CREATE TABLE `sonarqube_metric` (`id` VARCHAR(64) NOT NULL PRIMARY KEY,`checksum` CHAR(64),`customer_id` VARCHAR(64) NOT NULL,`project_ext_id` VARCHAR(255) NOT NULL,`project_id` VARCHAR(64) NOT NULL,`project_key_ext_id` VARCHAR(255) NOT NULL,`ext_id`VARCHAR(255) NOT NULL,`date` BIGINT NOT NULL,`metric`VARCHAR(64) NOT NULL,`value` VARCHAR(64) NOT NULL,INDEX sonarqube_metric_customer_id_index (`customer_id`),INDEX sonarqube_metric_project_id_index (`project_id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBCreateSonarqubeMetricTableTx will create the SonarqubeMetric table using the provided transction
func DBCreateSonarqubeMetricTableTx(ctx context.Context, tx *sql.Tx) error {
	q := "CREATE TABLE `sonarqube_metric` (`id` VARCHAR(64) NOT NULL PRIMARY KEY,`checksum` CHAR(64),`customer_id` VARCHAR(64) NOT NULL,`project_ext_id` VARCHAR(255) NOT NULL,`project_id` VARCHAR(64) NOT NULL,`project_key_ext_id` VARCHAR(255) NOT NULL,`ext_id`VARCHAR(255) NOT NULL,`date` BIGINT NOT NULL,`metric`VARCHAR(64) NOT NULL,`value` VARCHAR(64) NOT NULL,INDEX sonarqube_metric_customer_id_index (`customer_id`),INDEX sonarqube_metric_project_id_index (`project_id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// DBDropSonarqubeMetricTable will drop the SonarqubeMetric table
func DBDropSonarqubeMetricTable(ctx context.Context, db *sql.DB) error {
	q := "DROP TABLE IF EXISTS `sonarqube_metric`"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBDropSonarqubeMetricTableTx will drop the SonarqubeMetric table using the provided transaction
func DBDropSonarqubeMetricTableTx(ctx context.Context, tx *sql.Tx) error {
	q := "DROP TABLE IF EXISTS `sonarqube_metric`"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// CalculateChecksum will calculate a checksum of the SHA1 of all field values
func (t *SonarqubeMetric) CalculateChecksum() string {
	return orm.HashStrings(
		orm.ToString(t.ID),
		orm.ToString(t.CustomerID),
		orm.ToString(t.ProjectExtID),
		orm.ToString(t.ProjectID),
		orm.ToString(t.ProjectKeyExtID),
		orm.ToString(t.ExtID),
		orm.ToString(t.Date),
		orm.ToString(t.Metric),
		orm.ToString(t.Value),
	)
}

// DBCreate will create a new SonarqubeMetric record in the database
func (t *SonarqubeMetric) DBCreate(ctx context.Context, db *sql.DB) (sql.Result, error) {
	q := "INSERT INTO `sonarqube_metric` (`sonarqube_metric`.`id`,`sonarqube_metric`.`checksum`,`sonarqube_metric`.`customer_id`,`sonarqube_metric`.`project_ext_id`,`sonarqube_metric`.`project_id`,`sonarqube_metric`.`project_key_ext_id`,`sonarqube_metric`.`ext_id`,`sonarqube_metric`.`date`,`sonarqube_metric`.`metric`,`sonarqube_metric`.`value`) VALUES (?,?,?,?,?,?,?,?,?,?)"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.ProjectExtID),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLString(t.ProjectKeyExtID),
		orm.ToSQLString(t.ExtID),
		orm.ToSQLInt64(t.Date),
		orm.ToSQLString(t.Metric),
		orm.ToSQLString(t.Value),
	)
}

// DBCreateTx will create a new SonarqubeMetric record in the database using the provided transaction
func (t *SonarqubeMetric) DBCreateTx(ctx context.Context, tx *sql.Tx) (sql.Result, error) {
	q := "INSERT INTO `sonarqube_metric` (`sonarqube_metric`.`id`,`sonarqube_metric`.`checksum`,`sonarqube_metric`.`customer_id`,`sonarqube_metric`.`project_ext_id`,`sonarqube_metric`.`project_id`,`sonarqube_metric`.`project_key_ext_id`,`sonarqube_metric`.`ext_id`,`sonarqube_metric`.`date`,`sonarqube_metric`.`metric`,`sonarqube_metric`.`value`) VALUES (?,?,?,?,?,?,?,?,?,?)"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.ProjectExtID),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLString(t.ProjectKeyExtID),
		orm.ToSQLString(t.ExtID),
		orm.ToSQLInt64(t.Date),
		orm.ToSQLString(t.Metric),
		orm.ToSQLString(t.Value),
	)
}

// DBCreateIgnoreDuplicate will upsert the SonarqubeMetric record in the database
func (t *SonarqubeMetric) DBCreateIgnoreDuplicate(ctx context.Context, db *sql.DB) (sql.Result, error) {
	q := "INSERT INTO `sonarqube_metric` (`sonarqube_metric`.`id`,`sonarqube_metric`.`checksum`,`sonarqube_metric`.`customer_id`,`sonarqube_metric`.`project_ext_id`,`sonarqube_metric`.`project_id`,`sonarqube_metric`.`project_key_ext_id`,`sonarqube_metric`.`ext_id`,`sonarqube_metric`.`date`,`sonarqube_metric`.`metric`,`sonarqube_metric`.`value`) VALUES (?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.ProjectExtID),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLString(t.ProjectKeyExtID),
		orm.ToSQLString(t.ExtID),
		orm.ToSQLInt64(t.Date),
		orm.ToSQLString(t.Metric),
		orm.ToSQLString(t.Value),
	)
}

// DBCreateIgnoreDuplicateTx will upsert the SonarqubeMetric record in the database using the provided transaction
func (t *SonarqubeMetric) DBCreateIgnoreDuplicateTx(ctx context.Context, tx *sql.Tx) (sql.Result, error) {
	q := "INSERT INTO `sonarqube_metric` (`sonarqube_metric`.`id`,`sonarqube_metric`.`checksum`,`sonarqube_metric`.`customer_id`,`sonarqube_metric`.`project_ext_id`,`sonarqube_metric`.`project_id`,`sonarqube_metric`.`project_key_ext_id`,`sonarqube_metric`.`ext_id`,`sonarqube_metric`.`date`,`sonarqube_metric`.`metric`,`sonarqube_metric`.`value`) VALUES (?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.ProjectExtID),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLString(t.ProjectKeyExtID),
		orm.ToSQLString(t.ExtID),
		orm.ToSQLInt64(t.Date),
		orm.ToSQLString(t.Metric),
		orm.ToSQLString(t.Value),
	)
}

// DeleteAllSonarqubeMetrics deletes all SonarqubeMetric records in the database with optional filters
func DeleteAllSonarqubeMetrics(ctx context.Context, db *sql.DB, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(SonarqubeMetricTableName),
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

// DeleteAllSonarqubeMetricsTx deletes all SonarqubeMetric records in the database with optional filters using the provided transaction
func DeleteAllSonarqubeMetricsTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(SonarqubeMetricTableName),
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

// DBDelete will delete this SonarqubeMetric record in the database
func (t *SonarqubeMetric) DBDelete(ctx context.Context, db *sql.DB) (bool, error) {
	q := "DELETE FROM `sonarqube_metric` WHERE `id` = ?"
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

// DBDeleteTx will delete this SonarqubeMetric record in the database using the provided transaction
func (t *SonarqubeMetric) DBDeleteTx(ctx context.Context, tx *sql.Tx) (bool, error) {
	q := "DELETE FROM `sonarqube_metric` WHERE `id` = ?"
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

// DBUpdate will update the SonarqubeMetric record in the database
func (t *SonarqubeMetric) DBUpdate(ctx context.Context, db *sql.DB) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `sonarqube_metric` SET `checksum`=?,`customer_id`=?,`project_ext_id`=?,`project_id`=?,`project_key_ext_id`=?,`ext_id`=?,`date`=?,`metric`=?,`value`=? WHERE `id`=?"
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.ProjectExtID),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLString(t.ProjectKeyExtID),
		orm.ToSQLString(t.ExtID),
		orm.ToSQLInt64(t.Date),
		orm.ToSQLString(t.Metric),
		orm.ToSQLString(t.Value),
		orm.ToSQLString(t.ID),
	)
}

// DBUpdateTx will update the SonarqubeMetric record in the database using the provided transaction
func (t *SonarqubeMetric) DBUpdateTx(ctx context.Context, tx *sql.Tx) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `sonarqube_metric` SET `checksum`=?,`customer_id`=?,`project_ext_id`=?,`project_id`=?,`project_key_ext_id`=?,`ext_id`=?,`date`=?,`metric`=?,`value`=? WHERE `id`=?"
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.ProjectExtID),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLString(t.ProjectKeyExtID),
		orm.ToSQLString(t.ExtID),
		orm.ToSQLInt64(t.Date),
		orm.ToSQLString(t.Metric),
		orm.ToSQLString(t.Value),
		orm.ToSQLString(t.ID),
	)
}

// DBUpsert will upsert the SonarqubeMetric record in the database
func (t *SonarqubeMetric) DBUpsert(ctx context.Context, db *sql.DB, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `sonarqube_metric` (`sonarqube_metric`.`id`,`sonarqube_metric`.`checksum`,`sonarqube_metric`.`customer_id`,`sonarqube_metric`.`project_ext_id`,`sonarqube_metric`.`project_id`,`sonarqube_metric`.`project_key_ext_id`,`sonarqube_metric`.`ext_id`,`sonarqube_metric`.`date`,`sonarqube_metric`.`metric`,`sonarqube_metric`.`value`) VALUES (?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `sonarqube_metric` (`sonarqube_metric`.`id`,`sonarqube_metric`.`checksum`,`sonarqube_metric`.`customer_id`,`sonarqube_metric`.`project_ext_id`,`sonarqube_metric`.`project_id`,`sonarqube_metric`.`project_key_ext_id`,`sonarqube_metric`.`ext_id`,`sonarqube_metric`.`date`,`sonarqube_metric`.`metric`,`sonarqube_metric`.`value`) VALUES (?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`customer_id`=VALUES(`customer_id`),`project_ext_id`=VALUES(`project_ext_id`),`project_id`=VALUES(`project_id`),`project_key_ext_id`=VALUES(`project_key_ext_id`),`ext_id`=VALUES(`ext_id`),`date`=VALUES(`date`),`metric`=VALUES(`metric`),`value`=VALUES(`value`)"
	}
	r, err := db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.ProjectExtID),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLString(t.ProjectKeyExtID),
		orm.ToSQLString(t.ExtID),
		orm.ToSQLInt64(t.Date),
		orm.ToSQLString(t.Metric),
		orm.ToSQLString(t.Value),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBUpsertTx will upsert the SonarqubeMetric record in the database using the provided transaction
func (t *SonarqubeMetric) DBUpsertTx(ctx context.Context, tx *sql.Tx, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `sonarqube_metric` (`sonarqube_metric`.`id`,`sonarqube_metric`.`checksum`,`sonarqube_metric`.`customer_id`,`sonarqube_metric`.`project_ext_id`,`sonarqube_metric`.`project_id`,`sonarqube_metric`.`project_key_ext_id`,`sonarqube_metric`.`ext_id`,`sonarqube_metric`.`date`,`sonarqube_metric`.`metric`,`sonarqube_metric`.`value`) VALUES (?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `sonarqube_metric` (`sonarqube_metric`.`id`,`sonarqube_metric`.`checksum`,`sonarqube_metric`.`customer_id`,`sonarqube_metric`.`project_ext_id`,`sonarqube_metric`.`project_id`,`sonarqube_metric`.`project_key_ext_id`,`sonarqube_metric`.`ext_id`,`sonarqube_metric`.`date`,`sonarqube_metric`.`metric`,`sonarqube_metric`.`value`) VALUES (?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`customer_id`=VALUES(`customer_id`),`project_ext_id`=VALUES(`project_ext_id`),`project_id`=VALUES(`project_id`),`project_key_ext_id`=VALUES(`project_key_ext_id`),`ext_id`=VALUES(`ext_id`),`date`=VALUES(`date`),`metric`=VALUES(`metric`),`value`=VALUES(`value`)"
	}
	r, err := tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.ProjectExtID),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLString(t.ProjectKeyExtID),
		orm.ToSQLString(t.ExtID),
		orm.ToSQLInt64(t.Date),
		orm.ToSQLString(t.Metric),
		orm.ToSQLString(t.Value),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBFindOne will find a SonarqubeMetric record in the database with the primary key
func (t *SonarqubeMetric) DBFindOne(ctx context.Context, db *sql.DB, value string) (bool, error) {
	q := "SELECT `sonarqube_metric`.`id`,`sonarqube_metric`.`checksum`,`sonarqube_metric`.`customer_id`,`sonarqube_metric`.`project_ext_id`,`sonarqube_metric`.`project_id`,`sonarqube_metric`.`project_key_ext_id`,`sonarqube_metric`.`ext_id`,`sonarqube_metric`.`date`,`sonarqube_metric`.`metric`,`sonarqube_metric`.`value` FROM `sonarqube_metric` WHERE `id` = ? LIMIT 1"
	row := db.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CustomerID sql.NullString
	var _ProjectExtID sql.NullString
	var _ProjectID sql.NullString
	var _ProjectKeyExtID sql.NullString
	var _ExtID sql.NullString
	var _Date sql.NullInt64
	var _Metric sql.NullString
	var _Value sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_ProjectExtID,
		&_ProjectID,
		&_ProjectKeyExtID,
		&_ExtID,
		&_Date,
		&_Metric,
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
	if _ProjectExtID.Valid {
		t.SetProjectExtID(_ProjectExtID.String)
	}
	if _ProjectID.Valid {
		t.SetProjectID(_ProjectID.String)
	}
	if _ProjectKeyExtID.Valid {
		t.SetProjectKeyExtID(_ProjectKeyExtID.String)
	}
	if _ExtID.Valid {
		t.SetExtID(_ExtID.String)
	}
	if _Date.Valid {
		t.SetDate(_Date.Int64)
	}
	if _Metric.Valid {
		t.SetMetric(_Metric.String)
	}
	if _Value.Valid {
		t.SetValue(_Value.String)
	}
	return true, nil
}

// DBFindOneTx will find a SonarqubeMetric record in the database with the primary key using the provided transaction
func (t *SonarqubeMetric) DBFindOneTx(ctx context.Context, tx *sql.Tx, value string) (bool, error) {
	q := "SELECT `sonarqube_metric`.`id`,`sonarqube_metric`.`checksum`,`sonarqube_metric`.`customer_id`,`sonarqube_metric`.`project_ext_id`,`sonarqube_metric`.`project_id`,`sonarqube_metric`.`project_key_ext_id`,`sonarqube_metric`.`ext_id`,`sonarqube_metric`.`date`,`sonarqube_metric`.`metric`,`sonarqube_metric`.`value` FROM `sonarqube_metric` WHERE `id` = ? LIMIT 1"
	row := tx.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CustomerID sql.NullString
	var _ProjectExtID sql.NullString
	var _ProjectID sql.NullString
	var _ProjectKeyExtID sql.NullString
	var _ExtID sql.NullString
	var _Date sql.NullInt64
	var _Metric sql.NullString
	var _Value sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_ProjectExtID,
		&_ProjectID,
		&_ProjectKeyExtID,
		&_ExtID,
		&_Date,
		&_Metric,
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
	if _ProjectExtID.Valid {
		t.SetProjectExtID(_ProjectExtID.String)
	}
	if _ProjectID.Valid {
		t.SetProjectID(_ProjectID.String)
	}
	if _ProjectKeyExtID.Valid {
		t.SetProjectKeyExtID(_ProjectKeyExtID.String)
	}
	if _ExtID.Valid {
		t.SetExtID(_ExtID.String)
	}
	if _Date.Valid {
		t.SetDate(_Date.Int64)
	}
	if _Metric.Valid {
		t.SetMetric(_Metric.String)
	}
	if _Value.Valid {
		t.SetValue(_Value.String)
	}
	return true, nil
}

// FindSonarqubeMetrics will find a SonarqubeMetric record in the database with the provided parameters
func FindSonarqubeMetrics(ctx context.Context, db *sql.DB, _params ...interface{}) ([]*SonarqubeMetric, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("customer_id"),
		orm.Column("project_ext_id"),
		orm.Column("project_id"),
		orm.Column("project_key_ext_id"),
		orm.Column("ext_id"),
		orm.Column("date"),
		orm.Column("metric"),
		orm.Column("value"),
		orm.Table(SonarqubeMetricTableName),
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
	results := make([]*SonarqubeMetric, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _ProjectExtID sql.NullString
		var _ProjectID sql.NullString
		var _ProjectKeyExtID sql.NullString
		var _ExtID sql.NullString
		var _Date sql.NullInt64
		var _Metric sql.NullString
		var _Value sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_ProjectExtID,
			&_ProjectID,
			&_ProjectKeyExtID,
			&_ExtID,
			&_Date,
			&_Metric,
			&_Value,
		)
		if err != nil {
			return nil, err
		}
		t := &SonarqubeMetric{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _ProjectExtID.Valid {
			t.SetProjectExtID(_ProjectExtID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _ProjectKeyExtID.Valid {
			t.SetProjectKeyExtID(_ProjectKeyExtID.String)
		}
		if _ExtID.Valid {
			t.SetExtID(_ExtID.String)
		}
		if _Date.Valid {
			t.SetDate(_Date.Int64)
		}
		if _Metric.Valid {
			t.SetMetric(_Metric.String)
		}
		if _Value.Valid {
			t.SetValue(_Value.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindSonarqubeMetricsTx will find a SonarqubeMetric record in the database with the provided parameters using the provided transaction
func FindSonarqubeMetricsTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) ([]*SonarqubeMetric, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("customer_id"),
		orm.Column("project_ext_id"),
		orm.Column("project_id"),
		orm.Column("project_key_ext_id"),
		orm.Column("ext_id"),
		orm.Column("date"),
		orm.Column("metric"),
		orm.Column("value"),
		orm.Table(SonarqubeMetricTableName),
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
	results := make([]*SonarqubeMetric, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _ProjectExtID sql.NullString
		var _ProjectID sql.NullString
		var _ProjectKeyExtID sql.NullString
		var _ExtID sql.NullString
		var _Date sql.NullInt64
		var _Metric sql.NullString
		var _Value sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_ProjectExtID,
			&_ProjectID,
			&_ProjectKeyExtID,
			&_ExtID,
			&_Date,
			&_Metric,
			&_Value,
		)
		if err != nil {
			return nil, err
		}
		t := &SonarqubeMetric{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _ProjectExtID.Valid {
			t.SetProjectExtID(_ProjectExtID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _ProjectKeyExtID.Valid {
			t.SetProjectKeyExtID(_ProjectKeyExtID.String)
		}
		if _ExtID.Valid {
			t.SetExtID(_ExtID.String)
		}
		if _Date.Valid {
			t.SetDate(_Date.Int64)
		}
		if _Metric.Valid {
			t.SetMetric(_Metric.String)
		}
		if _Value.Valid {
			t.SetValue(_Value.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// DBFind will find a SonarqubeMetric record in the database with the provided parameters
func (t *SonarqubeMetric) DBFind(ctx context.Context, db *sql.DB, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("customer_id"),
		orm.Column("project_ext_id"),
		orm.Column("project_id"),
		orm.Column("project_key_ext_id"),
		orm.Column("ext_id"),
		orm.Column("date"),
		orm.Column("metric"),
		orm.Column("value"),
		orm.Table(SonarqubeMetricTableName),
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
	var _ProjectExtID sql.NullString
	var _ProjectID sql.NullString
	var _ProjectKeyExtID sql.NullString
	var _ExtID sql.NullString
	var _Date sql.NullInt64
	var _Metric sql.NullString
	var _Value sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_ProjectExtID,
		&_ProjectID,
		&_ProjectKeyExtID,
		&_ExtID,
		&_Date,
		&_Metric,
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
	if _ProjectExtID.Valid {
		t.SetProjectExtID(_ProjectExtID.String)
	}
	if _ProjectID.Valid {
		t.SetProjectID(_ProjectID.String)
	}
	if _ProjectKeyExtID.Valid {
		t.SetProjectKeyExtID(_ProjectKeyExtID.String)
	}
	if _ExtID.Valid {
		t.SetExtID(_ExtID.String)
	}
	if _Date.Valid {
		t.SetDate(_Date.Int64)
	}
	if _Metric.Valid {
		t.SetMetric(_Metric.String)
	}
	if _Value.Valid {
		t.SetValue(_Value.String)
	}
	return true, nil
}

// DBFindTx will find a SonarqubeMetric record in the database with the provided parameters using the provided transaction
func (t *SonarqubeMetric) DBFindTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("customer_id"),
		orm.Column("project_ext_id"),
		orm.Column("project_id"),
		orm.Column("project_key_ext_id"),
		orm.Column("ext_id"),
		orm.Column("date"),
		orm.Column("metric"),
		orm.Column("value"),
		orm.Table(SonarqubeMetricTableName),
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
	var _ProjectExtID sql.NullString
	var _ProjectID sql.NullString
	var _ProjectKeyExtID sql.NullString
	var _ExtID sql.NullString
	var _Date sql.NullInt64
	var _Metric sql.NullString
	var _Value sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_ProjectExtID,
		&_ProjectID,
		&_ProjectKeyExtID,
		&_ExtID,
		&_Date,
		&_Metric,
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
	if _ProjectExtID.Valid {
		t.SetProjectExtID(_ProjectExtID.String)
	}
	if _ProjectID.Valid {
		t.SetProjectID(_ProjectID.String)
	}
	if _ProjectKeyExtID.Valid {
		t.SetProjectKeyExtID(_ProjectKeyExtID.String)
	}
	if _ExtID.Valid {
		t.SetExtID(_ExtID.String)
	}
	if _Date.Valid {
		t.SetDate(_Date.Int64)
	}
	if _Metric.Valid {
		t.SetMetric(_Metric.String)
	}
	if _Value.Valid {
		t.SetValue(_Value.String)
	}
	return true, nil
}

// CountSonarqubeMetrics will find the count of SonarqubeMetric records in the database
func CountSonarqubeMetrics(ctx context.Context, db *sql.DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(SonarqubeMetricTableName),
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

// CountSonarqubeMetricsTx will find the count of SonarqubeMetric records in the database using the provided transaction
func CountSonarqubeMetricsTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(SonarqubeMetricTableName),
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

// DBCount will find the count of SonarqubeMetric records in the database
func (t *SonarqubeMetric) DBCount(ctx context.Context, db *sql.DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(SonarqubeMetricTableName),
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

// DBCountTx will find the count of SonarqubeMetric records in the database using the provided transaction
func (t *SonarqubeMetric) DBCountTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(SonarqubeMetricTableName),
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

// DBExists will return true if the SonarqubeMetric record exists in the database
func (t *SonarqubeMetric) DBExists(ctx context.Context, db *sql.DB) (bool, error) {
	q := "SELECT `id` FROM `sonarqube_metric` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := db.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// DBExistsTx will return true if the SonarqubeMetric record exists in the database using the provided transaction
func (t *SonarqubeMetric) DBExistsTx(ctx context.Context, tx *sql.Tx) (bool, error) {
	q := "SELECT `id` FROM `sonarqube_metric` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := tx.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// PrimaryKeyColumn returns the column name for the primary key
func (t *SonarqubeMetric) PrimaryKeyColumn() string {
	return SonarqubeMetricColumnID
}

// PrimaryKeyColumnType returns the primary key column Go type as a string
func (t *SonarqubeMetric) PrimaryKeyColumnType() string {
	return "string"
}

// PrimaryKey returns the primary key column value
func (t *SonarqubeMetric) PrimaryKey() interface{} {
	return t.ID
}
