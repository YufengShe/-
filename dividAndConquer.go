package main

import (
	"fmt"
)

//分治思想：归并排序
func MergeSort(arr []int, left int, right int) {
	//递归边界
	if left >= right {
		return
	}
	//分解原问题
	mid := (left + right)/2

	//解决子问题
	MergeSort(arr, left, mid)
	MergeSort(arr, mid+1, right)

	//合并子问题的解得到原问题的解
	Merge(arr, left, mid, right)
}

//归并排序合并子问题的解
func Merge(arr []int, left int, mid int, right int)  {

	tempArr := make([]int, right+1)
	ptemp := left
	pleft := left
	pright := mid+1

	for  {

		if pleft > mid || pright > right {
			break
		}

		if arr[pleft] <= arr[pright] {
			tempArr[ptemp] = arr[pleft]
			pleft++
		} else {
			tempArr[ptemp] = arr[pright]
			pright++
		}

		ptemp++

	}

	for ; pleft<=mid; pleft++ {
		tempArr[ptemp] = arr[pleft]
		ptemp++
	}

	for ; pright<=right; pright++{
		tempArr[ptemp] = arr[pright]
		ptemp++
	}

	for i:=left; i<=right; i++  {
		arr[i] = tempArr[i]
	}

}

//分治思想：快排
func QuickSort(arr []int, left int, right int)  {

	//递归边界
	if left >= right {
		return
	}

	//分:寻找分界点

	//以最右元素作为基准值
	iqvt := arr[right]
	//数组划分 使得pleft左面的数组元素均小于分界点元素的值，右面元素均大于基准值(Partition)
	pleft := left-1
	pright := left


	//开始进行数组划分
	for ; pright<right; pright++ {
		if arr[pright] < iqvt{
			temp := arr[pright]
			arr[pright] = arr[pleft+1]
			arr[pleft+1] = temp
			pleft++
		}

	}
	arr[pright] = arr[pleft+1]
	arr[pleft+1] = iqvt

	//数组划分结束 此时iqvt完成了排序 坐标值为mid
	mid := pleft+1

	//治：解决子问题
	QuickSort(arr, left, mid-1)
	QuickSort(arr, mid+1, right)
}



//分治思想：逆序对计数问题
func ReverseCount(arr []int, left int, right int) int{
	//递归边界
	if left >= right {
		return 0
	}

	//分-治
	mid := (left + right)/2
	n1 := ReverseCount(arr, left, mid)
	n2 := ReverseCount(arr, mid+1, right)

	//合
	n3 := CrossingCount(arr, left, mid, right)
	n := n1 + n2 + n3
	return  n
}

//逆序对计数问题—过界逆序对计数
func CrossingCount(arr []int, left int, mid int, right int)  int {

	count := 0
	tempArr := make([]int, right+1)
	ptemp := left
	pleft := left
	pright := mid+1

	for ; pleft<=mid && pright<=right;  {
		if arr[pleft] <= arr[pright] {
			tempArr[ptemp] = arr[pleft]

			for i:=mid+1; i<pright; i++ {
				fmt.Println("reverse: ", arr[pleft], "——", arr[i])
			}
			count = count + pright - (mid+1)
			pleft++
		} else {
			tempArr[ptemp] =arr[pright]
			pright++
		}
		ptemp++
	}

	for ; pleft<=mid; pleft++ {
		tempArr[ptemp] = arr[pleft]

		for i:=mid+1; i<=right; i++ {
			fmt.Println("reverse: ", arr[pleft], "——", arr[i])
		}

		count = count + right - mid

		ptemp++
	}

	for ; pright<=right; pright++ {
		tempArr[ptemp] = arr[pright]
		ptemp++
	}

	for i:=left; i<=right; i++ {
		arr[i] = tempArr[i]
	}
	return count
}

