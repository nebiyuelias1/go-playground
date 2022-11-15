// ex3
package main

import "fmt"

func pay (money, bills float32)  float32 {
	return money - bills
}

func main ()  {
	var paycheck, bills float32 = 4000, 5000.0
	fmt.Printf("bank account balance: %f\n", pay(paycheck, bills))
}