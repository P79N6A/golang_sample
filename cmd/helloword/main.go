package main

import (
	"bufio"
	"bytes"
	"container/heap"
	"container/list"
	"container/ring"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"index/suffixarray"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"reflect"
	"regexp"
	"sort"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/lxb31/golang_sample/internal/utils"
	"github.com/lxb31/golang_sample/pkg/hello"
)

const (
	_ = iota
	// WINDOWS = 1
	WINDOWS
	// LINUX = 2
	LINUX
	// MAC = 3
	MAC
)

const (
	// TYPE1 = 1
	TYPE1 = 1 << iota
	// TYPE2 = 2
	TYPE2
	// TYPE3 = 4
	TYPE3
)

var (
	a = hello.Vertex{}
	b = hello.Vertex{X: 1, Y: 2}
	c = hello.Vertex{Y: 1}
	d = &hello.Vertex{X: 1, Y: 2}
)

//closure
//Go 函数可以是闭包的。闭包是一个函数值，它来自函数体的外部的变量引用。
//函数可以对这个引用值进行访问和赋值；换句话说这个函数被“绑定”在这个变量上。
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func handle(x interface{}) {
	time.Sleep(10 * time.Millisecond)
	return
}

// I is a interface
type I interface {
	Echo(string) string
}

// A A is a struct implements I
type A struct {
	Name   string `tag1:"tagvalue1" tag2:"tagvalue2"`
	Number int32
}

// Echo implements I
func (a *A) Echo(s string) string {
	return fmt.Sprintf("%v,%v: %v", a.Name, a.Number, s)
}

// Noop just use varialbe
func Noop(...interface{}) {}

