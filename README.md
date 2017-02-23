# struct2lua

golang struct 结构体转成 lua文本文件

不支持struct直接嵌套的模式 例如:

type B struct {
	BB string
}

type A struct {
	AA B
}
