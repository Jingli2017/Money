package _62__Design_Hit_Counter_

type hit struct {
	count int
	time  int
}

type HitCounter struct {
	deque []hit
	total int
}

func Constructor() HitCounter {
	deque := make([]hit, 0)
	total := 0
	return HitCounter{
		deque: deque,
		total: total,
	}
}

func (this *HitCounter) Hit(timestamp int) {
	if len(this.deque) != 0 && this.deque[len(this.deque)-1].time == timestamp {
		lasthit := this.deque[len(this.deque)-1]
		lasthit.count++
		this.deque = this.deque[:len(this.deque)-1]
		this.deque = append(this.deque, lasthit)
	}else{
		this.deque = append(this.deque, hit{1, timestamp})
	}
	this.total++
}

func (this *HitCounter) GetHits(timestamp int) int {
	for len(this.deque) != 0 && timestamp - this.deque[0].time >= 300 {
		fronthit := this.deque[0]
		this.deque = this.deque[1:]
		this.total -= fronthit.count
	}
	return this.total
}
