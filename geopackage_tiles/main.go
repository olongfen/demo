package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
	"strings"
)

func main() {
	app := gin.Default()
	app.GET("/:z/:x/:y.png", func(c *gin.Context) {
		db, err := sql.Open("sqlite3", "/home/jelly/Documents/tegola_server/hk.gpkg")
		if err != nil {
			panic(err)
		}
		defer db.Close()
		z := c.Param("z")
		x := c.Param("x")
		y := strings.Split(c.Request.URL.Path, "/")[2]

		rows, err := db.Query("SELECT tile_data FROM hk WHERE zoom_level=? and tile_column=? and tile_row=?", 18, 4, 4)
		if err != nil {
			panic(err)
		}
		var (
			data []byte
		)
		defer rows.Close()
		for rows.Next() {
			if err = rows.Scan(&data); err != nil {
				panic(err)
			}
		}
		fmt.Println(z, x, y, len(data))
		c.Header("Content-Type", "image/png")
		c.Data(http.StatusOK, "image/png", data)
	})
	app.Run(":9999")
}
