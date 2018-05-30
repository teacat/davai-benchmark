package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"time"

	"github.com/astaxie/beego"
	"github.com/bmizerany/pat"
	"github.com/dimfeld/httptreemux"
	"github.com/gin-gonic/gin"
	"github.com/go-martini/martini"
	"github.com/gorilla/mux"
	"github.com/gramework/gramework"
	"github.com/julienschmidt/httprouter"
	"github.com/labstack/echo"
	"github.com/teacat/davai"
	"github.com/valyala/fasthttp"
	"github.com/zenazn/goji"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Ctx.WriteString("hello world")
}

func serverStartup() {
	go nethttpServer([]string{
		"/",
		"/users",
		"/user/:name",
		"/user/:name/:name2",
		"/user/:name/:name2/:name3",
	})
	go httprouterServer([]string{
		"/",
		"/users",
		"/user/:name",
		"/user/:name/:name2",
		"/user/:name/:name2/:name3",
	})
	go martiniServer([]string{
		"/",
		"/users",
		"/user/:name",
		"/user/:name/:name2",
		"/user/:name/:name2/:name3",
	})
	go gojiServer([]string{
		"/",
		"/users",
		"/user/:name",
		"/user/:name/:name2",
		"/user/:name/:name2/:name3",
	})
	go patServer([]string{
		"/",
		"/users",
		"/user/:name",
		"/user/:name/:name2",
		"/user/:name/:name2/:name3",
	})
	go fasthttpServer([]string{
		"/",
		"/users",
		"/user/:name",
		"/user/:name/:name2",
		"/user/:name/:name2/:name3",
	})
	go ginServer([]string{
		"/",
		"/users",
		"/user/:name",
		"/user/:name/:name2",
		"/user/:name/:name2/:name3",
	})
	go muxServer([]string{
		"/",
		"/users",
		"/user/{name}",
		"/user/{name}/{name2}",
		"/user/{name}/{name2}/{name3}",
	})
	go grameworkServer([]string{
		"/",
		"/users",
		"/user/:name",
		"/user/:name/:name2",
		"/user/:name/:name2/:name3",
	})
	go httptreemuxServer([]string{
		"/",
		"/users",
		"/user/:name",
		"/user/:name/:name2",
		"/user/:name/:name2/:name3",
	})
	go echoServer([]string{
		"/",
		"/users",
		"/user/:name",
		"/user/:name/:name2",
		"/user/:name/:name2/:name3",
	})
	go davaiServer([]string{
		"/",
		"/users",
		"/user/{name}",
		"/user/{name}/{name2}",
		"/user/{name}/{name2}/{name3}",
	})
	go beegoServer([]string{
		"/",
		"/users",
		"/user/:name",
		"/user/:name/:name2",
		"/user/:name/:name2/:name3",
	})

	log.Printf("等待所有伺服器啟動⋯")
	<-time.After(time.Second * 2)
}

var (
	testURLs = []string{
		// /
		//"/",
		// /users
		//"/users",
		// /user/:name
		//"/user/yamiodymel",
		// /user/:name/:name2
		"/user/yamiodymel/admin",
	}
	testServers = []string{
		"davai",
		"gramework",
		//"nethttp",
		"httprouter",
		"martini",
		//"goji",
		"pat",
		//"fasthttp",
		"gin",
		"mux",
		"httptreemux",
		"echo",
		"beego",
	}
)

func serverBenchmark() {
	for _, v := range testServers {
		for _, vv := range testURLs {
			run(v, vv)
			<-time.After(time.Millisecond * 100)
		}
	}
}

func main() {
	serverStartup()
	serverBenchmark()
	serverReport()
}

