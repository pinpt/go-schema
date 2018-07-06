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

var _ Model = (*JiraProjectComponent)(nil)
var _ CSVWriter = (*JiraProjectComponent)(nil)
var _ JSONWriter = (*JiraProjectComponent)(nil)
var _ Checksum = (*JiraProjectComponent)(nil)

// JiraProjectComponentTableName is the name of the table in SQL
const JiraProjectComponentTableName = "jira_project_component"

var JiraProjectComponentColumns = []string{
	"id",
	"checksum",
	"component_id",
	"project_id",
	"name",
	"customer_id",
}

// JiraProjectComponent table
type JiraProjectComponent struct {
	Checksum    *string `json:"checksum,omitempty"`
	ComponentID string  `json:"component_id"`
	CustomerID  string  `json:"customer_id"`
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	ProjectID   string  `json:"project_id"`
}

// TableName returns the SQL table name for JiraProjectComponent and satifies the Model interface
func (t *JiraProjectComponent) TableName() string {
	return JiraProjectComponentTableName
}

// ToCSV will serialize the JiraProjectComponent instance to a CSV compatible array of strings
func (t *JiraProjectComponent) ToCSV() []string {
	return []string{
		t.ID,
		t.CalculateChecksum(),
		t.ComponentID,
		t.ProjectID,
		t.Name,
		t.CustomerID,
	}
}

// WriteCSV will serialize the JiraProjectComponent instance to the writer as CSV and satisfies the CSVWriter interface
func (t *JiraProjectComponent) WriteCSV(w *csv.Writer) error {
	return w.Write(t.ToCSV())
}

