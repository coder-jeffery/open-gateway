open-gateway/
├── cmd
│   └── server
│       └── main.go          # 程序入口、信号监听、启动 http.Server、组装依赖
├── internal
│   ├── app                  # App 容器：全局依赖（logger、otel、redis、kms）
│   ├── config               # 配置加载：yaml / env / flag
│   ├── delivery
│   │   └── http
│   │       ├── router.go    # chi 路由注册、分组
│   │       ├── handler      # http handler：入参校验、返回封装
│   │       ├── middleware   # 全局/路由中间件：auth、rate-limit、cors、logging、recovery、guardrail
│   │       └── render       # 统一响应、错误封装（配合 go-chi/render）
│   ├── service              # 业务层：网关核心逻辑
│   │   ├── llm              # LLM Gateway：模型路由、fallback、token 记账、预算
│   │   ├── proxy            # 通用代理转发、流式处理、header transform
│   │   └── consumer         # 租户、虚拟 key、配额
│   ├── repository           # 数据层
│   │   ├── kv               # redis：限流、缓存、token counter
│   │   ├── audit            # prompt / 调用审计存储
│   │   └── modelregistry    # 模型元数据、厂商密钥
│   ├── provider             # 外部上游适配器
│   │   ├── openai
│   │   ├── anthropic
│   │   └── localvllm
│   └── domain               # 领域模型、常量、自定义 error、接口定义
│       ├── model.go
│       ├── errors.go
│       └── types.go
├── pkg
│   ├── pool                 # sync.Pool、对象复用
│   ├── tokenmeter           # token 计算器
│   └── retry                # 重试、hedge、熔断
├── api
│   └── openapi.yaml         # OpenAPI 文档
├── migrations               # 如果有数据库：审计表、租户表
├── scripts
│   └── run.sh
├── test
│   └── integration
├── go.mod
├── go.sum
└── config.yaml
