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

var _ Model = (*ACLGrant)(nil)
var _ CSVWriter = (*ACLGrant)(nil)
var _ JSONWriter = (*ACLGrant)(nil)
var _ Checksum = (*ACLGrant)(nil)

// ACLGrantTableName is the name of the table in SQL
const ACLGrantTableName = "acl_grant"

var ACLGrantColumns = []string{
	"id",
	"checksum",
	"resource_id",
	"role_id",
	"permission",
	"admin_user_id",
	"customer_id",
	"created_at",
	"updated_at",
}

type ACLGrant_ACLGrantPermission string

const (
	ACLGrantPermission_READ      ACLGrant_ACLGrantPermission = "read"
	ACLGrantPermission_READWRITE ACLGrant_ACLGrantPermission = "readwrite"
	ACLGrantPermission_ADMIN     ACLGrant_ACLGrantPermission = "admin"
)

func (x ACLGrant_ACLGrantPermission) String() string {
	return string(x)
}

func enumACLGrantPermissionToString(v *ACLGrant_ACLGrantPermission) string {
	if v == nil {
		return ""
	}
	return v.String()
}

func toACLGrantPermission(v string) *ACLGrant_ACLGrantPermission {
	var ev *ACLGrant_ACLGrantPermission
	switch v {
	case "READ", "read":
		{
			v := ACLGrantPermission_READ
			ev = &v
		}
	case "READWRITE", "readwrite":
		{
			v := ACLGrantPermission_READWRITE
			ev = &v
		}
	case "ADMIN", "admin":
		{
			v := ACLGrantPermission_ADMIN
			ev = &v
		}
	}
	return ev
}

// ACLGrant table
type ACLGrant struct {
	AdminUserID string                      `json:"admin_user_id"`
	Checksum    *string                     `json:"checksum,omitempty"`
	CreatedAt   int64                       `json:"created_at"`
	CustomerID  string                      `json:"customer_id"`
	ID          string                      `json:"id"`
	Permission  ACLGrant_ACLGrantPermission `json:"permission"`
	ResourceID  string                      `json:"resource_id"`
	RoleID      string                      `json:"role_id"`
	UpdatedAt   *int64                      `json:"updated_at,omitempty"`
}

// TableName returns the SQL table name for ACLGrant and satifies the Model interface
func (t *ACLGrant) TableName() string {
	return ACLGrantTableName
}

// ToCSV will serialize the ACLGrant instance to a CSV compatible array of strings
func (t *ACLGrant) ToCSV() []string {
	return []string{
		t.ID,
		t.CalculateChecksum(),
		t.ResourceID,
		t.RoleID,
		toCSVString(t.Permission),
		t.AdminUserID,
		t.CustomerID,
		toCSVString(t.CreatedAt),
		toCSVString(t.UpdatedAt),
	}
}

// WriteCSV will serialize the ACLGrant instance to the writer as CSV and satisfies the CSVWriter interface
func (t *ACLGrant) WriteCSV(w *csv.Writer) error {
	return w.Write(t.ToCSV())
}

