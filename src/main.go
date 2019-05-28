package main

import (
	"core"
	"fmt"
)

func main() {
	var config core.Images
	conf := config.GetConf("/Users/zhuluchi/go/mirror_checker/src/default.yaml")

	for k, v := range conf.Images {
		switch k {
		case "debian":
			fmt.Println("debian")
			fmt.Println(v.Crontab, v.LocalPath, v.RemotePath, v.ImageName)
			debian := &core.Basic{v.RemotePath, v.CostomFileName, v.LocalPath, v.ImageName}
			if debian.Checker() {
				fmt.Println("sync is true")
			} else {
				fmt.Println("sync is false")
			}

		case "ubuntu":
			fmt.Println("ubuntu")
			fmt.Println(v.Crontab, v.LocalPath, v.RemotePath, v.ImageName)
			ubuntu := &core.Basic{v.RemotePath, v.CostomFileName, v.LocalPath, v.ImageName}
			ubuntu.Checker()

		case "pypi":
			fmt.Println("pypi")
			fmt.Println(v.Crontab, v.LocalPath, v.RemotePath, v.ImageName)
			// pypi := &core.Basic{v.RemotePath, v.CostomFileName, v.LocalPath}
		}
	}
}
