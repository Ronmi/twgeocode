// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package main

import (
	"encoding/json"
	"log"
	"os"
)

var debug bool

func main() {
	debug = (os.Getenv("DEBUG") == "1")
	db := &dbWriter{}
	err := db.init()
	if err != nil {
		log.Fatal("cannot init db:", err)
	}

	for _, typ := range cityTypes {
		data, err := parseCityCSV("0.csv", typ)
		if err != nil {
			log.Printf("解析 %s 格式的縣市資料失敗: %v", typ, err)
			os.Exit(1)
		}

		fn := "city-" + typ + ".json"
		if err = saveArrayTo(fn, data); err != nil {
			log.Printf("儲存 %s 失敗: %v", fn, err)
			os.Exit(11)
		}

		err = db.save(data, typ)
		if err != nil {
			log.Print("無法儲存進 sqlite:", err)
			os.Exit(111)
		}
	}

	for _, typ := range districtTypes {
		data, err := parseDistrictCSV("1.csv", typ)
		if err != nil {
			log.Printf("解析 %s 格式的鄉鎮資料失敗: %v", typ, err)
			os.Exit(1)
		}

		fn := "district-" + typ + ".json"
		if err = saveArrayTo(fn, data); err != nil {
			log.Printf("儲存 %s 失敗: %v", fn, err)
			os.Exit(11)
		}

		err = db.save(data, typ)
		if err != nil {
			log.Print("無法儲存進 sqlite:", err)
			os.Exit(111)
		}
	}

	for _, typ := range villageTypes {
		data, err := parseVillageCSV("2.csv", typ)
		if err != nil {
			log.Printf("解析 %s 格式的村里資料失敗: %v", typ, err)
			os.Exit(1)
		}

		fn := "village-" + typ + ".json"
		if err = saveArrayTo(fn, data); err != nil {
			log.Printf("儲存 %s 失敗: %v", fn, err)
			os.Exit(11)
		}

		err = db.save(data, typ)
		if err != nil {
			log.Print("無法儲存進 sqlite:", err)
			os.Exit(111)
		}
	}

	log.Print("done")
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
