// GENERATED CODE. DO NOT EDIT

package schema

import (
	"context"
	"testing"

	"github.com/jhaynie/go-gator/orm"
)

func TestCreateIssueTable(t *testing.T) {
	tx, err := GetDatabase().Begin()
	if err != nil {
		t.Fatal(err)
	}
	err = DBCreateIssueTableTx(context.Background(), tx)
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

func TestCreateIssueDelete(t *testing.T) {
	r := &Issue{
		ID:             "f91a7b249091d0aa",
		Checksum:       nil,
		UserID:         nil,
		ProjectID:      "c7d8c3e33f8ae702",
		Title:          "957df6ec71df0b47",
		Body:           nil,
		State:          IssueState_OPEN,
		Type:           IssueType_BUG,
		Resolution:     IssueResolution_NONE,
		CreatedAt:      int64(64),
		UpdatedAt:      nil,
		ClosedAt:       nil,
		ClosedbyUserID: nil,
		URL:            "6a2681890a057738",
		CustomerID:     "03737274e6ed7b09",
		RefType:        "d47a72392e37c70d",
		RefID:          "feb9612b3839ed37",
		Metadata:       nil,
	}
	ctx := context.Background()
	db := GetDatabase()
	DeleteAllIssues(ctx, db)
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
	found, err := FindIssueByID(ctx, db, r.ID)
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
	results, err := FindIssues(ctx, db)
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
	r.SetProjectID("0c1a9e4f2fcd95c5")

	r.SetTitle("512912e5c54273ee")

	r.SetState(IssueState_OPEN)

	r.SetType(IssueType_BUG)

	r.SetResolution(IssueResolution_NONE)

	r.SetCreatedAt(int64(640))

	r.SetURL("9c2d71cda660f7dc")

	r.SetCustomerID("e75bc467aec3b3ad")

	r.SetRefType("e3447fa038425687")

	r.SetRefID("cf9693d4853dd00a")

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

func TestCreateIssueDeleteTx(t *testing.T) {
	r := &Issue{
		ID:             "f91a7b249091d0aa",
		Checksum:       nil,
		UserID:         nil,
		ProjectID:      "c7d8c3e33f8ae702",
		Title:          "957df6ec71df0b47",
		Body:           nil,
		State:          IssueState_OPEN,
		Type:           IssueType_BUG,
		Resolution:     IssueResolution_NONE,
		CreatedAt:      int64(64),
		UpdatedAt:      nil,
		ClosedAt:       nil,
		ClosedbyUserID: nil,
		URL:            "6a2681890a057738",
		CustomerID:     "03737274e6ed7b09",
		RefType:        "d47a72392e37c70d",
		RefID:          "feb9612b3839ed37",
		Metadata:       nil,
	}
	ctx := context.Background()
	db := GetDatabase()
	DeleteAllIssues(ctx, db)
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
	found, err := FindIssueByIDTx(ctx, tx, r.ID)
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
	results, err := FindIssuesTx(ctx, tx)
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
	r.SetProjectID("0c1a9e4f2fcd95c5")

	r.SetTitle("512912e5c54273ee")

	r.SetState(IssueState_OPEN)

	r.SetType(IssueType_BUG)

	r.SetResolution(IssueResolution_NONE)

	r.SetCreatedAt(int64(640))

	r.SetURL("9c2d71cda660f7dc")

	r.SetCustomerID("e75bc467aec3b3ad")

	r.SetRefType("e3447fa038425687")

	r.SetRefID("cf9693d4853dd00a")

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
