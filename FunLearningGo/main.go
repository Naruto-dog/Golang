package main

import (
	"fmt"
	"myGo/pack"
)

func main() {
	fmt.Println("Hello, Go")
	// pack.Malacandra()
	// pack.Guess(55) // 这里可以传入[1-100]之间的数字
	// pack.Random_dates() // 随机日期
	// pack.Ticket() // 太空旅行票
	// pack.Piggy() // 猪猪存钱罐
	// pack.Cains() // 用常量计算大数
	// pack.Rot13("Ubyn Rfgnpvóa Rfcnpvny Vagreanpvbany") // rot13加密
	// pack.Caesar("L fdph,L vdz,L frqtxhuhg.") // 凯撒解密
	// pack.Decipher("ECFRZKYGLGRMUSDHRXK")
	// pack.Cipher("your message goes here")
	// func() {
	// 	c := pack.Celsius(125)
	// 	f := pack.Fahrenheit(223)
	// 	k := pack.Kelvin(356)
	// 	ctf := c.ToFahrenheit()
	// 	ctk := c.ToKelvin()
	// 	ftc := f.ToCelsius()
	// 	ftk := f.ToKelvin()
	// 	ktc := k.ToCelsius()
	// 	ktf := k.ToFahrenheit()
	// 	fmt.Printf("%.2f ℃ Celsius to Fahrenheit: %.2f℉\n", c, ctf)
	// 	fmt.Printf("%.2f ℃ Celsius to Kelvin: %.2fK\n", c, ctk)
	// 	fmt.Printf("%.2f ℉ Fahrenheit to Celsius: %.2f℃\n", f, ftc)
	// 	fmt.Printf("%.2f ℉ Fahrenheit to Kelvin: %.2fK\n", f, ftk)
	// 	fmt.Printf("%.2f K Kelvin to Celsius: %.2f℃\n", k, ktc)
	// 	fmt.Printf("%.2f K Kelvin to Fahrenheit: %.2f℉\n", k, ktf)
	// }() // 温度转换
	// func() {
	// 	pack.DrawTable(pack.Ctof)
	// 	pack.DrawTable(pack.Ftoc)
	// }() // 温度转换表
	// pack.Chess()
	// fmt.Println(pack.Newearth.Terraform())
	// pack.Capacity()
	pack.Words()
}