func serverReport() {
	// Create a csv file
	f, err := os.Create("./benchmark.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	// Write Unmarshaled json data to CSV file
	w := csv.NewWriter(f)

	w.Write([]string{
		"RouterName",
		"TotalTime",
		"Location",
		"Goroutines",
		"TotalRequest",
		"ActualTime",
		"ReadSize",
		"RequestsPerSecond",
		"TransferPerSecond",
		"AverageRequestTime",
		"FastestRequestTime",
		"SlowestRequestTime",
		"NumberOfError",
	})

	var record []string
	for _, v := range results {
		record = append(record, v["RouterName"])
		record = append(record, v["TotalTime"])
		record = append(record, v["Location"])
		record = append(record, v["Goroutines"])
		record = append(record, v["TotalRequest"])
		record = append(record, v["ActualTime"])
		record = append(record, v["ReadSize"])
		record = append(record, v["RequestsPerSecond"])
		record = append(record, v["TransferPerSecond"])
		record = append(record, v["AverageRequestTime"])
		record = append(record, v["FastestRequestTime"])
		record = append(record, v["SlowestRequestTime"])
		record = append(record, v["NumberOfError"])
		w.Write(record)
		record = []string{}
	}

	w.Flush()
}

var (
	nethttpPort     = ":9086"
	httprouterPort  = ":9087"
	martiniPort     = ":9088"
	gojiPort        = ":9089"
	patPort         = ":9090"
	fasthttpPort    = ":9091"
	ginPort         = ":9092"
	muxPort         = ":9093"
	grameworkPort   = ":9094"
	httptreemuxPort = ":9095"
	echoPort        = ":9096"
	davaiPort       = ":9097"
	beegoPort       = ":9098"
	results         = []map[string]string{}
)

func run(server string, path string) {
	var port string
	switch server {
	case "nethttp":
		port = nethttpPort
	case "httprouter":
		port = httprouterPort
	case "martini":
		port = martiniPort
	case "goji":
		port = gojiPort
	case "pat":
		port = patPort
	case "fasthttp":
		port = fasthttpPort
	case "gin":
		port = ginPort
	case "mux":
		port = muxPort
	case "gramework":
		port = grameworkPort
	case "httptreemux":
		port = httptreemuxPort
	case "echo":
		port = echoPort
	case "davai":
		port = davaiPort
	case "beego":
		port = beegoPort
	}

	out, err := exec.Command("/Users/YamiOdymel/go/bin/go-wrk", "-c", "5", "-d", "1", "-redir", fmt.Sprintf("http://localhost%s%s", port, path)).Output()
	if err != nil {
		log.Fatal(err)
	}
	parsed := analyze(string(out))
	parsed["RouterName"] = server
	parsed["Location"] = path
	results = append(results, parsed)
	fmt.Printf("%s, %s \n", server, path)
}

// https://stackoverflow.com/a/30483899/5203951
func mapSubexpNames(m, n []string) map[string]string {
	if len(m) == 0 || len(n) == 0 {
		panic("錯誤")
	}
	m, n = m[1:], n[1:]
	r := make(map[string]string, len(m))
	for i := range n {
		r[n[i]] = m[i]
	}
	return r
}

func analyze(res string) map[string]string {
	r := regexp.MustCompile(`Running (?P<TotalTime>.*) test @ (?P<Location>.*)
  (?P<Goroutines>.*) goroutine\(s\) running concurrently
(?P<TotalRequest>.*) requests in (?P<ActualTime>.*), (?P<ReadSize>.*) read
Requests\/sec:		(?P<RequestsPerSecond>.*)
Transfer\/sec:		(?P<TransferPerSecond>.*)
Avg Req Time:		(?P<AverageRequestTime>.*)
Fastest Request:	(?P<FastestRequestTime>.*)
Slowest Request:	(?P<SlowestRequestTime>.*)
Number of Errors:	(?P<NumberOfError>.*)`)
	m := r.FindStringSubmatch(res)
	n := r.SubexpNames()
	v := mapSubexpNames(m, n)

	//fmt.Printf("%+v", res)

	fmt.Printf("%s (reqs/s) | ", v["RequestsPerSecond"])

	return mapSubexpNames(m, n)
}

func nethttpServer(paths []string) {
	for _, v := range paths {
		http.HandleFunc(v, func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello, world!\n")
		})
	}
	err := http.ListenAndServe(nethttpPort, nil)
	if err != nil {
		panic(err)
	}
}

