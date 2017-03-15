package main

import (
	"flag"
	"fmt"
	"github.com/Yuniii/plweb-api-server/controller"
	"gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/cors"
	"gopkg.in/kataras/iris.v6/middleware/logger"
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
