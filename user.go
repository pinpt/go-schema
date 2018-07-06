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

var _ Model = (*User)(nil)
var _ CSVWriter = (*User)(nil)
var _ JSONWriter = (*User)(nil)
var _ Checksum = (*User)(nil)

// UserTableName is the name of the table in SQL
const UserTableName = "user"

var UserColumns = []string{
	"id",
	"checksum",
	"customer_id",
	"email",
	"name",
	"username",
	"active",
	"autoenrolled",
	"visible",
	"avatar_url",
	"employee_id",
	"location",
	"region",
	"department",
	"reports_to_user_id",
	"title",
	"role",
	"created_at",
	"updated_at",
	"metadata",
	"ref_id",
	"ref_type",
}

// User table
type User struct {
	Active          bool    `json:"active"`
	Autoenrolled    bool    `json:"autoenrolled"`
	AvatarURL       string  `json:"avatar_url"`
	Checksum        *string `json:"checksum,omitempty"`
	CreatedAt       int64   `json:"created_at"`
	CustomerID      string  `json:"customer_id"`
	Department      *string `json:"department,omitempty"`
	Email           *string `json:"email,omitempty"`
	EmployeeID      *string `json:"employee_id,omitempty"`
	ID              string  `json:"id"`
	Location        *string `json:"location,omitempty"`
	Metadata        *string `json:"metadata,omitempty"`
	Name            *string `json:"name,omitempty"`
	RefID           *string `json:"ref_id,omitempty"`
	RefType         *string `json:"ref_type,omitempty"`
	Region          *string `json:"region,omitempty"`
	ReportsToUserID *string `json:"reports_to_user_id,omitempty"`
	Role            *string `json:"role,omitempty"`
	Title           *string `json:"title,omitempty"`
	UpdatedAt       *int64  `json:"updated_at,omitempty"`
	Username        string  `json:"username"`
	Visible         bool    `json:"visible"`
}

// TableName returns the SQL table name for User and satifies the Model interface
func (t *User) TableName() string {
	return UserTableName
}

// ToCSV will serialize the User instance to a CSV compatible array of strings
func (t *User) ToCSV() []string {
	return []string{
		t.ID,
		t.CalculateChecksum(),
		t.CustomerID,
		toCSVString(t.Email),
		toCSVString(t.Name),
		t.Username,
		toCSVBool(t.Active),
		toCSVBool(t.Autoenrolled),
		toCSVBool(t.Visible),
		t.AvatarURL,
		toCSVString(t.EmployeeID),
		toCSVString(t.Location),
		toCSVString(t.Region),
		toCSVString(t.Department),
		toCSVString(t.ReportsToUserID),
		toCSVString(t.Title),
		toCSVString(t.Role),
		toCSVString(t.CreatedAt),
		toCSVString(t.UpdatedAt),
		toCSVString(t.Metadata),
		toCSVString(t.RefID),
		toCSVString(t.RefType),
	}
}

// WriteCSV will serialize the User instance to the writer as CSV and satisfies the CSVWriter interface
func (t *User) WriteCSV(w *csv.Writer) error {
	return w.Write(t.ToCSV())
}

