package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

// Gitlab 结构体
type Gitlab struct {
	Id                  int    `json:"id"`
	Path_with_namespace string `json:"path_with_namespace"`
}

// 打印Body中的数据
func printBody(r *http.Response) {
	totalPages := r.Header.Get("X-Total-Pages")       // 取总页数
	page := fmt.Sprintf("%d", r.Header.Get("X-Page")) // 取当前页数
	n_totalPages, _ := strconv.Atoi(totalPages)
	n_page, _ := strconv.Atoi(page)

	for {
		if n_page < n_totalPages {
			fmt.Println(n_page)
			n_page += 1
			requestByParams(strconv.Itoa(n_page))
		} else if n_page == n_totalPages {
			break
		}
	}

	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var gitlab []Gitlab
	if err := json.Unmarshal(content, &gitlab); err != nil {
		fmt.Println(err)
		return
	}
	for k, v := range gitlab {
		fmt.Println(k, v.Id, v.Path_with_namespace)
	}

}

func requestByParams(page string) {
	request, err := http.NewRequest("GET", "http://code.gome.inc/gitlab/api/v4/projects", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	request.Header.Add("PRIVATE-TOKEN", "ve9ifEu9ZtAeqWik1L_y")
	params := make(url.Values)
	//params.Add("per_page", per_page)
	params.Add("page", page)
	params.Add("membership", "true")

	fmt.Println(params.Encode())           // 打印GET参数
	request.URL.RawQuery = params.Encode() // params.encode就是GET请求的参数进行URL解析
	r, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 关闭Body
	defer r.Body.Close()

	// 打印相关数据
	printBody(r)

}

func main() {

	requestByParams("1")

}
