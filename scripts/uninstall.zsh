#!/bin/zsh

# remove binary
rm $(which space)

# remove config directory
rm -rf $HOME/.config/space

echo "space removed"
