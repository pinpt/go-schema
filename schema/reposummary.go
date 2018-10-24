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

var _ Model = (*RepoSummary)(nil)
var _ CSVWriter = (*RepoSummary)(nil)
var _ JSONWriter = (*RepoSummary)(nil)

// RepoSummaryTableName is the name of the table in SQL
const RepoSummaryTableName = "repo_summary"

var RepoSummaryColumns = []string{
	"id",
	"name",
	"description",
	"ref_type",
	"data_group_id",
	"user_ids",
	"commits",
	"additions",
	"deletions",
	"latest_commit_date",
}

// RepoSummary table
type RepoSummary struct {
	Additions        int64   `json:"additions"`
	Commits          int64   `json:"commits"`
	DataGroupID      string  `json:"data_group_id"`
	Deletions        int64   `json:"deletions"`
	Description      *string `json:"description,omitempty"`
	ID               string  `json:"id"`
	LatestCommitDate int64   `json:"latest_commit_date"`
	Name             string  `json:"name"`
	RefType          string  `json:"ref_type"`
	UserIds          string  `json:"user_ids"`
}

// TableName returns the SQL table name for RepoSummary and satifies the Model interface
func (t *RepoSummary) TableName() string {
	return RepoSummaryTableName
}

// ToCSV will serialize the RepoSummary instance to a CSV compatible array of strings
func (t *RepoSummary) ToCSV() []string {
	return []string{
		t.ID,
		t.Name,
		toCSVString(t.Description),
		t.RefType,
		t.DataGroupID,
		t.UserIds,
		toCSVString(t.Commits),
		toCSVString(t.Additions),
		toCSVString(t.Deletions),
		toCSVString(t.LatestCommitDate),
	}
}

// WriteCSV will serialize the RepoSummary instance to the writer as CSV and satisfies the CSVWriter interface
func (t *RepoSummary) WriteCSV(w *csv.Writer) error {
	return w.Write(t.ToCSV())
}

