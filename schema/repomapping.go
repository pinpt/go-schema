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

var _ Model = (*RepoMapping)(nil)
var _ CSVWriter = (*RepoMapping)(nil)
var _ JSONWriter = (*RepoMapping)(nil)
var _ Checksum = (*RepoMapping)(nil)

// RepoMappingTableName is the name of the table in SQL
const RepoMappingTableName = "repo_mapping"

var RepoMappingColumns = []string{
	"id",
	"checksum",
	"customer_id",
	"repo_id",
	"ref_id",
	"ref_type",
}

// RepoMapping table
type RepoMapping struct {
	Checksum   *string `json:"checksum,omitempty"`
	CustomerID string  `json:"customer_id"`
	ID         string  `json:"id"`
	RefID      string  `json:"ref_id"`
	RefType    string  `json:"ref_type"`
	RepoID     string  `json:"repo_id"`
}

// TableName returns the SQL table name for RepoMapping and satifies the Model interface
func (t *RepoMapping) TableName() string {
	return RepoMappingTableName
}

// ToCSV will serialize the RepoMapping instance to a CSV compatible array of strings
func (t *RepoMapping) ToCSV() []string {
	return []string{
		t.ID,
		t.CalculateChecksum(),
		t.CustomerID,
		t.RepoID,
		t.RefID,
		t.RefType,
	}
}

// WriteCSV will serialize the RepoMapping instance to the writer as CSV and satisfies the CSVWriter interface
func (t *RepoMapping) WriteCSV(w *csv.Writer) error {
	return w.Write(t.ToCSV())
}

