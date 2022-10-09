package main

import "fmt"

/*
冲泡饮料: 咖啡和茶
公有步骤:
	煮水
	冲泡
	倒入杯中
	添加佐料(可选)
*/

// ----- 抽象层 -----
// 抽象饮料
type Beverage interface {
	BoilWater()
	Brew()
	PourIntoCup()
	AddSeasoning()
	WantSeasoning() bool
}

// 抽象模板
// 模板类
type template struct {
	b Beverage // 包裹抽象接口
}

// 模板方法
func (s *template) MakeBeverage() {
	if s == nil {
		return
	}
	if s.b == nil {
		return
	}

	s.b.BoilWater()
	s.b.Brew()
	s.b.PourIntoCup()
	if s.b.WantSeasoning() {
		s.b.AddSeasoning()
	}
}

// ----- 实现层 -----
// 茶
type Tea struct {
	template // 继承模板类
}

func NewTea() *Tea {
	tea := &Tea{}
	tea.b = tea
	return tea
}

func (s *Tea) BoilWater() {
	fmt.Println("boil water to 100 celsius")
}

func (s *Tea) Brew() {
	fmt.Println("brew tea with 100 celsius water")
}

func (s *Tea) PourIntoCup() {
	fmt.Println("pour tea water into a cup")
}

func (s *Tea) AddSeasoning() {
	fmt.Println("add solt into tea water")
}

func (s *Tea) WantSeasoning() bool {
	return true
}

// 咖啡(结构同上, 修改对应接口方法即可)

// ----- 业务层 -----
func main() {
	tea := NewTea()
	tea.MakeBeverage()
}
