package main

import (
	"fmt"
	"sync"
)

type singeltion struct{}

// 初始一个全局的实例
var instance *singeltion

// 创建one
var once sync.Once

// 实现单例的方法
func GetInstance() *singeltion {
	once.Do(func() {
		instance = new(singeltion)
	})

	return instance
}

// 方法
func (s *singeltion) SomeThing() {
	fmt.Println("单例模式的方法")
}

func main() {
	s1 := GetInstance()
	s1.SomeThing()

	s2 := GetInstance()

	if s1 == s2 {
		fmt.Println("s1 == s2")
	}
}
