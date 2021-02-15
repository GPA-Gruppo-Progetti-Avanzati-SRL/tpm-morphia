package schema

type ListFieldsVisitor struct {
	Attributes []*Field
}

func (lv *ListFieldsVisitor) visit(f *Field) {
	lv.Attributes = append(lv.Attributes, f)
}

func (lv *ListFieldsVisitor) startVisit(f *Field) {

}

func (lv *ListFieldsVisitor) endVisit(f *Field) {

}
