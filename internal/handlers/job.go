package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"

	"job-port-api/internal/auth"
	"job-port-api/internal/middleware"
	"job-port-api/internal/models"
)

func (h *handler) FetchJobById(c *gin.Context) {
	ctx := c.Request.Context()
	traceID, ok := ctx.Value(middleware.TraceIdKey).(string)
	if !ok {
		log.Error().Msg("traceid missing from context")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": http.StatusText(http.StatusInternalServerError),
		})
		return
	}
	_, ok = ctx.Value(auth.Ctxkey).(jwt.RegisteredClaims)
	if !ok {
		log.Error().Str("Trace Id", traceID).Msg("login required")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": http.StatusText(http.StatusUnauthorized)})
		return
	}
	id := c.Param("id")

	cid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	}

	jobData, err := h.service.FetchJobDetailsById(ctx, cid)
	if err != nil {
		log.Error().Err(err).Str("trace id", traceID)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, jobData)

}

func (h *handler) FetchAllJobs(c *gin.Context) {
	ctx := c.Request.Context()
	traceID, ok := ctx.Value(middleware.TraceIdKey).(string)
	if !ok {
		log.Error().Msg("traceid missing from context")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": http.StatusText(http.StatusInternalServerError),
		})
		return
	}
	_, ok = ctx.Value(auth.Ctxkey).(jwt.RegisteredClaims)
	if !ok {
		log.Error().Str("Trace Id", traceID).Msg("login required")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": http.StatusText(http.StatusUnauthorized)})
		return
	}

	jobDetails, err := h.service.FetchJobPosts(ctx)
	if err != nil {
		log.Error().Err(err).Str("trace id", traceID)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, jobDetails)
}

func (h *handler) FetchJobByCompanyId(c *gin.Context) {
	ctx := c.Request.Context()
	traceID, traceIDExist := ctx.Value(middleware.TraceIdKey).(string)
	if !traceIDExist {
		log.Error().Msg("traceid missing from context")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": http.StatusText(http.StatusInternalServerError),
		})
		return
	}
	_, traceIDExist = ctx.Value(auth.Ctxkey).(jwt.RegisteredClaims)
	if !traceIDExist {
		log.Error().Str("Trace Id", traceID).Msg("login required")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": http.StatusText(http.StatusUnauthorized)})
		return
	}
	id := c.Param("cid")

	cid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	}

	jobData, err := h.service.FetchJobDetails(ctx, cid)
	if err != nil {
		log.Error().Err(err).Str("trace id", traceID)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, jobData)

}

func (h *handler) AddJob(c *gin.Context) {
	ctx := c.Request.Context()
	traceID, traceIDExist := ctx.Value(middleware.TraceIdKey).(string)
	if !traceIDExist {
		log.Error().Msg("traceid missing from context")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": http.StatusText(http.StatusInternalServerError),
		})
		return
	}
	_, traceIDExist = ctx.Value(auth.Ctxkey).(jwt.RegisteredClaims)
	if !traceIDExist {
		log.Error().Str("Trace Id", traceID).Msg("login required")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": http.StatusText(http.StatusUnauthorized)})
		return
	}
	var jobData models.Job

	err := json.NewDecoder(c.Request.Body).Decode(&jobData)
	if err != nil {
		log.Error().Err(err).Str("trace id", traceID)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "provide valid id ,jobrole and salary",
		})
		return
	}
	jobData, err = h.service.AddJobDetails(ctx, jobData)
	if err != nil {
		log.Error().Err(err).Str("trace id", traceID)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, jobData)
}
