package main

import (
	"fmt"
	"reflect"
	"sort"
)

func main() {
	mySlice()

	// 使用sort包中已经实现了的Interface
	fmt.Println("\n使用sort包中已经实现了的Interface")
	myIntSlice := []int{1, 6, 2, 36, 8, 78, 6, 96}
	i := sort.IntSlice(myIntSlice)
	//i := sort.IntSlice{6, 9, 5, 7, 36, 15, 68, 6}
	fmt.Printf("type = %v, value = %v\n", reflect.TypeOf(i), i)
	sort.Sort(i)
	fmt.Printf("type = %v, value = %v\n", reflect.TypeOf([]int(i)), []int(i))

	// Float64s 方法按照升序的方法
	fmt.Println("\nFloat64s 方法按照升序的方法")
	f := []float64{36.02, 36.69, 22.0, 565.59, 2.0}
	fmt.Println(f)
	fmt.Println("是否已经排序 : ", sort.Float64sAreSorted(f))
	sort.Float64s(f)
	fmt.Println(f)
	fmt.Println("是否已经排序 : ", sort.Float64sAreSorted(f))
	fmt.Println("index = ", sort.SearchFloat64s(f, 2.0))

	// string 的升序排序
	fmt.Println("\nstring 的升序排序")
	s := []string{"asdf", "dfghdgf", "ADFsdfg", "SadsfDFa", "sfgdf65456"}
	fmt.Println(s)
	fmt.Println("是否已经排序 : ", sort.StringsAreSorted(s))
	sort.Strings(s)
	fmt.Println(s)
	fmt.Println("是否已经排序 : ", sort.StringsAreSorted(s))

	sort.Sort(sort.Reverse(sort.StringSlice(s)))
	fmt.Println(s)
	//sort.Search()
	//sort.SearchStrings(s, func(i int) bool {})
	//sort.Search()
}

type Student struct {
	name string
	sex  bool
	//...
}

type studetSlice []*Student

func (c studetSlice) Len() int {
	return len(c)
}

func (c studetSlice) Less(i, j int) bool {
	return c[i].name < c[j].name
}

func (c studetSlice) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c studetSlice) show() {
	for i := 0; i < c.Len(); i++ {
		fmt.Printf("%v ", c[i])
	}
}

func mySlice() {
	fmt.Println("自定义slice排序")
	var students studetSlice
	for i := 10; i >= 0; i-- {
		students = append(students, &Student{
			name: fmt.Sprintf("student %d ", i),
			sex:  true,
		})
	}

	fmt.Println("bofore, is sorted : ", sort.IsSorted(students))
	students.show()

	sort.Sort(students)

	fmt.Println("\nafter, is sorted : ", sort.IsSorted(students))
	students.show()
}
