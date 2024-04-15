package main

import "fmt"

// 抽象类
type Fruit interface {
	Show()
}

// 实现类
// 苹果
type Apple struct {
}

// 实现方法
func (apple *Apple) Show() {
	fmt.Println("我是苹果")
}

// 香蕉
type Banana struct{}

// 实现方法
func (banana *Banana) Show() {
	fmt.Println("我是香蕉")
}

// 梨类
type Pear struct{}

// 实现方法
func (pear *Pear) Show() {
	fmt.Println("我是梨")
}

// + 葡萄
type Grape struct{}

// + 实现方法
func (grape *Grape) Show() {
	fmt.Println("我是葡萄")
}

// 简单工厂
type Factor struct{}

// 工厂方法
func (factor *Factor) CreateFruit(kind string) Fruit {
	var fruit Fruit

	if kind == "apple" {
		fruit = new(Apple)
	} else if kind == "banana" {
		fruit = new(Banana)
	} else if kind == "pear" {
		fruit = new(Pear)
	} else if kind == "grape" {
		// + 添加新的功能，需要修改原始的代码
		fruit = new(Grape)
	}
	return fruit
}

func main() {
	// 创建工厂实例
	f := &Factor{}

	// 苹果
	apple := f.CreateFruit("apple")
	apple.Show()

	// 香蕉
	banana := f.CreateFruit("banana")
	banana.Show()

	// 梨
	pear := f.CreateFruit("pear")
	pear.Show()

	// 葡萄
	grape := f.CreateFruit("grape")
	grape.Show()
}
