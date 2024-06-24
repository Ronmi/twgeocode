package main

import (
	"bytes"
	"errors"
	"log"
	"os"
	"strings"
)

type Info struct {
	Code     string `json:"code"`
	English  string `json:"english"`
	Name     string `json:"name"`
	Fullname string `json:"fullname"`
}

var cache = map[string][]byte{}

func parseCSV(fn string, suffixes []string) (ret [][]string, err error) {
	buf, ok := cache[fn]
	if !ok {
		buf, err = os.ReadFile(fn)
		if err != nil {
			log.Printf("無法讀取原始檔案 %s: %v", fn, err)
			return
		}
		cache[fn] = buf
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

		m := make([]string, len(indexes))
		empty := false
		for key, idx := range indexes {
			if cols[idx] == "" {
				empty = true
				continue
			}
			m[key] = cols[idx]
		}
		if empty {
			// 有空白資料
			continue
		}
		ret = append(ret, m)
	}
	return
}

func parseCityCSV(fn, typ string) (ret []*Info, err error) {
	data, err := parseCSV(fn, []string{
		cityMap[typ],
		"Taiwan_Geocode_112_縣市英文名",
		"Taiwan_Geocode_112_縣市全名",
		"Taiwan_Geocode_112_縣市名",
	})
	if err != nil {
		return
	}

	ret = make([]*Info, len(data))
	for idx, record := range data {
		ret[idx] = &Info{
			Code:     record[0],
			English:  record[1],
			Fullname: record[2],
			Name:     record[3],
		}
	}

	return
}

func parseDistrictCSV(fn, typ string) (ret []*Info, err error) {
	data, err := parseCSV(fn, []string{
		districtMap[typ],
		"Taiwan_Geocode_107_鄉鎮英文名",
		"Taiwan_Geocode_107_縣市鄉鎮名",
		"Taiwan_Geocode_107_鄉鎮名",
	})
	if err != nil {
		return
	}

	ret = make([]*Info, len(data))
	for idx, record := range data {
		ret[idx] = &Info{
			Code:     record[0],
			English:  record[1],
			Fullname: record[2],
			Name:     record[3],
		}
	}

	return
}

func parseVillageCSV(fn, typ string) (ret []*Info, err error) {
	data, err := parseCSV(fn, []string{
		villageMap[typ],
		"Taiwan_Geocode_112_村里英文名稱",
		"Taiwan_Geocode_112_縣市名",
		"Taiwan_Geocode_112_鄉鎮名",
		"Taiwan_Geocode_112_村里名",
	})
	if err != nil {
		return
	}

	ret = make([]*Info, len(data))
	for idx, record := range data {
		ret[idx] = &Info{
			Code:    record[0],
			English: record[1],
			Fullname: record[2] +
				record[3] +
				record[4],
			Name: record[4],
		}
	}

	return
}
