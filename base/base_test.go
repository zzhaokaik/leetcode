package base

import (
	"errors"
	"fmt"
	"sync"
	"testing"
	"time"
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

func TestAA(t *testing.T){
	main1()
}


var mutxa sync.Mutex
var mutxb sync.Mutex
//死锁1
func main2(){

	//mutxa.Lock()
	//mutxb.Lock()
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		mutxa.Lock()
		time.Sleep(time.Second)
		fmt.Println("a")
		mutxb.Lock()
		wg.Done()

	}()

	go func() {
		mutxb.Lock()
		time.Sleep(time.Second)
		fmt.Println("b")
		mutxa.Lock()
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("ccc")

}


func TestSisuo(t *testing.T){

	main2()
}


func main3(){
	ch:=make(chan int)
	ch<-1
	a:= <- ch
	fmt.Println(a)


}


func main4(){
	ch:=make(chan int)
	ch<-1
	go func(){
		a:=<-ch
		fmt.Println(a)
	}()



}
func TestSiSuo2(t *testing.T){
	main3()
}