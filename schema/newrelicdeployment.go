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

var _ Model = (*NewrelicDeployment)(nil)
var _ CSVWriter = (*NewrelicDeployment)(nil)
var _ JSONWriter = (*NewrelicDeployment)(nil)
var _ Checksum = (*NewrelicDeployment)(nil)

// NewrelicDeploymentTableName is the name of the table in SQL
const NewrelicDeploymentTableName = "newrelic_deployment"

var NewrelicDeploymentColumns = []string{
	"id",
	"checksum",
	"customer_id",
	"application_ext_id",
	"application_id",
	"ext_id",
	"revision",
	"changelog",
	"description",
	"user",
	"timestamp",
}

// NewrelicDeployment table
type NewrelicDeployment struct {
	ApplicationExtID int64   `json:"application_ext_id"`
	ApplicationID    string  `json:"application_id"`
	Changelog        *string `json:"changelog,omitempty"`
	Checksum         *string `json:"checksum,omitempty"`
	CustomerID       string  `json:"customer_id"`
	Description      *string `json:"description,omitempty"`
	ExtID            int64   `json:"ext_id"`
	ID               string  `json:"id"`
	Revision         string  `json:"revision"`
	Timestamp        int64   `json:"timestamp"`
	User             string  `json:"user"`
}

// TableName returns the SQL table name for NewrelicDeployment and satifies the Model interface
func (t *NewrelicDeployment) TableName() string {
	return NewrelicDeploymentTableName
}

// ToCSV will serialize the NewrelicDeployment instance to a CSV compatible array of strings
func (t *NewrelicDeployment) ToCSV() []string {
	return []string{
		t.ID,
		t.CalculateChecksum(),
		t.CustomerID,
		toCSVString(t.ApplicationExtID),
		t.ApplicationID,
		toCSVString(t.ExtID),
		t.Revision,
		toCSVString(t.Changelog),
		toCSVString(t.Description),
		t.User,
		toCSVString(t.Timestamp),
	}
}

// WriteCSV will serialize the NewrelicDeployment instance to the writer as CSV and satisfies the CSVWriter interface
func (t *NewrelicDeployment) WriteCSV(w *csv.Writer) error {
	return w.Write(t.ToCSV())
}

