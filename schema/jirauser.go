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

var _ Model = (*JiraUser)(nil)
var _ CSVWriter = (*JiraUser)(nil)
var _ JSONWriter = (*JiraUser)(nil)
var _ Checksum = (*JiraUser)(nil)

// JiraUserTableName is the name of the table in SQL
const JiraUserTableName = "jira_user"

var JiraUserColumns = []string{
	"id",
	"checksum",
	"user_id",
	"username",
	"name",
	"email",
	"avatar_url",
	"url",
	"customer_id",
	"ref_id",
}

// JiraUser table
type JiraUser struct {
	AvatarURL  string  `json:"avatar_url"`
	Checksum   *string `json:"checksum,omitempty"`
	CustomerID string  `json:"customer_id"`
	Email      string  `json:"email"`
	ID         string  `json:"id"`
	Name       string  `json:"name"`
	RefID      string  `json:"ref_id"`
	URL        string  `json:"url"`
	UserID     string  `json:"user_id"`
	Username   string  `json:"username"`
}

// TableName returns the SQL table name for JiraUser and satifies the Model interface
func (t *JiraUser) TableName() string {
	return JiraUserTableName
}

// ToCSV will serialize the JiraUser instance to a CSV compatible array of strings
func (t *JiraUser) ToCSV() []string {
	return []string{
		t.ID,
		t.CalculateChecksum(),
		t.UserID,
		t.Username,
		t.Name,
		t.Email,
		t.AvatarURL,
		t.URL,
		t.CustomerID,
		t.RefID,
	}
}

// WriteCSV will serialize the JiraUser instance to the writer as CSV and satisfies the CSVWriter interface
func (t *JiraUser) WriteCSV(w *csv.Writer) error {
	return w.Write(t.ToCSV())
}

