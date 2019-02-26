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

	payload := `<!DOCTYPE html>
	<html lang="en">
	  <head>
		  <meta charset="utf-8">
		  <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
		  
		  <meta name="description" content="urip">
		  <meta name="author" content="urip">
  
		  <link rel="stylesheet" href="https://unpkg.com/@coreui/coreui/dist/css/coreui.min.css">
		  <link rel="stylesheet" href="https://unpkg.com/@coreui/icons/css/coreui-icons.min.css">
		  <title>urip.fyi</title>
	  </head> 
	  <body class="app header-fixed bg-white">
	  <div class="app-body">
	  <table>
	<tbody>
	<tr>
	<td>&nbsp;</td>
	<td>&nbsp;</td>
	<td>&nbsp;</td>
	</tr>
	<tr>
	<td>&nbsp;</td>
	<td><img src="./corgi.png" class="img-fluid" alt="Responsive image">`
	payload += c.Request().RemoteAddr
	payload += `</td>
	<td>&nbsp;</td>
	</tr>
	<tr>
	<td>&nbsp;</td>
	<td>&nbsp;</td>
	<td>&nbsp;</td>
	</tr>
	</tbody>
	</table></div>
	
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
