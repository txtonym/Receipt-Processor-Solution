# Receipt Processor

A RESTful API service built in Go to process and score retail receipts based on specific rules. This project demonstrates fundamental concepts of REST API development, including request handling, data processing, and in-memory storage.

## Table of Contents
- Project Overview
- Features
- Getting Started
  - Prerequisites
  - Installation
- Usage
  - Endpoints
  - Examples
- Scoring Rules
- Docker Setup
- Testing
- License

## Project Overview
The Receipt Processor API provides an interface for submitting retail receipts, calculating points based on defined criteria, and retrieving the score of each receipt. Each receipt is given a unique ID upon submission, which can later be used to retrieve its associated score.

## Features
- Process Receipts: Accepts receipt data and assigns a unique ID.
- Calculate Points: Scores the receipt based on criteria like retailer name, purchase total, item count, and purchase time.
- Retrieve Points: Allows querying the points associated with a receipt ID.

## Getting Started

### Prerequisites
- Go (version 1.18 or later)
- Git (for cloning the repository)
- Docker (optional, for running in a container)

### Installation
1. Clone the Repository: `git clone https://github.com/your-username/receipt-processor.git` then `cd receipt-processor`
2. Install Dependencies: `go mod download`
3. Run the Server: `go run .` The server will start on http://localhost:8080.

## Usage

### Endpoints
1. Process Receipt
   - Endpoint: POST /receipts/process
   - Description: Accepts a receipt JSON object, calculates points, and returns a unique receipt ID.
   - Request Body: JSON object representing the receipt.
2. Get Points
   - Endpoint: GET /receipts/{id}/points
   - Description: Returns the points for a receipt with the specified ID.
   - Path Parameter: id (string): Unique identifier for the receipt.

### Examples

#### Process a Receipt
Request: `curl -X POST http://localhost:8080/receipts/process -H "Content-Type: application/json" -d '{ "retailer": "Target", "purchaseDate": "2022-01-01", "purchaseTime": "13:01", "items": [ { "shortDescription": "Mountain Dew 12PK", "price": "6.49" }, { "shortDescription": "Emils Cheese Pizza", "price": "12.25" } ], "total": "35.35" }'`

Response: `{ "id": "7fb1377b-b223-49d9-a31a-5a02701dd310" }`

#### Get Points for a Receipt
Request: `curl -X GET http://localhost:8080/receipts/7fb1377b-b223-49d9-a31a-5a02701dd310/points`

Response: `{ "points": 28 }`

## Scoring Rules
The following rules are used to calculate points for each receipt:
1. Retailer Name: 1 point for every alphanumeric character in the retailer's name.
2. Round Dollar Total: 50 points if the total is a round dollar amount with no cents.
3. Multiple of 0.25: 25 points if the total is a multiple of 0.25.
4. Item Count: 5 points for every two items on the receipt.
5. Item Description Length: For items with a description length that is a multiple of 3, multiply the item's price by 0.2 and round up. This result is the number of points earned.
6. Odd Purchase Day: 6 points if the day in the purchase date is odd.
7. Purchase Time: 10 points if the purchase time is between 2:00 pm and 4:00 pm.

## Docker Setup
To make it easy to run the Receipt Processor in a consistent environment, you can use Docker.

### Steps
1. Build the Docker Image: In the root directory of the project, run `docker build -t receipt-processor .`
2. Run the Docker Container: Run `docker run -p 8080:8080 receipt-processor` The application will be available at http://localhost:8080.

## Testing
You can test the API locally using curl or a tool like Postman.

### Example Commands
1. Process a Receipt: `curl -X POST http://localhost:8080/receipts/process -H "Content-Type: application/json" -d '{ "retailer": "Target", "purchaseDate": "2022-01-01", "purchaseTime": "13:01", "items": [ { "shortDescription": "Mountain Dew 12PK", "price": "6.49" }, { "shortDescription": "Emils Cheese Pizza", "price": "12.25" } ], "total": "35.35" }'`
2. Get Points for a Receipt: `curl -X GET http://localhost:8080/receipts/{id}/points` Replace `{id}` with the actual ID returned from the POST request.