// WriteJSON will serialize the RepoMapping instance to the writer as JSON and satisfies the JSONWriter interface
func (t *RepoMapping) WriteJSON(w io.Writer, indent ...bool) error {
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

// NewRepoMappingReader creates a JSON reader which can read in RepoMapping objects serialized as JSON either as an array, single object or json new lines
// and writes each RepoMapping to the channel provided
func NewRepoMappingReader(r io.Reader, ch chan<- RepoMapping) error {
	return orm.Deserialize(r, func(buf json.RawMessage) error {
		dec := json.NewDecoder(bytes.NewBuffer(buf))
		e := RepoMapping{}
		if err := dec.Decode(&e); err != nil {
			return err
		}
		ch <- e
		return nil
	})
}

// NewCSVRepoMappingReaderDir will read the reader as CSV and emit each record to the channel provided
func NewCSVRepoMappingReader(r io.Reader, ch chan<- RepoMapping) error {
	cr := csv.NewReader(r)
	for {
		record, err := cr.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		ch <- RepoMapping{
			ID:         record[0],
			Checksum:   fromStringPointer(record[1]),
			CustomerID: record[2],
			RepoID:     record[3],
			RefID:      record[4],
			RefType:    record[5],
		}
	}
	return nil
}

// NewCSVRepoMappingReaderFile will read the file as a CSV and emit each record to the channel provided
func NewCSVRepoMappingReaderFile(fp string, ch chan<- RepoMapping) error {
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
	return NewCSVRepoMappingReader(fc, ch)
}

// NewCSVRepoMappingReaderDir will read the repo_mapping.csv.gz file as a CSV and emit each record to the channel provided
func NewCSVRepoMappingReaderDir(dir string, ch chan<- RepoMapping) error {
	return NewCSVRepoMappingReaderFile(filepath.Join(dir, "repo_mapping.csv.gz"), ch)
}

// RepoMappingCSVDeduper is a function callback which takes the existing value (a) and the new value (b)
// and the return value should be the one to use (or a new one, if applicable). return nil
// to skip processing of this record
type RepoMappingCSVDeduper func(a RepoMapping, b RepoMapping) *RepoMapping

// RepoMappingCSVDedupeDisabled is set on whether the CSV writer should de-dupe values by key
var RepoMappingCSVDedupeDisabled bool

// NewRepoMappingCSVWriterSize creates a batch writer that will write each RepoMapping into a CSV file
// this method will automatically de-duplicate entries using the primary key. if the checksum
// for a newer item with the same primary key doesn't match a previously sent item, the newer
// one will be used
func NewRepoMappingCSVWriterSize(w io.Writer, size int, dedupers ...RepoMappingCSVDeduper) (chan RepoMapping, chan bool, error) {
	cw := csv.NewWriter(w)
	ch := make(chan RepoMapping, size)
	done := make(chan bool)
	go func() {
		defer func() { done <- true }()
		dodedupe := !RepoMappingCSVDedupeDisabled
		var kv map[string]*RepoMapping
		var deduper RepoMappingCSVDeduper
		if dedupers != nil && len(dedupers) > 0 {
			deduper = dedupers[0]
			dodedupe = true
		}
		if dodedupe {
			kv = make(map[string]*RepoMapping)
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

// RepoMappingCSVDefaultSize is the default channel buffer size if not provided
var RepoMappingCSVDefaultSize = 100

// NewRepoMappingCSVWriter creates a batch writer that will write each RepoMapping into a CSV file
func NewRepoMappingCSVWriter(w io.Writer, dedupers ...RepoMappingCSVDeduper) (chan RepoMapping, chan bool, error) {
	return NewRepoMappingCSVWriterSize(w, RepoMappingCSVDefaultSize, dedupers...)
}

// NewRepoMappingCSVWriterDir creates a batch writer that will write each RepoMapping into a CSV file named repo_mapping.csv.gz in dir
func NewRepoMappingCSVWriterDir(dir string, dedupers ...RepoMappingCSVDeduper) (chan RepoMapping, chan bool, error) {
	return NewRepoMappingCSVWriterFile(filepath.Join(dir, "repo_mapping.csv.gz"), dedupers...)
}

// NewRepoMappingCSVWriterFile creates a batch writer that will write each RepoMapping into a CSV file
func NewRepoMappingCSVWriterFile(fn string, dedupers ...RepoMappingCSVDeduper) (chan RepoMapping, chan bool, error) {
	f, err := os.Create(fn)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening CSV file %s. %v", fn, err)
	}
	var fc io.WriteCloser = f
	if filepath.Ext(fn) == ".gz" {
		w, _ := gzip.NewWriterLevel(f, gzip.BestCompression)
		fc = w
	}
	ch, done, err := NewRepoMappingCSVWriter(fc, dedupers...)
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

type RepoMappingDBAction func(ctx context.Context, db DB, record RepoMapping) error

// NewRepoMappingDBWriterSize creates a DB writer that will write each issue into the DB
func NewRepoMappingDBWriterSize(ctx context.Context, db DB, errors chan<- error, size int, actions ...RepoMappingDBAction) (chan RepoMapping, chan bool, error) {
	ch := make(chan RepoMapping, size)
	done := make(chan bool)
	var action RepoMappingDBAction
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

// NewRepoMappingDBWriter creates a DB writer that will write each issue into the DB
func NewRepoMappingDBWriter(ctx context.Context, db DB, errors chan<- error, actions ...RepoMappingDBAction) (chan RepoMapping, chan bool, error) {
	return NewRepoMappingDBWriterSize(ctx, db, errors, 100, actions...)
}

// RepoMappingColumnID is the ID SQL column name for the RepoMapping table
const RepoMappingColumnID = "id"

// RepoMappingEscapedColumnID is the escaped ID SQL column name for the RepoMapping table
const RepoMappingEscapedColumnID = "`id`"

// RepoMappingColumnChecksum is the Checksum SQL column name for the RepoMapping table
const RepoMappingColumnChecksum = "checksum"

// RepoMappingEscapedColumnChecksum is the escaped Checksum SQL column name for the RepoMapping table
const RepoMappingEscapedColumnChecksum = "`checksum`"

// RepoMappingColumnCustomerID is the CustomerID SQL column name for the RepoMapping table
const RepoMappingColumnCustomerID = "customer_id"

// RepoMappingEscapedColumnCustomerID is the escaped CustomerID SQL column name for the RepoMapping table
const RepoMappingEscapedColumnCustomerID = "`customer_id`"

// RepoMappingColumnRepoID is the RepoID SQL column name for the RepoMapping table
const RepoMappingColumnRepoID = "repo_id"

// RepoMappingEscapedColumnRepoID is the escaped RepoID SQL column name for the RepoMapping table
const RepoMappingEscapedColumnRepoID = "`repo_id`"

// RepoMappingColumnRefID is the RefID SQL column name for the RepoMapping table
const RepoMappingColumnRefID = "ref_id"

// RepoMappingEscapedColumnRefID is the escaped RefID SQL column name for the RepoMapping table
const RepoMappingEscapedColumnRefID = "`ref_id`"

// RepoMappingColumnRefType is the RefType SQL column name for the RepoMapping table
const RepoMappingColumnRefType = "ref_type"

// RepoMappingEscapedColumnRefType is the escaped RefType SQL column name for the RepoMapping table
const RepoMappingEscapedColumnRefType = "`ref_type`"

// GetID will return the RepoMapping ID value
func (t *RepoMapping) GetID() string {
	return t.ID
}

// SetID will set the RepoMapping ID value
func (t *RepoMapping) SetID(v string) {
	t.ID = v
}

// FindRepoMappingByID will find a RepoMapping by ID
func FindRepoMappingByID(ctx context.Context, db DB, value string) (*RepoMapping, error) {
	q := "SELECT * FROM `repo_mapping` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CustomerID sql.NullString
	var _RepoID sql.NullString
	var _RefID sql.NullString
	var _RefType sql.NullString
	err := db.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_RepoID,
		&_RefID,
		&_RefType,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &RepoMapping{}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _Checksum.Valid {
		t.SetChecksum(_Checksum.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
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
	return t, nil
}

// FindRepoMappingByIDTx will find a RepoMapping by ID using the provided transaction
func FindRepoMappingByIDTx(ctx context.Context, tx Tx, value string) (*RepoMapping, error) {
	q := "SELECT * FROM `repo_mapping` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CustomerID sql.NullString
	var _RepoID sql.NullString
	var _RefID sql.NullString
	var _RefType sql.NullString
	err := tx.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_RepoID,
		&_RefID,
		&_RefType,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &RepoMapping{}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _Checksum.Valid {
		t.SetChecksum(_Checksum.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
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
	return t, nil
}

// GetChecksum will return the RepoMapping Checksum value
func (t *RepoMapping) GetChecksum() string {
	if t.Checksum == nil {
		return ""
	}
	return *t.Checksum
}

// SetChecksum will set the RepoMapping Checksum value
func (t *RepoMapping) SetChecksum(v string) {
	t.Checksum = &v
}

// GetCustomerID will return the RepoMapping CustomerID value
func (t *RepoMapping) GetCustomerID() string {
	return t.CustomerID
}

// SetCustomerID will set the RepoMapping CustomerID value
func (t *RepoMapping) SetCustomerID(v string) {
	t.CustomerID = v
}

// GetRepoID will return the RepoMapping RepoID value
func (t *RepoMapping) GetRepoID() string {
	return t.RepoID
}

// SetRepoID will set the RepoMapping RepoID value
func (t *RepoMapping) SetRepoID(v string) {
	t.RepoID = v
}

// FindRepoMappingsByRepoID will find all RepoMappings by the RepoID value
func FindRepoMappingsByRepoID(ctx context.Context, db DB, value string) ([]*RepoMapping, error) {
	q := "SELECT * FROM `repo_mapping` WHERE `repo_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*RepoMapping, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _RepoID sql.NullString
		var _RefID sql.NullString
		var _RefType sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_RepoID,
			&_RefID,
			&_RefType,
		)
		if err != nil {
			return nil, err
		}
		t := &RepoMapping{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
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
		results = append(results, t)
	}
	return results, nil
}

// FindRepoMappingsByRepoIDTx will find all RepoMappings by the RepoID value using the provided transaction
func FindRepoMappingsByRepoIDTx(ctx context.Context, tx Tx, value string) ([]*RepoMapping, error) {
	q := "SELECT * FROM `repo_mapping` WHERE `repo_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*RepoMapping, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _RepoID sql.NullString
		var _RefID sql.NullString
		var _RefType sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_RepoID,
			&_RefID,
			&_RefType,
		)
		if err != nil {
			return nil, err
		}
		t := &RepoMapping{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
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
		results = append(results, t)
	}
	return results, nil
}

// GetRefID will return the RepoMapping RefID value
func (t *RepoMapping) GetRefID() string {
	return t.RefID
}

// SetRefID will set the RepoMapping RefID value
func (t *RepoMapping) SetRefID(v string) {
	t.RefID = v
}

// FindRepoMappingsByRefID will find all RepoMappings by the RefID value
func FindRepoMappingsByRefID(ctx context.Context, db DB, value string) ([]*RepoMapping, error) {
	q := "SELECT * FROM `repo_mapping` WHERE `ref_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*RepoMapping, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _RepoID sql.NullString
		var _RefID sql.NullString
		var _RefType sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_RepoID,
			&_RefID,
			&_RefType,
		)
		if err != nil {
			return nil, err
		}
		t := &RepoMapping{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
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
		results = append(results, t)
	}
	return results, nil
}

// FindRepoMappingsByRefIDTx will find all RepoMappings by the RefID value using the provided transaction
func FindRepoMappingsByRefIDTx(ctx context.Context, tx Tx, value string) ([]*RepoMapping, error) {
	q := "SELECT * FROM `repo_mapping` WHERE `ref_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*RepoMapping, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _RepoID sql.NullString
		var _RefID sql.NullString
		var _RefType sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_RepoID,
			&_RefID,
			&_RefType,
		)
		if err != nil {
			return nil, err
		}
		t := &RepoMapping{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
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
		results = append(results, t)
	}
	return results, nil
}

// GetRefType will return the RepoMapping RefType value
func (t *RepoMapping) GetRefType() string {
	return t.RefType
}

// SetRefType will set the RepoMapping RefType value
func (t *RepoMapping) SetRefType(v string) {
	t.RefType = v
}

func (t *RepoMapping) toTimestamp(value time.Time) *timestamp.Timestamp {
	ts, _ := ptypes.TimestampProto(value)
	return ts
}

// DBCreateRepoMappingTable will create the RepoMapping table
func DBCreateRepoMappingTable(ctx context.Context, db DB) error {
	q := "CREATE TABLE `repo_mapping` (`id` VARCHAR(64) NOT NULL PRIMARY KEY,`checksum` CHAR(64),`customer_id` VARCHAR(64) NOT NULL,`repo_id`VARCHAR(64) NOT NULL,`ref_id` VARCHAR(64) NOT NULL,`ref_type` VARCHAR(20) NOT NULL,INDEX repo_mapping_repo_id_index (`repo_id`),INDEX repo_mapping_ref_id_index (`ref_id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBCreateRepoMappingTableTx will create the RepoMapping table using the provided transction
func DBCreateRepoMappingTableTx(ctx context.Context, tx Tx) error {
	q := "CREATE TABLE `repo_mapping` (`id` VARCHAR(64) NOT NULL PRIMARY KEY,`checksum` CHAR(64),`customer_id` VARCHAR(64) NOT NULL,`repo_id`VARCHAR(64) NOT NULL,`ref_id` VARCHAR(64) NOT NULL,`ref_type` VARCHAR(20) NOT NULL,INDEX repo_mapping_repo_id_index (`repo_id`),INDEX repo_mapping_ref_id_index (`ref_id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// DBDropRepoMappingTable will drop the RepoMapping table
func DBDropRepoMappingTable(ctx context.Context, db DB) error {
	q := "DROP TABLE IF EXISTS `repo_mapping`"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBDropRepoMappingTableTx will drop the RepoMapping table using the provided transaction
func DBDropRepoMappingTableTx(ctx context.Context, tx Tx) error {
	q := "DROP TABLE IF EXISTS `repo_mapping`"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// CalculateChecksum will calculate a checksum of the SHA1 of all field values
func (t *RepoMapping) CalculateChecksum() string {
	return orm.HashStrings(
		orm.ToString(t.ID),
		orm.ToString(t.CustomerID),
		orm.ToString(t.RepoID),
		orm.ToString(t.RefID),
		orm.ToString(t.RefType),
	)
}

// DBCreate will create a new RepoMapping record in the database
func (t *RepoMapping) DBCreate(ctx context.Context, db DB) (sql.Result, error) {
	q := "INSERT INTO `repo_mapping` (`repo_mapping`.`id`,`repo_mapping`.`checksum`,`repo_mapping`.`customer_id`,`repo_mapping`.`repo_id`,`repo_mapping`.`ref_id`,`repo_mapping`.`ref_type`) VALUES (?,?,?,?,?,?)"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RepoID),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.RefType),
	)
}

// DBCreateTx will create a new RepoMapping record in the database using the provided transaction
func (t *RepoMapping) DBCreateTx(ctx context.Context, tx Tx) (sql.Result, error) {
	q := "INSERT INTO `repo_mapping` (`repo_mapping`.`id`,`repo_mapping`.`checksum`,`repo_mapping`.`customer_id`,`repo_mapping`.`repo_id`,`repo_mapping`.`ref_id`,`repo_mapping`.`ref_type`) VALUES (?,?,?,?,?,?)"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RepoID),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.RefType),
	)
}

// DBCreateIgnoreDuplicate will upsert the RepoMapping record in the database
func (t *RepoMapping) DBCreateIgnoreDuplicate(ctx context.Context, db DB) (sql.Result, error) {
	q := "INSERT INTO `repo_mapping` (`repo_mapping`.`id`,`repo_mapping`.`checksum`,`repo_mapping`.`customer_id`,`repo_mapping`.`repo_id`,`repo_mapping`.`ref_id`,`repo_mapping`.`ref_type`) VALUES (?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RepoID),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.RefType),
	)
}

// DBCreateIgnoreDuplicateTx will upsert the RepoMapping record in the database using the provided transaction
func (t *RepoMapping) DBCreateIgnoreDuplicateTx(ctx context.Context, tx Tx) (sql.Result, error) {
	q := "INSERT INTO `repo_mapping` (`repo_mapping`.`id`,`repo_mapping`.`checksum`,`repo_mapping`.`customer_id`,`repo_mapping`.`repo_id`,`repo_mapping`.`ref_id`,`repo_mapping`.`ref_type`) VALUES (?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RepoID),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.RefType),
	)
}

// DeleteAllRepoMappings deletes all RepoMapping records in the database with optional filters
func DeleteAllRepoMappings(ctx context.Context, db DB, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(RepoMappingTableName),
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

// DeleteAllRepoMappingsTx deletes all RepoMapping records in the database with optional filters using the provided transaction
func DeleteAllRepoMappingsTx(ctx context.Context, tx Tx, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(RepoMappingTableName),
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

// DBDelete will delete this RepoMapping record in the database
func (t *RepoMapping) DBDelete(ctx context.Context, db DB) (bool, error) {
	q := "DELETE FROM `repo_mapping` WHERE `id` = ?"
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

// DBDeleteTx will delete this RepoMapping record in the database using the provided transaction
func (t *RepoMapping) DBDeleteTx(ctx context.Context, tx Tx) (bool, error) {
	q := "DELETE FROM `repo_mapping` WHERE `id` = ?"
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

// DBUpdate will update the RepoMapping record in the database
func (t *RepoMapping) DBUpdate(ctx context.Context, db DB) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `repo_mapping` SET `checksum`=?,`customer_id`=?,`repo_id`=?,`ref_id`=?,`ref_type`=? WHERE `id`=?"
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RepoID),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.RefType),
		orm.ToSQLString(t.ID),
	)
}

// DBUpdateTx will update the RepoMapping record in the database using the provided transaction
func (t *RepoMapping) DBUpdateTx(ctx context.Context, tx Tx) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `repo_mapping` SET `checksum`=?,`customer_id`=?,`repo_id`=?,`ref_id`=?,`ref_type`=? WHERE `id`=?"
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RepoID),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.RefType),
		orm.ToSQLString(t.ID),
	)
}

