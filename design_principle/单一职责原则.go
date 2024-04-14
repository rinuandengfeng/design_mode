package main

import "fmt"

// 工作类
type ClothesWork struct{}

// 工作穿衣方式
func (cw *ClothesWork) Style() {
	fmt.Println("逛街的穿着")
}

// 逛街类
type ClothesShop struct{}

// 逛街的穿着
func (cs *ClothesShop) Style() {
	fmt.Println("逛街的穿着")
}

func main() {

	// 工作
	cw := ClothesWork{}
	cw.Style()

	// 逛街
	cs := ClothesShop{}
	cs.Style()
}
