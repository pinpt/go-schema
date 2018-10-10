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

var _ Model = (*JiraIssueLabel)(nil)
var _ CSVWriter = (*JiraIssueLabel)(nil)
var _ JSONWriter = (*JiraIssueLabel)(nil)
var _ Checksum = (*JiraIssueLabel)(nil)

// JiraIssueLabelTableName is the name of the table in SQL
const JiraIssueLabelTableName = "jira_issue_label"

var JiraIssueLabelColumns = []string{
	"id",
	"checksum",
	"name",
	"customer_id",
}

// JiraIssueLabel table
type JiraIssueLabel struct {
	Checksum   *string `json:"checksum,omitempty"`
	CustomerID string  `json:"customer_id"`
	ID         string  `json:"id"`
	Name       string  `json:"name"`
}

// TableName returns the SQL table name for JiraIssueLabel and satifies the Model interface
func (t *JiraIssueLabel) TableName() string {
	return JiraIssueLabelTableName
}

// ToCSV will serialize the JiraIssueLabel instance to a CSV compatible array of strings
func (t *JiraIssueLabel) ToCSV() []string {
	return []string{
		t.ID,
		t.CalculateChecksum(),
		t.Name,
		t.CustomerID,
	}
}

// WriteCSV will serialize the JiraIssueLabel instance to the writer as CSV and satisfies the CSVWriter interface
func (t *JiraIssueLabel) WriteCSV(w *csv.Writer) error {
	return w.Write(t.ToCSV())
}

