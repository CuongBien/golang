package main // 1. Khai báo package

import "fmt" // 2. Import thư viện standard của Go

// 3. Hàm main là điểm bắt đầu của chương trình (giống C++)
func main() {
	c := Circle{
		radius: 2,
	}

	r := Rectangle{
		height: 2,
		width: 1,
	}

	PrintArea(r)
	PrintArea(c)

	fmt.Printf("Hello\n")
}