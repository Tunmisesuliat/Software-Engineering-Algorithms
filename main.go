package main

import "fmt"

//CHALLENGE 2
func nested(s rune) (output int){
	var string_arr []string
	var temp []string
	for _, v:= range s {
		arr:= [] string{'[', '{', '('}
		for _, x:= range arr {
			if v == x{
				temp=append(temp, v)
			}
		}
	}
}

// CHALLENGE 3
func dominator(b []int) (interface {}) {
	max :=1
	val:=0
	a := make(map[int]int)
	var temp []int
	for _, v:= range b{
		a[v] ++
		if a[v] > max{
			max = a[v]
			val = v
		}
	}
	if max > len(b)/2{
	for i, v :=range b{
		if v==val {
			temp=append(temp, i)
		}
	}
		return temp
	}		
	return -1
}
func main() {
	a := [] int{3,4,3,2,3, -1, 3, 3}
	fmt.Println(dominator(a))
}
