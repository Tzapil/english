package main

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tzapil/english/collections"
	"github.com/tzapil/english/common"
	"github.com/tzapil/english/ping"
	"github.com/tzapil/english/words"
)

type Collection struct {
	_id  uint
	Name string
	Date time.Time
}

type Trainer struct {
	name string
	age  uint
	city string
}

func serve() {
	// creating of new router
	r := gin.Default()

	// make all handlers v1 api version
	v1 := r.Group("/api/v1")

	collections.CollectionsRegister(v1)
	words.WordsRegister(v1)
	ping.PingRegister(v1)

	// run default listen and serve on 0.0.0.0:8080
	r.Run()
}

func main() {
	common.Init()
	defer common.Close()

	serve()
}
