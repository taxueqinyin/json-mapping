package json_mapping

import (
	"encoding/json"
	"errors"
	"github.com/tidwall/gjson"
	"log"
	"reflect"
	"strings"
)

const tagKey = "JsonMapping"

var (
	ErrInvalidStructure = errors.New("invalid structure")
	ErrStructureIsEmpty = errors.New("structure is empty")
	ErrCustomType       = errors.New("custom types are not supported")
	ErrIncorrectJson    = errors.New("json format is incorrect")
	ErrUnSupportType    = errors.New("types other than struct are not supported")
	ErrSrcNotAddress    = errors.New("src is not addressable")
)

//MappingStruct param src must be a structure pointer. The tag with label is 'label:path',
//and the tag without label is 'path'. If it is not found, it will degenerate into
//package encoding/json decoding
func MappingStruct(searchLabel string, jsn []byte, src any) error {
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
	//尝试自带库解析，忽略可能存在的错误类型解析
	_ = json.Unmarshal(jsn, src)
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
	if !srcValue.CanAddr() {
		return ErrSrcNotAddress
	}
	fieldNum := srcType.NumField()
	if fieldNum < 1 {
		return ErrStructureIsEmpty
	}
	//遍历字段,查找tag
FieldLoop:
	for i := 0; i < fieldNum; i++ {
		field := srcValue.Field(i)
		if !field.CanAddr() {
			continue
		}
		tagStr, ok := srcType.Field(i).Tag.Lookup(tagKey)
		if !ok {
			continue
		}
		var tagPath string
		//对tagKey的路径进行切割
		tagStrs := strings.Split(tagStr, ";")
		for _, path := range tagStrs {
			//找到searchKey对应的字符串
			keyDivs := strings.Split(path, ":")
			keyDivsLength := len(keyDivs)
			if keyDivsLength > 2 {
				//tagStr语法错误，略过本字段查找下一个字段
				continue FieldLoop
			}
			//不设置搜索标签，查找不带标签的
			if searchLabel == "" {
				if len(keyDivs) == 1 {
					tagPath = keyDivs[0]
					break
				} else {
					continue
				}
			}
			//设置了搜索标签，查找标签对应的路径
			if keyDivs[0] == searchLabel {
				tagPath = keyDivs[1]
				break
			}
			//找不到，下一个标签
			continue
		}
		//检索了本字段的所有标签，依然找不到符合要求的，不赋值，下一个字段
		if tagPath == "" {
			continue
		}
		jsRawRes := parsed.Get(tagPath)
		//在json里找不到path对应的值
		if !jsRawRes.Exists() {
			continue
		}
		jsRawStr := jsRawRes.Raw
		err = json.Unmarshal([]byte(jsRawStr), field.Addr().Interface())
		if err != nil {
			return err
		}
	}
	return err
}
