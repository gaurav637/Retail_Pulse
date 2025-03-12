# Kirana Club - Retail_Pulse

## Description
This project implements a RESTful API for asynchronous image processing. The API allows clients to submit jobs consisting of multiple store visits. Each visit includes a store ID and a list of image URLs. The API processes these jobs by downloading the images, calculating their perimeters, and returning the results.

## Setting up the project Locally
To set up the project paste the follwing commands in your terminal:

### 1. Clone the repository: 
   ```
    https://github.com/gaurav637/Retail_Pulse
   ```
### 2. Navigate to the project directory:

  ```
    cd KiranaClub
  ```

### 3. Open the project in your IDE: 

  Vs Code (recommended) or IntelliJ IDEA 
       
### 4. Configure the database connection in application.properties:

   MYSQL can be used as the database for this project. The database connection can be configured in the 
   database.go file
       
 ```
    DB_HOST=
    DB_PORT=
    DB_USER=
    DB_NAME=
    DB_PASSWORD=
 ```

### 5. Run the application:

  ``` 
   go run main.go
  ```

### 6. Access the application in your web browser at: 

  ``` 
  http://localhost:3000
 ```


## Features 
  - Job Processing System – The service can handle multiple jobs simultaneously
  - Job Status Tracking – Allows users to check the status of submitted jobs
  - Docker Support – Offers a Dockerized setup for easy deployment
  - Asynchronous Job Processing – Uses Goroutines (go keyword) to handle image processing tasks asynchronously, 
    improving efficiency.
  - Image Perimeter Calculation
  - MySQL Integration – Stores job details, image processing results, and status updates in a MySQL database.

    
## Prerequisites/Requirements
Ensure the following are installed on your system:

- Go (go1.20.1+)
- Docker
- Terminal
- Postman
- Vs Code
- MYSQL

## Skills and Expertise
- Go
- RESTs API
- Database (MYSQL)
- Version Control
- Echo
- Testing and Debugging
  
  
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

## Contributing

Contributions are welcome! Please follow these guidelines:

- Fork the repository
- Create a new branch (`git checkout -b feature`)
- Make changes and commit them (`git commit -am 'Add new feature'`)
- Push to the branch (`git push origin feature`)
- Create a pull request

## Contact
For any questions or feedback, please reach out to : negigaurav637@gmail.com , +91 9149102604

Happy coding!
