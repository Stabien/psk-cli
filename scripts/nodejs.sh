#!/bin/bash

folderName="template_nodejs"

if [ $# -eq 0 ]
  then
    git clone "git@github.com:Stabien/template_nodejs"
else
    folderName="$1"
    git clone "git@github.com:Stabien/template_nodejs" "$folderName"
fi

cd "$folderName"

npm install dotenv express jest jsonwebtoken prettier nodemon body-parser

rm -rf .git
git init

npm init -y