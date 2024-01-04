BINARY_NAME=nba

build:
	GOARCH=amd64 GOOS=darwin go build -o ${BINARY_NAME}-darwin main.go
	GOARCH=amd64 GOOS=linux go build -o ${BINARY_NAME}-linux main.go
	GOARCH=amd64 GOOS=windows go build -o ${BINARY_NAME}-windows main.go

run: 
	build ./${BINARY_NAME}

clean:
	go clean
	rm ${BINARY_NAME}-darwin
	rm ${BINARY_NAME}-linux
	rm ${BINARY_NAME}-windows

test:
	go test ./...

test_coverage:
	go test ./... -coverprofile=coverage.out

dep:
	go mod download

vet:
	go vet

lint:
	golangci-lint run --enable-all

mock:
	mockery --dir=internal/adapters --name=GameRepositoryInterface --filename=mock_game_repository.go --output=internal/adapters/mocks --outpkg=adapters_mock
	mockery --dir=internal/adapters --name=PlayerRepositoryInterface --filename=mock_player_repository.go --output=internal/adapters/mocks --outpkg=adapters_mock
	mockery --dir=internal/adapters --name=TeamRepositoryInterface --filename=mock_team_repository.go --output=internal/adapters/mocks --outpkg=adapters_mock

db:
	docker compose up

