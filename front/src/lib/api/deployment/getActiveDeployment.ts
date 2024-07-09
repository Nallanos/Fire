import { API_URL } from "../index";
import { errorsMessage } from "../error";

export async function getActiveDeployment(id: string) {
  const req = new Request(API_URL + `/apps/${id}/deployment/activeDeployment`, {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
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