func httprouterServer(paths []string) {
	httpRouter := httprouter.New()
	for _, v := range paths {
		httpRouter.GET(v, func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
			fmt.Fprint(w, "Hello, world!\n")
		})
	}
	err := http.ListenAndServe(httprouterPort, httpRouter)
	if err != nil {
		panic(err)
	}
}

func martiniServer(paths []string) {
	martiniRouter := martini.NewRouter()
	martiniServer := martini.New()
	for _, v := range paths {
		martiniRouter.Get(v, func() string {
			return "Hello, world!\n"
		})
	}
	martiniServer.Action(martiniRouter.Handle)
	martiniServer.RunOnAddr(martiniPort)
	//martiniRouter.RunOnAddr(martiniPort)
}

func gojiServer(paths []string) {
	for _, v := range paths {
		goji.Get(v, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, "Hello, world!\n")
		}))
	}
	err := http.ListenAndServe(gojiPort, goji.DefaultMux)
	if err != nil {
		panic(err)
	}
}

func patServer(paths []string) {
	patRouter := pat.New()
	for _, v := range paths {
		patRouter.Get(v, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, "Hello, world!\n")
		}))
	}
	err := http.ListenAndServe(patPort, patRouter)
	if err != nil {
		panic(err)
	}
}

func fasthttpServer(paths []string) {

	requestHandler := func(ctx *fasthttp.RequestCtx) {
		fmt.Fprintf(ctx, "Hello, world!\n")
	}

	err := fasthttp.ListenAndServe(fasthttpPort, requestHandler)
	if err != nil {
		panic(err)
	}

}

func ginServer(paths []string) {
	ginRouter := gin.New()
	for _, v := range paths {
		ginRouter.GET(v, func(c *gin.Context) {
			c.String(200, "Hello, world!\n")
		})
	}
	err := ginRouter.Run(ginPort)
	if err != nil {
		panic(err)
	}
}

func muxServer(paths []string) {
	muxRouter := mux.NewRouter()
	for _, v := range paths {
		muxRouter.HandleFunc(v, func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, "Hello, world!\n")
		})
	}
	err := http.ListenAndServe(muxPort, muxRouter)
	if err != nil {
		panic(err)
	}
}

func grameworkServer(paths []string) {
	grameRouter := gramework.New()
	for _, v := range paths {
		grameRouter.GET(v, "Hello, world!\n")
	}
	err := grameRouter.ListenAndServe(grameworkPort)
	if err != nil {
		panic(err)
	}
}

func httptreemuxServer(paths []string) {
	treemuxRouter := httptreemux.New()
	for _, v := range paths {
		treemuxRouter.GET(v, func(w http.ResponseWriter, r *http.Request, _ map[string]string) {
			fmt.Fprint(w, "Hello, world!\n")
		})
	}
	err := http.ListenAndServe(httptreemuxPort, treemuxRouter)
	if err != nil {
		panic(err)
	}
}

func echoServer(paths []string) {
	echoRouter := echo.New()
	for _, v := range paths {
		echoRouter.GET(v, func(c echo.Context) error {
			return c.String(200, "Hello, world!\n")
		})
	}
	err := echoRouter.Start(echoPort)
	if err != nil {
		panic(err)
	}
}

func davaiServer(paths []string) {
	davaiRouter := davai.New()
	for _, v := range paths {
		davaiRouter.Get(v, func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, "Hello, world!\n")
		})
	}
	err := davaiRouter.Run(davaiPort)
	if err != nil {
		panic(err)
	}
}

func beegoServer(paths []string) {
	for _, v := range paths {
		beego.Router(v, &MainController{})
	}
	beego.SetLogger("file", `{"filename":"logs/test.log"}`)
	beego.BeeLogger.DelLogger("console")
	beego.Run(beegoPort)
}
