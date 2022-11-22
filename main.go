package main

import "fmt"

func main() {
	// test2()
	// testPointer()
	// conditioinal()
	// arrayTest()
	// arrayTest2()
	// multidimensiArray()
	// loopingArray()
	// sliceArray()
	// sliceArray2()
	// belajarMap()
	// belajarMap2()
	loopingMap()
}

func loopingMap() {
	// cara horizontal
	var chicken1 = map[string]int{"januari": 50, "februari": 40, "maret": 35, "april": 70}
	for key, value := range chicken1 {
		fmt.Println(key, "  \t", value)
	}
}

func belajarMap2() {
	// cara horizontal
	var chicken1 = map[string]int{"januari": 50, "februari": 40}

	// cara vertical
	var chicken2 = map[string]int{
		"januari":  50,
		"februari": 40,
	}

	fmt.Println(chicken1)
	fmt.Println(chicken2)
}

func belajarMap() {
	// var chicken map[string]int
	chicken := map[string]int{}

	chicken["januari"] = 50
	chicken["februari"] = 40

	fmt.Println("januari", chicken["januari"]) // januari 50
	fmt.Println("mei", chicken["mei"])         // mei 0
}

func sliceArray2() {
	var fruits = []string{"apple", "grape", "banana", "melon"}

	var aFruits = fruits[0:3]
	var bFruits = fruits[1:4]

	var aaFruits = aFruits[1:2]
	var baFruits = bFruits[0:1]

	fmt.Println(fruits)   // [apple grape banana melon]
	fmt.Println(aFruits)  // [apple grape banana]
	fmt.Println(bFruits)  // [grape banana melon]
	fmt.Println(aaFruits) // [grape]
	fmt.Println(baFruits) // [grape]

	// Buah "grape" diubah menjadi "pinnaple"
	baFruits[0] = "pinnaple"

	fmt.Println(fruits)   // [apple pinnaple banana melon]
	fmt.Println(aFruits)  // [apple pinnaple banana]
	fmt.Println(bFruits)  // [pinnaple banana melon]
	fmt.Println(aaFruits) // [pinnaple]
	fmt.Println(baFruits) // [pinnaple]
}

func sliceArray() {
	// array1 := [...][3]string{{"fauzan", "akmal", "mahdi"}, {"lagi", "belajar golang"}}
	array1 := [...]string{"fauzan", "akmal", "mahdi", "lagi", "belajar golang"}

	fmt.Println("panjang array\t:", len(array1))
	fmt.Println("isi array\t:", array1[0:3])
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
