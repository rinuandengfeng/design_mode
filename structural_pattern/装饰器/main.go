package main

import "fmt"

// 抽象层
type Phone interface {
	Show() // 构建的功能
}

// 抽象的装饰器，装饰器的基础类（该类本应该是interface，
// 但是Golang的interface语法不可以有成员属性）
type Decorator struct {
	phone Phone
}

func (d *Decorator) Show() {}

// 实现层

// 具体的构建
// 华为手机
type Huawei struct{}

func (h Huawei) Show() {
	fmt.Println("我是华为手机")
}

// 小米手机
type XiaoMi struct{}

func (x *XiaoMi) Show() {
	fmt.Println("我是小米手机")
}

// 具体的装饰器
// 贴膜的装饰器
type MoPhone struct {
	Decorator
}

// 初始化方法
func NewMoPhone(phone Phone) Phone {
	return &MoPhone{Decorator{phone}}
}

func (mp MoPhone) Show() {
	// 实现show方法
	mp.phone.Show()
	// ++++贴膜
	fmt.Println("进行了贴膜....")
}

// 手机壳装饰器
type KePhone struct {
	Decorator
}

// 初始化方法
func NewKePhone(phone Phone) Phone {
	return &KePhone{Decorator{phone}}
}

func (kp KePhone) Show() {
	// 实现show方法
	kp.phone.Show()

	// +++ 使用手机壳
	fmt.Println("安装了手机壳")
}

// 业务逻辑层
func main() {

	// 创建华为手机
	hw := new(Huawei)
	hw.Show()

	// 创建小米手机
	xm := new(XiaoMi)
	xm.Show()

	fmt.Println("----------使用装饰器----------")
	// 使用贴膜装饰器
	hwmp := NewMoPhone(hw)
	hwmp.Show()

	// 使用手机壳装饰器
	hwkp := NewKePhone(hw)
	hwkp.Show()

	// 即贴膜又使用壳
	fmt.Println("--------即进行贴膜又使用手机壳----------")
	hwmk := NewKePhone(hwmp)
	hwmk.Show()
}
