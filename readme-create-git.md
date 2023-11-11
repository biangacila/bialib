
# let initialise our project
git init
# let add our remote git url e.g. https://github/biangacila/bialib.git
git add origin https://github/biangacila/bialib.git
# let now add all our file to the repository
git add .
# let commit our code
git commit -m "Initial"

# =============let call our module local ================================
go mod edit -replace github/biangacila/bialib=../bialib/dbcassandra

# let execute our module tidy command
go mod tidy


