package commands

type Value struct{
	str string
	val int
}

var commands = map[string]func([]Value) Value{
	"PING": ping,
	"ECHO": echo,
}

func ping(args []Value) Value{
	return Value{str: "+PONG\r\n"}
}

func echo(args []Value) Value{
	if len(args) > 0 {
		return args[0]
	}
	return Value{str: "+\r\n"}
}

