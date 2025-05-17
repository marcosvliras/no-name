package config

func InitConfig() {
	err := newResource()
	if err != nil {
		panic(err)
	}

	err = initConn()
	if err != nil {
		panic(err)
	}
}
