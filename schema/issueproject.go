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

var _ Model = (*IssueProject)(nil)
var _ CSVWriter = (*IssueProject)(nil)
var _ JSONWriter = (*IssueProject)(nil)
var _ Checksum = (*IssueProject)(nil)

// IssueProjectTableName is the name of the table in SQL
const IssueProjectTableName = "issue_project"

var IssueProjectColumns = []string{
	"id",
	"checksum",
	"name",
	"url",
	"created_at",
	"customer_id",
	"ref_type",
	"ref_id",
	"metadata",
}

// IssueProject table
type IssueProject struct {
	Checksum   *string `json:"checksum,omitempty"`
	CreatedAt  int64   `json:"created_at"`
	CustomerID string  `json:"customer_id"`
	ID         string  `json:"id"`
	Metadata   *string `json:"metadata,omitempty"`
	Name       string  `json:"name"`
	RefID      string  `json:"ref_id"`
	RefType    string  `json:"ref_type"`
	URL        string  `json:"url"`
}

// TableName returns the SQL table name for IssueProject and satifies the Model interface
func (t *IssueProject) TableName() string {
	return IssueProjectTableName
}

// ToCSV will serialize the IssueProject instance to a CSV compatible array of strings
func (t *IssueProject) ToCSV() []string {
	return []string{
		t.ID,
		t.CalculateChecksum(),
		t.Name,
		t.URL,
		toCSVString(t.CreatedAt),
		t.CustomerID,
		t.RefType,
		t.RefID,
		toCSVString(t.Metadata),
	}
}

// WriteCSV will serialize the IssueProject instance to the writer as CSV and satisfies the CSVWriter interface
func (t *IssueProject) WriteCSV(w *csv.Writer) error {
	return w.Write(t.ToCSV())
}

