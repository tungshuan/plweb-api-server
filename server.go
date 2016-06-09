package main

import (
	"flag"
	"fmt"
	"github.com/Yuniii/plweb-api-server/controller"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/cors"
	"github.com/kataras/iris/middleware/logger"
)

var (
	port = flag.String("port", ":8888", "port to listen to")
)

func main() {
	flag.Parse()

	iris.Use(logger.New(iris.Logger()))
	iris.Use(cors.DefaultCors())

	iris.Get("/", controller.Index)
	iris.Get("/course/:courseID/:lessonID", controller.GetCourse)

	iris.Post("/submit/:classID/:courseID/:lessonID/:qn", controller.SubmitCode)

	fmt.Printf("listen on %s\n", *port)
	iris.Listen(*port)
}
