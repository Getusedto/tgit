package myserver

import (
	"fmt"
	"net/http"
)

func mainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Main!")
	//w(writer)에 "문자열"을 출력
	//http.Request는 Request를 검토할 수 있게 한다
	//http.ResponseWriter는 HTTP Response에 무엇가 쓸 수 있게 함

	//html 파일을 Parsing
	//html, err := template.PraseFiles("view.html")
	//template.PraseFiles("view.html")
}
func dataHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "DATABASE")
}
func monitorHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Battery Status")
}
func controlltHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Controll battery")
}

/* <인스턴스 만들기>
type controlltHandler struct{}

//인스턴스에 해당하는 인터페이스를 구현. Serve인터페이스를 구현
func (f *controlltHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Controll battery")
}
*/

func HTTPHandler() http.Handler {
	//HTTP 요청에 대해 라우터 기능을 하는 mux를 생성
	mux := http.NewServeMux()

	/* <핸들러 만들기>
	1)HandleFunc를 이용해서 핸들러를 함수 형태로 등록 하는 방법
	2)Handle를 이용해서 인스턴스 형태로 핸들러를 등록하는 방법
	*/

	//Handle() and HandleFunc() is 클라이언트가 요청한 request path에 해당하는
	//request 핸들러를 지정해주는 라우터 역할을 한다
	//Request가 들어왔을때 "경로"에 해당하는 핸들러를 실행

	mux.HandleFunc("/main", mainHandler)
	mux.HandleFunc("/data", dataHandler)
	mux.HandleFunc("/monitor", monitorHandler)
	mux.HandleFunc("/controll", controlltHandler)

	//Handle 사용
	//mux.Handle("/controll", &controlltHandler{})
	return mux
}
