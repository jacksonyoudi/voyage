package simple_factory

/**
!简单工程模式
go 语言没有构造函数一说，所以一般会定义NewXXX函数来初始化相关类。 NewXXX 函数返回接口时就是简单工厂模式，也就是说Golang的一般推荐做法就是简单工厂。
在这个simplefactory包中只有API 接口和NewAPI函数为包外可见，封装了实现细节。
*/

type API interface {
	Say(name string) string
}

type HiAPI struct {
}

func (h *HiAPI) Say(name string) string {
	return "hi" + name
}

type HelloAPI struct{}

func (h *HelloAPI) Say(name string) string {
	return "hello" + name
}

// new API
func NewAPI(t int) API {
	if t == 1 {
		return &HiAPI{}
	}
	if t == 2 {
		return &HelloAPI{}
	}
	return nil
}
