package attributes

type MapAttribute struct {
	GoAttributeImpl
	Item GoAttribute
}

func (a *MapAttribute) GoType(currentPkg string) string {
	return "map[string]" + a.Item.GoType(currentPkg)
}

func (s *MapAttribute) Visit(visitor Visitor) {
	visitor.StartVisit(s)
	if visitor.Visit(s.Item) {
		s.Item.Visit(visitor)
	}
	visitor.EndVisit(s)
}

/*func (v *MapAttribute) StructQualifiedName() string {
	qn := "./" + v.AttrDefinition.Typ
	return qn
}
*/
