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

var _ Model = (*CommitActivity)(nil)
var _ CSVWriter = (*CommitActivity)(nil)
var _ JSONWriter = (*CommitActivity)(nil)

// CommitActivityTableName is the name of the table in SQL
const CommitActivityTableName = "commit_activity"

var CommitActivityColumns = []string{
	"id",
	"date",
	"sha",
	"user_id",
	"repo_id",
	"filename",
	"language",
	"ordinal",
	"loc",
	"sloc",
	"blanks",
	"comments",
}

// CommitActivity table
type CommitActivity struct {
	Blanks   int32  `json:"blanks"`
	Comments int32  `json:"comments"`
	Date     int64  `json:"date"`
	Filename string `json:"filename"`
	ID       string `json:"id"`
	Language string `json:"language"`
	Loc      int32  `json:"loc"`
	Ordinal  int64  `json:"ordinal"`
	RepoID   string `json:"repo_id"`
	Sha      string `json:"sha"`
	Sloc     int32  `json:"sloc"`
	UserID   string `json:"user_id"`
}

// TableName returns the SQL table name for CommitActivity and satifies the Model interface
func (t *CommitActivity) TableName() string {
	return CommitActivityTableName
}

// ToCSV will serialize the CommitActivity instance to a CSV compatible array of strings
func (t *CommitActivity) ToCSV() []string {
	return []string{
		t.ID,
		toCSVString(t.Date),
		t.Sha,
		t.UserID,
		t.RepoID,
		t.Filename,
		t.Language,
		toCSVString(t.Ordinal),
		toCSVString(t.Loc),
		toCSVString(t.Sloc),
		toCSVString(t.Blanks),
		toCSVString(t.Comments),
	}
}

// WriteCSV will serialize the CommitActivity instance to the writer as CSV and satisfies the CSVWriter interface
func (t *CommitActivity) WriteCSV(w *csv.Writer) error {
	return w.Write(t.ToCSV())
}

