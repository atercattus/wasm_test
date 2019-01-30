GO = go
GOFMT = gofmt

BIN = $(CURDIR)/target/build.wasm
SRC = *.go
FLAGS =

.PHONY: all clean

$(BIN): $(SRC)
	GOOS=js GOARCH=wasm $(GO) build $(FLAGS) -o $(BIN)

all: $(BIN)

clean:
	rm -f $(BIN)
