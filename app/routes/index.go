package routes

type Router interface {
	Register()
	RegisterMiddlewares()
}
