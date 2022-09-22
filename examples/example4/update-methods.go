package example4

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
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
	DefaultMode   UnsetMode
	OId           UnsetMode
	Ndg           UnsetMode
	CodiceFiscale UnsetMode
	PartitaIVA    UnsetMode
	Natura        UnsetMode
	Stato         UnsetMode
	Indirizzi     UnsetMode
	Legati        UnsetMode
	Leganti       UnsetMode
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
func WithOIdUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.OId = m
	}
}
func WithNdgUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Ndg = m
	}
}
func WithCodiceFiscaleUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.CodiceFiscale = m
	}
}
func WithPartitaIVAUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.PartitaIVA = m
	}
}
func WithNaturaUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Natura = m
	}
}
func WithStatoUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Stato = m
	}
}
func WithIndirizziUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Indirizzi = m
	}
}
func WithLegatiUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Legati = m
	}
}
func WithLegantiUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Leganti = m
	}
}

type UpdateOption func(ud *UpdateDocument)
type UpdateOptions []UpdateOption

// GetUpdateDocument convenience method to create an update document from single updates instead of a whole object
func (uopts UpdateOptions) GetUpdateDocument(opts ...UpdateOption) UpdateDocument {
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
func GetUpdateDocument(obj *Cliente, opts ...UnsetOption) UpdateDocument {

	uo := &UnsetOptions{DefaultMode: KeepCurrent}
	for _, o := range opts {
		o(uo)
	}

	ud := UpdateDocument{}
	ud.setOrUnsetNdg(obj.Ndg, uo.ResolveUnsetMode(uo.Ndg))
	ud.setOrUnsetCodiceFiscale(obj.CodiceFiscale, uo.ResolveUnsetMode(uo.CodiceFiscale))
	ud.setOrUnsetPartitaIVA(obj.PartitaIVA, uo.ResolveUnsetMode(uo.PartitaIVA))
	ud.setOrUnsetNatura(obj.Natura, uo.ResolveUnsetMode(uo.Natura))
	ud.setOrUnsetStato(obj.Stato, uo.ResolveUnsetMode(uo.Stato))
	// if len(obj.Indirizzi) > 0 {
	//   ud.SetIndirizzi ( obj.Indirizzi)
	// } else {
	ud.setOrUnsetIndirizzi(obj.Indirizzi, uo.ResolveUnsetMode(uo.Indirizzi))
	// }
	// if len(obj.Legati) > 0 {
	//   ud.SetLegati ( obj.Legati)
	// } else {
	ud.setOrUnsetLegati(obj.Legati, uo.ResolveUnsetMode(uo.Legati))
	// }
	// if len(obj.Leganti) > 0 {
	//   ud.SetLeganti ( obj.Leganti)
	// } else {
	ud.setOrUnsetLeganti(obj.Leganti, uo.ResolveUnsetMode(uo.Leganti))
	// }

	return ud
}

//----- oId - object-id -  [oId]

// SetOId No Remarks
func (ud *UpdateDocument) SetOId(p primitive.ObjectID) *UpdateDocument {
	mName := fmt.Sprintf(OID)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetOId No Remarks
func (ud *UpdateDocument) UnsetOId() *UpdateDocument {
	mName := fmt.Sprintf(OID)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetOId No Remarks
func (ud *UpdateDocument) setOrUnsetOId(p primitive.ObjectID, um UnsetMode) {
	if !p.IsZero() {
		ud.SetOId(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetOId()
		case SetData2Default:
			ud.UnsetOId()
		}
	}
}

//----- ndg - string -  [ndg]

// SetNdg No Remarks
func (ud *UpdateDocument) SetNdg(p string) *UpdateDocument {
	mName := fmt.Sprintf(NDG)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetNdg No Remarks
func (ud *UpdateDocument) UnsetNdg() *UpdateDocument {
	mName := fmt.Sprintf(NDG)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetNdg No Remarks
func (ud *UpdateDocument) setOrUnsetNdg(p string, um UnsetMode) {
	if p != "" {
		ud.SetNdg(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetNdg()
		case SetData2Default:
			ud.UnsetNdg()
		}
	}
}

func UpdateWithNdg(p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetNdg(p)
		} else {
			ud.UnsetNdg()
		}
	}
}

//----- codiceFiscale - string -  [codiceFiscale]

// SetCodiceFiscale No Remarks
func (ud *UpdateDocument) SetCodiceFiscale(p string) *UpdateDocument {
	mName := fmt.Sprintf(CODICEFISCALE)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetCodiceFiscale No Remarks
func (ud *UpdateDocument) UnsetCodiceFiscale() *UpdateDocument {
	mName := fmt.Sprintf(CODICEFISCALE)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetCodiceFiscale No Remarks
func (ud *UpdateDocument) setOrUnsetCodiceFiscale(p string, um UnsetMode) {
	if p != "" {
		ud.SetCodiceFiscale(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetCodiceFiscale()
		case SetData2Default:
			ud.UnsetCodiceFiscale()
		}
	}
}

func UpdateWithCodiceFiscale(p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetCodiceFiscale(p)
		} else {
			ud.UnsetCodiceFiscale()
		}
	}
}

//----- partitaIVA - string -  [partitaIVA]

// SetPartitaIVA No Remarks
func (ud *UpdateDocument) SetPartitaIVA(p string) *UpdateDocument {
	mName := fmt.Sprintf(PARTITAIVA)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetPartitaIVA No Remarks
func (ud *UpdateDocument) UnsetPartitaIVA() *UpdateDocument {
	mName := fmt.Sprintf(PARTITAIVA)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetPartitaIVA No Remarks
func (ud *UpdateDocument) setOrUnsetPartitaIVA(p string, um UnsetMode) {
	if p != "" {
		ud.SetPartitaIVA(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetPartitaIVA()
		case SetData2Default:
			ud.UnsetPartitaIVA()
		}
	}
}

func UpdateWithPartitaIVA(p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetPartitaIVA(p)
		} else {
			ud.UnsetPartitaIVA()
		}
	}
}

//----- natura - string -  [natura]

// SetNatura No Remarks
func (ud *UpdateDocument) SetNatura(p string) *UpdateDocument {
	mName := fmt.Sprintf(NATURA)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetNatura No Remarks
func (ud *UpdateDocument) UnsetNatura() *UpdateDocument {
	mName := fmt.Sprintf(NATURA)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetNatura No Remarks
func (ud *UpdateDocument) setOrUnsetNatura(p string, um UnsetMode) {
	if p != "" {
		ud.SetNatura(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetNatura()
		case SetData2Default:
			ud.UnsetNatura()
		}
	}
}

func UpdateWithNatura(p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetNatura(p)
		} else {
			ud.UnsetNatura()
		}
	}
}

//----- stato - string -  [stato]

// SetStato No Remarks
func (ud *UpdateDocument) SetStato(p string) *UpdateDocument {
	mName := fmt.Sprintf(STATO)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetStato No Remarks
func (ud *UpdateDocument) UnsetStato() *UpdateDocument {
	mName := fmt.Sprintf(STATO)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetStato No Remarks
func (ud *UpdateDocument) setOrUnsetStato(p string, um UnsetMode) {
	if p != "" {
		ud.SetStato(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetStato()
		case SetData2Default:
			ud.UnsetStato()
		}
	}
}

func UpdateWithStato(p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetStato(p)
		} else {
			ud.UnsetStato()
		}
	}
}

// ----- indirizzi - map -  [indirizzi]
// SetIndirizzi No Remarks
func (ud *UpdateDocument) SetIndirizzi(p map[string]Indirizzo) *UpdateDocument {
	mName := fmt.Sprintf(INDIRIZZI)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetIndirizzi No Remarks
func (ud *UpdateDocument) UnsetIndirizzi() *UpdateDocument {
	mName := fmt.Sprintf(INDIRIZZI)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetIndirizzi No Remarks - here2
func (ud *UpdateDocument) setOrUnsetIndirizzi(p map[string]Indirizzo, um UnsetMode) {

	//----- map\n

	if len(p) > 0 {
		ud.SetIndirizzi(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetIndirizzi()
		case SetData2Default:
			ud.UnsetIndirizzi()
		}
	}
}

func UpdateWithIndirizzi(p map[string]Indirizzo) UpdateOption {
	return func(ud *UpdateDocument) {
		if len(p) > 0 {
			ud.SetIndirizzi(p)
		} else {
			ud.UnsetIndirizzi()
		}
	}
}

// ----- %s - struct - Indirizzo [indirizzi.%s]
// SetIndirizziS No Remarks
func (ud *UpdateDocument) SetIndirizziS(keyS string, p Indirizzo) *UpdateDocument {
	mName := fmt.Sprintf(INDIRIZZI_S, keyS)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetIndirizziS No Remarks
func (ud *UpdateDocument) UnsetIndirizziS(keyS string) *UpdateDocument {
	mName := fmt.Sprintf(INDIRIZZI_S, keyS)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetIndirizziS No Remarks
func (ud *UpdateDocument) setOrUnsetIndirizziS(keyS string, p Indirizzo, um UnsetMode) {
	if !p.IsZero() {
		ud.SetIndirizziS(keyS, p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetIndirizziS(keyS)
		case SetData2Default:
			ud.UnsetIndirizziS(keyS)
		}
	}
}

//----- indirizzo - string -  [indirizzi.%s.indirizzo]

// SetIndirizziSIndirizzo No Remarks
func (ud *UpdateDocument) SetIndirizziSIndirizzo(keyS string, p string) *UpdateDocument {
	mName := fmt.Sprintf(INDIRIZZI_S_INDIRIZZO, keyS)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetIndirizziSIndirizzo No Remarks
func (ud *UpdateDocument) UnsetIndirizziSIndirizzo(keyS string) *UpdateDocument {
	mName := fmt.Sprintf(INDIRIZZI_S_INDIRIZZO, keyS)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetIndirizziSIndirizzo No Remarks
func (ud *UpdateDocument) setOrUnsetIndirizziSIndirizzo(keyS string, p string, um UnsetMode) {
	if p != "" {
		ud.SetIndirizziSIndirizzo(keyS, p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetIndirizziSIndirizzo(keyS)
		case SetData2Default:
			ud.UnsetIndirizziSIndirizzo(keyS)
		}
	}
}

func UpdateWithIndirizziSIndirizzo(keyS string, p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetIndirizziSIndirizzo(keyS, p)
		} else {
			ud.UnsetIndirizziSIndirizzo(keyS)
		}
	}
}

//----- cap - string -  [indirizzi.%s.cap]

// SetIndirizziSCap No Remarks
func (ud *UpdateDocument) SetIndirizziSCap(keyS string, p string) *UpdateDocument {
	mName := fmt.Sprintf(INDIRIZZI_S_CAP, keyS)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetIndirizziSCap No Remarks
func (ud *UpdateDocument) UnsetIndirizziSCap(keyS string) *UpdateDocument {
	mName := fmt.Sprintf(INDIRIZZI_S_CAP, keyS)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetIndirizziSCap No Remarks
func (ud *UpdateDocument) setOrUnsetIndirizziSCap(keyS string, p string, um UnsetMode) {
	if p != "" {
		ud.SetIndirizziSCap(keyS, p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetIndirizziSCap(keyS)
		case SetData2Default:
			ud.UnsetIndirizziSCap(keyS)
		}
	}
}

func UpdateWithIndirizziSCap(keyS string, p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetIndirizziSCap(keyS, p)
		} else {
			ud.UnsetIndirizziSCap(keyS)
		}
	}
}

//----- localita - string -  [indirizzi.%s.localita]

// SetIndirizziSLocalita No Remarks
func (ud *UpdateDocument) SetIndirizziSLocalita(keyS string, p string) *UpdateDocument {
	mName := fmt.Sprintf(INDIRIZZI_S_LOCALITA, keyS)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetIndirizziSLocalita No Remarks
func (ud *UpdateDocument) UnsetIndirizziSLocalita(keyS string) *UpdateDocument {
	mName := fmt.Sprintf(INDIRIZZI_S_LOCALITA, keyS)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetIndirizziSLocalita No Remarks
func (ud *UpdateDocument) setOrUnsetIndirizziSLocalita(keyS string, p string, um UnsetMode) {
	if p != "" {
		ud.SetIndirizziSLocalita(keyS, p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetIndirizziSLocalita(keyS)
		case SetData2Default:
			ud.UnsetIndirizziSLocalita(keyS)
		}
	}
}

func UpdateWithIndirizziSLocalita(keyS string, p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetIndirizziSLocalita(keyS, p)
		} else {
			ud.UnsetIndirizziSLocalita(keyS)
		}
	}
}

//----- provincia - string -  [indirizzi.%s.provincia]

// SetIndirizziSProvincia No Remarks
func (ud *UpdateDocument) SetIndirizziSProvincia(keyS string, p string) *UpdateDocument {
	mName := fmt.Sprintf(INDIRIZZI_S_PROVINCIA, keyS)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetIndirizziSProvincia No Remarks
func (ud *UpdateDocument) UnsetIndirizziSProvincia(keyS string) *UpdateDocument {
	mName := fmt.Sprintf(INDIRIZZI_S_PROVINCIA, keyS)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetIndirizziSProvincia No Remarks
func (ud *UpdateDocument) setOrUnsetIndirizziSProvincia(keyS string, p string, um UnsetMode) {
	if p != "" {
		ud.SetIndirizziSProvincia(keyS, p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetIndirizziSProvincia(keyS)
		case SetData2Default:
			ud.UnsetIndirizziSProvincia(keyS)
		}
	}
}

func UpdateWithIndirizziSProvincia(keyS string, p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetIndirizziSProvincia(keyS, p)
		} else {
			ud.UnsetIndirizziSProvincia(keyS)
		}
	}
}

//----- nazione - string -  [indirizzi.%s.nazione]

// SetIndirizziSNazione No Remarks
func (ud *UpdateDocument) SetIndirizziSNazione(keyS string, p string) *UpdateDocument {
	mName := fmt.Sprintf(INDIRIZZI_S_NAZIONE, keyS)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetIndirizziSNazione No Remarks
func (ud *UpdateDocument) UnsetIndirizziSNazione(keyS string) *UpdateDocument {
	mName := fmt.Sprintf(INDIRIZZI_S_NAZIONE, keyS)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetIndirizziSNazione No Remarks
func (ud *UpdateDocument) setOrUnsetIndirizziSNazione(keyS string, p string, um UnsetMode) {
	if p != "" {
		ud.SetIndirizziSNazione(keyS, p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetIndirizziSNazione(keyS)
		case SetData2Default:
			ud.UnsetIndirizziSNazione(keyS)
		}
	}
}

func UpdateWithIndirizziSNazione(keyS string, p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetIndirizziSNazione(keyS, p)
		} else {
			ud.UnsetIndirizziSNazione(keyS)
		}
	}
}

// ----- legati - array -  [legati]
// SetLegati No Remarks
func (ud *UpdateDocument) SetLegati(p []Legame) *UpdateDocument {
	mName := fmt.Sprintf(LEGATI)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetLegati No Remarks
func (ud *UpdateDocument) UnsetLegati() *UpdateDocument {
	mName := fmt.Sprintf(LEGATI)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetLegati No Remarks - here2
func (ud *UpdateDocument) setOrUnsetLegati(p []Legame, um UnsetMode) {

	//----- array\n

	if len(p) > 0 {
		ud.SetLegati(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetLegati()
		case SetData2Default:
			ud.UnsetLegati()
		}
	}
}

func UpdateWithLegati(p []Legame) UpdateOption {
	return func(ud *UpdateDocument) {
		if len(p) > 0 {
			ud.SetLegati(p)
		} else {
			ud.UnsetLegati()
		}
	}
}

// ----- [] - struct - Legame [legati.[]]
// SetLegatiI No Remarks
func (ud *UpdateDocument) SetLegatiI(ndxI int, p Legame) *UpdateDocument {
	mName := fmt.Sprintf(LEGATI_I, ndxI)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetLegatiI No Remarks
func (ud *UpdateDocument) UnsetLegatiI(ndxI int) *UpdateDocument {
	mName := fmt.Sprintf(LEGATI_I, ndxI)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetLegatiI No Remarks
func (ud *UpdateDocument) setOrUnsetLegatiI(ndxI int, p Legame, um UnsetMode) {
	if !p.IsZero() {
		ud.SetLegatiI(ndxI, p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetLegatiI(ndxI)
		case SetData2Default:
			ud.UnsetLegatiI(ndxI)
		}
	}
}

//----- ndg - string -  [legati.[].ndg legati.ndg]

// SetLegatiINdg No Remarks
func (ud *UpdateDocument) SetLegatiINdg(ndxI int, p string) *UpdateDocument {
	mName := fmt.Sprintf(LEGATI_I_NDG, ndxI)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetLegatiINdg No Remarks
func (ud *UpdateDocument) UnsetLegatiINdg(ndxI int) *UpdateDocument {
	mName := fmt.Sprintf(LEGATI_I_NDG, ndxI)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetLegatiINdg No Remarks
func (ud *UpdateDocument) setOrUnsetLegatiINdg(ndxI int, p string, um UnsetMode) {
	if p != "" {
		ud.SetLegatiINdg(ndxI, p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetLegatiINdg(ndxI)
		case SetData2Default:
			ud.UnsetLegatiINdg(ndxI)
		}
	}
}

func UpdateWithLegatiINdg(ndxI int, p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetLegatiINdg(ndxI, p)
		} else {
			ud.UnsetLegatiINdg(ndxI)
		}
	}
}

//----- cognome - string -  [legati.[].cognome legati.cognome]

// SetLegatiICognome No Remarks
func (ud *UpdateDocument) SetLegatiICognome(ndxI int, p string) *UpdateDocument {
	mName := fmt.Sprintf(LEGATI_I_COGNOME, ndxI)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetLegatiICognome No Remarks
func (ud *UpdateDocument) UnsetLegatiICognome(ndxI int) *UpdateDocument {
	mName := fmt.Sprintf(LEGATI_I_COGNOME, ndxI)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetLegatiICognome No Remarks
func (ud *UpdateDocument) setOrUnsetLegatiICognome(ndxI int, p string, um UnsetMode) {
	if p != "" {
		ud.SetLegatiICognome(ndxI, p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetLegatiICognome(ndxI)
		case SetData2Default:
			ud.UnsetLegatiICognome(ndxI)
		}
	}
}

func UpdateWithLegatiICognome(ndxI int, p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetLegatiICognome(ndxI, p)
		} else {
			ud.UnsetLegatiICognome(ndxI)
		}
	}
}

//----- nome - string -  [legati.[].nome legati.nome]

// SetLegatiINome No Remarks
func (ud *UpdateDocument) SetLegatiINome(ndxI int, p string) *UpdateDocument {
	mName := fmt.Sprintf(LEGATI_I_NOME, ndxI)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetLegatiINome No Remarks
func (ud *UpdateDocument) UnsetLegatiINome(ndxI int) *UpdateDocument {
	mName := fmt.Sprintf(LEGATI_I_NOME, ndxI)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetLegatiINome No Remarks
func (ud *UpdateDocument) setOrUnsetLegatiINome(ndxI int, p string, um UnsetMode) {
	if p != "" {
		ud.SetLegatiINome(ndxI, p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetLegatiINome(ndxI)
		case SetData2Default:
			ud.UnsetLegatiINome(ndxI)
		}
	}
}

func UpdateWithLegatiINome(ndxI int, p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetLegatiINome(ndxI, p)
		} else {
			ud.UnsetLegatiINome(ndxI)
		}
	}
}

//----- codiceFiscale - string -  [legati.[].codiceFiscale legati.codiceFiscale]

// SetLegatiICodiceFiscale No Remarks
func (ud *UpdateDocument) SetLegatiICodiceFiscale(ndxI int, p string) *UpdateDocument {
	mName := fmt.Sprintf(LEGATI_I_CODICEFISCALE, ndxI)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetLegatiICodiceFiscale No Remarks
func (ud *UpdateDocument) UnsetLegatiICodiceFiscale(ndxI int) *UpdateDocument {
	mName := fmt.Sprintf(LEGATI_I_CODICEFISCALE, ndxI)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetLegatiICodiceFiscale No Remarks
func (ud *UpdateDocument) setOrUnsetLegatiICodiceFiscale(ndxI int, p string, um UnsetMode) {
	if p != "" {
		ud.SetLegatiICodiceFiscale(ndxI, p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetLegatiICodiceFiscale(ndxI)
		case SetData2Default:
			ud.UnsetLegatiICodiceFiscale(ndxI)
		}
	}
}

func UpdateWithLegatiICodiceFiscale(ndxI int, p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetLegatiICodiceFiscale(ndxI, p)
		} else {
			ud.UnsetLegatiICodiceFiscale(ndxI)
		}
	}
}

//----- partitaIVA - string -  [legati.[].partitaIVA legati.partitaIVA]

// SetLegatiIPartitaIVA No Remarks
func (ud *UpdateDocument) SetLegatiIPartitaIVA(ndxI int, p string) *UpdateDocument {
	mName := fmt.Sprintf(LEGATI_I_PARTITAIVA, ndxI)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetLegatiIPartitaIVA No Remarks
func (ud *UpdateDocument) UnsetLegatiIPartitaIVA(ndxI int) *UpdateDocument {
	mName := fmt.Sprintf(LEGATI_I_PARTITAIVA, ndxI)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetLegatiIPartitaIVA No Remarks
func (ud *UpdateDocument) setOrUnsetLegatiIPartitaIVA(ndxI int, p string, um UnsetMode) {
	if p != "" {
		ud.SetLegatiIPartitaIVA(ndxI, p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetLegatiIPartitaIVA(ndxI)
		case SetData2Default:
			ud.UnsetLegatiIPartitaIVA(ndxI)
		}
	}
}

func UpdateWithLegatiIPartitaIVA(ndxI int, p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetLegatiIPartitaIVA(ndxI, p)
		} else {
			ud.UnsetLegatiIPartitaIVA(ndxI)
		}
	}
}

//----- natura - string -  [legati.[].natura legati.natura]

// SetLegatiINatura No Remarks
func (ud *UpdateDocument) SetLegatiINatura(ndxI int, p string) *UpdateDocument {
	mName := fmt.Sprintf(LEGATI_I_NATURA, ndxI)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetLegatiINatura No Remarks
func (ud *UpdateDocument) UnsetLegatiINatura(ndxI int) *UpdateDocument {
	mName := fmt.Sprintf(LEGATI_I_NATURA, ndxI)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetLegatiINatura No Remarks
func (ud *UpdateDocument) setOrUnsetLegatiINatura(ndxI int, p string, um UnsetMode) {
	if p != "" {
		ud.SetLegatiINatura(ndxI, p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetLegatiINatura(ndxI)
		case SetData2Default:
			ud.UnsetLegatiINatura(ndxI)
		}
	}
}

func UpdateWithLegatiINatura(ndxI int, p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetLegatiINatura(ndxI, p)
		} else {
			ud.UnsetLegatiINatura(ndxI)
		}
	}
}

// ----- leganti - array -  [leganti]
// SetLeganti No Remarks
func (ud *UpdateDocument) SetLeganti(p []Legame) *UpdateDocument {
	mName := fmt.Sprintf(LEGANTI)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetLeganti No Remarks
func (ud *UpdateDocument) UnsetLeganti() *UpdateDocument {
	mName := fmt.Sprintf(LEGANTI)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetLeganti No Remarks - here2
func (ud *UpdateDocument) setOrUnsetLeganti(p []Legame, um UnsetMode) {

	//----- array\n

	if len(p) > 0 {
		ud.SetLeganti(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetLeganti()
		case SetData2Default:
			ud.UnsetLeganti()
		}
	}
}

func UpdateWithLeganti(p []Legame) UpdateOption {
	return func(ud *UpdateDocument) {
		if len(p) > 0 {
			ud.SetLeganti(p)
		} else {
			ud.UnsetLeganti()
		}
	}
}

// ----- [] - ref-struct -  [leganti.[]]
// SetLegantiI No Remarks
func (ud *UpdateDocument) SetLegantiI(ndxI int, p Legame) *UpdateDocument {
	mName := fmt.Sprintf(LEGANTI_I, ndxI)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetLegantiI No Remarks
func (ud *UpdateDocument) UnsetLegantiI(ndxI int) *UpdateDocument {
	mName := fmt.Sprintf(LEGANTI_I, ndxI)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetLegantiI No Remarks
func (ud *UpdateDocument) setOrUnsetLegantiI(ndxI int, p Legame, um UnsetMode) {
	if !p.IsZero() {
		ud.SetLegantiI(ndxI, p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetLegantiI(ndxI)
		case SetData2Default:
			ud.UnsetLegantiI(ndxI)
		}
	}
}
