package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func attack(charName, charClass string) string {
	if charClass == "warrior" {
		return fmt.Sprintf("%s нанес урон противнику равный %d.", charName, randint(8, 10))
	}

	if charClass == "mage" {
		return fmt.Sprintf("%s нанес урон противнику равный %d.", charName, randint(10, 15))
	}

	if charClass == "healer" {
		return fmt.Sprintf("%s нанес урон противнику равный %d.", charName, randint(2, 4))
	}
	return "неизвестный класс персонажа"
}

// обратите внимание на "if else" и на "else"
func defence(charName, charClass string) string {
	if charClass == "warrior" {
		return fmt.Sprintf("%s блокировал %d урона.", charName, randint(15, 20))
	}

	if charClass == "mage" {
		return fmt.Sprintf("%s блокировал %d урона.", charName, randint(8, 12))
	}

	if charClass == "healer" {
		return fmt.Sprintf("%s блокировал %d урона.", charName, randint(12, 15))
	}

	return "неизвестный класс персонажа"

}

// обратите внимание на "if else" и на "else"
func special(charName, charClass string) string {
	if charClass == "warrior" {
		return fmt.Sprintf("%s применил специальное умение `Выносливость %d`", charName, 105)
	}

	if charClass == "mage" {
		return fmt.Sprintf("%s применил специальное умение `Атака %d`", charName, 45)
	}

	if charClass == "healer" {
		return fmt.Sprintf("%s применил специальное умение `Защита %d`", charName, 40)
	}

	return "неизвестный класс персонажа"

}

// здесь обратите внимание на имена параметров
func startTraining(charName, charClass string) string {
	if charClass == "warrior" {
		fmt.Printf("%s, ты Воитель - отличный боец ближнего боя.\n", charName)
	}

	if charClass == "mage" {
		fmt.Printf("%s, ты Маг - превосходный укротитель стихий.\n", charName)
	}

	if charClass == "healer" {
		fmt.Printf("%s, ты Лекарь - чародей, способный исцелять раны.\n", charName)
	}

	fmt.Println(`Потренируйся управлять своими навыками.
Введи одну из команд: attack — чтобы атаковать противника,
defence — чтобы блокировать атаку противника,
special — чтобы использовать свою суперсилу.
Если не хочешь тренироваться, введи команду skip.`)

	var cmd string
	for cmd != "skip" {
		fmt.Print("Введи команду: ")
		fmt.Scanf("%s\n", &cmd)

		switch cmd {
		case "attack":
			fmt.Println(attack(charName, charClass))
		case "defence":
			fmt.Println(defence(charName, charClass))
		case "special":
			fmt.Println(special(charName, charClass))
		case "skip":
			return "тренировка окончена"
		default:
			fmt.Println("неизвестная команда")
		}
	}

	return "тренировка окончена"
}

// обратите внимание на имя функции и имена переменных
func choiceCharClass() string {
	var approveChoice string
	var charClass string

	for approveChoice != "y" {
		fmt.Print("Введи название персонажа, за которого хочешь играть: Воитель — warrior, Маг — mage, Лекарь — healer: ")
		fmt.Scanf("%s\n", &charClass)

		if charClass == "warrior" {
			fmt.Println("Воитель — дерзкий воин ближнего боя. Сильный, выносливый и отважный.")
		}

		if charClass == "mage" {
			fmt.Println("Маг — находчивый воин дальнего боя. Обладает высоким интеллектом.")
		}

		if charClass == "healer" {
			fmt.Println("Лекарь — могущественный заклинатель. Черпает силы из природы, веры и духов.")
		}

		fmt.Print("Нажми (Y), чтобы подтвердить выбор, или любую другую кнопку, чтобы выбрать другого персонажа: ")
		fmt.Scanf("%s\n", &approveChoice)
		approveChoice = strings.ToLower(approveChoice)
	}
	return charClass
}

// обратите внимание на имена переменных
func main() {
	fmt.Println("Приветствую тебя, искатель приключений!")
	fmt.Println("Прежде чем начать игру...")

	var charName string
	fmt.Print("...назови себя: ")
	fmt.Scanf("%s\n", &charName)

	fmt.Printf("Здравствуй, %s\n", charName)
	fmt.Println(`Сейчас твоя выносливость — 80, атака — 5 и защита — 10.
Ты можешь выбрать один из трёх путей силы:
Воитель, Маг, Лекарь`)

	charClass := choiceCharClass()

	fmt.Println(startTraining(charName, charClass))
}

func randint(min, max int) int {
	return rand.Intn(max-min) + min
}
