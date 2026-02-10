package main
//import "fmt"
//func main(){
//fmt.Println(BalanceOut([]bool{true, false, false, false})) }

//package sprint
func BalanceOut(arr []bool) []bool {

var t int //counts true
var f  int //counts false
var bal[]bool //slice that holds extra values needed to balance ;currently its nil

//to count ;  range go through each array at a time , ignores index and stores at v
 for _, v:=  range arr{
 	if v {
  		t=t+1
  	} else {
   		f=f+1
   	}
 }
// to find the difference
if t > f {
	diff:= t-f
	bal = make([]bool, diff) //creates slice of false as bool does by default
} else if f > t {
	diff:= f-t
	bal = make([]bool, diff)
		for i:=0 ; i<diff ; i++{ //to convert false to true we loop and set each one
		bal[i] = true
		}
	}
	//adding elements to slice
	result := append(arr, bal...)
	return result
}
