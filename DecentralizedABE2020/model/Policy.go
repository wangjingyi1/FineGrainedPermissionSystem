package DecentralizedABE

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

/* has everything */
type Policy struct {
	PolicyDescription string //one user one policy or multiple?
	PolicyTreeRoot    *PolicyNode
	AccessStruct      *AccessStruct
}

func (this *Policy) Grow() *Policy {
	if this.PolicyDescription == "" {
		fmt.Printf("error:: user's policy description is EMPTY.\n")
		return nil
	} else {
		policyStr := this.PolicyDescription
		this.PolicyTreeRoot, this.AccessStruct = ParsePolicyStringToTree(&policyStr)
		//fmt.Println(this.AccessStruct)
		this.AccessStruct.genLsssMatrix()
		this.AccessStruct.padLsssMatrix()
		this.AccessStruct.genPolicyTreePath()
	}
	return this
}

/*  Policy Node */
type PolicyNode struct {
	Type      byte //0: node. 1: leaf
	Operation byte //1: and. 2: or
	Attr      string
	Max       int
	Min       int
	Children  []*PolicyNode
}

func NewPolicyNode(attr string, t byte) *PolicyNode {
	N := new(PolicyNode)
	N.Attr = attr
	N.Type = t
	N.Operation = 0
	N.Max = 0
	N.Min = 0
	N.Children = nil
	//fmt.Printf("NewPolicyNode:: %v\n",Attr)
	return N
}
func (this *PolicyNode) GetAttr() string {
	return this.Attr
}
func (this *PolicyNode) SetAttr(attr string) *PolicyNode {
	this.Attr = attr
	return this
}
func (this *PolicyNode) GetMax() int {
	return this.Max
}
func (this *PolicyNode) SetMax(max int) *PolicyNode {
	this.Max = max
	return this
}
func (this *PolicyNode) GetOperation() byte {
	return this.Operation
}
func (this *PolicyNode) SetOperation(o byte) *PolicyNode {
	this.Operation = o
	return this
}
func (this *PolicyNode) GetMin() int {
	return this.Min
}
func (this *PolicyNode) SetMin(min int) *PolicyNode {
	this.Min = min
	return this
}
func (this *PolicyNode) GetChildren() []*PolicyNode {
	return this.Children
}
func (this *PolicyNode) SetChildren(children []*PolicyNode) *PolicyNode {
	this.Children = children
	return this
}

type sendpon struct {
	Type      byte //0: node. 1: leaf
	Operation byte //1: and. 2: or
	Attr      string
	Max       int
	Min       int
	Children  [][]byte
}

func (b *PolicyNode) Serialize() []byte {
	var result bytes.Buffer
	var spon *sendpon
	spon = new(sendpon)
	spon.Type = b.Type
	spon.Operation = b.Operation
	spon.Attr = b.Attr
	spon.Max = b.Max
	spon.Min = b.Min
	for _, c := range b.Children {
		spon.Children = append(spon.Children, c.Serialize())
	}
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(spon)
	if err != nil {
		log.Panic(err)
	}

	return result.Bytes()
}

func DeserializePON(d []byte) *PolicyNode {
	pon := new(PolicyNode)
	var spon sendpon

	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&spon)
	if err != nil {
		log.Panic(err)
	}

	pon.Type = spon.Type
	pon.Operation = spon.Operation
	pon.Attr = spon.Attr
	pon.Max = spon.Max
	pon.Min = spon.Min
	for _, c := range spon.Children {
		pon.Children = append(pon.Children, DeserializePON(c))
	}

	return pon
}

type sendpy struct {
	PolicyDescription string //one user one policy or multiple?
	PolicyTreeRoot    []byte
	AccessStruct      []byte
}

func (b *Policy) Serialize() []byte {
	var result bytes.Buffer
	var spy *sendpy
	spy = new(sendpy)
	spy.PolicyDescription = b.PolicyDescription
	spy.PolicyTreeRoot = b.PolicyTreeRoot.Serialize()
	spy.AccessStruct = b.AccessStruct.Serialize()

	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(spy)
	if err != nil {
		log.Panic(err)
	}

	return result.Bytes()
}

func DeserializePY(d []byte, cp *CurveParam) *Policy {
	py := new(Policy)
	var spy sendpy

	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&spy)
	if err != nil {
		log.Panic(err)
	}

	py.PolicyDescription = spy.PolicyDescription
	py.PolicyTreeRoot = DeserializePON(spy.PolicyTreeRoot)
	py.AccessStruct = DeserializeACS(spy.AccessStruct, cp)

	return py
}
