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

var _ Model = (*JenkinsBuild)(nil)
var _ CSVWriter = (*JenkinsBuild)(nil)
var _ JSONWriter = (*JenkinsBuild)(nil)
var _ Checksum = (*JenkinsBuild)(nil)

// JenkinsBuildTableName is the name of the table in SQL
const JenkinsBuildTableName = "jenkins_build"

var JenkinsBuildColumns = []string{
	"id",
	"checksum",
	"customer_id",
	"jenkins_build_number",
	"job_id",
	"result",
	"created_at",
	"finished_at",
	"sha",
	"branch",
	"url",
}

type JenkinsBuild_JenkinsBuildResult string

const (
	JenkinsBuildResult_RESULT_UNKNOWN JenkinsBuild_JenkinsBuildResult = "result_unknown"
	JenkinsBuildResult_SUCCESS        JenkinsBuild_JenkinsBuildResult = "success"
	JenkinsBuildResult_FAILURE        JenkinsBuild_JenkinsBuildResult = "failure"
	JenkinsBuildResult_ABORTED        JenkinsBuild_JenkinsBuildResult = "aborted"
)

func (x JenkinsBuild_JenkinsBuildResult) String() string {
	return string(x)
}

func enumJenkinsBuildResultToString(v *JenkinsBuild_JenkinsBuildResult) string {
	if v == nil {
		return ""
	}
	return v.String()
}

func toJenkinsBuildResult(v string) *JenkinsBuild_JenkinsBuildResult {
	var ev *JenkinsBuild_JenkinsBuildResult
	switch v {
	case "RESULT_UNKNOWN", "result_unknown":
		{
			v := JenkinsBuildResult_RESULT_UNKNOWN
			ev = &v
		}
	case "SUCCESS", "success":
		{
			v := JenkinsBuildResult_SUCCESS
			ev = &v
		}
	case "FAILURE", "failure":
		{
			v := JenkinsBuildResult_FAILURE
			ev = &v
		}
	case "ABORTED", "aborted":
		{
			v := JenkinsBuildResult_ABORTED
			ev = &v
		}
	}
	return ev
}

// JenkinsBuild table
type JenkinsBuild struct {
	Branch             *string                         `json:"branch,omitempty"`
	Checksum           *string                         `json:"checksum,omitempty"`
	CreatedAt          int64                           `json:"created_at"`
	CustomerID         string                          `json:"customer_id"`
	FinishedAt         *int64                          `json:"finished_at,omitempty"`
	ID                 string                          `json:"id"`
	JenkinsBuildNumber int64                           `json:"jenkins_build_number"`
	JobID              string                          `json:"job_id"`
	Result             JenkinsBuild_JenkinsBuildResult `json:"result"`
	Sha                *string                         `json:"sha,omitempty"`
	URL                string                          `json:"url"`
}

// TableName returns the SQL table name for JenkinsBuild and satifies the Model interface
func (t *JenkinsBuild) TableName() string {
	return JenkinsBuildTableName
}

// ToCSV will serialize the JenkinsBuild instance to a CSV compatible array of strings
func (t *JenkinsBuild) ToCSV() []string {
	return []string{
		t.ID,
		t.CalculateChecksum(),
		t.CustomerID,
		toCSVString(t.JenkinsBuildNumber),
		t.JobID,
		toCSVString(t.Result),
		toCSVString(t.CreatedAt),
		toCSVString(t.FinishedAt),
		toCSVString(t.Sha),
		toCSVString(t.Branch),
		t.URL,
	}
}

// WriteCSV will serialize the JenkinsBuild instance to the writer as CSV and satisfies the CSVWriter interface
func (t *JenkinsBuild) WriteCSV(w *csv.Writer) error {
	return w.Write(t.ToCSV())
}

