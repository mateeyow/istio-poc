#!/usr/bin/env bash

while [ true ]; do
  curl https://istioweb.rapiscan.systems -Isk
  curl https://istioweb.rapiscan.systems/admin-api 
  echo -e "\nsleeping... \n"
  sleep 1
done