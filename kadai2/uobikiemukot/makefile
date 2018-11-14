GOCMD      = go
GOBUILD    = $(GOCMD) build
GOCLEAN    = $(GOCMD) clean
GOTEST     = $(GOCMD) test
GOGET      = $(GOCMD) get
GOVET      = $(GOCMD) vet
GOGENERATE = $(GOCMD) generate
GOTOOL     = $(GOCMD) tool
GOLINT     = golint
GODOC      = godoc

#BUILDENV = GOOS=linux GOARCH=amd64
BUILDENV =

TESTOPT = -v -parallel 4
LDFLAGS = -ldflags '-w -s'

DST  = convert
SRC  = main.go
TESTFILES = 

all: lint build test doc

lint: $(SRC)
	$(GOVET) ./...
	$(GOLINT) ./...

test: $(SRC)
	$(GOTEST) $(TESTOPT) ./imgconv

$(DST): $(SRC)
	$(BUILDENV) $(GOBUILD) $(LDFLAGS) -o $@

build: $(DST)

doc:
	$(GODOC) -ex ./imgconv > imgconv.doc

clean:
	rm -f $(DST) *.doc
