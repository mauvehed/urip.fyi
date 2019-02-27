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
	
		<!-- Bootstrap CSS -->
		<link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0/css/bootstrap.min.css" integrity="sha384-Gn5384xqQ1aoWXA+058RXPxPg6fy4IWvTNh0E263XmFcJlSAwiGgFAW/dAiS6JXm" crossorigin="anonymous">
		 <title>urip.fyi</title>
	  </head> 
	  <body>
	  <div class="container">
		  <div class="row">
		  	<div class="col-12">
			  <div class="d-inline-block position-relative">
			  <img src="./corgi.png" class="d-block" alt="corgi butts drive me nuts">
			  </div>
		  		<div class="position-relative">`
	payload += c.Request().RemoteAddr
	payload += `</div>
			</div>
		</div>
	</div>
	<!-- Optional JavaScript -->
    <!-- jQuery first, then Popper.js, then Bootstrap JS -->
    <script src="https://code.jquery.com/jquery-3.2.1.slim.min.js" integrity="sha384-KJ3o2DKtIkvYIK3UENzmM7KCkRr/rE9/Qpg6aAZGJwFDMVNA/GpGFF93hXpG5KkN" crossorigin="anonymous"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.12.9/umd/popper.min.js" integrity="sha384-ApNbgh9B+Y1QKtv3Rn7W3mgPxhU9K/ScQsAP7hUibX39j7fakFPskvXusvfa0b4Q" crossorigin="anonymous"></script>
    <script src="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0/js/bootstrap.min.js" integrity="sha384-JZR6Spejh4U02d8jOt6vLEHfe/JQGiRRSQQxSfFWpi1MquVdAyjUar5+76PVCmYl" crossorigin="anonymous"></script>
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
