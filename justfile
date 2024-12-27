code_folder := "./models"

test:
    @echo "testing"

clean:
    rm -rf bin/

check:
    @echo "comprobando sintaxis"
    @gofmt -e {{code_folder}}
    @echo "fin de comprobación de sintaxis"

build: 
    @echo "building"
    go build {{code_folder}}
    @echo "build terminada"