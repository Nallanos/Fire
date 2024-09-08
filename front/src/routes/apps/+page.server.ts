import type { PageServerLoad } from "./$types";
import { API_URL } from "lib/api";
import { error } from '@sveltejs/kit';
import type { Actions } from "@sveltejs/kit";
import { redirect } from "@sveltejs/kit";
import { type App } from "lib/api/apps/appTypes";
type bodyHttpError = {
  message: string
}

export type httpError = {
  status: string,
  body: bodyHttpError
}


export const load: PageServerLoad = async ({ fetch }) => {
  try {
    const req = new Request(API_URL + "/apps", {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
      },
    });

    const res = await fetch(req);

    if (!res.ok) {
      throw error(res.status, `Failed to load apps: ${res.statusText}`);
    }

    const apps: App[] = await res.json();

    return { apps };

  } catch (err: any) {
    console.error("Error loading apps:", err);
    throw error(500, err);
  }
};


export const actions: Actions = {
  deleteCookies: async ({ cookies }) => {
    console.log("Deleting cookies...");
    cookies.delete("token", { path: '/' });

    return redirect(303, "/");
  }
};
