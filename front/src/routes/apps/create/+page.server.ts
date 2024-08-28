import { API_URL } from 'lib/api/index.js';
import { redirect } from '@sveltejs/kit';
export const actions = {
    createApp: async ({ request, fetch }) => {
        const data = await request.formData();
        const appName = data.get('appName') as string;
        if (!appName) {
            console.error(`Impossible to createApp with ${appName} appName`);
            throw new Error(`Invalid appName: ${appName}`);
        } else {
            const req = new Request(`${API_URL}/apps`, {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify({ name: appName }),
            });
            try {
                const res = await fetch(req);
                if (res.ok) {
                    throw redirect(303, "/apps");
                }
            } catch (err) {
                console.error(err);
                throw err;
            }
        }
    }
};