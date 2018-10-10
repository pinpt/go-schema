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

var _ Model = (*JiraProject)(nil)
var _ CSVWriter = (*JiraProject)(nil)
var _ JSONWriter = (*JiraProject)(nil)
var _ Checksum = (*JiraProject)(nil)

// JiraProjectTableName is the name of the table in SQL
const JiraProjectTableName = "jira_project"

var JiraProjectColumns = []string{
	"id",
	"checksum",
	"project_id",
	"key",
	"avatar_url",
	"category_id",
	"customer_id",
	"ref_id",
}

// JiraProject table
type JiraProject struct {
	AvatarURL  string  `json:"avatar_url"`
	CategoryID *string `json:"category_id,omitempty"`
	Checksum   *string `json:"checksum,omitempty"`
	CustomerID string  `json:"customer_id"`
	ID         string  `json:"id"`
	Key        string  `json:"key"`
	ProjectID  string  `json:"project_id"`
	RefID      string  `json:"ref_id"`
}

// TableName returns the SQL table name for JiraProject and satifies the Model interface
func (t *JiraProject) TableName() string {
	return JiraProjectTableName
}

// ToCSV will serialize the JiraProject instance to a CSV compatible array of strings
func (t *JiraProject) ToCSV() []string {
	return []string{
		t.ID,
		t.CalculateChecksum(),
		t.ProjectID,
		t.Key,
		t.AvatarURL,
		toCSVString(t.CategoryID),
		t.CustomerID,
		t.RefID,
	}
}

// WriteCSV will serialize the JiraProject instance to the writer as CSV and satisfies the CSVWriter interface
func (t *JiraProject) WriteCSV(w *csv.Writer) error {
	return w.Write(t.ToCSV())
}

