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

var _ Model = (*CommitFile)(nil)
var _ CSVWriter = (*CommitFile)(nil)
var _ JSONWriter = (*CommitFile)(nil)
var _ Checksum = (*CommitFile)(nil)

// CommitFileTableName is the name of the table in SQL
const CommitFileTableName = "commit_file"

var CommitFileColumns = []string{
	"id",
	"checksum",
	"commit_id",
	"repo_id",
	"author_user_id",
	"committer_user_id",
	"filename",
	"language",
	"additions",
	"deletions",
	"size",
	"abinary",
	"date",
	"branch",
	"mergecommit",
	"excluded",
	"loc",
	"sloc",
	"comments",
	"blanks",
	"variance",
	"status",
	"renamed",
	"renamed_from",
	"renamed_to",
	"customer_id",
	"ref_type",
	"ref_id",
	"metadata",
}

type CommitFile_CommitFileStatus string

const (
	CommitFileStatus_ADDED    CommitFile_CommitFileStatus = "added"
	CommitFileStatus_MODIFIED CommitFile_CommitFileStatus = "modified"
	CommitFileStatus_DELETED  CommitFile_CommitFileStatus = "deleted"
)

func (x CommitFile_CommitFileStatus) String() string {
	return string(x)
}

func enumCommitFileStatusToString(v *CommitFile_CommitFileStatus) string {
	if v == nil {
		return ""
	}
	return v.String()
}

func toCommitFileStatus(v string) *CommitFile_CommitFileStatus {
	var ev *CommitFile_CommitFileStatus
	switch v {
	case "ADDED", "added":
		{
			v := CommitFileStatus_ADDED
			ev = &v
		}
	case "MODIFIED", "modified":
		{
			v := CommitFileStatus_MODIFIED
			ev = &v
		}
	case "DELETED", "deleted":
		{
			v := CommitFileStatus_DELETED
			ev = &v
		}
	}
	return ev
}

// CommitFile table
type CommitFile struct {
	Abinary         bool                        `json:"abinary"`
	Additions       int32                       `json:"additions"`
	AuthorUserID    string                      `json:"author_user_id"`
	Blanks          int32                       `json:"blanks"`
	Branch          string                      `json:"branch"`
	Checksum        *string                     `json:"checksum,omitempty"`
	Comments        int32                       `json:"comments"`
	CommitID        string                      `json:"commit_id"`
	CommitterUserID string                      `json:"committer_user_id"`
	CustomerID      string                      `json:"customer_id"`
	Date            int64                       `json:"date"`
	Deletions       int32                       `json:"deletions"`
	Excluded        bool                        `json:"excluded"`
	Filename        string                      `json:"filename"`
	ID              string                      `json:"id"`
	Language        string                      `json:"language"`
	Loc             int32                       `json:"loc"`
	Mergecommit     bool                        `json:"mergecommit"`
	Metadata        *string                     `json:"metadata,omitempty"`
	RefID           string                      `json:"ref_id"`
	RefType         string                      `json:"ref_type"`
	Renamed         bool                        `json:"renamed"`
	RenamedFrom     *string                     `json:"renamed_from,omitempty"`
	RenamedTo       *string                     `json:"renamed_to,omitempty"`
	RepoID          string                      `json:"repo_id"`
	Size            int32                       `json:"size"`
	Sloc            int32                       `json:"sloc"`
	Status          CommitFile_CommitFileStatus `json:"status"`
	Variance        int32                       `json:"variance"`
}

// TableName returns the SQL table name for CommitFile and satifies the Model interface
func (t *CommitFile) TableName() string {
	return CommitFileTableName
}

// ToCSV will serialize the CommitFile instance to a CSV compatible array of strings
func (t *CommitFile) ToCSV() []string {
	return []string{
		t.ID,
		t.CalculateChecksum(),
		t.CommitID,
		t.RepoID,
		t.AuthorUserID,
		t.CommitterUserID,
		t.Filename,
		t.Language,
		toCSVString(t.Additions),
		toCSVString(t.Deletions),
		toCSVString(t.Size),
		toCSVBool(t.Abinary),
		toCSVString(t.Date),
		t.Branch,
		toCSVBool(t.Mergecommit),
		toCSVBool(t.Excluded),
		toCSVString(t.Loc),
		toCSVString(t.Sloc),
		toCSVString(t.Comments),
		toCSVString(t.Blanks),
		toCSVString(t.Variance),
		toCSVString(t.Status),
		toCSVBool(t.Renamed),
		toCSVString(t.RenamedFrom),
		toCSVString(t.RenamedTo),
		t.CustomerID,
		t.RefType,
		t.RefID,
		toCSVString(t.Metadata),
	}
}

// WriteCSV will serialize the CommitFile instance to the writer as CSV and satisfies the CSVWriter interface
func (t *CommitFile) WriteCSV(w *csv.Writer) error {
	return w.Write(t.ToCSV())
}

