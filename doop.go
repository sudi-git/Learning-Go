package main

import "fmt"
//package sprint
func Doop(a int, op string, b int) int {
    switch op {
        case "+" :
            return a+b
        case "-" :
            return a-b
        case "/" :
            if b==0{
                return 0
            }
            return a/b
        case "*" :
            return a*b
        case "%" :
            if b==0{
            return 0
            }
            return a%b
        default :
            return 0
    }
    return 0
}


func main(){
	fmt.Println(Doop(5, "+", 3))
}
