package main

import "fmt"

/*
商场促销策略:
	上午: 策略A，打8折
	下午: 策略B，满200减100

抽象策略接口
具体策略A、策略B
使用环境类
*/

// ----- 抽象层 -----
type SellStrategy interface {
	GetSellPrice(price float64) float64
}

// ----- 实现层 -----
// 策略A
type StrategyA struct{}

func (s *StrategyA) GetSellPrice(price float64) float64 {
	fmt.Println("use strategy a: hit 8-fold")
	return price * 0.8
}

// 策略B
type StrategyB struct{}

func (s *StrategyB) GetSellPrice(price float64) float64 {
	fmt.Println("use strategy b: full 200 minus 100")
	if price >= 200 {
		price -= 100
	}
	return price
}

// 环境类: 商品
type Goods struct {
	Price    float64
	Strategy SellStrategy
}

func (s *Goods) SetStrategy(strategy SellStrategy) {
	s.Strategy = strategy
}

func (s *Goods) GetPrice() float64 {
	fmt.Println("origin price", s.Price, "")
	return s.Strategy.GetSellPrice(s.Price)
}

// ----- 业务层 -----
func main() {
	goods := &Goods{Price: 200.0}

	goods.SetStrategy(new(StrategyA))
	fmt.Println("sell price on am:", goods.GetPrice())
	fmt.Println("------------")
	goods.SetStrategy(new(StrategyB))
	fmt.Println("sell price on pm:", goods.GetPrice())
}
