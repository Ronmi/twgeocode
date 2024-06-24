package main

var (
	cityTypes = []string{
		"100", "103", "106", "112",
	}
	districtTypes = []string{
		"100", "103", "106", "107",
	}
	villageTypes = []string{
		"100", "103", "106", "107", "112",
	}
	cityMap = map[string]string{
		"100": "Taiwan_Geocode_100_縣市代碼",
		"103": "Taiwan_Geocode_103_縣市代碼",
		"106": "Taiwan_Geocode_106_縣市代碼",
		"112": "Taiwan_Geocode_112_縣市代碼",
	}

	districtMap = map[string]string{
		"100": "Taiwan_Geocode_100_鄉鎮代碼",
		"103": "Taiwan_Geocode_103_鄉鎮代碼",
		"106": "Taiwan_Geocode_106_鄉鎮代碼",
		"107": "Taiwan_Geocode_107_鄉鎮代碼",
	}

	villageMap = map[string]string{
		"100": "Taiwan_Geocode_100_村里代碼",
		"103": "Taiwan_Geocode_103_村里代碼",
		"106": "Taiwan_Geocode_106_村里代碼",
		"107": "Taiwan_Geocode_107_村里代碼",
		"112": "Taiwan_Geocode_112_村里代碼",
	}
)
