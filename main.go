// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"os"
	"strings"
)

var debug bool

func main() {
	debug = (os.Getenv("DEBUG") == "1")
	city, err := parseCityCSV("0.csv")
	if err != nil {
		os.Exit(1)
	}
	district, err := parseDistrictCSV("1.csv")
	if err != nil {
		os.Exit(2)
	}
	village, err := parseVillageCSV("2.csv")
	if err != nil {
		os.Exit(3)
	}
	all := make([]*Info, 0, len(city)+len(district)+len(village))
	all = append(all, city...)
	all = append(all, district...)
	all = append(all, village...)

	if err = saveArrayTo("city.json", city); err != nil {
		os.Exit(11)
	}
	if err = saveArrayTo("district.json", district); err != nil {
		os.Exit(12)
	}
	if err = saveArrayTo("village.json", village); err != nil {
		os.Exit(13)
	}
	if err = saveArrayTo("all.json", all); err != nil {
		os.Exit(14)
	}

	// 檢查是否有重覆的代碼，正常應該不會有
	m := map[string]*Info{}
	for _, i := range all {
		if v, ok := m[i.Code]; ok {
			log.Printf("抓到重覆代碼\n    old: %+v\n    new: %+v", v, i)
			continue
		}
		m[i.Code] = i
	}
	if len(m) != len(all) {
		os.Exit(999)
	}

	log.Print("done")
}

type Info struct {
	Code     string `json:"code"`
	English  string `json:"english"`
	Name     string `json:"name"`
	FullName string `json:"fullname"`
}

func parseCSV(fn string, suffixes []string) (ret []map[string]string, err error) {
	buf, err := os.ReadFile(fn)
	if err != nil {
		log.Printf("無法讀取原始檔案 %s: %v", fn, err)
		return
	}

	arr := bytes.Split(buf, []byte("\n"))
	if len(arr) < 2 {
		log.Print("原始檔案內容似乎不正確")
		return nil, errors.New("incorrect source")
	}

	// 分析 header
	indexes := make([]int, len(suffixes))
	for idx := range suffixes {
		indexes[idx] = -1
	}
	cols := bytes.Split(arr[0], []byte(","))

	for idx, col := range cols {
		str := string(col)
		for sidx, suffix := range suffixes {
			if str == suffix {
				if debug {
					log.Printf("DEBUG: %s: %d (%s)", suffix, idx, str)
				}
				indexes[sidx] = idx
				break
			}
		}
	}
	for _, v := range indexes {
		if v == -1 {
			log.Print("原始檔案的欄位格式不符預期")
			return nil, errors.New("incorrect column index")
		}
	}

	arr = arr[1:]
	for _, record := range arr {
		cols := strings.Split(string(record), ",")
		if len(cols) == 1 {
			// 空行
			continue
		}

		m := map[string]string{}
		empty := false
		for key, idx := range indexes {
			if cols[idx] == "" {
				empty = true
				continue
			}
			m[suffixes[key]] = cols[idx]
		}
		if empty {
			// 有空白資料
			continue
		}
		ret = append(ret, m)
	}
	return
}

func parseCityCSV(fn string) (ret []*Info, err error) {
	data, err := parseCSV(fn, []string{
		"Taiwan_Geocode_103_縣市代碼",
		"Taiwan_Geocode_103_縣市英文名",
		"Taiwan_Geocode_103_縣市全名",
		"Taiwan_Geocode_103_縣市名",
	})
	if err != nil {
		return
	}

	ret = make([]*Info, len(data))
	for idx, record := range data {
		ret[idx] = &Info{
			Code:     record["Taiwan_Geocode_103_縣市代碼"],
			English:  record["Taiwan_Geocode_103_縣市英文名"],
			FullName: record["Taiwan_Geocode_103_縣市全名"],
			Name:     record["Taiwan_Geocode_103_縣市名"],
		}
	}

	return
}

func parseDistrictCSV(fn string) (ret []*Info, err error) {
	data, err := parseCSV(fn, []string{
		"Taiwan_Geocode_103_鄉鎮代碼",
		"Taiwan_Geocode_103_鄉鎮英文名",
		"Taiwan_Geocode_103_縣市鄉鎮名",
		"Taiwan_Geocode_103_鄉鎮名",
	})
	if err != nil {
		return
	}

	ret = make([]*Info, len(data))
	for idx, record := range data {
		ret[idx] = &Info{
			Code:     record["Taiwan_Geocode_103_鄉鎮代碼"],
			English:  record["Taiwan_Geocode_103_鄉鎮英文名"],
			FullName: record["Taiwan_Geocode_103_縣市鄉鎮名"],
			Name:     record["Taiwan_Geocode_103_鄉鎮名"],
		}
	}

	return
}

func parseVillageCSV(fn string) (ret []*Info, err error) {
	data, err := parseCSV(fn, []string{
		"Taiwan_Geocode_103_村里代碼",
		"Taiwan_Geocode_107_村里英文名稱",
		"Taiwan_Geocode_103_縣市名",
		"Taiwan_Geocode_103_鄉鎮名",
		"Taiwan_Geocode_103_村里名",
	})
	if err != nil {
		return
	}

	ret = make([]*Info, len(data))
	for idx, record := range data {
		ret[idx] = &Info{
			Code:    record["Taiwan_Geocode_103_村里代碼"],
			English: record["Taiwan_Geocode_107_村里英文名稱"],
			FullName: record["Taiwan_Geocode_103_縣市名"] +
				record["Taiwan_Geocode_103_鄉鎮名"] +
				record["Taiwan_Geocode_103_村里名"],
			Name: record["Taiwan_Geocode_103_村里名"],
		}
	}

	return
}

func saveArrayTo(fn string, data []*Info) (err error) {
	f, err := os.Create(fn)
	if err != nil {
		return
	}
	defer f.Close()
	enc := json.NewEncoder(f)

	if debug {
		log.Printf("DEBUG: save to array: 共 %d 筆", len(data))
	}
	return enc.Encode(data)
}
