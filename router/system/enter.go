package system

type RouterGroup struct {
	ApiRouter
	AuthorityRouter
	BaseRouter
	CasbinRouter
	InitRouter
	JwtRouter
	MenuRouter
	OperationsRouter
	SysRouter
	UserRouter
}
