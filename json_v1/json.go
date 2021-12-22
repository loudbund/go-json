package json_v1

import (
	"encoding/json"
	"errors"
	jsonIter "github.com/json-iterator/go"
	log "github.com/sirupsen/logrus"
	"reflect"
	"strconv"
	"strings"
)

/**-------------------------
// 名称：map参数转换为字符串
// @maps={"ip": "127.0.0.1", "device": "ABESSF0023"}
// @maps=[{"ip": "127.0.0.1", "device": "ABESSF0023"}]
***-----------------------*/
func JsonEncode(maps interface{}) (string, error) {
	var (
		jsonIterator = jsonIter.ConfigCompatibleWithStandardLibrary
		b            []byte
		err          error
	)
	b, err = jsonIterator.Marshal(maps)
	if err != nil {
		// fmt.Printf("Marshal with error: %+v\n", err)
		return "", err
	}
	return string(b), nil
}

/**-------------------------
// 名称：获取将字符参数转换为map的参数
// @jsonStr=`{"ip": "127.0.0.1", "device": "ABESSF0023"}`
// @jsonStr=`[{"ip": "127.0.0.1", "device": "ABESSF0023"}]`
***-----------------------*/
func JsonDecode(jsonStr string) (interface{}, error) {
	var (
		d   interface{}
		err error
	)
	decoder := jsonIter.NewDecoder(strings.NewReader(jsonStr))
	decoder.UseNumber()
	err = decoder.Decode(&d)
	if err != nil {
		// fmt.Printf("Unmarshal with error: %+v\n", err)
		return nil, err
	}
	return d, nil
}

// JSON格式相关获取函数1: 获取一个节点内容
func GetJsonInterface(dJson interface{}, keys ...interface{}) (interface{}, error) {
	if len(keys) < 1 {
		return dJson, nil
	}
	kFinds := make([]string, 0)
	for i := 0; i < len(keys); i++ {
		k := keys[i]
		if reflect.TypeOf(k).Name() == "string" { // 针对string格式参数
			K := k.(string)
			kFinds = append(kFinds, `["`+K+`"]`)

			if d, ok := dJson.(map[string]interface{}); ok {
				if v, ok := d[K]; ok {
					dJson = v
				} else {
					return nil, errors.New(strings.Join(kFinds, "") + "未找到2")
				}
			} else {
				return nil, errors.New(strings.Join(kFinds, "") + "未找到1")
			}

		} else if reflect.TypeOf(k).Name() == "int" { // 针对int参数

			K := k.(int)
			kFinds = append(kFinds, "["+strconv.Itoa(K)+"]")

			if d, ok := dJson.([]interface{}); ok {
				if len(d) > K {
					dJson = d[K]
				} else {
					return nil, errors.New(strings.Join(kFinds, "") + "未找到2")
				}
			} else {
				return nil, errors.New(strings.Join(kFinds, "") + "未找到1")
			}
		} else {
			log.Panic("GetJsonInterface 参数只接受string或int")
		}
	}
	return dJson, nil
}

// JSON格式相关获取函数2: 获取一个节点string格式
func GetJsonString(dJson interface{}, keys ...interface{}) (string, error) {
	d, err := GetJsonInterface(dJson, keys...)
	if err != nil {
		return "", err
	}
	if ret, ok := d.(string); ok {
		return ret, nil
	} else {
		return "", errors.New("不是string格式")
	}
}

// JSON格式相关获取函数3: 获取一个节点int64格式
// fmt.Println(getJsonInterface(J, "def", "bbb", 3, "errorcode"))
func GetJsonInt64(dJson interface{}, keys ...interface{}) (int64, error) {
	d, err := GetJsonInterface(dJson, keys...)
	if err != nil {
		return -1, err
	}
	if dNumber, ok := d.(json.Number); ok {
		if dInt64, ok := dNumber.Int64(); ok == nil {
			return dInt64, nil
		} else {
			return -1, errors.New("不是int格式")
		}
	} else {
		return -1, errors.New("不是int格式")
	}
}

// JSON格式相关获取函数4: 获取一个节点int64格式，如果是字符串数字，则转换成int64后返回
func GetJsonInt64Force(dJson interface{}, keys ...interface{}) (int64, error) {
	d, err := GetJsonInterface(dJson, keys...)
	if err != nil {
		return -1, err
	}
	if dString, ok := d.(string); ok {
		if dInt, err := strconv.ParseInt(dString, 10, 64); err != nil {
			return -1, errors.New("非int格式字串")
		} else {
			return dInt, nil
		}
	} else if dNumber, ok := d.(json.Number); ok {
		if dInt64, ok := dNumber.Int64(); ok == nil {
			return dInt64, nil
		} else {
			return -1, errors.New("不是int格式")
		}
	} else {
		return -1, errors.New("不是int格式")
	}
}
