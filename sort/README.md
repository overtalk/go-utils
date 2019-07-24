> Golang 中的 sort
>
> sort包中实现了３种基本的排序算法：插入排序．快排和堆排序．和其他语言中一样，这三种方式都是不公开的，他们只在sort包内部使用．所以用户在使用sort包进行排序时无需考虑使用那种排序方式，sort.Interface定义的三个方法：获取数据集合长度的Len()方法、比较两个元素大小的Less()方法和交换两个元素位置的Swap()方法，就可以顺利对数据集合进行排序。sort包会根据实际数据自动选择高效的排序算法

## 核心介绍

- interface Interface
---
```bash
// A type, typically a collection, that satisfies sort.Interface can be
// sorted by the routines in this package. The methods require that the
// elements of the collection be enumerated by an integer index.
type Interface interface {
	// Len is the number of elements in the collection.
	Len() int
	// Less reports whether the element with
	// index i should sort before the element with index j.
	Less(i, j int) bool
	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
}
```

- func Sort(data Interface)
---
```bash
// Sort sorts data.
// It makes one call to data.Len to determine n, and O(n*log(n)) calls to
// data.Less and data.Swap. The sort is not guaranteed to be stable.
func Sort(data Interface) {
	n := data.Len()
	quickSort(data, 0, n, maxDepth(n))
}
```
- func Reverse(data Interface) Interface
    - 实现 Interface 的降序排序
---
```bash
// Reverse returns the reverse order for data.
func Reverse(data Interface) Interface {
    return &reverse{data}
} 
```

## 用法介绍

- 自定义slice进行排序的时候将 目标slice 定义成结构体，实现sort.Interface中的方法，即可以调用sort.Sort()方法进行升序排序，demo如下：
---
```bash
package main

import (
	"fmt"
	"sort"
)

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

func main() {
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
```

## 常用方法介绍
```bash
// 排序方法
func Sort(data Interface)                   // 将Interface进行排序
func IsSorted(data Interface) bool          // 检查Interface是否已经排序
func Float64s(a []float64)                  // 将[]float64按照升序排序
func Float64sAreSorted(a []float64) bool    // 检查[]float64是否已经排序
func Ints(a []int)                          // 将[]int以升序排序
func IntsAreSorted(a []int) bool            // 检查[]int是否已经排序
func Strings(a []string)                    // 将[]string以升序排序
func StringsAreSorted(a []string) bool      // 检查[]string是否已经排序

// 查找方法
func Search(n int, f func(int) bool) int        // 二分法遍历[0, n), 找到最小满足f的数
func SearchFloat64s(a []float64, x float64) int // 查找 x 在 a 中的index, 如果不存在, 则返回插入之后的index
func SearchInts(a []int, x int) int             // 查找 x 在 a 中的index, 如果不存在, 则返回插入之后的index
func SearchStrings(a []string, x string) int    // 查找 x 在 a 中的index, 如果不存在, 则返回插入之后的index

// 通用slice排序
func Slice(slice interface{}, less func(i, j int) bool)                 // slice 必须为slice，否则会panic；less为Less函数；不稳定排序；
func SliceIsSorted(slice interface{}, less func(i, j int) bool) bool    // 判断slice是否已经排序了
func SliceStable(slice interface{}, less func(i, j int) bool)           // Slice函数的稳定排序版本
```

## 常用slice封装

### IntSlice

- sort包封装好的int数组排序工具，核心函数如下所示：
---
```bash
// IntSlice attaches the methods of Interface to []int, sorting in increasing order.
type IntSlice []int
    // implement of Interface
    func (p IntSlice) Len() int           { return len(p) }
    func (p IntSlice) Less(i, j int) bool { return p[i] < p[j] }
    func (p IntSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
    
    func (p IntSlice) Sort() { Sort(p) }
    
    // Search returns the result of applying SearchInts to the receiver and x.
    func (p IntSlice) Search(x int) int { return SearchInts(p, x) }
    
    // SearchInts searches for x in a sorted slice of ints and returns the index
    // as specified by Search. The return value is the index to insert x if x is
    // not present (it could be len(a)).
    // The slice must be sorted in ascending order.
    //
    func SearchInts(a []int, x int) int {
    	return Search(len(a), func(i int) bool { return a[i] >= x })
    }
```

### Float64Slice

- sort包封装好的float64数组排序工具，核心函数如下所示：
---
```bash
// Float64Slice attaches the methods of Interface to []float64, sorting in increasing order
// (not-a-number values are treated as less than other values).
type Float64Slice []float64
    // implement of Interface
    func (p Float64Slice) Len() int           { return len(p) }
    func (p Float64Slice) Less(i, j int) bool { return p[i] < p[j] || isNaN(p[i]) && !isNaN(p[j]) }
    func (p Float64Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
    
    // Sort is a convenience method.
    func (p Float64Slice) Sort() { Sort(p) }

    // Search returns the result of applying SearchFloat64s to the receiver and x.
    func (p Float64Slice) Search(x float64) int { return SearchFloat64s(p, x) }
        
    // SearchFloat64s searches for x in a sorted slice of float64s and returns the index
    // as specified by Search. The return value is the index to insert x if x is not
    // present (it could be len(a)).
    // The slice must be sorted in ascending order.
    //
    func SearchFloat64s(a []float64, x float64) int {
    	return Search(len(a), func(i int) bool { return a[i] >= x })
    }
```

### StringSlice

- sort包封装好的string数组排序工具，核心函数如下所示：
---
```bash
// StringSlice attaches the methods of Interface to []string, sorting in increasing order.
type StringSlice []string
    // implement of Interface
    func (p StringSlice) Len() int           { return len(p) }
    func (p StringSlice) Less(i, j int) bool { return p[i] < p[j] }
    func (p StringSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
    
    // Sort is a convenience method.
    func (p StringSlice) Sort() { Sort(p) }

    // Search returns the result of applying SearchStrings to the receiver and x.
    func (p StringSlice) Search(x string) int { return SearchStrings(p, x) }
        
    // SearchStrings searches for x in a sorted slice of strings and returns the index
    // as specified by Search. The return value is the index to insert x if x is not
    // present (it could be len(a)).
    // The slice must be sorted in ascending order.
    //
    func SearchStrings(a []string, x string) int {
        return Search(len(a), func(i int) bool { return a[i] >= x })
    }
```