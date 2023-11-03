package handlers

import (
	"job-port-api/internal/auth"
	"job-port-api/internal/middleware"
	"job-port-api/internal/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func API(a auth.TokenAuth, sc service.UserService) *gin.Engine {
	r := gin.New()
	m, err := middleware.NewMid(a)
	if err != nil {
		log.Panic("middleware not setup")
		return nil
	}
	h, err := NewHandlerFunc(sc)
	if err != nil {
		log.Panic("handler not setup")
		return nil
	}
	r.Use(middleware.Log(), gin.Recovery())

	r.GET("/check", m.Authenticate((check)))
	r.POST("/register", h.SignUp)
	r.POST("/login", h.Login)
	r.POST("/companies/Addcomp", m.Authenticate(h.AddCompany))
	r.GET("/companies/viewCompByid/{id}", m.Authenticate(h.ViewCompany))
	r.GET("/companies/Allcomp", m.Authenticate(h.ViewAllCompanies))
	r.POST("/companies/{cid}/Addjobs", m.Authenticate(h.AddJob))
	r.GET("/companies/{cid}/ViewjobsBycompid", m.Authenticate(h.ViewJobByCompanyId))
	r.GET("/jobs/Alljobs", m.Authenticate(h.ViewAllJobs))
	r.GET("/jobs/viewjobsByid/{id}", m.Authenticate(h.ViewJobById))

	return r
}
func check(c *gin.Context) {
	c.JSON(http.StatusOK, "Msg :ok")

}
