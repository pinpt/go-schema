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

var _ Model = (*Customer)(nil)
var _ CSVWriter = (*Customer)(nil)
var _ JSONWriter = (*Customer)(nil)

// CustomerTableName is the name of the table in SQL
const CustomerTableName = "customer"

var CustomerColumns = []string{
	"id",
	"name",
	"active",
	"created_at",
	"updated_at",
}

// Customer table
type Customer struct {
	Active    bool   `json:"active"`
	CreatedAt int64  `json:"created_at"`
	ID        string `json:"id"`
	Name      string `json:"name"`
	UpdatedAt *int64 `json:"updated_at,omitempty"`
}

// TableName returns the SQL table name for Customer and satifies the Model interface
func (t *Customer) TableName() string {
	return CustomerTableName
}

// ToCSV will serialize the Customer instance to a CSV compatible array of strings
func (t *Customer) ToCSV() []string {
	return []string{
		t.ID,
		t.Name,
		toCSVBool(t.Active),
		toCSVString(t.CreatedAt),
		toCSVString(t.UpdatedAt),
	}
}

// WriteCSV will serialize the Customer instance to the writer as CSV and satisfies the CSVWriter interface
func (t *Customer) WriteCSV(w *csv.Writer) error {
	return w.Write(t.ToCSV())
}

