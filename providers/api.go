package providers

type IApi interface {

}

type Api struct {
	Config      IConfig
	Database    IDatabase
}

func NewApi(config IConfig, database IDatabase) IApi {
	return &Api{
		Config:      config,
		Database:    database,
	}
}
