# KiranaClub Task - Retail_Pulse

## Description
This project implements a RESTful API for asynchronous image processing. The API allows clients to submit jobs consisting of multiple store visits. Each visit includes a store ID and a list of image URLs. The API processes these jobs by downloading the images, calculating their perimeters, and returning the results.

## Setting up the project Locally
To set up the project paste the follwing commands in your terminal:

```
    git clone https://github.com/gaurav637/Retail_Pulse
    cd KiranaClub
    go mod tidy
````
This will install all the required dependencies for the project.

## Prerequisites
Ensure the following are installed on your system:

- Go (go1.20.1+)
- Docker
  
## API Endpoints 

  1. Submit Job
     - URL: POST ```/api/submit/```
     - Request Body:
```       
       
       {
          "count": 2,
          "visits": [
                {
                  "store_id": "S00339218",
                  "image_url": [
                    "https://www.gstatic.com/webp/gallery/2.jpg",
                    "https://www.gstatic.com/webp/gallery/3.jpg"
                         ],
                  "visit_time": "2025-03-11"
                }
           ]
       }
````
-  Response: 201 Created: 
```
job id 
````
Get Job Status
URL: GET /api/status/job?id=

## Development Tools
- IDEs: VS Code, GoLand, IntelliJ
- macOS, Windows  
- Postman
- Draw.io
- Docker Desktop (for logs)


## Future Improvements:- 

- Improve Documentation: Enhance the README with more details, examples, and diagrams.
- Add CI/CD Pipeline: Automate testing and deployment using GitHub Actions or Jenkins.
- Improve Code Quality: Refactor code for better maintainability and readability.
- Achieve 100% API Testing: Ensure complete test coverage using tools like Jest or Postman.
- Optimize API Latency: Improve response times using caching and database indexing.
- Enhance Security: Implement proper authentication, authorization, and data validation.
- Implement Rate Limiting: Prevent abuse by setting request limits per user.
- Improve Error Handling: Implement proper error messages and structured logging.
- Add Role-Based Access Control: Implement user roles and permissions for data access.
