package DecentralizedABE

import (
	"fmt"
	"strconv"
	"strings"
)

/* <Policy Parser SECTION */
func ParsePolicyStringToTree(s *string) (*PolicyNode, *AccessStruct) {
	ss := *s
	AS := NewAccessStruct()
	AS.ParsePolicyStringtoMap(&ss)

	*s = strings.Replace(*s, "AND", "&&", -1)
	*s = strings.Replace(*s, "OR", "||", -1)
	*s = strings.Replace(*s, " ", "", -1)
	MainPolicy, ID := ParsePolicyString(AS, s, 0, len(*s)-1)
	if ID == 0 {
	} //non sense
	return MainPolicy, AS
}

func ParsePolicyString(A *AccessStruct, s *string, startPos int, stopPos int) (*PolicyNode, int) {
	//leftPos := startPos+1+strings.Index((*s)[startPos+1:stopPos], "(")

	this := NewPolicyNode("ThreshHold", 0)

	A.A = append(A.A, make([]int, 2, 2))
	//_A := &(A.A[A.CurrentPointer])
	ID := A.CurrentPointer
	A.A[ID][0] = 0
	A.A[ID][1] = 0
	A.CurrentPointer++
	policy_children := make([]*PolicyNode, 0)

	var i int = startPos + 1
	var n int = 0
	var _n int = 0
	var leftPos int
	var rightPos int
	var trueChild string = ""

	for i <= stopPos {

		leftPos = strings.Index((*s)[i:stopPos], "(")

		if leftPos != -1 {
			trueChild += (*s)[i : i+leftPos]
			rightPos = LookForMyRightBraket(s, i+leftPos)
			tmpPolicy, tmpID := ParsePolicyString(A, s, i+leftPos, rightPos)
			policy_children = append(policy_children, tmpPolicy)
			A.A[ID] = append(A.A[ID], tmpID)
			n++
			i = rightPos + 1
		} else {
			trueChild += (*s)[i:stopPos]
			break
		}
	}

	var childAttr []string
	if strings.Index(trueChild, "&&") != -1 {
		childAttr = strings.Split(trueChild, "&&")

		for v := range childAttr {
			if childAttr[v] != "" {
				policy_children = append(policy_children, NewPolicyNode(childAttr[v], 1).SetMax(1).SetMin(1))
				A.A[ID] = append(A.A[ID], -A.PolicyMap[childAttr[v]])
				A.LeafID--
				n++
			}
		}
		_n = n
		this.SetOperation(1)
	} else if strings.Index(trueChild, "||") != -1 {
		childAttr = strings.Split(trueChild, "||")

		for v := range childAttr {
			if childAttr[v] != "" {
				policy_children = append(policy_children, NewPolicyNode(childAttr[v], 1).SetMax(1).SetMin(1))
				A.A[ID] = append(A.A[ID], -A.PolicyMap[childAttr[v]])
				A.LeafID--
				n++
			}
		}
		_n = 1
		this.SetOperation(2)
	}

	if n == 0 {
		fmt.Printf("Error:: bad description. \n")
	} else {
		this.SetChildren(policy_children)
		this.SetMax(n)
		this.SetMin(_n)
		A.A[ID][0] = n
		A.A[ID][1] = _n
	}

	return this, ID
}

func LookForMyRightBraket(s *string, posL int) int {
	rightPos := posL + strings.Index((*s)[posL:], ")")

	for true {
		if rightPos < posL {
			return -1
		} else {
			leftPos := posL + 1 + strings.Index((*s)[posL+1:rightPos], "(")
			if leftPos > posL+1 {
				posL = LookForMyRightBraket(s, leftPos)
				rightPos = posL + 1 + strings.Index((*s)[posL+1:], ")")
			} else {
				return rightPos
			}
		}
	}
	return 0
}

/* Policy Parser SECTION> */

/* <Utility SECTION */
//读取文件需要经常进行错误检查，这个帮助方法可以精简下面的错误检查过程。
func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

/* Utility SECTION> */

/* utils */
func CharToString(s string, t int) string {
	var sp string = ""
	for i := 0; i < t; i++ {
		sp += s
	}
	return sp
}

func GetPadding(m int, l int, depth int) string {
	var sp string = ""
	sp += CharToString("*", depth-l)
	if m == 0 {
		sp = CharToString("0", l-1) + sp
	} else {
		sp = CharToString("0", l-2) + strconv.FormatUint(uint64(m), 2) + sp
	}
	return sp[len(sp)-(depth-1):]
}

// 检查属性是否以组织/用户名称为前缀
func CheckAttrName(attrName, authorityName string) bool {
	splitN := strings.SplitN(attrName, ":", 2)
	if len(splitN) != 2 {
		return false
	}
	return authorityName == splitN[0]
}

// 根据属性名称获取组织/用户名，出错返回空字符串
func GetAuthorityNameFromAttrName(attrName string) string {
	splitN := strings.SplitN(attrName, ":", 2)
	if len(splitN) != 2 {
		return ""
	}
	return splitN[0]
}
