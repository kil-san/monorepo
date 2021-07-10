#!/usr/bin/env sh

while [ true ]
do
  kubectl exec svc/firebase-emulator -- firebase emulators:export firebase_data/export --project testing-microservices-1d7f8 --force
  sleep 60
done

