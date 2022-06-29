set -euxo pipefail

mkdir -p "$(pwd)/functions"
chmod +x "$(pwd)"/functions/*
GOBIN=$(pwd)/functions go install ./
go env