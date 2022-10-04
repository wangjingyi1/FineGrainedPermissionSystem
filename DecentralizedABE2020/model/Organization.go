package DecentralizedABE

import (
	"fmt"
	"github.com/Nik-U/pbc"
)

type Org struct {
	APKMap       map[string]*APK
	EGGAlpha     *pbc.Element `field:"2"`
	Name         string
	N            int                                 //总用户个数
	T            int                                 //门限阈值
	UserName2GID map[string]*pbc.Element `field:"3"` //用户的名称
}

func (o *Org) GetPK() *pbc.Element {
	return o.EGGAlpha
}

func (o *Org) GetAPKMap() map[string]*APK {
	return o.APKMap
}

//生成组织公钥
func (o *Org) GenerateOPK(names []string, pks []*pbc.Element, d *DABE) error {
	if len(pks) != o.T || len(names) != o.T {
		return fmt.Errorf("pks or names isn't eq t")
	}

	eGGAlpha := d.CurveParam.Get1FromGT()
	for i := 0; i < o.T; i++ {
		up := d.CurveParam.Get1FromZn()
		for j := 0; j < o.T; j++ {
			if i == j {
				continue
			}
			temp := d.CurveParam.Get0FromZn().Sub(o.UserName2GID[names[j]], o.UserName2GID[names[i]])
			temp = d.CurveParam.GetNewZn().Div(o.UserName2GID[names[j]], temp)
			up.ThenMul(temp)
		}
		eGGAlpha.ThenMul(d.CurveParam.Get0FromGT().PowZn(pks[i], up))
	}
	o.EGGAlpha = eGGAlpha
	return nil
}

//生成属性
func (o *Org) GenerateNewAttr(names []string, apks []*pbc.Element, attr string, d *DABE) error {
	if len(apks) != o.T || len(names) != o.T {
		return fmt.Errorf("pks or names isn't eq t")
	}
	if o.APKMap[attr] != nil {
		return fmt.Errorf("already has this attr")
	}

	gY := d.CurveParam.Get1FromG1()
	for i := 0; i < o.T; i++ {
		up := d.CurveParam.Get1FromZn()
		for j := 0; j < o.T; j++ {
			if i == j {
				continue
			}
			di := d.CurveParam.Get0FromZn().Sub(o.UserName2GID[names[j]], o.UserName2GID[names[i]])
			di = d.CurveParam.GetNewZn().Div(o.UserName2GID[names[j]], di)
			up.ThenMul(di)
		}
		gY.ThenMul(d.CurveParam.Get0FromG1().PowZn(apks[i], up))
	}
	o.APKMap[attr] = &APK{
		Gy: gY,
	}
	return nil
}

//组装其他用户给的key part(auth)
func (o *Org) AssembleKeyPart(names []string, keyParts []*pbc.Element, d *DABE) (*pbc.Element, error) {
	if len(keyParts) != o.T || len(names) != o.T {
		return nil, fmt.Errorf("pks or names isn't eq t")
	}

	key := d.CurveParam.Get1FromG1()
	for i := 0; i < o.T; i++ {
		up := d.CurveParam.Get1FromZn()
		for j := 0; j < o.T; j++ {
			if i == j {
				continue
			}
			di := d.CurveParam.Get0FromZn().Sub(o.UserName2GID[names[j]], o.UserName2GID[names[i]])
			di = d.CurveParam.GetNewZn().Div(o.UserName2GID[names[j]], di)
			up.ThenMul(di)
		}
		key.ThenMul(d.CurveParam.Get0FromG1().PowZn(keyParts[i], up))
	}
	return key, nil
}
