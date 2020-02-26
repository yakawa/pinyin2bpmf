package main

import (
	"bufio"
	"os/exec"
	"fmt"
	"os"
	"strings"
)

func next0(s []rune) string {
	if len(s) < 1 {
		return ""
	}
	return strings.ToLower(string(s[0]))
}

func next1(s []rune) string {
	if len(s) < 2 {
		return ""
	}
	return strings.ToLower(string(s[1]))
}

func next2(s []rune) string {
	if len(s) < 3 {
		return ""
	}
	return strings.ToLower(string(s[2]))
}

func main() {
	sc := bufio.NewScanner(os.Stdin)

	for sc.Scan() {
		t := sc.Text()
		u := []rune(t)
		bpmf := ""
		for i := 0; i < len(u); i++ {
			s, n := convert(u[i:])
			bpmf += s
			i += n
		}
		fmt.Println(bpmf)
		err := exec.Command("/usr/bin/say", "-v", "Mei-Jia", bpmf ).Run()
		if err != nil {
			fmt.Println(err)
		}
	}
}

func convert(u []rune) (string, int) {
	bpmf:= ""
	i := 0
	switch strings.ToLower(string(u[i])) {
	case "a", "ā", "á", "ǎ", "à":
		i += 1
		c, n := caseA(next0(u[i:]), next1(u[i:]))
		n--
		i += n
		bpmf += c
		if next0(u[i:]) == "1" || next0(u[i:]) == "2" || next0(u[i:]) == "3" || next0(u[i:]) == "4"  {
			bpmf += "˙ "
		}
	case "u", "ū", "ú", "ǔ", "ù":
		i += 1
		c, n := caseU(next0(u[i:]), next1(u[i:]), next2(u[i:]))
		n--
		i += n
		bpmf += c
		if next0(u[i:]) == "1" || next0(u[i:]) == "2" || next0(u[i:]) == "3" || next0(u[i:]) == "4"  {
			bpmf += "˙ "
		}
	case "v", "ü", "ǖ", "ǘ", "ǚ",  "ǜ":
		i += 1
		c, n := caseV(next0(u[i:]), next1(u[i:]))
		n--
		i += n
		bpmf += c
		if next0(u[i:]) == "1" || next0(u[i:]) == "2" || next0(u[i:]) == "3" || next0(u[i:]) == "4"  {
			bpmf += "˙ "
		}
	case "i", "ī", "í", "ǐ", "ì":
		i += 1
		c, n := caseI(next0(u[i:]), next1(u[i:]), next2(u[i:]))
		n--
		i += n
		bpmf += c
		if next0(u[i:]) == "1" || next0(u[i:]) == "2" || next0(u[i:]) == "3" || next0(u[i:]) == "4"  {
			bpmf += "˙ "
		}
	case "o", "ō", "ó", "ǒ", "ò":
		i += 1
		c, n := caseO(next0(u[i:]), next1(u[i:]))
		n--
		i += n
		bpmf += c
		if next0(u[i:]) == "1" || next0(u[i:]) == "2" || next0(u[i:]) == "3" || next0(u[i:]) == "4"  {
			bpmf += "˙ "
		}
	case "e", "ē", "é", "ě", "è":
		i += 1
		c, n := caseE(next0(u[i:]), next1(u[i:]))
		n--
		i += n
		bpmf += c
		if next0(u[i:]) == "1" || next0(u[i:]) == "2" || next0(u[i:]) == "3" || next0(u[i:]) == "4"  {
			bpmf += "˙ "
		}
	case "y":
		i += 1
		c, n := caseY(next0(u[i:]), next1(u[i:]), next2(u[i:]))
		n--
		i += n
		bpmf += c
		if next0(u[i:]) == "1" || next0(u[i:]) == "2" || next0(u[i:]) == "3" || next0(u[i:]) == "4"  {
			bpmf += "˙ "
		}
	case "w":
		i += 1
		c, n := caseW(next0(u[i:]), next1(u[i:]), next2(u[i:]))
		n--
		i += n
		bpmf += c
		if next0(u[i:]) == "1" || next0(u[i:]) == "2" || next0(u[i:]) == "3" || next0(u[i:]) == "4"  {
			bpmf += "˙ "
		}

	case "b":
		bpmf += "ㄅ"
	case "p":
		bpmf += "ㄆ"
	case "m":
		bpmf += "ㄇ"
	case "f":
		bpmf += "ㄈ"
	case "d":
		bpmf += "ㄉ"
	case "t":
		bpmf += "ㄊ"
	case "n":
		bpmf += "ㄋ"
	case "l":
		bpmf += "ㄌ"
	case "g":
		bpmf += "ㄍ"
	case "k":
		bpmf += "ㄎ"
	case "h":
		bpmf += "ㄏ"
	case "j":
		bpmf += "ㄐ"
	case "q":
		bpmf += "ㄑ"
	case "x":
		bpmf += "ㄒ"
	case "r":
		bpmf += "ㄖ"

	case "z":
		if strings.ToLower(string(u[i+1])) == "h" {
			i++
			bpmf += "ㄓ"
		} else {
			bpmf += "ㄗ"
		}
	case "c":
		if strings.ToLower(string(u[i+1])) == "h" {
			i++
			bpmf += "ㄔ"
		} else {
			bpmf += "ㄘ"
		}
	case "s":
		if strings.ToLower(string(u[i+1])) == "h" {
			i++
			bpmf += "ㄕ"
		} else {
			bpmf += "ㄙ"
		}
	case "1":
		bpmf += " "
	case "2":
		bpmf += "ˊ "
	case "3":
		bpmf += "ˇ "
	case "4":
		bpmf += "ˋ "
	case " ":
		break
	default:
		fmt.Printf("Unknown Pinyin %s\n", string(u[i]))
		return "", 0
	}
	return bpmf, i
}

