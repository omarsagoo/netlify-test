set -euxo pipefail

mkdir -p "$(pwd)/functions"
GOBIN=$(pwd)/functions 
go install ./...
str="$(pwd)"/functions/*
chmod +x "$str"
go env