package author

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"testing"
	"time"
)

func TestAuthorCollection(t *testing.T) {

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		t.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer client.Disconnect(ctx)

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		t.Fatal(err)
	}

	database := client.Database("tpm_morphia")
	authorCollection := database.Collection("authors")

	if err := removeAll(t, err, authorCollection, ctx); err != nil {
		t.Fatal(err)
	}

	if err := insertAFirstRecord(t, ctx, authorCollection); err != nil {
		t.Fatal(err)
	}

	if err := insertASecondRecord(t, ctx, authorCollection); err != nil {
		t.Fatal(err)
	}

	if err := queryUsingBsonD(t, err, authorCollection, ctx); err != nil {
		t.Fatal(err)
	}

	if err := queryUsingBsonM(t, err, authorCollection, ctx); err != nil {
		t.Fatal(err)
	}

	if err := queryUsingFilter(t, err, authorCollection, ctx); err != nil {
		t.Fatal(err)
	}
}

func queryUsingBsonD(t *testing.T, err error, authorCollection *mongo.Collection, ctx context.Context) error {

	/* (BadValue) $or must be an array
	filterD := bson.D{
		{LASTNAME, "Smith"},
		{"$or", bson.D{{ "fn", "John"}, { "fn", "Colin"}} },
	}
	*/

	/*
	 * https://stackoverflow.com/questions/26932298/mongodb-in-go-golang-with-mgo-how-to-use-logical-operators-to-query
	 * db.authors.find({ "ln" : "Smith", $or : [{ "fn": "Colin" }, { "fn": "John" }] })
	 * bson.A and []interface{} are equivalente
	 */
	filterD := bson.D{
		{LASTNAME, "Smith"},
		{"$or", bson.A{bson.D{{"fn", "John"}}, bson.D{{"fn", "Colin"}}}},
	}

	filterD = append(filterD, bson.E{ADDRESS_CITY, "Rome"})
	cur, err := authorCollection.Find(ctx, filterD)
	if err != nil {
		return err
	}

	defer cur.Close(ctx)
	consumeResultset(t, ctx, cur)
	return nil
}

func queryUsingBsonM(t *testing.T, err error, authorCollection *mongo.Collection, ctx context.Context) error {
	filterM := bson.M{FIRSTNAME: "John", LASTNAME: "Smith"}
	filterM[ADDRESS_CITY] = "Rome"

	cur, err := authorCollection.Find(ctx, filterM)
	if err != nil {
		return err
	}

	defer cur.Close(ctx)
	consumeResultset(t, ctx, cur)
	return nil
}

