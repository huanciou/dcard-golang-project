import http from 'k6/http';
import { sleep } from 'k6';

export const options = {
  vus: 5000,
  duration: '30s',
  thresholds: {
    http_reqs: ['rate>=10000'],
  },
};

export default function () {
  for (let i = 0; i < 10; i++) {
    http.get(
      'http://localhost:8080/api/v1/ad?offset=1&age=30&gender=f&country=tw&platform=ios',
    );
  }

  sleep(1);
}
