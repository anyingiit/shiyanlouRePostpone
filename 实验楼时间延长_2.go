package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	fmt.Println("Program is running!")
	dispatch()
}

func dispatch() {
	var next,surplusTime int
	var maxRunTime int
	fmt.Print("请输入最大运行次数:")
	fmt.Scanln(&maxRunTime)
	maxRunTime=int(maxRunTime)
	var url_select string = "https://www.shiyanlou.com/api/v2/labtask/"
	var url_repost string = "https://www.shiyanlou.com/api/v2/labtask/extend/"
	for{
		fmt.Println("dispatch is running!")
		surplusTime = examineAndGetTime(url_select)
		next = (surplusTime/60)/2 //下一次检查时间(分钟)
		fmt.Println("剩余时间:",surplusTime/60,"分","下一次检查:",next,"分")
		if maxRunTime>0 {
			if next>8 {
				time.Sleep(time.Minute*time.Duration(next))
			}else {
				RepostponeShiyanlou(url_repost)
				maxRunTime--
				fmt.Println("执行了时间延长,剩余运行次数:",maxRunTime,"次")
				fmt.Println("一分钟后进行下一次检查...")
				time.Sleep(time.Minute*1)
			}
		}else {
			return
		}

	}

}

func examineAndGetTime(url string) (surplusTime int) {
	datas :=map[string]interface{}{}

	req, _ := http.NewRequest("GET", url, nil)
***CREDENTIAL_REMOVED_BY_REPOCURATOR***
	req.Header.Add("User-Agent", "PostmanRuntime/7.15.2")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Postman-Token", "40494ecf-2643-4a25-8951-02518a6353c7,6bc64ce0-7f12-42b9-99b1-cbecd076ab69")
	req.Header.Add("Host", "www.shiyanlou.com")
	req.Header.Add("Accept-Encoding", "gzip, deflate")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("cache-control", "no-cache")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println("---------------------examineTime Start---------------------")
	fmt.Println(res)
	fmt.Println(string(body))
	if err := json.Unmarshal(body, &datas); err != nil {
		println("body to map error!",err)
	}

	result:= int(datas["ttl_seconds"].(float64))
	fmt.Println("The ttl_seconds is:", result)
	fmt.Println("---------------------examineTime End---------------------")
	return result
}

func RepostponeShiyanlou(url string) {
	fmt.Println("---------------------","RepostponeShiyanlou Start","---------------------")
	req, _ := http.NewRequest("POST", url, nil)

***CREDENTIAL_REMOVED_BY_REPOCURATOR***
	req.Header.Add("User-Agent", "PostmanRuntime/7.15.2")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Postman-Token", "ea016f2d-1bc8-483f-943c-69b8385d44f3,734295b3-fdd4-4d4b-adee-3d28f7805221")
	req.Header.Add("Host", "www.shiyanlou.com")
	req.Header.Add("Accept-Encoding", "gzip, deflate")
	req.Header.Add("Content-Length", "")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("cache-control", "no-cache")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
	fmt.Println("---------------------","RepostponeShiyanlou End","---------------------")
}