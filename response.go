package geo_search

import "strconv"

type ServerData struct {
	TotalResultsCount int             `json:"totalResultsData"`
	GeoNames          []ServerGeoName `json:"geonames"`
}

type ServerGeoName struct {
	AdminCode1  string `json:"adminCode1"`
	AdminCode2  string `json:"adminCode2"`
	AdminCode3  string `json:"adminCode3"`
	AdminCode4  string `json:"adminCode4"`
	AdminCode5  string `json:"adminCode5"`
	AdminCodes1 struct {
		ISO3166_2 string `json:"ISO3166_2"`
	} `json:"adminCodes1"`
	CountryCode string `json:"countryCode"`
	Longitude   string `json:"lng"`
	Latitude    string `json:"lat"`
	GeoNameID   int    `json:"geonameId"`
	FCL         string `json:"fcl"`
	Population  int    `json:"population"`
	Name        string `json:"name"`
	FCLName     string `json:"fclName"`
	CountryName string `json:"countryName"`
	FCodeName   string `json:"FCodeName"`
	AdminName1  string `json:"adminName1"`
	ToponymName string `json:"toponymName"`
	FCode       string `json:"fcode"`
}

func (g *ServerGeoName) Compile() GeoData {
	return GeoData{
		AdminCodes:  g.wrapAdminCodes(),
		ISO3166_2:   g.AdminCodes1.ISO3166_2,
		CountryCode: g.CountryCode,
		Name:        g.Name,
		Country:     g.CountryName,
		Admin:       g.AdminName1,
		Latitude:    stringToFload(g.Latitude),
		Longitude:   stringToFload(g.Longitude),
		Population:  g.Population,
	}
}

func (d *ServerData) Parse() []GeoData {
	var data []GeoData
	for _, r := range d.GeoNames {
		data = append(data, r.Compile())
	}
	return data
}

func stringToFload(s string) float64 {
	lat, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0
	}
	return lat
}

func (g *ServerGeoName) wrapAdminCodes() []string {
	var codes []string
	if g.AdminCode1 != "" {
		codes = append(codes, g.AdminCode1)
	}
	if g.AdminCode2 != "" {
		codes = append(codes, g.AdminCode2)
	}
	if g.AdminCode3 != "" {
		codes = append(codes, g.AdminCode3)
	}
	if g.AdminCode4 != "" {
		codes = append(codes, g.AdminCode4)
	}
	if g.AdminCode5 != "" {
		codes = append(codes, g.AdminCode5)
	}
	return codes
}
