package example5

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func FilterGoInfo() string {
	i := fmt.Sprintf("tpm_morphia query filter support generated for %s package on %s", "author", time.Now().String())
	return i
}

//----- sid of type string
//----- sid - string -  [sid]

// AndSidEqTo No Remarks
func (ca *Criteria) AndSidEqTo(p string) *Criteria {

	if p == "" {
		return ca
	}

	mName := fmt.Sprintf(SID)
	c := func() bson.E { return bson.E{Key: mName, Value: p} }
	*ca = append(*ca, c)
	return ca
}

// AndSidIsNullOrUnset No Remarks
func (ca *Criteria) AndSidIsNullOrUnset() *Criteria {

	mName := fmt.Sprintf(SID)
	c := func() bson.E { return bson.E{Key: mName, Value: nil} }
	*ca = append(*ca, c)
	return ca
}

func (ca *Criteria) AndSidIn(p []string) *Criteria {

	if len(p) == 0 {
		return ca
	}

	mName := fmt.Sprintf(SID)
	c := func() bson.E { return bson.E{Key: mName, Value: bson.D{{"$in", p}}} }
	*ca = append(*ca, c)
	return ca
}

//----- nickname of type string
//----- nickname - string -  [nickname]

// AndNicknameEqTo No Remarks
func (ca *Criteria) AndNicknameEqTo(p string) *Criteria {

	if p == "" {
		return ca
	}

	mName := fmt.Sprintf(NICKNAME)
	c := func() bson.E { return bson.E{Key: mName, Value: p} }
	*ca = append(*ca, c)
	return ca
}

// AndNicknameIsNullOrUnset No Remarks
func (ca *Criteria) AndNicknameIsNullOrUnset() *Criteria {

	mName := fmt.Sprintf(NICKNAME)
	c := func() bson.E { return bson.E{Key: mName, Value: nil} }
	*ca = append(*ca, c)
	return ca
}

func (ca *Criteria) AndNicknameIn(p []string) *Criteria {

	if len(p) == 0 {
		return ca
	}

	mName := fmt.Sprintf(NICKNAME)
	c := func() bson.E { return bson.E{Key: mName, Value: bson.D{{"$in", p}}} }
	*ca = append(*ca, c)
	return ca
}

//----- remoteaddr of type string
//----- remoteaddr - string -  [remoteaddr]

// AndRemoteaddrEqTo No Remarks
func (ca *Criteria) AndRemoteaddrEqTo(p string) *Criteria {

	if p == "" {
		return ca
	}

	mName := fmt.Sprintf(REMOTEADDR)
	c := func() bson.E { return bson.E{Key: mName, Value: p} }
	*ca = append(*ca, c)
	return ca
}

// AndRemoteaddrIsNullOrUnset No Remarks
func (ca *Criteria) AndRemoteaddrIsNullOrUnset() *Criteria {

	mName := fmt.Sprintf(REMOTEADDR)
	c := func() bson.E { return bson.E{Key: mName, Value: nil} }
	*ca = append(*ca, c)
	return ca
}

func (ca *Criteria) AndRemoteaddrIn(p []string) *Criteria {

	if len(p) == 0 {
		return ca
	}

	mName := fmt.Sprintf(REMOTEADDR)
	c := func() bson.E { return bson.E{Key: mName, Value: bson.D{{"$in", p}}} }
	*ca = append(*ca, c)
	return ca
}
