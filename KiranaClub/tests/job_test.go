package tests

import (
	"github.com/kiranaClub/internal/services"
	"testing"
	"github.com/kiranaClub/internal/dto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	
)

// Test job creation
type MockJobRepository struct {
	mock.Mock
}

func (m *MockJobRepository) InsertJobs(job dto.JobRequestDTO) (string, error) {
	args := m.Called(job)
	return args.String(0), args.Error(1)
}

func TestCreateJob(t *testing.T) {
	mockRepo := new(MockJobRepository)
	jobService := services.NewJobService(mockRepo)

	// Correct struct format
	request := dto.JobRequestDTO{
		Count: 3,
		Visits: []dto.visits{
			{
				StoreID:     "S00987654",
				url:    []string{"https://example.com/images/store1_1.jpg", "https://example.com/images/store1_2.jpg"},
				visitTime:   "2025-03-12T10:30:00Z",
				
			},
		},
	}

	expectedJobID := "123e4567-e89b-12d3-a456-426614174000"
	mockRepo.On("InsertJobs", request).Return(expectedJobID, nil)

	jobID, err := jobService.InsertJobs(request)

	assert.NoError(t, err)
	assert.Equal(t, expectedJobID, jobID)
	mockRepo.AssertExpectations(t)
}

