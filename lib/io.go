package lib

import (
	"io"
	"github.com/Tang-RoseChild/mahonia"
	"io/ioutil"
	"fmt"
)

func DecodeReader(rd io.Reader, charType string) ([]byte, error) {
	dec := mahonia.NewDecoder(charType)
	if dec == nil {
		return nil, fmt.Errorf("%s 编码格式不存在", charType)
	}
	body, err := ioutil.ReadAll(dec.NewReader(rd))
	if err != nil {
		return nil, err
	}
	return body, err
}

func Slice2String(data []string) string {
	var result string
	for index, d := range data {
		if index != len(data)-1 {
			result = result + "," + d
		}
	}
	return result
}
