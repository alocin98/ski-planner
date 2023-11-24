<script>
	import { t } from '@providers/language';

	export let label = '';
	export let optional = false;
	export let value = '';

	if (optional) label += ' (' + t('optional') + ')';
</script>

<label for="price-input" class="label">
	<span class="label-text">
		{label}
	</span>
</label>
<input
	type="text"
	class="bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
	bind:value
	{...$$restProps}
/>
