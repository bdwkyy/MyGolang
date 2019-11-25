/*
 *  Golang发送post表单请求
 */
/*
 * http 服务端 ，处理 Post 请求
 */
 
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/postpage", func(w http.ResponseWriter, r *http.Request) {
		//接受post请求，然后打印表单中key和value字段的值
		if r.Method == "POST" {
			var (
				key   string = r.PostFormValue("key")
				value string = r.PostFormValue("value")
			)
			fmt.Printf("key is  : %s\n", key)
			fmt.Printf("value is: %s\n", value)
		}
	})

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
