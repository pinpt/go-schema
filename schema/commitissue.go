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

var _ Model = (*CommitIssue)(nil)
var _ CSVWriter = (*CommitIssue)(nil)
var _ JSONWriter = (*CommitIssue)(nil)
var _ Checksum = (*CommitIssue)(nil)

// CommitIssueTableName is the name of the table in SQL
const CommitIssueTableName = "commit_issue"

var CommitIssueColumns = []string{
	"id",
	"checksum",
	"commit_id",
	"branch",
	"user_id",
	"repo_id",
	"issue_id",
	"date",
	"customer_id",
	"ref_type",
	"ref_commit_id",
	"ref_repo_id",
	"ref_issue_type",
	"ref_issue_id",
	"metadata",
}

// CommitIssue table
type CommitIssue struct {
	Branch       string  `json:"branch"`
	Checksum     *string `json:"checksum,omitempty"`
	CommitID     string  `json:"commit_id"`
	CustomerID   string  `json:"customer_id"`
	Date         int64   `json:"date"`
	ID           string  `json:"id"`
	IssueID      string  `json:"issue_id"`
	Metadata     *string `json:"metadata,omitempty"`
	RefCommitID  string  `json:"ref_commit_id"`
	RefIssueID   string  `json:"ref_issue_id"`
	RefIssueType string  `json:"ref_issue_type"`
	RefRepoID    string  `json:"ref_repo_id"`
	RefType      string  `json:"ref_type"`
	RepoID       string  `json:"repo_id"`
	UserID       string  `json:"user_id"`
}

// TableName returns the SQL table name for CommitIssue and satifies the Model interface
func (t *CommitIssue) TableName() string {
	return CommitIssueTableName
}

// ToCSV will serialize the CommitIssue instance to a CSV compatible array of strings
func (t *CommitIssue) ToCSV() []string {
	return []string{
		t.ID,
		t.CalculateChecksum(),
		t.CommitID,
		t.Branch,
		t.UserID,
		t.RepoID,
		t.IssueID,
		toCSVString(t.Date),
		t.CustomerID,
		t.RefType,
		t.RefCommitID,
		t.RefRepoID,
		t.RefIssueType,
		t.RefIssueID,
		toCSVString(t.Metadata),
	}
}

// WriteCSV will serialize the CommitIssue instance to the writer as CSV and satisfies the CSVWriter interface
func (t *CommitIssue) WriteCSV(w *csv.Writer) error {
	return w.Write(t.ToCSV())
}

