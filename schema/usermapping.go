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

var _ Model = (*UserMapping)(nil)
var _ CSVWriter = (*UserMapping)(nil)
var _ JSONWriter = (*UserMapping)(nil)
var _ Checksum = (*UserMapping)(nil)

// UserMappingTableName is the name of the table in SQL
const UserMappingTableName = "user_mapping"

var UserMappingColumns = []string{
	"id",
	"checksum",
	"customer_id",
	"user_id",
	"ref_id",
	"ref_type",
}

// UserMapping table
type UserMapping struct {
	Checksum   *string `json:"checksum,omitempty"`
	CustomerID string  `json:"customer_id"`
	ID         string  `json:"id"`
	RefID      string  `json:"ref_id"`
	RefType    string  `json:"ref_type"`
	UserID     string  `json:"user_id"`
}

// TableName returns the SQL table name for UserMapping and satifies the Model interface
func (t *UserMapping) TableName() string {
	return UserMappingTableName
}

// ToCSV will serialize the UserMapping instance to a CSV compatible array of strings
func (t *UserMapping) ToCSV() []string {
	return []string{
		t.ID,
		t.CalculateChecksum(),
		t.CustomerID,
		t.UserID,
		t.RefID,
		t.RefType,
	}
}

// WriteCSV will serialize the UserMapping instance to the writer as CSV and satisfies the CSVWriter interface
func (t *UserMapping) WriteCSV(w *csv.Writer) error {
	return w.Write(t.ToCSV())
}

