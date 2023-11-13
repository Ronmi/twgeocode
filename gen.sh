#!/usr/bin/sh
# This Source Code Form is subject to the terms of the Mozilla Public
# License, v. 2.0. If a copy of the MPL was not distributed with this
# file, You can obtain one at https://mozilla.org/MPL/2.0/.

set -ex

# install tool
go install github.com/tealeg/xlsx2csv@latest

# download source file
curl -sSL 'https://alerts.ncdr.nat.gov.tw/Document/%E8%A1%8C%E6%94%BF%E5%8D%80%E4%BB%A3%E7%A2%BC%E8%A1%A8_Taiwan_Geocode.xlsx' -o source.xlsx

# extract source csv and convert to json
xlsx2csv -d , -i 0 -o 0.csv source.xlsx
xlsx2csv -d , -i 1 -o 1.csv source.xlsx
xlsx2csv -d , -i 2 -o 2.csv source.xlsx

# run my prog to generate data
go run *.go
