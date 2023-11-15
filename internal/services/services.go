package services

import (
	"encoding/json"
	"fmt"
	// "log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/tunghauvan/nginx-backend-protocal/internal/models"
	httpupstream "github.com/tunghauvan/nginx-backend-protocal/packages/nginx/http_upstream"
	"github.com/tunghauvan/nginx-backend-protocal/packages/nginx/parser"

	package_parser "github.com/tufanbarisyildirim/gonginx/parser"

	core "github.com/tunghauvan/nginx-backend-protocal/packages/nginx/http_core"
)

type NginxService struct {
	directory       string
	config          string
	config_contents string
}

type NginxFile struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

func NewNginxService() *NginxService {
	return &NginxService{}
}

func (s *NginxService) SetDirectory(directory string) {
	s.directory = directory
}

func (s *NginxService) SetConfig(config string) {
	s.config = config
}

// Read all files's contents of the nginx configuration file
func (s *NginxService) ReadNginxConfiguration() error {
	// Get list of files in the nginx configuration directory
	files, err := os.ReadDir(s.directory)

	// Read all files's contents of the nginx configuration file
	for _, file := range files {
		// Get the file name
		file_name := file.Name()

		// Check if the file is a nginx configuration file
		if file_name == "nginx.conf" {
			// Set the nginx configuration file
			s.config = s.directory + "/" + file_name
		}
	}

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

func getFilePaths(dir string) ([]string, error) {
	var paths []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			paths = append(paths, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return paths, nil
}

func (s *NginxService) GetNginxFile(file_name string, format string) (core.HttpContext, error) {
	if format == "json" {
		p := package_parser.NewStringParser(s.config_contents)

		conf := p.Parse()

		// Get the http directive
		http_directive := conf.FindDirectives("http")[0]

		// Parse the http directive
		httpContext, err := core.ParseHttpDirective(http_directive)

		if err != nil {
			return core.HttpContext{}, err
		}

		return httpContext, nil

	}

	return core.HttpContext{}, nil
}

func (s *NginxService) GetNginxHttp() (models.NgxHttp, error) {
	// Implementation omitted for brevity
	return models.NgxHttp{}, nil
}

func (s *NginxService) ParseNginxConfiguration(w http.ResponseWriter, r *http.Request) {
	// Parse the "name" query parameter from the request URL
	name := r.URL.Query().Get("name")
	format := r.URL.Query().Get("format")

	// Format Json response
	w.Header().Add("Content-Type", "application/json")

	if format == "" {
		format = "json"
	}

	if format == "json" {
		contents, err := s.GetNginxFile(name, format)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Write response
		json.NewEncoder(w).Encode(contents)
	}

}

func (s *NginxService) GetNgxParser() (parser.NgxParser, error) {
	parser := parser.NewNgxParser()
	parser.SetConfig(s.config_contents)
	return parser, nil
}

func (s *NginxService) GetNginxHttpJson() ([]byte, error) {
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

func (s *NginxService) GetUpstreams() ([]httpupstream.Upstream, error) {
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

func (s *NginxService) GetServers() ([]models.NgxServer, error) {
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
