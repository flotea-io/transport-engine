#!/usr/bin/env bash

timezone=$(echo "$1")

function info {
  echo " "
  echo "-> $1"
  echo " "
}

#== Provision script ==
info "Provision-script user: `whoami`"
info "Configure timezone"
timedatectl set-timezone ${timezone} --no-ask-password
apt-get update

info "Install additional software"
sudo DEBIAN_FRONTEND=noninteractive apt-get install -y git curl ca-certificates apache2 --assume-yes phppgadmin --assume-yes build-essential
echo "Done!"
15.6. 10:30

info "Install PostgreSQL"
sudo sh -c 'echo "deb [arch=amd64] http://apt.postgresql.org/pub/repos/apt focal-pgdg main" >> /etc/apt/sources.list'
wget --quiet -O - http://apt.postgresql.org/pub/repos/apt/ACCC4CF8.asc | sudo apt-key add -
apt-get update
DEBIAN_FRONTEND=noninteractive apt-get install -yq postgresql-11-postgis-2.5
apt-get install -y postgresql-server-dev-11
echo "Done!"

info "Configure PostgreSQL database"
# fix permissions
sed -i "s/#listen_address.*/listen_addresses = '*'/" /etc/postgresql/11/main/postgresql.conf
#fixing postgres pg_hba.conf file
cat >> /etc/postgresql/11/main/pg_hba.conf <<EOF
#Accept all IPv4 connections - FOR DEVELOPMENT ONLY!!!
host    all         all         0.0.0.0/0             md5
EOF
echo "Done!"

info "PostgreSQL: Create User and Database for dev and tests"
# create user and database
su postgres -c "psql -c \"CREATE ROLE flotea SUPERUSER LOGIN PASSWORD 'Fckgw-Rhqq2'\" "
su postgres -c "createdb -E UTF8 -T template0 --locale=en_US.utf8 -O flotea flotea_platform"
#create database for tests
su postgres -c "createdb -E UTF8 -T template0 --locale=en_US.utf8 -O flotea flotea_platform_test"
echo "Done!"

info "Install PostGIS PgTimestamp Extension (@todo)"

/bin/mkdir -p '/usr/share/postgresql/11/extension'
/bin/mkdir -p '/usr/lib/postgresql/11/lib'
/usr/bin/install -c -m 644 /app/src/flt/utils/pgtimespan/build/pgtimespan.control '/usr/share/postgresql/11/extension/'
/usr/bin/install -c -m 644 /app/src/flt/utils/pgtimespan/build/pgtimespan--0.1.sql  '/usr/share/postgresql/11/extension/'
/usr/bin/install -c -m 755 /app/src/flt/utils/pgtimespan/build/pgtimespan.so '/usr/lib/postgresql/11/lib/'

su postgres -c 'psql -d flotea_platform -c "CREATE EXTENSION pgtimespan;"'


#info "Install PostGIS Extensions"
#su postgres -c 'psql -d flotea_platform -c "CREATE EXTENSION postgis;"'
#su postgres -c 'psql -d flotea_platform -c "CREATE EXTENSION postgis_topology;"'
#su postgres -c 'psql -d flotea_platform -c "CREATE EXTENSION postgis_sfcgal;"'
#su postgres -c 'psql -d flotea_platform -c "CREATE EXTENSION address_standardizer;"'
#su postgres -c 'psql -d flotea_platform -c "CREATE EXTENSION postgis_tiger_geocoder CASCADE;"'
#su postgres -c 'psql -d flotea_platform -c "CREATE EXTENSION \"uuid-ossp\";"'
#su postgres -c 'psql -d flotea_platform -c "CREATE EXTENSION pgcrypto;"'

# option CASCADE install also dependency extensions like fuzzystrmatch

echo "Configure phppgadmin /phppgadmin"
sed -i "s/Require local.*/Require all granted/" /etc/apache2/conf-available/phppgadmin.conf
echo "Done!"

#info "Install RabbitMQ"

