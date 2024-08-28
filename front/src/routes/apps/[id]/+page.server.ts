import { API_URL } from "$lib/api/index.ts";
import { redirect } from '@sveltejs/kit';

export const actions = {
  startContainer: async ({ fetch, params }) => {
    const req = new Request(API_URL + `/apps/${params.id}/start`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
    });
    try {
      const res = await fetch(req);
    } catch (err) {
      console.error(err);
      throw err;
    }
  },
  stopContainer: async ({ fetch, params }) => {
    const req = new Request(API_URL + `/apps/${params.id}/stop`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
    });
    try {
      await fetch(req);
    } catch (err) {
      console.error(err);
      throw err;
    }
  },
  deploy: async ({ fetch, params }) => {
    const req = new Request(API_URL + `/apps/${params.id}/deploy`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
    });
    try {
      await fetch(req);
    } catch (err) {
      console.error(err);
      throw err;
    }
  },
  deleteApp: async ({ fetch, params }) => {
    const req = new Request(API_URL + `/apps/${params.id}`, {
      method: "DELETE",
      headers: {
        "Content-Type": "application/json",
      },
    });
    try {
      const res = await fetch(req);
      if (res.ok) {
        throw redirect(303, '/apps');
      }
    } catch (err) {
      console.error(err);
      throw err;
    }
  }
};