package Sex

type Bullet struct {
    Message   string             `json:"message,omitempty"`
    Type      string             `json:"type,omitempty"`
    Data      interface{}        `json:"data,omitempty"`
}

type Route map[string] interface{}
type RouteDict map[string] Route

type RouteConf map[string] interface{}
type RouteConfDict map[string] RouteConf

type RawRouteFunc func(Request) ([]byte, int)
type StrRouteFunc func(Request) (string, int)
type ResRouteFunc func(Request) (*Response, int)
type PureResRouteFunc func(Request) (*Response)
type RouteFunc func(r Request) (interface{}, int)
