package main

import "fmt"

func main(){
	s:= "   fly me   to   the moon  "

	fmt.Printf(`Length of the last word in "%s" is %v`,s,lengthOfLastWord(s))
	fmt.Println()
}

func lengthOfLastWord(s string) int {
    var count int
    var lastIndex = len(s)-1

    for i:=lastIndex;i>=0;i--{
        if string(s[i]) != " "{
            count++
        }else if string(s[i]) == " " && count!=0{
            return count
        }
    }

    return count
}

