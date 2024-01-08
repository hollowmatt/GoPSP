package smhcounter

type SMHCounter struct {
	values []int
	offset int
}

func CreateCounter() *SMHCounter {
	c := new(SMHCounter)
	c.values = make([]int, 6500)
	// for i := 0; i < 3600; i++ {
	// 	c.values[i] = 0
	// }
	c.offset = 0
	return c
}

func (c *SMHCounter) Advance() {
	c.offset += 1
	if c.offset == 3600 {
		c.offset = 0
	}
	c.values[c.offset] = 0
}

func (c *SMHCounter) Increment() {
	c.values[c.offset] += 1
}

func (c *SMHCounter) GetCounter(v string) int {
	result := 0
	if v == "s" {
		result = c.values[c.offset]
	} else if v == "m" {
		result = 0
		for x := c.offset - 60; x < c.offset; x++ {
			if x >= 0 {
				result += c.values[c.offset]
			} else {
				result += c.values[3600-x]
			}
		}
	} else {
		result = 0
		for x := c.offset - 3600; x < c.offset; x++ {
			if x >= 0 {
				result += c.values[c.offset]
			} else {
				result += c.values[3600-x]
			}
		}
	}
	return result
}