// DBUpsert will upsert the RepoMapping record in the database
func (t *RepoMapping) DBUpsert(ctx context.Context, db DB, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `repo_mapping` (`repo_mapping`.`id`,`repo_mapping`.`checksum`,`repo_mapping`.`customer_id`,`repo_mapping`.`repo_id`,`repo_mapping`.`ref_id`,`repo_mapping`.`ref_type`) VALUES (?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `repo_mapping` (`repo_mapping`.`id`,`repo_mapping`.`checksum`,`repo_mapping`.`customer_id`,`repo_mapping`.`repo_id`,`repo_mapping`.`ref_id`,`repo_mapping`.`ref_type`) VALUES (?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`customer_id`=VALUES(`customer_id`),`repo_id`=VALUES(`repo_id`),`ref_id`=VALUES(`ref_id`),`ref_type`=VALUES(`ref_type`)"
	}
	r, err := db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RepoID),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.RefType),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBUpsertTx will upsert the RepoMapping record in the database using the provided transaction
func (t *RepoMapping) DBUpsertTx(ctx context.Context, tx Tx, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `repo_mapping` (`repo_mapping`.`id`,`repo_mapping`.`checksum`,`repo_mapping`.`customer_id`,`repo_mapping`.`repo_id`,`repo_mapping`.`ref_id`,`repo_mapping`.`ref_type`) VALUES (?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `repo_mapping` (`repo_mapping`.`id`,`repo_mapping`.`checksum`,`repo_mapping`.`customer_id`,`repo_mapping`.`repo_id`,`repo_mapping`.`ref_id`,`repo_mapping`.`ref_type`) VALUES (?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`customer_id`=VALUES(`customer_id`),`repo_id`=VALUES(`repo_id`),`ref_id`=VALUES(`ref_id`),`ref_type`=VALUES(`ref_type`)"
	}
	r, err := tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RepoID),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.RefType),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBFindOne will find a RepoMapping record in the database with the primary key
