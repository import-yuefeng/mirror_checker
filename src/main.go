package main

import (
	"core"
	"fmt"
	"mc_database"
)

func main() {
	var config core.Images
	conf := config.GetConf("/Users/zhuluchi/go/mirror_checker/src/default.yaml")

	mc_database.ConnDatabase()

	for k, v := range conf.Images {
		switch k {
		case "debian":
			fmt.Println("debian")
			fmt.Println(v)
			// debian := &core.Basic{v.RemotePath, v.CostomFileName, v.LocalPath, v.ImageName}
			// if result, err := debian.Checker(); err != nil {
			// 	log.Println(err)
			// } else {
			// 	if result {
			// 		fmt.Println("sync is true")
			// 	} else {
			// 		fmt.Println("sync is false")
			// 	}
			// }

		case "ubuntu":
			fmt.Println("ubuntu")
			// ubuntu := &core.Basic{v.RemotePath, v.CostomFileName, v.LocalPath, v.ImageName}
			// ubuntu.Checker()

		case "pypi":
			fmt.Println("pypi")
			// pypi := &core.Basic{v.RemotePath, v.CostomFileName, v.LocalPath}
		}
	}
}
