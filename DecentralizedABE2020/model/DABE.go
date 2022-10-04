package DecentralizedABE

import (
	"crypto/sha256"
	"fmt"
	"github.com/Nik-U/pbc"
	"github.com/thorweiyan/DecentralizedABE2020/model/AES"
)

type DABE struct {
	CurveParam *CurveParam
	G          *pbc.Element
	EGG        *pbc.Element
}

func (d *DABE) GlobalSetup() {
	fmt.Println("DABE GlobalSetup start")
	d.CurveParam = new(CurveParam)
	d.CurveParam.Initialize()
	d.G = d.CurveParam.GetNewG1()
	d.EGG = d.CurveParam.GetNewGT().Pair(d.G, d.G)
	fmt.Println("DABE GlobalSetup success")
}

func (d *DABE) UserSetup(name string) *User {
	fmt.Println("DABE UserSetup start")
	alpha := d.CurveParam.GetNewZn()
	eGGAlpha := d.EGG.NewFieldElement().PowZn(d.EGG, alpha)
	gAlpha := d.G.NewFieldElement().PowZn(d.G, alpha)
	fmt.Printf("DABE UserSetup success for %s\n", name)
	return &User{
		APKMap:   make(map[string]*APK),
		ASKMap:   make(map[string]*ASK),
		EGGAlpha: eGGAlpha,
		Alpha:    alpha,
		GAlpha:   gAlpha,
		Name:     name,
		OPKMap:   make(map[string]*OPKPart),
		OSKMap:   make(map[string]*OSKPart),
	}
}

func (d *DABE) OrgSetup(n, t int, name string, userNames []string) (*Org, error) {
	fmt.Println("DABE OrgSetup start")
	if t > n {
		return nil, fmt.Errorf("threshold can not bigger than n")
	}
	if len(userNames) != n {
		return nil, fmt.Errorf("userNames' length doesn't eq n")
	}
	hash := sha256.New()
	user2gid := make(map[string]*pbc.Element)
	for _, userName := range userNames {
		hashId := d.CurveParam.GetZnFromStringHash(userName, hash)
		user2gid[userName] = hashId
	}
	fmt.Printf("DABE OrgSetup success for %s\n", name)
	return &Org{
		APKMap:       make(map[string]*APK),
		EGGAlpha:     nil,
		Name:         name,
		N:            n,
		T:            t,
		UserName2GID: user2gid,
	}, nil
}

func (d *DABE) Encrypt(m string, uPolicy string, authorities map[string]Authority) (*Cipher, error) {
	fmt.Println("DABE Encrypt start")
	aesKey := d.EGG.NewFieldElement().Rand()
	aesCipherText, err := AES.AesEncrypt([]byte(m), (aesKey.Bytes())[0:32])
	if err != nil {
		return nil, fmt.Errorf("AES encrypt error\n")
	}

	policy := new(Policy)
	d.growNewPolicy(uPolicy, d.CurveParam.GetNewZn(), policy)

	n := len(policy.AccessStruct.LsssMatrix) - 1
	l := len(policy.AccessStruct.LsssMatrix[0])
	v := make([]*pbc.Element, l, l)
	w := make([]*pbc.Element, l, l)
	c1s := make([]*pbc.Element, n, n)
	c2s := make([]*pbc.Element, n, n)
	c3s := make([]*pbc.Element, n, n)
	s := d.CurveParam.GetNewZn()

	// c0 = M * e(g,g)^s
	c0 := aesKey.Mul(aesKey, d.EGG.NewFieldElement().PowZn(d.EGG, s))
	//generate v and w
	v[0] = s
	w[0] = s.NewFieldElement().Set0()
	for i := 1; i < l; i++ {
		v[i] = d.CurveParam.GetNewZn()
		w[i] = d.CurveParam.GetNewZn()
	}
	//generate c1s,c2s,c3s
	for i := 0; i < n; i++ {
		//attr
		attrStr := policy.AccessStruct.PolicyMaps[i+1]
		authorityName := GetAuthorityNameFromAttrName(attrStr)
		if authorities[authorityName] == nil {
			return nil, fmt.Errorf("authority not found, error when %s", attrStr)
		}
		authority := authorities[authorityName]
		pk := authority.GetAPKMap()[attrStr]
		if pk == nil {
			return nil, fmt.Errorf("pk not found, error when %s", attrStr)
		}
		//r
		r := d.CurveParam.GetNewZn()
		//c2 = g^r
		c2 := d.G.NewFieldElement().PowZn(d.G, r)

		//Ai*v
		AiV := policy.AccessStruct.LsssMatrixDotMulVector(i+1, v)
		//e(g,g)^(Ai*v)
		c1 := d.EGG.NewFieldElement().PowZn(d.EGG, AiV)
		//c1 = e(g,g)^(Ai*v) * e(g,g)^ ( alpha_p(x) * r_x )
		rightTemp := authority.GetPK().NewFieldElement().PowZn(authority.GetPK(), r)
		c1.Mul(c1, rightTemp)

		//Ai*w
		AiW := policy.AccessStruct.LsssMatrixDotMulVector(i+1, w)
		//g^(Ai*w)
		c3 := d.G.NewFieldElement().PowZn(d.G, AiW)
		//c3 = g^(y_p(x) * r) * g^(Ai*w)
		c3.Mul(c3, pk.Gy.NewFieldElement().PowZn(pk.Gy, r))

		c1s[i] = c1
		c2s[i] = c2
		c3s[i] = c3
	}
	fmt.Println("DABE Encrypt success")
	return &Cipher{
		C0:         c0,
		C1s:        c1s,
		C2s:        c2s,
		C3s:        c3s,
		CipherText: aesCipherText,
		Policy:     uPolicy,
	}, nil
}

