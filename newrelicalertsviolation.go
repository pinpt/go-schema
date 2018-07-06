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

var _ Model = (*NewrelicAlertsViolation)(nil)
var _ CSVWriter = (*NewrelicAlertsViolation)(nil)
var _ JSONWriter = (*NewrelicAlertsViolation)(nil)
var _ Checksum = (*NewrelicAlertsViolation)(nil)

// NewrelicAlertsViolationTableName is the name of the table in SQL
const NewrelicAlertsViolationTableName = "newrelic_alerts_violation"

var NewrelicAlertsViolationColumns = []string{
	"id",
	"checksum",
	"customer_id",
	"ext_id",
	"duration",
	"policy_name",
	"condition_name",
	"priority",
	"opened_at",
	"entity_product",
	"entity_type",
	"entity_group_id",
	"entity_id",
	"entity_name",
	"policy_ext_id",
	"policy_id",
	"condition_ext_id",
	"condition_id",
}

// NewrelicAlertsViolation table
type NewrelicAlertsViolation struct {
	Checksum       *string `json:"checksum,omitempty"`
	ConditionExtID int64   `json:"condition_ext_id"`
	ConditionID    string  `json:"condition_id"`
	ConditionName  string  `json:"condition_name"`
	CustomerID     string  `json:"customer_id"`
	Duration       int64   `json:"duration"`
	EntityGroupID  int64   `json:"entity_group_id"`
	EntityID       int64   `json:"entity_id"`
	EntityName     string  `json:"entity_name"`
	EntityProduct  string  `json:"entity_product"`
	EntityType     string  `json:"entity_type"`
	ExtID          int64   `json:"ext_id"`
	ID             string  `json:"id"`
	OpenedAt       int64   `json:"opened_at"`
	PolicyExtID    int64   `json:"policy_ext_id"`
	PolicyID       string  `json:"policy_id"`
	PolicyName     string  `json:"policy_name"`
	Priority       string  `json:"priority"`
}

// TableName returns the SQL table name for NewrelicAlertsViolation and satifies the Model interface
func (t *NewrelicAlertsViolation) TableName() string {
	return NewrelicAlertsViolationTableName
}

// ToCSV will serialize the NewrelicAlertsViolation instance to a CSV compatible array of strings
func (t *NewrelicAlertsViolation) ToCSV() []string {
	return []string{
		t.ID,
		t.CalculateChecksum(),
		t.CustomerID,
		toCSVString(t.ExtID),
		toCSVString(t.Duration),
		t.PolicyName,
		t.ConditionName,
		t.Priority,
		toCSVString(t.OpenedAt),
		t.EntityProduct,
		t.EntityType,
		toCSVString(t.EntityGroupID),
		toCSVString(t.EntityID),
		t.EntityName,
		toCSVString(t.PolicyExtID),
		t.PolicyID,
		toCSVString(t.ConditionExtID),
		t.ConditionID,
	}
}

// WriteCSV will serialize the NewrelicAlertsViolation instance to the writer as CSV and satisfies the CSVWriter interface
func (t *NewrelicAlertsViolation) WriteCSV(w *csv.Writer) error {
	return w.Write(t.ToCSV())
}

