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

var _ Model = (*ProcessingError)(nil)
var _ CSVWriter = (*ProcessingError)(nil)
var _ JSONWriter = (*ProcessingError)(nil)
var _ Checksum = (*ProcessingError)(nil)

// ProcessingErrorTableName is the name of the table in SQL
const ProcessingErrorTableName = "processing_error"

var ProcessingErrorColumns = []string{
	"id",
	"checksum",
	"created_at",
	"type",
	"message",
	"fatal",
	"customer_id",
}

// ProcessingError table
type ProcessingError struct {
	Checksum   *string `json:"checksum,omitempty"`
	CreatedAt  int64   `json:"created_at"`
	CustomerID string  `json:"customer_id"`
	Fatal      bool    `json:"fatal"`
	ID         string  `json:"id"`
	Message    string  `json:"message"`
	Type       string  `json:"type"`
}

// TableName returns the SQL table name for ProcessingError and satifies the Model interface
func (t *ProcessingError) TableName() string {
	return ProcessingErrorTableName
}

// ToCSV will serialize the ProcessingError instance to a CSV compatible array of strings
func (t *ProcessingError) ToCSV() []string {
	return []string{
		t.ID,
		t.CalculateChecksum(),
		toCSVString(t.CreatedAt),
		t.Type,
		t.Message,
		toCSVBool(t.Fatal),
		t.CustomerID,
	}
}

// WriteCSV will serialize the ProcessingError instance to the writer as CSV and satisfies the CSVWriter interface
func (t *ProcessingError) WriteCSV(w *csv.Writer) error {
	return w.Write(t.ToCSV())
}

