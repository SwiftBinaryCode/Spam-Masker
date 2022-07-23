//character set program

// package main

// import (
// 	"fmt"
// 	"unsafe"
// )

// func main() {

// 	//Section 8

// 	//string header
// 	// empty := ""
// 	// dump(empty)

// 	hello := "hello"
// 	dump(hello)
// 	dump("hello")
// 	dump("hello!")
// 	dump(hello[0:1])

// 	for i := range hello {
// 		dump(hello[i : i + 1])
// 	}

// 	dump(string([]byte(hello)))
// 	dump(string([]byte(hello)))
// 	dump(string([]rune(hello)))

// 	type StringHeader struct{
// 		//points to a backing arrays item
// 		pointer uintptr
// 		length int
// 	}

// 	//dump prints the string header of a string value
// 	func dump(s string) {
// 		ptr := *(*StringHeader)(unsafe.Pointer(&s))
// 		fmt.Printf("%q: %+v\n", s, ptr)
// 	}

// 	//Another example
// 	// word := []byte("öyku")

// 	// fmt.Printf("%s = % [1]x\n", word)

// 	// for i := range string(word) {

// 	// 	if i > 0 {
// 	// 		size = i
// 	// 		break
// 	// 	}
// 	// }

// 	// _, size := utf8.DecodeRune(word)

// copy(word[:size], bytes.ToUpper(word[:size]))
// fmt.Printf("%s = % [1]x\n", word)

//rune decoding automatically
// text := `hello my name is bob the builder`

// for _, r := range text {

// 	fmt.Printf("%c", r)
// }
// fmt.Println()

//Rune decoding manually
// text := `hello my name is bob the builder`

// for i := 0; i < len(text); {
// 	r, size := utf8.DecodeRuneInString(text[i:])
// 	fmt.Printf("%c", r)

// 	i += size
// }
// fmt.Println()

//converting
// str := "Yugen → →"

// bytes := []byte(str)
// str = string(bytes)

// fmt.Printf("%s\n", str)
// fmt.Printf("\t%d bytes\n", len(str))
// fmt.Printf("\t%d runes\n", utf8.RuneCountInString(str))

// fmt.Printf("% x\n", bytes)
// fmt.Printf("\t%d bytes\n", len(bytes))
// fmt.Printf("\t%d runes\n", utf8.RuneCount(bytes))

// fmt.Println()

// for i, r := range str {
// 	fmt.Printf("str[%2d] = %-12x = %q\n", i, string(r), r)

// }

// fmt.Println()
// fmt.Printf("1st byte      : %c\n", str[0])
// fmt.Printf("2nd byte      : %c\n", str[1])
// fmt.Printf("2nd rune      : %s\n", str[1:3])
// fmt.Printf("last rune      : %s\n", str[len(str)-4:])

// runes := []rune(str)

// fmt.Println()
// fmt.Printf("%s\n", string(runes))
// fmt.Printf("\t%d runes\n",
// 	int(unsafe.Sizeof(runes[0]))*len(runes))
// fmt.Printf("\t%d runes\n", len(runes))

// fmt.Printf("1st rune      : %c\n", runes[0])
// fmt.Printf("2nd rune     : %c\n", runes[1])
// fmt.Printf("first five     : %c\n", runes[:5])

//character set program
// var start, stop int

// if args := os.Args[1:]; len(args) == 2 {
// 	start, _ = strconv.Atoi(args[0])
// 	stop, _ = strconv.Atoi(args[1])
// }

// if start == 0 || stop == 0 {
// 	start, stop = 'A', 'Z'
// }

// fmt.Printf("%-10s %-10s %-10s %-12s\n%s\n",
// 	"literal", "dec", "hex", "encoded",
// 	strings.Repeat("-", 45))

// for n := start; n <= stop; n++ {

// 	fmt.Printf("%-10c => %-10[1]d %-10[1]x % -12x\n", n, string(n))

// }

// char := '→'
// _ = char
// }
