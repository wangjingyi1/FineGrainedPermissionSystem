package DecentralizedABE

import (
	"fmt"
	"github.com/Nik-U/pbc"
	"testing"
)

/*func TestReflect(t *testing.T) {
	helper = new(Helper)
	dabe := new(DABE)
	dabe.GlobalSetup()
	fudanUniversity := dabe.UserSetup("Fudan_University")
	fmt.Printf("%v\n", fudanUniversity)
	s := helper.Struct2Map(fudanUniversity)
	bytes, _ := json.Marshal(s)
	fmt.Println(bytes)
	user := new(User)
	helper.Str2Struct(bytes, user)
	fmt.Printf("%v\n", user)
	assert.Equal(t, fudanUniversity, user, "reflect error")
}*/

func TestReflect2(t *testing.T) {
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
	_, _ = user2.GenerateNewAttr("123456", dabe)

	bytes, err := Serialize2Bytes(user2)
	if err != nil {
		panic(err)
	}
	//fmt.Println(bytes)
	user := new(User)
	err = Deserialize2Struct(bytes, user)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Printf("%v\n", user2)
	fmt.Printf("%v\n", user)
	for k, v := range user2.OSKMap {
		fmt.Printf("%v\n%v\n", v, user.OSKMap[k])
	}
}
