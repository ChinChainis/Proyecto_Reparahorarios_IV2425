# Gestor de tareas

## Criterios de elección
- Debe tener un mantenimiento mínimo en la actualidad, no conviene utilizar herramientas de compilación desfasadas o sin un seguimiento activo de su correcto funcionamiento
## Opciones disponibles para gestor de tareas

### Make
** Según tutorialedge en el apartado de golang **

Es posiblemente la más extendida y compatible entre múltiples lenguajes, además de tener una estructura sencilla e intuitiva, 
es constantemente revisada y actualizada para su correcto uso. 
Sin embargo, puede quedarse corto dependiendo de con qué proyecto se trabaje, dado que origina de script sencillos en bash, 
por otro lado hay que declarar varios parámetros añadidos para que compile correctamente en todos los sistemas operativos como Windows0,
 pudiendo consumir tiempo extra de desarrollo.

### Mage
** Según la página oficial de magefile **

Magefile es una herramienta orientada a lenguaje Go compatible desde la 1.7, no requiere de dependencias adicionales y es operativo en la mayoría de sistemas operativos a diferencia de make.
Lo único negativo que veo es que tiene un ritmo de actualizaciones irregular (la última fue el 11/5/2023).

### Sage
** Según el repositorio oficial **

Sage es una herramienta de compilación inspirada en Mage que emplea sagefiles, permite generar otros makefiles y permite también 
declarar dependencias pudiendo servir como gestor de dependencias. Es actualizada recurrentemente, la última revisión fue hace 2 semanas.
Sin embargo no tiene mucha documentación ni veo que sea muy usado por los usuarios de golang, haciendo que pueda haber poco registro de fallos o taras.