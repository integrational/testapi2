export PORT=9004

echo
echo "health:"
curl -i http://localhost:$PORT/__health

./invoke-common.sh
