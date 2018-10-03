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

var _ Model = (*CostCenter)(nil)
var _ CSVWriter = (*CostCenter)(nil)
var _ JSONWriter = (*CostCenter)(nil)
var _ Checksum = (*CostCenter)(nil)

// CostCenterTableName is the name of the table in SQL
const CostCenterTableName = "cost_center"

var CostCenterColumns = []string{
	"id",
	"checksum",
	"customer_id",
	"name",
	"description",
	"identifier",
	"unit_cost",
	"value_type",
	"unit_type",
	"unit_currency",
	"allocation",
	"created_at",
	"updated_at",
}

type CostCenter_CostCenterValueType string

const (
	CostCenterValueType_ABSOLUTE   CostCenter_CostCenterValueType = "absolute"
	CostCenterValueType_PERCENTAGE CostCenter_CostCenterValueType = "percentage"
)

func (x CostCenter_CostCenterValueType) String() string {
	return string(x)
}

func enumCostCenterValueTypeToString(v *CostCenter_CostCenterValueType) string {
	if v == nil {
		return ""
	}
	return v.String()
}

func toCostCenterValueType(v string) *CostCenter_CostCenterValueType {
	var ev *CostCenter_CostCenterValueType
	switch v {
	case "ABSOLUTE", "absolute":
		{
			v := CostCenterValueType_ABSOLUTE
			ev = &v
		}
	case "PERCENTAGE", "percentage":
		{
			v := CostCenterValueType_PERCENTAGE
			ev = &v
		}
	}
	return ev
}

type CostCenter_CostCenterUnitCostType string

const (
	CostCenterUnitCostType_SALARY     CostCenter_CostCenterUnitCostType = "salary"
	CostCenterUnitCostType_HOURLY     CostCenter_CostCenterUnitCostType = "hourly"
	CostCenterUnitCostType_CONTRACTOR CostCenter_CostCenterUnitCostType = "contractor"
	CostCenterUnitCostType_CASUAL     CostCenter_CostCenterUnitCostType = "casual"
	CostCenterUnitCostType_OTHER      CostCenter_CostCenterUnitCostType = "other"
)

func (x CostCenter_CostCenterUnitCostType) String() string {
	return string(x)
}

func enumCostCenterUnitCostTypeToString(v *CostCenter_CostCenterUnitCostType) string {
	if v == nil {
		return ""
	}
	return v.String()
}

func toCostCenterUnitCostType(v string) *CostCenter_CostCenterUnitCostType {
	var ev *CostCenter_CostCenterUnitCostType
	switch v {
	case "SALARY", "salary":
		{
			v := CostCenterUnitCostType_SALARY
			ev = &v
		}
	case "HOURLY", "hourly":
		{
			v := CostCenterUnitCostType_HOURLY
			ev = &v
		}
	case "CONTRACTOR", "contractor":
		{
			v := CostCenterUnitCostType_CONTRACTOR
			ev = &v
		}
	case "CASUAL", "casual":
		{
			v := CostCenterUnitCostType_CASUAL
			ev = &v
		}
	case "OTHER", "other":
		{
			v := CostCenterUnitCostType_OTHER
			ev = &v
		}
	}
	return ev
}

// CostCenter table
type CostCenter struct {
	Allocation   float64                           `json:"allocation"`
	Checksum     *string                           `json:"checksum,omitempty"`
	CreatedAt    int64                             `json:"created_at"`
	CustomerID   string                            `json:"customer_id"`
	Description  *string                           `json:"description,omitempty"`
	ID           string                            `json:"id"`
	Identifier   *string                           `json:"identifier,omitempty"`
	Name         string                            `json:"name"`
	UnitCost     float64                           `json:"unit_cost"`
	UnitCurrency string                            `json:"unit_currency"`
	UnitType     CostCenter_CostCenterUnitCostType `json:"unit_type"`
	UpdatedAt    *int64                            `json:"updated_at,omitempty"`
	ValueType    CostCenter_CostCenterValueType    `json:"value_type"`
}

// TableName returns the SQL table name for CostCenter and satifies the Model interface
func (t *CostCenter) TableName() string {
	return CostCenterTableName
}

// ToCSV will serialize the CostCenter instance to a CSV compatible array of strings
func (t *CostCenter) ToCSV() []string {
	return []string{
		t.ID,
		t.CalculateChecksum(),
		t.CustomerID,
		t.Name,
		toCSVString(t.Description),
		toCSVString(t.Identifier),
		toCSVString(t.UnitCost),
		toCSVString(t.ValueType),
		toCSVString(t.UnitType),
		t.UnitCurrency,
		toCSVString(t.Allocation),
		toCSVString(t.CreatedAt),
		toCSVString(t.UpdatedAt),
	}
}

// WriteCSV will serialize the CostCenter instance to the writer as CSV and satisfies the CSVWriter interface
func (t *CostCenter) WriteCSV(w *csv.Writer) error {
	return w.Write(t.ToCSV())
}

