package geoip

/*
#include <stdio.h>
#include <errno.h>
#include <GeoIP.h>

//typedef GeoIP* GeoIP_pnt
*/
import "C"
import (
    "unsafe"
)

type GeoIP struct {
    gi *C.GeoIP
}

func GeoIP_Open(base string) (*GeoIP) {
    cbase := C.CString(base)
    gi := C.GeoIP_open(cbase, C.GEOIP_STANDARD | C.GEOIP_CHECK_CACHE);
    C.free(unsafe.Pointer(cbase))
    if (gi==nil) {
	return nil
    }
    return &GeoIP{gi}
}

func (gi *GeoIP) GetCountry(ip string) (*string) {
    if (gi==nil) {
	return nil
    }
    cip := C.CString(ip)
    ccountry := C.GeoIP_country_code_by_addr(gi.gi,cip)
    C.free(unsafe.Pointer(cip))
    if (ccountry!=nil) {
        rets := C.GoString(ccountry)
        return &rets
    }
    return nil
}

/*func (gi *GeoIP) GetCountry_v6(ip string) (*string) {
    if (gi==nil) {
	return nil
    }
    cip := C.CString(ip)
    ccountry := C.GeoIP_country_code_by_addr_v6(gi.gi,cip)
    C.free(unsafe.Pointer(cip))
    if (ccountry!=nil) {
        rets := C.GoString(ccountry)
        return &rets
    }
    return nil
}*/
