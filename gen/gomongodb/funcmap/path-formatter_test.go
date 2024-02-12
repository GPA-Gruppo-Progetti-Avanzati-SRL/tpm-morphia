package funcmap_test

import (
	"fmt"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-morphia/gen/gomongodb/attributes"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-morphia/gen/gomongodb/funcmap"
	"github.com/stretchr/testify/require"
	"testing"
)

type IW struct {
	path          string
	separator     string
	casing        funcmap.CasingMode
	indexHandling funcmap.IndexHandling
	indexFormat   funcmap.IndexFormat
	wanted        string
}

func TestFormatPath(t *testing.T) {

	iws := []IW{
		{path: "aBook", separator: ".", casing: "camelCase", indexHandling: "index", indexFormat: "indexIjk", wanted: "ABook"},
		{path: "aBook", separator: "_", casing: "classify", indexHandling: "index", indexFormat: "indexIjk", wanted: "ABook"},
		{path: "aBook", separator: "_", casing: "dasherize", indexHandling: "suppress", indexFormat: "indexIjk", wanted: "a-book"},
		{path: "aBook", separator: ".", casing: "none", indexHandling: "index", indexFormat: "indexSprintf", wanted: "aBook"},
		{path: "aBook.isbn", separator: "_", casing: "classify", indexHandling: "index", indexFormat: "indexIjk", wanted: "ABook_Isbn"},
		{path: "aBook.isbn", separator: ".", casing: "none", indexHandling: "index", indexFormat: "indexSprintf", wanted: "aBook.isbn"},
		{path: "aBook.title", separator: "_", casing: "classify", indexHandling: "index", indexFormat: "indexIjk", wanted: "ABook_Title"},
		{path: "aBook.title", separator: ".", casing: "none", indexHandling: "index", indexFormat: "indexSprintf", wanted: "aBook.title"},
		{path: "age", separator: ".", casing: "camelCase", indexHandling: "index", indexFormat: "indexIjk", wanted: "Age"},
		{path: "age", separator: "_", casing: "classify", indexHandling: "index", indexFormat: "indexIjk", wanted: "Age"},
		{path: "age", separator: "_", casing: "dasherize", indexHandling: "suppress", indexFormat: "indexIjk", wanted: "age"},
		{path: "age", separator: ".", casing: "none", indexHandling: "index", indexFormat: "indexSprintf", wanted: "age"},
		{path: "aBook.coAuthors", separator: "_", casing: "classify", indexHandling: "index", indexFormat: "indexIjk", wanted: "ABook_CoAuthors"},
		{path: "aBook.coAuthors.[]", separator: "_", casing: "classify", indexHandling: "index", indexFormat: "indexIjk", wanted: "ABook_CoAuthors_I"},
		{path: "aBook.coAuthors", separator: ".", casing: "none", indexHandling: "index", indexFormat: "indexSprintf", wanted: "aBook.coAuthors"},
		{path: "aBook.coAuthors.[]", separator: ".", casing: "none", indexHandling: "index", indexFormat: "indexSprintf", wanted: "aBook.coAuthors.%d"},
		{path: "arrayOfArrayOfBooks.[].coAuthors", separator: "_", casing: "classify", indexHandling: "index", indexFormat: "indexIjk", wanted: "ArrayOfArrayOfBooks_I_CoAuthors"},
		{path: "arrayOfArrayOfBooks.[].[].coAuthors", separator: "_", casing: "classify", indexHandling: "index", indexFormat: "indexIjk", wanted: "ArrayOfArrayOfBooks_I_J_CoAuthors"},
		{path: "arrayOfArrayOfBooks.[].[].coAuthors.[]", separator: "_", casing: "classify", indexHandling: "index", indexFormat: "indexIjk", wanted: "ArrayOfArrayOfBooks_I_J_CoAuthors_K"},
		{path: "arrayOfArrayOfBooks.[].coAuthors", separator: ".", casing: "none", indexHandling: "index", indexFormat: "indexSprintf", wanted: "arrayOfArrayOfBooks.%d.coAuthors"},
		{path: "arrayOfArrayOfBooks.[].[].coAuthors", separator: ".", casing: "none", indexHandling: "index", indexFormat: "indexSprintf", wanted: "arrayOfArrayOfBooks.%d.%d.coAuthors"},
		{path: "arrayOfArrayOfBooks.[].[].coAuthors.[]", separator: ".", casing: "none", indexHandling: "index", indexFormat: "indexSprintf", wanted: "arrayOfArrayOfBooks.%d.%d.coAuthors.%d"},
		{path: "arrayOfArrayOfBooks.[].isbn", separator: "_", casing: "classify", indexHandling: "index", indexFormat: "indexIjk", wanted: "ArrayOfArrayOfBooks_I_Isbn"},
		{path: "arrayOfArrayOfBooks.[].[].isbn", separator: "_", casing: "classify", indexHandling: "index", indexFormat: "indexIjk", wanted: "ArrayOfArrayOfBooks_I_J_Isbn"},
		{path: "arrayOfArrayOfBooks.[].[].isbn", separator: ".", casing: "none", indexHandling: "index", indexFormat: "indexSprintf", wanted: "arrayOfArrayOfBooks.%d.%d.isbn"},
		{path: "arrayOfArrayOfBooks.[].isbn", separator: ".", casing: "none", indexHandling: "index", indexFormat: "indexSprintf", wanted: "arrayOfArrayOfBooks.%d.isbn"},
		{path: "mapOfArrayOfBooks.%s.coAuthors", separator: "_", casing: "classify", indexHandling: "index", indexFormat: "indexIjk", wanted: "MapOfArrayOfBooks_S_CoAuthors"},
		{path: "mapOfArrayOfBooks.%s.[].coAuthors", separator: "_", casing: "classify", indexHandling: "index", indexFormat: "indexIjk", wanted: "MapOfArrayOfBooks_S_I_CoAuthors"},
		{path: "mapOfArrayOfBooks.%s.[].coAuthors.[]", separator: "_", casing: "classify", indexHandling: "index", indexFormat: "indexIjk", wanted: "MapOfArrayOfBooks_S_I_CoAuthors_J"},
		{path: "mapOfArrayOfBooks.%s.coAuthors", separator: ".", casing: "none", indexHandling: "index", indexFormat: "indexSprintf", wanted: "mapOfArrayOfBooks.%s.coAuthors"},
		{path: "mapOfArrayOfBooks.%s.[].coAuthors", separator: ".", casing: "none", indexHandling: "index", indexFormat: "indexSprintf", wanted: "mapOfArrayOfBooks.%s.%d.coAuthors"},
		{path: "mapOfArrayOfBooks.%s.[].coAuthors.[]", separator: ".", casing: "none", indexHandling: "index", indexFormat: "indexSprintf", wanted: "mapOfArrayOfBooks.%s.%d.coAuthors.%d"},
		{path: "mapOfArrayOfBooks.%s.[].isbn", separator: "_", casing: "classify", indexHandling: "index", indexFormat: "indexIjk", wanted: "MapOfArrayOfBooks_S_I_Isbn"},
		{path: "mapOfArrayOfBooks.%s.isbn", separator: "_", casing: "classify", indexHandling: "index", indexFormat: "indexIjk", wanted: "MapOfArrayOfBooks_S_Isbn"},
		{path: "mapOfArrayOfBooks.%s.[].isbn", separator: ".", casing: "none", indexHandling: "index", indexFormat: "indexSprintf", wanted: "mapOfArrayOfBooks.%s.%d.isbn"},
		{path: "mapOfArrayOfBooks.%s.isbn", separator: ".", casing: "none", indexHandling: "index", indexFormat: "indexSprintf", wanted: "mapOfArrayOfBooks.%s.isbn"},
		{path: "mapOfBooks.%s.coAuthors", separator: "_", casing: "classify", indexHandling: "index", indexFormat: "indexIjk", wanted: "MapOfBooks_S_CoAuthors"},
		{path: "mapOfBooks.%s.coAuthors.[]", separator: "_", casing: "classify", indexHandling: "index", indexFormat: "indexIjk", wanted: "MapOfBooks_S_CoAuthors_I"},
		{path: "mapOfBooks.%s.coAuthors", separator: ".", casing: "none", indexHandling: "index", indexFormat: "indexSprintf", wanted: "mapOfBooks.%s.coAuthors"},
		{path: "mapOfBooks.%s.coAuthors.[]", separator: ".", casing: "none", indexHandling: "index", indexFormat: "indexSprintf", wanted: "mapOfBooks.%s.coAuthors.%d"},
		{path: "mapOfBooks.%s.isbn", separator: "_", casing: "classify", indexHandling: "index", indexFormat: "indexIjk", wanted: "MapOfBooks_S_Isbn"},
		{path: "mapOfBooks.%s.isbn", separator: ".", casing: "none", indexHandling: "index", indexFormat: "indexSprintf", wanted: "mapOfBooks.%s.isbn"},
	}

	for i, iw := range iws {
		actual := funcmap.FormatPath(iw.path, iw.separator, iw.casing, iw.indexHandling, iw.indexFormat)
		require.Equal(t, iw.wanted, actual, fmt.Sprintf("[%d]", i))
	}
}