// WriteJSON will serialize the CommitFile instance to the writer as JSON and satisfies the JSONWriter interface
func (t *CommitFile) WriteJSON(w io.Writer, indent ...bool) error {
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

// NewCommitFileReader creates a JSON reader which can read in CommitFile objects serialized as JSON either as an array, single object or json new lines
// and writes each CommitFile to the channel provided
func NewCommitFileReader(r io.Reader, ch chan<- CommitFile) error {
	return orm.Deserialize(r, func(buf json.RawMessage) error {
		dec := json.NewDecoder(bytes.NewBuffer(buf))
		e := CommitFile{}
		if err := dec.Decode(&e); err != nil {
			return err
		}
		ch <- e
		return nil
	})
}

// NewCSVCommitFileReaderDir will read the reader as CSV and emit each record to the channel provided
func NewCSVCommitFileReader(r io.Reader, ch chan<- CommitFile) error {
	cr := csv.NewReader(r)
	for {
		record, err := cr.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		ch <- CommitFile{
			ID:              record[0],
			Checksum:        fromStringPointer(record[1]),
			CommitID:        record[2],
			RepoID:          record[3],
			AuthorUserID:    record[4],
			CommitterUserID: record[5],
			Filename:        record[6],
			Language:        record[7],
			Additions:       fromCSVInt32(record[8]),
			Deletions:       fromCSVInt32(record[9]),
			Size:            fromCSVInt32(record[10]),
			Abinary:         fromCSVBool(record[11]),
			Date:            fromCSVInt64(record[12]),
			Branch:          record[13],
			Mergecommit:     fromCSVBool(record[14]),
			Excluded:        fromCSVBool(record[15]),
			Loc:             fromCSVInt32(record[16]),
			Sloc:            fromCSVInt32(record[17]),
			Comments:        fromCSVInt32(record[18]),
			Blanks:          fromCSVInt32(record[19]),
			Variance:        fromCSVInt32(record[20]),
			Status:          *toCommitFileStatus(record[21]),
			Renamed:         fromCSVBool(record[22]),
			RenamedFrom:     fromStringPointer(record[23]),
			RenamedTo:       fromStringPointer(record[24]),
			CustomerID:      record[25],
			RefType:         record[26],
			RefID:           record[27],
			Metadata:        fromStringPointer(record[28]),
		}
	}
	return nil
}

// NewCSVCommitFileReaderFile will read the file as a CSV and emit each record to the channel provided
func NewCSVCommitFileReaderFile(fp string, ch chan<- CommitFile) error {
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
	return NewCSVCommitFileReader(fc, ch)
}

// NewCSVCommitFileReaderDir will read the commit_file.csv.gz file as a CSV and emit each record to the channel provided
func NewCSVCommitFileReaderDir(dir string, ch chan<- CommitFile) error {
	return NewCSVCommitFileReaderFile(filepath.Join(dir, "commit_file.csv.gz"), ch)
}

// CommitFileCSVDeduper is a function callback which takes the existing value (a) and the new value (b)
// and the return value should be the one to use (or a new one, if applicable). return nil
// to skip processing of this record
type CommitFileCSVDeduper func(a CommitFile, b CommitFile) *CommitFile

// CommitFileCSVDedupeDisabled is set on whether the CSV writer should de-dupe values by key
var CommitFileCSVDedupeDisabled bool

// NewCommitFileCSVWriterSize creates a batch writer that will write each CommitFile into a CSV file
// this method will automatically de-duplicate entries using the primary key. if the checksum
// for a newer item with the same primary key doesn't match a previously sent item, the newer
// one will be used
func NewCommitFileCSVWriterSize(w io.Writer, size int, dedupers ...CommitFileCSVDeduper) (chan CommitFile, chan bool, error) {
	cw := csv.NewWriter(w)
	ch := make(chan CommitFile, size)
	done := make(chan bool)
	go func() {
		defer func() { done <- true }()
		dodedupe := !CommitFileCSVDedupeDisabled
		var kv map[string]*CommitFile
		var deduper CommitFileCSVDeduper
		if dedupers != nil && len(dedupers) > 0 {
			deduper = dedupers[0]
			dodedupe = true
		}
		if dodedupe {
			kv = make(map[string]*CommitFile)
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

// CommitFileCSVDefaultSize is the default channel buffer size if not provided
var CommitFileCSVDefaultSize = 100

// NewCommitFileCSVWriter creates a batch writer that will write each CommitFile into a CSV file
func NewCommitFileCSVWriter(w io.Writer, dedupers ...CommitFileCSVDeduper) (chan CommitFile, chan bool, error) {
	return NewCommitFileCSVWriterSize(w, CommitFileCSVDefaultSize, dedupers...)
}

// NewCommitFileCSVWriterDir creates a batch writer that will write each CommitFile into a CSV file named commit_file.csv.gz in dir
func NewCommitFileCSVWriterDir(dir string, dedupers ...CommitFileCSVDeduper) (chan CommitFile, chan bool, error) {
	return NewCommitFileCSVWriterFile(filepath.Join(dir, "commit_file.csv.gz"), dedupers...)
}

// NewCommitFileCSVWriterFile creates a batch writer that will write each CommitFile into a CSV file
func NewCommitFileCSVWriterFile(fn string, dedupers ...CommitFileCSVDeduper) (chan CommitFile, chan bool, error) {
	f, err := os.Create(fn)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening CSV file %s. %v", fn, err)
	}
	var fc io.WriteCloser = f
	if filepath.Ext(fn) == ".gz" {
		w, _ := gzip.NewWriterLevel(f, gzip.BestCompression)
		fc = w
	}
	ch, done, err := NewCommitFileCSVWriter(fc, dedupers...)
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

type CommitFileDBAction func(ctx context.Context, db DB, record CommitFile) error

// NewCommitFileDBWriterSize creates a DB writer that will write each issue into the DB
func NewCommitFileDBWriterSize(ctx context.Context, db DB, errors chan<- error, size int, actions ...CommitFileDBAction) (chan CommitFile, chan bool, error) {
	ch := make(chan CommitFile, size)
	done := make(chan bool)
	var action CommitFileDBAction
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

// NewCommitFileDBWriter creates a DB writer that will write each issue into the DB
func NewCommitFileDBWriter(ctx context.Context, db DB, errors chan<- error, actions ...CommitFileDBAction) (chan CommitFile, chan bool, error) {
	return NewCommitFileDBWriterSize(ctx, db, errors, 100, actions...)
}

// CommitFileColumnID is the ID SQL column name for the CommitFile table
const CommitFileColumnID = "id"

// CommitFileEscapedColumnID is the escaped ID SQL column name for the CommitFile table
const CommitFileEscapedColumnID = "`id`"

// CommitFileColumnChecksum is the Checksum SQL column name for the CommitFile table
const CommitFileColumnChecksum = "checksum"

// CommitFileEscapedColumnChecksum is the escaped Checksum SQL column name for the CommitFile table
const CommitFileEscapedColumnChecksum = "`checksum`"

// CommitFileColumnCommitID is the CommitID SQL column name for the CommitFile table
const CommitFileColumnCommitID = "commit_id"

// CommitFileEscapedColumnCommitID is the escaped CommitID SQL column name for the CommitFile table
const CommitFileEscapedColumnCommitID = "`commit_id`"

// CommitFileColumnRepoID is the RepoID SQL column name for the CommitFile table
const CommitFileColumnRepoID = "repo_id"

// CommitFileEscapedColumnRepoID is the escaped RepoID SQL column name for the CommitFile table
const CommitFileEscapedColumnRepoID = "`repo_id`"

// CommitFileColumnAuthorUserID is the AuthorUserID SQL column name for the CommitFile table
const CommitFileColumnAuthorUserID = "author_user_id"

// CommitFileEscapedColumnAuthorUserID is the escaped AuthorUserID SQL column name for the CommitFile table
const CommitFileEscapedColumnAuthorUserID = "`author_user_id`"

// CommitFileColumnCommitterUserID is the CommitterUserID SQL column name for the CommitFile table
const CommitFileColumnCommitterUserID = "committer_user_id"

// CommitFileEscapedColumnCommitterUserID is the escaped CommitterUserID SQL column name for the CommitFile table
const CommitFileEscapedColumnCommitterUserID = "`committer_user_id`"

// CommitFileColumnFilename is the Filename SQL column name for the CommitFile table
const CommitFileColumnFilename = "filename"

// CommitFileEscapedColumnFilename is the escaped Filename SQL column name for the CommitFile table
const CommitFileEscapedColumnFilename = "`filename`"

// CommitFileColumnLanguage is the Language SQL column name for the CommitFile table
const CommitFileColumnLanguage = "language"

// CommitFileEscapedColumnLanguage is the escaped Language SQL column name for the CommitFile table
const CommitFileEscapedColumnLanguage = "`language`"

// CommitFileColumnAdditions is the Additions SQL column name for the CommitFile table
const CommitFileColumnAdditions = "additions"

// CommitFileEscapedColumnAdditions is the escaped Additions SQL column name for the CommitFile table
const CommitFileEscapedColumnAdditions = "`additions`"

// CommitFileColumnDeletions is the Deletions SQL column name for the CommitFile table
const CommitFileColumnDeletions = "deletions"

// CommitFileEscapedColumnDeletions is the escaped Deletions SQL column name for the CommitFile table
const CommitFileEscapedColumnDeletions = "`deletions`"

// CommitFileColumnSize is the Size SQL column name for the CommitFile table
const CommitFileColumnSize = "size"

// CommitFileEscapedColumnSize is the escaped Size SQL column name for the CommitFile table
const CommitFileEscapedColumnSize = "`size`"

// CommitFileColumnAbinary is the Abinary SQL column name for the CommitFile table
const CommitFileColumnAbinary = "abinary"

// CommitFileEscapedColumnAbinary is the escaped Abinary SQL column name for the CommitFile table
const CommitFileEscapedColumnAbinary = "`abinary`"

// CommitFileColumnDate is the Date SQL column name for the CommitFile table
const CommitFileColumnDate = "date"

// CommitFileEscapedColumnDate is the escaped Date SQL column name for the CommitFile table
const CommitFileEscapedColumnDate = "`date`"

// CommitFileColumnBranch is the Branch SQL column name for the CommitFile table
const CommitFileColumnBranch = "branch"

// CommitFileEscapedColumnBranch is the escaped Branch SQL column name for the CommitFile table
const CommitFileEscapedColumnBranch = "`branch`"

// CommitFileColumnMergecommit is the Mergecommit SQL column name for the CommitFile table
const CommitFileColumnMergecommit = "mergecommit"

// CommitFileEscapedColumnMergecommit is the escaped Mergecommit SQL column name for the CommitFile table
const CommitFileEscapedColumnMergecommit = "`mergecommit`"

// CommitFileColumnExcluded is the Excluded SQL column name for the CommitFile table
const CommitFileColumnExcluded = "excluded"

// CommitFileEscapedColumnExcluded is the escaped Excluded SQL column name for the CommitFile table
const CommitFileEscapedColumnExcluded = "`excluded`"

// CommitFileColumnLoc is the Loc SQL column name for the CommitFile table
const CommitFileColumnLoc = "loc"

// CommitFileEscapedColumnLoc is the escaped Loc SQL column name for the CommitFile table
const CommitFileEscapedColumnLoc = "`loc`"

// CommitFileColumnSloc is the Sloc SQL column name for the CommitFile table
const CommitFileColumnSloc = "sloc"

// CommitFileEscapedColumnSloc is the escaped Sloc SQL column name for the CommitFile table
const CommitFileEscapedColumnSloc = "`sloc`"

// CommitFileColumnComments is the Comments SQL column name for the CommitFile table
const CommitFileColumnComments = "comments"

// CommitFileEscapedColumnComments is the escaped Comments SQL column name for the CommitFile table
const CommitFileEscapedColumnComments = "`comments`"

// CommitFileColumnBlanks is the Blanks SQL column name for the CommitFile table
const CommitFileColumnBlanks = "blanks"

// CommitFileEscapedColumnBlanks is the escaped Blanks SQL column name for the CommitFile table
const CommitFileEscapedColumnBlanks = "`blanks`"

// CommitFileColumnVariance is the Variance SQL column name for the CommitFile table
const CommitFileColumnVariance = "variance"

// CommitFileEscapedColumnVariance is the escaped Variance SQL column name for the CommitFile table
const CommitFileEscapedColumnVariance = "`variance`"

// CommitFileColumnStatus is the Status SQL column name for the CommitFile table
const CommitFileColumnStatus = "status"

// CommitFileEscapedColumnStatus is the escaped Status SQL column name for the CommitFile table
const CommitFileEscapedColumnStatus = "`status`"

// CommitFileColumnRenamed is the Renamed SQL column name for the CommitFile table
const CommitFileColumnRenamed = "renamed"

// CommitFileEscapedColumnRenamed is the escaped Renamed SQL column name for the CommitFile table
const CommitFileEscapedColumnRenamed = "`renamed`"

// CommitFileColumnRenamedFrom is the RenamedFrom SQL column name for the CommitFile table
const CommitFileColumnRenamedFrom = "renamed_from"

// CommitFileEscapedColumnRenamedFrom is the escaped RenamedFrom SQL column name for the CommitFile table
const CommitFileEscapedColumnRenamedFrom = "`renamed_from`"

// CommitFileColumnRenamedTo is the RenamedTo SQL column name for the CommitFile table
const CommitFileColumnRenamedTo = "renamed_to"

// CommitFileEscapedColumnRenamedTo is the escaped RenamedTo SQL column name for the CommitFile table
const CommitFileEscapedColumnRenamedTo = "`renamed_to`"

// CommitFileColumnCustomerID is the CustomerID SQL column name for the CommitFile table
const CommitFileColumnCustomerID = "customer_id"

// CommitFileEscapedColumnCustomerID is the escaped CustomerID SQL column name for the CommitFile table
const CommitFileEscapedColumnCustomerID = "`customer_id`"

// CommitFileColumnRefType is the RefType SQL column name for the CommitFile table
const CommitFileColumnRefType = "ref_type"

// CommitFileEscapedColumnRefType is the escaped RefType SQL column name for the CommitFile table
const CommitFileEscapedColumnRefType = "`ref_type`"

// CommitFileColumnRefID is the RefID SQL column name for the CommitFile table
const CommitFileColumnRefID = "ref_id"

// CommitFileEscapedColumnRefID is the escaped RefID SQL column name for the CommitFile table
const CommitFileEscapedColumnRefID = "`ref_id`"

// CommitFileColumnMetadata is the Metadata SQL column name for the CommitFile table
const CommitFileColumnMetadata = "metadata"

// CommitFileEscapedColumnMetadata is the escaped Metadata SQL column name for the CommitFile table
const CommitFileEscapedColumnMetadata = "`metadata`"

// GetID will return the CommitFile ID value
func (t *CommitFile) GetID() string {
	return t.ID
}

// SetID will set the CommitFile ID value
func (t *CommitFile) SetID(v string) {
	t.ID = v
}

// FindCommitFileByID will find a CommitFile by ID
func FindCommitFileByID(ctx context.Context, db DB, value string) (*CommitFile, error) {
	q := "SELECT * FROM `commit_file` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CommitID sql.NullString
	var _RepoID sql.NullString
	var _AuthorUserID sql.NullString
	var _CommitterUserID sql.NullString
	var _Filename sql.NullString
	var _Language sql.NullString
	var _Additions sql.NullInt64
	var _Deletions sql.NullInt64
	var _Size sql.NullInt64
	var _Abinary sql.NullBool
	var _Date sql.NullInt64
	var _Branch sql.NullString
	var _Mergecommit sql.NullBool
	var _Excluded sql.NullBool
	var _Loc sql.NullInt64
	var _Sloc sql.NullInt64
	var _Comments sql.NullInt64
	var _Blanks sql.NullInt64
	var _Variance sql.NullInt64
	var _Status sql.NullString
	var _Renamed sql.NullBool
	var _RenamedFrom sql.NullString
	var _RenamedTo sql.NullString
	var _CustomerID sql.NullString
	var _RefType sql.NullString
	var _RefID sql.NullString
	var _Metadata sql.NullString
	err := db.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_CommitID,
		&_RepoID,
		&_AuthorUserID,
		&_CommitterUserID,
		&_Filename,
		&_Language,
		&_Additions,
		&_Deletions,
		&_Size,
		&_Abinary,
		&_Date,
		&_Branch,
		&_Mergecommit,
		&_Excluded,
		&_Loc,
		&_Sloc,
		&_Comments,
		&_Blanks,
		&_Variance,
		&_Status,
		&_Renamed,
		&_RenamedFrom,
		&_RenamedTo,
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
	t := &CommitFile{}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _Checksum.Valid {
		t.SetChecksum(_Checksum.String)
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
	if _CommitterUserID.Valid {
		t.SetCommitterUserID(_CommitterUserID.String)
	}
	if _Filename.Valid {
		t.SetFilename(_Filename.String)
	}
	if _Language.Valid {
		t.SetLanguage(_Language.String)
	}
	if _Additions.Valid {
		t.SetAdditions(int32(_Additions.Int64))
	}
	if _Deletions.Valid {
		t.SetDeletions(int32(_Deletions.Int64))
	}
	if _Size.Valid {
		t.SetSize(int32(_Size.Int64))
	}
	if _Abinary.Valid {
		t.SetAbinary(_Abinary.Bool)
	}
	if _Date.Valid {
		t.SetDate(_Date.Int64)
	}
	if _Branch.Valid {
		t.SetBranch(_Branch.String)
	}
	if _Mergecommit.Valid {
		t.SetMergecommit(_Mergecommit.Bool)
	}
	if _Excluded.Valid {
		t.SetExcluded(_Excluded.Bool)
	}
	if _Loc.Valid {
		t.SetLoc(int32(_Loc.Int64))
	}
	if _Sloc.Valid {
		t.SetSloc(int32(_Sloc.Int64))
	}
	if _Comments.Valid {
		t.SetComments(int32(_Comments.Int64))
	}
	if _Blanks.Valid {
		t.SetBlanks(int32(_Blanks.Int64))
	}
	if _Variance.Valid {
		t.SetVariance(int32(_Variance.Int64))
	}
	if _Status.Valid {
		t.SetStatusString(_Status.String)
	}
	if _Renamed.Valid {
		t.SetRenamed(_Renamed.Bool)
	}
	if _RenamedFrom.Valid {
		t.SetRenamedFrom(_RenamedFrom.String)
	}
	if _RenamedTo.Valid {
		t.SetRenamedTo(_RenamedTo.String)
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

// FindCommitFileByIDTx will find a CommitFile by ID using the provided transaction
func FindCommitFileByIDTx(ctx context.Context, tx Tx, value string) (*CommitFile, error) {
	q := "SELECT * FROM `commit_file` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CommitID sql.NullString
	var _RepoID sql.NullString
	var _AuthorUserID sql.NullString
	var _CommitterUserID sql.NullString
	var _Filename sql.NullString
	var _Language sql.NullString
	var _Additions sql.NullInt64
	var _Deletions sql.NullInt64
	var _Size sql.NullInt64
	var _Abinary sql.NullBool
	var _Date sql.NullInt64
	var _Branch sql.NullString
	var _Mergecommit sql.NullBool
	var _Excluded sql.NullBool
	var _Loc sql.NullInt64
	var _Sloc sql.NullInt64
	var _Comments sql.NullInt64
	var _Blanks sql.NullInt64
	var _Variance sql.NullInt64
	var _Status sql.NullString
	var _Renamed sql.NullBool
	var _RenamedFrom sql.NullString
	var _RenamedTo sql.NullString
	var _CustomerID sql.NullString
	var _RefType sql.NullString
	var _RefID sql.NullString
	var _Metadata sql.NullString
	err := tx.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_CommitID,
		&_RepoID,
		&_AuthorUserID,
		&_CommitterUserID,
		&_Filename,
		&_Language,
		&_Additions,
		&_Deletions,
		&_Size,
		&_Abinary,
		&_Date,
		&_Branch,
		&_Mergecommit,
		&_Excluded,
		&_Loc,
		&_Sloc,
		&_Comments,
		&_Blanks,
		&_Variance,
		&_Status,
		&_Renamed,
		&_RenamedFrom,
		&_RenamedTo,
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
	t := &CommitFile{}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _Checksum.Valid {
		t.SetChecksum(_Checksum.String)
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
	if _CommitterUserID.Valid {
		t.SetCommitterUserID(_CommitterUserID.String)
	}
	if _Filename.Valid {
		t.SetFilename(_Filename.String)
	}
	if _Language.Valid {
		t.SetLanguage(_Language.String)
	}
	if _Additions.Valid {
		t.SetAdditions(int32(_Additions.Int64))
	}
	if _Deletions.Valid {
		t.SetDeletions(int32(_Deletions.Int64))
	}
	if _Size.Valid {
		t.SetSize(int32(_Size.Int64))
	}
	if _Abinary.Valid {
		t.SetAbinary(_Abinary.Bool)
	}
	if _Date.Valid {
		t.SetDate(_Date.Int64)
	}
	if _Branch.Valid {
		t.SetBranch(_Branch.String)
	}
	if _Mergecommit.Valid {
		t.SetMergecommit(_Mergecommit.Bool)
	}
	if _Excluded.Valid {
		t.SetExcluded(_Excluded.Bool)
	}
	if _Loc.Valid {
		t.SetLoc(int32(_Loc.Int64))
	}
	if _Sloc.Valid {
		t.SetSloc(int32(_Sloc.Int64))
	}
	if _Comments.Valid {
		t.SetComments(int32(_Comments.Int64))
	}
	if _Blanks.Valid {
		t.SetBlanks(int32(_Blanks.Int64))
	}
	if _Variance.Valid {
		t.SetVariance(int32(_Variance.Int64))
	}
	if _Status.Valid {
		t.SetStatusString(_Status.String)
	}
	if _Renamed.Valid {
		t.SetRenamed(_Renamed.Bool)
	}
	if _RenamedFrom.Valid {
		t.SetRenamedFrom(_RenamedFrom.String)
	}
	if _RenamedTo.Valid {
		t.SetRenamedTo(_RenamedTo.String)
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

// GetChecksum will return the CommitFile Checksum value
func (t *CommitFile) GetChecksum() string {
	if t.Checksum == nil {
		return ""
	}
	return *t.Checksum
}

// SetChecksum will set the CommitFile Checksum value
func (t *CommitFile) SetChecksum(v string) {
	t.Checksum = &v
}

// GetCommitID will return the CommitFile CommitID value
func (t *CommitFile) GetCommitID() string {
	return t.CommitID
}

// SetCommitID will set the CommitFile CommitID value
func (t *CommitFile) SetCommitID(v string) {
	t.CommitID = v
}

// FindCommitFilesByCommitID will find all CommitFiles by the CommitID value
func FindCommitFilesByCommitID(ctx context.Context, db DB, value string) ([]*CommitFile, error) {
	q := "SELECT * FROM `commit_file` WHERE `commit_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*CommitFile, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CommitID sql.NullString
		var _RepoID sql.NullString
		var _AuthorUserID sql.NullString
		var _CommitterUserID sql.NullString
		var _Filename sql.NullString
		var _Language sql.NullString
		var _Additions sql.NullInt64
		var _Deletions sql.NullInt64
		var _Size sql.NullInt64
		var _Abinary sql.NullBool
		var _Date sql.NullInt64
		var _Branch sql.NullString
		var _Mergecommit sql.NullBool
		var _Excluded sql.NullBool
		var _Loc sql.NullInt64
		var _Sloc sql.NullInt64
		var _Comments sql.NullInt64
		var _Blanks sql.NullInt64
		var _Variance sql.NullInt64
		var _Status sql.NullString
		var _Renamed sql.NullBool
		var _RenamedFrom sql.NullString
		var _RenamedTo sql.NullString
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefID sql.NullString
		var _Metadata sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CommitID,
			&_RepoID,
			&_AuthorUserID,
			&_CommitterUserID,
			&_Filename,
			&_Language,
			&_Additions,
			&_Deletions,
			&_Size,
			&_Abinary,
			&_Date,
			&_Branch,
			&_Mergecommit,
			&_Excluded,
			&_Loc,
			&_Sloc,
			&_Comments,
			&_Blanks,
			&_Variance,
			&_Status,
			&_Renamed,
			&_RenamedFrom,
			&_RenamedTo,
			&_CustomerID,
			&_RefType,
			&_RefID,
			&_Metadata,
		)
		if err != nil {
			return nil, err
		}
		t := &CommitFile{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
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
		if _CommitterUserID.Valid {
			t.SetCommitterUserID(_CommitterUserID.String)
		}
		if _Filename.Valid {
			t.SetFilename(_Filename.String)
		}
		if _Language.Valid {
			t.SetLanguage(_Language.String)
		}
		if _Additions.Valid {
			t.SetAdditions(int32(_Additions.Int64))
		}
		if _Deletions.Valid {
			t.SetDeletions(int32(_Deletions.Int64))
		}
		if _Size.Valid {
			t.SetSize(int32(_Size.Int64))
		}
		if _Abinary.Valid {
			t.SetAbinary(_Abinary.Bool)
		}
		if _Date.Valid {
			t.SetDate(_Date.Int64)
		}
		if _Branch.Valid {
			t.SetBranch(_Branch.String)
		}
		if _Mergecommit.Valid {
			t.SetMergecommit(_Mergecommit.Bool)
		}
		if _Excluded.Valid {
			t.SetExcluded(_Excluded.Bool)
		}
		if _Loc.Valid {
			t.SetLoc(int32(_Loc.Int64))
		}
		if _Sloc.Valid {
			t.SetSloc(int32(_Sloc.Int64))
		}
		if _Comments.Valid {
			t.SetComments(int32(_Comments.Int64))
		}
		if _Blanks.Valid {
			t.SetBlanks(int32(_Blanks.Int64))
		}
		if _Variance.Valid {
			t.SetVariance(int32(_Variance.Int64))
		}
		if _Status.Valid {
			t.SetStatusString(_Status.String)
		}
		if _Renamed.Valid {
			t.SetRenamed(_Renamed.Bool)
		}
		if _RenamedFrom.Valid {
			t.SetRenamedFrom(_RenamedFrom.String)
		}
		if _RenamedTo.Valid {
			t.SetRenamedTo(_RenamedTo.String)
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

// FindCommitFilesByCommitIDTx will find all CommitFiles by the CommitID value using the provided transaction
func FindCommitFilesByCommitIDTx(ctx context.Context, tx Tx, value string) ([]*CommitFile, error) {
	q := "SELECT * FROM `commit_file` WHERE `commit_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*CommitFile, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CommitID sql.NullString
		var _RepoID sql.NullString
		var _AuthorUserID sql.NullString
		var _CommitterUserID sql.NullString
		var _Filename sql.NullString
		var _Language sql.NullString
		var _Additions sql.NullInt64
		var _Deletions sql.NullInt64
		var _Size sql.NullInt64
		var _Abinary sql.NullBool
		var _Date sql.NullInt64
		var _Branch sql.NullString
		var _Mergecommit sql.NullBool
		var _Excluded sql.NullBool
		var _Loc sql.NullInt64
		var _Sloc sql.NullInt64
		var _Comments sql.NullInt64
		var _Blanks sql.NullInt64
		var _Variance sql.NullInt64
		var _Status sql.NullString
		var _Renamed sql.NullBool
		var _RenamedFrom sql.NullString
		var _RenamedTo sql.NullString
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefID sql.NullString
		var _Metadata sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CommitID,
			&_RepoID,
			&_AuthorUserID,
			&_CommitterUserID,
			&_Filename,
			&_Language,
			&_Additions,
			&_Deletions,
			&_Size,
			&_Abinary,
			&_Date,
			&_Branch,
			&_Mergecommit,
			&_Excluded,
			&_Loc,
			&_Sloc,
			&_Comments,
			&_Blanks,
			&_Variance,
			&_Status,
			&_Renamed,
			&_RenamedFrom,
			&_RenamedTo,
			&_CustomerID,
			&_RefType,
			&_RefID,
			&_Metadata,
		)
		if err != nil {
			return nil, err
		}
		t := &CommitFile{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
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
		if _CommitterUserID.Valid {
			t.SetCommitterUserID(_CommitterUserID.String)
		}
		if _Filename.Valid {
			t.SetFilename(_Filename.String)
		}
		if _Language.Valid {
			t.SetLanguage(_Language.String)
		}
		if _Additions.Valid {
			t.SetAdditions(int32(_Additions.Int64))
		}
		if _Deletions.Valid {
			t.SetDeletions(int32(_Deletions.Int64))
		}
		if _Size.Valid {
			t.SetSize(int32(_Size.Int64))
		}
		if _Abinary.Valid {
			t.SetAbinary(_Abinary.Bool)
		}
		if _Date.Valid {
			t.SetDate(_Date.Int64)
		}
		if _Branch.Valid {
			t.SetBranch(_Branch.String)
		}
		if _Mergecommit.Valid {
			t.SetMergecommit(_Mergecommit.Bool)
		}
		if _Excluded.Valid {
			t.SetExcluded(_Excluded.Bool)
		}
		if _Loc.Valid {
			t.SetLoc(int32(_Loc.Int64))
		}
		if _Sloc.Valid {
			t.SetSloc(int32(_Sloc.Int64))
		}
		if _Comments.Valid {
			t.SetComments(int32(_Comments.Int64))
		}
		if _Blanks.Valid {
			t.SetBlanks(int32(_Blanks.Int64))
		}
		if _Variance.Valid {
			t.SetVariance(int32(_Variance.Int64))
		}
		if _Status.Valid {
			t.SetStatusString(_Status.String)
		}
		if _Renamed.Valid {
			t.SetRenamed(_Renamed.Bool)
		}
		if _RenamedFrom.Valid {
			t.SetRenamedFrom(_RenamedFrom.String)
		}
		if _RenamedTo.Valid {
			t.SetRenamedTo(_RenamedTo.String)
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

// GetRepoID will return the CommitFile RepoID value
func (t *CommitFile) GetRepoID() string {
	return t.RepoID
}

// SetRepoID will set the CommitFile RepoID value
func (t *CommitFile) SetRepoID(v string) {
	t.RepoID = v
}

// GetAuthorUserID will return the CommitFile AuthorUserID value
func (t *CommitFile) GetAuthorUserID() string {
	return t.AuthorUserID
}

// SetAuthorUserID will set the CommitFile AuthorUserID value
func (t *CommitFile) SetAuthorUserID(v string) {
	t.AuthorUserID = v
}

// GetCommitterUserID will return the CommitFile CommitterUserID value
func (t *CommitFile) GetCommitterUserID() string {
	return t.CommitterUserID
}

// SetCommitterUserID will set the CommitFile CommitterUserID value
func (t *CommitFile) SetCommitterUserID(v string) {
	t.CommitterUserID = v
}

// GetFilename will return the CommitFile Filename value
func (t *CommitFile) GetFilename() string {
	return t.Filename
}

// SetFilename will set the CommitFile Filename value
func (t *CommitFile) SetFilename(v string) {
	t.Filename = v
}

// GetLanguage will return the CommitFile Language value
func (t *CommitFile) GetLanguage() string {
	return t.Language
}

// SetLanguage will set the CommitFile Language value
func (t *CommitFile) SetLanguage(v string) {
	t.Language = v
}

// GetAdditions will return the CommitFile Additions value
func (t *CommitFile) GetAdditions() int32 {
	return t.Additions
}

// SetAdditions will set the CommitFile Additions value
func (t *CommitFile) SetAdditions(v int32) {
	t.Additions = v
}

// GetDeletions will return the CommitFile Deletions value
func (t *CommitFile) GetDeletions() int32 {
	return t.Deletions
}

// SetDeletions will set the CommitFile Deletions value
func (t *CommitFile) SetDeletions(v int32) {
	t.Deletions = v
}

// GetSize will return the CommitFile Size value
func (t *CommitFile) GetSize() int32 {
	return t.Size
}

// SetSize will set the CommitFile Size value
func (t *CommitFile) SetSize(v int32) {
	t.Size = v
}

// GetAbinary will return the CommitFile Abinary value
func (t *CommitFile) GetAbinary() bool {
	return t.Abinary
}

// SetAbinary will set the CommitFile Abinary value
func (t *CommitFile) SetAbinary(v bool) {
	t.Abinary = v
}

// GetDate will return the CommitFile Date value
func (t *CommitFile) GetDate() int64 {
	return t.Date
}

// SetDate will set the CommitFile Date value
func (t *CommitFile) SetDate(v int64) {
	t.Date = v
}

// GetBranch will return the CommitFile Branch value
func (t *CommitFile) GetBranch() string {
	return t.Branch
}

// SetBranch will set the CommitFile Branch value
func (t *CommitFile) SetBranch(v string) {
	t.Branch = v
}

// GetMergecommit will return the CommitFile Mergecommit value
func (t *CommitFile) GetMergecommit() bool {
	return t.Mergecommit
}

// SetMergecommit will set the CommitFile Mergecommit value
func (t *CommitFile) SetMergecommit(v bool) {
	t.Mergecommit = v
}

// GetExcluded will return the CommitFile Excluded value
func (t *CommitFile) GetExcluded() bool {
	return t.Excluded
}

// SetExcluded will set the CommitFile Excluded value
func (t *CommitFile) SetExcluded(v bool) {
	t.Excluded = v
}

// GetLoc will return the CommitFile Loc value
func (t *CommitFile) GetLoc() int32 {
	return t.Loc
}

// SetLoc will set the CommitFile Loc value
func (t *CommitFile) SetLoc(v int32) {
	t.Loc = v
}

// GetSloc will return the CommitFile Sloc value
func (t *CommitFile) GetSloc() int32 {
	return t.Sloc
}

// SetSloc will set the CommitFile Sloc value
func (t *CommitFile) SetSloc(v int32) {
	t.Sloc = v
}

// GetComments will return the CommitFile Comments value
func (t *CommitFile) GetComments() int32 {
	return t.Comments
}

// SetComments will set the CommitFile Comments value
func (t *CommitFile) SetComments(v int32) {
	t.Comments = v
}

// GetBlanks will return the CommitFile Blanks value
func (t *CommitFile) GetBlanks() int32 {
	return t.Blanks
}

// SetBlanks will set the CommitFile Blanks value
func (t *CommitFile) SetBlanks(v int32) {
	t.Blanks = v
}

// GetVariance will return the CommitFile Variance value
func (t *CommitFile) GetVariance() int32 {
	return t.Variance
}

// SetVariance will set the CommitFile Variance value
func (t *CommitFile) SetVariance(v int32) {
	t.Variance = v
}

// GetStatus will return the CommitFile Status value
func (t *CommitFile) GetStatus() CommitFile_CommitFileStatus {
	return t.Status
}

// SetStatus will set the CommitFile Status value
func (t *CommitFile) SetStatus(v CommitFile_CommitFileStatus) {
	t.Status = v
}

// GetStatusString will return the CommitFile Status value as a string
func (t *CommitFile) GetStatusString() string {
	return t.Status.String()
}

// SetStatusString will set the CommitFile Status value from a string
func (t *CommitFile) SetStatusString(v string) {
	var _Status = toCommitFileStatus(v)
	if _Status != nil {
		t.Status = *_Status
	}
}

// GetRenamed will return the CommitFile Renamed value
func (t *CommitFile) GetRenamed() bool {
	return t.Renamed
}

// SetRenamed will set the CommitFile Renamed value
func (t *CommitFile) SetRenamed(v bool) {
	t.Renamed = v
}

// GetRenamedFrom will return the CommitFile RenamedFrom value
func (t *CommitFile) GetRenamedFrom() string {
	if t.RenamedFrom == nil {
		return ""
	}
	return *t.RenamedFrom
}

// SetRenamedFrom will set the CommitFile RenamedFrom value
func (t *CommitFile) SetRenamedFrom(v string) {
	t.RenamedFrom = &v
}

// GetRenamedTo will return the CommitFile RenamedTo value
func (t *CommitFile) GetRenamedTo() string {
	if t.RenamedTo == nil {
		return ""
	}
	return *t.RenamedTo
}

// SetRenamedTo will set the CommitFile RenamedTo value
func (t *CommitFile) SetRenamedTo(v string) {
	t.RenamedTo = &v
}

// GetCustomerID will return the CommitFile CustomerID value
func (t *CommitFile) GetCustomerID() string {
	return t.CustomerID
}

// SetCustomerID will set the CommitFile CustomerID value
func (t *CommitFile) SetCustomerID(v string) {
	t.CustomerID = v
}

// FindCommitFilesByCustomerID will find all CommitFiles by the CustomerID value
func FindCommitFilesByCustomerID(ctx context.Context, db DB, value string) ([]*CommitFile, error) {
	q := "SELECT * FROM `commit_file` WHERE `customer_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*CommitFile, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CommitID sql.NullString
		var _RepoID sql.NullString
		var _AuthorUserID sql.NullString
		var _CommitterUserID sql.NullString
		var _Filename sql.NullString
		var _Language sql.NullString
		var _Additions sql.NullInt64
		var _Deletions sql.NullInt64
		var _Size sql.NullInt64
		var _Abinary sql.NullBool
		var _Date sql.NullInt64
		var _Branch sql.NullString
		var _Mergecommit sql.NullBool
		var _Excluded sql.NullBool
		var _Loc sql.NullInt64
		var _Sloc sql.NullInt64
		var _Comments sql.NullInt64
		var _Blanks sql.NullInt64
		var _Variance sql.NullInt64
		var _Status sql.NullString
		var _Renamed sql.NullBool
		var _RenamedFrom sql.NullString
		var _RenamedTo sql.NullString
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefID sql.NullString
		var _Metadata sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CommitID,
			&_RepoID,
			&_AuthorUserID,
			&_CommitterUserID,
			&_Filename,
			&_Language,
			&_Additions,
			&_Deletions,
			&_Size,
			&_Abinary,
			&_Date,
			&_Branch,
			&_Mergecommit,
			&_Excluded,
			&_Loc,
			&_Sloc,
			&_Comments,
			&_Blanks,
			&_Variance,
			&_Status,
			&_Renamed,
			&_RenamedFrom,
			&_RenamedTo,
			&_CustomerID,
			&_RefType,
			&_RefID,
			&_Metadata,
		)
		if err != nil {
			return nil, err
		}
		t := &CommitFile{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
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
		if _CommitterUserID.Valid {
			t.SetCommitterUserID(_CommitterUserID.String)
		}
		if _Filename.Valid {
			t.SetFilename(_Filename.String)
		}
		if _Language.Valid {
			t.SetLanguage(_Language.String)
		}
		if _Additions.Valid {
			t.SetAdditions(int32(_Additions.Int64))
		}
		if _Deletions.Valid {
			t.SetDeletions(int32(_Deletions.Int64))
		}
		if _Size.Valid {
			t.SetSize(int32(_Size.Int64))
		}
		if _Abinary.Valid {
			t.SetAbinary(_Abinary.Bool)
		}
		if _Date.Valid {
			t.SetDate(_Date.Int64)
		}
		if _Branch.Valid {
			t.SetBranch(_Branch.String)
		}
		if _Mergecommit.Valid {
			t.SetMergecommit(_Mergecommit.Bool)
		}
		if _Excluded.Valid {
			t.SetExcluded(_Excluded.Bool)
		}
		if _Loc.Valid {
			t.SetLoc(int32(_Loc.Int64))
		}
		if _Sloc.Valid {
			t.SetSloc(int32(_Sloc.Int64))
		}
		if _Comments.Valid {
			t.SetComments(int32(_Comments.Int64))
		}
		if _Blanks.Valid {
			t.SetBlanks(int32(_Blanks.Int64))
		}
		if _Variance.Valid {
			t.SetVariance(int32(_Variance.Int64))
		}
		if _Status.Valid {
			t.SetStatusString(_Status.String)
		}
		if _Renamed.Valid {
			t.SetRenamed(_Renamed.Bool)
		}
		if _RenamedFrom.Valid {
			t.SetRenamedFrom(_RenamedFrom.String)
		}
		if _RenamedTo.Valid {
			t.SetRenamedTo(_RenamedTo.String)
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

// FindCommitFilesByCustomerIDTx will find all CommitFiles by the CustomerID value using the provided transaction
func FindCommitFilesByCustomerIDTx(ctx context.Context, tx Tx, value string) ([]*CommitFile, error) {
	q := "SELECT * FROM `commit_file` WHERE `customer_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*CommitFile, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CommitID sql.NullString
		var _RepoID sql.NullString
		var _AuthorUserID sql.NullString
		var _CommitterUserID sql.NullString
		var _Filename sql.NullString
		var _Language sql.NullString
		var _Additions sql.NullInt64
		var _Deletions sql.NullInt64
		var _Size sql.NullInt64
		var _Abinary sql.NullBool
		var _Date sql.NullInt64
		var _Branch sql.NullString
		var _Mergecommit sql.NullBool
		var _Excluded sql.NullBool
		var _Loc sql.NullInt64
		var _Sloc sql.NullInt64
		var _Comments sql.NullInt64
		var _Blanks sql.NullInt64
		var _Variance sql.NullInt64
		var _Status sql.NullString
		var _Renamed sql.NullBool
		var _RenamedFrom sql.NullString
		var _RenamedTo sql.NullString
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefID sql.NullString
		var _Metadata sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CommitID,
			&_RepoID,
			&_AuthorUserID,
			&_CommitterUserID,
			&_Filename,
			&_Language,
			&_Additions,
			&_Deletions,
			&_Size,
			&_Abinary,
			&_Date,
			&_Branch,
			&_Mergecommit,
			&_Excluded,
			&_Loc,
			&_Sloc,
			&_Comments,
			&_Blanks,
			&_Variance,
			&_Status,
			&_Renamed,
			&_RenamedFrom,
			&_RenamedTo,
			&_CustomerID,
			&_RefType,
			&_RefID,
			&_Metadata,
		)
		if err != nil {
			return nil, err
		}
		t := &CommitFile{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
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
		if _CommitterUserID.Valid {
			t.SetCommitterUserID(_CommitterUserID.String)
		}
		if _Filename.Valid {
			t.SetFilename(_Filename.String)
		}
		if _Language.Valid {
			t.SetLanguage(_Language.String)
		}
		if _Additions.Valid {
			t.SetAdditions(int32(_Additions.Int64))
		}
		if _Deletions.Valid {
			t.SetDeletions(int32(_Deletions.Int64))
		}
		if _Size.Valid {
			t.SetSize(int32(_Size.Int64))
		}
		if _Abinary.Valid {
			t.SetAbinary(_Abinary.Bool)
		}
		if _Date.Valid {
			t.SetDate(_Date.Int64)
		}
		if _Branch.Valid {
			t.SetBranch(_Branch.String)
		}
		if _Mergecommit.Valid {
			t.SetMergecommit(_Mergecommit.Bool)
		}
		if _Excluded.Valid {
			t.SetExcluded(_Excluded.Bool)
		}
		if _Loc.Valid {
			t.SetLoc(int32(_Loc.Int64))
		}
		if _Sloc.Valid {
			t.SetSloc(int32(_Sloc.Int64))
		}
		if _Comments.Valid {
			t.SetComments(int32(_Comments.Int64))
		}
		if _Blanks.Valid {
			t.SetBlanks(int32(_Blanks.Int64))
		}
		if _Variance.Valid {
			t.SetVariance(int32(_Variance.Int64))
		}
		if _Status.Valid {
			t.SetStatusString(_Status.String)
		}
		if _Renamed.Valid {
			t.SetRenamed(_Renamed.Bool)
		}
		if _RenamedFrom.Valid {
			t.SetRenamedFrom(_RenamedFrom.String)
		}
		if _RenamedTo.Valid {
			t.SetRenamedTo(_RenamedTo.String)
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

// GetRefType will return the CommitFile RefType value
func (t *CommitFile) GetRefType() string {
	return t.RefType
}

// SetRefType will set the CommitFile RefType value
func (t *CommitFile) SetRefType(v string) {
	t.RefType = v
}

// FindCommitFilesByRefType will find all CommitFiles by the RefType value
func FindCommitFilesByRefType(ctx context.Context, db DB, value string) ([]*CommitFile, error) {
	q := "SELECT * FROM `commit_file` WHERE `ref_type` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*CommitFile, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CommitID sql.NullString
		var _RepoID sql.NullString
		var _AuthorUserID sql.NullString
		var _CommitterUserID sql.NullString
		var _Filename sql.NullString
		var _Language sql.NullString
		var _Additions sql.NullInt64
		var _Deletions sql.NullInt64
		var _Size sql.NullInt64
		var _Abinary sql.NullBool
		var _Date sql.NullInt64
		var _Branch sql.NullString
		var _Mergecommit sql.NullBool
		var _Excluded sql.NullBool
		var _Loc sql.NullInt64
		var _Sloc sql.NullInt64
		var _Comments sql.NullInt64
		var _Blanks sql.NullInt64
		var _Variance sql.NullInt64
		var _Status sql.NullString
		var _Renamed sql.NullBool
		var _RenamedFrom sql.NullString
		var _RenamedTo sql.NullString
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefID sql.NullString
		var _Metadata sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CommitID,
			&_RepoID,
			&_AuthorUserID,
			&_CommitterUserID,
			&_Filename,
			&_Language,
			&_Additions,
			&_Deletions,
			&_Size,
			&_Abinary,
			&_Date,
			&_Branch,
			&_Mergecommit,
			&_Excluded,
			&_Loc,
			&_Sloc,
			&_Comments,
			&_Blanks,
			&_Variance,
			&_Status,
			&_Renamed,
			&_RenamedFrom,
			&_RenamedTo,
			&_CustomerID,
			&_RefType,
			&_RefID,
			&_Metadata,
		)
		if err != nil {
			return nil, err
		}
		t := &CommitFile{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
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
		if _CommitterUserID.Valid {
			t.SetCommitterUserID(_CommitterUserID.String)
		}
		if _Filename.Valid {
			t.SetFilename(_Filename.String)
		}
		if _Language.Valid {
			t.SetLanguage(_Language.String)
		}
		if _Additions.Valid {
			t.SetAdditions(int32(_Additions.Int64))
		}
		if _Deletions.Valid {
			t.SetDeletions(int32(_Deletions.Int64))
		}
		if _Size.Valid {
			t.SetSize(int32(_Size.Int64))
		}
		if _Abinary.Valid {
			t.SetAbinary(_Abinary.Bool)
		}
		if _Date.Valid {
			t.SetDate(_Date.Int64)
		}
		if _Branch.Valid {
			t.SetBranch(_Branch.String)
		}
		if _Mergecommit.Valid {
			t.SetMergecommit(_Mergecommit.Bool)
		}
		if _Excluded.Valid {
			t.SetExcluded(_Excluded.Bool)
		}
		if _Loc.Valid {
			t.SetLoc(int32(_Loc.Int64))
		}
		if _Sloc.Valid {
			t.SetSloc(int32(_Sloc.Int64))
		}
		if _Comments.Valid {
			t.SetComments(int32(_Comments.Int64))
		}
		if _Blanks.Valid {
			t.SetBlanks(int32(_Blanks.Int64))
		}
		if _Variance.Valid {
			t.SetVariance(int32(_Variance.Int64))
		}
		if _Status.Valid {
			t.SetStatusString(_Status.String)
		}
		if _Renamed.Valid {
			t.SetRenamed(_Renamed.Bool)
		}
		if _RenamedFrom.Valid {
			t.SetRenamedFrom(_RenamedFrom.String)
		}
		if _RenamedTo.Valid {
			t.SetRenamedTo(_RenamedTo.String)
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

// FindCommitFilesByRefTypeTx will find all CommitFiles by the RefType value using the provided transaction
func FindCommitFilesByRefTypeTx(ctx context.Context, tx Tx, value string) ([]*CommitFile, error) {
	q := "SELECT * FROM `commit_file` WHERE `ref_type` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*CommitFile, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CommitID sql.NullString
		var _RepoID sql.NullString
		var _AuthorUserID sql.NullString
		var _CommitterUserID sql.NullString
		var _Filename sql.NullString
		var _Language sql.NullString
		var _Additions sql.NullInt64
		var _Deletions sql.NullInt64
		var _Size sql.NullInt64
		var _Abinary sql.NullBool
		var _Date sql.NullInt64
		var _Branch sql.NullString
		var _Mergecommit sql.NullBool
		var _Excluded sql.NullBool
		var _Loc sql.NullInt64
		var _Sloc sql.NullInt64
		var _Comments sql.NullInt64
		var _Blanks sql.NullInt64
		var _Variance sql.NullInt64
		var _Status sql.NullString
		var _Renamed sql.NullBool
		var _RenamedFrom sql.NullString
		var _RenamedTo sql.NullString
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefID sql.NullString
		var _Metadata sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CommitID,
			&_RepoID,
			&_AuthorUserID,
			&_CommitterUserID,
			&_Filename,
			&_Language,
			&_Additions,
			&_Deletions,
			&_Size,
			&_Abinary,
			&_Date,
			&_Branch,
			&_Mergecommit,
			&_Excluded,
			&_Loc,
			&_Sloc,
			&_Comments,
			&_Blanks,
			&_Variance,
			&_Status,
			&_Renamed,
			&_RenamedFrom,
			&_RenamedTo,
			&_CustomerID,
			&_RefType,
			&_RefID,
			&_Metadata,
		)
		if err != nil {
			return nil, err
		}
		t := &CommitFile{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
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
		if _CommitterUserID.Valid {
			t.SetCommitterUserID(_CommitterUserID.String)
		}
		if _Filename.Valid {
			t.SetFilename(_Filename.String)
		}
		if _Language.Valid {
			t.SetLanguage(_Language.String)
		}
		if _Additions.Valid {
			t.SetAdditions(int32(_Additions.Int64))
		}
		if _Deletions.Valid {
			t.SetDeletions(int32(_Deletions.Int64))
		}
		if _Size.Valid {
			t.SetSize(int32(_Size.Int64))
		}
		if _Abinary.Valid {
			t.SetAbinary(_Abinary.Bool)
		}
		if _Date.Valid {
			t.SetDate(_Date.Int64)
		}
		if _Branch.Valid {
			t.SetBranch(_Branch.String)
		}
		if _Mergecommit.Valid {
			t.SetMergecommit(_Mergecommit.Bool)
		}
		if _Excluded.Valid {
			t.SetExcluded(_Excluded.Bool)
		}
		if _Loc.Valid {
			t.SetLoc(int32(_Loc.Int64))
		}
		if _Sloc.Valid {
			t.SetSloc(int32(_Sloc.Int64))
		}
		if _Comments.Valid {
			t.SetComments(int32(_Comments.Int64))
		}
		if _Blanks.Valid {
			t.SetBlanks(int32(_Blanks.Int64))
		}
		if _Variance.Valid {
			t.SetVariance(int32(_Variance.Int64))
		}
		if _Status.Valid {
			t.SetStatusString(_Status.String)
		}
		if _Renamed.Valid {
			t.SetRenamed(_Renamed.Bool)
		}
		if _RenamedFrom.Valid {
			t.SetRenamedFrom(_RenamedFrom.String)
		}
		if _RenamedTo.Valid {
			t.SetRenamedTo(_RenamedTo.String)
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

// GetRefID will return the CommitFile RefID value
func (t *CommitFile) GetRefID() string {
	return t.RefID
}

// SetRefID will set the CommitFile RefID value
func (t *CommitFile) SetRefID(v string) {
	t.RefID = v
}

// FindCommitFilesByRefID will find all CommitFiles by the RefID value
func FindCommitFilesByRefID(ctx context.Context, db DB, value string) ([]*CommitFile, error) {
	q := "SELECT * FROM `commit_file` WHERE `ref_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*CommitFile, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CommitID sql.NullString
		var _RepoID sql.NullString
		var _AuthorUserID sql.NullString
		var _CommitterUserID sql.NullString
		var _Filename sql.NullString
		var _Language sql.NullString
		var _Additions sql.NullInt64
		var _Deletions sql.NullInt64
		var _Size sql.NullInt64
		var _Abinary sql.NullBool
		var _Date sql.NullInt64
		var _Branch sql.NullString
		var _Mergecommit sql.NullBool
		var _Excluded sql.NullBool
		var _Loc sql.NullInt64
		var _Sloc sql.NullInt64
		var _Comments sql.NullInt64
		var _Blanks sql.NullInt64
		var _Variance sql.NullInt64
		var _Status sql.NullString
		var _Renamed sql.NullBool
		var _RenamedFrom sql.NullString
		var _RenamedTo sql.NullString
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefID sql.NullString
		var _Metadata sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CommitID,
			&_RepoID,
			&_AuthorUserID,
			&_CommitterUserID,
			&_Filename,
			&_Language,
			&_Additions,
			&_Deletions,
			&_Size,
			&_Abinary,
			&_Date,
			&_Branch,
			&_Mergecommit,
			&_Excluded,
			&_Loc,
			&_Sloc,
			&_Comments,
			&_Blanks,
			&_Variance,
			&_Status,
			&_Renamed,
			&_RenamedFrom,
			&_RenamedTo,
			&_CustomerID,
			&_RefType,
			&_RefID,
			&_Metadata,
		)
		if err != nil {
			return nil, err
		}
		t := &CommitFile{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
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
		if _CommitterUserID.Valid {
			t.SetCommitterUserID(_CommitterUserID.String)
		}
		if _Filename.Valid {
			t.SetFilename(_Filename.String)
		}
		if _Language.Valid {
			t.SetLanguage(_Language.String)
		}
		if _Additions.Valid {
			t.SetAdditions(int32(_Additions.Int64))
		}
		if _Deletions.Valid {
			t.SetDeletions(int32(_Deletions.Int64))
		}
		if _Size.Valid {
			t.SetSize(int32(_Size.Int64))
		}
		if _Abinary.Valid {
			t.SetAbinary(_Abinary.Bool)
		}
		if _Date.Valid {
			t.SetDate(_Date.Int64)
		}
		if _Branch.Valid {
			t.SetBranch(_Branch.String)
		}
		if _Mergecommit.Valid {
			t.SetMergecommit(_Mergecommit.Bool)
		}
		if _Excluded.Valid {
			t.SetExcluded(_Excluded.Bool)
		}
		if _Loc.Valid {
			t.SetLoc(int32(_Loc.Int64))
		}
		if _Sloc.Valid {
			t.SetSloc(int32(_Sloc.Int64))
		}
		if _Comments.Valid {
			t.SetComments(int32(_Comments.Int64))
		}
		if _Blanks.Valid {
			t.SetBlanks(int32(_Blanks.Int64))
		}
		if _Variance.Valid {
			t.SetVariance(int32(_Variance.Int64))
		}
		if _Status.Valid {
			t.SetStatusString(_Status.String)
		}
		if _Renamed.Valid {
			t.SetRenamed(_Renamed.Bool)
		}
		if _RenamedFrom.Valid {
			t.SetRenamedFrom(_RenamedFrom.String)
		}
		if _RenamedTo.Valid {
			t.SetRenamedTo(_RenamedTo.String)
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

// FindCommitFilesByRefIDTx will find all CommitFiles by the RefID value using the provided transaction
func FindCommitFilesByRefIDTx(ctx context.Context, tx Tx, value string) ([]*CommitFile, error) {
	q := "SELECT * FROM `commit_file` WHERE `ref_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*CommitFile, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CommitID sql.NullString
		var _RepoID sql.NullString
		var _AuthorUserID sql.NullString
		var _CommitterUserID sql.NullString
		var _Filename sql.NullString
		var _Language sql.NullString
		var _Additions sql.NullInt64
		var _Deletions sql.NullInt64
		var _Size sql.NullInt64
		var _Abinary sql.NullBool
		var _Date sql.NullInt64
		var _Branch sql.NullString
		var _Mergecommit sql.NullBool
		var _Excluded sql.NullBool
		var _Loc sql.NullInt64
		var _Sloc sql.NullInt64
		var _Comments sql.NullInt64
		var _Blanks sql.NullInt64
		var _Variance sql.NullInt64
		var _Status sql.NullString
		var _Renamed sql.NullBool
		var _RenamedFrom sql.NullString
		var _RenamedTo sql.NullString
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefID sql.NullString
		var _Metadata sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CommitID,
			&_RepoID,
			&_AuthorUserID,
			&_CommitterUserID,
			&_Filename,
			&_Language,
			&_Additions,
			&_Deletions,
			&_Size,
			&_Abinary,
			&_Date,
			&_Branch,
			&_Mergecommit,
			&_Excluded,
			&_Loc,
			&_Sloc,
			&_Comments,
			&_Blanks,
			&_Variance,
			&_Status,
			&_Renamed,
			&_RenamedFrom,
			&_RenamedTo,
			&_CustomerID,
			&_RefType,
			&_RefID,
			&_Metadata,
		)
		if err != nil {
			return nil, err
		}
		t := &CommitFile{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
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
		if _CommitterUserID.Valid {
			t.SetCommitterUserID(_CommitterUserID.String)
		}
		if _Filename.Valid {
			t.SetFilename(_Filename.String)
		}
		if _Language.Valid {
			t.SetLanguage(_Language.String)
		}
		if _Additions.Valid {
			t.SetAdditions(int32(_Additions.Int64))
		}
		if _Deletions.Valid {
			t.SetDeletions(int32(_Deletions.Int64))
		}
		if _Size.Valid {
			t.SetSize(int32(_Size.Int64))
		}
		if _Abinary.Valid {
			t.SetAbinary(_Abinary.Bool)
		}
		if _Date.Valid {
			t.SetDate(_Date.Int64)
		}
		if _Branch.Valid {
			t.SetBranch(_Branch.String)
		}
		if _Mergecommit.Valid {
			t.SetMergecommit(_Mergecommit.Bool)
		}
		if _Excluded.Valid {
			t.SetExcluded(_Excluded.Bool)
		}
		if _Loc.Valid {
			t.SetLoc(int32(_Loc.Int64))
		}
		if _Sloc.Valid {
			t.SetSloc(int32(_Sloc.Int64))
		}
		if _Comments.Valid {
			t.SetComments(int32(_Comments.Int64))
		}
		if _Blanks.Valid {
			t.SetBlanks(int32(_Blanks.Int64))
		}
		if _Variance.Valid {
			t.SetVariance(int32(_Variance.Int64))
		}
		if _Status.Valid {
			t.SetStatusString(_Status.String)
		}
		if _Renamed.Valid {
			t.SetRenamed(_Renamed.Bool)
		}
		if _RenamedFrom.Valid {
			t.SetRenamedFrom(_RenamedFrom.String)
		}
		if _RenamedTo.Valid {
			t.SetRenamedTo(_RenamedTo.String)
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

// GetMetadata will return the CommitFile Metadata value
func (t *CommitFile) GetMetadata() string {
	if t.Metadata == nil {
		return ""
	}
	return *t.Metadata
}

// SetMetadata will set the CommitFile Metadata value
func (t *CommitFile) SetMetadata(v string) {
	t.Metadata = &v
}

func (t *CommitFile) toTimestamp(value time.Time) *timestamp.Timestamp {
	ts, _ := ptypes.TimestampProto(value)
	return ts
}

// DBCreateCommitFileTable will create the CommitFile table
func DBCreateCommitFileTable(ctx context.Context, db DB) error {
	q := "CREATE TABLE `commit_file` (`id` VARCHAR(64) NOT NULL PRIMARY KEY,`checksum` CHAR(64),`commit_id`VARCHAR(64) NOT NULL,`repo_id` VARCHAR(64) NOT NULL,`author_user_id` VARCHAR(64) NOT NULL,`committer_user_id` VARCHAR(64) NOT NULL,`filename` VARCHAR(1024) NOT NULL,`language` VARCHAR(100) NOT NULL DEFAULT \"unknown\",`additions`INT NOT NULL DEFAULT 0,`deletions`INT NOT NULL DEFAULT 0,`size` INT NOT NULL DEFAULT 0,`abinary` BOOL NOT NULL DEFAULT false,`date` BIGINT UNSIGNED NOT NULL,`branch`VARCHAR(255) NOT NULL DEFAULT \"master\",`mergecommit` BOOL NOT NULL DEFAULT false,`excluded` BOOL NOT NULL DEFAULT false,`loc`INT NOT NULL DEFAULT 0,`sloc` INT NOT NULL DEFAULT 0,`comments` INT NOT NULL DEFAULT 0,`blanks`INT NOT NULL DEFAULT 0,`variance` INT NOT NULL DEFAULT 0,`status`ENUM('added','modified','deleted') NOT NULL,`renamed` BOOL NOT NULL DEFAULT false,`renamed_from`TEXT,`renamed_to` TEXT,`customer_id` VARCHAR(64) NOT NULL,`ref_type` VARCHAR(20) NOT NULL,`ref_id`VARCHAR(64) NOT NULL,`metadata` JSON,INDEX commit_file_commit_id_index (`commit_id`),INDEX commit_file_customer_id_index (`customer_id`),INDEX commit_file_ref_type_index (`ref_type`),INDEX commit_file_ref_id_index (`ref_id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBCreateCommitFileTableTx will create the CommitFile table using the provided transction
func DBCreateCommitFileTableTx(ctx context.Context, tx Tx) error {
	q := "CREATE TABLE `commit_file` (`id` VARCHAR(64) NOT NULL PRIMARY KEY,`checksum` CHAR(64),`commit_id`VARCHAR(64) NOT NULL,`repo_id` VARCHAR(64) NOT NULL,`author_user_id` VARCHAR(64) NOT NULL,`committer_user_id` VARCHAR(64) NOT NULL,`filename` VARCHAR(1024) NOT NULL,`language` VARCHAR(100) NOT NULL DEFAULT \"unknown\",`additions`INT NOT NULL DEFAULT 0,`deletions`INT NOT NULL DEFAULT 0,`size` INT NOT NULL DEFAULT 0,`abinary` BOOL NOT NULL DEFAULT false,`date` BIGINT UNSIGNED NOT NULL,`branch`VARCHAR(255) NOT NULL DEFAULT \"master\",`mergecommit` BOOL NOT NULL DEFAULT false,`excluded` BOOL NOT NULL DEFAULT false,`loc`INT NOT NULL DEFAULT 0,`sloc` INT NOT NULL DEFAULT 0,`comments` INT NOT NULL DEFAULT 0,`blanks`INT NOT NULL DEFAULT 0,`variance` INT NOT NULL DEFAULT 0,`status`ENUM('added','modified','deleted') NOT NULL,`renamed` BOOL NOT NULL DEFAULT false,`renamed_from`TEXT,`renamed_to` TEXT,`customer_id` VARCHAR(64) NOT NULL,`ref_type` VARCHAR(20) NOT NULL,`ref_id`VARCHAR(64) NOT NULL,`metadata` JSON,INDEX commit_file_commit_id_index (`commit_id`),INDEX commit_file_customer_id_index (`customer_id`),INDEX commit_file_ref_type_index (`ref_type`),INDEX commit_file_ref_id_index (`ref_id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// DBDropCommitFileTable will drop the CommitFile table
func DBDropCommitFileTable(ctx context.Context, db DB) error {
	q := "DROP TABLE IF EXISTS `commit_file`"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBDropCommitFileTableTx will drop the CommitFile table using the provided transaction
func DBDropCommitFileTableTx(ctx context.Context, tx Tx) error {
	q := "DROP TABLE IF EXISTS `commit_file`"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// CalculateChecksum will calculate a checksum of the SHA1 of all field values
func (t *CommitFile) CalculateChecksum() string {
	return orm.HashStrings(
		orm.ToString(t.ID),
		orm.ToString(t.CommitID),
		orm.ToString(t.RepoID),
		orm.ToString(t.AuthorUserID),
		orm.ToString(t.CommitterUserID),
		orm.ToString(t.Filename),
		orm.ToString(t.Language),
		orm.ToString(t.Additions),
		orm.ToString(t.Deletions),
		orm.ToString(t.Size),
		orm.ToString(t.Abinary),
		orm.ToString(t.Date),
		orm.ToString(t.Branch),
		orm.ToString(t.Mergecommit),
		orm.ToString(t.Excluded),
		orm.ToString(t.Loc),
		orm.ToString(t.Sloc),
		orm.ToString(t.Comments),
		orm.ToString(t.Blanks),
		orm.ToString(t.Variance),
		enumCommitFileStatusToString(&t.Status),
		orm.ToString(t.Renamed),
		orm.ToString(t.RenamedFrom),
		orm.ToString(t.RenamedTo),
		orm.ToString(t.CustomerID),
		orm.ToString(t.RefType),
		orm.ToString(t.RefID),
		orm.ToString(t.Metadata),
	)
}

// DBCreate will create a new CommitFile record in the database
func (t *CommitFile) DBCreate(ctx context.Context, db DB) (sql.Result, error) {
	q := "INSERT INTO `commit_file` (`commit_file`.`id`,`commit_file`.`checksum`,`commit_file`.`commit_id`,`commit_file`.`repo_id`,`commit_file`.`author_user_id`,`commit_file`.`committer_user_id`,`commit_file`.`filename`,`commit_file`.`language`,`commit_file`.`additions`,`commit_file`.`deletions`,`commit_file`.`size`,`commit_file`.`abinary`,`commit_file`.`date`,`commit_file`.`branch`,`commit_file`.`mergecommit`,`commit_file`.`excluded`,`commit_file`.`loc`,`commit_file`.`sloc`,`commit_file`.`comments`,`commit_file`.`blanks`,`commit_file`.`variance`,`commit_file`.`status`,`commit_file`.`renamed`,`commit_file`.`renamed_from`,`commit_file`.`renamed_to`,`commit_file`.`customer_id`,`commit_file`.`ref_type`,`commit_file`.`ref_id`,`commit_file`.`metadata`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CommitID),
		orm.ToSQLString(t.RepoID),
		orm.ToSQLString(t.AuthorUserID),
		orm.ToSQLString(t.CommitterUserID),
		orm.ToSQLString(t.Filename),
		orm.ToSQLString(t.Language),
		orm.ToSQLInt64(t.Additions),
		orm.ToSQLInt64(t.Deletions),
		orm.ToSQLInt64(t.Size),
		orm.ToSQLBool(t.Abinary),
		orm.ToSQLInt64(t.Date),
		orm.ToSQLString(t.Branch),
		orm.ToSQLBool(t.Mergecommit),
		orm.ToSQLBool(t.Excluded),
		orm.ToSQLInt64(t.Loc),
		orm.ToSQLInt64(t.Sloc),
		orm.ToSQLInt64(t.Comments),
		orm.ToSQLInt64(t.Blanks),
		orm.ToSQLInt64(t.Variance),
		orm.ToSQLString(enumCommitFileStatusToString(&t.Status)),
		orm.ToSQLBool(t.Renamed),
		orm.ToSQLString(t.RenamedFrom),
		orm.ToSQLString(t.RenamedTo),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefType),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.Metadata),
	)
}

// DBCreateTx will create a new CommitFile record in the database using the provided transaction
func (t *CommitFile) DBCreateTx(ctx context.Context, tx Tx) (sql.Result, error) {
	q := "INSERT INTO `commit_file` (`commit_file`.`id`,`commit_file`.`checksum`,`commit_file`.`commit_id`,`commit_file`.`repo_id`,`commit_file`.`author_user_id`,`commit_file`.`committer_user_id`,`commit_file`.`filename`,`commit_file`.`language`,`commit_file`.`additions`,`commit_file`.`deletions`,`commit_file`.`size`,`commit_file`.`abinary`,`commit_file`.`date`,`commit_file`.`branch`,`commit_file`.`mergecommit`,`commit_file`.`excluded`,`commit_file`.`loc`,`commit_file`.`sloc`,`commit_file`.`comments`,`commit_file`.`blanks`,`commit_file`.`variance`,`commit_file`.`status`,`commit_file`.`renamed`,`commit_file`.`renamed_from`,`commit_file`.`renamed_to`,`commit_file`.`customer_id`,`commit_file`.`ref_type`,`commit_file`.`ref_id`,`commit_file`.`metadata`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CommitID),
		orm.ToSQLString(t.RepoID),
		orm.ToSQLString(t.AuthorUserID),
		orm.ToSQLString(t.CommitterUserID),
		orm.ToSQLString(t.Filename),
		orm.ToSQLString(t.Language),
		orm.ToSQLInt64(t.Additions),
		orm.ToSQLInt64(t.Deletions),
		orm.ToSQLInt64(t.Size),
		orm.ToSQLBool(t.Abinary),
		orm.ToSQLInt64(t.Date),
		orm.ToSQLString(t.Branch),
		orm.ToSQLBool(t.Mergecommit),
		orm.ToSQLBool(t.Excluded),
		orm.ToSQLInt64(t.Loc),
		orm.ToSQLInt64(t.Sloc),
		orm.ToSQLInt64(t.Comments),
		orm.ToSQLInt64(t.Blanks),
		orm.ToSQLInt64(t.Variance),
		orm.ToSQLString(enumCommitFileStatusToString(&t.Status)),
		orm.ToSQLBool(t.Renamed),
		orm.ToSQLString(t.RenamedFrom),
		orm.ToSQLString(t.RenamedTo),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefType),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.Metadata),
	)
}

// DBCreateIgnoreDuplicate will upsert the CommitFile record in the database
func (t *CommitFile) DBCreateIgnoreDuplicate(ctx context.Context, db DB) (sql.Result, error) {
	q := "INSERT INTO `commit_file` (`commit_file`.`id`,`commit_file`.`checksum`,`commit_file`.`commit_id`,`commit_file`.`repo_id`,`commit_file`.`author_user_id`,`commit_file`.`committer_user_id`,`commit_file`.`filename`,`commit_file`.`language`,`commit_file`.`additions`,`commit_file`.`deletions`,`commit_file`.`size`,`commit_file`.`abinary`,`commit_file`.`date`,`commit_file`.`branch`,`commit_file`.`mergecommit`,`commit_file`.`excluded`,`commit_file`.`loc`,`commit_file`.`sloc`,`commit_file`.`comments`,`commit_file`.`blanks`,`commit_file`.`variance`,`commit_file`.`status`,`commit_file`.`renamed`,`commit_file`.`renamed_from`,`commit_file`.`renamed_to`,`commit_file`.`customer_id`,`commit_file`.`ref_type`,`commit_file`.`ref_id`,`commit_file`.`metadata`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CommitID),
		orm.ToSQLString(t.RepoID),
		orm.ToSQLString(t.AuthorUserID),
		orm.ToSQLString(t.CommitterUserID),
		orm.ToSQLString(t.Filename),
		orm.ToSQLString(t.Language),
		orm.ToSQLInt64(t.Additions),
		orm.ToSQLInt64(t.Deletions),
		orm.ToSQLInt64(t.Size),
		orm.ToSQLBool(t.Abinary),
		orm.ToSQLInt64(t.Date),
		orm.ToSQLString(t.Branch),
		orm.ToSQLBool(t.Mergecommit),
		orm.ToSQLBool(t.Excluded),
		orm.ToSQLInt64(t.Loc),
		orm.ToSQLInt64(t.Sloc),
		orm.ToSQLInt64(t.Comments),
		orm.ToSQLInt64(t.Blanks),
		orm.ToSQLInt64(t.Variance),
		orm.ToSQLString(enumCommitFileStatusToString(&t.Status)),
		orm.ToSQLBool(t.Renamed),
		orm.ToSQLString(t.RenamedFrom),
		orm.ToSQLString(t.RenamedTo),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefType),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.Metadata),
	)
}

// DBCreateIgnoreDuplicateTx will upsert the CommitFile record in the database using the provided transaction
func (t *CommitFile) DBCreateIgnoreDuplicateTx(ctx context.Context, tx Tx) (sql.Result, error) {
	q := "INSERT INTO `commit_file` (`commit_file`.`id`,`commit_file`.`checksum`,`commit_file`.`commit_id`,`commit_file`.`repo_id`,`commit_file`.`author_user_id`,`commit_file`.`committer_user_id`,`commit_file`.`filename`,`commit_file`.`language`,`commit_file`.`additions`,`commit_file`.`deletions`,`commit_file`.`size`,`commit_file`.`abinary`,`commit_file`.`date`,`commit_file`.`branch`,`commit_file`.`mergecommit`,`commit_file`.`excluded`,`commit_file`.`loc`,`commit_file`.`sloc`,`commit_file`.`comments`,`commit_file`.`blanks`,`commit_file`.`variance`,`commit_file`.`status`,`commit_file`.`renamed`,`commit_file`.`renamed_from`,`commit_file`.`renamed_to`,`commit_file`.`customer_id`,`commit_file`.`ref_type`,`commit_file`.`ref_id`,`commit_file`.`metadata`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CommitID),
		orm.ToSQLString(t.RepoID),
		orm.ToSQLString(t.AuthorUserID),
		orm.ToSQLString(t.CommitterUserID),
		orm.ToSQLString(t.Filename),
		orm.ToSQLString(t.Language),
		orm.ToSQLInt64(t.Additions),
		orm.ToSQLInt64(t.Deletions),
		orm.ToSQLInt64(t.Size),
		orm.ToSQLBool(t.Abinary),
		orm.ToSQLInt64(t.Date),
		orm.ToSQLString(t.Branch),
		orm.ToSQLBool(t.Mergecommit),
		orm.ToSQLBool(t.Excluded),
		orm.ToSQLInt64(t.Loc),
		orm.ToSQLInt64(t.Sloc),
		orm.ToSQLInt64(t.Comments),
		orm.ToSQLInt64(t.Blanks),
		orm.ToSQLInt64(t.Variance),
		orm.ToSQLString(enumCommitFileStatusToString(&t.Status)),
		orm.ToSQLBool(t.Renamed),
		orm.ToSQLString(t.RenamedFrom),
		orm.ToSQLString(t.RenamedTo),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefType),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.Metadata),
	)
}

// DeleteAllCommitFiles deletes all CommitFile records in the database with optional filters
func DeleteAllCommitFiles(ctx context.Context, db DB, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(CommitFileTableName),
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

// DeleteAllCommitFilesTx deletes all CommitFile records in the database with optional filters using the provided transaction
func DeleteAllCommitFilesTx(ctx context.Context, tx Tx, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(CommitFileTableName),
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

// DBDelete will delete this CommitFile record in the database
func (t *CommitFile) DBDelete(ctx context.Context, db DB) (bool, error) {
	q := "DELETE FROM `commit_file` WHERE `id` = ?"
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

// DBDeleteTx will delete this CommitFile record in the database using the provided transaction
func (t *CommitFile) DBDeleteTx(ctx context.Context, tx Tx) (bool, error) {
	q := "DELETE FROM `commit_file` WHERE `id` = ?"
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

// DBUpdate will update the CommitFile record in the database
func (t *CommitFile) DBUpdate(ctx context.Context, db DB) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `commit_file` SET `checksum`=?,`commit_id`=?,`repo_id`=?,`author_user_id`=?,`committer_user_id`=?,`filename`=?,`language`=?,`additions`=?,`deletions`=?,`size`=?,`abinary`=?,`date`=?,`branch`=?,`mergecommit`=?,`excluded`=?,`loc`=?,`sloc`=?,`comments`=?,`blanks`=?,`variance`=?,`status`=?,`renamed`=?,`renamed_from`=?,`renamed_to`=?,`customer_id`=?,`ref_type`=?,`ref_id`=?,`metadata`=? WHERE `id`=?"
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CommitID),
		orm.ToSQLString(t.RepoID),
		orm.ToSQLString(t.AuthorUserID),
		orm.ToSQLString(t.CommitterUserID),
		orm.ToSQLString(t.Filename),
		orm.ToSQLString(t.Language),
		orm.ToSQLInt64(t.Additions),
		orm.ToSQLInt64(t.Deletions),
		orm.ToSQLInt64(t.Size),
		orm.ToSQLBool(t.Abinary),
		orm.ToSQLInt64(t.Date),
		orm.ToSQLString(t.Branch),
		orm.ToSQLBool(t.Mergecommit),
		orm.ToSQLBool(t.Excluded),
		orm.ToSQLInt64(t.Loc),
		orm.ToSQLInt64(t.Sloc),
		orm.ToSQLInt64(t.Comments),
		orm.ToSQLInt64(t.Blanks),
		orm.ToSQLInt64(t.Variance),
		orm.ToSQLString(enumCommitFileStatusToString(&t.Status)),
		orm.ToSQLBool(t.Renamed),
		orm.ToSQLString(t.RenamedFrom),
		orm.ToSQLString(t.RenamedTo),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefType),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.Metadata),
		orm.ToSQLString(t.ID),
	)
}

// DBUpdateTx will update the CommitFile record in the database using the provided transaction
func (t *CommitFile) DBUpdateTx(ctx context.Context, tx Tx) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `commit_file` SET `checksum`=?,`commit_id`=?,`repo_id`=?,`author_user_id`=?,`committer_user_id`=?,`filename`=?,`language`=?,`additions`=?,`deletions`=?,`size`=?,`abinary`=?,`date`=?,`branch`=?,`mergecommit`=?,`excluded`=?,`loc`=?,`sloc`=?,`comments`=?,`blanks`=?,`variance`=?,`status`=?,`renamed`=?,`renamed_from`=?,`renamed_to`=?,`customer_id`=?,`ref_type`=?,`ref_id`=?,`metadata`=? WHERE `id`=?"
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CommitID),
		orm.ToSQLString(t.RepoID),
		orm.ToSQLString(t.AuthorUserID),
		orm.ToSQLString(t.CommitterUserID),
		orm.ToSQLString(t.Filename),
		orm.ToSQLString(t.Language),
		orm.ToSQLInt64(t.Additions),
		orm.ToSQLInt64(t.Deletions),
		orm.ToSQLInt64(t.Size),
		orm.ToSQLBool(t.Abinary),
		orm.ToSQLInt64(t.Date),
		orm.ToSQLString(t.Branch),
		orm.ToSQLBool(t.Mergecommit),
		orm.ToSQLBool(t.Excluded),
		orm.ToSQLInt64(t.Loc),
		orm.ToSQLInt64(t.Sloc),
		orm.ToSQLInt64(t.Comments),
		orm.ToSQLInt64(t.Blanks),
		orm.ToSQLInt64(t.Variance),
		orm.ToSQLString(enumCommitFileStatusToString(&t.Status)),
		orm.ToSQLBool(t.Renamed),
		orm.ToSQLString(t.RenamedFrom),
		orm.ToSQLString(t.RenamedTo),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefType),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.Metadata),
		orm.ToSQLString(t.ID),
	)
}

// DBUpsert will upsert the CommitFile record in the database
func (t *CommitFile) DBUpsert(ctx context.Context, db DB, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `commit_file` (`commit_file`.`id`,`commit_file`.`checksum`,`commit_file`.`commit_id`,`commit_file`.`repo_id`,`commit_file`.`author_user_id`,`commit_file`.`committer_user_id`,`commit_file`.`filename`,`commit_file`.`language`,`commit_file`.`additions`,`commit_file`.`deletions`,`commit_file`.`size`,`commit_file`.`abinary`,`commit_file`.`date`,`commit_file`.`branch`,`commit_file`.`mergecommit`,`commit_file`.`excluded`,`commit_file`.`loc`,`commit_file`.`sloc`,`commit_file`.`comments`,`commit_file`.`blanks`,`commit_file`.`variance`,`commit_file`.`status`,`commit_file`.`renamed`,`commit_file`.`renamed_from`,`commit_file`.`renamed_to`,`commit_file`.`customer_id`,`commit_file`.`ref_type`,`commit_file`.`ref_id`,`commit_file`.`metadata`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `commit_file` (`commit_file`.`id`,`commit_file`.`checksum`,`commit_file`.`commit_id`,`commit_file`.`repo_id`,`commit_file`.`author_user_id`,`commit_file`.`committer_user_id`,`commit_file`.`filename`,`commit_file`.`language`,`commit_file`.`additions`,`commit_file`.`deletions`,`commit_file`.`size`,`commit_file`.`abinary`,`commit_file`.`date`,`commit_file`.`branch`,`commit_file`.`mergecommit`,`commit_file`.`excluded`,`commit_file`.`loc`,`commit_file`.`sloc`,`commit_file`.`comments`,`commit_file`.`blanks`,`commit_file`.`variance`,`commit_file`.`status`,`commit_file`.`renamed`,`commit_file`.`renamed_from`,`commit_file`.`renamed_to`,`commit_file`.`customer_id`,`commit_file`.`ref_type`,`commit_file`.`ref_id`,`commit_file`.`metadata`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`commit_id`=VALUES(`commit_id`),`repo_id`=VALUES(`repo_id`),`author_user_id`=VALUES(`author_user_id`),`committer_user_id`=VALUES(`committer_user_id`),`filename`=VALUES(`filename`),`language`=VALUES(`language`),`additions`=VALUES(`additions`),`deletions`=VALUES(`deletions`),`size`=VALUES(`size`),`abinary`=VALUES(`abinary`),`date`=VALUES(`date`),`branch`=VALUES(`branch`),`mergecommit`=VALUES(`mergecommit`),`excluded`=VALUES(`excluded`),`loc`=VALUES(`loc`),`sloc`=VALUES(`sloc`),`comments`=VALUES(`comments`),`blanks`=VALUES(`blanks`),`variance`=VALUES(`variance`),`status`=VALUES(`status`),`renamed`=VALUES(`renamed`),`renamed_from`=VALUES(`renamed_from`),`renamed_to`=VALUES(`renamed_to`),`customer_id`=VALUES(`customer_id`),`ref_type`=VALUES(`ref_type`),`ref_id`=VALUES(`ref_id`),`metadata`=VALUES(`metadata`)"
	}
	r, err := db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CommitID),
		orm.ToSQLString(t.RepoID),
		orm.ToSQLString(t.AuthorUserID),
		orm.ToSQLString(t.CommitterUserID),
		orm.ToSQLString(t.Filename),
		orm.ToSQLString(t.Language),
		orm.ToSQLInt64(t.Additions),
		orm.ToSQLInt64(t.Deletions),
		orm.ToSQLInt64(t.Size),
		orm.ToSQLBool(t.Abinary),
		orm.ToSQLInt64(t.Date),
		orm.ToSQLString(t.Branch),
		orm.ToSQLBool(t.Mergecommit),
		orm.ToSQLBool(t.Excluded),
		orm.ToSQLInt64(t.Loc),
		orm.ToSQLInt64(t.Sloc),
		orm.ToSQLInt64(t.Comments),
		orm.ToSQLInt64(t.Blanks),
		orm.ToSQLInt64(t.Variance),
		orm.ToSQLString(enumCommitFileStatusToString(&t.Status)),
		orm.ToSQLBool(t.Renamed),
		orm.ToSQLString(t.RenamedFrom),
		orm.ToSQLString(t.RenamedTo),
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

// DBUpsertTx will upsert the CommitFile record in the database using the provided transaction
func (t *CommitFile) DBUpsertTx(ctx context.Context, tx Tx, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `commit_file` (`commit_file`.`id`,`commit_file`.`checksum`,`commit_file`.`commit_id`,`commit_file`.`repo_id`,`commit_file`.`author_user_id`,`commit_file`.`committer_user_id`,`commit_file`.`filename`,`commit_file`.`language`,`commit_file`.`additions`,`commit_file`.`deletions`,`commit_file`.`size`,`commit_file`.`abinary`,`commit_file`.`date`,`commit_file`.`branch`,`commit_file`.`mergecommit`,`commit_file`.`excluded`,`commit_file`.`loc`,`commit_file`.`sloc`,`commit_file`.`comments`,`commit_file`.`blanks`,`commit_file`.`variance`,`commit_file`.`status`,`commit_file`.`renamed`,`commit_file`.`renamed_from`,`commit_file`.`renamed_to`,`commit_file`.`customer_id`,`commit_file`.`ref_type`,`commit_file`.`ref_id`,`commit_file`.`metadata`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `commit_file` (`commit_file`.`id`,`commit_file`.`checksum`,`commit_file`.`commit_id`,`commit_file`.`repo_id`,`commit_file`.`author_user_id`,`commit_file`.`committer_user_id`,`commit_file`.`filename`,`commit_file`.`language`,`commit_file`.`additions`,`commit_file`.`deletions`,`commit_file`.`size`,`commit_file`.`abinary`,`commit_file`.`date`,`commit_file`.`branch`,`commit_file`.`mergecommit`,`commit_file`.`excluded`,`commit_file`.`loc`,`commit_file`.`sloc`,`commit_file`.`comments`,`commit_file`.`blanks`,`commit_file`.`variance`,`commit_file`.`status`,`commit_file`.`renamed`,`commit_file`.`renamed_from`,`commit_file`.`renamed_to`,`commit_file`.`customer_id`,`commit_file`.`ref_type`,`commit_file`.`ref_id`,`commit_file`.`metadata`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`commit_id`=VALUES(`commit_id`),`repo_id`=VALUES(`repo_id`),`author_user_id`=VALUES(`author_user_id`),`committer_user_id`=VALUES(`committer_user_id`),`filename`=VALUES(`filename`),`language`=VALUES(`language`),`additions`=VALUES(`additions`),`deletions`=VALUES(`deletions`),`size`=VALUES(`size`),`abinary`=VALUES(`abinary`),`date`=VALUES(`date`),`branch`=VALUES(`branch`),`mergecommit`=VALUES(`mergecommit`),`excluded`=VALUES(`excluded`),`loc`=VALUES(`loc`),`sloc`=VALUES(`sloc`),`comments`=VALUES(`comments`),`blanks`=VALUES(`blanks`),`variance`=VALUES(`variance`),`status`=VALUES(`status`),`renamed`=VALUES(`renamed`),`renamed_from`=VALUES(`renamed_from`),`renamed_to`=VALUES(`renamed_to`),`customer_id`=VALUES(`customer_id`),`ref_type`=VALUES(`ref_type`),`ref_id`=VALUES(`ref_id`),`metadata`=VALUES(`metadata`)"
	}
	r, err := tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CommitID),
		orm.ToSQLString(t.RepoID),
		orm.ToSQLString(t.AuthorUserID),
		orm.ToSQLString(t.CommitterUserID),
		orm.ToSQLString(t.Filename),
		orm.ToSQLString(t.Language),
		orm.ToSQLInt64(t.Additions),
		orm.ToSQLInt64(t.Deletions),
		orm.ToSQLInt64(t.Size),
		orm.ToSQLBool(t.Abinary),
		orm.ToSQLInt64(t.Date),
		orm.ToSQLString(t.Branch),
		orm.ToSQLBool(t.Mergecommit),
		orm.ToSQLBool(t.Excluded),
		orm.ToSQLInt64(t.Loc),
		orm.ToSQLInt64(t.Sloc),
		orm.ToSQLInt64(t.Comments),
		orm.ToSQLInt64(t.Blanks),
		orm.ToSQLInt64(t.Variance),
		orm.ToSQLString(enumCommitFileStatusToString(&t.Status)),
		orm.ToSQLBool(t.Renamed),
		orm.ToSQLString(t.RenamedFrom),
		orm.ToSQLString(t.RenamedTo),
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

// DBFindOne will find a CommitFile record in the database with the primary key
func (t *CommitFile) DBFindOne(ctx context.Context, db DB, value string) (bool, error) {
	q := "SELECT * FROM `commit_file` WHERE `id` = ? LIMIT 1"
	row := db.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CommitID sql.NullString
	var _RepoID sql.NullString
	var _AuthorUserID sql.NullString
	var _CommitterUserID sql.NullString
	var _Filename sql.NullString
	var _Language sql.NullString
	var _Additions sql.NullInt64
	var _Deletions sql.NullInt64
	var _Size sql.NullInt64
	var _Abinary sql.NullBool
	var _Date sql.NullInt64
	var _Branch sql.NullString
	var _Mergecommit sql.NullBool
	var _Excluded sql.NullBool
	var _Loc sql.NullInt64
	var _Sloc sql.NullInt64
	var _Comments sql.NullInt64
	var _Blanks sql.NullInt64
	var _Variance sql.NullInt64
	var _Status sql.NullString
	var _Renamed sql.NullBool
	var _RenamedFrom sql.NullString
	var _RenamedTo sql.NullString
	var _CustomerID sql.NullString
	var _RefType sql.NullString
	var _RefID sql.NullString
	var _Metadata sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CommitID,
		&_RepoID,
		&_AuthorUserID,
		&_CommitterUserID,
		&_Filename,
		&_Language,
		&_Additions,
		&_Deletions,
		&_Size,
		&_Abinary,
		&_Date,
		&_Branch,
		&_Mergecommit,
		&_Excluded,
		&_Loc,
		&_Sloc,
		&_Comments,
		&_Blanks,
		&_Variance,
		&_Status,
		&_Renamed,
		&_RenamedFrom,
		&_RenamedTo,
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
	if _CommitID.Valid {
		t.SetCommitID(_CommitID.String)
	}
	if _RepoID.Valid {
		t.SetRepoID(_RepoID.String)
	}
	if _AuthorUserID.Valid {
		t.SetAuthorUserID(_AuthorUserID.String)
	}
	if _CommitterUserID.Valid {
		t.SetCommitterUserID(_CommitterUserID.String)
	}
	if _Filename.Valid {
		t.SetFilename(_Filename.String)
	}
	if _Language.Valid {
		t.SetLanguage(_Language.String)
	}
	if _Additions.Valid {
		t.SetAdditions(int32(_Additions.Int64))
	}
	if _Deletions.Valid {
		t.SetDeletions(int32(_Deletions.Int64))
	}
	if _Size.Valid {
		t.SetSize(int32(_Size.Int64))
	}
	if _Abinary.Valid {
		t.SetAbinary(_Abinary.Bool)
	}
	if _Date.Valid {
		t.SetDate(_Date.Int64)
	}
	if _Branch.Valid {
		t.SetBranch(_Branch.String)
	}
	if _Mergecommit.Valid {
		t.SetMergecommit(_Mergecommit.Bool)
	}
	if _Excluded.Valid {
		t.SetExcluded(_Excluded.Bool)
	}
	if _Loc.Valid {
		t.SetLoc(int32(_Loc.Int64))
	}
	if _Sloc.Valid {
		t.SetSloc(int32(_Sloc.Int64))
	}
	if _Comments.Valid {
		t.SetComments(int32(_Comments.Int64))
	}
	if _Blanks.Valid {
		t.SetBlanks(int32(_Blanks.Int64))
	}
	if _Variance.Valid {
		t.SetVariance(int32(_Variance.Int64))
	}
	if _Status.Valid {
		t.SetStatusString(_Status.String)
	}
	if _Renamed.Valid {
		t.SetRenamed(_Renamed.Bool)
	}
	if _RenamedFrom.Valid {
		t.SetRenamedFrom(_RenamedFrom.String)
	}
	if _RenamedTo.Valid {
		t.SetRenamedTo(_RenamedTo.String)
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

// DBFindOneTx will find a CommitFile record in the database with the primary key using the provided transaction
func (t *CommitFile) DBFindOneTx(ctx context.Context, tx Tx, value string) (bool, error) {
	q := "SELECT * FROM `commit_file` WHERE `id` = ? LIMIT 1"
	row := tx.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CommitID sql.NullString
	var _RepoID sql.NullString
	var _AuthorUserID sql.NullString
	var _CommitterUserID sql.NullString
	var _Filename sql.NullString
	var _Language sql.NullString
	var _Additions sql.NullInt64
	var _Deletions sql.NullInt64
	var _Size sql.NullInt64
	var _Abinary sql.NullBool
	var _Date sql.NullInt64
	var _Branch sql.NullString
	var _Mergecommit sql.NullBool
	var _Excluded sql.NullBool
	var _Loc sql.NullInt64
	var _Sloc sql.NullInt64
	var _Comments sql.NullInt64
	var _Blanks sql.NullInt64
	var _Variance sql.NullInt64
	var _Status sql.NullString
	var _Renamed sql.NullBool
	var _RenamedFrom sql.NullString
	var _RenamedTo sql.NullString
	var _CustomerID sql.NullString
	var _RefType sql.NullString
	var _RefID sql.NullString
	var _Metadata sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CommitID,
		&_RepoID,
		&_AuthorUserID,
		&_CommitterUserID,
		&_Filename,
		&_Language,
		&_Additions,
		&_Deletions,
		&_Size,
		&_Abinary,
		&_Date,
		&_Branch,
		&_Mergecommit,
		&_Excluded,
		&_Loc,
		&_Sloc,
		&_Comments,
		&_Blanks,
		&_Variance,
		&_Status,
		&_Renamed,
		&_RenamedFrom,
		&_RenamedTo,
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
	if _CommitID.Valid {
		t.SetCommitID(_CommitID.String)
	}
	if _RepoID.Valid {
		t.SetRepoID(_RepoID.String)
	}
	if _AuthorUserID.Valid {
		t.SetAuthorUserID(_AuthorUserID.String)
	}
	if _CommitterUserID.Valid {
		t.SetCommitterUserID(_CommitterUserID.String)
	}
	if _Filename.Valid {
		t.SetFilename(_Filename.String)
	}
	if _Language.Valid {
		t.SetLanguage(_Language.String)
	}
	if _Additions.Valid {
		t.SetAdditions(int32(_Additions.Int64))
	}
	if _Deletions.Valid {
		t.SetDeletions(int32(_Deletions.Int64))
	}
	if _Size.Valid {
		t.SetSize(int32(_Size.Int64))
	}
	if _Abinary.Valid {
		t.SetAbinary(_Abinary.Bool)
	}
	if _Date.Valid {
		t.SetDate(_Date.Int64)
	}
	if _Branch.Valid {
		t.SetBranch(_Branch.String)
	}
	if _Mergecommit.Valid {
		t.SetMergecommit(_Mergecommit.Bool)
	}
	if _Excluded.Valid {
		t.SetExcluded(_Excluded.Bool)
	}
	if _Loc.Valid {
		t.SetLoc(int32(_Loc.Int64))
	}
	if _Sloc.Valid {
		t.SetSloc(int32(_Sloc.Int64))
	}
	if _Comments.Valid {
		t.SetComments(int32(_Comments.Int64))
	}
	if _Blanks.Valid {
		t.SetBlanks(int32(_Blanks.Int64))
	}
	if _Variance.Valid {
		t.SetVariance(int32(_Variance.Int64))
	}
	if _Status.Valid {
		t.SetStatusString(_Status.String)
	}
	if _Renamed.Valid {
		t.SetRenamed(_Renamed.Bool)
	}
	if _RenamedFrom.Valid {
		t.SetRenamedFrom(_RenamedFrom.String)
	}
	if _RenamedTo.Valid {
		t.SetRenamedTo(_RenamedTo.String)
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

// FindCommitFiles will find a CommitFile record in the database with the provided parameters
func FindCommitFiles(ctx context.Context, db DB, _params ...interface{}) ([]*CommitFile, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("commit_id"),
		orm.Column("repo_id"),
		orm.Column("author_user_id"),
		orm.Column("committer_user_id"),
		orm.Column("filename"),
		orm.Column("language"),
		orm.Column("additions"),
		orm.Column("deletions"),
		orm.Column("size"),
		orm.Column("abinary"),
		orm.Column("date"),
		orm.Column("branch"),
		orm.Column("mergecommit"),
		orm.Column("excluded"),
		orm.Column("loc"),
		orm.Column("sloc"),
		orm.Column("comments"),
		orm.Column("blanks"),
		orm.Column("variance"),
		orm.Column("status"),
		orm.Column("renamed"),
		orm.Column("renamed_from"),
		orm.Column("renamed_to"),
		orm.Column("customer_id"),
		orm.Column("ref_type"),
		orm.Column("ref_id"),
		orm.Column("metadata"),
		orm.Table(CommitFileTableName),
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
	results := make([]*CommitFile, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CommitID sql.NullString
		var _RepoID sql.NullString
		var _AuthorUserID sql.NullString
		var _CommitterUserID sql.NullString
		var _Filename sql.NullString
		var _Language sql.NullString
		var _Additions sql.NullInt64
		var _Deletions sql.NullInt64
		var _Size sql.NullInt64
		var _Abinary sql.NullBool
		var _Date sql.NullInt64
		var _Branch sql.NullString
		var _Mergecommit sql.NullBool
		var _Excluded sql.NullBool
		var _Loc sql.NullInt64
		var _Sloc sql.NullInt64
		var _Comments sql.NullInt64
		var _Blanks sql.NullInt64
		var _Variance sql.NullInt64
		var _Status sql.NullString
		var _Renamed sql.NullBool
		var _RenamedFrom sql.NullString
		var _RenamedTo sql.NullString
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefID sql.NullString
		var _Metadata sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CommitID,
			&_RepoID,
			&_AuthorUserID,
			&_CommitterUserID,
			&_Filename,
			&_Language,
			&_Additions,
			&_Deletions,
			&_Size,
			&_Abinary,
			&_Date,
			&_Branch,
			&_Mergecommit,
			&_Excluded,
			&_Loc,
			&_Sloc,
			&_Comments,
			&_Blanks,
			&_Variance,
			&_Status,
			&_Renamed,
			&_RenamedFrom,
			&_RenamedTo,
			&_CustomerID,
			&_RefType,
			&_RefID,
			&_Metadata,
		)
		if err != nil {
			return nil, err
		}
		t := &CommitFile{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
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
		if _CommitterUserID.Valid {
			t.SetCommitterUserID(_CommitterUserID.String)
		}
		if _Filename.Valid {
			t.SetFilename(_Filename.String)
		}
		if _Language.Valid {
			t.SetLanguage(_Language.String)
		}
		if _Additions.Valid {
			t.SetAdditions(int32(_Additions.Int64))
		}
		if _Deletions.Valid {
			t.SetDeletions(int32(_Deletions.Int64))
		}
		if _Size.Valid {
			t.SetSize(int32(_Size.Int64))
		}
		if _Abinary.Valid {
			t.SetAbinary(_Abinary.Bool)
		}
		if _Date.Valid {
			t.SetDate(_Date.Int64)
		}
		if _Branch.Valid {
			t.SetBranch(_Branch.String)
		}
		if _Mergecommit.Valid {
			t.SetMergecommit(_Mergecommit.Bool)
		}
		if _Excluded.Valid {
			t.SetExcluded(_Excluded.Bool)
		}
		if _Loc.Valid {
			t.SetLoc(int32(_Loc.Int64))
		}
		if _Sloc.Valid {
			t.SetSloc(int32(_Sloc.Int64))
		}
		if _Comments.Valid {
			t.SetComments(int32(_Comments.Int64))
		}
		if _Blanks.Valid {
			t.SetBlanks(int32(_Blanks.Int64))
		}
		if _Variance.Valid {
			t.SetVariance(int32(_Variance.Int64))
		}
		if _Status.Valid {
			t.SetStatusString(_Status.String)
		}
		if _Renamed.Valid {
			t.SetRenamed(_Renamed.Bool)
		}
		if _RenamedFrom.Valid {
			t.SetRenamedFrom(_RenamedFrom.String)
		}
		if _RenamedTo.Valid {
			t.SetRenamedTo(_RenamedTo.String)
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

// FindCommitFilesTx will find a CommitFile record in the database with the provided parameters using the provided transaction
func FindCommitFilesTx(ctx context.Context, tx Tx, _params ...interface{}) ([]*CommitFile, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("commit_id"),
		orm.Column("repo_id"),
		orm.Column("author_user_id"),
		orm.Column("committer_user_id"),
		orm.Column("filename"),
		orm.Column("language"),
		orm.Column("additions"),
		orm.Column("deletions"),
		orm.Column("size"),
		orm.Column("abinary"),
		orm.Column("date"),
		orm.Column("branch"),
		orm.Column("mergecommit"),
		orm.Column("excluded"),
		orm.Column("loc"),
		orm.Column("sloc"),
		orm.Column("comments"),
		orm.Column("blanks"),
		orm.Column("variance"),
		orm.Column("status"),
		orm.Column("renamed"),
		orm.Column("renamed_from"),
		orm.Column("renamed_to"),
		orm.Column("customer_id"),
		orm.Column("ref_type"),
		orm.Column("ref_id"),
		orm.Column("metadata"),
		orm.Table(CommitFileTableName),
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
	results := make([]*CommitFile, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CommitID sql.NullString
		var _RepoID sql.NullString
		var _AuthorUserID sql.NullString
		var _CommitterUserID sql.NullString
		var _Filename sql.NullString
		var _Language sql.NullString
		var _Additions sql.NullInt64
		var _Deletions sql.NullInt64
		var _Size sql.NullInt64
		var _Abinary sql.NullBool
		var _Date sql.NullInt64
		var _Branch sql.NullString
		var _Mergecommit sql.NullBool
		var _Excluded sql.NullBool
		var _Loc sql.NullInt64
		var _Sloc sql.NullInt64
		var _Comments sql.NullInt64
		var _Blanks sql.NullInt64
		var _Variance sql.NullInt64
		var _Status sql.NullString
		var _Renamed sql.NullBool
		var _RenamedFrom sql.NullString
		var _RenamedTo sql.NullString
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefID sql.NullString
		var _Metadata sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CommitID,
			&_RepoID,
			&_AuthorUserID,
			&_CommitterUserID,
			&_Filename,
			&_Language,
			&_Additions,
			&_Deletions,
			&_Size,
			&_Abinary,
			&_Date,
			&_Branch,
			&_Mergecommit,
			&_Excluded,
			&_Loc,
			&_Sloc,
			&_Comments,
			&_Blanks,
			&_Variance,
			&_Status,
			&_Renamed,
			&_RenamedFrom,
			&_RenamedTo,
			&_CustomerID,
			&_RefType,
			&_RefID,
			&_Metadata,
		)
		if err != nil {
			return nil, err
		}
		t := &CommitFile{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
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
		if _CommitterUserID.Valid {
			t.SetCommitterUserID(_CommitterUserID.String)
		}
		if _Filename.Valid {
			t.SetFilename(_Filename.String)
		}
		if _Language.Valid {
			t.SetLanguage(_Language.String)
		}
		if _Additions.Valid {
			t.SetAdditions(int32(_Additions.Int64))
		}
		if _Deletions.Valid {
			t.SetDeletions(int32(_Deletions.Int64))
		}
		if _Size.Valid {
			t.SetSize(int32(_Size.Int64))
		}
		if _Abinary.Valid {
			t.SetAbinary(_Abinary.Bool)
		}
		if _Date.Valid {
			t.SetDate(_Date.Int64)
		}
		if _Branch.Valid {
			t.SetBranch(_Branch.String)
		}
		if _Mergecommit.Valid {
			t.SetMergecommit(_Mergecommit.Bool)
		}
		if _Excluded.Valid {
			t.SetExcluded(_Excluded.Bool)
		}
		if _Loc.Valid {
			t.SetLoc(int32(_Loc.Int64))
		}
		if _Sloc.Valid {
			t.SetSloc(int32(_Sloc.Int64))
		}
		if _Comments.Valid {
			t.SetComments(int32(_Comments.Int64))
		}
		if _Blanks.Valid {
			t.SetBlanks(int32(_Blanks.Int64))
		}
		if _Variance.Valid {
			t.SetVariance(int32(_Variance.Int64))
		}
		if _Status.Valid {
			t.SetStatusString(_Status.String)
		}
		if _Renamed.Valid {
			t.SetRenamed(_Renamed.Bool)
		}
		if _RenamedFrom.Valid {
			t.SetRenamedFrom(_RenamedFrom.String)
		}
		if _RenamedTo.Valid {
			t.SetRenamedTo(_RenamedTo.String)
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

// DBFind will find a CommitFile record in the database with the provided parameters
func (t *CommitFile) DBFind(ctx context.Context, db DB, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("commit_id"),
		orm.Column("repo_id"),
		orm.Column("author_user_id"),
		orm.Column("committer_user_id"),
		orm.Column("filename"),
		orm.Column("language"),
		orm.Column("additions"),
		orm.Column("deletions"),
		orm.Column("size"),
		orm.Column("abinary"),
		orm.Column("date"),
		orm.Column("branch"),
		orm.Column("mergecommit"),
		orm.Column("excluded"),
		orm.Column("loc"),
		orm.Column("sloc"),
		orm.Column("comments"),
		orm.Column("blanks"),
		orm.Column("variance"),
		orm.Column("status"),
		orm.Column("renamed"),
		orm.Column("renamed_from"),
		orm.Column("renamed_to"),
		orm.Column("customer_id"),
		orm.Column("ref_type"),
		orm.Column("ref_id"),
		orm.Column("metadata"),
		orm.Table(CommitFileTableName),
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
	var _RepoID sql.NullString
	var _AuthorUserID sql.NullString
	var _CommitterUserID sql.NullString
	var _Filename sql.NullString
	var _Language sql.NullString
	var _Additions sql.NullInt64
	var _Deletions sql.NullInt64
	var _Size sql.NullInt64
	var _Abinary sql.NullBool
	var _Date sql.NullInt64
	var _Branch sql.NullString
	var _Mergecommit sql.NullBool
	var _Excluded sql.NullBool
	var _Loc sql.NullInt64
	var _Sloc sql.NullInt64
	var _Comments sql.NullInt64
	var _Blanks sql.NullInt64
	var _Variance sql.NullInt64
	var _Status sql.NullString
	var _Renamed sql.NullBool
	var _RenamedFrom sql.NullString
	var _RenamedTo sql.NullString
	var _CustomerID sql.NullString
	var _RefType sql.NullString
	var _RefID sql.NullString
	var _Metadata sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CommitID,
		&_RepoID,
		&_AuthorUserID,
		&_CommitterUserID,
		&_Filename,
		&_Language,
		&_Additions,
		&_Deletions,
		&_Size,
		&_Abinary,
		&_Date,
		&_Branch,
		&_Mergecommit,
		&_Excluded,
		&_Loc,
		&_Sloc,
		&_Comments,
		&_Blanks,
		&_Variance,
		&_Status,
		&_Renamed,
		&_RenamedFrom,
		&_RenamedTo,
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
	if _CommitID.Valid {
		t.SetCommitID(_CommitID.String)
	}
	if _RepoID.Valid {
		t.SetRepoID(_RepoID.String)
	}
	if _AuthorUserID.Valid {
		t.SetAuthorUserID(_AuthorUserID.String)
	}
	if _CommitterUserID.Valid {
		t.SetCommitterUserID(_CommitterUserID.String)
	}
	if _Filename.Valid {
		t.SetFilename(_Filename.String)
	}
	if _Language.Valid {
		t.SetLanguage(_Language.String)
	}
	if _Additions.Valid {
		t.SetAdditions(int32(_Additions.Int64))
	}
	if _Deletions.Valid {
		t.SetDeletions(int32(_Deletions.Int64))
	}
	if _Size.Valid {
		t.SetSize(int32(_Size.Int64))
	}
	if _Abinary.Valid {
		t.SetAbinary(_Abinary.Bool)
	}
	if _Date.Valid {
		t.SetDate(_Date.Int64)
	}
	if _Branch.Valid {
		t.SetBranch(_Branch.String)
	}
	if _Mergecommit.Valid {
		t.SetMergecommit(_Mergecommit.Bool)
	}
	if _Excluded.Valid {
		t.SetExcluded(_Excluded.Bool)
	}
	if _Loc.Valid {
		t.SetLoc(int32(_Loc.Int64))
	}
	if _Sloc.Valid {
		t.SetSloc(int32(_Sloc.Int64))
	}
	if _Comments.Valid {
		t.SetComments(int32(_Comments.Int64))
	}
	if _Blanks.Valid {
		t.SetBlanks(int32(_Blanks.Int64))
	}
	if _Variance.Valid {
		t.SetVariance(int32(_Variance.Int64))
	}
	if _Status.Valid {
		t.SetStatusString(_Status.String)
	}
	if _Renamed.Valid {
		t.SetRenamed(_Renamed.Bool)
	}
	if _RenamedFrom.Valid {
		t.SetRenamedFrom(_RenamedFrom.String)
	}
	if _RenamedTo.Valid {
		t.SetRenamedTo(_RenamedTo.String)
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

// DBFindTx will find a CommitFile record in the database with the provided parameters using the provided transaction
func (t *CommitFile) DBFindTx(ctx context.Context, tx Tx, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("commit_id"),
		orm.Column("repo_id"),
		orm.Column("author_user_id"),
		orm.Column("committer_user_id"),
		orm.Column("filename"),
		orm.Column("language"),
		orm.Column("additions"),
		orm.Column("deletions"),
		orm.Column("size"),
		orm.Column("abinary"),
		orm.Column("date"),
		orm.Column("branch"),
		orm.Column("mergecommit"),
		orm.Column("excluded"),
		orm.Column("loc"),
		orm.Column("sloc"),
		orm.Column("comments"),
		orm.Column("blanks"),
		orm.Column("variance"),
		orm.Column("status"),
		orm.Column("renamed"),
		orm.Column("renamed_from"),
		orm.Column("renamed_to"),
		orm.Column("customer_id"),
		orm.Column("ref_type"),
		orm.Column("ref_id"),
		orm.Column("metadata"),
		orm.Table(CommitFileTableName),
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
	var _RepoID sql.NullString
	var _AuthorUserID sql.NullString
	var _CommitterUserID sql.NullString
	var _Filename sql.NullString
	var _Language sql.NullString
	var _Additions sql.NullInt64
	var _Deletions sql.NullInt64
	var _Size sql.NullInt64
	var _Abinary sql.NullBool
	var _Date sql.NullInt64
	var _Branch sql.NullString
	var _Mergecommit sql.NullBool
	var _Excluded sql.NullBool
	var _Loc sql.NullInt64
	var _Sloc sql.NullInt64
	var _Comments sql.NullInt64
	var _Blanks sql.NullInt64
	var _Variance sql.NullInt64
	var _Status sql.NullString
	var _Renamed sql.NullBool
	var _RenamedFrom sql.NullString
	var _RenamedTo sql.NullString
	var _CustomerID sql.NullString
	var _RefType sql.NullString
	var _RefID sql.NullString
	var _Metadata sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CommitID,
		&_RepoID,
		&_AuthorUserID,
		&_CommitterUserID,
		&_Filename,
		&_Language,
		&_Additions,
		&_Deletions,
		&_Size,
		&_Abinary,
		&_Date,
		&_Branch,
		&_Mergecommit,
		&_Excluded,
		&_Loc,
		&_Sloc,
		&_Comments,
		&_Blanks,
		&_Variance,
		&_Status,
		&_Renamed,
		&_RenamedFrom,
		&_RenamedTo,
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
	if _CommitID.Valid {
		t.SetCommitID(_CommitID.String)
	}
	if _RepoID.Valid {
		t.SetRepoID(_RepoID.String)
	}
	if _AuthorUserID.Valid {
		t.SetAuthorUserID(_AuthorUserID.String)
	}
	if _CommitterUserID.Valid {
		t.SetCommitterUserID(_CommitterUserID.String)
	}
	if _Filename.Valid {
		t.SetFilename(_Filename.String)
	}
	if _Language.Valid {
		t.SetLanguage(_Language.String)
	}
	if _Additions.Valid {
		t.SetAdditions(int32(_Additions.Int64))
	}
	if _Deletions.Valid {
		t.SetDeletions(int32(_Deletions.Int64))
	}
	if _Size.Valid {
		t.SetSize(int32(_Size.Int64))
	}
	if _Abinary.Valid {
		t.SetAbinary(_Abinary.Bool)
	}
	if _Date.Valid {
		t.SetDate(_Date.Int64)
	}
	if _Branch.Valid {
		t.SetBranch(_Branch.String)
	}
	if _Mergecommit.Valid {
		t.SetMergecommit(_Mergecommit.Bool)
	}
	if _Excluded.Valid {
		t.SetExcluded(_Excluded.Bool)
	}
	if _Loc.Valid {
		t.SetLoc(int32(_Loc.Int64))
	}
	if _Sloc.Valid {
		t.SetSloc(int32(_Sloc.Int64))
	}
	if _Comments.Valid {
		t.SetComments(int32(_Comments.Int64))
	}
	if _Blanks.Valid {
		t.SetBlanks(int32(_Blanks.Int64))
	}
	if _Variance.Valid {
		t.SetVariance(int32(_Variance.Int64))
	}
	if _Status.Valid {
		t.SetStatusString(_Status.String)
	}
	if _Renamed.Valid {
		t.SetRenamed(_Renamed.Bool)
	}
	if _RenamedFrom.Valid {
		t.SetRenamedFrom(_RenamedFrom.String)
	}
	if _RenamedTo.Valid {
		t.SetRenamedTo(_RenamedTo.String)
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

// CountCommitFiles will find the count of CommitFile records in the database
func CountCommitFiles(ctx context.Context, db DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(CommitFileTableName),
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

// CountCommitFilesTx will find the count of CommitFile records in the database using the provided transaction
func CountCommitFilesTx(ctx context.Context, tx Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(CommitFileTableName),
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

// DBCount will find the count of CommitFile records in the database
func (t *CommitFile) DBCount(ctx context.Context, db DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(CommitFileTableName),
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

// DBCountTx will find the count of CommitFile records in the database using the provided transaction
func (t *CommitFile) DBCountTx(ctx context.Context, tx Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(CommitFileTableName),
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

// DBExists will return true if the CommitFile record exists in the database
func (t *CommitFile) DBExists(ctx context.Context, db DB) (bool, error) {
	q := "SELECT * FROM `commit_file` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := db.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// DBExistsTx will return true if the CommitFile record exists in the database using the provided transaction
func (t *CommitFile) DBExistsTx(ctx context.Context, tx Tx) (bool, error) {
	q := "SELECT * FROM `commit_file` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := tx.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// PrimaryKeyColumn returns the column name for the primary key
func (t *CommitFile) PrimaryKeyColumn() string {
	return CommitFileColumnID
}

// PrimaryKeyColumnType returns the primary key column Go type as a string
func (t *CommitFile) PrimaryKeyColumnType() string {
	return "string"
}

// PrimaryKey returns the primary key column value
func (t *CommitFile) PrimaryKey() interface{} {
	return t.ID
}
