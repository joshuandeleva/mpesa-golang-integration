# M-Pesa STK Push and Callback Integration

This project is a Go-based implementation for integrating M-Pesa's STK Push and Callback functionalities. It includes components for initiating STK Push requests, handling M-Pesa callbacks, and saving callback data to a PostgreSQL database.

## Features

- **STK Push**: Initiate a payment request using M-Pesa's STK Push API.
- **Callback Handling**: Handle callback responses from M-Pesa and save them to the database.
- **Asynchronous Processing**: Save callback data in the background without blocking other processes.

## Prerequisites

- **Go**: Ensure you have Go installed. You can download it from [golang.org](https://golang.org).
- **PostgreSQL**: You need a PostgreSQL database to store callback responses.
- **Ngrok**: For local development and testing with M-Pesa, you may use Ngrok to tunnel requests.

## Environment Setup

Create a `.env` file at the root of the project with the following contents:
```
CONSUMER_KEY=your_consumer_key
CONSUMER_SECRET=your_consumer_secret
SHORT_CODE=your_short_code
PASS_KEY=your_pass_key
DB_USER=your_db_user
DB_PASSWORD=your_db_password
DB_NAME=your_db_name
DB_HOST=localhost
DB_PORT=5432
APP_PORT=6441
```
