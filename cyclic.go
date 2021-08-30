package main

import "fmt"

func cyclic_sort(nums []int, shortest_number int){
  index := 0;
  for(index < len(nums)){
    current := nums[index] - shortest_number;
    if (nums[current] != nums[index]){
      tmp := nums[current]
      nums[current] = nums[index]
      nums[index] = tmp
    }else{
      index += 1
    }
  }
}

func main(){
  nums := []int{5,4,3,2,1}
  cyclic_sort(nums, 1)
  fmt.Println("decl:", nums)
}
