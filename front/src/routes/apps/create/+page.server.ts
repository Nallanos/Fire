import { API_URL } from 'lib/api/index.js';
import { redirect } from '@sveltejs/kit';
import type { PageServerLoad } from "./$types";
export const actions = {
    createApp: async ({ request, fetch }) => {
        const data = await request.formData();
        const appName = data.get('appName') as string;
        const linkAppImage = data.get('imageName') as string;
        if (!appName) {
            console.error(`Impossible to createApp with ${appName} appName`);
            throw new Error(`Invalid appName: ${appName}`);
        } else {
            const req = new Request(`${API_URL}/apps`, {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify({ name: appName, image: linkAppImage }),
            });
            try {
                const res = await fetch(req);
                console.log("côté serveur", res)
                if (res.ok) {
                    throw redirect(303, "/apps");
                }
            } catch (err) {
                console.error(err);
                throw err;
            }
        }
    },
    searchDockerHubRepos: async ({ fetch, request }) => {
        const body = await request.formData();
        const searchTerm = body.get('searchTerm') as string;
        const apiUrl = `https://hub.docker.com/v2/search/repositories/?query=${searchTerm}`;

        try {
            const response = await fetch(apiUrl, {
                method: "GET",
                headers: {
                    "Content-Type": "application/json",
                },
            });

            if (!response.ok) {
                throw new Error(`Error fetching data: ${response.statusText}`);
            }
            const data = await response.json();
            return JSON.stringify(data);
        } catch (error) {
            console.error('Error fetching data:', error);
            return { images: [], error: 'Failed to fetch data from DockerHub.' };
        }
    }
};

export const load: PageServerLoad = async ({ fetch }) => {
    const url = `https://hub.docker.com/v2/repositories/library/`;
    try {
        const response = await fetch(url, {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json',
            },
        });

        if (!response.ok) {
            throw new Error(`Error fetching repositories: ${response.statusText}`);
        }
        const images = await response.json();
        return images;
    } catch (error) {
        console.error('Error fetching repositories:', error);
        return null;
    }
};
