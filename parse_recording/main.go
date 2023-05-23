package main

import (
	"encoding/json"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"strings"
	"time"
)

func main() {
	// 地域ID，固定值。
	const REGION_ID string = "cn-shanghai"
	const ENDPOINT_NAME string = "cn-shanghai"
	const PRODUCT string = "nls-filetrans"
	const DOMAIN string = "filetrans.cn-shanghai.aliyuncs.com"
	const API_VERSION string = "2018-08-17"
	const POST_REQUEST_ACTION string = "SubmitTask"
	const GET_REQUEST_ACTION string = "GetTaskResult"
	// 请求参数
	const KEY_APP_KEY string = "appkey" //此处appkey无需替换。
	const KEY_FILE_LINK string = "file_link"
	const KEY_VERSION string = "version"
	const KEY_ENABLE_WORDS string = "enable_words"
	const KEY_ENABLE_SAMPLE_RATE_ADAPTIVE string = "enable_sample_rate_adaptive"
	// 响应参数
	const KEY_TASK string = "Task"
	const KEY_TASK_ID string = "TaskId"
	const KEY_STATUS_TEXT string = "StatusText"
	const KEY_RESULT string = "Result"
	const KEY_SENTENCES string = "Sentences"
	const KEY_TEXT string = "Text"
	// 状态值
	const STATUS_SUCCESS string = "SUCCESS"
	const STATUS_RUNNING string = "RUNNING"
	const STATUS_QUEUEING string = "QUEUEING"
	var accessKeyId string = "" //获取AccessKey ID和AccessKey Secret请前往控制台：https://ram.console.aliyun.com/manage/ak
	var accessKeySecret string = ""
	var appKey string = ""
	var fileLink string = "https://gw.alipayobjects.com/os/bmw-prod/0574ee2e-f494-45a5-820f-63aee583045a.wav"
	client, err := sdk.NewClientWithAccessKey(REGION_ID, accessKeyId, accessKeySecret)
	if err != nil {
		panic(err)
	}
	postRequest := requests.NewCommonRequest()
	postRequest.Domain = DOMAIN
	postRequest.Version = API_VERSION
	postRequest.Product = PRODUCT
	postRequest.ApiName = POST_REQUEST_ACTION
	postRequest.Method = "POST"
	mapTask := make(map[string]string)
	mapTask[KEY_APP_KEY] = appKey
	mapTask[KEY_FILE_LINK] = fileLink
	// 新接入请使用4.0版本，已接入（默认2.0）如需维持现状，请注释掉该参数设置。
	mapTask[KEY_VERSION] = "4.0"
	// 设置是否输出词信息，默认为false。开启时需要设置version为4.0。
	mapTask[KEY_ENABLE_WORDS] = "false"
	// 大于16 kHz采样率的音频是否进行自动降采样（降为16 kHz），默认为false，开启时需要设置version为“4.0”。
	mapTask[KEY_ENABLE_SAMPLE_RATE_ADAPTIVE] = "true"
	task, err := json.Marshal(mapTask)
	if err != nil {
		panic(err)
	}
	postRequest.FormParams[KEY_TASK] = string(task)
	postResponse, err := client.ProcessCommonRequest(postRequest)
	if err != nil {
		panic(err)
	}
	postResponseContent := postResponse.GetHttpContentString()
	fmt.Println(postResponseContent)
	if postResponse.GetHttpStatus() != 200 {
		fmt.Println("录音文件识别请求失败，Http错误码: ", postResponse.GetHttpStatus())
		return
	}
	var postMapResult map[string]interface{}
	err = json.Unmarshal([]byte(postResponseContent), &postMapResult)
	if err != nil {
		panic(err)
	}
	var taskId string = ""
	var statusText string = ""
	statusText = postMapResult[KEY_STATUS_TEXT].(string)
	if statusText == STATUS_SUCCESS {
		fmt.Println("录音文件识别请求成功响应!")
		taskId = postMapResult[KEY_TASK_ID].(string)
	} else {
		fmt.Println("录音文件识别请求失败!")
		return
	}
	getRequest := requests.NewCommonRequest()
	getRequest.Domain = DOMAIN
	getRequest.Version = API_VERSION
	getRequest.Product = PRODUCT
	getRequest.ApiName = GET_REQUEST_ACTION
	getRequest.Method = "GET"
	getRequest.QueryParams[KEY_TASK_ID] = taskId
	statusText = ""
	now := time.Now()
	sb := strings.Builder{}
	for true {
		getResponse, err := client.ProcessCommonRequest(getRequest)
		if err != nil {
			panic(err)
		}
		getResponseContent := getResponse.GetHttpContentString()
		fmt.Println("识别查询结果：", getResponseContent)
		if getResponse.GetHttpStatus() != 200 {
			fmt.Println("识别结果查询请求失败，Http错误码：", getResponse.GetHttpStatus())
			break
		}
		var getMapResult map[string]interface{}
		err = json.Unmarshal([]byte(getResponseContent), &getMapResult)
		if err != nil {
			panic(err)
		}
		statusText = getMapResult[KEY_STATUS_TEXT].(string)
		if statusText == STATUS_RUNNING || statusText == STATUS_QUEUEING {
			time.Sleep(10 * time.Second)
		} else {
			resultContentMap := getMapResult[KEY_RESULT].(map[string]interface{})
			sentences := resultContentMap[KEY_SENTENCES].([]interface{})
			for _, rawSentence := range sentences {
				sentence := rawSentence.(map[string]interface{})
				sb.WriteString(sentence[KEY_TEXT].(string))
			}
			break
		}
	}
	if statusText == STATUS_SUCCESS {
		fmt.Println("录音文件识别成功！")
	} else {
		fmt.Println("录音文件识别失败！")
	}
	fmt.Println("识别结果：", sb.String())
	fmt.Println("耗时：", time.Now().Sub(now).Seconds())
}
