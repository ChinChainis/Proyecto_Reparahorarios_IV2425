# IV_AntonioSanchezAlcaraz_2425
Repositorio para la asignatura de Infraestructura Virtual

## Ejecución de tareas
Se requiere tener instalado la herramienta just para trabajar con el gestor de tareas.

Análisis de los task runner planteados en [gestor_de_tareas.md](https://github.com/ChinChainis/Proyecto_Reparahorarios_IV2425/blob/Objetivo-4/docs/gestor_tareas.md)

> `just clean`

Borramos ficheros binarios si se han generado.

> `just check`

Comprueba sintaxis de los ficheros con código.

> `just version`

Muestra la versión empleada de go.

> `just build`

Compila código de la carpeta de modelos.

## Test
Análisis de los test runner planteados en [test_runners.md](https://github.com/ChinChainis/Proyecto_Reparahorarios_IV2425/blob/Objetivo-4/docs/test_runners.md)

> `just test`

Ejecuta ficheros de testeo de las entidades del proyecto.

## Docker
Análisis de la imagen escogida para el contenedor en [imagen_contenedores.md](https://github.com/ChinChainis/Proyecto_Reparahorarios_IV2425/blob/Objetivo-5/docs/imagen_contenedores.md)

> `docker build . -t <repositorio>:<etiqueta>`

Ejecuta ficheros de testeo de las entidades del proyecto descargando remotamente la imagen del contenedor.

> `docker run -u 1001 -t -v `pwd`:/app/test antoniosanchz/reparahorarios:v1.0`