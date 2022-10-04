package DecentralizedABE

import (
	"github.com/Nik-U/pbc"
	"hash"
	"math/big"
)

type CurveParam struct {
	p *big.Int //Order of G, N=p1*p2*p3

	Param   *pbc.Params
	Pairing *pbc.Pairing
}

func (this *CurveParam) Initialize() {
	this.p = new(big.Int)
	p1 := new(big.Int)
	p2 := new(big.Int)
	p3 := new(big.Int)
	p1.SetString("242661090146032969904098483991985908921", 10) // octal
	p2.SetString("215662396313044988944834777682074105079", 10) // octal
	p3.SetString("253493408475411572624002367871313476827", 10) // octal
	this.p.Mul(p1, p2)
	this.p.Mul(this.p, p3)
	this.Param = pbc.GenerateA1(this.p)
	this.Pairing = this.Param.NewPairing()
}

func (this *CurveParam) GetP() *big.Int {
	N := new(big.Int)
	N.Set(this.p)
	return N
}
func (this *CurveParam) GetPairing() *pbc.Pairing {
	return this.Pairing
}

func (this *CurveParam) GetNewG1() *pbc.Element {
	g := this.Pairing.NewUncheckedElement(0).Rand()
	return g
}

func (this *CurveParam) GetNewGT() *pbc.Element {
	g := this.Pairing.NewUncheckedElement(2).Rand()
	return g
}

func (this *CurveParam) GetNewZn() *pbc.Element {
	g := this.Pairing.NewUncheckedElement(3).Rand()
	return g
}

func (this *CurveParam) GetG1FromStringHash(s string, hash hash.Hash) *pbc.Element {
	g := this.Pairing.NewUncheckedElement(0).SetFromStringHash(s, hash)
	return g
}

func (this *CurveParam) GetZnFromStringHash(s string, hash hash.Hash) *pbc.Element {
	g := this.Pairing.NewUncheckedElement(3).SetFromStringHash(s, hash)
	return g
}

func (this *CurveParam) Get0FromG1() *pbc.Element {
	g := this.Pairing.NewUncheckedElement(0).Set0()
	return g
}

func (this *CurveParam) Get0FromGT() *pbc.Element {
	g := this.Pairing.NewUncheckedElement(2).Set0()
	return g
}

func (this *CurveParam) Get0FromZn() *pbc.Element {
	g := this.Pairing.NewUncheckedElement(3).Set0()
	return g
}

func (this *CurveParam) Get1FromG1() *pbc.Element {
	g := this.Pairing.NewUncheckedElement(0).Set1()
	return g
}

func (this *CurveParam) Get1FromGT() *pbc.Element {
	g := this.Pairing.NewUncheckedElement(2).Set1()
	return g
}

func (this *CurveParam) Get1FromZn() *pbc.Element {
	g := this.Pairing.NewUncheckedElement(3).Set1()
	return g
}