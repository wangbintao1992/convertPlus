package main

import (
	"convertPlus/main/FieldHandler"
	"reflect"
)

type Convert struct {
	//源对象类型
	SourceType reflect.Type
	//转换对象类型
	TargetType reflect.Type
	//处理器链
	fieldHandlerChain []FieldHandler.FieldHander
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
	fieldNum := c.SourceType.NumField()

	for i := 0; i < fieldNum; i++ {
		//转换tag
	}
}
