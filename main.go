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
		// As the old confs wrongly capitalized this key. Would
		// be fixed by WriteConf below, but we want the JSON
		// schema to not flag this error.
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
	origListLen := len(depthList)
	var changesList []bool
	changesList = make([]bool, origListLen, origListLen)
	changesList[0] = false
	for i := 1; i < origListLen; i++ {
		if depthList[i] > depthList[i-1] {
			changesList[i] = true
		} else {
			changesList[i] = false
		}
	}
	fmt.Println("Answer: ",countBigger(changesList)) //1346

	//var countBiggers []int
	//countBiggers = make([]int, origListLen, origListLen)
	//countBiggers[0] = 0
	//for i := 1; i < origListLen; i++ {
	//	if changesList[i] {
	//		countBiggers[i] = countBiggers[i-1] + 1
	//	} else {
	//		countBiggers[i] = countBiggers[i-1]
	//	}
	//}
	//print results:
	//for i := 1 ; i< origListLen; i++ {
	//	fmt.Printf("%d) depth %d , changed %t, measurements larger than the previous %d\n",
	//		i, depthList[i], changesList[i], countBiggers[i])
	//}
	//fmt.Println("Answer: ", countBiggers[origListLen-1]) //1301
}

func largerThanPrevSlidingWindow(depthList []int) {

	origListLen := len(depthList)

	var tripletSums []int
	//	tripletSums = make([]int, slidingWindowLen, slidingWindowLen)

	i_tripletSums := 0
	for i_depthList := 0; i_depthList <= origListLen-3; i_depthList++ {
		tripletSums = append(tripletSums, depthList[i_depthList]+depthList[i_depthList+1]+depthList[i_depthList+2])
		i_tripletSums++
	}

	var changesList []bool
	slidingWindowLen := len(tripletSums)
	changesList = make([]bool, slidingWindowLen, slidingWindowLen)
	changesList[0] = false
	for i := 1; i < slidingWindowLen; i++ {
		if tripletSums[i] > tripletSums[i-1] {
			changesList[i] = true
		} else {
			changesList[i] = false
		}
	}
	fmt.Println("Answer: ",countBigger(changesList)) //1346

	//var countBiggers []int
	//countBiggers = make([]int, slidingWindowLen, slidingWindowLen)
	//countBiggers[0] = 0
	//for i := 1; i < slidingWindowLen; i++ {
	//	if changesList[i] {
	//		countBiggers[i] = countBiggers[i-1] + 1
	//	} else {
	//		countBiggers[i] = countBiggers[i-1]
	//	}
	//}
	//for i := 1; i < slidingWindowLen; i++ {
	//	fmt.Printf("%d) depth sum %d changed: %t, measurements larger than the previous %d\n",
	//		i, tripletSums[i], changesList[i], countBiggers[i])
	//}
	//fmt.Println("Answer: ", countBiggers[slidingWindowLen-1]) //1346
}

func countBigger(changesList []bool) (int){
	var countBiggers []int

	len := len(changesList)
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