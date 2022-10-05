package tag

import (
	"convertPlus/main/FieldHandler"
	"reflect"
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

func (convertTag *ConvertTag) SetFieldMeta(sourceField reflect.StructField, targetField reflect.StructField) {
	convertTag.sourceField = sourceField
	convertTag.targetField = targetField
}

func (convertTag *ConvertTag) SetFieldHandler(handler FieldHandler.FieldHander) {
	convertTag.fieldHandler = handler
}

func GetConvertTag(sourceField reflect.StructField) (convertTag ConvertTag) {
	return convertTag
}

// 字段元数据
type FieldMetaData struct {
	//字段绑定的处理器
	fieldHandler FieldHandler.FieldHander
	//源字段
	sourceField reflect.StructField
	//目标字段
	targetField reflect.StructField
}

type A struct {
	Name string `convert:cached=true,order=0,`
	Age  int    `convert:targetName=Age2,cached=false,exceptionHandler=ageExceptionHandler,handler=MysqlHandler`
}
