package main

import (
	"github.com/captncraig/seggzy/pkg/web"
)

func main() {
	web.ListenAndServe(":8000", "../../pkg/web")

}
