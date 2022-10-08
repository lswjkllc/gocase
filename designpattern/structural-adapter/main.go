package main

import "fmt"

/*
适配目标:
	5V 充电
适配者:
	220V 电源
*/

// ----- 抽象层 -----
type V5 interface {
	Use5V()
}

// ----- 实现层 -----
// 适配者
type V220 struct{}

func (s *V220) Use220V() {
	fmt.Println("use 220v charge")
}

// 适配器
type Adapter struct {
	v220 *V220
}

func NewAdapter(v220 *V220) *Adapter {
	return &Adapter{v220: v220}
}

func (s *Adapter) Use5V() {
	fmt.Println("transfor...")
	s.v220.Use220V()
}

// 使用者
type Phone struct {
	v5 V5
}

func NewPhone(v5 V5) *Phone {
	return &Phone{v5: v5}
}

func (s *Phone) Charge() {
	fmt.Println("start charge with 5v...")
	s.v5.Use5V()
}

// ----- 业务层 -----
func main() {
	phone := NewPhone(NewAdapter(new(V220)))
	phone.Charge()
}
