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

var _ Model = (*Issue)(nil)
var _ CSVWriter = (*Issue)(nil)
var _ JSONWriter = (*Issue)(nil)
var _ Checksum = (*Issue)(nil)

// IssueTableName is the name of the table in SQL
const IssueTableName = "issue"

var IssueColumns = []string{
	"id",
	"checksum",
	"user_id",
	"project_id",
	"title",
	"body",
	"state",
	"type",
	"resolution",
	"created_at",
	"updated_at",
	"closed_at",
	"closedby_user_id",
	"url",
	"customer_id",
	"ref_type",
	"ref_id",
	"metadata",
}

type Issue_IssueState string

const (
	IssueState_OPEN   Issue_IssueState = "open"
	IssueState_CLOSED Issue_IssueState = "closed"
)

func (x Issue_IssueState) String() string {
	return string(x)
}

func enumIssueStateToString(v *Issue_IssueState) string {
	if v == nil {
		return ""
	}
	return v.String()
}

func toIssueState(v string) *Issue_IssueState {
	var ev *Issue_IssueState
	switch v {
	case "OPEN", "open":
		{
			v := IssueState_OPEN
			ev = &v
		}
	case "CLOSED", "closed":
		{
			v := IssueState_CLOSED
			ev = &v
		}
	}
	return ev
}

type Issue_IssueType string

const (
	IssueType_BUG         Issue_IssueType = "bug"
	IssueType_FEATURE     Issue_IssueType = "feature"
	IssueType_ENHANCEMENT Issue_IssueType = "enhancement"
	IssueType_OTHER       Issue_IssueType = "other"
)

func (x Issue_IssueType) String() string {
	return string(x)
}

func enumIssueTypeToString(v *Issue_IssueType) string {
	if v == nil {
		return ""
	}
	return v.String()
}

func toIssueType(v string) *Issue_IssueType {
	var ev *Issue_IssueType
	switch v {
	case "BUG", "bug":
		{
			v := IssueType_BUG
			ev = &v
		}
	case "FEATURE", "feature":
		{
			v := IssueType_FEATURE
			ev = &v
		}
	case "ENHANCEMENT", "enhancement":
		{
			v := IssueType_ENHANCEMENT
			ev = &v
		}
	case "OTHER", "other":
		{
			v := IssueType_OTHER
			ev = &v
		}
	}
	return ev
}

type Issue_IssueResolution string

const (
	IssueResolution_NONE       Issue_IssueResolution = "none"
	IssueResolution_RESOLVED   Issue_IssueResolution = "resolved"
	IssueResolution_UNRESOLVED Issue_IssueResolution = "unresolved"
)

func (x Issue_IssueResolution) String() string {
	return string(x)
}

func enumIssueResolutionToString(v *Issue_IssueResolution) string {
	if v == nil {
		return ""
	}
	return v.String()
}

func toIssueResolution(v string) *Issue_IssueResolution {
	var ev *Issue_IssueResolution
	switch v {
	case "NONE", "none":
		{
			v := IssueResolution_NONE
			ev = &v
		}
	case "RESOLVED", "resolved":
		{
			v := IssueResolution_RESOLVED
			ev = &v
		}
	case "UNRESOLVED", "unresolved":
		{
			v := IssueResolution_UNRESOLVED
			ev = &v
		}
	}
	return ev
}

// Issue table
type Issue struct {
	Body           *string               `json:"body,omitempty"`
	Checksum       *string               `json:"checksum,omitempty"`
	ClosedAt       *int64                `json:"closed_at,omitempty"`
	ClosedbyUserID *string               `json:"closedby_user_id,omitempty"`
	CreatedAt      int64                 `json:"created_at"`
	CustomerID     string                `json:"customer_id"`
	ID             string                `json:"id"`
	Metadata       *string               `json:"metadata,omitempty"`
	ProjectID      string                `json:"project_id"`
	RefID          string                `json:"ref_id"`
	RefType        string                `json:"ref_type"`
	Resolution     Issue_IssueResolution `json:"resolution"`
	State          Issue_IssueState      `json:"state"`
	Title          string                `json:"title"`
	Type           Issue_IssueType       `json:"type"`
	UpdatedAt      *int64                `json:"updated_at,omitempty"`
	URL            string                `json:"url"`
	UserID         *string               `json:"user_id,omitempty"`
}

// TableName returns the SQL table name for Issue and satifies the Model interface
func (t *Issue) TableName() string {
	return IssueTableName
}

// ToCSV will serialize the Issue instance to a CSV compatible array of strings
func (t *Issue) ToCSV() []string {
	return []string{
		t.ID,
		t.CalculateChecksum(),
		toCSVString(t.UserID),
		t.ProjectID,
		t.Title,
		toCSVString(t.Body),
		toCSVString(t.State),
		toCSVString(t.Type),
		toCSVString(t.Resolution),
		toCSVString(t.CreatedAt),
		toCSVString(t.UpdatedAt),
		toCSVString(t.ClosedAt),
		toCSVString(t.ClosedbyUserID),
		t.URL,
		t.CustomerID,
		t.RefType,
		t.RefID,
		toCSVString(t.Metadata),
	}
}

// WriteCSV will serialize the Issue instance to the writer as CSV and satisfies the CSVWriter interface
func (t *Issue) WriteCSV(w *csv.Writer) error {
	return w.Write(t.ToCSV())
}

