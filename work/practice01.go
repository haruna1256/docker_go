// 問題１：配列nameに「スコティッシュ」「マンチカン」「ミケ」を代入
// 名前に「さん」を追加して出力する。出力結果は「スコティッシュさん」「マンチカンさん」「ミケさん」
package main
import "fmt"

func main() {
    // 配列nameに「スコティッシュ」、「マンチカン」、「ミケ」を代入
    name := []string{"スコティッシュ", "マンチカン", "ミケ"}
    // nameの要素数分繰り返す
    for i := 0; i < len(name); i++ {
        fmt.Printf("%sさん\n", name[i])
    }
}

