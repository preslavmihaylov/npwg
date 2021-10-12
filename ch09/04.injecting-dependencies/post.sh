#!/bin/bash

curl -i -X POST -H 'Content-Type: application/json' \
  --data '<world>' \
  http://localhost:18081

