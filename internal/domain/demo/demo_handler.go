package demo

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DemoHandler struct {
	demoService Service
}

func NewDemoHandler(s Service) *DemoHandler {
	return &DemoHandler{demoService: s}
}

func (h *DemoHandler) CreateDemo(c *gin.Context) {
	var demo Demo
	if err := c.ShouldBindJSON(&demo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.demoService.CreateDemo(&demo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, demo)
}

func (h *DemoHandler) GetDemo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	demo, err := h.demoService.GetDemo(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if demo == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Demo not found"})
		return
	}
	c.JSON(http.StatusOK, demo)
}

func (h *DemoHandler) GetAllDemos(c *gin.Context) {
	demos, err := h.demoService.GetAllDemos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, demos)
}

func (h *DemoHandler) UpdateDemo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	var demo Demo
	if err := c.ShouldBindJSON(&demo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	demo.ID = uint(id)
	if err := h.demoService.UpdateDemo(&demo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, demo)
}

func (h *DemoHandler) DeleteDemo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	if err := h.demoService.DeleteDemo(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Deleted successfully"})
}
