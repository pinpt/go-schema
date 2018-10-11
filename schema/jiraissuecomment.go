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

var _ Model = (*JiraIssueComment)(nil)
var _ CSVWriter = (*JiraIssueComment)(nil)
var _ JSONWriter = (*JiraIssueComment)(nil)
var _ Checksum = (*JiraIssueComment)(nil)

// JiraIssueCommentTableName is the name of the table in SQL
const JiraIssueCommentTableName = "jira_issue_comment"

var JiraIssueCommentColumns = []string{
	"id",
	"checksum",
	"comment_id",
	"issue_id",
	"project_id",
	"user_id",
	"updated_at",
	"customer_id",
	"ref_id",
}

// JiraIssueComment table
type JiraIssueComment struct {
	Checksum   *string `json:"checksum,omitempty"`
	CommentID  string  `json:"comment_id"`
	CustomerID string  `json:"customer_id"`
	ID         string  `json:"id"`
	IssueID    string  `json:"issue_id"`
	ProjectID  string  `json:"project_id"`
	RefID      string  `json:"ref_id"`
	UpdatedAt  int64   `json:"updated_at"`
	UserID     *string `json:"user_id,omitempty"`
}

// TableName returns the SQL table name for JiraIssueComment and satifies the Model interface
func (t *JiraIssueComment) TableName() string {
	return JiraIssueCommentTableName
}

// ToCSV will serialize the JiraIssueComment instance to a CSV compatible array of strings
func (t *JiraIssueComment) ToCSV() []string {
	return []string{
		t.ID,
		t.CalculateChecksum(),
		t.CommentID,
		t.IssueID,
		t.ProjectID,
		toCSVString(t.UserID),
		toCSVString(t.UpdatedAt),
		t.CustomerID,
		t.RefID,
	}
}

// WriteCSV will serialize the JiraIssueComment instance to the writer as CSV and satisfies the CSVWriter interface
func (t *JiraIssueComment) WriteCSV(w *csv.Writer) error {
	return w.Write(t.ToCSV())
}

