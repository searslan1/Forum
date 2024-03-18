package main

import (
	"fmt"
	"strings"
)

func main() {
	girdi := "Selam, asko!"
	parçalar := strings.Split(girdi, ",")
	for _, parça := range parçalar {
		fmt.Println(parça)
	}
}
