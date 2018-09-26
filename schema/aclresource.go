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

var _ Model = (*ACLResource)(nil)
var _ CSVWriter = (*ACLResource)(nil)
var _ JSONWriter = (*ACLResource)(nil)
var _ Checksum = (*ACLResource)(nil)

// ACLResourceTableName is the name of the table in SQL
const ACLResourceTableName = "acl_resource"

var ACLResourceColumns = []string{
	"id",
	"checksum",
	"urn",
	"description",
	"title",
	"public",
	"hidden",
	"admin",
	"created_at",
	"updated_at",
}

// ACLResource table
type ACLResource struct {
	Admin       bool    `json:"admin"`
	Checksum    *string `json:"checksum,omitempty"`
	CreatedAt   int64   `json:"created_at"`
	Description *string `json:"description,omitempty"`
	Hidden      bool    `json:"hidden"`
	ID          string  `json:"id"`
	Public      bool    `json:"public"`
	Title       *string `json:"title,omitempty"`
	UpdatedAt   *int64  `json:"updated_at,omitempty"`
	Urn         string  `json:"urn"`
}

// TableName returns the SQL table name for ACLResource and satifies the Model interface
func (t *ACLResource) TableName() string {
	return ACLResourceTableName
}

// ToCSV will serialize the ACLResource instance to a CSV compatible array of strings
func (t *ACLResource) ToCSV() []string {
	return []string{
		t.ID,
		t.CalculateChecksum(),
		t.Urn,
		toCSVString(t.Description),
		toCSVString(t.Title),
		toCSVBool(t.Public),
		toCSVBool(t.Hidden),
		toCSVBool(t.Admin),
		toCSVString(t.CreatedAt),
		toCSVString(t.UpdatedAt),
	}
}

// WriteCSV will serialize the ACLResource instance to the writer as CSV and satisfies the CSVWriter interface
func (t *ACLResource) WriteCSV(w *csv.Writer) error {
	return w.Write(t.ToCSV())
}

