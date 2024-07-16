# check if go.mod exists in current directory
if [ ! -f go.mod ]; then
    echo "go.mod not found"
    echo "Please make sure you are in the root ingest directory"
    exit 1
fi

go test ./... -race -covermode=atomic -coverprofile=coverage.out -v
go tool cover -html coverage.out -o coverage.html