import { API_URL } from "../index";

export async function getApp(id: string) {
  const req = new Request(`${API_URL}/apps/${id}`, {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
    },
  });
  try {
    const res = await fetch(req);
    if (!res.ok) {
      throw new Error(`HTTP error Status: ${res.status} at ${res.url}`);
    }
    const data = await res.json();
    return { data: data };
  } catch (err: any) {
    return { error: err.message };
  }
}
