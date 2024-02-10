package example5

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func UpdateMethodsGoInfo() string {
	i := fmt.Sprintf("tpm_morphia query filter support generated for %s package on %s", "author", time.Now().String())
	return i
}

type UnsetMode int64

const (
	UnSpecified     UnsetMode = 0
	KeepCurrent               = 1
	UnsetData                 = 2
	SetData2Default           = 3
)

type UnsetOption func(uopt *UnsetOptions)

type UnsetOptions struct {
	DefaultMode UnsetMode
	Sid         UnsetMode
	Nickname    UnsetMode
	Remoteaddr  UnsetMode
	Flags       UnsetMode
}

func (uo *UnsetOptions) ResolveUnsetMode(um UnsetMode) UnsetMode {
	if um == UnSpecified {
		um = uo.DefaultMode
	}

	return um
}

func WithDefaultUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.DefaultMode = m
	}
}
func WithSidUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Sid = m
	}
}
func WithNicknameUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Nickname = m
	}
}
func WithRemoteaddrUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Remoteaddr = m
	}
}
func WithFlagsUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Flags = m
	}
}

type UpdateOption func(ud *UpdateDocument)
type UpdateOptions []UpdateOption

// GetUpdateDocument convenience method to create an update document from single updates instead of a whole object
func GetUpdateDocumentFromOptions(opts ...UpdateOption) UpdateDocument {
	ud := UpdateDocument{}
	for _, o := range opts {
		o(&ud)
	}

	return ud
}

// GetUpdateDocument
// Convenience method to create an Update Document from the values of the top fields of the object. The convenience is in the handling
// the unset because if I pass an empty struct to the update it generates an empty object anyway in the db. Handling the unset eliminates
// the issue and delete an existing value without creating an empty struct.
func GetUpdateDocument(obj *Session, opts ...UnsetOption) UpdateDocument {

	uo := &UnsetOptions{DefaultMode: KeepCurrent}
	for _, o := range opts {
		o(uo)
	}

	ud := UpdateDocument{}
	ud.setOrUnsetSid(obj.Sid, uo.ResolveUnsetMode(uo.Sid))
	ud.setOrUnsetNickname(obj.Nickname, uo.ResolveUnsetMode(uo.Nickname))
	ud.setOrUnsetRemoteaddr(obj.Remoteaddr, uo.ResolveUnsetMode(uo.Remoteaddr))
	ud.setOrUnsetFlags(obj.Flags, uo.ResolveUnsetMode(uo.Flags))

	return ud
}

//----- sid - string -  [sid]

// SetSid No Remarks
func (ud *UpdateDocument) SetSid(p string) *UpdateDocument {
	mName := fmt.Sprintf(SID)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetSid No Remarks
func (ud *UpdateDocument) UnsetSid() *UpdateDocument {
	mName := fmt.Sprintf(SID)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetSid No Remarks
func (ud *UpdateDocument) setOrUnsetSid(p string, um UnsetMode) {
	if p != "" {
		ud.SetSid(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetSid()
		case SetData2Default:
			ud.UnsetSid()
		}
	}
}

func UpdateWithSid(p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetSid(p)
		} else {
			ud.UnsetSid()
		}
	}
}

//----- nickname - string -  [nickname]

// SetNickname No Remarks
func (ud *UpdateDocument) SetNickname(p string) *UpdateDocument {
	mName := fmt.Sprintf(NICKNAME)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetNickname No Remarks
func (ud *UpdateDocument) UnsetNickname() *UpdateDocument {
	mName := fmt.Sprintf(NICKNAME)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetNickname No Remarks
func (ud *UpdateDocument) setOrUnsetNickname(p string, um UnsetMode) {
	if p != "" {
		ud.SetNickname(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetNickname()
		case SetData2Default:
			ud.UnsetNickname()
		}
	}
}

func UpdateWithNickname(p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetNickname(p)
		} else {
			ud.UnsetNickname()
		}
	}
}

//----- remoteaddr - string -  [remoteaddr]

// SetRemoteaddr No Remarks
func (ud *UpdateDocument) SetRemoteaddr(p string) *UpdateDocument {
	mName := fmt.Sprintf(REMOTEADDR)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetRemoteaddr No Remarks
func (ud *UpdateDocument) UnsetRemoteaddr() *UpdateDocument {
	mName := fmt.Sprintf(REMOTEADDR)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetRemoteaddr No Remarks
func (ud *UpdateDocument) setOrUnsetRemoteaddr(p string, um UnsetMode) {
	if p != "" {
		ud.SetRemoteaddr(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetRemoteaddr()
		case SetData2Default:
			ud.UnsetRemoteaddr()
		}
	}
}

func UpdateWithRemoteaddr(p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetRemoteaddr(p)
		} else {
			ud.UnsetRemoteaddr()
		}
	}
}

//----- flags - string -  [flags]

// SetFlags No Remarks
func (ud *UpdateDocument) SetFlags(p string) *UpdateDocument {
	mName := fmt.Sprintf(FLAGS)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetFlags No Remarks
func (ud *UpdateDocument) UnsetFlags() *UpdateDocument {
	mName := fmt.Sprintf(FLAGS)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetFlags No Remarks
func (ud *UpdateDocument) setOrUnsetFlags(p string, um UnsetMode) {
	if p != "" {
		ud.SetFlags(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetFlags()
		case SetData2Default:
			ud.UnsetFlags()
		}
	}
}

func UpdateWithFlags(p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetFlags(p)
		} else {
			ud.UnsetFlags()
		}
	}
}
