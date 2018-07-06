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

var _ Model = (*ACLUserRole)(nil)
var _ CSVWriter = (*ACLUserRole)(nil)
var _ JSONWriter = (*ACLUserRole)(nil)
var _ Checksum = (*ACLUserRole)(nil)

// ACLUserRoleTableName is the name of the table in SQL
const ACLUserRoleTableName = "acl_user_role"

var ACLUserRoleColumns = []string{
	"id",
	"checksum",
	"user_id",
	"role_id",
	"customer_id",
	"admin_user_id",
	"created_at",
	"updated_at",
}

// ACLUserRole table
type ACLUserRole struct {
	AdminUserID string  `json:"admin_user_id"`
	Checksum    *string `json:"checksum,omitempty"`
	CreatedAt   int64   `json:"created_at"`
	CustomerID  string  `json:"customer_id"`
	ID          string  `json:"id"`
	RoleID      string  `json:"role_id"`
	UpdatedAt   *int64  `json:"updated_at,omitempty"`
	UserID      string  `json:"user_id"`
}

// TableName returns the SQL table name for ACLUserRole and satifies the Model interface
func (t *ACLUserRole) TableName() string {
	return ACLUserRoleTableName
}

// ToCSV will serialize the ACLUserRole instance to a CSV compatible array of strings
func (t *ACLUserRole) ToCSV() []string {
	return []string{
		t.ID,
		t.CalculateChecksum(),
		t.UserID,
		t.RoleID,
		t.CustomerID,
		t.AdminUserID,
		toCSVString(t.CreatedAt),
		toCSVString(t.UpdatedAt),
	}
}

// WriteCSV will serialize the ACLUserRole instance to the writer as CSV and satisfies the CSVWriter interface
func (t *ACLUserRole) WriteCSV(w *csv.Writer) error {
	return w.Write(t.ToCSV())
}

