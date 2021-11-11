package main
import "fmt"
type All interface {
	dove()
	repeater()(word string)
	lemon_monster()
	tasty_monster()
}
type person struct {
	name  string
	age int
	gender string
}
func (temp person) dove(){
	fmt.Println("gugugugugugu")
}
func (temp person) repeater(word string)  {
	fmt.Println(word)
}
func (temp person)lemon_monster(){
	fmt.Println("It's so sour!")
}
func (temp person)tasty_monster(){
	fmt.Println("It's so delicious!")
}
func main() {
	p:=person{
		name:   "小红",
		age:    18,
		gender: "男",
	}
	p.dove()//实现了鸽子接口————>dove()
	p.repeater("gogogo!")//实现了复读机接口————>repeater()
	p.lemon_monster()//实现了柠檬精接口————>lemon_monster()
	p.tasty_monster()//实现了真相怪接口————>tasty_monster()
}
