package main

import "fmt"

// 子系统A
type SubSystemA struct{}

func (ssa *SubSystemA) MethodA() {
	fmt.Println("我是方法A")
}

// 子系统B
type SubSystemB struct{}

func (ssb SubSystemB) MethodB() {
	fmt.Println("我是方法B")

}

// 子系统C
type SubSystemC struct{}

func (ssc *SubSystemC) MethodC() {
	fmt.Println("我是方法C")
}

// 外观模式
type Facade struct {
	a *SubSystemA
	b *SubSystemB
	c *SubSystemC
}

func (f *Facade) MethodOne() {
	f.a.MethodA()
	f.c.MethodC()
}

func (f *Facade) MethodTwo() {
	f.a.MethodA()
	f.b.MethodB()
}

func (f Facade) Method3() {
	f.b.MethodB()
	f.c.MethodC()
}

func main() {

	f := Facade{
		a: new(SubSystemA),
		b: new(SubSystemB),
		c: new(SubSystemC),
	}

	f.MethodOne()
	fmt.Println("---------")
	f.MethodTwo()
	fmt.Println("---------")
	f.Method3()
}
