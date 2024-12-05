// Webサーバを構築する
package main
// 必要なパッケージをインポート
import (
	"fmt"
	"net/http"
	"log"
	"html/template"
)

//handler関数
func helloHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templ.html")
	if err != nil {
		log.Println(err)
	}
	t.Execute(w, "Hello World!!")
}

func main(){
  ///helloの時helloHandlerを呼び出す
	http.HandleFunc("/hello", helloHandler)
	fmt.Println("Server Starting port:8000")
	http.ListenAndServe(":8000", nil)	
}

//templ.html
<html>
<head>
    <meta charset="utf8">
    <title>Go Web Programing</title>
</head>
<body>
    {{ . }}
</body>
</html>http.ListenAndServe(":8000", nil)	
}