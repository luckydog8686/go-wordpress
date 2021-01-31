package init

import (
	"github.com/luckydog8686/config"
	"github.com/luckydog8686/logs"
)

func init()  {
	config.LoadConfig(nil)
}

func Log()  {
	logs.Debug("fuckyou")
}