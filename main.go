package main

import "fmt"

//CHALLENGE 1
func triangular_triplet(b []int) (result int){
	length := len(b)
	if length < 3 {
		result = 0;
	}
	result = 0
	for i:= 0; i < length - 2; i++ {
		if b[i] + b[i + 1] > b[i + 2] {
			result = 1;
            break;
		}
	}
	return 
}

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
	i:=[] int{10, 2, 5, 1, 8, 20}
	a := [] int{3,4,3,2,3, -1, 3, 3}
	fmt.Println(triangular_triplet(a))
	fmt.Println(dominator(i))
	fmt.Println("")
}
