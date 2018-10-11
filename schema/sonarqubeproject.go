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

var _ Model = (*SonarqubeProject)(nil)
var _ CSVWriter = (*SonarqubeProject)(nil)
var _ JSONWriter = (*SonarqubeProject)(nil)
var _ Checksum = (*SonarqubeProject)(nil)

// SonarqubeProjectTableName is the name of the table in SQL
const SonarqubeProjectTableName = "sonarqube_project"

var SonarqubeProjectColumns = []string{
	"id",
	"checksum",
	"customer_id",
	"ext_id",
	"key",
	"name",
}

// SonarqubeProject table
type SonarqubeProject struct {
	Checksum   *string `json:"checksum,omitempty"`
	CustomerID string  `json:"customer_id"`
	ExtID      string  `json:"ext_id"`
	ID         string  `json:"id"`
	Key        string  `json:"key"`
	Name       string  `json:"name"`
}

// TableName returns the SQL table name for SonarqubeProject and satifies the Model interface
func (t *SonarqubeProject) TableName() string {
	return SonarqubeProjectTableName
}

// ToCSV will serialize the SonarqubeProject instance to a CSV compatible array of strings
func (t *SonarqubeProject) ToCSV() []string {
	return []string{
		t.ID,
		t.CalculateChecksum(),
		t.CustomerID,
		t.ExtID,
		t.Key,
		t.Name,
	}
}

// WriteCSV will serialize the SonarqubeProject instance to the writer as CSV and satisfies the CSVWriter interface
func (t *SonarqubeProject) WriteCSV(w *csv.Writer) error {
	return w.Write(t.ToCSV())
}

