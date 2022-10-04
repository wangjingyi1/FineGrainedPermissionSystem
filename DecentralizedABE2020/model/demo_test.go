package DecentralizedABE

import (
	"fmt"
	"github.com/Nik-U/pbc"
	"testing"
)

func TestDemo(t *testing.T) {
	//初始化和全局参数生成
	dabe := new(DABE)
	dabe.GlobalSetup()

	//初始化两个不同的权限管理机构，并保存
	authorityMap := make(map[string]Authority)
	fudanUniversity := dabe.UserSetup("Fudan_University")
	authorityMap["Fudan_University"] = fudanUniversity
	ageAuthority := dabe.UserSetup("Age_Authority")
	authorityMap["Age_Authority"] = ageAuthority

	//保存所有属性公钥
	pkMap := make(map[string]*APK)
	//生成属性私钥
	tempPk, err := fudanUniversity.GenerateNewAttr("Fudan_University:在读研究生", dabe)
	if err != nil {
		panic(err)
	}
	pkMap["Fudan_University:在读研究生"] = tempPk
	tempPk2, err := ageAuthority.GenerateNewAttr("Age_Authority:23", dabe)
	if err != nil {
		panic(err)
	}
	pkMap["Age_Authority:23"] = tempPk2
	tempPk3, err := ageAuthority.GenerateNewAttr("Age_Authority:24", dabe)
	if err != nil {
		panic(err)
	}
	pkMap["Age_Authority:24"] = tempPk3

	//用户申请密钥
	user1Privatekeys := make(map[string]*pbc.Element)
	user2Privatekeys := make(map[string]*pbc.Element)

	user1Privatekey1, err := fudanUniversity.KeyGenByUser("陈泽宁", "Fudan_University:在读研究生", dabe)
	if err != nil {
		panic(err)
	}
	user1Privatekey2, err := ageAuthority.KeyGenByUser("陈泽宁", "Age_Authority:23", dabe)
	if err != nil {
		panic(err)
	}
	user2Privatekey1, err := ageAuthority.KeyGenByUser("24岁的无名氏", "Age_Authority:24", dabe)
	if err != nil {
		panic(err)
	}
	user1Privatekeys["Fudan_University:在读研究生"] = user1Privatekey1
	user1Privatekeys["Age_Authority:23"] = user1Privatekey2
	user2Privatekeys["Age_Authority:24"] = user2Privatekey1

	//加密两个不同的明文,这里authorityMap应该不传入私钥相关，方便起见如此做
	m1 := "复旦的在读研究生或者24岁的人可以看见"
	m2 := "复旦的23岁在读研究生可以看见"
	cipher1, err := dabe.Encrypt(m1, "(Fudan_University:在读研究生 OR Age_Authority:24)", authorityMap)
	if err != nil {
		panic(err)
	}
	cipher2, err := dabe.Encrypt(m2, "(Fudan_University:在读研究生 AND Age_Authority:23)", authorityMap)
	if err != nil {
		panic(err)
	}

	//解密
	decrypt, err := dabe.Decrypt(cipher1, user1Privatekeys, "陈泽宁")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("陈泽宁解密出了： " + string(decrypt))
	}
	decrypt2, err := dabe.Decrypt(cipher2, user1Privatekeys, "陈泽宁")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("陈泽宁解密出了： " + string(decrypt2))
	}
	decrypt3, err := dabe.Decrypt(cipher1, user2Privatekeys, "24岁的无名氏")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("24岁的无名氏解密出了： " + string(decrypt3))
	}
	decrypt4, err := dabe.Decrypt(cipher2, user2Privatekeys, "24岁的无名氏")
	if err == nil {
		fmt.Println("24岁的无名氏错误解密出了： " + string(decrypt4))
	} else {
		fmt.Println("24岁的无名氏正常地失败于： " + err.Error())
	}
}

