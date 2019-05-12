package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

func Test_antiChristina(t *testing.T) {
	type args struct {
		realIP string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// Make sure it will validate correct IPv4 Address
		{
			name: "Check IPv4",
			args: args{
				realIP: "1.1.1.1",
			},
			want: "1.1.1.1",
		},
		// IPV6 is not currently cupported
		{
			name: "Check IPv6",
			args: args{
				realIP: "2606:4700:4700::1111",
			},
			want: "Bazinga!",
		},
		// Check for RFC1918 addresses
		// These should fail
		{
			name: "Check Not RFC1918",
			args: args{
				realIP: "192.168.1.1",
			},
			want: "Bazinga!",
		},
		{
			name: "Check Not RFC1918",
			args: args{
				realIP: "172.16.1.1",
			},
			want: "Bazinga!",
		},
		{
			name: "Check Not RFC1918",
			args: args{
				realIP: "10.0.1.1",
			},
			want: "Bazinga!",
		},
		// Check for XSS Style
		{
			name: "Check Not RFC1918",
			args: args{
				realIP: "<xss>",
			},
			want: "Bazinga!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := antiChristina(tt.args.realIP); got != tt.want {
				t.Errorf("antiChristina() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ip(t *testing.T) {
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "basic http response",
			args: args{
				c: &echo.Context{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(echo.GET, "/", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			if assert.NoError(t, ip(tt.args.c)) {
				assert.Equal(t, http.StatusOK, rec.Code)
				assert.HTTPBodyContains(t, ip, "GET", "/", nil, tt.args.c.RealIP)
			}

		})
	}
}

func Test_rawip(t *testing.T) {

	type args struct {
		c echo.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := rawip(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("rawip() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_jsonip(t *testing.T) {
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := jsonip(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("jsonip() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