// WriteJSON will serialize the JiraIssueComment instance to the writer as JSON and satisfies the JSONWriter interface
func (t *JiraIssueComment) WriteJSON(w io.Writer, indent ...bool) error {
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

// NewJiraIssueCommentReader creates a JSON reader which can read in JiraIssueComment objects serialized as JSON either as an array, single object or json new lines
// and writes each JiraIssueComment to the channel provided
func NewJiraIssueCommentReader(r io.Reader, ch chan<- JiraIssueComment) error {
	return orm.Deserialize(r, func(buf json.RawMessage) error {
		dec := json.NewDecoder(bytes.NewBuffer(buf))
		e := JiraIssueComment{}
		if err := dec.Decode(&e); err != nil {
			return err
		}
		ch <- e
		return nil
	})
}

// NewCSVJiraIssueCommentReaderDir will read the reader as CSV and emit each record to the channel provided
func NewCSVJiraIssueCommentReader(r io.Reader, ch chan<- JiraIssueComment) error {
	cr := csv.NewReader(r)
	for {
		record, err := cr.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		ch <- JiraIssueComment{
			ID:         record[0],
			Checksum:   fromStringPointer(record[1]),
			CommentID:  record[2],
			IssueID:    record[3],
			ProjectID:  record[4],
			UserID:     fromStringPointer(record[5]),
			UpdatedAt:  fromCSVInt64(record[6]),
			CustomerID: record[7],
			RefID:      record[8],
		}
	}
	return nil
}

// NewCSVJiraIssueCommentReaderFile will read the file as a CSV and emit each record to the channel provided
func NewCSVJiraIssueCommentReaderFile(fp string, ch chan<- JiraIssueComment) error {
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
	return NewCSVJiraIssueCommentReader(fc, ch)
}

// NewCSVJiraIssueCommentReaderDir will read the jira_issue_comment.csv.gz file as a CSV and emit each record to the channel provided
func NewCSVJiraIssueCommentReaderDir(dir string, ch chan<- JiraIssueComment) error {
	return NewCSVJiraIssueCommentReaderFile(filepath.Join(dir, "jira_issue_comment.csv.gz"), ch)
}

// JiraIssueCommentCSVDeduper is a function callback which takes the existing value (a) and the new value (b)
// and the return value should be the one to use (or a new one, if applicable). return nil
// to skip processing of this record
type JiraIssueCommentCSVDeduper func(a JiraIssueComment, b JiraIssueComment) *JiraIssueComment

// JiraIssueCommentCSVDedupeDisabled is set on whether the CSV writer should de-dupe values by key
var JiraIssueCommentCSVDedupeDisabled bool

// NewJiraIssueCommentCSVWriterSize creates a batch writer that will write each JiraIssueComment into a CSV file
// this method will automatically de-duplicate entries using the primary key. if the checksum
// for a newer item with the same primary key doesn't match a previously sent item, the newer
// one will be used
func NewJiraIssueCommentCSVWriterSize(w io.Writer, size int, dedupers ...JiraIssueCommentCSVDeduper) (chan JiraIssueComment, chan bool, error) {
	cw := csv.NewWriter(w)
	ch := make(chan JiraIssueComment, size)
	done := make(chan bool)
	go func() {
		defer func() { done <- true }()
		dodedupe := !JiraIssueCommentCSVDedupeDisabled
		var kv map[string]*JiraIssueComment
		var deduper JiraIssueCommentCSVDeduper
		if dedupers != nil && len(dedupers) > 0 {
			deduper = dedupers[0]
			dodedupe = true
		}
		if dodedupe {
			kv = make(map[string]*JiraIssueComment)
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

// JiraIssueCommentCSVDefaultSize is the default channel buffer size if not provided
var JiraIssueCommentCSVDefaultSize = 100

// NewJiraIssueCommentCSVWriter creates a batch writer that will write each JiraIssueComment into a CSV file
func NewJiraIssueCommentCSVWriter(w io.Writer, dedupers ...JiraIssueCommentCSVDeduper) (chan JiraIssueComment, chan bool, error) {
	return NewJiraIssueCommentCSVWriterSize(w, JiraIssueCommentCSVDefaultSize, dedupers...)
}

// NewJiraIssueCommentCSVWriterDir creates a batch writer that will write each JiraIssueComment into a CSV file named jira_issue_comment.csv.gz in dir
func NewJiraIssueCommentCSVWriterDir(dir string, dedupers ...JiraIssueCommentCSVDeduper) (chan JiraIssueComment, chan bool, error) {
	return NewJiraIssueCommentCSVWriterFile(filepath.Join(dir, "jira_issue_comment.csv.gz"), dedupers...)
}

// NewJiraIssueCommentCSVWriterFile creates a batch writer that will write each JiraIssueComment into a CSV file
func NewJiraIssueCommentCSVWriterFile(fn string, dedupers ...JiraIssueCommentCSVDeduper) (chan JiraIssueComment, chan bool, error) {
	f, err := os.Create(fn)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening CSV file %s. %v", fn, err)
	}
	var fc io.WriteCloser = f
	if filepath.Ext(fn) == ".gz" {
		w, _ := gzip.NewWriterLevel(f, gzip.BestCompression)
		fc = w
	}
	ch, done, err := NewJiraIssueCommentCSVWriter(fc, dedupers...)
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

type JiraIssueCommentDBAction func(ctx context.Context, db DB, record JiraIssueComment) error

// NewJiraIssueCommentDBWriterSize creates a DB writer that will write each issue into the DB
func NewJiraIssueCommentDBWriterSize(ctx context.Context, db DB, errors chan<- error, size int, actions ...JiraIssueCommentDBAction) (chan JiraIssueComment, chan bool, error) {
	ch := make(chan JiraIssueComment, size)
	done := make(chan bool)
	var action JiraIssueCommentDBAction
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

// NewJiraIssueCommentDBWriter creates a DB writer that will write each issue into the DB
func NewJiraIssueCommentDBWriter(ctx context.Context, db DB, errors chan<- error, actions ...JiraIssueCommentDBAction) (chan JiraIssueComment, chan bool, error) {
	return NewJiraIssueCommentDBWriterSize(ctx, db, errors, 100, actions...)
}

// JiraIssueCommentColumnID is the ID SQL column name for the JiraIssueComment table
const JiraIssueCommentColumnID = "id"

// JiraIssueCommentEscapedColumnID is the escaped ID SQL column name for the JiraIssueComment table
const JiraIssueCommentEscapedColumnID = "`id`"

// JiraIssueCommentColumnChecksum is the Checksum SQL column name for the JiraIssueComment table
const JiraIssueCommentColumnChecksum = "checksum"

// JiraIssueCommentEscapedColumnChecksum is the escaped Checksum SQL column name for the JiraIssueComment table
const JiraIssueCommentEscapedColumnChecksum = "`checksum`"

// JiraIssueCommentColumnCommentID is the CommentID SQL column name for the JiraIssueComment table
const JiraIssueCommentColumnCommentID = "comment_id"

// JiraIssueCommentEscapedColumnCommentID is the escaped CommentID SQL column name for the JiraIssueComment table
const JiraIssueCommentEscapedColumnCommentID = "`comment_id`"

// JiraIssueCommentColumnIssueID is the IssueID SQL column name for the JiraIssueComment table
const JiraIssueCommentColumnIssueID = "issue_id"

// JiraIssueCommentEscapedColumnIssueID is the escaped IssueID SQL column name for the JiraIssueComment table
const JiraIssueCommentEscapedColumnIssueID = "`issue_id`"

// JiraIssueCommentColumnProjectID is the ProjectID SQL column name for the JiraIssueComment table
const JiraIssueCommentColumnProjectID = "project_id"

// JiraIssueCommentEscapedColumnProjectID is the escaped ProjectID SQL column name for the JiraIssueComment table
const JiraIssueCommentEscapedColumnProjectID = "`project_id`"

// JiraIssueCommentColumnUserID is the UserID SQL column name for the JiraIssueComment table
const JiraIssueCommentColumnUserID = "user_id"

// JiraIssueCommentEscapedColumnUserID is the escaped UserID SQL column name for the JiraIssueComment table
const JiraIssueCommentEscapedColumnUserID = "`user_id`"

// JiraIssueCommentColumnUpdatedAt is the UpdatedAt SQL column name for the JiraIssueComment table
const JiraIssueCommentColumnUpdatedAt = "updated_at"

// JiraIssueCommentEscapedColumnUpdatedAt is the escaped UpdatedAt SQL column name for the JiraIssueComment table
const JiraIssueCommentEscapedColumnUpdatedAt = "`updated_at`"

// JiraIssueCommentColumnCustomerID is the CustomerID SQL column name for the JiraIssueComment table
const JiraIssueCommentColumnCustomerID = "customer_id"

// JiraIssueCommentEscapedColumnCustomerID is the escaped CustomerID SQL column name for the JiraIssueComment table
const JiraIssueCommentEscapedColumnCustomerID = "`customer_id`"

// JiraIssueCommentColumnRefID is the RefID SQL column name for the JiraIssueComment table
const JiraIssueCommentColumnRefID = "ref_id"

// JiraIssueCommentEscapedColumnRefID is the escaped RefID SQL column name for the JiraIssueComment table
const JiraIssueCommentEscapedColumnRefID = "`ref_id`"

// GetID will return the JiraIssueComment ID value
func (t *JiraIssueComment) GetID() string {
	return t.ID
}

// SetID will set the JiraIssueComment ID value
func (t *JiraIssueComment) SetID(v string) {
	t.ID = v
}

// FindJiraIssueCommentByID will find a JiraIssueComment by ID
func FindJiraIssueCommentByID(ctx context.Context, db DB, value string) (*JiraIssueComment, error) {
	q := "SELECT `jira_issue_comment`.`id`,`jira_issue_comment`.`checksum`,`jira_issue_comment`.`comment_id`,`jira_issue_comment`.`issue_id`,`jira_issue_comment`.`project_id`,`jira_issue_comment`.`user_id`,`jira_issue_comment`.`updated_at`,`jira_issue_comment`.`customer_id`,`jira_issue_comment`.`ref_id` FROM `jira_issue_comment` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CommentID sql.NullString
	var _IssueID sql.NullString
	var _ProjectID sql.NullString
	var _UserID sql.NullString
	var _UpdatedAt sql.NullInt64
	var _CustomerID sql.NullString
	var _RefID sql.NullString
	err := db.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_CommentID,
		&_IssueID,
		&_ProjectID,
		&_UserID,
		&_UpdatedAt,
		&_CustomerID,
		&_RefID,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &JiraIssueComment{}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _Checksum.Valid {
		t.SetChecksum(_Checksum.String)
	}
	if _CommentID.Valid {
		t.SetCommentID(_CommentID.String)
	}
	if _IssueID.Valid {
		t.SetIssueID(_IssueID.String)
	}
	if _ProjectID.Valid {
		t.SetProjectID(_ProjectID.String)
	}
	if _UserID.Valid {
		t.SetUserID(_UserID.String)
	}
	if _UpdatedAt.Valid {
		t.SetUpdatedAt(_UpdatedAt.Int64)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _RefID.Valid {
		t.SetRefID(_RefID.String)
	}
	return t, nil
}

// FindJiraIssueCommentByIDTx will find a JiraIssueComment by ID using the provided transaction
func FindJiraIssueCommentByIDTx(ctx context.Context, tx Tx, value string) (*JiraIssueComment, error) {
	q := "SELECT `jira_issue_comment`.`id`,`jira_issue_comment`.`checksum`,`jira_issue_comment`.`comment_id`,`jira_issue_comment`.`issue_id`,`jira_issue_comment`.`project_id`,`jira_issue_comment`.`user_id`,`jira_issue_comment`.`updated_at`,`jira_issue_comment`.`customer_id`,`jira_issue_comment`.`ref_id` FROM `jira_issue_comment` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CommentID sql.NullString
	var _IssueID sql.NullString
	var _ProjectID sql.NullString
	var _UserID sql.NullString
	var _UpdatedAt sql.NullInt64
	var _CustomerID sql.NullString
	var _RefID sql.NullString
	err := tx.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_CommentID,
		&_IssueID,
		&_ProjectID,
		&_UserID,
		&_UpdatedAt,
		&_CustomerID,
		&_RefID,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &JiraIssueComment{}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _Checksum.Valid {
		t.SetChecksum(_Checksum.String)
	}
	if _CommentID.Valid {
		t.SetCommentID(_CommentID.String)
	}
	if _IssueID.Valid {
		t.SetIssueID(_IssueID.String)
	}
	if _ProjectID.Valid {
		t.SetProjectID(_ProjectID.String)
	}
	if _UserID.Valid {
		t.SetUserID(_UserID.String)
	}
	if _UpdatedAt.Valid {
		t.SetUpdatedAt(_UpdatedAt.Int64)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _RefID.Valid {
		t.SetRefID(_RefID.String)
	}
	return t, nil
}

// GetChecksum will return the JiraIssueComment Checksum value
func (t *JiraIssueComment) GetChecksum() string {
	if t.Checksum == nil {
		return ""
	}
	return *t.Checksum
}

// SetChecksum will set the JiraIssueComment Checksum value
func (t *JiraIssueComment) SetChecksum(v string) {
	t.Checksum = &v
}

// GetCommentID will return the JiraIssueComment CommentID value
func (t *JiraIssueComment) GetCommentID() string {
	return t.CommentID
}

// SetCommentID will set the JiraIssueComment CommentID value
func (t *JiraIssueComment) SetCommentID(v string) {
	t.CommentID = v
}

// FindJiraIssueCommentsByCommentID will find all JiraIssueComments by the CommentID value
func FindJiraIssueCommentsByCommentID(ctx context.Context, db DB, value string) ([]*JiraIssueComment, error) {
	q := "SELECT `jira_issue_comment`.`id`,`jira_issue_comment`.`checksum`,`jira_issue_comment`.`comment_id`,`jira_issue_comment`.`issue_id`,`jira_issue_comment`.`project_id`,`jira_issue_comment`.`user_id`,`jira_issue_comment`.`updated_at`,`jira_issue_comment`.`customer_id`,`jira_issue_comment`.`ref_id` FROM `jira_issue_comment` WHERE `comment_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraIssueComment, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CommentID sql.NullString
		var _IssueID sql.NullString
		var _ProjectID sql.NullString
		var _UserID sql.NullString
		var _UpdatedAt sql.NullInt64
		var _CustomerID sql.NullString
		var _RefID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CommentID,
			&_IssueID,
			&_ProjectID,
			&_UserID,
			&_UpdatedAt,
			&_CustomerID,
			&_RefID,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraIssueComment{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CommentID.Valid {
			t.SetCommentID(_CommentID.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _UpdatedAt.Valid {
			t.SetUpdatedAt(_UpdatedAt.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindJiraIssueCommentsByCommentIDTx will find all JiraIssueComments by the CommentID value using the provided transaction
func FindJiraIssueCommentsByCommentIDTx(ctx context.Context, tx Tx, value string) ([]*JiraIssueComment, error) {
	q := "SELECT `jira_issue_comment`.`id`,`jira_issue_comment`.`checksum`,`jira_issue_comment`.`comment_id`,`jira_issue_comment`.`issue_id`,`jira_issue_comment`.`project_id`,`jira_issue_comment`.`user_id`,`jira_issue_comment`.`updated_at`,`jira_issue_comment`.`customer_id`,`jira_issue_comment`.`ref_id` FROM `jira_issue_comment` WHERE `comment_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraIssueComment, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CommentID sql.NullString
		var _IssueID sql.NullString
		var _ProjectID sql.NullString
		var _UserID sql.NullString
		var _UpdatedAt sql.NullInt64
		var _CustomerID sql.NullString
		var _RefID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CommentID,
			&_IssueID,
			&_ProjectID,
			&_UserID,
			&_UpdatedAt,
			&_CustomerID,
			&_RefID,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraIssueComment{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CommentID.Valid {
			t.SetCommentID(_CommentID.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _UpdatedAt.Valid {
			t.SetUpdatedAt(_UpdatedAt.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetIssueID will return the JiraIssueComment IssueID value
func (t *JiraIssueComment) GetIssueID() string {
	return t.IssueID
}

// SetIssueID will set the JiraIssueComment IssueID value
func (t *JiraIssueComment) SetIssueID(v string) {
	t.IssueID = v
}

// FindJiraIssueCommentsByIssueID will find all JiraIssueComments by the IssueID value
func FindJiraIssueCommentsByIssueID(ctx context.Context, db DB, value string) ([]*JiraIssueComment, error) {
	q := "SELECT `jira_issue_comment`.`id`,`jira_issue_comment`.`checksum`,`jira_issue_comment`.`comment_id`,`jira_issue_comment`.`issue_id`,`jira_issue_comment`.`project_id`,`jira_issue_comment`.`user_id`,`jira_issue_comment`.`updated_at`,`jira_issue_comment`.`customer_id`,`jira_issue_comment`.`ref_id` FROM `jira_issue_comment` WHERE `issue_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraIssueComment, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CommentID sql.NullString
		var _IssueID sql.NullString
		var _ProjectID sql.NullString
		var _UserID sql.NullString
		var _UpdatedAt sql.NullInt64
		var _CustomerID sql.NullString
		var _RefID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CommentID,
			&_IssueID,
			&_ProjectID,
			&_UserID,
			&_UpdatedAt,
			&_CustomerID,
			&_RefID,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraIssueComment{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CommentID.Valid {
			t.SetCommentID(_CommentID.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _UpdatedAt.Valid {
			t.SetUpdatedAt(_UpdatedAt.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindJiraIssueCommentsByIssueIDTx will find all JiraIssueComments by the IssueID value using the provided transaction
func FindJiraIssueCommentsByIssueIDTx(ctx context.Context, tx Tx, value string) ([]*JiraIssueComment, error) {
	q := "SELECT `jira_issue_comment`.`id`,`jira_issue_comment`.`checksum`,`jira_issue_comment`.`comment_id`,`jira_issue_comment`.`issue_id`,`jira_issue_comment`.`project_id`,`jira_issue_comment`.`user_id`,`jira_issue_comment`.`updated_at`,`jira_issue_comment`.`customer_id`,`jira_issue_comment`.`ref_id` FROM `jira_issue_comment` WHERE `issue_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraIssueComment, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CommentID sql.NullString
		var _IssueID sql.NullString
		var _ProjectID sql.NullString
		var _UserID sql.NullString
		var _UpdatedAt sql.NullInt64
		var _CustomerID sql.NullString
		var _RefID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CommentID,
			&_IssueID,
			&_ProjectID,
			&_UserID,
			&_UpdatedAt,
			&_CustomerID,
			&_RefID,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraIssueComment{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CommentID.Valid {
			t.SetCommentID(_CommentID.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _UpdatedAt.Valid {
			t.SetUpdatedAt(_UpdatedAt.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetProjectID will return the JiraIssueComment ProjectID value
func (t *JiraIssueComment) GetProjectID() string {
	return t.ProjectID
}

// SetProjectID will set the JiraIssueComment ProjectID value
func (t *JiraIssueComment) SetProjectID(v string) {
	t.ProjectID = v
}

// FindJiraIssueCommentsByProjectID will find all JiraIssueComments by the ProjectID value
func FindJiraIssueCommentsByProjectID(ctx context.Context, db DB, value string) ([]*JiraIssueComment, error) {
	q := "SELECT `jira_issue_comment`.`id`,`jira_issue_comment`.`checksum`,`jira_issue_comment`.`comment_id`,`jira_issue_comment`.`issue_id`,`jira_issue_comment`.`project_id`,`jira_issue_comment`.`user_id`,`jira_issue_comment`.`updated_at`,`jira_issue_comment`.`customer_id`,`jira_issue_comment`.`ref_id` FROM `jira_issue_comment` WHERE `project_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraIssueComment, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CommentID sql.NullString
		var _IssueID sql.NullString
		var _ProjectID sql.NullString
		var _UserID sql.NullString
		var _UpdatedAt sql.NullInt64
		var _CustomerID sql.NullString
		var _RefID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CommentID,
			&_IssueID,
			&_ProjectID,
			&_UserID,
			&_UpdatedAt,
			&_CustomerID,
			&_RefID,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraIssueComment{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CommentID.Valid {
			t.SetCommentID(_CommentID.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _UpdatedAt.Valid {
			t.SetUpdatedAt(_UpdatedAt.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindJiraIssueCommentsByProjectIDTx will find all JiraIssueComments by the ProjectID value using the provided transaction
func FindJiraIssueCommentsByProjectIDTx(ctx context.Context, tx Tx, value string) ([]*JiraIssueComment, error) {
	q := "SELECT `jira_issue_comment`.`id`,`jira_issue_comment`.`checksum`,`jira_issue_comment`.`comment_id`,`jira_issue_comment`.`issue_id`,`jira_issue_comment`.`project_id`,`jira_issue_comment`.`user_id`,`jira_issue_comment`.`updated_at`,`jira_issue_comment`.`customer_id`,`jira_issue_comment`.`ref_id` FROM `jira_issue_comment` WHERE `project_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraIssueComment, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CommentID sql.NullString
		var _IssueID sql.NullString
		var _ProjectID sql.NullString
		var _UserID sql.NullString
		var _UpdatedAt sql.NullInt64
		var _CustomerID sql.NullString
		var _RefID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CommentID,
			&_IssueID,
			&_ProjectID,
			&_UserID,
			&_UpdatedAt,
			&_CustomerID,
			&_RefID,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraIssueComment{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CommentID.Valid {
			t.SetCommentID(_CommentID.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _UpdatedAt.Valid {
			t.SetUpdatedAt(_UpdatedAt.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetUserID will return the JiraIssueComment UserID value
func (t *JiraIssueComment) GetUserID() string {
	if t.UserID == nil {
		return ""
	}
	return *t.UserID
}

// SetUserID will set the JiraIssueComment UserID value
func (t *JiraIssueComment) SetUserID(v string) {
	t.UserID = &v
}

// FindJiraIssueCommentsByUserID will find all JiraIssueComments by the UserID value
func FindJiraIssueCommentsByUserID(ctx context.Context, db DB, value string) ([]*JiraIssueComment, error) {
	q := "SELECT `jira_issue_comment`.`id`,`jira_issue_comment`.`checksum`,`jira_issue_comment`.`comment_id`,`jira_issue_comment`.`issue_id`,`jira_issue_comment`.`project_id`,`jira_issue_comment`.`user_id`,`jira_issue_comment`.`updated_at`,`jira_issue_comment`.`customer_id`,`jira_issue_comment`.`ref_id` FROM `jira_issue_comment` WHERE `user_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraIssueComment, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CommentID sql.NullString
		var _IssueID sql.NullString
		var _ProjectID sql.NullString
		var _UserID sql.NullString
		var _UpdatedAt sql.NullInt64
		var _CustomerID sql.NullString
		var _RefID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CommentID,
			&_IssueID,
			&_ProjectID,
			&_UserID,
			&_UpdatedAt,
			&_CustomerID,
			&_RefID,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraIssueComment{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CommentID.Valid {
			t.SetCommentID(_CommentID.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _UpdatedAt.Valid {
			t.SetUpdatedAt(_UpdatedAt.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindJiraIssueCommentsByUserIDTx will find all JiraIssueComments by the UserID value using the provided transaction
func FindJiraIssueCommentsByUserIDTx(ctx context.Context, tx Tx, value string) ([]*JiraIssueComment, error) {
	q := "SELECT `jira_issue_comment`.`id`,`jira_issue_comment`.`checksum`,`jira_issue_comment`.`comment_id`,`jira_issue_comment`.`issue_id`,`jira_issue_comment`.`project_id`,`jira_issue_comment`.`user_id`,`jira_issue_comment`.`updated_at`,`jira_issue_comment`.`customer_id`,`jira_issue_comment`.`ref_id` FROM `jira_issue_comment` WHERE `user_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraIssueComment, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CommentID sql.NullString
		var _IssueID sql.NullString
		var _ProjectID sql.NullString
		var _UserID sql.NullString
		var _UpdatedAt sql.NullInt64
		var _CustomerID sql.NullString
		var _RefID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CommentID,
			&_IssueID,
			&_ProjectID,
			&_UserID,
			&_UpdatedAt,
			&_CustomerID,
			&_RefID,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraIssueComment{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CommentID.Valid {
			t.SetCommentID(_CommentID.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _UpdatedAt.Valid {
			t.SetUpdatedAt(_UpdatedAt.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetUpdatedAt will return the JiraIssueComment UpdatedAt value
func (t *JiraIssueComment) GetUpdatedAt() int64 {
	return t.UpdatedAt
}

// SetUpdatedAt will set the JiraIssueComment UpdatedAt value
func (t *JiraIssueComment) SetUpdatedAt(v int64) {
	t.UpdatedAt = v
}

// GetCustomerID will return the JiraIssueComment CustomerID value
func (t *JiraIssueComment) GetCustomerID() string {
	return t.CustomerID
}

// SetCustomerID will set the JiraIssueComment CustomerID value
func (t *JiraIssueComment) SetCustomerID(v string) {
	t.CustomerID = v
}

// FindJiraIssueCommentsByCustomerID will find all JiraIssueComments by the CustomerID value
func FindJiraIssueCommentsByCustomerID(ctx context.Context, db DB, value string) ([]*JiraIssueComment, error) {
	q := "SELECT `jira_issue_comment`.`id`,`jira_issue_comment`.`checksum`,`jira_issue_comment`.`comment_id`,`jira_issue_comment`.`issue_id`,`jira_issue_comment`.`project_id`,`jira_issue_comment`.`user_id`,`jira_issue_comment`.`updated_at`,`jira_issue_comment`.`customer_id`,`jira_issue_comment`.`ref_id` FROM `jira_issue_comment` WHERE `customer_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraIssueComment, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CommentID sql.NullString
		var _IssueID sql.NullString
		var _ProjectID sql.NullString
		var _UserID sql.NullString
		var _UpdatedAt sql.NullInt64
		var _CustomerID sql.NullString
		var _RefID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CommentID,
			&_IssueID,
			&_ProjectID,
			&_UserID,
			&_UpdatedAt,
			&_CustomerID,
			&_RefID,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraIssueComment{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CommentID.Valid {
			t.SetCommentID(_CommentID.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _UpdatedAt.Valid {
			t.SetUpdatedAt(_UpdatedAt.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindJiraIssueCommentsByCustomerIDTx will find all JiraIssueComments by the CustomerID value using the provided transaction
func FindJiraIssueCommentsByCustomerIDTx(ctx context.Context, tx Tx, value string) ([]*JiraIssueComment, error) {
	q := "SELECT `jira_issue_comment`.`id`,`jira_issue_comment`.`checksum`,`jira_issue_comment`.`comment_id`,`jira_issue_comment`.`issue_id`,`jira_issue_comment`.`project_id`,`jira_issue_comment`.`user_id`,`jira_issue_comment`.`updated_at`,`jira_issue_comment`.`customer_id`,`jira_issue_comment`.`ref_id` FROM `jira_issue_comment` WHERE `customer_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraIssueComment, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CommentID sql.NullString
		var _IssueID sql.NullString
		var _ProjectID sql.NullString
		var _UserID sql.NullString
		var _UpdatedAt sql.NullInt64
		var _CustomerID sql.NullString
		var _RefID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CommentID,
			&_IssueID,
			&_ProjectID,
			&_UserID,
			&_UpdatedAt,
			&_CustomerID,
			&_RefID,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraIssueComment{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CommentID.Valid {
			t.SetCommentID(_CommentID.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _UpdatedAt.Valid {
			t.SetUpdatedAt(_UpdatedAt.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetRefID will return the JiraIssueComment RefID value
func (t *JiraIssueComment) GetRefID() string {
	return t.RefID
}

// SetRefID will set the JiraIssueComment RefID value
func (t *JiraIssueComment) SetRefID(v string) {
	t.RefID = v
}

// FindJiraIssueCommentsByRefID will find all JiraIssueComments by the RefID value
func FindJiraIssueCommentsByRefID(ctx context.Context, db DB, value string) ([]*JiraIssueComment, error) {
	q := "SELECT `jira_issue_comment`.`id`,`jira_issue_comment`.`checksum`,`jira_issue_comment`.`comment_id`,`jira_issue_comment`.`issue_id`,`jira_issue_comment`.`project_id`,`jira_issue_comment`.`user_id`,`jira_issue_comment`.`updated_at`,`jira_issue_comment`.`customer_id`,`jira_issue_comment`.`ref_id` FROM `jira_issue_comment` WHERE `ref_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraIssueComment, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CommentID sql.NullString
		var _IssueID sql.NullString
		var _ProjectID sql.NullString
		var _UserID sql.NullString
		var _UpdatedAt sql.NullInt64
		var _CustomerID sql.NullString
		var _RefID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CommentID,
			&_IssueID,
			&_ProjectID,
			&_UserID,
			&_UpdatedAt,
			&_CustomerID,
			&_RefID,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraIssueComment{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CommentID.Valid {
			t.SetCommentID(_CommentID.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _UpdatedAt.Valid {
			t.SetUpdatedAt(_UpdatedAt.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindJiraIssueCommentsByRefIDTx will find all JiraIssueComments by the RefID value using the provided transaction
func FindJiraIssueCommentsByRefIDTx(ctx context.Context, tx Tx, value string) ([]*JiraIssueComment, error) {
	q := "SELECT `jira_issue_comment`.`id`,`jira_issue_comment`.`checksum`,`jira_issue_comment`.`comment_id`,`jira_issue_comment`.`issue_id`,`jira_issue_comment`.`project_id`,`jira_issue_comment`.`user_id`,`jira_issue_comment`.`updated_at`,`jira_issue_comment`.`customer_id`,`jira_issue_comment`.`ref_id` FROM `jira_issue_comment` WHERE `ref_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraIssueComment, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CommentID sql.NullString
		var _IssueID sql.NullString
		var _ProjectID sql.NullString
		var _UserID sql.NullString
		var _UpdatedAt sql.NullInt64
		var _CustomerID sql.NullString
		var _RefID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CommentID,
			&_IssueID,
			&_ProjectID,
			&_UserID,
			&_UpdatedAt,
			&_CustomerID,
			&_RefID,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraIssueComment{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CommentID.Valid {
			t.SetCommentID(_CommentID.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _UpdatedAt.Valid {
			t.SetUpdatedAt(_UpdatedAt.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

func (t *JiraIssueComment) toTimestamp(value time.Time) *timestamp.Timestamp {
	ts, _ := ptypes.TimestampProto(value)
	return ts
}

// DBCreateJiraIssueCommentTable will create the JiraIssueComment table
func DBCreateJiraIssueCommentTable(ctx context.Context, db DB) error {
	q := "CREATE TABLE `jira_issue_comment` (`id` VARCHAR(64) NOT NULL PRIMARY KEY,`checksum` CHAR(64),`comment_id`VARCHAR(255) NOT NULL,`issue_id` VARCHAR(64) NOT NULL,`project_id`VARCHAR(64) NOT NULL,`user_id`VARCHAR(64),`updated_at`BIGINT UNSIGNED NOT NULL,`customer_id` VARCHAR(64) NOT NULL,`ref_id` VARCHAR(64) NOT NULL,INDEX jira_issue_comment_comment_id_index (`comment_id`),INDEX jira_issue_comment_issue_id_index (`issue_id`),INDEX jira_issue_comment_project_id_index (`project_id`),INDEX jira_issue_comment_user_id_index (`user_id`),INDEX jira_issue_comment_customer_id_index (`customer_id`),INDEX jira_issue_comment_ref_id_index (`ref_id`),INDEX jira_issue_comment_customer_id_project_id_user_id_updated_at_ind (`customer_id`,`project_id`,`user_id`,`updated_at`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBCreateJiraIssueCommentTableTx will create the JiraIssueComment table using the provided transction
func DBCreateJiraIssueCommentTableTx(ctx context.Context, tx Tx) error {
	q := "CREATE TABLE `jira_issue_comment` (`id` VARCHAR(64) NOT NULL PRIMARY KEY,`checksum` CHAR(64),`comment_id`VARCHAR(255) NOT NULL,`issue_id` VARCHAR(64) NOT NULL,`project_id`VARCHAR(64) NOT NULL,`user_id`VARCHAR(64),`updated_at`BIGINT UNSIGNED NOT NULL,`customer_id` VARCHAR(64) NOT NULL,`ref_id` VARCHAR(64) NOT NULL,INDEX jira_issue_comment_comment_id_index (`comment_id`),INDEX jira_issue_comment_issue_id_index (`issue_id`),INDEX jira_issue_comment_project_id_index (`project_id`),INDEX jira_issue_comment_user_id_index (`user_id`),INDEX jira_issue_comment_customer_id_index (`customer_id`),INDEX jira_issue_comment_ref_id_index (`ref_id`),INDEX jira_issue_comment_customer_id_project_id_user_id_updated_at_ind (`customer_id`,`project_id`,`user_id`,`updated_at`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// DBDropJiraIssueCommentTable will drop the JiraIssueComment table
func DBDropJiraIssueCommentTable(ctx context.Context, db DB) error {
	q := "DROP TABLE IF EXISTS `jira_issue_comment`"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBDropJiraIssueCommentTableTx will drop the JiraIssueComment table using the provided transaction
func DBDropJiraIssueCommentTableTx(ctx context.Context, tx Tx) error {
	q := "DROP TABLE IF EXISTS `jira_issue_comment`"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// CalculateChecksum will calculate a checksum of the SHA1 of all field values
func (t *JiraIssueComment) CalculateChecksum() string {
	return orm.HashStrings(
		orm.ToString(t.ID),
		orm.ToString(t.CommentID),
		orm.ToString(t.IssueID),
		orm.ToString(t.ProjectID),
		orm.ToString(t.UserID),
		orm.ToString(t.UpdatedAt),
		orm.ToString(t.CustomerID),
		orm.ToString(t.RefID),
	)
}

// DBCreate will create a new JiraIssueComment record in the database
func (t *JiraIssueComment) DBCreate(ctx context.Context, db DB) (sql.Result, error) {
	q := "INSERT INTO `jira_issue_comment` (`jira_issue_comment`.`id`,`jira_issue_comment`.`checksum`,`jira_issue_comment`.`comment_id`,`jira_issue_comment`.`issue_id`,`jira_issue_comment`.`project_id`,`jira_issue_comment`.`user_id`,`jira_issue_comment`.`updated_at`,`jira_issue_comment`.`customer_id`,`jira_issue_comment`.`ref_id`) VALUES (?,?,?,?,?,?,?,?,?)"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CommentID),
		orm.ToSQLString(t.IssueID),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLString(t.UserID),
		orm.ToSQLInt64(t.UpdatedAt),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefID),
	)
}

// DBCreateTx will create a new JiraIssueComment record in the database using the provided transaction
func (t *JiraIssueComment) DBCreateTx(ctx context.Context, tx Tx) (sql.Result, error) {
	q := "INSERT INTO `jira_issue_comment` (`jira_issue_comment`.`id`,`jira_issue_comment`.`checksum`,`jira_issue_comment`.`comment_id`,`jira_issue_comment`.`issue_id`,`jira_issue_comment`.`project_id`,`jira_issue_comment`.`user_id`,`jira_issue_comment`.`updated_at`,`jira_issue_comment`.`customer_id`,`jira_issue_comment`.`ref_id`) VALUES (?,?,?,?,?,?,?,?,?)"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CommentID),
		orm.ToSQLString(t.IssueID),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLString(t.UserID),
		orm.ToSQLInt64(t.UpdatedAt),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefID),
	)
}

// DBCreateIgnoreDuplicate will upsert the JiraIssueComment record in the database
func (t *JiraIssueComment) DBCreateIgnoreDuplicate(ctx context.Context, db DB) (sql.Result, error) {
	q := "INSERT INTO `jira_issue_comment` (`jira_issue_comment`.`id`,`jira_issue_comment`.`checksum`,`jira_issue_comment`.`comment_id`,`jira_issue_comment`.`issue_id`,`jira_issue_comment`.`project_id`,`jira_issue_comment`.`user_id`,`jira_issue_comment`.`updated_at`,`jira_issue_comment`.`customer_id`,`jira_issue_comment`.`ref_id`) VALUES (?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CommentID),
		orm.ToSQLString(t.IssueID),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLString(t.UserID),
		orm.ToSQLInt64(t.UpdatedAt),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefID),
	)
}

// DBCreateIgnoreDuplicateTx will upsert the JiraIssueComment record in the database using the provided transaction
func (t *JiraIssueComment) DBCreateIgnoreDuplicateTx(ctx context.Context, tx Tx) (sql.Result, error) {
	q := "INSERT INTO `jira_issue_comment` (`jira_issue_comment`.`id`,`jira_issue_comment`.`checksum`,`jira_issue_comment`.`comment_id`,`jira_issue_comment`.`issue_id`,`jira_issue_comment`.`project_id`,`jira_issue_comment`.`user_id`,`jira_issue_comment`.`updated_at`,`jira_issue_comment`.`customer_id`,`jira_issue_comment`.`ref_id`) VALUES (?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CommentID),
		orm.ToSQLString(t.IssueID),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLString(t.UserID),
		orm.ToSQLInt64(t.UpdatedAt),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefID),
	)
}

// DeleteAllJiraIssueComments deletes all JiraIssueComment records in the database with optional filters
func DeleteAllJiraIssueComments(ctx context.Context, db DB, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(JiraIssueCommentTableName),
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

// DeleteAllJiraIssueCommentsTx deletes all JiraIssueComment records in the database with optional filters using the provided transaction
func DeleteAllJiraIssueCommentsTx(ctx context.Context, tx Tx, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(JiraIssueCommentTableName),
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

// DBDelete will delete this JiraIssueComment record in the database
func (t *JiraIssueComment) DBDelete(ctx context.Context, db DB) (bool, error) {
	q := "DELETE FROM `jira_issue_comment` WHERE `id` = ?"
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

// DBDeleteTx will delete this JiraIssueComment record in the database using the provided transaction
func (t *JiraIssueComment) DBDeleteTx(ctx context.Context, tx Tx) (bool, error) {
	q := "DELETE FROM `jira_issue_comment` WHERE `id` = ?"
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

// DBUpdate will update the JiraIssueComment record in the database
func (t *JiraIssueComment) DBUpdate(ctx context.Context, db DB) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `jira_issue_comment` SET `checksum`=?,`comment_id`=?,`issue_id`=?,`project_id`=?,`user_id`=?,`updated_at`=?,`customer_id`=?,`ref_id`=? WHERE `id`=?"
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CommentID),
		orm.ToSQLString(t.IssueID),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLString(t.UserID),
		orm.ToSQLInt64(t.UpdatedAt),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.ID),
	)
}

// DBUpdateTx will update the JiraIssueComment record in the database using the provided transaction
func (t *JiraIssueComment) DBUpdateTx(ctx context.Context, tx Tx) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `jira_issue_comment` SET `checksum`=?,`comment_id`=?,`issue_id`=?,`project_id`=?,`user_id`=?,`updated_at`=?,`customer_id`=?,`ref_id`=? WHERE `id`=?"
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CommentID),
		orm.ToSQLString(t.IssueID),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLString(t.UserID),
		orm.ToSQLInt64(t.UpdatedAt),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.ID),
	)
}

// DBUpsert will upsert the JiraIssueComment record in the database
func (t *JiraIssueComment) DBUpsert(ctx context.Context, db DB, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `jira_issue_comment` (`jira_issue_comment`.`id`,`jira_issue_comment`.`checksum`,`jira_issue_comment`.`comment_id`,`jira_issue_comment`.`issue_id`,`jira_issue_comment`.`project_id`,`jira_issue_comment`.`user_id`,`jira_issue_comment`.`updated_at`,`jira_issue_comment`.`customer_id`,`jira_issue_comment`.`ref_id`) VALUES (?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `jira_issue_comment` (`jira_issue_comment`.`id`,`jira_issue_comment`.`checksum`,`jira_issue_comment`.`comment_id`,`jira_issue_comment`.`issue_id`,`jira_issue_comment`.`project_id`,`jira_issue_comment`.`user_id`,`jira_issue_comment`.`updated_at`,`jira_issue_comment`.`customer_id`,`jira_issue_comment`.`ref_id`) VALUES (?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`comment_id`=VALUES(`comment_id`),`issue_id`=VALUES(`issue_id`),`project_id`=VALUES(`project_id`),`user_id`=VALUES(`user_id`),`updated_at`=VALUES(`updated_at`),`customer_id`=VALUES(`customer_id`),`ref_id`=VALUES(`ref_id`)"
	}
	r, err := db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CommentID),
		orm.ToSQLString(t.IssueID),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLString(t.UserID),
		orm.ToSQLInt64(t.UpdatedAt),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefID),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBUpsertTx will upsert the JiraIssueComment record in the database using the provided transaction
func (t *JiraIssueComment) DBUpsertTx(ctx context.Context, tx Tx, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `jira_issue_comment` (`jira_issue_comment`.`id`,`jira_issue_comment`.`checksum`,`jira_issue_comment`.`comment_id`,`jira_issue_comment`.`issue_id`,`jira_issue_comment`.`project_id`,`jira_issue_comment`.`user_id`,`jira_issue_comment`.`updated_at`,`jira_issue_comment`.`customer_id`,`jira_issue_comment`.`ref_id`) VALUES (?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `jira_issue_comment` (`jira_issue_comment`.`id`,`jira_issue_comment`.`checksum`,`jira_issue_comment`.`comment_id`,`jira_issue_comment`.`issue_id`,`jira_issue_comment`.`project_id`,`jira_issue_comment`.`user_id`,`jira_issue_comment`.`updated_at`,`jira_issue_comment`.`customer_id`,`jira_issue_comment`.`ref_id`) VALUES (?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`comment_id`=VALUES(`comment_id`),`issue_id`=VALUES(`issue_id`),`project_id`=VALUES(`project_id`),`user_id`=VALUES(`user_id`),`updated_at`=VALUES(`updated_at`),`customer_id`=VALUES(`customer_id`),`ref_id`=VALUES(`ref_id`)"
	}
	r, err := tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CommentID),
		orm.ToSQLString(t.IssueID),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLString(t.UserID),
		orm.ToSQLInt64(t.UpdatedAt),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefID),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBFindOne will find a JiraIssueComment record in the database with the primary key
func (t *JiraIssueComment) DBFindOne(ctx context.Context, db DB, value string) (bool, error) {
	q := "SELECT `jira_issue_comment`.`id`,`jira_issue_comment`.`checksum`,`jira_issue_comment`.`comment_id`,`jira_issue_comment`.`issue_id`,`jira_issue_comment`.`project_id`,`jira_issue_comment`.`user_id`,`jira_issue_comment`.`updated_at`,`jira_issue_comment`.`customer_id`,`jira_issue_comment`.`ref_id` FROM `jira_issue_comment` WHERE `id` = ? LIMIT 1"
	row := db.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CommentID sql.NullString
	var _IssueID sql.NullString
	var _ProjectID sql.NullString
	var _UserID sql.NullString
	var _UpdatedAt sql.NullInt64
	var _CustomerID sql.NullString
	var _RefID sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CommentID,
		&_IssueID,
		&_ProjectID,
		&_UserID,
		&_UpdatedAt,
		&_CustomerID,
		&_RefID,
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
	if _CommentID.Valid {
		t.SetCommentID(_CommentID.String)
	}
	if _IssueID.Valid {
		t.SetIssueID(_IssueID.String)
	}
	if _ProjectID.Valid {
		t.SetProjectID(_ProjectID.String)
	}
	if _UserID.Valid {
		t.SetUserID(_UserID.String)
	}
	if _UpdatedAt.Valid {
		t.SetUpdatedAt(_UpdatedAt.Int64)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _RefID.Valid {
		t.SetRefID(_RefID.String)
	}
	return true, nil
}

// DBFindOneTx will find a JiraIssueComment record in the database with the primary key using the provided transaction
func (t *JiraIssueComment) DBFindOneTx(ctx context.Context, tx Tx, value string) (bool, error) {
	q := "SELECT `jira_issue_comment`.`id`,`jira_issue_comment`.`checksum`,`jira_issue_comment`.`comment_id`,`jira_issue_comment`.`issue_id`,`jira_issue_comment`.`project_id`,`jira_issue_comment`.`user_id`,`jira_issue_comment`.`updated_at`,`jira_issue_comment`.`customer_id`,`jira_issue_comment`.`ref_id` FROM `jira_issue_comment` WHERE `id` = ? LIMIT 1"
	row := tx.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CommentID sql.NullString
	var _IssueID sql.NullString
	var _ProjectID sql.NullString
	var _UserID sql.NullString
	var _UpdatedAt sql.NullInt64
	var _CustomerID sql.NullString
	var _RefID sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CommentID,
		&_IssueID,
		&_ProjectID,
		&_UserID,
		&_UpdatedAt,
		&_CustomerID,
		&_RefID,
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
	if _CommentID.Valid {
		t.SetCommentID(_CommentID.String)
	}
	if _IssueID.Valid {
		t.SetIssueID(_IssueID.String)
	}
	if _ProjectID.Valid {
		t.SetProjectID(_ProjectID.String)
	}
	if _UserID.Valid {
		t.SetUserID(_UserID.String)
	}
	if _UpdatedAt.Valid {
		t.SetUpdatedAt(_UpdatedAt.Int64)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _RefID.Valid {
		t.SetRefID(_RefID.String)
	}
	return true, nil
}

// FindJiraIssueComments will find a JiraIssueComment record in the database with the provided parameters
func FindJiraIssueComments(ctx context.Context, db DB, _params ...interface{}) ([]*JiraIssueComment, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("comment_id"),
		orm.Column("issue_id"),
		orm.Column("project_id"),
		orm.Column("user_id"),
		orm.Column("updated_at"),
		orm.Column("customer_id"),
		orm.Column("ref_id"),
		orm.Table(JiraIssueCommentTableName),
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
	results := make([]*JiraIssueComment, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CommentID sql.NullString
		var _IssueID sql.NullString
		var _ProjectID sql.NullString
		var _UserID sql.NullString
		var _UpdatedAt sql.NullInt64
		var _CustomerID sql.NullString
		var _RefID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CommentID,
			&_IssueID,
			&_ProjectID,
			&_UserID,
			&_UpdatedAt,
			&_CustomerID,
			&_RefID,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraIssueComment{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CommentID.Valid {
			t.SetCommentID(_CommentID.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _UpdatedAt.Valid {
			t.SetUpdatedAt(_UpdatedAt.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindJiraIssueCommentsTx will find a JiraIssueComment record in the database with the provided parameters using the provided transaction
func FindJiraIssueCommentsTx(ctx context.Context, tx Tx, _params ...interface{}) ([]*JiraIssueComment, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("comment_id"),
		orm.Column("issue_id"),
		orm.Column("project_id"),
		orm.Column("user_id"),
		orm.Column("updated_at"),
		orm.Column("customer_id"),
		orm.Column("ref_id"),
		orm.Table(JiraIssueCommentTableName),
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
	results := make([]*JiraIssueComment, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CommentID sql.NullString
		var _IssueID sql.NullString
		var _ProjectID sql.NullString
		var _UserID sql.NullString
		var _UpdatedAt sql.NullInt64
		var _CustomerID sql.NullString
		var _RefID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CommentID,
			&_IssueID,
			&_ProjectID,
			&_UserID,
			&_UpdatedAt,
			&_CustomerID,
			&_RefID,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraIssueComment{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CommentID.Valid {
			t.SetCommentID(_CommentID.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _UpdatedAt.Valid {
			t.SetUpdatedAt(_UpdatedAt.Int64)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _RefID.Valid {
			t.SetRefID(_RefID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// DBFind will find a JiraIssueComment record in the database with the provided parameters
func (t *JiraIssueComment) DBFind(ctx context.Context, db DB, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("comment_id"),
		orm.Column("issue_id"),
		orm.Column("project_id"),
		orm.Column("user_id"),
		orm.Column("updated_at"),
		orm.Column("customer_id"),
		orm.Column("ref_id"),
		orm.Table(JiraIssueCommentTableName),
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
	var _CommentID sql.NullString
	var _IssueID sql.NullString
	var _ProjectID sql.NullString
	var _UserID sql.NullString
	var _UpdatedAt sql.NullInt64
	var _CustomerID sql.NullString
	var _RefID sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CommentID,
		&_IssueID,
		&_ProjectID,
		&_UserID,
		&_UpdatedAt,
		&_CustomerID,
		&_RefID,
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
	if _CommentID.Valid {
		t.SetCommentID(_CommentID.String)
	}
	if _IssueID.Valid {
		t.SetIssueID(_IssueID.String)
	}
	if _ProjectID.Valid {
		t.SetProjectID(_ProjectID.String)
	}
	if _UserID.Valid {
		t.SetUserID(_UserID.String)
	}
	if _UpdatedAt.Valid {
		t.SetUpdatedAt(_UpdatedAt.Int64)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _RefID.Valid {
		t.SetRefID(_RefID.String)
	}
	return true, nil
}

// DBFindTx will find a JiraIssueComment record in the database with the provided parameters using the provided transaction
func (t *JiraIssueComment) DBFindTx(ctx context.Context, tx Tx, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("comment_id"),
		orm.Column("issue_id"),
		orm.Column("project_id"),
		orm.Column("user_id"),
		orm.Column("updated_at"),
		orm.Column("customer_id"),
		orm.Column("ref_id"),
		orm.Table(JiraIssueCommentTableName),
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
	var _CommentID sql.NullString
	var _IssueID sql.NullString
	var _ProjectID sql.NullString
	var _UserID sql.NullString
	var _UpdatedAt sql.NullInt64
	var _CustomerID sql.NullString
	var _RefID sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CommentID,
		&_IssueID,
		&_ProjectID,
		&_UserID,
		&_UpdatedAt,
		&_CustomerID,
		&_RefID,
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
	if _CommentID.Valid {
		t.SetCommentID(_CommentID.String)
	}
	if _IssueID.Valid {
		t.SetIssueID(_IssueID.String)
	}
	if _ProjectID.Valid {
		t.SetProjectID(_ProjectID.String)
	}
	if _UserID.Valid {
		t.SetUserID(_UserID.String)
	}
	if _UpdatedAt.Valid {
		t.SetUpdatedAt(_UpdatedAt.Int64)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _RefID.Valid {
		t.SetRefID(_RefID.String)
	}
	return true, nil
}

// CountJiraIssueComments will find the count of JiraIssueComment records in the database
func CountJiraIssueComments(ctx context.Context, db DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(JiraIssueCommentTableName),
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

// CountJiraIssueCommentsTx will find the count of JiraIssueComment records in the database using the provided transaction
func CountJiraIssueCommentsTx(ctx context.Context, tx Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(JiraIssueCommentTableName),
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

// DBCount will find the count of JiraIssueComment records in the database
func (t *JiraIssueComment) DBCount(ctx context.Context, db DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(JiraIssueCommentTableName),
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

// DBCountTx will find the count of JiraIssueComment records in the database using the provided transaction
func (t *JiraIssueComment) DBCountTx(ctx context.Context, tx Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(JiraIssueCommentTableName),
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

// DBExists will return true if the JiraIssueComment record exists in the database
func (t *JiraIssueComment) DBExists(ctx context.Context, db DB) (bool, error) {
	q := "SELECT `id` FROM `jira_issue_comment` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := db.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// DBExistsTx will return true if the JiraIssueComment record exists in the database using the provided transaction
func (t *JiraIssueComment) DBExistsTx(ctx context.Context, tx Tx) (bool, error) {
	q := "SELECT `id` FROM `jira_issue_comment` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := tx.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// PrimaryKeyColumn returns the column name for the primary key
func (t *JiraIssueComment) PrimaryKeyColumn() string {
	return JiraIssueCommentColumnID
}

// PrimaryKeyColumnType returns the primary key column Go type as a string
func (t *JiraIssueComment) PrimaryKeyColumnType() string {
	return "string"
}

// PrimaryKey returns the primary key column value
func (t *JiraIssueComment) PrimaryKey() interface{} {
	return t.ID
}
