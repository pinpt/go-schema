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

var _ Model = (*UserCostCenter)(nil)
var _ CSVWriter = (*UserCostCenter)(nil)
var _ JSONWriter = (*UserCostCenter)(nil)
var _ Checksum = (*UserCostCenter)(nil)

// UserCostCenterTableName is the name of the table in SQL
const UserCostCenterTableName = "user_cost_center"

var UserCostCenterColumns = []string{
	"id",
	"checksum",
	"customer_id",
	"user_id",
	"cost_center_id",
	"unit_cost",
	"unit_type",
	"unit_currency",
	"allocation",
	"created_at",
	"updated_at",
}

type UserCostCenter_UserUnitCostType string

const (
	UserUnitCostType_SALARY     UserCostCenter_UserUnitCostType = "salary"
	UserUnitCostType_HOURLY     UserCostCenter_UserUnitCostType = "hourly"
	UserUnitCostType_CONTRACTOR UserCostCenter_UserUnitCostType = "contractor"
	UserUnitCostType_CASUAL     UserCostCenter_UserUnitCostType = "casual"
	UserUnitCostType_OTHER      UserCostCenter_UserUnitCostType = "other"
)

func (x UserCostCenter_UserUnitCostType) String() string {
	return string(x)
}

func enumUserUnitCostTypeToString(v *UserCostCenter_UserUnitCostType) string {
	if v == nil {
		return ""
	}
	return v.String()
}

func toUserUnitCostType(v string) *UserCostCenter_UserUnitCostType {
	var ev *UserCostCenter_UserUnitCostType
	switch v {
	case "SALARY", "salary":
		{
			v := UserUnitCostType_SALARY
			ev = &v
		}
	case "HOURLY", "hourly":
		{
			v := UserUnitCostType_HOURLY
			ev = &v
		}
	case "CONTRACTOR", "contractor":
		{
			v := UserUnitCostType_CONTRACTOR
			ev = &v
		}
	case "CASUAL", "casual":
		{
			v := UserUnitCostType_CASUAL
			ev = &v
		}
	case "OTHER", "other":
		{
			v := UserUnitCostType_OTHER
			ev = &v
		}
	}
	return ev
}

// UserCostCenter table
type UserCostCenter struct {
	Allocation   *float64                         `json:"allocation,omitempty"`
	Checksum     *string                          `json:"checksum,omitempty"`
	CostCenterID *string                          `json:"cost_center_id,omitempty"`
	CreatedAt    int64                            `json:"created_at"`
	CustomerID   string                           `json:"customer_id"`
	ID           string                           `json:"id"`
	UnitCost     *float64                         `json:"unit_cost,omitempty"`
	UnitCurrency *string                          `json:"unit_currency,omitempty"`
	UnitType     *UserCostCenter_UserUnitCostType `json:"unit_type,omitempty"`
	UpdatedAt    *int64                           `json:"updated_at,omitempty"`
	UserID       string                           `json:"user_id"`
}

// TableName returns the SQL table name for UserCostCenter and satifies the Model interface
func (t *UserCostCenter) TableName() string {
	return UserCostCenterTableName
}

// ToCSV will serialize the UserCostCenter instance to a CSV compatible array of strings
func (t *UserCostCenter) ToCSV() []string {
	return []string{
		t.ID,
		t.CalculateChecksum(),
		t.CustomerID,
		t.UserID,
		toCSVString(t.CostCenterID),
		toCSVString(t.UnitCost),
		toCSVString(t.UnitType),
		toCSVString(t.UnitCurrency),
		toCSVString(t.Allocation),
		toCSVString(t.CreatedAt),
		toCSVString(t.UpdatedAt),
	}
}

// WriteCSV will serialize the UserCostCenter instance to the writer as CSV and satisfies the CSVWriter interface
func (t *UserCostCenter) WriteCSV(w *csv.Writer) error {
	return w.Write(t.ToCSV())
}