// WriteJSON will serialize the ProcessingError instance to the writer as JSON and satisfies the JSONWriter interface
func (t *ProcessingError) WriteJSON(w io.Writer, indent ...bool) error {
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

// NewProcessingErrorReader creates a JSON reader which can read in ProcessingError objects serialized as JSON either as an array, single object or json new lines
// and writes each ProcessingError to the channel provided
func NewProcessingErrorReader(r io.Reader, ch chan<- ProcessingError) error {
	return orm.Deserialize(r, func(buf json.RawMessage) error {
		dec := json.NewDecoder(bytes.NewBuffer(buf))
		e := ProcessingError{}
		if err := dec.Decode(&e); err != nil {
			return err
		}
		ch <- e
		return nil
	})
}

// NewCSVProcessingErrorReaderDir will read the reader as CSV and emit each record to the channel provided
func NewCSVProcessingErrorReader(r io.Reader, ch chan<- ProcessingError) error {
	cr := csv.NewReader(r)
	for {
		record, err := cr.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		ch <- ProcessingError{
			ID:         record[0],
			Checksum:   fromStringPointer(record[1]),
			CreatedAt:  fromCSVInt64(record[2]),
			Type:       record[3],
			Message:    record[4],
			Fatal:      fromCSVBool(record[5]),
			CustomerID: record[6],
		}
	}
	return nil
}

// NewCSVProcessingErrorReaderFile will read the file as a CSV and emit each record to the channel provided
func NewCSVProcessingErrorReaderFile(fp string, ch chan<- ProcessingError) error {
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
	return NewCSVProcessingErrorReader(fc, ch)
}

// NewCSVProcessingErrorReaderDir will read the processing_error.csv.gz file as a CSV and emit each record to the channel provided
func NewCSVProcessingErrorReaderDir(dir string, ch chan<- ProcessingError) error {
	return NewCSVProcessingErrorReaderFile(filepath.Join(dir, "processing_error.csv.gz"), ch)
}

// ProcessingErrorCSVDeduper is a function callback which takes the existing value (a) and the new value (b)
// and the return value should be the one to use (or a new one, if applicable). return nil
// to skip processing of this record
type ProcessingErrorCSVDeduper func(a ProcessingError, b ProcessingError) *ProcessingError

// ProcessingErrorCSVDedupeDisabled is set on whether the CSV writer should de-dupe values by key
var ProcessingErrorCSVDedupeDisabled bool

// NewProcessingErrorCSVWriterSize creates a batch writer that will write each ProcessingError into a CSV file
// this method will automatically de-duplicate entries using the primary key. if the checksum
// for a newer item with the same primary key doesn't match a previously sent item, the newer
// one will be used
func NewProcessingErrorCSVWriterSize(w io.Writer, size int, dedupers ...ProcessingErrorCSVDeduper) (chan ProcessingError, chan bool, error) {
	cw := csv.NewWriter(w)
	ch := make(chan ProcessingError, size)
	done := make(chan bool)
	go func() {
		defer func() { done <- true }()
		dodedupe := !ProcessingErrorCSVDedupeDisabled
		var kv map[string]*ProcessingError
		var deduper ProcessingErrorCSVDeduper
		if dedupers != nil && len(dedupers) > 0 {
			deduper = dedupers[0]
			dodedupe = true
		}
		if dodedupe {
			kv = make(map[string]*ProcessingError)
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

// ProcessingErrorCSVDefaultSize is the default channel buffer size if not provided
var ProcessingErrorCSVDefaultSize = 100

// NewProcessingErrorCSVWriter creates a batch writer that will write each ProcessingError into a CSV file
func NewProcessingErrorCSVWriter(w io.Writer, dedupers ...ProcessingErrorCSVDeduper) (chan ProcessingError, chan bool, error) {
	return NewProcessingErrorCSVWriterSize(w, ProcessingErrorCSVDefaultSize, dedupers...)
}

// NewProcessingErrorCSVWriterDir creates a batch writer that will write each ProcessingError into a CSV file named processing_error.csv.gz in dir
func NewProcessingErrorCSVWriterDir(dir string, dedupers ...ProcessingErrorCSVDeduper) (chan ProcessingError, chan bool, error) {
	return NewProcessingErrorCSVWriterFile(filepath.Join(dir, "processing_error.csv.gz"), dedupers...)
}

// NewProcessingErrorCSVWriterFile creates a batch writer that will write each ProcessingError into a CSV file
func NewProcessingErrorCSVWriterFile(fn string, dedupers ...ProcessingErrorCSVDeduper) (chan ProcessingError, chan bool, error) {
	f, err := os.Create(fn)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening CSV file %s. %v", fn, err)
	}
	var fc io.WriteCloser = f
	if filepath.Ext(fn) == ".gz" {
		w, _ := gzip.NewWriterLevel(f, gzip.BestCompression)
		fc = w
	}
	ch, done, err := NewProcessingErrorCSVWriter(fc, dedupers...)
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

type ProcessingErrorDBAction func(ctx context.Context, db *sql.DB, record ProcessingError) error

// NewProcessingErrorDBWriterSize creates a DB writer that will write each issue into the DB
func NewProcessingErrorDBWriterSize(ctx context.Context, db *sql.DB, errors chan<- error, size int, actions ...ProcessingErrorDBAction) (chan ProcessingError, chan bool, error) {
	ch := make(chan ProcessingError, size)
	done := make(chan bool)
	var action ProcessingErrorDBAction
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

// NewProcessingErrorDBWriter creates a DB writer that will write each issue into the DB
func NewProcessingErrorDBWriter(ctx context.Context, db *sql.DB, errors chan<- error, actions ...ProcessingErrorDBAction) (chan ProcessingError, chan bool, error) {
	return NewProcessingErrorDBWriterSize(ctx, db, errors, 100, actions...)
}

// ProcessingErrorColumnID is the ID SQL column name for the ProcessingError table
const ProcessingErrorColumnID = "id"

// ProcessingErrorEscapedColumnID is the escaped ID SQL column name for the ProcessingError table
const ProcessingErrorEscapedColumnID = "`id`"

// ProcessingErrorColumnChecksum is the Checksum SQL column name for the ProcessingError table
const ProcessingErrorColumnChecksum = "checksum"

// ProcessingErrorEscapedColumnChecksum is the escaped Checksum SQL column name for the ProcessingError table
const ProcessingErrorEscapedColumnChecksum = "`checksum`"

// ProcessingErrorColumnCreatedAt is the CreatedAt SQL column name for the ProcessingError table
const ProcessingErrorColumnCreatedAt = "created_at"

// ProcessingErrorEscapedColumnCreatedAt is the escaped CreatedAt SQL column name for the ProcessingError table
const ProcessingErrorEscapedColumnCreatedAt = "`created_at`"

// ProcessingErrorColumnType is the Type SQL column name for the ProcessingError table
const ProcessingErrorColumnType = "type"

// ProcessingErrorEscapedColumnType is the escaped Type SQL column name for the ProcessingError table
const ProcessingErrorEscapedColumnType = "`type`"

// ProcessingErrorColumnMessage is the Message SQL column name for the ProcessingError table
const ProcessingErrorColumnMessage = "message"

// ProcessingErrorEscapedColumnMessage is the escaped Message SQL column name for the ProcessingError table
const ProcessingErrorEscapedColumnMessage = "`message`"

// ProcessingErrorColumnFatal is the Fatal SQL column name for the ProcessingError table
const ProcessingErrorColumnFatal = "fatal"

// ProcessingErrorEscapedColumnFatal is the escaped Fatal SQL column name for the ProcessingError table
const ProcessingErrorEscapedColumnFatal = "`fatal`"

// ProcessingErrorColumnCustomerID is the CustomerID SQL column name for the ProcessingError table
const ProcessingErrorColumnCustomerID = "customer_id"

// ProcessingErrorEscapedColumnCustomerID is the escaped CustomerID SQL column name for the ProcessingError table
const ProcessingErrorEscapedColumnCustomerID = "`customer_id`"

// GetID will return the ProcessingError ID value
func (t *ProcessingError) GetID() string {
	return t.ID
}

// SetID will set the ProcessingError ID value
func (t *ProcessingError) SetID(v string) {
	t.ID = v
}

// FindProcessingErrorByID will find a ProcessingError by ID
func FindProcessingErrorByID(ctx context.Context, db *sql.DB, value string) (*ProcessingError, error) {
	q := "SELECT `processing_error`.`id`,`processing_error`.`checksum`,`processing_error`.`created_at`,`processing_error`.`type`,`processing_error`.`message`,`processing_error`.`fatal`,`processing_error`.`customer_id` FROM `processing_error` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CreatedAt sql.NullInt64
	var _Type sql.NullString
	var _Message sql.NullString
	var _Fatal sql.NullBool
	var _CustomerID sql.NullString
	err := db.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_CreatedAt,
		&_Type,
		&_Message,
		&_Fatal,
		&_CustomerID,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &ProcessingError{}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _Checksum.Valid {
		t.SetChecksum(_Checksum.String)
	}
	if _CreatedAt.Valid {
		t.SetCreatedAt(_CreatedAt.Int64)
	}
	if _Type.Valid {
		t.SetType(_Type.String)
	}
	if _Message.Valid {
		t.SetMessage(_Message.String)
	}
	if _Fatal.Valid {
		t.SetFatal(_Fatal.Bool)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	return t, nil
}

// FindProcessingErrorByIDTx will find a ProcessingError by ID using the provided transaction
func FindProcessingErrorByIDTx(ctx context.Context, tx *sql.Tx, value string) (*ProcessingError, error) {
	q := "SELECT `processing_error`.`id`,`processing_error`.`checksum`,`processing_error`.`created_at`,`processing_error`.`type`,`processing_error`.`message`,`processing_error`.`fatal`,`processing_error`.`customer_id` FROM `processing_error` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CreatedAt sql.NullInt64
	var _Type sql.NullString
	var _Message sql.NullString
	var _Fatal sql.NullBool
	var _CustomerID sql.NullString
	err := tx.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_CreatedAt,
		&_Type,
		&_Message,
		&_Fatal,
		&_CustomerID,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &ProcessingError{}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _Checksum.Valid {
		t.SetChecksum(_Checksum.String)
	}
	if _CreatedAt.Valid {
		t.SetCreatedAt(_CreatedAt.Int64)
	}
	if _Type.Valid {
		t.SetType(_Type.String)
	}
	if _Message.Valid {
		t.SetMessage(_Message.String)
	}
	if _Fatal.Valid {
		t.SetFatal(_Fatal.Bool)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	return t, nil
}

// GetChecksum will return the ProcessingError Checksum value
func (t *ProcessingError) GetChecksum() string {
	if t.Checksum == nil {
		return ""
	}
	return *t.Checksum
}

// SetChecksum will set the ProcessingError Checksum value
func (t *ProcessingError) SetChecksum(v string) {
	t.Checksum = &v
}

// GetCreatedAt will return the ProcessingError CreatedAt value
func (t *ProcessingError) GetCreatedAt() int64 {
	return t.CreatedAt
}

// SetCreatedAt will set the ProcessingError CreatedAt value
func (t *ProcessingError) SetCreatedAt(v int64) {
	t.CreatedAt = v
}

// GetType will return the ProcessingError Type value
func (t *ProcessingError) GetType() string {
	return t.Type
}

// SetType will set the ProcessingError Type value
func (t *ProcessingError) SetType(v string) {
	t.Type = v
}

// GetMessage will return the ProcessingError Message value
func (t *ProcessingError) GetMessage() string {
	return t.Message
}

// SetMessage will set the ProcessingError Message value
func (t *ProcessingError) SetMessage(v string) {
	t.Message = v
}

// GetFatal will return the ProcessingError Fatal value
func (t *ProcessingError) GetFatal() bool {
	return t.Fatal
}

// SetFatal will set the ProcessingError Fatal value
func (t *ProcessingError) SetFatal(v bool) {
	t.Fatal = v
}

// GetCustomerID will return the ProcessingError CustomerID value
func (t *ProcessingError) GetCustomerID() string {
	return t.CustomerID
}

// SetCustomerID will set the ProcessingError CustomerID value
func (t *ProcessingError) SetCustomerID(v string) {
	t.CustomerID = v
}

// FindProcessingErrorsByCustomerID will find all ProcessingErrors by the CustomerID value
func FindProcessingErrorsByCustomerID(ctx context.Context, db *sql.DB, value string) ([]*ProcessingError, error) {
	q := "SELECT `processing_error`.`id`,`processing_error`.`checksum`,`processing_error`.`created_at`,`processing_error`.`type`,`processing_error`.`message`,`processing_error`.`fatal`,`processing_error`.`customer_id` FROM `processing_error` WHERE `customer_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*ProcessingError, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CreatedAt sql.NullInt64
		var _Type sql.NullString
		var _Message sql.NullString
		var _Fatal sql.NullBool
		var _CustomerID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CreatedAt,
			&_Type,
			&_Message,
			&_Fatal,
			&_CustomerID,
		)
		if err != nil {
			return nil, err
		}
		t := &ProcessingError{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CreatedAt.Valid {
			t.SetCreatedAt(_CreatedAt.Int64)
		}
		if _Type.Valid {
			t.SetType(_Type.String)
		}
		if _Message.Valid {
			t.SetMessage(_Message.String)
		}
		if _Fatal.Valid {
			t.SetFatal(_Fatal.Bool)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindProcessingErrorsByCustomerIDTx will find all ProcessingErrors by the CustomerID value using the provided transaction
func FindProcessingErrorsByCustomerIDTx(ctx context.Context, tx *sql.Tx, value string) ([]*ProcessingError, error) {
	q := "SELECT `processing_error`.`id`,`processing_error`.`checksum`,`processing_error`.`created_at`,`processing_error`.`type`,`processing_error`.`message`,`processing_error`.`fatal`,`processing_error`.`customer_id` FROM `processing_error` WHERE `customer_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*ProcessingError, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CreatedAt sql.NullInt64
		var _Type sql.NullString
		var _Message sql.NullString
		var _Fatal sql.NullBool
		var _CustomerID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CreatedAt,
			&_Type,
			&_Message,
			&_Fatal,
			&_CustomerID,
		)
		if err != nil {
			return nil, err
		}
		t := &ProcessingError{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CreatedAt.Valid {
			t.SetCreatedAt(_CreatedAt.Int64)
		}
		if _Type.Valid {
			t.SetType(_Type.String)
		}
		if _Message.Valid {
			t.SetMessage(_Message.String)
		}
		if _Fatal.Valid {
			t.SetFatal(_Fatal.Bool)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

func (t *ProcessingError) toTimestamp(value time.Time) *timestamp.Timestamp {
	ts, _ := ptypes.TimestampProto(value)
	return ts
}

// DBCreateProcessingErrorTable will create the ProcessingError table
func DBCreateProcessingErrorTable(ctx context.Context, db *sql.DB) error {
	q := "CREATE TABLE `processing_error` (`id` VARCHAR(64) NOT NULL PRIMARY KEY,`checksum` CHAR(64),`created_at`BIGINT(20) UNSIGNED NOT NULL,`type`TEXT NOT NULL,`message`MEDIUMTEXT NOT NULL,`fatal` TINYINT(1) NOT NULL,`customer_id` VARCHAR(64) NOT NULL,INDEX processing_error_customer_id_index (`customer_id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBCreateProcessingErrorTableTx will create the ProcessingError table using the provided transction
func DBCreateProcessingErrorTableTx(ctx context.Context, tx *sql.Tx) error {
	q := "CREATE TABLE `processing_error` (`id` VARCHAR(64) NOT NULL PRIMARY KEY,`checksum` CHAR(64),`created_at`BIGINT(20) UNSIGNED NOT NULL,`type`TEXT NOT NULL,`message`MEDIUMTEXT NOT NULL,`fatal` TINYINT(1) NOT NULL,`customer_id` VARCHAR(64) NOT NULL,INDEX processing_error_customer_id_index (`customer_id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// DBDropProcessingErrorTable will drop the ProcessingError table
func DBDropProcessingErrorTable(ctx context.Context, db *sql.DB) error {
	q := "DROP TABLE IF EXISTS `processing_error`"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBDropProcessingErrorTableTx will drop the ProcessingError table using the provided transaction
func DBDropProcessingErrorTableTx(ctx context.Context, tx *sql.Tx) error {
	q := "DROP TABLE IF EXISTS `processing_error`"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// CalculateChecksum will calculate a checksum of the SHA1 of all field values
func (t *ProcessingError) CalculateChecksum() string {
	return orm.HashStrings(
		orm.ToString(t.ID),
		orm.ToString(t.CreatedAt),
		orm.ToString(t.Type),
		orm.ToString(t.Message),
		orm.ToString(t.Fatal),
		orm.ToString(t.CustomerID),
	)
}

// DBCreate will create a new ProcessingError record in the database
func (t *ProcessingError) DBCreate(ctx context.Context, db *sql.DB) (sql.Result, error) {
	q := "INSERT INTO `processing_error` (`processing_error`.`id`,`processing_error`.`checksum`,`processing_error`.`created_at`,`processing_error`.`type`,`processing_error`.`message`,`processing_error`.`fatal`,`processing_error`.`customer_id`) VALUES (?,?,?,?,?,?,?)"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLString(t.Type),
		orm.ToSQLString(t.Message),
		orm.ToSQLBool(t.Fatal),
		orm.ToSQLString(t.CustomerID),
	)
}

// DBCreateTx will create a new ProcessingError record in the database using the provided transaction
func (t *ProcessingError) DBCreateTx(ctx context.Context, tx *sql.Tx) (sql.Result, error) {
	q := "INSERT INTO `processing_error` (`processing_error`.`id`,`processing_error`.`checksum`,`processing_error`.`created_at`,`processing_error`.`type`,`processing_error`.`message`,`processing_error`.`fatal`,`processing_error`.`customer_id`) VALUES (?,?,?,?,?,?,?)"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLString(t.Type),
		orm.ToSQLString(t.Message),
		orm.ToSQLBool(t.Fatal),
		orm.ToSQLString(t.CustomerID),
	)
}

// DBCreateIgnoreDuplicate will upsert the ProcessingError record in the database
func (t *ProcessingError) DBCreateIgnoreDuplicate(ctx context.Context, db *sql.DB) (sql.Result, error) {
	q := "INSERT INTO `processing_error` (`processing_error`.`id`,`processing_error`.`checksum`,`processing_error`.`created_at`,`processing_error`.`type`,`processing_error`.`message`,`processing_error`.`fatal`,`processing_error`.`customer_id`) VALUES (?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLString(t.Type),
		orm.ToSQLString(t.Message),
		orm.ToSQLBool(t.Fatal),
		orm.ToSQLString(t.CustomerID),
	)
}

// DBCreateIgnoreDuplicateTx will upsert the ProcessingError record in the database using the provided transaction
func (t *ProcessingError) DBCreateIgnoreDuplicateTx(ctx context.Context, tx *sql.Tx) (sql.Result, error) {
	q := "INSERT INTO `processing_error` (`processing_error`.`id`,`processing_error`.`checksum`,`processing_error`.`created_at`,`processing_error`.`type`,`processing_error`.`message`,`processing_error`.`fatal`,`processing_error`.`customer_id`) VALUES (?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLString(t.Type),
		orm.ToSQLString(t.Message),
		orm.ToSQLBool(t.Fatal),
		orm.ToSQLString(t.CustomerID),
	)
}

// DeleteAllProcessingErrors deletes all ProcessingError records in the database with optional filters
func DeleteAllProcessingErrors(ctx context.Context, db *sql.DB, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(ProcessingErrorTableName),
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

// DeleteAllProcessingErrorsTx deletes all ProcessingError records in the database with optional filters using the provided transaction
func DeleteAllProcessingErrorsTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(ProcessingErrorTableName),
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

// DBDelete will delete this ProcessingError record in the database
func (t *ProcessingError) DBDelete(ctx context.Context, db *sql.DB) (bool, error) {
	q := "DELETE FROM `processing_error` WHERE `id` = ?"
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

// DBDeleteTx will delete this ProcessingError record in the database using the provided transaction
func (t *ProcessingError) DBDeleteTx(ctx context.Context, tx *sql.Tx) (bool, error) {
	q := "DELETE FROM `processing_error` WHERE `id` = ?"
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

// DBUpdate will update the ProcessingError record in the database
func (t *ProcessingError) DBUpdate(ctx context.Context, db *sql.DB) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `processing_error` SET `checksum`=?,`created_at`=?,`type`=?,`message`=?,`fatal`=?,`customer_id`=? WHERE `id`=?"
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLString(t.Type),
		orm.ToSQLString(t.Message),
		orm.ToSQLBool(t.Fatal),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.ID),
	)
}

// DBUpdateTx will update the ProcessingError record in the database using the provided transaction
func (t *ProcessingError) DBUpdateTx(ctx context.Context, tx *sql.Tx) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `processing_error` SET `checksum`=?,`created_at`=?,`type`=?,`message`=?,`fatal`=?,`customer_id`=? WHERE `id`=?"
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLString(t.Type),
		orm.ToSQLString(t.Message),
		orm.ToSQLBool(t.Fatal),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.ID),
	)
}

// DBUpsert will upsert the ProcessingError record in the database
func (t *ProcessingError) DBUpsert(ctx context.Context, db *sql.DB, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `processing_error` (`processing_error`.`id`,`processing_error`.`checksum`,`processing_error`.`created_at`,`processing_error`.`type`,`processing_error`.`message`,`processing_error`.`fatal`,`processing_error`.`customer_id`) VALUES (?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `processing_error` (`processing_error`.`id`,`processing_error`.`checksum`,`processing_error`.`created_at`,`processing_error`.`type`,`processing_error`.`message`,`processing_error`.`fatal`,`processing_error`.`customer_id`) VALUES (?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`created_at`=VALUES(`created_at`),`type`=VALUES(`type`),`message`=VALUES(`message`),`fatal`=VALUES(`fatal`),`customer_id`=VALUES(`customer_id`)"
	}
	r, err := db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLString(t.Type),
		orm.ToSQLString(t.Message),
		orm.ToSQLBool(t.Fatal),
		orm.ToSQLString(t.CustomerID),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBUpsertTx will upsert the ProcessingError record in the database using the provided transaction
func (t *ProcessingError) DBUpsertTx(ctx context.Context, tx *sql.Tx, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `processing_error` (`processing_error`.`id`,`processing_error`.`checksum`,`processing_error`.`created_at`,`processing_error`.`type`,`processing_error`.`message`,`processing_error`.`fatal`,`processing_error`.`customer_id`) VALUES (?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `processing_error` (`processing_error`.`id`,`processing_error`.`checksum`,`processing_error`.`created_at`,`processing_error`.`type`,`processing_error`.`message`,`processing_error`.`fatal`,`processing_error`.`customer_id`) VALUES (?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`created_at`=VALUES(`created_at`),`type`=VALUES(`type`),`message`=VALUES(`message`),`fatal`=VALUES(`fatal`),`customer_id`=VALUES(`customer_id`)"
	}
	r, err := tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLString(t.Type),
		orm.ToSQLString(t.Message),
		orm.ToSQLBool(t.Fatal),
		orm.ToSQLString(t.CustomerID),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBFindOne will find a ProcessingError record in the database with the primary key
func (t *ProcessingError) DBFindOne(ctx context.Context, db *sql.DB, value string) (bool, error) {
	q := "SELECT `processing_error`.`id`,`processing_error`.`checksum`,`processing_error`.`created_at`,`processing_error`.`type`,`processing_error`.`message`,`processing_error`.`fatal`,`processing_error`.`customer_id` FROM `processing_error` WHERE `id` = ? LIMIT 1"
	row := db.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CreatedAt sql.NullInt64
	var _Type sql.NullString
	var _Message sql.NullString
	var _Fatal sql.NullBool
	var _CustomerID sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CreatedAt,
		&_Type,
		&_Message,
		&_Fatal,
		&_CustomerID,
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
	if _CreatedAt.Valid {
		t.SetCreatedAt(_CreatedAt.Int64)
	}
	if _Type.Valid {
		t.SetType(_Type.String)
	}
	if _Message.Valid {
		t.SetMessage(_Message.String)
	}
	if _Fatal.Valid {
		t.SetFatal(_Fatal.Bool)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	return true, nil
}

// DBFindOneTx will find a ProcessingError record in the database with the primary key using the provided transaction
func (t *ProcessingError) DBFindOneTx(ctx context.Context, tx *sql.Tx, value string) (bool, error) {
	q := "SELECT `processing_error`.`id`,`processing_error`.`checksum`,`processing_error`.`created_at`,`processing_error`.`type`,`processing_error`.`message`,`processing_error`.`fatal`,`processing_error`.`customer_id` FROM `processing_error` WHERE `id` = ? LIMIT 1"
	row := tx.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CreatedAt sql.NullInt64
	var _Type sql.NullString
	var _Message sql.NullString
	var _Fatal sql.NullBool
	var _CustomerID sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CreatedAt,
		&_Type,
		&_Message,
		&_Fatal,
		&_CustomerID,
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
	if _CreatedAt.Valid {
		t.SetCreatedAt(_CreatedAt.Int64)
	}
	if _Type.Valid {
		t.SetType(_Type.String)
	}
	if _Message.Valid {
		t.SetMessage(_Message.String)
	}
	if _Fatal.Valid {
		t.SetFatal(_Fatal.Bool)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	return true, nil
}

// FindProcessingErrors will find a ProcessingError record in the database with the provided parameters
func FindProcessingErrors(ctx context.Context, db *sql.DB, _params ...interface{}) ([]*ProcessingError, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("created_at"),
		orm.Column("type"),
		orm.Column("message"),
		orm.Column("fatal"),
		orm.Column("customer_id"),
		orm.Table(ProcessingErrorTableName),
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
	results := make([]*ProcessingError, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CreatedAt sql.NullInt64
		var _Type sql.NullString
		var _Message sql.NullString
		var _Fatal sql.NullBool
		var _CustomerID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CreatedAt,
			&_Type,
			&_Message,
			&_Fatal,
			&_CustomerID,
		)
		if err != nil {
			return nil, err
		}
		t := &ProcessingError{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CreatedAt.Valid {
			t.SetCreatedAt(_CreatedAt.Int64)
		}
		if _Type.Valid {
			t.SetType(_Type.String)
		}
		if _Message.Valid {
			t.SetMessage(_Message.String)
		}
		if _Fatal.Valid {
			t.SetFatal(_Fatal.Bool)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindProcessingErrorsTx will find a ProcessingError record in the database with the provided parameters using the provided transaction
func FindProcessingErrorsTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) ([]*ProcessingError, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("created_at"),
		orm.Column("type"),
		orm.Column("message"),
		orm.Column("fatal"),
		orm.Column("customer_id"),
		orm.Table(ProcessingErrorTableName),
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
	results := make([]*ProcessingError, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CreatedAt sql.NullInt64
		var _Type sql.NullString
		var _Message sql.NullString
		var _Fatal sql.NullBool
		var _CustomerID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CreatedAt,
			&_Type,
			&_Message,
			&_Fatal,
			&_CustomerID,
		)
		if err != nil {
			return nil, err
		}
		t := &ProcessingError{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CreatedAt.Valid {
			t.SetCreatedAt(_CreatedAt.Int64)
		}
		if _Type.Valid {
			t.SetType(_Type.String)
		}
		if _Message.Valid {
			t.SetMessage(_Message.String)
		}
		if _Fatal.Valid {
			t.SetFatal(_Fatal.Bool)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// DBFind will find a ProcessingError record in the database with the provided parameters
func (t *ProcessingError) DBFind(ctx context.Context, db *sql.DB, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("created_at"),
		orm.Column("type"),
		orm.Column("message"),
		orm.Column("fatal"),
		orm.Column("customer_id"),
		orm.Table(ProcessingErrorTableName),
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
	var _CreatedAt sql.NullInt64
	var _Type sql.NullString
	var _Message sql.NullString
	var _Fatal sql.NullBool
	var _CustomerID sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CreatedAt,
		&_Type,
		&_Message,
		&_Fatal,
		&_CustomerID,
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
	if _CreatedAt.Valid {
		t.SetCreatedAt(_CreatedAt.Int64)
	}
	if _Type.Valid {
		t.SetType(_Type.String)
	}
	if _Message.Valid {
		t.SetMessage(_Message.String)
	}
	if _Fatal.Valid {
		t.SetFatal(_Fatal.Bool)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	return true, nil
}

// DBFindTx will find a ProcessingError record in the database with the provided parameters using the provided transaction
func (t *ProcessingError) DBFindTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("created_at"),
		orm.Column("type"),
		orm.Column("message"),
		orm.Column("fatal"),
		orm.Column("customer_id"),
		orm.Table(ProcessingErrorTableName),
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
	var _CreatedAt sql.NullInt64
	var _Type sql.NullString
	var _Message sql.NullString
	var _Fatal sql.NullBool
	var _CustomerID sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CreatedAt,
		&_Type,
		&_Message,
		&_Fatal,
		&_CustomerID,
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
	if _CreatedAt.Valid {
		t.SetCreatedAt(_CreatedAt.Int64)
	}
	if _Type.Valid {
		t.SetType(_Type.String)
	}
	if _Message.Valid {
		t.SetMessage(_Message.String)
	}
	if _Fatal.Valid {
		t.SetFatal(_Fatal.Bool)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	return true, nil
}

// CountProcessingErrors will find the count of ProcessingError records in the database
func CountProcessingErrors(ctx context.Context, db *sql.DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(ProcessingErrorTableName),
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

// CountProcessingErrorsTx will find the count of ProcessingError records in the database using the provided transaction
func CountProcessingErrorsTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(ProcessingErrorTableName),
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

// DBCount will find the count of ProcessingError records in the database
func (t *ProcessingError) DBCount(ctx context.Context, db *sql.DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(ProcessingErrorTableName),
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

// DBCountTx will find the count of ProcessingError records in the database using the provided transaction
func (t *ProcessingError) DBCountTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(ProcessingErrorTableName),
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

// DBExists will return true if the ProcessingError record exists in the database
func (t *ProcessingError) DBExists(ctx context.Context, db *sql.DB) (bool, error) {
	q := "SELECT `id` FROM `processing_error` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := db.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// DBExistsTx will return true if the ProcessingError record exists in the database using the provided transaction
func (t *ProcessingError) DBExistsTx(ctx context.Context, tx *sql.Tx) (bool, error) {
	q := "SELECT `id` FROM `processing_error` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := tx.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// PrimaryKeyColumn returns the column name for the primary key
func (t *ProcessingError) PrimaryKeyColumn() string {
	return ProcessingErrorColumnID
}

// PrimaryKeyColumnType returns the primary key column Go type as a string
func (t *ProcessingError) PrimaryKeyColumnType() string {
	return "string"
}

// PrimaryKey returns the primary key column value
func (t *ProcessingError) PrimaryKey() interface{} {
	return t.ID
}
