package crawler

import (
	"wallpager/lib"
	"net/http"
	"net/url"
	"strings"
	"encoding/json"
	"wallpager/db"
	"fmt"
	"github.com/op/go-logging"
	"strconv"
)

var log = logging.MustGetLogger("crawler")

func Crawl(count int) (error) {
	for i := 0; i <= count/21; i += 1 {
		err := Request(i * 21)
		if err != nil {
			return err
		}
	}
	return nil
}

func Request(skip int) (error) {
	params := url.Values{} // create URLParams，required params: phone, password
	params.Add("limit", "21")
	params.Add("adult", "false")
	params.Add("first", "0")
	params.Add("order", "new")
	params.Add("skip", strconv.Itoa(skip))

	headers := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	res, err := lib.Request(db.CRAWL_URL+"/wallpaper", http.MethodPost, headers, strings.NewReader(params.Encode()), "UTF-8")
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

func save(datas []Wallpager) {
	ch := make(chan int, 100)
	var insertStr string
	for index, data := range datas {
		var time int
		time = int(data.Atime * 1000)

		insertStr += "(" + "\"" + data.Id + "\"" + "," + strconv.Itoa(data.Rank) + "," + strconv.Itoa(data.Favs) + "," + "\"" + lib.Slice2String(data.Cid) + "\"" + "," + "\"" + lib.Slice2String(data.Tag) + "\"" + "," + "\"" + data.Wp + "\"" + "," + "\"" + data.Thumb + "\"" + "," + "\"" + data.Img + "\"" + "," + "\"" + data.Preview + "\"" + "," + "\"" + data.Desc + "\"" + "," + strconv.Itoa(time) + ")"
		if index != len(datas)-1 {
			insertStr += ","
		}
	}
	saveDb(insertStr, ch)
}

func saveDb(inStr string, ch chan int) {
	ch <- 1
	go func() {
		res, err := db.MySQL.Exec(`INSERT INTO wallpager (wid,rank,favs,cid,tag,wp,thumb,img,preview,wdesc,atime) VALUES` + inStr)
		if err != nil {
			log.Warningf("%s", err.Error())
			return
		}
		id, err := res.LastInsertId()
		if err != nil {
			fmt.Printf(" %d ,error :%s", id, err.Error())
		} else {
			fmt.Printf("insert %d \n", id)
		}
		<-ch
	}()
}

type Response struct {
	Res struct {
		Wallpaper []Wallpager
	}
}

type Wallpager struct {
	Id      string
	Rank    int
	Favs    int
	Atime   float32
	Desc    string
	Thumb   string
	Wp      string
	Img     string
	Preview string
	Cid     [] string
	Tag     []string
}
