/* 
 * Original author: Stiletto <blasux@blasux.ru>
 *
 * This program is free software. It comes without any warranty, to
 * the extent permitted by applicable law. You can redistribute it
 * and/or modify it under the terms of the Do What The Fuck You Want
 * To Public License, Version 2, as published by Sam Hocevar. See
 * http://sam.zoy.org/wtfpl/COPYING for more details. */

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

