<script lang="ts">
	import SpaceLayout from '@components/SpaceLayout.svelte';
	import '../../app.postcss';
	import LoggedInHeader from '@components/LoggedInHeader.svelte';
	import Footer from '@components/Footer.svelte';
	import WelcomeModal from '@features/user/components/welcome-modal.svelte';

	export let data;
	let drawer: HTMLInputElement;

	// close drawer when clicking outside

	//close drawer when on mobile
	const closeDrawer = () => {
		if (window.innerWidth > 1024) return;
		const drawer = document.querySelector('#my-drawer');
		drawer.checked = false;
	};
</script>

<div>
	<div class="drawer lg:drawer-open">
		<input bind:this={drawer} id="my-drawer" type="checkbox" class="drawer-toggle" />
		<div on:click={closeDrawer} class="drawer-content lg:pl-60">
			<LoggedInHeader />
			<SpaceLayout>
				<slot />
			</SpaceLayout>
			<Footer />
		</div>
		<div class="drawer-side text-base-100">
			<div class="p-4 min-h-full bg-primary fixed">
				<label for="my-drawer col-start-1" aria-label="close sidebar" class="drawer-overlay" />

				<div class="menu w-52">
					<p class="text-xl pb-8">⛷️ Skiyeti</p>
					<ul>
						<!-- Sidebar content here -->
						<li><a on:click={closeDrawer} href="/diary">Diary</a></li>
						<li><a on:click={closeDrawer} href="/stats">Stats</a></li>
						<li>
							<a on:click={closeDrawer} href="/coolfeatures">Some other really cool features</a>
						</li>
						<li><a on:click={closeDrawer} href="/settings">Settings</a></li>
					</ul>
				</div>
				<img
					class="absolute bottom-1 left-0 w-24"
					src="/branding/strava/powered-by-strava.png"
					alt="powered by strava"
				/>
			</div>
		</div>
	</div>
</div>
{#if data.user && data.user.isFirstLogin}
	<WelcomeModal user={data.user} />
{/if}