// WriteJSON will serialize the Customer instance to the writer as JSON and satisfies the JSONWriter interface
func (t *Customer) WriteJSON(w io.Writer, indent ...bool) error {
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

// NewCustomerReader creates a JSON reader which can read in Customer objects serialized as JSON either as an array, single object or json new lines
// and writes each Customer to the channel provided
func NewCustomerReader(r io.Reader, ch chan<- Customer) error {
	return orm.Deserialize(r, func(buf json.RawMessage) error {
		dec := json.NewDecoder(bytes.NewBuffer(buf))
		e := Customer{}
		if err := dec.Decode(&e); err != nil {
			return err
		}
		ch <- e
		return nil
	})
}

// NewCSVCustomerReaderDir will read the reader as CSV and emit each record to the channel provided
func NewCSVCustomerReader(r io.Reader, ch chan<- Customer) error {
	cr := csv.NewReader(r)
	for {
		record, err := cr.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		ch <- Customer{
			ID:        record[0],
			Name:      record[1],
			Active:    fromCSVBool(record[2]),
			CreatedAt: fromCSVInt64(record[3]),
			UpdatedAt: fromCSVInt64Pointer(record[4]),
		}
	}
	return nil
}

// NewCSVCustomerReaderFile will read the file as a CSV and emit each record to the channel provided
func NewCSVCustomerReaderFile(fp string, ch chan<- Customer) error {
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
	return NewCSVCustomerReader(fc, ch)
}

// NewCSVCustomerReaderDir will read the customer.csv.gz file as a CSV and emit each record to the channel provided
func NewCSVCustomerReaderDir(dir string, ch chan<- Customer) error {
	return NewCSVCustomerReaderFile(filepath.Join(dir, "customer.csv.gz"), ch)
}

// CustomerCSVDeduper is a function callback which takes the existing value (a) and the new value (b)
// and the return value should be the one to use (or a new one, if applicable). return nil
// to skip processing of this record
type CustomerCSVDeduper func(a Customer, b Customer) *Customer

// CustomerCSVDedupeDisabled is set on whether the CSV writer should de-dupe values by key
var CustomerCSVDedupeDisabled bool

// NewCustomerCSVWriterSize creates a batch writer that will write each Customer into a CSV file
func NewCustomerCSVWriterSize(w io.Writer, size int, dedupers ...CustomerCSVDeduper) (chan Customer, chan bool, error) {
	cw := csv.NewWriter(w)
	ch := make(chan Customer, size)
	done := make(chan bool)
	go func() {
		defer func() { done <- true }()
		dodedupe := !CustomerCSVDedupeDisabled
		var kv map[string]*Customer
		var deduper CustomerCSVDeduper
		if dedupers != nil && len(dedupers) > 0 {
			deduper = dedupers[0]
			dodedupe = true
		}
		if dodedupe {
			kv = make(map[string]*Customer)
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

// CustomerCSVDefaultSize is the default channel buffer size if not provided
var CustomerCSVDefaultSize = 100

// NewCustomerCSVWriter creates a batch writer that will write each Customer into a CSV file
func NewCustomerCSVWriter(w io.Writer, dedupers ...CustomerCSVDeduper) (chan Customer, chan bool, error) {
	return NewCustomerCSVWriterSize(w, CustomerCSVDefaultSize, dedupers...)
}

// NewCustomerCSVWriterDir creates a batch writer that will write each Customer into a CSV file named customer.csv.gz in dir
func NewCustomerCSVWriterDir(dir string, dedupers ...CustomerCSVDeduper) (chan Customer, chan bool, error) {
	return NewCustomerCSVWriterFile(filepath.Join(dir, "customer.csv.gz"), dedupers...)
}

// NewCustomerCSVWriterFile creates a batch writer that will write each Customer into a CSV file
func NewCustomerCSVWriterFile(fn string, dedupers ...CustomerCSVDeduper) (chan Customer, chan bool, error) {
	f, err := os.Create(fn)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening CSV file %s. %v", fn, err)
	}
	var fc io.WriteCloser = f
	if filepath.Ext(fn) == ".gz" {
		w, _ := gzip.NewWriterLevel(f, gzip.BestCompression)
		fc = w
	}
	ch, done, err := NewCustomerCSVWriter(fc, dedupers...)
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

type CustomerDBAction func(ctx context.Context, db *sql.DB, record Customer) error

// NewCustomerDBWriterSize creates a DB writer that will write each issue into the DB
func NewCustomerDBWriterSize(ctx context.Context, db *sql.DB, errors chan<- error, size int, actions ...CustomerDBAction) (chan Customer, chan bool, error) {
	ch := make(chan Customer, size)
	done := make(chan bool)
	var action CustomerDBAction
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

// NewCustomerDBWriter creates a DB writer that will write each issue into the DB
func NewCustomerDBWriter(ctx context.Context, db *sql.DB, errors chan<- error, actions ...CustomerDBAction) (chan Customer, chan bool, error) {
	return NewCustomerDBWriterSize(ctx, db, errors, 100, actions...)
}

// CustomerColumnID is the ID SQL column name for the Customer table
const CustomerColumnID = "id"

// CustomerEscapedColumnID is the escaped ID SQL column name for the Customer table
const CustomerEscapedColumnID = "`id`"

// CustomerColumnName is the Name SQL column name for the Customer table
const CustomerColumnName = "name"

// CustomerEscapedColumnName is the escaped Name SQL column name for the Customer table
const CustomerEscapedColumnName = "`name`"

// CustomerColumnActive is the Active SQL column name for the Customer table
const CustomerColumnActive = "active"

// CustomerEscapedColumnActive is the escaped Active SQL column name for the Customer table
const CustomerEscapedColumnActive = "`active`"

// CustomerColumnCreatedAt is the CreatedAt SQL column name for the Customer table
const CustomerColumnCreatedAt = "created_at"

// CustomerEscapedColumnCreatedAt is the escaped CreatedAt SQL column name for the Customer table
const CustomerEscapedColumnCreatedAt = "`created_at`"

// CustomerColumnUpdatedAt is the UpdatedAt SQL column name for the Customer table
const CustomerColumnUpdatedAt = "updated_at"

// CustomerEscapedColumnUpdatedAt is the escaped UpdatedAt SQL column name for the Customer table
const CustomerEscapedColumnUpdatedAt = "`updated_at`"

// GetID will return the Customer ID value
func (t *Customer) GetID() string {
	return t.ID
}

// SetID will set the Customer ID value
func (t *Customer) SetID(v string) {
	t.ID = v
}

// FindCustomerByID will find a Customer by ID
func FindCustomerByID(ctx context.Context, db *sql.DB, value string) (*Customer, error) {
	q := "SELECT `customer`.`id`,`customer`.`name`,`customer`.`active`,`customer`.`created_at`,`customer`.`updated_at` FROM `customer` WHERE `id` = ?"
	var _ID sql.NullString
	var _Name sql.NullString
	var _Active sql.NullBool
	var _CreatedAt sql.NullInt64
	var _UpdatedAt sql.NullInt64
	err := db.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Name,
		&_Active,
		&_CreatedAt,
		&_UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &Customer{}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _Active.Valid {
		t.SetActive(_Active.Bool)
	}
	if _CreatedAt.Valid {
		t.SetCreatedAt(_CreatedAt.Int64)
	}
	if _UpdatedAt.Valid {
		t.SetUpdatedAt(_UpdatedAt.Int64)
	}
	return t, nil
}

// FindCustomerByIDTx will find a Customer by ID using the provided transaction
func FindCustomerByIDTx(ctx context.Context, tx *sql.Tx, value string) (*Customer, error) {
	q := "SELECT `customer`.`id`,`customer`.`name`,`customer`.`active`,`customer`.`created_at`,`customer`.`updated_at` FROM `customer` WHERE `id` = ?"
	var _ID sql.NullString
	var _Name sql.NullString
	var _Active sql.NullBool
	var _CreatedAt sql.NullInt64
	var _UpdatedAt sql.NullInt64
	err := tx.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Name,
		&_Active,
		&_CreatedAt,
		&_UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &Customer{}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _Active.Valid {
		t.SetActive(_Active.Bool)
	}
	if _CreatedAt.Valid {
		t.SetCreatedAt(_CreatedAt.Int64)
	}
	if _UpdatedAt.Valid {
		t.SetUpdatedAt(_UpdatedAt.Int64)
	}
	return t, nil
}

// GetName will return the Customer Name value
func (t *Customer) GetName() string {
	return t.Name
}

// SetName will set the Customer Name value
func (t *Customer) SetName(v string) {
	t.Name = v
}

// GetActive will return the Customer Active value
func (t *Customer) GetActive() bool {
	return t.Active
}

// SetActive will set the Customer Active value
func (t *Customer) SetActive(v bool) {
	t.Active = v
}

// GetCreatedAt will return the Customer CreatedAt value
func (t *Customer) GetCreatedAt() int64 {
	return t.CreatedAt
}

// SetCreatedAt will set the Customer CreatedAt value
func (t *Customer) SetCreatedAt(v int64) {
	t.CreatedAt = v
}

// GetUpdatedAt will return the Customer UpdatedAt value
func (t *Customer) GetUpdatedAt() int64 {
	if t.UpdatedAt == nil {
		return int64(0)
	}
	return *t.UpdatedAt
}

// SetUpdatedAt will set the Customer UpdatedAt value
func (t *Customer) SetUpdatedAt(v int64) {
	t.UpdatedAt = &v
}

func (t *Customer) toTimestamp(value time.Time) *timestamp.Timestamp {
	ts, _ := ptypes.TimestampProto(value)
	return ts
}

// DBCreateCustomerTable will create the Customer table
func DBCreateCustomerTable(ctx context.Context, db *sql.DB) error {
	q := "CREATE TABLE `customer` (`id` VARCHAR(64) NOT NULL PRIMARY KEY,`name` TEXT NOT NULL,`active`TINYINT(1) NOT NULL,`created_at` BIGINT(20) UNSIGNED NOT NULL,`updated_at` BIGINT(20) UNSIGNED) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBCreateCustomerTableTx will create the Customer table using the provided transction
func DBCreateCustomerTableTx(ctx context.Context, tx *sql.Tx) error {
	q := "CREATE TABLE `customer` (`id` VARCHAR(64) NOT NULL PRIMARY KEY,`name` TEXT NOT NULL,`active`TINYINT(1) NOT NULL,`created_at` BIGINT(20) UNSIGNED NOT NULL,`updated_at` BIGINT(20) UNSIGNED) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// DBDropCustomerTable will drop the Customer table
func DBDropCustomerTable(ctx context.Context, db *sql.DB) error {
	q := "DROP TABLE IF EXISTS `customer`"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBDropCustomerTableTx will drop the Customer table using the provided transaction
func DBDropCustomerTableTx(ctx context.Context, tx *sql.Tx) error {
	q := "DROP TABLE IF EXISTS `customer`"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// DBCreate will create a new Customer record in the database
func (t *Customer) DBCreate(ctx context.Context, db *sql.DB) (sql.Result, error) {
	q := "INSERT INTO `customer` (`customer`.`id`,`customer`.`name`,`customer`.`active`,`customer`.`created_at`,`customer`.`updated_at`) VALUES (?,?,?,?,?)"
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Name),
		orm.ToSQLBool(t.Active),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
	)
}

// DBCreateTx will create a new Customer record in the database using the provided transaction
func (t *Customer) DBCreateTx(ctx context.Context, tx *sql.Tx) (sql.Result, error) {
	q := "INSERT INTO `customer` (`customer`.`id`,`customer`.`name`,`customer`.`active`,`customer`.`created_at`,`customer`.`updated_at`) VALUES (?,?,?,?,?)"
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Name),
		orm.ToSQLBool(t.Active),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
	)
}

// DBCreateIgnoreDuplicate will upsert the Customer record in the database
func (t *Customer) DBCreateIgnoreDuplicate(ctx context.Context, db *sql.DB) (sql.Result, error) {
	q := "INSERT INTO `customer` (`customer`.`id`,`customer`.`name`,`customer`.`active`,`customer`.`created_at`,`customer`.`updated_at`) VALUES (?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Name),
		orm.ToSQLBool(t.Active),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
	)
}

// DBCreateIgnoreDuplicateTx will upsert the Customer record in the database using the provided transaction
func (t *Customer) DBCreateIgnoreDuplicateTx(ctx context.Context, tx *sql.Tx) (sql.Result, error) {
	q := "INSERT INTO `customer` (`customer`.`id`,`customer`.`name`,`customer`.`active`,`customer`.`created_at`,`customer`.`updated_at`) VALUES (?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Name),
		orm.ToSQLBool(t.Active),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
	)
}

// DeleteAllCustomers deletes all Customer records in the database with optional filters
func DeleteAllCustomers(ctx context.Context, db *sql.DB, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(CustomerTableName),
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

// DeleteAllCustomersTx deletes all Customer records in the database with optional filters using the provided transaction
func DeleteAllCustomersTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(CustomerTableName),
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

// DBDelete will delete this Customer record in the database
func (t *Customer) DBDelete(ctx context.Context, db *sql.DB) (bool, error) {
	q := "DELETE FROM `customer` WHERE `id` = ?"
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

// DBDeleteTx will delete this Customer record in the database using the provided transaction
func (t *Customer) DBDeleteTx(ctx context.Context, tx *sql.Tx) (bool, error) {
	q := "DELETE FROM `customer` WHERE `id` = ?"
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

// DBUpdate will update the Customer record in the database
func (t *Customer) DBUpdate(ctx context.Context, db *sql.DB) (sql.Result, error) {
	q := "UPDATE `customer` SET `name`=?,`active`=?,`created_at`=?,`updated_at`=? WHERE `id`=?"
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.Name),
		orm.ToSQLBool(t.Active),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
		orm.ToSQLString(t.ID),
	)
}

// DBUpdateTx will update the Customer record in the database using the provided transaction
func (t *Customer) DBUpdateTx(ctx context.Context, tx *sql.Tx) (sql.Result, error) {
	q := "UPDATE `customer` SET `name`=?,`active`=?,`created_at`=?,`updated_at`=? WHERE `id`=?"
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.Name),
		orm.ToSQLBool(t.Active),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
		orm.ToSQLString(t.ID),
	)
}

// DBUpsert will upsert the Customer record in the database
func (t *Customer) DBUpsert(ctx context.Context, db *sql.DB, conditions ...interface{}) (bool, bool, error) {
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `customer` (`customer`.`id`,`customer`.`name`,`customer`.`active`,`customer`.`created_at`,`customer`.`updated_at`) VALUES (?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `customer` (`customer`.`id`,`customer`.`name`,`customer`.`active`,`customer`.`created_at`,`customer`.`updated_at`) VALUES (?,?,?,?,?) ON DUPLICATE KEY UPDATE `name`=VALUES(`name`),`active`=VALUES(`active`),`created_at`=VALUES(`created_at`),`updated_at`=VALUES(`updated_at`)"
	}
	r, err := db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Name),
		orm.ToSQLBool(t.Active),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBUpsertTx will upsert the Customer record in the database using the provided transaction
func (t *Customer) DBUpsertTx(ctx context.Context, tx *sql.Tx, conditions ...interface{}) (bool, bool, error) {
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `customer` (`customer`.`id`,`customer`.`name`,`customer`.`active`,`customer`.`created_at`,`customer`.`updated_at`) VALUES (?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `customer` (`customer`.`id`,`customer`.`name`,`customer`.`active`,`customer`.`created_at`,`customer`.`updated_at`) VALUES (?,?,?,?,?) ON DUPLICATE KEY UPDATE `name`=VALUES(`name`),`active`=VALUES(`active`),`created_at`=VALUES(`created_at`),`updated_at`=VALUES(`updated_at`)"
	}
	r, err := tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Name),
		orm.ToSQLBool(t.Active),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBFindOne will find a Customer record in the database with the primary key
func (t *Customer) DBFindOne(ctx context.Context, db *sql.DB, value string) (bool, error) {
	q := "SELECT `customer`.`id`,`customer`.`name`,`customer`.`active`,`customer`.`created_at`,`customer`.`updated_at` FROM `customer` WHERE `id` = ? LIMIT 1"
	row := db.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Name sql.NullString
	var _Active sql.NullBool
	var _CreatedAt sql.NullInt64
	var _UpdatedAt sql.NullInt64
	err := row.Scan(
		&_ID,
		&_Name,
		&_Active,
		&_CreatedAt,
		&_UpdatedAt,
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
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _Active.Valid {
		t.SetActive(_Active.Bool)
	}
	if _CreatedAt.Valid {
		t.SetCreatedAt(_CreatedAt.Int64)
	}
	if _UpdatedAt.Valid {
		t.SetUpdatedAt(_UpdatedAt.Int64)
	}
	return true, nil
}

// DBFindOneTx will find a Customer record in the database with the primary key using the provided transaction
func (t *Customer) DBFindOneTx(ctx context.Context, tx *sql.Tx, value string) (bool, error) {
	q := "SELECT `customer`.`id`,`customer`.`name`,`customer`.`active`,`customer`.`created_at`,`customer`.`updated_at` FROM `customer` WHERE `id` = ? LIMIT 1"
	row := tx.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Name sql.NullString
	var _Active sql.NullBool
	var _CreatedAt sql.NullInt64
	var _UpdatedAt sql.NullInt64
	err := row.Scan(
		&_ID,
		&_Name,
		&_Active,
		&_CreatedAt,
		&_UpdatedAt,
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
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _Active.Valid {
		t.SetActive(_Active.Bool)
	}
	if _CreatedAt.Valid {
		t.SetCreatedAt(_CreatedAt.Int64)
	}
	if _UpdatedAt.Valid {
		t.SetUpdatedAt(_UpdatedAt.Int64)
	}
	return true, nil
}

// FindCustomers will find a Customer record in the database with the provided parameters
func FindCustomers(ctx context.Context, db *sql.DB, _params ...interface{}) ([]*Customer, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("name"),
		orm.Column("active"),
		orm.Column("created_at"),
		orm.Column("updated_at"),
		orm.Table(CustomerTableName),
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
	results := make([]*Customer, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Name sql.NullString
		var _Active sql.NullBool
		var _CreatedAt sql.NullInt64
		var _UpdatedAt sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Name,
			&_Active,
			&_CreatedAt,
			&_UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		t := &Customer{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _Active.Valid {
			t.SetActive(_Active.Bool)
		}
		if _CreatedAt.Valid {
			t.SetCreatedAt(_CreatedAt.Int64)
		}
		if _UpdatedAt.Valid {
			t.SetUpdatedAt(_UpdatedAt.Int64)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindCustomersTx will find a Customer record in the database with the provided parameters using the provided transaction
func FindCustomersTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) ([]*Customer, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("name"),
		orm.Column("active"),
		orm.Column("created_at"),
		orm.Column("updated_at"),
		orm.Table(CustomerTableName),
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
	results := make([]*Customer, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Name sql.NullString
		var _Active sql.NullBool
		var _CreatedAt sql.NullInt64
		var _UpdatedAt sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Name,
			&_Active,
			&_CreatedAt,
			&_UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		t := &Customer{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _Active.Valid {
			t.SetActive(_Active.Bool)
		}
		if _CreatedAt.Valid {
			t.SetCreatedAt(_CreatedAt.Int64)
		}
		if _UpdatedAt.Valid {
			t.SetUpdatedAt(_UpdatedAt.Int64)
		}
		results = append(results, t)
	}
	return results, nil
}

// DBFind will find a Customer record in the database with the provided parameters
func (t *Customer) DBFind(ctx context.Context, db *sql.DB, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("name"),
		orm.Column("active"),
		orm.Column("created_at"),
		orm.Column("updated_at"),
		orm.Table(CustomerTableName),
	}
	if len(_params) > 0 {
		for _, param := range _params {
			params = append(params, param)
		}
	}
	q, p := orm.BuildQuery(params...)
	row := db.QueryRowContext(ctx, q, p...)
	var _ID sql.NullString
	var _Name sql.NullString
	var _Active sql.NullBool
	var _CreatedAt sql.NullInt64
	var _UpdatedAt sql.NullInt64
	err := row.Scan(
		&_ID,
		&_Name,
		&_Active,
		&_CreatedAt,
		&_UpdatedAt,
	)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _Active.Valid {
		t.SetActive(_Active.Bool)
	}
	if _CreatedAt.Valid {
		t.SetCreatedAt(_CreatedAt.Int64)
	}
	if _UpdatedAt.Valid {
		t.SetUpdatedAt(_UpdatedAt.Int64)
	}
	return true, nil
}

// DBFindTx will find a Customer record in the database with the provided parameters using the provided transaction
func (t *Customer) DBFindTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("name"),
		orm.Column("active"),
		orm.Column("created_at"),
		orm.Column("updated_at"),
		orm.Table(CustomerTableName),
	}
	if len(_params) > 0 {
		for _, param := range _params {
			params = append(params, param)
		}
	}
	q, p := orm.BuildQuery(params...)
	row := tx.QueryRowContext(ctx, q, p...)
	var _ID sql.NullString
	var _Name sql.NullString
	var _Active sql.NullBool
	var _CreatedAt sql.NullInt64
	var _UpdatedAt sql.NullInt64
	err := row.Scan(
		&_ID,
		&_Name,
		&_Active,
		&_CreatedAt,
		&_UpdatedAt,
	)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _Active.Valid {
		t.SetActive(_Active.Bool)
	}
	if _CreatedAt.Valid {
		t.SetCreatedAt(_CreatedAt.Int64)
	}
	if _UpdatedAt.Valid {
		t.SetUpdatedAt(_UpdatedAt.Int64)
	}
	return true, nil
}

// CountCustomers will find the count of Customer records in the database
func CountCustomers(ctx context.Context, db *sql.DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(CustomerTableName),
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

// CountCustomersTx will find the count of Customer records in the database using the provided transaction
func CountCustomersTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(CustomerTableName),
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

// DBCount will find the count of Customer records in the database
func (t *Customer) DBCount(ctx context.Context, db *sql.DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(CustomerTableName),
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

// DBCountTx will find the count of Customer records in the database using the provided transaction
func (t *Customer) DBCountTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(CustomerTableName),
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

// DBExists will return true if the Customer record exists in the database
func (t *Customer) DBExists(ctx context.Context, db *sql.DB) (bool, error) {
	q := "SELECT `id` FROM `customer` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := db.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// DBExistsTx will return true if the Customer record exists in the database using the provided transaction
func (t *Customer) DBExistsTx(ctx context.Context, tx *sql.Tx) (bool, error) {
	q := "SELECT `id` FROM `customer` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := tx.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// PrimaryKeyColumn returns the column name for the primary key
func (t *Customer) PrimaryKeyColumn() string {
	return CustomerColumnID
}

// PrimaryKeyColumnType returns the primary key column Go type as a string
func (t *Customer) PrimaryKeyColumnType() string {
	return "string"
}

// PrimaryKey returns the primary key column value
func (t *Customer) PrimaryKey() interface{} {
	return t.ID
}
