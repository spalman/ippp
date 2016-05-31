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

// атрибут ПО (элемент рабочей памяти)
type TRPObject struct {
	ident  string  // условный идентификатор атрибута в рамках МПО
	name   string  // реальное название атрибута ПО
	value  float32 // значение  атрибута ПО
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
func parseTxt(fileName string) {
	file, _ := os.Open(fileName)

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan() //should skip "#1"
	for scanner.Scan() {
		if scanner.Text() != "#2" {
			temp := strings.Split(scanner.Text(), ":")
			//Test things
			fmt.Printf("Temp:%s\n", temp)
			//(remove later)
			i, _ := strconv.Atoi(strings.Trim(temp[0], " "))
			RP[i] = TRPObject{ident: strings.Trim(temp[1], " "), name: strings.Trim(temp[2], " "), isCalc: false}
			//fmt.Printf("i:%d ident:%s, len:%d\n", i, RP[i].ident, len(RP[i].ident))
		} else {
			break
		}
	}
	for scanner.Scan() {
		tempObj := TModuleInterface{moduleIdent: "", moduleName: "", moduleOutputParam: "", moduleInputParam: ""}
		temp := strings.Split(scanner.Text(), ":")
		tempObj.moduleName = strings.Trim(temp[1], " ")
		//Test things
		//fmt.Printf("Temp:%s\n", temp)
		//(remove later)
		temp = strings.Split(temp[0], "=")
		//fmt.Printf("t0= %s", returnIndexInRP(temp[0]))
		tempObj.moduleOutputParam = returnIndexInRP(temp[0])
		//fmt.Printf("OUT:%s ", tempObj.moduleOutputParam)
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

type Stack struct {
	top  *Element
	size int
}

type Element struct {
	value interface{} // All types satisfy the empty interface, so we can store anything here.
	next  *Element
}

// Return the stack's length
func (s *Stack) Len() int {
	return s.size
}

// Push a new element onto the stack
func (s *Stack) Push(value interface{}) {
	s.top = &Element{value, s.top}
	s.size++
}

// Remove the top element from the stack and return it's value
// If the stack is empty, return nil
func (s *Stack) Pop() (value interface{}) {
	if s.size > 0 {
		value, s.top = s.top.value, s.top.next
		s.size--
		return
	}
	return nil
}

func Solver(s int) {
	if RP[s].isCalc {
		return
	}
	for i := range MI {
		outParam, _ := strconv.Atoi(MI[i].moduleOutputParam)
		if outParam == s {

			stack.Push(MI[i].moduleIdent)
			fmt.Printf("%s\n", MI[i].moduleIdent)
			for j := 0; j < len(MI[i].moduleInputParam); j++ {
				index, _ := strconv.Atoi(string(MI[i].moduleInputParam[j]))
				if !RP[index].isCalc {
					Solver(index)
				}
			}
			RP[s].isCalc = true
			return
		}
	}

}

func main() {
	parseTxt("1.txt")
	for i := range RP {
		fmt.Printf("%+v\n", RP[i])
	}
	for i := range MI {
		fmt.Printf("%+v\n", MI[i])
	}
	RP[9].isCalc = true
	RP[2].isCalc = true
	RP[4].isCalc = true
	Solver(8)
	for stack.Len() > 0 {
		// We have to do a type assertion because we get back a variable of type
		// interface{} while the underlying type is a string.
		fmt.Printf("%s ", stack.Pop().(string))
	}
}
