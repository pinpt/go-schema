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

var _ Model = (*Commit)(nil)
var _ CSVWriter = (*Commit)(nil)
var _ JSONWriter = (*Commit)(nil)
var _ Checksum = (*Commit)(nil)

// CommitTableName is the name of the table in SQL
const CommitTableName = "commit"

var CommitColumns = []string{
	"id",
	"checksum",
	"repo_id",
	"sha",
	"branch",
	"message",
	"mergecommit",
	"excluded",
	"parent",
	"parent_id",
	"date",
	"author_user_id",
	"committer_user_id",
	"ordinal",
	"customer_id",
	"ref_type",
	"ref_id",
	"metadata",
}

// Commit table
type Commit struct {
	AuthorUserID    string  `json:"author_user_id"`
	Branch          *string `json:"branch,omitempty"`
	Checksum        *string `json:"checksum,omitempty"`
	CommitterUserID string  `json:"committer_user_id"`
	CustomerID      string  `json:"customer_id"`
	Date            int64   `json:"date"`
	Excluded        *bool   `json:"excluded,omitempty"`
	ID              string  `json:"id"`
	Mergecommit     *bool   `json:"mergecommit,omitempty"`
	Message         *string `json:"message,omitempty"`
	Metadata        *string `json:"metadata,omitempty"`
	Ordinal         *int32  `json:"ordinal,omitempty"`
	Parent          *string `json:"parent,omitempty"`
	ParentID        *string `json:"parent_id,omitempty"`
	RefID           string  `json:"ref_id"`
	RefType         string  `json:"ref_type"`
	RepoID          string  `json:"repo_id"`
	Sha             string  `json:"sha"`
}

// TableName returns the SQL table name for Commit and satifies the Model interface
func (t *Commit) TableName() string {
	return CommitTableName
}

// ToCSV will serialize the Commit instance to a CSV compatible array of strings
func (t *Commit) ToCSV() []string {
	return []string{
		t.ID,
		t.CalculateChecksum(),
		t.RepoID,
		t.Sha,
		toCSVString(t.Branch),
		toCSVString(t.Message),
		toCSVString(t.Mergecommit),
		toCSVString(t.Excluded),
		toCSVString(t.Parent),
		toCSVString(t.ParentID),
		toCSVString(t.Date),
		t.AuthorUserID,
		t.CommitterUserID,
		toCSVString(t.Ordinal),
		t.CustomerID,
		t.RefType,
		t.RefID,
		toCSVString(t.Metadata),
	}
}

// WriteCSV will serialize the Commit instance to the writer as CSV and satisfies the CSVWriter interface
func (t *Commit) WriteCSV(w *csv.Writer) error {
	return w.Write(t.ToCSV())
}

