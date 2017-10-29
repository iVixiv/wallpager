package crawler

import (
	"wallpager/lib"
	"net/http"
	"net/url"
	"strings"
	"encoding/json"
	"strconv"
	"wallpager/db"
	"fmt"
)

var (
	types = []string{"4e4d610cdf714d2966000003",
		"4e4d610cdf714d2966000000",
		"4e4d610cdf714d2966000002",
		"4e4d610cdf714d2966000007",
		"5109e04e48d5b9364ae9ac45",
		"4fb479f75ba1c65561000027",
		"4ef0a35c0569795756000000",
		"4fb47a195ba1c60ca5000222",
		"5109e05248d5b9368bb559dc",
		"4fb47a465ba1c65561000028",
		"4ef0a3330569795757000000",
		"4e4d610cdf714d2966000006",
		"4e4d610cdf714d2966000004",
		"4e4d610cdf714d2966000005",
		"4fb47a305ba1c60ca5000223",
		"4e4d610cdf714d2966000001",
		"4ef0a34e0569795757000001",
		"4e58c2570569791a19000000"}
)

func Crawl_Type() (error) {
	for k := 0; k < len(types); k++ {
		sum := 1
		lastId := ""
		isNext := true
		for isNext {
			id, size := Request_Type(lastId, sum, db.CRAWL_URL+"/category/"+types[k]+"/wallpaper")
			sum += size
			if id == lastId {
				isNext = false
			} else {
				lastId = id
			}
		}
		fmt.Printf("request Over %s, Sum %d", types[k], sum)
	}
	return nil
}

func Request_Type(lastId string, skip int, urls string) (string, int) {
	params := url.Values{} // create URLParams，required params: phone, password
	params.Add("limit", "21")
	params.Add("adult", "false")
	params.Add("first", "0")
	params.Add("order", "new")
	params.Add("skip", strconv.Itoa(skip))

	headers := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	res, err := lib.Request(urls, http.MethodPost, headers, strings.NewReader(params.Encode()), "UTF-8")
	if err != nil {
		return "", 0
	}
	result := &Response{}
	err = json.Unmarshal([]byte(res), result)
	if err != nil {
		return "", 0
	}
	save(result.Res.Wallpaper)
	return result.Res.Wallpaper[len(result.Res.Wallpaper)-1].Id, len(result.Res.Wallpaper)
}
