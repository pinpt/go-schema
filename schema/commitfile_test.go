// GENERATED CODE. DO NOT EDIT

package schema

import (
	"context"
	"testing"

	"github.com/jhaynie/go-gator/orm"
)

func TestCreateCommitFileTable(t *testing.T) {
	tx, err := GetDatabase().Begin()
	if err != nil {
		t.Fatal(err)
	}
	err = DBCreateCommitFileTableTx(context.Background(), tx)
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

func TestCreateCommitFileDelete(t *testing.T) {
	r := &CommitFile{
		ID:              "b05054f5f126b172",
		Checksum:        nil,
		CommitID:        "a0c92c763354bc42",
		RepoID:          "6675bcbaa652ad71",
		AuthorUserID:    "397a1b09d7365bdb",
		CommitterUserID: "34bb596e7534adfb",
		Filename:        "eee763e402fc2d5d",
		Language:        "ed717ee5d78d036d",
		Additions:       int32(32),
		Deletions:       int32(32),
		Size:            int32(32),
		Abinary:         true,
		Date:            int64(64),
		Branch:          "839680718a3d4ec5",
		Mergecommit:     true,
		Excluded:        true,
		Loc:             int32(32),
		Sloc:            int32(32),
		Comments:        int32(32),
		Blanks:          int32(32),
		Variance:        int32(32),
		Status:          CommitFileStatus_ADDED,
		Renamed:         true,
		RenamedFrom:     nil,
		RenamedTo:       nil,
		CustomerID:      "bb7b1550fe7b707a",
		RefType:         "42b20cecc7e4ea7e",
		RefID:           "506dc97a45cfc18e",
		Metadata:        nil,
	}
	ctx := context.Background()
	db := GetDatabase()
	DeleteAllCommitFiles(ctx, db)
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
	found, err := FindCommitFileByID(ctx, db, r.ID)
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
	results, err := FindCommitFiles(ctx, db)
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
	r.SetCommitID("7e98582549a69e12")

	r.SetRepoID("92c3b6e04af50705")

	r.SetAuthorUserID("bc0b3b591f27ca06")

	r.SetCommitterUserID("38b667fb9fc47590")

	r.SetFilename("f20d44a4444b18b3")

	r.SetLanguage("3a27231e5fa28be9")

	r.SetAdditions(int32(320))

	r.SetDeletions(int32(320))

	r.SetSize(int32(320))

	r.SetAbinary(false)

	r.SetDate(int64(640))

	r.SetBranch("72f00b7b3aa04b0d")

	r.SetMergecommit(false)

	r.SetExcluded(false)

	r.SetLoc(int32(320))

	r.SetSloc(int32(320))

	r.SetComments(int32(320))

	r.SetBlanks(int32(320))

	r.SetVariance(int32(320))

	r.SetStatus(CommitFileStatus_ADDED)

	r.SetRenamed(false)

	r.SetCustomerID("765f739efa77bdfc")

	r.SetRefType("0cf83f710fabc6c8")

	r.SetRefID("15d3e1078f2e30fa")

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

func TestCreateCommitFileDeleteTx(t *testing.T) {
	r := &CommitFile{
		ID:              "b05054f5f126b172",
		Checksum:        nil,
		CommitID:        "a0c92c763354bc42",
		RepoID:          "6675bcbaa652ad71",
		AuthorUserID:    "397a1b09d7365bdb",
		CommitterUserID: "34bb596e7534adfb",
		Filename:        "eee763e402fc2d5d",
		Language:        "ed717ee5d78d036d",
		Additions:       int32(32),
		Deletions:       int32(32),
		Size:            int32(32),
		Abinary:         true,
		Date:            int64(64),
		Branch:          "839680718a3d4ec5",
		Mergecommit:     true,
		Excluded:        true,
		Loc:             int32(32),
		Sloc:            int32(32),
		Comments:        int32(32),
		Blanks:          int32(32),
		Variance:        int32(32),
		Status:          CommitFileStatus_ADDED,
		Renamed:         true,
		RenamedFrom:     nil,
		RenamedTo:       nil,
		CustomerID:      "bb7b1550fe7b707a",
		RefType:         "42b20cecc7e4ea7e",
		RefID:           "506dc97a45cfc18e",
		Metadata:        nil,
	}
	ctx := context.Background()
	db := GetDatabase()
	DeleteAllCommitFiles(ctx, db)
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
	found, err := FindCommitFileByIDTx(ctx, tx, r.ID)
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
	results, err := FindCommitFilesTx(ctx, tx)
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
	r.SetCommitID("7e98582549a69e12")

	r.SetRepoID("92c3b6e04af50705")

	r.SetAuthorUserID("bc0b3b591f27ca06")

	r.SetCommitterUserID("38b667fb9fc47590")

	r.SetFilename("f20d44a4444b18b3")

	r.SetLanguage("3a27231e5fa28be9")

	r.SetAdditions(int32(320))

	r.SetDeletions(int32(320))

	r.SetSize(int32(320))

	r.SetAbinary(false)

	r.SetDate(int64(640))

	r.SetBranch("72f00b7b3aa04b0d")

	r.SetMergecommit(false)

	r.SetExcluded(false)

	r.SetLoc(int32(320))

	r.SetSloc(int32(320))

	r.SetComments(int32(320))

	r.SetBlanks(int32(320))

	r.SetVariance(int32(320))

	r.SetStatus(CommitFileStatus_ADDED)

	r.SetRenamed(false)

	r.SetCustomerID("765f739efa77bdfc")

	r.SetRefType("0cf83f710fabc6c8")

	r.SetRefID("15d3e1078f2e30fa")

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
