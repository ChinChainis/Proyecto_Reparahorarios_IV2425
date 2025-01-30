# Elección de imagen

## Criterios de elección
- La imagen debe tener un mantenimiento en la actualidad para no aumentar la deuda técnica; comprobaremos su estado en la página de (DockerHub)[https://hub.docker.com/].
- Buscaremos un tamaño de la imagen mínimo o lo más bajo posible, para mayor manejo en la nube y de tratar el código contenido.

## Opciones disponibles de imágenes para Golang:

### Imagen oficial de (Golang)[https://hub.docker.com/_/golang]

Viene junto a la instalación del lenguaje, fácil de emplear debido a la falta de configuración o descarga de cualquier paquete externo, mediante pruebas vemos que la imagen ocupa unos 868MB y es actualizado regularmente según su página correspondiente de Dockerhub, siendo la última hace 9 días.

### Imagen oficial de (Ubuntu)[https://hub.docker.com/_/ubuntu]

Puesto que el sistema operativo que empleamos es Ubuntu también nos resulta fácil de preparar su instalación, la imagen ocupa 868MB similar a la de Golang y es revisada mensualmente.