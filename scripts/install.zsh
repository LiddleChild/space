#!/bin/zsh

# check for git
if ! command -v git &> /dev/null; then
  echo "git is required to install space"
  exit 1
fi

# check for go
if ! command -v go &> /dev/null; then
  echo "go is required to install space"
  exit 1
fi

# clone
git clone https://github.com/LiddleChild/space.git /tmp/space

# build and put binary in path
cd /tmp/space/cmd/space; go install

# copy necessary scripts
mkdir -p $HOME/.config/space/scripts
cp /tmp/space/scripts/* $HOME/.config/space/scripts/

# setup .zshrc
cat /tmp/space/scripts/rc.zsh >> $HOME/.zshrc

# clean up
rm -rf /tmp/space

echo "space is installed"