#sudo apt-key adv --keyserver "hkps.pool.sks-keyservers.net" --recv-keys "0x6B73A36E6026DFCA"
#sudo tee /etc/apt/sources.list.d/bintray.rabbitmq.list <<EOF
#deb https://dl.bintray.com/rabbitmq-erlang/debian disco erlang
#deb https://dl.bintray.com/rabbitmq/debian disco main
#EOF
# 16.04 echo "deb https://dl.bintray.com/rabbitmq/debian xenial main" | sudo tee /etc/apt/sources.list.d/bintray.rabbitmq.list
#sudo apt-get update -y
#sudo DEBIAN_FRONTEND=noninteractive apt-get install -y erlang-nox --assume-yes apt-transport-https --assume-yes
#sudo apt-get install -y rabbitmq-server

info "Install npm"
apt-get -y install npm

echo "Done!"

info "Install GOLANG"
apt-get install -y golang-go

echo "Done!"

info "Setting Apache"

rm /etc/apache2/sites-available/000-default.conf
rm /etc/apache2/sites-available/default-ssl.conf
rm /etc/apache2/sites-enabled/000-default.conf

cat <<EOF >/etc/apache2/sites-available/000-engine.devel.conf
<VirtualHost *:80>
        ServerName engine.devel
        Redirect permanent / https://engine.devel/
</VirtualHost>
<VirtualHost *:443>
    ServerName engine.devel
    DocumentRoot /app/src/frontend/dist
    SSLEngine on
    SSLCertificateFile /home/vagrant/engine.devel.cert
    SSLCertificateKeyFile /home/vagrant/engine.devel.key
    DirectoryIndex index.html
    <Directory /app/src/frontend/dist>
        Require all granted
    </directory>
</VirtualHost>
EOF

cat <<EOF >/etc/apache2/sites-available/001-api.engine.devel.conf
<VirtualHost *:80>
        ServerName api.engine.devel
        Redirect permanent / https://api.engine.devel/
</VirtualHost>
<VirtualHost *:443>
    ServerName api.engine.devel

    SSLEngine on
    SSLCertificateFile /home/vagrant/engine.devel.cert
    SSLCertificateKeyFile /home/vagrant/engine.devel.key

    ProxyRequests Off
    <Proxy *>
        Order deny,allow
        Allow from all
    </Proxy>
    ProxyPass / http://192.168.66.66:8010/
    ProxyPassReverse / http://192.168.66.66:8010/
</VirtualHost>
EOF

cat <<EOF >/etc/apache2/sites-available/002-wss.engine.devel.conf
<VirtualHost *:443>
    ServerName wss.engine.devel

    SSLEngine on
    SSLCertificateFile /home/vagrant/engine.devel.cert
    SSLCertificateKeyFile /home/vagrant/engine.devel.key

    ProxyRequests Off
    <Proxy *>
        Order deny,allow
        Allow from all
    </Proxy>
    ProxyPass / ws://192.168.66.66:8010/
    ProxyPassReverse / ws://192.168.66.66:8010/
</VirtualHost>
EOF

cat <<EOF >/etc/apache2/sites-available/003-phppgadmin.engine.devel.conf
<VirtualHost *:80>
    ServerName phppgadmin.engine.devel
    DocumentRoot /usr/share/phppgadmin

    <Directory /usr/share/phpmyadmin>
        Options FollowSymLinks
        DirectoryIndex index.php

        <IfModule mod_php5.c>
            <IfModule mod_mime.c>
                AddType application/x-httpd-php .php
            </IfModule>
            <FilesMatch ".+\.php$">
                SetHandler application/x-httpd-php
            </FilesMatch>

            php_flag magic_quotes_gpc Off
            php_flag track_vars On
            php_flag register_globals Off
            php_admin_flag allow_url_fopen Off
            php_value include_path .
            php_admin_value upload_tmp_dir /var/lib/phpmyadmin/tmp
            php_admin_value open_basedir /usr/share/phpmyadmin/:/etc/phpmyadmin/:/var/lib/phpmyadmin/:/usr/share/php/php-gettext/:/usr/share/javascript/:/usr/share/php/tcpdf/
        </IfModule>
    </Directory>

    ErrorLog ${APACHE_LOG_DIR}/error.log
    CustomLog ${APACHE_LOG_DIR}/access.log combined
