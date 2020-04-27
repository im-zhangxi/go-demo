package app

// Config 程序配置
type Config struct {
	App struct {
		Port int32 `iddd:"port"`
	} `iddd:"app"`
	DB struct {
		URL  string `iddd:"url"`
	} `iddd:"db"`
}
