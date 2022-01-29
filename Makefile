install:
	go build -o ~/.steampipe/plugins/hub.steampipe.io/plugins/svento/cf@latest/steampipe-plugin-cf.plugin *.go

install-debug:
	go build -gcflags=all="-N -l" -o ~/.steampipe/plugins/hub.steampipe.io/plugins/svento/cf@latest/steampipe-plugin-cf.plugin *.go