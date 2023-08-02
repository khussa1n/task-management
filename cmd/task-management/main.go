package task_management

import (
	"fmt"
	"github.com/khussa1n/task-management/internal/app"
	"github.com/khussa1n/task-management/internal/config"
)

func main() {
	cfg, err := config.InitConfig("config.yaml")
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%#v", cfg))

	err = app.Run(cfg)
	if err != nil {
		panic(err)
	}
}
