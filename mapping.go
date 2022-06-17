package json_mapping

import (
	"encoding/json"
	"errors"
	"github.com/tidwall/gjson"
	"log"
	"reflect"
)

const tagKey = "JsonMapping"

var (
	ErrInvalidStructure = errors.New("invalid structure")
	ErrStructureIsEmpty = errors.New("structure is empty")
	ErrCustomType       = errors.New("custom types are not supported")
	ErrIncorrectJson    = errors.New("json format is incorrect")
	ErrUnSupportType    = errors.New("types other than struct are not supported")
)

// MappingStruct src应该为结构体指针
func MappingStruct(jsn []byte, src any) error {
	defer func() {
		if pan := recover(); pan != nil {
			log.Println(pan)
			return
		}
	}()
	//验证解析json
	if ok := gjson.Valid(string(jsn)); !ok {
		return ErrIncorrectJson
	}
	parsed := gjson.Parse(string(jsn))
	var err error
	//尝试自带库解析
	err = json.Unmarshal(jsn, src)
	if err != nil {
		return err
	}
	srcType := reflect.TypeOf(src)
	for srcType.Kind() == reflect.Ptr {
		srcType = srcType.Elem()
	}
	if srcType.Kind() != reflect.Struct {
		return ErrUnSupportType
	}
	srcValue := reflect.ValueOf(src)
	for srcValue.Kind() == reflect.Ptr {
		srcValue = srcValue.Elem()
	}
	fieldNum := srcType.NumField()
	if fieldNum < 1 {
		return ErrStructureIsEmpty
	}
	//遍历字段,查找tag
	for i := 0; i < fieldNum; i++ {
		if !srcValue.CanAddr() {
			continue
		}
		tagStr, ok := srcType.Field(i).Tag.Lookup(tagKey)
		if !ok {
			continue
		}
		//tagsStr := strings.Split(tagStr,";")
		//for _,v :=range tagsStr{
		//	jsRawStr := parsed.Get(v)
		//	json.Unmarshal()
		//}
		jsRawStr := parsed.Get(tagStr).Raw
		err = json.Unmarshal([]byte(jsRawStr), srcValue.Field(i).Addr().Interface())
		if err != nil {
			continue
		}
	}
	return nil
}
