//go:build darwin && amd64
// +build darwin,amd64

package api

/*
#cgo darwin LDFLAGS: -L/usr/local/opt/unixodbc/lib -lodbc
#cgo darwin CFLAGS: -I/usr/local/opt/unixodbc/include
#include <sql.h>
#include <sqlext.h>
#include <stdint.h>
*/
import "C"
