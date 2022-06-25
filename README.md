# Image Docker for Machine Learning Model Scripts

This project allows its consumers to encapsulate their scripts at a call of an API Endpoint. 

### Run it in Docker
```
docker build -t mlcicd .
```
Use following command so your local docker daemon can be accessed from within the container

```
docker run -ti -v /var/run/docker.sock:/var/run/docker.sock -p 5433:5433 mlcicd
```

# API

POST /api/publish
    formdata