// WriteJSON will serialize the ACLUserRole instance to the writer as JSON and satisfies the JSONWriter interface
func (t *ACLUserRole) WriteJSON(w io.Writer, indent ...bool) error {
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

// NewACLUserRoleReader creates a JSON reader which can read in ACLUserRole objects serialized as JSON either as an array, single object or json new lines
// and writes each ACLUserRole to the channel provided
func NewACLUserRoleReader(r io.Reader, ch chan<- ACLUserRole) error {
	return orm.Deserialize(r, func(buf json.RawMessage) error {
		dec := json.NewDecoder(bytes.NewBuffer(buf))
		e := ACLUserRole{}
		if err := dec.Decode(&e); err != nil {
			return err
		}
		ch <- e
		return nil
	})
}

// NewCSVACLUserRoleReaderDir will read the reader as CSV and emit each record to the channel provided
func NewCSVACLUserRoleReader(r io.Reader, ch chan<- ACLUserRole) error {
	cr := csv.NewReader(r)
	for {
		record, err := cr.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		ch <- ACLUserRole{
			ID:          record[0],
			Checksum:    fromStringPointer(record[1]),
			UserID:      record[2],
			RoleID:      record[3],
			CustomerID:  record[4],
			AdminUserID: record[5],
			CreatedAt:   fromCSVInt64(record[6]),
			UpdatedAt:   fromCSVInt64Pointer(record[7]),
		}
	}
	return nil
}

// NewCSVACLUserRoleReaderFile will read the file as a CSV and emit each record to the channel provided
func NewCSVACLUserRoleReaderFile(fp string, ch chan<- ACLUserRole) error {
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
	return NewCSVACLUserRoleReader(fc, ch)
}

// NewCSVACLUserRoleReaderDir will read the acl_user_role.csv.gz file as a CSV and emit each record to the channel provided
func NewCSVACLUserRoleReaderDir(dir string, ch chan<- ACLUserRole) error {
	return NewCSVACLUserRoleReaderFile(filepath.Join(dir, "acl_user_role.csv.gz"), ch)
}

// ACLUserRoleCSVDeduper is a function callback which takes the existing value (a) and the new value (b)
// and the return value should be the one to use (or a new one, if applicable). return nil
// to skip processing of this record
type ACLUserRoleCSVDeduper func(a ACLUserRole, b ACLUserRole) *ACLUserRole

// ACLUserRoleCSVDedupeDisabled is set on whether the CSV writer should de-dupe values by key
var ACLUserRoleCSVDedupeDisabled bool

// NewACLUserRoleCSVWriterSize creates a batch writer that will write each ACLUserRole into a CSV file
// this method will automatically de-duplicate entries using the primary key. if the checksum
// for a newer item with the same primary key doesn't match a previously sent item, the newer
// one will be used
func NewACLUserRoleCSVWriterSize(w io.Writer, size int, dedupers ...ACLUserRoleCSVDeduper) (chan ACLUserRole, chan bool, error) {
	cw := csv.NewWriter(w)
	ch := make(chan ACLUserRole, size)
	done := make(chan bool)
	go func() {
		defer func() { done <- true }()
		dodedupe := !ACLUserRoleCSVDedupeDisabled
		var kv map[string]*ACLUserRole
		var deduper ACLUserRoleCSVDeduper
		if dedupers != nil && len(dedupers) > 0 {
			deduper = dedupers[0]
			dodedupe = true
		}
		if dodedupe {
			kv = make(map[string]*ACLUserRole)
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

// ACLUserRoleCSVDefaultSize is the default channel buffer size if not provided
var ACLUserRoleCSVDefaultSize = 100

// NewACLUserRoleCSVWriter creates a batch writer that will write each ACLUserRole into a CSV file
func NewACLUserRoleCSVWriter(w io.Writer, dedupers ...ACLUserRoleCSVDeduper) (chan ACLUserRole, chan bool, error) {
	return NewACLUserRoleCSVWriterSize(w, ACLUserRoleCSVDefaultSize, dedupers...)
}

// NewACLUserRoleCSVWriterDir creates a batch writer that will write each ACLUserRole into a CSV file named acl_user_role.csv.gz in dir
func NewACLUserRoleCSVWriterDir(dir string, dedupers ...ACLUserRoleCSVDeduper) (chan ACLUserRole, chan bool, error) {
	return NewACLUserRoleCSVWriterFile(filepath.Join(dir, "acl_user_role.csv.gz"), dedupers...)
}

// NewACLUserRoleCSVWriterFile creates a batch writer that will write each ACLUserRole into a CSV file
func NewACLUserRoleCSVWriterFile(fn string, dedupers ...ACLUserRoleCSVDeduper) (chan ACLUserRole, chan bool, error) {
	f, err := os.Create(fn)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening CSV file %s. %v", fn, err)
	}
	var fc io.WriteCloser = f
	if filepath.Ext(fn) == ".gz" {
		w, _ := gzip.NewWriterLevel(f, gzip.BestCompression)
		fc = w
	}
	ch, done, err := NewACLUserRoleCSVWriter(fc, dedupers...)
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

type ACLUserRoleDBAction func(ctx context.Context, db *sql.DB, record ACLUserRole) error

// NewACLUserRoleDBWriterSize creates a DB writer that will write each issue into the DB
func NewACLUserRoleDBWriterSize(ctx context.Context, db *sql.DB, errors chan<- error, size int, actions ...ACLUserRoleDBAction) (chan ACLUserRole, chan bool, error) {
	ch := make(chan ACLUserRole, size)
	done := make(chan bool)
	var action ACLUserRoleDBAction
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

// NewACLUserRoleDBWriter creates a DB writer that will write each issue into the DB
func NewACLUserRoleDBWriter(ctx context.Context, db *sql.DB, errors chan<- error, actions ...ACLUserRoleDBAction) (chan ACLUserRole, chan bool, error) {
	return NewACLUserRoleDBWriterSize(ctx, db, errors, 100, actions...)
}

// ACLUserRoleColumnID is the ID SQL column name for the ACLUserRole table
const ACLUserRoleColumnID = "id"

// ACLUserRoleEscapedColumnID is the escaped ID SQL column name for the ACLUserRole table
const ACLUserRoleEscapedColumnID = "`id`"

// ACLUserRoleColumnChecksum is the Checksum SQL column name for the ACLUserRole table
const ACLUserRoleColumnChecksum = "checksum"

// ACLUserRoleEscapedColumnChecksum is the escaped Checksum SQL column name for the ACLUserRole table
const ACLUserRoleEscapedColumnChecksum = "`checksum`"

// ACLUserRoleColumnUserID is the UserID SQL column name for the ACLUserRole table
const ACLUserRoleColumnUserID = "user_id"

// ACLUserRoleEscapedColumnUserID is the escaped UserID SQL column name for the ACLUserRole table
const ACLUserRoleEscapedColumnUserID = "`user_id`"

// ACLUserRoleColumnRoleID is the RoleID SQL column name for the ACLUserRole table
const ACLUserRoleColumnRoleID = "role_id"

// ACLUserRoleEscapedColumnRoleID is the escaped RoleID SQL column name for the ACLUserRole table
const ACLUserRoleEscapedColumnRoleID = "`role_id`"

// ACLUserRoleColumnCustomerID is the CustomerID SQL column name for the ACLUserRole table
const ACLUserRoleColumnCustomerID = "customer_id"

// ACLUserRoleEscapedColumnCustomerID is the escaped CustomerID SQL column name for the ACLUserRole table
const ACLUserRoleEscapedColumnCustomerID = "`customer_id`"

// ACLUserRoleColumnAdminUserID is the AdminUserID SQL column name for the ACLUserRole table
const ACLUserRoleColumnAdminUserID = "admin_user_id"

// ACLUserRoleEscapedColumnAdminUserID is the escaped AdminUserID SQL column name for the ACLUserRole table
const ACLUserRoleEscapedColumnAdminUserID = "`admin_user_id`"

// ACLUserRoleColumnCreatedAt is the CreatedAt SQL column name for the ACLUserRole table
const ACLUserRoleColumnCreatedAt = "created_at"

// ACLUserRoleEscapedColumnCreatedAt is the escaped CreatedAt SQL column name for the ACLUserRole table
const ACLUserRoleEscapedColumnCreatedAt = "`created_at`"

// ACLUserRoleColumnUpdatedAt is the UpdatedAt SQL column name for the ACLUserRole table
const ACLUserRoleColumnUpdatedAt = "updated_at"

// ACLUserRoleEscapedColumnUpdatedAt is the escaped UpdatedAt SQL column name for the ACLUserRole table
const ACLUserRoleEscapedColumnUpdatedAt = "`updated_at`"

// GetID will return the ACLUserRole ID value
func (t *ACLUserRole) GetID() string {
	return t.ID
}

// SetID will set the ACLUserRole ID value
func (t *ACLUserRole) SetID(v string) {
	t.ID = v
}

// FindACLUserRoleByID will find a ACLUserRole by ID
func FindACLUserRoleByID(ctx context.Context, db *sql.DB, value string) (*ACLUserRole, error) {
	q := "SELECT `acl_user_role`.`id`,`acl_user_role`.`checksum`,`acl_user_role`.`user_id`,`acl_user_role`.`role_id`,`acl_user_role`.`customer_id`,`acl_user_role`.`admin_user_id`,`acl_user_role`.`created_at`,`acl_user_role`.`updated_at` FROM `acl_user_role` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _UserID sql.NullString
	var _RoleID sql.NullString
	var _CustomerID sql.NullString
	var _AdminUserID sql.NullString
	var _CreatedAt sql.NullInt64
	var _UpdatedAt sql.NullInt64
	err := db.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_UserID,
		&_RoleID,
		&_CustomerID,
		&_AdminUserID,
		&_CreatedAt,
		&_UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &ACLUserRole{}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _Checksum.Valid {
		t.SetChecksum(_Checksum.String)
	}
	if _UserID.Valid {
		t.SetUserID(_UserID.String)
	}
	if _RoleID.Valid {
		t.SetRoleID(_RoleID.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _AdminUserID.Valid {
		t.SetAdminUserID(_AdminUserID.String)
	}
	if _CreatedAt.Valid {
		t.SetCreatedAt(_CreatedAt.Int64)
	}
	if _UpdatedAt.Valid {
		t.SetUpdatedAt(_UpdatedAt.Int64)
	}
	return t, nil
}

// FindACLUserRoleByIDTx will find a ACLUserRole by ID using the provided transaction
func FindACLUserRoleByIDTx(ctx context.Context, tx *sql.Tx, value string) (*ACLUserRole, error) {
	q := "SELECT `acl_user_role`.`id`,`acl_user_role`.`checksum`,`acl_user_role`.`user_id`,`acl_user_role`.`role_id`,`acl_user_role`.`customer_id`,`acl_user_role`.`admin_user_id`,`acl_user_role`.`created_at`,`acl_user_role`.`updated_at` FROM `acl_user_role` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _UserID sql.NullString
	var _RoleID sql.NullString
	var _CustomerID sql.NullString
	var _AdminUserID sql.NullString
	var _CreatedAt sql.NullInt64
	var _UpdatedAt sql.NullInt64
	err := tx.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_UserID,
		&_RoleID,
		&_CustomerID,
		&_AdminUserID,
		&_CreatedAt,
		&_UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &ACLUserRole{}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _Checksum.Valid {
		t.SetChecksum(_Checksum.String)
	}
	if _UserID.Valid {
		t.SetUserID(_UserID.String)
	}
	if _RoleID.Valid {
		t.SetRoleID(_RoleID.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _AdminUserID.Valid {
		t.SetAdminUserID(_AdminUserID.String)
	}
	if _CreatedAt.Valid {
		t.SetCreatedAt(_CreatedAt.Int64)
	}
	if _UpdatedAt.Valid {
		t.SetUpdatedAt(_UpdatedAt.Int64)
	}
	return t, nil
}

// GetChecksum will return the ACLUserRole Checksum value
func (t *ACLUserRole) GetChecksum() string {
	if t.Checksum == nil {
		return ""
	}
	return *t.Checksum
}

// SetChecksum will set the ACLUserRole Checksum value
func (t *ACLUserRole) SetChecksum(v string) {
	t.Checksum = &v
}

// GetUserID will return the ACLUserRole UserID value
func (t *ACLUserRole) GetUserID() string {
	return t.UserID
}

// SetUserID will set the ACLUserRole UserID value
func (t *ACLUserRole) SetUserID(v string) {
	t.UserID = v
}

// FindACLUserRolesByUserID will find all ACLUserRoles by the UserID value
func FindACLUserRolesByUserID(ctx context.Context, db *sql.DB, value string) ([]*ACLUserRole, error) {
	q := "SELECT `acl_user_role`.`id`,`acl_user_role`.`checksum`,`acl_user_role`.`user_id`,`acl_user_role`.`role_id`,`acl_user_role`.`customer_id`,`acl_user_role`.`admin_user_id`,`acl_user_role`.`created_at`,`acl_user_role`.`updated_at` FROM `acl_user_role` WHERE `user_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*ACLUserRole, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _UserID sql.NullString
		var _RoleID sql.NullString
		var _CustomerID sql.NullString
		var _AdminUserID sql.NullString
		var _CreatedAt sql.NullInt64
		var _UpdatedAt sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_UserID,
			&_RoleID,
			&_CustomerID,
			&_AdminUserID,
			&_CreatedAt,
			&_UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		t := &ACLUserRole{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _RoleID.Valid {
			t.SetRoleID(_RoleID.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _AdminUserID.Valid {
			t.SetAdminUserID(_AdminUserID.String)
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

// FindACLUserRolesByUserIDTx will find all ACLUserRoles by the UserID value using the provided transaction
func FindACLUserRolesByUserIDTx(ctx context.Context, tx *sql.Tx, value string) ([]*ACLUserRole, error) {
	q := "SELECT `acl_user_role`.`id`,`acl_user_role`.`checksum`,`acl_user_role`.`user_id`,`acl_user_role`.`role_id`,`acl_user_role`.`customer_id`,`acl_user_role`.`admin_user_id`,`acl_user_role`.`created_at`,`acl_user_role`.`updated_at` FROM `acl_user_role` WHERE `user_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*ACLUserRole, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _UserID sql.NullString
		var _RoleID sql.NullString
		var _CustomerID sql.NullString
		var _AdminUserID sql.NullString
		var _CreatedAt sql.NullInt64
		var _UpdatedAt sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_UserID,
			&_RoleID,
			&_CustomerID,
			&_AdminUserID,
			&_CreatedAt,
			&_UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		t := &ACLUserRole{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _RoleID.Valid {
			t.SetRoleID(_RoleID.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _AdminUserID.Valid {
			t.SetAdminUserID(_AdminUserID.String)
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

// GetRoleID will return the ACLUserRole RoleID value
func (t *ACLUserRole) GetRoleID() string {
	return t.RoleID
}

// SetRoleID will set the ACLUserRole RoleID value
func (t *ACLUserRole) SetRoleID(v string) {
	t.RoleID = v
}

// FindACLUserRolesByRoleID will find all ACLUserRoles by the RoleID value
func FindACLUserRolesByRoleID(ctx context.Context, db *sql.DB, value string) ([]*ACLUserRole, error) {
	q := "SELECT `acl_user_role`.`id`,`acl_user_role`.`checksum`,`acl_user_role`.`user_id`,`acl_user_role`.`role_id`,`acl_user_role`.`customer_id`,`acl_user_role`.`admin_user_id`,`acl_user_role`.`created_at`,`acl_user_role`.`updated_at` FROM `acl_user_role` WHERE `role_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*ACLUserRole, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _UserID sql.NullString
		var _RoleID sql.NullString
		var _CustomerID sql.NullString
		var _AdminUserID sql.NullString
		var _CreatedAt sql.NullInt64
		var _UpdatedAt sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_UserID,
			&_RoleID,
			&_CustomerID,
			&_AdminUserID,
			&_CreatedAt,
			&_UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		t := &ACLUserRole{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _RoleID.Valid {
			t.SetRoleID(_RoleID.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _AdminUserID.Valid {
			t.SetAdminUserID(_AdminUserID.String)
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

// FindACLUserRolesByRoleIDTx will find all ACLUserRoles by the RoleID value using the provided transaction
func FindACLUserRolesByRoleIDTx(ctx context.Context, tx *sql.Tx, value string) ([]*ACLUserRole, error) {
	q := "SELECT `acl_user_role`.`id`,`acl_user_role`.`checksum`,`acl_user_role`.`user_id`,`acl_user_role`.`role_id`,`acl_user_role`.`customer_id`,`acl_user_role`.`admin_user_id`,`acl_user_role`.`created_at`,`acl_user_role`.`updated_at` FROM `acl_user_role` WHERE `role_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*ACLUserRole, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _UserID sql.NullString
		var _RoleID sql.NullString
		var _CustomerID sql.NullString
		var _AdminUserID sql.NullString
		var _CreatedAt sql.NullInt64
		var _UpdatedAt sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_UserID,
			&_RoleID,
			&_CustomerID,
			&_AdminUserID,
			&_CreatedAt,
			&_UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		t := &ACLUserRole{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _RoleID.Valid {
			t.SetRoleID(_RoleID.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _AdminUserID.Valid {
			t.SetAdminUserID(_AdminUserID.String)
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

// GetCustomerID will return the ACLUserRole CustomerID value
func (t *ACLUserRole) GetCustomerID() string {
	return t.CustomerID
}

// SetCustomerID will set the ACLUserRole CustomerID value
func (t *ACLUserRole) SetCustomerID(v string) {
	t.CustomerID = v
}

// FindACLUserRolesByCustomerID will find all ACLUserRoles by the CustomerID value
func FindACLUserRolesByCustomerID(ctx context.Context, db *sql.DB, value string) ([]*ACLUserRole, error) {
	q := "SELECT `acl_user_role`.`id`,`acl_user_role`.`checksum`,`acl_user_role`.`user_id`,`acl_user_role`.`role_id`,`acl_user_role`.`customer_id`,`acl_user_role`.`admin_user_id`,`acl_user_role`.`created_at`,`acl_user_role`.`updated_at` FROM `acl_user_role` WHERE `customer_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*ACLUserRole, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _UserID sql.NullString
		var _RoleID sql.NullString
		var _CustomerID sql.NullString
		var _AdminUserID sql.NullString
		var _CreatedAt sql.NullInt64
		var _UpdatedAt sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_UserID,
			&_RoleID,
			&_CustomerID,
			&_AdminUserID,
			&_CreatedAt,
			&_UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		t := &ACLUserRole{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _RoleID.Valid {
			t.SetRoleID(_RoleID.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _AdminUserID.Valid {
			t.SetAdminUserID(_AdminUserID.String)
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

// FindACLUserRolesByCustomerIDTx will find all ACLUserRoles by the CustomerID value using the provided transaction
func FindACLUserRolesByCustomerIDTx(ctx context.Context, tx *sql.Tx, value string) ([]*ACLUserRole, error) {
	q := "SELECT `acl_user_role`.`id`,`acl_user_role`.`checksum`,`acl_user_role`.`user_id`,`acl_user_role`.`role_id`,`acl_user_role`.`customer_id`,`acl_user_role`.`admin_user_id`,`acl_user_role`.`created_at`,`acl_user_role`.`updated_at` FROM `acl_user_role` WHERE `customer_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*ACLUserRole, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _UserID sql.NullString
		var _RoleID sql.NullString
		var _CustomerID sql.NullString
		var _AdminUserID sql.NullString
		var _CreatedAt sql.NullInt64
		var _UpdatedAt sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_UserID,
			&_RoleID,
			&_CustomerID,
			&_AdminUserID,
			&_CreatedAt,
			&_UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		t := &ACLUserRole{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _RoleID.Valid {
			t.SetRoleID(_RoleID.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _AdminUserID.Valid {
			t.SetAdminUserID(_AdminUserID.String)
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

// GetAdminUserID will return the ACLUserRole AdminUserID value
func (t *ACLUserRole) GetAdminUserID() string {
	return t.AdminUserID
}

// SetAdminUserID will set the ACLUserRole AdminUserID value
func (t *ACLUserRole) SetAdminUserID(v string) {
	t.AdminUserID = v
}

// GetCreatedAt will return the ACLUserRole CreatedAt value
func (t *ACLUserRole) GetCreatedAt() int64 {
	return t.CreatedAt
}

// SetCreatedAt will set the ACLUserRole CreatedAt value
func (t *ACLUserRole) SetCreatedAt(v int64) {
	t.CreatedAt = v
}

// GetUpdatedAt will return the ACLUserRole UpdatedAt value
func (t *ACLUserRole) GetUpdatedAt() int64 {
	if t.UpdatedAt == nil {
		return int64(0)
	}
	return *t.UpdatedAt
}

// SetUpdatedAt will set the ACLUserRole UpdatedAt value
func (t *ACLUserRole) SetUpdatedAt(v int64) {
	t.UpdatedAt = &v
}

func (t *ACLUserRole) toTimestamp(value time.Time) *timestamp.Timestamp {
	ts, _ := ptypes.TimestampProto(value)
	return ts
}

// DBCreateACLUserRoleTable will create the ACLUserRole table
func DBCreateACLUserRoleTable(ctx context.Context, db *sql.DB) error {
	q := "CREATE TABLE `acl_user_role` (`id`VARCHAR(64) NOT NULL PRIMARY KEY,`checksum`CHAR(64),`user_id` VARCHAR(64) NOT NULL,`role_id` VARCHAR(64) NOT NULL,`customer_id`VARCHAR(64) NOT NULL,`admin_user_id` VARCHAR(64) NOT NULL,`created_at` BIGINT(20) UNSIGNED NOT NULL,`updated_at` BIGINT(20) UNSIGNED,INDEX acl_user_role_user_id_index (`user_id`),INDEX acl_user_role_role_id_index (`role_id`),INDEX acl_user_role_customer_id_index (`customer_id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBCreateACLUserRoleTableTx will create the ACLUserRole table using the provided transction
func DBCreateACLUserRoleTableTx(ctx context.Context, tx *sql.Tx) error {
	q := "CREATE TABLE `acl_user_role` (`id`VARCHAR(64) NOT NULL PRIMARY KEY,`checksum`CHAR(64),`user_id` VARCHAR(64) NOT NULL,`role_id` VARCHAR(64) NOT NULL,`customer_id`VARCHAR(64) NOT NULL,`admin_user_id` VARCHAR(64) NOT NULL,`created_at` BIGINT(20) UNSIGNED NOT NULL,`updated_at` BIGINT(20) UNSIGNED,INDEX acl_user_role_user_id_index (`user_id`),INDEX acl_user_role_role_id_index (`role_id`),INDEX acl_user_role_customer_id_index (`customer_id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// DBDropACLUserRoleTable will drop the ACLUserRole table
func DBDropACLUserRoleTable(ctx context.Context, db *sql.DB) error {
	q := "DROP TABLE IF EXISTS `acl_user_role`"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBDropACLUserRoleTableTx will drop the ACLUserRole table using the provided transaction
func DBDropACLUserRoleTableTx(ctx context.Context, tx *sql.Tx) error {
	q := "DROP TABLE IF EXISTS `acl_user_role`"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// CalculateChecksum will calculate a checksum of the SHA1 of all field values
func (t *ACLUserRole) CalculateChecksum() string {
	return orm.HashStrings(
		orm.ToString(t.ID),
		orm.ToString(t.UserID),
		orm.ToString(t.RoleID),
		orm.ToString(t.CustomerID),
		orm.ToString(t.AdminUserID),
		orm.ToString(t.CreatedAt),
		orm.ToString(t.UpdatedAt),
	)
}

// DBCreate will create a new ACLUserRole record in the database
func (t *ACLUserRole) DBCreate(ctx context.Context, db *sql.DB) (sql.Result, error) {
	q := "INSERT INTO `acl_user_role` (`acl_user_role`.`id`,`acl_user_role`.`checksum`,`acl_user_role`.`user_id`,`acl_user_role`.`role_id`,`acl_user_role`.`customer_id`,`acl_user_role`.`admin_user_id`,`acl_user_role`.`created_at`,`acl_user_role`.`updated_at`) VALUES (?,?,?,?,?,?,?,?)"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.UserID),
		orm.ToSQLString(t.RoleID),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.AdminUserID),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
	)
}

// DBCreateTx will create a new ACLUserRole record in the database using the provided transaction
func (t *ACLUserRole) DBCreateTx(ctx context.Context, tx *sql.Tx) (sql.Result, error) {
	q := "INSERT INTO `acl_user_role` (`acl_user_role`.`id`,`acl_user_role`.`checksum`,`acl_user_role`.`user_id`,`acl_user_role`.`role_id`,`acl_user_role`.`customer_id`,`acl_user_role`.`admin_user_id`,`acl_user_role`.`created_at`,`acl_user_role`.`updated_at`) VALUES (?,?,?,?,?,?,?,?)"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.UserID),
		orm.ToSQLString(t.RoleID),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.AdminUserID),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
	)
}

// DBCreateIgnoreDuplicate will upsert the ACLUserRole record in the database
func (t *ACLUserRole) DBCreateIgnoreDuplicate(ctx context.Context, db *sql.DB) (sql.Result, error) {
	q := "INSERT INTO `acl_user_role` (`acl_user_role`.`id`,`acl_user_role`.`checksum`,`acl_user_role`.`user_id`,`acl_user_role`.`role_id`,`acl_user_role`.`customer_id`,`acl_user_role`.`admin_user_id`,`acl_user_role`.`created_at`,`acl_user_role`.`updated_at`) VALUES (?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.UserID),
		orm.ToSQLString(t.RoleID),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.AdminUserID),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
	)
}

// DBCreateIgnoreDuplicateTx will upsert the ACLUserRole record in the database using the provided transaction
func (t *ACLUserRole) DBCreateIgnoreDuplicateTx(ctx context.Context, tx *sql.Tx) (sql.Result, error) {
	q := "INSERT INTO `acl_user_role` (`acl_user_role`.`id`,`acl_user_role`.`checksum`,`acl_user_role`.`user_id`,`acl_user_role`.`role_id`,`acl_user_role`.`customer_id`,`acl_user_role`.`admin_user_id`,`acl_user_role`.`created_at`,`acl_user_role`.`updated_at`) VALUES (?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.UserID),
		orm.ToSQLString(t.RoleID),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.AdminUserID),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
	)
}

// DeleteAllACLUserRoles deletes all ACLUserRole records in the database with optional filters
func DeleteAllACLUserRoles(ctx context.Context, db *sql.DB, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(ACLUserRoleTableName),
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

// DeleteAllACLUserRolesTx deletes all ACLUserRole records in the database with optional filters using the provided transaction
func DeleteAllACLUserRolesTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(ACLUserRoleTableName),
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

// DBDelete will delete this ACLUserRole record in the database
func (t *ACLUserRole) DBDelete(ctx context.Context, db *sql.DB) (bool, error) {
	q := "DELETE FROM `acl_user_role` WHERE `id` = ?"
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

// DBDeleteTx will delete this ACLUserRole record in the database using the provided transaction
func (t *ACLUserRole) DBDeleteTx(ctx context.Context, tx *sql.Tx) (bool, error) {
	q := "DELETE FROM `acl_user_role` WHERE `id` = ?"
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

// DBUpdate will update the ACLUserRole record in the database
func (t *ACLUserRole) DBUpdate(ctx context.Context, db *sql.DB) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `acl_user_role` SET `checksum`=?,`user_id`=?,`role_id`=?,`customer_id`=?,`admin_user_id`=?,`created_at`=?,`updated_at`=? WHERE `id`=?"
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.UserID),
		orm.ToSQLString(t.RoleID),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.AdminUserID),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
		orm.ToSQLString(t.ID),
	)
}

// DBUpdateTx will update the ACLUserRole record in the database using the provided transaction
func (t *ACLUserRole) DBUpdateTx(ctx context.Context, tx *sql.Tx) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `acl_user_role` SET `checksum`=?,`user_id`=?,`role_id`=?,`customer_id`=?,`admin_user_id`=?,`created_at`=?,`updated_at`=? WHERE `id`=?"
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.UserID),
		orm.ToSQLString(t.RoleID),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.AdminUserID),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
		orm.ToSQLString(t.ID),
	)
}

// DBUpsert will upsert the ACLUserRole record in the database
func (t *ACLUserRole) DBUpsert(ctx context.Context, db *sql.DB, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `acl_user_role` (`acl_user_role`.`id`,`acl_user_role`.`checksum`,`acl_user_role`.`user_id`,`acl_user_role`.`role_id`,`acl_user_role`.`customer_id`,`acl_user_role`.`admin_user_id`,`acl_user_role`.`created_at`,`acl_user_role`.`updated_at`) VALUES (?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `acl_user_role` (`acl_user_role`.`id`,`acl_user_role`.`checksum`,`acl_user_role`.`user_id`,`acl_user_role`.`role_id`,`acl_user_role`.`customer_id`,`acl_user_role`.`admin_user_id`,`acl_user_role`.`created_at`,`acl_user_role`.`updated_at`) VALUES (?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`user_id`=VALUES(`user_id`),`role_id`=VALUES(`role_id`),`customer_id`=VALUES(`customer_id`),`admin_user_id`=VALUES(`admin_user_id`),`created_at`=VALUES(`created_at`),`updated_at`=VALUES(`updated_at`)"
	}
	r, err := db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.UserID),
		orm.ToSQLString(t.RoleID),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.AdminUserID),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBUpsertTx will upsert the ACLUserRole record in the database using the provided transaction
func (t *ACLUserRole) DBUpsertTx(ctx context.Context, tx *sql.Tx, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `acl_user_role` (`acl_user_role`.`id`,`acl_user_role`.`checksum`,`acl_user_role`.`user_id`,`acl_user_role`.`role_id`,`acl_user_role`.`customer_id`,`acl_user_role`.`admin_user_id`,`acl_user_role`.`created_at`,`acl_user_role`.`updated_at`) VALUES (?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `acl_user_role` (`acl_user_role`.`id`,`acl_user_role`.`checksum`,`acl_user_role`.`user_id`,`acl_user_role`.`role_id`,`acl_user_role`.`customer_id`,`acl_user_role`.`admin_user_id`,`acl_user_role`.`created_at`,`acl_user_role`.`updated_at`) VALUES (?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`user_id`=VALUES(`user_id`),`role_id`=VALUES(`role_id`),`customer_id`=VALUES(`customer_id`),`admin_user_id`=VALUES(`admin_user_id`),`created_at`=VALUES(`created_at`),`updated_at`=VALUES(`updated_at`)"
	}
	r, err := tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.UserID),
		orm.ToSQLString(t.RoleID),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.AdminUserID),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBFindOne will find a ACLUserRole record in the database with the primary key
func (t *ACLUserRole) DBFindOne(ctx context.Context, db *sql.DB, value string) (bool, error) {
	q := "SELECT `acl_user_role`.`id`,`acl_user_role`.`checksum`,`acl_user_role`.`user_id`,`acl_user_role`.`role_id`,`acl_user_role`.`customer_id`,`acl_user_role`.`admin_user_id`,`acl_user_role`.`created_at`,`acl_user_role`.`updated_at` FROM `acl_user_role` WHERE `id` = ? LIMIT 1"
	row := db.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _UserID sql.NullString
	var _RoleID sql.NullString
	var _CustomerID sql.NullString
	var _AdminUserID sql.NullString
	var _CreatedAt sql.NullInt64
	var _UpdatedAt sql.NullInt64
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_UserID,
		&_RoleID,
		&_CustomerID,
		&_AdminUserID,
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
	if _UserID.Valid {
		t.SetUserID(_UserID.String)
	}
	if _RoleID.Valid {
		t.SetRoleID(_RoleID.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _AdminUserID.Valid {
		t.SetAdminUserID(_AdminUserID.String)
	}
	if _CreatedAt.Valid {
		t.SetCreatedAt(_CreatedAt.Int64)
	}
	if _UpdatedAt.Valid {
		t.SetUpdatedAt(_UpdatedAt.Int64)
	}
	return true, nil
}

// DBFindOneTx will find a ACLUserRole record in the database with the primary key using the provided transaction
func (t *ACLUserRole) DBFindOneTx(ctx context.Context, tx *sql.Tx, value string) (bool, error) {
	q := "SELECT `acl_user_role`.`id`,`acl_user_role`.`checksum`,`acl_user_role`.`user_id`,`acl_user_role`.`role_id`,`acl_user_role`.`customer_id`,`acl_user_role`.`admin_user_id`,`acl_user_role`.`created_at`,`acl_user_role`.`updated_at` FROM `acl_user_role` WHERE `id` = ? LIMIT 1"
	row := tx.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _UserID sql.NullString
	var _RoleID sql.NullString
	var _CustomerID sql.NullString
	var _AdminUserID sql.NullString
	var _CreatedAt sql.NullInt64
	var _UpdatedAt sql.NullInt64
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_UserID,
		&_RoleID,
		&_CustomerID,
		&_AdminUserID,
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
	if _UserID.Valid {
		t.SetUserID(_UserID.String)
	}
	if _RoleID.Valid {
		t.SetRoleID(_RoleID.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _AdminUserID.Valid {
		t.SetAdminUserID(_AdminUserID.String)
	}
	if _CreatedAt.Valid {
		t.SetCreatedAt(_CreatedAt.Int64)
	}
	if _UpdatedAt.Valid {
		t.SetUpdatedAt(_UpdatedAt.Int64)
	}
	return true, nil
}

// FindACLUserRoles will find a ACLUserRole record in the database with the provided parameters
func FindACLUserRoles(ctx context.Context, db *sql.DB, _params ...interface{}) ([]*ACLUserRole, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("user_id"),
		orm.Column("role_id"),
		orm.Column("customer_id"),
		orm.Column("admin_user_id"),
		orm.Column("created_at"),
		orm.Column("updated_at"),
		orm.Table(ACLUserRoleTableName),
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
	results := make([]*ACLUserRole, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _UserID sql.NullString
		var _RoleID sql.NullString
		var _CustomerID sql.NullString
		var _AdminUserID sql.NullString
		var _CreatedAt sql.NullInt64
		var _UpdatedAt sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_UserID,
			&_RoleID,
			&_CustomerID,
			&_AdminUserID,
			&_CreatedAt,
			&_UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		t := &ACLUserRole{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _RoleID.Valid {
			t.SetRoleID(_RoleID.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _AdminUserID.Valid {
			t.SetAdminUserID(_AdminUserID.String)
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

// FindACLUserRolesTx will find a ACLUserRole record in the database with the provided parameters using the provided transaction
func FindACLUserRolesTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) ([]*ACLUserRole, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("user_id"),
		orm.Column("role_id"),
		orm.Column("customer_id"),
		orm.Column("admin_user_id"),
		orm.Column("created_at"),
		orm.Column("updated_at"),
		orm.Table(ACLUserRoleTableName),
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
	results := make([]*ACLUserRole, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _UserID sql.NullString
		var _RoleID sql.NullString
		var _CustomerID sql.NullString
		var _AdminUserID sql.NullString
		var _CreatedAt sql.NullInt64
		var _UpdatedAt sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_UserID,
			&_RoleID,
			&_CustomerID,
			&_AdminUserID,
			&_CreatedAt,
			&_UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		t := &ACLUserRole{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _RoleID.Valid {
			t.SetRoleID(_RoleID.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _AdminUserID.Valid {
			t.SetAdminUserID(_AdminUserID.String)
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

// DBFind will find a ACLUserRole record in the database with the provided parameters
func (t *ACLUserRole) DBFind(ctx context.Context, db *sql.DB, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("user_id"),
		orm.Column("role_id"),
		orm.Column("customer_id"),
		orm.Column("admin_user_id"),
		orm.Column("created_at"),
		orm.Column("updated_at"),
		orm.Table(ACLUserRoleTableName),
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
	var _UserID sql.NullString
	var _RoleID sql.NullString
	var _CustomerID sql.NullString
	var _AdminUserID sql.NullString
	var _CreatedAt sql.NullInt64
	var _UpdatedAt sql.NullInt64
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_UserID,
		&_RoleID,
		&_CustomerID,
		&_AdminUserID,
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
	if _UserID.Valid {
		t.SetUserID(_UserID.String)
	}
	if _RoleID.Valid {
		t.SetRoleID(_RoleID.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _AdminUserID.Valid {
		t.SetAdminUserID(_AdminUserID.String)
	}
	if _CreatedAt.Valid {
		t.SetCreatedAt(_CreatedAt.Int64)
	}
	if _UpdatedAt.Valid {
		t.SetUpdatedAt(_UpdatedAt.Int64)
	}
	return true, nil
}

// DBFindTx will find a ACLUserRole record in the database with the provided parameters using the provided transaction
func (t *ACLUserRole) DBFindTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("user_id"),
		orm.Column("role_id"),
		orm.Column("customer_id"),
		orm.Column("admin_user_id"),
		orm.Column("created_at"),
		orm.Column("updated_at"),
		orm.Table(ACLUserRoleTableName),
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
	var _UserID sql.NullString
	var _RoleID sql.NullString
	var _CustomerID sql.NullString
	var _AdminUserID sql.NullString
	var _CreatedAt sql.NullInt64
	var _UpdatedAt sql.NullInt64
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_UserID,
		&_RoleID,
		&_CustomerID,
		&_AdminUserID,
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
	if _UserID.Valid {
		t.SetUserID(_UserID.String)
	}
	if _RoleID.Valid {
		t.SetRoleID(_RoleID.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _AdminUserID.Valid {
		t.SetAdminUserID(_AdminUserID.String)
	}
	if _CreatedAt.Valid {
		t.SetCreatedAt(_CreatedAt.Int64)
	}
	if _UpdatedAt.Valid {
		t.SetUpdatedAt(_UpdatedAt.Int64)
	}
	return true, nil
}

// CountACLUserRoles will find the count of ACLUserRole records in the database
func CountACLUserRoles(ctx context.Context, db *sql.DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(ACLUserRoleTableName),
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

// CountACLUserRolesTx will find the count of ACLUserRole records in the database using the provided transaction
func CountACLUserRolesTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(ACLUserRoleTableName),
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

// DBCount will find the count of ACLUserRole records in the database
func (t *ACLUserRole) DBCount(ctx context.Context, db *sql.DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(ACLUserRoleTableName),
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

// DBCountTx will find the count of ACLUserRole records in the database using the provided transaction
func (t *ACLUserRole) DBCountTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(ACLUserRoleTableName),
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

// DBExists will return true if the ACLUserRole record exists in the database
func (t *ACLUserRole) DBExists(ctx context.Context, db *sql.DB) (bool, error) {
	q := "SELECT `id` FROM `acl_user_role` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := db.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// DBExistsTx will return true if the ACLUserRole record exists in the database using the provided transaction
func (t *ACLUserRole) DBExistsTx(ctx context.Context, tx *sql.Tx) (bool, error) {
	q := "SELECT `id` FROM `acl_user_role` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := tx.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// PrimaryKeyColumn returns the column name for the primary key
func (t *ACLUserRole) PrimaryKeyColumn() string {
	return ACLUserRoleColumnID
}

// PrimaryKeyColumnType returns the primary key column Go type as a string
func (t *ACLUserRole) PrimaryKeyColumnType() string {
	return "string"
}

// PrimaryKey returns the primary key column value
func (t *ACLUserRole) PrimaryKey() interface{} {
	return t.ID
}
