package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Personaje struct {
	Nombre  string
	Vida    int
	Ataque  int
	Defensa int
}

type Habilidad struct {
	Nombre string
	Poder  int
}

var habilidadesHeroe = []Habilidad{
	{"Golpe normal", 0},
	{"Golpe fuerte", 10},
	{"Magia de fuego", 20},
}

var habilidadesDragon = []Habilidad{
	{"Zarpazo", 5},
	{"Llamarada", 15},
	{"Rugido Mortal", 25},
}

func (p *Personaje) EstaVivo() bool {
	return p.Vida > 0
}

func (p *Personaje) Atacar(objetivo *Personaje) {
	dano := p.Ataque - objetivo.Defensa
	if dano < 1 {
		dano = 1
	}
	variacion := rand.Intn(11) - 5
	dano += variacion
	if dano < 1 {
		dano = 1
	}
	objetivo.Vida -= dano
	fmt.Printf("%s ataca a %s por %d de daño!\n", p.Nombre, objetivo.Nombre, dano)
	fmt.Printf("%s tiene %d de vida restante.\n\n", objetivo.Nombre, objetivo.Vida)
}

func elegirHabilidad(habilidades []Habilidad) Habilidad {
	fmt.Println("Elige tu habilidad:")
	for i, h := range habilidades {
		fmt.Printf("%d. %s (poder: +%d)\n", i+1, h.Nombre, h.Poder)
	}
	var opcion int
	fmt.Scan(&opcion)
	if opcion < 1 || opcion > len(habilidades) {
		fmt.Println("Opción inválida, se usó:", habilidades[0].Nombre)
		return habilidades[0]
	}
	return habilidades[opcion-1]
}

func main() {
	rand.Seed(time.Now().UnixNano())

	heroe := Personaje{"Héroe", 100, 20, 5}
	monstruo := Personaje{"Dragón", 80, 15, 3}
	turno := 1

	for heroe.EstaVivo() && monstruo.EstaVivo() {
		fmt.Printf("\n========== TURNO %d ==========\n", turno)
		fmt.Printf("Tu vida: %d  |  Vida del Dragón: %d\n\n", heroe.Vida, monstruo.Vida)

		hab := elegirHabilidad(habilidadesHeroe)
		heroe.Ataque += hab.Poder
		heroe.Atacar(&monstruo)
		heroe.Ataque -= hab.Poder

		if monstruo.EstaVivo() {
			habDragon := habilidadesDragon[rand.Intn(len(habilidadesDragon))]
			fmt.Printf("▲ El Dragón prepara: %s (poder: +%d)\n\n", habDragon.Nombre, habDragon.Poder)
			monstruo.Ataque += habDragon.Poder
			monstruo.Atacar(&heroe)
			monstruo.Ataque -= habDragon.Poder
		}
		turno++
	}

	fmt.Println("\n==============================")
	if heroe.EstaVivo() {
		fmt.Println("¡Ganaste! El héroe venció al dragón. 🏆")
	} else {
		fmt.Println("El dragón ganó. Inténtalo de nuevo. 💀")
	}
}
