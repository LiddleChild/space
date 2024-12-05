if [[ -f $SPACE_DIR/scripts/zsh-completion.zsh && ! -v SPACE_SOURCED ]]; then
  source $SPACE_DIR/scripts/zsh-completion.zsh
  export SPACE_SOURCED=1
fi
