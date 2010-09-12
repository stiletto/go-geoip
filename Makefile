# Copyright 2009 The Go Authors.  All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

include $(GOROOT)/src/Make.inc

TARG=geoip
CGOFILES=\
	geoip.go

CGO_LDFLAGS=-lGeoIP $(LDPATH_$(GOOS))
#CLEANFILES+=hello fib chain run.out

include $(GOROOT)/src/Make.pkg

%: install %.go
	$(QUOTED_GOBIN)/$(GC) $*.go
	$(QUOTED_GOBIN)/$(LD) -o $@ $*.$O
