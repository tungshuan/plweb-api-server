package main

import (
	"flag"
	"fmt"
	"github.com/Yuniii/plweb-api-server/controller"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
)

var (
	port = flag.String("port", ":8888", "port to listen to")
)

func main() {
	flag.Parse()

	iris.UseFunc(logger.Default())

	iris.Get("/", controller.Index)
	iris.Get("/course/:courseId/:lessonId", controller.GetCourse)

	fmt.Printf("listen on %s\n", *port)
	iris.Listen(*port)
}
