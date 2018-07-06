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

var _ Model = (*IssueComment)(nil)
var _ CSVWriter = (*IssueComment)(nil)
var _ JSONWriter = (*IssueComment)(nil)
var _ Checksum = (*IssueComment)(nil)

// IssueCommentTableName is the name of the table in SQL
const IssueCommentTableName = "issue_comment"

var IssueCommentColumns = []string{
	"id",
	"checksum",
	"issue_id",
	"user_id",
	"project_id",
	"body",
	"deleted",
	"deleted_user_id",
	"created_at",
	"updated_at",
	"deleted_at",
	"url",
	"customer_id",
	"ref_type",
	"ref_id",
	"metadata",
}

// IssueComment table
type IssueComment struct {
	Body          string  `json:"body"`
	Checksum      *string `json:"checksum,omitempty"`
	CreatedAt     int64   `json:"created_at"`
	CustomerID    string  `json:"customer_id"`
	Deleted       bool    `json:"deleted"`
	DeletedAt     *int64  `json:"deleted_at,omitempty"`
	DeletedUserID *string `json:"deleted_user_id,omitempty"`
	ID            string  `json:"id"`
	IssueID       string  `json:"issue_id"`
	Metadata      *string `json:"metadata,omitempty"`
	ProjectID     string  `json:"project_id"`
	RefID         string  `json:"ref_id"`
	RefType       string  `json:"ref_type"`
	UpdatedAt     *int64  `json:"updated_at,omitempty"`
	URL           string  `json:"url"`
	UserID        *string `json:"user_id,omitempty"`
}

// TableName returns the SQL table name for IssueComment and satifies the Model interface
func (t *IssueComment) TableName() string {
	return IssueCommentTableName
}

// ToCSV will serialize the IssueComment instance to a CSV compatible array of strings
func (t *IssueComment) ToCSV() []string {
	return []string{
		t.ID,
		t.CalculateChecksum(),
		t.IssueID,
		toCSVString(t.UserID),
		t.ProjectID,
		t.Body,
		toCSVBool(t.Deleted),
		toCSVString(t.DeletedUserID),
		toCSVString(t.CreatedAt),
		toCSVString(t.UpdatedAt),
		toCSVString(t.DeletedAt),
		t.URL,
		t.CustomerID,
		t.RefType,
		t.RefID,
		toCSVString(t.Metadata),
	}
}

// WriteCSV will serialize the IssueComment instance to the writer as CSV and satisfies the CSVWriter interface
func (t *IssueComment) WriteCSV(w *csv.Writer) error {
	return w.Write(t.ToCSV())
}

