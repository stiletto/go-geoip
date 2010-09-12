
package main

import (
    "geoip"
    "fmt"
)

func main() {
    gi := geoip.GeoIP_Open("/usr/share/GeoIP/GeoIP.dat")
    if (gi==nil) {
	fmt.Printf("LOL GI NIL\n")
	return
    }
    
    country := gi.GetCountry("94.179.144.22")
    fmt.Printf("LOL COUNTRY %s\n",*country)
    
}

