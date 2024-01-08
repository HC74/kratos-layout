package logic

import super "kratos-layout/config"

type Logic struct {
	conf *super.Config
}

func NewLogic(conf *super.Config) *Logic {
	return &Logic{
		conf: conf,
	}
}
