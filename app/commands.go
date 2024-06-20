package main

type Value struct{
	str []string
}

type ReturnValue struct{
	str string
}

func ping(args []Value) ReturnValue{
	return ReturnValue{str: "+PONG\r\n"}
}

func echo(args []Value) ReturnValue{
	if len(args[0].str) > 1 {
		return ReturnValue{str: args[0].str[1]}
	}
	return ReturnValue{str: "+\r\n"}
}

