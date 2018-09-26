// GENERATED CODE. DO NOT EDIT

package schema

import (
	"context"
	"testing"

	"github.com/jhaynie/go-gator/orm"
)

func TestCreateJiraIssueTable(t *testing.T) {
	tx, err := GetDatabase().Begin()
	if err != nil {
		t.Fatal(err)
	}
	err = DBCreateJiraIssueTableTx(context.Background(), tx)
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

func TestCreateJiraIssueDelete(t *testing.T) {
	r := &JiraIssue{
		ID:                   "d39aed904c397351",
		Checksum:             nil,
		IssueID:              "c33d90a99282e380",
		ProjectID:            "95637b636f0438fe",
		IssueTypeID:          "c0473b0e91391aba",
		UserID:               nil,
		AssigneeID:           nil,
		PriorityID:           nil,
		StatusID:             "902048355bb93119",
		ResolutionID:         nil,
		FixVersionIds:        nil,
		VersionIds:           nil,
		Environment:          nil,
		ComponentIds:         nil,
		LabelIds:             nil,
		DuedateAt:            nil,
		PlannedStartAt:       nil,
		PlannedEndAt:         nil,
		Key:                  "be0f7c02e5ed9362",
		CustomFieldIds:       nil,
		SprintID:             nil,
		EpicID:               nil,
		ParentID:             nil,
		StrategicParentID:    nil,
		InProgressCount:      int32(32),
		ReopenCount:          int32(32),
		InProgressDuration:   int64(64),
		VerificationDuration: int64(64),
		InProgressStartAt:    nil,
		CustomerID:           "1cd2028d0bda0d12",
		RefID:                "88f68e6d04419dd6",
		UserRefID:            "1d56de159c195555",
		Cost:                 float64(6.4),
	}
	ctx := context.Background()
	db := GetDatabase()
	DeleteAllJiraIssues(ctx, db)
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
	found, err := FindJiraIssueByID(ctx, db, r.ID)
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
	results, err := FindJiraIssues(ctx, db)
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
	r.SetIssueID("5f93dbc727615f1c")

	r.SetProjectID("c80a30b41cb3b08a")

	r.SetIssueTypeID("f455cf0b1546c586")

	r.SetStatusID("1a357edbea0a5c5c")

	r.SetKey("3b7ff9b61809d86c")

	r.SetInProgressCount(int32(320))

	r.SetReopenCount(int32(320))

	r.SetInProgressDuration(int64(640))

	r.SetVerificationDuration(int64(640))

	r.SetCustomerID("1c48e1f5fb3d938f")

	r.SetRefID("01bd771c6f908687")

	r.SetUserRefID("f1047c15105d3575")

	r.SetCost(float64(64.1))

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

func TestCreateJiraIssueDeleteTx(t *testing.T) {
	r := &JiraIssue{
		ID:                   "d39aed904c397351",
		Checksum:             nil,
		IssueID:              "c33d90a99282e380",
		ProjectID:            "95637b636f0438fe",
		IssueTypeID:          "c0473b0e91391aba",
		UserID:               nil,
		AssigneeID:           nil,
		PriorityID:           nil,
		StatusID:             "902048355bb93119",
		ResolutionID:         nil,
		FixVersionIds:        nil,
		VersionIds:           nil,
		Environment:          nil,
		ComponentIds:         nil,
		LabelIds:             nil,
		DuedateAt:            nil,
		PlannedStartAt:       nil,
		PlannedEndAt:         nil,
		Key:                  "be0f7c02e5ed9362",
		CustomFieldIds:       nil,
		SprintID:             nil,
		EpicID:               nil,
		ParentID:             nil,
		StrategicParentID:    nil,
		InProgressCount:      int32(32),
		ReopenCount:          int32(32),
		InProgressDuration:   int64(64),
		VerificationDuration: int64(64),
		InProgressStartAt:    nil,
		CustomerID:           "1cd2028d0bda0d12",
		RefID:                "88f68e6d04419dd6",
		UserRefID:            "1d56de159c195555",
		Cost:                 float64(6.4),
	}
	ctx := context.Background()
	db := GetDatabase()
	DeleteAllJiraIssues(ctx, db)
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
	found, err := FindJiraIssueByIDTx(ctx, tx, r.ID)
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
	results, err := FindJiraIssuesTx(ctx, tx)
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
	r.SetIssueID("5f93dbc727615f1c")

	r.SetProjectID("c80a30b41cb3b08a")

	r.SetIssueTypeID("f455cf0b1546c586")

	r.SetStatusID("1a357edbea0a5c5c")

	r.SetKey("3b7ff9b61809d86c")

	r.SetInProgressCount(int32(320))

	r.SetReopenCount(int32(320))

	r.SetInProgressDuration(int64(640))

	r.SetVerificationDuration(int64(640))

	r.SetCustomerID("1c48e1f5fb3d938f")

	r.SetRefID("01bd771c6f908687")

	r.SetUserRefID("f1047c15105d3575")

	r.SetCost(float64(64.1))

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
