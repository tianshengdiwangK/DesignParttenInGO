package main

import "fmt"

//简单工厂模式有三个角色
//工厂角色：负责创建所有实例，工厂类可以直接被外界调用，创建所需的产品对象
//抽象产品角色:工厂创建出所有对象的父类，负责描述所有实例所共有的公共接口
//具体产品:工厂所创建的具体对象

type Fruit interface { //抽象产品角色
	Show()
}

type Apple struct { //具体产品
}
type Banana struct {
}

func (apple *Apple) Show() {
	fmt.Println("this is apple")
}
func (banana *Banana) Show() {
	fmt.Println("this is banana")
}

type Factory struct {
}

func (factory *Factory) CreateFruit(name string) Fruit {
	var fruit Fruit
	if name == "apple" {
		fruit = new(Apple)
	} else {
		fruit = new(Banana)
	}
	return fruit
}

func main() {
	factory := new(Factory)
	factory.CreateFruit("apple").Show()
}
