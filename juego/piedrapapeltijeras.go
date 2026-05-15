package main

import (
    "fmt"
    "math/rand"
    "strings"
    "time"
)

func main() {
    rand.Seed(time.Now().UnixNano())
    opciones := [3]string{"piedra", "papel", "tijera"}

    fmt.Println("¡Bienvenido a Piedra, Papel o Tijera!")

    for {
        fmt.Print("\nTu jugada (o escribe 'salir'): ")
        var usuario string
        fmt.Scan(&usuario)
        usuario = strings.ToLower(usuario)

        if usuario == "salir" {
            break
        }

        pc := opciones[rand.Intn(3)]
        fmt.Printf("PC eligió: %s\n", pc)

        if usuario == pc {
            fmt.Println("¡Empate! 😐")
        } else if (usuario == "piedra" && pc == "tijera") ||
            (usuario == "papel" && pc == "piedra") ||
            (usuario == "tijera" && pc == "papel") {
            fmt.Println("¡Ganaste! 🎉")
        } else {
            fmt.Println("Perdiste... 😞")
        }
    }
}