// WriteJSON will serialize the IssueComment instance to the writer as JSON and satisfies the JSONWriter interface
func (t *IssueComment) WriteJSON(w io.Writer, indent ...bool) error {
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

// NewIssueCommentReader creates a JSON reader which can read in IssueComment objects serialized as JSON either as an array, single object or json new lines
// and writes each IssueComment to the channel provided
func NewIssueCommentReader(r io.Reader, ch chan<- IssueComment) error {
	return orm.Deserialize(r, func(buf json.RawMessage) error {
		dec := json.NewDecoder(bytes.NewBuffer(buf))
		e := IssueComment{}
		if err := dec.Decode(&e); err != nil {
			return err
		}
		ch <- e
		return nil
	})
}

// NewCSVIssueCommentReaderDir will read the reader as CSV and emit each record to the channel provided
func NewCSVIssueCommentReader(r io.Reader, ch chan<- IssueComment) error {
	cr := csv.NewReader(r)
	for {
		record, err := cr.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		ch <- IssueComment{
			ID:            record[0],
			Checksum:      fromStringPointer(record[1]),
			IssueID:       record[2],
			UserID:        fromStringPointer(record[3]),
			ProjectID:     record[4],
			Body:          record[5],
			Deleted:       fromCSVBool(record[6]),
			DeletedUserID: fromStringPointer(record[7]),
			CreatedAt:     fromCSVInt64(record[8]),
			UpdatedAt:     fromCSVInt64Pointer(record[9]),
			DeletedAt:     fromCSVInt64Pointer(record[10]),
			URL:           record[11],
			CustomerID:    record[12],
			RefType:       record[13],
			RefID:         record[14],
			Metadata:      fromStringPointer(record[15]),
		}
	}
	return nil
}

// NewCSVIssueCommentReaderFile will read the file as a CSV and emit each record to the channel provided
func NewCSVIssueCommentReaderFile(fp string, ch chan<- IssueComment) error {
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
	return NewCSVIssueCommentReader(fc, ch)
}

// NewCSVIssueCommentReaderDir will read the issue_comment.csv.gz file as a CSV and emit each record to the channel provided
func NewCSVIssueCommentReaderDir(dir string, ch chan<- IssueComment) error {
	return NewCSVIssueCommentReaderFile(filepath.Join(dir, "issue_comment.csv.gz"), ch)
}

// IssueCommentCSVDeduper is a function callback which takes the existing value (a) and the new value (b)
// and the return value should be the one to use (or a new one, if applicable). return nil
// to skip processing of this record
type IssueCommentCSVDeduper func(a IssueComment, b IssueComment) *IssueComment

// IssueCommentCSVDedupeDisabled is set on whether the CSV writer should de-dupe values by key
var IssueCommentCSVDedupeDisabled bool

// NewIssueCommentCSVWriterSize creates a batch writer that will write each IssueComment into a CSV file
// this method will automatically de-duplicate entries using the primary key. if the checksum
// for a newer item with the same primary key doesn't match a previously sent item, the newer
// one will be used
func NewIssueCommentCSVWriterSize(w io.Writer, size int, dedupers ...IssueCommentCSVDeduper) (chan IssueComment, chan bool, error) {
	cw := csv.NewWriter(w)
	ch := make(chan IssueComment, size)
	done := make(chan bool)
	go func() {
		defer func() { done <- true }()
		dodedupe := !IssueCommentCSVDedupeDisabled
		var kv map[string]*IssueComment
		var deduper IssueCommentCSVDeduper
		if dedupers != nil && len(dedupers) > 0 {
			deduper = dedupers[0]
			dodedupe = true
		}
		if dodedupe {
			kv = make(map[string]*IssueComment)
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

// IssueCommentCSVDefaultSize is the default channel buffer size if not provided
var IssueCommentCSVDefaultSize = 100

// NewIssueCommentCSVWriter creates a batch writer that will write each IssueComment into a CSV file
func NewIssueCommentCSVWriter(w io.Writer, dedupers ...IssueCommentCSVDeduper) (chan IssueComment, chan bool, error) {
	return NewIssueCommentCSVWriterSize(w, IssueCommentCSVDefaultSize, dedupers...)
}

// NewIssueCommentCSVWriterDir creates a batch writer that will write each IssueComment into a CSV file named issue_comment.csv.gz in dir
func NewIssueCommentCSVWriterDir(dir string, dedupers ...IssueCommentCSVDeduper) (chan IssueComment, chan bool, error) {
	return NewIssueCommentCSVWriterFile(filepath.Join(dir, "issue_comment.csv.gz"), dedupers...)
}

// NewIssueCommentCSVWriterFile creates a batch writer that will write each IssueComment into a CSV file
func NewIssueCommentCSVWriterFile(fn string, dedupers ...IssueCommentCSVDeduper) (chan IssueComment, chan bool, error) {
	f, err := os.Create(fn)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening CSV file %s. %v", fn, err)
	}
	var fc io.WriteCloser = f
	if filepath.Ext(fn) == ".gz" {
		w, _ := gzip.NewWriterLevel(f, gzip.BestCompression)
		fc = w
	}
	ch, done, err := NewIssueCommentCSVWriter(fc, dedupers...)
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

type IssueCommentDBAction func(ctx context.Context, db *sql.DB, record IssueComment) error

// NewIssueCommentDBWriterSize creates a DB writer that will write each issue into the DB
func NewIssueCommentDBWriterSize(ctx context.Context, db *sql.DB, errors chan<- error, size int, actions ...IssueCommentDBAction) (chan IssueComment, chan bool, error) {
	ch := make(chan IssueComment, size)
	done := make(chan bool)
	var action IssueCommentDBAction
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

// NewIssueCommentDBWriter creates a DB writer that will write each issue into the DB
func NewIssueCommentDBWriter(ctx context.Context, db *sql.DB, errors chan<- error, actions ...IssueCommentDBAction) (chan IssueComment, chan bool, error) {
	return NewIssueCommentDBWriterSize(ctx, db, errors, 100, actions...)
}

// IssueCommentColumnID is the ID SQL column name for the IssueComment table
const IssueCommentColumnID = "id"

// IssueCommentEscapedColumnID is the escaped ID SQL column name for the IssueComment table
const IssueCommentEscapedColumnID = "`id`"

// IssueCommentColumnChecksum is the Checksum SQL column name for the IssueComment table
const IssueCommentColumnChecksum = "checksum"

// IssueCommentEscapedColumnChecksum is the escaped Checksum SQL column name for the IssueComment table
const IssueCommentEscapedColumnChecksum = "`checksum`"

// IssueCommentColumnIssueID is the IssueID SQL column name for the IssueComment table
const IssueCommentColumnIssueID = "issue_id"

// IssueCommentEscapedColumnIssueID is the escaped IssueID SQL column name for the IssueComment table
const IssueCommentEscapedColumnIssueID = "`issue_id`"

// IssueCommentColumnUserID is the UserID SQL column name for the IssueComment table
const IssueCommentColumnUserID = "user_id"

// IssueCommentEscapedColumnUserID is the escaped UserID SQL column name for the IssueComment table
const IssueCommentEscapedColumnUserID = "`user_id`"

// IssueCommentColumnProjectID is the ProjectID SQL column name for the IssueComment table
const IssueCommentColumnProjectID = "project_id"

// IssueCommentEscapedColumnProjectID is the escaped ProjectID SQL column name for the IssueComment table
const IssueCommentEscapedColumnProjectID = "`project_id`"

// IssueCommentColumnBody is the Body SQL column name for the IssueComment table
const IssueCommentColumnBody = "body"

// IssueCommentEscapedColumnBody is the escaped Body SQL column name for the IssueComment table
const IssueCommentEscapedColumnBody = "`body`"

// IssueCommentColumnDeleted is the Deleted SQL column name for the IssueComment table
const IssueCommentColumnDeleted = "deleted"

// IssueCommentEscapedColumnDeleted is the escaped Deleted SQL column name for the IssueComment table
const IssueCommentEscapedColumnDeleted = "`deleted`"

// IssueCommentColumnDeletedUserID is the DeletedUserID SQL column name for the IssueComment table
const IssueCommentColumnDeletedUserID = "deleted_user_id"

// IssueCommentEscapedColumnDeletedUserID is the escaped DeletedUserID SQL column name for the IssueComment table
const IssueCommentEscapedColumnDeletedUserID = "`deleted_user_id`"

// IssueCommentColumnCreatedAt is the CreatedAt SQL column name for the IssueComment table
const IssueCommentColumnCreatedAt = "created_at"

// IssueCommentEscapedColumnCreatedAt is the escaped CreatedAt SQL column name for the IssueComment table
const IssueCommentEscapedColumnCreatedAt = "`created_at`"

// IssueCommentColumnUpdatedAt is the UpdatedAt SQL column name for the IssueComment table
const IssueCommentColumnUpdatedAt = "updated_at"

// IssueCommentEscapedColumnUpdatedAt is the escaped UpdatedAt SQL column name for the IssueComment table
const IssueCommentEscapedColumnUpdatedAt = "`updated_at`"

// IssueCommentColumnDeletedAt is the DeletedAt SQL column name for the IssueComment table
const IssueCommentColumnDeletedAt = "deleted_at"

// IssueCommentEscapedColumnDeletedAt is the escaped DeletedAt SQL column name for the IssueComment table
const IssueCommentEscapedColumnDeletedAt = "`deleted_at`"

// IssueCommentColumnURL is the URL SQL column name for the IssueComment table
const IssueCommentColumnURL = "url"

// IssueCommentEscapedColumnURL is the escaped URL SQL column name for the IssueComment table
const IssueCommentEscapedColumnURL = "`url`"

// IssueCommentColumnCustomerID is the CustomerID SQL column name for the IssueComment table
const IssueCommentColumnCustomerID = "customer_id"

// IssueCommentEscapedColumnCustomerID is the escaped CustomerID SQL column name for the IssueComment table
const IssueCommentEscapedColumnCustomerID = "`customer_id`"

// IssueCommentColumnRefType is the RefType SQL column name for the IssueComment table
const IssueCommentColumnRefType = "ref_type"

// IssueCommentEscapedColumnRefType is the escaped RefType SQL column name for the IssueComment table
const IssueCommentEscapedColumnRefType = "`ref_type`"

// IssueCommentColumnRefID is the RefID SQL column name for the IssueComment table
const IssueCommentColumnRefID = "ref_id"

// IssueCommentEscapedColumnRefID is the escaped RefID SQL column name for the IssueComment table
const IssueCommentEscapedColumnRefID = "`ref_id`"

// IssueCommentColumnMetadata is the Metadata SQL column name for the IssueComment table
const IssueCommentColumnMetadata = "metadata"

// IssueCommentEscapedColumnMetadata is the escaped Metadata SQL column name for the IssueComment table
const IssueCommentEscapedColumnMetadata = "`metadata`"

// GetID will return the IssueComment ID value
func (t *IssueComment) GetID() string {
	return t.ID
}

// SetID will set the IssueComment ID value
func (t *IssueComment) SetID(v string) {
	t.ID = v
}

// FindIssueCommentByID will find a IssueComment by ID
func FindIssueCommentByID(ctx context.Context, db *sql.DB, value string) (*IssueComment, error) {
	q := "SELECT `issue_comment`.`id`,`issue_comment`.`checksum`,`issue_comment`.`issue_id`,`issue_comment`.`user_id`,`issue_comment`.`project_id`,`issue_comment`.`body`,`issue_comment`.`deleted`,`issue_comment`.`deleted_user_id`,`issue_comment`.`created_at`,`issue_comment`.`updated_at`,`issue_comment`.`deleted_at`,`issue_comment`.`url`,`issue_comment`.`customer_id`,`issue_comment`.`ref_type`,`issue_comment`.`ref_id`,`issue_comment`.`metadata` FROM `issue_comment` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _IssueID sql.NullString
	var _UserID sql.NullString
	var _ProjectID sql.NullString
	var _Body sql.NullString
	var _Deleted sql.NullBool
	var _DeletedUserID sql.NullString
	var _CreatedAt sql.NullInt64
	var _UpdatedAt sql.NullInt64
	var _DeletedAt sql.NullInt64
	var _URL sql.NullString
	var _CustomerID sql.NullString
	var _RefType sql.NullString
	var _RefID sql.NullString
	var _Metadata sql.NullString
	err := db.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_IssueID,
		&_UserID,
		&_ProjectID,
		&_Body,
		&_Deleted,
		&_DeletedUserID,
		&_CreatedAt,
		&_UpdatedAt,
		&_DeletedAt,
		&_URL,
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
	t := &IssueComment{}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _Checksum.Valid {
		t.SetChecksum(_Checksum.String)
	}
	if _IssueID.Valid {
		t.SetIssueID(_IssueID.String)
	}
	if _UserID.Valid {
		t.SetUserID(_UserID.String)
	}
	if _ProjectID.Valid {
		t.SetProjectID(_ProjectID.String)
	}
	if _Body.Valid {
		t.SetBody(_Body.String)
	}
	if _Deleted.Valid {
		t.SetDeleted(_Deleted.Bool)
	}
	if _DeletedUserID.Valid {
		t.SetDeletedUserID(_DeletedUserID.String)
	}
	if _CreatedAt.Valid {
		t.SetCreatedAt(_CreatedAt.Int64)
	}
	if _UpdatedAt.Valid {
		t.SetUpdatedAt(_UpdatedAt.Int64)
	}
	if _DeletedAt.Valid {
		t.SetDeletedAt(_DeletedAt.Int64)
	}
	if _URL.Valid {
		t.SetURL(_URL.String)
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

// FindIssueCommentByIDTx will find a IssueComment by ID using the provided transaction
func FindIssueCommentByIDTx(ctx context.Context, tx *sql.Tx, value string) (*IssueComment, error) {
	q := "SELECT `issue_comment`.`id`,`issue_comment`.`checksum`,`issue_comment`.`issue_id`,`issue_comment`.`user_id`,`issue_comment`.`project_id`,`issue_comment`.`body`,`issue_comment`.`deleted`,`issue_comment`.`deleted_user_id`,`issue_comment`.`created_at`,`issue_comment`.`updated_at`,`issue_comment`.`deleted_at`,`issue_comment`.`url`,`issue_comment`.`customer_id`,`issue_comment`.`ref_type`,`issue_comment`.`ref_id`,`issue_comment`.`metadata` FROM `issue_comment` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _IssueID sql.NullString
	var _UserID sql.NullString
	var _ProjectID sql.NullString
	var _Body sql.NullString
	var _Deleted sql.NullBool
	var _DeletedUserID sql.NullString
	var _CreatedAt sql.NullInt64
	var _UpdatedAt sql.NullInt64
	var _DeletedAt sql.NullInt64
	var _URL sql.NullString
	var _CustomerID sql.NullString
	var _RefType sql.NullString
	var _RefID sql.NullString
	var _Metadata sql.NullString
	err := tx.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_IssueID,
		&_UserID,
		&_ProjectID,
		&_Body,
		&_Deleted,
		&_DeletedUserID,
		&_CreatedAt,
		&_UpdatedAt,
		&_DeletedAt,
		&_URL,
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
	t := &IssueComment{}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _Checksum.Valid {
		t.SetChecksum(_Checksum.String)
	}
	if _IssueID.Valid {
		t.SetIssueID(_IssueID.String)
	}
	if _UserID.Valid {
		t.SetUserID(_UserID.String)
	}
	if _ProjectID.Valid {
		t.SetProjectID(_ProjectID.String)
	}
	if _Body.Valid {
		t.SetBody(_Body.String)
	}
	if _Deleted.Valid {
		t.SetDeleted(_Deleted.Bool)
	}
	if _DeletedUserID.Valid {
		t.SetDeletedUserID(_DeletedUserID.String)
	}
	if _CreatedAt.Valid {
		t.SetCreatedAt(_CreatedAt.Int64)
	}
	if _UpdatedAt.Valid {
		t.SetUpdatedAt(_UpdatedAt.Int64)
	}
	if _DeletedAt.Valid {
		t.SetDeletedAt(_DeletedAt.Int64)
	}
	if _URL.Valid {
		t.SetURL(_URL.String)
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

// GetChecksum will return the IssueComment Checksum value
func (t *IssueComment) GetChecksum() string {
	if t.Checksum == nil {
		return ""
	}
	return *t.Checksum
}

// SetChecksum will set the IssueComment Checksum value
func (t *IssueComment) SetChecksum(v string) {
	t.Checksum = &v
}

// GetIssueID will return the IssueComment IssueID value
func (t *IssueComment) GetIssueID() string {
	return t.IssueID
}

// SetIssueID will set the IssueComment IssueID value
func (t *IssueComment) SetIssueID(v string) {
	t.IssueID = v
}

// FindIssueCommentsByIssueID will find all IssueComments by the IssueID value
func FindIssueCommentsByIssueID(ctx context.Context, db *sql.DB, value string) ([]*IssueComment, error) {
	q := "SELECT `issue_comment`.`id`,`issue_comment`.`checksum`,`issue_comment`.`issue_id`,`issue_comment`.`user_id`,`issue_comment`.`project_id`,`issue_comment`.`body`,`issue_comment`.`deleted`,`issue_comment`.`deleted_user_id`,`issue_comment`.`created_at`,`issue_comment`.`updated_at`,`issue_comment`.`deleted_at`,`issue_comment`.`url`,`issue_comment`.`customer_id`,`issue_comment`.`ref_type`,`issue_comment`.`ref_id`,`issue_comment`.`metadata` FROM `issue_comment` WHERE `issue_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*IssueComment, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _IssueID sql.NullString
		var _UserID sql.NullString
		var _ProjectID sql.NullString
		var _Body sql.NullString
		var _Deleted sql.NullBool
		var _DeletedUserID sql.NullString
		var _CreatedAt sql.NullInt64
		var _UpdatedAt sql.NullInt64
		var _DeletedAt sql.NullInt64
		var _URL sql.NullString
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefID sql.NullString
		var _Metadata sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_IssueID,
			&_UserID,
			&_ProjectID,
			&_Body,
			&_Deleted,
			&_DeletedUserID,
			&_CreatedAt,
			&_UpdatedAt,
			&_DeletedAt,
			&_URL,
			&_CustomerID,
			&_RefType,
			&_RefID,
			&_Metadata,
		)
		if err != nil {
			return nil, err
		}
		t := &IssueComment{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _Body.Valid {
			t.SetBody(_Body.String)
		}
		if _Deleted.Valid {
			t.SetDeleted(_Deleted.Bool)
		}
		if _DeletedUserID.Valid {
			t.SetDeletedUserID(_DeletedUserID.String)
		}
		if _CreatedAt.Valid {
			t.SetCreatedAt(_CreatedAt.Int64)
		}
		if _UpdatedAt.Valid {
			t.SetUpdatedAt(_UpdatedAt.Int64)
		}
		if _DeletedAt.Valid {
			t.SetDeletedAt(_DeletedAt.Int64)
		}
		if _URL.Valid {
			t.SetURL(_URL.String)
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

// FindIssueCommentsByIssueIDTx will find all IssueComments by the IssueID value using the provided transaction
func FindIssueCommentsByIssueIDTx(ctx context.Context, tx *sql.Tx, value string) ([]*IssueComment, error) {
	q := "SELECT `issue_comment`.`id`,`issue_comment`.`checksum`,`issue_comment`.`issue_id`,`issue_comment`.`user_id`,`issue_comment`.`project_id`,`issue_comment`.`body`,`issue_comment`.`deleted`,`issue_comment`.`deleted_user_id`,`issue_comment`.`created_at`,`issue_comment`.`updated_at`,`issue_comment`.`deleted_at`,`issue_comment`.`url`,`issue_comment`.`customer_id`,`issue_comment`.`ref_type`,`issue_comment`.`ref_id`,`issue_comment`.`metadata` FROM `issue_comment` WHERE `issue_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*IssueComment, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _IssueID sql.NullString
		var _UserID sql.NullString
		var _ProjectID sql.NullString
		var _Body sql.NullString
		var _Deleted sql.NullBool
		var _DeletedUserID sql.NullString
		var _CreatedAt sql.NullInt64
		var _UpdatedAt sql.NullInt64
		var _DeletedAt sql.NullInt64
		var _URL sql.NullString
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefID sql.NullString
		var _Metadata sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_IssueID,
			&_UserID,
			&_ProjectID,
			&_Body,
			&_Deleted,
			&_DeletedUserID,
			&_CreatedAt,
			&_UpdatedAt,
			&_DeletedAt,
			&_URL,
			&_CustomerID,
			&_RefType,
			&_RefID,
			&_Metadata,
		)
		if err != nil {
			return nil, err
		}
		t := &IssueComment{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _Body.Valid {
			t.SetBody(_Body.String)
		}
		if _Deleted.Valid {
			t.SetDeleted(_Deleted.Bool)
		}
		if _DeletedUserID.Valid {
			t.SetDeletedUserID(_DeletedUserID.String)
		}
		if _CreatedAt.Valid {
			t.SetCreatedAt(_CreatedAt.Int64)
		}
		if _UpdatedAt.Valid {
			t.SetUpdatedAt(_UpdatedAt.Int64)
		}
		if _DeletedAt.Valid {
			t.SetDeletedAt(_DeletedAt.Int64)
		}
		if _URL.Valid {
			t.SetURL(_URL.String)
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

// GetUserID will return the IssueComment UserID value
func (t *IssueComment) GetUserID() string {
	if t.UserID == nil {
		return ""
	}
	return *t.UserID
}

// SetUserID will set the IssueComment UserID value
func (t *IssueComment) SetUserID(v string) {
	t.UserID = &v
}

// FindIssueCommentsByUserID will find all IssueComments by the UserID value
func FindIssueCommentsByUserID(ctx context.Context, db *sql.DB, value string) ([]*IssueComment, error) {
	q := "SELECT `issue_comment`.`id`,`issue_comment`.`checksum`,`issue_comment`.`issue_id`,`issue_comment`.`user_id`,`issue_comment`.`project_id`,`issue_comment`.`body`,`issue_comment`.`deleted`,`issue_comment`.`deleted_user_id`,`issue_comment`.`created_at`,`issue_comment`.`updated_at`,`issue_comment`.`deleted_at`,`issue_comment`.`url`,`issue_comment`.`customer_id`,`issue_comment`.`ref_type`,`issue_comment`.`ref_id`,`issue_comment`.`metadata` FROM `issue_comment` WHERE `user_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*IssueComment, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _IssueID sql.NullString
		var _UserID sql.NullString
		var _ProjectID sql.NullString
		var _Body sql.NullString
		var _Deleted sql.NullBool
		var _DeletedUserID sql.NullString
		var _CreatedAt sql.NullInt64
		var _UpdatedAt sql.NullInt64
		var _DeletedAt sql.NullInt64
		var _URL sql.NullString
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefID sql.NullString
		var _Metadata sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_IssueID,
			&_UserID,
			&_ProjectID,
			&_Body,
			&_Deleted,
			&_DeletedUserID,
			&_CreatedAt,
			&_UpdatedAt,
			&_DeletedAt,
			&_URL,
			&_CustomerID,
			&_RefType,
			&_RefID,
			&_Metadata,
		)
		if err != nil {
			return nil, err
		}
		t := &IssueComment{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _Body.Valid {
			t.SetBody(_Body.String)
		}
		if _Deleted.Valid {
			t.SetDeleted(_Deleted.Bool)
		}
		if _DeletedUserID.Valid {
			t.SetDeletedUserID(_DeletedUserID.String)
		}
		if _CreatedAt.Valid {
			t.SetCreatedAt(_CreatedAt.Int64)
		}
		if _UpdatedAt.Valid {
			t.SetUpdatedAt(_UpdatedAt.Int64)
		}
		if _DeletedAt.Valid {
			t.SetDeletedAt(_DeletedAt.Int64)
		}
		if _URL.Valid {
			t.SetURL(_URL.String)
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

// FindIssueCommentsByUserIDTx will find all IssueComments by the UserID value using the provided transaction
func FindIssueCommentsByUserIDTx(ctx context.Context, tx *sql.Tx, value string) ([]*IssueComment, error) {
	q := "SELECT `issue_comment`.`id`,`issue_comment`.`checksum`,`issue_comment`.`issue_id`,`issue_comment`.`user_id`,`issue_comment`.`project_id`,`issue_comment`.`body`,`issue_comment`.`deleted`,`issue_comment`.`deleted_user_id`,`issue_comment`.`created_at`,`issue_comment`.`updated_at`,`issue_comment`.`deleted_at`,`issue_comment`.`url`,`issue_comment`.`customer_id`,`issue_comment`.`ref_type`,`issue_comment`.`ref_id`,`issue_comment`.`metadata` FROM `issue_comment` WHERE `user_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*IssueComment, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _IssueID sql.NullString
		var _UserID sql.NullString
		var _ProjectID sql.NullString
		var _Body sql.NullString
		var _Deleted sql.NullBool
		var _DeletedUserID sql.NullString
		var _CreatedAt sql.NullInt64
		var _UpdatedAt sql.NullInt64
		var _DeletedAt sql.NullInt64
		var _URL sql.NullString
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefID sql.NullString
		var _Metadata sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_IssueID,
			&_UserID,
			&_ProjectID,
			&_Body,
			&_Deleted,
			&_DeletedUserID,
			&_CreatedAt,
			&_UpdatedAt,
			&_DeletedAt,
			&_URL,
			&_CustomerID,
			&_RefType,
			&_RefID,
			&_Metadata,
		)
		if err != nil {
			return nil, err
		}
		t := &IssueComment{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _Body.Valid {
			t.SetBody(_Body.String)
		}
		if _Deleted.Valid {
			t.SetDeleted(_Deleted.Bool)
		}
		if _DeletedUserID.Valid {
			t.SetDeletedUserID(_DeletedUserID.String)
		}
		if _CreatedAt.Valid {
			t.SetCreatedAt(_CreatedAt.Int64)
		}
		if _UpdatedAt.Valid {
			t.SetUpdatedAt(_UpdatedAt.Int64)
		}
		if _DeletedAt.Valid {
			t.SetDeletedAt(_DeletedAt.Int64)
		}
		if _URL.Valid {
			t.SetURL(_URL.String)
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

// GetProjectID will return the IssueComment ProjectID value
func (t *IssueComment) GetProjectID() string {
	return t.ProjectID
}

// SetProjectID will set the IssueComment ProjectID value
func (t *IssueComment) SetProjectID(v string) {
	t.ProjectID = v
}

// GetBody will return the IssueComment Body value
func (t *IssueComment) GetBody() string {
	return t.Body
}

// SetBody will set the IssueComment Body value
func (t *IssueComment) SetBody(v string) {
	t.Body = v
}

// GetDeleted will return the IssueComment Deleted value
func (t *IssueComment) GetDeleted() bool {
	return t.Deleted
}

// SetDeleted will set the IssueComment Deleted value
func (t *IssueComment) SetDeleted(v bool) {
	t.Deleted = v
}

// GetDeletedUserID will return the IssueComment DeletedUserID value
func (t *IssueComment) GetDeletedUserID() string {
	if t.DeletedUserID == nil {
		return ""
	}
	return *t.DeletedUserID
}

// SetDeletedUserID will set the IssueComment DeletedUserID value
func (t *IssueComment) SetDeletedUserID(v string) {
	t.DeletedUserID = &v
}

// GetCreatedAt will return the IssueComment CreatedAt value
func (t *IssueComment) GetCreatedAt() int64 {
	return t.CreatedAt
}

// SetCreatedAt will set the IssueComment CreatedAt value
func (t *IssueComment) SetCreatedAt(v int64) {
	t.CreatedAt = v
}

// GetUpdatedAt will return the IssueComment UpdatedAt value
func (t *IssueComment) GetUpdatedAt() int64 {
	if t.UpdatedAt == nil {
		return int64(0)
	}
	return *t.UpdatedAt
}

// SetUpdatedAt will set the IssueComment UpdatedAt value
func (t *IssueComment) SetUpdatedAt(v int64) {
	t.UpdatedAt = &v
}

// GetDeletedAt will return the IssueComment DeletedAt value
func (t *IssueComment) GetDeletedAt() int64 {
	if t.DeletedAt == nil {
		return int64(0)
	}
	return *t.DeletedAt
}

// SetDeletedAt will set the IssueComment DeletedAt value
func (t *IssueComment) SetDeletedAt(v int64) {
	t.DeletedAt = &v
}

// GetURL will return the IssueComment URL value
func (t *IssueComment) GetURL() string {
	return t.URL
}

// SetURL will set the IssueComment URL value
func (t *IssueComment) SetURL(v string) {
	t.URL = v
}

// GetCustomerID will return the IssueComment CustomerID value
func (t *IssueComment) GetCustomerID() string {
	return t.CustomerID
}

// SetCustomerID will set the IssueComment CustomerID value
func (t *IssueComment) SetCustomerID(v string) {
	t.CustomerID = v
}

// FindIssueCommentsByCustomerID will find all IssueComments by the CustomerID value
func FindIssueCommentsByCustomerID(ctx context.Context, db *sql.DB, value string) ([]*IssueComment, error) {
	q := "SELECT `issue_comment`.`id`,`issue_comment`.`checksum`,`issue_comment`.`issue_id`,`issue_comment`.`user_id`,`issue_comment`.`project_id`,`issue_comment`.`body`,`issue_comment`.`deleted`,`issue_comment`.`deleted_user_id`,`issue_comment`.`created_at`,`issue_comment`.`updated_at`,`issue_comment`.`deleted_at`,`issue_comment`.`url`,`issue_comment`.`customer_id`,`issue_comment`.`ref_type`,`issue_comment`.`ref_id`,`issue_comment`.`metadata` FROM `issue_comment` WHERE `customer_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*IssueComment, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _IssueID sql.NullString
		var _UserID sql.NullString
		var _ProjectID sql.NullString
		var _Body sql.NullString
		var _Deleted sql.NullBool
		var _DeletedUserID sql.NullString
		var _CreatedAt sql.NullInt64
		var _UpdatedAt sql.NullInt64
		var _DeletedAt sql.NullInt64
		var _URL sql.NullString
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefID sql.NullString
		var _Metadata sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_IssueID,
			&_UserID,
			&_ProjectID,
			&_Body,
			&_Deleted,
			&_DeletedUserID,
			&_CreatedAt,
			&_UpdatedAt,
			&_DeletedAt,
			&_URL,
			&_CustomerID,
			&_RefType,
			&_RefID,
			&_Metadata,
		)
		if err != nil {
			return nil, err
		}
		t := &IssueComment{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _Body.Valid {
			t.SetBody(_Body.String)
		}
		if _Deleted.Valid {
			t.SetDeleted(_Deleted.Bool)
		}
		if _DeletedUserID.Valid {
			t.SetDeletedUserID(_DeletedUserID.String)
		}
		if _CreatedAt.Valid {
			t.SetCreatedAt(_CreatedAt.Int64)
		}
		if _UpdatedAt.Valid {
			t.SetUpdatedAt(_UpdatedAt.Int64)
		}
		if _DeletedAt.Valid {
			t.SetDeletedAt(_DeletedAt.Int64)
		}
		if _URL.Valid {
			t.SetURL(_URL.String)
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

// FindIssueCommentsByCustomerIDTx will find all IssueComments by the CustomerID value using the provided transaction
func FindIssueCommentsByCustomerIDTx(ctx context.Context, tx *sql.Tx, value string) ([]*IssueComment, error) {
	q := "SELECT `issue_comment`.`id`,`issue_comment`.`checksum`,`issue_comment`.`issue_id`,`issue_comment`.`user_id`,`issue_comment`.`project_id`,`issue_comment`.`body`,`issue_comment`.`deleted`,`issue_comment`.`deleted_user_id`,`issue_comment`.`created_at`,`issue_comment`.`updated_at`,`issue_comment`.`deleted_at`,`issue_comment`.`url`,`issue_comment`.`customer_id`,`issue_comment`.`ref_type`,`issue_comment`.`ref_id`,`issue_comment`.`metadata` FROM `issue_comment` WHERE `customer_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*IssueComment, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _IssueID sql.NullString
		var _UserID sql.NullString
		var _ProjectID sql.NullString
		var _Body sql.NullString
		var _Deleted sql.NullBool
		var _DeletedUserID sql.NullString
		var _CreatedAt sql.NullInt64
		var _UpdatedAt sql.NullInt64
		var _DeletedAt sql.NullInt64
		var _URL sql.NullString
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefID sql.NullString
		var _Metadata sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_IssueID,
			&_UserID,
			&_ProjectID,
			&_Body,
			&_Deleted,
			&_DeletedUserID,
			&_CreatedAt,
			&_UpdatedAt,
			&_DeletedAt,
			&_URL,
			&_CustomerID,
			&_RefType,
			&_RefID,
			&_Metadata,
		)
		if err != nil {
			return nil, err
		}
		t := &IssueComment{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _Body.Valid {
			t.SetBody(_Body.String)
		}
		if _Deleted.Valid {
			t.SetDeleted(_Deleted.Bool)
		}
		if _DeletedUserID.Valid {
			t.SetDeletedUserID(_DeletedUserID.String)
		}
		if _CreatedAt.Valid {
			t.SetCreatedAt(_CreatedAt.Int64)
		}
		if _UpdatedAt.Valid {
			t.SetUpdatedAt(_UpdatedAt.Int64)
		}
		if _DeletedAt.Valid {
			t.SetDeletedAt(_DeletedAt.Int64)
		}
		if _URL.Valid {
			t.SetURL(_URL.String)
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

// GetRefType will return the IssueComment RefType value
func (t *IssueComment) GetRefType() string {
	return t.RefType
}

// SetRefType will set the IssueComment RefType value
func (t *IssueComment) SetRefType(v string) {
	t.RefType = v
}

// FindIssueCommentsByRefType will find all IssueComments by the RefType value
func FindIssueCommentsByRefType(ctx context.Context, db *sql.DB, value string) ([]*IssueComment, error) {
	q := "SELECT `issue_comment`.`id`,`issue_comment`.`checksum`,`issue_comment`.`issue_id`,`issue_comment`.`user_id`,`issue_comment`.`project_id`,`issue_comment`.`body`,`issue_comment`.`deleted`,`issue_comment`.`deleted_user_id`,`issue_comment`.`created_at`,`issue_comment`.`updated_at`,`issue_comment`.`deleted_at`,`issue_comment`.`url`,`issue_comment`.`customer_id`,`issue_comment`.`ref_type`,`issue_comment`.`ref_id`,`issue_comment`.`metadata` FROM `issue_comment` WHERE `ref_type` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*IssueComment, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _IssueID sql.NullString
		var _UserID sql.NullString
		var _ProjectID sql.NullString
		var _Body sql.NullString
		var _Deleted sql.NullBool
		var _DeletedUserID sql.NullString
		var _CreatedAt sql.NullInt64
		var _UpdatedAt sql.NullInt64
		var _DeletedAt sql.NullInt64
		var _URL sql.NullString
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefID sql.NullString
		var _Metadata sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_IssueID,
			&_UserID,
			&_ProjectID,
			&_Body,
			&_Deleted,
			&_DeletedUserID,
			&_CreatedAt,
			&_UpdatedAt,
			&_DeletedAt,
			&_URL,
			&_CustomerID,
			&_RefType,
			&_RefID,
			&_Metadata,
		)
		if err != nil {
			return nil, err
		}
		t := &IssueComment{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _Body.Valid {
			t.SetBody(_Body.String)
		}
		if _Deleted.Valid {
			t.SetDeleted(_Deleted.Bool)
		}
		if _DeletedUserID.Valid {
			t.SetDeletedUserID(_DeletedUserID.String)
		}
		if _CreatedAt.Valid {
			t.SetCreatedAt(_CreatedAt.Int64)
		}
		if _UpdatedAt.Valid {
			t.SetUpdatedAt(_UpdatedAt.Int64)
		}
		if _DeletedAt.Valid {
			t.SetDeletedAt(_DeletedAt.Int64)
		}
		if _URL.Valid {
			t.SetURL(_URL.String)
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

// FindIssueCommentsByRefTypeTx will find all IssueComments by the RefType value using the provided transaction
func FindIssueCommentsByRefTypeTx(ctx context.Context, tx *sql.Tx, value string) ([]*IssueComment, error) {
	q := "SELECT `issue_comment`.`id`,`issue_comment`.`checksum`,`issue_comment`.`issue_id`,`issue_comment`.`user_id`,`issue_comment`.`project_id`,`issue_comment`.`body`,`issue_comment`.`deleted`,`issue_comment`.`deleted_user_id`,`issue_comment`.`created_at`,`issue_comment`.`updated_at`,`issue_comment`.`deleted_at`,`issue_comment`.`url`,`issue_comment`.`customer_id`,`issue_comment`.`ref_type`,`issue_comment`.`ref_id`,`issue_comment`.`metadata` FROM `issue_comment` WHERE `ref_type` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*IssueComment, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _IssueID sql.NullString
		var _UserID sql.NullString
		var _ProjectID sql.NullString
		var _Body sql.NullString
		var _Deleted sql.NullBool
		var _DeletedUserID sql.NullString
		var _CreatedAt sql.NullInt64
		var _UpdatedAt sql.NullInt64
		var _DeletedAt sql.NullInt64
		var _URL sql.NullString
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefID sql.NullString
		var _Metadata sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_IssueID,
			&_UserID,
			&_ProjectID,
			&_Body,
			&_Deleted,
			&_DeletedUserID,
			&_CreatedAt,
			&_UpdatedAt,
			&_DeletedAt,
			&_URL,
			&_CustomerID,
			&_RefType,
			&_RefID,
			&_Metadata,
		)
		if err != nil {
			return nil, err
		}
		t := &IssueComment{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _Body.Valid {
			t.SetBody(_Body.String)
		}
		if _Deleted.Valid {
			t.SetDeleted(_Deleted.Bool)
		}
		if _DeletedUserID.Valid {
			t.SetDeletedUserID(_DeletedUserID.String)
		}
		if _CreatedAt.Valid {
			t.SetCreatedAt(_CreatedAt.Int64)
		}
		if _UpdatedAt.Valid {
			t.SetUpdatedAt(_UpdatedAt.Int64)
		}
		if _DeletedAt.Valid {
			t.SetDeletedAt(_DeletedAt.Int64)
		}
		if _URL.Valid {
			t.SetURL(_URL.String)
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

// GetRefID will return the IssueComment RefID value
func (t *IssueComment) GetRefID() string {
	return t.RefID
}

// SetRefID will set the IssueComment RefID value
func (t *IssueComment) SetRefID(v string) {
	t.RefID = v
}

// FindIssueCommentsByRefID will find all IssueComments by the RefID value
func FindIssueCommentsByRefID(ctx context.Context, db *sql.DB, value string) ([]*IssueComment, error) {
	q := "SELECT `issue_comment`.`id`,`issue_comment`.`checksum`,`issue_comment`.`issue_id`,`issue_comment`.`user_id`,`issue_comment`.`project_id`,`issue_comment`.`body`,`issue_comment`.`deleted`,`issue_comment`.`deleted_user_id`,`issue_comment`.`created_at`,`issue_comment`.`updated_at`,`issue_comment`.`deleted_at`,`issue_comment`.`url`,`issue_comment`.`customer_id`,`issue_comment`.`ref_type`,`issue_comment`.`ref_id`,`issue_comment`.`metadata` FROM `issue_comment` WHERE `ref_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*IssueComment, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _IssueID sql.NullString
		var _UserID sql.NullString
		var _ProjectID sql.NullString
		var _Body sql.NullString
		var _Deleted sql.NullBool
		var _DeletedUserID sql.NullString
		var _CreatedAt sql.NullInt64
		var _UpdatedAt sql.NullInt64
		var _DeletedAt sql.NullInt64
		var _URL sql.NullString
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefID sql.NullString
		var _Metadata sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_IssueID,
			&_UserID,
			&_ProjectID,
			&_Body,
			&_Deleted,
			&_DeletedUserID,
			&_CreatedAt,
			&_UpdatedAt,
			&_DeletedAt,
			&_URL,
			&_CustomerID,
			&_RefType,
			&_RefID,
			&_Metadata,
		)
		if err != nil {
			return nil, err
		}
		t := &IssueComment{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _Body.Valid {
			t.SetBody(_Body.String)
		}
		if _Deleted.Valid {
			t.SetDeleted(_Deleted.Bool)
		}
		if _DeletedUserID.Valid {
			t.SetDeletedUserID(_DeletedUserID.String)
		}
		if _CreatedAt.Valid {
			t.SetCreatedAt(_CreatedAt.Int64)
		}
		if _UpdatedAt.Valid {
			t.SetUpdatedAt(_UpdatedAt.Int64)
		}
		if _DeletedAt.Valid {
			t.SetDeletedAt(_DeletedAt.Int64)
		}
		if _URL.Valid {
			t.SetURL(_URL.String)
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

// FindIssueCommentsByRefIDTx will find all IssueComments by the RefID value using the provided transaction
func FindIssueCommentsByRefIDTx(ctx context.Context, tx *sql.Tx, value string) ([]*IssueComment, error) {
	q := "SELECT `issue_comment`.`id`,`issue_comment`.`checksum`,`issue_comment`.`issue_id`,`issue_comment`.`user_id`,`issue_comment`.`project_id`,`issue_comment`.`body`,`issue_comment`.`deleted`,`issue_comment`.`deleted_user_id`,`issue_comment`.`created_at`,`issue_comment`.`updated_at`,`issue_comment`.`deleted_at`,`issue_comment`.`url`,`issue_comment`.`customer_id`,`issue_comment`.`ref_type`,`issue_comment`.`ref_id`,`issue_comment`.`metadata` FROM `issue_comment` WHERE `ref_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*IssueComment, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _IssueID sql.NullString
		var _UserID sql.NullString
		var _ProjectID sql.NullString
		var _Body sql.NullString
		var _Deleted sql.NullBool
		var _DeletedUserID sql.NullString
		var _CreatedAt sql.NullInt64
		var _UpdatedAt sql.NullInt64
		var _DeletedAt sql.NullInt64
		var _URL sql.NullString
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefID sql.NullString
		var _Metadata sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_IssueID,
			&_UserID,
			&_ProjectID,
			&_Body,
			&_Deleted,
			&_DeletedUserID,
			&_CreatedAt,
			&_UpdatedAt,
			&_DeletedAt,
			&_URL,
			&_CustomerID,
			&_RefType,
			&_RefID,
			&_Metadata,
		)
		if err != nil {
			return nil, err
		}
		t := &IssueComment{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _Body.Valid {
			t.SetBody(_Body.String)
		}
		if _Deleted.Valid {
			t.SetDeleted(_Deleted.Bool)
		}
		if _DeletedUserID.Valid {
			t.SetDeletedUserID(_DeletedUserID.String)
		}
		if _CreatedAt.Valid {
			t.SetCreatedAt(_CreatedAt.Int64)
		}
		if _UpdatedAt.Valid {
			t.SetUpdatedAt(_UpdatedAt.Int64)
		}
		if _DeletedAt.Valid {
			t.SetDeletedAt(_DeletedAt.Int64)
		}
		if _URL.Valid {
			t.SetURL(_URL.String)
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

// GetMetadata will return the IssueComment Metadata value
func (t *IssueComment) GetMetadata() string {
	if t.Metadata == nil {
		return ""
	}
	return *t.Metadata
}

// SetMetadata will set the IssueComment Metadata value
func (t *IssueComment) SetMetadata(v string) {
	t.Metadata = &v
}

func (t *IssueComment) toTimestamp(value time.Time) *timestamp.Timestamp {
	ts, _ := ptypes.TimestampProto(value)
	return ts
}

// DBCreateIssueCommentTable will create the IssueComment table
func DBCreateIssueCommentTable(ctx context.Context, db *sql.DB) error {
	q := "CREATE TABLE `issue_comment` (`id` VARCHAR(64) NOT NULL PRIMARY KEY,`checksum` CHAR(64),`issue_id` VARCHAR(64) NOT NULL,`user_id`VARCHAR(64),`project_id`VARCHAR(64) NOT NULL,`body`LONGTEXT NOT NULL,`deleted`TINYINT(1) NOT NULL,`deleted_user_id` VARCHAR(64),`created_at`BIGINT(20) UNSIGNED NOT NULL,`updated_at`BIGINT(20) UNSIGNED,`deleted_at`BIGINT(20) UNSIGNED,`url` VARCHAR(255) NOT NULL,`customer_id` VARCHAR(64) NOT NULL,`ref_type` VARCHAR(20) NOT NULL,`ref_id` VARCHAR(64) NOT NULL,`metadata` JSON,INDEX issue_comment_issue_id_index (`issue_id`),INDEX issue_comment_user_id_index (`user_id`),INDEX issue_comment_customer_id_index (`customer_id`),INDEX issue_comment_ref_type_index (`ref_type`),INDEX issue_comment_ref_id_index (`ref_id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBCreateIssueCommentTableTx will create the IssueComment table using the provided transction
func DBCreateIssueCommentTableTx(ctx context.Context, tx *sql.Tx) error {
	q := "CREATE TABLE `issue_comment` (`id` VARCHAR(64) NOT NULL PRIMARY KEY,`checksum` CHAR(64),`issue_id` VARCHAR(64) NOT NULL,`user_id`VARCHAR(64),`project_id`VARCHAR(64) NOT NULL,`body`LONGTEXT NOT NULL,`deleted`TINYINT(1) NOT NULL,`deleted_user_id` VARCHAR(64),`created_at`BIGINT(20) UNSIGNED NOT NULL,`updated_at`BIGINT(20) UNSIGNED,`deleted_at`BIGINT(20) UNSIGNED,`url` VARCHAR(255) NOT NULL,`customer_id` VARCHAR(64) NOT NULL,`ref_type` VARCHAR(20) NOT NULL,`ref_id` VARCHAR(64) NOT NULL,`metadata` JSON,INDEX issue_comment_issue_id_index (`issue_id`),INDEX issue_comment_user_id_index (`user_id`),INDEX issue_comment_customer_id_index (`customer_id`),INDEX issue_comment_ref_type_index (`ref_type`),INDEX issue_comment_ref_id_index (`ref_id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// DBDropIssueCommentTable will drop the IssueComment table
func DBDropIssueCommentTable(ctx context.Context, db *sql.DB) error {
	q := "DROP TABLE IF EXISTS `issue_comment`"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBDropIssueCommentTableTx will drop the IssueComment table using the provided transaction
func DBDropIssueCommentTableTx(ctx context.Context, tx *sql.Tx) error {
	q := "DROP TABLE IF EXISTS `issue_comment`"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// CalculateChecksum will calculate a checksum of the SHA1 of all field values
func (t *IssueComment) CalculateChecksum() string {
	return orm.HashStrings(
		orm.ToString(t.ID),
		orm.ToString(t.IssueID),
		orm.ToString(t.UserID),
		orm.ToString(t.ProjectID),
		orm.ToString(t.Body),
		orm.ToString(t.Deleted),
		orm.ToString(t.DeletedUserID),
		orm.ToString(t.CreatedAt),
		orm.ToString(t.UpdatedAt),
		orm.ToString(t.DeletedAt),
		orm.ToString(t.URL),
		orm.ToString(t.CustomerID),
		orm.ToString(t.RefType),
		orm.ToString(t.RefID),
		orm.ToString(t.Metadata),
	)
}

// DBCreate will create a new IssueComment record in the database
func (t *IssueComment) DBCreate(ctx context.Context, db *sql.DB) (sql.Result, error) {
	q := "INSERT INTO `issue_comment` (`issue_comment`.`id`,`issue_comment`.`checksum`,`issue_comment`.`issue_id`,`issue_comment`.`user_id`,`issue_comment`.`project_id`,`issue_comment`.`body`,`issue_comment`.`deleted`,`issue_comment`.`deleted_user_id`,`issue_comment`.`created_at`,`issue_comment`.`updated_at`,`issue_comment`.`deleted_at`,`issue_comment`.`url`,`issue_comment`.`customer_id`,`issue_comment`.`ref_type`,`issue_comment`.`ref_id`,`issue_comment`.`metadata`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.IssueID),
		orm.ToSQLString(t.UserID),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLString(t.Body),
		orm.ToSQLBool(t.Deleted),
		orm.ToSQLString(t.DeletedUserID),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
		orm.ToSQLInt64(t.DeletedAt),
		orm.ToSQLString(t.URL),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefType),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.Metadata),
	)
}

// DBCreateTx will create a new IssueComment record in the database using the provided transaction
func (t *IssueComment) DBCreateTx(ctx context.Context, tx *sql.Tx) (sql.Result, error) {
	q := "INSERT INTO `issue_comment` (`issue_comment`.`id`,`issue_comment`.`checksum`,`issue_comment`.`issue_id`,`issue_comment`.`user_id`,`issue_comment`.`project_id`,`issue_comment`.`body`,`issue_comment`.`deleted`,`issue_comment`.`deleted_user_id`,`issue_comment`.`created_at`,`issue_comment`.`updated_at`,`issue_comment`.`deleted_at`,`issue_comment`.`url`,`issue_comment`.`customer_id`,`issue_comment`.`ref_type`,`issue_comment`.`ref_id`,`issue_comment`.`metadata`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.IssueID),
		orm.ToSQLString(t.UserID),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLString(t.Body),
		orm.ToSQLBool(t.Deleted),
		orm.ToSQLString(t.DeletedUserID),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
		orm.ToSQLInt64(t.DeletedAt),
		orm.ToSQLString(t.URL),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefType),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.Metadata),
	)
}

// DBCreateIgnoreDuplicate will upsert the IssueComment record in the database
func (t *IssueComment) DBCreateIgnoreDuplicate(ctx context.Context, db *sql.DB) (sql.Result, error) {
	q := "INSERT INTO `issue_comment` (`issue_comment`.`id`,`issue_comment`.`checksum`,`issue_comment`.`issue_id`,`issue_comment`.`user_id`,`issue_comment`.`project_id`,`issue_comment`.`body`,`issue_comment`.`deleted`,`issue_comment`.`deleted_user_id`,`issue_comment`.`created_at`,`issue_comment`.`updated_at`,`issue_comment`.`deleted_at`,`issue_comment`.`url`,`issue_comment`.`customer_id`,`issue_comment`.`ref_type`,`issue_comment`.`ref_id`,`issue_comment`.`metadata`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.IssueID),
		orm.ToSQLString(t.UserID),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLString(t.Body),
		orm.ToSQLBool(t.Deleted),
		orm.ToSQLString(t.DeletedUserID),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
		orm.ToSQLInt64(t.DeletedAt),
		orm.ToSQLString(t.URL),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefType),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.Metadata),
	)
}

// DBCreateIgnoreDuplicateTx will upsert the IssueComment record in the database using the provided transaction
func (t *IssueComment) DBCreateIgnoreDuplicateTx(ctx context.Context, tx *sql.Tx) (sql.Result, error) {
	q := "INSERT INTO `issue_comment` (`issue_comment`.`id`,`issue_comment`.`checksum`,`issue_comment`.`issue_id`,`issue_comment`.`user_id`,`issue_comment`.`project_id`,`issue_comment`.`body`,`issue_comment`.`deleted`,`issue_comment`.`deleted_user_id`,`issue_comment`.`created_at`,`issue_comment`.`updated_at`,`issue_comment`.`deleted_at`,`issue_comment`.`url`,`issue_comment`.`customer_id`,`issue_comment`.`ref_type`,`issue_comment`.`ref_id`,`issue_comment`.`metadata`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.IssueID),
		orm.ToSQLString(t.UserID),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLString(t.Body),
		orm.ToSQLBool(t.Deleted),
		orm.ToSQLString(t.DeletedUserID),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
		orm.ToSQLInt64(t.DeletedAt),
		orm.ToSQLString(t.URL),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefType),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.Metadata),
	)
}

// DeleteAllIssueComments deletes all IssueComment records in the database with optional filters
func DeleteAllIssueComments(ctx context.Context, db *sql.DB, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(IssueCommentTableName),
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

// DeleteAllIssueCommentsTx deletes all IssueComment records in the database with optional filters using the provided transaction
func DeleteAllIssueCommentsTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(IssueCommentTableName),
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

// DBDelete will delete this IssueComment record in the database
func (t *IssueComment) DBDelete(ctx context.Context, db *sql.DB) (bool, error) {
	q := "DELETE FROM `issue_comment` WHERE `id` = ?"
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

// DBDeleteTx will delete this IssueComment record in the database using the provided transaction
func (t *IssueComment) DBDeleteTx(ctx context.Context, tx *sql.Tx) (bool, error) {
	q := "DELETE FROM `issue_comment` WHERE `id` = ?"
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

// DBUpdate will update the IssueComment record in the database
func (t *IssueComment) DBUpdate(ctx context.Context, db *sql.DB) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `issue_comment` SET `checksum`=?,`issue_id`=?,`user_id`=?,`project_id`=?,`body`=?,`deleted`=?,`deleted_user_id`=?,`created_at`=?,`updated_at`=?,`deleted_at`=?,`url`=?,`customer_id`=?,`ref_type`=?,`ref_id`=?,`metadata`=? WHERE `id`=?"
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.IssueID),
		orm.ToSQLString(t.UserID),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLString(t.Body),
		orm.ToSQLBool(t.Deleted),
		orm.ToSQLString(t.DeletedUserID),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
		orm.ToSQLInt64(t.DeletedAt),
		orm.ToSQLString(t.URL),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefType),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.Metadata),
		orm.ToSQLString(t.ID),
	)
}

// DBUpdateTx will update the IssueComment record in the database using the provided transaction
func (t *IssueComment) DBUpdateTx(ctx context.Context, tx *sql.Tx) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `issue_comment` SET `checksum`=?,`issue_id`=?,`user_id`=?,`project_id`=?,`body`=?,`deleted`=?,`deleted_user_id`=?,`created_at`=?,`updated_at`=?,`deleted_at`=?,`url`=?,`customer_id`=?,`ref_type`=?,`ref_id`=?,`metadata`=? WHERE `id`=?"
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.IssueID),
		orm.ToSQLString(t.UserID),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLString(t.Body),
		orm.ToSQLBool(t.Deleted),
		orm.ToSQLString(t.DeletedUserID),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
		orm.ToSQLInt64(t.DeletedAt),
		orm.ToSQLString(t.URL),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefType),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.Metadata),
		orm.ToSQLString(t.ID),
	)
}

// DBUpsert will upsert the IssueComment record in the database
func (t *IssueComment) DBUpsert(ctx context.Context, db *sql.DB, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `issue_comment` (`issue_comment`.`id`,`issue_comment`.`checksum`,`issue_comment`.`issue_id`,`issue_comment`.`user_id`,`issue_comment`.`project_id`,`issue_comment`.`body`,`issue_comment`.`deleted`,`issue_comment`.`deleted_user_id`,`issue_comment`.`created_at`,`issue_comment`.`updated_at`,`issue_comment`.`deleted_at`,`issue_comment`.`url`,`issue_comment`.`customer_id`,`issue_comment`.`ref_type`,`issue_comment`.`ref_id`,`issue_comment`.`metadata`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `issue_comment` (`issue_comment`.`id`,`issue_comment`.`checksum`,`issue_comment`.`issue_id`,`issue_comment`.`user_id`,`issue_comment`.`project_id`,`issue_comment`.`body`,`issue_comment`.`deleted`,`issue_comment`.`deleted_user_id`,`issue_comment`.`created_at`,`issue_comment`.`updated_at`,`issue_comment`.`deleted_at`,`issue_comment`.`url`,`issue_comment`.`customer_id`,`issue_comment`.`ref_type`,`issue_comment`.`ref_id`,`issue_comment`.`metadata`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`issue_id`=VALUES(`issue_id`),`user_id`=VALUES(`user_id`),`project_id`=VALUES(`project_id`),`body`=VALUES(`body`),`deleted`=VALUES(`deleted`),`deleted_user_id`=VALUES(`deleted_user_id`),`created_at`=VALUES(`created_at`),`updated_at`=VALUES(`updated_at`),`deleted_at`=VALUES(`deleted_at`),`url`=VALUES(`url`),`customer_id`=VALUES(`customer_id`),`ref_type`=VALUES(`ref_type`),`ref_id`=VALUES(`ref_id`),`metadata`=VALUES(`metadata`)"
	}
	r, err := db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.IssueID),
		orm.ToSQLString(t.UserID),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLString(t.Body),
		orm.ToSQLBool(t.Deleted),
		orm.ToSQLString(t.DeletedUserID),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
		orm.ToSQLInt64(t.DeletedAt),
		orm.ToSQLString(t.URL),
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

// DBUpsertTx will upsert the IssueComment record in the database using the provided transaction
func (t *IssueComment) DBUpsertTx(ctx context.Context, tx *sql.Tx, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `issue_comment` (`issue_comment`.`id`,`issue_comment`.`checksum`,`issue_comment`.`issue_id`,`issue_comment`.`user_id`,`issue_comment`.`project_id`,`issue_comment`.`body`,`issue_comment`.`deleted`,`issue_comment`.`deleted_user_id`,`issue_comment`.`created_at`,`issue_comment`.`updated_at`,`issue_comment`.`deleted_at`,`issue_comment`.`url`,`issue_comment`.`customer_id`,`issue_comment`.`ref_type`,`issue_comment`.`ref_id`,`issue_comment`.`metadata`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `issue_comment` (`issue_comment`.`id`,`issue_comment`.`checksum`,`issue_comment`.`issue_id`,`issue_comment`.`user_id`,`issue_comment`.`project_id`,`issue_comment`.`body`,`issue_comment`.`deleted`,`issue_comment`.`deleted_user_id`,`issue_comment`.`created_at`,`issue_comment`.`updated_at`,`issue_comment`.`deleted_at`,`issue_comment`.`url`,`issue_comment`.`customer_id`,`issue_comment`.`ref_type`,`issue_comment`.`ref_id`,`issue_comment`.`metadata`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`issue_id`=VALUES(`issue_id`),`user_id`=VALUES(`user_id`),`project_id`=VALUES(`project_id`),`body`=VALUES(`body`),`deleted`=VALUES(`deleted`),`deleted_user_id`=VALUES(`deleted_user_id`),`created_at`=VALUES(`created_at`),`updated_at`=VALUES(`updated_at`),`deleted_at`=VALUES(`deleted_at`),`url`=VALUES(`url`),`customer_id`=VALUES(`customer_id`),`ref_type`=VALUES(`ref_type`),`ref_id`=VALUES(`ref_id`),`metadata`=VALUES(`metadata`)"
	}
	r, err := tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.IssueID),
		orm.ToSQLString(t.UserID),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLString(t.Body),
		orm.ToSQLBool(t.Deleted),
		orm.ToSQLString(t.DeletedUserID),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
		orm.ToSQLInt64(t.DeletedAt),
		orm.ToSQLString(t.URL),
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

// DBFindOne will find a IssueComment record in the database with the primary key
func (t *IssueComment) DBFindOne(ctx context.Context, db *sql.DB, value string) (bool, error) {
	q := "SELECT `issue_comment`.`id`,`issue_comment`.`checksum`,`issue_comment`.`issue_id`,`issue_comment`.`user_id`,`issue_comment`.`project_id`,`issue_comment`.`body`,`issue_comment`.`deleted`,`issue_comment`.`deleted_user_id`,`issue_comment`.`created_at`,`issue_comment`.`updated_at`,`issue_comment`.`deleted_at`,`issue_comment`.`url`,`issue_comment`.`customer_id`,`issue_comment`.`ref_type`,`issue_comment`.`ref_id`,`issue_comment`.`metadata` FROM `issue_comment` WHERE `id` = ? LIMIT 1"
	row := db.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _IssueID sql.NullString
	var _UserID sql.NullString
	var _ProjectID sql.NullString
	var _Body sql.NullString
	var _Deleted sql.NullBool
	var _DeletedUserID sql.NullString
	var _CreatedAt sql.NullInt64
	var _UpdatedAt sql.NullInt64
	var _DeletedAt sql.NullInt64
	var _URL sql.NullString
	var _CustomerID sql.NullString
	var _RefType sql.NullString
	var _RefID sql.NullString
	var _Metadata sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_IssueID,
		&_UserID,
		&_ProjectID,
		&_Body,
		&_Deleted,
		&_DeletedUserID,
		&_CreatedAt,
		&_UpdatedAt,
		&_DeletedAt,
		&_URL,
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
	if _IssueID.Valid {
		t.SetIssueID(_IssueID.String)
	}
	if _UserID.Valid {
		t.SetUserID(_UserID.String)
	}
	if _ProjectID.Valid {
		t.SetProjectID(_ProjectID.String)
	}
	if _Body.Valid {
		t.SetBody(_Body.String)
	}
	if _Deleted.Valid {
		t.SetDeleted(_Deleted.Bool)
	}
	if _DeletedUserID.Valid {
		t.SetDeletedUserID(_DeletedUserID.String)
	}
	if _CreatedAt.Valid {
		t.SetCreatedAt(_CreatedAt.Int64)
	}
	if _UpdatedAt.Valid {
		t.SetUpdatedAt(_UpdatedAt.Int64)
	}
	if _DeletedAt.Valid {
		t.SetDeletedAt(_DeletedAt.Int64)
	}
	if _URL.Valid {
		t.SetURL(_URL.String)
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

// DBFindOneTx will find a IssueComment record in the database with the primary key using the provided transaction
func (t *IssueComment) DBFindOneTx(ctx context.Context, tx *sql.Tx, value string) (bool, error) {
	q := "SELECT `issue_comment`.`id`,`issue_comment`.`checksum`,`issue_comment`.`issue_id`,`issue_comment`.`user_id`,`issue_comment`.`project_id`,`issue_comment`.`body`,`issue_comment`.`deleted`,`issue_comment`.`deleted_user_id`,`issue_comment`.`created_at`,`issue_comment`.`updated_at`,`issue_comment`.`deleted_at`,`issue_comment`.`url`,`issue_comment`.`customer_id`,`issue_comment`.`ref_type`,`issue_comment`.`ref_id`,`issue_comment`.`metadata` FROM `issue_comment` WHERE `id` = ? LIMIT 1"
	row := tx.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _IssueID sql.NullString
	var _UserID sql.NullString
	var _ProjectID sql.NullString
	var _Body sql.NullString
	var _Deleted sql.NullBool
	var _DeletedUserID sql.NullString
	var _CreatedAt sql.NullInt64
	var _UpdatedAt sql.NullInt64
	var _DeletedAt sql.NullInt64
	var _URL sql.NullString
	var _CustomerID sql.NullString
	var _RefType sql.NullString
	var _RefID sql.NullString
	var _Metadata sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_IssueID,
		&_UserID,
		&_ProjectID,
		&_Body,
		&_Deleted,
		&_DeletedUserID,
		&_CreatedAt,
		&_UpdatedAt,
		&_DeletedAt,
		&_URL,
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
	if _IssueID.Valid {
		t.SetIssueID(_IssueID.String)
	}
	if _UserID.Valid {
		t.SetUserID(_UserID.String)
	}
	if _ProjectID.Valid {
		t.SetProjectID(_ProjectID.String)
	}
	if _Body.Valid {
		t.SetBody(_Body.String)
	}
	if _Deleted.Valid {
		t.SetDeleted(_Deleted.Bool)
	}
	if _DeletedUserID.Valid {
		t.SetDeletedUserID(_DeletedUserID.String)
	}
	if _CreatedAt.Valid {
		t.SetCreatedAt(_CreatedAt.Int64)
	}
	if _UpdatedAt.Valid {
		t.SetUpdatedAt(_UpdatedAt.Int64)
	}
	if _DeletedAt.Valid {
		t.SetDeletedAt(_DeletedAt.Int64)
	}
	if _URL.Valid {
		t.SetURL(_URL.String)
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

// FindIssueComments will find a IssueComment record in the database with the provided parameters
func FindIssueComments(ctx context.Context, db *sql.DB, _params ...interface{}) ([]*IssueComment, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("issue_id"),
		orm.Column("user_id"),
		orm.Column("project_id"),
		orm.Column("body"),
		orm.Column("deleted"),
		orm.Column("deleted_user_id"),
		orm.Column("created_at"),
		orm.Column("updated_at"),
		orm.Column("deleted_at"),
		orm.Column("url"),
		orm.Column("customer_id"),
		orm.Column("ref_type"),
		orm.Column("ref_id"),
		orm.Column("metadata"),
		orm.Table(IssueCommentTableName),
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
	results := make([]*IssueComment, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _IssueID sql.NullString
		var _UserID sql.NullString
		var _ProjectID sql.NullString
		var _Body sql.NullString
		var _Deleted sql.NullBool
		var _DeletedUserID sql.NullString
		var _CreatedAt sql.NullInt64
		var _UpdatedAt sql.NullInt64
		var _DeletedAt sql.NullInt64
		var _URL sql.NullString
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefID sql.NullString
		var _Metadata sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_IssueID,
			&_UserID,
			&_ProjectID,
			&_Body,
			&_Deleted,
			&_DeletedUserID,
			&_CreatedAt,
			&_UpdatedAt,
			&_DeletedAt,
			&_URL,
			&_CustomerID,
			&_RefType,
			&_RefID,
			&_Metadata,
		)
		if err != nil {
			return nil, err
		}
		t := &IssueComment{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _Body.Valid {
			t.SetBody(_Body.String)
		}
		if _Deleted.Valid {
			t.SetDeleted(_Deleted.Bool)
		}
		if _DeletedUserID.Valid {
			t.SetDeletedUserID(_DeletedUserID.String)
		}
		if _CreatedAt.Valid {
			t.SetCreatedAt(_CreatedAt.Int64)
		}
		if _UpdatedAt.Valid {
			t.SetUpdatedAt(_UpdatedAt.Int64)
		}
		if _DeletedAt.Valid {
			t.SetDeletedAt(_DeletedAt.Int64)
		}
		if _URL.Valid {
			t.SetURL(_URL.String)
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

// FindIssueCommentsTx will find a IssueComment record in the database with the provided parameters using the provided transaction
func FindIssueCommentsTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) ([]*IssueComment, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("issue_id"),
		orm.Column("user_id"),
		orm.Column("project_id"),
		orm.Column("body"),
		orm.Column("deleted"),
		orm.Column("deleted_user_id"),
		orm.Column("created_at"),
		orm.Column("updated_at"),
		orm.Column("deleted_at"),
		orm.Column("url"),
		orm.Column("customer_id"),
		orm.Column("ref_type"),
		orm.Column("ref_id"),
		orm.Column("metadata"),
		orm.Table(IssueCommentTableName),
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
	results := make([]*IssueComment, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _IssueID sql.NullString
		var _UserID sql.NullString
		var _ProjectID sql.NullString
		var _Body sql.NullString
		var _Deleted sql.NullBool
		var _DeletedUserID sql.NullString
		var _CreatedAt sql.NullInt64
		var _UpdatedAt sql.NullInt64
		var _DeletedAt sql.NullInt64
		var _URL sql.NullString
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefID sql.NullString
		var _Metadata sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_IssueID,
			&_UserID,
			&_ProjectID,
			&_Body,
			&_Deleted,
			&_DeletedUserID,
			&_CreatedAt,
			&_UpdatedAt,
			&_DeletedAt,
			&_URL,
			&_CustomerID,
			&_RefType,
			&_RefID,
			&_Metadata,
		)
		if err != nil {
			return nil, err
		}
		t := &IssueComment{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _IssueID.Valid {
			t.SetIssueID(_IssueID.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _Body.Valid {
			t.SetBody(_Body.String)
		}
		if _Deleted.Valid {
			t.SetDeleted(_Deleted.Bool)
		}
		if _DeletedUserID.Valid {
			t.SetDeletedUserID(_DeletedUserID.String)
		}
		if _CreatedAt.Valid {
			t.SetCreatedAt(_CreatedAt.Int64)
		}
		if _UpdatedAt.Valid {
			t.SetUpdatedAt(_UpdatedAt.Int64)
		}
		if _DeletedAt.Valid {
			t.SetDeletedAt(_DeletedAt.Int64)
		}
		if _URL.Valid {
			t.SetURL(_URL.String)
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

// DBFind will find a IssueComment record in the database with the provided parameters
func (t *IssueComment) DBFind(ctx context.Context, db *sql.DB, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("issue_id"),
		orm.Column("user_id"),
		orm.Column("project_id"),
		orm.Column("body"),
		orm.Column("deleted"),
		orm.Column("deleted_user_id"),
		orm.Column("created_at"),
		orm.Column("updated_at"),
		orm.Column("deleted_at"),
		orm.Column("url"),
		orm.Column("customer_id"),
		orm.Column("ref_type"),
		orm.Column("ref_id"),
		orm.Column("metadata"),
		orm.Table(IssueCommentTableName),
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
	var _IssueID sql.NullString
	var _UserID sql.NullString
	var _ProjectID sql.NullString
	var _Body sql.NullString
	var _Deleted sql.NullBool
	var _DeletedUserID sql.NullString
	var _CreatedAt sql.NullInt64
	var _UpdatedAt sql.NullInt64
	var _DeletedAt sql.NullInt64
	var _URL sql.NullString
	var _CustomerID sql.NullString
	var _RefType sql.NullString
	var _RefID sql.NullString
	var _Metadata sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_IssueID,
		&_UserID,
		&_ProjectID,
		&_Body,
		&_Deleted,
		&_DeletedUserID,
		&_CreatedAt,
		&_UpdatedAt,
		&_DeletedAt,
		&_URL,
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
	if _IssueID.Valid {
		t.SetIssueID(_IssueID.String)
	}
	if _UserID.Valid {
		t.SetUserID(_UserID.String)
	}
	if _ProjectID.Valid {
		t.SetProjectID(_ProjectID.String)
	}
	if _Body.Valid {
		t.SetBody(_Body.String)
	}
	if _Deleted.Valid {
		t.SetDeleted(_Deleted.Bool)
	}
	if _DeletedUserID.Valid {
		t.SetDeletedUserID(_DeletedUserID.String)
	}
	if _CreatedAt.Valid {
		t.SetCreatedAt(_CreatedAt.Int64)
	}
	if _UpdatedAt.Valid {
		t.SetUpdatedAt(_UpdatedAt.Int64)
	}
	if _DeletedAt.Valid {
		t.SetDeletedAt(_DeletedAt.Int64)
	}
	if _URL.Valid {
		t.SetURL(_URL.String)
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

// DBFindTx will find a IssueComment record in the database with the provided parameters using the provided transaction
func (t *IssueComment) DBFindTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("issue_id"),
		orm.Column("user_id"),
		orm.Column("project_id"),
		orm.Column("body"),
		orm.Column("deleted"),
		orm.Column("deleted_user_id"),
		orm.Column("created_at"),
		orm.Column("updated_at"),
		orm.Column("deleted_at"),
		orm.Column("url"),
		orm.Column("customer_id"),
		orm.Column("ref_type"),
		orm.Column("ref_id"),
		orm.Column("metadata"),
		orm.Table(IssueCommentTableName),
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
	var _IssueID sql.NullString
	var _UserID sql.NullString
	var _ProjectID sql.NullString
	var _Body sql.NullString
	var _Deleted sql.NullBool
	var _DeletedUserID sql.NullString
	var _CreatedAt sql.NullInt64
	var _UpdatedAt sql.NullInt64
	var _DeletedAt sql.NullInt64
	var _URL sql.NullString
	var _CustomerID sql.NullString
	var _RefType sql.NullString
	var _RefID sql.NullString
	var _Metadata sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_IssueID,
		&_UserID,
		&_ProjectID,
		&_Body,
		&_Deleted,
		&_DeletedUserID,
		&_CreatedAt,
		&_UpdatedAt,
		&_DeletedAt,
		&_URL,
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
	if _IssueID.Valid {
		t.SetIssueID(_IssueID.String)
	}
	if _UserID.Valid {
		t.SetUserID(_UserID.String)
	}
	if _ProjectID.Valid {
		t.SetProjectID(_ProjectID.String)
	}
	if _Body.Valid {
		t.SetBody(_Body.String)
	}
	if _Deleted.Valid {
		t.SetDeleted(_Deleted.Bool)
	}
	if _DeletedUserID.Valid {
		t.SetDeletedUserID(_DeletedUserID.String)
	}
	if _CreatedAt.Valid {
		t.SetCreatedAt(_CreatedAt.Int64)
	}
	if _UpdatedAt.Valid {
		t.SetUpdatedAt(_UpdatedAt.Int64)
	}
	if _DeletedAt.Valid {
		t.SetDeletedAt(_DeletedAt.Int64)
	}
	if _URL.Valid {
		t.SetURL(_URL.String)
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

// CountIssueComments will find the count of IssueComment records in the database
func CountIssueComments(ctx context.Context, db *sql.DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(IssueCommentTableName),
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

// CountIssueCommentsTx will find the count of IssueComment records in the database using the provided transaction
func CountIssueCommentsTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(IssueCommentTableName),
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

// DBCount will find the count of IssueComment records in the database
func (t *IssueComment) DBCount(ctx context.Context, db *sql.DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(IssueCommentTableName),
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

// DBCountTx will find the count of IssueComment records in the database using the provided transaction
func (t *IssueComment) DBCountTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(IssueCommentTableName),
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

// DBExists will return true if the IssueComment record exists in the database
func (t *IssueComment) DBExists(ctx context.Context, db *sql.DB) (bool, error) {
	q := "SELECT `id` FROM `issue_comment` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := db.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// DBExistsTx will return true if the IssueComment record exists in the database using the provided transaction
func (t *IssueComment) DBExistsTx(ctx context.Context, tx *sql.Tx) (bool, error) {
	q := "SELECT `id` FROM `issue_comment` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := tx.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// PrimaryKeyColumn returns the column name for the primary key
func (t *IssueComment) PrimaryKeyColumn() string {
	return IssueCommentColumnID
}

// PrimaryKeyColumnType returns the primary key column Go type as a string
func (t *IssueComment) PrimaryKeyColumnType() string {
	return "string"
}

// PrimaryKey returns the primary key column value
func (t *IssueComment) PrimaryKey() interface{} {
	return t.ID
}
