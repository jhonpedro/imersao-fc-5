Important commands just for saving

Before start, run the `helper/createDb.go` helper for creating the base data base

`kafka-console-producer --bootstrap-server=localhost:9092 --topic=transactions`

`{"id":"1234","account_id":"1","credit_card_number":"123123123213123","credit_card_name":"Sample","credit_card_expiration_month":12,"credit_card_expiration_year":2024,"credit_card_expiration_cvv":122,"amount":1200}`

`kafka-console-consumer --bootstrap-server=localhost:9092 --topic=transactions_result`
