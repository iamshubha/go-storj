package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UploadData struct {
	Message string `json:"message" validate:"required"`
}

func (s *Server) Upload(c *gin.Context) {

	var uploadData UploadData

	if err := c.ShouldBindJSON(&uploadData); err != nil {
		// Return an error if the JSON decoding fails
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	v := validator.New()
	if err := v.Struct(uploadData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := s.Client.Upload(c.Request.Context(), uploadData.Message)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true})
}
