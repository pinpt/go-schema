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

var _ Model = (*NewrelicServer)(nil)
var _ CSVWriter = (*NewrelicServer)(nil)
var _ JSONWriter = (*NewrelicServer)(nil)
var _ Checksum = (*NewrelicServer)(nil)

// NewrelicServerTableName is the name of the table in SQL
const NewrelicServerTableName = "newrelic_server"

var NewrelicServerColumns = []string{
	"id",
	"checksum",
	"customer_id",
	"ext_id",
	"account_id",
	"name",
	"host",
	"health_status",
	"reporting",
	"last_reported_at",
	"summary_cpu",
	"summary_cpu_stolen",
	"summary_disk_io",
	"summary_memory",
	"summary_memory_used",
	"summary_memory_total",
	"summary_fullest_disk",
	"summary_fullest_disk_free",
}

// NewrelicServer table
type NewrelicServer struct {
	AccountID              int64    `json:"account_id"`
	Checksum               *string  `json:"checksum,omitempty"`
	CustomerID             string   `json:"customer_id"`
	ExtID                  int64    `json:"ext_id"`
	HealthStatus           *string  `json:"health_status,omitempty"`
	Host                   string   `json:"host"`
	ID                     string   `json:"id"`
	LastReportedAt         int64    `json:"last_reported_at"`
	Name                   string   `json:"name"`
	Reporting              bool     `json:"reporting"`
	SummaryCPU             *float64 `json:"summary_cpu,omitempty"`
	SummaryCPUStolen       *float64 `json:"summary_cpu_stolen,omitempty"`
	SummaryDiskIo          *float64 `json:"summary_disk_io,omitempty"`
	SummaryFullestDisk     *float64 `json:"summary_fullest_disk,omitempty"`
	SummaryFullestDiskFree *int64   `json:"summary_fullest_disk_free,omitempty"`
	SummaryMemory          *float64 `json:"summary_memory,omitempty"`
	SummaryMemoryTotal     *int64   `json:"summary_memory_total,omitempty"`
	SummaryMemoryUsed      *int64   `json:"summary_memory_used,omitempty"`
}

// TableName returns the SQL table name for NewrelicServer and satifies the Model interface
func (t *NewrelicServer) TableName() string {
	return NewrelicServerTableName
}

// ToCSV will serialize the NewrelicServer instance to a CSV compatible array of strings
func (t *NewrelicServer) ToCSV() []string {
	return []string{
		t.ID,
		t.CalculateChecksum(),
		t.CustomerID,
		toCSVString(t.ExtID),
		toCSVString(t.AccountID),
		t.Name,
		t.Host,
		toCSVString(t.HealthStatus),
		toCSVBool(t.Reporting),
		toCSVString(t.LastReportedAt),
		toCSVString(t.SummaryCPU),
		toCSVString(t.SummaryCPUStolen),
		toCSVString(t.SummaryDiskIo),
		toCSVString(t.SummaryMemory),
		toCSVString(t.SummaryMemoryUsed),
		toCSVString(t.SummaryMemoryTotal),
		toCSVString(t.SummaryFullestDisk),
		toCSVString(t.SummaryFullestDiskFree),
	}
}

// WriteCSV will serialize the NewrelicServer instance to the writer as CSV and satisfies the CSVWriter interface
func (t *NewrelicServer) WriteCSV(w *csv.Writer) error {
	return w.Write(t.ToCSV())
}

