package service

import "github.com/spf13/viper"

type Services struct {
	Scheduler  *Scheduler
	Monitoring *Monitoring
}

func GetServices(l Logger, c *viper.Viper) *Services {
	return &Services{
		Scheduler:  NewScheduler(l),
		Monitoring: NewMonitoring(l, c),
	}
}
