package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func action(charName, charClass, actionType string) string {
	var value int
	var specialSkill string
	switch actionType {
	case "attack":
		switch charClass {
		case "warrior":
			value = randint(8, 10)
		case "mage":
			value = randint(10, 15)
		case "healer":
			value = randint(2, 4)
		}
		return fmt.Sprintf("%s нанес урон противнику равный %d.", charName, value)
	case "defence":
		switch charClass {
		case "warrior":
			value = randint(15, 20)
		case "mage":
			value = randint(8, 12)
		case "healer":
			value = randint(12, 15)
		}
		return fmt.Sprintf("%s блокировал %d урона.", charName, value)
	case "special":
		switch charClass {
		case "warrior":
			value = 105
			specialSkill = "Выносливость"
		case "mage":
			value = 45
			specialSkill = "Атака"
		case "healer":
			value = 40
			specialSkill = "Защита"
		}
		return fmt.Sprintf("%s применил специальное умение `%s %d`", charName, specialSkill, value)
	default:
		return "неизвестная команда"
	}
}

func startTraining(charName, charClass string) string {
	fmt.Printf("%s, ты %s.\n", charName, charClass)

	fmt.Println(`Потренируйся управлять своими навыками.
Введи одну из команд: attack, defence, special.
Если не хочешь тренироваться, введи команду skip.`)

	for {
		var cmd string
		fmt.Print("Введи команду: ")
		fmt.Scanln(&cmd)

		switch cmd {
		case "attack", "defence", "special":
			fmt.Println(action(charName, charClass, cmd))
		case "skip":
			return "тренировка окончена"
		default:
			fmt.Println("неизвестная команда")
		}
	}
}

func choiceCharClass() string {
	var approveChoice string
	var charClass string

	for strings.ToLower(approveChoice) != "y" {
		fmt.Print("Введи название персонажа, за которого хочешь играть: Воитель — warrior, Маг — mage, Лекарь — healer: ")
		fmt.Scanln(&charClass)

		switch charClass {
		case "warrior":
			fmt.Println("Воитель — дерзкий воин ближнего боя. Сильный, выносливый и отважный.")
		case "mage":
			fmt.Println("Маг — находчивый воин дальнего боя. Обладает высоким интеллектом.")
		case "healer":
			fmt.Println("Лекарь — могущественный заклинатель. Черпает силы из природы, веры и духов.")
		default:
			fmt.Println("неизвестный класс персонажа")
			continue
		}

		fmt.Print("Нажми (Y), чтобы подтвердить выбор, или любую другую кнопку, чтобы выбрать другого персонажа: ")
		fmt.Scanln(&approveChoice)
	}
	return charClass
}

func randint(min, max int) int {
	return rand.Intn(max-min) + min
}

func main() {
	fmt.Println("Приветствую тебя, искатель приключений!")
	fmt.Println("Прежде чем начать игру...")

	var charName string
	fmt.Print("...назови себя: ")
	fmt.Scanln(&charName)

	fmt.Printf("Здравствуй, %s\n", charName)
	fmt.Println(`Сейчас твоя выносливость — 80, атака — 5 и защита — 10.
Ты можешь выбрать один из трёх путей силы:
Воитель, Маг, Лекарь`)

	charClass := choiceCharClass()

	fmt.Println(startTraining(charName, charClass))
}