func queryUsingFilter(t *testing.T, err error, authorCollection *mongo.Collection, ctx context.Context) error {

	t.Log("queryUsingFilter")

	/*
	 *
	 */
	var f Filter
	t.Log("######## fn = 'John' or fn = 'Colin'")
	f = Filter{}
	f.or().firstNameEqTo("John")
	f.or().firstNameEqTo("Colin")

	filterD := f.build()
	t.Log("resulting filter: ", filterD)
	cur, err := authorCollection.Find(ctx, filterD)
	if err != nil {
		return err
	}

	defer cur.Close(ctx)
	if nr := consumeResultset(t, ctx, cur); nr != 2 {
		t.Error("# records found: ", nr, " expected 2")
	}

	/*
	 *
	 */
	t.Log("######## fn = 'John' and ln = 'Smith'")
	f = Filter{}
	f.or().firstNameEqTo("John").lastNameEqTo("Smith")

	filterD = f.build()
	t.Log("resulting filter: ", filterD)
	cur, err = authorCollection.Find(ctx, filterD)
	if err != nil {
		return err
	}

	defer cur.Close(ctx)
	if nr := consumeResultset(t, ctx, cur); nr != 1 {
		t.Error("# records found: ", nr, " expected 1")
	}

	/*
	 *
	 */
	t.Log("######## ln = 'Smith' and age gt 29")
	f = Filter{}
	f.or().lastNameEqTo("Smith").ageGt(29)

	filterD = f.build()
	t.Log("resulting filter: ", filterD)
	cur, err = authorCollection.Find(ctx, filterD)
	if err != nil {
		return err
	}

	defer cur.Close(ctx)
	if nr := consumeResultset(t, ctx, cur); nr != 2 {
		t.Error("# records found: ", nr, " expected 2")
	}

	/*
	 *
	 */
	t.Log("######## (fn = 'John' or fn = 'Colin') and ln = 'Smith'")
	t.Log("that is: (fn = 'John'  and ln = 'Smith') or (fn = 'Colin'  and ln = 'Smith')")

	f = Filter{}
	f.or().firstNameEqTo("John").lastNameEqTo("Smith")
	f.or().firstNameEqTo("Colin").lastNameEqTo("Smith")

	filterD = f.build()
	t.Log("resulting filter: ", filterD)
	cur, err = authorCollection.Find(ctx, filterD)
	if err != nil {
		return err
	}

	defer cur.Close(ctx)
	if nr := consumeResultset(t, ctx, cur); nr != 2 {
		t.Error("# records found: ", nr, " expected 2")
	}

	/*
	 *
	 */
	t.Log("######## (addr.city = 'Rome' and addr.strt = 'St.Peter Square' and ln = 'Smith' and fn = 'Colin') or (fn = 'John' and books.title = 'My Best Seller')")

	f = Filter{}
	ca := f.or().lastNameEqTo("Smith").firstNameEqTo("Colin")
	ca.address().cityEqTo("Rome").strtEqTo("St.Peter Square")

	ca = f.or().firstNameEqTo("John")
	ca.books().titleEqTo("My Best Seller")
	// ca.books().titleEqTo("My Best Seller", 0)

	filterD = f.build()
	t.Log("resulting filter: ", filterD)
	cur, err = authorCollection.Find(ctx, filterD)
	if err != nil {
		return err
	}

	defer cur.Close(ctx)
	if nr := consumeResultset(t, ctx, cur); nr != 2 {
		t.Error("# records found: ", nr, " expected 2")
	}

	/*
	 *
	 */
	t.Log("######## (fn in [ 'Colin', 'John'])")

	f = Filter{}
	f.or().firstNameIn([]string{"John", "Colin"})

	filterD = f.build()
	t.Log("resulting filter: ", filterD)
	cur, err = authorCollection.Find(ctx, filterD)
	if err != nil {
		return err
	}

	defer cur.Close(ctx)
	if nr := consumeResultset(t, ctx, cur); nr != 2 {
		t.Error("# records found: ", nr, " expected 2")
	}
	return nil
}

func removeAll(t *testing.T, err error, authorCollection *mongo.Collection, ctx context.Context) error {
	deleteResult, err := authorCollection.DeleteMany(ctx, bson.D{})
	if err != nil {
		return err
	}
	t.Log(deleteResult)

	return nil
}

func consumeResultset(t *testing.T, ctx context.Context, cur *mongo.Cursor) int {

	numberOfRecords := 0
	for cur.Next(ctx) {
		numberOfRecords++
		var document bson.D
		err := cur.Decode(&document)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(document)
	}

	if err := cur.Err(); err != nil {
		t.Fatal(err)
	}

	return numberOfRecords
}

func insertAFirstRecord(t *testing.T, ctx context.Context, authorCollection *mongo.Collection) error {
	a := Author{
		FirstName: "John",
		LastName:  "Smith",
		Age:       30,
		Address:   Address{City: "Rome", Strt: "St.Peter Square"},
		Books:     []Book{{Title: "My Best Seller", Posts: []string{"post-1", "post-2"}}},
	}

	t.Log("Inserting a record: ", a)

	_, err := authorCollection.InsertOne(ctx, a)
	if err != nil {
		return err
	}

	return nil
}

func insertASecondRecord(t *testing.T, ctx context.Context, authorCollection *mongo.Collection) error {
	a := Author{
		FirstName: "Colin",
		LastName:  "Smith",
		Age:       40,
		Address:   Address{City: "Rome", Strt: "St.Peter Square"},
		Books:     []Book{{Title: "The life Journey"}},
	}

	t.Log("Inserting a record: ", a)

	_, err := authorCollection.InsertOne(ctx, a)
	if err != nil {
		return err
	}

	return nil
}
