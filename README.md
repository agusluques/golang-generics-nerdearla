# golang-generics-nerdearla

Este repositorio contiene el código que vamos a usar en el workshop de [Golang Generics de Nerdearla 2023](), dictado por Agustín Luques y Nicolás Del Piano.

## Contenido

1. [Introducción](#introduccion)
1. [Sistema de Tipos de Go](#type-system)
1. [Generics](#generics)
   1. [Generics en Go](#generics-en-go)
   1. [Notación](#notacion)
   1. [Funciones Genéricas](#funciones-genericas)
   1. [Tipos Genéricos](#tipos-genericos)
1. [Constraints](#constraints)
1. [Inferencia de Tipos](#inferencia-de-tipos)
1. [Encadenamiento de Tipos (Type Chaining)](#type-chaining)
1. [Generics Múltiples](#generics-multiples)
1. [Interfaces versus Generics](#interfaces-vs-generics)
1. [Reflection versus Generics](#reflection-vs-generics)
1. [Estado del Arte](#estado-del-arte)

<div id='introduccion'/>

## Introducción

**Go** ha ido evolucionando desde su versión inicial `1.0` hasta la más reciente `1.21` ([Go releases](https://go.dev/dl/)).

Uno de los grandes cambios introducidos últimamente en la version `1.18` es la implementación de _generics_ (tipos genéricos). La propuesta oficial fue conocida como [_type parameters_](https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md).

Pero... ¿para qué nos sirve esto?

<div id='type-system'/>

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

<div id='generics'/>

## Generics

Go es un lenguaje de tipado estático, por lo tanto el chequeo de tipos de las variables, funciones y parámetros se da en tiempo de compilación. Los tipos básicos junto con sus construcciones con `maps`, `slices` y `channels`, y las funciones asociadas como `len`, `cap`, o `make`, aceptan y retornan valores de diferentes tipos:

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

Los tipos genéricos (o también llamados _parámetros de tipo_) en Go nos permiten parametrizar el tipo de datos de los argumentos de una función para mantenerlos lo más abstractos que se puedan y poder definir funciones más genéricas, evitando la repetición de código.

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

Ahora bien, supongamos que tenemos otro tipo `Distancia` y queremos obtener la distancia total en un slice:

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

Notamos un patrón que se repite, dado un `T` con ciertas propiedades:

```go
func calcularTotal[T ???](t []T) T {
  var total T = 0

  for _, element := range t {
    total += element
  }

  return total
}
```

Si reemplazamos `???` por la interfaz `Sumable` definida como:

```go
type Sumable interface {
  ~int
}
```

Podemos directamente usar la función genérica `calcularTotal` gratis en los slices de ambos tipos `Precio` y `Distancia`:

```go
calcularTotal(precios)
calcularTotal(distancias)
```

💥 reduciendo así la repetición de código.

### Notación

```go
// Define la función f con el parámetro de tipo T
func f[T any](t T) {
  // ...
}

// Llama a la función f especificando el parámetro int
f[int](10)

// Define una estructura con parámetro de tipo T
type e[T any] struct {
  t T
}

// Instancia la estructura especificando el parámetro de tipo int
val := s[int]{t: 1}

// Define una interface con el parámetro de tipo T
type i[T any] interface {
  HacerAlgo(t T) T
}
```

<div id="funciones-genericas">

### Funciones Genéricas

Uno de los grandes problemas que teníamos en Go era la definición de funciones que hacian escencialmente lo mismo para distintos tipos, pero por una limitación del lenguaje, no podíamos generalizarla y terminábamos copiando y pegando la definición de `Map` en todos los lugares que necesitábamos.

Con generics, definir funciones generales se facilita bastante:

```go
// Map convierte []T1 a []T2 usando una función de mapeo.
// Esta función tiene dos parámetros de tipo, T1 and T2.
// Esto funciona con slices de cualquier tipo (especificado por "any").
func Map[T1, T2 any](s []T1, f func(T1) T2) []T2 {
    r := make([]T2, len(s))
    for i, v := range s {
        r[i] = f(v)
    }
    return r
}

words := []string{"hello", "world"}
lengths := Map[string, int](words, func(word string) int {
  return len(word)
})
fmt.Println(lengths) // [5 5]
```

<div id="tipos-genericos">

### Tipos Genéricos

Hay otro lugar donde vamos a ver parámetros de tipos, y es en definiciones de tipos o interfaces.

Por ejemplo, si queremos definir la estructura `Tupla` que contiene dos elementos del mismo tipo:

```go
type Tupla[T any] struct {
  t1 T
  t2 T
}

tupla := Tupla{t1: 1, t2: 2}
```

También podemos crear interfaces que contengan parámetros de tipos:

```go
type Sumable[T any] interface {
  fmt.Stringer
  Sumar(T) T
}

type Entero int

func (e Entero) String() string {
  return fmt.Sprintf("%d", e)
}

func (e Entero) Sumar(b Entero) Entero {
  return e + b
}

func Suma[T Sumable[T]](a, b T) T {
	return a.Sumar(b)
}

// Entero satisface la interfaz genérica Sumable
fmt.Println(Suma[Entero](1, 100)) // 101
```

## Constraints

Habrán notado en el ejemplo de la sección `Generics en Go` la siguiente línea en la función `calcularTotal`:

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

La nueva keyword `comparable` se introdujo en la versión de Go `1.18` y sirve para especificar tipos que pueden compararse, esto es, pueden usar los operadores `==` y `!=`.

Casi todos los tipos built-in implementan la interfaz `comparable` (booleanos, números, strings, punteros, canales, interfaces, arreglos de tipos comparables, etc).

Hay que tener en cuenta que solo puede usarse como una constraint en generics, y no como un tipo de una variable:

```go
var x comparable // error: cannot use type comparable outside a type constraint: interface is (or embeds) comparable
```

#### constraints package

The Go team created a package of constraints (constraints) that can be imported and used for the most generic of contraint types. One important constraint is constraints.Ordered, which allows the use of the <, <=, >, and >= operators.

```go
type Ordered interface {
        Integer | Float | ~string
}
```

The definition of the above Ordered constraint makes sense it would allow the use of the <, <=, >, and >= operators since these types are normally comparable outside of a generic scope.

#### Underlying Types

Habrás notado la notación `~T` en los ejemplos anteriores.

La tilde `~` significa que el tipo `T` es de tipo `T`, o es equivalente a `T`.

Por ejemplo, si definimos:

```go
type Precio int
```

Sabemos que "por debajo" `Precio` es lo mismo que un `int`: debería poder hacer las mismas cosas.

Entonces:

```go
func SumarMil[T ~int](v T) T {
  return v + 1_000
}

var p Precio = 100

fmt.Println(SumarMil(p)) // Imprime 1100
```

¿Qué pasaría si no usaramos la tilde?

```go
func SumarMil[T int](v T) T {
  return v + 1_000
}

var p Precio = 100

fmt.Println(SumarMil(p))
// Precio does not satisfy int (possibly missing ~ for int in int)
```

#### Unions

También podemos definir uniones de tipos, que representan la unión de los conjuntos de valores que puede aceptar un tipo `T`, por ejemplo:

```go
func Resta[T int | uint](a, b T) T {
	return a - b
}

var a uint = 10
var b uint = 2

Resta(a,b) // 8
```

Una notación conveniente es usar interfaces para expresar union types, por ejemplo:

```go
type Integer interface {
  int | uint
}

func Resta[T Integer](a, b T) T {
  return a - b
}
```

#### Combinaciones

Podemos combinar todas las constraints vistas hasta este momento en una interfaz:

```go
type ValorMoneda interface {
	~int | ~int64
}

type Moneda interface {
	ValorMoneda
	ISO4127Code() string
	Decimal() int
}

type ARS int64

func ImprimirBalance[T Moneda](m T) {
	balance := float64(m) / math.Pow10(m.Decimal())
	fmt.Printf("%.*f %s\n", m.Decimal(), balance, m.ISO4127Code())
}

func (a ARS) ISO4127Code() string {
	return "ARS"
}

func (a ARS) Decimal() int {
	return 2
}
```

### Uso de constantes

Las constantes que usemos dentro de funciones o métodos tienen que satisfacer al valor más general posible de todos los tipos que abarcan el tipo genérico.

Veamos un ejemplo:

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

### Ventajas de usar Constraints

Una de las principales ventajas que nos dan las constraints es que podemos definir la lista de tipos que permitimos en nuestra función en la misma declaración de generics:

```go
func Suma[A Sumable[A]](a, b A) A {
  // Sabemos que podemos sumar a + b y obtendremos un tipo A
}
```

En general, las constraints nos ayudan a escribir código más legible, testeable y mantenible.

## Inferencia de Tipos

Go ofrece inferencia de tipos cuando usamos el operador `:=`:

```go
x := 1
fmt.Println(reflect.TypeOf(x)) // int
```

Y también ofrece inferencia con Generics para simplificar las llamadas a funciones:

```go
func Contiene[T comparable](t []T, v T) bool {
  for _, x := range t {
    if x == v {
      return true
    }
  }

  return false
}

listaStrings := []string{"1", "2"}
Contiene(listaStrings, "1") // true

listaInts := []int{1, 2}
Contiene(listaInts, 100) // false
```

Vemos que podemos omitir los argumentos de tipo para `Contiene`: Go lo hace por nosotros.

En algunas situaciones sí hay que especificarlo, como por ejemplo, si el tipo genérico se usa en el valor de retorno:

```go
type Integer interface {
    int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64
}

func Convert[T1, T2 Integer](in T1) T2 {
    return T2(in)
}

func main() {
    var a int = 10
    b := Convert[int](a) // error: can't infer the return type
    fmt.Println(b)
}
```

## Encadenamiento de Tipos (Type Chaining)

_Type chaning_ es una técnica que nos permite definir tipos genéricos componiendo tipos definidos en la misma declaración.

Por ejemplo:

```go
func MapToString[L ~E[], E fmt.Stringer](l L) []string {
  resultado := make([]string, len(l))

  for i, e := range l {
    resultado[i] = e.String()
  }

  return resultado
}
```

Acá estamos requiriendo que el tipo de `L` depende del tipo de `E` que es un `fmt.Stringer`, por lo tanto estamos _encadenando_ los tipos.

### Inferencia de Tipos en Constraints (Constraint Type Inference)

En el ejemplo de arriba, podemos notar como el parámetro de tipo `L` está definido como `~E[]` (un compuesto de `E`) y el tipo de `E` es un `fmt.Stringer`. `L` puede ser inferido sabiendo el tipo de `E` cuando se llama a la función `MapToString`.

El compilador va a determinar el tipo de `L` cuando `MapToString` es llamado con el argumento. Por ejemplo, si implementásemos una struct `Persona` que satisfaga la interfaz `fmt.Stringer`:

```go
type Persona struct {
  Name string
}

func (p Persona) String() string {
  return fmt.Sprintf("%s", p.Name)
}

MapToString([]Persona{{Name: "Jack Sparrow"}})
```

El tipo de `L` será inferido como `[]Persona`.

## Generics Múltiples

Go permite especificar parámetros de tipo múltiples, por ejemplo:

```go
func ImprimirValores[A int, B, C any, D ~int](a A, b B, c1, c2 C, d D) {
  fmt.Printf("%v %v %v %v %v", a, b, c1, c2, d)
}
```

Acá las restricciones van a ser que `A` es un `int`, `B` y `C` pueden ser cualquier cosa y `D` de cualquier tipo subyacente que represente un `int`, pero ademas los parámetros `c1` y `c2` tienen que ser del mismo tipo:

```go
ImprimirValores(1, 2.0, "3", 4, 5) // error: mismatched types untyped string and untyped int (cannot infer C)
```

Dado que el tercer argumento `c1` es de tipo `string` y el cuarto `c2` de tipo `int`, pero los type parameters requieren que sean del mismo tipo:

```go
ImprimirValores(1, 2.0, "3", "4", 5) // 1 2 3 4 5
```

Notar que si especificamos tipos, debemos siempre completar los de la izquierda (no podemos saltearnos argumentos de tipos):

```go
ImprimirValores[int, float32, string](1, 2.0, "3", "4", 5) // 1 2 3 4 5
```

⚠️ Cuidado con definir varios parámetros de tipo si queremos especificar que sean iguales:

```go
func Iguales[T1, T2 comparable](a T1, b T2) bool {
  return a == b
}

var a int = 1
var b int = 2

Iguales(a, b) // error: invalid operation: a == b (mismatched types T1 and T2)
```

Una correcta definición sería garantizando que son del mismo tipo:

```go
func Iguales[T comparable](a, b T) bool {
  return a == b
}

var a int = 1
var b int = 2

Iguales(a, b) // false
Iguales(a, 1) // true
```

## Interfaces versus Generics

## Reflection versus Generics

## Estado del Arte

En esta sección veremos qué limitaciones tenemos con los Generics en la versión actual de Go `1.21`.

### ❌ Tipos genéricos en alias

Una declaración de alias de tipo no puede tener un parámetro de tipo:

```go
type T[X, Y any] func(X) Y

type A = T[int, string] // OK

type B[X any] = T[X, X] // Error: generic type cannot be alias
```

### ❌ Métodos Genéricos

Actualmente los [métodos no soportan parámetros de tipos](https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#methods-may-not-take-additional-type-arguments).

Es decir, el siguiente código no compila:

```go
type Moneda struct {
  valor int
}

func (m *Moneda) Valor[T ~int]() T {
  return m.valor // syntax error: method must have no type parameters
}
```

Existe una [issue en GitHub](https://github.com/golang/go/issues/49085) sobre esto.

### ❌ Embeber Genéricos

```go
type Derived[Base any] struct {
	Base // error: embedded field type cannot be a (pointer to a) type parameter
	x bool
}
```
