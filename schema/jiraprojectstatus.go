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

var _ Model = (*JiraProjectStatus)(nil)
var _ CSVWriter = (*JiraProjectStatus)(nil)
var _ JSONWriter = (*JiraProjectStatus)(nil)
var _ Checksum = (*JiraProjectStatus)(nil)

// JiraProjectStatusTableName is the name of the table in SQL
const JiraProjectStatusTableName = "jira_project_status"

var JiraProjectStatusColumns = []string{
	"id",
	"checksum",
	"status_id",
	"project_id",
	"name",
	"description",
	"icon_url",
	"category",
	"category_key",
	"category_color",
	"category_id",
	"issue_type_id",
	"customer_id",
}

// JiraProjectStatus table
type JiraProjectStatus struct {
	Category      *string `json:"category,omitempty"`
	CategoryColor *string `json:"category_color,omitempty"`
	CategoryID    *string `json:"category_id,omitempty"`
	CategoryKey   *string `json:"category_key,omitempty"`
	Checksum      *string `json:"checksum,omitempty"`
	CustomerID    string  `json:"customer_id"`
	Description   *string `json:"description,omitempty"`
	IconURL       *string `json:"icon_url,omitempty"`
	ID            string  `json:"id"`
	IssueTypeID   string  `json:"issue_type_id"`
	Name          string  `json:"name"`
	ProjectID     string  `json:"project_id"`
	StatusID      string  `json:"status_id"`
}

// TableName returns the SQL table name for JiraProjectStatus and satifies the Model interface
func (t *JiraProjectStatus) TableName() string {
	return JiraProjectStatusTableName
}

// ToCSV will serialize the JiraProjectStatus instance to a CSV compatible array of strings
func (t *JiraProjectStatus) ToCSV() []string {
	return []string{
		t.ID,
		t.CalculateChecksum(),
		t.StatusID,
		t.ProjectID,
		t.Name,
		toCSVString(t.Description),
		toCSVString(t.IconURL),
		toCSVString(t.Category),
		toCSVString(t.CategoryKey),
		toCSVString(t.CategoryColor),
		toCSVString(t.CategoryID),
		t.IssueTypeID,
		t.CustomerID,
	}
}

// WriteCSV will serialize the JiraProjectStatus instance to the writer as CSV and satisfies the CSVWriter interface
func (t *JiraProjectStatus) WriteCSV(w *csv.Writer) error {
	return w.Write(t.ToCSV())
}

