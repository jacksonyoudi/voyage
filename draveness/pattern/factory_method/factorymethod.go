package factory_method

/**
!工厂方法模式使用子类的方式延迟生成对象到子类中实现。
Go中不存在继承 所以使用匿名组合来实现
*/

type Operator interface {
	SetA(int)
	SetB(int)
	Result() int
}

type OperatorFactory interface {
	Create() Operator
}

type OperatorBase struct {
	a, b int
}

func (b *OperatorBase) SetA(i int) {
	b.a = i
}

func (b *OperatorBase) SetB(i int) {
	b.b = i
}

type PluginOperator struct {
	*OperatorBase
}

func (p *PluginOperator) Result() int {
	return p.a + p.b
}

type MinusOperator struct {
	*OperatorBase
}

func (p *MinusOperator) Result() int {
	return p.a - p.b
}

type PluginOperatorFactory struct {
}

func (f *PluginOperatorFactory) Create() Operator {
	return &PluginOperator{
		&OperatorBase{},
	}
}

type MinusOperatorFactory struct {
}

func (f *MinusOperatorFactory) Create() Operator {
	return &MinusOperator{
		&OperatorBase{},
	}
}
