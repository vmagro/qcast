package libqcast

type Storage interface {
	Get(key string)
	Set(key string, value interface{})
}

type boltsStorage struct {
}
