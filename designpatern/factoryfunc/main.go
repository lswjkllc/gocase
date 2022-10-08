package main

import "fmt"

/*
水果:
	苹果
	香蕉
	梨
*/

// 抽象层
type Fruit interface {
	Show()
}

type AbstractFactory interface {
	CreateFruit() Fruit
}

// 实现层
type Apple struct{}

func (s *Apple) Show() {
	fmt.Println("I'm apple")
}

type AppleFactory struct{}

func (s *AppleFactory) CreateFruit() Fruit {
	var fruit Fruit
	fruit = new(Apple)
	return fruit
}

type Banana struct{}

func (s *Banana) Show() {
	fmt.Println("I'm banana")
}

type BananaFactory struct{}

func (s *BananaFactory) CreateFruit() Fruit {
	var fruit Fruit
	fruit = new(Banana)
	return fruit
}

type Pear struct{}

func (s *Pear) Show() {
	fmt.Println("I'm pear")
}

type PearFactory struct{}

func (s *PearFactory) CreateFruit() Fruit {
	var fruit Fruit
	fruit = new(Pear)
	return fruit
}

// 业务逻辑层
func main() {
	var factory AbstractFactory
	var fruit Fruit

	// 1-apple 工厂
	factory = new(AppleFactory)
	// 1-生产 apple
	fruit = factory.CreateFruit()
	fruit.Show()

	// 2-banana 工厂
	factory = new(BananaFactory)
	// 2-生产 banana
	fruit = factory.CreateFruit()
	fruit.Show()

	// 3-pear 工厂
	factory = new(PearFactory)
	// 3-生产 pear
	fruit = factory.CreateFruit()
	fruit.Show()
}
