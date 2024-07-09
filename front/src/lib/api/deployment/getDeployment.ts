import { API_URL } from "../index";
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
