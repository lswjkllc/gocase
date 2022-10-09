package main

import "fmt"

/*
路边撸串烧烤场景:
	命令(抽象): 烤羊肉(具体)、烤鸡翅(具体)
	命令接受者(Receiver): 烤串师傅
	命令调用者(Invoker): 服务员
*/

// ----- 抽象层 -----
type BBCommand interface {
	Execute()
}

// ----- 实现层 -----
// 核心结构: 烤串师傅
type Cooker struct{}

func (s *Cooker) RoastLamb() {
	fmt.Println("roast lamb")
}

func (s *Cooker) RoastChickenwing() {
	fmt.Println("roas chicken wing")
}

// 具体命令
/* 烤羊肉命令 */
type RoastLambCmd struct {
	c *Cooker
}

func NewRoastLambCmd(c *Cooker) *RoastLambCmd {
	return &RoastLambCmd{c: c}
}

func (s *RoastLambCmd) Execute() {
	s.c.RoastLamb()
}

/* 烤鸡翅命令 */
type RoastChickenwingCmd struct {
	c *Cooker
}

func NewRoastChickenwingCmd(c *Cooker) *RoastChickenwingCmd {
	return &RoastChickenwingCmd{c: c}
}

func (s *RoastChickenwingCmd) Execute() {
	s.c.RoastChickenwing()
}

// 命令调用者: 服务员
type Waiter struct {
	Cmds []BBCommand
}

func (s *Waiter) Notify() {
	if s == nil {
		return
	}
	if len(s.Cmds) == 0 {
		return
	}

	for i := 0; i < len(s.Cmds); i++ {
		s.Cmds[i].Execute()
	}
}

// ----- 业务层 -----
func main() {
	cooker := new(Cooker)
	roastLambCmd := NewRoastLambCmd(cooker)
	roastChickenwing := NewRoastChickenwingCmd(cooker)

	waiter := new(Waiter)
	waiter.Cmds = append(waiter.Cmds, roastLambCmd)
	waiter.Cmds = append(waiter.Cmds, roastChickenwing)

	waiter.Notify()
}
