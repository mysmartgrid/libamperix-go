# libamperix-go

## Installing the libamperix software

If you're new to golang: an important concept is the "workspace" which contains all the libraries, 
project files etc. On a plain machine, I usually do the following:

	$ mkdir -p <dir>/go/{src|bin|pkg}

Now, add the following two lines to your .profile:

	export GOPATH="<dir>/go"
	export PATH="$GOPATH/bin:$PATH"


Your workspace is now set up. The dependencies of the software are managed using [godep](https://github.com/tools/godep). The installation is a good test whether your environment is set up correctly. Simply run

	go get github.com/tools/godep

You should now have a working godep binary in your path. Continue by checking out the defluxio source:
	
	$ cd <dir>/go/src
	$ mkdir github.com
	$ cd github.com
	$ git clone https://github.com/mysmartgrid/libamperix-go.git
	$ cd libamperix-go.git
	
Now, restore the libraries used by defluxio:

	$ godep restore
	
You're all set. You can build the project using 

	$ (cd cmd/amperix_get && go install )

	$ ls <dir>/go/bin
	$ cd <dir>/go/bin
	$ ./amperix_get -help