</VirtualHost>
EOF

cat <<EOF >/etc/apache2/sites-available/004-frontend.engine.devel.conf
<VirtualHost *:80>
        ServerName frontend.engine.devel
        Redirect permanent / https://frontend.engine.devel/
</VirtualHost>
<VirtualHost *:443>
    ServerName frontend.engine.devel

    SSLEngine on
    SSLCertificateFile /home/vagrant/engine.devel.cert
    SSLCertificateKeyFile /home/vagrant/engine.devel.key

    ProxyRequests Off
    <Proxy *>
        Order deny,allow
        Allow from all
    </Proxy>
    ProxyPass / http://192.168.66.66:3000/
    ProxyPassReverse / http://192.168.66.66:3000/
</VirtualHost>
EOF

cat <<EOF >/home/vagrant/engine.devel.cert
-----BEGIN CERTIFICATE-----
MIIDDzCCAfegAwIBAgIUbVM4t0I1Wvlh6S6lUl+Yb76HUB8wDQYJKoZIhvcNAQEL
BQAwFzEVMBMGA1UEAwwMZW5naW5lLmRldmVsMB4XDTE5MDgyODEzMjY1NFoXDTI5
MDgyNTEzMjY1NFowFzEVMBMGA1UEAwwMZW5naW5lLmRldmVsMIIBIjANBgkqhkiG
9w0BAQEFAAOCAQ8AMIIBCgKCAQEA7DQbY1+2Y489asXsykge6BYcjqcOGuiwN4Xd
+UicqhCjOZ591XO18SMJdvC+AIY3JVLaj+hiOdaaWAglfmMYH7Pga+c043Ch9IOG
q8AaxOCsxri+4wr39WH++35po2NaabXRdyhGMWE4WkNvoDj5STaW7U5mbEY/G5Sk
9WtQYA0oeU6aFhSDH+A0A4oDQ/Xln5vV2RZRQ8ECEmh5VVZ7mCF3qIGa9IQxaTU1
r1aE1xhK+/Z1nMkJJvM3KBRAGDC7cWZgcn1jrd6vuN458dkv4HkquJm4Lr/ZkZ0C
saIWI7OUyjb3JbIO8EpO52Wh4U5cd9NI6gNZJRMcwe1cXkHmCQIDAQABo1MwUTAd
BgNVHQ4EFgQUB3u+HZRObH4dOsPw1HJ2+SDYZK8wHwYDVR0jBBgwFoAUB3u+HZRO
bH4dOsPw1HJ2+SDYZK8wDwYDVR0TAQH/BAUwAwEB/zANBgkqhkiG9w0BAQsFAAOC
AQEAvvKMAL4tRYlIyRTXr6NnL0HXbAxdnKGHZjbP+IMWfN+BZTia0jeZdK2SF/Ku
2ecuB7JNWFq1pneQJ22/aTFAVw6eRi0KOy0bcRRM8xKQd7laQxy2sRA71/63dmQf
l7Pt8lXIrAGOKIcL3GSj/i5ZoKrayJLW3y0YWgsB4h62bq/UI4wgPE2y3c4mZ2d2
zXIPpcd7FWUXrbWJZZUofRaA42pZRoLyCEn3PZvIjMCIuzpHO6jdUqMHTVkuKuiq
XAS5NBe0qGT+xDsD9Au1Fgf1BWmjFuhDGrMfuBa1ZoCe9QYkKIv+epZU5USP81mD
7DxeIKrx2vMse5hkSjhtI7uM5Q==
-----END CERTIFICATE-----
EOF

