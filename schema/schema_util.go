// GENERATED CODE. DO NOT EDIT
package schema

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"encoding/binary"
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/jhaynie/go-gator/orm"
)

// NullTime taken from
// Go MySQL Driver - A MySQL-Driver for Go's database/sql package
//
// Copyright 2012 The Go-MySQL-Driver Authors. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.

// NullTime represents a time.Time that may be NULL.
// NullTime implements the Scanner interface so
// it can be used as a scan destination:
//
//  var nt NullTime
//  err := db.QueryRow("SELECT time FROM foo WHERE id=?", id).Scan(&nt)
//  ...
//  if nt.Valid {
//     // use nt.Time
//  } else {
//     // NULL value
//  }
//

const timeFormat = "2006-01-02 15:04:05.999999"

// This NullTime implementation is not driver-specific
type NullTime struct {
	Time  time.Time
	Valid bool // Valid is true if Time is not NULL
}

// Scan implements the Scanner interface.
// The value type must be time.Time or string / []byte (formatted time-string),
// otherwise Scan fails.
func (nt *NullTime) Scan(value interface{}) (err error) {
	if value == nil {
		nt.Time, nt.Valid = time.Time{}, false
		return
	}

	switch v := value.(type) {
	case time.Time:
		nt.Time, nt.Valid = v, true
		return
	case []byte:
		nt.Time, err = parseDateTime(string(v), time.UTC)
		nt.Valid = (err == nil)
		return
	case string:
		nt.Time, err = parseDateTime(v, time.UTC)
		nt.Valid = (err == nil)
		return
	}

	nt.Valid = false
	return fmt.Errorf("Can't convert %T to time.Time", value)
}

// Value implements the driver Valuer interface.
func (nt NullTime) Value() (driver.Value, error) {
	if !nt.Valid {
		return nil, nil
	}
	return nt.Time, nil
}

func parseDateTime(str string, loc *time.Location) (t time.Time, err error) {
	base := "0000-00-00 00:00:00.0000000"
	switch len(str) {
	case 10, 19, 21, 22, 23, 24, 25, 26: // up to "YYYY-MM-DD HH:MM:SS.MMMMMM"
		if str == base[:len(str)] {
			return
		}
		t, err = time.Parse(timeFormat[:len(str)], str)
	default:
		err = fmt.Errorf("invalid time string: %s", str)
		return
	}

	// Adjust location
	if err == nil && loc != time.UTC {
		y, mo, d := t.Date()
		h, mi, s := t.Clock()
		t, err = time.Date(y, mo, d, h, mi, s, t.Nanosecond(), loc), nil
	}

	return
}

func parseBinaryDateTime(num uint64, data []byte, loc *time.Location) (driver.Value, error) {
	switch num {
	case 0:
		return time.Time{}, nil
	case 4:
		return time.Date(
			int(binary.LittleEndian.Uint16(data[:2])), // year
			time.Month(data[2]),                       // month
			int(data[3]),                              // day
			0, 0, 0, 0,
			loc,
		), nil
	case 7:
		return time.Date(
			int(binary.LittleEndian.Uint16(data[:2])), // year
			time.Month(data[2]),                       // month
			int(data[3]),                              // day
			int(data[4]),                              // hour
			int(data[5]),                              // minutes
			int(data[6]),                              // seconds
			0,
			loc,
		), nil
	case 11:
		return time.Date(
			int(binary.LittleEndian.Uint16(data[:2])), // year
			time.Month(data[2]),                       // month
			int(data[3]),                              // day
			int(data[4]),                              // hour
			int(data[5]),                              // minutes
			int(data[6]),                              // seconds
			int(binary.LittleEndian.Uint32(data[7:11]))*1000, // nanoseconds
			loc,
		), nil
	}
	return nil, fmt.Errorf("invalid DATETIME packet length %d", num)
}

func toCSVBool(v bool) string {
	if v {
		return "1"
	}
	return "0"
}

func toCSVDate(t *timestamp.Timestamp) string {
	tv, err := ptypes.Timestamp(t)
	if err != nil {
		return "NULL"
	}
	return tv.UTC().Format(time.RFC3339)
}

func toCSVString(v interface{}) string {
	if v == nil {
		return "NULL"
	}
	if s, ok := v.(string); ok {
		return s
	}
	if s, ok := v.(*string); ok {
		if s == nil || *s == "" {
			return "NULL"
		}
		return *s
	}
	if i, ok := v.(*int); ok {
		if i == nil {
			return "NULL"
		}
		return fmt.Sprintf("%d", *i)
	}
	if i, ok := v.(*int32); ok {
		if i == nil {
			return "NULL"
		}
		return fmt.Sprintf("%d", *i)
	}
	if i, ok := v.(*int64); ok {
		if i == nil {
			return "NULL"
		}
		return fmt.Sprintf("%d", *i)
	}
	if i, ok := v.(*float32); ok {
		if i == nil {
			return "NULL"
		}
		return fmt.Sprintf("%f", *i)
	}
	if i, ok := v.(*float64); ok {
		if i == nil {
			return "NULL"
		}
		return fmt.Sprintf("%f", *i)
	}
	if i, ok := v.(*bool); ok {
		if i == nil {
			return "NULL"
		}
		return toCSVBool(*i)
	}
	if i, ok := v.(bool); ok {
		return toCSVBool(i)
	}
	return fmt.Sprintf("%v", v)
}

