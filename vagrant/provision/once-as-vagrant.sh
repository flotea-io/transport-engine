#!/usr/bin/env bash
function info {
  echo " "
  echo "-> $1"
  echo " "
}

info "Provision-script user: `whoami`"

info "Create bash-alias 'app' for vagrant user"
echo 'alias app="cd /app"' | tee /home/vagrant/.bash_aliases

info "Enabling colorized prompt for guest console"
sed -i "s/#force_color_prompt=yes/force_color_prompt=yes/" /home/vagrant/.bashrc

info "Instaling beego"
go get -u github.com/astaxie/beego
go get -u github.com/beego/bee

info "Instaling go packages"
echo "github.com/astaxie/beego/migration"
go get -u github.com/astaxie/beego/migration
echo "github.com/lib/pq"
go get -u github.com/lib/pq
echo "github.com/tidwall/gjson"
go get -u github.com/tidwall/gjson
echo "github.com/xeipuuv/gojsonschema"
go get -u github.com/xeipuuv/gojsonschema
echo "github.com/ethereum/go-ethereum"
go get -u github.com/ethereum/go-ethereum
echo "github.com/gorilla/websocket"
go get -u github.com/gorilla/websocket
echo "gitlab.com/microo8/plgo/plgo"
go get -u gitlab.com/microo8/plgo/plgo

go get -u "github.com/davecgh/go-spew/spew"
go get -u "github.com/deckarep/golang-set"
go get -u "github.com/edsrzf/mmap-go"
go get -u "github.com/gballet/go-libpcsclite"
go get -u "github.com/golang/protobuf/proto"
go get -u "github.com/hashicorp/golang-lru"
go get -u "github.com/huin/goupnp"
go get -u "github.com/huin/goupnp/dcps/internetgateway1"
go get -u "github.com/huin/goupnp/dcps/internetgateway2"
go get -u "github.com/jackpal/go-nat-pmp"
go get -u "github.com/karalabe/usb"
go get -u "github.com/olekukonko/tablewriter"
go get -u "github.com/pborman/uuid"
go get -u "github.com/prometheus/tsdb/fileutil"
go get -u "github.com/rjeczalik/notify"
go get -u "github.com/status-im/keycard-go/derivationpath"
go get -u "github.com/syndtr/goleveldb/leveldb"
go get -u "github.com/syndtr/goleveldb/leveldb/errors"
go get -u "github.com/syndtr/goleveldb/leveldb/filter"
go get -u "github.com/syndtr/goleveldb/leveldb/iterator"
go get -u "github.com/rs/cors"
go get -u "github.com/tyler-smith/go-bip39"
go get -u "github.com/wsddn/go-ecdh"

cat >> ~/.profile <<EOF
export PATH="/home/vagrant/go/bin:$PATH"
export GOPATH="/home/vagrant/go:/app"
EOF

. ~/.profile

chmod 777 /app/src/flt/migrate
. /app/src/flt/migrate


