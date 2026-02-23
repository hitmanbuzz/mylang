package lexer

import (
	"mylang/utils"
)

func (l *Lexer) scanComment() {
	for {
		if l.isAtEnd() {
			break
		}

		if l.peek() == '\n' {
			l.Line++
			l.advance(1)
			break
		}
		l.advance(1)
	}
}

func (l *Lexer) scanString() ([]byte, bool) {
	isString := false
	var str []byte
	l.advance(1)

	for {
		if l.isAtEnd() {
			break
		}

		if l.peek() == '"' {
			isString = true
			l.advance(1)
			break
		}

		if l.peek() == '\n' {
			l.Line++
		}

		str = append(str, l.peek())
		l.advance(1)
	}

	return str, isString
}

func (l *Lexer) scanNumber() []byte {
	var nums []byte
	isDot := false
	isNumberStart := false

	for {
		if l.isAtEnd() || l.peek() == '\n' {
			break
		}

		nextByte := l.Source[l.CurrIdx+1]

		if utils.IsNum(l.peek()) {
			nums = append(nums, l.peek())
			isNumberStart = true
		} else if l.peek() == '.' && utils.IsNum(nextByte) && !isDot && isNumberStart {
			isDot = true
			nums = append(nums, l.peek())
		} else {
			isNumberStart = false
			break
		}

		if l.peek() == '\n' {
			l.Line++
		}

		l.advance(1)
	}

	lastB := nums[len(nums)-1]
	if lastB != '.' {
		nums = append(nums, '.')
		nums = append(nums, '0')
	}

	return nums
}
