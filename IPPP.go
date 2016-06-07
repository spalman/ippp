package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var MI [50]TModuleInterface //массив  описаний интерфейсов модулей
var RP [50]TRPObject        //массив элементов рабочей памяти пакета
var stack Stack
var checkingArray [10]int

// атрибут ПО (элемент рабочей памяти)
type TRPObject struct {
	ident  string  // условный идентификатор атрибута в рамках МПО
	name   string  // реальное название атрибута ПО
	value  float64 // значение  атрибута ПО
	isCalc bool    // признак: вычислен или нет
}

type TModuleInterface struct {
	moduleIdent       string // условный идентификатор модуля
	moduleName        string // смысловое название модуля
	moduleOutputParam string // номера выходных параметров
	moduleInputParam  string // номера входных параметров
}

func returnIndexInRP(s string) string {
	//fmt.Printf("s:%s l:%d", s, len(s))
	for i := range RP {
		if RP[i].ident == s {
			return strconv.Itoa(i)
		}
	}
	return ""
}
func contains(e int) bool {
	for _, a := range checkingArray {
		if a == e {
			return true
		}
	}
	return false
}
func arrayPush(e int) {
	for i := range checkingArray {
		if checkingArray[i] == -1 {
			checkingArray[i] = e
			return
		}
	}
}
func clearArr() {
	for i := range checkingArray {
		checkingArray[i] = -1
	}
}
func parseTxt(fileName string) {
	file, _ := os.Open(fileName)

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan() //should skip "#1"
	for scanner.Scan() {
		if scanner.Text() != "#2" {
			temp := strings.Split(scanner.Text(), ":")
			i, _ := strconv.Atoi(strings.Trim(temp[0], " "))
			RP[i] = TRPObject{ident: strings.Trim(temp[1], " "), name: strings.Trim(temp[2], " "), isCalc: false}
		} else {
			break
		}
	}
	for scanner.Scan() {
		tempObj := TModuleInterface{moduleIdent: "", moduleName: "", moduleOutputParam: "", moduleInputParam: ""}
		temp := strings.Split(scanner.Text(), ":")
		tempObj.moduleName = strings.Trim(temp[1], " ")
		temp = strings.Split(temp[0], "=")
		tempObj.moduleOutputParam = returnIndexInRP(temp[0])
		temp = strings.Split(temp[1], "(")
		tempObj.moduleIdent = strings.Trim(temp[0], " ")
		temp = strings.Split(temp[1], ",")
		temp[len(temp)-1] = temp[len(temp)-1][:len(temp[len(temp)-1])-2]
		for i := range temp {
			tempObj.moduleInputParam += returnIndexInRP(temp[i])
		}
		index, _ := strconv.Atoi(tempObj.moduleIdent[1:])
		MI[index] = tempObj
	}

}

func Solver(s int) {
	cont := false

	if RP[s].isCalc {
		return
	}
	for i := range MI {
		cont = false
		outParam, _ := strconv.Atoi(MI[i].moduleOutputParam)
		if outParam == s {
			for j := 0; j < len(MI[i].moduleInputParam); j++ {
				index, _ := strconv.Atoi(string(MI[i].moduleInputParam[j]))
				if !RP[index].isCalc {
					cont = true
					break
				}
			}
			if !cont {
				RP[s].isCalc = true
				stack.Push(MI[i].moduleIdent)
				return
			}
		}

	}
	for i := range MI {
		outParam, _ := strconv.Atoi(MI[i].moduleOutputParam)
		arrayPush(s)
		if outParam == s {

			stack.Push(MI[i].moduleIdent)
			fmt.Printf("%s\n", MI[i].moduleIdent)
			for j := 0; j < len(MI[i].moduleInputParam); j++ {
				index, _ := strconv.Atoi(string(MI[i].moduleInputParam[j]))
				if contains(index) {
					stack.Pop()
					cont = true
					break
				}
				if !RP[index].isCalc {
					cont = false
					Solver(index)
				}
			}
			if cont {
				continue
			}
			RP[s].isCalc = true
			return
		}

	}

}
func parseRequest(fileName string) float64 {
	file, _ := os.Open(fileName)

	defer file.Close()
	clearArr()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() != "?" {
			s := strings.Split(scanner.Text(), "=")
			for i := range RP {
				if RP[i].ident == s[0] {
					value, _ := strconv.ParseFloat(s[1], 64)
					RP[i].value = value
					RP[i].isCalc = true
				}
			}
		} else {
			break
		}
	}
	scanner.Scan()
	rez, _ := strconv.Atoi(returnIndexInRP(scanner.Text()))
	Solver(rez)
	calcRes()
	return RP[rez].value
}
func calcRes() {
	for stack.Len() > 0 {
		switch stack.Pop() {
		case "F1":
			fmt.Printf("F1: %s\n", MI[1].moduleName)
			F1()
		case "F2":
			fmt.Printf("F2: %s\n", MI[2].moduleName)
			F2()
		case "F3":
			fmt.Printf("F3: %s\n", MI[3].moduleName)
			F3()
		case "F4":
			fmt.Printf("F4: %s\n", MI[4].moduleName)
			F4()
		case "F5":
			fmt.Printf("F5: %s\n", MI[5].moduleName)
			F5()
		case "F6":
			fmt.Printf("F6: %s\n", MI[6].moduleName)
			F7()
		case "F7":
			fmt.Printf("F7: %s\n", MI[7].moduleName)
			F7()
		case "F8":
			fmt.Printf("F8: %s\n", MI[8].moduleName)
			F8()
		case "F9":
			fmt.Printf("F9: %s\n", MI[9].moduleName)
			F9()
		case "F10":
			fmt.Printf("F10: %s\n", MI[10].moduleName)
			F10()
		case "F11":
			fmt.Printf("F11: %s\n", MI[11].moduleName)
			F11()
		case "F12":
			fmt.Printf("F12: %s\n", MI[12].moduleName)
			F12()
		}
	}
}

func main() {
	parseTxt("data/1.txt")
	fmt.Printf("%f", parseRequest("data/input.txt"))
}