// WriteJSON will serialize the NewrelicAlertsViolation instance to the writer as JSON and satisfies the JSONWriter interface
func (t *NewrelicAlertsViolation) WriteJSON(w io.Writer, indent ...bool) error {
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

// NewNewrelicAlertsViolationReader creates a JSON reader which can read in NewrelicAlertsViolation objects serialized as JSON either as an array, single object or json new lines
// and writes each NewrelicAlertsViolation to the channel provided
func NewNewrelicAlertsViolationReader(r io.Reader, ch chan<- NewrelicAlertsViolation) error {
	return orm.Deserialize(r, func(buf json.RawMessage) error {
		dec := json.NewDecoder(bytes.NewBuffer(buf))
		e := NewrelicAlertsViolation{}
		if err := dec.Decode(&e); err != nil {
			return err
		}
		ch <- e
		return nil
	})
}

// NewCSVNewrelicAlertsViolationReaderDir will read the reader as CSV and emit each record to the channel provided
func NewCSVNewrelicAlertsViolationReader(r io.Reader, ch chan<- NewrelicAlertsViolation) error {
	cr := csv.NewReader(r)
	for {
		record, err := cr.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		ch <- NewrelicAlertsViolation{
			ID:             record[0],
			Checksum:       fromStringPointer(record[1]),
			CustomerID:     record[2],
			ExtID:          fromCSVInt64(record[3]),
			Duration:       fromCSVInt64(record[4]),
			PolicyName:     record[5],
			ConditionName:  record[6],
			Priority:       record[7],
			OpenedAt:       fromCSVInt64(record[8]),
			EntityProduct:  record[9],
			EntityType:     record[10],
			EntityGroupID:  fromCSVInt64(record[11]),
			EntityID:       fromCSVInt64(record[12]),
			EntityName:     record[13],
			PolicyExtID:    fromCSVInt64(record[14]),
			PolicyID:       record[15],
			ConditionExtID: fromCSVInt64(record[16]),
			ConditionID:    record[17],
		}
	}
	return nil
}

// NewCSVNewrelicAlertsViolationReaderFile will read the file as a CSV and emit each record to the channel provided
func NewCSVNewrelicAlertsViolationReaderFile(fp string, ch chan<- NewrelicAlertsViolation) error {
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
	return NewCSVNewrelicAlertsViolationReader(fc, ch)
}

// NewCSVNewrelicAlertsViolationReaderDir will read the newrelic_alerts_violation.csv.gz file as a CSV and emit each record to the channel provided
func NewCSVNewrelicAlertsViolationReaderDir(dir string, ch chan<- NewrelicAlertsViolation) error {
	return NewCSVNewrelicAlertsViolationReaderFile(filepath.Join(dir, "newrelic_alerts_violation.csv.gz"), ch)
}

// NewrelicAlertsViolationCSVDeduper is a function callback which takes the existing value (a) and the new value (b)
// and the return value should be the one to use (or a new one, if applicable). return nil
// to skip processing of this record
type NewrelicAlertsViolationCSVDeduper func(a NewrelicAlertsViolation, b NewrelicAlertsViolation) *NewrelicAlertsViolation

// NewrelicAlertsViolationCSVDedupeDisabled is set on whether the CSV writer should de-dupe values by key
var NewrelicAlertsViolationCSVDedupeDisabled bool

// NewNewrelicAlertsViolationCSVWriterSize creates a batch writer that will write each NewrelicAlertsViolation into a CSV file
// this method will automatically de-duplicate entries using the primary key. if the checksum
// for a newer item with the same primary key doesn't match a previously sent item, the newer
// one will be used
func NewNewrelicAlertsViolationCSVWriterSize(w io.Writer, size int, dedupers ...NewrelicAlertsViolationCSVDeduper) (chan NewrelicAlertsViolation, chan bool, error) {
	cw := csv.NewWriter(w)
	ch := make(chan NewrelicAlertsViolation, size)
	done := make(chan bool)
	go func() {
		defer func() { done <- true }()
		dodedupe := !NewrelicAlertsViolationCSVDedupeDisabled
		var kv map[string]*NewrelicAlertsViolation
		var deduper NewrelicAlertsViolationCSVDeduper
		if dedupers != nil && len(dedupers) > 0 {
			deduper = dedupers[0]
			dodedupe = true
		}
		if dodedupe {
			kv = make(map[string]*NewrelicAlertsViolation)
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

// NewrelicAlertsViolationCSVDefaultSize is the default channel buffer size if not provided
var NewrelicAlertsViolationCSVDefaultSize = 100

// NewNewrelicAlertsViolationCSVWriter creates a batch writer that will write each NewrelicAlertsViolation into a CSV file
func NewNewrelicAlertsViolationCSVWriter(w io.Writer, dedupers ...NewrelicAlertsViolationCSVDeduper) (chan NewrelicAlertsViolation, chan bool, error) {
	return NewNewrelicAlertsViolationCSVWriterSize(w, NewrelicAlertsViolationCSVDefaultSize, dedupers...)
}

// NewNewrelicAlertsViolationCSVWriterDir creates a batch writer that will write each NewrelicAlertsViolation into a CSV file named newrelic_alerts_violation.csv.gz in dir
func NewNewrelicAlertsViolationCSVWriterDir(dir string, dedupers ...NewrelicAlertsViolationCSVDeduper) (chan NewrelicAlertsViolation, chan bool, error) {
	return NewNewrelicAlertsViolationCSVWriterFile(filepath.Join(dir, "newrelic_alerts_violation.csv.gz"), dedupers...)
}

// NewNewrelicAlertsViolationCSVWriterFile creates a batch writer that will write each NewrelicAlertsViolation into a CSV file
func NewNewrelicAlertsViolationCSVWriterFile(fn string, dedupers ...NewrelicAlertsViolationCSVDeduper) (chan NewrelicAlertsViolation, chan bool, error) {
	f, err := os.Create(fn)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening CSV file %s. %v", fn, err)
	}
	var fc io.WriteCloser = f
	if filepath.Ext(fn) == ".gz" {
		w, _ := gzip.NewWriterLevel(f, gzip.BestCompression)
		fc = w
	}
	ch, done, err := NewNewrelicAlertsViolationCSVWriter(fc, dedupers...)
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

type NewrelicAlertsViolationDBAction func(ctx context.Context, db *sql.DB, record NewrelicAlertsViolation) error

// NewNewrelicAlertsViolationDBWriterSize creates a DB writer that will write each issue into the DB
func NewNewrelicAlertsViolationDBWriterSize(ctx context.Context, db *sql.DB, errors chan<- error, size int, actions ...NewrelicAlertsViolationDBAction) (chan NewrelicAlertsViolation, chan bool, error) {
	ch := make(chan NewrelicAlertsViolation, size)
	done := make(chan bool)
	var action NewrelicAlertsViolationDBAction
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

// NewNewrelicAlertsViolationDBWriter creates a DB writer that will write each issue into the DB
func NewNewrelicAlertsViolationDBWriter(ctx context.Context, db *sql.DB, errors chan<- error, actions ...NewrelicAlertsViolationDBAction) (chan NewrelicAlertsViolation, chan bool, error) {
	return NewNewrelicAlertsViolationDBWriterSize(ctx, db, errors, 100, actions...)
}

// NewrelicAlertsViolationColumnID is the ID SQL column name for the NewrelicAlertsViolation table
const NewrelicAlertsViolationColumnID = "id"

// NewrelicAlertsViolationEscapedColumnID is the escaped ID SQL column name for the NewrelicAlertsViolation table
const NewrelicAlertsViolationEscapedColumnID = "`id`"

// NewrelicAlertsViolationColumnChecksum is the Checksum SQL column name for the NewrelicAlertsViolation table
const NewrelicAlertsViolationColumnChecksum = "checksum"

// NewrelicAlertsViolationEscapedColumnChecksum is the escaped Checksum SQL column name for the NewrelicAlertsViolation table
const NewrelicAlertsViolationEscapedColumnChecksum = "`checksum`"

// NewrelicAlertsViolationColumnCustomerID is the CustomerID SQL column name for the NewrelicAlertsViolation table
const NewrelicAlertsViolationColumnCustomerID = "customer_id"

// NewrelicAlertsViolationEscapedColumnCustomerID is the escaped CustomerID SQL column name for the NewrelicAlertsViolation table
const NewrelicAlertsViolationEscapedColumnCustomerID = "`customer_id`"

// NewrelicAlertsViolationColumnExtID is the ExtID SQL column name for the NewrelicAlertsViolation table
const NewrelicAlertsViolationColumnExtID = "ext_id"

// NewrelicAlertsViolationEscapedColumnExtID is the escaped ExtID SQL column name for the NewrelicAlertsViolation table
const NewrelicAlertsViolationEscapedColumnExtID = "`ext_id`"

// NewrelicAlertsViolationColumnDuration is the Duration SQL column name for the NewrelicAlertsViolation table
const NewrelicAlertsViolationColumnDuration = "duration"

// NewrelicAlertsViolationEscapedColumnDuration is the escaped Duration SQL column name for the NewrelicAlertsViolation table
const NewrelicAlertsViolationEscapedColumnDuration = "`duration`"

// NewrelicAlertsViolationColumnPolicyName is the PolicyName SQL column name for the NewrelicAlertsViolation table
const NewrelicAlertsViolationColumnPolicyName = "policy_name"

// NewrelicAlertsViolationEscapedColumnPolicyName is the escaped PolicyName SQL column name for the NewrelicAlertsViolation table
const NewrelicAlertsViolationEscapedColumnPolicyName = "`policy_name`"

// NewrelicAlertsViolationColumnConditionName is the ConditionName SQL column name for the NewrelicAlertsViolation table
const NewrelicAlertsViolationColumnConditionName = "condition_name"

// NewrelicAlertsViolationEscapedColumnConditionName is the escaped ConditionName SQL column name for the NewrelicAlertsViolation table
const NewrelicAlertsViolationEscapedColumnConditionName = "`condition_name`"

// NewrelicAlertsViolationColumnPriority is the Priority SQL column name for the NewrelicAlertsViolation table
const NewrelicAlertsViolationColumnPriority = "priority"

// NewrelicAlertsViolationEscapedColumnPriority is the escaped Priority SQL column name for the NewrelicAlertsViolation table
const NewrelicAlertsViolationEscapedColumnPriority = "`priority`"

// NewrelicAlertsViolationColumnOpenedAt is the OpenedAt SQL column name for the NewrelicAlertsViolation table
const NewrelicAlertsViolationColumnOpenedAt = "opened_at"

// NewrelicAlertsViolationEscapedColumnOpenedAt is the escaped OpenedAt SQL column name for the NewrelicAlertsViolation table
const NewrelicAlertsViolationEscapedColumnOpenedAt = "`opened_at`"

// NewrelicAlertsViolationColumnEntityProduct is the EntityProduct SQL column name for the NewrelicAlertsViolation table
const NewrelicAlertsViolationColumnEntityProduct = "entity_product"

// NewrelicAlertsViolationEscapedColumnEntityProduct is the escaped EntityProduct SQL column name for the NewrelicAlertsViolation table
const NewrelicAlertsViolationEscapedColumnEntityProduct = "`entity_product`"

// NewrelicAlertsViolationColumnEntityType is the EntityType SQL column name for the NewrelicAlertsViolation table
const NewrelicAlertsViolationColumnEntityType = "entity_type"

// NewrelicAlertsViolationEscapedColumnEntityType is the escaped EntityType SQL column name for the NewrelicAlertsViolation table
const NewrelicAlertsViolationEscapedColumnEntityType = "`entity_type`"

// NewrelicAlertsViolationColumnEntityGroupID is the EntityGroupID SQL column name for the NewrelicAlertsViolation table
const NewrelicAlertsViolationColumnEntityGroupID = "entity_group_id"

// NewrelicAlertsViolationEscapedColumnEntityGroupID is the escaped EntityGroupID SQL column name for the NewrelicAlertsViolation table
const NewrelicAlertsViolationEscapedColumnEntityGroupID = "`entity_group_id`"

// NewrelicAlertsViolationColumnEntityID is the EntityID SQL column name for the NewrelicAlertsViolation table
const NewrelicAlertsViolationColumnEntityID = "entity_id"

// NewrelicAlertsViolationEscapedColumnEntityID is the escaped EntityID SQL column name for the NewrelicAlertsViolation table
const NewrelicAlertsViolationEscapedColumnEntityID = "`entity_id`"

// NewrelicAlertsViolationColumnEntityName is the EntityName SQL column name for the NewrelicAlertsViolation table
const NewrelicAlertsViolationColumnEntityName = "entity_name"

// NewrelicAlertsViolationEscapedColumnEntityName is the escaped EntityName SQL column name for the NewrelicAlertsViolation table
const NewrelicAlertsViolationEscapedColumnEntityName = "`entity_name`"

// NewrelicAlertsViolationColumnPolicyExtID is the PolicyExtID SQL column name for the NewrelicAlertsViolation table
const NewrelicAlertsViolationColumnPolicyExtID = "policy_ext_id"

// NewrelicAlertsViolationEscapedColumnPolicyExtID is the escaped PolicyExtID SQL column name for the NewrelicAlertsViolation table
const NewrelicAlertsViolationEscapedColumnPolicyExtID = "`policy_ext_id`"

// NewrelicAlertsViolationColumnPolicyID is the PolicyID SQL column name for the NewrelicAlertsViolation table
const NewrelicAlertsViolationColumnPolicyID = "policy_id"

// NewrelicAlertsViolationEscapedColumnPolicyID is the escaped PolicyID SQL column name for the NewrelicAlertsViolation table
const NewrelicAlertsViolationEscapedColumnPolicyID = "`policy_id`"

// NewrelicAlertsViolationColumnConditionExtID is the ConditionExtID SQL column name for the NewrelicAlertsViolation table
const NewrelicAlertsViolationColumnConditionExtID = "condition_ext_id"

// NewrelicAlertsViolationEscapedColumnConditionExtID is the escaped ConditionExtID SQL column name for the NewrelicAlertsViolation table
const NewrelicAlertsViolationEscapedColumnConditionExtID = "`condition_ext_id`"

// NewrelicAlertsViolationColumnConditionID is the ConditionID SQL column name for the NewrelicAlertsViolation table
const NewrelicAlertsViolationColumnConditionID = "condition_id"

// NewrelicAlertsViolationEscapedColumnConditionID is the escaped ConditionID SQL column name for the NewrelicAlertsViolation table
const NewrelicAlertsViolationEscapedColumnConditionID = "`condition_id`"

// GetID will return the NewrelicAlertsViolation ID value
func (t *NewrelicAlertsViolation) GetID() string {
	return t.ID
}

// SetID will set the NewrelicAlertsViolation ID value
func (t *NewrelicAlertsViolation) SetID(v string) {
	t.ID = v
}

// FindNewrelicAlertsViolationByID will find a NewrelicAlertsViolation by ID
func FindNewrelicAlertsViolationByID(ctx context.Context, db *sql.DB, value string) (*NewrelicAlertsViolation, error) {
	q := "SELECT `newrelic_alerts_violation`.`id`,`newrelic_alerts_violation`.`checksum`,`newrelic_alerts_violation`.`customer_id`,`newrelic_alerts_violation`.`ext_id`,`newrelic_alerts_violation`.`duration`,`newrelic_alerts_violation`.`policy_name`,`newrelic_alerts_violation`.`condition_name`,`newrelic_alerts_violation`.`priority`,`newrelic_alerts_violation`.`opened_at`,`newrelic_alerts_violation`.`entity_product`,`newrelic_alerts_violation`.`entity_type`,`newrelic_alerts_violation`.`entity_group_id`,`newrelic_alerts_violation`.`entity_id`,`newrelic_alerts_violation`.`entity_name`,`newrelic_alerts_violation`.`policy_ext_id`,`newrelic_alerts_violation`.`policy_id`,`newrelic_alerts_violation`.`condition_ext_id`,`newrelic_alerts_violation`.`condition_id` FROM `newrelic_alerts_violation` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CustomerID sql.NullString
	var _ExtID sql.NullInt64
	var _Duration sql.NullInt64
	var _PolicyName sql.NullString
	var _ConditionName sql.NullString
	var _Priority sql.NullString
	var _OpenedAt sql.NullInt64
	var _EntityProduct sql.NullString
	var _EntityType sql.NullString
	var _EntityGroupID sql.NullInt64
	var _EntityID sql.NullInt64
	var _EntityName sql.NullString
	var _PolicyExtID sql.NullInt64
	var _PolicyID sql.NullString
	var _ConditionExtID sql.NullInt64
	var _ConditionID sql.NullString
	err := db.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_ExtID,
		&_Duration,
		&_PolicyName,
		&_ConditionName,
		&_Priority,
		&_OpenedAt,
		&_EntityProduct,
		&_EntityType,
		&_EntityGroupID,
		&_EntityID,
		&_EntityName,
		&_PolicyExtID,
		&_PolicyID,
		&_ConditionExtID,
		&_ConditionID,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &NewrelicAlertsViolation{}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _Checksum.Valid {
		t.SetChecksum(_Checksum.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _ExtID.Valid {
		t.SetExtID(_ExtID.Int64)
	}
	if _Duration.Valid {
		t.SetDuration(_Duration.Int64)
	}
	if _PolicyName.Valid {
		t.SetPolicyName(_PolicyName.String)
	}
	if _ConditionName.Valid {
		t.SetConditionName(_ConditionName.String)
	}
	if _Priority.Valid {
		t.SetPriority(_Priority.String)
	}
	if _OpenedAt.Valid {
		t.SetOpenedAt(_OpenedAt.Int64)
	}
	if _EntityProduct.Valid {
		t.SetEntityProduct(_EntityProduct.String)
	}
	if _EntityType.Valid {
		t.SetEntityType(_EntityType.String)
	}
	if _EntityGroupID.Valid {
		t.SetEntityGroupID(_EntityGroupID.Int64)
	}
	if _EntityID.Valid {
		t.SetEntityID(_EntityID.Int64)
	}
	if _EntityName.Valid {
		t.SetEntityName(_EntityName.String)
	}
	if _PolicyExtID.Valid {
		t.SetPolicyExtID(_PolicyExtID.Int64)
	}
	if _PolicyID.Valid {
		t.SetPolicyID(_PolicyID.String)
	}
	if _ConditionExtID.Valid {
		t.SetConditionExtID(_ConditionExtID.Int64)
	}
	if _ConditionID.Valid {
		t.SetConditionID(_ConditionID.String)
	}
	return t, nil
}

// FindNewrelicAlertsViolationByIDTx will find a NewrelicAlertsViolation by ID using the provided transaction
func FindNewrelicAlertsViolationByIDTx(ctx context.Context, tx *sql.Tx, value string) (*NewrelicAlertsViolation, error) {
	q := "SELECT `newrelic_alerts_violation`.`id`,`newrelic_alerts_violation`.`checksum`,`newrelic_alerts_violation`.`customer_id`,`newrelic_alerts_violation`.`ext_id`,`newrelic_alerts_violation`.`duration`,`newrelic_alerts_violation`.`policy_name`,`newrelic_alerts_violation`.`condition_name`,`newrelic_alerts_violation`.`priority`,`newrelic_alerts_violation`.`opened_at`,`newrelic_alerts_violation`.`entity_product`,`newrelic_alerts_violation`.`entity_type`,`newrelic_alerts_violation`.`entity_group_id`,`newrelic_alerts_violation`.`entity_id`,`newrelic_alerts_violation`.`entity_name`,`newrelic_alerts_violation`.`policy_ext_id`,`newrelic_alerts_violation`.`policy_id`,`newrelic_alerts_violation`.`condition_ext_id`,`newrelic_alerts_violation`.`condition_id` FROM `newrelic_alerts_violation` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CustomerID sql.NullString
	var _ExtID sql.NullInt64
	var _Duration sql.NullInt64
	var _PolicyName sql.NullString
	var _ConditionName sql.NullString
	var _Priority sql.NullString
	var _OpenedAt sql.NullInt64
	var _EntityProduct sql.NullString
	var _EntityType sql.NullString
	var _EntityGroupID sql.NullInt64
	var _EntityID sql.NullInt64
	var _EntityName sql.NullString
	var _PolicyExtID sql.NullInt64
	var _PolicyID sql.NullString
	var _ConditionExtID sql.NullInt64
	var _ConditionID sql.NullString
	err := tx.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_ExtID,
		&_Duration,
		&_PolicyName,
		&_ConditionName,
		&_Priority,
		&_OpenedAt,
		&_EntityProduct,
		&_EntityType,
		&_EntityGroupID,
		&_EntityID,
		&_EntityName,
		&_PolicyExtID,
		&_PolicyID,
		&_ConditionExtID,
		&_ConditionID,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &NewrelicAlertsViolation{}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _Checksum.Valid {
		t.SetChecksum(_Checksum.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _ExtID.Valid {
		t.SetExtID(_ExtID.Int64)
	}
	if _Duration.Valid {
		t.SetDuration(_Duration.Int64)
	}
	if _PolicyName.Valid {
		t.SetPolicyName(_PolicyName.String)
	}
	if _ConditionName.Valid {
		t.SetConditionName(_ConditionName.String)
	}
	if _Priority.Valid {
		t.SetPriority(_Priority.String)
	}
	if _OpenedAt.Valid {
		t.SetOpenedAt(_OpenedAt.Int64)
	}
	if _EntityProduct.Valid {
		t.SetEntityProduct(_EntityProduct.String)
	}
	if _EntityType.Valid {
		t.SetEntityType(_EntityType.String)
	}
	if _EntityGroupID.Valid {
		t.SetEntityGroupID(_EntityGroupID.Int64)
	}
	if _EntityID.Valid {
		t.SetEntityID(_EntityID.Int64)
	}
	if _EntityName.Valid {
		t.SetEntityName(_EntityName.String)
	}
	if _PolicyExtID.Valid {
		t.SetPolicyExtID(_PolicyExtID.Int64)
	}
	if _PolicyID.Valid {
		t.SetPolicyID(_PolicyID.String)
	}
	if _ConditionExtID.Valid {
		t.SetConditionExtID(_ConditionExtID.Int64)
	}
	if _ConditionID.Valid {
		t.SetConditionID(_ConditionID.String)
	}
	return t, nil
}

// GetChecksum will return the NewrelicAlertsViolation Checksum value
func (t *NewrelicAlertsViolation) GetChecksum() string {
	if t.Checksum == nil {
		return ""
	}
	return *t.Checksum
}

// SetChecksum will set the NewrelicAlertsViolation Checksum value
func (t *NewrelicAlertsViolation) SetChecksum(v string) {
	t.Checksum = &v
}

// GetCustomerID will return the NewrelicAlertsViolation CustomerID value
func (t *NewrelicAlertsViolation) GetCustomerID() string {
	return t.CustomerID
}

// SetCustomerID will set the NewrelicAlertsViolation CustomerID value
func (t *NewrelicAlertsViolation) SetCustomerID(v string) {
	t.CustomerID = v
}

// FindNewrelicAlertsViolationsByCustomerID will find all NewrelicAlertsViolations by the CustomerID value
func FindNewrelicAlertsViolationsByCustomerID(ctx context.Context, db *sql.DB, value string) ([]*NewrelicAlertsViolation, error) {
	q := "SELECT `newrelic_alerts_violation`.`id`,`newrelic_alerts_violation`.`checksum`,`newrelic_alerts_violation`.`customer_id`,`newrelic_alerts_violation`.`ext_id`,`newrelic_alerts_violation`.`duration`,`newrelic_alerts_violation`.`policy_name`,`newrelic_alerts_violation`.`condition_name`,`newrelic_alerts_violation`.`priority`,`newrelic_alerts_violation`.`opened_at`,`newrelic_alerts_violation`.`entity_product`,`newrelic_alerts_violation`.`entity_type`,`newrelic_alerts_violation`.`entity_group_id`,`newrelic_alerts_violation`.`entity_id`,`newrelic_alerts_violation`.`entity_name`,`newrelic_alerts_violation`.`policy_ext_id`,`newrelic_alerts_violation`.`policy_id`,`newrelic_alerts_violation`.`condition_ext_id`,`newrelic_alerts_violation`.`condition_id` FROM `newrelic_alerts_violation` WHERE `customer_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*NewrelicAlertsViolation, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _ExtID sql.NullInt64
		var _Duration sql.NullInt64
		var _PolicyName sql.NullString
		var _ConditionName sql.NullString
		var _Priority sql.NullString
		var _OpenedAt sql.NullInt64
		var _EntityProduct sql.NullString
		var _EntityType sql.NullString
		var _EntityGroupID sql.NullInt64
		var _EntityID sql.NullInt64
		var _EntityName sql.NullString
		var _PolicyExtID sql.NullInt64
		var _PolicyID sql.NullString
		var _ConditionExtID sql.NullInt64
		var _ConditionID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_ExtID,
			&_Duration,
			&_PolicyName,
			&_ConditionName,
			&_Priority,
			&_OpenedAt,
			&_EntityProduct,
			&_EntityType,
			&_EntityGroupID,
			&_EntityID,
			&_EntityName,
			&_PolicyExtID,
			&_PolicyID,
			&_ConditionExtID,
			&_ConditionID,
		)
		if err != nil {
			return nil, err
		}
		t := &NewrelicAlertsViolation{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _ExtID.Valid {
			t.SetExtID(_ExtID.Int64)
		}
		if _Duration.Valid {
			t.SetDuration(_Duration.Int64)
		}
		if _PolicyName.Valid {
			t.SetPolicyName(_PolicyName.String)
		}
		if _ConditionName.Valid {
			t.SetConditionName(_ConditionName.String)
		}
		if _Priority.Valid {
			t.SetPriority(_Priority.String)
		}
		if _OpenedAt.Valid {
			t.SetOpenedAt(_OpenedAt.Int64)
		}
		if _EntityProduct.Valid {
			t.SetEntityProduct(_EntityProduct.String)
		}
		if _EntityType.Valid {
			t.SetEntityType(_EntityType.String)
		}
		if _EntityGroupID.Valid {
			t.SetEntityGroupID(_EntityGroupID.Int64)
		}
		if _EntityID.Valid {
			t.SetEntityID(_EntityID.Int64)
		}
		if _EntityName.Valid {
			t.SetEntityName(_EntityName.String)
		}
		if _PolicyExtID.Valid {
			t.SetPolicyExtID(_PolicyExtID.Int64)
		}
		if _PolicyID.Valid {
			t.SetPolicyID(_PolicyID.String)
		}
		if _ConditionExtID.Valid {
			t.SetConditionExtID(_ConditionExtID.Int64)
		}
		if _ConditionID.Valid {
			t.SetConditionID(_ConditionID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindNewrelicAlertsViolationsByCustomerIDTx will find all NewrelicAlertsViolations by the CustomerID value using the provided transaction
func FindNewrelicAlertsViolationsByCustomerIDTx(ctx context.Context, tx *sql.Tx, value string) ([]*NewrelicAlertsViolation, error) {
	q := "SELECT `newrelic_alerts_violation`.`id`,`newrelic_alerts_violation`.`checksum`,`newrelic_alerts_violation`.`customer_id`,`newrelic_alerts_violation`.`ext_id`,`newrelic_alerts_violation`.`duration`,`newrelic_alerts_violation`.`policy_name`,`newrelic_alerts_violation`.`condition_name`,`newrelic_alerts_violation`.`priority`,`newrelic_alerts_violation`.`opened_at`,`newrelic_alerts_violation`.`entity_product`,`newrelic_alerts_violation`.`entity_type`,`newrelic_alerts_violation`.`entity_group_id`,`newrelic_alerts_violation`.`entity_id`,`newrelic_alerts_violation`.`entity_name`,`newrelic_alerts_violation`.`policy_ext_id`,`newrelic_alerts_violation`.`policy_id`,`newrelic_alerts_violation`.`condition_ext_id`,`newrelic_alerts_violation`.`condition_id` FROM `newrelic_alerts_violation` WHERE `customer_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*NewrelicAlertsViolation, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _ExtID sql.NullInt64
		var _Duration sql.NullInt64
		var _PolicyName sql.NullString
		var _ConditionName sql.NullString
		var _Priority sql.NullString
		var _OpenedAt sql.NullInt64
		var _EntityProduct sql.NullString
		var _EntityType sql.NullString
		var _EntityGroupID sql.NullInt64
		var _EntityID sql.NullInt64
		var _EntityName sql.NullString
		var _PolicyExtID sql.NullInt64
		var _PolicyID sql.NullString
		var _ConditionExtID sql.NullInt64
		var _ConditionID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_ExtID,
			&_Duration,
			&_PolicyName,
			&_ConditionName,
			&_Priority,
			&_OpenedAt,
			&_EntityProduct,
			&_EntityType,
			&_EntityGroupID,
			&_EntityID,
			&_EntityName,
			&_PolicyExtID,
			&_PolicyID,
			&_ConditionExtID,
			&_ConditionID,
		)
		if err != nil {
			return nil, err
		}
		t := &NewrelicAlertsViolation{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _ExtID.Valid {
			t.SetExtID(_ExtID.Int64)
		}
		if _Duration.Valid {
			t.SetDuration(_Duration.Int64)
		}
		if _PolicyName.Valid {
			t.SetPolicyName(_PolicyName.String)
		}
		if _ConditionName.Valid {
			t.SetConditionName(_ConditionName.String)
		}
		if _Priority.Valid {
			t.SetPriority(_Priority.String)
		}
		if _OpenedAt.Valid {
			t.SetOpenedAt(_OpenedAt.Int64)
		}
		if _EntityProduct.Valid {
			t.SetEntityProduct(_EntityProduct.String)
		}
		if _EntityType.Valid {
			t.SetEntityType(_EntityType.String)
		}
		if _EntityGroupID.Valid {
			t.SetEntityGroupID(_EntityGroupID.Int64)
		}
		if _EntityID.Valid {
			t.SetEntityID(_EntityID.Int64)
		}
		if _EntityName.Valid {
			t.SetEntityName(_EntityName.String)
		}
		if _PolicyExtID.Valid {
			t.SetPolicyExtID(_PolicyExtID.Int64)
		}
		if _PolicyID.Valid {
			t.SetPolicyID(_PolicyID.String)
		}
		if _ConditionExtID.Valid {
			t.SetConditionExtID(_ConditionExtID.Int64)
		}
		if _ConditionID.Valid {
			t.SetConditionID(_ConditionID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetExtID will return the NewrelicAlertsViolation ExtID value
func (t *NewrelicAlertsViolation) GetExtID() int64 {
	return t.ExtID
}

// SetExtID will set the NewrelicAlertsViolation ExtID value
func (t *NewrelicAlertsViolation) SetExtID(v int64) {
	t.ExtID = v
}

// GetDuration will return the NewrelicAlertsViolation Duration value
func (t *NewrelicAlertsViolation) GetDuration() int64 {
	return t.Duration
}

// SetDuration will set the NewrelicAlertsViolation Duration value
func (t *NewrelicAlertsViolation) SetDuration(v int64) {
	t.Duration = v
}

// GetPolicyName will return the NewrelicAlertsViolation PolicyName value
func (t *NewrelicAlertsViolation) GetPolicyName() string {
	return t.PolicyName
}

// SetPolicyName will set the NewrelicAlertsViolation PolicyName value
func (t *NewrelicAlertsViolation) SetPolicyName(v string) {
	t.PolicyName = v
}

// GetConditionName will return the NewrelicAlertsViolation ConditionName value
func (t *NewrelicAlertsViolation) GetConditionName() string {
	return t.ConditionName
}

// SetConditionName will set the NewrelicAlertsViolation ConditionName value
func (t *NewrelicAlertsViolation) SetConditionName(v string) {
	t.ConditionName = v
}

// GetPriority will return the NewrelicAlertsViolation Priority value
func (t *NewrelicAlertsViolation) GetPriority() string {
	return t.Priority
}

// SetPriority will set the NewrelicAlertsViolation Priority value
func (t *NewrelicAlertsViolation) SetPriority(v string) {
	t.Priority = v
}

// GetOpenedAt will return the NewrelicAlertsViolation OpenedAt value
func (t *NewrelicAlertsViolation) GetOpenedAt() int64 {
	return t.OpenedAt
}

// SetOpenedAt will set the NewrelicAlertsViolation OpenedAt value
func (t *NewrelicAlertsViolation) SetOpenedAt(v int64) {
	t.OpenedAt = v
}

// GetEntityProduct will return the NewrelicAlertsViolation EntityProduct value
func (t *NewrelicAlertsViolation) GetEntityProduct() string {
	return t.EntityProduct
}

// SetEntityProduct will set the NewrelicAlertsViolation EntityProduct value
func (t *NewrelicAlertsViolation) SetEntityProduct(v string) {
	t.EntityProduct = v
}

// GetEntityType will return the NewrelicAlertsViolation EntityType value
func (t *NewrelicAlertsViolation) GetEntityType() string {
	return t.EntityType
}

// SetEntityType will set the NewrelicAlertsViolation EntityType value
func (t *NewrelicAlertsViolation) SetEntityType(v string) {
	t.EntityType = v
}

// GetEntityGroupID will return the NewrelicAlertsViolation EntityGroupID value
func (t *NewrelicAlertsViolation) GetEntityGroupID() int64 {
	return t.EntityGroupID
}

// SetEntityGroupID will set the NewrelicAlertsViolation EntityGroupID value
func (t *NewrelicAlertsViolation) SetEntityGroupID(v int64) {
	t.EntityGroupID = v
}

// GetEntityID will return the NewrelicAlertsViolation EntityID value
func (t *NewrelicAlertsViolation) GetEntityID() int64 {
	return t.EntityID
}

// SetEntityID will set the NewrelicAlertsViolation EntityID value
func (t *NewrelicAlertsViolation) SetEntityID(v int64) {
	t.EntityID = v
}

// GetEntityName will return the NewrelicAlertsViolation EntityName value
func (t *NewrelicAlertsViolation) GetEntityName() string {
	return t.EntityName
}

// SetEntityName will set the NewrelicAlertsViolation EntityName value
func (t *NewrelicAlertsViolation) SetEntityName(v string) {
	t.EntityName = v
}

// GetPolicyExtID will return the NewrelicAlertsViolation PolicyExtID value
func (t *NewrelicAlertsViolation) GetPolicyExtID() int64 {
	return t.PolicyExtID
}

// SetPolicyExtID will set the NewrelicAlertsViolation PolicyExtID value
func (t *NewrelicAlertsViolation) SetPolicyExtID(v int64) {
	t.PolicyExtID = v
}

// FindNewrelicAlertsViolationsByPolicyExtID will find all NewrelicAlertsViolations by the PolicyExtID value
func FindNewrelicAlertsViolationsByPolicyExtID(ctx context.Context, db *sql.DB, value int64) ([]*NewrelicAlertsViolation, error) {
	q := "SELECT `newrelic_alerts_violation`.`id`,`newrelic_alerts_violation`.`checksum`,`newrelic_alerts_violation`.`customer_id`,`newrelic_alerts_violation`.`ext_id`,`newrelic_alerts_violation`.`duration`,`newrelic_alerts_violation`.`policy_name`,`newrelic_alerts_violation`.`condition_name`,`newrelic_alerts_violation`.`priority`,`newrelic_alerts_violation`.`opened_at`,`newrelic_alerts_violation`.`entity_product`,`newrelic_alerts_violation`.`entity_type`,`newrelic_alerts_violation`.`entity_group_id`,`newrelic_alerts_violation`.`entity_id`,`newrelic_alerts_violation`.`entity_name`,`newrelic_alerts_violation`.`policy_ext_id`,`newrelic_alerts_violation`.`policy_id`,`newrelic_alerts_violation`.`condition_ext_id`,`newrelic_alerts_violation`.`condition_id` FROM `newrelic_alerts_violation` WHERE `policy_ext_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*NewrelicAlertsViolation, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _ExtID sql.NullInt64
		var _Duration sql.NullInt64
		var _PolicyName sql.NullString
		var _ConditionName sql.NullString
		var _Priority sql.NullString
		var _OpenedAt sql.NullInt64
		var _EntityProduct sql.NullString
		var _EntityType sql.NullString
		var _EntityGroupID sql.NullInt64
		var _EntityID sql.NullInt64
		var _EntityName sql.NullString
		var _PolicyExtID sql.NullInt64
		var _PolicyID sql.NullString
		var _ConditionExtID sql.NullInt64
		var _ConditionID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_ExtID,
			&_Duration,
			&_PolicyName,
			&_ConditionName,
			&_Priority,
			&_OpenedAt,
			&_EntityProduct,
			&_EntityType,
			&_EntityGroupID,
			&_EntityID,
			&_EntityName,
			&_PolicyExtID,
			&_PolicyID,
			&_ConditionExtID,
			&_ConditionID,
		)
		if err != nil {
			return nil, err
		}
		t := &NewrelicAlertsViolation{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _ExtID.Valid {
			t.SetExtID(_ExtID.Int64)
		}
		if _Duration.Valid {
			t.SetDuration(_Duration.Int64)
		}
		if _PolicyName.Valid {
			t.SetPolicyName(_PolicyName.String)
		}
		if _ConditionName.Valid {
			t.SetConditionName(_ConditionName.String)
		}
		if _Priority.Valid {
			t.SetPriority(_Priority.String)
		}
		if _OpenedAt.Valid {
			t.SetOpenedAt(_OpenedAt.Int64)
		}
		if _EntityProduct.Valid {
			t.SetEntityProduct(_EntityProduct.String)
		}
		if _EntityType.Valid {
			t.SetEntityType(_EntityType.String)
		}
		if _EntityGroupID.Valid {
			t.SetEntityGroupID(_EntityGroupID.Int64)
		}
		if _EntityID.Valid {
			t.SetEntityID(_EntityID.Int64)
		}
		if _EntityName.Valid {
			t.SetEntityName(_EntityName.String)
		}
		if _PolicyExtID.Valid {
			t.SetPolicyExtID(_PolicyExtID.Int64)
		}
		if _PolicyID.Valid {
			t.SetPolicyID(_PolicyID.String)
		}
		if _ConditionExtID.Valid {
			t.SetConditionExtID(_ConditionExtID.Int64)
		}
		if _ConditionID.Valid {
			t.SetConditionID(_ConditionID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindNewrelicAlertsViolationsByPolicyExtIDTx will find all NewrelicAlertsViolations by the PolicyExtID value using the provided transaction
func FindNewrelicAlertsViolationsByPolicyExtIDTx(ctx context.Context, tx *sql.Tx, value int64) ([]*NewrelicAlertsViolation, error) {
	q := "SELECT `newrelic_alerts_violation`.`id`,`newrelic_alerts_violation`.`checksum`,`newrelic_alerts_violation`.`customer_id`,`newrelic_alerts_violation`.`ext_id`,`newrelic_alerts_violation`.`duration`,`newrelic_alerts_violation`.`policy_name`,`newrelic_alerts_violation`.`condition_name`,`newrelic_alerts_violation`.`priority`,`newrelic_alerts_violation`.`opened_at`,`newrelic_alerts_violation`.`entity_product`,`newrelic_alerts_violation`.`entity_type`,`newrelic_alerts_violation`.`entity_group_id`,`newrelic_alerts_violation`.`entity_id`,`newrelic_alerts_violation`.`entity_name`,`newrelic_alerts_violation`.`policy_ext_id`,`newrelic_alerts_violation`.`policy_id`,`newrelic_alerts_violation`.`condition_ext_id`,`newrelic_alerts_violation`.`condition_id` FROM `newrelic_alerts_violation` WHERE `policy_ext_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*NewrelicAlertsViolation, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _ExtID sql.NullInt64
		var _Duration sql.NullInt64
		var _PolicyName sql.NullString
		var _ConditionName sql.NullString
		var _Priority sql.NullString
		var _OpenedAt sql.NullInt64
		var _EntityProduct sql.NullString
		var _EntityType sql.NullString
		var _EntityGroupID sql.NullInt64
		var _EntityID sql.NullInt64
		var _EntityName sql.NullString
		var _PolicyExtID sql.NullInt64
		var _PolicyID sql.NullString
		var _ConditionExtID sql.NullInt64
		var _ConditionID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_ExtID,
			&_Duration,
			&_PolicyName,
			&_ConditionName,
			&_Priority,
			&_OpenedAt,
			&_EntityProduct,
			&_EntityType,
			&_EntityGroupID,
			&_EntityID,
			&_EntityName,
			&_PolicyExtID,
			&_PolicyID,
			&_ConditionExtID,
			&_ConditionID,
		)
		if err != nil {
			return nil, err
		}
		t := &NewrelicAlertsViolation{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _ExtID.Valid {
			t.SetExtID(_ExtID.Int64)
		}
		if _Duration.Valid {
			t.SetDuration(_Duration.Int64)
		}
		if _PolicyName.Valid {
			t.SetPolicyName(_PolicyName.String)
		}
		if _ConditionName.Valid {
			t.SetConditionName(_ConditionName.String)
		}
		if _Priority.Valid {
			t.SetPriority(_Priority.String)
		}
		if _OpenedAt.Valid {
			t.SetOpenedAt(_OpenedAt.Int64)
		}
		if _EntityProduct.Valid {
			t.SetEntityProduct(_EntityProduct.String)
		}
		if _EntityType.Valid {
			t.SetEntityType(_EntityType.String)
		}
		if _EntityGroupID.Valid {
			t.SetEntityGroupID(_EntityGroupID.Int64)
		}
		if _EntityID.Valid {
			t.SetEntityID(_EntityID.Int64)
		}
		if _EntityName.Valid {
			t.SetEntityName(_EntityName.String)
		}
		if _PolicyExtID.Valid {
			t.SetPolicyExtID(_PolicyExtID.Int64)
		}
		if _PolicyID.Valid {
			t.SetPolicyID(_PolicyID.String)
		}
		if _ConditionExtID.Valid {
			t.SetConditionExtID(_ConditionExtID.Int64)
		}
		if _ConditionID.Valid {
			t.SetConditionID(_ConditionID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetPolicyID will return the NewrelicAlertsViolation PolicyID value
func (t *NewrelicAlertsViolation) GetPolicyID() string {
	return t.PolicyID
}

// SetPolicyID will set the NewrelicAlertsViolation PolicyID value
func (t *NewrelicAlertsViolation) SetPolicyID(v string) {
	t.PolicyID = v
}

// FindNewrelicAlertsViolationsByPolicyID will find all NewrelicAlertsViolations by the PolicyID value
func FindNewrelicAlertsViolationsByPolicyID(ctx context.Context, db *sql.DB, value string) ([]*NewrelicAlertsViolation, error) {
	q := "SELECT `newrelic_alerts_violation`.`id`,`newrelic_alerts_violation`.`checksum`,`newrelic_alerts_violation`.`customer_id`,`newrelic_alerts_violation`.`ext_id`,`newrelic_alerts_violation`.`duration`,`newrelic_alerts_violation`.`policy_name`,`newrelic_alerts_violation`.`condition_name`,`newrelic_alerts_violation`.`priority`,`newrelic_alerts_violation`.`opened_at`,`newrelic_alerts_violation`.`entity_product`,`newrelic_alerts_violation`.`entity_type`,`newrelic_alerts_violation`.`entity_group_id`,`newrelic_alerts_violation`.`entity_id`,`newrelic_alerts_violation`.`entity_name`,`newrelic_alerts_violation`.`policy_ext_id`,`newrelic_alerts_violation`.`policy_id`,`newrelic_alerts_violation`.`condition_ext_id`,`newrelic_alerts_violation`.`condition_id` FROM `newrelic_alerts_violation` WHERE `policy_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*NewrelicAlertsViolation, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _ExtID sql.NullInt64
		var _Duration sql.NullInt64
		var _PolicyName sql.NullString
		var _ConditionName sql.NullString
		var _Priority sql.NullString
		var _OpenedAt sql.NullInt64
		var _EntityProduct sql.NullString
		var _EntityType sql.NullString
		var _EntityGroupID sql.NullInt64
		var _EntityID sql.NullInt64
		var _EntityName sql.NullString
		var _PolicyExtID sql.NullInt64
		var _PolicyID sql.NullString
		var _ConditionExtID sql.NullInt64
		var _ConditionID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_ExtID,
			&_Duration,
			&_PolicyName,
			&_ConditionName,
			&_Priority,
			&_OpenedAt,
			&_EntityProduct,
			&_EntityType,
			&_EntityGroupID,
			&_EntityID,
			&_EntityName,
			&_PolicyExtID,
			&_PolicyID,
			&_ConditionExtID,
			&_ConditionID,
		)
		if err != nil {
			return nil, err
		}
		t := &NewrelicAlertsViolation{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _ExtID.Valid {
			t.SetExtID(_ExtID.Int64)
		}
		if _Duration.Valid {
			t.SetDuration(_Duration.Int64)
		}
		if _PolicyName.Valid {
			t.SetPolicyName(_PolicyName.String)
		}
		if _ConditionName.Valid {
			t.SetConditionName(_ConditionName.String)
		}
		if _Priority.Valid {
			t.SetPriority(_Priority.String)
		}
		if _OpenedAt.Valid {
			t.SetOpenedAt(_OpenedAt.Int64)
		}
		if _EntityProduct.Valid {
			t.SetEntityProduct(_EntityProduct.String)
		}
		if _EntityType.Valid {
			t.SetEntityType(_EntityType.String)
		}
		if _EntityGroupID.Valid {
			t.SetEntityGroupID(_EntityGroupID.Int64)
		}
		if _EntityID.Valid {
			t.SetEntityID(_EntityID.Int64)
		}
		if _EntityName.Valid {
			t.SetEntityName(_EntityName.String)
		}
		if _PolicyExtID.Valid {
			t.SetPolicyExtID(_PolicyExtID.Int64)
		}
		if _PolicyID.Valid {
			t.SetPolicyID(_PolicyID.String)
		}
		if _ConditionExtID.Valid {
			t.SetConditionExtID(_ConditionExtID.Int64)
		}
		if _ConditionID.Valid {
			t.SetConditionID(_ConditionID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindNewrelicAlertsViolationsByPolicyIDTx will find all NewrelicAlertsViolations by the PolicyID value using the provided transaction
func FindNewrelicAlertsViolationsByPolicyIDTx(ctx context.Context, tx *sql.Tx, value string) ([]*NewrelicAlertsViolation, error) {
	q := "SELECT `newrelic_alerts_violation`.`id`,`newrelic_alerts_violation`.`checksum`,`newrelic_alerts_violation`.`customer_id`,`newrelic_alerts_violation`.`ext_id`,`newrelic_alerts_violation`.`duration`,`newrelic_alerts_violation`.`policy_name`,`newrelic_alerts_violation`.`condition_name`,`newrelic_alerts_violation`.`priority`,`newrelic_alerts_violation`.`opened_at`,`newrelic_alerts_violation`.`entity_product`,`newrelic_alerts_violation`.`entity_type`,`newrelic_alerts_violation`.`entity_group_id`,`newrelic_alerts_violation`.`entity_id`,`newrelic_alerts_violation`.`entity_name`,`newrelic_alerts_violation`.`policy_ext_id`,`newrelic_alerts_violation`.`policy_id`,`newrelic_alerts_violation`.`condition_ext_id`,`newrelic_alerts_violation`.`condition_id` FROM `newrelic_alerts_violation` WHERE `policy_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*NewrelicAlertsViolation, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _ExtID sql.NullInt64
		var _Duration sql.NullInt64
		var _PolicyName sql.NullString
		var _ConditionName sql.NullString
		var _Priority sql.NullString
		var _OpenedAt sql.NullInt64
		var _EntityProduct sql.NullString
		var _EntityType sql.NullString
		var _EntityGroupID sql.NullInt64
		var _EntityID sql.NullInt64
		var _EntityName sql.NullString
		var _PolicyExtID sql.NullInt64
		var _PolicyID sql.NullString
		var _ConditionExtID sql.NullInt64
		var _ConditionID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_ExtID,
			&_Duration,
			&_PolicyName,
			&_ConditionName,
			&_Priority,
			&_OpenedAt,
			&_EntityProduct,
			&_EntityType,
			&_EntityGroupID,
			&_EntityID,
			&_EntityName,
			&_PolicyExtID,
			&_PolicyID,
			&_ConditionExtID,
			&_ConditionID,
		)
		if err != nil {
			return nil, err
		}
		t := &NewrelicAlertsViolation{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _ExtID.Valid {
			t.SetExtID(_ExtID.Int64)
		}
		if _Duration.Valid {
			t.SetDuration(_Duration.Int64)
		}
		if _PolicyName.Valid {
			t.SetPolicyName(_PolicyName.String)
		}
		if _ConditionName.Valid {
			t.SetConditionName(_ConditionName.String)
		}
		if _Priority.Valid {
			t.SetPriority(_Priority.String)
		}
		if _OpenedAt.Valid {
			t.SetOpenedAt(_OpenedAt.Int64)
		}
		if _EntityProduct.Valid {
			t.SetEntityProduct(_EntityProduct.String)
		}
		if _EntityType.Valid {
			t.SetEntityType(_EntityType.String)
		}
		if _EntityGroupID.Valid {
			t.SetEntityGroupID(_EntityGroupID.Int64)
		}
		if _EntityID.Valid {
			t.SetEntityID(_EntityID.Int64)
		}
		if _EntityName.Valid {
			t.SetEntityName(_EntityName.String)
		}
		if _PolicyExtID.Valid {
			t.SetPolicyExtID(_PolicyExtID.Int64)
		}
		if _PolicyID.Valid {
			t.SetPolicyID(_PolicyID.String)
		}
		if _ConditionExtID.Valid {
			t.SetConditionExtID(_ConditionExtID.Int64)
		}
		if _ConditionID.Valid {
			t.SetConditionID(_ConditionID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetConditionExtID will return the NewrelicAlertsViolation ConditionExtID value
func (t *NewrelicAlertsViolation) GetConditionExtID() int64 {
	return t.ConditionExtID
}

// SetConditionExtID will set the NewrelicAlertsViolation ConditionExtID value
func (t *NewrelicAlertsViolation) SetConditionExtID(v int64) {
	t.ConditionExtID = v
}

// FindNewrelicAlertsViolationsByConditionExtID will find all NewrelicAlertsViolations by the ConditionExtID value
func FindNewrelicAlertsViolationsByConditionExtID(ctx context.Context, db *sql.DB, value int64) ([]*NewrelicAlertsViolation, error) {
	q := "SELECT `newrelic_alerts_violation`.`id`,`newrelic_alerts_violation`.`checksum`,`newrelic_alerts_violation`.`customer_id`,`newrelic_alerts_violation`.`ext_id`,`newrelic_alerts_violation`.`duration`,`newrelic_alerts_violation`.`policy_name`,`newrelic_alerts_violation`.`condition_name`,`newrelic_alerts_violation`.`priority`,`newrelic_alerts_violation`.`opened_at`,`newrelic_alerts_violation`.`entity_product`,`newrelic_alerts_violation`.`entity_type`,`newrelic_alerts_violation`.`entity_group_id`,`newrelic_alerts_violation`.`entity_id`,`newrelic_alerts_violation`.`entity_name`,`newrelic_alerts_violation`.`policy_ext_id`,`newrelic_alerts_violation`.`policy_id`,`newrelic_alerts_violation`.`condition_ext_id`,`newrelic_alerts_violation`.`condition_id` FROM `newrelic_alerts_violation` WHERE `condition_ext_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*NewrelicAlertsViolation, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _ExtID sql.NullInt64
		var _Duration sql.NullInt64
		var _PolicyName sql.NullString
		var _ConditionName sql.NullString
		var _Priority sql.NullString
		var _OpenedAt sql.NullInt64
		var _EntityProduct sql.NullString
		var _EntityType sql.NullString
		var _EntityGroupID sql.NullInt64
		var _EntityID sql.NullInt64
		var _EntityName sql.NullString
		var _PolicyExtID sql.NullInt64
		var _PolicyID sql.NullString
		var _ConditionExtID sql.NullInt64
		var _ConditionID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_ExtID,
			&_Duration,
			&_PolicyName,
			&_ConditionName,
			&_Priority,
			&_OpenedAt,
			&_EntityProduct,
			&_EntityType,
			&_EntityGroupID,
			&_EntityID,
			&_EntityName,
			&_PolicyExtID,
			&_PolicyID,
			&_ConditionExtID,
			&_ConditionID,
		)
		if err != nil {
			return nil, err
		}
		t := &NewrelicAlertsViolation{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _ExtID.Valid {
			t.SetExtID(_ExtID.Int64)
		}
		if _Duration.Valid {
			t.SetDuration(_Duration.Int64)
		}
		if _PolicyName.Valid {
			t.SetPolicyName(_PolicyName.String)
		}
		if _ConditionName.Valid {
			t.SetConditionName(_ConditionName.String)
		}
		if _Priority.Valid {
			t.SetPriority(_Priority.String)
		}
		if _OpenedAt.Valid {
			t.SetOpenedAt(_OpenedAt.Int64)
		}
		if _EntityProduct.Valid {
			t.SetEntityProduct(_EntityProduct.String)
		}
		if _EntityType.Valid {
			t.SetEntityType(_EntityType.String)
		}
		if _EntityGroupID.Valid {
			t.SetEntityGroupID(_EntityGroupID.Int64)
		}
		if _EntityID.Valid {
			t.SetEntityID(_EntityID.Int64)
		}
		if _EntityName.Valid {
			t.SetEntityName(_EntityName.String)
		}
		if _PolicyExtID.Valid {
			t.SetPolicyExtID(_PolicyExtID.Int64)
		}
		if _PolicyID.Valid {
			t.SetPolicyID(_PolicyID.String)
		}
		if _ConditionExtID.Valid {
			t.SetConditionExtID(_ConditionExtID.Int64)
		}
		if _ConditionID.Valid {
			t.SetConditionID(_ConditionID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindNewrelicAlertsViolationsByConditionExtIDTx will find all NewrelicAlertsViolations by the ConditionExtID value using the provided transaction
func FindNewrelicAlertsViolationsByConditionExtIDTx(ctx context.Context, tx *sql.Tx, value int64) ([]*NewrelicAlertsViolation, error) {
	q := "SELECT `newrelic_alerts_violation`.`id`,`newrelic_alerts_violation`.`checksum`,`newrelic_alerts_violation`.`customer_id`,`newrelic_alerts_violation`.`ext_id`,`newrelic_alerts_violation`.`duration`,`newrelic_alerts_violation`.`policy_name`,`newrelic_alerts_violation`.`condition_name`,`newrelic_alerts_violation`.`priority`,`newrelic_alerts_violation`.`opened_at`,`newrelic_alerts_violation`.`entity_product`,`newrelic_alerts_violation`.`entity_type`,`newrelic_alerts_violation`.`entity_group_id`,`newrelic_alerts_violation`.`entity_id`,`newrelic_alerts_violation`.`entity_name`,`newrelic_alerts_violation`.`policy_ext_id`,`newrelic_alerts_violation`.`policy_id`,`newrelic_alerts_violation`.`condition_ext_id`,`newrelic_alerts_violation`.`condition_id` FROM `newrelic_alerts_violation` WHERE `condition_ext_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*NewrelicAlertsViolation, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _ExtID sql.NullInt64
		var _Duration sql.NullInt64
		var _PolicyName sql.NullString
		var _ConditionName sql.NullString
		var _Priority sql.NullString
		var _OpenedAt sql.NullInt64
		var _EntityProduct sql.NullString
		var _EntityType sql.NullString
		var _EntityGroupID sql.NullInt64
		var _EntityID sql.NullInt64
		var _EntityName sql.NullString
		var _PolicyExtID sql.NullInt64
		var _PolicyID sql.NullString
		var _ConditionExtID sql.NullInt64
		var _ConditionID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_ExtID,
			&_Duration,
			&_PolicyName,
			&_ConditionName,
			&_Priority,
			&_OpenedAt,
			&_EntityProduct,
			&_EntityType,
			&_EntityGroupID,
			&_EntityID,
			&_EntityName,
			&_PolicyExtID,
			&_PolicyID,
			&_ConditionExtID,
			&_ConditionID,
		)
		if err != nil {
			return nil, err
		}
		t := &NewrelicAlertsViolation{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _ExtID.Valid {
			t.SetExtID(_ExtID.Int64)
		}
		if _Duration.Valid {
			t.SetDuration(_Duration.Int64)
		}
		if _PolicyName.Valid {
			t.SetPolicyName(_PolicyName.String)
		}
		if _ConditionName.Valid {
			t.SetConditionName(_ConditionName.String)
		}
		if _Priority.Valid {
			t.SetPriority(_Priority.String)
		}
		if _OpenedAt.Valid {
			t.SetOpenedAt(_OpenedAt.Int64)
		}
		if _EntityProduct.Valid {
			t.SetEntityProduct(_EntityProduct.String)
		}
		if _EntityType.Valid {
			t.SetEntityType(_EntityType.String)
		}
		if _EntityGroupID.Valid {
			t.SetEntityGroupID(_EntityGroupID.Int64)
		}
		if _EntityID.Valid {
			t.SetEntityID(_EntityID.Int64)
		}
		if _EntityName.Valid {
			t.SetEntityName(_EntityName.String)
		}
		if _PolicyExtID.Valid {
			t.SetPolicyExtID(_PolicyExtID.Int64)
		}
		if _PolicyID.Valid {
			t.SetPolicyID(_PolicyID.String)
		}
		if _ConditionExtID.Valid {
			t.SetConditionExtID(_ConditionExtID.Int64)
		}
		if _ConditionID.Valid {
			t.SetConditionID(_ConditionID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetConditionID will return the NewrelicAlertsViolation ConditionID value
func (t *NewrelicAlertsViolation) GetConditionID() string {
	return t.ConditionID
}

// SetConditionID will set the NewrelicAlertsViolation ConditionID value
func (t *NewrelicAlertsViolation) SetConditionID(v string) {
	t.ConditionID = v
}

// FindNewrelicAlertsViolationsByConditionID will find all NewrelicAlertsViolations by the ConditionID value
func FindNewrelicAlertsViolationsByConditionID(ctx context.Context, db *sql.DB, value string) ([]*NewrelicAlertsViolation, error) {
	q := "SELECT `newrelic_alerts_violation`.`id`,`newrelic_alerts_violation`.`checksum`,`newrelic_alerts_violation`.`customer_id`,`newrelic_alerts_violation`.`ext_id`,`newrelic_alerts_violation`.`duration`,`newrelic_alerts_violation`.`policy_name`,`newrelic_alerts_violation`.`condition_name`,`newrelic_alerts_violation`.`priority`,`newrelic_alerts_violation`.`opened_at`,`newrelic_alerts_violation`.`entity_product`,`newrelic_alerts_violation`.`entity_type`,`newrelic_alerts_violation`.`entity_group_id`,`newrelic_alerts_violation`.`entity_id`,`newrelic_alerts_violation`.`entity_name`,`newrelic_alerts_violation`.`policy_ext_id`,`newrelic_alerts_violation`.`policy_id`,`newrelic_alerts_violation`.`condition_ext_id`,`newrelic_alerts_violation`.`condition_id` FROM `newrelic_alerts_violation` WHERE `condition_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*NewrelicAlertsViolation, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _ExtID sql.NullInt64
		var _Duration sql.NullInt64
		var _PolicyName sql.NullString
		var _ConditionName sql.NullString
		var _Priority sql.NullString
		var _OpenedAt sql.NullInt64
		var _EntityProduct sql.NullString
		var _EntityType sql.NullString
		var _EntityGroupID sql.NullInt64
		var _EntityID sql.NullInt64
		var _EntityName sql.NullString
		var _PolicyExtID sql.NullInt64
		var _PolicyID sql.NullString
		var _ConditionExtID sql.NullInt64
		var _ConditionID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_ExtID,
			&_Duration,
			&_PolicyName,
			&_ConditionName,
			&_Priority,
			&_OpenedAt,
			&_EntityProduct,
			&_EntityType,
			&_EntityGroupID,
			&_EntityID,
			&_EntityName,
			&_PolicyExtID,
			&_PolicyID,
			&_ConditionExtID,
			&_ConditionID,
		)
		if err != nil {
			return nil, err
		}
		t := &NewrelicAlertsViolation{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _ExtID.Valid {
			t.SetExtID(_ExtID.Int64)
		}
		if _Duration.Valid {
			t.SetDuration(_Duration.Int64)
		}
		if _PolicyName.Valid {
			t.SetPolicyName(_PolicyName.String)
		}
		if _ConditionName.Valid {
			t.SetConditionName(_ConditionName.String)
		}
		if _Priority.Valid {
			t.SetPriority(_Priority.String)
		}
		if _OpenedAt.Valid {
			t.SetOpenedAt(_OpenedAt.Int64)
		}
		if _EntityProduct.Valid {
			t.SetEntityProduct(_EntityProduct.String)
		}
		if _EntityType.Valid {
			t.SetEntityType(_EntityType.String)
		}
		if _EntityGroupID.Valid {
			t.SetEntityGroupID(_EntityGroupID.Int64)
		}
		if _EntityID.Valid {
			t.SetEntityID(_EntityID.Int64)
		}
		if _EntityName.Valid {
			t.SetEntityName(_EntityName.String)
		}
		if _PolicyExtID.Valid {
			t.SetPolicyExtID(_PolicyExtID.Int64)
		}
		if _PolicyID.Valid {
			t.SetPolicyID(_PolicyID.String)
		}
		if _ConditionExtID.Valid {
			t.SetConditionExtID(_ConditionExtID.Int64)
		}
		if _ConditionID.Valid {
			t.SetConditionID(_ConditionID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindNewrelicAlertsViolationsByConditionIDTx will find all NewrelicAlertsViolations by the ConditionID value using the provided transaction
func FindNewrelicAlertsViolationsByConditionIDTx(ctx context.Context, tx *sql.Tx, value string) ([]*NewrelicAlertsViolation, error) {
	q := "SELECT `newrelic_alerts_violation`.`id`,`newrelic_alerts_violation`.`checksum`,`newrelic_alerts_violation`.`customer_id`,`newrelic_alerts_violation`.`ext_id`,`newrelic_alerts_violation`.`duration`,`newrelic_alerts_violation`.`policy_name`,`newrelic_alerts_violation`.`condition_name`,`newrelic_alerts_violation`.`priority`,`newrelic_alerts_violation`.`opened_at`,`newrelic_alerts_violation`.`entity_product`,`newrelic_alerts_violation`.`entity_type`,`newrelic_alerts_violation`.`entity_group_id`,`newrelic_alerts_violation`.`entity_id`,`newrelic_alerts_violation`.`entity_name`,`newrelic_alerts_violation`.`policy_ext_id`,`newrelic_alerts_violation`.`policy_id`,`newrelic_alerts_violation`.`condition_ext_id`,`newrelic_alerts_violation`.`condition_id` FROM `newrelic_alerts_violation` WHERE `condition_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*NewrelicAlertsViolation, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _ExtID sql.NullInt64
		var _Duration sql.NullInt64
		var _PolicyName sql.NullString
		var _ConditionName sql.NullString
		var _Priority sql.NullString
		var _OpenedAt sql.NullInt64
		var _EntityProduct sql.NullString
		var _EntityType sql.NullString
		var _EntityGroupID sql.NullInt64
		var _EntityID sql.NullInt64
		var _EntityName sql.NullString
		var _PolicyExtID sql.NullInt64
		var _PolicyID sql.NullString
		var _ConditionExtID sql.NullInt64
		var _ConditionID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_ExtID,
			&_Duration,
			&_PolicyName,
			&_ConditionName,
			&_Priority,
			&_OpenedAt,
			&_EntityProduct,
			&_EntityType,
			&_EntityGroupID,
			&_EntityID,
			&_EntityName,
			&_PolicyExtID,
			&_PolicyID,
			&_ConditionExtID,
			&_ConditionID,
		)
		if err != nil {
			return nil, err
		}
		t := &NewrelicAlertsViolation{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _ExtID.Valid {
			t.SetExtID(_ExtID.Int64)
		}
		if _Duration.Valid {
			t.SetDuration(_Duration.Int64)
		}
		if _PolicyName.Valid {
			t.SetPolicyName(_PolicyName.String)
		}
		if _ConditionName.Valid {
			t.SetConditionName(_ConditionName.String)
		}
		if _Priority.Valid {
			t.SetPriority(_Priority.String)
		}
		if _OpenedAt.Valid {
			t.SetOpenedAt(_OpenedAt.Int64)
		}
		if _EntityProduct.Valid {
			t.SetEntityProduct(_EntityProduct.String)
		}
		if _EntityType.Valid {
			t.SetEntityType(_EntityType.String)
		}
		if _EntityGroupID.Valid {
			t.SetEntityGroupID(_EntityGroupID.Int64)
		}
		if _EntityID.Valid {
			t.SetEntityID(_EntityID.Int64)
		}
		if _EntityName.Valid {
			t.SetEntityName(_EntityName.String)
		}
		if _PolicyExtID.Valid {
			t.SetPolicyExtID(_PolicyExtID.Int64)
		}
		if _PolicyID.Valid {
			t.SetPolicyID(_PolicyID.String)
		}
		if _ConditionExtID.Valid {
			t.SetConditionExtID(_ConditionExtID.Int64)
		}
		if _ConditionID.Valid {
			t.SetConditionID(_ConditionID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

func (t *NewrelicAlertsViolation) toTimestamp(value time.Time) *timestamp.Timestamp {
	ts, _ := ptypes.TimestampProto(value)
	return ts
}

// DBCreateNewrelicAlertsViolationTable will create the NewrelicAlertsViolation table
func DBCreateNewrelicAlertsViolationTable(ctx context.Context, db *sql.DB) error {
	q := "CREATE TABLE `newrelic_alerts_violation` (`id`VARCHAR(64) NOT NULL PRIMARY KEY,`checksum`CHAR(64),`customer_id`VARCHAR(64) NOT NULL,`ext_id` BIGINT(20) NOT NULL,`duration`BIGINT(20) NOT NULL,`policy_name`TEXT NOT NULL,`condition_name`TEXT NOT NULL,`priority`TEXT NOT NULL,`opened_at` BIGINT(20) NOT NULL,`entity_product`TEXT NOT NULL,`entity_type`TEXT NOT NULL,`entity_group_id` BIGINT(20) NOT NULL,`entity_id` BIGINT(20) NOT NULL,`entity_name`TEXT NOT NULL,`policy_ext_id` BIGINT(20) NOT NULL,`policy_id` VARCHAR(64) NOT NULL,`condition_ext_id` BIGINT(20) NOT NULL,`condition_id` VARCHAR(64) NOT NULL,INDEX newrelic_alerts_violation_customer_id_index (`customer_id`),INDEX newrelic_alerts_violation_policy_ext_id_index (`policy_ext_id`),INDEX newrelic_alerts_violation_policy_id_index (`policy_id`),INDEX newrelic_alerts_violation_condition_ext_id_index (`condition_ext_id`),INDEX newrelic_alerts_violation_condition_id_index (`condition_id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBCreateNewrelicAlertsViolationTableTx will create the NewrelicAlertsViolation table using the provided transction
func DBCreateNewrelicAlertsViolationTableTx(ctx context.Context, tx *sql.Tx) error {
	q := "CREATE TABLE `newrelic_alerts_violation` (`id`VARCHAR(64) NOT NULL PRIMARY KEY,`checksum`CHAR(64),`customer_id`VARCHAR(64) NOT NULL,`ext_id` BIGINT(20) NOT NULL,`duration`BIGINT(20) NOT NULL,`policy_name`TEXT NOT NULL,`condition_name`TEXT NOT NULL,`priority`TEXT NOT NULL,`opened_at` BIGINT(20) NOT NULL,`entity_product`TEXT NOT NULL,`entity_type`TEXT NOT NULL,`entity_group_id` BIGINT(20) NOT NULL,`entity_id` BIGINT(20) NOT NULL,`entity_name`TEXT NOT NULL,`policy_ext_id` BIGINT(20) NOT NULL,`policy_id` VARCHAR(64) NOT NULL,`condition_ext_id` BIGINT(20) NOT NULL,`condition_id` VARCHAR(64) NOT NULL,INDEX newrelic_alerts_violation_customer_id_index (`customer_id`),INDEX newrelic_alerts_violation_policy_ext_id_index (`policy_ext_id`),INDEX newrelic_alerts_violation_policy_id_index (`policy_id`),INDEX newrelic_alerts_violation_condition_ext_id_index (`condition_ext_id`),INDEX newrelic_alerts_violation_condition_id_index (`condition_id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// DBDropNewrelicAlertsViolationTable will drop the NewrelicAlertsViolation table
func DBDropNewrelicAlertsViolationTable(ctx context.Context, db *sql.DB) error {
	q := "DROP TABLE IF EXISTS `newrelic_alerts_violation`"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBDropNewrelicAlertsViolationTableTx will drop the NewrelicAlertsViolation table using the provided transaction
func DBDropNewrelicAlertsViolationTableTx(ctx context.Context, tx *sql.Tx) error {
	q := "DROP TABLE IF EXISTS `newrelic_alerts_violation`"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// CalculateChecksum will calculate a checksum of the SHA1 of all field values
func (t *NewrelicAlertsViolation) CalculateChecksum() string {
	return orm.HashStrings(
		orm.ToString(t.ID),
		orm.ToString(t.CustomerID),
		orm.ToString(t.ExtID),
		orm.ToString(t.Duration),
		orm.ToString(t.PolicyName),
		orm.ToString(t.ConditionName),
		orm.ToString(t.Priority),
		orm.ToString(t.OpenedAt),
		orm.ToString(t.EntityProduct),
		orm.ToString(t.EntityType),
		orm.ToString(t.EntityGroupID),
		orm.ToString(t.EntityID),
		orm.ToString(t.EntityName),
		orm.ToString(t.PolicyExtID),
		orm.ToString(t.PolicyID),
		orm.ToString(t.ConditionExtID),
		orm.ToString(t.ConditionID),
	)
}

// DBCreate will create a new NewrelicAlertsViolation record in the database
func (t *NewrelicAlertsViolation) DBCreate(ctx context.Context, db *sql.DB) (sql.Result, error) {
	q := "INSERT INTO `newrelic_alerts_violation` (`newrelic_alerts_violation`.`id`,`newrelic_alerts_violation`.`checksum`,`newrelic_alerts_violation`.`customer_id`,`newrelic_alerts_violation`.`ext_id`,`newrelic_alerts_violation`.`duration`,`newrelic_alerts_violation`.`policy_name`,`newrelic_alerts_violation`.`condition_name`,`newrelic_alerts_violation`.`priority`,`newrelic_alerts_violation`.`opened_at`,`newrelic_alerts_violation`.`entity_product`,`newrelic_alerts_violation`.`entity_type`,`newrelic_alerts_violation`.`entity_group_id`,`newrelic_alerts_violation`.`entity_id`,`newrelic_alerts_violation`.`entity_name`,`newrelic_alerts_violation`.`policy_ext_id`,`newrelic_alerts_violation`.`policy_id`,`newrelic_alerts_violation`.`condition_ext_id`,`newrelic_alerts_violation`.`condition_id`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLInt64(t.ExtID),
		orm.ToSQLInt64(t.Duration),
		orm.ToSQLString(t.PolicyName),
		orm.ToSQLString(t.ConditionName),
		orm.ToSQLString(t.Priority),
		orm.ToSQLInt64(t.OpenedAt),
		orm.ToSQLString(t.EntityProduct),
		orm.ToSQLString(t.EntityType),
		orm.ToSQLInt64(t.EntityGroupID),
		orm.ToSQLInt64(t.EntityID),
		orm.ToSQLString(t.EntityName),
		orm.ToSQLInt64(t.PolicyExtID),
		orm.ToSQLString(t.PolicyID),
		orm.ToSQLInt64(t.ConditionExtID),
		orm.ToSQLString(t.ConditionID),
	)
}

// DBCreateTx will create a new NewrelicAlertsViolation record in the database using the provided transaction
func (t *NewrelicAlertsViolation) DBCreateTx(ctx context.Context, tx *sql.Tx) (sql.Result, error) {
	q := "INSERT INTO `newrelic_alerts_violation` (`newrelic_alerts_violation`.`id`,`newrelic_alerts_violation`.`checksum`,`newrelic_alerts_violation`.`customer_id`,`newrelic_alerts_violation`.`ext_id`,`newrelic_alerts_violation`.`duration`,`newrelic_alerts_violation`.`policy_name`,`newrelic_alerts_violation`.`condition_name`,`newrelic_alerts_violation`.`priority`,`newrelic_alerts_violation`.`opened_at`,`newrelic_alerts_violation`.`entity_product`,`newrelic_alerts_violation`.`entity_type`,`newrelic_alerts_violation`.`entity_group_id`,`newrelic_alerts_violation`.`entity_id`,`newrelic_alerts_violation`.`entity_name`,`newrelic_alerts_violation`.`policy_ext_id`,`newrelic_alerts_violation`.`policy_id`,`newrelic_alerts_violation`.`condition_ext_id`,`newrelic_alerts_violation`.`condition_id`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLInt64(t.ExtID),
		orm.ToSQLInt64(t.Duration),
		orm.ToSQLString(t.PolicyName),
		orm.ToSQLString(t.ConditionName),
		orm.ToSQLString(t.Priority),
		orm.ToSQLInt64(t.OpenedAt),
		orm.ToSQLString(t.EntityProduct),
		orm.ToSQLString(t.EntityType),
		orm.ToSQLInt64(t.EntityGroupID),
		orm.ToSQLInt64(t.EntityID),
		orm.ToSQLString(t.EntityName),
		orm.ToSQLInt64(t.PolicyExtID),
		orm.ToSQLString(t.PolicyID),
		orm.ToSQLInt64(t.ConditionExtID),
		orm.ToSQLString(t.ConditionID),
	)
}

// DBCreateIgnoreDuplicate will upsert the NewrelicAlertsViolation record in the database
func (t *NewrelicAlertsViolation) DBCreateIgnoreDuplicate(ctx context.Context, db *sql.DB) (sql.Result, error) {
	q := "INSERT INTO `newrelic_alerts_violation` (`newrelic_alerts_violation`.`id`,`newrelic_alerts_violation`.`checksum`,`newrelic_alerts_violation`.`customer_id`,`newrelic_alerts_violation`.`ext_id`,`newrelic_alerts_violation`.`duration`,`newrelic_alerts_violation`.`policy_name`,`newrelic_alerts_violation`.`condition_name`,`newrelic_alerts_violation`.`priority`,`newrelic_alerts_violation`.`opened_at`,`newrelic_alerts_violation`.`entity_product`,`newrelic_alerts_violation`.`entity_type`,`newrelic_alerts_violation`.`entity_group_id`,`newrelic_alerts_violation`.`entity_id`,`newrelic_alerts_violation`.`entity_name`,`newrelic_alerts_violation`.`policy_ext_id`,`newrelic_alerts_violation`.`policy_id`,`newrelic_alerts_violation`.`condition_ext_id`,`newrelic_alerts_violation`.`condition_id`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLInt64(t.ExtID),
		orm.ToSQLInt64(t.Duration),
		orm.ToSQLString(t.PolicyName),
		orm.ToSQLString(t.ConditionName),
		orm.ToSQLString(t.Priority),
		orm.ToSQLInt64(t.OpenedAt),
		orm.ToSQLString(t.EntityProduct),
		orm.ToSQLString(t.EntityType),
		orm.ToSQLInt64(t.EntityGroupID),
		orm.ToSQLInt64(t.EntityID),
		orm.ToSQLString(t.EntityName),
		orm.ToSQLInt64(t.PolicyExtID),
		orm.ToSQLString(t.PolicyID),
		orm.ToSQLInt64(t.ConditionExtID),
		orm.ToSQLString(t.ConditionID),
	)
}

// DBCreateIgnoreDuplicateTx will upsert the NewrelicAlertsViolation record in the database using the provided transaction
func (t *NewrelicAlertsViolation) DBCreateIgnoreDuplicateTx(ctx context.Context, tx *sql.Tx) (sql.Result, error) {
	q := "INSERT INTO `newrelic_alerts_violation` (`newrelic_alerts_violation`.`id`,`newrelic_alerts_violation`.`checksum`,`newrelic_alerts_violation`.`customer_id`,`newrelic_alerts_violation`.`ext_id`,`newrelic_alerts_violation`.`duration`,`newrelic_alerts_violation`.`policy_name`,`newrelic_alerts_violation`.`condition_name`,`newrelic_alerts_violation`.`priority`,`newrelic_alerts_violation`.`opened_at`,`newrelic_alerts_violation`.`entity_product`,`newrelic_alerts_violation`.`entity_type`,`newrelic_alerts_violation`.`entity_group_id`,`newrelic_alerts_violation`.`entity_id`,`newrelic_alerts_violation`.`entity_name`,`newrelic_alerts_violation`.`policy_ext_id`,`newrelic_alerts_violation`.`policy_id`,`newrelic_alerts_violation`.`condition_ext_id`,`newrelic_alerts_violation`.`condition_id`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLInt64(t.ExtID),
		orm.ToSQLInt64(t.Duration),
		orm.ToSQLString(t.PolicyName),
		orm.ToSQLString(t.ConditionName),
		orm.ToSQLString(t.Priority),
		orm.ToSQLInt64(t.OpenedAt),
		orm.ToSQLString(t.EntityProduct),
		orm.ToSQLString(t.EntityType),
		orm.ToSQLInt64(t.EntityGroupID),
		orm.ToSQLInt64(t.EntityID),
		orm.ToSQLString(t.EntityName),
		orm.ToSQLInt64(t.PolicyExtID),
		orm.ToSQLString(t.PolicyID),
		orm.ToSQLInt64(t.ConditionExtID),
		orm.ToSQLString(t.ConditionID),
	)
}

// DeleteAllNewrelicAlertsViolations deletes all NewrelicAlertsViolation records in the database with optional filters
func DeleteAllNewrelicAlertsViolations(ctx context.Context, db *sql.DB, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(NewrelicAlertsViolationTableName),
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

// DeleteAllNewrelicAlertsViolationsTx deletes all NewrelicAlertsViolation records in the database with optional filters using the provided transaction
func DeleteAllNewrelicAlertsViolationsTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(NewrelicAlertsViolationTableName),
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

// DBDelete will delete this NewrelicAlertsViolation record in the database
func (t *NewrelicAlertsViolation) DBDelete(ctx context.Context, db *sql.DB) (bool, error) {
	q := "DELETE FROM `newrelic_alerts_violation` WHERE `id` = ?"
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

// DBDeleteTx will delete this NewrelicAlertsViolation record in the database using the provided transaction
func (t *NewrelicAlertsViolation) DBDeleteTx(ctx context.Context, tx *sql.Tx) (bool, error) {
	q := "DELETE FROM `newrelic_alerts_violation` WHERE `id` = ?"
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

// DBUpdate will update the NewrelicAlertsViolation record in the database
func (t *NewrelicAlertsViolation) DBUpdate(ctx context.Context, db *sql.DB) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `newrelic_alerts_violation` SET `checksum`=?,`customer_id`=?,`ext_id`=?,`duration`=?,`policy_name`=?,`condition_name`=?,`priority`=?,`opened_at`=?,`entity_product`=?,`entity_type`=?,`entity_group_id`=?,`entity_id`=?,`entity_name`=?,`policy_ext_id`=?,`policy_id`=?,`condition_ext_id`=?,`condition_id`=? WHERE `id`=?"
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLInt64(t.ExtID),
		orm.ToSQLInt64(t.Duration),
		orm.ToSQLString(t.PolicyName),
		orm.ToSQLString(t.ConditionName),
		orm.ToSQLString(t.Priority),
		orm.ToSQLInt64(t.OpenedAt),
		orm.ToSQLString(t.EntityProduct),
		orm.ToSQLString(t.EntityType),
		orm.ToSQLInt64(t.EntityGroupID),
		orm.ToSQLInt64(t.EntityID),
		orm.ToSQLString(t.EntityName),
		orm.ToSQLInt64(t.PolicyExtID),
		orm.ToSQLString(t.PolicyID),
		orm.ToSQLInt64(t.ConditionExtID),
		orm.ToSQLString(t.ConditionID),
		orm.ToSQLString(t.ID),
	)
}

// DBUpdateTx will update the NewrelicAlertsViolation record in the database using the provided transaction
func (t *NewrelicAlertsViolation) DBUpdateTx(ctx context.Context, tx *sql.Tx) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `newrelic_alerts_violation` SET `checksum`=?,`customer_id`=?,`ext_id`=?,`duration`=?,`policy_name`=?,`condition_name`=?,`priority`=?,`opened_at`=?,`entity_product`=?,`entity_type`=?,`entity_group_id`=?,`entity_id`=?,`entity_name`=?,`policy_ext_id`=?,`policy_id`=?,`condition_ext_id`=?,`condition_id`=? WHERE `id`=?"
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLInt64(t.ExtID),
		orm.ToSQLInt64(t.Duration),
		orm.ToSQLString(t.PolicyName),
		orm.ToSQLString(t.ConditionName),
		orm.ToSQLString(t.Priority),
		orm.ToSQLInt64(t.OpenedAt),
		orm.ToSQLString(t.EntityProduct),
		orm.ToSQLString(t.EntityType),
		orm.ToSQLInt64(t.EntityGroupID),
		orm.ToSQLInt64(t.EntityID),
		orm.ToSQLString(t.EntityName),
		orm.ToSQLInt64(t.PolicyExtID),
		orm.ToSQLString(t.PolicyID),
		orm.ToSQLInt64(t.ConditionExtID),
		orm.ToSQLString(t.ConditionID),
		orm.ToSQLString(t.ID),
	)
}

// DBUpsert will upsert the NewrelicAlertsViolation record in the database
func (t *NewrelicAlertsViolation) DBUpsert(ctx context.Context, db *sql.DB, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `newrelic_alerts_violation` (`newrelic_alerts_violation`.`id`,`newrelic_alerts_violation`.`checksum`,`newrelic_alerts_violation`.`customer_id`,`newrelic_alerts_violation`.`ext_id`,`newrelic_alerts_violation`.`duration`,`newrelic_alerts_violation`.`policy_name`,`newrelic_alerts_violation`.`condition_name`,`newrelic_alerts_violation`.`priority`,`newrelic_alerts_violation`.`opened_at`,`newrelic_alerts_violation`.`entity_product`,`newrelic_alerts_violation`.`entity_type`,`newrelic_alerts_violation`.`entity_group_id`,`newrelic_alerts_violation`.`entity_id`,`newrelic_alerts_violation`.`entity_name`,`newrelic_alerts_violation`.`policy_ext_id`,`newrelic_alerts_violation`.`policy_id`,`newrelic_alerts_violation`.`condition_ext_id`,`newrelic_alerts_violation`.`condition_id`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `newrelic_alerts_violation` (`newrelic_alerts_violation`.`id`,`newrelic_alerts_violation`.`checksum`,`newrelic_alerts_violation`.`customer_id`,`newrelic_alerts_violation`.`ext_id`,`newrelic_alerts_violation`.`duration`,`newrelic_alerts_violation`.`policy_name`,`newrelic_alerts_violation`.`condition_name`,`newrelic_alerts_violation`.`priority`,`newrelic_alerts_violation`.`opened_at`,`newrelic_alerts_violation`.`entity_product`,`newrelic_alerts_violation`.`entity_type`,`newrelic_alerts_violation`.`entity_group_id`,`newrelic_alerts_violation`.`entity_id`,`newrelic_alerts_violation`.`entity_name`,`newrelic_alerts_violation`.`policy_ext_id`,`newrelic_alerts_violation`.`policy_id`,`newrelic_alerts_violation`.`condition_ext_id`,`newrelic_alerts_violation`.`condition_id`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`customer_id`=VALUES(`customer_id`),`ext_id`=VALUES(`ext_id`),`duration`=VALUES(`duration`),`policy_name`=VALUES(`policy_name`),`condition_name`=VALUES(`condition_name`),`priority`=VALUES(`priority`),`opened_at`=VALUES(`opened_at`),`entity_product`=VALUES(`entity_product`),`entity_type`=VALUES(`entity_type`),`entity_group_id`=VALUES(`entity_group_id`),`entity_id`=VALUES(`entity_id`),`entity_name`=VALUES(`entity_name`),`policy_ext_id`=VALUES(`policy_ext_id`),`policy_id`=VALUES(`policy_id`),`condition_ext_id`=VALUES(`condition_ext_id`),`condition_id`=VALUES(`condition_id`)"
	}
	r, err := db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLInt64(t.ExtID),
		orm.ToSQLInt64(t.Duration),
		orm.ToSQLString(t.PolicyName),
		orm.ToSQLString(t.ConditionName),
		orm.ToSQLString(t.Priority),
		orm.ToSQLInt64(t.OpenedAt),
		orm.ToSQLString(t.EntityProduct),
		orm.ToSQLString(t.EntityType),
		orm.ToSQLInt64(t.EntityGroupID),
		orm.ToSQLInt64(t.EntityID),
		orm.ToSQLString(t.EntityName),
		orm.ToSQLInt64(t.PolicyExtID),
		orm.ToSQLString(t.PolicyID),
		orm.ToSQLInt64(t.ConditionExtID),
		orm.ToSQLString(t.ConditionID),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBUpsertTx will upsert the NewrelicAlertsViolation record in the database using the provided transaction
func (t *NewrelicAlertsViolation) DBUpsertTx(ctx context.Context, tx *sql.Tx, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `newrelic_alerts_violation` (`newrelic_alerts_violation`.`id`,`newrelic_alerts_violation`.`checksum`,`newrelic_alerts_violation`.`customer_id`,`newrelic_alerts_violation`.`ext_id`,`newrelic_alerts_violation`.`duration`,`newrelic_alerts_violation`.`policy_name`,`newrelic_alerts_violation`.`condition_name`,`newrelic_alerts_violation`.`priority`,`newrelic_alerts_violation`.`opened_at`,`newrelic_alerts_violation`.`entity_product`,`newrelic_alerts_violation`.`entity_type`,`newrelic_alerts_violation`.`entity_group_id`,`newrelic_alerts_violation`.`entity_id`,`newrelic_alerts_violation`.`entity_name`,`newrelic_alerts_violation`.`policy_ext_id`,`newrelic_alerts_violation`.`policy_id`,`newrelic_alerts_violation`.`condition_ext_id`,`newrelic_alerts_violation`.`condition_id`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `newrelic_alerts_violation` (`newrelic_alerts_violation`.`id`,`newrelic_alerts_violation`.`checksum`,`newrelic_alerts_violation`.`customer_id`,`newrelic_alerts_violation`.`ext_id`,`newrelic_alerts_violation`.`duration`,`newrelic_alerts_violation`.`policy_name`,`newrelic_alerts_violation`.`condition_name`,`newrelic_alerts_violation`.`priority`,`newrelic_alerts_violation`.`opened_at`,`newrelic_alerts_violation`.`entity_product`,`newrelic_alerts_violation`.`entity_type`,`newrelic_alerts_violation`.`entity_group_id`,`newrelic_alerts_violation`.`entity_id`,`newrelic_alerts_violation`.`entity_name`,`newrelic_alerts_violation`.`policy_ext_id`,`newrelic_alerts_violation`.`policy_id`,`newrelic_alerts_violation`.`condition_ext_id`,`newrelic_alerts_violation`.`condition_id`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`customer_id`=VALUES(`customer_id`),`ext_id`=VALUES(`ext_id`),`duration`=VALUES(`duration`),`policy_name`=VALUES(`policy_name`),`condition_name`=VALUES(`condition_name`),`priority`=VALUES(`priority`),`opened_at`=VALUES(`opened_at`),`entity_product`=VALUES(`entity_product`),`entity_type`=VALUES(`entity_type`),`entity_group_id`=VALUES(`entity_group_id`),`entity_id`=VALUES(`entity_id`),`entity_name`=VALUES(`entity_name`),`policy_ext_id`=VALUES(`policy_ext_id`),`policy_id`=VALUES(`policy_id`),`condition_ext_id`=VALUES(`condition_ext_id`),`condition_id`=VALUES(`condition_id`)"
	}
	r, err := tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLInt64(t.ExtID),
		orm.ToSQLInt64(t.Duration),
		orm.ToSQLString(t.PolicyName),
		orm.ToSQLString(t.ConditionName),
		orm.ToSQLString(t.Priority),
		orm.ToSQLInt64(t.OpenedAt),
		orm.ToSQLString(t.EntityProduct),
		orm.ToSQLString(t.EntityType),
		orm.ToSQLInt64(t.EntityGroupID),
		orm.ToSQLInt64(t.EntityID),
		orm.ToSQLString(t.EntityName),
		orm.ToSQLInt64(t.PolicyExtID),
		orm.ToSQLString(t.PolicyID),
		orm.ToSQLInt64(t.ConditionExtID),
		orm.ToSQLString(t.ConditionID),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBFindOne will find a NewrelicAlertsViolation record in the database with the primary key
func (t *NewrelicAlertsViolation) DBFindOne(ctx context.Context, db *sql.DB, value string) (bool, error) {
	q := "SELECT `newrelic_alerts_violation`.`id`,`newrelic_alerts_violation`.`checksum`,`newrelic_alerts_violation`.`customer_id`,`newrelic_alerts_violation`.`ext_id`,`newrelic_alerts_violation`.`duration`,`newrelic_alerts_violation`.`policy_name`,`newrelic_alerts_violation`.`condition_name`,`newrelic_alerts_violation`.`priority`,`newrelic_alerts_violation`.`opened_at`,`newrelic_alerts_violation`.`entity_product`,`newrelic_alerts_violation`.`entity_type`,`newrelic_alerts_violation`.`entity_group_id`,`newrelic_alerts_violation`.`entity_id`,`newrelic_alerts_violation`.`entity_name`,`newrelic_alerts_violation`.`policy_ext_id`,`newrelic_alerts_violation`.`policy_id`,`newrelic_alerts_violation`.`condition_ext_id`,`newrelic_alerts_violation`.`condition_id` FROM `newrelic_alerts_violation` WHERE `id` = ? LIMIT 1"
	row := db.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CustomerID sql.NullString
	var _ExtID sql.NullInt64
	var _Duration sql.NullInt64
	var _PolicyName sql.NullString
	var _ConditionName sql.NullString
	var _Priority sql.NullString
	var _OpenedAt sql.NullInt64
	var _EntityProduct sql.NullString
	var _EntityType sql.NullString
	var _EntityGroupID sql.NullInt64
	var _EntityID sql.NullInt64
	var _EntityName sql.NullString
	var _PolicyExtID sql.NullInt64
	var _PolicyID sql.NullString
	var _ConditionExtID sql.NullInt64
	var _ConditionID sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_ExtID,
		&_Duration,
		&_PolicyName,
		&_ConditionName,
		&_Priority,
		&_OpenedAt,
		&_EntityProduct,
		&_EntityType,
		&_EntityGroupID,
		&_EntityID,
		&_EntityName,
		&_PolicyExtID,
		&_PolicyID,
		&_ConditionExtID,
		&_ConditionID,
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
	if _ExtID.Valid {
		t.SetExtID(_ExtID.Int64)
	}
	if _Duration.Valid {
		t.SetDuration(_Duration.Int64)
	}
	if _PolicyName.Valid {
		t.SetPolicyName(_PolicyName.String)
	}
	if _ConditionName.Valid {
		t.SetConditionName(_ConditionName.String)
	}
	if _Priority.Valid {
		t.SetPriority(_Priority.String)
	}
	if _OpenedAt.Valid {
		t.SetOpenedAt(_OpenedAt.Int64)
	}
	if _EntityProduct.Valid {
		t.SetEntityProduct(_EntityProduct.String)
	}
	if _EntityType.Valid {
		t.SetEntityType(_EntityType.String)
	}
	if _EntityGroupID.Valid {
		t.SetEntityGroupID(_EntityGroupID.Int64)
	}
	if _EntityID.Valid {
		t.SetEntityID(_EntityID.Int64)
	}
	if _EntityName.Valid {
		t.SetEntityName(_EntityName.String)
	}
	if _PolicyExtID.Valid {
		t.SetPolicyExtID(_PolicyExtID.Int64)
	}
	if _PolicyID.Valid {
		t.SetPolicyID(_PolicyID.String)
	}
	if _ConditionExtID.Valid {
		t.SetConditionExtID(_ConditionExtID.Int64)
	}
	if _ConditionID.Valid {
		t.SetConditionID(_ConditionID.String)
	}
	return true, nil
}

// DBFindOneTx will find a NewrelicAlertsViolation record in the database with the primary key using the provided transaction
func (t *NewrelicAlertsViolation) DBFindOneTx(ctx context.Context, tx *sql.Tx, value string) (bool, error) {
	q := "SELECT `newrelic_alerts_violation`.`id`,`newrelic_alerts_violation`.`checksum`,`newrelic_alerts_violation`.`customer_id`,`newrelic_alerts_violation`.`ext_id`,`newrelic_alerts_violation`.`duration`,`newrelic_alerts_violation`.`policy_name`,`newrelic_alerts_violation`.`condition_name`,`newrelic_alerts_violation`.`priority`,`newrelic_alerts_violation`.`opened_at`,`newrelic_alerts_violation`.`entity_product`,`newrelic_alerts_violation`.`entity_type`,`newrelic_alerts_violation`.`entity_group_id`,`newrelic_alerts_violation`.`entity_id`,`newrelic_alerts_violation`.`entity_name`,`newrelic_alerts_violation`.`policy_ext_id`,`newrelic_alerts_violation`.`policy_id`,`newrelic_alerts_violation`.`condition_ext_id`,`newrelic_alerts_violation`.`condition_id` FROM `newrelic_alerts_violation` WHERE `id` = ? LIMIT 1"
	row := tx.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CustomerID sql.NullString
	var _ExtID sql.NullInt64
	var _Duration sql.NullInt64
	var _PolicyName sql.NullString
	var _ConditionName sql.NullString
	var _Priority sql.NullString
	var _OpenedAt sql.NullInt64
	var _EntityProduct sql.NullString
	var _EntityType sql.NullString
	var _EntityGroupID sql.NullInt64
	var _EntityID sql.NullInt64
	var _EntityName sql.NullString
	var _PolicyExtID sql.NullInt64
	var _PolicyID sql.NullString
	var _ConditionExtID sql.NullInt64
	var _ConditionID sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_ExtID,
		&_Duration,
		&_PolicyName,
		&_ConditionName,
		&_Priority,
		&_OpenedAt,
		&_EntityProduct,
		&_EntityType,
		&_EntityGroupID,
		&_EntityID,
		&_EntityName,
		&_PolicyExtID,
		&_PolicyID,
		&_ConditionExtID,
		&_ConditionID,
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
	if _ExtID.Valid {
		t.SetExtID(_ExtID.Int64)
	}
	if _Duration.Valid {
		t.SetDuration(_Duration.Int64)
	}
	if _PolicyName.Valid {
		t.SetPolicyName(_PolicyName.String)
	}
	if _ConditionName.Valid {
		t.SetConditionName(_ConditionName.String)
	}
	if _Priority.Valid {
		t.SetPriority(_Priority.String)
	}
	if _OpenedAt.Valid {
		t.SetOpenedAt(_OpenedAt.Int64)
	}
	if _EntityProduct.Valid {
		t.SetEntityProduct(_EntityProduct.String)
	}
	if _EntityType.Valid {
		t.SetEntityType(_EntityType.String)
	}
	if _EntityGroupID.Valid {
		t.SetEntityGroupID(_EntityGroupID.Int64)
	}
	if _EntityID.Valid {
		t.SetEntityID(_EntityID.Int64)
	}
	if _EntityName.Valid {
		t.SetEntityName(_EntityName.String)
	}
	if _PolicyExtID.Valid {
		t.SetPolicyExtID(_PolicyExtID.Int64)
	}
	if _PolicyID.Valid {
		t.SetPolicyID(_PolicyID.String)
	}
	if _ConditionExtID.Valid {
		t.SetConditionExtID(_ConditionExtID.Int64)
	}
	if _ConditionID.Valid {
		t.SetConditionID(_ConditionID.String)
	}
	return true, nil
}

// FindNewrelicAlertsViolations will find a NewrelicAlertsViolation record in the database with the provided parameters
func FindNewrelicAlertsViolations(ctx context.Context, db *sql.DB, _params ...interface{}) ([]*NewrelicAlertsViolation, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("customer_id"),
		orm.Column("ext_id"),
		orm.Column("duration"),
		orm.Column("policy_name"),
		orm.Column("condition_name"),
		orm.Column("priority"),
		orm.Column("opened_at"),
		orm.Column("entity_product"),
		orm.Column("entity_type"),
		orm.Column("entity_group_id"),
		orm.Column("entity_id"),
		orm.Column("entity_name"),
		orm.Column("policy_ext_id"),
		orm.Column("policy_id"),
		orm.Column("condition_ext_id"),
		orm.Column("condition_id"),
		orm.Table(NewrelicAlertsViolationTableName),
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
	results := make([]*NewrelicAlertsViolation, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _ExtID sql.NullInt64
		var _Duration sql.NullInt64
		var _PolicyName sql.NullString
		var _ConditionName sql.NullString
		var _Priority sql.NullString
		var _OpenedAt sql.NullInt64
		var _EntityProduct sql.NullString
		var _EntityType sql.NullString
		var _EntityGroupID sql.NullInt64
		var _EntityID sql.NullInt64
		var _EntityName sql.NullString
		var _PolicyExtID sql.NullInt64
		var _PolicyID sql.NullString
		var _ConditionExtID sql.NullInt64
		var _ConditionID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_ExtID,
			&_Duration,
			&_PolicyName,
			&_ConditionName,
			&_Priority,
			&_OpenedAt,
			&_EntityProduct,
			&_EntityType,
			&_EntityGroupID,
			&_EntityID,
			&_EntityName,
			&_PolicyExtID,
			&_PolicyID,
			&_ConditionExtID,
			&_ConditionID,
		)
		if err != nil {
			return nil, err
		}
		t := &NewrelicAlertsViolation{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _ExtID.Valid {
			t.SetExtID(_ExtID.Int64)
		}
		if _Duration.Valid {
			t.SetDuration(_Duration.Int64)
		}
		if _PolicyName.Valid {
			t.SetPolicyName(_PolicyName.String)
		}
		if _ConditionName.Valid {
			t.SetConditionName(_ConditionName.String)
		}
		if _Priority.Valid {
			t.SetPriority(_Priority.String)
		}
		if _OpenedAt.Valid {
			t.SetOpenedAt(_OpenedAt.Int64)
		}
		if _EntityProduct.Valid {
			t.SetEntityProduct(_EntityProduct.String)
		}
		if _EntityType.Valid {
			t.SetEntityType(_EntityType.String)
		}
		if _EntityGroupID.Valid {
			t.SetEntityGroupID(_EntityGroupID.Int64)
		}
		if _EntityID.Valid {
			t.SetEntityID(_EntityID.Int64)
		}
		if _EntityName.Valid {
			t.SetEntityName(_EntityName.String)
		}
		if _PolicyExtID.Valid {
			t.SetPolicyExtID(_PolicyExtID.Int64)
		}
		if _PolicyID.Valid {
			t.SetPolicyID(_PolicyID.String)
		}
		if _ConditionExtID.Valid {
			t.SetConditionExtID(_ConditionExtID.Int64)
		}
		if _ConditionID.Valid {
			t.SetConditionID(_ConditionID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindNewrelicAlertsViolationsTx will find a NewrelicAlertsViolation record in the database with the provided parameters using the provided transaction
func FindNewrelicAlertsViolationsTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) ([]*NewrelicAlertsViolation, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("customer_id"),
		orm.Column("ext_id"),
		orm.Column("duration"),
		orm.Column("policy_name"),
		orm.Column("condition_name"),
		orm.Column("priority"),
		orm.Column("opened_at"),
		orm.Column("entity_product"),
		orm.Column("entity_type"),
		orm.Column("entity_group_id"),
		orm.Column("entity_id"),
		orm.Column("entity_name"),
		orm.Column("policy_ext_id"),
		orm.Column("policy_id"),
		orm.Column("condition_ext_id"),
		orm.Column("condition_id"),
		orm.Table(NewrelicAlertsViolationTableName),
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
	results := make([]*NewrelicAlertsViolation, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _ExtID sql.NullInt64
		var _Duration sql.NullInt64
		var _PolicyName sql.NullString
		var _ConditionName sql.NullString
		var _Priority sql.NullString
		var _OpenedAt sql.NullInt64
		var _EntityProduct sql.NullString
		var _EntityType sql.NullString
		var _EntityGroupID sql.NullInt64
		var _EntityID sql.NullInt64
		var _EntityName sql.NullString
		var _PolicyExtID sql.NullInt64
		var _PolicyID sql.NullString
		var _ConditionExtID sql.NullInt64
		var _ConditionID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_ExtID,
			&_Duration,
			&_PolicyName,
			&_ConditionName,
			&_Priority,
			&_OpenedAt,
			&_EntityProduct,
			&_EntityType,
			&_EntityGroupID,
			&_EntityID,
			&_EntityName,
			&_PolicyExtID,
			&_PolicyID,
			&_ConditionExtID,
			&_ConditionID,
		)
		if err != nil {
			return nil, err
		}
		t := &NewrelicAlertsViolation{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _ExtID.Valid {
			t.SetExtID(_ExtID.Int64)
		}
		if _Duration.Valid {
			t.SetDuration(_Duration.Int64)
		}
		if _PolicyName.Valid {
			t.SetPolicyName(_PolicyName.String)
		}
		if _ConditionName.Valid {
			t.SetConditionName(_ConditionName.String)
		}
		if _Priority.Valid {
			t.SetPriority(_Priority.String)
		}
		if _OpenedAt.Valid {
			t.SetOpenedAt(_OpenedAt.Int64)
		}
		if _EntityProduct.Valid {
			t.SetEntityProduct(_EntityProduct.String)
		}
		if _EntityType.Valid {
			t.SetEntityType(_EntityType.String)
		}
		if _EntityGroupID.Valid {
			t.SetEntityGroupID(_EntityGroupID.Int64)
		}
		if _EntityID.Valid {
			t.SetEntityID(_EntityID.Int64)
		}
		if _EntityName.Valid {
			t.SetEntityName(_EntityName.String)
		}
		if _PolicyExtID.Valid {
			t.SetPolicyExtID(_PolicyExtID.Int64)
		}
		if _PolicyID.Valid {
			t.SetPolicyID(_PolicyID.String)
		}
		if _ConditionExtID.Valid {
			t.SetConditionExtID(_ConditionExtID.Int64)
		}
		if _ConditionID.Valid {
			t.SetConditionID(_ConditionID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// DBFind will find a NewrelicAlertsViolation record in the database with the provided parameters
func (t *NewrelicAlertsViolation) DBFind(ctx context.Context, db *sql.DB, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("customer_id"),
		orm.Column("ext_id"),
		orm.Column("duration"),
		orm.Column("policy_name"),
		orm.Column("condition_name"),
		orm.Column("priority"),
		orm.Column("opened_at"),
		orm.Column("entity_product"),
		orm.Column("entity_type"),
		orm.Column("entity_group_id"),
		orm.Column("entity_id"),
		orm.Column("entity_name"),
		orm.Column("policy_ext_id"),
		orm.Column("policy_id"),
		orm.Column("condition_ext_id"),
		orm.Column("condition_id"),
		orm.Table(NewrelicAlertsViolationTableName),
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
	var _ExtID sql.NullInt64
	var _Duration sql.NullInt64
	var _PolicyName sql.NullString
	var _ConditionName sql.NullString
	var _Priority sql.NullString
	var _OpenedAt sql.NullInt64
	var _EntityProduct sql.NullString
	var _EntityType sql.NullString
	var _EntityGroupID sql.NullInt64
	var _EntityID sql.NullInt64
	var _EntityName sql.NullString
	var _PolicyExtID sql.NullInt64
	var _PolicyID sql.NullString
	var _ConditionExtID sql.NullInt64
	var _ConditionID sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_ExtID,
		&_Duration,
		&_PolicyName,
		&_ConditionName,
		&_Priority,
		&_OpenedAt,
		&_EntityProduct,
		&_EntityType,
		&_EntityGroupID,
		&_EntityID,
		&_EntityName,
		&_PolicyExtID,
		&_PolicyID,
		&_ConditionExtID,
		&_ConditionID,
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
	if _ExtID.Valid {
		t.SetExtID(_ExtID.Int64)
	}
	if _Duration.Valid {
		t.SetDuration(_Duration.Int64)
	}
	if _PolicyName.Valid {
		t.SetPolicyName(_PolicyName.String)
	}
	if _ConditionName.Valid {
		t.SetConditionName(_ConditionName.String)
	}
	if _Priority.Valid {
		t.SetPriority(_Priority.String)
	}
	if _OpenedAt.Valid {
		t.SetOpenedAt(_OpenedAt.Int64)
	}
	if _EntityProduct.Valid {
		t.SetEntityProduct(_EntityProduct.String)
	}
	if _EntityType.Valid {
		t.SetEntityType(_EntityType.String)
	}
	if _EntityGroupID.Valid {
		t.SetEntityGroupID(_EntityGroupID.Int64)
	}
	if _EntityID.Valid {
		t.SetEntityID(_EntityID.Int64)
	}
	if _EntityName.Valid {
		t.SetEntityName(_EntityName.String)
	}
	if _PolicyExtID.Valid {
		t.SetPolicyExtID(_PolicyExtID.Int64)
	}
	if _PolicyID.Valid {
		t.SetPolicyID(_PolicyID.String)
	}
	if _ConditionExtID.Valid {
		t.SetConditionExtID(_ConditionExtID.Int64)
	}
	if _ConditionID.Valid {
		t.SetConditionID(_ConditionID.String)
	}
	return true, nil
}

// DBFindTx will find a NewrelicAlertsViolation record in the database with the provided parameters using the provided transaction
func (t *NewrelicAlertsViolation) DBFindTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("customer_id"),
		orm.Column("ext_id"),
		orm.Column("duration"),
		orm.Column("policy_name"),
		orm.Column("condition_name"),
		orm.Column("priority"),
		orm.Column("opened_at"),
		orm.Column("entity_product"),
		orm.Column("entity_type"),
		orm.Column("entity_group_id"),
		orm.Column("entity_id"),
		orm.Column("entity_name"),
		orm.Column("policy_ext_id"),
		orm.Column("policy_id"),
		orm.Column("condition_ext_id"),
		orm.Column("condition_id"),
		orm.Table(NewrelicAlertsViolationTableName),
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
	var _ExtID sql.NullInt64
	var _Duration sql.NullInt64
	var _PolicyName sql.NullString
	var _ConditionName sql.NullString
	var _Priority sql.NullString
	var _OpenedAt sql.NullInt64
	var _EntityProduct sql.NullString
	var _EntityType sql.NullString
	var _EntityGroupID sql.NullInt64
	var _EntityID sql.NullInt64
	var _EntityName sql.NullString
	var _PolicyExtID sql.NullInt64
	var _PolicyID sql.NullString
	var _ConditionExtID sql.NullInt64
	var _ConditionID sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_ExtID,
		&_Duration,
		&_PolicyName,
		&_ConditionName,
		&_Priority,
		&_OpenedAt,
		&_EntityProduct,
		&_EntityType,
		&_EntityGroupID,
		&_EntityID,
		&_EntityName,
		&_PolicyExtID,
		&_PolicyID,
		&_ConditionExtID,
		&_ConditionID,
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
	if _ExtID.Valid {
		t.SetExtID(_ExtID.Int64)
	}
	if _Duration.Valid {
		t.SetDuration(_Duration.Int64)
	}
	if _PolicyName.Valid {
		t.SetPolicyName(_PolicyName.String)
	}
	if _ConditionName.Valid {
		t.SetConditionName(_ConditionName.String)
	}
	if _Priority.Valid {
		t.SetPriority(_Priority.String)
	}
	if _OpenedAt.Valid {
		t.SetOpenedAt(_OpenedAt.Int64)
	}
	if _EntityProduct.Valid {
		t.SetEntityProduct(_EntityProduct.String)
	}
	if _EntityType.Valid {
		t.SetEntityType(_EntityType.String)
	}
	if _EntityGroupID.Valid {
		t.SetEntityGroupID(_EntityGroupID.Int64)
	}
	if _EntityID.Valid {
		t.SetEntityID(_EntityID.Int64)
	}
	if _EntityName.Valid {
		t.SetEntityName(_EntityName.String)
	}
	if _PolicyExtID.Valid {
		t.SetPolicyExtID(_PolicyExtID.Int64)
	}
	if _PolicyID.Valid {
		t.SetPolicyID(_PolicyID.String)
	}
	if _ConditionExtID.Valid {
		t.SetConditionExtID(_ConditionExtID.Int64)
	}
	if _ConditionID.Valid {
		t.SetConditionID(_ConditionID.String)
	}
	return true, nil
}

// CountNewrelicAlertsViolations will find the count of NewrelicAlertsViolation records in the database
func CountNewrelicAlertsViolations(ctx context.Context, db *sql.DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(NewrelicAlertsViolationTableName),
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

// CountNewrelicAlertsViolationsTx will find the count of NewrelicAlertsViolation records in the database using the provided transaction
func CountNewrelicAlertsViolationsTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(NewrelicAlertsViolationTableName),
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

// DBCount will find the count of NewrelicAlertsViolation records in the database
func (t *NewrelicAlertsViolation) DBCount(ctx context.Context, db *sql.DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(NewrelicAlertsViolationTableName),
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

// DBCountTx will find the count of NewrelicAlertsViolation records in the database using the provided transaction
func (t *NewrelicAlertsViolation) DBCountTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(NewrelicAlertsViolationTableName),
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

// DBExists will return true if the NewrelicAlertsViolation record exists in the database
func (t *NewrelicAlertsViolation) DBExists(ctx context.Context, db *sql.DB) (bool, error) {
	q := "SELECT `id` FROM `newrelic_alerts_violation` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := db.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// DBExistsTx will return true if the NewrelicAlertsViolation record exists in the database using the provided transaction
func (t *NewrelicAlertsViolation) DBExistsTx(ctx context.Context, tx *sql.Tx) (bool, error) {
	q := "SELECT `id` FROM `newrelic_alerts_violation` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := tx.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// PrimaryKeyColumn returns the column name for the primary key
func (t *NewrelicAlertsViolation) PrimaryKeyColumn() string {
	return NewrelicAlertsViolationColumnID
}

// PrimaryKeyColumnType returns the primary key column Go type as a string
func (t *NewrelicAlertsViolation) PrimaryKeyColumnType() string {
	return "string"
}

// PrimaryKey returns the primary key column value
func (t *NewrelicAlertsViolation) PrimaryKey() interface{} {
	return t.ID
}
