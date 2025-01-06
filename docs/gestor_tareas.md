# Gestor de tareas

## Criterios de elección
- Debe tener un mantenimiento en la actualidad, no conviene utilizar herramientas desfasadas dando lugar a un aumento de
la deuda técnica; estudiaremos el estado de los repositorios de las cada herramienta para ver su seguimiento y descartaremos aquellos con más de un año de inactividad.
- Debe tener una comunidad activa, múltiples instancias de gente usando estas herramientas para sus proyectos con dudas entre usuarios, 
artículos o discusiones para la mejora del desarrollo de la herramienta y su trabajo con ella.
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

### Task
**Según su [repositorio oficial](https://github.com/go-task/task) y un [artículo](https://tsh.io/blog/taskfile-or-gnu-make-for-automation/) en comparación con make**

Task sirve de task runner y gestor de tareas orientado a golang, basado en emplear archivos yaml con órdenes bash. Se puede emplear también
para otros lenguajes. Puede ser por temas subjetivos pero requiere una sintaxis de implementación completamente distinta y 
múltiples herramientras funcionan de forma distinta como los condicionales o variables de entorno; otros gestores son más directos.

### XC
**Según su [repositorio oficial](https://github.com/joerdav/xc)**

XC o Execute es un task runner similar a make y npm run que se centra en ser más accesible; emplea con ficheros markdown la filosofía 
de programación literaria, el uso de lenguaje natural para el desarrollo de scripts para intentar reducir la deuda técnica de aprendizaje.
En cuanto a aspectos negativos veo que xc no es muy empleado por los usuarios o prefieren usar herramientas como [mage](https://www.reddit.com/r/golang/comments/12rmz4p/selfdocumenting_task_runner_define_tasks_in_the/)que siguen una filosofía similar.


### Just
**Según su [repositorio oficial](https://github.com/casey/just)**

Muy similar a make pero con un formato más parecido a un archivo con main y funciones llamado justfile, 
siendo bastante fácil de abordar. Es una herramienta especialmente rápida siendo útil a la hora de largas sesiones de trabajo o múltiples archivos. Just está escrita en el mismo lenguaje que GO facilitando compatibilidad y salvando posibles errores con otros gestores.


**Vamos a emplear just con un justfile.**

Vamos a probar el uso de just como gestor de tareas, se ha decantado por esta herramienta por cumplir los criterios de selección
además de aportar múltiples qualities of life a la hora del tiempo de desarrollo minimizando la deuda técnica; además es una 
herramienta potente si se diera el caso de que el proyecto aumentara drásticamente de volumen de archivos.