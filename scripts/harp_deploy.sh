set -e
{{.SyncFiles}}
{{.SaveRelease}}
source $HOME/.envrc
{{.RestartServer}}
