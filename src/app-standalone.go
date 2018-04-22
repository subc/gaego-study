// +build !appengine,!appenginevm

package main

import (
	"github.com/labstack/echo"
)

func main() {
    gae_echo.Start()
}
