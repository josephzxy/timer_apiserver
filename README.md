# timer_apiserver
`timer_apiserver` is a demo Golang project that implements both external facing RESTful APIs and internal facing gRPC APIs for managing resource `timer`.
Resource `timer` will stored in `MySQL`

**Major Tech stacks:**
Golang, Gin, gRPC, MySQL, GORM

## Initial Design

### Model
`timer` has following structure
```go
// TimerCore contains fields that can be specified directly via APIs
type TimerCore struct {
    Name string `json:"name" gorm:"unique"`
    TriggerAt time.Time `json:"triggerAt" gorm:"index"`
}

// Model is the slightly-updated version of gorm.Model
// It will be managed automatically by gorm
type Model struct {
    ID uint `json:"-" gorm:"primarykey"`
    CreatedAt time.Time `json:"createdAt"`
    UpdatedAt time.Time `json:"-"`
    DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type Timer struct {
    Model
    TimerCore
}
```

The corresponding table in MySQL is like below:
```sql
CREATE TABLE IF NOT EXISTS `timer` (
    `id` BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT,
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `deleted_at` DATETIME DEFAULT NULL,
    `name` VARCHAR(255) NOT NULL,
    `trigger_at` DATETIME NOT NULL,
    PRIMARY KEY (`id`),
    KEY `idx_timer_deleted_at` (`deleted_at`),
    KEY `idx_timer_trigger_at` (`trigger_at`),
    CONSTRAINT `uniq_timer_name` UNIQUE (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4;
```

### RESTful API
Basic CRUD operations are supported.
#### Create
`POST /v1/timers`

Example request body
```json
{
    "name": "test_timer",
    "triggerAt": "2021-09-16T00:00:00+08:00"
}
```
(The timestamp is in RFC3339 format)

Example response
```json
// success
// HTTP status: 200
{
    "data": {
        "name": "test_timer",
        "triggerAt": "2021-09-16T00:00:00+08:00",
        "createdAt": "2021-09-15T00:00:00+08:00"
    },
    "err": null,
}
```
```json
// failure
// HTTP status: 400
{
    "data": null,
    "err": {
        "code": 100001,
        "msg": "Request validation failed"
    }
}
```

#### Read
`GET /v1/timers/{name}`

Example response is similar to the one in [RESTful API-Create](#Create)

#### Read all
`GET /v1/timers`

Example response
```json
// success
// HTTP status: 200
{
    "data": [
        {
            "name": "test_timer",
            "triggerAt": "2021-09-16T00:00:00+08:00", 
            "createdAt": "2021-09-15T00:00:00+08:00"
        },
        ...
    ],
    "err": null,
}
```
Error response is similar to the one in [RESTful API-Create](#Create)

#### Update
`PUT /v1/timers/{name}`

Example request body and response is similar to the one in [RESTful API-Create](#Create)

#### Delete
`DELETE /v1/timers/{name}`

Example response is similar to the one in [RESTful API-Create](#Create)

### gRPC API
The service definition file is like below.

```proto3
syntax = "proto3";

package proto;

service Resource {
    rpc ListPendingTimers(ListPendingTimersReq) returns (ListPendingTimersResp) {}
}

message TimerInfo {
    string name = 1;
    string trigger_at = 2;
}

message ListPendingTimersReq {
}

message ListPendingTimersResp {
    repeated TimerInfo items = 1;
}
```
A timer will be considered as "pending" if it is created, not deleted, and not triggerred yet.
