package osrm

import (
	"fmt"
	"github.com/codingsince1985/geo-golang"
	"github.com/codingsince1985/geo-golang/openstreetmap"
	"github.com/kr/pretty"
)

const (
	addr     = "Melbourne VIC"
	lat, lng = -37.813611, 144.963056
)

func ReverseGeocode() {
	try(openstreetmap.Geocoder())
}

func try(geocoder geo.Geocoder) {
	location, _ := geocoder.Geocode(addr)
	if location != nil {
		fmt.Printf("%s location is (%.6f, %.6f)\n", addr, location.Lat, location.Lng)
	} else {
		fmt.Println("got <nil> location")
	}
	address, _ := geocoder.ReverseGeocode(lat, lng)
	if address != nil {
		fmt.Printf("Address of (%.6f,%.6f) is %s\n", lat, lng, address.FormattedAddress)
		pretty.Println(address)
	} else {
		fmt.Println("got <nil> address")
	}
	fmt.Print("\n")
}