func (t *RepoMapping) DBFindOne(ctx context.Context, db DB, value string) (bool, error) {
	q := "SELECT * FROM `repo_mapping` WHERE `id` = ? LIMIT 1"
	row := db.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CustomerID sql.NullString
	var _RepoID sql.NullString
	var _RefID sql.NullString
	var _RefType sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_RepoID,
		&_RefID,
		&_RefType,
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
	if _RepoID.Valid {
		t.SetRepoID(_RepoID.String)
	}
	if _RefID.Valid {
		t.SetRefID(_RefID.String)
	}
	if _RefType.Valid {
		t.SetRefType(_RefType.String)
	}
	return true, nil
}

// DBFindOneTx will find a RepoMapping record in the database with the primary key using the provided transaction
func (t *RepoMapping) DBFindOneTx(ctx context.Context, tx Tx, value string) (bool, error) {
	q := "SELECT * FROM `repo_mapping` WHERE `id` = ? LIMIT 1"
	row := tx.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CustomerID sql.NullString
	var _RepoID sql.NullString
	var _RefID sql.NullString
	var _RefType sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_RepoID,
		&_RefID,
		&_RefType,
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
	if _RepoID.Valid {
		t.SetRepoID(_RepoID.String)
	}
	if _RefID.Valid {
		t.SetRefID(_RefID.String)
	}
	if _RefType.Valid {
		t.SetRefType(_RefType.String)
	}
	return true, nil
}

