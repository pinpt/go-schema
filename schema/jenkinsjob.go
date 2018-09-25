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

var _ Model = (*JenkinsJob)(nil)
var _ CSVWriter = (*JenkinsJob)(nil)
var _ JSONWriter = (*JenkinsJob)(nil)
var _ Checksum = (*JenkinsJob)(nil)

// JenkinsJobTableName is the name of the table in SQL
const JenkinsJobTableName = "jenkins_job"

var JenkinsJobColumns = []string{
	"id",
	"checksum",
	"customer_id",
	"name",
	"git_repo",
	"url",
}

// JenkinsJob table
type JenkinsJob struct {
	Checksum   *string `json:"checksum,omitempty"`
	CustomerID string  `json:"customer_id"`
	GitRepo    *string `json:"git_repo,omitempty"`
	ID         string  `json:"id"`
	Name       string  `json:"name"`
	URL        string  `json:"url"`
}

// TableName returns the SQL table name for JenkinsJob and satifies the Model interface
func (t *JenkinsJob) TableName() string {
	return JenkinsJobTableName
}

// ToCSV will serialize the JenkinsJob instance to a CSV compatible array of strings
func (t *JenkinsJob) ToCSV() []string {
	return []string{
		t.ID,
		t.CalculateChecksum(),
		t.CustomerID,
		t.Name,
		toCSVString(t.GitRepo),
		t.URL,
	}
}

// WriteCSV will serialize the JenkinsJob instance to the writer as CSV and satisfies the CSVWriter interface
func (t *JenkinsJob) WriteCSV(w *csv.Writer) error {
	return w.Write(t.ToCSV())
}