// WriteJSON will serialize the NewrelicServer instance to the writer as JSON and satisfies the JSONWriter interface
func (t *NewrelicServer) WriteJSON(w io.Writer, indent ...bool) error {
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

// NewNewrelicServerReader creates a JSON reader which can read in NewrelicServer objects serialized as JSON either as an array, single object or json new lines
// and writes each NewrelicServer to the channel provided
func NewNewrelicServerReader(r io.Reader, ch chan<- NewrelicServer) error {
	return orm.Deserialize(r, func(buf json.RawMessage) error {
		dec := json.NewDecoder(bytes.NewBuffer(buf))
		e := NewrelicServer{}
		if err := dec.Decode(&e); err != nil {
			return err
		}
		ch <- e
		return nil
	})
}

// NewCSVNewrelicServerReaderDir will read the reader as CSV and emit each record to the channel provided
func NewCSVNewrelicServerReader(r io.Reader, ch chan<- NewrelicServer) error {
	cr := csv.NewReader(r)
	for {
		record, err := cr.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		ch <- NewrelicServer{
			ID:                     record[0],
			Checksum:               fromStringPointer(record[1]),
			CustomerID:             record[2],
			ExtID:                  fromCSVInt64(record[3]),
			AccountID:              fromCSVInt64(record[4]),
			Name:                   record[5],
			Host:                   record[6],
			HealthStatus:           fromStringPointer(record[7]),
			Reporting:              fromCSVBool(record[8]),
			LastReportedAt:         fromCSVInt64(record[9]),
			SummaryCPU:             fromCSVFloat64Pointer(record[10]),
			SummaryCPUStolen:       fromCSVFloat64Pointer(record[11]),
			SummaryDiskIo:          fromCSVFloat64Pointer(record[12]),
			SummaryMemory:          fromCSVFloat64Pointer(record[13]),
			SummaryMemoryUsed:      fromCSVInt64Pointer(record[14]),
			SummaryMemoryTotal:     fromCSVInt64Pointer(record[15]),
			SummaryFullestDisk:     fromCSVFloat64Pointer(record[16]),
			SummaryFullestDiskFree: fromCSVInt64Pointer(record[17]),
		}
	}
	return nil
}

// NewCSVNewrelicServerReaderFile will read the file as a CSV and emit each record to the channel provided
func NewCSVNewrelicServerReaderFile(fp string, ch chan<- NewrelicServer) error {
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
	return NewCSVNewrelicServerReader(fc, ch)
}

// NewCSVNewrelicServerReaderDir will read the newrelic_server.csv.gz file as a CSV and emit each record to the channel provided
func NewCSVNewrelicServerReaderDir(dir string, ch chan<- NewrelicServer) error {
	return NewCSVNewrelicServerReaderFile(filepath.Join(dir, "newrelic_server.csv.gz"), ch)
}

// NewrelicServerCSVDeduper is a function callback which takes the existing value (a) and the new value (b)
// and the return value should be the one to use (or a new one, if applicable). return nil
// to skip processing of this record
type NewrelicServerCSVDeduper func(a NewrelicServer, b NewrelicServer) *NewrelicServer

// NewrelicServerCSVDedupeDisabled is set on whether the CSV writer should de-dupe values by key
var NewrelicServerCSVDedupeDisabled bool

// NewNewrelicServerCSVWriterSize creates a batch writer that will write each NewrelicServer into a CSV file
// this method will automatically de-duplicate entries using the primary key. if the checksum
// for a newer item with the same primary key doesn't match a previously sent item, the newer
// one will be used
func NewNewrelicServerCSVWriterSize(w io.Writer, size int, dedupers ...NewrelicServerCSVDeduper) (chan NewrelicServer, chan bool, error) {
	cw := csv.NewWriter(w)
	ch := make(chan NewrelicServer, size)
	done := make(chan bool)
	go func() {
		defer func() { done <- true }()
		dodedupe := !NewrelicServerCSVDedupeDisabled
		var kv map[string]*NewrelicServer
		var deduper NewrelicServerCSVDeduper
		if dedupers != nil && len(dedupers) > 0 {
			deduper = dedupers[0]
			dodedupe = true
		}
		if dodedupe {
			kv = make(map[string]*NewrelicServer)
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

// NewrelicServerCSVDefaultSize is the default channel buffer size if not provided
var NewrelicServerCSVDefaultSize = 100

// NewNewrelicServerCSVWriter creates a batch writer that will write each NewrelicServer into a CSV file
func NewNewrelicServerCSVWriter(w io.Writer, dedupers ...NewrelicServerCSVDeduper) (chan NewrelicServer, chan bool, error) {
	return NewNewrelicServerCSVWriterSize(w, NewrelicServerCSVDefaultSize, dedupers...)
}

// NewNewrelicServerCSVWriterDir creates a batch writer that will write each NewrelicServer into a CSV file named newrelic_server.csv.gz in dir
func NewNewrelicServerCSVWriterDir(dir string, dedupers ...NewrelicServerCSVDeduper) (chan NewrelicServer, chan bool, error) {
	return NewNewrelicServerCSVWriterFile(filepath.Join(dir, "newrelic_server.csv.gz"), dedupers...)
}

// NewNewrelicServerCSVWriterFile creates a batch writer that will write each NewrelicServer into a CSV file
func NewNewrelicServerCSVWriterFile(fn string, dedupers ...NewrelicServerCSVDeduper) (chan NewrelicServer, chan bool, error) {
	f, err := os.Create(fn)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening CSV file %s. %v", fn, err)
	}
	var fc io.WriteCloser = f
	if filepath.Ext(fn) == ".gz" {
		w, _ := gzip.NewWriterLevel(f, gzip.BestCompression)
		fc = w
	}
	ch, done, err := NewNewrelicServerCSVWriter(fc, dedupers...)
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

type NewrelicServerDBAction func(ctx context.Context, db *sql.DB, record NewrelicServer) error

// NewNewrelicServerDBWriterSize creates a DB writer that will write each issue into the DB
func NewNewrelicServerDBWriterSize(ctx context.Context, db *sql.DB, errors chan<- error, size int, actions ...NewrelicServerDBAction) (chan NewrelicServer, chan bool, error) {
	ch := make(chan NewrelicServer, size)
	done := make(chan bool)
	var action NewrelicServerDBAction
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

// NewNewrelicServerDBWriter creates a DB writer that will write each issue into the DB
func NewNewrelicServerDBWriter(ctx context.Context, db *sql.DB, errors chan<- error, actions ...NewrelicServerDBAction) (chan NewrelicServer, chan bool, error) {
	return NewNewrelicServerDBWriterSize(ctx, db, errors, 100, actions...)
}

// NewrelicServerColumnID is the ID SQL column name for the NewrelicServer table
const NewrelicServerColumnID = "id"

// NewrelicServerEscapedColumnID is the escaped ID SQL column name for the NewrelicServer table
const NewrelicServerEscapedColumnID = "`id`"

// NewrelicServerColumnChecksum is the Checksum SQL column name for the NewrelicServer table
const NewrelicServerColumnChecksum = "checksum"

// NewrelicServerEscapedColumnChecksum is the escaped Checksum SQL column name for the NewrelicServer table
const NewrelicServerEscapedColumnChecksum = "`checksum`"

// NewrelicServerColumnCustomerID is the CustomerID SQL column name for the NewrelicServer table
const NewrelicServerColumnCustomerID = "customer_id"

// NewrelicServerEscapedColumnCustomerID is the escaped CustomerID SQL column name for the NewrelicServer table
const NewrelicServerEscapedColumnCustomerID = "`customer_id`"

// NewrelicServerColumnExtID is the ExtID SQL column name for the NewrelicServer table
const NewrelicServerColumnExtID = "ext_id"

// NewrelicServerEscapedColumnExtID is the escaped ExtID SQL column name for the NewrelicServer table
const NewrelicServerEscapedColumnExtID = "`ext_id`"

// NewrelicServerColumnAccountID is the AccountID SQL column name for the NewrelicServer table
const NewrelicServerColumnAccountID = "account_id"

// NewrelicServerEscapedColumnAccountID is the escaped AccountID SQL column name for the NewrelicServer table
const NewrelicServerEscapedColumnAccountID = "`account_id`"

// NewrelicServerColumnName is the Name SQL column name for the NewrelicServer table
const NewrelicServerColumnName = "name"

// NewrelicServerEscapedColumnName is the escaped Name SQL column name for the NewrelicServer table
const NewrelicServerEscapedColumnName = "`name`"

// NewrelicServerColumnHost is the Host SQL column name for the NewrelicServer table
const NewrelicServerColumnHost = "host"

// NewrelicServerEscapedColumnHost is the escaped Host SQL column name for the NewrelicServer table
const NewrelicServerEscapedColumnHost = "`host`"

// NewrelicServerColumnHealthStatus is the HealthStatus SQL column name for the NewrelicServer table
const NewrelicServerColumnHealthStatus = "health_status"

// NewrelicServerEscapedColumnHealthStatus is the escaped HealthStatus SQL column name for the NewrelicServer table
const NewrelicServerEscapedColumnHealthStatus = "`health_status`"

// NewrelicServerColumnReporting is the Reporting SQL column name for the NewrelicServer table
const NewrelicServerColumnReporting = "reporting"

// NewrelicServerEscapedColumnReporting is the escaped Reporting SQL column name for the NewrelicServer table
const NewrelicServerEscapedColumnReporting = "`reporting`"

// NewrelicServerColumnLastReportedAt is the LastReportedAt SQL column name for the NewrelicServer table
const NewrelicServerColumnLastReportedAt = "last_reported_at"

// NewrelicServerEscapedColumnLastReportedAt is the escaped LastReportedAt SQL column name for the NewrelicServer table
const NewrelicServerEscapedColumnLastReportedAt = "`last_reported_at`"

// NewrelicServerColumnSummaryCPU is the SummaryCPU SQL column name for the NewrelicServer table
const NewrelicServerColumnSummaryCPU = "summary_cpu"

// NewrelicServerEscapedColumnSummaryCPU is the escaped SummaryCPU SQL column name for the NewrelicServer table
const NewrelicServerEscapedColumnSummaryCPU = "`summary_cpu`"

// NewrelicServerColumnSummaryCPUStolen is the SummaryCPUStolen SQL column name for the NewrelicServer table
const NewrelicServerColumnSummaryCPUStolen = "summary_cpu_stolen"

// NewrelicServerEscapedColumnSummaryCPUStolen is the escaped SummaryCPUStolen SQL column name for the NewrelicServer table
const NewrelicServerEscapedColumnSummaryCPUStolen = "`summary_cpu_stolen`"

// NewrelicServerColumnSummaryDiskIo is the SummaryDiskIo SQL column name for the NewrelicServer table
const NewrelicServerColumnSummaryDiskIo = "summary_disk_io"

// NewrelicServerEscapedColumnSummaryDiskIo is the escaped SummaryDiskIo SQL column name for the NewrelicServer table
const NewrelicServerEscapedColumnSummaryDiskIo = "`summary_disk_io`"

// NewrelicServerColumnSummaryMemory is the SummaryMemory SQL column name for the NewrelicServer table
const NewrelicServerColumnSummaryMemory = "summary_memory"

// NewrelicServerEscapedColumnSummaryMemory is the escaped SummaryMemory SQL column name for the NewrelicServer table
const NewrelicServerEscapedColumnSummaryMemory = "`summary_memory`"

// NewrelicServerColumnSummaryMemoryUsed is the SummaryMemoryUsed SQL column name for the NewrelicServer table
const NewrelicServerColumnSummaryMemoryUsed = "summary_memory_used"

// NewrelicServerEscapedColumnSummaryMemoryUsed is the escaped SummaryMemoryUsed SQL column name for the NewrelicServer table
const NewrelicServerEscapedColumnSummaryMemoryUsed = "`summary_memory_used`"

// NewrelicServerColumnSummaryMemoryTotal is the SummaryMemoryTotal SQL column name for the NewrelicServer table
const NewrelicServerColumnSummaryMemoryTotal = "summary_memory_total"

// NewrelicServerEscapedColumnSummaryMemoryTotal is the escaped SummaryMemoryTotal SQL column name for the NewrelicServer table
const NewrelicServerEscapedColumnSummaryMemoryTotal = "`summary_memory_total`"

// NewrelicServerColumnSummaryFullestDisk is the SummaryFullestDisk SQL column name for the NewrelicServer table
const NewrelicServerColumnSummaryFullestDisk = "summary_fullest_disk"

// NewrelicServerEscapedColumnSummaryFullestDisk is the escaped SummaryFullestDisk SQL column name for the NewrelicServer table
const NewrelicServerEscapedColumnSummaryFullestDisk = "`summary_fullest_disk`"

// NewrelicServerColumnSummaryFullestDiskFree is the SummaryFullestDiskFree SQL column name for the NewrelicServer table
const NewrelicServerColumnSummaryFullestDiskFree = "summary_fullest_disk_free"

// NewrelicServerEscapedColumnSummaryFullestDiskFree is the escaped SummaryFullestDiskFree SQL column name for the NewrelicServer table
const NewrelicServerEscapedColumnSummaryFullestDiskFree = "`summary_fullest_disk_free`"

// GetID will return the NewrelicServer ID value
func (t *NewrelicServer) GetID() string {
	return t.ID
}

// SetID will set the NewrelicServer ID value
func (t *NewrelicServer) SetID(v string) {
	t.ID = v
}

// FindNewrelicServerByID will find a NewrelicServer by ID
func FindNewrelicServerByID(ctx context.Context, db *sql.DB, value string) (*NewrelicServer, error) {
	q := "SELECT `newrelic_server`.`id`,`newrelic_server`.`checksum`,`newrelic_server`.`customer_id`,`newrelic_server`.`ext_id`,`newrelic_server`.`account_id`,`newrelic_server`.`name`,`newrelic_server`.`host`,`newrelic_server`.`health_status`,`newrelic_server`.`reporting`,`newrelic_server`.`last_reported_at`,`newrelic_server`.`summary_cpu`,`newrelic_server`.`summary_cpu_stolen`,`newrelic_server`.`summary_disk_io`,`newrelic_server`.`summary_memory`,`newrelic_server`.`summary_memory_used`,`newrelic_server`.`summary_memory_total`,`newrelic_server`.`summary_fullest_disk`,`newrelic_server`.`summary_fullest_disk_free` FROM `newrelic_server` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CustomerID sql.NullString
	var _ExtID sql.NullInt64
	var _AccountID sql.NullInt64
	var _Name sql.NullString
	var _Host sql.NullString
	var _HealthStatus sql.NullString
	var _Reporting sql.NullBool
	var _LastReportedAt sql.NullInt64
	var _SummaryCPU sql.NullFloat64
	var _SummaryCPUStolen sql.NullFloat64
	var _SummaryDiskIo sql.NullFloat64
	var _SummaryMemory sql.NullFloat64
	var _SummaryMemoryUsed sql.NullInt64
	var _SummaryMemoryTotal sql.NullInt64
	var _SummaryFullestDisk sql.NullFloat64
	var _SummaryFullestDiskFree sql.NullInt64
	err := db.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_ExtID,
		&_AccountID,
		&_Name,
		&_Host,
		&_HealthStatus,
		&_Reporting,
		&_LastReportedAt,
		&_SummaryCPU,
		&_SummaryCPUStolen,
		&_SummaryDiskIo,
		&_SummaryMemory,
		&_SummaryMemoryUsed,
		&_SummaryMemoryTotal,
		&_SummaryFullestDisk,
		&_SummaryFullestDiskFree,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &NewrelicServer{}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _Checksum.Valid {
		t.SetChecksum(_Checksum.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _ExtID.Valid {
		t.SetExtID(_ExtID.Int64)
	}
	if _AccountID.Valid {
		t.SetAccountID(_AccountID.Int64)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _Host.Valid {
		t.SetHost(_Host.String)
	}
	if _HealthStatus.Valid {
		t.SetHealthStatus(_HealthStatus.String)
	}
	if _Reporting.Valid {
		t.SetReporting(_Reporting.Bool)
	}
	if _LastReportedAt.Valid {
		t.SetLastReportedAt(_LastReportedAt.Int64)
	}
	if _SummaryCPU.Valid {
		t.SetSummaryCPU(_SummaryCPU.Float64)
	}
	if _SummaryCPUStolen.Valid {
		t.SetSummaryCPUStolen(_SummaryCPUStolen.Float64)
	}
	if _SummaryDiskIo.Valid {
		t.SetSummaryDiskIo(_SummaryDiskIo.Float64)
	}
	if _SummaryMemory.Valid {
		t.SetSummaryMemory(_SummaryMemory.Float64)
	}
	if _SummaryMemoryUsed.Valid {
		t.SetSummaryMemoryUsed(_SummaryMemoryUsed.Int64)
	}
	if _SummaryMemoryTotal.Valid {
		t.SetSummaryMemoryTotal(_SummaryMemoryTotal.Int64)
	}
	if _SummaryFullestDisk.Valid {
		t.SetSummaryFullestDisk(_SummaryFullestDisk.Float64)
	}
	if _SummaryFullestDiskFree.Valid {
		t.SetSummaryFullestDiskFree(_SummaryFullestDiskFree.Int64)
	}
	return t, nil
}

// FindNewrelicServerByIDTx will find a NewrelicServer by ID using the provided transaction
func FindNewrelicServerByIDTx(ctx context.Context, tx *sql.Tx, value string) (*NewrelicServer, error) {
	q := "SELECT `newrelic_server`.`id`,`newrelic_server`.`checksum`,`newrelic_server`.`customer_id`,`newrelic_server`.`ext_id`,`newrelic_server`.`account_id`,`newrelic_server`.`name`,`newrelic_server`.`host`,`newrelic_server`.`health_status`,`newrelic_server`.`reporting`,`newrelic_server`.`last_reported_at`,`newrelic_server`.`summary_cpu`,`newrelic_server`.`summary_cpu_stolen`,`newrelic_server`.`summary_disk_io`,`newrelic_server`.`summary_memory`,`newrelic_server`.`summary_memory_used`,`newrelic_server`.`summary_memory_total`,`newrelic_server`.`summary_fullest_disk`,`newrelic_server`.`summary_fullest_disk_free` FROM `newrelic_server` WHERE `id` = ?"
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CustomerID sql.NullString
	var _ExtID sql.NullInt64
	var _AccountID sql.NullInt64
	var _Name sql.NullString
	var _Host sql.NullString
	var _HealthStatus sql.NullString
	var _Reporting sql.NullBool
	var _LastReportedAt sql.NullInt64
	var _SummaryCPU sql.NullFloat64
	var _SummaryCPUStolen sql.NullFloat64
	var _SummaryDiskIo sql.NullFloat64
	var _SummaryMemory sql.NullFloat64
	var _SummaryMemoryUsed sql.NullInt64
	var _SummaryMemoryTotal sql.NullInt64
	var _SummaryFullestDisk sql.NullFloat64
	var _SummaryFullestDiskFree sql.NullInt64
	err := tx.QueryRowContext(ctx, q, value).Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_ExtID,
		&_AccountID,
		&_Name,
		&_Host,
		&_HealthStatus,
		&_Reporting,
		&_LastReportedAt,
		&_SummaryCPU,
		&_SummaryCPUStolen,
		&_SummaryDiskIo,
		&_SummaryMemory,
		&_SummaryMemoryUsed,
		&_SummaryMemoryTotal,
		&_SummaryFullestDisk,
		&_SummaryFullestDiskFree,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	t := &NewrelicServer{}
	if _ID.Valid {
		t.SetID(_ID.String)
	}
	if _Checksum.Valid {
		t.SetChecksum(_Checksum.String)
	}
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _ExtID.Valid {
		t.SetExtID(_ExtID.Int64)
	}
	if _AccountID.Valid {
		t.SetAccountID(_AccountID.Int64)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _Host.Valid {
		t.SetHost(_Host.String)
	}
	if _HealthStatus.Valid {
		t.SetHealthStatus(_HealthStatus.String)
	}
	if _Reporting.Valid {
		t.SetReporting(_Reporting.Bool)
	}
	if _LastReportedAt.Valid {
		t.SetLastReportedAt(_LastReportedAt.Int64)
	}
	if _SummaryCPU.Valid {
		t.SetSummaryCPU(_SummaryCPU.Float64)
	}
	if _SummaryCPUStolen.Valid {
		t.SetSummaryCPUStolen(_SummaryCPUStolen.Float64)
	}
	if _SummaryDiskIo.Valid {
		t.SetSummaryDiskIo(_SummaryDiskIo.Float64)
	}
	if _SummaryMemory.Valid {
		t.SetSummaryMemory(_SummaryMemory.Float64)
	}
	if _SummaryMemoryUsed.Valid {
		t.SetSummaryMemoryUsed(_SummaryMemoryUsed.Int64)
	}
	if _SummaryMemoryTotal.Valid {
		t.SetSummaryMemoryTotal(_SummaryMemoryTotal.Int64)
	}
	if _SummaryFullestDisk.Valid {
		t.SetSummaryFullestDisk(_SummaryFullestDisk.Float64)
	}
	if _SummaryFullestDiskFree.Valid {
		t.SetSummaryFullestDiskFree(_SummaryFullestDiskFree.Int64)
	}
	return t, nil
}

// GetChecksum will return the NewrelicServer Checksum value
func (t *NewrelicServer) GetChecksum() string {
	if t.Checksum == nil {
		return ""
	}
	return *t.Checksum
}

// SetChecksum will set the NewrelicServer Checksum value
func (t *NewrelicServer) SetChecksum(v string) {
	t.Checksum = &v
}

// GetCustomerID will return the NewrelicServer CustomerID value
func (t *NewrelicServer) GetCustomerID() string {
	return t.CustomerID
}

// SetCustomerID will set the NewrelicServer CustomerID value
func (t *NewrelicServer) SetCustomerID(v string) {
	t.CustomerID = v
}

// FindNewrelicServersByCustomerID will find all NewrelicServers by the CustomerID value
func FindNewrelicServersByCustomerID(ctx context.Context, db *sql.DB, value string) ([]*NewrelicServer, error) {
	q := "SELECT `newrelic_server`.`id`,`newrelic_server`.`checksum`,`newrelic_server`.`customer_id`,`newrelic_server`.`ext_id`,`newrelic_server`.`account_id`,`newrelic_server`.`name`,`newrelic_server`.`host`,`newrelic_server`.`health_status`,`newrelic_server`.`reporting`,`newrelic_server`.`last_reported_at`,`newrelic_server`.`summary_cpu`,`newrelic_server`.`summary_cpu_stolen`,`newrelic_server`.`summary_disk_io`,`newrelic_server`.`summary_memory`,`newrelic_server`.`summary_memory_used`,`newrelic_server`.`summary_memory_total`,`newrelic_server`.`summary_fullest_disk`,`newrelic_server`.`summary_fullest_disk_free` FROM `newrelic_server` WHERE `customer_id` = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*NewrelicServer, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _ExtID sql.NullInt64
		var _AccountID sql.NullInt64
		var _Name sql.NullString
		var _Host sql.NullString
		var _HealthStatus sql.NullString
		var _Reporting sql.NullBool
		var _LastReportedAt sql.NullInt64
		var _SummaryCPU sql.NullFloat64
		var _SummaryCPUStolen sql.NullFloat64
		var _SummaryDiskIo sql.NullFloat64
		var _SummaryMemory sql.NullFloat64
		var _SummaryMemoryUsed sql.NullInt64
		var _SummaryMemoryTotal sql.NullInt64
		var _SummaryFullestDisk sql.NullFloat64
		var _SummaryFullestDiskFree sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_ExtID,
			&_AccountID,
			&_Name,
			&_Host,
			&_HealthStatus,
			&_Reporting,
			&_LastReportedAt,
			&_SummaryCPU,
			&_SummaryCPUStolen,
			&_SummaryDiskIo,
			&_SummaryMemory,
			&_SummaryMemoryUsed,
			&_SummaryMemoryTotal,
			&_SummaryFullestDisk,
			&_SummaryFullestDiskFree,
		)
		if err != nil {
			return nil, err
		}
		t := &NewrelicServer{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _ExtID.Valid {
			t.SetExtID(_ExtID.Int64)
		}
		if _AccountID.Valid {
			t.SetAccountID(_AccountID.Int64)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _Host.Valid {
			t.SetHost(_Host.String)
		}
		if _HealthStatus.Valid {
			t.SetHealthStatus(_HealthStatus.String)
		}
		if _Reporting.Valid {
			t.SetReporting(_Reporting.Bool)
		}
		if _LastReportedAt.Valid {
			t.SetLastReportedAt(_LastReportedAt.Int64)
		}
		if _SummaryCPU.Valid {
			t.SetSummaryCPU(_SummaryCPU.Float64)
		}
		if _SummaryCPUStolen.Valid {
			t.SetSummaryCPUStolen(_SummaryCPUStolen.Float64)
		}
		if _SummaryDiskIo.Valid {
			t.SetSummaryDiskIo(_SummaryDiskIo.Float64)
		}
		if _SummaryMemory.Valid {
			t.SetSummaryMemory(_SummaryMemory.Float64)
		}
		if _SummaryMemoryUsed.Valid {
			t.SetSummaryMemoryUsed(_SummaryMemoryUsed.Int64)
		}
		if _SummaryMemoryTotal.Valid {
			t.SetSummaryMemoryTotal(_SummaryMemoryTotal.Int64)
		}
		if _SummaryFullestDisk.Valid {
			t.SetSummaryFullestDisk(_SummaryFullestDisk.Float64)
		}
		if _SummaryFullestDiskFree.Valid {
			t.SetSummaryFullestDiskFree(_SummaryFullestDiskFree.Int64)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindNewrelicServersByCustomerIDTx will find all NewrelicServers by the CustomerID value using the provided transaction
func FindNewrelicServersByCustomerIDTx(ctx context.Context, tx *sql.Tx, value string) ([]*NewrelicServer, error) {
	q := "SELECT `newrelic_server`.`id`,`newrelic_server`.`checksum`,`newrelic_server`.`customer_id`,`newrelic_server`.`ext_id`,`newrelic_server`.`account_id`,`newrelic_server`.`name`,`newrelic_server`.`host`,`newrelic_server`.`health_status`,`newrelic_server`.`reporting`,`newrelic_server`.`last_reported_at`,`newrelic_server`.`summary_cpu`,`newrelic_server`.`summary_cpu_stolen`,`newrelic_server`.`summary_disk_io`,`newrelic_server`.`summary_memory`,`newrelic_server`.`summary_memory_used`,`newrelic_server`.`summary_memory_total`,`newrelic_server`.`summary_fullest_disk`,`newrelic_server`.`summary_fullest_disk_free` FROM `newrelic_server` WHERE `customer_id` = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, q, orm.ToSQLString(value))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]*NewrelicServer, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _ExtID sql.NullInt64
		var _AccountID sql.NullInt64
		var _Name sql.NullString
		var _Host sql.NullString
		var _HealthStatus sql.NullString
		var _Reporting sql.NullBool
		var _LastReportedAt sql.NullInt64
		var _SummaryCPU sql.NullFloat64
		var _SummaryCPUStolen sql.NullFloat64
		var _SummaryDiskIo sql.NullFloat64
		var _SummaryMemory sql.NullFloat64
		var _SummaryMemoryUsed sql.NullInt64
		var _SummaryMemoryTotal sql.NullInt64
		var _SummaryFullestDisk sql.NullFloat64
		var _SummaryFullestDiskFree sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_ExtID,
			&_AccountID,
			&_Name,
			&_Host,
			&_HealthStatus,
			&_Reporting,
			&_LastReportedAt,
			&_SummaryCPU,
			&_SummaryCPUStolen,
			&_SummaryDiskIo,
			&_SummaryMemory,
			&_SummaryMemoryUsed,
			&_SummaryMemoryTotal,
			&_SummaryFullestDisk,
			&_SummaryFullestDiskFree,
		)
		if err != nil {
			return nil, err
		}
		t := &NewrelicServer{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _ExtID.Valid {
			t.SetExtID(_ExtID.Int64)
		}
		if _AccountID.Valid {
			t.SetAccountID(_AccountID.Int64)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _Host.Valid {
			t.SetHost(_Host.String)
		}
		if _HealthStatus.Valid {
			t.SetHealthStatus(_HealthStatus.String)
		}
		if _Reporting.Valid {
			t.SetReporting(_Reporting.Bool)
		}
		if _LastReportedAt.Valid {
			t.SetLastReportedAt(_LastReportedAt.Int64)
		}
		if _SummaryCPU.Valid {
			t.SetSummaryCPU(_SummaryCPU.Float64)
		}
		if _SummaryCPUStolen.Valid {
			t.SetSummaryCPUStolen(_SummaryCPUStolen.Float64)
		}
		if _SummaryDiskIo.Valid {
			t.SetSummaryDiskIo(_SummaryDiskIo.Float64)
		}
		if _SummaryMemory.Valid {
			t.SetSummaryMemory(_SummaryMemory.Float64)
		}
		if _SummaryMemoryUsed.Valid {
			t.SetSummaryMemoryUsed(_SummaryMemoryUsed.Int64)
		}
		if _SummaryMemoryTotal.Valid {
			t.SetSummaryMemoryTotal(_SummaryMemoryTotal.Int64)
		}
		if _SummaryFullestDisk.Valid {
			t.SetSummaryFullestDisk(_SummaryFullestDisk.Float64)
		}
		if _SummaryFullestDiskFree.Valid {
			t.SetSummaryFullestDiskFree(_SummaryFullestDiskFree.Int64)
		}
		results = append(results, t)
	}
	return results, nil
}

// GetExtID will return the NewrelicServer ExtID value
func (t *NewrelicServer) GetExtID() int64 {
	return t.ExtID
}

// SetExtID will set the NewrelicServer ExtID value
func (t *NewrelicServer) SetExtID(v int64) {
	t.ExtID = v
}

// GetAccountID will return the NewrelicServer AccountID value
func (t *NewrelicServer) GetAccountID() int64 {
	return t.AccountID
}

// SetAccountID will set the NewrelicServer AccountID value
func (t *NewrelicServer) SetAccountID(v int64) {
	t.AccountID = v
}

// GetName will return the NewrelicServer Name value
func (t *NewrelicServer) GetName() string {
	return t.Name
}

// SetName will set the NewrelicServer Name value
func (t *NewrelicServer) SetName(v string) {
	t.Name = v
}

// GetHost will return the NewrelicServer Host value
func (t *NewrelicServer) GetHost() string {
	return t.Host
}

// SetHost will set the NewrelicServer Host value
func (t *NewrelicServer) SetHost(v string) {
	t.Host = v
}

// GetHealthStatus will return the NewrelicServer HealthStatus value
func (t *NewrelicServer) GetHealthStatus() string {
	if t.HealthStatus == nil {
		return ""
	}
	return *t.HealthStatus
}

// SetHealthStatus will set the NewrelicServer HealthStatus value
func (t *NewrelicServer) SetHealthStatus(v string) {
	t.HealthStatus = &v
}

// GetReporting will return the NewrelicServer Reporting value
func (t *NewrelicServer) GetReporting() bool {
	return t.Reporting
}

// SetReporting will set the NewrelicServer Reporting value
func (t *NewrelicServer) SetReporting(v bool) {
	t.Reporting = v
}

// GetLastReportedAt will return the NewrelicServer LastReportedAt value
func (t *NewrelicServer) GetLastReportedAt() int64 {
	return t.LastReportedAt
}

// SetLastReportedAt will set the NewrelicServer LastReportedAt value
func (t *NewrelicServer) SetLastReportedAt(v int64) {
	t.LastReportedAt = v
}

// GetSummaryCPU will return the NewrelicServer SummaryCPU value
func (t *NewrelicServer) GetSummaryCPU() float64 {
	if t.SummaryCPU == nil {
		return float64(0.0)
	}
	return *t.SummaryCPU
}

// SetSummaryCPU will set the NewrelicServer SummaryCPU value
func (t *NewrelicServer) SetSummaryCPU(v float64) {
	t.SummaryCPU = &v
}

// GetSummaryCPUStolen will return the NewrelicServer SummaryCPUStolen value
func (t *NewrelicServer) GetSummaryCPUStolen() float64 {
	if t.SummaryCPUStolen == nil {
		return float64(0.0)
	}
	return *t.SummaryCPUStolen
}

// SetSummaryCPUStolen will set the NewrelicServer SummaryCPUStolen value
func (t *NewrelicServer) SetSummaryCPUStolen(v float64) {
	t.SummaryCPUStolen = &v
}

// GetSummaryDiskIo will return the NewrelicServer SummaryDiskIo value
func (t *NewrelicServer) GetSummaryDiskIo() float64 {
	if t.SummaryDiskIo == nil {
		return float64(0.0)
	}
	return *t.SummaryDiskIo
}

// SetSummaryDiskIo will set the NewrelicServer SummaryDiskIo value
func (t *NewrelicServer) SetSummaryDiskIo(v float64) {
	t.SummaryDiskIo = &v
}

// GetSummaryMemory will return the NewrelicServer SummaryMemory value
func (t *NewrelicServer) GetSummaryMemory() float64 {
	if t.SummaryMemory == nil {
		return float64(0.0)
	}
	return *t.SummaryMemory
}

// SetSummaryMemory will set the NewrelicServer SummaryMemory value
func (t *NewrelicServer) SetSummaryMemory(v float64) {
	t.SummaryMemory = &v
}

// GetSummaryMemoryUsed will return the NewrelicServer SummaryMemoryUsed value
func (t *NewrelicServer) GetSummaryMemoryUsed() int64 {
	if t.SummaryMemoryUsed == nil {
		return int64(0)
	}
	return *t.SummaryMemoryUsed
}

// SetSummaryMemoryUsed will set the NewrelicServer SummaryMemoryUsed value
func (t *NewrelicServer) SetSummaryMemoryUsed(v int64) {
	t.SummaryMemoryUsed = &v
}

// GetSummaryMemoryTotal will return the NewrelicServer SummaryMemoryTotal value
func (t *NewrelicServer) GetSummaryMemoryTotal() int64 {
	if t.SummaryMemoryTotal == nil {
		return int64(0)
	}
	return *t.SummaryMemoryTotal
}

// SetSummaryMemoryTotal will set the NewrelicServer SummaryMemoryTotal value
func (t *NewrelicServer) SetSummaryMemoryTotal(v int64) {
	t.SummaryMemoryTotal = &v
}

// GetSummaryFullestDisk will return the NewrelicServer SummaryFullestDisk value
func (t *NewrelicServer) GetSummaryFullestDisk() float64 {
	if t.SummaryFullestDisk == nil {
		return float64(0.0)
	}
	return *t.SummaryFullestDisk
}

// SetSummaryFullestDisk will set the NewrelicServer SummaryFullestDisk value
func (t *NewrelicServer) SetSummaryFullestDisk(v float64) {
	t.SummaryFullestDisk = &v
}

// GetSummaryFullestDiskFree will return the NewrelicServer SummaryFullestDiskFree value
func (t *NewrelicServer) GetSummaryFullestDiskFree() int64 {
	if t.SummaryFullestDiskFree == nil {
		return int64(0)
	}
	return *t.SummaryFullestDiskFree
}

// SetSummaryFullestDiskFree will set the NewrelicServer SummaryFullestDiskFree value
func (t *NewrelicServer) SetSummaryFullestDiskFree(v int64) {
	t.SummaryFullestDiskFree = &v
}

func (t *NewrelicServer) toTimestamp(value time.Time) *timestamp.Timestamp {
	ts, _ := ptypes.TimestampProto(value)
	return ts
}

// DBCreateNewrelicServerTable will create the NewrelicServer table
func DBCreateNewrelicServerTable(ctx context.Context, db *sql.DB) error {
	q := "CREATE TABLE `newrelic_server` (`id` VARCHAR(64) NOT NULL PRIMARY KEY,`checksum` CHAR(64),`customer_id` VARCHAR(64) NOT NULL,`ext_id` BIGINT(20) NOT NULL,`account_id`BIGINT(20) NOT NULL,`name`TEXT NOT NULL,`host`TEXT NOT NULL,`health_status`TEXT,`reporting` BOOL NOT NULL,`last_reported_at`BIGINT(20) NOT NULL,`summary_cpu` DOUBLE,`summary_cpu_stolen` DOUBLE,`summary_disk_io` DOUBLE,`summary_memory` DOUBLE,`summary_memory_used`BIGINT(20),`summary_memory_total` BIGINT(20),`summary_fullest_disk` DOUBLE,`summary_fullest_disk_free` BIGINT(20),INDEX newrelic_server_customer_id_index (`customer_id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBCreateNewrelicServerTableTx will create the NewrelicServer table using the provided transction
func DBCreateNewrelicServerTableTx(ctx context.Context, tx *sql.Tx) error {
	q := "CREATE TABLE `newrelic_server` (`id` VARCHAR(64) NOT NULL PRIMARY KEY,`checksum` CHAR(64),`customer_id` VARCHAR(64) NOT NULL,`ext_id` BIGINT(20) NOT NULL,`account_id`BIGINT(20) NOT NULL,`name`TEXT NOT NULL,`host`TEXT NOT NULL,`health_status`TEXT,`reporting` BOOL NOT NULL,`last_reported_at`BIGINT(20) NOT NULL,`summary_cpu` DOUBLE,`summary_cpu_stolen` DOUBLE,`summary_disk_io` DOUBLE,`summary_memory` DOUBLE,`summary_memory_used`BIGINT(20),`summary_memory_total` BIGINT(20),`summary_fullest_disk` DOUBLE,`summary_fullest_disk_free` BIGINT(20),INDEX newrelic_server_customer_id_index (`customer_id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// DBDropNewrelicServerTable will drop the NewrelicServer table
func DBDropNewrelicServerTable(ctx context.Context, db *sql.DB) error {
	q := "DROP TABLE IF EXISTS `newrelic_server`"
	_, err := db.ExecContext(ctx, q)
	return err
}

// DBDropNewrelicServerTableTx will drop the NewrelicServer table using the provided transaction
func DBDropNewrelicServerTableTx(ctx context.Context, tx *sql.Tx) error {
	q := "DROP TABLE IF EXISTS `newrelic_server`"
	_, err := tx.ExecContext(ctx, q)
	return err
}

// CalculateChecksum will calculate a checksum of the SHA1 of all field values
func (t *NewrelicServer) CalculateChecksum() string {
	return orm.HashStrings(
		orm.ToString(t.ID),
		orm.ToString(t.CustomerID),
		orm.ToString(t.ExtID),
		orm.ToString(t.AccountID),
		orm.ToString(t.Name),
		orm.ToString(t.Host),
		orm.ToString(t.HealthStatus),
		orm.ToString(t.Reporting),
		orm.ToString(t.LastReportedAt),
		orm.ToString(t.SummaryCPU),
		orm.ToString(t.SummaryCPUStolen),
		orm.ToString(t.SummaryDiskIo),
		orm.ToString(t.SummaryMemory),
		orm.ToString(t.SummaryMemoryUsed),
		orm.ToString(t.SummaryMemoryTotal),
		orm.ToString(t.SummaryFullestDisk),
		orm.ToString(t.SummaryFullestDiskFree),
	)
}

// DBCreate will create a new NewrelicServer record in the database
func (t *NewrelicServer) DBCreate(ctx context.Context, db *sql.DB) (sql.Result, error) {
	q := "INSERT INTO `newrelic_server` (`newrelic_server`.`id`,`newrelic_server`.`checksum`,`newrelic_server`.`customer_id`,`newrelic_server`.`ext_id`,`newrelic_server`.`account_id`,`newrelic_server`.`name`,`newrelic_server`.`host`,`newrelic_server`.`health_status`,`newrelic_server`.`reporting`,`newrelic_server`.`last_reported_at`,`newrelic_server`.`summary_cpu`,`newrelic_server`.`summary_cpu_stolen`,`newrelic_server`.`summary_disk_io`,`newrelic_server`.`summary_memory`,`newrelic_server`.`summary_memory_used`,`newrelic_server`.`summary_memory_total`,`newrelic_server`.`summary_fullest_disk`,`newrelic_server`.`summary_fullest_disk_free`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLInt64(t.ExtID),
		orm.ToSQLInt64(t.AccountID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.Host),
		orm.ToSQLString(t.HealthStatus),
		orm.ToSQLBool(t.Reporting),
		orm.ToSQLInt64(t.LastReportedAt),
		orm.ToSQLFloat64(t.SummaryCPU),
		orm.ToSQLFloat64(t.SummaryCPUStolen),
		orm.ToSQLFloat64(t.SummaryDiskIo),
		orm.ToSQLFloat64(t.SummaryMemory),
		orm.ToSQLInt64(t.SummaryMemoryUsed),
		orm.ToSQLInt64(t.SummaryMemoryTotal),
		orm.ToSQLFloat64(t.SummaryFullestDisk),
		orm.ToSQLInt64(t.SummaryFullestDiskFree),
	)
}

// DBCreateTx will create a new NewrelicServer record in the database using the provided transaction
func (t *NewrelicServer) DBCreateTx(ctx context.Context, tx *sql.Tx) (sql.Result, error) {
	q := "INSERT INTO `newrelic_server` (`newrelic_server`.`id`,`newrelic_server`.`checksum`,`newrelic_server`.`customer_id`,`newrelic_server`.`ext_id`,`newrelic_server`.`account_id`,`newrelic_server`.`name`,`newrelic_server`.`host`,`newrelic_server`.`health_status`,`newrelic_server`.`reporting`,`newrelic_server`.`last_reported_at`,`newrelic_server`.`summary_cpu`,`newrelic_server`.`summary_cpu_stolen`,`newrelic_server`.`summary_disk_io`,`newrelic_server`.`summary_memory`,`newrelic_server`.`summary_memory_used`,`newrelic_server`.`summary_memory_total`,`newrelic_server`.`summary_fullest_disk`,`newrelic_server`.`summary_fullest_disk_free`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLInt64(t.ExtID),
		orm.ToSQLInt64(t.AccountID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.Host),
		orm.ToSQLString(t.HealthStatus),
		orm.ToSQLBool(t.Reporting),
		orm.ToSQLInt64(t.LastReportedAt),
		orm.ToSQLFloat64(t.SummaryCPU),
		orm.ToSQLFloat64(t.SummaryCPUStolen),
		orm.ToSQLFloat64(t.SummaryDiskIo),
		orm.ToSQLFloat64(t.SummaryMemory),
		orm.ToSQLInt64(t.SummaryMemoryUsed),
		orm.ToSQLInt64(t.SummaryMemoryTotal),
		orm.ToSQLFloat64(t.SummaryFullestDisk),
		orm.ToSQLInt64(t.SummaryFullestDiskFree),
	)
}

// DBCreateIgnoreDuplicate will upsert the NewrelicServer record in the database
func (t *NewrelicServer) DBCreateIgnoreDuplicate(ctx context.Context, db *sql.DB) (sql.Result, error) {
	q := "INSERT INTO `newrelic_server` (`newrelic_server`.`id`,`newrelic_server`.`checksum`,`newrelic_server`.`customer_id`,`newrelic_server`.`ext_id`,`newrelic_server`.`account_id`,`newrelic_server`.`name`,`newrelic_server`.`host`,`newrelic_server`.`health_status`,`newrelic_server`.`reporting`,`newrelic_server`.`last_reported_at`,`newrelic_server`.`summary_cpu`,`newrelic_server`.`summary_cpu_stolen`,`newrelic_server`.`summary_disk_io`,`newrelic_server`.`summary_memory`,`newrelic_server`.`summary_memory_used`,`newrelic_server`.`summary_memory_total`,`newrelic_server`.`summary_fullest_disk`,`newrelic_server`.`summary_fullest_disk_free`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLInt64(t.ExtID),
		orm.ToSQLInt64(t.AccountID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.Host),
		orm.ToSQLString(t.HealthStatus),
		orm.ToSQLBool(t.Reporting),
		orm.ToSQLInt64(t.LastReportedAt),
		orm.ToSQLFloat64(t.SummaryCPU),
		orm.ToSQLFloat64(t.SummaryCPUStolen),
		orm.ToSQLFloat64(t.SummaryDiskIo),
		orm.ToSQLFloat64(t.SummaryMemory),
		orm.ToSQLInt64(t.SummaryMemoryUsed),
		orm.ToSQLInt64(t.SummaryMemoryTotal),
		orm.ToSQLFloat64(t.SummaryFullestDisk),
		orm.ToSQLInt64(t.SummaryFullestDiskFree),
	)
}

// DBCreateIgnoreDuplicateTx will upsert the NewrelicServer record in the database using the provided transaction
func (t *NewrelicServer) DBCreateIgnoreDuplicateTx(ctx context.Context, tx *sql.Tx) (sql.Result, error) {
	q := "INSERT INTO `newrelic_server` (`newrelic_server`.`id`,`newrelic_server`.`checksum`,`newrelic_server`.`customer_id`,`newrelic_server`.`ext_id`,`newrelic_server`.`account_id`,`newrelic_server`.`name`,`newrelic_server`.`host`,`newrelic_server`.`health_status`,`newrelic_server`.`reporting`,`newrelic_server`.`last_reported_at`,`newrelic_server`.`summary_cpu`,`newrelic_server`.`summary_cpu_stolen`,`newrelic_server`.`summary_disk_io`,`newrelic_server`.`summary_memory`,`newrelic_server`.`summary_memory_used`,`newrelic_server`.`summary_memory_total`,`newrelic_server`.`summary_fullest_disk`,`newrelic_server`.`summary_fullest_disk_free`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `id` = `id`"
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLInt64(t.ExtID),
		orm.ToSQLInt64(t.AccountID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.Host),
		orm.ToSQLString(t.HealthStatus),
		orm.ToSQLBool(t.Reporting),
		orm.ToSQLInt64(t.LastReportedAt),
		orm.ToSQLFloat64(t.SummaryCPU),
		orm.ToSQLFloat64(t.SummaryCPUStolen),
		orm.ToSQLFloat64(t.SummaryDiskIo),
		orm.ToSQLFloat64(t.SummaryMemory),
		orm.ToSQLInt64(t.SummaryMemoryUsed),
		orm.ToSQLInt64(t.SummaryMemoryTotal),
		orm.ToSQLFloat64(t.SummaryFullestDisk),
		orm.ToSQLInt64(t.SummaryFullestDiskFree),
	)
}

// DeleteAllNewrelicServers deletes all NewrelicServer records in the database with optional filters
func DeleteAllNewrelicServers(ctx context.Context, db *sql.DB, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(NewrelicServerTableName),
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

// DeleteAllNewrelicServersTx deletes all NewrelicServer records in the database with optional filters using the provided transaction
func DeleteAllNewrelicServersTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) error {
	params := []interface{}{
		orm.Table(NewrelicServerTableName),
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

// DBDelete will delete this NewrelicServer record in the database
func (t *NewrelicServer) DBDelete(ctx context.Context, db *sql.DB) (bool, error) {
	q := "DELETE FROM `newrelic_server` WHERE `id` = ?"
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

// DBDeleteTx will delete this NewrelicServer record in the database using the provided transaction
func (t *NewrelicServer) DBDeleteTx(ctx context.Context, tx *sql.Tx) (bool, error) {
	q := "DELETE FROM `newrelic_server` WHERE `id` = ?"
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

// DBUpdate will update the NewrelicServer record in the database
func (t *NewrelicServer) DBUpdate(ctx context.Context, db *sql.DB) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `newrelic_server` SET `checksum`=?,`customer_id`=?,`ext_id`=?,`account_id`=?,`name`=?,`host`=?,`health_status`=?,`reporting`=?,`last_reported_at`=?,`summary_cpu`=?,`summary_cpu_stolen`=?,`summary_disk_io`=?,`summary_memory`=?,`summary_memory_used`=?,`summary_memory_total`=?,`summary_fullest_disk`=?,`summary_fullest_disk_free`=? WHERE `id`=?"
	return db.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLInt64(t.ExtID),
		orm.ToSQLInt64(t.AccountID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.Host),
		orm.ToSQLString(t.HealthStatus),
		orm.ToSQLBool(t.Reporting),
		orm.ToSQLInt64(t.LastReportedAt),
		orm.ToSQLFloat64(t.SummaryCPU),
		orm.ToSQLFloat64(t.SummaryCPUStolen),
		orm.ToSQLFloat64(t.SummaryDiskIo),
		orm.ToSQLFloat64(t.SummaryMemory),
		orm.ToSQLInt64(t.SummaryMemoryUsed),
		orm.ToSQLInt64(t.SummaryMemoryTotal),
		orm.ToSQLFloat64(t.SummaryFullestDisk),
		orm.ToSQLInt64(t.SummaryFullestDiskFree),
		orm.ToSQLString(t.ID),
	)
}

// DBUpdateTx will update the NewrelicServer record in the database using the provided transaction
func (t *NewrelicServer) DBUpdateTx(ctx context.Context, tx *sql.Tx) (sql.Result, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return nil, nil
	}
	t.Checksum = &checksum
	q := "UPDATE `newrelic_server` SET `checksum`=?,`customer_id`=?,`ext_id`=?,`account_id`=?,`name`=?,`host`=?,`health_status`=?,`reporting`=?,`last_reported_at`=?,`summary_cpu`=?,`summary_cpu_stolen`=?,`summary_disk_io`=?,`summary_memory`=?,`summary_memory_used`=?,`summary_memory_total`=?,`summary_fullest_disk`=?,`summary_fullest_disk_free`=? WHERE `id`=?"
	return tx.ExecContext(ctx, q,
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLInt64(t.ExtID),
		orm.ToSQLInt64(t.AccountID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.Host),
		orm.ToSQLString(t.HealthStatus),
		orm.ToSQLBool(t.Reporting),
		orm.ToSQLInt64(t.LastReportedAt),
		orm.ToSQLFloat64(t.SummaryCPU),
		orm.ToSQLFloat64(t.SummaryCPUStolen),
		orm.ToSQLFloat64(t.SummaryDiskIo),
		orm.ToSQLFloat64(t.SummaryMemory),
		orm.ToSQLInt64(t.SummaryMemoryUsed),
		orm.ToSQLInt64(t.SummaryMemoryTotal),
		orm.ToSQLFloat64(t.SummaryFullestDisk),
		orm.ToSQLInt64(t.SummaryFullestDiskFree),
		orm.ToSQLString(t.ID),
	)
}

// DBUpsert will upsert the NewrelicServer record in the database
func (t *NewrelicServer) DBUpsert(ctx context.Context, db *sql.DB, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `newrelic_server` (`newrelic_server`.`id`,`newrelic_server`.`checksum`,`newrelic_server`.`customer_id`,`newrelic_server`.`ext_id`,`newrelic_server`.`account_id`,`newrelic_server`.`name`,`newrelic_server`.`host`,`newrelic_server`.`health_status`,`newrelic_server`.`reporting`,`newrelic_server`.`last_reported_at`,`newrelic_server`.`summary_cpu`,`newrelic_server`.`summary_cpu_stolen`,`newrelic_server`.`summary_disk_io`,`newrelic_server`.`summary_memory`,`newrelic_server`.`summary_memory_used`,`newrelic_server`.`summary_memory_total`,`newrelic_server`.`summary_fullest_disk`,`newrelic_server`.`summary_fullest_disk_free`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `newrelic_server` (`newrelic_server`.`id`,`newrelic_server`.`checksum`,`newrelic_server`.`customer_id`,`newrelic_server`.`ext_id`,`newrelic_server`.`account_id`,`newrelic_server`.`name`,`newrelic_server`.`host`,`newrelic_server`.`health_status`,`newrelic_server`.`reporting`,`newrelic_server`.`last_reported_at`,`newrelic_server`.`summary_cpu`,`newrelic_server`.`summary_cpu_stolen`,`newrelic_server`.`summary_disk_io`,`newrelic_server`.`summary_memory`,`newrelic_server`.`summary_memory_used`,`newrelic_server`.`summary_memory_total`,`newrelic_server`.`summary_fullest_disk`,`newrelic_server`.`summary_fullest_disk_free`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`customer_id`=VALUES(`customer_id`),`ext_id`=VALUES(`ext_id`),`account_id`=VALUES(`account_id`),`name`=VALUES(`name`),`host`=VALUES(`host`),`health_status`=VALUES(`health_status`),`reporting`=VALUES(`reporting`),`last_reported_at`=VALUES(`last_reported_at`),`summary_cpu`=VALUES(`summary_cpu`),`summary_cpu_stolen`=VALUES(`summary_cpu_stolen`),`summary_disk_io`=VALUES(`summary_disk_io`),`summary_memory`=VALUES(`summary_memory`),`summary_memory_used`=VALUES(`summary_memory_used`),`summary_memory_total`=VALUES(`summary_memory_total`),`summary_fullest_disk`=VALUES(`summary_fullest_disk`),`summary_fullest_disk_free`=VALUES(`summary_fullest_disk_free`)"
	}
	r, err := db.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLInt64(t.ExtID),
		orm.ToSQLInt64(t.AccountID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.Host),
		orm.ToSQLString(t.HealthStatus),
		orm.ToSQLBool(t.Reporting),
		orm.ToSQLInt64(t.LastReportedAt),
		orm.ToSQLFloat64(t.SummaryCPU),
		orm.ToSQLFloat64(t.SummaryCPUStolen),
		orm.ToSQLFloat64(t.SummaryDiskIo),
		orm.ToSQLFloat64(t.SummaryMemory),
		orm.ToSQLInt64(t.SummaryMemoryUsed),
		orm.ToSQLInt64(t.SummaryMemoryTotal),
		orm.ToSQLFloat64(t.SummaryFullestDisk),
		orm.ToSQLInt64(t.SummaryFullestDiskFree),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBUpsertTx will upsert the NewrelicServer record in the database using the provided transaction
func (t *NewrelicServer) DBUpsertTx(ctx context.Context, tx *sql.Tx, conditions ...interface{}) (bool, bool, error) {
	checksum := t.CalculateChecksum()
	if t.GetChecksum() == checksum {
		return false, false, nil
	}
	t.Checksum = &checksum
	var q string
	if conditions != nil && len(conditions) > 0 {
		q = "INSERT INTO `newrelic_server` (`newrelic_server`.`id`,`newrelic_server`.`checksum`,`newrelic_server`.`customer_id`,`newrelic_server`.`ext_id`,`newrelic_server`.`account_id`,`newrelic_server`.`name`,`newrelic_server`.`host`,`newrelic_server`.`health_status`,`newrelic_server`.`reporting`,`newrelic_server`.`last_reported_at`,`newrelic_server`.`summary_cpu`,`newrelic_server`.`summary_cpu_stolen`,`newrelic_server`.`summary_disk_io`,`newrelic_server`.`summary_memory`,`newrelic_server`.`summary_memory_used`,`newrelic_server`.`summary_memory_total`,`newrelic_server`.`summary_fullest_disk`,`newrelic_server`.`summary_fullest_disk_free`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE "
		for _, cond := range conditions {
			q = fmt.Sprintf("%s %v ", q, cond)
		}
	} else {
		q = "INSERT INTO `newrelic_server` (`newrelic_server`.`id`,`newrelic_server`.`checksum`,`newrelic_server`.`customer_id`,`newrelic_server`.`ext_id`,`newrelic_server`.`account_id`,`newrelic_server`.`name`,`newrelic_server`.`host`,`newrelic_server`.`health_status`,`newrelic_server`.`reporting`,`newrelic_server`.`last_reported_at`,`newrelic_server`.`summary_cpu`,`newrelic_server`.`summary_cpu_stolen`,`newrelic_server`.`summary_disk_io`,`newrelic_server`.`summary_memory`,`newrelic_server`.`summary_memory_used`,`newrelic_server`.`summary_memory_total`,`newrelic_server`.`summary_fullest_disk`,`newrelic_server`.`summary_fullest_disk_free`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `checksum`=VALUES(`checksum`),`customer_id`=VALUES(`customer_id`),`ext_id`=VALUES(`ext_id`),`account_id`=VALUES(`account_id`),`name`=VALUES(`name`),`host`=VALUES(`host`),`health_status`=VALUES(`health_status`),`reporting`=VALUES(`reporting`),`last_reported_at`=VALUES(`last_reported_at`),`summary_cpu`=VALUES(`summary_cpu`),`summary_cpu_stolen`=VALUES(`summary_cpu_stolen`),`summary_disk_io`=VALUES(`summary_disk_io`),`summary_memory`=VALUES(`summary_memory`),`summary_memory_used`=VALUES(`summary_memory_used`),`summary_memory_total`=VALUES(`summary_memory_total`),`summary_fullest_disk`=VALUES(`summary_fullest_disk`),`summary_fullest_disk_free`=VALUES(`summary_fullest_disk_free`)"
	}
	r, err := tx.ExecContext(ctx, q,
		orm.ToSQLString(t.ID),
		orm.ToSQLString(t.Checksum),
		orm.ToSQLString(t.CustomerID),
		orm.ToSQLInt64(t.ExtID),
		orm.ToSQLInt64(t.AccountID),
		orm.ToSQLString(t.Name),
		orm.ToSQLString(t.Host),
		orm.ToSQLString(t.HealthStatus),
		orm.ToSQLBool(t.Reporting),
		orm.ToSQLInt64(t.LastReportedAt),
		orm.ToSQLFloat64(t.SummaryCPU),
		orm.ToSQLFloat64(t.SummaryCPUStolen),
		orm.ToSQLFloat64(t.SummaryDiskIo),
		orm.ToSQLFloat64(t.SummaryMemory),
		orm.ToSQLInt64(t.SummaryMemoryUsed),
		orm.ToSQLInt64(t.SummaryMemoryTotal),
		orm.ToSQLFloat64(t.SummaryFullestDisk),
		orm.ToSQLInt64(t.SummaryFullestDiskFree),
	)
	if err != nil {
		return false, false, err
	}
	c, _ := r.RowsAffected()
	return c > 0, c == 0, nil
}

// DBFindOne will find a NewrelicServer record in the database with the primary key
func (t *NewrelicServer) DBFindOne(ctx context.Context, db *sql.DB, value string) (bool, error) {
	q := "SELECT `newrelic_server`.`id`,`newrelic_server`.`checksum`,`newrelic_server`.`customer_id`,`newrelic_server`.`ext_id`,`newrelic_server`.`account_id`,`newrelic_server`.`name`,`newrelic_server`.`host`,`newrelic_server`.`health_status`,`newrelic_server`.`reporting`,`newrelic_server`.`last_reported_at`,`newrelic_server`.`summary_cpu`,`newrelic_server`.`summary_cpu_stolen`,`newrelic_server`.`summary_disk_io`,`newrelic_server`.`summary_memory`,`newrelic_server`.`summary_memory_used`,`newrelic_server`.`summary_memory_total`,`newrelic_server`.`summary_fullest_disk`,`newrelic_server`.`summary_fullest_disk_free` FROM `newrelic_server` WHERE `id` = ? LIMIT 1"
	row := db.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CustomerID sql.NullString
	var _ExtID sql.NullInt64
	var _AccountID sql.NullInt64
	var _Name sql.NullString
	var _Host sql.NullString
	var _HealthStatus sql.NullString
	var _Reporting sql.NullBool
	var _LastReportedAt sql.NullInt64
	var _SummaryCPU sql.NullFloat64
	var _SummaryCPUStolen sql.NullFloat64
	var _SummaryDiskIo sql.NullFloat64
	var _SummaryMemory sql.NullFloat64
	var _SummaryMemoryUsed sql.NullInt64
	var _SummaryMemoryTotal sql.NullInt64
	var _SummaryFullestDisk sql.NullFloat64
	var _SummaryFullestDiskFree sql.NullInt64
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_ExtID,
		&_AccountID,
		&_Name,
		&_Host,
		&_HealthStatus,
		&_Reporting,
		&_LastReportedAt,
		&_SummaryCPU,
		&_SummaryCPUStolen,
		&_SummaryDiskIo,
		&_SummaryMemory,
		&_SummaryMemoryUsed,
		&_SummaryMemoryTotal,
		&_SummaryFullestDisk,
		&_SummaryFullestDiskFree,
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
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _ExtID.Valid {
		t.SetExtID(_ExtID.Int64)
	}
	if _AccountID.Valid {
		t.SetAccountID(_AccountID.Int64)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _Host.Valid {
		t.SetHost(_Host.String)
	}
	if _HealthStatus.Valid {
		t.SetHealthStatus(_HealthStatus.String)
	}
	if _Reporting.Valid {
		t.SetReporting(_Reporting.Bool)
	}
	if _LastReportedAt.Valid {
		t.SetLastReportedAt(_LastReportedAt.Int64)
	}
	if _SummaryCPU.Valid {
		t.SetSummaryCPU(_SummaryCPU.Float64)
	}
	if _SummaryCPUStolen.Valid {
		t.SetSummaryCPUStolen(_SummaryCPUStolen.Float64)
	}
	if _SummaryDiskIo.Valid {
		t.SetSummaryDiskIo(_SummaryDiskIo.Float64)
	}
	if _SummaryMemory.Valid {
		t.SetSummaryMemory(_SummaryMemory.Float64)
	}
	if _SummaryMemoryUsed.Valid {
		t.SetSummaryMemoryUsed(_SummaryMemoryUsed.Int64)
	}
	if _SummaryMemoryTotal.Valid {
		t.SetSummaryMemoryTotal(_SummaryMemoryTotal.Int64)
	}
	if _SummaryFullestDisk.Valid {
		t.SetSummaryFullestDisk(_SummaryFullestDisk.Float64)
	}
	if _SummaryFullestDiskFree.Valid {
		t.SetSummaryFullestDiskFree(_SummaryFullestDiskFree.Int64)
	}
	return true, nil
}

// DBFindOneTx will find a NewrelicServer record in the database with the primary key using the provided transaction
func (t *NewrelicServer) DBFindOneTx(ctx context.Context, tx *sql.Tx, value string) (bool, error) {
	q := "SELECT `newrelic_server`.`id`,`newrelic_server`.`checksum`,`newrelic_server`.`customer_id`,`newrelic_server`.`ext_id`,`newrelic_server`.`account_id`,`newrelic_server`.`name`,`newrelic_server`.`host`,`newrelic_server`.`health_status`,`newrelic_server`.`reporting`,`newrelic_server`.`last_reported_at`,`newrelic_server`.`summary_cpu`,`newrelic_server`.`summary_cpu_stolen`,`newrelic_server`.`summary_disk_io`,`newrelic_server`.`summary_memory`,`newrelic_server`.`summary_memory_used`,`newrelic_server`.`summary_memory_total`,`newrelic_server`.`summary_fullest_disk`,`newrelic_server`.`summary_fullest_disk_free` FROM `newrelic_server` WHERE `id` = ? LIMIT 1"
	row := tx.QueryRowContext(ctx, q, orm.ToSQLString(value))
	var _ID sql.NullString
	var _Checksum sql.NullString
	var _CustomerID sql.NullString
	var _ExtID sql.NullInt64
	var _AccountID sql.NullInt64
	var _Name sql.NullString
	var _Host sql.NullString
	var _HealthStatus sql.NullString
	var _Reporting sql.NullBool
	var _LastReportedAt sql.NullInt64
	var _SummaryCPU sql.NullFloat64
	var _SummaryCPUStolen sql.NullFloat64
	var _SummaryDiskIo sql.NullFloat64
	var _SummaryMemory sql.NullFloat64
	var _SummaryMemoryUsed sql.NullInt64
	var _SummaryMemoryTotal sql.NullInt64
	var _SummaryFullestDisk sql.NullFloat64
	var _SummaryFullestDiskFree sql.NullInt64
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_ExtID,
		&_AccountID,
		&_Name,
		&_Host,
		&_HealthStatus,
		&_Reporting,
		&_LastReportedAt,
		&_SummaryCPU,
		&_SummaryCPUStolen,
		&_SummaryDiskIo,
		&_SummaryMemory,
		&_SummaryMemoryUsed,
		&_SummaryMemoryTotal,
		&_SummaryFullestDisk,
		&_SummaryFullestDiskFree,
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
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _ExtID.Valid {
		t.SetExtID(_ExtID.Int64)
	}
	if _AccountID.Valid {
		t.SetAccountID(_AccountID.Int64)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _Host.Valid {
		t.SetHost(_Host.String)
	}
	if _HealthStatus.Valid {
		t.SetHealthStatus(_HealthStatus.String)
	}
	if _Reporting.Valid {
		t.SetReporting(_Reporting.Bool)
	}
	if _LastReportedAt.Valid {
		t.SetLastReportedAt(_LastReportedAt.Int64)
	}
	if _SummaryCPU.Valid {
		t.SetSummaryCPU(_SummaryCPU.Float64)
	}
	if _SummaryCPUStolen.Valid {
		t.SetSummaryCPUStolen(_SummaryCPUStolen.Float64)
	}
	if _SummaryDiskIo.Valid {
		t.SetSummaryDiskIo(_SummaryDiskIo.Float64)
	}
	if _SummaryMemory.Valid {
		t.SetSummaryMemory(_SummaryMemory.Float64)
	}
	if _SummaryMemoryUsed.Valid {
		t.SetSummaryMemoryUsed(_SummaryMemoryUsed.Int64)
	}
	if _SummaryMemoryTotal.Valid {
		t.SetSummaryMemoryTotal(_SummaryMemoryTotal.Int64)
	}
	if _SummaryFullestDisk.Valid {
		t.SetSummaryFullestDisk(_SummaryFullestDisk.Float64)
	}
	if _SummaryFullestDiskFree.Valid {
		t.SetSummaryFullestDiskFree(_SummaryFullestDiskFree.Int64)
	}
	return true, nil
}

// FindNewrelicServers will find a NewrelicServer record in the database with the provided parameters
func FindNewrelicServers(ctx context.Context, db *sql.DB, _params ...interface{}) ([]*NewrelicServer, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("customer_id"),
		orm.Column("ext_id"),
		orm.Column("account_id"),
		orm.Column("name"),
		orm.Column("host"),
		orm.Column("health_status"),
		orm.Column("reporting"),
		orm.Column("last_reported_at"),
		orm.Column("summary_cpu"),
		orm.Column("summary_cpu_stolen"),
		orm.Column("summary_disk_io"),
		orm.Column("summary_memory"),
		orm.Column("summary_memory_used"),
		orm.Column("summary_memory_total"),
		orm.Column("summary_fullest_disk"),
		orm.Column("summary_fullest_disk_free"),
		orm.Table(NewrelicServerTableName),
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
	results := make([]*NewrelicServer, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _ExtID sql.NullInt64
		var _AccountID sql.NullInt64
		var _Name sql.NullString
		var _Host sql.NullString
		var _HealthStatus sql.NullString
		var _Reporting sql.NullBool
		var _LastReportedAt sql.NullInt64
		var _SummaryCPU sql.NullFloat64
		var _SummaryCPUStolen sql.NullFloat64
		var _SummaryDiskIo sql.NullFloat64
		var _SummaryMemory sql.NullFloat64
		var _SummaryMemoryUsed sql.NullInt64
		var _SummaryMemoryTotal sql.NullInt64
		var _SummaryFullestDisk sql.NullFloat64
		var _SummaryFullestDiskFree sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_ExtID,
			&_AccountID,
			&_Name,
			&_Host,
			&_HealthStatus,
			&_Reporting,
			&_LastReportedAt,
			&_SummaryCPU,
			&_SummaryCPUStolen,
			&_SummaryDiskIo,
			&_SummaryMemory,
			&_SummaryMemoryUsed,
			&_SummaryMemoryTotal,
			&_SummaryFullestDisk,
			&_SummaryFullestDiskFree,
		)
		if err != nil {
			return nil, err
		}
		t := &NewrelicServer{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _ExtID.Valid {
			t.SetExtID(_ExtID.Int64)
		}
		if _AccountID.Valid {
			t.SetAccountID(_AccountID.Int64)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _Host.Valid {
			t.SetHost(_Host.String)
		}
		if _HealthStatus.Valid {
			t.SetHealthStatus(_HealthStatus.String)
		}
		if _Reporting.Valid {
			t.SetReporting(_Reporting.Bool)
		}
		if _LastReportedAt.Valid {
			t.SetLastReportedAt(_LastReportedAt.Int64)
		}
		if _SummaryCPU.Valid {
			t.SetSummaryCPU(_SummaryCPU.Float64)
		}
		if _SummaryCPUStolen.Valid {
			t.SetSummaryCPUStolen(_SummaryCPUStolen.Float64)
		}
		if _SummaryDiskIo.Valid {
			t.SetSummaryDiskIo(_SummaryDiskIo.Float64)
		}
		if _SummaryMemory.Valid {
			t.SetSummaryMemory(_SummaryMemory.Float64)
		}
		if _SummaryMemoryUsed.Valid {
			t.SetSummaryMemoryUsed(_SummaryMemoryUsed.Int64)
		}
		if _SummaryMemoryTotal.Valid {
			t.SetSummaryMemoryTotal(_SummaryMemoryTotal.Int64)
		}
		if _SummaryFullestDisk.Valid {
			t.SetSummaryFullestDisk(_SummaryFullestDisk.Float64)
		}
		if _SummaryFullestDiskFree.Valid {
			t.SetSummaryFullestDiskFree(_SummaryFullestDiskFree.Int64)
		}
		results = append(results, t)
	}
	return results, nil
}

// FindNewrelicServersTx will find a NewrelicServer record in the database with the provided parameters using the provided transaction
func FindNewrelicServersTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) ([]*NewrelicServer, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("customer_id"),
		orm.Column("ext_id"),
		orm.Column("account_id"),
		orm.Column("name"),
		orm.Column("host"),
		orm.Column("health_status"),
		orm.Column("reporting"),
		orm.Column("last_reported_at"),
		orm.Column("summary_cpu"),
		orm.Column("summary_cpu_stolen"),
		orm.Column("summary_disk_io"),
		orm.Column("summary_memory"),
		orm.Column("summary_memory_used"),
		orm.Column("summary_memory_total"),
		orm.Column("summary_fullest_disk"),
		orm.Column("summary_fullest_disk_free"),
		orm.Table(NewrelicServerTableName),
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
	results := make([]*NewrelicServer, 0)
	for rows.Next() {
		var _ID sql.NullString
		var _Checksum sql.NullString
		var _CustomerID sql.NullString
		var _ExtID sql.NullInt64
		var _AccountID sql.NullInt64
		var _Name sql.NullString
		var _Host sql.NullString
		var _HealthStatus sql.NullString
		var _Reporting sql.NullBool
		var _LastReportedAt sql.NullInt64
		var _SummaryCPU sql.NullFloat64
		var _SummaryCPUStolen sql.NullFloat64
		var _SummaryDiskIo sql.NullFloat64
		var _SummaryMemory sql.NullFloat64
		var _SummaryMemoryUsed sql.NullInt64
		var _SummaryMemoryTotal sql.NullInt64
		var _SummaryFullestDisk sql.NullFloat64
		var _SummaryFullestDiskFree sql.NullInt64
		err := rows.Scan(
			&_ID,
			&_Checksum,
			&_CustomerID,
			&_ExtID,
			&_AccountID,
			&_Name,
			&_Host,
			&_HealthStatus,
			&_Reporting,
			&_LastReportedAt,
			&_SummaryCPU,
			&_SummaryCPUStolen,
			&_SummaryDiskIo,
			&_SummaryMemory,
			&_SummaryMemoryUsed,
			&_SummaryMemoryTotal,
			&_SummaryFullestDisk,
			&_SummaryFullestDiskFree,
		)
		if err != nil {
			return nil, err
		}
		t := &NewrelicServer{}
		if _ID.Valid {
			t.SetID(_ID.String)
		}
		if _Checksum.Valid {
			t.SetChecksum(_Checksum.String)
		}
		if _CustomerID.Valid {
			t.SetCustomerID(_CustomerID.String)
		}
		if _ExtID.Valid {
			t.SetExtID(_ExtID.Int64)
		}
		if _AccountID.Valid {
			t.SetAccountID(_AccountID.Int64)
		}
		if _Name.Valid {
			t.SetName(_Name.String)
		}
		if _Host.Valid {
			t.SetHost(_Host.String)
		}
		if _HealthStatus.Valid {
			t.SetHealthStatus(_HealthStatus.String)
		}
		if _Reporting.Valid {
			t.SetReporting(_Reporting.Bool)
		}
		if _LastReportedAt.Valid {
			t.SetLastReportedAt(_LastReportedAt.Int64)
		}
		if _SummaryCPU.Valid {
			t.SetSummaryCPU(_SummaryCPU.Float64)
		}
		if _SummaryCPUStolen.Valid {
			t.SetSummaryCPUStolen(_SummaryCPUStolen.Float64)
		}
		if _SummaryDiskIo.Valid {
			t.SetSummaryDiskIo(_SummaryDiskIo.Float64)
		}
		if _SummaryMemory.Valid {
			t.SetSummaryMemory(_SummaryMemory.Float64)
		}
		if _SummaryMemoryUsed.Valid {
			t.SetSummaryMemoryUsed(_SummaryMemoryUsed.Int64)
		}
		if _SummaryMemoryTotal.Valid {
			t.SetSummaryMemoryTotal(_SummaryMemoryTotal.Int64)
		}
		if _SummaryFullestDisk.Valid {
			t.SetSummaryFullestDisk(_SummaryFullestDisk.Float64)
		}
		if _SummaryFullestDiskFree.Valid {
			t.SetSummaryFullestDiskFree(_SummaryFullestDiskFree.Int64)
		}
		results = append(results, t)
	}
	return results, nil
}

// DBFind will find a NewrelicServer record in the database with the provided parameters
func (t *NewrelicServer) DBFind(ctx context.Context, db *sql.DB, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("customer_id"),
		orm.Column("ext_id"),
		orm.Column("account_id"),
		orm.Column("name"),
		orm.Column("host"),
		orm.Column("health_status"),
		orm.Column("reporting"),
		orm.Column("last_reported_at"),
		orm.Column("summary_cpu"),
		orm.Column("summary_cpu_stolen"),
		orm.Column("summary_disk_io"),
		orm.Column("summary_memory"),
		orm.Column("summary_memory_used"),
		orm.Column("summary_memory_total"),
		orm.Column("summary_fullest_disk"),
		orm.Column("summary_fullest_disk_free"),
		orm.Table(NewrelicServerTableName),
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
	var _CustomerID sql.NullString
	var _ExtID sql.NullInt64
	var _AccountID sql.NullInt64
	var _Name sql.NullString
	var _Host sql.NullString
	var _HealthStatus sql.NullString
	var _Reporting sql.NullBool
	var _LastReportedAt sql.NullInt64
	var _SummaryCPU sql.NullFloat64
	var _SummaryCPUStolen sql.NullFloat64
	var _SummaryDiskIo sql.NullFloat64
	var _SummaryMemory sql.NullFloat64
	var _SummaryMemoryUsed sql.NullInt64
	var _SummaryMemoryTotal sql.NullInt64
	var _SummaryFullestDisk sql.NullFloat64
	var _SummaryFullestDiskFree sql.NullInt64
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_ExtID,
		&_AccountID,
		&_Name,
		&_Host,
		&_HealthStatus,
		&_Reporting,
		&_LastReportedAt,
		&_SummaryCPU,
		&_SummaryCPUStolen,
		&_SummaryDiskIo,
		&_SummaryMemory,
		&_SummaryMemoryUsed,
		&_SummaryMemoryTotal,
		&_SummaryFullestDisk,
		&_SummaryFullestDiskFree,
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
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _ExtID.Valid {
		t.SetExtID(_ExtID.Int64)
	}
	if _AccountID.Valid {
		t.SetAccountID(_AccountID.Int64)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _Host.Valid {
		t.SetHost(_Host.String)
	}
	if _HealthStatus.Valid {
		t.SetHealthStatus(_HealthStatus.String)
	}
	if _Reporting.Valid {
		t.SetReporting(_Reporting.Bool)
	}
	if _LastReportedAt.Valid {
		t.SetLastReportedAt(_LastReportedAt.Int64)
	}
	if _SummaryCPU.Valid {
		t.SetSummaryCPU(_SummaryCPU.Float64)
	}
	if _SummaryCPUStolen.Valid {
		t.SetSummaryCPUStolen(_SummaryCPUStolen.Float64)
	}
	if _SummaryDiskIo.Valid {
		t.SetSummaryDiskIo(_SummaryDiskIo.Float64)
	}
	if _SummaryMemory.Valid {
		t.SetSummaryMemory(_SummaryMemory.Float64)
	}
	if _SummaryMemoryUsed.Valid {
		t.SetSummaryMemoryUsed(_SummaryMemoryUsed.Int64)
	}
	if _SummaryMemoryTotal.Valid {
		t.SetSummaryMemoryTotal(_SummaryMemoryTotal.Int64)
	}
	if _SummaryFullestDisk.Valid {
		t.SetSummaryFullestDisk(_SummaryFullestDisk.Float64)
	}
	if _SummaryFullestDiskFree.Valid {
		t.SetSummaryFullestDiskFree(_SummaryFullestDiskFree.Int64)
	}
	return true, nil
}

// DBFindTx will find a NewrelicServer record in the database with the provided parameters using the provided transaction
func (t *NewrelicServer) DBFindTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) (bool, error) {
	params := []interface{}{
		orm.Column("id"),
		orm.Column("checksum"),
		orm.Column("customer_id"),
		orm.Column("ext_id"),
		orm.Column("account_id"),
		orm.Column("name"),
		orm.Column("host"),
		orm.Column("health_status"),
		orm.Column("reporting"),
		orm.Column("last_reported_at"),
		orm.Column("summary_cpu"),
		orm.Column("summary_cpu_stolen"),
		orm.Column("summary_disk_io"),
		orm.Column("summary_memory"),
		orm.Column("summary_memory_used"),
		orm.Column("summary_memory_total"),
		orm.Column("summary_fullest_disk"),
		orm.Column("summary_fullest_disk_free"),
		orm.Table(NewrelicServerTableName),
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
	var _CustomerID sql.NullString
	var _ExtID sql.NullInt64
	var _AccountID sql.NullInt64
	var _Name sql.NullString
	var _Host sql.NullString
	var _HealthStatus sql.NullString
	var _Reporting sql.NullBool
	var _LastReportedAt sql.NullInt64
	var _SummaryCPU sql.NullFloat64
	var _SummaryCPUStolen sql.NullFloat64
	var _SummaryDiskIo sql.NullFloat64
	var _SummaryMemory sql.NullFloat64
	var _SummaryMemoryUsed sql.NullInt64
	var _SummaryMemoryTotal sql.NullInt64
	var _SummaryFullestDisk sql.NullFloat64
	var _SummaryFullestDiskFree sql.NullInt64
	err := row.Scan(
		&_ID,
		&_Checksum,
		&_CustomerID,
		&_ExtID,
		&_AccountID,
		&_Name,
		&_Host,
		&_HealthStatus,
		&_Reporting,
		&_LastReportedAt,
		&_SummaryCPU,
		&_SummaryCPUStolen,
		&_SummaryDiskIo,
		&_SummaryMemory,
		&_SummaryMemoryUsed,
		&_SummaryMemoryTotal,
		&_SummaryFullestDisk,
		&_SummaryFullestDiskFree,
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
	if _CustomerID.Valid {
		t.SetCustomerID(_CustomerID.String)
	}
	if _ExtID.Valid {
		t.SetExtID(_ExtID.Int64)
	}
	if _AccountID.Valid {
		t.SetAccountID(_AccountID.Int64)
	}
	if _Name.Valid {
		t.SetName(_Name.String)
	}
	if _Host.Valid {
		t.SetHost(_Host.String)
	}
	if _HealthStatus.Valid {
		t.SetHealthStatus(_HealthStatus.String)
	}
	if _Reporting.Valid {
		t.SetReporting(_Reporting.Bool)
	}
	if _LastReportedAt.Valid {
		t.SetLastReportedAt(_LastReportedAt.Int64)
	}
	if _SummaryCPU.Valid {
		t.SetSummaryCPU(_SummaryCPU.Float64)
	}
	if _SummaryCPUStolen.Valid {
		t.SetSummaryCPUStolen(_SummaryCPUStolen.Float64)
	}
	if _SummaryDiskIo.Valid {
		t.SetSummaryDiskIo(_SummaryDiskIo.Float64)
	}
	if _SummaryMemory.Valid {
		t.SetSummaryMemory(_SummaryMemory.Float64)
	}
	if _SummaryMemoryUsed.Valid {
		t.SetSummaryMemoryUsed(_SummaryMemoryUsed.Int64)
	}
	if _SummaryMemoryTotal.Valid {
		t.SetSummaryMemoryTotal(_SummaryMemoryTotal.Int64)
	}
	if _SummaryFullestDisk.Valid {
		t.SetSummaryFullestDisk(_SummaryFullestDisk.Float64)
	}
	if _SummaryFullestDiskFree.Valid {
		t.SetSummaryFullestDiskFree(_SummaryFullestDiskFree.Int64)
	}
	return true, nil
}

// CountNewrelicServers will find the count of NewrelicServer records in the database
func CountNewrelicServers(ctx context.Context, db *sql.DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(NewrelicServerTableName),
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

// CountNewrelicServersTx will find the count of NewrelicServer records in the database using the provided transaction
func CountNewrelicServersTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.Count("*"),
		orm.Table(NewrelicServerTableName),
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

// DBCount will find the count of NewrelicServer records in the database
func (t *NewrelicServer) DBCount(ctx context.Context, db *sql.DB, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(NewrelicServerTableName),
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

// DBCountTx will find the count of NewrelicServer records in the database using the provided transaction
func (t *NewrelicServer) DBCountTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) (int64, error) {
	params := []interface{}{
		orm.CountAlias("*", "count"),
		orm.Table(NewrelicServerTableName),
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

// DBExists will return true if the NewrelicServer record exists in the database
func (t *NewrelicServer) DBExists(ctx context.Context, db *sql.DB) (bool, error) {
	q := "SELECT `id` FROM `newrelic_server` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := db.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// DBExistsTx will return true if the NewrelicServer record exists in the database using the provided transaction
func (t *NewrelicServer) DBExistsTx(ctx context.Context, tx *sql.Tx) (bool, error) {
	q := "SELECT `id` FROM `newrelic_server` WHERE `id` = ? LIMIT 1"
	var _ID sql.NullString
	err := tx.QueryRowContext(ctx, q, orm.ToSQLString(t.ID)).Scan(&_ID)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return _ID.Valid, nil
}

// PrimaryKeyColumn returns the column name for the primary key
func (t *NewrelicServer) PrimaryKeyColumn() string {
	return NewrelicServerColumnID
}

// PrimaryKeyColumnType returns the primary key column Go type as a string
func (t *NewrelicServer) PrimaryKeyColumnType() string {
	return "string"
}

// PrimaryKey returns the primary key column value
func (t *NewrelicServer) PrimaryKey() interface{} {
	return t.ID
}
