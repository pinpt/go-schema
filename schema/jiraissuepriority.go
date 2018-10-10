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

var _ Model = (*JiraIssuePriority)(nil)
var _ CSVWriter = (*JiraIssuePriority)(nil)
var _ JSONWriter = (*JiraIssuePriority)(nil)
var _ Checksum = (*JiraIssuePriority)(nil)

// JiraIssuePriorityTableName is the name of the table in SQL
const JiraIssuePriorityTableName = "jira_issue_priority"

var JiraIssuePriorityColumns = []string{
	"id",
	"checksum",
	"issue_priority_id",
	"name",
	"icon_url",
	"customer_id",
}

// JiraIssuePriority table
type JiraIssuePriority struct {
	Checksum        *string `json:"checksum,omitempty"`
	CustomerID      string  `json:"customer_id"`
	IconURL         string  `json:"icon_url"`
	ID              string  `json:"id"`
	IssuePriorityID string  `json:"issue_priority_id"`
	Name            string  `json:"name"`
}

// TableName returns the SQL table name for JiraIssuePriority and satifies the Model interface
func (t *JiraIssuePriority) TableName() string {
	return JiraIssuePriorityTableName
}

// ToCSV will serialize the JiraIssuePriority instance to a CSV compatible array of strings
func (t *JiraIssuePriority) ToCSV() []string {
	return []string{
		t.ID,
		t.CalculateChecksum(),
		t.IssuePriorityID,
		t.Name,
		t.IconURL,
		t.CustomerID,
	}
}

// WriteCSV will serialize the JiraIssuePriority instance to the writer as CSV and satisfies the CSVWriter interface
func (t *JiraIssuePriority) WriteCSV(w *csv.Writer) error {
	return w.Write(t.ToCSV())
}

