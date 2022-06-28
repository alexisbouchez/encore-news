import type { RequestHandler } from '@sveltejs/kit';
import Client, { Local } from '../encore';

export const post: RequestHandler = async ({ request }) => {
	const formData = await request.formData();

	const client = new Client(Local);
	const news = await client.news.Submit({
		Title: formData.get('title') as string,
		URL: formData.get('url') as string,
		Text: formData.get('text') as string
	});

	return { headers: { Location: '/' }, status: 302 };
};