// WriteJSON will serialize the Commit instance to the writer as JSON and satisfies the JSONWriter interface
func (t *Commit) WriteJSON(w io.Writer, indent ...bool) error {
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

// NewCommitReader creates a JSON reader which can read in Commit objects serialized as JSON either as an array, single object or json new lines
// and writes each Commit to the channel provided
func NewCommitReader(r io.Reader, ch chan<- Commit) error {
	return orm.Deserialize(r, func(buf json.RawMessage) error {
		dec := json.NewDecoder(bytes.NewBuffer(buf))
		e := Commit{}
		if err := dec.Decode(&e); err != nil {
			return err
		}
		ch <- e
		return nil
	})
}

// NewCSVCommitReaderDir will read the reader as CSV and emit each record to the channel provided
func NewCSVCommitReader(r io.Reader, ch chan<- Commit) error {
	cr := csv.NewReader(r)
	for {
		record, err := cr.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		ch <- Commit{
			ID:              record[0],
			Checksum:        fromStringPointer(record[1]),
			RepoID:          record[2],
			Sha:             record[3],
			Branch:          fromStringPointer(record[4]),
			Message:         fromStringPointer(record[5]),
			Mergecommit:     fromCSVBoolPointer(record[6]),
			Excluded:        fromCSVBoolPointer(record[7]),
			Parent:          fromStringPointer(record[8]),
			ParentID:        fromStringPointer(record[9]),
			Date:            fromCSVInt64(record[10]),
			AuthorUserID:    record[11],
			CommitterUserID: record[12],
			Ordinal:         fromCSVInt32Pointer(record[13]),
			CustomerID:      record[14],
			RefType:         record[15],
			RefID:           record[16],
			Metadata:        fromStringPointer(record[17]),
		}
	}
	return nil
}

// NewCSVCommitReaderFile will read the file as a CSV and emit each record to the channel provided
func NewCSVCommitReaderFile(fp string, ch chan<- Commit) error {
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
	return NewCSVCommitReader(fc, ch)
}

// NewCSVCommitReaderDir will read the commit.csv.gz file as a CSV and emit each record to the channel provided
func NewCSVCommitReaderDir(dir string, ch chan<- Commit) error {
	return NewCSVCommitReaderFile(filepath.Join(dir, "commit.csv.gz"), ch)
}

// CommitCSVDeduper is a function callback which takes the existing value (a) and the new value (b)
// and the return value should be the one to use (or a new one, if applicable). return nil
// to skip processing of this record
type CommitCSVDeduper func(a Commit, b Commit) *Commit

// CommitCSVDedupeDisabled is set on whether the CSV writer should de-dupe values by key
var CommitCSVDedupeDisabled bool

// NewCommitCSVWriterSize creates a batch writer that will write each Commit into a CSV file
// this method will automatically de-duplicate entries using the primary key. if the checksum
// for a newer item with the same primary key doesn't match a previously sent item, the newer
// one will be used
func NewCommitCSVWriterSize(w io.Writer, size int, dedupers ...CommitCSVDeduper) (chan Commit, chan bool, error) {
	cw := csv.NewWriter(w)
	ch := make(chan Commit, size)
	done := make(chan bool)
	go func() {
		defer func() { done <- true }()
		dodedupe := !CommitCSVDedupeDisabled
		var kv map[string]*Commit
		var deduper CommitCSVDeduper
		if dedupers != nil && len(dedupers) > 0 {
			deduper = dedupers[0]
			dodedupe = true
		}
		if dodedupe {
			kv = make(map[string]*Commit)
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

// CommitCSVDefaultSize is the default channel buffer size if not provided
var CommitCSVDefaultSize = 100

// NewCommitCSVWriter creates a batch writer that will write each Commit into a CSV file
func NewCommitCSVWriter(w io.Writer, dedupers ...CommitCSVDeduper) (chan Commit, chan bool, error) {
	return NewCommitCSVWriterSize(w, CommitCSVDefaultSize, dedupers...)
}

// NewCommitCSVWriterDir creates a batch writer that will write each Commit into a CSV file named commit.csv.gz in dir
func NewCommitCSVWriterDir(dir string, dedupers ...CommitCSVDeduper) (chan Commit, chan bool, error) {
	return NewCommitCSVWriterFile(filepath.Join(dir, "commit.csv.gz"), dedupers...)
}

// NewCommitCSVWriterFile creates a batch writer that will write each Commit into a CSV file
func NewCommitCSVWriterFile(fn string, dedupers ...CommitCSVDeduper) (chan Commit, chan bool, error) {
	f, err := os.Create(fn)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening CSV file %s. %v", fn, err)
	}
	var fc io.WriteCloser = f
	if filepath.Ext(fn) == ".gz" {
		w, _ := gzip.NewWriterLevel(f, gzip.BestCompression)
		fc = w
	}
	ch, done, err := NewCommitCSVWriter(fc, dedupers...)
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

type CommitDBAction func(ctx context.Context, db DB, record Commit) error

// NewCommitDBWriterSize creates a DB writer that will write each issue into the DB
func NewCommitDBWriterSize(ctx context.Context, db DB, errors chan<- error, size int, actions ...CommitDBAction) (chan Commit, chan bool, error) {
	ch := make(chan Commit, size)
	done := make(chan bool)
	var action CommitDBAction
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

// NewCommitDBWriter creates a DB writer that will write each issue into the DB
func NewCommitDBWriter(ctx context.Context, db DB, errors chan<- error, actions ...CommitDBAction) (chan Commit, chan bool, error) {
	return NewCommitDBWriterSize(ctx, db, errors, 100, actions...)
}

// CommitColumnID is the ID SQL column name for the Commit table
const CommitColumnID = "id"

// CommitEscapedColumnID is the escaped ID SQL column name for the Commit table
const CommitEscapedColumnID = "`id`"

// CommitColumnChecksum is the Checksum SQL column name for the Commit table
const CommitColumnChecksum = "checksum"

// CommitEscapedColumnChecksum is the escaped Checksum SQL column name for the Commit table
const CommitEscapedColumnChecksum = "`checksum`"

// CommitColumnRepoID is the RepoID SQL column name for the Commit table
const CommitColumnRepoID = "repo_id"

// CommitEscapedColumnRepoID is the escaped RepoID SQL column name for the Commit table
const CommitEscapedColumnRepoID = "`repo_id`"

// CommitColumnSha is the Sha SQL column name for the Commit table
const CommitColumnSha = "sha"

// CommitEscapedColumnSha is the escaped Sha SQL column name for the Commit table
const CommitEscapedColumnSha = "`sha`"

// CommitColumnBranch is the Branch SQL column name for the Commit table
const CommitColumnBranch = "branch"

// CommitEscapedColumnBranch is the escaped Branch SQL column name for the Commit table
const CommitEscapedColumnBranch = "`branch`"

// CommitColumnMessage is the Message SQL column name for the Commit table
const CommitColumnMessage = "message"

// CommitEscapedColumnMessage is the escaped Message SQL column name for the Commit table
const CommitEscapedColumnMessage = "`message`"

// CommitColumnMergecommit is the Mergecommit SQL column name for the Commit table
const CommitColumnMergecommit = "mergecommit"

// CommitEscapedColumnMergecommit is the escaped Mergecommit SQL column name for the Commit table
const CommitEscapedColumnMergecommit = "`mergecommit`"

// CommitColumnExcluded is the Excluded SQL column name for the Commit table
const CommitColumnExcluded = "excluded"

// CommitEscapedColumnExcluded is the escaped Excluded SQL column name for the Commit table
const CommitEscapedColumnExcluded = "`excluded`"

// CommitColumnParent is the Parent SQL column name for the Commit table
const CommitColumnParent = "parent"

// CommitEscapedColumnParent is the escaped Parent SQL column name for the Commit table
const CommitEscapedColumnParent = "`parent`"

// CommitColumnParentID is the ParentID SQL column name for the Commit table
const CommitColumnParentID = "parent_id"

// CommitEscapedColumnParentID is the escaped ParentID SQL column name for the Commit table
const CommitEscapedColumnParentID = "`parent_id`"

// CommitColumnDate is the Date SQL column name for the Commit table
const CommitColumnDate = "date"

// CommitEscapedColumnDate is the escaped Date SQL column name for the Commit table
const CommitEscapedColumnDate = "`date`"

// CommitColumnAuthorUserID is the AuthorUserID SQL column name for the Commit table
const CommitColumnAuthorUserID = "author_user_id"

// CommitEscapedColumnAuthorUserID is the escaped AuthorUserID SQL column name for the Commit table
const CommitEscapedColumnAuthorUserID = "`author_user_id`"

// CommitColumnCommitterUserID is the CommitterUserID SQL column name for the Commit table
const CommitColumnCommitterUserID = "committer_user_id"

// CommitEscapedColumnCommitterUserID is the escaped CommitterUserID SQL column name for the Commit table
const CommitEscapedColumnCommitterUserID = "`committer_user_id`"

// CommitColumnOrdinal is the Ordinal SQL column name for the Commit table
const CommitColumnOrdinal = "ordinal"

// CommitEscapedColumnOrdinal is the escaped Ordinal SQL column name for the Commit table
const CommitEscapedColumnOrdinal = "`ordinal`"

// CommitColumnCustomerID is the CustomerID SQL column name for the Commit table
const CommitColumnCustomerID = "customer_id"

// CommitEscapedColumnCustomerID is the escaped CustomerID SQL column name for the Commit table
const CommitEscapedColumnCustomerID = "`customer_id`"

// CommitColumnRefType is the RefType SQL column name for the Commit table
const CommitColumnRefType = "ref_type"

// CommitEscapedColumnRefType is the escaped RefType SQL column name for the Commit table
const CommitEscapedColumnRefType = "`ref_type`"

// CommitColumnRefID is the RefID SQL column name for the Commit table
const CommitColumnRefID = "ref_id"

// CommitEscapedColumnRefID is the escaped RefID SQL column name for the Commit table
const CommitEscapedColumnRefID = "`ref_id`"

// CommitColumnMetadata is the Metadata SQL column name for the Commit table
const CommitColumnMetadata = "metadata"

// CommitEscapedColumnMetadata is the escaped Metadata SQL column name for the Commit table
const CommitEscapedColumnMetadata = "`metadata`"

// GetID will return the Commit ID value
func (t *Commit) GetID() string {
	return t.ID
}

// SetID will set the Commit ID value
func (t *Commit) SetID(v string) {
	t.ID = v
}

// FindCommitByID will find a Commit by ID
func FindCommitByID(ctx context.Context, db DB, value string) (*Commit, error) {
	q := "SELECT `commit`.`id`,`commit`.`checksum`,`commit`.`repo_id`,`commit`.`sha`,`commit`.`branch`,`commit`.`message`,`commit`.`mergecommit`,`commit`.`excluded`,`commit`.`parent`,`commit`.`parent_id`,`commit`.`date`,`commit`.`author_user_id`,`commit`.`committer_user_id`,`commit`.`ordinal`,`commit`.`customer_id`,`commit`.`ref_type`,`commit`.`ref_id`,`commit`.`metadata` FROM `commit` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _RepoID sql.NullString
	var _Sha sql.NullString
	var _Branch sql.NullString
	var _Message sql.NullString
	var _Mergecommit sql.NullBool
	var _Excluded sql.NullBool
	var _Parent sql.NullString
	var _ParentID sql.NullString
	var _Date sql.NullInt64
	var _AuthorUserID sql.NullString
	var _CommitterUserID sql.NullString
	var _Ordinal sql.NullInt64
	var _CustomerID sql.NullString
	var _RefType sql.NullString
	var _RefID sql.NullString
	var _Metadata sql.NullString
	err := db.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_RepoID,
		&_Sha,
		&_Branch,
		&_Message,
		&_Mergecommit,
		&_Excluded,
		&_Parent,
		&_ParentID,
		&_Date,
		&_AuthorUserID,
		&_CommitterUserID,
		&_Ordinal,
		&_CustomerID,
		&_RefType,
		&_RefID,
		&_Metadata,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &Commit{}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _Checksum.Valid {
		t.SetChecksum(_Checksum.String)
	}
	if _RepoID.Valid {
		t.SetRepoID(_RepoID.String)
	}
	if _Sha.Valid {
		t.SetSha(_Sha.String)
	}
	if _Branch.Valid {
		t.SetBranch(_Branch.String)
	}
	if _Message.Valid {
		t.SetMessage(_Message.String)
	}
	if _Mergecommit.Valid {
		t.SetMergecommit(_Mergecommit.Bool)
	}
	if _Excluded.Valid {
		t.SetExcluded(_Excluded.Bool)
	}
	if _Parent.Valid {
		t.SetParent(_Parent.String)
	}
	if _ParentID.Valid {
		t.SetParentID(_ParentID.String)
	}
	if _Date.Valid {
		t.SetDate(_Date.Int64)
	}
	if _AuthorUserID.Valid {
		t.SetAuthorUserID(_AuthorUserID.String)
	}
	if _CommitterUserID.Valid {
		t.SetCommitterUserID(_CommitterUserID.String)
	}
	if _Ordinal.Valid {
		t.SetOrdinal(int32(_Ordinal.Int64))
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _RefType.Valid {
		t.SetRefType(_RefType.String)
	}
	if _RefID.Valid {
		t.SetRefID(_RefID.String)
	}
	if _Metadata.Valid {
		t.SetMetadata(_Metadata.String)
	}
	return t, nil
}

// FindCommitByIDTx will find a Commit by ID using the provided transaction
func FindCommitByIDTx(ctx context.Context, tx Tx, value string) (*Commit, error) {
	q := "SELECT `commit`.`id`,`commit`.`checksum`,`commit`.`repo_id`,`commit`.`sha`,`commit`.`branch`,`commit`.`message`,`commit`.`mergecommit`,`commit`.`excluded`,`commit`.`parent`,`commit`.`parent_id`,`commit`.`date`,`commit`.`author_user_id`,`commit`.`committer_user_id`,`commit`.`ordinal`,`commit`.`customer_id`,`commit`.`ref_type`,`commit`.`ref_id`,`commit`.`metadata` FROM `commit` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _RepoID sql.NullString
	var _Sha sql.NullString
	var _Branch sql.NullString
	var _Message sql.NullString
	var _Mergecommit sql.NullBool
	var _Excluded sql.NullBool
	var _Parent sql.NullString
	var _ParentID sql.NullString
	var _Date sql.NullInt64
	var _AuthorUserID sql.NullString
	var _CommitterUserID sql.NullString
	var _Ordinal sql.NullInt64
	var _CustomerID sql.NullString
	var _RefType sql.NullString
	var _RefID sql.NullString
	var _Metadata sql.NullString
	err := tx.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_RepoID,
		&_Sha,
		&_Branch,
		&_Message,
		&_Mergecommit,
		&_Excluded,
		&_Parent,
		&_ParentID,
		&_Date,
		&_AuthorUserID,
		&_CommitterUserID,
		&_Ordinal,
		&_CustomerID,
		&_RefType,
		&_RefID,
		&_Metadata,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &Commit{}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _Checksum.Valid {
		t.SetChecksum(_Checksum.String)
	}
	if _RepoID.Valid {
		t.SetRepoID(_RepoID.String)
	}
	if _Sha.Valid {
		t.SetSha(_Sha.String)
	}
	if _Branch.Valid {
		t.SetBranch(_Branch.String)
	}
	if _Message.Valid {
		t.SetMessage(_Message.String)
	}
	if _Mergecommit.Valid {
		t.SetMergecommit(_Mergecommit.Bool)
	}
	if _Excluded.Valid {
		t.SetExcluded(_Excluded.Bool)
	}
	if _Parent.Valid {
		t.SetParent(_Parent.String)
	}
	if _ParentID.Valid {
		t.SetParentID(_ParentID.String)
	}
	if _Date.Valid {
		t.SetDate(_Date.Int64)
	}
	if _AuthorUserID.Valid {
		t.SetAuthorUserID(_AuthorUserID.String)
	}
	if _CommitterUserID.Valid {
		t.SetCommitterUserID(_CommitterUserID.String)
	}
	if _Ordinal.Valid {
		t.SetOrdinal(int32(_Ordinal.Int64))
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _RefType.Valid {
		t.SetRefType(_RefType.String)
	}
	if _RefID.Valid {
		t.SetRefID(_RefID.String)
	}
	if _Metadata.Valid {
		t.SetMetadata(_Metadata.String)
	}
	return t, nil
}

// GetChecksum will return the Commit Checksum value
func (t *Commit) GetChecksum() string {
	if t.Checksum == nil {
		return ""
	}
	return *t.Checksum
}

// SetChecksum will set the Commit Checksum value
func (t *Commit) SetChecksum(v string) {
	t.Checksum = &v
}

// GetRepoID will return the Commit RepoID value
func (t *Commit) GetRepoID() string {
	return t.RepoID
}

// SetRepoID will set the Commit RepoID value
func (t *Commit) SetRepoID(v string) {
	t.RepoID = v
}

// FindCommitsByRepoID will find all Commits by the RepoID value
func FindCommitsByRepoID(ctx context.Context, db DB, value string) ([]*Commit, error) {
	q := "SELECT `commit`.`id`,`commit`.`checksum`,`commit`.`repo_id`,`commit`.`sha`,`commit`.`branch`,`commit`.`message`,`commit`.`mergecommit`,`commit`.`excluded`,`commit`.`parent`,`commit`.`parent_id`,`commit`.`date`,`commit`.`author_user_id`,`commit`.`committer_user_id`,`commit`.`ordinal`,`commit`.`customer_id`,`commit`.`ref_type`,`commit`.`ref_id`,`commit`.`metadata` FROM `commit` WHERE `repo_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*Commit, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _RepoID sql.NullString
		var _Sha sql.NullString
		var _Branch sql.NullString
		var _Message sql.NullString
		var _Mergecommit sql.NullBool
		var _Excluded sql.NullBool
		var _Parent sql.NullString
		var _ParentID sql.NullString
		var _Date sql.NullInt64
		var _AuthorUserID sql.NullString
		var _CommitterUserID sql.NullString
		var _Ordinal sql.NullInt64
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefID sql.NullString
		var _Metadata sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_RepoID,
			&_Sha,
			&_Branch,
			&_Message,
			&_Mergecommit,
			&_Excluded,
			&_Parent,
			&_ParentID,
			&_Date,
			&_AuthorUserID,
			&_CommitterUserID,
			&_Ordinal,
			&_CustomerID,
			&_RefType,
			&_RefID,
			&_Metadata,
		)
		if err != nil {
			return nil, err
		}
		t := &Commit{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _RepoID.Valid {
			t.SetRepoID(_RepoID.String)
		}
		if _Sha.Valid {
			t.SetSha(_Sha.String)
		}
		if _Branch.Valid {
			t.SetBranch(_Branch.String)
		}
		if _Message.Valid {
			t.SetMessage(_Message.String)
		}
		if _Mergecommit.Valid {
			t.SetMergecommit(_Mergecommit.Bool)
		}
		if _Excluded.Valid {
			t.SetExcluded(_Excluded.Bool)
		}
		if _Parent.Valid {
			t.SetParent(_Parent.String)
		}
		if _ParentID.Valid {
			t.SetParentID(_ParentID.String)
		}
		if _Date.Valid {
			t.SetDate(_Date.Int64)
		}
		if _AuthorUserID.Valid {
			t.SetAuthorUserID(_AuthorUserID.String)
		}
		if _CommitterUserID.Valid {
			t.SetCommitterUserID(_CommitterUserID.String)
		}
		if _Ordinal.Valid {
			t.SetOrdinal(int32(_Ordinal.Int64))
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefType.Valid {
			t.SetRefType(_RefType.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		if _Metadata.Valid {
			t.SetMetadata(_Metadata.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindCommitsByRepoIDTx will find all Commits by the RepoID value using the provided transaction
func FindCommitsByRepoIDTx(ctx context.Context, tx Tx, value string) ([]*Commit, error) {
	q := "SELECT `commit`.`id`,`commit`.`checksum`,`commit`.`repo_id`,`commit`.`sha`,`commit`.`branch`,`commit`.`message`,`commit`.`mergecommit`,`commit`.`excluded`,`commit`.`parent`,`commit`.`parent_id`,`commit`.`date`,`commit`.`author_user_id`,`commit`.`committer_user_id`,`commit`.`ordinal`,`commit`.`customer_id`,`commit`.`ref_type`,`commit`.`ref_id`,`commit`.`metadata` FROM `commit` WHERE `repo_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*Commit, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _RepoID sql.NullString
		var _Sha sql.NullString
		var _Branch sql.NullString
		var _Message sql.NullString
		var _Mergecommit sql.NullBool
		var _Excluded sql.NullBool
		var _Parent sql.NullString
		var _ParentID sql.NullString
		var _Date sql.NullInt64
		var _AuthorUserID sql.NullString
		var _CommitterUserID sql.NullString
		var _Ordinal sql.NullInt64
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefID sql.NullString
		var _Metadata sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_RepoID,
			&_Sha,
			&_Branch,
			&_Message,
			&_Mergecommit,
			&_Excluded,
			&_Parent,
			&_ParentID,
			&_Date,
			&_AuthorUserID,
			&_CommitterUserID,
			&_Ordinal,
			&_CustomerID,
			&_RefType,
			&_RefID,
			&_Metadata,
		)
		if err != nil {
			return nil, err
		}
		t := &Commit{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _RepoID.Valid {
			t.SetRepoID(_RepoID.String)
		}
		if _Sha.Valid {
			t.SetSha(_Sha.String)
		}
		if _Branch.Valid {
			t.SetBranch(_Branch.String)
		}
		if _Message.Valid {
			t.SetMessage(_Message.String)
		}
		if _Mergecommit.Valid {
			t.SetMergecommit(_Mergecommit.Bool)
		}
		if _Excluded.Valid {
			t.SetExcluded(_Excluded.Bool)
		}
		if _Parent.Valid {
			t.SetParent(_Parent.String)
		}
		if _ParentID.Valid {
			t.SetParentID(_ParentID.String)
		}
		if _Date.Valid {
			t.SetDate(_Date.Int64)
		}
		if _AuthorUserID.Valid {
			t.SetAuthorUserID(_AuthorUserID.String)
		}
		if _CommitterUserID.Valid {
			t.SetCommitterUserID(_CommitterUserID.String)
		}
		if _Ordinal.Valid {
			t.SetOrdinal(int32(_Ordinal.Int64))
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefType.Valid {
			t.SetRefType(_RefType.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		if _Metadata.Valid {
			t.SetMetadata(_Metadata.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetSha will return the Commit Sha value
func (t *Commit) GetSha() string {
	return t.Sha
}

// SetSha will set the Commit Sha value
func (t *Commit) SetSha(v string) {
	t.Sha = v
}

// FindCommitsBySha will find all Commits by the Sha value
func FindCommitsBySha(ctx context.Context, db DB, value string) ([]*Commit, error) {
	q := "SELECT `commit`.`id`,`commit`.`checksum`,`commit`.`repo_id`,`commit`.`sha`,`commit`.`branch`,`commit`.`message`,`commit`.`mergecommit`,`commit`.`excluded`,`commit`.`parent`,`commit`.`parent_id`,`commit`.`date`,`commit`.`author_user_id`,`commit`.`committer_user_id`,`commit`.`ordinal`,`commit`.`customer_id`,`commit`.`ref_type`,`commit`.`ref_id`,`commit`.`metadata` FROM `commit` WHERE `sha` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*Commit, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _RepoID sql.NullString
		var _Sha sql.NullString
		var _Branch sql.NullString
		var _Message sql.NullString
		var _Mergecommit sql.NullBool
		var _Excluded sql.NullBool
		var _Parent sql.NullString
		var _ParentID sql.NullString
		var _Date sql.NullInt64
		var _AuthorUserID sql.NullString
		var _CommitterUserID sql.NullString
		var _Ordinal sql.NullInt64
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefID sql.NullString
		var _Metadata sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_RepoID,
			&_Sha,
			&_Branch,
			&_Message,
			&_Mergecommit,
			&_Excluded,
			&_Parent,
			&_ParentID,
			&_Date,
			&_AuthorUserID,
			&_CommitterUserID,
			&_Ordinal,
			&_CustomerID,
			&_RefType,
			&_RefID,
			&_Metadata,
		)
		if err != nil {
			return nil, err
		}
		t := &Commit{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _RepoID.Valid {
			t.SetRepoID(_RepoID.String)
		}
		if _Sha.Valid {
			t.SetSha(_Sha.String)
		}
		if _Branch.Valid {
			t.SetBranch(_Branch.String)
		}
		if _Message.Valid {
			t.SetMessage(_Message.String)
		}
		if _Mergecommit.Valid {
			t.SetMergecommit(_Mergecommit.Bool)
		}
		if _Excluded.Valid {
			t.SetExcluded(_Excluded.Bool)
		}
		if _Parent.Valid {
			t.SetParent(_Parent.String)
		}
		if _ParentID.Valid {
			t.SetParentID(_ParentID.String)
		}
		if _Date.Valid {
			t.SetDate(_Date.Int64)
		}
		if _AuthorUserID.Valid {
			t.SetAuthorUserID(_AuthorUserID.String)
		}
		if _CommitterUserID.Valid {
			t.SetCommitterUserID(_CommitterUserID.String)
		}
		if _Ordinal.Valid {
			t.SetOrdinal(int32(_Ordinal.Int64))
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefType.Valid {
			t.SetRefType(_RefType.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		if _Metadata.Valid {
			t.SetMetadata(_Metadata.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindCommitsByShaTx will find all Commits by the Sha value using the provided transaction
func FindCommitsByShaTx(ctx context.Context, tx Tx, value string) ([]*Commit, error) {
	q := "SELECT `commit`.`id`,`commit`.`checksum`,`commit`.`repo_id`,`commit`.`sha`,`commit`.`branch`,`commit`.`message`,`commit`.`mergecommit`,`commit`.`excluded`,`commit`.`parent`,`commit`.`parent_id`,`commit`.`date`,`commit`.`author_user_id`,`commit`.`committer_user_id`,`commit`.`ordinal`,`commit`.`customer_id`,`commit`.`ref_type`,`commit`.`ref_id`,`commit`.`metadata` FROM `commit` WHERE `sha` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*Commit, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _RepoID sql.NullString
		var _Sha sql.NullString
		var _Branch sql.NullString
		var _Message sql.NullString
		var _Mergecommit sql.NullBool
		var _Excluded sql.NullBool
		var _Parent sql.NullString
		var _ParentID sql.NullString
		var _Date sql.NullInt64
		var _AuthorUserID sql.NullString
		var _CommitterUserID sql.NullString
		var _Ordinal sql.NullInt64
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefID sql.NullString
		var _Metadata sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_RepoID,
			&_Sha,
			&_Branch,
			&_Message,
			&_Mergecommit,
			&_Excluded,
			&_Parent,
			&_ParentID,
			&_Date,
			&_AuthorUserID,
			&_CommitterUserID,
			&_Ordinal,
			&_CustomerID,
			&_RefType,
			&_RefID,
			&_Metadata,
		)
		if err != nil {
			return nil, err
		}
		t := &Commit{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _RepoID.Valid {
			t.SetRepoID(_RepoID.String)
		}
		if _Sha.Valid {
			t.SetSha(_Sha.String)
		}
		if _Branch.Valid {
			t.SetBranch(_Branch.String)
		}
		if _Message.Valid {
			t.SetMessage(_Message.String)
		}
		if _Mergecommit.Valid {
			t.SetMergecommit(_Mergecommit.Bool)
		}
		if _Excluded.Valid {
			t.SetExcluded(_Excluded.Bool)
		}
		if _Parent.Valid {
			t.SetParent(_Parent.String)
		}
		if _ParentID.Valid {
			t.SetParentID(_ParentID.String)
		}
		if _Date.Valid {
			t.SetDate(_Date.Int64)
		}
		if _AuthorUserID.Valid {
			t.SetAuthorUserID(_AuthorUserID.String)
		}
		if _CommitterUserID.Valid {
			t.SetCommitterUserID(_CommitterUserID.String)
		}
		if _Ordinal.Valid {
			t.SetOrdinal(int32(_Ordinal.Int64))
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefType.Valid {
			t.SetRefType(_RefType.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		if _Metadata.Valid {
			t.SetMetadata(_Metadata.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetBranch will return the Commit Branch value
func (t *Commit) GetBranch() string {
	if t.Branch == nil {
		return ""
	}
	return *t.Branch
}

// SetBranch will set the Commit Branch value
func (t *Commit) SetBranch(v string) {
	t.Branch = &v
}

// GetMessage will return the Commit Message value
func (t *Commit) GetMessage() string {
	if t.Message == nil {
		return ""
	}
	return *t.Message
}

// SetMessage will set the Commit Message value
func (t *Commit) SetMessage(v string) {
	t.Message = &v
}

// GetMergecommit will return the Commit Mergecommit value
func (t *Commit) GetMergecommit() bool {
	if t.Mergecommit == nil {
		return false
	}
	return *t.Mergecommit
}

// SetMergecommit will set the Commit Mergecommit value
func (t *Commit) SetMergecommit(v bool) {
	t.Mergecommit = &v
}

// GetExcluded will return the Commit Excluded value
func (t *Commit) GetExcluded() bool {
	if t.Excluded == nil {
		return false
	}
	return *t.Excluded
}

// SetExcluded will set the Commit Excluded value
func (t *Commit) SetExcluded(v bool) {
	t.Excluded = &v
}

// GetParent will return the Commit Parent value
func (t *Commit) GetParent() string {
	if t.Parent == nil {
		return ""
	}
	return *t.Parent
}

// SetParent will set the Commit Parent value
func (t *Commit) SetParent(v string) {
	t.Parent = &v
}

// GetParentID will return the Commit ParentID value
func (t *Commit) GetParentID() string {
	if t.ParentID == nil {
		return ""
	}
	return *t.ParentID
}

// SetParentID will set the Commit ParentID value
func (t *Commit) SetParentID(v string) {
	t.ParentID = &v
}

// GetDate will return the Commit Date value
func (t *Commit) GetDate() int64 {
	return t.Date
}

// SetDate will set the Commit Date value
func (t *Commit) SetDate(v int64) {
	t.Date = v
}

// GetAuthorUserID will return the Commit AuthorUserID value
func (t *Commit) GetAuthorUserID() string {
	return t.AuthorUserID
}

// SetAuthorUserID will set the Commit AuthorUserID value
func (t *Commit) SetAuthorUserID(v string) {
	t.AuthorUserID = v
}

// FindCommitsByAuthorUserID will find all Commits by the AuthorUserID value
func FindCommitsByAuthorUserID(ctx context.Context, db DB, value string) ([]*Commit, error) {
	q := "SELECT `commit`.`id`,`commit`.`checksum`,`commit`.`repo_id`,`commit`.`sha`,`commit`.`branch`,`commit`.`message`,`commit`.`mergecommit`,`commit`.`excluded`,`commit`.`parent`,`commit`.`parent_id`,`commit`.`date`,`commit`.`author_user_id`,`commit`.`committer_user_id`,`commit`.`ordinal`,`commit`.`customer_id`,`commit`.`ref_type`,`commit`.`ref_id`,`commit`.`metadata` FROM `commit` WHERE `author_user_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*Commit, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _RepoID sql.NullString
		var _Sha sql.NullString
		var _Branch sql.NullString
		var _Message sql.NullString
		var _Mergecommit sql.NullBool
		var _Excluded sql.NullBool
		var _Parent sql.NullString
		var _ParentID sql.NullString
		var _Date sql.NullInt64
		var _AuthorUserID sql.NullString
		var _CommitterUserID sql.NullString
		var _Ordinal sql.NullInt64
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefID sql.NullString
		var _Metadata sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_RepoID,
			&_Sha,
			&_Branch,
			&_Message,
			&_Mergecommit,
			&_Excluded,
			&_Parent,
			&_ParentID,
			&_Date,
			&_AuthorUserID,
			&_CommitterUserID,
			&_Ordinal,
			&_CustomerID,
			&_RefType,
			&_RefID,
			&_Metadata,
		)
		if err != nil {
			return nil, err
		}
		t := &Commit{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _RepoID.Valid {
			t.SetRepoID(_RepoID.String)
		}
		if _Sha.Valid {
			t.SetSha(_Sha.String)
		}
		if _Branch.Valid {
			t.SetBranch(_Branch.String)
		}
		if _Message.Valid {
			t.SetMessage(_Message.String)
		}
		if _Mergecommit.Valid {
			t.SetMergecommit(_Mergecommit.Bool)
		}
		if _Excluded.Valid {
			t.SetExcluded(_Excluded.Bool)
		}
		if _Parent.Valid {
			t.SetParent(_Parent.String)
		}
		if _ParentID.Valid {
			t.SetParentID(_ParentID.String)
		}
		if _Date.Valid {
			t.SetDate(_Date.Int64)
		}
		if _AuthorUserID.Valid {
			t.SetAuthorUserID(_AuthorUserID.String)
		}
		if _CommitterUserID.Valid {
			t.SetCommitterUserID(_CommitterUserID.String)
		}
		if _Ordinal.Valid {
			t.SetOrdinal(int32(_Ordinal.Int64))
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefType.Valid {
			t.SetRefType(_RefType.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		if _Metadata.Valid {
			t.SetMetadata(_Metadata.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindCommitsByAuthorUserIDTx will find all Commits by the AuthorUserID value using the provided transaction
func FindCommitsByAuthorUserIDTx(ctx context.Context, tx Tx, value string) ([]*Commit, error) {
	q := "SELECT `commit`.`id`,`commit`.`checksum`,`commit`.`repo_id`,`commit`.`sha`,`commit`.`branch`,`commit`.`message`,`commit`.`mergecommit`,`commit`.`excluded`,`commit`.`parent`,`commit`.`parent_id`,`commit`.`date`,`commit`.`author_user_id`,`commit`.`committer_user_id`,`commit`.`ordinal`,`commit`.`customer_id`,`commit`.`ref_type`,`commit`.`ref_id`,`commit`.`metadata` FROM `commit` WHERE `author_user_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*Commit, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _RepoID sql.NullString
		var _Sha sql.NullString
		var _Branch sql.NullString
		var _Message sql.NullString
		var _Mergecommit sql.NullBool
		var _Excluded sql.NullBool
		var _Parent sql.NullString
		var _ParentID sql.NullString
		var _Date sql.NullInt64
		var _AuthorUserID sql.NullString
		var _CommitterUserID sql.NullString
		var _Ordinal sql.NullInt64
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefID sql.NullString
		var _Metadata sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_RepoID,
			&_Sha,
			&_Branch,
			&_Message,
			&_Mergecommit,
			&_Excluded,
			&_Parent,
			&_ParentID,
			&_Date,
			&_AuthorUserID,
			&_CommitterUserID,
			&_Ordinal,
			&_CustomerID,
			&_RefType,
			&_RefID,
			&_Metadata,
		)
		if err != nil {
			return nil, err
		}
		t := &Commit{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _RepoID.Valid {
			t.SetRepoID(_RepoID.String)
		}
		if _Sha.Valid {
			t.SetSha(_Sha.String)
		}
		if _Branch.Valid {
			t.SetBranch(_Branch.String)
		}
		if _Message.Valid {
			t.SetMessage(_Message.String)
		}
		if _Mergecommit.Valid {
			t.SetMergecommit(_Mergecommit.Bool)
		}
		if _Excluded.Valid {
			t.SetExcluded(_Excluded.Bool)
		}
		if _Parent.Valid {
			t.SetParent(_Parent.String)
		}
		if _ParentID.Valid {
			t.SetParentID(_ParentID.String)
		}
		if _Date.Valid {
			t.SetDate(_Date.Int64)
		}
		if _AuthorUserID.Valid {
			t.SetAuthorUserID(_AuthorUserID.String)
		}
		if _CommitterUserID.Valid {
			t.SetCommitterUserID(_CommitterUserID.String)
		}
		if _Ordinal.Valid {
			t.SetOrdinal(int32(_Ordinal.Int64))
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefType.Valid {
			t.SetRefType(_RefType.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		if _Metadata.Valid {
			t.SetMetadata(_Metadata.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetCommitterUserID will return the Commit CommitterUserID value
func (t *Commit) GetCommitterUserID() string {
	return t.CommitterUserID
}

// SetCommitterUserID will set the Commit CommitterUserID value
func (t *Commit) SetCommitterUserID(v string) {
	t.CommitterUserID = v
}

// FindCommitsByCommitterUserID will find all Commits by the CommitterUserID value
func FindCommitsByCommitterUserID(ctx context.Context, db DB, value string) ([]*Commit, error) {
	q := "SELECT `commit`.`id`,`commit`.`checksum`,`commit`.`repo_id`,`commit`.`sha`,`commit`.`branch`,`commit`.`message`,`commit`.`mergecommit`,`commit`.`excluded`,`commit`.`parent`,`commit`.`parent_id`,`commit`.`date`,`commit`.`author_user_id`,`commit`.`committer_user_id`,`commit`.`ordinal`,`commit`.`customer_id`,`commit`.`ref_type`,`commit`.`ref_id`,`commit`.`metadata` FROM `commit` WHERE `committer_user_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*Commit, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _RepoID sql.NullString
		var _Sha sql.NullString
		var _Branch sql.NullString
		var _Message sql.NullString
		var _Mergecommit sql.NullBool
		var _Excluded sql.NullBool
		var _Parent sql.NullString
		var _ParentID sql.NullString
		var _Date sql.NullInt64
		var _AuthorUserID sql.NullString
		var _CommitterUserID sql.NullString
		var _Ordinal sql.NullInt64
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefID sql.NullString
		var _Metadata sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_RepoID,
			&_Sha,
			&_Branch,
			&_Message,
			&_Mergecommit,
			&_Excluded,
			&_Parent,
			&_ParentID,
			&_Date,
			&_AuthorUserID,
			&_CommitterUserID,
			&_Ordinal,
			&_CustomerID,
			&_RefType,
			&_RefID,
			&_Metadata,
		)
		if err != nil {
			return nil, err
		}
		t := &Commit{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _RepoID.Valid {
			t.SetRepoID(_RepoID.String)
		}
		if _Sha.Valid {
			t.SetSha(_Sha.String)
		}
		if _Branch.Valid {
			t.SetBranch(_Branch.String)
		}
		if _Message.Valid {
			t.SetMessage(_Message.String)
		}
		if _Mergecommit.Valid {
			t.SetMergecommit(_Mergecommit.Bool)
		}
		if _Excluded.Valid {
			t.SetExcluded(_Excluded.Bool)
		}
		if _Parent.Valid {
			t.SetParent(_Parent.String)
		}
		if _ParentID.Valid {
			t.SetParentID(_ParentID.String)
		}
		if _Date.Valid {
			t.SetDate(_Date.Int64)
		}
		if _AuthorUserID.Valid {
			t.SetAuthorUserID(_AuthorUserID.String)
		}
		if _CommitterUserID.Valid {
			t.SetCommitterUserID(_CommitterUserID.String)
		}
		if _Ordinal.Valid {
			t.SetOrdinal(int32(_Ordinal.Int64))
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefType.Valid {
			t.SetRefType(_RefType.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		if _Metadata.Valid {
			t.SetMetadata(_Metadata.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindCommitsByCommitterUserIDTx will find all Commits by the CommitterUserID value using the provided transaction
func FindCommitsByCommitterUserIDTx(ctx context.Context, tx Tx, value string) ([]*Commit, error) {
	q := "SELECT `commit`.`id`,`commit`.`checksum`,`commit`.`repo_id`,`commit`.`sha`,`commit`.`branch`,`commit`.`message`,`commit`.`mergecommit`,`commit`.`excluded`,`commit`.`parent`,`commit`.`parent_id`,`commit`.`date`,`commit`.`author_user_id`,`commit`.`committer_user_id`,`commit`.`ordinal`,`commit`.`customer_id`,`commit`.`ref_type`,`commit`.`ref_id`,`commit`.`metadata` FROM `commit` WHERE `committer_user_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*Commit, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _RepoID sql.NullString
		var _Sha sql.NullString
		var _Branch sql.NullString
		var _Message sql.NullString
		var _Mergecommit sql.NullBool
		var _Excluded sql.NullBool
		var _Parent sql.NullString
		var _ParentID sql.NullString
		var _Date sql.NullInt64
		var _AuthorUserID sql.NullString
		var _CommitterUserID sql.NullString
		var _Ordinal sql.NullInt64
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefID sql.NullString
		var _Metadata sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_RepoID,
			&_Sha,
			&_Branch,
			&_Message,
			&_Mergecommit,
			&_Excluded,
			&_Parent,
			&_ParentID,
			&_Date,
			&_AuthorUserID,
			&_CommitterUserID,
			&_Ordinal,
			&_CustomerID,
			&_RefType,
			&_RefID,
			&_Metadata,
		)
		if err != nil {
			return nil, err
		}
		t := &Commit{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _RepoID.Valid {
			t.SetRepoID(_RepoID.String)
		}
		if _Sha.Valid {
			t.SetSha(_Sha.String)
		}
		if _Branch.Valid {
			t.SetBranch(_Branch.String)
		}
		if _Message.Valid {
			t.SetMessage(_Message.String)
		}
		if _Mergecommit.Valid {
			t.SetMergecommit(_Mergecommit.Bool)
		}
		if _Excluded.Valid {
			t.SetExcluded(_Excluded.Bool)
		}
		if _Parent.Valid {
			t.SetParent(_Parent.String)
		}
		if _ParentID.Valid {
			t.SetParentID(_ParentID.String)
		}
		if _Date.Valid {
			t.SetDate(_Date.Int64)
		}
		if _AuthorUserID.Valid {
			t.SetAuthorUserID(_AuthorUserID.String)
		}
		if _CommitterUserID.Valid {
			t.SetCommitterUserID(_CommitterUserID.String)
		}
		if _Ordinal.Valid {
			t.SetOrdinal(int32(_Ordinal.Int64))
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefType.Valid {
			t.SetRefType(_RefType.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		if _Metadata.Valid {
			t.SetMetadata(_Metadata.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetOrdinal will return the Commit Ordinal value
func (t *Commit) GetOrdinal() int32 {
	if t.Ordinal == nil {
		return int32(0)
	}
	return *t.Ordinal
}

// SetOrdinal will set the Commit Ordinal value
func (t *Commit) SetOrdinal(v int32) {
	t.Ordinal = &v
}

// GetCustomerID will return the Commit CustomerID value
func (t *Commit) GetCustomerID() string {
	return t.CustomerID
}

// SetCustomerID will set the Commit CustomerID value
func (t *Commit) SetCustomerID(v string) {
	t.CustomerID = v
}

// FindCommitsByCustomerID will find all Commits by the CustomerID value
func FindCommitsByCustomerID(ctx context.Context, db DB, value string) ([]*Commit, error) {
	q := "SELECT `commit`.`id`,`commit`.`checksum`,`commit`.`repo_id`,`commit`.`sha`,`commit`.`branch`,`commit`.`message`,`commit`.`mergecommit`,`commit`.`excluded`,`commit`.`parent`,`commit`.`parent_id`,`commit`.`date`,`commit`.`author_user_id`,`commit`.`committer_user_id`,`commit`.`ordinal`,`commit`.`customer_id`,`commit`.`ref_type`,`commit`.`ref_id`,`commit`.`metadata` FROM `commit` WHERE `customer_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*Commit, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _RepoID sql.NullString
		var _Sha sql.NullString
		var _Branch sql.NullString
		var _Message sql.NullString
		var _Mergecommit sql.NullBool
		var _Excluded sql.NullBool
		var _Parent sql.NullString
		var _ParentID sql.NullString
		var _Date sql.NullInt64
		var _AuthorUserID sql.NullString
		var _CommitterUserID sql.NullString
		var _Ordinal sql.NullInt64
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefID sql.NullString
		var _Metadata sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_RepoID,
			&_Sha,
			&_Branch,
			&_Message,
			&_Mergecommit,
			&_Excluded,
			&_Parent,
			&_ParentID,
			&_Date,
			&_AuthorUserID,
			&_CommitterUserID,
			&_Ordinal,
			&_CustomerID,
			&_RefType,
			&_RefID,
			&_Metadata,
		)
		if err != nil {
			return nil, err
		}
		t := &Commit{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _RepoID.Valid {
			t.SetRepoID(_RepoID.String)
		}
		if _Sha.Valid {
			t.SetSha(_Sha.String)
		}
		if _Branch.Valid {
			t.SetBranch(_Branch.String)
		}
		if _Message.Valid {
			t.SetMessage(_Message.String)
		}
		if _Mergecommit.Valid {
			t.SetMergecommit(_Mergecommit.Bool)
		}
		if _Excluded.Valid {
			t.SetExcluded(_Excluded.Bool)
		}
		if _Parent.Valid {
			t.SetParent(_Parent.String)
		}
		if _ParentID.Valid {
			t.SetParentID(_ParentID.String)
		}
		if _Date.Valid {
			t.SetDate(_Date.Int64)
		}
		if _AuthorUserID.Valid {
			t.SetAuthorUserID(_AuthorUserID.String)
		}
		if _CommitterUserID.Valid {
			t.SetCommitterUserID(_CommitterUserID.String)
		}
		if _Ordinal.Valid {
			t.SetOrdinal(int32(_Ordinal.Int64))
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefType.Valid {
			t.SetRefType(_RefType.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		if _Metadata.Valid {
			t.SetMetadata(_Metadata.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindCommitsByCustomerIDTx will find all Commits by the CustomerID value using the provided transaction
func FindCommitsByCustomerIDTx(ctx context.Context, tx Tx, value string) ([]*Commit, error) {
	q := "SELECT `commit`.`id`,`commit`.`checksum`,`commit`.`repo_id`,`commit`.`sha`,`commit`.`branch`,`commit`.`message`,`commit`.`mergecommit`,`commit`.`excluded`,`commit`.`parent`,`commit`.`parent_id`,`commit`.`date`,`commit`.`author_user_id`,`commit`.`committer_user_id`,`commit`.`ordinal`,`commit`.`customer_id`,`commit`.`ref_type`,`commit`.`ref_id`,`commit`.`metadata` FROM `commit` WHERE `customer_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*Commit, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _RepoID sql.NullString
		var _Sha sql.NullString
		var _Branch sql.NullString
		var _Message sql.NullString
		var _Mergecommit sql.NullBool
		var _Excluded sql.NullBool
		var _Parent sql.NullString
		var _ParentID sql.NullString
		var _Date sql.NullInt64
		var _AuthorUserID sql.NullString
		var _CommitterUserID sql.NullString
		var _Ordinal sql.NullInt64
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefID sql.NullString
		var _Metadata sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_RepoID,
			&_Sha,
			&_Branch,
			&_Message,
			&_Mergecommit,
			&_Excluded,
			&_Parent,
			&_ParentID,
			&_Date,
			&_AuthorUserID,
			&_CommitterUserID,
			&_Ordinal,
			&_CustomerID,
			&_RefType,
			&_RefID,
			&_Metadata,
		)
		if err != nil {
			return nil, err
		}
		t := &Commit{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _RepoID.Valid {
			t.SetRepoID(_RepoID.String)
		}
		if _Sha.Valid {
			t.SetSha(_Sha.String)
		}
		if _Branch.Valid {
			t.SetBranch(_Branch.String)
		}
		if _Message.Valid {
			t.SetMessage(_Message.String)
		}
		if _Mergecommit.Valid {
			t.SetMergecommit(_Mergecommit.Bool)
		}
		if _Excluded.Valid {
			t.SetExcluded(_Excluded.Bool)
		}
		if _Parent.Valid {
			t.SetParent(_Parent.String)
		}
		if _ParentID.Valid {
			t.SetParentID(_ParentID.String)
		}
		if _Date.Valid {
			t.SetDate(_Date.Int64)
		}
		if _AuthorUserID.Valid {
			t.SetAuthorUserID(_AuthorUserID.String)
		}
		if _CommitterUserID.Valid {
			t.SetCommitterUserID(_CommitterUserID.String)
		}
		if _Ordinal.Valid {
			t.SetOrdinal(int32(_Ordinal.Int64))
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefType.Valid {
			t.SetRefType(_RefType.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		if _Metadata.Valid {
			t.SetMetadata(_Metadata.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetRefType will return the Commit RefType value
func (t *Commit) GetRefType() string {
	return t.RefType
}

// SetRefType will set the Commit RefType value
func (t *Commit) SetRefType(v string) {
	t.RefType = v
}

// FindCommitsByRefType will find all Commits by the RefType value
func FindCommitsByRefType(ctx context.Context, db DB, value string) ([]*Commit, error) {
	q := "SELECT `commit`.`id`,`commit`.`checksum`,`commit`.`repo_id`,`commit`.`sha`,`commit`.`branch`,`commit`.`message`,`commit`.`mergecommit`,`commit`.`excluded`,`commit`.`parent`,`commit`.`parent_id`,`commit`.`date`,`commit`.`author_user_id`,`commit`.`committer_user_id`,`commit`.`ordinal`,`commit`.`customer_id`,`commit`.`ref_type`,`commit`.`ref_id`,`commit`.`metadata` FROM `commit` WHERE `ref_type` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*Commit, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _RepoID sql.NullString
		var _Sha sql.NullString
		var _Branch sql.NullString
		var _Message sql.NullString
		var _Mergecommit sql.NullBool
		var _Excluded sql.NullBool
		var _Parent sql.NullString
		var _ParentID sql.NullString
		var _Date sql.NullInt64
		var _AuthorUserID sql.NullString
		var _CommitterUserID sql.NullString
		var _Ordinal sql.NullInt64
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefID sql.NullString
		var _Metadata sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_RepoID,
			&_Sha,
			&_Branch,
			&_Message,
			&_Mergecommit,
			&_Excluded,
			&_Parent,
			&_ParentID,
			&_Date,
			&_AuthorUserID,
			&_CommitterUserID,
			&_Ordinal,
			&_CustomerID,
			&_RefType,
			&_RefID,
			&_Metadata,
		)
		if err != nil {
			return nil, err
		}
		t := &Commit{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _RepoID.Valid {
			t.SetRepoID(_RepoID.String)
		}
		if _Sha.Valid {
			t.SetSha(_Sha.String)
		}
		if _Branch.Valid {
			t.SetBranch(_Branch.String)
		}
		if _Message.Valid {
			t.SetMessage(_Message.String)
		}
		if _Mergecommit.Valid {
			t.SetMergecommit(_Mergecommit.Bool)
		}
		if _Excluded.Valid {
			t.SetExcluded(_Excluded.Bool)
		}
		if _Parent.Valid {
			t.SetParent(_Parent.String)
		}
		if _ParentID.Valid {
			t.SetParentID(_ParentID.String)
		}
		if _Date.Valid {
			t.SetDate(_Date.Int64)
		}
		if _AuthorUserID.Valid {
			t.SetAuthorUserID(_AuthorUserID.String)
		}
		if _CommitterUserID.Valid {
			t.SetCommitterUserID(_CommitterUserID.String)
		}
		if _Ordinal.Valid {
			t.SetOrdinal(int32(_Ordinal.Int64))
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefType.Valid {
			t.SetRefType(_RefType.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		if _Metadata.Valid {
			t.SetMetadata(_Metadata.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindCommitsByRefTypeTx will find all Commits by the RefType value using the provided transaction
func FindCommitsByRefTypeTx(ctx context.Context, tx Tx, value string) ([]*Commit, error) {
	q := "SELECT `commit`.`id`,`commit`.`checksum`,`commit`.`repo_id`,`commit`.`sha`,`commit`.`branch`,`commit`.`message`,`commit`.`mergecommit`,`commit`.`excluded`,`commit`.`parent`,`commit`.`parent_id`,`commit`.`date`,`commit`.`author_user_id`,`commit`.`committer_user_id`,`commit`.`ordinal`,`commit`.`customer_id`,`commit`.`ref_type`,`commit`.`ref_id`,`commit`.`metadata` FROM `commit` WHERE `ref_type` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*Commit, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _RepoID sql.NullString
		var _Sha sql.NullString
		var _Branch sql.NullString
		var _Message sql.NullString
		var _Mergecommit sql.NullBool
		var _Excluded sql.NullBool
		var _Parent sql.NullString
		var _ParentID sql.NullString
		var _Date sql.NullInt64
		var _AuthorUserID sql.NullString
		var _CommitterUserID sql.NullString
		var _Ordinal sql.NullInt64
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefID sql.NullString
		var _Metadata sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_RepoID,
			&_Sha,
			&_Branch,
			&_Message,
			&_Mergecommit,
			&_Excluded,
			&_Parent,
			&_ParentID,
			&_Date,
			&_AuthorUserID,
			&_CommitterUserID,
			&_Ordinal,
			&_CustomerID,
			&_RefType,
			&_RefID,
			&_Metadata,
		)
		if err != nil {
			return nil, err
		}
		t := &Commit{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _RepoID.Valid {
			t.SetRepoID(_RepoID.String)
		}
		if _Sha.Valid {
			t.SetSha(_Sha.String)
		}
		if _Branch.Valid {
			t.SetBranch(_Branch.String)
		}
		if _Message.Valid {
			t.SetMessage(_Message.String)
		}
		if _Mergecommit.Valid {
			t.SetMergecommit(_Mergecommit.Bool)
		}
		if _Excluded.Valid {
			t.SetExcluded(_Excluded.Bool)
		}
		if _Parent.Valid {
			t.SetParent(_Parent.String)
		}
		if _ParentID.Valid {
			t.SetParentID(_ParentID.String)
		}
		if _Date.Valid {
			t.SetDate(_Date.Int64)
		}
		if _AuthorUserID.Valid {
			t.SetAuthorUserID(_AuthorUserID.String)
		}
		if _CommitterUserID.Valid {
			t.SetCommitterUserID(_CommitterUserID.String)
		}
		if _Ordinal.Valid {
			t.SetOrdinal(int32(_Ordinal.Int64))
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefType.Valid {
			t.SetRefType(_RefType.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		if _Metadata.Valid {
			t.SetMetadata(_Metadata.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetRefID will return the Commit RefID value
func (t *Commit) GetRefID() string {
	return t.RefID
}

// SetRefID will set the Commit RefID value
func (t *Commit) SetRefID(v string) {
	t.RefID = v
}

// FindCommitsByRefID will find all Commits by the RefID value
func FindCommitsByRefID(ctx context.Context, db DB, value string) ([]*Commit, error) {
	q := "SELECT `commit`.`id`,`commit`.`checksum`,`commit`.`repo_id`,`commit`.`sha`,`commit`.`branch`,`commit`.`message`,`commit`.`mergecommit`,`commit`.`excluded`,`commit`.`parent`,`commit`.`parent_id`,`commit`.`date`,`commit`.`author_user_id`,`commit`.`committer_user_id`,`commit`.`ordinal`,`commit`.`customer_id`,`commit`.`ref_type`,`commit`.`ref_id`,`commit`.`metadata` FROM `commit` WHERE `ref_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*Commit, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _RepoID sql.NullString
		var _Sha sql.NullString
		var _Branch sql.NullString
		var _Message sql.NullString
		var _Mergecommit sql.NullBool
		var _Excluded sql.NullBool
		var _Parent sql.NullString
		var _ParentID sql.NullString
		var _Date sql.NullInt64
		var _AuthorUserID sql.NullString
		var _CommitterUserID sql.NullString
		var _Ordinal sql.NullInt64
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefID sql.NullString
		var _Metadata sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_RepoID,
			&_Sha,
			&_Branch,
			&_Message,
			&_Mergecommit,
			&_Excluded,
			&_Parent,
			&_ParentID,
			&_Date,
			&_AuthorUserID,
			&_CommitterUserID,
			&_Ordinal,
			&_CustomerID,
			&_RefType,
			&_RefID,
			&_Metadata,
		)
		if err != nil {
			return nil, err
		}
		t := &Commit{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _RepoID.Valid {
			t.SetRepoID(_RepoID.String)
		}
		if _Sha.Valid {
			t.SetSha(_Sha.String)
		}
		if _Branch.Valid {
			t.SetBranch(_Branch.String)
		}
		if _Message.Valid {
			t.SetMessage(_Message.String)
		}
		if _Mergecommit.Valid {
			t.SetMergecommit(_Mergecommit.Bool)
		}
		if _Excluded.Valid {
			t.SetExcluded(_Excluded.Bool)
		}
		if _Parent.Valid {
			t.SetParent(_Parent.String)
		}
		if _ParentID.Valid {
			t.SetParentID(_ParentID.String)
		}
		if _Date.Valid {
			t.SetDate(_Date.Int64)
		}
		if _AuthorUserID.Valid {
			t.SetAuthorUserID(_AuthorUserID.String)
		}
		if _CommitterUserID.Valid {
			t.SetCommitterUserID(_CommitterUserID.String)
		}
		if _Ordinal.Valid {
			t.SetOrdinal(int32(_Ordinal.Int64))
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefType.Valid {
			t.SetRefType(_RefType.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		if _Metadata.Valid {
			t.SetMetadata(_Metadata.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindCommitsByRefIDTx will find all Commits by the RefID value using the provided transaction
func FindCommitsByRefIDTx(ctx context.Context, tx Tx, value string) ([]*Commit, error) {
	q := "SELECT `commit`.`id`,`commit`.`checksum`,`commit`.`repo_id`,`commit`.`sha`,`commit`.`branch`,`commit`.`message`,`commit`.`mergecommit`,`commit`.`excluded`,`commit`.`parent`,`commit`.`parent_id`,`commit`.`date`,`commit`.`author_user_id`,`commit`.`committer_user_id`,`commit`.`ordinal`,`commit`.`customer_id`,`commit`.`ref_type`,`commit`.`ref_id`,`commit`.`metadata` FROM `commit` WHERE `ref_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*Commit, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _RepoID sql.NullString
		var _Sha sql.NullString
		var _Branch sql.NullString
		var _Message sql.NullString
		var _Mergecommit sql.NullBool
		var _Excluded sql.NullBool
		var _Parent sql.NullString
		var _ParentID sql.NullString
		var _Date sql.NullInt64
		var _AuthorUserID sql.NullString
		var _CommitterUserID sql.NullString
		var _Ordinal sql.NullInt64
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefID sql.NullString
		var _Metadata sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_RepoID,
			&_Sha,
			&_Branch,
			&_Message,
			&_Mergecommit,
			&_Excluded,
			&_Parent,
			&_ParentID,
			&_Date,
			&_AuthorUserID,
			&_CommitterUserID,
			&_Ordinal,
			&_CustomerID,
			&_RefType,
			&_RefID,
			&_Metadata,
		)
		if err != nil {
			return nil, err
		}
		t := &Commit{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _RepoID.Valid {
			t.SetRepoID(_RepoID.String)
		}
		if _Sha.Valid {
			t.SetSha(_Sha.String)
		}
		if _Branch.Valid {
			t.SetBranch(_Branch.String)
		}
		if _Message.Valid {
			t.SetMessage(_Message.String)
		}
		if _Mergecommit.Valid {
			t.SetMergecommit(_Mergecommit.Bool)
		}
		if _Excluded.Valid {
			t.SetExcluded(_Excluded.Bool)
		}
		if _Parent.Valid {
			t.SetParent(_Parent.String)
		}
		if _ParentID.Valid {
			t.SetParentID(_ParentID.String)
		}
		if _Date.Valid {
			t.SetDate(_Date.Int64)
		}
		if _AuthorUserID.Valid {
			t.SetAuthorUserID(_AuthorUserID.String)
		}
		if _CommitterUserID.Valid {
			t.SetCommitterUserID(_CommitterUserID.String)
		}
		if _Ordinal.Valid {
			t.SetOrdinal(int32(_Ordinal.Int64))
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefType.Valid {
			t.SetRefType(_RefType.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		if _Metadata.Valid {
			t.SetMetadata(_Metadata.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetMetadata will return the Commit Metadata value
func (t *Commit) GetMetadata() string {
	if t.Metadata == nil {
		return ""
	}
	return *t.Metadata
}

// SetMetadata will set the Commit Metadata value
func (t *Commit) SetMetadata(v string) {
	t.Metadata = &v
}

func (t *Commit) toTimestamp(value time.Time) *timestamp.Timestamp {
	ts, _ := ptypes.TimestampProto(value)
	return ts
}

// DBCreateCommitTable will create the Commit table
func DBCreateCommitTable(ctx context.Context, db DB) error {
	q := "CREATE TABLE `commit` (`id` VARCHAR(64) NOT NULL PRIMARY KEY,`checksum` CHAR(64),`repo_id` VARCHAR(64) NOT NULL,`sha`VARCHAR(64) NOT NULL,`branch`VARCHAR(255) DEFAULT \"master\",`message` LONGTEXT,`mergecommit` BOOL DEFAULT false,`excluded` BOOL DEFAULT false,`parent`VARCHAR(64),`parent_id`VARCHAR(64),`date` BIGINT UNSIGNED NOT NULL,`author_user_id` VARCHAR(64) NOT NULL,`committer_user_id` VARCHAR(64) NOT NULL,`ordinal` INT,`customer_id` VARCHAR(64) NOT NULL,`ref_type` VARCHAR(20) NOT NULL,`ref_id`VARCHAR(64) NOT NULL,`metadata` JSON,INDEX commit_repo_id_index (`repo_id`),INDEX commit_sha_index (`sha`),INDEX commit_author_user_id_index (`author_user_id`),INDEX commit_committer_user_id_index (`committer_user_id`),INDEX commit_customer_id_index (`customer_id`),INDEX commit_ref_type_index (`ref_type`),INDEX commit_ref_id_index (`ref_id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBCreateCommitTableTx will create the Commit table using the provided transction
func DBCreateCommitTableTx(ctx context.Context, tx Tx) error {
	q := "CREATE TABLE `commit` (`id` VARCHAR(64) NOT NULL PRIMARY KEY,`checksum` CHAR(64),`repo_id` VARCHAR(64) NOT NULL,`sha`VARCHAR(64) NOT NULL,`branch`VARCHAR(255) DEFAULT \"master\",`message` LONGTEXT,`mergecommit` BOOL DEFAULT false,`excluded` BOOL DEFAULT false,`parent`VARCHAR(64),`parent_id`VARCHAR(64),`date` BIGINT UNSIGNED NOT NULL,`author_user_id` VARCHAR(64) NOT NULL,`committer_user_id` VARCHAR(64) NOT NULL,`ordinal` INT,`customer_id` VARCHAR(64) NOT NULL,`ref_type` VARCHAR(20) NOT NULL,`ref_id`VARCHAR(64) NOT NULL,`metadata` JSON,INDEX commit_repo_id_index (`repo_id`),INDEX commit_sha_index (`sha`),INDEX commit_author_user_id_index (`author_user_id`),INDEX commit_committer_user_id_index (`committer_user_id`),INDEX commit_customer_id_index (`customer_id`),INDEX commit_ref_type_index (`ref_type`),INDEX commit_ref_id_index (`ref_id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// DBDropCommitTable will drop the Commit table
func DBDropCommitTable(ctx context.Context, db DB) error {
	q := "DROP TABLE IF EXISTS `commit`"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBDropCommitTableTx will drop the Commit table using the provided transaction
func DBDropCommitTableTx(ctx context.Context, tx Tx) error {
	q := "DROP TABLE IF EXISTS `commit`"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// CalculateChecksum will calculate a checksum of the SHA1 of all field values
func (t *Commit) CalculateChecksum() string {
	return orm.HashStrings(
		orm.ToString(t.ID),
		orm.ToString(t.RepoID),
		orm.ToString(t.Sha),
		orm.ToString(t.Branch),
		orm.ToString(t.Message),
		orm.ToString(t.Mergecommit),
		orm.ToString(t.Excluded),
		orm.ToString(t.Parent),
		orm.ToString(t.ParentID),
		orm.ToString(t.Date),
		orm.ToString(t.AuthorUserID),
		orm.ToString(t.CommitterUserID),
		orm.ToString(t.Ordinal),
		orm.ToString(t.CustomerID),
		orm.ToString(t.RefType),
		orm.ToString(t.RefID),
		orm.ToString(t.Metadata),
	)
}

// DBCreate will create a new Commit record in the database
func (t *Commit) DBCreate(ctx context.Context, db DB) (sql.Result, error) {
	q := "INSERT INTO `commit` (`commit`.`id`,`commit`.`checksum`,`commit`.`repo_id`,`commit`.`sha`,`commit`.`branch`,`commit`.`message`,`commit`.`mergecommit`,`commit`.`excluded`,`commit`.`parent`,`commit`.`parent_id`,`commit`.`date`,`commit`.`author_user_id`,`commit`.`committer_user_id`,`commit`.`ordinal`,`commit`.`customer_id`,`commit`.`ref_type`,`commit`.`ref_id`,`commit`.`metadata`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.RepoID),
		orm.ToSQLString(t.Sha),
		orm.ToSQLString(t.Branch),
		orm.ToSQLString(t.Message),
		orm.ToSQLBool(t.Mergecommit),
		orm.ToSQLBool(t.Excluded),
		orm.ToSQLString(t.Parent),
		orm.ToSQLString(t.ParentID),
		orm.ToSQLInt64(t.Date),
		orm.ToSQLString(t.AuthorUserID),
		orm.ToSQLString(t.CommitterUserID),
		orm.ToSQLInt64(t.Ordinal),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefType),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.Metadata),
	)
}

// DBCreateTx will create a new Commit record in the database using the provided transaction
func (t *Commit) DBCreateTx(ctx context.Context, tx Tx) (sql.Result, error) {
	q := "INSERT INTO `commit` (`commit`.`id`,`commit`.`checksum`,`commit`.`repo_id`,`commit`.`sha`,`commit`.`branch`,`commit`.`message`,`commit`.`mergecommit`,`commit`.`excluded`,`commit`.`parent`,`commit`.`parent_id`,`commit`.`date`,`commit`.`author_user_id`,`commit`.`committer_user_id`,`commit`.`ordinal`,`commit`.`customer_id`,`commit`.`ref_type`,`commit`.`ref_id`,`commit`.`metadata`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.RepoID),
		orm.ToSQLString(t.Sha),
		orm.ToSQLString(t.Branch),
		orm.ToSQLString(t.Message),
		orm.ToSQLBool(t.Mergecommit),
		orm.ToSQLBool(t.Excluded),
		orm.ToSQLString(t.Parent),
		orm.ToSQLString(t.ParentID),
		orm.ToSQLInt64(t.Date),
		orm.ToSQLString(t.AuthorUserID),
		orm.ToSQLString(t.CommitterUserID),
		orm.ToSQLInt64(t.Ordinal),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefType),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.Metadata),
	)
}

// DBCreateIgnoreDuplicate will upsert the Commit record in the database
func (t *Commit) DBCreateIgnoreDuplicate(ctx context.Context, db DB) (sql.Result, error) {
	q := "INSERT INTO `commit` (`commit`.`id`,`commit`.`checksum`,`commit`.`repo_id`,`commit`.`sha`,`commit`.`branch`,`commit`.`message`,`commit`.`mergecommit`,`commit`.`excluded`,`commit`.`parent`,`commit`.`parent_id`,`commit`.`date`,`commit`.`author_user_id`,`commit`.`committer_user_id`,`commit`.`ordinal`,`commit`.`customer_id`,`commit`.`ref_type`,`commit`.`ref_id`,`commit`.`metadata`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.RepoID),
		orm.ToSQLString(t.Sha),
		orm.ToSQLString(t.Branch),
		orm.ToSQLString(t.Message),
		orm.ToSQLBool(t.Mergecommit),
		orm.ToSQLBool(t.Excluded),
		orm.ToSQLString(t.Parent),
		orm.ToSQLString(t.ParentID),
		orm.ToSQLInt64(t.Date),
		orm.ToSQLString(t.AuthorUserID),
		orm.ToSQLString(t.CommitterUserID),
		orm.ToSQLInt64(t.Ordinal),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefType),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.Metadata),
	)
}

// DBCreateIgnoreDuplicateTx will upsert the Commit record in the database using the provided transaction
func (t *Commit) DBCreateIgnoreDuplicateTx(ctx context.Context, tx Tx) (sql.Result, error) {
	q := "INSERT INTO `commit` (`commit`.`id`,`commit`.`checksum`,`commit`.`repo_id`,`commit`.`sha`,`commit`.`branch`,`commit`.`message`,`commit`.`mergecommit`,`commit`.`excluded`,`commit`.`parent`,`commit`.`parent_id`,`commit`.`date`,`commit`.`author_user_id`,`commit`.`committer_user_id`,`commit`.`ordinal`,`commit`.`customer_id`,`commit`.`ref_type`,`commit`.`ref_id`,`commit`.`metadata`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.RepoID),
		orm.ToSQLString(t.Sha),
		orm.ToSQLString(t.Branch),
		orm.ToSQLString(t.Message),
		orm.ToSQLBool(t.Mergecommit),
		orm.ToSQLBool(t.Excluded),
		orm.ToSQLString(t.Parent),
		orm.ToSQLString(t.ParentID),
		orm.ToSQLInt64(t.Date),
		orm.ToSQLString(t.AuthorUserID),
		orm.ToSQLString(t.CommitterUserID),
		orm.ToSQLInt64(t.Ordinal),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefType),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.Metadata),
	)
}

// DeleteAllCommits deletes all Commit records in the database with optional filters
func DeleteAllCommits(ctx context.Context, db DB, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(CommitTableName),
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

// DeleteAllCommitsTx deletes all Commit records in the database with optional filters using the provided transaction
func DeleteAllCommitsTx(ctx context.Context, tx Tx, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(CommitTableName),
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

// DBDelete will delete this Commit record in the database
func (t *Commit) DBDelete(ctx context.Context, db DB) (bool, error) {
	q := "DELETE FROM `commit` WHERE `id` = ?"
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

// DBDeleteTx will delete this Commit record in the database using the provided transaction
func (t *Commit) DBDeleteTx(ctx context.Context, tx Tx) (bool, error) {
	q := "DELETE FROM `commit` WHERE `id` = ?"
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

// DBUpdate will update the Commit record in the database
func (t *Commit) DBUpdate(ctx context.Context, db DB) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `commit` SET `checksum`=?,`repo_id`=?,`sha`=?,`branch`=?,`message`=?,`mergecommit`=?,`excluded`=?,`parent`=?,`parent_id`=?,`date`=?,`author_user_id`=?,`committer_user_id`=?,`ordinal`=?,`customer_id`=?,`ref_type`=?,`ref_id`=?,`metadata`=? WHERE `id`=?"
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.RepoID),
		orm.ToSQLString(t.Sha),
		orm.ToSQLString(t.Branch),
		orm.ToSQLString(t.Message),
		orm.ToSQLBool(t.Mergecommit),
		orm.ToSQLBool(t.Excluded),
		orm.ToSQLString(t.Parent),
		orm.ToSQLString(t.ParentID),
		orm.ToSQLInt64(t.Date),
		orm.ToSQLString(t.AuthorUserID),
		orm.ToSQLString(t.CommitterUserID),
		orm.ToSQLInt64(t.Ordinal),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefType),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.Metadata),
		orm.ToSQLString(t.ID),
	)
}

// DBUpdateTx will update the Commit record in the database using the provided transaction
func (t *Commit) DBUpdateTx(ctx context.Context, tx Tx) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `commit` SET `checksum`=?,`repo_id`=?,`sha`=?,`branch`=?,`message`=?,`mergecommit`=?,`excluded`=?,`parent`=?,`parent_id`=?,`date`=?,`author_user_id`=?,`committer_user_id`=?,`ordinal`=?,`customer_id`=?,`ref_type`=?,`ref_id`=?,`metadata`=? WHERE `id`=?"
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.RepoID),
		orm.ToSQLString(t.Sha),
		orm.ToSQLString(t.Branch),
		orm.ToSQLString(t.Message),
		orm.ToSQLBool(t.Mergecommit),
		orm.ToSQLBool(t.Excluded),
		orm.ToSQLString(t.Parent),
		orm.ToSQLString(t.ParentID),
		orm.ToSQLInt64(t.Date),
		orm.ToSQLString(t.AuthorUserID),
		orm.ToSQLString(t.CommitterUserID),
		orm.ToSQLInt64(t.Ordinal),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefType),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.Metadata),
		orm.ToSQLString(t.ID),
	)
}

// DBUpsert will upsert the Commit record in the database
func (t *Commit) DBUpsert(ctx context.Context, db DB, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `commit` (`commit`.`id`,`commit`.`checksum`,`commit`.`repo_id`,`commit`.`sha`,`commit`.`branch`,`commit`.`message`,`commit`.`mergecommit`,`commit`.`excluded`,`commit`.`parent`,`commit`.`parent_id`,`commit`.`date`,`commit`.`author_user_id`,`commit`.`committer_user_id`,`commit`.`ordinal`,`commit`.`customer_id`,`commit`.`ref_type`,`commit`.`ref_id`,`commit`.`metadata`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `commit` (`commit`.`id`,`commit`.`checksum`,`commit`.`repo_id`,`commit`.`sha`,`commit`.`branch`,`commit`.`message`,`commit`.`mergecommit`,`commit`.`excluded`,`commit`.`parent`,`commit`.`parent_id`,`commit`.`date`,`commit`.`author_user_id`,`commit`.`committer_user_id`,`commit`.`ordinal`,`commit`.`customer_id`,`commit`.`ref_type`,`commit`.`ref_id`,`commit`.`metadata`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`repo_id`=VALUES(`repo_id`),`sha`=VALUES(`sha`),`branch`=VALUES(`branch`),`message`=VALUES(`message`),`mergecommit`=VALUES(`mergecommit`),`excluded`=VALUES(`excluded`),`parent`=VALUES(`parent`),`parent_id`=VALUES(`parent_id`),`date`=VALUES(`date`),`author_user_id`=VALUES(`author_user_id`),`committer_user_id`=VALUES(`committer_user_id`),`ordinal`=VALUES(`ordinal`),`customer_id`=VALUES(`customer_id`),`ref_type`=VALUES(`ref_type`),`ref_id`=VALUES(`ref_id`),`metadata`=VALUES(`metadata`)"
	}
	r, err := db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.RepoID),
		orm.ToSQLString(t.Sha),
		orm.ToSQLString(t.Branch),
		orm.ToSQLString(t.Message),
		orm.ToSQLBool(t.Mergecommit),
		orm.ToSQLBool(t.Excluded),
		orm.ToSQLString(t.Parent),
		orm.ToSQLString(t.ParentID),
		orm.ToSQLInt64(t.Date),
		orm.ToSQLString(t.AuthorUserID),
		orm.ToSQLString(t.CommitterUserID),
		orm.ToSQLInt64(t.Ordinal),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefType),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.Metadata),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBUpsertTx will upsert the Commit record in the database using the provided transaction
func (t *Commit) DBUpsertTx(ctx context.Context, tx Tx, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `commit` (`commit`.`id`,`commit`.`checksum`,`commit`.`repo_id`,`commit`.`sha`,`commit`.`branch`,`commit`.`message`,`commit`.`mergecommit`,`commit`.`excluded`,`commit`.`parent`,`commit`.`parent_id`,`commit`.`date`,`commit`.`author_user_id`,`commit`.`committer_user_id`,`commit`.`ordinal`,`commit`.`customer_id`,`commit`.`ref_type`,`commit`.`ref_id`,`commit`.`metadata`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `commit` (`commit`.`id`,`commit`.`checksum`,`commit`.`repo_id`,`commit`.`sha`,`commit`.`branch`,`commit`.`message`,`commit`.`mergecommit`,`commit`.`excluded`,`commit`.`parent`,`commit`.`parent_id`,`commit`.`date`,`commit`.`author_user_id`,`commit`.`committer_user_id`,`commit`.`ordinal`,`commit`.`customer_id`,`commit`.`ref_type`,`commit`.`ref_id`,`commit`.`metadata`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`repo_id`=VALUES(`repo_id`),`sha`=VALUES(`sha`),`branch`=VALUES(`branch`),`message`=VALUES(`message`),`mergecommit`=VALUES(`mergecommit`),`excluded`=VALUES(`excluded`),`parent`=VALUES(`parent`),`parent_id`=VALUES(`parent_id`),`date`=VALUES(`date`),`author_user_id`=VALUES(`author_user_id`),`committer_user_id`=VALUES(`committer_user_id`),`ordinal`=VALUES(`ordinal`),`customer_id`=VALUES(`customer_id`),`ref_type`=VALUES(`ref_type`),`ref_id`=VALUES(`ref_id`),`metadata`=VALUES(`metadata`)"
	}
	r, err := tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.RepoID),
		orm.ToSQLString(t.Sha),
		orm.ToSQLString(t.Branch),
		orm.ToSQLString(t.Message),
		orm.ToSQLBool(t.Mergecommit),
		orm.ToSQLBool(t.Excluded),
		orm.ToSQLString(t.Parent),
		orm.ToSQLString(t.ParentID),
		orm.ToSQLInt64(t.Date),
		orm.ToSQLString(t.AuthorUserID),
		orm.ToSQLString(t.CommitterUserID),
		orm.ToSQLInt64(t.Ordinal),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefType),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.Metadata),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBFindOne will find a Commit record in the database with the primary key
func (t *Commit) DBFindOne(ctx context.Context, db DB, value string) (bool, error) {
	q := "SELECT `commit`.`id`,`commit`.`checksum`,`commit`.`repo_id`,`commit`.`sha`,`commit`.`branch`,`commit`.`message`,`commit`.`mergecommit`,`commit`.`excluded`,`commit`.`parent`,`commit`.`parent_id`,`commit`.`date`,`commit`.`author_user_id`,`commit`.`committer_user_id`,`commit`.`ordinal`,`commit`.`customer_id`,`commit`.`ref_type`,`commit`.`ref_id`,`commit`.`metadata` FROM `commit` WHERE `id` = ? LIMIT 1"
	row := db.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _RepoID sql.NullString
	var _Sha sql.NullString
	var _Branch sql.NullString
	var _Message sql.NullString
	var _Mergecommit sql.NullBool
	var _Excluded sql.NullBool
	var _Parent sql.NullString
	var _ParentID sql.NullString
	var _Date sql.NullInt64
	var _AuthorUserID sql.NullString
	var _CommitterUserID sql.NullString
	var _Ordinal sql.NullInt64
	var _CustomerID sql.NullString
	var _RefType sql.NullString
	var _RefID sql.NullString
	var _Metadata sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_RepoID,
		&_Sha,
		&_Branch,
		&_Message,
		&_Mergecommit,
		&_Excluded,
		&_Parent,
		&_ParentID,
		&_Date,
		&_AuthorUserID,
		&_CommitterUserID,
		&_Ordinal,
		&_CustomerID,
		&_RefType,
		&_RefID,
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
	if _RepoID.Valid {
		t.SetRepoID(_RepoID.String)
	}
	if _Sha.Valid {
		t.SetSha(_Sha.String)
	}
	if _Branch.Valid {
		t.SetBranch(_Branch.String)
	}
	if _Message.Valid {
		t.SetMessage(_Message.String)
	}
	if _Mergecommit.Valid {
		t.SetMergecommit(_Mergecommit.Bool)
	}
	if _Excluded.Valid {
		t.SetExcluded(_Excluded.Bool)
	}
	if _Parent.Valid {
		t.SetParent(_Parent.String)
	}
	if _ParentID.Valid {
		t.SetParentID(_ParentID.String)
	}
	if _Date.Valid {
		t.SetDate(_Date.Int64)
	}
	if _AuthorUserID.Valid {
		t.SetAuthorUserID(_AuthorUserID.String)
	}
	if _CommitterUserID.Valid {
		t.SetCommitterUserID(_CommitterUserID.String)
	}
	if _Ordinal.Valid {
		t.SetOrdinal(int32(_Ordinal.Int64))
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _RefType.Valid {
		t.SetRefType(_RefType.String)
	}
	if _RefID.Valid {
		t.SetRefID(_RefID.String)
	}
	if _Metadata.Valid {
		t.SetMetadata(_Metadata.String)
	}
	return true, nil
}

// DBFindOneTx will find a Commit record in the database with the primary key using the provided transaction
func (t *Commit) DBFindOneTx(ctx context.Context, tx Tx, value string) (bool, error) {
	q := "SELECT `commit`.`id`,`commit`.`checksum`,`commit`.`repo_id`,`commit`.`sha`,`commit`.`branch`,`commit`.`message`,`commit`.`mergecommit`,`commit`.`excluded`,`commit`.`parent`,`commit`.`parent_id`,`commit`.`date`,`commit`.`author_user_id`,`commit`.`committer_user_id`,`commit`.`ordinal`,`commit`.`customer_id`,`commit`.`ref_type`,`commit`.`ref_id`,`commit`.`metadata` FROM `commit` WHERE `id` = ? LIMIT 1"
	row := tx.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _RepoID sql.NullString
	var _Sha sql.NullString
	var _Branch sql.NullString
	var _Message sql.NullString
	var _Mergecommit sql.NullBool
	var _Excluded sql.NullBool
	var _Parent sql.NullString
	var _ParentID sql.NullString
	var _Date sql.NullInt64
	var _AuthorUserID sql.NullString
	var _CommitterUserID sql.NullString
	var _Ordinal sql.NullInt64
	var _CustomerID sql.NullString
	var _RefType sql.NullString
	var _RefID sql.NullString
	var _Metadata sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_RepoID,
		&_Sha,
		&_Branch,
		&_Message,
		&_Mergecommit,
		&_Excluded,
		&_Parent,
		&_ParentID,
		&_Date,
		&_AuthorUserID,
		&_CommitterUserID,
		&_Ordinal,
		&_CustomerID,
		&_RefType,
		&_RefID,
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
	if _RepoID.Valid {
		t.SetRepoID(_RepoID.String)
	}
	if _Sha.Valid {
		t.SetSha(_Sha.String)
	}
	if _Branch.Valid {
		t.SetBranch(_Branch.String)
	}
	if _Message.Valid {
		t.SetMessage(_Message.String)
	}
	if _Mergecommit.Valid {
		t.SetMergecommit(_Mergecommit.Bool)
	}
	if _Excluded.Valid {
		t.SetExcluded(_Excluded.Bool)
	}
	if _Parent.Valid {
		t.SetParent(_Parent.String)
	}
	if _ParentID.Valid {
		t.SetParentID(_ParentID.String)
	}
	if _Date.Valid {
		t.SetDate(_Date.Int64)
	}
	if _AuthorUserID.Valid {
		t.SetAuthorUserID(_AuthorUserID.String)
	}
	if _CommitterUserID.Valid {
		t.SetCommitterUserID(_CommitterUserID.String)
	}
	if _Ordinal.Valid {
		t.SetOrdinal(int32(_Ordinal.Int64))
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _RefType.Valid {
		t.SetRefType(_RefType.String)
	}
	if _RefID.Valid {
		t.SetRefID(_RefID.String)
	}
	if _Metadata.Valid {
		t.SetMetadata(_Metadata.String)
	}
	return true, nil
}

// FindCommits will find a Commit record in the database with the provided parameters
func FindCommits(ctx context.Context, db DB, _params ...interface{}) ([]*Commit, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("repo_id"),
		orm.Column("sha"),
		orm.Column("branch"),
		orm.Column("message"),
		orm.Column("mergecommit"),
		orm.Column("excluded"),
		orm.Column("parent"),
		orm.Column("parent_id"),
		orm.Column("date"),
		orm.Column("author_user_id"),
		orm.Column("committer_user_id"),
		orm.Column("ordinal"),
		orm.Column("customer_id"),
		orm.Column("ref_type"),
		orm.Column("ref_id"),
		orm.Column("metadata"),
		orm.Table(CommitTableName),
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
	results := make([]*Commit, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _RepoID sql.NullString
		var _Sha sql.NullString
		var _Branch sql.NullString
		var _Message sql.NullString
		var _Mergecommit sql.NullBool
		var _Excluded sql.NullBool
		var _Parent sql.NullString
		var _ParentID sql.NullString
		var _Date sql.NullInt64
		var _AuthorUserID sql.NullString
		var _CommitterUserID sql.NullString
		var _Ordinal sql.NullInt64
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefID sql.NullString
		var _Metadata sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_RepoID,
			&_Sha,
			&_Branch,
			&_Message,
			&_Mergecommit,
			&_Excluded,
			&_Parent,
			&_ParentID,
			&_Date,
			&_AuthorUserID,
			&_CommitterUserID,
			&_Ordinal,
			&_CustomerID,
			&_RefType,
			&_RefID,
			&_Metadata,
		)
		if err != nil {
			return nil, err
		}
		t := &Commit{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _RepoID.Valid {
			t.SetRepoID(_RepoID.String)
		}
		if _Sha.Valid {
			t.SetSha(_Sha.String)
		}
		if _Branch.Valid {
			t.SetBranch(_Branch.String)
		}
		if _Message.Valid {
			t.SetMessage(_Message.String)
		}
		if _Mergecommit.Valid {
			t.SetMergecommit(_Mergecommit.Bool)
		}
		if _Excluded.Valid {
			t.SetExcluded(_Excluded.Bool)
		}
		if _Parent.Valid {
			t.SetParent(_Parent.String)
		}
		if _ParentID.Valid {
			t.SetParentID(_ParentID.String)
		}
		if _Date.Valid {
			t.SetDate(_Date.Int64)
		}
		if _AuthorUserID.Valid {
			t.SetAuthorUserID(_AuthorUserID.String)
		}
		if _CommitterUserID.Valid {
			t.SetCommitterUserID(_CommitterUserID.String)
		}
		if _Ordinal.Valid {
			t.SetOrdinal(int32(_Ordinal.Int64))
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefType.Valid {
			t.SetRefType(_RefType.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		if _Metadata.Valid {
			t.SetMetadata(_Metadata.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindCommitsTx will find a Commit record in the database with the provided parameters using the provided transaction
func FindCommitsTx(ctx context.Context, tx Tx, _params ...interface{}) ([]*Commit, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("repo_id"),
		orm.Column("sha"),
		orm.Column("branch"),
		orm.Column("message"),
		orm.Column("mergecommit"),
		orm.Column("excluded"),
		orm.Column("parent"),
		orm.Column("parent_id"),
		orm.Column("date"),
		orm.Column("author_user_id"),
		orm.Column("committer_user_id"),
		orm.Column("ordinal"),
		orm.Column("customer_id"),
		orm.Column("ref_type"),
		orm.Column("ref_id"),
		orm.Column("metadata"),
		orm.Table(CommitTableName),
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
	results := make([]*Commit, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _RepoID sql.NullString
		var _Sha sql.NullString
		var _Branch sql.NullString
		var _Message sql.NullString
		var _Mergecommit sql.NullBool
		var _Excluded sql.NullBool
		var _Parent sql.NullString
		var _ParentID sql.NullString
		var _Date sql.NullInt64
		var _AuthorUserID sql.NullString
		var _CommitterUserID sql.NullString
		var _Ordinal sql.NullInt64
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefID sql.NullString
		var _Metadata sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_RepoID,
			&_Sha,
			&_Branch,
			&_Message,
			&_Mergecommit,
			&_Excluded,
			&_Parent,
			&_ParentID,
			&_Date,
			&_AuthorUserID,
			&_CommitterUserID,
			&_Ordinal,
			&_CustomerID,
			&_RefType,
			&_RefID,
			&_Metadata,
		)
		if err != nil {
			return nil, err
		}
		t := &Commit{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _RepoID.Valid {
			t.SetRepoID(_RepoID.String)
		}
		if _Sha.Valid {
			t.SetSha(_Sha.String)
		}
		if _Branch.Valid {
			t.SetBranch(_Branch.String)
		}
		if _Message.Valid {
			t.SetMessage(_Message.String)
		}
		if _Mergecommit.Valid {
			t.SetMergecommit(_Mergecommit.Bool)
		}
		if _Excluded.Valid {
			t.SetExcluded(_Excluded.Bool)
		}
		if _Parent.Valid {
			t.SetParent(_Parent.String)
		}
		if _ParentID.Valid {
			t.SetParentID(_ParentID.String)
		}
		if _Date.Valid {
			t.SetDate(_Date.Int64)
		}
		if _AuthorUserID.Valid {
			t.SetAuthorUserID(_AuthorUserID.String)
		}
		if _CommitterUserID.Valid {
			t.SetCommitterUserID(_CommitterUserID.String)
		}
		if _Ordinal.Valid {
			t.SetOrdinal(int32(_Ordinal.Int64))
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefType.Valid {
			t.SetRefType(_RefType.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		if _Metadata.Valid {
			t.SetMetadata(_Metadata.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// DBFind will find a Commit record in the database with the provided parameters
func (t *Commit) DBFind(ctx context.Context, db DB, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("repo_id"),
		orm.Column("sha"),
		orm.Column("branch"),
		orm.Column("message"),
		orm.Column("mergecommit"),
		orm.Column("excluded"),
		orm.Column("parent"),
		orm.Column("parent_id"),
		orm.Column("date"),
		orm.Column("author_user_id"),
		orm.Column("committer_user_id"),
		orm.Column("ordinal"),
		orm.Column("customer_id"),
		orm.Column("ref_type"),
		orm.Column("ref_id"),
		orm.Column("metadata"),
		orm.Table(CommitTableName),
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
	var _RepoID sql.NullString
	var _Sha sql.NullString
	var _Branch sql.NullString
	var _Message sql.NullString
	var _Mergecommit sql.NullBool
	var _Excluded sql.NullBool
	var _Parent sql.NullString
	var _ParentID sql.NullString
	var _Date sql.NullInt64
	var _AuthorUserID sql.NullString
	var _CommitterUserID sql.NullString
	var _Ordinal sql.NullInt64
	var _CustomerID sql.NullString
	var _RefType sql.NullString
	var _RefID sql.NullString
	var _Metadata sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_RepoID,
		&_Sha,
		&_Branch,
		&_Message,
		&_Mergecommit,
		&_Excluded,
		&_Parent,
		&_ParentID,
		&_Date,
		&_AuthorUserID,
		&_CommitterUserID,
		&_Ordinal,
		&_CustomerID,
		&_RefType,
		&_RefID,
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
	if _RepoID.Valid {
		t.SetRepoID(_RepoID.String)
	}
	if _Sha.Valid {
		t.SetSha(_Sha.String)
	}
	if _Branch.Valid {
		t.SetBranch(_Branch.String)
	}
	if _Message.Valid {
		t.SetMessage(_Message.String)
	}
	if _Mergecommit.Valid {
		t.SetMergecommit(_Mergecommit.Bool)
	}
	if _Excluded.Valid {
		t.SetExcluded(_Excluded.Bool)
	}
	if _Parent.Valid {
		t.SetParent(_Parent.String)
	}
	if _ParentID.Valid {
		t.SetParentID(_ParentID.String)
	}
	if _Date.Valid {
		t.SetDate(_Date.Int64)
	}
	if _AuthorUserID.Valid {
		t.SetAuthorUserID(_AuthorUserID.String)
	}
	if _CommitterUserID.Valid {
		t.SetCommitterUserID(_CommitterUserID.String)
	}
	if _Ordinal.Valid {
		t.SetOrdinal(int32(_Ordinal.Int64))
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _RefType.Valid {
		t.SetRefType(_RefType.String)
	}
	if _RefID.Valid {
		t.SetRefID(_RefID.String)
	}
	if _Metadata.Valid {
		t.SetMetadata(_Metadata.String)
	}
	return true, nil
}

// DBFindTx will find a Commit record in the database with the provided parameters using the provided transaction
func (t *Commit) DBFindTx(ctx context.Context, tx Tx, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("repo_id"),
		orm.Column("sha"),
		orm.Column("branch"),
		orm.Column("message"),
		orm.Column("mergecommit"),
		orm.Column("excluded"),
		orm.Column("parent"),
		orm.Column("parent_id"),
		orm.Column("date"),
		orm.Column("author_user_id"),
		orm.Column("committer_user_id"),
		orm.Column("ordinal"),
		orm.Column("customer_id"),
		orm.Column("ref_type"),
		orm.Column("ref_id"),
		orm.Column("metadata"),
		orm.Table(CommitTableName),
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
	var _RepoID sql.NullString
	var _Sha sql.NullString
	var _Branch sql.NullString
	var _Message sql.NullString
	var _Mergecommit sql.NullBool
	var _Excluded sql.NullBool
	var _Parent sql.NullString
	var _ParentID sql.NullString
	var _Date sql.NullInt64
	var _AuthorUserID sql.NullString
	var _CommitterUserID sql.NullString
	var _Ordinal sql.NullInt64
	var _CustomerID sql.NullString
	var _RefType sql.NullString
	var _RefID sql.NullString
	var _Metadata sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_RepoID,
		&_Sha,
		&_Branch,
		&_Message,
		&_Mergecommit,
		&_Excluded,
		&_Parent,
		&_ParentID,
		&_Date,
		&_AuthorUserID,
		&_CommitterUserID,
		&_Ordinal,
		&_CustomerID,
		&_RefType,
		&_RefID,
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
	if _RepoID.Valid {
		t.SetRepoID(_RepoID.String)
	}
	if _Sha.Valid {
		t.SetSha(_Sha.String)
	}
	if _Branch.Valid {
		t.SetBranch(_Branch.String)
	}
	if _Message.Valid {
		t.SetMessage(_Message.String)
	}
	if _Mergecommit.Valid {
		t.SetMergecommit(_Mergecommit.Bool)
	}
	if _Excluded.Valid {
		t.SetExcluded(_Excluded.Bool)
	}
	if _Parent.Valid {
		t.SetParent(_Parent.String)
	}
	if _ParentID.Valid {
		t.SetParentID(_ParentID.String)
	}
	if _Date.Valid {
		t.SetDate(_Date.Int64)
	}
	if _AuthorUserID.Valid {
		t.SetAuthorUserID(_AuthorUserID.String)
	}
	if _CommitterUserID.Valid {
		t.SetCommitterUserID(_CommitterUserID.String)
	}
	if _Ordinal.Valid {
		t.SetOrdinal(int32(_Ordinal.Int64))
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _RefType.Valid {
		t.SetRefType(_RefType.String)
	}
	if _RefID.Valid {
		t.SetRefID(_RefID.String)
	}
	if _Metadata.Valid {
		t.SetMetadata(_Metadata.String)
	}
	return true, nil
}

// CountCommits will find the count of Commit records in the database
func CountCommits(ctx context.Context, db DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(CommitTableName),
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

// CountCommitsTx will find the count of Commit records in the database using the provided transaction
func CountCommitsTx(ctx context.Context, tx Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(CommitTableName),
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

// DBCount will find the count of Commit records in the database
func (t *Commit) DBCount(ctx context.Context, db DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(CommitTableName),
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

// DBCountTx will find the count of Commit records in the database using the provided transaction
func (t *Commit) DBCountTx(ctx context.Context, tx Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(CommitTableName),
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

// DBExists will return true if the Commit record exists in the database
func (t *Commit) DBExists(ctx context.Context, db DB) (bool, error) {
	q := "SELECT `id` FROM `commit` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := db.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// DBExistsTx will return true if the Commit record exists in the database using the provided transaction
func (t *Commit) DBExistsTx(ctx context.Context, tx Tx) (bool, error) {
	q := "SELECT `id` FROM `commit` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := tx.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// PrimaryKeyColumn returns the column name for the primary key
func (t *Commit) PrimaryKeyColumn() string {
	return CommitColumnID
}

// PrimaryKeyColumnType returns the primary key column Go type as a string
func (t *Commit) PrimaryKeyColumnType() string {
	return "string"
}

// PrimaryKey returns the primary key column value
func (t *Commit) PrimaryKey() interface{} {
	return t.ID
}
