// GENERATED CODE. DO NOT EDIT

package schema

import (
	"context"
	"testing"

	"github.com/jhaynie/go-gator/orm"
)

func TestCreateCommitTable(t *testing.T) {
	tx, err := GetDatabase().Begin()
	if err != nil {
		t.Fatal(err)
	}
	err = DBCreateCommitTableTx(context.Background(), tx)
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

func TestCreateCommitDelete(t *testing.T) {
	r := &Commit{
		ID:              "00e4372ec5f3e99e",
		Checksum:        nil,
		RepoID:          "a30cf7f3b4b6a004",
		Sha:             "6ed922f58d7b9f27",
		Branch:          nil,
		Message:         nil,
		Mergecommit:     nil,
		Excluded:        nil,
		Parent:          nil,
		ParentID:        nil,
		Date:            int64(64),
		AuthorUserID:    "6256657d519ef4c2",
		CommitterUserID: "0a410dbb35d8c6fc",
		Ordinal:         nil,
		CustomerID:      "ac4f16679d54b54d",
		RefType:         "8d72d8f36ad18a74",
		RefID:           "950c4a9b32683ad4",
		Metadata:        nil,
	}
	ctx := context.Background()
	db := GetDatabase()
	DeleteAllCommits(ctx, db)
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
	found, err := FindCommitByID(ctx, db, r.ID)
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
	results, err := FindCommits(ctx, db)
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
	r.SetRepoID("2278ddc9b9c468e5")

	r.SetSha("e49e1f004ff97b20")

	r.SetDate(int64(640))

	r.SetAuthorUserID("45849e4dcb2adcc3")

	r.SetCommitterUserID("5e6201dad9437e80")

	r.SetCustomerID("8f1666ddf82b39de")

	r.SetRefType("94ea1aade9a442f9")

	r.SetRefID("c0e51339082249c8")

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

func TestCreateCommitDeleteTx(t *testing.T) {
	r := &Commit{
		ID:              "00e4372ec5f3e99e",
		Checksum:        nil,
		RepoID:          "a30cf7f3b4b6a004",
		Sha:             "6ed922f58d7b9f27",
		Branch:          nil,
		Message:         nil,
		Mergecommit:     nil,
		Excluded:        nil,
		Parent:          nil,
		ParentID:        nil,
		Date:            int64(64),
		AuthorUserID:    "6256657d519ef4c2",
		CommitterUserID: "0a410dbb35d8c6fc",
		Ordinal:         nil,
		CustomerID:      "ac4f16679d54b54d",
		RefType:         "8d72d8f36ad18a74",
		RefID:           "950c4a9b32683ad4",
		Metadata:        nil,
	}
	ctx := context.Background()
	db := GetDatabase()
	DeleteAllCommits(ctx, db)
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
	found, err := FindCommitByIDTx(ctx, tx, r.ID)
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
	results, err := FindCommitsTx(ctx, tx)
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
	r.SetRepoID("2278ddc9b9c468e5")

	r.SetSha("e49e1f004ff97b20")

	r.SetDate(int64(640))

	r.SetAuthorUserID("45849e4dcb2adcc3")

	r.SetCommitterUserID("5e6201dad9437e80")

	r.SetCustomerID("8f1666ddf82b39de")

	r.SetRefType("94ea1aade9a442f9")

	r.SetRefID("c0e51339082249c8")

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
