import { API_URL } from "..";
import type { User } from "../user";
type getUserParams = {
    token: string;
};
export async function getUserByToken(token: getUserParams): Promise<User | undefined> {
    const req = new Request(API_URL + "/user", {
        method: "GET",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(token),
    })
    try {
        const res = await fetch(req);
        const data: User = await res.json();
        return data
    } catch (err) {
        console.error(err);
    }
} 