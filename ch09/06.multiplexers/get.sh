#!/bin/bash

echo "# Getting / fetches the root handler"
curl -i -X GET -H 'Content-Type: application/json' http://localhost:18081

echo ""
echo ""
echo "# Getting /hello fetches the /hello handler"
curl -i -X GET -H 'Content-Type: application/json' http://localhost:18081/hello

echo ""
echo ""
echo "# Getting /hello/there/ fetches the /hello/there/ handler"
curl -i -X GET -H 'Content-Type: application/json' http://localhost:18081/hello/there/

echo ""
echo ""
echo "# Getting /hello/there redirects to the /hello/there/ handler"
curl -i -X GET -H 'Content-Type: application/json' http://localhost:18081/hello/there

echo ""
echo ""
echo "# Getting /hello/there/you redirects to the /hello/there/ handler"
curl -i -X GET -H 'Content-Type: application/json' http://localhost:18081/hello/there/you

echo ""
echo ""
echo "# Getting /hello/and/goodbye redirects to the / handler"
curl -i -X GET -H 'Content-Type: application/json' http://localhost:18081/hello/and/goodbye
