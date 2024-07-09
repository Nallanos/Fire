import { API_URL } from "../index";

export async function listApps() {
  const req = new Request(API_URL + "/apps", {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
    },
  });

  const res = await fetch(req);
  const data = await res.json();
  console.log(data);
  return data;
}
