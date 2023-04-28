package main

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
)

// 原文: https://studygolang.com/articles/35401

// 工作流结构: 开始节点、边、结束节点
type WorkFlow struct {
	alreadyDone bool
	done        chan struct{}
	s           *sync.Once
	root        *Node
	end         *Node
}

func NewWorkFlow() *WorkFlow {
	wf := &WorkFlow{
		root: NewNode(nil),
		done: make(chan struct{}, 1),
		s:    &sync.Once{},
	}
	endTask := &EndAction{done: wf.done, s: wf.s}
	wf.end = NewNode(endTask)

	return wf
}

func (s *WorkFlow) StartWithContext(ctx context.Context, i interface{}) {
	s.root.ExecuteWithContext(ctx, s, i)
}

func (s *WorkFlow) WaitDone() {
	<-s.done
	close(s.done)
}

func (s *WorkFlow) interrupDone() {
	s.alreadyDone = true
	s.s.Do(func() { s.done <- struct{}{} })
}

func (s *WorkFlow) AddEdge(from *Node, to *Node) {
	AddEdge(from, to)
}

func (s *WorkFlow) ConnectToRoot(node *Node) {
	AddEdge(s.root, node)
}

func (s *WorkFlow) ConnectToEnd(node *Node) {
	AddEdge(node, s.end)
}

// 边: 从哪个节点来、到哪个节点去
type Edge struct {
	from *Node
	to   *Node
}

// 节点: 依赖边（入边）、当前任务、孩子边（出边）
type Node struct {
	dependencies   []*Edge
	dependCompletd int32
	task           Runnable
	children       []*Edge
}

func NewNode(task Runnable) *Node {
	return &Node{task: task}
}

func (s *Node) ExecuteWithContext(ctx context.Context, wf *WorkFlow, i interface{}) {
	// 判断当前节点前置依赖是否完成
	if !s.dependencyHasDone() {
		return
	}
	// 如果发生了错误，终止流程的执行
	if ctx.Err() != nil {
		wf.interrupDone()
		return
	}
	// 执行当前任务
	if s.task != nil {
		s.task.Run(i)
	}

	// 执行后续任务
	if len(s.children) > 0 {
		// 其他每个任务使用单独的线程完成
		for i := 1; i < len(s.children); i++ {
			go func(edge *Edge, i interface{}) {
				edge.to.ExecuteWithContext(ctx, wf, i)
			}(s.children[i], i)
		}
		// 当前线程执行第一个任务
		s.children[0].to.ExecuteWithContext(ctx, wf, i)
	}
}

func (s *Node) dependencyHasDone() bool {
	// 不存在前置节点（边），不需要等待，直接返回
	if len(s.dependencies) == 0 {
		return true
	}
	// 只有一个前置节点（边），不需要等待，直接返回
	if len(s.dependencies) == 1 {
		return true
	}
	// 将依赖节点（边）+1，表明有一个依赖的节点（边）完成了
	atomic.AddInt32(&s.dependCompletd, 1)

	return s.dependCompletd == int32(len(s.dependencies))
}

// 抽象任务接口
type Runnable interface {
	Run(i interface{})
}

// 结束任务
type EndAction struct {
	done chan struct{}
	s    *sync.Once
}

func (s *EndAction) Run(i interface{}) {
	fmt.Println("穿戴完成")
	s.s.Do(func() { s.done <- struct{}{} })
}

// 穿内裤
type PantiesAction struct{}

func (s *PantiesAction) Run(i interface{}) {
	fmt.Println("穿内裤...")
}

// 穿袜子
type SockAction struct{}

func (s *SockAction) Run(i interface{}) {
	fmt.Println("穿袜子...")
}

// 穿衬衣
type ShirtAction struct{}

func (s *ShirtAction) Run(i interface{}) {
	fmt.Println("穿衬衣...")
}

// 戴手表
type WatchAction struct{}

func (s *WatchAction) Run(i interface{}) {
	fmt.Println("带手表...")
}

// 穿裤子
type PantsAction struct{}

func (s *PantsAction) Run(i interface{}) {
	fmt.Println("穿裤子...")
}

// 穿鞋
type ShoeAction struct{}

func (s *ShoeAction) Run(i interface{}) {
	fmt.Println("穿鞋子...")
}

// 穿外套
type CoatAction struct{}

func (s *CoatAction) Run(i interface{}) {
	fmt.Println("穿外套...")
}

// 工具方法
func AddEdge(from *Node, to *Node) *Edge {
	// 初始化
	edge := &Edge{from: from, to: to}
	// 处理前置节点
	from.children = append(from.children, edge)
	// 处理后置节点
	to.dependencies = append(to.dependencies, edge)

	return edge
}

// 业务逻辑
func main() {
	// 初始化工作流
	ctx := context.Background()
	wf := NewWorkFlow()
	// 构建节点
	pantiesNode := NewNode(&PantiesAction{})
	sockNode := NewNode(&SockAction{})
	shirtNode := NewNode(&ShirtAction{})
	watchNode := NewNode(&WatchAction{})
	pantsNode := NewNode(&PantsAction{})
	shoeNode := NewNode(&ShoeAction{})
	coatNode := NewNode(&CoatAction{})
	// 构建节点之间的关系
	wf.ConnectToRoot(pantiesNode)
	wf.ConnectToRoot(sockNode)
	wf.ConnectToRoot(shirtNode)
	wf.ConnectToRoot(watchNode)

	wf.AddEdge(pantiesNode, pantsNode)
	wf.AddEdge(sockNode, shoeNode)
	wf.AddEdge(pantsNode, shoeNode)
	wf.AddEdge(shirtNode, coatNode)
	wf.AddEdge(watchNode, coatNode)

	wf.ConnectToEnd(shoeNode)
	wf.ConnectToEnd(coatNode)
	// 开始执行
	wf.StartWithContext(ctx, []string{})
	wf.WaitDone()
}
