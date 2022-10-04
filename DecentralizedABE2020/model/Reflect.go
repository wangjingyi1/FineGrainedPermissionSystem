package DecentralizedABE

import (
	"encoding/json"
	"fmt"
	"github.com/Nik-U/pbc"
	"math/big"
	"reflect"
)

type Helper struct {
	P       *big.Int
	Param   *pbc.Params
	Pairing *pbc.Pairing
	G       *pbc.Element
	EGG     *pbc.Element
}

var helper = new(Helper)

func (d *Helper) Setup() {
	d.P = new(big.Int)
	d.P.SetString("242661090146032969904098483991985908921", 10)
	d.Param = pbc.GenerateA1(d.P)
	d.Pairing = d.Param.NewPairing()
	d.G = d.Pairing.NewUncheckedElement(0).Rand()
	d.EGG = d.Pairing.NewUncheckedElement(2).Rand()
}

func init() {
	helper.Setup()
}

func (d *Helper) Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj).Elem()
	v := reflect.ValueOf(obj).Elem()

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).Type.String() == "*pbc.Params" {
			data[t.Field(i).Name] = ((v.Field(i).Interface()).(*pbc.Params)).String()
		} else if t.Field(i).Type.String() == "*pbc.Pairing" {
			data[t.Field(i).Name] = ""
		} else if t.Field(i).Type.String() == "*pbc.Element" {
			data[t.Field(i).Name] = ((v.Field(i).Interface()).(*pbc.Element)).String()
		} else if t.Field(i).Type.String() == "*big.Int" {
			data[t.Field(i).Name] = ((v.Field(i).Interface()).(*big.Int)).String()
		} else if t.Field(i).Type.String() == "[]*pbc.Element" {
			elem := (v.Field(i).Interface()).([]*pbc.Element)
			elemStr := make([]string, len(elem))
			for i, v := range elem {
				elemStr[i] = v.String()
			}
			data[t.Field(i).Name] = elemStr
		} else if t.Field(i).Type.String() == "[]uint8" {
			data[t.Field(i).Name] = string((v.Field(i).Interface()).([]uint8))
		} else if t.Field(i).Type.String() == "int" {
			data[t.Field(i).Name] = (v.Field(i).Interface()).(int)
		} else if t.Field(i).Type.String() == "map[string]*pbc.Element" {
			tem := ((v.Field(i).Interface()).(map[string]*pbc.Element))
			elemStr := make(map[string]string)
			for k, v := range tem {
				elemStr[k] = v.String()
			}
			data[t.Field(i).Name] = elemStr
		} else if t.Field(i).Type.String() == "map[string]*DecentralizedABE.APK" {
			tem := ((v.Field(i).Interface()).(map[string]*APK))
			elemStr := make(map[string]string)
			for k, v := range tem {
				elemStr[k] = v.Gy.String()
			}
			data[t.Field(i).Name] = elemStr
		} else if t.Field(i).Type.String() == "map[string]*DecentralizedABE.ASK" {
			tem := ((v.Field(i).Interface()).(map[string]*ASK))
			elemStr := make(map[string]string)
			for k, v := range tem {
				elemStr[k] = v.Y.String()
			}
			data[t.Field(i).Name] = elemStr
		} else if t.Field(i).Type.String() == "map[string]*DecentralizedABE.OPKPart" {
			tem := ((v.Field(i).Interface()).(map[string]*OPKPart))
			elemStr := make(map[string]string)
			for k, v := range tem {
				opk := v.OPK.String()
				apkmap := make(map[string]string)
				for k1, v1 := range v.APKMap {
					apkmap[k1] = v1.String()
				}
				opkpart := &struct {
					OPK    string;
					APKMap map[string]string
				}{opk, apkmap}
				raw, _ := json.Marshal(opkpart)
				elemStr[k] = string(raw)
			}
			data[t.Field(i).Name] = elemStr
		} else if t.Field(i).Type.String() == "map[string]*DecentralizedABE.OSKPart" {
			tem := ((v.Field(i).Interface()).(map[string]*OSKPart))
			elemStr := make(map[string]string)
			for k, v := range tem {
				alphapart := v.AlphaPart.String()
				osk := v.OSK.String()
				gosk := v.GOSK.String()
				n := int(v.N)
				t := int(v.T)
				f := v.F
				fStr := make([]string, len(f))
				for i1, v1 := range f {
					fStr[i1] = v1.String()
				}
				others := v.OthersShare
				otherStr := make([]string, len(others))
				for i2, v2 := range others {
					otherStr[i2] = v2.String()
				}
				askmap := make(map[string]string)
				for k3, v3 := range v.ASKMap {
					af := v3.F
					afStr := make([]string, len(af))
					for i4, v4 := range af {
						afStr[i4] = v4.String()
					}
					aothers := v3.OthersShare
					aotherStr := make([]string, len(aothers))
					for i5, v5 := range aothers {
						aotherStr[i5] = v5.String()
					}
					ypart := v3.YPart.String()
					ask := v3.ASK.String()
					askpart := &struct {
						F           []string;
						OthersShare []string;
						YPart       string;
						ASK         string
					}{afStr, aotherStr, ypart, ask}
					raw1, _ := json.Marshal(askpart)
					askmap[k3] = string(raw1)

				}
				oskpart := &struct {
					AlphaPart   string;
					ASKMap      map[string]string;
					F           []string;
					N           int;
					T           int;
					OthersShare []string;
					OSK         string;
					GOSK        string
				}{alphapart, askmap, fStr, n, t, otherStr, osk, gosk}
				raw2, _ := json.Marshal(oskpart)
				elemStr[k] = string(raw2)
			}
			data[t.Field(i).Name] = elemStr
			//fmt.Println(elemStr)
		} else {
			data[t.Field(i).Name] = v.Field(i).Interface()
		}
		//fmt.Println(t.Field(i).Type.String())
		//map[string]*DecentralizedABE.APK
	}
	return data
}

