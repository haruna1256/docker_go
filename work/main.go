package main

// 出力できるようにする
import "fmt"



func Double(i * int) {	// 引数はint型
	*i = *i * 2			// アドレスを参照して演算する
}


// はじめに出力
func main(){
	fmt.Println("Hello World!!!!")
	fmt.Println("test Go!")
	fmt.Println("Hello Go!!!!")
}
