package settings

import (
	"github.com/mailchain/mailchain/cmd/mailchain/internal/settings/defaults"
	"github.com/mailchain/mailchain/cmd/mailchain/internal/settings/output"
	"github.com/mailchain/mailchain/cmd/mailchain/internal/settings/values"
)

func server(s values.Store) *Server {
	return &Server{
		Port: values.NewDefaultInt(defaults.Port, s, "server.port"),
		CORS: cors(s),
	}
}

type Server struct {
	Port values.Int
	CORS CORS
}

func (o Server) Output() output.Element {
	return output.Element{
		FullName: "server",
		Attributes: []output.Attribute{
			o.Port.Attribute(),
		},
		Elements: []output.Element{
			o.CORS.Output(),
		},
	}
}

func cors(s values.Store) CORS {
	return CORS{
		AllowedOrigins: values.NewDefaultStringSlice([]string{"*"}, s, "server.cors.allowedOrigins"),
		Disabled:       values.NewDefaultBool(defaults.CORSDisabled, s, "server.cors.disabled"),
	}
}

type CORS struct {
	AllowedOrigins values.StringSlice
	Disabled       values.Bool
}

func (o CORS) Output() output.Element {
	return output.Element{
		FullName: "server.cors",
		Attributes: []output.Attribute{
			o.AllowedOrigins.Attribute(),
			o.Disabled.Attribute(),
		},
	}
}
