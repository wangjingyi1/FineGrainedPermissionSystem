package DecentralizedABE2020

import (
	"fmt"
	DecentralizedABE "DecentralizedABE2020/model"
)

func main() {
	//初始化和全局参数生成
	dabe := new(DecentralizedABE.DABE)
	dabe.GlobalSetup()

	//初始化两个不同的权限管理机构，并保存
	authorityMap := make(map[string]DecentralizedABE.Authority)
	fudanUniversity := dabe.UserSetup("Fudan_University")
	authorityMap["Fudan_University"] = fudanUniversity
	ageAuthority := dabe.UserSetup("Age_Authority")
	authorityMap["Age_Authority"] = ageAuthority

	//保存所有属性公钥
	pkMap := make(map[string]*DecentralizedABE.APK)
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
