package example5

const (
	SID        = "_id"
	NICKNAME   = "nickname"
	REMOTEADDR = "remoteaddr"
	FLAGS      = "flags"
)

type Session struct {
	Sid        string `json:"-" bson:"_id,omitempty"`
	Nickname   string `json:"nickname,omitempty" bson:"nickname,omitempty"`
	Remoteaddr string `json:"remoteaddr,omitempty" bson:"remoteaddr,omitempty"`
	Flags      string `json:"flags,omitempty" bson:"flags,omitempty"`
}

func (s Session) IsZero() bool {
	/*
	   if s.Sid == "" {
	       return false
	   }
	   if s.Nickname == "" {
	       return false
	   }
	   if s.Remoteaddr == "" {
	       return false
	   }
	   if s.Flags == "" {
	       return false
	   }
	       return true
	*/

	return s.Sid == "" && s.Nickname == "" && s.Remoteaddr == "" && s.Flags == ""
}
