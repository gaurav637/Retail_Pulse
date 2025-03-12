package repository

import (
	"fmt"
	"image"
	_ "image/jpeg" // JPEG decoding photos
	_ "image/png"  // PNG decoding photos
	"log"
	"math/rand"
	"net/http"
	"time"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/kiranaClub/internal/dto"
	"github.com/kiranaClub/internal/enums"
	"github.com/kiranaClub/internal/models"
)

type JobRepository struct {
	DB *sqlx.DB
}

func NewJobRepository(db *sqlx.DB) *JobRepository {
	return &JobRepository{DB: db}
}

// this logic is used insert the job in your table and starting async processing
func (r *JobRepository) InsertJobs(request dto.JobRequestDTO) (string, error) {
	log.Printf(`request30  -> `, request)
	// we can generate uuid and used as a jobID
	jobID := uuid.New().String()
	// this qeuery insert the job entry into job table
	query := `INSERT INTO jobs (ID, created_at, updated_at, processing_status) VALUES (?, ?, ?, ?)`
	_, err := r.DB.Exec(query, jobID, time.Now(), time.Now(), enums.JobStatusOngoing)
	if err != nil {
		return "", err
	}

	// starting asyncprocessing here...
	go r.processImages(jobID, request)
	// return the jobId
	return jobID, nil
}

// this logic to used process the image and perform the operation in your image-> downloads and processes each image asynchronously
func (r *JobRepository) processImages(jobID string, request dto.JobRequestDTO) {
	for _, visit := range request.Visits { // Iterate over visits
		storeID := visit.StoreID
		for _, imageURL := range visit.ImageURLs { // Iterate over images
			go func(url, storeID string) {
				// Simulate image processing delay (random sleep time 0.1 - 0.4 sec)
				sleepTime := time.Duration(100+rand.Intn(300)) * time.Millisecond
				time.Sleep(sleepTime)
				// Download the image and get its dimensions
				width, height, err := DownloadImage(url)
				if err != nil {
					log.Printf("Failed to download image: %s, error: %v\n", url, err)
					return
				}

				// Calculate the perimeter
				perimeter := 2 * (width + height)
				// Store the image details in the database
				query := `INSERT INTO images (job_id, store_id, image_url, width, height, parameter, VisitTime) 
						  VALUES (?, ?, ?, ?, ?, ?, ?)`
				_, err = r.DB.Exec(query, jobID, storeID, url, width, height, perimeter, time.Now())
				if err != nil {
					log.Printf("Error saving image result: %v\n", err)
				}
			}(imageURL, storeID) // Pass variables explicitly to goroutine
		}
	}

	// Wait for some time to allow Goroutines to finish before updating job status
	time.Sleep(1 * time.Second)

	// Once all images are processed, update job status
	query := `UPDATE jobs SET processing_status = ? WHERE id = ?`
	_, err := r.DB.Exec(query, enums.JobStatusCompleted, jobID)
	if err != nil {
		_, err := r.DB.Exec(query, enums.JobStatusFailed, jobID)
		log.Printf("Error updating job status: %v\n", err)
	}
}

// downloadImage this logic is used to download iamge
func DownloadImage(imageURL string) (int, int, error) {
	// Send an HTTP GET request to fetch the image
	resp, err := http.Get(imageURL)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to download image: %v", err)
	}
	defer resp.Body.Close()

	// Decode the image
	img, _, err := image.Decode(resp.Body)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to decode image: %v", err)
	}

	// Get the dimensions of the image
	width := img.Bounds().Dx()
	height := img.Bounds().Dy()

	return width, height, nil
}

// Get all jobs
func (r *JobRepository) GetAllJobs() ([]models.Job, error) {
	var jobs []models.Job
	query := "SELECT * FROM jobs"
	err := r.DB.Select(&jobs, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get jobs: %v", err)
	}
	return jobs, nil
}

// Get a job by ID
func (r *JobRepository) GetJobByID(id string) (*models.Job, error) {
	var job models.Job
	query := "SELECT * FROM jobs WHERE id = ?"
	err := r.DB.Get(&job, query, id)
	log.Printf("err job 128 -> ", err);
	if err != nil {
		return nil, fmt.Errorf("failed to get job: %v", err)
	}
	return &job, nil
}

// Delete a job
func (r *JobRepository) DeleteJob(id int) error {
	query := "DELETE FROM jobs WHERE id=?"
	_, err := r.DB.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete job: %v", err)
	}
	return nil
}