// WriteJSON will serialize the CommitActivity instance to the writer as JSON and satisfies the JSONWriter interface
func (t *CommitActivity) WriteJSON(w io.Writer, indent ...bool) error {
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

// NewCommitActivityReader creates a JSON reader which can read in CommitActivity objects serialized as JSON either as an array, single object or json new lines
// and writes each CommitActivity to the channel provided
func NewCommitActivityReader(r io.Reader, ch chan<- CommitActivity) error {
	return orm.Deserialize(r, func(buf json.RawMessage) error {
		dec := json.NewDecoder(bytes.NewBuffer(buf))
		e := CommitActivity{}
		if err := dec.Decode(&e); err != nil {
			return err
		}
		ch <- e
		return nil
	})
}

// NewCSVCommitActivityReaderDir will read the reader as CSV and emit each record to the channel provided
func NewCSVCommitActivityReader(r io.Reader, ch chan<- CommitActivity) error {
	cr := csv.NewReader(r)
	for {
		record, err := cr.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		ch <- CommitActivity{
			ID:       record[0],
			Date:     fromCSVInt64(record[1]),
			Sha:      record[2],
			UserID:   record[3],
			RepoID:   record[4],
			Filename: record[5],
			Language: record[6],
			Ordinal:  fromCSVInt64(record[7]),
			Loc:      fromCSVInt32(record[8]),
			Sloc:     fromCSVInt32(record[9]),
			Blanks:   fromCSVInt32(record[10]),
			Comments: fromCSVInt32(record[11]),
		}
	}
	return nil
}

// NewCSVCommitActivityReaderFile will read the file as a CSV and emit each record to the channel provided
func NewCSVCommitActivityReaderFile(fp string, ch chan<- CommitActivity) error {
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
	return NewCSVCommitActivityReader(fc, ch)
}

// NewCSVCommitActivityReaderDir will read the commit_activity.csv.gz file as a CSV and emit each record to the channel provided
func NewCSVCommitActivityReaderDir(dir string, ch chan<- CommitActivity) error {
	return NewCSVCommitActivityReaderFile(filepath.Join(dir, "commit_activity.csv.gz"), ch)
}

// CommitActivityCSVDeduper is a function callback which takes the existing value (a) and the new value (b)
// and the return value should be the one to use (or a new one, if applicable). return nil
// to skip processing of this record
type CommitActivityCSVDeduper func(a CommitActivity, b CommitActivity) *CommitActivity

// CommitActivityCSVDedupeDisabled is set on whether the CSV writer should de-dupe values by key
var CommitActivityCSVDedupeDisabled bool

// NewCommitActivityCSVWriterSize creates a batch writer that will write each CommitActivity into a CSV file
func NewCommitActivityCSVWriterSize(w io.Writer, size int, dedupers ...CommitActivityCSVDeduper) (chan CommitActivity, chan bool, error) {
	cw := csv.NewWriter(w)
	ch := make(chan CommitActivity, size)
	done := make(chan bool)
	go func() {
		defer func() { done <- true }()
		dodedupe := !CommitActivityCSVDedupeDisabled
		var kv map[string]*CommitActivity
		var deduper CommitActivityCSVDeduper
		if dedupers != nil && len(dedupers) > 0 {
			deduper = dedupers[0]
			dodedupe = true
		}
		if dodedupe {
			kv = make(map[string]*CommitActivity)
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

// CommitActivityCSVDefaultSize is the default channel buffer size if not provided
var CommitActivityCSVDefaultSize = 100

// NewCommitActivityCSVWriter creates a batch writer that will write each CommitActivity into a CSV file
func NewCommitActivityCSVWriter(w io.Writer, dedupers ...CommitActivityCSVDeduper) (chan CommitActivity, chan bool, error) {
	return NewCommitActivityCSVWriterSize(w, CommitActivityCSVDefaultSize, dedupers...)
}

// NewCommitActivityCSVWriterDir creates a batch writer that will write each CommitActivity into a CSV file named commit_activity.csv.gz in dir
func NewCommitActivityCSVWriterDir(dir string, dedupers ...CommitActivityCSVDeduper) (chan CommitActivity, chan bool, error) {
	return NewCommitActivityCSVWriterFile(filepath.Join(dir, "commit_activity.csv.gz"), dedupers...)
}

// NewCommitActivityCSVWriterFile creates a batch writer that will write each CommitActivity into a CSV file
func NewCommitActivityCSVWriterFile(fn string, dedupers ...CommitActivityCSVDeduper) (chan CommitActivity, chan bool, error) {
	f, err := os.Create(fn)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening CSV file %s. %v", fn, err)
	}
	var fc io.WriteCloser = f
	if filepath.Ext(fn) == ".gz" {
		w, _ := gzip.NewWriterLevel(f, gzip.BestCompression)
		fc = w
	}
	ch, done, err := NewCommitActivityCSVWriter(fc, dedupers...)
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

type CommitActivityDBAction func(ctx context.Context, db DB, record CommitActivity) error

// NewCommitActivityDBWriterSize creates a DB writer that will write each issue into the DB
func NewCommitActivityDBWriterSize(ctx context.Context, db DB, errors chan<- error, size int, actions ...CommitActivityDBAction) (chan CommitActivity, chan bool, error) {
	ch := make(chan CommitActivity, size)
	done := make(chan bool)
	var action CommitActivityDBAction
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

// NewCommitActivityDBWriter creates a DB writer that will write each issue into the DB
func NewCommitActivityDBWriter(ctx context.Context, db DB, errors chan<- error, actions ...CommitActivityDBAction) (chan CommitActivity, chan bool, error) {
	return NewCommitActivityDBWriterSize(ctx, db, errors, 100, actions...)
}

// CommitActivityColumnID is the ID SQL column name for the CommitActivity table
const CommitActivityColumnID = "id"

// CommitActivityEscapedColumnID is the escaped ID SQL column name for the CommitActivity table
const CommitActivityEscapedColumnID = "`id`"

// CommitActivityColumnDate is the Date SQL column name for the CommitActivity table
const CommitActivityColumnDate = "date"

// CommitActivityEscapedColumnDate is the escaped Date SQL column name for the CommitActivity table
const CommitActivityEscapedColumnDate = "`date`"

// CommitActivityColumnSha is the Sha SQL column name for the CommitActivity table
const CommitActivityColumnSha = "sha"

// CommitActivityEscapedColumnSha is the escaped Sha SQL column name for the CommitActivity table
const CommitActivityEscapedColumnSha = "`sha`"

// CommitActivityColumnUserID is the UserID SQL column name for the CommitActivity table
const CommitActivityColumnUserID = "user_id"

// CommitActivityEscapedColumnUserID is the escaped UserID SQL column name for the CommitActivity table
const CommitActivityEscapedColumnUserID = "`user_id`"

// CommitActivityColumnRepoID is the RepoID SQL column name for the CommitActivity table
const CommitActivityColumnRepoID = "repo_id"

// CommitActivityEscapedColumnRepoID is the escaped RepoID SQL column name for the CommitActivity table
const CommitActivityEscapedColumnRepoID = "`repo_id`"

// CommitActivityColumnFilename is the Filename SQL column name for the CommitActivity table
const CommitActivityColumnFilename = "filename"

// CommitActivityEscapedColumnFilename is the escaped Filename SQL column name for the CommitActivity table
const CommitActivityEscapedColumnFilename = "`filename`"

// CommitActivityColumnLanguage is the Language SQL column name for the CommitActivity table
const CommitActivityColumnLanguage = "language"

// CommitActivityEscapedColumnLanguage is the escaped Language SQL column name for the CommitActivity table
const CommitActivityEscapedColumnLanguage = "`language`"

// CommitActivityColumnOrdinal is the Ordinal SQL column name for the CommitActivity table
const CommitActivityColumnOrdinal = "ordinal"

// CommitActivityEscapedColumnOrdinal is the escaped Ordinal SQL column name for the CommitActivity table
const CommitActivityEscapedColumnOrdinal = "`ordinal`"

// CommitActivityColumnLoc is the Loc SQL column name for the CommitActivity table
const CommitActivityColumnLoc = "loc"

// CommitActivityEscapedColumnLoc is the escaped Loc SQL column name for the CommitActivity table
const CommitActivityEscapedColumnLoc = "`loc`"

// CommitActivityColumnSloc is the Sloc SQL column name for the CommitActivity table
const CommitActivityColumnSloc = "sloc"

// CommitActivityEscapedColumnSloc is the escaped Sloc SQL column name for the CommitActivity table
const CommitActivityEscapedColumnSloc = "`sloc`"

// CommitActivityColumnBlanks is the Blanks SQL column name for the CommitActivity table
const CommitActivityColumnBlanks = "blanks"

// CommitActivityEscapedColumnBlanks is the escaped Blanks SQL column name for the CommitActivity table
const CommitActivityEscapedColumnBlanks = "`blanks`"

// CommitActivityColumnComments is the Comments SQL column name for the CommitActivity table
const CommitActivityColumnComments = "comments"

// CommitActivityEscapedColumnComments is the escaped Comments SQL column name for the CommitActivity table
const CommitActivityEscapedColumnComments = "`comments`"

// GetID will return the CommitActivity ID value
func (t *CommitActivity) GetID() string {
	return t.ID
}

// SetID will set the CommitActivity ID value
func (t *CommitActivity) SetID(v string) {
	t.ID = v
}

// FindCommitActivityByID will find a CommitActivity by ID
func FindCommitActivityByID(ctx context.Context, db DB, value string) (*CommitActivity, error) {
	q := "SELECT * FROM `commit_activity` WHERE `id` = ?"
	var _ID sql.NullString
	var _Date sql.NullInt64
	var _Sha sql.NullString
	var _UserID sql.NullString
	var _RepoID sql.NullString
	var _Filename sql.NullString
	var _Language sql.NullString
	var _Ordinal sql.NullInt64
	var _Loc sql.NullInt64
	var _Sloc sql.NullInt64
	var _Blanks sql.NullInt64
	var _Comments sql.NullInt64
	err := db.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Date,
		&_Sha,
		&_UserID,
		&_RepoID,
		&_Filename,
		&_Language,
		&_Ordinal,
		&_Loc,
		&_Sloc,
		&_Blanks,
		&_Comments,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &CommitActivity{}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _Date.Valid {
		t.SetDate(_Date.Int64)
	}
	if _Sha.Valid {
		t.SetSha(_Sha.String)
	}
	if _UserID.Valid {
		t.SetUserID(_UserID.String)
	}
	if _RepoID.Valid {
		t.SetRepoID(_RepoID.String)
	}
	if _Filename.Valid {
		t.SetFilename(_Filename.String)
	}
	if _Language.Valid {
		t.SetLanguage(_Language.String)
	}
	if _Ordinal.Valid {
		t.SetOrdinal(_Ordinal.Int64)
	}
	if _Loc.Valid {
		t.SetLoc(int32(_Loc.Int64))
	}
	if _Sloc.Valid {
		t.SetSloc(int32(_Sloc.Int64))
	}
	if _Blanks.Valid {
		t.SetBlanks(int32(_Blanks.Int64))
	}
	if _Comments.Valid {
		t.SetComments(int32(_Comments.Int64))
	}
	return t, nil
}

// FindCommitActivityByIDTx will find a CommitActivity by ID using the provided transaction
func FindCommitActivityByIDTx(ctx context.Context, tx Tx, value string) (*CommitActivity, error) {
	q := "SELECT * FROM `commit_activity` WHERE `id` = ?"
	var _ID sql.NullString
	var _Date sql.NullInt64
	var _Sha sql.NullString
	var _UserID sql.NullString
	var _RepoID sql.NullString
	var _Filename sql.NullString
	var _Language sql.NullString
	var _Ordinal sql.NullInt64
	var _Loc sql.NullInt64
	var _Sloc sql.NullInt64
	var _Blanks sql.NullInt64
	var _Comments sql.NullInt64
	err := tx.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Date,
		&_Sha,
		&_UserID,
		&_RepoID,
		&_Filename,
		&_Language,
		&_Ordinal,
		&_Loc,
		&_Sloc,
		&_Blanks,
		&_Comments,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &CommitActivity{}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _Date.Valid {
		t.SetDate(_Date.Int64)
	}
	if _Sha.Valid {
		t.SetSha(_Sha.String)
	}
	if _UserID.Valid {
		t.SetUserID(_UserID.String)
	}
	if _RepoID.Valid {
		t.SetRepoID(_RepoID.String)
	}
	if _Filename.Valid {
		t.SetFilename(_Filename.String)
	}
	if _Language.Valid {
		t.SetLanguage(_Language.String)
	}
	if _Ordinal.Valid {
		t.SetOrdinal(_Ordinal.Int64)
	}
	if _Loc.Valid {
		t.SetLoc(int32(_Loc.Int64))
	}
	if _Sloc.Valid {
		t.SetSloc(int32(_Sloc.Int64))
	}
	if _Blanks.Valid {
		t.SetBlanks(int32(_Blanks.Int64))
	}
	if _Comments.Valid {
		t.SetComments(int32(_Comments.Int64))
	}
	return t, nil
}

// GetDate will return the CommitActivity Date value
func (t *CommitActivity) GetDate() int64 {
	return t.Date
}

// SetDate will set the CommitActivity Date value
func (t *CommitActivity) SetDate(v int64) {
	t.Date = v
}

// GetSha will return the CommitActivity Sha value
func (t *CommitActivity) GetSha() string {
	return t.Sha
}

// SetSha will set the CommitActivity Sha value
func (t *CommitActivity) SetSha(v string) {
	t.Sha = v
}

// FindCommitActivitiesBySha will find all CommitActivitys by the Sha value
func FindCommitActivitiesBySha(ctx context.Context, db DB, value string) ([]*CommitActivity, error) {
	q := "SELECT * FROM `commit_activity` WHERE `sha` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*CommitActivity, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Date sql.NullInt64
		var _Sha sql.NullString
		var _UserID sql.NullString
		var _RepoID sql.NullString
		var _Filename sql.NullString
		var _Language sql.NullString
		var _Ordinal sql.NullInt64
		var _Loc sql.NullInt64
		var _Sloc sql.NullInt64
		var _Blanks sql.NullInt64
		var _Comments sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Date,
			&_Sha,
			&_UserID,
			&_RepoID,
			&_Filename,
			&_Language,
			&_Ordinal,
			&_Loc,
			&_Sloc,
			&_Blanks,
			&_Comments,
		)
		if err != nil {
			return nil, err
		}
		t := &CommitActivity{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Date.Valid {
			t.SetDate(_Date.Int64)
		}
		if _Sha.Valid {
			t.SetSha(_Sha.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _RepoID.Valid {
			t.SetRepoID(_RepoID.String)
		}
		if _Filename.Valid {
			t.SetFilename(_Filename.String)
		}
		if _Language.Valid {
			t.SetLanguage(_Language.String)
		}
		if _Ordinal.Valid {
			t.SetOrdinal(_Ordinal.Int64)
		}
		if _Loc.Valid {
			t.SetLoc(int32(_Loc.Int64))
		}
		if _Sloc.Valid {
			t.SetSloc(int32(_Sloc.Int64))
		}
		if _Blanks.Valid {
			t.SetBlanks(int32(_Blanks.Int64))
		}
		if _Comments.Valid {
			t.SetComments(int32(_Comments.Int64))
		}
		results = append(results, t)
	}
	return results, nil
}

// FindCommitActivitiesByShaTx will find all CommitActivitys by the Sha value using the provided transaction
func FindCommitActivitiesByShaTx(ctx context.Context, tx Tx, value string) ([]*CommitActivity, error) {
	q := "SELECT * FROM `commit_activity` WHERE `sha` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*CommitActivity, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Date sql.NullInt64
		var _Sha sql.NullString
		var _UserID sql.NullString
		var _RepoID sql.NullString
		var _Filename sql.NullString
		var _Language sql.NullString
		var _Ordinal sql.NullInt64
		var _Loc sql.NullInt64
		var _Sloc sql.NullInt64
		var _Blanks sql.NullInt64
		var _Comments sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Date,
			&_Sha,
			&_UserID,
			&_RepoID,
			&_Filename,
			&_Language,
			&_Ordinal,
			&_Loc,
			&_Sloc,
			&_Blanks,
			&_Comments,
		)
		if err != nil {
			return nil, err
		}
		t := &CommitActivity{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Date.Valid {
			t.SetDate(_Date.Int64)
		}
		if _Sha.Valid {
			t.SetSha(_Sha.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _RepoID.Valid {
			t.SetRepoID(_RepoID.String)
		}
		if _Filename.Valid {
			t.SetFilename(_Filename.String)
		}
		if _Language.Valid {
			t.SetLanguage(_Language.String)
		}
		if _Ordinal.Valid {
			t.SetOrdinal(_Ordinal.Int64)
		}
		if _Loc.Valid {
			t.SetLoc(int32(_Loc.Int64))
		}
		if _Sloc.Valid {
			t.SetSloc(int32(_Sloc.Int64))
		}
		if _Blanks.Valid {
			t.SetBlanks(int32(_Blanks.Int64))
		}
		if _Comments.Valid {
			t.SetComments(int32(_Comments.Int64))
		}
		results = append(results, t)
	}
	return results, nil
}

// GetUserID will return the CommitActivity UserID value
func (t *CommitActivity) GetUserID() string {
	return t.UserID
}

// SetUserID will set the CommitActivity UserID value
func (t *CommitActivity) SetUserID(v string) {
	t.UserID = v
}

// FindCommitActivitiesByUserID will find all CommitActivitys by the UserID value
func FindCommitActivitiesByUserID(ctx context.Context, db DB, value string) ([]*CommitActivity, error) {
	q := "SELECT * FROM `commit_activity` WHERE `user_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*CommitActivity, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Date sql.NullInt64
		var _Sha sql.NullString
		var _UserID sql.NullString
		var _RepoID sql.NullString
		var _Filename sql.NullString
		var _Language sql.NullString
		var _Ordinal sql.NullInt64
		var _Loc sql.NullInt64
		var _Sloc sql.NullInt64
		var _Blanks sql.NullInt64
		var _Comments sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Date,
			&_Sha,
			&_UserID,
			&_RepoID,
			&_Filename,
			&_Language,
			&_Ordinal,
			&_Loc,
			&_Sloc,
			&_Blanks,
			&_Comments,
		)
		if err != nil {
			return nil, err
		}
		t := &CommitActivity{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Date.Valid {
			t.SetDate(_Date.Int64)
		}
		if _Sha.Valid {
			t.SetSha(_Sha.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _RepoID.Valid {
			t.SetRepoID(_RepoID.String)
		}
		if _Filename.Valid {
			t.SetFilename(_Filename.String)
		}
		if _Language.Valid {
			t.SetLanguage(_Language.String)
		}
		if _Ordinal.Valid {
			t.SetOrdinal(_Ordinal.Int64)
		}
		if _Loc.Valid {
			t.SetLoc(int32(_Loc.Int64))
		}
		if _Sloc.Valid {
			t.SetSloc(int32(_Sloc.Int64))
		}
		if _Blanks.Valid {
			t.SetBlanks(int32(_Blanks.Int64))
		}
		if _Comments.Valid {
			t.SetComments(int32(_Comments.Int64))
		}
		results = append(results, t)
	}
	return results, nil
}

// FindCommitActivitiesByUserIDTx will find all CommitActivitys by the UserID value using the provided transaction
func FindCommitActivitiesByUserIDTx(ctx context.Context, tx Tx, value string) ([]*CommitActivity, error) {
	q := "SELECT * FROM `commit_activity` WHERE `user_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*CommitActivity, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Date sql.NullInt64
		var _Sha sql.NullString
		var _UserID sql.NullString
		var _RepoID sql.NullString
		var _Filename sql.NullString
		var _Language sql.NullString
		var _Ordinal sql.NullInt64
		var _Loc sql.NullInt64
		var _Sloc sql.NullInt64
		var _Blanks sql.NullInt64
		var _Comments sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Date,
			&_Sha,
			&_UserID,
			&_RepoID,
			&_Filename,
			&_Language,
			&_Ordinal,
			&_Loc,
			&_Sloc,
			&_Blanks,
			&_Comments,
		)
		if err != nil {
			return nil, err
		}
		t := &CommitActivity{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Date.Valid {
			t.SetDate(_Date.Int64)
		}
		if _Sha.Valid {
			t.SetSha(_Sha.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _RepoID.Valid {
			t.SetRepoID(_RepoID.String)
		}
		if _Filename.Valid {
			t.SetFilename(_Filename.String)
		}
		if _Language.Valid {
			t.SetLanguage(_Language.String)
		}
		if _Ordinal.Valid {
			t.SetOrdinal(_Ordinal.Int64)
		}
		if _Loc.Valid {
			t.SetLoc(int32(_Loc.Int64))
		}
		if _Sloc.Valid {
			t.SetSloc(int32(_Sloc.Int64))
		}
		if _Blanks.Valid {
			t.SetBlanks(int32(_Blanks.Int64))
		}
		if _Comments.Valid {
			t.SetComments(int32(_Comments.Int64))
		}
		results = append(results, t)
	}
	return results, nil
}

// GetRepoID will return the CommitActivity RepoID value
func (t *CommitActivity) GetRepoID() string {
	return t.RepoID
}

// SetRepoID will set the CommitActivity RepoID value
func (t *CommitActivity) SetRepoID(v string) {
	t.RepoID = v
}

// FindCommitActivitiesByRepoID will find all CommitActivitys by the RepoID value
func FindCommitActivitiesByRepoID(ctx context.Context, db DB, value string) ([]*CommitActivity, error) {
	q := "SELECT * FROM `commit_activity` WHERE `repo_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*CommitActivity, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Date sql.NullInt64
		var _Sha sql.NullString
		var _UserID sql.NullString
		var _RepoID sql.NullString
		var _Filename sql.NullString
		var _Language sql.NullString
		var _Ordinal sql.NullInt64
		var _Loc sql.NullInt64
		var _Sloc sql.NullInt64
		var _Blanks sql.NullInt64
		var _Comments sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Date,
			&_Sha,
			&_UserID,
			&_RepoID,
			&_Filename,
			&_Language,
			&_Ordinal,
			&_Loc,
			&_Sloc,
			&_Blanks,
			&_Comments,
		)
		if err != nil {
			return nil, err
		}
		t := &CommitActivity{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Date.Valid {
			t.SetDate(_Date.Int64)
		}
		if _Sha.Valid {
			t.SetSha(_Sha.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _RepoID.Valid {
			t.SetRepoID(_RepoID.String)
		}
		if _Filename.Valid {
			t.SetFilename(_Filename.String)
		}
		if _Language.Valid {
			t.SetLanguage(_Language.String)
		}
		if _Ordinal.Valid {
			t.SetOrdinal(_Ordinal.Int64)
		}
		if _Loc.Valid {
			t.SetLoc(int32(_Loc.Int64))
		}
		if _Sloc.Valid {
			t.SetSloc(int32(_Sloc.Int64))
		}
		if _Blanks.Valid {
			t.SetBlanks(int32(_Blanks.Int64))
		}
		if _Comments.Valid {
			t.SetComments(int32(_Comments.Int64))
		}
		results = append(results, t)
	}
	return results, nil
}

// FindCommitActivitiesByRepoIDTx will find all CommitActivitys by the RepoID value using the provided transaction
func FindCommitActivitiesByRepoIDTx(ctx context.Context, tx Tx, value string) ([]*CommitActivity, error) {
	q := "SELECT * FROM `commit_activity` WHERE `repo_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*CommitActivity, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Date sql.NullInt64
		var _Sha sql.NullString
		var _UserID sql.NullString
		var _RepoID sql.NullString
		var _Filename sql.NullString
		var _Language sql.NullString
		var _Ordinal sql.NullInt64
		var _Loc sql.NullInt64
		var _Sloc sql.NullInt64
		var _Blanks sql.NullInt64
		var _Comments sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Date,
			&_Sha,
			&_UserID,
			&_RepoID,
			&_Filename,
			&_Language,
			&_Ordinal,
			&_Loc,
			&_Sloc,
			&_Blanks,
			&_Comments,
		)
		if err != nil {
			return nil, err
		}
		t := &CommitActivity{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Date.Valid {
			t.SetDate(_Date.Int64)
		}
		if _Sha.Valid {
			t.SetSha(_Sha.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _RepoID.Valid {
			t.SetRepoID(_RepoID.String)
		}
		if _Filename.Valid {
			t.SetFilename(_Filename.String)
		}
		if _Language.Valid {
			t.SetLanguage(_Language.String)
		}
		if _Ordinal.Valid {
			t.SetOrdinal(_Ordinal.Int64)
		}
		if _Loc.Valid {
			t.SetLoc(int32(_Loc.Int64))
		}
		if _Sloc.Valid {
			t.SetSloc(int32(_Sloc.Int64))
		}
		if _Blanks.Valid {
			t.SetBlanks(int32(_Blanks.Int64))
		}
		if _Comments.Valid {
			t.SetComments(int32(_Comments.Int64))
		}
		results = append(results, t)
	}
	return results, nil
}

// GetFilename will return the CommitActivity Filename value
func (t *CommitActivity) GetFilename() string {
	return t.Filename
}

// SetFilename will set the CommitActivity Filename value
func (t *CommitActivity) SetFilename(v string) {
	t.Filename = v
}

// FindCommitActivitiesByFilename will find all CommitActivitys by the Filename value
func FindCommitActivitiesByFilename(ctx context.Context, db DB, value string) ([]*CommitActivity, error) {
	q := "SELECT * FROM `commit_activity` WHERE `filename` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*CommitActivity, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Date sql.NullInt64
		var _Sha sql.NullString
		var _UserID sql.NullString
		var _RepoID sql.NullString
		var _Filename sql.NullString
		var _Language sql.NullString
		var _Ordinal sql.NullInt64
		var _Loc sql.NullInt64
		var _Sloc sql.NullInt64
		var _Blanks sql.NullInt64
		var _Comments sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Date,
			&_Sha,
			&_UserID,
			&_RepoID,
			&_Filename,
			&_Language,
			&_Ordinal,
			&_Loc,
			&_Sloc,
			&_Blanks,
			&_Comments,
		)
		if err != nil {
			return nil, err
		}
		t := &CommitActivity{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Date.Valid {
			t.SetDate(_Date.Int64)
		}
		if _Sha.Valid {
			t.SetSha(_Sha.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _RepoID.Valid {
			t.SetRepoID(_RepoID.String)
		}
		if _Filename.Valid {
			t.SetFilename(_Filename.String)
		}
		if _Language.Valid {
			t.SetLanguage(_Language.String)
		}
		if _Ordinal.Valid {
			t.SetOrdinal(_Ordinal.Int64)
		}
		if _Loc.Valid {
			t.SetLoc(int32(_Loc.Int64))
		}
		if _Sloc.Valid {
			t.SetSloc(int32(_Sloc.Int64))
		}
		if _Blanks.Valid {
			t.SetBlanks(int32(_Blanks.Int64))
		}
		if _Comments.Valid {
			t.SetComments(int32(_Comments.Int64))
		}
		results = append(results, t)
	}
	return results, nil
}

// FindCommitActivitiesByFilenameTx will find all CommitActivitys by the Filename value using the provided transaction
func FindCommitActivitiesByFilenameTx(ctx context.Context, tx Tx, value string) ([]*CommitActivity, error) {
	q := "SELECT * FROM `commit_activity` WHERE `filename` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*CommitActivity, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Date sql.NullInt64
		var _Sha sql.NullString
		var _UserID sql.NullString
		var _RepoID sql.NullString
		var _Filename sql.NullString
		var _Language sql.NullString
		var _Ordinal sql.NullInt64
		var _Loc sql.NullInt64
		var _Sloc sql.NullInt64
		var _Blanks sql.NullInt64
		var _Comments sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Date,
			&_Sha,
			&_UserID,
			&_RepoID,
			&_Filename,
			&_Language,
			&_Ordinal,
			&_Loc,
			&_Sloc,
			&_Blanks,
			&_Comments,
		)
		if err != nil {
			return nil, err
		}
		t := &CommitActivity{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Date.Valid {
			t.SetDate(_Date.Int64)
		}
		if _Sha.Valid {
			t.SetSha(_Sha.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _RepoID.Valid {
			t.SetRepoID(_RepoID.String)
		}
		if _Filename.Valid {
			t.SetFilename(_Filename.String)
		}
		if _Language.Valid {
			t.SetLanguage(_Language.String)
		}
		if _Ordinal.Valid {
			t.SetOrdinal(_Ordinal.Int64)
		}
		if _Loc.Valid {
			t.SetLoc(int32(_Loc.Int64))
		}
		if _Sloc.Valid {
			t.SetSloc(int32(_Sloc.Int64))
		}
		if _Blanks.Valid {
			t.SetBlanks(int32(_Blanks.Int64))
		}
		if _Comments.Valid {
			t.SetComments(int32(_Comments.Int64))
		}
		results = append(results, t)
	}
	return results, nil
}

// GetLanguage will return the CommitActivity Language value
func (t *CommitActivity) GetLanguage() string {
	return t.Language
}

// SetLanguage will set the CommitActivity Language value
func (t *CommitActivity) SetLanguage(v string) {
	t.Language = v
}

// GetOrdinal will return the CommitActivity Ordinal value
func (t *CommitActivity) GetOrdinal() int64 {
	return t.Ordinal
}

// SetOrdinal will set the CommitActivity Ordinal value
func (t *CommitActivity) SetOrdinal(v int64) {
	t.Ordinal = v
}

// GetLoc will return the CommitActivity Loc value
func (t *CommitActivity) GetLoc() int32 {
	return t.Loc
}

// SetLoc will set the CommitActivity Loc value
func (t *CommitActivity) SetLoc(v int32) {
	t.Loc = v
}

// GetSloc will return the CommitActivity Sloc value
func (t *CommitActivity) GetSloc() int32 {
	return t.Sloc
}

// SetSloc will set the CommitActivity Sloc value
func (t *CommitActivity) SetSloc(v int32) {
	t.Sloc = v
}

// GetBlanks will return the CommitActivity Blanks value
func (t *CommitActivity) GetBlanks() int32 {
	return t.Blanks
}

// SetBlanks will set the CommitActivity Blanks value
func (t *CommitActivity) SetBlanks(v int32) {
	t.Blanks = v
}

// GetComments will return the CommitActivity Comments value
func (t *CommitActivity) GetComments() int32 {
	return t.Comments
}

// SetComments will set the CommitActivity Comments value
func (t *CommitActivity) SetComments(v int32) {
	t.Comments = v
}

func (t *CommitActivity) toTimestamp(value time.Time) *timestamp.Timestamp {
	ts, _ := ptypes.TimestampProto(value)
	return ts
}

// DBCreateCommitActivityTable will create the CommitActivity table
func DBCreateCommitActivityTable(ctx context.Context, db DB) error {
	q := "CREATE TABLE `commit_activity` (`id`VARCHAR(64) NOT NULL PRIMARY KEY,`date` BIGINT UNSIGNED NOT NULL,`sha` VARCHAR(64) NOT NULL,`user_id` VARCHAR(64) NOT NULL,`repo_id` VARCHAR(64) NOT NULL,`filename`VARCHAR(700) NOT NULL,`language`VARCHAR(100) NOT NULL DEFAULT \"unknown\",`ordinal` BIGINT UNSIGNED NOT NULL,`loc` INT NOT NULL DEFAULT 0,`sloc` INT NOT NULL DEFAULT 0,`blanks` INT NOT NULL DEFAULT 0,`comments`INT NOT NULL DEFAULT 0,INDEX commit_activity_sha_index (`sha`),INDEX commit_activity_user_id_index (`user_id`),INDEX commit_activity_repo_id_index (`repo_id`),INDEX commit_activity_filename_index (`filename`),INDEX commit_activity_filename_repo_id_date_ordinal_index (`filename`,`repo_id`,`date`,`ordinal`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBCreateCommitActivityTableTx will create the CommitActivity table using the provided transction
func DBCreateCommitActivityTableTx(ctx context.Context, tx Tx) error {
	q := "CREATE TABLE `commit_activity` (`id`VARCHAR(64) NOT NULL PRIMARY KEY,`date` BIGINT UNSIGNED NOT NULL,`sha` VARCHAR(64) NOT NULL,`user_id` VARCHAR(64) NOT NULL,`repo_id` VARCHAR(64) NOT NULL,`filename`VARCHAR(700) NOT NULL,`language`VARCHAR(100) NOT NULL DEFAULT \"unknown\",`ordinal` BIGINT UNSIGNED NOT NULL,`loc` INT NOT NULL DEFAULT 0,`sloc` INT NOT NULL DEFAULT 0,`blanks` INT NOT NULL DEFAULT 0,`comments`INT NOT NULL DEFAULT 0,INDEX commit_activity_sha_index (`sha`),INDEX commit_activity_user_id_index (`user_id`),INDEX commit_activity_repo_id_index (`repo_id`),INDEX commit_activity_filename_index (`filename`),INDEX commit_activity_filename_repo_id_date_ordinal_index (`filename`,`repo_id`,`date`,`ordinal`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// DBDropCommitActivityTable will drop the CommitActivity table
func DBDropCommitActivityTable(ctx context.Context, db DB) error {
	q := "DROP TABLE IF EXISTS `commit_activity`"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBDropCommitActivityTableTx will drop the CommitActivity table using the provided transaction
func DBDropCommitActivityTableTx(ctx context.Context, tx Tx) error {
	q := "DROP TABLE IF EXISTS `commit_activity`"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// DBCreate will create a new CommitActivity record in the database
func (t *CommitActivity) DBCreate(ctx context.Context, db DB) (sql.Result, error) {
	q := "INSERT INTO `commit_activity` (`commit_activity`.`id`,`commit_activity`.`date`,`commit_activity`.`sha`,`commit_activity`.`user_id`,`commit_activity`.`repo_id`,`commit_activity`.`filename`,`commit_activity`.`language`,`commit_activity`.`ordinal`,`commit_activity`.`loc`,`commit_activity`.`sloc`,`commit_activity`.`blanks`,`commit_activity`.`comments`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?)"
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLInt64(t.Date),
		orm.ToSQLString(t.Sha),
		orm.ToSQLString(t.UserID),
		orm.ToSQLString(t.RepoID),
		orm.ToSQLString(t.Filename),
		orm.ToSQLString(t.Language),
		orm.ToSQLInt64(t.Ordinal),
		orm.ToSQLInt64(t.Loc),
		orm.ToSQLInt64(t.Sloc),
		orm.ToSQLInt64(t.Blanks),
		orm.ToSQLInt64(t.Comments),
	)
}

// DBCreateTx will create a new CommitActivity record in the database using the provided transaction
func (t *CommitActivity) DBCreateTx(ctx context.Context, tx Tx) (sql.Result, error) {
	q := "INSERT INTO `commit_activity` (`commit_activity`.`id`,`commit_activity`.`date`,`commit_activity`.`sha`,`commit_activity`.`user_id`,`commit_activity`.`repo_id`,`commit_activity`.`filename`,`commit_activity`.`language`,`commit_activity`.`ordinal`,`commit_activity`.`loc`,`commit_activity`.`sloc`,`commit_activity`.`blanks`,`commit_activity`.`comments`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?)"
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLInt64(t.Date),
		orm.ToSQLString(t.Sha),
		orm.ToSQLString(t.UserID),
		orm.ToSQLString(t.RepoID),
		orm.ToSQLString(t.Filename),
		orm.ToSQLString(t.Language),
		orm.ToSQLInt64(t.Ordinal),
		orm.ToSQLInt64(t.Loc),
		orm.ToSQLInt64(t.Sloc),
		orm.ToSQLInt64(t.Blanks),
		orm.ToSQLInt64(t.Comments),
	)
}

// DBCreateIgnoreDuplicate will upsert the CommitActivity record in the database
func (t *CommitActivity) DBCreateIgnoreDuplicate(ctx context.Context, db DB) (sql.Result, error) {
	q := "INSERT INTO `commit_activity` (`commit_activity`.`id`,`commit_activity`.`date`,`commit_activity`.`sha`,`commit_activity`.`user_id`,`commit_activity`.`repo_id`,`commit_activity`.`filename`,`commit_activity`.`language`,`commit_activity`.`ordinal`,`commit_activity`.`loc`,`commit_activity`.`sloc`,`commit_activity`.`blanks`,`commit_activity`.`comments`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLInt64(t.Date),
		orm.ToSQLString(t.Sha),
		orm.ToSQLString(t.UserID),
		orm.ToSQLString(t.RepoID),
		orm.ToSQLString(t.Filename),
		orm.ToSQLString(t.Language),
		orm.ToSQLInt64(t.Ordinal),
		orm.ToSQLInt64(t.Loc),
		orm.ToSQLInt64(t.Sloc),
		orm.ToSQLInt64(t.Blanks),
		orm.ToSQLInt64(t.Comments),
	)
}

// DBCreateIgnoreDuplicateTx will upsert the CommitActivity record in the database using the provided transaction
func (t *CommitActivity) DBCreateIgnoreDuplicateTx(ctx context.Context, tx Tx) (sql.Result, error) {
	q := "INSERT INTO `commit_activity` (`commit_activity`.`id`,`commit_activity`.`date`,`commit_activity`.`sha`,`commit_activity`.`user_id`,`commit_activity`.`repo_id`,`commit_activity`.`filename`,`commit_activity`.`language`,`commit_activity`.`ordinal`,`commit_activity`.`loc`,`commit_activity`.`sloc`,`commit_activity`.`blanks`,`commit_activity`.`comments`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLInt64(t.Date),
		orm.ToSQLString(t.Sha),
		orm.ToSQLString(t.UserID),
		orm.ToSQLString(t.RepoID),
		orm.ToSQLString(t.Filename),
		orm.ToSQLString(t.Language),
		orm.ToSQLInt64(t.Ordinal),
		orm.ToSQLInt64(t.Loc),
		orm.ToSQLInt64(t.Sloc),
		orm.ToSQLInt64(t.Blanks),
		orm.ToSQLInt64(t.Comments),
	)
}

// DeleteAllCommitActivities deletes all CommitActivity records in the database with optional filters
func DeleteAllCommitActivities(ctx context.Context, db DB, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(CommitActivityTableName),
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

// DeleteAllCommitActivitiesTx deletes all CommitActivity records in the database with optional filters using the provided transaction
func DeleteAllCommitActivitiesTx(ctx context.Context, tx Tx, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(CommitActivityTableName),
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

// DBDelete will delete this CommitActivity record in the database
func (t *CommitActivity) DBDelete(ctx context.Context, db DB) (bool, error) {
	q := "DELETE FROM `commit_activity` WHERE `id` = ?"
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

// DBDeleteTx will delete this CommitActivity record in the database using the provided transaction
func (t *CommitActivity) DBDeleteTx(ctx context.Context, tx Tx) (bool, error) {
	q := "DELETE FROM `commit_activity` WHERE `id` = ?"
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

// DBUpdate will update the CommitActivity record in the database
func (t *CommitActivity) DBUpdate(ctx context.Context, db DB) (sql.Result, error) {
	q := "UPDATE `commit_activity` SET `date`=?,`sha`=?,`user_id`=?,`repo_id`=?,`filename`=?,`language`=?,`ordinal`=?,`loc`=?,`sloc`=?,`blanks`=?,`comments`=? WHERE `id`=?"
	return db.ExecContext(ctx, q,
		orm.ToSQLInt64(t.Date),
		orm.ToSQLString(t.Sha),
		orm.ToSQLString(t.UserID),
		orm.ToSQLString(t.RepoID),
		orm.ToSQLString(t.Filename),
		orm.ToSQLString(t.Language),
		orm.ToSQLInt64(t.Ordinal),
		orm.ToSQLInt64(t.Loc),
		orm.ToSQLInt64(t.Sloc),
		orm.ToSQLInt64(t.Blanks),
		orm.ToSQLInt64(t.Comments),
		orm.ToSQLString(t.ID),
	)
}

// DBUpdateTx will update the CommitActivity record in the database using the provided transaction
func (t *CommitActivity) DBUpdateTx(ctx context.Context, tx Tx) (sql.Result, error) {
	q := "UPDATE `commit_activity` SET `date`=?,`sha`=?,`user_id`=?,`repo_id`=?,`filename`=?,`language`=?,`ordinal`=?,`loc`=?,`sloc`=?,`blanks`=?,`comments`=? WHERE `id`=?"
	return tx.ExecContext(ctx, q,
		orm.ToSQLInt64(t.Date),
		orm.ToSQLString(t.Sha),
		orm.ToSQLString(t.UserID),
		orm.ToSQLString(t.RepoID),
		orm.ToSQLString(t.Filename),
		orm.ToSQLString(t.Language),
		orm.ToSQLInt64(t.Ordinal),
		orm.ToSQLInt64(t.Loc),
		orm.ToSQLInt64(t.Sloc),
		orm.ToSQLInt64(t.Blanks),
		orm.ToSQLInt64(t.Comments),
		orm.ToSQLString(t.ID),
	)
}

// DBUpsert will upsert the CommitActivity record in the database
func (t *CommitActivity) DBUpsert(ctx context.Context, db DB, conditions ...interface{}) (bool, bool, error) {
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `commit_activity` (`commit_activity`.`id`,`commit_activity`.`date`,`commit_activity`.`sha`,`commit_activity`.`user_id`,`commit_activity`.`repo_id`,`commit_activity`.`filename`,`commit_activity`.`language`,`commit_activity`.`ordinal`,`commit_activity`.`loc`,`commit_activity`.`sloc`,`commit_activity`.`blanks`,`commit_activity`.`comments`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `commit_activity` (`commit_activity`.`id`,`commit_activity`.`date`,`commit_activity`.`sha`,`commit_activity`.`user_id`,`commit_activity`.`repo_id`,`commit_activity`.`filename`,`commit_activity`.`language`,`commit_activity`.`ordinal`,`commit_activity`.`loc`,`commit_activity`.`sloc`,`commit_activity`.`blanks`,`commit_activity`.`comments`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `date`=VALUES(`date`),`sha`=VALUES(`sha`),`user_id`=VALUES(`user_id`),`repo_id`=VALUES(`repo_id`),`filename`=VALUES(`filename`),`language`=VALUES(`language`),`ordinal`=VALUES(`ordinal`),`loc`=VALUES(`loc`),`sloc`=VALUES(`sloc`),`blanks`=VALUES(`blanks`),`comments`=VALUES(`comments`)"
	}
	r, err := db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLInt64(t.Date),
		orm.ToSQLString(t.Sha),
		orm.ToSQLString(t.UserID),
		orm.ToSQLString(t.RepoID),
		orm.ToSQLString(t.Filename),
		orm.ToSQLString(t.Language),
		orm.ToSQLInt64(t.Ordinal),
		orm.ToSQLInt64(t.Loc),
		orm.ToSQLInt64(t.Sloc),
		orm.ToSQLInt64(t.Blanks),
		orm.ToSQLInt64(t.Comments),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBUpsertTx will upsert the CommitActivity record in the database using the provided transaction
func (t *CommitActivity) DBUpsertTx(ctx context.Context, tx Tx, conditions ...interface{}) (bool, bool, error) {
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `commit_activity` (`commit_activity`.`id`,`commit_activity`.`date`,`commit_activity`.`sha`,`commit_activity`.`user_id`,`commit_activity`.`repo_id`,`commit_activity`.`filename`,`commit_activity`.`language`,`commit_activity`.`ordinal`,`commit_activity`.`loc`,`commit_activity`.`sloc`,`commit_activity`.`blanks`,`commit_activity`.`comments`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `commit_activity` (`commit_activity`.`id`,`commit_activity`.`date`,`commit_activity`.`sha`,`commit_activity`.`user_id`,`commit_activity`.`repo_id`,`commit_activity`.`filename`,`commit_activity`.`language`,`commit_activity`.`ordinal`,`commit_activity`.`loc`,`commit_activity`.`sloc`,`commit_activity`.`blanks`,`commit_activity`.`comments`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `date`=VALUES(`date`),`sha`=VALUES(`sha`),`user_id`=VALUES(`user_id`),`repo_id`=VALUES(`repo_id`),`filename`=VALUES(`filename`),`language`=VALUES(`language`),`ordinal`=VALUES(`ordinal`),`loc`=VALUES(`loc`),`sloc`=VALUES(`sloc`),`blanks`=VALUES(`blanks`),`comments`=VALUES(`comments`)"
	}
	r, err := tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLInt64(t.Date),
		orm.ToSQLString(t.Sha),
		orm.ToSQLString(t.UserID),
		orm.ToSQLString(t.RepoID),
		orm.ToSQLString(t.Filename),
		orm.ToSQLString(t.Language),
		orm.ToSQLInt64(t.Ordinal),
		orm.ToSQLInt64(t.Loc),
		orm.ToSQLInt64(t.Sloc),
		orm.ToSQLInt64(t.Blanks),
		orm.ToSQLInt64(t.Comments),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBFindOne will find a CommitActivity record in the database with the primary key
func (t *CommitActivity) DBFindOne(ctx context.Context, db DB, value string) (bool, error) {
	q := "SELECT * FROM `commit_activity` WHERE `id` = ? LIMIT 1"
	row := db.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Date sql.NullInt64
	var _Sha sql.NullString
	var _UserID sql.NullString
	var _RepoID sql.NullString
	var _Filename sql.NullString
	var _Language sql.NullString
	var _Ordinal sql.NullInt64
	var _Loc sql.NullInt64
	var _Sloc sql.NullInt64
	var _Blanks sql.NullInt64
	var _Comments sql.NullInt64
	err := row.Scan(
		&_ID,
		&_Date,
		&_Sha,
		&_UserID,
		&_RepoID,
		&_Filename,
		&_Language,
		&_Ordinal,
		&_Loc,
		&_Sloc,
		&_Blanks,
		&_Comments,
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
	if _Date.Valid {
		t.SetDate(_Date.Int64)
	}
	if _Sha.Valid {
		t.SetSha(_Sha.String)
	}
	if _UserID.Valid {
		t.SetUserID(_UserID.String)
	}
	if _RepoID.Valid {
		t.SetRepoID(_RepoID.String)
	}
	if _Filename.Valid {
		t.SetFilename(_Filename.String)
	}
	if _Language.Valid {
		t.SetLanguage(_Language.String)
	}
	if _Ordinal.Valid {
		t.SetOrdinal(_Ordinal.Int64)
	}
	if _Loc.Valid {
		t.SetLoc(int32(_Loc.Int64))
	}
	if _Sloc.Valid {
		t.SetSloc(int32(_Sloc.Int64))
	}
	if _Blanks.Valid {
		t.SetBlanks(int32(_Blanks.Int64))
	}
	if _Comments.Valid {
		t.SetComments(int32(_Comments.Int64))
	}
	return true, nil
}

// DBFindOneTx will find a CommitActivity record in the database with the primary key using the provided transaction
func (t *CommitActivity) DBFindOneTx(ctx context.Context, tx Tx, value string) (bool, error) {
	q := "SELECT * FROM `commit_activity` WHERE `id` = ? LIMIT 1"
	row := tx.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Date sql.NullInt64
	var _Sha sql.NullString
	var _UserID sql.NullString
	var _RepoID sql.NullString
	var _Filename sql.NullString
	var _Language sql.NullString
	var _Ordinal sql.NullInt64
	var _Loc sql.NullInt64
	var _Sloc sql.NullInt64
	var _Blanks sql.NullInt64
	var _Comments sql.NullInt64
	err := row.Scan(
		&_ID,
		&_Date,
		&_Sha,
		&_UserID,
		&_RepoID,
		&_Filename,
		&_Language,
		&_Ordinal,
		&_Loc,
		&_Sloc,
		&_Blanks,
		&_Comments,
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
	if _Date.Valid {
		t.SetDate(_Date.Int64)
	}
	if _Sha.Valid {
		t.SetSha(_Sha.String)
	}
	if _UserID.Valid {
		t.SetUserID(_UserID.String)
	}
	if _RepoID.Valid {
		t.SetRepoID(_RepoID.String)
	}
	if _Filename.Valid {
		t.SetFilename(_Filename.String)
	}
	if _Language.Valid {
		t.SetLanguage(_Language.String)
	}
	if _Ordinal.Valid {
		t.SetOrdinal(_Ordinal.Int64)
	}
	if _Loc.Valid {
		t.SetLoc(int32(_Loc.Int64))
	}
	if _Sloc.Valid {
		t.SetSloc(int32(_Sloc.Int64))
	}
	if _Blanks.Valid {
		t.SetBlanks(int32(_Blanks.Int64))
	}
	if _Comments.Valid {
		t.SetComments(int32(_Comments.Int64))
	}
	return true, nil
}

// FindCommitActivities will find a CommitActivity record in the database with the provided parameters
func FindCommitActivities(ctx context.Context, db DB, _params ...interface{}) ([]*CommitActivity, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("date"),
		orm.Column("sha"),
		orm.Column("user_id"),
		orm.Column("repo_id"),
		orm.Column("filename"),
		orm.Column("language"),
		orm.Column("ordinal"),
		orm.Column("loc"),
		orm.Column("sloc"),
		orm.Column("blanks"),
		orm.Column("comments"),
		orm.Table(CommitActivityTableName),
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
	results := make([]*CommitActivity, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Date sql.NullInt64
		var _Sha sql.NullString
		var _UserID sql.NullString
		var _RepoID sql.NullString
		var _Filename sql.NullString
		var _Language sql.NullString
		var _Ordinal sql.NullInt64
		var _Loc sql.NullInt64
		var _Sloc sql.NullInt64
		var _Blanks sql.NullInt64
		var _Comments sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Date,
			&_Sha,
			&_UserID,
			&_RepoID,
			&_Filename,
			&_Language,
			&_Ordinal,
			&_Loc,
			&_Sloc,
			&_Blanks,
			&_Comments,
		)
		if err != nil {
			return nil, err
		}
		t := &CommitActivity{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Date.Valid {
			t.SetDate(_Date.Int64)
		}
		if _Sha.Valid {
			t.SetSha(_Sha.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _RepoID.Valid {
			t.SetRepoID(_RepoID.String)
		}
		if _Filename.Valid {
			t.SetFilename(_Filename.String)
		}
		if _Language.Valid {
			t.SetLanguage(_Language.String)
		}
		if _Ordinal.Valid {
			t.SetOrdinal(_Ordinal.Int64)
		}
		if _Loc.Valid {
			t.SetLoc(int32(_Loc.Int64))
		}
		if _Sloc.Valid {
			t.SetSloc(int32(_Sloc.Int64))
		}
		if _Blanks.Valid {
			t.SetBlanks(int32(_Blanks.Int64))
		}
		if _Comments.Valid {
			t.SetComments(int32(_Comments.Int64))
		}
		results = append(results, t)
	}
	return results, nil
}

// FindCommitActivitiesTx will find a CommitActivity record in the database with the provided parameters using the provided transaction
func FindCommitActivitiesTx(ctx context.Context, tx Tx, _params ...interface{}) ([]*CommitActivity, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("date"),
		orm.Column("sha"),
		orm.Column("user_id"),
		orm.Column("repo_id"),
		orm.Column("filename"),
		orm.Column("language"),
		orm.Column("ordinal"),
		orm.Column("loc"),
		orm.Column("sloc"),
		orm.Column("blanks"),
		orm.Column("comments"),
		orm.Table(CommitActivityTableName),
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
	results := make([]*CommitActivity, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Date sql.NullInt64
		var _Sha sql.NullString
		var _UserID sql.NullString
		var _RepoID sql.NullString
		var _Filename sql.NullString
		var _Language sql.NullString
		var _Ordinal sql.NullInt64
		var _Loc sql.NullInt64
		var _Sloc sql.NullInt64
		var _Blanks sql.NullInt64
		var _Comments sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Date,
			&_Sha,
			&_UserID,
			&_RepoID,
			&_Filename,
			&_Language,
			&_Ordinal,
			&_Loc,
			&_Sloc,
			&_Blanks,
			&_Comments,
		)
		if err != nil {
			return nil, err
		}
		t := &CommitActivity{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Date.Valid {
			t.SetDate(_Date.Int64)
		}
		if _Sha.Valid {
			t.SetSha(_Sha.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _RepoID.Valid {
			t.SetRepoID(_RepoID.String)
		}
		if _Filename.Valid {
			t.SetFilename(_Filename.String)
		}
		if _Language.Valid {
			t.SetLanguage(_Language.String)
		}
		if _Ordinal.Valid {
			t.SetOrdinal(_Ordinal.Int64)
		}
		if _Loc.Valid {
			t.SetLoc(int32(_Loc.Int64))
		}
		if _Sloc.Valid {
			t.SetSloc(int32(_Sloc.Int64))
		}
		if _Blanks.Valid {
			t.SetBlanks(int32(_Blanks.Int64))
		}
		if _Comments.Valid {
			t.SetComments(int32(_Comments.Int64))
		}
		results = append(results, t)
	}
	return results, nil
}

// DBFind will find a CommitActivity record in the database with the provided parameters
func (t *CommitActivity) DBFind(ctx context.Context, db DB, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("date"),
		orm.Column("sha"),
		orm.Column("user_id"),
		orm.Column("repo_id"),
		orm.Column("filename"),
		orm.Column("language"),
		orm.Column("ordinal"),
		orm.Column("loc"),
		orm.Column("sloc"),
		orm.Column("blanks"),
		orm.Column("comments"),
		orm.Table(CommitActivityTableName),
	}
	if len(_params) > 0 {
		for _, param := range _params {
			params = append(params, param)
		}
	}
	q, p := orm.BuildQuery(params...)
	row := db.QueryRowContext(ctx, q, p...)
	var _ID sql.NullString
	var _Date sql.NullInt64
	var _Sha sql.NullString
	var _UserID sql.NullString
	var _RepoID sql.NullString
	var _Filename sql.NullString
	var _Language sql.NullString
	var _Ordinal sql.NullInt64
	var _Loc sql.NullInt64
	var _Sloc sql.NullInt64
	var _Blanks sql.NullInt64
	var _Comments sql.NullInt64
	err := row.Scan(
		&_ID,
		&_Date,
		&_Sha,
		&_UserID,
		&_RepoID,
		&_Filename,
		&_Language,
		&_Ordinal,
		&_Loc,
		&_Sloc,
		&_Blanks,
		&_Comments,
	)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _Date.Valid {
		t.SetDate(_Date.Int64)
	}
	if _Sha.Valid {
		t.SetSha(_Sha.String)
	}
	if _UserID.Valid {
		t.SetUserID(_UserID.String)
	}
	if _RepoID.Valid {
		t.SetRepoID(_RepoID.String)
	}
	if _Filename.Valid {
		t.SetFilename(_Filename.String)
	}
	if _Language.Valid {
		t.SetLanguage(_Language.String)
	}
	if _Ordinal.Valid {
		t.SetOrdinal(_Ordinal.Int64)
	}
	if _Loc.Valid {
		t.SetLoc(int32(_Loc.Int64))
	}
	if _Sloc.Valid {
		t.SetSloc(int32(_Sloc.Int64))
	}
	if _Blanks.Valid {
		t.SetBlanks(int32(_Blanks.Int64))
	}
	if _Comments.Valid {
		t.SetComments(int32(_Comments.Int64))
	}
	return true, nil
}

// DBFindTx will find a CommitActivity record in the database with the provided parameters using the provided transaction
func (t *CommitActivity) DBFindTx(ctx context.Context, tx Tx, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("date"),
		orm.Column("sha"),
		orm.Column("user_id"),
		orm.Column("repo_id"),
		orm.Column("filename"),
		orm.Column("language"),
		orm.Column("ordinal"),
		orm.Column("loc"),
		orm.Column("sloc"),
		orm.Column("blanks"),
		orm.Column("comments"),
		orm.Table(CommitActivityTableName),
	}
	if len(_params) > 0 {
		for _, param := range _params {
			params = append(params, param)
		}
	}
	q, p := orm.BuildQuery(params...)
	row := tx.QueryRowContext(ctx, q, p...)
	var _ID sql.NullString
	var _Date sql.NullInt64
	var _Sha sql.NullString
	var _UserID sql.NullString
	var _RepoID sql.NullString
	var _Filename sql.NullString
	var _Language sql.NullString
	var _Ordinal sql.NullInt64
	var _Loc sql.NullInt64
	var _Sloc sql.NullInt64
	var _Blanks sql.NullInt64
	var _Comments sql.NullInt64
	err := row.Scan(
		&_ID,
		&_Date,
		&_Sha,
		&_UserID,
		&_RepoID,
		&_Filename,
		&_Language,
		&_Ordinal,
		&_Loc,
		&_Sloc,
		&_Blanks,
		&_Comments,
	)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _Date.Valid {
		t.SetDate(_Date.Int64)
	}
	if _Sha.Valid {
		t.SetSha(_Sha.String)
	}
	if _UserID.Valid {
		t.SetUserID(_UserID.String)
	}
	if _RepoID.Valid {
		t.SetRepoID(_RepoID.String)
	}
	if _Filename.Valid {
		t.SetFilename(_Filename.String)
	}
	if _Language.Valid {
		t.SetLanguage(_Language.String)
	}
	if _Ordinal.Valid {
		t.SetOrdinal(_Ordinal.Int64)
	}
	if _Loc.Valid {
		t.SetLoc(int32(_Loc.Int64))
	}
	if _Sloc.Valid {
		t.SetSloc(int32(_Sloc.Int64))
	}
	if _Blanks.Valid {
		t.SetBlanks(int32(_Blanks.Int64))
	}
	if _Comments.Valid {
		t.SetComments(int32(_Comments.Int64))
	}
	return true, nil
}

// CountCommitActivities will find the count of CommitActivity records in the database
func CountCommitActivities(ctx context.Context, db DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(CommitActivityTableName),
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

// CountCommitActivitiesTx will find the count of CommitActivity records in the database using the provided transaction
func CountCommitActivitiesTx(ctx context.Context, tx Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(CommitActivityTableName),
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

// DBCount will find the count of CommitActivity records in the database
func (t *CommitActivity) DBCount(ctx context.Context, db DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(CommitActivityTableName),
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

// DBCountTx will find the count of CommitActivity records in the database using the provided transaction
func (t *CommitActivity) DBCountTx(ctx context.Context, tx Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(CommitActivityTableName),
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

// DBExists will return true if the CommitActivity record exists in the database
func (t *CommitActivity) DBExists(ctx context.Context, db DB) (bool, error) {
	q := "SELECT * FROM `commit_activity` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := db.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// DBExistsTx will return true if the CommitActivity record exists in the database using the provided transaction
func (t *CommitActivity) DBExistsTx(ctx context.Context, tx Tx) (bool, error) {
	q := "SELECT * FROM `commit_activity` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := tx.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// PrimaryKeyColumn returns the column name for the primary key
func (t *CommitActivity) PrimaryKeyColumn() string {
	return CommitActivityColumnID
}

// PrimaryKeyColumnType returns the primary key column Go type as a string
func (t *CommitActivity) PrimaryKeyColumnType() string {
	return "string"
}

// PrimaryKey returns the primary key column value
func (t *CommitActivity) PrimaryKey() interface{} {
	return t.ID
}
