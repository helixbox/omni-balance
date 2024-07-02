package bot

var (
	monitors map[string]Bot
)

func Register(name string, monitor Bot) {
	if monitors == nil {
		monitors = make(map[string]Bot)
	}
	monitors[name] = monitor
}

func ListMonitors() map[string]Bot {
	return monitors
}

func GetMonitor(name string) Bot {
	return monitors[name]
}