func (d *DABE) Decrypt(cipher *Cipher, privateKeys map[string]*pbc.Element, gid string) ([]byte, error) {
	fmt.Println("DABE Decrypt start")
	hashGid := d.CurveParam.GetG1FromStringHash(gid, sha256.New())

	policy := new(Policy)
	d.growNewPolicy(cipher.Policy, d.CurveParam.GetNewZn(), policy)
	n := len(policy.AccessStruct.LsssMatrix) - 1
	attrs := make([]string, 0, 0)
	for key, _ := range privateKeys {
		attrs = append(attrs, key)
	}
	// sum(cx * Ax) = (1,0,0,0...)
	cxs, err := d.genCoefficient(attrs, policy)
	if err != nil {
		return nil, err
	}

	// (c1 * e(HGID,c3) / e(key, c2)) ^ cx  累×后得到 e(g,g)^s
	result := d.EGG.NewFieldElement().Set1()
	for i := 0; i < n; i++ {
		if cxs[i+1] == nil {
			continue
		}
		//attr
		attrStr := policy.AccessStruct.PolicyMaps[i+1]
		// c1 * e(HGID,c3)
		temp := d.EGG.NewFieldElement().Pair(hashGid, cipher.C3s[i]).ThenMul(cipher.C1s[i])
		// e(key, c2)
		temp2 := d.EGG.NewFieldElement().Pair(privateKeys[attrStr], cipher.C2s[i])
		// (c1 * e(HGID,c3) / e(key, c2)) ^ cx
		temp.ThenDiv(temp2)
		temp.ThenPowZn(cxs[i+1])
		// 累×
		result.ThenMul(temp)
	}
	aesKey := d.EGG.NewFieldElement().Set(cipher.C0).ThenDiv(result)
	if aesKey == nil {
		return nil, fmt.Errorf("User policy not match,decrypt failed.\n")
	}
	if len(aesKey.Bytes()) <= 32 {
		return nil, fmt.Errorf("invalid aeskey:: decrypt failed.\n")
	}
	M, err := AES.AesDecrypt(cipher.CipherText, (aesKey.Bytes())[0:32])
	if err != nil || M == nil {
		return nil, fmt.Errorf("aes error:: decrypt failed.\n")
	}
	fmt.Println("DABE Decrypt success")
	return M, nil
}

func (d *DABE) genCoefficient(attrs []string, policy *Policy) ([]*pbc.Element, error) {
	n := len(attrs)
	attrMap := make(map[string]int)
	for i := 0; i < n; i++ {
		attrMap[attrs[i]] = 1
	}

	nodeLine := make([]int, len(policy.AccessStruct.A))
	leafLine := make([]int, len(policy.AccessStruct.PolicyMaps))
	nodeLine[0] = 1
	isSatisfy := policy.AccessStruct.isSatisfy(&nodeLine, &leafLine, &attrMap, 0)
	if !isSatisfy {
		return nil, fmt.Errorf("User attrs not satisfy the policy: %s", policy.PolicyDescription)
	}

	for i := 0; i < len(policy.AccessStruct.A); i++ {
		for j := 2; j < len(policy.AccessStruct.A[i]); j++ {
			if policy.AccessStruct.A[i][j] > 0 {
				nodeLine[policy.AccessStruct.A[i][j]] *= nodeLine[i]
			} else {
				leafLine[-policy.AccessStruct.A[i][j]] *= nodeLine[i]
			}
		}
	}

	w := make([]*pbc.Element, len(leafLine), len(leafLine))
	for i := 1; i < len(leafLine); i++ {
		if leafLine[i] != 0 {
			w[i] = d.CurveParam.GetNewZn().Set1().ThenDiv(d.CurveParam.GetNewZn().SetInt32(int32(leafLine[i])))
		} else {
			//如果不需要就返回nil
			w[i] = nil
		}

	}
	return w, nil
}

func (d *DABE) growNewPolicy(s string, p *pbc.Element, policy *Policy) {
	policy.PolicyDescription = s
	policy.Grow()
	policy.AccessStruct.genElementLsssMatrix(p)
}
