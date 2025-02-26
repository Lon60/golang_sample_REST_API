package demo

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	demoService Service
}

func NewDemoHandler(s Service) *Handler {
	return &Handler{demoService: s}
}

// CreateDemo godoc
// @Summary Create a new Demo
// @Description Create a new demo entry
// @Tags demos
// @Accept json
// @Produce json
// @Param demo body Demo true "Demo to create"
// @Success 201 {object} Demo
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /demos/ [post]
func (h *Handler) CreateDemo(c *gin.Context) {
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

// GetDemo godoc
// @Summary Get a Demo by ID
// @Description Get details of a demo entry by ID
// @Tags demos
// @Produce json
// @Param id path int true "Demo ID"
// @Success 200 {object} Demo
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Router /demos/{id} [get]
func (h *Handler) GetDemo(c *gin.Context) {
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

// GetAllDemos godoc
// @Summary Get all Demos
// @Description Retrieve all demo entries
// @Tags demos
// @Produce json
// @Success 200 {array} Demo
// @Failure 500 {object} gin.H
// @Router /demos/ [get]
func (h *Handler) GetAllDemos(c *gin.Context) {
	demos, err := h.demoService.GetAllDemos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, demos)
}

// UpdateDemo godoc
// @Summary Update an existing Demo
// @Description Update a demo entry by ID
// @Tags demos
// @Accept json
// @Produce json
// @Param id path int true "Demo ID"
// @Param demo body Demo true "Updated demo"
// @Success 200 {object} Demo
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /demos/{id} [put]
func (h *Handler) UpdateDemo(c *gin.Context) {
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

// DeleteDemo godoc
// @Summary Delete a Demo
// @Description Delete a demo entry by ID
// @Tags demos
// @Produce json
// @Param id path int true "Demo ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /demos/{id} [delete]
func (h *Handler) DeleteDemo(c *gin.Context) {
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