// WriteJSON will serialize the JiraProjectStatus instance to the writer as JSON and satisfies the JSONWriter interface
func (t *JiraProjectStatus) WriteJSON(w io.Writer, indent ...bool) error {
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

// NewJiraProjectStatusReader creates a JSON reader which can read in JiraProjectStatus objects serialized as JSON either as an array, single object or json new lines
// and writes each JiraProjectStatus to the channel provided
func NewJiraProjectStatusReader(r io.Reader, ch chan<- JiraProjectStatus) error {
	return orm.Deserialize(r, func(buf json.RawMessage) error {
		dec := json.NewDecoder(bytes.NewBuffer(buf))
		e := JiraProjectStatus{}
		if err := dec.Decode(&e); err != nil {
			return err
		}
		ch <- e
		return nil
	})
}

// NewCSVJiraProjectStatusReaderDir will read the reader as CSV and emit each record to the channel provided
func NewCSVJiraProjectStatusReader(r io.Reader, ch chan<- JiraProjectStatus) error {
	cr := csv.NewReader(r)
	for {
		record, err := cr.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		ch <- JiraProjectStatus{
			ID:            record[0],
			Checksum:      fromStringPointer(record[1]),
			StatusID:      record[2],
			ProjectID:     record[3],
			Name:          record[4],
			Description:   fromStringPointer(record[5]),
			IconURL:       fromStringPointer(record[6]),
			Category:      fromStringPointer(record[7]),
			CategoryKey:   fromStringPointer(record[8]),
			CategoryColor: fromStringPointer(record[9]),
			CategoryID:    fromStringPointer(record[10]),
			IssueTypeID:   record[11],
			CustomerID:    record[12],
		}
	}
	return nil
}

// NewCSVJiraProjectStatusReaderFile will read the file as a CSV and emit each record to the channel provided
func NewCSVJiraProjectStatusReaderFile(fp string, ch chan<- JiraProjectStatus) error {
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
	return NewCSVJiraProjectStatusReader(fc, ch)
}

// NewCSVJiraProjectStatusReaderDir will read the jira_project_status.csv.gz file as a CSV and emit each record to the channel provided
func NewCSVJiraProjectStatusReaderDir(dir string, ch chan<- JiraProjectStatus) error {
	return NewCSVJiraProjectStatusReaderFile(filepath.Join(dir, "jira_project_status.csv.gz"), ch)
}

// JiraProjectStatusCSVDeduper is a function callback which takes the existing value (a) and the new value (b)
// and the return value should be the one to use (or a new one, if applicable). return nil
// to skip processing of this record
type JiraProjectStatusCSVDeduper func(a JiraProjectStatus, b JiraProjectStatus) *JiraProjectStatus

// JiraProjectStatusCSVDedupeDisabled is set on whether the CSV writer should de-dupe values by key
var JiraProjectStatusCSVDedupeDisabled bool

// NewJiraProjectStatusCSVWriterSize creates a batch writer that will write each JiraProjectStatus into a CSV file
// this method will automatically de-duplicate entries using the primary key. if the checksum
// for a newer item with the same primary key doesn't match a previously sent item, the newer
// one will be used
func NewJiraProjectStatusCSVWriterSize(w io.Writer, size int, dedupers ...JiraProjectStatusCSVDeduper) (chan JiraProjectStatus, chan bool, error) {
	cw := csv.NewWriter(w)
	ch := make(chan JiraProjectStatus, size)
	done := make(chan bool)
	go func() {
		defer func() { done <- true }()
		dodedupe := !JiraProjectStatusCSVDedupeDisabled
		var kv map[string]*JiraProjectStatus
		var deduper JiraProjectStatusCSVDeduper
		if dedupers != nil && len(dedupers) > 0 {
			deduper = dedupers[0]
			dodedupe = true
		}
		if dodedupe {
			kv = make(map[string]*JiraProjectStatus)
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

// JiraProjectStatusCSVDefaultSize is the default channel buffer size if not provided
var JiraProjectStatusCSVDefaultSize = 100

// NewJiraProjectStatusCSVWriter creates a batch writer that will write each JiraProjectStatus into a CSV file
func NewJiraProjectStatusCSVWriter(w io.Writer, dedupers ...JiraProjectStatusCSVDeduper) (chan JiraProjectStatus, chan bool, error) {
	return NewJiraProjectStatusCSVWriterSize(w, JiraProjectStatusCSVDefaultSize, dedupers...)
}

// NewJiraProjectStatusCSVWriterDir creates a batch writer that will write each JiraProjectStatus into a CSV file named jira_project_status.csv.gz in dir
func NewJiraProjectStatusCSVWriterDir(dir string, dedupers ...JiraProjectStatusCSVDeduper) (chan JiraProjectStatus, chan bool, error) {
	return NewJiraProjectStatusCSVWriterFile(filepath.Join(dir, "jira_project_status.csv.gz"), dedupers...)
}

// NewJiraProjectStatusCSVWriterFile creates a batch writer that will write each JiraProjectStatus into a CSV file
func NewJiraProjectStatusCSVWriterFile(fn string, dedupers ...JiraProjectStatusCSVDeduper) (chan JiraProjectStatus, chan bool, error) {
	f, err := os.Create(fn)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening CSV file %s. %v", fn, err)
	}
	var fc io.WriteCloser = f
	if filepath.Ext(fn) == ".gz" {
		w, _ := gzip.NewWriterLevel(f, gzip.BestCompression)
		fc = w
	}
	ch, done, err := NewJiraProjectStatusCSVWriter(fc, dedupers...)
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

type JiraProjectStatusDBAction func(ctx context.Context, db *sql.DB, record JiraProjectStatus) error

// NewJiraProjectStatusDBWriterSize creates a DB writer that will write each issue into the DB
func NewJiraProjectStatusDBWriterSize(ctx context.Context, db *sql.DB, errors chan<- error, size int, actions ...JiraProjectStatusDBAction) (chan JiraProjectStatus, chan bool, error) {
	ch := make(chan JiraProjectStatus, size)
	done := make(chan bool)
	var action JiraProjectStatusDBAction
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

// NewJiraProjectStatusDBWriter creates a DB writer that will write each issue into the DB
func NewJiraProjectStatusDBWriter(ctx context.Context, db *sql.DB, errors chan<- error, actions ...JiraProjectStatusDBAction) (chan JiraProjectStatus, chan bool, error) {
	return NewJiraProjectStatusDBWriterSize(ctx, db, errors, 100, actions...)
}

// JiraProjectStatusColumnID is the ID SQL column name for the JiraProjectStatus table
const JiraProjectStatusColumnID = "id"

// JiraProjectStatusEscapedColumnID is the escaped ID SQL column name for the JiraProjectStatus table
const JiraProjectStatusEscapedColumnID = "`id`"

// JiraProjectStatusColumnChecksum is the Checksum SQL column name for the JiraProjectStatus table
const JiraProjectStatusColumnChecksum = "checksum"

// JiraProjectStatusEscapedColumnChecksum is the escaped Checksum SQL column name for the JiraProjectStatus table
const JiraProjectStatusEscapedColumnChecksum = "`checksum`"

// JiraProjectStatusColumnStatusID is the StatusID SQL column name for the JiraProjectStatus table
const JiraProjectStatusColumnStatusID = "status_id"

// JiraProjectStatusEscapedColumnStatusID is the escaped StatusID SQL column name for the JiraProjectStatus table
const JiraProjectStatusEscapedColumnStatusID = "`status_id`"

// JiraProjectStatusColumnProjectID is the ProjectID SQL column name for the JiraProjectStatus table
const JiraProjectStatusColumnProjectID = "project_id"

// JiraProjectStatusEscapedColumnProjectID is the escaped ProjectID SQL column name for the JiraProjectStatus table
const JiraProjectStatusEscapedColumnProjectID = "`project_id`"

// JiraProjectStatusColumnName is the Name SQL column name for the JiraProjectStatus table
const JiraProjectStatusColumnName = "name"

// JiraProjectStatusEscapedColumnName is the escaped Name SQL column name for the JiraProjectStatus table
const JiraProjectStatusEscapedColumnName = "`name`"

// JiraProjectStatusColumnDescription is the Description SQL column name for the JiraProjectStatus table
const JiraProjectStatusColumnDescription = "description"

// JiraProjectStatusEscapedColumnDescription is the escaped Description SQL column name for the JiraProjectStatus table
const JiraProjectStatusEscapedColumnDescription = "`description`"

// JiraProjectStatusColumnIconURL is the IconURL SQL column name for the JiraProjectStatus table
const JiraProjectStatusColumnIconURL = "icon_url"

// JiraProjectStatusEscapedColumnIconURL is the escaped IconURL SQL column name for the JiraProjectStatus table
const JiraProjectStatusEscapedColumnIconURL = "`icon_url`"

// JiraProjectStatusColumnCategory is the Category SQL column name for the JiraProjectStatus table
const JiraProjectStatusColumnCategory = "category"

// JiraProjectStatusEscapedColumnCategory is the escaped Category SQL column name for the JiraProjectStatus table
const JiraProjectStatusEscapedColumnCategory = "`category`"

// JiraProjectStatusColumnCategoryKey is the CategoryKey SQL column name for the JiraProjectStatus table
const JiraProjectStatusColumnCategoryKey = "category_key"

// JiraProjectStatusEscapedColumnCategoryKey is the escaped CategoryKey SQL column name for the JiraProjectStatus table
const JiraProjectStatusEscapedColumnCategoryKey = "`category_key`"

// JiraProjectStatusColumnCategoryColor is the CategoryColor SQL column name for the JiraProjectStatus table
const JiraProjectStatusColumnCategoryColor = "category_color"

// JiraProjectStatusEscapedColumnCategoryColor is the escaped CategoryColor SQL column name for the JiraProjectStatus table
const JiraProjectStatusEscapedColumnCategoryColor = "`category_color`"

// JiraProjectStatusColumnCategoryID is the CategoryID SQL column name for the JiraProjectStatus table
const JiraProjectStatusColumnCategoryID = "category_id"

// JiraProjectStatusEscapedColumnCategoryID is the escaped CategoryID SQL column name for the JiraProjectStatus table
const JiraProjectStatusEscapedColumnCategoryID = "`category_id`"

// JiraProjectStatusColumnIssueTypeID is the IssueTypeID SQL column name for the JiraProjectStatus table
const JiraProjectStatusColumnIssueTypeID = "issue_type_id"

// JiraProjectStatusEscapedColumnIssueTypeID is the escaped IssueTypeID SQL column name for the JiraProjectStatus table
const JiraProjectStatusEscapedColumnIssueTypeID = "`issue_type_id`"

// JiraProjectStatusColumnCustomerID is the CustomerID SQL column name for the JiraProjectStatus table
const JiraProjectStatusColumnCustomerID = "customer_id"

// JiraProjectStatusEscapedColumnCustomerID is the escaped CustomerID SQL column name for the JiraProjectStatus table
const JiraProjectStatusEscapedColumnCustomerID = "`customer_id`"

// GetID will return the JiraProjectStatus ID value
func (t *JiraProjectStatus) GetID() string {
	return t.ID
}

// SetID will set the JiraProjectStatus ID value
func (t *JiraProjectStatus) SetID(v string) {
	t.ID = v
}

// FindJiraProjectStatusByID will find a JiraProjectStatus by ID
func FindJiraProjectStatusByID(ctx context.Context, db *sql.DB, value string) (*JiraProjectStatus, error) {
	q := "SELECT `jira_project_status`.`id`,`jira_project_status`.`checksum`,`jira_project_status`.`status_id`,`jira_project_status`.`project_id`,`jira_project_status`.`name`,`jira_project_status`.`description`,`jira_project_status`.`icon_url`,`jira_project_status`.`category`,`jira_project_status`.`category_key`,`jira_project_status`.`category_color`,`jira_project_status`.`category_id`,`jira_project_status`.`issue_type_id`,`jira_project_status`.`customer_id` FROM `jira_project_status` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _StatusID sql.NullString
	var _ProjectID sql.NullString
	var _Name sql.NullString
	var _Description sql.NullString
	var _IconURL sql.NullString
	var _Category sql.NullString
	var _CategoryKey sql.NullString
	var _CategoryColor sql.NullString
	var _CategoryID sql.NullString
	var _IssueTypeID sql.NullString
	var _CustomerID sql.NullString
	err := db.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_StatusID,
		&_ProjectID,
		&_Name,
		&_Description,
		&_IconURL,
		&_Category,
		&_CategoryKey,
		&_CategoryColor,
		&_CategoryID,
		&_IssueTypeID,
		&_CustomerID,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &JiraProjectStatus{}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _Checksum.Valid {
		t.SetChecksum(_Checksum.String)
	}
	if _StatusID.Valid {
		t.SetStatusID(_StatusID.String)
	}
	if _ProjectID.Valid {
		t.SetProjectID(_ProjectID.String)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _Description.Valid {
		t.SetDescription(_Description.String)
	}
	if _IconURL.Valid {
		t.SetIconURL(_IconURL.String)
	}
	if _Category.Valid {
		t.SetCategory(_Category.String)
	}
	if _CategoryKey.Valid {
		t.SetCategoryKey(_CategoryKey.String)
	}
	if _CategoryColor.Valid {
		t.SetCategoryColor(_CategoryColor.String)
	}
	if _CategoryID.Valid {
		t.SetCategoryID(_CategoryID.String)
	}
	if _IssueTypeID.Valid {
		t.SetIssueTypeID(_IssueTypeID.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	return t, nil
}

// FindJiraProjectStatusByIDTx will find a JiraProjectStatus by ID using the provided transaction
func FindJiraProjectStatusByIDTx(ctx context.Context, tx *sql.Tx, value string) (*JiraProjectStatus, error) {
	q := "SELECT `jira_project_status`.`id`,`jira_project_status`.`checksum`,`jira_project_status`.`status_id`,`jira_project_status`.`project_id`,`jira_project_status`.`name`,`jira_project_status`.`description`,`jira_project_status`.`icon_url`,`jira_project_status`.`category`,`jira_project_status`.`category_key`,`jira_project_status`.`category_color`,`jira_project_status`.`category_id`,`jira_project_status`.`issue_type_id`,`jira_project_status`.`customer_id` FROM `jira_project_status` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _StatusID sql.NullString
	var _ProjectID sql.NullString
	var _Name sql.NullString
	var _Description sql.NullString
	var _IconURL sql.NullString
	var _Category sql.NullString
	var _CategoryKey sql.NullString
	var _CategoryColor sql.NullString
	var _CategoryID sql.NullString
	var _IssueTypeID sql.NullString
	var _CustomerID sql.NullString
	err := tx.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_StatusID,
		&_ProjectID,
		&_Name,
		&_Description,
		&_IconURL,
		&_Category,
		&_CategoryKey,
		&_CategoryColor,
		&_CategoryID,
		&_IssueTypeID,
		&_CustomerID,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &JiraProjectStatus{}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _Checksum.Valid {
		t.SetChecksum(_Checksum.String)
	}
	if _StatusID.Valid {
		t.SetStatusID(_StatusID.String)
	}
	if _ProjectID.Valid {
		t.SetProjectID(_ProjectID.String)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _Description.Valid {
		t.SetDescription(_Description.String)
	}
	if _IconURL.Valid {
		t.SetIconURL(_IconURL.String)
	}
	if _Category.Valid {
		t.SetCategory(_Category.String)
	}
	if _CategoryKey.Valid {
		t.SetCategoryKey(_CategoryKey.String)
	}
	if _CategoryColor.Valid {
		t.SetCategoryColor(_CategoryColor.String)
	}
	if _CategoryID.Valid {
		t.SetCategoryID(_CategoryID.String)
	}
	if _IssueTypeID.Valid {
		t.SetIssueTypeID(_IssueTypeID.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	return t, nil
}

// GetChecksum will return the JiraProjectStatus Checksum value
func (t *JiraProjectStatus) GetChecksum() string {
	if t.Checksum == nil {
		return ""
	}
	return *t.Checksum
}

// SetChecksum will set the JiraProjectStatus Checksum value
func (t *JiraProjectStatus) SetChecksum(v string) {
	t.Checksum = &v
}

// GetStatusID will return the JiraProjectStatus StatusID value
func (t *JiraProjectStatus) GetStatusID() string {
	return t.StatusID
}

// SetStatusID will set the JiraProjectStatus StatusID value
func (t *JiraProjectStatus) SetStatusID(v string) {
	t.StatusID = v
}

// GetProjectID will return the JiraProjectStatus ProjectID value
func (t *JiraProjectStatus) GetProjectID() string {
	return t.ProjectID
}

// SetProjectID will set the JiraProjectStatus ProjectID value
func (t *JiraProjectStatus) SetProjectID(v string) {
	t.ProjectID = v
}

// FindJiraProjectStatusesByProjectID will find all JiraProjectStatuss by the ProjectID value
func FindJiraProjectStatusesByProjectID(ctx context.Context, db *sql.DB, value string) ([]*JiraProjectStatus, error) {
	q := "SELECT `jira_project_status`.`id`,`jira_project_status`.`checksum`,`jira_project_status`.`status_id`,`jira_project_status`.`project_id`,`jira_project_status`.`name`,`jira_project_status`.`description`,`jira_project_status`.`icon_url`,`jira_project_status`.`category`,`jira_project_status`.`category_key`,`jira_project_status`.`category_color`,`jira_project_status`.`category_id`,`jira_project_status`.`issue_type_id`,`jira_project_status`.`customer_id` FROM `jira_project_status` WHERE `project_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraProjectStatus, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _StatusID sql.NullString
		var _ProjectID sql.NullString
		var _Name sql.NullString
		var _Description sql.NullString
		var _IconURL sql.NullString
		var _Category sql.NullString
		var _CategoryKey sql.NullString
		var _CategoryColor sql.NullString
		var _CategoryID sql.NullString
		var _IssueTypeID sql.NullString
		var _CustomerID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_StatusID,
			&_ProjectID,
			&_Name,
			&_Description,
			&_IconURL,
			&_Category,
			&_CategoryKey,
			&_CategoryColor,
			&_CategoryID,
			&_IssueTypeID,
			&_CustomerID,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraProjectStatus{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _StatusID.Valid {
			t.SetStatusID(_StatusID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _Description.Valid {
			t.SetDescription(_Description.String)
		}
		if _IconURL.Valid {
			t.SetIconURL(_IconURL.String)
		}
		if _Category.Valid {
			t.SetCategory(_Category.String)
		}
		if _CategoryKey.Valid {
			t.SetCategoryKey(_CategoryKey.String)
		}
		if _CategoryColor.Valid {
			t.SetCategoryColor(_CategoryColor.String)
		}
		if _CategoryID.Valid {
			t.SetCategoryID(_CategoryID.String)
		}
		if _IssueTypeID.Valid {
			t.SetIssueTypeID(_IssueTypeID.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindJiraProjectStatusesByProjectIDTx will find all JiraProjectStatuss by the ProjectID value using the provided transaction
func FindJiraProjectStatusesByProjectIDTx(ctx context.Context, tx *sql.Tx, value string) ([]*JiraProjectStatus, error) {
	q := "SELECT `jira_project_status`.`id`,`jira_project_status`.`checksum`,`jira_project_status`.`status_id`,`jira_project_status`.`project_id`,`jira_project_status`.`name`,`jira_project_status`.`description`,`jira_project_status`.`icon_url`,`jira_project_status`.`category`,`jira_project_status`.`category_key`,`jira_project_status`.`category_color`,`jira_project_status`.`category_id`,`jira_project_status`.`issue_type_id`,`jira_project_status`.`customer_id` FROM `jira_project_status` WHERE `project_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraProjectStatus, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _StatusID sql.NullString
		var _ProjectID sql.NullString
		var _Name sql.NullString
		var _Description sql.NullString
		var _IconURL sql.NullString
		var _Category sql.NullString
		var _CategoryKey sql.NullString
		var _CategoryColor sql.NullString
		var _CategoryID sql.NullString
		var _IssueTypeID sql.NullString
		var _CustomerID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_StatusID,
			&_ProjectID,
			&_Name,
			&_Description,
			&_IconURL,
			&_Category,
			&_CategoryKey,
			&_CategoryColor,
			&_CategoryID,
			&_IssueTypeID,
			&_CustomerID,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraProjectStatus{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _StatusID.Valid {
			t.SetStatusID(_StatusID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _Description.Valid {
			t.SetDescription(_Description.String)
		}
		if _IconURL.Valid {
			t.SetIconURL(_IconURL.String)
		}
		if _Category.Valid {
			t.SetCategory(_Category.String)
		}
		if _CategoryKey.Valid {
			t.SetCategoryKey(_CategoryKey.String)
		}
		if _CategoryColor.Valid {
			t.SetCategoryColor(_CategoryColor.String)
		}
		if _CategoryID.Valid {
			t.SetCategoryID(_CategoryID.String)
		}
		if _IssueTypeID.Valid {
			t.SetIssueTypeID(_IssueTypeID.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetName will return the JiraProjectStatus Name value
func (t *JiraProjectStatus) GetName() string {
	return t.Name
}

// SetName will set the JiraProjectStatus Name value
func (t *JiraProjectStatus) SetName(v string) {
	t.Name = v
}

// FindJiraProjectStatusesByName will find all JiraProjectStatuss by the Name value
func FindJiraProjectStatusesByName(ctx context.Context, db *sql.DB, value string) ([]*JiraProjectStatus, error) {
	q := "SELECT `jira_project_status`.`id`,`jira_project_status`.`checksum`,`jira_project_status`.`status_id`,`jira_project_status`.`project_id`,`jira_project_status`.`name`,`jira_project_status`.`description`,`jira_project_status`.`icon_url`,`jira_project_status`.`category`,`jira_project_status`.`category_key`,`jira_project_status`.`category_color`,`jira_project_status`.`category_id`,`jira_project_status`.`issue_type_id`,`jira_project_status`.`customer_id` FROM `jira_project_status` WHERE `name` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraProjectStatus, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _StatusID sql.NullString
		var _ProjectID sql.NullString
		var _Name sql.NullString
		var _Description sql.NullString
		var _IconURL sql.NullString
		var _Category sql.NullString
		var _CategoryKey sql.NullString
		var _CategoryColor sql.NullString
		var _CategoryID sql.NullString
		var _IssueTypeID sql.NullString
		var _CustomerID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_StatusID,
			&_ProjectID,
			&_Name,
			&_Description,
			&_IconURL,
			&_Category,
			&_CategoryKey,
			&_CategoryColor,
			&_CategoryID,
			&_IssueTypeID,
			&_CustomerID,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraProjectStatus{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _StatusID.Valid {
			t.SetStatusID(_StatusID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _Description.Valid {
			t.SetDescription(_Description.String)
		}
		if _IconURL.Valid {
			t.SetIconURL(_IconURL.String)
		}
		if _Category.Valid {
			t.SetCategory(_Category.String)
		}
		if _CategoryKey.Valid {
			t.SetCategoryKey(_CategoryKey.String)
		}
		if _CategoryColor.Valid {
			t.SetCategoryColor(_CategoryColor.String)
		}
		if _CategoryID.Valid {
			t.SetCategoryID(_CategoryID.String)
		}
		if _IssueTypeID.Valid {
			t.SetIssueTypeID(_IssueTypeID.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindJiraProjectStatusesByNameTx will find all JiraProjectStatuss by the Name value using the provided transaction
func FindJiraProjectStatusesByNameTx(ctx context.Context, tx *sql.Tx, value string) ([]*JiraProjectStatus, error) {
	q := "SELECT `jira_project_status`.`id`,`jira_project_status`.`checksum`,`jira_project_status`.`status_id`,`jira_project_status`.`project_id`,`jira_project_status`.`name`,`jira_project_status`.`description`,`jira_project_status`.`icon_url`,`jira_project_status`.`category`,`jira_project_status`.`category_key`,`jira_project_status`.`category_color`,`jira_project_status`.`category_id`,`jira_project_status`.`issue_type_id`,`jira_project_status`.`customer_id` FROM `jira_project_status` WHERE `name` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraProjectStatus, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _StatusID sql.NullString
		var _ProjectID sql.NullString
		var _Name sql.NullString
		var _Description sql.NullString
		var _IconURL sql.NullString
		var _Category sql.NullString
		var _CategoryKey sql.NullString
		var _CategoryColor sql.NullString
		var _CategoryID sql.NullString
		var _IssueTypeID sql.NullString
		var _CustomerID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_StatusID,
			&_ProjectID,
			&_Name,
			&_Description,
			&_IconURL,
			&_Category,
			&_CategoryKey,
			&_CategoryColor,
			&_CategoryID,
			&_IssueTypeID,
			&_CustomerID,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraProjectStatus{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _StatusID.Valid {
			t.SetStatusID(_StatusID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _Description.Valid {
			t.SetDescription(_Description.String)
		}
		if _IconURL.Valid {
			t.SetIconURL(_IconURL.String)
		}
		if _Category.Valid {
			t.SetCategory(_Category.String)
		}
		if _CategoryKey.Valid {
			t.SetCategoryKey(_CategoryKey.String)
		}
		if _CategoryColor.Valid {
			t.SetCategoryColor(_CategoryColor.String)
		}
		if _CategoryID.Valid {
			t.SetCategoryID(_CategoryID.String)
		}
		if _IssueTypeID.Valid {
			t.SetIssueTypeID(_IssueTypeID.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetDescription will return the JiraProjectStatus Description value
func (t *JiraProjectStatus) GetDescription() string {
	if t.Description == nil {
		return ""
	}
	return *t.Description
}

// SetDescription will set the JiraProjectStatus Description value
func (t *JiraProjectStatus) SetDescription(v string) {
	t.Description = &v
}

// GetIconURL will return the JiraProjectStatus IconURL value
func (t *JiraProjectStatus) GetIconURL() string {
	if t.IconURL == nil {
		return ""
	}
	return *t.IconURL
}

// SetIconURL will set the JiraProjectStatus IconURL value
func (t *JiraProjectStatus) SetIconURL(v string) {
	t.IconURL = &v
}

// GetCategory will return the JiraProjectStatus Category value
func (t *JiraProjectStatus) GetCategory() string {
	if t.Category == nil {
		return ""
	}
	return *t.Category
}

// SetCategory will set the JiraProjectStatus Category value
func (t *JiraProjectStatus) SetCategory(v string) {
	t.Category = &v
}

// GetCategoryKey will return the JiraProjectStatus CategoryKey value
func (t *JiraProjectStatus) GetCategoryKey() string {
	if t.CategoryKey == nil {
		return ""
	}
	return *t.CategoryKey
}

// SetCategoryKey will set the JiraProjectStatus CategoryKey value
func (t *JiraProjectStatus) SetCategoryKey(v string) {
	t.CategoryKey = &v
}

// GetCategoryColor will return the JiraProjectStatus CategoryColor value
func (t *JiraProjectStatus) GetCategoryColor() string {
	if t.CategoryColor == nil {
		return ""
	}
	return *t.CategoryColor
}

// SetCategoryColor will set the JiraProjectStatus CategoryColor value
func (t *JiraProjectStatus) SetCategoryColor(v string) {
	t.CategoryColor = &v
}

// GetCategoryID will return the JiraProjectStatus CategoryID value
func (t *JiraProjectStatus) GetCategoryID() string {
	if t.CategoryID == nil {
		return ""
	}
	return *t.CategoryID
}

// SetCategoryID will set the JiraProjectStatus CategoryID value
func (t *JiraProjectStatus) SetCategoryID(v string) {
	t.CategoryID = &v
}

// GetIssueTypeID will return the JiraProjectStatus IssueTypeID value
func (t *JiraProjectStatus) GetIssueTypeID() string {
	return t.IssueTypeID
}

// SetIssueTypeID will set the JiraProjectStatus IssueTypeID value
func (t *JiraProjectStatus) SetIssueTypeID(v string) {
	t.IssueTypeID = v
}

// GetCustomerID will return the JiraProjectStatus CustomerID value
func (t *JiraProjectStatus) GetCustomerID() string {
	return t.CustomerID
}

// SetCustomerID will set the JiraProjectStatus CustomerID value
func (t *JiraProjectStatus) SetCustomerID(v string) {
	t.CustomerID = v
}

// FindJiraProjectStatusesByCustomerID will find all JiraProjectStatuss by the CustomerID value
func FindJiraProjectStatusesByCustomerID(ctx context.Context, db *sql.DB, value string) ([]*JiraProjectStatus, error) {
	q := "SELECT `jira_project_status`.`id`,`jira_project_status`.`checksum`,`jira_project_status`.`status_id`,`jira_project_status`.`project_id`,`jira_project_status`.`name`,`jira_project_status`.`description`,`jira_project_status`.`icon_url`,`jira_project_status`.`category`,`jira_project_status`.`category_key`,`jira_project_status`.`category_color`,`jira_project_status`.`category_id`,`jira_project_status`.`issue_type_id`,`jira_project_status`.`customer_id` FROM `jira_project_status` WHERE `customer_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraProjectStatus, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _StatusID sql.NullString
		var _ProjectID sql.NullString
		var _Name sql.NullString
		var _Description sql.NullString
		var _IconURL sql.NullString
		var _Category sql.NullString
		var _CategoryKey sql.NullString
		var _CategoryColor sql.NullString
		var _CategoryID sql.NullString
		var _IssueTypeID sql.NullString
		var _CustomerID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_StatusID,
			&_ProjectID,
			&_Name,
			&_Description,
			&_IconURL,
			&_Category,
			&_CategoryKey,
			&_CategoryColor,
			&_CategoryID,
			&_IssueTypeID,
			&_CustomerID,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraProjectStatus{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _StatusID.Valid {
			t.SetStatusID(_StatusID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _Description.Valid {
			t.SetDescription(_Description.String)
		}
		if _IconURL.Valid {
			t.SetIconURL(_IconURL.String)
		}
		if _Category.Valid {
			t.SetCategory(_Category.String)
		}
		if _CategoryKey.Valid {
			t.SetCategoryKey(_CategoryKey.String)
		}
		if _CategoryColor.Valid {
			t.SetCategoryColor(_CategoryColor.String)
		}
		if _CategoryID.Valid {
			t.SetCategoryID(_CategoryID.String)
		}
		if _IssueTypeID.Valid {
			t.SetIssueTypeID(_IssueTypeID.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindJiraProjectStatusesByCustomerIDTx will find all JiraProjectStatuss by the CustomerID value using the provided transaction
func FindJiraProjectStatusesByCustomerIDTx(ctx context.Context, tx *sql.Tx, value string) ([]*JiraProjectStatus, error) {
	q := "SELECT `jira_project_status`.`id`,`jira_project_status`.`checksum`,`jira_project_status`.`status_id`,`jira_project_status`.`project_id`,`jira_project_status`.`name`,`jira_project_status`.`description`,`jira_project_status`.`icon_url`,`jira_project_status`.`category`,`jira_project_status`.`category_key`,`jira_project_status`.`category_color`,`jira_project_status`.`category_id`,`jira_project_status`.`issue_type_id`,`jira_project_status`.`customer_id` FROM `jira_project_status` WHERE `customer_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*JiraProjectStatus, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _StatusID sql.NullString
		var _ProjectID sql.NullString
		var _Name sql.NullString
		var _Description sql.NullString
		var _IconURL sql.NullString
		var _Category sql.NullString
		var _CategoryKey sql.NullString
		var _CategoryColor sql.NullString
		var _CategoryID sql.NullString
		var _IssueTypeID sql.NullString
		var _CustomerID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_StatusID,
			&_ProjectID,
			&_Name,
			&_Description,
			&_IconURL,
			&_Category,
			&_CategoryKey,
			&_CategoryColor,
			&_CategoryID,
			&_IssueTypeID,
			&_CustomerID,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraProjectStatus{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _StatusID.Valid {
			t.SetStatusID(_StatusID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _Description.Valid {
			t.SetDescription(_Description.String)
		}
		if _IconURL.Valid {
			t.SetIconURL(_IconURL.String)
		}
		if _Category.Valid {
			t.SetCategory(_Category.String)
		}
		if _CategoryKey.Valid {
			t.SetCategoryKey(_CategoryKey.String)
		}
		if _CategoryColor.Valid {
			t.SetCategoryColor(_CategoryColor.String)
		}
		if _CategoryID.Valid {
			t.SetCategoryID(_CategoryID.String)
		}
		if _IssueTypeID.Valid {
			t.SetIssueTypeID(_IssueTypeID.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

func (t *JiraProjectStatus) toTimestamp(value time.Time) *timestamp.Timestamp {
	ts, _ := ptypes.TimestampProto(value)
	return ts
}

// DBCreateJiraProjectStatusTable will create the JiraProjectStatus table
func DBCreateJiraProjectStatusTable(ctx context.Context, db *sql.DB) error {
	q := "CREATE TABLE `jira_project_status` (`id` VARCHAR(64) NOT NULL PRIMARY KEY,`checksum` CHAR(64),`status_id` VARCHAR(64) NOT NULL,`project_id`VARCHAR(64) NOT NULL,`name`VARCHAR(255) NOT NULL,`description` TEXT,`icon_url` TEXT,`category` TEXT,`category_key` TEXT,`category_color` TEXT,`category_id` VARCHAR(64),`issue_type_id`VARCHAR(64) NOT NULL,`customer_id` VARCHAR(64) NOT NULL,INDEX jira_project_status_project_id_index (`project_id`),INDEX jira_project_status_name_index (`name`),INDEX jira_project_status_customer_id_index (`customer_id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBCreateJiraProjectStatusTableTx will create the JiraProjectStatus table using the provided transction
func DBCreateJiraProjectStatusTableTx(ctx context.Context, tx *sql.Tx) error {
	q := "CREATE TABLE `jira_project_status` (`id` VARCHAR(64) NOT NULL PRIMARY KEY,`checksum` CHAR(64),`status_id` VARCHAR(64) NOT NULL,`project_id`VARCHAR(64) NOT NULL,`name`VARCHAR(255) NOT NULL,`description` TEXT,`icon_url` TEXT,`category` TEXT,`category_key` TEXT,`category_color` TEXT,`category_id` VARCHAR(64),`issue_type_id`VARCHAR(64) NOT NULL,`customer_id` VARCHAR(64) NOT NULL,INDEX jira_project_status_project_id_index (`project_id`),INDEX jira_project_status_name_index (`name`),INDEX jira_project_status_customer_id_index (`customer_id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// DBDropJiraProjectStatusTable will drop the JiraProjectStatus table
func DBDropJiraProjectStatusTable(ctx context.Context, db *sql.DB) error {
	q := "DROP TABLE IF EXISTS `jira_project_status`"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBDropJiraProjectStatusTableTx will drop the JiraProjectStatus table using the provided transaction
func DBDropJiraProjectStatusTableTx(ctx context.Context, tx *sql.Tx) error {
	q := "DROP TABLE IF EXISTS `jira_project_status`"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// CalculateChecksum will calculate a checksum of the SHA1 of all field values
func (t *JiraProjectStatus) CalculateChecksum() string {
	return orm.HashStrings(
		orm.ToString(t.ID),
		orm.ToString(t.StatusID),
		orm.ToString(t.ProjectID),
		orm.ToString(t.Name),
		orm.ToString(t.Description),
		orm.ToString(t.IconURL),
		orm.ToString(t.Category),
		orm.ToString(t.CategoryKey),
		orm.ToString(t.CategoryColor),
		orm.ToString(t.CategoryID),
		orm.ToString(t.IssueTypeID),
		orm.ToString(t.CustomerID),
	)
}

// DBCreate will create a new JiraProjectStatus record in the database
func (t *JiraProjectStatus) DBCreate(ctx context.Context, db *sql.DB) (sql.Result, error) {
	q := "INSERT INTO `jira_project_status` (`jira_project_status`.`id`,`jira_project_status`.`checksum`,`jira_project_status`.`status_id`,`jira_project_status`.`project_id`,`jira_project_status`.`name`,`jira_project_status`.`description`,`jira_project_status`.`icon_url`,`jira_project_status`.`category`,`jira_project_status`.`category_key`,`jira_project_status`.`category_color`,`jira_project_status`.`category_id`,`jira_project_status`.`issue_type_id`,`jira_project_status`.`customer_id`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?)"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.StatusID),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.Description),
		orm.ToSQLString(t.IconURL),
		orm.ToSQLString(t.Category),
		orm.ToSQLString(t.CategoryKey),
		orm.ToSQLString(t.CategoryColor),
		orm.ToSQLString(t.CategoryID),
		orm.ToSQLString(t.IssueTypeID),
		orm.ToSQLString(t.CustomerID),
	)
}

// DBCreateTx will create a new JiraProjectStatus record in the database using the provided transaction
func (t *JiraProjectStatus) DBCreateTx(ctx context.Context, tx *sql.Tx) (sql.Result, error) {
	q := "INSERT INTO `jira_project_status` (`jira_project_status`.`id`,`jira_project_status`.`checksum`,`jira_project_status`.`status_id`,`jira_project_status`.`project_id`,`jira_project_status`.`name`,`jira_project_status`.`description`,`jira_project_status`.`icon_url`,`jira_project_status`.`category`,`jira_project_status`.`category_key`,`jira_project_status`.`category_color`,`jira_project_status`.`category_id`,`jira_project_status`.`issue_type_id`,`jira_project_status`.`customer_id`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?)"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.StatusID),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.Description),
		orm.ToSQLString(t.IconURL),
		orm.ToSQLString(t.Category),
		orm.ToSQLString(t.CategoryKey),
		orm.ToSQLString(t.CategoryColor),
		orm.ToSQLString(t.CategoryID),
		orm.ToSQLString(t.IssueTypeID),
		orm.ToSQLString(t.CustomerID),
	)
}

// DBCreateIgnoreDuplicate will upsert the JiraProjectStatus record in the database
func (t *JiraProjectStatus) DBCreateIgnoreDuplicate(ctx context.Context, db *sql.DB) (sql.Result, error) {
	q := "INSERT INTO `jira_project_status` (`jira_project_status`.`id`,`jira_project_status`.`checksum`,`jira_project_status`.`status_id`,`jira_project_status`.`project_id`,`jira_project_status`.`name`,`jira_project_status`.`description`,`jira_project_status`.`icon_url`,`jira_project_status`.`category`,`jira_project_status`.`category_key`,`jira_project_status`.`category_color`,`jira_project_status`.`category_id`,`jira_project_status`.`issue_type_id`,`jira_project_status`.`customer_id`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.StatusID),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.Description),
		orm.ToSQLString(t.IconURL),
		orm.ToSQLString(t.Category),
		orm.ToSQLString(t.CategoryKey),
		orm.ToSQLString(t.CategoryColor),
		orm.ToSQLString(t.CategoryID),
		orm.ToSQLString(t.IssueTypeID),
		orm.ToSQLString(t.CustomerID),
	)
}

// DBCreateIgnoreDuplicateTx will upsert the JiraProjectStatus record in the database using the provided transaction
func (t *JiraProjectStatus) DBCreateIgnoreDuplicateTx(ctx context.Context, tx *sql.Tx) (sql.Result, error) {
	q := "INSERT INTO `jira_project_status` (`jira_project_status`.`id`,`jira_project_status`.`checksum`,`jira_project_status`.`status_id`,`jira_project_status`.`project_id`,`jira_project_status`.`name`,`jira_project_status`.`description`,`jira_project_status`.`icon_url`,`jira_project_status`.`category`,`jira_project_status`.`category_key`,`jira_project_status`.`category_color`,`jira_project_status`.`category_id`,`jira_project_status`.`issue_type_id`,`jira_project_status`.`customer_id`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.StatusID),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.Description),
		orm.ToSQLString(t.IconURL),
		orm.ToSQLString(t.Category),
		orm.ToSQLString(t.CategoryKey),
		orm.ToSQLString(t.CategoryColor),
		orm.ToSQLString(t.CategoryID),
		orm.ToSQLString(t.IssueTypeID),
		orm.ToSQLString(t.CustomerID),
	)
}

// DeleteAllJiraProjectStatuses deletes all JiraProjectStatus records in the database with optional filters
func DeleteAllJiraProjectStatuses(ctx context.Context, db *sql.DB, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(JiraProjectStatusTableName),
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

// DeleteAllJiraProjectStatusesTx deletes all JiraProjectStatus records in the database with optional filters using the provided transaction
func DeleteAllJiraProjectStatusesTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(JiraProjectStatusTableName),
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

// DBDelete will delete this JiraProjectStatus record in the database
func (t *JiraProjectStatus) DBDelete(ctx context.Context, db *sql.DB) (bool, error) {
	q := "DELETE FROM `jira_project_status` WHERE `id` = ?"
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

// DBDeleteTx will delete this JiraProjectStatus record in the database using the provided transaction
func (t *JiraProjectStatus) DBDeleteTx(ctx context.Context, tx *sql.Tx) (bool, error) {
	q := "DELETE FROM `jira_project_status` WHERE `id` = ?"
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

// DBUpdate will update the JiraProjectStatus record in the database
func (t *JiraProjectStatus) DBUpdate(ctx context.Context, db *sql.DB) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `jira_project_status` SET `checksum`=?,`status_id`=?,`project_id`=?,`name`=?,`description`=?,`icon_url`=?,`category`=?,`category_key`=?,`category_color`=?,`category_id`=?,`issue_type_id`=?,`customer_id`=? WHERE `id`=?"
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.StatusID),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.Description),
		orm.ToSQLString(t.IconURL),
		orm.ToSQLString(t.Category),
		orm.ToSQLString(t.CategoryKey),
		orm.ToSQLString(t.CategoryColor),
		orm.ToSQLString(t.CategoryID),
		orm.ToSQLString(t.IssueTypeID),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.ID),
	)
}

// DBUpdateTx will update the JiraProjectStatus record in the database using the provided transaction
func (t *JiraProjectStatus) DBUpdateTx(ctx context.Context, tx *sql.Tx) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `jira_project_status` SET `checksum`=?,`status_id`=?,`project_id`=?,`name`=?,`description`=?,`icon_url`=?,`category`=?,`category_key`=?,`category_color`=?,`category_id`=?,`issue_type_id`=?,`customer_id`=? WHERE `id`=?"
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.StatusID),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.Description),
		orm.ToSQLString(t.IconURL),
		orm.ToSQLString(t.Category),
		orm.ToSQLString(t.CategoryKey),
		orm.ToSQLString(t.CategoryColor),
		orm.ToSQLString(t.CategoryID),
		orm.ToSQLString(t.IssueTypeID),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLString(t.ID),
	)
}

// DBUpsert will upsert the JiraProjectStatus record in the database
func (t *JiraProjectStatus) DBUpsert(ctx context.Context, db *sql.DB, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `jira_project_status` (`jira_project_status`.`id`,`jira_project_status`.`checksum`,`jira_project_status`.`status_id`,`jira_project_status`.`project_id`,`jira_project_status`.`name`,`jira_project_status`.`description`,`jira_project_status`.`icon_url`,`jira_project_status`.`category`,`jira_project_status`.`category_key`,`jira_project_status`.`category_color`,`jira_project_status`.`category_id`,`jira_project_status`.`issue_type_id`,`jira_project_status`.`customer_id`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `jira_project_status` (`jira_project_status`.`id`,`jira_project_status`.`checksum`,`jira_project_status`.`status_id`,`jira_project_status`.`project_id`,`jira_project_status`.`name`,`jira_project_status`.`description`,`jira_project_status`.`icon_url`,`jira_project_status`.`category`,`jira_project_status`.`category_key`,`jira_project_status`.`category_color`,`jira_project_status`.`category_id`,`jira_project_status`.`issue_type_id`,`jira_project_status`.`customer_id`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`status_id`=VALUES(`status_id`),`project_id`=VALUES(`project_id`),`name`=VALUES(`name`),`description`=VALUES(`description`),`icon_url`=VALUES(`icon_url`),`category`=VALUES(`category`),`category_key`=VALUES(`category_key`),`category_color`=VALUES(`category_color`),`category_id`=VALUES(`category_id`),`issue_type_id`=VALUES(`issue_type_id`),`customer_id`=VALUES(`customer_id`)"
	}
	r, err := db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.StatusID),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.Description),
		orm.ToSQLString(t.IconURL),
		orm.ToSQLString(t.Category),
		orm.ToSQLString(t.CategoryKey),
		orm.ToSQLString(t.CategoryColor),
		orm.ToSQLString(t.CategoryID),
		orm.ToSQLString(t.IssueTypeID),
		orm.ToSQLString(t.CustomerID),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBUpsertTx will upsert the JiraProjectStatus record in the database using the provided transaction
func (t *JiraProjectStatus) DBUpsertTx(ctx context.Context, tx *sql.Tx, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `jira_project_status` (`jira_project_status`.`id`,`jira_project_status`.`checksum`,`jira_project_status`.`status_id`,`jira_project_status`.`project_id`,`jira_project_status`.`name`,`jira_project_status`.`description`,`jira_project_status`.`icon_url`,`jira_project_status`.`category`,`jira_project_status`.`category_key`,`jira_project_status`.`category_color`,`jira_project_status`.`category_id`,`jira_project_status`.`issue_type_id`,`jira_project_status`.`customer_id`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `jira_project_status` (`jira_project_status`.`id`,`jira_project_status`.`checksum`,`jira_project_status`.`status_id`,`jira_project_status`.`project_id`,`jira_project_status`.`name`,`jira_project_status`.`description`,`jira_project_status`.`icon_url`,`jira_project_status`.`category`,`jira_project_status`.`category_key`,`jira_project_status`.`category_color`,`jira_project_status`.`category_id`,`jira_project_status`.`issue_type_id`,`jira_project_status`.`customer_id`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`status_id`=VALUES(`status_id`),`project_id`=VALUES(`project_id`),`name`=VALUES(`name`),`description`=VALUES(`description`),`icon_url`=VALUES(`icon_url`),`category`=VALUES(`category`),`category_key`=VALUES(`category_key`),`category_color`=VALUES(`category_color`),`category_id`=VALUES(`category_id`),`issue_type_id`=VALUES(`issue_type_id`),`customer_id`=VALUES(`customer_id`)"
	}
	r, err := tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.StatusID),
		orm.ToSQLString(t.ProjectID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.Description),
		orm.ToSQLString(t.IconURL),
		orm.ToSQLString(t.Category),
		orm.ToSQLString(t.CategoryKey),
		orm.ToSQLString(t.CategoryColor),
		orm.ToSQLString(t.CategoryID),
		orm.ToSQLString(t.IssueTypeID),
		orm.ToSQLString(t.CustomerID),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBFindOne will find a JiraProjectStatus record in the database with the primary key
func (t *JiraProjectStatus) DBFindOne(ctx context.Context, db *sql.DB, value string) (bool, error) {
	q := "SELECT `jira_project_status`.`id`,`jira_project_status`.`checksum`,`jira_project_status`.`status_id`,`jira_project_status`.`project_id`,`jira_project_status`.`name`,`jira_project_status`.`description`,`jira_project_status`.`icon_url`,`jira_project_status`.`category`,`jira_project_status`.`category_key`,`jira_project_status`.`category_color`,`jira_project_status`.`category_id`,`jira_project_status`.`issue_type_id`,`jira_project_status`.`customer_id` FROM `jira_project_status` WHERE `id` = ? LIMIT 1"
	row := db.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _StatusID sql.NullString
	var _ProjectID sql.NullString
	var _Name sql.NullString
	var _Description sql.NullString
	var _IconURL sql.NullString
	var _Category sql.NullString
	var _CategoryKey sql.NullString
	var _CategoryColor sql.NullString
	var _CategoryID sql.NullString
	var _IssueTypeID sql.NullString
	var _CustomerID sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_StatusID,
		&_ProjectID,
		&_Name,
		&_Description,
		&_IconURL,
		&_Category,
		&_CategoryKey,
		&_CategoryColor,
		&_CategoryID,
		&_IssueTypeID,
		&_CustomerID,
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
	if _StatusID.Valid {
		t.SetStatusID(_StatusID.String)
	}
	if _ProjectID.Valid {
		t.SetProjectID(_ProjectID.String)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _Description.Valid {
		t.SetDescription(_Description.String)
	}
	if _IconURL.Valid {
		t.SetIconURL(_IconURL.String)
	}
	if _Category.Valid {
		t.SetCategory(_Category.String)
	}
	if _CategoryKey.Valid {
		t.SetCategoryKey(_CategoryKey.String)
	}
	if _CategoryColor.Valid {
		t.SetCategoryColor(_CategoryColor.String)
	}
	if _CategoryID.Valid {
		t.SetCategoryID(_CategoryID.String)
	}
	if _IssueTypeID.Valid {
		t.SetIssueTypeID(_IssueTypeID.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	return true, nil
}

// DBFindOneTx will find a JiraProjectStatus record in the database with the primary key using the provided transaction
func (t *JiraProjectStatus) DBFindOneTx(ctx context.Context, tx *sql.Tx, value string) (bool, error) {
	q := "SELECT `jira_project_status`.`id`,`jira_project_status`.`checksum`,`jira_project_status`.`status_id`,`jira_project_status`.`project_id`,`jira_project_status`.`name`,`jira_project_status`.`description`,`jira_project_status`.`icon_url`,`jira_project_status`.`category`,`jira_project_status`.`category_key`,`jira_project_status`.`category_color`,`jira_project_status`.`category_id`,`jira_project_status`.`issue_type_id`,`jira_project_status`.`customer_id` FROM `jira_project_status` WHERE `id` = ? LIMIT 1"
	row := tx.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _StatusID sql.NullString
	var _ProjectID sql.NullString
	var _Name sql.NullString
	var _Description sql.NullString
	var _IconURL sql.NullString
	var _Category sql.NullString
	var _CategoryKey sql.NullString
	var _CategoryColor sql.NullString
	var _CategoryID sql.NullString
	var _IssueTypeID sql.NullString
	var _CustomerID sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_StatusID,
		&_ProjectID,
		&_Name,
		&_Description,
		&_IconURL,
		&_Category,
		&_CategoryKey,
		&_CategoryColor,
		&_CategoryID,
		&_IssueTypeID,
		&_CustomerID,
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
	if _StatusID.Valid {
		t.SetStatusID(_StatusID.String)
	}
	if _ProjectID.Valid {
		t.SetProjectID(_ProjectID.String)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _Description.Valid {
		t.SetDescription(_Description.String)
	}
	if _IconURL.Valid {
		t.SetIconURL(_IconURL.String)
	}
	if _Category.Valid {
		t.SetCategory(_Category.String)
	}
	if _CategoryKey.Valid {
		t.SetCategoryKey(_CategoryKey.String)
	}
	if _CategoryColor.Valid {
		t.SetCategoryColor(_CategoryColor.String)
	}
	if _CategoryID.Valid {
		t.SetCategoryID(_CategoryID.String)
	}
	if _IssueTypeID.Valid {
		t.SetIssueTypeID(_IssueTypeID.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	return true, nil
}

// FindJiraProjectStatuses will find a JiraProjectStatus record in the database with the provided parameters
func FindJiraProjectStatuses(ctx context.Context, db *sql.DB, _params ...interface{}) ([]*JiraProjectStatus, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("status_id"),
		orm.Column("project_id"),
		orm.Column("name"),
		orm.Column("description"),
		orm.Column("icon_url"),
		orm.Column("category"),
		orm.Column("category_key"),
		orm.Column("category_color"),
		orm.Column("category_id"),
		orm.Column("issue_type_id"),
		orm.Column("customer_id"),
		orm.Table(JiraProjectStatusTableName),
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
	results := make([]*JiraProjectStatus, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _StatusID sql.NullString
		var _ProjectID sql.NullString
		var _Name sql.NullString
		var _Description sql.NullString
		var _IconURL sql.NullString
		var _Category sql.NullString
		var _CategoryKey sql.NullString
		var _CategoryColor sql.NullString
		var _CategoryID sql.NullString
		var _IssueTypeID sql.NullString
		var _CustomerID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_StatusID,
			&_ProjectID,
			&_Name,
			&_Description,
			&_IconURL,
			&_Category,
			&_CategoryKey,
			&_CategoryColor,
			&_CategoryID,
			&_IssueTypeID,
			&_CustomerID,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraProjectStatus{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _StatusID.Valid {
			t.SetStatusID(_StatusID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _Description.Valid {
			t.SetDescription(_Description.String)
		}
		if _IconURL.Valid {
			t.SetIconURL(_IconURL.String)
		}
		if _Category.Valid {
			t.SetCategory(_Category.String)
		}
		if _CategoryKey.Valid {
			t.SetCategoryKey(_CategoryKey.String)
		}
		if _CategoryColor.Valid {
			t.SetCategoryColor(_CategoryColor.String)
		}
		if _CategoryID.Valid {
			t.SetCategoryID(_CategoryID.String)
		}
		if _IssueTypeID.Valid {
			t.SetIssueTypeID(_IssueTypeID.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindJiraProjectStatusesTx will find a JiraProjectStatus record in the database with the provided parameters using the provided transaction
func FindJiraProjectStatusesTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) ([]*JiraProjectStatus, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("status_id"),
		orm.Column("project_id"),
		orm.Column("name"),
		orm.Column("description"),
		orm.Column("icon_url"),
		orm.Column("category"),
		orm.Column("category_key"),
		orm.Column("category_color"),
		orm.Column("category_id"),
		orm.Column("issue_type_id"),
		orm.Column("customer_id"),
		orm.Table(JiraProjectStatusTableName),
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
	results := make([]*JiraProjectStatus, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _StatusID sql.NullString
		var _ProjectID sql.NullString
		var _Name sql.NullString
		var _Description sql.NullString
		var _IconURL sql.NullString
		var _Category sql.NullString
		var _CategoryKey sql.NullString
		var _CategoryColor sql.NullString
		var _CategoryID sql.NullString
		var _IssueTypeID sql.NullString
		var _CustomerID sql.NullString
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_StatusID,
			&_ProjectID,
			&_Name,
			&_Description,
			&_IconURL,
			&_Category,
			&_CategoryKey,
			&_CategoryColor,
			&_CategoryID,
			&_IssueTypeID,
			&_CustomerID,
		)
		if err != nil {
			return nil, err
		}
		t := &JiraProjectStatus{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _StatusID.Valid {
			t.SetStatusID(_StatusID.String)
		}
		if _ProjectID.Valid {
			t.SetProjectID(_ProjectID.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _Description.Valid {
			t.SetDescription(_Description.String)
		}
		if _IconURL.Valid {
			t.SetIconURL(_IconURL.String)
		}
		if _Category.Valid {
			t.SetCategory(_Category.String)
		}
		if _CategoryKey.Valid {
			t.SetCategoryKey(_CategoryKey.String)
		}
		if _CategoryColor.Valid {
			t.SetCategoryColor(_CategoryColor.String)
		}
		if _CategoryID.Valid {
			t.SetCategoryID(_CategoryID.String)
		}
		if _IssueTypeID.Valid {
			t.SetIssueTypeID(_IssueTypeID.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		results = append(results, t)
	}
	return results, nil
}

// DBFind will find a JiraProjectStatus record in the database with the provided parameters
func (t *JiraProjectStatus) DBFind(ctx context.Context, db *sql.DB, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("status_id"),
		orm.Column("project_id"),
		orm.Column("name"),
		orm.Column("description"),
		orm.Column("icon_url"),
		orm.Column("category"),
		orm.Column("category_key"),
		orm.Column("category_color"),
		orm.Column("category_id"),
		orm.Column("issue_type_id"),
		orm.Column("customer_id"),
		orm.Table(JiraProjectStatusTableName),
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
	var _StatusID sql.NullString
	var _ProjectID sql.NullString
	var _Name sql.NullString
	var _Description sql.NullString
	var _IconURL sql.NullString
	var _Category sql.NullString
	var _CategoryKey sql.NullString
	var _CategoryColor sql.NullString
	var _CategoryID sql.NullString
	var _IssueTypeID sql.NullString
	var _CustomerID sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_StatusID,
		&_ProjectID,
		&_Name,
		&_Description,
		&_IconURL,
		&_Category,
		&_CategoryKey,
		&_CategoryColor,
		&_CategoryID,
		&_IssueTypeID,
		&_CustomerID,
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
	if _StatusID.Valid {
		t.SetStatusID(_StatusID.String)
	}
	if _ProjectID.Valid {
		t.SetProjectID(_ProjectID.String)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _Description.Valid {
		t.SetDescription(_Description.String)
	}
	if _IconURL.Valid {
		t.SetIconURL(_IconURL.String)
	}
	if _Category.Valid {
		t.SetCategory(_Category.String)
	}
	if _CategoryKey.Valid {
		t.SetCategoryKey(_CategoryKey.String)
	}
	if _CategoryColor.Valid {
		t.SetCategoryColor(_CategoryColor.String)
	}
	if _CategoryID.Valid {
		t.SetCategoryID(_CategoryID.String)
	}
	if _IssueTypeID.Valid {
		t.SetIssueTypeID(_IssueTypeID.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	return true, nil
}

// DBFindTx will find a JiraProjectStatus record in the database with the provided parameters using the provided transaction
func (t *JiraProjectStatus) DBFindTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("status_id"),
		orm.Column("project_id"),
		orm.Column("name"),
		orm.Column("description"),
		orm.Column("icon_url"),
		orm.Column("category"),
		orm.Column("category_key"),
		orm.Column("category_color"),
		orm.Column("category_id"),
		orm.Column("issue_type_id"),
		orm.Column("customer_id"),
		orm.Table(JiraProjectStatusTableName),
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
	var _StatusID sql.NullString
	var _ProjectID sql.NullString
	var _Name sql.NullString
	var _Description sql.NullString
	var _IconURL sql.NullString
	var _Category sql.NullString
	var _CategoryKey sql.NullString
	var _CategoryColor sql.NullString
	var _CategoryID sql.NullString
	var _IssueTypeID sql.NullString
	var _CustomerID sql.NullString
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_StatusID,
		&_ProjectID,
		&_Name,
		&_Description,
		&_IconURL,
		&_Category,
		&_CategoryKey,
		&_CategoryColor,
		&_CategoryID,
		&_IssueTypeID,
		&_CustomerID,
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
	if _StatusID.Valid {
		t.SetStatusID(_StatusID.String)
	}
	if _ProjectID.Valid {
		t.SetProjectID(_ProjectID.String)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _Description.Valid {
		t.SetDescription(_Description.String)
	}
	if _IconURL.Valid {
		t.SetIconURL(_IconURL.String)
	}
	if _Category.Valid {
		t.SetCategory(_Category.String)
	}
	if _CategoryKey.Valid {
		t.SetCategoryKey(_CategoryKey.String)
	}
	if _CategoryColor.Valid {
		t.SetCategoryColor(_CategoryColor.String)
	}
	if _CategoryID.Valid {
		t.SetCategoryID(_CategoryID.String)
	}
	if _IssueTypeID.Valid {
		t.SetIssueTypeID(_IssueTypeID.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	return true, nil
}

// CountJiraProjectStatuses will find the count of JiraProjectStatus records in the database
func CountJiraProjectStatuses(ctx context.Context, db *sql.DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(JiraProjectStatusTableName),
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

// CountJiraProjectStatusesTx will find the count of JiraProjectStatus records in the database using the provided transaction
func CountJiraProjectStatusesTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(JiraProjectStatusTableName),
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

// DBCount will find the count of JiraProjectStatus records in the database
func (t *JiraProjectStatus) DBCount(ctx context.Context, db *sql.DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(JiraProjectStatusTableName),
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

// DBCountTx will find the count of JiraProjectStatus records in the database using the provided transaction
func (t *JiraProjectStatus) DBCountTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(JiraProjectStatusTableName),
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

// DBExists will return true if the JiraProjectStatus record exists in the database
func (t *JiraProjectStatus) DBExists(ctx context.Context, db *sql.DB) (bool, error) {
	q := "SELECT `id` FROM `jira_project_status` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := db.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// DBExistsTx will return true if the JiraProjectStatus record exists in the database using the provided transaction
func (t *JiraProjectStatus) DBExistsTx(ctx context.Context, tx *sql.Tx) (bool, error) {
	q := "SELECT `id` FROM `jira_project_status` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := tx.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// PrimaryKeyColumn returns the column name for the primary key
func (t *JiraProjectStatus) PrimaryKeyColumn() string {
	return JiraProjectStatusColumnID
}

// PrimaryKeyColumnType returns the primary key column Go type as a string
func (t *JiraProjectStatus) PrimaryKeyColumnType() string {
	return "string"
}

// PrimaryKey returns the primary key column value
func (t *JiraProjectStatus) PrimaryKey() interface{} {
	return t.ID
}
