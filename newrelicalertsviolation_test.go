// GENERATED CODE. DO NOT EDIT

package schema

import (
	"context"
	"testing"

	"github.com/jhaynie/go-gator/orm"
)

func TestCreateNewrelicAlertsViolationTable(t *testing.T) {
	tx, err := GetDatabase().Begin()
	if err != nil {
		t.Fatal(err)
	}
	err = DBCreateNewrelicAlertsViolationTableTx(context.Background(), tx)
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

func TestCreateNewrelicAlertsViolationDelete(t *testing.T) {
	r := &NewrelicAlertsViolation{
		ID:             "1bbfea8eb565f89e",
		Checksum:       nil,
		CustomerID:     "36e5a52c1c7a41ca",
		ExtID:          int64(64),
		Duration:       int64(64),
		PolicyName:     "8fa049a598283197",
		ConditionName:  "ed6c174c2c2e1838",
		Priority:       "2f9153d92148417c",
		OpenedAt:       int64(64),
		EntityProduct:  "40f4a05b4446e603",
		EntityType:     "1db48e24759bf37b",
		EntityGroupID:  int64(64),
		EntityID:       int64(64),
		EntityName:     "50533483d3bb07a1",
		PolicyExtID:    int64(64),
		PolicyID:       "fbd7c1b0b4aa1aaa",
		ConditionExtID: int64(64),
		ConditionID:    "bb56d4759310cfe9",
	}
	ctx := context.Background()
	db := GetDatabase()
	DeleteAllNewrelicAlertsViolations(ctx, db)
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
	found, err := FindNewrelicAlertsViolationByID(ctx, db, r.ID)
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
	results, err := FindNewrelicAlertsViolations(ctx, db)
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
	r.SetCustomerID("d05daf8a02530a4a")

	r.SetExtID(int64(640))

	r.SetDuration(int64(640))

	r.SetPolicyName("f117c0bb721aa163")

	r.SetConditionName("ba9c099f575d4c75")

	r.SetPriority("7cd68d5a05593e43")

	r.SetOpenedAt(int64(640))

	r.SetEntityProduct("f8a3e8aa1ecc812b")

	r.SetEntityType("249610b3021ac52c")

	r.SetEntityGroupID(int64(640))

	r.SetEntityID(int64(640))

	r.SetEntityName("2a4074eaf31e044b")

	r.SetPolicyExtID(int64(640))

	r.SetPolicyID("4af8210c42ad8af7")

	r.SetConditionExtID(int64(640))

	r.SetConditionID("6c7dd26dbf6b4997")

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

func TestCreateNewrelicAlertsViolationDeleteTx(t *testing.T) {
	r := &NewrelicAlertsViolation{
		ID:             "1bbfea8eb565f89e",
		Checksum:       nil,
		CustomerID:     "36e5a52c1c7a41ca",
		ExtID:          int64(64),
		Duration:       int64(64),
		PolicyName:     "8fa049a598283197",
		ConditionName:  "ed6c174c2c2e1838",
		Priority:       "2f9153d92148417c",
		OpenedAt:       int64(64),
		EntityProduct:  "40f4a05b4446e603",
		EntityType:     "1db48e24759bf37b",
		EntityGroupID:  int64(64),
		EntityID:       int64(64),
		EntityName:     "50533483d3bb07a1",
		PolicyExtID:    int64(64),
		PolicyID:       "fbd7c1b0b4aa1aaa",
		ConditionExtID: int64(64),
		ConditionID:    "bb56d4759310cfe9",
	}
	ctx := context.Background()
	db := GetDatabase()
	DeleteAllNewrelicAlertsViolations(ctx, db)
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
	found, err := FindNewrelicAlertsViolationByIDTx(ctx, tx, r.ID)
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
	results, err := FindNewrelicAlertsViolationsTx(ctx, tx)
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
	r.SetCustomerID("d05daf8a02530a4a")

	r.SetExtID(int64(640))

	r.SetDuration(int64(640))

	r.SetPolicyName("f117c0bb721aa163")

	r.SetConditionName("ba9c099f575d4c75")

	r.SetPriority("7cd68d5a05593e43")

	r.SetOpenedAt(int64(640))

	r.SetEntityProduct("f8a3e8aa1ecc812b")

	r.SetEntityType("249610b3021ac52c")

	r.SetEntityGroupID(int64(640))

	r.SetEntityID(int64(640))

	r.SetEntityName("2a4074eaf31e044b")

	r.SetPolicyExtID(int64(640))

	r.SetPolicyID("4af8210c42ad8af7")

	r.SetConditionExtID(int64(640))

	r.SetConditionID("6c7dd26dbf6b4997")

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
