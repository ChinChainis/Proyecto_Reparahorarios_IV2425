# Gestor de dependencias
Go emplea los denominados módulos para gestionar las dependencias de versión de los programas, no existen los gestores de dependencias canónicos sino que se emplean colecciones de paquetes que son lanzados, versionados y distribuidos.

## Descripción del gestor
Vamos a emplear Go Module para el gestor de dependencias, el cual es recogido en un archivo *go.mod* que contiene:
- El path del proyecto que va a tratar
- Dependencias necesarias especificando las versiones
