# account-producer-service

## About
Service responsible for account management, integration with partners, such as Via Cep and with the producer service - the only one that has a direct connection to the database.

## Technologies
* Golang 1.18

## Development requirements
* Docker Compose
* Visual Studio Code

## Directory Structure
- `api`
    - OpenAPI/Swagger specs, JSON schema files, protocol definition files.
- `build`
    - It has all cloud package, container (Docker), operating system (deb, rpm, pkg) and scripts settings.
- `cmd`
    - It has the `main` function that imports and invokes code from the `/internal` and `/pkg` directories.
- `internal`
    - It has all the code that is not available for import.
- `local-dev`
    - Possui toda configuração do docker.

## Running
- `Docker`
    - Run the following command: docker-compose -f local-dev/docker-compose.yaml --profile infra up -d
- `Run the project`
    - Just to create the topics: run -> start debugging -> to allow -> perform account creation via Postman -> stop
- `Run the project`
    - Just to create the topics: run -> start debugging -> to allow -> stop




