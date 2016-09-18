package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type BianminResult struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

func SayHello(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte(`{"errcode":3107,"errmsg":"来源ip异常"}`))
}

func SayHello2(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte(`test`))
}

func main() {
	response, _ := http.Get("http://bianmin.loc.php.yggx.com/api/order_query?uid=211686&pwd=123456&ptOrderNo=123&sign=35b2793f71a1ba183dc5982b583e768f")
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(body))

	str := `{"errcode":3107,"errmsg":"来源ip异常"}`
	res := &BianminResult{}
	fmt.Println(res)
	fmt.Println([]byte(str))
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)

	//	client := &http.Client{}
	//	reqest, _ := http.NewRequest("GET", "http://www.baidu.com", nil)

	//	reqest.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	//	reqest.Header.Set("Accept-Charset", "GBK,utf-8;q=0.7,*;q=0.3")
	//	reqest.Header.Set("Accept-Encoding", "gzip,deflate,sdch")
	//	reqest.Header.Set("Accept-Language", "zh-CN,zh;q=0.8")
	//	reqest.Header.Set("Cache-Control", "max-age=0")
	//	reqest.Header.Set("Connection", "keep-alive")

	//	response2, _ := client.Do(reqest)
	//	if response2.StatusCode == 200 {
	//		body, _ := ioutil.ReadAll(response2.Body)
	//		bodystr := string(body)
	//		fmt.Println(bodystr)
	//	}

	//	http.HandleFunc("/hello", SayHello)
	//	http.HandleFunc("/hello2", SayHello2)
	//	http.ListenAndServe(":8001", nil)

	result := &BianminResult{100, "hi"}
	//result.Errcode = 100
	//result.Errmsg = "登录成功"
	bytes, _ := json.Marshal(result)
	fmt.Println(string(bytes))

}
