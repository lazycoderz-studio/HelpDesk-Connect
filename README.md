# HelpDesk-Connect
## Overview
HelpDesk Connect is a real-time chat application designed to connect customers with agents for assistance. The system includes a queue management mechanism to handle multiple customer requests and ensure efficient agent-customer interactions.

## Backend Design (Golang)

####  WebSocket Server: Manages real-time communication between customers and agents.
####  Queue Management: Handles customer queueing and agent availability.
####  REST API: Provides endpoints for agent actions, such as picking a customer from the queue.
####  Database Interaction: Stores and retrieves data related to agents, customers, and chat sessions.


## Database (PostgreSQL):

####  Agents Table: Stores information about agents.
####  Customers Table: Stores session details for customers.
####  ChatSessions Table: Stores details of chat sessions, including timestamps and statuses. 

## Architectural Design
```
 +-----------+           +---------------------+            +------------+
 |           |           |                     |            |            |
 | Customer  +<--------->+  WebSocket Server   +<---------->+ PostgreSQL |
 | Interface |           |                     |            |  Database  |
 +-----------+           +---------------------+            +------------+
                             ^            ^  |                  ^
                             |            |  |                  |
                             v            v  v                  |
                         +-----------+   +------------+     +----------+
                         | Agent     |   | REST API   |     | Queue    |
                         | Interface |   |            |     | Management|
                         +-----------+   +------------+     +----------+

```
