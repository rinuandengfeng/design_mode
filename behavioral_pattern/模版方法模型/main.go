package main

import "fmt"

// 抽象接口
type Beverage interface {
	// 煮开水
	BoilWater()

	// 冲泡
	Brew()

	// 倒入杯中
	PourInCup()

	// 添加佐料
	AddThings()

	// 添加佐料的HOOK
	WantAddTings() bool
}

// 模版
type Template struct {
	b Beverage
}

// MakeBeverage 封装固定制作饮料的流程
func (t *Template) MakeBeverage() {
	if t == nil {
		return
	}

	// 制作流程  // 真实执行子类具体的实现
	t.b.BoilWater() // 烧水
	t.b.Brew()      // 冲泡
	t.b.PourInCup() // 倒入杯中
	// 根据用户需求 添加佐料
	if t.b.WantAddTings() {
		// 添加佐料
		t.b.AddThings()
	}
}

type MakeCoffee struct {
	Template //继承模版
}

func (mc *MakeCoffee) BoilWater() {
	fmt.Println("将水煮到100摄氏度")
}

func (mc *MakeCoffee) Brew() {
	fmt.Println("用水冲咖啡豆")
}
func (mc *MakeCoffee) PourInCup() {
	fmt.Println("将充好的咖啡倒入陶瓷杯中")
}
func (mc *MakeCoffee) AddThings() {
	fmt.Println("添加牛奶喝糖")
}

func (mc *MakeCoffee) WantAddTings() bool {
	return true
}

func NewMakeCoffee() *MakeCoffee {
	makeCoffee := new(MakeCoffee)

	// b是Beverage，是MakeCoffee的接口，这里需要给接口赋值，让b指向具体的子类
	// 来触发b的全部方法的多台特性
	makeCoffee.b = makeCoffee
	return makeCoffee
}

// 制作茶的流程
type MakeTea struct {
	Template
}

func (mb *MakeTea) BoilWater() {
	fmt.Println("将水煮到100摄氏度")
}

func (mb *MakeTea) Brew() {
	fmt.Println("用水冲茶")
}
func (mb *MakeTea) PourInCup() {
	fmt.Println("将冲好的茶倒入陶瓷杯中")
}
func (mb *MakeTea) AddThings() {
	fmt.Println("添加糖")
}
func (mb *MakeTea) WantAddTings() bool {
	return false
}

func NewMakeTea() *MakeTea {
	makeTea := new(MakeTea)

	// b是Beverage，是MakeCoffee的接口，这里需要给接口赋值，让b指向具体的子类
	// 来触发b的全部方法的多台特性
	makeTea.b = makeTea
	return makeTea
}

func main() {
	mc := NewMakeCoffee()
	mc.Template.MakeBeverage()

	fmt.Println("——---------------")
	mb := NewMakeTea()
	mb.Template.MakeBeverage()
}
