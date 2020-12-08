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
var slicestr = make([]int, 0)

// 打印Body中的数据
func printBody(r *http.Response) {

	if r.Body != nil {
		defer r.Body.Close()
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
	for _, v := range gitlab {
		slicestr = append(slicestr, v.Id)
		//fmt.Println(k, v.Id, v.Path_with_namespace)
	}
	fmt.Println(slicestr)
	/*
	 */
}

// requestByParams 获取所有项目信息
func requestByParams(page string) *http.Response {
	request, err := http.NewRequest("GET", "http://code.com.com/gitlab/api/v4/projects", nil)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	request.Header.Add("PRIVATE-TOKEN", "ve9ifEu9ZtAeqWik1L")
	params := make(url.Values)
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

}

// 获取项下目的分支

func main() {
	r := requestByParams("1")
	printBody(r)
	// 取总页数
	totalPages := r.Header.Get("X-Total-Pages")
	nPage, err := strconv.Atoi(totalPages)
	if err != nil {
		fmt.Println(err)
		return
	}
	//获取以后页数
	for i := 2; i <= nPage; i++ {
		r = requestByParams(strconv.Itoa(i))
		printBody(r)
	}

}
