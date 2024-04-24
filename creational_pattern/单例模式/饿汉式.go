package main

import "fmt"

/*
三个要点：
		一个某个类只要有一个实例
        二它必须自行创建这个实例
        三是它必须自行向整个系统提供这个实例
*/

/*
   保证一个类永远只能有一个对象，这个对象还能被系统的其他模块使用
*/

// 1.保证这个类非公有化，外界不能通过这个类 直接创建一个对象
//
//	那么这个类就应该变得非公有访问，类的名字小写
type singelton struct{}

//  2. 但是还要有一个指针可以指向这个唯一对象，但是这个指针永远不能改变方向
//     Golang中没有常指针概念，所以只能通过这个指针私有化不让外部模块访问
var instance *singelton = new(singelton)

// 3. 如果全部都是私有化，那么外部模块将永远无法访问到这个对象
// 所以需要提供一个方法来获取这个对象
func GetInstance() *singelton {
	return instance
}

func (s *singelton) SomeThing() {
	fmt.Println("单例的一个方法")
}

func main() {
	s1 := GetInstance()
	s1.SomeThing()

	s2 := GetInstance()

	if s1 == s2 {
		fmt.Println("s1 == s2")
	}
}
