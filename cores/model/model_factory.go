package model

type ModelFactory struct {
}

func (m *ModelFactory) ModelClientFactory(modelConfStr string) *ModelClient {

	fi := new(ModelClient)
	fi.ConfigLoad("", "", modelConfStr)

	return fi
}
