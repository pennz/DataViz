package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	_ "github.com/heroku/x/hmetrics/onload"
	"github.com/stretchr/testify/assert"
)

func TestGetThings(t *testing.T) {
	tests := []struct {
		name    string
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Init", 301, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			router := setupRouter()

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/Viz", nil)

			router.ServeHTTP(w, req)
			assert.Equal(t, tt.want, w.Code)
		})
	}
}

func TestCompile(t *testing.T) {
	tests := []struct {
		name    string
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
		{"SimpleGraph", 200, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			router := setupRouter()

			w := httptest.NewRecorder()
			bodyString := "body=%2F%2F%20!!!%20use%20Ctrl%2BShift%2BC%20to%20compile%20code%20to%20graph%0A%2F%2F%20credits%20goes%20to%20https%3A%2F%2Fplay.golang.org%2F%20for%20compiling%20and%20runing%20our%20code.%0A%2F%2F%0A%2F%2F%20Package%20binaryheap%20implements%20a%20binary%20heap%20backed%20by%20array%20list.%0A%2F%2F%0A%2F%2F%20Comparator%20defines%20this%20heap%20as%20either%20min%20or%20max%20heap.%0A%2F%2F%0A%2F%2F%20Structure%20is%20not%20thread%20safe.%0A%2F%2F%0A%2F%2F%20References%3A%20http%3A%2F%2Fen.wikipedia.org%2Fwiki%2FBinary_heap%0Apackage%20main%0A%0Aimport%20(%0A%09%22github.com%2Fpennz%2FDataViz%2Flists%2Farraylist%22%0A%09%22github.com%2Fpennz%2FDataViz%2Fviz%22%0A)%0A%0Afunc%20main()%20%7B%0A%09l%20%3A%3D%20arraylist.New()%0A%09l.Add(-1%2C0%2C1%2C-2)%0A%09avw%20%3A%3D%20viz.NewAlgVisualWrapper()%0A%09%0A%09avw.Wrap(l)%20%2F%2F%20already%20a%20pointer%20now...%0A%09avw.Call(%22Add%22%2C%203)%0A%09avw.Call(%22Add%22%2C%204)%0A%09avw.Call(%22Add%22%2C%205)%0A%09avw.Call(%22Swap%22%2C%200%2C%201)%0A%09%2F%2Flog.Println(avw.Visualize())%0A%09avw.Visualize()%0A%7D"
			req, _ := http.NewRequest("POST", "/compile", strings.NewReader(bodyString))
			req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:75.0) Gecko/20100101 Firefox/75.0")
			req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
			req.Header.Set("Accept-Language", "en-US,en;q=0.7,zh;q=0.3")
			req.Header.Set("Referer", "https://go-algorithm.herokuapp.com/Viz/")
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
			req.Header.Set("X-Requested-With", "XMLHttpRequest")
			req.Header.Set("Pragma", "no-cache")
			req.Header.Set("Cache-Control", "no-cache")
			req.Header.Set("Origin", "https://go-algorithm.herokuapp.com")
			req.Header.Set("DNT", "1")
			req.Header.Set("Connection", "keep-alive")

			router.ServeHTTP(w, req)
			assert.Equal(t, tt.want, w.Code)
			res := w.Result()
			resBytes, err := ioutil.ReadAll(res.Body)
			assert.Equal(t, err, nil)
			//log.Println(string(resBytes))
			assert.Equal(t, len(resBytes) > 5, true)
		})
	}
}

func TestPingRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
