package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"golangstudy/jike/gingormpractise/common"
)

func main() {
	db := common.InitDB()
	defer db.Close()
	r := gin.Default()
	r=CollectRoute(r)
	panic(r.Run())
}