func main() {
	msg := hello.NewMessage("hello")
	//msg.
	r, err := hello.SayHello(msg)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(&r)
	}

	// internal
	utils.InternalSayHello()

	{
		/*
			[值类型]
			整形:  int8、byte、int16、uint、uintprt等，零值为0
			浮点: float32、float64，零值为0
			布尔: bool, 零值为false
			复数: complex64、complex128，零值为0+0i
			字符串: string, 零值为”“
			复合类型: 数组和struct, 递归赋零值
			[指针和引用] nil
			pointer, function, interface, error, slice, map, channel
		*/
		//slice内部有3个变量, 指向数组的指针/len/cap, 只有这3个变量都为零值时, slice才是零值
		//通过声明得到的是零值
		var s []string
		fmt.Println(len(s), s, s == nil)
		//通过赋值得到的不是零值
		var s1 = []string{}
		fmt.Println(len(s1), s1, s1 == nil)
	}
	// const
	{
		fmt.Println("[const] =====")
		fmt.Printf("WINDOWS=%v, LINUX=%v, MAC=%v \n", WINDOWS, LINUX, MAC)
		fmt.Printf("TYPE1=%v, TYPE2=%v, TYPE3=%v \n", TYPE1, TYPE2, TYPE3)
	}

	//Stringer
	v := hello.Vertex{X: 1, Y: 3}
	fmt.Println(v.String())
	fmt.Println(&v) //Print严格遵守接受者的方法集, 而不会进行(&v).Sting()的语法糖扩展

	/*
	* time
	 */
	{
		time := time.Now()
		fmt.Println(time)
		fmt.Println(time.Second())
		fmt.Println(time.Unix())
		fmt.Println(time.UnixNano() / 1000000)
	}

	/*
	* array
	 */
	{
		a1 := [2]string{}
		a2 := [2]string{"hello", "world"}
		a3 := [...]string{"s1", "s2", "s3"}

		a1[0] = "hello"
		a1[1] = "world"
		fmt.Printf("a1=%v, a2=%v, a3=%v, len(a1)=%v, len(a2)=%v\n", a1, a2, a3, len(a1), len(a2))
	}

	/*
	* slice
	 */
	{
		s1 := []int{}
		s2 := []int{1, 2, 3, 4}
		s3 := make([]int, 4) //len=4, cap=4

		s1 = append(s1, 1, 2)
		s3[0] = 1
		s3[1] = 2
		s3 = append(s3, s2...)

		// 浅拷贝
		s4 := s3
		// 深拷贝
		s5 := make([]int, len(s3))
		copy(s5, s3)

		fmt.Printf("s1=%v, s2=%v, s3=%v, s4=%v, s5=%v\n", s1, s2, s3, s4, s5)
		fmt.Printf("len(s1)=%v, len(s2)=%v\n", len(s1), len(s2))
		fmt.Printf("cap(s1)=%v, cap(s2)=%v\n", cap(s1), cap(s2))
		fmt.Printf("s2[2:3]=%v\n", s2[2:3])
	}

	/*
	* range
	 */
	{
		s := []int{1, 2, 3, 4}
		for i, v := range s {
			fmt.Println(i, v)
		}
		for _, v := range s {
			fmt.Println(v)
		}
	}

	/*
	* map
	 */
	{
		m1 := map[string]*hello.Vertex{}
		m2 := map[string]*hello.Vertex{
			"key1": {X: 10, Y: 11},
			"key2": {X: 30, Y: 40},
		}
		m3 := make(map[string]*hello.Vertex)

		//赋值
		m3["key1"] = &hello.Vertex{X: 1, Y: 2}
		m3["key2"] = &hello.Vertex{X: 3, Y: 4}
		//删除
		delete(m3, "key1")
		//遍历
		for key, value := range m3 {
			fmt.Println(key, value)
		}
		//判断是否存在
		if elem, ok := m3["key1"]; ok {
			Noop(elem)
		}
		//判断存在并且非零值
		if elem := m3["key1"]; elem != nil {
			Noop(elem)
		}
		fmt.Println(m1, m2, m3)
	}

	// Heap
	{
		h := &hello.IntHeap{2, 5, 3, 7, 1}
		heap.Init(h)
		heap.Push(h, 10)
		fmt.Println(h)
		heap.Pop(h)
		fmt.Println(h)
	}
	//list
	{
		fmt.Println("[list] =====")
		l := list.New()
		l.PushFront(1)
		l.PushFront(3)
		l.PushFront(5)
		l.Front()
		l.PushBack(2)
		l.PushBack(4)
		l.PushBack(6)
		l.Back()
		l.Len()
		for e := l.Front(); e != nil; e = e.Next() {
			fmt.Println(e.Value)
		}
	}
	// ring
	{
		fmt.Println("[ring] =====")
		r := ring.New(5)
		n := r.Len()
		for i := 0; i < n; i++ {
			r.Value = i
			r = r.Next()
		}
		e := r
		for i := 0; i < n; i++ {
			fmt.Println(e.Value)
			e = e.Next()
		}
	}

	//function pointer
	{
		fmt.Println("[function pointer] =====")
		max := func(i1 int, i2 int) int {
			if i1 > i2 {
				return i1
			}
			return i2
		}
		imax := max(2, 3)
		fmt.Printf("imax=%v \n", imax)
	}

	//closure
	{
		counter := adder()
		for i := 0; i < 10; i++ {
			fmt.Println(counter(i))
		}
	}

	//type assertion
	{
		var i interface{} = "hello"
		s, ok := i.(string)
		fmt.Println(s, ok)
		f, ok := i.(float64)
		fmt.Println(f, ok)
	}

	//type switch
	{
		var v interface{} = "value"
		switch value := v.(type) {
		case int:
			fmt.Println("type int", value)
		case string:
			fmt.Println("type string", value)
		default:
			fmt.Println("type unknown")
		}

	}

	//error
	{
		err := hello.Run()
		if err != nil {
			fmt.Println(err)
		}
	}

	//Reader
	{
		reader := strings.NewReader("hello, world")
		bts := make([]byte, 8)
		for {
			n, err := reader.Read(bts)
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Println(err)
			}
			fmt.Printf("read %v bytes \n", n)
		}
	}

	// channel
	// range/close, select

	// sync map
	{
		m := sync.Map{}
		m.Store("key1", "value1")
		m.Store("key2", "value2")
		m.Store(3, 4.0)
		e, ok := m.Load("key1")
		e, ok = m.Load("key1")
		e, ok = m.Load("key1")
		fmt.Println(e, ok)
	}
	// sort
	{
		a := []int{2, 3, 4, 7, 3, 1, 66, 34, 8, 3, 9}
		sort.Sort(sort.IntSlice(a))
		sort.Ints(a)
		sort.Sort(sort.Reverse(sort.IntSlice(a)))
		sort.Stable(sort.IntSlice(a))

		persons := []hello.Person{
			{Name: "bob", Age: 21},
			{Name: "tom", Age: 22},
			{Name: "frank", Age: 23},
			{Name: "lxb31", Age: 21},
		}
		sort.Sort(hello.ByAge(persons))
		fmt.Println(persons)

		sort.Slice(persons, func(i, j int) bool {
			return persons[i].Age < persons[j].Age
		})
	}
	// sync.atomic
	{
		fmt.Println("[sync.atomic] =====")
		var i int32 = 1
		for {
			ok := atomic.CompareAndSwapInt32(&i, i, 2)
			if ok {
				break
			}
		}

		atomic.AddInt32(&i, 3)
		atomic.StoreInt32(&i, 10)
		i = atomic.LoadInt32(&i)

		v := new(atomic.Value)
		v.Store(hello.Vertex{X: 1, Y: 2})
		vertex := v.Load()
		fmt.Println(vertex)
	}
	// sync
	// sync.Mutex和sync.RWMutex不可重入!
	{
		fmt.Println("[sync] =====")
		//sync.Mutex
		m := new(sync.Mutex)
		m.Lock()
		m.Unlock()

		//sync.RWMutex
		rw := new(sync.RWMutex)
		rw.RLock()
		rw.RUnlock()
		rw.Lock()
		rw.Unlock()

		// 	//sync.Cond
		// 	//不建议使用Cond, 多数情况下应该使用channel
		// 	done := false
		// 	locker := new(sync.Mutex)
		// 	condition := sync.NewCond(locker)
		// 	go func() {
		// 		condition.L.Lock()
		// 		for !done {
		// 			condition.Wait()
		// 		}
		// 		fmt.Println("done")
		// 		condition.L.Unlock()
		// 	}()
		// 	go func() {
		// 		condition.L.Lock()
		// 		done = true
		// 		condition.Signal()
		// 		//condition.Broadcast()
		// 		condition.L.Unlock()
		// 	}()
		// 	time.Sleep(1 * time.Second)

		// 	//sync.Once
		// 	action := new(sync.Once)
		// 	action.Do(func() {
		// 		fmt.Println("do once")
		// 	})

		// 	// sync.Pool
		// 	// Pool用于存放已经分配但没有用的临时对象, 方便以后重复使用
		// 	// 在GC时Pool内的对象会被清空
		// 	// sync.Pool的定位不是做类似连接池的东西，它的用途仅仅是增加对象重用的几率，减少gc的负担
		// 	pool := &sync.Pool{
		// 		New: func() interface{} {
		// 			return fmt.Sprintf("%v", time.Now().Format("2006-01-02 15:04:05"))
		// 		},
		// 	}
		// 	pool.Put(time.Now().Format("2006-01-02 15:04:05"))
		// 	v := pool.Get().(string)
		// 	fmt.Println(v)

		// 	// sync.WaitGroup
		// 	// 类似Java的Phaser和CountDownLatch
		// 	wg := new(sync.WaitGroup)
		// 	for index := 0; index < 3; index++ {
		// 		wg.Add(1)
		// 		go func() {
		// 			time.Sleep(1 * time.Second)
		// 			wg.Done()
		// 		}()
		// 	}
		// 	wg.Wait()
	}
	// strings.Builder
	{
		builder := &strings.Builder{}
		builder.WriteString("hello")
		builder.WriteString("world")
		fmt.Fprintf(builder, "hello")
		fmt.Fprintf(builder, "world")
		fmt.Println(builder.String())
	}
	// gob serialization
	// struct序列化会存储名称和类型, 只要两个struct有交集即可交换
	// 名称不能冲突
	// 整形使用变长编码, 有符号整型之间兼容, 无符号整型之间兼容
	// 指针类型自动解析与赋值
	{
		type P struct {
			X, Y, Z int
			Name    string
		}

		type Q struct {
			X, Y *int32
			Name string
		}
		obj1 := P{1, 2, 3, "hello"}
		buf := &bytes.Buffer{}
		encoder := gob.NewEncoder(buf)
		encoder.Encode(obj1)
		data := buf.Bytes()

		obj2 := Q{}
		decoder := gob.NewDecoder(bytes.NewBuffer(data))
		decoder.Decode(&obj2)
		fmt.Println(obj2)
	}
	// json serialization
	{
		jsontext := `
			{"people_name": "tom", "people_age": 21}
			`
		type People struct {
			Name string `json:"people_name"`
			Age  int32  `json:"people_age"`
		}
		// decoder & encoder
		decoder := json.NewDecoder(strings.NewReader(jsontext))
		people := People{}
		decoder.Decode(&people)
		buf := bytes.Buffer{}
		encoder := json.NewEncoder(&buf)
		encoder.Encode(people)

		//Marshal & Unmarshal
		person := People{}
		json.Unmarshal([]byte(jsontext), &person)
		jsonbytes, _ := json.Marshal(person)
		fmt.Println(string(jsonbytes))
	}
	// context
	// context是在发生API调用时传递的上下文
	// 本质上是channel的封装, 可以组成一个上下文树, 进行cancel操作
	{
		// ctx, cancel := context.WithCancel(context.Background())
		// go func() {
		// 	for range time.Tick(time.Second) {
		// 		select {
		// 		case <-ctx.Done():
		// 			fmt.Println("Goroutine Done")
		// 			return
		// 		default:
		// 			fmt.Println("Goroutine Tick")
		// 		}
		// 	}
		// }()

		// time.Sleep(3 * time.Second)
		// cancel()
		// time.Sleep(3 * time.Second)
	}
	// file
	{
		root := `/Users/liufuliang/project/golang_sample/src/github.com/lxb31/golang_sample/test/testdata`
		filename := path.Join(root, "file.txt")
		// ioutil
		ioutil.ReadFile(filename)
		// io.Reader
		file, err := os.Open(filename)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		// Read specific len bytes
		// buf := make([]byte, 512)
		// file.Read(buf)

		// Fscanf format type
		// var value int
		// fmt.Fscanf(file, "%d", &value)

		// bufio.NewScanner line by line
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			scanner.Text()
		}
	}
	//index.SuffixArray
	{
		text := `
			Suffix tree preprocesses the text, to make any pattern searching better. 
			You can preprocess the text and search any pattern in O(m) time, 
			where m denotes the length of the pattern you are searching for.
		`
		// 构建
		index := suffixarray.New([]byte(text))
		offsets := index.Lookup([]byte("search"), -1)
		handle(offsets)

		// 持久化
		buf := bytes.Buffer{}
		writer := bufio.NewWriter(&buf)
		index.Write(writer)
		writer.Flush()

		// 加载
		index1 := suffixarray.New(nil)
		index1.Read(&buf)
	}
	// regexp
	{
		// MatchString
		ok, _ := regexp.MatchString(`windows[0-9]+`, "windows2000")
		if ok {
		}
		//Compile
		re, _ := regexp.Compile(`windows[0-9]+`)
		re.MatchString("windows2000")
		re.MatchString("windows98")
		// FindAllString
		// [windows98 windows2000]
		re, _ = regexp.Compile(`windows[0-9]+`)
		re.FindAllString("windows98&windows2000&windowsxp", -1)
		// FindAllStringSubmatch, 额外返回捕获组
		// [[windows98 98] [windows2000 2000]]
		re, _ = regexp.Compile(`windows([0-9]+)`)
		re.FindAllStringSubmatch("windows98&windows2000&windowsxp", -1)
		// FindString
		// windows98
		re, _ = regexp.Compile(`windows[0-9]+`)
		re.FindString("windows98&windows2000&windowsxp")
		// ReplaceAllLiteralString
		// linux&linux&windowsxp
		re, _ = regexp.Compile(`windows([0-9]+)`)
		re.ReplaceAllLiteralString("windows98&windows2000&windowsxp", "linux")
		// ReplaceAllString, 支持捕获组"$"扩展
		// linux98&linux2000&windowsxp
		re, _ = regexp.Compile(`windows([0-9]+)`)
		re.ReplaceAllString("windows98&windows2000&windowsxp", "linux$1")
	}
	// time
	{
		// time.After 一次性定时器
		c := make(chan int)
		select {
		case m := <-c:
			handle(m)
		case <-time.After(10 * time.Millisecond):
			log.Println("timeout")
		}
		//time.Tick 周期定时器
		// for now := range time.Tick(1 * time.Second) {
		// 	fmt.Println(now)
		// }
		//time.Time
		t := time.Now()
		t.Year()
		t.Month()
		t.Day()
		t.Unix()
		// Duration
		t1 := time.Now()
		handle(nil)
		t2 := time.Now()
		t2.Sub(t1).Nanoseconds()
		// Format
		layout := "2006-01-02 15:04:05 -07:00"
		timestring := time.Now().Format(layout)
		handle(timestring)
		// Round
		time.Now().Round(time.Minute)
		time.Now().Truncate(time.Minute)
		// time.Location
		zone, _ := time.LoadLocation("Asia/Shanghai")
		zone = time.Local
		// time.Date
		date := time.Date(2019, 7, 5, 15, 35, 0, 0, zone)
		handle(date)
	}
	// reflect
	{
		obj := &A{Name: "bob", Number: 31}
		// TypeOf
		typ := reflect.TypeOf(obj).Elem()
		field, _ := typ.FieldByName("Name")
		field.Tag.Get("tag1")
		// ValueOf
		val := reflect.ValueOf(obj)
		methodValue := val.MethodByName("Echo")
		args := []reflect.Value{reflect.ValueOf("hello")}
		methodValue.Call(args)
	}

	{
		// bad case
		// 列表遍历时a是同一个内存, 所以取地址内容会变
		as := []A{
			{Name: "1", Number: 1},
			{Name: "2", Number: 2},
		}

		for _, a := range as {
			a.Number++
		}
		fmt.Println(as[0], as[1])
	}

	{
		type P struct {
			v int
		}
		type Q struct {
			p P
		}
		q := Q{p: P{1}}

		(&q.p).v = 2
		fmt.Println(&q.p)
	}

}
