# HelpDesk-Connect
## Overview
HelpDesk Connect is a real-time chat application designed to connect customers with agents for assistance. The system includes a queue management mechanism to handle multiple customer requests and ensure efficient agent-customer interactions.

## Backend Design (Golang)

####  WebSocket Server: Manages real-time communication between customers and agents.
####  Queue Management: Handles customer queueing and agent availability.
####  REST API: Provides endpoints for agent actions, such as picking a customer from the queue.
####  Database Interaction: Stores and retrieves data related to agents, customers, and chat sessions.


## Database (PostgreSQL):

####  Users Table: Stores information about agents.
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
## Detailed Workflow
### Customer Workflow:

#### Initiate Chat:

Customer opens the chat page, inputs their name and email, and clicks "Start Chat."
Frontend establishes a WebSocket connection with the server using a unique customer_id.
Join Queue:

The backend server generates a unique customer_id and places the customer in the waiting queue.
The customer receives a message indicating they are in the queue.
#### Agent Workflow:

Login:

Agent logs into their dashboard and establishes a WebSocket connection using their agent_id.
View Queue:

Agent views the queue of waiting customers via their dashboard.
Pick Customer:

Agent picks a customer from the queue.
Backend server notifies the customer and opens the chat interface for both the agent and the customer.
####  Chat Session:

Real-Time Communication:

Messages are exchanged in real-time via WebSocket.
End Session:

When the chat is complete, the agent marks the session as complete.
Backend server updates the database, marking the session as complete and removing the customer from the queue.
