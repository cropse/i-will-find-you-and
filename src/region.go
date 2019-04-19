package main

import (
	"log"
	"net"

	geoip2 "github.com/oschwald/geoip2-golang"
)

type regionInfo struct {
	Code int `json:"code"`
	Data `json:"data"`
}

// Data fuck you
type Data struct {
	Area      string `json:"area"`
	AreaID    string `json:"area_id"`
	Country   string `json:"country"`
	CountryID string `json:"country_id"`
	Region    string `json:"region"`
	RegionID  string `json:"region_id"`
	City      string `json:"city"`
	CityID    string `json:"city_id"`
	IP        string `json:"ip"`
}

func getRegionByIP(rawIP string) (*regionInfo, error) {
	db, err := geoip2.Open("GeoLite2-City.mmdb")
	ip := net.ParseIP(rawIP)
	record, err := db.City(ip)
	defer db.Close()
	if err != nil {
		log.Printf("%s", err)
		// return empty for invalid ip
		return &regionInfo{}, err
	}
	log.Printf("%+v\n", record)
	region := &regionInfo{
		Code: 1,
		Data: Data{
			Area:      record.Continent.Names["zh-CN"],
			AreaID:    record.Continent.Code,
			Region:    "",
			RegionID:  "",
			Country:   record.Country.Names["zh-CN"],
			CountryID: record.Country.IsoCode,
			City:      record.City.Names["zh-CN"],
			CityID:    record.City.Names["en"],
			IP:        rawIP,
		},
	}
	if len(record.Subdivisions) > 0 {
		region.Data.Region = record.Subdivisions[0].Names["zh-CN"]
		region.Data.RegionID = record.Subdivisions[0].IsoCode
	}
	return region, nil
}
