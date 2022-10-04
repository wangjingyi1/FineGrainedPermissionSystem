package DecentralizedABE

import (
	"github.com/Nik-U/pbc"
)

/* Public Key Structure */
type APK struct {
	Gy *pbc.Element `field:"0"` //G^y, y from Zp
}

func (p *APK) Initialize(gy *pbc.Element) {
	p.Gy = gy
}

func (p *APK) getGy() *pbc.Element {
	return p.Gy.NewFieldElement().Set(p.Gy)
}

type ASK struct {
	Y *pbc.Element `field:"3"`
}

func (s *ASK) Initialize(y *pbc.Element) {
	s.Y = y
}

func (s *ASK) getY() *pbc.Element {
	return s.Y.NewFieldElement().Set(s.Y)
}

type OPKPart struct {
	OPK    *pbc.Element            `field:"2"` //part of org's EGGAlpha
	APKMap map[string]*pbc.Element `field:"0"` //part of org attrs' gy
}
type OSKPart struct {
	AlphaPart   *pbc.Element `field:"3"`   //part of org's Alpha
	ASKMap      map[string]*ASKPart        //part of org attrs' y
	F           []*pbc.Element `field:"3"` //for shamir's share
	N           int
	T           int
	OthersShare []*pbc.Element `field:"3"` //for some special time
	OSK         *pbc.Element   `field:"3"` //mul shares
	GOSK        *pbc.Element   `field:"0"`
}
type ASKPart struct {
	F           []*pbc.Element `field:"3"` //for shamir's share
	OthersShare []*pbc.Element `field:"3"` //for some special time
	YPart       *pbc.Element   `field:"3"`
	ASK         *pbc.Element   `field:"3"` //mul shares
}
