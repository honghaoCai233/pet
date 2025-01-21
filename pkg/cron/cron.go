package cron

import "github.com/robfig/cron/v3"

const stdCronParse = cron.SecondOptional |
	cron.Minute |
	cron.Hour |
	cron.Dom |
	cron.Month |
	cron.Dow

func New() *cron.Cron {
	return cron.New(
		cron.WithParser(
			cron.NewParser(
				stdCronParse,
			),
		),
	)
}
