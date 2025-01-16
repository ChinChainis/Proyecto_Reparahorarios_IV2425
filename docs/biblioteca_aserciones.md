# Biblioteca de aserciones

## Criterios de elección
- Debe tener un mantenimiento en la actualidad para no aumentar la deuda técnica; estudiaremos el estado de los repositorios de las cada herramienta para ver su seguimiento y descartaremos aquellos con más de 6 meses de inactividad.
- Debe necesitar el mínimo de paquetes, bibliotecas o cualquier software secundario para su funcionamiento, se descartará aquellos que necesiten de múltiples elementos externos para el testeo de código.
## Opciones disponibles para aserciones:

### Testify
**Según el [repositorio oficial](https://github.com/stretchr/testify)**

Testify es una serie de paquetes de uso bastante extendido poseyendo uno específico para el testeo con aserciones, tiene una instalación sencilla con solo escribir `$go get github.com/stretchr/testify` se instalan los 5 paquetes. Tiene un mantenimiento cuidado regularmente según su repositorio.

### Ghost
**Según el [repositorio oficial](https://github.com/rliebz/ghost)**

Sigue una orientación al lenguaje natural, definiendo con verbos qué testear y qué debería salir, necesita de incluir el paquete en el proyecto y su última revisión fue recientemente. 

### Assert
**Según el [repositorio oficial](https://github.com/negrel/assert)**

Es un paquete para proyectos en Go basado en el paquete de assert de testify. Para su uso se necesita importar al proyecto con `import{"github.com/negrel/assert"}` y fue revisada hace un mes desde la escritura de este apartado del documento.