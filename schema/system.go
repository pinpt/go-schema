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

var _ Model = (*System)(nil)
var _ CSVWriter = (*System)(nil)
var _ JSONWriter = (*System)(nil)
var _ Checksum = (*System)(nil)

// SystemTableName is the name of the table in SQL
const SystemTableName = "system"

var SystemColumns = []string{
	"id",
	"checksum",
	"repo_id",
	"ref_id",
	"ref_type",
	"customer_id",
}

// System table
type System struct {
	Checksum   *string `json:"checksum,omitempty"`
	CustomerID string  `json:"customer_id"`
	ID         string  `json:"id"`
	RefID      string  `json:"ref_id"`
	RefType    string  `json:"ref_type"`
	RepoID     string  `json:"repo_id"`
}

// TableName returns the SQL table name for System and satifies the Model interface
func (t *System) TableName() string {
	return SystemTableName
}

// ToCSV will serialize the System instance to a CSV compatible array of strings
func (t *System) ToCSV() []string {
	return []string{
		t.ID,
		t.CalculateChecksum(),
		t.RepoID,
		t.RefID,
		t.RefType,
		t.CustomerID,
	}
}

// WriteCSV will serialize the System instance to the writer as CSV and satisfies the CSVWriter interface
func (t *System) WriteCSV(w *csv.Writer) error {
	return w.Write(t.ToCSV())
}

