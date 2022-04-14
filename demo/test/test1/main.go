package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	//defer_call()
	//pase_student()
	//addOne()
	//functionExtend()
	//ChanlePainc()
	//sliceInit()
	//mapLock()
	//safeSet()
	//speak()
	//nilInterface()
	setMap()
}

func defer_call() {
	//defer --先进后出，类似栈
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()

	panic("触发异常")
}


type student struct {
	Name string
	Age  int
}

func pase_student() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	//map循环时，value使用的是副本。所以m[stu.Name]=&stu实际上一致指向同一个指针， 最终该指针的值为遍历的最后一个struct的值拷贝
	for _, stu := range stus {
		fmt.Println(stu.Name)
		m[stu.Name] = &stu
	}
	fmt.Println(m["wang"], m["li"], m["zhou"])
	//正确写法。直接用index去map中取值，不仅可以避免上面的问题，而且每次循环都减少了一次复制操作
	for i := range stus {
		m[stus[i].Name] = &stus[i]
	}
	fmt.Println(m["wang"], m["li"], m["zhou"])
}

func addOne()  {
	runtime.GOMAXPROCS(1) //cpu1,保证所有goroutine按顺序执行
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("i: ", i)
			wg.Done()
		}()
	}
	fmt.Println("-----------------")
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("i: ", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func functionExtend()  {
	t := Teacher{}
	t.ShowA()
}
type People struct{}
func (p *People) ShowA() {
	fmt.Println("showA")
	p.ShowB() //因为这里调用的是people的showB方法，所以只会输出showB，不会被Teacher的showB方法重载
}
func (p *People) ShowB() {
	fmt.Println("showB")
}
type Teacher struct {
	People
}
func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}

func ChanlePainc() {
	runtime.GOMAXPROCS(1)
	int_chan := make(chan int, 1)
	string_chan := make(chan string, 1)
	int_chan <- 1
	string_chan <- "hello"
	//golang select当条件都满足时，会随机选择一条case执行，所以1和panic都可能出现
	select {
	case value := <-int_chan:
		fmt.Println(value)
	case value := <-string_chan:
		panic(value)
	}
}

func sliceInit() {
	//初始化一个长度为5的切片，值都默认为0值
	s := make([]int, 5)
	s = append(s, 1, 2, 3)
	//[0 0 0 0 0 1 2 3]
	fmt.Println(s)
}

func mapLock() {
	//如果声明时不初始化map，后面会报panic: assignment to entry in nil map
	//ua := UserAges{}
	ua := UserAges{ages: map[string]int{}}
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i< 1000; i++ {
			ua.Add("Tom", 11)
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i< 1000; i++ {
			fmt.Println(ua.Get("Tom"))
		}
	}()
	wg.Wait()
}
type UserAges struct {
	ages map[string]int
	sync.Mutex
}
func (ua *UserAges) Add(name string, age int) {
	ua.Lock()
	defer ua.Unlock()
	ua.ages[name] = age
}
func (ua *UserAges) Get(name string) int {
	//读取时如果不加锁，map也有可能出现线程安全问题，并发读写时报错：fatal error: concurrent map read and map write
	//ua.Lock()
	//defer ua.Unlock()
	if age, ok := ua.ages[name]; ok {
		return age
	}
	return -1
}

func safeSet()  {
	th:=threadSafeSet{
		s:[]interface{}{1,2},
	}
	v:=<-th.Iter()
	fmt.Println(fmt.Sprintf("%s%v","ch",v))
}
type threadSafeSet struct {
	sync.RWMutex
	s []interface{}
}
func (set *threadSafeSet) Iter() <-chan interface{} {
	ch := make(chan interface{}) // 解除注释看看！
	//ch := make(chan interface{},len(set.s))
	go func() {
		set.RLock()

		for elem,value := range set.s {
			ch <- elem
			fmt.Printf("Iter:%d, %v\n",elem,value)
		}

		close(ch)
		set.RUnlock()

	}()
	return ch
}

//考点：接口实现
func speak()  {
	//指针类型方法集合包含了某接口的所有方法, 则说该类型的指针类型实现了该接口，此时值类型不一定实现了该接口。
	//值类型方法集合包含了某接口的所有方法, 则说该类型的值类型实现了该接口，此时指针类型一定实现了该接口。
	//var peo Man = Stduent{}
	var peo Man = &Stduent{}
	think := "bitch"
	fmt.Println(peo.Speak(think))
}
type Man interface {
	Speak(string) string
}
type Stduent struct{}

func (stu *Stduent) Speak(think string) (talk string) {
	if think == "bitch" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	return
}

func nilInterface()  {
	if live() == nil {
		fmt.Println("AAAAAAA")
	} else {
		fmt.Println("BBBBBBB")
	}
}
type People2 interface {
	Show()
}
type Student struct{}
func (stu *Student) Show() {}
func live() People2 {
	var stu *Student
	return stu
}
/*
//可以看出iface比eface 中间多了一层itab结构。
itab 存储_type信息和[]fun方法集，从上面的结构我们就可得出，因为data指向了nil 并不代表interface 是nil， 所以返回值并不为空
type eface struct {      //空接口
	_type *_type         //类型信息
	data  unsafe.Pointer //指向数据的指针(go语言中特殊的指针类型unsafe.Pointer类似于c语言中的void*)
}
type iface struct {      //带有方法的接口
	tab  *itab           //存储type信息还有结构实现方法的集合
	data unsafe.Pointer  //指向数据的指针(go语言中特殊的指针类型unsafe.Pointer类似于c语言中的void*)
}
type _type struct {
    size       uintptr  //类型大小
    ptrdata    uintptr  //前缀持有所有指针的内存大小
    hash       uint32   //数据hash值
    tflag      tflag
    align      uint8    //对齐
    fieldalign uint8    //嵌入结构体时的对齐
    kind       uint8    //kind 有些枚举值kind等于0是无效的
    alg        *typeAlg //函数指针数组，类型实现的所有方法
    gcdata    *byte
    str       nameOff
    ptrToThis typeOff
}
type itab struct {
    inter  *interfacetype  //接口类型
    _type  *_type          //结构类型
    link   *itab
    bad    int32
    inhash int32
    fun    [1]uintptr      //可变大小 方法集合
}
 */

func setMap()  {
	//a := make(map[int]string)
	var a sync.Map
	var wg sync.WaitGroup
	for i := 0; i <= 10; i++ {
		wg.Add(1)
		go func() {
			defer  wg.Done()
			a[i] = "hello world"
		}()
	}
	wg.Wait()
}