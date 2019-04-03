package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	valid "github.com/asaskevich/govalidator"
)

func antiChristina(realIP string) string {
	if valid.IsIPv4(realIP) {
		return realIP
	}
	return "Bazinga!"
}

func ip(c echo.Context) error {

	// dump, err := httputil.DumpRequest(c.Request(), false)
	// if err != nil {
	// 	c.Logger().Info("oops")
	// }

	payload := `<!doctype html>
	<html lang="en">
	  <head>
		<!-- Global site tag (gtag.js) - Google Analytics -->
		<script async src="https://www.googletagmanager.com/gtag/js?id=UA-75116-16"></script>
		<script>
			window.dataLayer = window.dataLayer || [];
			function gtag(){dataLayer.push(arguments);}
			gtag('js', new Date());
		
			gtag('config', 'UA-75116-16');
		</script>
		<!-- Required meta tags -->
		<link rel="apple-touch-icon" sizes="57x57" href="/apple-icon-57x57.png">
		<link rel="apple-touch-icon" sizes="60x60" href="/apple-icon-60x60.png">
		<link rel="apple-touch-icon" sizes="72x72" href="/apple-icon-72x72.png">
		<link rel="apple-touch-icon" sizes="76x76" href="/apple-icon-76x76.png">
		<link rel="apple-touch-icon" sizes="114x114" href="/apple-icon-114x114.png">
		<link rel="apple-touch-icon" sizes="120x120" href="/apple-icon-120x120.png">
		<link rel="apple-touch-icon" sizes="144x144" href="/apple-icon-144x144.png">
		<link rel="apple-touch-icon" sizes="152x152" href="/apple-icon-152x152.png">
		<link rel="apple-touch-icon" sizes="180x180" href="/apple-icon-180x180.png">
		<link rel="icon" type="image/png" sizes="192x192"  href="/android-icon-192x192.png">
		<link rel="icon" type="image/png" sizes="32x32" href="/favicon-32x32.png">
		<link rel="icon" type="image/png" sizes="96x96" href="/favicon-96x96.png">
		<link rel="icon" type="image/png" sizes="16x16" href="/favicon-16x16.png">
		<link rel="manifest" href="/manifest.json">
		<meta name="msapplication-TileColor" content="#ffffff">
		<meta name="msapplication-TileImage" content="/ms-icon-144x144.png">
		<meta name="theme-color" content="#ffffff">
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
	payload += antiChristina(c.RealIP())
	payload += `</h2></div>

	</div>
</body>
</html>`

	return c.HTMLBlob(http.StatusOK, []byte(payload))
}

func rawip(c echo.Context) error {

	return c.String(http.StatusOK, antiChristina(c.RealIP()))
}

func jsonip(c echo.Context) error {

	return c.JSONPretty(http.StatusOK, antiChristina(c.RealIP()), "  ")
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)

	e := echo.New()
	e.File("/favicon.ico", "favicon.ico")
	e.File("/corgi.png", "corgi.png")
	e.File("/favicon-16x16.png", "favicon-16x16.png")
	e.File("/favicon-32x32.png", "favicon-32x32.png")
	e.File("/favicon-96x96.png", "favicon-96x96.png")
	e.File("/manifest.json", "manifest.json")
	e.File("/ms-icon-70x70.png", "ms-icon-70x70.png")
	e.File("/ms-icon-144x144.png", "ms-icon-144x144.png")
	e.File("/ms-icon-150x150.png", "ms-icon-150x150.png")
	e.File("/ms-icon-310x310.png", "ms-icon-310x310.png")
	e.File("/android-icon-36x36.png", "android-icon-36x36.png")
	e.File("/android-icon-48x48.png", "android-icon-48x48.png")
	e.File("/android-icon-72x72.png", "android-icon-72x72.png")
	e.File("/android-icon-96x96.png", "android-icon-96x96.png")
	e.File("/android-icon-144x144.png", "android-icon-144x144.png")
	e.File("/android-icon-192x192.png", "android-icon-192x192.png")
	e.File("/apple-icon-57x57.png", "apple-icon-57x57.png")
	e.File("/apple-icon-60x60.png", "apple-icon-60x60.png")
	e.File("/apple-icon-72x72.png", "apple-icon-72x72.png")
	e.File("/apple-icon-76x76.png", "apple-icon-76x76.png")
	e.File("/apple-icon-114x114.png", "apple-icon-114x114.png")
	e.File("/apple-icon-120x120.png", "apple-icon-120x120.png")
	e.File("/apple-icon-144x144.png", "apple-icon-144x144.png")
	e.File("/apple-icon-152x152.png", "apple-icon-152x152.png")
	e.File("/apple-icon-180x180.png", "apple-icon-180x180.png")
	e.File("/apple-icon-precomposed.png", "apple-icon-precomposed.png")
	e.File("/apple-icon.png", "apple-icon.png")
	e.File("/browserconfig.xml", "browserconfig.xml")
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.HideBanner = true
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET},
	}))
	e.Use(middleware.Secure())

	e.GET("/", ip)
	e.GET("/raw", rawip)
	e.GET("/json", jsonip)

	e.Logger.Fatal(e.Start(":3000"))
}
