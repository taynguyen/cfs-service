Simple Call For Service (CFS) service


## Start dev
MySQL installed

Edit file dev/reset-db.sh with your DB account and password
```
make reset-db-dev
make dev
```


Example APIs:
### Health check
```
curl localhost:8080/healthz
```

### POST events
```
curl --request POST \
  --url http://localhost:8080/api/v1/events \
  --header 'Content-Type: application/json' \
  --cookie auth-token=eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJhX2lkIjoiNGY5Yjk5ZWItNDkwYS00ODRlLWJhZGUtMTVlMzg0MWRmZGE5IiwiYyI6MTYxODE0OTcxNjIwMiwiZXhwIjoxNjIwNzQxNzE2MjAyfQ.bry5uYJAWn1WqOhdD4KY4DRZMPTR-6q3uUlfsQnWrkmE_CTWHIM3aFHQV49ilpBKGG0ft1otF4Z-5xzIvhM52nAACd_A7sEXRy8sgq5mOEYmjBO9xW2ygL83Wd7hTJA7mN6vy8RCeb3lGLMLoMmuMFFNbeuNKhmZpqIsD0Myj93S7CUsRHVL3gysbBpSY4iySKPI4SSK5d8fMd6DRQwiZxjHHHh-VYP6c5HwbuJnGX-9YJjOOc-9SiJvAHa8FIJJxochGvOtxV3i_qlC0MVDM7AfkKb0WSQSJvmoTImsguw6_fycvOOOPeVzI6VSZK6DXXU2oZvAt2HIZkaTgeegled9u3F9Zu0vjT5sZ_SdOc8MZKO9pRa2wD2bQUKPTQverkE5NYnbiHOlfl_n22EdFbAfmjyLqdrJauAk4A2YyDhcXOUrvZEsqafH2D76ssDonjbYuCG9oqyek_0CTdw6r4ODHn_zXBoSWsgd3xG-IxPwk1bHtrtS8InECEi66H3b \
  --data '{
	"agency_id": "4f9b99eb-490a-484e-bade-15e3841dfda9",
	"event_id": "562c89de-f140-4482-8ef5-5f1703b286b6",
	"event_number": "3234019",
	"event_type_code": "SMO",
	"event_time": "2020-11-25 07:36:04.193",
	"dispatch_time": "2020-11-26 13:55:46.466",
	"responder": "OFFICER_001"
}'
```

### Search request
```
curl --request GET \
  --url 'http://localhost:8080/api/v1/events?from=2020-11-25%2007%3A36%3A04.000&to=2020-11-25%2007%3A36%3A05.000&offset=0' \
  --cookie auth-token=eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJhX2lkIjoiNGY5Yjk5ZWItNDkwYS00ODRlLWJhZGUtMTVlMzg0MWRmZGE5IiwiYyI6MTYxODE0OTcxNjIwMiwiZXhwIjoxNjIwNzQxNzE2MjAyfQ.bry5uYJAWn1WqOhdD4KY4DRZMPTR-6q3uUlfsQnWrkmE_CTWHIM3aFHQV49ilpBKGG0ft1otF4Z-5xzIvhM52nAACd_A7sEXRy8sgq5mOEYmjBO9xW2ygL83Wd7hTJA7mN6vy8RCeb3lGLMLoMmuMFFNbeuNKhmZpqIsD0Myj93S7CUsRHVL3gysbBpSY4iySKPI4SSK5d8fMd6DRQwiZxjHHHh-VYP6c5HwbuJnGX-9YJjOOc-9SiJvAHa8FIJJxochGvOtxV3i_qlC0MVDM7AfkKb0WSQSJvmoTImsguw6_fycvOOOPeVzI6VSZK6DXXU2oZvAt2HIZkaTgeegled9u3F9Zu0vjT5sZ_SdOc8MZKO9pRa2wD2bQUKPTQverkE5NYnbiHOlfl_n22EdFbAfmjyLqdrJauAk4A2YyDhcXOUrvZEsqafH2D76ssDonjbYuCG9oqyek_0CTdw6r4ODHn_zXBoSWsgd3xG-IxPwk1bHtrtS8InECEi66H3b
```