// WriteJSON will serialize the JiraProjectComponent instance to the writer as JSON and satisfies the JSONWriter interface
func (t *JiraProjectComponent) WriteJSON(w io.Writer, indent ...bool) error {
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

// NewJiraProjectComponentReader creates a JSON reader which can read in JiraProjectComponent objects serialized as JSON either as an array, single object or json new lines
// and writes each JiraProjectComponent to the channel provided
func NewJiraProjectComponentReader(r io.Reader, ch chan<- JiraProjectComponent) error {
	return orm.Deserialize(r, func(buf json.RawMessage) error {
		dec := json.NewDecoder(bytes.NewBuffer(buf))
		e := JiraProjectComponent{}
		if err := dec.Decode(&e); err != nil {
			return err
		}
		ch <- e
		return nil
	})
}

// NewCSVJiraProjectComponentReaderDir will read the reader as CSV and emit each record to the channel provided
func NewCSVJiraProjectComponentReader(r io.Reader, ch chan<- JiraProjectComponent) error {
	cr := csv.NewReader(r)
	for {
		record, err := cr.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		ch <- JiraProjectComponent{
			ID:          record[0],
			Checksum:    fromStringPointer(record[1]),
			ComponentID: record[2],
			ProjectID:   record[3],
			Name:        record[4],
			CustomerID:  record[5],
		}
	}
	return nil
}

// NewCSVJiraProjectComponentReaderFile will read the file as a CSV and emit each record to the channel provided
func NewCSVJiraProjectComponentReaderFile(fp string, ch chan<- JiraProjectComponent) error {
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
	return NewCSVJiraProjectComponentReader(fc, ch)
}

// NewCSVJiraProjectComponentReaderDir will read the jira_project_component.csv.gz file as a CSV and emit each record to the channel provided
func NewCSVJiraProjectComponentReaderDir(dir string, ch chan<- JiraProjectComponent) error {
	return NewCSVJiraProjectComponentReaderFile(filepath.Join(dir, "jira_project_component.csv.gz"), ch)
}

// JiraProjectComponentCSVDeduper is a function callback which takes the existing value (a) and the new value (b)
// and the return value should be the one to use (or a new one, if applicable). return nil
// to skip processing of this record
type JiraProjectComponentCSVDeduper func(a JiraProjectComponent, b JiraProjectComponent) *JiraProjectComponent

// JiraProjectComponentCSVDedupeDisabled is set on whether the CSV writer should de-dupe values by key
var JiraProjectComponentCSVDedupeDisabled bool

// NewJiraProjectComponentCSVWriterSize creates a batch writer that will write each JiraProjectComponent into a CSV file
// this method will automatically de-duplicate entries using the primary key. if the checksum
// for a newer item with the same primary key doesn't match a previously sent item, the newer
// one will be used
func NewJiraProjectComponentCSVWriterSize(w io.Writer, size int, dedupers ...JiraProjectComponentCSVDeduper) (chan JiraProjectComponent, chan bool, error) {
	cw := csv.NewWriter(w)
	ch := make(chan JiraProjectComponent, size)
	done := make(chan bool)
	go func() {
		defer func() { done <- true }()
		dodedupe := !JiraProjectComponentCSVDedupeDisabled
		var kv map[string]*JiraProjectComponent
		var deduper JiraProjectComponentCSVDeduper
		if dedupers != nil && len(dedupers) > 0 {
			deduper = dedupers[0]
			dodedupe = true
		}
		if dodedupe {
			kv = make(map[string]*JiraProjectComponent)
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

// JiraProjectComponentCSVDefaultSize is the default channel buffer size if not provided
var JiraProjectComponentCSVDefaultSize = 100

// NewJiraProjectComponentCSVWriter creates a batch writer that will write each JiraProjectComponent into a CSV file
func NewJiraProjectComponentCSVWriter(w io.Writer, dedupers ...JiraProjectComponentCSVDeduper) (chan JiraProjectComponent, chan bool, error) {
	return NewJiraProjectComponentCSVWriterSize(w, JiraProjectComponentCSVDefaultSize, dedupers...)
}

// NewJiraProjectComponentCSVWriterDir creates a batch writer that will write each JiraProjectComponent into a CSV file named jira_project_component.csv.gz in dir
func NewJiraProjectComponentCSVWriterDir(dir string, dedupers ...JiraProjectComponentCSVDeduper) (chan JiraProjectComponent, chan bool, error) {
	return NewJiraProjectComponentCSVWriterFile(filepath.Join(dir, "jira_project_component.csv.gz"), dedupers...)
}

// NewJiraProjectComponentCSVWriterFile creates a batch writer that will write each JiraProjectComponent into a CSV file
func NewJiraProjectComponentCSVWriterFile(fn string, dedupers ...JiraProjectComponentCSVDeduper) (chan JiraProjectComponent, chan bool, error) {
	f, err := os.Create(fn)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening CSV file %s. %v", fn, err)
	}
	var fc io.WriteCloser = f
	if filepath.Ext(fn) == ".gz" {
		w, _ := gzip.NewWriterLevel(f, gzip.BestCompression)
		fc = w
	}
	ch, done, err := NewJiraProjectComponentCSVWriter(fc, dedupers...)
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

type JiraProjectComponentDBAction func(ctx context.Context, db *sql.DB, record JiraProjectComponent) error

// NewJiraProjectComponentDBWriterSize creates a DB writer that will write each issue into the DB
func NewJiraProjectComponentDBWriterSize(ctx context.Context, db *sql.DB, errors chan<- error, size int, actions ...JiraProjectComponentDBAction) (chan JiraProjectComponent, chan bool, error) {
	ch := make(chan JiraProjectComponent, size)
	done := make(chan bool)
	var action JiraProjectComponentDBAction
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

// NewJiraProjectComponentDBWriter creates a DB writer that will write each issue into the DB
func NewJiraProjectComponentDBWriter(ctx context.Context, db *sql.DB, errors chan<- error, actions ...JiraProjectComponentDBAction) (chan JiraProjectComponent, chan bool, error) {
	return NewJiraProjectComponentDBWriterSize(ctx, db, errors, 100, actions...)
}

// JiraProjectComponentColumnID is the ID SQL column name for the JiraProjectComponent table
const JiraProjectComponentColumnID = "id"

// JiraProjectComponentEscapedColumnID is the escaped ID SQL column name for the JiraProjectComponent table
const JiraProjectComponentEscapedColumnID = "`id`"

// JiraProjectComponentColumnChecksum is the Checksum SQL column name for the JiraProjectComponent table
const JiraProjectComponentColumnChecksum = "checksum"

// JiraProjectComponentEscapedColumnChecksum is the escaped Checksum SQL column name for the JiraProjectComponent table
const JiraProjectComponentEscapedColumnChecksum = "`checksum`"

// JiraProjectComponentColumnComponentID is the ComponentID SQL column name for the JiraProjectComponent table
const JiraProjectComponentColumnComponentID = "component_id"

// JiraProjectComponentEscapedColumnComponentID is the escaped ComponentID SQL column name for the JiraProjectComponent table
const JiraProjectComponentEscapedColumnComponentID = "`component_id`"

// JiraProjectComponentColumnProjectID is the ProjectID SQL column name for the JiraProjectComponent table
const JiraProjectComponentColumnProjectID = "project_id"

// JiraProjectComponentEscapedColumnProjectID is the escaped ProjectID SQL column name for the JiraProjectComponent table
const JiraProjectComponentEscapedColumnProjectID = "`project_id`"

// JiraProjectComponentColumnName is the Name SQL column name for the JiraProjectComponent table
const JiraProjectComponentColumnName = "name"

// JiraProjectComponentEscapedColumnName is the escaped Name SQL column name for the JiraProjectComponent table
const JiraProjectComponentEscapedColumnName = "`name`"

// JiraProjectComponentColumnCustomerID is the CustomerID SQL column name for the JiraProjectComponent table
const JiraProjectComponentColumnCustomerID = "customer_id"

// JiraProjectComponentEscapedColumnCustomerID is the escaped CustomerID SQL column name for the JiraProjectComponent table
const JiraProjectComponentEscapedColumnCustomerID = "`customer_id`"

// GetID will return the JiraProjectComponent ID value
func (t *JiraProjectComponent) GetID() string {
	return t.ID
}

// SetID will set the JiraProjectComponent ID value
func (t *JiraProjectComponent) SetID(v string) {
	t.ID = v
}

// FindJiraProjectComponentByID will find a JiraProjectComponent by ID
func FindJiraProjectComponentByID(ctx context.Context, db *sql.DB, value string) (*JiraProjectComponent, error) {
	q := "SELECT `jira_project_component`.`id`,`jira_project_component`.`checksum`,`jira_project_component`.`component_id`,`jira_project_component`.`project_id`,`jira_project_component`.`name`,`jira_project_component`.`customer_id` FROM `jira_project_component` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _ComponentID sql.NullString
	var _ProjectID sql.NullString
	var _Name sql.NullString
	var _CustomerID sql.NullString
	err := db.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_ComponentID,
		&_ProjectID,
		&_Name,
		&_CustomerID,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &JiraProjectComponent{}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _Checksum.Valid {
		t.SetChecksum(_Checksum.String)
	}
	if _ComponentID.Valid {
		t.SetComponentID(_ComponentID.String)
	}
	if _ProjectID.Valid {
		t.SetProjectID(_ProjectID.String)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	return t, nil
}

// FindJiraProjectComponentByIDTx will find a JiraProjectComponent by ID using the provided transaction
func FindJiraProjectComponentByIDTx(ctx context.Context, tx *sql.Tx, value string) (*JiraProjectComponent, error) {
	q := "SELECT `jira_project_component`.`id`,`jira_project_component`.`checksum`,`jira_project_component`.`component_id`,`jira_project_component`.`project_id`,`jira_project_component`.`name`,`jira_project_component`.`customer_id` FROM `jira_project_component` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _ComponentID sql.NullString
	var _ProjectID sql.NullString
	var _Name sql.NullString
	var _CustomerID sql.NullString
	err := tx.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_ComponentID,
		&_ProjectID,
		&_Name,
		&_CustomerID,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &JiraProjectComponent{}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _Checksum.Valid {
		t.SetChecksum(_Checksum.String)
	}
	if _ComponentID.Valid {
		t.SetComponentID(_ComponentID.String)
	}
	if _ProjectID.Valid {
		t.SetProjectID(_ProjectID.String)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	return t, nil
}

// GetChecksum will return the JiraProjectComponent Checksum value
func (t *JiraProjectComponent) GetChecksum() string {
	if t.Checksum == nil {
		return ""
	}
	return *t.Checksum
}

// SetChecksum will set the JiraProjectComponent Checksum value
func (t *JiraProjectComponent) SetChecksum(v string) {
	t.Checksum = &v
}

// GetComponentID will return the JiraProjectComponent ComponentID value
func (t *JiraProjectComponent) GetComponentID() string {
	return t.ComponentID
}

// SetComponentID will set the JiraProjectComponent ComponentID value
func (t *JiraProjectComponent) SetComponentID(v string) {
	t.ComponentID = v
}

// GetProjectID will return the JiraProjectComponent ProjectID value
func (t *JiraProjectComponent) GetProjectID() string {
	return t.ProjectID
}

// SetProjectID will set the JiraProjectComponent ProjectID value
func (t *JiraProjectComponent) SetProjectID(v string) {
	t.ProjectID = v
}

// FindJiraProjectComponentsByProjectID will find all JiraProjectComponents by the ProjectID value
func FindJiraProjectComponentsByProjectID(ctx context.Context, db *sql.DB, value string) ([]*JiraProjectComponent, error) {
	q := "SELECT `jira_project_component`.`id`,`jira_project_component`.`checksum`,`jira_project_component`.`component_id`,`jira_project_component`.`project_id`,`jira_project_component`.`name`,`jira_project_component`.`customer_id` FROM `jira_project_component` WHERE `project_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraProjectComponent, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _ComponentID sql.NullString
		var _ProjectID sql.NullString
		var _Name sql.NullString
		var _CustomerID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_ComponentID,
			&_ProjectID,
			&_Name,
			&_CustomerID,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraProjectComponent{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _ComponentID.Valid {
			t.SetComponentID(_ComponentID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
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

// FindJiraProjectComponentsByProjectIDTx will find all JiraProjectComponents by the ProjectID value using the provided transaction
func FindJiraProjectComponentsByProjectIDTx(ctx context.Context, tx *sql.Tx, value string) ([]*JiraProjectComponent, error) {
	q := "SELECT `jira_project_component`.`id`,`jira_project_component`.`checksum`,`jira_project_component`.`component_id`,`jira_project_component`.`project_id`,`jira_project_component`.`name`,`jira_project_component`.`customer_id` FROM `jira_project_component` WHERE `project_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraProjectComponent, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _ComponentID sql.NullString
		var _ProjectID sql.NullString
		var _Name sql.NullString
		var _CustomerID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_ComponentID,
			&_ProjectID,
			&_Name,
			&_CustomerID,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraProjectComponent{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _ComponentID.Valid {
			t.SetComponentID(_ComponentID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
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

// GetName will return the JiraProjectComponent Name value
func (t *JiraProjectComponent) GetName() string {
	return t.Name
}

// SetName will set the JiraProjectComponent Name value
func (t *JiraProjectComponent) SetName(v string) {
	t.Name = v
}

// FindJiraProjectComponentsByName will find all JiraProjectComponents by the Name value
func FindJiraProjectComponentsByName(ctx context.Context, db *sql.DB, value string) ([]*JiraProjectComponent, error) {
	q := "SELECT `jira_project_component`.`id`,`jira_project_component`.`checksum`,`jira_project_component`.`component_id`,`jira_project_component`.`project_id`,`jira_project_component`.`name`,`jira_project_component`.`customer_id` FROM `jira_project_component` WHERE `name` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraProjectComponent, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _ComponentID sql.NullString
		var _ProjectID sql.NullString
		var _Name sql.NullString
		var _CustomerID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_ComponentID,
			&_ProjectID,
			&_Name,
			&_CustomerID,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraProjectComponent{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _ComponentID.Valid {
			t.SetComponentID(_ComponentID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
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

// FindJiraProjectComponentsByNameTx will find all JiraProjectComponents by the Name value using the provided transaction
func FindJiraProjectComponentsByNameTx(ctx context.Context, tx *sql.Tx, value string) ([]*JiraProjectComponent, error) {
	q := "SELECT `jira_project_component`.`id`,`jira_project_component`.`checksum`,`jira_project_component`.`component_id`,`jira_project_component`.`project_id`,`jira_project_component`.`name`,`jira_project_component`.`customer_id` FROM `jira_project_component` WHERE `name` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraProjectComponent, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _ComponentID sql.NullString
		var _ProjectID sql.NullString
		var _Name sql.NullString
		var _CustomerID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_ComponentID,
			&_ProjectID,
			&_Name,
			&_CustomerID,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraProjectComponent{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _ComponentID.Valid {
			t.SetComponentID(_ComponentID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
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

// GetCustomerID will return the JiraProjectComponent CustomerID value
func (t *JiraProjectComponent) GetCustomerID() string {
	return t.CustomerID
}

// SetCustomerID will set the JiraProjectComponent CustomerID value
func (t *JiraProjectComponent) SetCustomerID(v string) {
	t.CustomerID = v
}

// FindJiraProjectComponentsByCustomerID will find all JiraProjectComponents by the CustomerID value
func FindJiraProjectComponentsByCustomerID(ctx context.Context, db *sql.DB, value string) ([]*JiraProjectComponent, error) {
	q := "SELECT `jira_project_component`.`id`,`jira_project_component`.`checksum`,`jira_project_component`.`component_id`,`jira_project_component`.`project_id`,`jira_project_component`.`name`,`jira_project_component`.`customer_id` FROM `jira_project_component` WHERE `customer_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraProjectComponent, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _ComponentID sql.NullString
		var _ProjectID sql.NullString
		var _Name sql.NullString
		var _CustomerID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_ComponentID,
			&_ProjectID,
			&_Name,
			&_CustomerID,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraProjectComponent{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _ComponentID.Valid {
			t.SetComponentID(_ComponentID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
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

// FindJiraProjectComponentsByCustomerIDTx will find all JiraProjectComponents by the CustomerID value using the provided transaction
func FindJiraProjectComponentsByCustomerIDTx(ctx context.Context, tx *sql.Tx, value string) ([]*JiraProjectComponent, error) {
	q := "SELECT `jira_project_component`.`id`,`jira_project_component`.`checksum`,`jira_project_component`.`component_id`,`jira_project_component`.`project_id`,`jira_project_component`.`name`,`jira_project_component`.`customer_id` FROM `jira_project_component` WHERE `customer_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraProjectComponent, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _ComponentID sql.NullString
		var _ProjectID sql.NullString
		var _Name sql.NullString
		var _CustomerID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_ComponentID,
			&_ProjectID,
			&_Name,
			&_CustomerID,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraProjectComponent{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _ComponentID.Valid {
			t.SetComponentID(_ComponentID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
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

func (t *JiraProjectComponent) toTimestamp(value time.Time) *timestamp.Timestamp {
	ts, _ := ptypes.TimestampProto(value)
	return ts
}

// DBCreateJiraProjectComponentTable will create the JiraProjectComponent table
func DBCreateJiraProjectComponentTable(ctx context.Context, db *sql.DB) error {
	q := "CREATE TABLE `jira_project_component` (`id`VARCHAR(64) NOT NULL PRIMARY KEY,`checksum`CHAR(64),`component_id` VARCHAR(64) NOT NULL,`project_id` VARCHAR(64) NOT NULL,`name` VARCHAR(255) NOT NULL,`customer_id`VARCHAR(64) NOT NULL,INDEX jira_project_component_project_id_index (`project_id`),INDEX jira_project_component_name_index (`name`),INDEX jira_project_component_customer_id_index (`customer_id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBCreateJiraProjectComponentTableTx will create the JiraProjectComponent table using the provided transction
func DBCreateJiraProjectComponentTableTx(ctx context.Context, tx *sql.Tx) error {
	q := "CREATE TABLE `jira_project_component` (`id`VARCHAR(64) NOT NULL PRIMARY KEY,`checksum`CHAR(64),`component_id` VARCHAR(64) NOT NULL,`project_id` VARCHAR(64) NOT NULL,`name` VARCHAR(255) NOT NULL,`customer_id`VARCHAR(64) NOT NULL,INDEX jira_project_component_project_id_index (`project_id`),INDEX jira_project_component_name_index (`name`),INDEX jira_project_component_customer_id_index (`customer_id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// DBDropJiraProjectComponentTable will drop the JiraProjectComponent table
func DBDropJiraProjectComponentTable(ctx context.Context, db *sql.DB) error {
	q := "DROP TABLE IF EXISTS `jira_project_component`"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBDropJiraProjectComponentTableTx will drop the JiraProjectComponent table using the provided transaction
func DBDropJiraProjectComponentTableTx(ctx context.Context, tx *sql.Tx) error {
	q := "DROP TABLE IF EXISTS `jira_project_component`"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// CalculateChecksum will calculate a checksum of the SHA1 of all field values
func (t *JiraProjectComponent) CalculateChecksum() string {
	return orm.HashStrings(
		orm.ToString(t.ID),
		orm.ToString(t.ComponentID),
		orm.ToString(t.ProjectID),
		orm.ToString(t.Name),
		orm.ToString(t.CustomerID),
	)
}

// DBCreate will create a new JiraProjectComponent record in the database
func (t *JiraProjectComponent) DBCreate(ctx context.Context, db *sql.DB) (sql.Result, error) {
	q := "INSERT INTO `jira_project_component` (`jira_project_component`.`id`,`jira_project_component`.`checksum`,`jira_project_component`.`component_id`,`jira_project_component`.`project_id`,`jira_project_component`.`name`,`jira_project_component`.`customer_id`) VALUES (?,?,?,?,?,?)"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.ComponentID),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.CustomerID),
	)
}

// DBCreateTx will create a new JiraProjectComponent record in the database using the provided transaction
func (t *JiraProjectComponent) DBCreateTx(ctx context.Context, tx *sql.Tx) (sql.Result, error) {
	q := "INSERT INTO `jira_project_component` (`jira_project_component`.`id`,`jira_project_component`.`checksum`,`jira_project_component`.`component_id`,`jira_project_component`.`project_id`,`jira_project_component`.`name`,`jira_project_component`.`customer_id`) VALUES (?,?,?,?,?,?)"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.ComponentID),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.CustomerID),
	)
}

// DBCreateIgnoreDuplicate will upsert the JiraProjectComponent record in the database
func (t *JiraProjectComponent) DBCreateIgnoreDuplicate(ctx context.Context, db *sql.DB) (sql.Result, error) {
	q := "INSERT INTO `jira_project_component` (`jira_project_component`.`id`,`jira_project_component`.`checksum`,`jira_project_component`.`component_id`,`jira_project_component`.`project_id`,`jira_project_component`.`name`,`jira_project_component`.`customer_id`) VALUES (?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.ComponentID),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.CustomerID),
	)
}

// DBCreateIgnoreDuplicateTx will upsert the JiraProjectComponent record in the database using the provided transaction
func (t *JiraProjectComponent) DBCreateIgnoreDuplicateTx(ctx context.Context, tx *sql.Tx) (sql.Result, error) {
	q := "INSERT INTO `jira_project_component` (`jira_project_component`.`id`,`jira_project_component`.`checksum`,`jira_project_component`.`component_id`,`jira_project_component`.`project_id`,`jira_project_component`.`name`,`jira_project_component`.`customer_id`) VALUES (?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.ComponentID),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.CustomerID),
	)
}

// DeleteAllJiraProjectComponents deletes all JiraProjectComponent records in the database with optional filters
func DeleteAllJiraProjectComponents(ctx context.Context, db *sql.DB, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(JiraProjectComponentTableName),
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

// DeleteAllJiraProjectComponentsTx deletes all JiraProjectComponent records in the database with optional filters using the provided transaction
func DeleteAllJiraProjectComponentsTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(JiraProjectComponentTableName),
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

// DBDelete will delete this JiraProjectComponent record in the database
func (t *JiraProjectComponent) DBDelete(ctx context.Context, db *sql.DB) (bool, error) {
	q := "DELETE FROM `jira_project_component` WHERE `id` = ?"
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

// DBDeleteTx will delete this JiraProjectComponent record in the database using the provided transaction
func (t *JiraProjectComponent) DBDeleteTx(ctx context.Context, tx *sql.Tx) (bool, error) {
	q := "DELETE FROM `jira_project_component` WHERE `id` = ?"
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

// DBUpdate will update the JiraProjectComponent record in the database
func (t *JiraProjectComponent) DBUpdate(ctx context.Context, db *sql.DB) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `jira_project_component` SET `checksum`=?,`component_id`=?,`project_id`=?,`name`=?,`customer_id`=? WHERE `id`=?"
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.ComponentID),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.ID),
	)
}

// DBUpdateTx will update the JiraProjectComponent record in the database using the provided transaction
func (t *JiraProjectComponent) DBUpdateTx(ctx context.Context, tx *sql.Tx) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `jira_project_component` SET `checksum`=?,`component_id`=?,`project_id`=?,`name`=?,`customer_id`=? WHERE `id`=?"
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.ComponentID),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.ID),
	)
}

// DBUpsert will upsert the JiraProjectComponent record in the database
func (t *JiraProjectComponent) DBUpsert(ctx context.Context, db *sql.DB, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `jira_project_component` (`jira_project_component`.`id`,`jira_project_component`.`checksum`,`jira_project_component`.`component_id`,`jira_project_component`.`project_id`,`jira_project_component`.`name`,`jira_project_component`.`customer_id`) VALUES (?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `jira_project_component` (`jira_project_component`.`id`,`jira_project_component`.`checksum`,`jira_project_component`.`component_id`,`jira_project_component`.`project_id`,`jira_project_component`.`name`,`jira_project_component`.`customer_id`) VALUES (?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`component_id`=VALUES(`component_id`),`project_id`=VALUES(`project_id`),`name`=VALUES(`name`),`customer_id`=VALUES(`customer_id`)"
	}
	r, err := db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.ComponentID),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.CustomerID),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBUpsertTx will upsert the JiraProjectComponent record in the database using the provided transaction
func (t *JiraProjectComponent) DBUpsertTx(ctx context.Context, tx *sql.Tx, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `jira_project_component` (`jira_project_component`.`id`,`jira_project_component`.`checksum`,`jira_project_component`.`component_id`,`jira_project_component`.`project_id`,`jira_project_component`.`name`,`jira_project_component`.`customer_id`) VALUES (?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `jira_project_component` (`jira_project_component`.`id`,`jira_project_component`.`checksum`,`jira_project_component`.`component_id`,`jira_project_component`.`project_id`,`jira_project_component`.`name`,`jira_project_component`.`customer_id`) VALUES (?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`component_id`=VALUES(`component_id`),`project_id`=VALUES(`project_id`),`name`=VALUES(`name`),`customer_id`=VALUES(`customer_id`)"
	}
	r, err := tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.ComponentID),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.CustomerID),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBFindOne will find a JiraProjectComponent record in the database with the primary key
func (t *JiraProjectComponent) DBFindOne(ctx context.Context, db *sql.DB, value string) (bool, error) {
	q := "SELECT `jira_project_component`.`id`,`jira_project_component`.`checksum`,`jira_project_component`.`component_id`,`jira_project_component`.`project_id`,`jira_project_component`.`name`,`jira_project_component`.`customer_id` FROM `jira_project_component` WHERE `id` = ? LIMIT 1"
	row := db.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _ComponentID sql.NullString
	var _ProjectID sql.NullString
	var _Name sql.NullString
	var _CustomerID sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_ComponentID,
		&_ProjectID,
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
	if _ComponentID.Valid {
		t.SetComponentID(_ComponentID.String)
	}
	if _ProjectID.Valid {
		t.SetProjectID(_ProjectID.String)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	return true, nil
}

// DBFindOneTx will find a JiraProjectComponent record in the database with the primary key using the provided transaction
func (t *JiraProjectComponent) DBFindOneTx(ctx context.Context, tx *sql.Tx, value string) (bool, error) {
	q := "SELECT `jira_project_component`.`id`,`jira_project_component`.`checksum`,`jira_project_component`.`component_id`,`jira_project_component`.`project_id`,`jira_project_component`.`name`,`jira_project_component`.`customer_id` FROM `jira_project_component` WHERE `id` = ? LIMIT 1"
	row := tx.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _ComponentID sql.NullString
	var _ProjectID sql.NullString
	var _Name sql.NullString
	var _CustomerID sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_ComponentID,
		&_ProjectID,
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
	if _ComponentID.Valid {
		t.SetComponentID(_ComponentID.String)
	}
	if _ProjectID.Valid {
		t.SetProjectID(_ProjectID.String)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	return true, nil
}

// FindJiraProjectComponents will find a JiraProjectComponent record in the database with the provided parameters
func FindJiraProjectComponents(ctx context.Context, db *sql.DB, _params ...interface{}) ([]*JiraProjectComponent, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("component_id"),
		orm.Column("project_id"),
		orm.Column("name"),
		orm.Column("customer_id"),
		orm.Table(JiraProjectComponentTableName),
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
	results := make([]*JiraProjectComponent, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _ComponentID sql.NullString
		var _ProjectID sql.NullString
		var _Name sql.NullString
		var _CustomerID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_ComponentID,
			&_ProjectID,
			&_Name,
			&_CustomerID,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraProjectComponent{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _ComponentID.Valid {
			t.SetComponentID(_ComponentID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
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

// FindJiraProjectComponentsTx will find a JiraProjectComponent record in the database with the provided parameters using the provided transaction
func FindJiraProjectComponentsTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) ([]*JiraProjectComponent, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("component_id"),
		orm.Column("project_id"),
		orm.Column("name"),
		orm.Column("customer_id"),
		orm.Table(JiraProjectComponentTableName),
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
	results := make([]*JiraProjectComponent, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _ComponentID sql.NullString
		var _ProjectID sql.NullString
		var _Name sql.NullString
		var _CustomerID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_ComponentID,
			&_ProjectID,
			&_Name,
			&_CustomerID,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraProjectComponent{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _ComponentID.Valid {
			t.SetComponentID(_ComponentID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
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

// DBFind will find a JiraProjectComponent record in the database with the provided parameters
func (t *JiraProjectComponent) DBFind(ctx context.Context, db *sql.DB, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("component_id"),
		orm.Column("project_id"),
		orm.Column("name"),
		orm.Column("customer_id"),
		orm.Table(JiraProjectComponentTableName),
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
	var _ComponentID sql.NullString
	var _ProjectID sql.NullString
	var _Name sql.NullString
	var _CustomerID sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_ComponentID,
		&_ProjectID,
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
	if _ComponentID.Valid {
		t.SetComponentID(_ComponentID.String)
	}
	if _ProjectID.Valid {
		t.SetProjectID(_ProjectID.String)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	return true, nil
}

// DBFindTx will find a JiraProjectComponent record in the database with the provided parameters using the provided transaction
func (t *JiraProjectComponent) DBFindTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("component_id"),
		orm.Column("project_id"),
		orm.Column("name"),
		orm.Column("customer_id"),
		orm.Table(JiraProjectComponentTableName),
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
	var _ComponentID sql.NullString
	var _ProjectID sql.NullString
	var _Name sql.NullString
	var _CustomerID sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_ComponentID,
		&_ProjectID,
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
	if _ComponentID.Valid {
		t.SetComponentID(_ComponentID.String)
	}
	if _ProjectID.Valid {
		t.SetProjectID(_ProjectID.String)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	return true, nil
}

// CountJiraProjectComponents will find the count of JiraProjectComponent records in the database
func CountJiraProjectComponents(ctx context.Context, db *sql.DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(JiraProjectComponentTableName),
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

// CountJiraProjectComponentsTx will find the count of JiraProjectComponent records in the database using the provided transaction
func CountJiraProjectComponentsTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(JiraProjectComponentTableName),
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

// DBCount will find the count of JiraProjectComponent records in the database
func (t *JiraProjectComponent) DBCount(ctx context.Context, db *sql.DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(JiraProjectComponentTableName),
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

// DBCountTx will find the count of JiraProjectComponent records in the database using the provided transaction
func (t *JiraProjectComponent) DBCountTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(JiraProjectComponentTableName),
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

// DBExists will return true if the JiraProjectComponent record exists in the database
func (t *JiraProjectComponent) DBExists(ctx context.Context, db *sql.DB) (bool, error) {
	q := "SELECT `id` FROM `jira_project_component` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := db.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// DBExistsTx will return true if the JiraProjectComponent record exists in the database using the provided transaction
func (t *JiraProjectComponent) DBExistsTx(ctx context.Context, tx *sql.Tx) (bool, error) {
	q := "SELECT `id` FROM `jira_project_component` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := tx.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// PrimaryKeyColumn returns the column name for the primary key
func (t *JiraProjectComponent) PrimaryKeyColumn() string {
	return JiraProjectComponentColumnID
}

// PrimaryKeyColumnType returns the primary key column Go type as a string
func (t *JiraProjectComponent) PrimaryKeyColumnType() string {
	return "string"
}

// PrimaryKey returns the primary key column value
func (t *JiraProjectComponent) PrimaryKey() interface{} {
	return t.ID
}
