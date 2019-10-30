package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/titasp/gin-starter-project/models"
	"net/http"
)

type UserCertificatesController struct{}

var certificateModel = new(models.Certificate)

func (u UserCertificatesController) One(c *gin.Context) {
	cert, err := certificateModel.GetByID(c.Param("userid"), c.Param("orgid"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	if cert == nil {
		c.AbortWithStatus(http.StatusNoContent)
		return
	}

	c.JSON(http.StatusOK, cert)
}

func (u UserCertificatesController) All(c *gin.Context) {

}

func (u UserCertificatesController) Create(c *gin.Context) {

}

func (u UserCertificatesController) Update(c *gin.Context) {

}

func (u UserCertificatesController) Delete(c *gin.Context) {

}
