// GENERATED CODE. DO NOT EDIT

package schema

import (
	"context"
	"testing"

	"github.com/jhaynie/go-gator/orm"
)

func TestCreateIssueCommentTable(t *testing.T) {
	tx, err := GetDatabase().Begin()
	if err != nil {
		t.Fatal(err)
	}
	err = DBCreateIssueCommentTableTx(context.Background(), tx)
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

func TestCreateIssueCommentDelete(t *testing.T) {
	r := &IssueComment{
		ID:            "b12ccc9d86c34d52",
		Checksum:      nil,
		IssueID:       "7681c6f33112efd6",
		UserID:        nil,
		ProjectID:     "5eb9248c5c82caaa",
		Body:          "016bf69bb7584b7d",
		Deleted:       true,
		DeletedUserID: nil,
		CreatedAt:     int64(64),
		UpdatedAt:     nil,
		DeletedAt:     nil,
		URL:           "291e95d433bcba53",
		CustomerID:    "b260b457501cdb91",
		RefType:       "ae1c7a4514172f69",
		RefID:         "e712ddab8c137f53",
		Metadata:      nil,
	}
	ctx := context.Background()
	db := GetDatabase()
	DeleteAllIssueComments(ctx, db)
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
	found, err := FindIssueCommentByID(ctx, db, r.ID)
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
	results, err := FindIssueComments(ctx, db)
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
	r.SetIssueID("7d09fd3f4fe8401e")

	r.SetProjectID("b1e84e670a114a36")

	r.SetBody("9e849130e0156cab")

	r.SetDeleted(false)

	r.SetCreatedAt(int64(640))

	r.SetURL("cfc22999a8fc0a81")

	r.SetCustomerID("939a6572c67c5958")

	r.SetRefType("3842174d705eb59b")

	r.SetRefID("c8e6d27dfd5f5aca")

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

func TestCreateIssueCommentDeleteTx(t *testing.T) {
	r := &IssueComment{
		ID:            "b12ccc9d86c34d52",
		Checksum:      nil,
		IssueID:       "7681c6f33112efd6",
		UserID:        nil,
		ProjectID:     "5eb9248c5c82caaa",
		Body:          "016bf69bb7584b7d",
		Deleted:       true,
		DeletedUserID: nil,
		CreatedAt:     int64(64),
		UpdatedAt:     nil,
		DeletedAt:     nil,
		URL:           "291e95d433bcba53",
		CustomerID:    "b260b457501cdb91",
		RefType:       "ae1c7a4514172f69",
		RefID:         "e712ddab8c137f53",
		Metadata:      nil,
	}
	ctx := context.Background()
	db := GetDatabase()
	DeleteAllIssueComments(ctx, db)
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
	found, err := FindIssueCommentByIDTx(ctx, tx, r.ID)
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
	results, err := FindIssueCommentsTx(ctx, tx)
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
	r.SetIssueID("7d09fd3f4fe8401e")

	r.SetProjectID("b1e84e670a114a36")

	r.SetBody("9e849130e0156cab")

	r.SetDeleted(false)

	r.SetCreatedAt(int64(640))

	r.SetURL("cfc22999a8fc0a81")

	r.SetCustomerID("939a6572c67c5958")

	r.SetRefType("3842174d705eb59b")

	r.SetRefID("c8e6d27dfd5f5aca")

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
