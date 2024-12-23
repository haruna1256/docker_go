// 問題3　問題2で作成したプログラムを関数を使用して書き直す

package main
import "fmt"

func main() {
	// 結果の配列を定義する関数
	result := result_data()
	// データ入力、名前と点数の関数
	name, score := input_data()
	// 点数の判定
    index := judge_point(score)
	// 結果を表示する関数
	result_view(name,index,result)

}

// 結果の配列を定義する関数
func result_data() []string {
	result := []string{"NG!", "Good!", "Very Good!", "Excellent!"}
	return result
}

// データ入力する関数
func input_data() (string, int) {
	// 名前、点数の変数を定義
	var name string
	var score int

	// 名前を入力する
	fmt.Print("名前:")
	fmt.Scanf("%s", &name)	

	// 点数を入力する
	fmt.Print("点数:")
	fmt.Scanf("%d", &score)

	return name, score
}

// 点数を判定する関数
func judge_point(score int) int {
	// 結果を表示する配列の場所を指定する変数
	var index int
	// 入力された点数が数字かどうかを確認する
	if score < 0 || score > 100 {
		fmt.Println("点数は0〜100までの数字を入力してください")
		return 0
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

	return index
}


// 結果を表示する関数
func result_view(name string, index int, result []string) {
	// 結果を表示する
	fmt.Println(name,"さん")
	fmt.Println(result[index])

}

