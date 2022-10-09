package main

import "fmt"

/*
武林群侠传

观察者:
	丐帮:
		洪七公
		黄蓉
		乔峰
	明教:
		张无忌
		谢逊
		阳顶天
通知者:
	江湖百晓生
事件:
	`'江湖百晓生' 通知: '[丐帮]黄蓉' 消灭了 '[明教]张无忌'`
*/

type Event struct {
	Message string   // 消息: [丐帮]黄蓉 消灭了 [明教]张无忌
	Winner  Listener // 胜利者: [丐帮]黄蓉
	Loser   Listener // 失败者: [明教]张无忌
	Notifer Notifer  // 通知者: 江湖百晓生
}

// ----- 抽象层 -----
type Listener interface {
	GetName() string
	GetParty() string
	GetTitle() string
	Execute(event *Event)
}

type Notifer interface {
	AddListener(listener Listener)
	DelListener(listener Listener)
	Notify(event *Event)
}

// ----- 实现层 -----
// 观察者
type Hero struct {
	Name  string // 姓名
	Party string // 帮派
}

func (s *Hero) GetName() string {
	return s.Name
}

func (s *Hero) GetParty() string {
	return s.Party
}

func (s *Hero) GetTitle() string {
	return fmt.Sprintf("[%s]%s", s.Party, s.Name)
}

func (s *Hero) Execute(event *Event) {
	// 当事件为当事人时, 直接返回
	if s == event.Winner || s == event.Loser {
		return
	}

	// 如果和当前胜利者同属一个帮派, 拍手叫好
	if s.GetParty() == event.Winner.GetParty() {
		fmt.Printf("%s 得知 %s, 拍手叫好...\n", s.GetTitle(), event.Message)
		return
	}
	// 如果和当前失败者同属一个帮派, ...
	if s.GetParty() == event.Loser.GetParty() {
		fmt.Printf("%s 得知 %s, 准备报仇...\n", s.GetTitle(), event.Message)
		s.Fight(event.Winner, event.Notifer) // TODO 未避免发生重复循环, 此处方法需谨慎
		return
	}
}

func (s *Hero) Fight(other Listener, notifer Notifer) {
	// 产生事件
	event := &Event{
		Message: fmt.Sprintf("%s 消灭了 %s", s.GetTitle(), other.GetTitle()),
		Winner:  s,
		Loser:   other,
		Notifer: notifer,
	}
	// 广播事件
	fmt.Printf("[世界消息] %s ------------------------------\n", event.Message)
	notifer.Notify(event)
}

// 观察者
type Baixiaosheng struct {
	listeners []Listener
}

func (s *Baixiaosheng) AddListener(listener Listener) {
	// 检查观察者是否存在, 存在则返回
	for i := 0; i < len(s.listeners); i++ {
		if s.listeners[i] == listener {
			break
		}
	}
	// 添加观察者
	s.listeners = append(s.listeners, listener)
}

func (s *Baixiaosheng) DelListener(listener Listener) {
	for i := 0; i < len(s.listeners); i++ {
		if s.listeners[i] == listener {
			s.listeners = append(s.listeners[:i], s.listeners[i+1:]...)
			break
		}
	}
}

func (s *Baixiaosheng) Notify(event *Event) {
	// 删除失败的观察者
	s.DelListener(event.Loser) // TODO 此处如果不删除失败者, 会引起重复循环
	// 循环通知其他观察者
	for i := 0; i < len(s.listeners); i++ {
		s.listeners[i].Execute(event) // 该方法(Execute)可能会再次调用本方法(Notify)
	}
}

// ----- 业务层 -----
func main() {
	// 产生观察者
	hero1 := &Hero{"洪七公", "丐帮"}
	hero2 := &Hero{"黄蓉", "丐帮"}
	hero3 := &Hero{"乔峰", "丐帮"}
	hero4 := &Hero{"张无忌", "明教"}
	hero5 := &Hero{"谢逊", "明教"}
	hero6 := &Hero{"阳顶天", "明教"}
	// 产生通知者
	notifer := &Baixiaosheng{}
	// 向通知者中添加观察者
	notifer.AddListener(hero1)
	notifer.AddListener(hero2)
	notifer.AddListener(hero3)
	notifer.AddListener(hero4)
	notifer.AddListener(hero5)
	notifer.AddListener(hero6)
	// 初始状态
	fmt.Println("江湖一片太平...")
	// 触发事件
	hero1.Fight(hero4, notifer)
}