// WriteJSON will serialize the SonarqubeProject instance to the writer as JSON and satisfies the JSONWriter interface
func (t *SonarqubeProject) WriteJSON(w io.Writer, indent ...bool) error {
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

// NewSonarqubeProjectReader creates a JSON reader which can read in SonarqubeProject objects serialized as JSON either as an array, single object or json new lines
// and writes each SonarqubeProject to the channel provided
func NewSonarqubeProjectReader(r io.Reader, ch chan<- SonarqubeProject) error {
	return orm.Deserialize(r, func(buf json.RawMessage) error {
		dec := json.NewDecoder(bytes.NewBuffer(buf))
		e := SonarqubeProject{}
		if err := dec.Decode(&e); err != nil {
			return err
		}
		ch <- e
		return nil
	})
}

// NewCSVSonarqubeProjectReaderDir will read the reader as CSV and emit each record to the channel provided
func NewCSVSonarqubeProjectReader(r io.Reader, ch chan<- SonarqubeProject) error {
	cr := csv.NewReader(r)
	for {
		record, err := cr.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		ch <- SonarqubeProject{
			ID:         record[0],
			Checksum:   fromStringPointer(record[1]),
			CustomerID: record[2],
			ExtID:      record[3],
			Key:        record[4],
			Name:       record[5],
		}
	}
	return nil
}

// NewCSVSonarqubeProjectReaderFile will read the file as a CSV and emit each record to the channel provided
func NewCSVSonarqubeProjectReaderFile(fp string, ch chan<- SonarqubeProject) error {
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
	return NewCSVSonarqubeProjectReader(fc, ch)
}

// NewCSVSonarqubeProjectReaderDir will read the sonarqube_project.csv.gz file as a CSV and emit each record to the channel provided
func NewCSVSonarqubeProjectReaderDir(dir string, ch chan<- SonarqubeProject) error {
	return NewCSVSonarqubeProjectReaderFile(filepath.Join(dir, "sonarqube_project.csv.gz"), ch)
}

// SonarqubeProjectCSVDeduper is a function callback which takes the existing value (a) and the new value (b)
// and the return value should be the one to use (or a new one, if applicable). return nil
// to skip processing of this record
type SonarqubeProjectCSVDeduper func(a SonarqubeProject, b SonarqubeProject) *SonarqubeProject

// SonarqubeProjectCSVDedupeDisabled is set on whether the CSV writer should de-dupe values by key
var SonarqubeProjectCSVDedupeDisabled bool

// NewSonarqubeProjectCSVWriterSize creates a batch writer that will write each SonarqubeProject into a CSV file
// this method will automatically de-duplicate entries using the primary key. if the checksum
// for a newer item with the same primary key doesn't match a previously sent item, the newer
// one will be used
func NewSonarqubeProjectCSVWriterSize(w io.Writer, size int, dedupers ...SonarqubeProjectCSVDeduper) (chan SonarqubeProject, chan bool, error) {
	cw := csv.NewWriter(w)
	ch := make(chan SonarqubeProject, size)
	done := make(chan bool)
	go func() {
		defer func() { done <- true }()
		dodedupe := !SonarqubeProjectCSVDedupeDisabled
		var kv map[string]*SonarqubeProject
		var deduper SonarqubeProjectCSVDeduper
		if dedupers != nil && len(dedupers) > 0 {
			deduper = dedupers[0]
			dodedupe = true
		}
		if dodedupe {
			kv = make(map[string]*SonarqubeProject)
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

// SonarqubeProjectCSVDefaultSize is the default channel buffer size if not provided
var SonarqubeProjectCSVDefaultSize = 100

// NewSonarqubeProjectCSVWriter creates a batch writer that will write each SonarqubeProject into a CSV file
func NewSonarqubeProjectCSVWriter(w io.Writer, dedupers ...SonarqubeProjectCSVDeduper) (chan SonarqubeProject, chan bool, error) {
	return NewSonarqubeProjectCSVWriterSize(w, SonarqubeProjectCSVDefaultSize, dedupers...)
}

// NewSonarqubeProjectCSVWriterDir creates a batch writer that will write each SonarqubeProject into a CSV file named sonarqube_project.csv.gz in dir
func NewSonarqubeProjectCSVWriterDir(dir string, dedupers ...SonarqubeProjectCSVDeduper) (chan SonarqubeProject, chan bool, error) {
	return NewSonarqubeProjectCSVWriterFile(filepath.Join(dir, "sonarqube_project.csv.gz"), dedupers...)
}

// NewSonarqubeProjectCSVWriterFile creates a batch writer that will write each SonarqubeProject into a CSV file
func NewSonarqubeProjectCSVWriterFile(fn string, dedupers ...SonarqubeProjectCSVDeduper) (chan SonarqubeProject, chan bool, error) {
	f, err := os.Create(fn)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening CSV file %s. %v", fn, err)
	}
	var fc io.WriteCloser = f
	if filepath.Ext(fn) == ".gz" {
		w, _ := gzip.NewWriterLevel(f, gzip.BestCompression)
		fc = w
	}
	ch, done, err := NewSonarqubeProjectCSVWriter(fc, dedupers...)
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

type SonarqubeProjectDBAction func(ctx context.Context, db DB, record SonarqubeProject) error

// NewSonarqubeProjectDBWriterSize creates a DB writer that will write each issue into the DB
func NewSonarqubeProjectDBWriterSize(ctx context.Context, db DB, errors chan<- error, size int, actions ...SonarqubeProjectDBAction) (chan SonarqubeProject, chan bool, error) {
	ch := make(chan SonarqubeProject, size)
	done := make(chan bool)
	var action SonarqubeProjectDBAction
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

// NewSonarqubeProjectDBWriter creates a DB writer that will write each issue into the DB
func NewSonarqubeProjectDBWriter(ctx context.Context, db DB, errors chan<- error, actions ...SonarqubeProjectDBAction) (chan SonarqubeProject, chan bool, error) {
	return NewSonarqubeProjectDBWriterSize(ctx, db, errors, 100, actions...)
}

// SonarqubeProjectColumnID is the ID SQL column name for the SonarqubeProject table
const SonarqubeProjectColumnID = "id"

// SonarqubeProjectEscapedColumnID is the escaped ID SQL column name for the SonarqubeProject table
const SonarqubeProjectEscapedColumnID = "`id`"

// SonarqubeProjectColumnChecksum is the Checksum SQL column name for the SonarqubeProject table
const SonarqubeProjectColumnChecksum = "checksum"

// SonarqubeProjectEscapedColumnChecksum is the escaped Checksum SQL column name for the SonarqubeProject table
const SonarqubeProjectEscapedColumnChecksum = "`checksum`"

// SonarqubeProjectColumnCustomerID is the CustomerID SQL column name for the SonarqubeProject table
const SonarqubeProjectColumnCustomerID = "customer_id"

// SonarqubeProjectEscapedColumnCustomerID is the escaped CustomerID SQL column name for the SonarqubeProject table
const SonarqubeProjectEscapedColumnCustomerID = "`customer_id`"

// SonarqubeProjectColumnExtID is the ExtID SQL column name for the SonarqubeProject table
const SonarqubeProjectColumnExtID = "ext_id"

// SonarqubeProjectEscapedColumnExtID is the escaped ExtID SQL column name for the SonarqubeProject table
const SonarqubeProjectEscapedColumnExtID = "`ext_id`"

// SonarqubeProjectColumnKey is the Key SQL column name for the SonarqubeProject table
const SonarqubeProjectColumnKey = "key"

// SonarqubeProjectEscapedColumnKey is the escaped Key SQL column name for the SonarqubeProject table
const SonarqubeProjectEscapedColumnKey = "`key`"

// SonarqubeProjectColumnName is the Name SQL column name for the SonarqubeProject table
const SonarqubeProjectColumnName = "name"

// SonarqubeProjectEscapedColumnName is the escaped Name SQL column name for the SonarqubeProject table
const SonarqubeProjectEscapedColumnName = "`name`"

// GetID will return the SonarqubeProject ID value
func (t *SonarqubeProject) GetID() string {
	return t.ID
}

// SetID will set the SonarqubeProject ID value
func (t *SonarqubeProject) SetID(v string) {
	t.ID = v
}

// FindSonarqubeProjectByID will find a SonarqubeProject by ID
func FindSonarqubeProjectByID(ctx context.Context, db DB, value string) (*SonarqubeProject, error) {
	q := "SELECT `sonarqube_project`.`id`,`sonarqube_project`.`checksum`,`sonarqube_project`.`customer_id`,`sonarqube_project`.`ext_id`,`sonarqube_project`.`key`,`sonarqube_project`.`name` FROM `sonarqube_project` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CustomerID sql.NullString
	var _ExtID sql.NullString
	var _Key sql.NullString
	var _Name sql.NullString
	err := db.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_ExtID,
		&_Key,
		&_Name,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &SonarqubeProject{}
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
		t.SetExtID(_ExtID.String)
	}
	if _Key.Valid {
		t.SetKey(_Key.String)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	return t, nil
}

// FindSonarqubeProjectByIDTx will find a SonarqubeProject by ID using the provided transaction
func FindSonarqubeProjectByIDTx(ctx context.Context, tx Tx, value string) (*SonarqubeProject, error) {
	q := "SELECT `sonarqube_project`.`id`,`sonarqube_project`.`checksum`,`sonarqube_project`.`customer_id`,`sonarqube_project`.`ext_id`,`sonarqube_project`.`key`,`sonarqube_project`.`name` FROM `sonarqube_project` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CustomerID sql.NullString
	var _ExtID sql.NullString
	var _Key sql.NullString
	var _Name sql.NullString
	err := tx.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_ExtID,
		&_Key,
		&_Name,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &SonarqubeProject{}
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
		t.SetExtID(_ExtID.String)
	}
	if _Key.Valid {
		t.SetKey(_Key.String)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	return t, nil
}

// GetChecksum will return the SonarqubeProject Checksum value
func (t *SonarqubeProject) GetChecksum() string {
	if t.Checksum == nil {
		return ""
	}
	return *t.Checksum
}

// SetChecksum will set the SonarqubeProject Checksum value
func (t *SonarqubeProject) SetChecksum(v string) {
	t.Checksum = &v
}

// GetCustomerID will return the SonarqubeProject CustomerID value
func (t *SonarqubeProject) GetCustomerID() string {
	return t.CustomerID
}

// SetCustomerID will set the SonarqubeProject CustomerID value
func (t *SonarqubeProject) SetCustomerID(v string) {
	t.CustomerID = v
}

// FindSonarqubeProjectsByCustomerID will find all SonarqubeProjects by the CustomerID value
func FindSonarqubeProjectsByCustomerID(ctx context.Context, db DB, value string) ([]*SonarqubeProject, error) {
	q := "SELECT `sonarqube_project`.`id`,`sonarqube_project`.`checksum`,`sonarqube_project`.`customer_id`,`sonarqube_project`.`ext_id`,`sonarqube_project`.`key`,`sonarqube_project`.`name` FROM `sonarqube_project` WHERE `customer_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*SonarqubeProject, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _ExtID sql.NullString
		var _Key sql.NullString
		var _Name sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_ExtID,
			&_Key,
			&_Name,
		)
		if err != nil {
			return nil, err
		}
		t := &SonarqubeProject{}
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
			t.SetExtID(_ExtID.String)
		}
		if _Key.Valid {
			t.SetKey(_Key.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindSonarqubeProjectsByCustomerIDTx will find all SonarqubeProjects by the CustomerID value using the provided transaction
func FindSonarqubeProjectsByCustomerIDTx(ctx context.Context, tx Tx, value string) ([]*SonarqubeProject, error) {
	q := "SELECT `sonarqube_project`.`id`,`sonarqube_project`.`checksum`,`sonarqube_project`.`customer_id`,`sonarqube_project`.`ext_id`,`sonarqube_project`.`key`,`sonarqube_project`.`name` FROM `sonarqube_project` WHERE `customer_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*SonarqubeProject, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _ExtID sql.NullString
		var _Key sql.NullString
		var _Name sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_ExtID,
			&_Key,
			&_Name,
		)
		if err != nil {
			return nil, err
		}
		t := &SonarqubeProject{}
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
			t.SetExtID(_ExtID.String)
		}
		if _Key.Valid {
			t.SetKey(_Key.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetExtID will return the SonarqubeProject ExtID value
func (t *SonarqubeProject) GetExtID() string {
	return t.ExtID
}

// SetExtID will set the SonarqubeProject ExtID value
func (t *SonarqubeProject) SetExtID(v string) {
	t.ExtID = v
}

// GetKey will return the SonarqubeProject Key value
func (t *SonarqubeProject) GetKey() string {
	return t.Key
}

// SetKey will set the SonarqubeProject Key value
func (t *SonarqubeProject) SetKey(v string) {
	t.Key = v
}

// GetName will return the SonarqubeProject Name value
func (t *SonarqubeProject) GetName() string {
	return t.Name
}

// SetName will set the SonarqubeProject Name value
func (t *SonarqubeProject) SetName(v string) {
	t.Name = v
}

func (t *SonarqubeProject) toTimestamp(value time.Time) *timestamp.Timestamp {
	ts, _ := ptypes.TimestampProto(value)
	return ts
}

// DBCreateSonarqubeProjectTable will create the SonarqubeProject table
func DBCreateSonarqubeProjectTable(ctx context.Context, db DB) error {
	q := "CREATE TABLE `sonarqube_project` (`id` VARCHAR(64) NOT NULL PRIMARY KEY,`checksum` CHAR(64),`customer_id` VARCHAR(64) NOT NULL,`ext_id` TEXT NOT NULL,`key` TEXT NOT NULL,`name`TEXT NOT NULL,INDEX sonarqube_project_customer_id_index (`customer_id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBCreateSonarqubeProjectTableTx will create the SonarqubeProject table using the provided transction
func DBCreateSonarqubeProjectTableTx(ctx context.Context, tx Tx) error {
	q := "CREATE TABLE `sonarqube_project` (`id` VARCHAR(64) NOT NULL PRIMARY KEY,`checksum` CHAR(64),`customer_id` VARCHAR(64) NOT NULL,`ext_id` TEXT NOT NULL,`key` TEXT NOT NULL,`name`TEXT NOT NULL,INDEX sonarqube_project_customer_id_index (`customer_id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// DBDropSonarqubeProjectTable will drop the SonarqubeProject table
func DBDropSonarqubeProjectTable(ctx context.Context, db DB) error {
	q := "DROP TABLE IF EXISTS `sonarqube_project`"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBDropSonarqubeProjectTableTx will drop the SonarqubeProject table using the provided transaction
func DBDropSonarqubeProjectTableTx(ctx context.Context, tx Tx) error {
	q := "DROP TABLE IF EXISTS `sonarqube_project`"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// CalculateChecksum will calculate a checksum of the SHA1 of all field values
func (t *SonarqubeProject) CalculateChecksum() string {
	return orm.HashStrings(
		orm.ToString(t.ID),
		orm.ToString(t.CustomerID),
		orm.ToString(t.ExtID),
		orm.ToString(t.Key),
		orm.ToString(t.Name),
	)
}

// DBCreate will create a new SonarqubeProject record in the database
func (t *SonarqubeProject) DBCreate(ctx context.Context, db DB) (sql.Result, error) {
	q := "INSERT INTO `sonarqube_project` (`sonarqube_project`.`id`,`sonarqube_project`.`checksum`,`sonarqube_project`.`customer_id`,`sonarqube_project`.`ext_id`,`sonarqube_project`.`key`,`sonarqube_project`.`name`) VALUES (?,?,?,?,?,?)"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.ExtID),
		orm.ToSQLString(t.Key),
		orm.ToSQLString(t.Name),
	)
}

// DBCreateTx will create a new SonarqubeProject record in the database using the provided transaction
func (t *SonarqubeProject) DBCreateTx(ctx context.Context, tx Tx) (sql.Result, error) {
	q := "INSERT INTO `sonarqube_project` (`sonarqube_project`.`id`,`sonarqube_project`.`checksum`,`sonarqube_project`.`customer_id`,`sonarqube_project`.`ext_id`,`sonarqube_project`.`key`,`sonarqube_project`.`name`) VALUES (?,?,?,?,?,?)"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.ExtID),
		orm.ToSQLString(t.Key),
		orm.ToSQLString(t.Name),
	)
}

// DBCreateIgnoreDuplicate will upsert the SonarqubeProject record in the database
func (t *SonarqubeProject) DBCreateIgnoreDuplicate(ctx context.Context, db DB) (sql.Result, error) {
	q := "INSERT INTO `sonarqube_project` (`sonarqube_project`.`id`,`sonarqube_project`.`checksum`,`sonarqube_project`.`customer_id`,`sonarqube_project`.`ext_id`,`sonarqube_project`.`key`,`sonarqube_project`.`name`) VALUES (?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.ExtID),
		orm.ToSQLString(t.Key),
		orm.ToSQLString(t.Name),
	)
}

// DBCreateIgnoreDuplicateTx will upsert the SonarqubeProject record in the database using the provided transaction
func (t *SonarqubeProject) DBCreateIgnoreDuplicateTx(ctx context.Context, tx Tx) (sql.Result, error) {
	q := "INSERT INTO `sonarqube_project` (`sonarqube_project`.`id`,`sonarqube_project`.`checksum`,`sonarqube_project`.`customer_id`,`sonarqube_project`.`ext_id`,`sonarqube_project`.`key`,`sonarqube_project`.`name`) VALUES (?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.ExtID),
		orm.ToSQLString(t.Key),
		orm.ToSQLString(t.Name),
	)
}

// DeleteAllSonarqubeProjects deletes all SonarqubeProject records in the database with optional filters
func DeleteAllSonarqubeProjects(ctx context.Context, db DB, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(SonarqubeProjectTableName),
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

// DeleteAllSonarqubeProjectsTx deletes all SonarqubeProject records in the database with optional filters using the provided transaction
func DeleteAllSonarqubeProjectsTx(ctx context.Context, tx Tx, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(SonarqubeProjectTableName),
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

// DBDelete will delete this SonarqubeProject record in the database
func (t *SonarqubeProject) DBDelete(ctx context.Context, db DB) (bool, error) {
	q := "DELETE FROM `sonarqube_project` WHERE `id` = ?"
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

// DBDeleteTx will delete this SonarqubeProject record in the database using the provided transaction
func (t *SonarqubeProject) DBDeleteTx(ctx context.Context, tx Tx) (bool, error) {
	q := "DELETE FROM `sonarqube_project` WHERE `id` = ?"
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

// DBUpdate will update the SonarqubeProject record in the database
func (t *SonarqubeProject) DBUpdate(ctx context.Context, db DB) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `sonarqube_project` SET `checksum`=?,`customer_id`=?,`ext_id`=?,`key`=?,`name`=? WHERE `id`=?"
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.ExtID),
		orm.ToSQLString(t.Key),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.ID),
	)
}

// DBUpdateTx will update the SonarqubeProject record in the database using the provided transaction
func (t *SonarqubeProject) DBUpdateTx(ctx context.Context, tx Tx) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `sonarqube_project` SET `checksum`=?,`customer_id`=?,`ext_id`=?,`key`=?,`name`=? WHERE `id`=?"
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.ExtID),
		orm.ToSQLString(t.Key),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.ID),
	)
}

// DBUpsert will upsert the SonarqubeProject record in the database
func (t *SonarqubeProject) DBUpsert(ctx context.Context, db DB, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `sonarqube_project` (`sonarqube_project`.`id`,`sonarqube_project`.`checksum`,`sonarqube_project`.`customer_id`,`sonarqube_project`.`ext_id`,`sonarqube_project`.`key`,`sonarqube_project`.`name`) VALUES (?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `sonarqube_project` (`sonarqube_project`.`id`,`sonarqube_project`.`checksum`,`sonarqube_project`.`customer_id`,`sonarqube_project`.`ext_id`,`sonarqube_project`.`key`,`sonarqube_project`.`name`) VALUES (?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`customer_id`=VALUES(`customer_id`),`ext_id`=VALUES(`ext_id`),`key`=VALUES(`key`),`name`=VALUES(`name`)"
	}
	r, err := db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.ExtID),
		orm.ToSQLString(t.Key),
		orm.ToSQLString(t.Name),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBUpsertTx will upsert the SonarqubeProject record in the database using the provided transaction
func (t *SonarqubeProject) DBUpsertTx(ctx context.Context, tx Tx, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `sonarqube_project` (`sonarqube_project`.`id`,`sonarqube_project`.`checksum`,`sonarqube_project`.`customer_id`,`sonarqube_project`.`ext_id`,`sonarqube_project`.`key`,`sonarqube_project`.`name`) VALUES (?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `sonarqube_project` (`sonarqube_project`.`id`,`sonarqube_project`.`checksum`,`sonarqube_project`.`customer_id`,`sonarqube_project`.`ext_id`,`sonarqube_project`.`key`,`sonarqube_project`.`name`) VALUES (?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`customer_id`=VALUES(`customer_id`),`ext_id`=VALUES(`ext_id`),`key`=VALUES(`key`),`name`=VALUES(`name`)"
	}
	r, err := tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.ExtID),
		orm.ToSQLString(t.Key),
		orm.ToSQLString(t.Name),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBFindOne will find a SonarqubeProject record in the database with the primary key
func (t *SonarqubeProject) DBFindOne(ctx context.Context, db DB, value string) (bool, error) {
	q := "SELECT `sonarqube_project`.`id`,`sonarqube_project`.`checksum`,`sonarqube_project`.`customer_id`,`sonarqube_project`.`ext_id`,`sonarqube_project`.`key`,`sonarqube_project`.`name` FROM `sonarqube_project` WHERE `id` = ? LIMIT 1"
	row := db.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CustomerID sql.NullString
	var _ExtID sql.NullString
	var _Key sql.NullString
	var _Name sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_ExtID,
		&_Key,
		&_Name,
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
		t.SetExtID(_ExtID.String)
	}
	if _Key.Valid {
		t.SetKey(_Key.String)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	return true, nil
}

// DBFindOneTx will find a SonarqubeProject record in the database with the primary key using the provided transaction
func (t *SonarqubeProject) DBFindOneTx(ctx context.Context, tx Tx, value string) (bool, error) {
	q := "SELECT `sonarqube_project`.`id`,`sonarqube_project`.`checksum`,`sonarqube_project`.`customer_id`,`sonarqube_project`.`ext_id`,`sonarqube_project`.`key`,`sonarqube_project`.`name` FROM `sonarqube_project` WHERE `id` = ? LIMIT 1"
	row := tx.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CustomerID sql.NullString
	var _ExtID sql.NullString
	var _Key sql.NullString
	var _Name sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_ExtID,
		&_Key,
		&_Name,
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
		t.SetExtID(_ExtID.String)
	}
	if _Key.Valid {
		t.SetKey(_Key.String)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	return true, nil
}

// FindSonarqubeProjects will find a SonarqubeProject record in the database with the provided parameters
func FindSonarqubeProjects(ctx context.Context, db DB, _params ...interface{}) ([]*SonarqubeProject, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("customer_id"),
		orm.Column("ext_id"),
		orm.Column("key"),
		orm.Column("name"),
		orm.Table(SonarqubeProjectTableName),
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
	results := make([]*SonarqubeProject, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _ExtID sql.NullString
		var _Key sql.NullString
		var _Name sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_ExtID,
			&_Key,
			&_Name,
		)
		if err != nil {
			return nil, err
		}
		t := &SonarqubeProject{}
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
			t.SetExtID(_ExtID.String)
		}
		if _Key.Valid {
			t.SetKey(_Key.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindSonarqubeProjectsTx will find a SonarqubeProject record in the database with the provided parameters using the provided transaction
func FindSonarqubeProjectsTx(ctx context.Context, tx Tx, _params ...interface{}) ([]*SonarqubeProject, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("customer_id"),
		orm.Column("ext_id"),
		orm.Column("key"),
		orm.Column("name"),
		orm.Table(SonarqubeProjectTableName),
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
	results := make([]*SonarqubeProject, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _ExtID sql.NullString
		var _Key sql.NullString
		var _Name sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_ExtID,
			&_Key,
			&_Name,
		)
		if err != nil {
			return nil, err
		}
		t := &SonarqubeProject{}
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
			t.SetExtID(_ExtID.String)
		}
		if _Key.Valid {
			t.SetKey(_Key.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// DBFind will find a SonarqubeProject record in the database with the provided parameters
func (t *SonarqubeProject) DBFind(ctx context.Context, db DB, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("customer_id"),
		orm.Column("ext_id"),
		orm.Column("key"),
		orm.Column("name"),
		orm.Table(SonarqubeProjectTableName),
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
	var _ExtID sql.NullString
	var _Key sql.NullString
	var _Name sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_ExtID,
		&_Key,
		&_Name,
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
		t.SetExtID(_ExtID.String)
	}
	if _Key.Valid {
		t.SetKey(_Key.String)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	return true, nil
}

// DBFindTx will find a SonarqubeProject record in the database with the provided parameters using the provided transaction
func (t *SonarqubeProject) DBFindTx(ctx context.Context, tx Tx, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("customer_id"),
		orm.Column("ext_id"),
		orm.Column("key"),
		orm.Column("name"),
		orm.Table(SonarqubeProjectTableName),
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
	var _ExtID sql.NullString
	var _Key sql.NullString
	var _Name sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_ExtID,
		&_Key,
		&_Name,
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
		t.SetExtID(_ExtID.String)
	}
	if _Key.Valid {
		t.SetKey(_Key.String)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	return true, nil
}

// CountSonarqubeProjects will find the count of SonarqubeProject records in the database
func CountSonarqubeProjects(ctx context.Context, db DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(SonarqubeProjectTableName),
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

// CountSonarqubeProjectsTx will find the count of SonarqubeProject records in the database using the provided transaction
func CountSonarqubeProjectsTx(ctx context.Context, tx Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(SonarqubeProjectTableName),
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

// DBCount will find the count of SonarqubeProject records in the database
func (t *SonarqubeProject) DBCount(ctx context.Context, db DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(SonarqubeProjectTableName),
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

// DBCountTx will find the count of SonarqubeProject records in the database using the provided transaction
func (t *SonarqubeProject) DBCountTx(ctx context.Context, tx Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(SonarqubeProjectTableName),
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

// DBExists will return true if the SonarqubeProject record exists in the database
func (t *SonarqubeProject) DBExists(ctx context.Context, db DB) (bool, error) {
	q := "SELECT `id` FROM `sonarqube_project` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := db.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// DBExistsTx will return true if the SonarqubeProject record exists in the database using the provided transaction
func (t *SonarqubeProject) DBExistsTx(ctx context.Context, tx Tx) (bool, error) {
	q := "SELECT `id` FROM `sonarqube_project` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := tx.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// PrimaryKeyColumn returns the column name for the primary key
func (t *SonarqubeProject) PrimaryKeyColumn() string {
	return SonarqubeProjectColumnID
}

// PrimaryKeyColumnType returns the primary key column Go type as a string
func (t *SonarqubeProject) PrimaryKeyColumnType() string {
	return "string"
}

// PrimaryKey returns the primary key column value
func (t *SonarqubeProject) PrimaryKey() interface{} {
	return t.ID
}
