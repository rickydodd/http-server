package config

import (
	"fmt"
	"reflect"
	"testing"
)

var defaultAddr = "0.0.0.0"
var defaultPort uint16 = 80
var customAddr = "1.1.1.1"
var customPort uint16 = 81

func TestConf(t *testing.T) {
	conf := Config()

	if conf == nil {
		t.Fatalf("error - *configBuilder not returned")
	}

	if fmt.Sprintf("%v", reflect.TypeOf(conf)) != "*config.configBuilder" {
		t.Fatalf("error - incorrect type of *configBuilder: want \"*config.configBuilder\", got: \"%v\"", reflect.TypeOf(conf))
	}
}

func TestBuild(t *testing.T) {
	conf, err := Config().Build()
	if err != nil {
		t.Fatalf("error - %s", err.Error())
	}

	if fmt.Sprintf("%v", reflect.TypeOf(conf)) != "config.config" {
		t.Fatalf("error - incorrect type of config: want \"config.config\", got: \"%v\"", reflect.TypeOf(conf))
	}
}

func TestAddr(t *testing.T) {
	testCases := []struct {
		name string
		addr *string
		want string
	}{
		{
			name: "expect addr field set to default address",
			want: defaultAddr,
		},
		{
			name: "expect addr field set to custom address",
			addr: &customAddr,
			want: customAddr,
		},
	}

	for _, tc := range testCases {
		var got string
		if tc.addr == nil {
			conf, err := Config().Build()
			if err != nil {
				t.Fatalf("error - %s", err.Error())
			}

			got = conf.Addr
		} else {
			conf, err := Config().Addr(customAddr).Build()
			if err != nil {
				t.Fatalf("error - %s", err.Error())
			}

			got = conf.Addr
		}

		if got != tc.want {
			t.Errorf("error - %s: want \"%s\", got \"%s\"", tc.name, tc.want, got)
		}
	}
}

func TestPort(t *testing.T) {
	testCases := []struct {
		name string
		port *uint16
		want uint16
	}{
		{
			name: "expect port field set to default port",
			want: defaultPort,
		},
		{
			name: "expect port field set to custom port",
			port: &customPort,
			want: customPort,
		},
	}

	for _, tc := range testCases {
		var got uint16
		if tc.port == nil {
			conf, _ := Config().Build()
			got = conf.Port
		} else {
			conf, _ := Config().Port(customPort).Build()
			got = conf.Port
		}

		if got != tc.want {
			t.Errorf("error - %s: want \"%d\", got \"%d\"", tc.name, tc.want, got)
		}
	}
}
