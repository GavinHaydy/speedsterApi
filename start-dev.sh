#!/bin/bash

case "$1" in
  start)
    tmux new-session -d -s speedster -n gateway

    tmux send-keys -t speedster:gateway \
      'cd app/gateway && air' C-m

    tmux new-window -t speedster -n user-api
    tmux send-keys -t speedster:user-api \
      'cd app/user/api && air' C-m

    tmux new-window -t speedster -n user-rpc
    tmux send-keys -t speedster:user-rpc \
      'cd app/user/user && air' C-m

    tmux new-window -t speedster -n role-api
    tmux send-keys -t speedster:role-api \
      'cd app/role/api && air' C-m




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