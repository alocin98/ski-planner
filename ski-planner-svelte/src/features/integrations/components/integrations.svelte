<script>
	import { StravaService } from '../strava/strava.service';

	export const onClickAuthorize = () => {};
	export let user;

	export let isConnectedToStrava = user?.stravaInfo?.stravaConnected;
	export let isConnectedToPolar = user?.polarInfo?.polarConnected;
	export let isConnectedToGarmin = user?.garminInfo?.garminConnected;

	export const isConnected = isConnectedToStrava || isConnectedToPolar || isConnectedToGarmin;

	const authorizeStrava = () => {
		let authRedirectUrl = window.location.href;
		StravaService.authorize(authRedirectUrl);
	};
</script>

<table class="table w-8/12">
	<!-- head -->

	<tbody>
		<!-- row 1 -->
		<tr>
			<th>Strava</th>
			<td class="flex justify-end">
				<button on:click={authorizeStrava}
					><img src="branding/strava/btn-connect-with-strava.png" alt="strava-authorize" /></button
				>
			</td>
		</tr>
		<!-- row 2 -->
		<tr>
			<th>Polar</th>
			<td class="flex justify-end">
				<button class="btn btn-primary w-48" disabled>Comming soon...</button>
			</td>
		</tr>
		<!-- row 3 -->
		<tr>
			<th>Garmin</th>
			<td class="flex justify-end">
				<button class="btn btn-primary w-48" disabled>Comming soon...</button>
			</td>
		</tr>
	</tbody>
</table>