// WriteJSON will serialize the RepoSummary instance to the writer as JSON and satisfies the JSONWriter interface
func (t *RepoSummary) WriteJSON(w io.Writer, indent ...bool) error {
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

// NewRepoSummaryReader creates a JSON reader which can read in RepoSummary objects serialized as JSON either as an array, single object or json new lines
// and writes each RepoSummary to the channel provided
func NewRepoSummaryReader(r io.Reader, ch chan<- RepoSummary) error {
	return orm.Deserialize(r, func(buf json.RawMessage) error {
		dec := json.NewDecoder(bytes.NewBuffer(buf))
		e := RepoSummary{}
		if err := dec.Decode(&e); err != nil {
			return err
		}
		ch <- e
		return nil
	})
}

// NewCSVRepoSummaryReaderDir will read the reader as CSV and emit each record to the channel provided
func NewCSVRepoSummaryReader(r io.Reader, ch chan<- RepoSummary) error {
	cr := csv.NewReader(r)
	for {
		record, err := cr.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		ch <- RepoSummary{
			ID:               record[0],
			Name:             record[1],
			Description:      fromStringPointer(record[2]),
			RefType:          record[3],
			DataGroupID:      record[4],
			UserIds:          record[5],
			Commits:          fromCSVInt64(record[6]),
			Additions:        fromCSVInt64(record[7]),
			Deletions:        fromCSVInt64(record[8]),
			LatestCommitDate: fromCSVInt64(record[9]),
		}
	}
	return nil
}

// NewCSVRepoSummaryReaderFile will read the file as a CSV and emit each record to the channel provided
func NewCSVRepoSummaryReaderFile(fp string, ch chan<- RepoSummary) error {
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
	return NewCSVRepoSummaryReader(fc, ch)
}

// NewCSVRepoSummaryReaderDir will read the repo_summary.csv.gz file as a CSV and emit each record to the channel provided
func NewCSVRepoSummaryReaderDir(dir string, ch chan<- RepoSummary) error {
	return NewCSVRepoSummaryReaderFile(filepath.Join(dir, "repo_summary.csv.gz"), ch)
}

// RepoSummaryCSVDeduper is a function callback which takes the existing value (a) and the new value (b)
// and the return value should be the one to use (or a new one, if applicable). return nil
// to skip processing of this record
type RepoSummaryCSVDeduper func(a RepoSummary, b RepoSummary) *RepoSummary

// RepoSummaryCSVDedupeDisabled is set on whether the CSV writer should de-dupe values by key
var RepoSummaryCSVDedupeDisabled bool

// NewRepoSummaryCSVWriterSize creates a batch writer that will write each RepoSummary into a CSV file
func NewRepoSummaryCSVWriterSize(w io.Writer, size int, dedupers ...RepoSummaryCSVDeduper) (chan RepoSummary, chan bool, error) {
	cw := csv.NewWriter(w)
	ch := make(chan RepoSummary, size)
	done := make(chan bool)
	go func() {
		defer func() { done <- true }()
		dodedupe := !RepoSummaryCSVDedupeDisabled
		var kv map[string]*RepoSummary
		var deduper RepoSummaryCSVDeduper
		if dedupers != nil && len(dedupers) > 0 {
			deduper = dedupers[0]
			dodedupe = true
		}
		if dodedupe {
			kv = make(map[string]*RepoSummary)
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

// RepoSummaryCSVDefaultSize is the default channel buffer size if not provided
var RepoSummaryCSVDefaultSize = 100

// NewRepoSummaryCSVWriter creates a batch writer that will write each RepoSummary into a CSV file
func NewRepoSummaryCSVWriter(w io.Writer, dedupers ...RepoSummaryCSVDeduper) (chan RepoSummary, chan bool, error) {
	return NewRepoSummaryCSVWriterSize(w, RepoSummaryCSVDefaultSize, dedupers...)
}

// NewRepoSummaryCSVWriterDir creates a batch writer that will write each RepoSummary into a CSV file named repo_summary.csv.gz in dir
func NewRepoSummaryCSVWriterDir(dir string, dedupers ...RepoSummaryCSVDeduper) (chan RepoSummary, chan bool, error) {
	return NewRepoSummaryCSVWriterFile(filepath.Join(dir, "repo_summary.csv.gz"), dedupers...)
}

// NewRepoSummaryCSVWriterFile creates a batch writer that will write each RepoSummary into a CSV file
func NewRepoSummaryCSVWriterFile(fn string, dedupers ...RepoSummaryCSVDeduper) (chan RepoSummary, chan bool, error) {
	f, err := os.Create(fn)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening CSV file %s. %v", fn, err)
	}
	var fc io.WriteCloser = f
	if filepath.Ext(fn) == ".gz" {
		w, _ := gzip.NewWriterLevel(f, gzip.BestCompression)
		fc = w
	}
	ch, done, err := NewRepoSummaryCSVWriter(fc, dedupers...)
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

type RepoSummaryDBAction func(ctx context.Context, db DB, record RepoSummary) error

// NewRepoSummaryDBWriterSize creates a DB writer that will write each issue into the DB
func NewRepoSummaryDBWriterSize(ctx context.Context, db DB, errors chan<- error, size int, actions ...RepoSummaryDBAction) (chan RepoSummary, chan bool, error) {
	ch := make(chan RepoSummary, size)
	done := make(chan bool)
	var action RepoSummaryDBAction
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

// NewRepoSummaryDBWriter creates a DB writer that will write each issue into the DB
func NewRepoSummaryDBWriter(ctx context.Context, db DB, errors chan<- error, actions ...RepoSummaryDBAction) (chan RepoSummary, chan bool, error) {
	return NewRepoSummaryDBWriterSize(ctx, db, errors, 100, actions...)
}

// RepoSummaryColumnID is the ID SQL column name for the RepoSummary table
const RepoSummaryColumnID = "id"

// RepoSummaryEscapedColumnID is the escaped ID SQL column name for the RepoSummary table
const RepoSummaryEscapedColumnID = "`id`"

// RepoSummaryColumnName is the Name SQL column name for the RepoSummary table
const RepoSummaryColumnName = "name"

// RepoSummaryEscapedColumnName is the escaped Name SQL column name for the RepoSummary table
const RepoSummaryEscapedColumnName = "`name`"

// RepoSummaryColumnDescription is the Description SQL column name for the RepoSummary table
const RepoSummaryColumnDescription = "description"

// RepoSummaryEscapedColumnDescription is the escaped Description SQL column name for the RepoSummary table
const RepoSummaryEscapedColumnDescription = "`description`"

// RepoSummaryColumnRefType is the RefType SQL column name for the RepoSummary table
const RepoSummaryColumnRefType = "ref_type"

// RepoSummaryEscapedColumnRefType is the escaped RefType SQL column name for the RepoSummary table
const RepoSummaryEscapedColumnRefType = "`ref_type`"

// RepoSummaryColumnDataGroupID is the DataGroupID SQL column name for the RepoSummary table
const RepoSummaryColumnDataGroupID = "data_group_id"

// RepoSummaryEscapedColumnDataGroupID is the escaped DataGroupID SQL column name for the RepoSummary table
const RepoSummaryEscapedColumnDataGroupID = "`data_group_id`"

// RepoSummaryColumnUserIds is the UserIds SQL column name for the RepoSummary table
const RepoSummaryColumnUserIds = "user_ids"

// RepoSummaryEscapedColumnUserIds is the escaped UserIds SQL column name for the RepoSummary table
const RepoSummaryEscapedColumnUserIds = "`user_ids`"

// RepoSummaryColumnCommits is the Commits SQL column name for the RepoSummary table
const RepoSummaryColumnCommits = "commits"

// RepoSummaryEscapedColumnCommits is the escaped Commits SQL column name for the RepoSummary table
const RepoSummaryEscapedColumnCommits = "`commits`"

// RepoSummaryColumnAdditions is the Additions SQL column name for the RepoSummary table
const RepoSummaryColumnAdditions = "additions"

// RepoSummaryEscapedColumnAdditions is the escaped Additions SQL column name for the RepoSummary table
const RepoSummaryEscapedColumnAdditions = "`additions`"

// RepoSummaryColumnDeletions is the Deletions SQL column name for the RepoSummary table
const RepoSummaryColumnDeletions = "deletions"

// RepoSummaryEscapedColumnDeletions is the escaped Deletions SQL column name for the RepoSummary table
const RepoSummaryEscapedColumnDeletions = "`deletions`"

// RepoSummaryColumnLatestCommitDate is the LatestCommitDate SQL column name for the RepoSummary table
const RepoSummaryColumnLatestCommitDate = "latest_commit_date"

// RepoSummaryEscapedColumnLatestCommitDate is the escaped LatestCommitDate SQL column name for the RepoSummary table
const RepoSummaryEscapedColumnLatestCommitDate = "`latest_commit_date`"

// GetID will return the RepoSummary ID value
func (t *RepoSummary) GetID() string {
	return t.ID
}

// SetID will set the RepoSummary ID value
func (t *RepoSummary) SetID(v string) {
	t.ID = v
}

// FindRepoSummaryByID will find a RepoSummary by ID
func FindRepoSummaryByID(ctx context.Context, db DB, value string) (*RepoSummary, error) {
	q := "SELECT `repo_summary`.`id`,`repo_summary`.`name`,`repo_summary`.`description`,`repo_summary`.`ref_type`,`repo_summary`.`data_group_id`,`repo_summary`.`user_ids`,`repo_summary`.`commits`,`repo_summary`.`additions`,`repo_summary`.`deletions`,`repo_summary`.`latest_commit_date` FROM `repo_summary` WHERE `id` = ?"
	var _ID sql.NullString
	var _Name sql.NullString
	var _Description sql.NullString
	var _RefType sql.NullString
	var _DataGroupID sql.NullString
	var _UserIds sql.NullString
	var _Commits sql.NullInt64
	var _Additions sql.NullInt64
	var _Deletions sql.NullInt64
	var _LatestCommitDate sql.NullInt64
	err := db.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Name,
		&_Description,
		&_RefType,
		&_DataGroupID,
		&_UserIds,
		&_Commits,
		&_Additions,
		&_Deletions,
		&_LatestCommitDate,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &RepoSummary{}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _Description.Valid {
		t.SetDescription(_Description.String)
	}
	if _RefType.Valid {
		t.SetRefType(_RefType.String)
	}
	if _DataGroupID.Valid {
		t.SetDataGroupID(_DataGroupID.String)
	}
	if _UserIds.Valid {
		t.SetUserIds(_UserIds.String)
	}
	if _Commits.Valid {
		t.SetCommits(_Commits.Int64)
	}
	if _Additions.Valid {
		t.SetAdditions(_Additions.Int64)
	}
	if _Deletions.Valid {
		t.SetDeletions(_Deletions.Int64)
	}
	if _LatestCommitDate.Valid {
		t.SetLatestCommitDate(_LatestCommitDate.Int64)
	}
	return t, nil
}

// FindRepoSummaryByIDTx will find a RepoSummary by ID using the provided transaction
func FindRepoSummaryByIDTx(ctx context.Context, tx Tx, value string) (*RepoSummary, error) {
	q := "SELECT `repo_summary`.`id`,`repo_summary`.`name`,`repo_summary`.`description`,`repo_summary`.`ref_type`,`repo_summary`.`data_group_id`,`repo_summary`.`user_ids`,`repo_summary`.`commits`,`repo_summary`.`additions`,`repo_summary`.`deletions`,`repo_summary`.`latest_commit_date` FROM `repo_summary` WHERE `id` = ?"
	var _ID sql.NullString
	var _Name sql.NullString
	var _Description sql.NullString
	var _RefType sql.NullString
	var _DataGroupID sql.NullString
	var _UserIds sql.NullString
	var _Commits sql.NullInt64
	var _Additions sql.NullInt64
	var _Deletions sql.NullInt64
	var _LatestCommitDate sql.NullInt64
	err := tx.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Name,
		&_Description,
		&_RefType,
		&_DataGroupID,
		&_UserIds,
		&_Commits,
		&_Additions,
		&_Deletions,
		&_LatestCommitDate,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &RepoSummary{}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _Description.Valid {
		t.SetDescription(_Description.String)
	}
	if _RefType.Valid {
		t.SetRefType(_RefType.String)
	}
	if _DataGroupID.Valid {
		t.SetDataGroupID(_DataGroupID.String)
	}
	if _UserIds.Valid {
		t.SetUserIds(_UserIds.String)
	}
	if _Commits.Valid {
		t.SetCommits(_Commits.Int64)
	}
	if _Additions.Valid {
		t.SetAdditions(_Additions.Int64)
	}
	if _Deletions.Valid {
		t.SetDeletions(_Deletions.Int64)
	}
	if _LatestCommitDate.Valid {
		t.SetLatestCommitDate(_LatestCommitDate.Int64)
	}
	return t, nil
}

// GetName will return the RepoSummary Name value
func (t *RepoSummary) GetName() string {
	return t.Name
}

// SetName will set the RepoSummary Name value
func (t *RepoSummary) SetName(v string) {
	t.Name = v
}

// FindRepoSummariesByName will find all RepoSummarys by the Name value
func FindRepoSummariesByName(ctx context.Context, db DB, value string) ([]*RepoSummary, error) {
	q := "SELECT `repo_summary`.`id`,`repo_summary`.`name`,`repo_summary`.`description`,`repo_summary`.`ref_type`,`repo_summary`.`data_group_id`,`repo_summary`.`user_ids`,`repo_summary`.`commits`,`repo_summary`.`additions`,`repo_summary`.`deletions`,`repo_summary`.`latest_commit_date` FROM `repo_summary` WHERE `name` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*RepoSummary, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Name sql.NullString
		var _Description sql.NullString
		var _RefType sql.NullString
		var _DataGroupID sql.NullString
		var _UserIds sql.NullString
		var _Commits sql.NullInt64
		var _Additions sql.NullInt64
		var _Deletions sql.NullInt64
		var _LatestCommitDate sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Name,
			&_Description,
			&_RefType,
			&_DataGroupID,
			&_UserIds,
			&_Commits,
			&_Additions,
			&_Deletions,
			&_LatestCommitDate,
		)
		if err != nil {
			return nil, err
		}
		t := &RepoSummary{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _Description.Valid {
			t.SetDescription(_Description.String)
		}
		if _RefType.Valid {
			t.SetRefType(_RefType.String)
		}
		if _DataGroupID.Valid {
			t.SetDataGroupID(_DataGroupID.String)
		}
		if _UserIds.Valid {
			t.SetUserIds(_UserIds.String)
		}
		if _Commits.Valid {
			t.SetCommits(_Commits.Int64)
		}
		if _Additions.Valid {
			t.SetAdditions(_Additions.Int64)
		}
		if _Deletions.Valid {
			t.SetDeletions(_Deletions.Int64)
		}
		if _LatestCommitDate.Valid {
			t.SetLatestCommitDate(_LatestCommitDate.Int64)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindRepoSummariesByNameTx will find all RepoSummarys by the Name value using the provided transaction
func FindRepoSummariesByNameTx(ctx context.Context, tx Tx, value string) ([]*RepoSummary, error) {
	q := "SELECT `repo_summary`.`id`,`repo_summary`.`name`,`repo_summary`.`description`,`repo_summary`.`ref_type`,`repo_summary`.`data_group_id`,`repo_summary`.`user_ids`,`repo_summary`.`commits`,`repo_summary`.`additions`,`repo_summary`.`deletions`,`repo_summary`.`latest_commit_date` FROM `repo_summary` WHERE `name` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*RepoSummary, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Name sql.NullString
		var _Description sql.NullString
		var _RefType sql.NullString
		var _DataGroupID sql.NullString
		var _UserIds sql.NullString
		var _Commits sql.NullInt64
		var _Additions sql.NullInt64
		var _Deletions sql.NullInt64
		var _LatestCommitDate sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Name,
			&_Description,
			&_RefType,
			&_DataGroupID,
			&_UserIds,
			&_Commits,
			&_Additions,
			&_Deletions,
			&_LatestCommitDate,
		)
		if err != nil {
			return nil, err
		}
		t := &RepoSummary{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _Description.Valid {
			t.SetDescription(_Description.String)
		}
		if _RefType.Valid {
			t.SetRefType(_RefType.String)
		}
		if _DataGroupID.Valid {
			t.SetDataGroupID(_DataGroupID.String)
		}
		if _UserIds.Valid {
			t.SetUserIds(_UserIds.String)
		}
		if _Commits.Valid {
			t.SetCommits(_Commits.Int64)
		}
		if _Additions.Valid {
			t.SetAdditions(_Additions.Int64)
		}
		if _Deletions.Valid {
			t.SetDeletions(_Deletions.Int64)
		}
		if _LatestCommitDate.Valid {
			t.SetLatestCommitDate(_LatestCommitDate.Int64)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetDescription will return the RepoSummary Description value
func (t *RepoSummary) GetDescription() string {
	if t.Description == nil {
		return ""
	}
	return *t.Description
}

// SetDescription will set the RepoSummary Description value
func (t *RepoSummary) SetDescription(v string) {
	t.Description = &v
}

// GetRefType will return the RepoSummary RefType value
func (t *RepoSummary) GetRefType() string {
	return t.RefType
}

// SetRefType will set the RepoSummary RefType value
func (t *RepoSummary) SetRefType(v string) {
	t.RefType = v
}

// FindRepoSummariesByRefType will find all RepoSummarys by the RefType value
func FindRepoSummariesByRefType(ctx context.Context, db DB, value string) ([]*RepoSummary, error) {
	q := "SELECT `repo_summary`.`id`,`repo_summary`.`name`,`repo_summary`.`description`,`repo_summary`.`ref_type`,`repo_summary`.`data_group_id`,`repo_summary`.`user_ids`,`repo_summary`.`commits`,`repo_summary`.`additions`,`repo_summary`.`deletions`,`repo_summary`.`latest_commit_date` FROM `repo_summary` WHERE `ref_type` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*RepoSummary, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Name sql.NullString
		var _Description sql.NullString
		var _RefType sql.NullString
		var _DataGroupID sql.NullString
		var _UserIds sql.NullString
		var _Commits sql.NullInt64
		var _Additions sql.NullInt64
		var _Deletions sql.NullInt64
		var _LatestCommitDate sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Name,
			&_Description,
			&_RefType,
			&_DataGroupID,
			&_UserIds,
			&_Commits,
			&_Additions,
			&_Deletions,
			&_LatestCommitDate,
		)
		if err != nil {
			return nil, err
		}
		t := &RepoSummary{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _Description.Valid {
			t.SetDescription(_Description.String)
		}
		if _RefType.Valid {
			t.SetRefType(_RefType.String)
		}
		if _DataGroupID.Valid {
			t.SetDataGroupID(_DataGroupID.String)
		}
		if _UserIds.Valid {
			t.SetUserIds(_UserIds.String)
		}
		if _Commits.Valid {
			t.SetCommits(_Commits.Int64)
		}
		if _Additions.Valid {
			t.SetAdditions(_Additions.Int64)
		}
		if _Deletions.Valid {
			t.SetDeletions(_Deletions.Int64)
		}
		if _LatestCommitDate.Valid {
			t.SetLatestCommitDate(_LatestCommitDate.Int64)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindRepoSummariesByRefTypeTx will find all RepoSummarys by the RefType value using the provided transaction
func FindRepoSummariesByRefTypeTx(ctx context.Context, tx Tx, value string) ([]*RepoSummary, error) {
	q := "SELECT `repo_summary`.`id`,`repo_summary`.`name`,`repo_summary`.`description`,`repo_summary`.`ref_type`,`repo_summary`.`data_group_id`,`repo_summary`.`user_ids`,`repo_summary`.`commits`,`repo_summary`.`additions`,`repo_summary`.`deletions`,`repo_summary`.`latest_commit_date` FROM `repo_summary` WHERE `ref_type` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*RepoSummary, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Name sql.NullString
		var _Description sql.NullString
		var _RefType sql.NullString
		var _DataGroupID sql.NullString
		var _UserIds sql.NullString
		var _Commits sql.NullInt64
		var _Additions sql.NullInt64
		var _Deletions sql.NullInt64
		var _LatestCommitDate sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Name,
			&_Description,
			&_RefType,
			&_DataGroupID,
			&_UserIds,
			&_Commits,
			&_Additions,
			&_Deletions,
			&_LatestCommitDate,
		)
		if err != nil {
			return nil, err
		}
		t := &RepoSummary{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _Description.Valid {
			t.SetDescription(_Description.String)
		}
		if _RefType.Valid {
			t.SetRefType(_RefType.String)
		}
		if _DataGroupID.Valid {
			t.SetDataGroupID(_DataGroupID.String)
		}
		if _UserIds.Valid {
			t.SetUserIds(_UserIds.String)
		}
		if _Commits.Valid {
			t.SetCommits(_Commits.Int64)
		}
		if _Additions.Valid {
			t.SetAdditions(_Additions.Int64)
		}
		if _Deletions.Valid {
			t.SetDeletions(_Deletions.Int64)
		}
		if _LatestCommitDate.Valid {
			t.SetLatestCommitDate(_LatestCommitDate.Int64)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetDataGroupID will return the RepoSummary DataGroupID value
func (t *RepoSummary) GetDataGroupID() string {
	return t.DataGroupID
}

// SetDataGroupID will set the RepoSummary DataGroupID value
func (t *RepoSummary) SetDataGroupID(v string) {
	t.DataGroupID = v
}

// FindRepoSummariesByDataGroupID will find all RepoSummarys by the DataGroupID value
func FindRepoSummariesByDataGroupID(ctx context.Context, db DB, value string) ([]*RepoSummary, error) {
	q := "SELECT `repo_summary`.`id`,`repo_summary`.`name`,`repo_summary`.`description`,`repo_summary`.`ref_type`,`repo_summary`.`data_group_id`,`repo_summary`.`user_ids`,`repo_summary`.`commits`,`repo_summary`.`additions`,`repo_summary`.`deletions`,`repo_summary`.`latest_commit_date` FROM `repo_summary` WHERE `data_group_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*RepoSummary, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Name sql.NullString
		var _Description sql.NullString
		var _RefType sql.NullString
		var _DataGroupID sql.NullString
		var _UserIds sql.NullString
		var _Commits sql.NullInt64
		var _Additions sql.NullInt64
		var _Deletions sql.NullInt64
		var _LatestCommitDate sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Name,
			&_Description,
			&_RefType,
			&_DataGroupID,
			&_UserIds,
			&_Commits,
			&_Additions,
			&_Deletions,
			&_LatestCommitDate,
		)
		if err != nil {
			return nil, err
		}
		t := &RepoSummary{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _Description.Valid {
			t.SetDescription(_Description.String)
		}
		if _RefType.Valid {
			t.SetRefType(_RefType.String)
		}
		if _DataGroupID.Valid {
			t.SetDataGroupID(_DataGroupID.String)
		}
		if _UserIds.Valid {
			t.SetUserIds(_UserIds.String)
		}
		if _Commits.Valid {
			t.SetCommits(_Commits.Int64)
		}
		if _Additions.Valid {
			t.SetAdditions(_Additions.Int64)
		}
		if _Deletions.Valid {
			t.SetDeletions(_Deletions.Int64)
		}
		if _LatestCommitDate.Valid {
			t.SetLatestCommitDate(_LatestCommitDate.Int64)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindRepoSummariesByDataGroupIDTx will find all RepoSummarys by the DataGroupID value using the provided transaction
func FindRepoSummariesByDataGroupIDTx(ctx context.Context, tx Tx, value string) ([]*RepoSummary, error) {
	q := "SELECT `repo_summary`.`id`,`repo_summary`.`name`,`repo_summary`.`description`,`repo_summary`.`ref_type`,`repo_summary`.`data_group_id`,`repo_summary`.`user_ids`,`repo_summary`.`commits`,`repo_summary`.`additions`,`repo_summary`.`deletions`,`repo_summary`.`latest_commit_date` FROM `repo_summary` WHERE `data_group_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*RepoSummary, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Name sql.NullString
		var _Description sql.NullString
		var _RefType sql.NullString
		var _DataGroupID sql.NullString
		var _UserIds sql.NullString
		var _Commits sql.NullInt64
		var _Additions sql.NullInt64
		var _Deletions sql.NullInt64
		var _LatestCommitDate sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Name,
			&_Description,
			&_RefType,
			&_DataGroupID,
			&_UserIds,
			&_Commits,
			&_Additions,
			&_Deletions,
			&_LatestCommitDate,
		)
		if err != nil {
			return nil, err
		}
		t := &RepoSummary{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _Description.Valid {
			t.SetDescription(_Description.String)
		}
		if _RefType.Valid {
			t.SetRefType(_RefType.String)
		}
		if _DataGroupID.Valid {
			t.SetDataGroupID(_DataGroupID.String)
		}
		if _UserIds.Valid {
			t.SetUserIds(_UserIds.String)
		}
		if _Commits.Valid {
			t.SetCommits(_Commits.Int64)
		}
		if _Additions.Valid {
			t.SetAdditions(_Additions.Int64)
		}
		if _Deletions.Valid {
			t.SetDeletions(_Deletions.Int64)
		}
		if _LatestCommitDate.Valid {
			t.SetLatestCommitDate(_LatestCommitDate.Int64)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetUserIds will return the RepoSummary UserIds value
func (t *RepoSummary) GetUserIds() string {
	return t.UserIds
}

// SetUserIds will set the RepoSummary UserIds value
func (t *RepoSummary) SetUserIds(v string) {
	t.UserIds = v
}

// GetCommits will return the RepoSummary Commits value
func (t *RepoSummary) GetCommits() int64 {
	return t.Commits
}

// SetCommits will set the RepoSummary Commits value
func (t *RepoSummary) SetCommits(v int64) {
	t.Commits = v
}

// GetAdditions will return the RepoSummary Additions value
func (t *RepoSummary) GetAdditions() int64 {
	return t.Additions
}

// SetAdditions will set the RepoSummary Additions value
func (t *RepoSummary) SetAdditions(v int64) {
	t.Additions = v
}

// GetDeletions will return the RepoSummary Deletions value
func (t *RepoSummary) GetDeletions() int64 {
	return t.Deletions
}

// SetDeletions will set the RepoSummary Deletions value
func (t *RepoSummary) SetDeletions(v int64) {
	t.Deletions = v
}

// GetLatestCommitDate will return the RepoSummary LatestCommitDate value
func (t *RepoSummary) GetLatestCommitDate() int64 {
	return t.LatestCommitDate
}

// SetLatestCommitDate will set the RepoSummary LatestCommitDate value
func (t *RepoSummary) SetLatestCommitDate(v int64) {
	t.LatestCommitDate = v
}

func (t *RepoSummary) toTimestamp(value time.Time) *timestamp.Timestamp {
	ts, _ := ptypes.TimestampProto(value)
	return ts
}

// DBCreateRepoSummaryTable will create the RepoSummary table
func DBCreateRepoSummaryTable(ctx context.Context, db DB) error {
	q := "CREATE TABLE `repo_summary` (`id` VARCHAR(64) NOT NULL PRIMARY KEY,`name`VARCHAR(255) NOT NULL,`description` TEXT,`ref_type` VARCHAR(20) NOT NULL,`data_group_id`VARCHAR(20) NOT NULL,`user_ids` JSON NOT NULL,`commits`BIGINT UNSIGNED NOT NULL,`additions` BIGINT UNSIGNED NOT NULL,`deletions` BIGINT UNSIGNED NOT NULL,`latest_commit_date` BIGINT UNSIGNED NOT NULL,INDEX repo_summary_name_index (`name`),INDEX repo_summary_ref_type_index (`ref_type`),INDEX repo_summary_data_group_id_index (`data_group_id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBCreateRepoSummaryTableTx will create the RepoSummary table using the provided transction
func DBCreateRepoSummaryTableTx(ctx context.Context, tx Tx) error {
	q := "CREATE TABLE `repo_summary` (`id` VARCHAR(64) NOT NULL PRIMARY KEY,`name`VARCHAR(255) NOT NULL,`description` TEXT,`ref_type` VARCHAR(20) NOT NULL,`data_group_id`VARCHAR(20) NOT NULL,`user_ids` JSON NOT NULL,`commits`BIGINT UNSIGNED NOT NULL,`additions` BIGINT UNSIGNED NOT NULL,`deletions` BIGINT UNSIGNED NOT NULL,`latest_commit_date` BIGINT UNSIGNED NOT NULL,INDEX repo_summary_name_index (`name`),INDEX repo_summary_ref_type_index (`ref_type`),INDEX repo_summary_data_group_id_index (`data_group_id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// DBDropRepoSummaryTable will drop the RepoSummary table
func DBDropRepoSummaryTable(ctx context.Context, db DB) error {
	q := "DROP TABLE IF EXISTS `repo_summary`"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBDropRepoSummaryTableTx will drop the RepoSummary table using the provided transaction
func DBDropRepoSummaryTableTx(ctx context.Context, tx Tx) error {
	q := "DROP TABLE IF EXISTS `repo_summary`"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// DBCreate will create a new RepoSummary record in the database
func (t *RepoSummary) DBCreate(ctx context.Context, db DB) (sql.Result, error) {
	q := "INSERT INTO `repo_summary` (`repo_summary`.`id`,`repo_summary`.`name`,`repo_summary`.`description`,`repo_summary`.`ref_type`,`repo_summary`.`data_group_id`,`repo_summary`.`user_ids`,`repo_summary`.`commits`,`repo_summary`.`additions`,`repo_summary`.`deletions`,`repo_summary`.`latest_commit_date`) VALUES (?,?,?,?,?,?,?,?,?,?)"
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.Description),
		orm.ToSQLString(t.RefType),
		orm.ToSQLString(t.DataGroupID),
		orm.ToSQLString(t.UserIds),
		orm.ToSQLInt64(t.Commits),
		orm.ToSQLInt64(t.Additions),
		orm.ToSQLInt64(t.Deletions),
		orm.ToSQLInt64(t.LatestCommitDate),
	)
}

// DBCreateTx will create a new RepoSummary record in the database using the provided transaction
func (t *RepoSummary) DBCreateTx(ctx context.Context, tx Tx) (sql.Result, error) {
	q := "INSERT INTO `repo_summary` (`repo_summary`.`id`,`repo_summary`.`name`,`repo_summary`.`description`,`repo_summary`.`ref_type`,`repo_summary`.`data_group_id`,`repo_summary`.`user_ids`,`repo_summary`.`commits`,`repo_summary`.`additions`,`repo_summary`.`deletions`,`repo_summary`.`latest_commit_date`) VALUES (?,?,?,?,?,?,?,?,?,?)"
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.Description),
		orm.ToSQLString(t.RefType),
		orm.ToSQLString(t.DataGroupID),
		orm.ToSQLString(t.UserIds),
		orm.ToSQLInt64(t.Commits),
		orm.ToSQLInt64(t.Additions),
		orm.ToSQLInt64(t.Deletions),
		orm.ToSQLInt64(t.LatestCommitDate),
	)
}

// DBCreateIgnoreDuplicate will upsert the RepoSummary record in the database
func (t *RepoSummary) DBCreateIgnoreDuplicate(ctx context.Context, db DB) (sql.Result, error) {
	q := "INSERT INTO `repo_summary` (`repo_summary`.`id`,`repo_summary`.`name`,`repo_summary`.`description`,`repo_summary`.`ref_type`,`repo_summary`.`data_group_id`,`repo_summary`.`user_ids`,`repo_summary`.`commits`,`repo_summary`.`additions`,`repo_summary`.`deletions`,`repo_summary`.`latest_commit_date`) VALUES (?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.Description),
		orm.ToSQLString(t.RefType),
		orm.ToSQLString(t.DataGroupID),
		orm.ToSQLString(t.UserIds),
		orm.ToSQLInt64(t.Commits),
		orm.ToSQLInt64(t.Additions),
		orm.ToSQLInt64(t.Deletions),
		orm.ToSQLInt64(t.LatestCommitDate),
	)
}

// DBCreateIgnoreDuplicateTx will upsert the RepoSummary record in the database using the provided transaction
func (t *RepoSummary) DBCreateIgnoreDuplicateTx(ctx context.Context, tx Tx) (sql.Result, error) {
	q := "INSERT INTO `repo_summary` (`repo_summary`.`id`,`repo_summary`.`name`,`repo_summary`.`description`,`repo_summary`.`ref_type`,`repo_summary`.`data_group_id`,`repo_summary`.`user_ids`,`repo_summary`.`commits`,`repo_summary`.`additions`,`repo_summary`.`deletions`,`repo_summary`.`latest_commit_date`) VALUES (?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.Description),
		orm.ToSQLString(t.RefType),
		orm.ToSQLString(t.DataGroupID),
		orm.ToSQLString(t.UserIds),
		orm.ToSQLInt64(t.Commits),
		orm.ToSQLInt64(t.Additions),
		orm.ToSQLInt64(t.Deletions),
		orm.ToSQLInt64(t.LatestCommitDate),
	)
}

// DeleteAllRepoSummaries deletes all RepoSummary records in the database with optional filters
func DeleteAllRepoSummaries(ctx context.Context, db DB, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(RepoSummaryTableName),
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

// DeleteAllRepoSummariesTx deletes all RepoSummary records in the database with optional filters using the provided transaction
func DeleteAllRepoSummariesTx(ctx context.Context, tx Tx, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(RepoSummaryTableName),
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

// DBDelete will delete this RepoSummary record in the database
func (t *RepoSummary) DBDelete(ctx context.Context, db DB) (bool, error) {
	q := "DELETE FROM `repo_summary` WHERE `id` = ?"
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

// DBDeleteTx will delete this RepoSummary record in the database using the provided transaction
func (t *RepoSummary) DBDeleteTx(ctx context.Context, tx Tx) (bool, error) {
	q := "DELETE FROM `repo_summary` WHERE `id` = ?"
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

// DBUpdate will update the RepoSummary record in the database
func (t *RepoSummary) DBUpdate(ctx context.Context, db DB) (sql.Result, error) {
	q := "UPDATE `repo_summary` SET `name`=?,`description`=?,`ref_type`=?,`data_group_id`=?,`user_ids`=?,`commits`=?,`additions`=?,`deletions`=?,`latest_commit_date`=? WHERE `id`=?"
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.Description),
		orm.ToSQLString(t.RefType),
		orm.ToSQLString(t.DataGroupID),
		orm.ToSQLString(t.UserIds),
		orm.ToSQLInt64(t.Commits),
		orm.ToSQLInt64(t.Additions),
		orm.ToSQLInt64(t.Deletions),
		orm.ToSQLInt64(t.LatestCommitDate),
		orm.ToSQLString(t.ID),
	)
}

// DBUpdateTx will update the RepoSummary record in the database using the provided transaction
func (t *RepoSummary) DBUpdateTx(ctx context.Context, tx Tx) (sql.Result, error) {
	q := "UPDATE `repo_summary` SET `name`=?,`description`=?,`ref_type`=?,`data_group_id`=?,`user_ids`=?,`commits`=?,`additions`=?,`deletions`=?,`latest_commit_date`=? WHERE `id`=?"
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.Description),
		orm.ToSQLString(t.RefType),
		orm.ToSQLString(t.DataGroupID),
		orm.ToSQLString(t.UserIds),
		orm.ToSQLInt64(t.Commits),
		orm.ToSQLInt64(t.Additions),
		orm.ToSQLInt64(t.Deletions),
		orm.ToSQLInt64(t.LatestCommitDate),
		orm.ToSQLString(t.ID),
	)
}

// DBUpsert will upsert the RepoSummary record in the database
func (t *RepoSummary) DBUpsert(ctx context.Context, db DB, conditions ...interface{}) (bool, bool, error) {
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `repo_summary` (`repo_summary`.`id`,`repo_summary`.`name`,`repo_summary`.`description`,`repo_summary`.`ref_type`,`repo_summary`.`data_group_id`,`repo_summary`.`user_ids`,`repo_summary`.`commits`,`repo_summary`.`additions`,`repo_summary`.`deletions`,`repo_summary`.`latest_commit_date`) VALUES (?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `repo_summary` (`repo_summary`.`id`,`repo_summary`.`name`,`repo_summary`.`description`,`repo_summary`.`ref_type`,`repo_summary`.`data_group_id`,`repo_summary`.`user_ids`,`repo_summary`.`commits`,`repo_summary`.`additions`,`repo_summary`.`deletions`,`repo_summary`.`latest_commit_date`) VALUES (?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `name`=VALUES(`name`),`description`=VALUES(`description`),`ref_type`=VALUES(`ref_type`),`data_group_id`=VALUES(`data_group_id`),`user_ids`=VALUES(`user_ids`),`commits`=VALUES(`commits`),`additions`=VALUES(`additions`),`deletions`=VALUES(`deletions`),`latest_commit_date`=VALUES(`latest_commit_date`)"
	}
	r, err := db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.Description),
		orm.ToSQLString(t.RefType),
		orm.ToSQLString(t.DataGroupID),
		orm.ToSQLString(t.UserIds),
		orm.ToSQLInt64(t.Commits),
		orm.ToSQLInt64(t.Additions),
		orm.ToSQLInt64(t.Deletions),
		orm.ToSQLInt64(t.LatestCommitDate),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBUpsertTx will upsert the RepoSummary record in the database using the provided transaction
func (t *RepoSummary) DBUpsertTx(ctx context.Context, tx Tx, conditions ...interface{}) (bool, bool, error) {
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `repo_summary` (`repo_summary`.`id`,`repo_summary`.`name`,`repo_summary`.`description`,`repo_summary`.`ref_type`,`repo_summary`.`data_group_id`,`repo_summary`.`user_ids`,`repo_summary`.`commits`,`repo_summary`.`additions`,`repo_summary`.`deletions`,`repo_summary`.`latest_commit_date`) VALUES (?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `repo_summary` (`repo_summary`.`id`,`repo_summary`.`name`,`repo_summary`.`description`,`repo_summary`.`ref_type`,`repo_summary`.`data_group_id`,`repo_summary`.`user_ids`,`repo_summary`.`commits`,`repo_summary`.`additions`,`repo_summary`.`deletions`,`repo_summary`.`latest_commit_date`) VALUES (?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `name`=VALUES(`name`),`description`=VALUES(`description`),`ref_type`=VALUES(`ref_type`),`data_group_id`=VALUES(`data_group_id`),`user_ids`=VALUES(`user_ids`),`commits`=VALUES(`commits`),`additions`=VALUES(`additions`),`deletions`=VALUES(`deletions`),`latest_commit_date`=VALUES(`latest_commit_date`)"
	}
	r, err := tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.Description),
		orm.ToSQLString(t.RefType),
		orm.ToSQLString(t.DataGroupID),
		orm.ToSQLString(t.UserIds),
		orm.ToSQLInt64(t.Commits),
		orm.ToSQLInt64(t.Additions),
		orm.ToSQLInt64(t.Deletions),
		orm.ToSQLInt64(t.LatestCommitDate),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBFindOne will find a RepoSummary record in the database with the primary key
func (t *RepoSummary) DBFindOne(ctx context.Context, db DB, value string) (bool, error) {
	q := "SELECT `repo_summary`.`id`,`repo_summary`.`name`,`repo_summary`.`description`,`repo_summary`.`ref_type`,`repo_summary`.`data_group_id`,`repo_summary`.`user_ids`,`repo_summary`.`commits`,`repo_summary`.`additions`,`repo_summary`.`deletions`,`repo_summary`.`latest_commit_date` FROM `repo_summary` WHERE `id` = ? LIMIT 1"
	row := db.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Name sql.NullString
	var _Description sql.NullString
	var _RefType sql.NullString
	var _DataGroupID sql.NullString
	var _UserIds sql.NullString
	var _Commits sql.NullInt64
	var _Additions sql.NullInt64
	var _Deletions sql.NullInt64
	var _LatestCommitDate sql.NullInt64
	err := row.Scan(
		&_ID,
		&_Name,
		&_Description,
		&_RefType,
		&_DataGroupID,
		&_UserIds,
		&_Commits,
		&_Additions,
		&_Deletions,
		&_LatestCommitDate,
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
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _Description.Valid {
		t.SetDescription(_Description.String)
	}
	if _RefType.Valid {
		t.SetRefType(_RefType.String)
	}
	if _DataGroupID.Valid {
		t.SetDataGroupID(_DataGroupID.String)
	}
	if _UserIds.Valid {
		t.SetUserIds(_UserIds.String)
	}
	if _Commits.Valid {
		t.SetCommits(_Commits.Int64)
	}
	if _Additions.Valid {
		t.SetAdditions(_Additions.Int64)
	}
	if _Deletions.Valid {
		t.SetDeletions(_Deletions.Int64)
	}
	if _LatestCommitDate.Valid {
		t.SetLatestCommitDate(_LatestCommitDate.Int64)
	}
	return true, nil
}

// DBFindOneTx will find a RepoSummary record in the database with the primary key using the provided transaction
func (t *RepoSummary) DBFindOneTx(ctx context.Context, tx Tx, value string) (bool, error) {
	q := "SELECT `repo_summary`.`id`,`repo_summary`.`name`,`repo_summary`.`description`,`repo_summary`.`ref_type`,`repo_summary`.`data_group_id`,`repo_summary`.`user_ids`,`repo_summary`.`commits`,`repo_summary`.`additions`,`repo_summary`.`deletions`,`repo_summary`.`latest_commit_date` FROM `repo_summary` WHERE `id` = ? LIMIT 1"
	row := tx.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Name sql.NullString
	var _Description sql.NullString
	var _RefType sql.NullString
	var _DataGroupID sql.NullString
	var _UserIds sql.NullString
	var _Commits sql.NullInt64
	var _Additions sql.NullInt64
	var _Deletions sql.NullInt64
	var _LatestCommitDate sql.NullInt64
	err := row.Scan(
		&_ID,
		&_Name,
		&_Description,
		&_RefType,
		&_DataGroupID,
		&_UserIds,
		&_Commits,
		&_Additions,
		&_Deletions,
		&_LatestCommitDate,
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
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _Description.Valid {
		t.SetDescription(_Description.String)
	}
	if _RefType.Valid {
		t.SetRefType(_RefType.String)
	}
	if _DataGroupID.Valid {
		t.SetDataGroupID(_DataGroupID.String)
	}
	if _UserIds.Valid {
		t.SetUserIds(_UserIds.String)
	}
	if _Commits.Valid {
		t.SetCommits(_Commits.Int64)
	}
	if _Additions.Valid {
		t.SetAdditions(_Additions.Int64)
	}
	if _Deletions.Valid {
		t.SetDeletions(_Deletions.Int64)
	}
	if _LatestCommitDate.Valid {
		t.SetLatestCommitDate(_LatestCommitDate.Int64)
	}
	return true, nil
}

// FindRepoSummaries will find a RepoSummary record in the database with the provided parameters
func FindRepoSummaries(ctx context.Context, db DB, _params ...interface{}) ([]*RepoSummary, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("name"),
		orm.Column("description"),
		orm.Column("ref_type"),
		orm.Column("data_group_id"),
		orm.Column("user_ids"),
		orm.Column("commits"),
		orm.Column("additions"),
		orm.Column("deletions"),
		orm.Column("latest_commit_date"),
		orm.Table(RepoSummaryTableName),
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
	results := make([]*RepoSummary, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Name sql.NullString
		var _Description sql.NullString
		var _RefType sql.NullString
		var _DataGroupID sql.NullString
		var _UserIds sql.NullString
		var _Commits sql.NullInt64
		var _Additions sql.NullInt64
		var _Deletions sql.NullInt64
		var _LatestCommitDate sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Name,
			&_Description,
			&_RefType,
			&_DataGroupID,
			&_UserIds,
			&_Commits,
			&_Additions,
			&_Deletions,
			&_LatestCommitDate,
		)
		if err != nil {
			return nil, err
		}
		t := &RepoSummary{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _Description.Valid {
			t.SetDescription(_Description.String)
		}
		if _RefType.Valid {
			t.SetRefType(_RefType.String)
		}
		if _DataGroupID.Valid {
			t.SetDataGroupID(_DataGroupID.String)
		}
		if _UserIds.Valid {
			t.SetUserIds(_UserIds.String)
		}
		if _Commits.Valid {
			t.SetCommits(_Commits.Int64)
		}
		if _Additions.Valid {
			t.SetAdditions(_Additions.Int64)
		}
		if _Deletions.Valid {
			t.SetDeletions(_Deletions.Int64)
		}
		if _LatestCommitDate.Valid {
			t.SetLatestCommitDate(_LatestCommitDate.Int64)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindRepoSummariesTx will find a RepoSummary record in the database with the provided parameters using the provided transaction
func FindRepoSummariesTx(ctx context.Context, tx Tx, _params ...interface{}) ([]*RepoSummary, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("name"),
		orm.Column("description"),
		orm.Column("ref_type"),
		orm.Column("data_group_id"),
		orm.Column("user_ids"),
		orm.Column("commits"),
		orm.Column("additions"),
		orm.Column("deletions"),
		orm.Column("latest_commit_date"),
		orm.Table(RepoSummaryTableName),
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
	results := make([]*RepoSummary, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Name sql.NullString
		var _Description sql.NullString
		var _RefType sql.NullString
		var _DataGroupID sql.NullString
		var _UserIds sql.NullString
		var _Commits sql.NullInt64
		var _Additions sql.NullInt64
		var _Deletions sql.NullInt64
		var _LatestCommitDate sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Name,
			&_Description,
			&_RefType,
			&_DataGroupID,
			&_UserIds,
			&_Commits,
			&_Additions,
			&_Deletions,
			&_LatestCommitDate,
		)
		if err != nil {
			return nil, err
		}
		t := &RepoSummary{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _Description.Valid {
			t.SetDescription(_Description.String)
		}
		if _RefType.Valid {
			t.SetRefType(_RefType.String)
		}
		if _DataGroupID.Valid {
			t.SetDataGroupID(_DataGroupID.String)
		}
		if _UserIds.Valid {
			t.SetUserIds(_UserIds.String)
		}
		if _Commits.Valid {
			t.SetCommits(_Commits.Int64)
		}
		if _Additions.Valid {
			t.SetAdditions(_Additions.Int64)
		}
		if _Deletions.Valid {
			t.SetDeletions(_Deletions.Int64)
		}
		if _LatestCommitDate.Valid {
			t.SetLatestCommitDate(_LatestCommitDate.Int64)
		}
		results = append(results, t)
	}
	return results, nil
}

// DBFind will find a RepoSummary record in the database with the provided parameters
func (t *RepoSummary) DBFind(ctx context.Context, db DB, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("name"),
		orm.Column("description"),
		orm.Column("ref_type"),
		orm.Column("data_group_id"),
		orm.Column("user_ids"),
		orm.Column("commits"),
		orm.Column("additions"),
		orm.Column("deletions"),
		orm.Column("latest_commit_date"),
		orm.Table(RepoSummaryTableName),
	}
	if len(_params) > 0 {
		for _, param := range _params {
			params = append(params, param)
		}
	}
	q, p := orm.BuildQuery(params...)
	row := db.QueryRowContext(ctx, q, p...)
	var _ID sql.NullString
	var _Name sql.NullString
	var _Description sql.NullString
	var _RefType sql.NullString
	var _DataGroupID sql.NullString
	var _UserIds sql.NullString
	var _Commits sql.NullInt64
	var _Additions sql.NullInt64
	var _Deletions sql.NullInt64
	var _LatestCommitDate sql.NullInt64
	err := row.Scan(
		&_ID,
		&_Name,
		&_Description,
		&_RefType,
		&_DataGroupID,
		&_UserIds,
		&_Commits,
		&_Additions,
		&_Deletions,
		&_LatestCommitDate,
	)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _Description.Valid {
		t.SetDescription(_Description.String)
	}
	if _RefType.Valid {
		t.SetRefType(_RefType.String)
	}
	if _DataGroupID.Valid {
		t.SetDataGroupID(_DataGroupID.String)
	}
	if _UserIds.Valid {
		t.SetUserIds(_UserIds.String)
	}
	if _Commits.Valid {
		t.SetCommits(_Commits.Int64)
	}
	if _Additions.Valid {
		t.SetAdditions(_Additions.Int64)
	}
	if _Deletions.Valid {
		t.SetDeletions(_Deletions.Int64)
	}
	if _LatestCommitDate.Valid {
		t.SetLatestCommitDate(_LatestCommitDate.Int64)
	}
	return true, nil
}

// DBFindTx will find a RepoSummary record in the database with the provided parameters using the provided transaction
func (t *RepoSummary) DBFindTx(ctx context.Context, tx Tx, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("name"),
		orm.Column("description"),
		orm.Column("ref_type"),
		orm.Column("data_group_id"),
		orm.Column("user_ids"),
		orm.Column("commits"),
		orm.Column("additions"),
		orm.Column("deletions"),
		orm.Column("latest_commit_date"),
		orm.Table(RepoSummaryTableName),
	}
	if len(_params) > 0 {
		for _, param := range _params {
			params = append(params, param)
		}
	}
	q, p := orm.BuildQuery(params...)
	row := tx.QueryRowContext(ctx, q, p...)
	var _ID sql.NullString
	var _Name sql.NullString
	var _Description sql.NullString
	var _RefType sql.NullString
	var _DataGroupID sql.NullString
	var _UserIds sql.NullString
	var _Commits sql.NullInt64
	var _Additions sql.NullInt64
	var _Deletions sql.NullInt64
	var _LatestCommitDate sql.NullInt64
	err := row.Scan(
		&_ID,
		&_Name,
		&_Description,
		&_RefType,
		&_DataGroupID,
		&_UserIds,
		&_Commits,
		&_Additions,
		&_Deletions,
		&_LatestCommitDate,
	)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _Description.Valid {
		t.SetDescription(_Description.String)
	}
	if _RefType.Valid {
		t.SetRefType(_RefType.String)
	}
	if _DataGroupID.Valid {
		t.SetDataGroupID(_DataGroupID.String)
	}
	if _UserIds.Valid {
		t.SetUserIds(_UserIds.String)
	}
	if _Commits.Valid {
		t.SetCommits(_Commits.Int64)
	}
	if _Additions.Valid {
		t.SetAdditions(_Additions.Int64)
	}
	if _Deletions.Valid {
		t.SetDeletions(_Deletions.Int64)
	}
	if _LatestCommitDate.Valid {
		t.SetLatestCommitDate(_LatestCommitDate.Int64)
	}
	return true, nil
}

// CountRepoSummaries will find the count of RepoSummary records in the database
func CountRepoSummaries(ctx context.Context, db DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(RepoSummaryTableName),
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

// CountRepoSummariesTx will find the count of RepoSummary records in the database using the provided transaction
func CountRepoSummariesTx(ctx context.Context, tx Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(RepoSummaryTableName),
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

// DBCount will find the count of RepoSummary records in the database
func (t *RepoSummary) DBCount(ctx context.Context, db DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(RepoSummaryTableName),
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

// DBCountTx will find the count of RepoSummary records in the database using the provided transaction
func (t *RepoSummary) DBCountTx(ctx context.Context, tx Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(RepoSummaryTableName),
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

// DBExists will return true if the RepoSummary record exists in the database
func (t *RepoSummary) DBExists(ctx context.Context, db DB) (bool, error) {
	q := "SELECT `id` FROM `repo_summary` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := db.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// DBExistsTx will return true if the RepoSummary record exists in the database using the provided transaction
func (t *RepoSummary) DBExistsTx(ctx context.Context, tx Tx) (bool, error) {
	q := "SELECT `id` FROM `repo_summary` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := tx.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// PrimaryKeyColumn returns the column name for the primary key
func (t *RepoSummary) PrimaryKeyColumn() string {
	return RepoSummaryColumnID
}

// PrimaryKeyColumnType returns the primary key column Go type as a string
func (t *RepoSummary) PrimaryKeyColumnType() string {
	return "string"
}

// PrimaryKey returns the primary key column value
func (t *RepoSummary) PrimaryKey() interface{} {
	return t.ID
}
