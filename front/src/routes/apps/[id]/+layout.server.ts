import type { LayoutServerLoad } from './$types';
import { API_URL } from '$lib/api';
import { errorsMessage } from '$lib/api/error';

async function getActiveDeployment(id: string, fetch: (input: RequestInfo | URL, init?: RequestInit) => Promise<Response>) {
  const req = new Request(`${API_URL}/apps/${id}/deployment/activeDeployment`, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
    },
  });

  try {
    const res = await fetch(req);
    if (res.status === 404) {
      return { error: new errorsMessage().sqlNoRow };
    }
    if (!res.ok) {
      throw new Error(`HTTP error! Status: ${res.status}`);
    }
    const data = await res.json();
    return { data };
  } catch (err: any) {
    return { error: err.message };
  }
}

async function listDeployments(id: string, fetch: (input: RequestInfo | URL, init?: RequestInit) => Promise<Response>) {
  const req = new Request(`${API_URL}/apps/${id}/deployment`, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
    },
  });

  try {
    const res = await fetch(req);
    if (res.status === 404) {
      return { error: new errorsMessage().sqlNoRow };
    }
    if (!res.ok) {
      throw new Error(`HTTP error! Status: ${res.status}`);
    }
    const data = await res.json();
    return { data };
  } catch (err: any) {
    return { error: err.message };
  }
}

async function getApp(id: string, fetch: (input: RequestInfo | URL, init?: RequestInit) => Promise<Response>) {
  const req = new Request(`${API_URL}/apps/${id}`, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
    },
  });

  try {
    const res = await fetch(req);
    if (!res.ok) {
      throw new Error(`HTTP error! Status: ${res.status}`);
    }
    const data = await res.json();
    return { data };
  } catch (err: any) {
    return { error: err.message };
  }
}

export const load: LayoutServerLoad = async ({ fetch, params }) => {
  const { id } = params;

  const appPromise = getApp(id, fetch);
  const activeDeploymentPromise = getActiveDeployment(id, fetch);
  const deploymentsPromise = listDeployments(id, fetch);
  const [app, activeDeployment, deployments] = await Promise.all([
    appPromise,
    activeDeploymentPromise,
    deploymentsPromise
  ]);
  return {
    app: app.data,
    activeDeployment: activeDeployment.data,
    deployments: deployments.data,
    errors: {
      app: app.error,
      activeDeployment: activeDeployment.error,
      deployments: deployments.error,
    }
  };
};

