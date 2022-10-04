package DecentralizedABE

import (
	"crypto/sha256"
	"fmt"
	"github.com/Nik-U/pbc"
)

type User struct {
	APKMap   map[string]*APK
	ASKMap   map[string]*ASK
	EGGAlpha *pbc.Element	`field:"2"`
	Alpha    *pbc.Element	`field:"3"`
	GAlpha   *pbc.Element	`field:"0"`
	Name     string
	OPKMap   map[string]*OPKPart
	OSKMap   map[string]*OSKPart
}

func (u *User) GetPK() *pbc.Element {
	return u.EGGAlpha
}

func (u *User) GetAPKMap() map[string]*APK {
	return u.APKMap
}

func (u *User) GenerateNewAttr(attr string, d *DABE) (*APK, error) {
	if u.APKMap[attr] != nil || u.ASKMap[attr] != nil {
		return nil, fmt.Errorf("already has this attr:%s", attr)
	}
	y := d.CurveParam.GetNewZn()
	sk := ASK{y}
	gy := d.G.NewFieldElement().PowZn(d.G, y)
	pk := APK{gy}
	u.APKMap[attr] = &pk
	u.ASKMap[attr] = &sk
	return &pk, nil
}

//授权用户属性
func (u *User) KeyGenByUser(gid string, attr string, d *DABE) (*pbc.Element, error) {
	if u.ASKMap[attr] == nil {
		return nil, fmt.Errorf("don't have this attr, error when %s", attr)
	}
	hashGid := d.CurveParam.GetG1FromStringHash(gid, sha256.New())
	key := hashGid.
		ThenPowZn(u.ASKMap[attr].Y).
		ThenMul(u.GAlpha)
	return key, nil
}

//授权组织属性
func (u *User) KeyGenByOrg(gid string, attr string, d *DABE, orgName string) (*pbc.Element, error) {
	if u.OSKMap[orgName] == nil || u.OSKMap[orgName].ASKMap[attr] == nil {
		return nil, fmt.Errorf("don't have this attr, error when %s", attr)
	}
	hashGid := d.CurveParam.GetG1FromStringHash(gid, sha256.New())
	key := hashGid.
		ThenPowZn(u.OSKMap[orgName].ASKMap[attr].ASK).
		ThenMul(u.OSKMap[orgName].GOSK)
	return key, nil
}

//创建Org所需的秘密share
func (u *User) GenerateOrgShare(n, t int, userNames map[string]*pbc.Element, orgName string, d *DABE) (
	map[string]*pbc.Element, error) {

	if u.OSKMap[orgName] != nil || u.OPKMap[orgName] != nil {
		return nil, fmt.Errorf("already has this org")
	}
	alphaPart := d.CurveParam.GetNewZn()
	f := make([]*pbc.Element, 0, 0)
	f = append(f, alphaPart)
	for i := 1; i < t; i++ {
		f = append(f, d.CurveParam.GetNewZn())
	}
	oskPart := &OSKPart{
		AlphaPart:   alphaPart,
		ASKMap:      make(map[string]*ASKPart),
		F:           f,
		N:           n,
		T:           t,
		OthersShare: make([]*pbc.Element, 0, 0),
	}
	opkPart := &OPKPart{
		APKMap: make(map[string]*pbc.Element),
	}
	u.OSKMap[orgName] = oskPart
	u.OPKMap[orgName] = opkPart

	shares := make(map[string]*pbc.Element)
	for name, hGID := range userNames {
		shares[name] = u.share(hGID, d, n, t, f)
	}
	return shares, nil
}

//创建Org属性所需的秘密share
func (u *User) GenerateOrgAttrShare(n, t int, org *Org, d *DABE, attrName string) (
	map[string]*pbc.Element, error) {

	if !CheckAttrName(attrName, org.Name) {
		return nil, fmt.Errorf("attrName is invalid")
	}
	if u.OSKMap[org.Name] == nil || u.OPKMap[org.Name] == nil {
		return nil, fmt.Errorf("doesn't has this org")
	}
	if u.OSKMap[org.Name].ASKMap[attrName] != nil || u.OPKMap[org.Name].APKMap[attrName] != nil {
		return nil, fmt.Errorf("already has this attr")
	}
	yPart := d.CurveParam.GetNewZn()
	f := make([]*pbc.Element, 0, 0)
	f = append(f, yPart)
	for i := 1; i < t; i++ {
		f = append(f, d.CurveParam.GetNewZn())
	}

	askPart := &ASKPart{
		F:     f,
		YPart: yPart,
	}
	u.OSKMap[org.Name].ASKMap[attrName] = askPart

	shares := make(map[string]*pbc.Element)
	for name, hGID := range org.UserName2GID {
		shares[name] = u.share(hGID, d, n, t, f)
	}
	return shares, nil
}

//组装其他用户的share，传入aid为0表示为了生成opk，为1表示为了生成apk
func (u *User) AssembleShare(names []string, name2share map[string]*pbc.Element, d *DABE,
	n int, aid int, orgName string, attrName string) (*pbc.Element, error) {

	if aid < 0 || aid > 1 {
		return nil, fmt.Errorf("wrong aid")
	}
	if len(name2share) != n || len(names) != n {
		return nil, fmt.Errorf("length not enough")
	}
	key := d.CurveParam.Get0FromZn()
	for _, name := range names {
		key.ThenAdd(name2share[name])
		//fmt.Println("---"+key.String())
	}
	if aid == 0 {
		u.OSKMap[orgName].OSK = key
		u.OSKMap[orgName].GOSK = d.CurveParam.Get0FromG1().PowZn(d.G, key)
		u.OPKMap[orgName].OPK = d.CurveParam.Get0FromGT().PowZn(d.EGG, key)
		return u.OPKMap[orgName].OPK, nil
	} else {
		u.OSKMap[orgName].ASKMap[attrName].ASK = key
		u.OPKMap[orgName].APKMap[attrName] = d.CurveParam.Get0FromG1().PowZn(d.G, key)
		return u.OPKMap[orgName].APKMap[attrName], nil
	}
}

//get sij
func (u *User) share(otherHGID *pbc.Element, d *DABE, n, t int, f []*pbc.Element) *pbc.Element {
	sij := d.CurveParam.Get0FromZn()
	//from t-1 -> 1, O(t)
	for index := t - 1; index >= 1; index-- {
		sij.ThenAdd(f[index])
		sij.ThenMul(otherHGID)
	}
	sij.ThenAdd(f[0])
	return sij
}
