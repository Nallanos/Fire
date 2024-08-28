import { fail, redirect, type Actions } from '@sveltejs/kit';
import { Login } from '$lib/api/auth';

export const actions: Actions = {
    default: async ({ request, cookies }) => {
        const data = await request.formData();
        const email = data.get('email') as string;
        const password = data.get('password') as string;
        const username = data.get('username') as string;

        if (!email || !password || !username) {
            return fail(400, { missing: true });
        }
        const loginResult = await Login({ email, password, name: username });
        if (loginResult.success && loginResult.data != undefined) {
            cookies.set('token', loginResult.data, {
                path: '/',
                httpOnly: true,
                sameSite: 'strict',
                secure: process.env.NODE_ENV === 'production',
                maxAge: 60 * 60 * 24 * 7
            });
            return redirect(303, '/apps');
        } else {
            return fail(400, { credentials: true });
        }
    }
};