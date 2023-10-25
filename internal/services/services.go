package services

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/tunghauvan/nginx-backend-protocal/internal/models"
	httpupstream "github.com/tunghauvan/nginx-backend-protocal/packages/nginx/http_upstream"
	"github.com/tunghauvan/nginx-backend-protocal/packages/nginx/parser"
)

type NgxService struct {
	directory       string
	config          string
	config_contents string
}

func NewNgxService() *NgxService {
	return &NgxService{}
}

func (s *NgxService) SetDirectory(directory string) {
	s.directory = directory
}

func (s *NgxService) SetConfig(config string) {
	s.config = config
}

// Read all files's contents of the nginx configuration file
func (s *NgxService) ReadNginxConfiguration() error {
	// Read nginx config file and return the contents
	file_contents, err := os.ReadFile(s.config)
	if err != nil {
		// Print the error and return an empty string
		fmt.Println(err)
		return err
	}

	file_contents_string := string(file_contents)

	s.config_contents = file_contents_string

	return nil
}

func (s *NgxService) GetNginxHttp() (models.NgxHttp, error) {
	// Implementation omitted for brevity
	return models.NgxHttp{}, nil
}

func (s *NgxService) ParseNginxConfiguration() {
	// Implementation omitted for brevity
}

func (s *NgxService) GetNgxParser() (parser.NgxParser, error) {
	parser := parser.NewNgxParser()
	parser.SetConfig(s.config_contents)
	return parser, nil
}

func (s *NgxService) GetNginxHttpJson() ([]byte, error) {
	// Get the nginx http configuration
	nginxHttp, err := s.GetNginxHttp()
	if err != nil {
		return nil, err
	}

	// Marshal the nginx http configuration to JSON
	jsonBytes, err := json.MarshalIndent(nginxHttp, "", "  ")
	if err != nil {
		return nil, err
	}

	return jsonBytes, nil
}

func (s *NgxService) GetUpstreams() ([]httpupstream.Upstream, error) {
	// Get the nginx parser
	ngxParser, err := s.GetNgxParser()
	if err != nil {
		return nil, err
	}

	// Parse the nginx configuration
	ngxParser.InitParse()

	// Parse the upstreams
	ngxUpstreams, err := ngxParser.ParseUpstreams()
	if err != nil {
		return nil, err
	}

	return ngxUpstreams, nil
}

func (s *NgxService) GetServers() ([]models.NgxServer, error) {
	// Get the nginx parser
	ngxParser, err := s.GetNgxParser()
	if err != nil {
		return nil, err
	}

	// Parse the nginx configuration
	ngxParser.InitParse()

	// Parse the servers
	ngxServers, err := ngxParser.ParseServers()
	if err != nil {
		return nil, err
	}

	return ngxServers, nil
}
