// To run just type: k6 run --vus 100 --iterations 10000  integration-tests/main.js
import http from 'k6/http';

export default function () {
    http.get('http://localhost:8080/countfs');
}
