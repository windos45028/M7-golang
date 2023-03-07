package main

import (
	"encoding/json"
	"fmt"
	"math/big"
	"sort"
	"strconv"
	"strings"
)

type BetInfo struct {
	Odds        BetInfoOdds   `json:"odds"`
	Items       []string      `json:"items"`
	OddsOptions []BetInfoOdds `json:"odds_options,omitempty"`
	Label       string        `json:"label"`
}

type BetInfoOdds struct {
	Key   string  `json:"key"`
	Value float64 `json:"value"`
}

func main() {
	chek()
}

func chek() (err error) {
	// x, err := chek1()
	// if err != nil {
	// 	return
	// }

	// if y, err := chek2(); err != nil {
	// 	return
	// }

	// fmt.Println(x, y)
	m := [][]int{
		[]int{2, 2, 2, 2, 2, 2, 2},
		[]int{2, 0, 0, 0, 0, 0, 2},
		[]int{2, 0, 2, 0, 2, 0, 2},
		[]int{2, 0, 0, 2, 0, 2, 2},
		[]int{2, 2, 0, 2, 0, 2, 2},
		[]int{2, 0, 0, 0, 0, 0, 2},
		[]int{2, 2, 2, 2, 2, 2, 2},
	}
	ttt(m, 1, 1, 2, 5)

	var str string
	for _, v := range m {
		for _, v2 := range v {
			switch v2 {
			case 0:
				str += "   "
			case 1:
				str += " ◇ "
			case 2:
				str += " ▫ "
			}
		}
		str += "\n"
	}

	fmt.Println(m)
	fmt.Println(str)
	return
}

func ttt(m [][]int, i, j, end_i, end_j int) bool {

	if m[i][j] == 0 {
		fmt.Println(strconv.Itoa(i) + "-" + strconv.Itoa(j))
		m[i][j] = 1
		if !(m[end_i][end_j] == 1) && !(ttt(m, i, j+1, end_i, end_j) || ttt(m, i+1, j, end_i, end_j) || ttt(m, i, j-1, end_i, end_j) || ttt(m, i-1, j, end_i, end_j)) {
			m[i][j] = 0
		}
	}

	return m[end_i][end_j] == 1
}

func chek1() (string, error) {
	return "2", nil
}

func chek2() (string, error) {
	return "3", nil
}

// 取得台號各位置的號碼出現次數
func CombinationNumberCount(combinations []int) map[int]int {
	numberCount := make(map[int]int)

	for _, combination := range combinations {
		numberCount[combination]++
	}

	return numberCount
}

// 取得 賽果轉換台號 特殊玩法
func ResultCombination(result map[string]int, locates []string) ([]int, string) {
	var newResult []int
	var head string

	// 賽果先由小排到大
	resultAscNumber := ResultASCNumbers(result, locates)
	for index, value := range resultAscNumber {
		tail := strconv.Itoa(CalcTail(value))

		switch index {
		case 0:
			head = tail
		default:
			number, err := strconv.Atoi(head + tail)
			if err != nil {
				return nil, "errr"
			}
			newResult = append(newResult, number)
			head = tail
		}
	}

	return newResult, ""
}

// 取得賽果指定位置的升序排序獎號
func ResultASCNumbers(result map[string]int, locates []string) []int {
	numbers := ResultNumbers(result, locates)
	sort.Ints(numbers)

	return numbers
}

// 取得賽果指定位置的獎號
func ResultNumbers(result map[string]int, locates []string) []int {
	numbers := make([]int, 0)

	for _, locate := range locates {
		numbers = append(numbers, result[locate])
	}

	return numbers
}

func MapStringThanSize(data map[string]float64) string {
	var minOdds float64
	var minOddsKey string
	for oddsKey, odds := range data {
		if odds < minOdds || minOdds == 0 {
			minOdds = odds
			minOddsKey = oddsKey
		}
	}

	return minOddsKey
}

// items add string zero (used in string sorting)
func AddStringArrayPreixZero(items []string) []string {
	newItems := []string{}
	for _, item := range items {
		number, _ := strconv.Atoi(item)
		if number < 10 {
			newItems = append(newItems, "0"+strconv.Itoa(number))
		} else {
			newItems = append(newItems, item)
		}

	}

	return newItems
}

