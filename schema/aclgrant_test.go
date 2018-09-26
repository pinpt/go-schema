// GENERATED CODE. DO NOT EDIT

package schema

import (
	"context"
	"testing"

	"github.com/jhaynie/go-gator/orm"
)

func TestCreateACLGrantTable(t *testing.T) {
	tx, err := GetDatabase().Begin()
	if err != nil {
		t.Fatal(err)
	}
	err = DBCreateACLGrantTableTx(context.Background(), tx)
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

func TestCreateACLGrantDelete(t *testing.T) {
	r := &ACLGrant{
		ID:          "f531e753b78e4604",
		Checksum:    nil,
		ResourceID:  "dafeb4406b7615d8",
		RoleID:      "fc3ada3ff861e784",
		Permission:  ACLGrantPermission_READ,
		AdminUserID: "c8ac234d97fc55d7",
		CustomerID:  "54d65040f54385cf",
		CreatedAt:   int64(64),
		UpdatedAt:   nil,
	}
	ctx := context.Background()
	db := GetDatabase()
	DeleteAllACLGrants(ctx, db)
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
	found, err := FindACLGrantByID(ctx, db, r.ID)
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
	results, err := FindACLGrants(ctx, db)
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
	r.SetResourceID("eade060a487bceda")

	r.SetRoleID("a9f5e3f911ba20c4")

	r.SetPermission(ACLGrantPermission_READ)

	r.SetAdminUserID("64ea389acb09d71a")

	r.SetCustomerID("a98db6b659d27138")

	r.SetCreatedAt(int64(640))

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

func TestCreateACLGrantDeleteTx(t *testing.T) {
	r := &ACLGrant{
		ID:          "f531e753b78e4604",
		Checksum:    nil,
		ResourceID:  "dafeb4406b7615d8",
		RoleID:      "fc3ada3ff861e784",
		Permission:  ACLGrantPermission_READ,
		AdminUserID: "c8ac234d97fc55d7",
		CustomerID:  "54d65040f54385cf",
		CreatedAt:   int64(64),
		UpdatedAt:   nil,
	}
	ctx := context.Background()
	db := GetDatabase()
	DeleteAllACLGrants(ctx, db)
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
	found, err := FindACLGrantByIDTx(ctx, tx, r.ID)
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
	results, err := FindACLGrantsTx(ctx, tx)
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
	r.SetResourceID("eade060a487bceda")

	r.SetRoleID("a9f5e3f911ba20c4")

	r.SetPermission(ACLGrantPermission_READ)

	r.SetAdminUserID("64ea389acb09d71a")

	r.SetCustomerID("a98db6b659d27138")

	r.SetCreatedAt(int64(640))

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
