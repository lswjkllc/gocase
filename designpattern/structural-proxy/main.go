package main

import "fmt"

/*
王婆代理潘金莲和人约会

被代理者: 潘金莲
代理者: 王婆
*/

// ----- 接口层 -----
type BeautyWoman interface {
	MakeEyesWithMan()
	MakeDateWithMan()
}

// ----- 实现层 -----
// 被代理者
type PanJinLian struct{}

func (s *PanJinLian) MakeEyesWithMan() {
	fmt.Println("panjinlian make eyes with man")
}

func (s *PanJinLian) MakeDateWithMan() {
	fmt.Println("panjinlian make date with man")
}

// 代理者
type WangPoProxy struct {
	woman BeautyWoman
}

func NewWangPoProxy(woman BeautyWoman) BeautyWoman {
	return &WangPoProxy{woman: woman}
}

func (s *WangPoProxy) MakeEyesWithMan() {
	s.woman.MakeEyesWithMan()
}

func (s *WangPoProxy) MakeDateWithMan() {
	s.woman.MakeDateWithMan()
}

// ----- 业务层 -----
func main() {
	wangpo := NewWangPoProxy(new(PanJinLian))
	wangpo.MakeEyesWithMan()
	wangpo.MakeDateWithMan()
}
