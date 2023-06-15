package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goweb/send"
  "goweb/compute"
	"net/http"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob(
		"views/HTML/*",
	)
	r.StaticFS("/views",http.Dir("./views"))
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "Login.html", nil)
	})
 	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://127.0.0.1:7995/login")
	})
	r.POST("/login", func(c *gin.Context) {
		id := c.PostForm("idnumber")
		username := c.PostForm("username")
		password := c.PostForm("password")
		user := &send.User{
			Id:       id,
			Name:     username,
			Password: password,
		}
		flag := send.Ask(id, username, password)
		if flag == 1 {
			key, _ := send.GenerateToken(*user)
			fmt.Println(key)
			//c.JSON(http.StatusOK, gin.H{"token": key})
			c.Redirect(http.StatusMovedPermanently, "http://127.0.0.1:7995/map")
		} else {
			//send.Insert(id,username,password)
			//c.JSON(http.StatusOK, gin.H{"token": nil})
			c.Redirect(http.StatusMovedPermanently, "http://127.0.0.1:7995/login")
		}
	})
	r.GET("/map", func(c *gin.Context) {
		c.HTML(http.StatusOK, "map.html", nil)
	})
	r.POST("/map", func(c *gin.Context) {
		var now send.Date
		err := c.BindJSON(&now)
		fmt.Println(err)
		fmt.Println(now)
		fmt.Println(now.Etime)
		fmt.Println(len(now.Node))
	  ca := compute.Count(now)
		//c.HTML(http.StatusOK, "state_red.html", nil)
		if ca=="A"{
			c.JSON(302, gin.H{"location": "http://127.0.0.1:7995/sate1"})
		}
		if ca=="B"{
			c.JSON(302, gin.H{"location": "http://127.0.0.1:7995/sate2"})
		}
		if ca=="C"{
			c.JSON(302, gin.H{"location": "http://127.0.0.1:7995/sate3"})
		}
		//c.JSON(302, gin.H{"location": "http://1.15.146.175:7995/sate1"})
		//c.Redirect(302, "http://10.0.4.15:7995/sate1")
		//if err != nil {
		//	//log.Info(err)
		//	c.JSON(200, gin.H{"errcode": 400, "description": "Post Data Err"})
		//	return
		//} else {
		//	//fmt.Println(reqInfo.Data)
		//}
	})
	r.GET("/sate1", func(c *gin.Context) {
		//c.JSON(200, gin.H{"data": "1"})
		c.HTML(http.StatusOK, "state_red.html", nil)

	})
	r.POST("/sate1", func(c *gin.Context) {
		//c.JSON(200, gin.H{"data": "1"})
		c.HTML(http.StatusOK, "state_red.html", nil)

	})
	r.GET("/sate2", func(c *gin.Context) {
		//c.JSON(200, gin.H{"data": "1"})
		c.HTML(http.StatusOK, "state_yellow.html", nil)

	})
	r.GET("/sate3", func(c *gin.Context) {
		//c.JSON(200, gin.H{"data": "1"})
		c.HTML(http.StatusOK, "state_green.html", nil)

	})

	//r.POST("/register", func(c *gin.Context) {
	//	id := c.PostForm("id")
	//	username := c.PostForm("username")
	//	password := c.PostForm("password")
	//	mark := send.Insert(id, username, password)
	//	if mark == 1 {
	//		c.Redirect(http.StatusMovedPermanently, "http://127.0.0.1:8080/login")
	//		//c.Request.URL.Path = "/login"
	//		//r.HandleContext(c)
	//	}
	//})
	r.Run("127.0.0.1:7995")
}
