echo "Check fast_fake_rest image exists.";
if [[ "$(docker images -q fast_fake_rest 2> /dev/null)" == "" ]]; then
  echo "Building fast_fake_rest image.";
  docker build --no-cache -t zoer/fast_fake_rest .;
fi
echo "OK";


echo "Start docker-compose up";
docker-compose up -d;
echo "OK";

