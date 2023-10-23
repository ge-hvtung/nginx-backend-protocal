package nginx

type Conf struct {
}

type Http struct {
	// Create
}

// Create defintion allowed directive for specific context (http, server, location)

type HttpContext struct {
	Server []*ServerContext `nginx:"server"`
}

type ServerContext struct {
	Location []*LocationContext `nginx:"location"`
}

// define location object with allowed directive is net.IPNet
type LocationContext struct {
}
