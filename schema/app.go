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

var _ Model = (*App)(nil)
var _ CSVWriter = (*App)(nil)
var _ JSONWriter = (*App)(nil)
var _ Checksum = (*App)(nil)

// AppTableName is the name of the table in SQL
const AppTableName = "app"

var AppColumns = []string{
	"id",
	"checksum",
	"name",
	"description",
	"active",
	"repo_ids",
	"created_at",
	"updated_at",
	"customer_id",
}

// App table
type App struct {
	Active      bool    `json:"active"`
	Checksum    *string `json:"checksum,omitempty"`
	CreatedAt   int64   `json:"created_at"`
	CustomerID  string  `json:"customer_id"`
	Description *string `json:"description,omitempty"`
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	RepoIds     *string `json:"repo_ids,omitempty"`
	UpdatedAt   *int64  `json:"updated_at,omitempty"`
}

// TableName returns the SQL table name for App and satifies the Model interface
func (t *App) TableName() string {
	return AppTableName
}

// ToCSV will serialize the App instance to a CSV compatible array of strings
func (t *App) ToCSV() []string {
	return []string{
		t.ID,
		t.CalculateChecksum(),
		t.Name,
		toCSVString(t.Description),
		toCSVBool(t.Active),
		toCSVString(t.RepoIds),
		toCSVString(t.CreatedAt),
		toCSVString(t.UpdatedAt),
		t.CustomerID,
	}
}

// WriteCSV will serialize the App instance to the writer as CSV and satisfies the CSVWriter interface
func (t *App) WriteCSV(w *csv.Writer) error {
	return w.Write(t.ToCSV())
}