func fromStringPointer(v string) *string {
	if v == "" || v == "NULL" {
		return nil
	}
	return &v
}

func fromCSVBool(v string) bool {
	if v == "1" {
		return true
	}
	return false
}

func fromCSVBoolPointer(v string) *bool {
	if v == "" || v == "NULL" {
		return nil
	}
	b := fromCSVBool(v)
	return &b
}

func fromCSVInt32(v string) int32 {
	if v == "" {
		return int32(0)
	}
	i, _ := strconv.ParseInt(v, 10, 32)
	return int32(i)
}

func fromCSVInt32Pointer(v string) *int32 {
	if v == "" || v == "NULL" {
		return nil
	}
	i := fromCSVInt32(v)
	return &i
}

func fromCSVInt64(v string) int64 {
	if v == "" {
		return int64(0)
	}
	i, _ := strconv.ParseInt(v, 10, 64)
	return int64(i)
}

func fromCSVInt64Pointer(v string) *int64 {
	if v == "" || v == "NULL" {
		return nil
	}
	i := fromCSVInt64(v)
	return &i
}

func fromCSVUint32(v string) uint32 {
	if v == "" {
		return uint32(0)
	}
	i, _ := strconv.ParseUint(v, 10, 32)
	return uint32(i)
}

func fromCSVUint64(v string) uint64 {
	if v == "" {
		return uint64(0)
	}
	i, _ := strconv.ParseUint(v, 10, 64)
	return uint64(i)
}

func fromCSVUint32Pointer(v string) *uint32 {
	if v == "" || v == "NULL" {
		return nil
	}
	i := fromCSVUint32(v)
	return &i
}

func fromCSVUint64Pointer(v string) *uint64 {
	if v == "" || v == "NULL" {
		return nil
	}
	i := fromCSVUint64(v)
	return &i
}

func fromCSVFloat32(v string) float32 {
	if v == "" {
		return float32(0)
	}
	f, _ := strconv.ParseFloat(v, 32)
	return float32(f)
}

func fromCSVFloat64(v string) float64 {
	if v == "" {
		return float64(0)
	}
	f, _ := strconv.ParseFloat(v, 64)
	return f
}

func fromCSVFloat32Pointer(v string) *float32 {
	if v == "" || v == "NULL" {
		return nil
	}
	f := fromCSVFloat32(v)
	return &f
}

func fromCSVFloat64Pointer(v string) *float64 {
	if v == "" || v == "NULL" {
		return nil
	}
	f := fromCSVFloat64(v)
	return &f
}

func fromCSVDate(v string) *timestamp.Timestamp {
	if v == "" || v == "NULL" {
		return nil
	}
	tv, err := time.Parse("2006-01-02T15:04:05Z", v)
	if err != nil {
		return nil
	}
	ts, err := ptypes.TimestampProto(tv)
	if err != nil {
		return nil
	}
	return ts
}

// Deserializer is a callback which will take a json RawMessage for processing
type Deserializer = orm.Deserializer

// Deserialize will return a function which will Deserialize in a flexible way the JSON in reader
func Deserialize(r io.Reader, dser Deserializer) error {
	return orm.Deserialize(r, dser)
}

// Model is an interface for describing a DB model object
type Model interface {

	// TableName is the SQL name of the table
	TableName() string

	DBCreate(ctx context.Context, db *sql.DB) (sql.Result, error)
	DBCreateTx(ctx context.Context, tx *sql.Tx) (sql.Result, error)

	DBUpsert(ctx context.Context, db *sql.DB, conditions ...interface{}) (bool, bool, error)
	DBUpsertTx(ctx context.Context, tx *sql.Tx, conditions ...interface{}) (bool, bool, error)

	DBUpdate(ctx context.Context, db *sql.DB) (sql.Result, error)
	DBUpdateTx(ctx context.Context, tx *sql.Tx) (sql.Result, error)
}

// ModelWithPrimaryKey is an interface for describing a DB model object that has a primary key
type ModelWithPrimaryKey interface {
	PrimaryKeyColumn() string
	PrimaryKeyColumnType() string
	PrimaryKey() interface{}

	DBDelete(ctx context.Context, db *sql.DB) (bool, error)
	DBDeleteTx(ctx context.Context, tx *sql.Tx) (bool, error)

	DBExists(ctx context.Context, db *sql.DB) (bool, error)
	DBExistsTx(ctx context.Context, tx *sql.Tx) (bool, error)

	DBFind(ctx context.Context, db *sql.DB, _params ...interface{}) (bool, error)
	DBFindTx(ctx context.Context, tx *sql.Tx, _params ...interface{}) (bool, error)

	DBCount(ctx context.Context, db *sql.DB, _params ...interface{})
	DBCountTx(ctx context.Context, tx *sql.Tx, _params ...interface{})
}

// Checksum is an interface for describing checking the contents of the model for equality
type Checksum interface {

	// CalculateChecksum will return the checksum of the model. this can be used to compare identical objects as a hash identity
	CalculateChecksum() string
}

// CSVWriter is an interface for implementing CSV writer output
type CSVWriter interface {

	// WriteCSV will write the instance to the writer as CSV
	WriteCSV(w *csv.Writer) error
}

// JSONWriter is an interface for implementing JSON writer output
type JSONWriter interface {

	// WriteJSON will write the instance to the writer as JSON
	WriteJSON(w io.Writer, indent ...bool) error
}
