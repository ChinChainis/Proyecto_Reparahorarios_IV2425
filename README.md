# IV_AntonioSanchezAlcaraz_2425
Repositorio para la asignatura de Infraestructura Virtual

## *Aplicación de administración de horarios*

**El cliente tiene un problema de incompatibilidad con los horarios de la carrera**

El cliente es un estudiante universitario en periodo de alteración de matrícula se plantea una aplicación web para administrar un horario compatible con las asignaturas elegidas y acorde a los criterios de matriculación del alumno, dado que cabe la posibilidad de matricularse de alguna asignatura y que coincida en horario con otra.

El programa debe tener en cuenta las asignaturas que quiera el alumno o en su defecto, no tener en cuenta las que ya haya cursado o no le interesen a este, la forma de introducir estos datos es de forma indiviual según el perfil académico de cada usuario.

También debe tener en cuenta las preferencias de turnos de por la mañana o tarde o ramas a excluir para no tener que introducir las asignaturas una a una de la mención a evitar.

Con un mayor avance del proyecto se podría implementar una función de obtener los examenes finales/entregas finales para comprobar si algún día coincide asemejándose a una aplicación de calendario pero siempre manteniendo informado al usuario de las asignaturas afectadas de forma personalizada.

Documentación del proyecto en la carpeta [docs](https://github.com/ChinChainis/Proyecto_Reparahorarios_IV2425/blob/Objetivo-4/docs/README.md)

Archivos de seguridad [docs](https://github.com/ChinChainis/Proyecto_Reparahorarios_IV2425/tree/Objetivo-4/docs/docs_seguridad_repositorio)


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


Análisis de los test runner planteados en [test_runners.md](https://github.com/ChinChainis/Proyecto_Reparahorarios_IV2425/blob/Objetivo-4/docs/test_runners.md)

> `go test -v`
En fichero de models