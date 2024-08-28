import type { PageServerLoad } from "./$types";
import { API_URL } from "lib/api";
import { type App } from "lib/api/apps/appTypes";
export const load: PageServerLoad = async ({ fetch }) => {
  const req = new Request(API_URL + "/apps", {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
    },
  });
  const res = await fetch(req);
  const apps: App[] = await res.json()
  if (apps) {
    return { apps }
  }
};