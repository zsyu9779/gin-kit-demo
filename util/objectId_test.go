package util

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"go.doglobal.net/duapps/go-basic/obs"
	"go.doglobal.net/duapps/go-basic/util"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

func BenchmarkNewObjectID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewObjectID()
	}
}

func TestSplit(t *testing.T) {
	//pwd,_ := os.Getwd()
	src := "http://xiyou.sc.diyixin.com/pro-img-xiyou/20191022/15717552322363.jpg"
	arr := strings.Split(src, "/")
	logs.Info(arr)
	logs.Info(arr[len(arr)-1])
}

type fields struct {
	httpClient *util.HttpClient
	obsClient  *obs.ObsClient
	production string
	endPoint   string
}
type args struct {
	bucket string
	file   string
}

func TestDownloadImageFromLink(t *testing.T) {
	//url := "http://xiyou.sc.diyixin.com/pro-img-xiyou/20191222/15769813086072.jpg"
	//path := DownloadImageFromLink(url)
	//logs.Info(path)
	//fields := fields{production: "sandbox", endPoint: "obs.cn-south-1.myhuaweicloud.com"}
	//args := args{bucket: "material-feed", file: path}
	//h := service.NewHuaweiObs(fields.production, fields.endPoint)
	//got := h.PutFile(args.bucket, args.file)
	//logs.Info(got)
	//final:=strings.ReplaceAll(got,"https://obs.cn-south-1.myhuaweicloud.com/material-feed/","http://mv-res.xdplt.com/")
	//logs.Info(final)
}

func TestDownloadImageFromLink2(t *testing.T) {
	values, err := url.ParseRequestURI("https://www.baidu.com/s?wd=%E6%90%9C%E7%B4%A2&rsv_spt=1&issp=1&f=8&rsv_bp=0&rsv_idx=2&ie=utf-8&tn=baiduhome_pg&rsv_enter=1&rsv_sug3=7&rsv_sug1=6")
	fmt.Println(values)
	if err != nil {
		fmt.Println(err)
	}
	urlParam := values.RawQuery

	fmt.Println(urlParam)
	fmt.Println(values.Query().Get("rsv_spt"))
}



func TestDownloadImageFromLink3(t *testing.T) {

	url := "http://xy.uheixia.com/api/video/getShareVideos"

	payload := strings.NewReader(fmt.Sprintf("------WebKitFormBoundary7MA4YWxkTrZu0gW\r\nContent-Disposition: form-data; " +
		"name=\"type\"\r\n\r\n%d\r\n------WebKitFormBoundary7MA4YWxkTrZu0gW\r\nContent-Disposition: form-data; " +
		"name=\"userId\"\r\n\r\n%d\r\n------WebKitFormBoundary7MA4YWxkTrZu0gW\r\nContent-Disposition: form-data; " +
		"name=\"videoId\"\r\n\r\n%d\r\n------WebKitFormBoundary7MA4YWxkTrZu0gW\r\nContent-Disposition: form-data; " +
		"name=\"videoType\"\r\n\r\n%d\r\n------WebKitFormBoundary7MA4YWxkTrZu0gW--",10,8,8164,20))

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "multipart/form-data; boundary=----WebKitFormBoundary7MA4YWxkTrZu0gW")
	req.Header.Add("cache-control", "no-cache")
	//req.Header.Add("Postman-Token", "0df31b81-36bc-4c18-9b2f-3997e0b318f8")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}