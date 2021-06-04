# apiMatriz
 Backend - Rotar matriz en sentido horario

Se tiene el requerimiento de crear un nuevo backend interno Restful, con diversos endpoints que realizan cálculos distintos.

La primera versión del backend empezará implementando el primero de todos los servicios (MVP), porque los demás servicios aún están en definición.

En base a ello, se requiere implementar un backend con un API Rest (Golang) que reciba un array de  arrays  de  números  que  conformen  una  matriz  de  NxN  representando  una  imagen,  y  se devuelva  la  misma matriz que  represente  la  imagen pero  rotada  en  sentido anti-horario  (90 grados). La API debe ser de tipo POST, el Content-Type de los request y response deben ser application/json y se debe controlar correctamente los errores. Adicional, el backend debe estar dockerizado. Considerar que esta primera versión será la base para que posteriormente otros desarrolladores puedan incluir los demás endpoints que aún están en definición, por lo que la solución debe garantizar la mantenibilidad y escalamiento.

Se evaluará el correcto funcionamiento del API según las especificaciones siendo eficiente y usando las mejores prácticas de desarrollo. Enviar la solución completa (como zip o la URL del repo) y un ejemplo de ejecución.




Ejemplo 1

Input: [ [1,2], [3,4] ]
Output: [ [2,4], [1,3] ]

