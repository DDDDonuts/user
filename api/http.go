package api

import "github.com/gin-gonic/gin"

type HTTPTransport struct {
	*gin.Engine
	svc UserService
}

type UserService interface {
	GetById()
	GetAll() []string
	Create()
}

const (
	CreateUserPath  = "/user/create"
	ListUsersPath   = "/user/list"
	HealthCheckPath = "/health"
)

func NewHTTPTransport(us UserService) HTTPTransport {
	g := gin.Default()
	ht := HTTPTransport{g, us}
	ht.POST(CreateUserPath, ht.CreateNewUser)
	ht.GET(ListUsersPath, ht.ListUser)
	ht.GET(HealthCheckPath, ht.Health)
	return ht
}

func (ht HTTPTransport) CreateNewUser(ctx *gin.Context) {
	ht.svc.Create()
	ctx.JSON(200, gin.H{"status": "OK"})
}

func (ht HTTPTransport) ListUser(ctx *gin.Context) {
	users := ht.svc.GetAll()
	ctx.JSON(200, gin.H{"users": users})
}

func (ht HTTPTransport) Health(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"status": "OK"})
}
