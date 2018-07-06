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

var _ Model = (*JiraProjectCategory)(nil)
var _ CSVWriter = (*JiraProjectCategory)(nil)
var _ JSONWriter = (*JiraProjectCategory)(nil)
var _ Checksum = (*JiraProjectCategory)(nil)

// JiraProjectCategoryTableName is the name of the table in SQL
const JiraProjectCategoryTableName = "jira_project_category"

var JiraProjectCategoryColumns = []string{
	"id",
	"checksum",
	"name",
	"description",
	"customer_id",
}

// JiraProjectCategory table
type JiraProjectCategory struct {
	Checksum    *string `json:"checksum,omitempty"`
	CustomerID  string  `json:"customer_id"`
	Description *string `json:"description,omitempty"`
	ID          string  `json:"id"`
	Name        *string `json:"name,omitempty"`
}

// TableName returns the SQL table name for JiraProjectCategory and satifies the Model interface
func (t *JiraProjectCategory) TableName() string {
	return JiraProjectCategoryTableName
}

// ToCSV will serialize the JiraProjectCategory instance to a CSV compatible array of strings
func (t *JiraProjectCategory) ToCSV() []string {
	return []string{
		t.ID,
		t.CalculateChecksum(),
		toCSVString(t.Name),
		toCSVString(t.Description),
		t.CustomerID,
	}
}

// WriteCSV will serialize the JiraProjectCategory instance to the writer as CSV and satisfies the CSVWriter interface
func (t *JiraProjectCategory) WriteCSV(w *csv.Writer) error {
	return w.Write(t.ToCSV())
}