// WriteJSON will serialize the NewrelicDeployment instance to the writer as JSON and satisfies the JSONWriter interface
func (t *NewrelicDeployment) WriteJSON(w io.Writer, indent ...bool) error {
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

// NewNewrelicDeploymentReader creates a JSON reader which can read in NewrelicDeployment objects serialized as JSON either as an array, single object or json new lines
// and writes each NewrelicDeployment to the channel provided
func NewNewrelicDeploymentReader(r io.Reader, ch chan<- NewrelicDeployment) error {
	return orm.Deserialize(r, func(buf json.RawMessage) error {
		dec := json.NewDecoder(bytes.NewBuffer(buf))
		e := NewrelicDeployment{}
		if err := dec.Decode(&e); err != nil {
			return err
		}
		ch <- e
		return nil
	})
}

// NewCSVNewrelicDeploymentReaderDir will read the reader as CSV and emit each record to the channel provided
func NewCSVNewrelicDeploymentReader(r io.Reader, ch chan<- NewrelicDeployment) error {
	cr := csv.NewReader(r)
	for {
		record, err := cr.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		ch <- NewrelicDeployment{
			ID:               record[0],
			Checksum:         fromStringPointer(record[1]),
			CustomerID:       record[2],
			ApplicationExtID: fromCSVInt64(record[3]),
			ApplicationID:    record[4],
			ExtID:            fromCSVInt64(record[5]),
			Revision:         record[6],
			Changelog:        fromStringPointer(record[7]),
			Description:      fromStringPointer(record[8]),
			User:             record[9],
			Timestamp:        fromCSVInt64(record[10]),
		}
	}
	return nil
}

// NewCSVNewrelicDeploymentReaderFile will read the file as a CSV and emit each record to the channel provided
func NewCSVNewrelicDeploymentReaderFile(fp string, ch chan<- NewrelicDeployment) error {
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
	return NewCSVNewrelicDeploymentReader(fc, ch)
}

// NewCSVNewrelicDeploymentReaderDir will read the newrelic_deployment.csv.gz file as a CSV and emit each record to the channel provided
func NewCSVNewrelicDeploymentReaderDir(dir string, ch chan<- NewrelicDeployment) error {
	return NewCSVNewrelicDeploymentReaderFile(filepath.Join(dir, "newrelic_deployment.csv.gz"), ch)
}

// NewrelicDeploymentCSVDeduper is a function callback which takes the existing value (a) and the new value (b)
// and the return value should be the one to use (or a new one, if applicable). return nil
// to skip processing of this record
type NewrelicDeploymentCSVDeduper func(a NewrelicDeployment, b NewrelicDeployment) *NewrelicDeployment

// NewrelicDeploymentCSVDedupeDisabled is set on whether the CSV writer should de-dupe values by key
var NewrelicDeploymentCSVDedupeDisabled bool

// NewNewrelicDeploymentCSVWriterSize creates a batch writer that will write each NewrelicDeployment into a CSV file
// this method will automatically de-duplicate entries using the primary key. if the checksum
// for a newer item with the same primary key doesn't match a previously sent item, the newer
// one will be used
func NewNewrelicDeploymentCSVWriterSize(w io.Writer, size int, dedupers ...NewrelicDeploymentCSVDeduper) (chan NewrelicDeployment, chan bool, error) {
	cw := csv.NewWriter(w)
	ch := make(chan NewrelicDeployment, size)
	done := make(chan bool)
	go func() {
		defer func() { done <- true }()
		dodedupe := !NewrelicDeploymentCSVDedupeDisabled
		var kv map[string]*NewrelicDeployment
		var deduper NewrelicDeploymentCSVDeduper
		if dedupers != nil && len(dedupers) > 0 {
			deduper = dedupers[0]
			dodedupe = true
		}
		if dodedupe {
			kv = make(map[string]*NewrelicDeployment)
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

// NewrelicDeploymentCSVDefaultSize is the default channel buffer size if not provided
var NewrelicDeploymentCSVDefaultSize = 100

// NewNewrelicDeploymentCSVWriter creates a batch writer that will write each NewrelicDeployment into a CSV file
func NewNewrelicDeploymentCSVWriter(w io.Writer, dedupers ...NewrelicDeploymentCSVDeduper) (chan NewrelicDeployment, chan bool, error) {
	return NewNewrelicDeploymentCSVWriterSize(w, NewrelicDeploymentCSVDefaultSize, dedupers...)
}

// NewNewrelicDeploymentCSVWriterDir creates a batch writer that will write each NewrelicDeployment into a CSV file named newrelic_deployment.csv.gz in dir
func NewNewrelicDeploymentCSVWriterDir(dir string, dedupers ...NewrelicDeploymentCSVDeduper) (chan NewrelicDeployment, chan bool, error) {
	return NewNewrelicDeploymentCSVWriterFile(filepath.Join(dir, "newrelic_deployment.csv.gz"), dedupers...)
}

// NewNewrelicDeploymentCSVWriterFile creates a batch writer that will write each NewrelicDeployment into a CSV file
func NewNewrelicDeploymentCSVWriterFile(fn string, dedupers ...NewrelicDeploymentCSVDeduper) (chan NewrelicDeployment, chan bool, error) {
	f, err := os.Create(fn)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening CSV file %s. %v", fn, err)
	}
	var fc io.WriteCloser = f
	if filepath.Ext(fn) == ".gz" {
		w, _ := gzip.NewWriterLevel(f, gzip.BestCompression)
		fc = w
	}
	ch, done, err := NewNewrelicDeploymentCSVWriter(fc, dedupers...)
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

type NewrelicDeploymentDBAction func(ctx context.Context, db *sql.DB, record NewrelicDeployment) error

// NewNewrelicDeploymentDBWriterSize creates a DB writer that will write each issue into the DB
func NewNewrelicDeploymentDBWriterSize(ctx context.Context, db *sql.DB, errors chan<- error, size int, actions ...NewrelicDeploymentDBAction) (chan NewrelicDeployment, chan bool, error) {
	ch := make(chan NewrelicDeployment, size)
	done := make(chan bool)
	var action NewrelicDeploymentDBAction
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

// NewNewrelicDeploymentDBWriter creates a DB writer that will write each issue into the DB
func NewNewrelicDeploymentDBWriter(ctx context.Context, db *sql.DB, errors chan<- error, actions ...NewrelicDeploymentDBAction) (chan NewrelicDeployment, chan bool, error) {
	return NewNewrelicDeploymentDBWriterSize(ctx, db, errors, 100, actions...)
}

// NewrelicDeploymentColumnID is the ID SQL column name for the NewrelicDeployment table
const NewrelicDeploymentColumnID = "id"

// NewrelicDeploymentEscapedColumnID is the escaped ID SQL column name for the NewrelicDeployment table
const NewrelicDeploymentEscapedColumnID = "`id`"

// NewrelicDeploymentColumnChecksum is the Checksum SQL column name for the NewrelicDeployment table
const NewrelicDeploymentColumnChecksum = "checksum"

// NewrelicDeploymentEscapedColumnChecksum is the escaped Checksum SQL column name for the NewrelicDeployment table
const NewrelicDeploymentEscapedColumnChecksum = "`checksum`"

// NewrelicDeploymentColumnCustomerID is the CustomerID SQL column name for the NewrelicDeployment table
const NewrelicDeploymentColumnCustomerID = "customer_id"

// NewrelicDeploymentEscapedColumnCustomerID is the escaped CustomerID SQL column name for the NewrelicDeployment table
const NewrelicDeploymentEscapedColumnCustomerID = "`customer_id`"

// NewrelicDeploymentColumnApplicationExtID is the ApplicationExtID SQL column name for the NewrelicDeployment table
const NewrelicDeploymentColumnApplicationExtID = "application_ext_id"

// NewrelicDeploymentEscapedColumnApplicationExtID is the escaped ApplicationExtID SQL column name for the NewrelicDeployment table
const NewrelicDeploymentEscapedColumnApplicationExtID = "`application_ext_id`"

// NewrelicDeploymentColumnApplicationID is the ApplicationID SQL column name for the NewrelicDeployment table
const NewrelicDeploymentColumnApplicationID = "application_id"

// NewrelicDeploymentEscapedColumnApplicationID is the escaped ApplicationID SQL column name for the NewrelicDeployment table
const NewrelicDeploymentEscapedColumnApplicationID = "`application_id`"

// NewrelicDeploymentColumnExtID is the ExtID SQL column name for the NewrelicDeployment table
const NewrelicDeploymentColumnExtID = "ext_id"

// NewrelicDeploymentEscapedColumnExtID is the escaped ExtID SQL column name for the NewrelicDeployment table
const NewrelicDeploymentEscapedColumnExtID = "`ext_id`"

// NewrelicDeploymentColumnRevision is the Revision SQL column name for the NewrelicDeployment table
const NewrelicDeploymentColumnRevision = "revision"

// NewrelicDeploymentEscapedColumnRevision is the escaped Revision SQL column name for the NewrelicDeployment table
const NewrelicDeploymentEscapedColumnRevision = "`revision`"

// NewrelicDeploymentColumnChangelog is the Changelog SQL column name for the NewrelicDeployment table
const NewrelicDeploymentColumnChangelog = "changelog"

// NewrelicDeploymentEscapedColumnChangelog is the escaped Changelog SQL column name for the NewrelicDeployment table
const NewrelicDeploymentEscapedColumnChangelog = "`changelog`"

// NewrelicDeploymentColumnDescription is the Description SQL column name for the NewrelicDeployment table
const NewrelicDeploymentColumnDescription = "description"

// NewrelicDeploymentEscapedColumnDescription is the escaped Description SQL column name for the NewrelicDeployment table
const NewrelicDeploymentEscapedColumnDescription = "`description`"

// NewrelicDeploymentColumnUser is the User SQL column name for the NewrelicDeployment table
const NewrelicDeploymentColumnUser = "user"

// NewrelicDeploymentEscapedColumnUser is the escaped User SQL column name for the NewrelicDeployment table
const NewrelicDeploymentEscapedColumnUser = "`user`"

// NewrelicDeploymentColumnTimestamp is the Timestamp SQL column name for the NewrelicDeployment table
const NewrelicDeploymentColumnTimestamp = "timestamp"

// NewrelicDeploymentEscapedColumnTimestamp is the escaped Timestamp SQL column name for the NewrelicDeployment table
const NewrelicDeploymentEscapedColumnTimestamp = "`timestamp`"

// GetID will return the NewrelicDeployment ID value
func (t *NewrelicDeployment) GetID() string {
	return t.ID
}

// SetID will set the NewrelicDeployment ID value
func (t *NewrelicDeployment) SetID(v string) {
	t.ID = v
}

// FindNewrelicDeploymentByID will find a NewrelicDeployment by ID
func FindNewrelicDeploymentByID(ctx context.Context, db *sql.DB, value string) (*NewrelicDeployment, error) {
	q := "SELECT `newrelic_deployment`.`id`,`newrelic_deployment`.`checksum`,`newrelic_deployment`.`customer_id`,`newrelic_deployment`.`application_ext_id`,`newrelic_deployment`.`application_id`,`newrelic_deployment`.`ext_id`,`newrelic_deployment`.`revision`,`newrelic_deployment`.`changelog`,`newrelic_deployment`.`description`,`newrelic_deployment`.`user`,`newrelic_deployment`.`timestamp` FROM `newrelic_deployment` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CustomerID sql.NullString
	var _ApplicationExtID sql.NullInt64
	var _ApplicationID sql.NullString
	var _ExtID sql.NullInt64
	var _Revision sql.NullString
	var _Changelog sql.NullString
	var _Description sql.NullString
	var _User sql.NullString
	var _Timestamp sql.NullInt64
	err := db.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_ApplicationExtID,
		&_ApplicationID,
		&_ExtID,
		&_Revision,
		&_Changelog,
		&_Description,
		&_User,
		&_Timestamp,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &NewrelicDeployment{}
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
	if _ExtID.Valid {
		t.SetExtID(_ExtID.Int64)
	}
	if _Revision.Valid {
		t.SetRevision(_Revision.String)
	}
	if _Changelog.Valid {
		t.SetChangelog(_Changelog.String)
	}
	if _Description.Valid {
		t.SetDescription(_Description.String)
	}
	if _User.Valid {
		t.SetUser(_User.String)
	}
	if _Timestamp.Valid {
		t.SetTimestamp(_Timestamp.Int64)
	}
	return t, nil
}

// FindNewrelicDeploymentByIDTx will find a NewrelicDeployment by ID using the provided transaction
func FindNewrelicDeploymentByIDTx(ctx context.Context, tx *sql.Tx, value string) (*NewrelicDeployment, error) {
	q := "SELECT `newrelic_deployment`.`id`,`newrelic_deployment`.`checksum`,`newrelic_deployment`.`customer_id`,`newrelic_deployment`.`application_ext_id`,`newrelic_deployment`.`application_id`,`newrelic_deployment`.`ext_id`,`newrelic_deployment`.`revision`,`newrelic_deployment`.`changelog`,`newrelic_deployment`.`description`,`newrelic_deployment`.`user`,`newrelic_deployment`.`timestamp` FROM `newrelic_deployment` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CustomerID sql.NullString
	var _ApplicationExtID sql.NullInt64
	var _ApplicationID sql.NullString
	var _ExtID sql.NullInt64
	var _Revision sql.NullString
	var _Changelog sql.NullString
	var _Description sql.NullString
	var _User sql.NullString
	var _Timestamp sql.NullInt64
	err := tx.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_ApplicationExtID,
		&_ApplicationID,
		&_ExtID,
		&_Revision,
		&_Changelog,
		&_Description,
		&_User,
		&_Timestamp,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &NewrelicDeployment{}
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
	if _ExtID.Valid {
		t.SetExtID(_ExtID.Int64)
	}
	if _Revision.Valid {
		t.SetRevision(_Revision.String)
	}
	if _Changelog.Valid {
		t.SetChangelog(_Changelog.String)
	}
	if _Description.Valid {
		t.SetDescription(_Description.String)
	}
	if _User.Valid {
		t.SetUser(_User.String)
	}
	if _Timestamp.Valid {
		t.SetTimestamp(_Timestamp.Int64)
	}
	return t, nil
}

// GetChecksum will return the NewrelicDeployment Checksum value
func (t *NewrelicDeployment) GetChecksum() string {
	if t.Checksum == nil {
		return ""
	}
	return *t.Checksum
}

// SetChecksum will set the NewrelicDeployment Checksum value
func (t *NewrelicDeployment) SetChecksum(v string) {
	t.Checksum = &v
}

// GetCustomerID will return the NewrelicDeployment CustomerID value
func (t *NewrelicDeployment) GetCustomerID() string {
	return t.CustomerID
}

// SetCustomerID will set the NewrelicDeployment CustomerID value
func (t *NewrelicDeployment) SetCustomerID(v string) {
	t.CustomerID = v
}

// FindNewrelicDeploymentsByCustomerID will find all NewrelicDeployments by the CustomerID value
func FindNewrelicDeploymentsByCustomerID(ctx context.Context, db *sql.DB, value string) ([]*NewrelicDeployment, error) {
	q := "SELECT `newrelic_deployment`.`id`,`newrelic_deployment`.`checksum`,`newrelic_deployment`.`customer_id`,`newrelic_deployment`.`application_ext_id`,`newrelic_deployment`.`application_id`,`newrelic_deployment`.`ext_id`,`newrelic_deployment`.`revision`,`newrelic_deployment`.`changelog`,`newrelic_deployment`.`description`,`newrelic_deployment`.`user`,`newrelic_deployment`.`timestamp` FROM `newrelic_deployment` WHERE `customer_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*NewrelicDeployment, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _ApplicationExtID sql.NullInt64
		var _ApplicationID sql.NullString
		var _ExtID sql.NullInt64
		var _Revision sql.NullString
		var _Changelog sql.NullString
		var _Description sql.NullString
		var _User sql.NullString
		var _Timestamp sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_ApplicationExtID,
			&_ApplicationID,
			&_ExtID,
			&_Revision,
			&_Changelog,
			&_Description,
			&_User,
			&_Timestamp,
		)
		if err != nil {
			return nil, err
		}
		t := &NewrelicDeployment{}
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
		if _ExtID.Valid {
			t.SetExtID(_ExtID.Int64)
		}
		if _Revision.Valid {
			t.SetRevision(_Revision.String)
		}
		if _Changelog.Valid {
			t.SetChangelog(_Changelog.String)
		}
		if _Description.Valid {
			t.SetDescription(_Description.String)
		}
		if _User.Valid {
			t.SetUser(_User.String)
		}
		if _Timestamp.Valid {
			t.SetTimestamp(_Timestamp.Int64)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindNewrelicDeploymentsByCustomerIDTx will find all NewrelicDeployments by the CustomerID value using the provided transaction
func FindNewrelicDeploymentsByCustomerIDTx(ctx context.Context, tx *sql.Tx, value string) ([]*NewrelicDeployment, error) {
	q := "SELECT `newrelic_deployment`.`id`,`newrelic_deployment`.`checksum`,`newrelic_deployment`.`customer_id`,`newrelic_deployment`.`application_ext_id`,`newrelic_deployment`.`application_id`,`newrelic_deployment`.`ext_id`,`newrelic_deployment`.`revision`,`newrelic_deployment`.`changelog`,`newrelic_deployment`.`description`,`newrelic_deployment`.`user`,`newrelic_deployment`.`timestamp` FROM `newrelic_deployment` WHERE `customer_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*NewrelicDeployment, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _ApplicationExtID sql.NullInt64
		var _ApplicationID sql.NullString
		var _ExtID sql.NullInt64
		var _Revision sql.NullString
		var _Changelog sql.NullString
		var _Description sql.NullString
		var _User sql.NullString
		var _Timestamp sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_ApplicationExtID,
			&_ApplicationID,
			&_ExtID,
			&_Revision,
			&_Changelog,
			&_Description,
			&_User,
			&_Timestamp,
		)
		if err != nil {
			return nil, err
		}
		t := &NewrelicDeployment{}
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
		if _ExtID.Valid {
			t.SetExtID(_ExtID.Int64)
		}
		if _Revision.Valid {
			t.SetRevision(_Revision.String)
		}
		if _Changelog.Valid {
			t.SetChangelog(_Changelog.String)
		}
		if _Description.Valid {
			t.SetDescription(_Description.String)
		}
		if _User.Valid {
			t.SetUser(_User.String)
		}
		if _Timestamp.Valid {
			t.SetTimestamp(_Timestamp.Int64)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetApplicationExtID will return the NewrelicDeployment ApplicationExtID value
func (t *NewrelicDeployment) GetApplicationExtID() int64 {
	return t.ApplicationExtID
}

// SetApplicationExtID will set the NewrelicDeployment ApplicationExtID value
func (t *NewrelicDeployment) SetApplicationExtID(v int64) {
	t.ApplicationExtID = v
}

// FindNewrelicDeploymentsByApplicationExtID will find all NewrelicDeployments by the ApplicationExtID value
func FindNewrelicDeploymentsByApplicationExtID(ctx context.Context, db *sql.DB, value int64) ([]*NewrelicDeployment, error) {
	q := "SELECT `newrelic_deployment`.`id`,`newrelic_deployment`.`checksum`,`newrelic_deployment`.`customer_id`,`newrelic_deployment`.`application_ext_id`,`newrelic_deployment`.`application_id`,`newrelic_deployment`.`ext_id`,`newrelic_deployment`.`revision`,`newrelic_deployment`.`changelog`,`newrelic_deployment`.`description`,`newrelic_deployment`.`user`,`newrelic_deployment`.`timestamp` FROM `newrelic_deployment` WHERE `application_ext_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*NewrelicDeployment, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _ApplicationExtID sql.NullInt64
		var _ApplicationID sql.NullString
		var _ExtID sql.NullInt64
		var _Revision sql.NullString
		var _Changelog sql.NullString
		var _Description sql.NullString
		var _User sql.NullString
		var _Timestamp sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_ApplicationExtID,
			&_ApplicationID,
			&_ExtID,
			&_Revision,
			&_Changelog,
			&_Description,
			&_User,
			&_Timestamp,
		)
		if err != nil {
			return nil, err
		}
		t := &NewrelicDeployment{}
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
		if _ExtID.Valid {
			t.SetExtID(_ExtID.Int64)
		}
		if _Revision.Valid {
			t.SetRevision(_Revision.String)
		}
		if _Changelog.Valid {
			t.SetChangelog(_Changelog.String)
		}
		if _Description.Valid {
			t.SetDescription(_Description.String)
		}
		if _User.Valid {
			t.SetUser(_User.String)
		}
		if _Timestamp.Valid {
			t.SetTimestamp(_Timestamp.Int64)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindNewrelicDeploymentsByApplicationExtIDTx will find all NewrelicDeployments by the ApplicationExtID value using the provided transaction
func FindNewrelicDeploymentsByApplicationExtIDTx(ctx context.Context, tx *sql.Tx, value int64) ([]*NewrelicDeployment, error) {
	q := "SELECT `newrelic_deployment`.`id`,`newrelic_deployment`.`checksum`,`newrelic_deployment`.`customer_id`,`newrelic_deployment`.`application_ext_id`,`newrelic_deployment`.`application_id`,`newrelic_deployment`.`ext_id`,`newrelic_deployment`.`revision`,`newrelic_deployment`.`changelog`,`newrelic_deployment`.`description`,`newrelic_deployment`.`user`,`newrelic_deployment`.`timestamp` FROM `newrelic_deployment` WHERE `application_ext_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*NewrelicDeployment, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _ApplicationExtID sql.NullInt64
		var _ApplicationID sql.NullString
		var _ExtID sql.NullInt64
		var _Revision sql.NullString
		var _Changelog sql.NullString
		var _Description sql.NullString
		var _User sql.NullString
		var _Timestamp sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_ApplicationExtID,
			&_ApplicationID,
			&_ExtID,
			&_Revision,
			&_Changelog,
			&_Description,
			&_User,
			&_Timestamp,
		)
		if err != nil {
			return nil, err
		}
		t := &NewrelicDeployment{}
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
		if _ExtID.Valid {
			t.SetExtID(_ExtID.Int64)
		}
		if _Revision.Valid {
			t.SetRevision(_Revision.String)
		}
		if _Changelog.Valid {
			t.SetChangelog(_Changelog.String)
		}
		if _Description.Valid {
			t.SetDescription(_Description.String)
		}
		if _User.Valid {
			t.SetUser(_User.String)
		}
		if _Timestamp.Valid {
			t.SetTimestamp(_Timestamp.Int64)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetApplicationID will return the NewrelicDeployment ApplicationID value
func (t *NewrelicDeployment) GetApplicationID() string {
	return t.ApplicationID
}

// SetApplicationID will set the NewrelicDeployment ApplicationID value
func (t *NewrelicDeployment) SetApplicationID(v string) {
	t.ApplicationID = v
}

// FindNewrelicDeploymentsByApplicationID will find all NewrelicDeployments by the ApplicationID value
func FindNewrelicDeploymentsByApplicationID(ctx context.Context, db *sql.DB, value string) ([]*NewrelicDeployment, error) {
	q := "SELECT `newrelic_deployment`.`id`,`newrelic_deployment`.`checksum`,`newrelic_deployment`.`customer_id`,`newrelic_deployment`.`application_ext_id`,`newrelic_deployment`.`application_id`,`newrelic_deployment`.`ext_id`,`newrelic_deployment`.`revision`,`newrelic_deployment`.`changelog`,`newrelic_deployment`.`description`,`newrelic_deployment`.`user`,`newrelic_deployment`.`timestamp` FROM `newrelic_deployment` WHERE `application_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*NewrelicDeployment, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _ApplicationExtID sql.NullInt64
		var _ApplicationID sql.NullString
		var _ExtID sql.NullInt64
		var _Revision sql.NullString
		var _Changelog sql.NullString
		var _Description sql.NullString
		var _User sql.NullString
		var _Timestamp sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_ApplicationExtID,
			&_ApplicationID,
			&_ExtID,
			&_Revision,
			&_Changelog,
			&_Description,
			&_User,
			&_Timestamp,
		)
		if err != nil {
			return nil, err
		}
		t := &NewrelicDeployment{}
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
		if _ExtID.Valid {
			t.SetExtID(_ExtID.Int64)
		}
		if _Revision.Valid {
			t.SetRevision(_Revision.String)
		}
		if _Changelog.Valid {
			t.SetChangelog(_Changelog.String)
		}
		if _Description.Valid {
			t.SetDescription(_Description.String)
		}
		if _User.Valid {
			t.SetUser(_User.String)
		}
		if _Timestamp.Valid {
			t.SetTimestamp(_Timestamp.Int64)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindNewrelicDeploymentsByApplicationIDTx will find all NewrelicDeployments by the ApplicationID value using the provided transaction
func FindNewrelicDeploymentsByApplicationIDTx(ctx context.Context, tx *sql.Tx, value string) ([]*NewrelicDeployment, error) {
	q := "SELECT `newrelic_deployment`.`id`,`newrelic_deployment`.`checksum`,`newrelic_deployment`.`customer_id`,`newrelic_deployment`.`application_ext_id`,`newrelic_deployment`.`application_id`,`newrelic_deployment`.`ext_id`,`newrelic_deployment`.`revision`,`newrelic_deployment`.`changelog`,`newrelic_deployment`.`description`,`newrelic_deployment`.`user`,`newrelic_deployment`.`timestamp` FROM `newrelic_deployment` WHERE `application_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*NewrelicDeployment, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _ApplicationExtID sql.NullInt64
		var _ApplicationID sql.NullString
		var _ExtID sql.NullInt64
		var _Revision sql.NullString
		var _Changelog sql.NullString
		var _Description sql.NullString
		var _User sql.NullString
		var _Timestamp sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_ApplicationExtID,
			&_ApplicationID,
			&_ExtID,
			&_Revision,
			&_Changelog,
			&_Description,
			&_User,
			&_Timestamp,
		)
		if err != nil {
			return nil, err
		}
		t := &NewrelicDeployment{}
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
		if _ExtID.Valid {
			t.SetExtID(_ExtID.Int64)
		}
		if _Revision.Valid {
			t.SetRevision(_Revision.String)
		}
		if _Changelog.Valid {
			t.SetChangelog(_Changelog.String)
		}
		if _Description.Valid {
			t.SetDescription(_Description.String)
		}
		if _User.Valid {
			t.SetUser(_User.String)
		}
		if _Timestamp.Valid {
			t.SetTimestamp(_Timestamp.Int64)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetExtID will return the NewrelicDeployment ExtID value
func (t *NewrelicDeployment) GetExtID() int64 {
	return t.ExtID
}

// SetExtID will set the NewrelicDeployment ExtID value
func (t *NewrelicDeployment) SetExtID(v int64) {
	t.ExtID = v
}

// GetRevision will return the NewrelicDeployment Revision value
func (t *NewrelicDeployment) GetRevision() string {
	return t.Revision
}

// SetRevision will set the NewrelicDeployment Revision value
func (t *NewrelicDeployment) SetRevision(v string) {
	t.Revision = v
}

// GetChangelog will return the NewrelicDeployment Changelog value
func (t *NewrelicDeployment) GetChangelog() string {
	if t.Changelog == nil {
		return ""
	}
	return *t.Changelog
}

// SetChangelog will set the NewrelicDeployment Changelog value
func (t *NewrelicDeployment) SetChangelog(v string) {
	t.Changelog = &v
}

// GetDescription will return the NewrelicDeployment Description value
func (t *NewrelicDeployment) GetDescription() string {
	if t.Description == nil {
		return ""
	}
	return *t.Description
}

// SetDescription will set the NewrelicDeployment Description value
func (t *NewrelicDeployment) SetDescription(v string) {
	t.Description = &v
}

// GetUser will return the NewrelicDeployment User value
func (t *NewrelicDeployment) GetUser() string {
	return t.User
}

// SetUser will set the NewrelicDeployment User value
func (t *NewrelicDeployment) SetUser(v string) {
	t.User = v
}

// GetTimestamp will return the NewrelicDeployment Timestamp value
func (t *NewrelicDeployment) GetTimestamp() int64 {
	return t.Timestamp
}

// SetTimestamp will set the NewrelicDeployment Timestamp value
func (t *NewrelicDeployment) SetTimestamp(v int64) {
	t.Timestamp = v
}

func (t *NewrelicDeployment) toTimestamp(value time.Time) *timestamp.Timestamp {
	ts, _ := ptypes.TimestampProto(value)
	return ts
}

// DBCreateNewrelicDeploymentTable will create the NewrelicDeployment table
func DBCreateNewrelicDeploymentTable(ctx context.Context, db *sql.DB) error {
	q := "CREATE TABLE `newrelic_deployment` (`id` VARCHAR(64) NOT NULL PRIMARY KEY,`checksum` CHAR(64),`customer_id` VARCHAR(64) NOT NULL,`application_ext_id` BIGINT(20) NOT NULL,`application_id` VARCHAR(64) NOT NULL,`ext_id` BIGINT(20) NOT NULL,`revision` TEXT NOT NULL,`changelog` TEXT,`description` TEXT,`user`TEXT NOT NULL,`timestamp` BIGINT(20) NOT NULL,INDEX newrelic_deployment_customer_id_index (`customer_id`),INDEX newrelic_deployment_application_ext_id_index (`application_ext_id`),INDEX newrelic_deployment_application_id_index (`application_id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBCreateNewrelicDeploymentTableTx will create the NewrelicDeployment table using the provided transction
func DBCreateNewrelicDeploymentTableTx(ctx context.Context, tx *sql.Tx) error {
	q := "CREATE TABLE `newrelic_deployment` (`id` VARCHAR(64) NOT NULL PRIMARY KEY,`checksum` CHAR(64),`customer_id` VARCHAR(64) NOT NULL,`application_ext_id` BIGINT(20) NOT NULL,`application_id` VARCHAR(64) NOT NULL,`ext_id` BIGINT(20) NOT NULL,`revision` TEXT NOT NULL,`changelog` TEXT,`description` TEXT,`user`TEXT NOT NULL,`timestamp` BIGINT(20) NOT NULL,INDEX newrelic_deployment_customer_id_index (`customer_id`),INDEX newrelic_deployment_application_ext_id_index (`application_ext_id`),INDEX newrelic_deployment_application_id_index (`application_id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// DBDropNewrelicDeploymentTable will drop the NewrelicDeployment table
func DBDropNewrelicDeploymentTable(ctx context.Context, db *sql.DB) error {
	q := "DROP TABLE IF EXISTS `newrelic_deployment`"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBDropNewrelicDeploymentTableTx will drop the NewrelicDeployment table using the provided transaction
func DBDropNewrelicDeploymentTableTx(ctx context.Context, tx *sql.Tx) error {
	q := "DROP TABLE IF EXISTS `newrelic_deployment`"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// CalculateChecksum will calculate a checksum of the SHA1 of all field values
func (t *NewrelicDeployment) CalculateChecksum() string {
	return orm.HashStrings(
		orm.ToString(t.ID),
		orm.ToString(t.CustomerID),
		orm.ToString(t.ApplicationExtID),
		orm.ToString(t.ApplicationID),
		orm.ToString(t.ExtID),
		orm.ToString(t.Revision),
		orm.ToString(t.Changelog),
		orm.ToString(t.Description),
		orm.ToString(t.User),
		orm.ToString(t.Timestamp),
	)
}

// DBCreate will create a new NewrelicDeployment record in the database
func (t *NewrelicDeployment) DBCreate(ctx context.Context, db *sql.DB) (sql.Result, error) {
	q := "INSERT INTO `newrelic_deployment` (`newrelic_deployment`.`id`,`newrelic_deployment`.`checksum`,`newrelic_deployment`.`customer_id`,`newrelic_deployment`.`application_ext_id`,`newrelic_deployment`.`application_id`,`newrelic_deployment`.`ext_id`,`newrelic_deployment`.`revision`,`newrelic_deployment`.`changelog`,`newrelic_deployment`.`description`,`newrelic_deployment`.`user`,`newrelic_deployment`.`timestamp`) VALUES (?,?,?,?,?,?,?,?,?,?,?)"
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
		orm.ToSQLInt64(t.ExtID),
		orm.ToSQLString(t.Revision),
		orm.ToSQLString(t.Changelog),
		orm.ToSQLString(t.Description),
		orm.ToSQLString(t.User),
		orm.ToSQLInt64(t.Timestamp),
	)
}

// DBCreateTx will create a new NewrelicDeployment record in the database using the provided transaction
func (t *NewrelicDeployment) DBCreateTx(ctx context.Context, tx *sql.Tx) (sql.Result, error) {
	q := "INSERT INTO `newrelic_deployment` (`newrelic_deployment`.`id`,`newrelic_deployment`.`checksum`,`newrelic_deployment`.`customer_id`,`newrelic_deployment`.`application_ext_id`,`newrelic_deployment`.`application_id`,`newrelic_deployment`.`ext_id`,`newrelic_deployment`.`revision`,`newrelic_deployment`.`changelog`,`newrelic_deployment`.`description`,`newrelic_deployment`.`user`,`newrelic_deployment`.`timestamp`) VALUES (?,?,?,?,?,?,?,?,?,?,?)"
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
		orm.ToSQLInt64(t.ExtID),
		orm.ToSQLString(t.Revision),
		orm.ToSQLString(t.Changelog),
		orm.ToSQLString(t.Description),
		orm.ToSQLString(t.User),
		orm.ToSQLInt64(t.Timestamp),
	)
}

// DBCreateIgnoreDuplicate will upsert the NewrelicDeployment record in the database
func (t *NewrelicDeployment) DBCreateIgnoreDuplicate(ctx context.Context, db *sql.DB) (sql.Result, error) {
	q := "INSERT INTO `newrelic_deployment` (`newrelic_deployment`.`id`,`newrelic_deployment`.`checksum`,`newrelic_deployment`.`customer_id`,`newrelic_deployment`.`application_ext_id`,`newrelic_deployment`.`application_id`,`newrelic_deployment`.`ext_id`,`newrelic_deployment`.`revision`,`newrelic_deployment`.`changelog`,`newrelic_deployment`.`description`,`newrelic_deployment`.`user`,`newrelic_deployment`.`timestamp`) VALUES (?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
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
		orm.ToSQLInt64(t.ExtID),
		orm.ToSQLString(t.Revision),
		orm.ToSQLString(t.Changelog),
		orm.ToSQLString(t.Description),
		orm.ToSQLString(t.User),
		orm.ToSQLInt64(t.Timestamp),
	)
}

// DBCreateIgnoreDuplicateTx will upsert the NewrelicDeployment record in the database using the provided transaction
func (t *NewrelicDeployment) DBCreateIgnoreDuplicateTx(ctx context.Context, tx *sql.Tx) (sql.Result, error) {
	q := "INSERT INTO `newrelic_deployment` (`newrelic_deployment`.`id`,`newrelic_deployment`.`checksum`,`newrelic_deployment`.`customer_id`,`newrelic_deployment`.`application_ext_id`,`newrelic_deployment`.`application_id`,`newrelic_deployment`.`ext_id`,`newrelic_deployment`.`revision`,`newrelic_deployment`.`changelog`,`newrelic_deployment`.`description`,`newrelic_deployment`.`user`,`newrelic_deployment`.`timestamp`) VALUES (?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
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
		orm.ToSQLInt64(t.ExtID),
		orm.ToSQLString(t.Revision),
		orm.ToSQLString(t.Changelog),
		orm.ToSQLString(t.Description),
		orm.ToSQLString(t.User),
		orm.ToSQLInt64(t.Timestamp),
	)
}

// DeleteAllNewrelicDeployments deletes all NewrelicDeployment records in the database with optional filters
func DeleteAllNewrelicDeployments(ctx context.Context, db *sql.DB, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(NewrelicDeploymentTableName),
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

// DeleteAllNewrelicDeploymentsTx deletes all NewrelicDeployment records in the database with optional filters using the provided transaction
func DeleteAllNewrelicDeploymentsTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(NewrelicDeploymentTableName),
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

// DBDelete will delete this NewrelicDeployment record in the database
func (t *NewrelicDeployment) DBDelete(ctx context.Context, db *sql.DB) (bool, error) {
	q := "DELETE FROM `newrelic_deployment` WHERE `id` = ?"
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

// DBDeleteTx will delete this NewrelicDeployment record in the database using the provided transaction
func (t *NewrelicDeployment) DBDeleteTx(ctx context.Context, tx *sql.Tx) (bool, error) {
	q := "DELETE FROM `newrelic_deployment` WHERE `id` = ?"
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

// DBUpdate will update the NewrelicDeployment record in the database
func (t *NewrelicDeployment) DBUpdate(ctx context.Context, db *sql.DB) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `newrelic_deployment` SET `checksum`=?,`customer_id`=?,`application_ext_id`=?,`application_id`=?,`ext_id`=?,`revision`=?,`changelog`=?,`description`=?,`user`=?,`timestamp`=? WHERE `id`=?"
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLInt64(t.ApplicationExtID),
		orm.ToSQLString(t.ApplicationID),
		orm.ToSQLInt64(t.ExtID),
		orm.ToSQLString(t.Revision),
		orm.ToSQLString(t.Changelog),
		orm.ToSQLString(t.Description),
		orm.ToSQLString(t.User),
		orm.ToSQLInt64(t.Timestamp),
		orm.ToSQLString(t.ID),
	)
}

// DBUpdateTx will update the NewrelicDeployment record in the database using the provided transaction
func (t *NewrelicDeployment) DBUpdateTx(ctx context.Context, tx *sql.Tx) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `newrelic_deployment` SET `checksum`=?,`customer_id`=?,`application_ext_id`=?,`application_id`=?,`ext_id`=?,`revision`=?,`changelog`=?,`description`=?,`user`=?,`timestamp`=? WHERE `id`=?"
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLInt64(t.ApplicationExtID),
		orm.ToSQLString(t.ApplicationID),
		orm.ToSQLInt64(t.ExtID),
		orm.ToSQLString(t.Revision),
		orm.ToSQLString(t.Changelog),
		orm.ToSQLString(t.Description),
		orm.ToSQLString(t.User),
		orm.ToSQLInt64(t.Timestamp),
		orm.ToSQLString(t.ID),
	)
}

// DBUpsert will upsert the NewrelicDeployment record in the database
func (t *NewrelicDeployment) DBUpsert(ctx context.Context, db *sql.DB, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `newrelic_deployment` (`newrelic_deployment`.`id`,`newrelic_deployment`.`checksum`,`newrelic_deployment`.`customer_id`,`newrelic_deployment`.`application_ext_id`,`newrelic_deployment`.`application_id`,`newrelic_deployment`.`ext_id`,`newrelic_deployment`.`revision`,`newrelic_deployment`.`changelog`,`newrelic_deployment`.`description`,`newrelic_deployment`.`user`,`newrelic_deployment`.`timestamp`) VALUES (?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `newrelic_deployment` (`newrelic_deployment`.`id`,`newrelic_deployment`.`checksum`,`newrelic_deployment`.`customer_id`,`newrelic_deployment`.`application_ext_id`,`newrelic_deployment`.`application_id`,`newrelic_deployment`.`ext_id`,`newrelic_deployment`.`revision`,`newrelic_deployment`.`changelog`,`newrelic_deployment`.`description`,`newrelic_deployment`.`user`,`newrelic_deployment`.`timestamp`) VALUES (?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`customer_id`=VALUES(`customer_id`),`application_ext_id`=VALUES(`application_ext_id`),`application_id`=VALUES(`application_id`),`ext_id`=VALUES(`ext_id`),`revision`=VALUES(`revision`),`changelog`=VALUES(`changelog`),`description`=VALUES(`description`),`user`=VALUES(`user`),`timestamp`=VALUES(`timestamp`)"
	}
	r, err := db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLInt64(t.ApplicationExtID),
		orm.ToSQLString(t.ApplicationID),
		orm.ToSQLInt64(t.ExtID),
		orm.ToSQLString(t.Revision),
		orm.ToSQLString(t.Changelog),
		orm.ToSQLString(t.Description),
		orm.ToSQLString(t.User),
		orm.ToSQLInt64(t.Timestamp),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBUpsertTx will upsert the NewrelicDeployment record in the database using the provided transaction
func (t *NewrelicDeployment) DBUpsertTx(ctx context.Context, tx *sql.Tx, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `newrelic_deployment` (`newrelic_deployment`.`id`,`newrelic_deployment`.`checksum`,`newrelic_deployment`.`customer_id`,`newrelic_deployment`.`application_ext_id`,`newrelic_deployment`.`application_id`,`newrelic_deployment`.`ext_id`,`newrelic_deployment`.`revision`,`newrelic_deployment`.`changelog`,`newrelic_deployment`.`description`,`newrelic_deployment`.`user`,`newrelic_deployment`.`timestamp`) VALUES (?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `newrelic_deployment` (`newrelic_deployment`.`id`,`newrelic_deployment`.`checksum`,`newrelic_deployment`.`customer_id`,`newrelic_deployment`.`application_ext_id`,`newrelic_deployment`.`application_id`,`newrelic_deployment`.`ext_id`,`newrelic_deployment`.`revision`,`newrelic_deployment`.`changelog`,`newrelic_deployment`.`description`,`newrelic_deployment`.`user`,`newrelic_deployment`.`timestamp`) VALUES (?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`customer_id`=VALUES(`customer_id`),`application_ext_id`=VALUES(`application_ext_id`),`application_id`=VALUES(`application_id`),`ext_id`=VALUES(`ext_id`),`revision`=VALUES(`revision`),`changelog`=VALUES(`changelog`),`description`=VALUES(`description`),`user`=VALUES(`user`),`timestamp`=VALUES(`timestamp`)"
	}
	r, err := tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLInt64(t.ApplicationExtID),
		orm.ToSQLString(t.ApplicationID),
		orm.ToSQLInt64(t.ExtID),
		orm.ToSQLString(t.Revision),
		orm.ToSQLString(t.Changelog),
		orm.ToSQLString(t.Description),
		orm.ToSQLString(t.User),
		orm.ToSQLInt64(t.Timestamp),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBFindOne will find a NewrelicDeployment record in the database with the primary key
func (t *NewrelicDeployment) DBFindOne(ctx context.Context, db *sql.DB, value string) (bool, error) {
	q := "SELECT `newrelic_deployment`.`id`,`newrelic_deployment`.`checksum`,`newrelic_deployment`.`customer_id`,`newrelic_deployment`.`application_ext_id`,`newrelic_deployment`.`application_id`,`newrelic_deployment`.`ext_id`,`newrelic_deployment`.`revision`,`newrelic_deployment`.`changelog`,`newrelic_deployment`.`description`,`newrelic_deployment`.`user`,`newrelic_deployment`.`timestamp` FROM `newrelic_deployment` WHERE `id` = ? LIMIT 1"
	row := db.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CustomerID sql.NullString
	var _ApplicationExtID sql.NullInt64
	var _ApplicationID sql.NullString
	var _ExtID sql.NullInt64
	var _Revision sql.NullString
	var _Changelog sql.NullString
	var _Description sql.NullString
	var _User sql.NullString
	var _Timestamp sql.NullInt64
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_ApplicationExtID,
		&_ApplicationID,
		&_ExtID,
		&_Revision,
		&_Changelog,
		&_Description,
		&_User,
		&_Timestamp,
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
	if _ExtID.Valid {
		t.SetExtID(_ExtID.Int64)
	}
	if _Revision.Valid {
		t.SetRevision(_Revision.String)
	}
	if _Changelog.Valid {
		t.SetChangelog(_Changelog.String)
	}
	if _Description.Valid {
		t.SetDescription(_Description.String)
	}
	if _User.Valid {
		t.SetUser(_User.String)
	}
	if _Timestamp.Valid {
		t.SetTimestamp(_Timestamp.Int64)
	}
	return true, nil
}

// DBFindOneTx will find a NewrelicDeployment record in the database with the primary key using the provided transaction
func (t *NewrelicDeployment) DBFindOneTx(ctx context.Context, tx *sql.Tx, value string) (bool, error) {
	q := "SELECT `newrelic_deployment`.`id`,`newrelic_deployment`.`checksum`,`newrelic_deployment`.`customer_id`,`newrelic_deployment`.`application_ext_id`,`newrelic_deployment`.`application_id`,`newrelic_deployment`.`ext_id`,`newrelic_deployment`.`revision`,`newrelic_deployment`.`changelog`,`newrelic_deployment`.`description`,`newrelic_deployment`.`user`,`newrelic_deployment`.`timestamp` FROM `newrelic_deployment` WHERE `id` = ? LIMIT 1"
	row := tx.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CustomerID sql.NullString
	var _ApplicationExtID sql.NullInt64
	var _ApplicationID sql.NullString
	var _ExtID sql.NullInt64
	var _Revision sql.NullString
	var _Changelog sql.NullString
	var _Description sql.NullString
	var _User sql.NullString
	var _Timestamp sql.NullInt64
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_ApplicationExtID,
		&_ApplicationID,
		&_ExtID,
		&_Revision,
		&_Changelog,
		&_Description,
		&_User,
		&_Timestamp,
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
	if _ExtID.Valid {
		t.SetExtID(_ExtID.Int64)
	}
	if _Revision.Valid {
		t.SetRevision(_Revision.String)
	}
	if _Changelog.Valid {
		t.SetChangelog(_Changelog.String)
	}
	if _Description.Valid {
		t.SetDescription(_Description.String)
	}
	if _User.Valid {
		t.SetUser(_User.String)
	}
	if _Timestamp.Valid {
		t.SetTimestamp(_Timestamp.Int64)
	}
	return true, nil
}

// FindNewrelicDeployments will find a NewrelicDeployment record in the database with the provided parameters
func FindNewrelicDeployments(ctx context.Context, db *sql.DB, _params ...interface{}) ([]*NewrelicDeployment, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("customer_id"),
		orm.Column("application_ext_id"),
		orm.Column("application_id"),
		orm.Column("ext_id"),
		orm.Column("revision"),
		orm.Column("changelog"),
		orm.Column("description"),
		orm.Column("user"),
		orm.Column("timestamp"),
		orm.Table(NewrelicDeploymentTableName),
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
	results := make([]*NewrelicDeployment, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _ApplicationExtID sql.NullInt64
		var _ApplicationID sql.NullString
		var _ExtID sql.NullInt64
		var _Revision sql.NullString
		var _Changelog sql.NullString
		var _Description sql.NullString
		var _User sql.NullString
		var _Timestamp sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_ApplicationExtID,
			&_ApplicationID,
			&_ExtID,
			&_Revision,
			&_Changelog,
			&_Description,
			&_User,
			&_Timestamp,
		)
		if err != nil {
			return nil, err
		}
		t := &NewrelicDeployment{}
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
		if _ExtID.Valid {
			t.SetExtID(_ExtID.Int64)
		}
		if _Revision.Valid {
			t.SetRevision(_Revision.String)
		}
		if _Changelog.Valid {
			t.SetChangelog(_Changelog.String)
		}
		if _Description.Valid {
			t.SetDescription(_Description.String)
		}
		if _User.Valid {
			t.SetUser(_User.String)
		}
		if _Timestamp.Valid {
			t.SetTimestamp(_Timestamp.Int64)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindNewrelicDeploymentsTx will find a NewrelicDeployment record in the database with the provided parameters using the provided transaction
func FindNewrelicDeploymentsTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) ([]*NewrelicDeployment, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("customer_id"),
		orm.Column("application_ext_id"),
		orm.Column("application_id"),
		orm.Column("ext_id"),
		orm.Column("revision"),
		orm.Column("changelog"),
		orm.Column("description"),
		orm.Column("user"),
		orm.Column("timestamp"),
		orm.Table(NewrelicDeploymentTableName),
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
	results := make([]*NewrelicDeployment, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _ApplicationExtID sql.NullInt64
		var _ApplicationID sql.NullString
		var _ExtID sql.NullInt64
		var _Revision sql.NullString
		var _Changelog sql.NullString
		var _Description sql.NullString
		var _User sql.NullString
		var _Timestamp sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_ApplicationExtID,
			&_ApplicationID,
			&_ExtID,
			&_Revision,
			&_Changelog,
			&_Description,
			&_User,
			&_Timestamp,
		)
		if err != nil {
			return nil, err
		}
		t := &NewrelicDeployment{}
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
		if _ExtID.Valid {
			t.SetExtID(_ExtID.Int64)
		}
		if _Revision.Valid {
			t.SetRevision(_Revision.String)
		}
		if _Changelog.Valid {
			t.SetChangelog(_Changelog.String)
		}
		if _Description.Valid {
			t.SetDescription(_Description.String)
		}
		if _User.Valid {
			t.SetUser(_User.String)
		}
		if _Timestamp.Valid {
			t.SetTimestamp(_Timestamp.Int64)
		}
		results = append(results, t)
	}
	return results, nil
}

// DBFind will find a NewrelicDeployment record in the database with the provided parameters
func (t *NewrelicDeployment) DBFind(ctx context.Context, db *sql.DB, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("customer_id"),
		orm.Column("application_ext_id"),
		orm.Column("application_id"),
		orm.Column("ext_id"),
		orm.Column("revision"),
		orm.Column("changelog"),
		orm.Column("description"),
		orm.Column("user"),
		orm.Column("timestamp"),
		orm.Table(NewrelicDeploymentTableName),
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
	var _ExtID sql.NullInt64
	var _Revision sql.NullString
	var _Changelog sql.NullString
	var _Description sql.NullString
	var _User sql.NullString
	var _Timestamp sql.NullInt64
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_ApplicationExtID,
		&_ApplicationID,
		&_ExtID,
		&_Revision,
		&_Changelog,
		&_Description,
		&_User,
		&_Timestamp,
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
	if _ExtID.Valid {
		t.SetExtID(_ExtID.Int64)
	}
	if _Revision.Valid {
		t.SetRevision(_Revision.String)
	}
	if _Changelog.Valid {
		t.SetChangelog(_Changelog.String)
	}
	if _Description.Valid {
		t.SetDescription(_Description.String)
	}
	if _User.Valid {
		t.SetUser(_User.String)
	}
	if _Timestamp.Valid {
		t.SetTimestamp(_Timestamp.Int64)
	}
	return true, nil
}

// DBFindTx will find a NewrelicDeployment record in the database with the provided parameters using the provided transaction
func (t *NewrelicDeployment) DBFindTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("customer_id"),
		orm.Column("application_ext_id"),
		orm.Column("application_id"),
		orm.Column("ext_id"),
		orm.Column("revision"),
		orm.Column("changelog"),
		orm.Column("description"),
		orm.Column("user"),
		orm.Column("timestamp"),
		orm.Table(NewrelicDeploymentTableName),
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
	var _ExtID sql.NullInt64
	var _Revision sql.NullString
	var _Changelog sql.NullString
	var _Description sql.NullString
	var _User sql.NullString
	var _Timestamp sql.NullInt64
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_ApplicationExtID,
		&_ApplicationID,
		&_ExtID,
		&_Revision,
		&_Changelog,
		&_Description,
		&_User,
		&_Timestamp,
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
	if _ExtID.Valid {
		t.SetExtID(_ExtID.Int64)
	}
	if _Revision.Valid {
		t.SetRevision(_Revision.String)
	}
	if _Changelog.Valid {
		t.SetChangelog(_Changelog.String)
	}
	if _Description.Valid {
		t.SetDescription(_Description.String)
	}
	if _User.Valid {
		t.SetUser(_User.String)
	}
	if _Timestamp.Valid {
		t.SetTimestamp(_Timestamp.Int64)
	}
	return true, nil
}

// CountNewrelicDeployments will find the count of NewrelicDeployment records in the database
func CountNewrelicDeployments(ctx context.Context, db *sql.DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(NewrelicDeploymentTableName),
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

// CountNewrelicDeploymentsTx will find the count of NewrelicDeployment records in the database using the provided transaction
func CountNewrelicDeploymentsTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(NewrelicDeploymentTableName),
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

// DBCount will find the count of NewrelicDeployment records in the database
func (t *NewrelicDeployment) DBCount(ctx context.Context, db *sql.DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(NewrelicDeploymentTableName),
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

// DBCountTx will find the count of NewrelicDeployment records in the database using the provided transaction
func (t *NewrelicDeployment) DBCountTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(NewrelicDeploymentTableName),
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

// DBExists will return true if the NewrelicDeployment record exists in the database
func (t *NewrelicDeployment) DBExists(ctx context.Context, db *sql.DB) (bool, error) {
	q := "SELECT `id` FROM `newrelic_deployment` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := db.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// DBExistsTx will return true if the NewrelicDeployment record exists in the database using the provided transaction
func (t *NewrelicDeployment) DBExistsTx(ctx context.Context, tx *sql.Tx) (bool, error) {
	q := "SELECT `id` FROM `newrelic_deployment` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := tx.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// PrimaryKeyColumn returns the column name for the primary key
func (t *NewrelicDeployment) PrimaryKeyColumn() string {
	return NewrelicDeploymentColumnID
}

// PrimaryKeyColumnType returns the primary key column Go type as a string
func (t *NewrelicDeployment) PrimaryKeyColumnType() string {
	return "string"
}

// PrimaryKey returns the primary key column value
func (t *NewrelicDeployment) PrimaryKey() interface{} {
	return t.ID
}
