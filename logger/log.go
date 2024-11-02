package logger

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/go-redis-v1/pkg/utils"
)

func Info(message string) {
	fmt.Println(color.GreenString("[INFO]: "), color.CyanString(utils.GetNow()+" * "+message))
}

func Warning(message string) {
	fmt.Println(color.YellowString("[WARNING]: "), color.BlueString(utils.GetNow()+" * "+message))
}
