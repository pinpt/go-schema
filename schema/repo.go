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

var _ Model = (*Repo)(nil)
var _ CSVWriter = (*Repo)(nil)
var _ JSONWriter = (*Repo)(nil)
var _ Checksum = (*Repo)(nil)

// RepoTableName is the name of the table in SQL
const RepoTableName = "repo"

var RepoColumns = []string{
	"id",
	"checksum",
	"name",
	"url",
	"active",
	"fork",
	"parent_id",
	"description",
	"created_at",
	"updated_at",
	"customer_id",
	"ref_type",
	"ref_id",
	"metadata",
}

// Repo table
type Repo struct {
	Active      bool    `json:"active"`
	Checksum    *string `json:"checksum,omitempty"`
	CreatedAt   int64   `json:"created_at"`
	CustomerID  string  `json:"customer_id"`
	Description *string `json:"description,omitempty"`
	Fork        bool    `json:"fork"`
	ID          string  `json:"id"`
	Metadata    *string `json:"metadata,omitempty"`
	Name        string  `json:"name"`
	ParentID    *string `json:"parent_id,omitempty"`
	RefID       string  `json:"ref_id"`
	RefType     string  `json:"ref_type"`
	UpdatedAt   *int64  `json:"updated_at,omitempty"`
	URL         string  `json:"url"`
}

// TableName returns the SQL table name for Repo and satifies the Model interface
func (t *Repo) TableName() string {
	return RepoTableName
}

// ToCSV will serialize the Repo instance to a CSV compatible array of strings
func (t *Repo) ToCSV() []string {
	return []string{
		t.ID,
		t.CalculateChecksum(),
		t.Name,
		t.URL,
		toCSVBool(t.Active),
		toCSVBool(t.Fork),
		toCSVString(t.ParentID),
		toCSVString(t.Description),
		toCSVString(t.CreatedAt),
		toCSVString(t.UpdatedAt),
		t.CustomerID,
		t.RefType,
		t.RefID,
		toCSVString(t.Metadata),
	}
}

// WriteCSV will serialize the Repo instance to the writer as CSV and satisfies the CSVWriter interface
func (t *Repo) WriteCSV(w *csv.Writer) error {
	return w.Write(t.ToCSV())
}

