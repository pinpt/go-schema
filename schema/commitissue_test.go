// GENERATED CODE. DO NOT EDIT

package schema

import (
	"context"
	"testing"

	"github.com/jhaynie/go-gator/orm"
)

func TestCreateCommitIssueTable(t *testing.T) {
	tx, err := GetDatabase().Begin()
	if err != nil {
		t.Fatal(err)
	}
	err = DBCreateCommitIssueTableTx(context.Background(), tx)
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

func TestCreateCommitIssueDelete(t *testing.T) {
	r := &CommitIssue{
		ID:           "5c490bbaa95db827",
		Checksum:     nil,
		CommitID:     "1e61d814f625b711",
		Branch:       "3978fb9b3b344152",
		UserID:       "fc2f9a1706b3ff87",
		RepoID:       "7fcbe20f311e2e12",
		IssueID:      "12b720395714c71a",
		Date:         int64(64),
		CustomerID:   "ae0f2fbc80d95d28",
		RefType:      "10cec76d5943b66a",
		RefCommitID:  "7c471507f80a4b9f",
		RefRepoID:    "bba16596ad6a1347",
		RefIssueType: "786ac4b859bd924a",
		RefIssueID:   "8f18cebef83c63ab",
		Metadata:     nil,
	}
	ctx := context.Background()
	db := GetDatabase()
	DeleteAllCommitIssues(ctx, db)
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
	found, err := FindCommitIssueByID(ctx, db, r.ID)
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
	results, err := FindCommitIssues(ctx, db)
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
	r.SetCommitID("9371ff6c1d9edef6")

	r.SetBranch("7546bbcb6232a4db")

	r.SetUserID("a9a354a6bed677eb")

	r.SetRepoID("cd9df9285adfe2ed")

	r.SetIssueID("fd8b7ffcafb0a74e")

	r.SetDate(int64(640))

	r.SetCustomerID("db7b1edd2b4d7657")

	r.SetRefType("826e1b9076ce56fa")

	r.SetRefCommitID("2df4195f6f13b731")

	r.SetRefRepoID("1c606da31076252c")

	r.SetRefIssueType("9e2541a962200f7d")

	r.SetRefIssueID("6bbfc799c96c5bfe")

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

func TestCreateCommitIssueDeleteTx(t *testing.T) {
	r := &CommitIssue{
		ID:           "5c490bbaa95db827",
		Checksum:     nil,
		CommitID:     "1e61d814f625b711",
		Branch:       "3978fb9b3b344152",
		UserID:       "fc2f9a1706b3ff87",
		RepoID:       "7fcbe20f311e2e12",
		IssueID:      "12b720395714c71a",
		Date:         int64(64),
		CustomerID:   "ae0f2fbc80d95d28",
		RefType:      "10cec76d5943b66a",
		RefCommitID:  "7c471507f80a4b9f",
		RefRepoID:    "bba16596ad6a1347",
		RefIssueType: "786ac4b859bd924a",
		RefIssueID:   "8f18cebef83c63ab",
		Metadata:     nil,
	}
	ctx := context.Background()
	db := GetDatabase()
	DeleteAllCommitIssues(ctx, db)
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
	found, err := FindCommitIssueByIDTx(ctx, tx, r.ID)
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
	results, err := FindCommitIssuesTx(ctx, tx)
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
	r.SetCommitID("9371ff6c1d9edef6")

	r.SetBranch("7546bbcb6232a4db")

	r.SetUserID("a9a354a6bed677eb")

	r.SetRepoID("cd9df9285adfe2ed")

	r.SetIssueID("fd8b7ffcafb0a74e")

	r.SetDate(int64(640))

	r.SetCustomerID("db7b1edd2b4d7657")

	r.SetRefType("826e1b9076ce56fa")

	r.SetRefCommitID("2df4195f6f13b731")

	r.SetRefRepoID("1c606da31076252c")

	r.SetRefIssueType("9e2541a962200f7d")

	r.SetRefIssueID("6bbfc799c96c5bfe")

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
