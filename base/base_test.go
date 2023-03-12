package base

import (
	"errors"
	"fmt"
	"testing"
)

func TestPanic(t *testing.T){
	err:=test()
	fmt.Println(err)
}


// err 如果在方法名显示命名则defer 可以修改
// err 如果在函数中单独定义 则不会
func test()(err error){
	//var err error
	err=errors.New("bbbb")
	defer func(){
		if r:=recover();r!=nil{
			err=fmt.Errorf("aaaa")
		}
	}()
	rasePainc()
	return err
}
func rasePainc(){
	panic("错误")
}


func TestLoop(t *testing.T){
	a:=[]int{1,2,3}
	for k,v:=range a{
		if k==0{
			a[0],a[1]=100,200
			fmt.Println(a)
		}
		a[k] = v+100

	}
	fmt.Println(a)
}

func TestLoop2(t *testing.T){
	a:=[3]int{1,2,3}
	for k,v:=range a{
		if k==0{
			a[0],a[1]=100,200
			fmt.Println(a)
		}
		a[k] = v+100

	}
	fmt.Println(a)
}

func TestCopy(t *testing.T){

		slice1 := []int{1, 2, 3, 4, 5}
		slice2 := make([]int, 5)
		fmt.Printf("slice1: %v, %p\n", slice1, slice1)
	fmt.Printf("slice2: %v, %p\n", slice2, slice2)
		copy(slice2, slice1)
		fmt.Printf("slice2: %v, %p\n", slice2, slice2)
		slice3 := make([]int, 0, 5)
		for _, v := range slice1 {
			slice3 = append(slice3, v)
		}
		fmt.Printf("slice3: %v, %p\n", slice3, slice3)

//
//slice1: [1 2 3 4 5], 0xc0000b0030
//	slice2: [1 2 3 4 5], 0xc0000b0060
//	slice3: [1 2 3 4 5], 0xc0000b0090

}