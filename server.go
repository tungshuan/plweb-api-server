package main

import (
	"flag"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/yuniii/plweb-api-server/controller"
)

var (
	port = flag.String("port", ":8888", "port to listen to")
)

func main() {
	flag.Parse()

	iris.UseFunc(logger.Default())
	iris.Get("/", controller.Index)
	iris.Get("/course/:courseId/:lessonId", controller.GetCourse)
	iris.Listen(*port)
}
