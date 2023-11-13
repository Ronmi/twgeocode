把 [民生示警平台](https://alerts.ncdr.nat.gov.tw/CAPfiledownload.aspx) 提供的行政區代碼表轉換成易於在程式中使用的格式

# 提供的格式及版本

此專案目前僅提供 JSON array 格式的 Taiwan Geocode 103 代碼對照表，依內容共分四種

- city: 相當於原檔案中的縣市分頁
- district: 相當於原檔案中的鄉鎮分頁
- village: 相當於原檔案中的村里分頁
- all: 以上三個檔案合在一起

四個 JSON 檔案都是從 [民生示警平台](https://alerts.ncdr.nat.gov.tw/CAPfiledownload.aspx) 提供的行政區代碼表轉換而來，原檔案的更新日期為 2023.09.06

# 檔案是如何產生的

## 使用 docker (推薦)

在 Linux 環境下可以直接執行 `docker.sh`

Windows 及 MacOS 使用者可以自行調整參數或手動下指令，例如拿掉 user 相關的參數並寫死路徑

```sh
docker run -it --rm -v "/path/to/colned/repo:/twgeocode" --workdir /twgeocode golang /twgeocode/gen.sh
```

## 直接產生

在 Linux 環境下，安裝好 curl 及 golang 環境後，執行 `gen.sh` 即可

# 版權宣告

- 程式碼部份 (Go 語言及 shell script) 使用 MPLv2 授權
- 原始對照表檔案的授權依 [民生示警平台之授權規定](https://alerts.ncdr.nat.gov.tw/usestandard.aspx)
- 執行本專案的過程產生的 csv 與 json 檔案使用 CC0 授權
