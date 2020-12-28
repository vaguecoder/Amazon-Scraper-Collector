# Amazon-Scraper-and-Collector-Using-Go-REST-API

> :warning: This is intended for the purpose of demonstration of interaction between multiple REST APIs that are developed using Go, and basic CRUD operations on MongoDB. The output of the application might not be useful, real-time.

## What does this project do? :bulb:
This is a simple Go project that scrapes data (item name, image URL, price, description and total number of reviews posted so far) from the largest E-commerce website, [Amazon.com](https://www.amazon.com/) and stores in database. Following are the modules of this project:

#### 1. [Scraper-API](https://github.com/VagueCoder/Amazon-Scraper-Collector/tree/master/scraper-api)
Scraper-API, as explained, takes Amazon page URL as input using POST request and scrapes the required details. And makes an internal call to next module, [Collector-API](https://github.com/VagueCoder/Amazon-Scraper-Collector/tree/master/collector-api) using another POST request with the scraped details as form data.

#### 2. [Collector-API](https://github.com/VagueCoder/Amazon-Scraper-Collector/tree/master/collector-api)
Collector-API gets triggered internally by [Scraper-API](https://github.com/VagueCoder/Amazon-Scraper-Collector/tree/master/scraper-api) to collect the so fetched details and place in database. This makes calls to [MongoDB](https://hub.docker.com/_/mongo) module internally for storing in database.

#### 3. [MongoDB](https://hub.docker.com/_/mongo)
MongoDB module is container created of the [official mongo SDK image](https://hub.docker.com/_/mongo) that is available on [Docker Hub](https://hub.docker.com/). The [Collector-API](https://github.com/VagueCoder/Amazon-Scraper-Collector/tree/master/collector-api) calls the mongo functions to insert or retrieve the data from database.

## Knowledge (or) Technologies Used :books:
**Sno.** | **Name** | **Usage**
-------: | :------: | :--------
1 | Go (Golang) | Go is a statically typed, compiled programming language designed at Google. Helps in speed and concurrency. More info at [golang.org](https://golang.org/doc/).
2 | REST API | Representational State Transfer (REST) is a software architectural style that defines a set of constraints to be used for creating Web services. The rule Zero of using this is, you'll send everything as JSON objects over the application endpoints (API socket ports). More info at [restfulapi.net](https://restfulapi.net/).
3 | MongoDB | MongoDB is a cross-platform document-oriented database program. Classified as a NoSQL database program, MongoDB uses JSON-like documents with optional schemas. More info at [mongodb.com](https://www.mongodb.com/).
4 | Docker | Docker is a set of platform as a service products that use OS-level virtualization to deliver software in packages called containers. These containers are lightweight virtual-machine kind of platforms which make use of same host OS kernel, but are very scalable and efficient. More info at [docker.com](https://www.docker.com/).


## Base System Configurations :wrench:
**Sno.** | **Name** | **Version/Config.**
-------: | :------: | :------------------
1 | Operating System | Windows 10 x64 bit + WSL2 Ubuntu-20.04 
2 | Language | Go Version 1.14.7 Windows/amd64
3 | IDE | Visual Studio Code Version 1.52.1
4 | Containerization | Docker Version 20.10.0, Docker-Compose Version 1.27.4
5 | Database | MongoDB Version 4.4.2

> This doesn't probably make any difference in usage if having a different version, as the application works in Docker containers. But of course, the development steps may differ as per your configuration. The required softwares/configurations are mentioned under **Prerequisites** section.

## Prerequisites :file_folder:
**Sno.** | **Software** | **Detail** | **Download Links/Steps** |
-------: | :----------: | :--------: | :----------------------: |
1 | Docker Version 20.10.0 or Higher | Containerizes the application modules for using them as services. This also creates containers of Golang and MongoDB avoiding to download stand-alones for the same. | [docker.com/products/docker-desktop](https://www.docker.com/products/docker-desktop)
2 | Docker-Compose Version 1.27.4 or Higher | A CLI of Docker that helps to run [docker-compose.yml](https://github.com/VagueCoder/Amazon-Scraper-Collector/blob/master/docker-compose.yml) which is helps in building/starting/stopping all the containers at once and with ease. If using Windows (or) Mac, the Docker-Compose automatically gets downloaded along with Docker. | [docs.docker.com/compose/install/](https://docs.docker.com/compose/install/)
3 | Postman (or any equivalent) | Making the GET requests, and most importantly, the POST requests to the API are made easy with Postman. | [postman.com/downloads/](https://www.postman.com/downloads/)

> If you're using Postman, for avoiding the issue with Proxy while running this application, go to **Postman -> File -> Settings -> Proxy** and uncheck "Use the system proxy" option.

## Useful Socket-Ports: :handshake:
**Sno.** | **Port Number** | **Endpoint** | **Defined Calls**
-------: | :-------------: | :----------: | :----------------
1 | 8080 | Scraper-API | 1. [http://localhost:8080/scraper](http://localhost:8080/scraper)
2 | 8081 | Collector-API | 1. [http://localhost:8081/collector](http://localhost:8081/collector)
3 | 27017 | MongoDB | 1. mongodb://localhost:27017

> These ports are hard coded for now, but might be dynamically binded in future developments.

## Setup Application in Local :bookmark_tabs:
Following the steps to recreate the application in your local (in Docker Containers) to scrape and save to database.
1. Download the whole repo and place anywhere in the local. Go's location constraint doesn't apply as application runs in containers. 
2. Open command prompt on the same location and build using docker-compose.
```
docker-compose build
```
3. Run the application (includes all module containers along with network) at once using the command:
```
docker-compose up -d
```
The option `-d` (short form of `--detach`) here means the container runs in background.

Using docker-compose here has the following advantages:
  - It builds the services from the containers, which in turn, builds the services from our modules (Scraper-API & Collector-API) and skips to build mongo as it can access the official image (MongoDB) all by itself.
  - It runs all the services as we expect from the command.
  - Creates a network by default which helps in interaction between all the containers in the same app.
  - Takes relative paths and root accesses of the project by default. Hence the project location is not a constraint unlike the native Go workspace.
  - Downloads the images that are mentioned in Dockerfile(s) for creating containers of/out of it. Like in this app, we had made use of [mongo](https://hub.docker.com/_/mongo) and [golang](https://hub.docker.com/_/golang) official SDK images from [Docker Hub](https://hub.docker.com).
  - Also, we've included the `:latest` tag for the mongo and golang images in Dockerfiles. This checks and keeps the containers up-to-date.

Last lines if output should be:
```
Creating Collector-API ... done
Creating Scraper-API   ... done
Creating mongodb       ... done
```

4. To verify the services/containers processes running, run the following commands:
```
docker-compose ps
docker ps
```
The first command, `docker-compose ps` gives the status of all the services running in the same directory and the second, `docker ps` shows all the docker processes running on the desktop.

Output:
```
Amazon-Scraper-and-Collector-Using-Go-REST-API>docker-compose ps
    Name                  Command             State                Ports
--------------------------------------------------------------------------------------
Collector-API   ./collector-api               Up      0.0.0.0:8081->8081/tcp
Scraper-API     ./scraper-api                 Up      0.0.0.0:8080->8080/tcp, 8081/tcp
mongodb         docker-entrypoint.sh mongod   Up      0.0.0.0:27017->27017/tcp

Amazon-Scraper-and-Collector-Using-Go-REST-API>docker ps
CONTAINER ID        IMAGE               COMMAND                  CREATED             STATUS              PORTS                              NAMES
44fc708b30cc        scraper-api         "./scraper-api"          14 minutes ago      Up 13 minutes       0.0.0.0:8080->8080/tcp, 8081/tcp   Scraper-API
ab8af239575a        collector-api       "./collector-api"        14 minutes ago      Up 13 minutes       0.0.0.0:8081->8081/tcp             Collector-API
d28f9f2d0338        mongo:latest        "docker-entrypoint.s…"   14 minutes ago      Up 13 minutes       0.0.0.0:27017->27017/tcp           mongodb
```

> Make use of the application after this. The application should function. How to use, is explained under `Making the Calls` and `Using Postman` sections below.

4. Close the services/application:
```
docker-compose down
```
This closes all the services gracefully (the word it uses for non-force shutting). And removes the containers and network so created.

`Note: The automatic removal process wipes out only the containers and network that got created, and leaves the 2 images that are built (Scraper-API & Collector-API) and of course, the downloaded images (mongo & golang) remain. This is good. If not removed, the containers might collide with the new containers that will be build and might also lead to build failures.`

Output:
```
Stopping Scraper-API   ... done
Stopping Collector-API ... done
Stopping mongodb       ... done
Removing Scraper-API   ... done
Removing Collector-API ... done
Removing mongodb       ... done
Removing network amazon-scraper-and-collector-using-go-rest-api_default
```

## Making the Calls :calling:
The 4 possible calls to this application are as follows:
**Sno.** | **Port** | **Method** | **URL** | **Form Data** | **Details** | **Importance**
-------: | :------: | :--------- | :-----: | :-----------: | :---------- | :------------- 
1 | 8080 | POST | [localhost:8080/scraper](http://localhost:8080/scraper) | Yes, send Amazon page URL | The main call for the application to scrape the data from Amazon page and save to database. All other calls in this operation are internal. | HIGH
2 | 8081 | GET | [localhost:8081/collector](http://localhost:8081/collector) | Not applicable | This returns all the records (documents in Mongo terminology) in database onto page body, for cross-checking the data loading and can be used for future endpoints to retrieve data. | MEDIUM
3 | 8080 | GET | [localhost:8080/scraper](http://localhost:8080/scraper) | Not applicable | Only for debugging the endpoint. | LOW
4 | 8081 | POST | [localhost:8081/collector](http://localhost:8081/collector) | Yes, all the product details | Only for debugging the endpoint. | LOW

#### i. POST request to Scraper-API
URL: [localhost:8080/scraper](http://localhost:8080/scraper)

Form Data:
```
{
    "url":"https://www.amazon.com/PlayStation-4-Pro-1TB-Console/dp/B01LOP8EZC/"
}
```

That is just a sample URL by the time of development of this project. Use any page of Amazon that has only 1 product in it to make sure the prices and details won't collide.

Sample Output:
```
For URL: https://www.amazon.com/PlayStation-4-Pro-1TB-Console/dp/B01LOP8EZC/

Product details scraped and stored in database with ID: 5fc40bdc092ea8b4c0c2c094
```
There can be 3 possible types of output for this:
1. New Record: This inserts the data and returns the ID.
1. Existing Record: This confirms on existance of record.
1. Existing Record but updated data on Amazon page: This confirms on existance, updates the record in database and confirms that too.

#### ii. GET request to Colector-API
URL: [localhost:8080/scraper](http://localhost:8080/scraper)

Form Data: NA

Sample Output:
```
[
    {
        "_id": "5fc40bdc092ea8b4c0c2c094",
        "url": "https://www.amazon.com/PlayStation-4-Pro-1TB-Console/dp/B01LOP8EZC/",
        "product": {
            "name": "PlayStation 4 Pro 1TB Console",
            "imageURL": "https://images-na.ssl-images-amazon.com/images/I/41GGPRqTZtL._AC_SX355_.jpg",
            "description": "Heighten your experiences. Enrich your adventures. Let the super charged PS4 Pro lead the way. 4K TV Gaming : PS4 Pro outputs gameplay to your 4K TV. More HD Power: Turn on Boost Mode to give PS4 games access to the increased power of PS4 Pro. HDR Technology : With an HDR TV, compatible PS4 games display an unbelievably vibrant and life like range of colors. ",
            "price": "$339.00",
            "totalReviews": 8725
        },
        "last_update": "2020-11-29T21:00:12.958Z"
    }
]
```

## Using Postman :email:
As explained above, postman helps in making the calls, especially the POST method calls to hosts. Make the proxy disable as explained under `Prerequisites` section.
The steps are pretty simple.
1. Download Postman from official site [postman.com/downloads/](https://www.postman.com/downloads/).
2. Install and launch the application in local.
3. Select the method (GET/POST) from the drop-down and enter the above mentioned URLs (one at a time) in the corresponding space.
4. If POST requests, you'll find the `Body` option just below the URL. Click and go as follows:<br>
   **Body (Sub Menu) -> Raw (Radio Button) -> JSON (from Drop-down)** <br>
   This will enable the description form where you can copy-paste JSON data. This works more efficiently than selecting individual Key-Value pairs.
Use the URLs, methods, form data and check for outputs in this with that of mentioned under `Making the Calls` section above.


#### This concludes everything that is required to check and make use of the [Amazon-Scraper-Collector](https://github.com/VagueCoder/Amazon-Scraper-Collector). The code walk-throughs will be added in the future developments on this. For any issues, queries or discussions, please update in issues menu or write to `vaguecoder0to.n@gmail.com`.

## Happy Coding !! :metal:
