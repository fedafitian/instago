# instago Usage

## InstaGo

Tool for collecting data analytics for an instagram account

Usage:
  instago [command]

Available Commands:
  help        Help about any command
  login       Logs into an instagram accout with user credentials
  track       Track a specific metric related to your account

Flags:
      --config string   config file (default is $HOME/.instago.env)
  -h, --help            help for instago
  -t, --toggle          Help message for toggle

Use "instago [command] --help" for more information about a command.

## Track

Track a specific metric related to your account

Usage:
  instago track [flags]
  instago track [command]

Available Commands:
  fans        Checks which users follow you that you don't follow back
  ghosts      Checks which users don't follow you back

Flags:
  -h, --help   help for track

Global Flags:
      --config string   config file (default is $HOME/.instago.env)

Use "instago track [command] --help" for more information about a command.