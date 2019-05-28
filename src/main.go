package main

import (
	"core"
	"fmt"
)

func main() {
	// fmt.Println(xpath.Request("https://mirrors.xtom.com/debian/dists/"))

	var config core.Images
	conf := config.GetConf("/Users/zhuluchi/go/mirror_checker/src/default.yaml")

	for k, v := range conf.Images {
		switch k {
		case "debian":
			fmt.Println("debian")
			fmt.Println(v.Crontab, v.CostomPath, v.CostomFileName)
			debian := &core.Basic{v.CostomPath, v.CostomFileName}
			debian.Checker()

		case "ubuntu":
			fmt.Println("ubuntu")
			fmt.Println(v.Crontab, v.CostomPath, v.CostomFileName)
			// ubuntu := &Basic{v.CostomPath, v.CostomFileName}
		case "pypi":
			fmt.Println("pypi")
			fmt.Println(v.Crontab, v.CostomPath, v.CostomFileName)
			// pypi := &Basic{v.CostomPath, v.CostomFileName}
		}
	}
}
