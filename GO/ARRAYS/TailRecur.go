package main

import (
	"fmt"
)
func print_num(n int){
	if n>0{
		fmt.Println(n)
		print_num(n-1)
	}
}
//print_num( n int)
func main(){
	print_num(5)

}