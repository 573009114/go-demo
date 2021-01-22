package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/573009114/go-demo/tree/main/internal/pkg"
)

//KmrMetrics 是KMR的结构化的接口返回信息, 自动生成结构体网站 https://transform.tools/json-to-go
type KmrMetrics struct {
	Apps struct {
		App []struct {
			ID                         string  `json:"id"`
			User                       string  `json:"user"`
			Name                       string  `json:"name"`
			Queue                      string  `json:"queue"`
			State                      string  `json:"state"`
			FinalStatus                string  `json:"finalStatus"`
			Progress                   float64 `json:"progress"`
			TrackingUI                 string  `json:"trackingUI"`
			TrackingURL                string  `json:"trackingUrl"`
			Diagnostics                string  `json:"diagnostics"`
			ClusterID                  int64   `json:"clusterId"`
			ApplicationType            string  `json:"applicationType"`
			ApplicationTags            string  `json:"applicationTags"`
			Priority                   int     `json:"priority"`
			StartedTime                int64   `json:"startedTime"`
			FinishedTime               int64   `json:"finishedTime"`
			ElapsedTime                int     `json:"elapsedTime"`
			AmContainerLogs            string  `json:"amContainerLogs"`
			AmHostHTTPAddress          string  `json:"amHostHttpAddress"`
			AllocatedMB                int     `json:"allocatedMB"`
			AllocatedVCores            int     `json:"allocatedVCores"`
			RunningContainers          int     `json:"runningContainers"`
			MemorySeconds              int     `json:"memorySeconds"`
			VcoreSeconds               int     `json:"vcoreSeconds"`
			QueueUsagePercentage       float64 `json:"queueUsagePercentage"`
			ClusterUsagePercentage     float64 `json:"clusterUsagePercentage"`
			PreemptedResourceMB        int     `json:"preemptedResourceMB"`
			PreemptedResourceVCores    int     `json:"preemptedResourceVCores"`
			NumNonAMContainerPreempted int     `json:"numNonAMContainerPreempted"`
			NumAMContainerPreempted    int     `json:"numAMContainerPreempted"`
			LogAggregationStatus       string  `json:"logAggregationStatus"`
			UnmanagedApplication       bool    `json:"unmanagedApplication"`
			AmNodeLabelExpression      string  `json:"amNodeLabelExpression"`
		} `json:"app"`
	} `json:"apps"`
}

//Gettestcpu 通过本地程序包，获取CPU信息
func Gettestcpu() {
	s, _ := pkg.GetCPUInfo()
	fmt.Println(s[0])
}

//GetURL 获取URL地址
func GetURL(urls string) string {
	res, _ := url.Parse(urls)
	fmt.Printf("%s", res.Query().Encode())
	return res.Query().Encode()
}

//GetPageStr 获取页面内容
func GetPageStr(urls string) (*KmrMetrics, error) {

	res, _ := http.Get(urls)
	response, err := ioutil.ReadAll(res.Body)

	defer res.Body.Close() //函数结束后执行defer
	// fmt.Println(reflect.TypeOf(response))

	data := KmrMetrics{} //指向kmrMetrics struct 格式化数据

	err = json.Unmarshal(response, &data) //转换数据类型，将当前response结果转成KmrMetrics 结构体

	//循环遍历结构体数据
	for d := range data.Apps.App {
		if err == nil {
			fmt.Println(data.Apps.App[d].Name)
			fmt.Println(data.Apps.App[d].State)
			// panic("something wrong") //阻断功能
		}
	}

	return &data, err
}

func main() {
	dateTime := time.Now().Unix()
	urls := fmt.Sprintf("%s%d", "http://192.168.1.1:8019/ws/v1/cluster/apps?startedTimeBegin=", dateTime) //地址拼接
	GetPageStr(urls)
}
