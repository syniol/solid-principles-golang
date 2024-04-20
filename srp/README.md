# Single Responsibility Principle
In this example we are creating a backend example service for a companies house. To demonstrate 
the difference in code with & without utilising of Single Responsibility Principle.


## Overview
As part of public API offered to our client, we require to create a new endpoint to 
display a list of companies matching any registered company in our database. New endpoint 
is divided into two tasks:
 * New RESTful API endpoint `/v1/company` with `GET` method
 * __Service to be injected inside the endpoint to service the information__

In this User story we will deliver the second bullet-point to create a service that 
will later be injected inside the new endpoint

## Requirements
 * Create a new service/method to find a company by name
 * Only accept company name and it should not be empty
 * Connect to database to find and list companies matching the name


## Acceptance Criteria
 * New method similar to `FindCompany` that accept company name as a parameter
 * Method should list companies with similar name in an array of strings
 * Unit Test


##### Credits
Copyright &copy; 2024, Syniol Limited. All rights reserved.

Author: [Hadi Tajallaei](mailto:hadi@syniol.com)
