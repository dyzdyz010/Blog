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
    sudo apt-get install golang
If you are running in a lower version, you could manually add the ppa of golang and install from that repository as follows:
    sudo add-apt-repository ppa:gophers/go
    sudo apt-get update
    sudo apt-get install golang-stable
(If you don't have add-apt-repository, run "sudo apt-get install python-software-properties".) 

### Install Bower
Bower depends on Node and npm, hence we need to resolve these first.


