package main

import (
	"html/template"
	"io"
	"./handlers"
	"./mocacinno"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"fmt"
	"gopkg.in/ini.v1"
	"os"
)

type TemplateRegistry struct {
	templates *template.Template
  }

func (t *TemplateRegistry) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
  }

func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "MocacinnoServer 1.0")
		return next(c)
	}
}

  func main() {


	cfg, err := ini.Load("config.ini")
    if err != nil {
        fmt.Printf("Fail to read ini file: %v", err)
        os.Exit(1)
    }


	e := echo.New()
	logging, err := cfg.Section("server").Key("logging").Bool()
	if logging {
		e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
			Format: "method=${method}, uri=${uri}, status=${status}, time=${time_unix}, remote_ip=${remote_ip}, host=${host}, path=${path}, protocol=${protocol}, user_agent=${user_agent}, error=${error}, bytes_in=${bytes_in}, bytes_out=${bytes_out}, header=${header}, query=${query}, form=${form}\n\n",
		}))
	}
	e.Pre(middleware.AddTrailingSlash())
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))

	e.Use(ServerHeader)
	
	e.Use(middleware.SecureWithConfig(middleware.SecureConfig{
	XSSProtection:         "1; mode=block",
	ContentTypeNosniff:    "nosniff",
	ReferrerPolicy:    "same-origin",
	XFrameOptions:         "DENY",
	HSTSMaxAge:            3600,
	ContentSecurityPolicy: "default-src 'self'",
	}))
	
	e.Renderer = &TemplateRegistry{
		templates: template.Must(template.ParseGlob("views/*.html")),
	  }
	
	
	e.GET("/", handler.HomeHandler)
	e.GET("/page/feechecker/", handler.FeeCheckerHandler)
	e.POST("/page/feechecker/", mocacinno.FeeCheckerHandler)
	/*
	client := mocacinno.Client(cfg)
	blockcount := mocacinno.Blockcount(client)
	fmt.Printf("%v\n", blockcount) 
	*/
	/*
	e.GET("/page/getraw/", handler.GetRawHandler)
	e.GET("/page/rawdecode/", handler.RawDecodeHandler)
    e.GET("/page/createraw/", handler.CreateRawHandler)
    e.GET("/page/txpusher/", handler.TxPusherHandler)
	e.GET("/page/feeestimate/", handler.FeeEstimateHandler)
	e.GET("/page/block/", handler.BlockHandler)
	e.GET("/page/sigcheck/", handler.SigCheckHandler)
	e.GET("/page/validate/", handler.ValidateSigHandler)
	e.GET("/page/timestamp/", handler.TimestampedMessageHandler)
    e.GET("/page/links/", handler.UsefullLinksHandler)
	*/
	e.Static("/static", "staticfiles")
	e.Static("/fonts", "staticfiles")
	e.Logger.Fatal(e.StartTLS(cfg.Section("server").Key("ip").String() + ":" + cfg.Section("server").Key("port").String(), cfg.Section("server").Key("cert").String() , cfg.Section("server").Key("privkey").String()))
}