// FindRepoMappings will find a RepoMapping record in the database with the provided parameters
func FindRepoMappings(ctx context.Context, db DB, _params ...interface{}) ([]*RepoMapping, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("customer_id"),
		orm.Column("repo_id"),
		orm.Column("ref_id"),
		orm.Column("ref_type"),
		orm.Table(RepoMappingTableName),
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
	results := make([]*RepoMapping, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _RepoID sql.NullString
		var _RefID sql.NullString
		var _RefType sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_RepoID,
			&_RefID,
			&_RefType,
		)
		if err != nil {
			return nil, err
		}
		t := &RepoMapping{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
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
		results = append(results, t)
	}
	return results, nil
}

// FindRepoMappingsTx will find a RepoMapping record in the database with the provided parameters using the provided transaction
func FindRepoMappingsTx(ctx context.Context, tx Tx, _params ...interface{}) ([]*RepoMapping, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("customer_id"),
		orm.Column("repo_id"),
		orm.Column("ref_id"),
		orm.Column("ref_type"),
		orm.Table(RepoMappingTableName),
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
	results := make([]*RepoMapping, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _RepoID sql.NullString
		var _RefID sql.NullString
		var _RefType sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_RepoID,
			&_RefID,
			&_RefType,
		)
		if err != nil {
			return nil, err
		}
		t := &RepoMapping{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
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
		results = append(results, t)
	}
	return results, nil
}

