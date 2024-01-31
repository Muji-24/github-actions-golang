// Copyright (c) 2019, Daniel Martí <mvdan@mvdan.cc>
// See LICENSE for licensing information

package actions

import (
	"fmt"
	"runtime"

	"rsc.io/quote"
)

func Demo() {
	fmt.Printf("Go version: %s\n", runtime.Version())
	fmt.Printf("GOOS: %s\n", runtime.GOOS)
	fmt.Printf("GOARCH: %s\n", runtime.GOARCH)
	fmt.Printf("FYP Pipeline Testing 1.0")
	fmt.Printf("FYP Pipeline Testing 2.0")
	fmt.Printf("FYP Pipeline Testing 3.0")
	fmt.Printf("FYP Pipeline Testing 4.0")

	fmt.Println(quote.Go())
}
