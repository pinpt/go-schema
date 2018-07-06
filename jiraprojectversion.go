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

var _ Model = (*JiraProjectVersion)(nil)
var _ CSVWriter = (*JiraProjectVersion)(nil)
var _ JSONWriter = (*JiraProjectVersion)(nil)
var _ Checksum = (*JiraProjectVersion)(nil)

// JiraProjectVersionTableName is the name of the table in SQL
const JiraProjectVersionTableName = "jira_project_version"

var JiraProjectVersionColumns = []string{
	"id",
	"checksum",
	"version_id",
	"project_id",
	"name",
	"description",
	"archived",
	"released",
	"release_date",
	"overdue",
	"user_release_date",
	"customer_id",
}

// JiraProjectVersion table
type JiraProjectVersion struct {
	Archived        bool    `json:"archived"`
	Checksum        *string `json:"checksum,omitempty"`
	CustomerID      string  `json:"customer_id"`
	Description     *string `json:"description,omitempty"`
	ID              string  `json:"id"`
	Name            string  `json:"name"`
	Overdue         bool    `json:"overdue"`
	ProjectID       string  `json:"project_id"`
	ReleaseDate     *int64  `json:"release_date,omitempty"`
	Released        bool    `json:"released"`
	UserReleaseDate *int64  `json:"user_release_date,omitempty"`
	VersionID       string  `json:"version_id"`
}

// TableName returns the SQL table name for JiraProjectVersion and satifies the Model interface
func (t *JiraProjectVersion) TableName() string {
	return JiraProjectVersionTableName
}

// ToCSV will serialize the JiraProjectVersion instance to a CSV compatible array of strings
func (t *JiraProjectVersion) ToCSV() []string {
	return []string{
		t.ID,
		t.CalculateChecksum(),
		t.VersionID,
		t.ProjectID,
		t.Name,
		toCSVString(t.Description),
		toCSVBool(t.Archived),
		toCSVBool(t.Released),
		toCSVString(t.ReleaseDate),
		toCSVBool(t.Overdue),
		toCSVString(t.UserReleaseDate),
		t.CustomerID,
	}
}

// WriteCSV will serialize the JiraProjectVersion instance to the writer as CSV and satisfies the CSVWriter interface
func (t *JiraProjectVersion) WriteCSV(w *csv.Writer) error {
	return w.Write(t.ToCSV())
}

