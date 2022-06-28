<script lang="ts">
	import type { news } from '../encore';

	export let news: news.News[];
</script>

<ul>
	{#each news as n}
		<li>
			<a href={n.ID}>
				{n.Title}
			</a>
		</li>
	{/each}
</ul>
