// GENERATED CODE. DO NOT EDIT

package schema

import (
	"context"
	"testing"

	"github.com/jhaynie/go-gator/orm"
)

func TestCreateJiraProjectSprintTable(t *testing.T) {
	tx, err := GetDatabase().Begin()
	if err != nil {
		t.Fatal(err)
	}
	err = DBCreateJiraProjectSprintTableTx(context.Background(), tx)
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

func TestCreateJiraProjectSprintDelete(t *testing.T) {
	r := &JiraProjectSprint{
		ID:                  "b4e2106af80bbc51",
		Checksum:            nil,
		ProjectID:           "1257165aa92f1df0",
		SprintID:            "05601d7cb7f70818",
		Name:                "c0e4be0c39931c1f",
		CustomerID:          "702b480582b3792e",
		CompleteDate:        nil,
		EndDate:             nil,
		StartDate:           nil,
		State:               nil,
		Goal:                nil,
		InitialIssueIds:     nil,
		FinalIssueIds:       nil,
		FinalClosedIssueIds: nil,
		InitialIssueCount:   int32(32),
		FinalIssueCount:     int32(32),
		ClosedIssueCount:    int32(32),
		AddedIssueCount:     int32(32),
		RemovedIssueCount:   int32(32),
		InitialIssuesClosed: int32(32),
	}
	ctx := context.Background()
	db := GetDatabase()
	DeleteAllJiraProjectSprints(ctx, db)
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
	found, err := FindJiraProjectSprintByID(ctx, db, r.ID)
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
	results, err := FindJiraProjectSprints(ctx, db)
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
	r.SetProjectID("a6c31e89f11d939a")

	r.SetSprintID("5a76bba80d3e3d77")

	r.SetName("46ab31c1afd6a943")

	r.SetCustomerID("8bbd1860cb6e7c19")

	r.SetInitialIssueCount(int32(320))

	r.SetFinalIssueCount(int32(320))

	r.SetClosedIssueCount(int32(320))

	r.SetAddedIssueCount(int32(320))

	r.SetRemovedIssueCount(int32(320))

	r.SetInitialIssuesClosed(int32(320))

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

func TestCreateJiraProjectSprintDeleteTx(t *testing.T) {
	r := &JiraProjectSprint{
		ID:                  "b4e2106af80bbc51",
		Checksum:            nil,
		ProjectID:           "1257165aa92f1df0",
		SprintID:            "05601d7cb7f70818",
		Name:                "c0e4be0c39931c1f",
		CustomerID:          "702b480582b3792e",
		CompleteDate:        nil,
		EndDate:             nil,
		StartDate:           nil,
		State:               nil,
		Goal:                nil,
		InitialIssueIds:     nil,
		FinalIssueIds:       nil,
		FinalClosedIssueIds: nil,
		InitialIssueCount:   int32(32),
		FinalIssueCount:     int32(32),
		ClosedIssueCount:    int32(32),
		AddedIssueCount:     int32(32),
		RemovedIssueCount:   int32(32),
		InitialIssuesClosed: int32(32),
	}
	ctx := context.Background()
	db := GetDatabase()
	DeleteAllJiraProjectSprints(ctx, db)
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
	found, err := FindJiraProjectSprintByIDTx(ctx, tx, r.ID)
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
	results, err := FindJiraProjectSprintsTx(ctx, tx)
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
	r.SetProjectID("a6c31e89f11d939a")

	r.SetSprintID("5a76bba80d3e3d77")

	r.SetName("46ab31c1afd6a943")

	r.SetCustomerID("8bbd1860cb6e7c19")

	r.SetInitialIssueCount(int32(320))

	r.SetFinalIssueCount(int32(320))

	r.SetClosedIssueCount(int32(320))

	r.SetAddedIssueCount(int32(320))

	r.SetRemovedIssueCount(int32(320))

	r.SetInitialIssuesClosed(int32(320))

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
