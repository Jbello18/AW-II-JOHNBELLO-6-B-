# REFLEXION.md - Taller Semana 1

**1. ¿Cuántas líneas tiene tu función main al final del taller?**
Mi función "main" tiene aproximadamente 75 líneas de código. Esto incluye la declaración de variables, el bucle de control, el switch de operaciones y la gestión del historial.

**2. Si tuvieras que agregar 5 operaciones más (raíz cuadrada, logaritmo, seno, coseno, módulo), ¿qué tan grande se haría tu main?**
El "main" crecería de forma desproporcionada. Cada nueva operación requeriría un nuevo "case" en el "switch" con su lógica de cálculo y formateo de texto, superando fácilmente las 120 líneas de código.

**3. ¿Sería fácil de leer para alguien que vea el código por primera vez?**
No sería fácil de leer. Al tener toda la lógica de entrada de datos, procesamiento matemático y salida de historial en un solo bloque, el código se vuelve denso y difícil de seguir para un tercero.

**4. Notaste que el código para 'pedir un número al usuario' o 'imprimir el resultado' se repite varias veces. ¿No sería mejor escribirlo una sola vez y reutilizarlo en muchos lugares?**
Sí, sería mucho mejor. La repetición de código "fmt.Scanln" y "fmt.Printf" aumenta el riesgo de errores. Lo ideal sería hacer que esas acciones en funciones reutilizables para mantener el "main" limpio y organizado.

**5. Después de este taller, ¿qué fue lo más difícil de Go para ti? ¿Y lo más interesante?**
* Lo más difícil:Entender el uso de los punteros "&" en la función "Scanln" para que Go sepa exactamente dónde guardar el valor ingresado en la memoria.
* Lo más interesante: La versatilidad del bucle "for" para crear tanto ciclos de repetición controlados (potencia/factorial) como bucles infinitos para mantener el programa activo.