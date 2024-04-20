package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"TugasTestingRG/helper" //import dari folder helper
)

//EXIT CONTROL C

func Multiplication(operand1 string, operand2 string) (string, error) {
	if len(operand1) == 0 || len(operand2) == 0 {
		return "", fmt.Errorf("operand1 or operand2 is empty")
	}
	num1, err := strconv.Atoi(operand1)
	if err != nil {
		return "", fmt.Errorf("operand1 is not a valid number")
	}
	num2, err := strconv.Atoi(operand2)
	if err != nil {
		return "", fmt.Errorf("operand2 is not a valid number")
	}
	result := float64(num1) * float64(num2)
	resultStr := fmt.Sprintf("%.2f", result)
	return "Hasil adalah : " + resultStr, nil

}

func Division(operand1 string, operand2 string) (string, error) {
	if len(operand1) == 0 || len(operand2) == 0 {
		return "", fmt.Errorf("operand1 or operand2 is empty")
	}
	num1, err := strconv.Atoi(operand1)
	if err != nil {
		return "", fmt.Errorf("operand1 is not a valid number")
	}
	num2, err := strconv.Atoi(operand2)
	if err != nil {
		return "", fmt.Errorf("operand2 is not a valid number")
	}
	result := float64(num1) / float64(num2)
	// fmt.Println(result)
	resultStr := fmt.Sprintf("%.2f", result)
	return "Hasil adalah : " + resultStr, nil

}

func Exponen(operand1 string, operand2 string) (string, error) {
	if len(operand1) == 0 || len(operand2) == 0 {
		return "", fmt.Errorf("operand1 or operand2 is empty")
	}
	num1, err := strconv.Atoi(operand1)
	if err != nil {
		return "", fmt.Errorf("operand1 is not a valid number")
	}
	num2, err := strconv.Atoi(operand2)
	if err != nil {
		return "", fmt.Errorf("operand2 is not a valid number")
	}
	result := math.Pow(float64(num1), float64(num2))
	resultStr := fmt.Sprintf("%.2f", result) //https://yourbasic.org/golang/convert-string-to-float/
	return "Hasil adalah : " + resultStr, nil

}

func Sqrt(operand1 string) (string, error) {
	if len(operand1) == 0 {
		return "", fmt.Errorf("operand is empty")
	}
	num1, err := strconv.Atoi(operand1)
	if err != nil {
		return "", fmt.Errorf("operand is not a valid number")
	}

	result := math.Sqrt(float64(num1))       //fungsi math mengaruskan input float
	resultStr := fmt.Sprintf("%.2f", result) //SUPAYA HNY 2 ANGKA DIBELAKANG KOMA  //https://yourbasic.org/golang/convert-string-to-float/
	return "Hasil adalah : " + resultStr, nil

}

func main() {
	for {
		helper.ClearScreen()
		fmt.Println("Selamat datang di Kalkulator Tersederhana :)!")
		fmt.Println("1. Perkalian")
		fmt.Println("2. Pembagian")
		fmt.Println("3. Perpangkatan")
		fmt.Println("4. Akar Bilangan")
		fmt.Println("5. Exit")

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Pilih menu: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			helper.ClearScreen() //membersihkan terminal
			fmt.Println("=== Perkalian ===")
			fmt.Print("Masukkan angka pertama: ")
			operand1, _ := reader.ReadString('\n')
			operand1 = strings.TrimSpace(operand1)

			fmt.Print("Masukkan angka kedua: ")
			operand2, _ := reader.ReadString('\n')
			operand2 = strings.TrimSpace(operand2)

			result, err := Multiplication(operand1, operand2)
			if err != nil {
				fmt.Printf("Error: %s\n", err.Error())
			}
			fmt.Println(result)
			helper.Delay(5) //menunggu agar pesan muncul terlebih dahulu

		case "2":
			helper.ClearScreen()
			fmt.Println("=== Pembagian ===")
			fmt.Print("Masukkan angka pertama: ")
			operand1, _ := reader.ReadString('\n')
			operand1 = strings.TrimSpace(operand1)

			fmt.Print("Masukkan angka kedua: ")
			operand2, _ := reader.ReadString('\n')
			operand2 = strings.TrimSpace(operand2)
			result, err := Division(operand1, operand2)
			if err != nil {
				fmt.Printf("Error: %s\n", err.Error())
			}
			fmt.Println(result)
			helper.Delay(5)
		case "3":
			helper.ClearScreen()
			fmt.Println("=== Perpangkatan ===")
			fmt.Print("Masukkan angka pertama: ")
			operand1, _ := reader.ReadString('\n')
			operand1 = strings.TrimSpace(operand1)

			fmt.Print("Masukkan angka kedua: ")
			operand2, _ := reader.ReadString('\n')
			operand2 = strings.TrimSpace(operand2)
			result, err := Exponen(operand1, operand2)
			if err != nil {
				fmt.Printf("Error: %s\n", err.Error())
			}
			fmt.Println(result)
			helper.Delay(5)

		case "4":
			helper.ClearScreen()
			fmt.Println("=== Akar Kuadrat ===")
			fmt.Print("Masukkan angka : ")
			operand1, _ := reader.ReadString('\n')
			operand1 = strings.TrimSpace(operand1)

			result, err := Sqrt(operand1)
			if err != nil {
				fmt.Printf("Error: %s\n", err.Error())
			}
			fmt.Println(result)
			helper.Delay(5)
		case "5":
			helper.ClearScreen()
			fmt.Println("Terima kasih telah berkunjung :)")
			return
		}
	}
}
