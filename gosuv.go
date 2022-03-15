package main

import (
	"fmt"
	"github.com/urfave/cli"
	"gosuv/server"
	"gosuv/utils"
	"os"
	"path/filepath"
)

var cl = &Client{}

func main() {

	//初始global 变量
	server.CfgDir = utils.GetExecPath()
	server.CurrentDir = utils.GetCurrentPath()
	server.CmdDir = utils.GetExecPath()

	app := cli.NewApp()
	app.Name = server.AppName
	app.Version = server.Version
	app.Usage = "golang supervisor"
	app.Before = func(c *cli.Context) error {
		var err error
		server.CfgFile = c.GlobalString("conf")

		if filepath.IsAbs(server.CfgFile) {
			server.CfgDir = filepath.Dir(server.CfgFile)
		} else {
			server.CfgDir = filepath.Dir(filepath.Join(utils.GetExecPath(), server.CfgFile))
		}
		server.Cfg, err = server.ReadConf(server.CfgFile)
		if err != nil {
			fmt.Printf("read conf failed: %s", err)
			os.Exit(-1)
		}
		//加载client配置
		cl = NewClient()
		return nil
	}
	app.Authors = []cli.Author{{
		Name:  server.Author,
		Email: server.Email,
	},
	}
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "conf, c",
			Usage: "config file",
			Value: server.CfgFile + server.DefaultConfig,
		},
	}
	app.Commands = []cli.Command{
		{
			Name:  "start-server",
			Usage: "Start server and run in background",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "foreground, f",
					Usage: "start in foreground",
				},
				cli.StringFlag{
					Name:  "conf, c",
					Usage: "config file",
					Value: server.CfgFile + server.DefaultConfig,
				},
			},
			Action: actionStartServer,
		},
		{
			Name:   "status-server",
			Usage:  "Status server",
			Action: actionStatus,
		},
		{
			Name:   "stop-server",
			Usage:  "Stop server",
			Action: actionShutdown,
		},
		{
			Name:   "restart-server",
			Usage:  "Restart server",
			Action: actionRestart,
		},
		{
			Name:   "kill-server",
			Usage:  "Kill server by pid file.",
			Action: actionKill,
		},
		{
			Name:   "status",
			Usage:  "Show program status",
			Action: actionProgramStatus,
		},
		{
			Name:   "start",
			Usage:  "Start program",
			Action: actionStart,
		},
		{
			Name:   "stop",
			Usage:  "Stop program",
			Action: actionStop,
		},
		{
			Name:   "restart",
			Usage:  "Restart program",
			Action: actionRestartProgram,
		},
		{
			Name:   "reload",
			Usage:  "Reload config file",
			Action: actionReload,
		},
		{
			Name:    "conftest",
			Aliases: []string{"t"},
			Usage:   "Test if config file is valid",
			Action:  actionConfigTest,
		},
		{
			Name:    "edit",
			Aliases: []string{"e"},
			Usage:   "Edit config file",
			Action:  actionEdit,
		},
		{
			Name:    "version",
			Usage:   "Show version",
			Aliases: []string{"v"},
			Action:  actionVersion,
		},
	}
	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
