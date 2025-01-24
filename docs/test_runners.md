# Test runners

## Criterios de elección
- Debe tener un mantenimiento en la actualidad para no aumentar la deuda técnica.
- Debe necesitar el mínimo de paquetes, bibliotecas o cualquier software secundario para su funcionamiento, se descartará aquellos que necesiten de múltiples elementos externos para el testeo de código.
## Opciones disponibles para test runners o frameworks:

### Go testing
**Según la página de documentación oficial de [golang](https://pkg.go.dev/testing) y su [repositorio](https://github.com/golang/go/blob/master/src/testing/testing.go)**

Go testing es una función de un paquete nativo del lenguaje Go para testeo automático de código. Es actualizado regularmente dado que va directamente asociado al lenguaje y no necesita de instalación ni implementación de ningún paquete, para su uso solo hay que importar "testing" en el archivo a usar y emplea la sintaxis inherente de lenguaje Go.

### Testify
**Según el estudio de [google trends](https://trends.google.com/trends/explore?date=today%205-y&q=golang%20testify,golang%20goconvey,golang%20ginkgo,golang%20httpexpect,golang%20gomega) y su [repositorio](https://github.com/stretchr/testify)**

Testify es una serie de paquetes de uso bastante extendido, comprende 5 paquetes siendo interesante su paquete para implementar mockups de los datos del proyecto, muy popular en desarrollo de proyectos grandes; sin embargo requieren de (estudiar)[https://pkg.go.dev/github.com/stretchr/testify/mock] su sintaxis y correcto funcionamiento para un buen desarrollo de estos test por mockup. Respecto a su seguimiento fue actualizado hace un mes a día de escritura de este apartado, es revisado mensualmente para mantener su uso correcto.

### Ginkgo
**Según la [página oficial](https://onsi.github.io/ginkgo/) y su [repositorio](https://github.com/onsi/ginkgo)**

Diseñado exclusivamente para go, es un framework de testeo que emplea test centrados en el uso de DSL(Domain Specific Language),
según su repositorio fue actualizado hace 2 semanas según se escribe este apartado.

### Maelstrom
**Según una discusión sobre el tema en [reddit](https://www.reddit.com/r/golang/comments/t29c4d/looking_for_a_test_runner_like_pytest/) y su [repositorio](https://github.com/maelstrom-software/maelstrom)**

Conjunto de herramientas de ejecución de test de forma aislada en 'micro-contenedores', explicado su proceso interno en su (página web)[https://maelstrom-software.com/], requiere de la instalación del nodo de Maelstrom que contiene varios paquetes. Sin embargo su repositorio oficial es actualizado regularmente, desde la última revisión de este apartado se actualizó hace 3 días.


## Conclusión final

Go testing al ser una función que es parte del propio lenguaje es actualizada a la par que el lenguaje asegurando que no quede desfasada, es la herramienta más directa y no requiere dependencias externas; permitiendo instalación mínima para el testeo automático de los archivos del proyecto.

**Vamos a emplear go testing como test runner.**
