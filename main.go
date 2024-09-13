package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// random - генератор псевдослучайных чисел
var random = rand.New(rand.NewSource(time.Now().UnixNano()))

var availableAttacks = map[string]func() int{
	"lite": getLiteAttack,
	"mid":  getMidAttack,
	"hard": getHardAttack,
}

// randNum возвращает случайное число в интервале [min, max]
func randNum(min, max int) int {
	return random.Intn(max-min+1) + min
}

// askUserValue запрашивает и возвращает ввод пользователя в консоли
func askUserValue(explanation string) string {
	fmt.Print(explanation)

	var userInput string
	_, err := fmt.Scanln(&userInput)

	if err != nil {
		fmt.Println(err)
	}

	return userInput
}

func setEnemyHealth() int {
	return randNum(80, 120)
}

func getLiteAttack() int {
	return randNum(2, 5)
}

func getMidAttack() int {
	return randNum(15, 25)
}

func getHardAttack() int {
	return randNum(30, 40)
}

func compareValues(enemyHealth, userTotalAttack int) bool {
	pointDifference := enemyHealth - userTotalAttack

	if pointDifference < 0 {
		pointDifference = -pointDifference
	}

	return pointDifference <= 10
}

func getUserAttack() int {
	var total int

	for i := 0; i < 5; {
		selectedAttack := askUserValue("Введи тип атаки: ")

		var attack, isExist = availableAttacks[selectedAttack]

		if !isExist {
			fmt.Println("Неизвестный тип атаки:", selectedAttack)
			continue
		}

		var damage = attack()

		fmt.Println("Количество очков твоей атаки:", damage)

		total += damage

		i++
	}

	return total
}

func runGame() bool {
	enemyHealth := setEnemyHealth()
	userTotalAttack := getUserAttack()

	fmt.Println("Тобой нанесён урон противнику равный", userTotalAttack)
	fmt.Println("Очки здоровья противника до твоей атаки", enemyHealth)

	if compareValues(enemyHealth, userTotalAttack) {
		fmt.Println("Ура! Победа за тобой!")
	} else {
		fmt.Println("В этот раз не повезло :( Бой проигран.")
	}

	answer := askUserValue("Чтобы сыграть ещё раз, введи букву [y] или [Y]: ")

	return strings.ToLower(answer) == "y"
}

func main() {
	intro := `РАССЧИТАЙ И ПОБЕДИ!
Загрузка...
	
Твоя цель — за 5 ходов набрать такое количество очков урона противнику,
которое попадет в диапазон +– 10 от значения здоровья противника.
	
Значение здоровья противника генерируется случайным образом
в диапазоне от 80 до 120 очков.
	
В твоём распоряжении три вида атак:
lite — урон от 2 до 5 очков;
mid — урон от 15 до 25 очков;
hard — урон от 30 до 40 очков.
ВПЕРЁД К ПОБЕДЕ!!!
`
	fmt.Println(intro)

	for runGame() {
	}
}
