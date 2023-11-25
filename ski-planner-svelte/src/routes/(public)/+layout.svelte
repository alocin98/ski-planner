<script>
	import '../../app.postcss';
	import Header from '@components/Header.svelte';
</script>

<Header />
<section class="bg-white dark:bg-gray-900">
	<div class="max-w-screen-xl px-4 py-8 mx-auto">
		<slot />
	</div>
</section>
