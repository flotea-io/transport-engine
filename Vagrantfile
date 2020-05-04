require 'yaml'
require 'fileutils'

required_plugins = %w( vagrant-hostmanager vagrant-vbguest )
required_plugins.each do |plugin|
    system("vagrant plugin install #{plugin}", :chdir=>"/tmp") || exit! unless Vagrant.has_plugin?(plugin)
end

domains = {
  phppgyadmin: 'phppgadmin.engine.devel',
  api: 'api.engine.devel',
  wss: 'wss.engine.devel',
  frontendDevelop: 'frontend.engine.devel'
}

vagrantfile_dir_path = File.dirname(__FILE__)

config = {
  local: vagrantfile_dir_path + '/vagrant/config/vagrant-local.yml',
}

options = YAML.load_file config[:local]

Vagrant.configure(2) do |config|
  #config.vm.box = 'ubuntu/disco64'
  config.vm.box = 'geerlingguy/ubuntu2004'
  #config.vm.box_url = "http://cloud-images.ubuntu.com/releases/disco/release/ubuntu-19.04-server-cloudimg-amd64-vagrant.box"
#  config.vm.box_version = "20191217.0.0"
config.vm.box_check_update = options['box_check_update']
config.vbguest.auto_update = false
  config.vm.provider 'virtualbox' do |vb|
    vb.cpus = options['cpus']
    vb.memory = options['memory']
    vb.name = options['machine_name']
  end

  config.vm.define options['machine_name']
  config.vm.hostname = options['machine_name']
  config.vm.network 'private_network', ip: options['ip']
  config.vm.synced_folder './', '/app', owner: 'vagrant', group: 'vagrant'
  #config.vm.synced_folder './src/flt', '/app/src/flt', owner: 'vagrant', group: 'vagrant', type: "rsync"
  config.vm.synced_folder './', '/vagrant', disabled: true

  config.vm.provision :hostmanager
  config.hostmanager.enabled            = true
  config.hostmanager.manage_host        = true
  config.hostmanager.ignore_private_ip  = false
  config.hostmanager.include_offline    = true
  config.hostmanager.aliases            = domains.values

  config.vm.provision 'shell', path: './vagrant/provision/once-as-root.sh', args: [options['timezone']]
  config.vm.provision 'shell', path: './vagrant/provision/once-as-vagrant.sh', privileged: false
  config.vm.provision 'shell', path: './vagrant/provision/always-as-root.sh', run: 'always'

end
