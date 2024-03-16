const API_URL = "http://localhost:8080";

export type App = {
  id: string;
  name: string;
  image: string;
  port: number;
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
