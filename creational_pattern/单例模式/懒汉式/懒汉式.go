package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 定义结果体
type singeltion struct{}

// 定义一个互斥锁
var lock sync.Mutex

// 声明实例
var instance *singeltion

// 方式 1 获取实例 使用锁的方式
// func GetInstance() *singeltion {
// 	// 加锁
// 	lock.Lock()
// 	defer lock.Unlock()
// 	if instance == nil {
// 		instance = &singeltion{}
// 	}
// 	return instance
// }

// 使用原子操作 创建一个标记
var initialized uint32

func GetInstance() *singeltion {
	if atomic.LoadUint32(&initialized) == 1 {
		return instance
	}

	// 如果没有，则加锁进行创建
	lock.Lock()

	// 执行完毕关闭锁
	defer lock.Unlock()

	// 只有首次GetInstance()方法被调用，才会生成这个实例
	if instance == nil {
		instance = new(singeltion)
		// 设置标记为1
		atomic.StoreUint32(&initialized, 1)
	}

	return instance
}

// singeltion的方法
func (s *singeltion) SomeThing() {
	fmt.Println("单例的一个方法")
}

func main() {
	s1 := GetInstance()
	s1.SomeThing()

	s2 := GetInstance()
	// 判断两个实例是否相等
	if s1 == s2 {
		fmt.Println("s1 == s2")
	}
}
