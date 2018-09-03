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

var _ Model = (*CommitSummary)(nil)
var _ CSVWriter = (*CommitSummary)(nil)
var _ JSONWriter = (*CommitSummary)(nil)

// CommitSummaryTableName is the name of the table in SQL
const CommitSummaryTableName = "commit_summary"

var CommitSummaryColumns = []string{
	"id",
	"commit_id",
	"repo_id",
	"author_user_id",
	"customer_id",
	"ref_type",
	"additions",
	"deletions",
	"files_changed",
	"branch",
	"language",
	"date",
	"message",
}

// CommitSummary table
type CommitSummary struct {
	Additions    int32   `json:"additions"`
	AuthorUserID *string `json:"author_user_id,omitempty"`
	Branch       *string `json:"branch,omitempty"`
	CommitID     string  `json:"commit_id"`
	CustomerID   string  `json:"customer_id"`
	Date         int64   `json:"date"`
	Deletions    int32   `json:"deletions"`
	FilesChanged int32   `json:"files_changed"`
	ID           string  `json:"id"`
	Language     string  `json:"language"`
	Message      *string `json:"message,omitempty"`
	RefType      string  `json:"ref_type"`
	RepoID       string  `json:"repo_id"`
}

// TableName returns the SQL table name for CommitSummary and satifies the Model interface
func (t *CommitSummary) TableName() string {
	return CommitSummaryTableName
}

// ToCSV will serialize the CommitSummary instance to a CSV compatible array of strings
func (t *CommitSummary) ToCSV() []string {
	return []string{
		t.ID,
		t.CommitID,
		t.RepoID,
		toCSVString(t.AuthorUserID),
		t.CustomerID,
		t.RefType,
		toCSVString(t.Additions),
		toCSVString(t.Deletions),
		toCSVString(t.FilesChanged),
		toCSVString(t.Branch),
		t.Language,
		toCSVString(t.Date),
		toCSVString(t.Message),
	}
}

// WriteCSV will serialize the CommitSummary instance to the writer as CSV and satisfies the CSVWriter interface
func (t *CommitSummary) WriteCSV(w *csv.Writer) error {
	return w.Write(t.ToCSV())
}

