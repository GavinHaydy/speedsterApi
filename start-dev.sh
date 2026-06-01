#!/bin/bash

case "$1" in
  start)
    tmux new-session -d -s speedster

    tmux send-keys -t speedster:0 \
      'cd app/user && air' C-m

    tmux new-window -t speedster:1
    tmux send-keys -t speedster:1 \
      'cd app/role && air' C-m

    tmux new-window -t speedster:2
    tmux send-keys -t speedster:2 \
      'cd app/gateway && air' C-m

    tmux attach -t speedster
    ;;

  stop)
    tmux kill-session -t speedster
    ;;

  restart)
    tmux kill-session -t speedster 2>/dev/null
    "$0" start
    ;;
esac