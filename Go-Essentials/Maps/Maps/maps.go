package maps

import "fmt"

func main() {
	websites := map[string]string{
		"Google":              "https://google.com",
		"Amazon Web Services": "https://aws.com",
	}

	fmt.Println(websites)

	fmt.Println(websites["Amazon Web Services"])
	fmt.Println(websites["Google"])
	
	websites["LinkedIn"] = "https://linkedin.com"
	fmt.Println(websites)
	fmt.Println(websites["LinkedIn"])

	delete(websites, "Google")
	fmt.Println(websites)


}