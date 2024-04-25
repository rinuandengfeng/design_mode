package main

import "fmt"

// 适配目标
type V5 interface {
	Use5V()
}

// 被适配的角色，适配者
type V220 struct{}

func (v22 *V220) Use220V() {
	fmt.Println("使用220v的电压")
}

// 适配器
type Adapter struct {
	v200 *V220
}

// 初始化方法
func NewAdapter(v220 *V220) *Adapter {
	return &Adapter{v220}
}

func (a *Adapter) Use5V() {
	fmt.Println("使用适配器进行充电...")

	// 调用适配器的方法
	a.v200.Use220V()

}

// 业务类
type Phone struct {
	v V5
}

func NewPhone(v V5) *Phone {
	return &Phone{v}
}

// 普通业务逻辑
func (p *Phone) Charge() {
	fmt.Println("Phone进行了充电....")
	p.v.Use5V()
}

// 普通业务功能

func main() {
	p := NewPhone(NewAdapter(new(V220)))
	p.Charge()
}
