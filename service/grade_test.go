package service_test

import (
	"fmt"
	"testing"
	"unittest/service"

	"github.com/stretchr/testify/assert"
	// "golang.org/x/tools/cmd/godoc"
)

// go test ./...  | test ทุกตัวทุกที่
// go test unittest/service -v | เทส ทั้งหมดใน package service
//go test unittest/service -v -run=Test_Hello | เทสบาง Func
// ไฟล์ test มี error แต่  Go จะยังรันได้ *****
// ใส  -cover จะช้วยบอกว่า test cover กี่ %  | coverage: 42.9% of statements

// ******** go test -v คำสั่ง run ********
//go test "unittest/service" -v (go test ตามด้วยที่อยู่ func test)
// func สำหรับ test จะอยู๋ใน service_test ก็ตาม มันจะทำงานได้

func Test_CheckGrade(t *testing.T) {
	fmt.Println("")

	// inpotGrade:= []int{80,70,60,50,40}
	// expectedGrade:= []string{"A","B","C","D","F"}

	// 	for i,v:=range inpotGrade {
	// 		grade := service.CheckGrade(v)
	// 		if grade != expectedGrade[i] {
	// 			t.Errorf("got %v expected %v", grade, expectedGrade[i])
	// 		}
	// 	}

	// ******** ทำ sub test ****
	// test sub test ตามชื่อ ที่ต้องการ เทส เท่านั้น
	//**** go test unittest/service -v -run="Test_CheckGrade/A" | -run="ชื่อ Func test และ ชื่อ sub test" ******
	// t.Run("ชื่อ sub test")

	// t.Run("A",func(t *testing.T) {
	// 	grade := service.CheckGrade(80)
	// 	expected := "A"
	// 	if grade !=expected {
	// 		t.Errorf("got %v expected %v", grade, expected)
	// 	}
	// })

	// t.Run("B",func(t *testing.T) {
	// 	grade := service.CheckGrade(78)
	// 	expected := "B"
	// 	if grade !=expected {
	// 		t.Errorf("got %v expected %v", grade, expected)
	// 	}
	// })

	type testCase struct {
		name     string
		score    int
		expected string
	}

	// ********* ใช้ loop ช้วยในการ test ********
	cases := []testCase{
		{name: "a", score: 80, expected: "A"},
		{name: "b", score: 70, expected: "B"},
		{name: "c", score: 60, expected: "C"},
		{name: "d", score: 50, expected: "D"},
		{name: "f", score: 40, expected: "F"},
	}

	for _, v := range cases {
		t.Run(v.name, func(t *testing.T) {
			grade := service.CheckGrade(v.score)

			// if grade != v.expected {
			// 	t.Errorf("got %v expected %v", grade, v.expected)
			// }

			// ใช้ lib github.com/stretchr/testify ช้วย
			assert.Equal(t,v.expected,grade)
		})
	}

}

// ถ้ามีหลาย test ให้ใช้คำสั่ง run และตามด้วนชื่อ Func
// go test unittest/service -v -run=Test_Hello
func Test_Hello(t *testing.T) {
	fmt.Println("Func Test")
}

// ***********  Benchmark ********
// go test unittest/service -bench=Benchmark_CheckGrade | Benchmark เฉพาะชื่อ
// go test unittest/service -bench=. | Benchmark ทั้งหมด ใส .


// *********** check mem  *********
// go test unittest/service -bench=. -benchmem 
func Benchmark_CheckGrade(b *testing.B) {

	for i := 0; i < b.N; i++ {
		service.CheckGrade(80)
	}

}

func Example_CheckGrade() {
	grade := service.CheckGrade(80)
	fmt.Println(grade) // Output: A
}
