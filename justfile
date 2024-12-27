code_folder := "./models"

test:
    @echo "testing"

check:
    @echo "comprobando sintaxis"
    @gofmt -e {{code_folder}}
    @echo "fin de comprobación de sintaxis"

build: 
    clean
    @echo "building"
    go build {{code_folder}}
    @echo "build terminada"