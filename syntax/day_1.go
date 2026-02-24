package syntax

import "fmt"

var i int
var f float64
var s string = "hello world"
var b bool
var p *int = nil
var sl = []int{1, 2, 3}
var m = map[int]string{1: "one", 2: "two"}
var ch chan int = make(chan int)
var fn func()
var iface interface{}

func printCommand[T any](name string, v T) {
	fmt.Printf("%s: type=%T value=%v\n", name, v, v)
}

func PrintVariables() {
	printCommand("i", i)
	printCommand("f", f)
	printCommand("s", s)
	printCommand("b", b)
	printCommand("p", p)
	printCommand("sl", sl)
	printCommand("m", m)
	printCommand("ch", ch)
	printCommand("fn", fn)
	printCommand("iface", iface)
}

func Converting() {
	var a int = 100
	printCommand("a", a)
	printCommand("converting to int64 a", int64(a))
	fmt.Println()

	var b int64 = 100
	printCommand("b", b)
	printCommand("converting to int a", int(b))
	fmt.Println()

	var d byte = 64
	printCommand("d", d)
	printCommand("converting to int d", int(d))
	fmt.Println()

	var r rune = '&'
	printCommand("r", r)
	printCommand("converting to int r", int(r))

	d = byte(255)
	fmt.Println(b)

	d++
	fmt.Println(b)
}
