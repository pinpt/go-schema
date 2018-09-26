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

var _ Model = (*ACLRole)(nil)
var _ CSVWriter = (*ACLRole)(nil)
var _ JSONWriter = (*ACLRole)(nil)
var _ Checksum = (*ACLRole)(nil)

// ACLRoleTableName is the name of the table in SQL
const ACLRoleTableName = "acl_role"

var ACLRoleColumns = []string{
	"id",
	"checksum",
	"name",
	"description",
	"admin_user_id",
	"customer_id",
	"created_at",
	"updated_at",
}

// ACLRole table
type ACLRole struct {
	AdminUserID *string `json:"admin_user_id,omitempty"`
	Checksum    *string `json:"checksum,omitempty"`
	CreatedAt   int64   `json:"created_at"`
	CustomerID  *string `json:"customer_id,omitempty"`
	Description *string `json:"description,omitempty"`
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	UpdatedAt   *int64  `json:"updated_at,omitempty"`
}

// TableName returns the SQL table name for ACLRole and satifies the Model interface
func (t *ACLRole) TableName() string {
	return ACLRoleTableName
}

// ToCSV will serialize the ACLRole instance to a CSV compatible array of strings
func (t *ACLRole) ToCSV() []string {
	return []string{
		t.ID,
		t.CalculateChecksum(),
		t.Name,
		toCSVString(t.Description),
		toCSVString(t.AdminUserID),
		toCSVString(t.CustomerID),
		toCSVString(t.CreatedAt),
		toCSVString(t.UpdatedAt),
	}
}

// WriteCSV will serialize the ACLRole instance to the writer as CSV and satisfies the CSVWriter interface
func (t *ACLRole) WriteCSV(w *csv.Writer) error {
	return w.Write(t.ToCSV())
}

