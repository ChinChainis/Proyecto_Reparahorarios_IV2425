## Milestones/Hitos

### [Milestone 0] Modelo del problema
Análisis detallado de las Historias de Usuario (HU), estos clientes no pueden hacer sus horarios de clases sin solapamientos o no saber cómo organizarse las horas de clases.

PMV: Código que atienda a las necesidades de los clientes, programa de organización de horarios.

El programa tendrá en cuenta todos esos elementos con sus atributos individuales (horas, grupos...) para la relación entre ellos para asegurar la mínima alteración.

Viabilidad: Deducir una solución para el problema presentado en las HUs.

Entidades encontradas en los HUs son implementadas en código fielmente.

### [Milestone 1] Implementación lógica de negocio
Necesito generar un nuevo horario u horarios partiendo de asignaturas, grupos, turnos y cursos concretos de un grado que sea factible, es decir que no haya ningún impedimento a la hora de asistir a una asginatura por coincidir con otra como en #2 y que cumpla los criterios pedidos por el cliente.

PMV: Creación de horarios por combinatoria de asignaturas, meter datos tipo diccionario de ejemplo y elaborar un horario coherente sin solapamientos.

Viabilidad: Cuando genere un horario fiel a los datos originales y sin que coincidan.

### [Milestone 2] Filtro de los distintos valores de un horario
Leer y categorizar la información obtenida de un horario para estudiar su correcta orgranización, obtener sus parámetros de horas a cursar, grado y curso al que pertenecen cada una.
Por ejemplo #2 solo quiere asignaturas con horario de turno de tarde.

PMV: filtro según criterios del cliente

Viabilidad: No descarta ningún elemento que cumpla las condiciones del filtro.

### [Milestone 3] Portal web abierto al público
Ofreciendo un servicio a distintos alumnos independientemente del grado como #3 y soportado online, solo buscaría los datos con una dirección a sus respectivos horarios. Servicio web similar al proceso de alteración de matrícula de un grado.

PMV: Página web de acceso por usuario externos. 

Viabilidad: Cuando sea ejecutable en un servidor y accesible mediante un navegador.

Terminología empleada [terminos.md](https://github.com/ChinChainis/Proyecto_Reparahorarios_IV2425/blob/Objetivo-1/docs/terminos.md)

User stories [user_journeys.md](https://github.com/ChinChainis/Proyecto_Reparahorarios_IV2425/blob/Objetivo-1/docs/user_stories.md)