// WriteJSON will serialize the App instance to the writer as JSON and satisfies the JSONWriter interface
func (t *App) WriteJSON(w io.Writer, indent ...bool) error {
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

// NewAppReader creates a JSON reader which can read in App objects serialized as JSON either as an array, single object or json new lines
// and writes each App to the channel provided
func NewAppReader(r io.Reader, ch chan<- App) error {
	return orm.Deserialize(r, func(buf json.RawMessage) error {
		dec := json.NewDecoder(bytes.NewBuffer(buf))
		e := App{}
		if err := dec.Decode(&e); err != nil {
			return err
		}
		ch <- e
		return nil
	})
}

// NewCSVAppReaderDir will read the reader as CSV and emit each record to the channel provided
func NewCSVAppReader(r io.Reader, ch chan<- App) error {
	cr := csv.NewReader(r)
	for {
		record, err := cr.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		ch <- App{
			ID:          record[0],
			Checksum:    fromStringPointer(record[1]),
			Name:        record[2],
			Description: fromStringPointer(record[3]),
			Active:      fromCSVBool(record[4]),
			RepoIds:     fromStringPointer(record[5]),
			CreatedAt:   fromCSVInt64(record[6]),
			UpdatedAt:   fromCSVInt64Pointer(record[7]),
			CustomerID:  record[8],
		}
	}
	return nil
}

// NewCSVAppReaderFile will read the file as a CSV and emit each record to the channel provided
func NewCSVAppReaderFile(fp string, ch chan<- App) error {
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
	return NewCSVAppReader(fc, ch)
}

// NewCSVAppReaderDir will read the app.csv.gz file as a CSV and emit each record to the channel provided
func NewCSVAppReaderDir(dir string, ch chan<- App) error {
	return NewCSVAppReaderFile(filepath.Join(dir, "app.csv.gz"), ch)
}

// AppCSVDeduper is a function callback which takes the existing value (a) and the new value (b)
// and the return value should be the one to use (or a new one, if applicable). return nil
// to skip processing of this record
type AppCSVDeduper func(a App, b App) *App

// AppCSVDedupeDisabled is set on whether the CSV writer should de-dupe values by key
var AppCSVDedupeDisabled bool

// NewAppCSVWriterSize creates a batch writer that will write each App into a CSV file
// this method will automatically de-duplicate entries using the primary key. if the checksum
// for a newer item with the same primary key doesn't match a previously sent item, the newer
// one will be used
func NewAppCSVWriterSize(w io.Writer, size int, dedupers ...AppCSVDeduper) (chan App, chan bool, error) {
	cw := csv.NewWriter(w)
	ch := make(chan App, size)
	done := make(chan bool)
	go func() {
		defer func() { done <- true }()
		dodedupe := !AppCSVDedupeDisabled
		var kv map[string]*App
		var deduper AppCSVDeduper
		if dedupers != nil && len(dedupers) > 0 {
			deduper = dedupers[0]
			dodedupe = true
		}
		if dodedupe {
			kv = make(map[string]*App)
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

// AppCSVDefaultSize is the default channel buffer size if not provided
var AppCSVDefaultSize = 100

// NewAppCSVWriter creates a batch writer that will write each App into a CSV file
func NewAppCSVWriter(w io.Writer, dedupers ...AppCSVDeduper) (chan App, chan bool, error) {
	return NewAppCSVWriterSize(w, AppCSVDefaultSize, dedupers...)
}

// NewAppCSVWriterDir creates a batch writer that will write each App into a CSV file named app.csv.gz in dir
func NewAppCSVWriterDir(dir string, dedupers ...AppCSVDeduper) (chan App, chan bool, error) {
	return NewAppCSVWriterFile(filepath.Join(dir, "app.csv.gz"), dedupers...)
}

// NewAppCSVWriterFile creates a batch writer that will write each App into a CSV file
func NewAppCSVWriterFile(fn string, dedupers ...AppCSVDeduper) (chan App, chan bool, error) {
	f, err := os.Create(fn)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening CSV file %s. %v", fn, err)
	}
	var fc io.WriteCloser = f
	if filepath.Ext(fn) == ".gz" {
		w, _ := gzip.NewWriterLevel(f, gzip.BestCompression)
		fc = w
	}
	ch, done, err := NewAppCSVWriter(fc, dedupers...)
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

type AppDBAction func(ctx context.Context, db *sql.DB, record App) error

// NewAppDBWriterSize creates a DB writer that will write each issue into the DB
func NewAppDBWriterSize(ctx context.Context, db *sql.DB, errors chan<- error, size int, actions ...AppDBAction) (chan App, chan bool, error) {
	ch := make(chan App, size)
	done := make(chan bool)
	var action AppDBAction
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

// NewAppDBWriter creates a DB writer that will write each issue into the DB
func NewAppDBWriter(ctx context.Context, db *sql.DB, errors chan<- error, actions ...AppDBAction) (chan App, chan bool, error) {
	return NewAppDBWriterSize(ctx, db, errors, 100, actions...)
}

// AppColumnID is the ID SQL column name for the App table
const AppColumnID = "id"

// AppEscapedColumnID is the escaped ID SQL column name for the App table
const AppEscapedColumnID = "`id`"

// AppColumnChecksum is the Checksum SQL column name for the App table
const AppColumnChecksum = "checksum"

// AppEscapedColumnChecksum is the escaped Checksum SQL column name for the App table
const AppEscapedColumnChecksum = "`checksum`"

// AppColumnName is the Name SQL column name for the App table
const AppColumnName = "name"

// AppEscapedColumnName is the escaped Name SQL column name for the App table
const AppEscapedColumnName = "`name`"

// AppColumnDescription is the Description SQL column name for the App table
const AppColumnDescription = "description"

// AppEscapedColumnDescription is the escaped Description SQL column name for the App table
const AppEscapedColumnDescription = "`description`"

// AppColumnActive is the Active SQL column name for the App table
const AppColumnActive = "active"

// AppEscapedColumnActive is the escaped Active SQL column name for the App table
const AppEscapedColumnActive = "`active`"

// AppColumnRepoIds is the RepoIds SQL column name for the App table
const AppColumnRepoIds = "repo_ids"

// AppEscapedColumnRepoIds is the escaped RepoIds SQL column name for the App table
const AppEscapedColumnRepoIds = "`repo_ids`"

// AppColumnCreatedAt is the CreatedAt SQL column name for the App table
const AppColumnCreatedAt = "created_at"

// AppEscapedColumnCreatedAt is the escaped CreatedAt SQL column name for the App table
const AppEscapedColumnCreatedAt = "`created_at`"

// AppColumnUpdatedAt is the UpdatedAt SQL column name for the App table
const AppColumnUpdatedAt = "updated_at"

// AppEscapedColumnUpdatedAt is the escaped UpdatedAt SQL column name for the App table
const AppEscapedColumnUpdatedAt = "`updated_at`"

// AppColumnCustomerID is the CustomerID SQL column name for the App table
const AppColumnCustomerID = "customer_id"

// AppEscapedColumnCustomerID is the escaped CustomerID SQL column name for the App table
const AppEscapedColumnCustomerID = "`customer_id`"

// GetID will return the App ID value
func (t *App) GetID() string {
	return t.ID
}

// SetID will set the App ID value
func (t *App) SetID(v string) {
	t.ID = v
}

// FindAppByID will find a App by ID
func FindAppByID(ctx context.Context, db *sql.DB, value string) (*App, error) {
	q := "SELECT `app`.`id`,`app`.`checksum`,`app`.`name`,`app`.`description`,`app`.`active`,`app`.`repo_ids`,`app`.`created_at`,`app`.`updated_at`,`app`.`customer_id` FROM `app` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _Name sql.NullString
	var _Description sql.NullString
	var _Active sql.NullBool
	var _RepoIds sql.NullString
	var _CreatedAt sql.NullInt64
	var _UpdatedAt sql.NullInt64
	var _CustomerID sql.NullString
	err := db.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_Name,
		&_Description,
		&_Active,
		&_RepoIds,
		&_CreatedAt,
		&_UpdatedAt,
		&_CustomerID,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &App{}
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
	if _Active.Valid {
		t.SetActive(_Active.Bool)
	}
	if _RepoIds.Valid {
		t.SetRepoIds(_RepoIds.String)
	}
	if _CreatedAt.Valid {
		t.SetCreatedAt(_CreatedAt.Int64)
	}
	if _UpdatedAt.Valid {
		t.SetUpdatedAt(_UpdatedAt.Int64)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	return t, nil
}

// FindAppByIDTx will find a App by ID using the provided transaction
func FindAppByIDTx(ctx context.Context, tx *sql.Tx, value string) (*App, error) {
	q := "SELECT `app`.`id`,`app`.`checksum`,`app`.`name`,`app`.`description`,`app`.`active`,`app`.`repo_ids`,`app`.`created_at`,`app`.`updated_at`,`app`.`customer_id` FROM `app` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _Name sql.NullString
	var _Description sql.NullString
	var _Active sql.NullBool
	var _RepoIds sql.NullString
	var _CreatedAt sql.NullInt64
	var _UpdatedAt sql.NullInt64
	var _CustomerID sql.NullString
	err := tx.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_Name,
		&_Description,
		&_Active,
		&_RepoIds,
		&_CreatedAt,
		&_UpdatedAt,
		&_CustomerID,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &App{}
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
	if _Active.Valid {
		t.SetActive(_Active.Bool)
	}
	if _RepoIds.Valid {
		t.SetRepoIds(_RepoIds.String)
	}
	if _CreatedAt.Valid {
		t.SetCreatedAt(_CreatedAt.Int64)
	}
	if _UpdatedAt.Valid {
		t.SetUpdatedAt(_UpdatedAt.Int64)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	return t, nil
}

// GetChecksum will return the App Checksum value
func (t *App) GetChecksum() string {
	if t.Checksum == nil {
		return ""
	}
	return *t.Checksum
}

// SetChecksum will set the App Checksum value
func (t *App) SetChecksum(v string) {
	t.Checksum = &v
}

// GetName will return the App Name value
func (t *App) GetName() string {
	return t.Name
}

// SetName will set the App Name value
func (t *App) SetName(v string) {
	t.Name = v
}

// FindAppsByName will find all Apps by the Name value
func FindAppsByName(ctx context.Context, db *sql.DB, value string) ([]*App, error) {
	q := "SELECT `app`.`id`,`app`.`checksum`,`app`.`name`,`app`.`description`,`app`.`active`,`app`.`repo_ids`,`app`.`created_at`,`app`.`updated_at`,`app`.`customer_id` FROM `app` WHERE `name` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*App, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _Name sql.NullString
		var _Description sql.NullString
		var _Active sql.NullBool
		var _RepoIds sql.NullString
		var _CreatedAt sql.NullInt64
		var _UpdatedAt sql.NullInt64
		var _CustomerID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_Name,
			&_Description,
			&_Active,
			&_RepoIds,
			&_CreatedAt,
			&_UpdatedAt,
			&_CustomerID,
		)
		if err != nil {
			return nil, err
		}
		t := &App{}
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
		if _Active.Valid {
			t.SetActive(_Active.Bool)
		}
		if _RepoIds.Valid {
			t.SetRepoIds(_RepoIds.String)
		}
		if _CreatedAt.Valid {
			t.SetCreatedAt(_CreatedAt.Int64)
		}
		if _UpdatedAt.Valid {
			t.SetUpdatedAt(_UpdatedAt.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindAppsByNameTx will find all Apps by the Name value using the provided transaction
func FindAppsByNameTx(ctx context.Context, tx *sql.Tx, value string) ([]*App, error) {
	q := "SELECT `app`.`id`,`app`.`checksum`,`app`.`name`,`app`.`description`,`app`.`active`,`app`.`repo_ids`,`app`.`created_at`,`app`.`updated_at`,`app`.`customer_id` FROM `app` WHERE `name` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*App, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _Name sql.NullString
		var _Description sql.NullString
		var _Active sql.NullBool
		var _RepoIds sql.NullString
		var _CreatedAt sql.NullInt64
		var _UpdatedAt sql.NullInt64
		var _CustomerID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_Name,
			&_Description,
			&_Active,
			&_RepoIds,
			&_CreatedAt,
			&_UpdatedAt,
			&_CustomerID,
		)
		if err != nil {
			return nil, err
		}
		t := &App{}
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
		if _Active.Valid {
			t.SetActive(_Active.Bool)
		}
		if _RepoIds.Valid {
			t.SetRepoIds(_RepoIds.String)
		}
		if _CreatedAt.Valid {
			t.SetCreatedAt(_CreatedAt.Int64)
		}
		if _UpdatedAt.Valid {
			t.SetUpdatedAt(_UpdatedAt.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetDescription will return the App Description value
func (t *App) GetDescription() string {
	if t.Description == nil {
		return ""
	}
	return *t.Description
}

// SetDescription will set the App Description value
func (t *App) SetDescription(v string) {
	t.Description = &v
}

// GetActive will return the App Active value
func (t *App) GetActive() bool {
	return t.Active
}

// SetActive will set the App Active value
func (t *App) SetActive(v bool) {
	t.Active = v
}

// GetRepoIds will return the App RepoIds value
func (t *App) GetRepoIds() string {
	if t.RepoIds == nil {
		return ""
	}
	return *t.RepoIds
}

// SetRepoIds will set the App RepoIds value
func (t *App) SetRepoIds(v string) {
	t.RepoIds = &v
}

// GetCreatedAt will return the App CreatedAt value
func (t *App) GetCreatedAt() int64 {
	return t.CreatedAt
}

// SetCreatedAt will set the App CreatedAt value
func (t *App) SetCreatedAt(v int64) {
	t.CreatedAt = v
}

// GetUpdatedAt will return the App UpdatedAt value
func (t *App) GetUpdatedAt() int64 {
	if t.UpdatedAt == nil {
		return int64(0)
	}
	return *t.UpdatedAt
}

// SetUpdatedAt will set the App UpdatedAt value
func (t *App) SetUpdatedAt(v int64) {
	t.UpdatedAt = &v
}

// GetCustomerID will return the App CustomerID value
func (t *App) GetCustomerID() string {
	return t.CustomerID
}

// SetCustomerID will set the App CustomerID value
func (t *App) SetCustomerID(v string) {
	t.CustomerID = v
}

// FindAppsByCustomerID will find all Apps by the CustomerID value
func FindAppsByCustomerID(ctx context.Context, db *sql.DB, value string) ([]*App, error) {
	q := "SELECT `app`.`id`,`app`.`checksum`,`app`.`name`,`app`.`description`,`app`.`active`,`app`.`repo_ids`,`app`.`created_at`,`app`.`updated_at`,`app`.`customer_id` FROM `app` WHERE `customer_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*App, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _Name sql.NullString
		var _Description sql.NullString
		var _Active sql.NullBool
		var _RepoIds sql.NullString
		var _CreatedAt sql.NullInt64
		var _UpdatedAt sql.NullInt64
		var _CustomerID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_Name,
			&_Description,
			&_Active,
			&_RepoIds,
			&_CreatedAt,
			&_UpdatedAt,
			&_CustomerID,
		)
		if err != nil {
			return nil, err
		}
		t := &App{}
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
		if _Active.Valid {
			t.SetActive(_Active.Bool)
		}
		if _RepoIds.Valid {
			t.SetRepoIds(_RepoIds.String)
		}
		if _CreatedAt.Valid {
			t.SetCreatedAt(_CreatedAt.Int64)
		}
		if _UpdatedAt.Valid {
			t.SetUpdatedAt(_UpdatedAt.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindAppsByCustomerIDTx will find all Apps by the CustomerID value using the provided transaction
func FindAppsByCustomerIDTx(ctx context.Context, tx *sql.Tx, value string) ([]*App, error) {
	q := "SELECT `app`.`id`,`app`.`checksum`,`app`.`name`,`app`.`description`,`app`.`active`,`app`.`repo_ids`,`app`.`created_at`,`app`.`updated_at`,`app`.`customer_id` FROM `app` WHERE `customer_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*App, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _Name sql.NullString
		var _Description sql.NullString
		var _Active sql.NullBool
		var _RepoIds sql.NullString
		var _CreatedAt sql.NullInt64
		var _UpdatedAt sql.NullInt64
		var _CustomerID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_Name,
			&_Description,
			&_Active,
			&_RepoIds,
			&_CreatedAt,
			&_UpdatedAt,
			&_CustomerID,
		)
		if err != nil {
			return nil, err
		}
		t := &App{}
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
		if _Active.Valid {
			t.SetActive(_Active.Bool)
		}
		if _RepoIds.Valid {
			t.SetRepoIds(_RepoIds.String)
		}
		if _CreatedAt.Valid {
			t.SetCreatedAt(_CreatedAt.Int64)
		}
		if _UpdatedAt.Valid {
			t.SetUpdatedAt(_UpdatedAt.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

func (t *App) toTimestamp(value time.Time) *timestamp.Timestamp {
	ts, _ := ptypes.TimestampProto(value)
	return ts
}

// DBCreateAppTable will create the App table
func DBCreateAppTable(ctx context.Context, db *sql.DB) error {
	q := "CREATE TABLE `app` (`id`VARCHAR(64) NOT NULL PRIMARY KEY,`checksum`CHAR(64),`name` VARCHAR(255) NOT NULL,`description`TEXT,`active` BOOL NOT NULL,`repo_ids`JSON,`created_at` BIGINT UNSIGNED NOT NULL,`updated_at` BIGINT UNSIGNED,`customer_id`VARCHAR(64) NOT NULL,INDEX app_name_index (`name`),INDEX app_customer_id_index (`customer_id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBCreateAppTableTx will create the App table using the provided transction
func DBCreateAppTableTx(ctx context.Context, tx *sql.Tx) error {
	q := "CREATE TABLE `app` (`id`VARCHAR(64) NOT NULL PRIMARY KEY,`checksum`CHAR(64),`name` VARCHAR(255) NOT NULL,`description`TEXT,`active` BOOL NOT NULL,`repo_ids`JSON,`created_at` BIGINT UNSIGNED NOT NULL,`updated_at` BIGINT UNSIGNED,`customer_id`VARCHAR(64) NOT NULL,INDEX app_name_index (`name`),INDEX app_customer_id_index (`customer_id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// DBDropAppTable will drop the App table
func DBDropAppTable(ctx context.Context, db *sql.DB) error {
	q := "DROP TABLE IF EXISTS `app`"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBDropAppTableTx will drop the App table using the provided transaction
func DBDropAppTableTx(ctx context.Context, tx *sql.Tx) error {
	q := "DROP TABLE IF EXISTS `app`"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// CalculateChecksum will calculate a checksum of the SHA1 of all field values
func (t *App) CalculateChecksum() string {
	return orm.HashStrings(
		orm.ToString(t.ID),
		orm.ToString(t.Name),
		orm.ToString(t.Description),
		orm.ToString(t.Active),
		orm.ToString(t.RepoIds),
		orm.ToString(t.CreatedAt),
		orm.ToString(t.UpdatedAt),
		orm.ToString(t.CustomerID),
	)
}

// DBCreate will create a new App record in the database
func (t *App) DBCreate(ctx context.Context, db *sql.DB) (sql.Result, error) {
	q := "INSERT INTO `app` (`app`.`id`,`app`.`checksum`,`app`.`name`,`app`.`description`,`app`.`active`,`app`.`repo_ids`,`app`.`created_at`,`app`.`updated_at`,`app`.`customer_id`) VALUES (?,?,?,?,?,?,?,?,?)"
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
		orm.ToSQLBool(t.Active),
		orm.ToSQLString(t.RepoIds),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
		orm.ToSQLString(t.CustomerID),
	)
}

// DBCreateTx will create a new App record in the database using the provided transaction
func (t *App) DBCreateTx(ctx context.Context, tx *sql.Tx) (sql.Result, error) {
	q := "INSERT INTO `app` (`app`.`id`,`app`.`checksum`,`app`.`name`,`app`.`description`,`app`.`active`,`app`.`repo_ids`,`app`.`created_at`,`app`.`updated_at`,`app`.`customer_id`) VALUES (?,?,?,?,?,?,?,?,?)"
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
		orm.ToSQLBool(t.Active),
		orm.ToSQLString(t.RepoIds),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
		orm.ToSQLString(t.CustomerID),
	)
}

// DBCreateIgnoreDuplicate will upsert the App record in the database
func (t *App) DBCreateIgnoreDuplicate(ctx context.Context, db *sql.DB) (sql.Result, error) {
	q := "INSERT INTO `app` (`app`.`id`,`app`.`checksum`,`app`.`name`,`app`.`description`,`app`.`active`,`app`.`repo_ids`,`app`.`created_at`,`app`.`updated_at`,`app`.`customer_id`) VALUES (?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
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
		orm.ToSQLBool(t.Active),
		orm.ToSQLString(t.RepoIds),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
		orm.ToSQLString(t.CustomerID),
	)
}

// DBCreateIgnoreDuplicateTx will upsert the App record in the database using the provided transaction
func (t *App) DBCreateIgnoreDuplicateTx(ctx context.Context, tx *sql.Tx) (sql.Result, error) {
	q := "INSERT INTO `app` (`app`.`id`,`app`.`checksum`,`app`.`name`,`app`.`description`,`app`.`active`,`app`.`repo_ids`,`app`.`created_at`,`app`.`updated_at`,`app`.`customer_id`) VALUES (?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
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
		orm.ToSQLBool(t.Active),
		orm.ToSQLString(t.RepoIds),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
		orm.ToSQLString(t.CustomerID),
	)
}

// DeleteAllApps deletes all App records in the database with optional filters
func DeleteAllApps(ctx context.Context, db *sql.DB, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(AppTableName),
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

// DeleteAllAppsTx deletes all App records in the database with optional filters using the provided transaction
func DeleteAllAppsTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(AppTableName),
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

// DBDelete will delete this App record in the database
func (t *App) DBDelete(ctx context.Context, db *sql.DB) (bool, error) {
	q := "DELETE FROM `app` WHERE `id` = ?"
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

// DBDeleteTx will delete this App record in the database using the provided transaction
func (t *App) DBDeleteTx(ctx context.Context, tx *sql.Tx) (bool, error) {
	q := "DELETE FROM `app` WHERE `id` = ?"
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

// DBUpdate will update the App record in the database
func (t *App) DBUpdate(ctx context.Context, db *sql.DB) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `app` SET `checksum`=?,`name`=?,`description`=?,`active`=?,`repo_ids`=?,`created_at`=?,`updated_at`=?,`customer_id`=? WHERE `id`=?"
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.Description),
		orm.ToSQLBool(t.Active),
		orm.ToSQLString(t.RepoIds),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.ID),
	)
}

// DBUpdateTx will update the App record in the database using the provided transaction
func (t *App) DBUpdateTx(ctx context.Context, tx *sql.Tx) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `app` SET `checksum`=?,`name`=?,`description`=?,`active`=?,`repo_ids`=?,`created_at`=?,`updated_at`=?,`customer_id`=? WHERE `id`=?"
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.Description),
		orm.ToSQLBool(t.Active),
		orm.ToSQLString(t.RepoIds),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.ID),
	)
}

// DBUpsert will upsert the App record in the database
func (t *App) DBUpsert(ctx context.Context, db *sql.DB, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `app` (`app`.`id`,`app`.`checksum`,`app`.`name`,`app`.`description`,`app`.`active`,`app`.`repo_ids`,`app`.`created_at`,`app`.`updated_at`,`app`.`customer_id`) VALUES (?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `app` (`app`.`id`,`app`.`checksum`,`app`.`name`,`app`.`description`,`app`.`active`,`app`.`repo_ids`,`app`.`created_at`,`app`.`updated_at`,`app`.`customer_id`) VALUES (?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`name`=VALUES(`name`),`description`=VALUES(`description`),`active`=VALUES(`active`),`repo_ids`=VALUES(`repo_ids`),`created_at`=VALUES(`created_at`),`updated_at`=VALUES(`updated_at`),`customer_id`=VALUES(`customer_id`)"
	}
	r, err := db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.Description),
		orm.ToSQLBool(t.Active),
		orm.ToSQLString(t.RepoIds),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
		orm.ToSQLString(t.CustomerID),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBUpsertTx will upsert the App record in the database using the provided transaction
func (t *App) DBUpsertTx(ctx context.Context, tx *sql.Tx, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `app` (`app`.`id`,`app`.`checksum`,`app`.`name`,`app`.`description`,`app`.`active`,`app`.`repo_ids`,`app`.`created_at`,`app`.`updated_at`,`app`.`customer_id`) VALUES (?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `app` (`app`.`id`,`app`.`checksum`,`app`.`name`,`app`.`description`,`app`.`active`,`app`.`repo_ids`,`app`.`created_at`,`app`.`updated_at`,`app`.`customer_id`) VALUES (?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`name`=VALUES(`name`),`description`=VALUES(`description`),`active`=VALUES(`active`),`repo_ids`=VALUES(`repo_ids`),`created_at`=VALUES(`created_at`),`updated_at`=VALUES(`updated_at`),`customer_id`=VALUES(`customer_id`)"
	}
	r, err := tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.Description),
		orm.ToSQLBool(t.Active),
		orm.ToSQLString(t.RepoIds),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
		orm.ToSQLString(t.CustomerID),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBFindOne will find a App record in the database with the primary key
func (t *App) DBFindOne(ctx context.Context, db *sql.DB, value string) (bool, error) {
	q := "SELECT `app`.`id`,`app`.`checksum`,`app`.`name`,`app`.`description`,`app`.`active`,`app`.`repo_ids`,`app`.`created_at`,`app`.`updated_at`,`app`.`customer_id` FROM `app` WHERE `id` = ? LIMIT 1"
	row := db.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _Name sql.NullString
	var _Description sql.NullString
	var _Active sql.NullBool
	var _RepoIds sql.NullString
	var _CreatedAt sql.NullInt64
	var _UpdatedAt sql.NullInt64
	var _CustomerID sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_Name,
		&_Description,
		&_Active,
		&_RepoIds,
		&_CreatedAt,
		&_UpdatedAt,
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
	if _Active.Valid {
		t.SetActive(_Active.Bool)
	}
	if _RepoIds.Valid {
		t.SetRepoIds(_RepoIds.String)
	}
	if _CreatedAt.Valid {
		t.SetCreatedAt(_CreatedAt.Int64)
	}
	if _UpdatedAt.Valid {
		t.SetUpdatedAt(_UpdatedAt.Int64)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	return true, nil
}

// DBFindOneTx will find a App record in the database with the primary key using the provided transaction
func (t *App) DBFindOneTx(ctx context.Context, tx *sql.Tx, value string) (bool, error) {
	q := "SELECT `app`.`id`,`app`.`checksum`,`app`.`name`,`app`.`description`,`app`.`active`,`app`.`repo_ids`,`app`.`created_at`,`app`.`updated_at`,`app`.`customer_id` FROM `app` WHERE `id` = ? LIMIT 1"
	row := tx.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _Name sql.NullString
	var _Description sql.NullString
	var _Active sql.NullBool
	var _RepoIds sql.NullString
	var _CreatedAt sql.NullInt64
	var _UpdatedAt sql.NullInt64
	var _CustomerID sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_Name,
		&_Description,
		&_Active,
		&_RepoIds,
		&_CreatedAt,
		&_UpdatedAt,
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
	if _Active.Valid {
		t.SetActive(_Active.Bool)
	}
	if _RepoIds.Valid {
		t.SetRepoIds(_RepoIds.String)
	}
	if _CreatedAt.Valid {
		t.SetCreatedAt(_CreatedAt.Int64)
	}
	if _UpdatedAt.Valid {
		t.SetUpdatedAt(_UpdatedAt.Int64)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	return true, nil
}

// FindApps will find a App record in the database with the provided parameters
func FindApps(ctx context.Context, db *sql.DB, _params ...interface{}) ([]*App, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("name"),
		orm.Column("description"),
		orm.Column("active"),
		orm.Column("repo_ids"),
		orm.Column("created_at"),
		orm.Column("updated_at"),
		orm.Column("customer_id"),
		orm.Table(AppTableName),
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
	results := make([]*App, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _Name sql.NullString
		var _Description sql.NullString
		var _Active sql.NullBool
		var _RepoIds sql.NullString
		var _CreatedAt sql.NullInt64
		var _UpdatedAt sql.NullInt64
		var _CustomerID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_Name,
			&_Description,
			&_Active,
			&_RepoIds,
			&_CreatedAt,
			&_UpdatedAt,
			&_CustomerID,
		)
		if err != nil {
			return nil, err
		}
		t := &App{}
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
		if _Active.Valid {
			t.SetActive(_Active.Bool)
		}
		if _RepoIds.Valid {
			t.SetRepoIds(_RepoIds.String)
		}
		if _CreatedAt.Valid {
			t.SetCreatedAt(_CreatedAt.Int64)
		}
		if _UpdatedAt.Valid {
			t.SetUpdatedAt(_UpdatedAt.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindAppsTx will find a App record in the database with the provided parameters using the provided transaction
func FindAppsTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) ([]*App, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("name"),
		orm.Column("description"),
		orm.Column("active"),
		orm.Column("repo_ids"),
		orm.Column("created_at"),
		orm.Column("updated_at"),
		orm.Column("customer_id"),
		orm.Table(AppTableName),
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
	results := make([]*App, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _Name sql.NullString
		var _Description sql.NullString
		var _Active sql.NullBool
		var _RepoIds sql.NullString
		var _CreatedAt sql.NullInt64
		var _UpdatedAt sql.NullInt64
		var _CustomerID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_Name,
			&_Description,
			&_Active,
			&_RepoIds,
			&_CreatedAt,
			&_UpdatedAt,
			&_CustomerID,
		)
		if err != nil {
			return nil, err
		}
		t := &App{}
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
		if _Active.Valid {
			t.SetActive(_Active.Bool)
		}
		if _RepoIds.Valid {
			t.SetRepoIds(_RepoIds.String)
		}
		if _CreatedAt.Valid {
			t.SetCreatedAt(_CreatedAt.Int64)
		}
		if _UpdatedAt.Valid {
			t.SetUpdatedAt(_UpdatedAt.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// DBFind will find a App record in the database with the provided parameters
func (t *App) DBFind(ctx context.Context, db *sql.DB, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("name"),
		orm.Column("description"),
		orm.Column("active"),
		orm.Column("repo_ids"),
		orm.Column("created_at"),
		orm.Column("updated_at"),
		orm.Column("customer_id"),
		orm.Table(AppTableName),
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
	var _Active sql.NullBool
	var _RepoIds sql.NullString
	var _CreatedAt sql.NullInt64
	var _UpdatedAt sql.NullInt64
	var _CustomerID sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_Name,
		&_Description,
		&_Active,
		&_RepoIds,
		&_CreatedAt,
		&_UpdatedAt,
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
	if _Active.Valid {
		t.SetActive(_Active.Bool)
	}
	if _RepoIds.Valid {
		t.SetRepoIds(_RepoIds.String)
	}
	if _CreatedAt.Valid {
		t.SetCreatedAt(_CreatedAt.Int64)
	}
	if _UpdatedAt.Valid {
		t.SetUpdatedAt(_UpdatedAt.Int64)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	return true, nil
}

// DBFindTx will find a App record in the database with the provided parameters using the provided transaction
func (t *App) DBFindTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("name"),
		orm.Column("description"),
		orm.Column("active"),
		orm.Column("repo_ids"),
		orm.Column("created_at"),
		orm.Column("updated_at"),
		orm.Column("customer_id"),
		orm.Table(AppTableName),
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
	var _Active sql.NullBool
	var _RepoIds sql.NullString
	var _CreatedAt sql.NullInt64
	var _UpdatedAt sql.NullInt64
	var _CustomerID sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_Name,
		&_Description,
		&_Active,
		&_RepoIds,
		&_CreatedAt,
		&_UpdatedAt,
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
	if _Active.Valid {
		t.SetActive(_Active.Bool)
	}
	if _RepoIds.Valid {
		t.SetRepoIds(_RepoIds.String)
	}
	if _CreatedAt.Valid {
		t.SetCreatedAt(_CreatedAt.Int64)
	}
	if _UpdatedAt.Valid {
		t.SetUpdatedAt(_UpdatedAt.Int64)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	return true, nil
}

// CountApps will find the count of App records in the database
func CountApps(ctx context.Context, db *sql.DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(AppTableName),
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

// CountAppsTx will find the count of App records in the database using the provided transaction
func CountAppsTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(AppTableName),
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

// DBCount will find the count of App records in the database
func (t *App) DBCount(ctx context.Context, db *sql.DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(AppTableName),
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

// DBCountTx will find the count of App records in the database using the provided transaction
func (t *App) DBCountTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(AppTableName),
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

// DBExists will return true if the App record exists in the database
func (t *App) DBExists(ctx context.Context, db *sql.DB) (bool, error) {
	q := "SELECT `id` FROM `app` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := db.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// DBExistsTx will return true if the App record exists in the database using the provided transaction
func (t *App) DBExistsTx(ctx context.Context, tx *sql.Tx) (bool, error) {
	q := "SELECT `id` FROM `app` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := tx.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// PrimaryKeyColumn returns the column name for the primary key
func (t *App) PrimaryKeyColumn() string {
	return AppColumnID
}

// PrimaryKeyColumnType returns the primary key column Go type as a string
func (t *App) PrimaryKeyColumnType() string {
	return "string"
}

// PrimaryKey returns the primary key column value
func (t *App) PrimaryKey() interface{} {
	return t.ID
}
