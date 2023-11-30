<script lang="ts">
	import Loader from '@components/loader.svelte';
	import { t } from '@providers/language';
	import { UserService } from '../service/user-service';

	export let afterRegister: (user: any) => void | Promise<void>;
	export let loginUrl = '/login';

	const data = { email: '', password: '' };
	let error: string | null = null;
	let loadingRegister = false;

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

<div class="card max-w-96 bg-base-100 shadow-xl">
	<form class="w-full card-body">
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
			<button on:click={register} class="btn btn-secondary">
				{#if loadingRegister}
					<Loader className="fill-current" size={30} />
				{:else}
					{t('register')}
				{/if}</button
			>
			<p>{t('login.alreadyHaveAnAccount')}</p>
			<a href={loginUrl} class="link">
				{t('login.loginLink')}
			</a>
		</div>
	</form>
</div>