// DBFind will find a RepoMapping record in the database with the provided parameters
func (t *RepoMapping) DBFind(ctx context.Context, db DB, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("customer_id"),
		orm.Column("repo_id"),
		orm.Column("ref_id"),
		orm.Column("ref_type"),
		orm.Table(RepoMappingTableName),
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
	var _RepoID sql.NullString
	var _RefID sql.NullString
	var _RefType sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_RepoID,
		&_RefID,
		&_RefType,
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
	if _RepoID.Valid {
		t.SetRepoID(_RepoID.String)
	}
	if _RefID.Valid {
		t.SetRefID(_RefID.String)
	}
	if _RefType.Valid {
		t.SetRefType(_RefType.String)
	}
	return true, nil
}

// DBFindTx will find a RepoMapping record in the database with the provided parameters using the provided transaction
func (t *RepoMapping) DBFindTx(ctx context.Context, tx Tx, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("customer_id"),
		orm.Column("repo_id"),
		orm.Column("ref_id"),
		orm.Column("ref_type"),
		orm.Table(RepoMappingTableName),
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
	var _RepoID sql.NullString
	var _RefID sql.NullString
	var _RefType sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_RepoID,
		&_RefID,
		&_RefType,
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
	if _RepoID.Valid {
		t.SetRepoID(_RepoID.String)
	}
	if _RefID.Valid {
		t.SetRefID(_RefID.String)
	}
	if _RefType.Valid {
		t.SetRefType(_RefType.String)
	}
	return true, nil
}

