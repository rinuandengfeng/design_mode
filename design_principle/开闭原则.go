package main

import "fmt"

// 抽象工人类
type AbstractWorker interface {
	Work()
}

// 海底捞工作
type HaiDiLaoWorker struct {
	// AbstractWorker 的子类
}

// 实现接口的方法 多态
func (hw *HaiDiLaoWorker) Work() {
	fmt.Println("在海底捞工作...")
}

// 仓库工作
type CangKuWorker struct {
	// AbstractWorker 的子类
}

// a实现接口的方法，实现多态
func (cw *CangKuWorker) Work() {
	fmt.Println("在仓库工作...")
}

// 阿里工作
type AliWorker struct {
	// AbstractWorker 的子类
}

// 实现接口的方法，实现多态
func (aw *AliWorker) Work() {
	fmt.Println("在阿里工作...")
}

// 增加新功能  在腾讯工作
type TengXunWoker struct{}

// 实现接口的方法，实现多态
func (tw *TengXunWoker) Work() {
	fmt.Println("在腾讯工作...")
}

// 添加架构  实现架构层 （基于抽象层进行业务封装 -- 针对 interface 接口进行封装）
func XiaoMingWork(work AbstractWorker) {

	// 通过接口向下调用，（多态现象）
	work.Work()
}

func main() {

	/*
		// 在海底捞工作
		hw := HaiDiLaoWorker{}
		hw.Work()

		// 在仓库工作
		cw := CangKuWorker{}
		cw.Work()

		// 阿里工作
		aw := AliWorker{}
		aw.Work()

		// 增加新功能  腾讯工作
		tw := TengXunWoker{}
		tw.Work()
	*/

	// 在海底捞工作
	XiaoMingWork(&HaiDiLaoWorker{})

	// 在仓库工作
	XiaoMingWork(&CangKuWorker{})

	// 在阿里工作
	XiaoMingWork(&AliWorker{})

	// 在腾讯工作
	XiaoMingWork(&TengXunWoker{})
}
