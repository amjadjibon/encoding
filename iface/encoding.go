package iface

type IMarshal interface {
	Marshal(v interface{}) ([]byte, error)
}

type IUnmarshal interface {
	Unmarshal(b []byte, v interface{}) error
}

type IEncoding interface {
	IMarshal
	IUnmarshal
}
