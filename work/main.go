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


	var n int =100
	var p * int = &n	// nのアドレスをポインタに返す
	
	fmt.Println(*p)		// nの値を出力
	fmt.Println(p)		// nのアドレスを出力
}
