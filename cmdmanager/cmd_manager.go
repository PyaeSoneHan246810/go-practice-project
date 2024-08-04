package cmdmanager

import "fmt"

type CMDManager struct {
}

func New() *CMDManager {
	return &CMDManager{}
}

func (cmdManager CMDManager) ReadLines() ([]string, error) {
	fmt.Println("Please enter your prices. Enter 0 to stop.")
	var prices []string
	for {
		var price string
		fmt.Print("Price: ")
		fmt.Scan(&price)
		if price == "0" {
			break
		}
		prices = append(prices, price)
	}
	return prices, nil
}

func (cmdManager CMDManager) WriteData(data any) error {
	fmt.Println(data)
	return nil
}
