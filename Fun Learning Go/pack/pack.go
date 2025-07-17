package pack

import (
	"fmt"
	"math/rand"
)

// 到火星的速度
func Malacandra() {
	const (
		distance   = 56000000 //km
		nedtime    = 28       //days
		hourPerDay = 24       //hours
	)
	var num = distance / nedtime / hourPerDay //km/h
	fmt.Print(num)
}

// 猜数字
func Guess(n int) {
	if n < 1 || n > 100 {
		fmt.Println("请输入1-100之间的数字,包括1和100")
		fmt.Printf("you are a bad boy,you input: %v", n)
		return
	}
	var guessNum = rand.Intn(100) + 1
	var guessmin, guessmax = 0, 101 //猜测范围，开区间
	fmt.Println("猜一个数字")
	for {
		if guessNum == guessmax {
			guessNum--
		}
		fmt.Printf("猜测是%v，", guessNum)
		switch {
		case guessNum < n:
			fmt.Printf("猜小了!\n")
			guessmin = guessNum
			diff := guessmax - guessNum //计算差值
			guessNum = rand.Intn(diff) + guessNum + 1
		case guessNum > n:
			fmt.Printf("猜大了!\n")
			guessmax = guessNum
			diff := guessNum - guessmin //计算差值
			guessNum = rand.Intn(diff) + guessmin + 1
		case guessNum == n:
			fmt.Printf("猜对了!\n")
			fmt.Printf("Hi! it found your number: %v", n)
			return
		}
	}
}

// 随机日期
func Random_dates() {
	era := "AD"
	daysInMonth := 31
	for i := 0; i < 10; i++ {
		year := rand.Intn(5000) + 1
		month := rand.Intn(12) + 1
		switch month {
		case 4, 6, 9, 11:
			daysInMonth = 30
		case 2:
			if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
				daysInMonth = 29
			}
			daysInMonth = 28
		}
		day := rand.Intn(daysInMonth) + 1
		fmt.Println(era, year, month, day)
	}
}

// 生成太空航行票
func Ticket() {
	const (
		distance   = 62100000 // km
		sPerhour   = 60 * 60
		hourPerDay = 24
	)
	var year, month, day = 2020, 10, 13
	var company, gotype = "SpaceX", "单程"
	fmt.Printf("出发日期：%v/%v/%v\n", year, month, day)
	fmt.Printf("%-20s%-20s%-20s%-20s\n", "太空航行公司", "飞行天数", "飞行类型", "价格（百万美元）")
	for i := 0; i < 10; i++ {
		var speed = rand.Intn(15) + 16 //km/s
		var price = 50 - (30 - speed)  //百万美元
		var days = distance / (speed * sPerhour * hourPerDay)
		switch n := rand.Intn(3); n {
		case 0:
			company = "SpaceX"
		case 1:
			company = "Virgin Galactic"
		case 2:
			company = "Space Adventures"
		}
		switch n := rand.Intn(2); n {
		case 0:
			gotype = "单程"
		case 1:
			gotype = "往返"
			price *= 2
			days *= 2
		}
		fmt.Printf("%-26s%5v%23s%26v\n", company, days, gotype, price)
	}

}

// piggy储蓄
func Piggy() {
	var pig float64 = 0
	var pigint int = 0
	for pig < 20 {
		coin := 0
		switch n := rand.Intn(3); n {
		case 0:
			coin = 5
		case 1:
			coin = 10
		case 2:
			coin = 25
		}
		pigint += coin
		pig = float64(pigint) / 100.0 // 转换为元
		fmt.Printf("当前余额$%5.2f元\n", pig)
		// time.Sleep(time.Second) // 模拟存钱的延时
	}
}

// 大数计算
func Cains() {
	const lightYear = 299792458 * 60 * 60 * 24 * 365 / 1000 // m
	const distance = 236000000000000000                     // km
	const lightYears = distance / lightYear
	fmt.Println(lightYears)
}
