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

var totalPages uint
var slicestr = make([]string, 0)

// 打印Body中的数据
func printBody() {
	r := requestByParams("1")
	//获取第一页内容

	r.Body.Close()
	// 取总页数
	totalPages := r.Header.Get("X-Total-Pages")
	n, err := strconv.Atoi(totalPages)
	if err != nil {
		fmt.Println(err)
		return
	}

	//获取以后页数
	for i := 2; i <= n; i++ {
		r = requestByParams(strconv.Itoa(i))
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
		for _, v := range gitlab {
			slicestr = append(slicestr, v.Path_with_namespace)
			//fmt.Println(k, v.Id, v.Path_with_namespace)
		}
		r.Body.Close()
	}
	fmt.Println(len(slicestr))

}

func requestByParams(page string) *http.Response {
	request, err := http.NewRequest("GET", "http://code.gome.inc/gitlab/api/v4/projects", nil)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	request.Header.Add("PRIVATE-TOKEN", "ve9ifEu9ZtAeqWik1L_y")
	params := make(url.Values)
	//params.Add("per_page", per_page)
	params.Add("page", page)
	params.Add("membership", "true")

	//fmt.Println(params.Encode())           // 打印GET参数
	request.URL.RawQuery = params.Encode() // params.encode就是GET请求的参数进行URL解析
	r, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return r
	// 关闭Body

	// 打印相关数据
	//printBody(r)

}

func main() {

	printBody()

}