// WriteJSON will serialize the CommitSummary instance to the writer as JSON and satisfies the JSONWriter interface
func (t *CommitSummary) WriteJSON(w io.Writer, indent ...bool) error {
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

// NewCommitSummaryReader creates a JSON reader which can read in CommitSummary objects serialized as JSON either as an array, single object or json new lines
// and writes each CommitSummary to the channel provided
func NewCommitSummaryReader(r io.Reader, ch chan<- CommitSummary) error {
	return orm.Deserialize(r, func(buf json.RawMessage) error {
		dec := json.NewDecoder(bytes.NewBuffer(buf))
		e := CommitSummary{}
		if err := dec.Decode(&e); err != nil {
			return err
		}
		ch <- e
		return nil
	})
}

// NewCSVCommitSummaryReaderDir will read the reader as CSV and emit each record to the channel provided
func NewCSVCommitSummaryReader(r io.Reader, ch chan<- CommitSummary) error {
	cr := csv.NewReader(r)
	for {
		record, err := cr.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		ch <- CommitSummary{
			ID:           record[0],
			CommitID:     record[1],
			RepoID:       record[2],
			AuthorUserID: fromStringPointer(record[3]),
			CustomerID:   record[4],
			RefType:      record[5],
			Additions:    fromCSVInt32(record[6]),
			Deletions:    fromCSVInt32(record[7]),
			FilesChanged: fromCSVInt32(record[8]),
			Branch:       fromStringPointer(record[9]),
			Language:     record[10],
			Date:         fromCSVInt64(record[11]),
			Message:      fromStringPointer(record[12]),
		}
	}
	return nil
}

// NewCSVCommitSummaryReaderFile will read the file as a CSV and emit each record to the channel provided
func NewCSVCommitSummaryReaderFile(fp string, ch chan<- CommitSummary) error {
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
	return NewCSVCommitSummaryReader(fc, ch)
}

// NewCSVCommitSummaryReaderDir will read the commit_summary.csv.gz file as a CSV and emit each record to the channel provided
func NewCSVCommitSummaryReaderDir(dir string, ch chan<- CommitSummary) error {
	return NewCSVCommitSummaryReaderFile(filepath.Join(dir, "commit_summary.csv.gz"), ch)
}

// CommitSummaryCSVDeduper is a function callback which takes the existing value (a) and the new value (b)
// and the return value should be the one to use (or a new one, if applicable). return nil
// to skip processing of this record
type CommitSummaryCSVDeduper func(a CommitSummary, b CommitSummary) *CommitSummary

// CommitSummaryCSVDedupeDisabled is set on whether the CSV writer should de-dupe values by key
var CommitSummaryCSVDedupeDisabled bool

// NewCommitSummaryCSVWriterSize creates a batch writer that will write each CommitSummary into a CSV file
func NewCommitSummaryCSVWriterSize(w io.Writer, size int, dedupers ...CommitSummaryCSVDeduper) (chan CommitSummary, chan bool, error) {
	cw := csv.NewWriter(w)
	ch := make(chan CommitSummary, size)
	done := make(chan bool)
	go func() {
		defer func() { done <- true }()
		dodedupe := !CommitSummaryCSVDedupeDisabled
		var kv map[string]*CommitSummary
		var deduper CommitSummaryCSVDeduper
		if dedupers != nil && len(dedupers) > 0 {
			deduper = dedupers[0]
			dodedupe = true
		}
		if dodedupe {
			kv = make(map[string]*CommitSummary)
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

// CommitSummaryCSVDefaultSize is the default channel buffer size if not provided
var CommitSummaryCSVDefaultSize = 100

// NewCommitSummaryCSVWriter creates a batch writer that will write each CommitSummary into a CSV file
func NewCommitSummaryCSVWriter(w io.Writer, dedupers ...CommitSummaryCSVDeduper) (chan CommitSummary, chan bool, error) {
	return NewCommitSummaryCSVWriterSize(w, CommitSummaryCSVDefaultSize, dedupers...)
}

// NewCommitSummaryCSVWriterDir creates a batch writer that will write each CommitSummary into a CSV file named commit_summary.csv.gz in dir
func NewCommitSummaryCSVWriterDir(dir string, dedupers ...CommitSummaryCSVDeduper) (chan CommitSummary, chan bool, error) {
	return NewCommitSummaryCSVWriterFile(filepath.Join(dir, "commit_summary.csv.gz"), dedupers...)
}

// NewCommitSummaryCSVWriterFile creates a batch writer that will write each CommitSummary into a CSV file
func NewCommitSummaryCSVWriterFile(fn string, dedupers ...CommitSummaryCSVDeduper) (chan CommitSummary, chan bool, error) {
	f, err := os.Create(fn)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening CSV file %s. %v", fn, err)
	}
	var fc io.WriteCloser = f
	if filepath.Ext(fn) == ".gz" {
		w, _ := gzip.NewWriterLevel(f, gzip.BestCompression)
		fc = w
	}
	ch, done, err := NewCommitSummaryCSVWriter(fc, dedupers...)
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

type CommitSummaryDBAction func(ctx context.Context, db *sql.DB, record CommitSummary) error

// NewCommitSummaryDBWriterSize creates a DB writer that will write each issue into the DB
func NewCommitSummaryDBWriterSize(ctx context.Context, db *sql.DB, errors chan<- error, size int, actions ...CommitSummaryDBAction) (chan CommitSummary, chan bool, error) {
	ch := make(chan CommitSummary, size)
	done := make(chan bool)
	var action CommitSummaryDBAction
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

// NewCommitSummaryDBWriter creates a DB writer that will write each issue into the DB
func NewCommitSummaryDBWriter(ctx context.Context, db *sql.DB, errors chan<- error, actions ...CommitSummaryDBAction) (chan CommitSummary, chan bool, error) {
	return NewCommitSummaryDBWriterSize(ctx, db, errors, 100, actions...)
}

// CommitSummaryColumnID is the ID SQL column name for the CommitSummary table
const CommitSummaryColumnID = "id"

// CommitSummaryEscapedColumnID is the escaped ID SQL column name for the CommitSummary table
const CommitSummaryEscapedColumnID = "`id`"

// CommitSummaryColumnCommitID is the CommitID SQL column name for the CommitSummary table
const CommitSummaryColumnCommitID = "commit_id"

// CommitSummaryEscapedColumnCommitID is the escaped CommitID SQL column name for the CommitSummary table
const CommitSummaryEscapedColumnCommitID = "`commit_id`"

// CommitSummaryColumnRepoID is the RepoID SQL column name for the CommitSummary table
const CommitSummaryColumnRepoID = "repo_id"

// CommitSummaryEscapedColumnRepoID is the escaped RepoID SQL column name for the CommitSummary table
const CommitSummaryEscapedColumnRepoID = "`repo_id`"

// CommitSummaryColumnAuthorUserID is the AuthorUserID SQL column name for the CommitSummary table
const CommitSummaryColumnAuthorUserID = "author_user_id"

// CommitSummaryEscapedColumnAuthorUserID is the escaped AuthorUserID SQL column name for the CommitSummary table
const CommitSummaryEscapedColumnAuthorUserID = "`author_user_id`"

// CommitSummaryColumnCustomerID is the CustomerID SQL column name for the CommitSummary table
const CommitSummaryColumnCustomerID = "customer_id"

// CommitSummaryEscapedColumnCustomerID is the escaped CustomerID SQL column name for the CommitSummary table
const CommitSummaryEscapedColumnCustomerID = "`customer_id`"

// CommitSummaryColumnRefType is the RefType SQL column name for the CommitSummary table
const CommitSummaryColumnRefType = "ref_type"

// CommitSummaryEscapedColumnRefType is the escaped RefType SQL column name for the CommitSummary table
const CommitSummaryEscapedColumnRefType = "`ref_type`"

// CommitSummaryColumnAdditions is the Additions SQL column name for the CommitSummary table
const CommitSummaryColumnAdditions = "additions"

// CommitSummaryEscapedColumnAdditions is the escaped Additions SQL column name for the CommitSummary table
const CommitSummaryEscapedColumnAdditions = "`additions`"

// CommitSummaryColumnDeletions is the Deletions SQL column name for the CommitSummary table
const CommitSummaryColumnDeletions = "deletions"

// CommitSummaryEscapedColumnDeletions is the escaped Deletions SQL column name for the CommitSummary table
const CommitSummaryEscapedColumnDeletions = "`deletions`"

// CommitSummaryColumnFilesChanged is the FilesChanged SQL column name for the CommitSummary table
const CommitSummaryColumnFilesChanged = "files_changed"

// CommitSummaryEscapedColumnFilesChanged is the escaped FilesChanged SQL column name for the CommitSummary table
const CommitSummaryEscapedColumnFilesChanged = "`files_changed`"

// CommitSummaryColumnBranch is the Branch SQL column name for the CommitSummary table
const CommitSummaryColumnBranch = "branch"

// CommitSummaryEscapedColumnBranch is the escaped Branch SQL column name for the CommitSummary table
const CommitSummaryEscapedColumnBranch = "`branch`"

// CommitSummaryColumnLanguage is the Language SQL column name for the CommitSummary table
const CommitSummaryColumnLanguage = "language"

// CommitSummaryEscapedColumnLanguage is the escaped Language SQL column name for the CommitSummary table
const CommitSummaryEscapedColumnLanguage = "`language`"

// CommitSummaryColumnDate is the Date SQL column name for the CommitSummary table
const CommitSummaryColumnDate = "date"

// CommitSummaryEscapedColumnDate is the escaped Date SQL column name for the CommitSummary table
const CommitSummaryEscapedColumnDate = "`date`"

// CommitSummaryColumnMessage is the Message SQL column name for the CommitSummary table
const CommitSummaryColumnMessage = "message"

// CommitSummaryEscapedColumnMessage is the escaped Message SQL column name for the CommitSummary table
const CommitSummaryEscapedColumnMessage = "`message`"

// GetID will return the CommitSummary ID value
func (t *CommitSummary) GetID() string {
	return t.ID
}

// SetID will set the CommitSummary ID value
func (t *CommitSummary) SetID(v string) {
	t.ID = v
}

// FindCommitSummaryByID will find a CommitSummary by ID
func FindCommitSummaryByID(ctx context.Context, db *sql.DB, value string) (*CommitSummary, error) {
	q := "SELECT `commit_summary`.`id`,`commit_summary`.`commit_id`,`commit_summary`.`repo_id`,`commit_summary`.`author_user_id`,`commit_summary`.`customer_id`,`commit_summary`.`ref_type`,`commit_summary`.`additions`,`commit_summary`.`deletions`,`commit_summary`.`files_changed`,`commit_summary`.`branch`,`commit_summary`.`language`,`commit_summary`.`date`,`commit_summary`.`message` FROM `commit_summary` WHERE `id` = ?"
	var _ID sql.NullString
	var _CommitID sql.NullString
	var _RepoID sql.NullString
	var _AuthorUserID sql.NullString
	var _CustomerID sql.NullString
	var _RefType sql.NullString
	var _Additions sql.NullInt64
	var _Deletions sql.NullInt64
	var _FilesChanged sql.NullInt64
	var _Branch sql.NullString
	var _Language sql.NullString
	var _Date sql.NullInt64
	var _Message sql.NullString
	err := db.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_CommitID,
		&_RepoID,
		&_AuthorUserID,
		&_CustomerID,
		&_RefType,
		&_Additions,
		&_Deletions,
		&_FilesChanged,
		&_Branch,
		&_Language,
		&_Date,
		&_Message,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &CommitSummary{}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _CommitID.Valid {
		t.SetCommitID(_CommitID.String)
	}
	if _RepoID.Valid {
		t.SetRepoID(_RepoID.String)
	}
	if _AuthorUserID.Valid {
		t.SetAuthorUserID(_AuthorUserID.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _RefType.Valid {
		t.SetRefType(_RefType.String)
	}
	if _Additions.Valid {
		t.SetAdditions(int32(_Additions.Int64))
	}
	if _Deletions.Valid {
		t.SetDeletions(int32(_Deletions.Int64))
	}
	if _FilesChanged.Valid {
		t.SetFilesChanged(int32(_FilesChanged.Int64))
	}
	if _Branch.Valid {
		t.SetBranch(_Branch.String)
	}
	if _Language.Valid {
		t.SetLanguage(_Language.String)
	}
	if _Date.Valid {
		t.SetDate(_Date.Int64)
	}
	if _Message.Valid {
		t.SetMessage(_Message.String)
	}
	return t, nil
}

// FindCommitSummaryByIDTx will find a CommitSummary by ID using the provided transaction
func FindCommitSummaryByIDTx(ctx context.Context, tx *sql.Tx, value string) (*CommitSummary, error) {
	q := "SELECT `commit_summary`.`id`,`commit_summary`.`commit_id`,`commit_summary`.`repo_id`,`commit_summary`.`author_user_id`,`commit_summary`.`customer_id`,`commit_summary`.`ref_type`,`commit_summary`.`additions`,`commit_summary`.`deletions`,`commit_summary`.`files_changed`,`commit_summary`.`branch`,`commit_summary`.`language`,`commit_summary`.`date`,`commit_summary`.`message` FROM `commit_summary` WHERE `id` = ?"
	var _ID sql.NullString
	var _CommitID sql.NullString
	var _RepoID sql.NullString
	var _AuthorUserID sql.NullString
	var _CustomerID sql.NullString
	var _RefType sql.NullString
	var _Additions sql.NullInt64
	var _Deletions sql.NullInt64
	var _FilesChanged sql.NullInt64
	var _Branch sql.NullString
	var _Language sql.NullString
	var _Date sql.NullInt64
	var _Message sql.NullString
	err := tx.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_CommitID,
		&_RepoID,
		&_AuthorUserID,
		&_CustomerID,
		&_RefType,
		&_Additions,
		&_Deletions,
		&_FilesChanged,
		&_Branch,
		&_Language,
		&_Date,
		&_Message,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &CommitSummary{}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _CommitID.Valid {
		t.SetCommitID(_CommitID.String)
	}
	if _RepoID.Valid {
		t.SetRepoID(_RepoID.String)
	}
	if _AuthorUserID.Valid {
		t.SetAuthorUserID(_AuthorUserID.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _RefType.Valid {
		t.SetRefType(_RefType.String)
	}
	if _Additions.Valid {
		t.SetAdditions(int32(_Additions.Int64))
	}
	if _Deletions.Valid {
		t.SetDeletions(int32(_Deletions.Int64))
	}
	if _FilesChanged.Valid {
		t.SetFilesChanged(int32(_FilesChanged.Int64))
	}
	if _Branch.Valid {
		t.SetBranch(_Branch.String)
	}
	if _Language.Valid {
		t.SetLanguage(_Language.String)
	}
	if _Date.Valid {
		t.SetDate(_Date.Int64)
	}
	if _Message.Valid {
		t.SetMessage(_Message.String)
	}
	return t, nil
}

// GetCommitID will return the CommitSummary CommitID value
func (t *CommitSummary) GetCommitID() string {
	return t.CommitID
}

// SetCommitID will set the CommitSummary CommitID value
func (t *CommitSummary) SetCommitID(v string) {
	t.CommitID = v
}

// FindCommitSummariesByCommitID will find all CommitSummarys by the CommitID value
func FindCommitSummariesByCommitID(ctx context.Context, db *sql.DB, value string) ([]*CommitSummary, error) {
	q := "SELECT `commit_summary`.`id`,`commit_summary`.`commit_id`,`commit_summary`.`repo_id`,`commit_summary`.`author_user_id`,`commit_summary`.`customer_id`,`commit_summary`.`ref_type`,`commit_summary`.`additions`,`commit_summary`.`deletions`,`commit_summary`.`files_changed`,`commit_summary`.`branch`,`commit_summary`.`language`,`commit_summary`.`date`,`commit_summary`.`message` FROM `commit_summary` WHERE `commit_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*CommitSummary, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _CommitID sql.NullString
		var _RepoID sql.NullString
		var _AuthorUserID sql.NullString
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _Additions sql.NullInt64
		var _Deletions sql.NullInt64
		var _FilesChanged sql.NullInt64
		var _Branch sql.NullString
		var _Language sql.NullString
		var _Date sql.NullInt64
		var _Message sql.NullString
		err := rows.Scan(
			&_ID,
			&_CommitID,
			&_RepoID,
			&_AuthorUserID,
			&_CustomerID,
			&_RefType,
			&_Additions,
			&_Deletions,
			&_FilesChanged,
			&_Branch,
			&_Language,
			&_Date,
			&_Message,
		)
		if err != nil {
			return nil, err
		}
		t := &CommitSummary{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _CommitID.Valid {
			t.SetCommitID(_CommitID.String)
		}
		if _RepoID.Valid {
			t.SetRepoID(_RepoID.String)
		}
		if _AuthorUserID.Valid {
			t.SetAuthorUserID(_AuthorUserID.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefType.Valid {
			t.SetRefType(_RefType.String)
		}
		if _Additions.Valid {
			t.SetAdditions(int32(_Additions.Int64))
		}
		if _Deletions.Valid {
			t.SetDeletions(int32(_Deletions.Int64))
		}
		if _FilesChanged.Valid {
			t.SetFilesChanged(int32(_FilesChanged.Int64))
		}
		if _Branch.Valid {
			t.SetBranch(_Branch.String)
		}
		if _Language.Valid {
			t.SetLanguage(_Language.String)
		}
		if _Date.Valid {
			t.SetDate(_Date.Int64)
		}
		if _Message.Valid {
			t.SetMessage(_Message.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindCommitSummariesByCommitIDTx will find all CommitSummarys by the CommitID value using the provided transaction
func FindCommitSummariesByCommitIDTx(ctx context.Context, tx *sql.Tx, value string) ([]*CommitSummary, error) {
	q := "SELECT `commit_summary`.`id`,`commit_summary`.`commit_id`,`commit_summary`.`repo_id`,`commit_summary`.`author_user_id`,`commit_summary`.`customer_id`,`commit_summary`.`ref_type`,`commit_summary`.`additions`,`commit_summary`.`deletions`,`commit_summary`.`files_changed`,`commit_summary`.`branch`,`commit_summary`.`language`,`commit_summary`.`date`,`commit_summary`.`message` FROM `commit_summary` WHERE `commit_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*CommitSummary, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _CommitID sql.NullString
		var _RepoID sql.NullString
		var _AuthorUserID sql.NullString
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _Additions sql.NullInt64
		var _Deletions sql.NullInt64
		var _FilesChanged sql.NullInt64
		var _Branch sql.NullString
		var _Language sql.NullString
		var _Date sql.NullInt64
		var _Message sql.NullString
		err := rows.Scan(
			&_ID,
			&_CommitID,
			&_RepoID,
			&_AuthorUserID,
			&_CustomerID,
			&_RefType,
			&_Additions,
			&_Deletions,
			&_FilesChanged,
			&_Branch,
			&_Language,
			&_Date,
			&_Message,
		)
		if err != nil {
			return nil, err
		}
		t := &CommitSummary{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _CommitID.Valid {
			t.SetCommitID(_CommitID.String)
		}
		if _RepoID.Valid {
			t.SetRepoID(_RepoID.String)
		}
		if _AuthorUserID.Valid {
			t.SetAuthorUserID(_AuthorUserID.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefType.Valid {
			t.SetRefType(_RefType.String)
		}
		if _Additions.Valid {
			t.SetAdditions(int32(_Additions.Int64))
		}
		if _Deletions.Valid {
			t.SetDeletions(int32(_Deletions.Int64))
		}
		if _FilesChanged.Valid {
			t.SetFilesChanged(int32(_FilesChanged.Int64))
		}
		if _Branch.Valid {
			t.SetBranch(_Branch.String)
		}
		if _Language.Valid {
			t.SetLanguage(_Language.String)
		}
		if _Date.Valid {
			t.SetDate(_Date.Int64)
		}
		if _Message.Valid {
			t.SetMessage(_Message.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetRepoID will return the CommitSummary RepoID value
func (t *CommitSummary) GetRepoID() string {
	return t.RepoID
}

// SetRepoID will set the CommitSummary RepoID value
func (t *CommitSummary) SetRepoID(v string) {
	t.RepoID = v
}

// FindCommitSummariesByRepoID will find all CommitSummarys by the RepoID value
func FindCommitSummariesByRepoID(ctx context.Context, db *sql.DB, value string) ([]*CommitSummary, error) {
	q := "SELECT `commit_summary`.`id`,`commit_summary`.`commit_id`,`commit_summary`.`repo_id`,`commit_summary`.`author_user_id`,`commit_summary`.`customer_id`,`commit_summary`.`ref_type`,`commit_summary`.`additions`,`commit_summary`.`deletions`,`commit_summary`.`files_changed`,`commit_summary`.`branch`,`commit_summary`.`language`,`commit_summary`.`date`,`commit_summary`.`message` FROM `commit_summary` WHERE `repo_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*CommitSummary, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _CommitID sql.NullString
		var _RepoID sql.NullString
		var _AuthorUserID sql.NullString
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _Additions sql.NullInt64
		var _Deletions sql.NullInt64
		var _FilesChanged sql.NullInt64
		var _Branch sql.NullString
		var _Language sql.NullString
		var _Date sql.NullInt64
		var _Message sql.NullString
		err := rows.Scan(
			&_ID,
			&_CommitID,
			&_RepoID,
			&_AuthorUserID,
			&_CustomerID,
			&_RefType,
			&_Additions,
			&_Deletions,
			&_FilesChanged,
			&_Branch,
			&_Language,
			&_Date,
			&_Message,
		)
		if err != nil {
			return nil, err
		}
		t := &CommitSummary{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _CommitID.Valid {
			t.SetCommitID(_CommitID.String)
		}
		if _RepoID.Valid {
			t.SetRepoID(_RepoID.String)
		}
		if _AuthorUserID.Valid {
			t.SetAuthorUserID(_AuthorUserID.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefType.Valid {
			t.SetRefType(_RefType.String)
		}
		if _Additions.Valid {
			t.SetAdditions(int32(_Additions.Int64))
		}
		if _Deletions.Valid {
			t.SetDeletions(int32(_Deletions.Int64))
		}
		if _FilesChanged.Valid {
			t.SetFilesChanged(int32(_FilesChanged.Int64))
		}
		if _Branch.Valid {
			t.SetBranch(_Branch.String)
		}
		if _Language.Valid {
			t.SetLanguage(_Language.String)
		}
		if _Date.Valid {
			t.SetDate(_Date.Int64)
		}
		if _Message.Valid {
			t.SetMessage(_Message.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindCommitSummariesByRepoIDTx will find all CommitSummarys by the RepoID value using the provided transaction
func FindCommitSummariesByRepoIDTx(ctx context.Context, tx *sql.Tx, value string) ([]*CommitSummary, error) {
	q := "SELECT `commit_summary`.`id`,`commit_summary`.`commit_id`,`commit_summary`.`repo_id`,`commit_summary`.`author_user_id`,`commit_summary`.`customer_id`,`commit_summary`.`ref_type`,`commit_summary`.`additions`,`commit_summary`.`deletions`,`commit_summary`.`files_changed`,`commit_summary`.`branch`,`commit_summary`.`language`,`commit_summary`.`date`,`commit_summary`.`message` FROM `commit_summary` WHERE `repo_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*CommitSummary, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _CommitID sql.NullString
		var _RepoID sql.NullString
		var _AuthorUserID sql.NullString
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _Additions sql.NullInt64
		var _Deletions sql.NullInt64
		var _FilesChanged sql.NullInt64
		var _Branch sql.NullString
		var _Language sql.NullString
		var _Date sql.NullInt64
		var _Message sql.NullString
		err := rows.Scan(
			&_ID,
			&_CommitID,
			&_RepoID,
			&_AuthorUserID,
			&_CustomerID,
			&_RefType,
			&_Additions,
			&_Deletions,
			&_FilesChanged,
			&_Branch,
			&_Language,
			&_Date,
			&_Message,
		)
		if err != nil {
			return nil, err
		}
		t := &CommitSummary{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _CommitID.Valid {
			t.SetCommitID(_CommitID.String)
		}
		if _RepoID.Valid {
			t.SetRepoID(_RepoID.String)
		}
		if _AuthorUserID.Valid {
			t.SetAuthorUserID(_AuthorUserID.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefType.Valid {
			t.SetRefType(_RefType.String)
		}
		if _Additions.Valid {
			t.SetAdditions(int32(_Additions.Int64))
		}
		if _Deletions.Valid {
			t.SetDeletions(int32(_Deletions.Int64))
		}
		if _FilesChanged.Valid {
			t.SetFilesChanged(int32(_FilesChanged.Int64))
		}
		if _Branch.Valid {
			t.SetBranch(_Branch.String)
		}
		if _Language.Valid {
			t.SetLanguage(_Language.String)
		}
		if _Date.Valid {
			t.SetDate(_Date.Int64)
		}
		if _Message.Valid {
			t.SetMessage(_Message.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetAuthorUserID will return the CommitSummary AuthorUserID value
func (t *CommitSummary) GetAuthorUserID() string {
	if t.AuthorUserID == nil {
		return ""
	}
	return *t.AuthorUserID
}

// SetAuthorUserID will set the CommitSummary AuthorUserID value
func (t *CommitSummary) SetAuthorUserID(v string) {
	t.AuthorUserID = &v
}

// FindCommitSummariesByAuthorUserID will find all CommitSummarys by the AuthorUserID value
func FindCommitSummariesByAuthorUserID(ctx context.Context, db *sql.DB, value string) ([]*CommitSummary, error) {
	q := "SELECT `commit_summary`.`id`,`commit_summary`.`commit_id`,`commit_summary`.`repo_id`,`commit_summary`.`author_user_id`,`commit_summary`.`customer_id`,`commit_summary`.`ref_type`,`commit_summary`.`additions`,`commit_summary`.`deletions`,`commit_summary`.`files_changed`,`commit_summary`.`branch`,`commit_summary`.`language`,`commit_summary`.`date`,`commit_summary`.`message` FROM `commit_summary` WHERE `author_user_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*CommitSummary, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _CommitID sql.NullString
		var _RepoID sql.NullString
		var _AuthorUserID sql.NullString
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _Additions sql.NullInt64
		var _Deletions sql.NullInt64
		var _FilesChanged sql.NullInt64
		var _Branch sql.NullString
		var _Language sql.NullString
		var _Date sql.NullInt64
		var _Message sql.NullString
		err := rows.Scan(
			&_ID,
			&_CommitID,
			&_RepoID,
			&_AuthorUserID,
			&_CustomerID,
			&_RefType,
			&_Additions,
			&_Deletions,
			&_FilesChanged,
			&_Branch,
			&_Language,
			&_Date,
			&_Message,
		)
		if err != nil {
			return nil, err
		}
		t := &CommitSummary{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _CommitID.Valid {
			t.SetCommitID(_CommitID.String)
		}
		if _RepoID.Valid {
			t.SetRepoID(_RepoID.String)
		}
		if _AuthorUserID.Valid {
			t.SetAuthorUserID(_AuthorUserID.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefType.Valid {
			t.SetRefType(_RefType.String)
		}
		if _Additions.Valid {
			t.SetAdditions(int32(_Additions.Int64))
		}
		if _Deletions.Valid {
			t.SetDeletions(int32(_Deletions.Int64))
		}
		if _FilesChanged.Valid {
			t.SetFilesChanged(int32(_FilesChanged.Int64))
		}
		if _Branch.Valid {
			t.SetBranch(_Branch.String)
		}
		if _Language.Valid {
			t.SetLanguage(_Language.String)
		}
		if _Date.Valid {
			t.SetDate(_Date.Int64)
		}
		if _Message.Valid {
			t.SetMessage(_Message.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindCommitSummariesByAuthorUserIDTx will find all CommitSummarys by the AuthorUserID value using the provided transaction
func FindCommitSummariesByAuthorUserIDTx(ctx context.Context, tx *sql.Tx, value string) ([]*CommitSummary, error) {
	q := "SELECT `commit_summary`.`id`,`commit_summary`.`commit_id`,`commit_summary`.`repo_id`,`commit_summary`.`author_user_id`,`commit_summary`.`customer_id`,`commit_summary`.`ref_type`,`commit_summary`.`additions`,`commit_summary`.`deletions`,`commit_summary`.`files_changed`,`commit_summary`.`branch`,`commit_summary`.`language`,`commit_summary`.`date`,`commit_summary`.`message` FROM `commit_summary` WHERE `author_user_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*CommitSummary, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _CommitID sql.NullString
		var _RepoID sql.NullString
		var _AuthorUserID sql.NullString
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _Additions sql.NullInt64
		var _Deletions sql.NullInt64
		var _FilesChanged sql.NullInt64
		var _Branch sql.NullString
		var _Language sql.NullString
		var _Date sql.NullInt64
		var _Message sql.NullString
		err := rows.Scan(
			&_ID,
			&_CommitID,
			&_RepoID,
			&_AuthorUserID,
			&_CustomerID,
			&_RefType,
			&_Additions,
			&_Deletions,
			&_FilesChanged,
			&_Branch,
			&_Language,
			&_Date,
			&_Message,
		)
		if err != nil {
			return nil, err
		}
		t := &CommitSummary{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _CommitID.Valid {
			t.SetCommitID(_CommitID.String)
		}
		if _RepoID.Valid {
			t.SetRepoID(_RepoID.String)
		}
		if _AuthorUserID.Valid {
			t.SetAuthorUserID(_AuthorUserID.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefType.Valid {
			t.SetRefType(_RefType.String)
		}
		if _Additions.Valid {
			t.SetAdditions(int32(_Additions.Int64))
		}
		if _Deletions.Valid {
			t.SetDeletions(int32(_Deletions.Int64))
		}
		if _FilesChanged.Valid {
			t.SetFilesChanged(int32(_FilesChanged.Int64))
		}
		if _Branch.Valid {
			t.SetBranch(_Branch.String)
		}
		if _Language.Valid {
			t.SetLanguage(_Language.String)
		}
		if _Date.Valid {
			t.SetDate(_Date.Int64)
		}
		if _Message.Valid {
			t.SetMessage(_Message.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetCustomerID will return the CommitSummary CustomerID value
func (t *CommitSummary) GetCustomerID() string {
	return t.CustomerID
}

// SetCustomerID will set the CommitSummary CustomerID value
func (t *CommitSummary) SetCustomerID(v string) {
	t.CustomerID = v
}

// FindCommitSummariesByCustomerID will find all CommitSummarys by the CustomerID value
func FindCommitSummariesByCustomerID(ctx context.Context, db *sql.DB, value string) ([]*CommitSummary, error) {
	q := "SELECT `commit_summary`.`id`,`commit_summary`.`commit_id`,`commit_summary`.`repo_id`,`commit_summary`.`author_user_id`,`commit_summary`.`customer_id`,`commit_summary`.`ref_type`,`commit_summary`.`additions`,`commit_summary`.`deletions`,`commit_summary`.`files_changed`,`commit_summary`.`branch`,`commit_summary`.`language`,`commit_summary`.`date`,`commit_summary`.`message` FROM `commit_summary` WHERE `customer_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*CommitSummary, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _CommitID sql.NullString
		var _RepoID sql.NullString
		var _AuthorUserID sql.NullString
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _Additions sql.NullInt64
		var _Deletions sql.NullInt64
		var _FilesChanged sql.NullInt64
		var _Branch sql.NullString
		var _Language sql.NullString
		var _Date sql.NullInt64
		var _Message sql.NullString
		err := rows.Scan(
			&_ID,
			&_CommitID,
			&_RepoID,
			&_AuthorUserID,
			&_CustomerID,
			&_RefType,
			&_Additions,
			&_Deletions,
			&_FilesChanged,
			&_Branch,
			&_Language,
			&_Date,
			&_Message,
		)
		if err != nil {
			return nil, err
		}
		t := &CommitSummary{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _CommitID.Valid {
			t.SetCommitID(_CommitID.String)
		}
		if _RepoID.Valid {
			t.SetRepoID(_RepoID.String)
		}
		if _AuthorUserID.Valid {
			t.SetAuthorUserID(_AuthorUserID.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefType.Valid {
			t.SetRefType(_RefType.String)
		}
		if _Additions.Valid {
			t.SetAdditions(int32(_Additions.Int64))
		}
		if _Deletions.Valid {
			t.SetDeletions(int32(_Deletions.Int64))
		}
		if _FilesChanged.Valid {
			t.SetFilesChanged(int32(_FilesChanged.Int64))
		}
		if _Branch.Valid {
			t.SetBranch(_Branch.String)
		}
		if _Language.Valid {
			t.SetLanguage(_Language.String)
		}
		if _Date.Valid {
			t.SetDate(_Date.Int64)
		}
		if _Message.Valid {
			t.SetMessage(_Message.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindCommitSummariesByCustomerIDTx will find all CommitSummarys by the CustomerID value using the provided transaction
func FindCommitSummariesByCustomerIDTx(ctx context.Context, tx *sql.Tx, value string) ([]*CommitSummary, error) {
	q := "SELECT `commit_summary`.`id`,`commit_summary`.`commit_id`,`commit_summary`.`repo_id`,`commit_summary`.`author_user_id`,`commit_summary`.`customer_id`,`commit_summary`.`ref_type`,`commit_summary`.`additions`,`commit_summary`.`deletions`,`commit_summary`.`files_changed`,`commit_summary`.`branch`,`commit_summary`.`language`,`commit_summary`.`date`,`commit_summary`.`message` FROM `commit_summary` WHERE `customer_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*CommitSummary, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _CommitID sql.NullString
		var _RepoID sql.NullString
		var _AuthorUserID sql.NullString
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _Additions sql.NullInt64
		var _Deletions sql.NullInt64
		var _FilesChanged sql.NullInt64
		var _Branch sql.NullString
		var _Language sql.NullString
		var _Date sql.NullInt64
		var _Message sql.NullString
		err := rows.Scan(
			&_ID,
			&_CommitID,
			&_RepoID,
			&_AuthorUserID,
			&_CustomerID,
			&_RefType,
			&_Additions,
			&_Deletions,
			&_FilesChanged,
			&_Branch,
			&_Language,
			&_Date,
			&_Message,
		)
		if err != nil {
			return nil, err
		}
		t := &CommitSummary{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _CommitID.Valid {
			t.SetCommitID(_CommitID.String)
		}
		if _RepoID.Valid {
			t.SetRepoID(_RepoID.String)
		}
		if _AuthorUserID.Valid {
			t.SetAuthorUserID(_AuthorUserID.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefType.Valid {
			t.SetRefType(_RefType.String)
		}
		if _Additions.Valid {
			t.SetAdditions(int32(_Additions.Int64))
		}
		if _Deletions.Valid {
			t.SetDeletions(int32(_Deletions.Int64))
		}
		if _FilesChanged.Valid {
			t.SetFilesChanged(int32(_FilesChanged.Int64))
		}
		if _Branch.Valid {
			t.SetBranch(_Branch.String)
		}
		if _Language.Valid {
			t.SetLanguage(_Language.String)
		}
		if _Date.Valid {
			t.SetDate(_Date.Int64)
		}
		if _Message.Valid {
			t.SetMessage(_Message.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetRefType will return the CommitSummary RefType value
func (t *CommitSummary) GetRefType() string {
	return t.RefType
}

// SetRefType will set the CommitSummary RefType value
func (t *CommitSummary) SetRefType(v string) {
	t.RefType = v
}

// FindCommitSummariesByRefType will find all CommitSummarys by the RefType value
func FindCommitSummariesByRefType(ctx context.Context, db *sql.DB, value string) ([]*CommitSummary, error) {
	q := "SELECT `commit_summary`.`id`,`commit_summary`.`commit_id`,`commit_summary`.`repo_id`,`commit_summary`.`author_user_id`,`commit_summary`.`customer_id`,`commit_summary`.`ref_type`,`commit_summary`.`additions`,`commit_summary`.`deletions`,`commit_summary`.`files_changed`,`commit_summary`.`branch`,`commit_summary`.`language`,`commit_summary`.`date`,`commit_summary`.`message` FROM `commit_summary` WHERE `ref_type` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*CommitSummary, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _CommitID sql.NullString
		var _RepoID sql.NullString
		var _AuthorUserID sql.NullString
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _Additions sql.NullInt64
		var _Deletions sql.NullInt64
		var _FilesChanged sql.NullInt64
		var _Branch sql.NullString
		var _Language sql.NullString
		var _Date sql.NullInt64
		var _Message sql.NullString
		err := rows.Scan(
			&_ID,
			&_CommitID,
			&_RepoID,
			&_AuthorUserID,
			&_CustomerID,
			&_RefType,
			&_Additions,
			&_Deletions,
			&_FilesChanged,
			&_Branch,
			&_Language,
			&_Date,
			&_Message,
		)
		if err != nil {
			return nil, err
		}
		t := &CommitSummary{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _CommitID.Valid {
			t.SetCommitID(_CommitID.String)
		}
		if _RepoID.Valid {
			t.SetRepoID(_RepoID.String)
		}
		if _AuthorUserID.Valid {
			t.SetAuthorUserID(_AuthorUserID.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefType.Valid {
			t.SetRefType(_RefType.String)
		}
		if _Additions.Valid {
			t.SetAdditions(int32(_Additions.Int64))
		}
		if _Deletions.Valid {
			t.SetDeletions(int32(_Deletions.Int64))
		}
		if _FilesChanged.Valid {
			t.SetFilesChanged(int32(_FilesChanged.Int64))
		}
		if _Branch.Valid {
			t.SetBranch(_Branch.String)
		}
		if _Language.Valid {
			t.SetLanguage(_Language.String)
		}
		if _Date.Valid {
			t.SetDate(_Date.Int64)
		}
		if _Message.Valid {
			t.SetMessage(_Message.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindCommitSummariesByRefTypeTx will find all CommitSummarys by the RefType value using the provided transaction
func FindCommitSummariesByRefTypeTx(ctx context.Context, tx *sql.Tx, value string) ([]*CommitSummary, error) {
	q := "SELECT `commit_summary`.`id`,`commit_summary`.`commit_id`,`commit_summary`.`repo_id`,`commit_summary`.`author_user_id`,`commit_summary`.`customer_id`,`commit_summary`.`ref_type`,`commit_summary`.`additions`,`commit_summary`.`deletions`,`commit_summary`.`files_changed`,`commit_summary`.`branch`,`commit_summary`.`language`,`commit_summary`.`date`,`commit_summary`.`message` FROM `commit_summary` WHERE `ref_type` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*CommitSummary, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _CommitID sql.NullString
		var _RepoID sql.NullString
		var _AuthorUserID sql.NullString
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _Additions sql.NullInt64
		var _Deletions sql.NullInt64
		var _FilesChanged sql.NullInt64
		var _Branch sql.NullString
		var _Language sql.NullString
		var _Date sql.NullInt64
		var _Message sql.NullString
		err := rows.Scan(
			&_ID,
			&_CommitID,
			&_RepoID,
			&_AuthorUserID,
			&_CustomerID,
			&_RefType,
			&_Additions,
			&_Deletions,
			&_FilesChanged,
			&_Branch,
			&_Language,
			&_Date,
			&_Message,
		)
		if err != nil {
			return nil, err
		}
		t := &CommitSummary{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _CommitID.Valid {
			t.SetCommitID(_CommitID.String)
		}
		if _RepoID.Valid {
			t.SetRepoID(_RepoID.String)
		}
		if _AuthorUserID.Valid {
			t.SetAuthorUserID(_AuthorUserID.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefType.Valid {
			t.SetRefType(_RefType.String)
		}
		if _Additions.Valid {
			t.SetAdditions(int32(_Additions.Int64))
		}
		if _Deletions.Valid {
			t.SetDeletions(int32(_Deletions.Int64))
		}
		if _FilesChanged.Valid {
			t.SetFilesChanged(int32(_FilesChanged.Int64))
		}
		if _Branch.Valid {
			t.SetBranch(_Branch.String)
		}
		if _Language.Valid {
			t.SetLanguage(_Language.String)
		}
		if _Date.Valid {
			t.SetDate(_Date.Int64)
		}
		if _Message.Valid {
			t.SetMessage(_Message.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetAdditions will return the CommitSummary Additions value
func (t *CommitSummary) GetAdditions() int32 {
	return t.Additions
}

// SetAdditions will set the CommitSummary Additions value
func (t *CommitSummary) SetAdditions(v int32) {
	t.Additions = v
}

// GetDeletions will return the CommitSummary Deletions value
func (t *CommitSummary) GetDeletions() int32 {
	return t.Deletions
}

// SetDeletions will set the CommitSummary Deletions value
func (t *CommitSummary) SetDeletions(v int32) {
	t.Deletions = v
}

// GetFilesChanged will return the CommitSummary FilesChanged value
func (t *CommitSummary) GetFilesChanged() int32 {
	return t.FilesChanged
}

// SetFilesChanged will set the CommitSummary FilesChanged value
func (t *CommitSummary) SetFilesChanged(v int32) {
	t.FilesChanged = v
}

// GetBranch will return the CommitSummary Branch value
func (t *CommitSummary) GetBranch() string {
	if t.Branch == nil {
		return ""
	}
	return *t.Branch
}

// SetBranch will set the CommitSummary Branch value
func (t *CommitSummary) SetBranch(v string) {
	t.Branch = &v
}

// GetLanguage will return the CommitSummary Language value
func (t *CommitSummary) GetLanguage() string {
	return t.Language
}

// SetLanguage will set the CommitSummary Language value
func (t *CommitSummary) SetLanguage(v string) {
	t.Language = v
}

// GetDate will return the CommitSummary Date value
func (t *CommitSummary) GetDate() int64 {
	return t.Date
}

// SetDate will set the CommitSummary Date value
func (t *CommitSummary) SetDate(v int64) {
	t.Date = v
}

// GetMessage will return the CommitSummary Message value
func (t *CommitSummary) GetMessage() string {
	if t.Message == nil {
		return ""
	}
	return *t.Message
}

// SetMessage will set the CommitSummary Message value
func (t *CommitSummary) SetMessage(v string) {
	t.Message = &v
}

func (t *CommitSummary) toTimestamp(value time.Time) *timestamp.Timestamp {
	ts, _ := ptypes.TimestampProto(value)
	return ts
}

// DBCreateCommitSummaryTable will create the CommitSummary table
func DBCreateCommitSummaryTable(ctx context.Context, db *sql.DB) error {
	q := "CREATE TABLE `commit_summary` (`id` VARCHAR(64) NOT NULL PRIMARY KEY,`commit_id`VARCHAR(64) NOT NULL,`repo_id` VARCHAR(64) NOT NULL,`author_user_id` VARCHAR(64),`customer_id` VARCHAR(64) NOT NULL,`ref_type` VARCHAR(20) NOT NULL,`additions`INT NOT NULL DEFAULT 0,`deletions`INT NOT NULL DEFAULT 0,`files_changed` INT NOT NULL DEFAULT 0,`branch`VARCHAR(255) DEFAULT \"master\",`language` VARCHAR(500) NOT NULL DEFAULT \"unknown\",`date` BIGINT UNSIGNED NOT NULL,`message` LONGTEXT,INDEX commit_summary_commit_id_index (`commit_id`),INDEX commit_summary_repo_id_index (`repo_id`),INDEX commit_summary_author_user_id_index (`author_user_id`),INDEX commit_summary_customer_id_index (`customer_id`),INDEX commit_summary_ref_type_index (`ref_type`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBCreateCommitSummaryTableTx will create the CommitSummary table using the provided transction
func DBCreateCommitSummaryTableTx(ctx context.Context, tx *sql.Tx) error {
	q := "CREATE TABLE `commit_summary` (`id` VARCHAR(64) NOT NULL PRIMARY KEY,`commit_id`VARCHAR(64) NOT NULL,`repo_id` VARCHAR(64) NOT NULL,`author_user_id` VARCHAR(64),`customer_id` VARCHAR(64) NOT NULL,`ref_type` VARCHAR(20) NOT NULL,`additions`INT NOT NULL DEFAULT 0,`deletions`INT NOT NULL DEFAULT 0,`files_changed` INT NOT NULL DEFAULT 0,`branch`VARCHAR(255) DEFAULT \"master\",`language` VARCHAR(500) NOT NULL DEFAULT \"unknown\",`date` BIGINT UNSIGNED NOT NULL,`message` LONGTEXT,INDEX commit_summary_commit_id_index (`commit_id`),INDEX commit_summary_repo_id_index (`repo_id`),INDEX commit_summary_author_user_id_index (`author_user_id`),INDEX commit_summary_customer_id_index (`customer_id`),INDEX commit_summary_ref_type_index (`ref_type`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// DBDropCommitSummaryTable will drop the CommitSummary table
func DBDropCommitSummaryTable(ctx context.Context, db *sql.DB) error {
	q := "DROP TABLE IF EXISTS `commit_summary`"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBDropCommitSummaryTableTx will drop the CommitSummary table using the provided transaction
func DBDropCommitSummaryTableTx(ctx context.Context, tx *sql.Tx) error {
	q := "DROP TABLE IF EXISTS `commit_summary`"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// DBCreate will create a new CommitSummary record in the database
func (t *CommitSummary) DBCreate(ctx context.Context, db *sql.DB) (sql.Result, error) {
	q := "INSERT INTO `commit_summary` (`commit_summary`.`id`,`commit_summary`.`commit_id`,`commit_summary`.`repo_id`,`commit_summary`.`author_user_id`,`commit_summary`.`customer_id`,`commit_summary`.`ref_type`,`commit_summary`.`additions`,`commit_summary`.`deletions`,`commit_summary`.`files_changed`,`commit_summary`.`branch`,`commit_summary`.`language`,`commit_summary`.`date`,`commit_summary`.`message`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?)"
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.CommitID),
		orm.ToSQLString(t.RepoID),
		orm.ToSQLString(t.AuthorUserID),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefType),
		orm.ToSQLInt64(t.Additions),
		orm.ToSQLInt64(t.Deletions),
		orm.ToSQLInt64(t.FilesChanged),
		orm.ToSQLString(t.Branch),
		orm.ToSQLString(t.Language),
		orm.ToSQLInt64(t.Date),
		orm.ToSQLString(t.Message),
	)
}

// DBCreateTx will create a new CommitSummary record in the database using the provided transaction
func (t *CommitSummary) DBCreateTx(ctx context.Context, tx *sql.Tx) (sql.Result, error) {
	q := "INSERT INTO `commit_summary` (`commit_summary`.`id`,`commit_summary`.`commit_id`,`commit_summary`.`repo_id`,`commit_summary`.`author_user_id`,`commit_summary`.`customer_id`,`commit_summary`.`ref_type`,`commit_summary`.`additions`,`commit_summary`.`deletions`,`commit_summary`.`files_changed`,`commit_summary`.`branch`,`commit_summary`.`language`,`commit_summary`.`date`,`commit_summary`.`message`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?)"
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.CommitID),
		orm.ToSQLString(t.RepoID),
		orm.ToSQLString(t.AuthorUserID),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefType),
		orm.ToSQLInt64(t.Additions),
		orm.ToSQLInt64(t.Deletions),
		orm.ToSQLInt64(t.FilesChanged),
		orm.ToSQLString(t.Branch),
		orm.ToSQLString(t.Language),
		orm.ToSQLInt64(t.Date),
		orm.ToSQLString(t.Message),
	)
}

// DBCreateIgnoreDuplicate will upsert the CommitSummary record in the database
func (t *CommitSummary) DBCreateIgnoreDuplicate(ctx context.Context, db *sql.DB) (sql.Result, error) {
	q := "INSERT INTO `commit_summary` (`commit_summary`.`id`,`commit_summary`.`commit_id`,`commit_summary`.`repo_id`,`commit_summary`.`author_user_id`,`commit_summary`.`customer_id`,`commit_summary`.`ref_type`,`commit_summary`.`additions`,`commit_summary`.`deletions`,`commit_summary`.`files_changed`,`commit_summary`.`branch`,`commit_summary`.`language`,`commit_summary`.`date`,`commit_summary`.`message`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.CommitID),
		orm.ToSQLString(t.RepoID),
		orm.ToSQLString(t.AuthorUserID),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefType),
		orm.ToSQLInt64(t.Additions),
		orm.ToSQLInt64(t.Deletions),
		orm.ToSQLInt64(t.FilesChanged),
		orm.ToSQLString(t.Branch),
		orm.ToSQLString(t.Language),
		orm.ToSQLInt64(t.Date),
		orm.ToSQLString(t.Message),
	)
}

// DBCreateIgnoreDuplicateTx will upsert the CommitSummary record in the database using the provided transaction
func (t *CommitSummary) DBCreateIgnoreDuplicateTx(ctx context.Context, tx *sql.Tx) (sql.Result, error) {
	q := "INSERT INTO `commit_summary` (`commit_summary`.`id`,`commit_summary`.`commit_id`,`commit_summary`.`repo_id`,`commit_summary`.`author_user_id`,`commit_summary`.`customer_id`,`commit_summary`.`ref_type`,`commit_summary`.`additions`,`commit_summary`.`deletions`,`commit_summary`.`files_changed`,`commit_summary`.`branch`,`commit_summary`.`language`,`commit_summary`.`date`,`commit_summary`.`message`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.CommitID),
		orm.ToSQLString(t.RepoID),
		orm.ToSQLString(t.AuthorUserID),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefType),
		orm.ToSQLInt64(t.Additions),
		orm.ToSQLInt64(t.Deletions),
		orm.ToSQLInt64(t.FilesChanged),
		orm.ToSQLString(t.Branch),
		orm.ToSQLString(t.Language),
		orm.ToSQLInt64(t.Date),
		orm.ToSQLString(t.Message),
	)
}

// DeleteAllCommitSummaries deletes all CommitSummary records in the database with optional filters
func DeleteAllCommitSummaries(ctx context.Context, db *sql.DB, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(CommitSummaryTableName),
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

// DeleteAllCommitSummariesTx deletes all CommitSummary records in the database with optional filters using the provided transaction
func DeleteAllCommitSummariesTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(CommitSummaryTableName),
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

// DBDelete will delete this CommitSummary record in the database
func (t *CommitSummary) DBDelete(ctx context.Context, db *sql.DB) (bool, error) {
	q := "DELETE FROM `commit_summary` WHERE `id` = ?"
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

// DBDeleteTx will delete this CommitSummary record in the database using the provided transaction
func (t *CommitSummary) DBDeleteTx(ctx context.Context, tx *sql.Tx) (bool, error) {
	q := "DELETE FROM `commit_summary` WHERE `id` = ?"
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

// DBUpdate will update the CommitSummary record in the database
func (t *CommitSummary) DBUpdate(ctx context.Context, db *sql.DB) (sql.Result, error) {
	q := "UPDATE `commit_summary` SET `commit_id`=?,`repo_id`=?,`author_user_id`=?,`customer_id`=?,`ref_type`=?,`additions`=?,`deletions`=?,`files_changed`=?,`branch`=?,`language`=?,`date`=?,`message`=? WHERE `id`=?"
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.CommitID),
		orm.ToSQLString(t.RepoID),
		orm.ToSQLString(t.AuthorUserID),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefType),
		orm.ToSQLInt64(t.Additions),
		orm.ToSQLInt64(t.Deletions),
		orm.ToSQLInt64(t.FilesChanged),
		orm.ToSQLString(t.Branch),
		orm.ToSQLString(t.Language),
		orm.ToSQLInt64(t.Date),
		orm.ToSQLString(t.Message),
		orm.ToSQLString(t.ID),
	)
}

// DBUpdateTx will update the CommitSummary record in the database using the provided transaction
func (t *CommitSummary) DBUpdateTx(ctx context.Context, tx *sql.Tx) (sql.Result, error) {
	q := "UPDATE `commit_summary` SET `commit_id`=?,`repo_id`=?,`author_user_id`=?,`customer_id`=?,`ref_type`=?,`additions`=?,`deletions`=?,`files_changed`=?,`branch`=?,`language`=?,`date`=?,`message`=? WHERE `id`=?"
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.CommitID),
		orm.ToSQLString(t.RepoID),
		orm.ToSQLString(t.AuthorUserID),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefType),
		orm.ToSQLInt64(t.Additions),
		orm.ToSQLInt64(t.Deletions),
		orm.ToSQLInt64(t.FilesChanged),
		orm.ToSQLString(t.Branch),
		orm.ToSQLString(t.Language),
		orm.ToSQLInt64(t.Date),
		orm.ToSQLString(t.Message),
		orm.ToSQLString(t.ID),
	)
}

// DBUpsert will upsert the CommitSummary record in the database
func (t *CommitSummary) DBUpsert(ctx context.Context, db *sql.DB, conditions ...interface{}) (bool, bool, error) {
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `commit_summary` (`commit_summary`.`id`,`commit_summary`.`commit_id`,`commit_summary`.`repo_id`,`commit_summary`.`author_user_id`,`commit_summary`.`customer_id`,`commit_summary`.`ref_type`,`commit_summary`.`additions`,`commit_summary`.`deletions`,`commit_summary`.`files_changed`,`commit_summary`.`branch`,`commit_summary`.`language`,`commit_summary`.`date`,`commit_summary`.`message`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `commit_summary` (`commit_summary`.`id`,`commit_summary`.`commit_id`,`commit_summary`.`repo_id`,`commit_summary`.`author_user_id`,`commit_summary`.`customer_id`,`commit_summary`.`ref_type`,`commit_summary`.`additions`,`commit_summary`.`deletions`,`commit_summary`.`files_changed`,`commit_summary`.`branch`,`commit_summary`.`language`,`commit_summary`.`date`,`commit_summary`.`message`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `commit_id`=VALUES(`commit_id`),`repo_id`=VALUES(`repo_id`),`author_user_id`=VALUES(`author_user_id`),`customer_id`=VALUES(`customer_id`),`ref_type`=VALUES(`ref_type`),`additions`=VALUES(`additions`),`deletions`=VALUES(`deletions`),`files_changed`=VALUES(`files_changed`),`branch`=VALUES(`branch`),`language`=VALUES(`language`),`date`=VALUES(`date`),`message`=VALUES(`message`)"
	}
	r, err := db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.CommitID),
		orm.ToSQLString(t.RepoID),
		orm.ToSQLString(t.AuthorUserID),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefType),
		orm.ToSQLInt64(t.Additions),
		orm.ToSQLInt64(t.Deletions),
		orm.ToSQLInt64(t.FilesChanged),
		orm.ToSQLString(t.Branch),
		orm.ToSQLString(t.Language),
		orm.ToSQLInt64(t.Date),
		orm.ToSQLString(t.Message),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBUpsertTx will upsert the CommitSummary record in the database using the provided transaction
func (t *CommitSummary) DBUpsertTx(ctx context.Context, tx *sql.Tx, conditions ...interface{}) (bool, bool, error) {
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `commit_summary` (`commit_summary`.`id`,`commit_summary`.`commit_id`,`commit_summary`.`repo_id`,`commit_summary`.`author_user_id`,`commit_summary`.`customer_id`,`commit_summary`.`ref_type`,`commit_summary`.`additions`,`commit_summary`.`deletions`,`commit_summary`.`files_changed`,`commit_summary`.`branch`,`commit_summary`.`language`,`commit_summary`.`date`,`commit_summary`.`message`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `commit_summary` (`commit_summary`.`id`,`commit_summary`.`commit_id`,`commit_summary`.`repo_id`,`commit_summary`.`author_user_id`,`commit_summary`.`customer_id`,`commit_summary`.`ref_type`,`commit_summary`.`additions`,`commit_summary`.`deletions`,`commit_summary`.`files_changed`,`commit_summary`.`branch`,`commit_summary`.`language`,`commit_summary`.`date`,`commit_summary`.`message`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `commit_id`=VALUES(`commit_id`),`repo_id`=VALUES(`repo_id`),`author_user_id`=VALUES(`author_user_id`),`customer_id`=VALUES(`customer_id`),`ref_type`=VALUES(`ref_type`),`additions`=VALUES(`additions`),`deletions`=VALUES(`deletions`),`files_changed`=VALUES(`files_changed`),`branch`=VALUES(`branch`),`language`=VALUES(`language`),`date`=VALUES(`date`),`message`=VALUES(`message`)"
	}
	r, err := tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.CommitID),
		orm.ToSQLString(t.RepoID),
		orm.ToSQLString(t.AuthorUserID),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefType),
		orm.ToSQLInt64(t.Additions),
		orm.ToSQLInt64(t.Deletions),
		orm.ToSQLInt64(t.FilesChanged),
		orm.ToSQLString(t.Branch),
		orm.ToSQLString(t.Language),
		orm.ToSQLInt64(t.Date),
		orm.ToSQLString(t.Message),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBFindOne will find a CommitSummary record in the database with the primary key
func (t *CommitSummary) DBFindOne(ctx context.Context, db *sql.DB, value string) (bool, error) {
	q := "SELECT `commit_summary`.`id`,`commit_summary`.`commit_id`,`commit_summary`.`repo_id`,`commit_summary`.`author_user_id`,`commit_summary`.`customer_id`,`commit_summary`.`ref_type`,`commit_summary`.`additions`,`commit_summary`.`deletions`,`commit_summary`.`files_changed`,`commit_summary`.`branch`,`commit_summary`.`language`,`commit_summary`.`date`,`commit_summary`.`message` FROM `commit_summary` WHERE `id` = ? LIMIT 1"
	row := db.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _CommitID sql.NullString
	var _RepoID sql.NullString
	var _AuthorUserID sql.NullString
	var _CustomerID sql.NullString
	var _RefType sql.NullString
	var _Additions sql.NullInt64
	var _Deletions sql.NullInt64
	var _FilesChanged sql.NullInt64
	var _Branch sql.NullString
	var _Language sql.NullString
	var _Date sql.NullInt64
	var _Message sql.NullString
	err := row.Scan(
		&_ID,
		&_CommitID,
		&_RepoID,
		&_AuthorUserID,
		&_CustomerID,
		&_RefType,
		&_Additions,
		&_Deletions,
		&_FilesChanged,
		&_Branch,
		&_Language,
		&_Date,
		&_Message,
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
	if _CommitID.Valid {
		t.SetCommitID(_CommitID.String)
	}
	if _RepoID.Valid {
		t.SetRepoID(_RepoID.String)
	}
	if _AuthorUserID.Valid {
		t.SetAuthorUserID(_AuthorUserID.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _RefType.Valid {
		t.SetRefType(_RefType.String)
	}
	if _Additions.Valid {
		t.SetAdditions(int32(_Additions.Int64))
	}
	if _Deletions.Valid {
		t.SetDeletions(int32(_Deletions.Int64))
	}
	if _FilesChanged.Valid {
		t.SetFilesChanged(int32(_FilesChanged.Int64))
	}
	if _Branch.Valid {
		t.SetBranch(_Branch.String)
	}
	if _Language.Valid {
		t.SetLanguage(_Language.String)
	}
	if _Date.Valid {
		t.SetDate(_Date.Int64)
	}
	if _Message.Valid {
		t.SetMessage(_Message.String)
	}
	return true, nil
}

// DBFindOneTx will find a CommitSummary record in the database with the primary key using the provided transaction
func (t *CommitSummary) DBFindOneTx(ctx context.Context, tx *sql.Tx, value string) (bool, error) {
	q := "SELECT `commit_summary`.`id`,`commit_summary`.`commit_id`,`commit_summary`.`repo_id`,`commit_summary`.`author_user_id`,`commit_summary`.`customer_id`,`commit_summary`.`ref_type`,`commit_summary`.`additions`,`commit_summary`.`deletions`,`commit_summary`.`files_changed`,`commit_summary`.`branch`,`commit_summary`.`language`,`commit_summary`.`date`,`commit_summary`.`message` FROM `commit_summary` WHERE `id` = ? LIMIT 1"
	row := tx.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _CommitID sql.NullString
	var _RepoID sql.NullString
	var _AuthorUserID sql.NullString
	var _CustomerID sql.NullString
	var _RefType sql.NullString
	var _Additions sql.NullInt64
	var _Deletions sql.NullInt64
	var _FilesChanged sql.NullInt64
	var _Branch sql.NullString
	var _Language sql.NullString
	var _Date sql.NullInt64
	var _Message sql.NullString
	err := row.Scan(
		&_ID,
		&_CommitID,
		&_RepoID,
		&_AuthorUserID,
		&_CustomerID,
		&_RefType,
		&_Additions,
		&_Deletions,
		&_FilesChanged,
		&_Branch,
		&_Language,
		&_Date,
		&_Message,
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
	if _CommitID.Valid {
		t.SetCommitID(_CommitID.String)
	}
	if _RepoID.Valid {
		t.SetRepoID(_RepoID.String)
	}
	if _AuthorUserID.Valid {
		t.SetAuthorUserID(_AuthorUserID.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _RefType.Valid {
		t.SetRefType(_RefType.String)
	}
	if _Additions.Valid {
		t.SetAdditions(int32(_Additions.Int64))
	}
	if _Deletions.Valid {
		t.SetDeletions(int32(_Deletions.Int64))
	}
	if _FilesChanged.Valid {
		t.SetFilesChanged(int32(_FilesChanged.Int64))
	}
	if _Branch.Valid {
		t.SetBranch(_Branch.String)
	}
	if _Language.Valid {
		t.SetLanguage(_Language.String)
	}
	if _Date.Valid {
		t.SetDate(_Date.Int64)
	}
	if _Message.Valid {
		t.SetMessage(_Message.String)
	}
	return true, nil
}

// FindCommitSummaries will find a CommitSummary record in the database with the provided parameters
func FindCommitSummaries(ctx context.Context, db *sql.DB, _params ...interface{}) ([]*CommitSummary, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("commit_id"),
		orm.Column("repo_id"),
		orm.Column("author_user_id"),
		orm.Column("customer_id"),
		orm.Column("ref_type"),
		orm.Column("additions"),
		orm.Column("deletions"),
		orm.Column("files_changed"),
		orm.Column("branch"),
		orm.Column("language"),
		orm.Column("date"),
		orm.Column("message"),
		orm.Table(CommitSummaryTableName),
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
	results := make([]*CommitSummary, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _CommitID sql.NullString
		var _RepoID sql.NullString
		var _AuthorUserID sql.NullString
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _Additions sql.NullInt64
		var _Deletions sql.NullInt64
		var _FilesChanged sql.NullInt64
		var _Branch sql.NullString
		var _Language sql.NullString
		var _Date sql.NullInt64
		var _Message sql.NullString
		err := rows.Scan(
			&_ID,
			&_CommitID,
			&_RepoID,
			&_AuthorUserID,
			&_CustomerID,
			&_RefType,
			&_Additions,
			&_Deletions,
			&_FilesChanged,
			&_Branch,
			&_Language,
			&_Date,
			&_Message,
		)
		if err != nil {
			return nil, err
		}
		t := &CommitSummary{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _CommitID.Valid {
			t.SetCommitID(_CommitID.String)
		}
		if _RepoID.Valid {
			t.SetRepoID(_RepoID.String)
		}
		if _AuthorUserID.Valid {
			t.SetAuthorUserID(_AuthorUserID.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefType.Valid {
			t.SetRefType(_RefType.String)
		}
		if _Additions.Valid {
			t.SetAdditions(int32(_Additions.Int64))
		}
		if _Deletions.Valid {
			t.SetDeletions(int32(_Deletions.Int64))
		}
		if _FilesChanged.Valid {
			t.SetFilesChanged(int32(_FilesChanged.Int64))
		}
		if _Branch.Valid {
			t.SetBranch(_Branch.String)
		}
		if _Language.Valid {
			t.SetLanguage(_Language.String)
		}
		if _Date.Valid {
			t.SetDate(_Date.Int64)
		}
		if _Message.Valid {
			t.SetMessage(_Message.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindCommitSummariesTx will find a CommitSummary record in the database with the provided parameters using the provided transaction
func FindCommitSummariesTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) ([]*CommitSummary, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("commit_id"),
		orm.Column("repo_id"),
		orm.Column("author_user_id"),
		orm.Column("customer_id"),
		orm.Column("ref_type"),
		orm.Column("additions"),
		orm.Column("deletions"),
		orm.Column("files_changed"),
		orm.Column("branch"),
		orm.Column("language"),
		orm.Column("date"),
		orm.Column("message"),
		orm.Table(CommitSummaryTableName),
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
	results := make([]*CommitSummary, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _CommitID sql.NullString
		var _RepoID sql.NullString
		var _AuthorUserID sql.NullString
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _Additions sql.NullInt64
		var _Deletions sql.NullInt64
		var _FilesChanged sql.NullInt64
		var _Branch sql.NullString
		var _Language sql.NullString
		var _Date sql.NullInt64
		var _Message sql.NullString
		err := rows.Scan(
			&_ID,
			&_CommitID,
			&_RepoID,
			&_AuthorUserID,
			&_CustomerID,
			&_RefType,
			&_Additions,
			&_Deletions,
			&_FilesChanged,
			&_Branch,
			&_Language,
			&_Date,
			&_Message,
		)
		if err != nil {
			return nil, err
		}
		t := &CommitSummary{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _CommitID.Valid {
			t.SetCommitID(_CommitID.String)
		}
		if _RepoID.Valid {
			t.SetRepoID(_RepoID.String)
		}
		if _AuthorUserID.Valid {
			t.SetAuthorUserID(_AuthorUserID.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefType.Valid {
			t.SetRefType(_RefType.String)
		}
		if _Additions.Valid {
			t.SetAdditions(int32(_Additions.Int64))
		}
		if _Deletions.Valid {
			t.SetDeletions(int32(_Deletions.Int64))
		}
		if _FilesChanged.Valid {
			t.SetFilesChanged(int32(_FilesChanged.Int64))
		}
		if _Branch.Valid {
			t.SetBranch(_Branch.String)
		}
		if _Language.Valid {
			t.SetLanguage(_Language.String)
		}
		if _Date.Valid {
			t.SetDate(_Date.Int64)
		}
		if _Message.Valid {
			t.SetMessage(_Message.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// DBFind will find a CommitSummary record in the database with the provided parameters
func (t *CommitSummary) DBFind(ctx context.Context, db *sql.DB, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("commit_id"),
		orm.Column("repo_id"),
		orm.Column("author_user_id"),
		orm.Column("customer_id"),
		orm.Column("ref_type"),
		orm.Column("additions"),
		orm.Column("deletions"),
		orm.Column("files_changed"),
		orm.Column("branch"),
		orm.Column("language"),
		orm.Column("date"),
		orm.Column("message"),
		orm.Table(CommitSummaryTableName),
	}
	if len(_params) > 0 {
		for _, param := range _params {
			params = append(params, param)
		}
	}
	q, p := orm.BuildQuery(params...)
	row := db.QueryRowContext(ctx, q, p...)
	var _ID sql.NullString
	var _CommitID sql.NullString
	var _RepoID sql.NullString
	var _AuthorUserID sql.NullString
	var _CustomerID sql.NullString
	var _RefType sql.NullString
	var _Additions sql.NullInt64
	var _Deletions sql.NullInt64
	var _FilesChanged sql.NullInt64
	var _Branch sql.NullString
	var _Language sql.NullString
	var _Date sql.NullInt64
	var _Message sql.NullString
	err := row.Scan(
		&_ID,
		&_CommitID,
		&_RepoID,
		&_AuthorUserID,
		&_CustomerID,
		&_RefType,
		&_Additions,
		&_Deletions,
		&_FilesChanged,
		&_Branch,
		&_Language,
		&_Date,
		&_Message,
	)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _CommitID.Valid {
		t.SetCommitID(_CommitID.String)
	}
	if _RepoID.Valid {
		t.SetRepoID(_RepoID.String)
	}
	if _AuthorUserID.Valid {
		t.SetAuthorUserID(_AuthorUserID.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _RefType.Valid {
		t.SetRefType(_RefType.String)
	}
	if _Additions.Valid {
		t.SetAdditions(int32(_Additions.Int64))
	}
	if _Deletions.Valid {
		t.SetDeletions(int32(_Deletions.Int64))
	}
	if _FilesChanged.Valid {
		t.SetFilesChanged(int32(_FilesChanged.Int64))
	}
	if _Branch.Valid {
		t.SetBranch(_Branch.String)
	}
	if _Language.Valid {
		t.SetLanguage(_Language.String)
	}
	if _Date.Valid {
		t.SetDate(_Date.Int64)
	}
	if _Message.Valid {
		t.SetMessage(_Message.String)
	}
	return true, nil
}

// DBFindTx will find a CommitSummary record in the database with the provided parameters using the provided transaction
func (t *CommitSummary) DBFindTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("commit_id"),
		orm.Column("repo_id"),
		orm.Column("author_user_id"),
		orm.Column("customer_id"),
		orm.Column("ref_type"),
		orm.Column("additions"),
		orm.Column("deletions"),
		orm.Column("files_changed"),
		orm.Column("branch"),
		orm.Column("language"),
		orm.Column("date"),
		orm.Column("message"),
		orm.Table(CommitSummaryTableName),
	}
	if len(_params) > 0 {
		for _, param := range _params {
			params = append(params, param)
		}
	}
	q, p := orm.BuildQuery(params...)
	row := tx.QueryRowContext(ctx, q, p...)
	var _ID sql.NullString
	var _CommitID sql.NullString
	var _RepoID sql.NullString
	var _AuthorUserID sql.NullString
	var _CustomerID sql.NullString
	var _RefType sql.NullString
	var _Additions sql.NullInt64
	var _Deletions sql.NullInt64
	var _FilesChanged sql.NullInt64
	var _Branch sql.NullString
	var _Language sql.NullString
	var _Date sql.NullInt64
	var _Message sql.NullString
	err := row.Scan(
		&_ID,
		&_CommitID,
		&_RepoID,
		&_AuthorUserID,
		&_CustomerID,
		&_RefType,
		&_Additions,
		&_Deletions,
		&_FilesChanged,
		&_Branch,
		&_Language,
		&_Date,
		&_Message,
	)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _CommitID.Valid {
		t.SetCommitID(_CommitID.String)
	}
	if _RepoID.Valid {
		t.SetRepoID(_RepoID.String)
	}
	if _AuthorUserID.Valid {
		t.SetAuthorUserID(_AuthorUserID.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _RefType.Valid {
		t.SetRefType(_RefType.String)
	}
	if _Additions.Valid {
		t.SetAdditions(int32(_Additions.Int64))
	}
	if _Deletions.Valid {
		t.SetDeletions(int32(_Deletions.Int64))
	}
	if _FilesChanged.Valid {
		t.SetFilesChanged(int32(_FilesChanged.Int64))
	}
	if _Branch.Valid {
		t.SetBranch(_Branch.String)
	}
	if _Language.Valid {
		t.SetLanguage(_Language.String)
	}
	if _Date.Valid {
		t.SetDate(_Date.Int64)
	}
	if _Message.Valid {
		t.SetMessage(_Message.String)
	}
	return true, nil
}

// CountCommitSummaries will find the count of CommitSummary records in the database
func CountCommitSummaries(ctx context.Context, db *sql.DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(CommitSummaryTableName),
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

// CountCommitSummariesTx will find the count of CommitSummary records in the database using the provided transaction
func CountCommitSummariesTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(CommitSummaryTableName),
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

// DBCount will find the count of CommitSummary records in the database
func (t *CommitSummary) DBCount(ctx context.Context, db *sql.DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(CommitSummaryTableName),
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

// DBCountTx will find the count of CommitSummary records in the database using the provided transaction
func (t *CommitSummary) DBCountTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(CommitSummaryTableName),
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

// DBExists will return true if the CommitSummary record exists in the database
func (t *CommitSummary) DBExists(ctx context.Context, db *sql.DB) (bool, error) {
	q := "SELECT `id` FROM `commit_summary` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := db.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// DBExistsTx will return true if the CommitSummary record exists in the database using the provided transaction
func (t *CommitSummary) DBExistsTx(ctx context.Context, tx *sql.Tx) (bool, error) {
	q := "SELECT `id` FROM `commit_summary` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := tx.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// PrimaryKeyColumn returns the column name for the primary key
func (t *CommitSummary) PrimaryKeyColumn() string {
	return CommitSummaryColumnID
}

// PrimaryKeyColumnType returns the primary key column Go type as a string
func (t *CommitSummary) PrimaryKeyColumnType() string {
	return "string"
}

// PrimaryKey returns the primary key column value
func (t *CommitSummary) PrimaryKey() interface{} {
	return t.ID
}
