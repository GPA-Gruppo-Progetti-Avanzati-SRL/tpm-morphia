package author

const (
	// author struct: Author
	FIRSTNAME = "fn"
	LASTNAME  = "ln"
	AGE       = "age"
	ADDRESS   = "addr"
	BOOKS     = "books"
	BOOKS_i   = "books.%d"

	// address struct: Address
	ADDRESS_CITY = "addr.city"
	ADDRESS_STRT = "addr.strt"

	// books.[] struct: Book
	BOOKS_TITLE   = "books.title"
	BOOKS_i_TITLE = "books.%d.title"
	BOOKS_ISBN    = "books.isbn"
	BOOKS_i_ISBN  = "books.%d.isbn"
	BOOKS_POSTS   = "books.posts"
	BOOKS_i_POSTS = "books.%d.posts"
)

type Author struct {
	FirstName string  `json:"fn,omitempty" bson:"fn,omitempty"`
	LastName  string  `json:"ln,omitempty" bson:"ln,omitempty"`
	Age       int32   `json:"age,omitempty" bson:"age,omitempty"`
	Address   Address `json:"addr,omitempty" bson:"addr,omitempty"`
	Books     []Book  `json:"books,omitempty" bson:"books,omitempty"`
}

type Address struct {
	City string `json:"city,omitempty" bson:"city,omitempty"`
	Strt string `json:"strt,omitempty" bson:"strt,omitempty"`
}

type Book struct {
	Title string   `json:"title,omitempty" bson:"title,omitempty"`
	Isbn  string   `json:"isbn,omitempty" bson:"isbn,omitempty"`
	Posts []string `json:"posts,omitempty" bson:"posts,omitempty"`
}