// WriteJSON will serialize the ACLRole instance to the writer as JSON and satisfies the JSONWriter interface
func (t *ACLRole) WriteJSON(w io.Writer, indent ...bool) error {
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

// NewACLRoleReader creates a JSON reader which can read in ACLRole objects serialized as JSON either as an array, single object or json new lines
// and writes each ACLRole to the channel provided
func NewACLRoleReader(r io.Reader, ch chan<- ACLRole) error {
	return orm.Deserialize(r, func(buf json.RawMessage) error {
		dec := json.NewDecoder(bytes.NewBuffer(buf))
		e := ACLRole{}
		if err := dec.Decode(&e); err != nil {
			return err
		}
		ch <- e
		return nil
	})
}

// NewCSVACLRoleReaderDir will read the reader as CSV and emit each record to the channel provided
func NewCSVACLRoleReader(r io.Reader, ch chan<- ACLRole) error {
	cr := csv.NewReader(r)
	for {
		record, err := cr.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		ch <- ACLRole{
			ID:          record[0],
			Checksum:    fromStringPointer(record[1]),
			Name:        record[2],
			Description: fromStringPointer(record[3]),
			AdminUserID: fromStringPointer(record[4]),
			CustomerID:  fromStringPointer(record[5]),
			CreatedAt:   fromCSVInt64(record[6]),
			UpdatedAt:   fromCSVInt64Pointer(record[7]),
		}
	}
	return nil
}

// NewCSVACLRoleReaderFile will read the file as a CSV and emit each record to the channel provided
func NewCSVACLRoleReaderFile(fp string, ch chan<- ACLRole) error {
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
	return NewCSVACLRoleReader(fc, ch)
}

// NewCSVACLRoleReaderDir will read the acl_role.csv.gz file as a CSV and emit each record to the channel provided
func NewCSVACLRoleReaderDir(dir string, ch chan<- ACLRole) error {
	return NewCSVACLRoleReaderFile(filepath.Join(dir, "acl_role.csv.gz"), ch)
}

// ACLRoleCSVDeduper is a function callback which takes the existing value (a) and the new value (b)
// and the return value should be the one to use (or a new one, if applicable). return nil
// to skip processing of this record
type ACLRoleCSVDeduper func(a ACLRole, b ACLRole) *ACLRole

// ACLRoleCSVDedupeDisabled is set on whether the CSV writer should de-dupe values by key
var ACLRoleCSVDedupeDisabled bool

// NewACLRoleCSVWriterSize creates a batch writer that will write each ACLRole into a CSV file
// this method will automatically de-duplicate entries using the primary key. if the checksum
// for a newer item with the same primary key doesn't match a previously sent item, the newer
// one will be used
func NewACLRoleCSVWriterSize(w io.Writer, size int, dedupers ...ACLRoleCSVDeduper) (chan ACLRole, chan bool, error) {
	cw := csv.NewWriter(w)
	ch := make(chan ACLRole, size)
	done := make(chan bool)
	go func() {
		defer func() { done <- true }()
		dodedupe := !ACLRoleCSVDedupeDisabled
		var kv map[string]*ACLRole
		var deduper ACLRoleCSVDeduper
		if dedupers != nil && len(dedupers) > 0 {
			deduper = dedupers[0]
			dodedupe = true
		}
		if dodedupe {
			kv = make(map[string]*ACLRole)
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

// ACLRoleCSVDefaultSize is the default channel buffer size if not provided
var ACLRoleCSVDefaultSize = 100

// NewACLRoleCSVWriter creates a batch writer that will write each ACLRole into a CSV file
func NewACLRoleCSVWriter(w io.Writer, dedupers ...ACLRoleCSVDeduper) (chan ACLRole, chan bool, error) {
	return NewACLRoleCSVWriterSize(w, ACLRoleCSVDefaultSize, dedupers...)
}

// NewACLRoleCSVWriterDir creates a batch writer that will write each ACLRole into a CSV file named acl_role.csv.gz in dir
func NewACLRoleCSVWriterDir(dir string, dedupers ...ACLRoleCSVDeduper) (chan ACLRole, chan bool, error) {
	return NewACLRoleCSVWriterFile(filepath.Join(dir, "acl_role.csv.gz"), dedupers...)
}

// NewACLRoleCSVWriterFile creates a batch writer that will write each ACLRole into a CSV file
func NewACLRoleCSVWriterFile(fn string, dedupers ...ACLRoleCSVDeduper) (chan ACLRole, chan bool, error) {
	f, err := os.Create(fn)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening CSV file %s. %v", fn, err)
	}
	var fc io.WriteCloser = f
	if filepath.Ext(fn) == ".gz" {
		w, _ := gzip.NewWriterLevel(f, gzip.BestCompression)
		fc = w
	}
	ch, done, err := NewACLRoleCSVWriter(fc, dedupers...)
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

type ACLRoleDBAction func(ctx context.Context, db *sql.DB, record ACLRole) error

// NewACLRoleDBWriterSize creates a DB writer that will write each issue into the DB
func NewACLRoleDBWriterSize(ctx context.Context, db *sql.DB, errors chan<- error, size int, actions ...ACLRoleDBAction) (chan ACLRole, chan bool, error) {
	ch := make(chan ACLRole, size)
	done := make(chan bool)
	var action ACLRoleDBAction
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

// NewACLRoleDBWriter creates a DB writer that will write each issue into the DB
func NewACLRoleDBWriter(ctx context.Context, db *sql.DB, errors chan<- error, actions ...ACLRoleDBAction) (chan ACLRole, chan bool, error) {
	return NewACLRoleDBWriterSize(ctx, db, errors, 100, actions...)
}

// ACLRoleColumnID is the ID SQL column name for the ACLRole table
const ACLRoleColumnID = "id"

// ACLRoleEscapedColumnID is the escaped ID SQL column name for the ACLRole table
const ACLRoleEscapedColumnID = "`id`"

// ACLRoleColumnChecksum is the Checksum SQL column name for the ACLRole table
const ACLRoleColumnChecksum = "checksum"

// ACLRoleEscapedColumnChecksum is the escaped Checksum SQL column name for the ACLRole table
const ACLRoleEscapedColumnChecksum = "`checksum`"

// ACLRoleColumnName is the Name SQL column name for the ACLRole table
const ACLRoleColumnName = "name"

// ACLRoleEscapedColumnName is the escaped Name SQL column name for the ACLRole table
const ACLRoleEscapedColumnName = "`name`"

// ACLRoleColumnDescription is the Description SQL column name for the ACLRole table
const ACLRoleColumnDescription = "description"

// ACLRoleEscapedColumnDescription is the escaped Description SQL column name for the ACLRole table
const ACLRoleEscapedColumnDescription = "`description`"

// ACLRoleColumnAdminUserID is the AdminUserID SQL column name for the ACLRole table
const ACLRoleColumnAdminUserID = "admin_user_id"

// ACLRoleEscapedColumnAdminUserID is the escaped AdminUserID SQL column name for the ACLRole table
const ACLRoleEscapedColumnAdminUserID = "`admin_user_id`"

// ACLRoleColumnCustomerID is the CustomerID SQL column name for the ACLRole table
const ACLRoleColumnCustomerID = "customer_id"

// ACLRoleEscapedColumnCustomerID is the escaped CustomerID SQL column name for the ACLRole table
const ACLRoleEscapedColumnCustomerID = "`customer_id`"

// ACLRoleColumnCreatedAt is the CreatedAt SQL column name for the ACLRole table
const ACLRoleColumnCreatedAt = "created_at"

// ACLRoleEscapedColumnCreatedAt is the escaped CreatedAt SQL column name for the ACLRole table
const ACLRoleEscapedColumnCreatedAt = "`created_at`"

// ACLRoleColumnUpdatedAt is the UpdatedAt SQL column name for the ACLRole table
const ACLRoleColumnUpdatedAt = "updated_at"

// ACLRoleEscapedColumnUpdatedAt is the escaped UpdatedAt SQL column name for the ACLRole table
const ACLRoleEscapedColumnUpdatedAt = "`updated_at`"

// GetID will return the ACLRole ID value
func (t *ACLRole) GetID() string {
	return t.ID
}

// SetID will set the ACLRole ID value
func (t *ACLRole) SetID(v string) {
	t.ID = v
}

// FindACLRoleByID will find a ACLRole by ID
func FindACLRoleByID(ctx context.Context, db *sql.DB, value string) (*ACLRole, error) {
	q := "SELECT `acl_role`.`id`,`acl_role`.`checksum`,`acl_role`.`name`,`acl_role`.`description`,`acl_role`.`admin_user_id`,`acl_role`.`customer_id`,`acl_role`.`created_at`,`acl_role`.`updated_at` FROM `acl_role` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _Name sql.NullString
	var _Description sql.NullString
	var _AdminUserID sql.NullString
	var _CustomerID sql.NullString
	var _CreatedAt sql.NullInt64
	var _UpdatedAt sql.NullInt64
	err := db.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_Name,
		&_Description,
		&_AdminUserID,
		&_CustomerID,
		&_CreatedAt,
		&_UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &ACLRole{}
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
	if _AdminUserID.Valid {
		t.SetAdminUserID(_AdminUserID.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _CreatedAt.Valid {
		t.SetCreatedAt(_CreatedAt.Int64)
	}
	if _UpdatedAt.Valid {
		t.SetUpdatedAt(_UpdatedAt.Int64)
	}
	return t, nil
}

// FindACLRoleByIDTx will find a ACLRole by ID using the provided transaction
func FindACLRoleByIDTx(ctx context.Context, tx *sql.Tx, value string) (*ACLRole, error) {
	q := "SELECT `acl_role`.`id`,`acl_role`.`checksum`,`acl_role`.`name`,`acl_role`.`description`,`acl_role`.`admin_user_id`,`acl_role`.`customer_id`,`acl_role`.`created_at`,`acl_role`.`updated_at` FROM `acl_role` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _Name sql.NullString
	var _Description sql.NullString
	var _AdminUserID sql.NullString
	var _CustomerID sql.NullString
	var _CreatedAt sql.NullInt64
	var _UpdatedAt sql.NullInt64
	err := tx.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_Name,
		&_Description,
		&_AdminUserID,
		&_CustomerID,
		&_CreatedAt,
		&_UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &ACLRole{}
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
	if _AdminUserID.Valid {
		t.SetAdminUserID(_AdminUserID.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _CreatedAt.Valid {
		t.SetCreatedAt(_CreatedAt.Int64)
	}
	if _UpdatedAt.Valid {
		t.SetUpdatedAt(_UpdatedAt.Int64)
	}
	return t, nil
}

// GetChecksum will return the ACLRole Checksum value
func (t *ACLRole) GetChecksum() string {
	if t.Checksum == nil {
		return ""
	}
	return *t.Checksum
}

// SetChecksum will set the ACLRole Checksum value
func (t *ACLRole) SetChecksum(v string) {
	t.Checksum = &v
}

// GetName will return the ACLRole Name value
func (t *ACLRole) GetName() string {
	return t.Name
}

// SetName will set the ACLRole Name value
func (t *ACLRole) SetName(v string) {
	t.Name = v
}

// FindACLRolesByName will find all ACLRoles by the Name value
func FindACLRolesByName(ctx context.Context, db *sql.DB, value string) ([]*ACLRole, error) {
	q := "SELECT `acl_role`.`id`,`acl_role`.`checksum`,`acl_role`.`name`,`acl_role`.`description`,`acl_role`.`admin_user_id`,`acl_role`.`customer_id`,`acl_role`.`created_at`,`acl_role`.`updated_at` FROM `acl_role` WHERE `name` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*ACLRole, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _Name sql.NullString
		var _Description sql.NullString
		var _AdminUserID sql.NullString
		var _CustomerID sql.NullString
		var _CreatedAt sql.NullInt64
		var _UpdatedAt sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_Name,
			&_Description,
			&_AdminUserID,
			&_CustomerID,
			&_CreatedAt,
			&_UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		t := &ACLRole{}
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
		if _AdminUserID.Valid {
			t.SetAdminUserID(_AdminUserID.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
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

// FindACLRolesByNameTx will find all ACLRoles by the Name value using the provided transaction
func FindACLRolesByNameTx(ctx context.Context, tx *sql.Tx, value string) ([]*ACLRole, error) {
	q := "SELECT `acl_role`.`id`,`acl_role`.`checksum`,`acl_role`.`name`,`acl_role`.`description`,`acl_role`.`admin_user_id`,`acl_role`.`customer_id`,`acl_role`.`created_at`,`acl_role`.`updated_at` FROM `acl_role` WHERE `name` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*ACLRole, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _Name sql.NullString
		var _Description sql.NullString
		var _AdminUserID sql.NullString
		var _CustomerID sql.NullString
		var _CreatedAt sql.NullInt64
		var _UpdatedAt sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_Name,
			&_Description,
			&_AdminUserID,
			&_CustomerID,
			&_CreatedAt,
			&_UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		t := &ACLRole{}
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
		if _AdminUserID.Valid {
			t.SetAdminUserID(_AdminUserID.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
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

// GetDescription will return the ACLRole Description value
func (t *ACLRole) GetDescription() string {
	if t.Description == nil {
		return ""
	}
	return *t.Description
}

// SetDescription will set the ACLRole Description value
func (t *ACLRole) SetDescription(v string) {
	t.Description = &v
}

// GetAdminUserID will return the ACLRole AdminUserID value
func (t *ACLRole) GetAdminUserID() string {
	if t.AdminUserID == nil {
		return ""
	}
	return *t.AdminUserID
}

// SetAdminUserID will set the ACLRole AdminUserID value
func (t *ACLRole) SetAdminUserID(v string) {
	t.AdminUserID = &v
}

// GetCustomerID will return the ACLRole CustomerID value
func (t *ACLRole) GetCustomerID() string {
	if t.CustomerID == nil {
		return ""
	}
	return *t.CustomerID
}

// SetCustomerID will set the ACLRole CustomerID value
func (t *ACLRole) SetCustomerID(v string) {
	t.CustomerID = &v
}

// FindACLRolesByCustomerID will find all ACLRoles by the CustomerID value
func FindACLRolesByCustomerID(ctx context.Context, db *sql.DB, value string) ([]*ACLRole, error) {
	q := "SELECT `acl_role`.`id`,`acl_role`.`checksum`,`acl_role`.`name`,`acl_role`.`description`,`acl_role`.`admin_user_id`,`acl_role`.`customer_id`,`acl_role`.`created_at`,`acl_role`.`updated_at` FROM `acl_role` WHERE `customer_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*ACLRole, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _Name sql.NullString
		var _Description sql.NullString
		var _AdminUserID sql.NullString
		var _CustomerID sql.NullString
		var _CreatedAt sql.NullInt64
		var _UpdatedAt sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_Name,
			&_Description,
			&_AdminUserID,
			&_CustomerID,
			&_CreatedAt,
			&_UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		t := &ACLRole{}
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
		if _AdminUserID.Valid {
			t.SetAdminUserID(_AdminUserID.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
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

// FindACLRolesByCustomerIDTx will find all ACLRoles by the CustomerID value using the provided transaction
func FindACLRolesByCustomerIDTx(ctx context.Context, tx *sql.Tx, value string) ([]*ACLRole, error) {
	q := "SELECT `acl_role`.`id`,`acl_role`.`checksum`,`acl_role`.`name`,`acl_role`.`description`,`acl_role`.`admin_user_id`,`acl_role`.`customer_id`,`acl_role`.`created_at`,`acl_role`.`updated_at` FROM `acl_role` WHERE `customer_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*ACLRole, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _Name sql.NullString
		var _Description sql.NullString
		var _AdminUserID sql.NullString
		var _CustomerID sql.NullString
		var _CreatedAt sql.NullInt64
		var _UpdatedAt sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_Name,
			&_Description,
			&_AdminUserID,
			&_CustomerID,
			&_CreatedAt,
			&_UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		t := &ACLRole{}
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
		if _AdminUserID.Valid {
			t.SetAdminUserID(_AdminUserID.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
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

// GetCreatedAt will return the ACLRole CreatedAt value
func (t *ACLRole) GetCreatedAt() int64 {
	return t.CreatedAt
}

// SetCreatedAt will set the ACLRole CreatedAt value
func (t *ACLRole) SetCreatedAt(v int64) {
	t.CreatedAt = v
}

// GetUpdatedAt will return the ACLRole UpdatedAt value
func (t *ACLRole) GetUpdatedAt() int64 {
	if t.UpdatedAt == nil {
		return int64(0)
	}
	return *t.UpdatedAt
}

// SetUpdatedAt will set the ACLRole UpdatedAt value
func (t *ACLRole) SetUpdatedAt(v int64) {
	t.UpdatedAt = &v
}

func (t *ACLRole) toTimestamp(value time.Time) *timestamp.Timestamp {
	ts, _ := ptypes.TimestampProto(value)
	return ts
}

// DBCreateACLRoleTable will create the ACLRole table
func DBCreateACLRoleTable(ctx context.Context, db *sql.DB) error {
	q := "CREATE TABLE `acl_role` (`id`VARCHAR(64) NOT NULL PRIMARY KEY,`checksum`CHAR(64),`name` VARCHAR(100) NOT NULL,`description`TEXT,`admin_user_id` VARCHAR(64),`customer_id`VARCHAR(64),`created_at` BIGINT UNSIGNED NOT NULL,`updated_at` BIGINT UNSIGNED,INDEX acl_role_name_index (`name`),INDEX acl_role_customer_id_index (`customer_id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBCreateACLRoleTableTx will create the ACLRole table using the provided transction
func DBCreateACLRoleTableTx(ctx context.Context, tx *sql.Tx) error {
	q := "CREATE TABLE `acl_role` (`id`VARCHAR(64) NOT NULL PRIMARY KEY,`checksum`CHAR(64),`name` VARCHAR(100) NOT NULL,`description`TEXT,`admin_user_id` VARCHAR(64),`customer_id`VARCHAR(64),`created_at` BIGINT UNSIGNED NOT NULL,`updated_at` BIGINT UNSIGNED,INDEX acl_role_name_index (`name`),INDEX acl_role_customer_id_index (`customer_id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// DBDropACLRoleTable will drop the ACLRole table
func DBDropACLRoleTable(ctx context.Context, db *sql.DB) error {
	q := "DROP TABLE IF EXISTS `acl_role`"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBDropACLRoleTableTx will drop the ACLRole table using the provided transaction
func DBDropACLRoleTableTx(ctx context.Context, tx *sql.Tx) error {
	q := "DROP TABLE IF EXISTS `acl_role`"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// CalculateChecksum will calculate a checksum of the SHA1 of all field values
func (t *ACLRole) CalculateChecksum() string {
	return orm.HashStrings(
		orm.ToString(t.ID),
		orm.ToString(t.Name),
		orm.ToString(t.Description),
		orm.ToString(t.AdminUserID),
		orm.ToString(t.CustomerID),
		orm.ToString(t.CreatedAt),
		orm.ToString(t.UpdatedAt),
	)
}

// DBCreate will create a new ACLRole record in the database
func (t *ACLRole) DBCreate(ctx context.Context, db *sql.DB) (sql.Result, error) {
	q := "INSERT INTO `acl_role` (`acl_role`.`id`,`acl_role`.`checksum`,`acl_role`.`name`,`acl_role`.`description`,`acl_role`.`admin_user_id`,`acl_role`.`customer_id`,`acl_role`.`created_at`,`acl_role`.`updated_at`) VALUES (?,?,?,?,?,?,?,?)"
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
		orm.ToSQLString(t.AdminUserID),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
	)
}

// DBCreateTx will create a new ACLRole record in the database using the provided transaction
func (t *ACLRole) DBCreateTx(ctx context.Context, tx *sql.Tx) (sql.Result, error) {
	q := "INSERT INTO `acl_role` (`acl_role`.`id`,`acl_role`.`checksum`,`acl_role`.`name`,`acl_role`.`description`,`acl_role`.`admin_user_id`,`acl_role`.`customer_id`,`acl_role`.`created_at`,`acl_role`.`updated_at`) VALUES (?,?,?,?,?,?,?,?)"
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
		orm.ToSQLString(t.AdminUserID),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
	)
}

// DBCreateIgnoreDuplicate will upsert the ACLRole record in the database
func (t *ACLRole) DBCreateIgnoreDuplicate(ctx context.Context, db *sql.DB) (sql.Result, error) {
	q := "INSERT INTO `acl_role` (`acl_role`.`id`,`acl_role`.`checksum`,`acl_role`.`name`,`acl_role`.`description`,`acl_role`.`admin_user_id`,`acl_role`.`customer_id`,`acl_role`.`created_at`,`acl_role`.`updated_at`) VALUES (?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
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
		orm.ToSQLString(t.AdminUserID),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
	)
}

// DBCreateIgnoreDuplicateTx will upsert the ACLRole record in the database using the provided transaction
func (t *ACLRole) DBCreateIgnoreDuplicateTx(ctx context.Context, tx *sql.Tx) (sql.Result, error) {
	q := "INSERT INTO `acl_role` (`acl_role`.`id`,`acl_role`.`checksum`,`acl_role`.`name`,`acl_role`.`description`,`acl_role`.`admin_user_id`,`acl_role`.`customer_id`,`acl_role`.`created_at`,`acl_role`.`updated_at`) VALUES (?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
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
		orm.ToSQLString(t.AdminUserID),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
	)
}

// DeleteAllACLRoles deletes all ACLRole records in the database with optional filters
func DeleteAllACLRoles(ctx context.Context, db *sql.DB, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(ACLRoleTableName),
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

// DeleteAllACLRolesTx deletes all ACLRole records in the database with optional filters using the provided transaction
func DeleteAllACLRolesTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(ACLRoleTableName),
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

// DBDelete will delete this ACLRole record in the database
func (t *ACLRole) DBDelete(ctx context.Context, db *sql.DB) (bool, error) {
	q := "DELETE FROM `acl_role` WHERE `id` = ?"
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

// DBDeleteTx will delete this ACLRole record in the database using the provided transaction
func (t *ACLRole) DBDeleteTx(ctx context.Context, tx *sql.Tx) (bool, error) {
	q := "DELETE FROM `acl_role` WHERE `id` = ?"
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

// DBUpdate will update the ACLRole record in the database
func (t *ACLRole) DBUpdate(ctx context.Context, db *sql.DB) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `acl_role` SET `checksum`=?,`name`=?,`description`=?,`admin_user_id`=?,`customer_id`=?,`created_at`=?,`updated_at`=? WHERE `id`=?"
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.Description),
		orm.ToSQLString(t.AdminUserID),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
		orm.ToSQLString(t.ID),
	)
}

// DBUpdateTx will update the ACLRole record in the database using the provided transaction
func (t *ACLRole) DBUpdateTx(ctx context.Context, tx *sql.Tx) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `acl_role` SET `checksum`=?,`name`=?,`description`=?,`admin_user_id`=?,`customer_id`=?,`created_at`=?,`updated_at`=? WHERE `id`=?"
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.Description),
		orm.ToSQLString(t.AdminUserID),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
		orm.ToSQLString(t.ID),
	)
}

// DBUpsert will upsert the ACLRole record in the database
func (t *ACLRole) DBUpsert(ctx context.Context, db *sql.DB, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `acl_role` (`acl_role`.`id`,`acl_role`.`checksum`,`acl_role`.`name`,`acl_role`.`description`,`acl_role`.`admin_user_id`,`acl_role`.`customer_id`,`acl_role`.`created_at`,`acl_role`.`updated_at`) VALUES (?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `acl_role` (`acl_role`.`id`,`acl_role`.`checksum`,`acl_role`.`name`,`acl_role`.`description`,`acl_role`.`admin_user_id`,`acl_role`.`customer_id`,`acl_role`.`created_at`,`acl_role`.`updated_at`) VALUES (?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`name`=VALUES(`name`),`description`=VALUES(`description`),`admin_user_id`=VALUES(`admin_user_id`),`customer_id`=VALUES(`customer_id`),`created_at`=VALUES(`created_at`),`updated_at`=VALUES(`updated_at`)"
	}
	r, err := db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.Description),
		orm.ToSQLString(t.AdminUserID),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBUpsertTx will upsert the ACLRole record in the database using the provided transaction
func (t *ACLRole) DBUpsertTx(ctx context.Context, tx *sql.Tx, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `acl_role` (`acl_role`.`id`,`acl_role`.`checksum`,`acl_role`.`name`,`acl_role`.`description`,`acl_role`.`admin_user_id`,`acl_role`.`customer_id`,`acl_role`.`created_at`,`acl_role`.`updated_at`) VALUES (?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `acl_role` (`acl_role`.`id`,`acl_role`.`checksum`,`acl_role`.`name`,`acl_role`.`description`,`acl_role`.`admin_user_id`,`acl_role`.`customer_id`,`acl_role`.`created_at`,`acl_role`.`updated_at`) VALUES (?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`name`=VALUES(`name`),`description`=VALUES(`description`),`admin_user_id`=VALUES(`admin_user_id`),`customer_id`=VALUES(`customer_id`),`created_at`=VALUES(`created_at`),`updated_at`=VALUES(`updated_at`)"
	}
	r, err := tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.Description),
		orm.ToSQLString(t.AdminUserID),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBFindOne will find a ACLRole record in the database with the primary key
func (t *ACLRole) DBFindOne(ctx context.Context, db *sql.DB, value string) (bool, error) {
	q := "SELECT `acl_role`.`id`,`acl_role`.`checksum`,`acl_role`.`name`,`acl_role`.`description`,`acl_role`.`admin_user_id`,`acl_role`.`customer_id`,`acl_role`.`created_at`,`acl_role`.`updated_at` FROM `acl_role` WHERE `id` = ? LIMIT 1"
	row := db.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _Name sql.NullString
	var _Description sql.NullString
	var _AdminUserID sql.NullString
	var _CustomerID sql.NullString
	var _CreatedAt sql.NullInt64
	var _UpdatedAt sql.NullInt64
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_Name,
		&_Description,
		&_AdminUserID,
		&_CustomerID,
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
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _Description.Valid {
		t.SetDescription(_Description.String)
	}
	if _AdminUserID.Valid {
		t.SetAdminUserID(_AdminUserID.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _CreatedAt.Valid {
		t.SetCreatedAt(_CreatedAt.Int64)
	}
	if _UpdatedAt.Valid {
		t.SetUpdatedAt(_UpdatedAt.Int64)
	}
	return true, nil
}

// DBFindOneTx will find a ACLRole record in the database with the primary key using the provided transaction
func (t *ACLRole) DBFindOneTx(ctx context.Context, tx *sql.Tx, value string) (bool, error) {
	q := "SELECT `acl_role`.`id`,`acl_role`.`checksum`,`acl_role`.`name`,`acl_role`.`description`,`acl_role`.`admin_user_id`,`acl_role`.`customer_id`,`acl_role`.`created_at`,`acl_role`.`updated_at` FROM `acl_role` WHERE `id` = ? LIMIT 1"
	row := tx.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _Name sql.NullString
	var _Description sql.NullString
	var _AdminUserID sql.NullString
	var _CustomerID sql.NullString
	var _CreatedAt sql.NullInt64
	var _UpdatedAt sql.NullInt64
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_Name,
		&_Description,
		&_AdminUserID,
		&_CustomerID,
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
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _Description.Valid {
		t.SetDescription(_Description.String)
	}
	if _AdminUserID.Valid {
		t.SetAdminUserID(_AdminUserID.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _CreatedAt.Valid {
		t.SetCreatedAt(_CreatedAt.Int64)
	}
	if _UpdatedAt.Valid {
		t.SetUpdatedAt(_UpdatedAt.Int64)
	}
	return true, nil
}

// FindACLRoles will find a ACLRole record in the database with the provided parameters
func FindACLRoles(ctx context.Context, db *sql.DB, _params ...interface{}) ([]*ACLRole, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("name"),
		orm.Column("description"),
		orm.Column("admin_user_id"),
		orm.Column("customer_id"),
		orm.Column("created_at"),
		orm.Column("updated_at"),
		orm.Table(ACLRoleTableName),
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
	results := make([]*ACLRole, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _Name sql.NullString
		var _Description sql.NullString
		var _AdminUserID sql.NullString
		var _CustomerID sql.NullString
		var _CreatedAt sql.NullInt64
		var _UpdatedAt sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_Name,
			&_Description,
			&_AdminUserID,
			&_CustomerID,
			&_CreatedAt,
			&_UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		t := &ACLRole{}
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
		if _AdminUserID.Valid {
			t.SetAdminUserID(_AdminUserID.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
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

// FindACLRolesTx will find a ACLRole record in the database with the provided parameters using the provided transaction
func FindACLRolesTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) ([]*ACLRole, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("name"),
		orm.Column("description"),
		orm.Column("admin_user_id"),
		orm.Column("customer_id"),
		orm.Column("created_at"),
		orm.Column("updated_at"),
		orm.Table(ACLRoleTableName),
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
	results := make([]*ACLRole, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _Name sql.NullString
		var _Description sql.NullString
		var _AdminUserID sql.NullString
		var _CustomerID sql.NullString
		var _CreatedAt sql.NullInt64
		var _UpdatedAt sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_Name,
			&_Description,
			&_AdminUserID,
			&_CustomerID,
			&_CreatedAt,
			&_UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		t := &ACLRole{}
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
		if _AdminUserID.Valid {
			t.SetAdminUserID(_AdminUserID.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
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

// DBFind will find a ACLRole record in the database with the provided parameters
func (t *ACLRole) DBFind(ctx context.Context, db *sql.DB, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("name"),
		orm.Column("description"),
		orm.Column("admin_user_id"),
		orm.Column("customer_id"),
		orm.Column("created_at"),
		orm.Column("updated_at"),
		orm.Table(ACLRoleTableName),
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
	var _AdminUserID sql.NullString
	var _CustomerID sql.NullString
	var _CreatedAt sql.NullInt64
	var _UpdatedAt sql.NullInt64
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_Name,
		&_Description,
		&_AdminUserID,
		&_CustomerID,
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
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _Description.Valid {
		t.SetDescription(_Description.String)
	}
	if _AdminUserID.Valid {
		t.SetAdminUserID(_AdminUserID.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _CreatedAt.Valid {
		t.SetCreatedAt(_CreatedAt.Int64)
	}
	if _UpdatedAt.Valid {
		t.SetUpdatedAt(_UpdatedAt.Int64)
	}
	return true, nil
}

// DBFindTx will find a ACLRole record in the database with the provided parameters using the provided transaction
func (t *ACLRole) DBFindTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("name"),
		orm.Column("description"),
		orm.Column("admin_user_id"),
		orm.Column("customer_id"),
		orm.Column("created_at"),
		orm.Column("updated_at"),
		orm.Table(ACLRoleTableName),
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
	var _AdminUserID sql.NullString
	var _CustomerID sql.NullString
	var _CreatedAt sql.NullInt64
	var _UpdatedAt sql.NullInt64
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_Name,
		&_Description,
		&_AdminUserID,
		&_CustomerID,
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
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _Description.Valid {
		t.SetDescription(_Description.String)
	}
	if _AdminUserID.Valid {
		t.SetAdminUserID(_AdminUserID.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _CreatedAt.Valid {
		t.SetCreatedAt(_CreatedAt.Int64)
	}
	if _UpdatedAt.Valid {
		t.SetUpdatedAt(_UpdatedAt.Int64)
	}
	return true, nil
}

// CountACLRoles will find the count of ACLRole records in the database
func CountACLRoles(ctx context.Context, db *sql.DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(ACLRoleTableName),
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

// CountACLRolesTx will find the count of ACLRole records in the database using the provided transaction
func CountACLRolesTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(ACLRoleTableName),
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

// DBCount will find the count of ACLRole records in the database
func (t *ACLRole) DBCount(ctx context.Context, db *sql.DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(ACLRoleTableName),
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

// DBCountTx will find the count of ACLRole records in the database using the provided transaction
func (t *ACLRole) DBCountTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(ACLRoleTableName),
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

// DBExists will return true if the ACLRole record exists in the database
func (t *ACLRole) DBExists(ctx context.Context, db *sql.DB) (bool, error) {
	q := "SELECT `id` FROM `acl_role` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := db.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// DBExistsTx will return true if the ACLRole record exists in the database using the provided transaction
func (t *ACLRole) DBExistsTx(ctx context.Context, tx *sql.Tx) (bool, error) {
	q := "SELECT `id` FROM `acl_role` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := tx.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// PrimaryKeyColumn returns the column name for the primary key
func (t *ACLRole) PrimaryKeyColumn() string {
	return ACLRoleColumnID
}

// PrimaryKeyColumnType returns the primary key column Go type as a string
func (t *ACLRole) PrimaryKeyColumnType() string {
	return "string"
}

// PrimaryKey returns the primary key column value
func (t *ACLRole) PrimaryKey() interface{} {
	return t.ID
}
