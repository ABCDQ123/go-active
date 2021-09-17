package learn

import (
	"fmt"
	"runtime"
	"strconv"
	"time"
)

var key_go, value_go = "golang", "_with_no_this_point"

func New_const() {
	var key = "var key"
	value := key
	fmt.Println("slice: string :" + value + "/" + key_go + "/" + value_go)

	var key_int int
	key_int = 10
	fmt.Println("slice: int :" + strconv.Itoa(key_int))

	const key_const = 100
	fmt.Println("slice: const :" + strconv.Itoa(key_const))

	var key_bool = false
	key_bool = true
	fmt.Println("slice: bool :" + strconv.FormatBool(key_bool))
}

func New_enum() {
	const (
		x = iota
		y = iota
		z = iota
	)
	fmt.Println("slice: enum :" + strconv.Itoa(x) + " " + strconv.Itoa(y) + " " + strconv.Itoa(z))
}

func New_array() {
	var array_string [2]string
	array_string[0] = "array_string1"
	array_string[1] = "array_string1"
	var array_int = [2]int{1, 2}
	fmt.Println("slice: array :"+array_string[0], array_int[0])
	for index := 0; index < 100; index++ {

	}
	var array_int2 [100][100]int
	array_int2[99][1] = 1
	fmt.Println("slice: array2 ???:" + strconv.Itoa(array_int2[99][1]))
}

func New_slice_map() {
	var slice_string []string
	slice_string = append(slice_string, "slice1", "slice2")
	slice_string[0] = "slice0"
	fmt.Println("slice: slice append:"+slice_string[0], slice_string[1], len(slice_string), cap(slice_string))

	var map_int = make(map[int]int)
	map_int[11] = 11
	fmt.Println("slice: map ???:" + strconv.Itoa(map_int[11]))
}

type basic_student struct {
	name string
}

type student struct {
	basic_student
	school string
}

func New_struct() {
	student_a := student{basic_student{"name"}, "school_a"}
	fmt.Println("slice: struct :" + student_a.name)
}

type student_impl interface {
	name_m() string
	school_m() string
}

func (std *student) name_m() string {
	return std.basic_student.name
}

func (std *student) school_m() string {
	return std.school
}

func New_interface() {
	student_a := student{basic_student{""}, "school_b"}
	student_a.name = "Cc"
	fmt.Println("slice: interface :" + student_a.name_m() + " " + student_a.school_m())
}

func new_go(is int) {
	for i := 0; i < 10; i++ {
		runtime.Gosched()
		fmt.Println("slice: new_go" + strconv.Itoa(is) + ": " + strconv.Itoa(i))
		time.Sleep(1) //主协程 太快，执行完退出后，副协程无法执行
	}
}

func New_goroutine() {
	go new_go(1)
	new_go(2)
}

func new_channel(is int, c chan int) {
	for i := 0; i < 10; i++ {
		fmt.Println("slice: new_chan" + strconv.Itoa(is) + ": " + strconv.Itoa(i))
		time.Sleep(time.Second * 1)
	}
	c <- 10 + is
}

func New_channels() {
	c := make(chan int)
	go new_channel(1, c)
	go new_channel(2, c)
	x, y := <-c, <-c
	fmt.Println("slice: channel", strconv.Itoa(x), strconv.Itoa(y))
}