cat <<EOF >/home/vagrant/engine.devel.key
-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEA7DQbY1+2Y489asXsykge6BYcjqcOGuiwN4Xd+UicqhCjOZ59
1XO18SMJdvC+AIY3JVLaj+hiOdaaWAglfmMYH7Pga+c043Ch9IOGq8AaxOCsxri+
4wr39WH++35po2NaabXRdyhGMWE4WkNvoDj5STaW7U5mbEY/G5Sk9WtQYA0oeU6a
FhSDH+A0A4oDQ/Xln5vV2RZRQ8ECEmh5VVZ7mCF3qIGa9IQxaTU1r1aE1xhK+/Z1
nMkJJvM3KBRAGDC7cWZgcn1jrd6vuN458dkv4HkquJm4Lr/ZkZ0CsaIWI7OUyjb3
JbIO8EpO52Wh4U5cd9NI6gNZJRMcwe1cXkHmCQIDAQABAoIBAEs4wOwhxAzqxg1l
4OX+l9EjkY6ghu5s4gmcmTdVN4c1azXjUoGSqwOTO+Vj/65dD9zUCQTBZd6ziE22
snIQjtMxzquuYvp3mSwsYsv5jszaHEvg8/GMQbEktzZSaMP+1QFLRNzkt43sPy0H
wtnMqx70wpXqXNXofRqEeRArT9p2WYspKOJ5NF12B4cCUdjxlvMWF0lviqEAwHQX
sPJyj/N0bxajyq1nqj+r7kx2sOJ0LjkuxUAyOKbQy0U1Z+2G+m9MqpIsjoqK84UK
KG4JlAf1rOhbFRxcJGjInmQIAgPyEjaZSRcyu6DCaBdHHorpiyW9Czx+VQ9+Y21a
ODcIhQ0CgYEA/4p7n7soauOBFJYx7By8avq5JA9+TRxhfIcEh6VFI9/qRNhtYt0Q
3eYhAXW1vxh+R/13qoUU9s8qAFhJpTZ1R/ml9k9G6E5a3IqFAq7TxBK3fUhSoTWD
CcfZTGkQDIxKYR1gDRBLcatTo5/v5t7pGDuAwux6m+HAqYTncY6jcA8CgYEA7KC7
NMtGrgSU7t3DItO5x+kVP+e1YTuCSifA23K4WWZZSnK5qz2hOqtrRGYdxOifnwLM
+GpLsP0WR3sctfQIs6B8NLjLAACSSRrTRG7ShGC6sAa6+9kyNF5p5kiwL7COCL+y
vwyVeM8vyvczt9HiGFhVzHaUHX2RJElUf1WzMGcCgYEAr9B1kcZQ3om8a4+StWJn
+CqcasUGHsOiBu1WHoWE9lO+eLN3Pwfd7O3CU6DK9LlDWQB47qc1b0HFIgucT8ES
G/0VREvy+y29sfLT5K6oIETIXbf+sc+BYp9PUBUQ57E2E28EgLhvT7uShgcqJWvE
b9mwk2vkR4+g4IECQ8zhgd0CgYEApDiMlAA7pkvBvOElK0d7mNUXjUlW5Qsyyho6
y1IbxNe2QmfKmaBlA92xzT99i4nSNb3w2LqUm1maG7PCYjjhzSAk80YIhJDg9WFr
nAO5kGu9RKr3HBBJunSH1G0/8Wj7ufCosdSe68AGsDrNptY/rpunuAXXl3vsvh0y
SkkE2iMCgYB45/skWZeWSZjvCWjOTa11a5xGfaOHierSegnNPXdP/2tGeVaQU2ac
G8XTwuIqKwgfsnJCBGsv74F4adyquYbab9Tzeuu7qdywt5khKwPHnwPnYk/uirKp
5MXfxvK3eT5D/Vm1QGB8OlM2IW3YCHTLx3qM7/JCr+XlVSlMrAqtcg==
-----END RSA PRIVATE KEY-----
EOF


a2ensite 000-engine.devel.conf
a2ensite 001-api.engine.devel.conf
a2ensite 002-wss.engine.devel.conf
a2ensite 003-phppgadmin.engine.devel.conf
a2enmod ssl
a2enmod proxy
a2enmod proxy_http

service apache2 restart

echo "Done!"