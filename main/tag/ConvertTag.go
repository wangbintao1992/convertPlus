package tag

//tag实体
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

/*type A struct {
	反射tag
	Name     string `convert:cached=true,order=0,defaultValue:{Age:123},doubleFormat:0.00`
	Age      int    `convert:targetName=Age2,cached=false,exceptionHandler=ageExceptionHandler,handler=MysqlHandler`
	ClassAge []C    `convert`
}*/
