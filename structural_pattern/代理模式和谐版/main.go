package main

import "fmt"

// 抽象主题
type BeautyMoman interface {
	// 对男人抛媚眼
	MakeEyesWithMan()
	// 和男人共度美好的时光
	HappyWithMan()
}

// 具体主题
type PanJinLian struct {
}

// 对男人抛媚眼
func (p *PanJinLian) MakeEyesWithMan() {
	fmt.Println("潘金莲对本官抛媚眼")
}

// 和男人共度美好的时光
func (p *PanJinLian) HappyWithMan() {
	fmt.Println("潘金莲和本官共度美好时光")
}

// 王婆
type WangPo struct {
	Woman BeautyMoman
}

// 创建方法
func NewWangPo(woman BeautyMoman) BeautyMoman {
	return &WangPo{woman}
}

// 对男人抛媚眼
func (w *WangPo) MakeEyesWithMan() {
	w.Woman.MakeEyesWithMan()
}

// 和男人共度美好时光
func (w *WangPo) HappyWithMan() {
	w.Woman.HappyWithMan()
}

// 业务逻辑
func main() {

	pjl := new(PanJinLian)

	// 使用代理
	wp := NewWangPo(pjl)
	wp.MakeEyesWithMan()
	wp.HappyWithMan()
}