// WriteJSON will serialize the CostCenter instance to the writer as JSON and satisfies the JSONWriter interface
func (t *CostCenter) WriteJSON(w io.Writer, indent ...bool) error {
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

// NewCostCenterReader creates a JSON reader which can read in CostCenter objects serialized as JSON either as an array, single object or json new lines
// and writes each CostCenter to the channel provided
func NewCostCenterReader(r io.Reader, ch chan<- CostCenter) error {
	return orm.Deserialize(r, func(buf json.RawMessage) error {
		dec := json.NewDecoder(bytes.NewBuffer(buf))
		e := CostCenter{}
		if err := dec.Decode(&e); err != nil {
			return err
		}
		ch <- e
		return nil
	})
}

// NewCSVCostCenterReaderDir will read the reader as CSV and emit each record to the channel provided
func NewCSVCostCenterReader(r io.Reader, ch chan<- CostCenter) error {
	cr := csv.NewReader(r)
	for {
		record, err := cr.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		ch <- CostCenter{
			ID:           record[0],
			Checksum:     fromStringPointer(record[1]),
			CustomerID:   record[2],
			Name:         record[3],
			Description:  fromStringPointer(record[4]),
			Identifier:   fromStringPointer(record[5]),
			UnitCost:     fromCSVFloat64(record[6]),
			ValueType:    *toCostCenterValueType(record[7]),
			UnitType:     *toCostCenterUnitCostType(record[8]),
			UnitCurrency: record[9],
			Allocation:   fromCSVFloat64(record[10]),
			CreatedAt:    fromCSVInt64(record[11]),
			UpdatedAt:    fromCSVInt64Pointer(record[12]),
		}
	}
	return nil
}

// NewCSVCostCenterReaderFile will read the file as a CSV and emit each record to the channel provided
func NewCSVCostCenterReaderFile(fp string, ch chan<- CostCenter) error {
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
	return NewCSVCostCenterReader(fc, ch)
}

// NewCSVCostCenterReaderDir will read the cost_center.csv.gz file as a CSV and emit each record to the channel provided
func NewCSVCostCenterReaderDir(dir string, ch chan<- CostCenter) error {
	return NewCSVCostCenterReaderFile(filepath.Join(dir, "cost_center.csv.gz"), ch)
}

// CostCenterCSVDeduper is a function callback which takes the existing value (a) and the new value (b)
// and the return value should be the one to use (or a new one, if applicable). return nil
// to skip processing of this record
type CostCenterCSVDeduper func(a CostCenter, b CostCenter) *CostCenter

// CostCenterCSVDedupeDisabled is set on whether the CSV writer should de-dupe values by key
var CostCenterCSVDedupeDisabled bool

// NewCostCenterCSVWriterSize creates a batch writer that will write each CostCenter into a CSV file
// this method will automatically de-duplicate entries using the primary key. if the checksum
// for a newer item with the same primary key doesn't match a previously sent item, the newer
// one will be used
func NewCostCenterCSVWriterSize(w io.Writer, size int, dedupers ...CostCenterCSVDeduper) (chan CostCenter, chan bool, error) {
	cw := csv.NewWriter(w)
	ch := make(chan CostCenter, size)
	done := make(chan bool)
	go func() {
		defer func() { done <- true }()
		dodedupe := !CostCenterCSVDedupeDisabled
		var kv map[string]*CostCenter
		var deduper CostCenterCSVDeduper
		if dedupers != nil && len(dedupers) > 0 {
			deduper = dedupers[0]
			dodedupe = true
		}
		if dodedupe {
			kv = make(map[string]*CostCenter)
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

// CostCenterCSVDefaultSize is the default channel buffer size if not provided
var CostCenterCSVDefaultSize = 100

// NewCostCenterCSVWriter creates a batch writer that will write each CostCenter into a CSV file
func NewCostCenterCSVWriter(w io.Writer, dedupers ...CostCenterCSVDeduper) (chan CostCenter, chan bool, error) {
	return NewCostCenterCSVWriterSize(w, CostCenterCSVDefaultSize, dedupers...)
}

// NewCostCenterCSVWriterDir creates a batch writer that will write each CostCenter into a CSV file named cost_center.csv.gz in dir
func NewCostCenterCSVWriterDir(dir string, dedupers ...CostCenterCSVDeduper) (chan CostCenter, chan bool, error) {
	return NewCostCenterCSVWriterFile(filepath.Join(dir, "cost_center.csv.gz"), dedupers...)
}

// NewCostCenterCSVWriterFile creates a batch writer that will write each CostCenter into a CSV file
func NewCostCenterCSVWriterFile(fn string, dedupers ...CostCenterCSVDeduper) (chan CostCenter, chan bool, error) {
	f, err := os.Create(fn)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening CSV file %s. %v", fn, err)
	}
	var fc io.WriteCloser = f
	if filepath.Ext(fn) == ".gz" {
		w, _ := gzip.NewWriterLevel(f, gzip.BestCompression)
		fc = w
	}
	ch, done, err := NewCostCenterCSVWriter(fc, dedupers...)
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

type CostCenterDBAction func(ctx context.Context, db *sql.DB, record CostCenter) error

// NewCostCenterDBWriterSize creates a DB writer that will write each issue into the DB
func NewCostCenterDBWriterSize(ctx context.Context, db *sql.DB, errors chan<- error, size int, actions ...CostCenterDBAction) (chan CostCenter, chan bool, error) {
	ch := make(chan CostCenter, size)
	done := make(chan bool)
	var action CostCenterDBAction
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

// NewCostCenterDBWriter creates a DB writer that will write each issue into the DB
func NewCostCenterDBWriter(ctx context.Context, db *sql.DB, errors chan<- error, actions ...CostCenterDBAction) (chan CostCenter, chan bool, error) {
	return NewCostCenterDBWriterSize(ctx, db, errors, 100, actions...)
}

// CostCenterColumnID is the ID SQL column name for the CostCenter table
const CostCenterColumnID = "id"

// CostCenterEscapedColumnID is the escaped ID SQL column name for the CostCenter table
const CostCenterEscapedColumnID = "`id`"

// CostCenterColumnChecksum is the Checksum SQL column name for the CostCenter table
const CostCenterColumnChecksum = "checksum"

// CostCenterEscapedColumnChecksum is the escaped Checksum SQL column name for the CostCenter table
const CostCenterEscapedColumnChecksum = "`checksum`"

// CostCenterColumnCustomerID is the CustomerID SQL column name for the CostCenter table
const CostCenterColumnCustomerID = "customer_id"

// CostCenterEscapedColumnCustomerID is the escaped CustomerID SQL column name for the CostCenter table
const CostCenterEscapedColumnCustomerID = "`customer_id`"

// CostCenterColumnName is the Name SQL column name for the CostCenter table
const CostCenterColumnName = "name"

// CostCenterEscapedColumnName is the escaped Name SQL column name for the CostCenter table
const CostCenterEscapedColumnName = "`name`"

// CostCenterColumnDescription is the Description SQL column name for the CostCenter table
const CostCenterColumnDescription = "description"

// CostCenterEscapedColumnDescription is the escaped Description SQL column name for the CostCenter table
const CostCenterEscapedColumnDescription = "`description`"

// CostCenterColumnIdentifier is the Identifier SQL column name for the CostCenter table
const CostCenterColumnIdentifier = "identifier"

// CostCenterEscapedColumnIdentifier is the escaped Identifier SQL column name for the CostCenter table
const CostCenterEscapedColumnIdentifier = "`identifier`"

// CostCenterColumnUnitCost is the UnitCost SQL column name for the CostCenter table
const CostCenterColumnUnitCost = "unit_cost"

// CostCenterEscapedColumnUnitCost is the escaped UnitCost SQL column name for the CostCenter table
const CostCenterEscapedColumnUnitCost = "`unit_cost`"

// CostCenterColumnValueType is the ValueType SQL column name for the CostCenter table
const CostCenterColumnValueType = "value_type"

// CostCenterEscapedColumnValueType is the escaped ValueType SQL column name for the CostCenter table
const CostCenterEscapedColumnValueType = "`value_type`"

// CostCenterColumnUnitType is the UnitType SQL column name for the CostCenter table
const CostCenterColumnUnitType = "unit_type"

// CostCenterEscapedColumnUnitType is the escaped UnitType SQL column name for the CostCenter table
const CostCenterEscapedColumnUnitType = "`unit_type`"

// CostCenterColumnUnitCurrency is the UnitCurrency SQL column name for the CostCenter table
const CostCenterColumnUnitCurrency = "unit_currency"

// CostCenterEscapedColumnUnitCurrency is the escaped UnitCurrency SQL column name for the CostCenter table
const CostCenterEscapedColumnUnitCurrency = "`unit_currency`"

// CostCenterColumnAllocation is the Allocation SQL column name for the CostCenter table
const CostCenterColumnAllocation = "allocation"

// CostCenterEscapedColumnAllocation is the escaped Allocation SQL column name for the CostCenter table
const CostCenterEscapedColumnAllocation = "`allocation`"

// CostCenterColumnCreatedAt is the CreatedAt SQL column name for the CostCenter table
const CostCenterColumnCreatedAt = "created_at"

// CostCenterEscapedColumnCreatedAt is the escaped CreatedAt SQL column name for the CostCenter table
const CostCenterEscapedColumnCreatedAt = "`created_at`"

// CostCenterColumnUpdatedAt is the UpdatedAt SQL column name for the CostCenter table
const CostCenterColumnUpdatedAt = "updated_at"

// CostCenterEscapedColumnUpdatedAt is the escaped UpdatedAt SQL column name for the CostCenter table
const CostCenterEscapedColumnUpdatedAt = "`updated_at`"

// GetID will return the CostCenter ID value
func (t *CostCenter) GetID() string {
	return t.ID
}

// SetID will set the CostCenter ID value
func (t *CostCenter) SetID(v string) {
	t.ID = v
}

// FindCostCenterByID will find a CostCenter by ID
func FindCostCenterByID(ctx context.Context, db *sql.DB, value string) (*CostCenter, error) {
	q := "SELECT `cost_center`.`id`,`cost_center`.`checksum`,`cost_center`.`customer_id`,`cost_center`.`name`,`cost_center`.`description`,`cost_center`.`identifier`,`cost_center`.`unit_cost`,`cost_center`.`value_type`,`cost_center`.`unit_type`,`cost_center`.`unit_currency`,`cost_center`.`allocation`,`cost_center`.`created_at`,`cost_center`.`updated_at` FROM `cost_center` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CustomerID sql.NullString
	var _Name sql.NullString
	var _Description sql.NullString
	var _Identifier sql.NullString
	var _UnitCost sql.NullFloat64
	var _ValueType sql.NullString
	var _UnitType sql.NullString
	var _UnitCurrency sql.NullString
	var _Allocation sql.NullFloat64
	var _CreatedAt sql.NullInt64
	var _UpdatedAt sql.NullInt64
	err := db.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_Name,
		&_Description,
		&_Identifier,
		&_UnitCost,
		&_ValueType,
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
	t := &CostCenter{}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _Checksum.Valid {
		t.SetChecksum(_Checksum.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _Description.Valid {
		t.SetDescription(_Description.String)
	}
	if _Identifier.Valid {
		t.SetIdentifier(_Identifier.String)
	}
	if _UnitCost.Valid {
		t.SetUnitCost(_UnitCost.Float64)
	}
	if _ValueType.Valid {
		t.SetValueTypeString(_ValueType.String)
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

// FindCostCenterByIDTx will find a CostCenter by ID using the provided transaction
func FindCostCenterByIDTx(ctx context.Context, tx *sql.Tx, value string) (*CostCenter, error) {
	q := "SELECT `cost_center`.`id`,`cost_center`.`checksum`,`cost_center`.`customer_id`,`cost_center`.`name`,`cost_center`.`description`,`cost_center`.`identifier`,`cost_center`.`unit_cost`,`cost_center`.`value_type`,`cost_center`.`unit_type`,`cost_center`.`unit_currency`,`cost_center`.`allocation`,`cost_center`.`created_at`,`cost_center`.`updated_at` FROM `cost_center` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CustomerID sql.NullString
	var _Name sql.NullString
	var _Description sql.NullString
	var _Identifier sql.NullString
	var _UnitCost sql.NullFloat64
	var _ValueType sql.NullString
	var _UnitType sql.NullString
	var _UnitCurrency sql.NullString
	var _Allocation sql.NullFloat64
	var _CreatedAt sql.NullInt64
	var _UpdatedAt sql.NullInt64
	err := tx.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_Name,
		&_Description,
		&_Identifier,
		&_UnitCost,
		&_ValueType,
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
	t := &CostCenter{}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _Checksum.Valid {
		t.SetChecksum(_Checksum.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _Description.Valid {
		t.SetDescription(_Description.String)
	}
	if _Identifier.Valid {
		t.SetIdentifier(_Identifier.String)
	}
	if _UnitCost.Valid {
		t.SetUnitCost(_UnitCost.Float64)
	}
	if _ValueType.Valid {
		t.SetValueTypeString(_ValueType.String)
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

// GetChecksum will return the CostCenter Checksum value
func (t *CostCenter) GetChecksum() string {
	if t.Checksum == nil {
		return ""
	}
	return *t.Checksum
}

// SetChecksum will set the CostCenter Checksum value
func (t *CostCenter) SetChecksum(v string) {
	t.Checksum = &v
}

// GetCustomerID will return the CostCenter CustomerID value
func (t *CostCenter) GetCustomerID() string {
	return t.CustomerID
}

// SetCustomerID will set the CostCenter CustomerID value
func (t *CostCenter) SetCustomerID(v string) {
	t.CustomerID = v
}

// GetName will return the CostCenter Name value
func (t *CostCenter) GetName() string {
	return t.Name
}

// SetName will set the CostCenter Name value
func (t *CostCenter) SetName(v string) {
	t.Name = v
}

// GetDescription will return the CostCenter Description value
func (t *CostCenter) GetDescription() string {
	if t.Description == nil {
		return ""
	}
	return *t.Description
}

// SetDescription will set the CostCenter Description value
func (t *CostCenter) SetDescription(v string) {
	t.Description = &v
}

// GetIdentifier will return the CostCenter Identifier value
func (t *CostCenter) GetIdentifier() string {
	if t.Identifier == nil {
		return ""
	}
	return *t.Identifier
}

// SetIdentifier will set the CostCenter Identifier value
func (t *CostCenter) SetIdentifier(v string) {
	t.Identifier = &v
}

// GetUnitCost will return the CostCenter UnitCost value
func (t *CostCenter) GetUnitCost() float64 {
	return t.UnitCost
}

// SetUnitCost will set the CostCenter UnitCost value
func (t *CostCenter) SetUnitCost(v float64) {
	t.UnitCost = v
}

// GetValueType will return the CostCenter ValueType value
func (t *CostCenter) GetValueType() CostCenter_CostCenterValueType {
	return t.ValueType
}

// SetValueType will set the CostCenter ValueType value
func (t *CostCenter) SetValueType(v CostCenter_CostCenterValueType) {
	t.ValueType = v
}

// GetValueTypeString will return the CostCenter ValueType value as a string
func (t *CostCenter) GetValueTypeString() string {
	return t.ValueType.String()
}

// SetValueTypeString will set the CostCenter ValueType value from a string
func (t *CostCenter) SetValueTypeString(v string) {
	var _ValueType = toCostCenterValueType(v)
	if _ValueType != nil {
		t.ValueType = *_ValueType
	}
}

// GetUnitType will return the CostCenter UnitType value
func (t *CostCenter) GetUnitType() CostCenter_CostCenterUnitCostType {
	return t.UnitType
}

// SetUnitType will set the CostCenter UnitType value
func (t *CostCenter) SetUnitType(v CostCenter_CostCenterUnitCostType) {
	t.UnitType = v
}

// GetUnitTypeString will return the CostCenter UnitType value as a string
func (t *CostCenter) GetUnitTypeString() string {
	return t.UnitType.String()
}

// SetUnitTypeString will set the CostCenter UnitType value from a string
func (t *CostCenter) SetUnitTypeString(v string) {
	var _UnitType = toCostCenterUnitCostType(v)
	if _UnitType != nil {
		t.UnitType = *_UnitType
	}
}

// GetUnitCurrency will return the CostCenter UnitCurrency value
func (t *CostCenter) GetUnitCurrency() string {
	return t.UnitCurrency
}

// SetUnitCurrency will set the CostCenter UnitCurrency value
func (t *CostCenter) SetUnitCurrency(v string) {
	t.UnitCurrency = v
}

// GetAllocation will return the CostCenter Allocation value
func (t *CostCenter) GetAllocation() float64 {
	return t.Allocation
}

// SetAllocation will set the CostCenter Allocation value
func (t *CostCenter) SetAllocation(v float64) {
	t.Allocation = v
}

// GetCreatedAt will return the CostCenter CreatedAt value
func (t *CostCenter) GetCreatedAt() int64 {
	return t.CreatedAt
}

// SetCreatedAt will set the CostCenter CreatedAt value
func (t *CostCenter) SetCreatedAt(v int64) {
	t.CreatedAt = v
}

// GetUpdatedAt will return the CostCenter UpdatedAt value
func (t *CostCenter) GetUpdatedAt() int64 {
	if t.UpdatedAt == nil {
		return int64(0)
	}
	return *t.UpdatedAt
}

// SetUpdatedAt will set the CostCenter UpdatedAt value
func (t *CostCenter) SetUpdatedAt(v int64) {
	t.UpdatedAt = &v
}

func (t *CostCenter) toTimestamp(value time.Time) *timestamp.Timestamp {
	ts, _ := ptypes.TimestampProto(value)
	return ts
}

// DBCreateCostCenterTable will create the CostCenter table
func DBCreateCostCenterTable(ctx context.Context, db *sql.DB) error {
	q := "CREATE TABLE `cost_center` (`id` VARCHAR(64) NOT NULL PRIMARY KEY,`checksum` CHAR(64),`customer_id` VARCHAR(64) NOT NULL,`name` TEXT NOT NULL,`description` TEXT,`identifier` TEXT,`unit_cost`FLOAT NOT NULL,`value_type` ENUM('absolute','percentage') NOT NULL,`unit_type`ENUM('salary','hourly','contractor','casual','other') NOT NULL,`unit_currency` CHAR(3) NOT NULL,`allocation` FLOAT NOT NULL,`created_at` BIGINT UNSIGNED NOT NULL,`updated_at` BIGINT UNSIGNED) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBCreateCostCenterTableTx will create the CostCenter table using the provided transction
func DBCreateCostCenterTableTx(ctx context.Context, tx *sql.Tx) error {
	q := "CREATE TABLE `cost_center` (`id` VARCHAR(64) NOT NULL PRIMARY KEY,`checksum` CHAR(64),`customer_id` VARCHAR(64) NOT NULL,`name` TEXT NOT NULL,`description` TEXT,`identifier` TEXT,`unit_cost`FLOAT NOT NULL,`value_type` ENUM('absolute','percentage') NOT NULL,`unit_type`ENUM('salary','hourly','contractor','casual','other') NOT NULL,`unit_currency` CHAR(3) NOT NULL,`allocation` FLOAT NOT NULL,`created_at` BIGINT UNSIGNED NOT NULL,`updated_at` BIGINT UNSIGNED) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// DBDropCostCenterTable will drop the CostCenter table
func DBDropCostCenterTable(ctx context.Context, db *sql.DB) error {
	q := "DROP TABLE IF EXISTS `cost_center`"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBDropCostCenterTableTx will drop the CostCenter table using the provided transaction
func DBDropCostCenterTableTx(ctx context.Context, tx *sql.Tx) error {
	q := "DROP TABLE IF EXISTS `cost_center`"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// CalculateChecksum will calculate a checksum of the SHA1 of all field values
func (t *CostCenter) CalculateChecksum() string {
	return orm.HashStrings(
		orm.ToString(t.ID),
		orm.ToString(t.CustomerID),
		orm.ToString(t.Name),
		orm.ToString(t.Description),
		orm.ToString(t.Identifier),
		orm.ToString(t.UnitCost),
		enumCostCenterValueTypeToString(&t.ValueType),
		enumCostCenterUnitCostTypeToString(&t.UnitType),
		orm.ToString(t.UnitCurrency),
		orm.ToString(t.Allocation),
		orm.ToString(t.CreatedAt),
		orm.ToString(t.UpdatedAt),
	)
}

// DBCreate will create a new CostCenter record in the database
func (t *CostCenter) DBCreate(ctx context.Context, db *sql.DB) (sql.Result, error) {
	q := "INSERT INTO `cost_center` (`cost_center`.`id`,`cost_center`.`checksum`,`cost_center`.`customer_id`,`cost_center`.`name`,`cost_center`.`description`,`cost_center`.`identifier`,`cost_center`.`unit_cost`,`cost_center`.`value_type`,`cost_center`.`unit_type`,`cost_center`.`unit_currency`,`cost_center`.`allocation`,`cost_center`.`created_at`,`cost_center`.`updated_at`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?)"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.Description),
		orm.ToSQLString(t.Identifier),
		orm.ToSQLFloat64(t.UnitCost),
		orm.ToSQLString(enumCostCenterValueTypeToString(&t.ValueType)),
		orm.ToSQLString(enumCostCenterUnitCostTypeToString(&t.UnitType)),
		orm.ToSQLString(t.UnitCurrency),
		orm.ToSQLFloat64(t.Allocation),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
	)
}

// DBCreateTx will create a new CostCenter record in the database using the provided transaction
func (t *CostCenter) DBCreateTx(ctx context.Context, tx *sql.Tx) (sql.Result, error) {
	q := "INSERT INTO `cost_center` (`cost_center`.`id`,`cost_center`.`checksum`,`cost_center`.`customer_id`,`cost_center`.`name`,`cost_center`.`description`,`cost_center`.`identifier`,`cost_center`.`unit_cost`,`cost_center`.`value_type`,`cost_center`.`unit_type`,`cost_center`.`unit_currency`,`cost_center`.`allocation`,`cost_center`.`created_at`,`cost_center`.`updated_at`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?)"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.Description),
		orm.ToSQLString(t.Identifier),
		orm.ToSQLFloat64(t.UnitCost),
		orm.ToSQLString(enumCostCenterValueTypeToString(&t.ValueType)),
		orm.ToSQLString(enumCostCenterUnitCostTypeToString(&t.UnitType)),
		orm.ToSQLString(t.UnitCurrency),
		orm.ToSQLFloat64(t.Allocation),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
	)
}

// DBCreateIgnoreDuplicate will upsert the CostCenter record in the database
func (t *CostCenter) DBCreateIgnoreDuplicate(ctx context.Context, db *sql.DB) (sql.Result, error) {
	q := "INSERT INTO `cost_center` (`cost_center`.`id`,`cost_center`.`checksum`,`cost_center`.`customer_id`,`cost_center`.`name`,`cost_center`.`description`,`cost_center`.`identifier`,`cost_center`.`unit_cost`,`cost_center`.`value_type`,`cost_center`.`unit_type`,`cost_center`.`unit_currency`,`cost_center`.`allocation`,`cost_center`.`created_at`,`cost_center`.`updated_at`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.Description),
		orm.ToSQLString(t.Identifier),
		orm.ToSQLFloat64(t.UnitCost),
		orm.ToSQLString(enumCostCenterValueTypeToString(&t.ValueType)),
		orm.ToSQLString(enumCostCenterUnitCostTypeToString(&t.UnitType)),
		orm.ToSQLString(t.UnitCurrency),
		orm.ToSQLFloat64(t.Allocation),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
	)
}

// DBCreateIgnoreDuplicateTx will upsert the CostCenter record in the database using the provided transaction
func (t *CostCenter) DBCreateIgnoreDuplicateTx(ctx context.Context, tx *sql.Tx) (sql.Result, error) {
	q := "INSERT INTO `cost_center` (`cost_center`.`id`,`cost_center`.`checksum`,`cost_center`.`customer_id`,`cost_center`.`name`,`cost_center`.`description`,`cost_center`.`identifier`,`cost_center`.`unit_cost`,`cost_center`.`value_type`,`cost_center`.`unit_type`,`cost_center`.`unit_currency`,`cost_center`.`allocation`,`cost_center`.`created_at`,`cost_center`.`updated_at`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.Description),
		orm.ToSQLString(t.Identifier),
		orm.ToSQLFloat64(t.UnitCost),
		orm.ToSQLString(enumCostCenterValueTypeToString(&t.ValueType)),
		orm.ToSQLString(enumCostCenterUnitCostTypeToString(&t.UnitType)),
		orm.ToSQLString(t.UnitCurrency),
		orm.ToSQLFloat64(t.Allocation),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
	)
}

// DeleteAllCostCenters deletes all CostCenter records in the database with optional filters
func DeleteAllCostCenters(ctx context.Context, db *sql.DB, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(CostCenterTableName),
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

// DeleteAllCostCentersTx deletes all CostCenter records in the database with optional filters using the provided transaction
func DeleteAllCostCentersTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(CostCenterTableName),
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

// DBDelete will delete this CostCenter record in the database
func (t *CostCenter) DBDelete(ctx context.Context, db *sql.DB) (bool, error) {
	q := "DELETE FROM `cost_center` WHERE `id` = ?"
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

// DBDeleteTx will delete this CostCenter record in the database using the provided transaction
func (t *CostCenter) DBDeleteTx(ctx context.Context, tx *sql.Tx) (bool, error) {
	q := "DELETE FROM `cost_center` WHERE `id` = ?"
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

// DBUpdate will update the CostCenter record in the database
func (t *CostCenter) DBUpdate(ctx context.Context, db *sql.DB) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `cost_center` SET `checksum`=?,`customer_id`=?,`name`=?,`description`=?,`identifier`=?,`unit_cost`=?,`value_type`=?,`unit_type`=?,`unit_currency`=?,`allocation`=?,`created_at`=?,`updated_at`=? WHERE `id`=?"
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.Description),
		orm.ToSQLString(t.Identifier),
		orm.ToSQLFloat64(t.UnitCost),
		orm.ToSQLString(enumCostCenterValueTypeToString(&t.ValueType)),
		orm.ToSQLString(enumCostCenterUnitCostTypeToString(&t.UnitType)),
		orm.ToSQLString(t.UnitCurrency),
		orm.ToSQLFloat64(t.Allocation),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
		orm.ToSQLString(t.ID),
	)
}

// DBUpdateTx will update the CostCenter record in the database using the provided transaction
func (t *CostCenter) DBUpdateTx(ctx context.Context, tx *sql.Tx) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `cost_center` SET `checksum`=?,`customer_id`=?,`name`=?,`description`=?,`identifier`=?,`unit_cost`=?,`value_type`=?,`unit_type`=?,`unit_currency`=?,`allocation`=?,`created_at`=?,`updated_at`=? WHERE `id`=?"
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.Description),
		orm.ToSQLString(t.Identifier),
		orm.ToSQLFloat64(t.UnitCost),
		orm.ToSQLString(enumCostCenterValueTypeToString(&t.ValueType)),
		orm.ToSQLString(enumCostCenterUnitCostTypeToString(&t.UnitType)),
		orm.ToSQLString(t.UnitCurrency),
		orm.ToSQLFloat64(t.Allocation),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
		orm.ToSQLString(t.ID),
	)
}

// DBUpsert will upsert the CostCenter record in the database
func (t *CostCenter) DBUpsert(ctx context.Context, db *sql.DB, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `cost_center` (`cost_center`.`id`,`cost_center`.`checksum`,`cost_center`.`customer_id`,`cost_center`.`name`,`cost_center`.`description`,`cost_center`.`identifier`,`cost_center`.`unit_cost`,`cost_center`.`value_type`,`cost_center`.`unit_type`,`cost_center`.`unit_currency`,`cost_center`.`allocation`,`cost_center`.`created_at`,`cost_center`.`updated_at`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `cost_center` (`cost_center`.`id`,`cost_center`.`checksum`,`cost_center`.`customer_id`,`cost_center`.`name`,`cost_center`.`description`,`cost_center`.`identifier`,`cost_center`.`unit_cost`,`cost_center`.`value_type`,`cost_center`.`unit_type`,`cost_center`.`unit_currency`,`cost_center`.`allocation`,`cost_center`.`created_at`,`cost_center`.`updated_at`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`customer_id`=VALUES(`customer_id`),`name`=VALUES(`name`),`description`=VALUES(`description`),`identifier`=VALUES(`identifier`),`unit_cost`=VALUES(`unit_cost`),`value_type`=VALUES(`value_type`),`unit_type`=VALUES(`unit_type`),`unit_currency`=VALUES(`unit_currency`),`allocation`=VALUES(`allocation`),`created_at`=VALUES(`created_at`),`updated_at`=VALUES(`updated_at`)"
	}
	r, err := db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.Description),
		orm.ToSQLString(t.Identifier),
		orm.ToSQLFloat64(t.UnitCost),
		orm.ToSQLString(enumCostCenterValueTypeToString(&t.ValueType)),
		orm.ToSQLString(enumCostCenterUnitCostTypeToString(&t.UnitType)),
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

// DBUpsertTx will upsert the CostCenter record in the database using the provided transaction
func (t *CostCenter) DBUpsertTx(ctx context.Context, tx *sql.Tx, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `cost_center` (`cost_center`.`id`,`cost_center`.`checksum`,`cost_center`.`customer_id`,`cost_center`.`name`,`cost_center`.`description`,`cost_center`.`identifier`,`cost_center`.`unit_cost`,`cost_center`.`value_type`,`cost_center`.`unit_type`,`cost_center`.`unit_currency`,`cost_center`.`allocation`,`cost_center`.`created_at`,`cost_center`.`updated_at`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `cost_center` (`cost_center`.`id`,`cost_center`.`checksum`,`cost_center`.`customer_id`,`cost_center`.`name`,`cost_center`.`description`,`cost_center`.`identifier`,`cost_center`.`unit_cost`,`cost_center`.`value_type`,`cost_center`.`unit_type`,`cost_center`.`unit_currency`,`cost_center`.`allocation`,`cost_center`.`created_at`,`cost_center`.`updated_at`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`customer_id`=VALUES(`customer_id`),`name`=VALUES(`name`),`description`=VALUES(`description`),`identifier`=VALUES(`identifier`),`unit_cost`=VALUES(`unit_cost`),`value_type`=VALUES(`value_type`),`unit_type`=VALUES(`unit_type`),`unit_currency`=VALUES(`unit_currency`),`allocation`=VALUES(`allocation`),`created_at`=VALUES(`created_at`),`updated_at`=VALUES(`updated_at`)"
	}
	r, err := tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.Description),
		orm.ToSQLString(t.Identifier),
		orm.ToSQLFloat64(t.UnitCost),
		orm.ToSQLString(enumCostCenterValueTypeToString(&t.ValueType)),
		orm.ToSQLString(enumCostCenterUnitCostTypeToString(&t.UnitType)),
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

// DBFindOne will find a CostCenter record in the database with the primary key
func (t *CostCenter) DBFindOne(ctx context.Context, db *sql.DB, value string) (bool, error) {
	q := "SELECT `cost_center`.`id`,`cost_center`.`checksum`,`cost_center`.`customer_id`,`cost_center`.`name`,`cost_center`.`description`,`cost_center`.`identifier`,`cost_center`.`unit_cost`,`cost_center`.`value_type`,`cost_center`.`unit_type`,`cost_center`.`unit_currency`,`cost_center`.`allocation`,`cost_center`.`created_at`,`cost_center`.`updated_at` FROM `cost_center` WHERE `id` = ? LIMIT 1"
	row := db.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CustomerID sql.NullString
	var _Name sql.NullString
	var _Description sql.NullString
	var _Identifier sql.NullString
	var _UnitCost sql.NullFloat64
	var _ValueType sql.NullString
	var _UnitType sql.NullString
	var _UnitCurrency sql.NullString
	var _Allocation sql.NullFloat64
	var _CreatedAt sql.NullInt64
	var _UpdatedAt sql.NullInt64
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_Name,
		&_Description,
		&_Identifier,
		&_UnitCost,
		&_ValueType,
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
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _Description.Valid {
		t.SetDescription(_Description.String)
	}
	if _Identifier.Valid {
		t.SetIdentifier(_Identifier.String)
	}
	if _UnitCost.Valid {
		t.SetUnitCost(_UnitCost.Float64)
	}
	if _ValueType.Valid {
		t.SetValueTypeString(_ValueType.String)
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

// DBFindOneTx will find a CostCenter record in the database with the primary key using the provided transaction
func (t *CostCenter) DBFindOneTx(ctx context.Context, tx *sql.Tx, value string) (bool, error) {
	q := "SELECT `cost_center`.`id`,`cost_center`.`checksum`,`cost_center`.`customer_id`,`cost_center`.`name`,`cost_center`.`description`,`cost_center`.`identifier`,`cost_center`.`unit_cost`,`cost_center`.`value_type`,`cost_center`.`unit_type`,`cost_center`.`unit_currency`,`cost_center`.`allocation`,`cost_center`.`created_at`,`cost_center`.`updated_at` FROM `cost_center` WHERE `id` = ? LIMIT 1"
	row := tx.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CustomerID sql.NullString
	var _Name sql.NullString
	var _Description sql.NullString
	var _Identifier sql.NullString
	var _UnitCost sql.NullFloat64
	var _ValueType sql.NullString
	var _UnitType sql.NullString
	var _UnitCurrency sql.NullString
	var _Allocation sql.NullFloat64
	var _CreatedAt sql.NullInt64
	var _UpdatedAt sql.NullInt64
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_Name,
		&_Description,
		&_Identifier,
		&_UnitCost,
		&_ValueType,
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
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _Description.Valid {
		t.SetDescription(_Description.String)
	}
	if _Identifier.Valid {
		t.SetIdentifier(_Identifier.String)
	}
	if _UnitCost.Valid {
		t.SetUnitCost(_UnitCost.Float64)
	}
	if _ValueType.Valid {
		t.SetValueTypeString(_ValueType.String)
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

// FindCostCenters will find a CostCenter record in the database with the provided parameters
func FindCostCenters(ctx context.Context, db *sql.DB, _params ...interface{}) ([]*CostCenter, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("customer_id"),
		orm.Column("name"),
		orm.Column("description"),
		orm.Column("identifier"),
		orm.Column("unit_cost"),
		orm.Column("value_type"),
		orm.Column("unit_type"),
		orm.Column("unit_currency"),
		orm.Column("allocation"),
		orm.Column("created_at"),
		orm.Column("updated_at"),
		orm.Table(CostCenterTableName),
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
	results := make([]*CostCenter, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _Name sql.NullString
		var _Description sql.NullString
		var _Identifier sql.NullString
		var _UnitCost sql.NullFloat64
		var _ValueType sql.NullString
		var _UnitType sql.NullString
		var _UnitCurrency sql.NullString
		var _Allocation sql.NullFloat64
		var _CreatedAt sql.NullInt64
		var _UpdatedAt sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_Name,
			&_Description,
			&_Identifier,
			&_UnitCost,
			&_ValueType,
			&_UnitType,
			&_UnitCurrency,
			&_Allocation,
			&_CreatedAt,
			&_UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		t := &CostCenter{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _Description.Valid {
			t.SetDescription(_Description.String)
		}
		if _Identifier.Valid {
			t.SetIdentifier(_Identifier.String)
		}
		if _UnitCost.Valid {
			t.SetUnitCost(_UnitCost.Float64)
		}
		if _ValueType.Valid {
			t.SetValueTypeString(_ValueType.String)
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

// FindCostCentersTx will find a CostCenter record in the database with the provided parameters using the provided transaction
func FindCostCentersTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) ([]*CostCenter, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("customer_id"),
		orm.Column("name"),
		orm.Column("description"),
		orm.Column("identifier"),
		orm.Column("unit_cost"),
		orm.Column("value_type"),
		orm.Column("unit_type"),
		orm.Column("unit_currency"),
		orm.Column("allocation"),
		orm.Column("created_at"),
		orm.Column("updated_at"),
		orm.Table(CostCenterTableName),
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
	results := make([]*CostCenter, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _Name sql.NullString
		var _Description sql.NullString
		var _Identifier sql.NullString
		var _UnitCost sql.NullFloat64
		var _ValueType sql.NullString
		var _UnitType sql.NullString
		var _UnitCurrency sql.NullString
		var _Allocation sql.NullFloat64
		var _CreatedAt sql.NullInt64
		var _UpdatedAt sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_Name,
			&_Description,
			&_Identifier,
			&_UnitCost,
			&_ValueType,
			&_UnitType,
			&_UnitCurrency,
			&_Allocation,
			&_CreatedAt,
			&_UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		t := &CostCenter{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _Description.Valid {
			t.SetDescription(_Description.String)
		}
		if _Identifier.Valid {
			t.SetIdentifier(_Identifier.String)
		}
		if _UnitCost.Valid {
			t.SetUnitCost(_UnitCost.Float64)
		}
		if _ValueType.Valid {
			t.SetValueTypeString(_ValueType.String)
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

// DBFind will find a CostCenter record in the database with the provided parameters
func (t *CostCenter) DBFind(ctx context.Context, db *sql.DB, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("customer_id"),
		orm.Column("name"),
		orm.Column("description"),
		orm.Column("identifier"),
		orm.Column("unit_cost"),
		orm.Column("value_type"),
		orm.Column("unit_type"),
		orm.Column("unit_currency"),
		orm.Column("allocation"),
		orm.Column("created_at"),
		orm.Column("updated_at"),
		orm.Table(CostCenterTableName),
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
	var _Name sql.NullString
	var _Description sql.NullString
	var _Identifier sql.NullString
	var _UnitCost sql.NullFloat64
	var _ValueType sql.NullString
	var _UnitType sql.NullString
	var _UnitCurrency sql.NullString
	var _Allocation sql.NullFloat64
	var _CreatedAt sql.NullInt64
	var _UpdatedAt sql.NullInt64
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_Name,
		&_Description,
		&_Identifier,
		&_UnitCost,
		&_ValueType,
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
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _Description.Valid {
		t.SetDescription(_Description.String)
	}
	if _Identifier.Valid {
		t.SetIdentifier(_Identifier.String)
	}
	if _UnitCost.Valid {
		t.SetUnitCost(_UnitCost.Float64)
	}
	if _ValueType.Valid {
		t.SetValueTypeString(_ValueType.String)
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

// DBFindTx will find a CostCenter record in the database with the provided parameters using the provided transaction
func (t *CostCenter) DBFindTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("customer_id"),
		orm.Column("name"),
		orm.Column("description"),
		orm.Column("identifier"),
		orm.Column("unit_cost"),
		orm.Column("value_type"),
		orm.Column("unit_type"),
		orm.Column("unit_currency"),
		orm.Column("allocation"),
		orm.Column("created_at"),
		orm.Column("updated_at"),
		orm.Table(CostCenterTableName),
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
	var _Name sql.NullString
	var _Description sql.NullString
	var _Identifier sql.NullString
	var _UnitCost sql.NullFloat64
	var _ValueType sql.NullString
	var _UnitType sql.NullString
	var _UnitCurrency sql.NullString
	var _Allocation sql.NullFloat64
	var _CreatedAt sql.NullInt64
	var _UpdatedAt sql.NullInt64
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_Name,
		&_Description,
		&_Identifier,
		&_UnitCost,
		&_ValueType,
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
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _Description.Valid {
		t.SetDescription(_Description.String)
	}
	if _Identifier.Valid {
		t.SetIdentifier(_Identifier.String)
	}
	if _UnitCost.Valid {
		t.SetUnitCost(_UnitCost.Float64)
	}
	if _ValueType.Valid {
		t.SetValueTypeString(_ValueType.String)
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

// CountCostCenters will find the count of CostCenter records in the database
func CountCostCenters(ctx context.Context, db *sql.DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(CostCenterTableName),
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

// CountCostCentersTx will find the count of CostCenter records in the database using the provided transaction
func CountCostCentersTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(CostCenterTableName),
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

// DBCount will find the count of CostCenter records in the database
func (t *CostCenter) DBCount(ctx context.Context, db *sql.DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(CostCenterTableName),
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

// DBCountTx will find the count of CostCenter records in the database using the provided transaction
func (t *CostCenter) DBCountTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(CostCenterTableName),
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

// DBExists will return true if the CostCenter record exists in the database
func (t *CostCenter) DBExists(ctx context.Context, db *sql.DB) (bool, error) {
	q := "SELECT `id` FROM `cost_center` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := db.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// DBExistsTx will return true if the CostCenter record exists in the database using the provided transaction
func (t *CostCenter) DBExistsTx(ctx context.Context, tx *sql.Tx) (bool, error) {
	q := "SELECT `id` FROM `cost_center` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := tx.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// PrimaryKeyColumn returns the column name for the primary key
func (t *CostCenter) PrimaryKeyColumn() string {
	return CostCenterColumnID
}

// PrimaryKeyColumnType returns the primary key column Go type as a string
func (t *CostCenter) PrimaryKeyColumnType() string {
	return "string"
}

// PrimaryKey returns the primary key column value
func (t *CostCenter) PrimaryKey() interface{} {
	return t.ID
}
