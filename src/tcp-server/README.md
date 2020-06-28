# Go TCP server

https://www.infoq.com/articles/build-a-container-golang/
https://github.com/mactsouk/opensource.com/blob/master/concTCP.go

namespaces, cgroups and layered filesystems

## Run ubuntu
multipass launch focal -n ubuntudev --cloud-init ../multipass/dev.yml
multipass mount $HOME/dev ubuntudev:dev
multipass exec ubuntudev -- /bin/bash

or 

vagrant up
