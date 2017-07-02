# Go Dev in Docker Container

## Requirements
- Local Golang setup with $GOPATH defined
- Docker

## Commands
*Run program in container:*
    `docker-compose up`

*Run tests:*
    `docker-compose run app go test --code-coverage -v ./app/...`

## Vendor packages
- Install Govendor locally: `go get -u github.com/kardianos/govendor`
- Make sure you're in the correct project directory (./app)
- First time only: `govendor init`
- Then: 
    - `govendor add +external` to add packages from local $GOPATH
    - and `govendor update +vendor` to update those packages
    - Or `govendor fetch +out` to fetch packages directly from repos

