package main

func GetLastRune(s string) rune {
 len:= len(s)
 last := rune(s[len-1])
 return last
 
}