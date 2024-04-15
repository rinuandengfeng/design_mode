package main

import "fmt"

// 创建猫的类
type Cat struct {
}

// 实现吃饭的方法
func (c *Cat) Eat() {
	fmt.Println("小猫正在吃饭...")
}

// 添加睡觉的功能，（使用继承的方式）
type CatB struct {
	// 继承Cat类
	Cat
}

// 实现睡觉的方法
func (cb *CatB) Sleep() {
	fmt.Println("小猫正在睡觉...")
}

// 添加睡觉功能，（使用组合的方式）
type CatC struct {
	C *Cat
}

// 吃饭功能
//func (cc *CatC) Eat() {
//	cc.C.Eat()
//}

// 或者这样实现吃饭功能  这样比上面的依赖性更小
func (cc *CatC) Eat(c *Cat) {
	c.Eat()
}

// 实现睡觉功能
func (cc *CatC) Sleep() {
	fmt.Println("小猫正在睡觉....")
}

func main() {

	c := &Cat{}
	c.Eat()

	fmt.Println("-----------")
	cb := &CatB{}
	cb.Eat()
	cb.Sleep()

	fmt.Println("------------")
	cc := &CatC{}
	cc.Eat(c)
	cc.Sleep()
}