// CountRepoMappings will find the count of RepoMapping records in the database
func CountRepoMappings(ctx context.Context, db DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(RepoMappingTableName),
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

// CountRepoMappingsTx will find the count of RepoMapping records in the database using the provided transaction
func CountRepoMappingsTx(ctx context.Context, tx Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(RepoMappingTableName),
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

// DBCount will find the count of RepoMapping records in the database
func (t *RepoMapping) DBCount(ctx context.Context, db DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(RepoMappingTableName),
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

// DBCountTx will find the count of RepoMapping records in the database using the provided transaction
func (t *RepoMapping) DBCountTx(ctx context.Context, tx Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(RepoMappingTableName),
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

// DBExists will return true if the RepoMapping record exists in the database
func (t *RepoMapping) DBExists(ctx context.Context, db DB) (bool, error) {
	q := "SELECT * FROM `repo_mapping` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := db.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// DBExistsTx will return true if the RepoMapping record exists in the database using the provided transaction
func (t *RepoMapping) DBExistsTx(ctx context.Context, tx Tx) (bool, error) {
	q := "SELECT * FROM `repo_mapping` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := tx.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// PrimaryKeyColumn returns the column name for the primary key
func (t *RepoMapping) PrimaryKeyColumn() string {
	return RepoMappingColumnID
}

// PrimaryKeyColumnType returns the primary key column Go type as a string
func (t *RepoMapping) PrimaryKeyColumnType() string {
	return "string"
}

// PrimaryKey returns the primary key column value
func (t *RepoMapping) PrimaryKey() interface{} {
	return t.ID
}
