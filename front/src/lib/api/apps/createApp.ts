import { API_URL } from "../index";
type CreateAppParams = {
  name: string;
};

export async function createApp(app: CreateAppParams) {
  if (app.name == "" || app.name == undefined || app.name == null) {
    console.error(`Impossible to createApp with ${app.name} appName`);
  } else {
    const req = new Request(`${API_URL}/apps`, {
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
}
