# Rede Social

Welcome to the Rede Social development repository!

* [Getting Started](#getting-started)
* [Credentials](#credentials)

## Getting Started

Social Network is a simple social network, which contains an `API` developed in GoLang and Mysql, responsible for communicating with the database (back-end), and a web application (`WebApp`) developed in GoLang, Javascript and Bootstrap that communicates with the API and renders the information to the user (front-end). In this way, for the web application to work correctly, the API must be running.

A local development environment is available to quickly get up and running. You will need a basic understanding of how to use the command line on your computer. This will allow you to set up the local development environment, to start it and stop it when necessary, and to run the tests.

You will need Go installed on your computer. Go was created at Google in 2007, and since then, engineering teams across Google have adopted Go to build products and services at massive scale. If you don't have Go installed, [click here](https://go.dev/dl/) to download and install.


### Folder structure
[api](https://github.com/douglasfelc/golang-mysql-redeSocial/tree/main/api): responsible for communicating with the database
[webapp](https://github.com/douglasfelc/golang-mysql-redeSocial/tree/main/webapp): web application that renders information to the user


### Development Environment Commands

Ensure [Docker](https://www.docker.com/products/docker-desktop) is running before using these commands.

#### To start the development environment for the first time

Clone the current repository using:

```
git clone https://github.com/douglasfelc/golang-mysql-redeSocial.git
```

Then in your terminal move to the repository folder `cd golang-mysql-redeSocial`.

After that, first access the API folder `cd api` and run the following commands:

#### To build and start using [Docker Compose](https://docs.docker.com/compose/reference/)

```
docker-compose up --build
```

Access the MySQL database through your manager or terminal, and execute the SQLs from the `/api/sql/sql.sql` file.

If you want to populate the table with some examples, also run the file `/api/sql/exampleData.sql`.

Now you will have the environment necessary to run the API, with Go, MySQL and PHPMyAdmin. At this point you can choose to run the API directly from the executable, running:

```
./api
```

Or if you have made any changes to the code, and want to compile, run:

```
go run api
```

And to generate a new executable:

```
go build
```

The Rede Social API is accessible at http://localhost:5000. You can see or change configurations in the `.env` file located at the root of the project directory.

Now all that's left to do is run the web application. Assuming you are in the api folder, you must first go back one level in the folder by typing `cd ..` in your terminal, and now access the web application folder by typing `cd webapp`. Likewise, you can directly run the executable by running:

```
./webapp
```

Or if you have made any changes to the code, and want to compile, run:

```
go run webapp
```

And to generate a new executable:

```
go build
```

The Rede Social WebApp is accessible at http://localhost:3000. You can see or change configurations in the `.env` file located at the root of the project directory.

## Credentials

These are the default environment credentials of API:

* API port: `5000`
* Secret key: `8mYRgJV6mpZ++vYVJqiilDHes2VUyY0SQp8gLDaN+JmTOccFOqJqENRrBQlqq0C911MGmRNj+5sleh+2+vVlyA==`
* ------------------
* MySQL Port: `3306`
* MySQL Root Password: `senharoot`
* Database Host: `db`
* Database Name: `redeSocial`
* Database Username: `golang`
* Database Password: `golang`

These are the default environment credentials of WebApp:

* API URL: `http://localhost:5000`
* APP Port: `3000`
* Hash Key: `643462c178346172b9497907dedcd1ea`
* Block Key: `f70717a9a0d384e6869d2c024597909e`

To create an account, navigate to http://localhost:3000/signup