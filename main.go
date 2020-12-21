package main

import (
	"myserver"
	"net/http"
)

/*
1) http://localhost:8080/(path)
	path에 "main", "data", "monitor", "controll"를 입력
2) 주소에 따라 텍스트 출력
3) myserver폴더에서 servemux_test.go 파일로 요청/응답을 테스트
	$ go test -v
*/

func main() {
	/* 방법 1)
	myserver.NewHTTPHandler()
	//코드가 실행되면 서버가 실행되며 8080포트에서 Request를 기다리는 상태가 된다
	http.ListenAndServe(":8080", nil)
	*/
	//방법 2)
	//코드가 실행되면 서버가 실행되며 8080포트에서 Request를 기다리는 상태가 된다\
	//8080 포트에서 웹 서버를 열고 클라이언트 request를 받아서 새go루틴에 작업을 할당

	//8080 포트에서 서버 실행
	//http.ListenAndServe(":8080", myserver.NewHTTPHandler())

	//mux를 사용
	//8080 포트에서 서버 실행
	http.ListenAndServe(":8080", myserver.HTTPHandler())

}
