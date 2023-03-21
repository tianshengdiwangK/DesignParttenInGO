package main

import "fmt"

//● 简单工厂模式：一个工厂负责创建所有产品
//  ○ 违反“开闭原则”，添加新产品需要修改工厂逻辑，工厂越来越复杂
//● 工厂方法模式：一个工厂创建一个产品
//  ○ 系统的可扩展性也就变得非常好，无需修改接口和原类
//  ○ 增加系统中类的个数，复杂度和理解度增加（一个具体产品就需要对应一个具体工厂）
//● 抽象方法模式：一个工厂创建一系列（一个产品族）的产品
//  ○ 增加新的产品族很方便，无须修改已有系统，符合“开闭原则”
//  ○ 增加新的产品等级结构麻烦，需要对原有系统进行较大的修改，违背了“开闭原则”
//  ○ 相当于在工厂方法模式的基础下进行了折中
//    ■ 对于产品族来说遵循了开闭原则
//    ■ 对于产品等级结构来说没有遵循开闭原则
//    ■ 如果产品结构等级稳定，那么就相当于完全遵循开闭原则
type Apple interface {
	ShowApple()
}
type Banana interface {
	ShowBanana()
}

type AbstractFactory interface {
	CreateApple() Apple
	CreateBanana() Banana
}

type ChinaApple struct {
}
type JapanApple struct {
}
type ChinaBanana struct {
}
type JapanBanana struct {
}

func (chinaApple *ChinaApple) ShowApple() {
	fmt.Println("this is china apple")
}
func (japanApple *JapanApple) ShowApple() {
	fmt.Println("this is japan Apple")
}
func (chinaBanana *ChinaBanana) ShowBanana() {
	fmt.Println("this is China Banana")
}
func (japanBanana *JapanBanana) ShowBanana() {
	fmt.Println("this is Japan Banana")
}

type ChinaFactory struct {
}

func (chinaFactory *ChinaFactory) CreateApple() Apple {
	return new(ChinaApple)
}
func (chinaFactory *ChinaFactory) CreateBanana() Banana {
	return new(ChinaBanana)
}

type JapanFactory struct {
}

func (japanFactory *JapanFactory) CreateApple() Apple {
	return new(JapanApple)
}
func (japanFactory *JapanFactory) CreateBanana() Banana {
	return new(JapanBanana)
}

func main() {
	var chinaFactory AbstractFactory
	chinaFactory = new(ChinaFactory)
	chinaFactory.CreateApple().ShowApple()
	chinaFactory.CreateBanana().ShowBanana()
}