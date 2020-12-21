package myserver

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHTTPHandler(t *testing.T) {
	//Response를 record하기 위한 ResponseRecoder를 만듬
	//NewRecorder는 초기화된 ResponseRecorder를 반환
	res := httptest.NewRecorder()
	//handler에 전달할 Reqeuest를 만듬, 지금은 쿼리 파라미터가 없음
	//NewRequest는 백그라운드 컨텍스트를 사용하여 NewRequestWithContext를 래핑

	// target을 /main , /data , /monitor , /controll 로 변경하며 test
	//$ go test -v
	req := httptest.NewRequest("GET", "/controll", nil)
	mux := HTTPHandler()
	mux.ServeHTTP(res, req)

	//HTTP 요청에 대한 응답확인. StatusOK = 200
	if res.Code != http.StatusOK {
		t.Fatal("Failed:", res.Code)
	} else {
		fmt.Printf("HTTP StatusOK")
	}

	fmt.Printf(" :")

	/* Response 확인
	/monitor -> Battery Status!
	/main -> Main!
	/data -> DATABASE
	/controll -> Controll battery
	*/
	expected := "Controll battery"
	data, _ := ioutil.ReadAll(res.Body)
	if string(data) != expected {
		t.Errorf("Failed: Got <%v>, Expected <%v> \n", string(data), expected)
	} else {
		fmt.Printf("Response OK")
	}
}