func TestNewFormatPath(t *testing.T) {

	iws := []IW{
		{path: "aBook", separator: ".", casing: "camelCase", indexHandling: "index", indexFormat: "indexIjk", wanted: "ABook"},
		{path: "aBook", separator: "_", casing: "classify", indexHandling: "index", indexFormat: "indexIjk", wanted: "ABook"},
		{path: "aBook", separator: "_", casing: "dasherize", indexHandling: "suppress", indexFormat: "indexIjk", wanted: "a-book"},
		{path: "aBook", separator: ".", casing: "none", indexHandling: "index", indexFormat: "indexSprintf", wanted: "aBook"},
		{path: "aBook.isbn", separator: "_", casing: "classify", indexHandling: "index", indexFormat: "indexIjk", wanted: "ABook_Isbn"},
		{path: "aBook.isbn", separator: ".", casing: "none", indexHandling: "index", indexFormat: "indexSprintf", wanted: "aBook.isbn"},
		{path: "aBook.title", separator: "_", casing: "classify", indexHandling: "index", indexFormat: "indexIjk", wanted: "ABook_Title"},
		{path: "aBook.title", separator: ".", casing: "none", indexHandling: "index", indexFormat: "indexSprintf", wanted: "aBook.title"},
		{path: "age", separator: ".", casing: "camelCase", indexHandling: "index", indexFormat: "indexIjk", wanted: "Age"},
		{path: "age", separator: "_", casing: "classify", indexHandling: "index", indexFormat: "indexIjk", wanted: "Age"},
		{path: "age", separator: "_", casing: "dasherize", indexHandling: "suppress", indexFormat: "indexIjk", wanted: "age"},
		{path: "age", separator: ".", casing: "none", indexHandling: "index", indexFormat: "indexSprintf", wanted: "age"},
		{path: "aBook.coAuthors", separator: "_", casing: "classify", indexHandling: "index", indexFormat: "indexIjk", wanted: "ABook_CoAuthors"},
		{path: "aBook.coAuthors.[]", separator: "_", casing: "classify", indexHandling: "index", indexFormat: "indexIjk", wanted: "ABook_CoAuthors_I"},
		{path: "aBook.coAuthors", separator: ".", casing: "none", indexHandling: "index", indexFormat: "indexSprintf", wanted: "aBook.coAuthors"},
		{path: "aBook.coAuthors.[]", separator: ".", casing: "none", indexHandling: "index", indexFormat: "indexSprintf", wanted: "aBook.coAuthors.%d"},
		{path: "arrayOfArrayOfBooks.[].coAuthors", separator: "_", casing: "classify", indexHandling: "index", indexFormat: "indexIjk", wanted: "ArrayOfArrayOfBooks_I_CoAuthors"},
		{path: "arrayOfArrayOfBooks.[].[].coAuthors", separator: "_", casing: "classify", indexHandling: "index", indexFormat: "indexIjk", wanted: "ArrayOfArrayOfBooks_I_J_CoAuthors"},
		{path: "arrayOfArrayOfBooks.[].[].coAuthors.[]", separator: "_", casing: "classify", indexHandling: "index", indexFormat: "indexIjk", wanted: "ArrayOfArrayOfBooks_I_J_CoAuthors_K"},
		{path: "arrayOfArrayOfBooks.[].coAuthors", separator: ".", casing: "none", indexHandling: "index", indexFormat: "indexSprintf", wanted: "arrayOfArrayOfBooks.%d.coAuthors"},
		{path: "arrayOfArrayOfBooks.[].[].coAuthors", separator: ".", casing: "none", indexHandling: "index", indexFormat: "indexSprintf", wanted: "arrayOfArrayOfBooks.%d.%d.coAuthors"},
		{path: "arrayOfArrayOfBooks.[].[].coAuthors.[]", separator: ".", casing: "none", indexHandling: "index", indexFormat: "indexSprintf", wanted: "arrayOfArrayOfBooks.%d.%d.coAuthors.%d"},
		{path: "arrayOfArrayOfBooks.[].isbn", separator: "_", casing: "classify", indexHandling: "index", indexFormat: "indexIjk", wanted: "ArrayOfArrayOfBooks_I_Isbn"},
		{path: "arrayOfArrayOfBooks.[].[].isbn", separator: "_", casing: "classify", indexHandling: "index", indexFormat: "indexIjk", wanted: "ArrayOfArrayOfBooks_I_J_Isbn"},
		{path: "arrayOfArrayOfBooks.[].[].isbn", separator: ".", casing: "none", indexHandling: "index", indexFormat: "indexSprintf", wanted: "arrayOfArrayOfBooks.%d.%d.isbn"},
		{path: "arrayOfArrayOfBooks.[].isbn", separator: ".", casing: "none", indexHandling: "index", indexFormat: "indexSprintf", wanted: "arrayOfArrayOfBooks.%d.isbn"},
		{path: "mapOfArrayOfBooks.%s.coAuthors", separator: "_", casing: "classify", indexHandling: "index", indexFormat: "indexIjk", wanted: "MapOfArrayOfBooks_S_CoAuthors"},
		{path: "mapOfArrayOfBooks.%s.[].coAuthors", separator: "_", casing: "classify", indexHandling: "index", indexFormat: "indexIjk", wanted: "MapOfArrayOfBooks_S_I_CoAuthors"},
		{path: "mapOfArrayOfBooks.%s.[].coAuthors.[]", separator: "_", casing: "classify", indexHandling: "index", indexFormat: "indexIjk", wanted: "MapOfArrayOfBooks_S_I_CoAuthors_J"},
		{path: "mapOfArrayOfBooks.%s.coAuthors", separator: ".", casing: "none", indexHandling: "index", indexFormat: "indexSprintf", wanted: "mapOfArrayOfBooks.%s.coAuthors"},
		{path: "mapOfArrayOfBooks.%s.[].coAuthors", separator: ".", casing: "none", indexHandling: "index", indexFormat: "indexSprintf", wanted: "mapOfArrayOfBooks.%s.%d.coAuthors"},
		{path: "mapOfArrayOfBooks.%s.[].coAuthors.[]", separator: ".", casing: "none", indexHandling: "index", indexFormat: "indexSprintf", wanted: "mapOfArrayOfBooks.%s.%d.coAuthors.%d"},
		{path: "mapOfArrayOfBooks.%s.[].isbn", separator: "_", casing: "classify", indexHandling: "index", indexFormat: "indexIjk", wanted: "MapOfArrayOfBooks_S_I_Isbn"},
		{path: "mapOfArrayOfBooks.%s.isbn", separator: "_", casing: "classify", indexHandling: "index", indexFormat: "indexIjk", wanted: "MapOfArrayOfBooks_S_Isbn"},
		{path: "mapOfArrayOfBooks.%s.[].isbn", separator: ".", casing: "none", indexHandling: "index", indexFormat: "indexSprintf", wanted: "mapOfArrayOfBooks.%s.%d.isbn"},
		{path: "mapOfArrayOfBooks.%s.isbn", separator: ".", casing: "none", indexHandling: "index", indexFormat: "indexSprintf", wanted: "mapOfArrayOfBooks.%s.isbn"},
		{path: "mapOfBooks.%s.coAuthors", separator: "_", casing: "classify", indexHandling: "index", indexFormat: "indexIjk", wanted: "MapOfBooks_S_CoAuthors"},
		{path: "mapOfBooks.%s.coAuthors.[]", separator: "_", casing: "classify", indexHandling: "index", indexFormat: "indexIjk", wanted: "MapOfBooks_S_CoAuthors_I"},
		{path: "mapOfBooks.%s.coAuthors", separator: ".", casing: "none", indexHandling: "index", indexFormat: "indexSprintf", wanted: "mapOfBooks.%s.coAuthors"},
		{path: "mapOfBooks.%s.coAuthors.[]", separator: ".", casing: "none", indexHandling: "index", indexFormat: "indexSprintf", wanted: "mapOfBooks.%s.coAuthors.%d"},
		{path: "mapOfBooks.%s.isbn", separator: "_", casing: "classify", indexHandling: "index", indexFormat: "indexIjk", wanted: "MapOfBooks_S_Isbn"},
		{path: "mapOfBooks.%s.isbn", separator: ".", casing: "none", indexHandling: "index", indexFormat: "indexSprintf", wanted: "mapOfBooks.%s.isbn"},
	}

	for i, iw := range iws {

		pf := funcmap.NewPathFormatter(attributes.PathInfo{Path: iw.path})
		pf.WithCasing(string(iw.casing)).WithSeparator(iw.separator)
		if iw.indexFormat == "indexSprintf" {
			pf.WithSprintfFormat()
		}
		if iw.indexHandling == "suppress" {
			pf.WithSuppressCollectionsPlaceHolders()
		}

		actual := pf.FormatPath()
		require.Equal(t, iw.wanted, actual, fmt.Sprintf("[%d]", i))
	}

}
