set -euxo pipefail

mkdir -p "$(pwd)/functions"
GOBIN=$(pwd)/functions 
go install ./...
str="$(pwd)"/functions/*
find "$(dirname "$str")" -name "$(basename "$str")" 2>/dev/null | while read file
do
  chmod +x "$file"
done
go env