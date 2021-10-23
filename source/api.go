package source

import (
	"gin-starter/global"
	"gin-starter/model/system"
	"github.com/gookit/color"
	"time"

	"gorm.io/gorm"
)

var Api = new(api)

type api struct{}

var apis = []system.Api{
	{MODEL: global.MODEL{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/base/login", Description: "用户登录（必选）", ApiGroup: "base", Method: "POST"},
	{MODEL: global.MODEL{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/user/register", Description: "用户注册（必选）", ApiGroup: "user", Method: "POST"},
	{MODEL: global.MODEL{ID: 3, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/api/createApi", Description: "创建api", ApiGroup: "api", Method: "POST"},
	{MODEL: global.MODEL{ID: 4, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/api/getApiList", Description: "获取api列表", ApiGroup: "api", Method: "POST"},
	{MODEL: global.MODEL{ID: 5, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/api/getApiById", Description: "获取api详细信息", ApiGroup: "api", Method: "POST"},
	{MODEL: global.MODEL{ID: 6, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/api/deleteApi", Description: "删除Api", ApiGroup: "api", Method: "POST"},
	{MODEL: global.MODEL{ID: 7, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/api/updateApi", Description: "更新Api", ApiGroup: "api", Method: "POST"},
	{MODEL: global.MODEL{ID: 8, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/api/getAllApis", Description: "获取所有api", ApiGroup: "api", Method: "POST"},
	{MODEL: global.MODEL{ID: 9, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/authority/createAuthority", Description: "创建角色", ApiGroup: "authority", Method: "POST"},
	{MODEL: global.MODEL{ID: 10, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/authority/deleteAuthority", Description: "删除角色", ApiGroup: "authority", Method: "POST"},
	{MODEL: global.MODEL{ID: 11, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/authority/getAuthorityList", Description: "获取角色列表", ApiGroup: "authority", Method: "POST"},
	{MODEL: global.MODEL{ID: 12, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/menu/getMenu", Description: "获取菜单树（必选）", ApiGroup: "menu", Method: "POST"},
	{MODEL: global.MODEL{ID: 13, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/menu/getMenuList", Description: "分页获取基础menu列表", ApiGroup: "menu", Method: "POST"},
	{MODEL: global.MODEL{ID: 14, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/menu/addBaseMenu", Description: "新增菜单", ApiGroup: "menu", Method: "POST"},
	{MODEL: global.MODEL{ID: 15, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/menu/getBaseMenuTree", Description: "获取用户动态路由", ApiGroup: "menu", Method: "POST"},
	{MODEL: global.MODEL{ID: 16, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/menu/addMenuAuthority", Description: "增加menu和角色关联关系", ApiGroup: "menu", Method: "POST"},
	{MODEL: global.MODEL{ID: 17, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/menu/getMenuAuthority", Description: "获取指定角色menu", ApiGroup: "menu", Method: "POST"},
	{MODEL: global.MODEL{ID: 18, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/menu/deleteBaseMenu", Description: "删除菜单", ApiGroup: "menu", Method: "POST"},
	{MODEL: global.MODEL{ID: 19, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/menu/updateBaseMenu", Description: "更新菜单", ApiGroup: "menu", Method: "POST"},
	{MODEL: global.MODEL{ID: 20, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/menu/getBaseMenuById", Description: "根据id获取菜单", ApiGroup: "menu", Method: "POST"},
	{MODEL: global.MODEL{ID: 21, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/user/changePassword", Description: "修改密码（建议选择）", ApiGroup: "user", Method: "POST"},
	{MODEL: global.MODEL{ID: 23, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/user/getUserList", Description: "获取用户列表", ApiGroup: "user", Method: "POST"},
	{MODEL: global.MODEL{ID: 24, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/user/setUserAuthority", Description: "修改用户角色（必选）", ApiGroup: "user", Method: "POST"},
	{MODEL: global.MODEL{ID: 25, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/fileTransfer/upload", Description: "文件上传示例", ApiGroup: "fileTransfer", Method: "POST"},
	{MODEL: global.MODEL{ID: 26, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/fileTransfer/getFileList", Description: "获取上传文件列表", ApiGroup: "fileTransfer", Method: "POST"},
	{MODEL: global.MODEL{ID: 27, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/casbin/updateCasbin", Description: "更改角色api权限", ApiGroup: "casbin", Method: "POST"},
	{MODEL: global.MODEL{ID: 28, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/casbin/getPolicyPathByAuthorityId", Description: "获取权限列表", ApiGroup: "casbin", Method: "POST"},
	{MODEL: global.MODEL{ID: 29, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/fileTransfer/deleteFile", Description: "删除文件", ApiGroup: "fileTransfer", Method: "POST"},
	{MODEL: global.MODEL{ID: 30, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/jwt/jsonInBlacklist", Description: "jwt加入黑名单(退出，必选)", ApiGroup: "jwt", Method: "POST"},
	{MODEL: global.MODEL{ID: 31, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/authority/setDataAuthority", Description: "设置角色资源权限", ApiGroup: "authority", Method: "POST"},
	{MODEL: global.MODEL{ID: 32, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/system/getSystemConfig", Description: "获取配置文件内容", ApiGroup: "system", Method: "POST"},
	{MODEL: global.MODEL{ID: 33, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/system/setSystemConfig", Description: "设置配置文件内容", ApiGroup: "system", Method: "POST"},
	{MODEL: global.MODEL{ID: 34, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/customer/customer", Description: "创建客户", ApiGroup: "customer", Method: "POST"},
	{MODEL: global.MODEL{ID: 35, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/customer/customer", Description: "更新客户", ApiGroup: "customer", Method: "PUT"},
	{MODEL: global.MODEL{ID: 36, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/customer/customer", Description: "删除客户", ApiGroup: "customer", Method: "DELETE"},
	{MODEL: global.MODEL{ID: 37, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/customer/customer", Description: "获取单一客户", ApiGroup: "customer", Method: "GET"},
	{MODEL: global.MODEL{ID: 38, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/customer/customerList", Description: "获取客户列表", ApiGroup: "customer", Method: "GET"},
	{MODEL: global.MODEL{ID: 39, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/casbin/casbinTest/:pathParam", Description: "RESTFUL模式测试", ApiGroup: "casbin", Method: "GET"},
	{MODEL: global.MODEL{ID: 41, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/authority/updateAuthority", Description: "更新角色信息", ApiGroup: "authority", Method: "PUT"},
	{MODEL: global.MODEL{ID: 42, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/authority/copyAuthority", Description: "拷贝角色", ApiGroup: "authority", Method: "POST"},
	{MODEL: global.MODEL{ID: 43, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/user/deleteUser", Description: "删除用户", ApiGroup: "user", Method: "DELETE"},

	{MODEL: global.MODEL{ID: 54, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/operations/createOperations", Description: "新增操作记录", ApiGroup: "operations", Method: "POST"},
	{MODEL: global.MODEL{ID: 55, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/operations/deleteOperations", Description: "删除操作记录", ApiGroup: "operations", Method: "DELETE"},
	{MODEL: global.MODEL{ID: 56, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/operations/findOperations", Description: "根据ID获取操作记录", ApiGroup: "operations", Method: "GET"},
	{MODEL: global.MODEL{ID: 57, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/operations/getOperationsList", Description: "获取操作记录列表", ApiGroup: "operations", Method: "GET"},

	{MODEL: global.MODEL{ID: 61, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/operations/deleteOperationsByIds", Description: "批量删除操作历史", ApiGroup: "operations", Method: "DELETE"},
	{MODEL: global.MODEL{ID: 62, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/simpleUpload/upload", Description: "插件版分片上传", ApiGroup: "simpleUpload", Method: "POST"},
	{MODEL: global.MODEL{ID: 63, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/simpleUpload/checkFileMd5", Description: "文件完整度验证", ApiGroup: "simpleUpload", Method: "GET"},
	{MODEL: global.MODEL{ID: 64, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/simpleUpload/mergeFileMd5", Description: "上传完成合并文件", ApiGroup: "simpleUpload", Method: "GET"},
	{MODEL: global.MODEL{ID: 65, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/user/setUserInfo", Description: "设置用户信息（必选）", ApiGroup: "user", Method: "PUT"},
	{MODEL: global.MODEL{ID: 66, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/system/getServerInfo", Description: "获取服务器信息", ApiGroup: "system", Method: "POST"},
	{MODEL: global.MODEL{ID: 67, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/email/emailTest", Description: "发送测试邮件", ApiGroup: "email", Method: "POST"},

	{MODEL: global.MODEL{ID: 81, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/excel/importExcel", Description: "导入excel", ApiGroup: "excel", Method: "POST"},
	{MODEL: global.MODEL{ID: 82, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/excel/loadExcel", Description: "下载excel", ApiGroup: "excel", Method: "GET"},
	{MODEL: global.MODEL{ID: 83, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/excel/exportExcel", Description: "导出excel", ApiGroup: "excel", Method: "POST"},
	{MODEL: global.MODEL{ID: 84, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/excel/downloadTemplate", Description: "下载excel模板", ApiGroup: "excel", Method: "GET"},
	{MODEL: global.MODEL{ID: 85, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/api/deleteApisByIds", Description: "批量删除api", ApiGroup: "api", Method: "DELETE"},

	{MODEL: global.MODEL{ID: 90, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/user/setUserAuthorities", Description: "设置权限组", ApiGroup: "user", Method: "POST"},
	{MODEL: global.MODEL{ID: 91, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/user/getUserInfo", Description: "获取自身信息（必选）", ApiGroup: "user", Method: "GET"},
}

// Init @author: Flame
// @description: apis 表数据初始化1
func (a *api) Init() error {
	return global.DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1, 67}).Find(&[]system.Api{}).RowsAffected == 2 {
			color.Danger.Println("\n[Mysql] --> apis 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&apis).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> apis 表初始数据成功!")
		return nil
	})
}