// items del string zero (used in string sorting)
func DelStringArrayPreixZero(items []string) []string {
	newItems := []string{}
	for _, item := range items {
		number, _ := strconv.Atoi(item)
		newItems = append(newItems, strconv.Itoa(number))
	}

	return newItems
}

func ItemsHitCount(hitItems []string, items []string) int {
	var hit int
	for _, val := range hitItems {
		if InStringSlice(items, val) {
			hit++
		}
	}

	return hit
}

func Combin(x, y int64) int64 {
	b := new(big.Int).Binomial(x, y)
	return b.Int64()
}

func RemoveIndex(s []string, index int) ([]string, string) {
	remain := s[index : index+1][0]
	return append(s[:index], s[index+1:]...), remain
}

// 計算尾數（個位數）
func CalcTail(number int) int {
	return number % 10
}

// 計算頭數（十位數）
func CalcHead(number int) int {
	return number / 10 % 10
}

// string slice 裡是否有該值
func InStringSlice(haystack []string, needle string) bool {
	for _, h := range haystack {
		if h == needle {
			return true
		}
	}

	return false
}

func A1() {
	items := []string{"1,2,3", "5,6,8", "9"}
	countItems := len(items)
	// items := []string{"1", "5", "9", "16", "19,20"}

	type detail struct {
		Size int
		Item []string
	}

	var itemOfBets, max int
	var ttt []detail
	for _, item := range items {
		str := strings.Split(item, ",")
		if len(str) > max {
			max = len(str)
		}
		ttt = append(ttt, detail{
			Size: len(str),
			Item: str,
		})
		itemOfBets += len(str)
	}

	var delItems []string
	for {
		var b bool
		if len(delItems) >= itemOfBets-6 {
			break
		}

		for k, c := range ttt {
			if c.Size > max {
				max = c.Size
			}
			if c.Size == max && !b {
				items, delItem := RemoveIndex(c.Item, c.Size-1)
				ttt[k] = detail{
					Item: items,
					Size: c.Size - 1,
				}

				delItems = append(delItems, delItem)
				max -= 1
				b = true
			}
		}
	}

	fmt.Println(delItems)

	s, _ := json.Marshal(ttt)
	var b string
	if len(delItems) > 0 {
		delItems, b = RemoveIndex(delItems, len(delItems)-1)
	}

	fmt.Println(delItems, b, string(s))

	tmp := make(map[string]string, 0)
	var c, d int
	for index := range items {
		if index >= countItems-1 {
			break
		}
		itemNl1 := strings.Split(items[index], ",")
		for index2 := range items {
			end := index + index2 + 1
			if end >= countItems {
				break
			}
			itemNl2 := strings.Split(items[end], ",")
			for _, nl1 := range itemNl1 {
				for _, nl2 := range itemNl2 {
					betinfoItems := []string{nl1, nl2}
					sort.Strings(betinfoItems)
					betinfoItem := strings.Join(betinfoItems, "")
					if _, ok := tmp[betinfoItem]; ok {
						continue
					}

					itemsNotHitCount := ItemsHitCount(betinfoItems, delItems)
					fmt.Println(betinfoItems, itemsNotHitCount)
					itemsHitSpCount := ItemsHitCount(betinfoItems, []string{b})
					// 二中二
					if itemsNotHitCount >= 1 {
						c++
					} else if itemsHitSpCount == 1 {
						d++
					}

					tmp[betinfoItem] = betinfoItem
				}
			}
		}
	}
	fmt.Println(len(tmp), c, d)
	fmt.Println("總碰=>", len(tmp))
	fmt.Println("二中二碰=>", len(tmp)-c-d)
	fmt.Println("二中特碰=>", d)
	fmt.Println("不中碰=>", c)
}