// WriteJSON will serialize the UserMapping instance to the writer as JSON and satisfies the JSONWriter interface
func (t *UserMapping) WriteJSON(w io.Writer, indent ...bool) error {
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

// NewUserMappingReader creates a JSON reader which can read in UserMapping objects serialized as JSON either as an array, single object or json new lines
// and writes each UserMapping to the channel provided
func NewUserMappingReader(r io.Reader, ch chan<- UserMapping) error {
	return orm.Deserialize(r, func(buf json.RawMessage) error {
		dec := json.NewDecoder(bytes.NewBuffer(buf))
		e := UserMapping{}
		if err := dec.Decode(&e); err != nil {
			return err
		}
		ch <- e
		return nil
	})
}

// NewCSVUserMappingReaderDir will read the reader as CSV and emit each record to the channel provided
func NewCSVUserMappingReader(r io.Reader, ch chan<- UserMapping) error {
	cr := csv.NewReader(r)
	for {
		record, err := cr.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		ch <- UserMapping{
			ID:         record[0],
			Checksum:   fromStringPointer(record[1]),
			CustomerID: record[2],
			UserID:     record[3],
			RefID:      record[4],
			RefType:    record[5],
		}
	}
	return nil
}

// NewCSVUserMappingReaderFile will read the file as a CSV and emit each record to the channel provided
func NewCSVUserMappingReaderFile(fp string, ch chan<- UserMapping) error {
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
	return NewCSVUserMappingReader(fc, ch)
}

// NewCSVUserMappingReaderDir will read the user_mapping.csv.gz file as a CSV and emit each record to the channel provided
func NewCSVUserMappingReaderDir(dir string, ch chan<- UserMapping) error {
	return NewCSVUserMappingReaderFile(filepath.Join(dir, "user_mapping.csv.gz"), ch)
}

// UserMappingCSVDeduper is a function callback which takes the existing value (a) and the new value (b)
// and the return value should be the one to use (or a new one, if applicable). return nil
// to skip processing of this record
type UserMappingCSVDeduper func(a UserMapping, b UserMapping) *UserMapping

// UserMappingCSVDedupeDisabled is set on whether the CSV writer should de-dupe values by key
var UserMappingCSVDedupeDisabled bool

// NewUserMappingCSVWriterSize creates a batch writer that will write each UserMapping into a CSV file
// this method will automatically de-duplicate entries using the primary key. if the checksum
// for a newer item with the same primary key doesn't match a previously sent item, the newer
// one will be used
func NewUserMappingCSVWriterSize(w io.Writer, size int, dedupers ...UserMappingCSVDeduper) (chan UserMapping, chan bool, error) {
	cw := csv.NewWriter(w)
	ch := make(chan UserMapping, size)
	done := make(chan bool)
	go func() {
		defer func() { done <- true }()
		dodedupe := !UserMappingCSVDedupeDisabled
		var kv map[string]*UserMapping
		var deduper UserMappingCSVDeduper
		if dedupers != nil && len(dedupers) > 0 {
			deduper = dedupers[0]
			dodedupe = true
		}
		if dodedupe {
			kv = make(map[string]*UserMapping)
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

// UserMappingCSVDefaultSize is the default channel buffer size if not provided
var UserMappingCSVDefaultSize = 100

// NewUserMappingCSVWriter creates a batch writer that will write each UserMapping into a CSV file
func NewUserMappingCSVWriter(w io.Writer, dedupers ...UserMappingCSVDeduper) (chan UserMapping, chan bool, error) {
	return NewUserMappingCSVWriterSize(w, UserMappingCSVDefaultSize, dedupers...)
}

// NewUserMappingCSVWriterDir creates a batch writer that will write each UserMapping into a CSV file named user_mapping.csv.gz in dir
func NewUserMappingCSVWriterDir(dir string, dedupers ...UserMappingCSVDeduper) (chan UserMapping, chan bool, error) {
	return NewUserMappingCSVWriterFile(filepath.Join(dir, "user_mapping.csv.gz"), dedupers...)
}

// NewUserMappingCSVWriterFile creates a batch writer that will write each UserMapping into a CSV file
func NewUserMappingCSVWriterFile(fn string, dedupers ...UserMappingCSVDeduper) (chan UserMapping, chan bool, error) {
	f, err := os.Create(fn)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening CSV file %s. %v", fn, err)
	}
	var fc io.WriteCloser = f
	if filepath.Ext(fn) == ".gz" {
		w, _ := gzip.NewWriterLevel(f, gzip.BestCompression)
		fc = w
	}
	ch, done, err := NewUserMappingCSVWriter(fc, dedupers...)
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

type UserMappingDBAction func(ctx context.Context, db *sql.DB, record UserMapping) error

// NewUserMappingDBWriterSize creates a DB writer that will write each issue into the DB
func NewUserMappingDBWriterSize(ctx context.Context, db *sql.DB, errors chan<- error, size int, actions ...UserMappingDBAction) (chan UserMapping, chan bool, error) {
	ch := make(chan UserMapping, size)
	done := make(chan bool)
	var action UserMappingDBAction
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

// NewUserMappingDBWriter creates a DB writer that will write each issue into the DB
func NewUserMappingDBWriter(ctx context.Context, db *sql.DB, errors chan<- error, actions ...UserMappingDBAction) (chan UserMapping, chan bool, error) {
	return NewUserMappingDBWriterSize(ctx, db, errors, 100, actions...)
}

// UserMappingColumnID is the ID SQL column name for the UserMapping table
const UserMappingColumnID = "id"

// UserMappingEscapedColumnID is the escaped ID SQL column name for the UserMapping table
const UserMappingEscapedColumnID = "`id`"

// UserMappingColumnChecksum is the Checksum SQL column name for the UserMapping table
const UserMappingColumnChecksum = "checksum"

// UserMappingEscapedColumnChecksum is the escaped Checksum SQL column name for the UserMapping table
const UserMappingEscapedColumnChecksum = "`checksum`"

// UserMappingColumnCustomerID is the CustomerID SQL column name for the UserMapping table
const UserMappingColumnCustomerID = "customer_id"

// UserMappingEscapedColumnCustomerID is the escaped CustomerID SQL column name for the UserMapping table
const UserMappingEscapedColumnCustomerID = "`customer_id`"

// UserMappingColumnUserID is the UserID SQL column name for the UserMapping table
const UserMappingColumnUserID = "user_id"

// UserMappingEscapedColumnUserID is the escaped UserID SQL column name for the UserMapping table
const UserMappingEscapedColumnUserID = "`user_id`"

// UserMappingColumnRefID is the RefID SQL column name for the UserMapping table
const UserMappingColumnRefID = "ref_id"

// UserMappingEscapedColumnRefID is the escaped RefID SQL column name for the UserMapping table
const UserMappingEscapedColumnRefID = "`ref_id`"

// UserMappingColumnRefType is the RefType SQL column name for the UserMapping table
const UserMappingColumnRefType = "ref_type"

// UserMappingEscapedColumnRefType is the escaped RefType SQL column name for the UserMapping table
const UserMappingEscapedColumnRefType = "`ref_type`"

// GetID will return the UserMapping ID value
func (t *UserMapping) GetID() string {
	return t.ID
}

// SetID will set the UserMapping ID value
func (t *UserMapping) SetID(v string) {
	t.ID = v
}

// FindUserMappingByID will find a UserMapping by ID
func FindUserMappingByID(ctx context.Context, db *sql.DB, value string) (*UserMapping, error) {
	q := "SELECT `user_mapping`.`id`,`user_mapping`.`checksum`,`user_mapping`.`customer_id`,`user_mapping`.`user_id`,`user_mapping`.`ref_id`,`user_mapping`.`ref_type` FROM `user_mapping` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CustomerID sql.NullString
	var _UserID sql.NullString
	var _RefID sql.NullString
	var _RefType sql.NullString
	err := db.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_UserID,
		&_RefID,
		&_RefType,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &UserMapping{}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _Checksum.Valid {
		t.SetChecksum(_Checksum.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _UserID.Valid {
		t.SetUserID(_UserID.String)
	}
	if _RefID.Valid {
		t.SetRefID(_RefID.String)
	}
	if _RefType.Valid {
		t.SetRefType(_RefType.String)
	}
	return t, nil
}

// FindUserMappingByIDTx will find a UserMapping by ID using the provided transaction
func FindUserMappingByIDTx(ctx context.Context, tx *sql.Tx, value string) (*UserMapping, error) {
	q := "SELECT `user_mapping`.`id`,`user_mapping`.`checksum`,`user_mapping`.`customer_id`,`user_mapping`.`user_id`,`user_mapping`.`ref_id`,`user_mapping`.`ref_type` FROM `user_mapping` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CustomerID sql.NullString
	var _UserID sql.NullString
	var _RefID sql.NullString
	var _RefType sql.NullString
	err := tx.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_UserID,
		&_RefID,
		&_RefType,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &UserMapping{}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _Checksum.Valid {
		t.SetChecksum(_Checksum.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _UserID.Valid {
		t.SetUserID(_UserID.String)
	}
	if _RefID.Valid {
		t.SetRefID(_RefID.String)
	}
	if _RefType.Valid {
		t.SetRefType(_RefType.String)
	}
	return t, nil
}

// GetChecksum will return the UserMapping Checksum value
func (t *UserMapping) GetChecksum() string {
	if t.Checksum == nil {
		return ""
	}
	return *t.Checksum
}

// SetChecksum will set the UserMapping Checksum value
func (t *UserMapping) SetChecksum(v string) {
	t.Checksum = &v
}

// GetCustomerID will return the UserMapping CustomerID value
func (t *UserMapping) GetCustomerID() string {
	return t.CustomerID
}

// SetCustomerID will set the UserMapping CustomerID value
func (t *UserMapping) SetCustomerID(v string) {
	t.CustomerID = v
}

// GetUserID will return the UserMapping UserID value
func (t *UserMapping) GetUserID() string {
	return t.UserID
}

// SetUserID will set the UserMapping UserID value
func (t *UserMapping) SetUserID(v string) {
	t.UserID = v
}

// FindUserMappingsByUserID will find all UserMappings by the UserID value
func FindUserMappingsByUserID(ctx context.Context, db *sql.DB, value string) ([]*UserMapping, error) {
	q := "SELECT `user_mapping`.`id`,`user_mapping`.`checksum`,`user_mapping`.`customer_id`,`user_mapping`.`user_id`,`user_mapping`.`ref_id`,`user_mapping`.`ref_type` FROM `user_mapping` WHERE `user_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*UserMapping, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _UserID sql.NullString
		var _RefID sql.NullString
		var _RefType sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_UserID,
			&_RefID,
			&_RefType,
		)
		if err != nil {
			return nil, err
		}
		t := &UserMapping{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
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

// FindUserMappingsByUserIDTx will find all UserMappings by the UserID value using the provided transaction
func FindUserMappingsByUserIDTx(ctx context.Context, tx *sql.Tx, value string) ([]*UserMapping, error) {
	q := "SELECT `user_mapping`.`id`,`user_mapping`.`checksum`,`user_mapping`.`customer_id`,`user_mapping`.`user_id`,`user_mapping`.`ref_id`,`user_mapping`.`ref_type` FROM `user_mapping` WHERE `user_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*UserMapping, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _UserID sql.NullString
		var _RefID sql.NullString
		var _RefType sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_UserID,
			&_RefID,
			&_RefType,
		)
		if err != nil {
			return nil, err
		}
		t := &UserMapping{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
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

// GetRefID will return the UserMapping RefID value
func (t *UserMapping) GetRefID() string {
	return t.RefID
}

// SetRefID will set the UserMapping RefID value
func (t *UserMapping) SetRefID(v string) {
	t.RefID = v
}

// FindUserMappingsByRefID will find all UserMappings by the RefID value
func FindUserMappingsByRefID(ctx context.Context, db *sql.DB, value string) ([]*UserMapping, error) {
	q := "SELECT `user_mapping`.`id`,`user_mapping`.`checksum`,`user_mapping`.`customer_id`,`user_mapping`.`user_id`,`user_mapping`.`ref_id`,`user_mapping`.`ref_type` FROM `user_mapping` WHERE `ref_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*UserMapping, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _UserID sql.NullString
		var _RefID sql.NullString
		var _RefType sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_UserID,
			&_RefID,
			&_RefType,
		)
		if err != nil {
			return nil, err
		}
		t := &UserMapping{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
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

// FindUserMappingsByRefIDTx will find all UserMappings by the RefID value using the provided transaction
func FindUserMappingsByRefIDTx(ctx context.Context, tx *sql.Tx, value string) ([]*UserMapping, error) {
	q := "SELECT `user_mapping`.`id`,`user_mapping`.`checksum`,`user_mapping`.`customer_id`,`user_mapping`.`user_id`,`user_mapping`.`ref_id`,`user_mapping`.`ref_type` FROM `user_mapping` WHERE `ref_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*UserMapping, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _UserID sql.NullString
		var _RefID sql.NullString
		var _RefType sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_UserID,
			&_RefID,
			&_RefType,
		)
		if err != nil {
			return nil, err
		}
		t := &UserMapping{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
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

// GetRefType will return the UserMapping RefType value
func (t *UserMapping) GetRefType() string {
	return t.RefType
}

// SetRefType will set the UserMapping RefType value
func (t *UserMapping) SetRefType(v string) {
	t.RefType = v
}

func (t *UserMapping) toTimestamp(value time.Time) *timestamp.Timestamp {
	ts, _ := ptypes.TimestampProto(value)
	return ts
}

// DBCreateUserMappingTable will create the UserMapping table
func DBCreateUserMappingTable(ctx context.Context, db *sql.DB) error {
	q := "CREATE TABLE `user_mapping` (`id` VARCHAR(64) NOT NULL PRIMARY KEY,`checksum` CHAR(64),`customer_id` VARCHAR(64) NOT NULL,`user_id`VARCHAR(64) NOT NULL,`ref_id` VARCHAR(64) NOT NULL,`ref_type` VARCHAR(20) NOT NULL,INDEX user_mapping_user_id_index (`user_id`),INDEX user_mapping_ref_id_index (`ref_id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBCreateUserMappingTableTx will create the UserMapping table using the provided transction
func DBCreateUserMappingTableTx(ctx context.Context, tx *sql.Tx) error {
	q := "CREATE TABLE `user_mapping` (`id` VARCHAR(64) NOT NULL PRIMARY KEY,`checksum` CHAR(64),`customer_id` VARCHAR(64) NOT NULL,`user_id`VARCHAR(64) NOT NULL,`ref_id` VARCHAR(64) NOT NULL,`ref_type` VARCHAR(20) NOT NULL,INDEX user_mapping_user_id_index (`user_id`),INDEX user_mapping_ref_id_index (`ref_id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// DBDropUserMappingTable will drop the UserMapping table
func DBDropUserMappingTable(ctx context.Context, db *sql.DB) error {
	q := "DROP TABLE IF EXISTS `user_mapping`"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBDropUserMappingTableTx will drop the UserMapping table using the provided transaction
func DBDropUserMappingTableTx(ctx context.Context, tx *sql.Tx) error {
	q := "DROP TABLE IF EXISTS `user_mapping`"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// CalculateChecksum will calculate a checksum of the SHA1 of all field values
func (t *UserMapping) CalculateChecksum() string {
	return orm.HashStrings(
		orm.ToString(t.ID),
		orm.ToString(t.CustomerID),
		orm.ToString(t.UserID),
		orm.ToString(t.RefID),
		orm.ToString(t.RefType),
	)
}

// DBCreate will create a new UserMapping record in the database
func (t *UserMapping) DBCreate(ctx context.Context, db *sql.DB) (sql.Result, error) {
	q := "INSERT INTO `user_mapping` (`user_mapping`.`id`,`user_mapping`.`checksum`,`user_mapping`.`customer_id`,`user_mapping`.`user_id`,`user_mapping`.`ref_id`,`user_mapping`.`ref_type`) VALUES (?,?,?,?,?,?)"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.UserID),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.RefType),
	)
}

// DBCreateTx will create a new UserMapping record in the database using the provided transaction
func (t *UserMapping) DBCreateTx(ctx context.Context, tx *sql.Tx) (sql.Result, error) {
	q := "INSERT INTO `user_mapping` (`user_mapping`.`id`,`user_mapping`.`checksum`,`user_mapping`.`customer_id`,`user_mapping`.`user_id`,`user_mapping`.`ref_id`,`user_mapping`.`ref_type`) VALUES (?,?,?,?,?,?)"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.UserID),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.RefType),
	)
}

// DBCreateIgnoreDuplicate will upsert the UserMapping record in the database
func (t *UserMapping) DBCreateIgnoreDuplicate(ctx context.Context, db *sql.DB) (sql.Result, error) {
	q := "INSERT INTO `user_mapping` (`user_mapping`.`id`,`user_mapping`.`checksum`,`user_mapping`.`customer_id`,`user_mapping`.`user_id`,`user_mapping`.`ref_id`,`user_mapping`.`ref_type`) VALUES (?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.UserID),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.RefType),
	)
}

// DBCreateIgnoreDuplicateTx will upsert the UserMapping record in the database using the provided transaction
func (t *UserMapping) DBCreateIgnoreDuplicateTx(ctx context.Context, tx *sql.Tx) (sql.Result, error) {
	q := "INSERT INTO `user_mapping` (`user_mapping`.`id`,`user_mapping`.`checksum`,`user_mapping`.`customer_id`,`user_mapping`.`user_id`,`user_mapping`.`ref_id`,`user_mapping`.`ref_type`) VALUES (?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.UserID),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.RefType),
	)
}

// DeleteAllUserMappings deletes all UserMapping records in the database with optional filters
func DeleteAllUserMappings(ctx context.Context, db *sql.DB, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(UserMappingTableName),
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

// DeleteAllUserMappingsTx deletes all UserMapping records in the database with optional filters using the provided transaction
func DeleteAllUserMappingsTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(UserMappingTableName),
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

// DBDelete will delete this UserMapping record in the database
func (t *UserMapping) DBDelete(ctx context.Context, db *sql.DB) (bool, error) {
	q := "DELETE FROM `user_mapping` WHERE `id` = ?"
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

// DBDeleteTx will delete this UserMapping record in the database using the provided transaction
func (t *UserMapping) DBDeleteTx(ctx context.Context, tx *sql.Tx) (bool, error) {
	q := "DELETE FROM `user_mapping` WHERE `id` = ?"
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

// DBUpdate will update the UserMapping record in the database
func (t *UserMapping) DBUpdate(ctx context.Context, db *sql.DB) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `user_mapping` SET `checksum`=?,`customer_id`=?,`user_id`=?,`ref_id`=?,`ref_type`=? WHERE `id`=?"
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.UserID),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.RefType),
		orm.ToSQLString(t.ID),
	)
}

// DBUpdateTx will update the UserMapping record in the database using the provided transaction
func (t *UserMapping) DBUpdateTx(ctx context.Context, tx *sql.Tx) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `user_mapping` SET `checksum`=?,`customer_id`=?,`user_id`=?,`ref_id`=?,`ref_type`=? WHERE `id`=?"
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.UserID),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.RefType),
		orm.ToSQLString(t.ID),
	)
}

// DBUpsert will upsert the UserMapping record in the database
func (t *UserMapping) DBUpsert(ctx context.Context, db *sql.DB, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `user_mapping` (`user_mapping`.`id`,`user_mapping`.`checksum`,`user_mapping`.`customer_id`,`user_mapping`.`user_id`,`user_mapping`.`ref_id`,`user_mapping`.`ref_type`) VALUES (?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `user_mapping` (`user_mapping`.`id`,`user_mapping`.`checksum`,`user_mapping`.`customer_id`,`user_mapping`.`user_id`,`user_mapping`.`ref_id`,`user_mapping`.`ref_type`) VALUES (?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`customer_id`=VALUES(`customer_id`),`user_id`=VALUES(`user_id`),`ref_id`=VALUES(`ref_id`),`ref_type`=VALUES(`ref_type`)"
	}
	r, err := db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.UserID),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.RefType),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBUpsertTx will upsert the UserMapping record in the database using the provided transaction
func (t *UserMapping) DBUpsertTx(ctx context.Context, tx *sql.Tx, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `user_mapping` (`user_mapping`.`id`,`user_mapping`.`checksum`,`user_mapping`.`customer_id`,`user_mapping`.`user_id`,`user_mapping`.`ref_id`,`user_mapping`.`ref_type`) VALUES (?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `user_mapping` (`user_mapping`.`id`,`user_mapping`.`checksum`,`user_mapping`.`customer_id`,`user_mapping`.`user_id`,`user_mapping`.`ref_id`,`user_mapping`.`ref_type`) VALUES (?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`customer_id`=VALUES(`customer_id`),`user_id`=VALUES(`user_id`),`ref_id`=VALUES(`ref_id`),`ref_type`=VALUES(`ref_type`)"
	}
	r, err := tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.UserID),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.RefType),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBFindOne will find a UserMapping record in the database with the primary key
func (t *UserMapping) DBFindOne(ctx context.Context, db *sql.DB, value string) (bool, error) {
	q := "SELECT `user_mapping`.`id`,`user_mapping`.`checksum`,`user_mapping`.`customer_id`,`user_mapping`.`user_id`,`user_mapping`.`ref_id`,`user_mapping`.`ref_type` FROM `user_mapping` WHERE `id` = ? LIMIT 1"
	row := db.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CustomerID sql.NullString
	var _UserID sql.NullString
	var _RefID sql.NullString
	var _RefType sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_UserID,
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
	if _UserID.Valid {
		t.SetUserID(_UserID.String)
	}
	if _RefID.Valid {
		t.SetRefID(_RefID.String)
	}
	if _RefType.Valid {
		t.SetRefType(_RefType.String)
	}
	return true, nil
}

// DBFindOneTx will find a UserMapping record in the database with the primary key using the provided transaction
func (t *UserMapping) DBFindOneTx(ctx context.Context, tx *sql.Tx, value string) (bool, error) {
	q := "SELECT `user_mapping`.`id`,`user_mapping`.`checksum`,`user_mapping`.`customer_id`,`user_mapping`.`user_id`,`user_mapping`.`ref_id`,`user_mapping`.`ref_type` FROM `user_mapping` WHERE `id` = ? LIMIT 1"
	row := tx.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CustomerID sql.NullString
	var _UserID sql.NullString
	var _RefID sql.NullString
	var _RefType sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_UserID,
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
	if _UserID.Valid {
		t.SetUserID(_UserID.String)
	}
	if _RefID.Valid {
		t.SetRefID(_RefID.String)
	}
	if _RefType.Valid {
		t.SetRefType(_RefType.String)
	}
	return true, nil
}

// FindUserMappings will find a UserMapping record in the database with the provided parameters
func FindUserMappings(ctx context.Context, db *sql.DB, _params ...interface{}) ([]*UserMapping, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("customer_id"),
		orm.Column("user_id"),
		orm.Column("ref_id"),
		orm.Column("ref_type"),
		orm.Table(UserMappingTableName),
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
	results := make([]*UserMapping, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _UserID sql.NullString
		var _RefID sql.NullString
		var _RefType sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_UserID,
			&_RefID,
			&_RefType,
		)
		if err != nil {
			return nil, err
		}
		t := &UserMapping{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
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

// FindUserMappingsTx will find a UserMapping record in the database with the provided parameters using the provided transaction
func FindUserMappingsTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) ([]*UserMapping, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("customer_id"),
		orm.Column("user_id"),
		orm.Column("ref_id"),
		orm.Column("ref_type"),
		orm.Table(UserMappingTableName),
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
	results := make([]*UserMapping, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _UserID sql.NullString
		var _RefID sql.NullString
		var _RefType sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_UserID,
			&_RefID,
			&_RefType,
		)
		if err != nil {
			return nil, err
		}
		t := &UserMapping{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
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

// DBFind will find a UserMapping record in the database with the provided parameters
func (t *UserMapping) DBFind(ctx context.Context, db *sql.DB, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("customer_id"),
		orm.Column("user_id"),
		orm.Column("ref_id"),
		orm.Column("ref_type"),
		orm.Table(UserMappingTableName),
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
	var _UserID sql.NullString
	var _RefID sql.NullString
	var _RefType sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_UserID,
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
	if _UserID.Valid {
		t.SetUserID(_UserID.String)
	}
	if _RefID.Valid {
		t.SetRefID(_RefID.String)
	}
	if _RefType.Valid {
		t.SetRefType(_RefType.String)
	}
	return true, nil
}

// DBFindTx will find a UserMapping record in the database with the provided parameters using the provided transaction
func (t *UserMapping) DBFindTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("customer_id"),
		orm.Column("user_id"),
		orm.Column("ref_id"),
		orm.Column("ref_type"),
		orm.Table(UserMappingTableName),
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
	var _UserID sql.NullString
	var _RefID sql.NullString
	var _RefType sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_UserID,
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
	if _UserID.Valid {
		t.SetUserID(_UserID.String)
	}
	if _RefID.Valid {
		t.SetRefID(_RefID.String)
	}
	if _RefType.Valid {
		t.SetRefType(_RefType.String)
	}
	return true, nil
}

// CountUserMappings will find the count of UserMapping records in the database
func CountUserMappings(ctx context.Context, db *sql.DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(UserMappingTableName),
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

// CountUserMappingsTx will find the count of UserMapping records in the database using the provided transaction
func CountUserMappingsTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(UserMappingTableName),
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

// DBCount will find the count of UserMapping records in the database
func (t *UserMapping) DBCount(ctx context.Context, db *sql.DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(UserMappingTableName),
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

// DBCountTx will find the count of UserMapping records in the database using the provided transaction
func (t *UserMapping) DBCountTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(UserMappingTableName),
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

// DBExists will return true if the UserMapping record exists in the database
func (t *UserMapping) DBExists(ctx context.Context, db *sql.DB) (bool, error) {
	q := "SELECT `id` FROM `user_mapping` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := db.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// DBExistsTx will return true if the UserMapping record exists in the database using the provided transaction
func (t *UserMapping) DBExistsTx(ctx context.Context, tx *sql.Tx) (bool, error) {
	q := "SELECT `id` FROM `user_mapping` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := tx.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// PrimaryKeyColumn returns the column name for the primary key
func (t *UserMapping) PrimaryKeyColumn() string {
	return UserMappingColumnID
}

// PrimaryKeyColumnType returns the primary key column Go type as a string
func (t *UserMapping) PrimaryKeyColumnType() string {
	return "string"
}

// PrimaryKey returns the primary key column value
func (t *UserMapping) PrimaryKey() interface{} {
	return t.ID
}