func caseA(c1, c2 string) (string, int) {
	if c1 == "n" && c2 == "g" {
		return "ㄤ", 2
	} else if c1 == "n" {
		return "ㄢ", 1
	} else if c1 == "o" {
		return "ㄠ", 1
	} else if c1 == "i" {
		return "ㄞ", 1
	} else {
		return "ㄚ", 0
	}
}

func caseE(c1, c2 string) (string, int) {
	if c1 == "n" && c2 == "g" {
		return "ㄥ", 2
	} else if c1 == "n" {
		return "ㄣ", 1
	} else if c1 == "i" {
		return "ㄟ", 1
	} else if c1 == "r" {
		return "ㄦ", 1
	} else {
		return "ㄜ", 0
	}
}

func caseI(c1, c2, c3 string) (string, int) {
	if c1 == "e" {
		return "ㄧㄝ", 1
	} else if c1 == "a" && c2 == "n" && c3 == "g" {
		return "ㄧㄤ", 3
	} else if c1 == "a" && c2 == "n" {
		return "ㄧㄢ", 2
	} else if c1 == "a" && c2 == "o" {
		return "ㄧㄠ", 2
	} else if c1 == "a" {
		return "ㄧㄚ", 1
	} else if c1 == "e" {
		return "ㄧㄝ", 1
	} else if c1 == "u" {
		return "ㄧㄡ", 1
	} else if c1 == "n" && c2 == "g" {
		return "ㄧㄥ", 2
	} else if c1 == "n" {
		return "ㄧㄣ", 1
	} else if c1 == "o" && c2 == "n" && c3 == "g" {
		return "ㄩㄥ", 3
	} else {
		return "ㄧ", 0
	}
}

func caseO(c1, c2 string) (string, int) {
	if c1 == "u" {
		return "ㄡ", 1
	} else if c1 == "n" && c2 == "g" {
		return "ㄨㄥ", 2
	} else {
		return "ㄛ", 0
	}
}

func caseV(c1, c2 string) (string, int) {
	if c1 == "e" {
		return "ㄩㄝ", 1
	} else if c1 == "a" && c2 == "n" {
		return "ㄩㄢ", 2
	} else if c1 == "n" {
		return "ㄩㄣ", 1
	} else {
		return "ㄩ", 0
	}
}

func caseU(c1, c2, c3 string) (string, int) {
	if c1 == "o" {
		return "ㄨㄛ", 1
	} else if c1 == "a" && c2 == "n" && c3 == "g" {
		return "ㄨㄤ", 3
	} else if c1 == "a" && c2 == "n" {
		return "ㄨㄢ", 2
	} else if c1 == "a" && c2 == "i" {
		return "ㄨㄞ", 2
	} else if c1 == "a" {
		return "ㄨㄚ", 1
	} else if c1 == "i" {
		return "ㄨㄟ", 1
	} else if c1 == "n" {
		return "ㄨㄣ", 1
	} else {
		return "ㄨ", 0
	}
}


func caseW(c1, c2, c3 string) (string, int) {
	if c1 == "o" {
		return "ㄨㄛ", 1
	} else if c1 == "a" && c2 == "n" && c3 == "g" {
		return "ㄨㄤ", 3
	} else if c1 == "a" && c2 == "n" {
		return "ㄨㄢ", 2
	} else if c1 == "a" && c2 == "i" {
		return "ㄨㄞ", 2
	} else if c1 == "a" {
		return "ㄨㄚ", 1
	} else if c1 == "e" && c2 == "i" {
		return "ㄨㄟ", 2
	} else if c1 == "e" && c2 == "n" && c3 == "g"{
		return "ㄨㄥ", 3
	} else if c1 == "e" && c2 == "n" {
		return "ㄨㄣ", 2
	} else {
		return "ㄨ", 0
	}
}

func caseY(c1, c2, c3 string) (string, int) {
	if c1 == "e" {
		return "ㄧㄝ", 1
	} else if c1 == "a" && c2 == "n" && c3 == "g" {
		return "ㄧㄤ", 3
	} else if c1 == "a" && c2 == "n" {
		return "ㄧㄢ", 2
	} else if c1 == "a" && c2 == "o" {
		return "ㄧㄠ", 2
	} else if c1 == "a" {
		return "ㄧㄚ", 1
	} else if c1 == "e" {
		return "ㄧㄝ", 1
	} else if c1 == "u" && c2 == "a" && c3 == "n" {
		return "ㄩㄢ", 3
	} else if c1 == "u" && c2 == "n" {
		return "ㄩㄣ", 2
	} else if c1 == "u" && c2 == "e" {
		return "ㄩㄝ", 3
	} else if c1 == "u" {
		return "ㄩ", 1
	} else if c1 == "i" && c2 == "n" && c3 == "g" {
		return "ㄧㄥ", 3
	} else if c1 == "i" && c2 == "n" {
		return "ㄧㄣ", 1
	} else if c1 == "o" && c2 == "n" && c3 == "g" {
		return "ㄩㄥ", 3
	} else if c1 == "o" && c2 == "u" {
		return "ㄧㄡ", 3
	} else {
		return "ㄧ", 0
	}
}