// WriteJSON will serialize the JiraProjectCategory instance to the writer as JSON and satisfies the JSONWriter interface
func (t *JiraProjectCategory) WriteJSON(w io.Writer, indent ...bool) error {
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

// NewJiraProjectCategoryReader creates a JSON reader which can read in JiraProjectCategory objects serialized as JSON either as an array, single object or json new lines
// and writes each JiraProjectCategory to the channel provided
func NewJiraProjectCategoryReader(r io.Reader, ch chan<- JiraProjectCategory) error {
	return orm.Deserialize(r, func(buf json.RawMessage) error {
		dec := json.NewDecoder(bytes.NewBuffer(buf))
		e := JiraProjectCategory{}
		if err := dec.Decode(&e); err != nil {
			return err
		}
		ch <- e
		return nil
	})
}

// NewCSVJiraProjectCategoryReaderDir will read the reader as CSV and emit each record to the channel provided
func NewCSVJiraProjectCategoryReader(r io.Reader, ch chan<- JiraProjectCategory) error {
	cr := csv.NewReader(r)
	for {
		record, err := cr.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		ch <- JiraProjectCategory{
			ID:          record[0],
			Checksum:    fromStringPointer(record[1]),
			Name:        fromStringPointer(record[2]),
			Description: fromStringPointer(record[3]),
			CustomerID:  record[4],
		}
	}
	return nil
}

// NewCSVJiraProjectCategoryReaderFile will read the file as a CSV and emit each record to the channel provided
func NewCSVJiraProjectCategoryReaderFile(fp string, ch chan<- JiraProjectCategory) error {
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
	return NewCSVJiraProjectCategoryReader(fc, ch)
}

// NewCSVJiraProjectCategoryReaderDir will read the jira_project_category.csv.gz file as a CSV and emit each record to the channel provided
func NewCSVJiraProjectCategoryReaderDir(dir string, ch chan<- JiraProjectCategory) error {
	return NewCSVJiraProjectCategoryReaderFile(filepath.Join(dir, "jira_project_category.csv.gz"), ch)
}

// JiraProjectCategoryCSVDeduper is a function callback which takes the existing value (a) and the new value (b)
// and the return value should be the one to use (or a new one, if applicable). return nil
// to skip processing of this record
type JiraProjectCategoryCSVDeduper func(a JiraProjectCategory, b JiraProjectCategory) *JiraProjectCategory

// JiraProjectCategoryCSVDedupeDisabled is set on whether the CSV writer should de-dupe values by key
var JiraProjectCategoryCSVDedupeDisabled bool

// NewJiraProjectCategoryCSVWriterSize creates a batch writer that will write each JiraProjectCategory into a CSV file
// this method will automatically de-duplicate entries using the primary key. if the checksum
// for a newer item with the same primary key doesn't match a previously sent item, the newer
// one will be used
func NewJiraProjectCategoryCSVWriterSize(w io.Writer, size int, dedupers ...JiraProjectCategoryCSVDeduper) (chan JiraProjectCategory, chan bool, error) {
	cw := csv.NewWriter(w)
	ch := make(chan JiraProjectCategory, size)
	done := make(chan bool)
	go func() {
		defer func() { done <- true }()
		dodedupe := !JiraProjectCategoryCSVDedupeDisabled
		var kv map[string]*JiraProjectCategory
		var deduper JiraProjectCategoryCSVDeduper
		if dedupers != nil && len(dedupers) > 0 {
			deduper = dedupers[0]
			dodedupe = true
		}
		if dodedupe {
			kv = make(map[string]*JiraProjectCategory)
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

// JiraProjectCategoryCSVDefaultSize is the default channel buffer size if not provided
var JiraProjectCategoryCSVDefaultSize = 100

// NewJiraProjectCategoryCSVWriter creates a batch writer that will write each JiraProjectCategory into a CSV file
func NewJiraProjectCategoryCSVWriter(w io.Writer, dedupers ...JiraProjectCategoryCSVDeduper) (chan JiraProjectCategory, chan bool, error) {
	return NewJiraProjectCategoryCSVWriterSize(w, JiraProjectCategoryCSVDefaultSize, dedupers...)
}

// NewJiraProjectCategoryCSVWriterDir creates a batch writer that will write each JiraProjectCategory into a CSV file named jira_project_category.csv.gz in dir
func NewJiraProjectCategoryCSVWriterDir(dir string, dedupers ...JiraProjectCategoryCSVDeduper) (chan JiraProjectCategory, chan bool, error) {
	return NewJiraProjectCategoryCSVWriterFile(filepath.Join(dir, "jira_project_category.csv.gz"), dedupers...)
}

// NewJiraProjectCategoryCSVWriterFile creates a batch writer that will write each JiraProjectCategory into a CSV file
func NewJiraProjectCategoryCSVWriterFile(fn string, dedupers ...JiraProjectCategoryCSVDeduper) (chan JiraProjectCategory, chan bool, error) {
	f, err := os.Create(fn)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening CSV file %s. %v", fn, err)
	}
	var fc io.WriteCloser = f
	if filepath.Ext(fn) == ".gz" {
		w, _ := gzip.NewWriterLevel(f, gzip.BestCompression)
		fc = w
	}
	ch, done, err := NewJiraProjectCategoryCSVWriter(fc, dedupers...)
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

type JiraProjectCategoryDBAction func(ctx context.Context, db *sql.DB, record JiraProjectCategory) error

// NewJiraProjectCategoryDBWriterSize creates a DB writer that will write each issue into the DB
func NewJiraProjectCategoryDBWriterSize(ctx context.Context, db *sql.DB, errors chan<- error, size int, actions ...JiraProjectCategoryDBAction) (chan JiraProjectCategory, chan bool, error) {
	ch := make(chan JiraProjectCategory, size)
	done := make(chan bool)
	var action JiraProjectCategoryDBAction
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

// NewJiraProjectCategoryDBWriter creates a DB writer that will write each issue into the DB
func NewJiraProjectCategoryDBWriter(ctx context.Context, db *sql.DB, errors chan<- error, actions ...JiraProjectCategoryDBAction) (chan JiraProjectCategory, chan bool, error) {
	return NewJiraProjectCategoryDBWriterSize(ctx, db, errors, 100, actions...)
}

// JiraProjectCategoryColumnID is the ID SQL column name for the JiraProjectCategory table
const JiraProjectCategoryColumnID = "id"

// JiraProjectCategoryEscapedColumnID is the escaped ID SQL column name for the JiraProjectCategory table
const JiraProjectCategoryEscapedColumnID = "`id`"

// JiraProjectCategoryColumnChecksum is the Checksum SQL column name for the JiraProjectCategory table
const JiraProjectCategoryColumnChecksum = "checksum"

// JiraProjectCategoryEscapedColumnChecksum is the escaped Checksum SQL column name for the JiraProjectCategory table
const JiraProjectCategoryEscapedColumnChecksum = "`checksum`"

// JiraProjectCategoryColumnName is the Name SQL column name for the JiraProjectCategory table
const JiraProjectCategoryColumnName = "name"

// JiraProjectCategoryEscapedColumnName is the escaped Name SQL column name for the JiraProjectCategory table
const JiraProjectCategoryEscapedColumnName = "`name`"

// JiraProjectCategoryColumnDescription is the Description SQL column name for the JiraProjectCategory table
const JiraProjectCategoryColumnDescription = "description"

// JiraProjectCategoryEscapedColumnDescription is the escaped Description SQL column name for the JiraProjectCategory table
const JiraProjectCategoryEscapedColumnDescription = "`description`"

// JiraProjectCategoryColumnCustomerID is the CustomerID SQL column name for the JiraProjectCategory table
const JiraProjectCategoryColumnCustomerID = "customer_id"

// JiraProjectCategoryEscapedColumnCustomerID is the escaped CustomerID SQL column name for the JiraProjectCategory table
const JiraProjectCategoryEscapedColumnCustomerID = "`customer_id`"

// GetID will return the JiraProjectCategory ID value
func (t *JiraProjectCategory) GetID() string {
	return t.ID
}

// SetID will set the JiraProjectCategory ID value
func (t *JiraProjectCategory) SetID(v string) {
	t.ID = v
}

// FindJiraProjectCategoryByID will find a JiraProjectCategory by ID
func FindJiraProjectCategoryByID(ctx context.Context, db *sql.DB, value string) (*JiraProjectCategory, error) {
	q := "SELECT `jira_project_category`.`id`,`jira_project_category`.`checksum`,`jira_project_category`.`name`,`jira_project_category`.`description`,`jira_project_category`.`customer_id` FROM `jira_project_category` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _Name sql.NullString
	var _Description sql.NullString
	var _CustomerID sql.NullString
	err := db.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_Name,
		&_Description,
		&_CustomerID,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &JiraProjectCategory{}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _Checksum.Valid {
		t.SetChecksum(_Checksum.String)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _Description.Valid {
		t.SetDescription(_Description.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	return t, nil
}

// FindJiraProjectCategoryByIDTx will find a JiraProjectCategory by ID using the provided transaction
func FindJiraProjectCategoryByIDTx(ctx context.Context, tx *sql.Tx, value string) (*JiraProjectCategory, error) {
	q := "SELECT `jira_project_category`.`id`,`jira_project_category`.`checksum`,`jira_project_category`.`name`,`jira_project_category`.`description`,`jira_project_category`.`customer_id` FROM `jira_project_category` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _Name sql.NullString
	var _Description sql.NullString
	var _CustomerID sql.NullString
	err := tx.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_Name,
		&_Description,
		&_CustomerID,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &JiraProjectCategory{}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _Checksum.Valid {
		t.SetChecksum(_Checksum.String)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _Description.Valid {
		t.SetDescription(_Description.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	return t, nil
}

// GetChecksum will return the JiraProjectCategory Checksum value
func (t *JiraProjectCategory) GetChecksum() string {
	if t.Checksum == nil {
		return ""
	}
	return *t.Checksum
}

// SetChecksum will set the JiraProjectCategory Checksum value
func (t *JiraProjectCategory) SetChecksum(v string) {
	t.Checksum = &v
}

// GetName will return the JiraProjectCategory Name value
func (t *JiraProjectCategory) GetName() string {
	if t.Name == nil {
		return ""
	}
	return *t.Name
}

// SetName will set the JiraProjectCategory Name value
func (t *JiraProjectCategory) SetName(v string) {
	t.Name = &v
}

// GetDescription will return the JiraProjectCategory Description value
func (t *JiraProjectCategory) GetDescription() string {
	if t.Description == nil {
		return ""
	}
	return *t.Description
}

// SetDescription will set the JiraProjectCategory Description value
func (t *JiraProjectCategory) SetDescription(v string) {
	t.Description = &v
}

// GetCustomerID will return the JiraProjectCategory CustomerID value
func (t *JiraProjectCategory) GetCustomerID() string {
	return t.CustomerID
}

// SetCustomerID will set the JiraProjectCategory CustomerID value
func (t *JiraProjectCategory) SetCustomerID(v string) {
	t.CustomerID = v
}

// FindJiraProjectCategoriesByCustomerID will find all JiraProjectCategorys by the CustomerID value
func FindJiraProjectCategoriesByCustomerID(ctx context.Context, db *sql.DB, value string) ([]*JiraProjectCategory, error) {
	q := "SELECT `jira_project_category`.`id`,`jira_project_category`.`checksum`,`jira_project_category`.`name`,`jira_project_category`.`description`,`jira_project_category`.`customer_id` FROM `jira_project_category` WHERE `customer_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraProjectCategory, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _Name sql.NullString
		var _Description sql.NullString
		var _CustomerID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_Name,
			&_Description,
			&_CustomerID,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraProjectCategory{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _Description.Valid {
			t.SetDescription(_Description.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindJiraProjectCategoriesByCustomerIDTx will find all JiraProjectCategorys by the CustomerID value using the provided transaction
func FindJiraProjectCategoriesByCustomerIDTx(ctx context.Context, tx *sql.Tx, value string) ([]*JiraProjectCategory, error) {
	q := "SELECT `jira_project_category`.`id`,`jira_project_category`.`checksum`,`jira_project_category`.`name`,`jira_project_category`.`description`,`jira_project_category`.`customer_id` FROM `jira_project_category` WHERE `customer_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraProjectCategory, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _Name sql.NullString
		var _Description sql.NullString
		var _CustomerID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_Name,
			&_Description,
			&_CustomerID,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraProjectCategory{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _Description.Valid {
			t.SetDescription(_Description.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

func (t *JiraProjectCategory) toTimestamp(value time.Time) *timestamp.Timestamp {
	ts, _ := ptypes.TimestampProto(value)
	return ts
}

// DBCreateJiraProjectCategoryTable will create the JiraProjectCategory table
func DBCreateJiraProjectCategoryTable(ctx context.Context, db *sql.DB) error {
	q := "CREATE TABLE `jira_project_category` (`id`VARCHAR(64) NOT NULL PRIMARY KEY,`checksum`CHAR(64),`name` TEXT,`description`TEXT,`customer_id`VARCHAR(64) NOT NULL,INDEX jira_project_category_customer_id_index (`customer_id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBCreateJiraProjectCategoryTableTx will create the JiraProjectCategory table using the provided transction
func DBCreateJiraProjectCategoryTableTx(ctx context.Context, tx *sql.Tx) error {
	q := "CREATE TABLE `jira_project_category` (`id`VARCHAR(64) NOT NULL PRIMARY KEY,`checksum`CHAR(64),`name` TEXT,`description`TEXT,`customer_id`VARCHAR(64) NOT NULL,INDEX jira_project_category_customer_id_index (`customer_id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// DBDropJiraProjectCategoryTable will drop the JiraProjectCategory table
func DBDropJiraProjectCategoryTable(ctx context.Context, db *sql.DB) error {
	q := "DROP TABLE IF EXISTS `jira_project_category`"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBDropJiraProjectCategoryTableTx will drop the JiraProjectCategory table using the provided transaction
func DBDropJiraProjectCategoryTableTx(ctx context.Context, tx *sql.Tx) error {
	q := "DROP TABLE IF EXISTS `jira_project_category`"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// CalculateChecksum will calculate a checksum of the SHA1 of all field values
func (t *JiraProjectCategory) CalculateChecksum() string {
	return orm.HashStrings(
		orm.ToString(t.ID),
		orm.ToString(t.Name),
		orm.ToString(t.Description),
		orm.ToString(t.CustomerID),
	)
}

// DBCreate will create a new JiraProjectCategory record in the database
func (t *JiraProjectCategory) DBCreate(ctx context.Context, db *sql.DB) (sql.Result, error) {
	q := "INSERT INTO `jira_project_category` (`jira_project_category`.`id`,`jira_project_category`.`checksum`,`jira_project_category`.`name`,`jira_project_category`.`description`,`jira_project_category`.`customer_id`) VALUES (?,?,?,?,?)"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.Description),
		orm.ToSQLString(t.CustomerID),
	)
}

// DBCreateTx will create a new JiraProjectCategory record in the database using the provided transaction
func (t *JiraProjectCategory) DBCreateTx(ctx context.Context, tx *sql.Tx) (sql.Result, error) {
	q := "INSERT INTO `jira_project_category` (`jira_project_category`.`id`,`jira_project_category`.`checksum`,`jira_project_category`.`name`,`jira_project_category`.`description`,`jira_project_category`.`customer_id`) VALUES (?,?,?,?,?)"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.Description),
		orm.ToSQLString(t.CustomerID),
	)
}

// DBCreateIgnoreDuplicate will upsert the JiraProjectCategory record in the database
func (t *JiraProjectCategory) DBCreateIgnoreDuplicate(ctx context.Context, db *sql.DB) (sql.Result, error) {
	q := "INSERT INTO `jira_project_category` (`jira_project_category`.`id`,`jira_project_category`.`checksum`,`jira_project_category`.`name`,`jira_project_category`.`description`,`jira_project_category`.`customer_id`) VALUES (?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.Description),
		orm.ToSQLString(t.CustomerID),
	)
}

// DBCreateIgnoreDuplicateTx will upsert the JiraProjectCategory record in the database using the provided transaction
func (t *JiraProjectCategory) DBCreateIgnoreDuplicateTx(ctx context.Context, tx *sql.Tx) (sql.Result, error) {
	q := "INSERT INTO `jira_project_category` (`jira_project_category`.`id`,`jira_project_category`.`checksum`,`jira_project_category`.`name`,`jira_project_category`.`description`,`jira_project_category`.`customer_id`) VALUES (?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.Description),
		orm.ToSQLString(t.CustomerID),
	)
}

// DeleteAllJiraProjectCategories deletes all JiraProjectCategory records in the database with optional filters
func DeleteAllJiraProjectCategories(ctx context.Context, db *sql.DB, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(JiraProjectCategoryTableName),
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

// DeleteAllJiraProjectCategoriesTx deletes all JiraProjectCategory records in the database with optional filters using the provided transaction
func DeleteAllJiraProjectCategoriesTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(JiraProjectCategoryTableName),
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

// DBDelete will delete this JiraProjectCategory record in the database
func (t *JiraProjectCategory) DBDelete(ctx context.Context, db *sql.DB) (bool, error) {
	q := "DELETE FROM `jira_project_category` WHERE `id` = ?"
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

// DBDeleteTx will delete this JiraProjectCategory record in the database using the provided transaction
func (t *JiraProjectCategory) DBDeleteTx(ctx context.Context, tx *sql.Tx) (bool, error) {
	q := "DELETE FROM `jira_project_category` WHERE `id` = ?"
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

// DBUpdate will update the JiraProjectCategory record in the database
func (t *JiraProjectCategory) DBUpdate(ctx context.Context, db *sql.DB) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `jira_project_category` SET `checksum`=?,`name`=?,`description`=?,`customer_id`=? WHERE `id`=?"
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.Description),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.ID),
	)
}

// DBUpdateTx will update the JiraProjectCategory record in the database using the provided transaction
func (t *JiraProjectCategory) DBUpdateTx(ctx context.Context, tx *sql.Tx) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `jira_project_category` SET `checksum`=?,`name`=?,`description`=?,`customer_id`=? WHERE `id`=?"
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.Description),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.ID),
	)
}

// DBUpsert will upsert the JiraProjectCategory record in the database
func (t *JiraProjectCategory) DBUpsert(ctx context.Context, db *sql.DB, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `jira_project_category` (`jira_project_category`.`id`,`jira_project_category`.`checksum`,`jira_project_category`.`name`,`jira_project_category`.`description`,`jira_project_category`.`customer_id`) VALUES (?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `jira_project_category` (`jira_project_category`.`id`,`jira_project_category`.`checksum`,`jira_project_category`.`name`,`jira_project_category`.`description`,`jira_project_category`.`customer_id`) VALUES (?,?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`name`=VALUES(`name`),`description`=VALUES(`description`),`customer_id`=VALUES(`customer_id`)"
	}
	r, err := db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.Description),
		orm.ToSQLString(t.CustomerID),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBUpsertTx will upsert the JiraProjectCategory record in the database using the provided transaction
func (t *JiraProjectCategory) DBUpsertTx(ctx context.Context, tx *sql.Tx, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `jira_project_category` (`jira_project_category`.`id`,`jira_project_category`.`checksum`,`jira_project_category`.`name`,`jira_project_category`.`description`,`jira_project_category`.`customer_id`) VALUES (?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `jira_project_category` (`jira_project_category`.`id`,`jira_project_category`.`checksum`,`jira_project_category`.`name`,`jira_project_category`.`description`,`jira_project_category`.`customer_id`) VALUES (?,?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`name`=VALUES(`name`),`description`=VALUES(`description`),`customer_id`=VALUES(`customer_id`)"
	}
	r, err := tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.Description),
		orm.ToSQLString(t.CustomerID),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBFindOne will find a JiraProjectCategory record in the database with the primary key
func (t *JiraProjectCategory) DBFindOne(ctx context.Context, db *sql.DB, value string) (bool, error) {
	q := "SELECT `jira_project_category`.`id`,`jira_project_category`.`checksum`,`jira_project_category`.`name`,`jira_project_category`.`description`,`jira_project_category`.`customer_id` FROM `jira_project_category` WHERE `id` = ? LIMIT 1"
	row := db.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _Name sql.NullString
	var _Description sql.NullString
	var _CustomerID sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_Name,
		&_Description,
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
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _Description.Valid {
		t.SetDescription(_Description.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	return true, nil
}

// DBFindOneTx will find a JiraProjectCategory record in the database with the primary key using the provided transaction
func (t *JiraProjectCategory) DBFindOneTx(ctx context.Context, tx *sql.Tx, value string) (bool, error) {
	q := "SELECT `jira_project_category`.`id`,`jira_project_category`.`checksum`,`jira_project_category`.`name`,`jira_project_category`.`description`,`jira_project_category`.`customer_id` FROM `jira_project_category` WHERE `id` = ? LIMIT 1"
	row := tx.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _Name sql.NullString
	var _Description sql.NullString
	var _CustomerID sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_Name,
		&_Description,
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
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _Description.Valid {
		t.SetDescription(_Description.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	return true, nil
}

// FindJiraProjectCategories will find a JiraProjectCategory record in the database with the provided parameters
func FindJiraProjectCategories(ctx context.Context, db *sql.DB, _params ...interface{}) ([]*JiraProjectCategory, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("name"),
		orm.Column("description"),
		orm.Column("customer_id"),
		orm.Table(JiraProjectCategoryTableName),
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
	results := make([]*JiraProjectCategory, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _Name sql.NullString
		var _Description sql.NullString
		var _CustomerID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_Name,
			&_Description,
			&_CustomerID,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraProjectCategory{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _Description.Valid {
			t.SetDescription(_Description.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindJiraProjectCategoriesTx will find a JiraProjectCategory record in the database with the provided parameters using the provided transaction
func FindJiraProjectCategoriesTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) ([]*JiraProjectCategory, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("name"),
		orm.Column("description"),
		orm.Column("customer_id"),
		orm.Table(JiraProjectCategoryTableName),
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
	results := make([]*JiraProjectCategory, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _Name sql.NullString
		var _Description sql.NullString
		var _CustomerID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_Name,
			&_Description,
			&_CustomerID,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraProjectCategory{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _Description.Valid {
			t.SetDescription(_Description.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// DBFind will find a JiraProjectCategory record in the database with the provided parameters
func (t *JiraProjectCategory) DBFind(ctx context.Context, db *sql.DB, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("name"),
		orm.Column("description"),
		orm.Column("customer_id"),
		orm.Table(JiraProjectCategoryTableName),
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
	var _Name sql.NullString
	var _Description sql.NullString
	var _CustomerID sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_Name,
		&_Description,
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
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _Description.Valid {
		t.SetDescription(_Description.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	return true, nil
}

// DBFindTx will find a JiraProjectCategory record in the database with the provided parameters using the provided transaction
func (t *JiraProjectCategory) DBFindTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("name"),
		orm.Column("description"),
		orm.Column("customer_id"),
		orm.Table(JiraProjectCategoryTableName),
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
	var _Name sql.NullString
	var _Description sql.NullString
	var _CustomerID sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_Name,
		&_Description,
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
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _Description.Valid {
		t.SetDescription(_Description.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	return true, nil
}

// CountJiraProjectCategories will find the count of JiraProjectCategory records in the database
func CountJiraProjectCategories(ctx context.Context, db *sql.DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(JiraProjectCategoryTableName),
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

// CountJiraProjectCategoriesTx will find the count of JiraProjectCategory records in the database using the provided transaction
func CountJiraProjectCategoriesTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(JiraProjectCategoryTableName),
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

// DBCount will find the count of JiraProjectCategory records in the database
func (t *JiraProjectCategory) DBCount(ctx context.Context, db *sql.DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(JiraProjectCategoryTableName),
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

// DBCountTx will find the count of JiraProjectCategory records in the database using the provided transaction
func (t *JiraProjectCategory) DBCountTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(JiraProjectCategoryTableName),
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

// DBExists will return true if the JiraProjectCategory record exists in the database
func (t *JiraProjectCategory) DBExists(ctx context.Context, db *sql.DB) (bool, error) {
	q := "SELECT `id` FROM `jira_project_category` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := db.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// DBExistsTx will return true if the JiraProjectCategory record exists in the database using the provided transaction
func (t *JiraProjectCategory) DBExistsTx(ctx context.Context, tx *sql.Tx) (bool, error) {
	q := "SELECT `id` FROM `jira_project_category` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := tx.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// PrimaryKeyColumn returns the column name for the primary key
func (t *JiraProjectCategory) PrimaryKeyColumn() string {
	return JiraProjectCategoryColumnID
}

// PrimaryKeyColumnType returns the primary key column Go type as a string
func (t *JiraProjectCategory) PrimaryKeyColumnType() string {
	return "string"
}

// PrimaryKey returns the primary key column value
func (t *JiraProjectCategory) PrimaryKey() interface{} {
	return t.ID
}
