# golang-generics-nerdearla

Este repositorio contiene el código que vamos a usar en el workshop de [Golang Generics de Nerdearla 2023](), dictado por Agustín Luques y Nicolás Del Piano.

## Introducción

**Go** ha ido evolucionando desde su versión inicial `1.0` hasta la más reciente `1.21` ([Go releases](https://go.dev/dl/)).

Uno de los grandes cambios introducidos últimamente en la version `1.18` es la implementación de _generics_ (tipos genéricos). La propuesta oficial fue conocida como [_type parameters_](https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md).

Pero... ¿qué es esto?

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

```go

```

## Interfaces

## Inferencia de Tipos
