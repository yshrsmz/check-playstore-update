# go get github.com/urfave/cli
# go get github.com/nlopes/slalck
# go get gopkg.in/yaml.v2
#
# for GOOS in darwin linux; do
#     export GOOS=$GOOS
#     for GOARCH in 386 amd64; do
#         export GOARCH=$GOARCH
#         go build -v -o bin/check-playstore-update-$GOOS-$GOARCH
#     done
# done

export GOARCH=386
go build -v -o bin/check-playstore-update-linux-386
