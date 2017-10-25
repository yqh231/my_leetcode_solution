/*
Judge Route Circle - LeetCode
Initially, there is a Robot at position (0, 0). Given a sequence of its moves, judge if this robot makes a circle, which means it moves back to the original place. 



The move sequence is represented by a string. And each move is represent by a character. The valid robot moves are R (Right), L (Left), U (Up) and D (down). The output should be true or false representing whether the robot makes a circle.


Example 1:

Input: "UD"
Output: true



Example 2:

Input: "LL"
Output: false

*/
package main

import(
	"fmt"
)
func judgeCircle(moves string) bool{
	i, j := 0, 0
	for _, step := range moves{
		if step == 'U'{
			i += 1
		}else if step == 'D'{
			i -= 1
		}else if step == 'L'{
			j -= 1
		}else if step == 'R'{
			j += 1
		}

	}
	return i == 0 && j == 0
}

func main(){
	var a string = "UUUUUDDDD"
	for i := range a{
		fmt.Println(i)	
	}
}