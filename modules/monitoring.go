package modules

type Monitoring struct {
	url  string
	port int
}

func NewMonitoring() Monitoring {
	mon := Monitoring{url: "localhost", port: 4533}
	return mon
}
