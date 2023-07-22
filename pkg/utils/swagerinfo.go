package utils

import (
	"gateway/cmd/docs"
)

func NewSwaggerInfo() {
	info, _ := LoadConfig("./")
	docs.SwaggerInfo.Title = info.Title
	docs.SwaggerInfo.Description = info.Description
	docs.SwaggerInfo.Version = info.Version
	docs.SwaggerInfo.InfoInstanceName = "swagger"
	docs.SwaggerInfo.Host = info.Host
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
}
