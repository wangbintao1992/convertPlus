package tag

type ConvertTag struct {
	//是否缓存
	Cached bool
	//异常回调
	ExceptionCallBack error
	//排序
	Order int
	//别名
	Target string
}
