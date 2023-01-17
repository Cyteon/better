## Getting Started

### Prerequisites

- GoLang 1.19 installed
- A mongodb database (local or cloud hosted)

### Installing

0. Install extra packages:
    ```go install github.com/cosmtrek/air@latest```
    ```go install github.com/swaggo/swag/cmd/swag@latest```
1. Clone the repo
2. Create your own .env file
3. ```air```
4. View docs at http://localhost:8080/docs

### Scripts

- ```air``` - Runs the server in development mode
- ```swag init``` - Generates the swagger docs