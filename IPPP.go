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

func parseTxt(fileName string) {
	file, _ := os.Open(fileName)

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan() //should skip "#1"
	for scanner.Scan() {
		if scanner.Text() != "#2" {
			temp := strings.Split(scanner.Text(), ":")
			i, _ := strconv.Atoi(temp[0])
			baseTerms[i] = TRPObject{ident: temp[1], name: temp[2], isCalc: false}
		} else {
			break
		}
	}
	for scanner.Scan() {
		temp := strings.Split(scanner.Text(), ":")

	}
	/*
		for scanner.Scan() {
			temp := strings.Split(scanner.Text(), ":")
			i, _ := strconv.Atoi(temp[0])
			j, _ := strconv.Atoi(temp[1])
			k, _ := strconv.Atoi(temp[2])
			adjacencyMatrix[i][k] = j
		}
	*/
}
