// GENERATED CODE. DO NOT EDIT

package schema

import (
	"context"
	"testing"

	"github.com/jhaynie/go-gator/orm"
)

func TestCreateIssueSummaryTable(t *testing.T) {
	tx, err := GetDatabase().Begin()
	if err != nil {
		t.Fatal(err)
	}
	err = DBCreateIssueSummaryTableTx(context.Background(), tx)
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

func TestCreateIssueSummaryDelete(t *testing.T) {
	r := &IssueSummary{
		ID:                        "e7e39abcec0c73bb",
		Checksum:                  nil,
		IssueID:                   "7b265bb3e2ee9fb3",
		TotalIssues:               int32(32),
		New30Days:                 int32(32),
		TotalClosed:               int32(32),
		Closed30Days:              int32(32),
		EstimatedWorkMonths:       float64(6.4),
		EstimatedWorkMonths30Days: float64(6.4),
		Title:                     "00f4a631851d85fe",
		URL:                       nil,
		Priority:                  nil,
		PriorityID:                nil,
		Status:                    nil,
		StatusID:                  nil,
		IssueType:                 "c3ec5ef2557f6a5f",
		IssueTypeID:               nil,
		Resolution:                nil,
		ResolutionID:              nil,
		State:                     "20b2753657",
		CustomFieldIds:            nil,
		Teams:                     nil,
		ParentIssueID:             nil,
		ParentsIssueIds:           nil,
		Metadata:                  nil,
		ProjectID:                 "f4662184b65cb607",
		Sprints:                   nil,
		Labels:                    nil,
		TopLevel:                  true,
		IsLeaf:                    true,
		Path:                      "02acd163487bcc55",
		InProgressDuration:        int64(64),
		VerificationDuration:      int64(64),
		InProgressCount:           int32(32),
		ReopenCount:               int32(32),
		MappedType:                "958efdd335b319af",
		StrategicParentID:         nil,
		SprintID:                  "{}",
		IssueProjectName:          "a0d3e41b208b8436",
		Users:                     nil,
		InitialStartDate:          int64(64),
		TotalDuration:             int64(64),
		CreatedAt:                 nil,
		ClosedAt:                  nil,
		PlannedEndDate:            nil,
		CustomerID:                "7e0e01202491f425",
		RefType:                   "13cf4b36049dfe2d",
		RefID:                     "1d293261eb293454",
		CustomFieldIdsVirtual:     nil,
		ReleaseDuration:           int64(64),
	}
	ctx := context.Background()
	db := GetDatabase()
	DeleteAllIssueSummaries(ctx, db)
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
	found, err := FindIssueSummaryByID(ctx, db, r.ID)
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
	results, err := FindIssueSummaries(ctx, db)
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
	r.SetIssueID("616142b410120e75")

	r.SetTotalIssues(int32(320))

	r.SetNew30Days(int32(320))

	r.SetTotalClosed(int32(320))

	r.SetClosed30Days(int32(320))

	r.SetEstimatedWorkMonths(float64(64.1))

	r.SetEstimatedWorkMonths30Days(float64(64.1))

	r.SetTitle("7f7461b968286171")

	r.SetIssueType("daa2c182cd7c5dba")

	r.SetState("da11e14679")

	r.SetProjectID("c9b8962d56d68088")

	r.SetTopLevel(false)

	r.SetIsLeaf(false)

	r.SetPath("1cdbbacbdd05c738")

	r.SetInProgressDuration(int64(640))

	r.SetVerificationDuration(int64(640))

	r.SetInProgressCount(int32(320))

	r.SetReopenCount(int32(320))

	r.SetMappedType("499e6c5529670287")

	r.SetSprintID("{}")

	r.SetIssueProjectName("26594eda04a886ec")

	r.SetInitialStartDate(int64(640))

	r.SetTotalDuration(int64(640))

	r.SetCustomerID("21b82b783819469c")

	r.SetRefType("2edf32c7c3242e46")

	r.SetRefID("d4433b91865f8626")

	r.SetReleaseDuration(int64(640))

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

func TestCreateIssueSummaryDeleteTx(t *testing.T) {
	r := &IssueSummary{
		ID:                        "e7e39abcec0c73bb",
		Checksum:                  nil,
		IssueID:                   "7b265bb3e2ee9fb3",
		TotalIssues:               int32(32),
		New30Days:                 int32(32),
		TotalClosed:               int32(32),
		Closed30Days:              int32(32),
		EstimatedWorkMonths:       float64(6.4),
		EstimatedWorkMonths30Days: float64(6.4),
		Title:                     "00f4a631851d85fe",
		URL:                       nil,
		Priority:                  nil,
		PriorityID:                nil,
		Status:                    nil,
		StatusID:                  nil,
		IssueType:                 "c3ec5ef2557f6a5f",
		IssueTypeID:               nil,
		Resolution:                nil,
		ResolutionID:              nil,
		State:                     "20b2753657",
		CustomFieldIds:            nil,
		Teams:                     nil,
		ParentIssueID:             nil,
		ParentsIssueIds:           nil,
		Metadata:                  nil,
		ProjectID:                 "f4662184b65cb607",
		Sprints:                   nil,
		Labels:                    nil,
		TopLevel:                  true,
		IsLeaf:                    true,
		Path:                      "02acd163487bcc55",
		InProgressDuration:        int64(64),
		VerificationDuration:      int64(64),
		InProgressCount:           int32(32),
		ReopenCount:               int32(32),
		MappedType:                "958efdd335b319af",
		StrategicParentID:         nil,
		SprintID:                  "{}",
		IssueProjectName:          "a0d3e41b208b8436",
		Users:                     nil,
		InitialStartDate:          int64(64),
		TotalDuration:             int64(64),
		CreatedAt:                 nil,
		ClosedAt:                  nil,
		PlannedEndDate:            nil,
		CustomerID:                "7e0e01202491f425",
		RefType:                   "13cf4b36049dfe2d",
		RefID:                     "1d293261eb293454",
		CustomFieldIdsVirtual:     nil,
		ReleaseDuration:           int64(64),
	}
	ctx := context.Background()
	db := GetDatabase()
	DeleteAllIssueSummaries(ctx, db)
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
	found, err := FindIssueSummaryByIDTx(ctx, tx, r.ID)
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
	results, err := FindIssueSummariesTx(ctx, tx)
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
	r.SetIssueID("616142b410120e75")

	r.SetTotalIssues(int32(320))

	r.SetNew30Days(int32(320))

	r.SetTotalClosed(int32(320))

	r.SetClosed30Days(int32(320))

	r.SetEstimatedWorkMonths(float64(64.1))

	r.SetEstimatedWorkMonths30Days(float64(64.1))

	r.SetTitle("7f7461b968286171")

	r.SetIssueType("daa2c182cd7c5dba")

	r.SetState("da11e14679")

	r.SetProjectID("c9b8962d56d68088")

	r.SetTopLevel(false)

	r.SetIsLeaf(false)

	r.SetPath("1cdbbacbdd05c738")

	r.SetInProgressDuration(int64(640))

	r.SetVerificationDuration(int64(640))

	r.SetInProgressCount(int32(320))

	r.SetReopenCount(int32(320))

	r.SetMappedType("499e6c5529670287")

	r.SetSprintID("{}")

	r.SetIssueProjectName("26594eda04a886ec")

	r.SetInitialStartDate(int64(640))

	r.SetTotalDuration(int64(640))

	r.SetCustomerID("21b82b783819469c")

	r.SetRefType("2edf32c7c3242e46")

	r.SetRefID("d4433b91865f8626")

	r.SetReleaseDuration(int64(640))

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
