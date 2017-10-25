package crawler

import (
	"wallpager/lib"
	"net/http"
	"net/url"
	"strings"
	"encoding/json"
	"strconv"
	"wallpager/db"
)

func Crawl_Type(count int) (error) {
	for i := 0; i <= count/21; i += 1 {
		err := Request_Type(i*21, db.CRAWL_URL+"/category/4e4d610cdf714d2966000003/wallpaper")
		if err != nil {
			return err
		}
		err = Request_Type(i*21, db.CRAWL_URL+"/category/4e4d610cdf714d2966000000/wallpaper")
		if err != nil {
			return err
		}
		err = Request_Type(i*21, db.CRAWL_URL+"/category/4e4d610cdf714d2966000002/wallpaper")
		if err != nil {
			return err
		}
		err = Request_Type(i*21, db.CRAWL_URL+"/category/4e4d610cdf714d2966000007/wallpaper")
		if err != nil {
			return err
		}
		err = Request_Type(i*21, db.CRAWL_URL+"/category/5109e04e48d5b9364ae9ac45/wallpaper")
		if err != nil {
			return err
		}
		err = Request_Type(i*21, db.CRAWL_URL+"/category/4fb479f75ba1c65561000027/wallpaper")
		if err != nil {
			return err
		}
		err = Request_Type(i*21, db.CRAWL_URL+"/category/4ef0a35c0569795756000000/wallpaper")
		if err != nil {
			return err
		}
		err = Request_Type(i*21, db.CRAWL_URL+"/category/4fb47a195ba1c60ca5000222/wallpaper")
		if err != nil {
			return err
		}
		err = Request_Type(i*21, db.CRAWL_URL+"/category/5109e05248d5b9368bb559dc/wallpaper")
		if err != nil {
			return err
		}

		err = Request_Type(i*21, db.CRAWL_URL+"/category/4fb47a465ba1c65561000028/wallpaper")
		if err != nil {
			return err
		}

		err = Request_Type(i*21, db.CRAWL_URL+"/category/4ef0a3330569795757000000/wallpaper")
		if err != nil {
			return err
		}
		err = Request_Type(i*21, db.CRAWL_URL+"/category/4e4d610cdf714d2966000006/wallpaper")
		if err != nil {
			return err
		}
		err = Request_Type(i*21, db.CRAWL_URL+"/category/4e4d610cdf714d2966000004/wallpaper")
		if err != nil {
			return err
		}
		err = Request_Type(i*21, db.CRAWL_URL+"/category/4e4d610cdf714d2966000005/wallpaper")
		if err != nil {
			return err
		}
		err = Request_Type(i*21, db.CRAWL_URL+"/category/4fb47a305ba1c60ca5000223/wallpaper")
		if err != nil {
			return err
		}
		err = Request_Type(i*21, db.CRAWL_URL+"/category/4e4d610cdf714d2966000001/wallpaper")
		if err != nil {
			return err
		}
		err = Request_Type(i*21, db.CRAWL_URL+"/category/4ef0a34e0569795757000001/wallpaper")
		if err != nil {
			return err
		}
		err = Request_Type(i*21, db.CRAWL_URL+"/category/4e58c2570569791a19000000/wallpaper")
		if err != nil {
			return err
		}
	}
	return nil
}

func Request_Type(skip int, urls string) (error) {
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
		return err
	}
	result := &Response{}
	err = json.Unmarshal([]byte(res), result)
	if err != nil {
		return err
	}
	save(result.Res.Wallpaper)
	return nil
}
