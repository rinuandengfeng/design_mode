package main

import "fmt"

// 抽象层
// 抽象苹果
type AbstractApple interface{
	ShowApple()
}

// 抽象香蕉
type AbstractBanana interface{
	ShowBanana()
}

// 抽象梨
type AbstractPear interface{
	ShowPear()
}

// 工厂抽象类
type AbstractFactory interface{
	// 创建苹果
	CreateApple()AbstractApple
	// 创建香蕉
	CreateBanana()AbstractBanana
	// 创建梨
	CreatePear()AbstractPear
}



// 实现层
// 中国苹果
type ChinaApple struct{}

// 生产中国苹果
func(ca *ChinaApple)ShowApple(){
	fmt.Println("我是中国苹果")
}

// 生成中国香蕉
type ChinaBanana struct{}
func(cb *ChinaBanana)ShowBanana(){
	fmt.Println("我是中国香蕉")
}

// 生成中国梨
type ChinaPear struct{}
func(cp *ChinaPear)ShowPear(){
	fmt.Println("我是中国梨")
}

// 中国工厂
type ChinaFactory struct{}

// 创建中国苹果
func(cf *ChinaFactory) CreateApple()AbstractApple{
	var chinaApple AbstractApple

	// 创建中国苹果
	chinaApple = new(ChinaApple)

	return chinaApple
}

// 创建中国香蕉
func (cf *ChinaFactory)CreateBanana()AbstractBanana{
	var chinaBanana AbstractBanana

	// 创建中国香蕉
	chinaBanana = new(ChinaBanana)
	
	return chinaBanana
}


// 创建中国梨
func(cf *ChinaFactory) CreatePear() AbstractPear{
	var chinaPear AbstractPear

	// 创建中国梨
	chinaPear = new(ChinaPear)
	
	return chinaPear
}


// 日本
// 生成日本苹果
type JanpenApple struct{}
func(ja *JanpenApple) ShowApple(){
	fmt.Println("我是日本苹果")
}

// 生成日本香蕉
type JanpenBanana struct{}
func(jb *JanpenBanana) ShowBanana(){
	fmt.Println("我是日本香蕉")
}

// 生成日本梨
type JanpenPear struct{}
func(jb *JanpenPear) ShowPear(){
	fmt.Println("我是日本梨")
}

// 创建日本工厂
type JanpenFactory struct{}

// 创建日本苹果
func(jf *JanpenFactory)CreateApple()AbstractApple{
	var janpenApple AbstractApple

	// 创建日本苹果
	janpenApple = new(JanpenApple)

	return janpenApple
}

// 创建日本香蕉
func(jf *JanpenFactory)CreateBanana()AbstractBanana{
	var janpenBanana AbstractBanana

	// 创建日本香蕉
	janpenBanana = new(JanpenBanana)

	return janpenBanana
}

// 创建日本梨
func(jf *JanpenFactory)CreatePear()AbstractPear{
	var janpenPear AbstractPear

	// 创建日本梨
	janpenPear = new(JanpenPear)

	return janpenPear
}

// 逻辑层
 func main(){
	// 需求1 : 拿到中国生产的苹果、香蕉、梨
	var chinaFactory AbstractFactory
	// 创建中国工厂
	chinaFactory = new(ChinaFactory)

	// 生产中国苹果
	var chinaApple AbstractApple
	chinaApple = chinaFactory.CreateApple()
	chinaApple.ShowApple()

	// 生产中国香蕉
	var chinaBanana AbstractBanana
	chinaBanana = chinaFactory.CreateBanana()
	chinaBanana.ShowBanana()

	// 生产中国梨
	var chinaPear AbstractPear
	chinaPear = chinaFactory.CreatePear()
	chinaPear.ShowPear()



	// 需求2： 拿到日本生产的苹果、香蕉、梨
	//创建日本工厂
	var janpenFactory AbstractFactory
	janpenFactory = new(JanpenFactory)

	// 生产日本苹果
	var janpenApple AbstractApple
	janpenApple = janpenFactory.CreateApple()
	janpenApple.ShowApple()

	// 生产日本香蕉
	var janpenBanana AbstractBanana
	janpenBanana = janpenFactory.CreateBanana()
	janpenBanana.ShowBanana()

	// 生产日本梨
	var janpenPear AbstractPear
	janpenPear = janpenFactory.CreatePear()
	janpenPear.ShowPear()
 }