// WriteJSON will serialize the JiraIssueLabel instance to the writer as JSON and satisfies the JSONWriter interface
func (t *JiraIssueLabel) WriteJSON(w io.Writer, indent ...bool) error {
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

// NewJiraIssueLabelReader creates a JSON reader which can read in JiraIssueLabel objects serialized as JSON either as an array, single object or json new lines
// and writes each JiraIssueLabel to the channel provided
func NewJiraIssueLabelReader(r io.Reader, ch chan<- JiraIssueLabel) error {
	return orm.Deserialize(r, func(buf json.RawMessage) error {
		dec := json.NewDecoder(bytes.NewBuffer(buf))
		e := JiraIssueLabel{}
		if err := dec.Decode(&e); err != nil {
			return err
		}
		ch <- e
		return nil
	})
}

// NewCSVJiraIssueLabelReaderDir will read the reader as CSV and emit each record to the channel provided
func NewCSVJiraIssueLabelReader(r io.Reader, ch chan<- JiraIssueLabel) error {
	cr := csv.NewReader(r)
	for {
		record, err := cr.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		ch <- JiraIssueLabel{
			ID:         record[0],
			Checksum:   fromStringPointer(record[1]),
			Name:       record[2],
			CustomerID: record[3],
		}
	}
	return nil
}

// NewCSVJiraIssueLabelReaderFile will read the file as a CSV and emit each record to the channel provided
func NewCSVJiraIssueLabelReaderFile(fp string, ch chan<- JiraIssueLabel) error {
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
	return NewCSVJiraIssueLabelReader(fc, ch)
}

// NewCSVJiraIssueLabelReaderDir will read the jira_issue_label.csv.gz file as a CSV and emit each record to the channel provided
func NewCSVJiraIssueLabelReaderDir(dir string, ch chan<- JiraIssueLabel) error {
	return NewCSVJiraIssueLabelReaderFile(filepath.Join(dir, "jira_issue_label.csv.gz"), ch)
}

// JiraIssueLabelCSVDeduper is a function callback which takes the existing value (a) and the new value (b)
// and the return value should be the one to use (or a new one, if applicable). return nil
// to skip processing of this record
type JiraIssueLabelCSVDeduper func(a JiraIssueLabel, b JiraIssueLabel) *JiraIssueLabel

// JiraIssueLabelCSVDedupeDisabled is set on whether the CSV writer should de-dupe values by key
var JiraIssueLabelCSVDedupeDisabled bool

// NewJiraIssueLabelCSVWriterSize creates a batch writer that will write each JiraIssueLabel into a CSV file
// this method will automatically de-duplicate entries using the primary key. if the checksum
// for a newer item with the same primary key doesn't match a previously sent item, the newer
// one will be used
func NewJiraIssueLabelCSVWriterSize(w io.Writer, size int, dedupers ...JiraIssueLabelCSVDeduper) (chan JiraIssueLabel, chan bool, error) {
	cw := csv.NewWriter(w)
	ch := make(chan JiraIssueLabel, size)
	done := make(chan bool)
	go func() {
		defer func() { done <- true }()
		dodedupe := !JiraIssueLabelCSVDedupeDisabled
		var kv map[string]*JiraIssueLabel
		var deduper JiraIssueLabelCSVDeduper
		if dedupers != nil && len(dedupers) > 0 {
			deduper = dedupers[0]
			dodedupe = true
		}
		if dodedupe {
			kv = make(map[string]*JiraIssueLabel)
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

// JiraIssueLabelCSVDefaultSize is the default channel buffer size if not provided
var JiraIssueLabelCSVDefaultSize = 100

// NewJiraIssueLabelCSVWriter creates a batch writer that will write each JiraIssueLabel into a CSV file
func NewJiraIssueLabelCSVWriter(w io.Writer, dedupers ...JiraIssueLabelCSVDeduper) (chan JiraIssueLabel, chan bool, error) {
	return NewJiraIssueLabelCSVWriterSize(w, JiraIssueLabelCSVDefaultSize, dedupers...)
}

// NewJiraIssueLabelCSVWriterDir creates a batch writer that will write each JiraIssueLabel into a CSV file named jira_issue_label.csv.gz in dir
func NewJiraIssueLabelCSVWriterDir(dir string, dedupers ...JiraIssueLabelCSVDeduper) (chan JiraIssueLabel, chan bool, error) {
	return NewJiraIssueLabelCSVWriterFile(filepath.Join(dir, "jira_issue_label.csv.gz"), dedupers...)
}

// NewJiraIssueLabelCSVWriterFile creates a batch writer that will write each JiraIssueLabel into a CSV file
func NewJiraIssueLabelCSVWriterFile(fn string, dedupers ...JiraIssueLabelCSVDeduper) (chan JiraIssueLabel, chan bool, error) {
	f, err := os.Create(fn)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening CSV file %s. %v", fn, err)
	}
	var fc io.WriteCloser = f
	if filepath.Ext(fn) == ".gz" {
		w, _ := gzip.NewWriterLevel(f, gzip.BestCompression)
		fc = w
	}
	ch, done, err := NewJiraIssueLabelCSVWriter(fc, dedupers...)
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

type JiraIssueLabelDBAction func(ctx context.Context, db DB, record JiraIssueLabel) error

// NewJiraIssueLabelDBWriterSize creates a DB writer that will write each issue into the DB
func NewJiraIssueLabelDBWriterSize(ctx context.Context, db DB, errors chan<- error, size int, actions ...JiraIssueLabelDBAction) (chan JiraIssueLabel, chan bool, error) {
	ch := make(chan JiraIssueLabel, size)
	done := make(chan bool)
	var action JiraIssueLabelDBAction
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

// NewJiraIssueLabelDBWriter creates a DB writer that will write each issue into the DB
func NewJiraIssueLabelDBWriter(ctx context.Context, db DB, errors chan<- error, actions ...JiraIssueLabelDBAction) (chan JiraIssueLabel, chan bool, error) {
	return NewJiraIssueLabelDBWriterSize(ctx, db, errors, 100, actions...)
}

// JiraIssueLabelColumnID is the ID SQL column name for the JiraIssueLabel table
const JiraIssueLabelColumnID = "id"

// JiraIssueLabelEscapedColumnID is the escaped ID SQL column name for the JiraIssueLabel table
const JiraIssueLabelEscapedColumnID = "`id`"

// JiraIssueLabelColumnChecksum is the Checksum SQL column name for the JiraIssueLabel table
const JiraIssueLabelColumnChecksum = "checksum"

// JiraIssueLabelEscapedColumnChecksum is the escaped Checksum SQL column name for the JiraIssueLabel table
const JiraIssueLabelEscapedColumnChecksum = "`checksum`"

// JiraIssueLabelColumnName is the Name SQL column name for the JiraIssueLabel table
const JiraIssueLabelColumnName = "name"

// JiraIssueLabelEscapedColumnName is the escaped Name SQL column name for the JiraIssueLabel table
const JiraIssueLabelEscapedColumnName = "`name`"

// JiraIssueLabelColumnCustomerID is the CustomerID SQL column name for the JiraIssueLabel table
const JiraIssueLabelColumnCustomerID = "customer_id"

// JiraIssueLabelEscapedColumnCustomerID is the escaped CustomerID SQL column name for the JiraIssueLabel table
const JiraIssueLabelEscapedColumnCustomerID = "`customer_id`"

// GetID will return the JiraIssueLabel ID value
func (t *JiraIssueLabel) GetID() string {
	return t.ID
}

// SetID will set the JiraIssueLabel ID value
func (t *JiraIssueLabel) SetID(v string) {
	t.ID = v
}

// FindJiraIssueLabelByID will find a JiraIssueLabel by ID
func FindJiraIssueLabelByID(ctx context.Context, db DB, value string) (*JiraIssueLabel, error) {
	q := "SELECT * FROM `jira_issue_label` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _Name sql.NullString
	var _CustomerID sql.NullString
	err := db.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_Name,
		&_CustomerID,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &JiraIssueLabel{}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _Checksum.Valid {
		t.SetChecksum(_Checksum.String)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	return t, nil
}

// FindJiraIssueLabelByIDTx will find a JiraIssueLabel by ID using the provided transaction
func FindJiraIssueLabelByIDTx(ctx context.Context, tx Tx, value string) (*JiraIssueLabel, error) {
	q := "SELECT * FROM `jira_issue_label` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _Name sql.NullString
	var _CustomerID sql.NullString
	err := tx.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_Name,
		&_CustomerID,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &JiraIssueLabel{}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _Checksum.Valid {
		t.SetChecksum(_Checksum.String)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	return t, nil
}

// GetChecksum will return the JiraIssueLabel Checksum value
func (t *JiraIssueLabel) GetChecksum() string {
	if t.Checksum == nil {
		return ""
	}
	return *t.Checksum
}

// SetChecksum will set the JiraIssueLabel Checksum value
func (t *JiraIssueLabel) SetChecksum(v string) {
	t.Checksum = &v
}

// GetName will return the JiraIssueLabel Name value
func (t *JiraIssueLabel) GetName() string {
	return t.Name
}

// SetName will set the JiraIssueLabel Name value
func (t *JiraIssueLabel) SetName(v string) {
	t.Name = v
}

// FindJiraIssueLabelsByName will find all JiraIssueLabels by the Name value
func FindJiraIssueLabelsByName(ctx context.Context, db DB, value string) ([]*JiraIssueLabel, error) {
	q := "SELECT * FROM `jira_issue_label` WHERE `name` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraIssueLabel, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _Name sql.NullString
		var _CustomerID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_Name,
			&_CustomerID,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraIssueLabel{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindJiraIssueLabelsByNameTx will find all JiraIssueLabels by the Name value using the provided transaction
func FindJiraIssueLabelsByNameTx(ctx context.Context, tx Tx, value string) ([]*JiraIssueLabel, error) {
	q := "SELECT * FROM `jira_issue_label` WHERE `name` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraIssueLabel, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _Name sql.NullString
		var _CustomerID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_Name,
			&_CustomerID,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraIssueLabel{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetCustomerID will return the JiraIssueLabel CustomerID value
func (t *JiraIssueLabel) GetCustomerID() string {
	return t.CustomerID
}

// SetCustomerID will set the JiraIssueLabel CustomerID value
func (t *JiraIssueLabel) SetCustomerID(v string) {
	t.CustomerID = v
}

// FindJiraIssueLabelsByCustomerID will find all JiraIssueLabels by the CustomerID value
func FindJiraIssueLabelsByCustomerID(ctx context.Context, db DB, value string) ([]*JiraIssueLabel, error) {
	q := "SELECT * FROM `jira_issue_label` WHERE `customer_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraIssueLabel, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _Name sql.NullString
		var _CustomerID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_Name,
			&_CustomerID,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraIssueLabel{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindJiraIssueLabelsByCustomerIDTx will find all JiraIssueLabels by the CustomerID value using the provided transaction
func FindJiraIssueLabelsByCustomerIDTx(ctx context.Context, tx Tx, value string) ([]*JiraIssueLabel, error) {
	q := "SELECT * FROM `jira_issue_label` WHERE `customer_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraIssueLabel, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _Name sql.NullString
		var _CustomerID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_Name,
			&_CustomerID,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraIssueLabel{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

func (t *JiraIssueLabel) toTimestamp(value time.Time) *timestamp.Timestamp {
	ts, _ := ptypes.TimestampProto(value)
	return ts
}

// DBCreateJiraIssueLabelTable will create the JiraIssueLabel table
func DBCreateJiraIssueLabelTable(ctx context.Context, db DB) error {
	q := "CREATE TABLE `jira_issue_label` (`id` VARCHAR(64) NOT NULL PRIMARY KEY,`checksum` CHAR(64),`name`VARCHAR(255) NOT NULL,`customer_id` VARCHAR(64) NOT NULL,INDEX jira_issue_label_name_index (`name`),INDEX jira_issue_label_customer_id_index (`customer_id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBCreateJiraIssueLabelTableTx will create the JiraIssueLabel table using the provided transction
func DBCreateJiraIssueLabelTableTx(ctx context.Context, tx Tx) error {
	q := "CREATE TABLE `jira_issue_label` (`id` VARCHAR(64) NOT NULL PRIMARY KEY,`checksum` CHAR(64),`name`VARCHAR(255) NOT NULL,`customer_id` VARCHAR(64) NOT NULL,INDEX jira_issue_label_name_index (`name`),INDEX jira_issue_label_customer_id_index (`customer_id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// DBDropJiraIssueLabelTable will drop the JiraIssueLabel table
func DBDropJiraIssueLabelTable(ctx context.Context, db DB) error {
	q := "DROP TABLE IF EXISTS `jira_issue_label`"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBDropJiraIssueLabelTableTx will drop the JiraIssueLabel table using the provided transaction
func DBDropJiraIssueLabelTableTx(ctx context.Context, tx Tx) error {
	q := "DROP TABLE IF EXISTS `jira_issue_label`"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// CalculateChecksum will calculate a checksum of the SHA1 of all field values
func (t *JiraIssueLabel) CalculateChecksum() string {
	return orm.HashStrings(
		orm.ToString(t.ID),
		orm.ToString(t.Name),
		orm.ToString(t.CustomerID),
	)
}

// DBCreate will create a new JiraIssueLabel record in the database
func (t *JiraIssueLabel) DBCreate(ctx context.Context, db DB) (sql.Result, error) {
	q := "INSERT INTO `jira_issue_label` (`jira_issue_label`.`id`,`jira_issue_label`.`checksum`,`jira_issue_label`.`name`,`jira_issue_label`.`customer_id`) VALUES (?,?,?,?)"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.CustomerID),
	)
}

// DBCreateTx will create a new JiraIssueLabel record in the database using the provided transaction
func (t *JiraIssueLabel) DBCreateTx(ctx context.Context, tx Tx) (sql.Result, error) {
	q := "INSERT INTO `jira_issue_label` (`jira_issue_label`.`id`,`jira_issue_label`.`checksum`,`jira_issue_label`.`name`,`jira_issue_label`.`customer_id`) VALUES (?,?,?,?)"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.CustomerID),
	)
}

// DBCreateIgnoreDuplicate will upsert the JiraIssueLabel record in the database
func (t *JiraIssueLabel) DBCreateIgnoreDuplicate(ctx context.Context, db DB) (sql.Result, error) {
	q := "INSERT INTO `jira_issue_label` (`jira_issue_label`.`id`,`jira_issue_label`.`checksum`,`jira_issue_label`.`name`,`jira_issue_label`.`customer_id`) VALUES (?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.CustomerID),
	)
}

// DBCreateIgnoreDuplicateTx will upsert the JiraIssueLabel record in the database using the provided transaction
func (t *JiraIssueLabel) DBCreateIgnoreDuplicateTx(ctx context.Context, tx Tx) (sql.Result, error) {
	q := "INSERT INTO `jira_issue_label` (`jira_issue_label`.`id`,`jira_issue_label`.`checksum`,`jira_issue_label`.`name`,`jira_issue_label`.`customer_id`) VALUES (?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.CustomerID),
	)
}

// DeleteAllJiraIssueLabels deletes all JiraIssueLabel records in the database with optional filters
func DeleteAllJiraIssueLabels(ctx context.Context, db DB, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(JiraIssueLabelTableName),
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

// DeleteAllJiraIssueLabelsTx deletes all JiraIssueLabel records in the database with optional filters using the provided transaction
func DeleteAllJiraIssueLabelsTx(ctx context.Context, tx Tx, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(JiraIssueLabelTableName),
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

// DBDelete will delete this JiraIssueLabel record in the database
func (t *JiraIssueLabel) DBDelete(ctx context.Context, db DB) (bool, error) {
	q := "DELETE FROM `jira_issue_label` WHERE `id` = ?"
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

// DBDeleteTx will delete this JiraIssueLabel record in the database using the provided transaction
func (t *JiraIssueLabel) DBDeleteTx(ctx context.Context, tx Tx) (bool, error) {
	q := "DELETE FROM `jira_issue_label` WHERE `id` = ?"
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

// DBUpdate will update the JiraIssueLabel record in the database
func (t *JiraIssueLabel) DBUpdate(ctx context.Context, db DB) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `jira_issue_label` SET `checksum`=?,`name`=?,`customer_id`=? WHERE `id`=?"
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.ID),
	)
}

// DBUpdateTx will update the JiraIssueLabel record in the database using the provided transaction
func (t *JiraIssueLabel) DBUpdateTx(ctx context.Context, tx Tx) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `jira_issue_label` SET `checksum`=?,`name`=?,`customer_id`=? WHERE `id`=?"
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.ID),
	)
}

// DBUpsert will upsert the JiraIssueLabel record in the database
func (t *JiraIssueLabel) DBUpsert(ctx context.Context, db DB, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `jira_issue_label` (`jira_issue_label`.`id`,`jira_issue_label`.`checksum`,`jira_issue_label`.`name`,`jira_issue_label`.`customer_id`) VALUES (?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `jira_issue_label` (`jira_issue_label`.`id`,`jira_issue_label`.`checksum`,`jira_issue_label`.`name`,`jira_issue_label`.`customer_id`) VALUES (?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`name`=VALUES(`name`),`customer_id`=VALUES(`customer_id`)"
	}
	r, err := db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.CustomerID),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBUpsertTx will upsert the JiraIssueLabel record in the database using the provided transaction
func (t *JiraIssueLabel) DBUpsertTx(ctx context.Context, tx Tx, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `jira_issue_label` (`jira_issue_label`.`id`,`jira_issue_label`.`checksum`,`jira_issue_label`.`name`,`jira_issue_label`.`customer_id`) VALUES (?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `jira_issue_label` (`jira_issue_label`.`id`,`jira_issue_label`.`checksum`,`jira_issue_label`.`name`,`jira_issue_label`.`customer_id`) VALUES (?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`name`=VALUES(`name`),`customer_id`=VALUES(`customer_id`)"
	}
	r, err := tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.CustomerID),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBFindOne will find a JiraIssueLabel record in the database with the primary key
func (t *JiraIssueLabel) DBFindOne(ctx context.Context, db DB, value string) (bool, error) {
	q := "SELECT * FROM `jira_issue_label` WHERE `id` = ? LIMIT 1"
	row := db.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _Name sql.NullString
	var _CustomerID sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_Name,
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
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	return true, nil
}

// DBFindOneTx will find a JiraIssueLabel record in the database with the primary key using the provided transaction
func (t *JiraIssueLabel) DBFindOneTx(ctx context.Context, tx Tx, value string) (bool, error) {
	q := "SELECT * FROM `jira_issue_label` WHERE `id` = ? LIMIT 1"
	row := tx.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _Name sql.NullString
	var _CustomerID sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_Name,
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
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	return true, nil
}

// FindJiraIssueLabels will find a JiraIssueLabel record in the database with the provided parameters
func FindJiraIssueLabels(ctx context.Context, db DB, _params ...interface{}) ([]*JiraIssueLabel, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("name"),
		orm.Column("customer_id"),
		orm.Table(JiraIssueLabelTableName),
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
	results := make([]*JiraIssueLabel, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _Name sql.NullString
		var _CustomerID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_Name,
			&_CustomerID,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraIssueLabel{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindJiraIssueLabelsTx will find a JiraIssueLabel record in the database with the provided parameters using the provided transaction
func FindJiraIssueLabelsTx(ctx context.Context, tx Tx, _params ...interface{}) ([]*JiraIssueLabel, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("name"),
		orm.Column("customer_id"),
		orm.Table(JiraIssueLabelTableName),
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
	results := make([]*JiraIssueLabel, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _Name sql.NullString
		var _CustomerID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_Name,
			&_CustomerID,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraIssueLabel{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// DBFind will find a JiraIssueLabel record in the database with the provided parameters
func (t *JiraIssueLabel) DBFind(ctx context.Context, db DB, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("name"),
		orm.Column("customer_id"),
		orm.Table(JiraIssueLabelTableName),
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
	var _CustomerID sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_Name,
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
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	return true, nil
}

// DBFindTx will find a JiraIssueLabel record in the database with the provided parameters using the provided transaction
func (t *JiraIssueLabel) DBFindTx(ctx context.Context, tx Tx, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("name"),
		orm.Column("customer_id"),
		orm.Table(JiraIssueLabelTableName),
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
	var _CustomerID sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_Name,
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
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	return true, nil
}

// CountJiraIssueLabels will find the count of JiraIssueLabel records in the database
func CountJiraIssueLabels(ctx context.Context, db DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(JiraIssueLabelTableName),
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

// CountJiraIssueLabelsTx will find the count of JiraIssueLabel records in the database using the provided transaction
func CountJiraIssueLabelsTx(ctx context.Context, tx Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(JiraIssueLabelTableName),
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

// DBCount will find the count of JiraIssueLabel records in the database
func (t *JiraIssueLabel) DBCount(ctx context.Context, db DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(JiraIssueLabelTableName),
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

// DBCountTx will find the count of JiraIssueLabel records in the database using the provided transaction
func (t *JiraIssueLabel) DBCountTx(ctx context.Context, tx Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(JiraIssueLabelTableName),
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

// DBExists will return true if the JiraIssueLabel record exists in the database
func (t *JiraIssueLabel) DBExists(ctx context.Context, db DB) (bool, error) {
	q := "SELECT * FROM `jira_issue_label` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := db.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// DBExistsTx will return true if the JiraIssueLabel record exists in the database using the provided transaction
func (t *JiraIssueLabel) DBExistsTx(ctx context.Context, tx Tx) (bool, error) {
	q := "SELECT * FROM `jira_issue_label` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := tx.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// PrimaryKeyColumn returns the column name for the primary key
func (t *JiraIssueLabel) PrimaryKeyColumn() string {
	return JiraIssueLabelColumnID
}

// PrimaryKeyColumnType returns the primary key column Go type as a string
func (t *JiraIssueLabel) PrimaryKeyColumnType() string {
	return "string"
}

// PrimaryKey returns the primary key column value
func (t *JiraIssueLabel) PrimaryKey() interface{} {
	return t.ID
}
