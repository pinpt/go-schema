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

var _ Model = (*UserLogin)(nil)
var _ CSVWriter = (*UserLogin)(nil)
var _ JSONWriter = (*UserLogin)(nil)
var _ Checksum = (*UserLogin)(nil)

// UserLoginTableName is the name of the table in SQL
const UserLoginTableName = "user_login"

var UserLoginColumns = []string{
	"id",
	"checksum",
	"customer_id",
	"user_id",
	"browser",
	"date",
}

// UserLogin table
type UserLogin struct {
	Browser    *string `json:"browser,omitempty"`
	Checksum   *string `json:"checksum,omitempty"`
	CustomerID string  `json:"customer_id"`
	Date       int64   `json:"date"`
	ID         string  `json:"id"`
	UserID     string  `json:"user_id"`
}

// TableName returns the SQL table name for UserLogin and satifies the Model interface
func (t *UserLogin) TableName() string {
	return UserLoginTableName
}

// ToCSV will serialize the UserLogin instance to a CSV compatible array of strings
func (t *UserLogin) ToCSV() []string {
	return []string{
		t.ID,
		t.CalculateChecksum(),
		t.CustomerID,
		t.UserID,
		toCSVString(t.Browser),
		toCSVString(t.Date),
	}
}

// WriteCSV will serialize the UserLogin instance to the writer as CSV and satisfies the CSVWriter interface
func (t *UserLogin) WriteCSV(w *csv.Writer) error {
	return w.Write(t.ToCSV())
}

