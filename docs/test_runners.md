# Test runners

## Criterios de elección
- Debe tener un mantenimiento en la actualidad para no aumentar la deuda técnica; estudiaremos el estado de los repositorios de las cada herramienta para ver su seguimiento y descartaremos aquellos con más de 6 meses de inactividad.
- Debe necesitar el mínimo de paquetes, bibliotecas o cualquier software secundario para su funcionamiento, se descartará aquellos que necesiten de múltiples elementos externos para el testeo de código.
## Opciones disponibles para test runners o frameworks:

### Go testing
**Según la página de documentación oficial de [golang](https://pkg.go.dev/testing) y su [repositorio](https://github.com/golang/go/blob/master/src/testing/testing.go)**

Go testing es una función de un paquete nativo del lenguaje Go para testeo automático de código. Es actualizado regularmente dado que va directamente asociado al lenguaje y no necesita de instalación, para su uso solo hay que importar "testing" en el archivo a usar.

### Testify
**Según el estudio de [google trends](https://trends.google.com/trends/explore?date=today%205-y&q=golang%20testify,golang%20goconvey,golang%20ginkgo,golang%20httpexpect,golang%20gomega) y su [repositorio](https://github.com/stretchr/testify)**

Testify es una serie de paquetes de uso bastante extendido, tiene una instalación sencilla con solo escribir `$go get github.com/stretchr/testify` y comprende 5 paquetes; respecto a su seguimiento es revisado mensualmente para mantener su uso correcto.

### Ginkgo
**Según la [página oficial](https://onsi.github.io/ginkgo/) y su [repositorio](https://github.com/onsi/ginkgo)**

Diseñado exclusivamente para go, es un framework de testeo que emplea test centrados en el uso de DSL(Domain Specific Language) e, para instalarse solo necesita escribirse `$go install github.com/onsi/ginkgo/v2/ginkgo` y según su repositorio fue actualizado hace 2 semanas según se escribe este apartado.

### Maelstrom
**Según una discusión sobre el tema en [reddit](https://www.reddit.com/r/golang/comments/t29c4d/looking_for_a_test_runner_like_pytest/) y su [repositorio](https://github.com/maelstrom-software/maelstrom)**

Herramienta de testeo de forma aislada en contenedores, para instalarse solo necesita escribirse `$wget -q -O - https://github.com/maelstrom-software/maelstrom/releases/latest/download/cargo-maelstrom-x86_64-unknown-linux-gnu.tgz | tar xzf -` que da lugar a múltiples paquetes para emplear este software. Sin embargo su repositorio es actualizado regularmente, hace escasas horas desde la redacción de este párrafo que se actualizó.


**Vamos a emplear go testing como test runner.**

Go testing es la herramienta más directa y es nativa del lenguaje Go, permitiendo instalación mínima para el testeo automático de los archivos del proyecto.