OpenShift cartridge Go, Martini and mgo MongoDB driver. 
======================

Runs [Go 1.3](http://golang.org), [Martini](https://github.com/go-martini/martini) and [mgo](https://gopkg.in/mgo.v2) on [OpenShift](https://openshift.redhat.com/app/login) using downloadable cartridge support.  To install to OpenShift from the CLI (you'll need version 1.9 or later of rhc), run:

    rhc create-app mygo https://cartreflect-claytondev.rhcloud.com/reflect?github=thomasmodeneis/openshift-go-cart

Once the app is created, you'll need to create and add a ".godir" file in your repo to tell the cartridge what the package of your Go code is.  A typical .godir file might contain:

    github.com/thomasmodeneis/openshift-go-cart

which would tell OpenShift to place all of the files in the root of the Git repository inside of the <code>github.com/thomasmodeneis/openshift-go-cart</code> package prior to compilation.

When you push code to the repo, the cart will compile your package into <code>$OPENSHIFT_REPO_DIR/bin/</code>, with the last segment of the .godir being the name of the executable.  For the above .godir, your executable will be:

    $OPENSHIFT_REPO_DIR/bin/goexample

If you want to serve web requests (vs. running in the background), you'll need to listen on the ip address and port that OpenShift allocates - those are available as HOST and PORT in the environment.

The repository contains a sample go file which will print "hello, world" when someone hits your web application - see [web.go](https://github.com/thomasmodeneis/openshift-go-cart/blob/master/template/web.go).

Any log output will be generated to <code>$OPENSHIFT_GO_LOG_DIR</code>



MongoDB
------------
Note that you need to install mongodb locally or update web.go file with a valid mongodb url credential --> mongodb://user:pass@mongolab.com mongodb://user:pass@mongohq.com or any other provider).


How it Works
------------

When you push code to your repo, a Git postreceive hook runs and invokes the bin/compile script.  This attempts to download a Go environment for you into $OPENSHIFT_GO_DIR/cache.  Once the environment is setup, the cart runs

    go get -tags openshift ./...

on a working copy of your source.  The main file that you run will have access to two environment variables, $HOST and $PORT, which contain the internal address you must listen on to receive HTTP requests to your application.


Credits
-------

The bin/compile script is based on the [Heroku Go buildpack](https://github.com/kr/heroku-buildpack-go), adapted for OpenShift cartridges.
