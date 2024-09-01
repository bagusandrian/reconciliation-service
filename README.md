            # Reconciliation-Service

## About Reconciliation-Service
This service simplifies the reconciliation process between your internal system transactions and bank statements.   

It helps you identify discrepancies, missing transactions, and errors, saving you time and ensuring financial accuracy. 

## Processing flow

### Sequence Diagram

![image](https://github.com/user-attachments/assets/31ba852f-c603-4eb3-ba4b-2d4a64673c2d)

### Flow chart

![image](https://github.com/user-attachments/assets/6e7e498d-9cc6-45ac-b1a9-6ac245767a9e)


## Tech Stack

**Web framework:** gofiber [![Gofiber Doc](https://camo.githubusercontent.com/5140b60155572122a7d51faf0ddfdcab144059c3b3f2293c548fe94dc3842b53/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f254630253946253933253941253230676f646f632d706b672d3030414344372e7376673f636f6c6f723d303041434437267374796c653d666c61742d737175617265)](https://pkg.go.dev/github.com/gofiber/fiber/v3#pkg-overview)



## Pre-requisite

1. Golang version min 1.20
## Installation

1. Clone this repo into your local laptop
```bash
    git clone https://github.com/bagusandrian/reconciliation-service.git
```
2. If u want to run in docker, just running this command on your terminal (on path repo that u already clone before): 
```bash
make run-docker
```
3. If you want to run locally, u need download dependecy library first:
```bash
go mod tidy
```
4. After done, u can run this command on your terminal: 
```bash
make run-http
```
This log will showing on your terminal, indicate service is running. 
```bash
$ make run
 > Building [app]...
 > Finished building [app]
 > Running [app]...
2024/09/01 14:43:59 Starting reconciliation-service http
2024/09/01 14:43:59 init config done
2024/09/01 14:43:59 starting http server 2024-09-01 14:43:59.741702 +0700 WIB m=+0.002544835

 ┌───────────────────────────────────────────────────┐ 
 │              reconciliation Service               │ 
 │                   Fiber v2.52.5                   │ 
 │               http://127.0.0.1:9000               │ 
 │       (bound on host 0.0.0.0 and port 9000)       │ 
 │                                                   │ 
 │ Handlers ............. 1  Processes ........... 1 │ 
 │ Prefork ....... Disabled  PID ............. 47863 │ 
 └───────────────────────────────────────────────────┘ 
```
This service running on port `:9000`, u can change port if u want on `files/etc/reconciliation-service/reconciliation-service.development.yaml`

### Reference of Makefile commands

| Command |  Action |
|:-----|:--------:|
| `make build` | will build binary |
| `make run` | this will build the binary first, after that run the binary |
| `make run-docker` | this will run service using docker | 
| `make test-coverage` | will run go test with total of coverage unit test |
| `make test-fail` | will run go test and only showing log if any error |


## API Reference
This service only has 1 method
### Reconcilition data

```http
  POST /reconciliation
```

#### Request
request parameter put in body as json

| Param Request | Type     | **Required** | Description |
| :-------- | :------- | :----------- | :----------- |
| `system_transaction_csv_file_path` | `string` | **Required** | File path of system transaction file. File must be csv |
| `bank_statements` | []bank_statement | **Required** | list of bank_statement |
| bank_statement -> `bank_name` | `string` | **Required** | name of bank |
| bank_statement -> `csv_file_path` | `string` | **Required** | File path of bank statements file. File must be csv |
| `reconciliaton_start_date` | `date` | **Required** | start date of reconciliation | 
| `reconciliaton_end_date` | `date` | **Required** | start date of reconciliation | 


example curl: 
```curl
curl --location 'localhost:9000/reconciliation' \
--header 'Content-Type: application/json' \
--data '{
    "system_transaction_csv_file_path": "/Users/bagusandrian/go/src/github.com/bagusandrian/reconciliation-service/files/csv/system.csv",
    "bank_statements": [
        {
            "bank_name": "bca",
            "csv_file_path": "/Users/bagusandrian/go/src/github.com/bagusandrian/reconciliation-service/files/csv/bca.csv"
        },
        {
            "bank_name": "danamon",
            "csv_file_path": "/Users/bagusandrian/go/src/github.com/bagusandrian/reconciliation-service/files/csv/danamon.csv"
        },
        {
            "bank_name": "bri",
            "csv_file_path": "/Users/bagusandrian/go/src/github.com/bagusandrian/reconciliation-service/files/csv/bri.csv"
        }
    ],
    "reconciliaton_start_date": "2024-01-20",
    "reconciliaton_end_date": "2024-01-22"
}'
```

#### Response

| Param Response | Type | Description |
| :------------- | :--- | :---------- |
| `total_transactions_processed` | int64 | total of transactions processed |
| `total_number_matched_transactions` | int64 | total number of match transactions |
| `detail_of_matched_transactions` | map[string]DetailMatchedTransaction |  detail information of match transactions |
| `total_number_unmatched_transactions` | int64 |  total number of unmatch transaction |
| `detail_of_unmatched_transactions` | map[string]DetailUnmatchedTransaction |  detail information of unmatch transactions |
| `total_discrepancies_amount` | float64 | total discrepancies amount match transaction |

##### DetailMatchedTransaction

| Param Response | Type | Description |
| :------------- | :--- | :---------- |
| total_number_matched_transactions | int64 | total number of match transactions |
| detail_transaction | []DetailTransaction |  detail of match transaction |

##### DetailUnmatchedTransaction

| Param Response | Type | Description |
| :------------- | :--- | :---------- |
| info | string | information of unmatch transaction |
| detail_transaction | []DetailTransaction |  detail of unmatch transaction |

##### DetailTransaction

| Param Response | Type | Description |
| :------------- | :--- | :---------- |
| trx_id | string | trx_id from system transaction |
| unique_identifier | string | unique_identifier from bank statement |
| amount | float64 | amount of transaction (positive == `debit`  negative == `credit`) |
| date | string | date of transaction |
| type | TypeTransaction | type of transaction (enum DEBIT and CREDIT) |
| transaction_time | string | transaction time from system transaction |

##### example Response
``` json
{
    "data": {
        "total_transactions_processed": 21,
        "total_number_matched_transactions": 18,
        "detail_of_matched_transactions": {
            "bca": {
                "total_number_matched_transactions": 4,
                "detail_transaction": [
                    {
                        "trx_id": "trx001",
                        "unique_identifier": "BCA-4",
                        "amount": -10000,
                        "date": "2024-01-20",
                        "transaction_time": "2024-01-20 10:00:00"
                    },
                    {
                        "trx_id": "trx002",
                        "unique_identifier": "BCA-1",
                        "amount": 10000,
                        "date": "2024-01-20",
                        "transaction_time": "2024-01-20 10:00:00"
                    },
                    {
                        "trx_id": "trx004",
                        "unique_identifier": "BCA-2",
                        "amount": 10000,
                        "date": "2024-01-20",
                        "transaction_time": "2024-01-20 10:00:00"
                    },
                    {
                        "trx_id": "trx006",
                        "unique_identifier": "BCA-3",
                        "amount": 10000,
                        "date": "2024-01-20",
                        "transaction_time": "2024-01-20 10:00:00"
                    }
                ]
            },
            "bri": {
                "total_number_matched_transactions": 2,
                "detail_transaction": [
                    {
                        "trx_id": "trx007",
                        "unique_identifier": "BRI-001",
                        "amount": -10000,
                        "date": "2024-01-20",
                        "transaction_time": "2024-01-20 10:00:00"
                    },
                    {
                        "trx_id": "trx008",
                        "unique_identifier": "BRI-002",
                        "amount": 10000,
                        "date": "2024-01-20",
                        "transaction_time": "2024-01-20 10:00:00"
                    }
                ]
            },
            "danamon": {
                "total_number_matched_transactions": 3,
                "detail_transaction": [
                    {
                        "trx_id": "trx003",
                        "unique_identifier": "DANAMON-001",
                        "amount": -10000,
                        "date": "2024-01-20",
                        "transaction_time": "2024-01-20 10:00:00"
                    },
                    {
                        "trx_id": "trx005",
                        "unique_identifier": "DANAMON-003",
                        "amount": -10000,
                        "date": "2024-01-20",
                        "transaction_time": "2024-01-20 10:00:00"
                    },
                    {
                        "trx_id": "trx010",
                        "unique_identifier": "DANAMON-002",
                        "amount": 10000,
                        "date": "2024-01-20",
                        "transaction_time": "2024-01-20 10:00:00"
                    }
                ]
            }
        },
        "total_number_unmatched_transactions": 3,
        "detail_of_unmatched_transactions": {
            "danamon": {
                "info": "bank statement not found on any system transaction",
                "detail_transaction": [
                    {
                        "unique_identifier": "DANAMON-004",
                        "amount": -90000,
                        "date": "2024-01-20",
                        "type": 2
                    }
                ]
            },
            "system": {
                "info": "system transaction not found on any bank statement",
                "detail_transaction": [
                    {
                        "trx_id": "trx009",
                        "amount": -10000,
                        "type": 2,
                        "transaction_time": "2024-01-20 10:00:00"
                    },
                    {
                        "trx_id": "trx011",
                        "amount": 10000,
                        "type": 1,
                        "transaction_time": "2024-01-20 10:00:00"
                    }
                ]
            }
        },
        "total_discrepancies_amount": 0
    },
    "header": {
        "status": 200,
        "processing_time": "1.578917ms"
    }
}
```

# Next Steps
## Enhancements:
1. Performance optimization: Optimize the reconciliation process for large datasets to improve efficiency.
2. Additional features: Consider adding features like:
   - Customizable reconciliation rules: Allow users to define their own reconciliation rules.
   - Integration with other systems: Integrate with accounting software or other financial systems for seamless data flow.
   - Scheduled reconciliations: Set up automatic reconciliation schedules.
3. Expand test coverage: Write more comprehensive unit and integration tests to ensure code quality.
