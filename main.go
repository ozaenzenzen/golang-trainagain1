package main

import "fmt"

func main() {
	// test2()
	// testPointer()
	// conditioinal()
	// arrayTest()
	// arrayTest2()
	// multidimensiArray()
	loopingArray()
}

func loopingArray() {
	array1 := [...][3]string{{"fauzan", "akmal", "mahdi"}, {"lagi", "belajar golang"}}
	// array1 := [...]string{"fauzan", "akmal", "mahdi", "lagi", "belajar golang"}

	for i := 0; i < len(array1); i++ {
		fmt.Println(array1[i])
	}

	for i, arrays1 := range array1 {
		fmt.Printf("elemen %d : %s\n", i, arrays1[i])
	}
	// fmt.Println("panjang array\t:", len(array1))
	// fmt.Println("isi array\t:", array1)
}

func multidimensiArray() {
	array1 := [...][3]string{{"fauzan", "akmal", "mahdi"}, {"lagi", "belajar golang"}}
	fmt.Println("panjang array\t:", len(array1))
	fmt.Println("isi array\t:", array1)
}

func arrayTest2() {
	array1 := [...]string{"fauzan", "akmal", "mahdi", "lagi", "belajar golang"}
	fmt.Println("panjang array\t:", len(array1))
	fmt.Println("isi array\t:", array1)
}

func arrayTest() {
	array1 := [3]string{"fauzan", "akmal", "mahdi"}
	fmt.Println(len(array1))
	fmt.Println(array1)
}

func test1() {
	// fmt.Println("Hello World")
	// fmt.Println("Isi pesannya: apa")
	var firstName string = "Fauzan"

	var lastName string
	lastName = "Akmal Mahdi"

	fmt.Printf("halo %s %s!\n", firstName, lastName)
}

func test2() {

	var firstName string = "Fauzan"
	lastName := "Akmal Mahdi"

	fmt.Printf("halo %s %s!\n", firstName, lastName+"OK")
}

func testPointer() {
	name := new(string)

	fmt.Println(*name)
}

func conditioinal() {
	var point = 8840.0

	if percent := point / 100; percent >= 100 {
		fmt.Printf("%.1f%s perfect!\n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s good\n", percent, "%")
	} else {
		fmt.Printf("%.1f%s not bad\n", percent, "%")
	}
}
