package main

import "fmt"

// ------>抽象层<------
// 车的抽象接口
type Car interface {
	// 开车父类，有一个Run方法
	Run()
}

// 驾驶员的抽象接口
type Driver interface {
	// 驾驶员父类,实现驾驶功能
	Drive(car Car)
}

// ------>实现层<------

// 车的子类
// 奔驰汽车
type Benz struct {
	// 奔驰汽车的属性
}

// 实现父类接口的方法
func (bz *Benz) Run() {
	fmt.Println("正在开奔驰...")
}

// 宝马汽车
type BWM struct {
	// 宝马汽车的属性
}

// 实现父类接口的方法
func (bwm *BWM) Run() {
	fmt.Println("正在开宝马...")
}

// 小米su7
type XiaoMisu7 struct {
	// 小米su7属性
}

// 实现父类接口的方法
func (xm *XiaoMisu7) Run() {
	fmt.Println("正在开小米su7...")
}

// 驾驶员的子类
// 张三
type ZhangSan struct {
	// 张三的属性
}

// 张三开车
func (zs *ZhangSan) Drive(car Car) {
	// 张三开车
	fmt.Println("我是驾驶员张三....")
	// 张三开的车
	car.Run()
}

// 李四
type LiSi struct {
	// 李四的属性
}

// 李四开车
func (ls *LiSi) Drive(car Car) {
	// 李四开车
	fmt.Println("我是驾驶员李四...")
	// 李四开的车
	car.Run()
}

// + 王五
type WangWu struct {
	// 王五的属性
}

// +王五驾驶汽车
func (w *WangWu) Drive(car Car) {
	// 王五正在开汽车
	fmt.Println("我是驾驶员王五...")
	// 王五开的车
	car.Run()
}

// ------>业务逻辑层<------
func main() {

	// 创建奔驰实例
	var benz Car
	benz = new(Benz)

	// 创建宝马实例
	var bwm Car
	bwm = new(BWM)

	// 张三实例
	var zs Driver
	zs = new(ZhangSan)
	// 张三开宝马
	zs.Drive(bwm)

	// 李四实例
	var ls Driver
	ls = new(LiSi)
	// 李四驾驶奔驰
	ls.Drive(benz)

	// + 添加小米su7
	var xm Car
	xm = new(XiaoMisu7)

	// + 添加王五实例
	var w5 Driver
	w5 = new(WangWu)
	w5.Drive(xm)
}
