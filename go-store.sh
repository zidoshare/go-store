#! /bin/bash

SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ]; do 
  DIR="$( cd -P "$( dirname "$SOURCE" )" && pwd )"
  SOURCE="$(readlink "$SOURCE")"
  [[ $SOURCE != /* ]] && SOURCE="$DIR/$SOURCE" 
done
DIR="$( cd -P "$( dirname "$SOURCE" )" && pwd )"

echo "current dir : "$DIR
sub(){
  echo "checking submodules..."
  cd $DIR
  if [[ ! -d "go-store-client" || ! -d "go-store-admin-client" ]] 
  then
  echo "submodules are not exists,init submodules"
  git submodule init
  git submodule update
  fi

  if [[ $(ls -A go-store-client) == "" || $(ls -A go-store-client) == "" ]] 
  then
  echo "update submodules"
  git submodule init
  git submodule update
  fi
}

checkNodeModules(){
  cd $DIR
  cd $1
  if [[ ! -d "node_modules" ]]
  then
  echo "download update node modules"
  yarn install
  fi
}

runServer(){
  cd $DIR
  go run main.go
}

runClient(){
  sub
  checkNodeModules go-store-client
  cd $DIR/go-store-client
  npm run start
}

runAdminClient(){
  sub
  checkNodeModules go-store-admin-client
  cd $DIR/go-store-admin-client
  npm run start
}

pushCommit(){
  cd $DIR/
  git push --recurse-submodules=on-demand
}

case "$1" in 
  sub) 
  sub
  exit 1
  ;;
  run:client)
  runClient
  ;;
  push)
  pushCommit
  ;;
  run:server)
  runServer
  ;; 
  run:client:admin)
  runAdminClient
  ;; 
  *) 
  echo "command: sub|run:client|run:client:admin|run:server|push" 
  exit 1 
  ;;
  esac