// WriteJSON will serialize the ACLGrant instance to the writer as JSON and satisfies the JSONWriter interface
func (t *ACLGrant) WriteJSON(w io.Writer, indent ...bool) error {
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

// NewACLGrantReader creates a JSON reader which can read in ACLGrant objects serialized as JSON either as an array, single object or json new lines
// and writes each ACLGrant to the channel provided
func NewACLGrantReader(r io.Reader, ch chan<- ACLGrant) error {
	return orm.Deserialize(r, func(buf json.RawMessage) error {
		dec := json.NewDecoder(bytes.NewBuffer(buf))
		e := ACLGrant{}
		if err := dec.Decode(&e); err != nil {
			return err
		}
		ch <- e
		return nil
	})
}

// NewCSVACLGrantReaderDir will read the reader as CSV and emit each record to the channel provided
func NewCSVACLGrantReader(r io.Reader, ch chan<- ACLGrant) error {
	cr := csv.NewReader(r)
	for {
		record, err := cr.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		ch <- ACLGrant{
			ID:          record[0],
			Checksum:    fromStringPointer(record[1]),
			ResourceID:  record[2],
			RoleID:      record[3],
			Permission:  *toACLGrantPermission(record[4]),
			AdminUserID: record[5],
			CustomerID:  record[6],
			CreatedAt:   fromCSVInt64(record[7]),
			UpdatedAt:   fromCSVInt64Pointer(record[8]),
		}
	}
	return nil
}

// NewCSVACLGrantReaderFile will read the file as a CSV and emit each record to the channel provided
func NewCSVACLGrantReaderFile(fp string, ch chan<- ACLGrant) error {
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
	return NewCSVACLGrantReader(fc, ch)
}

// NewCSVACLGrantReaderDir will read the acl_grant.csv.gz file as a CSV and emit each record to the channel provided
func NewCSVACLGrantReaderDir(dir string, ch chan<- ACLGrant) error {
	return NewCSVACLGrantReaderFile(filepath.Join(dir, "acl_grant.csv.gz"), ch)
}

// ACLGrantCSVDeduper is a function callback which takes the existing value (a) and the new value (b)
// and the return value should be the one to use (or a new one, if applicable). return nil
// to skip processing of this record
type ACLGrantCSVDeduper func(a ACLGrant, b ACLGrant) *ACLGrant

// ACLGrantCSVDedupeDisabled is set on whether the CSV writer should de-dupe values by key
var ACLGrantCSVDedupeDisabled bool

// NewACLGrantCSVWriterSize creates a batch writer that will write each ACLGrant into a CSV file
// this method will automatically de-duplicate entries using the primary key. if the checksum
// for a newer item with the same primary key doesn't match a previously sent item, the newer
// one will be used
func NewACLGrantCSVWriterSize(w io.Writer, size int, dedupers ...ACLGrantCSVDeduper) (chan ACLGrant, chan bool, error) {
	cw := csv.NewWriter(w)
	ch := make(chan ACLGrant, size)
	done := make(chan bool)
	go func() {
		defer func() { done <- true }()
		dodedupe := !ACLGrantCSVDedupeDisabled
		var kv map[string]*ACLGrant
		var deduper ACLGrantCSVDeduper
		if dedupers != nil && len(dedupers) > 0 {
			deduper = dedupers[0]
			dodedupe = true
		}
		if dodedupe {
			kv = make(map[string]*ACLGrant)
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

// ACLGrantCSVDefaultSize is the default channel buffer size if not provided
var ACLGrantCSVDefaultSize = 100

// NewACLGrantCSVWriter creates a batch writer that will write each ACLGrant into a CSV file
func NewACLGrantCSVWriter(w io.Writer, dedupers ...ACLGrantCSVDeduper) (chan ACLGrant, chan bool, error) {
	return NewACLGrantCSVWriterSize(w, ACLGrantCSVDefaultSize, dedupers...)
}

// NewACLGrantCSVWriterDir creates a batch writer that will write each ACLGrant into a CSV file named acl_grant.csv.gz in dir
func NewACLGrantCSVWriterDir(dir string, dedupers ...ACLGrantCSVDeduper) (chan ACLGrant, chan bool, error) {
	return NewACLGrantCSVWriterFile(filepath.Join(dir, "acl_grant.csv.gz"), dedupers...)
}

// NewACLGrantCSVWriterFile creates a batch writer that will write each ACLGrant into a CSV file
func NewACLGrantCSVWriterFile(fn string, dedupers ...ACLGrantCSVDeduper) (chan ACLGrant, chan bool, error) {
	f, err := os.Create(fn)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening CSV file %s. %v", fn, err)
	}
	var fc io.WriteCloser = f
	if filepath.Ext(fn) == ".gz" {
		w, _ := gzip.NewWriterLevel(f, gzip.BestCompression)
		fc = w
	}
	ch, done, err := NewACLGrantCSVWriter(fc, dedupers...)
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

type ACLGrantDBAction func(ctx context.Context, db *sql.DB, record ACLGrant) error

// NewACLGrantDBWriterSize creates a DB writer that will write each issue into the DB
func NewACLGrantDBWriterSize(ctx context.Context, db *sql.DB, errors chan<- error, size int, actions ...ACLGrantDBAction) (chan ACLGrant, chan bool, error) {
	ch := make(chan ACLGrant, size)
	done := make(chan bool)
	var action ACLGrantDBAction
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

// NewACLGrantDBWriter creates a DB writer that will write each issue into the DB
func NewACLGrantDBWriter(ctx context.Context, db *sql.DB, errors chan<- error, actions ...ACLGrantDBAction) (chan ACLGrant, chan bool, error) {
	return NewACLGrantDBWriterSize(ctx, db, errors, 100, actions...)
}

// ACLGrantColumnID is the ID SQL column name for the ACLGrant table
const ACLGrantColumnID = "id"

// ACLGrantEscapedColumnID is the escaped ID SQL column name for the ACLGrant table
const ACLGrantEscapedColumnID = "`id`"

// ACLGrantColumnChecksum is the Checksum SQL column name for the ACLGrant table
const ACLGrantColumnChecksum = "checksum"

// ACLGrantEscapedColumnChecksum is the escaped Checksum SQL column name for the ACLGrant table
const ACLGrantEscapedColumnChecksum = "`checksum`"

// ACLGrantColumnResourceID is the ResourceID SQL column name for the ACLGrant table
const ACLGrantColumnResourceID = "resource_id"

// ACLGrantEscapedColumnResourceID is the escaped ResourceID SQL column name for the ACLGrant table
const ACLGrantEscapedColumnResourceID = "`resource_id`"

// ACLGrantColumnRoleID is the RoleID SQL column name for the ACLGrant table
const ACLGrantColumnRoleID = "role_id"

// ACLGrantEscapedColumnRoleID is the escaped RoleID SQL column name for the ACLGrant table
const ACLGrantEscapedColumnRoleID = "`role_id`"

// ACLGrantColumnPermission is the Permission SQL column name for the ACLGrant table
const ACLGrantColumnPermission = "permission"

// ACLGrantEscapedColumnPermission is the escaped Permission SQL column name for the ACLGrant table
const ACLGrantEscapedColumnPermission = "`permission`"

// ACLGrantColumnAdminUserID is the AdminUserID SQL column name for the ACLGrant table
const ACLGrantColumnAdminUserID = "admin_user_id"

// ACLGrantEscapedColumnAdminUserID is the escaped AdminUserID SQL column name for the ACLGrant table
const ACLGrantEscapedColumnAdminUserID = "`admin_user_id`"

// ACLGrantColumnCustomerID is the CustomerID SQL column name for the ACLGrant table
const ACLGrantColumnCustomerID = "customer_id"

// ACLGrantEscapedColumnCustomerID is the escaped CustomerID SQL column name for the ACLGrant table
const ACLGrantEscapedColumnCustomerID = "`customer_id`"

// ACLGrantColumnCreatedAt is the CreatedAt SQL column name for the ACLGrant table
const ACLGrantColumnCreatedAt = "created_at"

// ACLGrantEscapedColumnCreatedAt is the escaped CreatedAt SQL column name for the ACLGrant table
const ACLGrantEscapedColumnCreatedAt = "`created_at`"

// ACLGrantColumnUpdatedAt is the UpdatedAt SQL column name for the ACLGrant table
const ACLGrantColumnUpdatedAt = "updated_at"

// ACLGrantEscapedColumnUpdatedAt is the escaped UpdatedAt SQL column name for the ACLGrant table
const ACLGrantEscapedColumnUpdatedAt = "`updated_at`"

// GetID will return the ACLGrant ID value
func (t *ACLGrant) GetID() string {
	return t.ID
}

// SetID will set the ACLGrant ID value
func (t *ACLGrant) SetID(v string) {
	t.ID = v
}

// FindACLGrantByID will find a ACLGrant by ID
func FindACLGrantByID(ctx context.Context, db *sql.DB, value string) (*ACLGrant, error) {
	q := "SELECT `acl_grant`.`id`,`acl_grant`.`checksum`,`acl_grant`.`resource_id`,`acl_grant`.`role_id`,`acl_grant`.`permission`,`acl_grant`.`admin_user_id`,`acl_grant`.`customer_id`,`acl_grant`.`created_at`,`acl_grant`.`updated_at` FROM `acl_grant` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _ResourceID sql.NullString
	var _RoleID sql.NullString
	var _Permission sql.NullString
	var _AdminUserID sql.NullString
	var _CustomerID sql.NullString
	var _CreatedAt sql.NullInt64
	var _UpdatedAt sql.NullInt64
	err := db.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_ResourceID,
		&_RoleID,
		&_Permission,
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
	t := &ACLGrant{}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _Checksum.Valid {
		t.SetChecksum(_Checksum.String)
	}
	if _ResourceID.Valid {
		t.SetResourceID(_ResourceID.String)
	}
	if _RoleID.Valid {
		t.SetRoleID(_RoleID.String)
	}
	if _Permission.Valid {
		t.SetPermissionString(_Permission.String)
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

// FindACLGrantByIDTx will find a ACLGrant by ID using the provided transaction
func FindACLGrantByIDTx(ctx context.Context, tx *sql.Tx, value string) (*ACLGrant, error) {
	q := "SELECT `acl_grant`.`id`,`acl_grant`.`checksum`,`acl_grant`.`resource_id`,`acl_grant`.`role_id`,`acl_grant`.`permission`,`acl_grant`.`admin_user_id`,`acl_grant`.`customer_id`,`acl_grant`.`created_at`,`acl_grant`.`updated_at` FROM `acl_grant` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _ResourceID sql.NullString
	var _RoleID sql.NullString
	var _Permission sql.NullString
	var _AdminUserID sql.NullString
	var _CustomerID sql.NullString
	var _CreatedAt sql.NullInt64
	var _UpdatedAt sql.NullInt64
	err := tx.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_ResourceID,
		&_RoleID,
		&_Permission,
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
	t := &ACLGrant{}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _Checksum.Valid {
		t.SetChecksum(_Checksum.String)
	}
	if _ResourceID.Valid {
		t.SetResourceID(_ResourceID.String)
	}
	if _RoleID.Valid {
		t.SetRoleID(_RoleID.String)
	}
	if _Permission.Valid {
		t.SetPermissionString(_Permission.String)
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

// GetChecksum will return the ACLGrant Checksum value
func (t *ACLGrant) GetChecksum() string {
	if t.Checksum == nil {
		return ""
	}
	return *t.Checksum
}

// SetChecksum will set the ACLGrant Checksum value
func (t *ACLGrant) SetChecksum(v string) {
	t.Checksum = &v
}

// GetResourceID will return the ACLGrant ResourceID value
func (t *ACLGrant) GetResourceID() string {
	return t.ResourceID
}

// SetResourceID will set the ACLGrant ResourceID value
func (t *ACLGrant) SetResourceID(v string) {
	t.ResourceID = v
}

// FindACLGrantsByResourceID will find all ACLGrants by the ResourceID value
func FindACLGrantsByResourceID(ctx context.Context, db *sql.DB, value string) ([]*ACLGrant, error) {
	q := "SELECT `acl_grant`.`id`,`acl_grant`.`checksum`,`acl_grant`.`resource_id`,`acl_grant`.`role_id`,`acl_grant`.`permission`,`acl_grant`.`admin_user_id`,`acl_grant`.`customer_id`,`acl_grant`.`created_at`,`acl_grant`.`updated_at` FROM `acl_grant` WHERE `resource_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*ACLGrant, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _ResourceID sql.NullString
		var _RoleID sql.NullString
		var _Permission sql.NullString
		var _AdminUserID sql.NullString
		var _CustomerID sql.NullString
		var _CreatedAt sql.NullInt64
		var _UpdatedAt sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_ResourceID,
			&_RoleID,
			&_Permission,
			&_AdminUserID,
			&_CustomerID,
			&_CreatedAt,
			&_UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		t := &ACLGrant{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _ResourceID.Valid {
			t.SetResourceID(_ResourceID.String)
		}
		if _RoleID.Valid {
			t.SetRoleID(_RoleID.String)
		}
		if _Permission.Valid {
			t.SetPermissionString(_Permission.String)
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

// FindACLGrantsByResourceIDTx will find all ACLGrants by the ResourceID value using the provided transaction
func FindACLGrantsByResourceIDTx(ctx context.Context, tx *sql.Tx, value string) ([]*ACLGrant, error) {
	q := "SELECT `acl_grant`.`id`,`acl_grant`.`checksum`,`acl_grant`.`resource_id`,`acl_grant`.`role_id`,`acl_grant`.`permission`,`acl_grant`.`admin_user_id`,`acl_grant`.`customer_id`,`acl_grant`.`created_at`,`acl_grant`.`updated_at` FROM `acl_grant` WHERE `resource_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*ACLGrant, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _ResourceID sql.NullString
		var _RoleID sql.NullString
		var _Permission sql.NullString
		var _AdminUserID sql.NullString
		var _CustomerID sql.NullString
		var _CreatedAt sql.NullInt64
		var _UpdatedAt sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_ResourceID,
			&_RoleID,
			&_Permission,
			&_AdminUserID,
			&_CustomerID,
			&_CreatedAt,
			&_UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		t := &ACLGrant{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _ResourceID.Valid {
			t.SetResourceID(_ResourceID.String)
		}
		if _RoleID.Valid {
			t.SetRoleID(_RoleID.String)
		}
		if _Permission.Valid {
			t.SetPermissionString(_Permission.String)
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

// GetRoleID will return the ACLGrant RoleID value
func (t *ACLGrant) GetRoleID() string {
	return t.RoleID
}

// SetRoleID will set the ACLGrant RoleID value
func (t *ACLGrant) SetRoleID(v string) {
	t.RoleID = v
}

// FindACLGrantsByRoleID will find all ACLGrants by the RoleID value
func FindACLGrantsByRoleID(ctx context.Context, db *sql.DB, value string) ([]*ACLGrant, error) {
	q := "SELECT `acl_grant`.`id`,`acl_grant`.`checksum`,`acl_grant`.`resource_id`,`acl_grant`.`role_id`,`acl_grant`.`permission`,`acl_grant`.`admin_user_id`,`acl_grant`.`customer_id`,`acl_grant`.`created_at`,`acl_grant`.`updated_at` FROM `acl_grant` WHERE `role_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*ACLGrant, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _ResourceID sql.NullString
		var _RoleID sql.NullString
		var _Permission sql.NullString
		var _AdminUserID sql.NullString
		var _CustomerID sql.NullString
		var _CreatedAt sql.NullInt64
		var _UpdatedAt sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_ResourceID,
			&_RoleID,
			&_Permission,
			&_AdminUserID,
			&_CustomerID,
			&_CreatedAt,
			&_UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		t := &ACLGrant{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _ResourceID.Valid {
			t.SetResourceID(_ResourceID.String)
		}
		if _RoleID.Valid {
			t.SetRoleID(_RoleID.String)
		}
		if _Permission.Valid {
			t.SetPermissionString(_Permission.String)
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

// FindACLGrantsByRoleIDTx will find all ACLGrants by the RoleID value using the provided transaction
func FindACLGrantsByRoleIDTx(ctx context.Context, tx *sql.Tx, value string) ([]*ACLGrant, error) {
	q := "SELECT `acl_grant`.`id`,`acl_grant`.`checksum`,`acl_grant`.`resource_id`,`acl_grant`.`role_id`,`acl_grant`.`permission`,`acl_grant`.`admin_user_id`,`acl_grant`.`customer_id`,`acl_grant`.`created_at`,`acl_grant`.`updated_at` FROM `acl_grant` WHERE `role_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*ACLGrant, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _ResourceID sql.NullString
		var _RoleID sql.NullString
		var _Permission sql.NullString
		var _AdminUserID sql.NullString
		var _CustomerID sql.NullString
		var _CreatedAt sql.NullInt64
		var _UpdatedAt sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_ResourceID,
			&_RoleID,
			&_Permission,
			&_AdminUserID,
			&_CustomerID,
			&_CreatedAt,
			&_UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		t := &ACLGrant{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _ResourceID.Valid {
			t.SetResourceID(_ResourceID.String)
		}
		if _RoleID.Valid {
			t.SetRoleID(_RoleID.String)
		}
		if _Permission.Valid {
			t.SetPermissionString(_Permission.String)
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

// GetPermission will return the ACLGrant Permission value
func (t *ACLGrant) GetPermission() ACLGrant_ACLGrantPermission {
	return t.Permission
}

// SetPermission will set the ACLGrant Permission value
func (t *ACLGrant) SetPermission(v ACLGrant_ACLGrantPermission) {
	t.Permission = v
}

// GetPermissionString will return the ACLGrant Permission value as a string
func (t *ACLGrant) GetPermissionString() string {
	return t.Permission.String()
}

// SetPermissionString will set the ACLGrant Permission value from a string
func (t *ACLGrant) SetPermissionString(v string) {
	var _Permission = toACLGrantPermission(v)
	if _Permission != nil {
		t.Permission = *_Permission
	}
}

// GetAdminUserID will return the ACLGrant AdminUserID value
func (t *ACLGrant) GetAdminUserID() string {
	return t.AdminUserID
}

// SetAdminUserID will set the ACLGrant AdminUserID value
func (t *ACLGrant) SetAdminUserID(v string) {
	t.AdminUserID = v
}

// GetCustomerID will return the ACLGrant CustomerID value
func (t *ACLGrant) GetCustomerID() string {
	return t.CustomerID
}

// SetCustomerID will set the ACLGrant CustomerID value
func (t *ACLGrant) SetCustomerID(v string) {
	t.CustomerID = v
}

// FindACLGrantsByCustomerID will find all ACLGrants by the CustomerID value
func FindACLGrantsByCustomerID(ctx context.Context, db *sql.DB, value string) ([]*ACLGrant, error) {
	q := "SELECT `acl_grant`.`id`,`acl_grant`.`checksum`,`acl_grant`.`resource_id`,`acl_grant`.`role_id`,`acl_grant`.`permission`,`acl_grant`.`admin_user_id`,`acl_grant`.`customer_id`,`acl_grant`.`created_at`,`acl_grant`.`updated_at` FROM `acl_grant` WHERE `customer_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*ACLGrant, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _ResourceID sql.NullString
		var _RoleID sql.NullString
		var _Permission sql.NullString
		var _AdminUserID sql.NullString
		var _CustomerID sql.NullString
		var _CreatedAt sql.NullInt64
		var _UpdatedAt sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_ResourceID,
			&_RoleID,
			&_Permission,
			&_AdminUserID,
			&_CustomerID,
			&_CreatedAt,
			&_UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		t := &ACLGrant{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _ResourceID.Valid {
			t.SetResourceID(_ResourceID.String)
		}
		if _RoleID.Valid {
			t.SetRoleID(_RoleID.String)
		}
		if _Permission.Valid {
			t.SetPermissionString(_Permission.String)
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

// FindACLGrantsByCustomerIDTx will find all ACLGrants by the CustomerID value using the provided transaction
func FindACLGrantsByCustomerIDTx(ctx context.Context, tx *sql.Tx, value string) ([]*ACLGrant, error) {
	q := "SELECT `acl_grant`.`id`,`acl_grant`.`checksum`,`acl_grant`.`resource_id`,`acl_grant`.`role_id`,`acl_grant`.`permission`,`acl_grant`.`admin_user_id`,`acl_grant`.`customer_id`,`acl_grant`.`created_at`,`acl_grant`.`updated_at` FROM `acl_grant` WHERE `customer_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*ACLGrant, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _ResourceID sql.NullString
		var _RoleID sql.NullString
		var _Permission sql.NullString
		var _AdminUserID sql.NullString
		var _CustomerID sql.NullString
		var _CreatedAt sql.NullInt64
		var _UpdatedAt sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_ResourceID,
			&_RoleID,
			&_Permission,
			&_AdminUserID,
			&_CustomerID,
			&_CreatedAt,
			&_UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		t := &ACLGrant{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _ResourceID.Valid {
			t.SetResourceID(_ResourceID.String)
		}
		if _RoleID.Valid {
			t.SetRoleID(_RoleID.String)
		}
		if _Permission.Valid {
			t.SetPermissionString(_Permission.String)
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

// GetCreatedAt will return the ACLGrant CreatedAt value
func (t *ACLGrant) GetCreatedAt() int64 {
	return t.CreatedAt
}

// SetCreatedAt will set the ACLGrant CreatedAt value
func (t *ACLGrant) SetCreatedAt(v int64) {
	t.CreatedAt = v
}

// GetUpdatedAt will return the ACLGrant UpdatedAt value
func (t *ACLGrant) GetUpdatedAt() int64 {
	if t.UpdatedAt == nil {
		return int64(0)
	}
	return *t.UpdatedAt
}

// SetUpdatedAt will set the ACLGrant UpdatedAt value
func (t *ACLGrant) SetUpdatedAt(v int64) {
	t.UpdatedAt = &v
}

func (t *ACLGrant) toTimestamp(value time.Time) *timestamp.Timestamp {
	ts, _ := ptypes.TimestampProto(value)
	return ts
}

// DBCreateACLGrantTable will create the ACLGrant table
func DBCreateACLGrantTable(ctx context.Context, db *sql.DB) error {
	q := "CREATE TABLE `acl_grant` (`id`VARCHAR(64) NOT NULL PRIMARY KEY,`checksum`CHAR(64),`resource_id`VARCHAR(64) NOT NULL,`role_id` VARCHAR(64) NOT NULL,`permission` ENUM('read','readwrite','admin') NOT NULL,`admin_user_id` VARCHAR(64) NOT NULL,`customer_id`VARCHAR(64) NOT NULL,`created_at` BIGINT UNSIGNED NOT NULL,`updated_at` BIGINT UNSIGNED,INDEX acl_grant_resource_id_index (`resource_id`),INDEX acl_grant_role_id_index (`role_id`),INDEX acl_grant_customer_id_index (`customer_id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBCreateACLGrantTableTx will create the ACLGrant table using the provided transction
func DBCreateACLGrantTableTx(ctx context.Context, tx *sql.Tx) error {
	q := "CREATE TABLE `acl_grant` (`id`VARCHAR(64) NOT NULL PRIMARY KEY,`checksum`CHAR(64),`resource_id`VARCHAR(64) NOT NULL,`role_id` VARCHAR(64) NOT NULL,`permission` ENUM('read','readwrite','admin') NOT NULL,`admin_user_id` VARCHAR(64) NOT NULL,`customer_id`VARCHAR(64) NOT NULL,`created_at` BIGINT UNSIGNED NOT NULL,`updated_at` BIGINT UNSIGNED,INDEX acl_grant_resource_id_index (`resource_id`),INDEX acl_grant_role_id_index (`role_id`),INDEX acl_grant_customer_id_index (`customer_id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// DBDropACLGrantTable will drop the ACLGrant table
func DBDropACLGrantTable(ctx context.Context, db *sql.DB) error {
	q := "DROP TABLE IF EXISTS `acl_grant`"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBDropACLGrantTableTx will drop the ACLGrant table using the provided transaction
func DBDropACLGrantTableTx(ctx context.Context, tx *sql.Tx) error {
	q := "DROP TABLE IF EXISTS `acl_grant`"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// CalculateChecksum will calculate a checksum of the SHA1 of all field values
func (t *ACLGrant) CalculateChecksum() string {
	return orm.HashStrings(
		orm.ToString(t.ID),
		orm.ToString(t.ResourceID),
		orm.ToString(t.RoleID),
		enumACLGrantPermissionToString(&t.Permission),
		orm.ToString(t.AdminUserID),
		orm.ToString(t.CustomerID),
		orm.ToString(t.CreatedAt),
		orm.ToString(t.UpdatedAt),
	)
}

// DBCreate will create a new ACLGrant record in the database
func (t *ACLGrant) DBCreate(ctx context.Context, db *sql.DB) (sql.Result, error) {
	q := "INSERT INTO `acl_grant` (`acl_grant`.`id`,`acl_grant`.`checksum`,`acl_grant`.`resource_id`,`acl_grant`.`role_id`,`acl_grant`.`permission`,`acl_grant`.`admin_user_id`,`acl_grant`.`customer_id`,`acl_grant`.`created_at`,`acl_grant`.`updated_at`) VALUES (?,?,?,?,?,?,?,?,?)"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.ResourceID),
		orm.ToSQLString(t.RoleID),
		orm.ToSQLString(enumACLGrantPermissionToString(&t.Permission)),
		orm.ToSQLString(t.AdminUserID),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
	)
}

// DBCreateTx will create a new ACLGrant record in the database using the provided transaction
func (t *ACLGrant) DBCreateTx(ctx context.Context, tx *sql.Tx) (sql.Result, error) {
	q := "INSERT INTO `acl_grant` (`acl_grant`.`id`,`acl_grant`.`checksum`,`acl_grant`.`resource_id`,`acl_grant`.`role_id`,`acl_grant`.`permission`,`acl_grant`.`admin_user_id`,`acl_grant`.`customer_id`,`acl_grant`.`created_at`,`acl_grant`.`updated_at`) VALUES (?,?,?,?,?,?,?,?,?)"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.ResourceID),
		orm.ToSQLString(t.RoleID),
		orm.ToSQLString(enumACLGrantPermissionToString(&t.Permission)),
		orm.ToSQLString(t.AdminUserID),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
	)
}

// DBCreateIgnoreDuplicate will upsert the ACLGrant record in the database
func (t *ACLGrant) DBCreateIgnoreDuplicate(ctx context.Context, db *sql.DB) (sql.Result, error) {
	q := "INSERT INTO `acl_grant` (`acl_grant`.`id`,`acl_grant`.`checksum`,`acl_grant`.`resource_id`,`acl_grant`.`role_id`,`acl_grant`.`permission`,`acl_grant`.`admin_user_id`,`acl_grant`.`customer_id`,`acl_grant`.`created_at`,`acl_grant`.`updated_at`) VALUES (?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.ResourceID),
		orm.ToSQLString(t.RoleID),
		orm.ToSQLString(enumACLGrantPermissionToString(&t.Permission)),
		orm.ToSQLString(t.AdminUserID),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
	)
}

// DBCreateIgnoreDuplicateTx will upsert the ACLGrant record in the database using the provided transaction
func (t *ACLGrant) DBCreateIgnoreDuplicateTx(ctx context.Context, tx *sql.Tx) (sql.Result, error) {
	q := "INSERT INTO `acl_grant` (`acl_grant`.`id`,`acl_grant`.`checksum`,`acl_grant`.`resource_id`,`acl_grant`.`role_id`,`acl_grant`.`permission`,`acl_grant`.`admin_user_id`,`acl_grant`.`customer_id`,`acl_grant`.`created_at`,`acl_grant`.`updated_at`) VALUES (?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.ResourceID),
		orm.ToSQLString(t.RoleID),
		orm.ToSQLString(enumACLGrantPermissionToString(&t.Permission)),
		orm.ToSQLString(t.AdminUserID),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
	)
}

// DeleteAllACLGrants deletes all ACLGrant records in the database with optional filters
func DeleteAllACLGrants(ctx context.Context, db *sql.DB, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(ACLGrantTableName),
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

// DeleteAllACLGrantsTx deletes all ACLGrant records in the database with optional filters using the provided transaction
func DeleteAllACLGrantsTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(ACLGrantTableName),
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

// DBDelete will delete this ACLGrant record in the database
func (t *ACLGrant) DBDelete(ctx context.Context, db *sql.DB) (bool, error) {
	q := "DELETE FROM `acl_grant` WHERE `id` = ?"
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

// DBDeleteTx will delete this ACLGrant record in the database using the provided transaction
func (t *ACLGrant) DBDeleteTx(ctx context.Context, tx *sql.Tx) (bool, error) {
	q := "DELETE FROM `acl_grant` WHERE `id` = ?"
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

// DBUpdate will update the ACLGrant record in the database
func (t *ACLGrant) DBUpdate(ctx context.Context, db *sql.DB) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `acl_grant` SET `checksum`=?,`resource_id`=?,`role_id`=?,`permission`=?,`admin_user_id`=?,`customer_id`=?,`created_at`=?,`updated_at`=? WHERE `id`=?"
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.ResourceID),
		orm.ToSQLString(t.RoleID),
		orm.ToSQLString(enumACLGrantPermissionToString(&t.Permission)),
		orm.ToSQLString(t.AdminUserID),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
		orm.ToSQLString(t.ID),
	)
}

// DBUpdateTx will update the ACLGrant record in the database using the provided transaction
func (t *ACLGrant) DBUpdateTx(ctx context.Context, tx *sql.Tx) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `acl_grant` SET `checksum`=?,`resource_id`=?,`role_id`=?,`permission`=?,`admin_user_id`=?,`customer_id`=?,`created_at`=?,`updated_at`=? WHERE `id`=?"
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.ResourceID),
		orm.ToSQLString(t.RoleID),
		orm.ToSQLString(enumACLGrantPermissionToString(&t.Permission)),
		orm.ToSQLString(t.AdminUserID),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
		orm.ToSQLString(t.ID),
	)
}

// DBUpsert will upsert the ACLGrant record in the database
func (t *ACLGrant) DBUpsert(ctx context.Context, db *sql.DB, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `acl_grant` (`acl_grant`.`id`,`acl_grant`.`checksum`,`acl_grant`.`resource_id`,`acl_grant`.`role_id`,`acl_grant`.`permission`,`acl_grant`.`admin_user_id`,`acl_grant`.`customer_id`,`acl_grant`.`created_at`,`acl_grant`.`updated_at`) VALUES (?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `acl_grant` (`acl_grant`.`id`,`acl_grant`.`checksum`,`acl_grant`.`resource_id`,`acl_grant`.`role_id`,`acl_grant`.`permission`,`acl_grant`.`admin_user_id`,`acl_grant`.`customer_id`,`acl_grant`.`created_at`,`acl_grant`.`updated_at`) VALUES (?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`resource_id`=VALUES(`resource_id`),`role_id`=VALUES(`role_id`),`permission`=VALUES(`permission`),`admin_user_id`=VALUES(`admin_user_id`),`customer_id`=VALUES(`customer_id`),`created_at`=VALUES(`created_at`),`updated_at`=VALUES(`updated_at`)"
	}
	r, err := db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.ResourceID),
		orm.ToSQLString(t.RoleID),
		orm.ToSQLString(enumACLGrantPermissionToString(&t.Permission)),
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

// DBUpsertTx will upsert the ACLGrant record in the database using the provided transaction
func (t *ACLGrant) DBUpsertTx(ctx context.Context, tx *sql.Tx, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `acl_grant` (`acl_grant`.`id`,`acl_grant`.`checksum`,`acl_grant`.`resource_id`,`acl_grant`.`role_id`,`acl_grant`.`permission`,`acl_grant`.`admin_user_id`,`acl_grant`.`customer_id`,`acl_grant`.`created_at`,`acl_grant`.`updated_at`) VALUES (?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `acl_grant` (`acl_grant`.`id`,`acl_grant`.`checksum`,`acl_grant`.`resource_id`,`acl_grant`.`role_id`,`acl_grant`.`permission`,`acl_grant`.`admin_user_id`,`acl_grant`.`customer_id`,`acl_grant`.`created_at`,`acl_grant`.`updated_at`) VALUES (?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`resource_id`=VALUES(`resource_id`),`role_id`=VALUES(`role_id`),`permission`=VALUES(`permission`),`admin_user_id`=VALUES(`admin_user_id`),`customer_id`=VALUES(`customer_id`),`created_at`=VALUES(`created_at`),`updated_at`=VALUES(`updated_at`)"
	}
	r, err := tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.ResourceID),
		orm.ToSQLString(t.RoleID),
		orm.ToSQLString(enumACLGrantPermissionToString(&t.Permission)),
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

// DBFindOne will find a ACLGrant record in the database with the primary key
func (t *ACLGrant) DBFindOne(ctx context.Context, db *sql.DB, value string) (bool, error) {
	q := "SELECT `acl_grant`.`id`,`acl_grant`.`checksum`,`acl_grant`.`resource_id`,`acl_grant`.`role_id`,`acl_grant`.`permission`,`acl_grant`.`admin_user_id`,`acl_grant`.`customer_id`,`acl_grant`.`created_at`,`acl_grant`.`updated_at` FROM `acl_grant` WHERE `id` = ? LIMIT 1"
	row := db.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _ResourceID sql.NullString
	var _RoleID sql.NullString
	var _Permission sql.NullString
	var _AdminUserID sql.NullString
	var _CustomerID sql.NullString
	var _CreatedAt sql.NullInt64
	var _UpdatedAt sql.NullInt64
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_ResourceID,
		&_RoleID,
		&_Permission,
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
	if _ResourceID.Valid {
		t.SetResourceID(_ResourceID.String)
	}
	if _RoleID.Valid {
		t.SetRoleID(_RoleID.String)
	}
	if _Permission.Valid {
		t.SetPermissionString(_Permission.String)
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

// DBFindOneTx will find a ACLGrant record in the database with the primary key using the provided transaction
func (t *ACLGrant) DBFindOneTx(ctx context.Context, tx *sql.Tx, value string) (bool, error) {
	q := "SELECT `acl_grant`.`id`,`acl_grant`.`checksum`,`acl_grant`.`resource_id`,`acl_grant`.`role_id`,`acl_grant`.`permission`,`acl_grant`.`admin_user_id`,`acl_grant`.`customer_id`,`acl_grant`.`created_at`,`acl_grant`.`updated_at` FROM `acl_grant` WHERE `id` = ? LIMIT 1"
	row := tx.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _ResourceID sql.NullString
	var _RoleID sql.NullString
	var _Permission sql.NullString
	var _AdminUserID sql.NullString
	var _CustomerID sql.NullString
	var _CreatedAt sql.NullInt64
	var _UpdatedAt sql.NullInt64
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_ResourceID,
		&_RoleID,
		&_Permission,
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
	if _ResourceID.Valid {
		t.SetResourceID(_ResourceID.String)
	}
	if _RoleID.Valid {
		t.SetRoleID(_RoleID.String)
	}
	if _Permission.Valid {
		t.SetPermissionString(_Permission.String)
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

// FindACLGrants will find a ACLGrant record in the database with the provided parameters
func FindACLGrants(ctx context.Context, db *sql.DB, _params ...interface{}) ([]*ACLGrant, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("resource_id"),
		orm.Column("role_id"),
		orm.Column("permission"),
		orm.Column("admin_user_id"),
		orm.Column("customer_id"),
		orm.Column("created_at"),
		orm.Column("updated_at"),
		orm.Table(ACLGrantTableName),
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
	results := make([]*ACLGrant, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _ResourceID sql.NullString
		var _RoleID sql.NullString
		var _Permission sql.NullString
		var _AdminUserID sql.NullString
		var _CustomerID sql.NullString
		var _CreatedAt sql.NullInt64
		var _UpdatedAt sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_ResourceID,
			&_RoleID,
			&_Permission,
			&_AdminUserID,
			&_CustomerID,
			&_CreatedAt,
			&_UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		t := &ACLGrant{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _ResourceID.Valid {
			t.SetResourceID(_ResourceID.String)
		}
		if _RoleID.Valid {
			t.SetRoleID(_RoleID.String)
		}
		if _Permission.Valid {
			t.SetPermissionString(_Permission.String)
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

// FindACLGrantsTx will find a ACLGrant record in the database with the provided parameters using the provided transaction
func FindACLGrantsTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) ([]*ACLGrant, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("resource_id"),
		orm.Column("role_id"),
		orm.Column("permission"),
		orm.Column("admin_user_id"),
		orm.Column("customer_id"),
		orm.Column("created_at"),
		orm.Column("updated_at"),
		orm.Table(ACLGrantTableName),
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
	results := make([]*ACLGrant, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _ResourceID sql.NullString
		var _RoleID sql.NullString
		var _Permission sql.NullString
		var _AdminUserID sql.NullString
		var _CustomerID sql.NullString
		var _CreatedAt sql.NullInt64
		var _UpdatedAt sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_ResourceID,
			&_RoleID,
			&_Permission,
			&_AdminUserID,
			&_CustomerID,
			&_CreatedAt,
			&_UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		t := &ACLGrant{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _ResourceID.Valid {
			t.SetResourceID(_ResourceID.String)
		}
		if _RoleID.Valid {
			t.SetRoleID(_RoleID.String)
		}
		if _Permission.Valid {
			t.SetPermissionString(_Permission.String)
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

// DBFind will find a ACLGrant record in the database with the provided parameters
func (t *ACLGrant) DBFind(ctx context.Context, db *sql.DB, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("resource_id"),
		orm.Column("role_id"),
		orm.Column("permission"),
		orm.Column("admin_user_id"),
		orm.Column("customer_id"),
		orm.Column("created_at"),
		orm.Column("updated_at"),
		orm.Table(ACLGrantTableName),
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
	var _ResourceID sql.NullString
	var _RoleID sql.NullString
	var _Permission sql.NullString
	var _AdminUserID sql.NullString
	var _CustomerID sql.NullString
	var _CreatedAt sql.NullInt64
	var _UpdatedAt sql.NullInt64
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_ResourceID,
		&_RoleID,
		&_Permission,
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
	if _ResourceID.Valid {
		t.SetResourceID(_ResourceID.String)
	}
	if _RoleID.Valid {
		t.SetRoleID(_RoleID.String)
	}
	if _Permission.Valid {
		t.SetPermissionString(_Permission.String)
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

// DBFindTx will find a ACLGrant record in the database with the provided parameters using the provided transaction
func (t *ACLGrant) DBFindTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("resource_id"),
		orm.Column("role_id"),
		orm.Column("permission"),
		orm.Column("admin_user_id"),
		orm.Column("customer_id"),
		orm.Column("created_at"),
		orm.Column("updated_at"),
		orm.Table(ACLGrantTableName),
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
	var _ResourceID sql.NullString
	var _RoleID sql.NullString
	var _Permission sql.NullString
	var _AdminUserID sql.NullString
	var _CustomerID sql.NullString
	var _CreatedAt sql.NullInt64
	var _UpdatedAt sql.NullInt64
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_ResourceID,
		&_RoleID,
		&_Permission,
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
	if _ResourceID.Valid {
		t.SetResourceID(_ResourceID.String)
	}
	if _RoleID.Valid {
		t.SetRoleID(_RoleID.String)
	}
	if _Permission.Valid {
		t.SetPermissionString(_Permission.String)
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

// CountACLGrants will find the count of ACLGrant records in the database
func CountACLGrants(ctx context.Context, db *sql.DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(ACLGrantTableName),
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

// CountACLGrantsTx will find the count of ACLGrant records in the database using the provided transaction
func CountACLGrantsTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(ACLGrantTableName),
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

// DBCount will find the count of ACLGrant records in the database
func (t *ACLGrant) DBCount(ctx context.Context, db *sql.DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(ACLGrantTableName),
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

// DBCountTx will find the count of ACLGrant records in the database using the provided transaction
func (t *ACLGrant) DBCountTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(ACLGrantTableName),
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

// DBExists will return true if the ACLGrant record exists in the database
func (t *ACLGrant) DBExists(ctx context.Context, db *sql.DB) (bool, error) {
	q := "SELECT `id` FROM `acl_grant` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := db.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// DBExistsTx will return true if the ACLGrant record exists in the database using the provided transaction
func (t *ACLGrant) DBExistsTx(ctx context.Context, tx *sql.Tx) (bool, error) {
	q := "SELECT `id` FROM `acl_grant` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := tx.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// PrimaryKeyColumn returns the column name for the primary key
func (t *ACLGrant) PrimaryKeyColumn() string {
	return ACLGrantColumnID
}

// PrimaryKeyColumnType returns the primary key column Go type as a string
func (t *ACLGrant) PrimaryKeyColumnType() string {
	return "string"
}

// PrimaryKey returns the primary key column value
func (t *ACLGrant) PrimaryKey() interface{} {
	return t.ID
}