func TestDemo2(t *testing.T) {
	//初始化和全局参数生成
	dabe := new(DABE)
	dabe.GlobalSetup()

	user1Name := "user1"
	user2Name := "user2"
	user3Name := "user3"
	user4Name := "user4"
	org1Name := "org1"

	//初始化4个不同用户
	user1 := dabe.UserSetup(user1Name)
	user2 := dabe.UserSetup(user2Name)
	user3 := dabe.UserSetup(user3Name)
	user4 := dabe.UserSetup(user4Name)

	//初始化两个不同的权限管理机构，并保存
	authorityMap := make(map[string]Authority)
	authorityMap[user1Name] = user1

	userNames := []string{user2Name, user3Name, user4Name}
	org1, err := dabe.OrgSetup(3, 2, org1Name, userNames)
	if err != nil {
		panic(err)
	}

	//生成share
	user2Shares, err := user2.GenerateOrgShare(3, 2, org1.UserName2GID, org1Name, dabe)
	if err != nil {
		panic(err)
	}
	user3Shares, err := user3.GenerateOrgShare(3, 2, org1.UserName2GID, org1Name, dabe)
	if err != nil {
		panic(err)
	}
	user4Shares, err := user4.GenerateOrgShare(3, 2, org1.UserName2GID, org1Name, dabe)
	if err != nil {
		panic(err)
	}

	//交换share
	sharesForUser2 := make(map[string]*pbc.Element)
	sharesForUser2[user2Name] = user2Shares[user2Name]
	sharesForUser2[user3Name] = user3Shares[user2Name]
	sharesForUser2[user4Name] = user4Shares[user2Name]
	user2PK, err := user2.AssembleShare(userNames, sharesForUser2, dabe, 3, 0, org1Name, "")
	if err != nil {
		panic(err)
	}
	sharesForUser3 := make(map[string]*pbc.Element)
	sharesForUser3[user2Name] = user2Shares[user3Name]
	sharesForUser3[user3Name] = user3Shares[user3Name]
	sharesForUser3[user4Name] = user4Shares[user3Name]
	user3PK, err := user3.AssembleShare(userNames, sharesForUser3, dabe, 3, 0, org1Name, "")
	if err != nil {
		panic(err)
	}
	sharesForUser4 := make(map[string]*pbc.Element)
	sharesForUser4[user2Name] = user2Shares[user4Name]
	sharesForUser4[user3Name] = user3Shares[user4Name]
	sharesForUser4[user4Name] = user4Shares[user4Name]
	user4PK, err := user4.AssembleShare(userNames, sharesForUser4, dabe, 3, 0, org1Name, "")
	if err != nil {
		panic(err)
	}
	pks := []*pbc.Element{user2PK, user3PK, user4PK}
	err = org1.GenerateOPK(userNames[:2], pks[:2], dabe)
	if err != nil {
		panic(err)
	}

	authorityMap["org1"] = org1

	//保存所有属性公钥
	pkMap := make(map[string]*APK)
	user1Attr1 := "user1:好朋友"
	user1Attr2 := "user1:仇家"
	org1Attr1 := "org1:正式员工"

	//生成属性私钥
	tempPk, err := user1.GenerateNewAttr(user1Attr1, dabe)
	if err != nil {
		panic(err)
	}
	pkMap[user1Attr1] = tempPk
	tempPk2, err := user1.GenerateNewAttr(user1Attr2, dabe)
	if err != nil {
		panic(err)
	}
	pkMap[user1Attr2] = tempPk2

	//生成share
	user2Shares_, err := user2.GenerateOrgAttrShare(3, 2, org1, dabe, org1Attr1)
	if err != nil {
		panic(err)
	}
	user3Shares_, err := user3.GenerateOrgAttrShare(3, 2, org1, dabe, org1Attr1)
	if err != nil {
		panic(err)
	}
	user4Shares_, err := user4.GenerateOrgAttrShare(3, 2, org1, dabe, org1Attr1)
	if err != nil {
		panic(err)
	}

	//交换share
	sharesForUser2_ := make(map[string]*pbc.Element)
	sharesForUser2_[user2Name] = user2Shares_[user2Name]
	sharesForUser2_[user3Name] = user3Shares_[user2Name]
	sharesForUser2_[user4Name] = user4Shares_[user2Name]
	user2PK_, err := user2.AssembleShare(userNames, sharesForUser2_, dabe, 3, 1, org1Name, org1Attr1)
	if err != nil {
		panic(err)
	}
	sharesForUser3_ := make(map[string]*pbc.Element)
	sharesForUser3_[user2Name] = user2Shares_[user3Name]
	sharesForUser3_[user3Name] = user3Shares_[user3Name]
	sharesForUser3_[user4Name] = user4Shares_[user3Name]
	user3PK_, err := user3.AssembleShare(userNames, sharesForUser3_, dabe, 3, 1, org1Name, org1Attr1)
	if err != nil {
		panic(err)
	}
	sharesForUser4_ := make(map[string]*pbc.Element)
	sharesForUser4_[user2Name] = user2Shares_[user4Name]
	sharesForUser4_[user3Name] = user3Shares_[user4Name]
	sharesForUser4_[user4Name] = user4Shares_[user4Name]
	user4PK_, err := user4.AssembleShare(userNames, sharesForUser4_, dabe, 3, 1, org1Name, org1Attr1)
	if err != nil {
		panic(err)
	}
	apks := []*pbc.Element{user2PK_, user3PK_, user4PK_}
	err = org1.GenerateNewAttr(userNames[1:3], apks[1:3], org1Attr1, dabe)
	if err != nil {
		panic(err)
	}
	pkMap[org1Attr1] = org1.APKMap[org1Attr1]

	//用户申请密钥
	goodManPrivatekeys := make(map[string]*pbc.Element)
	badManPrivatekeys := make(map[string]*pbc.Element)

	goodManPrivatekey1, err := user1.KeyGenByUser("好人", user1Attr1, dabe)
	if err != nil {
		panic(err)
	}
	badManPrivatekey1, err := user1.KeyGenByUser("坏人", user1Attr2, dabe)
	if err != nil {
		panic(err)
	}
	goodManPrivatekeys[user1Attr1] = goodManPrivatekey1
	badManPrivatekeys[user1Attr2] = badManPrivatekey1

	partKey1, err := user2.KeyGenByOrg("好人", org1Attr1, dabe, org1.Name)
	if err != nil {
		panic(err)
	}
	partKey2, err := user3.KeyGenByOrg("好人", org1Attr1, dabe, org1.Name)
	if err != nil {
		panic(err)
	}
	partKey3, err := user3.KeyGenByOrg("坏人", org1Attr1, dabe, org1.Name)
	if err != nil {
		panic(err)
	}
	partKey4, err := user4.KeyGenByOrg("坏人", org1Attr1, dabe, org1.Name)
	if err != nil {
		panic(err)
	}
	goodManPrivatekey2, err := org1.AssembleKeyPart([]string{user2Name, user3Name},
		[]*pbc.Element{partKey1, partKey2}, dabe)
	badManPrivatekey2, err := org1.AssembleKeyPart([]string{user3Name, user4Name},
		[]*pbc.Element{partKey3, partKey4}, dabe)
	goodManPrivatekeys[org1Attr1] = goodManPrivatekey2
	badManPrivatekeys[org1Attr1] = badManPrivatekey2

	//加密两个不同的明文,这里authorityMap应该不传入私钥相关，方便起见如此做
	m1 := "只给正式员工且是user1的好朋友看的密语"
	m2 := "好朋友和仇人都能看见的宣言"
	cipher1, err := dabe.Encrypt(m1, "(user1:好朋友 AND org1:正式员工)", authorityMap)
	if err != nil {
		panic(err)
	}
	cipher2, err := dabe.Encrypt(m2, "(user1:好朋友 OR user1:仇家)", authorityMap)
	if err != nil {
		panic(err)
	}

	//解密
	decrypt, err := dabe.Decrypt(cipher1, goodManPrivatekeys, "好人")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("好人解密出了： " + string(decrypt))
	}
	decrypt2, err := dabe.Decrypt(cipher2, goodManPrivatekeys, "好人")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("好人解密出了： " + string(decrypt2))
	}
	decrypt3, err := dabe.Decrypt(cipher2, badManPrivatekeys, "坏人")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("坏人解密出了： " + string(decrypt3))
	}
	decrypt4, err := dabe.Decrypt(cipher1, badManPrivatekeys, "坏人")
	if err == nil {
		fmt.Println("坏人错误解密出了： " + string(decrypt4))
	} else {
		fmt.Println("坏人正常地失败于： " + err.Error())
	}
}
