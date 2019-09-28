# Golang Bcrypt-Blowfish

Generando password seguros Bcrypt/Blowfish en GO.

## Instalación

Instalamos Go y seteamos la variable de entorno.

	$ sudo yum install golang
	$ mkdir $HOME/go
	$ export GOPATH=$HOME/go
	$ export PATH=$PATH:$GOPATH/bin
  
Instalamos mercurial, para poder descargar el package bcrypt del repositorio de Google.

  $ sudo yum install mercurial

Si se presenta un problema para instalar Mercurial, puedes hacer lo siguiente:

  $ sudo yum install python-setuptools python-devel gcc -y
  $ sudo easy_install Mercurial
