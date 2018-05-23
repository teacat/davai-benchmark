package benchmark

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/astaxie/beego"
	"github.com/dimfeld/httptreemux"
	"github.com/gin-gonic/gin"
	"github.com/go-martini/martini"
	"github.com/gorilla/mux"
	"github.com/gramework/gramework"
	"github.com/julienschmidt/httprouter"
	"github.com/labstack/echo"
	davai "github.com/teacat/go-davai"
	"github.com/valyala/fasthttp"
	// The MySQL driver.
	_ "github.com/go-sql-driver/mysql"
)

var (
	dsn = "root:root@/test?charset=utf8"
	i   = 0
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
)

type user struct {
	ID       int    `xorm:"ID" db:"ID"`
	Username string `xorm:"Username" db:"Username"`
	Password string `xorm:"Password" db:"Password"`
}

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Ctx.WriteString("hello world")
}

func init() {
	var err error

	go func() {
		httpRouter = httprouter.New()
		httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
			fmt.Fprint(w, "Hello, world!\n")
		})
		err = http.ListenAndServe(":9090", httpRouter)
		if err != nil {
			panic(err)
		}
	}()

	go func() {
		martiniRouter = martini.Classic()
		martiniRouter.Get("/", func() string {
			return "Hello, world!\n"
		})
		martiniRouter.Run()
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
	}()

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

	go func() {

		echoRouter = echo.New()
		echoRouter.GET("/", func(c echo.Context) error {
			return c.String(200, "Hello, world!\n")
		})
		err = echoRouter.Start(":9096")
		if err != nil {
			panic(err)
		}
	}()

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

	beego.Router("/", &MainController{})
	beego.Run()

	<-time.After(1 * time.Second)
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

func BenchmarkDavai_Static(b *testing.B) {
	w := new(mockResponseWriter)
	r, _ := http.NewRequest("GET", "/", nil)
	for x := 0; x < b.N; x++ {
		davaiRouter.ServeHTTP(w, r)
	}
}

func BenchmarkTreeMux_Static(b *testing.B) {
	w := new(mockResponseWriter)
	r, _ := http.NewRequest("GET", "/", nil)
	for x := 0; x < b.N; x++ {
		treemuxRouter.ServeHTTP(w, r)
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
