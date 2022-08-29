package _80__Insert_Delete_GetRandom_O_1_

import "math/rand"

type RandomizedSet struct {
	nums       []int
	valToIndex map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		valToIndex: make(map[int]int),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.valToIndex[val]; ok {
		return false
	}
	this.valToIndex[val] = len(this.nums)
	this.nums = append(this.nums, val)
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.valToIndex[val]; !ok {
		return false
	}

	index := this.valToIndex[val]
	lastVal := this.nums[len(this.nums)-1]
	this.nums[index] = lastVal
	this.nums = this.nums[:len(this.nums)-1]
	this.valToIndex[lastVal] = index
	delete(this.valToIndex, val)

	return true
}

func (this *RandomizedSet) GetRandom() int {
	n := len(this.nums)
	return this.nums[rand.Intn(n)]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
