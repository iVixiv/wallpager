package download

import (
	"fmt"
	"os"
	"net/http"
	"io/ioutil"
	"io"
	"bytes"
)

func DownloadImg(id string, url string) () {
	go func() {
		out, err := os.Create("H:\\wallpager\\" + id + ".jpg")
		if err != nil {
			fmt.Printf("download err %s \n", err.Error())
		}
		defer out.Close()
		resp, err := http.Get(url)
		defer resp.Body.Close()
		pix, err := ioutil.ReadAll(resp.Body)
		_, err = io.Copy(out, bytes.NewReader(pix))
		if err != nil {
			fmt.Printf("download err %s \n", err.Error())
		}
	}()
}