// WriteJSON will serialize the JiraIssuePriority instance to the writer as JSON and satisfies the JSONWriter interface
func (t *JiraIssuePriority) WriteJSON(w io.Writer, indent ...bool) error {
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

// NewJiraIssuePriorityReader creates a JSON reader which can read in JiraIssuePriority objects serialized as JSON either as an array, single object or json new lines
// and writes each JiraIssuePriority to the channel provided
func NewJiraIssuePriorityReader(r io.Reader, ch chan<- JiraIssuePriority) error {
	return orm.Deserialize(r, func(buf json.RawMessage) error {
		dec := json.NewDecoder(bytes.NewBuffer(buf))
		e := JiraIssuePriority{}
		if err := dec.Decode(&e); err != nil {
			return err
		}
		ch <- e
		return nil
	})
}

// NewCSVJiraIssuePriorityReaderDir will read the reader as CSV and emit each record to the channel provided
func NewCSVJiraIssuePriorityReader(r io.Reader, ch chan<- JiraIssuePriority) error {
	cr := csv.NewReader(r)
	for {
		record, err := cr.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		ch <- JiraIssuePriority{
			ID:              record[0],
			Checksum:        fromStringPointer(record[1]),
			IssuePriorityID: record[2],
			Name:            record[3],
			IconURL:         record[4],
			CustomerID:      record[5],
		}
	}
	return nil
}

// NewCSVJiraIssuePriorityReaderFile will read the file as a CSV and emit each record to the channel provided
func NewCSVJiraIssuePriorityReaderFile(fp string, ch chan<- JiraIssuePriority) error {
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
	return NewCSVJiraIssuePriorityReader(fc, ch)
}

// NewCSVJiraIssuePriorityReaderDir will read the jira_issue_priority.csv.gz file as a CSV and emit each record to the channel provided
func NewCSVJiraIssuePriorityReaderDir(dir string, ch chan<- JiraIssuePriority) error {
	return NewCSVJiraIssuePriorityReaderFile(filepath.Join(dir, "jira_issue_priority.csv.gz"), ch)
}

// JiraIssuePriorityCSVDeduper is a function callback which takes the existing value (a) and the new value (b)
// and the return value should be the one to use (or a new one, if applicable). return nil
// to skip processing of this record
type JiraIssuePriorityCSVDeduper func(a JiraIssuePriority, b JiraIssuePriority) *JiraIssuePriority

// JiraIssuePriorityCSVDedupeDisabled is set on whether the CSV writer should de-dupe values by key
var JiraIssuePriorityCSVDedupeDisabled bool

// NewJiraIssuePriorityCSVWriterSize creates a batch writer that will write each JiraIssuePriority into a CSV file
// this method will automatically de-duplicate entries using the primary key. if the checksum
// for a newer item with the same primary key doesn't match a previously sent item, the newer
// one will be used
func NewJiraIssuePriorityCSVWriterSize(w io.Writer, size int, dedupers ...JiraIssuePriorityCSVDeduper) (chan JiraIssuePriority, chan bool, error) {
	cw := csv.NewWriter(w)
	ch := make(chan JiraIssuePriority, size)
	done := make(chan bool)
	go func() {
		defer func() { done <- true }()
		dodedupe := !JiraIssuePriorityCSVDedupeDisabled
		var kv map[string]*JiraIssuePriority
		var deduper JiraIssuePriorityCSVDeduper
		if dedupers != nil && len(dedupers) > 0 {
			deduper = dedupers[0]
			dodedupe = true
		}
		if dodedupe {
			kv = make(map[string]*JiraIssuePriority)
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

// JiraIssuePriorityCSVDefaultSize is the default channel buffer size if not provided
var JiraIssuePriorityCSVDefaultSize = 100

// NewJiraIssuePriorityCSVWriter creates a batch writer that will write each JiraIssuePriority into a CSV file
func NewJiraIssuePriorityCSVWriter(w io.Writer, dedupers ...JiraIssuePriorityCSVDeduper) (chan JiraIssuePriority, chan bool, error) {
	return NewJiraIssuePriorityCSVWriterSize(w, JiraIssuePriorityCSVDefaultSize, dedupers...)
}

// NewJiraIssuePriorityCSVWriterDir creates a batch writer that will write each JiraIssuePriority into a CSV file named jira_issue_priority.csv.gz in dir
func NewJiraIssuePriorityCSVWriterDir(dir string, dedupers ...JiraIssuePriorityCSVDeduper) (chan JiraIssuePriority, chan bool, error) {
	return NewJiraIssuePriorityCSVWriterFile(filepath.Join(dir, "jira_issue_priority.csv.gz"), dedupers...)
}

// NewJiraIssuePriorityCSVWriterFile creates a batch writer that will write each JiraIssuePriority into a CSV file
func NewJiraIssuePriorityCSVWriterFile(fn string, dedupers ...JiraIssuePriorityCSVDeduper) (chan JiraIssuePriority, chan bool, error) {
	f, err := os.Create(fn)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening CSV file %s. %v", fn, err)
	}
	var fc io.WriteCloser = f
	if filepath.Ext(fn) == ".gz" {
		w, _ := gzip.NewWriterLevel(f, gzip.BestCompression)
		fc = w
	}
	ch, done, err := NewJiraIssuePriorityCSVWriter(fc, dedupers...)
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

type JiraIssuePriorityDBAction func(ctx context.Context, db DB, record JiraIssuePriority) error

// NewJiraIssuePriorityDBWriterSize creates a DB writer that will write each issue into the DB
func NewJiraIssuePriorityDBWriterSize(ctx context.Context, db DB, errors chan<- error, size int, actions ...JiraIssuePriorityDBAction) (chan JiraIssuePriority, chan bool, error) {
	ch := make(chan JiraIssuePriority, size)
	done := make(chan bool)
	var action JiraIssuePriorityDBAction
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

// NewJiraIssuePriorityDBWriter creates a DB writer that will write each issue into the DB
func NewJiraIssuePriorityDBWriter(ctx context.Context, db DB, errors chan<- error, actions ...JiraIssuePriorityDBAction) (chan JiraIssuePriority, chan bool, error) {
	return NewJiraIssuePriorityDBWriterSize(ctx, db, errors, 100, actions...)
}

// JiraIssuePriorityColumnID is the ID SQL column name for the JiraIssuePriority table
const JiraIssuePriorityColumnID = "id"

// JiraIssuePriorityEscapedColumnID is the escaped ID SQL column name for the JiraIssuePriority table
const JiraIssuePriorityEscapedColumnID = "`id`"

// JiraIssuePriorityColumnChecksum is the Checksum SQL column name for the JiraIssuePriority table
const JiraIssuePriorityColumnChecksum = "checksum"

// JiraIssuePriorityEscapedColumnChecksum is the escaped Checksum SQL column name for the JiraIssuePriority table
const JiraIssuePriorityEscapedColumnChecksum = "`checksum`"

// JiraIssuePriorityColumnIssuePriorityID is the IssuePriorityID SQL column name for the JiraIssuePriority table
const JiraIssuePriorityColumnIssuePriorityID = "issue_priority_id"

// JiraIssuePriorityEscapedColumnIssuePriorityID is the escaped IssuePriorityID SQL column name for the JiraIssuePriority table
const JiraIssuePriorityEscapedColumnIssuePriorityID = "`issue_priority_id`"

// JiraIssuePriorityColumnName is the Name SQL column name for the JiraIssuePriority table
const JiraIssuePriorityColumnName = "name"

// JiraIssuePriorityEscapedColumnName is the escaped Name SQL column name for the JiraIssuePriority table
const JiraIssuePriorityEscapedColumnName = "`name`"

// JiraIssuePriorityColumnIconURL is the IconURL SQL column name for the JiraIssuePriority table
const JiraIssuePriorityColumnIconURL = "icon_url"

// JiraIssuePriorityEscapedColumnIconURL is the escaped IconURL SQL column name for the JiraIssuePriority table
const JiraIssuePriorityEscapedColumnIconURL = "`icon_url`"

// JiraIssuePriorityColumnCustomerID is the CustomerID SQL column name for the JiraIssuePriority table
const JiraIssuePriorityColumnCustomerID = "customer_id"

// JiraIssuePriorityEscapedColumnCustomerID is the escaped CustomerID SQL column name for the JiraIssuePriority table
const JiraIssuePriorityEscapedColumnCustomerID = "`customer_id`"

// GetID will return the JiraIssuePriority ID value
func (t *JiraIssuePriority) GetID() string {
	return t.ID
}

// SetID will set the JiraIssuePriority ID value
func (t *JiraIssuePriority) SetID(v string) {
	t.ID = v
}

// FindJiraIssuePriorityByID will find a JiraIssuePriority by ID
func FindJiraIssuePriorityByID(ctx context.Context, db DB, value string) (*JiraIssuePriority, error) {
	q := "SELECT * FROM `jira_issue_priority` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _IssuePriorityID sql.NullString
	var _Name sql.NullString
	var _IconURL sql.NullString
	var _CustomerID sql.NullString
	err := db.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_IssuePriorityID,
		&_Name,
		&_IconURL,
		&_CustomerID,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &JiraIssuePriority{}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _Checksum.Valid {
		t.SetChecksum(_Checksum.String)
	}
	if _IssuePriorityID.Valid {
		t.SetIssuePriorityID(_IssuePriorityID.String)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _IconURL.Valid {
		t.SetIconURL(_IconURL.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	return t, nil
}

// FindJiraIssuePriorityByIDTx will find a JiraIssuePriority by ID using the provided transaction
func FindJiraIssuePriorityByIDTx(ctx context.Context, tx Tx, value string) (*JiraIssuePriority, error) {
	q := "SELECT * FROM `jira_issue_priority` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _IssuePriorityID sql.NullString
	var _Name sql.NullString
	var _IconURL sql.NullString
	var _CustomerID sql.NullString
	err := tx.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_IssuePriorityID,
		&_Name,
		&_IconURL,
		&_CustomerID,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &JiraIssuePriority{}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _Checksum.Valid {
		t.SetChecksum(_Checksum.String)
	}
	if _IssuePriorityID.Valid {
		t.SetIssuePriorityID(_IssuePriorityID.String)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _IconURL.Valid {
		t.SetIconURL(_IconURL.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	return t, nil
}

// GetChecksum will return the JiraIssuePriority Checksum value
func (t *JiraIssuePriority) GetChecksum() string {
	if t.Checksum == nil {
		return ""
	}
	return *t.Checksum
}

// SetChecksum will set the JiraIssuePriority Checksum value
func (t *JiraIssuePriority) SetChecksum(v string) {
	t.Checksum = &v
}

// GetIssuePriorityID will return the JiraIssuePriority IssuePriorityID value
func (t *JiraIssuePriority) GetIssuePriorityID() string {
	return t.IssuePriorityID
}

// SetIssuePriorityID will set the JiraIssuePriority IssuePriorityID value
func (t *JiraIssuePriority) SetIssuePriorityID(v string) {
	t.IssuePriorityID = v
}

// FindJiraIssuePrioritiesByIssuePriorityID will find all JiraIssuePrioritys by the IssuePriorityID value
func FindJiraIssuePrioritiesByIssuePriorityID(ctx context.Context, db DB, value string) ([]*JiraIssuePriority, error) {
	q := "SELECT * FROM `jira_issue_priority` WHERE `issue_priority_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraIssuePriority, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _IssuePriorityID sql.NullString
		var _Name sql.NullString
		var _IconURL sql.NullString
		var _CustomerID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_IssuePriorityID,
			&_Name,
			&_IconURL,
			&_CustomerID,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraIssuePriority{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _IssuePriorityID.Valid {
			t.SetIssuePriorityID(_IssuePriorityID.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _IconURL.Valid {
			t.SetIconURL(_IconURL.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindJiraIssuePrioritiesByIssuePriorityIDTx will find all JiraIssuePrioritys by the IssuePriorityID value using the provided transaction
func FindJiraIssuePrioritiesByIssuePriorityIDTx(ctx context.Context, tx Tx, value string) ([]*JiraIssuePriority, error) {
	q := "SELECT * FROM `jira_issue_priority` WHERE `issue_priority_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraIssuePriority, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _IssuePriorityID sql.NullString
		var _Name sql.NullString
		var _IconURL sql.NullString
		var _CustomerID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_IssuePriorityID,
			&_Name,
			&_IconURL,
			&_CustomerID,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraIssuePriority{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _IssuePriorityID.Valid {
			t.SetIssuePriorityID(_IssuePriorityID.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _IconURL.Valid {
			t.SetIconURL(_IconURL.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetName will return the JiraIssuePriority Name value
func (t *JiraIssuePriority) GetName() string {
	return t.Name
}

// SetName will set the JiraIssuePriority Name value
func (t *JiraIssuePriority) SetName(v string) {
	t.Name = v
}

// GetIconURL will return the JiraIssuePriority IconURL value
func (t *JiraIssuePriority) GetIconURL() string {
	return t.IconURL
}

// SetIconURL will set the JiraIssuePriority IconURL value
func (t *JiraIssuePriority) SetIconURL(v string) {
	t.IconURL = v
}

// GetCustomerID will return the JiraIssuePriority CustomerID value
func (t *JiraIssuePriority) GetCustomerID() string {
	return t.CustomerID
}

// SetCustomerID will set the JiraIssuePriority CustomerID value
func (t *JiraIssuePriority) SetCustomerID(v string) {
	t.CustomerID = v
}

// FindJiraIssuePrioritiesByCustomerID will find all JiraIssuePrioritys by the CustomerID value
func FindJiraIssuePrioritiesByCustomerID(ctx context.Context, db DB, value string) ([]*JiraIssuePriority, error) {
	q := "SELECT * FROM `jira_issue_priority` WHERE `customer_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraIssuePriority, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _IssuePriorityID sql.NullString
		var _Name sql.NullString
		var _IconURL sql.NullString
		var _CustomerID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_IssuePriorityID,
			&_Name,
			&_IconURL,
			&_CustomerID,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraIssuePriority{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _IssuePriorityID.Valid {
			t.SetIssuePriorityID(_IssuePriorityID.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _IconURL.Valid {
			t.SetIconURL(_IconURL.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindJiraIssuePrioritiesByCustomerIDTx will find all JiraIssuePrioritys by the CustomerID value using the provided transaction
func FindJiraIssuePrioritiesByCustomerIDTx(ctx context.Context, tx Tx, value string) ([]*JiraIssuePriority, error) {
	q := "SELECT * FROM `jira_issue_priority` WHERE `customer_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraIssuePriority, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _IssuePriorityID sql.NullString
		var _Name sql.NullString
		var _IconURL sql.NullString
		var _CustomerID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_IssuePriorityID,
			&_Name,
			&_IconURL,
			&_CustomerID,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraIssuePriority{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _IssuePriorityID.Valid {
			t.SetIssuePriorityID(_IssuePriorityID.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _IconURL.Valid {
			t.SetIconURL(_IconURL.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

func (t *JiraIssuePriority) toTimestamp(value time.Time) *timestamp.Timestamp {
	ts, _ := ptypes.TimestampProto(value)
	return ts
}

// DBCreateJiraIssuePriorityTable will create the JiraIssuePriority table
func DBCreateJiraIssuePriorityTable(ctx context.Context, db DB) error {
	q := "CREATE TABLE `jira_issue_priority` (`id` VARCHAR(64) NOT NULL PRIMARY KEY,`checksum` CHAR(64),`issue_priority_id` VARCHAR(255) NOT NULL,`name` TEXT NOT NULL,`icon_url` TEXT NOT NULL,`customer_id` VARCHAR(64) NOT NULL,INDEX jira_issue_priority_issue_priority_id_index (`issue_priority_id`),INDEX jira_issue_priority_customer_id_index (`customer_id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBCreateJiraIssuePriorityTableTx will create the JiraIssuePriority table using the provided transction
func DBCreateJiraIssuePriorityTableTx(ctx context.Context, tx Tx) error {
	q := "CREATE TABLE `jira_issue_priority` (`id` VARCHAR(64) NOT NULL PRIMARY KEY,`checksum` CHAR(64),`issue_priority_id` VARCHAR(255) NOT NULL,`name` TEXT NOT NULL,`icon_url` TEXT NOT NULL,`customer_id` VARCHAR(64) NOT NULL,INDEX jira_issue_priority_issue_priority_id_index (`issue_priority_id`),INDEX jira_issue_priority_customer_id_index (`customer_id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// DBDropJiraIssuePriorityTable will drop the JiraIssuePriority table
func DBDropJiraIssuePriorityTable(ctx context.Context, db DB) error {
	q := "DROP TABLE IF EXISTS `jira_issue_priority`"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBDropJiraIssuePriorityTableTx will drop the JiraIssuePriority table using the provided transaction
func DBDropJiraIssuePriorityTableTx(ctx context.Context, tx Tx) error {
	q := "DROP TABLE IF EXISTS `jira_issue_priority`"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// CalculateChecksum will calculate a checksum of the SHA1 of all field values
func (t *JiraIssuePriority) CalculateChecksum() string {
	return orm.HashStrings(
		orm.ToString(t.ID),
		orm.ToString(t.IssuePriorityID),
		orm.ToString(t.Name),
		orm.ToString(t.IconURL),
		orm.ToString(t.CustomerID),
	)
}

// DBCreate will create a new JiraIssuePriority record in the database
func (t *JiraIssuePriority) DBCreate(ctx context.Context, db DB) (sql.Result, error) {
	q := "INSERT INTO `jira_issue_priority` (`jira_issue_priority`.`id`,`jira_issue_priority`.`checksum`,`jira_issue_priority`.`issue_priority_id`,`jira_issue_priority`.`name`,`jira_issue_priority`.`icon_url`,`jira_issue_priority`.`customer_id`) VALUES (?,?,?,?,?,?)"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.IssuePriorityID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.IconURL),
		orm.ToSQLString(t.CustomerID),
	)
}

// DBCreateTx will create a new JiraIssuePriority record in the database using the provided transaction
func (t *JiraIssuePriority) DBCreateTx(ctx context.Context, tx Tx) (sql.Result, error) {
	q := "INSERT INTO `jira_issue_priority` (`jira_issue_priority`.`id`,`jira_issue_priority`.`checksum`,`jira_issue_priority`.`issue_priority_id`,`jira_issue_priority`.`name`,`jira_issue_priority`.`icon_url`,`jira_issue_priority`.`customer_id`) VALUES (?,?,?,?,?,?)"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.IssuePriorityID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.IconURL),
		orm.ToSQLString(t.CustomerID),
	)
}

// DBCreateIgnoreDuplicate will upsert the JiraIssuePriority record in the database
func (t *JiraIssuePriority) DBCreateIgnoreDuplicate(ctx context.Context, db DB) (sql.Result, error) {
	q := "INSERT INTO `jira_issue_priority` (`jira_issue_priority`.`id`,`jira_issue_priority`.`checksum`,`jira_issue_priority`.`issue_priority_id`,`jira_issue_priority`.`name`,`jira_issue_priority`.`icon_url`,`jira_issue_priority`.`customer_id`) VALUES (?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.IssuePriorityID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.IconURL),
		orm.ToSQLString(t.CustomerID),
	)
}

// DBCreateIgnoreDuplicateTx will upsert the JiraIssuePriority record in the database using the provided transaction
func (t *JiraIssuePriority) DBCreateIgnoreDuplicateTx(ctx context.Context, tx Tx) (sql.Result, error) {
	q := "INSERT INTO `jira_issue_priority` (`jira_issue_priority`.`id`,`jira_issue_priority`.`checksum`,`jira_issue_priority`.`issue_priority_id`,`jira_issue_priority`.`name`,`jira_issue_priority`.`icon_url`,`jira_issue_priority`.`customer_id`) VALUES (?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.IssuePriorityID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.IconURL),
		orm.ToSQLString(t.CustomerID),
	)
}

// DeleteAllJiraIssuePriorities deletes all JiraIssuePriority records in the database with optional filters
func DeleteAllJiraIssuePriorities(ctx context.Context, db DB, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(JiraIssuePriorityTableName),
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

// DeleteAllJiraIssuePrioritiesTx deletes all JiraIssuePriority records in the database with optional filters using the provided transaction
func DeleteAllJiraIssuePrioritiesTx(ctx context.Context, tx Tx, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(JiraIssuePriorityTableName),
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

// DBDelete will delete this JiraIssuePriority record in the database
func (t *JiraIssuePriority) DBDelete(ctx context.Context, db DB) (bool, error) {
	q := "DELETE FROM `jira_issue_priority` WHERE `id` = ?"
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

// DBDeleteTx will delete this JiraIssuePriority record in the database using the provided transaction
func (t *JiraIssuePriority) DBDeleteTx(ctx context.Context, tx Tx) (bool, error) {
	q := "DELETE FROM `jira_issue_priority` WHERE `id` = ?"
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

// DBUpdate will update the JiraIssuePriority record in the database
func (t *JiraIssuePriority) DBUpdate(ctx context.Context, db DB) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `jira_issue_priority` SET `checksum`=?,`issue_priority_id`=?,`name`=?,`icon_url`=?,`customer_id`=? WHERE `id`=?"
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.IssuePriorityID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.IconURL),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.ID),
	)
}

// DBUpdateTx will update the JiraIssuePriority record in the database using the provided transaction
func (t *JiraIssuePriority) DBUpdateTx(ctx context.Context, tx Tx) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `jira_issue_priority` SET `checksum`=?,`issue_priority_id`=?,`name`=?,`icon_url`=?,`customer_id`=? WHERE `id`=?"
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.IssuePriorityID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.IconURL),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.ID),
	)
}

// DBUpsert will upsert the JiraIssuePriority record in the database
func (t *JiraIssuePriority) DBUpsert(ctx context.Context, db DB, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `jira_issue_priority` (`jira_issue_priority`.`id`,`jira_issue_priority`.`checksum`,`jira_issue_priority`.`issue_priority_id`,`jira_issue_priority`.`name`,`jira_issue_priority`.`icon_url`,`jira_issue_priority`.`customer_id`) VALUES (?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `jira_issue_priority` (`jira_issue_priority`.`id`,`jira_issue_priority`.`checksum`,`jira_issue_priority`.`issue_priority_id`,`jira_issue_priority`.`name`,`jira_issue_priority`.`icon_url`,`jira_issue_priority`.`customer_id`) VALUES (?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`issue_priority_id`=VALUES(`issue_priority_id`),`name`=VALUES(`name`),`icon_url`=VALUES(`icon_url`),`customer_id`=VALUES(`customer_id`)"
	}
	r, err := db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.IssuePriorityID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.IconURL),
		orm.ToSQLString(t.CustomerID),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBUpsertTx will upsert the JiraIssuePriority record in the database using the provided transaction
func (t *JiraIssuePriority) DBUpsertTx(ctx context.Context, tx Tx, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `jira_issue_priority` (`jira_issue_priority`.`id`,`jira_issue_priority`.`checksum`,`jira_issue_priority`.`issue_priority_id`,`jira_issue_priority`.`name`,`jira_issue_priority`.`icon_url`,`jira_issue_priority`.`customer_id`) VALUES (?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `jira_issue_priority` (`jira_issue_priority`.`id`,`jira_issue_priority`.`checksum`,`jira_issue_priority`.`issue_priority_id`,`jira_issue_priority`.`name`,`jira_issue_priority`.`icon_url`,`jira_issue_priority`.`customer_id`) VALUES (?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`issue_priority_id`=VALUES(`issue_priority_id`),`name`=VALUES(`name`),`icon_url`=VALUES(`icon_url`),`customer_id`=VALUES(`customer_id`)"
	}
	r, err := tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.IssuePriorityID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.IconURL),
		orm.ToSQLString(t.CustomerID),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBFindOne will find a JiraIssuePriority record in the database with the primary key
func (t *JiraIssuePriority) DBFindOne(ctx context.Context, db DB, value string) (bool, error) {
	q := "SELECT * FROM `jira_issue_priority` WHERE `id` = ? LIMIT 1"
	row := db.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _IssuePriorityID sql.NullString
	var _Name sql.NullString
	var _IconURL sql.NullString
	var _CustomerID sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_IssuePriorityID,
		&_Name,
		&_IconURL,
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
	if _IssuePriorityID.Valid {
		t.SetIssuePriorityID(_IssuePriorityID.String)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _IconURL.Valid {
		t.SetIconURL(_IconURL.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	return true, nil
}

// DBFindOneTx will find a JiraIssuePriority record in the database with the primary key using the provided transaction
func (t *JiraIssuePriority) DBFindOneTx(ctx context.Context, tx Tx, value string) (bool, error) {
	q := "SELECT * FROM `jira_issue_priority` WHERE `id` = ? LIMIT 1"
	row := tx.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _IssuePriorityID sql.NullString
	var _Name sql.NullString
	var _IconURL sql.NullString
	var _CustomerID sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_IssuePriorityID,
		&_Name,
		&_IconURL,
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
	if _IssuePriorityID.Valid {
		t.SetIssuePriorityID(_IssuePriorityID.String)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _IconURL.Valid {
		t.SetIconURL(_IconURL.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	return true, nil
}

// FindJiraIssuePriorities will find a JiraIssuePriority record in the database with the provided parameters
func FindJiraIssuePriorities(ctx context.Context, db DB, _params ...interface{}) ([]*JiraIssuePriority, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("issue_priority_id"),
		orm.Column("name"),
		orm.Column("icon_url"),
		orm.Column("customer_id"),
		orm.Table(JiraIssuePriorityTableName),
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
	results := make([]*JiraIssuePriority, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _IssuePriorityID sql.NullString
		var _Name sql.NullString
		var _IconURL sql.NullString
		var _CustomerID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_IssuePriorityID,
			&_Name,
			&_IconURL,
			&_CustomerID,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraIssuePriority{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _IssuePriorityID.Valid {
			t.SetIssuePriorityID(_IssuePriorityID.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _IconURL.Valid {
			t.SetIconURL(_IconURL.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindJiraIssuePrioritiesTx will find a JiraIssuePriority record in the database with the provided parameters using the provided transaction
func FindJiraIssuePrioritiesTx(ctx context.Context, tx Tx, _params ...interface{}) ([]*JiraIssuePriority, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("issue_priority_id"),
		orm.Column("name"),
		orm.Column("icon_url"),
		orm.Column("customer_id"),
		orm.Table(JiraIssuePriorityTableName),
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
	results := make([]*JiraIssuePriority, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _IssuePriorityID sql.NullString
		var _Name sql.NullString
		var _IconURL sql.NullString
		var _CustomerID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_IssuePriorityID,
			&_Name,
			&_IconURL,
			&_CustomerID,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraIssuePriority{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _IssuePriorityID.Valid {
			t.SetIssuePriorityID(_IssuePriorityID.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _IconURL.Valid {
			t.SetIconURL(_IconURL.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// DBFind will find a JiraIssuePriority record in the database with the provided parameters
func (t *JiraIssuePriority) DBFind(ctx context.Context, db DB, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("issue_priority_id"),
		orm.Column("name"),
		orm.Column("icon_url"),
		orm.Column("customer_id"),
		orm.Table(JiraIssuePriorityTableName),
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
	var _IssuePriorityID sql.NullString
	var _Name sql.NullString
	var _IconURL sql.NullString
	var _CustomerID sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_IssuePriorityID,
		&_Name,
		&_IconURL,
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
	if _IssuePriorityID.Valid {
		t.SetIssuePriorityID(_IssuePriorityID.String)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _IconURL.Valid {
		t.SetIconURL(_IconURL.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	return true, nil
}

// DBFindTx will find a JiraIssuePriority record in the database with the provided parameters using the provided transaction
func (t *JiraIssuePriority) DBFindTx(ctx context.Context, tx Tx, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("issue_priority_id"),
		orm.Column("name"),
		orm.Column("icon_url"),
		orm.Column("customer_id"),
		orm.Table(JiraIssuePriorityTableName),
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
	var _IssuePriorityID sql.NullString
	var _Name sql.NullString
	var _IconURL sql.NullString
	var _CustomerID sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_IssuePriorityID,
		&_Name,
		&_IconURL,
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
	if _IssuePriorityID.Valid {
		t.SetIssuePriorityID(_IssuePriorityID.String)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _IconURL.Valid {
		t.SetIconURL(_IconURL.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	return true, nil
}

// CountJiraIssuePriorities will find the count of JiraIssuePriority records in the database
func CountJiraIssuePriorities(ctx context.Context, db DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(JiraIssuePriorityTableName),
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

// CountJiraIssuePrioritiesTx will find the count of JiraIssuePriority records in the database using the provided transaction
func CountJiraIssuePrioritiesTx(ctx context.Context, tx Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(JiraIssuePriorityTableName),
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

// DBCount will find the count of JiraIssuePriority records in the database
func (t *JiraIssuePriority) DBCount(ctx context.Context, db DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(JiraIssuePriorityTableName),
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

// DBCountTx will find the count of JiraIssuePriority records in the database using the provided transaction
func (t *JiraIssuePriority) DBCountTx(ctx context.Context, tx Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(JiraIssuePriorityTableName),
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

// DBExists will return true if the JiraIssuePriority record exists in the database
func (t *JiraIssuePriority) DBExists(ctx context.Context, db DB) (bool, error) {
	q := "SELECT * FROM `jira_issue_priority` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := db.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// DBExistsTx will return true if the JiraIssuePriority record exists in the database using the provided transaction
func (t *JiraIssuePriority) DBExistsTx(ctx context.Context, tx Tx) (bool, error) {
	q := "SELECT * FROM `jira_issue_priority` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := tx.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// PrimaryKeyColumn returns the column name for the primary key
func (t *JiraIssuePriority) PrimaryKeyColumn() string {
	return JiraIssuePriorityColumnID
}

// PrimaryKeyColumnType returns the primary key column Go type as a string
func (t *JiraIssuePriority) PrimaryKeyColumnType() string {
	return "string"
}

// PrimaryKey returns the primary key column value
func (t *JiraIssuePriority) PrimaryKey() interface{} {
	return t.ID
}
