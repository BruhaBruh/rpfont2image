package entity

type BitMapProvider struct {
	Type  string   `json:"type"`
	File  string   `json:"file"`
	Chars []string `json:"chars"`
}

func BitMapProviders(providers []map[string]interface{}) []BitMapProvider {
	var bitMapProviders []BitMapProvider

	for _, provider := range providers {
		if provider["type"] != "bitmap" {
			continue
		}
		bitMapProvider := BitMapProvider{
			Type: "bitmap",
		}

		switch file := provider["file"].(type) {
		case string:
			bitMapProvider.File = file
		}

		switch chars := provider["chars"].(type) {
		case []interface{}:
			bitMapProvider.Chars = toStringArray(chars)
		}

		bitMapProviders = append(bitMapProviders, bitMapProvider)
	}

	return bitMapProviders
}

func toStringArray(arr []interface{}) []string {
	var a []string

	for _, v := range arr {
		switch c := v.(type) {
		case string:
			a = append(a, c)
		}
	}

	return a
}
