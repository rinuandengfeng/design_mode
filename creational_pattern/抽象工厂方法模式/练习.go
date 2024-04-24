package main

import (
	"fmt"
)

// 抽象层
/*
产品结构
*/

// 抽象显卡
type AbstractGraphicsCard interface{
	Display()
}

// 抽象内存
type AbstractMemory interface{
	Storage() bool
}

// 抽象CPU
type AbstractCPU interface{
	Calculate()
}


/*
产品族
*/
// 抽象创建工厂
type AbstractFactory interface{
	CreateGraphicsCard()AbstractGraphicsCard
	CreateMemory()AbstractMemory
	CreateCPU()AbstractCPU
}

// 实现类
// intel
// 显卡
type IntelGraphicsCard struct{}
// 实现抽象显卡的功能
func (igc *IntelGraphicsCard)Display(){
	fmt.Println("我是因特尔显卡")
}

// 内存
type IntelMemory struct{}
// 实现内存的抽象功能
func(im *IntelMemory)Storage()bool{
	fmt.Println("我是因特尔的内存")
	return true
}

// CPU
type IntelCPU struct{}
func(Icpu *IntelCPU)Calculate(){
	fmt.Println("我是因特尔的CPU")
}


// Nvdia
// 显卡
type NvdiaGraphicsCard struct{}
// 实现抽象显卡的功能
func (ngc *NvdiaGraphicsCard)Display(){
	fmt.Println("我是英伟达显卡")
}

// 内存
type NvdiaMemory struct{}
// 实现内存的抽象功能
func(nm *NvdiaMemory)Storage()bool{
	fmt.Println("我是英伟达的内存")
	return true
}

// CPU
type NvdiaCPU struct{}
func(ncpu *NvdiaCPU)Calculate(){
	fmt.Println("我是英伟达的CPU")
}


// Kingston
// 显卡
type KingstonGraphicsCard struct{}
// 实现抽象显卡的功能
func (kgc *KingstonGraphicsCard)Display(){
	fmt.Println("我是Kingston显卡")
}

// 内存
type KingstonMemory struct{}
// 实现内存的抽象功能
func(km *KingstonMemory)Storage()bool{
	fmt.Println("我是Kingston的内存")
	return true
}

// CPU
type KingstonCPU struct{}
func(kcpu *KingstonCPU)Calculate(){
	fmt.Println("我是Kingston的CPU")
}


// intel厂商工厂
type IntelFactory struct{}

// 创建intel显卡
func(inf *IntelFactory)CreateGraphicsCard()AbstractGraphicsCard{
	var intelGraphicsCard AbstractGraphicsCard

	// 创建intel显卡
	intelGraphicsCard = new(IntelGraphicsCard)

	return intelGraphicsCard
}

// 创建intel内存
func(inf *IntelFactory)CreateMemory()AbstractMemory{
	var intelMemory AbstractMemory

	// 创建intel内存
	intelMemory = new(IntelMemory)
	return intelMemory
}

// 创建intelcpu
func(inf *IntelFactory)CreateCPU()AbstractCPU{
	var intelCpu AbstractCPU

	// 创建intelCPU
	intelCpu = new(IntelCPU)
	return intelCpu
}

// 英伟达厂商工厂
type NvidiaFactory struct{}

// 创建英伟达显卡
func(nf *NvidiaFactory)CreateGraphicsCard()AbstractGraphicsCard{
	var nvidiaGraphicsCard AbstractGraphicsCard

	// 创建英伟达显卡
	nvidiaGraphicsCard = new(NvdiaGraphicsCard)

	return nvidiaGraphicsCard
}

// 创建英伟达内存
func(nf *NvidiaFactory)CreateMemory()AbstractMemory{
	var nvidiaMemory AbstractMemory

	// 创建英伟达内存
	nvidiaMemory = new(NvdiaMemory)
	return nvidiaMemory
}

// 创建英伟达cpu
func(nf *NvidiaFactory)CreateCPU()AbstractCPU{
	var nvidiaCpu AbstractCPU

	// 创建英伟达CPU
	nvidiaCpu = new(NvdiaCPU)
	return nvidiaCpu
}

// Kingston厂商工厂
type KingstonFactory struct{}

// 创建Kingston显卡
func(kf *KingstonFactory)CreateGraphicsCard()AbstractGraphicsCard{
	var kingstonGraphicsCard AbstractGraphicsCard

	// 创建Kingston显卡
	kingstonGraphicsCard = new(KingstonGraphicsCard)

	return kingstonGraphicsCard
}

// 创建Kingston内存
func(kf *KingstonFactory)CreateMemory()AbstractMemory{
	var kingstonMemory AbstractMemory

	// 创建Kingston内存
	kingstonMemory = new(KingstonMemory)
	return kingstonMemory
}

// 创建Kingston cpu
func(kf *KingstonFactory)CreateCPU()AbstractCPU{
	var kingstonCpu AbstractCPU

	// 创建KingstonCPU
	kingstonCpu = new(KingstonCPU)
	return kingstonCpu
}


// 业务层
func main(){
	// 组装1台 intel的CPU，intel的显卡，intel的内存
	// 1. 创建intel工厂
	var intelFactory AbstractFactory
	intelFactory = new(IntelFactory)

	// 获取intelCpu
	var intelCPU AbstractCPU
	intelCPU = intelFactory.CreateCPU()
	intelCPU.Calculate()

	// 创建intel显卡
	var intelGraphCard AbstractGraphicsCard
	intelGraphCard = intelFactory.CreateGraphicsCard() 
	intelGraphCard.Display()

	// 创建intel内存
	var intelMemory AbstractMemory
	intelMemory = intelFactory.CreateMemory()
	intelMemory.Storage()

	// 组装1台 Intel的CPU， nvidia的显卡，Kingston的内存
	var intelCPU2 AbstractCPU
	// 创建intelCPU
	intelCPU2 = intelFactory.CreateCPU()
	intelCPU2.Calculate()

	// 创建nvidia厂商工厂
	var nvidiaFactory AbstractFactory
	nvidiaFactory = new(NvidiaFactory)

	// 创建nvidia的显卡
	var navidiaGraphCard AbstractGraphicsCard
	navidiaGraphCard = nvidiaFactory.CreateGraphicsCard()
	navidiaGraphCard.Display()

	// 创建Kingston厂商工厂
	var kingstonFactory AbstractFactory
	kingstonFactory = new(KingstonFactory)

	// 创建Kingston的内存
	var kingstonMemory AbstractMemory
	kingstonMemory = kingstonFactory.CreateMemory()
	kingstonMemory.Storage()

}


