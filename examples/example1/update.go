package example1

import "go.mongodb.org/mongo-driver/bson"

var emptyUpdate = bson.D{}

type UpdateOperator string

const (
	Set   UpdateOperator = "$set"
	Unset UpdateOperator = "$unset"
)

type Update func() bson.E

type UpdateDocument struct {
	Ops map[UpdateOperator]*Updates
}

func (uo *UpdateDocument) Set() *Updates {
	return uo.op(Set)
}

func (uo *UpdateDocument) op(op UpdateOperator) *Updates {

	if len(uo.Ops) == 0 {
		uo.Ops = make(map[UpdateOperator]*Updates)
	}

	if u, ok := uo.Ops[op]; ok {
		return u
	}

	uopi := &Updates{operator: "$set"}
	uo.Ops[op] = uopi

	return uopi
}

func (uo *UpdateDocument) Build() bson.D {
	if len(uo.Ops) == 0 {
		return emptyFilter
	}

	docAll := bson.D{}
	for _, cas := range uo.Ops {
		doc := bson.D{}
		for _, c := range cas.updates {
			if u := c(); u.Key != "" {
				doc = append(doc, u)
			}
		}
		if len(doc) > 0 {
			opu := bson.E{Key: cas.operator, Value: doc}
			docAll = append(docAll, opu)
		}
	}

	return docAll
}

type Updates struct {
	operator string
	updates  []Update
}

func (ui *Updates) Add(fu Update) {
	ui.updates = append(ui.updates, fu)
}
