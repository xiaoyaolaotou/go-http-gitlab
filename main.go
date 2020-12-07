package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"

	"net/http"
)

// Gitlab 结构体
type Gitlab struct {
	Id                  int    `json:"id"`
	Path_with_namespace string `json:"path_with_namespace"`
}

// 打印Body中的数据
func printBody(r *http.Response) {

	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var gitlab []Gitlab
	if err := json.Unmarshal(content, &gitlab); err != nil {
		fmt.Println(err, "xxxx")
		return
	}
	for k, v := range gitlab {
		fmt.Println(k, v.Id, v.Path_with_namespace)
	}

}

func requestByParams() {
	request, err := http.NewRequest("GET", "http://code.com.com/gitlab/api/v4/projects", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	request.Header.Add("private-token", "ve9ifEu9ZtAeqWik1L_y")
	params := make(url.Values)
	params.Add("membership", "true")
	//fmt.Println(params.Encode())
	request.URL.RawQuery = params.Encode() // params.encode就是GET请求的参数
	r, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 关闭Body
	defer r.Body.Close()

	printBody(r)

}

func main() {

	requestByParams()

}
