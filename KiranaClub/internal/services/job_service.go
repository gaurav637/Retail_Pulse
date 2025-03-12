package services

import (
	"github.com/kiranaClub/internal/models"
	"github.com/kiranaClub/internal/repository"
	"github.com/kiranaClub/internal/dto"
)

type JobService struct {
	Repo *repository.JobRepository
}

func NewJobService(repo *repository.JobRepository) *JobService {
	return &JobService{Repo: repo}
}

func (s *JobService) InsertJobs(jobs dto.JobRequestDTO) (string, error) {
	return s.Repo.InsertJobs(jobs)
}

func (s *JobService) GetAllJobs() ([]models.Job, error) {
	return s.Repo.GetAllJobs()
}

func (s *JobService) GetJobByID(id string) (*models.Job, error) {
	return s.Repo.GetJobByID(id)
}


func (s *JobService) DeleteJob(id int) error {
	return s.Repo.DeleteJob(id)
}
