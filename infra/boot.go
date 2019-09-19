package infra

import "github.com/tietang/props/kvs"

type BootApplication struct {
	conf           kvs.ConfigSource
	starterContext StarterContext
}

func New(conf kvs.ConfigSource) *BootApplication {
	application := &BootApplication{
		conf:           conf,
		starterContext: StarterContext{},
	}
	application.starterContext[KeyProps] = conf
	return application
}

func (b *BootApplication) Start() {
	//1 。 初始化所有starter
	b.init()
	//2 。 安装所有的starter
	b.setup()
	//3 。 启动starter
}

func (b *BootApplication) init() {
	for _, starter := range StarterRegister.AllStarters() {
		starter.Init(b.starterContext)
	}
}

func (b *BootApplication) setup() {
	for _, starter := range StarterRegister.AllStarters() {
		starter.Setup(b.starterContext)
	}
}

func (b *BootApplication) start() {
	for i, starter := range StarterRegister.AllStarters() {
		starter.Start(b.starterContext)

		//如果是阻塞性质
		if starter.StartBlocking() {
			//如果是最后一个要启动的，直接启动
			if i+1 == len(StarterRegister.AllStarters()) {
				starter.Start(b.starterContext)
				//如果不是，则使用goroutine 异步启动,防止阻塞后面的starter
			} else {
				go starter.Start(b.starterContext)
			}
		} else {
			starter.Start(b.starterContext)
		}
	}
}