// WriteJSON will serialize the CommitIssue instance to the writer as JSON and satisfies the JSONWriter interface
func (t *CommitIssue) WriteJSON(w io.Writer, indent ...bool) error {
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

// NewCommitIssueReader creates a JSON reader which can read in CommitIssue objects serialized as JSON either as an array, single object or json new lines
// and writes each CommitIssue to the channel provided
func NewCommitIssueReader(r io.Reader, ch chan<- CommitIssue) error {
	return orm.Deserialize(r, func(buf json.RawMessage) error {
		dec := json.NewDecoder(bytes.NewBuffer(buf))
		e := CommitIssue{}
		if err := dec.Decode(&e); err != nil {
			return err
		}
		ch <- e
		return nil
	})
}

// NewCSVCommitIssueReaderDir will read the reader as CSV and emit each record to the channel provided
func NewCSVCommitIssueReader(r io.Reader, ch chan<- CommitIssue) error {
	cr := csv.NewReader(r)
	for {
		record, err := cr.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		ch <- CommitIssue{
			ID:           record[0],
			Checksum:     fromStringPointer(record[1]),
			CommitID:     record[2],
			Branch:       record[3],
			UserID:       record[4],
			RepoID:       record[5],
			IssueID:      record[6],
			Date:         fromCSVInt64(record[7]),
			CustomerID:   record[8],
			RefType:      record[9],
			RefCommitID:  record[10],
			RefRepoID:    record[11],
			RefIssueType: record[12],
			RefIssueID:   record[13],
			Metadata:     fromStringPointer(record[14]),
		}
	}
	return nil
}

// NewCSVCommitIssueReaderFile will read the file as a CSV and emit each record to the channel provided
func NewCSVCommitIssueReaderFile(fp string, ch chan<- CommitIssue) error {
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
	return NewCSVCommitIssueReader(fc, ch)
}

// NewCSVCommitIssueReaderDir will read the commit_issue.csv.gz file as a CSV and emit each record to the channel provided
func NewCSVCommitIssueReaderDir(dir string, ch chan<- CommitIssue) error {
	return NewCSVCommitIssueReaderFile(filepath.Join(dir, "commit_issue.csv.gz"), ch)
}

// CommitIssueCSVDeduper is a function callback which takes the existing value (a) and the new value (b)
// and the return value should be the one to use (or a new one, if applicable). return nil
// to skip processing of this record
type CommitIssueCSVDeduper func(a CommitIssue, b CommitIssue) *CommitIssue

// CommitIssueCSVDedupeDisabled is set on whether the CSV writer should de-dupe values by key
var CommitIssueCSVDedupeDisabled bool

// NewCommitIssueCSVWriterSize creates a batch writer that will write each CommitIssue into a CSV file
// this method will automatically de-duplicate entries using the primary key. if the checksum
// for a newer item with the same primary key doesn't match a previously sent item, the newer
// one will be used
func NewCommitIssueCSVWriterSize(w io.Writer, size int, dedupers ...CommitIssueCSVDeduper) (chan CommitIssue, chan bool, error) {
	cw := csv.NewWriter(w)
	ch := make(chan CommitIssue, size)
	done := make(chan bool)
	go func() {
		defer func() { done <- true }()
		dodedupe := !CommitIssueCSVDedupeDisabled
		var kv map[string]*CommitIssue
		var deduper CommitIssueCSVDeduper
		if dedupers != nil && len(dedupers) > 0 {
			deduper = dedupers[0]
			dodedupe = true
		}
		if dodedupe {
			kv = make(map[string]*CommitIssue)
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

// CommitIssueCSVDefaultSize is the default channel buffer size if not provided
var CommitIssueCSVDefaultSize = 100

// NewCommitIssueCSVWriter creates a batch writer that will write each CommitIssue into a CSV file
func NewCommitIssueCSVWriter(w io.Writer, dedupers ...CommitIssueCSVDeduper) (chan CommitIssue, chan bool, error) {
	return NewCommitIssueCSVWriterSize(w, CommitIssueCSVDefaultSize, dedupers...)
}

// NewCommitIssueCSVWriterDir creates a batch writer that will write each CommitIssue into a CSV file named commit_issue.csv.gz in dir
func NewCommitIssueCSVWriterDir(dir string, dedupers ...CommitIssueCSVDeduper) (chan CommitIssue, chan bool, error) {
	return NewCommitIssueCSVWriterFile(filepath.Join(dir, "commit_issue.csv.gz"), dedupers...)
}

// NewCommitIssueCSVWriterFile creates a batch writer that will write each CommitIssue into a CSV file
func NewCommitIssueCSVWriterFile(fn string, dedupers ...CommitIssueCSVDeduper) (chan CommitIssue, chan bool, error) {
	f, err := os.Create(fn)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening CSV file %s. %v", fn, err)
	}
	var fc io.WriteCloser = f
	if filepath.Ext(fn) == ".gz" {
		w, _ := gzip.NewWriterLevel(f, gzip.BestCompression)
		fc = w
	}
	ch, done, err := NewCommitIssueCSVWriter(fc, dedupers...)
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

type CommitIssueDBAction func(ctx context.Context, db *sql.DB, record CommitIssue) error

// NewCommitIssueDBWriterSize creates a DB writer that will write each issue into the DB
func NewCommitIssueDBWriterSize(ctx context.Context, db *sql.DB, errors chan<- error, size int, actions ...CommitIssueDBAction) (chan CommitIssue, chan bool, error) {
	ch := make(chan CommitIssue, size)
	done := make(chan bool)
	var action CommitIssueDBAction
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

// NewCommitIssueDBWriter creates a DB writer that will write each issue into the DB
func NewCommitIssueDBWriter(ctx context.Context, db *sql.DB, errors chan<- error, actions ...CommitIssueDBAction) (chan CommitIssue, chan bool, error) {
	return NewCommitIssueDBWriterSize(ctx, db, errors, 100, actions...)
}

// CommitIssueColumnID is the ID SQL column name for the CommitIssue table
const CommitIssueColumnID = "id"

// CommitIssueEscapedColumnID is the escaped ID SQL column name for the CommitIssue table
const CommitIssueEscapedColumnID = "`id`"

// CommitIssueColumnChecksum is the Checksum SQL column name for the CommitIssue table
const CommitIssueColumnChecksum = "checksum"

// CommitIssueEscapedColumnChecksum is the escaped Checksum SQL column name for the CommitIssue table
const CommitIssueEscapedColumnChecksum = "`checksum`"

// CommitIssueColumnCommitID is the CommitID SQL column name for the CommitIssue table
const CommitIssueColumnCommitID = "commit_id"

// CommitIssueEscapedColumnCommitID is the escaped CommitID SQL column name for the CommitIssue table
const CommitIssueEscapedColumnCommitID = "`commit_id`"

// CommitIssueColumnBranch is the Branch SQL column name for the CommitIssue table
const CommitIssueColumnBranch = "branch"

// CommitIssueEscapedColumnBranch is the escaped Branch SQL column name for the CommitIssue table
const CommitIssueEscapedColumnBranch = "`branch`"

// CommitIssueColumnUserID is the UserID SQL column name for the CommitIssue table
const CommitIssueColumnUserID = "user_id"

// CommitIssueEscapedColumnUserID is the escaped UserID SQL column name for the CommitIssue table
const CommitIssueEscapedColumnUserID = "`user_id`"

// CommitIssueColumnRepoID is the RepoID SQL column name for the CommitIssue table
const CommitIssueColumnRepoID = "repo_id"

// CommitIssueEscapedColumnRepoID is the escaped RepoID SQL column name for the CommitIssue table
const CommitIssueEscapedColumnRepoID = "`repo_id`"

// CommitIssueColumnIssueID is the IssueID SQL column name for the CommitIssue table
const CommitIssueColumnIssueID = "issue_id"

// CommitIssueEscapedColumnIssueID is the escaped IssueID SQL column name for the CommitIssue table
const CommitIssueEscapedColumnIssueID = "`issue_id`"

// CommitIssueColumnDate is the Date SQL column name for the CommitIssue table
const CommitIssueColumnDate = "date"

// CommitIssueEscapedColumnDate is the escaped Date SQL column name for the CommitIssue table
const CommitIssueEscapedColumnDate = "`date`"

// CommitIssueColumnCustomerID is the CustomerID SQL column name for the CommitIssue table
const CommitIssueColumnCustomerID = "customer_id"

// CommitIssueEscapedColumnCustomerID is the escaped CustomerID SQL column name for the CommitIssue table
const CommitIssueEscapedColumnCustomerID = "`customer_id`"

// CommitIssueColumnRefType is the RefType SQL column name for the CommitIssue table
const CommitIssueColumnRefType = "ref_type"

// CommitIssueEscapedColumnRefType is the escaped RefType SQL column name for the CommitIssue table
const CommitIssueEscapedColumnRefType = "`ref_type`"

// CommitIssueColumnRefCommitID is the RefCommitID SQL column name for the CommitIssue table
const CommitIssueColumnRefCommitID = "ref_commit_id"

// CommitIssueEscapedColumnRefCommitID is the escaped RefCommitID SQL column name for the CommitIssue table
const CommitIssueEscapedColumnRefCommitID = "`ref_commit_id`"

// CommitIssueColumnRefRepoID is the RefRepoID SQL column name for the CommitIssue table
const CommitIssueColumnRefRepoID = "ref_repo_id"

// CommitIssueEscapedColumnRefRepoID is the escaped RefRepoID SQL column name for the CommitIssue table
const CommitIssueEscapedColumnRefRepoID = "`ref_repo_id`"

// CommitIssueColumnRefIssueType is the RefIssueType SQL column name for the CommitIssue table
const CommitIssueColumnRefIssueType = "ref_issue_type"

// CommitIssueEscapedColumnRefIssueType is the escaped RefIssueType SQL column name for the CommitIssue table
const CommitIssueEscapedColumnRefIssueType = "`ref_issue_type`"

// CommitIssueColumnRefIssueID is the RefIssueID SQL column name for the CommitIssue table
const CommitIssueColumnRefIssueID = "ref_issue_id"

// CommitIssueEscapedColumnRefIssueID is the escaped RefIssueID SQL column name for the CommitIssue table
const CommitIssueEscapedColumnRefIssueID = "`ref_issue_id`"

// CommitIssueColumnMetadata is the Metadata SQL column name for the CommitIssue table
const CommitIssueColumnMetadata = "metadata"

// CommitIssueEscapedColumnMetadata is the escaped Metadata SQL column name for the CommitIssue table
const CommitIssueEscapedColumnMetadata = "`metadata`"

// GetID will return the CommitIssue ID value
func (t *CommitIssue) GetID() string {
	return t.ID
}

// SetID will set the CommitIssue ID value
func (t *CommitIssue) SetID(v string) {
	t.ID = v
}

// FindCommitIssueByID will find a CommitIssue by ID
func FindCommitIssueByID(ctx context.Context, db *sql.DB, value string) (*CommitIssue, error) {
	q := "SELECT `commit_issue`.`id`,`commit_issue`.`checksum`,`commit_issue`.`commit_id`,`commit_issue`.`branch`,`commit_issue`.`user_id`,`commit_issue`.`repo_id`,`commit_issue`.`issue_id`,`commit_issue`.`date`,`commit_issue`.`customer_id`,`commit_issue`.`ref_type`,`commit_issue`.`ref_commit_id`,`commit_issue`.`ref_repo_id`,`commit_issue`.`ref_issue_type`,`commit_issue`.`ref_issue_id`,`commit_issue`.`metadata` FROM `commit_issue` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CommitID sql.NullString
	var _Branch sql.NullString
	var _UserID sql.NullString
	var _RepoID sql.NullString
	var _IssueID sql.NullString
	var _Date sql.NullInt64
	var _CustomerID sql.NullString
	var _RefType sql.NullString
	var _RefCommitID sql.NullString
	var _RefRepoID sql.NullString
	var _RefIssueType sql.NullString
	var _RefIssueID sql.NullString
	var _Metadata sql.NullString
	err := db.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_CommitID,
		&_Branch,
		&_UserID,
		&_RepoID,
		&_IssueID,
		&_Date,
		&_CustomerID,
		&_RefType,
		&_RefCommitID,
		&_RefRepoID,
		&_RefIssueType,
		&_RefIssueID,
		&_Metadata,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &CommitIssue{}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _Checksum.Valid {
		t.SetChecksum(_Checksum.String)
	}
	if _CommitID.Valid {
		t.SetCommitID(_CommitID.String)
	}
	if _Branch.Valid {
		t.SetBranch(_Branch.String)
	}
	if _UserID.Valid {
		t.SetUserID(_UserID.String)
	}
	if _RepoID.Valid {
		t.SetRepoID(_RepoID.String)
	}
	if _IssueID.Valid {
		t.SetIssueID(_IssueID.String)
	}
	if _Date.Valid {
		t.SetDate(_Date.Int64)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _RefType.Valid {
		t.SetRefType(_RefType.String)
	}
	if _RefCommitID.Valid {
		t.SetRefCommitID(_RefCommitID.String)
	}
	if _RefRepoID.Valid {
		t.SetRefRepoID(_RefRepoID.String)
	}
	if _RefIssueType.Valid {
		t.SetRefIssueType(_RefIssueType.String)
	}
	if _RefIssueID.Valid {
		t.SetRefIssueID(_RefIssueID.String)
	}
	if _Metadata.Valid {
		t.SetMetadata(_Metadata.String)
	}
	return t, nil
}

// FindCommitIssueByIDTx will find a CommitIssue by ID using the provided transaction
func FindCommitIssueByIDTx(ctx context.Context, tx *sql.Tx, value string) (*CommitIssue, error) {
	q := "SELECT `commit_issue`.`id`,`commit_issue`.`checksum`,`commit_issue`.`commit_id`,`commit_issue`.`branch`,`commit_issue`.`user_id`,`commit_issue`.`repo_id`,`commit_issue`.`issue_id`,`commit_issue`.`date`,`commit_issue`.`customer_id`,`commit_issue`.`ref_type`,`commit_issue`.`ref_commit_id`,`commit_issue`.`ref_repo_id`,`commit_issue`.`ref_issue_type`,`commit_issue`.`ref_issue_id`,`commit_issue`.`metadata` FROM `commit_issue` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CommitID sql.NullString
	var _Branch sql.NullString
	var _UserID sql.NullString
	var _RepoID sql.NullString
	var _IssueID sql.NullString
	var _Date sql.NullInt64
	var _CustomerID sql.NullString
	var _RefType sql.NullString
	var _RefCommitID sql.NullString
	var _RefRepoID sql.NullString
	var _RefIssueType sql.NullString
	var _RefIssueID sql.NullString
	var _Metadata sql.NullString
	err := tx.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_CommitID,
		&_Branch,
		&_UserID,
		&_RepoID,
		&_IssueID,
		&_Date,
		&_CustomerID,
		&_RefType,
		&_RefCommitID,
		&_RefRepoID,
		&_RefIssueType,
		&_RefIssueID,
		&_Metadata,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &CommitIssue{}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _Checksum.Valid {
		t.SetChecksum(_Checksum.String)
	}
	if _CommitID.Valid {
		t.SetCommitID(_CommitID.String)
	}
	if _Branch.Valid {
		t.SetBranch(_Branch.String)
	}
	if _UserID.Valid {
		t.SetUserID(_UserID.String)
	}
	if _RepoID.Valid {
		t.SetRepoID(_RepoID.String)
	}
	if _IssueID.Valid {
		t.SetIssueID(_IssueID.String)
	}
	if _Date.Valid {
		t.SetDate(_Date.Int64)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _RefType.Valid {
		t.SetRefType(_RefType.String)
	}
	if _RefCommitID.Valid {
		t.SetRefCommitID(_RefCommitID.String)
	}
	if _RefRepoID.Valid {
		t.SetRefRepoID(_RefRepoID.String)
	}
	if _RefIssueType.Valid {
		t.SetRefIssueType(_RefIssueType.String)
	}
	if _RefIssueID.Valid {
		t.SetRefIssueID(_RefIssueID.String)
	}
	if _Metadata.Valid {
		t.SetMetadata(_Metadata.String)
	}
	return t, nil
}

// GetChecksum will return the CommitIssue Checksum value
func (t *CommitIssue) GetChecksum() string {
	if t.Checksum == nil {
		return ""
	}
	return *t.Checksum
}

// SetChecksum will set the CommitIssue Checksum value
func (t *CommitIssue) SetChecksum(v string) {
	t.Checksum = &v
}

// GetCommitID will return the CommitIssue CommitID value
func (t *CommitIssue) GetCommitID() string {
	return t.CommitID
}

// SetCommitID will set the CommitIssue CommitID value
func (t *CommitIssue) SetCommitID(v string) {
	t.CommitID = v
}

// FindCommitIssuesByCommitID will find all CommitIssues by the CommitID value
func FindCommitIssuesByCommitID(ctx context.Context, db *sql.DB, value string) ([]*CommitIssue, error) {
	q := "SELECT `commit_issue`.`id`,`commit_issue`.`checksum`,`commit_issue`.`commit_id`,`commit_issue`.`branch`,`commit_issue`.`user_id`,`commit_issue`.`repo_id`,`commit_issue`.`issue_id`,`commit_issue`.`date`,`commit_issue`.`customer_id`,`commit_issue`.`ref_type`,`commit_issue`.`ref_commit_id`,`commit_issue`.`ref_repo_id`,`commit_issue`.`ref_issue_type`,`commit_issue`.`ref_issue_id`,`commit_issue`.`metadata` FROM `commit_issue` WHERE `commit_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*CommitIssue, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CommitID sql.NullString
		var _Branch sql.NullString
		var _UserID sql.NullString
		var _RepoID sql.NullString
		var _IssueID sql.NullString
		var _Date sql.NullInt64
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefCommitID sql.NullString
		var _RefRepoID sql.NullString
		var _RefIssueType sql.NullString
		var _RefIssueID sql.NullString
		var _Metadata sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CommitID,
			&_Branch,
			&_UserID,
			&_RepoID,
			&_IssueID,
			&_Date,
			&_CustomerID,
			&_RefType,
			&_RefCommitID,
			&_RefRepoID,
			&_RefIssueType,
			&_RefIssueID,
			&_Metadata,
		)
		if err != nil {
			return nil, err
		}
		t := &CommitIssue{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CommitID.Valid {
			t.SetCommitID(_CommitID.String)
		}
		if _Branch.Valid {
			t.SetBranch(_Branch.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _RepoID.Valid {
			t.SetRepoID(_RepoID.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _Date.Valid {
			t.SetDate(_Date.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefType.Valid {
			t.SetRefType(_RefType.String)
		}
		if _RefCommitID.Valid {
			t.SetRefCommitID(_RefCommitID.String)
		}
		if _RefRepoID.Valid {
			t.SetRefRepoID(_RefRepoID.String)
		}
		if _RefIssueType.Valid {
			t.SetRefIssueType(_RefIssueType.String)
		}
		if _RefIssueID.Valid {
			t.SetRefIssueID(_RefIssueID.String)
		}
		if _Metadata.Valid {
			t.SetMetadata(_Metadata.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindCommitIssuesByCommitIDTx will find all CommitIssues by the CommitID value using the provided transaction
func FindCommitIssuesByCommitIDTx(ctx context.Context, tx *sql.Tx, value string) ([]*CommitIssue, error) {
	q := "SELECT `commit_issue`.`id`,`commit_issue`.`checksum`,`commit_issue`.`commit_id`,`commit_issue`.`branch`,`commit_issue`.`user_id`,`commit_issue`.`repo_id`,`commit_issue`.`issue_id`,`commit_issue`.`date`,`commit_issue`.`customer_id`,`commit_issue`.`ref_type`,`commit_issue`.`ref_commit_id`,`commit_issue`.`ref_repo_id`,`commit_issue`.`ref_issue_type`,`commit_issue`.`ref_issue_id`,`commit_issue`.`metadata` FROM `commit_issue` WHERE `commit_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*CommitIssue, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CommitID sql.NullString
		var _Branch sql.NullString
		var _UserID sql.NullString
		var _RepoID sql.NullString
		var _IssueID sql.NullString
		var _Date sql.NullInt64
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefCommitID sql.NullString
		var _RefRepoID sql.NullString
		var _RefIssueType sql.NullString
		var _RefIssueID sql.NullString
		var _Metadata sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CommitID,
			&_Branch,
			&_UserID,
			&_RepoID,
			&_IssueID,
			&_Date,
			&_CustomerID,
			&_RefType,
			&_RefCommitID,
			&_RefRepoID,
			&_RefIssueType,
			&_RefIssueID,
			&_Metadata,
		)
		if err != nil {
			return nil, err
		}
		t := &CommitIssue{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CommitID.Valid {
			t.SetCommitID(_CommitID.String)
		}
		if _Branch.Valid {
			t.SetBranch(_Branch.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _RepoID.Valid {
			t.SetRepoID(_RepoID.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _Date.Valid {
			t.SetDate(_Date.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefType.Valid {
			t.SetRefType(_RefType.String)
		}
		if _RefCommitID.Valid {
			t.SetRefCommitID(_RefCommitID.String)
		}
		if _RefRepoID.Valid {
			t.SetRefRepoID(_RefRepoID.String)
		}
		if _RefIssueType.Valid {
			t.SetRefIssueType(_RefIssueType.String)
		}
		if _RefIssueID.Valid {
			t.SetRefIssueID(_RefIssueID.String)
		}
		if _Metadata.Valid {
			t.SetMetadata(_Metadata.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetBranch will return the CommitIssue Branch value
func (t *CommitIssue) GetBranch() string {
	return t.Branch
}

// SetBranch will set the CommitIssue Branch value
func (t *CommitIssue) SetBranch(v string) {
	t.Branch = v
}

// GetUserID will return the CommitIssue UserID value
func (t *CommitIssue) GetUserID() string {
	return t.UserID
}

// SetUserID will set the CommitIssue UserID value
func (t *CommitIssue) SetUserID(v string) {
	t.UserID = v
}

// GetRepoID will return the CommitIssue RepoID value
func (t *CommitIssue) GetRepoID() string {
	return t.RepoID
}

// SetRepoID will set the CommitIssue RepoID value
func (t *CommitIssue) SetRepoID(v string) {
	t.RepoID = v
}

// GetIssueID will return the CommitIssue IssueID value
func (t *CommitIssue) GetIssueID() string {
	return t.IssueID
}

// SetIssueID will set the CommitIssue IssueID value
func (t *CommitIssue) SetIssueID(v string) {
	t.IssueID = v
}

// GetDate will return the CommitIssue Date value
func (t *CommitIssue) GetDate() int64 {
	return t.Date
}

// SetDate will set the CommitIssue Date value
func (t *CommitIssue) SetDate(v int64) {
	t.Date = v
}

// GetCustomerID will return the CommitIssue CustomerID value
func (t *CommitIssue) GetCustomerID() string {
	return t.CustomerID
}

// SetCustomerID will set the CommitIssue CustomerID value
func (t *CommitIssue) SetCustomerID(v string) {
	t.CustomerID = v
}

// FindCommitIssuesByCustomerID will find all CommitIssues by the CustomerID value
func FindCommitIssuesByCustomerID(ctx context.Context, db *sql.DB, value string) ([]*CommitIssue, error) {
	q := "SELECT `commit_issue`.`id`,`commit_issue`.`checksum`,`commit_issue`.`commit_id`,`commit_issue`.`branch`,`commit_issue`.`user_id`,`commit_issue`.`repo_id`,`commit_issue`.`issue_id`,`commit_issue`.`date`,`commit_issue`.`customer_id`,`commit_issue`.`ref_type`,`commit_issue`.`ref_commit_id`,`commit_issue`.`ref_repo_id`,`commit_issue`.`ref_issue_type`,`commit_issue`.`ref_issue_id`,`commit_issue`.`metadata` FROM `commit_issue` WHERE `customer_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*CommitIssue, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CommitID sql.NullString
		var _Branch sql.NullString
		var _UserID sql.NullString
		var _RepoID sql.NullString
		var _IssueID sql.NullString
		var _Date sql.NullInt64
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefCommitID sql.NullString
		var _RefRepoID sql.NullString
		var _RefIssueType sql.NullString
		var _RefIssueID sql.NullString
		var _Metadata sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CommitID,
			&_Branch,
			&_UserID,
			&_RepoID,
			&_IssueID,
			&_Date,
			&_CustomerID,
			&_RefType,
			&_RefCommitID,
			&_RefRepoID,
			&_RefIssueType,
			&_RefIssueID,
			&_Metadata,
		)
		if err != nil {
			return nil, err
		}
		t := &CommitIssue{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CommitID.Valid {
			t.SetCommitID(_CommitID.String)
		}
		if _Branch.Valid {
			t.SetBranch(_Branch.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _RepoID.Valid {
			t.SetRepoID(_RepoID.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _Date.Valid {
			t.SetDate(_Date.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefType.Valid {
			t.SetRefType(_RefType.String)
		}
		if _RefCommitID.Valid {
			t.SetRefCommitID(_RefCommitID.String)
		}
		if _RefRepoID.Valid {
			t.SetRefRepoID(_RefRepoID.String)
		}
		if _RefIssueType.Valid {
			t.SetRefIssueType(_RefIssueType.String)
		}
		if _RefIssueID.Valid {
			t.SetRefIssueID(_RefIssueID.String)
		}
		if _Metadata.Valid {
			t.SetMetadata(_Metadata.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindCommitIssuesByCustomerIDTx will find all CommitIssues by the CustomerID value using the provided transaction
func FindCommitIssuesByCustomerIDTx(ctx context.Context, tx *sql.Tx, value string) ([]*CommitIssue, error) {
	q := "SELECT `commit_issue`.`id`,`commit_issue`.`checksum`,`commit_issue`.`commit_id`,`commit_issue`.`branch`,`commit_issue`.`user_id`,`commit_issue`.`repo_id`,`commit_issue`.`issue_id`,`commit_issue`.`date`,`commit_issue`.`customer_id`,`commit_issue`.`ref_type`,`commit_issue`.`ref_commit_id`,`commit_issue`.`ref_repo_id`,`commit_issue`.`ref_issue_type`,`commit_issue`.`ref_issue_id`,`commit_issue`.`metadata` FROM `commit_issue` WHERE `customer_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*CommitIssue, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CommitID sql.NullString
		var _Branch sql.NullString
		var _UserID sql.NullString
		var _RepoID sql.NullString
		var _IssueID sql.NullString
		var _Date sql.NullInt64
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefCommitID sql.NullString
		var _RefRepoID sql.NullString
		var _RefIssueType sql.NullString
		var _RefIssueID sql.NullString
		var _Metadata sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CommitID,
			&_Branch,
			&_UserID,
			&_RepoID,
			&_IssueID,
			&_Date,
			&_CustomerID,
			&_RefType,
			&_RefCommitID,
			&_RefRepoID,
			&_RefIssueType,
			&_RefIssueID,
			&_Metadata,
		)
		if err != nil {
			return nil, err
		}
		t := &CommitIssue{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CommitID.Valid {
			t.SetCommitID(_CommitID.String)
		}
		if _Branch.Valid {
			t.SetBranch(_Branch.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _RepoID.Valid {
			t.SetRepoID(_RepoID.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _Date.Valid {
			t.SetDate(_Date.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefType.Valid {
			t.SetRefType(_RefType.String)
		}
		if _RefCommitID.Valid {
			t.SetRefCommitID(_RefCommitID.String)
		}
		if _RefRepoID.Valid {
			t.SetRefRepoID(_RefRepoID.String)
		}
		if _RefIssueType.Valid {
			t.SetRefIssueType(_RefIssueType.String)
		}
		if _RefIssueID.Valid {
			t.SetRefIssueID(_RefIssueID.String)
		}
		if _Metadata.Valid {
			t.SetMetadata(_Metadata.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetRefType will return the CommitIssue RefType value
func (t *CommitIssue) GetRefType() string {
	return t.RefType
}

// SetRefType will set the CommitIssue RefType value
func (t *CommitIssue) SetRefType(v string) {
	t.RefType = v
}

// FindCommitIssuesByRefType will find all CommitIssues by the RefType value
func FindCommitIssuesByRefType(ctx context.Context, db *sql.DB, value string) ([]*CommitIssue, error) {
	q := "SELECT `commit_issue`.`id`,`commit_issue`.`checksum`,`commit_issue`.`commit_id`,`commit_issue`.`branch`,`commit_issue`.`user_id`,`commit_issue`.`repo_id`,`commit_issue`.`issue_id`,`commit_issue`.`date`,`commit_issue`.`customer_id`,`commit_issue`.`ref_type`,`commit_issue`.`ref_commit_id`,`commit_issue`.`ref_repo_id`,`commit_issue`.`ref_issue_type`,`commit_issue`.`ref_issue_id`,`commit_issue`.`metadata` FROM `commit_issue` WHERE `ref_type` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*CommitIssue, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CommitID sql.NullString
		var _Branch sql.NullString
		var _UserID sql.NullString
		var _RepoID sql.NullString
		var _IssueID sql.NullString
		var _Date sql.NullInt64
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefCommitID sql.NullString
		var _RefRepoID sql.NullString
		var _RefIssueType sql.NullString
		var _RefIssueID sql.NullString
		var _Metadata sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CommitID,
			&_Branch,
			&_UserID,
			&_RepoID,
			&_IssueID,
			&_Date,
			&_CustomerID,
			&_RefType,
			&_RefCommitID,
			&_RefRepoID,
			&_RefIssueType,
			&_RefIssueID,
			&_Metadata,
		)
		if err != nil {
			return nil, err
		}
		t := &CommitIssue{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CommitID.Valid {
			t.SetCommitID(_CommitID.String)
		}
		if _Branch.Valid {
			t.SetBranch(_Branch.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _RepoID.Valid {
			t.SetRepoID(_RepoID.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _Date.Valid {
			t.SetDate(_Date.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefType.Valid {
			t.SetRefType(_RefType.String)
		}
		if _RefCommitID.Valid {
			t.SetRefCommitID(_RefCommitID.String)
		}
		if _RefRepoID.Valid {
			t.SetRefRepoID(_RefRepoID.String)
		}
		if _RefIssueType.Valid {
			t.SetRefIssueType(_RefIssueType.String)
		}
		if _RefIssueID.Valid {
			t.SetRefIssueID(_RefIssueID.String)
		}
		if _Metadata.Valid {
			t.SetMetadata(_Metadata.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindCommitIssuesByRefTypeTx will find all CommitIssues by the RefType value using the provided transaction
func FindCommitIssuesByRefTypeTx(ctx context.Context, tx *sql.Tx, value string) ([]*CommitIssue, error) {
	q := "SELECT `commit_issue`.`id`,`commit_issue`.`checksum`,`commit_issue`.`commit_id`,`commit_issue`.`branch`,`commit_issue`.`user_id`,`commit_issue`.`repo_id`,`commit_issue`.`issue_id`,`commit_issue`.`date`,`commit_issue`.`customer_id`,`commit_issue`.`ref_type`,`commit_issue`.`ref_commit_id`,`commit_issue`.`ref_repo_id`,`commit_issue`.`ref_issue_type`,`commit_issue`.`ref_issue_id`,`commit_issue`.`metadata` FROM `commit_issue` WHERE `ref_type` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*CommitIssue, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CommitID sql.NullString
		var _Branch sql.NullString
		var _UserID sql.NullString
		var _RepoID sql.NullString
		var _IssueID sql.NullString
		var _Date sql.NullInt64
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefCommitID sql.NullString
		var _RefRepoID sql.NullString
		var _RefIssueType sql.NullString
		var _RefIssueID sql.NullString
		var _Metadata sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CommitID,
			&_Branch,
			&_UserID,
			&_RepoID,
			&_IssueID,
			&_Date,
			&_CustomerID,
			&_RefType,
			&_RefCommitID,
			&_RefRepoID,
			&_RefIssueType,
			&_RefIssueID,
			&_Metadata,
		)
		if err != nil {
			return nil, err
		}
		t := &CommitIssue{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CommitID.Valid {
			t.SetCommitID(_CommitID.String)
		}
		if _Branch.Valid {
			t.SetBranch(_Branch.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _RepoID.Valid {
			t.SetRepoID(_RepoID.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _Date.Valid {
			t.SetDate(_Date.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefType.Valid {
			t.SetRefType(_RefType.String)
		}
		if _RefCommitID.Valid {
			t.SetRefCommitID(_RefCommitID.String)
		}
		if _RefRepoID.Valid {
			t.SetRefRepoID(_RefRepoID.String)
		}
		if _RefIssueType.Valid {
			t.SetRefIssueType(_RefIssueType.String)
		}
		if _RefIssueID.Valid {
			t.SetRefIssueID(_RefIssueID.String)
		}
		if _Metadata.Valid {
			t.SetMetadata(_Metadata.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetRefCommitID will return the CommitIssue RefCommitID value
func (t *CommitIssue) GetRefCommitID() string {
	return t.RefCommitID
}

// SetRefCommitID will set the CommitIssue RefCommitID value
func (t *CommitIssue) SetRefCommitID(v string) {
	t.RefCommitID = v
}

// FindCommitIssuesByRefCommitID will find all CommitIssues by the RefCommitID value
func FindCommitIssuesByRefCommitID(ctx context.Context, db *sql.DB, value string) ([]*CommitIssue, error) {
	q := "SELECT `commit_issue`.`id`,`commit_issue`.`checksum`,`commit_issue`.`commit_id`,`commit_issue`.`branch`,`commit_issue`.`user_id`,`commit_issue`.`repo_id`,`commit_issue`.`issue_id`,`commit_issue`.`date`,`commit_issue`.`customer_id`,`commit_issue`.`ref_type`,`commit_issue`.`ref_commit_id`,`commit_issue`.`ref_repo_id`,`commit_issue`.`ref_issue_type`,`commit_issue`.`ref_issue_id`,`commit_issue`.`metadata` FROM `commit_issue` WHERE `ref_commit_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*CommitIssue, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CommitID sql.NullString
		var _Branch sql.NullString
		var _UserID sql.NullString
		var _RepoID sql.NullString
		var _IssueID sql.NullString
		var _Date sql.NullInt64
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefCommitID sql.NullString
		var _RefRepoID sql.NullString
		var _RefIssueType sql.NullString
		var _RefIssueID sql.NullString
		var _Metadata sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CommitID,
			&_Branch,
			&_UserID,
			&_RepoID,
			&_IssueID,
			&_Date,
			&_CustomerID,
			&_RefType,
			&_RefCommitID,
			&_RefRepoID,
			&_RefIssueType,
			&_RefIssueID,
			&_Metadata,
		)
		if err != nil {
			return nil, err
		}
		t := &CommitIssue{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CommitID.Valid {
			t.SetCommitID(_CommitID.String)
		}
		if _Branch.Valid {
			t.SetBranch(_Branch.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _RepoID.Valid {
			t.SetRepoID(_RepoID.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _Date.Valid {
			t.SetDate(_Date.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefType.Valid {
			t.SetRefType(_RefType.String)
		}
		if _RefCommitID.Valid {
			t.SetRefCommitID(_RefCommitID.String)
		}
		if _RefRepoID.Valid {
			t.SetRefRepoID(_RefRepoID.String)
		}
		if _RefIssueType.Valid {
			t.SetRefIssueType(_RefIssueType.String)
		}
		if _RefIssueID.Valid {
			t.SetRefIssueID(_RefIssueID.String)
		}
		if _Metadata.Valid {
			t.SetMetadata(_Metadata.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindCommitIssuesByRefCommitIDTx will find all CommitIssues by the RefCommitID value using the provided transaction
func FindCommitIssuesByRefCommitIDTx(ctx context.Context, tx *sql.Tx, value string) ([]*CommitIssue, error) {
	q := "SELECT `commit_issue`.`id`,`commit_issue`.`checksum`,`commit_issue`.`commit_id`,`commit_issue`.`branch`,`commit_issue`.`user_id`,`commit_issue`.`repo_id`,`commit_issue`.`issue_id`,`commit_issue`.`date`,`commit_issue`.`customer_id`,`commit_issue`.`ref_type`,`commit_issue`.`ref_commit_id`,`commit_issue`.`ref_repo_id`,`commit_issue`.`ref_issue_type`,`commit_issue`.`ref_issue_id`,`commit_issue`.`metadata` FROM `commit_issue` WHERE `ref_commit_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*CommitIssue, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CommitID sql.NullString
		var _Branch sql.NullString
		var _UserID sql.NullString
		var _RepoID sql.NullString
		var _IssueID sql.NullString
		var _Date sql.NullInt64
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefCommitID sql.NullString
		var _RefRepoID sql.NullString
		var _RefIssueType sql.NullString
		var _RefIssueID sql.NullString
		var _Metadata sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CommitID,
			&_Branch,
			&_UserID,
			&_RepoID,
			&_IssueID,
			&_Date,
			&_CustomerID,
			&_RefType,
			&_RefCommitID,
			&_RefRepoID,
			&_RefIssueType,
			&_RefIssueID,
			&_Metadata,
		)
		if err != nil {
			return nil, err
		}
		t := &CommitIssue{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CommitID.Valid {
			t.SetCommitID(_CommitID.String)
		}
		if _Branch.Valid {
			t.SetBranch(_Branch.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _RepoID.Valid {
			t.SetRepoID(_RepoID.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _Date.Valid {
			t.SetDate(_Date.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefType.Valid {
			t.SetRefType(_RefType.String)
		}
		if _RefCommitID.Valid {
			t.SetRefCommitID(_RefCommitID.String)
		}
		if _RefRepoID.Valid {
			t.SetRefRepoID(_RefRepoID.String)
		}
		if _RefIssueType.Valid {
			t.SetRefIssueType(_RefIssueType.String)
		}
		if _RefIssueID.Valid {
			t.SetRefIssueID(_RefIssueID.String)
		}
		if _Metadata.Valid {
			t.SetMetadata(_Metadata.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetRefRepoID will return the CommitIssue RefRepoID value
func (t *CommitIssue) GetRefRepoID() string {
	return t.RefRepoID
}

// SetRefRepoID will set the CommitIssue RefRepoID value
func (t *CommitIssue) SetRefRepoID(v string) {
	t.RefRepoID = v
}

// FindCommitIssuesByRefRepoID will find all CommitIssues by the RefRepoID value
func FindCommitIssuesByRefRepoID(ctx context.Context, db *sql.DB, value string) ([]*CommitIssue, error) {
	q := "SELECT `commit_issue`.`id`,`commit_issue`.`checksum`,`commit_issue`.`commit_id`,`commit_issue`.`branch`,`commit_issue`.`user_id`,`commit_issue`.`repo_id`,`commit_issue`.`issue_id`,`commit_issue`.`date`,`commit_issue`.`customer_id`,`commit_issue`.`ref_type`,`commit_issue`.`ref_commit_id`,`commit_issue`.`ref_repo_id`,`commit_issue`.`ref_issue_type`,`commit_issue`.`ref_issue_id`,`commit_issue`.`metadata` FROM `commit_issue` WHERE `ref_repo_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*CommitIssue, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CommitID sql.NullString
		var _Branch sql.NullString
		var _UserID sql.NullString
		var _RepoID sql.NullString
		var _IssueID sql.NullString
		var _Date sql.NullInt64
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefCommitID sql.NullString
		var _RefRepoID sql.NullString
		var _RefIssueType sql.NullString
		var _RefIssueID sql.NullString
		var _Metadata sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CommitID,
			&_Branch,
			&_UserID,
			&_RepoID,
			&_IssueID,
			&_Date,
			&_CustomerID,
			&_RefType,
			&_RefCommitID,
			&_RefRepoID,
			&_RefIssueType,
			&_RefIssueID,
			&_Metadata,
		)
		if err != nil {
			return nil, err
		}
		t := &CommitIssue{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CommitID.Valid {
			t.SetCommitID(_CommitID.String)
		}
		if _Branch.Valid {
			t.SetBranch(_Branch.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _RepoID.Valid {
			t.SetRepoID(_RepoID.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _Date.Valid {
			t.SetDate(_Date.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefType.Valid {
			t.SetRefType(_RefType.String)
		}
		if _RefCommitID.Valid {
			t.SetRefCommitID(_RefCommitID.String)
		}
		if _RefRepoID.Valid {
			t.SetRefRepoID(_RefRepoID.String)
		}
		if _RefIssueType.Valid {
			t.SetRefIssueType(_RefIssueType.String)
		}
		if _RefIssueID.Valid {
			t.SetRefIssueID(_RefIssueID.String)
		}
		if _Metadata.Valid {
			t.SetMetadata(_Metadata.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindCommitIssuesByRefRepoIDTx will find all CommitIssues by the RefRepoID value using the provided transaction
func FindCommitIssuesByRefRepoIDTx(ctx context.Context, tx *sql.Tx, value string) ([]*CommitIssue, error) {
	q := "SELECT `commit_issue`.`id`,`commit_issue`.`checksum`,`commit_issue`.`commit_id`,`commit_issue`.`branch`,`commit_issue`.`user_id`,`commit_issue`.`repo_id`,`commit_issue`.`issue_id`,`commit_issue`.`date`,`commit_issue`.`customer_id`,`commit_issue`.`ref_type`,`commit_issue`.`ref_commit_id`,`commit_issue`.`ref_repo_id`,`commit_issue`.`ref_issue_type`,`commit_issue`.`ref_issue_id`,`commit_issue`.`metadata` FROM `commit_issue` WHERE `ref_repo_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*CommitIssue, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CommitID sql.NullString
		var _Branch sql.NullString
		var _UserID sql.NullString
		var _RepoID sql.NullString
		var _IssueID sql.NullString
		var _Date sql.NullInt64
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefCommitID sql.NullString
		var _RefRepoID sql.NullString
		var _RefIssueType sql.NullString
		var _RefIssueID sql.NullString
		var _Metadata sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CommitID,
			&_Branch,
			&_UserID,
			&_RepoID,
			&_IssueID,
			&_Date,
			&_CustomerID,
			&_RefType,
			&_RefCommitID,
			&_RefRepoID,
			&_RefIssueType,
			&_RefIssueID,
			&_Metadata,
		)
		if err != nil {
			return nil, err
		}
		t := &CommitIssue{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CommitID.Valid {
			t.SetCommitID(_CommitID.String)
		}
		if _Branch.Valid {
			t.SetBranch(_Branch.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _RepoID.Valid {
			t.SetRepoID(_RepoID.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _Date.Valid {
			t.SetDate(_Date.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefType.Valid {
			t.SetRefType(_RefType.String)
		}
		if _RefCommitID.Valid {
			t.SetRefCommitID(_RefCommitID.String)
		}
		if _RefRepoID.Valid {
			t.SetRefRepoID(_RefRepoID.String)
		}
		if _RefIssueType.Valid {
			t.SetRefIssueType(_RefIssueType.String)
		}
		if _RefIssueID.Valid {
			t.SetRefIssueID(_RefIssueID.String)
		}
		if _Metadata.Valid {
			t.SetMetadata(_Metadata.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetRefIssueType will return the CommitIssue RefIssueType value
func (t *CommitIssue) GetRefIssueType() string {
	return t.RefIssueType
}

// SetRefIssueType will set the CommitIssue RefIssueType value
func (t *CommitIssue) SetRefIssueType(v string) {
	t.RefIssueType = v
}

// FindCommitIssuesByRefIssueType will find all CommitIssues by the RefIssueType value
func FindCommitIssuesByRefIssueType(ctx context.Context, db *sql.DB, value string) ([]*CommitIssue, error) {
	q := "SELECT `commit_issue`.`id`,`commit_issue`.`checksum`,`commit_issue`.`commit_id`,`commit_issue`.`branch`,`commit_issue`.`user_id`,`commit_issue`.`repo_id`,`commit_issue`.`issue_id`,`commit_issue`.`date`,`commit_issue`.`customer_id`,`commit_issue`.`ref_type`,`commit_issue`.`ref_commit_id`,`commit_issue`.`ref_repo_id`,`commit_issue`.`ref_issue_type`,`commit_issue`.`ref_issue_id`,`commit_issue`.`metadata` FROM `commit_issue` WHERE `ref_issue_type` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*CommitIssue, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CommitID sql.NullString
		var _Branch sql.NullString
		var _UserID sql.NullString
		var _RepoID sql.NullString
		var _IssueID sql.NullString
		var _Date sql.NullInt64
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefCommitID sql.NullString
		var _RefRepoID sql.NullString
		var _RefIssueType sql.NullString
		var _RefIssueID sql.NullString
		var _Metadata sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CommitID,
			&_Branch,
			&_UserID,
			&_RepoID,
			&_IssueID,
			&_Date,
			&_CustomerID,
			&_RefType,
			&_RefCommitID,
			&_RefRepoID,
			&_RefIssueType,
			&_RefIssueID,
			&_Metadata,
		)
		if err != nil {
			return nil, err
		}
		t := &CommitIssue{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CommitID.Valid {
			t.SetCommitID(_CommitID.String)
		}
		if _Branch.Valid {
			t.SetBranch(_Branch.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _RepoID.Valid {
			t.SetRepoID(_RepoID.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _Date.Valid {
			t.SetDate(_Date.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefType.Valid {
			t.SetRefType(_RefType.String)
		}
		if _RefCommitID.Valid {
			t.SetRefCommitID(_RefCommitID.String)
		}
		if _RefRepoID.Valid {
			t.SetRefRepoID(_RefRepoID.String)
		}
		if _RefIssueType.Valid {
			t.SetRefIssueType(_RefIssueType.String)
		}
		if _RefIssueID.Valid {
			t.SetRefIssueID(_RefIssueID.String)
		}
		if _Metadata.Valid {
			t.SetMetadata(_Metadata.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindCommitIssuesByRefIssueTypeTx will find all CommitIssues by the RefIssueType value using the provided transaction
func FindCommitIssuesByRefIssueTypeTx(ctx context.Context, tx *sql.Tx, value string) ([]*CommitIssue, error) {
	q := "SELECT `commit_issue`.`id`,`commit_issue`.`checksum`,`commit_issue`.`commit_id`,`commit_issue`.`branch`,`commit_issue`.`user_id`,`commit_issue`.`repo_id`,`commit_issue`.`issue_id`,`commit_issue`.`date`,`commit_issue`.`customer_id`,`commit_issue`.`ref_type`,`commit_issue`.`ref_commit_id`,`commit_issue`.`ref_repo_id`,`commit_issue`.`ref_issue_type`,`commit_issue`.`ref_issue_id`,`commit_issue`.`metadata` FROM `commit_issue` WHERE `ref_issue_type` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*CommitIssue, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CommitID sql.NullString
		var _Branch sql.NullString
		var _UserID sql.NullString
		var _RepoID sql.NullString
		var _IssueID sql.NullString
		var _Date sql.NullInt64
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefCommitID sql.NullString
		var _RefRepoID sql.NullString
		var _RefIssueType sql.NullString
		var _RefIssueID sql.NullString
		var _Metadata sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CommitID,
			&_Branch,
			&_UserID,
			&_RepoID,
			&_IssueID,
			&_Date,
			&_CustomerID,
			&_RefType,
			&_RefCommitID,
			&_RefRepoID,
			&_RefIssueType,
			&_RefIssueID,
			&_Metadata,
		)
		if err != nil {
			return nil, err
		}
		t := &CommitIssue{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CommitID.Valid {
			t.SetCommitID(_CommitID.String)
		}
		if _Branch.Valid {
			t.SetBranch(_Branch.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _RepoID.Valid {
			t.SetRepoID(_RepoID.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _Date.Valid {
			t.SetDate(_Date.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefType.Valid {
			t.SetRefType(_RefType.String)
		}
		if _RefCommitID.Valid {
			t.SetRefCommitID(_RefCommitID.String)
		}
		if _RefRepoID.Valid {
			t.SetRefRepoID(_RefRepoID.String)
		}
		if _RefIssueType.Valid {
			t.SetRefIssueType(_RefIssueType.String)
		}
		if _RefIssueID.Valid {
			t.SetRefIssueID(_RefIssueID.String)
		}
		if _Metadata.Valid {
			t.SetMetadata(_Metadata.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetRefIssueID will return the CommitIssue RefIssueID value
func (t *CommitIssue) GetRefIssueID() string {
	return t.RefIssueID
}

// SetRefIssueID will set the CommitIssue RefIssueID value
func (t *CommitIssue) SetRefIssueID(v string) {
	t.RefIssueID = v
}

// FindCommitIssuesByRefIssueID will find all CommitIssues by the RefIssueID value
func FindCommitIssuesByRefIssueID(ctx context.Context, db *sql.DB, value string) ([]*CommitIssue, error) {
	q := "SELECT `commit_issue`.`id`,`commit_issue`.`checksum`,`commit_issue`.`commit_id`,`commit_issue`.`branch`,`commit_issue`.`user_id`,`commit_issue`.`repo_id`,`commit_issue`.`issue_id`,`commit_issue`.`date`,`commit_issue`.`customer_id`,`commit_issue`.`ref_type`,`commit_issue`.`ref_commit_id`,`commit_issue`.`ref_repo_id`,`commit_issue`.`ref_issue_type`,`commit_issue`.`ref_issue_id`,`commit_issue`.`metadata` FROM `commit_issue` WHERE `ref_issue_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*CommitIssue, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CommitID sql.NullString
		var _Branch sql.NullString
		var _UserID sql.NullString
		var _RepoID sql.NullString
		var _IssueID sql.NullString
		var _Date sql.NullInt64
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefCommitID sql.NullString
		var _RefRepoID sql.NullString
		var _RefIssueType sql.NullString
		var _RefIssueID sql.NullString
		var _Metadata sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CommitID,
			&_Branch,
			&_UserID,
			&_RepoID,
			&_IssueID,
			&_Date,
			&_CustomerID,
			&_RefType,
			&_RefCommitID,
			&_RefRepoID,
			&_RefIssueType,
			&_RefIssueID,
			&_Metadata,
		)
		if err != nil {
			return nil, err
		}
		t := &CommitIssue{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CommitID.Valid {
			t.SetCommitID(_CommitID.String)
		}
		if _Branch.Valid {
			t.SetBranch(_Branch.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _RepoID.Valid {
			t.SetRepoID(_RepoID.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _Date.Valid {
			t.SetDate(_Date.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefType.Valid {
			t.SetRefType(_RefType.String)
		}
		if _RefCommitID.Valid {
			t.SetRefCommitID(_RefCommitID.String)
		}
		if _RefRepoID.Valid {
			t.SetRefRepoID(_RefRepoID.String)
		}
		if _RefIssueType.Valid {
			t.SetRefIssueType(_RefIssueType.String)
		}
		if _RefIssueID.Valid {
			t.SetRefIssueID(_RefIssueID.String)
		}
		if _Metadata.Valid {
			t.SetMetadata(_Metadata.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindCommitIssuesByRefIssueIDTx will find all CommitIssues by the RefIssueID value using the provided transaction
func FindCommitIssuesByRefIssueIDTx(ctx context.Context, tx *sql.Tx, value string) ([]*CommitIssue, error) {
	q := "SELECT `commit_issue`.`id`,`commit_issue`.`checksum`,`commit_issue`.`commit_id`,`commit_issue`.`branch`,`commit_issue`.`user_id`,`commit_issue`.`repo_id`,`commit_issue`.`issue_id`,`commit_issue`.`date`,`commit_issue`.`customer_id`,`commit_issue`.`ref_type`,`commit_issue`.`ref_commit_id`,`commit_issue`.`ref_repo_id`,`commit_issue`.`ref_issue_type`,`commit_issue`.`ref_issue_id`,`commit_issue`.`metadata` FROM `commit_issue` WHERE `ref_issue_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*CommitIssue, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CommitID sql.NullString
		var _Branch sql.NullString
		var _UserID sql.NullString
		var _RepoID sql.NullString
		var _IssueID sql.NullString
		var _Date sql.NullInt64
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefCommitID sql.NullString
		var _RefRepoID sql.NullString
		var _RefIssueType sql.NullString
		var _RefIssueID sql.NullString
		var _Metadata sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CommitID,
			&_Branch,
			&_UserID,
			&_RepoID,
			&_IssueID,
			&_Date,
			&_CustomerID,
			&_RefType,
			&_RefCommitID,
			&_RefRepoID,
			&_RefIssueType,
			&_RefIssueID,
			&_Metadata,
		)
		if err != nil {
			return nil, err
		}
		t := &CommitIssue{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CommitID.Valid {
			t.SetCommitID(_CommitID.String)
		}
		if _Branch.Valid {
			t.SetBranch(_Branch.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _RepoID.Valid {
			t.SetRepoID(_RepoID.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _Date.Valid {
			t.SetDate(_Date.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefType.Valid {
			t.SetRefType(_RefType.String)
		}
		if _RefCommitID.Valid {
			t.SetRefCommitID(_RefCommitID.String)
		}
		if _RefRepoID.Valid {
			t.SetRefRepoID(_RefRepoID.String)
		}
		if _RefIssueType.Valid {
			t.SetRefIssueType(_RefIssueType.String)
		}
		if _RefIssueID.Valid {
			t.SetRefIssueID(_RefIssueID.String)
		}
		if _Metadata.Valid {
			t.SetMetadata(_Metadata.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetMetadata will return the CommitIssue Metadata value
func (t *CommitIssue) GetMetadata() string {
	if t.Metadata == nil {
		return ""
	}
	return *t.Metadata
}

// SetMetadata will set the CommitIssue Metadata value
func (t *CommitIssue) SetMetadata(v string) {
	t.Metadata = &v
}

func (t *CommitIssue) toTimestamp(value time.Time) *timestamp.Timestamp {
	ts, _ := ptypes.TimestampProto(value)
	return ts
}

// DBCreateCommitIssueTable will create the CommitIssue table
func DBCreateCommitIssueTable(ctx context.Context, db *sql.DB) error {
	q := "CREATE TABLE `commit_issue` (`id` VARCHAR(64) NOT NULL PRIMARY KEY,`checksum` CHAR(64),`commit_id`VARCHAR(64) NOT NULL,`branch`VARCHAR(255) NOT NULL DEFAULT \"master\",`user_id` VARCHAR(64) NOT NULL,`repo_id` VARCHAR(64) NOT NULL,`issue_id` VARCHAR(64) NOT NULL,`date` BIGINT UNSIGNED NOT NULL,`customer_id` VARCHAR(64) NOT NULL,`ref_type` VARCHAR(20) NOT NULL,`ref_commit_id` VARCHAR(64) NOT NULL,`ref_repo_id` VARCHAR(64) NOT NULL,`ref_issue_type` VARCHAR(20) NOT NULL,`ref_issue_id`VARCHAR(64) NOT NULL,`metadata` JSON,INDEX commit_issue_commit_id_index (`commit_id`),INDEX commit_issue_customer_id_index (`customer_id`),INDEX commit_issue_ref_type_index (`ref_type`),INDEX commit_issue_ref_commit_id_index (`ref_commit_id`),INDEX commit_issue_ref_repo_id_index (`ref_repo_id`),INDEX commit_issue_ref_issue_type_index (`ref_issue_type`),INDEX commit_issue_ref_issue_id_index (`ref_issue_id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBCreateCommitIssueTableTx will create the CommitIssue table using the provided transction
func DBCreateCommitIssueTableTx(ctx context.Context, tx *sql.Tx) error {
	q := "CREATE TABLE `commit_issue` (`id` VARCHAR(64) NOT NULL PRIMARY KEY,`checksum` CHAR(64),`commit_id`VARCHAR(64) NOT NULL,`branch`VARCHAR(255) NOT NULL DEFAULT \"master\",`user_id` VARCHAR(64) NOT NULL,`repo_id` VARCHAR(64) NOT NULL,`issue_id` VARCHAR(64) NOT NULL,`date` BIGINT UNSIGNED NOT NULL,`customer_id` VARCHAR(64) NOT NULL,`ref_type` VARCHAR(20) NOT NULL,`ref_commit_id` VARCHAR(64) NOT NULL,`ref_repo_id` VARCHAR(64) NOT NULL,`ref_issue_type` VARCHAR(20) NOT NULL,`ref_issue_id`VARCHAR(64) NOT NULL,`metadata` JSON,INDEX commit_issue_commit_id_index (`commit_id`),INDEX commit_issue_customer_id_index (`customer_id`),INDEX commit_issue_ref_type_index (`ref_type`),INDEX commit_issue_ref_commit_id_index (`ref_commit_id`),INDEX commit_issue_ref_repo_id_index (`ref_repo_id`),INDEX commit_issue_ref_issue_type_index (`ref_issue_type`),INDEX commit_issue_ref_issue_id_index (`ref_issue_id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// DBDropCommitIssueTable will drop the CommitIssue table
func DBDropCommitIssueTable(ctx context.Context, db *sql.DB) error {
	q := "DROP TABLE IF EXISTS `commit_issue`"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBDropCommitIssueTableTx will drop the CommitIssue table using the provided transaction
func DBDropCommitIssueTableTx(ctx context.Context, tx *sql.Tx) error {
	q := "DROP TABLE IF EXISTS `commit_issue`"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// CalculateChecksum will calculate a checksum of the SHA1 of all field values
func (t *CommitIssue) CalculateChecksum() string {
	return orm.HashStrings(
		orm.ToString(t.ID),
		orm.ToString(t.CommitID),
		orm.ToString(t.Branch),
		orm.ToString(t.UserID),
		orm.ToString(t.RepoID),
		orm.ToString(t.IssueID),
		orm.ToString(t.Date),
		orm.ToString(t.CustomerID),
		orm.ToString(t.RefType),
		orm.ToString(t.RefCommitID),
		orm.ToString(t.RefRepoID),
		orm.ToString(t.RefIssueType),
		orm.ToString(t.RefIssueID),
		orm.ToString(t.Metadata),
	)
}

// DBCreate will create a new CommitIssue record in the database
func (t *CommitIssue) DBCreate(ctx context.Context, db *sql.DB) (sql.Result, error) {
	q := "INSERT INTO `commit_issue` (`commit_issue`.`id`,`commit_issue`.`checksum`,`commit_issue`.`commit_id`,`commit_issue`.`branch`,`commit_issue`.`user_id`,`commit_issue`.`repo_id`,`commit_issue`.`issue_id`,`commit_issue`.`date`,`commit_issue`.`customer_id`,`commit_issue`.`ref_type`,`commit_issue`.`ref_commit_id`,`commit_issue`.`ref_repo_id`,`commit_issue`.`ref_issue_type`,`commit_issue`.`ref_issue_id`,`commit_issue`.`metadata`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CommitID),
		orm.ToSQLString(t.Branch),
		orm.ToSQLString(t.UserID),
		orm.ToSQLString(t.RepoID),
		orm.ToSQLString(t.IssueID),
		orm.ToSQLInt64(t.Date),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefType),
		orm.ToSQLString(t.RefCommitID),
		orm.ToSQLString(t.RefRepoID),
		orm.ToSQLString(t.RefIssueType),
		orm.ToSQLString(t.RefIssueID),
		orm.ToSQLString(t.Metadata),
	)
}

// DBCreateTx will create a new CommitIssue record in the database using the provided transaction
func (t *CommitIssue) DBCreateTx(ctx context.Context, tx *sql.Tx) (sql.Result, error) {
	q := "INSERT INTO `commit_issue` (`commit_issue`.`id`,`commit_issue`.`checksum`,`commit_issue`.`commit_id`,`commit_issue`.`branch`,`commit_issue`.`user_id`,`commit_issue`.`repo_id`,`commit_issue`.`issue_id`,`commit_issue`.`date`,`commit_issue`.`customer_id`,`commit_issue`.`ref_type`,`commit_issue`.`ref_commit_id`,`commit_issue`.`ref_repo_id`,`commit_issue`.`ref_issue_type`,`commit_issue`.`ref_issue_id`,`commit_issue`.`metadata`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CommitID),
		orm.ToSQLString(t.Branch),
		orm.ToSQLString(t.UserID),
		orm.ToSQLString(t.RepoID),
		orm.ToSQLString(t.IssueID),
		orm.ToSQLInt64(t.Date),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefType),
		orm.ToSQLString(t.RefCommitID),
		orm.ToSQLString(t.RefRepoID),
		orm.ToSQLString(t.RefIssueType),
		orm.ToSQLString(t.RefIssueID),
		orm.ToSQLString(t.Metadata),
	)
}

// DBCreateIgnoreDuplicate will upsert the CommitIssue record in the database
func (t *CommitIssue) DBCreateIgnoreDuplicate(ctx context.Context, db *sql.DB) (sql.Result, error) {
	q := "INSERT INTO `commit_issue` (`commit_issue`.`id`,`commit_issue`.`checksum`,`commit_issue`.`commit_id`,`commit_issue`.`branch`,`commit_issue`.`user_id`,`commit_issue`.`repo_id`,`commit_issue`.`issue_id`,`commit_issue`.`date`,`commit_issue`.`customer_id`,`commit_issue`.`ref_type`,`commit_issue`.`ref_commit_id`,`commit_issue`.`ref_repo_id`,`commit_issue`.`ref_issue_type`,`commit_issue`.`ref_issue_id`,`commit_issue`.`metadata`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CommitID),
		orm.ToSQLString(t.Branch),
		orm.ToSQLString(t.UserID),
		orm.ToSQLString(t.RepoID),
		orm.ToSQLString(t.IssueID),
		orm.ToSQLInt64(t.Date),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefType),
		orm.ToSQLString(t.RefCommitID),
		orm.ToSQLString(t.RefRepoID),
		orm.ToSQLString(t.RefIssueType),
		orm.ToSQLString(t.RefIssueID),
		orm.ToSQLString(t.Metadata),
	)
}

// DBCreateIgnoreDuplicateTx will upsert the CommitIssue record in the database using the provided transaction
func (t *CommitIssue) DBCreateIgnoreDuplicateTx(ctx context.Context, tx *sql.Tx) (sql.Result, error) {
	q := "INSERT INTO `commit_issue` (`commit_issue`.`id`,`commit_issue`.`checksum`,`commit_issue`.`commit_id`,`commit_issue`.`branch`,`commit_issue`.`user_id`,`commit_issue`.`repo_id`,`commit_issue`.`issue_id`,`commit_issue`.`date`,`commit_issue`.`customer_id`,`commit_issue`.`ref_type`,`commit_issue`.`ref_commit_id`,`commit_issue`.`ref_repo_id`,`commit_issue`.`ref_issue_type`,`commit_issue`.`ref_issue_id`,`commit_issue`.`metadata`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CommitID),
		orm.ToSQLString(t.Branch),
		orm.ToSQLString(t.UserID),
		orm.ToSQLString(t.RepoID),
		orm.ToSQLString(t.IssueID),
		orm.ToSQLInt64(t.Date),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefType),
		orm.ToSQLString(t.RefCommitID),
		orm.ToSQLString(t.RefRepoID),
		orm.ToSQLString(t.RefIssueType),
		orm.ToSQLString(t.RefIssueID),
		orm.ToSQLString(t.Metadata),
	)
}

// DeleteAllCommitIssues deletes all CommitIssue records in the database with optional filters
func DeleteAllCommitIssues(ctx context.Context, db *sql.DB, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(CommitIssueTableName),
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

// DeleteAllCommitIssuesTx deletes all CommitIssue records in the database with optional filters using the provided transaction
func DeleteAllCommitIssuesTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(CommitIssueTableName),
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

// DBDelete will delete this CommitIssue record in the database
func (t *CommitIssue) DBDelete(ctx context.Context, db *sql.DB) (bool, error) {
	q := "DELETE FROM `commit_issue` WHERE `id` = ?"
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

// DBDeleteTx will delete this CommitIssue record in the database using the provided transaction
func (t *CommitIssue) DBDeleteTx(ctx context.Context, tx *sql.Tx) (bool, error) {
	q := "DELETE FROM `commit_issue` WHERE `id` = ?"
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

// DBUpdate will update the CommitIssue record in the database
func (t *CommitIssue) DBUpdate(ctx context.Context, db *sql.DB) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `commit_issue` SET `checksum`=?,`commit_id`=?,`branch`=?,`user_id`=?,`repo_id`=?,`issue_id`=?,`date`=?,`customer_id`=?,`ref_type`=?,`ref_commit_id`=?,`ref_repo_id`=?,`ref_issue_type`=?,`ref_issue_id`=?,`metadata`=? WHERE `id`=?"
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CommitID),
		orm.ToSQLString(t.Branch),
		orm.ToSQLString(t.UserID),
		orm.ToSQLString(t.RepoID),
		orm.ToSQLString(t.IssueID),
		orm.ToSQLInt64(t.Date),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefType),
		orm.ToSQLString(t.RefCommitID),
		orm.ToSQLString(t.RefRepoID),
		orm.ToSQLString(t.RefIssueType),
		orm.ToSQLString(t.RefIssueID),
		orm.ToSQLString(t.Metadata),
		orm.ToSQLString(t.ID),
	)
}

// DBUpdateTx will update the CommitIssue record in the database using the provided transaction
func (t *CommitIssue) DBUpdateTx(ctx context.Context, tx *sql.Tx) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `commit_issue` SET `checksum`=?,`commit_id`=?,`branch`=?,`user_id`=?,`repo_id`=?,`issue_id`=?,`date`=?,`customer_id`=?,`ref_type`=?,`ref_commit_id`=?,`ref_repo_id`=?,`ref_issue_type`=?,`ref_issue_id`=?,`metadata`=? WHERE `id`=?"
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CommitID),
		orm.ToSQLString(t.Branch),
		orm.ToSQLString(t.UserID),
		orm.ToSQLString(t.RepoID),
		orm.ToSQLString(t.IssueID),
		orm.ToSQLInt64(t.Date),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefType),
		orm.ToSQLString(t.RefCommitID),
		orm.ToSQLString(t.RefRepoID),
		orm.ToSQLString(t.RefIssueType),
		orm.ToSQLString(t.RefIssueID),
		orm.ToSQLString(t.Metadata),
		orm.ToSQLString(t.ID),
	)
}

// DBUpsert will upsert the CommitIssue record in the database
func (t *CommitIssue) DBUpsert(ctx context.Context, db *sql.DB, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `commit_issue` (`commit_issue`.`id`,`commit_issue`.`checksum`,`commit_issue`.`commit_id`,`commit_issue`.`branch`,`commit_issue`.`user_id`,`commit_issue`.`repo_id`,`commit_issue`.`issue_id`,`commit_issue`.`date`,`commit_issue`.`customer_id`,`commit_issue`.`ref_type`,`commit_issue`.`ref_commit_id`,`commit_issue`.`ref_repo_id`,`commit_issue`.`ref_issue_type`,`commit_issue`.`ref_issue_id`,`commit_issue`.`metadata`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `commit_issue` (`commit_issue`.`id`,`commit_issue`.`checksum`,`commit_issue`.`commit_id`,`commit_issue`.`branch`,`commit_issue`.`user_id`,`commit_issue`.`repo_id`,`commit_issue`.`issue_id`,`commit_issue`.`date`,`commit_issue`.`customer_id`,`commit_issue`.`ref_type`,`commit_issue`.`ref_commit_id`,`commit_issue`.`ref_repo_id`,`commit_issue`.`ref_issue_type`,`commit_issue`.`ref_issue_id`,`commit_issue`.`metadata`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`commit_id`=VALUES(`commit_id`),`branch`=VALUES(`branch`),`user_id`=VALUES(`user_id`),`repo_id`=VALUES(`repo_id`),`issue_id`=VALUES(`issue_id`),`date`=VALUES(`date`),`customer_id`=VALUES(`customer_id`),`ref_type`=VALUES(`ref_type`),`ref_commit_id`=VALUES(`ref_commit_id`),`ref_repo_id`=VALUES(`ref_repo_id`),`ref_issue_type`=VALUES(`ref_issue_type`),`ref_issue_id`=VALUES(`ref_issue_id`),`metadata`=VALUES(`metadata`)"
	}
	r, err := db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CommitID),
		orm.ToSQLString(t.Branch),
		orm.ToSQLString(t.UserID),
		orm.ToSQLString(t.RepoID),
		orm.ToSQLString(t.IssueID),
		orm.ToSQLInt64(t.Date),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefType),
		orm.ToSQLString(t.RefCommitID),
		orm.ToSQLString(t.RefRepoID),
		orm.ToSQLString(t.RefIssueType),
		orm.ToSQLString(t.RefIssueID),
		orm.ToSQLString(t.Metadata),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBUpsertTx will upsert the CommitIssue record in the database using the provided transaction
func (t *CommitIssue) DBUpsertTx(ctx context.Context, tx *sql.Tx, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `commit_issue` (`commit_issue`.`id`,`commit_issue`.`checksum`,`commit_issue`.`commit_id`,`commit_issue`.`branch`,`commit_issue`.`user_id`,`commit_issue`.`repo_id`,`commit_issue`.`issue_id`,`commit_issue`.`date`,`commit_issue`.`customer_id`,`commit_issue`.`ref_type`,`commit_issue`.`ref_commit_id`,`commit_issue`.`ref_repo_id`,`commit_issue`.`ref_issue_type`,`commit_issue`.`ref_issue_id`,`commit_issue`.`metadata`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `commit_issue` (`commit_issue`.`id`,`commit_issue`.`checksum`,`commit_issue`.`commit_id`,`commit_issue`.`branch`,`commit_issue`.`user_id`,`commit_issue`.`repo_id`,`commit_issue`.`issue_id`,`commit_issue`.`date`,`commit_issue`.`customer_id`,`commit_issue`.`ref_type`,`commit_issue`.`ref_commit_id`,`commit_issue`.`ref_repo_id`,`commit_issue`.`ref_issue_type`,`commit_issue`.`ref_issue_id`,`commit_issue`.`metadata`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`commit_id`=VALUES(`commit_id`),`branch`=VALUES(`branch`),`user_id`=VALUES(`user_id`),`repo_id`=VALUES(`repo_id`),`issue_id`=VALUES(`issue_id`),`date`=VALUES(`date`),`customer_id`=VALUES(`customer_id`),`ref_type`=VALUES(`ref_type`),`ref_commit_id`=VALUES(`ref_commit_id`),`ref_repo_id`=VALUES(`ref_repo_id`),`ref_issue_type`=VALUES(`ref_issue_type`),`ref_issue_id`=VALUES(`ref_issue_id`),`metadata`=VALUES(`metadata`)"
	}
	r, err := tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CommitID),
		orm.ToSQLString(t.Branch),
		orm.ToSQLString(t.UserID),
		orm.ToSQLString(t.RepoID),
		orm.ToSQLString(t.IssueID),
		orm.ToSQLInt64(t.Date),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefType),
		orm.ToSQLString(t.RefCommitID),
		orm.ToSQLString(t.RefRepoID),
		orm.ToSQLString(t.RefIssueType),
		orm.ToSQLString(t.RefIssueID),
		orm.ToSQLString(t.Metadata),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBFindOne will find a CommitIssue record in the database with the primary key
func (t *CommitIssue) DBFindOne(ctx context.Context, db *sql.DB, value string) (bool, error) {
	q := "SELECT `commit_issue`.`id`,`commit_issue`.`checksum`,`commit_issue`.`commit_id`,`commit_issue`.`branch`,`commit_issue`.`user_id`,`commit_issue`.`repo_id`,`commit_issue`.`issue_id`,`commit_issue`.`date`,`commit_issue`.`customer_id`,`commit_issue`.`ref_type`,`commit_issue`.`ref_commit_id`,`commit_issue`.`ref_repo_id`,`commit_issue`.`ref_issue_type`,`commit_issue`.`ref_issue_id`,`commit_issue`.`metadata` FROM `commit_issue` WHERE `id` = ? LIMIT 1"
	row := db.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CommitID sql.NullString
	var _Branch sql.NullString
	var _UserID sql.NullString
	var _RepoID sql.NullString
	var _IssueID sql.NullString
	var _Date sql.NullInt64
	var _CustomerID sql.NullString
	var _RefType sql.NullString
	var _RefCommitID sql.NullString
	var _RefRepoID sql.NullString
	var _RefIssueType sql.NullString
	var _RefIssueID sql.NullString
	var _Metadata sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CommitID,
		&_Branch,
		&_UserID,
		&_RepoID,
		&_IssueID,
		&_Date,
		&_CustomerID,
		&_RefType,
		&_RefCommitID,
		&_RefRepoID,
		&_RefIssueType,
		&_RefIssueID,
		&_Metadata,
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
	if _CommitID.Valid {
		t.SetCommitID(_CommitID.String)
	}
	if _Branch.Valid {
		t.SetBranch(_Branch.String)
	}
	if _UserID.Valid {
		t.SetUserID(_UserID.String)
	}
	if _RepoID.Valid {
		t.SetRepoID(_RepoID.String)
	}
	if _IssueID.Valid {
		t.SetIssueID(_IssueID.String)
	}
	if _Date.Valid {
		t.SetDate(_Date.Int64)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _RefType.Valid {
		t.SetRefType(_RefType.String)
	}
	if _RefCommitID.Valid {
		t.SetRefCommitID(_RefCommitID.String)
	}
	if _RefRepoID.Valid {
		t.SetRefRepoID(_RefRepoID.String)
	}
	if _RefIssueType.Valid {
		t.SetRefIssueType(_RefIssueType.String)
	}
	if _RefIssueID.Valid {
		t.SetRefIssueID(_RefIssueID.String)
	}
	if _Metadata.Valid {
		t.SetMetadata(_Metadata.String)
	}
	return true, nil
}

// DBFindOneTx will find a CommitIssue record in the database with the primary key using the provided transaction
func (t *CommitIssue) DBFindOneTx(ctx context.Context, tx *sql.Tx, value string) (bool, error) {
	q := "SELECT `commit_issue`.`id`,`commit_issue`.`checksum`,`commit_issue`.`commit_id`,`commit_issue`.`branch`,`commit_issue`.`user_id`,`commit_issue`.`repo_id`,`commit_issue`.`issue_id`,`commit_issue`.`date`,`commit_issue`.`customer_id`,`commit_issue`.`ref_type`,`commit_issue`.`ref_commit_id`,`commit_issue`.`ref_repo_id`,`commit_issue`.`ref_issue_type`,`commit_issue`.`ref_issue_id`,`commit_issue`.`metadata` FROM `commit_issue` WHERE `id` = ? LIMIT 1"
	row := tx.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CommitID sql.NullString
	var _Branch sql.NullString
	var _UserID sql.NullString
	var _RepoID sql.NullString
	var _IssueID sql.NullString
	var _Date sql.NullInt64
	var _CustomerID sql.NullString
	var _RefType sql.NullString
	var _RefCommitID sql.NullString
	var _RefRepoID sql.NullString
	var _RefIssueType sql.NullString
	var _RefIssueID sql.NullString
	var _Metadata sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CommitID,
		&_Branch,
		&_UserID,
		&_RepoID,
		&_IssueID,
		&_Date,
		&_CustomerID,
		&_RefType,
		&_RefCommitID,
		&_RefRepoID,
		&_RefIssueType,
		&_RefIssueID,
		&_Metadata,
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
	if _CommitID.Valid {
		t.SetCommitID(_CommitID.String)
	}
	if _Branch.Valid {
		t.SetBranch(_Branch.String)
	}
	if _UserID.Valid {
		t.SetUserID(_UserID.String)
	}
	if _RepoID.Valid {
		t.SetRepoID(_RepoID.String)
	}
	if _IssueID.Valid {
		t.SetIssueID(_IssueID.String)
	}
	if _Date.Valid {
		t.SetDate(_Date.Int64)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _RefType.Valid {
		t.SetRefType(_RefType.String)
	}
	if _RefCommitID.Valid {
		t.SetRefCommitID(_RefCommitID.String)
	}
	if _RefRepoID.Valid {
		t.SetRefRepoID(_RefRepoID.String)
	}
	if _RefIssueType.Valid {
		t.SetRefIssueType(_RefIssueType.String)
	}
	if _RefIssueID.Valid {
		t.SetRefIssueID(_RefIssueID.String)
	}
	if _Metadata.Valid {
		t.SetMetadata(_Metadata.String)
	}
	return true, nil
}

// FindCommitIssues will find a CommitIssue record in the database with the provided parameters
func FindCommitIssues(ctx context.Context, db *sql.DB, _params ...interface{}) ([]*CommitIssue, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("commit_id"),
		orm.Column("branch"),
		orm.Column("user_id"),
		orm.Column("repo_id"),
		orm.Column("issue_id"),
		orm.Column("date"),
		orm.Column("customer_id"),
		orm.Column("ref_type"),
		orm.Column("ref_commit_id"),
		orm.Column("ref_repo_id"),
		orm.Column("ref_issue_type"),
		orm.Column("ref_issue_id"),
		orm.Column("metadata"),
		orm.Table(CommitIssueTableName),
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
	results := make([]*CommitIssue, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CommitID sql.NullString
		var _Branch sql.NullString
		var _UserID sql.NullString
		var _RepoID sql.NullString
		var _IssueID sql.NullString
		var _Date sql.NullInt64
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefCommitID sql.NullString
		var _RefRepoID sql.NullString
		var _RefIssueType sql.NullString
		var _RefIssueID sql.NullString
		var _Metadata sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CommitID,
			&_Branch,
			&_UserID,
			&_RepoID,
			&_IssueID,
			&_Date,
			&_CustomerID,
			&_RefType,
			&_RefCommitID,
			&_RefRepoID,
			&_RefIssueType,
			&_RefIssueID,
			&_Metadata,
		)
		if err != nil {
			return nil, err
		}
		t := &CommitIssue{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CommitID.Valid {
			t.SetCommitID(_CommitID.String)
		}
		if _Branch.Valid {
			t.SetBranch(_Branch.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _RepoID.Valid {
			t.SetRepoID(_RepoID.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _Date.Valid {
			t.SetDate(_Date.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefType.Valid {
			t.SetRefType(_RefType.String)
		}
		if _RefCommitID.Valid {
			t.SetRefCommitID(_RefCommitID.String)
		}
		if _RefRepoID.Valid {
			t.SetRefRepoID(_RefRepoID.String)
		}
		if _RefIssueType.Valid {
			t.SetRefIssueType(_RefIssueType.String)
		}
		if _RefIssueID.Valid {
			t.SetRefIssueID(_RefIssueID.String)
		}
		if _Metadata.Valid {
			t.SetMetadata(_Metadata.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindCommitIssuesTx will find a CommitIssue record in the database with the provided parameters using the provided transaction
func FindCommitIssuesTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) ([]*CommitIssue, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("commit_id"),
		orm.Column("branch"),
		orm.Column("user_id"),
		orm.Column("repo_id"),
		orm.Column("issue_id"),
		orm.Column("date"),
		orm.Column("customer_id"),
		orm.Column("ref_type"),
		orm.Column("ref_commit_id"),
		orm.Column("ref_repo_id"),
		orm.Column("ref_issue_type"),
		orm.Column("ref_issue_id"),
		orm.Column("metadata"),
		orm.Table(CommitIssueTableName),
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
	results := make([]*CommitIssue, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CommitID sql.NullString
		var _Branch sql.NullString
		var _UserID sql.NullString
		var _RepoID sql.NullString
		var _IssueID sql.NullString
		var _Date sql.NullInt64
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefCommitID sql.NullString
		var _RefRepoID sql.NullString
		var _RefIssueType sql.NullString
		var _RefIssueID sql.NullString
		var _Metadata sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CommitID,
			&_Branch,
			&_UserID,
			&_RepoID,
			&_IssueID,
			&_Date,
			&_CustomerID,
			&_RefType,
			&_RefCommitID,
			&_RefRepoID,
			&_RefIssueType,
			&_RefIssueID,
			&_Metadata,
		)
		if err != nil {
			return nil, err
		}
		t := &CommitIssue{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CommitID.Valid {
			t.SetCommitID(_CommitID.String)
		}
		if _Branch.Valid {
			t.SetBranch(_Branch.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _RepoID.Valid {
			t.SetRepoID(_RepoID.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _Date.Valid {
			t.SetDate(_Date.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefType.Valid {
			t.SetRefType(_RefType.String)
		}
		if _RefCommitID.Valid {
			t.SetRefCommitID(_RefCommitID.String)
		}
		if _RefRepoID.Valid {
			t.SetRefRepoID(_RefRepoID.String)
		}
		if _RefIssueType.Valid {
			t.SetRefIssueType(_RefIssueType.String)
		}
		if _RefIssueID.Valid {
			t.SetRefIssueID(_RefIssueID.String)
		}
		if _Metadata.Valid {
			t.SetMetadata(_Metadata.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// DBFind will find a CommitIssue record in the database with the provided parameters
func (t *CommitIssue) DBFind(ctx context.Context, db *sql.DB, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("commit_id"),
		orm.Column("branch"),
		orm.Column("user_id"),
		orm.Column("repo_id"),
		orm.Column("issue_id"),
		orm.Column("date"),
		orm.Column("customer_id"),
		orm.Column("ref_type"),
		orm.Column("ref_commit_id"),
		orm.Column("ref_repo_id"),
		orm.Column("ref_issue_type"),
		orm.Column("ref_issue_id"),
		orm.Column("metadata"),
		orm.Table(CommitIssueTableName),
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
	var _CommitID sql.NullString
	var _Branch sql.NullString
	var _UserID sql.NullString
	var _RepoID sql.NullString
	var _IssueID sql.NullString
	var _Date sql.NullInt64
	var _CustomerID sql.NullString
	var _RefType sql.NullString
	var _RefCommitID sql.NullString
	var _RefRepoID sql.NullString
	var _RefIssueType sql.NullString
	var _RefIssueID sql.NullString
	var _Metadata sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CommitID,
		&_Branch,
		&_UserID,
		&_RepoID,
		&_IssueID,
		&_Date,
		&_CustomerID,
		&_RefType,
		&_RefCommitID,
		&_RefRepoID,
		&_RefIssueType,
		&_RefIssueID,
		&_Metadata,
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
	if _CommitID.Valid {
		t.SetCommitID(_CommitID.String)
	}
	if _Branch.Valid {
		t.SetBranch(_Branch.String)
	}
	if _UserID.Valid {
		t.SetUserID(_UserID.String)
	}
	if _RepoID.Valid {
		t.SetRepoID(_RepoID.String)
	}
	if _IssueID.Valid {
		t.SetIssueID(_IssueID.String)
	}
	if _Date.Valid {
		t.SetDate(_Date.Int64)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _RefType.Valid {
		t.SetRefType(_RefType.String)
	}
	if _RefCommitID.Valid {
		t.SetRefCommitID(_RefCommitID.String)
	}
	if _RefRepoID.Valid {
		t.SetRefRepoID(_RefRepoID.String)
	}
	if _RefIssueType.Valid {
		t.SetRefIssueType(_RefIssueType.String)
	}
	if _RefIssueID.Valid {
		t.SetRefIssueID(_RefIssueID.String)
	}
	if _Metadata.Valid {
		t.SetMetadata(_Metadata.String)
	}
	return true, nil
}

// DBFindTx will find a CommitIssue record in the database with the provided parameters using the provided transaction
func (t *CommitIssue) DBFindTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("commit_id"),
		orm.Column("branch"),
		orm.Column("user_id"),
		orm.Column("repo_id"),
		orm.Column("issue_id"),
		orm.Column("date"),
		orm.Column("customer_id"),
		orm.Column("ref_type"),
		orm.Column("ref_commit_id"),
		orm.Column("ref_repo_id"),
		orm.Column("ref_issue_type"),
		orm.Column("ref_issue_id"),
		orm.Column("metadata"),
		orm.Table(CommitIssueTableName),
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
	var _CommitID sql.NullString
	var _Branch sql.NullString
	var _UserID sql.NullString
	var _RepoID sql.NullString
	var _IssueID sql.NullString
	var _Date sql.NullInt64
	var _CustomerID sql.NullString
	var _RefType sql.NullString
	var _RefCommitID sql.NullString
	var _RefRepoID sql.NullString
	var _RefIssueType sql.NullString
	var _RefIssueID sql.NullString
	var _Metadata sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CommitID,
		&_Branch,
		&_UserID,
		&_RepoID,
		&_IssueID,
		&_Date,
		&_CustomerID,
		&_RefType,
		&_RefCommitID,
		&_RefRepoID,
		&_RefIssueType,
		&_RefIssueID,
		&_Metadata,
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
	if _CommitID.Valid {
		t.SetCommitID(_CommitID.String)
	}
	if _Branch.Valid {
		t.SetBranch(_Branch.String)
	}
	if _UserID.Valid {
		t.SetUserID(_UserID.String)
	}
	if _RepoID.Valid {
		t.SetRepoID(_RepoID.String)
	}
	if _IssueID.Valid {
		t.SetIssueID(_IssueID.String)
	}
	if _Date.Valid {
		t.SetDate(_Date.Int64)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _RefType.Valid {
		t.SetRefType(_RefType.String)
	}
	if _RefCommitID.Valid {
		t.SetRefCommitID(_RefCommitID.String)
	}
	if _RefRepoID.Valid {
		t.SetRefRepoID(_RefRepoID.String)
	}
	if _RefIssueType.Valid {
		t.SetRefIssueType(_RefIssueType.String)
	}
	if _RefIssueID.Valid {
		t.SetRefIssueID(_RefIssueID.String)
	}
	if _Metadata.Valid {
		t.SetMetadata(_Metadata.String)
	}
	return true, nil
}

// CountCommitIssues will find the count of CommitIssue records in the database
func CountCommitIssues(ctx context.Context, db *sql.DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(CommitIssueTableName),
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

// CountCommitIssuesTx will find the count of CommitIssue records in the database using the provided transaction
func CountCommitIssuesTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(CommitIssueTableName),
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

// DBCount will find the count of CommitIssue records in the database
func (t *CommitIssue) DBCount(ctx context.Context, db *sql.DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(CommitIssueTableName),
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

// DBCountTx will find the count of CommitIssue records in the database using the provided transaction
func (t *CommitIssue) DBCountTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(CommitIssueTableName),
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

// DBExists will return true if the CommitIssue record exists in the database
func (t *CommitIssue) DBExists(ctx context.Context, db *sql.DB) (bool, error) {
	q := "SELECT `id` FROM `commit_issue` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := db.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// DBExistsTx will return true if the CommitIssue record exists in the database using the provided transaction
func (t *CommitIssue) DBExistsTx(ctx context.Context, tx *sql.Tx) (bool, error) {
	q := "SELECT `id` FROM `commit_issue` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := tx.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// PrimaryKeyColumn returns the column name for the primary key
func (t *CommitIssue) PrimaryKeyColumn() string {
	return CommitIssueColumnID
}

// PrimaryKeyColumnType returns the primary key column Go type as a string
func (t *CommitIssue) PrimaryKeyColumnType() string {
	return "string"
}

// PrimaryKey returns the primary key column value
func (t *CommitIssue) PrimaryKey() interface{} {
	return t.ID
}
