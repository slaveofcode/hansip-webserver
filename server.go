package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

//go:embed all:assets/web/**
var webHtml embed.FS

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")             // working directory
	viper.AddConfigPath("$HOME/.hansip") // hansip app directory
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic(fmt.Errorf("please create the config file [config.yaml]"))
		} else {
			panic(fmt.Errorf("unable to read config file [config.yaml]: %w", err))
		}
	}

	viper.SetDefault("server_api.secure", false)
	viper.SetDefault("server_api.host", "localhost")
	viper.SetDefault("server_api.port", "8080")

	viper.SetDefault("server_web.secure", false)
	viper.SetDefault("server_web.host", "localhost")
	viper.SetDefault("server_web.port", "8181")
}

func main() {
	web, err := fs.Sub(webHtml, "assets/web")
	if err != nil {
		log.Fatal(err.Error())
	}

	gin.SetMode(gin.ReleaseMode)
	webStatic := gin.Default()
	webStatic.Use(gin.Recovery())
	// webStatic.StaticFS("/", http.FS(web))

	fs := http.FS(web)
	fileServer := http.StripPrefix("/", http.FileServer(fs))
	webStatic.Any("/*filepath", func(c *gin.Context) {
		file := c.Param("filepath")

		// Check if file exists and/or if we have permission to access it
		f, err := fs.Open(file)
		if err != nil {
			log.Println("fff", file)
			if file == "/server.json" {
				serveServerConfig(c)
				return
			}

			// Fallback to index
			c.FileFromFS("/", fs)
			return
		}
		f.Close()

		fileServer.ServeHTTP(c.Writer, c.Request)
	})

	serverProtocol := "http://"
	if viper.GetBool("server_web.secure") {
		serverProtocol = "https://"
	}

	serverAddr := serverProtocol + viper.GetString("server_web.host")
	serverPort := viper.GetString("server_web.port")

	if serverPort != "80" {
		serverAddr = fmt.Sprintf("%s:%s", serverAddr, serverPort)
	}

	log.Println("Web Server Started at:", serverAddr)

	webStatic.Run(":" + viper.GetString("server_web.port"))
}

func serveServerConfig(c *gin.Context) {
	protocol := "http://"
	if viper.GetBool("server_api.secure") {
		protocol = "https://"
	}

	host := viper.GetString("server_api.host")
	port := viper.GetString("server_api.port")

	serverBaseURL := fmt.Sprintf(`%s%s`, protocol, host)
	if port != "" {
		serverBaseURL = fmt.Sprintf("%s:%s", serverBaseURL, port)
	}

	c.JSON(200, gin.H{
		"baseURL": serverBaseURL,
	})
}