// WriteJSON will serialize the JiraUser instance to the writer as JSON and satisfies the JSONWriter interface
func (t *JiraUser) WriteJSON(w io.Writer, indent ...bool) error {
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

// NewJiraUserReader creates a JSON reader which can read in JiraUser objects serialized as JSON either as an array, single object or json new lines
// and writes each JiraUser to the channel provided
func NewJiraUserReader(r io.Reader, ch chan<- JiraUser) error {
	return orm.Deserialize(r, func(buf json.RawMessage) error {
		dec := json.NewDecoder(bytes.NewBuffer(buf))
		e := JiraUser{}
		if err := dec.Decode(&e); err != nil {
			return err
		}
		ch <- e
		return nil
	})
}

// NewCSVJiraUserReaderDir will read the reader as CSV and emit each record to the channel provided
func NewCSVJiraUserReader(r io.Reader, ch chan<- JiraUser) error {
	cr := csv.NewReader(r)
	for {
		record, err := cr.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		ch <- JiraUser{
			ID:         record[0],
			Checksum:   fromStringPointer(record[1]),
			UserID:     record[2],
			Username:   record[3],
			Name:       record[4],
			Email:      record[5],
			AvatarURL:  record[6],
			URL:        record[7],
			CustomerID: record[8],
			RefID:      record[9],
		}
	}
	return nil
}

// NewCSVJiraUserReaderFile will read the file as a CSV and emit each record to the channel provided
func NewCSVJiraUserReaderFile(fp string, ch chan<- JiraUser) error {
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
	return NewCSVJiraUserReader(fc, ch)
}

// NewCSVJiraUserReaderDir will read the jira_user.csv.gz file as a CSV and emit each record to the channel provided
func NewCSVJiraUserReaderDir(dir string, ch chan<- JiraUser) error {
	return NewCSVJiraUserReaderFile(filepath.Join(dir, "jira_user.csv.gz"), ch)
}

// JiraUserCSVDeduper is a function callback which takes the existing value (a) and the new value (b)
// and the return value should be the one to use (or a new one, if applicable). return nil
// to skip processing of this record
type JiraUserCSVDeduper func(a JiraUser, b JiraUser) *JiraUser

// JiraUserCSVDedupeDisabled is set on whether the CSV writer should de-dupe values by key
var JiraUserCSVDedupeDisabled bool

// NewJiraUserCSVWriterSize creates a batch writer that will write each JiraUser into a CSV file
// this method will automatically de-duplicate entries using the primary key. if the checksum
// for a newer item with the same primary key doesn't match a previously sent item, the newer
// one will be used
func NewJiraUserCSVWriterSize(w io.Writer, size int, dedupers ...JiraUserCSVDeduper) (chan JiraUser, chan bool, error) {
	cw := csv.NewWriter(w)
	ch := make(chan JiraUser, size)
	done := make(chan bool)
	go func() {
		defer func() { done <- true }()
		dodedupe := !JiraUserCSVDedupeDisabled
		var kv map[string]*JiraUser
		var deduper JiraUserCSVDeduper
		if dedupers != nil && len(dedupers) > 0 {
			deduper = dedupers[0]
			dodedupe = true
		}
		if dodedupe {
			kv = make(map[string]*JiraUser)
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

// JiraUserCSVDefaultSize is the default channel buffer size if not provided
var JiraUserCSVDefaultSize = 100

// NewJiraUserCSVWriter creates a batch writer that will write each JiraUser into a CSV file
func NewJiraUserCSVWriter(w io.Writer, dedupers ...JiraUserCSVDeduper) (chan JiraUser, chan bool, error) {
	return NewJiraUserCSVWriterSize(w, JiraUserCSVDefaultSize, dedupers...)
}

// NewJiraUserCSVWriterDir creates a batch writer that will write each JiraUser into a CSV file named jira_user.csv.gz in dir
func NewJiraUserCSVWriterDir(dir string, dedupers ...JiraUserCSVDeduper) (chan JiraUser, chan bool, error) {
	return NewJiraUserCSVWriterFile(filepath.Join(dir, "jira_user.csv.gz"), dedupers...)
}

// NewJiraUserCSVWriterFile creates a batch writer that will write each JiraUser into a CSV file
func NewJiraUserCSVWriterFile(fn string, dedupers ...JiraUserCSVDeduper) (chan JiraUser, chan bool, error) {
	f, err := os.Create(fn)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening CSV file %s. %v", fn, err)
	}
	var fc io.WriteCloser = f
	if filepath.Ext(fn) == ".gz" {
		w, _ := gzip.NewWriterLevel(f, gzip.BestCompression)
		fc = w
	}
	ch, done, err := NewJiraUserCSVWriter(fc, dedupers...)
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

type JiraUserDBAction func(ctx context.Context, db DB, record JiraUser) error

// NewJiraUserDBWriterSize creates a DB writer that will write each issue into the DB
func NewJiraUserDBWriterSize(ctx context.Context, db DB, errors chan<- error, size int, actions ...JiraUserDBAction) (chan JiraUser, chan bool, error) {
	ch := make(chan JiraUser, size)
	done := make(chan bool)
	var action JiraUserDBAction
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

// NewJiraUserDBWriter creates a DB writer that will write each issue into the DB
func NewJiraUserDBWriter(ctx context.Context, db DB, errors chan<- error, actions ...JiraUserDBAction) (chan JiraUser, chan bool, error) {
	return NewJiraUserDBWriterSize(ctx, db, errors, 100, actions...)
}

// JiraUserColumnID is the ID SQL column name for the JiraUser table
const JiraUserColumnID = "id"

// JiraUserEscapedColumnID is the escaped ID SQL column name for the JiraUser table
const JiraUserEscapedColumnID = "`id`"

// JiraUserColumnChecksum is the Checksum SQL column name for the JiraUser table
const JiraUserColumnChecksum = "checksum"

// JiraUserEscapedColumnChecksum is the escaped Checksum SQL column name for the JiraUser table
const JiraUserEscapedColumnChecksum = "`checksum`"

// JiraUserColumnUserID is the UserID SQL column name for the JiraUser table
const JiraUserColumnUserID = "user_id"

// JiraUserEscapedColumnUserID is the escaped UserID SQL column name for the JiraUser table
const JiraUserEscapedColumnUserID = "`user_id`"

// JiraUserColumnUsername is the Username SQL column name for the JiraUser table
const JiraUserColumnUsername = "username"

// JiraUserEscapedColumnUsername is the escaped Username SQL column name for the JiraUser table
const JiraUserEscapedColumnUsername = "`username`"

// JiraUserColumnName is the Name SQL column name for the JiraUser table
const JiraUserColumnName = "name"

// JiraUserEscapedColumnName is the escaped Name SQL column name for the JiraUser table
const JiraUserEscapedColumnName = "`name`"

// JiraUserColumnEmail is the Email SQL column name for the JiraUser table
const JiraUserColumnEmail = "email"

// JiraUserEscapedColumnEmail is the escaped Email SQL column name for the JiraUser table
const JiraUserEscapedColumnEmail = "`email`"

// JiraUserColumnAvatarURL is the AvatarURL SQL column name for the JiraUser table
const JiraUserColumnAvatarURL = "avatar_url"

// JiraUserEscapedColumnAvatarURL is the escaped AvatarURL SQL column name for the JiraUser table
const JiraUserEscapedColumnAvatarURL = "`avatar_url`"

// JiraUserColumnURL is the URL SQL column name for the JiraUser table
const JiraUserColumnURL = "url"

// JiraUserEscapedColumnURL is the escaped URL SQL column name for the JiraUser table
const JiraUserEscapedColumnURL = "`url`"

// JiraUserColumnCustomerID is the CustomerID SQL column name for the JiraUser table
const JiraUserColumnCustomerID = "customer_id"

// JiraUserEscapedColumnCustomerID is the escaped CustomerID SQL column name for the JiraUser table
const JiraUserEscapedColumnCustomerID = "`customer_id`"

// JiraUserColumnRefID is the RefID SQL column name for the JiraUser table
const JiraUserColumnRefID = "ref_id"

// JiraUserEscapedColumnRefID is the escaped RefID SQL column name for the JiraUser table
const JiraUserEscapedColumnRefID = "`ref_id`"

// GetID will return the JiraUser ID value
func (t *JiraUser) GetID() string {
	return t.ID
}

// SetID will set the JiraUser ID value
func (t *JiraUser) SetID(v string) {
	t.ID = v
}

// FindJiraUserByID will find a JiraUser by ID
func FindJiraUserByID(ctx context.Context, db DB, value string) (*JiraUser, error) {
	q := "SELECT * FROM `jira_user` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _UserID sql.NullString
	var _Username sql.NullString
	var _Name sql.NullString
	var _Email sql.NullString
	var _AvatarURL sql.NullString
	var _URL sql.NullString
	var _CustomerID sql.NullString
	var _RefID sql.NullString
	err := db.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_UserID,
		&_Username,
		&_Name,
		&_Email,
		&_AvatarURL,
		&_URL,
		&_CustomerID,
		&_RefID,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &JiraUser{}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _Checksum.Valid {
		t.SetChecksum(_Checksum.String)
	}
	if _UserID.Valid {
		t.SetUserID(_UserID.String)
	}
	if _Username.Valid {
		t.SetUsername(_Username.String)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _Email.Valid {
		t.SetEmail(_Email.String)
	}
	if _AvatarURL.Valid {
		t.SetAvatarURL(_AvatarURL.String)
	}
	if _URL.Valid {
		t.SetURL(_URL.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _RefID.Valid {
		t.SetRefID(_RefID.String)
	}
	return t, nil
}

// FindJiraUserByIDTx will find a JiraUser by ID using the provided transaction
func FindJiraUserByIDTx(ctx context.Context, tx Tx, value string) (*JiraUser, error) {
	q := "SELECT * FROM `jira_user` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _UserID sql.NullString
	var _Username sql.NullString
	var _Name sql.NullString
	var _Email sql.NullString
	var _AvatarURL sql.NullString
	var _URL sql.NullString
	var _CustomerID sql.NullString
	var _RefID sql.NullString
	err := tx.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_UserID,
		&_Username,
		&_Name,
		&_Email,
		&_AvatarURL,
		&_URL,
		&_CustomerID,
		&_RefID,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &JiraUser{}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _Checksum.Valid {
		t.SetChecksum(_Checksum.String)
	}
	if _UserID.Valid {
		t.SetUserID(_UserID.String)
	}
	if _Username.Valid {
		t.SetUsername(_Username.String)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _Email.Valid {
		t.SetEmail(_Email.String)
	}
	if _AvatarURL.Valid {
		t.SetAvatarURL(_AvatarURL.String)
	}
	if _URL.Valid {
		t.SetURL(_URL.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _RefID.Valid {
		t.SetRefID(_RefID.String)
	}
	return t, nil
}

// GetChecksum will return the JiraUser Checksum value
func (t *JiraUser) GetChecksum() string {
	if t.Checksum == nil {
		return ""
	}
	return *t.Checksum
}

// SetChecksum will set the JiraUser Checksum value
func (t *JiraUser) SetChecksum(v string) {
	t.Checksum = &v
}

// GetUserID will return the JiraUser UserID value
func (t *JiraUser) GetUserID() string {
	return t.UserID
}

// SetUserID will set the JiraUser UserID value
func (t *JiraUser) SetUserID(v string) {
	t.UserID = v
}

// FindJiraUsersByUserID will find all JiraUsers by the UserID value
func FindJiraUsersByUserID(ctx context.Context, db DB, value string) ([]*JiraUser, error) {
	q := "SELECT * FROM `jira_user` WHERE `user_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraUser, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _UserID sql.NullString
		var _Username sql.NullString
		var _Name sql.NullString
		var _Email sql.NullString
		var _AvatarURL sql.NullString
		var _URL sql.NullString
		var _CustomerID sql.NullString
		var _RefID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_UserID,
			&_Username,
			&_Name,
			&_Email,
			&_AvatarURL,
			&_URL,
			&_CustomerID,
			&_RefID,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraUser{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _Username.Valid {
			t.SetUsername(_Username.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _Email.Valid {
			t.SetEmail(_Email.String)
		}
		if _AvatarURL.Valid {
			t.SetAvatarURL(_AvatarURL.String)
		}
		if _URL.Valid {
			t.SetURL(_URL.String)
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

// FindJiraUsersByUserIDTx will find all JiraUsers by the UserID value using the provided transaction
func FindJiraUsersByUserIDTx(ctx context.Context, tx Tx, value string) ([]*JiraUser, error) {
	q := "SELECT * FROM `jira_user` WHERE `user_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraUser, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _UserID sql.NullString
		var _Username sql.NullString
		var _Name sql.NullString
		var _Email sql.NullString
		var _AvatarURL sql.NullString
		var _URL sql.NullString
		var _CustomerID sql.NullString
		var _RefID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_UserID,
			&_Username,
			&_Name,
			&_Email,
			&_AvatarURL,
			&_URL,
			&_CustomerID,
			&_RefID,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraUser{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _Username.Valid {
			t.SetUsername(_Username.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _Email.Valid {
			t.SetEmail(_Email.String)
		}
		if _AvatarURL.Valid {
			t.SetAvatarURL(_AvatarURL.String)
		}
		if _URL.Valid {
			t.SetURL(_URL.String)
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

// GetUsername will return the JiraUser Username value
func (t *JiraUser) GetUsername() string {
	return t.Username
}

// SetUsername will set the JiraUser Username value
func (t *JiraUser) SetUsername(v string) {
	t.Username = v
}

// GetName will return the JiraUser Name value
func (t *JiraUser) GetName() string {
	return t.Name
}

// SetName will set the JiraUser Name value
func (t *JiraUser) SetName(v string) {
	t.Name = v
}

// GetEmail will return the JiraUser Email value
func (t *JiraUser) GetEmail() string {
	return t.Email
}

// SetEmail will set the JiraUser Email value
func (t *JiraUser) SetEmail(v string) {
	t.Email = v
}

// GetAvatarURL will return the JiraUser AvatarURL value
func (t *JiraUser) GetAvatarURL() string {
	return t.AvatarURL
}

// SetAvatarURL will set the JiraUser AvatarURL value
func (t *JiraUser) SetAvatarURL(v string) {
	t.AvatarURL = v
}

// GetURL will return the JiraUser URL value
func (t *JiraUser) GetURL() string {
	return t.URL
}

// SetURL will set the JiraUser URL value
func (t *JiraUser) SetURL(v string) {
	t.URL = v
}

// GetCustomerID will return the JiraUser CustomerID value
func (t *JiraUser) GetCustomerID() string {
	return t.CustomerID
}

// SetCustomerID will set the JiraUser CustomerID value
func (t *JiraUser) SetCustomerID(v string) {
	t.CustomerID = v
}

// FindJiraUsersByCustomerID will find all JiraUsers by the CustomerID value
func FindJiraUsersByCustomerID(ctx context.Context, db DB, value string) ([]*JiraUser, error) {
	q := "SELECT * FROM `jira_user` WHERE `customer_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraUser, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _UserID sql.NullString
		var _Username sql.NullString
		var _Name sql.NullString
		var _Email sql.NullString
		var _AvatarURL sql.NullString
		var _URL sql.NullString
		var _CustomerID sql.NullString
		var _RefID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_UserID,
			&_Username,
			&_Name,
			&_Email,
			&_AvatarURL,
			&_URL,
			&_CustomerID,
			&_RefID,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraUser{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _Username.Valid {
			t.SetUsername(_Username.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _Email.Valid {
			t.SetEmail(_Email.String)
		}
		if _AvatarURL.Valid {
			t.SetAvatarURL(_AvatarURL.String)
		}
		if _URL.Valid {
			t.SetURL(_URL.String)
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

// FindJiraUsersByCustomerIDTx will find all JiraUsers by the CustomerID value using the provided transaction
func FindJiraUsersByCustomerIDTx(ctx context.Context, tx Tx, value string) ([]*JiraUser, error) {
	q := "SELECT * FROM `jira_user` WHERE `customer_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraUser, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _UserID sql.NullString
		var _Username sql.NullString
		var _Name sql.NullString
		var _Email sql.NullString
		var _AvatarURL sql.NullString
		var _URL sql.NullString
		var _CustomerID sql.NullString
		var _RefID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_UserID,
			&_Username,
			&_Name,
			&_Email,
			&_AvatarURL,
			&_URL,
			&_CustomerID,
			&_RefID,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraUser{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _Username.Valid {
			t.SetUsername(_Username.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _Email.Valid {
			t.SetEmail(_Email.String)
		}
		if _AvatarURL.Valid {
			t.SetAvatarURL(_AvatarURL.String)
		}
		if _URL.Valid {
			t.SetURL(_URL.String)
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

// GetRefID will return the JiraUser RefID value
func (t *JiraUser) GetRefID() string {
	return t.RefID
}

// SetRefID will set the JiraUser RefID value
func (t *JiraUser) SetRefID(v string) {
	t.RefID = v
}

// FindJiraUsersByRefID will find all JiraUsers by the RefID value
func FindJiraUsersByRefID(ctx context.Context, db DB, value string) ([]*JiraUser, error) {
	q := "SELECT * FROM `jira_user` WHERE `ref_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraUser, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _UserID sql.NullString
		var _Username sql.NullString
		var _Name sql.NullString
		var _Email sql.NullString
		var _AvatarURL sql.NullString
		var _URL sql.NullString
		var _CustomerID sql.NullString
		var _RefID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_UserID,
			&_Username,
			&_Name,
			&_Email,
			&_AvatarURL,
			&_URL,
			&_CustomerID,
			&_RefID,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraUser{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _Username.Valid {
			t.SetUsername(_Username.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _Email.Valid {
			t.SetEmail(_Email.String)
		}
		if _AvatarURL.Valid {
			t.SetAvatarURL(_AvatarURL.String)
		}
		if _URL.Valid {
			t.SetURL(_URL.String)
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

// FindJiraUsersByRefIDTx will find all JiraUsers by the RefID value using the provided transaction
func FindJiraUsersByRefIDTx(ctx context.Context, tx Tx, value string) ([]*JiraUser, error) {
	q := "SELECT * FROM `jira_user` WHERE `ref_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraUser, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _UserID sql.NullString
		var _Username sql.NullString
		var _Name sql.NullString
		var _Email sql.NullString
		var _AvatarURL sql.NullString
		var _URL sql.NullString
		var _CustomerID sql.NullString
		var _RefID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_UserID,
			&_Username,
			&_Name,
			&_Email,
			&_AvatarURL,
			&_URL,
			&_CustomerID,
			&_RefID,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraUser{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _Username.Valid {
			t.SetUsername(_Username.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _Email.Valid {
			t.SetEmail(_Email.String)
		}
		if _AvatarURL.Valid {
			t.SetAvatarURL(_AvatarURL.String)
		}
		if _URL.Valid {
			t.SetURL(_URL.String)
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

func (t *JiraUser) toTimestamp(value time.Time) *timestamp.Timestamp {
	ts, _ := ptypes.TimestampProto(value)
	return ts
}

// DBCreateJiraUserTable will create the JiraUser table
func DBCreateJiraUserTable(ctx context.Context, db DB) error {
	q := "CREATE TABLE `jira_user` (`id` VARCHAR(64) NOT NULL PRIMARY KEY,`checksum` CHAR(64),`user_id`VARCHAR(255) NOT NULL,`username` TEXT NOT NULL,`name`TEXT NOT NULL,`email` VARCHAR(255) NOT NULL,`avatar_url`VARCHAR(255) NOT NULL,`url` VARCHAR(255) NOT NULL,`customer_id` VARCHAR(64) NOT NULL,`ref_id` VARCHAR(64) NOT NULL,INDEX jira_user_user_id_index (`user_id`),INDEX jira_user_customer_id_index (`customer_id`),INDEX jira_user_ref_id_index (`ref_id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBCreateJiraUserTableTx will create the JiraUser table using the provided transction
func DBCreateJiraUserTableTx(ctx context.Context, tx Tx) error {
	q := "CREATE TABLE `jira_user` (`id` VARCHAR(64) NOT NULL PRIMARY KEY,`checksum` CHAR(64),`user_id`VARCHAR(255) NOT NULL,`username` TEXT NOT NULL,`name`TEXT NOT NULL,`email` VARCHAR(255) NOT NULL,`avatar_url`VARCHAR(255) NOT NULL,`url` VARCHAR(255) NOT NULL,`customer_id` VARCHAR(64) NOT NULL,`ref_id` VARCHAR(64) NOT NULL,INDEX jira_user_user_id_index (`user_id`),INDEX jira_user_customer_id_index (`customer_id`),INDEX jira_user_ref_id_index (`ref_id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// DBDropJiraUserTable will drop the JiraUser table
func DBDropJiraUserTable(ctx context.Context, db DB) error {
	q := "DROP TABLE IF EXISTS `jira_user`"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBDropJiraUserTableTx will drop the JiraUser table using the provided transaction
func DBDropJiraUserTableTx(ctx context.Context, tx Tx) error {
	q := "DROP TABLE IF EXISTS `jira_user`"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// CalculateChecksum will calculate a checksum of the SHA1 of all field values
func (t *JiraUser) CalculateChecksum() string {
	return orm.HashStrings(
		orm.ToString(t.ID),
		orm.ToString(t.UserID),
		orm.ToString(t.Username),
		orm.ToString(t.Name),
		orm.ToString(t.Email),
		orm.ToString(t.AvatarURL),
		orm.ToString(t.URL),
		orm.ToString(t.CustomerID),
		orm.ToString(t.RefID),
	)
}

// DBCreate will create a new JiraUser record in the database
func (t *JiraUser) DBCreate(ctx context.Context, db DB) (sql.Result, error) {
	q := "INSERT INTO `jira_user` (`jira_user`.`id`,`jira_user`.`checksum`,`jira_user`.`user_id`,`jira_user`.`username`,`jira_user`.`name`,`jira_user`.`email`,`jira_user`.`avatar_url`,`jira_user`.`url`,`jira_user`.`customer_id`,`jira_user`.`ref_id`) VALUES (?,?,?,?,?,?,?,?,?,?)"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.UserID),
		orm.ToSQLString(t.Username),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.Email),
		orm.ToSQLString(t.AvatarURL),
		orm.ToSQLString(t.URL),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefID),
	)
}

// DBCreateTx will create a new JiraUser record in the database using the provided transaction
func (t *JiraUser) DBCreateTx(ctx context.Context, tx Tx) (sql.Result, error) {
	q := "INSERT INTO `jira_user` (`jira_user`.`id`,`jira_user`.`checksum`,`jira_user`.`user_id`,`jira_user`.`username`,`jira_user`.`name`,`jira_user`.`email`,`jira_user`.`avatar_url`,`jira_user`.`url`,`jira_user`.`customer_id`,`jira_user`.`ref_id`) VALUES (?,?,?,?,?,?,?,?,?,?)"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.UserID),
		orm.ToSQLString(t.Username),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.Email),
		orm.ToSQLString(t.AvatarURL),
		orm.ToSQLString(t.URL),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefID),
	)
}

// DBCreateIgnoreDuplicate will upsert the JiraUser record in the database
func (t *JiraUser) DBCreateIgnoreDuplicate(ctx context.Context, db DB) (sql.Result, error) {
	q := "INSERT INTO `jira_user` (`jira_user`.`id`,`jira_user`.`checksum`,`jira_user`.`user_id`,`jira_user`.`username`,`jira_user`.`name`,`jira_user`.`email`,`jira_user`.`avatar_url`,`jira_user`.`url`,`jira_user`.`customer_id`,`jira_user`.`ref_id`) VALUES (?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.UserID),
		orm.ToSQLString(t.Username),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.Email),
		orm.ToSQLString(t.AvatarURL),
		orm.ToSQLString(t.URL),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefID),
	)
}

// DBCreateIgnoreDuplicateTx will upsert the JiraUser record in the database using the provided transaction
func (t *JiraUser) DBCreateIgnoreDuplicateTx(ctx context.Context, tx Tx) (sql.Result, error) {
	q := "INSERT INTO `jira_user` (`jira_user`.`id`,`jira_user`.`checksum`,`jira_user`.`user_id`,`jira_user`.`username`,`jira_user`.`name`,`jira_user`.`email`,`jira_user`.`avatar_url`,`jira_user`.`url`,`jira_user`.`customer_id`,`jira_user`.`ref_id`) VALUES (?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.UserID),
		orm.ToSQLString(t.Username),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.Email),
		orm.ToSQLString(t.AvatarURL),
		orm.ToSQLString(t.URL),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefID),
	)
}

// DeleteAllJiraUsers deletes all JiraUser records in the database with optional filters
func DeleteAllJiraUsers(ctx context.Context, db DB, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(JiraUserTableName),
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

// DeleteAllJiraUsersTx deletes all JiraUser records in the database with optional filters using the provided transaction
func DeleteAllJiraUsersTx(ctx context.Context, tx Tx, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(JiraUserTableName),
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

// DBDelete will delete this JiraUser record in the database
func (t *JiraUser) DBDelete(ctx context.Context, db DB) (bool, error) {
	q := "DELETE FROM `jira_user` WHERE `id` = ?"
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

// DBDeleteTx will delete this JiraUser record in the database using the provided transaction
func (t *JiraUser) DBDeleteTx(ctx context.Context, tx Tx) (bool, error) {
	q := "DELETE FROM `jira_user` WHERE `id` = ?"
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

// DBUpdate will update the JiraUser record in the database
func (t *JiraUser) DBUpdate(ctx context.Context, db DB) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `jira_user` SET `checksum`=?,`user_id`=?,`username`=?,`name`=?,`email`=?,`avatar_url`=?,`url`=?,`customer_id`=?,`ref_id`=? WHERE `id`=?"
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.UserID),
		orm.ToSQLString(t.Username),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.Email),
		orm.ToSQLString(t.AvatarURL),
		orm.ToSQLString(t.URL),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.ID),
	)
}

// DBUpdateTx will update the JiraUser record in the database using the provided transaction
func (t *JiraUser) DBUpdateTx(ctx context.Context, tx Tx) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `jira_user` SET `checksum`=?,`user_id`=?,`username`=?,`name`=?,`email`=?,`avatar_url`=?,`url`=?,`customer_id`=?,`ref_id`=? WHERE `id`=?"
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.UserID),
		orm.ToSQLString(t.Username),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.Email),
		orm.ToSQLString(t.AvatarURL),
		orm.ToSQLString(t.URL),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefID),
		orm.ToSQLString(t.ID),
	)
}

// DBUpsert will upsert the JiraUser record in the database
func (t *JiraUser) DBUpsert(ctx context.Context, db DB, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `jira_user` (`jira_user`.`id`,`jira_user`.`checksum`,`jira_user`.`user_id`,`jira_user`.`username`,`jira_user`.`name`,`jira_user`.`email`,`jira_user`.`avatar_url`,`jira_user`.`url`,`jira_user`.`customer_id`,`jira_user`.`ref_id`) VALUES (?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `jira_user` (`jira_user`.`id`,`jira_user`.`checksum`,`jira_user`.`user_id`,`jira_user`.`username`,`jira_user`.`name`,`jira_user`.`email`,`jira_user`.`avatar_url`,`jira_user`.`url`,`jira_user`.`customer_id`,`jira_user`.`ref_id`) VALUES (?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`user_id`=VALUES(`user_id`),`username`=VALUES(`username`),`name`=VALUES(`name`),`email`=VALUES(`email`),`avatar_url`=VALUES(`avatar_url`),`url`=VALUES(`url`),`customer_id`=VALUES(`customer_id`),`ref_id`=VALUES(`ref_id`)"
	}
	r, err := db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.UserID),
		orm.ToSQLString(t.Username),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.Email),
		orm.ToSQLString(t.AvatarURL),
		orm.ToSQLString(t.URL),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefID),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBUpsertTx will upsert the JiraUser record in the database using the provided transaction
func (t *JiraUser) DBUpsertTx(ctx context.Context, tx Tx, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `jira_user` (`jira_user`.`id`,`jira_user`.`checksum`,`jira_user`.`user_id`,`jira_user`.`username`,`jira_user`.`name`,`jira_user`.`email`,`jira_user`.`avatar_url`,`jira_user`.`url`,`jira_user`.`customer_id`,`jira_user`.`ref_id`) VALUES (?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `jira_user` (`jira_user`.`id`,`jira_user`.`checksum`,`jira_user`.`user_id`,`jira_user`.`username`,`jira_user`.`name`,`jira_user`.`email`,`jira_user`.`avatar_url`,`jira_user`.`url`,`jira_user`.`customer_id`,`jira_user`.`ref_id`) VALUES (?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`user_id`=VALUES(`user_id`),`username`=VALUES(`username`),`name`=VALUES(`name`),`email`=VALUES(`email`),`avatar_url`=VALUES(`avatar_url`),`url`=VALUES(`url`),`customer_id`=VALUES(`customer_id`),`ref_id`=VALUES(`ref_id`)"
	}
	r, err := tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.UserID),
		orm.ToSQLString(t.Username),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.Email),
		orm.ToSQLString(t.AvatarURL),
		orm.ToSQLString(t.URL),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.RefID),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBFindOne will find a JiraUser record in the database with the primary key
func (t *JiraUser) DBFindOne(ctx context.Context, db DB, value string) (bool, error) {
	q := "SELECT * FROM `jira_user` WHERE `id` = ? LIMIT 1"
	row := db.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _UserID sql.NullString
	var _Username sql.NullString
	var _Name sql.NullString
	var _Email sql.NullString
	var _AvatarURL sql.NullString
	var _URL sql.NullString
	var _CustomerID sql.NullString
	var _RefID sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_UserID,
		&_Username,
		&_Name,
		&_Email,
		&_AvatarURL,
		&_URL,
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
	if _UserID.Valid {
		t.SetUserID(_UserID.String)
	}
	if _Username.Valid {
		t.SetUsername(_Username.String)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _Email.Valid {
		t.SetEmail(_Email.String)
	}
	if _AvatarURL.Valid {
		t.SetAvatarURL(_AvatarURL.String)
	}
	if _URL.Valid {
		t.SetURL(_URL.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _RefID.Valid {
		t.SetRefID(_RefID.String)
	}
	return true, nil
}

// DBFindOneTx will find a JiraUser record in the database with the primary key using the provided transaction
func (t *JiraUser) DBFindOneTx(ctx context.Context, tx Tx, value string) (bool, error) {
	q := "SELECT * FROM `jira_user` WHERE `id` = ? LIMIT 1"
	row := tx.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _UserID sql.NullString
	var _Username sql.NullString
	var _Name sql.NullString
	var _Email sql.NullString
	var _AvatarURL sql.NullString
	var _URL sql.NullString
	var _CustomerID sql.NullString
	var _RefID sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_UserID,
		&_Username,
		&_Name,
		&_Email,
		&_AvatarURL,
		&_URL,
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
	if _UserID.Valid {
		t.SetUserID(_UserID.String)
	}
	if _Username.Valid {
		t.SetUsername(_Username.String)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _Email.Valid {
		t.SetEmail(_Email.String)
	}
	if _AvatarURL.Valid {
		t.SetAvatarURL(_AvatarURL.String)
	}
	if _URL.Valid {
		t.SetURL(_URL.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _RefID.Valid {
		t.SetRefID(_RefID.String)
	}
	return true, nil
}

// FindJiraUsers will find a JiraUser record in the database with the provided parameters
func FindJiraUsers(ctx context.Context, db DB, _params ...interface{}) ([]*JiraUser, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("user_id"),
		orm.Column("username"),
		orm.Column("name"),
		orm.Column("email"),
		orm.Column("avatar_url"),
		orm.Column("url"),
		orm.Column("customer_id"),
		orm.Column("ref_id"),
		orm.Table(JiraUserTableName),
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
	results := make([]*JiraUser, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _UserID sql.NullString
		var _Username sql.NullString
		var _Name sql.NullString
		var _Email sql.NullString
		var _AvatarURL sql.NullString
		var _URL sql.NullString
		var _CustomerID sql.NullString
		var _RefID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_UserID,
			&_Username,
			&_Name,
			&_Email,
			&_AvatarURL,
			&_URL,
			&_CustomerID,
			&_RefID,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraUser{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _Username.Valid {
			t.SetUsername(_Username.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _Email.Valid {
			t.SetEmail(_Email.String)
		}
		if _AvatarURL.Valid {
			t.SetAvatarURL(_AvatarURL.String)
		}
		if _URL.Valid {
			t.SetURL(_URL.String)
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

// FindJiraUsersTx will find a JiraUser record in the database with the provided parameters using the provided transaction
func FindJiraUsersTx(ctx context.Context, tx Tx, _params ...interface{}) ([]*JiraUser, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("user_id"),
		orm.Column("username"),
		orm.Column("name"),
		orm.Column("email"),
		orm.Column("avatar_url"),
		orm.Column("url"),
		orm.Column("customer_id"),
		orm.Column("ref_id"),
		orm.Table(JiraUserTableName),
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
	results := make([]*JiraUser, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _UserID sql.NullString
		var _Username sql.NullString
		var _Name sql.NullString
		var _Email sql.NullString
		var _AvatarURL sql.NullString
		var _URL sql.NullString
		var _CustomerID sql.NullString
		var _RefID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_UserID,
			&_Username,
			&_Name,
			&_Email,
			&_AvatarURL,
			&_URL,
			&_CustomerID,
			&_RefID,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraUser{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _UserID.Valid {
			t.SetUserID(_UserID.String)
		}
		if _Username.Valid {
			t.SetUsername(_Username.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _Email.Valid {
			t.SetEmail(_Email.String)
		}
		if _AvatarURL.Valid {
			t.SetAvatarURL(_AvatarURL.String)
		}
		if _URL.Valid {
			t.SetURL(_URL.String)
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

// DBFind will find a JiraUser record in the database with the provided parameters
func (t *JiraUser) DBFind(ctx context.Context, db DB, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("user_id"),
		orm.Column("username"),
		orm.Column("name"),
		orm.Column("email"),
		orm.Column("avatar_url"),
		orm.Column("url"),
		orm.Column("customer_id"),
		orm.Column("ref_id"),
		orm.Table(JiraUserTableName),
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
	var _Username sql.NullString
	var _Name sql.NullString
	var _Email sql.NullString
	var _AvatarURL sql.NullString
	var _URL sql.NullString
	var _CustomerID sql.NullString
	var _RefID sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_UserID,
		&_Username,
		&_Name,
		&_Email,
		&_AvatarURL,
		&_URL,
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
	if _UserID.Valid {
		t.SetUserID(_UserID.String)
	}
	if _Username.Valid {
		t.SetUsername(_Username.String)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _Email.Valid {
		t.SetEmail(_Email.String)
	}
	if _AvatarURL.Valid {
		t.SetAvatarURL(_AvatarURL.String)
	}
	if _URL.Valid {
		t.SetURL(_URL.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _RefID.Valid {
		t.SetRefID(_RefID.String)
	}
	return true, nil
}

// DBFindTx will find a JiraUser record in the database with the provided parameters using the provided transaction
func (t *JiraUser) DBFindTx(ctx context.Context, tx Tx, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("user_id"),
		orm.Column("username"),
		orm.Column("name"),
		orm.Column("email"),
		orm.Column("avatar_url"),
		orm.Column("url"),
		orm.Column("customer_id"),
		orm.Column("ref_id"),
		orm.Table(JiraUserTableName),
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
	var _Username sql.NullString
	var _Name sql.NullString
	var _Email sql.NullString
	var _AvatarURL sql.NullString
	var _URL sql.NullString
	var _CustomerID sql.NullString
	var _RefID sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_UserID,
		&_Username,
		&_Name,
		&_Email,
		&_AvatarURL,
		&_URL,
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
	if _UserID.Valid {
		t.SetUserID(_UserID.String)
	}
	if _Username.Valid {
		t.SetUsername(_Username.String)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _Email.Valid {
		t.SetEmail(_Email.String)
	}
	if _AvatarURL.Valid {
		t.SetAvatarURL(_AvatarURL.String)
	}
	if _URL.Valid {
		t.SetURL(_URL.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _RefID.Valid {
		t.SetRefID(_RefID.String)
	}
	return true, nil
}

// CountJiraUsers will find the count of JiraUser records in the database
func CountJiraUsers(ctx context.Context, db DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(JiraUserTableName),
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

// CountJiraUsersTx will find the count of JiraUser records in the database using the provided transaction
func CountJiraUsersTx(ctx context.Context, tx Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(JiraUserTableName),
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

// DBCount will find the count of JiraUser records in the database
func (t *JiraUser) DBCount(ctx context.Context, db DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(JiraUserTableName),
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

// DBCountTx will find the count of JiraUser records in the database using the provided transaction
func (t *JiraUser) DBCountTx(ctx context.Context, tx Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(JiraUserTableName),
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

// DBExists will return true if the JiraUser record exists in the database
func (t *JiraUser) DBExists(ctx context.Context, db DB) (bool, error) {
	q := "SELECT * FROM `jira_user` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := db.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// DBExistsTx will return true if the JiraUser record exists in the database using the provided transaction
func (t *JiraUser) DBExistsTx(ctx context.Context, tx Tx) (bool, error) {
	q := "SELECT * FROM `jira_user` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := tx.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// PrimaryKeyColumn returns the column name for the primary key
func (t *JiraUser) PrimaryKeyColumn() string {
	return JiraUserColumnID
}

// PrimaryKeyColumnType returns the primary key column Go type as a string
func (t *JiraUser) PrimaryKeyColumnType() string {
	return "string"
}

// PrimaryKey returns the primary key column value
func (t *JiraUser) PrimaryKey() interface{} {
	return t.ID
}