func (d *Helper) Str2Struct(str []byte, obj interface{}) {
	var data = make(map[string]interface{})
	err := json.Unmarshal(str, &data)
	if err != nil {
		fmt.Println(err)
	}

	//fmt.Println(reflect.TypeOf(data["OSKMap"]))

	t := reflect.TypeOf(obj).Elem()
	v := reflect.ValueOf(obj).Elem()

	var tem interface{}
	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).Type.String() == "*big.Int" {
			tem = new(big.Int)
			tem.(*big.Int).SetString(data[t.Field(i).Name].(string), 10)
		} else if t.Field(i).Type.String() == "*pbc.Params" {
			tem = pbc.GenerateA1(v.FieldByName("P").Interface().(*big.Int))
		} else if t.Field(i).Type.String() == "*pbc.Pairing" {
			tem = v.FieldByName("Param").Interface().(*pbc.Params).NewPairing()
		} else if t.Field(i).Type.String() == "*pbc.Element" {
			if t.Field(i).Name == "G" {
				tem, _ = helper.G.NewFieldElement().SetString(data["G"].(string), 10)
			} else if t.Field(i).Name == "EGG" {
				tem, _ = helper.EGG.NewFieldElement().SetString(data["EGG"].(string), 10)
			}
		} else if t.Field(i).Type.String() == "[]uint8" {
			tem = []byte(data[t.Field(i).Name].(string))
		} else if t.Field(i).Type.String() == "string" {
			tem = data[t.Field(i).Name].(string)
		} else if t.Field(i).Type.String() == "int" {
			tem = int(data[t.Field(i).Name].(float64))
		} else if t.Field(i).Type.String() == "[]*pbc.Element" {
			raw := data[t.Field(i).Name].([]interface{})
			c1s := make([]*pbc.Element, len(raw), len(raw))
			for i, v := range raw {
				c1s[i], _ = helper.G.NewFieldElement().SetString(v.(string), 10)
			}
			tem = c1s
		} else if t.Field(i).Type.String() == "map[string]*pbc.Element" {
			raw := data[t.Field(i).Name].(map[string]interface{})
			elem := make(map[string]*pbc.Element)
			for k, v := range raw {
				elem[k], _ = helper.G.NewFieldElement().SetString(v.(string), 10)
			}
			tem = elem
		} else if t.Field(i).Type.String() == "map[string]*DecentralizedABE.APK" {
			raw := data[t.Field(i).Name].(map[string]interface{})
			elem := make(map[string]*APK)
			for k, v := range raw {
				gy, _ := helper.G.NewFieldElement().SetString(v.(string), 10)
				elem[k] = &APK{gy}
			}
			tem = elem
		} else if t.Field(i).Type.String() == "map[string]*DecentralizedABE.ASK" {
			raw := data[t.Field(i).Name].(map[string]interface{})
			elem := make(map[string]*ASK)
			for k, v := range raw {
				y, _ := helper.G.NewFieldElement().SetString(v.(string), 10)
				elem[k] = &ASK{y}
			}
			tem = elem
		} else if t.Field(i).Type.String() == "map[string]*DecentralizedABE.OPKPart" {
			raw := data[t.Field(i).Name].(map[string]interface{})
			elem := make(map[string]*OPKPart)
			for k, v := range raw {
				value := &struct {
					OPK    string;
					APKMap map[string]string
				}{}
				_ = json.Unmarshal([]byte(v.(string)), value)
				//fmt.Println(reflect.TypeOf(v))
				opk, _ := helper.G.NewFieldElement().SetString(value.OPK, 10)
				apkmap := make(map[string]*pbc.Element)
				//value1:=value["APKMap"].(map[string]interface{})
				for k1, v1 := range value.APKMap {
					apk, _ := helper.G.NewFieldElement().SetString(v1, 10)
					apkmap[k1] = apk
				}
				elem[k] = &OPKPart{opk, apkmap}
			}
			tem = elem
		} else if t.Field(i).Type.String() == "map[string]*DecentralizedABE.OSKPart" {
			raw := data[t.Field(i).Name].(map[string]interface{})
			elem := make(map[string]*OSKPart)
			for k, v := range raw {
				value := &struct {
					AlphaPart   string;
					ASKMap      map[string]string;
					F           []string;
					N           int;
					T           int;
					OthersShare []string;
					OSK         string;
					GOSK        string
				}{}
				_ = json.Unmarshal([]byte(v.(string)), value)
				//fmt.Println(reflect.TypeOf(v))
				alphapart, _ := helper.G.NewFieldElement().SetString(value.AlphaPart, 10)
				osk, _ := helper.G.NewFieldElement().SetString(value.OSK, 10)
				gosk, _ := helper.G.NewFieldElement().SetString(value.GOSK, 10)
				f := make([]*pbc.Element, len(value.F), len(value.F))
				for i1, v1 := range value.F {
					f[i1], _ = helper.G.NewFieldElement().SetString(v1, 10)
				}
				other := make([]*pbc.Element, len(value.OthersShare), len(value.OthersShare))
				for i2, v2 := range value.OthersShare {
					other[i2], _ = helper.G.NewFieldElement().SetString(v2, 10)
				}
				n := int(value.N)
				t := int(value.T)
				askmap := make(map[string]*ASKPart)
				//value1:=value["APKMap"].(map[string]interface{})
				for k3, v3 := range value.ASKMap {
					askpart := &struct {
						F           []string;
						OthersShare []string;
						YPart       string;
						ASK         string
					}{}
					_ = json.Unmarshal([]byte(v3), askpart)
					ypart, _ := helper.G.NewFieldElement().SetString(askpart.YPart, 10)
					ask, _ := helper.G.NewFieldElement().SetString(askpart.ASK, 10)
					af := make([]*pbc.Element, len(askpart.F), len(askpart.F))
					for i4, v4 := range askpart.F {
						af[i4], _ = helper.G.NewFieldElement().SetString(v4, 10)
					}
					aother := make([]*pbc.Element, len(askpart.OthersShare), len(askpart.OthersShare))
					for i5, v5 := range askpart.OthersShare {
						aother[i5], _ = helper.G.NewFieldElement().SetString(v5, 10)
					}
					askmap[k3] = &ASKPart{af, aother, ypart, ask}
				}
				elem[k] = &OSKPart{alphapart, askmap, f, n, t, other, osk, gosk}
			}
			tem = elem
		}
		v.Field(i).Set(reflect.ValueOf(tem))
	}
}
