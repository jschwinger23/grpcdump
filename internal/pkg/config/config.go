package config

import (
	"flag"
	"strings"
)

//Config represents config model
type Config struct {
	Iface          string
	Port           uint
	LogMetaHeaders string
	LoggerLevel    string
	ColorOutput    bool
	JSONOutput     bool
	ProtoPaths     string
	ProtoFiles     []string
	OutputGrpcurl  bool
}

var config *Config

var (
	iface          = flag.String("i", "eth0", "Interface to get packets from")
	port           = flag.Uint("p", 80, "TCP port")
	logMetaHeaders = flag.String("meta", "*", "Comma separated list of properties meta info for http2")
	loggerLevel    = flag.String("log-level", "info", "Logger level")
	colorOutput    = flag.Bool("color", false, "Output with color")
	jsonOutput     = flag.Bool("json", false, "Json output")
	protoPaths     = flag.String("proto-path", "", "Paths with proto files")
	protoFiles     = flag.String("proto-files", "", "Names of proto files")
	outputGrpcurl  = flag.Bool("output-as-grpcurl", false, "Output each request as grpcurl")
)

//Init inits config
func Init() {
	flag.Parse()

	config = &Config{
		*iface,
		*port,
		*logMetaHeaders,
		*loggerLevel,
		*colorOutput,
		*jsonOutput,
		*protoPaths,
		strings.Split(*protoFiles, ","),
		*outputGrpcurl,
	}
}

//GetConfig returns config
func GetConfig() *Config {
	return config
}

//GetLogMetaHeaders ...
func (config *Config) GetLogMetaHeaders() map[string]struct{} {
	result := make(map[string]struct{})

	logMetaHeaders := strings.TrimSpace(config.LogMetaHeaders)
	metaHeaders := strings.Split(logMetaHeaders, ",")

	for _, metaHeader := range metaHeaders {
		result[metaHeader] = struct{}{}
	}

	return result
}
