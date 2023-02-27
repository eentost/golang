package ma

import "fmt"

func MakeNumIn() [3]int {
	// 0~9사이 겹차지 않는 숫자 3개 반환
	var rst [3]int
	return rst
}

func UserInputNum() [3]int {

	// 키보드로부터 0~9 사이의 겹치지 않는 숫자 3개 입력받아 반환
	var rst [3]int
	return rst
}

func CompareNum(MakeNumIn, UserInputNum [3]int) bool {
	// 두개의 숫자 3개를 비교해서 결과를 반환한다.
	return true
}

func PrintRsult(result bool) {
	fmt.Println(result)
}

func IsGameEnd(result bool) bool {
	//비교 결과가 3s 인지 확인
	return true
}

func main() {

	// 무작위 숫자 3개를 만든다.
	numbers := MakeNumIn()

	cnt := 0
	for {
		cnt++
		// 사용자의 입력을 받는다.
		inNumbers := UserInputNum()
		// 결과 비교
		result := CompareNum(numbers, inNumbers)
		PrintRsult(result)
		// 3S 인가?
		if IsGameEnd(result) {
			// 게임끝
			break
		}
	}
	// 몇번만에 맞췄는지 출력
	fmt.Printf("%d번 만에 맞췄습니다.\n", cnt)
}
