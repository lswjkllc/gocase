package main

import "fmt"

/*
产品:
	显卡: display
	内存: storage
	CPU: calculate
厂商:
	Intel
	Nvidia
	Kingston
*/

// 抽象层
type AbstractGraphcard interface {
	Display()
}

type AbstractMemory interface {
	Storage()
}

type AbstractCpu interface {
	Calculate()
}

type AbstractFactory interface {
	CreateGraphcard() AbstractGraphcard
	CreateMemory() AbstractMemory
	CreateCpu() AbstractCpu
}

// 实现层
/* Intel 产品族*/
type IntelGraphcard struct{}

func (s IntelGraphcard) Display() {
	fmt.Println("Intel graphcard")
}

type IntelMemory struct{}

func (s *IntelMemory) Storage() {
	fmt.Println("Intel memory")
}

type IntelCpu struct{}

func (s *IntelCpu) Calculate() {
	fmt.Println("Intel cpu")
}

type IntelFactory struct{}

func (s *IntelFactory) CreateGraphcard() AbstractGraphcard {
	var graphCard AbstractGraphcard
	graphCard = new(IntelGraphcard)
	return graphCard
}

func (s *IntelFactory) CreateMemory() AbstractMemory {
	var memory AbstractMemory
	memory = new(IntelMemory)
	return memory
}

func (s *IntelFactory) CreateCpu() AbstractCpu {
	var cpu AbstractCpu
	cpu = new(IntelCpu)
	return cpu
}

/* Nvidia 产品族*/
type NvidiaGraphcard struct{}

func (s NvidiaGraphcard) Display() {
	fmt.Println("Nvidia graphcard")
}

type NvidiaMemory struct{}

func (s *NvidiaMemory) Storage() {
	fmt.Println("Nvidia memory")
}

type NvidiaCpu struct{}

func (s *NvidiaCpu) Calculate() {
	fmt.Println("Nvidia cpu")
}

type NvidiaFactory struct{}

func (s *NvidiaFactory) CreateGraphcard() AbstractGraphcard {
	var graphCard AbstractGraphcard
	graphCard = new(NvidiaGraphcard)
	return graphCard
}

func (s *NvidiaFactory) CreateMemory() AbstractMemory {
	var memory AbstractMemory
	memory = new(NvidiaMemory)
	return memory
}

func (s *NvidiaFactory) CreateCpu() AbstractCpu {
	var cpu AbstractCpu
	cpu = new(NvidiaCpu)
	return cpu
}

/* Kingston 产品族*/
type KingstonGraphcard struct{}

func (s KingstonGraphcard) Display() {
	fmt.Println("Kingston graphcard")
}

type KingstonMemory struct{}

func (s *KingstonMemory) Storage() {
	fmt.Println("Kingston memory")
}

type KingstonCpu struct{}

func (s *KingstonCpu) Calculate() {
	fmt.Println("Kingston cpu")
}

type KingstonFactory struct{}

func (s *KingstonFactory) CreateGraphcard() AbstractGraphcard {
	var graphCard AbstractGraphcard
	graphCard = new(KingstonGraphcard)
	return graphCard
}

func (s *KingstonFactory) CreateMemory() AbstractMemory {
	var memory AbstractMemory
	memory = new(KingstonMemory)
	return memory
}

func (s *KingstonFactory) CreateCpu() AbstractCpu {
	var cpu AbstractCpu
	cpu = new(KingstonCpu)
	return cpu
}

// 电脑
type Computer struct {
	Graphcard AbstractGraphcard
	Memory    AbstractMemory
	Cpu       AbstractCpu
}

func (s Computer) Show() {
	s.Graphcard.Display()
	s.Memory.Storage()
	s.Cpu.Calculate()
}

// 业务逻辑层
func main() {
	// Intel 工厂
	var intelFac AbstractFactory
	intelFac = new(IntelFactory)
	// Nvidia 工厂
	var nvidiaFac AbstractFactory
	nvidiaFac = new(NvidiaFactory)
	// Kingston 工厂
	var kingstonFac AbstractFactory
	kingstonFac = new(KingstonFactory)

	// Intel 显卡、Intel 内存、Intel CPU
	cpr1 := Computer{
		Graphcard: intelFac.CreateGraphcard(),
		Memory:    intelFac.CreateMemory(),
		Cpu:       intelFac.CreateCpu(),
	}
	cpr1.Show()
	// Nvidia 显卡、Kingston 内存、Intel CPU
	cpr2 := Computer{
		Graphcard: nvidiaFac.CreateGraphcard(),
		Memory:    kingstonFac.CreateMemory(),
		Cpu:       intelFac.CreateCpu(),
	}
	cpr2.Show()
}
