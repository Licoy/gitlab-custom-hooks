#!/bin/bash

GITLAB_HOME="/home/gitlab/data"
RULES_HOME="/home/commit-rules"

run(){
        name=`echo -n $1 | sha256sum`
        name=${name:0:64}
        path1=${name:0:2}
        path2=${name:2:2}
        path="$GITLAB_HOME/git-data/repositories/@hashed/${path1}/${path2}/${name}.git/custom_hooks/"
        rm -rf "$path"
        mkdir "$path"
        cp -rf "$RULES_HOME/pre-receive" "$path"
        echo "Success to: ${path}"
}

run "$1"