package main

import (
	shortener "link-shortener/internal/app/link_shortener"
	"link-shortener/internal/pkg/utils"
)

func init() {
	utils.LoadEnv()
}

func main() {
	shortener.Start()
}
