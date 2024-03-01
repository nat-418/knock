#!/usr/bin/env tclsh

set oses {freebsd openbsd netbsd darwin linux windows}
set arches {386 arm amd64 arm64}

file mkdir dist

foreach os $oses {
	foreach arch $arches {
		if {$os eq "darwin" && $arch eq "386"} { continue }
		exec env GOOS=linux GOARCH=${arch} go build -o ./dist/knock-${os}-${arch}
	}
}

