package main

import (
	"log"
	"strconv"
)

type Value struct{
	str []string
}

var tree = &AVLTree{}

type ReturnValue struct{
	str string
}

func ping(args []Value) ReturnValue{
	return ReturnValue{str: "+PONG"}
}

func echo(args []Value) ReturnValue{
	if len(args[0].str) > 1 {
		BulkStringSize := "$" + strconv.Itoa(len(args[0].str[1])) + "\r\n"
		return ReturnValue{str: BulkStringSize + args[0].str[1]}
	}
	return ReturnValue{str: "+\r\n"}
}

// map[raspberry] = grape
func set(args []Value) ReturnValue{

	if len(args[0].str) != 3 {
		return ReturnValue{str: "-ERR wrong number of arguments for 'set' command"}
	}

	tree.Insert(args[0].str[1] ,args[0].str[2])

	return ReturnValue{str: "+OK"}
}

func get(args []Value) ReturnValue{
	log.Printf("args: %v", args)
	if len(args) != 1 {
		return ReturnValue{str: "-ERR wrong number of arguments for 'get' command"}
	}
	if val, ok := tree.Search(args[0].str[1]); ok {
		BulkStringSize := "$" + strconv.Itoa(len(val)) + "\r\n"
		return ReturnValue{str: BulkStringSize + val}
	}
	return ReturnValue{str: "$-1"}
}
