package baeksolution

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type card struct {
	color string
	num   int
}

//Solution0 No.2621
func Solution2621() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	hand := make([]card, 5)
	for i := 0; i < 5; i++ {
		fmt.Fscanln(reader, &hand[i].color, &hand[i].num)
	}

	// 스트레이트플러시 높은숫자 900 - 스트레이트 같은색 o
	// 마운틴 숫자4개같음 같은숫자800 - 같은숫자 4
	// 풀하우스 3숫자*10 2숫자 700 - 같은숫자 3 같은숫자 2
	// 플러시 높은숫자 600 - 같은색
	// 스트레이트 높으숫자 500 - 스트레이트 o
	// 트리플 같은숫자 400 - 같은숫자 3
	// 투페어 큰수*10 작은수 300 - 같은숫자 2 2
	// 원페어 같은수 200 - 같은숫자 2
	// 가장큰수 100

	// 낚시 가능성도 항상 생각해야 한다! 예를들어 높은 숫자가 떠서 상위 족보를 이기는 경우?

	sort.SliceStable(hand, func(i, j int) bool {
		return hand[i].num < hand[j].num
	})

	res := 0

	flush := isFlush(hand)
	straight := isStraight(hand)
	if straight {
		if flush {
			// 스트플
			res += 900 + getHigh(hand)
		} else {
			//스트레이트
			res += 500 + getHigh(hand)
		}
	} else {
		nums := countNums(hand)
		pairs := [][]int{}
		// 숫자 i 가 v번 나온다
		for i, v := range nums {
			if v >= 2 {
				pairs = append(pairs, []int{i, v})
			}
		}
		if len(pairs) < 2 {
			pairs = append(pairs, []int{0, 1})
		}
		if len(pairs) < 2 {
			pairs = append(pairs, []int{0, 1})
		}
		if pairs[0][1] < pairs[1][1] {
			pairs = [][]int{pairs[1], pairs[0]}
		}

		if pairs[0][1] == 4 {
			// mountain
			res += pairs[0][0] + 800
		} else if pairs[0][1] == 3 {
			if pairs[1][1] == 2 {
				// full
				res += 700 + pairs[0][0]*10 + pairs[1][0]
			} else {
				// triple
				res += 400 + pairs[0][0]
			}
		} else if pairs[0][1] == 2 {
			if pairs[1][1] == 2 {
				//two pair
				if pairs[0][0] > pairs[1][0] {
					res += 300 + pairs[0][0]*10 + pairs[1][0]
				} else {
					res += 300 + pairs[1][0]*10 + pairs[0][0]
				}
			} else {
				//one pair
				res += 200 + pairs[0][0]
			}
		} else {
			if flush {
				//flush
				res += 600 + getHigh(hand)
			} else {
				//high
				res += 100 + getHigh(hand)
			}
		}
	}
	fmt.Fprint(writer, res)
}

func isFlush(hand []card) bool {
	color := hand[0].color
	for i := 1; i < 5; i++ {
		if color != hand[i].color {
			return false
		}
	}
	return true
}

func isStraight(hand []card) bool {
	low := hand[0].num
	for i := 1; i < 5; i++ {
		if low+i != hand[i].num {
			return false
		}
	}
	return true
}

func countNums(hand []card) []int {
	l := make([]int, 10)
	for _, v := range hand {
		l[v.num]++
	}
	return l
}

func getHigh(hand []card) int {
	high := hand[0].num
	for _, v := range hand {
		if high < v.num {
			high = v.num
		}
	}
	return high
}
