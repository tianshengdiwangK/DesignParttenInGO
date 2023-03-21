package main

import "fmt"

//工厂方法相对于简单工厂来说多了一个抽象工厂角色
//抽象工厂是工厂方法的核心，任何工厂类都必须实现这个接口
//相对于抽象工厂，更符合开闭原则。
type Fruit interface {
	Show()
}

//抽象工厂
type AbstractFactory interface {
	CreateFruit() Fruit
}
type Apple struct {
}
type Banana struct {
}

func (apple *Apple) Show() {
	fmt.Println("this is apple")
}
func (banana *Banana) Show() {
	fmt.Println("this is banana")
}

//工厂角色
type AppleFactory struct {
}

func (appleFactory *AppleFactory) CreateFruit() Fruit {
	apple := new(Apple)
	return apple
}

type BananaFactory struct {
}

func (bananaFactory *BananaFactory) CreateFruit() Fruit {
	banana := new(Banana)
	return banana
}

func main() {
	//想要一个具体的苹果对象
	//依赖倒转原则，依托于接口编程
	//相对于简单工厂模式，更符合开闭原则，当新增一个水果类时，不需要修改原有的代码，只需要新增代码
	var appleFactory AbstractFactory
	appleFactory = new(AppleFactory)
	appleFactory.CreateFruit().Show()
}