// WriteJSON will serialize the JenkinsBuild instance to the writer as JSON and satisfies the JSONWriter interface
func (t *JenkinsBuild) WriteJSON(w io.Writer, indent ...bool) error {
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

// NewJenkinsBuildReader creates a JSON reader which can read in JenkinsBuild objects serialized as JSON either as an array, single object or json new lines
// and writes each JenkinsBuild to the channel provided
func NewJenkinsBuildReader(r io.Reader, ch chan<- JenkinsBuild) error {
	return orm.Deserialize(r, func(buf json.RawMessage) error {
		dec := json.NewDecoder(bytes.NewBuffer(buf))
		e := JenkinsBuild{}
		if err := dec.Decode(&e); err != nil {
			return err
		}
		ch <- e
		return nil
	})
}

// NewCSVJenkinsBuildReaderDir will read the reader as CSV and emit each record to the channel provided
func NewCSVJenkinsBuildReader(r io.Reader, ch chan<- JenkinsBuild) error {
	cr := csv.NewReader(r)
	for {
		record, err := cr.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		ch <- JenkinsBuild{
			ID:                 record[0],
			Checksum:           fromStringPointer(record[1]),
			CustomerID:         record[2],
			JenkinsBuildNumber: fromCSVInt64(record[3]),
			JobID:              record[4],
			Result:             *toJenkinsBuildResult(record[5]),
			CreatedAt:          fromCSVInt64(record[6]),
			FinishedAt:         fromCSVInt64Pointer(record[7]),
			Sha:                fromStringPointer(record[8]),
			Branch:             fromStringPointer(record[9]),
			URL:                record[10],
		}
	}
	return nil
}

// NewCSVJenkinsBuildReaderFile will read the file as a CSV and emit each record to the channel provided
func NewCSVJenkinsBuildReaderFile(fp string, ch chan<- JenkinsBuild) error {
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
	return NewCSVJenkinsBuildReader(fc, ch)
}

// NewCSVJenkinsBuildReaderDir will read the jenkins_build.csv.gz file as a CSV and emit each record to the channel provided
func NewCSVJenkinsBuildReaderDir(dir string, ch chan<- JenkinsBuild) error {
	return NewCSVJenkinsBuildReaderFile(filepath.Join(dir, "jenkins_build.csv.gz"), ch)
}

// JenkinsBuildCSVDeduper is a function callback which takes the existing value (a) and the new value (b)
// and the return value should be the one to use (or a new one, if applicable). return nil
// to skip processing of this record
type JenkinsBuildCSVDeduper func(a JenkinsBuild, b JenkinsBuild) *JenkinsBuild

// JenkinsBuildCSVDedupeDisabled is set on whether the CSV writer should de-dupe values by key
var JenkinsBuildCSVDedupeDisabled bool

// NewJenkinsBuildCSVWriterSize creates a batch writer that will write each JenkinsBuild into a CSV file
// this method will automatically de-duplicate entries using the primary key. if the checksum
// for a newer item with the same primary key doesn't match a previously sent item, the newer
// one will be used
func NewJenkinsBuildCSVWriterSize(w io.Writer, size int, dedupers ...JenkinsBuildCSVDeduper) (chan JenkinsBuild, chan bool, error) {
	cw := csv.NewWriter(w)
	ch := make(chan JenkinsBuild, size)
	done := make(chan bool)
	go func() {
		defer func() { done <- true }()
		dodedupe := !JenkinsBuildCSVDedupeDisabled
		var kv map[string]*JenkinsBuild
		var deduper JenkinsBuildCSVDeduper
		if dedupers != nil && len(dedupers) > 0 {
			deduper = dedupers[0]
			dodedupe = true
		}
		if dodedupe {
			kv = make(map[string]*JenkinsBuild)
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

// JenkinsBuildCSVDefaultSize is the default channel buffer size if not provided
var JenkinsBuildCSVDefaultSize = 100

// NewJenkinsBuildCSVWriter creates a batch writer that will write each JenkinsBuild into a CSV file
func NewJenkinsBuildCSVWriter(w io.Writer, dedupers ...JenkinsBuildCSVDeduper) (chan JenkinsBuild, chan bool, error) {
	return NewJenkinsBuildCSVWriterSize(w, JenkinsBuildCSVDefaultSize, dedupers...)
}

// NewJenkinsBuildCSVWriterDir creates a batch writer that will write each JenkinsBuild into a CSV file named jenkins_build.csv.gz in dir
func NewJenkinsBuildCSVWriterDir(dir string, dedupers ...JenkinsBuildCSVDeduper) (chan JenkinsBuild, chan bool, error) {
	return NewJenkinsBuildCSVWriterFile(filepath.Join(dir, "jenkins_build.csv.gz"), dedupers...)
}

// NewJenkinsBuildCSVWriterFile creates a batch writer that will write each JenkinsBuild into a CSV file
func NewJenkinsBuildCSVWriterFile(fn string, dedupers ...JenkinsBuildCSVDeduper) (chan JenkinsBuild, chan bool, error) {
	f, err := os.Create(fn)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening CSV file %s. %v", fn, err)
	}
	var fc io.WriteCloser = f
	if filepath.Ext(fn) == ".gz" {
		w, _ := gzip.NewWriterLevel(f, gzip.BestCompression)
		fc = w
	}
	ch, done, err := NewJenkinsBuildCSVWriter(fc, dedupers...)
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

type JenkinsBuildDBAction func(ctx context.Context, db *sql.DB, record JenkinsBuild) error

// NewJenkinsBuildDBWriterSize creates a DB writer that will write each issue into the DB
func NewJenkinsBuildDBWriterSize(ctx context.Context, db *sql.DB, errors chan<- error, size int, actions ...JenkinsBuildDBAction) (chan JenkinsBuild, chan bool, error) {
	ch := make(chan JenkinsBuild, size)
	done := make(chan bool)
	var action JenkinsBuildDBAction
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

// NewJenkinsBuildDBWriter creates a DB writer that will write each issue into the DB
func NewJenkinsBuildDBWriter(ctx context.Context, db *sql.DB, errors chan<- error, actions ...JenkinsBuildDBAction) (chan JenkinsBuild, chan bool, error) {
	return NewJenkinsBuildDBWriterSize(ctx, db, errors, 100, actions...)
}

// JenkinsBuildColumnID is the ID SQL column name for the JenkinsBuild table
const JenkinsBuildColumnID = "id"

// JenkinsBuildEscapedColumnID is the escaped ID SQL column name for the JenkinsBuild table
const JenkinsBuildEscapedColumnID = "`id`"

// JenkinsBuildColumnChecksum is the Checksum SQL column name for the JenkinsBuild table
const JenkinsBuildColumnChecksum = "checksum"

// JenkinsBuildEscapedColumnChecksum is the escaped Checksum SQL column name for the JenkinsBuild table
const JenkinsBuildEscapedColumnChecksum = "`checksum`"

// JenkinsBuildColumnCustomerID is the CustomerID SQL column name for the JenkinsBuild table
const JenkinsBuildColumnCustomerID = "customer_id"

// JenkinsBuildEscapedColumnCustomerID is the escaped CustomerID SQL column name for the JenkinsBuild table
const JenkinsBuildEscapedColumnCustomerID = "`customer_id`"

// JenkinsBuildColumnJenkinsBuildNumber is the JenkinsBuildNumber SQL column name for the JenkinsBuild table
const JenkinsBuildColumnJenkinsBuildNumber = "jenkins_build_number"

// JenkinsBuildEscapedColumnJenkinsBuildNumber is the escaped JenkinsBuildNumber SQL column name for the JenkinsBuild table
const JenkinsBuildEscapedColumnJenkinsBuildNumber = "`jenkins_build_number`"

// JenkinsBuildColumnJobID is the JobID SQL column name for the JenkinsBuild table
const JenkinsBuildColumnJobID = "job_id"

// JenkinsBuildEscapedColumnJobID is the escaped JobID SQL column name for the JenkinsBuild table
const JenkinsBuildEscapedColumnJobID = "`job_id`"

// JenkinsBuildColumnResult is the Result SQL column name for the JenkinsBuild table
const JenkinsBuildColumnResult = "result"

// JenkinsBuildEscapedColumnResult is the escaped Result SQL column name for the JenkinsBuild table
const JenkinsBuildEscapedColumnResult = "`result`"

// JenkinsBuildColumnCreatedAt is the CreatedAt SQL column name for the JenkinsBuild table
const JenkinsBuildColumnCreatedAt = "created_at"

// JenkinsBuildEscapedColumnCreatedAt is the escaped CreatedAt SQL column name for the JenkinsBuild table
const JenkinsBuildEscapedColumnCreatedAt = "`created_at`"

// JenkinsBuildColumnFinishedAt is the FinishedAt SQL column name for the JenkinsBuild table
const JenkinsBuildColumnFinishedAt = "finished_at"

// JenkinsBuildEscapedColumnFinishedAt is the escaped FinishedAt SQL column name for the JenkinsBuild table
const JenkinsBuildEscapedColumnFinishedAt = "`finished_at`"

// JenkinsBuildColumnSha is the Sha SQL column name for the JenkinsBuild table
const JenkinsBuildColumnSha = "sha"

// JenkinsBuildEscapedColumnSha is the escaped Sha SQL column name for the JenkinsBuild table
const JenkinsBuildEscapedColumnSha = "`sha`"

// JenkinsBuildColumnBranch is the Branch SQL column name for the JenkinsBuild table
const JenkinsBuildColumnBranch = "branch"

// JenkinsBuildEscapedColumnBranch is the escaped Branch SQL column name for the JenkinsBuild table
const JenkinsBuildEscapedColumnBranch = "`branch`"

// JenkinsBuildColumnURL is the URL SQL column name for the JenkinsBuild table
const JenkinsBuildColumnURL = "url"

// JenkinsBuildEscapedColumnURL is the escaped URL SQL column name for the JenkinsBuild table
const JenkinsBuildEscapedColumnURL = "`url`"

// GetID will return the JenkinsBuild ID value
func (t *JenkinsBuild) GetID() string {
	return t.ID
}

// SetID will set the JenkinsBuild ID value
func (t *JenkinsBuild) SetID(v string) {
	t.ID = v
}

// FindJenkinsBuildByID will find a JenkinsBuild by ID
func FindJenkinsBuildByID(ctx context.Context, db *sql.DB, value string) (*JenkinsBuild, error) {
	q := "SELECT `jenkins_build`.`id`,`jenkins_build`.`checksum`,`jenkins_build`.`customer_id`,`jenkins_build`.`jenkins_build_number`,`jenkins_build`.`job_id`,`jenkins_build`.`result`,`jenkins_build`.`created_at`,`jenkins_build`.`finished_at`,`jenkins_build`.`sha`,`jenkins_build`.`branch`,`jenkins_build`.`url` FROM `jenkins_build` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CustomerID sql.NullString
	var _JenkinsBuildNumber sql.NullInt64
	var _JobID sql.NullString
	var _Result sql.NullString
	var _CreatedAt sql.NullInt64
	var _FinishedAt sql.NullInt64
	var _Sha sql.NullString
	var _Branch sql.NullString
	var _URL sql.NullString
	err := db.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_JenkinsBuildNumber,
		&_JobID,
		&_Result,
		&_CreatedAt,
		&_FinishedAt,
		&_Sha,
		&_Branch,
		&_URL,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &JenkinsBuild{}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _Checksum.Valid {
		t.SetChecksum(_Checksum.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _JenkinsBuildNumber.Valid {
		t.SetJenkinsBuildNumber(_JenkinsBuildNumber.Int64)
	}
	if _JobID.Valid {
		t.SetJobID(_JobID.String)
	}
	if _Result.Valid {
		t.SetResultString(_Result.String)
	}
	if _CreatedAt.Valid {
		t.SetCreatedAt(_CreatedAt.Int64)
	}
	if _FinishedAt.Valid {
		t.SetFinishedAt(_FinishedAt.Int64)
	}
	if _Sha.Valid {
		t.SetSha(_Sha.String)
	}
	if _Branch.Valid {
		t.SetBranch(_Branch.String)
	}
	if _URL.Valid {
		t.SetURL(_URL.String)
	}
	return t, nil
}

// FindJenkinsBuildByIDTx will find a JenkinsBuild by ID using the provided transaction
func FindJenkinsBuildByIDTx(ctx context.Context, tx *sql.Tx, value string) (*JenkinsBuild, error) {
	q := "SELECT `jenkins_build`.`id`,`jenkins_build`.`checksum`,`jenkins_build`.`customer_id`,`jenkins_build`.`jenkins_build_number`,`jenkins_build`.`job_id`,`jenkins_build`.`result`,`jenkins_build`.`created_at`,`jenkins_build`.`finished_at`,`jenkins_build`.`sha`,`jenkins_build`.`branch`,`jenkins_build`.`url` FROM `jenkins_build` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CustomerID sql.NullString
	var _JenkinsBuildNumber sql.NullInt64
	var _JobID sql.NullString
	var _Result sql.NullString
	var _CreatedAt sql.NullInt64
	var _FinishedAt sql.NullInt64
	var _Sha sql.NullString
	var _Branch sql.NullString
	var _URL sql.NullString
	err := tx.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_JenkinsBuildNumber,
		&_JobID,
		&_Result,
		&_CreatedAt,
		&_FinishedAt,
		&_Sha,
		&_Branch,
		&_URL,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &JenkinsBuild{}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _Checksum.Valid {
		t.SetChecksum(_Checksum.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _JenkinsBuildNumber.Valid {
		t.SetJenkinsBuildNumber(_JenkinsBuildNumber.Int64)
	}
	if _JobID.Valid {
		t.SetJobID(_JobID.String)
	}
	if _Result.Valid {
		t.SetResultString(_Result.String)
	}
	if _CreatedAt.Valid {
		t.SetCreatedAt(_CreatedAt.Int64)
	}
	if _FinishedAt.Valid {
		t.SetFinishedAt(_FinishedAt.Int64)
	}
	if _Sha.Valid {
		t.SetSha(_Sha.String)
	}
	if _Branch.Valid {
		t.SetBranch(_Branch.String)
	}
	if _URL.Valid {
		t.SetURL(_URL.String)
	}
	return t, nil
}

// GetChecksum will return the JenkinsBuild Checksum value
func (t *JenkinsBuild) GetChecksum() string {
	if t.Checksum == nil {
		return ""
	}
	return *t.Checksum
}

// SetChecksum will set the JenkinsBuild Checksum value
func (t *JenkinsBuild) SetChecksum(v string) {
	t.Checksum = &v
}

// GetCustomerID will return the JenkinsBuild CustomerID value
func (t *JenkinsBuild) GetCustomerID() string {
	return t.CustomerID
}

// SetCustomerID will set the JenkinsBuild CustomerID value
func (t *JenkinsBuild) SetCustomerID(v string) {
	t.CustomerID = v
}

// FindJenkinsBuildsByCustomerID will find all JenkinsBuilds by the CustomerID value
func FindJenkinsBuildsByCustomerID(ctx context.Context, db *sql.DB, value string) ([]*JenkinsBuild, error) {
	q := "SELECT `jenkins_build`.`id`,`jenkins_build`.`checksum`,`jenkins_build`.`customer_id`,`jenkins_build`.`jenkins_build_number`,`jenkins_build`.`job_id`,`jenkins_build`.`result`,`jenkins_build`.`created_at`,`jenkins_build`.`finished_at`,`jenkins_build`.`sha`,`jenkins_build`.`branch`,`jenkins_build`.`url` FROM `jenkins_build` WHERE `customer_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JenkinsBuild, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _JenkinsBuildNumber sql.NullInt64
		var _JobID sql.NullString
		var _Result sql.NullString
		var _CreatedAt sql.NullInt64
		var _FinishedAt sql.NullInt64
		var _Sha sql.NullString
		var _Branch sql.NullString
		var _URL sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_JenkinsBuildNumber,
			&_JobID,
			&_Result,
			&_CreatedAt,
			&_FinishedAt,
			&_Sha,
			&_Branch,
			&_URL,
		)
		if err != nil {
			return nil, err
		}
		t := &JenkinsBuild{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _JenkinsBuildNumber.Valid {
			t.SetJenkinsBuildNumber(_JenkinsBuildNumber.Int64)
		}
		if _JobID.Valid {
			t.SetJobID(_JobID.String)
		}
		if _Result.Valid {
			t.SetResultString(_Result.String)
		}
		if _CreatedAt.Valid {
			t.SetCreatedAt(_CreatedAt.Int64)
		}
		if _FinishedAt.Valid {
			t.SetFinishedAt(_FinishedAt.Int64)
		}
		if _Sha.Valid {
			t.SetSha(_Sha.String)
		}
		if _Branch.Valid {
			t.SetBranch(_Branch.String)
		}
		if _URL.Valid {
			t.SetURL(_URL.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindJenkinsBuildsByCustomerIDTx will find all JenkinsBuilds by the CustomerID value using the provided transaction
func FindJenkinsBuildsByCustomerIDTx(ctx context.Context, tx *sql.Tx, value string) ([]*JenkinsBuild, error) {
	q := "SELECT `jenkins_build`.`id`,`jenkins_build`.`checksum`,`jenkins_build`.`customer_id`,`jenkins_build`.`jenkins_build_number`,`jenkins_build`.`job_id`,`jenkins_build`.`result`,`jenkins_build`.`created_at`,`jenkins_build`.`finished_at`,`jenkins_build`.`sha`,`jenkins_build`.`branch`,`jenkins_build`.`url` FROM `jenkins_build` WHERE `customer_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JenkinsBuild, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _JenkinsBuildNumber sql.NullInt64
		var _JobID sql.NullString
		var _Result sql.NullString
		var _CreatedAt sql.NullInt64
		var _FinishedAt sql.NullInt64
		var _Sha sql.NullString
		var _Branch sql.NullString
		var _URL sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_JenkinsBuildNumber,
			&_JobID,
			&_Result,
			&_CreatedAt,
			&_FinishedAt,
			&_Sha,
			&_Branch,
			&_URL,
		)
		if err != nil {
			return nil, err
		}
		t := &JenkinsBuild{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _JenkinsBuildNumber.Valid {
			t.SetJenkinsBuildNumber(_JenkinsBuildNumber.Int64)
		}
		if _JobID.Valid {
			t.SetJobID(_JobID.String)
		}
		if _Result.Valid {
			t.SetResultString(_Result.String)
		}
		if _CreatedAt.Valid {
			t.SetCreatedAt(_CreatedAt.Int64)
		}
		if _FinishedAt.Valid {
			t.SetFinishedAt(_FinishedAt.Int64)
		}
		if _Sha.Valid {
			t.SetSha(_Sha.String)
		}
		if _Branch.Valid {
			t.SetBranch(_Branch.String)
		}
		if _URL.Valid {
			t.SetURL(_URL.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetJenkinsBuildNumber will return the JenkinsBuild JenkinsBuildNumber value
func (t *JenkinsBuild) GetJenkinsBuildNumber() int64 {
	return t.JenkinsBuildNumber
}

// SetJenkinsBuildNumber will set the JenkinsBuild JenkinsBuildNumber value
func (t *JenkinsBuild) SetJenkinsBuildNumber(v int64) {
	t.JenkinsBuildNumber = v
}

// FindJenkinsBuildsByJenkinsBuildNumber will find all JenkinsBuilds by the JenkinsBuildNumber value
func FindJenkinsBuildsByJenkinsBuildNumber(ctx context.Context, db *sql.DB, value int64) ([]*JenkinsBuild, error) {
	q := "SELECT `jenkins_build`.`id`,`jenkins_build`.`checksum`,`jenkins_build`.`customer_id`,`jenkins_build`.`jenkins_build_number`,`jenkins_build`.`job_id`,`jenkins_build`.`result`,`jenkins_build`.`created_at`,`jenkins_build`.`finished_at`,`jenkins_build`.`sha`,`jenkins_build`.`branch`,`jenkins_build`.`url` FROM `jenkins_build` WHERE `jenkins_build_number` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JenkinsBuild, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _JenkinsBuildNumber sql.NullInt64
		var _JobID sql.NullString
		var _Result sql.NullString
		var _CreatedAt sql.NullInt64
		var _FinishedAt sql.NullInt64
		var _Sha sql.NullString
		var _Branch sql.NullString
		var _URL sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_JenkinsBuildNumber,
			&_JobID,
			&_Result,
			&_CreatedAt,
			&_FinishedAt,
			&_Sha,
			&_Branch,
			&_URL,
		)
		if err != nil {
			return nil, err
		}
		t := &JenkinsBuild{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _JenkinsBuildNumber.Valid {
			t.SetJenkinsBuildNumber(_JenkinsBuildNumber.Int64)
		}
		if _JobID.Valid {
			t.SetJobID(_JobID.String)
		}
		if _Result.Valid {
			t.SetResultString(_Result.String)
		}
		if _CreatedAt.Valid {
			t.SetCreatedAt(_CreatedAt.Int64)
		}
		if _FinishedAt.Valid {
			t.SetFinishedAt(_FinishedAt.Int64)
		}
		if _Sha.Valid {
			t.SetSha(_Sha.String)
		}
		if _Branch.Valid {
			t.SetBranch(_Branch.String)
		}
		if _URL.Valid {
			t.SetURL(_URL.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindJenkinsBuildsByJenkinsBuildNumberTx will find all JenkinsBuilds by the JenkinsBuildNumber value using the provided transaction
func FindJenkinsBuildsByJenkinsBuildNumberTx(ctx context.Context, tx *sql.Tx, value int64) ([]*JenkinsBuild, error) {
	q := "SELECT `jenkins_build`.`id`,`jenkins_build`.`checksum`,`jenkins_build`.`customer_id`,`jenkins_build`.`jenkins_build_number`,`jenkins_build`.`job_id`,`jenkins_build`.`result`,`jenkins_build`.`created_at`,`jenkins_build`.`finished_at`,`jenkins_build`.`sha`,`jenkins_build`.`branch`,`jenkins_build`.`url` FROM `jenkins_build` WHERE `jenkins_build_number` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JenkinsBuild, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _JenkinsBuildNumber sql.NullInt64
		var _JobID sql.NullString
		var _Result sql.NullString
		var _CreatedAt sql.NullInt64
		var _FinishedAt sql.NullInt64
		var _Sha sql.NullString
		var _Branch sql.NullString
		var _URL sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_JenkinsBuildNumber,
			&_JobID,
			&_Result,
			&_CreatedAt,
			&_FinishedAt,
			&_Sha,
			&_Branch,
			&_URL,
		)
		if err != nil {
			return nil, err
		}
		t := &JenkinsBuild{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _JenkinsBuildNumber.Valid {
			t.SetJenkinsBuildNumber(_JenkinsBuildNumber.Int64)
		}
		if _JobID.Valid {
			t.SetJobID(_JobID.String)
		}
		if _Result.Valid {
			t.SetResultString(_Result.String)
		}
		if _CreatedAt.Valid {
			t.SetCreatedAt(_CreatedAt.Int64)
		}
		if _FinishedAt.Valid {
			t.SetFinishedAt(_FinishedAt.Int64)
		}
		if _Sha.Valid {
			t.SetSha(_Sha.String)
		}
		if _Branch.Valid {
			t.SetBranch(_Branch.String)
		}
		if _URL.Valid {
			t.SetURL(_URL.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetJobID will return the JenkinsBuild JobID value
func (t *JenkinsBuild) GetJobID() string {
	return t.JobID
}

// SetJobID will set the JenkinsBuild JobID value
func (t *JenkinsBuild) SetJobID(v string) {
	t.JobID = v
}

// FindJenkinsBuildsByJobID will find all JenkinsBuilds by the JobID value
func FindJenkinsBuildsByJobID(ctx context.Context, db *sql.DB, value string) ([]*JenkinsBuild, error) {
	q := "SELECT `jenkins_build`.`id`,`jenkins_build`.`checksum`,`jenkins_build`.`customer_id`,`jenkins_build`.`jenkins_build_number`,`jenkins_build`.`job_id`,`jenkins_build`.`result`,`jenkins_build`.`created_at`,`jenkins_build`.`finished_at`,`jenkins_build`.`sha`,`jenkins_build`.`branch`,`jenkins_build`.`url` FROM `jenkins_build` WHERE `job_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JenkinsBuild, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _JenkinsBuildNumber sql.NullInt64
		var _JobID sql.NullString
		var _Result sql.NullString
		var _CreatedAt sql.NullInt64
		var _FinishedAt sql.NullInt64
		var _Sha sql.NullString
		var _Branch sql.NullString
		var _URL sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_JenkinsBuildNumber,
			&_JobID,
			&_Result,
			&_CreatedAt,
			&_FinishedAt,
			&_Sha,
			&_Branch,
			&_URL,
		)
		if err != nil {
			return nil, err
		}
		t := &JenkinsBuild{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _JenkinsBuildNumber.Valid {
			t.SetJenkinsBuildNumber(_JenkinsBuildNumber.Int64)
		}
		if _JobID.Valid {
			t.SetJobID(_JobID.String)
		}
		if _Result.Valid {
			t.SetResultString(_Result.String)
		}
		if _CreatedAt.Valid {
			t.SetCreatedAt(_CreatedAt.Int64)
		}
		if _FinishedAt.Valid {
			t.SetFinishedAt(_FinishedAt.Int64)
		}
		if _Sha.Valid {
			t.SetSha(_Sha.String)
		}
		if _Branch.Valid {
			t.SetBranch(_Branch.String)
		}
		if _URL.Valid {
			t.SetURL(_URL.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindJenkinsBuildsByJobIDTx will find all JenkinsBuilds by the JobID value using the provided transaction
func FindJenkinsBuildsByJobIDTx(ctx context.Context, tx *sql.Tx, value string) ([]*JenkinsBuild, error) {
	q := "SELECT `jenkins_build`.`id`,`jenkins_build`.`checksum`,`jenkins_build`.`customer_id`,`jenkins_build`.`jenkins_build_number`,`jenkins_build`.`job_id`,`jenkins_build`.`result`,`jenkins_build`.`created_at`,`jenkins_build`.`finished_at`,`jenkins_build`.`sha`,`jenkins_build`.`branch`,`jenkins_build`.`url` FROM `jenkins_build` WHERE `job_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JenkinsBuild, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _JenkinsBuildNumber sql.NullInt64
		var _JobID sql.NullString
		var _Result sql.NullString
		var _CreatedAt sql.NullInt64
		var _FinishedAt sql.NullInt64
		var _Sha sql.NullString
		var _Branch sql.NullString
		var _URL sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_JenkinsBuildNumber,
			&_JobID,
			&_Result,
			&_CreatedAt,
			&_FinishedAt,
			&_Sha,
			&_Branch,
			&_URL,
		)
		if err != nil {
			return nil, err
		}
		t := &JenkinsBuild{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _JenkinsBuildNumber.Valid {
			t.SetJenkinsBuildNumber(_JenkinsBuildNumber.Int64)
		}
		if _JobID.Valid {
			t.SetJobID(_JobID.String)
		}
		if _Result.Valid {
			t.SetResultString(_Result.String)
		}
		if _CreatedAt.Valid {
			t.SetCreatedAt(_CreatedAt.Int64)
		}
		if _FinishedAt.Valid {
			t.SetFinishedAt(_FinishedAt.Int64)
		}
		if _Sha.Valid {
			t.SetSha(_Sha.String)
		}
		if _Branch.Valid {
			t.SetBranch(_Branch.String)
		}
		if _URL.Valid {
			t.SetURL(_URL.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetResult will return the JenkinsBuild Result value
func (t *JenkinsBuild) GetResult() JenkinsBuild_JenkinsBuildResult {
	return t.Result
}

// SetResult will set the JenkinsBuild Result value
func (t *JenkinsBuild) SetResult(v JenkinsBuild_JenkinsBuildResult) {
	t.Result = v
}

// GetResultString will return the JenkinsBuild Result value as a string
func (t *JenkinsBuild) GetResultString() string {
	return t.Result.String()
}

// SetResultString will set the JenkinsBuild Result value from a string
func (t *JenkinsBuild) SetResultString(v string) {
	var _Result = toJenkinsBuildResult(v)
	if _Result != nil {
		t.Result = *_Result
	}
}

// GetCreatedAt will return the JenkinsBuild CreatedAt value
func (t *JenkinsBuild) GetCreatedAt() int64 {
	return t.CreatedAt
}

// SetCreatedAt will set the JenkinsBuild CreatedAt value
func (t *JenkinsBuild) SetCreatedAt(v int64) {
	t.CreatedAt = v
}

// GetFinishedAt will return the JenkinsBuild FinishedAt value
func (t *JenkinsBuild) GetFinishedAt() int64 {
	if t.FinishedAt == nil {
		return int64(0)
	}
	return *t.FinishedAt
}

// SetFinishedAt will set the JenkinsBuild FinishedAt value
func (t *JenkinsBuild) SetFinishedAt(v int64) {
	t.FinishedAt = &v
}

// GetSha will return the JenkinsBuild Sha value
func (t *JenkinsBuild) GetSha() string {
	if t.Sha == nil {
		return ""
	}
	return *t.Sha
}

// SetSha will set the JenkinsBuild Sha value
func (t *JenkinsBuild) SetSha(v string) {
	t.Sha = &v
}

// GetBranch will return the JenkinsBuild Branch value
func (t *JenkinsBuild) GetBranch() string {
	if t.Branch == nil {
		return ""
	}
	return *t.Branch
}

// SetBranch will set the JenkinsBuild Branch value
func (t *JenkinsBuild) SetBranch(v string) {
	t.Branch = &v
}

// GetURL will return the JenkinsBuild URL value
func (t *JenkinsBuild) GetURL() string {
	return t.URL
}

// SetURL will set the JenkinsBuild URL value
func (t *JenkinsBuild) SetURL(v string) {
	t.URL = v
}

func (t *JenkinsBuild) toTimestamp(value time.Time) *timestamp.Timestamp {
	ts, _ := ptypes.TimestampProto(value)
	return ts
}

// DBCreateJenkinsBuildTable will create the JenkinsBuild table
func DBCreateJenkinsBuildTable(ctx context.Context, db *sql.DB) error {
	q := "CREATE TABLE `jenkins_build` (`id` VARCHAR(64) NOT NULL PRIMARY KEY,`checksum` CHAR(64),`customer_id` VARCHAR(64) NOT NULL,`jenkins_build_number` BIGINT NOT NULL,`job_id`VARCHAR(64) NOT NULL,`result`ENUM('result_unknown','success','failure','aborted') NOT NULL,`created_at` BIGINT UNSIGNED NOT NULL,`finished_at` BIGINT UNSIGNED,`sha`VARCHAR(255),`branch`VARCHAR(255),`url`VARCHAR(255) NOT NULL,INDEX jenkins_build_customer_id_index (`customer_id`),INDEX jenkins_build_jenkins_build_number_index (`jenkins_build_number`),INDEX jenkins_build_job_id_index (`job_id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBCreateJenkinsBuildTableTx will create the JenkinsBuild table using the provided transction
func DBCreateJenkinsBuildTableTx(ctx context.Context, tx *sql.Tx) error {
	q := "CREATE TABLE `jenkins_build` (`id` VARCHAR(64) NOT NULL PRIMARY KEY,`checksum` CHAR(64),`customer_id` VARCHAR(64) NOT NULL,`jenkins_build_number` BIGINT NOT NULL,`job_id`VARCHAR(64) NOT NULL,`result`ENUM('result_unknown','success','failure','aborted') NOT NULL,`created_at` BIGINT UNSIGNED NOT NULL,`finished_at` BIGINT UNSIGNED,`sha`VARCHAR(255),`branch`VARCHAR(255),`url`VARCHAR(255) NOT NULL,INDEX jenkins_build_customer_id_index (`customer_id`),INDEX jenkins_build_jenkins_build_number_index (`jenkins_build_number`),INDEX jenkins_build_job_id_index (`job_id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// DBDropJenkinsBuildTable will drop the JenkinsBuild table
func DBDropJenkinsBuildTable(ctx context.Context, db *sql.DB) error {
	q := "DROP TABLE IF EXISTS `jenkins_build`"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBDropJenkinsBuildTableTx will drop the JenkinsBuild table using the provided transaction
func DBDropJenkinsBuildTableTx(ctx context.Context, tx *sql.Tx) error {
	q := "DROP TABLE IF EXISTS `jenkins_build`"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// CalculateChecksum will calculate a checksum of the SHA1 of all field values
func (t *JenkinsBuild) CalculateChecksum() string {
	return orm.HashStrings(
		orm.ToString(t.ID),
		orm.ToString(t.CustomerID),
		orm.ToString(t.JenkinsBuildNumber),
		orm.ToString(t.JobID),
		enumJenkinsBuildResultToString(&t.Result),
		orm.ToString(t.CreatedAt),
		orm.ToString(t.FinishedAt),
		orm.ToString(t.Sha),
		orm.ToString(t.Branch),
		orm.ToString(t.URL),
	)
}

// DBCreate will create a new JenkinsBuild record in the database
func (t *JenkinsBuild) DBCreate(ctx context.Context, db *sql.DB) (sql.Result, error) {
	q := "INSERT INTO `jenkins_build` (`jenkins_build`.`id`,`jenkins_build`.`checksum`,`jenkins_build`.`customer_id`,`jenkins_build`.`jenkins_build_number`,`jenkins_build`.`job_id`,`jenkins_build`.`result`,`jenkins_build`.`created_at`,`jenkins_build`.`finished_at`,`jenkins_build`.`sha`,`jenkins_build`.`branch`,`jenkins_build`.`url`) VALUES (?,?,?,?,?,?,?,?,?,?,?)"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLInt64(t.JenkinsBuildNumber),
		orm.ToSQLString(t.JobID),
		orm.ToSQLString(enumJenkinsBuildResultToString(&t.Result)),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.FinishedAt),
		orm.ToSQLString(t.Sha),
		orm.ToSQLString(t.Branch),
		orm.ToSQLString(t.URL),
	)
}

// DBCreateTx will create a new JenkinsBuild record in the database using the provided transaction
func (t *JenkinsBuild) DBCreateTx(ctx context.Context, tx *sql.Tx) (sql.Result, error) {
	q := "INSERT INTO `jenkins_build` (`jenkins_build`.`id`,`jenkins_build`.`checksum`,`jenkins_build`.`customer_id`,`jenkins_build`.`jenkins_build_number`,`jenkins_build`.`job_id`,`jenkins_build`.`result`,`jenkins_build`.`created_at`,`jenkins_build`.`finished_at`,`jenkins_build`.`sha`,`jenkins_build`.`branch`,`jenkins_build`.`url`) VALUES (?,?,?,?,?,?,?,?,?,?,?)"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLInt64(t.JenkinsBuildNumber),
		orm.ToSQLString(t.JobID),
		orm.ToSQLString(enumJenkinsBuildResultToString(&t.Result)),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.FinishedAt),
		orm.ToSQLString(t.Sha),
		orm.ToSQLString(t.Branch),
		orm.ToSQLString(t.URL),
	)
}

// DBCreateIgnoreDuplicate will upsert the JenkinsBuild record in the database
func (t *JenkinsBuild) DBCreateIgnoreDuplicate(ctx context.Context, db *sql.DB) (sql.Result, error) {
	q := "INSERT INTO `jenkins_build` (`jenkins_build`.`id`,`jenkins_build`.`checksum`,`jenkins_build`.`customer_id`,`jenkins_build`.`jenkins_build_number`,`jenkins_build`.`job_id`,`jenkins_build`.`result`,`jenkins_build`.`created_at`,`jenkins_build`.`finished_at`,`jenkins_build`.`sha`,`jenkins_build`.`branch`,`jenkins_build`.`url`) VALUES (?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLInt64(t.JenkinsBuildNumber),
		orm.ToSQLString(t.JobID),
		orm.ToSQLString(enumJenkinsBuildResultToString(&t.Result)),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.FinishedAt),
		orm.ToSQLString(t.Sha),
		orm.ToSQLString(t.Branch),
		orm.ToSQLString(t.URL),
	)
}

// DBCreateIgnoreDuplicateTx will upsert the JenkinsBuild record in the database using the provided transaction
func (t *JenkinsBuild) DBCreateIgnoreDuplicateTx(ctx context.Context, tx *sql.Tx) (sql.Result, error) {
	q := "INSERT INTO `jenkins_build` (`jenkins_build`.`id`,`jenkins_build`.`checksum`,`jenkins_build`.`customer_id`,`jenkins_build`.`jenkins_build_number`,`jenkins_build`.`job_id`,`jenkins_build`.`result`,`jenkins_build`.`created_at`,`jenkins_build`.`finished_at`,`jenkins_build`.`sha`,`jenkins_build`.`branch`,`jenkins_build`.`url`) VALUES (?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLInt64(t.JenkinsBuildNumber),
		orm.ToSQLString(t.JobID),
		orm.ToSQLString(enumJenkinsBuildResultToString(&t.Result)),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.FinishedAt),
		orm.ToSQLString(t.Sha),
		orm.ToSQLString(t.Branch),
		orm.ToSQLString(t.URL),
	)
}

// DeleteAllJenkinsBuilds deletes all JenkinsBuild records in the database with optional filters
func DeleteAllJenkinsBuilds(ctx context.Context, db *sql.DB, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(JenkinsBuildTableName),
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

// DeleteAllJenkinsBuildsTx deletes all JenkinsBuild records in the database with optional filters using the provided transaction
func DeleteAllJenkinsBuildsTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(JenkinsBuildTableName),
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

// DBDelete will delete this JenkinsBuild record in the database
func (t *JenkinsBuild) DBDelete(ctx context.Context, db *sql.DB) (bool, error) {
	q := "DELETE FROM `jenkins_build` WHERE `id` = ?"
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

// DBDeleteTx will delete this JenkinsBuild record in the database using the provided transaction
func (t *JenkinsBuild) DBDeleteTx(ctx context.Context, tx *sql.Tx) (bool, error) {
	q := "DELETE FROM `jenkins_build` WHERE `id` = ?"
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

// DBUpdate will update the JenkinsBuild record in the database
func (t *JenkinsBuild) DBUpdate(ctx context.Context, db *sql.DB) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `jenkins_build` SET `checksum`=?,`customer_id`=?,`jenkins_build_number`=?,`job_id`=?,`result`=?,`created_at`=?,`finished_at`=?,`sha`=?,`branch`=?,`url`=? WHERE `id`=?"
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLInt64(t.JenkinsBuildNumber),
		orm.ToSQLString(t.JobID),
		orm.ToSQLString(enumJenkinsBuildResultToString(&t.Result)),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.FinishedAt),
		orm.ToSQLString(t.Sha),
		orm.ToSQLString(t.Branch),
		orm.ToSQLString(t.URL),
		orm.ToSQLString(t.ID),
	)
}

// DBUpdateTx will update the JenkinsBuild record in the database using the provided transaction
func (t *JenkinsBuild) DBUpdateTx(ctx context.Context, tx *sql.Tx) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `jenkins_build` SET `checksum`=?,`customer_id`=?,`jenkins_build_number`=?,`job_id`=?,`result`=?,`created_at`=?,`finished_at`=?,`sha`=?,`branch`=?,`url`=? WHERE `id`=?"
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLInt64(t.JenkinsBuildNumber),
		orm.ToSQLString(t.JobID),
		orm.ToSQLString(enumJenkinsBuildResultToString(&t.Result)),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.FinishedAt),
		orm.ToSQLString(t.Sha),
		orm.ToSQLString(t.Branch),
		orm.ToSQLString(t.URL),
		orm.ToSQLString(t.ID),
	)
}

// DBUpsert will upsert the JenkinsBuild record in the database
func (t *JenkinsBuild) DBUpsert(ctx context.Context, db *sql.DB, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `jenkins_build` (`jenkins_build`.`id`,`jenkins_build`.`checksum`,`jenkins_build`.`customer_id`,`jenkins_build`.`jenkins_build_number`,`jenkins_build`.`job_id`,`jenkins_build`.`result`,`jenkins_build`.`created_at`,`jenkins_build`.`finished_at`,`jenkins_build`.`sha`,`jenkins_build`.`branch`,`jenkins_build`.`url`) VALUES (?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `jenkins_build` (`jenkins_build`.`id`,`jenkins_build`.`checksum`,`jenkins_build`.`customer_id`,`jenkins_build`.`jenkins_build_number`,`jenkins_build`.`job_id`,`jenkins_build`.`result`,`jenkins_build`.`created_at`,`jenkins_build`.`finished_at`,`jenkins_build`.`sha`,`jenkins_build`.`branch`,`jenkins_build`.`url`) VALUES (?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`customer_id`=VALUES(`customer_id`),`jenkins_build_number`=VALUES(`jenkins_build_number`),`job_id`=VALUES(`job_id`),`result`=VALUES(`result`),`created_at`=VALUES(`created_at`),`finished_at`=VALUES(`finished_at`),`sha`=VALUES(`sha`),`branch`=VALUES(`branch`),`url`=VALUES(`url`)"
	}
	r, err := db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLInt64(t.JenkinsBuildNumber),
		orm.ToSQLString(t.JobID),
		orm.ToSQLString(enumJenkinsBuildResultToString(&t.Result)),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.FinishedAt),
		orm.ToSQLString(t.Sha),
		orm.ToSQLString(t.Branch),
		orm.ToSQLString(t.URL),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBUpsertTx will upsert the JenkinsBuild record in the database using the provided transaction
func (t *JenkinsBuild) DBUpsertTx(ctx context.Context, tx *sql.Tx, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `jenkins_build` (`jenkins_build`.`id`,`jenkins_build`.`checksum`,`jenkins_build`.`customer_id`,`jenkins_build`.`jenkins_build_number`,`jenkins_build`.`job_id`,`jenkins_build`.`result`,`jenkins_build`.`created_at`,`jenkins_build`.`finished_at`,`jenkins_build`.`sha`,`jenkins_build`.`branch`,`jenkins_build`.`url`) VALUES (?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `jenkins_build` (`jenkins_build`.`id`,`jenkins_build`.`checksum`,`jenkins_build`.`customer_id`,`jenkins_build`.`jenkins_build_number`,`jenkins_build`.`job_id`,`jenkins_build`.`result`,`jenkins_build`.`created_at`,`jenkins_build`.`finished_at`,`jenkins_build`.`sha`,`jenkins_build`.`branch`,`jenkins_build`.`url`) VALUES (?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`customer_id`=VALUES(`customer_id`),`jenkins_build_number`=VALUES(`jenkins_build_number`),`job_id`=VALUES(`job_id`),`result`=VALUES(`result`),`created_at`=VALUES(`created_at`),`finished_at`=VALUES(`finished_at`),`sha`=VALUES(`sha`),`branch`=VALUES(`branch`),`url`=VALUES(`url`)"
	}
	r, err := tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLInt64(t.JenkinsBuildNumber),
		orm.ToSQLString(t.JobID),
		orm.ToSQLString(enumJenkinsBuildResultToString(&t.Result)),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.FinishedAt),
		orm.ToSQLString(t.Sha),
		orm.ToSQLString(t.Branch),
		orm.ToSQLString(t.URL),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBFindOne will find a JenkinsBuild record in the database with the primary key
func (t *JenkinsBuild) DBFindOne(ctx context.Context, db *sql.DB, value string) (bool, error) {
	q := "SELECT `jenkins_build`.`id`,`jenkins_build`.`checksum`,`jenkins_build`.`customer_id`,`jenkins_build`.`jenkins_build_number`,`jenkins_build`.`job_id`,`jenkins_build`.`result`,`jenkins_build`.`created_at`,`jenkins_build`.`finished_at`,`jenkins_build`.`sha`,`jenkins_build`.`branch`,`jenkins_build`.`url` FROM `jenkins_build` WHERE `id` = ? LIMIT 1"
	row := db.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CustomerID sql.NullString
	var _JenkinsBuildNumber sql.NullInt64
	var _JobID sql.NullString
	var _Result sql.NullString
	var _CreatedAt sql.NullInt64
	var _FinishedAt sql.NullInt64
	var _Sha sql.NullString
	var _Branch sql.NullString
	var _URL sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_JenkinsBuildNumber,
		&_JobID,
		&_Result,
		&_CreatedAt,
		&_FinishedAt,
		&_Sha,
		&_Branch,
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
	if _JenkinsBuildNumber.Valid {
		t.SetJenkinsBuildNumber(_JenkinsBuildNumber.Int64)
	}
	if _JobID.Valid {
		t.SetJobID(_JobID.String)
	}
	if _Result.Valid {
		t.SetResultString(_Result.String)
	}
	if _CreatedAt.Valid {
		t.SetCreatedAt(_CreatedAt.Int64)
	}
	if _FinishedAt.Valid {
		t.SetFinishedAt(_FinishedAt.Int64)
	}
	if _Sha.Valid {
		t.SetSha(_Sha.String)
	}
	if _Branch.Valid {
		t.SetBranch(_Branch.String)
	}
	if _URL.Valid {
		t.SetURL(_URL.String)
	}
	return true, nil
}

// DBFindOneTx will find a JenkinsBuild record in the database with the primary key using the provided transaction
func (t *JenkinsBuild) DBFindOneTx(ctx context.Context, tx *sql.Tx, value string) (bool, error) {
	q := "SELECT `jenkins_build`.`id`,`jenkins_build`.`checksum`,`jenkins_build`.`customer_id`,`jenkins_build`.`jenkins_build_number`,`jenkins_build`.`job_id`,`jenkins_build`.`result`,`jenkins_build`.`created_at`,`jenkins_build`.`finished_at`,`jenkins_build`.`sha`,`jenkins_build`.`branch`,`jenkins_build`.`url` FROM `jenkins_build` WHERE `id` = ? LIMIT 1"
	row := tx.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CustomerID sql.NullString
	var _JenkinsBuildNumber sql.NullInt64
	var _JobID sql.NullString
	var _Result sql.NullString
	var _CreatedAt sql.NullInt64
	var _FinishedAt sql.NullInt64
	var _Sha sql.NullString
	var _Branch sql.NullString
	var _URL sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_JenkinsBuildNumber,
		&_JobID,
		&_Result,
		&_CreatedAt,
		&_FinishedAt,
		&_Sha,
		&_Branch,
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
	if _JenkinsBuildNumber.Valid {
		t.SetJenkinsBuildNumber(_JenkinsBuildNumber.Int64)
	}
	if _JobID.Valid {
		t.SetJobID(_JobID.String)
	}
	if _Result.Valid {
		t.SetResultString(_Result.String)
	}
	if _CreatedAt.Valid {
		t.SetCreatedAt(_CreatedAt.Int64)
	}
	if _FinishedAt.Valid {
		t.SetFinishedAt(_FinishedAt.Int64)
	}
	if _Sha.Valid {
		t.SetSha(_Sha.String)
	}
	if _Branch.Valid {
		t.SetBranch(_Branch.String)
	}
	if _URL.Valid {
		t.SetURL(_URL.String)
	}
	return true, nil
}

// FindJenkinsBuilds will find a JenkinsBuild record in the database with the provided parameters
func FindJenkinsBuilds(ctx context.Context, db *sql.DB, _params ...interface{}) ([]*JenkinsBuild, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("customer_id"),
		orm.Column("jenkins_build_number"),
		orm.Column("job_id"),
		orm.Column("result"),
		orm.Column("created_at"),
		orm.Column("finished_at"),
		orm.Column("sha"),
		orm.Column("branch"),
		orm.Column("url"),
		orm.Table(JenkinsBuildTableName),
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
	results := make([]*JenkinsBuild, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _JenkinsBuildNumber sql.NullInt64
		var _JobID sql.NullString
		var _Result sql.NullString
		var _CreatedAt sql.NullInt64
		var _FinishedAt sql.NullInt64
		var _Sha sql.NullString
		var _Branch sql.NullString
		var _URL sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_JenkinsBuildNumber,
			&_JobID,
			&_Result,
			&_CreatedAt,
			&_FinishedAt,
			&_Sha,
			&_Branch,
			&_URL,
		)
		if err != nil {
			return nil, err
		}
		t := &JenkinsBuild{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _JenkinsBuildNumber.Valid {
			t.SetJenkinsBuildNumber(_JenkinsBuildNumber.Int64)
		}
		if _JobID.Valid {
			t.SetJobID(_JobID.String)
		}
		if _Result.Valid {
			t.SetResultString(_Result.String)
		}
		if _CreatedAt.Valid {
			t.SetCreatedAt(_CreatedAt.Int64)
		}
		if _FinishedAt.Valid {
			t.SetFinishedAt(_FinishedAt.Int64)
		}
		if _Sha.Valid {
			t.SetSha(_Sha.String)
		}
		if _Branch.Valid {
			t.SetBranch(_Branch.String)
		}
		if _URL.Valid {
			t.SetURL(_URL.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindJenkinsBuildsTx will find a JenkinsBuild record in the database with the provided parameters using the provided transaction
func FindJenkinsBuildsTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) ([]*JenkinsBuild, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("customer_id"),
		orm.Column("jenkins_build_number"),
		orm.Column("job_id"),
		orm.Column("result"),
		orm.Column("created_at"),
		orm.Column("finished_at"),
		orm.Column("sha"),
		orm.Column("branch"),
		orm.Column("url"),
		orm.Table(JenkinsBuildTableName),
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
	results := make([]*JenkinsBuild, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _JenkinsBuildNumber sql.NullInt64
		var _JobID sql.NullString
		var _Result sql.NullString
		var _CreatedAt sql.NullInt64
		var _FinishedAt sql.NullInt64
		var _Sha sql.NullString
		var _Branch sql.NullString
		var _URL sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_JenkinsBuildNumber,
			&_JobID,
			&_Result,
			&_CreatedAt,
			&_FinishedAt,
			&_Sha,
			&_Branch,
			&_URL,
		)
		if err != nil {
			return nil, err
		}
		t := &JenkinsBuild{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _JenkinsBuildNumber.Valid {
			t.SetJenkinsBuildNumber(_JenkinsBuildNumber.Int64)
		}
		if _JobID.Valid {
			t.SetJobID(_JobID.String)
		}
		if _Result.Valid {
			t.SetResultString(_Result.String)
		}
		if _CreatedAt.Valid {
			t.SetCreatedAt(_CreatedAt.Int64)
		}
		if _FinishedAt.Valid {
			t.SetFinishedAt(_FinishedAt.Int64)
		}
		if _Sha.Valid {
			t.SetSha(_Sha.String)
		}
		if _Branch.Valid {
			t.SetBranch(_Branch.String)
		}
		if _URL.Valid {
			t.SetURL(_URL.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// DBFind will find a JenkinsBuild record in the database with the provided parameters
func (t *JenkinsBuild) DBFind(ctx context.Context, db *sql.DB, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("customer_id"),
		orm.Column("jenkins_build_number"),
		orm.Column("job_id"),
		orm.Column("result"),
		orm.Column("created_at"),
		orm.Column("finished_at"),
		orm.Column("sha"),
		orm.Column("branch"),
		orm.Column("url"),
		orm.Table(JenkinsBuildTableName),
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
	var _JenkinsBuildNumber sql.NullInt64
	var _JobID sql.NullString
	var _Result sql.NullString
	var _CreatedAt sql.NullInt64
	var _FinishedAt sql.NullInt64
	var _Sha sql.NullString
	var _Branch sql.NullString
	var _URL sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_JenkinsBuildNumber,
		&_JobID,
		&_Result,
		&_CreatedAt,
		&_FinishedAt,
		&_Sha,
		&_Branch,
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
	if _JenkinsBuildNumber.Valid {
		t.SetJenkinsBuildNumber(_JenkinsBuildNumber.Int64)
	}
	if _JobID.Valid {
		t.SetJobID(_JobID.String)
	}
	if _Result.Valid {
		t.SetResultString(_Result.String)
	}
	if _CreatedAt.Valid {
		t.SetCreatedAt(_CreatedAt.Int64)
	}
	if _FinishedAt.Valid {
		t.SetFinishedAt(_FinishedAt.Int64)
	}
	if _Sha.Valid {
		t.SetSha(_Sha.String)
	}
	if _Branch.Valid {
		t.SetBranch(_Branch.String)
	}
	if _URL.Valid {
		t.SetURL(_URL.String)
	}
	return true, nil
}

// DBFindTx will find a JenkinsBuild record in the database with the provided parameters using the provided transaction
func (t *JenkinsBuild) DBFindTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("customer_id"),
		orm.Column("jenkins_build_number"),
		orm.Column("job_id"),
		orm.Column("result"),
		orm.Column("created_at"),
		orm.Column("finished_at"),
		orm.Column("sha"),
		orm.Column("branch"),
		orm.Column("url"),
		orm.Table(JenkinsBuildTableName),
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
	var _JenkinsBuildNumber sql.NullInt64
	var _JobID sql.NullString
	var _Result sql.NullString
	var _CreatedAt sql.NullInt64
	var _FinishedAt sql.NullInt64
	var _Sha sql.NullString
	var _Branch sql.NullString
	var _URL sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_JenkinsBuildNumber,
		&_JobID,
		&_Result,
		&_CreatedAt,
		&_FinishedAt,
		&_Sha,
		&_Branch,
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
	if _JenkinsBuildNumber.Valid {
		t.SetJenkinsBuildNumber(_JenkinsBuildNumber.Int64)
	}
	if _JobID.Valid {
		t.SetJobID(_JobID.String)
	}
	if _Result.Valid {
		t.SetResultString(_Result.String)
	}
	if _CreatedAt.Valid {
		t.SetCreatedAt(_CreatedAt.Int64)
	}
	if _FinishedAt.Valid {
		t.SetFinishedAt(_FinishedAt.Int64)
	}
	if _Sha.Valid {
		t.SetSha(_Sha.String)
	}
	if _Branch.Valid {
		t.SetBranch(_Branch.String)
	}
	if _URL.Valid {
		t.SetURL(_URL.String)
	}
	return true, nil
}

// CountJenkinsBuilds will find the count of JenkinsBuild records in the database
func CountJenkinsBuilds(ctx context.Context, db *sql.DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(JenkinsBuildTableName),
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

// CountJenkinsBuildsTx will find the count of JenkinsBuild records in the database using the provided transaction
func CountJenkinsBuildsTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(JenkinsBuildTableName),
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

// DBCount will find the count of JenkinsBuild records in the database
func (t *JenkinsBuild) DBCount(ctx context.Context, db *sql.DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(JenkinsBuildTableName),
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

// DBCountTx will find the count of JenkinsBuild records in the database using the provided transaction
func (t *JenkinsBuild) DBCountTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(JenkinsBuildTableName),
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

// DBExists will return true if the JenkinsBuild record exists in the database
func (t *JenkinsBuild) DBExists(ctx context.Context, db *sql.DB) (bool, error) {
	q := "SELECT `id` FROM `jenkins_build` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := db.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// DBExistsTx will return true if the JenkinsBuild record exists in the database using the provided transaction
func (t *JenkinsBuild) DBExistsTx(ctx context.Context, tx *sql.Tx) (bool, error) {
	q := "SELECT `id` FROM `jenkins_build` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := tx.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// PrimaryKeyColumn returns the column name for the primary key
func (t *JenkinsBuild) PrimaryKeyColumn() string {
	return JenkinsBuildColumnID
}

// PrimaryKeyColumnType returns the primary key column Go type as a string
func (t *JenkinsBuild) PrimaryKeyColumnType() string {
	return "string"
}

// PrimaryKey returns the primary key column value
func (t *JenkinsBuild) PrimaryKey() interface{} {
	return t.ID
}
