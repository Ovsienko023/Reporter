package configuration

import "flag"

var ConfigPathFlag = flag.String("configuration", DefaultConfigPath, "Path to the configuration file. -configuration=/path/to/configuration")