func A2() {
	items := []string{"1,2,3", "6,7,8"}
	countItems := len(items)
	// it := []string{"1", "2", "3", "4", "5"}
	// fmt.Println(Combin(5, 2))
	// betInfos := []obj.SubBetInfo{}
	type detail struct {
		Size int
		Item []string
	}

	var itemOfBets, max int
	var ttt []detail
	for k, item := range items {
		if k == 0 {
			continue
		}
		str := strings.Split(item, ",")
		if len(str) > max {
			max = len(str)
		}
		ttt = append(ttt, detail{
			Size: len(str),
			Item: str,
		})
		itemOfBets += len(str)
	}

	var delItems []string
	for {
		var b bool
		if len(delItems) >= itemOfBets-6 {
			break
		}

		for k, c := range ttt {
			if c.Size > max {
				max = c.Size
			}
			if c.Size == max && !b {
				items, delItem := RemoveIndex(c.Item, c.Size-1)
				ttt[k] = detail{
					Item: items,
					Size: c.Size - 1,
				}

				delItems = append(delItems, delItem)
				max -= 1
				b = true
			}
		}
	}

	var t []string
	var a int
	// vvv := map[string]int{}
	itemSp := strings.Split(items[0], ",")
	_, delSp := RemoveIndex(itemSp, len(itemSp)-1)
	fmt.Println(delSp)
	for index := range items {
		if index == 0 || index >= countItems-1 {
			continue
		}
		itemNl1 := strings.Split(items[index], ",")
		for index2 := range items {
			end := index + index2 + 1
			if end >= countItems {
				break
			}
			itemNl2 := strings.Split(items[end], ",")
			for _, sp := range itemSp {
				for _, nl1 := range itemNl1 {
					for _, nl2 := range itemNl2 {
						// 只排序正碼
						nlItems := []string{nl1, nl2}

						betinfoItems := []string{sp}
						betinfoItems = append(betinfoItems, nlItems...)
						if delSp == sp && nl1 != "8" && nl2 != "8" && nl1 != "13" && nl2 != "13" {
							a++
						}
						fmt.Println(betinfoItems)
						t = append(t, strings.Join(betinfoItems, ","))

					}
				}
			}
		}
	}
	fmt.Println(len(t), a, delItems)
}

func A3() {
	items := []string{"1,2,3", "6,7,8", "11,12", "13,14,15"}
	type detail struct {
		Size int
		Item []string
	}

	var itemOfBets, max, min, minLocation int
	for k, item := range items {
		str := strings.Split(item, ",")
		// fmt.Println(len(str), min, k, minLocation)
		if len(str) < min || min == 0 {
			min = len(str)
			minLocation = k
		}
	}
	newItems, spItem := RemoveIndex(items, minLocation)
	spItems := strings.Split(spItem, ",")
	_, delItem := RemoveIndex(spItems, len(spItems)-1)
	fmt.Println(delItem)
	var ttt []detail
	for _, item := range newItems {
		str := strings.Split(item, ",")
		if len(str) > max {
			max = len(str)
		}

		ttt = append(ttt, detail{
			Size: len(str),
			Item: str,
		})
		itemOfBets += len(str)
	}

	var delItems []string
	for {
		var b bool
		if len(delItems) >= itemOfBets-6 {
			break
		}

		for k, c := range ttt {
			if c.Size > max {
				max = c.Size
			}
			if c.Size == max && !b {
				items, delItem := RemoveIndex(c.Item, c.Size-1)
				ttt[k] = detail{
					Item: items,
					Size: c.Size - 1,
				}

				delItems = append(delItems, delItem)
				max -= 1
				b = true
			}
		}
	}
	fmt.Println(delItems)
}

func A4() {
	items := []string{"RAT", "PIG", "OX", "DOG", "B", "C", "D", "E", "F", "G"}
	oddsTable := map[string]float64{
		"RAT": 50.2,
		"PIG": 60.2,
		"OX":  70.2,
		"DOG": 80.2,
		"B":   40.2,
		"C":   30.2,
		"D":   30.2,
		"E":   30.2,
		"F":   30.2,
		"G":   20.2,
	}
	type kv struct {
		Key   string
		Value float64
	}
	var ss []kv
	for k, v := range oddsTable {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	var getItems []string
	for _, kv := range ss {
		fmt.Printf("%s, %d\n", kv.Key, kv.Value)
		if InStringSlice(items, kv.Key) && len(getItems) < 7 {
			getItems = append(getItems, kv.Key)
		}
	}

	fmt.Println(getItems)

	// if len(items) > 7 {
	// 	// for _, item :=
	// } else {

	// }
}