// WriteJSON will serialize the JiraProjectVersion instance to the writer as JSON and satisfies the JSONWriter interface
func (t *JiraProjectVersion) WriteJSON(w io.Writer, indent ...bool) error {
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

// NewJiraProjectVersionReader creates a JSON reader which can read in JiraProjectVersion objects serialized as JSON either as an array, single object or json new lines
// and writes each JiraProjectVersion to the channel provided
func NewJiraProjectVersionReader(r io.Reader, ch chan<- JiraProjectVersion) error {
	return orm.Deserialize(r, func(buf json.RawMessage) error {
		dec := json.NewDecoder(bytes.NewBuffer(buf))
		e := JiraProjectVersion{}
		if err := dec.Decode(&e); err != nil {
			return err
		}
		ch <- e
		return nil
	})
}

// NewCSVJiraProjectVersionReaderDir will read the reader as CSV and emit each record to the channel provided
func NewCSVJiraProjectVersionReader(r io.Reader, ch chan<- JiraProjectVersion) error {
	cr := csv.NewReader(r)
	for {
		record, err := cr.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		ch <- JiraProjectVersion{
			ID:              record[0],
			Checksum:        fromStringPointer(record[1]),
			VersionID:       record[2],
			ProjectID:       record[3],
			Name:            record[4],
			Description:     fromStringPointer(record[5]),
			Archived:        fromCSVBool(record[6]),
			Released:        fromCSVBool(record[7]),
			ReleaseDate:     fromCSVInt64Pointer(record[8]),
			Overdue:         fromCSVBool(record[9]),
			UserReleaseDate: fromCSVInt64Pointer(record[10]),
			CustomerID:      record[11],
		}
	}
	return nil
}

// NewCSVJiraProjectVersionReaderFile will read the file as a CSV and emit each record to the channel provided
func NewCSVJiraProjectVersionReaderFile(fp string, ch chan<- JiraProjectVersion) error {
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
	return NewCSVJiraProjectVersionReader(fc, ch)
}

// NewCSVJiraProjectVersionReaderDir will read the jira_project_version.csv.gz file as a CSV and emit each record to the channel provided
func NewCSVJiraProjectVersionReaderDir(dir string, ch chan<- JiraProjectVersion) error {
	return NewCSVJiraProjectVersionReaderFile(filepath.Join(dir, "jira_project_version.csv.gz"), ch)
}

// JiraProjectVersionCSVDeduper is a function callback which takes the existing value (a) and the new value (b)
// and the return value should be the one to use (or a new one, if applicable). return nil
// to skip processing of this record
type JiraProjectVersionCSVDeduper func(a JiraProjectVersion, b JiraProjectVersion) *JiraProjectVersion

// JiraProjectVersionCSVDedupeDisabled is set on whether the CSV writer should de-dupe values by key
var JiraProjectVersionCSVDedupeDisabled bool

// NewJiraProjectVersionCSVWriterSize creates a batch writer that will write each JiraProjectVersion into a CSV file
// this method will automatically de-duplicate entries using the primary key. if the checksum
// for a newer item with the same primary key doesn't match a previously sent item, the newer
// one will be used
func NewJiraProjectVersionCSVWriterSize(w io.Writer, size int, dedupers ...JiraProjectVersionCSVDeduper) (chan JiraProjectVersion, chan bool, error) {
	cw := csv.NewWriter(w)
	ch := make(chan JiraProjectVersion, size)
	done := make(chan bool)
	go func() {
		defer func() { done <- true }()
		dodedupe := !JiraProjectVersionCSVDedupeDisabled
		var kv map[string]*JiraProjectVersion
		var deduper JiraProjectVersionCSVDeduper
		if dedupers != nil && len(dedupers) > 0 {
			deduper = dedupers[0]
			dodedupe = true
		}
		if dodedupe {
			kv = make(map[string]*JiraProjectVersion)
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

// JiraProjectVersionCSVDefaultSize is the default channel buffer size if not provided
var JiraProjectVersionCSVDefaultSize = 100

// NewJiraProjectVersionCSVWriter creates a batch writer that will write each JiraProjectVersion into a CSV file
func NewJiraProjectVersionCSVWriter(w io.Writer, dedupers ...JiraProjectVersionCSVDeduper) (chan JiraProjectVersion, chan bool, error) {
	return NewJiraProjectVersionCSVWriterSize(w, JiraProjectVersionCSVDefaultSize, dedupers...)
}

// NewJiraProjectVersionCSVWriterDir creates a batch writer that will write each JiraProjectVersion into a CSV file named jira_project_version.csv.gz in dir
func NewJiraProjectVersionCSVWriterDir(dir string, dedupers ...JiraProjectVersionCSVDeduper) (chan JiraProjectVersion, chan bool, error) {
	return NewJiraProjectVersionCSVWriterFile(filepath.Join(dir, "jira_project_version.csv.gz"), dedupers...)
}

// NewJiraProjectVersionCSVWriterFile creates a batch writer that will write each JiraProjectVersion into a CSV file
func NewJiraProjectVersionCSVWriterFile(fn string, dedupers ...JiraProjectVersionCSVDeduper) (chan JiraProjectVersion, chan bool, error) {
	f, err := os.Create(fn)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening CSV file %s. %v", fn, err)
	}
	var fc io.WriteCloser = f
	if filepath.Ext(fn) == ".gz" {
		w, _ := gzip.NewWriterLevel(f, gzip.BestCompression)
		fc = w
	}
	ch, done, err := NewJiraProjectVersionCSVWriter(fc, dedupers...)
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

type JiraProjectVersionDBAction func(ctx context.Context, db *sql.DB, record JiraProjectVersion) error

// NewJiraProjectVersionDBWriterSize creates a DB writer that will write each issue into the DB
func NewJiraProjectVersionDBWriterSize(ctx context.Context, db *sql.DB, errors chan<- error, size int, actions ...JiraProjectVersionDBAction) (chan JiraProjectVersion, chan bool, error) {
	ch := make(chan JiraProjectVersion, size)
	done := make(chan bool)
	var action JiraProjectVersionDBAction
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

// NewJiraProjectVersionDBWriter creates a DB writer that will write each issue into the DB
func NewJiraProjectVersionDBWriter(ctx context.Context, db *sql.DB, errors chan<- error, actions ...JiraProjectVersionDBAction) (chan JiraProjectVersion, chan bool, error) {
	return NewJiraProjectVersionDBWriterSize(ctx, db, errors, 100, actions...)
}

// JiraProjectVersionColumnID is the ID SQL column name for the JiraProjectVersion table
const JiraProjectVersionColumnID = "id"

// JiraProjectVersionEscapedColumnID is the escaped ID SQL column name for the JiraProjectVersion table
const JiraProjectVersionEscapedColumnID = "`id`"

// JiraProjectVersionColumnChecksum is the Checksum SQL column name for the JiraProjectVersion table
const JiraProjectVersionColumnChecksum = "checksum"

// JiraProjectVersionEscapedColumnChecksum is the escaped Checksum SQL column name for the JiraProjectVersion table
const JiraProjectVersionEscapedColumnChecksum = "`checksum`"

// JiraProjectVersionColumnVersionID is the VersionID SQL column name for the JiraProjectVersion table
const JiraProjectVersionColumnVersionID = "version_id"

// JiraProjectVersionEscapedColumnVersionID is the escaped VersionID SQL column name for the JiraProjectVersion table
const JiraProjectVersionEscapedColumnVersionID = "`version_id`"

// JiraProjectVersionColumnProjectID is the ProjectID SQL column name for the JiraProjectVersion table
const JiraProjectVersionColumnProjectID = "project_id"

// JiraProjectVersionEscapedColumnProjectID is the escaped ProjectID SQL column name for the JiraProjectVersion table
const JiraProjectVersionEscapedColumnProjectID = "`project_id`"

// JiraProjectVersionColumnName is the Name SQL column name for the JiraProjectVersion table
const JiraProjectVersionColumnName = "name"

// JiraProjectVersionEscapedColumnName is the escaped Name SQL column name for the JiraProjectVersion table
const JiraProjectVersionEscapedColumnName = "`name`"

// JiraProjectVersionColumnDescription is the Description SQL column name for the JiraProjectVersion table
const JiraProjectVersionColumnDescription = "description"

// JiraProjectVersionEscapedColumnDescription is the escaped Description SQL column name for the JiraProjectVersion table
const JiraProjectVersionEscapedColumnDescription = "`description`"

// JiraProjectVersionColumnArchived is the Archived SQL column name for the JiraProjectVersion table
const JiraProjectVersionColumnArchived = "archived"

// JiraProjectVersionEscapedColumnArchived is the escaped Archived SQL column name for the JiraProjectVersion table
const JiraProjectVersionEscapedColumnArchived = "`archived`"

// JiraProjectVersionColumnReleased is the Released SQL column name for the JiraProjectVersion table
const JiraProjectVersionColumnReleased = "released"

// JiraProjectVersionEscapedColumnReleased is the escaped Released SQL column name for the JiraProjectVersion table
const JiraProjectVersionEscapedColumnReleased = "`released`"

// JiraProjectVersionColumnReleaseDate is the ReleaseDate SQL column name for the JiraProjectVersion table
const JiraProjectVersionColumnReleaseDate = "release_date"

// JiraProjectVersionEscapedColumnReleaseDate is the escaped ReleaseDate SQL column name for the JiraProjectVersion table
const JiraProjectVersionEscapedColumnReleaseDate = "`release_date`"

// JiraProjectVersionColumnOverdue is the Overdue SQL column name for the JiraProjectVersion table
const JiraProjectVersionColumnOverdue = "overdue"

// JiraProjectVersionEscapedColumnOverdue is the escaped Overdue SQL column name for the JiraProjectVersion table
const JiraProjectVersionEscapedColumnOverdue = "`overdue`"

// JiraProjectVersionColumnUserReleaseDate is the UserReleaseDate SQL column name for the JiraProjectVersion table
const JiraProjectVersionColumnUserReleaseDate = "user_release_date"

// JiraProjectVersionEscapedColumnUserReleaseDate is the escaped UserReleaseDate SQL column name for the JiraProjectVersion table
const JiraProjectVersionEscapedColumnUserReleaseDate = "`user_release_date`"

// JiraProjectVersionColumnCustomerID is the CustomerID SQL column name for the JiraProjectVersion table
const JiraProjectVersionColumnCustomerID = "customer_id"

// JiraProjectVersionEscapedColumnCustomerID is the escaped CustomerID SQL column name for the JiraProjectVersion table
const JiraProjectVersionEscapedColumnCustomerID = "`customer_id`"

// GetID will return the JiraProjectVersion ID value
func (t *JiraProjectVersion) GetID() string {
	return t.ID
}

// SetID will set the JiraProjectVersion ID value
func (t *JiraProjectVersion) SetID(v string) {
	t.ID = v
}

// FindJiraProjectVersionByID will find a JiraProjectVersion by ID
func FindJiraProjectVersionByID(ctx context.Context, db *sql.DB, value string) (*JiraProjectVersion, error) {
	q := "SELECT `jira_project_version`.`id`,`jira_project_version`.`checksum`,`jira_project_version`.`version_id`,`jira_project_version`.`project_id`,`jira_project_version`.`name`,`jira_project_version`.`description`,`jira_project_version`.`archived`,`jira_project_version`.`released`,`jira_project_version`.`release_date`,`jira_project_version`.`overdue`,`jira_project_version`.`user_release_date`,`jira_project_version`.`customer_id` FROM `jira_project_version` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _VersionID sql.NullString
	var _ProjectID sql.NullString
	var _Name sql.NullString
	var _Description sql.NullString
	var _Archived sql.NullBool
	var _Released sql.NullBool
	var _ReleaseDate sql.NullInt64
	var _Overdue sql.NullBool
	var _UserReleaseDate sql.NullInt64
	var _CustomerID sql.NullString
	err := db.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_VersionID,
		&_ProjectID,
		&_Name,
		&_Description,
		&_Archived,
		&_Released,
		&_ReleaseDate,
		&_Overdue,
		&_UserReleaseDate,
		&_CustomerID,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &JiraProjectVersion{}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _Checksum.Valid {
		t.SetChecksum(_Checksum.String)
	}
	if _VersionID.Valid {
		t.SetVersionID(_VersionID.String)
	}
	if _ProjectID.Valid {
		t.SetProjectID(_ProjectID.String)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _Description.Valid {
		t.SetDescription(_Description.String)
	}
	if _Archived.Valid {
		t.SetArchived(_Archived.Bool)
	}
	if _Released.Valid {
		t.SetReleased(_Released.Bool)
	}
	if _ReleaseDate.Valid {
		t.SetReleaseDate(_ReleaseDate.Int64)
	}
	if _Overdue.Valid {
		t.SetOverdue(_Overdue.Bool)
	}
	if _UserReleaseDate.Valid {
		t.SetUserReleaseDate(_UserReleaseDate.Int64)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	return t, nil
}

// FindJiraProjectVersionByIDTx will find a JiraProjectVersion by ID using the provided transaction
func FindJiraProjectVersionByIDTx(ctx context.Context, tx *sql.Tx, value string) (*JiraProjectVersion, error) {
	q := "SELECT `jira_project_version`.`id`,`jira_project_version`.`checksum`,`jira_project_version`.`version_id`,`jira_project_version`.`project_id`,`jira_project_version`.`name`,`jira_project_version`.`description`,`jira_project_version`.`archived`,`jira_project_version`.`released`,`jira_project_version`.`release_date`,`jira_project_version`.`overdue`,`jira_project_version`.`user_release_date`,`jira_project_version`.`customer_id` FROM `jira_project_version` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _VersionID sql.NullString
	var _ProjectID sql.NullString
	var _Name sql.NullString
	var _Description sql.NullString
	var _Archived sql.NullBool
	var _Released sql.NullBool
	var _ReleaseDate sql.NullInt64
	var _Overdue sql.NullBool
	var _UserReleaseDate sql.NullInt64
	var _CustomerID sql.NullString
	err := tx.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_VersionID,
		&_ProjectID,
		&_Name,
		&_Description,
		&_Archived,
		&_Released,
		&_ReleaseDate,
		&_Overdue,
		&_UserReleaseDate,
		&_CustomerID,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &JiraProjectVersion{}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _Checksum.Valid {
		t.SetChecksum(_Checksum.String)
	}
	if _VersionID.Valid {
		t.SetVersionID(_VersionID.String)
	}
	if _ProjectID.Valid {
		t.SetProjectID(_ProjectID.String)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _Description.Valid {
		t.SetDescription(_Description.String)
	}
	if _Archived.Valid {
		t.SetArchived(_Archived.Bool)
	}
	if _Released.Valid {
		t.SetReleased(_Released.Bool)
	}
	if _ReleaseDate.Valid {
		t.SetReleaseDate(_ReleaseDate.Int64)
	}
	if _Overdue.Valid {
		t.SetOverdue(_Overdue.Bool)
	}
	if _UserReleaseDate.Valid {
		t.SetUserReleaseDate(_UserReleaseDate.Int64)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	return t, nil
}

// GetChecksum will return the JiraProjectVersion Checksum value
func (t *JiraProjectVersion) GetChecksum() string {
	if t.Checksum == nil {
		return ""
	}
	return *t.Checksum
}

// SetChecksum will set the JiraProjectVersion Checksum value
func (t *JiraProjectVersion) SetChecksum(v string) {
	t.Checksum = &v
}

// GetVersionID will return the JiraProjectVersion VersionID value
func (t *JiraProjectVersion) GetVersionID() string {
	return t.VersionID
}

// SetVersionID will set the JiraProjectVersion VersionID value
func (t *JiraProjectVersion) SetVersionID(v string) {
	t.VersionID = v
}

// FindJiraProjectVersionsByVersionID will find all JiraProjectVersions by the VersionID value
func FindJiraProjectVersionsByVersionID(ctx context.Context, db *sql.DB, value string) ([]*JiraProjectVersion, error) {
	q := "SELECT `jira_project_version`.`id`,`jira_project_version`.`checksum`,`jira_project_version`.`version_id`,`jira_project_version`.`project_id`,`jira_project_version`.`name`,`jira_project_version`.`description`,`jira_project_version`.`archived`,`jira_project_version`.`released`,`jira_project_version`.`release_date`,`jira_project_version`.`overdue`,`jira_project_version`.`user_release_date`,`jira_project_version`.`customer_id` FROM `jira_project_version` WHERE `version_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraProjectVersion, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _VersionID sql.NullString
		var _ProjectID sql.NullString
		var _Name sql.NullString
		var _Description sql.NullString
		var _Archived sql.NullBool
		var _Released sql.NullBool
		var _ReleaseDate sql.NullInt64
		var _Overdue sql.NullBool
		var _UserReleaseDate sql.NullInt64
		var _CustomerID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_VersionID,
			&_ProjectID,
			&_Name,
			&_Description,
			&_Archived,
			&_Released,
			&_ReleaseDate,
			&_Overdue,
			&_UserReleaseDate,
			&_CustomerID,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraProjectVersion{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _VersionID.Valid {
			t.SetVersionID(_VersionID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _Description.Valid {
			t.SetDescription(_Description.String)
		}
		if _Archived.Valid {
			t.SetArchived(_Archived.Bool)
		}
		if _Released.Valid {
			t.SetReleased(_Released.Bool)
		}
		if _ReleaseDate.Valid {
			t.SetReleaseDate(_ReleaseDate.Int64)
		}
		if _Overdue.Valid {
			t.SetOverdue(_Overdue.Bool)
		}
		if _UserReleaseDate.Valid {
			t.SetUserReleaseDate(_UserReleaseDate.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindJiraProjectVersionsByVersionIDTx will find all JiraProjectVersions by the VersionID value using the provided transaction
func FindJiraProjectVersionsByVersionIDTx(ctx context.Context, tx *sql.Tx, value string) ([]*JiraProjectVersion, error) {
	q := "SELECT `jira_project_version`.`id`,`jira_project_version`.`checksum`,`jira_project_version`.`version_id`,`jira_project_version`.`project_id`,`jira_project_version`.`name`,`jira_project_version`.`description`,`jira_project_version`.`archived`,`jira_project_version`.`released`,`jira_project_version`.`release_date`,`jira_project_version`.`overdue`,`jira_project_version`.`user_release_date`,`jira_project_version`.`customer_id` FROM `jira_project_version` WHERE `version_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraProjectVersion, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _VersionID sql.NullString
		var _ProjectID sql.NullString
		var _Name sql.NullString
		var _Description sql.NullString
		var _Archived sql.NullBool
		var _Released sql.NullBool
		var _ReleaseDate sql.NullInt64
		var _Overdue sql.NullBool
		var _UserReleaseDate sql.NullInt64
		var _CustomerID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_VersionID,
			&_ProjectID,
			&_Name,
			&_Description,
			&_Archived,
			&_Released,
			&_ReleaseDate,
			&_Overdue,
			&_UserReleaseDate,
			&_CustomerID,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraProjectVersion{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _VersionID.Valid {
			t.SetVersionID(_VersionID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _Description.Valid {
			t.SetDescription(_Description.String)
		}
		if _Archived.Valid {
			t.SetArchived(_Archived.Bool)
		}
		if _Released.Valid {
			t.SetReleased(_Released.Bool)
		}
		if _ReleaseDate.Valid {
			t.SetReleaseDate(_ReleaseDate.Int64)
		}
		if _Overdue.Valid {
			t.SetOverdue(_Overdue.Bool)
		}
		if _UserReleaseDate.Valid {
			t.SetUserReleaseDate(_UserReleaseDate.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetProjectID will return the JiraProjectVersion ProjectID value
func (t *JiraProjectVersion) GetProjectID() string {
	return t.ProjectID
}

// SetProjectID will set the JiraProjectVersion ProjectID value
func (t *JiraProjectVersion) SetProjectID(v string) {
	t.ProjectID = v
}

// FindJiraProjectVersionsByProjectID will find all JiraProjectVersions by the ProjectID value
func FindJiraProjectVersionsByProjectID(ctx context.Context, db *sql.DB, value string) ([]*JiraProjectVersion, error) {
	q := "SELECT `jira_project_version`.`id`,`jira_project_version`.`checksum`,`jira_project_version`.`version_id`,`jira_project_version`.`project_id`,`jira_project_version`.`name`,`jira_project_version`.`description`,`jira_project_version`.`archived`,`jira_project_version`.`released`,`jira_project_version`.`release_date`,`jira_project_version`.`overdue`,`jira_project_version`.`user_release_date`,`jira_project_version`.`customer_id` FROM `jira_project_version` WHERE `project_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraProjectVersion, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _VersionID sql.NullString
		var _ProjectID sql.NullString
		var _Name sql.NullString
		var _Description sql.NullString
		var _Archived sql.NullBool
		var _Released sql.NullBool
		var _ReleaseDate sql.NullInt64
		var _Overdue sql.NullBool
		var _UserReleaseDate sql.NullInt64
		var _CustomerID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_VersionID,
			&_ProjectID,
			&_Name,
			&_Description,
			&_Archived,
			&_Released,
			&_ReleaseDate,
			&_Overdue,
			&_UserReleaseDate,
			&_CustomerID,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraProjectVersion{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _VersionID.Valid {
			t.SetVersionID(_VersionID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _Description.Valid {
			t.SetDescription(_Description.String)
		}
		if _Archived.Valid {
			t.SetArchived(_Archived.Bool)
		}
		if _Released.Valid {
			t.SetReleased(_Released.Bool)
		}
		if _ReleaseDate.Valid {
			t.SetReleaseDate(_ReleaseDate.Int64)
		}
		if _Overdue.Valid {
			t.SetOverdue(_Overdue.Bool)
		}
		if _UserReleaseDate.Valid {
			t.SetUserReleaseDate(_UserReleaseDate.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindJiraProjectVersionsByProjectIDTx will find all JiraProjectVersions by the ProjectID value using the provided transaction
func FindJiraProjectVersionsByProjectIDTx(ctx context.Context, tx *sql.Tx, value string) ([]*JiraProjectVersion, error) {
	q := "SELECT `jira_project_version`.`id`,`jira_project_version`.`checksum`,`jira_project_version`.`version_id`,`jira_project_version`.`project_id`,`jira_project_version`.`name`,`jira_project_version`.`description`,`jira_project_version`.`archived`,`jira_project_version`.`released`,`jira_project_version`.`release_date`,`jira_project_version`.`overdue`,`jira_project_version`.`user_release_date`,`jira_project_version`.`customer_id` FROM `jira_project_version` WHERE `project_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraProjectVersion, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _VersionID sql.NullString
		var _ProjectID sql.NullString
		var _Name sql.NullString
		var _Description sql.NullString
		var _Archived sql.NullBool
		var _Released sql.NullBool
		var _ReleaseDate sql.NullInt64
		var _Overdue sql.NullBool
		var _UserReleaseDate sql.NullInt64
		var _CustomerID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_VersionID,
			&_ProjectID,
			&_Name,
			&_Description,
			&_Archived,
			&_Released,
			&_ReleaseDate,
			&_Overdue,
			&_UserReleaseDate,
			&_CustomerID,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraProjectVersion{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _VersionID.Valid {
			t.SetVersionID(_VersionID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _Description.Valid {
			t.SetDescription(_Description.String)
		}
		if _Archived.Valid {
			t.SetArchived(_Archived.Bool)
		}
		if _Released.Valid {
			t.SetReleased(_Released.Bool)
		}
		if _ReleaseDate.Valid {
			t.SetReleaseDate(_ReleaseDate.Int64)
		}
		if _Overdue.Valid {
			t.SetOverdue(_Overdue.Bool)
		}
		if _UserReleaseDate.Valid {
			t.SetUserReleaseDate(_UserReleaseDate.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetName will return the JiraProjectVersion Name value
func (t *JiraProjectVersion) GetName() string {
	return t.Name
}

// SetName will set the JiraProjectVersion Name value
func (t *JiraProjectVersion) SetName(v string) {
	t.Name = v
}

// GetDescription will return the JiraProjectVersion Description value
func (t *JiraProjectVersion) GetDescription() string {
	if t.Description == nil {
		return ""
	}
	return *t.Description
}

// SetDescription will set the JiraProjectVersion Description value
func (t *JiraProjectVersion) SetDescription(v string) {
	t.Description = &v
}

// GetArchived will return the JiraProjectVersion Archived value
func (t *JiraProjectVersion) GetArchived() bool {
	return t.Archived
}

// SetArchived will set the JiraProjectVersion Archived value
func (t *JiraProjectVersion) SetArchived(v bool) {
	t.Archived = v
}

// GetReleased will return the JiraProjectVersion Released value
func (t *JiraProjectVersion) GetReleased() bool {
	return t.Released
}

// SetReleased will set the JiraProjectVersion Released value
func (t *JiraProjectVersion) SetReleased(v bool) {
	t.Released = v
}

// GetReleaseDate will return the JiraProjectVersion ReleaseDate value
func (t *JiraProjectVersion) GetReleaseDate() int64 {
	if t.ReleaseDate == nil {
		return int64(0)
	}
	return *t.ReleaseDate
}

// SetReleaseDate will set the JiraProjectVersion ReleaseDate value
func (t *JiraProjectVersion) SetReleaseDate(v int64) {
	t.ReleaseDate = &v
}

// GetOverdue will return the JiraProjectVersion Overdue value
func (t *JiraProjectVersion) GetOverdue() bool {
	return t.Overdue
}

// SetOverdue will set the JiraProjectVersion Overdue value
func (t *JiraProjectVersion) SetOverdue(v bool) {
	t.Overdue = v
}

// GetUserReleaseDate will return the JiraProjectVersion UserReleaseDate value
func (t *JiraProjectVersion) GetUserReleaseDate() int64 {
	if t.UserReleaseDate == nil {
		return int64(0)
	}
	return *t.UserReleaseDate
}

// SetUserReleaseDate will set the JiraProjectVersion UserReleaseDate value
func (t *JiraProjectVersion) SetUserReleaseDate(v int64) {
	t.UserReleaseDate = &v
}

// GetCustomerID will return the JiraProjectVersion CustomerID value
func (t *JiraProjectVersion) GetCustomerID() string {
	return t.CustomerID
}

// SetCustomerID will set the JiraProjectVersion CustomerID value
func (t *JiraProjectVersion) SetCustomerID(v string) {
	t.CustomerID = v
}

// FindJiraProjectVersionsByCustomerID will find all JiraProjectVersions by the CustomerID value
func FindJiraProjectVersionsByCustomerID(ctx context.Context, db *sql.DB, value string) ([]*JiraProjectVersion, error) {
	q := "SELECT `jira_project_version`.`id`,`jira_project_version`.`checksum`,`jira_project_version`.`version_id`,`jira_project_version`.`project_id`,`jira_project_version`.`name`,`jira_project_version`.`description`,`jira_project_version`.`archived`,`jira_project_version`.`released`,`jira_project_version`.`release_date`,`jira_project_version`.`overdue`,`jira_project_version`.`user_release_date`,`jira_project_version`.`customer_id` FROM `jira_project_version` WHERE `customer_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraProjectVersion, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _VersionID sql.NullString
		var _ProjectID sql.NullString
		var _Name sql.NullString
		var _Description sql.NullString
		var _Archived sql.NullBool
		var _Released sql.NullBool
		var _ReleaseDate sql.NullInt64
		var _Overdue sql.NullBool
		var _UserReleaseDate sql.NullInt64
		var _CustomerID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_VersionID,
			&_ProjectID,
			&_Name,
			&_Description,
			&_Archived,
			&_Released,
			&_ReleaseDate,
			&_Overdue,
			&_UserReleaseDate,
			&_CustomerID,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraProjectVersion{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _VersionID.Valid {
			t.SetVersionID(_VersionID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _Description.Valid {
			t.SetDescription(_Description.String)
		}
		if _Archived.Valid {
			t.SetArchived(_Archived.Bool)
		}
		if _Released.Valid {
			t.SetReleased(_Released.Bool)
		}
		if _ReleaseDate.Valid {
			t.SetReleaseDate(_ReleaseDate.Int64)
		}
		if _Overdue.Valid {
			t.SetOverdue(_Overdue.Bool)
		}
		if _UserReleaseDate.Valid {
			t.SetUserReleaseDate(_UserReleaseDate.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindJiraProjectVersionsByCustomerIDTx will find all JiraProjectVersions by the CustomerID value using the provided transaction
func FindJiraProjectVersionsByCustomerIDTx(ctx context.Context, tx *sql.Tx, value string) ([]*JiraProjectVersion, error) {
	q := "SELECT `jira_project_version`.`id`,`jira_project_version`.`checksum`,`jira_project_version`.`version_id`,`jira_project_version`.`project_id`,`jira_project_version`.`name`,`jira_project_version`.`description`,`jira_project_version`.`archived`,`jira_project_version`.`released`,`jira_project_version`.`release_date`,`jira_project_version`.`overdue`,`jira_project_version`.`user_release_date`,`jira_project_version`.`customer_id` FROM `jira_project_version` WHERE `customer_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraProjectVersion, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _VersionID sql.NullString
		var _ProjectID sql.NullString
		var _Name sql.NullString
		var _Description sql.NullString
		var _Archived sql.NullBool
		var _Released sql.NullBool
		var _ReleaseDate sql.NullInt64
		var _Overdue sql.NullBool
		var _UserReleaseDate sql.NullInt64
		var _CustomerID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_VersionID,
			&_ProjectID,
			&_Name,
			&_Description,
			&_Archived,
			&_Released,
			&_ReleaseDate,
			&_Overdue,
			&_UserReleaseDate,
			&_CustomerID,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraProjectVersion{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _VersionID.Valid {
			t.SetVersionID(_VersionID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _Description.Valid {
			t.SetDescription(_Description.String)
		}
		if _Archived.Valid {
			t.SetArchived(_Archived.Bool)
		}
		if _Released.Valid {
			t.SetReleased(_Released.Bool)
		}
		if _ReleaseDate.Valid {
			t.SetReleaseDate(_ReleaseDate.Int64)
		}
		if _Overdue.Valid {
			t.SetOverdue(_Overdue.Bool)
		}
		if _UserReleaseDate.Valid {
			t.SetUserReleaseDate(_UserReleaseDate.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

func (t *JiraProjectVersion) toTimestamp(value time.Time) *timestamp.Timestamp {
	ts, _ := ptypes.TimestampProto(value)
	return ts
}

// DBCreateJiraProjectVersionTable will create the JiraProjectVersion table
func DBCreateJiraProjectVersionTable(ctx context.Context, db *sql.DB) error {
	q := "CREATE TABLE `jira_project_version` (`id` VARCHAR(64) NOT NULL PRIMARY KEY,`checksum` CHAR(64),`version_id` VARCHAR(255) NOT NULL,`project_id` VARCHAR(64) NOT NULL,`name` TEXT NOT NULL,`description` TEXT,`archived` BOOL NOT NULL,`released` BOOL NOT NULL,`release_date`BIGINT(20) UNSIGNED,`overdue` BOOL NOT NULL,`user_release_date` BIGINT(20) UNSIGNED,`customer_id` VARCHAR(64) NOT NULL,INDEX jira_project_version_version_id_index (`version_id`),INDEX jira_project_version_project_id_index (`project_id`),INDEX jira_project_version_customer_id_index (`customer_id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBCreateJiraProjectVersionTableTx will create the JiraProjectVersion table using the provided transction
func DBCreateJiraProjectVersionTableTx(ctx context.Context, tx *sql.Tx) error {
	q := "CREATE TABLE `jira_project_version` (`id` VARCHAR(64) NOT NULL PRIMARY KEY,`checksum` CHAR(64),`version_id` VARCHAR(255) NOT NULL,`project_id` VARCHAR(64) NOT NULL,`name` TEXT NOT NULL,`description` TEXT,`archived` BOOL NOT NULL,`released` BOOL NOT NULL,`release_date`BIGINT(20) UNSIGNED,`overdue` BOOL NOT NULL,`user_release_date` BIGINT(20) UNSIGNED,`customer_id` VARCHAR(64) NOT NULL,INDEX jira_project_version_version_id_index (`version_id`),INDEX jira_project_version_project_id_index (`project_id`),INDEX jira_project_version_customer_id_index (`customer_id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// DBDropJiraProjectVersionTable will drop the JiraProjectVersion table
func DBDropJiraProjectVersionTable(ctx context.Context, db *sql.DB) error {
	q := "DROP TABLE IF EXISTS `jira_project_version`"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBDropJiraProjectVersionTableTx will drop the JiraProjectVersion table using the provided transaction
func DBDropJiraProjectVersionTableTx(ctx context.Context, tx *sql.Tx) error {
	q := "DROP TABLE IF EXISTS `jira_project_version`"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// CalculateChecksum will calculate a checksum of the SHA1 of all field values
func (t *JiraProjectVersion) CalculateChecksum() string {
	return orm.HashStrings(
		orm.ToString(t.ID),
		orm.ToString(t.VersionID),
		orm.ToString(t.ProjectID),
		orm.ToString(t.Name),
		orm.ToString(t.Description),
		orm.ToString(t.Archived),
		orm.ToString(t.Released),
		orm.ToString(t.ReleaseDate),
		orm.ToString(t.Overdue),
		orm.ToString(t.UserReleaseDate),
		orm.ToString(t.CustomerID),
	)
}

// DBCreate will create a new JiraProjectVersion record in the database
func (t *JiraProjectVersion) DBCreate(ctx context.Context, db *sql.DB) (sql.Result, error) {
	q := "INSERT INTO `jira_project_version` (`jira_project_version`.`id`,`jira_project_version`.`checksum`,`jira_project_version`.`version_id`,`jira_project_version`.`project_id`,`jira_project_version`.`name`,`jira_project_version`.`description`,`jira_project_version`.`archived`,`jira_project_version`.`released`,`jira_project_version`.`release_date`,`jira_project_version`.`overdue`,`jira_project_version`.`user_release_date`,`jira_project_version`.`customer_id`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?)"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.VersionID),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.Description),
		orm.ToSQLBool(t.Archived),
		orm.ToSQLBool(t.Released),
		orm.ToSQLInt64(t.ReleaseDate),
		orm.ToSQLBool(t.Overdue),
		orm.ToSQLInt64(t.UserReleaseDate),
		orm.ToSQLString(t.CustomerID),
	)
}

// DBCreateTx will create a new JiraProjectVersion record in the database using the provided transaction
func (t *JiraProjectVersion) DBCreateTx(ctx context.Context, tx *sql.Tx) (sql.Result, error) {
	q := "INSERT INTO `jira_project_version` (`jira_project_version`.`id`,`jira_project_version`.`checksum`,`jira_project_version`.`version_id`,`jira_project_version`.`project_id`,`jira_project_version`.`name`,`jira_project_version`.`description`,`jira_project_version`.`archived`,`jira_project_version`.`released`,`jira_project_version`.`release_date`,`jira_project_version`.`overdue`,`jira_project_version`.`user_release_date`,`jira_project_version`.`customer_id`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?)"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.VersionID),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.Description),
		orm.ToSQLBool(t.Archived),
		orm.ToSQLBool(t.Released),
		orm.ToSQLInt64(t.ReleaseDate),
		orm.ToSQLBool(t.Overdue),
		orm.ToSQLInt64(t.UserReleaseDate),
		orm.ToSQLString(t.CustomerID),
	)
}

// DBCreateIgnoreDuplicate will upsert the JiraProjectVersion record in the database
func (t *JiraProjectVersion) DBCreateIgnoreDuplicate(ctx context.Context, db *sql.DB) (sql.Result, error) {
	q := "INSERT INTO `jira_project_version` (`jira_project_version`.`id`,`jira_project_version`.`checksum`,`jira_project_version`.`version_id`,`jira_project_version`.`project_id`,`jira_project_version`.`name`,`jira_project_version`.`description`,`jira_project_version`.`archived`,`jira_project_version`.`released`,`jira_project_version`.`release_date`,`jira_project_version`.`overdue`,`jira_project_version`.`user_release_date`,`jira_project_version`.`customer_id`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.VersionID),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.Description),
		orm.ToSQLBool(t.Archived),
		orm.ToSQLBool(t.Released),
		orm.ToSQLInt64(t.ReleaseDate),
		orm.ToSQLBool(t.Overdue),
		orm.ToSQLInt64(t.UserReleaseDate),
		orm.ToSQLString(t.CustomerID),
	)
}

// DBCreateIgnoreDuplicateTx will upsert the JiraProjectVersion record in the database using the provided transaction
func (t *JiraProjectVersion) DBCreateIgnoreDuplicateTx(ctx context.Context, tx *sql.Tx) (sql.Result, error) {
	q := "INSERT INTO `jira_project_version` (`jira_project_version`.`id`,`jira_project_version`.`checksum`,`jira_project_version`.`version_id`,`jira_project_version`.`project_id`,`jira_project_version`.`name`,`jira_project_version`.`description`,`jira_project_version`.`archived`,`jira_project_version`.`released`,`jira_project_version`.`release_date`,`jira_project_version`.`overdue`,`jira_project_version`.`user_release_date`,`jira_project_version`.`customer_id`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.VersionID),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.Description),
		orm.ToSQLBool(t.Archived),
		orm.ToSQLBool(t.Released),
		orm.ToSQLInt64(t.ReleaseDate),
		orm.ToSQLBool(t.Overdue),
		orm.ToSQLInt64(t.UserReleaseDate),
		orm.ToSQLString(t.CustomerID),
	)
}

// DeleteAllJiraProjectVersions deletes all JiraProjectVersion records in the database with optional filters
func DeleteAllJiraProjectVersions(ctx context.Context, db *sql.DB, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(JiraProjectVersionTableName),
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

// DeleteAllJiraProjectVersionsTx deletes all JiraProjectVersion records in the database with optional filters using the provided transaction
func DeleteAllJiraProjectVersionsTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(JiraProjectVersionTableName),
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

// DBDelete will delete this JiraProjectVersion record in the database
func (t *JiraProjectVersion) DBDelete(ctx context.Context, db *sql.DB) (bool, error) {
	q := "DELETE FROM `jira_project_version` WHERE `id` = ?"
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

// DBDeleteTx will delete this JiraProjectVersion record in the database using the provided transaction
func (t *JiraProjectVersion) DBDeleteTx(ctx context.Context, tx *sql.Tx) (bool, error) {
	q := "DELETE FROM `jira_project_version` WHERE `id` = ?"
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

// DBUpdate will update the JiraProjectVersion record in the database
func (t *JiraProjectVersion) DBUpdate(ctx context.Context, db *sql.DB) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `jira_project_version` SET `checksum`=?,`version_id`=?,`project_id`=?,`name`=?,`description`=?,`archived`=?,`released`=?,`release_date`=?,`overdue`=?,`user_release_date`=?,`customer_id`=? WHERE `id`=?"
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.VersionID),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.Description),
		orm.ToSQLBool(t.Archived),
		orm.ToSQLBool(t.Released),
		orm.ToSQLInt64(t.ReleaseDate),
		orm.ToSQLBool(t.Overdue),
		orm.ToSQLInt64(t.UserReleaseDate),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.ID),
	)
}

// DBUpdateTx will update the JiraProjectVersion record in the database using the provided transaction
func (t *JiraProjectVersion) DBUpdateTx(ctx context.Context, tx *sql.Tx) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `jira_project_version` SET `checksum`=?,`version_id`=?,`project_id`=?,`name`=?,`description`=?,`archived`=?,`released`=?,`release_date`=?,`overdue`=?,`user_release_date`=?,`customer_id`=? WHERE `id`=?"
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.VersionID),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.Description),
		orm.ToSQLBool(t.Archived),
		orm.ToSQLBool(t.Released),
		orm.ToSQLInt64(t.ReleaseDate),
		orm.ToSQLBool(t.Overdue),
		orm.ToSQLInt64(t.UserReleaseDate),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.ID),
	)
}

// DBUpsert will upsert the JiraProjectVersion record in the database
func (t *JiraProjectVersion) DBUpsert(ctx context.Context, db *sql.DB, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `jira_project_version` (`jira_project_version`.`id`,`jira_project_version`.`checksum`,`jira_project_version`.`version_id`,`jira_project_version`.`project_id`,`jira_project_version`.`name`,`jira_project_version`.`description`,`jira_project_version`.`archived`,`jira_project_version`.`released`,`jira_project_version`.`release_date`,`jira_project_version`.`overdue`,`jira_project_version`.`user_release_date`,`jira_project_version`.`customer_id`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `jira_project_version` (`jira_project_version`.`id`,`jira_project_version`.`checksum`,`jira_project_version`.`version_id`,`jira_project_version`.`project_id`,`jira_project_version`.`name`,`jira_project_version`.`description`,`jira_project_version`.`archived`,`jira_project_version`.`released`,`jira_project_version`.`release_date`,`jira_project_version`.`overdue`,`jira_project_version`.`user_release_date`,`jira_project_version`.`customer_id`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`version_id`=VALUES(`version_id`),`project_id`=VALUES(`project_id`),`name`=VALUES(`name`),`description`=VALUES(`description`),`archived`=VALUES(`archived`),`released`=VALUES(`released`),`release_date`=VALUES(`release_date`),`overdue`=VALUES(`overdue`),`user_release_date`=VALUES(`user_release_date`),`customer_id`=VALUES(`customer_id`)"
	}
	r, err := db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.VersionID),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.Description),
		orm.ToSQLBool(t.Archived),
		orm.ToSQLBool(t.Released),
		orm.ToSQLInt64(t.ReleaseDate),
		orm.ToSQLBool(t.Overdue),
		orm.ToSQLInt64(t.UserReleaseDate),
		orm.ToSQLString(t.CustomerID),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBUpsertTx will upsert the JiraProjectVersion record in the database using the provided transaction
func (t *JiraProjectVersion) DBUpsertTx(ctx context.Context, tx *sql.Tx, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `jira_project_version` (`jira_project_version`.`id`,`jira_project_version`.`checksum`,`jira_project_version`.`version_id`,`jira_project_version`.`project_id`,`jira_project_version`.`name`,`jira_project_version`.`description`,`jira_project_version`.`archived`,`jira_project_version`.`released`,`jira_project_version`.`release_date`,`jira_project_version`.`overdue`,`jira_project_version`.`user_release_date`,`jira_project_version`.`customer_id`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `jira_project_version` (`jira_project_version`.`id`,`jira_project_version`.`checksum`,`jira_project_version`.`version_id`,`jira_project_version`.`project_id`,`jira_project_version`.`name`,`jira_project_version`.`description`,`jira_project_version`.`archived`,`jira_project_version`.`released`,`jira_project_version`.`release_date`,`jira_project_version`.`overdue`,`jira_project_version`.`user_release_date`,`jira_project_version`.`customer_id`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`version_id`=VALUES(`version_id`),`project_id`=VALUES(`project_id`),`name`=VALUES(`name`),`description`=VALUES(`description`),`archived`=VALUES(`archived`),`released`=VALUES(`released`),`release_date`=VALUES(`release_date`),`overdue`=VALUES(`overdue`),`user_release_date`=VALUES(`user_release_date`),`customer_id`=VALUES(`customer_id`)"
	}
	r, err := tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.VersionID),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.Description),
		orm.ToSQLBool(t.Archived),
		orm.ToSQLBool(t.Released),
		orm.ToSQLInt64(t.ReleaseDate),
		orm.ToSQLBool(t.Overdue),
		orm.ToSQLInt64(t.UserReleaseDate),
		orm.ToSQLString(t.CustomerID),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBFindOne will find a JiraProjectVersion record in the database with the primary key
func (t *JiraProjectVersion) DBFindOne(ctx context.Context, db *sql.DB, value string) (bool, error) {
	q := "SELECT `jira_project_version`.`id`,`jira_project_version`.`checksum`,`jira_project_version`.`version_id`,`jira_project_version`.`project_id`,`jira_project_version`.`name`,`jira_project_version`.`description`,`jira_project_version`.`archived`,`jira_project_version`.`released`,`jira_project_version`.`release_date`,`jira_project_version`.`overdue`,`jira_project_version`.`user_release_date`,`jira_project_version`.`customer_id` FROM `jira_project_version` WHERE `id` = ? LIMIT 1"
	row := db.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _VersionID sql.NullString
	var _ProjectID sql.NullString
	var _Name sql.NullString
	var _Description sql.NullString
	var _Archived sql.NullBool
	var _Released sql.NullBool
	var _ReleaseDate sql.NullInt64
	var _Overdue sql.NullBool
	var _UserReleaseDate sql.NullInt64
	var _CustomerID sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_VersionID,
		&_ProjectID,
		&_Name,
		&_Description,
		&_Archived,
		&_Released,
		&_ReleaseDate,
		&_Overdue,
		&_UserReleaseDate,
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
	if _VersionID.Valid {
		t.SetVersionID(_VersionID.String)
	}
	if _ProjectID.Valid {
		t.SetProjectID(_ProjectID.String)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _Description.Valid {
		t.SetDescription(_Description.String)
	}
	if _Archived.Valid {
		t.SetArchived(_Archived.Bool)
	}
	if _Released.Valid {
		t.SetReleased(_Released.Bool)
	}
	if _ReleaseDate.Valid {
		t.SetReleaseDate(_ReleaseDate.Int64)
	}
	if _Overdue.Valid {
		t.SetOverdue(_Overdue.Bool)
	}
	if _UserReleaseDate.Valid {
		t.SetUserReleaseDate(_UserReleaseDate.Int64)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	return true, nil
}

// DBFindOneTx will find a JiraProjectVersion record in the database with the primary key using the provided transaction
func (t *JiraProjectVersion) DBFindOneTx(ctx context.Context, tx *sql.Tx, value string) (bool, error) {
	q := "SELECT `jira_project_version`.`id`,`jira_project_version`.`checksum`,`jira_project_version`.`version_id`,`jira_project_version`.`project_id`,`jira_project_version`.`name`,`jira_project_version`.`description`,`jira_project_version`.`archived`,`jira_project_version`.`released`,`jira_project_version`.`release_date`,`jira_project_version`.`overdue`,`jira_project_version`.`user_release_date`,`jira_project_version`.`customer_id` FROM `jira_project_version` WHERE `id` = ? LIMIT 1"
	row := tx.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _VersionID sql.NullString
	var _ProjectID sql.NullString
	var _Name sql.NullString
	var _Description sql.NullString
	var _Archived sql.NullBool
	var _Released sql.NullBool
	var _ReleaseDate sql.NullInt64
	var _Overdue sql.NullBool
	var _UserReleaseDate sql.NullInt64
	var _CustomerID sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_VersionID,
		&_ProjectID,
		&_Name,
		&_Description,
		&_Archived,
		&_Released,
		&_ReleaseDate,
		&_Overdue,
		&_UserReleaseDate,
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
	if _VersionID.Valid {
		t.SetVersionID(_VersionID.String)
	}
	if _ProjectID.Valid {
		t.SetProjectID(_ProjectID.String)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _Description.Valid {
		t.SetDescription(_Description.String)
	}
	if _Archived.Valid {
		t.SetArchived(_Archived.Bool)
	}
	if _Released.Valid {
		t.SetReleased(_Released.Bool)
	}
	if _ReleaseDate.Valid {
		t.SetReleaseDate(_ReleaseDate.Int64)
	}
	if _Overdue.Valid {
		t.SetOverdue(_Overdue.Bool)
	}
	if _UserReleaseDate.Valid {
		t.SetUserReleaseDate(_UserReleaseDate.Int64)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	return true, nil
}

// FindJiraProjectVersions will find a JiraProjectVersion record in the database with the provided parameters
func FindJiraProjectVersions(ctx context.Context, db *sql.DB, _params ...interface{}) ([]*JiraProjectVersion, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("version_id"),
		orm.Column("project_id"),
		orm.Column("name"),
		orm.Column("description"),
		orm.Column("archived"),
		orm.Column("released"),
		orm.Column("release_date"),
		orm.Column("overdue"),
		orm.Column("user_release_date"),
		orm.Column("customer_id"),
		orm.Table(JiraProjectVersionTableName),
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
	results := make([]*JiraProjectVersion, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _VersionID sql.NullString
		var _ProjectID sql.NullString
		var _Name sql.NullString
		var _Description sql.NullString
		var _Archived sql.NullBool
		var _Released sql.NullBool
		var _ReleaseDate sql.NullInt64
		var _Overdue sql.NullBool
		var _UserReleaseDate sql.NullInt64
		var _CustomerID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_VersionID,
			&_ProjectID,
			&_Name,
			&_Description,
			&_Archived,
			&_Released,
			&_ReleaseDate,
			&_Overdue,
			&_UserReleaseDate,
			&_CustomerID,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraProjectVersion{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _VersionID.Valid {
			t.SetVersionID(_VersionID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _Description.Valid {
			t.SetDescription(_Description.String)
		}
		if _Archived.Valid {
			t.SetArchived(_Archived.Bool)
		}
		if _Released.Valid {
			t.SetReleased(_Released.Bool)
		}
		if _ReleaseDate.Valid {
			t.SetReleaseDate(_ReleaseDate.Int64)
		}
		if _Overdue.Valid {
			t.SetOverdue(_Overdue.Bool)
		}
		if _UserReleaseDate.Valid {
			t.SetUserReleaseDate(_UserReleaseDate.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindJiraProjectVersionsTx will find a JiraProjectVersion record in the database with the provided parameters using the provided transaction
func FindJiraProjectVersionsTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) ([]*JiraProjectVersion, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("version_id"),
		orm.Column("project_id"),
		orm.Column("name"),
		orm.Column("description"),
		orm.Column("archived"),
		orm.Column("released"),
		orm.Column("release_date"),
		orm.Column("overdue"),
		orm.Column("user_release_date"),
		orm.Column("customer_id"),
		orm.Table(JiraProjectVersionTableName),
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
	results := make([]*JiraProjectVersion, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _VersionID sql.NullString
		var _ProjectID sql.NullString
		var _Name sql.NullString
		var _Description sql.NullString
		var _Archived sql.NullBool
		var _Released sql.NullBool
		var _ReleaseDate sql.NullInt64
		var _Overdue sql.NullBool
		var _UserReleaseDate sql.NullInt64
		var _CustomerID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_VersionID,
			&_ProjectID,
			&_Name,
			&_Description,
			&_Archived,
			&_Released,
			&_ReleaseDate,
			&_Overdue,
			&_UserReleaseDate,
			&_CustomerID,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraProjectVersion{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _VersionID.Valid {
			t.SetVersionID(_VersionID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _Description.Valid {
			t.SetDescription(_Description.String)
		}
		if _Archived.Valid {
			t.SetArchived(_Archived.Bool)
		}
		if _Released.Valid {
			t.SetReleased(_Released.Bool)
		}
		if _ReleaseDate.Valid {
			t.SetReleaseDate(_ReleaseDate.Int64)
		}
		if _Overdue.Valid {
			t.SetOverdue(_Overdue.Bool)
		}
		if _UserReleaseDate.Valid {
			t.SetUserReleaseDate(_UserReleaseDate.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// DBFind will find a JiraProjectVersion record in the database with the provided parameters
func (t *JiraProjectVersion) DBFind(ctx context.Context, db *sql.DB, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("version_id"),
		orm.Column("project_id"),
		orm.Column("name"),
		orm.Column("description"),
		orm.Column("archived"),
		orm.Column("released"),
		orm.Column("release_date"),
		orm.Column("overdue"),
		orm.Column("user_release_date"),
		orm.Column("customer_id"),
		orm.Table(JiraProjectVersionTableName),
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
	var _VersionID sql.NullString
	var _ProjectID sql.NullString
	var _Name sql.NullString
	var _Description sql.NullString
	var _Archived sql.NullBool
	var _Released sql.NullBool
	var _ReleaseDate sql.NullInt64
	var _Overdue sql.NullBool
	var _UserReleaseDate sql.NullInt64
	var _CustomerID sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_VersionID,
		&_ProjectID,
		&_Name,
		&_Description,
		&_Archived,
		&_Released,
		&_ReleaseDate,
		&_Overdue,
		&_UserReleaseDate,
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
	if _VersionID.Valid {
		t.SetVersionID(_VersionID.String)
	}
	if _ProjectID.Valid {
		t.SetProjectID(_ProjectID.String)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _Description.Valid {
		t.SetDescription(_Description.String)
	}
	if _Archived.Valid {
		t.SetArchived(_Archived.Bool)
	}
	if _Released.Valid {
		t.SetReleased(_Released.Bool)
	}
	if _ReleaseDate.Valid {
		t.SetReleaseDate(_ReleaseDate.Int64)
	}
	if _Overdue.Valid {
		t.SetOverdue(_Overdue.Bool)
	}
	if _UserReleaseDate.Valid {
		t.SetUserReleaseDate(_UserReleaseDate.Int64)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	return true, nil
}

// DBFindTx will find a JiraProjectVersion record in the database with the provided parameters using the provided transaction
func (t *JiraProjectVersion) DBFindTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("version_id"),
		orm.Column("project_id"),
		orm.Column("name"),
		orm.Column("description"),
		orm.Column("archived"),
		orm.Column("released"),
		orm.Column("release_date"),
		orm.Column("overdue"),
		orm.Column("user_release_date"),
		orm.Column("customer_id"),
		orm.Table(JiraProjectVersionTableName),
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
	var _VersionID sql.NullString
	var _ProjectID sql.NullString
	var _Name sql.NullString
	var _Description sql.NullString
	var _Archived sql.NullBool
	var _Released sql.NullBool
	var _ReleaseDate sql.NullInt64
	var _Overdue sql.NullBool
	var _UserReleaseDate sql.NullInt64
	var _CustomerID sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_VersionID,
		&_ProjectID,
		&_Name,
		&_Description,
		&_Archived,
		&_Released,
		&_ReleaseDate,
		&_Overdue,
		&_UserReleaseDate,
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
	if _VersionID.Valid {
		t.SetVersionID(_VersionID.String)
	}
	if _ProjectID.Valid {
		t.SetProjectID(_ProjectID.String)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _Description.Valid {
		t.SetDescription(_Description.String)
	}
	if _Archived.Valid {
		t.SetArchived(_Archived.Bool)
	}
	if _Released.Valid {
		t.SetReleased(_Released.Bool)
	}
	if _ReleaseDate.Valid {
		t.SetReleaseDate(_ReleaseDate.Int64)
	}
	if _Overdue.Valid {
		t.SetOverdue(_Overdue.Bool)
	}
	if _UserReleaseDate.Valid {
		t.SetUserReleaseDate(_UserReleaseDate.Int64)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	return true, nil
}

// CountJiraProjectVersions will find the count of JiraProjectVersion records in the database
func CountJiraProjectVersions(ctx context.Context, db *sql.DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(JiraProjectVersionTableName),
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

// CountJiraProjectVersionsTx will find the count of JiraProjectVersion records in the database using the provided transaction
func CountJiraProjectVersionsTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(JiraProjectVersionTableName),
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

// DBCount will find the count of JiraProjectVersion records in the database
func (t *JiraProjectVersion) DBCount(ctx context.Context, db *sql.DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(JiraProjectVersionTableName),
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

// DBCountTx will find the count of JiraProjectVersion records in the database using the provided transaction
func (t *JiraProjectVersion) DBCountTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(JiraProjectVersionTableName),
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

// DBExists will return true if the JiraProjectVersion record exists in the database
func (t *JiraProjectVersion) DBExists(ctx context.Context, db *sql.DB) (bool, error) {
	q := "SELECT `id` FROM `jira_project_version` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := db.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// DBExistsTx will return true if the JiraProjectVersion record exists in the database using the provided transaction
func (t *JiraProjectVersion) DBExistsTx(ctx context.Context, tx *sql.Tx) (bool, error) {
	q := "SELECT `id` FROM `jira_project_version` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := tx.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// PrimaryKeyColumn returns the column name for the primary key
func (t *JiraProjectVersion) PrimaryKeyColumn() string {
	return JiraProjectVersionColumnID
}

// PrimaryKeyColumnType returns the primary key column Go type as a string
func (t *JiraProjectVersion) PrimaryKeyColumnType() string {
	return "string"
}

// PrimaryKey returns the primary key column value
func (t *JiraProjectVersion) PrimaryKey() interface{} {
	return t.ID
}
