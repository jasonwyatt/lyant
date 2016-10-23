# Authentication

This file will discuss the architecture and design of signup and login for lyant.

Some key points:

* Authentication/Sign-up will be done via "Log in with github"
* Users will be stored in neo4j as User nodes.
* auth tokens and refresh tokens will be stored in Auth nodes, linked to Users
* both auth tokens and refresh tokens will need to be encrypted
* ... other key points?

## User Nodes

## Auth Nodes

### Encrypting the auth/refresh tokens

## User -> Auth Relationships