// WriteJSON will serialize the Issue instance to the writer as JSON and satisfies the JSONWriter interface
func (t *Issue) WriteJSON(w io.Writer, indent ...bool) error {
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

// NewIssueReader creates a JSON reader which can read in Issue objects serialized as JSON either as an array, single object or json new lines
// and writes each Issue to the channel provided
func NewIssueReader(r io.Reader, ch chan<- Issue) error {
	return orm.Deserialize(r, func(buf json.RawMessage) error {
		dec := json.NewDecoder(bytes.NewBuffer(buf))
		e := Issue{}
		if err := dec.Decode(&e); err != nil {
			return err
		}
		ch <- e
		return nil
	})
}

// NewCSVIssueReaderDir will read the reader as CSV and emit each record to the channel provided
func NewCSVIssueReader(r io.Reader, ch chan<- Issue) error {
	cr := csv.NewReader(r)
	for {
		record, err := cr.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		ch <- Issue{
			ID:             record[0],
			Checksum:       fromStringPointer(record[1]),
			UserID:         fromStringPointer(record[2]),
			ProjectID:      record[3],
			Title:          record[4],
			Body:           fromStringPointer(record[5]),
			State:          *toIssueState(record[6]),
			Type:           *toIssueType(record[7]),
			Resolution:     *toIssueResolution(record[8]),
			CreatedAt:      fromCSVInt64(record[9]),
			UpdatedAt:      fromCSVInt64Pointer(record[10]),
			ClosedAt:       fromCSVInt64Pointer(record[11]),
			ClosedbyUserID: fromStringPointer(record[12]),
			URL:            record[13],
			CustomerID:     record[14],
			RefType:        record[15],
			RefID:          record[16],
			Metadata:       fromStringPointer(record[17]),
		}
	}
	return nil
}

// NewCSVIssueReaderFile will read the file as a CSV and emit each record to the channel provided
func NewCSVIssueReaderFile(fp string, ch chan<- Issue) error {
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
	return NewCSVIssueReader(fc, ch)
}

// NewCSVIssueReaderDir will read the issue.csv.gz file as a CSV and emit each record to the channel provided
func NewCSVIssueReaderDir(dir string, ch chan<- Issue) error {
	return NewCSVIssueReaderFile(filepath.Join(dir, "issue.csv.gz"), ch)
}

// IssueCSVDeduper is a function callback which takes the existing value (a) and the new value (b)
// and the return value should be the one to use (or a new one, if applicable). return nil
// to skip processing of this record
type IssueCSVDeduper func(a Issue, b Issue) *Issue

// IssueCSVDedupeDisabled is set on whether the CSV writer should de-dupe values by key
var IssueCSVDedupeDisabled bool

// NewIssueCSVWriterSize creates a batch writer that will write each Issue into a CSV file
// this method will automatically de-duplicate entries using the primary key. if the checksum
// for a newer item with the same primary key doesn't match a previously sent item, the newer
// one will be used
func NewIssueCSVWriterSize(w io.Writer, size int, dedupers ...IssueCSVDeduper) (chan Issue, chan bool, error) {
	cw := csv.NewWriter(w)
	ch := make(chan Issue, size)
	done := make(chan bool)
	go func() {
		defer func() { done <- true }()
		dodedupe := !IssueCSVDedupeDisabled
		var kv map[string]*Issue
		var deduper IssueCSVDeduper
		if dedupers != nil && len(dedupers) > 0 {
			deduper = dedupers[0]
			dodedupe = true
		}
		if dodedupe {
			kv = make(map[string]*Issue)
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

// IssueCSVDefaultSize is the default channel buffer size if not provided
var IssueCSVDefaultSize = 100

// NewIssueCSVWriter creates a batch writer that will write each Issue into a CSV file
func NewIssueCSVWriter(w io.Writer, dedupers ...IssueCSVDeduper) (chan Issue, chan bool, error) {
	return NewIssueCSVWriterSize(w, IssueCSVDefaultSize, dedupers...)
}

// NewIssueCSVWriterDir creates a batch writer that will write each Issue into a CSV file named issue.csv.gz in dir
func NewIssueCSVWriterDir(dir string, dedupers ...IssueCSVDeduper) (chan Issue, chan bool, error) {
	return NewIssueCSVWriterFile(filepath.Join(dir, "issue.csv.gz"), dedupers...)
}

// NewIssueCSVWriterFile creates a batch writer that will write each Issue into a CSV file
func NewIssueCSVWriterFile(fn string, dedupers ...IssueCSVDeduper) (chan Issue, chan bool, error) {
	f, err := os.Create(fn)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening CSV file %s. %v", fn, err)
	}
	var fc io.WriteCloser = f
	if filepath.Ext(fn) == ".gz" {
		w, _ := gzip.NewWriterLevel(f, gzip.BestCompression)
		fc = w
	}
	ch, done, err := NewIssueCSVWriter(fc, dedupers...)
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

type IssueDBAction func(ctx context.Context, db DB, record Issue) error

// NewIssueDBWriterSize creates a DB writer that will write each issue into the DB
func NewIssueDBWriterSize(ctx context.Context, db DB, errors chan<- error, size int, actions ...IssueDBAction) (chan Issue, chan bool, error) {
	ch := make(chan Issue, size)
	done := make(chan bool)
	var action IssueDBAction
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

// NewIssueDBWriter creates a DB writer that will write each issue into the DB
func NewIssueDBWriter(ctx context.Context, db DB, errors chan<- error, actions ...IssueDBAction) (chan Issue, chan bool, error) {
	return NewIssueDBWriterSize(ctx, db, errors, 100, actions...)
}

// IssueColumnID is the ID SQL column name for the Issue table
const IssueColumnID = "id"

// IssueEscapedColumnID is the escaped ID SQL column name for the Issue table
const IssueEscapedColumnID = "`id`"

// IssueColumnChecksum is the Checksum SQL column name for the Issue table
const IssueColumnChecksum = "checksum"

// IssueEscapedColumnChecksum is the escaped Checksum SQL column name for the Issue table
const IssueEscapedColumnChecksum = "`checksum`"

// IssueColumnUserID is the UserID SQL column name for the Issue table
const IssueColumnUserID = "user_id"

// IssueEscapedColumnUserID is the escaped UserID SQL column name for the Issue table
const IssueEscapedColumnUserID = "`user_id`"

// IssueColumnProjectID is the ProjectID SQL column name for the Issue table
const IssueColumnProjectID = "project_id"

// IssueEscapedColumnProjectID is the escaped ProjectID SQL column name for the Issue table
const IssueEscapedColumnProjectID = "`project_id`"

// IssueColumnTitle is the Title SQL column name for the Issue table
const IssueColumnTitle = "title"

// IssueEscapedColumnTitle is the escaped Title SQL column name for the Issue table
const IssueEscapedColumnTitle = "`title`"

// IssueColumnBody is the Body SQL column name for the Issue table
const IssueColumnBody = "body"

// IssueEscapedColumnBody is the escaped Body SQL column name for the Issue table
const IssueEscapedColumnBody = "`body`"

// IssueColumnState is the State SQL column name for the Issue table
const IssueColumnState = "state"

// IssueEscapedColumnState is the escaped State SQL column name for the Issue table
const IssueEscapedColumnState = "`state`"

// IssueColumnType is the Type SQL column name for the Issue table
const IssueColumnType = "type"

// IssueEscapedColumnType is the escaped Type SQL column name for the Issue table
const IssueEscapedColumnType = "`type`"

// IssueColumnResolution is the Resolution SQL column name for the Issue table
const IssueColumnResolution = "resolution"

// IssueEscapedColumnResolution is the escaped Resolution SQL column name for the Issue table
const IssueEscapedColumnResolution = "`resolution`"

// IssueColumnCreatedAt is the CreatedAt SQL column name for the Issue table
const IssueColumnCreatedAt = "created_at"

// IssueEscapedColumnCreatedAt is the escaped CreatedAt SQL column name for the Issue table
const IssueEscapedColumnCreatedAt = "`created_at`"

// IssueColumnUpdatedAt is the UpdatedAt SQL column name for the Issue table
const IssueColumnUpdatedAt = "updated_at"

// IssueEscapedColumnUpdatedAt is the escaped UpdatedAt SQL column name for the Issue table
const IssueEscapedColumnUpdatedAt = "`updated_at`"

// IssueColumnClosedAt is the ClosedAt SQL column name for the Issue table
const IssueColumnClosedAt = "closed_at"

// IssueEscapedColumnClosedAt is the escaped ClosedAt SQL column name for the Issue table
const IssueEscapedColumnClosedAt = "`closed_at`"

// IssueColumnClosedbyUserID is the ClosedbyUserID SQL column name for the Issue table
const IssueColumnClosedbyUserID = "closedby_user_id"

// IssueEscapedColumnClosedbyUserID is the escaped ClosedbyUserID SQL column name for the Issue table
const IssueEscapedColumnClosedbyUserID = "`closedby_user_id`"

// IssueColumnURL is the URL SQL column name for the Issue table
const IssueColumnURL = "url"

// IssueEscapedColumnURL is the escaped URL SQL column name for the Issue table
const IssueEscapedColumnURL = "`url`"

// IssueColumnCustomerID is the CustomerID SQL column name for the Issue table
const IssueColumnCustomerID = "customer_id"

// IssueEscapedColumnCustomerID is the escaped CustomerID SQL column name for the Issue table
const IssueEscapedColumnCustomerID = "`customer_id`"

// IssueColumnRefType is the RefType SQL column name for the Issue table
const IssueColumnRefType = "ref_type"

// IssueEscapedColumnRefType is the escaped RefType SQL column name for the Issue table
const IssueEscapedColumnRefType = "`ref_type`"

// IssueColumnRefID is the RefID SQL column name for the Issue table
const IssueColumnRefID = "ref_id"

// IssueEscapedColumnRefID is the escaped RefID SQL column name for the Issue table
const IssueEscapedColumnRefID = "`ref_id`"

// IssueColumnMetadata is the Metadata SQL column name for the Issue table
const IssueColumnMetadata = "metadata"

// IssueEscapedColumnMetadata is the escaped Metadata SQL column name for the Issue table
const IssueEscapedColumnMetadata = "`metadata`"

// GetID will return the Issue ID value
func (t *Issue) GetID() string {
	return t.ID
}

// SetID will set the Issue ID value
func (t *Issue) SetID(v string) {
	t.ID = v
}

// FindIssueByID will find a Issue by ID
func FindIssueByID(ctx context.Context, db DB, value string) (*Issue, error) {
	q := "SELECT `issue`.`id`,`issue`.`checksum`,`issue`.`user_id`,`issue`.`project_id`,`issue`.`title`,`issue`.`body`,`issue`.`state`,`issue`.`type`,`issue`.`resolution`,`issue`.`created_at`,`issue`.`updated_at`,`issue`.`closed_at`,`issue`.`closedby_user_id`,`issue`.`url`,`issue`.`customer_id`,`issue`.`ref_type`,`issue`.`ref_id`,`issue`.`metadata` FROM `issue` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _UserID sql.NullString
	var _ProjectID sql.NullString
	var _Title sql.NullString
	var _Body sql.NullString
	var _State sql.NullString
	var _Type sql.NullString
	var _Resolution sql.NullString
	var _CreatedAt sql.NullInt64
	var _UpdatedAt sql.NullInt64
	var _ClosedAt sql.NullInt64
	var _ClosedbyUserID sql.NullString
	var _URL sql.NullString
	var _CustomerID sql.NullString
	var _RefType sql.NullString
	var _RefID sql.NullString
	var _Metadata sql.NullString
	err := db.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_UserID,
		&_ProjectID,
		&_Title,
		&_Body,
		&_State,
		&_Type,
		&_Resolution,
		&_CreatedAt,
		&_UpdatedAt,
		&_ClosedAt,
		&_ClosedbyUserID,
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
	t := &Issue{}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _Checksum.Valid {
		t.SetChecksum(_Checksum.String)
	}
	if _UserID.Valid {
		t.SetUserID(_UserID.String)
	}
	if _ProjectID.Valid {
		t.SetProjectID(_ProjectID.String)
	}
	if _Title.Valid {
		t.SetTitle(_Title.String)
	}
	if _Body.Valid {
		t.SetBody(_Body.String)
	}
	if _State.Valid {
		t.SetStateString(_State.String)
	}
	if _Type.Valid {
		t.SetTypeString(_Type.String)
	}
	if _Resolution.Valid {
		t.SetResolutionString(_Resolution.String)
	}
	if _CreatedAt.Valid {
		t.SetCreatedAt(_CreatedAt.Int64)
	}
	if _UpdatedAt.Valid {
		t.SetUpdatedAt(_UpdatedAt.Int64)
	}
	if _ClosedAt.Valid {
		t.SetClosedAt(_ClosedAt.Int64)
	}
	if _ClosedbyUserID.Valid {
		t.SetClosedbyUserID(_ClosedbyUserID.String)
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

// FindIssueByIDTx will find a Issue by ID using the provided transaction
func FindIssueByIDTx(ctx context.Context, tx Tx, value string) (*Issue, error) {
	q := "SELECT `issue`.`id`,`issue`.`checksum`,`issue`.`user_id`,`issue`.`project_id`,`issue`.`title`,`issue`.`body`,`issue`.`state`,`issue`.`type`,`issue`.`resolution`,`issue`.`created_at`,`issue`.`updated_at`,`issue`.`closed_at`,`issue`.`closedby_user_id`,`issue`.`url`,`issue`.`customer_id`,`issue`.`ref_type`,`issue`.`ref_id`,`issue`.`metadata` FROM `issue` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _UserID sql.NullString
	var _ProjectID sql.NullString
	var _Title sql.NullString
	var _Body sql.NullString
	var _State sql.NullString
	var _Type sql.NullString
	var _Resolution sql.NullString
	var _CreatedAt sql.NullInt64
	var _UpdatedAt sql.NullInt64
	var _ClosedAt sql.NullInt64
	var _ClosedbyUserID sql.NullString
	var _URL sql.NullString
	var _CustomerID sql.NullString
	var _RefType sql.NullString
	var _RefID sql.NullString
	var _Metadata sql.NullString
	err := tx.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_UserID,
		&_ProjectID,
		&_Title,
		&_Body,
		&_State,
		&_Type,
		&_Resolution,
		&_CreatedAt,
		&_UpdatedAt,
		&_ClosedAt,
		&_ClosedbyUserID,
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
	t := &Issue{}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _Checksum.Valid {
		t.SetChecksum(_Checksum.String)
	}
	if _UserID.Valid {
		t.SetUserID(_UserID.String)
	}
	if _ProjectID.Valid {
		t.SetProjectID(_ProjectID.String)
	}
	if _Title.Valid {
		t.SetTitle(_Title.String)
	}
	if _Body.Valid {
		t.SetBody(_Body.String)
	}
	if _State.Valid {
		t.SetStateString(_State.String)
	}
	if _Type.Valid {
		t.SetTypeString(_Type.String)
	}
	if _Resolution.Valid {
		t.SetResolutionString(_Resolution.String)
	}
	if _CreatedAt.Valid {
		t.SetCreatedAt(_CreatedAt.Int64)
	}
	if _UpdatedAt.Valid {
		t.SetUpdatedAt(_UpdatedAt.Int64)
	}
	if _ClosedAt.Valid {
		t.SetClosedAt(_ClosedAt.Int64)
	}
	if _ClosedbyUserID.Valid {
		t.SetClosedbyUserID(_ClosedbyUserID.String)
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

// GetChecksum will return the Issue Checksum value
func (t *Issue) GetChecksum() string {
	if t.Checksum == nil {
		return ""
	}
	return *t.Checksum
}

// SetChecksum will set the Issue Checksum value
func (t *Issue) SetChecksum(v string) {
	t.Checksum = &v
}

// GetUserID will return the Issue UserID value
func (t *Issue) GetUserID() string {
	if t.UserID == nil {
		return ""
	}
	return *t.UserID
}

// SetUserID will set the Issue UserID value
func (t *Issue) SetUserID(v string) {
	t.UserID = &v
}

// GetProjectID will return the Issue ProjectID value
func (t *Issue) GetProjectID() string {
	return t.ProjectID
}

// SetProjectID will set the Issue ProjectID value
func (t *Issue) SetProjectID(v string) {
	t.ProjectID = v
}

// GetTitle will return the Issue Title value
func (t *Issue) GetTitle() string {
	return t.Title
}

// SetTitle will set the Issue Title value
func (t *Issue) SetTitle(v string) {
	t.Title = v
}

// GetBody will return the Issue Body value
func (t *Issue) GetBody() string {
	if t.Body == nil {
		return ""
	}
	return *t.Body
}

// SetBody will set the Issue Body value
func (t *Issue) SetBody(v string) {
	t.Body = &v
}

// GetState will return the Issue State value
func (t *Issue) GetState() Issue_IssueState {
	return t.State
}

// SetState will set the Issue State value
func (t *Issue) SetState(v Issue_IssueState) {
	t.State = v
}

// GetStateString will return the Issue State value as a string
func (t *Issue) GetStateString() string {
	return t.State.String()
}

// SetStateString will set the Issue State value from a string
func (t *Issue) SetStateString(v string) {
	var _State = toIssueState(v)
	if _State != nil {
		t.State = *_State
	}
}

// GetType will return the Issue Type value
func (t *Issue) GetType() Issue_IssueType {
	return t.Type
}

// SetType will set the Issue Type value
func (t *Issue) SetType(v Issue_IssueType) {
	t.Type = v
}

// GetTypeString will return the Issue Type value as a string
func (t *Issue) GetTypeString() string {
	return t.Type.String()
}

// SetTypeString will set the Issue Type value from a string
func (t *Issue) SetTypeString(v string) {
	var _Type = toIssueType(v)
	if _Type != nil {
		t.Type = *_Type
	}
}

// GetResolution will return the Issue Resolution value
func (t *Issue) GetResolution() Issue_IssueResolution {
	return t.Resolution
}

// SetResolution will set the Issue Resolution value
func (t *Issue) SetResolution(v Issue_IssueResolution) {
	t.Resolution = v
}

// GetResolutionString will return the Issue Resolution value as a string
func (t *Issue) GetResolutionString() string {
	return t.Resolution.String()
}

// SetResolutionString will set the Issue Resolution value from a string
func (t *Issue) SetResolutionString(v string) {
	var _Resolution = toIssueResolution(v)
	if _Resolution != nil {
		t.Resolution = *_Resolution
	}
}

// GetCreatedAt will return the Issue CreatedAt value
func (t *Issue) GetCreatedAt() int64 {
	return t.CreatedAt
}

// SetCreatedAt will set the Issue CreatedAt value
func (t *Issue) SetCreatedAt(v int64) {
	t.CreatedAt = v
}

// GetUpdatedAt will return the Issue UpdatedAt value
func (t *Issue) GetUpdatedAt() int64 {
	if t.UpdatedAt == nil {
		return int64(0)
	}
	return *t.UpdatedAt
}

// SetUpdatedAt will set the Issue UpdatedAt value
func (t *Issue) SetUpdatedAt(v int64) {
	t.UpdatedAt = &v
}

// GetClosedAt will return the Issue ClosedAt value
func (t *Issue) GetClosedAt() int64 {
	if t.ClosedAt == nil {
		return int64(0)
	}
	return *t.ClosedAt
}

// SetClosedAt will set the Issue ClosedAt value
func (t *Issue) SetClosedAt(v int64) {
	t.ClosedAt = &v
}

// GetClosedbyUserID will return the Issue ClosedbyUserID value
func (t *Issue) GetClosedbyUserID() string {
	if t.ClosedbyUserID == nil {
		return ""
	}
	return *t.ClosedbyUserID
}

// SetClosedbyUserID will set the Issue ClosedbyUserID value
func (t *Issue) SetClosedbyUserID(v string) {
	t.ClosedbyUserID = &v
}

// GetURL will return the Issue URL value
func (t *Issue) GetURL() string {
	return t.URL
}

// SetURL will set the Issue URL value
func (t *Issue) SetURL(v string) {
	t.URL = v
}

// GetCustomerID will return the Issue CustomerID value
func (t *Issue) GetCustomerID() string {
	return t.CustomerID
}

// SetCustomerID will set the Issue CustomerID value
func (t *Issue) SetCustomerID(v string) {
	t.CustomerID = v
}

// FindIssuesByCustomerID will find all Issues by the CustomerID value
func FindIssuesByCustomerID(ctx context.Context, db DB, value string) ([]*Issue, error) {
	q := "SELECT `issue`.`id`,`issue`.`checksum`,`issue`.`user_id`,`issue`.`project_id`,`issue`.`title`,`issue`.`body`,`issue`.`state`,`issue`.`type`,`issue`.`resolution`,`issue`.`created_at`,`issue`.`updated_at`,`issue`.`closed_at`,`issue`.`closedby_user_id`,`issue`.`url`,`issue`.`customer_id`,`issue`.`ref_type`,`issue`.`ref_id`,`issue`.`metadata` FROM `issue` WHERE `customer_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*Issue, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _UserID sql.NullString
		var _ProjectID sql.NullString
		var _Title sql.NullString
		var _Body sql.NullString
		var _State sql.NullString
		var _Type sql.NullString
		var _Resolution sql.NullString
		var _CreatedAt sql.NullInt64
		var _UpdatedAt sql.NullInt64
		var _ClosedAt sql.NullInt64
		var _ClosedbyUserID sql.NullString
		var _URL sql.NullString
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefID sql.NullString
		var _Metadata sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_UserID,
			&_ProjectID,
			&_Title,
			&_Body,
			&_State,
			&_Type,
			&_Resolution,
			&_CreatedAt,
			&_UpdatedAt,
			&_ClosedAt,
			&_ClosedbyUserID,
			&_URL,
			&_CustomerID,
			&_RefType,
			&_RefID,
			&_Metadata,
		)
		if err != nil {
			return nil, err
		}
		t := &Issue{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _Title.Valid {
			t.SetTitle(_Title.String)
		}
		if _Body.Valid {
			t.SetBody(_Body.String)
		}
		if _State.Valid {
			t.SetStateString(_State.String)
		}
		if _Type.Valid {
			t.SetTypeString(_Type.String)
		}
		if _Resolution.Valid {
			t.SetResolutionString(_Resolution.String)
		}
		if _CreatedAt.Valid {
			t.SetCreatedAt(_CreatedAt.Int64)
		}
		if _UpdatedAt.Valid {
			t.SetUpdatedAt(_UpdatedAt.Int64)
		}
		if _ClosedAt.Valid {
			t.SetClosedAt(_ClosedAt.Int64)
		}
		if _ClosedbyUserID.Valid {
			t.SetClosedbyUserID(_ClosedbyUserID.String)
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

// FindIssuesByCustomerIDTx will find all Issues by the CustomerID value using the provided transaction
func FindIssuesByCustomerIDTx(ctx context.Context, tx Tx, value string) ([]*Issue, error) {
	q := "SELECT `issue`.`id`,`issue`.`checksum`,`issue`.`user_id`,`issue`.`project_id`,`issue`.`title`,`issue`.`body`,`issue`.`state`,`issue`.`type`,`issue`.`resolution`,`issue`.`created_at`,`issue`.`updated_at`,`issue`.`closed_at`,`issue`.`closedby_user_id`,`issue`.`url`,`issue`.`customer_id`,`issue`.`ref_type`,`issue`.`ref_id`,`issue`.`metadata` FROM `issue` WHERE `customer_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*Issue, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _UserID sql.NullString
		var _ProjectID sql.NullString
		var _Title sql.NullString
		var _Body sql.NullString
		var _State sql.NullString
		var _Type sql.NullString
		var _Resolution sql.NullString
		var _CreatedAt sql.NullInt64
		var _UpdatedAt sql.NullInt64
		var _ClosedAt sql.NullInt64
		var _ClosedbyUserID sql.NullString
		var _URL sql.NullString
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefID sql.NullString
		var _Metadata sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_UserID,
			&_ProjectID,
			&_Title,
			&_Body,
			&_State,
			&_Type,
			&_Resolution,
			&_CreatedAt,
			&_UpdatedAt,
			&_ClosedAt,
			&_ClosedbyUserID,
			&_URL,
			&_CustomerID,
			&_RefType,
			&_RefID,
			&_Metadata,
		)
		if err != nil {
			return nil, err
		}
		t := &Issue{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _Title.Valid {
			t.SetTitle(_Title.String)
		}
		if _Body.Valid {
			t.SetBody(_Body.String)
		}
		if _State.Valid {
			t.SetStateString(_State.String)
		}
		if _Type.Valid {
			t.SetTypeString(_Type.String)
		}
		if _Resolution.Valid {
			t.SetResolutionString(_Resolution.String)
		}
		if _CreatedAt.Valid {
			t.SetCreatedAt(_CreatedAt.Int64)
		}
		if _UpdatedAt.Valid {
			t.SetUpdatedAt(_UpdatedAt.Int64)
		}
		if _ClosedAt.Valid {
			t.SetClosedAt(_ClosedAt.Int64)
		}
		if _ClosedbyUserID.Valid {
			t.SetClosedbyUserID(_ClosedbyUserID.String)
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

// GetRefType will return the Issue RefType value
func (t *Issue) GetRefType() string {
	return t.RefType
}

// SetRefType will set the Issue RefType value
func (t *Issue) SetRefType(v string) {
	t.RefType = v
}

// FindIssuesByRefType will find all Issues by the RefType value
func FindIssuesByRefType(ctx context.Context, db DB, value string) ([]*Issue, error) {
	q := "SELECT `issue`.`id`,`issue`.`checksum`,`issue`.`user_id`,`issue`.`project_id`,`issue`.`title`,`issue`.`body`,`issue`.`state`,`issue`.`type`,`issue`.`resolution`,`issue`.`created_at`,`issue`.`updated_at`,`issue`.`closed_at`,`issue`.`closedby_user_id`,`issue`.`url`,`issue`.`customer_id`,`issue`.`ref_type`,`issue`.`ref_id`,`issue`.`metadata` FROM `issue` WHERE `ref_type` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*Issue, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _UserID sql.NullString
		var _ProjectID sql.NullString
		var _Title sql.NullString
		var _Body sql.NullString
		var _State sql.NullString
		var _Type sql.NullString
		var _Resolution sql.NullString
		var _CreatedAt sql.NullInt64
		var _UpdatedAt sql.NullInt64
		var _ClosedAt sql.NullInt64
		var _ClosedbyUserID sql.NullString
		var _URL sql.NullString
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefID sql.NullString
		var _Metadata sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_UserID,
			&_ProjectID,
			&_Title,
			&_Body,
			&_State,
			&_Type,
			&_Resolution,
			&_CreatedAt,
			&_UpdatedAt,
			&_ClosedAt,
			&_ClosedbyUserID,
			&_URL,
			&_CustomerID,
			&_RefType,
			&_RefID,
			&_Metadata,
		)
		if err != nil {
			return nil, err
		}
		t := &Issue{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _Title.Valid {
			t.SetTitle(_Title.String)
		}
		if _Body.Valid {
			t.SetBody(_Body.String)
		}
		if _State.Valid {
			t.SetStateString(_State.String)
		}
		if _Type.Valid {
			t.SetTypeString(_Type.String)
		}
		if _Resolution.Valid {
			t.SetResolutionString(_Resolution.String)
		}
		if _CreatedAt.Valid {
			t.SetCreatedAt(_CreatedAt.Int64)
		}
		if _UpdatedAt.Valid {
			t.SetUpdatedAt(_UpdatedAt.Int64)
		}
		if _ClosedAt.Valid {
			t.SetClosedAt(_ClosedAt.Int64)
		}
		if _ClosedbyUserID.Valid {
			t.SetClosedbyUserID(_ClosedbyUserID.String)
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

// FindIssuesByRefTypeTx will find all Issues by the RefType value using the provided transaction
func FindIssuesByRefTypeTx(ctx context.Context, tx Tx, value string) ([]*Issue, error) {
	q := "SELECT `issue`.`id`,`issue`.`checksum`,`issue`.`user_id`,`issue`.`project_id`,`issue`.`title`,`issue`.`body`,`issue`.`state`,`issue`.`type`,`issue`.`resolution`,`issue`.`created_at`,`issue`.`updated_at`,`issue`.`closed_at`,`issue`.`closedby_user_id`,`issue`.`url`,`issue`.`customer_id`,`issue`.`ref_type`,`issue`.`ref_id`,`issue`.`metadata` FROM `issue` WHERE `ref_type` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*Issue, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _UserID sql.NullString
		var _ProjectID sql.NullString
		var _Title sql.NullString
		var _Body sql.NullString
		var _State sql.NullString
		var _Type sql.NullString
		var _Resolution sql.NullString
		var _CreatedAt sql.NullInt64
		var _UpdatedAt sql.NullInt64
		var _ClosedAt sql.NullInt64
		var _ClosedbyUserID sql.NullString
		var _URL sql.NullString
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefID sql.NullString
		var _Metadata sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_UserID,
			&_ProjectID,
			&_Title,
			&_Body,
			&_State,
			&_Type,
			&_Resolution,
			&_CreatedAt,
			&_UpdatedAt,
			&_ClosedAt,
			&_ClosedbyUserID,
			&_URL,
			&_CustomerID,
			&_RefType,
			&_RefID,
			&_Metadata,
		)
		if err != nil {
			return nil, err
		}
		t := &Issue{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _Title.Valid {
			t.SetTitle(_Title.String)
		}
		if _Body.Valid {
			t.SetBody(_Body.String)
		}
		if _State.Valid {
			t.SetStateString(_State.String)
		}
		if _Type.Valid {
			t.SetTypeString(_Type.String)
		}
		if _Resolution.Valid {
			t.SetResolutionString(_Resolution.String)
		}
		if _CreatedAt.Valid {
			t.SetCreatedAt(_CreatedAt.Int64)
		}
		if _UpdatedAt.Valid {
			t.SetUpdatedAt(_UpdatedAt.Int64)
		}
		if _ClosedAt.Valid {
			t.SetClosedAt(_ClosedAt.Int64)
		}
		if _ClosedbyUserID.Valid {
			t.SetClosedbyUserID(_ClosedbyUserID.String)
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

// GetRefID will return the Issue RefID value
func (t *Issue) GetRefID() string {
	return t.RefID
}

// SetRefID will set the Issue RefID value
func (t *Issue) SetRefID(v string) {
	t.RefID = v
}

// FindIssuesByRefID will find all Issues by the RefID value
func FindIssuesByRefID(ctx context.Context, db DB, value string) ([]*Issue, error) {
	q := "SELECT `issue`.`id`,`issue`.`checksum`,`issue`.`user_id`,`issue`.`project_id`,`issue`.`title`,`issue`.`body`,`issue`.`state`,`issue`.`type`,`issue`.`resolution`,`issue`.`created_at`,`issue`.`updated_at`,`issue`.`closed_at`,`issue`.`closedby_user_id`,`issue`.`url`,`issue`.`customer_id`,`issue`.`ref_type`,`issue`.`ref_id`,`issue`.`metadata` FROM `issue` WHERE `ref_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*Issue, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _UserID sql.NullString
		var _ProjectID sql.NullString
		var _Title sql.NullString
		var _Body sql.NullString
		var _State sql.NullString
		var _Type sql.NullString
		var _Resolution sql.NullString
		var _CreatedAt sql.NullInt64
		var _UpdatedAt sql.NullInt64
		var _ClosedAt sql.NullInt64
		var _ClosedbyUserID sql.NullString
		var _URL sql.NullString
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefID sql.NullString
		var _Metadata sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_UserID,
			&_ProjectID,
			&_Title,
			&_Body,
			&_State,
			&_Type,
			&_Resolution,
			&_CreatedAt,
			&_UpdatedAt,
			&_ClosedAt,
			&_ClosedbyUserID,
			&_URL,
			&_CustomerID,
			&_RefType,
			&_RefID,
			&_Metadata,
		)
		if err != nil {
			return nil, err
		}
		t := &Issue{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _Title.Valid {
			t.SetTitle(_Title.String)
		}
		if _Body.Valid {
			t.SetBody(_Body.String)
		}
		if _State.Valid {
			t.SetStateString(_State.String)
		}
		if _Type.Valid {
			t.SetTypeString(_Type.String)
		}
		if _Resolution.Valid {
			t.SetResolutionString(_Resolution.String)
		}
		if _CreatedAt.Valid {
			t.SetCreatedAt(_CreatedAt.Int64)
		}
		if _UpdatedAt.Valid {
			t.SetUpdatedAt(_UpdatedAt.Int64)
		}
		if _ClosedAt.Valid {
			t.SetClosedAt(_ClosedAt.Int64)
		}
		if _ClosedbyUserID.Valid {
			t.SetClosedbyUserID(_ClosedbyUserID.String)
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

// FindIssuesByRefIDTx will find all Issues by the RefID value using the provided transaction
func FindIssuesByRefIDTx(ctx context.Context, tx Tx, value string) ([]*Issue, error) {
	q := "SELECT `issue`.`id`,`issue`.`checksum`,`issue`.`user_id`,`issue`.`project_id`,`issue`.`title`,`issue`.`body`,`issue`.`state`,`issue`.`type`,`issue`.`resolution`,`issue`.`created_at`,`issue`.`updated_at`,`issue`.`closed_at`,`issue`.`closedby_user_id`,`issue`.`url`,`issue`.`customer_id`,`issue`.`ref_type`,`issue`.`ref_id`,`issue`.`metadata` FROM `issue` WHERE `ref_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*Issue, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _UserID sql.NullString
		var _ProjectID sql.NullString
		var _Title sql.NullString
		var _Body sql.NullString
		var _State sql.NullString
		var _Type sql.NullString
		var _Resolution sql.NullString
		var _CreatedAt sql.NullInt64
		var _UpdatedAt sql.NullInt64
		var _ClosedAt sql.NullInt64
		var _ClosedbyUserID sql.NullString
		var _URL sql.NullString
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefID sql.NullString
		var _Metadata sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_UserID,
			&_ProjectID,
			&_Title,
			&_Body,
			&_State,
			&_Type,
			&_Resolution,
			&_CreatedAt,
			&_UpdatedAt,
			&_ClosedAt,
			&_ClosedbyUserID,
			&_URL,
			&_CustomerID,
			&_RefType,
			&_RefID,
			&_Metadata,
		)
		if err != nil {
			return nil, err
		}
		t := &Issue{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _Title.Valid {
			t.SetTitle(_Title.String)
		}
		if _Body.Valid {
			t.SetBody(_Body.String)
		}
		if _State.Valid {
			t.SetStateString(_State.String)
		}
		if _Type.Valid {
			t.SetTypeString(_Type.String)
		}
		if _Resolution.Valid {
			t.SetResolutionString(_Resolution.String)
		}
		if _CreatedAt.Valid {
			t.SetCreatedAt(_CreatedAt.Int64)
		}
		if _UpdatedAt.Valid {
			t.SetUpdatedAt(_UpdatedAt.Int64)
		}
		if _ClosedAt.Valid {
			t.SetClosedAt(_ClosedAt.Int64)
		}
		if _ClosedbyUserID.Valid {
			t.SetClosedbyUserID(_ClosedbyUserID.String)
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

// GetMetadata will return the Issue Metadata value
func (t *Issue) GetMetadata() string {
	if t.Metadata == nil {
		return ""
	}
	return *t.Metadata
}

// SetMetadata will set the Issue Metadata value
func (t *Issue) SetMetadata(v string) {
	t.Metadata = &v
}

func (t *Issue) toTimestamp(value time.Time) *timestamp.Timestamp {
	ts, _ := ptypes.TimestampProto(value)
	return ts
}

// DBCreateIssueTable will create the Issue table
func DBCreateIssueTable(ctx context.Context, db DB) error {
	q := "CREATE TABLE `issue` (`id`VARCHAR(64) NOT NULL PRIMARY KEY,`checksum`CHAR(64),`user_id` VARCHAR(64),`project_id` VARCHAR(64) NOT NULL,`title`TEXT NOT NULL,`body` LONGTEXT,`state`ENUM('open','closed') NOT NULL,`type` ENUM('bug','feature','enhancement','other') NOT NULL,`resolution` ENUM('none','resolved','unresolved') NOT NULL,`created_at` BIGINT UNSIGNED NOT NULL,`updated_at` BIGINT UNSIGNED,`closed_at` BIGINT UNSIGNED,`closedby_user_id` VARCHAR(64),`url` VARCHAR(255) NOT NULL,`customer_id`VARCHAR(64) NOT NULL,`ref_type`VARCHAR(20) NOT NULL,`ref_id` VARCHAR(64) NOT NULL,`metadata`JSON,INDEX issue_customer_id_index (`customer_id`),INDEX issue_ref_type_index (`ref_type`),INDEX issue_ref_id_index (`ref_id`),INDEX issue_customer_id_user_id_project_id_index (`customer_id`,`user_id`,`project_id`),INDEX issue_state_created_at_customer_id_index (`state`,`created_at`,`customer_id`),INDEX issue_type_customer_id_project_id_index (`type`,`customer_id`,`project_id`),INDEX issue_ref_type_state_project_id_customer_id_index (`ref_type`,`state`,`project_id`,`customer_id`),INDEX issue_ref_type_state_closed_at_project_id_customer_id_index (`ref_type`,`state`,`closed_at`,`project_id`,`customer_id`),INDEX issue_ref_type_state_created_at_project_id_customer_id_index (`ref_type`,`state`,`created_at`,`project_id`,`customer_id`),INDEX issue_state_customer_id_project_id_index (`state`,`customer_id`,`project_id`),INDEX issue_state_ref_id_closed_at_index (`state`,`ref_id`,`closed_at`),INDEX issue_state_customer_id_closed_at_ref_id_type_index (`state`,`customer_id`,`closed_at`,`ref_id`,`type`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBCreateIssueTableTx will create the Issue table using the provided transction
func DBCreateIssueTableTx(ctx context.Context, tx Tx) error {
	q := "CREATE TABLE `issue` (`id`VARCHAR(64) NOT NULL PRIMARY KEY,`checksum`CHAR(64),`user_id` VARCHAR(64),`project_id` VARCHAR(64) NOT NULL,`title`TEXT NOT NULL,`body` LONGTEXT,`state`ENUM('open','closed') NOT NULL,`type` ENUM('bug','feature','enhancement','other') NOT NULL,`resolution` ENUM('none','resolved','unresolved') NOT NULL,`created_at` BIGINT UNSIGNED NOT NULL,`updated_at` BIGINT UNSIGNED,`closed_at` BIGINT UNSIGNED,`closedby_user_id` VARCHAR(64),`url` VARCHAR(255) NOT NULL,`customer_id`VARCHAR(64) NOT NULL,`ref_type`VARCHAR(20) NOT NULL,`ref_id` VARCHAR(64) NOT NULL,`metadata`JSON,INDEX issue_customer_id_index (`customer_id`),INDEX issue_ref_type_index (`ref_type`),INDEX issue_ref_id_index (`ref_id`),INDEX issue_customer_id_user_id_project_id_index (`customer_id`,`user_id`,`project_id`),INDEX issue_state_created_at_customer_id_index (`state`,`created_at`,`customer_id`),INDEX issue_type_customer_id_project_id_index (`type`,`customer_id`,`project_id`),INDEX issue_ref_type_state_project_id_customer_id_index (`ref_type`,`state`,`project_id`,`customer_id`),INDEX issue_ref_type_state_closed_at_project_id_customer_id_index (`ref_type`,`state`,`closed_at`,`project_id`,`customer_id`),INDEX issue_ref_type_state_created_at_project_id_customer_id_index (`ref_type`,`state`,`created_at`,`project_id`,`customer_id`),INDEX issue_state_customer_id_project_id_index (`state`,`customer_id`,`project_id`),INDEX issue_state_ref_id_closed_at_index (`state`,`ref_id`,`closed_at`),INDEX issue_state_customer_id_closed_at_ref_id_type_index (`state`,`customer_id`,`closed_at`,`ref_id`,`type`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// DBDropIssueTable will drop the Issue table
func DBDropIssueTable(ctx context.Context, db DB) error {
	q := "DROP TABLE IF EXISTS `issue`"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBDropIssueTableTx will drop the Issue table using the provided transaction
func DBDropIssueTableTx(ctx context.Context, tx Tx) error {
	q := "DROP TABLE IF EXISTS `issue`"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// CalculateChecksum will calculate a checksum of the SHA1 of all field values
func (t *Issue) CalculateChecksum() string {
	return orm.HashStrings(
		orm.ToString(t.ID),
		orm.ToString(t.UserID),
		orm.ToString(t.ProjectID),
		orm.ToString(t.Title),
		orm.ToString(t.Body),
		enumIssueStateToString(&t.State),
		enumIssueTypeToString(&t.Type),
		enumIssueResolutionToString(&t.Resolution),
		orm.ToString(t.CreatedAt),
		orm.ToString(t.UpdatedAt),
		orm.ToString(t.ClosedAt),
		orm.ToString(t.ClosedbyUserID),
		orm.ToString(t.URL),
		orm.ToString(t.CustomerID),
		orm.ToString(t.RefType),
		orm.ToString(t.RefID),
		orm.ToString(t.Metadata),
	)
}

// DBCreate will create a new Issue record in the database
func (t *Issue) DBCreate(ctx context.Context, db DB) (sql.Result, error) {
	q := "INSERT INTO `issue` (`issue`.`id`,`issue`.`checksum`,`issue`.`user_id`,`issue`.`project_id`,`issue`.`title`,`issue`.`body`,`issue`.`state`,`issue`.`type`,`issue`.`resolution`,`issue`.`created_at`,`issue`.`updated_at`,`issue`.`closed_at`,`issue`.`closedby_user_id`,`issue`.`url`,`issue`.`customer_id`,`issue`.`ref_type`,`issue`.`ref_id`,`issue`.`metadata`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.UserID),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLString(t.Title),
		orm.ToSQLString(t.Body),
		orm.ToSQLString(enumIssueStateToString(&t.State)),
		orm.ToSQLString(enumIssueTypeToString(&t.Type)),
		orm.ToSQLString(enumIssueResolutionToString(&t.Resolution)),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
		orm.ToSQLInt64(t.ClosedAt),
		orm.ToSQLString(t.ClosedbyUserID),
		orm.ToSQLString(t.URL),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefType),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.Metadata),
	)
}

// DBCreateTx will create a new Issue record in the database using the provided transaction
func (t *Issue) DBCreateTx(ctx context.Context, tx Tx) (sql.Result, error) {
	q := "INSERT INTO `issue` (`issue`.`id`,`issue`.`checksum`,`issue`.`user_id`,`issue`.`project_id`,`issue`.`title`,`issue`.`body`,`issue`.`state`,`issue`.`type`,`issue`.`resolution`,`issue`.`created_at`,`issue`.`updated_at`,`issue`.`closed_at`,`issue`.`closedby_user_id`,`issue`.`url`,`issue`.`customer_id`,`issue`.`ref_type`,`issue`.`ref_id`,`issue`.`metadata`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.UserID),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLString(t.Title),
		orm.ToSQLString(t.Body),
		orm.ToSQLString(enumIssueStateToString(&t.State)),
		orm.ToSQLString(enumIssueTypeToString(&t.Type)),
		orm.ToSQLString(enumIssueResolutionToString(&t.Resolution)),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
		orm.ToSQLInt64(t.ClosedAt),
		orm.ToSQLString(t.ClosedbyUserID),
		orm.ToSQLString(t.URL),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefType),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.Metadata),
	)
}

// DBCreateIgnoreDuplicate will upsert the Issue record in the database
func (t *Issue) DBCreateIgnoreDuplicate(ctx context.Context, db DB) (sql.Result, error) {
	q := "INSERT INTO `issue` (`issue`.`id`,`issue`.`checksum`,`issue`.`user_id`,`issue`.`project_id`,`issue`.`title`,`issue`.`body`,`issue`.`state`,`issue`.`type`,`issue`.`resolution`,`issue`.`created_at`,`issue`.`updated_at`,`issue`.`closed_at`,`issue`.`closedby_user_id`,`issue`.`url`,`issue`.`customer_id`,`issue`.`ref_type`,`issue`.`ref_id`,`issue`.`metadata`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.UserID),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLString(t.Title),
		orm.ToSQLString(t.Body),
		orm.ToSQLString(enumIssueStateToString(&t.State)),
		orm.ToSQLString(enumIssueTypeToString(&t.Type)),
		orm.ToSQLString(enumIssueResolutionToString(&t.Resolution)),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
		orm.ToSQLInt64(t.ClosedAt),
		orm.ToSQLString(t.ClosedbyUserID),
		orm.ToSQLString(t.URL),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefType),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.Metadata),
	)
}

// DBCreateIgnoreDuplicateTx will upsert the Issue record in the database using the provided transaction
func (t *Issue) DBCreateIgnoreDuplicateTx(ctx context.Context, tx Tx) (sql.Result, error) {
	q := "INSERT INTO `issue` (`issue`.`id`,`issue`.`checksum`,`issue`.`user_id`,`issue`.`project_id`,`issue`.`title`,`issue`.`body`,`issue`.`state`,`issue`.`type`,`issue`.`resolution`,`issue`.`created_at`,`issue`.`updated_at`,`issue`.`closed_at`,`issue`.`closedby_user_id`,`issue`.`url`,`issue`.`customer_id`,`issue`.`ref_type`,`issue`.`ref_id`,`issue`.`metadata`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.UserID),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLString(t.Title),
		orm.ToSQLString(t.Body),
		orm.ToSQLString(enumIssueStateToString(&t.State)),
		orm.ToSQLString(enumIssueTypeToString(&t.Type)),
		orm.ToSQLString(enumIssueResolutionToString(&t.Resolution)),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
		orm.ToSQLInt64(t.ClosedAt),
		orm.ToSQLString(t.ClosedbyUserID),
		orm.ToSQLString(t.URL),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefType),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.Metadata),
	)
}

// DeleteAllIssues deletes all Issue records in the database with optional filters
func DeleteAllIssues(ctx context.Context, db DB, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(IssueTableName),
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

// DeleteAllIssuesTx deletes all Issue records in the database with optional filters using the provided transaction
func DeleteAllIssuesTx(ctx context.Context, tx Tx, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(IssueTableName),
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

// DBDelete will delete this Issue record in the database
func (t *Issue) DBDelete(ctx context.Context, db DB) (bool, error) {
	q := "DELETE FROM `issue` WHERE `id` = ?"
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

// DBDeleteTx will delete this Issue record in the database using the provided transaction
func (t *Issue) DBDeleteTx(ctx context.Context, tx Tx) (bool, error) {
	q := "DELETE FROM `issue` WHERE `id` = ?"
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

// DBUpdate will update the Issue record in the database
func (t *Issue) DBUpdate(ctx context.Context, db DB) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `issue` SET `checksum`=?,`user_id`=?,`project_id`=?,`title`=?,`body`=?,`state`=?,`type`=?,`resolution`=?,`created_at`=?,`updated_at`=?,`closed_at`=?,`closedby_user_id`=?,`url`=?,`customer_id`=?,`ref_type`=?,`ref_id`=?,`metadata`=? WHERE `id`=?"
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.UserID),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLString(t.Title),
		orm.ToSQLString(t.Body),
		orm.ToSQLString(enumIssueStateToString(&t.State)),
		orm.ToSQLString(enumIssueTypeToString(&t.Type)),
		orm.ToSQLString(enumIssueResolutionToString(&t.Resolution)),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
		orm.ToSQLInt64(t.ClosedAt),
		orm.ToSQLString(t.ClosedbyUserID),
		orm.ToSQLString(t.URL),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefType),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.Metadata),
		orm.ToSQLString(t.ID),
	)
}

// DBUpdateTx will update the Issue record in the database using the provided transaction
func (t *Issue) DBUpdateTx(ctx context.Context, tx Tx) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `issue` SET `checksum`=?,`user_id`=?,`project_id`=?,`title`=?,`body`=?,`state`=?,`type`=?,`resolution`=?,`created_at`=?,`updated_at`=?,`closed_at`=?,`closedby_user_id`=?,`url`=?,`customer_id`=?,`ref_type`=?,`ref_id`=?,`metadata`=? WHERE `id`=?"
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.UserID),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLString(t.Title),
		orm.ToSQLString(t.Body),
		orm.ToSQLString(enumIssueStateToString(&t.State)),
		orm.ToSQLString(enumIssueTypeToString(&t.Type)),
		orm.ToSQLString(enumIssueResolutionToString(&t.Resolution)),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
		orm.ToSQLInt64(t.ClosedAt),
		orm.ToSQLString(t.ClosedbyUserID),
		orm.ToSQLString(t.URL),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefType),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.Metadata),
		orm.ToSQLString(t.ID),
	)
}

// DBUpsert will upsert the Issue record in the database
func (t *Issue) DBUpsert(ctx context.Context, db DB, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `issue` (`issue`.`id`,`issue`.`checksum`,`issue`.`user_id`,`issue`.`project_id`,`issue`.`title`,`issue`.`body`,`issue`.`state`,`issue`.`type`,`issue`.`resolution`,`issue`.`created_at`,`issue`.`updated_at`,`issue`.`closed_at`,`issue`.`closedby_user_id`,`issue`.`url`,`issue`.`customer_id`,`issue`.`ref_type`,`issue`.`ref_id`,`issue`.`metadata`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `issue` (`issue`.`id`,`issue`.`checksum`,`issue`.`user_id`,`issue`.`project_id`,`issue`.`title`,`issue`.`body`,`issue`.`state`,`issue`.`type`,`issue`.`resolution`,`issue`.`created_at`,`issue`.`updated_at`,`issue`.`closed_at`,`issue`.`closedby_user_id`,`issue`.`url`,`issue`.`customer_id`,`issue`.`ref_type`,`issue`.`ref_id`,`issue`.`metadata`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`user_id`=VALUES(`user_id`),`project_id`=VALUES(`project_id`),`title`=VALUES(`title`),`body`=VALUES(`body`),`state`=VALUES(`state`),`type`=VALUES(`type`),`resolution`=VALUES(`resolution`),`created_at`=VALUES(`created_at`),`updated_at`=VALUES(`updated_at`),`closed_at`=VALUES(`closed_at`),`closedby_user_id`=VALUES(`closedby_user_id`),`url`=VALUES(`url`),`customer_id`=VALUES(`customer_id`),`ref_type`=VALUES(`ref_type`),`ref_id`=VALUES(`ref_id`),`metadata`=VALUES(`metadata`)"
	}
	r, err := db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.UserID),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLString(t.Title),
		orm.ToSQLString(t.Body),
		orm.ToSQLString(enumIssueStateToString(&t.State)),
		orm.ToSQLString(enumIssueTypeToString(&t.Type)),
		orm.ToSQLString(enumIssueResolutionToString(&t.Resolution)),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
		orm.ToSQLInt64(t.ClosedAt),
		orm.ToSQLString(t.ClosedbyUserID),
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

// DBUpsertTx will upsert the Issue record in the database using the provided transaction
func (t *Issue) DBUpsertTx(ctx context.Context, tx Tx, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `issue` (`issue`.`id`,`issue`.`checksum`,`issue`.`user_id`,`issue`.`project_id`,`issue`.`title`,`issue`.`body`,`issue`.`state`,`issue`.`type`,`issue`.`resolution`,`issue`.`created_at`,`issue`.`updated_at`,`issue`.`closed_at`,`issue`.`closedby_user_id`,`issue`.`url`,`issue`.`customer_id`,`issue`.`ref_type`,`issue`.`ref_id`,`issue`.`metadata`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `issue` (`issue`.`id`,`issue`.`checksum`,`issue`.`user_id`,`issue`.`project_id`,`issue`.`title`,`issue`.`body`,`issue`.`state`,`issue`.`type`,`issue`.`resolution`,`issue`.`created_at`,`issue`.`updated_at`,`issue`.`closed_at`,`issue`.`closedby_user_id`,`issue`.`url`,`issue`.`customer_id`,`issue`.`ref_type`,`issue`.`ref_id`,`issue`.`metadata`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`user_id`=VALUES(`user_id`),`project_id`=VALUES(`project_id`),`title`=VALUES(`title`),`body`=VALUES(`body`),`state`=VALUES(`state`),`type`=VALUES(`type`),`resolution`=VALUES(`resolution`),`created_at`=VALUES(`created_at`),`updated_at`=VALUES(`updated_at`),`closed_at`=VALUES(`closed_at`),`closedby_user_id`=VALUES(`closedby_user_id`),`url`=VALUES(`url`),`customer_id`=VALUES(`customer_id`),`ref_type`=VALUES(`ref_type`),`ref_id`=VALUES(`ref_id`),`metadata`=VALUES(`metadata`)"
	}
	r, err := tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.UserID),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLString(t.Title),
		orm.ToSQLString(t.Body),
		orm.ToSQLString(enumIssueStateToString(&t.State)),
		orm.ToSQLString(enumIssueTypeToString(&t.Type)),
		orm.ToSQLString(enumIssueResolutionToString(&t.Resolution)),
		orm.ToSQLInt64(t.CreatedAt),
		orm.ToSQLInt64(t.UpdatedAt),
		orm.ToSQLInt64(t.ClosedAt),
		orm.ToSQLString(t.ClosedbyUserID),
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

// DBFindOne will find a Issue record in the database with the primary key
func (t *Issue) DBFindOne(ctx context.Context, db DB, value string) (bool, error) {
	q := "SELECT `issue`.`id`,`issue`.`checksum`,`issue`.`user_id`,`issue`.`project_id`,`issue`.`title`,`issue`.`body`,`issue`.`state`,`issue`.`type`,`issue`.`resolution`,`issue`.`created_at`,`issue`.`updated_at`,`issue`.`closed_at`,`issue`.`closedby_user_id`,`issue`.`url`,`issue`.`customer_id`,`issue`.`ref_type`,`issue`.`ref_id`,`issue`.`metadata` FROM `issue` WHERE `id` = ? LIMIT 1"
	row := db.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _UserID sql.NullString
	var _ProjectID sql.NullString
	var _Title sql.NullString
	var _Body sql.NullString
	var _State sql.NullString
	var _Type sql.NullString
	var _Resolution sql.NullString
	var _CreatedAt sql.NullInt64
	var _UpdatedAt sql.NullInt64
	var _ClosedAt sql.NullInt64
	var _ClosedbyUserID sql.NullString
	var _URL sql.NullString
	var _CustomerID sql.NullString
	var _RefType sql.NullString
	var _RefID sql.NullString
	var _Metadata sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_UserID,
		&_ProjectID,
		&_Title,
		&_Body,
		&_State,
		&_Type,
		&_Resolution,
		&_CreatedAt,
		&_UpdatedAt,
		&_ClosedAt,
		&_ClosedbyUserID,
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
	if _UserID.Valid {
		t.SetUserID(_UserID.String)
	}
	if _ProjectID.Valid {
		t.SetProjectID(_ProjectID.String)
	}
	if _Title.Valid {
		t.SetTitle(_Title.String)
	}
	if _Body.Valid {
		t.SetBody(_Body.String)
	}
	if _State.Valid {
		t.SetStateString(_State.String)
	}
	if _Type.Valid {
		t.SetTypeString(_Type.String)
	}
	if _Resolution.Valid {
		t.SetResolutionString(_Resolution.String)
	}
	if _CreatedAt.Valid {
		t.SetCreatedAt(_CreatedAt.Int64)
	}
	if _UpdatedAt.Valid {
		t.SetUpdatedAt(_UpdatedAt.Int64)
	}
	if _ClosedAt.Valid {
		t.SetClosedAt(_ClosedAt.Int64)
	}
	if _ClosedbyUserID.Valid {
		t.SetClosedbyUserID(_ClosedbyUserID.String)
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

// DBFindOneTx will find a Issue record in the database with the primary key using the provided transaction
func (t *Issue) DBFindOneTx(ctx context.Context, tx Tx, value string) (bool, error) {
	q := "SELECT `issue`.`id`,`issue`.`checksum`,`issue`.`user_id`,`issue`.`project_id`,`issue`.`title`,`issue`.`body`,`issue`.`state`,`issue`.`type`,`issue`.`resolution`,`issue`.`created_at`,`issue`.`updated_at`,`issue`.`closed_at`,`issue`.`closedby_user_id`,`issue`.`url`,`issue`.`customer_id`,`issue`.`ref_type`,`issue`.`ref_id`,`issue`.`metadata` FROM `issue` WHERE `id` = ? LIMIT 1"
	row := tx.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _UserID sql.NullString
	var _ProjectID sql.NullString
	var _Title sql.NullString
	var _Body sql.NullString
	var _State sql.NullString
	var _Type sql.NullString
	var _Resolution sql.NullString
	var _CreatedAt sql.NullInt64
	var _UpdatedAt sql.NullInt64
	var _ClosedAt sql.NullInt64
	var _ClosedbyUserID sql.NullString
	var _URL sql.NullString
	var _CustomerID sql.NullString
	var _RefType sql.NullString
	var _RefID sql.NullString
	var _Metadata sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_UserID,
		&_ProjectID,
		&_Title,
		&_Body,
		&_State,
		&_Type,
		&_Resolution,
		&_CreatedAt,
		&_UpdatedAt,
		&_ClosedAt,
		&_ClosedbyUserID,
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
	if _UserID.Valid {
		t.SetUserID(_UserID.String)
	}
	if _ProjectID.Valid {
		t.SetProjectID(_ProjectID.String)
	}
	if _Title.Valid {
		t.SetTitle(_Title.String)
	}
	if _Body.Valid {
		t.SetBody(_Body.String)
	}
	if _State.Valid {
		t.SetStateString(_State.String)
	}
	if _Type.Valid {
		t.SetTypeString(_Type.String)
	}
	if _Resolution.Valid {
		t.SetResolutionString(_Resolution.String)
	}
	if _CreatedAt.Valid {
		t.SetCreatedAt(_CreatedAt.Int64)
	}
	if _UpdatedAt.Valid {
		t.SetUpdatedAt(_UpdatedAt.Int64)
	}
	if _ClosedAt.Valid {
		t.SetClosedAt(_ClosedAt.Int64)
	}
	if _ClosedbyUserID.Valid {
		t.SetClosedbyUserID(_ClosedbyUserID.String)
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

// FindIssues will find a Issue record in the database with the provided parameters
func FindIssues(ctx context.Context, db DB, _params ...interface{}) ([]*Issue, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("user_id"),
		orm.Column("project_id"),
		orm.Column("title"),
		orm.Column("body"),
		orm.Column("state"),
		orm.Column("type"),
		orm.Column("resolution"),
		orm.Column("created_at"),
		orm.Column("updated_at"),
		orm.Column("closed_at"),
		orm.Column("closedby_user_id"),
		orm.Column("url"),
		orm.Column("customer_id"),
		orm.Column("ref_type"),
		orm.Column("ref_id"),
		orm.Column("metadata"),
		orm.Table(IssueTableName),
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
	results := make([]*Issue, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _UserID sql.NullString
		var _ProjectID sql.NullString
		var _Title sql.NullString
		var _Body sql.NullString
		var _State sql.NullString
		var _Type sql.NullString
		var _Resolution sql.NullString
		var _CreatedAt sql.NullInt64
		var _UpdatedAt sql.NullInt64
		var _ClosedAt sql.NullInt64
		var _ClosedbyUserID sql.NullString
		var _URL sql.NullString
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefID sql.NullString
		var _Metadata sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_UserID,
			&_ProjectID,
			&_Title,
			&_Body,
			&_State,
			&_Type,
			&_Resolution,
			&_CreatedAt,
			&_UpdatedAt,
			&_ClosedAt,
			&_ClosedbyUserID,
			&_URL,
			&_CustomerID,
			&_RefType,
			&_RefID,
			&_Metadata,
		)
		if err != nil {
			return nil, err
		}
		t := &Issue{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _Title.Valid {
			t.SetTitle(_Title.String)
		}
		if _Body.Valid {
			t.SetBody(_Body.String)
		}
		if _State.Valid {
			t.SetStateString(_State.String)
		}
		if _Type.Valid {
			t.SetTypeString(_Type.String)
		}
		if _Resolution.Valid {
			t.SetResolutionString(_Resolution.String)
		}
		if _CreatedAt.Valid {
			t.SetCreatedAt(_CreatedAt.Int64)
		}
		if _UpdatedAt.Valid {
			t.SetUpdatedAt(_UpdatedAt.Int64)
		}
		if _ClosedAt.Valid {
			t.SetClosedAt(_ClosedAt.Int64)
		}
		if _ClosedbyUserID.Valid {
			t.SetClosedbyUserID(_ClosedbyUserID.String)
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

// FindIssuesTx will find a Issue record in the database with the provided parameters using the provided transaction
func FindIssuesTx(ctx context.Context, tx Tx, _params ...interface{}) ([]*Issue, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("user_id"),
		orm.Column("project_id"),
		orm.Column("title"),
		orm.Column("body"),
		orm.Column("state"),
		orm.Column("type"),
		orm.Column("resolution"),
		orm.Column("created_at"),
		orm.Column("updated_at"),
		orm.Column("closed_at"),
		orm.Column("closedby_user_id"),
		orm.Column("url"),
		orm.Column("customer_id"),
		orm.Column("ref_type"),
		orm.Column("ref_id"),
		orm.Column("metadata"),
		orm.Table(IssueTableName),
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
	results := make([]*Issue, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _UserID sql.NullString
		var _ProjectID sql.NullString
		var _Title sql.NullString
		var _Body sql.NullString
		var _State sql.NullString
		var _Type sql.NullString
		var _Resolution sql.NullString
		var _CreatedAt sql.NullInt64
		var _UpdatedAt sql.NullInt64
		var _ClosedAt sql.NullInt64
		var _ClosedbyUserID sql.NullString
		var _URL sql.NullString
		var _CustomerID sql.NullString
		var _RefType sql.NullString
		var _RefID sql.NullString
		var _Metadata sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_UserID,
			&_ProjectID,
			&_Title,
			&_Body,
			&_State,
			&_Type,
			&_Resolution,
			&_CreatedAt,
			&_UpdatedAt,
			&_ClosedAt,
			&_ClosedbyUserID,
			&_URL,
			&_CustomerID,
			&_RefType,
			&_RefID,
			&_Metadata,
		)
		if err != nil {
			return nil, err
		}
		t := &Issue{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _Title.Valid {
			t.SetTitle(_Title.String)
		}
		if _Body.Valid {
			t.SetBody(_Body.String)
		}
		if _State.Valid {
			t.SetStateString(_State.String)
		}
		if _Type.Valid {
			t.SetTypeString(_Type.String)
		}
		if _Resolution.Valid {
			t.SetResolutionString(_Resolution.String)
		}
		if _CreatedAt.Valid {
			t.SetCreatedAt(_CreatedAt.Int64)
		}
		if _UpdatedAt.Valid {
			t.SetUpdatedAt(_UpdatedAt.Int64)
		}
		if _ClosedAt.Valid {
			t.SetClosedAt(_ClosedAt.Int64)
		}
		if _ClosedbyUserID.Valid {
			t.SetClosedbyUserID(_ClosedbyUserID.String)
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

// DBFind will find a Issue record in the database with the provided parameters
func (t *Issue) DBFind(ctx context.Context, db DB, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("user_id"),
		orm.Column("project_id"),
		orm.Column("title"),
		orm.Column("body"),
		orm.Column("state"),
		orm.Column("type"),
		orm.Column("resolution"),
		orm.Column("created_at"),
		orm.Column("updated_at"),
		orm.Column("closed_at"),
		orm.Column("closedby_user_id"),
		orm.Column("url"),
		orm.Column("customer_id"),
		orm.Column("ref_type"),
		orm.Column("ref_id"),
		orm.Column("metadata"),
		orm.Table(IssueTableName),
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
	var _UserID sql.NullString
	var _ProjectID sql.NullString
	var _Title sql.NullString
	var _Body sql.NullString
	var _State sql.NullString
	var _Type sql.NullString
	var _Resolution sql.NullString
	var _CreatedAt sql.NullInt64
	var _UpdatedAt sql.NullInt64
	var _ClosedAt sql.NullInt64
	var _ClosedbyUserID sql.NullString
	var _URL sql.NullString
	var _CustomerID sql.NullString
	var _RefType sql.NullString
	var _RefID sql.NullString
	var _Metadata sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_UserID,
		&_ProjectID,
		&_Title,
		&_Body,
		&_State,
		&_Type,
		&_Resolution,
		&_CreatedAt,
		&_UpdatedAt,
		&_ClosedAt,
		&_ClosedbyUserID,
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
	if _UserID.Valid {
		t.SetUserID(_UserID.String)
	}
	if _ProjectID.Valid {
		t.SetProjectID(_ProjectID.String)
	}
	if _Title.Valid {
		t.SetTitle(_Title.String)
	}
	if _Body.Valid {
		t.SetBody(_Body.String)
	}
	if _State.Valid {
		t.SetStateString(_State.String)
	}
	if _Type.Valid {
		t.SetTypeString(_Type.String)
	}
	if _Resolution.Valid {
		t.SetResolutionString(_Resolution.String)
	}
	if _CreatedAt.Valid {
		t.SetCreatedAt(_CreatedAt.Int64)
	}
	if _UpdatedAt.Valid {
		t.SetUpdatedAt(_UpdatedAt.Int64)
	}
	if _ClosedAt.Valid {
		t.SetClosedAt(_ClosedAt.Int64)
	}
	if _ClosedbyUserID.Valid {
		t.SetClosedbyUserID(_ClosedbyUserID.String)
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

// DBFindTx will find a Issue record in the database with the provided parameters using the provided transaction
func (t *Issue) DBFindTx(ctx context.Context, tx Tx, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("user_id"),
		orm.Column("project_id"),
		orm.Column("title"),
		orm.Column("body"),
		orm.Column("state"),
		orm.Column("type"),
		orm.Column("resolution"),
		orm.Column("created_at"),
		orm.Column("updated_at"),
		orm.Column("closed_at"),
		orm.Column("closedby_user_id"),
		orm.Column("url"),
		orm.Column("customer_id"),
		orm.Column("ref_type"),
		orm.Column("ref_id"),
		orm.Column("metadata"),
		orm.Table(IssueTableName),
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
	var _UserID sql.NullString
	var _ProjectID sql.NullString
	var _Title sql.NullString
	var _Body sql.NullString
	var _State sql.NullString
	var _Type sql.NullString
	var _Resolution sql.NullString
	var _CreatedAt sql.NullInt64
	var _UpdatedAt sql.NullInt64
	var _ClosedAt sql.NullInt64
	var _ClosedbyUserID sql.NullString
	var _URL sql.NullString
	var _CustomerID sql.NullString
	var _RefType sql.NullString
	var _RefID sql.NullString
	var _Metadata sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_UserID,
		&_ProjectID,
		&_Title,
		&_Body,
		&_State,
		&_Type,
		&_Resolution,
		&_CreatedAt,
		&_UpdatedAt,
		&_ClosedAt,
		&_ClosedbyUserID,
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
	if _UserID.Valid {
		t.SetUserID(_UserID.String)
	}
	if _ProjectID.Valid {
		t.SetProjectID(_ProjectID.String)
	}
	if _Title.Valid {
		t.SetTitle(_Title.String)
	}
	if _Body.Valid {
		t.SetBody(_Body.String)
	}
	if _State.Valid {
		t.SetStateString(_State.String)
	}
	if _Type.Valid {
		t.SetTypeString(_Type.String)
	}
	if _Resolution.Valid {
		t.SetResolutionString(_Resolution.String)
	}
	if _CreatedAt.Valid {
		t.SetCreatedAt(_CreatedAt.Int64)
	}
	if _UpdatedAt.Valid {
		t.SetUpdatedAt(_UpdatedAt.Int64)
	}
	if _ClosedAt.Valid {
		t.SetClosedAt(_ClosedAt.Int64)
	}
	if _ClosedbyUserID.Valid {
		t.SetClosedbyUserID(_ClosedbyUserID.String)
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

// CountIssues will find the count of Issue records in the database
func CountIssues(ctx context.Context, db DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(IssueTableName),
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

// CountIssuesTx will find the count of Issue records in the database using the provided transaction
func CountIssuesTx(ctx context.Context, tx Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(IssueTableName),
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

// DBCount will find the count of Issue records in the database
func (t *Issue) DBCount(ctx context.Context, db DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(IssueTableName),
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

// DBCountTx will find the count of Issue records in the database using the provided transaction
func (t *Issue) DBCountTx(ctx context.Context, tx Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(IssueTableName),
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

// DBExists will return true if the Issue record exists in the database
func (t *Issue) DBExists(ctx context.Context, db DB) (bool, error) {
	q := "SELECT `id` FROM `issue` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := db.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// DBExistsTx will return true if the Issue record exists in the database using the provided transaction
func (t *Issue) DBExistsTx(ctx context.Context, tx Tx) (bool, error) {
	q := "SELECT `id` FROM `issue` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := tx.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// PrimaryKeyColumn returns the column name for the primary key
func (t *Issue) PrimaryKeyColumn() string {
	return IssueColumnID
}

// PrimaryKeyColumnType returns the primary key column Go type as a string
func (t *Issue) PrimaryKeyColumnType() string {
	return "string"
}

// PrimaryKey returns the primary key column value
func (t *Issue) PrimaryKey() interface{} {
	return t.ID
}
