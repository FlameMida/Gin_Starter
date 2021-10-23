package system

type ServiceGroup struct {
	JwtService
	ApiService
	AuthorityService
	BaseMenuService
	CasbinService
	InitDBService
	MenuService
	OperationsService
	ConfigService
	UserService
}
