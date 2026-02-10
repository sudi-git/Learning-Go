package main
import "fmt"

//package sprint
func BetweenLimits(from, to rune) string {
//just to make a copy of the initial
start := from
end := to
    if from > to{
        start = to
        end =from
    }
str := ""
for r:= start+1 ; r< end ; r++ {
    str = str + string(r)
}
return str 
}


func main(){
	 fmt.Println(BetweenLimits('j', 'f'))
}
