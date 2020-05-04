package compress

import (
	"math"
	"strconv"
)

func FromHexString(str string) string {

	var commands = []string{"{\"or\":[", "{\"and\":[", "{\"not\":[", "\"wd\":[", "\"bhr\":[", "\"hr\":[", "\"bev\":[", "\"eev\":[", "\"bdate\":[", "\"edate\":[", "{", "", "", "", "", ","}

	if str[len(str)-1] == '0' {
		for i := len(str) - 1; i >= 0; i-- {
			if str[i] != '0' {
				str = str[0:i+1] + "K"
				break
			}
		}
	}
	//fmt.Println("decimals", str)
	var rec string
	var strLength = len(str)
	var reverseStr = reverse(str)
	var lastNumber = 0
	var lastCommandPosition = 0
	var par int64
	for i := 0; i < strLength; i++ {
		var command = string(str[i])
		value, _ := strconv.ParseInt(command, 16, 64)
		if lastNumber != 0 {
			i = strLength - 1
		}
		if string(str[i]) != "K" {
			switch command {
			case "0", "1", "2":
				rec += commands[str[i]-48]
				par++
				break
			case "3", "4", "5", "6", "7", "8", "9":
				rec += commands[value]
				lastCommandPosition = i
				var newI, re, newLastNumber = findNumbers(i, str)
				i = newI
				lastNumber = newLastNumber
				rec += re
				break
			case "a", "f":
				rec += commands[value]
				break
			case "d":
				pom := indexOf(reverseStr, "f", len(reverseStr)-i)
				if pom == -1 {
					pom = lastCommandPosition + 1
				} else {
					pom = int(math.Max(float64(strLength-pom), float64(lastCommandPosition+1)))
				}
				intNumber, _ := strconv.ParseInt(string(str[pom:i]), 13, 64)
				rec += strconv.FormatInt(intNumber, 10)

				i++
				var f = indexOf(str, "f", i)
				var d = indexOf(str, "d", i)
				if d == -1 {
					d = i + 3
				}
				var e = indexOf(str, "e", i)
				//fmt.Println("f", i, d)
				repeat, _ := strconv.ParseInt(string(str[i:d]), 13, 64)
				if (f < d && f != -1) || (e < d && e != -1) || d-i > 2 {
					repeat = 1
					i--
				} else {
					i = d
				}
				par -= repeat - 1
				for r := 0; r < int(repeat); r++ {
					rec += "]}"
				}
				rec += ","
				break
			case "e":
				pom := indexOf(reverseStr, "f", len(reverseStr)-i)
				if pom == -1 {
					pom = lastCommandPosition + 1
				} else {
					pom = int(math.Max(float64(strLength-pom), float64(lastCommandPosition+1)))
				}
				intNumber, _ := strconv.ParseInt(string(str[pom:i]), 13, 64)
				rec += strconv.FormatInt(intNumber, 10) + "],"
				break
			}
		}
		//fmt.Println(command, rec)
		if string(str[i]) == "K" {
			intNumber, _ := strconv.ParseInt(string(str[lastNumber:i]), 13, 64)
			rec += strconv.FormatInt(intNumber, 10)
			for p := 0; p <= int(par); p++ {
				rec += "]}"
			}
		}
	}
	return rec
}

func findNumbers(i int, str string) (int, string, int) {
	var numberPosition = -1
	var jsonString = ""
	var lastNumberPosition = 0
	var end = 6
	i++
	for {
		numberPosition = indexOf(str, "f", i)

		var d = indexOf(str, "d", i)
		var e = indexOf(str, "e", i)
		var c = 0
		if d != -1 {
			c = numberPosition
			if numberPosition == -1 {
				c = d
			}
			numberPosition = int(math.Min(float64(c), float64(d)))
		}
		if e != -1 {
			c = numberPosition
			if numberPosition == -1 {
				c = e
			}
			numberPosition = int(math.Min(float64(c), float64(e)))
		}
		if numberPosition != -1 {
			if str[numberPosition] == 'f' {
				intNumber, _ := strconv.ParseInt(string(str[i:numberPosition]), 13, 64)
				jsonString += strconv.FormatInt(intNumber, 10) + ","
				i = numberPosition + 1
			} else {
				i = numberPosition - 1
			}
		} else {
			lastNumberPosition = i
		}
		end -= 1

		if end < 0 {
			break
		}
		if numberPosition == -1 || str[numberPosition] != 'f' {
			break
		}
	}
	return i, jsonString, lastNumberPosition
}

func indexOf(str string, find string, from int) int {
	var findByte = find[0]
	for i := from; i < len(str); i++ {
		if str[i] == findByte {
			return i
		}
	}
	return -1
}

func reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}
