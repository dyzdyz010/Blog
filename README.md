Blog
====

A blog platform.

Warning
------

**This repo is under active development, and not fully functional yet, don't use it for production environment.**

Prerequisites
-------

1. ruby-compass
2. SSDB
3. go
4. bower

Installation
-------
> WARNING  
> Only have tested on Ubuntu 14.04, which has relatively new version of the
> prerequisites.

### Install Go Language
The golang Debian package have already been included in Ubuntu 14.04 LTS.

``` 
sudo apt-get install golang
```
If you are running in a lower version, you could manually add the ppa of golang and install from that repository as follows:
```
sudo add-apt-repository ppa:gophers/go
sudo apt-get update
sudo apt-get install golang-stable
```
(If you don't have add-apt-repository, run "`sudo apt-get install python-software-properties`".) 

### Install Bower
Bower depends on Node and npm, hence we need to resolve these first.
Ubuntu 14.04 ships with node.js, just install it via
```
sudo apt-get install nodejs
```
> WARNING  
> node in Ubuntu 14.04 is a completely different package which has absolutely
> no connection with nodejs. Make sure you did not install the package node
> previously.
In Ubuntu 14.04, node.js uses the binary command `nodejs`, not just `node` like
in some other platforms. However, bower really needs nodejs to be node to
function correctly. Hence, you need to manually create a symbolic link
`/usr/bin/node`
```
sudo ln -s /usr/bin/nodejs /usr/bin/node
```

