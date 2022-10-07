package main

import (
	"convertPlus/main/Control"
	"convertPlus/main/FieldHandler"
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

	fieldHandlerCache map[string]FieldHandler.FieldHander
}

func (c *Convert) registHandler(handler ...FieldHandler.FieldHander) {
	for _, r := range handler {
		c.fieldHandlerCache[reflect.TypeOf(r).Name()] = r
	}

}

// 初始化
func (c *Convert) convert(source any, target any) any {
	c.TargetType = reflect.TypeOf(target).Elem()
	c.SourceType = reflect.TypeOf(source).Elem()

	c.scanSourceField()

	//TODO初始化控制器 context初始化

	return c.convert0(source, target)
}

// 转换
func (c *Convert) convert0(source any, target any) any {
	c.controller.Control(source, target, c.ConvertTagChain)

	return target
}

func (c *Convert) scanSourceField() {

	//目标对象字段映射 map  fieldName -> field
	targetFieldMap := make(map[string]reflect.StructField)

	fieldNum := c.TargetType.NumField()

	for i := 0; i < fieldNum; i++ {
		tempField := c.TargetType.Field(i)
		targetFieldMap[tempField.Name] = tempField
	}

	sourceFieldNum := c.SourceType.NumField()

	for i := 0; i < sourceFieldNum; i++ {
		field := c.SourceType.Field(i)
		convertTag := tag.GetConvertTag(field)

		//是否目标对象有映射字段
		if targetField, ok := targetFieldMap[convertTag.Target]; ok {

			convertTag.SetFieldMeta(field, targetField)

			if targetHandler, findHandler := c.fieldHandlerCache[convertTag.FieldHandlerName]; findHandler {
				convertTag.SetFieldHandler(targetHandler)
			}
		}
	}
}

//func main() {
//	var u User
//	t := reflect.TypeOf(u)
//	for i := 0; i < t.NumField(); i++ {
//		sf := t.Field(i)
//		tag := string(sf.Tag)
//		// if head convert:
//		convertRe := regexp.MustCompile("convert:")
//		//isConvert := strings.Contains(string(tag), "convert:")
//		//fmt.Println("test: ", string(tag))
//		isConvert := convertRe.FindStringIndex(tag)
//		fmt.Println("isConvert", isConvert[0])
//		if len(isConvert) > 0 && isConvert[0] == 0 {
//			// 处理
//			fmt.Println("56", tag)
//		} else {
//			// 抛出异常
//			panic("convert format error")
//		}
//	}
//}
//
//type User struct {
//	Name string `1convert:cached=true,order=0,defaultValue:{Age:123},doubleFormat:0.00`
//	Age  int    `json:"age" bson:"b_age"`
//}
