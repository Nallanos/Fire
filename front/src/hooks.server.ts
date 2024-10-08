import { parse } from 'cookie';
import type { HandleFetch } from '@sveltejs/kit';

export const handleFetch: HandleFetch = async ({ event, request, fetch }) => {
    const cookies = parse(event.request.headers.get('cookie') || '');
    const token = cookies['token'];
    if (token) {
        request.headers.set('Authorization', `${token}`);
    }
    return fetch(request);

};