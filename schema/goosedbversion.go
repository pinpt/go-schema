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

var _ Model = (*GooseDbVersion)(nil)
var _ CSVWriter = (*GooseDbVersion)(nil)
var _ JSONWriter = (*GooseDbVersion)(nil)

// GooseDbVersionTableName is the name of the table in SQL
const GooseDbVersionTableName = "goose_db_version"

var GooseDbVersionColumns = []string{
	"id",
	"version_id",
	"is_applied",
	"tstamp",
}

// GooseDbVersion table
type GooseDbVersion struct {
	ID        int32                `json:"id"`
	IsApplied bool                 `json:"is_applied"`
	Tstamp    *timestamp.Timestamp `json:"tstamp"`
	VersionID int32                `json:"version_id"`
}

// TableName returns the SQL table name for GooseDbVersion and satifies the Model interface
func (t *GooseDbVersion) TableName() string {
	return GooseDbVersionTableName
}

// ToCSV will serialize the GooseDbVersion instance to a CSV compatible array of strings
func (t *GooseDbVersion) ToCSV() []string {
	return []string{
		toCSVString(t.ID),
		toCSVString(t.VersionID),
		toCSVBool(t.IsApplied),
		toCSVDate(t.Tstamp),
	}
}

// WriteCSV will serialize the GooseDbVersion instance to the writer as CSV and satisfies the CSVWriter interface
func (t *GooseDbVersion) WriteCSV(w *csv.Writer) error {
	return w.Write(t.ToCSV())
}