// WriteJSON will serialize the UserLogin instance to the writer as JSON and satisfies the JSONWriter interface
func (t *UserLogin) WriteJSON(w io.Writer, indent ...bool) error {
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

// NewUserLoginReader creates a JSON reader which can read in UserLogin objects serialized as JSON either as an array, single object or json new lines
// and writes each UserLogin to the channel provided
func NewUserLoginReader(r io.Reader, ch chan<- UserLogin) error {
	return orm.Deserialize(r, func(buf json.RawMessage) error {
		dec := json.NewDecoder(bytes.NewBuffer(buf))
		e := UserLogin{}
		if err := dec.Decode(&e); err != nil {
			return err
		}
		ch <- e
		return nil
	})
}

// NewCSVUserLoginReaderDir will read the reader as CSV and emit each record to the channel provided
func NewCSVUserLoginReader(r io.Reader, ch chan<- UserLogin) error {
	cr := csv.NewReader(r)
	for {
		record, err := cr.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		ch <- UserLogin{
			ID:         record[0],
			Checksum:   fromStringPointer(record[1]),
			CustomerID: record[2],
			UserID:     record[3],
			Browser:    fromStringPointer(record[4]),
			Date:       fromCSVInt64(record[5]),
		}
	}
	return nil
}

// NewCSVUserLoginReaderFile will read the file as a CSV and emit each record to the channel provided
func NewCSVUserLoginReaderFile(fp string, ch chan<- UserLogin) error {
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
	return NewCSVUserLoginReader(fc, ch)
}

// NewCSVUserLoginReaderDir will read the user_login.csv.gz file as a CSV and emit each record to the channel provided
func NewCSVUserLoginReaderDir(dir string, ch chan<- UserLogin) error {
	return NewCSVUserLoginReaderFile(filepath.Join(dir, "user_login.csv.gz"), ch)
}

// UserLoginCSVDeduper is a function callback which takes the existing value (a) and the new value (b)
// and the return value should be the one to use (or a new one, if applicable). return nil
// to skip processing of this record
type UserLoginCSVDeduper func(a UserLogin, b UserLogin) *UserLogin

// UserLoginCSVDedupeDisabled is set on whether the CSV writer should de-dupe values by key
var UserLoginCSVDedupeDisabled bool

// NewUserLoginCSVWriterSize creates a batch writer that will write each UserLogin into a CSV file
// this method will automatically de-duplicate entries using the primary key. if the checksum
// for a newer item with the same primary key doesn't match a previously sent item, the newer
// one will be used
func NewUserLoginCSVWriterSize(w io.Writer, size int, dedupers ...UserLoginCSVDeduper) (chan UserLogin, chan bool, error) {
	cw := csv.NewWriter(w)
	ch := make(chan UserLogin, size)
	done := make(chan bool)
	go func() {
		defer func() { done <- true }()
		dodedupe := !UserLoginCSVDedupeDisabled
		var kv map[string]*UserLogin
		var deduper UserLoginCSVDeduper
		if dedupers != nil && len(dedupers) > 0 {
			deduper = dedupers[0]
			dodedupe = true
		}
		if dodedupe {
			kv = make(map[string]*UserLogin)
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

// UserLoginCSVDefaultSize is the default channel buffer size if not provided
var UserLoginCSVDefaultSize = 100

// NewUserLoginCSVWriter creates a batch writer that will write each UserLogin into a CSV file
func NewUserLoginCSVWriter(w io.Writer, dedupers ...UserLoginCSVDeduper) (chan UserLogin, chan bool, error) {
	return NewUserLoginCSVWriterSize(w, UserLoginCSVDefaultSize, dedupers...)
}

// NewUserLoginCSVWriterDir creates a batch writer that will write each UserLogin into a CSV file named user_login.csv.gz in dir
func NewUserLoginCSVWriterDir(dir string, dedupers ...UserLoginCSVDeduper) (chan UserLogin, chan bool, error) {
	return NewUserLoginCSVWriterFile(filepath.Join(dir, "user_login.csv.gz"), dedupers...)
}

// NewUserLoginCSVWriterFile creates a batch writer that will write each UserLogin into a CSV file
func NewUserLoginCSVWriterFile(fn string, dedupers ...UserLoginCSVDeduper) (chan UserLogin, chan bool, error) {
	f, err := os.Create(fn)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening CSV file %s. %v", fn, err)
	}
	var fc io.WriteCloser = f
	if filepath.Ext(fn) == ".gz" {
		w, _ := gzip.NewWriterLevel(f, gzip.BestCompression)
		fc = w
	}
	ch, done, err := NewUserLoginCSVWriter(fc, dedupers...)
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

type UserLoginDBAction func(ctx context.Context, db *sql.DB, record UserLogin) error

// NewUserLoginDBWriterSize creates a DB writer that will write each issue into the DB
func NewUserLoginDBWriterSize(ctx context.Context, db *sql.DB, errors chan<- error, size int, actions ...UserLoginDBAction) (chan UserLogin, chan bool, error) {
	ch := make(chan UserLogin, size)
	done := make(chan bool)
	var action UserLoginDBAction
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

// NewUserLoginDBWriter creates a DB writer that will write each issue into the DB
func NewUserLoginDBWriter(ctx context.Context, db *sql.DB, errors chan<- error, actions ...UserLoginDBAction) (chan UserLogin, chan bool, error) {
	return NewUserLoginDBWriterSize(ctx, db, errors, 100, actions...)
}

// UserLoginColumnID is the ID SQL column name for the UserLogin table
const UserLoginColumnID = "id"

// UserLoginEscapedColumnID is the escaped ID SQL column name for the UserLogin table
const UserLoginEscapedColumnID = "`id`"

// UserLoginColumnChecksum is the Checksum SQL column name for the UserLogin table
const UserLoginColumnChecksum = "checksum"

// UserLoginEscapedColumnChecksum is the escaped Checksum SQL column name for the UserLogin table
const UserLoginEscapedColumnChecksum = "`checksum`"

// UserLoginColumnCustomerID is the CustomerID SQL column name for the UserLogin table
const UserLoginColumnCustomerID = "customer_id"

// UserLoginEscapedColumnCustomerID is the escaped CustomerID SQL column name for the UserLogin table
const UserLoginEscapedColumnCustomerID = "`customer_id`"

// UserLoginColumnUserID is the UserID SQL column name for the UserLogin table
const UserLoginColumnUserID = "user_id"

// UserLoginEscapedColumnUserID is the escaped UserID SQL column name for the UserLogin table
const UserLoginEscapedColumnUserID = "`user_id`"

// UserLoginColumnBrowser is the Browser SQL column name for the UserLogin table
const UserLoginColumnBrowser = "browser"

// UserLoginEscapedColumnBrowser is the escaped Browser SQL column name for the UserLogin table
const UserLoginEscapedColumnBrowser = "`browser`"

// UserLoginColumnDate is the Date SQL column name for the UserLogin table
const UserLoginColumnDate = "date"

// UserLoginEscapedColumnDate is the escaped Date SQL column name for the UserLogin table
const UserLoginEscapedColumnDate = "`date`"

// GetID will return the UserLogin ID value
func (t *UserLogin) GetID() string {
	return t.ID
}

// SetID will set the UserLogin ID value
func (t *UserLogin) SetID(v string) {
	t.ID = v
}

// FindUserLoginByID will find a UserLogin by ID
func FindUserLoginByID(ctx context.Context, db *sql.DB, value string) (*UserLogin, error) {
	q := "SELECT `user_login`.`id`,`user_login`.`checksum`,`user_login`.`customer_id`,`user_login`.`user_id`,`user_login`.`browser`,`user_login`.`date` FROM `user_login` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CustomerID sql.NullString
	var _UserID sql.NullString
	var _Browser sql.NullString
	var _Date sql.NullInt64
	err := db.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_UserID,
		&_Browser,
		&_Date,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &UserLogin{}
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
	if _Browser.Valid {
		t.SetBrowser(_Browser.String)
	}
	if _Date.Valid {
		t.SetDate(_Date.Int64)
	}
	return t, nil
}

// FindUserLoginByIDTx will find a UserLogin by ID using the provided transaction
func FindUserLoginByIDTx(ctx context.Context, tx *sql.Tx, value string) (*UserLogin, error) {
	q := "SELECT `user_login`.`id`,`user_login`.`checksum`,`user_login`.`customer_id`,`user_login`.`user_id`,`user_login`.`browser`,`user_login`.`date` FROM `user_login` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CustomerID sql.NullString
	var _UserID sql.NullString
	var _Browser sql.NullString
	var _Date sql.NullInt64
	err := tx.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_UserID,
		&_Browser,
		&_Date,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &UserLogin{}
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
	if _Browser.Valid {
		t.SetBrowser(_Browser.String)
	}
	if _Date.Valid {
		t.SetDate(_Date.Int64)
	}
	return t, nil
}

// GetChecksum will return the UserLogin Checksum value
func (t *UserLogin) GetChecksum() string {
	if t.Checksum == nil {
		return ""
	}
	return *t.Checksum
}

// SetChecksum will set the UserLogin Checksum value
func (t *UserLogin) SetChecksum(v string) {
	t.Checksum = &v
}

// GetCustomerID will return the UserLogin CustomerID value
func (t *UserLogin) GetCustomerID() string {
	return t.CustomerID
}

// SetCustomerID will set the UserLogin CustomerID value
func (t *UserLogin) SetCustomerID(v string) {
	t.CustomerID = v
}

// GetUserID will return the UserLogin UserID value
func (t *UserLogin) GetUserID() string {
	return t.UserID
}

// SetUserID will set the UserLogin UserID value
func (t *UserLogin) SetUserID(v string) {
	t.UserID = v
}

// GetBrowser will return the UserLogin Browser value
func (t *UserLogin) GetBrowser() string {
	if t.Browser == nil {
		return ""
	}
	return *t.Browser
}

// SetBrowser will set the UserLogin Browser value
func (t *UserLogin) SetBrowser(v string) {
	t.Browser = &v
}

// GetDate will return the UserLogin Date value
func (t *UserLogin) GetDate() int64 {
	return t.Date
}

// SetDate will set the UserLogin Date value
func (t *UserLogin) SetDate(v int64) {
	t.Date = v
}

func (t *UserLogin) toTimestamp(value time.Time) *timestamp.Timestamp {
	ts, _ := ptypes.TimestampProto(value)
	return ts
}

// DBCreateUserLoginTable will create the UserLogin table
func DBCreateUserLoginTable(ctx context.Context, db *sql.DB) error {
	q := "CREATE TABLE `user_login` (`id` VARCHAR(64) NOT NULL PRIMARY KEY,`checksum` CHAR(64),`customer_id` VARCHAR(64) NOT NULL,`user_id`VARCHAR(64) NOT NULL,`browser`TEXT,`date`BIGINT(20) UNSIGNED NOT NULL) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBCreateUserLoginTableTx will create the UserLogin table using the provided transction
func DBCreateUserLoginTableTx(ctx context.Context, tx *sql.Tx) error {
	q := "CREATE TABLE `user_login` (`id` VARCHAR(64) NOT NULL PRIMARY KEY,`checksum` CHAR(64),`customer_id` VARCHAR(64) NOT NULL,`user_id`VARCHAR(64) NOT NULL,`browser`TEXT,`date`BIGINT(20) UNSIGNED NOT NULL) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// DBDropUserLoginTable will drop the UserLogin table
func DBDropUserLoginTable(ctx context.Context, db *sql.DB) error {
	q := "DROP TABLE IF EXISTS `user_login`"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBDropUserLoginTableTx will drop the UserLogin table using the provided transaction
func DBDropUserLoginTableTx(ctx context.Context, tx *sql.Tx) error {
	q := "DROP TABLE IF EXISTS `user_login`"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// CalculateChecksum will calculate a checksum of the SHA1 of all field values
func (t *UserLogin) CalculateChecksum() string {
	return orm.HashStrings(
		orm.ToString(t.ID),
		orm.ToString(t.CustomerID),
		orm.ToString(t.UserID),
		orm.ToString(t.Browser),
		orm.ToString(t.Date),
	)
}

// DBCreate will create a new UserLogin record in the database
func (t *UserLogin) DBCreate(ctx context.Context, db *sql.DB) (sql.Result, error) {
	q := "INSERT INTO `user_login` (`user_login`.`id`,`user_login`.`checksum`,`user_login`.`customer_id`,`user_login`.`user_id`,`user_login`.`browser`,`user_login`.`date`) VALUES (?,?,?,?,?,?)"
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
		orm.ToSQLString(t.Browser),
		orm.ToSQLInt64(t.Date),
	)
}

// DBCreateTx will create a new UserLogin record in the database using the provided transaction
func (t *UserLogin) DBCreateTx(ctx context.Context, tx *sql.Tx) (sql.Result, error) {
	q := "INSERT INTO `user_login` (`user_login`.`id`,`user_login`.`checksum`,`user_login`.`customer_id`,`user_login`.`user_id`,`user_login`.`browser`,`user_login`.`date`) VALUES (?,?,?,?,?,?)"
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
		orm.ToSQLString(t.Browser),
		orm.ToSQLInt64(t.Date),
	)
}

// DBCreateIgnoreDuplicate will upsert the UserLogin record in the database
func (t *UserLogin) DBCreateIgnoreDuplicate(ctx context.Context, db *sql.DB) (sql.Result, error) {
	q := "INSERT INTO `user_login` (`user_login`.`id`,`user_login`.`checksum`,`user_login`.`customer_id`,`user_login`.`user_id`,`user_login`.`browser`,`user_login`.`date`) VALUES (?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
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
		orm.ToSQLString(t.Browser),
		orm.ToSQLInt64(t.Date),
	)
}

// DBCreateIgnoreDuplicateTx will upsert the UserLogin record in the database using the provided transaction
func (t *UserLogin) DBCreateIgnoreDuplicateTx(ctx context.Context, tx *sql.Tx) (sql.Result, error) {
	q := "INSERT INTO `user_login` (`user_login`.`id`,`user_login`.`checksum`,`user_login`.`customer_id`,`user_login`.`user_id`,`user_login`.`browser`,`user_login`.`date`) VALUES (?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
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
		orm.ToSQLString(t.Browser),
		orm.ToSQLInt64(t.Date),
	)
}

// DeleteAllUserLogins deletes all UserLogin records in the database with optional filters
func DeleteAllUserLogins(ctx context.Context, db *sql.DB, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(UserLoginTableName),
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

// DeleteAllUserLoginsTx deletes all UserLogin records in the database with optional filters using the provided transaction
func DeleteAllUserLoginsTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(UserLoginTableName),
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

// DBDelete will delete this UserLogin record in the database
func (t *UserLogin) DBDelete(ctx context.Context, db *sql.DB) (bool, error) {
	q := "DELETE FROM `user_login` WHERE `id` = ?"
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

// DBDeleteTx will delete this UserLogin record in the database using the provided transaction
func (t *UserLogin) DBDeleteTx(ctx context.Context, tx *sql.Tx) (bool, error) {
	q := "DELETE FROM `user_login` WHERE `id` = ?"
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

// DBUpdate will update the UserLogin record in the database
func (t *UserLogin) DBUpdate(ctx context.Context, db *sql.DB) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `user_login` SET `checksum`=?,`customer_id`=?,`user_id`=?,`browser`=?,`date`=? WHERE `id`=?"
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.UserID),
		orm.ToSQLString(t.Browser),
		orm.ToSQLInt64(t.Date),
		orm.ToSQLString(t.ID),
	)
}

// DBUpdateTx will update the UserLogin record in the database using the provided transaction
func (t *UserLogin) DBUpdateTx(ctx context.Context, tx *sql.Tx) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `user_login` SET `checksum`=?,`customer_id`=?,`user_id`=?,`browser`=?,`date`=? WHERE `id`=?"
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.UserID),
		orm.ToSQLString(t.Browser),
		orm.ToSQLInt64(t.Date),
		orm.ToSQLString(t.ID),
	)
}

// DBUpsert will upsert the UserLogin record in the database
func (t *UserLogin) DBUpsert(ctx context.Context, db *sql.DB, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `user_login` (`user_login`.`id`,`user_login`.`checksum`,`user_login`.`customer_id`,`user_login`.`user_id`,`user_login`.`browser`,`user_login`.`date`) VALUES (?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `user_login` (`user_login`.`id`,`user_login`.`checksum`,`user_login`.`customer_id`,`user_login`.`user_id`,`user_login`.`browser`,`user_login`.`date`) VALUES (?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`customer_id`=VALUES(`customer_id`),`user_id`=VALUES(`user_id`),`browser`=VALUES(`browser`),`date`=VALUES(`date`)"
	}
	r, err := db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.UserID),
		orm.ToSQLString(t.Browser),
		orm.ToSQLInt64(t.Date),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBUpsertTx will upsert the UserLogin record in the database using the provided transaction
func (t *UserLogin) DBUpsertTx(ctx context.Context, tx *sql.Tx, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `user_login` (`user_login`.`id`,`user_login`.`checksum`,`user_login`.`customer_id`,`user_login`.`user_id`,`user_login`.`browser`,`user_login`.`date`) VALUES (?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `user_login` (`user_login`.`id`,`user_login`.`checksum`,`user_login`.`customer_id`,`user_login`.`user_id`,`user_login`.`browser`,`user_login`.`date`) VALUES (?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`customer_id`=VALUES(`customer_id`),`user_id`=VALUES(`user_id`),`browser`=VALUES(`browser`),`date`=VALUES(`date`)"
	}
	r, err := tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.UserID),
		orm.ToSQLString(t.Browser),
		orm.ToSQLInt64(t.Date),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBFindOne will find a UserLogin record in the database with the primary key
func (t *UserLogin) DBFindOne(ctx context.Context, db *sql.DB, value string) (bool, error) {
	q := "SELECT `user_login`.`id`,`user_login`.`checksum`,`user_login`.`customer_id`,`user_login`.`user_id`,`user_login`.`browser`,`user_login`.`date` FROM `user_login` WHERE `id` = ? LIMIT 1"
	row := db.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CustomerID sql.NullString
	var _UserID sql.NullString
	var _Browser sql.NullString
	var _Date sql.NullInt64
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_UserID,
		&_Browser,
		&_Date,
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
	if _Browser.Valid {
		t.SetBrowser(_Browser.String)
	}
	if _Date.Valid {
		t.SetDate(_Date.Int64)
	}
	return true, nil
}

// DBFindOneTx will find a UserLogin record in the database with the primary key using the provided transaction
func (t *UserLogin) DBFindOneTx(ctx context.Context, tx *sql.Tx, value string) (bool, error) {
	q := "SELECT `user_login`.`id`,`user_login`.`checksum`,`user_login`.`customer_id`,`user_login`.`user_id`,`user_login`.`browser`,`user_login`.`date` FROM `user_login` WHERE `id` = ? LIMIT 1"
	row := tx.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CustomerID sql.NullString
	var _UserID sql.NullString
	var _Browser sql.NullString
	var _Date sql.NullInt64
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_UserID,
		&_Browser,
		&_Date,
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
	if _Browser.Valid {
		t.SetBrowser(_Browser.String)
	}
	if _Date.Valid {
		t.SetDate(_Date.Int64)
	}
	return true, nil
}

// FindUserLogins will find a UserLogin record in the database with the provided parameters
func FindUserLogins(ctx context.Context, db *sql.DB, _params ...interface{}) ([]*UserLogin, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("customer_id"),
		orm.Column("user_id"),
		orm.Column("browser"),
		orm.Column("date"),
		orm.Table(UserLoginTableName),
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
	results := make([]*UserLogin, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _UserID sql.NullString
		var _Browser sql.NullString
		var _Date sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_UserID,
			&_Browser,
			&_Date,
		)
		if err != nil {
			return nil, err
		}
		t := &UserLogin{}
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
		if _Browser.Valid {
			t.SetBrowser(_Browser.String)
		}
		if _Date.Valid {
			t.SetDate(_Date.Int64)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindUserLoginsTx will find a UserLogin record in the database with the provided parameters using the provided transaction
func FindUserLoginsTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) ([]*UserLogin, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("customer_id"),
		orm.Column("user_id"),
		orm.Column("browser"),
		orm.Column("date"),
		orm.Table(UserLoginTableName),
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
	results := make([]*UserLogin, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _UserID sql.NullString
		var _Browser sql.NullString
		var _Date sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_UserID,
			&_Browser,
			&_Date,
		)
		if err != nil {
			return nil, err
		}
		t := &UserLogin{}
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
		if _Browser.Valid {
			t.SetBrowser(_Browser.String)
		}
		if _Date.Valid {
			t.SetDate(_Date.Int64)
		}
		results = append(results, t)
	}
	return results, nil
}

// DBFind will find a UserLogin record in the database with the provided parameters
func (t *UserLogin) DBFind(ctx context.Context, db *sql.DB, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("customer_id"),
		orm.Column("user_id"),
		orm.Column("browser"),
		orm.Column("date"),
		orm.Table(UserLoginTableName),
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
	var _Browser sql.NullString
	var _Date sql.NullInt64
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_UserID,
		&_Browser,
		&_Date,
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
	if _Browser.Valid {
		t.SetBrowser(_Browser.String)
	}
	if _Date.Valid {
		t.SetDate(_Date.Int64)
	}
	return true, nil
}

// DBFindTx will find a UserLogin record in the database with the provided parameters using the provided transaction
func (t *UserLogin) DBFindTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("customer_id"),
		orm.Column("user_id"),
		orm.Column("browser"),
		orm.Column("date"),
		orm.Table(UserLoginTableName),
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
	var _Browser sql.NullString
	var _Date sql.NullInt64
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_UserID,
		&_Browser,
		&_Date,
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
	if _Browser.Valid {
		t.SetBrowser(_Browser.String)
	}
	if _Date.Valid {
		t.SetDate(_Date.Int64)
	}
	return true, nil
}

// CountUserLogins will find the count of UserLogin records in the database
func CountUserLogins(ctx context.Context, db *sql.DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(UserLoginTableName),
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

// CountUserLoginsTx will find the count of UserLogin records in the database using the provided transaction
func CountUserLoginsTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(UserLoginTableName),
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

// DBCount will find the count of UserLogin records in the database
func (t *UserLogin) DBCount(ctx context.Context, db *sql.DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(UserLoginTableName),
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

// DBCountTx will find the count of UserLogin records in the database using the provided transaction
func (t *UserLogin) DBCountTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(UserLoginTableName),
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

// DBExists will return true if the UserLogin record exists in the database
func (t *UserLogin) DBExists(ctx context.Context, db *sql.DB) (bool, error) {
	q := "SELECT `id` FROM `user_login` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := db.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// DBExistsTx will return true if the UserLogin record exists in the database using the provided transaction
func (t *UserLogin) DBExistsTx(ctx context.Context, tx *sql.Tx) (bool, error) {
	q := "SELECT `id` FROM `user_login` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := tx.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// PrimaryKeyColumn returns the column name for the primary key
func (t *UserLogin) PrimaryKeyColumn() string {
	return UserLoginColumnID
}

// PrimaryKeyColumnType returns the primary key column Go type as a string
func (t *UserLogin) PrimaryKeyColumnType() string {
	return "string"
}

// PrimaryKey returns the primary key column value
func (t *UserLogin) PrimaryKey() interface{} {
	return t.ID
}
