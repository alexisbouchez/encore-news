import type { RequestHandler } from '@sveltejs/kit';
import type { news } from '../encore';
import Client, { Local } from '../encore';

type Params = { id: string };
type Output = { news: news.News };

export const get: RequestHandler<Params, Output> = async ({ params }) => {
	try {
		const client = new Client(Local);
		const news = await client.news.Get(params.id);

		if (news.URL) {
			return { headers: { Location: news.URL }, status: 302 };
		}

		return { body: { news } };
	} catch {
		return { status: 400 };
	}
};
