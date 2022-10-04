package DecentralizedABE

import (
	"encoding/json"
	"fmt"
	"github.com/Nik-U/pbc"
	"math/big"
	"reflect"
	"strconv"
)

var specialHandle map[string]struct{}
var curve CurveParam

func init() {
	specialHandle = make(map[string]struct{})
	specialHandle["*pbc.Params"] = struct{}{}
	specialHandle["*pbc.Pairing"] = struct{}{}
	specialHandle["*pbc.Element"] = struct{}{}
	specialHandle["*big.Int"] = struct{}{}
	specialHandle["int"] = struct{}{}
	curve.Initialize()
}

func Serialize2Bytes(obj interface{}) ([]byte, error) {
	serialize2Map, err := Serialize2Map(obj)
	if err != nil {
		return nil, err
	}
	bytes, err := json.Marshal(serialize2Map)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func Serialize2Map(obj interface{}) (map[string]interface{}, error) {
	var err error
	if obj == nil {
		return nil, fmt.Errorf("nil input")
	}
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	data := make(map[string]interface{})
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)
		if _, exist := specialHandle[field.Type.String()]; exist {
			data[field.Name] = serializeHandle(field.Type, value)
			continue
		}
		switch field.Type.Kind() {
		case reflect.Slice:
			nestedName := field.Type.Elem().String()
			tempData := make([]interface{}, value.Len())
			if _, exist := specialHandle[nestedName]; !exist {
				for i := 0; i < value.Len(); i++ {
					tempData[i], err = Serialize2Map(value.Index(i).Interface())
					if err != nil {
						return nil, err
					}
				}
			} else {
				for i := 0; i < value.Len(); i++ {
					tempData[i] = serializeHandle(field.Type.Elem(), value.Index(i))
				}
			}
			data[field.Name] = tempData
			continue
		case reflect.Map:
			tempData := make(map[string]interface{}, len(value.MapKeys()))
			for _, key := range value.MapKeys() {
				innerVal := value.MapIndex(key)
				tempData[key.Interface().(string)], err = Serialize2Map(innerVal.Interface())
				if err != nil {
					return nil, err
				}
			}
			data[field.Name] = tempData
			continue
		case reflect.Struct:
			tempData, err := Serialize2Map(value)
			if err != nil {
				return nil, err
			}
			data[field.Name] = tempData
			continue
		default:
			data[field.Name] = value.Interface()
		}
	}
	return data, nil
}

func Deserialize2Struct(bytes []byte, obj interface{}) error {
	data := make(map[string]interface{})
	if err := json.Unmarshal(bytes, &data); err != nil {
		fmt.Println(err.Error())
		return err
	}
	obj, e := deserialize2Struct(data, obj)
	return e
}

func deserialize2Struct(data map[string]interface{}, obj interface{}) (interface{}, error) {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)
		if _, exist := specialHandle[field.Type.String()]; exist {
			result, err := deserializeHandle(field.Type, data[field.Name], field.Tag)
			if err != nil {
				return nil, err
			}
			value.Set(reflect.ValueOf(result))
			continue
		}

		switch field.Type.Kind() {
		case reflect.Slice:
			innerType := field.Type.Elem()
			tempArray := data[field.Name].([]interface{})
			tempData := reflect.MakeSlice(field.Type, len(tempArray), len(tempArray))
			if _, exist := specialHandle[innerType.String()]; exist {
				for i, v := range tempArray {
					result, err := deserializeHandle(innerType, v, field.Tag)
					if err != nil {
						return nil, err
					}
					tempData.Index(i).Set(reflect.ValueOf(result))
				}
			} else {
				if innerType.Kind() == reflect.Ptr {
					innerType = innerType.Elem()
				}
				for i, v := range tempArray {
					result, err := deserialize2Struct(v.(map[string]interface{}), reflect.New(innerType).Interface())
					if err != nil {
						return nil, err
					}
					tempData.Index(i).Set(reflect.ValueOf(result))
				}
			}
			value.Set(tempData)
			continue
		case reflect.Map:
			innerType := field.Type.Elem()
			tempMap := data[field.Name].(map[string]interface{})
			tempData := reflect.MakeMap(field.Type)
			//tempData := make(map[string]interface{}, len(tempMap))
			if _, exist := specialHandle[innerType.String()]; exist {
				for k, v := range tempMap {
					result, err := deserializeHandle(innerType, v, field.Tag)
					if err != nil {
						return nil, err
					}
					tempData.SetMapIndex(reflect.ValueOf(k), reflect.ValueOf(result))
					//tempData[k] = result
				}
			} else {
				for k, v := range tempMap {
					if innerType.Kind() == reflect.Ptr {
						innerType = innerType.Elem()
					}
					result, err := deserialize2Struct(v.(map[string]interface{}), reflect.New(innerType).Interface())
					if err != nil {
						return nil, err
					}
					tempData.SetMapIndex(reflect.ValueOf(k), reflect.ValueOf(result))
					//tempData[k] = result
				}
			}
			value.Set(tempData)
			continue
		case reflect.Struct:
			result, err := deserialize2Struct(data[field.Name].(map[string]interface{}), reflect.New(field.Type))
			if err != nil {
				return nil, err
			}
			value.Set(reflect.ValueOf(result))
			continue
		default:
			value.Set(reflect.ValueOf(data[field.Name]))
		}
	}

	return obj, nil
}

func serializeHandle(fieldType reflect.Type, val reflect.Value) interface{} {
	switch fieldType.String() {
	case "*pbc.Params":
		return ""
	case "*pbc.Pairing":
		return ""
	case "*pbc.Element":
		return (val.Interface().(*pbc.Element)).String()
	case "*big.Int":
		return (val.Interface().(*big.Int)).String()
	default:
		return val.Interface()
	}
}

func deserializeHandle(fieldType reflect.Type, obj interface{}, tag reflect.StructTag) (interface{}, error) {
	switch fieldType.String() {
	case "*pbc.Params":
		return curve.Param, nil
	case "*pbc.Pairing":
		return curve.Pairing, nil
	case "*pbc.Element":
		fieldStr := tag.Get("field")
		i, err := strconv.Atoi(fieldStr)
		if err != nil {
			return nil, err
		}
		element, b := curve.Pairing.NewUncheckedElement(pbc.Field(i)).SetString(obj.(string), 10)
		if !b {
			return nil, fmt.Errorf("deserialze pbc.Element error with" + obj.(string))
		}
		return element, nil
	case "*big.Int":
		result, b := new(big.Int).SetString(obj.(string), 10)
		if !b {
			return nil, fmt.Errorf("deserialze big.Int error with" + obj.(string))
		}
		return result, nil
	case "int":
		return int(obj.(float64)), nil
	default:
		if fieldType.Kind() == reflect.Struct {
			return deserialize2Struct(obj.(map[string]interface{}), reflect.New(fieldType))
		}
		return obj, nil
	}
}
