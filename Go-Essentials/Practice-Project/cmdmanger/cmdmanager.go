package cmdmanger

import "fmt"

type CMDManger struct {
}

func (cmd CMDManger) ReadLines() ([]string, error) {

	fmt.Println("Enter you Prices. Conform your prices with ENTER!")
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

func (cmd CMDManger) WriteResult(data interface{}) error {
	fmt.Println(data)
	return nil
}


func New()CMDManger{
	return CMDManger{}
}
