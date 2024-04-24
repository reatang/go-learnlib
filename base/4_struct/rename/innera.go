package rename

// 检测其他包的类型，是否可以因为添加了别名而添加方法

type InnerA struct {
	A string
}

func (a *InnerA) MyName() string {
	return a.A
}
