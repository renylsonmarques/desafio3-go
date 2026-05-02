package main

import (
	"fmt"
	"time"
)

// Nesse desafio do curso de GO na DIO.me, vou criar uma aplicação chamada ping pong
// Onde irei usar concorrência para criar uma aplicação que simule um jogo de ping pong, onde duas goroutines irão se comunicar usando canais para simular o jogo.
// No segundo 1 será ping e no segundo 2 será pong, e assim por diante, até o segundo 10, onde o jogo irá terminar.

func main() {
	ping := make(chan bool) // Criando um canal para a goroutine de ping
	pong := make(chan bool) // Criando um canal para a goroutine de pong

	// Criando a goroutine de ping
	go func() {
		for i := 1; i <= 10; i++ {
			if i%2 != 0 {
				fmt.Println("Ping") // Imprimindo "Ping" para os segundos ímpares
				ping <- true        // Enviando um valor para o canal de ping para indicar que o ping foi feito
			}
			time.Sleep(1 * time.Second)
		}
	}()

	// Criando a goroutine de pong
	go func() {
		for i := 1; i <= 10; i++ {
			if i%2 == 0 {
				fmt.Println("Pong") // Imprimindo "Pong" para os segundos pares
				pong <- true        // Enviando um valor para o canal de pong para indicar que o pong foi feito
			}
			time.Sleep(1 * time.Second)
		}
	}()

	// Loop para receber os valores dos canais de ping e pong, garantindo que o jogo seja sincronizado
	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			<-ping // Recebendo um valor do canal de ping para indicar que o ping foi feito
		} else {
			<-pong // Recebendo um valor do canal de pong para indicar que o pong foi feito
		}
	}

	fmt.Println("Fim do jogo!") // Imprimindo "Fim do jogo!" após o loop para indicar que o jogo terminou
}
