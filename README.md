
# Image Processing Service

## Overview

This service processes thousands of images collected from stores. It allows users to submit jobs containing images to be processed, calculates the perimeter of each image, and provides job status updates. The job details and status can be accessed via two primary APIs.

---

## Assumptions

- The service is expected to run locally or in a Docker container.
- The system can handle multiple jobs and process images concurrently.
- Each job will contain multiple images, and each image's perimeter is calculated.
- The job and image processing is done asynchronously.
- We simulate GPU processing with a random delay of 0.1 to 0.4 seconds after each image's perimeter is calculated.

---

## Prerequisites

Before running the application, make sure you have the following installed on your local machine:

1. **Go** (1.20 or higher)
   - You can download Go from [here](https://golang.org/dl/).
2. **Git** (for version control)
   - You can download Git from [here](https://git-scm.com/downloads).
3. **Docker** (optional, for running the app in a container)
   - You can download Docker from [here](https://www.docker.com/get-started).

---

## Installation

### Without Docker

1. Clone the repository:

   ```bash
   git clone https://github.com/Margesh06/Retail-pulse.git
   ```

2. Download dependencies:

   ```bash
   go mod tidy
   ```

3. Run the application:

   ```bash
   go run main.go
   ```

4. The application will run on `http://localhost:8080`.

### With Docker

1. Clone the repository:

   ```bash
   git clone <your-repo-url>
   cd retail-pulse
   ```

2. Build the Docker image:

   ```bash
   docker build -t retail-pulse .
   ```

3. Run the Docker container:

   ```bash
   docker run -p 8081:8080 retail-pulse
   ```

4. The application will run on `http://localhost:8080` within the Docker container.

---

## API Endpoints

### 1. Submit Job

**URL**: `/api/submit/`  
**Method**: `POST`  
**Request Payload**:

```json
{
   "count": 2,
   "visits": [
      {
         "store_id": "S00339218",
         "image_url": [
            "https://www.gstatic.com/webp/gallery/2.jpg",
            "https://www.gstatic.com/webp/gallery/3.jpg"
         ],
         "visit_time": "time of store visit"
      },
      {
         "store_id": "S01408764",
         "image_url": [
            "https://www.gstatic.com/webp/gallery/3.jpg"
         ],
         "visit_time": "time of store visit"
      }
   ]
}
```

**Success Response**:
- **Code**: `201 Created`
- **Content**:

```json
{
    "job_id": "job-80375"
}
```

**Error Response**:
- **Code**: `400 Bad Request`
- **Content**:

```json
{
    "error": "Invalid request payload"
}
```

---

### 2. Get Job Info

**URL**: `/api/status?jobid=<job_id>`  
**Method**: `GET`  
**URL Parameter**: `jobid` (Job ID received while creating the job)

**Success Response**:
- **Code**: `200 OK`
- **Content**:

```json
{
    "status": "completed",
    "job_id": "job-80375"
}
```

**Job Status - Failed**:
If a store_id does not exist or an image download fails for any given URL, the error message contains only the failed store_id.

- **Code**: `200 OK`
- **Content**:

```json
{
    "status": "failed",
    "job_id": "job-80375",
    "error": [{
        "store_id": "S00339218",
        "error": "Image download failed"
    }]
}
```

**Error Response**:
- **Code**: `400 Bad Request`
- **Content**:

```json
{
    "error": "Job ID not found"
}
```

---

## Example Test Cases in Postman

### Submit Job

In the **Tests** tab of the **Submit Job** request, add the following tests:

```javascript
pm.test("Status code is 201", function () {
    pm.response.to.have.status(201);
});

pm.test("Job ID is returned", function () {
    var jsonData = pm.response.json();
    pm.expect(jsonData.job_id).to.exist;
});

```

### Get Job Info

In the **Tests** tab of the **Get Job Info** request, add the following tests:

```javascript
pm.test("Status code is 200", function () {
    pm.response.to.have.status(200);
});

pm.test("Job status is completed or failed", function () {
    var jsonData = pm.response.json();
    pm.expect(jsonData.status).to.be.oneOf(["completed", "failed"]);
});

```

---
