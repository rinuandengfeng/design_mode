package main

import "fmt"

type Goods struct {
	Kind string // 商品的种类
	Fact bool   // 商品的真伪
}

// ======== 抽象层 =========
type Shopping interface {
	Buy(goods *Goods)
}

// ======== 实现层 =========
// 美国代理
type AmericaShopping struct{}

func (as *AmericaShopping) Buy(goods *Goods) {
	fmt.Println("美国代理，购买" + goods.Kind)
}

// KoreaShopping 韩国代理
type KoreaShopping struct{}

func (ks *KoreaShopping) Buy(goods *Goods) {
	fmt.Println("韩国代理，购买" + goods.Kind)
}

// AfricaShopping 非洲代理
type AfricaShopping struct{}

func (afs *AfricaShopping) Buy(goods *Goods) {
	fmt.Println("非洲代理，购买" + goods.Kind)
}

// ProxyShopping 代理
type ProxyShopping struct {
	shopping Shopping
}

func (p *ProxyShopping) Buy(goods *Goods) {
	// 1. 判断真伪
	if p.IsFact(goods) {
		// 2.购买商品
		p.shopping.Buy(goods)
		// 3. 过海关
		p.Customs(goods)
	}

}

// IsFact 判断真伪的方法
func (p *ProxyShopping) IsFact(goods *Goods) bool {
	if !goods.Fact {
		fmt.Println(goods.Kind + " 是假的，不建议购买")
		return false
	}
	return goods.Fact
}

// Customs 过海关
func (p *ProxyShopping) Customs(goods *Goods) {
	fmt.Println(goods.Kind + " 将商品带回祖国")
}

func NewProxyShopping(shopping Shopping) Shopping {
	return &ProxyShopping{shopping}
}

// ======== 业务逻辑层 =========
func main() {
	// 创建商品
	g1 := Goods{
		Kind: "韩国面膜",
		Fact: true,
	}

	g2 := Goods{
		Kind: "CET4",
		Fact: false,
	}

	// 创建美国代理

	as := new(AmericaShopping)

	ps := NewProxyShopping(as)
	ps.Buy(&g1)

	// 创建韩国代理

	ks := new(KoreaShopping)
	var ps2 Shopping

	ps2 = NewProxyShopping(ks)
	ps2.Buy(&g2)
	// 使用代理进行购买

	ps3 := NewProxyShopping(as)
	ps3.Buy(&g1)

}
