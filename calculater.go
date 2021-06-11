package main

import (
	"fmt"
)

func main() {
	var command int
	var operand_1 float64
	var operand_2 float64
	for {
		fmt.Println("1.Сложение")
		fmt.Println("2.Вычитание")
		fmt.Println("3.Деление")
		fmt.Println("4.Умножение")
		fmt.Println("0.Выход")

		fmt.Println("Введите номер операции калькулятора: ")

		if _, err := fmt.Scanln(&command); err != nil {
			fmt.Println("Ошибка ввода команды, проверьте правильность ввода, необходимо ввести номер операции!")
		} else {
			if command >= 0 && command <= 4 {
				if command == 0 {
					fmt.Println("Да прибудет с Вами сила!")
					break
				}
				for {
					fmt.Printf("Введите первое число: ")
					_, err_op1 := fmt.Scanln(&operand_1)
					fmt.Printf("Введите второе число: ")
					_, err_op2 := fmt.Scanln(&operand_2)
					if err_op1 != nil || err_op2 != nil {
						fmt.Println("Ошибка ввода чисел, повторите ввод")
					} else {
						break
					}
				}
				switch command {
				case 1:
					fmt.Println("Результат операции:", operand_1, "+", operand_2, "=", operand_1+operand_2)
				case 2:
					fmt.Println("Результат операции:", operand_1, "-", operand_2, "=", operand_1-operand_2)
				case 3:
					fmt.Println("Результат операции:", operand_1, "/", operand_2, "=", operand_1/operand_2)
				case 4:
					fmt.Println("Результат операции:", operand_1, "*", operand_2, "=", operand_1*operand_2)
				}

			} else {
				fmt.Println("Операция с таким номером отсутствует, попробуйте еще раз")
			}

		}

	}

}
