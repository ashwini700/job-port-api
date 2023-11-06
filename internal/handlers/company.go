package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"

	"job-port-api/internal/auth"
	"job-port-api/internal/middleware"
	"job-port-api/internal/models"

)

// AddCompany handles the addition of a company
func (h *handler) AddCompany(c *gin.Context) {
	ctx := c.Request.Context()
	traceID, traceIDExists := ctx.Value(middleware.TraceIdKey).(string)
	if !traceIDExists {
		log.Error().Msg("traceid missing from context")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": http.StatusText(http.StatusInternalServerError),
		})
		return
	}
	_, traceIDExists = ctx.Value(auth.Ctxkey).(jwt.RegisteredClaims)
	if !traceIDExists {
		log.Error().Str("Trace Id", traceID).Msg("login first")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": http.StatusText(http.StatusUnauthorized)})
		return
	}
	var companyData models.Company

	err := json.NewDecoder(c.Request.Body).Decode(&companyData)
	if err != nil {
		log.Error().Err(err).Str("trace id", traceID)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "please provide valid name and location",
		})
		return
	}
	validate := validator.New()
	err = validate.Struct(companyData)
	if err != nil {
		log.Error().Err(err).Str("trace id", traceID)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "please provide valid name and location",
		})
		return
	}

	companyData, err = h.service.AddCompanyDetails(ctx, companyData)
	if err != nil {
		log.Error().Err(err).Str("trace id", traceID)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, companyData)

}

// ViewCompany handles viewing details of a specific company.
func (h *handler) FetchCompany(c *gin.Context) {
	ctx := c.Request.Context()
	traceID, traceIDExists := ctx.Value(middleware.TraceIdKey).(string)
	if !traceIDExists {
		log.Error().Msg("traceid missing from context")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": http.StatusText(http.StatusInternalServerError),
		})
		return
	}
	_, traceIDExists = ctx.Value(auth.Ctxkey).(jwt.RegisteredClaims)
	if !traceIDExists {
		log.Error().Str("Trace Id", traceID).Msg("login first")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": http.StatusText(http.StatusUnauthorized)})
		return
	}
	id := c.Param("id")

	cid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	companyData, err := h.service.FetchCompanyDetails(ctx, cid)
	if err != nil {
		log.Error().Err(err).Str("trace id", traceID)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, companyData)
}

// ViewAllCompanies lists all available companies.
func (h *handler) FetchAllCompanies(c *gin.Context) {
	ctx := c.Request.Context()
	traceID, traceIDExists := ctx.Value(middleware.TraceIdKey).(string)
	if !traceIDExists {
		log.Error().Msg("traceid missing from context")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": http.StatusText(http.StatusInternalServerError),
		})
		return
	}
	_, traceIDExists = ctx.Value(auth.Ctxkey).(jwt.RegisteredClaims)
	if !traceIDExists {
		log.Error().Str("Trace Id", traceID).Msg("login first")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": http.StatusText(http.StatusUnauthorized)})
		return
	}

	companyDetails, err := h.service.FetchAllCompanies(ctx)
	if err != nil {
		log.Error().Err(err).Str("trace id", traceID)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, companyDetails)
}
