# golang-generics-nerdearla

Este repositorio contiene el código que vamos a usar en el workshop de [Golang Generics de Nerdearla 2023](), dictado por Agustín Luques y Nicolás Del Piano.

## Introducción

**Go** ha ido evolucionando desde su versión inicial `1.0` hasta la más reciente `1.21` ([Go releases](https://go.dev/dl/)).

Uno de los grandes cambios introducidos últimamente en la version `1.18` es la implementación de _generics_ (tipos genéricos). La propuesta oficial fue conocida como [_type parameters_](https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md).

Pero... ¿para qué?

## Sistema de Tipos de Go

Un repaso rápido del sistema de tipos de Go:

### Tipos Básicos

Los tipos básicos _built-in_ que podemos encontrar en el lenguaje son:

Tipos numéricos:

- `int8`, `uint8` (`byte`), `int16`, `uint16`, `int32` (`rune`), `uint32`, `int64`, `uint64`, `int`, `uint`, `uintptr`
- `float32`, `float64`
- `complex64`, `complex128`

Booleanos:

- `bool`

Strings:

- `string`

Cada uno de estos tipos pueden ser usados en código Go sin importar ningún paquete externo.

### Zero Values

Cada tipo tiene un "valor cero" (_zero value_), el cual puede ser pensado como el valor por defecto del tipo, es decir, si no inicializamos la variable con un valor.

- El valor cero de un tipo booleano es `false`
- El valor cero de un tipo numérico es `0` (el tamaño en memoria puede variar de acuerdo al tipo)
- El valor cero de un tipo string es `""`

### Tipos Compuestos

## Generics

Go es un lenguaje de tipado estático, lo que significa que significa que el chequeo de tipos de las variables, funciones y parámetros se da en tiempo de compilación. Los tipos básicos junto con sus construcciones con `maps`, `slices` y `channels`, y las funciones asociadas como `len`, `cap`, o `make`, aceptan y retornan valores de diferentes tipos:

```go
arrayOfInts := []int{1,2,3}
arrayOfStrings := []string{"1","2","3"}

fmt.Println("El tamaño del arreglo de enteros es: " + len(arrayOfInts))
// Imprime: "El tamaño del arreglo de enteros es: 3"

fmt.Println("El tamaño del arreglo de strings es: " + len(arrayOfStrings))
// Imprime: "El tamaño del arreglo de strings es: 3"
```

Lo cual nos dice que tenemos soporte de genéricos para tipos built-in con las funciones ya definidas en el lenguaje.

¿Pero qué pasa con los tipos y funciones que definimos nosotros como programadores?

### Generics en Go

Los genéricos (o parámetros de tipo) en Go son una herramienta que nos da el lenguaje para poder especificar libremente el tipo de datos de los parámetros de una función o método, básicamente porque queremos mantenerlas lo más genéricas que se puedan.

Consideremos un caso de uso bastante simple. Tenemos una lista de precios y queremos conocer la suma total:

```go
type Precio int // Precio en centavos de alguna moneda

precios := []Precio{1, 2, 1000, 50}
```

Una solución probable es:

```go
func calcularTotalPrecios(precios []Precio) Precio {
  total := 0

  for _, precio := range precios {
    total += precio
  }

  return total
}
```

Ahora bien, supongamos que queremos obtener la distancia total:

```go
type Distancia int // Distancia en metros

distancias := []Distancia{100, 2000, 50}
```

Volvemos a definir la misma funcion, pero para el tipo `Distancia`:

```go
func calcularTotalDistancias(distancias []Distancia) Distancia {
  total := 0

  for _, distancia := range distancias {
    total += distancia
  }

  return total
}
```

Notamos un patrón que se repite:

```go
func calcularTotal[T Sumable](t []T) T {
  var total T = 0

  for _, element := range t {
    total += element
  }

  return total
}
```

Donde sumable esta definido por la siguiente interfaz:

```go
type Sumable interface {
  ~int
}
```

Y ahora, podemos directamente usar la función genérica `calcularTotal` gratis:

```go
calcularTotal(precios)
calcularTotal(distancias)
```

Utilizando el mismo nombre para ambos slices de diferentes tipos, reduciendo así el código repetitivo.

Otra ventaja es que los tests de `calcularTotal` van a cubrir ambos tipos.

## Constraints

Habrán notado en el ejemplo anterior la siguiente línea en la función `calcularTotal`:

```go
func calcularTotal[T Sumable](t []T)  T {
```

¿Qué es ese `Sumable` ahí?

Lo que le estamos diciendo al compilador con `[T Sumable]`, es que el `T` que le estamos pasando satisface la interfaz `Sumable`, lo que significa que podemos usar el operador `+` en elementos del tipo `T`.

Claramente `int` es `Sumable`.

Pero veámoslo con un ejemplo más concreto:

```go
type Sumable[T any] interface {
  Sumar(b T) T
}
```

Ahora, redefinamos `calcularTotal` como:

```go
func calcularTotal[T Sumable[T]](t []T) T {
  var total T // Usa el zero-value de T

  for _, e := range t {
    total = total.Sumar(e)
  }

  return total
}
```

Sabiendo que nuestro `T` satisface la interfaz `Sumable`, podemos utilizar la función `Sumar` en elementos de ese tipo.

### Constraints existentes

#### any

Recién vimos el ejemplo de una constraint hecha por nosotros `[T Sumable]`, pero si vemos cómo definimos nuestra interfaz:

```go
type Sumable[T any] interface {
  Sumar(b T) T
}
```

encontramos la keyword `any`. `any` no es mas que un alias a nuestro querido `interface{}` (el tipo que abarca todos los tipos en Go).

Sería como el tipo "bottom", no nos dice nada, no nos aporta mucha información más que podemos meter cualquier tipo ahí.

Supongamos que queremos redefinir `calcularTotal` usando `any`:

```go
func calcularTotal[T any](t []T) T {
  var total T

  for _, e := range t {
    total += e
  }

  return total
}
```

El compilador va a tirar el siguiente error:

    invalid operation: operator + not defined on total (variable of type T constrained by any)

El punto débil de usar `any` como constraint es que no podemos asumir nada sobre el tipo: puede ser cualquiera, lo cual nos limita a la hora de definir una función.

Las operaciones permitidas para variables de tipo `any` son:

- Declarar variables de ese tipo
- Asignaciones
- Usarlas en parámetros y valores de retorno
- Obtener la dirección de esas variables
- convert or assign values of those types to the type interface{}
- convert a value of type T to type T (permitted but useless)
- use a type assertion to convert an interface value to the type
- use the type as a case in a type switch
- define and use composite types that use those types, such as a slice of that type
- pass the type to some predeclared functions such as new

#### comparable

The new comparable keyword, in Go 1.18, was added for specifying types that can be compared with the == and != operators.

comparable is an interface that is implemented by all comparable types (booleans, numbers, strings, pointers, channels, interfaces, arrays of comparable types, structs whose fields are all comparable types). The comparable interface may only be used as a type parameter constraint, not as the type of a variable.

Comparable types include: structs, pointers, interfaces, channels, and builtin types. comparable can also be embedded in other constraints since it is a constraint.

#### constraints package

The Go team created a package of constraints (constraints) that can be imported and used for the most generic of contraint types. One important constraint is constraints.Ordered, which allows the use of the <, <=, >, and >= operators.

```go
type Ordered interface {
        Integer | Float | ~string
}
```

The definition of the above Ordered constraint makes sense it would allow the use of the <, <=, >, and >= operators since these types are normally comparable outside of a generic scope.

### Uso de constantes

Las constantes que usemos dentro de funciones o métodos tienen que satisfacer al valor más general posible de todos los tipos que abarcan el tipo genérico. Veamos un ejemplo:

```go
// No es válido!
func SumarMil[T Integer](v T) T {
  return v + 1_000
}

fmt.Println(SumarMil(100))
```

Retorna:

    ./prog.go:47:13: cannot convert 1_000 (untyped int constant 1000) to type T

ya que el tipo `int8` no puede representar ese valor.

Pero:

```go
// Válido!
func SumarCien[T Integer](v T) T {
  return v + 100
}

fmt.Println(SumarCien(100))
```

Retorna "200" como se esperaría.

## Interfaces versus Generics

## Reflection versus Generics

## Inferencia de Tipos

## Definiendo nuevos tipos de datos con Generics
