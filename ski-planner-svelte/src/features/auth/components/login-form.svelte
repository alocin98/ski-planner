<script lang="ts">
	import Loader from '@components/loader.svelte';
	import { t } from '@root/src/providers/language';
	import { UserService } from '../service/user-service';

	export let afterLogin: (user: any) => void | Promise<void>;
	export let afterRegister: (user: any) => void | Promise<void>;

	const data = { email: '', password: '' };
	let error: string | null = null;
	let loadingLogin = false;
	let loadingRegister = false;

	function login() {
		loadingLogin = true;
		UserService.loginWithEmailndPassword(data.email, data.password)
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

	function register() {
		loadingRegister = true;
		UserService.newUserWithEmailAndPassword(data.email, data.password)
			.catch((e: any) => {
				console.log('STH WENT WRONG');
				error = e.code;
			})
			.then((user) => {
				if (user) afterRegister(user);
			})

			.finally(() => {
				loadingRegister = false;
			});
	}
</script>

<form class="card-body">
	<div class="form-control">
		<label for="email-input" class="label">
			<span class="label-text">{t('mail')}</span>
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
			<span class="label-text">{t('password')}</span>
		</label>
		<input
			id="password-input"
			bind:value={data.password}
			type="password"
			placeholder="password"
			class="input input-bordered"
		/>
	</div>
	<div class="form-control mt-6 gap-2">
		<button on:submit={login} on:click={login} class="btn btn-primary">
			{#if loadingLogin}
				<Loader className="fill-current" size={30} />
			{:else}
				{t('login')}
			{/if}
		</button>
		<p>{t('login-form.noAccountYet')}</p>
		<button on:click={register} class="btn btn-secondary">
			{#if loadingRegister}
				<Loader className="fill-current" size={30} />
			{:else}
				{t('register')}
			{/if}</button
		>
	</div>

	{#if error !== null}
		<p>{error}</p>
		<p>{t('firebaseErrors.' + error)}</p>
	{/if}
</form>
