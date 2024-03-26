import http from 'k6/http';
import { sleep } from 'k6';

export const options = {
  vus: 10000,
  duration: '10s',
  thresholds: {
    http_reqs: ['rate>=10000'],
  },
};

export default function () {
  for (let i = 0; i < 1; i++) {
    http.get(
      'http://localhost:8080/api/v1/ad?offset=5&age=30&gender=F&country=TW&platform=IOS',
    );
  }

  sleep(1);
}
