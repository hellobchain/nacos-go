package constant

const (
	DEFAULT_SERVER_PORT = 8848
)

// router
const (
	// 配置相关
	CONFIGS_ROUTER = "/nacos/v1/cs/configs"
	// 监听配置
	LISTEN_CONFIGS = "/nacos/v1/cs/configs/listener"

	// 监听配置
	LIST_CONFIGS = "/nacos/v1/cs/configs/list"
)

// 服务
const (
	// 注册服务
	REGISTER_SERVICE_ROUTER = "/nacos/v1/ns/instance"
	// 获取服务列表
	LIST_SERVICE_ROUTER = "/nacos/v1/ns/instance/list"
	// 心跳服务
	HEARTBEAT_SERVICE_ROUTER = "/nacos/v1/ns/instance/beat"
)

// 用户
const (
	// 用户登录
	AUTH_LOGIN = "/nacos/v1/auth/login"
	USER_INFO  = "/nacos/v1/auth/user"
)

const (
	SRC_USER = "srcUser"
	SRC_IP   = "srcIp"
)

// tenant
const (
	TENANT_ROUTER = "/nacos/v1/cs/tenants"
)
