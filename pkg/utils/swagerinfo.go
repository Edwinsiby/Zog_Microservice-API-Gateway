package utils

import "gateway/cmd/docs"

func NewSwaggerInfo() {
	info, _ := LoadConfig("./")
	docs.SwaggerInfo.Title = info.INFO.Title
	docs.SwaggerInfo.Description = info.INFO.Description
	docs.SwaggerInfo.Version = info.INFO.Version
	docs.SwaggerInfo.InfoInstanceName = "swagger"
	docs.SwaggerInfo.Host = info.INFO.Host
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
}
