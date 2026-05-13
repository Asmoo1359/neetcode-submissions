func hasDuplicate(nums []int) bool {
    arr := nums
	toSet := make(map[int]struct{})
	// Length of Array
	arrSize := len(arr)
	//set status of duplicate in array
	var status bool
	// Check if Element in array contains duplicate
	for _, num := range arr {
		toSet[num] = struct{}{}
	}
	// Length of toSet
	toSetSize := len(toSet)
	
	if arrSize != toSetSize {
		status = true
		fmt.Printf("%v \n contains duplicate and has status : %t \n", arr, status)
	} else {
		status = false
		fmt.Printf("%v \n does not contain duplicate and has status : %t \n", arr, status)
	}
	return status
}