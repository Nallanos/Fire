const API_URL = "http://localhost:8080";

export type App = {
  id: string;
  name: string;
  image: string;
  port: number;
  status: string;
};
export type finished_at = {
  Time: string;
  Valid: boolean;
};
export type Deployment = {
  id: string;
  app_id: string;
  status: string;
  created_at: string;
  finished_at: finished_at;
};
type CreateAppPayload = {
  name: string;
};

export async function createApp(app: CreateAppPayload) {
  if (app.name == "" || app.name == undefined || app.name == null) {
    console.error(`Impossible to createApp with ${app.name} appName`);
  }
  const req = new Request(API_URL + "/apps", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(app),
  });

  try {
    const res = await fetch(req);
    console.log(res.status);
    const data = await res.json();
    console.log(data);
  } catch (err) {
    console.error(err);
  }
}

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

export async function getApp(id: string) {
  const req = new Request(API_URL + "/apps/" + id, {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
    },
  });

  const res = await fetch(req);
  const data = await res.json();
  return data;
}

export async function deleteApp(id: string) {
  const req = new Request(API_URL + "/apps/" + id, {
    method: "DELETE",
    headers: {
      "Content-Type": "application/json",
    },
  });
  try {
    const res = await fetch(req);
    console.log(res.status);
  } catch (err) {
    console.error(err);
  }
}

export async function deployApp(id: string) {
  const req = new Request(API_URL + `/apps/${id}/deploy`, {
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
}

export async function stopContainer(id: string) {
  const req = new Request(API_URL + `/apps/${id}/stop`, {
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
}

export async function listDeployments(id: string) {
  const req = new Request(API_URL + `/apps/${id}/deployment`, {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
    },
  });

  const res = await fetch(req);
  const data = await res.json();
  return data;
}

export async function getActiveDeployment(id: string) {
  const req = new Request(API_URL + `/apps/${id}/deployment/activeDeployment`, {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
    },
  });

  try {
    const res = await fetch(req);
    const data = await res.json();
    return data;
  } catch (err) {
    throw err;
  }
}
export async function StartContainer(id: string) {
  const req = new Request(API_URL + `/apps/${id}/start`, {
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
}

export async function getDeployment(id: string, deploymentId: string) {
  const req = new Request(
    API_URL + `/apps/${id}/getDeployment/${deploymentId}`,
    {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
      },
    }
  );

  try {
    const res = await fetch(req);
    const data = await res.json();
    return data;
  } catch (err) {
    throw err;
  }
}
