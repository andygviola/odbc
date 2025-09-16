//go:build darwin && arm64
// +build darwin,arm64

package api

/*
#cgo darwin LDFLAGS: -L/opt/homebrew/opt/unixodbc/lib -lodbc
#cgo darwin CFLAGS: -I/opt/homebrew/opt/unixodbc/include
#include <sql.h>
#include <sqlext.h>
#include <stdint.h>
*/
import "C"
