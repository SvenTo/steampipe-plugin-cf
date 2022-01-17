connection "cf" {
  plugin = "hub.steampipe.io/plugins/svento/cf@latest"

  # Currently, no configuration parameters are implemented
  # The plugin is using the access token from the cf CLI configuration
  # file (typically ``~/.cf/config.json``).
  #
  # Run ``cf login`` before steampipe to get an access token for the plugin.
}
