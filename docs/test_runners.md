# Test runners

## Criterios de elección
- Debe tener un mantenimiento en la actualidad para no aumentar la deuda técnica; estudiaremos el estado de los repositorios de las cada herramienta para ver su seguimiento y descartaremos aquellos con más de 6 meses de inactividad.
- Debe necesitar el mínimo de paquetes, bibliotecas o cualquier software secundario para su funcionamiento, se descartará aquellos que necesiten de múltiples elementos externos para el testeo de código.
## Opciones disponibles para test runners o frameworks:

### Go testing
**Según la página de documentación oficial de [golang](https://pkg.go.dev/testing) y su [repositorio](https://github.com/golang/go/blob/master/src/testing/testing.go)**

Go testing es una función de un paquete nativo del lenguaje Go para testeo automático de código. Es actualizado regularmente dado que va directamente asociado al lenguaje y no necesita de instalación, para su uso solo hay que importar "testing" en el archivo a usar.

