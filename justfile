test:
    @echo "testing"

check:
    @echo "comprobando sintaxis"
    @gofmt -e ./models
    @echo "fin de comprobación de sintaxis"

build: 
    clean
    @echo "building"
    go build ./models
    @echo "build terminada"