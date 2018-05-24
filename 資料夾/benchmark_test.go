package benchmark

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/astaxie/beego"
	"github.com/dimfeld/httptreemux"
	"github.com/gin-gonic/gin"
	"github.com/go-martini/martini"
	"github.com/gorilla/mux"
	"github.com/gramework/gramework"
	"github.com/julienschmidt/httprouter"
	"github.com/labstack/echo"
	davai "github.com/teacat/go-davai"
)

var (
	httpRouter    *httprouter.Router
	martiniRouter *martini.ClassicMartini
	ginRouter     *gin.Engine
	muxRouter     *mux.Router
	grameRouter   *gramework.App
	treemuxRouter *httptreemux.TreeMux
	echoRouter    *echo.Echo
	davaiRouter   *davai.Router
	beegoRouter   *beego.App
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Ctx.WriteString("hello world")
}

func init() {
	var err error

	go func() {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello, world!\n")
		})
		err = http.ListenAndServe(":9088", nil)
		if err != nil {
			panic(err)
		}
	}()

	/*go func() {
		httpRouter = httprouter.New()
		httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
			fmt.Fprint(w, "Hello, world!\n")
		})
		err = http.ListenAndServe(":9089", httpRouter)
		if err != nil {
			panic(err)
		}
	}()

	go func() {
		martiniRouter = martini.Classic()
		martiniRouter.Get("/", func() string {
			return "Hello, world!\n"
		})
		martiniRouter.RunOnAddr(":9090")
	}()

	go func() {
		requestHandler := func(ctx *fasthttp.RequestCtx) {
			fmt.Fprintf(ctx, "Hello, world!\n")
		}

		err = fasthttp.ListenAndServe(":9091", requestHandler)
		if err != nil {
			panic(err)
		}
	}()

	go func() {

		ginRouter = gin.New()
		ginRouter.GET("/", func(c *gin.Context) {
			c.String(200, "Hello, world!\n")
		})
		err = ginRouter.Run(":9092")
		if err != nil {
			panic(err)
		}
	}()

	go func() {
		muxRouter = mux.NewRouter()
		muxRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, "Hello, world!\n")
		})
		err = http.ListenAndServe(":9093", muxRouter)
		if err != nil {
			panic(err)
		}
	}()

	go func() {
		grameRouter = gramework.New()
		grameRouter.GET("/", "Hello, world!\n")
		err = grameRouter.ListenAndServe(":9094")
		if err != nil {
			panic(err)
		}
	}()*/

	go func() {

		treemuxRouter = httptreemux.New()
		treemuxRouter.GET("/", func(w http.ResponseWriter, r *http.Request, _ map[string]string) {
			fmt.Fprint(w, "Hello, world!\n")
		})
		err = http.ListenAndServe(":9095", treemuxRouter)
		if err != nil {
			panic(err)
		}
	}()
	/*
		go func() {

			echoRouter = echo.New()
			echoRouter.GET("/", func(c echo.Context) error {
				return c.String(200, "Hello, world!\n")
			})
			err = echoRouter.Start(":9096")
			if err != nil {
				panic(err)
			}
		}()*/

	go func() {
		davaiRouter = davai.New()
		davaiRouter.Get("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, "Hello, world!\n")
		})
		err = davaiRouter.Run(":9097")
		if err != nil {
			panic(err)
		}
	}()

	/*go func() {
		beegoRouter = beego.Router("/", &MainController{})
		beego.Run(":9098")
	}()*/

	//<-time.After(1 * time.Second)
}

type mockResponseWriter struct{}

func (m *mockResponseWriter) Header() (h http.Header) {
	return http.Header{}
}

func (m *mockResponseWriter) Write(p []byte) (n int, err error) {
	return len(p), nil
}

func (m *mockResponseWriter) WriteString(s string) (n int, err error) {
	return len(s), nil
}

func (m *mockResponseWriter) WriteHeader(int) {}

func BenchmarkNetHTTP_Static(b *testing.B) {
	for x := 0; x < b.N; x++ {
		_, err := http.Get("http://localhost:9088/")
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkDavai_Static(b *testing.B) {
	for x := 0; x < b.N; x++ {
		resp, err := http.Get("http://localhost:9097/")
		if err != nil {
			panic(err)
		}
		fmt.Printf("%+v", resp)
	}
}

func BenchmarkTreeMux_Static(b *testing.B) {
	for x := 0; x < b.N; x++ {
		_, err := http.Get("http://localhost:9095/")
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkEcho_Static(b *testing.B) {
	w := new(mockResponseWriter)
	r, _ := http.NewRequest("GET", "/", nil)
	for x := 0; x < b.N; x++ {
		echoRouter.ServeHTTP(w, r)
	}
}

func BenchmarkGin_Static(b *testing.B) {
	w := new(mockResponseWriter)
	r, _ := http.NewRequest("GET", "/", nil)
	for x := 0; x < b.N; x++ {
		ginRouter.ServeHTTP(w, r)
	}
}

func BenchmarkBeego_Static(b *testing.B) {
	w := new(mockResponseWriter)
	r, _ := http.NewRequest("GET", "/", nil)
	for x := 0; x < b.N; x++ {
		beegoRouter.Server.Handler.ServeHTTP(w, r)
	}
}
