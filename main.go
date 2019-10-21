package main

import (
	"fmt"

	"github.com/gogf/gf/os/gcmd"
	"github.com/gogf/gf/text/gstr"
	"github.com/zhaopengme/sdkvm/commands/help"
	"github.com/zhaopengme/sdkvm/commands/ls"
)

const (
	VERSION = "v0.0.1"
)

var (
	verContent  = fmt.Sprintf("sdkvm version %s, https://sdkvm.github.comg", VERSION)
	helpContent = gstr.TrimLeft(`
USAGE
	sdkvm COMMAND [ARGUMENT] [OPTION]
COMMAND
	version    show version info
	help       show more information about a specified command
OPTION
	-?,-h      show this help or detail for specified command
	-v,-i      show version information
ADDITIONAL
	Use 'sdkvm help COMMAND' or 'sdkvm COMMAND -h' for detail about a command, which has '...' in the tail of their comments.
	`)
)

func main() {
	command := gcmd.GetArg(1)
	fmt.Println(gcmd.GetArgAll())
	fmt.Println(command)
	if gcmd.ContainsOpt("h") && command != "" {
		help.Run(command)
		return
	}
	switch command {
	case "ls":
		ls.Run()
	case "help":
		help.Run(gcmd.GetArg(2))
	case "version":
		fmt.Println(verContent)
	default:
		fmt.Println(helpContent)
	}

}
