package main

import (
	piscinegotests "piscine-go-tests"
	"piscine-go-tests/lib/challenge"
	"piscine-go-tests/lib/random"
	"strconv"
	"strings"
)

func main() {
	piscinegotests.ClearBinaries()

	operatorsTable := []string{"+", "-", "*", "/", "%"}

	table := []string{}
	firstArg := strconv.Itoa(random.IntBetween(-1000, 1000))
	secondArg := strconv.Itoa(random.IntBetween(0, 1000))
	for _, operator := range operatorsTable {
		table = append(table, firstArg+" "+operator+" "+secondArg)
	}

	table = append(table, "1 + 1")
	table = append(table, "hello + 1")
	table = append(table, "1 p 1")
	table = append(table, "1 / 0")
	table = append(table, "1 # 1")
	table = append(table, "1 % 0")
	table = append(table, "1 * 1")
	table = append(table, "1argument")
	table = append(table, "2 arguments")
	table = append(table, "4 arguments so invalid")
	table = append(table, "9223372036854775807 + 1")
	table = append(table, "9223372036854775809 - 3")
	table = append(table, "9223372036854775807 * 3")
	for _, s := range table {
		challenge.Program("doop", strings.Fields(s)...)
	}
}
