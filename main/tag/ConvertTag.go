package tag

import (
	"convertPlus/main/FieldHandler"
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

// tag实体
type ConvertTag struct {
	// tag是否缓存
	Cached bool
	//tag 异常回调
	ExceptionCallBack error
	//tag 排序
	Order int
	//tag 别名
	Target string
	//tag 字段绑定的处理器名称
	FieldHandlerName string
	//字段元数据
	FieldMetaData
}

func getKeyAndValue(str string) (key string, value string) {
	// 支持等号和冒号
	denghao := regexp.MustCompile("=")
	isDenghao := denghao.FindStringIndex(str)
	maohao := regexp.MustCompile(":")
	isMaohao := maohao.FindStringIndex(str)

	// 等号
	if len(isDenghao) > 0 {
		return string([]byte(str)[0:isDenghao[0]]), string([]byte(str)[isDenghao[1]:len(str)])
	}
	if len(isMaohao) > 0 {
		// 冒号
		fmt.Println("isMaohao", isMaohao)
		return string([]byte(str)[0:isMaohao[0]]), string([]byte(str)[isMaohao[1]:len(str)])
	}

	return key, value
}

func GetConvertTag(sourceField reflect.StructField) (convertTag ConvertTag, err error) {
	tag := string(sourceField.Tag)
	// if head convert:
	convertRe := regexp.MustCompile("convert:")
	//isConvert := strings.Contains(string(tag), "convert:")
	//fmt.Println("test: ", tag)
	isConvert := convertRe.FindStringIndex(tag)
	fmt.Println("isConvert", isConvert)
	if len(isConvert) > 0 && isConvert[0] == 0 {
		// 处理
		convertTagStr := string([]byte(tag)[isConvert[1]:len(tag)])
		res1 := strings.Split(convertTagStr, ",")
		//fmt.Println("3636", res1)
		for i := 0; i < len(res1); i++ {
			key, value := getKeyAndValue(res1[i])
			fmt.Println("obj", key, value)
			//convertTag := &ConvertTag{
			//	key: value,
			//}
			//convertTag[key] := value
		}
		// convert:cached=true,order=0,defaultValue:{Age:123},doubleFormat

	} else {
		// 抛出异常
		return convertTag, errors.New("not have convert")
		//panic("convert format error")
	}
	return convertTag, nil
func (convertTag *ConvertTag) SetFieldMeta(sourceField reflect.StructField, targetField reflect.StructField) {
	convertTag.SourceField = sourceField
	convertTag.TargetField = targetField
}

func (convertTag *ConvertTag) SetFieldHandler(handler FieldHandler.FieldHander) {
	convertTag.FieldHandler = handler
}

func GetConvertTag(sourceField reflect.StructField) (convertTag ConvertTag) {
	return convertTag
}

// 字段元数据
type FieldMetaData struct {
	//字段绑定的处理器
	FieldHandler FieldHandler.FieldHander
	//源字段
	SourceField reflect.StructField
	//目标字段
	TargetField reflect.StructField
}

type A struct {
	Name string `convert:cached=true,order=0,`
	Age  int    `convert:targetName=Age2,cached=false,exceptionHandler=ageExceptionHandler,handler=MysqlHandler`
}
