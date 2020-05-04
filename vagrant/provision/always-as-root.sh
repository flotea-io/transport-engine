#!/usr/bin/env bash

function info {
  echo " "
  echo "--> $1"
  echo " "
}

info "Restart apps"
sudo systemctl restart postgresql
sudo systemctl restart apache2
echo "PostgreSQL... Done!"

info "Main - Engine Dev - You probably can try do something now"
echo "Provision-script user: `whoami`"
echo "IP: 192.168.66.66"
echo "PhpPgAdmin: http://phppgadmin.engine.devel/"
