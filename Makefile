# This is how we want to name the binary output
OUTPUT=common-srv-atomicid

SHELL=C:/Windows/System32/cmd.exe

# These are the values we want to pass for Version and BuildTime
GITTAG=`git describe --tags`
BUILD_TIME=`date +%FT%T%z`

# Setup the -ldflags option for go build here, interpolate the variable values
# LDFLAGS=-ldflags "-X main.GitTag=${GITTAG} -X main.BuildTime=${BUILD_TIME}"
LDFLAGS=-ldflags "-X main.BuildTime=${BUILD_TIME}"

.PHONY:all clean release
all:clean release

clean:
	rm -f ${OUTPUT}

release:
	go build -o ${OUTPUT} main.go



