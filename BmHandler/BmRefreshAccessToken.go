package BmHandler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	
	"reflect"
	"strings"

	"github.com/alfredyang1986/BmServiceDef/BmDaemons"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmMongodb"

	"github.com/julienschmidt/httprouter"
)
type RefreshAccessTokenHandler struct {
	Method     string
	HttpMethod string
	Args       []string
	db         *BmMongodb.BmMongodb
}


func (h RefreshAccessTokenHandler) NewRefreshAccessTokenHandler(args ...interface{}) RefreshAccessTokenHandler {
	var m *BmMongodb.BmMongodb
	var hm string
	var md string
	var ag []string
	for i, arg := range args {
		if i == 0 {
			sts := arg.([]BmDaemons.BmDaemon)
			for _, dm := range sts {
				tp := reflect.ValueOf(dm).Interface()
				tm := reflect.ValueOf(tp).Elem().Type()
				if tm.Name() == "BmMongodb" {
					m = dm.(*BmMongodb.BmMongodb)
				}
			}
		} else if i == 1 {
			md = arg.(string)
		} else if i == 2 {
			hm = arg.(string)
		} else if i == 3 {
			lst := arg.([]string)
			for _, str := range lst {
				ag = append(ag, str)
			}
		} else {
		}
	}
	return RefreshAccessTokenHandler{Method: md, HttpMethod: hm, Args: ag, db: m}
}

func (h RefreshAccessTokenHandler) RefreshAccessToken(w http.ResponseWriter, r *http.Request, _ httprouter.Params) int {
	// 拼接转发的URL
	scheme := "http://"
	if r.TLS != nil {
		scheme = "https://"
	}
	version := "v0"
	resource := fmt.Sprint("192.168.100.174:9096", "/"+version+"/", "RefreshAccessToken", "?", r.URL.RawQuery)
	mergeURL := strings.Join([]string{scheme, resource}, "")

	// 转发
	client := &http.Client{}
	req, _ := http.NewRequest("GET", mergeURL, nil)
	for k, v := range r.Header {
		req.Header.Add(k, v[0])
	}
	response, err := client.Do(req)
	if err != nil {
		fmt.Println("Fuck Error")
	}
	result, err := ioutil.ReadAll(response.Body)
	data := map[string]interface{}{}
	err = json.Unmarshal(result, &data)
	if err != nil {
		fmt.Println("RefreshAccessToken Error")
	}

	enc := json.NewEncoder(w)
	enc.Encode(data)
	return 0
}

func (h RefreshAccessTokenHandler) GetHttpMethod() string {
	return h.HttpMethod
}

func (h RefreshAccessTokenHandler) GetHandlerMethod() string {
	return h.Method
}
