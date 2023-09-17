
# let initialise our project
git init
# let add our remote git url e.g. https://biacibenga.com/lib/go/bialib.git
git add origin https://biacibenga.com/lib/go/bialib.git
# let now add all our file to the repository
git add .
# let commit our code
git commit -m "Initial"

# =============let call our module local ================================
go mod edit -replace biacibenga.com/lib/go/bialib=../bialib/dbcassandra

# let execute our module tidy command
go mod tidy


