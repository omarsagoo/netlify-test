mkdir -p "$(pwd)/functions"
GOBIN=$(pwd)/functions go install ./
go env