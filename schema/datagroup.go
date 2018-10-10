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

var _ Model = (*DataGroup)(nil)
var _ CSVWriter = (*DataGroup)(nil)
var _ JSONWriter = (*DataGroup)(nil)
var _ Checksum = (*DataGroup)(nil)

// DataGroupTableName is the name of the table in SQL
const DataGroupTableName = "data_group"

var DataGroupColumns = []string{
	"id",
	"checksum",
	"customer_id",
	"name",
	"description",
	"parent_id",
	"issue_project_ids",
	"user_ids",
	"repo_ids",
	"created_at",
	"updated_at",
}

// DataGroup table
type DataGroup struct {
	Checksum        *string `json:"checksum,omitempty"`
	CreatedAt       int64   `json:"created_at"`
	CustomerID      string  `json:"customer_id"`
	Description     *string `json:"description,omitempty"`
	ID              string  `json:"id"`
	IssueProjectIds *string `json:"issue_project_ids,omitempty"`
	Name            string  `json:"name"`
	ParentID        *string `json:"parent_id,omitempty"`
	RepoIds         *string `json:"repo_ids,omitempty"`
	UpdatedAt       *int64  `json:"updated_at,omitempty"`
	UserIds         *string `json:"user_ids,omitempty"`
}

// TableName returns the SQL table name for DataGroup and satifies the Model interface
func (t *DataGroup) TableName() string {
	return DataGroupTableName
}

// ToCSV will serialize the DataGroup instance to a CSV compatible array of strings
func (t *DataGroup) ToCSV() []string {
	return []string{
		t.ID,
		t.CalculateChecksum(),
		t.CustomerID,
		t.Name,
		toCSVString(t.Description),
		toCSVString(t.ParentID),
		toCSVString(t.IssueProjectIds),
		toCSVString(t.UserIds),
		toCSVString(t.RepoIds),
		toCSVString(t.CreatedAt),
		toCSVString(t.UpdatedAt),
	}
}

// WriteCSV will serialize the DataGroup instance to the writer as CSV and satisfies the CSVWriter interface
func (t *DataGroup) WriteCSV(w *csv.Writer) error {
	return w.Write(t.ToCSV())
}

