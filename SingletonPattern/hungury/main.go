package hungury

//含义是，在初始化单例唯一指针的时候，就已经提前开辟好了一个对象，申请了内存。
//饿汉式的好处是，不会出现线程并发创建，导致多个单例的出现，但是缺点是如果这个单例对象在业务逻辑没有被使用，也会客观的创建一块内存对象。

type singleton struct {
}

func (single *singleton) doSomething() {

}

var single *singleton = new(singleton)

func GetInstance() *singleton {
	return single
}

func main() {
	single := GetInstance()
	single.doSomething()
}
