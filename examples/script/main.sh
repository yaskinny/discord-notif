#!/usr/bin/env bash

set -euo pipefail

export ServerIP=`hostname -I | awk '{print $1}'`
export Owner='Yaser'
export NOTIF_FIELDS='USER,ServerIP,Owner'

_break(){
  if (( ${#} > 0 )); then
    printf >&2 '%s\n' "${@}"
  fi
  discord-notif backup script error
  exit 1
}
trap "_break" ERR

discord-notif backup script start


# this will make err
SALAM_ERRRRR
