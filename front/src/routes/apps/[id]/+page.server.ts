import { getApp } from "$lib/api/apps/getApp.js";
import { listDeployments } from "$lib/api/deployment/listDeployments.js";
import { getActiveDeployment } from "$lib/api/deployment/getActiveDeployment.js";
import { API_URL } from "$lib/api/index.ts";
import type { PageServerLoad } from "./$types.js";
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
    } catch (err) {
      console.error(err);
      throw err;
    }
  },
  stopContainer: async (e) => {
    console.log(`stopping deployment of ${e.params.id} at ${API_URL}/apps/${e.params.id}/stop`)
    const req = new Request(API_URL + `/apps/${e.params.id}/stop`, {
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
  deploy: async (e) => {
    const req = new Request(API_URL + `/apps/${e.params.id}/deploy`, {
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

};
export const load: PageServerLoad = async ({ params }) => {
  const appResult = await getApp(params.id);

  const deploymentResult = await listDeployments(params.id);


  const activeDeploymentResult = await getActiveDeployment(params.id);

  return ({ appResult, deploymentResult, activeDeploymentResult })
}
