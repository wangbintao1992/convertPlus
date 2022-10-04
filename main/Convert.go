package main

import (
	"convertPlus/main/Control"
	"convertPlus/main/tag"
	"reflect"
)

type Convert struct {
	//源对象类型
	SourceType reflect.Type
	//转换对象类型
	TargetType reflect.Type
	//处理器链
	ConvertTagChain []tag.ConvertTag
	//持有controller
	controller Control.Controller
}

//初始化
func (c *Convert) convert(source any, targetType reflect.Type) {
	c.TargetType = targetType
	c.SourceType = reflect.TypeOf(source)

	c.convert0(source)
}

//转换
func (c *Convert) convert0(source any) {
	c.scanSourceField()
}

func (c *Convert) scanSourceField() {

	//sourceFieldMap := make(map[string]reflect.StructField)

	//目标对象字段映射 map  fieldName -> field
	targetFieldMap := make(map[string]reflect.StructField)

	fieldNum := c.TargetType.NumField()

	for i := 0; i < fieldNum; i++ {
		tempField := c.TargetType.Field(i)
		targetFieldMap[tempField.Name] = tempField
	}

	sourceFieldNum := c.SourceType.NumField()

	for i := 0; i < sourceFieldNum; i++ {
		//sourceTempField := c.SourceType.Field(i)
	}
}
