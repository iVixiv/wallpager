package wallpager

import (
	"wallpager/db"
	"fmt"
)

func Select(start string, mType string) ([]*WallPager, error) {
	fmt.Println("start ：", start)
	res := make([]*WallPager, 0)
	sql := `SELECT wp,tag FROM wallpager ORDER BY %s LIMIT %s,30 `
	if mType == "rank" {
		sql = fmt.Sprintf(sql, "rank", start)
	} else {
		sql = fmt.Sprintf(sql, "uid", start)
	}
	rows, err := db.MySQL.Query(sql)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var tag string
		var wp string
		var wdesc string
		err = rows.Scan(&tag, &wp)
		res = append(res, &WallPager{
			Tag:  tag,
			Pic:  wp,
			Desc: wdesc,
		})
	}
	fmt.Println("end ：", len(res))
	return res, nil
}

type WallPager struct {
	Pic  string `json:"pic"`
	Tag  string `json:"tag"`
	Desc string `json:"desc"`
}
