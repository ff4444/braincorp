# Robot API Documentation

Welcome to the documentation for My Awesome API. This API allows you to interact with various resources and perform CRUD operations.

## Authentication

All endpoints require authentication using API keys. You can obtain your API key by signing up on our customer solution team.

## Base URL

The base URL for all API endpoints is `https://api.braincorp.com`.

## Endpoints

### 1. Create New Event List

Post a list of events that have occured recently. Events are time stamped so that chronology is preserved even when internet connectivity may be compromised.

#### Request

- **URL:** `/events`
- **Method:** `POST`
- **Authentication:** Required (API key)

```json
{
   "username" : "my_username",
   "robotID" : "my_id",
   "sw version" : "2.03",
   "fw version" : "2.02",
   "hw version" : "1.1",
   "event name" : "event1029: Left drive moter over-current warning",
   "epoch time" : 1715286332,
   "validation-factors" : {
      "validationFactors" : [
         {
            "name" : "remote_address",
            "value" : "127.0.0.1",
            "timekey" : "my_key"
         }
      ]
   }
   "eventsArray" : [
       "event" : {
           "event name" : "event1029: Left drive moter over-current warning",
           "epoch time" : 1715286332
           }
   ]
}
```
#### Response

- **Status Code:** 200 OK
- **Body:**

```json
[
    {
        "name": "EventsReceived",
        "count": "1"
    }
]
```