// WriteJSON will serialize the ACLResource instance to the writer as JSON and satisfies the JSONWriter interface
func (t *ACLResource) WriteJSON(w io.Writer, indent ...bool) error {
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

// NewACLResourceReader creates a JSON reader which can read in ACLResource objects serialized as JSON either as an array, single object or json new lines
// and writes each ACLResource to the channel provided
func NewACLResourceReader(r io.Reader, ch chan<- ACLResource) error {
	return orm.Deserialize(r, func(buf json.RawMessage) error {
		dec := json.NewDecoder(bytes.NewBuffer(buf))
		e := ACLResource{}
		if err := dec.Decode(&e); err != nil {
			return err
		}
		ch <- e
		return nil
	})
}

// NewCSVACLResourceReaderDir will read the reader as CSV and emit each record to the channel provided
func NewCSVACLResourceReader(r io.Reader, ch chan<- ACLResource) error {
	cr := csv.NewReader(r)
	for {
		record, err := cr.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		ch <- ACLResource{
			ID:          record[0],
			Checksum:    fromStringPointer(record[1]),
			Urn:         record[2],
			Description: fromStringPointer(record[3]),
			Title:       fromStringPointer(record[4]),
			Public:      fromCSVBool(record[5]),
			Hidden:      fromCSVBool(record[6]),
			Admin:       fromCSVBool(record[7]),
			CreatedAt:   fromCSVInt64(record[8]),
			UpdatedAt:   fromCSVInt64Pointer(record[9]),
		}
	}
	return nil
}

// NewCSVACLResourceReaderFile will read the file as a CSV and emit each record to the channel provided
func NewCSVACLResourceReaderFile(fp string, ch chan<- ACLResource) error {
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
	return NewCSVACLResourceReader(fc, ch)
}

// NewCSVACLResourceReaderDir will read the acl_resource.csv.gz file as a CSV and emit each record to the channel provided
func NewCSVACLResourceReaderDir(dir string, ch chan<- ACLResource) error {
	return NewCSVACLResourceReaderFile(filepath.Join(dir, "acl_resource.csv.gz"), ch)
}

// ACLResourceCSVDeduper is a function callback which takes the existing value (a) and the new value (b)
// and the return value should be the one to use (or a new one, if applicable). return nil
// to skip processing of this record
type ACLResourceCSVDeduper func(a ACLResource, b ACLResource) *ACLResource

// ACLResourceCSVDedupeDisabled is set on whether the CSV writer should de-dupe values by key
var ACLResourceCSVDedupeDisabled bool

// NewACLResourceCSVWriterSize creates a batch writer that will write each ACLResource into a CSV file
// this method will automatically de-duplicate entries using the primary key. if the checksum
// for a newer item with the same primary key doesn't match a previously sent item, the newer
// one will be used
func NewACLResourceCSVWriterSize(w io.Writer, size int, dedupers ...ACLResourceCSVDeduper) (chan ACLResource, chan bool, error) {
	cw := csv.NewWriter(w)
	ch := make(chan ACLResource, size)
	done := make(chan bool)
	go func() {
		defer func() { done <- true }()
		dodedupe := !ACLResourceCSVDedupeDisabled
		var kv map[string]*ACLResource
		var deduper ACLResourceCSVDeduper
		if dedupers != nil && len(dedupers) > 0 {
			deduper = dedupers[0]
			dodedupe = true
		}
		if dodedupe {
			kv = make(map[string]*ACLResource)
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

// ACLResourceCSVDefaultSize is the default channel buffer size if not provided
var ACLResourceCSVDefaultSize = 100

// NewACLResourceCSVWriter creates a batch writer that will write each ACLResource into a CSV file
func NewACLResourceCSVWriter(w io.Writer, dedupers ...ACLResourceCSVDeduper) (chan ACLResource, chan bool, error) {
	return NewACLResourceCSVWriterSize(w, ACLResourceCSVDefaultSize, dedupers...)
}

// NewACLResourceCSVWriterDir creates a batch writer that will write each ACLResource into a CSV file named acl_resource.csv.gz in dir
func NewACLResourceCSVWriterDir(dir string, dedupers ...ACLResourceCSVDeduper) (chan ACLResource, chan bool, error) {
	return NewACLResourceCSVWriterFile(filepath.Join(dir, "acl_resource.csv.gz"), dedupers...)
}

// NewACLResourceCSVWriterFile creates a batch writer that will write each ACLResource into a CSV file
func NewACLResourceCSVWriterFile(fn string, dedupers ...ACLResourceCSVDeduper) (chan ACLResource, chan bool, error) {
	f, err := os.Create(fn)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening CSV file %s. %v", fn, err)
	}
	var fc io.WriteCloser = f
	if filepath.Ext(fn) == ".gz" {
		w, _ := gzip.NewWriterLevel(f, gzip.BestCompression)
		fc = w
	}
	ch, done, err := NewACLResourceCSVWriter(fc, dedupers...)
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

type ACLResourceDBAction func(ctx context.Context, db *sql.DB, record ACLResource) error

// NewACLResourceDBWriterSize creates a DB writer that will write each issue into the DB
func NewACLResourceDBWriterSize(ctx context.Context, db *sql.DB, errors chan<- error, size int, actions ...ACLResourceDBAction) (chan ACLResource, chan bool, error) {
	ch := make(chan ACLResource, size)
	done := make(chan bool)
	var action ACLResourceDBAction
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

// NewACLResourceDBWriter creates a DB writer that will write each issue into the DB
func NewACLResourceDBWriter(ctx context.Context, db *sql.DB, errors chan<- error, actions ...ACLResourceDBAction) (chan ACLResource, chan bool, error) {
	return NewACLResourceDBWriterSize(ctx, db, errors, 100, actions...)
}

// ACLResourceColumnID is the ID SQL column name for the ACLResource table
const ACLResourceColumnID = "id"

// ACLResourceEscapedColumnID is the escaped ID SQL column name for the ACLResource table
const ACLResourceEscapedColumnID = "`id`"

// ACLResourceColumnChecksum is the Checksum SQL column name for the ACLResource table
const ACLResourceColumnChecksum = "checksum"

// ACLResourceEscapedColumnChecksum is the escaped Checksum SQL column name for the ACLResource table
const ACLResourceEscapedColumnChecksum = "`checksum`"

// ACLResourceColumnUrn is the Urn SQL column name for the ACLResource table
const ACLResourceColumnUrn = "urn"

// ACLResourceEscapedColumnUrn is the escaped Urn SQL column name for the ACLResource table
const ACLResourceEscapedColumnUrn = "`urn`"

// ACLResourceColumnDescription is the Description SQL column name for the ACLResource table
const ACLResourceColumnDescription = "description"

// ACLResourceEscapedColumnDescription is the escaped Description SQL column name for the ACLResource table
const ACLResourceEscapedColumnDescription = "`description`"

// ACLResourceColumnTitle is the Title SQL column name for the ACLResource table
const ACLResourceColumnTitle = "title"

// ACLResourceEscapedColumnTitle is the escaped Title SQL column name for the ACLResource table
const ACLResourceEscapedColumnTitle = "`title`"

// ACLResourceColumnPublic is the Public SQL column name for the ACLResource table
const ACLResourceColumnPublic = "public"

// ACLResourceEscapedColumnPublic is the escaped Public SQL column name for the ACLResource table
const ACLResourceEscapedColumnPublic = "`public`"

// ACLResourceColumnHidden is the Hidden SQL column name for the ACLResource table
const ACLResourceColumnHidden = "hidden"

// ACLResourceEscapedColumnHidden is the escaped Hidden SQL column name for the ACLResource table
const ACLResourceEscapedColumnHidden = "`hidden`"

// ACLResourceColumnAdmin is the Admin SQL column name for the ACLResource table
const ACLResourceColumnAdmin = "admin"

// ACLResourceEscapedColumnAdmin is the escaped Admin SQL column name for the ACLResource table
const ACLResourceEscapedColumnAdmin = "`admin`"

// ACLResourceColumnCreatedAt is the CreatedAt SQL column name for the ACLResource table
const ACLResourceColumnCreatedAt = "created_at"

// ACLResourceEscapedColumnCreatedAt is the escaped CreatedAt SQL column name for the ACLResource table
const ACLResourceEscapedColumnCreatedAt = "`created_at`"

// ACLResourceColumnUpdatedAt is the UpdatedAt SQL column name for the ACLResource table
const ACLResourceColumnUpdatedAt = "updated_at"

// ACLResourceEscapedColumnUpdatedAt is the escaped UpdatedAt SQL column name for the ACLResource table
const ACLResourceEscapedColumnUpdatedAt = "`updated_at`"

// GetID will return the ACLResource ID value
func (t *ACLResource) GetID() string {
	return t.ID
}

// SetID will set the ACLResource ID value
func (t *ACLResource) SetID(v string) {
	t.ID = v
}

// FindACLResourceByID will find a ACLResource by ID
func FindACLResourceByID(ctx context.Context, db *sql.DB, value string) (*ACLResource, error) {
	q := "SELECT `acl_resource`.`id`,`acl_resource`.`checksum`,`acl_resource`.`urn`,`acl_resource`.`description`,`acl_resource`.`title`,`acl_resource`.`public`,`acl_resource`.`hidden`,`acl_resource`.`admin`,`acl_resource`.`created_at`,`acl_resource`.`updated_at` FROM `acl_resource` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _Urn sql.NullString
	var _Description sql.NullString
	var _Title sql.NullString
	var _Public sql.NullBool
	var _Hidden sql.NullBool
	var _Admin sql.NullBool
	var _CreatedAt sql.NullInt64
	var _UpdatedAt sql.NullInt64
	err := db.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_Urn,
		&_Description,
		&_Title,
		&_Public,
		&_Hidden,
		&_Admin,
		&_CreatedAt,
		&_UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &ACLResource{}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _Checksum.Valid {
		t.SetChecksum(_Checksum.String)
	}
	if _Urn.Valid {
		t.SetUrn(_Urn.String)
	}
	if _Description.Valid {
		t.SetDescription(_Description.String)
	}
	if _Title.Valid {
		t.SetTitle(_Title.String)
	}
	if _Public.Valid {
		t.SetPublic(_Public.Bool)
	}
	if _Hidden.Valid {
		t.SetHidden(_Hidden.Bool)
	}
	if _Admin.Valid {
		t.SetAdmin(_Admin.Bool)
	}
	if _CreatedAt.Valid {
		t.SetCreatedAt(_CreatedAt.Int64)
	}
	if _UpdatedAt.Valid {
		t.SetUpdatedAt(_UpdatedAt.Int64)
	}
	return t, nil
}

// FindACLResourceByIDTx will find a ACLResource by ID using the provided transaction
func FindACLResourceByIDTx(ctx context.Context, tx *sql.Tx, value string) (*ACLResource, error) {
	q := "SELECT `acl_resource`.`id`,`acl_resource`.`checksum`,`acl_resource`.`urn`,`acl_resource`.`description`,`acl_resource`.`title`,`acl_resource`.`public`,`acl_resource`.`hidden`,`acl_resource`.`admin`,`acl_resource`.`created_at`,`acl_resource`.`updated_at` FROM `acl_resource` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _Urn sql.NullString
	var _Description sql.NullString
	var _Title sql.NullString
	var _Public sql.NullBool
	var _Hidden sql.NullBool
	var _Admin sql.NullBool
	var _CreatedAt sql.NullInt64
	var _UpdatedAt sql.NullInt64
	err := tx.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_Urn,
		&_Description,
		&_Title,
		&_Public,
		&_Hidden,
		&_Admin,
		&_CreatedAt,
		&_UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &ACLResource{}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _Checksum.Valid {
		t.SetChecksum(_Checksum.String)
	}
	if _Urn.Valid {
		t.SetUrn(_Urn.String)
	}
	if _Description.Valid {
		t.SetDescription(_Description.String)
	}
	if _Title.Valid {
		t.SetTitle(_Title.String)
	}
	if _Public.Valid {
		t.SetPublic(_Public.Bool)
	}
	if _Hidden.Valid {
		t.SetHidden(_Hidden.Bool)
	}
	if _Admin.Valid {
		t.SetAdmin(_Admin.Bool)
	}
	if _CreatedAt.Valid {
		t.SetCreatedAt(_CreatedAt.Int64)
	}
	if _UpdatedAt.Valid {
		t.SetUpdatedAt(_UpdatedAt.Int64)
	}
	return t, nil
}

// GetChecksum will return the ACLResource Checksum value
func (t *ACLResource) GetChecksum() string {
	if t.Checksum == nil {
		return ""
	}
	return *t.Checksum
}

// SetChecksum will set the ACLResource Checksum value
func (t *ACLResource) SetChecksum(v string) {
	t.Checksum = &v
}

// GetUrn will return the ACLResource Urn value
func (t *ACLResource) GetUrn() string {
	return t.Urn
}

// SetUrn will set the ACLResource Urn value
func (t *ACLResource) SetUrn(v string) {
	t.Urn = v
}

// FindACLResourcesByUrn will find all ACLResources by the Urn value
func FindACLResourcesByUrn(ctx context.Context, db *sql.DB, value string) ([]*ACLResource, error) {
	q := "SELECT `acl_resource`.`id`,`acl_resource`.`checksum`,`acl_resource`.`urn`,`acl_resource`.`description`,`acl_resource`.`title`,`acl_resource`.`public`,`acl_resource`.`hidden`,`acl_resource`.`admin`,`acl_resource`.`created_at`,`acl_resource`.`updated_at` FROM `acl_resource` WHERE `urn` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*ACLResource, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _Urn sql.NullString
		var _Description sql.NullString
		var _Title sql.NullString
		var _Public sql.NullBool
		var _Hidden sql.NullBool
		var _Admin sql.NullBool
		var _CreatedAt sql.NullInt64
		var _UpdatedAt sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_Urn,
			&_Description,
			&_Title,
			&_Public,
			&_Hidden,
			&_Admin,
			&_CreatedAt,
			&_UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		t := &ACLResource{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _Urn.Valid {
			t.SetUrn(_Urn.String)
		}
		if _Description.Valid {
			t.SetDescription(_Description.String)
		}
		if _Title.Valid {
			t.SetTitle(_Title.String)
		}
		if _Public.Valid {
			t.SetPublic(_Public.Bool)
		}
		if _Hidden.Valid {
			t.SetHidden(_Hidden.Bool)
		}
		if _Admin.Valid {
			t.SetAdmin(_Admin.Bool)
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

// FindACLResourcesByUrnTx will find all ACLResources by the Urn value using the provided transaction
func FindACLResourcesByUrnTx(ctx context.Context, tx *sql.Tx, value string) ([]*ACLResource, error) {
	q := "SELECT `acl_resource`.`id`,`acl_resource`.`checksum`,`acl_resource`.`urn`,`acl_resource`.`description`,`acl_resource`.`title`,`acl_resource`.`public`,`acl_resource`.`hidden`,`acl_resource`.`admin`,`acl_resource`.`created_at`,`acl_resource`.`updated_at` FROM `acl_resource` WHERE `urn` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*ACLResource, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _Urn sql.NullString
		var _Description sql.NullString
		var _Title sql.NullString
		var _Public sql.NullBool
		var _Hidden sql.NullBool
		var _Admin sql.NullBool
		var _CreatedAt sql.NullInt64
		var _UpdatedAt sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_Urn,
			&_Description,
			&_Title,
			&_Public,
			&_Hidden,
			&_Admin,
			&_CreatedAt,
			&_UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		t := &ACLResource{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _Urn.Valid {
			t.SetUrn(_Urn.String)
		}
		if _Description.Valid {
			t.SetDescription(_Description.String)
		}
		if _Title.Valid {
			t.SetTitle(_Title.String)
		}
		if _Public.Valid {
			t.SetPublic(_Public.Bool)
		}
		if _Hidden.Valid {
			t.SetHidden(_Hidden.Bool)
		}
		if _Admin.Valid {
			t.SetAdmin(_Admin.Bool)
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

// GetDescription will return the ACLResource Description value
func (t *ACLResource) GetDescription() string {
	if t.Description == nil {
		return ""
	}
	return *t.Description
}

// SetDescription will set the ACLResource Description value
func (t *ACLResource) SetDescription(v string) {
	t.Description = &v
}

// GetTitle will return the ACLResource Title value
func (t *ACLResource) GetTitle() string {
	if t.Title == nil {
		return ""
	}
	return *t.Title
}

// SetTitle will set the ACLResource Title value
func (t *ACLResource) SetTitle(v string) {
	t.Title = &v
}

// GetPublic will return the ACLResource Public value
func (t *ACLResource) GetPublic() bool {
	return t.Public
}

// SetPublic will set the ACLResource Public value
func (t *ACLResource) SetPublic(v bool) {
	t.Public = v
}

// GetHidden will return the ACLResource Hidden value
func (t *ACLResource) GetHidden() bool {
	return t.Hidden
}

// SetHidden will set the ACLResource Hidden value
func (t *ACLResource) SetHidden(v bool) {
	t.Hidden = v
}

// GetAdmin will return the ACLResource Admin value
func (t *ACLResource) GetAdmin() bool {
	return t.Admin
}

// SetAdmin will set the ACLResource Admin value
func (t *ACLResource) SetAdmin(v bool) {
	t.Admin = v
}

// GetCreatedAt will return the ACLResource CreatedAt value
func (t *ACLResource) GetCreatedAt() int64 {
	return t.CreatedAt
}

// SetCreatedAt will set the ACLResource CreatedAt value
func (t *ACLResource) SetCreatedAt(v int64) {
	t.CreatedAt = v
}

// GetUpdatedAt will return the ACLResource UpdatedAt value
func (t *ACLResource) GetUpdatedAt() int64 {
	if t.UpdatedAt == nil {
		return int64(0)
	}
	return *t.UpdatedAt
}

// SetUpdatedAt will set the ACLResource UpdatedAt value
func (t *ACLResource) SetUpdatedAt(v int64) {
	t.UpdatedAt = &v
}

func (t *ACLResource) toTimestamp(value time.Time) *timestamp.Timestamp {
	ts, _ := ptypes.TimestampProto(value)
	return ts
}

// DBCreateACLResourceTable will create the ACLResource table
func DBCreateACLResourceTable(ctx context.Context, db *sql.DB) error {
	q := "CREATE TABLE `acl_resource` (`id`VARCHAR(64) NOT NULL PRIMARY KEY,`checksum`CHAR(64),`urn` VARCHAR(255) NOT NULL,`description`TEXT,`title`TEXT,`public` TINYINT(1) NOT NULL,`hidden` TINYINT(1) NOT NULL,`admin`TINYINT(1) NOT NULL,`created_at` BIGINT UNSIGNED NOT NULL,`updated_at` BIGINT UNSIGNED,INDEX acl_resource_urn_index (`urn`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBCreateACLResourceTableTx will create the ACLResource table using the provided transction
func DBCreateACLResourceTableTx(ctx context.Context, tx *sql.Tx) error {
	q := "CREATE TABLE `acl_resource` (`id`VARCHAR(64) NOT NULL PRIMARY KEY,`checksum`CHAR(64),`urn` VARCHAR(255) NOT NULL,`description`TEXT,`title`TEXT,`public` TINYINT(1) NOT NULL,`hidden` TINYINT(1) NOT NULL,`admin`TINYINT(1) NOT NULL,`created_at` BIGINT UNSIGNED NOT NULL,`updated_at` BIGINT UNSIGNED,INDEX acl_resource_urn_index (`urn`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// DBDropACLResourceTable will drop the ACLResource table
func DBDropACLResourceTable(ctx context.Context, db *sql.DB) error {
	q := "DROP TABLE IF EXISTS `acl_resource`"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBDropACLResourceTableTx will drop the ACLResource table using the provided transaction
func DBDropACLResourceTableTx(ctx context.Context, tx *sql.Tx) error {
	q := "DROP TABLE IF EXISTS `acl_resource`"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// CalculateChecksum will calculate a checksum of the SHA1 of all field values
func (t *ACLResource) CalculateChecksum() string {
	return orm.HashStrings(
		orm.ToString(t.ID),
		orm.ToString(t.Urn),
		orm.ToString(t.Description),
		orm.ToString(t.Title),
		orm.ToString(t.Public),
		orm.ToString(t.Hidden),
		orm.ToString(t.Admin),
		orm.ToString(t.CreatedAt),
		orm.ToString(t.UpdatedAt),
	)
}

// DBCreate will create a new ACLResource record in the database
func (t *ACLResource) DBCreate(ctx context.Context, db *sql.DB) (sql.Result, error) {
	q := "INSERT INTO `acl_resource` (`acl_resource`.`id`,`acl_resource`.`checksum`,`acl_resource`.`urn`,`acl_resource`.`description`,`acl_resource`.`title`,`acl_resource`.`public`,`acl_resource`.`hidden`,`acl_resource`.`admin`,`acl_resource`.`created_at`,`acl_resource`.`updated_at`) VALUES (?,?,?,?,?,?,?,?,?,?)"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.Urn),
		orm.ToSQLString(t.Description),
		orm.ToSQLString(t.Title),
		orm.ToSQLBool(t.Public),
		orm.ToSQLBool(t.Hidden),
		orm.ToSQLBool(t.Admin),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
	)
}

// DBCreateTx will create a new ACLResource record in the database using the provided transaction
func (t *ACLResource) DBCreateTx(ctx context.Context, tx *sql.Tx) (sql.Result, error) {
	q := "INSERT INTO `acl_resource` (`acl_resource`.`id`,`acl_resource`.`checksum`,`acl_resource`.`urn`,`acl_resource`.`description`,`acl_resource`.`title`,`acl_resource`.`public`,`acl_resource`.`hidden`,`acl_resource`.`admin`,`acl_resource`.`created_at`,`acl_resource`.`updated_at`) VALUES (?,?,?,?,?,?,?,?,?,?)"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.Urn),
		orm.ToSQLString(t.Description),
		orm.ToSQLString(t.Title),
		orm.ToSQLBool(t.Public),
		orm.ToSQLBool(t.Hidden),
		orm.ToSQLBool(t.Admin),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
	)
}

// DBCreateIgnoreDuplicate will upsert the ACLResource record in the database
func (t *ACLResource) DBCreateIgnoreDuplicate(ctx context.Context, db *sql.DB) (sql.Result, error) {
	q := "INSERT INTO `acl_resource` (`acl_resource`.`id`,`acl_resource`.`checksum`,`acl_resource`.`urn`,`acl_resource`.`description`,`acl_resource`.`title`,`acl_resource`.`public`,`acl_resource`.`hidden`,`acl_resource`.`admin`,`acl_resource`.`created_at`,`acl_resource`.`updated_at`) VALUES (?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.Urn),
		orm.ToSQLString(t.Description),
		orm.ToSQLString(t.Title),
		orm.ToSQLBool(t.Public),
		orm.ToSQLBool(t.Hidden),
		orm.ToSQLBool(t.Admin),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
	)
}

// DBCreateIgnoreDuplicateTx will upsert the ACLResource record in the database using the provided transaction
func (t *ACLResource) DBCreateIgnoreDuplicateTx(ctx context.Context, tx *sql.Tx) (sql.Result, error) {
	q := "INSERT INTO `acl_resource` (`acl_resource`.`id`,`acl_resource`.`checksum`,`acl_resource`.`urn`,`acl_resource`.`description`,`acl_resource`.`title`,`acl_resource`.`public`,`acl_resource`.`hidden`,`acl_resource`.`admin`,`acl_resource`.`created_at`,`acl_resource`.`updated_at`) VALUES (?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.Urn),
		orm.ToSQLString(t.Description),
		orm.ToSQLString(t.Title),
		orm.ToSQLBool(t.Public),
		orm.ToSQLBool(t.Hidden),
		orm.ToSQLBool(t.Admin),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
	)
}

// DeleteAllACLResources deletes all ACLResource records in the database with optional filters
func DeleteAllACLResources(ctx context.Context, db *sql.DB, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(ACLResourceTableName),
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

// DeleteAllACLResourcesTx deletes all ACLResource records in the database with optional filters using the provided transaction
func DeleteAllACLResourcesTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(ACLResourceTableName),
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

// DBDelete will delete this ACLResource record in the database
func (t *ACLResource) DBDelete(ctx context.Context, db *sql.DB) (bool, error) {
	q := "DELETE FROM `acl_resource` WHERE `id` = ?"
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

// DBDeleteTx will delete this ACLResource record in the database using the provided transaction
func (t *ACLResource) DBDeleteTx(ctx context.Context, tx *sql.Tx) (bool, error) {
	q := "DELETE FROM `acl_resource` WHERE `id` = ?"
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

// DBUpdate will update the ACLResource record in the database
func (t *ACLResource) DBUpdate(ctx context.Context, db *sql.DB) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `acl_resource` SET `checksum`=?,`urn`=?,`description`=?,`title`=?,`public`=?,`hidden`=?,`admin`=?,`created_at`=?,`updated_at`=? WHERE `id`=?"
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.Urn),
		orm.ToSQLString(t.Description),
		orm.ToSQLString(t.Title),
		orm.ToSQLBool(t.Public),
		orm.ToSQLBool(t.Hidden),
		orm.ToSQLBool(t.Admin),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
		orm.ToSQLString(t.ID),
	)
}

// DBUpdateTx will update the ACLResource record in the database using the provided transaction
func (t *ACLResource) DBUpdateTx(ctx context.Context, tx *sql.Tx) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `acl_resource` SET `checksum`=?,`urn`=?,`description`=?,`title`=?,`public`=?,`hidden`=?,`admin`=?,`created_at`=?,`updated_at`=? WHERE `id`=?"
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.Urn),
		orm.ToSQLString(t.Description),
		orm.ToSQLString(t.Title),
		orm.ToSQLBool(t.Public),
		orm.ToSQLBool(t.Hidden),
		orm.ToSQLBool(t.Admin),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
		orm.ToSQLString(t.ID),
	)
}

// DBUpsert will upsert the ACLResource record in the database
func (t *ACLResource) DBUpsert(ctx context.Context, db *sql.DB, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `acl_resource` (`acl_resource`.`id`,`acl_resource`.`checksum`,`acl_resource`.`urn`,`acl_resource`.`description`,`acl_resource`.`title`,`acl_resource`.`public`,`acl_resource`.`hidden`,`acl_resource`.`admin`,`acl_resource`.`created_at`,`acl_resource`.`updated_at`) VALUES (?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `acl_resource` (`acl_resource`.`id`,`acl_resource`.`checksum`,`acl_resource`.`urn`,`acl_resource`.`description`,`acl_resource`.`title`,`acl_resource`.`public`,`acl_resource`.`hidden`,`acl_resource`.`admin`,`acl_resource`.`created_at`,`acl_resource`.`updated_at`) VALUES (?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`urn`=VALUES(`urn`),`description`=VALUES(`description`),`title`=VALUES(`title`),`public`=VALUES(`public`),`hidden`=VALUES(`hidden`),`admin`=VALUES(`admin`),`created_at`=VALUES(`created_at`),`updated_at`=VALUES(`updated_at`)"
	}
	r, err := db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.Urn),
		orm.ToSQLString(t.Description),
		orm.ToSQLString(t.Title),
		orm.ToSQLBool(t.Public),
		orm.ToSQLBool(t.Hidden),
		orm.ToSQLBool(t.Admin),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBUpsertTx will upsert the ACLResource record in the database using the provided transaction
func (t *ACLResource) DBUpsertTx(ctx context.Context, tx *sql.Tx, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `acl_resource` (`acl_resource`.`id`,`acl_resource`.`checksum`,`acl_resource`.`urn`,`acl_resource`.`description`,`acl_resource`.`title`,`acl_resource`.`public`,`acl_resource`.`hidden`,`acl_resource`.`admin`,`acl_resource`.`created_at`,`acl_resource`.`updated_at`) VALUES (?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `acl_resource` (`acl_resource`.`id`,`acl_resource`.`checksum`,`acl_resource`.`urn`,`acl_resource`.`description`,`acl_resource`.`title`,`acl_resource`.`public`,`acl_resource`.`hidden`,`acl_resource`.`admin`,`acl_resource`.`created_at`,`acl_resource`.`updated_at`) VALUES (?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`urn`=VALUES(`urn`),`description`=VALUES(`description`),`title`=VALUES(`title`),`public`=VALUES(`public`),`hidden`=VALUES(`hidden`),`admin`=VALUES(`admin`),`created_at`=VALUES(`created_at`),`updated_at`=VALUES(`updated_at`)"
	}
	r, err := tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.Urn),
		orm.ToSQLString(t.Description),
		orm.ToSQLString(t.Title),
		orm.ToSQLBool(t.Public),
		orm.ToSQLBool(t.Hidden),
		orm.ToSQLBool(t.Admin),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBFindOne will find a ACLResource record in the database with the primary key
func (t *ACLResource) DBFindOne(ctx context.Context, db *sql.DB, value string) (bool, error) {
	q := "SELECT `acl_resource`.`id`,`acl_resource`.`checksum`,`acl_resource`.`urn`,`acl_resource`.`description`,`acl_resource`.`title`,`acl_resource`.`public`,`acl_resource`.`hidden`,`acl_resource`.`admin`,`acl_resource`.`created_at`,`acl_resource`.`updated_at` FROM `acl_resource` WHERE `id` = ? LIMIT 1"
	row := db.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _Urn sql.NullString
	var _Description sql.NullString
	var _Title sql.NullString
	var _Public sql.NullBool
	var _Hidden sql.NullBool
	var _Admin sql.NullBool
	var _CreatedAt sql.NullInt64
	var _UpdatedAt sql.NullInt64
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_Urn,
		&_Description,
		&_Title,
		&_Public,
		&_Hidden,
		&_Admin,
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
	if _Checksum.Valid {
		t.SetChecksum(_Checksum.String)
	}
	if _Urn.Valid {
		t.SetUrn(_Urn.String)
	}
	if _Description.Valid {
		t.SetDescription(_Description.String)
	}
	if _Title.Valid {
		t.SetTitle(_Title.String)
	}
	if _Public.Valid {
		t.SetPublic(_Public.Bool)
	}
	if _Hidden.Valid {
		t.SetHidden(_Hidden.Bool)
	}
	if _Admin.Valid {
		t.SetAdmin(_Admin.Bool)
	}
	if _CreatedAt.Valid {
		t.SetCreatedAt(_CreatedAt.Int64)
	}
	if _UpdatedAt.Valid {
		t.SetUpdatedAt(_UpdatedAt.Int64)
	}
	return true, nil
}

// DBFindOneTx will find a ACLResource record in the database with the primary key using the provided transaction
func (t *ACLResource) DBFindOneTx(ctx context.Context, tx *sql.Tx, value string) (bool, error) {
	q := "SELECT `acl_resource`.`id`,`acl_resource`.`checksum`,`acl_resource`.`urn`,`acl_resource`.`description`,`acl_resource`.`title`,`acl_resource`.`public`,`acl_resource`.`hidden`,`acl_resource`.`admin`,`acl_resource`.`created_at`,`acl_resource`.`updated_at` FROM `acl_resource` WHERE `id` = ? LIMIT 1"
	row := tx.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _Urn sql.NullString
	var _Description sql.NullString
	var _Title sql.NullString
	var _Public sql.NullBool
	var _Hidden sql.NullBool
	var _Admin sql.NullBool
	var _CreatedAt sql.NullInt64
	var _UpdatedAt sql.NullInt64
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_Urn,
		&_Description,
		&_Title,
		&_Public,
		&_Hidden,
		&_Admin,
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
	if _Checksum.Valid {
		t.SetChecksum(_Checksum.String)
	}
	if _Urn.Valid {
		t.SetUrn(_Urn.String)
	}
	if _Description.Valid {
		t.SetDescription(_Description.String)
	}
	if _Title.Valid {
		t.SetTitle(_Title.String)
	}
	if _Public.Valid {
		t.SetPublic(_Public.Bool)
	}
	if _Hidden.Valid {
		t.SetHidden(_Hidden.Bool)
	}
	if _Admin.Valid {
		t.SetAdmin(_Admin.Bool)
	}
	if _CreatedAt.Valid {
		t.SetCreatedAt(_CreatedAt.Int64)
	}
	if _UpdatedAt.Valid {
		t.SetUpdatedAt(_UpdatedAt.Int64)
	}
	return true, nil
}

// FindACLResources will find a ACLResource record in the database with the provided parameters
func FindACLResources(ctx context.Context, db *sql.DB, _params ...interface{}) ([]*ACLResource, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("urn"),
		orm.Column("description"),
		orm.Column("title"),
		orm.Column("public"),
		orm.Column("hidden"),
		orm.Column("admin"),
		orm.Column("created_at"),
		orm.Column("updated_at"),
		orm.Table(ACLResourceTableName),
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
	results := make([]*ACLResource, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _Urn sql.NullString
		var _Description sql.NullString
		var _Title sql.NullString
		var _Public sql.NullBool
		var _Hidden sql.NullBool
		var _Admin sql.NullBool
		var _CreatedAt sql.NullInt64
		var _UpdatedAt sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_Urn,
			&_Description,
			&_Title,
			&_Public,
			&_Hidden,
			&_Admin,
			&_CreatedAt,
			&_UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		t := &ACLResource{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _Urn.Valid {
			t.SetUrn(_Urn.String)
		}
		if _Description.Valid {
			t.SetDescription(_Description.String)
		}
		if _Title.Valid {
			t.SetTitle(_Title.String)
		}
		if _Public.Valid {
			t.SetPublic(_Public.Bool)
		}
		if _Hidden.Valid {
			t.SetHidden(_Hidden.Bool)
		}
		if _Admin.Valid {
			t.SetAdmin(_Admin.Bool)
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

// FindACLResourcesTx will find a ACLResource record in the database with the provided parameters using the provided transaction
func FindACLResourcesTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) ([]*ACLResource, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("urn"),
		orm.Column("description"),
		orm.Column("title"),
		orm.Column("public"),
		orm.Column("hidden"),
		orm.Column("admin"),
		orm.Column("created_at"),
		orm.Column("updated_at"),
		orm.Table(ACLResourceTableName),
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
	results := make([]*ACLResource, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _Urn sql.NullString
		var _Description sql.NullString
		var _Title sql.NullString
		var _Public sql.NullBool
		var _Hidden sql.NullBool
		var _Admin sql.NullBool
		var _CreatedAt sql.NullInt64
		var _UpdatedAt sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_Urn,
			&_Description,
			&_Title,
			&_Public,
			&_Hidden,
			&_Admin,
			&_CreatedAt,
			&_UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		t := &ACLResource{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _Urn.Valid {
			t.SetUrn(_Urn.String)
		}
		if _Description.Valid {
			t.SetDescription(_Description.String)
		}
		if _Title.Valid {
			t.SetTitle(_Title.String)
		}
		if _Public.Valid {
			t.SetPublic(_Public.Bool)
		}
		if _Hidden.Valid {
			t.SetHidden(_Hidden.Bool)
		}
		if _Admin.Valid {
			t.SetAdmin(_Admin.Bool)
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

// DBFind will find a ACLResource record in the database with the provided parameters
func (t *ACLResource) DBFind(ctx context.Context, db *sql.DB, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("urn"),
		orm.Column("description"),
		orm.Column("title"),
		orm.Column("public"),
		orm.Column("hidden"),
		orm.Column("admin"),
		orm.Column("created_at"),
		orm.Column("updated_at"),
		orm.Table(ACLResourceTableName),
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
	var _Urn sql.NullString
	var _Description sql.NullString
	var _Title sql.NullString
	var _Public sql.NullBool
	var _Hidden sql.NullBool
	var _Admin sql.NullBool
	var _CreatedAt sql.NullInt64
	var _UpdatedAt sql.NullInt64
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_Urn,
		&_Description,
		&_Title,
		&_Public,
		&_Hidden,
		&_Admin,
		&_CreatedAt,
		&_UpdatedAt,
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
	if _Urn.Valid {
		t.SetUrn(_Urn.String)
	}
	if _Description.Valid {
		t.SetDescription(_Description.String)
	}
	if _Title.Valid {
		t.SetTitle(_Title.String)
	}
	if _Public.Valid {
		t.SetPublic(_Public.Bool)
	}
	if _Hidden.Valid {
		t.SetHidden(_Hidden.Bool)
	}
	if _Admin.Valid {
		t.SetAdmin(_Admin.Bool)
	}
	if _CreatedAt.Valid {
		t.SetCreatedAt(_CreatedAt.Int64)
	}
	if _UpdatedAt.Valid {
		t.SetUpdatedAt(_UpdatedAt.Int64)
	}
	return true, nil
}

// DBFindTx will find a ACLResource record in the database with the provided parameters using the provided transaction
func (t *ACLResource) DBFindTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("urn"),
		orm.Column("description"),
		orm.Column("title"),
		orm.Column("public"),
		orm.Column("hidden"),
		orm.Column("admin"),
		orm.Column("created_at"),
		orm.Column("updated_at"),
		orm.Table(ACLResourceTableName),
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
	var _Urn sql.NullString
	var _Description sql.NullString
	var _Title sql.NullString
	var _Public sql.NullBool
	var _Hidden sql.NullBool
	var _Admin sql.NullBool
	var _CreatedAt sql.NullInt64
	var _UpdatedAt sql.NullInt64
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_Urn,
		&_Description,
		&_Title,
		&_Public,
		&_Hidden,
		&_Admin,
		&_CreatedAt,
		&_UpdatedAt,
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
	if _Urn.Valid {
		t.SetUrn(_Urn.String)
	}
	if _Description.Valid {
		t.SetDescription(_Description.String)
	}
	if _Title.Valid {
		t.SetTitle(_Title.String)
	}
	if _Public.Valid {
		t.SetPublic(_Public.Bool)
	}
	if _Hidden.Valid {
		t.SetHidden(_Hidden.Bool)
	}
	if _Admin.Valid {
		t.SetAdmin(_Admin.Bool)
	}
	if _CreatedAt.Valid {
		t.SetCreatedAt(_CreatedAt.Int64)
	}
	if _UpdatedAt.Valid {
		t.SetUpdatedAt(_UpdatedAt.Int64)
	}
	return true, nil
}

// CountACLResources will find the count of ACLResource records in the database
func CountACLResources(ctx context.Context, db *sql.DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(ACLResourceTableName),
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

// CountACLResourcesTx will find the count of ACLResource records in the database using the provided transaction
func CountACLResourcesTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(ACLResourceTableName),
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

// DBCount will find the count of ACLResource records in the database
func (t *ACLResource) DBCount(ctx context.Context, db *sql.DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(ACLResourceTableName),
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

// DBCountTx will find the count of ACLResource records in the database using the provided transaction
func (t *ACLResource) DBCountTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(ACLResourceTableName),
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

// DBExists will return true if the ACLResource record exists in the database
func (t *ACLResource) DBExists(ctx context.Context, db *sql.DB) (bool, error) {
	q := "SELECT `id` FROM `acl_resource` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := db.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// DBExistsTx will return true if the ACLResource record exists in the database using the provided transaction
func (t *ACLResource) DBExistsTx(ctx context.Context, tx *sql.Tx) (bool, error) {
	q := "SELECT `id` FROM `acl_resource` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := tx.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// PrimaryKeyColumn returns the column name for the primary key
func (t *ACLResource) PrimaryKeyColumn() string {
	return ACLResourceColumnID
}

// PrimaryKeyColumnType returns the primary key column Go type as a string
func (t *ACLResource) PrimaryKeyColumnType() string {
	return "string"
}

// PrimaryKey returns the primary key column value
func (t *ACLResource) PrimaryKey() interface{} {
	return t.ID
}
