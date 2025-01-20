# Biblioteca de aserciones

## Criterios de elección
- Debe tener un mantenimiento en la actualidad para no aumentar la deuda técnica; descartaremos aquellos cuyos repositorios no hayan tenido ninguna revisión o commit nuevo en más de 6 meses.
- Debe ser compatible con el test runner asignado, en este caso es la herramienta nativa del lenguaje go testing.
- Evitar aquellas que empleen múltiples instalaciones de librerías externas, daremos prioridad a aquellas optimizadas para el lenguaje go para evitar posible deuda técnica que genere las incompatibilidades.

## Opciones disponibles para aserciones:

### Testify
**Según el [repositorio oficial](https://github.com/stretchr/testify)**

Testify es una serie de paquetes de uso bastante extendido poseyendo uno específico para el testeo con aserciones llamado `assert package`. Este paquete es revisado regularmente según el seguimiento de commits de su repositorio y testify es completamente orientado a lenguaje Go, testify/assert asegura que emplea salidas compatible con las funciones de go test lo cual nos interesa si lo tenemos como test runner.

### Ghost
**Según el [repositorio oficial](https://github.com/rliebz/ghost)**

Sigue una orientación al lenguaje natural, tiene múltiples palabras reservadas para intentar mejorar de forma intuitiva qué emplea en cada momento, necesita de incluir la librería en el proyecto; tiene un arduo seguimiento, hace 2 semanas fue revisado, así que está en constante actualización en caso de encontrarse algun error.

### Assert
**Según el [repositorio oficial](https://github.com/negrel/assert)**

Es un paquete para proyectos en Go basado en el paquete de assert de testify pero más especializado. Necesita importar al proyecto con `import{"github.com/negrel/assert"}` y fue revisada hace un mes desde la escritura de este apartado del documento.

**De momento vamos a emplear go testify.**

Es la biblioteca más extendida además de tener gran compatiblidad con el test runner empleado, probaremos sus métodos para comparar las entradas y salidas de estructuras de datos.