import { API_URL } from "../index";

export async function deleteApp(id: string) {
  const req = new Request(`${API_URL}/apps/${id}`, {
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
