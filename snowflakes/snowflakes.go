package snowflakes

import "fmt"

// https://dmoj.ca/problem/cco07p2
// TLE

type snowflake struct {
	arms [6]int
}

func equals(s1, s2 *snowflake) bool {
	m := make(map[int]bool, 6)
	for _, v := range s1.arms {
		m[v] = true
	}

	for _, v := range s2.arms {
		_, ok := m[v]
		if !ok {
			return false
		}
	}
	return true
}

func (s snowflake) calculateHash() int {
	sum := 0
	for _, v := range s.arms {
		sum += v
	}
	return sum
}

func newSnowflake(arm [6]int) *snowflake {
	return &snowflake{arms: arm}
}

type Snowflakes struct {
	flakes map[int]*snowflake
	number int
	twins  bool
}

func Solve() {
	var s Snowflakes
	fmt.Scanf("%d", &s.number)
	s.twins = false
	s.flakes = make(map[int]*snowflake, s.number)

	for i := 0; i < s.number; i++ {
		var currentArm [6]int
		for j := 0; j < 6; j++ {
			fmt.Scan(&currentArm[j])
		}

		f := newSnowflake(currentArm)
		hash := f.calculateHash()
		val, exists := s.flakes[hash]
		if !exists {
			s.flakes[hash] = f
		} else {
			twins := equals(val, f)
			if twins {
				fmt.Println("Twin snowflakes found.")
				s.twins = true
				break
			}
		}
	}
	if !s.twins {
		fmt.Println("No two snowflakes are alike.")
	}
}
