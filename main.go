package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	depthListFileName := "./depthList"
	file, err := os.Open(depthListFileName)
	if err != nil {
		fmt.Println("Error ", err.Error())
		return
	}
	defer file.Close()

	//depthList := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
	var perline int
	var depthList []int
	for {
		_, err = fmt.Fscanf(file, "%d\n", &perline) // give a patter to scan
		if err != nil {
			if err == io.EOF {
				break // stop reading the file
			}
			fmt.Println(err)
			os.Exit(1)
		}
		depthList = append(depthList, perline)
	}

	origListLen := len(depthList)
	if origListLen == 0 {
		fmt.Println("No data. How many measurements are larger than the previous measurement?", 0)
		return
	}
	if origListLen == 1 {
		fmt.Println("1 input. How many measurements are larger than the previous measurement?", 0)
		return
	}
	//fmt.Println("Data:", origList)

	//0 decreased
	//0 no prev
	// 1 increased
	largerThanPrev(depthList)
	largerThanPrevSlidingWindow(depthList)

}

func largerThanPrev(depthList []int) {

	var changesList []bool
	setChangesList(depthList, &changesList)
	fmt.Println("Answer: ",countBigger(changesList)) //1346
}

func largerThanPrevSlidingWindow(depthList []int) {

	origListLen := len(depthList)

	var tripletSums []int
	i_tripletSums := 0
	for i_depthList := 0; i_depthList <= origListLen-3; i_depthList++ {
		tripletSums = append(tripletSums, depthList[i_depthList]+depthList[i_depthList+1]+depthList[i_depthList+2])
		i_tripletSums++
	}

	var changesList []bool
	setChangesList(tripletSums, &changesList)
	fmt.Println("Answer: ",countBigger(changesList)) //1346
}

func setChangesList(numbersList []int, changesList *[]bool ){
	len := len(numbersList)
	*changesList = append (*changesList, false)
	for i := 1; i < len; i++ {
		if numbersList[i] > numbersList[i-1] {
			*changesList = append(*changesList,true)
		} else {
			*changesList = append(*changesList,false)
		}
	}
}

func countBigger(changesList []bool) (int){
	var countBiggers []int

	len := len(changesList)
	if len == 0 {
		return 0
	}
	countBiggers = make([]int, len, len)
	countBiggers[0] = 0
	for i := 1; i < len; i++ {
		if changesList[i] {
			countBiggers[i] = countBiggers[i-1] + 1
		} else {
			countBiggers[i] = countBiggers[i-1]
		}
	}
	return countBiggers[len-1]
}