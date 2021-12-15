# Skyline Commander

## Requirements

  * Golang
  * libncurses

## Source code compile
  
### Prepare Debian bases distributions

    $ sudo apt-get install pkg-config
    $ go build
	
### Prepare MacOS X system

    $ brew install ncurses
    $ export PKG_CONFIG_PATH="/usr/local/opt/ncurses/lib/pkgconfig"
    $ go build
    