//分治思想：最大子数组问题
func MaxSum(arr []int, left int, right int)  int{

	//递归边界
	if left >= right {
		return arr[left]
	}

	//分治
	mid := (left + right)/2
	s1 := MaxSum(arr, left, mid)
	s2 := MaxSum(arr, mid+1, right)

	//合
	s3 := MaxCrossing(arr, left, mid, right)
	max := Max(s1, s2, s3)
	return max
}

func Max(s1 int, s2 int, s3 int) int {
	max := s1
	if s2 > max {
		max = s2
	}

	if s3 > max {
		max = s3
	}

	return max
}

func MaxCrossing(arr []int, left int, mid int, right int)  int{

	//求左边以mid结尾的最大子数组和 右边以mid+1开头的最大子数组和 加起来就是crossMax值

	//MaxLeftArry
	maxleft := arr[mid]
	sumleft := arr[mid]
	for i:=mid-1; i>=left; i-- {
		sumleft := sumleft + arr[i]
		if sumleft > maxleft {
			maxleft = sumleft
		}
	}

	//MaxRightArry
	maxright := arr[mid+1]
	sumright := arr[mid+1]
	for i:=mid+2; i<=right; i++ {
		sumright := sumright + arr[i]
		if sumright > maxright {
			maxright = sumright
		}
	}

	return maxright+maxleft
}


//分治思想：次序选择问题(寻找第K小的元素)
func Select(arr []int, left int, right int, k int)  int{

	//递归边界
	if k > right+1 {
		return -1
	}

	//就是找落在k-1坐标上的元素并返回
	pivot := arr[right] //基准元素
	pleft := left-1
	pright := left

	//Partition
	for ; pright<right; pright++ {
		if arr[pright] < pivot {
			temp := arr[pright]
			arr[pright] = arr[pleft+1]
			arr[pleft+1] = temp
			pleft++
		}
	}
	arr[right] = arr[pleft+1]
	arr[pleft+1] = pivot

	//此时确定了pivot元素排好序的位置 下标为pleft+1
	mid := pleft+1
	if mid == k-1 {
		return pivot
	} else if mid < k-1 {
		return Select(arr, mid+1, right, k)
	} else {
		return Select(arr, left, mid-1, k)
	}

}




func main() {


	//归并排序
	arr1 := []int{-500,10,5,5,2,-3,-29,-10,-50,22,-23,10,150,1,9,-900,-26,3,99,7}
	MergeSort(arr1,0, len(arr1)-1)
	fmt.Println("MergeSort Result：")
	for _,v := range arr1{
		fmt.Printf("%d, ", v)
	}
	fmt.Print("\n")

	//快速排序
	arr2 := []int{-500,10,5,5,2,-3,-29,-10,-50,22,-23,10,150,1,9,-900,-26,3,99,7}
	QuickSort(arr2, 0, len(arr2)-1)
	fmt.Println("QuickSort Result：")
	for _,v := range arr2{
		fmt.Printf("%d, ", v)
	}
	fmt.Print("\n")

	//逆序对计数问题
	fmt.Println("ReverseCount Result：")
	arr3 := []int{-500,10,5,5,2,-3,-29,-10,-50,22,-23,10,150,1,9,-900,-26,3,99,7}
	count := ReverseCount(arr3, 0, len(arr3)-1)
	for _,v := range arr3{
		fmt.Printf("%d, ", v)
	}
	fmt.Println("")
	fmt.Printf("逆序对的个数为%d\n", count)

	//最大子数组和
	arr4 := []int{-500,10,5,5,2,-3,-29,-10,-50,22,-23,10,150,1,9,-900,-26,3,99,7}
	maxSum := MaxSum(arr4, 0, len(arr4)-1)
	fmt.Println("MaxSum Result: ")
	fmt.Println("最大子数组和为", maxSum)

	//次序选择问题
	arr5 := []int{-500,10,5,5,2,-3,-29,-10,-50,22,-23,10,150,1,9,-900,-26,3,99,7}
	k := 6
	elem := Select(arr5, 0, len(arr5)-1, 3)
	fmt.Println("Select Orderer:")
	fmt.Printf("第%d小的数为%d\n", k, elem)
}
