package main

import "fmt"

func Receiver(v interface{}) {
	switch v.(type){
	case string: fmt.Println("这个是string")
	case int :fmt.Println("这个是int")
	case bool :fmt.Println("这个是bool")
	}

}
func main() {
	a:="I love you!"
	b:=520
	c:=true
	Receiver(a)
	Receiver(b)
	Receiver(c)
}