// WriteJSON will serialize the Repo instance to the writer as JSON and satisfies the JSONWriter interface
func (t *Repo) WriteJSON(w io.Writer, indent ...bool) error {
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

// NewRepoReader creates a JSON reader which can read in Repo objects serialized as JSON either as an array, single object or json new lines
// and writes each Repo to the channel provided
func NewRepoReader(r io.Reader, ch chan<- Repo) error {
	return orm.Deserialize(r, func(buf json.RawMessage) error {
		dec := json.NewDecoder(bytes.NewBuffer(buf))
		e := Repo{}
		if err := dec.Decode(&e); err != nil {
			return err
		}
		ch <- e
		return nil
	})
}

// NewCSVRepoReaderDir will read the reader as CSV and emit each record to the channel provided
func NewCSVRepoReader(r io.Reader, ch chan<- Repo) error {
	cr := csv.NewReader(r)
	for {
		record, err := cr.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		ch <- Repo{
			ID:          record[0],
			Checksum:    fromStringPointer(record[1]),
			Name:        record[2],
			URL:         record[3],
			Active:      fromCSVBool(record[4]),
			Fork:        fromCSVBool(record[5]),
			ParentID:    fromStringPointer(record[6]),
			Description: fromStringPointer(record[7]),
			CreatedAt:   fromCSVInt64(record[8]),
			UpdatedAt:   fromCSVInt64Pointer(record[9]),
			CustomerID:  record[10],
			RefType:     record[11],
			RefID:       record[12],
			Metadata:    fromStringPointer(record[13]),
		}
	}
	return nil
}

// NewCSVRepoReaderFile will read the file as a CSV and emit each record to the channel provided
func NewCSVRepoReaderFile(fp string, ch chan<- Repo) error {
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
	return NewCSVRepoReader(fc, ch)
}

// NewCSVRepoReaderDir will read the repo.csv.gz file as a CSV and emit each record to the channel provided
func NewCSVRepoReaderDir(dir string, ch chan<- Repo) error {
	return NewCSVRepoReaderFile(filepath.Join(dir, "repo.csv.gz"), ch)
}

// RepoCSVDeduper is a function callback which takes the existing value (a) and the new value (b)
// and the return value should be the one to use (or a new one, if applicable). return nil
// to skip processing of this record
type RepoCSVDeduper func(a Repo, b Repo) *Repo

// RepoCSVDedupeDisabled is set on whether the CSV writer should de-dupe values by key
var RepoCSVDedupeDisabled bool

// NewRepoCSVWriterSize creates a batch writer that will write each Repo into a CSV file
// this method will automatically de-duplicate entries using the primary key. if the checksum
// for a newer item with the same primary key doesn't match a previously sent item, the newer
// one will be used
func NewRepoCSVWriterSize(w io.Writer, size int, dedupers ...RepoCSVDeduper) (chan Repo, chan bool, error) {
	cw := csv.NewWriter(w)
	ch := make(chan Repo, size)
	done := make(chan bool)
	go func() {
		defer func() { done <- true }()
		dodedupe := !RepoCSVDedupeDisabled
		var kv map[string]*Repo
		var deduper RepoCSVDeduper
		if dedupers != nil && len(dedupers) > 0 {
			deduper = dedupers[0]
			dodedupe = true
		}
		if dodedupe {
			kv = make(map[string]*Repo)
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

// RepoCSVDefaultSize is the default channel buffer size if not provided
var RepoCSVDefaultSize = 100

// NewRepoCSVWriter creates a batch writer that will write each Repo into a CSV file
func NewRepoCSVWriter(w io.Writer, dedupers ...RepoCSVDeduper) (chan Repo, chan bool, error) {
	return NewRepoCSVWriterSize(w, RepoCSVDefaultSize, dedupers...)
}

// NewRepoCSVWriterDir creates a batch writer that will write each Repo into a CSV file named repo.csv.gz in dir
func NewRepoCSVWriterDir(dir string, dedupers ...RepoCSVDeduper) (chan Repo, chan bool, error) {
	return NewRepoCSVWriterFile(filepath.Join(dir, "repo.csv.gz"), dedupers...)
}

// NewRepoCSVWriterFile creates a batch writer that will write each Repo into a CSV file
func NewRepoCSVWriterFile(fn string, dedupers ...RepoCSVDeduper) (chan Repo, chan bool, error) {
	f, err := os.Create(fn)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening CSV file %s. %v", fn, err)
	}
	var fc io.WriteCloser = f
	if filepath.Ext(fn) == ".gz" {
		w, _ := gzip.NewWriterLevel(f, gzip.BestCompression)
		fc = w
	}
	ch, done, err := NewRepoCSVWriter(fc, dedupers...)
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

type RepoDBAction func(ctx context.Context, db DB, record Repo) error

// NewRepoDBWriterSize creates a DB writer that will write each issue into the DB
func NewRepoDBWriterSize(ctx context.Context, db DB, errors chan<- error, size int, actions ...RepoDBAction) (chan Repo, chan bool, error) {
	ch := make(chan Repo, size)
	done := make(chan bool)
	var action RepoDBAction
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

// NewRepoDBWriter creates a DB writer that will write each issue into the DB
func NewRepoDBWriter(ctx context.Context, db DB, errors chan<- error, actions ...RepoDBAction) (chan Repo, chan bool, error) {
	return NewRepoDBWriterSize(ctx, db, errors, 100, actions...)
}

// RepoColumnID is the ID SQL column name for the Repo table
const RepoColumnID = "id"

// RepoEscapedColumnID is the escaped ID SQL column name for the Repo table
const RepoEscapedColumnID = "`id`"

// RepoColumnChecksum is the Checksum SQL column name for the Repo table
const RepoColumnChecksum = "checksum"

// RepoEscapedColumnChecksum is the escaped Checksum SQL column name for the Repo table
const RepoEscapedColumnChecksum = "`checksum`"

// RepoColumnName is the Name SQL column name for the Repo table
const RepoColumnName = "name"

// RepoEscapedColumnName is the escaped Name SQL column name for the Repo table
const RepoEscapedColumnName = "`name`"

// RepoColumnURL is the URL SQL column name for the Repo table
const RepoColumnURL = "url"

// RepoEscapedColumnURL is the escaped URL SQL column name for the Repo table
const RepoEscapedColumnURL = "`url`"

// RepoColumnActive is the Active SQL column name for the Repo table
const RepoColumnActive = "active"

// RepoEscapedColumnActive is the escaped Active SQL column name for the Repo table
const RepoEscapedColumnActive = "`active`"

// RepoColumnFork is the Fork SQL column name for the Repo table
const RepoColumnFork = "fork"

// RepoEscapedColumnFork is the escaped Fork SQL column name for the Repo table
const RepoEscapedColumnFork = "`fork`"

// RepoColumnParentID is the ParentID SQL column name for the Repo table
const RepoColumnParentID = "parent_id"

// RepoEscapedColumnParentID is the escaped ParentID SQL column name for the Repo table
const RepoEscapedColumnParentID = "`parent_id`"

// RepoColumnDescription is the Description SQL column name for the Repo table
const RepoColumnDescription = "description"

// RepoEscapedColumnDescription is the escaped Description SQL column name for the Repo table
const RepoEscapedColumnDescription = "`description`"

// RepoColumnCreatedAt is the CreatedAt SQL column name for the Repo table
const RepoColumnCreatedAt = "created_at"

// RepoEscapedColumnCreatedAt is the escaped CreatedAt SQL column name for the Repo table
const RepoEscapedColumnCreatedAt = "`created_at`"

// RepoColumnUpdatedAt is the UpdatedAt SQL column name for the Repo table
const RepoColumnUpdatedAt = "updated_at"

// RepoEscapedColumnUpdatedAt is the escaped UpdatedAt SQL column name for the Repo table
const RepoEscapedColumnUpdatedAt = "`updated_at`"

// RepoColumnCustomerID is the CustomerID SQL column name for the Repo table
const RepoColumnCustomerID = "customer_id"

// RepoEscapedColumnCustomerID is the escaped CustomerID SQL column name for the Repo table
const RepoEscapedColumnCustomerID = "`customer_id`"

// RepoColumnRefType is the RefType SQL column name for the Repo table
const RepoColumnRefType = "ref_type"

// RepoEscapedColumnRefType is the escaped RefType SQL column name for the Repo table
const RepoEscapedColumnRefType = "`ref_type`"

// RepoColumnRefID is the RefID SQL column name for the Repo table
const RepoColumnRefID = "ref_id"

// RepoEscapedColumnRefID is the escaped RefID SQL column name for the Repo table
const RepoEscapedColumnRefID = "`ref_id`"

// RepoColumnMetadata is the Metadata SQL column name for the Repo table
const RepoColumnMetadata = "metadata"

// RepoEscapedColumnMetadata is the escaped Metadata SQL column name for the Repo table
const RepoEscapedColumnMetadata = "`metadata`"

// GetID will return the Repo ID value
func (t *Repo) GetID() string {
	return t.ID
}

// SetID will set the Repo ID value
func (t *Repo) SetID(v string) {
	t.ID = v
}

// FindRepoByID will find a Repo by ID
func FindRepoByID(ctx context.Context, db DB, value string) (*Repo, error) {
	q := "SELECT `repo`.`id`,`repo`.`checksum`,`repo`.`name`,`repo`.`url`,`repo`.`active`,`repo`.`fork`,`repo`.`parent_id`,`repo`.`description`,`repo`.`created_at`,`repo`.`updated_at`,`repo`.`customer_id`,`repo`.`ref_type`,`repo`.`ref_id`,`repo`.`metadata` FROM `repo` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _Name sql.NullString
	var _URL sql.NullString
	var _Active sql.NullBool
	var _Fork sql.NullBool
	var _ParentID sql.NullString
	var _Description sql.NullString
	var _CreatedAt sql.NullInt64
	var _UpdatedAt sql.NullInt64
	var _CustomerID sql.NullString
	var _RefType sql.NullString
	var _RefID sql.NullString
	var _Metadata sql.NullString
	err := db.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_Name,
		&_URL,
		&_Active,
		&_Fork,
		&_ParentID,
		&_Description,
		&_CreatedAt,
		&_UpdatedAt,
		&_CustomerID,
		&_RefType,
		&_RefID,
		&_Metadata,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &Repo{}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _Checksum.Valid {
		t.SetChecksum(_Checksum.String)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _URL.Valid {
		t.SetURL(_URL.String)
	}
	if _Active.Valid {
		t.SetActive(_Active.Bool)
	}
	if _Fork.Valid {
		t.SetFork(_Fork.Bool)
	}
	if _ParentID.Valid {
		t.SetParentID(_ParentID.String)
	}
	if _Description.Valid {
		t.SetDescription(_Description.String)
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
	if _RefType.Valid {
		t.SetRefType(_RefType.String)
	}
	if _RefID.Valid {
		t.SetRefID(_RefID.String)
	}
	if _Metadata.Valid {
		t.SetMetadata(_Metadata.String)
	}
	return t, nil
}

// FindRepoByIDTx will find a Repo by ID using the provided transaction
func FindRepoByIDTx(ctx context.Context, tx Tx, value string) (*Repo, error) {
	q := "SELECT `repo`.`id`,`repo`.`checksum`,`repo`.`name`,`repo`.`url`,`repo`.`active`,`repo`.`fork`,`repo`.`parent_id`,`repo`.`description`,`repo`.`created_at`,`repo`.`updated_at`,`repo`.`customer_id`,`repo`.`ref_type`,`repo`.`ref_id`,`repo`.`metadata` FROM `repo` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _Name sql.NullString
	var _URL sql.NullString
	var _Active sql.NullBool
	var _Fork sql.NullBool
	var _ParentID sql.NullString
	var _Description sql.NullString
	var _CreatedAt sql.NullInt64
	var _UpdatedAt sql.NullInt64
	var _CustomerID sql.NullString
	var _RefType sql.NullString
	var _RefID sql.NullString
	var _Metadata sql.NullString
	err := tx.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_Name,
		&_URL,
		&_Active,
		&_Fork,
		&_ParentID,
		&_Description,
		&_CreatedAt,
		&_UpdatedAt,
		&_CustomerID,
		&_RefType,
		&_RefID,
		&_Metadata,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &Repo{}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _Checksum.Valid {
		t.SetChecksum(_Checksum.String)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _URL.Valid {
		t.SetURL(_URL.String)
	}
	if _Active.Valid {
		t.SetActive(_Active.Bool)
	}
	if _Fork.Valid {
		t.SetFork(_Fork.Bool)
	}
	if _ParentID.Valid {
		t.SetParentID(_ParentID.String)
	}
	if _Description.Valid {
		t.SetDescription(_Description.String)
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
	if _RefType.Valid {
		t.SetRefType(_RefType.String)
	}
	if _RefID.Valid {
		t.SetRefID(_RefID.String)
	}
	if _Metadata.Valid {
		t.SetMetadata(_Metadata.String)
	}
	return t, nil
}

// GetChecksum will return the Repo Checksum value
func (t *Repo) GetChecksum() string {
	if t.Checksum == nil {
		return ""
	}
	return *t.Checksum
}

// SetChecksum will set the Repo Checksum value
func (t *Repo) SetChecksum(v string) {
	t.Checksum = &v
}

// GetName will return the Repo Name value
func (t *Repo) GetName() string {
	return t.Name
}

// SetName will set the Repo Name value
func (t *Repo) SetName(v string) {
	t.Name = v
}

// FindReposByName will find all Repos by the Name value
func FindReposByName(ctx context.Context, db DB, value string) ([]*Repo, error) {
	q := "SELECT `repo`.`id`,`repo`.`checksum`,`repo`.`name`,`repo`.`url`,`repo`.`active`,`repo`.`fork`,`repo`.`parent_id`,`repo`.`description`,`repo`.`created_at`,`repo`.`updated_at`,`repo`.`customer_id`,`repo`.`ref_type`,`repo`.`ref_id`,`repo`.`metadata` FROM `repo` WHERE `name` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*Repo, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _Name sql.NullString
		var _URL sql.NullString
		var _Active sql.NullBool
		var _Fork sql.NullBool
		var _ParentID sql.NullString
		var _Description sql.NullString
		var _CreatedAt sql.NullInt64
		var _UpdatedAt sql.NullInt64
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefID sql.NullString
		var _Metadata sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_Name,
			&_URL,
			&_Active,
			&_Fork,
			&_ParentID,
			&_Description,
			&_CreatedAt,
			&_UpdatedAt,
			&_CustomerID,
			&_RefType,
			&_RefID,
			&_Metadata,
		)
		if err != nil {
			return nil, err
		}
		t := &Repo{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _URL.Valid {
			t.SetURL(_URL.String)
		}
		if _Active.Valid {
			t.SetActive(_Active.Bool)
		}
		if _Fork.Valid {
			t.SetFork(_Fork.Bool)
		}
		if _ParentID.Valid {
			t.SetParentID(_ParentID.String)
		}
		if _Description.Valid {
			t.SetDescription(_Description.String)
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
		if _RefType.Valid {
			t.SetRefType(_RefType.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		if _Metadata.Valid {
			t.SetMetadata(_Metadata.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindReposByNameTx will find all Repos by the Name value using the provided transaction
func FindReposByNameTx(ctx context.Context, tx Tx, value string) ([]*Repo, error) {
	q := "SELECT `repo`.`id`,`repo`.`checksum`,`repo`.`name`,`repo`.`url`,`repo`.`active`,`repo`.`fork`,`repo`.`parent_id`,`repo`.`description`,`repo`.`created_at`,`repo`.`updated_at`,`repo`.`customer_id`,`repo`.`ref_type`,`repo`.`ref_id`,`repo`.`metadata` FROM `repo` WHERE `name` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*Repo, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _Name sql.NullString
		var _URL sql.NullString
		var _Active sql.NullBool
		var _Fork sql.NullBool
		var _ParentID sql.NullString
		var _Description sql.NullString
		var _CreatedAt sql.NullInt64
		var _UpdatedAt sql.NullInt64
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefID sql.NullString
		var _Metadata sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_Name,
			&_URL,
			&_Active,
			&_Fork,
			&_ParentID,
			&_Description,
			&_CreatedAt,
			&_UpdatedAt,
			&_CustomerID,
			&_RefType,
			&_RefID,
			&_Metadata,
		)
		if err != nil {
			return nil, err
		}
		t := &Repo{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _URL.Valid {
			t.SetURL(_URL.String)
		}
		if _Active.Valid {
			t.SetActive(_Active.Bool)
		}
		if _Fork.Valid {
			t.SetFork(_Fork.Bool)
		}
		if _ParentID.Valid {
			t.SetParentID(_ParentID.String)
		}
		if _Description.Valid {
			t.SetDescription(_Description.String)
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
		if _RefType.Valid {
			t.SetRefType(_RefType.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		if _Metadata.Valid {
			t.SetMetadata(_Metadata.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetURL will return the Repo URL value
func (t *Repo) GetURL() string {
	return t.URL
}

// SetURL will set the Repo URL value
func (t *Repo) SetURL(v string) {
	t.URL = v
}

// GetActive will return the Repo Active value
func (t *Repo) GetActive() bool {
	return t.Active
}

// SetActive will set the Repo Active value
func (t *Repo) SetActive(v bool) {
	t.Active = v
}

// GetFork will return the Repo Fork value
func (t *Repo) GetFork() bool {
	return t.Fork
}

// SetFork will set the Repo Fork value
func (t *Repo) SetFork(v bool) {
	t.Fork = v
}

// GetParentID will return the Repo ParentID value
func (t *Repo) GetParentID() string {
	if t.ParentID == nil {
		return ""
	}
	return *t.ParentID
}

// SetParentID will set the Repo ParentID value
func (t *Repo) SetParentID(v string) {
	t.ParentID = &v
}

// GetDescription will return the Repo Description value
func (t *Repo) GetDescription() string {
	if t.Description == nil {
		return ""
	}
	return *t.Description
}

// SetDescription will set the Repo Description value
func (t *Repo) SetDescription(v string) {
	t.Description = &v
}

// GetCreatedAt will return the Repo CreatedAt value
func (t *Repo) GetCreatedAt() int64 {
	return t.CreatedAt
}

// SetCreatedAt will set the Repo CreatedAt value
func (t *Repo) SetCreatedAt(v int64) {
	t.CreatedAt = v
}

// GetUpdatedAt will return the Repo UpdatedAt value
func (t *Repo) GetUpdatedAt() int64 {
	if t.UpdatedAt == nil {
		return int64(0)
	}
	return *t.UpdatedAt
}

// SetUpdatedAt will set the Repo UpdatedAt value
func (t *Repo) SetUpdatedAt(v int64) {
	t.UpdatedAt = &v
}

// GetCustomerID will return the Repo CustomerID value
func (t *Repo) GetCustomerID() string {
	return t.CustomerID
}

// SetCustomerID will set the Repo CustomerID value
func (t *Repo) SetCustomerID(v string) {
	t.CustomerID = v
}

// FindReposByCustomerID will find all Repos by the CustomerID value
func FindReposByCustomerID(ctx context.Context, db DB, value string) ([]*Repo, error) {
	q := "SELECT `repo`.`id`,`repo`.`checksum`,`repo`.`name`,`repo`.`url`,`repo`.`active`,`repo`.`fork`,`repo`.`parent_id`,`repo`.`description`,`repo`.`created_at`,`repo`.`updated_at`,`repo`.`customer_id`,`repo`.`ref_type`,`repo`.`ref_id`,`repo`.`metadata` FROM `repo` WHERE `customer_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*Repo, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _Name sql.NullString
		var _URL sql.NullString
		var _Active sql.NullBool
		var _Fork sql.NullBool
		var _ParentID sql.NullString
		var _Description sql.NullString
		var _CreatedAt sql.NullInt64
		var _UpdatedAt sql.NullInt64
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefID sql.NullString
		var _Metadata sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_Name,
			&_URL,
			&_Active,
			&_Fork,
			&_ParentID,
			&_Description,
			&_CreatedAt,
			&_UpdatedAt,
			&_CustomerID,
			&_RefType,
			&_RefID,
			&_Metadata,
		)
		if err != nil {
			return nil, err
		}
		t := &Repo{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _URL.Valid {
			t.SetURL(_URL.String)
		}
		if _Active.Valid {
			t.SetActive(_Active.Bool)
		}
		if _Fork.Valid {
			t.SetFork(_Fork.Bool)
		}
		if _ParentID.Valid {
			t.SetParentID(_ParentID.String)
		}
		if _Description.Valid {
			t.SetDescription(_Description.String)
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
		if _RefType.Valid {
			t.SetRefType(_RefType.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		if _Metadata.Valid {
			t.SetMetadata(_Metadata.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindReposByCustomerIDTx will find all Repos by the CustomerID value using the provided transaction
func FindReposByCustomerIDTx(ctx context.Context, tx Tx, value string) ([]*Repo, error) {
	q := "SELECT `repo`.`id`,`repo`.`checksum`,`repo`.`name`,`repo`.`url`,`repo`.`active`,`repo`.`fork`,`repo`.`parent_id`,`repo`.`description`,`repo`.`created_at`,`repo`.`updated_at`,`repo`.`customer_id`,`repo`.`ref_type`,`repo`.`ref_id`,`repo`.`metadata` FROM `repo` WHERE `customer_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*Repo, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _Name sql.NullString
		var _URL sql.NullString
		var _Active sql.NullBool
		var _Fork sql.NullBool
		var _ParentID sql.NullString
		var _Description sql.NullString
		var _CreatedAt sql.NullInt64
		var _UpdatedAt sql.NullInt64
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefID sql.NullString
		var _Metadata sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_Name,
			&_URL,
			&_Active,
			&_Fork,
			&_ParentID,
			&_Description,
			&_CreatedAt,
			&_UpdatedAt,
			&_CustomerID,
			&_RefType,
			&_RefID,
			&_Metadata,
		)
		if err != nil {
			return nil, err
		}
		t := &Repo{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _URL.Valid {
			t.SetURL(_URL.String)
		}
		if _Active.Valid {
			t.SetActive(_Active.Bool)
		}
		if _Fork.Valid {
			t.SetFork(_Fork.Bool)
		}
		if _ParentID.Valid {
			t.SetParentID(_ParentID.String)
		}
		if _Description.Valid {
			t.SetDescription(_Description.String)
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
		if _RefType.Valid {
			t.SetRefType(_RefType.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		if _Metadata.Valid {
			t.SetMetadata(_Metadata.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetRefType will return the Repo RefType value
func (t *Repo) GetRefType() string {
	return t.RefType
}

// SetRefType will set the Repo RefType value
func (t *Repo) SetRefType(v string) {
	t.RefType = v
}

// FindReposByRefType will find all Repos by the RefType value
func FindReposByRefType(ctx context.Context, db DB, value string) ([]*Repo, error) {
	q := "SELECT `repo`.`id`,`repo`.`checksum`,`repo`.`name`,`repo`.`url`,`repo`.`active`,`repo`.`fork`,`repo`.`parent_id`,`repo`.`description`,`repo`.`created_at`,`repo`.`updated_at`,`repo`.`customer_id`,`repo`.`ref_type`,`repo`.`ref_id`,`repo`.`metadata` FROM `repo` WHERE `ref_type` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*Repo, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _Name sql.NullString
		var _URL sql.NullString
		var _Active sql.NullBool
		var _Fork sql.NullBool
		var _ParentID sql.NullString
		var _Description sql.NullString
		var _CreatedAt sql.NullInt64
		var _UpdatedAt sql.NullInt64
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefID sql.NullString
		var _Metadata sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_Name,
			&_URL,
			&_Active,
			&_Fork,
			&_ParentID,
			&_Description,
			&_CreatedAt,
			&_UpdatedAt,
			&_CustomerID,
			&_RefType,
			&_RefID,
			&_Metadata,
		)
		if err != nil {
			return nil, err
		}
		t := &Repo{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _URL.Valid {
			t.SetURL(_URL.String)
		}
		if _Active.Valid {
			t.SetActive(_Active.Bool)
		}
		if _Fork.Valid {
			t.SetFork(_Fork.Bool)
		}
		if _ParentID.Valid {
			t.SetParentID(_ParentID.String)
		}
		if _Description.Valid {
			t.SetDescription(_Description.String)
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
		if _RefType.Valid {
			t.SetRefType(_RefType.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		if _Metadata.Valid {
			t.SetMetadata(_Metadata.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindReposByRefTypeTx will find all Repos by the RefType value using the provided transaction
func FindReposByRefTypeTx(ctx context.Context, tx Tx, value string) ([]*Repo, error) {
	q := "SELECT `repo`.`id`,`repo`.`checksum`,`repo`.`name`,`repo`.`url`,`repo`.`active`,`repo`.`fork`,`repo`.`parent_id`,`repo`.`description`,`repo`.`created_at`,`repo`.`updated_at`,`repo`.`customer_id`,`repo`.`ref_type`,`repo`.`ref_id`,`repo`.`metadata` FROM `repo` WHERE `ref_type` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*Repo, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _Name sql.NullString
		var _URL sql.NullString
		var _Active sql.NullBool
		var _Fork sql.NullBool
		var _ParentID sql.NullString
		var _Description sql.NullString
		var _CreatedAt sql.NullInt64
		var _UpdatedAt sql.NullInt64
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefID sql.NullString
		var _Metadata sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_Name,
			&_URL,
			&_Active,
			&_Fork,
			&_ParentID,
			&_Description,
			&_CreatedAt,
			&_UpdatedAt,
			&_CustomerID,
			&_RefType,
			&_RefID,
			&_Metadata,
		)
		if err != nil {
			return nil, err
		}
		t := &Repo{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _URL.Valid {
			t.SetURL(_URL.String)
		}
		if _Active.Valid {
			t.SetActive(_Active.Bool)
		}
		if _Fork.Valid {
			t.SetFork(_Fork.Bool)
		}
		if _ParentID.Valid {
			t.SetParentID(_ParentID.String)
		}
		if _Description.Valid {
			t.SetDescription(_Description.String)
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
		if _RefType.Valid {
			t.SetRefType(_RefType.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		if _Metadata.Valid {
			t.SetMetadata(_Metadata.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetRefID will return the Repo RefID value
func (t *Repo) GetRefID() string {
	return t.RefID
}

// SetRefID will set the Repo RefID value
func (t *Repo) SetRefID(v string) {
	t.RefID = v
}

// FindReposByRefID will find all Repos by the RefID value
func FindReposByRefID(ctx context.Context, db DB, value string) ([]*Repo, error) {
	q := "SELECT `repo`.`id`,`repo`.`checksum`,`repo`.`name`,`repo`.`url`,`repo`.`active`,`repo`.`fork`,`repo`.`parent_id`,`repo`.`description`,`repo`.`created_at`,`repo`.`updated_at`,`repo`.`customer_id`,`repo`.`ref_type`,`repo`.`ref_id`,`repo`.`metadata` FROM `repo` WHERE `ref_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*Repo, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _Name sql.NullString
		var _URL sql.NullString
		var _Active sql.NullBool
		var _Fork sql.NullBool
		var _ParentID sql.NullString
		var _Description sql.NullString
		var _CreatedAt sql.NullInt64
		var _UpdatedAt sql.NullInt64
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefID sql.NullString
		var _Metadata sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_Name,
			&_URL,
			&_Active,
			&_Fork,
			&_ParentID,
			&_Description,
			&_CreatedAt,
			&_UpdatedAt,
			&_CustomerID,
			&_RefType,
			&_RefID,
			&_Metadata,
		)
		if err != nil {
			return nil, err
		}
		t := &Repo{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _URL.Valid {
			t.SetURL(_URL.String)
		}
		if _Active.Valid {
			t.SetActive(_Active.Bool)
		}
		if _Fork.Valid {
			t.SetFork(_Fork.Bool)
		}
		if _ParentID.Valid {
			t.SetParentID(_ParentID.String)
		}
		if _Description.Valid {
			t.SetDescription(_Description.String)
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
		if _RefType.Valid {
			t.SetRefType(_RefType.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		if _Metadata.Valid {
			t.SetMetadata(_Metadata.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindReposByRefIDTx will find all Repos by the RefID value using the provided transaction
func FindReposByRefIDTx(ctx context.Context, tx Tx, value string) ([]*Repo, error) {
	q := "SELECT `repo`.`id`,`repo`.`checksum`,`repo`.`name`,`repo`.`url`,`repo`.`active`,`repo`.`fork`,`repo`.`parent_id`,`repo`.`description`,`repo`.`created_at`,`repo`.`updated_at`,`repo`.`customer_id`,`repo`.`ref_type`,`repo`.`ref_id`,`repo`.`metadata` FROM `repo` WHERE `ref_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*Repo, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _Name sql.NullString
		var _URL sql.NullString
		var _Active sql.NullBool
		var _Fork sql.NullBool
		var _ParentID sql.NullString
		var _Description sql.NullString
		var _CreatedAt sql.NullInt64
		var _UpdatedAt sql.NullInt64
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefID sql.NullString
		var _Metadata sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_Name,
			&_URL,
			&_Active,
			&_Fork,
			&_ParentID,
			&_Description,
			&_CreatedAt,
			&_UpdatedAt,
			&_CustomerID,
			&_RefType,
			&_RefID,
			&_Metadata,
		)
		if err != nil {
			return nil, err
		}
		t := &Repo{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _URL.Valid {
			t.SetURL(_URL.String)
		}
		if _Active.Valid {
			t.SetActive(_Active.Bool)
		}
		if _Fork.Valid {
			t.SetFork(_Fork.Bool)
		}
		if _ParentID.Valid {
			t.SetParentID(_ParentID.String)
		}
		if _Description.Valid {
			t.SetDescription(_Description.String)
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
		if _RefType.Valid {
			t.SetRefType(_RefType.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		if _Metadata.Valid {
			t.SetMetadata(_Metadata.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetMetadata will return the Repo Metadata value
func (t *Repo) GetMetadata() string {
	if t.Metadata == nil {
		return ""
	}
	return *t.Metadata
}

// SetMetadata will set the Repo Metadata value
func (t *Repo) SetMetadata(v string) {
	t.Metadata = &v
}

func (t *Repo) toTimestamp(value time.Time) *timestamp.Timestamp {
	ts, _ := ptypes.TimestampProto(value)
	return ts
}

// DBCreateRepoTable will create the Repo table
func DBCreateRepoTable(ctx context.Context, db DB) error {
	q := "CREATE TABLE `repo` (`id`VARCHAR(64) NOT NULL PRIMARY KEY,`checksum`CHAR(64),`name` VARCHAR(255) NOT NULL,`url` VARCHAR(255) NOT NULL,`active` BOOL NOT NULL DEFAULT true,`fork` BOOL NOT NULL DEFAULT false,`parent_id` VARCHAR(64),`description`TEXT,`created_at` BIGINT UNSIGNED NOT NULL,`updated_at` BIGINT UNSIGNED,`customer_id`VARCHAR(64) NOT NULL,`ref_type`VARCHAR(20) NOT NULL,`ref_id` VARCHAR(64) NOT NULL,`metadata`JSON,INDEX repo_name_index (`name`),INDEX repo_customer_id_index (`customer_id`),INDEX repo_ref_type_index (`ref_type`),INDEX repo_ref_id_index (`ref_id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBCreateRepoTableTx will create the Repo table using the provided transction
func DBCreateRepoTableTx(ctx context.Context, tx Tx) error {
	q := "CREATE TABLE `repo` (`id`VARCHAR(64) NOT NULL PRIMARY KEY,`checksum`CHAR(64),`name` VARCHAR(255) NOT NULL,`url` VARCHAR(255) NOT NULL,`active` BOOL NOT NULL DEFAULT true,`fork` BOOL NOT NULL DEFAULT false,`parent_id` VARCHAR(64),`description`TEXT,`created_at` BIGINT UNSIGNED NOT NULL,`updated_at` BIGINT UNSIGNED,`customer_id`VARCHAR(64) NOT NULL,`ref_type`VARCHAR(20) NOT NULL,`ref_id` VARCHAR(64) NOT NULL,`metadata`JSON,INDEX repo_name_index (`name`),INDEX repo_customer_id_index (`customer_id`),INDEX repo_ref_type_index (`ref_type`),INDEX repo_ref_id_index (`ref_id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// DBDropRepoTable will drop the Repo table
func DBDropRepoTable(ctx context.Context, db DB) error {
	q := "DROP TABLE IF EXISTS `repo`"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBDropRepoTableTx will drop the Repo table using the provided transaction
func DBDropRepoTableTx(ctx context.Context, tx Tx) error {
	q := "DROP TABLE IF EXISTS `repo`"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// CalculateChecksum will calculate a checksum of the SHA1 of all field values
func (t *Repo) CalculateChecksum() string {
	return orm.HashStrings(
		orm.ToString(t.ID),
		orm.ToString(t.Name),
		orm.ToString(t.URL),
		orm.ToString(t.Active),
		orm.ToString(t.Fork),
		orm.ToString(t.ParentID),
		orm.ToString(t.Description),
		orm.ToString(t.CreatedAt),
		orm.ToString(t.UpdatedAt),
		orm.ToString(t.CustomerID),
		orm.ToString(t.RefType),
		orm.ToString(t.RefID),
		orm.ToString(t.Metadata),
	)
}

// DBCreate will create a new Repo record in the database
func (t *Repo) DBCreate(ctx context.Context, db DB) (sql.Result, error) {
	q := "INSERT INTO `repo` (`repo`.`id`,`repo`.`checksum`,`repo`.`name`,`repo`.`url`,`repo`.`active`,`repo`.`fork`,`repo`.`parent_id`,`repo`.`description`,`repo`.`created_at`,`repo`.`updated_at`,`repo`.`customer_id`,`repo`.`ref_type`,`repo`.`ref_id`,`repo`.`metadata`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.URL),
		orm.ToSQLBool(t.Active),
		orm.ToSQLBool(t.Fork),
		orm.ToSQLString(t.ParentID),
		orm.ToSQLString(t.Description),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefType),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.Metadata),
	)
}

// DBCreateTx will create a new Repo record in the database using the provided transaction
func (t *Repo) DBCreateTx(ctx context.Context, tx Tx) (sql.Result, error) {
	q := "INSERT INTO `repo` (`repo`.`id`,`repo`.`checksum`,`repo`.`name`,`repo`.`url`,`repo`.`active`,`repo`.`fork`,`repo`.`parent_id`,`repo`.`description`,`repo`.`created_at`,`repo`.`updated_at`,`repo`.`customer_id`,`repo`.`ref_type`,`repo`.`ref_id`,`repo`.`metadata`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.URL),
		orm.ToSQLBool(t.Active),
		orm.ToSQLBool(t.Fork),
		orm.ToSQLString(t.ParentID),
		orm.ToSQLString(t.Description),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefType),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.Metadata),
	)
}

// DBCreateIgnoreDuplicate will upsert the Repo record in the database
func (t *Repo) DBCreateIgnoreDuplicate(ctx context.Context, db DB) (sql.Result, error) {
	q := "INSERT INTO `repo` (`repo`.`id`,`repo`.`checksum`,`repo`.`name`,`repo`.`url`,`repo`.`active`,`repo`.`fork`,`repo`.`parent_id`,`repo`.`description`,`repo`.`created_at`,`repo`.`updated_at`,`repo`.`customer_id`,`repo`.`ref_type`,`repo`.`ref_id`,`repo`.`metadata`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.URL),
		orm.ToSQLBool(t.Active),
		orm.ToSQLBool(t.Fork),
		orm.ToSQLString(t.ParentID),
		orm.ToSQLString(t.Description),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefType),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.Metadata),
	)
}

// DBCreateIgnoreDuplicateTx will upsert the Repo record in the database using the provided transaction
func (t *Repo) DBCreateIgnoreDuplicateTx(ctx context.Context, tx Tx) (sql.Result, error) {
	q := "INSERT INTO `repo` (`repo`.`id`,`repo`.`checksum`,`repo`.`name`,`repo`.`url`,`repo`.`active`,`repo`.`fork`,`repo`.`parent_id`,`repo`.`description`,`repo`.`created_at`,`repo`.`updated_at`,`repo`.`customer_id`,`repo`.`ref_type`,`repo`.`ref_id`,`repo`.`metadata`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.URL),
		orm.ToSQLBool(t.Active),
		orm.ToSQLBool(t.Fork),
		orm.ToSQLString(t.ParentID),
		orm.ToSQLString(t.Description),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefType),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.Metadata),
	)
}

// DeleteAllRepos deletes all Repo records in the database with optional filters
func DeleteAllRepos(ctx context.Context, db DB, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(RepoTableName),
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

// DeleteAllReposTx deletes all Repo records in the database with optional filters using the provided transaction
func DeleteAllReposTx(ctx context.Context, tx Tx, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(RepoTableName),
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

// DBDelete will delete this Repo record in the database
func (t *Repo) DBDelete(ctx context.Context, db DB) (bool, error) {
	q := "DELETE FROM `repo` WHERE `id` = ?"
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

// DBDeleteTx will delete this Repo record in the database using the provided transaction
func (t *Repo) DBDeleteTx(ctx context.Context, tx Tx) (bool, error) {
	q := "DELETE FROM `repo` WHERE `id` = ?"
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

// DBUpdate will update the Repo record in the database
func (t *Repo) DBUpdate(ctx context.Context, db DB) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `repo` SET `checksum`=?,`name`=?,`url`=?,`active`=?,`fork`=?,`parent_id`=?,`description`=?,`created_at`=?,`updated_at`=?,`customer_id`=?,`ref_type`=?,`ref_id`=?,`metadata`=? WHERE `id`=?"
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.URL),
		orm.ToSQLBool(t.Active),
		orm.ToSQLBool(t.Fork),
		orm.ToSQLString(t.ParentID),
		orm.ToSQLString(t.Description),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefType),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.Metadata),
		orm.ToSQLString(t.ID),
	)
}

// DBUpdateTx will update the Repo record in the database using the provided transaction
func (t *Repo) DBUpdateTx(ctx context.Context, tx Tx) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `repo` SET `checksum`=?,`name`=?,`url`=?,`active`=?,`fork`=?,`parent_id`=?,`description`=?,`created_at`=?,`updated_at`=?,`customer_id`=?,`ref_type`=?,`ref_id`=?,`metadata`=? WHERE `id`=?"
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.URL),
		orm.ToSQLBool(t.Active),
		orm.ToSQLBool(t.Fork),
		orm.ToSQLString(t.ParentID),
		orm.ToSQLString(t.Description),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefType),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.Metadata),
		orm.ToSQLString(t.ID),
	)
}

// DBUpsert will upsert the Repo record in the database
func (t *Repo) DBUpsert(ctx context.Context, db DB, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `repo` (`repo`.`id`,`repo`.`checksum`,`repo`.`name`,`repo`.`url`,`repo`.`active`,`repo`.`fork`,`repo`.`parent_id`,`repo`.`description`,`repo`.`created_at`,`repo`.`updated_at`,`repo`.`customer_id`,`repo`.`ref_type`,`repo`.`ref_id`,`repo`.`metadata`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `repo` (`repo`.`id`,`repo`.`checksum`,`repo`.`name`,`repo`.`url`,`repo`.`active`,`repo`.`fork`,`repo`.`parent_id`,`repo`.`description`,`repo`.`created_at`,`repo`.`updated_at`,`repo`.`customer_id`,`repo`.`ref_type`,`repo`.`ref_id`,`repo`.`metadata`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`name`=VALUES(`name`),`url`=VALUES(`url`),`active`=VALUES(`active`),`fork`=VALUES(`fork`),`parent_id`=VALUES(`parent_id`),`description`=VALUES(`description`),`created_at`=VALUES(`created_at`),`updated_at`=VALUES(`updated_at`),`customer_id`=VALUES(`customer_id`),`ref_type`=VALUES(`ref_type`),`ref_id`=VALUES(`ref_id`),`metadata`=VALUES(`metadata`)"
	}
	r, err := db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.URL),
		orm.ToSQLBool(t.Active),
		orm.ToSQLBool(t.Fork),
		orm.ToSQLString(t.ParentID),
		orm.ToSQLString(t.Description),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefType),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.Metadata),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBUpsertTx will upsert the Repo record in the database using the provided transaction
func (t *Repo) DBUpsertTx(ctx context.Context, tx Tx, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `repo` (`repo`.`id`,`repo`.`checksum`,`repo`.`name`,`repo`.`url`,`repo`.`active`,`repo`.`fork`,`repo`.`parent_id`,`repo`.`description`,`repo`.`created_at`,`repo`.`updated_at`,`repo`.`customer_id`,`repo`.`ref_type`,`repo`.`ref_id`,`repo`.`metadata`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `repo` (`repo`.`id`,`repo`.`checksum`,`repo`.`name`,`repo`.`url`,`repo`.`active`,`repo`.`fork`,`repo`.`parent_id`,`repo`.`description`,`repo`.`created_at`,`repo`.`updated_at`,`repo`.`customer_id`,`repo`.`ref_type`,`repo`.`ref_id`,`repo`.`metadata`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`name`=VALUES(`name`),`url`=VALUES(`url`),`active`=VALUES(`active`),`fork`=VALUES(`fork`),`parent_id`=VALUES(`parent_id`),`description`=VALUES(`description`),`created_at`=VALUES(`created_at`),`updated_at`=VALUES(`updated_at`),`customer_id`=VALUES(`customer_id`),`ref_type`=VALUES(`ref_type`),`ref_id`=VALUES(`ref_id`),`metadata`=VALUES(`metadata`)"
	}
	r, err := tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.URL),
		orm.ToSQLBool(t.Active),
		orm.ToSQLBool(t.Fork),
		orm.ToSQLString(t.ParentID),
		orm.ToSQLString(t.Description),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefType),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.Metadata),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBFindOne will find a Repo record in the database with the primary key
func (t *Repo) DBFindOne(ctx context.Context, db DB, value string) (bool, error) {
	q := "SELECT `repo`.`id`,`repo`.`checksum`,`repo`.`name`,`repo`.`url`,`repo`.`active`,`repo`.`fork`,`repo`.`parent_id`,`repo`.`description`,`repo`.`created_at`,`repo`.`updated_at`,`repo`.`customer_id`,`repo`.`ref_type`,`repo`.`ref_id`,`repo`.`metadata` FROM `repo` WHERE `id` = ? LIMIT 1"
	row := db.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _Name sql.NullString
	var _URL sql.NullString
	var _Active sql.NullBool
	var _Fork sql.NullBool
	var _ParentID sql.NullString
	var _Description sql.NullString
	var _CreatedAt sql.NullInt64
	var _UpdatedAt sql.NullInt64
	var _CustomerID sql.NullString
	var _RefType sql.NullString
	var _RefID sql.NullString
	var _Metadata sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_Name,
		&_URL,
		&_Active,
		&_Fork,
		&_ParentID,
		&_Description,
		&_CreatedAt,
		&_UpdatedAt,
		&_CustomerID,
		&_RefType,
		&_RefID,
		&_Metadata,
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
	if _URL.Valid {
		t.SetURL(_URL.String)
	}
	if _Active.Valid {
		t.SetActive(_Active.Bool)
	}
	if _Fork.Valid {
		t.SetFork(_Fork.Bool)
	}
	if _ParentID.Valid {
		t.SetParentID(_ParentID.String)
	}
	if _Description.Valid {
		t.SetDescription(_Description.String)
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
	if _RefType.Valid {
		t.SetRefType(_RefType.String)
	}
	if _RefID.Valid {
		t.SetRefID(_RefID.String)
	}
	if _Metadata.Valid {
		t.SetMetadata(_Metadata.String)
	}
	return true, nil
}

// DBFindOneTx will find a Repo record in the database with the primary key using the provided transaction
func (t *Repo) DBFindOneTx(ctx context.Context, tx Tx, value string) (bool, error) {
	q := "SELECT `repo`.`id`,`repo`.`checksum`,`repo`.`name`,`repo`.`url`,`repo`.`active`,`repo`.`fork`,`repo`.`parent_id`,`repo`.`description`,`repo`.`created_at`,`repo`.`updated_at`,`repo`.`customer_id`,`repo`.`ref_type`,`repo`.`ref_id`,`repo`.`metadata` FROM `repo` WHERE `id` = ? LIMIT 1"
	row := tx.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _Name sql.NullString
	var _URL sql.NullString
	var _Active sql.NullBool
	var _Fork sql.NullBool
	var _ParentID sql.NullString
	var _Description sql.NullString
	var _CreatedAt sql.NullInt64
	var _UpdatedAt sql.NullInt64
	var _CustomerID sql.NullString
	var _RefType sql.NullString
	var _RefID sql.NullString
	var _Metadata sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_Name,
		&_URL,
		&_Active,
		&_Fork,
		&_ParentID,
		&_Description,
		&_CreatedAt,
		&_UpdatedAt,
		&_CustomerID,
		&_RefType,
		&_RefID,
		&_Metadata,
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
	if _URL.Valid {
		t.SetURL(_URL.String)
	}
	if _Active.Valid {
		t.SetActive(_Active.Bool)
	}
	if _Fork.Valid {
		t.SetFork(_Fork.Bool)
	}
	if _ParentID.Valid {
		t.SetParentID(_ParentID.String)
	}
	if _Description.Valid {
		t.SetDescription(_Description.String)
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
	if _RefType.Valid {
		t.SetRefType(_RefType.String)
	}
	if _RefID.Valid {
		t.SetRefID(_RefID.String)
	}
	if _Metadata.Valid {
		t.SetMetadata(_Metadata.String)
	}
	return true, nil
}

// FindRepos will find a Repo record in the database with the provided parameters
func FindRepos(ctx context.Context, db DB, _params ...interface{}) ([]*Repo, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("name"),
		orm.Column("url"),
		orm.Column("active"),
		orm.Column("fork"),
		orm.Column("parent_id"),
		orm.Column("description"),
		orm.Column("created_at"),
		orm.Column("updated_at"),
		orm.Column("customer_id"),
		orm.Column("ref_type"),
		orm.Column("ref_id"),
		orm.Column("metadata"),
		orm.Table(RepoTableName),
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
	results := make([]*Repo, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _Name sql.NullString
		var _URL sql.NullString
		var _Active sql.NullBool
		var _Fork sql.NullBool
		var _ParentID sql.NullString
		var _Description sql.NullString
		var _CreatedAt sql.NullInt64
		var _UpdatedAt sql.NullInt64
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefID sql.NullString
		var _Metadata sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_Name,
			&_URL,
			&_Active,
			&_Fork,
			&_ParentID,
			&_Description,
			&_CreatedAt,
			&_UpdatedAt,
			&_CustomerID,
			&_RefType,
			&_RefID,
			&_Metadata,
		)
		if err != nil {
			return nil, err
		}
		t := &Repo{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _URL.Valid {
			t.SetURL(_URL.String)
		}
		if _Active.Valid {
			t.SetActive(_Active.Bool)
		}
		if _Fork.Valid {
			t.SetFork(_Fork.Bool)
		}
		if _ParentID.Valid {
			t.SetParentID(_ParentID.String)
		}
		if _Description.Valid {
			t.SetDescription(_Description.String)
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
		if _RefType.Valid {
			t.SetRefType(_RefType.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		if _Metadata.Valid {
			t.SetMetadata(_Metadata.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindReposTx will find a Repo record in the database with the provided parameters using the provided transaction
func FindReposTx(ctx context.Context, tx Tx, _params ...interface{}) ([]*Repo, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("name"),
		orm.Column("url"),
		orm.Column("active"),
		orm.Column("fork"),
		orm.Column("parent_id"),
		orm.Column("description"),
		orm.Column("created_at"),
		orm.Column("updated_at"),
		orm.Column("customer_id"),
		orm.Column("ref_type"),
		orm.Column("ref_id"),
		orm.Column("metadata"),
		orm.Table(RepoTableName),
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
	results := make([]*Repo, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _Name sql.NullString
		var _URL sql.NullString
		var _Active sql.NullBool
		var _Fork sql.NullBool
		var _ParentID sql.NullString
		var _Description sql.NullString
		var _CreatedAt sql.NullInt64
		var _UpdatedAt sql.NullInt64
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefID sql.NullString
		var _Metadata sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_Name,
			&_URL,
			&_Active,
			&_Fork,
			&_ParentID,
			&_Description,
			&_CreatedAt,
			&_UpdatedAt,
			&_CustomerID,
			&_RefType,
			&_RefID,
			&_Metadata,
		)
		if err != nil {
			return nil, err
		}
		t := &Repo{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _URL.Valid {
			t.SetURL(_URL.String)
		}
		if _Active.Valid {
			t.SetActive(_Active.Bool)
		}
		if _Fork.Valid {
			t.SetFork(_Fork.Bool)
		}
		if _ParentID.Valid {
			t.SetParentID(_ParentID.String)
		}
		if _Description.Valid {
			t.SetDescription(_Description.String)
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
		if _RefType.Valid {
			t.SetRefType(_RefType.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		if _Metadata.Valid {
			t.SetMetadata(_Metadata.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// DBFind will find a Repo record in the database with the provided parameters
func (t *Repo) DBFind(ctx context.Context, db DB, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("name"),
		orm.Column("url"),
		orm.Column("active"),
		orm.Column("fork"),
		orm.Column("parent_id"),
		orm.Column("description"),
		orm.Column("created_at"),
		orm.Column("updated_at"),
		orm.Column("customer_id"),
		orm.Column("ref_type"),
		orm.Column("ref_id"),
		orm.Column("metadata"),
		orm.Table(RepoTableName),
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
	var _URL sql.NullString
	var _Active sql.NullBool
	var _Fork sql.NullBool
	var _ParentID sql.NullString
	var _Description sql.NullString
	var _CreatedAt sql.NullInt64
	var _UpdatedAt sql.NullInt64
	var _CustomerID sql.NullString
	var _RefType sql.NullString
	var _RefID sql.NullString
	var _Metadata sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_Name,
		&_URL,
		&_Active,
		&_Fork,
		&_ParentID,
		&_Description,
		&_CreatedAt,
		&_UpdatedAt,
		&_CustomerID,
		&_RefType,
		&_RefID,
		&_Metadata,
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
	if _URL.Valid {
		t.SetURL(_URL.String)
	}
	if _Active.Valid {
		t.SetActive(_Active.Bool)
	}
	if _Fork.Valid {
		t.SetFork(_Fork.Bool)
	}
	if _ParentID.Valid {
		t.SetParentID(_ParentID.String)
	}
	if _Description.Valid {
		t.SetDescription(_Description.String)
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
	if _RefType.Valid {
		t.SetRefType(_RefType.String)
	}
	if _RefID.Valid {
		t.SetRefID(_RefID.String)
	}
	if _Metadata.Valid {
		t.SetMetadata(_Metadata.String)
	}
	return true, nil
}

// DBFindTx will find a Repo record in the database with the provided parameters using the provided transaction
func (t *Repo) DBFindTx(ctx context.Context, tx Tx, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("name"),
		orm.Column("url"),
		orm.Column("active"),
		orm.Column("fork"),
		orm.Column("parent_id"),
		orm.Column("description"),
		orm.Column("created_at"),
		orm.Column("updated_at"),
		orm.Column("customer_id"),
		orm.Column("ref_type"),
		orm.Column("ref_id"),
		orm.Column("metadata"),
		orm.Table(RepoTableName),
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
	var _URL sql.NullString
	var _Active sql.NullBool
	var _Fork sql.NullBool
	var _ParentID sql.NullString
	var _Description sql.NullString
	var _CreatedAt sql.NullInt64
	var _UpdatedAt sql.NullInt64
	var _CustomerID sql.NullString
	var _RefType sql.NullString
	var _RefID sql.NullString
	var _Metadata sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_Name,
		&_URL,
		&_Active,
		&_Fork,
		&_ParentID,
		&_Description,
		&_CreatedAt,
		&_UpdatedAt,
		&_CustomerID,
		&_RefType,
		&_RefID,
		&_Metadata,
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
	if _URL.Valid {
		t.SetURL(_URL.String)
	}
	if _Active.Valid {
		t.SetActive(_Active.Bool)
	}
	if _Fork.Valid {
		t.SetFork(_Fork.Bool)
	}
	if _ParentID.Valid {
		t.SetParentID(_ParentID.String)
	}
	if _Description.Valid {
		t.SetDescription(_Description.String)
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
	if _RefType.Valid {
		t.SetRefType(_RefType.String)
	}
	if _RefID.Valid {
		t.SetRefID(_RefID.String)
	}
	if _Metadata.Valid {
		t.SetMetadata(_Metadata.String)
	}
	return true, nil
}

// CountRepos will find the count of Repo records in the database
func CountRepos(ctx context.Context, db DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(RepoTableName),
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

// CountReposTx will find the count of Repo records in the database using the provided transaction
func CountReposTx(ctx context.Context, tx Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(RepoTableName),
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

// DBCount will find the count of Repo records in the database
func (t *Repo) DBCount(ctx context.Context, db DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(RepoTableName),
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

// DBCountTx will find the count of Repo records in the database using the provided transaction
func (t *Repo) DBCountTx(ctx context.Context, tx Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(RepoTableName),
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

// DBExists will return true if the Repo record exists in the database
func (t *Repo) DBExists(ctx context.Context, db DB) (bool, error) {
	q := "SELECT `id` FROM `repo` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := db.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// DBExistsTx will return true if the Repo record exists in the database using the provided transaction
func (t *Repo) DBExistsTx(ctx context.Context, tx Tx) (bool, error) {
	q := "SELECT `id` FROM `repo` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := tx.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// PrimaryKeyColumn returns the column name for the primary key
func (t *Repo) PrimaryKeyColumn() string {
	return RepoColumnID
}

// PrimaryKeyColumnType returns the primary key column Go type as a string
func (t *Repo) PrimaryKeyColumnType() string {
	return "string"
}

// PrimaryKey returns the primary key column value
func (t *Repo) PrimaryKey() interface{} {
	return t.ID
}