// WriteJSON will serialize the GooseDbVersion instance to the writer as JSON and satisfies the JSONWriter interface
func (t *GooseDbVersion) WriteJSON(w io.Writer, indent ...bool) error {
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

// NewGooseDbVersionReader creates a JSON reader which can read in GooseDbVersion objects serialized as JSON either as an array, single object or json new lines
// and writes each GooseDbVersion to the channel provided
func NewGooseDbVersionReader(r io.Reader, ch chan<- GooseDbVersion) error {
	return orm.Deserialize(r, func(buf json.RawMessage) error {
		dec := json.NewDecoder(bytes.NewBuffer(buf))
		e := GooseDbVersion{}
		if err := dec.Decode(&e); err != nil {
			return err
		}
		ch <- e
		return nil
	})
}

// NewCSVGooseDbVersionReaderDir will read the reader as CSV and emit each record to the channel provided
func NewCSVGooseDbVersionReader(r io.Reader, ch chan<- GooseDbVersion) error {
	cr := csv.NewReader(r)
	for {
		record, err := cr.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		ch <- GooseDbVersion{
			ID:        fromCSVInt32(record[0]),
			VersionID: fromCSVInt32(record[1]),
			IsApplied: fromCSVBool(record[2]),
			Tstamp:    fromCSVDate(record[3]),
		}
	}
	return nil
}

// NewCSVGooseDbVersionReaderFile will read the file as a CSV and emit each record to the channel provided
func NewCSVGooseDbVersionReaderFile(fp string, ch chan<- GooseDbVersion) error {
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
	return NewCSVGooseDbVersionReader(fc, ch)
}

// NewCSVGooseDbVersionReaderDir will read the goose_db_version.csv.gz file as a CSV and emit each record to the channel provided
func NewCSVGooseDbVersionReaderDir(dir string, ch chan<- GooseDbVersion) error {
	return NewCSVGooseDbVersionReaderFile(filepath.Join(dir, "goose_db_version.csv.gz"), ch)
}

// GooseDbVersionCSVDeduper is a function callback which takes the existing value (a) and the new value (b)
// and the return value should be the one to use (or a new one, if applicable). return nil
// to skip processing of this record
type GooseDbVersionCSVDeduper func(a GooseDbVersion, b GooseDbVersion) *GooseDbVersion

// GooseDbVersionCSVDedupeDisabled is set on whether the CSV writer should de-dupe values by key
var GooseDbVersionCSVDedupeDisabled bool

// NewGooseDbVersionCSVWriterSize creates a batch writer that will write each GooseDbVersion into a CSV file
func NewGooseDbVersionCSVWriterSize(w io.Writer, size int, dedupers ...GooseDbVersionCSVDeduper) (chan GooseDbVersion, chan bool, error) {
	cw := csv.NewWriter(w)
	ch := make(chan GooseDbVersion, size)
	done := make(chan bool)
	go func() {
		defer func() { done <- true }()
		dodedupe := !GooseDbVersionCSVDedupeDisabled
		var kv map[int32]*GooseDbVersion
		var deduper GooseDbVersionCSVDeduper
		if dedupers != nil && len(dedupers) > 0 {
			deduper = dedupers[0]
			dodedupe = true
		}
		if dodedupe {
			kv = make(map[int32]*GooseDbVersion)
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

// GooseDbVersionCSVDefaultSize is the default channel buffer size if not provided
var GooseDbVersionCSVDefaultSize = 100

// NewGooseDbVersionCSVWriter creates a batch writer that will write each GooseDbVersion into a CSV file
func NewGooseDbVersionCSVWriter(w io.Writer, dedupers ...GooseDbVersionCSVDeduper) (chan GooseDbVersion, chan bool, error) {
	return NewGooseDbVersionCSVWriterSize(w, GooseDbVersionCSVDefaultSize, dedupers...)
}

// NewGooseDbVersionCSVWriterDir creates a batch writer that will write each GooseDbVersion into a CSV file named goose_db_version.csv.gz in dir
func NewGooseDbVersionCSVWriterDir(dir string, dedupers ...GooseDbVersionCSVDeduper) (chan GooseDbVersion, chan bool, error) {
	return NewGooseDbVersionCSVWriterFile(filepath.Join(dir, "goose_db_version.csv.gz"), dedupers...)
}

// NewGooseDbVersionCSVWriterFile creates a batch writer that will write each GooseDbVersion into a CSV file
func NewGooseDbVersionCSVWriterFile(fn string, dedupers ...GooseDbVersionCSVDeduper) (chan GooseDbVersion, chan bool, error) {
	f, err := os.Create(fn)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening CSV file %s. %v", fn, err)
	}
	var fc io.WriteCloser = f
	if filepath.Ext(fn) == ".gz" {
		w, _ := gzip.NewWriterLevel(f, gzip.BestCompression)
		fc = w
	}
	ch, done, err := NewGooseDbVersionCSVWriter(fc, dedupers...)
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

type GooseDbVersionDBAction func(ctx context.Context, db DB, record GooseDbVersion) error

// NewGooseDbVersionDBWriterSize creates a DB writer that will write each issue into the DB
func NewGooseDbVersionDBWriterSize(ctx context.Context, db DB, errors chan<- error, size int, actions ...GooseDbVersionDBAction) (chan GooseDbVersion, chan bool, error) {
	ch := make(chan GooseDbVersion, size)
	done := make(chan bool)
	var action GooseDbVersionDBAction
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

// NewGooseDbVersionDBWriter creates a DB writer that will write each issue into the DB
func NewGooseDbVersionDBWriter(ctx context.Context, db DB, errors chan<- error, actions ...GooseDbVersionDBAction) (chan GooseDbVersion, chan bool, error) {
	return NewGooseDbVersionDBWriterSize(ctx, db, errors, 100, actions...)
}

// GooseDbVersionColumnID is the ID SQL column name for the GooseDbVersion table
const GooseDbVersionColumnID = "id"

// GooseDbVersionEscapedColumnID is the escaped ID SQL column name for the GooseDbVersion table
const GooseDbVersionEscapedColumnID = "`id`"

// GooseDbVersionColumnVersionID is the VersionID SQL column name for the GooseDbVersion table
const GooseDbVersionColumnVersionID = "version_id"

// GooseDbVersionEscapedColumnVersionID is the escaped VersionID SQL column name for the GooseDbVersion table
const GooseDbVersionEscapedColumnVersionID = "`version_id`"

// GooseDbVersionColumnIsApplied is the IsApplied SQL column name for the GooseDbVersion table
const GooseDbVersionColumnIsApplied = "is_applied"

// GooseDbVersionEscapedColumnIsApplied is the escaped IsApplied SQL column name for the GooseDbVersion table
const GooseDbVersionEscapedColumnIsApplied = "`is_applied`"

// GooseDbVersionColumnTstamp is the Tstamp SQL column name for the GooseDbVersion table
const GooseDbVersionColumnTstamp = "tstamp"

// GooseDbVersionEscapedColumnTstamp is the escaped Tstamp SQL column name for the GooseDbVersion table
const GooseDbVersionEscapedColumnTstamp = "`tstamp`"

// GetID will return the GooseDbVersion ID value
func (t *GooseDbVersion) GetID() int32 {
	return t.ID
}

// SetID will set the GooseDbVersion ID value
func (t *GooseDbVersion) SetID(v int32) {
	t.ID = v
}

// FindGooseDbVersionByID will find a GooseDbVersion by ID
func FindGooseDbVersionByID(ctx context.Context, db DB, value int32) (*GooseDbVersion, error) {
	q := "SELECT * FROM `goose_db_version` WHERE `id` = ?"
	var _ID sql.NullInt64
	var _VersionID sql.NullInt64
	var _IsApplied sql.NullBool
	var _Tstamp NullTime
	err := db.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_VersionID,
		&_IsApplied,
		&_Tstamp,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &GooseDbVersion{}
	if _ID.Valid {
		t.SetID(int32(_ID.Int64))
	}
	if _VersionID.Valid {
		t.SetVersionID(int32(_VersionID.Int64))
	}
	if _IsApplied.Valid {
		t.SetIsApplied(_IsApplied.Bool)
	}
	if _Tstamp.Valid {
		t.SetTstamp(t.toTimestamp(_Tstamp.Time))
	}
	return t, nil
}

// FindGooseDbVersionByIDTx will find a GooseDbVersion by ID using the provided transaction
func FindGooseDbVersionByIDTx(ctx context.Context, tx Tx, value int32) (*GooseDbVersion, error) {
	q := "SELECT * FROM `goose_db_version` WHERE `id` = ?"
	var _ID sql.NullInt64
	var _VersionID sql.NullInt64
	var _IsApplied sql.NullBool
	var _Tstamp NullTime
	err := tx.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_VersionID,
		&_IsApplied,
		&_Tstamp,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &GooseDbVersion{}
	if _ID.Valid {
		t.SetID(int32(_ID.Int64))
	}
	if _VersionID.Valid {
		t.SetVersionID(int32(_VersionID.Int64))
	}
	if _IsApplied.Valid {
		t.SetIsApplied(_IsApplied.Bool)
	}
	if _Tstamp.Valid {
		t.SetTstamp(t.toTimestamp(_Tstamp.Time))
	}
	return t, nil
}

// GetVersionID will return the GooseDbVersion VersionID value
func (t *GooseDbVersion) GetVersionID() int32 {
	return t.VersionID
}

// SetVersionID will set the GooseDbVersion VersionID value
func (t *GooseDbVersion) SetVersionID(v int32) {
	t.VersionID = v
}

// GetIsApplied will return the GooseDbVersion IsApplied value
func (t *GooseDbVersion) GetIsApplied() bool {
	return t.IsApplied
}

// SetIsApplied will set the GooseDbVersion IsApplied value
func (t *GooseDbVersion) SetIsApplied(v bool) {
	t.IsApplied = v
}

// GetTstamp will return the GooseDbVersion Tstamp value
func (t *GooseDbVersion) GetTstamp() *timestamp.Timestamp {
	return t.Tstamp
}

// SetTstamp will set the GooseDbVersion Tstamp value
func (t *GooseDbVersion) SetTstamp(v *timestamp.Timestamp) {
	t.Tstamp = v
}

func (t *GooseDbVersion) toTimestamp(value time.Time) *timestamp.Timestamp {
	ts, _ := ptypes.TimestampProto(value)
	return ts
}

// DBCreateGooseDbVersionTable will create the GooseDbVersion table
func DBCreateGooseDbVersionTable(ctx context.Context, db DB) error {
	q := "CREATE TABLE `goose_db_version` (`id` INT NOT NULL PRIMARY KEY,`version_id` INT NOT NULL,`is_applied` BOOL NOT NULL,`tstamp`DATETIME NOT NULL) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBCreateGooseDbVersionTableTx will create the GooseDbVersion table using the provided transction
func DBCreateGooseDbVersionTableTx(ctx context.Context, tx Tx) error {
	q := "CREATE TABLE `goose_db_version` (`id` INT NOT NULL PRIMARY KEY,`version_id` INT NOT NULL,`is_applied` BOOL NOT NULL,`tstamp`DATETIME NOT NULL) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// DBDropGooseDbVersionTable will drop the GooseDbVersion table
func DBDropGooseDbVersionTable(ctx context.Context, db DB) error {
	q := "DROP TABLE IF EXISTS `goose_db_version`"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBDropGooseDbVersionTableTx will drop the GooseDbVersion table using the provided transaction
func DBDropGooseDbVersionTableTx(ctx context.Context, tx Tx) error {
	q := "DROP TABLE IF EXISTS `goose_db_version`"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// DBCreate will create a new GooseDbVersion record in the database
func (t *GooseDbVersion) DBCreate(ctx context.Context, db DB) (sql.Result, error) {
	q := "INSERT INTO `goose_db_version` (`goose_db_version`.`id`,`goose_db_version`.`version_id`,`goose_db_version`.`is_applied`,`goose_db_version`.`tstamp`) VALUES (?,?,?,?)"
	return db.ExecContext(ctx, q,
		orm.ToSQLInt64(t.ID),
		orm.ToSQLInt64(t.VersionID),
		orm.ToSQLBool(t.IsApplied),
		orm.ToSQLDate(t.Tstamp),
	)
}

// DBCreateTx will create a new GooseDbVersion record in the database using the provided transaction
func (t *GooseDbVersion) DBCreateTx(ctx context.Context, tx Tx) (sql.Result, error) {
	q := "INSERT INTO `goose_db_version` (`goose_db_version`.`id`,`goose_db_version`.`version_id`,`goose_db_version`.`is_applied`,`goose_db_version`.`tstamp`) VALUES (?,?,?,?)"
	return tx.ExecContext(ctx, q,
		orm.ToSQLInt64(t.ID),
		orm.ToSQLInt64(t.VersionID),
		orm.ToSQLBool(t.IsApplied),
		orm.ToSQLDate(t.Tstamp),
	)
}

// DBCreateIgnoreDuplicate will upsert the GooseDbVersion record in the database
func (t *GooseDbVersion) DBCreateIgnoreDuplicate(ctx context.Context, db DB) (sql.Result, error) {
	q := "INSERT INTO `goose_db_version` (`goose_db_version`.`id`,`goose_db_version`.`version_id`,`goose_db_version`.`is_applied`,`goose_db_version`.`tstamp`) VALUES (?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
	return db.ExecContext(ctx, q,
		orm.ToSQLInt64(t.ID),
		orm.ToSQLInt64(t.VersionID),
		orm.ToSQLBool(t.IsApplied),
		orm.ToSQLDate(t.Tstamp),
	)
}

// DBCreateIgnoreDuplicateTx will upsert the GooseDbVersion record in the database using the provided transaction
func (t *GooseDbVersion) DBCreateIgnoreDuplicateTx(ctx context.Context, tx Tx) (sql.Result, error) {
	q := "INSERT INTO `goose_db_version` (`goose_db_version`.`id`,`goose_db_version`.`version_id`,`goose_db_version`.`is_applied`,`goose_db_version`.`tstamp`) VALUES (?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
	return tx.ExecContext(ctx, q,
		orm.ToSQLInt64(t.ID),
		orm.ToSQLInt64(t.VersionID),
		orm.ToSQLBool(t.IsApplied),
		orm.ToSQLDate(t.Tstamp),
	)
}

// DeleteAllGooseDbVersions deletes all GooseDbVersion records in the database with optional filters
func DeleteAllGooseDbVersions(ctx context.Context, db DB, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(GooseDbVersionTableName),
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

// DeleteAllGooseDbVersionsTx deletes all GooseDbVersion records in the database with optional filters using the provided transaction
func DeleteAllGooseDbVersionsTx(ctx context.Context, tx Tx, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(GooseDbVersionTableName),
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

// DBDelete will delete this GooseDbVersion record in the database
func (t *GooseDbVersion) DBDelete(ctx context.Context, db DB) (bool, error) {
	q := "DELETE FROM `goose_db_version` WHERE `id` = ?"
	r, err := db.ExecContext(ctx, q, orm.ToSQLInt64(t.ID))
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	if err == sql.ErrNoRows {
		return false, nil
	}
	c, _ := r.RowsAffected()
	return c > 0, nil
}

// DBDeleteTx will delete this GooseDbVersion record in the database using the provided transaction
func (t *GooseDbVersion) DBDeleteTx(ctx context.Context, tx Tx) (bool, error) {
	q := "DELETE FROM `goose_db_version` WHERE `id` = ?"
	r, err := tx.ExecContext(ctx, q, orm.ToSQLInt64(t.ID))
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	if err == sql.ErrNoRows {
		return false, nil
	}
	c, _ := r.RowsAffected()
	return c > 0, nil
}

// DBUpdate will update the GooseDbVersion record in the database
func (t *GooseDbVersion) DBUpdate(ctx context.Context, db DB) (sql.Result, error) {
	q := "UPDATE `goose_db_version` SET `version_id`=?,`is_applied`=?,`tstamp`=? WHERE `id`=?"
	return db.ExecContext(ctx, q,
		orm.ToSQLInt64(t.VersionID),
		orm.ToSQLBool(t.IsApplied),
		orm.ToSQLDate(t.Tstamp),
		orm.ToSQLInt64(t.ID),
	)
}

// DBUpdateTx will update the GooseDbVersion record in the database using the provided transaction
func (t *GooseDbVersion) DBUpdateTx(ctx context.Context, tx Tx) (sql.Result, error) {
	q := "UPDATE `goose_db_version` SET `version_id`=?,`is_applied`=?,`tstamp`=? WHERE `id`=?"
	return tx.ExecContext(ctx, q,
		orm.ToSQLInt64(t.VersionID),
		orm.ToSQLBool(t.IsApplied),
		orm.ToSQLDate(t.Tstamp),
		orm.ToSQLInt64(t.ID),
	)
}

// DBUpsert will upsert the GooseDbVersion record in the database
func (t *GooseDbVersion) DBUpsert(ctx context.Context, db DB, conditions ...interface{}) (bool, bool, error) {
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `goose_db_version` (`goose_db_version`.`id`,`goose_db_version`.`version_id`,`goose_db_version`.`is_applied`,`goose_db_version`.`tstamp`) VALUES (?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `goose_db_version` (`goose_db_version`.`id`,`goose_db_version`.`version_id`,`goose_db_version`.`is_applied`,`goose_db_version`.`tstamp`) VALUES (?,?,?,?) ON DUPLICATE KEY UPDATE `version_id`=VALUES(`version_id`),`is_applied`=VALUES(`is_applied`),`tstamp`=VALUES(`tstamp`)"
	}
	r, err := db.ExecContext(ctx, q,
		orm.ToSQLInt64(t.ID),
		orm.ToSQLInt64(t.VersionID),
		orm.ToSQLBool(t.IsApplied),
		orm.ToSQLDate(t.Tstamp),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBUpsertTx will upsert the GooseDbVersion record in the database using the provided transaction
func (t *GooseDbVersion) DBUpsertTx(ctx context.Context, tx Tx, conditions ...interface{}) (bool, bool, error) {
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `goose_db_version` (`goose_db_version`.`id`,`goose_db_version`.`version_id`,`goose_db_version`.`is_applied`,`goose_db_version`.`tstamp`) VALUES (?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `goose_db_version` (`goose_db_version`.`id`,`goose_db_version`.`version_id`,`goose_db_version`.`is_applied`,`goose_db_version`.`tstamp`) VALUES (?,?,?,?) ON DUPLICATE KEY UPDATE `version_id`=VALUES(`version_id`),`is_applied`=VALUES(`is_applied`),`tstamp`=VALUES(`tstamp`)"
	}
	r, err := tx.ExecContext(ctx, q,
		orm.ToSQLInt64(t.ID),
		orm.ToSQLInt64(t.VersionID),
		orm.ToSQLBool(t.IsApplied),
		orm.ToSQLDate(t.Tstamp),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBFindOne will find a GooseDbVersion record in the database with the primary key
func (t *GooseDbVersion) DBFindOne(ctx context.Context, db DB, value int32) (bool, error) {
	q := "SELECT * FROM `goose_db_version` WHERE `id` = ? LIMIT 1"
	row := db.QueryRowContext(ctx, q, orm.ToSQLInt64(value))
	var _ID sql.NullInt64
	var _VersionID sql.NullInt64
	var _IsApplied sql.NullBool
	var _Tstamp NullTime
	err := row.Scan(
		&_ID,
		&_VersionID,
		&_IsApplied,
		&_Tstamp,
	)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	if _ID.Valid == false {
		return false, nil
	}
	if _ID.Valid {
		t.SetID(int32(_ID.Int64))
	}
	if _VersionID.Valid {
		t.SetVersionID(int32(_VersionID.Int64))
	}
	if _IsApplied.Valid {
		t.SetIsApplied(_IsApplied.Bool)
	}
	if _Tstamp.Valid {
		t.SetTstamp(t.toTimestamp(_Tstamp.Time))
	}
	return true, nil
}

// DBFindOneTx will find a GooseDbVersion record in the database with the primary key using the provided transaction
func (t *GooseDbVersion) DBFindOneTx(ctx context.Context, tx Tx, value int32) (bool, error) {
	q := "SELECT * FROM `goose_db_version` WHERE `id` = ? LIMIT 1"
	row := tx.QueryRowContext(ctx, q, orm.ToSQLInt64(value))
	var _ID sql.NullInt64
	var _VersionID sql.NullInt64
	var _IsApplied sql.NullBool
	var _Tstamp NullTime
	err := row.Scan(
		&_ID,
		&_VersionID,
		&_IsApplied,
		&_Tstamp,
	)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	if _ID.Valid == false {
		return false, nil
	}
	if _ID.Valid {
		t.SetID(int32(_ID.Int64))
	}
	if _VersionID.Valid {
		t.SetVersionID(int32(_VersionID.Int64))
	}
	if _IsApplied.Valid {
		t.SetIsApplied(_IsApplied.Bool)
	}
	if _Tstamp.Valid {
		t.SetTstamp(t.toTimestamp(_Tstamp.Time))
	}
	return true, nil
}

// FindGooseDbVersions will find a GooseDbVersion record in the database with the provided parameters
func FindGooseDbVersions(ctx context.Context, db DB, _params ...interface{}) ([]*GooseDbVersion, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("version_id"),
		orm.Column("is_applied"),
		orm.Column("tstamp"),
		orm.Table(GooseDbVersionTableName),
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
	results := make([]*GooseDbVersion, 0)
	for rows.Next() {
		var _ID sql.NullInt64
		var _VersionID sql.NullInt64
		var _IsApplied sql.NullBool
		var _Tstamp NullTime
		err := rows.Scan(
			&_ID,
			&_VersionID,
			&_IsApplied,
			&_Tstamp,
		)
		if err != nil {
			return nil, err
		}
		t := &GooseDbVersion{}
		if _ID.Valid {
			t.SetID(int32(_ID.Int64))
		}
		if _VersionID.Valid {
			t.SetVersionID(int32(_VersionID.Int64))
		}
		if _IsApplied.Valid {
			t.SetIsApplied(_IsApplied.Bool)
		}
		if _Tstamp.Valid {
			t.SetTstamp(t.toTimestamp(_Tstamp.Time))
		}
		results = append(results, t)
	}
	return results, nil
}

// FindGooseDbVersionsTx will find a GooseDbVersion record in the database with the provided parameters using the provided transaction
func FindGooseDbVersionsTx(ctx context.Context, tx Tx, _params ...interface{}) ([]*GooseDbVersion, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("version_id"),
		orm.Column("is_applied"),
		orm.Column("tstamp"),
		orm.Table(GooseDbVersionTableName),
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
	results := make([]*GooseDbVersion, 0)
	for rows.Next() {
		var _ID sql.NullInt64
		var _VersionID sql.NullInt64
		var _IsApplied sql.NullBool
		var _Tstamp NullTime
		err := rows.Scan(
			&_ID,
			&_VersionID,
			&_IsApplied,
			&_Tstamp,
		)
		if err != nil {
			return nil, err
		}
		t := &GooseDbVersion{}
		if _ID.Valid {
			t.SetID(int32(_ID.Int64))
		}
		if _VersionID.Valid {
			t.SetVersionID(int32(_VersionID.Int64))
		}
		if _IsApplied.Valid {
			t.SetIsApplied(_IsApplied.Bool)
		}
		if _Tstamp.Valid {
			t.SetTstamp(t.toTimestamp(_Tstamp.Time))
		}
		results = append(results, t)
	}
	return results, nil
}

// DBFind will find a GooseDbVersion record in the database with the provided parameters
func (t *GooseDbVersion) DBFind(ctx context.Context, db DB, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("version_id"),
		orm.Column("is_applied"),
		orm.Column("tstamp"),
		orm.Table(GooseDbVersionTableName),
	}
	if len(_params) > 0 {
		for _, param := range _params {
			params = append(params, param)
		}
	}
	q, p := orm.BuildQuery(params...)
	row := db.QueryRowContext(ctx, q, p...)
	var _ID sql.NullInt64
	var _VersionID sql.NullInt64
	var _IsApplied sql.NullBool
	var _Tstamp NullTime
	err := row.Scan(
		&_ID,
		&_VersionID,
		&_IsApplied,
		&_Tstamp,
	)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	if _ID.Valid {
		t.SetID(int32(_ID.Int64))
	}
	if _VersionID.Valid {
		t.SetVersionID(int32(_VersionID.Int64))
	}
	if _IsApplied.Valid {
		t.SetIsApplied(_IsApplied.Bool)
	}
	if _Tstamp.Valid {
		t.SetTstamp(t.toTimestamp(_Tstamp.Time))
	}
	return true, nil
}

// DBFindTx will find a GooseDbVersion record in the database with the provided parameters using the provided transaction
func (t *GooseDbVersion) DBFindTx(ctx context.Context, tx Tx, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("version_id"),
		orm.Column("is_applied"),
		orm.Column("tstamp"),
		orm.Table(GooseDbVersionTableName),
	}
	if len(_params) > 0 {
		for _, param := range _params {
			params = append(params, param)
		}
	}
	q, p := orm.BuildQuery(params...)
	row := tx.QueryRowContext(ctx, q, p...)
	var _ID sql.NullInt64
	var _VersionID sql.NullInt64
	var _IsApplied sql.NullBool
	var _Tstamp NullTime
	err := row.Scan(
		&_ID,
		&_VersionID,
		&_IsApplied,
		&_Tstamp,
	)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	if _ID.Valid {
		t.SetID(int32(_ID.Int64))
	}
	if _VersionID.Valid {
		t.SetVersionID(int32(_VersionID.Int64))
	}
	if _IsApplied.Valid {
		t.SetIsApplied(_IsApplied.Bool)
	}
	if _Tstamp.Valid {
		t.SetTstamp(t.toTimestamp(_Tstamp.Time))
	}
	return true, nil
}

// CountGooseDbVersions will find the count of GooseDbVersion records in the database
func CountGooseDbVersions(ctx context.Context, db DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(GooseDbVersionTableName),
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

// CountGooseDbVersionsTx will find the count of GooseDbVersion records in the database using the provided transaction
func CountGooseDbVersionsTx(ctx context.Context, tx Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(GooseDbVersionTableName),
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

// DBCount will find the count of GooseDbVersion records in the database
func (t *GooseDbVersion) DBCount(ctx context.Context, db DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(GooseDbVersionTableName),
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

// DBCountTx will find the count of GooseDbVersion records in the database using the provided transaction
func (t *GooseDbVersion) DBCountTx(ctx context.Context, tx Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(GooseDbVersionTableName),
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

// DBExists will return true if the GooseDbVersion record exists in the database
func (t *GooseDbVersion) DBExists(ctx context.Context, db DB) (bool, error) {
	q := "SELECT * FROM `goose_db_version` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := db.QueryRowContext(ctx, q, orm.ToSQLInt64(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// DBExistsTx will return true if the GooseDbVersion record exists in the database using the provided transaction
func (t *GooseDbVersion) DBExistsTx(ctx context.Context, tx Tx) (bool, error) {
	q := "SELECT * FROM `goose_db_version` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := tx.QueryRowContext(ctx, q, orm.ToSQLInt64(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// PrimaryKeyColumn returns the column name for the primary key
func (t *GooseDbVersion) PrimaryKeyColumn() string {
	return GooseDbVersionColumnID
}

// PrimaryKeyColumnType returns the primary key column Go type as a string
func (t *GooseDbVersion) PrimaryKeyColumnType() string {
	return "int32"
}

// PrimaryKey returns the primary key column value
func (t *GooseDbVersion) PrimaryKey() interface{} {
	return t.ID
}
