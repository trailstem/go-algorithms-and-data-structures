package main

import (
	"fmt"
	"time"
)

func main() {
	// 1. 変数宣言と初期化
	fmt.Println("### 1. 変数宣言と初期化 ###")
	var a int = 10
	b := 20
	var c, d string = "Hello", "World"
	fmt.Println("Values:", a, b, c, d)

	// 2. 定数
	fmt.Println("\n### 2. 定数 ###")
	const Pi float64 = 3.14159
	fmt.Println("Value of Pi:", Pi)

	// 3. 基本データ型
	fmt.Println("\n### 3. 基本データ型 ###")
	var i int = 42
	var f float64 = 3.14
	var s string = "go language"
	var k bool = true
	fmt.Println("Integer:", i, "Float:", f, "String:", s, "Boolean:", k)

	// 4. 配列
	fmt.Println("\n### 4. 配列 ###")
	var arr [3]int = [3]int{1, 2, 3}
	fmt.Println("Array:", arr)

	// 5. スライス
	fmt.Println("\n### 5. スライス ###")
	slice := []int{1, 2, 3}
	slice = append(slice, 4)
	fmt.Println("Slice:", slice)

	// 6. マップ
	fmt.Println("\n### 6. マップ ###")
	m := map[string]int{"alice": 25, "bov": 30}
	m["charlie"] = 35
	fmt.Println("Map:", m)

	age, exists := m["alice"]
	if exists {
		fmt.Println("Alice's age:", age, "Exists:", exists)
	}

	// 7. 条件分岐
	fmt.Println("\n### 7. 条件分岐 ###")
	x := 4
	if x > 5 {
		fmt.Println("x is greater than 5")
	} else {
		fmt.Println("x is less than or equal to 5")
	}

	// 8. switch文
	fmt.Println("\n### 8. Switch文 ###")
	switch x {
	case 10:
		fmt.Println("x is 10")
	default:
		fmt.Println("x is not 10")
	}

	// 9. forループ
	fmt.Println("\n### 9. Forループ ###")
	for i := 0; i < 5; i++ {
		fmt.Println("Counter:", i)
	}

	// 10. rangeを使用したループ
	fmt.Println("\n### 10. Rangeを使用したループ ###")
	nums := []int{1, 2, 3, 4, 5}
	for i, v := range nums {
		fmt.Println("Index:", i, "Value:", v)
	}

	// 11. 関数の呼び出し
	fmt.Println("\n### 11. 関数の呼び出し ###")
	fmt.Println("Addition result:", add(2, 3))

	quotient, remainder := divide(5, 2)
	fmt.Println("Quotient:", quotient, "Remainder:", remainder)

	// 12. 構造体とメソッド
	fmt.Println("\n### 12. 構造体とメソッド ###")
	p := Person{"Alice", 25}
	fmt.Println("Person:", p)
	p = Person{"Bob", 30}
	p.Greet()

	// 13. ポインタ
	fmt.Println("\n### 13. ポインタ ###")
	num := 5
	increment(&num)
	fmt.Println("Incremented value:", num)

	// 14. ゴルーチン
	fmt.Println("\n### 14. ゴルーチン ###")
	// ゴルーチンで並行実行
	go task1()
	go task2()
	// メイン関数が終了しないように待機
	time.Sleep(2 * time.Second)
	fmt.Println("Main Function Finished")

	// 15. チャネル
	fmt.Println("\n### 15. チャネル ###")
	ch := make(chan int)
	go func() {
		time.Sleep(10000000000)
		ch <- 42
	}()
	val := <-ch
	fmt.Println("Value received from channel:", val)
}

func add(a int, b int) int {
	return a + b
}

func divide(a, b int) (int, int) {
	return a / b, a % b
}

type Person struct {
	Name string
	Age  int
}

func (p Person) Greet() {
	fmt.Println("Hello", p.Name)
}

func increment(x *int) {
	*x = *x + 1
}

func task1() {
	for i := 1; i <= 5; i++ {
		fmt.Println("Task 1 - Step", i)
		time.Sleep(200 * time.Millisecond) // 200ms待機
	}
}

func task2() {
	for i := 1; i <= 5; i++ {
		fmt.Println("Task 2 - Step", i)
		time.Sleep(300 * time.Millisecond) // 300ms待機
	}
}
