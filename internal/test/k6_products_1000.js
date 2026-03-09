import http from "k6/http";
import { check, sleep } from "k6";

export let options = {
  stages: [
    { duration: "30s", target: 1000 }, // ramp up to 10 users
    { duration: "1m", target: 1000 }, // stay at 10 users
    { duration: "30s", target: 0 }, // ramp down
  ],
};

export default function () {
  let res = http.get("http://localhost:8080/products");
  check(res, {
    "status is 200": (r) => r.status === 200,
  });
  sleep(1);
}
