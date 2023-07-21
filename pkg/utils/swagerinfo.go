package utils

import (
	"gateway/cmd/docs"
)

func NewSwaggerInfo(title, description, version, host string) {

	docs.SwaggerInfo.Title = title
	docs.SwaggerInfo.Description = description
	docs.SwaggerInfo.Version = version
	docs.SwaggerInfo.InfoInstanceName = "swagger"
	docs.SwaggerInfo.Host = host
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http"}

}