// WriteJSON will serialize the DataGroup instance to the writer as JSON and satisfies the JSONWriter interface
func (t *DataGroup) WriteJSON(w io.Writer, indent ...bool) error {
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

// NewDataGroupReader creates a JSON reader which can read in DataGroup objects serialized as JSON either as an array, single object or json new lines
// and writes each DataGroup to the channel provided
func NewDataGroupReader(r io.Reader, ch chan<- DataGroup) error {
	return orm.Deserialize(r, func(buf json.RawMessage) error {
		dec := json.NewDecoder(bytes.NewBuffer(buf))
		e := DataGroup{}
		if err := dec.Decode(&e); err != nil {
			return err
		}
		ch <- e
		return nil
	})
}

// NewCSVDataGroupReaderDir will read the reader as CSV and emit each record to the channel provided
func NewCSVDataGroupReader(r io.Reader, ch chan<- DataGroup) error {
	cr := csv.NewReader(r)
	for {
		record, err := cr.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		ch <- DataGroup{
			ID:              record[0],
			Checksum:        fromStringPointer(record[1]),
			CustomerID:      record[2],
			Name:            record[3],
			Description:     fromStringPointer(record[4]),
			ParentID:        fromStringPointer(record[5]),
			IssueProjectIds: fromStringPointer(record[6]),
			UserIds:         fromStringPointer(record[7]),
			RepoIds:         fromStringPointer(record[8]),
			CreatedAt:       fromCSVInt64(record[9]),
			UpdatedAt:       fromCSVInt64Pointer(record[10]),
		}
	}
	return nil
}

// NewCSVDataGroupReaderFile will read the file as a CSV and emit each record to the channel provided
func NewCSVDataGroupReaderFile(fp string, ch chan<- DataGroup) error {
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
	return NewCSVDataGroupReader(fc, ch)
}

// NewCSVDataGroupReaderDir will read the data_group.csv.gz file as a CSV and emit each record to the channel provided
func NewCSVDataGroupReaderDir(dir string, ch chan<- DataGroup) error {
	return NewCSVDataGroupReaderFile(filepath.Join(dir, "data_group.csv.gz"), ch)
}

// DataGroupCSVDeduper is a function callback which takes the existing value (a) and the new value (b)
// and the return value should be the one to use (or a new one, if applicable). return nil
// to skip processing of this record
type DataGroupCSVDeduper func(a DataGroup, b DataGroup) *DataGroup

// DataGroupCSVDedupeDisabled is set on whether the CSV writer should de-dupe values by key
var DataGroupCSVDedupeDisabled bool

// NewDataGroupCSVWriterSize creates a batch writer that will write each DataGroup into a CSV file
// this method will automatically de-duplicate entries using the primary key. if the checksum
// for a newer item with the same primary key doesn't match a previously sent item, the newer
// one will be used
func NewDataGroupCSVWriterSize(w io.Writer, size int, dedupers ...DataGroupCSVDeduper) (chan DataGroup, chan bool, error) {
	cw := csv.NewWriter(w)
	ch := make(chan DataGroup, size)
	done := make(chan bool)
	go func() {
		defer func() { done <- true }()
		dodedupe := !DataGroupCSVDedupeDisabled
		var kv map[string]*DataGroup
		var deduper DataGroupCSVDeduper
		if dedupers != nil && len(dedupers) > 0 {
			deduper = dedupers[0]
			dodedupe = true
		}
		if dodedupe {
			kv = make(map[string]*DataGroup)
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

// DataGroupCSVDefaultSize is the default channel buffer size if not provided
var DataGroupCSVDefaultSize = 100

// NewDataGroupCSVWriter creates a batch writer that will write each DataGroup into a CSV file
func NewDataGroupCSVWriter(w io.Writer, dedupers ...DataGroupCSVDeduper) (chan DataGroup, chan bool, error) {
	return NewDataGroupCSVWriterSize(w, DataGroupCSVDefaultSize, dedupers...)
}

// NewDataGroupCSVWriterDir creates a batch writer that will write each DataGroup into a CSV file named data_group.csv.gz in dir
func NewDataGroupCSVWriterDir(dir string, dedupers ...DataGroupCSVDeduper) (chan DataGroup, chan bool, error) {
	return NewDataGroupCSVWriterFile(filepath.Join(dir, "data_group.csv.gz"), dedupers...)
}

// NewDataGroupCSVWriterFile creates a batch writer that will write each DataGroup into a CSV file
func NewDataGroupCSVWriterFile(fn string, dedupers ...DataGroupCSVDeduper) (chan DataGroup, chan bool, error) {
	f, err := os.Create(fn)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening CSV file %s. %v", fn, err)
	}
	var fc io.WriteCloser = f
	if filepath.Ext(fn) == ".gz" {
		w, _ := gzip.NewWriterLevel(f, gzip.BestCompression)
		fc = w
	}
	ch, done, err := NewDataGroupCSVWriter(fc, dedupers...)
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

type DataGroupDBAction func(ctx context.Context, db DB, record DataGroup) error

// NewDataGroupDBWriterSize creates a DB writer that will write each issue into the DB
func NewDataGroupDBWriterSize(ctx context.Context, db DB, errors chan<- error, size int, actions ...DataGroupDBAction) (chan DataGroup, chan bool, error) {
	ch := make(chan DataGroup, size)
	done := make(chan bool)
	var action DataGroupDBAction
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

// NewDataGroupDBWriter creates a DB writer that will write each issue into the DB
func NewDataGroupDBWriter(ctx context.Context, db DB, errors chan<- error, actions ...DataGroupDBAction) (chan DataGroup, chan bool, error) {
	return NewDataGroupDBWriterSize(ctx, db, errors, 100, actions...)
}

// DataGroupColumnID is the ID SQL column name for the DataGroup table
const DataGroupColumnID = "id"

// DataGroupEscapedColumnID is the escaped ID SQL column name for the DataGroup table
const DataGroupEscapedColumnID = "`id`"

// DataGroupColumnChecksum is the Checksum SQL column name for the DataGroup table
const DataGroupColumnChecksum = "checksum"

// DataGroupEscapedColumnChecksum is the escaped Checksum SQL column name for the DataGroup table
const DataGroupEscapedColumnChecksum = "`checksum`"

// DataGroupColumnCustomerID is the CustomerID SQL column name for the DataGroup table
const DataGroupColumnCustomerID = "customer_id"

// DataGroupEscapedColumnCustomerID is the escaped CustomerID SQL column name for the DataGroup table
const DataGroupEscapedColumnCustomerID = "`customer_id`"

// DataGroupColumnName is the Name SQL column name for the DataGroup table
const DataGroupColumnName = "name"

// DataGroupEscapedColumnName is the escaped Name SQL column name for the DataGroup table
const DataGroupEscapedColumnName = "`name`"

// DataGroupColumnDescription is the Description SQL column name for the DataGroup table
const DataGroupColumnDescription = "description"

// DataGroupEscapedColumnDescription is the escaped Description SQL column name for the DataGroup table
const DataGroupEscapedColumnDescription = "`description`"

// DataGroupColumnParentID is the ParentID SQL column name for the DataGroup table
const DataGroupColumnParentID = "parent_id"

// DataGroupEscapedColumnParentID is the escaped ParentID SQL column name for the DataGroup table
const DataGroupEscapedColumnParentID = "`parent_id`"

// DataGroupColumnIssueProjectIds is the IssueProjectIds SQL column name for the DataGroup table
const DataGroupColumnIssueProjectIds = "issue_project_ids"

// DataGroupEscapedColumnIssueProjectIds is the escaped IssueProjectIds SQL column name for the DataGroup table
const DataGroupEscapedColumnIssueProjectIds = "`issue_project_ids`"

// DataGroupColumnUserIds is the UserIds SQL column name for the DataGroup table
const DataGroupColumnUserIds = "user_ids"

// DataGroupEscapedColumnUserIds is the escaped UserIds SQL column name for the DataGroup table
const DataGroupEscapedColumnUserIds = "`user_ids`"

// DataGroupColumnRepoIds is the RepoIds SQL column name for the DataGroup table
const DataGroupColumnRepoIds = "repo_ids"

// DataGroupEscapedColumnRepoIds is the escaped RepoIds SQL column name for the DataGroup table
const DataGroupEscapedColumnRepoIds = "`repo_ids`"

// DataGroupColumnCreatedAt is the CreatedAt SQL column name for the DataGroup table
const DataGroupColumnCreatedAt = "created_at"

// DataGroupEscapedColumnCreatedAt is the escaped CreatedAt SQL column name for the DataGroup table
const DataGroupEscapedColumnCreatedAt = "`created_at`"

// DataGroupColumnUpdatedAt is the UpdatedAt SQL column name for the DataGroup table
const DataGroupColumnUpdatedAt = "updated_at"

// DataGroupEscapedColumnUpdatedAt is the escaped UpdatedAt SQL column name for the DataGroup table
const DataGroupEscapedColumnUpdatedAt = "`updated_at`"

// GetID will return the DataGroup ID value
func (t *DataGroup) GetID() string {
	return t.ID
}

// SetID will set the DataGroup ID value
func (t *DataGroup) SetID(v string) {
	t.ID = v
}

// FindDataGroupByID will find a DataGroup by ID
func FindDataGroupByID(ctx context.Context, db DB, value string) (*DataGroup, error) {
	q := "SELECT * FROM `data_group` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CustomerID sql.NullString
	var _Name sql.NullString
	var _Description sql.NullString
	var _ParentID sql.NullString
	var _IssueProjectIds sql.NullString
	var _UserIds sql.NullString
	var _RepoIds sql.NullString
	var _CreatedAt sql.NullInt64
	var _UpdatedAt sql.NullInt64
	err := db.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_Name,
		&_Description,
		&_ParentID,
		&_IssueProjectIds,
		&_UserIds,
		&_RepoIds,
		&_CreatedAt,
		&_UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &DataGroup{}
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
	if _ParentID.Valid {
		t.SetParentID(_ParentID.String)
	}
	if _IssueProjectIds.Valid {
		t.SetIssueProjectIds(_IssueProjectIds.String)
	}
	if _UserIds.Valid {
		t.SetUserIds(_UserIds.String)
	}
	if _RepoIds.Valid {
		t.SetRepoIds(_RepoIds.String)
	}
	if _CreatedAt.Valid {
		t.SetCreatedAt(_CreatedAt.Int64)
	}
	if _UpdatedAt.Valid {
		t.SetUpdatedAt(_UpdatedAt.Int64)
	}
	return t, nil
}

// FindDataGroupByIDTx will find a DataGroup by ID using the provided transaction
func FindDataGroupByIDTx(ctx context.Context, tx Tx, value string) (*DataGroup, error) {
	q := "SELECT * FROM `data_group` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CustomerID sql.NullString
	var _Name sql.NullString
	var _Description sql.NullString
	var _ParentID sql.NullString
	var _IssueProjectIds sql.NullString
	var _UserIds sql.NullString
	var _RepoIds sql.NullString
	var _CreatedAt sql.NullInt64
	var _UpdatedAt sql.NullInt64
	err := tx.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_Name,
		&_Description,
		&_ParentID,
		&_IssueProjectIds,
		&_UserIds,
		&_RepoIds,
		&_CreatedAt,
		&_UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &DataGroup{}
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
	if _ParentID.Valid {
		t.SetParentID(_ParentID.String)
	}
	if _IssueProjectIds.Valid {
		t.SetIssueProjectIds(_IssueProjectIds.String)
	}
	if _UserIds.Valid {
		t.SetUserIds(_UserIds.String)
	}
	if _RepoIds.Valid {
		t.SetRepoIds(_RepoIds.String)
	}
	if _CreatedAt.Valid {
		t.SetCreatedAt(_CreatedAt.Int64)
	}
	if _UpdatedAt.Valid {
		t.SetUpdatedAt(_UpdatedAt.Int64)
	}
	return t, nil
}

// GetChecksum will return the DataGroup Checksum value
func (t *DataGroup) GetChecksum() string {
	if t.Checksum == nil {
		return ""
	}
	return *t.Checksum
}

// SetChecksum will set the DataGroup Checksum value
func (t *DataGroup) SetChecksum(v string) {
	t.Checksum = &v
}

// GetCustomerID will return the DataGroup CustomerID value
func (t *DataGroup) GetCustomerID() string {
	return t.CustomerID
}

// SetCustomerID will set the DataGroup CustomerID value
func (t *DataGroup) SetCustomerID(v string) {
	t.CustomerID = v
}

// GetName will return the DataGroup Name value
func (t *DataGroup) GetName() string {
	return t.Name
}

// SetName will set the DataGroup Name value
func (t *DataGroup) SetName(v string) {
	t.Name = v
}

// GetDescription will return the DataGroup Description value
func (t *DataGroup) GetDescription() string {
	if t.Description == nil {
		return ""
	}
	return *t.Description
}

// SetDescription will set the DataGroup Description value
func (t *DataGroup) SetDescription(v string) {
	t.Description = &v
}

// GetParentID will return the DataGroup ParentID value
func (t *DataGroup) GetParentID() string {
	if t.ParentID == nil {
		return ""
	}
	return *t.ParentID
}

// SetParentID will set the DataGroup ParentID value
func (t *DataGroup) SetParentID(v string) {
	t.ParentID = &v
}

// GetIssueProjectIds will return the DataGroup IssueProjectIds value
func (t *DataGroup) GetIssueProjectIds() string {
	if t.IssueProjectIds == nil {
		return ""
	}
	return *t.IssueProjectIds
}

// SetIssueProjectIds will set the DataGroup IssueProjectIds value
func (t *DataGroup) SetIssueProjectIds(v string) {
	t.IssueProjectIds = &v
}

// GetUserIds will return the DataGroup UserIds value
func (t *DataGroup) GetUserIds() string {
	if t.UserIds == nil {
		return ""
	}
	return *t.UserIds
}

// SetUserIds will set the DataGroup UserIds value
func (t *DataGroup) SetUserIds(v string) {
	t.UserIds = &v
}

// GetRepoIds will return the DataGroup RepoIds value
func (t *DataGroup) GetRepoIds() string {
	if t.RepoIds == nil {
		return ""
	}
	return *t.RepoIds
}

// SetRepoIds will set the DataGroup RepoIds value
func (t *DataGroup) SetRepoIds(v string) {
	t.RepoIds = &v
}

// GetCreatedAt will return the DataGroup CreatedAt value
func (t *DataGroup) GetCreatedAt() int64 {
	return t.CreatedAt
}

// SetCreatedAt will set the DataGroup CreatedAt value
func (t *DataGroup) SetCreatedAt(v int64) {
	t.CreatedAt = v
}

// GetUpdatedAt will return the DataGroup UpdatedAt value
func (t *DataGroup) GetUpdatedAt() int64 {
	if t.UpdatedAt == nil {
		return int64(0)
	}
	return *t.UpdatedAt
}

// SetUpdatedAt will set the DataGroup UpdatedAt value
func (t *DataGroup) SetUpdatedAt(v int64) {
	t.UpdatedAt = &v
}

func (t *DataGroup) toTimestamp(value time.Time) *timestamp.Timestamp {
	ts, _ := ptypes.TimestampProto(value)
	return ts
}

// DBCreateDataGroupTable will create the DataGroup table
func DBCreateDataGroupTable(ctx context.Context, db DB) error {
	q := "CREATE TABLE `data_group` (`id` VARCHAR(64) NOT NULL PRIMARY KEY,`checksum` CHAR(64),`customer_id` VARCHAR(64) NOT NULL,`name` TEXT NOT NULL,`description` TEXT,`parent_id`VARCHAR(64),`issue_project_ids` JSON,`user_ids` JSON,`repo_ids` JSON,`created_at` BIGINT UNSIGNED NOT NULL,`updated_at` BIGINT UNSIGNED) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBCreateDataGroupTableTx will create the DataGroup table using the provided transction
func DBCreateDataGroupTableTx(ctx context.Context, tx Tx) error {
	q := "CREATE TABLE `data_group` (`id` VARCHAR(64) NOT NULL PRIMARY KEY,`checksum` CHAR(64),`customer_id` VARCHAR(64) NOT NULL,`name` TEXT NOT NULL,`description` TEXT,`parent_id`VARCHAR(64),`issue_project_ids` JSON,`user_ids` JSON,`repo_ids` JSON,`created_at` BIGINT UNSIGNED NOT NULL,`updated_at` BIGINT UNSIGNED) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// DBDropDataGroupTable will drop the DataGroup table
func DBDropDataGroupTable(ctx context.Context, db DB) error {
	q := "DROP TABLE IF EXISTS `data_group`"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBDropDataGroupTableTx will drop the DataGroup table using the provided transaction
func DBDropDataGroupTableTx(ctx context.Context, tx Tx) error {
	q := "DROP TABLE IF EXISTS `data_group`"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// CalculateChecksum will calculate a checksum of the SHA1 of all field values
func (t *DataGroup) CalculateChecksum() string {
	return orm.HashStrings(
		orm.ToString(t.ID),
		orm.ToString(t.CustomerID),
		orm.ToString(t.Name),
		orm.ToString(t.Description),
		orm.ToString(t.ParentID),
		orm.ToString(t.IssueProjectIds),
		orm.ToString(t.UserIds),
		orm.ToString(t.RepoIds),
		orm.ToString(t.CreatedAt),
		orm.ToString(t.UpdatedAt),
	)
}

// DBCreate will create a new DataGroup record in the database
func (t *DataGroup) DBCreate(ctx context.Context, db DB) (sql.Result, error) {
	q := "INSERT INTO `data_group` (`data_group`.`id`,`data_group`.`checksum`,`data_group`.`customer_id`,`data_group`.`name`,`data_group`.`description`,`data_group`.`parent_id`,`data_group`.`issue_project_ids`,`data_group`.`user_ids`,`data_group`.`repo_ids`,`data_group`.`created_at`,`data_group`.`updated_at`) VALUES (?,?,?,?,?,?,?,?,?,?,?)"
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
		orm.ToSQLString(t.ParentID),
		orm.ToSQLString(t.IssueProjectIds),
		orm.ToSQLString(t.UserIds),
		orm.ToSQLString(t.RepoIds),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
	)
}

// DBCreateTx will create a new DataGroup record in the database using the provided transaction
func (t *DataGroup) DBCreateTx(ctx context.Context, tx Tx) (sql.Result, error) {
	q := "INSERT INTO `data_group` (`data_group`.`id`,`data_group`.`checksum`,`data_group`.`customer_id`,`data_group`.`name`,`data_group`.`description`,`data_group`.`parent_id`,`data_group`.`issue_project_ids`,`data_group`.`user_ids`,`data_group`.`repo_ids`,`data_group`.`created_at`,`data_group`.`updated_at`) VALUES (?,?,?,?,?,?,?,?,?,?,?)"
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
		orm.ToSQLString(t.ParentID),
		orm.ToSQLString(t.IssueProjectIds),
		orm.ToSQLString(t.UserIds),
		orm.ToSQLString(t.RepoIds),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
	)
}

// DBCreateIgnoreDuplicate will upsert the DataGroup record in the database
func (t *DataGroup) DBCreateIgnoreDuplicate(ctx context.Context, db DB) (sql.Result, error) {
	q := "INSERT INTO `data_group` (`data_group`.`id`,`data_group`.`checksum`,`data_group`.`customer_id`,`data_group`.`name`,`data_group`.`description`,`data_group`.`parent_id`,`data_group`.`issue_project_ids`,`data_group`.`user_ids`,`data_group`.`repo_ids`,`data_group`.`created_at`,`data_group`.`updated_at`) VALUES (?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
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
		orm.ToSQLString(t.ParentID),
		orm.ToSQLString(t.IssueProjectIds),
		orm.ToSQLString(t.UserIds),
		orm.ToSQLString(t.RepoIds),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
	)
}

// DBCreateIgnoreDuplicateTx will upsert the DataGroup record in the database using the provided transaction
func (t *DataGroup) DBCreateIgnoreDuplicateTx(ctx context.Context, tx Tx) (sql.Result, error) {
	q := "INSERT INTO `data_group` (`data_group`.`id`,`data_group`.`checksum`,`data_group`.`customer_id`,`data_group`.`name`,`data_group`.`description`,`data_group`.`parent_id`,`data_group`.`issue_project_ids`,`data_group`.`user_ids`,`data_group`.`repo_ids`,`data_group`.`created_at`,`data_group`.`updated_at`) VALUES (?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
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
		orm.ToSQLString(t.ParentID),
		orm.ToSQLString(t.IssueProjectIds),
		orm.ToSQLString(t.UserIds),
		orm.ToSQLString(t.RepoIds),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
	)
}

// DeleteAllDataGroups deletes all DataGroup records in the database with optional filters
func DeleteAllDataGroups(ctx context.Context, db DB, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(DataGroupTableName),
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

// DeleteAllDataGroupsTx deletes all DataGroup records in the database with optional filters using the provided transaction
func DeleteAllDataGroupsTx(ctx context.Context, tx Tx, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(DataGroupTableName),
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

// DBDelete will delete this DataGroup record in the database
func (t *DataGroup) DBDelete(ctx context.Context, db DB) (bool, error) {
	q := "DELETE FROM `data_group` WHERE `id` = ?"
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

// DBDeleteTx will delete this DataGroup record in the database using the provided transaction
func (t *DataGroup) DBDeleteTx(ctx context.Context, tx Tx) (bool, error) {
	q := "DELETE FROM `data_group` WHERE `id` = ?"
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

// DBUpdate will update the DataGroup record in the database
func (t *DataGroup) DBUpdate(ctx context.Context, db DB) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `data_group` SET `checksum`=?,`customer_id`=?,`name`=?,`description`=?,`parent_id`=?,`issue_project_ids`=?,`user_ids`=?,`repo_ids`=?,`created_at`=?,`updated_at`=? WHERE `id`=?"
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.Description),
		orm.ToSQLString(t.ParentID),
		orm.ToSQLString(t.IssueProjectIds),
		orm.ToSQLString(t.UserIds),
		orm.ToSQLString(t.RepoIds),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
		orm.ToSQLString(t.ID),
	)
}

// DBUpdateTx will update the DataGroup record in the database using the provided transaction
func (t *DataGroup) DBUpdateTx(ctx context.Context, tx Tx) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `data_group` SET `checksum`=?,`customer_id`=?,`name`=?,`description`=?,`parent_id`=?,`issue_project_ids`=?,`user_ids`=?,`repo_ids`=?,`created_at`=?,`updated_at`=? WHERE `id`=?"
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.Description),
		orm.ToSQLString(t.ParentID),
		orm.ToSQLString(t.IssueProjectIds),
		orm.ToSQLString(t.UserIds),
		orm.ToSQLString(t.RepoIds),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
		orm.ToSQLString(t.ID),
	)
}

// DBUpsert will upsert the DataGroup record in the database
func (t *DataGroup) DBUpsert(ctx context.Context, db DB, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `data_group` (`data_group`.`id`,`data_group`.`checksum`,`data_group`.`customer_id`,`data_group`.`name`,`data_group`.`description`,`data_group`.`parent_id`,`data_group`.`issue_project_ids`,`data_group`.`user_ids`,`data_group`.`repo_ids`,`data_group`.`created_at`,`data_group`.`updated_at`) VALUES (?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `data_group` (`data_group`.`id`,`data_group`.`checksum`,`data_group`.`customer_id`,`data_group`.`name`,`data_group`.`description`,`data_group`.`parent_id`,`data_group`.`issue_project_ids`,`data_group`.`user_ids`,`data_group`.`repo_ids`,`data_group`.`created_at`,`data_group`.`updated_at`) VALUES (?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`customer_id`=VALUES(`customer_id`),`name`=VALUES(`name`),`description`=VALUES(`description`),`parent_id`=VALUES(`parent_id`),`issue_project_ids`=VALUES(`issue_project_ids`),`user_ids`=VALUES(`user_ids`),`repo_ids`=VALUES(`repo_ids`),`created_at`=VALUES(`created_at`),`updated_at`=VALUES(`updated_at`)"
	}
	r, err := db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.Description),
		orm.ToSQLString(t.ParentID),
		orm.ToSQLString(t.IssueProjectIds),
		orm.ToSQLString(t.UserIds),
		orm.ToSQLString(t.RepoIds),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBUpsertTx will upsert the DataGroup record in the database using the provided transaction
func (t *DataGroup) DBUpsertTx(ctx context.Context, tx Tx, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `data_group` (`data_group`.`id`,`data_group`.`checksum`,`data_group`.`customer_id`,`data_group`.`name`,`data_group`.`description`,`data_group`.`parent_id`,`data_group`.`issue_project_ids`,`data_group`.`user_ids`,`data_group`.`repo_ids`,`data_group`.`created_at`,`data_group`.`updated_at`) VALUES (?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `data_group` (`data_group`.`id`,`data_group`.`checksum`,`data_group`.`customer_id`,`data_group`.`name`,`data_group`.`description`,`data_group`.`parent_id`,`data_group`.`issue_project_ids`,`data_group`.`user_ids`,`data_group`.`repo_ids`,`data_group`.`created_at`,`data_group`.`updated_at`) VALUES (?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`customer_id`=VALUES(`customer_id`),`name`=VALUES(`name`),`description`=VALUES(`description`),`parent_id`=VALUES(`parent_id`),`issue_project_ids`=VALUES(`issue_project_ids`),`user_ids`=VALUES(`user_ids`),`repo_ids`=VALUES(`repo_ids`),`created_at`=VALUES(`created_at`),`updated_at`=VALUES(`updated_at`)"
	}
	r, err := tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.Description),
		orm.ToSQLString(t.ParentID),
		orm.ToSQLString(t.IssueProjectIds),
		orm.ToSQLString(t.UserIds),
		orm.ToSQLString(t.RepoIds),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBFindOne will find a DataGroup record in the database with the primary key
func (t *DataGroup) DBFindOne(ctx context.Context, db DB, value string) (bool, error) {
	q := "SELECT * FROM `data_group` WHERE `id` = ? LIMIT 1"
	row := db.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CustomerID sql.NullString
	var _Name sql.NullString
	var _Description sql.NullString
	var _ParentID sql.NullString
	var _IssueProjectIds sql.NullString
	var _UserIds sql.NullString
	var _RepoIds sql.NullString
	var _CreatedAt sql.NullInt64
	var _UpdatedAt sql.NullInt64
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_Name,
		&_Description,
		&_ParentID,
		&_IssueProjectIds,
		&_UserIds,
		&_RepoIds,
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
	if _ParentID.Valid {
		t.SetParentID(_ParentID.String)
	}
	if _IssueProjectIds.Valid {
		t.SetIssueProjectIds(_IssueProjectIds.String)
	}
	if _UserIds.Valid {
		t.SetUserIds(_UserIds.String)
	}
	if _RepoIds.Valid {
		t.SetRepoIds(_RepoIds.String)
	}
	if _CreatedAt.Valid {
		t.SetCreatedAt(_CreatedAt.Int64)
	}
	if _UpdatedAt.Valid {
		t.SetUpdatedAt(_UpdatedAt.Int64)
	}
	return true, nil
}

// DBFindOneTx will find a DataGroup record in the database with the primary key using the provided transaction
func (t *DataGroup) DBFindOneTx(ctx context.Context, tx Tx, value string) (bool, error) {
	q := "SELECT * FROM `data_group` WHERE `id` = ? LIMIT 1"
	row := tx.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CustomerID sql.NullString
	var _Name sql.NullString
	var _Description sql.NullString
	var _ParentID sql.NullString
	var _IssueProjectIds sql.NullString
	var _UserIds sql.NullString
	var _RepoIds sql.NullString
	var _CreatedAt sql.NullInt64
	var _UpdatedAt sql.NullInt64
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_Name,
		&_Description,
		&_ParentID,
		&_IssueProjectIds,
		&_UserIds,
		&_RepoIds,
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
	if _ParentID.Valid {
		t.SetParentID(_ParentID.String)
	}
	if _IssueProjectIds.Valid {
		t.SetIssueProjectIds(_IssueProjectIds.String)
	}
	if _UserIds.Valid {
		t.SetUserIds(_UserIds.String)
	}
	if _RepoIds.Valid {
		t.SetRepoIds(_RepoIds.String)
	}
	if _CreatedAt.Valid {
		t.SetCreatedAt(_CreatedAt.Int64)
	}
	if _UpdatedAt.Valid {
		t.SetUpdatedAt(_UpdatedAt.Int64)
	}
	return true, nil
}

// FindDataGroups will find a DataGroup record in the database with the provided parameters
func FindDataGroups(ctx context.Context, db DB, _params ...interface{}) ([]*DataGroup, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("customer_id"),
		orm.Column("name"),
		orm.Column("description"),
		orm.Column("parent_id"),
		orm.Column("issue_project_ids"),
		orm.Column("user_ids"),
		orm.Column("repo_ids"),
		orm.Column("created_at"),
		orm.Column("updated_at"),
		orm.Table(DataGroupTableName),
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
	results := make([]*DataGroup, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _Name sql.NullString
		var _Description sql.NullString
		var _ParentID sql.NullString
		var _IssueProjectIds sql.NullString
		var _UserIds sql.NullString
		var _RepoIds sql.NullString
		var _CreatedAt sql.NullInt64
		var _UpdatedAt sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_Name,
			&_Description,
			&_ParentID,
			&_IssueProjectIds,
			&_UserIds,
			&_RepoIds,
			&_CreatedAt,
			&_UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		t := &DataGroup{}
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
		if _ParentID.Valid {
			t.SetParentID(_ParentID.String)
		}
		if _IssueProjectIds.Valid {
			t.SetIssueProjectIds(_IssueProjectIds.String)
		}
		if _UserIds.Valid {
			t.SetUserIds(_UserIds.String)
		}
		if _RepoIds.Valid {
			t.SetRepoIds(_RepoIds.String)
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

// FindDataGroupsTx will find a DataGroup record in the database with the provided parameters using the provided transaction
func FindDataGroupsTx(ctx context.Context, tx Tx, _params ...interface{}) ([]*DataGroup, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("customer_id"),
		orm.Column("name"),
		orm.Column("description"),
		orm.Column("parent_id"),
		orm.Column("issue_project_ids"),
		orm.Column("user_ids"),
		orm.Column("repo_ids"),
		orm.Column("created_at"),
		orm.Column("updated_at"),
		orm.Table(DataGroupTableName),
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
	results := make([]*DataGroup, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _Name sql.NullString
		var _Description sql.NullString
		var _ParentID sql.NullString
		var _IssueProjectIds sql.NullString
		var _UserIds sql.NullString
		var _RepoIds sql.NullString
		var _CreatedAt sql.NullInt64
		var _UpdatedAt sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_Name,
			&_Description,
			&_ParentID,
			&_IssueProjectIds,
			&_UserIds,
			&_RepoIds,
			&_CreatedAt,
			&_UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		t := &DataGroup{}
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
		if _ParentID.Valid {
			t.SetParentID(_ParentID.String)
		}
		if _IssueProjectIds.Valid {
			t.SetIssueProjectIds(_IssueProjectIds.String)
		}
		if _UserIds.Valid {
			t.SetUserIds(_UserIds.String)
		}
		if _RepoIds.Valid {
			t.SetRepoIds(_RepoIds.String)
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

// DBFind will find a DataGroup record in the database with the provided parameters
func (t *DataGroup) DBFind(ctx context.Context, db DB, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("customer_id"),
		orm.Column("name"),
		orm.Column("description"),
		orm.Column("parent_id"),
		orm.Column("issue_project_ids"),
		orm.Column("user_ids"),
		orm.Column("repo_ids"),
		orm.Column("created_at"),
		orm.Column("updated_at"),
		orm.Table(DataGroupTableName),
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
	var _ParentID sql.NullString
	var _IssueProjectIds sql.NullString
	var _UserIds sql.NullString
	var _RepoIds sql.NullString
	var _CreatedAt sql.NullInt64
	var _UpdatedAt sql.NullInt64
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_Name,
		&_Description,
		&_ParentID,
		&_IssueProjectIds,
		&_UserIds,
		&_RepoIds,
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
	if _ParentID.Valid {
		t.SetParentID(_ParentID.String)
	}
	if _IssueProjectIds.Valid {
		t.SetIssueProjectIds(_IssueProjectIds.String)
	}
	if _UserIds.Valid {
		t.SetUserIds(_UserIds.String)
	}
	if _RepoIds.Valid {
		t.SetRepoIds(_RepoIds.String)
	}
	if _CreatedAt.Valid {
		t.SetCreatedAt(_CreatedAt.Int64)
	}
	if _UpdatedAt.Valid {
		t.SetUpdatedAt(_UpdatedAt.Int64)
	}
	return true, nil
}

// DBFindTx will find a DataGroup record in the database with the provided parameters using the provided transaction
func (t *DataGroup) DBFindTx(ctx context.Context, tx Tx, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("customer_id"),
		orm.Column("name"),
		orm.Column("description"),
		orm.Column("parent_id"),
		orm.Column("issue_project_ids"),
		orm.Column("user_ids"),
		orm.Column("repo_ids"),
		orm.Column("created_at"),
		orm.Column("updated_at"),
		orm.Table(DataGroupTableName),
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
	var _ParentID sql.NullString
	var _IssueProjectIds sql.NullString
	var _UserIds sql.NullString
	var _RepoIds sql.NullString
	var _CreatedAt sql.NullInt64
	var _UpdatedAt sql.NullInt64
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_Name,
		&_Description,
		&_ParentID,
		&_IssueProjectIds,
		&_UserIds,
		&_RepoIds,
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
	if _ParentID.Valid {
		t.SetParentID(_ParentID.String)
	}
	if _IssueProjectIds.Valid {
		t.SetIssueProjectIds(_IssueProjectIds.String)
	}
	if _UserIds.Valid {
		t.SetUserIds(_UserIds.String)
	}
	if _RepoIds.Valid {
		t.SetRepoIds(_RepoIds.String)
	}
	if _CreatedAt.Valid {
		t.SetCreatedAt(_CreatedAt.Int64)
	}
	if _UpdatedAt.Valid {
		t.SetUpdatedAt(_UpdatedAt.Int64)
	}
	return true, nil
}

// CountDataGroups will find the count of DataGroup records in the database
func CountDataGroups(ctx context.Context, db DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(DataGroupTableName),
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

// CountDataGroupsTx will find the count of DataGroup records in the database using the provided transaction
func CountDataGroupsTx(ctx context.Context, tx Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(DataGroupTableName),
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

// DBCount will find the count of DataGroup records in the database
func (t *DataGroup) DBCount(ctx context.Context, db DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(DataGroupTableName),
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

// DBCountTx will find the count of DataGroup records in the database using the provided transaction
func (t *DataGroup) DBCountTx(ctx context.Context, tx Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(DataGroupTableName),
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

// DBExists will return true if the DataGroup record exists in the database
func (t *DataGroup) DBExists(ctx context.Context, db DB) (bool, error) {
	q := "SELECT * FROM `data_group` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := db.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// DBExistsTx will return true if the DataGroup record exists in the database using the provided transaction
func (t *DataGroup) DBExistsTx(ctx context.Context, tx Tx) (bool, error) {
	q := "SELECT * FROM `data_group` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := tx.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// PrimaryKeyColumn returns the column name for the primary key
func (t *DataGroup) PrimaryKeyColumn() string {
	return DataGroupColumnID
}

// PrimaryKeyColumnType returns the primary key column Go type as a string
func (t *DataGroup) PrimaryKeyColumnType() string {
	return "string"
}

// PrimaryKey returns the primary key column value
func (t *DataGroup) PrimaryKey() interface{} {
	return t.ID
}
