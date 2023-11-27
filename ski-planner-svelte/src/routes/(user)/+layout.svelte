<script>
	import SpaceLayout from '@components/SpaceLayout.svelte';
	import '../../app.postcss';
	import LoggedInHeader from '@components/LoggedInHeader.svelte';

	export let data;
</script>

<div class="grid grid-rows-applayout">
	<div class="drawer lg:drawer-open">
		<input id="my-drawer" type="checkbox" class="drawer-toggle" />
		<div class="drawer-content col-start-2">
			<LoggedInHeader />
			<SpaceLayout>
				<slot />
			</SpaceLayout>
		</div>
		<div class="drawer-side text-base-100">
			<div class="p-4 min-h-full bg-primary">
				<label for="my-drawer col-start-1" aria-label="close sidebar" class="drawer-overlay" />

				<div class="menu w-52">
					<label for="my-drawer" tabindex="0" class="btn btn-ghost md:hidden">
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
					<p class="text-xl pb-8">⛷️ Skiyeti</p>
					<ul>
						<!-- Sidebar content here -->
						<li><a href="/strava">Strava</a></li>
						<li><a href="/trainings">Trainings</a></li>
						<li><a href="/planner">Maybe a training planner</a></li>
						<li><a href="/coolfeatures">Some other really cool features</a></li>
						<li><a href="/settings">Settings</a></li>
					</ul>
				</div>
			</div>
		</div>
	</div>
</div>
