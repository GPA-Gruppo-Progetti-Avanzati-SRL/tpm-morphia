package schema

import (
	"fmt"
	"strings"
)

type PathFinderVisitor struct {
	Attributes []*Field
	Paths      []string
	BsonPaths  []string
}

func (pv *PathFinderVisitor) visit(f *Field) {
	pv.Attributes = append(pv.Attributes, f)
}

func (pv *PathFinderVisitor) startVisit(f *Field) {

	pv.Paths = append(pv.Paths, f.Name)
	pv.BsonPaths = append(pv.BsonPaths, f.GetTagNameValue("bson"))

	p := buildPath(pv.Paths)
	f.Paths = append(f.Paths, p)

	p1 := clearIndexPlaceHolders(p)
	if p1 != "" {
		f.Paths = append(f.Paths, p1)
	}

	bp := buildPath(pv.BsonPaths)
	f.BsonPaths = append(f.BsonPaths, bp)
	bp1 := clearIndexPlaceHolders(bp)
	if bp1 != "" {
		f.BsonPaths = append(f.BsonPaths, bp1)
	}

	/*
		if f.Paths == "" {
			f.Paths = p
		} else {
			f.Paths = strings.Join([]string{ f.Paths, p }, ";")
		}
	*/
	fmt.Printf("Path of %s of type %s = %v %v\n", f.Name, f.Typ, f.Paths, f.BsonPaths)
}

func (pv *PathFinderVisitor) endVisit(f *Field) {
	if len(pv.Paths) > 0 {
		pv.Paths = pv.Paths[:len(pv.Paths)-1]
	}

	if len(pv.BsonPaths) > 0 {
		pv.BsonPaths = pv.BsonPaths[:len(pv.BsonPaths)-1]
	}
}

func buildPath(paths []string) string {

	var b strings.Builder
	for i := range paths {
		if i > 0 {
			b.WriteRune('.')
		}
		b.WriteString(paths[i])
	}

	return b.String()
}

func clearIndexPlaceHolders(p string) string {

	if strings.HasSuffix(p, "[]") {
		return ""
	}

	// In here i consider the removal of index place holder.
	p1 := strings.ReplaceAll(p, ".[].", ".")
	// And here I check that something has been done...
	if p1 != p {
		return p1
	}

	return ""
}
