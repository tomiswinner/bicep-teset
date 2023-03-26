// ソースコード元
// https://qiita.com/ozora/items/0597e52b3f9c1759e292
// FUNCTIONS_CUSTOMHANDLER_PORT は、Azure Functions の環境変数
// https://learn.microsoft.com/ja-jp/azure/azure-functions/create-first-function-vs-code-other?tabs=go%2Cwindows#create-and-build-your-function


package main

import (
    "os"
    "github.com/gin-gonic/gin"
    _ "github.com/microsoft/go-mssqldb"
	"database/sql"
	"context"
	"log"
	"fmt"
)

var db *sql.DB
var server = os.Getenv("SQLSERVER")
var port = 1433
var user = os.Getenv("SQLUSER")
var password = os.Getenv("SQLPASSWORD")
var database = os.Getenv("SQLDATABASE")

func main() {
    r := gin.Default()
    r.GET("/api/message1", func(c *gin.Context) {
        var ping = dbconnect() + "1"
        c.JSON(200, gin.H{
            "message1": ping,
        })
    })
    r.GET("/api/message2", func(c *gin.Context) {
        var pong = dbconnect() + "2"
        c.JSON(200, gin.H{
            "message2": pong,
        })
    })
		r.POST("/api/slacktest", func(c *gin.Context) {
			var req map[string]string
			c.BindJSON(&req)
			c.String(200, req["challenge"])
		})


    port := os.Getenv("FUNCTIONS_CUSTOMHANDLER_PORT")
    r.Run(":" + port)
}



func dbconnect() string {
	// Build connection string
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		server, user, password, port, database)
	var err error
	// Create connection pool
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}
	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	return "Connected" 
}
