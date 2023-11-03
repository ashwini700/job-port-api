package handlers

import (
	"encoding/json"
	"job-port-api/internal/auth"
	"job-port-api/internal/middleware"
	"job-port-api/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
)

func (h *handler) ViewJobById(c *gin.Context) {
	ctx := c.Request.Context()
	traceid, ok := ctx.Value(middleware.TraceIdKey).(string)
	if !ok {
		log.Error().Msg("traceid missing from context")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": http.StatusText(http.StatusInternalServerError),
		})
		return
	}
	_, ok = ctx.Value(auth.Ctxkey).(jwt.RegisteredClaims)
	if !ok {
		log.Error().Str("Trace Id", traceid).Msg("login required")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": http.StatusText(http.StatusUnauthorized)})
		return
	}
	id := c.Param("id")

	cid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	}

	jobData, err := h.service.ViewJobDetailsById(ctx, cid)
	if err != nil {
		log.Error().Err(err).Str("trace id", traceid)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, jobData)

}

func (h *handler) ViewAllJobs(c *gin.Context) {
	ctx := c.Request.Context()
	traceid, ok := ctx.Value(middleware.TraceIdKey).(string)
	if !ok {
		log.Error().Msg("traceid missing from context")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": http.StatusText(http.StatusInternalServerError),
		})
		return
	}
	_, ok = ctx.Value(auth.Ctxkey).(jwt.RegisteredClaims)
	if !ok {
		log.Error().Str("Trace Id", traceid).Msg("login required")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": http.StatusText(http.StatusUnauthorized)})
		return
	}

	jobDetails, err := h.service.ViewAllJobPostings(ctx)
	if err != nil {
		log.Error().Err(err).Str("trace id", traceid)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, jobDetails)
}

func (h *handler) ViewJobByCompanyId(c *gin.Context) {
	ctx := c.Request.Context()
	traceid, ok := ctx.Value(middleware.TraceIdKey).(string)
	if !ok {
		log.Error().Msg("traceid missing from context")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": http.StatusText(http.StatusInternalServerError),
		})
		return
	}
	_, ok = ctx.Value(auth.Ctxkey).(jwt.RegisteredClaims)
	if !ok {
		log.Error().Str("Trace Id", traceid).Msg("login required")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": http.StatusText(http.StatusUnauthorized)})
		return
	}
	id := c.Param("cid")

	cid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	}

	jobData, err := h.service.ViewJobDetails(ctx, cid)
	if err != nil {
		log.Error().Err(err).Str("trace id", traceid)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, jobData)

}

func (h *handler) AddJob(c *gin.Context) {
	ctx := c.Request.Context()
	traceid, ok := ctx.Value(middleware.TraceIdKey).(string)
	if !ok {
		log.Error().Msg("traceid missing from context")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": http.StatusText(http.StatusInternalServerError),
		})
		return
	}
	_, ok = ctx.Value(auth.Ctxkey).(jwt.RegisteredClaims)
	if !ok {
		log.Error().Str("Trace Id", traceid).Msg("login required")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": http.StatusText(http.StatusUnauthorized)})
		return
	}
	var jobData models.Job

	err := json.NewDecoder(c.Request.Body).Decode(&jobData)
	if err != nil {
		log.Error().Err(err).Str("trace id", traceid)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "provide valid id ,jobrole and salary",
		})
		return
	}
	jobData, err = h.service.AddJobDetails(ctx, jobData)
	if err != nil {
		log.Error().Err(err).Str("trace id", traceid)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, jobData)
}
