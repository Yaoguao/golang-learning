package main

import "fmt"

//	Задание 8
//	Дана переменная int64. Разработать программу которая устанавливает i-й бит в
//	1 или 0.
//	* Для меня очень плохо сформулировано задание, не понимаю

// пока 8 бит для теста, а то много нулей будет ([]uint8, 8)
func ConvertIntToBits(n int) []uint8 {
	bits := make([]uint8, 8)
	for i := 7; i >= 0; i-- {
		bits[7-i] = uint8((n >> i) & 1)
	}
	return bits
}

func ConvertBitsToInt(bits []uint8) int {
	if len(bits) != 8 {
		panic("bits must be 8 bits long")
	}
	length := 0
	for i := 0; i < 8; i++ {
		length = (length << 1) | int(bits[i])
	}
	return length
}

// Ну тут вообще пушка, надеюсь я правильно понял задание :D
func main() {
	var res, index int
	fmt.Println("Ведите число(до 255):")
	_, err := fmt.Scanln(&res)
	if err != nil {
		return
	}

	sb := ConvertIntToBits(res)
	fmt.Printf("Ваше число в двоичном виде: %v\n", sb) // [0 0 0 0 0 1 0 0]
	fmt.Println("Ведите индекс который хотите изменить от 0 до 7 с слева на право")
	_, err = fmt.Scanln(&index)
	if err != nil {
		return
	}

	if index < 0 || index > 8 {
		fmt.Println("Опа, давай заного")
		return
	}

	if sb[index] != 0 {
		sb[index] = 0
	} else {
		sb[index] = 1
	}
	fmt.Printf("Ваше измененное число в двоичном виде: %v\n", sb)
	res = ConvertBitsToInt(sb)
	fmt.Printf("Ваше число: %d\n", res)
	fmt.Println("ПОКА!!")
}
