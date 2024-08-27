# Loan Tracker Application

## Overview

The Loan Tracker Application is designed to manage loan applications. It supports user registration, email verification, loan application, loan status retrieval, and admin-specific operations like viewing, approving, and rejecting loan applications.

## Features

### User Management
- **User Registration**: Users can register by providing their details. An email verification token is sent to the user's email.
- **Email Verification**: Users verify their email using a token sent to their email address.
- **Login/Logout**: Users can log in and out, obtaining a JWT token for accessing protected resources.

### Loan Management
- **Apply for a Loan**: Users can apply for loans by providing necessary details.
- **View Loan Status**: Users can view the status of their loan applications.
- **Admin View All Loans**: Admins can view all loan applications, with options to filter by status and order.
- **Approve/Reject Loan**: Admins can approve or reject loan applications.

## Project Structure

- `domain/`: Contains business logic interfaces and domain models.
- `repository/`: Contains implementations for interacting with the MongoDB database.
- `usecase/`: Contains use case logic that interacts with the domain and repositories.
- `delivery/`: Contains controllers for handling HTTP requests.
- `middleware/`: Contains middleware for handling JWT authentication.
- `infrastructure/`: Contains configurations for authentication, database, and email services.

## Setup Instructions

### Prerequisites

- Go 1.18 or higher installed
- MongoDB instance running
- Access to an SMTP service like Mailtrap for sending emails

### Environment Variables

Create a `.env` file in the root directory with the following environment variables:

```dotenv
# MongoDB
DB_NAME=loan_tracker
DB_URI=mongodb://localhost:27017

# Server
SERVER_ADDRESS=:8080

# JWT Secrets
ACCESS_TOKEN_SECRET=youraccesstokensecret
REFRESH_TOKEN_SECRET=yourrefreshtokensecret
RESET_TOKEN_SECRET=yourresettokensecret

# SMTP (Mailtrap)
SMTP_SERVER=smtp.mailtrap.io
SMTP_PORT=587
SMTP_USER=yourmailtrapusername
SMTP_PASSWORD=yourmailtrappassword
FROM_ADDRESS=youremail@example.com

# Context Timeout
CONTEXT_TIMEOUT=2m
