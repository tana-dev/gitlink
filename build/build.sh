# set
OS=${1}

# resource file copy
cp -R ../src/resources ../bin/gitlocal/
cp ../src/config/user.json ../bin/gitlocal/config/

# remove binary file
rm ../bin/gitlocal/gitlocal*

# compile
cd ../src/main/
case ${OS} in
    "linux" ) GOOS=linux GOARCH=amd64 go build -o ../../bin/gitlocal/gitlocal main.go
    ;;
    "windows" ) GOOS=windows GOARCH=amd64 go build -o ../../bin/gitlocal/gitlocal.exe main.go
    ;;
    "mac" ) GOOS=darwin GOARCH=amd64 go build -o ../../bin/gitlocal/gitlocal main.go
    ;;
    * ) echo "No compile"
    ;;
esac
