package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func ip(c echo.Context) error {

	// dump, err := httputil.DumpRequest(c.Request(), false)
	// if err != nil {
	// 	c.Logger().Info("oops")
	// }

	payload := `<!doctype html>
	<html lang="en">
	  <head>
		<!-- Required meta tags -->
		<meta charset="utf-8">
		<meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
	
		<title>urip.fyi</title>
		<style>
		body {text-align: center;}
		.wrapper {display: inline-block;margin-top: 25px;position: relative;}
		
		.wrapper img {
			display: block;
			max-width: 100%;
		}
		
		.wrapper .overlay {
			position: absolute;
			top: 70%;
			left: 38%;
			transform: translate(-50%, -50%);
			color: black;
		}
		</style>
	  </head> 
	  <body>
	  <div class="wrapper">
			  <img src="./corgi.png" alt="corgi butts drive me nuts">
					<div class="overlay"><h2>`
	payload += c.RealIP()
	payload += `</h2></div>

	</div>
</body>
</html>`

	return c.HTMLBlob(http.StatusOK, []byte(payload))
}

func rawip(c echo.Context) error {

	return c.String(http.StatusOK, c.RealIP())
}

func jsonip(c echo.Context) error {

	return c.JSON(http.StatusOK, c.RealIP())
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)

	e := echo.New()
	e.File("/favicon.ico", "favicon.ico")
	e.File("/corgi.png", "corgi.png")
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.HideBanner = true
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET},
	}))

	e.GET("/", ip)
	e.GET("/raw", rawip)
	e.GET("/json", jsonip)

	e.Logger.Fatal(e.Start(":3000"))
}