// WriteJSON will serialize the JenkinsJob instance to the writer as JSON and satisfies the JSONWriter interface
func (t *JenkinsJob) WriteJSON(w io.Writer, indent ...bool) error {
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

// NewJenkinsJobReader creates a JSON reader which can read in JenkinsJob objects serialized as JSON either as an array, single object or json new lines
// and writes each JenkinsJob to the channel provided
func NewJenkinsJobReader(r io.Reader, ch chan<- JenkinsJob) error {
	return orm.Deserialize(r, func(buf json.RawMessage) error {
		dec := json.NewDecoder(bytes.NewBuffer(buf))
		e := JenkinsJob{}
		if err := dec.Decode(&e); err != nil {
			return err
		}
		ch <- e
		return nil
	})
}

// NewCSVJenkinsJobReaderDir will read the reader as CSV and emit each record to the channel provided
func NewCSVJenkinsJobReader(r io.Reader, ch chan<- JenkinsJob) error {
	cr := csv.NewReader(r)
	for {
		record, err := cr.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		ch <- JenkinsJob{
			ID:         record[0],
			Checksum:   fromStringPointer(record[1]),
			CustomerID: record[2],
			Name:       record[3],
			GitRepo:    fromStringPointer(record[4]),
			URL:        record[5],
		}
	}
	return nil
}

// NewCSVJenkinsJobReaderFile will read the file as a CSV and emit each record to the channel provided
func NewCSVJenkinsJobReaderFile(fp string, ch chan<- JenkinsJob) error {
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
	return NewCSVJenkinsJobReader(fc, ch)
}

// NewCSVJenkinsJobReaderDir will read the jenkins_job.csv.gz file as a CSV and emit each record to the channel provided
func NewCSVJenkinsJobReaderDir(dir string, ch chan<- JenkinsJob) error {
	return NewCSVJenkinsJobReaderFile(filepath.Join(dir, "jenkins_job.csv.gz"), ch)
}

// JenkinsJobCSVDeduper is a function callback which takes the existing value (a) and the new value (b)
// and the return value should be the one to use (or a new one, if applicable). return nil
// to skip processing of this record
type JenkinsJobCSVDeduper func(a JenkinsJob, b JenkinsJob) *JenkinsJob

// JenkinsJobCSVDedupeDisabled is set on whether the CSV writer should de-dupe values by key
var JenkinsJobCSVDedupeDisabled bool

// NewJenkinsJobCSVWriterSize creates a batch writer that will write each JenkinsJob into a CSV file
// this method will automatically de-duplicate entries using the primary key. if the checksum
// for a newer item with the same primary key doesn't match a previously sent item, the newer
// one will be used
func NewJenkinsJobCSVWriterSize(w io.Writer, size int, dedupers ...JenkinsJobCSVDeduper) (chan JenkinsJob, chan bool, error) {
	cw := csv.NewWriter(w)
	ch := make(chan JenkinsJob, size)
	done := make(chan bool)
	go func() {
		defer func() { done <- true }()
		dodedupe := !JenkinsJobCSVDedupeDisabled
		var kv map[string]*JenkinsJob
		var deduper JenkinsJobCSVDeduper
		if dedupers != nil && len(dedupers) > 0 {
			deduper = dedupers[0]
			dodedupe = true
		}
		if dodedupe {
			kv = make(map[string]*JenkinsJob)
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

// JenkinsJobCSVDefaultSize is the default channel buffer size if not provided
var JenkinsJobCSVDefaultSize = 100

// NewJenkinsJobCSVWriter creates a batch writer that will write each JenkinsJob into a CSV file
func NewJenkinsJobCSVWriter(w io.Writer, dedupers ...JenkinsJobCSVDeduper) (chan JenkinsJob, chan bool, error) {
	return NewJenkinsJobCSVWriterSize(w, JenkinsJobCSVDefaultSize, dedupers...)
}

// NewJenkinsJobCSVWriterDir creates a batch writer that will write each JenkinsJob into a CSV file named jenkins_job.csv.gz in dir
func NewJenkinsJobCSVWriterDir(dir string, dedupers ...JenkinsJobCSVDeduper) (chan JenkinsJob, chan bool, error) {
	return NewJenkinsJobCSVWriterFile(filepath.Join(dir, "jenkins_job.csv.gz"), dedupers...)
}

// NewJenkinsJobCSVWriterFile creates a batch writer that will write each JenkinsJob into a CSV file
func NewJenkinsJobCSVWriterFile(fn string, dedupers ...JenkinsJobCSVDeduper) (chan JenkinsJob, chan bool, error) {
	f, err := os.Create(fn)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening CSV file %s. %v", fn, err)
	}
	var fc io.WriteCloser = f
	if filepath.Ext(fn) == ".gz" {
		w, _ := gzip.NewWriterLevel(f, gzip.BestCompression)
		fc = w
	}
	ch, done, err := NewJenkinsJobCSVWriter(fc, dedupers...)
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

type JenkinsJobDBAction func(ctx context.Context, db *sql.DB, record JenkinsJob) error

// NewJenkinsJobDBWriterSize creates a DB writer that will write each issue into the DB
func NewJenkinsJobDBWriterSize(ctx context.Context, db *sql.DB, errors chan<- error, size int, actions ...JenkinsJobDBAction) (chan JenkinsJob, chan bool, error) {
	ch := make(chan JenkinsJob, size)
	done := make(chan bool)
	var action JenkinsJobDBAction
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

// NewJenkinsJobDBWriter creates a DB writer that will write each issue into the DB
func NewJenkinsJobDBWriter(ctx context.Context, db *sql.DB, errors chan<- error, actions ...JenkinsJobDBAction) (chan JenkinsJob, chan bool, error) {
	return NewJenkinsJobDBWriterSize(ctx, db, errors, 100, actions...)
}

// JenkinsJobColumnID is the ID SQL column name for the JenkinsJob table
const JenkinsJobColumnID = "id"

// JenkinsJobEscapedColumnID is the escaped ID SQL column name for the JenkinsJob table
const JenkinsJobEscapedColumnID = "`id`"

// JenkinsJobColumnChecksum is the Checksum SQL column name for the JenkinsJob table
const JenkinsJobColumnChecksum = "checksum"

// JenkinsJobEscapedColumnChecksum is the escaped Checksum SQL column name for the JenkinsJob table
const JenkinsJobEscapedColumnChecksum = "`checksum`"

// JenkinsJobColumnCustomerID is the CustomerID SQL column name for the JenkinsJob table
const JenkinsJobColumnCustomerID = "customer_id"

// JenkinsJobEscapedColumnCustomerID is the escaped CustomerID SQL column name for the JenkinsJob table
const JenkinsJobEscapedColumnCustomerID = "`customer_id`"

// JenkinsJobColumnName is the Name SQL column name for the JenkinsJob table
const JenkinsJobColumnName = "name"

// JenkinsJobEscapedColumnName is the escaped Name SQL column name for the JenkinsJob table
const JenkinsJobEscapedColumnName = "`name`"

// JenkinsJobColumnGitRepo is the GitRepo SQL column name for the JenkinsJob table
const JenkinsJobColumnGitRepo = "git_repo"

// JenkinsJobEscapedColumnGitRepo is the escaped GitRepo SQL column name for the JenkinsJob table
const JenkinsJobEscapedColumnGitRepo = "`git_repo`"

// JenkinsJobColumnURL is the URL SQL column name for the JenkinsJob table
const JenkinsJobColumnURL = "url"

// JenkinsJobEscapedColumnURL is the escaped URL SQL column name for the JenkinsJob table
const JenkinsJobEscapedColumnURL = "`url`"

// GetID will return the JenkinsJob ID value
func (t *JenkinsJob) GetID() string {
	return t.ID
}

// SetID will set the JenkinsJob ID value
func (t *JenkinsJob) SetID(v string) {
	t.ID = v
}

// FindJenkinsJobByID will find a JenkinsJob by ID
func FindJenkinsJobByID(ctx context.Context, db *sql.DB, value string) (*JenkinsJob, error) {
	q := "SELECT `jenkins_job`.`id`,`jenkins_job`.`checksum`,`jenkins_job`.`customer_id`,`jenkins_job`.`name`,`jenkins_job`.`git_repo`,`jenkins_job`.`url` FROM `jenkins_job` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CustomerID sql.NullString
	var _Name sql.NullString
	var _GitRepo sql.NullString
	var _URL sql.NullString
	err := db.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_Name,
		&_GitRepo,
		&_URL,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &JenkinsJob{}
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
	if _GitRepo.Valid {
		t.SetGitRepo(_GitRepo.String)
	}
	if _URL.Valid {
		t.SetURL(_URL.String)
	}
	return t, nil
}

// FindJenkinsJobByIDTx will find a JenkinsJob by ID using the provided transaction
func FindJenkinsJobByIDTx(ctx context.Context, tx *sql.Tx, value string) (*JenkinsJob, error) {
	q := "SELECT `jenkins_job`.`id`,`jenkins_job`.`checksum`,`jenkins_job`.`customer_id`,`jenkins_job`.`name`,`jenkins_job`.`git_repo`,`jenkins_job`.`url` FROM `jenkins_job` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CustomerID sql.NullString
	var _Name sql.NullString
	var _GitRepo sql.NullString
	var _URL sql.NullString
	err := tx.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_Name,
		&_GitRepo,
		&_URL,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &JenkinsJob{}
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
	if _GitRepo.Valid {
		t.SetGitRepo(_GitRepo.String)
	}
	if _URL.Valid {
		t.SetURL(_URL.String)
	}
	return t, nil
}

// GetChecksum will return the JenkinsJob Checksum value
func (t *JenkinsJob) GetChecksum() string {
	if t.Checksum == nil {
		return ""
	}
	return *t.Checksum
}

// SetChecksum will set the JenkinsJob Checksum value
func (t *JenkinsJob) SetChecksum(v string) {
	t.Checksum = &v
}

// GetCustomerID will return the JenkinsJob CustomerID value
func (t *JenkinsJob) GetCustomerID() string {
	return t.CustomerID
}

// SetCustomerID will set the JenkinsJob CustomerID value
func (t *JenkinsJob) SetCustomerID(v string) {
	t.CustomerID = v
}

// FindJenkinsJobsByCustomerID will find all JenkinsJobs by the CustomerID value
func FindJenkinsJobsByCustomerID(ctx context.Context, db *sql.DB, value string) ([]*JenkinsJob, error) {
	q := "SELECT `jenkins_job`.`id`,`jenkins_job`.`checksum`,`jenkins_job`.`customer_id`,`jenkins_job`.`name`,`jenkins_job`.`git_repo`,`jenkins_job`.`url` FROM `jenkins_job` WHERE `customer_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JenkinsJob, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _Name sql.NullString
		var _GitRepo sql.NullString
		var _URL sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_Name,
			&_GitRepo,
			&_URL,
		)
		if err != nil {
			return nil, err
		}
		t := &JenkinsJob{}
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
		if _GitRepo.Valid {
			t.SetGitRepo(_GitRepo.String)
		}
		if _URL.Valid {
			t.SetURL(_URL.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindJenkinsJobsByCustomerIDTx will find all JenkinsJobs by the CustomerID value using the provided transaction
func FindJenkinsJobsByCustomerIDTx(ctx context.Context, tx *sql.Tx, value string) ([]*JenkinsJob, error) {
	q := "SELECT `jenkins_job`.`id`,`jenkins_job`.`checksum`,`jenkins_job`.`customer_id`,`jenkins_job`.`name`,`jenkins_job`.`git_repo`,`jenkins_job`.`url` FROM `jenkins_job` WHERE `customer_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JenkinsJob, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _Name sql.NullString
		var _GitRepo sql.NullString
		var _URL sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_Name,
			&_GitRepo,
			&_URL,
		)
		if err != nil {
			return nil, err
		}
		t := &JenkinsJob{}
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
		if _GitRepo.Valid {
			t.SetGitRepo(_GitRepo.String)
		}
		if _URL.Valid {
			t.SetURL(_URL.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetName will return the JenkinsJob Name value
func (t *JenkinsJob) GetName() string {
	return t.Name
}

// SetName will set the JenkinsJob Name value
func (t *JenkinsJob) SetName(v string) {
	t.Name = v
}

// FindJenkinsJobsByName will find all JenkinsJobs by the Name value
func FindJenkinsJobsByName(ctx context.Context, db *sql.DB, value string) ([]*JenkinsJob, error) {
	q := "SELECT `jenkins_job`.`id`,`jenkins_job`.`checksum`,`jenkins_job`.`customer_id`,`jenkins_job`.`name`,`jenkins_job`.`git_repo`,`jenkins_job`.`url` FROM `jenkins_job` WHERE `name` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JenkinsJob, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _Name sql.NullString
		var _GitRepo sql.NullString
		var _URL sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_Name,
			&_GitRepo,
			&_URL,
		)
		if err != nil {
			return nil, err
		}
		t := &JenkinsJob{}
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
		if _GitRepo.Valid {
			t.SetGitRepo(_GitRepo.String)
		}
		if _URL.Valid {
			t.SetURL(_URL.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindJenkinsJobsByNameTx will find all JenkinsJobs by the Name value using the provided transaction
func FindJenkinsJobsByNameTx(ctx context.Context, tx *sql.Tx, value string) ([]*JenkinsJob, error) {
	q := "SELECT `jenkins_job`.`id`,`jenkins_job`.`checksum`,`jenkins_job`.`customer_id`,`jenkins_job`.`name`,`jenkins_job`.`git_repo`,`jenkins_job`.`url` FROM `jenkins_job` WHERE `name` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JenkinsJob, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _Name sql.NullString
		var _GitRepo sql.NullString
		var _URL sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_Name,
			&_GitRepo,
			&_URL,
		)
		if err != nil {
			return nil, err
		}
		t := &JenkinsJob{}
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
		if _GitRepo.Valid {
			t.SetGitRepo(_GitRepo.String)
		}
		if _URL.Valid {
			t.SetURL(_URL.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetGitRepo will return the JenkinsJob GitRepo value
func (t *JenkinsJob) GetGitRepo() string {
	if t.GitRepo == nil {
		return ""
	}
	return *t.GitRepo
}

// SetGitRepo will set the JenkinsJob GitRepo value
func (t *JenkinsJob) SetGitRepo(v string) {
	t.GitRepo = &v
}

// GetURL will return the JenkinsJob URL value
func (t *JenkinsJob) GetURL() string {
	return t.URL
}

// SetURL will set the JenkinsJob URL value
func (t *JenkinsJob) SetURL(v string) {
	t.URL = v
}

func (t *JenkinsJob) toTimestamp(value time.Time) *timestamp.Timestamp {
	ts, _ := ptypes.TimestampProto(value)
	return ts
}

// DBCreateJenkinsJobTable will create the JenkinsJob table
func DBCreateJenkinsJobTable(ctx context.Context, db *sql.DB) error {
	q := "CREATE TABLE `jenkins_job` (`id` VARCHAR(64) NOT NULL PRIMARY KEY,`checksum` CHAR(64),`customer_id` VARCHAR(64) NOT NULL,`name`VARCHAR(255) NOT NULL,`git_repo` VARCHAR(255),`url` VARCHAR(255) NOT NULL,INDEX jenkins_job_customer_id_index (`customer_id`),INDEX jenkins_job_name_index (`name`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBCreateJenkinsJobTableTx will create the JenkinsJob table using the provided transction
func DBCreateJenkinsJobTableTx(ctx context.Context, tx *sql.Tx) error {
	q := "CREATE TABLE `jenkins_job` (`id` VARCHAR(64) NOT NULL PRIMARY KEY,`checksum` CHAR(64),`customer_id` VARCHAR(64) NOT NULL,`name`VARCHAR(255) NOT NULL,`git_repo` VARCHAR(255),`url` VARCHAR(255) NOT NULL,INDEX jenkins_job_customer_id_index (`customer_id`),INDEX jenkins_job_name_index (`name`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// DBDropJenkinsJobTable will drop the JenkinsJob table
func DBDropJenkinsJobTable(ctx context.Context, db *sql.DB) error {
	q := "DROP TABLE IF EXISTS `jenkins_job`"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBDropJenkinsJobTableTx will drop the JenkinsJob table using the provided transaction
func DBDropJenkinsJobTableTx(ctx context.Context, tx *sql.Tx) error {
	q := "DROP TABLE IF EXISTS `jenkins_job`"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// CalculateChecksum will calculate a checksum of the SHA1 of all field values
func (t *JenkinsJob) CalculateChecksum() string {
	return orm.HashStrings(
		orm.ToString(t.ID),
		orm.ToString(t.CustomerID),
		orm.ToString(t.Name),
		orm.ToString(t.GitRepo),
		orm.ToString(t.URL),
	)
}

// DBCreate will create a new JenkinsJob record in the database
func (t *JenkinsJob) DBCreate(ctx context.Context, db *sql.DB) (sql.Result, error) {
	q := "INSERT INTO `jenkins_job` (`jenkins_job`.`id`,`jenkins_job`.`checksum`,`jenkins_job`.`customer_id`,`jenkins_job`.`name`,`jenkins_job`.`git_repo`,`jenkins_job`.`url`) VALUES (?,?,?,?,?,?)"
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
		orm.ToSQLString(t.GitRepo),
		orm.ToSQLString(t.URL),
	)
}

// DBCreateTx will create a new JenkinsJob record in the database using the provided transaction
func (t *JenkinsJob) DBCreateTx(ctx context.Context, tx *sql.Tx) (sql.Result, error) {
	q := "INSERT INTO `jenkins_job` (`jenkins_job`.`id`,`jenkins_job`.`checksum`,`jenkins_job`.`customer_id`,`jenkins_job`.`name`,`jenkins_job`.`git_repo`,`jenkins_job`.`url`) VALUES (?,?,?,?,?,?)"
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
		orm.ToSQLString(t.GitRepo),
		orm.ToSQLString(t.URL),
	)
}

// DBCreateIgnoreDuplicate will upsert the JenkinsJob record in the database
func (t *JenkinsJob) DBCreateIgnoreDuplicate(ctx context.Context, db *sql.DB) (sql.Result, error) {
	q := "INSERT INTO `jenkins_job` (`jenkins_job`.`id`,`jenkins_job`.`checksum`,`jenkins_job`.`customer_id`,`jenkins_job`.`name`,`jenkins_job`.`git_repo`,`jenkins_job`.`url`) VALUES (?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
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
		orm.ToSQLString(t.GitRepo),
		orm.ToSQLString(t.URL),
	)
}

// DBCreateIgnoreDuplicateTx will upsert the JenkinsJob record in the database using the provided transaction
func (t *JenkinsJob) DBCreateIgnoreDuplicateTx(ctx context.Context, tx *sql.Tx) (sql.Result, error) {
	q := "INSERT INTO `jenkins_job` (`jenkins_job`.`id`,`jenkins_job`.`checksum`,`jenkins_job`.`customer_id`,`jenkins_job`.`name`,`jenkins_job`.`git_repo`,`jenkins_job`.`url`) VALUES (?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
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
		orm.ToSQLString(t.GitRepo),
		orm.ToSQLString(t.URL),
	)
}

// DeleteAllJenkinsJobs deletes all JenkinsJob records in the database with optional filters
func DeleteAllJenkinsJobs(ctx context.Context, db *sql.DB, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(JenkinsJobTableName),
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

// DeleteAllJenkinsJobsTx deletes all JenkinsJob records in the database with optional filters using the provided transaction
func DeleteAllJenkinsJobsTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(JenkinsJobTableName),
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

// DBDelete will delete this JenkinsJob record in the database
func (t *JenkinsJob) DBDelete(ctx context.Context, db *sql.DB) (bool, error) {
	q := "DELETE FROM `jenkins_job` WHERE `id` = ?"
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

// DBDeleteTx will delete this JenkinsJob record in the database using the provided transaction
func (t *JenkinsJob) DBDeleteTx(ctx context.Context, tx *sql.Tx) (bool, error) {
	q := "DELETE FROM `jenkins_job` WHERE `id` = ?"
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

// DBUpdate will update the JenkinsJob record in the database
func (t *JenkinsJob) DBUpdate(ctx context.Context, db *sql.DB) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `jenkins_job` SET `checksum`=?,`customer_id`=?,`name`=?,`git_repo`=?,`url`=? WHERE `id`=?"
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.GitRepo),
		orm.ToSQLString(t.URL),
		orm.ToSQLString(t.ID),
	)
}

// DBUpdateTx will update the JenkinsJob record in the database using the provided transaction
func (t *JenkinsJob) DBUpdateTx(ctx context.Context, tx *sql.Tx) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `jenkins_job` SET `checksum`=?,`customer_id`=?,`name`=?,`git_repo`=?,`url`=? WHERE `id`=?"
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.GitRepo),
		orm.ToSQLString(t.URL),
		orm.ToSQLString(t.ID),
	)
}

// DBUpsert will upsert the JenkinsJob record in the database
func (t *JenkinsJob) DBUpsert(ctx context.Context, db *sql.DB, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `jenkins_job` (`jenkins_job`.`id`,`jenkins_job`.`checksum`,`jenkins_job`.`customer_id`,`jenkins_job`.`name`,`jenkins_job`.`git_repo`,`jenkins_job`.`url`) VALUES (?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `jenkins_job` (`jenkins_job`.`id`,`jenkins_job`.`checksum`,`jenkins_job`.`customer_id`,`jenkins_job`.`name`,`jenkins_job`.`git_repo`,`jenkins_job`.`url`) VALUES (?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`customer_id`=VALUES(`customer_id`),`name`=VALUES(`name`),`git_repo`=VALUES(`git_repo`),`url`=VALUES(`url`)"
	}
	r, err := db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.GitRepo),
		orm.ToSQLString(t.URL),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBUpsertTx will upsert the JenkinsJob record in the database using the provided transaction
func (t *JenkinsJob) DBUpsertTx(ctx context.Context, tx *sql.Tx, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `jenkins_job` (`jenkins_job`.`id`,`jenkins_job`.`checksum`,`jenkins_job`.`customer_id`,`jenkins_job`.`name`,`jenkins_job`.`git_repo`,`jenkins_job`.`url`) VALUES (?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `jenkins_job` (`jenkins_job`.`id`,`jenkins_job`.`checksum`,`jenkins_job`.`customer_id`,`jenkins_job`.`name`,`jenkins_job`.`git_repo`,`jenkins_job`.`url`) VALUES (?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`customer_id`=VALUES(`customer_id`),`name`=VALUES(`name`),`git_repo`=VALUES(`git_repo`),`url`=VALUES(`url`)"
	}
	r, err := tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.GitRepo),
		orm.ToSQLString(t.URL),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBFindOne will find a JenkinsJob record in the database with the primary key
func (t *JenkinsJob) DBFindOne(ctx context.Context, db *sql.DB, value string) (bool, error) {
	q := "SELECT `jenkins_job`.`id`,`jenkins_job`.`checksum`,`jenkins_job`.`customer_id`,`jenkins_job`.`name`,`jenkins_job`.`git_repo`,`jenkins_job`.`url` FROM `jenkins_job` WHERE `id` = ? LIMIT 1"
	row := db.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CustomerID sql.NullString
	var _Name sql.NullString
	var _GitRepo sql.NullString
	var _URL sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_Name,
		&_GitRepo,
		&_URL,
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
	if _GitRepo.Valid {
		t.SetGitRepo(_GitRepo.String)
	}
	if _URL.Valid {
		t.SetURL(_URL.String)
	}
	return true, nil
}

// DBFindOneTx will find a JenkinsJob record in the database with the primary key using the provided transaction
func (t *JenkinsJob) DBFindOneTx(ctx context.Context, tx *sql.Tx, value string) (bool, error) {
	q := "SELECT `jenkins_job`.`id`,`jenkins_job`.`checksum`,`jenkins_job`.`customer_id`,`jenkins_job`.`name`,`jenkins_job`.`git_repo`,`jenkins_job`.`url` FROM `jenkins_job` WHERE `id` = ? LIMIT 1"
	row := tx.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CustomerID sql.NullString
	var _Name sql.NullString
	var _GitRepo sql.NullString
	var _URL sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_Name,
		&_GitRepo,
		&_URL,
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
	if _GitRepo.Valid {
		t.SetGitRepo(_GitRepo.String)
	}
	if _URL.Valid {
		t.SetURL(_URL.String)
	}
	return true, nil
}

// FindJenkinsJobs will find a JenkinsJob record in the database with the provided parameters
func FindJenkinsJobs(ctx context.Context, db *sql.DB, _params ...interface{}) ([]*JenkinsJob, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("customer_id"),
		orm.Column("name"),
		orm.Column("git_repo"),
		orm.Column("url"),
		orm.Table(JenkinsJobTableName),
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
	results := make([]*JenkinsJob, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _Name sql.NullString
		var _GitRepo sql.NullString
		var _URL sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_Name,
			&_GitRepo,
			&_URL,
		)
		if err != nil {
			return nil, err
		}
		t := &JenkinsJob{}
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
		if _GitRepo.Valid {
			t.SetGitRepo(_GitRepo.String)
		}
		if _URL.Valid {
			t.SetURL(_URL.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindJenkinsJobsTx will find a JenkinsJob record in the database with the provided parameters using the provided transaction
func FindJenkinsJobsTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) ([]*JenkinsJob, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("customer_id"),
		orm.Column("name"),
		orm.Column("git_repo"),
		orm.Column("url"),
		orm.Table(JenkinsJobTableName),
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
	results := make([]*JenkinsJob, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _Name sql.NullString
		var _GitRepo sql.NullString
		var _URL sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_Name,
			&_GitRepo,
			&_URL,
		)
		if err != nil {
			return nil, err
		}
		t := &JenkinsJob{}
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
		if _GitRepo.Valid {
			t.SetGitRepo(_GitRepo.String)
		}
		if _URL.Valid {
			t.SetURL(_URL.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// DBFind will find a JenkinsJob record in the database with the provided parameters
func (t *JenkinsJob) DBFind(ctx context.Context, db *sql.DB, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("customer_id"),
		orm.Column("name"),
		orm.Column("git_repo"),
		orm.Column("url"),
		orm.Table(JenkinsJobTableName),
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
	var _GitRepo sql.NullString
	var _URL sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_Name,
		&_GitRepo,
		&_URL,
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
	if _GitRepo.Valid {
		t.SetGitRepo(_GitRepo.String)
	}
	if _URL.Valid {
		t.SetURL(_URL.String)
	}
	return true, nil
}

// DBFindTx will find a JenkinsJob record in the database with the provided parameters using the provided transaction
func (t *JenkinsJob) DBFindTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("customer_id"),
		orm.Column("name"),
		orm.Column("git_repo"),
		orm.Column("url"),
		orm.Table(JenkinsJobTableName),
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
	var _GitRepo sql.NullString
	var _URL sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_Name,
		&_GitRepo,
		&_URL,
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
	if _GitRepo.Valid {
		t.SetGitRepo(_GitRepo.String)
	}
	if _URL.Valid {
		t.SetURL(_URL.String)
	}
	return true, nil
}

// CountJenkinsJobs will find the count of JenkinsJob records in the database
func CountJenkinsJobs(ctx context.Context, db *sql.DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(JenkinsJobTableName),
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

// CountJenkinsJobsTx will find the count of JenkinsJob records in the database using the provided transaction
func CountJenkinsJobsTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(JenkinsJobTableName),
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

// DBCount will find the count of JenkinsJob records in the database
func (t *JenkinsJob) DBCount(ctx context.Context, db *sql.DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(JenkinsJobTableName),
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

// DBCountTx will find the count of JenkinsJob records in the database using the provided transaction
func (t *JenkinsJob) DBCountTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(JenkinsJobTableName),
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

// DBExists will return true if the JenkinsJob record exists in the database
func (t *JenkinsJob) DBExists(ctx context.Context, db *sql.DB) (bool, error) {
	q := "SELECT `id` FROM `jenkins_job` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := db.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// DBExistsTx will return true if the JenkinsJob record exists in the database using the provided transaction
func (t *JenkinsJob) DBExistsTx(ctx context.Context, tx *sql.Tx) (bool, error) {
	q := "SELECT `id` FROM `jenkins_job` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := tx.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// PrimaryKeyColumn returns the column name for the primary key
func (t *JenkinsJob) PrimaryKeyColumn() string {
	return JenkinsJobColumnID
}

// PrimaryKeyColumnType returns the primary key column Go type as a string
func (t *JenkinsJob) PrimaryKeyColumnType() string {
	return "string"
}

// PrimaryKey returns the primary key column value
func (t *JenkinsJob) PrimaryKey() interface{} {
	return t.ID
}
