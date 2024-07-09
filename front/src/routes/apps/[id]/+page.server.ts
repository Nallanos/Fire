import { API_URL } from "$lib/api/index.ts";
export const actions = {
  startContainer: async (e) => {
    const req = new Request(API_URL + `/apps/${e.params.id}/start`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
    });
    try {
      const res = await fetch(req);
      console.log("application started", res.status);
    } catch (err) {
      console.error(err);
      throw err;
    }
  },
  stopContainer: async (e) => {
    const req = new Request(API_URL + `/apps/${e.params.id}/stop`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
    });
    try {
      const res = await fetch(req);
      console.log("application stopped", res.status);
    } catch (err) {
      console.error(err);
      throw err;
    }
  },
  deploy: async (e) => {
    console.log("deploying application");
    const req = new Request(API_URL + `/apps/${e.params.id}/deploy`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
    });
    try {
      const res = await fetch(req);
      console.log("application deployed", res.status);
      if (!res.ok) {
        throw new Error(`HTTP error! status: ${res.status}`);
      }
    } catch (err) {
      console.error(err);
      throw err;
    }
  },
};
