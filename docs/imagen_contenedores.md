# Elección de imagen

## Criterios de elección
- La imagen debe tener un mantenimiento en la actualidad para no aumentar la deuda técnica; comprobaremos su estado en la página de (DockerHub)[https://hub.docker.com/].
- Buscaremos un tamaño de la imagen mínimo o lo más bajo posible, para mayor manejo en la nube y de tratar el código contenido.

## Opciones disponibles de imágenes para Golang:

### Imagen oficial de (Golang)[https://hub.docker.com/_/golang]

Viene junto a la instalación del lenguaje, fácil de emplear debido a la falta de configuración o descarga de cualquier paquete externo, mediante pruebas vemos que la imagen ocupa unos 868MB y es actualizado regularmente según su página correspondiente de Dockerhub, siendo la última hace 9 días.

### Imagen oficial de (Ubuntu)[https://hub.docker.com/_/ubuntu]

Puesto que el sistema operativo que empleamos es Ubuntu también nos resulta fácil de preparar su instalación, la imagen ocupa 868MB similar a la de Golang y es revisada mensualmente.

### Golang:(Alpine)[https://hub.docker.com/_/ubuntu]

Alpine es un sistema operativo centrado en la seguridad y emplea tablas musl para codificar información, permitiendo un método de compresión mejor obteniendo un tamaño de 249MB (para este proyecto va bien pero según discusiones en (reddit)[https://www.reddit.com/r/golang/comments/t7hsps/building_a_docker_image_for_a_go_programm/] este tipo de tablas pueden dar lugar a bastantes problemas a la hora de almacenar según qué datos) y cuenta con revisiones recientes, la última siendo hace 22 días.

## Conclusión final:

A pesar de leer múltiples discusiones en foros sobre los problemas de la imagen he decidido que vamos a emplear golang:alpine ; la mayoría suelen venir de que tienen problemas de incompatiblidad con ciertos archivos por su sistema de cifrado sin embargo es con diferencia la que menos ocupa y tiene un mantenimiento ideal para ser empleado.