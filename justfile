test:
    @echo "testing"

build: clean
    @echo "building"
    go build ./models
    @echo "build terminada"