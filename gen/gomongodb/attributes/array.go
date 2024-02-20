package attributes

type ArrayAttribute struct {
	GoAttributeImpl
	Item GoAttribute
}

func (a *ArrayAttribute) GoType(currentPkg string) string {
	return "[]" + a.Item.GoType(currentPkg)
}

func (s *ArrayAttribute) Visit(visitor Visitor) {
	visitor.StartVisit(s)
	if visitor.Visit(s.Item) {
		s.Item.Visit(visitor)
	}
	visitor.EndVisit(s)
}

/*func (v *ArrayAttribute) StructQualifiedName() string {
	qn := "./" + v.AttrDefinition.Typ
	return qn
}
*/
