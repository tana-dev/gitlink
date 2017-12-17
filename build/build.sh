# resource file copy
cp -R ../src/resources ../bin/gitlocal/
cp ../src/config/user.json ../bin/gitlocal/config/

# compile
cd ../src/main/
GOOS=linux GOARCH=amd64 go build -o ../../bin/gitlocal/gitlocal main.go

