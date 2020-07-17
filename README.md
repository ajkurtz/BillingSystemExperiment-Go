## Billing System Experiment

This is a coding exercise to experiment with APIs, a database, and AWS SQS.  It is implemented in Go.

### Component Descriptions

The system consists of two APIs and two CLI applications.

#### CustomersAPI

An API that performs CRUD operations on the customers resource.

#### PaymentsAPI

An API that performs CRUD operations on the payments resource.

#### BillScheduling Service

A CLI application that searches the database for customers who should be billed "today" and adds them to a queue for billing.

#### Billing Service

A CLI application that reads customers from a queue and creates a payment for them based on the data in their customer record.


