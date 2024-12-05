// 問題2　名前と点数を入力すると、名前と結果を表示
// 60未満・・・NG、60～69・・・Good、70～85・・・Very Good、85以上・・・Excellent

package main
import "fmt"

func main() {
	// 点数によって表示する結果を配列に代入する
	result := []string{"NG", "Good", "Very Good", "Excellent"}
	// 結果を表示する配列の場所を指定する変数
	var index int

	// 名前を入力する
	fmt.Print("名前:")
	var name string
	fmt.Scanf("%s", &name)	// 入力された文字をうnameに代入

	// 点数を入力する
	fmt.Print("点数:")
	var score int
	fmt.Scanf("%d", &score)	// 入力された数字をscoreに代入

	// 入力された点数が数字かどうかを確認する
	if score < 0 || score > 100 {
		fmt.Println("点数は0〜100までの数字を入力してください")
		return
	}

	// 点数を判定する
	if score < 60 {
		index = 0
	} else if score < 70 {
		index = 1
	} else if score < 85 {
		index = 2
	} else {
		index = 3
	}

	// 結果を表示する
	fmt.println(name,"さん")
	fmt.println(result[index])

}		