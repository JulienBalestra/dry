package exit

import "go.uber.org/zap"

func Exit(err error) int {
	defer zap.L().Sync()
	exit := 0
	if err != nil {
		exit = 1
		zap.L().Error("exit", zap.Error(err), zap.Int("code", exit))
		return exit
	}
	zap.L().Info("exit", zap.Int("code", exit))
	return exit
}
