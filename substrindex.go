package main
//import "fmt"
//func main() {
//fmt.Println(SubstrIndex("How are you?", "o"))
//fmt.Println(SubstrIndex("How are you doing?", "ou"))
//fmt.Println(SubstrIndex("You can do it!", " od"))
//}

func SubstrIndex(s string, toFind string) int {

    if toFind == "" || toFind == s {
        return 0
    }
    if s == "" {
        return -1
    }

    // To avoid the outbound we subtract len of toFind
    for i:=0; i < len(s) - len(toFind); i++ {
        if s[i:i+len(toFind)] == toFind { // this expression makes the i out of range if we keep the loop until the end
            return i
        }
    }
    return -1
}
