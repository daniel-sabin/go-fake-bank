# Generate statik file

### Install statik

    go install github.com/rakyll/statik

### Generate file

Update your PATH

     export PATH=$PATH:$(go env GOPATH)/bin

If you are in the root of the project :

    statik -src=/<your-absolute-local-path>/server/swaggerui

Otherwise :

    statik -src=/<your-absolute-local-path>/server/swaggerui -dest=/<your-absolute-local-path>
