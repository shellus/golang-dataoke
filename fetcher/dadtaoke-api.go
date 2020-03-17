package fetcher

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"sync"
)

type apiConfig struct {
	App_secret string
	App_key    string
}

var config *apiConfig

func SetConfig(s string, k string) {
	config.App_secret = s
	config.App_key = k
}

const url = "http://openapi.dataoke.com"

// 只允许 10,50,100,200
const pageSize = 100

func GetTotalPage() int {
	items, err := Paginator(1)
	if err != nil {
		panic(err.Error())
	}
	totalNum := items.TotalNum
	total := int(math.Ceil(float64(totalNum / pageSize)))
	return total
}

func PaginatorIterator(i chan int, o chan *Items, wg *sync.WaitGroup) {
	for p := range i {
		items, err := Paginator(p)
		if err != nil {
			fmt.Printf("抓取异常: page %s", p)
			continue
		}
		o <- items
	}
	fmt.Println("抓取线程退出")
	wg.Done()
}

func Paginator(p int) (*Items, error) {
	options := map[string]string{
		"pageId":   strconv.Itoa(p),
		"pageSize": strconv.Itoa(pageSize),
	}
	res, err := Get("/api/goods/get-goods-list", options)
	if err != nil {
		return nil, err
	}
	var items *Items
	err = json.Unmarshal(*res.Data, &items)
	if err != nil {
		return nil, errors.New("ITEMS JSON ERR:" + err.Error())
	}
	return items, nil
}

func makeSign(data map[string]string) string {
	var arr []string
	for k, v := range data {
		arr = append(arr, k+"="+v)
	}

	sort.Sort(sort.StringSlice(arr))

	arr = append(arr, "key="+config.App_secret)
	str := strings.Join(arr, "&")
	//str = str + "&key=" + app_secret
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}

type RequestErr error

/**
 * @param string $uri
 * @param array $params
 */
func Get(uri string, params map[string]string) (*Result, RequestErr) {
	params["appKey"] = config.App_key
	params["version"] = "v1.2.0"
	params["sign"] = makeSign(params)
	return request("GET", uri, options{query: params})
}

type options struct {
	query map[string]string
	json  map[string]string
}

func request(method string, uri string, options options) (*Result, RequestErr) {

	isJSON := false
	var body []byte
	if options.json != nil {
		isJSON = true
		var err error
		body, err = json.Marshal(options.json)
		if err != nil {
			panic("JSON ERR:" + err.Error())
		}
	}

	request, err := http.NewRequest(method, url+uri, bytes.NewReader(body))
	if err != nil {
		panic("JSON ERR:" + err.Error())
	}
	query := request.URL.Query()
	for k, v := range options.query {
		query.Add(k, v)
	}
	request.URL.RawQuery = query.Encode()

	if isJSON == true {
		request.Header.Set("Content-Type", "application/json")
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		panic("JSON ERR:" + err.Error())
	}

	//log.Print(fmt.Sprintf("API请求：%s, %s, %s", method, url+uri, options))

	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic("BODY ERR:" + err.Error())
	}

	var res *Result

	err = json.Unmarshal(bodyBytes, &res)
	if err != nil {
		panic("BODY JSON ERR:" + err.Error())
	}
	if res.Code != 0 {
		return nil, fmt.Errorf("BODY JSON CODE != 0: %s", res.Msg)
	}

	return res, nil
}
