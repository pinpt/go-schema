// GENERATED CODE. DO NOT EDIT

package schema

import (
	"context"
	"testing"

	"github.com/jhaynie/go-gator/orm"
)

func TestCreateNewrelicServerTable(t *testing.T) {
	tx, err := GetDatabase().Begin()
	if err != nil {
		t.Fatal(err)
	}
	err = DBCreateNewrelicServerTableTx(context.Background(), tx)
	if err != nil {
		tx.Rollback()
		t.Fatal(err)
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		t.Fatal(err)
	}
}

func TestCreateNewrelicServerDelete(t *testing.T) {
	r := &NewrelicServer{
		ID:                     "e2e8c0ee40bc31d4",
		Checksum:               nil,
		CustomerID:             "6a966002bdf72a80",
		ExtID:                  int64(64),
		AccountID:              int64(64),
		Name:                   "eb54768a9dcb3a78",
		Host:                   "3f73ff30adc6ed48",
		HealthStatus:           nil,
		Reporting:              true,
		LastReportedAt:         int64(64),
		SummaryCPU:             nil,
		SummaryCPUStolen:       nil,
		SummaryDiskIo:          nil,
		SummaryMemory:          nil,
		SummaryMemoryUsed:      nil,
		SummaryMemoryTotal:     nil,
		SummaryFullestDisk:     nil,
		SummaryFullestDiskFree: nil,
	}
	ctx := context.Background()
	db := GetDatabase()
	DeleteAllNewrelicServers(ctx, db)
	result, err := r.DBCreate(ctx, db)
	if err != nil {
		t.Fatal(err)
	}
	if result == nil {
		t.Fatal("expected result to be non-nil")
	}
	count, err := r.DBCount(ctx, db)
	if err != nil {
		t.Fatal(err)
	}
	if count != 1 {
		t.Fatalf("count should have been 1 but was %d", count)
	}
	exists, err := r.DBExists(ctx, db)
	if err != nil {
		t.Fatal(err)
	}
	if !exists {
		t.Fatal("exists should have been true but was false")
	}
	found, err := FindNewrelicServerByID(ctx, db, r.ID)
	if err != nil {
		t.Fatal(err)
	}
	if found == nil {
		t.Fatal("expected found to be a value but was nil")
	}
	if found.ID != r.ID {
		t.Fatalf("expected found primary key to be %v but was %v", r.ID, found.ID)
	}
	if orm.Stringify(r) != orm.Stringify(found) {
		t.Fatalf("expected r to be found but was different")
	}
	results, err := FindNewrelicServers(ctx, db)
	if err != nil {
		t.Fatal(err)
	}
	if results == nil {
		t.Fatal("expected results to be a value but was nil")
	}
	if len(results) != 1 {
		t.Log(orm.Stringify(results))
		t.Fatalf("expected results length to be 1 but was %d", len(results))
	}
	f, err := r.DBFindOne(ctx, db, r.GetID())
	if err != nil {
		t.Fatal(err)
	}
	if f == false {
		t.Fatal("expected found to be a true but was false")
	}
	a, b, err := r.DBUpsert(ctx, db)
	if err != nil {
		t.Fatal(err)
	}
	if a {
		t.Fatal("expected a to be false but was true")
	}
	if b {
		t.Fatal("expected b to be false but was true")
	}
	r.SetCustomerID("cfbf4dddc542df59")

	r.SetExtID(int64(640))

	r.SetAccountID(int64(640))

	r.SetName("3f461e6ff1408a71")

	r.SetHost("f556add5ef1e0ba0")

	r.SetReporting(false)

	r.SetLastReportedAt(int64(640))

	a, b, err = r.DBUpsert(ctx, db)
	if err != nil {
		t.Fatal(err)
	}
	if !a {
		t.Fatal("expected a to be true but was false")
	}
	if b {
		t.Fatal("expected b to be false but was true")
	}
	_, err = r.DBDelete(ctx, db)
	if err != nil {
		t.Fatal(err)
	}
	count, err = r.DBCount(ctx, db)
	if err != nil {
		t.Fatal(err)
	}
	if count != 0 {
		t.Fatalf("count should have been 0 but was %d", count)
	}
	exists, err = r.DBExists(ctx, db)
	if err != nil {
		t.Fatal(err)
	}
	if exists {
		t.Fatal("exists should have been false but was true")
	}
}

