package adapter

/**
适配器模式用于转换一种接口适配另一种接口。
实际使用中Adaptee一般为接口，并且使用工厂函数生成实例。
在Adapter中匿名组合Adaptee接口，所以Adapter类也拥有SpecificRequest实例方法，又因为Go语言中非入侵式接口特征，其实Adapter也适配Adaptee接口。
*/

// 适配的目标接口
type Target interface {
	Request() string
}

//  被适配的的目标接口
type Adaptee interface {
	SpecificRequest() string
}

// 被适配的接口的工厂函数
func NewAdaptee() Adaptee {
	return &adapteeImpl{}
}

type adapteeImpl struct{}

func (a *adapteeImpl) SpecificRequest() string {
	return "adaptee method"
}

func NewAdapter(adaptee Adaptee) Target {
	return &adapter{
		Adaptee: adaptee,
	}
}

type adapter struct {
	Adaptee
}

func (a *adapter) Request() string {
	return a.SpecificRequest()
}
