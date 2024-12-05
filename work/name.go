
// 問題１：配列nameに「スコティッシュ」「マンチカン」「ミケ」を代入
// 名前に「さん」を追加して出力する。出力結果は「スコティッシュさん」「マンチカンさん」「ミケさん」
package main
import "fmt"

func main() {
	// 配列nameに「スコティッシュ」、「マンチカン」、「ミケ」を代入
	var name [3]string = [3]string{"スコティッシュ", "マンチカン", "ミケ"}
	// nameの要素数分繰り返す
	for i := 0; i < len(name); i++ {
		fmt.Println(name[i] + "さん")
	}
}