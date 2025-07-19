package pack

import (
	"fmt"
	"math/rand"
	"strings"
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

// rot13加密解密
func Rot13(s string) {
	// message := "uv vagreangvbany fcnpr fgngvba" // 示例字符串
	for _, v := range s {
		if v >= 'a' && v <= 'z' {
			v += 13
			if v > 'z' {
				v -= 26
			}
		}
		if v >= 'A' && v <= 'Z' {
			v += 13
			if v > 'Z' {
				v -= 26
			}
		}
		fmt.Printf("%c", v)
	}
}

// 解密凯撒密码
func Caesar(s string) {
	for i := 0; i < len(s); i++ {
		a := s[i]
		if a >= 'a' && a <= 'z' {
			a -= 3
			if a < 'a' {
				a += 26
			}
		}
		if a >= 'A' && a <= 'Z' {
			a -= 3
			if a < 'A' {
				a += 26
			}
			fmt.Printf("%c", a)
		}
	}
}

// 维吉尼亚解密(默认大写密文，关键字为GOLANG)
func Decipher(s string) {
	cipherText := s // "CSOITEUIWUIZNSROCNKFD"
	keyword := "GOLANG"
	for i, c := range cipherText {
		j := i % len(keyword)           // 计算当前字符对应的关键字字符索引
		diff := rune(keyword[j]) - 'A'  // 计算偏移量
		res := (c-diff-'A'+26)%26 + 'A' // 解密公式
		fmt.Printf("%c\n", res)         // 输出解密后的字符
	}
}

// 维吉尼亚加密
func Cipher(s string) {
	cipherText := s // "your message goes here"
	keyword := "GOLANG"
	cipherText = strings.Replace(cipherText, " ", "", -1) // 去除空格
	cipherText = strings.ToUpper(cipherText)              // 转为大写
	// fmt.Printf("Ciphertext: %s\n", cipherText)
	for i := 0; i < len(cipherText); i++ {
		c := cipherText[i]
		j := i % len(keyword)              // 计算当前字符对应的关键字字符索引
		diff := rune(keyword[j]) - 'A'     // 计算偏移量
		res := (rune(c)+diff-'A')%26 + 'A' // 加密公式
		fmt.Printf("%c\n", res)            // 输出加密后的字符
	}
}

// ////////////////////////////////////////////////////////////////////
// 温度转换
type Celsius float64
type Fahrenheit float64
type Kelvin float64

// 为celsius类型定义两个转换方法 ℃
func (c Celsius) ToFahrenheit() Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}
func (c Celsius) ToKelvin() Kelvin {
	return Kelvin(c + 273.15)
}

// 为fahrenheit类型定义两个转换方法 ℉
func (f Fahrenheit) ToCelsius() Celsius {
	return Celsius((f - 32) * 5 / 9)
}
func (f Fahrenheit) ToKelvin() Kelvin {
	return Kelvin((f-32)*5/9 + 273.15)
}

// 为kelvin类型定义两个转换方法 K
func (k Kelvin) ToCelsius() Celsius {
	return Celsius(k - 273.15)
}
func (k Kelvin) ToFahrenheit() Fahrenheit {
	return Fahrenheit((k-273.15)*9/5 + 32)
}

////////////////////////////////////////////////////////////////

// ////////////////////////////////////////////////////////////////////////
// 温度表
func draw(n int) {
	for i := 0; i < n; i++ {
		fmt.Print("=")
	}
	fmt.Print("\n")
}

func Ctof(n int) {
	draw(n)
	fmt.Printf("|℃        |℉        |\n")
	draw(n)
	for i := -40.0; i <= 101.0; i += 5 {
		fmt.Printf("|%-9.2f", i)
		fmt.Printf("|%-9.2f|\n", i*9/5+32)
	}
	draw(n)
}
func Ftoc(n int) {
	draw(n)
	fmt.Printf("|℉        |℃        |\n")
	draw(n)
	for i := -40.0; i <= 101.0; i += 5 {
		fmt.Printf("|%-9.2f", i)
		fmt.Printf("|%-9.2f|\n", (i-32)*5/9)
	}
	draw(n)
}
func DrawTable(fn func(int)) {
	fn(22)
}

/////////////////////////////////////////////////////////////////////

// 棋盘
func Chess() {
	var board = [8][8]string{}
	board[0][0] = "R" // 车
	board[0][1] = "N" // 马
	board[0][2] = "B" // 象
	board[0][3] = "Q" // 后
	board[0][4] = "K" // 王
	board[0][5] = "B" // 象
	board[0][6] = "N" // 马
	board[0][7] = "R" // 车
	board[7][0] = "r" // 车
	board[7][1] = "n" // 马
	board[7][2] = "b" // 象
	board[7][3] = "q" // 后
	board[7][4] = "k" // 王
	board[7][5] = "b" // 象
	board[7][6] = "n" // 马
	board[7][7] = "r" // 车

	for c := range board[1] {
		board[1][c] = "P" // 兵
		board[6][c] = "p" // 兵
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == "" && (i+j)%2 == 0 {
				board[i][j] = "·"
			} else if board[i][j] == "" && (i+j)%2 != 0 {
				board[i][j] = "-"
			}
			fmt.Printf("|%s", board[i][j])
		}
		fmt.Println("|")
		fmt.Println("----------------")
	}
}

// 行星地球化
type Planets []string

var Newearth = Planets{
	"Mars",
	"Uranus",
	"Neptune",
}

func (s Planets) Terraform() Planets {
	for i := range s {
		s[i] += "New"
	}
	return s
}

//////////////////////////////////////////////////////////////////

// append增加数组容量
func Capacity() {
	var words = [1]string{"whatbehind"}
	newwords := words[:]
	slice := []string{}
	for i := 0; i < 5; i++ {
		slice = append(slice, "itsme", "itsyou")
		fmt.Printf("Slice:%s\n", slice)
		newwords = append(newwords, slice...)
		fmt.Printf("newwords:%s\n len:%v--cap:%v\n", newwords, len(newwords), cap(newwords))
		fmt.Println("-----------------------------------------")
	}
	fmt.Println(words, len(words), cap(words))
}

// 统计单词出现评率
func Words() {
	var paragraph string = `
				As far as eye could reach he saw nothing but the stems of the great plants about him
				receding in the violet shade,andfar overhead the multiple transparency of huge leaves
				filtering the sunshine to the solemn splendour of twilight in which he walked.Whenever
				he felt able he ran again;the ground continued soft and springy,covered with the same
				resilient weed which was the first thing his hands had touched in Malacandra. Once or
				twice a small red creature scuttled across his path,but otherwise there seemed to be 
				no life stirring in the wood;nothing to fear—except the fact of wandering unprovisioned 
				and alone in a forest of unknown vegetation thousands or millions of miles beyond the 
				reach or knowledge of man
				`
	frequency := make(map[string]int)
	paragraph = strings.ToLower(paragraph)
	var word = ""
	for i := 0; i < len(paragraph); i++ {
		switch paragraph[i] {
		case ',', ' ', '\n', '.', '-', ';', '\t':
			frequency[word]++
			word = ""
		default:
			word = word + string(paragraph[i])

		}
	}
	delete(frequency, "")
	fmt.Printf("一共有%v个单词，大于一个的单词如下：\n", len(frequency))
	fmt.Printf("单词          频次\n")
	fmt.Printf("=================\n")
	for k, v := range frequency {
		if v > 1 && k != "" {
			fmt.Printf("%-15s%-2d\n", k, v)
		}
	}
}
