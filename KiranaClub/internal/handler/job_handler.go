package handler

import (
	"encoding/json"
	"net/http"
	"github.com/labstack/echo/v4"
	"strconv"
	"github.com/kiranaClub/internal/services"
	"github.com/kiranaClub/internal/dto"
	"log"

)

type JobHandler struct {
	Service *services.JobService
}

func NewJobHandler(service *services.JobService) *JobHandler {
	return &JobHandler{Service: service}
}

func (h *JobHandler) InsertJob(c echo.Context) error {
	var jobDTO dto.JobRequestDTO

	// Bind request body to DTO
	if err := c.Bind(&jobDTO); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid request payload",
			"code":    http.StatusBadRequest,
			"data":    nil,
		})
	}
	// Call service layer to insert job
	insertedJob, err := h.Service.InsertJobs(jobDTO)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to insert job",
			"code":    http.StatusInternalServerError,
			"data":    nil,
		})
	}

	// Return response
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Job inserted successfully",
		"code":    http.StatusCreated,
		"JobId":    insertedJob,
	})
}

func (h *JobHandler) GetJobByID(c echo.Context) error {
	// Extract job ID from query parameter
	id := c.QueryParam("id")
	log.Printf("id - 45 -> %s", id)

	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Missing job ID",
			"code":    http.StatusBadRequest,
			"data":    nil,
		})
	}

	job, err := h.Service.GetJobByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "Job not found",
			"code":    http.StatusNotFound,
			"data":    nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Job retrieved successfully",
		"code":    http.StatusOK,
		"data":    job,
	})
}


func (h *JobHandler) GetAllJobs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	jobs, err := h.Service.GetAllJobs()
	if err != nil {
		response := map[string]interface{}{
			"message": "Failed to retrieve jobs",
			"code":    http.StatusInternalServerError,
			"data":    nil,
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}

	response := map[string]interface{}{
		"message": "Jobs retrieved successfully",
		"code":    http.StatusOK,
		"data":    jobs,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (h *JobHandler) DeleteJob(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Extract job ID from query parameter
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		response := map[string]interface{}{
			"message": "Missing job ID",
			"code":    http.StatusBadRequest,
			"data":    nil,
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		response := map[string]interface{}{
			"message": "Invalid job ID",
			"code":    http.StatusBadRequest,
			"data":    nil,
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	err = h.Service.DeleteJob(id)
	if err != nil {
		response := map[string]interface{}{
			"message": "Failed to delete job",
			"code":    http.StatusInternalServerError,
			"data":    nil,
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}

	response := map[string]interface{}{
		"message": "Job deleted successfully",
		"code":    http.StatusOK,
		"data":    nil,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

