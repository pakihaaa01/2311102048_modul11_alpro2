package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Masukkan suara (pisahkan dengan spasi, akhiri dengan 0):")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	data := strings.Split(input, " ")

	totalSuara := 0
	validSuara := 0
	count := make([]int, 20)

	for _, value := range data {
		num, err := strconv.Atoi(value)
		if err != nil {
			continue 
		}
		if num == 0 {
			break 
		}
		totalSuara++
		if num >= 1 && num <= 20 {
			validSuara++
			count[num-1]++
		}
	}

	fmt.Printf("Total suara terbaca: %d\n", totalSuara)
	fmt.Printf("Total suara valid: %d\n", validSuara)
	for i, suara := range count {
		if suara > 0 {
			fmt.Printf("Calon nomor %d: %d suara\n", i+1, suara)
		}
	}
}
