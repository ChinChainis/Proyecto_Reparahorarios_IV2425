# Biblioteca de aserciones

## Criterios de elección
- Debe tener un mantenimiento en la actualidad para no aumentar la deuda técnica; descartaremos aquellos cuyos repositorios no hayan tenido actividad de revisión reciente o cuidado con respecto a la actualización del lenguaje Go u otras dependencias a usar.
- Evitar aquellas que empleen instalaciones de librerías externas, daremos prioridad a aquellas optimizadas para el lenguaje go para evitar depender de paquetes o dependencias adicionales.

## Opciones disponibles para aserciones:

### Paquete Assert de Testify
**Según el [repositorio oficial](https://github.com/stretchr/testify)**

Testify es una serie de paquetes de uso bastante extendido, pero nos centraremos en uno específico para el testeo con aserciones llamado `assert package`. Este paquete es revisado regularmente según el seguimiento de commits de su repositorio y testify es completamente orientado a lenguaje Go, testify/assert asegura que emplea salidas compatible con las funciones de go test lo cual nos interesa si lo tenemos como test runner.

### Ghost
**Según el [repositorio oficial](https://github.com/rliebz/ghost)**

Sigue una orientación al lenguaje natural, tiene múltiples palabras reservadas para intentar mejorar de forma intuitiva qué emplea en cada momento, necesita de incluir las dependencias en el proyecto; tiene un arduo seguimiento, hace 2 semanas fue revisado, así que está en constante actualización en caso de encontrarse algun error. Sin embargo declara en su [página oficial](https://pkg.go.dev/github.com/rliebz/ghost) que está en un estado temprano de desarrollo.

### Assert
**Según el [repositorio oficial](https://github.com/negrel/assert)**

Es un paquete para proyectos en Go basado en el paquete de assert de testify pero más especializado. Según su (blog de desarrollo e implementación)[https://www.negrel.dev/blog/zero-cost-debug-assertions-in-go/] esta herramienta facilita el uso de variables que no encajan en un tipo de datos específico, ofreciendo solución a los errores de programación más que de tiempo de ejecución. 
Fue revisada hace un mes desde la escritura de este apartado del documento, se puede saber más de su funcionamineto en la (página)[https://pkg.go.dev/gotest.tools/v3/assert] del paquete.

## Conclusión
**De momento vamos a emplear go testify, concretamente el paquete assert de go testify.**

Es la biblioteca más extendida además de tener gran compatiblidad con el test runner empleado, probaremos sus métodos para comparar las entradas y salidas de estructuras de datos.