connection "cf" {
  plugin = "hub.steampipe.io/plugins/svento/cf@latest"

  # You may connect to the Cloud Foundry API using more than one option:
  #
  # 1. Empty configuration: The plugin is using the access token from the
  # cf CLI configuration file (typically ``~/.cf/config.json``) in default.
  #
  # Run ``cf login`` before Steampipe to get an access token for the plugin.

  # 2. Credentials from the config.json in the directory specified by the ``cf_home_dir`` parameter.
  # cf_home_dir         = "/path/to/my/cf/home/directory"

  # 3. Use credentials explicitly set in a Steampipe config file by setting the option:
  # api_url = "https://api.cf.example.com"
  # 
  # and either one of the following authentication mechanisms:
  # 
  # 3.1. with username and password:
  # user = "user@example.com"
  # password = "example"

  # 3.2. with client_id / client_secret:
  # client_id = "myclientid"
  # client_secret = "38df0e66e4904a75ac951720e4b9df02"

  # 3.3. auth_token:
  # auth_token = "ey[...]"

  # Misc options:
  # skip_ssl_validation = false
  # user_agent = "Steampipe"
}