// WriteJSON will serialize the UserCostCenter instance to the writer as JSON and satisfies the JSONWriter interface
func (t *UserCostCenter) WriteJSON(w io.Writer, indent ...bool) error {
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

// NewUserCostCenterReader creates a JSON reader which can read in UserCostCenter objects serialized as JSON either as an array, single object or json new lines
// and writes each UserCostCenter to the channel provided
func NewUserCostCenterReader(r io.Reader, ch chan<- UserCostCenter) error {
	return orm.Deserialize(r, func(buf json.RawMessage) error {
		dec := json.NewDecoder(bytes.NewBuffer(buf))
		e := UserCostCenter{}
		if err := dec.Decode(&e); err != nil {
			return err
		}
		ch <- e
		return nil
	})
}

// NewCSVUserCostCenterReaderDir will read the reader as CSV and emit each record to the channel provided
func NewCSVUserCostCenterReader(r io.Reader, ch chan<- UserCostCenter) error {
	cr := csv.NewReader(r)
	for {
		record, err := cr.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		ch <- UserCostCenter{
			ID:           record[0],
			Checksum:     fromStringPointer(record[1]),
			CustomerID:   record[2],
			UserID:       record[3],
			CostCenterID: fromStringPointer(record[4]),
			UnitCost:     fromCSVFloat64Pointer(record[5]),
			UnitType:     toUserUnitCostType(record[6]),
			UnitCurrency: fromStringPointer(record[7]),
			Allocation:   fromCSVFloat64Pointer(record[8]),
			CreatedAt:    fromCSVInt64(record[9]),
			UpdatedAt:    fromCSVInt64Pointer(record[10]),
		}
	}
	return nil
}

// NewCSVUserCostCenterReaderFile will read the file as a CSV and emit each record to the channel provided
func NewCSVUserCostCenterReaderFile(fp string, ch chan<- UserCostCenter) error {
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
	return NewCSVUserCostCenterReader(fc, ch)
}

// NewCSVUserCostCenterReaderDir will read the user_cost_center.csv.gz file as a CSV and emit each record to the channel provided
func NewCSVUserCostCenterReaderDir(dir string, ch chan<- UserCostCenter) error {
	return NewCSVUserCostCenterReaderFile(filepath.Join(dir, "user_cost_center.csv.gz"), ch)
}

// UserCostCenterCSVDeduper is a function callback which takes the existing value (a) and the new value (b)
// and the return value should be the one to use (or a new one, if applicable). return nil
// to skip processing of this record
type UserCostCenterCSVDeduper func(a UserCostCenter, b UserCostCenter) *UserCostCenter

// UserCostCenterCSVDedupeDisabled is set on whether the CSV writer should de-dupe values by key
var UserCostCenterCSVDedupeDisabled bool

// NewUserCostCenterCSVWriterSize creates a batch writer that will write each UserCostCenter into a CSV file
// this method will automatically de-duplicate entries using the primary key. if the checksum
// for a newer item with the same primary key doesn't match a previously sent item, the newer
// one will be used
func NewUserCostCenterCSVWriterSize(w io.Writer, size int, dedupers ...UserCostCenterCSVDeduper) (chan UserCostCenter, chan bool, error) {
	cw := csv.NewWriter(w)
	ch := make(chan UserCostCenter, size)
	done := make(chan bool)
	go func() {
		defer func() { done <- true }()
		dodedupe := !UserCostCenterCSVDedupeDisabled
		var kv map[string]*UserCostCenter
		var deduper UserCostCenterCSVDeduper
		if dedupers != nil && len(dedupers) > 0 {
			deduper = dedupers[0]
			dodedupe = true
		}
		if dodedupe {
			kv = make(map[string]*UserCostCenter)
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

// UserCostCenterCSVDefaultSize is the default channel buffer size if not provided
var UserCostCenterCSVDefaultSize = 100

// NewUserCostCenterCSVWriter creates a batch writer that will write each UserCostCenter into a CSV file
func NewUserCostCenterCSVWriter(w io.Writer, dedupers ...UserCostCenterCSVDeduper) (chan UserCostCenter, chan bool, error) {
	return NewUserCostCenterCSVWriterSize(w, UserCostCenterCSVDefaultSize, dedupers...)
}

// NewUserCostCenterCSVWriterDir creates a batch writer that will write each UserCostCenter into a CSV file named user_cost_center.csv.gz in dir
func NewUserCostCenterCSVWriterDir(dir string, dedupers ...UserCostCenterCSVDeduper) (chan UserCostCenter, chan bool, error) {
	return NewUserCostCenterCSVWriterFile(filepath.Join(dir, "user_cost_center.csv.gz"), dedupers...)
}

// NewUserCostCenterCSVWriterFile creates a batch writer that will write each UserCostCenter into a CSV file
func NewUserCostCenterCSVWriterFile(fn string, dedupers ...UserCostCenterCSVDeduper) (chan UserCostCenter, chan bool, error) {
	f, err := os.Create(fn)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening CSV file %s. %v", fn, err)
	}
	var fc io.WriteCloser = f
	if filepath.Ext(fn) == ".gz" {
		w, _ := gzip.NewWriterLevel(f, gzip.BestCompression)
		fc = w
	}
	ch, done, err := NewUserCostCenterCSVWriter(fc, dedupers...)
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

type UserCostCenterDBAction func(ctx context.Context, db DB, record UserCostCenter) error

// NewUserCostCenterDBWriterSize creates a DB writer that will write each issue into the DB
func NewUserCostCenterDBWriterSize(ctx context.Context, db DB, errors chan<- error, size int, actions ...UserCostCenterDBAction) (chan UserCostCenter, chan bool, error) {
	ch := make(chan UserCostCenter, size)
	done := make(chan bool)
	var action UserCostCenterDBAction
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

// NewUserCostCenterDBWriter creates a DB writer that will write each issue into the DB
func NewUserCostCenterDBWriter(ctx context.Context, db DB, errors chan<- error, actions ...UserCostCenterDBAction) (chan UserCostCenter, chan bool, error) {
	return NewUserCostCenterDBWriterSize(ctx, db, errors, 100, actions...)
}

// UserCostCenterColumnID is the ID SQL column name for the UserCostCenter table
const UserCostCenterColumnID = "id"

// UserCostCenterEscapedColumnID is the escaped ID SQL column name for the UserCostCenter table
const UserCostCenterEscapedColumnID = "`id`"

// UserCostCenterColumnChecksum is the Checksum SQL column name for the UserCostCenter table
const UserCostCenterColumnChecksum = "checksum"

// UserCostCenterEscapedColumnChecksum is the escaped Checksum SQL column name for the UserCostCenter table
const UserCostCenterEscapedColumnChecksum = "`checksum`"

// UserCostCenterColumnCustomerID is the CustomerID SQL column name for the UserCostCenter table
const UserCostCenterColumnCustomerID = "customer_id"

// UserCostCenterEscapedColumnCustomerID is the escaped CustomerID SQL column name for the UserCostCenter table
const UserCostCenterEscapedColumnCustomerID = "`customer_id`"

// UserCostCenterColumnUserID is the UserID SQL column name for the UserCostCenter table
const UserCostCenterColumnUserID = "user_id"

// UserCostCenterEscapedColumnUserID is the escaped UserID SQL column name for the UserCostCenter table
const UserCostCenterEscapedColumnUserID = "`user_id`"

// UserCostCenterColumnCostCenterID is the CostCenterID SQL column name for the UserCostCenter table
const UserCostCenterColumnCostCenterID = "cost_center_id"

// UserCostCenterEscapedColumnCostCenterID is the escaped CostCenterID SQL column name for the UserCostCenter table
const UserCostCenterEscapedColumnCostCenterID = "`cost_center_id`"

// UserCostCenterColumnUnitCost is the UnitCost SQL column name for the UserCostCenter table
const UserCostCenterColumnUnitCost = "unit_cost"

// UserCostCenterEscapedColumnUnitCost is the escaped UnitCost SQL column name for the UserCostCenter table
const UserCostCenterEscapedColumnUnitCost = "`unit_cost`"

// UserCostCenterColumnUnitType is the UnitType SQL column name for the UserCostCenter table
const UserCostCenterColumnUnitType = "unit_type"

// UserCostCenterEscapedColumnUnitType is the escaped UnitType SQL column name for the UserCostCenter table
const UserCostCenterEscapedColumnUnitType = "`unit_type`"

// UserCostCenterColumnUnitCurrency is the UnitCurrency SQL column name for the UserCostCenter table
const UserCostCenterColumnUnitCurrency = "unit_currency"

// UserCostCenterEscapedColumnUnitCurrency is the escaped UnitCurrency SQL column name for the UserCostCenter table
const UserCostCenterEscapedColumnUnitCurrency = "`unit_currency`"

// UserCostCenterColumnAllocation is the Allocation SQL column name for the UserCostCenter table
const UserCostCenterColumnAllocation = "allocation"

// UserCostCenterEscapedColumnAllocation is the escaped Allocation SQL column name for the UserCostCenter table
const UserCostCenterEscapedColumnAllocation = "`allocation`"

// UserCostCenterColumnCreatedAt is the CreatedAt SQL column name for the UserCostCenter table
const UserCostCenterColumnCreatedAt = "created_at"

// UserCostCenterEscapedColumnCreatedAt is the escaped CreatedAt SQL column name for the UserCostCenter table
const UserCostCenterEscapedColumnCreatedAt = "`created_at`"

// UserCostCenterColumnUpdatedAt is the UpdatedAt SQL column name for the UserCostCenter table
const UserCostCenterColumnUpdatedAt = "updated_at"

// UserCostCenterEscapedColumnUpdatedAt is the escaped UpdatedAt SQL column name for the UserCostCenter table
const UserCostCenterEscapedColumnUpdatedAt = "`updated_at`"

// GetID will return the UserCostCenter ID value
func (t *UserCostCenter) GetID() string {
	return t.ID
}

// SetID will set the UserCostCenter ID value
func (t *UserCostCenter) SetID(v string) {
	t.ID = v
}

// FindUserCostCenterByID will find a UserCostCenter by ID
func FindUserCostCenterByID(ctx context.Context, db DB, value string) (*UserCostCenter, error) {
	q := "SELECT `user_cost_center`.`id`,`user_cost_center`.`checksum`,`user_cost_center`.`customer_id`,`user_cost_center`.`user_id`,`user_cost_center`.`cost_center_id`,`user_cost_center`.`unit_cost`,`user_cost_center`.`unit_type`,`user_cost_center`.`unit_currency`,`user_cost_center`.`allocation`,`user_cost_center`.`created_at`,`user_cost_center`.`updated_at` FROM `user_cost_center` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CustomerID sql.NullString
	var _UserID sql.NullString
	var _CostCenterID sql.NullString
	var _UnitCost sql.NullFloat64
	var _UnitType sql.NullString
	var _UnitCurrency sql.NullString
	var _Allocation sql.NullFloat64
	var _CreatedAt sql.NullInt64
	var _UpdatedAt sql.NullInt64
	err := db.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_UserID,
		&_CostCenterID,
		&_UnitCost,
		&_UnitType,
		&_UnitCurrency,
		&_Allocation,
		&_CreatedAt,
		&_UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &UserCostCenter{}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _Checksum.Valid {
		t.SetChecksum(_Checksum.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _UserID.Valid {
		t.SetUserID(_UserID.String)
	}
	if _CostCenterID.Valid {
		t.SetCostCenterID(_CostCenterID.String)
	}
	if _UnitCost.Valid {
		t.SetUnitCost(_UnitCost.Float64)
	}
	if _UnitType.Valid {
		t.SetUnitTypeString(_UnitType.String)
	}
	if _UnitCurrency.Valid {
		t.SetUnitCurrency(_UnitCurrency.String)
	}
	if _Allocation.Valid {
		t.SetAllocation(_Allocation.Float64)
	}
	if _CreatedAt.Valid {
		t.SetCreatedAt(_CreatedAt.Int64)
	}
	if _UpdatedAt.Valid {
		t.SetUpdatedAt(_UpdatedAt.Int64)
	}
	return t, nil
}

// FindUserCostCenterByIDTx will find a UserCostCenter by ID using the provided transaction
func FindUserCostCenterByIDTx(ctx context.Context, tx Tx, value string) (*UserCostCenter, error) {
	q := "SELECT `user_cost_center`.`id`,`user_cost_center`.`checksum`,`user_cost_center`.`customer_id`,`user_cost_center`.`user_id`,`user_cost_center`.`cost_center_id`,`user_cost_center`.`unit_cost`,`user_cost_center`.`unit_type`,`user_cost_center`.`unit_currency`,`user_cost_center`.`allocation`,`user_cost_center`.`created_at`,`user_cost_center`.`updated_at` FROM `user_cost_center` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CustomerID sql.NullString
	var _UserID sql.NullString
	var _CostCenterID sql.NullString
	var _UnitCost sql.NullFloat64
	var _UnitType sql.NullString
	var _UnitCurrency sql.NullString
	var _Allocation sql.NullFloat64
	var _CreatedAt sql.NullInt64
	var _UpdatedAt sql.NullInt64
	err := tx.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_UserID,
		&_CostCenterID,
		&_UnitCost,
		&_UnitType,
		&_UnitCurrency,
		&_Allocation,
		&_CreatedAt,
		&_UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &UserCostCenter{}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _Checksum.Valid {
		t.SetChecksum(_Checksum.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _UserID.Valid {
		t.SetUserID(_UserID.String)
	}
	if _CostCenterID.Valid {
		t.SetCostCenterID(_CostCenterID.String)
	}
	if _UnitCost.Valid {
		t.SetUnitCost(_UnitCost.Float64)
	}
	if _UnitType.Valid {
		t.SetUnitTypeString(_UnitType.String)
	}
	if _UnitCurrency.Valid {
		t.SetUnitCurrency(_UnitCurrency.String)
	}
	if _Allocation.Valid {
		t.SetAllocation(_Allocation.Float64)
	}
	if _CreatedAt.Valid {
		t.SetCreatedAt(_CreatedAt.Int64)
	}
	if _UpdatedAt.Valid {
		t.SetUpdatedAt(_UpdatedAt.Int64)
	}
	return t, nil
}

// GetChecksum will return the UserCostCenter Checksum value
func (t *UserCostCenter) GetChecksum() string {
	if t.Checksum == nil {
		return ""
	}
	return *t.Checksum
}

// SetChecksum will set the UserCostCenter Checksum value
func (t *UserCostCenter) SetChecksum(v string) {
	t.Checksum = &v
}

// GetCustomerID will return the UserCostCenter CustomerID value
func (t *UserCostCenter) GetCustomerID() string {
	return t.CustomerID
}

// SetCustomerID will set the UserCostCenter CustomerID value
func (t *UserCostCenter) SetCustomerID(v string) {
	t.CustomerID = v
}

// GetUserID will return the UserCostCenter UserID value
func (t *UserCostCenter) GetUserID() string {
	return t.UserID
}

// SetUserID will set the UserCostCenter UserID value
func (t *UserCostCenter) SetUserID(v string) {
	t.UserID = v
}

// GetCostCenterID will return the UserCostCenter CostCenterID value
func (t *UserCostCenter) GetCostCenterID() string {
	if t.CostCenterID == nil {
		return ""
	}
	return *t.CostCenterID
}

// SetCostCenterID will set the UserCostCenter CostCenterID value
func (t *UserCostCenter) SetCostCenterID(v string) {
	t.CostCenterID = &v
}

// GetUnitCost will return the UserCostCenter UnitCost value
func (t *UserCostCenter) GetUnitCost() float64 {
	if t.UnitCost == nil {
		return float64(0.0)
	}
	return *t.UnitCost
}

// SetUnitCost will set the UserCostCenter UnitCost value
func (t *UserCostCenter) SetUnitCost(v float64) {
	t.UnitCost = &v
}

// GetUnitType will return the UserCostCenter UnitType value
func (t *UserCostCenter) GetUnitType() UserCostCenter_UserUnitCostType {
	if t.UnitType == nil {
		return UserUnitCostType_SALARY
	}
	return *t.UnitType
}

// SetUnitType will set the UserCostCenter UnitType value
func (t *UserCostCenter) SetUnitType(v UserCostCenter_UserUnitCostType) {
	t.UnitType = &v
}

// GetUnitTypeString will return the UserCostCenter UnitType value as a string
func (t *UserCostCenter) GetUnitTypeString() string {
	if t.UnitType == nil {
		return ""
	}
	return t.UnitType.String()
}

// SetUnitTypeString will set the UserCostCenter UnitType value from a string
func (t *UserCostCenter) SetUnitTypeString(v string) {
	t.UnitType = toUserUnitCostType(v)
}

// GetUnitCurrency will return the UserCostCenter UnitCurrency value
func (t *UserCostCenter) GetUnitCurrency() string {
	if t.UnitCurrency == nil {
		return ""
	}
	return *t.UnitCurrency
}

// SetUnitCurrency will set the UserCostCenter UnitCurrency value
func (t *UserCostCenter) SetUnitCurrency(v string) {
	t.UnitCurrency = &v
}

// GetAllocation will return the UserCostCenter Allocation value
func (t *UserCostCenter) GetAllocation() float64 {
	if t.Allocation == nil {
		return float64(0.0)
	}
	return *t.Allocation
}

// SetAllocation will set the UserCostCenter Allocation value
func (t *UserCostCenter) SetAllocation(v float64) {
	t.Allocation = &v
}

// GetCreatedAt will return the UserCostCenter CreatedAt value
func (t *UserCostCenter) GetCreatedAt() int64 {
	return t.CreatedAt
}

// SetCreatedAt will set the UserCostCenter CreatedAt value
func (t *UserCostCenter) SetCreatedAt(v int64) {
	t.CreatedAt = v
}

// GetUpdatedAt will return the UserCostCenter UpdatedAt value
func (t *UserCostCenter) GetUpdatedAt() int64 {
	if t.UpdatedAt == nil {
		return int64(0)
	}
	return *t.UpdatedAt
}

// SetUpdatedAt will set the UserCostCenter UpdatedAt value
func (t *UserCostCenter) SetUpdatedAt(v int64) {
	t.UpdatedAt = &v
}

func (t *UserCostCenter) toTimestamp(value time.Time) *timestamp.Timestamp {
	ts, _ := ptypes.TimestampProto(value)
	return ts
}

// DBCreateUserCostCenterTable will create the UserCostCenter table
func DBCreateUserCostCenterTable(ctx context.Context, db DB) error {
	q := "CREATE TABLE `user_cost_center` (`id` VARCHAR(64) NOT NULL PRIMARY KEY,`checksum` CHAR(64),`customer_id` VARCHAR(64) NOT NULL,`user_id` VARCHAR(64) NOT NULL,`cost_center_id` VARCHAR(64),`unit_cost`FLOAT,`unit_type`ENUM('salary','hourly','contractor','casual','other'),`unit_currency` CHAR(3),`allocation` FLOAT,`created_at` BIGINT UNSIGNED NOT NULL,`updated_at` BIGINT UNSIGNED) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBCreateUserCostCenterTableTx will create the UserCostCenter table using the provided transction
func DBCreateUserCostCenterTableTx(ctx context.Context, tx Tx) error {
	q := "CREATE TABLE `user_cost_center` (`id` VARCHAR(64) NOT NULL PRIMARY KEY,`checksum` CHAR(64),`customer_id` VARCHAR(64) NOT NULL,`user_id` VARCHAR(64) NOT NULL,`cost_center_id` VARCHAR(64),`unit_cost`FLOAT,`unit_type`ENUM('salary','hourly','contractor','casual','other'),`unit_currency` CHAR(3),`allocation` FLOAT,`created_at` BIGINT UNSIGNED NOT NULL,`updated_at` BIGINT UNSIGNED) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// DBDropUserCostCenterTable will drop the UserCostCenter table
func DBDropUserCostCenterTable(ctx context.Context, db DB) error {
	q := "DROP TABLE IF EXISTS `user_cost_center`"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBDropUserCostCenterTableTx will drop the UserCostCenter table using the provided transaction
func DBDropUserCostCenterTableTx(ctx context.Context, tx Tx) error {
	q := "DROP TABLE IF EXISTS `user_cost_center`"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// CalculateChecksum will calculate a checksum of the SHA1 of all field values
func (t *UserCostCenter) CalculateChecksum() string {
	return orm.HashStrings(
		orm.ToString(t.ID),
		orm.ToString(t.CustomerID),
		orm.ToString(t.UserID),
		orm.ToString(t.CostCenterID),
		orm.ToString(t.UnitCost),
		enumUserUnitCostTypeToString(t.UnitType),
		orm.ToString(t.UnitCurrency),
		orm.ToString(t.Allocation),
		orm.ToString(t.CreatedAt),
		orm.ToString(t.UpdatedAt),
	)
}

// DBCreate will create a new UserCostCenter record in the database
func (t *UserCostCenter) DBCreate(ctx context.Context, db DB) (sql.Result, error) {
	q := "INSERT INTO `user_cost_center` (`user_cost_center`.`id`,`user_cost_center`.`checksum`,`user_cost_center`.`customer_id`,`user_cost_center`.`user_id`,`user_cost_center`.`cost_center_id`,`user_cost_center`.`unit_cost`,`user_cost_center`.`unit_type`,`user_cost_center`.`unit_currency`,`user_cost_center`.`allocation`,`user_cost_center`.`created_at`,`user_cost_center`.`updated_at`) VALUES (?,?,?,?,?,?,?,?,?,?,?)"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.UserID),
		orm.ToSQLString(t.CostCenterID),
		orm.ToSQLFloat64(t.UnitCost),
		orm.ToSQLString(enumUserUnitCostTypeToString(t.UnitType)),
		orm.ToSQLString(t.UnitCurrency),
		orm.ToSQLFloat64(t.Allocation),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
	)
}

// DBCreateTx will create a new UserCostCenter record in the database using the provided transaction
func (t *UserCostCenter) DBCreateTx(ctx context.Context, tx Tx) (sql.Result, error) {
	q := "INSERT INTO `user_cost_center` (`user_cost_center`.`id`,`user_cost_center`.`checksum`,`user_cost_center`.`customer_id`,`user_cost_center`.`user_id`,`user_cost_center`.`cost_center_id`,`user_cost_center`.`unit_cost`,`user_cost_center`.`unit_type`,`user_cost_center`.`unit_currency`,`user_cost_center`.`allocation`,`user_cost_center`.`created_at`,`user_cost_center`.`updated_at`) VALUES (?,?,?,?,?,?,?,?,?,?,?)"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.UserID),
		orm.ToSQLString(t.CostCenterID),
		orm.ToSQLFloat64(t.UnitCost),
		orm.ToSQLString(enumUserUnitCostTypeToString(t.UnitType)),
		orm.ToSQLString(t.UnitCurrency),
		orm.ToSQLFloat64(t.Allocation),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
	)
}

// DBCreateIgnoreDuplicate will upsert the UserCostCenter record in the database
func (t *UserCostCenter) DBCreateIgnoreDuplicate(ctx context.Context, db DB) (sql.Result, error) {
	q := "INSERT INTO `user_cost_center` (`user_cost_center`.`id`,`user_cost_center`.`checksum`,`user_cost_center`.`customer_id`,`user_cost_center`.`user_id`,`user_cost_center`.`cost_center_id`,`user_cost_center`.`unit_cost`,`user_cost_center`.`unit_type`,`user_cost_center`.`unit_currency`,`user_cost_center`.`allocation`,`user_cost_center`.`created_at`,`user_cost_center`.`updated_at`) VALUES (?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.UserID),
		orm.ToSQLString(t.CostCenterID),
		orm.ToSQLFloat64(t.UnitCost),
		orm.ToSQLString(enumUserUnitCostTypeToString(t.UnitType)),
		orm.ToSQLString(t.UnitCurrency),
		orm.ToSQLFloat64(t.Allocation),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
	)
}

// DBCreateIgnoreDuplicateTx will upsert the UserCostCenter record in the database using the provided transaction
func (t *UserCostCenter) DBCreateIgnoreDuplicateTx(ctx context.Context, tx Tx) (sql.Result, error) {
	q := "INSERT INTO `user_cost_center` (`user_cost_center`.`id`,`user_cost_center`.`checksum`,`user_cost_center`.`customer_id`,`user_cost_center`.`user_id`,`user_cost_center`.`cost_center_id`,`user_cost_center`.`unit_cost`,`user_cost_center`.`unit_type`,`user_cost_center`.`unit_currency`,`user_cost_center`.`allocation`,`user_cost_center`.`created_at`,`user_cost_center`.`updated_at`) VALUES (?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.UserID),
		orm.ToSQLString(t.CostCenterID),
		orm.ToSQLFloat64(t.UnitCost),
		orm.ToSQLString(enumUserUnitCostTypeToString(t.UnitType)),
		orm.ToSQLString(t.UnitCurrency),
		orm.ToSQLFloat64(t.Allocation),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
	)
}

// DeleteAllUserCostCenters deletes all UserCostCenter records in the database with optional filters
func DeleteAllUserCostCenters(ctx context.Context, db DB, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(UserCostCenterTableName),
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

// DeleteAllUserCostCentersTx deletes all UserCostCenter records in the database with optional filters using the provided transaction
func DeleteAllUserCostCentersTx(ctx context.Context, tx Tx, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(UserCostCenterTableName),
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

// DBDelete will delete this UserCostCenter record in the database
func (t *UserCostCenter) DBDelete(ctx context.Context, db DB) (bool, error) {
	q := "DELETE FROM `user_cost_center` WHERE `id` = ?"
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

// DBDeleteTx will delete this UserCostCenter record in the database using the provided transaction
func (t *UserCostCenter) DBDeleteTx(ctx context.Context, tx Tx) (bool, error) {
	q := "DELETE FROM `user_cost_center` WHERE `id` = ?"
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

// DBUpdate will update the UserCostCenter record in the database
func (t *UserCostCenter) DBUpdate(ctx context.Context, db DB) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `user_cost_center` SET `checksum`=?,`customer_id`=?,`user_id`=?,`cost_center_id`=?,`unit_cost`=?,`unit_type`=?,`unit_currency`=?,`allocation`=?,`created_at`=?,`updated_at`=? WHERE `id`=?"
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.UserID),
		orm.ToSQLString(t.CostCenterID),
		orm.ToSQLFloat64(t.UnitCost),
		orm.ToSQLString(enumUserUnitCostTypeToString(t.UnitType)),
		orm.ToSQLString(t.UnitCurrency),
		orm.ToSQLFloat64(t.Allocation),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
		orm.ToSQLString(t.ID),
	)
}

// DBUpdateTx will update the UserCostCenter record in the database using the provided transaction
func (t *UserCostCenter) DBUpdateTx(ctx context.Context, tx Tx) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `user_cost_center` SET `checksum`=?,`customer_id`=?,`user_id`=?,`cost_center_id`=?,`unit_cost`=?,`unit_type`=?,`unit_currency`=?,`allocation`=?,`created_at`=?,`updated_at`=? WHERE `id`=?"
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.UserID),
		orm.ToSQLString(t.CostCenterID),
		orm.ToSQLFloat64(t.UnitCost),
		orm.ToSQLString(enumUserUnitCostTypeToString(t.UnitType)),
		orm.ToSQLString(t.UnitCurrency),
		orm.ToSQLFloat64(t.Allocation),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
		orm.ToSQLString(t.ID),
	)
}

// DBUpsert will upsert the UserCostCenter record in the database
func (t *UserCostCenter) DBUpsert(ctx context.Context, db DB, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `user_cost_center` (`user_cost_center`.`id`,`user_cost_center`.`checksum`,`user_cost_center`.`customer_id`,`user_cost_center`.`user_id`,`user_cost_center`.`cost_center_id`,`user_cost_center`.`unit_cost`,`user_cost_center`.`unit_type`,`user_cost_center`.`unit_currency`,`user_cost_center`.`allocation`,`user_cost_center`.`created_at`,`user_cost_center`.`updated_at`) VALUES (?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `user_cost_center` (`user_cost_center`.`id`,`user_cost_center`.`checksum`,`user_cost_center`.`customer_id`,`user_cost_center`.`user_id`,`user_cost_center`.`cost_center_id`,`user_cost_center`.`unit_cost`,`user_cost_center`.`unit_type`,`user_cost_center`.`unit_currency`,`user_cost_center`.`allocation`,`user_cost_center`.`created_at`,`user_cost_center`.`updated_at`) VALUES (?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`customer_id`=VALUES(`customer_id`),`user_id`=VALUES(`user_id`),`cost_center_id`=VALUES(`cost_center_id`),`unit_cost`=VALUES(`unit_cost`),`unit_type`=VALUES(`unit_type`),`unit_currency`=VALUES(`unit_currency`),`allocation`=VALUES(`allocation`),`created_at`=VALUES(`created_at`),`updated_at`=VALUES(`updated_at`)"
	}
	r, err := db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.UserID),
		orm.ToSQLString(t.CostCenterID),
		orm.ToSQLFloat64(t.UnitCost),
		orm.ToSQLString(enumUserUnitCostTypeToString(t.UnitType)),
		orm.ToSQLString(t.UnitCurrency),
		orm.ToSQLFloat64(t.Allocation),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBUpsertTx will upsert the UserCostCenter record in the database using the provided transaction
func (t *UserCostCenter) DBUpsertTx(ctx context.Context, tx Tx, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `user_cost_center` (`user_cost_center`.`id`,`user_cost_center`.`checksum`,`user_cost_center`.`customer_id`,`user_cost_center`.`user_id`,`user_cost_center`.`cost_center_id`,`user_cost_center`.`unit_cost`,`user_cost_center`.`unit_type`,`user_cost_center`.`unit_currency`,`user_cost_center`.`allocation`,`user_cost_center`.`created_at`,`user_cost_center`.`updated_at`) VALUES (?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `user_cost_center` (`user_cost_center`.`id`,`user_cost_center`.`checksum`,`user_cost_center`.`customer_id`,`user_cost_center`.`user_id`,`user_cost_center`.`cost_center_id`,`user_cost_center`.`unit_cost`,`user_cost_center`.`unit_type`,`user_cost_center`.`unit_currency`,`user_cost_center`.`allocation`,`user_cost_center`.`created_at`,`user_cost_center`.`updated_at`) VALUES (?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`customer_id`=VALUES(`customer_id`),`user_id`=VALUES(`user_id`),`cost_center_id`=VALUES(`cost_center_id`),`unit_cost`=VALUES(`unit_cost`),`unit_type`=VALUES(`unit_type`),`unit_currency`=VALUES(`unit_currency`),`allocation`=VALUES(`allocation`),`created_at`=VALUES(`created_at`),`updated_at`=VALUES(`updated_at`)"
	}
	r, err := tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.UserID),
		orm.ToSQLString(t.CostCenterID),
		orm.ToSQLFloat64(t.UnitCost),
		orm.ToSQLString(enumUserUnitCostTypeToString(t.UnitType)),
		orm.ToSQLString(t.UnitCurrency),
		orm.ToSQLFloat64(t.Allocation),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBFindOne will find a UserCostCenter record in the database with the primary key
func (t *UserCostCenter) DBFindOne(ctx context.Context, db DB, value string) (bool, error) {
	q := "SELECT `user_cost_center`.`id`,`user_cost_center`.`checksum`,`user_cost_center`.`customer_id`,`user_cost_center`.`user_id`,`user_cost_center`.`cost_center_id`,`user_cost_center`.`unit_cost`,`user_cost_center`.`unit_type`,`user_cost_center`.`unit_currency`,`user_cost_center`.`allocation`,`user_cost_center`.`created_at`,`user_cost_center`.`updated_at` FROM `user_cost_center` WHERE `id` = ? LIMIT 1"
	row := db.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CustomerID sql.NullString
	var _UserID sql.NullString
	var _CostCenterID sql.NullString
	var _UnitCost sql.NullFloat64
	var _UnitType sql.NullString
	var _UnitCurrency sql.NullString
	var _Allocation sql.NullFloat64
	var _CreatedAt sql.NullInt64
	var _UpdatedAt sql.NullInt64
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_UserID,
		&_CostCenterID,
		&_UnitCost,
		&_UnitType,
		&_UnitCurrency,
		&_Allocation,
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
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _UserID.Valid {
		t.SetUserID(_UserID.String)
	}
	if _CostCenterID.Valid {
		t.SetCostCenterID(_CostCenterID.String)
	}
	if _UnitCost.Valid {
		t.SetUnitCost(_UnitCost.Float64)
	}
	if _UnitType.Valid {
		t.SetUnitTypeString(_UnitType.String)
	}
	if _UnitCurrency.Valid {
		t.SetUnitCurrency(_UnitCurrency.String)
	}
	if _Allocation.Valid {
		t.SetAllocation(_Allocation.Float64)
	}
	if _CreatedAt.Valid {
		t.SetCreatedAt(_CreatedAt.Int64)
	}
	if _UpdatedAt.Valid {
		t.SetUpdatedAt(_UpdatedAt.Int64)
	}
	return true, nil
}

// DBFindOneTx will find a UserCostCenter record in the database with the primary key using the provided transaction
func (t *UserCostCenter) DBFindOneTx(ctx context.Context, tx Tx, value string) (bool, error) {
	q := "SELECT `user_cost_center`.`id`,`user_cost_center`.`checksum`,`user_cost_center`.`customer_id`,`user_cost_center`.`user_id`,`user_cost_center`.`cost_center_id`,`user_cost_center`.`unit_cost`,`user_cost_center`.`unit_type`,`user_cost_center`.`unit_currency`,`user_cost_center`.`allocation`,`user_cost_center`.`created_at`,`user_cost_center`.`updated_at` FROM `user_cost_center` WHERE `id` = ? LIMIT 1"
	row := tx.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CustomerID sql.NullString
	var _UserID sql.NullString
	var _CostCenterID sql.NullString
	var _UnitCost sql.NullFloat64
	var _UnitType sql.NullString
	var _UnitCurrency sql.NullString
	var _Allocation sql.NullFloat64
	var _CreatedAt sql.NullInt64
	var _UpdatedAt sql.NullInt64
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_UserID,
		&_CostCenterID,
		&_UnitCost,
		&_UnitType,
		&_UnitCurrency,
		&_Allocation,
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
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _UserID.Valid {
		t.SetUserID(_UserID.String)
	}
	if _CostCenterID.Valid {
		t.SetCostCenterID(_CostCenterID.String)
	}
	if _UnitCost.Valid {
		t.SetUnitCost(_UnitCost.Float64)
	}
	if _UnitType.Valid {
		t.SetUnitTypeString(_UnitType.String)
	}
	if _UnitCurrency.Valid {
		t.SetUnitCurrency(_UnitCurrency.String)
	}
	if _Allocation.Valid {
		t.SetAllocation(_Allocation.Float64)
	}
	if _CreatedAt.Valid {
		t.SetCreatedAt(_CreatedAt.Int64)
	}
	if _UpdatedAt.Valid {
		t.SetUpdatedAt(_UpdatedAt.Int64)
	}
	return true, nil
}

// FindUserCostCenters will find a UserCostCenter record in the database with the provided parameters
func FindUserCostCenters(ctx context.Context, db DB, _params ...interface{}) ([]*UserCostCenter, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("customer_id"),
		orm.Column("user_id"),
		orm.Column("cost_center_id"),
		orm.Column("unit_cost"),
		orm.Column("unit_type"),
		orm.Column("unit_currency"),
		orm.Column("allocation"),
		orm.Column("created_at"),
		orm.Column("updated_at"),
		orm.Table(UserCostCenterTableName),
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
	results := make([]*UserCostCenter, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _UserID sql.NullString
		var _CostCenterID sql.NullString
		var _UnitCost sql.NullFloat64
		var _UnitType sql.NullString
		var _UnitCurrency sql.NullString
		var _Allocation sql.NullFloat64
		var _CreatedAt sql.NullInt64
		var _UpdatedAt sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_UserID,
			&_CostCenterID,
			&_UnitCost,
			&_UnitType,
			&_UnitCurrency,
			&_Allocation,
			&_CreatedAt,
			&_UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		t := &UserCostCenter{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _CostCenterID.Valid {
			t.SetCostCenterID(_CostCenterID.String)
		}
		if _UnitCost.Valid {
			t.SetUnitCost(_UnitCost.Float64)
		}
		if _UnitType.Valid {
			t.SetUnitTypeString(_UnitType.String)
		}
		if _UnitCurrency.Valid {
			t.SetUnitCurrency(_UnitCurrency.String)
		}
		if _Allocation.Valid {
			t.SetAllocation(_Allocation.Float64)
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

// FindUserCostCentersTx will find a UserCostCenter record in the database with the provided parameters using the provided transaction
func FindUserCostCentersTx(ctx context.Context, tx Tx, _params ...interface{}) ([]*UserCostCenter, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("customer_id"),
		orm.Column("user_id"),
		orm.Column("cost_center_id"),
		orm.Column("unit_cost"),
		orm.Column("unit_type"),
		orm.Column("unit_currency"),
		orm.Column("allocation"),
		orm.Column("created_at"),
		orm.Column("updated_at"),
		orm.Table(UserCostCenterTableName),
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
	results := make([]*UserCostCenter, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _UserID sql.NullString
		var _CostCenterID sql.NullString
		var _UnitCost sql.NullFloat64
		var _UnitType sql.NullString
		var _UnitCurrency sql.NullString
		var _Allocation sql.NullFloat64
		var _CreatedAt sql.NullInt64
		var _UpdatedAt sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_UserID,
			&_CostCenterID,
			&_UnitCost,
			&_UnitType,
			&_UnitCurrency,
			&_Allocation,
			&_CreatedAt,
			&_UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		t := &UserCostCenter{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _CostCenterID.Valid {
			t.SetCostCenterID(_CostCenterID.String)
		}
		if _UnitCost.Valid {
			t.SetUnitCost(_UnitCost.Float64)
		}
		if _UnitType.Valid {
			t.SetUnitTypeString(_UnitType.String)
		}
		if _UnitCurrency.Valid {
			t.SetUnitCurrency(_UnitCurrency.String)
		}
		if _Allocation.Valid {
			t.SetAllocation(_Allocation.Float64)
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

// DBFind will find a UserCostCenter record in the database with the provided parameters
func (t *UserCostCenter) DBFind(ctx context.Context, db DB, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("customer_id"),
		orm.Column("user_id"),
		orm.Column("cost_center_id"),
		orm.Column("unit_cost"),
		orm.Column("unit_type"),
		orm.Column("unit_currency"),
		orm.Column("allocation"),
		orm.Column("created_at"),
		orm.Column("updated_at"),
		orm.Table(UserCostCenterTableName),
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
	var _UserID sql.NullString
	var _CostCenterID sql.NullString
	var _UnitCost sql.NullFloat64
	var _UnitType sql.NullString
	var _UnitCurrency sql.NullString
	var _Allocation sql.NullFloat64
	var _CreatedAt sql.NullInt64
	var _UpdatedAt sql.NullInt64
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_UserID,
		&_CostCenterID,
		&_UnitCost,
		&_UnitType,
		&_UnitCurrency,
		&_Allocation,
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
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _UserID.Valid {
		t.SetUserID(_UserID.String)
	}
	if _CostCenterID.Valid {
		t.SetCostCenterID(_CostCenterID.String)
	}
	if _UnitCost.Valid {
		t.SetUnitCost(_UnitCost.Float64)
	}
	if _UnitType.Valid {
		t.SetUnitTypeString(_UnitType.String)
	}
	if _UnitCurrency.Valid {
		t.SetUnitCurrency(_UnitCurrency.String)
	}
	if _Allocation.Valid {
		t.SetAllocation(_Allocation.Float64)
	}
	if _CreatedAt.Valid {
		t.SetCreatedAt(_CreatedAt.Int64)
	}
	if _UpdatedAt.Valid {
		t.SetUpdatedAt(_UpdatedAt.Int64)
	}
	return true, nil
}

// DBFindTx will find a UserCostCenter record in the database with the provided parameters using the provided transaction
func (t *UserCostCenter) DBFindTx(ctx context.Context, tx Tx, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("customer_id"),
		orm.Column("user_id"),
		orm.Column("cost_center_id"),
		orm.Column("unit_cost"),
		orm.Column("unit_type"),
		orm.Column("unit_currency"),
		orm.Column("allocation"),
		orm.Column("created_at"),
		orm.Column("updated_at"),
		orm.Table(UserCostCenterTableName),
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
	var _UserID sql.NullString
	var _CostCenterID sql.NullString
	var _UnitCost sql.NullFloat64
	var _UnitType sql.NullString
	var _UnitCurrency sql.NullString
	var _Allocation sql.NullFloat64
	var _CreatedAt sql.NullInt64
	var _UpdatedAt sql.NullInt64
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_UserID,
		&_CostCenterID,
		&_UnitCost,
		&_UnitType,
		&_UnitCurrency,
		&_Allocation,
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
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _UserID.Valid {
		t.SetUserID(_UserID.String)
	}
	if _CostCenterID.Valid {
		t.SetCostCenterID(_CostCenterID.String)
	}
	if _UnitCost.Valid {
		t.SetUnitCost(_UnitCost.Float64)
	}
	if _UnitType.Valid {
		t.SetUnitTypeString(_UnitType.String)
	}
	if _UnitCurrency.Valid {
		t.SetUnitCurrency(_UnitCurrency.String)
	}
	if _Allocation.Valid {
		t.SetAllocation(_Allocation.Float64)
	}
	if _CreatedAt.Valid {
		t.SetCreatedAt(_CreatedAt.Int64)
	}
	if _UpdatedAt.Valid {
		t.SetUpdatedAt(_UpdatedAt.Int64)
	}
	return true, nil
}

// CountUserCostCenters will find the count of UserCostCenter records in the database
func CountUserCostCenters(ctx context.Context, db DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(UserCostCenterTableName),
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

// CountUserCostCentersTx will find the count of UserCostCenter records in the database using the provided transaction
func CountUserCostCentersTx(ctx context.Context, tx Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(UserCostCenterTableName),
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

// DBCount will find the count of UserCostCenter records in the database
func (t *UserCostCenter) DBCount(ctx context.Context, db DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(UserCostCenterTableName),
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

// DBCountTx will find the count of UserCostCenter records in the database using the provided transaction
func (t *UserCostCenter) DBCountTx(ctx context.Context, tx Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(UserCostCenterTableName),
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

// DBExists will return true if the UserCostCenter record exists in the database
func (t *UserCostCenter) DBExists(ctx context.Context, db DB) (bool, error) {
	q := "SELECT `id` FROM `user_cost_center` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := db.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// DBExistsTx will return true if the UserCostCenter record exists in the database using the provided transaction
func (t *UserCostCenter) DBExistsTx(ctx context.Context, tx Tx) (bool, error) {
	q := "SELECT `id` FROM `user_cost_center` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := tx.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// PrimaryKeyColumn returns the column name for the primary key
func (t *UserCostCenter) PrimaryKeyColumn() string {
	return UserCostCenterColumnID
}

// PrimaryKeyColumnType returns the primary key column Go type as a string
func (t *UserCostCenter) PrimaryKeyColumnType() string {
	return "string"
}

// PrimaryKey returns the primary key column value
func (t *UserCostCenter) PrimaryKey() interface{} {
	return t.ID
}
