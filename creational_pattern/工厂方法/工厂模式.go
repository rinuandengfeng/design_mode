package main

import "fmt"

// 抽象层
//水果类
type AbstractFruit interface{
	Show()
}

// 工厂类
type AbstractFactory interface{
	CreateFruit() AbstractFruit
}


// 基础类
// 苹果
type Apple struct{}

// 实现方法
func (apple *Apple) Show(){
	fmt.Println("我是apple")	
}

// 香蕉
type Banana struct{}

// 实现方法
func (banana *Banana) Show(){
	fmt.Println("我是banana")
}

// 梨
type Pear struct{}

// 实现方法
func (pear *Pear) Show(){
	fmt.Println("我是pear")
}


// 具体工厂类
// 苹果类的工厂
type AppleFactory struct{}

// 工厂方法
func (af *AppleFactory) CreateFruit() AbstractFruit{

	// 创建一个抽象的fruit
	var fruit AbstractFruit

	fruit = new(Apple)
	return fruit
}

// 香蕉类的工厂
type BananaFactory struct{}

// 工厂方法
func(bf *BananaFactory) CreateFruit()AbstractFruit{
	//创建抽象的fruit
	var fruit AbstractFruit

	fruit = new(Banana)

	return fruit
}

// 梨类工厂
type PearFactory struct{}

// 工厂方法
func (pf *PearFactory) CreateFruit() AbstractFruit{
	
	// 创建抽象的fruit
	var fruit AbstractFruit

	fruit = new(Pear)

	return fruit
}



func main(){

	// 获取一个苹果
	var applefac  AbstractFactory

	// 获取苹果工厂
	applefac = new(AppleFactory)

	// 创建苹果
	apple := applefac.CreateFruit()

	apple.Show()


	// 创建梨的工厂
	var pearFac AbstractFactory

	// 生成梨工厂
	pearFac  = new(PearFactory)

	// 生产梨
	pear := pearFac.CreateFruit()

	// 打印梨
	pear.Show()

	// 创建香蕉的工厂
	var bananaFac AbstractFactory

	bananaFac = new(BananaFactory)

	// 生产香蕉的工厂
	banana  := bananaFac.CreateFruit()

	// 打印香蕉
	banana.Show()
	
}