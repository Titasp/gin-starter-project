package server

import (
	"github.com/gin-gonic/contrib/ginrus"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/titasp/gin-starter-project/controllers"
	"time"
)

func New() *gin.Engine {
	r := gin.New()
	r.Use(ginrus.Ginrus(logrus.StandardLogger(), time.RFC3339, true))
	r.Use(gin.Recovery())

	certificates := new(controllers.UserCertificatesController)

	r.GET("/user-certificates/:userid", certificates.All)
	r.GET("/user-certificates/:userid/organization/:orgid", certificates.One)
	r.POST("/user-certificates/:userid/organization/:orgid", certificates.Create)
	r.PATCH("/user-certificates/:userid/organization/:orgid", certificates.Update)
	r.DELETE("/user-certificates/:userid/organization/:orgid", certificates.Delete)

	health := new(controllers.HealthController)
	r.GET("/status", health.Status)

	return r
}