// WriteJSON will serialize the JiraProject instance to the writer as JSON and satisfies the JSONWriter interface
func (t *JiraProject) WriteJSON(w io.Writer, indent ...bool) error {
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

// NewJiraProjectReader creates a JSON reader which can read in JiraProject objects serialized as JSON either as an array, single object or json new lines
// and writes each JiraProject to the channel provided
func NewJiraProjectReader(r io.Reader, ch chan<- JiraProject) error {
	return orm.Deserialize(r, func(buf json.RawMessage) error {
		dec := json.NewDecoder(bytes.NewBuffer(buf))
		e := JiraProject{}
		if err := dec.Decode(&e); err != nil {
			return err
		}
		ch <- e
		return nil
	})
}

// NewCSVJiraProjectReaderDir will read the reader as CSV and emit each record to the channel provided
func NewCSVJiraProjectReader(r io.Reader, ch chan<- JiraProject) error {
	cr := csv.NewReader(r)
	for {
		record, err := cr.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		ch <- JiraProject{
			ID:         record[0],
			Checksum:   fromStringPointer(record[1]),
			ProjectID:  record[2],
			Key:        record[3],
			AvatarURL:  record[4],
			CategoryID: fromStringPointer(record[5]),
			CustomerID: record[6],
			RefID:      record[7],
		}
	}
	return nil
}

// NewCSVJiraProjectReaderFile will read the file as a CSV and emit each record to the channel provided
func NewCSVJiraProjectReaderFile(fp string, ch chan<- JiraProject) error {
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
	return NewCSVJiraProjectReader(fc, ch)
}

// NewCSVJiraProjectReaderDir will read the jira_project.csv.gz file as a CSV and emit each record to the channel provided
func NewCSVJiraProjectReaderDir(dir string, ch chan<- JiraProject) error {
	return NewCSVJiraProjectReaderFile(filepath.Join(dir, "jira_project.csv.gz"), ch)
}

// JiraProjectCSVDeduper is a function callback which takes the existing value (a) and the new value (b)
// and the return value should be the one to use (or a new one, if applicable). return nil
// to skip processing of this record
type JiraProjectCSVDeduper func(a JiraProject, b JiraProject) *JiraProject

// JiraProjectCSVDedupeDisabled is set on whether the CSV writer should de-dupe values by key
var JiraProjectCSVDedupeDisabled bool

// NewJiraProjectCSVWriterSize creates a batch writer that will write each JiraProject into a CSV file
// this method will automatically de-duplicate entries using the primary key. if the checksum
// for a newer item with the same primary key doesn't match a previously sent item, the newer
// one will be used
func NewJiraProjectCSVWriterSize(w io.Writer, size int, dedupers ...JiraProjectCSVDeduper) (chan JiraProject, chan bool, error) {
	cw := csv.NewWriter(w)
	ch := make(chan JiraProject, size)
	done := make(chan bool)
	go func() {
		defer func() { done <- true }()
		dodedupe := !JiraProjectCSVDedupeDisabled
		var kv map[string]*JiraProject
		var deduper JiraProjectCSVDeduper
		if dedupers != nil && len(dedupers) > 0 {
			deduper = dedupers[0]
			dodedupe = true
		}
		if dodedupe {
			kv = make(map[string]*JiraProject)
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

// JiraProjectCSVDefaultSize is the default channel buffer size if not provided
var JiraProjectCSVDefaultSize = 100

// NewJiraProjectCSVWriter creates a batch writer that will write each JiraProject into a CSV file
func NewJiraProjectCSVWriter(w io.Writer, dedupers ...JiraProjectCSVDeduper) (chan JiraProject, chan bool, error) {
	return NewJiraProjectCSVWriterSize(w, JiraProjectCSVDefaultSize, dedupers...)
}

// NewJiraProjectCSVWriterDir creates a batch writer that will write each JiraProject into a CSV file named jira_project.csv.gz in dir
func NewJiraProjectCSVWriterDir(dir string, dedupers ...JiraProjectCSVDeduper) (chan JiraProject, chan bool, error) {
	return NewJiraProjectCSVWriterFile(filepath.Join(dir, "jira_project.csv.gz"), dedupers...)
}

// NewJiraProjectCSVWriterFile creates a batch writer that will write each JiraProject into a CSV file
func NewJiraProjectCSVWriterFile(fn string, dedupers ...JiraProjectCSVDeduper) (chan JiraProject, chan bool, error) {
	f, err := os.Create(fn)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening CSV file %s. %v", fn, err)
	}
	var fc io.WriteCloser = f
	if filepath.Ext(fn) == ".gz" {
		w, _ := gzip.NewWriterLevel(f, gzip.BestCompression)
		fc = w
	}
	ch, done, err := NewJiraProjectCSVWriter(fc, dedupers...)
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

type JiraProjectDBAction func(ctx context.Context, db DB, record JiraProject) error

// NewJiraProjectDBWriterSize creates a DB writer that will write each issue into the DB
func NewJiraProjectDBWriterSize(ctx context.Context, db DB, errors chan<- error, size int, actions ...JiraProjectDBAction) (chan JiraProject, chan bool, error) {
	ch := make(chan JiraProject, size)
	done := make(chan bool)
	var action JiraProjectDBAction
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

// NewJiraProjectDBWriter creates a DB writer that will write each issue into the DB
func NewJiraProjectDBWriter(ctx context.Context, db DB, errors chan<- error, actions ...JiraProjectDBAction) (chan JiraProject, chan bool, error) {
	return NewJiraProjectDBWriterSize(ctx, db, errors, 100, actions...)
}

// JiraProjectColumnID is the ID SQL column name for the JiraProject table
const JiraProjectColumnID = "id"

// JiraProjectEscapedColumnID is the escaped ID SQL column name for the JiraProject table
const JiraProjectEscapedColumnID = "`id`"

// JiraProjectColumnChecksum is the Checksum SQL column name for the JiraProject table
const JiraProjectColumnChecksum = "checksum"

// JiraProjectEscapedColumnChecksum is the escaped Checksum SQL column name for the JiraProject table
const JiraProjectEscapedColumnChecksum = "`checksum`"

// JiraProjectColumnProjectID is the ProjectID SQL column name for the JiraProject table
const JiraProjectColumnProjectID = "project_id"

// JiraProjectEscapedColumnProjectID is the escaped ProjectID SQL column name for the JiraProject table
const JiraProjectEscapedColumnProjectID = "`project_id`"

// JiraProjectColumnKey is the Key SQL column name for the JiraProject table
const JiraProjectColumnKey = "key"

// JiraProjectEscapedColumnKey is the escaped Key SQL column name for the JiraProject table
const JiraProjectEscapedColumnKey = "`key`"

// JiraProjectColumnAvatarURL is the AvatarURL SQL column name for the JiraProject table
const JiraProjectColumnAvatarURL = "avatar_url"

// JiraProjectEscapedColumnAvatarURL is the escaped AvatarURL SQL column name for the JiraProject table
const JiraProjectEscapedColumnAvatarURL = "`avatar_url`"

// JiraProjectColumnCategoryID is the CategoryID SQL column name for the JiraProject table
const JiraProjectColumnCategoryID = "category_id"

// JiraProjectEscapedColumnCategoryID is the escaped CategoryID SQL column name for the JiraProject table
const JiraProjectEscapedColumnCategoryID = "`category_id`"

// JiraProjectColumnCustomerID is the CustomerID SQL column name for the JiraProject table
const JiraProjectColumnCustomerID = "customer_id"

// JiraProjectEscapedColumnCustomerID is the escaped CustomerID SQL column name for the JiraProject table
const JiraProjectEscapedColumnCustomerID = "`customer_id`"

// JiraProjectColumnRefID is the RefID SQL column name for the JiraProject table
const JiraProjectColumnRefID = "ref_id"

// JiraProjectEscapedColumnRefID is the escaped RefID SQL column name for the JiraProject table
const JiraProjectEscapedColumnRefID = "`ref_id`"

// GetID will return the JiraProject ID value
func (t *JiraProject) GetID() string {
	return t.ID
}

// SetID will set the JiraProject ID value
func (t *JiraProject) SetID(v string) {
	t.ID = v
}

// FindJiraProjectByID will find a JiraProject by ID
func FindJiraProjectByID(ctx context.Context, db DB, value string) (*JiraProject, error) {
	q := "SELECT * FROM `jira_project` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _ProjectID sql.NullString
	var _Key sql.NullString
	var _AvatarURL sql.NullString
	var _CategoryID sql.NullString
	var _CustomerID sql.NullString
	var _RefID sql.NullString
	err := db.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_ProjectID,
		&_Key,
		&_AvatarURL,
		&_CategoryID,
		&_CustomerID,
		&_RefID,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &JiraProject{}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _Checksum.Valid {
		t.SetChecksum(_Checksum.String)
	}
	if _ProjectID.Valid {
		t.SetProjectID(_ProjectID.String)
	}
	if _Key.Valid {
		t.SetKey(_Key.String)
	}
	if _AvatarURL.Valid {
		t.SetAvatarURL(_AvatarURL.String)
	}
	if _CategoryID.Valid {
		t.SetCategoryID(_CategoryID.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _RefID.Valid {
		t.SetRefID(_RefID.String)
	}
	return t, nil
}

// FindJiraProjectByIDTx will find a JiraProject by ID using the provided transaction
func FindJiraProjectByIDTx(ctx context.Context, tx Tx, value string) (*JiraProject, error) {
	q := "SELECT * FROM `jira_project` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _ProjectID sql.NullString
	var _Key sql.NullString
	var _AvatarURL sql.NullString
	var _CategoryID sql.NullString
	var _CustomerID sql.NullString
	var _RefID sql.NullString
	err := tx.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_ProjectID,
		&_Key,
		&_AvatarURL,
		&_CategoryID,
		&_CustomerID,
		&_RefID,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &JiraProject{}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _Checksum.Valid {
		t.SetChecksum(_Checksum.String)
	}
	if _ProjectID.Valid {
		t.SetProjectID(_ProjectID.String)
	}
	if _Key.Valid {
		t.SetKey(_Key.String)
	}
	if _AvatarURL.Valid {
		t.SetAvatarURL(_AvatarURL.String)
	}
	if _CategoryID.Valid {
		t.SetCategoryID(_CategoryID.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _RefID.Valid {
		t.SetRefID(_RefID.String)
	}
	return t, nil
}

// GetChecksum will return the JiraProject Checksum value
func (t *JiraProject) GetChecksum() string {
	if t.Checksum == nil {
		return ""
	}
	return *t.Checksum
}

// SetChecksum will set the JiraProject Checksum value
func (t *JiraProject) SetChecksum(v string) {
	t.Checksum = &v
}

// GetProjectID will return the JiraProject ProjectID value
func (t *JiraProject) GetProjectID() string {
	return t.ProjectID
}

// SetProjectID will set the JiraProject ProjectID value
func (t *JiraProject) SetProjectID(v string) {
	t.ProjectID = v
}

// FindJiraProjectsByProjectID will find all JiraProjects by the ProjectID value
func FindJiraProjectsByProjectID(ctx context.Context, db DB, value string) ([]*JiraProject, error) {
	q := "SELECT * FROM `jira_project` WHERE `project_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraProject, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _ProjectID sql.NullString
		var _Key sql.NullString
		var _AvatarURL sql.NullString
		var _CategoryID sql.NullString
		var _CustomerID sql.NullString
		var _RefID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_ProjectID,
			&_Key,
			&_AvatarURL,
			&_CategoryID,
			&_CustomerID,
			&_RefID,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraProject{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _Key.Valid {
			t.SetKey(_Key.String)
		}
		if _AvatarURL.Valid {
			t.SetAvatarURL(_AvatarURL.String)
		}
		if _CategoryID.Valid {
			t.SetCategoryID(_CategoryID.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindJiraProjectsByProjectIDTx will find all JiraProjects by the ProjectID value using the provided transaction
func FindJiraProjectsByProjectIDTx(ctx context.Context, tx Tx, value string) ([]*JiraProject, error) {
	q := "SELECT * FROM `jira_project` WHERE `project_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraProject, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _ProjectID sql.NullString
		var _Key sql.NullString
		var _AvatarURL sql.NullString
		var _CategoryID sql.NullString
		var _CustomerID sql.NullString
		var _RefID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_ProjectID,
			&_Key,
			&_AvatarURL,
			&_CategoryID,
			&_CustomerID,
			&_RefID,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraProject{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _Key.Valid {
			t.SetKey(_Key.String)
		}
		if _AvatarURL.Valid {
			t.SetAvatarURL(_AvatarURL.String)
		}
		if _CategoryID.Valid {
			t.SetCategoryID(_CategoryID.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetKey will return the JiraProject Key value
func (t *JiraProject) GetKey() string {
	return t.Key
}

// SetKey will set the JiraProject Key value
func (t *JiraProject) SetKey(v string) {
	t.Key = v
}

// GetAvatarURL will return the JiraProject AvatarURL value
func (t *JiraProject) GetAvatarURL() string {
	return t.AvatarURL
}

// SetAvatarURL will set the JiraProject AvatarURL value
func (t *JiraProject) SetAvatarURL(v string) {
	t.AvatarURL = v
}

// GetCategoryID will return the JiraProject CategoryID value
func (t *JiraProject) GetCategoryID() string {
	if t.CategoryID == nil {
		return ""
	}
	return *t.CategoryID
}

// SetCategoryID will set the JiraProject CategoryID value
func (t *JiraProject) SetCategoryID(v string) {
	t.CategoryID = &v
}

// GetCustomerID will return the JiraProject CustomerID value
func (t *JiraProject) GetCustomerID() string {
	return t.CustomerID
}

// SetCustomerID will set the JiraProject CustomerID value
func (t *JiraProject) SetCustomerID(v string) {
	t.CustomerID = v
}

// FindJiraProjectsByCustomerID will find all JiraProjects by the CustomerID value
func FindJiraProjectsByCustomerID(ctx context.Context, db DB, value string) ([]*JiraProject, error) {
	q := "SELECT * FROM `jira_project` WHERE `customer_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraProject, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _ProjectID sql.NullString
		var _Key sql.NullString
		var _AvatarURL sql.NullString
		var _CategoryID sql.NullString
		var _CustomerID sql.NullString
		var _RefID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_ProjectID,
			&_Key,
			&_AvatarURL,
			&_CategoryID,
			&_CustomerID,
			&_RefID,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraProject{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _Key.Valid {
			t.SetKey(_Key.String)
		}
		if _AvatarURL.Valid {
			t.SetAvatarURL(_AvatarURL.String)
		}
		if _CategoryID.Valid {
			t.SetCategoryID(_CategoryID.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindJiraProjectsByCustomerIDTx will find all JiraProjects by the CustomerID value using the provided transaction
func FindJiraProjectsByCustomerIDTx(ctx context.Context, tx Tx, value string) ([]*JiraProject, error) {
	q := "SELECT * FROM `jira_project` WHERE `customer_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraProject, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _ProjectID sql.NullString
		var _Key sql.NullString
		var _AvatarURL sql.NullString
		var _CategoryID sql.NullString
		var _CustomerID sql.NullString
		var _RefID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_ProjectID,
			&_Key,
			&_AvatarURL,
			&_CategoryID,
			&_CustomerID,
			&_RefID,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraProject{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _Key.Valid {
			t.SetKey(_Key.String)
		}
		if _AvatarURL.Valid {
			t.SetAvatarURL(_AvatarURL.String)
		}
		if _CategoryID.Valid {
			t.SetCategoryID(_CategoryID.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetRefID will return the JiraProject RefID value
func (t *JiraProject) GetRefID() string {
	return t.RefID
}

// SetRefID will set the JiraProject RefID value
func (t *JiraProject) SetRefID(v string) {
	t.RefID = v
}

// FindJiraProjectsByRefID will find all JiraProjects by the RefID value
func FindJiraProjectsByRefID(ctx context.Context, db DB, value string) ([]*JiraProject, error) {
	q := "SELECT * FROM `jira_project` WHERE `ref_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraProject, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _ProjectID sql.NullString
		var _Key sql.NullString
		var _AvatarURL sql.NullString
		var _CategoryID sql.NullString
		var _CustomerID sql.NullString
		var _RefID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_ProjectID,
			&_Key,
			&_AvatarURL,
			&_CategoryID,
			&_CustomerID,
			&_RefID,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraProject{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _Key.Valid {
			t.SetKey(_Key.String)
		}
		if _AvatarURL.Valid {
			t.SetAvatarURL(_AvatarURL.String)
		}
		if _CategoryID.Valid {
			t.SetCategoryID(_CategoryID.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindJiraProjectsByRefIDTx will find all JiraProjects by the RefID value using the provided transaction
func FindJiraProjectsByRefIDTx(ctx context.Context, tx Tx, value string) ([]*JiraProject, error) {
	q := "SELECT * FROM `jira_project` WHERE `ref_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraProject, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _ProjectID sql.NullString
		var _Key sql.NullString
		var _AvatarURL sql.NullString
		var _CategoryID sql.NullString
		var _CustomerID sql.NullString
		var _RefID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_ProjectID,
			&_Key,
			&_AvatarURL,
			&_CategoryID,
			&_CustomerID,
			&_RefID,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraProject{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _Key.Valid {
			t.SetKey(_Key.String)
		}
		if _AvatarURL.Valid {
			t.SetAvatarURL(_AvatarURL.String)
		}
		if _CategoryID.Valid {
			t.SetCategoryID(_CategoryID.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

func (t *JiraProject) toTimestamp(value time.Time) *timestamp.Timestamp {
	ts, _ := ptypes.TimestampProto(value)
	return ts
}

// DBCreateJiraProjectTable will create the JiraProject table
func DBCreateJiraProjectTable(ctx context.Context, db DB) error {
	q := "CREATE TABLE `jira_project` (`id` VARCHAR(64) NOT NULL PRIMARY KEY,`checksum` CHAR(64),`project_id`VARCHAR(255) NOT NULL,`key` VARCHAR(100) NOT NULL,`avatar_url`VARCHAR(255) NOT NULL,`category_id` VARCHAR(64),`customer_id` VARCHAR(64) NOT NULL,`ref_id` VARCHAR(64) NOT NULL,INDEX jira_project_project_id_index (`project_id`),INDEX jira_project_customer_id_index (`customer_id`),INDEX jira_project_ref_id_index (`ref_id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBCreateJiraProjectTableTx will create the JiraProject table using the provided transction
func DBCreateJiraProjectTableTx(ctx context.Context, tx Tx) error {
	q := "CREATE TABLE `jira_project` (`id` VARCHAR(64) NOT NULL PRIMARY KEY,`checksum` CHAR(64),`project_id`VARCHAR(255) NOT NULL,`key` VARCHAR(100) NOT NULL,`avatar_url`VARCHAR(255) NOT NULL,`category_id` VARCHAR(64),`customer_id` VARCHAR(64) NOT NULL,`ref_id` VARCHAR(64) NOT NULL,INDEX jira_project_project_id_index (`project_id`),INDEX jira_project_customer_id_index (`customer_id`),INDEX jira_project_ref_id_index (`ref_id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// DBDropJiraProjectTable will drop the JiraProject table
func DBDropJiraProjectTable(ctx context.Context, db DB) error {
	q := "DROP TABLE IF EXISTS `jira_project`"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBDropJiraProjectTableTx will drop the JiraProject table using the provided transaction
func DBDropJiraProjectTableTx(ctx context.Context, tx Tx) error {
	q := "DROP TABLE IF EXISTS `jira_project`"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// CalculateChecksum will calculate a checksum of the SHA1 of all field values
func (t *JiraProject) CalculateChecksum() string {
	return orm.HashStrings(
		orm.ToString(t.ID),
		orm.ToString(t.ProjectID),
		orm.ToString(t.Key),
		orm.ToString(t.AvatarURL),
		orm.ToString(t.CategoryID),
		orm.ToString(t.CustomerID),
		orm.ToString(t.RefID),
	)
}

// DBCreate will create a new JiraProject record in the database
func (t *JiraProject) DBCreate(ctx context.Context, db DB) (sql.Result, error) {
	q := "INSERT INTO `jira_project` (`jira_project`.`id`,`jira_project`.`checksum`,`jira_project`.`project_id`,`jira_project`.`key`,`jira_project`.`avatar_url`,`jira_project`.`category_id`,`jira_project`.`customer_id`,`jira_project`.`ref_id`) VALUES (?,?,?,?,?,?,?,?)"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLString(t.Key),
		orm.ToSQLString(t.AvatarURL),
		orm.ToSQLString(t.CategoryID),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefID),
	)
}

// DBCreateTx will create a new JiraProject record in the database using the provided transaction
func (t *JiraProject) DBCreateTx(ctx context.Context, tx Tx) (sql.Result, error) {
	q := "INSERT INTO `jira_project` (`jira_project`.`id`,`jira_project`.`checksum`,`jira_project`.`project_id`,`jira_project`.`key`,`jira_project`.`avatar_url`,`jira_project`.`category_id`,`jira_project`.`customer_id`,`jira_project`.`ref_id`) VALUES (?,?,?,?,?,?,?,?)"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLString(t.Key),
		orm.ToSQLString(t.AvatarURL),
		orm.ToSQLString(t.CategoryID),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefID),
	)
}

// DBCreateIgnoreDuplicate will upsert the JiraProject record in the database
func (t *JiraProject) DBCreateIgnoreDuplicate(ctx context.Context, db DB) (sql.Result, error) {
	q := "INSERT INTO `jira_project` (`jira_project`.`id`,`jira_project`.`checksum`,`jira_project`.`project_id`,`jira_project`.`key`,`jira_project`.`avatar_url`,`jira_project`.`category_id`,`jira_project`.`customer_id`,`jira_project`.`ref_id`) VALUES (?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLString(t.Key),
		orm.ToSQLString(t.AvatarURL),
		orm.ToSQLString(t.CategoryID),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefID),
	)
}

// DBCreateIgnoreDuplicateTx will upsert the JiraProject record in the database using the provided transaction
func (t *JiraProject) DBCreateIgnoreDuplicateTx(ctx context.Context, tx Tx) (sql.Result, error) {
	q := "INSERT INTO `jira_project` (`jira_project`.`id`,`jira_project`.`checksum`,`jira_project`.`project_id`,`jira_project`.`key`,`jira_project`.`avatar_url`,`jira_project`.`category_id`,`jira_project`.`customer_id`,`jira_project`.`ref_id`) VALUES (?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLString(t.Key),
		orm.ToSQLString(t.AvatarURL),
		orm.ToSQLString(t.CategoryID),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefID),
	)
}

// DeleteAllJiraProjects deletes all JiraProject records in the database with optional filters
func DeleteAllJiraProjects(ctx context.Context, db DB, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(JiraProjectTableName),
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

// DeleteAllJiraProjectsTx deletes all JiraProject records in the database with optional filters using the provided transaction
func DeleteAllJiraProjectsTx(ctx context.Context, tx Tx, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(JiraProjectTableName),
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

// DBDelete will delete this JiraProject record in the database
func (t *JiraProject) DBDelete(ctx context.Context, db DB) (bool, error) {
	q := "DELETE FROM `jira_project` WHERE `id` = ?"
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

// DBDeleteTx will delete this JiraProject record in the database using the provided transaction
func (t *JiraProject) DBDeleteTx(ctx context.Context, tx Tx) (bool, error) {
	q := "DELETE FROM `jira_project` WHERE `id` = ?"
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

// DBUpdate will update the JiraProject record in the database
func (t *JiraProject) DBUpdate(ctx context.Context, db DB) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `jira_project` SET `checksum`=?,`project_id`=?,`key`=?,`avatar_url`=?,`category_id`=?,`customer_id`=?,`ref_id`=? WHERE `id`=?"
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLString(t.Key),
		orm.ToSQLString(t.AvatarURL),
		orm.ToSQLString(t.CategoryID),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.ID),
	)
}

// DBUpdateTx will update the JiraProject record in the database using the provided transaction
func (t *JiraProject) DBUpdateTx(ctx context.Context, tx Tx) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `jira_project` SET `checksum`=?,`project_id`=?,`key`=?,`avatar_url`=?,`category_id`=?,`customer_id`=?,`ref_id`=? WHERE `id`=?"
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLString(t.Key),
		orm.ToSQLString(t.AvatarURL),
		orm.ToSQLString(t.CategoryID),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.ID),
	)
}

// DBUpsert will upsert the JiraProject record in the database
func (t *JiraProject) DBUpsert(ctx context.Context, db DB, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `jira_project` (`jira_project`.`id`,`jira_project`.`checksum`,`jira_project`.`project_id`,`jira_project`.`key`,`jira_project`.`avatar_url`,`jira_project`.`category_id`,`jira_project`.`customer_id`,`jira_project`.`ref_id`) VALUES (?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `jira_project` (`jira_project`.`id`,`jira_project`.`checksum`,`jira_project`.`project_id`,`jira_project`.`key`,`jira_project`.`avatar_url`,`jira_project`.`category_id`,`jira_project`.`customer_id`,`jira_project`.`ref_id`) VALUES (?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`project_id`=VALUES(`project_id`),`key`=VALUES(`key`),`avatar_url`=VALUES(`avatar_url`),`category_id`=VALUES(`category_id`),`customer_id`=VALUES(`customer_id`),`ref_id`=VALUES(`ref_id`)"
	}
	r, err := db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLString(t.Key),
		orm.ToSQLString(t.AvatarURL),
		orm.ToSQLString(t.CategoryID),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefID),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBUpsertTx will upsert the JiraProject record in the database using the provided transaction
func (t *JiraProject) DBUpsertTx(ctx context.Context, tx Tx, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `jira_project` (`jira_project`.`id`,`jira_project`.`checksum`,`jira_project`.`project_id`,`jira_project`.`key`,`jira_project`.`avatar_url`,`jira_project`.`category_id`,`jira_project`.`customer_id`,`jira_project`.`ref_id`) VALUES (?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `jira_project` (`jira_project`.`id`,`jira_project`.`checksum`,`jira_project`.`project_id`,`jira_project`.`key`,`jira_project`.`avatar_url`,`jira_project`.`category_id`,`jira_project`.`customer_id`,`jira_project`.`ref_id`) VALUES (?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`project_id`=VALUES(`project_id`),`key`=VALUES(`key`),`avatar_url`=VALUES(`avatar_url`),`category_id`=VALUES(`category_id`),`customer_id`=VALUES(`customer_id`),`ref_id`=VALUES(`ref_id`)"
	}
	r, err := tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLString(t.Key),
		orm.ToSQLString(t.AvatarURL),
		orm.ToSQLString(t.CategoryID),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefID),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBFindOne will find a JiraProject record in the database with the primary key
func (t *JiraProject) DBFindOne(ctx context.Context, db DB, value string) (bool, error) {
	q := "SELECT * FROM `jira_project` WHERE `id` = ? LIMIT 1"
	row := db.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _ProjectID sql.NullString
	var _Key sql.NullString
	var _AvatarURL sql.NullString
	var _CategoryID sql.NullString
	var _CustomerID sql.NullString
	var _RefID sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_ProjectID,
		&_Key,
		&_AvatarURL,
		&_CategoryID,
		&_CustomerID,
		&_RefID,
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
	if _ProjectID.Valid {
		t.SetProjectID(_ProjectID.String)
	}
	if _Key.Valid {
		t.SetKey(_Key.String)
	}
	if _AvatarURL.Valid {
		t.SetAvatarURL(_AvatarURL.String)
	}
	if _CategoryID.Valid {
		t.SetCategoryID(_CategoryID.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _RefID.Valid {
		t.SetRefID(_RefID.String)
	}
	return true, nil
}

// DBFindOneTx will find a JiraProject record in the database with the primary key using the provided transaction
func (t *JiraProject) DBFindOneTx(ctx context.Context, tx Tx, value string) (bool, error) {
	q := "SELECT * FROM `jira_project` WHERE `id` = ? LIMIT 1"
	row := tx.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _ProjectID sql.NullString
	var _Key sql.NullString
	var _AvatarURL sql.NullString
	var _CategoryID sql.NullString
	var _CustomerID sql.NullString
	var _RefID sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_ProjectID,
		&_Key,
		&_AvatarURL,
		&_CategoryID,
		&_CustomerID,
		&_RefID,
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
	if _ProjectID.Valid {
		t.SetProjectID(_ProjectID.String)
	}
	if _Key.Valid {
		t.SetKey(_Key.String)
	}
	if _AvatarURL.Valid {
		t.SetAvatarURL(_AvatarURL.String)
	}
	if _CategoryID.Valid {
		t.SetCategoryID(_CategoryID.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _RefID.Valid {
		t.SetRefID(_RefID.String)
	}
	return true, nil
}

// FindJiraProjects will find a JiraProject record in the database with the provided parameters
func FindJiraProjects(ctx context.Context, db DB, _params ...interface{}) ([]*JiraProject, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("project_id"),
		orm.Column("key"),
		orm.Column("avatar_url"),
		orm.Column("category_id"),
		orm.Column("customer_id"),
		orm.Column("ref_id"),
		orm.Table(JiraProjectTableName),
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
	results := make([]*JiraProject, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _ProjectID sql.NullString
		var _Key sql.NullString
		var _AvatarURL sql.NullString
		var _CategoryID sql.NullString
		var _CustomerID sql.NullString
		var _RefID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_ProjectID,
			&_Key,
			&_AvatarURL,
			&_CategoryID,
			&_CustomerID,
			&_RefID,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraProject{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _Key.Valid {
			t.SetKey(_Key.String)
		}
		if _AvatarURL.Valid {
			t.SetAvatarURL(_AvatarURL.String)
		}
		if _CategoryID.Valid {
			t.SetCategoryID(_CategoryID.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindJiraProjectsTx will find a JiraProject record in the database with the provided parameters using the provided transaction
func FindJiraProjectsTx(ctx context.Context, tx Tx, _params ...interface{}) ([]*JiraProject, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("project_id"),
		orm.Column("key"),
		orm.Column("avatar_url"),
		orm.Column("category_id"),
		orm.Column("customer_id"),
		orm.Column("ref_id"),
		orm.Table(JiraProjectTableName),
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
	results := make([]*JiraProject, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _ProjectID sql.NullString
		var _Key sql.NullString
		var _AvatarURL sql.NullString
		var _CategoryID sql.NullString
		var _CustomerID sql.NullString
		var _RefID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_ProjectID,
			&_Key,
			&_AvatarURL,
			&_CategoryID,
			&_CustomerID,
			&_RefID,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraProject{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _Key.Valid {
			t.SetKey(_Key.String)
		}
		if _AvatarURL.Valid {
			t.SetAvatarURL(_AvatarURL.String)
		}
		if _CategoryID.Valid {
			t.SetCategoryID(_CategoryID.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// DBFind will find a JiraProject record in the database with the provided parameters
func (t *JiraProject) DBFind(ctx context.Context, db DB, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("project_id"),
		orm.Column("key"),
		orm.Column("avatar_url"),
		orm.Column("category_id"),
		orm.Column("customer_id"),
		orm.Column("ref_id"),
		orm.Table(JiraProjectTableName),
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
	var _ProjectID sql.NullString
	var _Key sql.NullString
	var _AvatarURL sql.NullString
	var _CategoryID sql.NullString
	var _CustomerID sql.NullString
	var _RefID sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_ProjectID,
		&_Key,
		&_AvatarURL,
		&_CategoryID,
		&_CustomerID,
		&_RefID,
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
	if _ProjectID.Valid {
		t.SetProjectID(_ProjectID.String)
	}
	if _Key.Valid {
		t.SetKey(_Key.String)
	}
	if _AvatarURL.Valid {
		t.SetAvatarURL(_AvatarURL.String)
	}
	if _CategoryID.Valid {
		t.SetCategoryID(_CategoryID.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _RefID.Valid {
		t.SetRefID(_RefID.String)
	}
	return true, nil
}

// DBFindTx will find a JiraProject record in the database with the provided parameters using the provided transaction
func (t *JiraProject) DBFindTx(ctx context.Context, tx Tx, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("project_id"),
		orm.Column("key"),
		orm.Column("avatar_url"),
		orm.Column("category_id"),
		orm.Column("customer_id"),
		orm.Column("ref_id"),
		orm.Table(JiraProjectTableName),
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
	var _ProjectID sql.NullString
	var _Key sql.NullString
	var _AvatarURL sql.NullString
	var _CategoryID sql.NullString
	var _CustomerID sql.NullString
	var _RefID sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_ProjectID,
		&_Key,
		&_AvatarURL,
		&_CategoryID,
		&_CustomerID,
		&_RefID,
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
	if _ProjectID.Valid {
		t.SetProjectID(_ProjectID.String)
	}
	if _Key.Valid {
		t.SetKey(_Key.String)
	}
	if _AvatarURL.Valid {
		t.SetAvatarURL(_AvatarURL.String)
	}
	if _CategoryID.Valid {
		t.SetCategoryID(_CategoryID.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _RefID.Valid {
		t.SetRefID(_RefID.String)
	}
	return true, nil
}

// CountJiraProjects will find the count of JiraProject records in the database
func CountJiraProjects(ctx context.Context, db DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(JiraProjectTableName),
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

// CountJiraProjectsTx will find the count of JiraProject records in the database using the provided transaction
func CountJiraProjectsTx(ctx context.Context, tx Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(JiraProjectTableName),
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

// DBCount will find the count of JiraProject records in the database
func (t *JiraProject) DBCount(ctx context.Context, db DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(JiraProjectTableName),
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

// DBCountTx will find the count of JiraProject records in the database using the provided transaction
func (t *JiraProject) DBCountTx(ctx context.Context, tx Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(JiraProjectTableName),
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

// DBExists will return true if the JiraProject record exists in the database
func (t *JiraProject) DBExists(ctx context.Context, db DB) (bool, error) {
	q := "SELECT * FROM `jira_project` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := db.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// DBExistsTx will return true if the JiraProject record exists in the database using the provided transaction
func (t *JiraProject) DBExistsTx(ctx context.Context, tx Tx) (bool, error) {
	q := "SELECT * FROM `jira_project` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := tx.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// PrimaryKeyColumn returns the column name for the primary key
func (t *JiraProject) PrimaryKeyColumn() string {
	return JiraProjectColumnID
}

// PrimaryKeyColumnType returns the primary key column Go type as a string
func (t *JiraProject) PrimaryKeyColumnType() string {
	return "string"
}

// PrimaryKey returns the primary key column value
func (t *JiraProject) PrimaryKey() interface{} {
	return t.ID
}
