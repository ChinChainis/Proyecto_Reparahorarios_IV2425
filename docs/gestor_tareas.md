# Gestor de tareas

## Criterios de elección
- Debe tener un mantenimiento mínimo en la actualidad, no conviene utilizar herramientas de compilación desfasadas o 
sin un seguimiento activo de su correcto funcionamiento
- Para evitar el aumento de deuda técnica, las herramienta debe tener buena compatibilidad con el lenguaje para implementar buenas prácticas del lenguaje sin afectar al funcionamiento de compilación o chequeo de sintaxis.
## Opciones disponibles para gestor de tareas

### Make
**Según tutorialedge en el apartado de [golang](https://tutorialedge.net/golang/makefiles-for-go-developers/).**

Es posiblemente la más extendida y compatible entre múltiples lenguajes, además de tener una estructura sencilla e intuitiva, 
es constantemente revisada y actualizada para su correcto uso. 
Sin embargo, puede quedarse corto dependiendo de con qué proyecto se trabaje, dado que origina de script sencillos en bash, 
por otro lado hay que declarar varios parámetros añadidos para que compile correctamente en todos los sistemas operativos como Windows, pudiendo consumir tiempo extra de desarrollo.

### Mage
**Según la página oficial de [magefile](https://magefile.org/)**

Magefile es una herramienta orientada a lenguaje Go compatible desde la 1.7, no requiere de dependencias adicionales y es operativo en la mayoría de sistemas operativos a diferencia de make.

### Sage
**Según el repositorio oficial de [sage](https://github.com/einride/sage)**

Sage es una herramienta de compilación inspirada en Mage que emplea sagefiles, permite generar otros makefiles y permite también 
declarar dependencias pudiendo servir como gestor de dependencias. Es actualizada recurrentemente, la última revisión fue hace 2 semanas.
Sin embargo no tiene mucha documentación ni veo que sea muy usado por los usuarios de golang, haciendo que pueda haber poco registro de fallos o taras.

### Bazel
**Según [bazel.build](https://bazel.build/start/go?hl=es-419), en la página oficial de descarga**

Bazel es un software enfocado en la automatización de builds y test para proyectos, funciona en múltiples sistemas operativos. Es actualizado regularmente, según su repositorio hace tan solo 2 semanas que fue revisado.
También sirve como gestor de dependencias, haciendo que sus compilaciones son incrementales y siempre producirán el mismo resultado y
tiene muy buena escalabilidad pudiendo trabajar con proyectos de gran escala.
Según múltiples usuario Bazel está enfocados a proyectos grandes de empresas multinacionales u otros lenguajes y puede dar lugar 
a un grado de dificultad mayor que cualquier otra herramienta de compilación. Según este [artículo](https://medium.com/windmill-engineering/bazel-is-the-worst-build-system-except-for-all-the-others-b369396a9e26) explica su experiencia intentando usar bazel para proyectos de código abierto.

### Task
**Según su [repositorio oficial](https://github.com/go-task/task) y un [artículo](https://tsh.io/blog/taskfile-or-gnu-make-for-automation/) en comparación con make**

Task sirve de task runner y gestor de tareas orientado a golang, basado en emplear archivos yaml con órdenes bash. Se puede emplear también
para otros lenguajes. Puede ser por temas subjetivos pero requiere una sintaxis de implementación completamente distinta y 
múltiples herramientras funcionan de forma distinta como los condicionales o variables de entorno; otros gestores son más directos.

### Just
**Según su [repositorio oficial](https://github.com/casey/just)**

Muy similar a make pero con un formato más parecido a un archivo con main y funciones llamado justfile, 
siendo bastante fácil de abordar. Es una herramienta especialmente rápida siendo útil a la hora de largas sesiones de trabajo o múltiples archivos. Just está escrita en el mismo lenguaje que GO facilitando compatibilidad y salvando posibles errores con otros gestores.


**Vamos a emplear just con un justfile.**

Vamos a probar el uso de just como gestor de tareas, se ha decantado por esta herramienta por cumplir los criterios de selección
además de aportar múltiples qualities of life a la hora del tiempo de desarrollo minimizando la deuda técnica; además es una 
herramienta potente si se diera el caso de que el proyecto aumentara drásticamente de volumen de archivos.