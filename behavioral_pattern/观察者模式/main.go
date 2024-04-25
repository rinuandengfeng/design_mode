package main

import "fmt"

// ----------抽象层----------
// 抽象的观察者
type Listener interface {
	OnTeacherComming() //观察者得到通知后要触发具体动作
}

type Notifier interface {
	AddListenner(listener Listener)
	RemoveListener(listener Listener)
	Notifier()
}

// ---------实现层---------
// 观察者具体的学生
type StuZhang3 struct {
	Badthing string
}

func (s *StuZhang3) OnTeacherComming() {
	fmt.Println("张3 停止  ", s.Badthing)
}
func (s *StuZhang3) DoBadthing() {
	fmt.Println("张4 正在", s.Badthing)
}

// 赵四
type StuLi4 struct {
	Badthing string
}

func (l *StuLi4) OnTeacherComming() {
	fmt.Println("李四 停止  ", l.Badthing)
}

func (l *StuLi4) DoBadthing() {
	fmt.Println("李四 正在", l.Badthing)
}

// 王五
type StuWang5 struct {
	Badthing string
}

func (w *StuWang5) OnTeacherComming() {
	fmt.Println("王五 停止  ", w.Badthing)
}

func (w *StuWang5) DoBadthing() {
	fmt.Println("王五 正在", w.Badthing)
}

// 统治者班长
type ClassMonitor struct {
	listenerList []Listener // 需要通知的全部观察者集合
}

func (m *ClassMonitor) AddListenner(listener Listener) {
	m.listenerList = append(m.listenerList, listener)
}
func (m *ClassMonitor) RemoveListener(listener Listener) {
	for index, l := range m.listenerList {
		// 找到要删除的元素位置
		if listener == l {
			// 将删除的元素的前后位置链接起来
			m.listenerList = append(m.listenerList[:index], m.listenerList[index+1:]...)
		}
	}
}

func (m *ClassMonitor) Notifier() {
	for _, listener := range m.listenerList {
		listener.OnTeacherComming() // 多态现象
	}
}
func main() {

	s1 := &StuZhang3{
		Badthing: "抄作业",
	}

	s2 := &StuLi4{
		Badthing: "玩王者荣耀",
	}

	s3 := &StuWang5{
		Badthing: "看赵4玩王者荣耀",
	}

	classMonitor := new(ClassMonitor)
	classMonitor.AddListenner(s1)
	classMonitor.AddListenner(s2)
	classMonitor.AddListenner(s3)

	fmt.Println("上课了，但是老师没有来，学生们都开始干坏事")
	s1.DoBadthing()
	s2.DoBadthing()
	s3.DoBadthing()

	fmt.Println("这时候老师来了，班长给学生们使了个眼神,.....")
	classMonitor.Notifier()
}
