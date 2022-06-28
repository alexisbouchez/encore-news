import type { RequestHandler } from '@sveltejs/kit';
import type { news } from '../encore';
import Client, { Local } from '../encore';

type Params = { id: string };
type Output = { news: news.News[] };

export const get: RequestHandler<Params, Output> = async ({ params }) => {
	try {
		const client = new Client(Local);
		const result = await client.news.Query();

		let { News: news } = result;

		if (!news) {
			news = [];
		}

		return { body: { news } };
	} catch {
		return { status: 500 };
	}
};
