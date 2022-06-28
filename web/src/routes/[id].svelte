<script lang="ts">
	import type { news } from '../encore';

	export let news: news.News;
</script>

<h1>
	{news.Title}
</h1>

{news.Text}

{news.URL}