// WriteJSON will serialize the System instance to the writer as JSON and satisfies the JSONWriter interface
func (t *System) WriteJSON(w io.Writer, indent ...bool) error {
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

// NewSystemReader creates a JSON reader which can read in System objects serialized as JSON either as an array, single object or json new lines
// and writes each System to the channel provided
func NewSystemReader(r io.Reader, ch chan<- System) error {
	return orm.Deserialize(r, func(buf json.RawMessage) error {
		dec := json.NewDecoder(bytes.NewBuffer(buf))
		e := System{}
		if err := dec.Decode(&e); err != nil {
			return err
		}
		ch <- e
		return nil
	})
}

// NewCSVSystemReaderDir will read the reader as CSV and emit each record to the channel provided
func NewCSVSystemReader(r io.Reader, ch chan<- System) error {
	cr := csv.NewReader(r)
	for {
		record, err := cr.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		ch <- System{
			ID:         record[0],
			Checksum:   fromStringPointer(record[1]),
			RepoID:     record[2],
			RefID:      record[3],
			RefType:    record[4],
			CustomerID: record[5],
		}
	}
	return nil
}

// NewCSVSystemReaderFile will read the file as a CSV and emit each record to the channel provided
func NewCSVSystemReaderFile(fp string, ch chan<- System) error {
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
	return NewCSVSystemReader(fc, ch)
}

// NewCSVSystemReaderDir will read the system.csv.gz file as a CSV and emit each record to the channel provided
func NewCSVSystemReaderDir(dir string, ch chan<- System) error {
	return NewCSVSystemReaderFile(filepath.Join(dir, "system.csv.gz"), ch)
}

// SystemCSVDeduper is a function callback which takes the existing value (a) and the new value (b)
// and the return value should be the one to use (or a new one, if applicable). return nil
// to skip processing of this record
type SystemCSVDeduper func(a System, b System) *System

// SystemCSVDedupeDisabled is set on whether the CSV writer should de-dupe values by key
var SystemCSVDedupeDisabled bool

// NewSystemCSVWriterSize creates a batch writer that will write each System into a CSV file
// this method will automatically de-duplicate entries using the primary key. if the checksum
// for a newer item with the same primary key doesn't match a previously sent item, the newer
// one will be used
func NewSystemCSVWriterSize(w io.Writer, size int, dedupers ...SystemCSVDeduper) (chan System, chan bool, error) {
	cw := csv.NewWriter(w)
	ch := make(chan System, size)
	done := make(chan bool)
	go func() {
		defer func() { done <- true }()
		dodedupe := !SystemCSVDedupeDisabled
		var kv map[string]*System
		var deduper SystemCSVDeduper
		if dedupers != nil && len(dedupers) > 0 {
			deduper = dedupers[0]
			dodedupe = true
		}
		if dodedupe {
			kv = make(map[string]*System)
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

// SystemCSVDefaultSize is the default channel buffer size if not provided
var SystemCSVDefaultSize = 100

// NewSystemCSVWriter creates a batch writer that will write each System into a CSV file
func NewSystemCSVWriter(w io.Writer, dedupers ...SystemCSVDeduper) (chan System, chan bool, error) {
	return NewSystemCSVWriterSize(w, SystemCSVDefaultSize, dedupers...)
}

// NewSystemCSVWriterDir creates a batch writer that will write each System into a CSV file named system.csv.gz in dir
func NewSystemCSVWriterDir(dir string, dedupers ...SystemCSVDeduper) (chan System, chan bool, error) {
	return NewSystemCSVWriterFile(filepath.Join(dir, "system.csv.gz"), dedupers...)
}

// NewSystemCSVWriterFile creates a batch writer that will write each System into a CSV file
func NewSystemCSVWriterFile(fn string, dedupers ...SystemCSVDeduper) (chan System, chan bool, error) {
	f, err := os.Create(fn)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening CSV file %s. %v", fn, err)
	}
	var fc io.WriteCloser = f
	if filepath.Ext(fn) == ".gz" {
		w, _ := gzip.NewWriterLevel(f, gzip.BestCompression)
		fc = w
	}
	ch, done, err := NewSystemCSVWriter(fc, dedupers...)
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

type SystemDBAction func(ctx context.Context, db DB, record System) error

// NewSystemDBWriterSize creates a DB writer that will write each issue into the DB
func NewSystemDBWriterSize(ctx context.Context, db DB, errors chan<- error, size int, actions ...SystemDBAction) (chan System, chan bool, error) {
	ch := make(chan System, size)
	done := make(chan bool)
	var action SystemDBAction
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

// NewSystemDBWriter creates a DB writer that will write each issue into the DB
func NewSystemDBWriter(ctx context.Context, db DB, errors chan<- error, actions ...SystemDBAction) (chan System, chan bool, error) {
	return NewSystemDBWriterSize(ctx, db, errors, 100, actions...)
}

// SystemColumnID is the ID SQL column name for the System table
const SystemColumnID = "id"

// SystemEscapedColumnID is the escaped ID SQL column name for the System table
const SystemEscapedColumnID = "`id`"

// SystemColumnChecksum is the Checksum SQL column name for the System table
const SystemColumnChecksum = "checksum"

// SystemEscapedColumnChecksum is the escaped Checksum SQL column name for the System table
const SystemEscapedColumnChecksum = "`checksum`"

// SystemColumnRepoID is the RepoID SQL column name for the System table
const SystemColumnRepoID = "repo_id"

// SystemEscapedColumnRepoID is the escaped RepoID SQL column name for the System table
const SystemEscapedColumnRepoID = "`repo_id`"

// SystemColumnRefID is the RefID SQL column name for the System table
const SystemColumnRefID = "ref_id"

// SystemEscapedColumnRefID is the escaped RefID SQL column name for the System table
const SystemEscapedColumnRefID = "`ref_id`"

// SystemColumnRefType is the RefType SQL column name for the System table
const SystemColumnRefType = "ref_type"

// SystemEscapedColumnRefType is the escaped RefType SQL column name for the System table
const SystemEscapedColumnRefType = "`ref_type`"

// SystemColumnCustomerID is the CustomerID SQL column name for the System table
const SystemColumnCustomerID = "customer_id"

// SystemEscapedColumnCustomerID is the escaped CustomerID SQL column name for the System table
const SystemEscapedColumnCustomerID = "`customer_id`"

// GetID will return the System ID value
func (t *System) GetID() string {
	return t.ID
}

// SetID will set the System ID value
func (t *System) SetID(v string) {
	t.ID = v
}

// FindSystemByID will find a System by ID
func FindSystemByID(ctx context.Context, db DB, value string) (*System, error) {
	q := "SELECT `system`.`id`,`system`.`checksum`,`system`.`repo_id`,`system`.`ref_id`,`system`.`ref_type`,`system`.`customer_id` FROM `system` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _RepoID sql.NullString
	var _RefID sql.NullString
	var _RefType sql.NullString
	var _CustomerID sql.NullString
	err := db.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_RepoID,
		&_RefID,
		&_RefType,
		&_CustomerID,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &System{}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _Checksum.Valid {
		t.SetChecksum(_Checksum.String)
	}
	if _RepoID.Valid {
		t.SetRepoID(_RepoID.String)
	}
	if _RefID.Valid {
		t.SetRefID(_RefID.String)
	}
	if _RefType.Valid {
		t.SetRefType(_RefType.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	return t, nil
}

// FindSystemByIDTx will find a System by ID using the provided transaction
func FindSystemByIDTx(ctx context.Context, tx Tx, value string) (*System, error) {
	q := "SELECT `system`.`id`,`system`.`checksum`,`system`.`repo_id`,`system`.`ref_id`,`system`.`ref_type`,`system`.`customer_id` FROM `system` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _RepoID sql.NullString
	var _RefID sql.NullString
	var _RefType sql.NullString
	var _CustomerID sql.NullString
	err := tx.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_RepoID,
		&_RefID,
		&_RefType,
		&_CustomerID,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &System{}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _Checksum.Valid {
		t.SetChecksum(_Checksum.String)
	}
	if _RepoID.Valid {
		t.SetRepoID(_RepoID.String)
	}
	if _RefID.Valid {
		t.SetRefID(_RefID.String)
	}
	if _RefType.Valid {
		t.SetRefType(_RefType.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	return t, nil
}

// GetChecksum will return the System Checksum value
func (t *System) GetChecksum() string {
	if t.Checksum == nil {
		return ""
	}
	return *t.Checksum
}

// SetChecksum will set the System Checksum value
func (t *System) SetChecksum(v string) {
	t.Checksum = &v
}

// GetRepoID will return the System RepoID value
func (t *System) GetRepoID() string {
	return t.RepoID
}

// SetRepoID will set the System RepoID value
func (t *System) SetRepoID(v string) {
	t.RepoID = v
}

// FindSystemsByRepoID will find all Systems by the RepoID value
func FindSystemsByRepoID(ctx context.Context, db DB, value string) ([]*System, error) {
	q := "SELECT `system`.`id`,`system`.`checksum`,`system`.`repo_id`,`system`.`ref_id`,`system`.`ref_type`,`system`.`customer_id` FROM `system` WHERE `repo_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*System, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _RepoID sql.NullString
		var _RefID sql.NullString
		var _RefType sql.NullString
		var _CustomerID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_RepoID,
			&_RefID,
			&_RefType,
			&_CustomerID,
		)
		if err != nil {
			return nil, err
		}
		t := &System{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _RepoID.Valid {
			t.SetRepoID(_RepoID.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		if _RefType.Valid {
			t.SetRefType(_RefType.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindSystemsByRepoIDTx will find all Systems by the RepoID value using the provided transaction
func FindSystemsByRepoIDTx(ctx context.Context, tx Tx, value string) ([]*System, error) {
	q := "SELECT `system`.`id`,`system`.`checksum`,`system`.`repo_id`,`system`.`ref_id`,`system`.`ref_type`,`system`.`customer_id` FROM `system` WHERE `repo_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*System, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _RepoID sql.NullString
		var _RefID sql.NullString
		var _RefType sql.NullString
		var _CustomerID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_RepoID,
			&_RefID,
			&_RefType,
			&_CustomerID,
		)
		if err != nil {
			return nil, err
		}
		t := &System{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _RepoID.Valid {
			t.SetRepoID(_RepoID.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		if _RefType.Valid {
			t.SetRefType(_RefType.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetRefID will return the System RefID value
func (t *System) GetRefID() string {
	return t.RefID
}

// SetRefID will set the System RefID value
func (t *System) SetRefID(v string) {
	t.RefID = v
}

// GetRefType will return the System RefType value
func (t *System) GetRefType() string {
	return t.RefType
}

// SetRefType will set the System RefType value
func (t *System) SetRefType(v string) {
	t.RefType = v
}

// GetCustomerID will return the System CustomerID value
func (t *System) GetCustomerID() string {
	return t.CustomerID
}

// SetCustomerID will set the System CustomerID value
func (t *System) SetCustomerID(v string) {
	t.CustomerID = v
}

// FindSystemsByCustomerID will find all Systems by the CustomerID value
func FindSystemsByCustomerID(ctx context.Context, db DB, value string) ([]*System, error) {
	q := "SELECT `system`.`id`,`system`.`checksum`,`system`.`repo_id`,`system`.`ref_id`,`system`.`ref_type`,`system`.`customer_id` FROM `system` WHERE `customer_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*System, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _RepoID sql.NullString
		var _RefID sql.NullString
		var _RefType sql.NullString
		var _CustomerID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_RepoID,
			&_RefID,
			&_RefType,
			&_CustomerID,
		)
		if err != nil {
			return nil, err
		}
		t := &System{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _RepoID.Valid {
			t.SetRepoID(_RepoID.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		if _RefType.Valid {
			t.SetRefType(_RefType.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindSystemsByCustomerIDTx will find all Systems by the CustomerID value using the provided transaction
func FindSystemsByCustomerIDTx(ctx context.Context, tx Tx, value string) ([]*System, error) {
	q := "SELECT `system`.`id`,`system`.`checksum`,`system`.`repo_id`,`system`.`ref_id`,`system`.`ref_type`,`system`.`customer_id` FROM `system` WHERE `customer_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*System, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _RepoID sql.NullString
		var _RefID sql.NullString
		var _RefType sql.NullString
		var _CustomerID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_RepoID,
			&_RefID,
			&_RefType,
			&_CustomerID,
		)
		if err != nil {
			return nil, err
		}
		t := &System{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _RepoID.Valid {
			t.SetRepoID(_RepoID.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		if _RefType.Valid {
			t.SetRefType(_RefType.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

func (t *System) toTimestamp(value time.Time) *timestamp.Timestamp {
	ts, _ := ptypes.TimestampProto(value)
	return ts
}

// DBCreateSystemTable will create the System table
func DBCreateSystemTable(ctx context.Context, db DB) error {
	q := "CREATE TABLE `system` (`id` VARCHAR(64) NOT NULL PRIMARY KEY,`checksum` CHAR(64),`repo_id`VARCHAR(64) NOT NULL,`ref_id` VARCHAR(64) NOT NULL,`ref_type` VARCHAR(20) NOT NULL,`customer_id` VARCHAR(64) NOT NULL,INDEX system_repo_id_index (`repo_id`),INDEX system_customer_id_index (`customer_id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBCreateSystemTableTx will create the System table using the provided transction
func DBCreateSystemTableTx(ctx context.Context, tx Tx) error {
	q := "CREATE TABLE `system` (`id` VARCHAR(64) NOT NULL PRIMARY KEY,`checksum` CHAR(64),`repo_id`VARCHAR(64) NOT NULL,`ref_id` VARCHAR(64) NOT NULL,`ref_type` VARCHAR(20) NOT NULL,`customer_id` VARCHAR(64) NOT NULL,INDEX system_repo_id_index (`repo_id`),INDEX system_customer_id_index (`customer_id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// DBDropSystemTable will drop the System table
func DBDropSystemTable(ctx context.Context, db DB) error {
	q := "DROP TABLE IF EXISTS `system`"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBDropSystemTableTx will drop the System table using the provided transaction
func DBDropSystemTableTx(ctx context.Context, tx Tx) error {
	q := "DROP TABLE IF EXISTS `system`"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// CalculateChecksum will calculate a checksum of the SHA1 of all field values
func (t *System) CalculateChecksum() string {
	return orm.HashStrings(
		orm.ToString(t.ID),
		orm.ToString(t.RepoID),
		orm.ToString(t.RefID),
		orm.ToString(t.RefType),
		orm.ToString(t.CustomerID),
	)
}

// DBCreate will create a new System record in the database
func (t *System) DBCreate(ctx context.Context, db DB) (sql.Result, error) {
	q := "INSERT INTO `system` (`system`.`id`,`system`.`checksum`,`system`.`repo_id`,`system`.`ref_id`,`system`.`ref_type`,`system`.`customer_id`) VALUES (?,?,?,?,?,?)"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.RepoID),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.RefType),
		orm.ToSQLString(t.CustomerID),
	)
}

// DBCreateTx will create a new System record in the database using the provided transaction
func (t *System) DBCreateTx(ctx context.Context, tx Tx) (sql.Result, error) {
	q := "INSERT INTO `system` (`system`.`id`,`system`.`checksum`,`system`.`repo_id`,`system`.`ref_id`,`system`.`ref_type`,`system`.`customer_id`) VALUES (?,?,?,?,?,?)"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.RepoID),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.RefType),
		orm.ToSQLString(t.CustomerID),
	)
}

// DBCreateIgnoreDuplicate will upsert the System record in the database
func (t *System) DBCreateIgnoreDuplicate(ctx context.Context, db DB) (sql.Result, error) {
	q := "INSERT INTO `system` (`system`.`id`,`system`.`checksum`,`system`.`repo_id`,`system`.`ref_id`,`system`.`ref_type`,`system`.`customer_id`) VALUES (?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.RepoID),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.RefType),
		orm.ToSQLString(t.CustomerID),
	)
}

// DBCreateIgnoreDuplicateTx will upsert the System record in the database using the provided transaction
func (t *System) DBCreateIgnoreDuplicateTx(ctx context.Context, tx Tx) (sql.Result, error) {
	q := "INSERT INTO `system` (`system`.`id`,`system`.`checksum`,`system`.`repo_id`,`system`.`ref_id`,`system`.`ref_type`,`system`.`customer_id`) VALUES (?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.RepoID),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.RefType),
		orm.ToSQLString(t.CustomerID),
	)
}

// DeleteAllSystems deletes all System records in the database with optional filters
func DeleteAllSystems(ctx context.Context, db DB, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(SystemTableName),
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

// DeleteAllSystemsTx deletes all System records in the database with optional filters using the provided transaction
func DeleteAllSystemsTx(ctx context.Context, tx Tx, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(SystemTableName),
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

// DBDelete will delete this System record in the database
func (t *System) DBDelete(ctx context.Context, db DB) (bool, error) {
	q := "DELETE FROM `system` WHERE `id` = ?"
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

// DBDeleteTx will delete this System record in the database using the provided transaction
func (t *System) DBDeleteTx(ctx context.Context, tx Tx) (bool, error) {
	q := "DELETE FROM `system` WHERE `id` = ?"
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

// DBUpdate will update the System record in the database
func (t *System) DBUpdate(ctx context.Context, db DB) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `system` SET `checksum`=?,`repo_id`=?,`ref_id`=?,`ref_type`=?,`customer_id`=? WHERE `id`=?"
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.RepoID),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.RefType),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.ID),
	)
}

// DBUpdateTx will update the System record in the database using the provided transaction
func (t *System) DBUpdateTx(ctx context.Context, tx Tx) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `system` SET `checksum`=?,`repo_id`=?,`ref_id`=?,`ref_type`=?,`customer_id`=? WHERE `id`=?"
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.RepoID),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.RefType),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.ID),
	)
}

// DBUpsert will upsert the System record in the database
func (t *System) DBUpsert(ctx context.Context, db DB, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `system` (`system`.`id`,`system`.`checksum`,`system`.`repo_id`,`system`.`ref_id`,`system`.`ref_type`,`system`.`customer_id`) VALUES (?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `system` (`system`.`id`,`system`.`checksum`,`system`.`repo_id`,`system`.`ref_id`,`system`.`ref_type`,`system`.`customer_id`) VALUES (?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`repo_id`=VALUES(`repo_id`),`ref_id`=VALUES(`ref_id`),`ref_type`=VALUES(`ref_type`),`customer_id`=VALUES(`customer_id`)"
	}
	r, err := db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.RepoID),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.RefType),
		orm.ToSQLString(t.CustomerID),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBUpsertTx will upsert the System record in the database using the provided transaction
func (t *System) DBUpsertTx(ctx context.Context, tx Tx, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `system` (`system`.`id`,`system`.`checksum`,`system`.`repo_id`,`system`.`ref_id`,`system`.`ref_type`,`system`.`customer_id`) VALUES (?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `system` (`system`.`id`,`system`.`checksum`,`system`.`repo_id`,`system`.`ref_id`,`system`.`ref_type`,`system`.`customer_id`) VALUES (?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`repo_id`=VALUES(`repo_id`),`ref_id`=VALUES(`ref_id`),`ref_type`=VALUES(`ref_type`),`customer_id`=VALUES(`customer_id`)"
	}
	r, err := tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.RepoID),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.RefType),
		orm.ToSQLString(t.CustomerID),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBFindOne will find a System record in the database with the primary key
func (t *System) DBFindOne(ctx context.Context, db DB, value string) (bool, error) {
	q := "SELECT `system`.`id`,`system`.`checksum`,`system`.`repo_id`,`system`.`ref_id`,`system`.`ref_type`,`system`.`customer_id` FROM `system` WHERE `id` = ? LIMIT 1"
	row := db.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _RepoID sql.NullString
	var _RefID sql.NullString
	var _RefType sql.NullString
	var _CustomerID sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_RepoID,
		&_RefID,
		&_RefType,
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
	if _RepoID.Valid {
		t.SetRepoID(_RepoID.String)
	}
	if _RefID.Valid {
		t.SetRefID(_RefID.String)
	}
	if _RefType.Valid {
		t.SetRefType(_RefType.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	return true, nil
}

// DBFindOneTx will find a System record in the database with the primary key using the provided transaction
func (t *System) DBFindOneTx(ctx context.Context, tx Tx, value string) (bool, error) {
	q := "SELECT `system`.`id`,`system`.`checksum`,`system`.`repo_id`,`system`.`ref_id`,`system`.`ref_type`,`system`.`customer_id` FROM `system` WHERE `id` = ? LIMIT 1"
	row := tx.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _RepoID sql.NullString
	var _RefID sql.NullString
	var _RefType sql.NullString
	var _CustomerID sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_RepoID,
		&_RefID,
		&_RefType,
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
	if _RepoID.Valid {
		t.SetRepoID(_RepoID.String)
	}
	if _RefID.Valid {
		t.SetRefID(_RefID.String)
	}
	if _RefType.Valid {
		t.SetRefType(_RefType.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	return true, nil
}

// FindSystems will find a System record in the database with the provided parameters
func FindSystems(ctx context.Context, db DB, _params ...interface{}) ([]*System, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("repo_id"),
		orm.Column("ref_id"),
		orm.Column("ref_type"),
		orm.Column("customer_id"),
		orm.Table(SystemTableName),
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
	results := make([]*System, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _RepoID sql.NullString
		var _RefID sql.NullString
		var _RefType sql.NullString
		var _CustomerID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_RepoID,
			&_RefID,
			&_RefType,
			&_CustomerID,
		)
		if err != nil {
			return nil, err
		}
		t := &System{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _RepoID.Valid {
			t.SetRepoID(_RepoID.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		if _RefType.Valid {
			t.SetRefType(_RefType.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindSystemsTx will find a System record in the database with the provided parameters using the provided transaction
func FindSystemsTx(ctx context.Context, tx Tx, _params ...interface{}) ([]*System, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("repo_id"),
		orm.Column("ref_id"),
		orm.Column("ref_type"),
		orm.Column("customer_id"),
		orm.Table(SystemTableName),
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
	results := make([]*System, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _RepoID sql.NullString
		var _RefID sql.NullString
		var _RefType sql.NullString
		var _CustomerID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_RepoID,
			&_RefID,
			&_RefType,
			&_CustomerID,
		)
		if err != nil {
			return nil, err
		}
		t := &System{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _RepoID.Valid {
			t.SetRepoID(_RepoID.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		if _RefType.Valid {
			t.SetRefType(_RefType.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// DBFind will find a System record in the database with the provided parameters
func (t *System) DBFind(ctx context.Context, db DB, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("repo_id"),
		orm.Column("ref_id"),
		orm.Column("ref_type"),
		orm.Column("customer_id"),
		orm.Table(SystemTableName),
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
	var _RepoID sql.NullString
	var _RefID sql.NullString
	var _RefType sql.NullString
	var _CustomerID sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_RepoID,
		&_RefID,
		&_RefType,
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
	if _RepoID.Valid {
		t.SetRepoID(_RepoID.String)
	}
	if _RefID.Valid {
		t.SetRefID(_RefID.String)
	}
	if _RefType.Valid {
		t.SetRefType(_RefType.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	return true, nil
}

// DBFindTx will find a System record in the database with the provided parameters using the provided transaction
func (t *System) DBFindTx(ctx context.Context, tx Tx, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("repo_id"),
		orm.Column("ref_id"),
		orm.Column("ref_type"),
		orm.Column("customer_id"),
		orm.Table(SystemTableName),
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
	var _RepoID sql.NullString
	var _RefID sql.NullString
	var _RefType sql.NullString
	var _CustomerID sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_RepoID,
		&_RefID,
		&_RefType,
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
	if _RepoID.Valid {
		t.SetRepoID(_RepoID.String)
	}
	if _RefID.Valid {
		t.SetRefID(_RefID.String)
	}
	if _RefType.Valid {
		t.SetRefType(_RefType.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	return true, nil
}

// CountSystems will find the count of System records in the database
func CountSystems(ctx context.Context, db DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(SystemTableName),
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

// CountSystemsTx will find the count of System records in the database using the provided transaction
func CountSystemsTx(ctx context.Context, tx Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(SystemTableName),
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

// DBCount will find the count of System records in the database
func (t *System) DBCount(ctx context.Context, db DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(SystemTableName),
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

// DBCountTx will find the count of System records in the database using the provided transaction
func (t *System) DBCountTx(ctx context.Context, tx Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(SystemTableName),
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

// DBExists will return true if the System record exists in the database
func (t *System) DBExists(ctx context.Context, db DB) (bool, error) {
	q := "SELECT `id` FROM `system` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := db.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// DBExistsTx will return true if the System record exists in the database using the provided transaction
func (t *System) DBExistsTx(ctx context.Context, tx Tx) (bool, error) {
	q := "SELECT `id` FROM `system` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := tx.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// PrimaryKeyColumn returns the column name for the primary key
func (t *System) PrimaryKeyColumn() string {
	return SystemColumnID
}

// PrimaryKeyColumnType returns the primary key column Go type as a string
func (t *System) PrimaryKeyColumnType() string {
	return "string"
}

// PrimaryKey returns the primary key column value
func (t *System) PrimaryKey() interface{} {
	return t.ID
}