// WriteJSON will serialize the User instance to the writer as JSON and satisfies the JSONWriter interface
func (t *User) WriteJSON(w io.Writer, indent ...bool) error {
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

// NewUserReader creates a JSON reader which can read in User objects serialized as JSON either as an array, single object or json new lines
// and writes each User to the channel provided
func NewUserReader(r io.Reader, ch chan<- User) error {
	return orm.Deserialize(r, func(buf json.RawMessage) error {
		dec := json.NewDecoder(bytes.NewBuffer(buf))
		e := User{}
		if err := dec.Decode(&e); err != nil {
			return err
		}
		ch <- e
		return nil
	})
}

// NewCSVUserReaderDir will read the reader as CSV and emit each record to the channel provided
func NewCSVUserReader(r io.Reader, ch chan<- User) error {
	cr := csv.NewReader(r)
	for {
		record, err := cr.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		ch <- User{
			ID:              record[0],
			Checksum:        fromStringPointer(record[1]),
			CustomerID:      record[2],
			Email:           fromStringPointer(record[3]),
			Name:            fromStringPointer(record[4]),
			Username:        record[5],
			Active:          fromCSVBool(record[6]),
			Autoenrolled:    fromCSVBool(record[7]),
			Visible:         fromCSVBool(record[8]),
			AvatarURL:       record[9],
			EmployeeID:      fromStringPointer(record[10]),
			Location:        fromStringPointer(record[11]),
			Region:          fromStringPointer(record[12]),
			Department:      fromStringPointer(record[13]),
			ReportsToUserID: fromStringPointer(record[14]),
			Title:           fromStringPointer(record[15]),
			Role:            fromStringPointer(record[16]),
			CreatedAt:       fromCSVInt64(record[17]),
			UpdatedAt:       fromCSVInt64Pointer(record[18]),
			Metadata:        fromStringPointer(record[19]),
			RefID:           fromStringPointer(record[20]),
			RefType:         fromStringPointer(record[21]),
		}
	}
	return nil
}

// NewCSVUserReaderFile will read the file as a CSV and emit each record to the channel provided
func NewCSVUserReaderFile(fp string, ch chan<- User) error {
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
	return NewCSVUserReader(fc, ch)
}

// NewCSVUserReaderDir will read the user.csv.gz file as a CSV and emit each record to the channel provided
func NewCSVUserReaderDir(dir string, ch chan<- User) error {
	return NewCSVUserReaderFile(filepath.Join(dir, "user.csv.gz"), ch)
}

// UserCSVDeduper is a function callback which takes the existing value (a) and the new value (b)
// and the return value should be the one to use (or a new one, if applicable). return nil
// to skip processing of this record
type UserCSVDeduper func(a User, b User) *User

// UserCSVDedupeDisabled is set on whether the CSV writer should de-dupe values by key
var UserCSVDedupeDisabled bool

// NewUserCSVWriterSize creates a batch writer that will write each User into a CSV file
// this method will automatically de-duplicate entries using the primary key. if the checksum
// for a newer item with the same primary key doesn't match a previously sent item, the newer
// one will be used
func NewUserCSVWriterSize(w io.Writer, size int, dedupers ...UserCSVDeduper) (chan User, chan bool, error) {
	cw := csv.NewWriter(w)
	ch := make(chan User, size)
	done := make(chan bool)
	go func() {
		defer func() { done <- true }()
		dodedupe := !UserCSVDedupeDisabled
		var kv map[string]*User
		var deduper UserCSVDeduper
		if dedupers != nil && len(dedupers) > 0 {
			deduper = dedupers[0]
			dodedupe = true
		}
		if dodedupe {
			kv = make(map[string]*User)
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

// UserCSVDefaultSize is the default channel buffer size if not provided
var UserCSVDefaultSize = 100

// NewUserCSVWriter creates a batch writer that will write each User into a CSV file
func NewUserCSVWriter(w io.Writer, dedupers ...UserCSVDeduper) (chan User, chan bool, error) {
	return NewUserCSVWriterSize(w, UserCSVDefaultSize, dedupers...)
}

// NewUserCSVWriterDir creates a batch writer that will write each User into a CSV file named user.csv.gz in dir
func NewUserCSVWriterDir(dir string, dedupers ...UserCSVDeduper) (chan User, chan bool, error) {
	return NewUserCSVWriterFile(filepath.Join(dir, "user.csv.gz"), dedupers...)
}

// NewUserCSVWriterFile creates a batch writer that will write each User into a CSV file
func NewUserCSVWriterFile(fn string, dedupers ...UserCSVDeduper) (chan User, chan bool, error) {
	f, err := os.Create(fn)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening CSV file %s. %v", fn, err)
	}
	var fc io.WriteCloser = f
	if filepath.Ext(fn) == ".gz" {
		w, _ := gzip.NewWriterLevel(f, gzip.BestCompression)
		fc = w
	}
	ch, done, err := NewUserCSVWriter(fc, dedupers...)
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

type UserDBAction func(ctx context.Context, db *sql.DB, record User) error

// NewUserDBWriterSize creates a DB writer that will write each issue into the DB
func NewUserDBWriterSize(ctx context.Context, db *sql.DB, errors chan<- error, size int, actions ...UserDBAction) (chan User, chan bool, error) {
	ch := make(chan User, size)
	done := make(chan bool)
	var action UserDBAction
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

// NewUserDBWriter creates a DB writer that will write each issue into the DB
func NewUserDBWriter(ctx context.Context, db *sql.DB, errors chan<- error, actions ...UserDBAction) (chan User, chan bool, error) {
	return NewUserDBWriterSize(ctx, db, errors, 100, actions...)
}

// UserColumnID is the ID SQL column name for the User table
const UserColumnID = "id"

// UserEscapedColumnID is the escaped ID SQL column name for the User table
const UserEscapedColumnID = "`id`"

// UserColumnChecksum is the Checksum SQL column name for the User table
const UserColumnChecksum = "checksum"

// UserEscapedColumnChecksum is the escaped Checksum SQL column name for the User table
const UserEscapedColumnChecksum = "`checksum`"

// UserColumnCustomerID is the CustomerID SQL column name for the User table
const UserColumnCustomerID = "customer_id"

// UserEscapedColumnCustomerID is the escaped CustomerID SQL column name for the User table
const UserEscapedColumnCustomerID = "`customer_id`"

// UserColumnEmail is the Email SQL column name for the User table
const UserColumnEmail = "email"

// UserEscapedColumnEmail is the escaped Email SQL column name for the User table
const UserEscapedColumnEmail = "`email`"

// UserColumnName is the Name SQL column name for the User table
const UserColumnName = "name"

// UserEscapedColumnName is the escaped Name SQL column name for the User table
const UserEscapedColumnName = "`name`"

// UserColumnUsername is the Username SQL column name for the User table
const UserColumnUsername = "username"

// UserEscapedColumnUsername is the escaped Username SQL column name for the User table
const UserEscapedColumnUsername = "`username`"

// UserColumnActive is the Active SQL column name for the User table
const UserColumnActive = "active"

// UserEscapedColumnActive is the escaped Active SQL column name for the User table
const UserEscapedColumnActive = "`active`"

// UserColumnAutoenrolled is the Autoenrolled SQL column name for the User table
const UserColumnAutoenrolled = "autoenrolled"

// UserEscapedColumnAutoenrolled is the escaped Autoenrolled SQL column name for the User table
const UserEscapedColumnAutoenrolled = "`autoenrolled`"

// UserColumnVisible is the Visible SQL column name for the User table
const UserColumnVisible = "visible"

// UserEscapedColumnVisible is the escaped Visible SQL column name for the User table
const UserEscapedColumnVisible = "`visible`"

// UserColumnAvatarURL is the AvatarURL SQL column name for the User table
const UserColumnAvatarURL = "avatar_url"

// UserEscapedColumnAvatarURL is the escaped AvatarURL SQL column name for the User table
const UserEscapedColumnAvatarURL = "`avatar_url`"

// UserColumnEmployeeID is the EmployeeID SQL column name for the User table
const UserColumnEmployeeID = "employee_id"

// UserEscapedColumnEmployeeID is the escaped EmployeeID SQL column name for the User table
const UserEscapedColumnEmployeeID = "`employee_id`"

// UserColumnLocation is the Location SQL column name for the User table
const UserColumnLocation = "location"

// UserEscapedColumnLocation is the escaped Location SQL column name for the User table
const UserEscapedColumnLocation = "`location`"

// UserColumnRegion is the Region SQL column name for the User table
const UserColumnRegion = "region"

// UserEscapedColumnRegion is the escaped Region SQL column name for the User table
const UserEscapedColumnRegion = "`region`"

// UserColumnDepartment is the Department SQL column name for the User table
const UserColumnDepartment = "department"

// UserEscapedColumnDepartment is the escaped Department SQL column name for the User table
const UserEscapedColumnDepartment = "`department`"

// UserColumnReportsToUserID is the ReportsToUserID SQL column name for the User table
const UserColumnReportsToUserID = "reports_to_user_id"

// UserEscapedColumnReportsToUserID is the escaped ReportsToUserID SQL column name for the User table
const UserEscapedColumnReportsToUserID = "`reports_to_user_id`"

// UserColumnTitle is the Title SQL column name for the User table
const UserColumnTitle = "title"

// UserEscapedColumnTitle is the escaped Title SQL column name for the User table
const UserEscapedColumnTitle = "`title`"

// UserColumnRole is the Role SQL column name for the User table
const UserColumnRole = "role"

// UserEscapedColumnRole is the escaped Role SQL column name for the User table
const UserEscapedColumnRole = "`role`"

// UserColumnCreatedAt is the CreatedAt SQL column name for the User table
const UserColumnCreatedAt = "created_at"

// UserEscapedColumnCreatedAt is the escaped CreatedAt SQL column name for the User table
const UserEscapedColumnCreatedAt = "`created_at`"

// UserColumnUpdatedAt is the UpdatedAt SQL column name for the User table
const UserColumnUpdatedAt = "updated_at"

// UserEscapedColumnUpdatedAt is the escaped UpdatedAt SQL column name for the User table
const UserEscapedColumnUpdatedAt = "`updated_at`"

// UserColumnMetadata is the Metadata SQL column name for the User table
const UserColumnMetadata = "metadata"

// UserEscapedColumnMetadata is the escaped Metadata SQL column name for the User table
const UserEscapedColumnMetadata = "`metadata`"

// UserColumnRefID is the RefID SQL column name for the User table
const UserColumnRefID = "ref_id"

// UserEscapedColumnRefID is the escaped RefID SQL column name for the User table
const UserEscapedColumnRefID = "`ref_id`"

// UserColumnRefType is the RefType SQL column name for the User table
const UserColumnRefType = "ref_type"

// UserEscapedColumnRefType is the escaped RefType SQL column name for the User table
const UserEscapedColumnRefType = "`ref_type`"

// GetID will return the User ID value
func (t *User) GetID() string {
	return t.ID
}

// SetID will set the User ID value
func (t *User) SetID(v string) {
	t.ID = v
}

// FindUserByID will find a User by ID
func FindUserByID(ctx context.Context, db *sql.DB, value string) (*User, error) {
	q := "SELECT `user`.`id`,`user`.`checksum`,`user`.`customer_id`,`user`.`email`,`user`.`name`,`user`.`username`,`user`.`active`,`user`.`autoenrolled`,`user`.`visible`,`user`.`avatar_url`,`user`.`employee_id`,`user`.`location`,`user`.`region`,`user`.`department`,`user`.`reports_to_user_id`,`user`.`title`,`user`.`role`,`user`.`created_at`,`user`.`updated_at`,`user`.`metadata`,`user`.`ref_id`,`user`.`ref_type` FROM `user` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CustomerID sql.NullString
	var _Email sql.NullString
	var _Name sql.NullString
	var _Username sql.NullString
	var _Active sql.NullBool
	var _Autoenrolled sql.NullBool
	var _Visible sql.NullBool
	var _AvatarURL sql.NullString
	var _EmployeeID sql.NullString
	var _Location sql.NullString
	var _Region sql.NullString
	var _Department sql.NullString
	var _ReportsToUserID sql.NullString
	var _Title sql.NullString
	var _Role sql.NullString
	var _CreatedAt sql.NullInt64
	var _UpdatedAt sql.NullInt64
	var _Metadata sql.NullString
	var _RefID sql.NullString
	var _RefType sql.NullString
	err := db.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_Email,
		&_Name,
		&_Username,
		&_Active,
		&_Autoenrolled,
		&_Visible,
		&_AvatarURL,
		&_EmployeeID,
		&_Location,
		&_Region,
		&_Department,
		&_ReportsToUserID,
		&_Title,
		&_Role,
		&_CreatedAt,
		&_UpdatedAt,
		&_Metadata,
		&_RefID,
		&_RefType,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &User{}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _Checksum.Valid {
		t.SetChecksum(_Checksum.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _Email.Valid {
		t.SetEmail(_Email.String)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _Username.Valid {
		t.SetUsername(_Username.String)
	}
	if _Active.Valid {
		t.SetActive(_Active.Bool)
	}
	if _Autoenrolled.Valid {
		t.SetAutoenrolled(_Autoenrolled.Bool)
	}
	if _Visible.Valid {
		t.SetVisible(_Visible.Bool)
	}
	if _AvatarURL.Valid {
		t.SetAvatarURL(_AvatarURL.String)
	}
	if _EmployeeID.Valid {
		t.SetEmployeeID(_EmployeeID.String)
	}
	if _Location.Valid {
		t.SetLocation(_Location.String)
	}
	if _Region.Valid {
		t.SetRegion(_Region.String)
	}
	if _Department.Valid {
		t.SetDepartment(_Department.String)
	}
	if _ReportsToUserID.Valid {
		t.SetReportsToUserID(_ReportsToUserID.String)
	}
	if _Title.Valid {
		t.SetTitle(_Title.String)
	}
	if _Role.Valid {
		t.SetRole(_Role.String)
	}
	if _CreatedAt.Valid {
		t.SetCreatedAt(_CreatedAt.Int64)
	}
	if _UpdatedAt.Valid {
		t.SetUpdatedAt(_UpdatedAt.Int64)
	}
	if _Metadata.Valid {
		t.SetMetadata(_Metadata.String)
	}
	if _RefID.Valid {
		t.SetRefID(_RefID.String)
	}
	if _RefType.Valid {
		t.SetRefType(_RefType.String)
	}
	return t, nil
}

// FindUserByIDTx will find a User by ID using the provided transaction
func FindUserByIDTx(ctx context.Context, tx *sql.Tx, value string) (*User, error) {
	q := "SELECT `user`.`id`,`user`.`checksum`,`user`.`customer_id`,`user`.`email`,`user`.`name`,`user`.`username`,`user`.`active`,`user`.`autoenrolled`,`user`.`visible`,`user`.`avatar_url`,`user`.`employee_id`,`user`.`location`,`user`.`region`,`user`.`department`,`user`.`reports_to_user_id`,`user`.`title`,`user`.`role`,`user`.`created_at`,`user`.`updated_at`,`user`.`metadata`,`user`.`ref_id`,`user`.`ref_type` FROM `user` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CustomerID sql.NullString
	var _Email sql.NullString
	var _Name sql.NullString
	var _Username sql.NullString
	var _Active sql.NullBool
	var _Autoenrolled sql.NullBool
	var _Visible sql.NullBool
	var _AvatarURL sql.NullString
	var _EmployeeID sql.NullString
	var _Location sql.NullString
	var _Region sql.NullString
	var _Department sql.NullString
	var _ReportsToUserID sql.NullString
	var _Title sql.NullString
	var _Role sql.NullString
	var _CreatedAt sql.NullInt64
	var _UpdatedAt sql.NullInt64
	var _Metadata sql.NullString
	var _RefID sql.NullString
	var _RefType sql.NullString
	err := tx.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_Email,
		&_Name,
		&_Username,
		&_Active,
		&_Autoenrolled,
		&_Visible,
		&_AvatarURL,
		&_EmployeeID,
		&_Location,
		&_Region,
		&_Department,
		&_ReportsToUserID,
		&_Title,
		&_Role,
		&_CreatedAt,
		&_UpdatedAt,
		&_Metadata,
		&_RefID,
		&_RefType,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &User{}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _Checksum.Valid {
		t.SetChecksum(_Checksum.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _Email.Valid {
		t.SetEmail(_Email.String)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _Username.Valid {
		t.SetUsername(_Username.String)
	}
	if _Active.Valid {
		t.SetActive(_Active.Bool)
	}
	if _Autoenrolled.Valid {
		t.SetAutoenrolled(_Autoenrolled.Bool)
	}
	if _Visible.Valid {
		t.SetVisible(_Visible.Bool)
	}
	if _AvatarURL.Valid {
		t.SetAvatarURL(_AvatarURL.String)
	}
	if _EmployeeID.Valid {
		t.SetEmployeeID(_EmployeeID.String)
	}
	if _Location.Valid {
		t.SetLocation(_Location.String)
	}
	if _Region.Valid {
		t.SetRegion(_Region.String)
	}
	if _Department.Valid {
		t.SetDepartment(_Department.String)
	}
	if _ReportsToUserID.Valid {
		t.SetReportsToUserID(_ReportsToUserID.String)
	}
	if _Title.Valid {
		t.SetTitle(_Title.String)
	}
	if _Role.Valid {
		t.SetRole(_Role.String)
	}
	if _CreatedAt.Valid {
		t.SetCreatedAt(_CreatedAt.Int64)
	}
	if _UpdatedAt.Valid {
		t.SetUpdatedAt(_UpdatedAt.Int64)
	}
	if _Metadata.Valid {
		t.SetMetadata(_Metadata.String)
	}
	if _RefID.Valid {
		t.SetRefID(_RefID.String)
	}
	if _RefType.Valid {
		t.SetRefType(_RefType.String)
	}
	return t, nil
}

// GetChecksum will return the User Checksum value
func (t *User) GetChecksum() string {
	if t.Checksum == nil {
		return ""
	}
	return *t.Checksum
}

// SetChecksum will set the User Checksum value
func (t *User) SetChecksum(v string) {
	t.Checksum = &v
}

// GetCustomerID will return the User CustomerID value
func (t *User) GetCustomerID() string {
	return t.CustomerID
}

// SetCustomerID will set the User CustomerID value
func (t *User) SetCustomerID(v string) {
	t.CustomerID = v
}

// GetEmail will return the User Email value
func (t *User) GetEmail() string {
	if t.Email == nil {
		return ""
	}
	return *t.Email
}

// SetEmail will set the User Email value
func (t *User) SetEmail(v string) {
	t.Email = &v
}

// FindUsersByEmail will find all Users by the Email value
func FindUsersByEmail(ctx context.Context, db *sql.DB, value string) ([]*User, error) {
	q := "SELECT `user`.`id`,`user`.`checksum`,`user`.`customer_id`,`user`.`email`,`user`.`name`,`user`.`username`,`user`.`active`,`user`.`autoenrolled`,`user`.`visible`,`user`.`avatar_url`,`user`.`employee_id`,`user`.`location`,`user`.`region`,`user`.`department`,`user`.`reports_to_user_id`,`user`.`title`,`user`.`role`,`user`.`created_at`,`user`.`updated_at`,`user`.`metadata`,`user`.`ref_id`,`user`.`ref_type` FROM `user` WHERE `email` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*User, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _Email sql.NullString
		var _Name sql.NullString
		var _Username sql.NullString
		var _Active sql.NullBool
		var _Autoenrolled sql.NullBool
		var _Visible sql.NullBool
		var _AvatarURL sql.NullString
		var _EmployeeID sql.NullString
		var _Location sql.NullString
		var _Region sql.NullString
		var _Department sql.NullString
		var _ReportsToUserID sql.NullString
		var _Title sql.NullString
		var _Role sql.NullString
		var _CreatedAt sql.NullInt64
		var _UpdatedAt sql.NullInt64
		var _Metadata sql.NullString
		var _RefID sql.NullString
		var _RefType sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_Email,
			&_Name,
			&_Username,
			&_Active,
			&_Autoenrolled,
			&_Visible,
			&_AvatarURL,
			&_EmployeeID,
			&_Location,
			&_Region,
			&_Department,
			&_ReportsToUserID,
			&_Title,
			&_Role,
			&_CreatedAt,
			&_UpdatedAt,
			&_Metadata,
			&_RefID,
			&_RefType,
		)
		if err != nil {
			return nil, err
		}
		t := &User{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _Email.Valid {
			t.SetEmail(_Email.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _Username.Valid {
			t.SetUsername(_Username.String)
		}
		if _Active.Valid {
			t.SetActive(_Active.Bool)
		}
		if _Autoenrolled.Valid {
			t.SetAutoenrolled(_Autoenrolled.Bool)
		}
		if _Visible.Valid {
			t.SetVisible(_Visible.Bool)
		}
		if _AvatarURL.Valid {
			t.SetAvatarURL(_AvatarURL.String)
		}
		if _EmployeeID.Valid {
			t.SetEmployeeID(_EmployeeID.String)
		}
		if _Location.Valid {
			t.SetLocation(_Location.String)
		}
		if _Region.Valid {
			t.SetRegion(_Region.String)
		}
		if _Department.Valid {
			t.SetDepartment(_Department.String)
		}
		if _ReportsToUserID.Valid {
			t.SetReportsToUserID(_ReportsToUserID.String)
		}
		if _Title.Valid {
			t.SetTitle(_Title.String)
		}
		if _Role.Valid {
			t.SetRole(_Role.String)
		}
		if _CreatedAt.Valid {
			t.SetCreatedAt(_CreatedAt.Int64)
		}
		if _UpdatedAt.Valid {
			t.SetUpdatedAt(_UpdatedAt.Int64)
		}
		if _Metadata.Valid {
			t.SetMetadata(_Metadata.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		if _RefType.Valid {
			t.SetRefType(_RefType.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindUsersByEmailTx will find all Users by the Email value using the provided transaction
func FindUsersByEmailTx(ctx context.Context, tx *sql.Tx, value string) ([]*User, error) {
	q := "SELECT `user`.`id`,`user`.`checksum`,`user`.`customer_id`,`user`.`email`,`user`.`name`,`user`.`username`,`user`.`active`,`user`.`autoenrolled`,`user`.`visible`,`user`.`avatar_url`,`user`.`employee_id`,`user`.`location`,`user`.`region`,`user`.`department`,`user`.`reports_to_user_id`,`user`.`title`,`user`.`role`,`user`.`created_at`,`user`.`updated_at`,`user`.`metadata`,`user`.`ref_id`,`user`.`ref_type` FROM `user` WHERE `email` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*User, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _Email sql.NullString
		var _Name sql.NullString
		var _Username sql.NullString
		var _Active sql.NullBool
		var _Autoenrolled sql.NullBool
		var _Visible sql.NullBool
		var _AvatarURL sql.NullString
		var _EmployeeID sql.NullString
		var _Location sql.NullString
		var _Region sql.NullString
		var _Department sql.NullString
		var _ReportsToUserID sql.NullString
		var _Title sql.NullString
		var _Role sql.NullString
		var _CreatedAt sql.NullInt64
		var _UpdatedAt sql.NullInt64
		var _Metadata sql.NullString
		var _RefID sql.NullString
		var _RefType sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_Email,
			&_Name,
			&_Username,
			&_Active,
			&_Autoenrolled,
			&_Visible,
			&_AvatarURL,
			&_EmployeeID,
			&_Location,
			&_Region,
			&_Department,
			&_ReportsToUserID,
			&_Title,
			&_Role,
			&_CreatedAt,
			&_UpdatedAt,
			&_Metadata,
			&_RefID,
			&_RefType,
		)
		if err != nil {
			return nil, err
		}
		t := &User{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _Email.Valid {
			t.SetEmail(_Email.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _Username.Valid {
			t.SetUsername(_Username.String)
		}
		if _Active.Valid {
			t.SetActive(_Active.Bool)
		}
		if _Autoenrolled.Valid {
			t.SetAutoenrolled(_Autoenrolled.Bool)
		}
		if _Visible.Valid {
			t.SetVisible(_Visible.Bool)
		}
		if _AvatarURL.Valid {
			t.SetAvatarURL(_AvatarURL.String)
		}
		if _EmployeeID.Valid {
			t.SetEmployeeID(_EmployeeID.String)
		}
		if _Location.Valid {
			t.SetLocation(_Location.String)
		}
		if _Region.Valid {
			t.SetRegion(_Region.String)
		}
		if _Department.Valid {
			t.SetDepartment(_Department.String)
		}
		if _ReportsToUserID.Valid {
			t.SetReportsToUserID(_ReportsToUserID.String)
		}
		if _Title.Valid {
			t.SetTitle(_Title.String)
		}
		if _Role.Valid {
			t.SetRole(_Role.String)
		}
		if _CreatedAt.Valid {
			t.SetCreatedAt(_CreatedAt.Int64)
		}
		if _UpdatedAt.Valid {
			t.SetUpdatedAt(_UpdatedAt.Int64)
		}
		if _Metadata.Valid {
			t.SetMetadata(_Metadata.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		if _RefType.Valid {
			t.SetRefType(_RefType.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetName will return the User Name value
func (t *User) GetName() string {
	if t.Name == nil {
		return ""
	}
	return *t.Name
}

// SetName will set the User Name value
func (t *User) SetName(v string) {
	t.Name = &v
}

// GetUsername will return the User Username value
func (t *User) GetUsername() string {
	return t.Username
}

// SetUsername will set the User Username value
func (t *User) SetUsername(v string) {
	t.Username = v
}

// FindUsersByUsername will find all Users by the Username value
func FindUsersByUsername(ctx context.Context, db *sql.DB, value string) ([]*User, error) {
	q := "SELECT `user`.`id`,`user`.`checksum`,`user`.`customer_id`,`user`.`email`,`user`.`name`,`user`.`username`,`user`.`active`,`user`.`autoenrolled`,`user`.`visible`,`user`.`avatar_url`,`user`.`employee_id`,`user`.`location`,`user`.`region`,`user`.`department`,`user`.`reports_to_user_id`,`user`.`title`,`user`.`role`,`user`.`created_at`,`user`.`updated_at`,`user`.`metadata`,`user`.`ref_id`,`user`.`ref_type` FROM `user` WHERE `username` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*User, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _Email sql.NullString
		var _Name sql.NullString
		var _Username sql.NullString
		var _Active sql.NullBool
		var _Autoenrolled sql.NullBool
		var _Visible sql.NullBool
		var _AvatarURL sql.NullString
		var _EmployeeID sql.NullString
		var _Location sql.NullString
		var _Region sql.NullString
		var _Department sql.NullString
		var _ReportsToUserID sql.NullString
		var _Title sql.NullString
		var _Role sql.NullString
		var _CreatedAt sql.NullInt64
		var _UpdatedAt sql.NullInt64
		var _Metadata sql.NullString
		var _RefID sql.NullString
		var _RefType sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_Email,
			&_Name,
			&_Username,
			&_Active,
			&_Autoenrolled,
			&_Visible,
			&_AvatarURL,
			&_EmployeeID,
			&_Location,
			&_Region,
			&_Department,
			&_ReportsToUserID,
			&_Title,
			&_Role,
			&_CreatedAt,
			&_UpdatedAt,
			&_Metadata,
			&_RefID,
			&_RefType,
		)
		if err != nil {
			return nil, err
		}
		t := &User{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _Email.Valid {
			t.SetEmail(_Email.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _Username.Valid {
			t.SetUsername(_Username.String)
		}
		if _Active.Valid {
			t.SetActive(_Active.Bool)
		}
		if _Autoenrolled.Valid {
			t.SetAutoenrolled(_Autoenrolled.Bool)
		}
		if _Visible.Valid {
			t.SetVisible(_Visible.Bool)
		}
		if _AvatarURL.Valid {
			t.SetAvatarURL(_AvatarURL.String)
		}
		if _EmployeeID.Valid {
			t.SetEmployeeID(_EmployeeID.String)
		}
		if _Location.Valid {
			t.SetLocation(_Location.String)
		}
		if _Region.Valid {
			t.SetRegion(_Region.String)
		}
		if _Department.Valid {
			t.SetDepartment(_Department.String)
		}
		if _ReportsToUserID.Valid {
			t.SetReportsToUserID(_ReportsToUserID.String)
		}
		if _Title.Valid {
			t.SetTitle(_Title.String)
		}
		if _Role.Valid {
			t.SetRole(_Role.String)
		}
		if _CreatedAt.Valid {
			t.SetCreatedAt(_CreatedAt.Int64)
		}
		if _UpdatedAt.Valid {
			t.SetUpdatedAt(_UpdatedAt.Int64)
		}
		if _Metadata.Valid {
			t.SetMetadata(_Metadata.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		if _RefType.Valid {
			t.SetRefType(_RefType.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindUsersByUsernameTx will find all Users by the Username value using the provided transaction
func FindUsersByUsernameTx(ctx context.Context, tx *sql.Tx, value string) ([]*User, error) {
	q := "SELECT `user`.`id`,`user`.`checksum`,`user`.`customer_id`,`user`.`email`,`user`.`name`,`user`.`username`,`user`.`active`,`user`.`autoenrolled`,`user`.`visible`,`user`.`avatar_url`,`user`.`employee_id`,`user`.`location`,`user`.`region`,`user`.`department`,`user`.`reports_to_user_id`,`user`.`title`,`user`.`role`,`user`.`created_at`,`user`.`updated_at`,`user`.`metadata`,`user`.`ref_id`,`user`.`ref_type` FROM `user` WHERE `username` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*User, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _Email sql.NullString
		var _Name sql.NullString
		var _Username sql.NullString
		var _Active sql.NullBool
		var _Autoenrolled sql.NullBool
		var _Visible sql.NullBool
		var _AvatarURL sql.NullString
		var _EmployeeID sql.NullString
		var _Location sql.NullString
		var _Region sql.NullString
		var _Department sql.NullString
		var _ReportsToUserID sql.NullString
		var _Title sql.NullString
		var _Role sql.NullString
		var _CreatedAt sql.NullInt64
		var _UpdatedAt sql.NullInt64
		var _Metadata sql.NullString
		var _RefID sql.NullString
		var _RefType sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_Email,
			&_Name,
			&_Username,
			&_Active,
			&_Autoenrolled,
			&_Visible,
			&_AvatarURL,
			&_EmployeeID,
			&_Location,
			&_Region,
			&_Department,
			&_ReportsToUserID,
			&_Title,
			&_Role,
			&_CreatedAt,
			&_UpdatedAt,
			&_Metadata,
			&_RefID,
			&_RefType,
		)
		if err != nil {
			return nil, err
		}
		t := &User{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _Email.Valid {
			t.SetEmail(_Email.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _Username.Valid {
			t.SetUsername(_Username.String)
		}
		if _Active.Valid {
			t.SetActive(_Active.Bool)
		}
		if _Autoenrolled.Valid {
			t.SetAutoenrolled(_Autoenrolled.Bool)
		}
		if _Visible.Valid {
			t.SetVisible(_Visible.Bool)
		}
		if _AvatarURL.Valid {
			t.SetAvatarURL(_AvatarURL.String)
		}
		if _EmployeeID.Valid {
			t.SetEmployeeID(_EmployeeID.String)
		}
		if _Location.Valid {
			t.SetLocation(_Location.String)
		}
		if _Region.Valid {
			t.SetRegion(_Region.String)
		}
		if _Department.Valid {
			t.SetDepartment(_Department.String)
		}
		if _ReportsToUserID.Valid {
			t.SetReportsToUserID(_ReportsToUserID.String)
		}
		if _Title.Valid {
			t.SetTitle(_Title.String)
		}
		if _Role.Valid {
			t.SetRole(_Role.String)
		}
		if _CreatedAt.Valid {
			t.SetCreatedAt(_CreatedAt.Int64)
		}
		if _UpdatedAt.Valid {
			t.SetUpdatedAt(_UpdatedAt.Int64)
		}
		if _Metadata.Valid {
			t.SetMetadata(_Metadata.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		if _RefType.Valid {
			t.SetRefType(_RefType.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetActive will return the User Active value
func (t *User) GetActive() bool {
	return t.Active
}

// SetActive will set the User Active value
func (t *User) SetActive(v bool) {
	t.Active = v
}

// GetAutoenrolled will return the User Autoenrolled value
func (t *User) GetAutoenrolled() bool {
	return t.Autoenrolled
}

// SetAutoenrolled will set the User Autoenrolled value
func (t *User) SetAutoenrolled(v bool) {
	t.Autoenrolled = v
}

// GetVisible will return the User Visible value
func (t *User) GetVisible() bool {
	return t.Visible
}

// SetVisible will set the User Visible value
func (t *User) SetVisible(v bool) {
	t.Visible = v
}

// GetAvatarURL will return the User AvatarURL value
func (t *User) GetAvatarURL() string {
	return t.AvatarURL
}

// SetAvatarURL will set the User AvatarURL value
func (t *User) SetAvatarURL(v string) {
	t.AvatarURL = v
}

// GetEmployeeID will return the User EmployeeID value
func (t *User) GetEmployeeID() string {
	if t.EmployeeID == nil {
		return ""
	}
	return *t.EmployeeID
}

// SetEmployeeID will set the User EmployeeID value
func (t *User) SetEmployeeID(v string) {
	t.EmployeeID = &v
}

// GetLocation will return the User Location value
func (t *User) GetLocation() string {
	if t.Location == nil {
		return ""
	}
	return *t.Location
}

// SetLocation will set the User Location value
func (t *User) SetLocation(v string) {
	t.Location = &v
}

// GetRegion will return the User Region value
func (t *User) GetRegion() string {
	if t.Region == nil {
		return ""
	}
	return *t.Region
}

// SetRegion will set the User Region value
func (t *User) SetRegion(v string) {
	t.Region = &v
}

// GetDepartment will return the User Department value
func (t *User) GetDepartment() string {
	if t.Department == nil {
		return ""
	}
	return *t.Department
}

// SetDepartment will set the User Department value
func (t *User) SetDepartment(v string) {
	t.Department = &v
}

// GetReportsToUserID will return the User ReportsToUserID value
func (t *User) GetReportsToUserID() string {
	if t.ReportsToUserID == nil {
		return ""
	}
	return *t.ReportsToUserID
}

// SetReportsToUserID will set the User ReportsToUserID value
func (t *User) SetReportsToUserID(v string) {
	t.ReportsToUserID = &v
}

// GetTitle will return the User Title value
func (t *User) GetTitle() string {
	if t.Title == nil {
		return ""
	}
	return *t.Title
}

// SetTitle will set the User Title value
func (t *User) SetTitle(v string) {
	t.Title = &v
}

// GetRole will return the User Role value
func (t *User) GetRole() string {
	if t.Role == nil {
		return ""
	}
	return *t.Role
}

// SetRole will set the User Role value
func (t *User) SetRole(v string) {
	t.Role = &v
}

// GetCreatedAt will return the User CreatedAt value
func (t *User) GetCreatedAt() int64 {
	return t.CreatedAt
}

// SetCreatedAt will set the User CreatedAt value
func (t *User) SetCreatedAt(v int64) {
	t.CreatedAt = v
}

// GetUpdatedAt will return the User UpdatedAt value
func (t *User) GetUpdatedAt() int64 {
	if t.UpdatedAt == nil {
		return int64(0)
	}
	return *t.UpdatedAt
}

// SetUpdatedAt will set the User UpdatedAt value
func (t *User) SetUpdatedAt(v int64) {
	t.UpdatedAt = &v
}

// GetMetadata will return the User Metadata value
func (t *User) GetMetadata() string {
	if t.Metadata == nil {
		return ""
	}
	return *t.Metadata
}

// SetMetadata will set the User Metadata value
func (t *User) SetMetadata(v string) {
	t.Metadata = &v
}

// GetRefID will return the User RefID value
func (t *User) GetRefID() string {
	if t.RefID == nil {
		return ""
	}
	return *t.RefID
}

// SetRefID will set the User RefID value
func (t *User) SetRefID(v string) {
	t.RefID = &v
}

// GetRefType will return the User RefType value
func (t *User) GetRefType() string {
	if t.RefType == nil {
		return ""
	}
	return *t.RefType
}

// SetRefType will set the User RefType value
func (t *User) SetRefType(v string) {
	t.RefType = &v
}

func (t *User) toTimestamp(value time.Time) *timestamp.Timestamp {
	ts, _ := ptypes.TimestampProto(value)
	return ts
}

// DBCreateUserTable will create the User table
func DBCreateUserTable(ctx context.Context, db *sql.DB) error {
	q := "CREATE TABLE `user` (`id` VARCHAR(64) NOT NULL PRIMARY KEY,`checksum` CHAR(64),`customer_id` VARCHAR(64) NOT NULL,`email` VARCHAR(255),`name` TEXT,`username` VARCHAR(100) NOT NULL,`active`BOOL NOT NULL,`autoenrolled`BOOL NOT NULL,`visible` BOOL NOT NULL,`avatar_url` VARCHAR(255) NOT NULL,`employee_id` TEXT,`location` TEXT,`region`TEXT,`department` TEXT,`reports_to_user_id` VARCHAR(64),`title` TEXT,`role` TEXT,`created_at` BIGINT(20) UNSIGNED NOT NULL,`updated_at` BIGINT(20) UNSIGNED,`metadata` JSON,`ref_id`VARCHAR(64),`ref_type` VARCHAR(20),INDEX user_email_index (`email`),INDEX user_username_index (`username`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBCreateUserTableTx will create the User table using the provided transction
func DBCreateUserTableTx(ctx context.Context, tx *sql.Tx) error {
	q := "CREATE TABLE `user` (`id` VARCHAR(64) NOT NULL PRIMARY KEY,`checksum` CHAR(64),`customer_id` VARCHAR(64) NOT NULL,`email` VARCHAR(255),`name` TEXT,`username` VARCHAR(100) NOT NULL,`active`BOOL NOT NULL,`autoenrolled`BOOL NOT NULL,`visible` BOOL NOT NULL,`avatar_url` VARCHAR(255) NOT NULL,`employee_id` TEXT,`location` TEXT,`region`TEXT,`department` TEXT,`reports_to_user_id` VARCHAR(64),`title` TEXT,`role` TEXT,`created_at` BIGINT(20) UNSIGNED NOT NULL,`updated_at` BIGINT(20) UNSIGNED,`metadata` JSON,`ref_id`VARCHAR(64),`ref_type` VARCHAR(20),INDEX user_email_index (`email`),INDEX user_username_index (`username`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// DBDropUserTable will drop the User table
func DBDropUserTable(ctx context.Context, db *sql.DB) error {
	q := "DROP TABLE IF EXISTS `user`"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBDropUserTableTx will drop the User table using the provided transaction
func DBDropUserTableTx(ctx context.Context, tx *sql.Tx) error {
	q := "DROP TABLE IF EXISTS `user`"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// CalculateChecksum will calculate a checksum of the SHA1 of all field values
func (t *User) CalculateChecksum() string {
	return orm.HashStrings(
		orm.ToString(t.ID),
		orm.ToString(t.CustomerID),
		orm.ToString(t.Email),
		orm.ToString(t.Name),
		orm.ToString(t.Username),
		orm.ToString(t.Active),
		orm.ToString(t.Autoenrolled),
		orm.ToString(t.Visible),
		orm.ToString(t.AvatarURL),
		orm.ToString(t.EmployeeID),
		orm.ToString(t.Location),
		orm.ToString(t.Region),
		orm.ToString(t.Department),
		orm.ToString(t.ReportsToUserID),
		orm.ToString(t.Title),
		orm.ToString(t.Role),
		orm.ToString(t.CreatedAt),
		orm.ToString(t.UpdatedAt),
		orm.ToString(t.Metadata),
		orm.ToString(t.RefID),
		orm.ToString(t.RefType),
	)
}

// DBCreate will create a new User record in the database
func (t *User) DBCreate(ctx context.Context, db *sql.DB) (sql.Result, error) {
	q := "INSERT INTO `user` (`user`.`id`,`user`.`checksum`,`user`.`customer_id`,`user`.`email`,`user`.`name`,`user`.`username`,`user`.`active`,`user`.`autoenrolled`,`user`.`visible`,`user`.`avatar_url`,`user`.`employee_id`,`user`.`location`,`user`.`region`,`user`.`department`,`user`.`reports_to_user_id`,`user`.`title`,`user`.`role`,`user`.`created_at`,`user`.`updated_at`,`user`.`metadata`,`user`.`ref_id`,`user`.`ref_type`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.Email),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.Username),
		orm.ToSQLBool(t.Active),
		orm.ToSQLBool(t.Autoenrolled),
		orm.ToSQLBool(t.Visible),
		orm.ToSQLString(t.AvatarURL),
		orm.ToSQLString(t.EmployeeID),
		orm.ToSQLString(t.Location),
		orm.ToSQLString(t.Region),
		orm.ToSQLString(t.Department),
		orm.ToSQLString(t.ReportsToUserID),
		orm.ToSQLString(t.Title),
		orm.ToSQLString(t.Role),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
		orm.ToSQLString(t.Metadata),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.RefType),
	)
}

// DBCreateTx will create a new User record in the database using the provided transaction
func (t *User) DBCreateTx(ctx context.Context, tx *sql.Tx) (sql.Result, error) {
	q := "INSERT INTO `user` (`user`.`id`,`user`.`checksum`,`user`.`customer_id`,`user`.`email`,`user`.`name`,`user`.`username`,`user`.`active`,`user`.`autoenrolled`,`user`.`visible`,`user`.`avatar_url`,`user`.`employee_id`,`user`.`location`,`user`.`region`,`user`.`department`,`user`.`reports_to_user_id`,`user`.`title`,`user`.`role`,`user`.`created_at`,`user`.`updated_at`,`user`.`metadata`,`user`.`ref_id`,`user`.`ref_type`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.Email),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.Username),
		orm.ToSQLBool(t.Active),
		orm.ToSQLBool(t.Autoenrolled),
		orm.ToSQLBool(t.Visible),
		orm.ToSQLString(t.AvatarURL),
		orm.ToSQLString(t.EmployeeID),
		orm.ToSQLString(t.Location),
		orm.ToSQLString(t.Region),
		orm.ToSQLString(t.Department),
		orm.ToSQLString(t.ReportsToUserID),
		orm.ToSQLString(t.Title),
		orm.ToSQLString(t.Role),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
		orm.ToSQLString(t.Metadata),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.RefType),
	)
}

// DBCreateIgnoreDuplicate will upsert the User record in the database
func (t *User) DBCreateIgnoreDuplicate(ctx context.Context, db *sql.DB) (sql.Result, error) {
	q := "INSERT INTO `user` (`user`.`id`,`user`.`checksum`,`user`.`customer_id`,`user`.`email`,`user`.`name`,`user`.`username`,`user`.`active`,`user`.`autoenrolled`,`user`.`visible`,`user`.`avatar_url`,`user`.`employee_id`,`user`.`location`,`user`.`region`,`user`.`department`,`user`.`reports_to_user_id`,`user`.`title`,`user`.`role`,`user`.`created_at`,`user`.`updated_at`,`user`.`metadata`,`user`.`ref_id`,`user`.`ref_type`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.Email),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.Username),
		orm.ToSQLBool(t.Active),
		orm.ToSQLBool(t.Autoenrolled),
		orm.ToSQLBool(t.Visible),
		orm.ToSQLString(t.AvatarURL),
		orm.ToSQLString(t.EmployeeID),
		orm.ToSQLString(t.Location),
		orm.ToSQLString(t.Region),
		orm.ToSQLString(t.Department),
		orm.ToSQLString(t.ReportsToUserID),
		orm.ToSQLString(t.Title),
		orm.ToSQLString(t.Role),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
		orm.ToSQLString(t.Metadata),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.RefType),
	)
}

// DBCreateIgnoreDuplicateTx will upsert the User record in the database using the provided transaction
func (t *User) DBCreateIgnoreDuplicateTx(ctx context.Context, tx *sql.Tx) (sql.Result, error) {
	q := "INSERT INTO `user` (`user`.`id`,`user`.`checksum`,`user`.`customer_id`,`user`.`email`,`user`.`name`,`user`.`username`,`user`.`active`,`user`.`autoenrolled`,`user`.`visible`,`user`.`avatar_url`,`user`.`employee_id`,`user`.`location`,`user`.`region`,`user`.`department`,`user`.`reports_to_user_id`,`user`.`title`,`user`.`role`,`user`.`created_at`,`user`.`updated_at`,`user`.`metadata`,`user`.`ref_id`,`user`.`ref_type`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.Email),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.Username),
		orm.ToSQLBool(t.Active),
		orm.ToSQLBool(t.Autoenrolled),
		orm.ToSQLBool(t.Visible),
		orm.ToSQLString(t.AvatarURL),
		orm.ToSQLString(t.EmployeeID),
		orm.ToSQLString(t.Location),
		orm.ToSQLString(t.Region),
		orm.ToSQLString(t.Department),
		orm.ToSQLString(t.ReportsToUserID),
		orm.ToSQLString(t.Title),
		orm.ToSQLString(t.Role),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
		orm.ToSQLString(t.Metadata),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.RefType),
	)
}

// DeleteAllUsers deletes all User records in the database with optional filters
func DeleteAllUsers(ctx context.Context, db *sql.DB, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(UserTableName),
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

// DeleteAllUsersTx deletes all User records in the database with optional filters using the provided transaction
func DeleteAllUsersTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(UserTableName),
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

// DBDelete will delete this User record in the database
func (t *User) DBDelete(ctx context.Context, db *sql.DB) (bool, error) {
	q := "DELETE FROM `user` WHERE `id` = ?"
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

// DBDeleteTx will delete this User record in the database using the provided transaction
func (t *User) DBDeleteTx(ctx context.Context, tx *sql.Tx) (bool, error) {
	q := "DELETE FROM `user` WHERE `id` = ?"
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

// DBUpdate will update the User record in the database
func (t *User) DBUpdate(ctx context.Context, db *sql.DB) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `user` SET `checksum`=?,`customer_id`=?,`email`=?,`name`=?,`username`=?,`active`=?,`autoenrolled`=?,`visible`=?,`avatar_url`=?,`employee_id`=?,`location`=?,`region`=?,`department`=?,`reports_to_user_id`=?,`title`=?,`role`=?,`created_at`=?,`updated_at`=?,`metadata`=?,`ref_id`=?,`ref_type`=? WHERE `id`=?"
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.Email),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.Username),
		orm.ToSQLBool(t.Active),
		orm.ToSQLBool(t.Autoenrolled),
		orm.ToSQLBool(t.Visible),
		orm.ToSQLString(t.AvatarURL),
		orm.ToSQLString(t.EmployeeID),
		orm.ToSQLString(t.Location),
		orm.ToSQLString(t.Region),
		orm.ToSQLString(t.Department),
		orm.ToSQLString(t.ReportsToUserID),
		orm.ToSQLString(t.Title),
		orm.ToSQLString(t.Role),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
		orm.ToSQLString(t.Metadata),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.RefType),
		orm.ToSQLString(t.ID),
	)
}

// DBUpdateTx will update the User record in the database using the provided transaction
func (t *User) DBUpdateTx(ctx context.Context, tx *sql.Tx) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `user` SET `checksum`=?,`customer_id`=?,`email`=?,`name`=?,`username`=?,`active`=?,`autoenrolled`=?,`visible`=?,`avatar_url`=?,`employee_id`=?,`location`=?,`region`=?,`department`=?,`reports_to_user_id`=?,`title`=?,`role`=?,`created_at`=?,`updated_at`=?,`metadata`=?,`ref_id`=?,`ref_type`=? WHERE `id`=?"
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.Email),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.Username),
		orm.ToSQLBool(t.Active),
		orm.ToSQLBool(t.Autoenrolled),
		orm.ToSQLBool(t.Visible),
		orm.ToSQLString(t.AvatarURL),
		orm.ToSQLString(t.EmployeeID),
		orm.ToSQLString(t.Location),
		orm.ToSQLString(t.Region),
		orm.ToSQLString(t.Department),
		orm.ToSQLString(t.ReportsToUserID),
		orm.ToSQLString(t.Title),
		orm.ToSQLString(t.Role),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
		orm.ToSQLString(t.Metadata),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.RefType),
		orm.ToSQLString(t.ID),
	)
}

// DBUpsert will upsert the User record in the database
func (t *User) DBUpsert(ctx context.Context, db *sql.DB, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `user` (`user`.`id`,`user`.`checksum`,`user`.`customer_id`,`user`.`email`,`user`.`name`,`user`.`username`,`user`.`active`,`user`.`autoenrolled`,`user`.`visible`,`user`.`avatar_url`,`user`.`employee_id`,`user`.`location`,`user`.`region`,`user`.`department`,`user`.`reports_to_user_id`,`user`.`title`,`user`.`role`,`user`.`created_at`,`user`.`updated_at`,`user`.`metadata`,`user`.`ref_id`,`user`.`ref_type`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `user` (`user`.`id`,`user`.`checksum`,`user`.`customer_id`,`user`.`email`,`user`.`name`,`user`.`username`,`user`.`active`,`user`.`autoenrolled`,`user`.`visible`,`user`.`avatar_url`,`user`.`employee_id`,`user`.`location`,`user`.`region`,`user`.`department`,`user`.`reports_to_user_id`,`user`.`title`,`user`.`role`,`user`.`created_at`,`user`.`updated_at`,`user`.`metadata`,`user`.`ref_id`,`user`.`ref_type`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`customer_id`=VALUES(`customer_id`),`email`=VALUES(`email`),`name`=VALUES(`name`),`username`=VALUES(`username`),`active`=VALUES(`active`),`autoenrolled`=VALUES(`autoenrolled`),`visible`=VALUES(`visible`),`avatar_url`=VALUES(`avatar_url`),`employee_id`=VALUES(`employee_id`),`location`=VALUES(`location`),`region`=VALUES(`region`),`department`=VALUES(`department`),`reports_to_user_id`=VALUES(`reports_to_user_id`),`title`=VALUES(`title`),`role`=VALUES(`role`),`created_at`=VALUES(`created_at`),`updated_at`=VALUES(`updated_at`),`metadata`=VALUES(`metadata`),`ref_id`=VALUES(`ref_id`),`ref_type`=VALUES(`ref_type`)"
	}
	r, err := db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.Email),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.Username),
		orm.ToSQLBool(t.Active),
		orm.ToSQLBool(t.Autoenrolled),
		orm.ToSQLBool(t.Visible),
		orm.ToSQLString(t.AvatarURL),
		orm.ToSQLString(t.EmployeeID),
		orm.ToSQLString(t.Location),
		orm.ToSQLString(t.Region),
		orm.ToSQLString(t.Department),
		orm.ToSQLString(t.ReportsToUserID),
		orm.ToSQLString(t.Title),
		orm.ToSQLString(t.Role),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
		orm.ToSQLString(t.Metadata),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.RefType),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBUpsertTx will upsert the User record in the database using the provided transaction
func (t *User) DBUpsertTx(ctx context.Context, tx *sql.Tx, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `user` (`user`.`id`,`user`.`checksum`,`user`.`customer_id`,`user`.`email`,`user`.`name`,`user`.`username`,`user`.`active`,`user`.`autoenrolled`,`user`.`visible`,`user`.`avatar_url`,`user`.`employee_id`,`user`.`location`,`user`.`region`,`user`.`department`,`user`.`reports_to_user_id`,`user`.`title`,`user`.`role`,`user`.`created_at`,`user`.`updated_at`,`user`.`metadata`,`user`.`ref_id`,`user`.`ref_type`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `user` (`user`.`id`,`user`.`checksum`,`user`.`customer_id`,`user`.`email`,`user`.`name`,`user`.`username`,`user`.`active`,`user`.`autoenrolled`,`user`.`visible`,`user`.`avatar_url`,`user`.`employee_id`,`user`.`location`,`user`.`region`,`user`.`department`,`user`.`reports_to_user_id`,`user`.`title`,`user`.`role`,`user`.`created_at`,`user`.`updated_at`,`user`.`metadata`,`user`.`ref_id`,`user`.`ref_type`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`customer_id`=VALUES(`customer_id`),`email`=VALUES(`email`),`name`=VALUES(`name`),`username`=VALUES(`username`),`active`=VALUES(`active`),`autoenrolled`=VALUES(`autoenrolled`),`visible`=VALUES(`visible`),`avatar_url`=VALUES(`avatar_url`),`employee_id`=VALUES(`employee_id`),`location`=VALUES(`location`),`region`=VALUES(`region`),`department`=VALUES(`department`),`reports_to_user_id`=VALUES(`reports_to_user_id`),`title`=VALUES(`title`),`role`=VALUES(`role`),`created_at`=VALUES(`created_at`),`updated_at`=VALUES(`updated_at`),`metadata`=VALUES(`metadata`),`ref_id`=VALUES(`ref_id`),`ref_type`=VALUES(`ref_type`)"
	}
	r, err := tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.Email),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.Username),
		orm.ToSQLBool(t.Active),
		orm.ToSQLBool(t.Autoenrolled),
		orm.ToSQLBool(t.Visible),
		orm.ToSQLString(t.AvatarURL),
		orm.ToSQLString(t.EmployeeID),
		orm.ToSQLString(t.Location),
		orm.ToSQLString(t.Region),
		orm.ToSQLString(t.Department),
		orm.ToSQLString(t.ReportsToUserID),
		orm.ToSQLString(t.Title),
		orm.ToSQLString(t.Role),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
		orm.ToSQLString(t.Metadata),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.RefType),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBFindOne will find a User record in the database with the primary key
func (t *User) DBFindOne(ctx context.Context, db *sql.DB, value string) (bool, error) {
	q := "SELECT `user`.`id`,`user`.`checksum`,`user`.`customer_id`,`user`.`email`,`user`.`name`,`user`.`username`,`user`.`active`,`user`.`autoenrolled`,`user`.`visible`,`user`.`avatar_url`,`user`.`employee_id`,`user`.`location`,`user`.`region`,`user`.`department`,`user`.`reports_to_user_id`,`user`.`title`,`user`.`role`,`user`.`created_at`,`user`.`updated_at`,`user`.`metadata`,`user`.`ref_id`,`user`.`ref_type` FROM `user` WHERE `id` = ? LIMIT 1"
	row := db.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CustomerID sql.NullString
	var _Email sql.NullString
	var _Name sql.NullString
	var _Username sql.NullString
	var _Active sql.NullBool
	var _Autoenrolled sql.NullBool
	var _Visible sql.NullBool
	var _AvatarURL sql.NullString
	var _EmployeeID sql.NullString
	var _Location sql.NullString
	var _Region sql.NullString
	var _Department sql.NullString
	var _ReportsToUserID sql.NullString
	var _Title sql.NullString
	var _Role sql.NullString
	var _CreatedAt sql.NullInt64
	var _UpdatedAt sql.NullInt64
	var _Metadata sql.NullString
	var _RefID sql.NullString
	var _RefType sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_Email,
		&_Name,
		&_Username,
		&_Active,
		&_Autoenrolled,
		&_Visible,
		&_AvatarURL,
		&_EmployeeID,
		&_Location,
		&_Region,
		&_Department,
		&_ReportsToUserID,
		&_Title,
		&_Role,
		&_CreatedAt,
		&_UpdatedAt,
		&_Metadata,
		&_RefID,
		&_RefType,
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
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _Email.Valid {
		t.SetEmail(_Email.String)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _Username.Valid {
		t.SetUsername(_Username.String)
	}
	if _Active.Valid {
		t.SetActive(_Active.Bool)
	}
	if _Autoenrolled.Valid {
		t.SetAutoenrolled(_Autoenrolled.Bool)
	}
	if _Visible.Valid {
		t.SetVisible(_Visible.Bool)
	}
	if _AvatarURL.Valid {
		t.SetAvatarURL(_AvatarURL.String)
	}
	if _EmployeeID.Valid {
		t.SetEmployeeID(_EmployeeID.String)
	}
	if _Location.Valid {
		t.SetLocation(_Location.String)
	}
	if _Region.Valid {
		t.SetRegion(_Region.String)
	}
	if _Department.Valid {
		t.SetDepartment(_Department.String)
	}
	if _ReportsToUserID.Valid {
		t.SetReportsToUserID(_ReportsToUserID.String)
	}
	if _Title.Valid {
		t.SetTitle(_Title.String)
	}
	if _Role.Valid {
		t.SetRole(_Role.String)
	}
	if _CreatedAt.Valid {
		t.SetCreatedAt(_CreatedAt.Int64)
	}
	if _UpdatedAt.Valid {
		t.SetUpdatedAt(_UpdatedAt.Int64)
	}
	if _Metadata.Valid {
		t.SetMetadata(_Metadata.String)
	}
	if _RefID.Valid {
		t.SetRefID(_RefID.String)
	}
	if _RefType.Valid {
		t.SetRefType(_RefType.String)
	}
	return true, nil
}

// DBFindOneTx will find a User record in the database with the primary key using the provided transaction
func (t *User) DBFindOneTx(ctx context.Context, tx *sql.Tx, value string) (bool, error) {
	q := "SELECT `user`.`id`,`user`.`checksum`,`user`.`customer_id`,`user`.`email`,`user`.`name`,`user`.`username`,`user`.`active`,`user`.`autoenrolled`,`user`.`visible`,`user`.`avatar_url`,`user`.`employee_id`,`user`.`location`,`user`.`region`,`user`.`department`,`user`.`reports_to_user_id`,`user`.`title`,`user`.`role`,`user`.`created_at`,`user`.`updated_at`,`user`.`metadata`,`user`.`ref_id`,`user`.`ref_type` FROM `user` WHERE `id` = ? LIMIT 1"
	row := tx.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CustomerID sql.NullString
	var _Email sql.NullString
	var _Name sql.NullString
	var _Username sql.NullString
	var _Active sql.NullBool
	var _Autoenrolled sql.NullBool
	var _Visible sql.NullBool
	var _AvatarURL sql.NullString
	var _EmployeeID sql.NullString
	var _Location sql.NullString
	var _Region sql.NullString
	var _Department sql.NullString
	var _ReportsToUserID sql.NullString
	var _Title sql.NullString
	var _Role sql.NullString
	var _CreatedAt sql.NullInt64
	var _UpdatedAt sql.NullInt64
	var _Metadata sql.NullString
	var _RefID sql.NullString
	var _RefType sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_Email,
		&_Name,
		&_Username,
		&_Active,
		&_Autoenrolled,
		&_Visible,
		&_AvatarURL,
		&_EmployeeID,
		&_Location,
		&_Region,
		&_Department,
		&_ReportsToUserID,
		&_Title,
		&_Role,
		&_CreatedAt,
		&_UpdatedAt,
		&_Metadata,
		&_RefID,
		&_RefType,
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
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _Email.Valid {
		t.SetEmail(_Email.String)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _Username.Valid {
		t.SetUsername(_Username.String)
	}
	if _Active.Valid {
		t.SetActive(_Active.Bool)
	}
	if _Autoenrolled.Valid {
		t.SetAutoenrolled(_Autoenrolled.Bool)
	}
	if _Visible.Valid {
		t.SetVisible(_Visible.Bool)
	}
	if _AvatarURL.Valid {
		t.SetAvatarURL(_AvatarURL.String)
	}
	if _EmployeeID.Valid {
		t.SetEmployeeID(_EmployeeID.String)
	}
	if _Location.Valid {
		t.SetLocation(_Location.String)
	}
	if _Region.Valid {
		t.SetRegion(_Region.String)
	}
	if _Department.Valid {
		t.SetDepartment(_Department.String)
	}
	if _ReportsToUserID.Valid {
		t.SetReportsToUserID(_ReportsToUserID.String)
	}
	if _Title.Valid {
		t.SetTitle(_Title.String)
	}
	if _Role.Valid {
		t.SetRole(_Role.String)
	}
	if _CreatedAt.Valid {
		t.SetCreatedAt(_CreatedAt.Int64)
	}
	if _UpdatedAt.Valid {
		t.SetUpdatedAt(_UpdatedAt.Int64)
	}
	if _Metadata.Valid {
		t.SetMetadata(_Metadata.String)
	}
	if _RefID.Valid {
		t.SetRefID(_RefID.String)
	}
	if _RefType.Valid {
		t.SetRefType(_RefType.String)
	}
	return true, nil
}

// FindUsers will find a User record in the database with the provided parameters
func FindUsers(ctx context.Context, db *sql.DB, _params ...interface{}) ([]*User, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("customer_id"),
		orm.Column("email"),
		orm.Column("name"),
		orm.Column("username"),
		orm.Column("active"),
		orm.Column("autoenrolled"),
		orm.Column("visible"),
		orm.Column("avatar_url"),
		orm.Column("employee_id"),
		orm.Column("location"),
		orm.Column("region"),
		orm.Column("department"),
		orm.Column("reports_to_user_id"),
		orm.Column("title"),
		orm.Column("role"),
		orm.Column("created_at"),
		orm.Column("updated_at"),
		orm.Column("metadata"),
		orm.Column("ref_id"),
		orm.Column("ref_type"),
		orm.Table(UserTableName),
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
	results := make([]*User, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _Email sql.NullString
		var _Name sql.NullString
		var _Username sql.NullString
		var _Active sql.NullBool
		var _Autoenrolled sql.NullBool
		var _Visible sql.NullBool
		var _AvatarURL sql.NullString
		var _EmployeeID sql.NullString
		var _Location sql.NullString
		var _Region sql.NullString
		var _Department sql.NullString
		var _ReportsToUserID sql.NullString
		var _Title sql.NullString
		var _Role sql.NullString
		var _CreatedAt sql.NullInt64
		var _UpdatedAt sql.NullInt64
		var _Metadata sql.NullString
		var _RefID sql.NullString
		var _RefType sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_Email,
			&_Name,
			&_Username,
			&_Active,
			&_Autoenrolled,
			&_Visible,
			&_AvatarURL,
			&_EmployeeID,
			&_Location,
			&_Region,
			&_Department,
			&_ReportsToUserID,
			&_Title,
			&_Role,
			&_CreatedAt,
			&_UpdatedAt,
			&_Metadata,
			&_RefID,
			&_RefType,
		)
		if err != nil {
			return nil, err
		}
		t := &User{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _Email.Valid {
			t.SetEmail(_Email.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _Username.Valid {
			t.SetUsername(_Username.String)
		}
		if _Active.Valid {
			t.SetActive(_Active.Bool)
		}
		if _Autoenrolled.Valid {
			t.SetAutoenrolled(_Autoenrolled.Bool)
		}
		if _Visible.Valid {
			t.SetVisible(_Visible.Bool)
		}
		if _AvatarURL.Valid {
			t.SetAvatarURL(_AvatarURL.String)
		}
		if _EmployeeID.Valid {
			t.SetEmployeeID(_EmployeeID.String)
		}
		if _Location.Valid {
			t.SetLocation(_Location.String)
		}
		if _Region.Valid {
			t.SetRegion(_Region.String)
		}
		if _Department.Valid {
			t.SetDepartment(_Department.String)
		}
		if _ReportsToUserID.Valid {
			t.SetReportsToUserID(_ReportsToUserID.String)
		}
		if _Title.Valid {
			t.SetTitle(_Title.String)
		}
		if _Role.Valid {
			t.SetRole(_Role.String)
		}
		if _CreatedAt.Valid {
			t.SetCreatedAt(_CreatedAt.Int64)
		}
		if _UpdatedAt.Valid {
			t.SetUpdatedAt(_UpdatedAt.Int64)
		}
		if _Metadata.Valid {
			t.SetMetadata(_Metadata.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		if _RefType.Valid {
			t.SetRefType(_RefType.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindUsersTx will find a User record in the database with the provided parameters using the provided transaction
func FindUsersTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) ([]*User, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("customer_id"),
		orm.Column("email"),
		orm.Column("name"),
		orm.Column("username"),
		orm.Column("active"),
		orm.Column("autoenrolled"),
		orm.Column("visible"),
		orm.Column("avatar_url"),
		orm.Column("employee_id"),
		orm.Column("location"),
		orm.Column("region"),
		orm.Column("department"),
		orm.Column("reports_to_user_id"),
		orm.Column("title"),
		orm.Column("role"),
		orm.Column("created_at"),
		orm.Column("updated_at"),
		orm.Column("metadata"),
		orm.Column("ref_id"),
		orm.Column("ref_type"),
		orm.Table(UserTableName),
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
	results := make([]*User, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _Email sql.NullString
		var _Name sql.NullString
		var _Username sql.NullString
		var _Active sql.NullBool
		var _Autoenrolled sql.NullBool
		var _Visible sql.NullBool
		var _AvatarURL sql.NullString
		var _EmployeeID sql.NullString
		var _Location sql.NullString
		var _Region sql.NullString
		var _Department sql.NullString
		var _ReportsToUserID sql.NullString
		var _Title sql.NullString
		var _Role sql.NullString
		var _CreatedAt sql.NullInt64
		var _UpdatedAt sql.NullInt64
		var _Metadata sql.NullString
		var _RefID sql.NullString
		var _RefType sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_Email,
			&_Name,
			&_Username,
			&_Active,
			&_Autoenrolled,
			&_Visible,
			&_AvatarURL,
			&_EmployeeID,
			&_Location,
			&_Region,
			&_Department,
			&_ReportsToUserID,
			&_Title,
			&_Role,
			&_CreatedAt,
			&_UpdatedAt,
			&_Metadata,
			&_RefID,
			&_RefType,
		)
		if err != nil {
			return nil, err
		}
		t := &User{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _Email.Valid {
			t.SetEmail(_Email.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _Username.Valid {
			t.SetUsername(_Username.String)
		}
		if _Active.Valid {
			t.SetActive(_Active.Bool)
		}
		if _Autoenrolled.Valid {
			t.SetAutoenrolled(_Autoenrolled.Bool)
		}
		if _Visible.Valid {
			t.SetVisible(_Visible.Bool)
		}
		if _AvatarURL.Valid {
			t.SetAvatarURL(_AvatarURL.String)
		}
		if _EmployeeID.Valid {
			t.SetEmployeeID(_EmployeeID.String)
		}
		if _Location.Valid {
			t.SetLocation(_Location.String)
		}
		if _Region.Valid {
			t.SetRegion(_Region.String)
		}
		if _Department.Valid {
			t.SetDepartment(_Department.String)
		}
		if _ReportsToUserID.Valid {
			t.SetReportsToUserID(_ReportsToUserID.String)
		}
		if _Title.Valid {
			t.SetTitle(_Title.String)
		}
		if _Role.Valid {
			t.SetRole(_Role.String)
		}
		if _CreatedAt.Valid {
			t.SetCreatedAt(_CreatedAt.Int64)
		}
		if _UpdatedAt.Valid {
			t.SetUpdatedAt(_UpdatedAt.Int64)
		}
		if _Metadata.Valid {
			t.SetMetadata(_Metadata.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		if _RefType.Valid {
			t.SetRefType(_RefType.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// DBFind will find a User record in the database with the provided parameters
func (t *User) DBFind(ctx context.Context, db *sql.DB, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("customer_id"),
		orm.Column("email"),
		orm.Column("name"),
		orm.Column("username"),
		orm.Column("active"),
		orm.Column("autoenrolled"),
		orm.Column("visible"),
		orm.Column("avatar_url"),
		orm.Column("employee_id"),
		orm.Column("location"),
		orm.Column("region"),
		orm.Column("department"),
		orm.Column("reports_to_user_id"),
		orm.Column("title"),
		orm.Column("role"),
		orm.Column("created_at"),
		orm.Column("updated_at"),
		orm.Column("metadata"),
		orm.Column("ref_id"),
		orm.Column("ref_type"),
		orm.Table(UserTableName),
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
	var _CustomerID sql.NullString
	var _Email sql.NullString
	var _Name sql.NullString
	var _Username sql.NullString
	var _Active sql.NullBool
	var _Autoenrolled sql.NullBool
	var _Visible sql.NullBool
	var _AvatarURL sql.NullString
	var _EmployeeID sql.NullString
	var _Location sql.NullString
	var _Region sql.NullString
	var _Department sql.NullString
	var _ReportsToUserID sql.NullString
	var _Title sql.NullString
	var _Role sql.NullString
	var _CreatedAt sql.NullInt64
	var _UpdatedAt sql.NullInt64
	var _Metadata sql.NullString
	var _RefID sql.NullString
	var _RefType sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_Email,
		&_Name,
		&_Username,
		&_Active,
		&_Autoenrolled,
		&_Visible,
		&_AvatarURL,
		&_EmployeeID,
		&_Location,
		&_Region,
		&_Department,
		&_ReportsToUserID,
		&_Title,
		&_Role,
		&_CreatedAt,
		&_UpdatedAt,
		&_Metadata,
		&_RefID,
		&_RefType,
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
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _Email.Valid {
		t.SetEmail(_Email.String)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _Username.Valid {
		t.SetUsername(_Username.String)
	}
	if _Active.Valid {
		t.SetActive(_Active.Bool)
	}
	if _Autoenrolled.Valid {
		t.SetAutoenrolled(_Autoenrolled.Bool)
	}
	if _Visible.Valid {
		t.SetVisible(_Visible.Bool)
	}
	if _AvatarURL.Valid {
		t.SetAvatarURL(_AvatarURL.String)
	}
	if _EmployeeID.Valid {
		t.SetEmployeeID(_EmployeeID.String)
	}
	if _Location.Valid {
		t.SetLocation(_Location.String)
	}
	if _Region.Valid {
		t.SetRegion(_Region.String)
	}
	if _Department.Valid {
		t.SetDepartment(_Department.String)
	}
	if _ReportsToUserID.Valid {
		t.SetReportsToUserID(_ReportsToUserID.String)
	}
	if _Title.Valid {
		t.SetTitle(_Title.String)
	}
	if _Role.Valid {
		t.SetRole(_Role.String)
	}
	if _CreatedAt.Valid {
		t.SetCreatedAt(_CreatedAt.Int64)
	}
	if _UpdatedAt.Valid {
		t.SetUpdatedAt(_UpdatedAt.Int64)
	}
	if _Metadata.Valid {
		t.SetMetadata(_Metadata.String)
	}
	if _RefID.Valid {
		t.SetRefID(_RefID.String)
	}
	if _RefType.Valid {
		t.SetRefType(_RefType.String)
	}
	return true, nil
}

// DBFindTx will find a User record in the database with the provided parameters using the provided transaction
func (t *User) DBFindTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("customer_id"),
		orm.Column("email"),
		orm.Column("name"),
		orm.Column("username"),
		orm.Column("active"),
		orm.Column("autoenrolled"),
		orm.Column("visible"),
		orm.Column("avatar_url"),
		orm.Column("employee_id"),
		orm.Column("location"),
		orm.Column("region"),
		orm.Column("department"),
		orm.Column("reports_to_user_id"),
		orm.Column("title"),
		orm.Column("role"),
		orm.Column("created_at"),
		orm.Column("updated_at"),
		orm.Column("metadata"),
		orm.Column("ref_id"),
		orm.Column("ref_type"),
		orm.Table(UserTableName),
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
	var _CustomerID sql.NullString
	var _Email sql.NullString
	var _Name sql.NullString
	var _Username sql.NullString
	var _Active sql.NullBool
	var _Autoenrolled sql.NullBool
	var _Visible sql.NullBool
	var _AvatarURL sql.NullString
	var _EmployeeID sql.NullString
	var _Location sql.NullString
	var _Region sql.NullString
	var _Department sql.NullString
	var _ReportsToUserID sql.NullString
	var _Title sql.NullString
	var _Role sql.NullString
	var _CreatedAt sql.NullInt64
	var _UpdatedAt sql.NullInt64
	var _Metadata sql.NullString
	var _RefID sql.NullString
	var _RefType sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_Email,
		&_Name,
		&_Username,
		&_Active,
		&_Autoenrolled,
		&_Visible,
		&_AvatarURL,
		&_EmployeeID,
		&_Location,
		&_Region,
		&_Department,
		&_ReportsToUserID,
		&_Title,
		&_Role,
		&_CreatedAt,
		&_UpdatedAt,
		&_Metadata,
		&_RefID,
		&_RefType,
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
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _Email.Valid {
		t.SetEmail(_Email.String)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _Username.Valid {
		t.SetUsername(_Username.String)
	}
	if _Active.Valid {
		t.SetActive(_Active.Bool)
	}
	if _Autoenrolled.Valid {
		t.SetAutoenrolled(_Autoenrolled.Bool)
	}
	if _Visible.Valid {
		t.SetVisible(_Visible.Bool)
	}
	if _AvatarURL.Valid {
		t.SetAvatarURL(_AvatarURL.String)
	}
	if _EmployeeID.Valid {
		t.SetEmployeeID(_EmployeeID.String)
	}
	if _Location.Valid {
		t.SetLocation(_Location.String)
	}
	if _Region.Valid {
		t.SetRegion(_Region.String)
	}
	if _Department.Valid {
		t.SetDepartment(_Department.String)
	}
	if _ReportsToUserID.Valid {
		t.SetReportsToUserID(_ReportsToUserID.String)
	}
	if _Title.Valid {
		t.SetTitle(_Title.String)
	}
	if _Role.Valid {
		t.SetRole(_Role.String)
	}
	if _CreatedAt.Valid {
		t.SetCreatedAt(_CreatedAt.Int64)
	}
	if _UpdatedAt.Valid {
		t.SetUpdatedAt(_UpdatedAt.Int64)
	}
	if _Metadata.Valid {
		t.SetMetadata(_Metadata.String)
	}
	if _RefID.Valid {
		t.SetRefID(_RefID.String)
	}
	if _RefType.Valid {
		t.SetRefType(_RefType.String)
	}
	return true, nil
}

// CountUsers will find the count of User records in the database
func CountUsers(ctx context.Context, db *sql.DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(UserTableName),
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

// CountUsersTx will find the count of User records in the database using the provided transaction
func CountUsersTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(UserTableName),
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

// DBCount will find the count of User records in the database
func (t *User) DBCount(ctx context.Context, db *sql.DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(UserTableName),
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

// DBCountTx will find the count of User records in the database using the provided transaction
func (t *User) DBCountTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(UserTableName),
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

// DBExists will return true if the User record exists in the database
func (t *User) DBExists(ctx context.Context, db *sql.DB) (bool, error) {
	q := "SELECT `id` FROM `user` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := db.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// DBExistsTx will return true if the User record exists in the database using the provided transaction
func (t *User) DBExistsTx(ctx context.Context, tx *sql.Tx) (bool, error) {
	q := "SELECT `id` FROM `user` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := tx.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// PrimaryKeyColumn returns the column name for the primary key
func (t *User) PrimaryKeyColumn() string {
	return UserColumnID
}

// PrimaryKeyColumnType returns the primary key column Go type as a string
func (t *User) PrimaryKeyColumnType() string {
	return "string"
}

// PrimaryKey returns the primary key column value
func (t *User) PrimaryKey() interface{} {
	return t.ID
}
