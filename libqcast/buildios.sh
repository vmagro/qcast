#!/bin/bash
set -e
shopt -s extglob

GODEPS=$1
OUT=$2

# this is gross hack to trick gomobile into seeing it at the right place on the gopath but we don't want to clutter the worktree
# make a fake gopath for the libqcast sources by making a src directory and symlinking the current dir under it
libqcast="$(pwd)"
rm -rf "$libqcast/src/"
mkdir -p "$libqcast/src/"
ln -s "$libqcast" "$libqcast/src/libqcast"

# make a fake gopath for the third-party godeps also
pushd "$GODEPS"
rm -rf src/
mkdir src/
pushd src
# ignore src, . and pkg when making symlinks
find .. ! -name src ! -name .. ! -name pkg -maxdepth 1 -exec ln -s {} ./ \;
third_party="$GODEPS"
popd
popd

# run gomobile with our hacked together gopaths
env GOPATH="$third_party:$libqcast" gomobile bind -tags lldb -target=ios libqcast libqcast/mobile

# make sure the file is where buck wants it
mv Libqcast.framework $OUT

# clean up our hacky symlink afterwards
rm -r src/
