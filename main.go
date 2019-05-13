package main

import (
	"log"
	"net"
	"net/http"

	valid "github.com/asaskevich/govalidator"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	geoip2 "github.com/oschwald/geoip2-golang"
)

var db *geoip2.Reader

// type City struct {
// 	City struct {
// 		GeoNameID uint              `maxminddb:"geoname_id"`
// 		Names     map[string]string `maxminddb:"names"`
// 	} `maxminddb:"city"`
// 	Continent struct {
// 		Code      string            `maxminddb:"code"`
// 		GeoNameID uint              `maxminddb:"geoname_id"`
// 		Names     map[string]string `maxminddb:"names"`
// 	} `maxminddb:"continent"`
// 	Country struct {
// 		GeoNameID         uint              `maxminddb:"geoname_id"`
// 		IsInEuropeanUnion bool              `maxminddb:"is_in_european_union"`
// 		IsoCode           string            `maxminddb:"iso_code"`
// 		Names             map[string]string `maxminddb:"names"`
// 	} `maxminddb:"country"`
// 	Location struct {
// 		AccuracyRadius uint16  `maxminddb:"accuracy_radius"`
// 		Latitude       float64 `maxminddb:"latitude"`
// 		Longitude      float64 `maxminddb:"longitude"`
// 		MetroCode      uint    `maxminddb:"metro_code"`
// 		TimeZone       string  `maxminddb:"time_zone"`
// 	} `maxminddb:"location"`
// 	Postal struct {
// 		Code string `maxminddb:"code"`
// 	} `maxminddb:"postal"`
// 	RegisteredCountry struct {
// 		GeoNameID         uint              `maxminddb:"geoname_id"`
// 		IsInEuropeanUnion bool              `maxminddb:"is_in_european_union"`
// 		IsoCode           string            `maxminddb:"iso_code"`
// 		Names             map[string]string `maxminddb:"names"`
// 	} `maxminddb:"registered_country"`
// 	RepresentedCountry struct {
// 		GeoNameID         uint              `maxminddb:"geoname_id"`
// 		IsInEuropeanUnion bool              `maxminddb:"is_in_european_union"`
// 		IsoCode           string            `maxminddb:"iso_code"`
// 		Names             map[string]string `maxminddb:"names"`
// 		Type              string            `maxminddb:"type"`
// 	} `maxminddb:"represented_country"`
// 	Subdivisions []struct {
// 		GeoNameID uint              `maxminddb:"geoname_id"`
// 		IsoCode   string            `maxminddb:"iso_code"`
// 		Names     map[string]string `maxminddb:"names"`
// 	} `maxminddb:"subdivisions"`
// 	Traits struct {
// 		IsAnonymousProxy    bool `maxminddb:"is_anonymous_proxy"`
// 		IsSatelliteProvider bool `maxminddb:"is_satellite_provider"`
// 	} `maxminddb:"traits"`
// }

type urip struct {
	IP  string  `json:"ip,omitempty"`
	AR  uint16  `json:"accuracy_radius,omitempty"`
	Lat float64 `json:"latitude,omitempty"`
	Lon float64 `json:"longitude,omitempty"`
	MC  uint    `json:"metro_code,omitempty"`
	TZ  string  `json:"time_zone,omitempty"`
}

func antiChristina(realIP string) string {
	c := "192.168.0.0/16"
	b := "172.16.0.0/12"
	a := "10.0.0.0/8"

	ip := net.ParseIP(realIP)

	_, neta, err := net.ParseCIDR(a)
	if err != nil {
		log.Panic(err)
	}

	_, netb, err := net.ParseCIDR(b)
	if err != nil {
		log.Panic(err)
	}

	_, netc, err := net.ParseCIDR(c)
	if err != nil {
		log.Panic(err)
	}

	if valid.IsIPv4(realIP) && !(neta.Contains(ip) || netb.Contains(ip) || netc.Contains(ip)) {

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
	ip := net.ParseIP(antiChristina(c.RealIP()))
	record, err := db.City(ip)
	if err != nil {
		log.Fatal(err)
	}
	urip := &urip{
		IP:  antiChristina(c.RealIP()),
		AR:  record.Location.AccuracyRadius,
		Lat: record.Location.Latitude,
		Lon: record.Location.Longitude,
		MC:  record.Location.MetroCode,
		TZ:  record.Location.TimeZone,
	}
	return c.JSONPretty(http.StatusOK, urip, "  ")
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)

	db, err := geoip2.Open("GeoLite2-City.mmdb")
	if err != nil {
		log.Fatalf("Error : %v", err)
	}
	defer db.Close()
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
