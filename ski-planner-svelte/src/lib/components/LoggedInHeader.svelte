<script>
	import { goto } from '$app/navigation';
	import Cookies from 'js-cookie';

	const logout = () => {
		Cookies.remove('AccessToken');
		goto('/');
	};
</script>

<div class="navbar bg-base-100 h-10">
	<div class="navbar-start">
		<div class="dropdown">
			<!-- svelte-ignore a11y-no-noninteractive-tabindex -->
			<!-- svelte-ignore a11y-label-has-associated-control -->
			<label for="my-drawer" tabindex="0" class="btn btn-ghost lg:hidden">
				<svg
					xmlns="http://www.w3.org/2000/svg"
					class="h-5 w-5"
					fill="none"
					viewBox="0 0 24 24"
					stroke="currentColor"
					><path
						stroke-linecap="round"
						stroke-linejoin="round"
						stroke-width="2"
						d="M4 6h16M4 12h8m-8 6h16"
					/></svg
				>
			</label>
		</div>
	</div>
	<div class="navbar-end">
		<button on:click={logout}>Logout</button>
	</div>
</div>
