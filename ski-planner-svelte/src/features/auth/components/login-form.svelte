<script lang="ts">
	import Loader from '@components/loader.svelte';
	import { t } from '@providers/language';
	import { UserService } from '../service/user-service';

	export let afterLogin: (user: any) => void | Promise<void>;
	export let registerUrl = '/register';

	const data = { email: '', password: '' };
	let error: string | null = null;
	let loadingLogin = false;

	function login() {
		loadingLogin = true;
		UserService.loginWithEmailAndPassword(data.email, data.password)
			.then((user) => {
				afterLogin(user);
			})
			.catch((e: Error) => {
				error = e.message;
				console.error(e.message, e.cause);
			})
			.finally(() => {
				loadingLogin = false;
			});
	}
</script>

<div class="card max-w-96 bg-base-100 shadow-xl">
	<form class="card-body">
		<div class="form-control">
			<label for="email-input" class="label">
				<span class="label-text">{t('login.email')}</span>
			</label>
			<input
				id="email-input"
				bind:value={data.email}
				type="email"
				autocomplete="email"
				placeholder="email"
				class="input input-bordered"
			/>
		</div>
		<div class="form-control">
			<label for="password-input" class="label">
				<span class="label-text">{t('login.password')}</span>
			</label>
			<input
				id="password-input"
				bind:value={data.password}
				type="password"
				placeholder="password"
				class="input input-bordered"
			/>
		</div>
		{#if error !== null}
			<p class="text-error">{error}</p>
		{/if}
		<div class="form-control mt-2 gap-2">
			<button on:submit={login} on:click={login} class="btn btn-primary">
				{#if loadingLogin}
					<Loader className="fill-current" size={30} />
				{:else}
					{t('login.loginButton')}
				{/if}
			</button>
			<p>{t('login.noAccountYet')}</p>
			<a class="link" href={registerUrl}>
				{t('register')}
			</a>
		</div>
	</form>
</div>
