# Ina-Rec

This repository contains a Go application called Ina-Rec. Ina-Rec is a RESTful API for a task application that allows registered users to perform CRUD operations on task posts.

## Prerequisites

Before you can run this application locally, make sure you have the following prerequisites installed on your system:

- Go (version 1.18): [Installation Guide](https://golang.org/doc/install)
- PostgreSQL (version 14): [Installation Guide](https://www.postgresql.org/download/)

## Getting Started

Follow these steps to run the Ina-Rec application locally:

1. Clone this repository to your local machine:

   ```bash
   git clone https://github.com/priyanfadhil/ina-rec.git
2. Change your current directory to the project folder:

   ```bash
   cd ina-rec
3. Build the Go application:

   ```bash
   go build
3. Run the application:

   ```bash
   go run main.go
The application should now be running locally on your machine.

Create the PostgreSQL database table for user and sequence table. You can use the following SQL command as an example:

    CREATE TABLE public.users
    (
        id integer NOT NULL DEFAULT nextval('users_id_seq'::regclass),
        name character varying,
        email character varying,
        password character varying,
        created_at character varying,
        created_by character varying,
        updated_at character varying,
        updated_by character varying,
        PRIMARY KEY (id)
    );
    
    CREATE SEQUENCE public.users_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 2147483647
    CACHE 1
    OWNED BY users.id;

Create the PostgreSQL database table for tasks and sequence table. You can use the following SQL command as an example:

    CREATE TABLE public.tasks
    (
        id integer NOT NULL DEFAULT nextval('tasks_id_seq'::regclass),
        user_id integer,
        title character varying,
        description text,
        status character varying,
        created_at character varying,
        created_by character varying,
        updated_at character varying,
        updated_by character varying,
        PRIMARY KEY (id)
    );
    
    CREATE SEQUENCE public.tasks_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 2147483647
    CACHE 1
    OWNED BY tasks.id;

## Configuration
You may need to configure the application by create .env file by duplicating the contents of the .envexample file and filling in its variables.

## Usage
To use the Ina-Rec application, you must first register as a user:
1. Make a POST request to the `{localhost}/api/v1/user/register` endpoint with the following JSON payload:

   ```bash
    {
        "name": "your_name",
        "password": "your_password",
        "email": "your_email",
    }
2. If the registration is successful, you will receive a confirmation message.
3. After registering, you can log in to obtain a JWT token for authentication. Make a POST request to the `{localhost}/api/v1/user/login` endpoint with the following JSON payload
    ```bash
    {
        "name": "your_name",
        "password": "your_password"
    }
4. If the login is successful, you will receive a JWT token in the response.

You can use the JWT token to access other protected endpoints by including it in the Bearer Token Authorization of your requests.