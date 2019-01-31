package zap

func (c *Core) Debug(args ...interface{}) {
	c.locker.RLock()
	c.locker.RUnlock()
	c.globalLogger.Debug(args...)
}

func (c *Core) Info(args ...interface{}) {
	c.locker.RLock()
	c.locker.RUnlock()
	c.globalLogger.Info(args...)
}

func (c *Core) Warn(args ...interface{}) {
	c.globalLogger.Warn(args...)
}

func (c *Core) Error(args ...interface{}) {
	c.globalLogger.Error(args...)
}

func (c *Core) Fatal(args ...interface{}) {
	c.globalLogger.Fatal(args...)
}

func (c *Core) Panic(args ...interface{}) {
	c.globalLogger.Panic(args...)
}

func (c *Core) DPanic(args ...interface{}) {
	c.globalLogger.DPanic(args...)
}

func (c *Core) Debugf(format string, args ...interface{}) {
	c.globalLogger.Debugf(format, args...)
}

func (c *Core) Infof(format string, args ...interface{}) {
	c.globalLogger.Infof(format, args...)
}

func (c *Core) Warnf(format string, args ...interface{}) {
	c.globalLogger.Warnf(format, args...)
}

func (c *Core) Errorf(format string, args ...interface{}) {
	c.globalLogger.Errorf(format, args...)
}

func (c *Core) Fatalf(format string, args ...interface{}) {
	c.globalLogger.Fatalf(format, args...)
}

func (c *Core) Panicf(format string, args ...interface{}) {
	c.globalLogger.Panicf(format, args...)
}

func (c *Core) DPanicf(format string, args ...interface{}) {
	c.globalLogger.DPanicf(format, args...)
}

func (c *Core) Debugw(msg string, keysAndValues ...interface{}) {
	c.globalLogger.Debugw(msg, keysAndValues...)
}

func (c *Core) Infow(msg string, keysAndValues ...interface{}) {
	c.globalLogger.Infow(msg, keysAndValues...)
}

func (c *Core) Warnw(msg string, keysAndValues ...interface{}) {
	c.globalLogger.Warnw(msg, keysAndValues...)
}

func (c *Core) Errorw(msg string, keysAndValues ...interface{}) {
	c.globalLogger.Errorw(msg, keysAndValues...)
}

func (c *Core) Fatalw(msg string, keysAndValues ...interface{}) {
	c.globalLogger.Fatalw(msg, keysAndValues...)
}

func (c *Core) Panicw(msg string, keysAndValues ...interface{}) {
	c.globalLogger.Panicw(msg, keysAndValues...)
}

func (c *Core) DPanicw(msg string, keysAndValues ...interface{}) {
	c.globalLogger.DPanicw(msg, keysAndValues...)
}
