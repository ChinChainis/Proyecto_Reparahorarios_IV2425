code_folder := "./models"
go := env_var_or_default("GO", "go")

test:
    @echo "testing"
    go test -v {{code_folder}}

clean:
    rm -rf bin/

check:
    @echo "comprobando sintaxis"
    @gofmt -e {{code_folder}}
    @echo "fin de comprobación de sintaxis"

version:
    @echo "Go version: $({{go}} version)"
    
build: 
    @echo "building"
    go build {{code_folder}}
    @echo "build terminada"