// WriteJSON will serialize the IssueProject instance to the writer as JSON and satisfies the JSONWriter interface
func (t *IssueProject) WriteJSON(w io.Writer, indent ...bool) error {
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

// NewIssueProjectReader creates a JSON reader which can read in IssueProject objects serialized as JSON either as an array, single object or json new lines
// and writes each IssueProject to the channel provided
func NewIssueProjectReader(r io.Reader, ch chan<- IssueProject) error {
	return orm.Deserialize(r, func(buf json.RawMessage) error {
		dec := json.NewDecoder(bytes.NewBuffer(buf))
		e := IssueProject{}
		if err := dec.Decode(&e); err != nil {
			return err
		}
		ch <- e
		return nil
	})
}

// NewCSVIssueProjectReaderDir will read the reader as CSV and emit each record to the channel provided
func NewCSVIssueProjectReader(r io.Reader, ch chan<- IssueProject) error {
	cr := csv.NewReader(r)
	for {
		record, err := cr.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		ch <- IssueProject{
			ID:         record[0],
			Checksum:   fromStringPointer(record[1]),
			Name:       record[2],
			URL:        record[3],
			CreatedAt:  fromCSVInt64(record[4]),
			CustomerID: record[5],
			RefType:    record[6],
			RefID:      record[7],
			Metadata:   fromStringPointer(record[8]),
		}
	}
	return nil
}

// NewCSVIssueProjectReaderFile will read the file as a CSV and emit each record to the channel provided
func NewCSVIssueProjectReaderFile(fp string, ch chan<- IssueProject) error {
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
	return NewCSVIssueProjectReader(fc, ch)
}

// NewCSVIssueProjectReaderDir will read the issue_project.csv.gz file as a CSV and emit each record to the channel provided
func NewCSVIssueProjectReaderDir(dir string, ch chan<- IssueProject) error {
	return NewCSVIssueProjectReaderFile(filepath.Join(dir, "issue_project.csv.gz"), ch)
}

// IssueProjectCSVDeduper is a function callback which takes the existing value (a) and the new value (b)
// and the return value should be the one to use (or a new one, if applicable). return nil
// to skip processing of this record
type IssueProjectCSVDeduper func(a IssueProject, b IssueProject) *IssueProject

// IssueProjectCSVDedupeDisabled is set on whether the CSV writer should de-dupe values by key
var IssueProjectCSVDedupeDisabled bool

// NewIssueProjectCSVWriterSize creates a batch writer that will write each IssueProject into a CSV file
// this method will automatically de-duplicate entries using the primary key. if the checksum
// for a newer item with the same primary key doesn't match a previously sent item, the newer
// one will be used
func NewIssueProjectCSVWriterSize(w io.Writer, size int, dedupers ...IssueProjectCSVDeduper) (chan IssueProject, chan bool, error) {
	cw := csv.NewWriter(w)
	ch := make(chan IssueProject, size)
	done := make(chan bool)
	go func() {
		defer func() { done <- true }()
		dodedupe := !IssueProjectCSVDedupeDisabled
		var kv map[string]*IssueProject
		var deduper IssueProjectCSVDeduper
		if dedupers != nil && len(dedupers) > 0 {
			deduper = dedupers[0]
			dodedupe = true
		}
		if dodedupe {
			kv = make(map[string]*IssueProject)
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

// IssueProjectCSVDefaultSize is the default channel buffer size if not provided
var IssueProjectCSVDefaultSize = 100

// NewIssueProjectCSVWriter creates a batch writer that will write each IssueProject into a CSV file
func NewIssueProjectCSVWriter(w io.Writer, dedupers ...IssueProjectCSVDeduper) (chan IssueProject, chan bool, error) {
	return NewIssueProjectCSVWriterSize(w, IssueProjectCSVDefaultSize, dedupers...)
}

// NewIssueProjectCSVWriterDir creates a batch writer that will write each IssueProject into a CSV file named issue_project.csv.gz in dir
func NewIssueProjectCSVWriterDir(dir string, dedupers ...IssueProjectCSVDeduper) (chan IssueProject, chan bool, error) {
	return NewIssueProjectCSVWriterFile(filepath.Join(dir, "issue_project.csv.gz"), dedupers...)
}

// NewIssueProjectCSVWriterFile creates a batch writer that will write each IssueProject into a CSV file
func NewIssueProjectCSVWriterFile(fn string, dedupers ...IssueProjectCSVDeduper) (chan IssueProject, chan bool, error) {
	f, err := os.Create(fn)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening CSV file %s. %v", fn, err)
	}
	var fc io.WriteCloser = f
	if filepath.Ext(fn) == ".gz" {
		w, _ := gzip.NewWriterLevel(f, gzip.BestCompression)
		fc = w
	}
	ch, done, err := NewIssueProjectCSVWriter(fc, dedupers...)
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

type IssueProjectDBAction func(ctx context.Context, db DB, record IssueProject) error

// NewIssueProjectDBWriterSize creates a DB writer that will write each issue into the DB
func NewIssueProjectDBWriterSize(ctx context.Context, db DB, errors chan<- error, size int, actions ...IssueProjectDBAction) (chan IssueProject, chan bool, error) {
	ch := make(chan IssueProject, size)
	done := make(chan bool)
	var action IssueProjectDBAction
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

// NewIssueProjectDBWriter creates a DB writer that will write each issue into the DB
func NewIssueProjectDBWriter(ctx context.Context, db DB, errors chan<- error, actions ...IssueProjectDBAction) (chan IssueProject, chan bool, error) {
	return NewIssueProjectDBWriterSize(ctx, db, errors, 100, actions...)
}

// IssueProjectColumnID is the ID SQL column name for the IssueProject table
const IssueProjectColumnID = "id"

// IssueProjectEscapedColumnID is the escaped ID SQL column name for the IssueProject table
const IssueProjectEscapedColumnID = "`id`"

// IssueProjectColumnChecksum is the Checksum SQL column name for the IssueProject table
const IssueProjectColumnChecksum = "checksum"

// IssueProjectEscapedColumnChecksum is the escaped Checksum SQL column name for the IssueProject table
const IssueProjectEscapedColumnChecksum = "`checksum`"

// IssueProjectColumnName is the Name SQL column name for the IssueProject table
const IssueProjectColumnName = "name"

// IssueProjectEscapedColumnName is the escaped Name SQL column name for the IssueProject table
const IssueProjectEscapedColumnName = "`name`"

// IssueProjectColumnURL is the URL SQL column name for the IssueProject table
const IssueProjectColumnURL = "url"

// IssueProjectEscapedColumnURL is the escaped URL SQL column name for the IssueProject table
const IssueProjectEscapedColumnURL = "`url`"

// IssueProjectColumnCreatedAt is the CreatedAt SQL column name for the IssueProject table
const IssueProjectColumnCreatedAt = "created_at"

// IssueProjectEscapedColumnCreatedAt is the escaped CreatedAt SQL column name for the IssueProject table
const IssueProjectEscapedColumnCreatedAt = "`created_at`"

// IssueProjectColumnCustomerID is the CustomerID SQL column name for the IssueProject table
const IssueProjectColumnCustomerID = "customer_id"

// IssueProjectEscapedColumnCustomerID is the escaped CustomerID SQL column name for the IssueProject table
const IssueProjectEscapedColumnCustomerID = "`customer_id`"

// IssueProjectColumnRefType is the RefType SQL column name for the IssueProject table
const IssueProjectColumnRefType = "ref_type"

// IssueProjectEscapedColumnRefType is the escaped RefType SQL column name for the IssueProject table
const IssueProjectEscapedColumnRefType = "`ref_type`"

// IssueProjectColumnRefID is the RefID SQL column name for the IssueProject table
const IssueProjectColumnRefID = "ref_id"

// IssueProjectEscapedColumnRefID is the escaped RefID SQL column name for the IssueProject table
const IssueProjectEscapedColumnRefID = "`ref_id`"

// IssueProjectColumnMetadata is the Metadata SQL column name for the IssueProject table
const IssueProjectColumnMetadata = "metadata"

// IssueProjectEscapedColumnMetadata is the escaped Metadata SQL column name for the IssueProject table
const IssueProjectEscapedColumnMetadata = "`metadata`"

// GetID will return the IssueProject ID value
func (t *IssueProject) GetID() string {
	return t.ID
}

// SetID will set the IssueProject ID value
func (t *IssueProject) SetID(v string) {
	t.ID = v
}

// FindIssueProjectByID will find a IssueProject by ID
func FindIssueProjectByID(ctx context.Context, db DB, value string) (*IssueProject, error) {
	q := "SELECT * FROM `issue_project` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _Name sql.NullString
	var _URL sql.NullString
	var _CreatedAt sql.NullInt64
	var _CustomerID sql.NullString
	var _RefType sql.NullString
	var _RefID sql.NullString
	var _Metadata sql.NullString
	err := db.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_Name,
		&_URL,
		&_CreatedAt,
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
	t := &IssueProject{}
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
	if _CreatedAt.Valid {
		t.SetCreatedAt(_CreatedAt.Int64)
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

// FindIssueProjectByIDTx will find a IssueProject by ID using the provided transaction
func FindIssueProjectByIDTx(ctx context.Context, tx Tx, value string) (*IssueProject, error) {
	q := "SELECT * FROM `issue_project` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _Name sql.NullString
	var _URL sql.NullString
	var _CreatedAt sql.NullInt64
	var _CustomerID sql.NullString
	var _RefType sql.NullString
	var _RefID sql.NullString
	var _Metadata sql.NullString
	err := tx.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_Name,
		&_URL,
		&_CreatedAt,
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
	t := &IssueProject{}
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
	if _CreatedAt.Valid {
		t.SetCreatedAt(_CreatedAt.Int64)
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

// GetChecksum will return the IssueProject Checksum value
func (t *IssueProject) GetChecksum() string {
	if t.Checksum == nil {
		return ""
	}
	return *t.Checksum
}

// SetChecksum will set the IssueProject Checksum value
func (t *IssueProject) SetChecksum(v string) {
	t.Checksum = &v
}

// GetName will return the IssueProject Name value
func (t *IssueProject) GetName() string {
	return t.Name
}

// SetName will set the IssueProject Name value
func (t *IssueProject) SetName(v string) {
	t.Name = v
}

// GetURL will return the IssueProject URL value
func (t *IssueProject) GetURL() string {
	return t.URL
}

// SetURL will set the IssueProject URL value
func (t *IssueProject) SetURL(v string) {
	t.URL = v
}

// GetCreatedAt will return the IssueProject CreatedAt value
func (t *IssueProject) GetCreatedAt() int64 {
	return t.CreatedAt
}

// SetCreatedAt will set the IssueProject CreatedAt value
func (t *IssueProject) SetCreatedAt(v int64) {
	t.CreatedAt = v
}

// GetCustomerID will return the IssueProject CustomerID value
func (t *IssueProject) GetCustomerID() string {
	return t.CustomerID
}

// SetCustomerID will set the IssueProject CustomerID value
func (t *IssueProject) SetCustomerID(v string) {
	t.CustomerID = v
}

// FindIssueProjectsByCustomerID will find all IssueProjects by the CustomerID value
func FindIssueProjectsByCustomerID(ctx context.Context, db DB, value string) ([]*IssueProject, error) {
	q := "SELECT * FROM `issue_project` WHERE `customer_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*IssueProject, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _Name sql.NullString
		var _URL sql.NullString
		var _CreatedAt sql.NullInt64
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefID sql.NullString
		var _Metadata sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_Name,
			&_URL,
			&_CreatedAt,
			&_CustomerID,
			&_RefType,
			&_RefID,
			&_Metadata,
		)
		if err != nil {
			return nil, err
		}
		t := &IssueProject{}
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
		if _CreatedAt.Valid {
			t.SetCreatedAt(_CreatedAt.Int64)
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

// FindIssueProjectsByCustomerIDTx will find all IssueProjects by the CustomerID value using the provided transaction
func FindIssueProjectsByCustomerIDTx(ctx context.Context, tx Tx, value string) ([]*IssueProject, error) {
	q := "SELECT * FROM `issue_project` WHERE `customer_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*IssueProject, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _Name sql.NullString
		var _URL sql.NullString
		var _CreatedAt sql.NullInt64
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefID sql.NullString
		var _Metadata sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_Name,
			&_URL,
			&_CreatedAt,
			&_CustomerID,
			&_RefType,
			&_RefID,
			&_Metadata,
		)
		if err != nil {
			return nil, err
		}
		t := &IssueProject{}
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
		if _CreatedAt.Valid {
			t.SetCreatedAt(_CreatedAt.Int64)
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

// GetRefType will return the IssueProject RefType value
func (t *IssueProject) GetRefType() string {
	return t.RefType
}

// SetRefType will set the IssueProject RefType value
func (t *IssueProject) SetRefType(v string) {
	t.RefType = v
}

// FindIssueProjectsByRefType will find all IssueProjects by the RefType value
func FindIssueProjectsByRefType(ctx context.Context, db DB, value string) ([]*IssueProject, error) {
	q := "SELECT * FROM `issue_project` WHERE `ref_type` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*IssueProject, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _Name sql.NullString
		var _URL sql.NullString
		var _CreatedAt sql.NullInt64
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefID sql.NullString
		var _Metadata sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_Name,
			&_URL,
			&_CreatedAt,
			&_CustomerID,
			&_RefType,
			&_RefID,
			&_Metadata,
		)
		if err != nil {
			return nil, err
		}
		t := &IssueProject{}
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
		if _CreatedAt.Valid {
			t.SetCreatedAt(_CreatedAt.Int64)
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

// FindIssueProjectsByRefTypeTx will find all IssueProjects by the RefType value using the provided transaction
func FindIssueProjectsByRefTypeTx(ctx context.Context, tx Tx, value string) ([]*IssueProject, error) {
	q := "SELECT * FROM `issue_project` WHERE `ref_type` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*IssueProject, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _Name sql.NullString
		var _URL sql.NullString
		var _CreatedAt sql.NullInt64
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefID sql.NullString
		var _Metadata sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_Name,
			&_URL,
			&_CreatedAt,
			&_CustomerID,
			&_RefType,
			&_RefID,
			&_Metadata,
		)
		if err != nil {
			return nil, err
		}
		t := &IssueProject{}
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
		if _CreatedAt.Valid {
			t.SetCreatedAt(_CreatedAt.Int64)
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

// GetRefID will return the IssueProject RefID value
func (t *IssueProject) GetRefID() string {
	return t.RefID
}

// SetRefID will set the IssueProject RefID value
func (t *IssueProject) SetRefID(v string) {
	t.RefID = v
}

// FindIssueProjectsByRefID will find all IssueProjects by the RefID value
func FindIssueProjectsByRefID(ctx context.Context, db DB, value string) ([]*IssueProject, error) {
	q := "SELECT * FROM `issue_project` WHERE `ref_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*IssueProject, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _Name sql.NullString
		var _URL sql.NullString
		var _CreatedAt sql.NullInt64
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefID sql.NullString
		var _Metadata sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_Name,
			&_URL,
			&_CreatedAt,
			&_CustomerID,
			&_RefType,
			&_RefID,
			&_Metadata,
		)
		if err != nil {
			return nil, err
		}
		t := &IssueProject{}
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
		if _CreatedAt.Valid {
			t.SetCreatedAt(_CreatedAt.Int64)
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

// FindIssueProjectsByRefIDTx will find all IssueProjects by the RefID value using the provided transaction
func FindIssueProjectsByRefIDTx(ctx context.Context, tx Tx, value string) ([]*IssueProject, error) {
	q := "SELECT * FROM `issue_project` WHERE `ref_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*IssueProject, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _Name sql.NullString
		var _URL sql.NullString
		var _CreatedAt sql.NullInt64
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefID sql.NullString
		var _Metadata sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_Name,
			&_URL,
			&_CreatedAt,
			&_CustomerID,
			&_RefType,
			&_RefID,
			&_Metadata,
		)
		if err != nil {
			return nil, err
		}
		t := &IssueProject{}
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
		if _CreatedAt.Valid {
			t.SetCreatedAt(_CreatedAt.Int64)
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

// GetMetadata will return the IssueProject Metadata value
func (t *IssueProject) GetMetadata() string {
	if t.Metadata == nil {
		return ""
	}
	return *t.Metadata
}

// SetMetadata will set the IssueProject Metadata value
func (t *IssueProject) SetMetadata(v string) {
	t.Metadata = &v
}

func (t *IssueProject) toTimestamp(value time.Time) *timestamp.Timestamp {
	ts, _ := ptypes.TimestampProto(value)
	return ts
}

// DBCreateIssueProjectTable will create the IssueProject table
func DBCreateIssueProjectTable(ctx context.Context, db DB) error {
	q := "CREATE TABLE `issue_project` (`id` VARCHAR(64) NOT NULL PRIMARY KEY,`checksum` CHAR(64),`name`TEXT NOT NULL,`url` TEXT NOT NULL,`created_at`BIGINT UNSIGNED NOT NULL,`customer_id` VARCHAR(64) NOT NULL,`ref_type` VARCHAR(20) NOT NULL,`ref_id` VARCHAR(64) NOT NULL,`metadata` JSON,INDEX issue_project_customer_id_index (`customer_id`),INDEX issue_project_ref_type_index (`ref_type`),INDEX issue_project_ref_id_index (`ref_id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBCreateIssueProjectTableTx will create the IssueProject table using the provided transction
func DBCreateIssueProjectTableTx(ctx context.Context, tx Tx) error {
	q := "CREATE TABLE `issue_project` (`id` VARCHAR(64) NOT NULL PRIMARY KEY,`checksum` CHAR(64),`name`TEXT NOT NULL,`url` TEXT NOT NULL,`created_at`BIGINT UNSIGNED NOT NULL,`customer_id` VARCHAR(64) NOT NULL,`ref_type` VARCHAR(20) NOT NULL,`ref_id` VARCHAR(64) NOT NULL,`metadata` JSON,INDEX issue_project_customer_id_index (`customer_id`),INDEX issue_project_ref_type_index (`ref_type`),INDEX issue_project_ref_id_index (`ref_id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// DBDropIssueProjectTable will drop the IssueProject table
func DBDropIssueProjectTable(ctx context.Context, db DB) error {
	q := "DROP TABLE IF EXISTS `issue_project`"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBDropIssueProjectTableTx will drop the IssueProject table using the provided transaction
func DBDropIssueProjectTableTx(ctx context.Context, tx Tx) error {
	q := "DROP TABLE IF EXISTS `issue_project`"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// CalculateChecksum will calculate a checksum of the SHA1 of all field values
func (t *IssueProject) CalculateChecksum() string {
	return orm.HashStrings(
		orm.ToString(t.ID),
		orm.ToString(t.Name),
		orm.ToString(t.URL),
		orm.ToString(t.CreatedAt),
		orm.ToString(t.CustomerID),
		orm.ToString(t.RefType),
		orm.ToString(t.RefID),
		orm.ToString(t.Metadata),
	)
}

// DBCreate will create a new IssueProject record in the database
func (t *IssueProject) DBCreate(ctx context.Context, db DB) (sql.Result, error) {
	q := "INSERT INTO `issue_project` (`issue_project`.`id`,`issue_project`.`checksum`,`issue_project`.`name`,`issue_project`.`url`,`issue_project`.`created_at`,`issue_project`.`customer_id`,`issue_project`.`ref_type`,`issue_project`.`ref_id`,`issue_project`.`metadata`) VALUES (?,?,?,?,?,?,?,?,?)"
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
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefType),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.Metadata),
	)
}

// DBCreateTx will create a new IssueProject record in the database using the provided transaction
func (t *IssueProject) DBCreateTx(ctx context.Context, tx Tx) (sql.Result, error) {
	q := "INSERT INTO `issue_project` (`issue_project`.`id`,`issue_project`.`checksum`,`issue_project`.`name`,`issue_project`.`url`,`issue_project`.`created_at`,`issue_project`.`customer_id`,`issue_project`.`ref_type`,`issue_project`.`ref_id`,`issue_project`.`metadata`) VALUES (?,?,?,?,?,?,?,?,?)"
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
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefType),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.Metadata),
	)
}

// DBCreateIgnoreDuplicate will upsert the IssueProject record in the database
func (t *IssueProject) DBCreateIgnoreDuplicate(ctx context.Context, db DB) (sql.Result, error) {
	q := "INSERT INTO `issue_project` (`issue_project`.`id`,`issue_project`.`checksum`,`issue_project`.`name`,`issue_project`.`url`,`issue_project`.`created_at`,`issue_project`.`customer_id`,`issue_project`.`ref_type`,`issue_project`.`ref_id`,`issue_project`.`metadata`) VALUES (?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
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
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefType),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.Metadata),
	)
}

// DBCreateIgnoreDuplicateTx will upsert the IssueProject record in the database using the provided transaction
func (t *IssueProject) DBCreateIgnoreDuplicateTx(ctx context.Context, tx Tx) (sql.Result, error) {
	q := "INSERT INTO `issue_project` (`issue_project`.`id`,`issue_project`.`checksum`,`issue_project`.`name`,`issue_project`.`url`,`issue_project`.`created_at`,`issue_project`.`customer_id`,`issue_project`.`ref_type`,`issue_project`.`ref_id`,`issue_project`.`metadata`) VALUES (?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
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
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefType),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.Metadata),
	)
}

// DeleteAllIssueProjects deletes all IssueProject records in the database with optional filters
func DeleteAllIssueProjects(ctx context.Context, db DB, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(IssueProjectTableName),
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

// DeleteAllIssueProjectsTx deletes all IssueProject records in the database with optional filters using the provided transaction
func DeleteAllIssueProjectsTx(ctx context.Context, tx Tx, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(IssueProjectTableName),
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

// DBDelete will delete this IssueProject record in the database
func (t *IssueProject) DBDelete(ctx context.Context, db DB) (bool, error) {
	q := "DELETE FROM `issue_project` WHERE `id` = ?"
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

// DBDeleteTx will delete this IssueProject record in the database using the provided transaction
func (t *IssueProject) DBDeleteTx(ctx context.Context, tx Tx) (bool, error) {
	q := "DELETE FROM `issue_project` WHERE `id` = ?"
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

// DBUpdate will update the IssueProject record in the database
func (t *IssueProject) DBUpdate(ctx context.Context, db DB) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `issue_project` SET `checksum`=?,`name`=?,`url`=?,`created_at`=?,`customer_id`=?,`ref_type`=?,`ref_id`=?,`metadata`=? WHERE `id`=?"
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.URL),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefType),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.Metadata),
		orm.ToSQLString(t.ID),
	)
}

// DBUpdateTx will update the IssueProject record in the database using the provided transaction
func (t *IssueProject) DBUpdateTx(ctx context.Context, tx Tx) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `issue_project` SET `checksum`=?,`name`=?,`url`=?,`created_at`=?,`customer_id`=?,`ref_type`=?,`ref_id`=?,`metadata`=? WHERE `id`=?"
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.URL),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefType),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.Metadata),
		orm.ToSQLString(t.ID),
	)
}

// DBUpsert will upsert the IssueProject record in the database
func (t *IssueProject) DBUpsert(ctx context.Context, db DB, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `issue_project` (`issue_project`.`id`,`issue_project`.`checksum`,`issue_project`.`name`,`issue_project`.`url`,`issue_project`.`created_at`,`issue_project`.`customer_id`,`issue_project`.`ref_type`,`issue_project`.`ref_id`,`issue_project`.`metadata`) VALUES (?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `issue_project` (`issue_project`.`id`,`issue_project`.`checksum`,`issue_project`.`name`,`issue_project`.`url`,`issue_project`.`created_at`,`issue_project`.`customer_id`,`issue_project`.`ref_type`,`issue_project`.`ref_id`,`issue_project`.`metadata`) VALUES (?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`name`=VALUES(`name`),`url`=VALUES(`url`),`created_at`=VALUES(`created_at`),`customer_id`=VALUES(`customer_id`),`ref_type`=VALUES(`ref_type`),`ref_id`=VALUES(`ref_id`),`metadata`=VALUES(`metadata`)"
	}
	r, err := db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.URL),
		orm.ToSQLInt64(t.CreatedAt),
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

// DBUpsertTx will upsert the IssueProject record in the database using the provided transaction
func (t *IssueProject) DBUpsertTx(ctx context.Context, tx Tx, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `issue_project` (`issue_project`.`id`,`issue_project`.`checksum`,`issue_project`.`name`,`issue_project`.`url`,`issue_project`.`created_at`,`issue_project`.`customer_id`,`issue_project`.`ref_type`,`issue_project`.`ref_id`,`issue_project`.`metadata`) VALUES (?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `issue_project` (`issue_project`.`id`,`issue_project`.`checksum`,`issue_project`.`name`,`issue_project`.`url`,`issue_project`.`created_at`,`issue_project`.`customer_id`,`issue_project`.`ref_type`,`issue_project`.`ref_id`,`issue_project`.`metadata`) VALUES (?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`name`=VALUES(`name`),`url`=VALUES(`url`),`created_at`=VALUES(`created_at`),`customer_id`=VALUES(`customer_id`),`ref_type`=VALUES(`ref_type`),`ref_id`=VALUES(`ref_id`),`metadata`=VALUES(`metadata`)"
	}
	r, err := tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.URL),
		orm.ToSQLInt64(t.CreatedAt),
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

// DBFindOne will find a IssueProject record in the database with the primary key
func (t *IssueProject) DBFindOne(ctx context.Context, db DB, value string) (bool, error) {
	q := "SELECT * FROM `issue_project` WHERE `id` = ? LIMIT 1"
	row := db.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _Name sql.NullString
	var _URL sql.NullString
	var _CreatedAt sql.NullInt64
	var _CustomerID sql.NullString
	var _RefType sql.NullString
	var _RefID sql.NullString
	var _Metadata sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_Name,
		&_URL,
		&_CreatedAt,
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
	if _CreatedAt.Valid {
		t.SetCreatedAt(_CreatedAt.Int64)
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

// DBFindOneTx will find a IssueProject record in the database with the primary key using the provided transaction
func (t *IssueProject) DBFindOneTx(ctx context.Context, tx Tx, value string) (bool, error) {
	q := "SELECT * FROM `issue_project` WHERE `id` = ? LIMIT 1"
	row := tx.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _Name sql.NullString
	var _URL sql.NullString
	var _CreatedAt sql.NullInt64
	var _CustomerID sql.NullString
	var _RefType sql.NullString
	var _RefID sql.NullString
	var _Metadata sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_Name,
		&_URL,
		&_CreatedAt,
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
	if _CreatedAt.Valid {
		t.SetCreatedAt(_CreatedAt.Int64)
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

// FindIssueProjects will find a IssueProject record in the database with the provided parameters
func FindIssueProjects(ctx context.Context, db DB, _params ...interface{}) ([]*IssueProject, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("name"),
		orm.Column("url"),
		orm.Column("created_at"),
		orm.Column("customer_id"),
		orm.Column("ref_type"),
		orm.Column("ref_id"),
		orm.Column("metadata"),
		orm.Table(IssueProjectTableName),
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
	results := make([]*IssueProject, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _Name sql.NullString
		var _URL sql.NullString
		var _CreatedAt sql.NullInt64
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefID sql.NullString
		var _Metadata sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_Name,
			&_URL,
			&_CreatedAt,
			&_CustomerID,
			&_RefType,
			&_RefID,
			&_Metadata,
		)
		if err != nil {
			return nil, err
		}
		t := &IssueProject{}
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
		if _CreatedAt.Valid {
			t.SetCreatedAt(_CreatedAt.Int64)
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

// FindIssueProjectsTx will find a IssueProject record in the database with the provided parameters using the provided transaction
func FindIssueProjectsTx(ctx context.Context, tx Tx, _params ...interface{}) ([]*IssueProject, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("name"),
		orm.Column("url"),
		orm.Column("created_at"),
		orm.Column("customer_id"),
		orm.Column("ref_type"),
		orm.Column("ref_id"),
		orm.Column("metadata"),
		orm.Table(IssueProjectTableName),
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
	results := make([]*IssueProject, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _Name sql.NullString
		var _URL sql.NullString
		var _CreatedAt sql.NullInt64
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefID sql.NullString
		var _Metadata sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_Name,
			&_URL,
			&_CreatedAt,
			&_CustomerID,
			&_RefType,
			&_RefID,
			&_Metadata,
		)
		if err != nil {
			return nil, err
		}
		t := &IssueProject{}
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
		if _CreatedAt.Valid {
			t.SetCreatedAt(_CreatedAt.Int64)
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

// DBFind will find a IssueProject record in the database with the provided parameters
func (t *IssueProject) DBFind(ctx context.Context, db DB, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("name"),
		orm.Column("url"),
		orm.Column("created_at"),
		orm.Column("customer_id"),
		orm.Column("ref_type"),
		orm.Column("ref_id"),
		orm.Column("metadata"),
		orm.Table(IssueProjectTableName),
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
	var _CreatedAt sql.NullInt64
	var _CustomerID sql.NullString
	var _RefType sql.NullString
	var _RefID sql.NullString
	var _Metadata sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_Name,
		&_URL,
		&_CreatedAt,
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
	if _CreatedAt.Valid {
		t.SetCreatedAt(_CreatedAt.Int64)
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

// DBFindTx will find a IssueProject record in the database with the provided parameters using the provided transaction
func (t *IssueProject) DBFindTx(ctx context.Context, tx Tx, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("name"),
		orm.Column("url"),
		orm.Column("created_at"),
		orm.Column("customer_id"),
		orm.Column("ref_type"),
		orm.Column("ref_id"),
		orm.Column("metadata"),
		orm.Table(IssueProjectTableName),
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
	var _CreatedAt sql.NullInt64
	var _CustomerID sql.NullString
	var _RefType sql.NullString
	var _RefID sql.NullString
	var _Metadata sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_Name,
		&_URL,
		&_CreatedAt,
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
	if _CreatedAt.Valid {
		t.SetCreatedAt(_CreatedAt.Int64)
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

// CountIssueProjects will find the count of IssueProject records in the database
func CountIssueProjects(ctx context.Context, db DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(IssueProjectTableName),
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

// CountIssueProjectsTx will find the count of IssueProject records in the database using the provided transaction
func CountIssueProjectsTx(ctx context.Context, tx Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(IssueProjectTableName),
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

// DBCount will find the count of IssueProject records in the database
func (t *IssueProject) DBCount(ctx context.Context, db DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(IssueProjectTableName),
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

// DBCountTx will find the count of IssueProject records in the database using the provided transaction
func (t *IssueProject) DBCountTx(ctx context.Context, tx Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(IssueProjectTableName),
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

// DBExists will return true if the IssueProject record exists in the database
func (t *IssueProject) DBExists(ctx context.Context, db DB) (bool, error) {
	q := "SELECT * FROM `issue_project` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := db.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// DBExistsTx will return true if the IssueProject record exists in the database using the provided transaction
func (t *IssueProject) DBExistsTx(ctx context.Context, tx Tx) (bool, error) {
	q := "SELECT * FROM `issue_project` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := tx.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// PrimaryKeyColumn returns the column name for the primary key
func (t *IssueProject) PrimaryKeyColumn() string {
	return IssueProjectColumnID
}

// PrimaryKeyColumnType returns the primary key column Go type as a string
func (t *IssueProject) PrimaryKeyColumnType() string {
	return "string"
}

// PrimaryKey returns the primary key column value
func (t *IssueProject) PrimaryKey() interface{} {
	return t.ID
}
