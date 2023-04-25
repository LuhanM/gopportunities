package handler

import (
	"fmt"
	"net/http"

	schemas "github.com/LuhanM/gopportunities/Schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Delete opening
// @Description Delete a job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param id query string true "Opening identification"
// @Success 200 {object} DeleteOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /opening [delete]
func DeleteOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		logger.Errorf("id to delete not defined")
		sendError(ctx, http.StatusBadRequest,
			errParamIsRequired("id", "queryParameter").Error())
		return
	}

	opening := schemas.Opening{}
	if err := db.First(&opening, id).Error; err != nil {
		logger.Errorf(fmt.Sprintf("opening with id: %s not found", id))
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("opening with id: %s not found", id))
		return
	}

	if err := db.Delete(&opening).Error; err != nil {
		logger.Errorf(fmt.Sprintf("error deleting opening with id: %s ", id))
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("error deleting opening with id: %s ", id))
		return
	}

	sendSuccess(ctx, "delete-opening", opening)
}