func TestCreateNewrelicServerDeleteTx(t *testing.T) {
	r := &NewrelicServer{
		ID:                     "e2e8c0ee40bc31d4",
		Checksum:               nil,
		CustomerID:             "6a966002bdf72a80",
		ExtID:                  int64(64),
		AccountID:              int64(64),
		Name:                   "eb54768a9dcb3a78",
		Host:                   "3f73ff30adc6ed48",
		HealthStatus:           nil,
		Reporting:              true,
		LastReportedAt:         int64(64),
		SummaryCPU:             nil,
		SummaryCPUStolen:       nil,
		SummaryDiskIo:          nil,
		SummaryMemory:          nil,
		SummaryMemoryUsed:      nil,
		SummaryMemoryTotal:     nil,
		SummaryFullestDisk:     nil,
		SummaryFullestDiskFree: nil,
	}
	ctx := context.Background()
	db := GetDatabase()
	DeleteAllNewrelicServers(ctx, db)
	tx, err := db.Begin()
	if err != nil {
		t.Fatal(err)
	}
	result, err := r.DBCreateTx(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if result == nil {
		t.Fatal("expected result to be non-nil")
	}
	count, err := r.DBCountTx(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 1 {
		t.Fatalf("count should have been 1 but was %d", count)
	}
	exists, err := r.DBExistsTx(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if !exists {
		t.Fatal("exists should have been true but was false")
	}
	found, err := FindNewrelicServerByIDTx(ctx, tx, r.ID)
	if err != nil {
		t.Fatal(err)
	}
	if found == nil {
		t.Fatal("expected found to be a value but was nil")
	}
	if found.ID != r.ID {
		t.Fatalf("expected found primary key to be %v but was %v", r.ID, found.ID)
	}
	if orm.Stringify(r) != orm.Stringify(found) {
		t.Fatalf("expected r to be found but was different")
	}
	results, err := FindNewrelicServersTx(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if results == nil {
		t.Fatal("expected results to be a value but was nil")
	}
	if len(results) != 1 {
		t.Log(orm.Stringify(results))
		t.Fatalf("expected results length to be 1 but was %d", len(results))
	}
	f, err := r.DBFindOneTx(ctx, tx, r.GetID())
	if err != nil {
		t.Fatal(err)
	}
	if f == false {
		t.Fatal("expected found to be a true but was false")
	}
	a, b, err := r.DBUpsertTx(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if a {
		t.Fatal("expected a to be false but was true")
	}
	if b {
		t.Fatal("expected b to be false but was true")
	}
	r.SetCustomerID("cfbf4dddc542df59")

	r.SetExtID(int64(640))

	r.SetAccountID(int64(640))

	r.SetName("3f461e6ff1408a71")

	r.SetHost("f556add5ef1e0ba0")

	r.SetReporting(false)

	r.SetLastReportedAt(int64(640))

	a, b, err = r.DBUpsertTx(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if !a {
		t.Fatal("expected a to be true but was false")
	}
	if b {
		t.Fatal("expected b to be false but was true")
	}
	_, err = r.DBDeleteTx(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	count, err = r.DBCountTx(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 0 {
		t.Fatalf("count should have been 0 but was %d", count)
	}
	exists, err = r.DBExistsTx(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if exists {
		t.Fatal("exists should have been false but was true")
	}
	err = tx.Commit()
	if err != nil {
		t.Fatal(err)
	}
}
