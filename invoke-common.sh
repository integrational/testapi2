# assume PORT

echo
echo "findPets:"
curl -i http://localhost:$PORT/pets

echo
echo "findPets with query params:"
curl -i http://localhost:$PORT/pets?limit=1\&tags=cat\&tags=dog

echo
echo "addPet:"
curl -i http://localhost:$PORT/pets -X POST -H 'Content-Type: application/json' -d '{"name":"Moritz","tag":"cat"}'

echo
echo "find pet by id:"
curl -i http://localhost:$PORT/pets/1

echo
echo "deletePet:"
curl -i http://localhost:$PORT/pets/1 -X DELETE
