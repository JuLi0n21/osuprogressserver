import http from 'k6/http';
import { check, sleep } from 'k6';
import { Rate } from 'k6/metrics';

// Define custom metrics
let errorRate = new Rate('errors');

export let options = {
    stages: [
        { duration: '30s', target: 50 },  // Ramp-up to 50 users over 30 seconds
        { duration: '30s', target: 200 },  // Ramp-up to 50 users over 30 seconds
        { duration: '1m', target: 50 },   // Stay at 50 users for 1 minute
        { duration: '30s', target: 0 },   // Ramp-down to 0 users over 30 seconds
    ],
    thresholds: {
        'http_req_duration': ['p(95)<500'], // 95% of requests should be below 500ms
        'errors': ['rate<0.01'],            // Error rate should be less than 1%
    },
};

const BASE_URL = 'https://progress.illegalesachen.download';

const endpoints = [
    { url: `${BASE_URL}/api/scoresearch?query=apo&from=&to=09:33 20/05/2024&limit=10&offset=0`, weight: 50 },
    { url: `${BASE_URL}/me`, weight: 20 },
    { url: `${BASE_URL}/score/2000`, weight: 20 },
    { url: `${BASE_URL}/`, weight: 10 },
];

function getWeightedRandomEndpoint() {
    let totalWeight = 0;
    for (let endpoint of endpoints) {
        totalWeight += endpoint.weight;
    }

    let randomWeight = Math.random() * totalWeight;
    for (let endpoint of endpoints) {
        if (randomWeight < endpoint.weight) {
            return endpoint.url;
        }
        randomWeight -= endpoint.weight;
    }
}

export default function () {
    let url = getWeightedRandomEndpoint();
    let res = http.get(url);

    // Check if the response status is 200
    let checkRes = check(res, {
        'is status 200': (r) => r.status === 200,
        'is duration < 500ms': (r) => r.timings.duration < 500,
    });

    // Record errors if any check failed
    if (!checkRes) {
        errorRate.add(1);
    }

    // Sleep for 1 second between iterations
    sleep(1);
}

