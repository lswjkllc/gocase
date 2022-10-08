package main

import "fmt"

/*
手机:
	华为
	小米
装饰器:
	贴膜功能
	带壳功能
*/

// ----- 抽象层 -----
// 抽象的构件
type Phone interface {
	Call()
}

// 抽象的装饰器
// 由于 Golang 的 interface 无法包含成员变量, 此处抽象装饰器只能使用 struct 结构
type Decorator struct {
	phone Phone
}

func (s *Decorator) Call() {
	s.phone.Call()
}

// ----- 实现层 -----
// 具体的构件
type HuaweiPhone struct{}

func (s *HuaweiPhone) Call() {
	fmt.Println("use huawei phone call")
}

type XiaomiPhone struct{}

func (s *XiaomiPhone) Call() {
	fmt.Println("use xiaomi phone call")
}

// 具体的装饰器
type FilmDecorator struct {
	Decorator // 继承抽象装饰器
}

func NewFilmDecorator(phone Phone) Phone {
	return &FilmDecorator{Decorator: Decorator{phone: phone}}
}

func (s *FilmDecorator) Call() {
	s.phone.Call()
	fmt.Println("film for phone")
}

type ShellDecorator struct {
	Decorator // 继承抽象装饰器
}

func NewShellDecorator(phone Phone) Phone {
	return &ShellDecorator{Decorator: Decorator{phone: phone}}
}

func (s *ShellDecorator) Call() {
	s.phone.Call()
	fmt.Println("shell for phone")
}

// ----- 业务层 -----
func main() {
	var phone Phone

	fmt.Println("----- huawei -----")
	phone = new(HuaweiPhone)
	phone.Call()
	fmt.Println("------------------")
	phone = NewFilmDecorator(phone)
	phone.Call()
	fmt.Println("------------------")
	phone = NewShellDecorator(phone)
	phone.Call()

	fmt.Println("----- xiaomi -----")
	phone = new(XiaomiPhone)
	phone.Call()
	fmt.Println("------------------")
	phone = NewFilmDecorator(phone)
	phone.Call()
	fmt.Println("------------------")
	phone = NewShellDecorator(phone)
	phone.Call()
}
