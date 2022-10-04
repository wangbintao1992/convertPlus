package FieldHandler

//字段处理器顶级接口
type FieldHander interface {
	//前置
	Before(source any) any
	//抽象转换
	convert(source any) any
	//后置
	After(value any) any
}
