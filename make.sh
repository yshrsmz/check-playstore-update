# go get github.com/urfave/cli
# go get github.com/nlopes/slack
# go get github.com/roylee0704/gron
# go get gopkg.in/yaml.v2
#
for GOOS in darwin linux; do
    export GOOS=$GOOS
    for GOARCH in 386 amd64; do
        export GOARCH=$GOARCH
        go build -v -o bin/check-playstore-update-$GOOS-$GOARCH
    done
done

# debug
# export GOARCH=386
# go build -v -o bin/check-playstore-update-linux-386
